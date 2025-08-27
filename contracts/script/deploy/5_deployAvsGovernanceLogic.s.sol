// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "forge-std/console.sol";
import {AvsGovernance} from "@othentic-core-contracts/NetworkManagement/L1/AvsGovernance.sol";
import {IAvsGovernanceLogic} from "@othentic-core-contracts/NetworkManagement/L1/interfaces/IAvsGovernanceLogic.sol";
import {AvsGovernanceLogic} from "../../src/AvsGovernanceLogic.sol";
import {TaskExecutionHub} from "../../src/lz/TaskExecutionHub.sol";

contract DeployAvsGovernanceLogic is Script {  
    uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
    address deployer = vm.addr(deployerPrivateKey);

    function run() external {
        // Create a fork using the Holesky RPC.
        vm.createSelectFork(vm.envString("ETH_RPC"));
        vm.startBroadcast(deployerPrivateKey);

        // Deploy AvsGovernanceLogic
        console.log("\n=== Deploying AvsGovernanceLogic on Ethereum ===");
        AvsGovernanceLogic avsGovernance = new AvsGovernanceLogic(
            vm.envAddress("LZ_ENDPOINT_ETH"),    // LayerZero endpoint
            vm.envAddress("TASK_EXECUTION_ADDRESS"), // TaskExecutionHub address
            uint32(vm.envUint("LZ_EID_BASE")),  // Destination chain ID (Base)
            deployer,                            // Owner address
            vm.envAddress("AVS_GOVERNANCE_ADDRESS")      // AVS Governance address
        );

        // Transfer ETH to AvsGovernanceLogic
        payable(address(avsGovernance)).transfer(0.1 ether);

        // Set the AvsGovernanceLogic on the AvsGovernance contract
        AvsGovernance avsGovernanceContract = AvsGovernance(payable(vm.envAddress("AVS_GOVERNANCE_ADDRESS")));
        avsGovernanceContract.setAvsGovernanceLogic(IAvsGovernanceLogic(address(avsGovernance)));
        
        vm.stopBroadcast();

        // Print final deployment summary
        console.log("");
        console.log("=== AVS GOVERNANCE LOGIC DEPLOYMENT SUMMARY ===");
        console.log("Implementation:", address(avsGovernance));
        console.log("Deployer:", deployer);
        console.log("");
    }
}
