const PongGame = artifacts.require("PongGame");

contract("PongGame", accounts => {
  const [owner, player1, player2] = accounts;
  let pongGame;
  const betAmount = web3.utils.toWei("1", "ether");

  beforeEach(async () => {
    pongGame = await PongGame.new();
  });

  describe("Game Creation", () => {
    it("should emit a GameCreated event when a game is created", async () => {
      const result = await pongGame.createGame(player2, betAmount, { from: player1, value: betAmount });
      assert.equal(result.logs[0].event, "GameCreated", "GameCreated event should be emitted");
    });
  });

  describe("Betting", () => {
    it("should allow player2 to place a bet", async () => {
      await pongGame.createGame(player2, betAmount, { from: player1, value: betAmount });
      const result = await pongGame.placeBet(1, { from: player2, value: betAmount });
      assert.ok(result, "Bet should be placed successfully");
    });
  });

  describe("Winner Declaration", () => {
    it("should payout correctly to the winner", async () => {
      await pongGame.createGame(player2, betAmount, { from: player1, value: betAmount });
      await pongGame.placeBet(1, { from: player2, value: betAmount });
      await pongGame.setWinner(1, player1, { from: player1 });
      let balanceAfter = await web3.eth.getBalance(player1);
      balanceAfter = new web3.utils.BN(balanceAfter);
      assert(balanceAfter.gt(new web3.utils.BN(betAmount)), "Winner should receive the bet amount");
    });
  });
});
