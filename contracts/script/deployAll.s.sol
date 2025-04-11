// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "forge-std/console.sol";
import {CREATE3} from "solady/utils/CREATE3.sol";
import {ProxyHub} from "../src/lz/ProxyHub.sol";
import {ProxySpoke} from "../src/lz/ProxySpoke.sol";
import {Target} from "../src/Target.sol";

contract DeployAll is Script {
    // --- Configuration (Update if needed) ---
    // LayerZero Endpoints
    address constant LZ_ENDPOINT_OP_SEPOLIA = 0x6EDCE65403992e310A62460808c4b910D972f10f;
    address constant LZ_ENDPOINT_BASE_SEPOLIA = 0x6EDCE65403992e310A62460808c4b910D972f10f;
    address constant LZ_ENDPOINT_HOLESKY = 0x6EDCE65403992e310A62460808c4b910D972f10f;

    // Endpoint IDs (EIDs for LayerZero)
    uint32 constant OP_SEPOLIA_EID = 40232; // Optimism Sepolia Endpoint ID
    uint32 constant BASE_SEPOLIA_EID = 40245; // Base Sepolia Endpoint ID
    uint32 constant HOLESKY_EID = 40217; // Holesky Endpoint ID

    // Deployment Salt (Must be the same for Hub and Spoke)
    bytes32 constant SALT = "KeeperProxy";

    function run() external {
        // Fetch deployer information from environment variables.
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployerAddress = vm.addr(deployerPrivateKey);
        console.log("Deploying contracts using deployer:", deployerAddress);

        // --- Deploy Hub on Base Sepolia ---
        console.log("\n=== Deploying ProxyHub on Base Sepolia ===");

        // Create a fork using the Base Sepolia RPC.
        vm.createSelectFork(vm.envString("BASE_SEPOLIA_RPC"));
        vm.startBroadcast(deployerPrivateKey);

        // Prepare the bytecode for ProxyHub.
        // ProxyHub's constructor takes: (address _endpoint, address _owner, uint32 _srcEid, uint32 _thisChainEid)
        bytes memory hubBytecode = abi.encodePacked(
            type(ProxyHub).creationCode,
            abi.encode(LZ_ENDPOINT_BASE_SEPOLIA, deployerAddress, HOLESKY_EID, BASE_SEPOLIA_EID)
        );
        
        // Deploy using CREATE3 with no ETH value.
        address hubAddr = CREATE3.deployDeterministic(hubBytecode, SALT);
        ProxyHub hub = ProxyHub(payable(hubAddr));
        console.log("ProxyHub deployed at:", hubAddr);

        // Add spoke endpoints to Hub
        uint32[] memory spokeEids = new uint32[](1);
        spokeEids[0] = OP_SEPOLIA_EID;
        hub.addSpokes(spokeEids);
        console.log("Added spoke endpoint:", OP_SEPOLIA_EID);

        // Send ETH to Hub contract to cover LayerZero fees
        vm.deal(address(hub), 1 ether);
        console.log("Sent 1 ETH to Hub contract at:", address(hub));

        vm.stopBroadcast();

        // --- Deploy Spoke on OP Sepolia ---
        console.log("\n=== Deploying ProxySpoke on OP Sepolia ===");

        // Create a fork using the OP Sepolia RPC.
        vm.createSelectFork(vm.envString("OPSEPOLIA_RPC"));
        vm.startBroadcast(deployerPrivateKey);

        // Prepare the bytecode for ProxySpoke.
        // ProxySpoke's constructor takes: (address _endpoint, address _owner)
        bytes memory spokeBytecode = abi.encodePacked(
            type(ProxySpoke).creationCode,
            abi.encode(LZ_ENDPOINT_OP_SEPOLIA, deployerAddress)
        );
        
        // Deploy using CREATE3 with the same salt.
        address spokeAddr = CREATE3.deployDeterministic(spokeBytecode, SALT);
        ProxySpoke spoke = ProxySpoke(payable(spokeAddr));
        console.log("ProxySpoke deployed at:", spokeAddr);

        vm.stopBroadcast();

        // --- Verify final state ---
        console.log("\n--- Deployment Complete ---");
  
        console.log("Hub Address:", address(hub));
        console.log("Hub Owner:", hub.owner());
        console.log("Hub AVSGovernance:", deployerAddress);
        console.log("Spoke Address:", address(spoke));
        console.log("Spoke Owner:", spoke.owner());
        console.log("---------------------------");
    }
} 