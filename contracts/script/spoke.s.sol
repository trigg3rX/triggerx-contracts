// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "forge-std/Test.sol";
import "forge-std/console.sol";
import {Vm} from "forge-std/Vm.sol";
import {CREATE3} from "solady/utils/CREATE3.sol";

import {ProxyHub} from "../src/lz/ProxyHub.sol";
import {ProxySpoke} from "../src/lz/ProxySpoke.sol";
import {Target} from "../src/Target.sol";
import {IAttestationCenter} from "../src/interfaces/IAttestationCenter.sol";
import {ILayerZeroEndpointV2, Origin, MessagingParams, MessagingFee} from "@layerzero-v2/contracts/interfaces/ILayerZeroEndpointV2.sol";

contract DeployAndTestFlow is Script, Test {
    // --- Configuration ---
    // Endpoint IDs (EIDs for LayerZero)
    uint32 constant OP_SEPOLIA_EID = 40232; // Optimism Sepolia Endpoint ID
    uint32 constant BASE_SEPOLIA_EID = 40245; // Base Sepolia Endpoint ID (Your Hub Endpoint ID)

    // LayerZero Endpoints
    address constant LZ_ENDPOINT_OP_SEPOLIA = 0x6EDCE65403992e310A62460808c4b910D972f10f; // Optimism Sepolia Endpoint
    address constant LZ_ENDPOINT_BASE_SEPOLIA = 0x6EDCE65403992e310A62460808c4b910D972f10f; // Base Sepolia Endpoint

    // Attestation Center on Base Sepolia
    address constant ATTESTATION_CENTER_ADDR = 0x710DAb96f318b16F0fC9962D3466C00275414Ff0;

    // RPC URLs
    string opSepoliaRpcUrl = vm.envString("OPSEPOLIA_RPC");
    string baseSepoliaRpcUrl = vm.envString("BASE_SEPOLIA_RPC");

    // Keeper Private Key
    uint256 keeperPrivateKey = vm.envUint("KEEPER_PRIVATE_KEY");
    address keeperAddress = vm.addr(keeperPrivateKey);

    // Deployed Contract Instances
    ProxyHub hub;
    ProxySpoke spoke;
    Target target;

    // Fork IDs
    uint256 opForkId;
    uint256 baseForkId;

    function run() external {
        console.log("Starting cross-chain flow test with CREATE3 deployment...");
        console.log("Keeper Address:", keeperAddress);

        // Use the same salt for both deployments to get the same address
        // bytes32 SALT = keccak256(abi.encodePacked("Account-abstraction"));
        bytes32 SALT = "Account-abstraction"; // Use keccak256 for bytes32 salt

        // === 1. Deploy Hub on Base Sepolia Fork using CREATE3 ===
        console.log("\n=== Deploying Hub on Base Sepolia Fork via CREATE3 ===");
        baseForkId = vm.createSelectFork(baseSepoliaRpcUrl);
        console.log("Selected Base Sepolia Fork ID:", baseForkId);

        // Encode constructor arguments for Hub
        bytes memory hubConstructorArgs = abi.encode(LZ_ENDPOINT_BASE_SEPOLIA, ATTESTATION_CENTER_ADDR);
        bytes memory hubCreationCode = abi.encodePacked(
            type(ProxyHub).creationCode,
            hubConstructorArgs
        );

        // Encode constructor arguments for Spoke
        bytes memory spokeConstructorArgs = abi.encode(LZ_ENDPOINT_OP_SEPOLIA, BASE_SEPOLIA_EID);
        bytes memory spokeCreationCode = abi.encodePacked(
            type(ProxySpoke).creationCode,
            spokeConstructorArgs
        );

        vm.startBroadcast();
        // Deploy using CREATE3 with deployDeterministic
        address deployedHubAddr = CREATE3.deployDeterministic(hubCreationCode, SALT);
        hub = ProxyHub(payable(deployedHubAddr)); // Cast address to contract type
        vm.stopBroadcast();
        console.log("Attestation Center used by Hub:", address(hub.attestationCenter()));

        console.log("Hub deployed via CREATE3 at:", address(hub));

        // Send ETH to Hub contract to cover LayerZero fees
        vm.deal(address(hub), 1 ether);
        console.log("Sent 1 ETH to Hub contract at:", address(hub));

        // === 2. Deploy Spoke and Target on Optimism Sepolia Fork ===
        console.log("\n=== Deploying Spoke (CREATE3) & Target (new) on Optimism Sepolia Fork ===");
        opForkId = vm.createSelectFork(opSepoliaRpcUrl);
        console.log("Selected Optimism Sepolia Fork ID:", opForkId);

        vm.startBroadcast();
        // Deploy Spoke using CREATE3 with deployDeterministic
        address deployedSpokeAddr = CREATE3.deployDeterministic(spokeCreationCode, SALT);
        spoke = ProxySpoke(payable(deployedSpokeAddr)); // Cast address to contract type

        // Deploy Target using standard 'new'
        target = new Target();
        vm.stopBroadcast();

        console.log("Spoke deployed via CREATE3 at:", address(spoke));
        console.log("Target deployed at:", address(target));
        console.log("Hub Endpoint ID set in Spoke:", spoke.hubEid());
    
        // === 3. Initiate Call from Keeper on Spoke (Optimism Sepolia) ===
        console.log("\n=== Initiating call from Keeper on Spoke (Op Sepolia) ===");
        vm.selectFork(opForkId);
        vm.startPrank(keeperAddress);

        uint256 valueToSet = 12345;
        bytes memory targetCalldata = abi.encodeWithSignature("setValue", valueToSet);
        console.logBytes(targetCalldata);
        bytes memory expectedMsgToHub = abi.encode(keeperAddress, address(target), targetCalldata);

        // Send with enough value to cover LayerZero fees (0.001 ETH to be safe)
        spoke.executeFunction{value: 1000000000000000}(address(target), targetCalldata);
        console.log("Called spoke.executeFunction as keeper.");
        vm.stopPrank();

        // === 4. Simulate Hub Receiving Message (Base Sepolia) ===
        console.log("\n=== Simulating Hub receiving message (Base Sepolia) ===");
        vm.selectFork(baseForkId);

        Origin memory originFromSpoke = Origin({
            srcEid: OP_SEPOLIA_EID,
            sender: bytes32(uint256(uint160(address(spoke)))),
            nonce: 0
        });

        // Estimate the fee
        MessagingFee memory fee = ILayerZeroEndpointV2(LZ_ENDPOINT_BASE_SEPOLIA).quote(
            MessagingParams({
                dstEid: OP_SEPOLIA_EID,
                receiver: bytes32(uint256(uint160(address(spoke)))),
                message: abi.encode(address(target), targetCalldata),
                options: abi.encodePacked(uint16(1), uint256(1000000)),
                payInLzToken: false
            }),
            address(hub)
        );
        uint256 feeWithBuffer = fee.nativeFee * 120 / 100; // 20% buffer

        vm.prank(LZ_ENDPOINT_BASE_SEPOLIA);
        vm.deal(LZ_ENDPOINT_BASE_SEPOLIA, feeWithBuffer);
        hub.lzReceive{value: feeWithBuffer}(originFromSpoke, bytes32(0), expectedMsgToHub, address(0), bytes(""));
        console.log("Simulated Hub receiving message from Spoke via LZ Endpoint.");

        // === 5. Simulate Spoke Receiving Return Message (Optimism Sepolia) ===
        console.log("\n=== Simulating Spoke receiving return message (Op Sepolia) ===");
        vm.selectFork(opForkId);

        Origin memory originFromHub = Origin({
            srcEid: BASE_SEPOLIA_EID,
            sender: bytes32(uint256(uint160(address(hub)))), // Use the CREATE3 deployed Hub address
            nonce: 0
        });

        assertEq(target.value(), 0, "Target value should be 0 before final execution");

        // vm.expectEmit(true, true, true, true, address(spoke));
        // emit ProxySpoke.FunctionExecuted(LZ_ENDPOINT_OP_SEPOLIA, address(target), targetCalldata, 0);

        vm.prank(LZ_ENDPOINT_OP_SEPOLIA);
        spoke.lzReceive(originFromHub, bytes32(0), abi.encode(address(target), targetCalldata), address(0), bytes(""));
        console.log("Simulated Spoke receiving message from Hub via LZ Endpoint.");

        // === 6. Verify Final State ===
        console.log("\n=== Verifying final state ===");
        assertEq(target.value(), valueToSet, "Target value mismatch after execution");
        console.log("Target value correctly set to:", target.value());
        assertEq(target.lastCaller(), address(spoke), "Target last caller mismatch");
        console.log("Target last caller correctly set to Spoke address:", target.lastCaller());

        console.log("\n Cross-chain flow simulation successful with CREATE3!");
    }
}