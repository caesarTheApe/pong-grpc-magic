// SPDX-License-Identifier: MIT
pragma solidity >=0.8.24 <0.9.0;

contract PongGame {
    struct Game {
        address player1;
        address player2;
        uint betAmount;
        address winner;
        bool isComplete;
        bool fundsLocked;
    }

    uint public gameIdCounter = 1;
    mapping(uint => Game) public games;
    mapping(address => uint) public withdrawals;

    event GameCreated(uint gameId, address player1, uint betAmount);
    event PlayerJoined(uint gameId, address player2);
    event GameCompleted(uint gameId, address winner);
    event FundsUnlocked(uint gameId);
    event FundsClaimed(address claimant, uint amount);

    // Function to lock funds as the first player or join as the second player
    function lockOrJoinGame(uint _betAmount) external payable {
        require(msg.value == _betAmount, "Incorrect bet amount");

        uint gameId = findOpenGame(_betAmount);
        if (gameId != 0) {
            Game storage game = games[gameId];
            require(game.player2 == address(0), "Game already has two players");
            game.player2 = msg.sender;
            emit PlayerJoined(gameId, msg.sender);
        } else {
            gameId = gameIdCounter++;
            games[gameId] = Game({
                player1: msg.sender,
                player2: address(0),
                betAmount: _betAmount,
                winner: address(0),
                isComplete: false,
                fundsLocked: true
            });
            emit GameCreated(gameId, msg.sender, _betAmount);
        }
    }

    function findOpenGame(uint _betAmount) private view returns (uint) {
        for (uint i = 1; i < gameIdCounter; i++) {
            if (games[i].player2 == address(0) && games[i].betAmount == _betAmount && !games[i].isComplete) {
                return i;
            }
        }
        return 0;
    }

    // Function to set the winner, should be called by the game server
    function setWinner(uint _gameId, address _winner) external {
        Game storage game = games[_gameId];
        require(!game.isComplete, "Game is already complete");
        require(msg.sender == game.player1 || msg.sender == game.player2, "Only game participants can set the winner");

        game.winner = _winner;
        game.isComplete = true;
        withdrawals[_winner] += game.betAmount * 2;

        emit GameCompleted(_gameId, _winner);
    }

    // Function to unlock funds if a game is aborted or a player leaves
    function unlockFunds(uint _gameId) external {
        Game storage game = games[_gameId];
        require(msg.sender == game.player1 || msg.sender == game.player2, "Only game participants can unlock funds");
        require(!game.isComplete, "Game is already complete");
        require(game.fundsLocked, "Funds are not locked");

        if (game.player1 != address(0)) {
            withdrawals[game.player1] += game.betAmount;
        }
        if (game.player2 != address(0)) {
            withdrawals[game.player2] += game.betAmount;
        }

        game.fundsLocked = false;
        emit FundsUnlocked(_gameId);
    }

    function claimFunds() external {
        uint amount = withdrawals[msg.sender];
        require(amount > 0, "No funds available to withdraw");

        withdrawals[msg.sender] = 0;
        payable(msg.sender).transfer(amount);
        emit FundsClaimed(msg.sender, amount);
    }
}
