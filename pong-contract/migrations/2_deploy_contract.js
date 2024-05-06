const PongGame = artifacts.require("PongGame");

module.exports = function(deployer) {
  deployer.deploy(PongGame);
};
