// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "forge-std/console.sol";
import {CREATE3} from "solady/utils/CREATE3.sol";
import {ProxyHub} from "../src/lz/ProxyHub.sol";
import {ProxySpoke} from "../src/lz/ProxySpoke.sol";
import {Target} from "../src/Target.sol";
import {ILayerZeroEndpointV2} from "@layerzero-v2/contracts/interfaces/ILayerZeroEndpointV2.sol";

contract DeployAll is Script {
    // --- Configuration (Update if needed) ---
    // LayerZero Endpoints
    address constant LZ_ENDPOINT_OP_SEPOLIA = 0x6EDCE65403992e310A62460808c4b910D972f10f;
    address constant LZ_ENDPOINT_BASE_SEPOLIA = 0x6EDCE65403992e310A62460808c4b910D972f10f;

    // Endpoint IDs (EIDs for LayerZero)
    uint32 constant OP_SEPOLIA_EID = 40232; // Optimism Sepolia Endpoint ID
    uint32 constant BASE_SEPOLIA_EID = 40245; // Base Sepolia Endpoint ID

    // Attestation Center on Base Sepolia
    address constant ATTESTATION_CENTER_ADDR = 0x710DAb96f318b16F0fC9962D3466C00275414Ff0;

    // Deployment Salt (Must be the same for Hub and Spoke)
    bytes32 constant SALT = "Account-abstraction";

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployerAddress = vm.addr(deployerPrivateKey);
        console.log("Deploying contracts using deployer:", deployerAddress);

        // --- Deploy Hub on Base Sepolia ---
        console.log("\n=== Deploying Hub on Base Sepolia ===");
        vm.createSelectFork(vm.envString("BASE_SEPOLIA_RPC"));
        vm.startBroadcast(deployerPrivateKey);

        // Encode constructor arguments for Hub
        bytes memory hubConstructorArgs = abi.encode(LZ_ENDPOINT_BASE_SEPOLIA, ATTESTATION_CENTER_ADDR);
        bytes memory hubCreationCode = abi.encodePacked(
            type(ProxyHub).creationCode,
            hubConstructorArgs
        );

        // Deploy Hub using CREATE3
        address deployedHubAddr = CREATE3.deployDeterministic(hubCreationCode, SALT);
        ProxyHub hub = ProxyHub(payable(deployedHubAddr)); // Cast address to contract type

        console.log("ProxyHub deployed at:", deployedHubAddr);
        console.log("Attestation Center used by Hub:", address(hub.attestationCenter()));

        // Send ETH to Hub contract to cover LayerZero fees
        vm.deal(address(hub), 1 ether);
        console.log("Sent 1 ETH to Hub contract at:", address(hub));

        vm.stopBroadcast();

        // --- Deploy Spoke and Target on Optimism Sepolia ---
        console.log("\n=== Deploying Spoke and Target on Optimism Sepolia ===");
        vm.createSelectFork(vm.envString("OPSEPOLIA_RPC"));
        vm.startBroadcast(deployerPrivateKey);

        // Encode constructor arguments for Spoke
        bytes memory spokeConstructorArgs = abi.encode(LZ_ENDPOINT_OP_SEPOLIA, BASE_SEPOLIA_EID);
        bytes memory spokeCreationCode = abi.encodePacked(
            type(ProxySpoke).creationCode,
            spokeConstructorArgs
        );

        // Deploy Spoke using CREATE3
        address deployedSpokeAddr = CREATE3.deployDeterministic(spokeCreationCode, SALT);
        ProxySpoke spoke = ProxySpoke(payable(deployedSpokeAddr)); // Cast address to contract type
        console.log("ProxySpoke deployed at:", deployedSpokeAddr);

        // Deploy Target using standard 'new'
        Target target = new Target();
        address targetAddr = address(target);
        console.log("Target contract deployed at:", targetAddr);

        // Set Hub Address on Spoke
        // !! This function doesn't exist yet in your ProxySpoke contract. You'll need to add it.
        // spoke.setHubAddress(deployedHubAddr);
        // console.log("Set Hub Address on Spoke.");

        vm.stopBroadcast();

        console.log("\n--- Deployment Complete ---");
        console.log("Hub Address:", deployedHubAddr);
        console.log("Spoke Address:", deployedSpokeAddr);
        console.log("Target Address:", targetAddr);
        console.log("---------------------------");
    }
}