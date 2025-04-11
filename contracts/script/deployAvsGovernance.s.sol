// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "forge-std/console.sol";
import {AvsGovernanceLogic} from "../src/AvsGovernanceLogic.sol";
import {ProxyHub} from "../src/lz/ProxyHub.sol";

contract DeployAvsGovernance is Script {
 
    address constant PROXY_HUB = 0x9b34C613AfC61725C4650109b02cCc62518cc159;
     // LayerZero Endpoints
    address constant LZ_ENDPOINT_OP_SEPOLIA = 0x6EDCE65403992e310A62460808c4b910D972f10f;
    address constant LZ_ENDPOINT_BASE_SEPOLIA = 0x6EDCE65403992e310A62460808c4b910D972f10f;
    address constant LZ_ENDPOINT_HOLESKY = 0x6EDCE65403992e310A62460808c4b910D972f10f; // Update with actual Holesky endpoint

    // Endpoint IDs (EIDs for LayerZero)
    uint32 constant OP_SEPOLIA_EID = 40232; // Optimism Sepolia Endpoint ID
    uint32 constant BASE_SEPOLIA_EID = 40245; // Base Sepolia Endpoint ID
    uint32 constant HOLESKY_EID = 40217; 

    function run() external {
        // Fetch deployer information from environment variables.
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployerAddress = vm.addr(deployerPrivateKey);
        console.log("Deploying contracts using deployer:", deployerAddress);

        // Create a fork using the Holesky RPC.
        vm.createSelectFork(vm.envString("HOLESKY_RPC"));
        vm.startBroadcast(deployerPrivateKey);

        // Deploy AvsGovernanceLogic
        console.log("\n=== Deploying AvsGovernanceLogic on Holesky ===");
        AvsGovernanceLogic avsGovernance = new AvsGovernanceLogic(
            LZ_ENDPOINT_HOLESKY,  // LayerZero endpoint
            PROXY_HUB,            // ProxyHub address
            BASE_SEPOLIA_EID,     // Destination chain ID (Base Sepolia)
            deployerAddress       // Owner address
        );
        console.log("AvsGovernanceLogic deployed at:", address(avsGovernance));


    
        vm.stopBroadcast();

        console.log("\n--- Deployment Complete ---");
        console.log("AvsGovernanceLogic Address:", address(avsGovernance));
        console.log("---------------------------");
    }
} 