// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "forge-std/console.sol";
import {AvsGovernanceLogic} from "../src/AvsGovernanceLogic.sol";
import {ProxyHub} from "../src/lz/ProxyHub.sol";

contract DeployAvsGovernance is Script {
 
    address payable constant PROXY_HUB = payable(0x96c5F575940DBe98fd9600F74F0c36139A7Be2A2);

    address payable constant AVS_GOVERNANCE = payable(0xBDd47006B79675274959fE8cA13470ed206a836D);
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
        // vm.createSelectFork(vm.envString("L2_RPC"));
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
        // use call to send the 0.5 eth to the deployed contract from the deployer address
        (bool success, ) = address(avsGovernance).call{value: 0.5 ether}("");
        require(success, "Failed to send 0.5 eth to the deployed contract");
        console.log("0.5 eth sent to the deployed contract");

        // setProxyHub on Holesky
        // AvsGovernanceLogic avsGovernance = AvsGovernanceLogic(AVS_GOVERNANCE);
        console.log("\n=== Setting proxyHub on Holesky ===");
        avsGovernance.setProxyHub(PROXY_HUB);
        console.log("ProxyHub set successfully on Holesky");
        
        vm.stopBroadcast();

        // Switch to Base Sepolia to set peer on ProxyHub
        console.log("\n=== Switching to Base Sepolia to set peer on ProxyHub ===");
        vm.createSelectFork(vm.envString("BASE_SEPOLIA_RPC"));
        vm.startBroadcast(deployerPrivateKey);
        
        ProxyHub proxyHub = ProxyHub(PROXY_HUB);
        proxyHub.setPeer(HOLESKY_EID, bytes32(uint256(uint160(address(avsGovernance)))));
        console.log("Peer set successfully on ProxyHub");
        vm.stopBroadcast();
        // // Call afterOperatorRegistered with the deployer address
        console.log("\n=== Calling afterOperatorRegistered ===");
        avsGovernance.afterOperatorRegistered(
            deployerAddress,  // operator
            100,             // votingPower (example value)
            [uint256(0), 0, 0, 0],  // blsKey (example value)
            deployerAddress  // rewardsReceiver
        );
        console.log("afterOperatorRegistered called successfully");

       
    

        console.log("\n--- Deployment Complete ---");
        console.log("AvsGovernanceLogic Address:", address(avsGovernance));
        console.log("---------------------------");
    }
} 