// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import {TriggerXSafeFactory} from "../../src/SafeWalletFactory.sol";

/**
 * @title DeploySafeFactory
 * @notice Script to deploy TriggerXSafeFactory for different networks
 * @dev Usage:
 *      For Ethereum Mainnet:
 *        forge script script/DeploySafeFactory.s.sol:DeploySafeFactory --rpc-url $ETH_RPC_URL --broadcast
 *      
 *      For L2 (Optimism, Arbitrum, Base, etc.):
 *        forge script script/DeploySafeFactory.s.sol:DeploySafeFactoryL2 --rpc-url $L2_RPC_URL --broadcast
 */
contract DeploySafeFactory is Script {
    // Safe Core Contracts (deterministic deployments across all networks)
    address constant SAFE_PROXY_FACTORY = 0x4e1DCf7AD4e460CfD30791CCC4F9c8a4f820ec67;
    
    // Safe v1.4.1 - Optimized for L1
    address constant SAFE_SINGLETON = 0x41675C099F32341bf84BFc5382aF534df5C7461a;
    
    // SafeL2 v1.4.1 - Optimize d for L2
    address constant SAFE_L2_SINGLETON = 0x29fcB43b46531BcA003ddC8FCB67FFE91900C762;
    
    // TriggerX contracts (to be set based on your deployment)
    address taskExecutionHub;

    function setUp() public {
        // Set your TriggerX contract addresses here
        // These should be loaded from environment variables or config
        taskExecutionHub = vm.envOr("TASK_EXECUTION_HUB", 0x179c62e83c3f90981B65bc12176FdFB0f2efAD54);
        
        require(taskExecutionHub != address(0), "TASK_EXECUTION_HUB not set");
    }

    function run() public {
        vm.startBroadcast();
        
        // Deploy factory with Safe singleton (for L1)
        TriggerXSafeFactory factory = new TriggerXSafeFactory(
            taskExecutionHub,
            SAFE_PROXY_FACTORY,
            SAFE_SINGLETON
        ); 
        
        console.log("TriggerXSafeFactory deployed at:", address(factory));
        console.log("Using Safe singleton:", SAFE_SINGLETON);
        console.log("Using SafeProxyFactory:", SAFE_PROXY_FACTORY);
        console.log("TaskExecutionHub:", taskExecutionHub);
        
        vm.stopBroadcast();
    }
}

/**
 * @title DeploySafeFactoryL2
 * @notice Deploy TriggerXSafeFactory with SafeL2 singleton for L2 networks
 */
contract DeploySafeFactoryL2 is Script {
    address constant SAFE_PROXY_FACTORY = 0x4e1DCf7AD4e460CfD30791CCC4F9c8a4f820ec67;
    address constant SAFE_L2_SINGLETON = 0x29fcB43b46531BcA003ddC8FCB67FFE91900C762;
    
    address taskExecutionHub;

    function setUp() public {
        taskExecutionHub = vm.envOr("TASK_EXECUTION_HUB", 0x179c62e83c3f90981B65bc12176FdFB0f2efAD54);
        
        require(taskExecutionHub != address(0), "TASK_EXECUTION_HUB not set");
    }

    function run() public {
        vm.startBroadcast();
        
        // Deploy factory with SafeL2 singleton (for L2) using CREATE2
        bytes32 salt = "TriggerXSafeFactory";
        TriggerXSafeFactory factory = new TriggerXSafeFactory{salt: salt}(
            taskExecutionHub,
            SAFE_PROXY_FACTORY,
            SAFE_L2_SINGLETON
        );
        console.log("TriggerXSafeFactory (L2) deployed at:", address(factory));
        console.log("Using SafeL2 singleton:", SAFE_L2_SINGLETON);
        console.log("Using SafeProxyFactory:", SAFE_PROXY_FACTORY);
        console.log("TaskExecutionHub:", taskExecutionHub);
        
        vm.stopBroadcast();
    }
}

/**
 * @title TestSafeCreation
 * @notice Test script to create a Safe wallet using the factory
 */
contract TestSafeCreation is Script {
    function run() public {
        address factoryAddress = vm.envAddress("SAFE_FACTORY_ADDRESS");
        address userAddress = vm.envAddress("USER_ADDRESS");
        
        TriggerXSafeFactory factory = TriggerXSafeFactory(factoryAddress);
        
        vm.startBroadcast();
        
        // Predict the Safe address before creation
        address predictedAddress = factory.predictSafeAddress(userAddress);
        console.log("Predicted Safe address:", predictedAddress);
        
        // Create the Safe wallet
        address actualAddress = factory.createSafeWallet(userAddress);
        console.log("Actual Safe address:", actualAddress);
        
        // Verify they match
        require(actualAddress == predictedAddress, "Address mismatch!");
        console.log("Success! Addresses match.");
        
        // Check wallet info
        address[] memory wallets = factory.getSafeWallets(userAddress);
        console.log("Total wallets for user:", wallets.length);
        
        uint256 nextNonce = factory.getUserSaltNonce(userAddress);
        console.log("Next saltNonce:", nextNonce);
        
        vm.stopBroadcast();
    }
}

