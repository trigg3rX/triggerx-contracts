var TriggerXJobManager = artifacts.require("TriggerXJobManager");

module.exports = function(deployer) {
  deployer.deploy(TriggerXJobManager);
};