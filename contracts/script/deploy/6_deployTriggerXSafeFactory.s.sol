// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Script} from "forge-std/Script.sol";
import {console} from "forge-std/console.sol";
import {TriggerXSafeFactory} from "../../src/TriggerXSafeFactory.sol";
import {CREATE3} from "@solady/utils/CREATE3.sol";

contract DeployTriggerXSafeFactory is Script {
    uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
    address deployer = vm.addr(deployerPrivateKey);

    address SAFE_PROXY_FACTORY = vm.envAddress("SAFE_PROXY_FACTORY");
    address SAFE_SINGLETON = vm.envAddress("SAFE_SINGLETON");
    address SAFE_L2_SINGLETON = vm.envAddress("SAFE_L2_SINGLETON");

    bytes32 SALT = keccak256(abi.encodePacked("put_salt_here"));

    function run() public {
        // Create fork for this chain
        vm.createSelectFork(vm.envString("ETH_RPC"));

        vm.startBroadcast(deployerPrivateKey);
        
        // Deploy factory with Safe singleton (for L1) using CREATE3 for deterministic address
        bytes memory bytecode = abi.encodePacked(
            type(TriggerXSafeFactory).creationCode,
            abi.encode(SAFE_PROXY_FACTORY, SAFE_SINGLETON)
        );
        address factoryAddress = CREATE3.deployDeterministic(bytecode, SALT);
        TriggerXSafeFactory factory = TriggerXSafeFactory(factoryAddress);
        
        vm.stopBroadcast();

        // Print summary
        console.log("");
        console.log("=== TRIGGERX SAFE FACTORY DEPLOYMENT SUMMARY (L1) ===");
        console.log("TriggerX Safe Factory:", address(factory));
        console.log("Deployer:", deployer);
        console.log("");
    }
}

/**
 * @title DeployTriggerXSafeFactoryL2
 * @notice Deploy TriggerXSafeFactory with SafeL2 singleton for L2 networks
 */
contract DeployTriggerXSafeFactoryL2 is Script {
    uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
    address deployer = vm.addr(deployerPrivateKey);

    address SAFE_PROXY_FACTORY = vm.envAddress("SAFE_PROXY_FACTORY");
    address SAFE_L2_SINGLETON = vm.envAddress("SAFE_L2_SINGLETON");

    bytes32 SALT = keccak256(abi.encodePacked("put_salt_here"));

    function run() public {
        // Create fork for this chain
        // vm.createSelectFork(vm.envString("BASE_RPC"));
        // vm.createSelectFork(vm.envString("OP_RPC"));
        vm.createSelectFork(vm.envString("ARB_RPC"));

        vm.startBroadcast(deployerPrivateKey);
        
        // Deploy factory with SafeL2 singleton (for L2) using CREATE3 for deterministic address
        bytes memory bytecode = abi.encodePacked(
            type(TriggerXSafeFactory).creationCode,
            abi.encode(SAFE_PROXY_FACTORY, SAFE_L2_SINGLETON)
        );
        address factoryAddress = CREATE3.deployDeterministic(bytecode, SALT);
        TriggerXSafeFactory factory = TriggerXSafeFactory(factoryAddress);
        
        vm.stopBroadcast();

        // Print summary
        console.log("");
        console.log("=== TRIGGERX SAFE FACTORY DEPLOYMENT SUMMARY (L2) ===");
        console.log("TriggerX Safe Factory:", address(factory));
        console.log("Deployer:", deployer);
        console.log("");
    }
}

/**
 * @title TestTriggerXSafeCreation
 * @notice Test script to create a Safe wallet using the factory
 */
contract TestTriggerXSafeCreation is Script {
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

