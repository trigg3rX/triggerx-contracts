// SPDX-License-Identifier: MIT
pragma solidity >=0.8.25;

import "forge-std/Script.sol";
import "../src/ProxyFactory.sol";

/// @title Deploy ProxyFactory Script
/// @notice Deployment script for ProxyFactory contract
/// @dev Uses environment variables for configuration
contract DeployProxyFactory is Script {
    // Config struct to hold deployment parameters
    struct DeployConfig {
        address avsGovernance;
        uint256 deployerKey;
    }

    /// @dev Loads configuration from environment
    function _loadConfig() internal view returns (DeployConfig memory) {
        DeployConfig memory config;
        
        // Try to load AVS governance from env, otherwise use default for testnet
        bytes memory avsAddr = vm.envOr("AVS_GOVERNANCE_ADDRESS", bytes(""));
        if (avsAddr.length > 0) {
            config.avsGovernance = abi.decode(avsAddr, (address));
        } else {
            // Default Holesky testnet address
            config.avsGovernance = address(0x0C77B6273F4852200b17193837960b2f253518FC);
        }

        // Load deployer private key - will revert if not set
        config.deployerKey = vm.envUint("KEEPER_PRIVATE_KEY");
        
        return config;
    }

    function run() external {
        // Load deployment configuration
        DeployConfig memory config = _loadConfig();
        
        // Log pre-deployment information
        console.log("Deploying ProxyFactory with configuration:");
        console.log("AVS Governance:", config.avsGovernance);
        console.log("Deployer:", vm.addr(config.deployerKey));

        // Start broadcasting transactions
        vm.startBroadcast(config.deployerKey);

        // Deploy ProxyFactory
        ProxyFactory factory = new ProxyFactory(config.avsGovernance);
        
        // Log deployment results
        console.log("\nDeployment Results:");
        console.log("ProxyFactory deployed to:", address(factory));
        console.log("AVS Governance set to:", factory.avsGovernance());

        vm.stopBroadcast();
    }


}