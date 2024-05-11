const PongGame = artifacts.require("PongGame");

contract("PongGame", accounts => {
  let pongGame;

  beforeEach(async () => {
    pongGame = await PongGame.new(); // Deploy a new instance before each test
  });

  function logEvents(transaction) {
    transaction.logs.forEach(log => {
      console.log(`Event: ${log.event}`);
      console.log('Arguments:', log.args);
    });
  }

  describe("Game creation and participation", () => {
    it("should allow a player to lock a bet and create a game", async () => {
      let betAmount = web3.utils.toWei("0.1", "ether");
      let result = await pongGame.lockOrJoinGame(betAmount, { from: accounts[0], value: betAmount });

      assert.equal(result.logs[0].event, "GameCreated", "Game should be created");
      let gameInfo = await pongGame.games(1);
      assert.equal(gameInfo.player1, accounts[0], "Player 1 should be set correctly");
      assert.equal(gameInfo.betAmount.toString(), betAmount, "Bet amount should be recorded correctly");
    });

    it("should allow a second player to join a game", async () => {
      // First, create a game
      let betAmount = web3.utils.toWei("0.1", "ether");
      await pongGame.lockOrJoinGame(betAmount, { from: accounts[0], value: betAmount });

      // Now, let another player join
      let result = await pongGame.lockOrJoinGame(betAmount, { from: accounts[1], value: betAmount });
      assert.equal(result.logs[0].event, "PlayerJoined", "Player should be able to join");
      let gameInfo = await pongGame.games(1);
      assert.equal(gameInfo.player2, accounts[1], "Player 2 should be set correctly");
    });
  });

  describe("Game completion and fund claiming", () => {
    it("should allow setting a winner and claiming funds", async () => {
      // Assume game 1 is created and both players have joined in the setup of this test
      let betAmount = web3.utils.toWei("0.1", "ether");
      await pongGame.lockOrJoinGame(betAmount, { from: accounts[0], value: betAmount });
      await pongGame.lockOrJoinGame(betAmount, { from: accounts[1], value: betAmount });

      // Set winner
      await pongGame.setWinner(1, accounts[0], { from: accounts[0] });
      let winnerInitialBalance = await web3.eth.getBalance(accounts[0]);

      // Winner claims funds
      let claimResult = await pongGame.claimFunds({ from: accounts[0] });
      assert.equal(claimResult.logs[0].event, "FundsClaimed", "Funds should be claimed by the winner");

      let winnerNewBalance = await web3.eth.getBalance(accounts[0]);
      assert.isTrue(new web3.utils.BN(winnerNewBalance).gt(new web3.utils.BN(winnerInitialBalance)), "Winner's balance should increase");
    });
  });

  describe("Game uncompleted and unlocking funds", () => {
    it("should allow funds to be unlocked if the game is not completed", async () => {
      let betAmount = web3.utils.toWei("0.1", "ether");
      await pongGame.lockOrJoinGame(betAmount, { from: accounts[0], value: betAmount });
      await pongGame.lockOrJoinGame(betAmount, { from: accounts[1], value: betAmount });

      // Attempt to unlock funds before the game is set as complete
      let unlockResult = await pongGame.unlockFunds(1, { from: accounts[0] });
      assert.equal(unlockResult.logs[0].event, "FundsUnlocked", "Funds should be unlocked");

      let gameInfo = await pongGame.games(1);
      assert.equal(gameInfo.fundsLocked, false, "Funds should be marked as unlocked");
      assert.equal(gameInfo.isComplete, false, "Game should not be marked as complete");
    });
  });
});
