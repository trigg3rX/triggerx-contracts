var txJobCreator = artifacts.require("txJobCreator");

module.exports = function(deployer) {
  deployer.then(async () => {
    await deployer.deploy(txJobCreator);

    const receipt = await txJobCreator.deployed();

    console.log("TronLzApp deployed at:", receipt.address);

  });

};
