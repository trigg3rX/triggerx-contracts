var txJobCreator = artifacts.require("txJobCreator");

module.exports = function(deployer) {
  deployer.deploy(txJobCreator);
};