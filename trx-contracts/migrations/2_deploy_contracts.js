var txJobCreator = artifacts.require("txJobCreator");

module.exports = function(deployer) {
  deployer.then(async () => {
    await deployer.deploy(txJobCreator);
  });
};