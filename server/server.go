package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/big"
	"net"
	"os"
	"sync"

	"github.com/vctt94/pong-grpc-eth/pongrpc/grpc/pong"

	canvas "github.com/vctt94/pong-grpc-eth/pong"

	PongGame "github.com/vctt94/pong-grpc-eth/contract-go"

	"github.com/decred/slog"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/ndabAP/ping-pong/engine"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func init() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
}

var (
	serverLogger = log.New(os.Stdout, "[SERVER] ", 0)
	debug        = flag.Bool("debug", false, "")
	fps          = flag.Uint("fps", canvas.DEFAULT_FPS, "")
)

type GameStartNotification struct {
	GameID  string
	Players []*Player
}

type gameInstance struct {
	engine   *canvas.CanvasEngine
	framesch chan []byte
	inputch  chan []byte
	players  []*Player
}

type server struct {
	pong.UnimplementedPongGameServer
	ID                 string
	mu                 sync.Mutex
	clientReady        chan string              // Channel to signal a client is ready
	gameStartNotifiers map[string]chan string   // channel to signal to players a game started
	games              map[string]*gameInstance // Map to hold game instances, indexed by a game ID
	waitingRoom        *WaitingRoom
	playerSessions     *PlayerSessions
	ethClient          *ethclient.Client
	contract           *PongGame.PongGameTransactor
}

func setupEthereumClient() (*ethclient.Client, error) {
	// Fetch the Infura project ID from environment variables
	apiKey := os.Getenv("ROCKX_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("ROCKX_API_KEY is not set in the environment")
	}

	// Use the project ID in the Ethereum client dial
	client, err := ethclient.Dial(fmt.Sprintf("https://sepolia-eth.w3node.com/%s", apiKey))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum client: %v", err)
	}

	return client, nil
}

func (s *server) SendInput(ctx context.Context, in *pong.PlayerInput) (*pong.GameUpdate, error) {
	// Example: Determine client ID and game instance (implementation depends on your client ID strategy)
	clientID, err := getClientIDFromContext(ctx) // Implement this based on your authentication/identification scheme
	if err != nil {
		return nil, err
	}
	gameInstance, player, exists := s.findGameInstanceAndPlayerByClientID(clientID)
	if !exists {
		return nil, fmt.Errorf("game instance not found for client ID %s", clientID)
	}

	in.PlayerNumber = player.PlayerNumber
	inputBytes, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize input: %v", err)
	}
	// Forward input to the correct game instance

	gameInstance.inputch <- inputBytes

	return &pong.GameUpdate{}, nil
}

func (s *server) findPlayer(clientID string) (*Player, bool) {
	for _, player := range s.waitingRoom.queue {
		if player.ID == clientID {
			return player, true
		}
	}
	return nil, false
}

func (s *server) StreamUpdates(req *pong.GameStreamRequest, stream pong.PongGame_StreamUpdatesServer) error {
	clientID, err := getClientIDFromContext(stream.Context())
	if err != nil {
		return err
	}
	ctx := stream.Context()
	s.mu.Lock()
	player, exists := s.playerSessions.GetPlayer(clientID)
	if !exists {
		player = NewPlayer(clientID, stream)
		s.playerSessions.AddOrUpdatePlayer(player)
	} else {
		player.stream = stream
		s.playerSessions.AddOrUpdatePlayer(player)
	}
	s.mu.Unlock()
	// Subscribe to game updates.
	gameInstance, _, exists := s.findGameInstanceAndPlayerByClientID(clientID)
	if !exists {
		return fmt.Errorf("no game instance found for client ID %s", clientID)
	}
	for {
		select {
		case <-ctx.Done():
			s.handleDisconnect(clientID)
			return ctx.Err()
		case frame, ok := <-gameInstance.framesch:
			if !ok {
				return nil // Game has ended or error occurred.
			}
			if err := stream.Send(&pong.GameUpdateBytes{Data: frame}); err != nil {
				s.handleDisconnect(clientID)
				return err
			}
		}
	}
}

func (s *server) handleDisconnect(clientID string) {
	// Handle the client's disconnection by cleaning up resources, notifying other players, etc.
	// Update internal structures to reflect that the client has disconnected.
}

func (s *server) SignalReady(ctx context.Context, req *pong.SignalReadyRequest) (*pong.SignalReadyResponse, error) {
	clientID := req.ClientId
	serverLogger.Printf("SignalReady called by client ID: %s", clientID)

	player, exists := s.playerSessions.GetPlayer(clientID)
	if !exists {
		return nil, fmt.Errorf("player %s has not established a connection", clientID)
	}

	// Mark the player as ready in the waiting room
	s.waitingRoom.AddPlayer(player)
	s.signalClientReady(clientID)

	return &pong.SignalReadyResponse{}, nil
}

func (s *server) findGameInstanceAndPlayerByClientID(clientID string) (*gameInstance, *Player, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, game := range s.games {
		for _, player := range game.players {
			if player.ID == clientID {
				return game, player, true
			}
		}
	}
	return nil, nil, false
}

func newServer(id string) *server {
	return &server{
		ID:                 id,
		clientReady:        make(chan string, 10), // Buffer based on expected simultaneous ready signals
		games:              make(map[string]*gameInstance),
		gameStartNotifiers: make(map[string]chan string),
		waitingRoom:        NewWaitingRoom(),
		playerSessions:     NewPlayerSessions(),
	}
}

func getClientIDFromContext(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	fmt.Printf("md: %+v\n\n", md)
	if !ok {
		return "", fmt.Errorf("no metadata found in context")
	}

	clientIDs, ok := md["client-id"] // Assuming the client ID is passed under the key "client-id"
	if !ok || len(clientIDs) == 0 {
		return "", fmt.Errorf("client-id not found in metadata")
	}

	return clientIDs[0], nil // Return the first client ID from the metadata
}

func (s *server) signalClientReady(clientID string) {
	fmt.Printf("Client %s is ready\n", clientID)
	s.clientReady <- clientID // Simply mark the client as ready
}

func (s *server) manageGames(ctx context.Context) {
	for {
		select {
		case clientID := <-s.clientReady:
			if player, exists := s.findPlayer(clientID); exists && player.stream != nil {
				// Check if enough players are ready to start a game
				if players, ready := s.waitingRoom.ReadyPlayers(); ready {
					s.startGameWithPlayers(ctx, players)
				}
			} else {
				fmt.Printf("Waiting for player %s to initialize stream\n", clientID)
			}
		case <-ctx.Done():
			return
		}
	}
}

func (s *server) getGameStartNotifier(clientID string) (chan string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	notifier, exists := s.gameStartNotifiers[clientID]
	if !exists {
		// If no notifier exists for this client, create one
		notifier = make(chan string, 1) // Buffer of 1 to avoid blocking senders
		s.gameStartNotifiers[clientID] = notifier
		exists = true
	}
	return notifier, exists
}

func (s *server) NotifyGameStarted(req *pong.GameStartedRequest, stream pong.PongGame_NotifyGameStartedServer) error {
	clientID := req.ClientId
	gameStartNotifier, exists := s.getGameStartNotifier(clientID)
	if !exists {
		return fmt.Errorf("no game notifier available for client ID %s", clientID)
	}

	ctx := stream.Context()
	select {
	case gameID := <-gameStartNotifier:
		if err := stream.Send(&pong.GameStartedResponse{Message: "Game has started with ID: " + gameID}); err != nil {
			return err
		}
	case <-ctx.Done():
		return ctx.Err()
	}
	return nil
}

func (s *server) startGameWithPlayers(ctx context.Context, players []*Player) {
	gameID := generateGameID()
	newGameInstance := s.startNewGame(ctx)
	players[0].PlayerNumber = 1
	players[1].PlayerNumber = 2
	newGameInstance.players = players
	s.games[gameID] = newGameInstance

	for _, player := range players {
		if player.stream == nil {
			fmt.Printf("Player %s stream is not initialized\n", player.ID)
		} else {
			fmt.Printf("Notifying player %s game started\n", player.ID)
			notifier, ok := s.gameStartNotifiers[player.ID]
			if ok {
				notifier <- gameID // Assuming each player has a channel listening for their game start events
			}
		}
	}
}

func generateGameID() string {
	return uuid.New().String()
}

func (s *server) startNewGame(ctx context.Context) *gameInstance {
	// Initialize game engine
	game := engine.NewGame(
		800, 400,
		engine.NewPlayer(5, 100),
		engine.NewPlayer(5, 100),
		engine.NewBall(10, 10),
	)

	canvasEngine := canvas.New(game)
	canvasEngine.SetDebug(*debug).SetFPS(*fps)

	// Set up channels for the new game instance
	framesch := make(chan []byte, 100)
	inputch := make(chan []byte, 10)

	// Start the game engine for this instance
	go canvasEngine.NewRound(ctx, framesch, inputch)

	// Create a new game instance and add it to the map
	gameID := generateGameID()
	instance := &gameInstance{
		engine:   canvasEngine,
		framesch: framesch,
		inputch:  inputch,
	}
	s.games[gameID] = instance

	return instance
}

func (s *server) getTransactor(betAmount *big.Int) (*bind.TransactOpts, error) {
	privateKeyString := os.Getenv("ETH_PRIVATE_KEY") // Ensure this environment variable is set securely
	if privateKeyString == "" {
		return nil, fmt.Errorf("ethereum private key not set")
	}

	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %v", err)
	}

	// NewKeyedTransactor is deprecated and will be removed in future releases, use NewKeyedTransactorWithChainID instead.
	chainID, err := s.ethClient.NetworkID(context.Background()) // Make sure to handle this according to your Ethereum network
	if err != nil {
		return nil, fmt.Errorf("failed to get network ID: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create keyed transactor: %v", err)
	}

	auth.Value = betAmount // Set the transaction value to the bet amount

	return auth, nil
}

func (s *server) lockOrJoinGame(playerID string, betAmount *big.Int) error {
	auth, err := s.getTransactor(betAmount)
	if err != nil {
		return err
	}

	// Call the lockOrJoinGame function from the contract
	tx, err := s.contract.LockOrJoinGame(auth, betAmount)
	if err != nil {
		return fmt.Errorf("failed to lock or join game: %v", err)
	}

	fmt.Printf("Lock or Join Game transaction sent: %s\n", tx.Hash().Hex())
	return nil
}

func setupContract(client *ethclient.Client, contractAddress common.Address) (*PongGame.PongGameTransactor, error) {
	// Assuming you have both a transactor and a caller, you might want to choose based on usage context
	contract, err := PongGame.NewPongGameTransactor(contractAddress, client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate contract: %v", err)
	}
	return contract, nil
}

func realMain() error {
	flag.Parse()

	contractAddress := common.HexToAddress("0x653Def72c6b9Fdf6BaA5675542D2a4cB43137563")
	// Initialize and start managing games in the background
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// g, gctx := errgroup.WithContext(ctx)
	bknd := slog.NewBackend(os.Stderr)
	log := bknd.Logger("EXMP")
	log.SetLevel(slog.LevelInfo)

	serverId := uuid.New().String()
	srv := newServer(serverId)

	ethc, err := setupEthereumClient()
	if err != nil {
		log.Errorf("failed to init eth client: %v", err)
		return err
	}

	srv.ethClient = ethc
	contract, err := setupContract(srv.ethClient, contractAddress)
	if err != nil {
		log.Errorf("failed to setup contract: %v", err)
		return err
	}
	srv.contract = contract

	go srv.manageGames(ctx)

	// Set up gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Errorf("failed to listen: %v", err)
		return err
	}
	grpcServer := grpc.NewServer()
	pong.RegisterPongGameServer(grpcServer, srv)
	fmt.Println("server listening at", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Errorf("failed to serve: %v", err)
		return err
	}
	return nil
}

func main() {
	err := realMain()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
