package main

import (
	"sync"

	"github.com/vctt94/pong-grpc-eth/pongrpc/grpc/pong"
)

type Player struct {
	ID           string
	PlayerNumber int32 // 1 for player 1, 2 for player 2
	stream       pong.PongGame_StreamUpdatesServer
}

func NewPlayer(id string, stream pong.PongGame_StreamUpdatesServer) *Player {
	return &Player{
		ID:     id,
		stream: stream,
	}
}

type PlayerSessions struct {
	mu       sync.Mutex
	sessions map[string]*Player // Map client ID to Player
}

func NewPlayerSessions() *PlayerSessions {
	return &PlayerSessions{
		sessions: make(map[string]*Player),
	}
}

func (ps *PlayerSessions) AddOrUpdatePlayer(player *Player) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ps.sessions[player.ID] = player
}

func (ps *PlayerSessions) RemovePlayer(clientID string) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	delete(ps.sessions, clientID)
}

func (ps *PlayerSessions) GetPlayer(clientID string) (*Player, bool) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	player, exists := ps.sessions[clientID]
	return player, exists
}

type WaitingRoom struct {
	mu    sync.Mutex
	queue []*Player
}

func NewWaitingRoom() *WaitingRoom {
	return &WaitingRoom{
		queue: make([]*Player, 0),
	}
}

func (wr *WaitingRoom) AddPlayer(player *Player) {
	wr.mu.Lock()
	defer wr.mu.Unlock()
	wr.queue = append(wr.queue, player)
}

func (wr *WaitingRoom) ReadyPlayers() ([]*Player, bool) {
	wr.mu.Lock()
	defer wr.mu.Unlock()
	if len(wr.queue) >= 2 {
		players := wr.queue[:2]
		wr.queue = wr.queue[2:]
		return players, true
	}
	return nil, false
}
