var TriggerXJobCreator = artifacts.require("TriggerXJobCreator");

module.exports = function(deployer) {
  deployer.deploy(TriggerXJobCreator);
};