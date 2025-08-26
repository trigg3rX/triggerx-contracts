// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "forge-std/console.sol";
import {AvsGovernanceLogic} from "../../src/AvsGovernanceLogic.sol";
import {TaskExecutionHub} from "../../src/lz/TaskExecutionHub.sol";

contract DeployAvsGovernance is Script {
 
    address payable constant TASK_EXECUTION_HUB = payable(0x7a2fC7bBE6c8a2248d651FdE1b1bD7d9509F6bfc);

    address payable constant AVS_GOVERNANCE = payable(0x530D21739B4B2cDccEcda80DCe5e1731BDBB9ed8); // TestnetProduction AVS Governance address

    // address payable constant AVS_GOVERNANCE = payable(0xBDd47006B79675274959fE8cA13470ed206a836D);
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
        uint256 forkId = vm.createSelectFork(vm.envString("SEPOLIA_RPC"));
        vm.startBroadcast(deployerPrivateKey);

        // Deploy AvsGovernanceLogic
        console.log("\n=== Deploying AvsGovernanceLogic on Sepolia ===");
        AvsGovernanceLogic avsGovernance = new AvsGovernanceLogic(
            LZ_ENDPOINT_HOLESKY,  // LayerZero endpoint
            TASK_EXECUTION_HUB,            // TaskExecutionHub address
            BASE_SEPOLIA_EID,     // Destination chain ID (Base Sepolia)
            deployerAddress,      // Owner address
            AVS_GOVERNANCE       // AVS Governance address
        );
        console.log("AvsGovernanceLogic deployed at:", address(avsGovernance));
        // use call to send the 0.5 eth to the deployed contract from the deployer address
        // (bool success, ) = address(avsGovernance).call{value: 0.5 ether}("");
        // require(success, "Failed to send 0.5 eth to the deployed contract");
        // console.log("0.5 eth sent to the deployed contract");

        // setTaskExecutionHub on Holesky
        // AvsGovernanceLogic avsGovernance = AvsGovernanceLogic(AVS_GOVERNANCE);
        console.log("\n=== Setting TaskExecutionHub on Holesky ===");
        avsGovernance.setTaskExecutionHub(TASK_EXECUTION_HUB);
        console.log("TaskExecutionHub set successfully on Holesky");
        
        vm.stopBroadcast();

        // Switch to Base Sepolia to set peer on TaskExecutionHub
        console.log("\n=== Switching to Base Sepolia to set peer on TaskExecutionHub ===");
        vm.createSelectFork(vm.envString("BASE_SEPOLIA_RPC"));
        vm.startBroadcast(deployerPrivateKey);
        
        TaskExecutionHub taskExecutionHub = TaskExecutionHub(TASK_EXECUTION_HUB);
        taskExecutionHub.setPeer(HOLESKY_EID, bytes32(uint256(uint160(address(avsGovernance)))));
        console.log("Peer set successfully on TaskExecutionHub");
        vm.stopBroadcast();
        // Call afterOperatorRegistered with the deployer address
        
        // console.log("\n=== Calling afterOperatorRegistered ===");
        // vm.selectFork(forkId);

        // AvsGovernanceLogic avsGovernance = AvsGovernanceLogic(payable(0xCd40429c3A85550EC75E3153F5fb7c0F06826EE5));
        // avsGovernance.afterOperatorRegistered(
        //     0x72461158B7abbd3741773F7F3BA145cE02F5177c,  // operator
        //     100,             // votingPower (example value)
        //     [uint256(0), 0, 0, 0],  // blsKey (example value)
        //     0x72461158B7abbd3741773F7F3BA145cE02F5177c  // rewardsReceiver
        // );
        // vm.stopBroadcast();
        // console.log("afterOperatorRegistered called successfully");

       
    

        console.log("\n--- Deployment Complete ---");
        console.log("AvsGovernanceLogic Address:", address(avsGovernance));
        console.log("---------------------------");
    }
} 