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

        address avsGovernanceAddress = vm.envAddress("AVS_GOVERNANCE_ADDRESS");
        address avsGovernanceLogicAddress = vm.envAddress("AVS_GOVERNANCE_LOGIC_ADDRESS");

        // Transaction 1: Deploy AvsGovernanceLogic
        // Separate broadcast block ensures this transaction is confirmed before proceeding
        console.log("\n=== Deploying AvsGovernanceLogic on Ethereum ===");
        vm.startBroadcast(deployerPrivateKey);
        AvsGovernanceLogic deploymentAddress = new AvsGovernanceLogic(
            vm.envAddress("LZ_ENDPOINT_ETH"),    // LayerZero endpoint
            vm.envAddress("TASK_EXECUTION_ADDRESS"), // TaskExecutionHub address
            uint32(vm.envUint("LZ_EID_BASE")),  // Destination chain ID (Base)
            deployer,                            // Owner address
            avsGovernanceAddress      // AVS Governance address
        );
        avsGovernanceLogicAddress = address(deploymentAddress);
        vm.stopBroadcast();

        // Transaction 2: Transfer ETH to AvsGovernanceLogic
        // Separate broadcast block ensures previous transaction is confirmed
        console.log("\n=== Transferring ETH to AvsGovernanceLogic ===");
        vm.startBroadcast(deployerPrivateKey);
        payable(avsGovernanceLogicAddress).transfer(0.1 ether);
        vm.stopBroadcast();

        // Transaction 3: Set the AvsGovernanceLogic on the AvsGovernance contract
        // Separate broadcast block ensures previous transaction is confirmed
        console.log("\n=== Setting AvsGovernanceLogic on AvsGovernance contract ===");
        vm.startBroadcast(deployerPrivateKey);
        AvsGovernance avsGovernanceContract = AvsGovernance(payable(vm.envAddress("AVS_GOVERNANCE_ADDRESS")));
        avsGovernanceContract.setAvsGovernanceLogic(IAvsGovernanceLogic(avsGovernanceLogicAddress));
        vm.stopBroadcast();

        // Print final deployment summary
        console.log("");
        console.log("=== AVS GOVERNANCE LOGIC DEPLOYMENT SUMMARY ===");
        console.log("Implementation:", avsGovernanceLogicAddress);
        console.log("Deployer:", deployer);
        console.log("");
    }
}
