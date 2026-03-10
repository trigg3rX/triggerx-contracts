// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Script} from "forge-std/Script.sol";
import {console} from "forge-std/console.sol";
import {TriggerXSafeModule} from "../../src/TriggerXSafeModule.sol";
import {CREATE3} from "@solady/utils/CREATE3.sol";

contract DeployTriggerXSafeModule is Script {
    uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
    address deployer = vm.addr(deployerPrivateKey);
    address taskExecutionHub = vm.envAddress("TASK_EXECUTION_ADDRESS");

    bytes32 SALT = keccak256(abi.encodePacked("triggerx-safe-module-v0-1"));

    function run() public {
        // Create fork for this chain
        // vm.createSelectFork(vm.envString("BASE_RPC"));
        // vm.createSelectFork(vm.envString("OP_RPC"));
        vm.createSelectFork(vm.envString("ARB_RPC"));

        vm.startBroadcast(deployerPrivateKey);
        
        // Deploy module using CREATE3 for deterministic address across all chains
        bytes memory bytecode = abi.encodePacked(
            type(TriggerXSafeModule).creationCode,
            abi.encode(taskExecutionHub)
        );
        address moduleAddress = CREATE3.deployDeterministic(bytecode, SALT);
        TriggerXSafeModule module = TriggerXSafeModule(moduleAddress);
        
        vm.stopBroadcast();

        // Print summary
        console.log("");
        console.log("=== TRIGGERX SAFE MODULE DEPLOYMENT SUMMARY ===");
        console.log("TriggerX Safe Module:", address(module));
        console.log("Task Execution Hub:", taskExecutionHub);
        console.log("Deployer:", deployer);
        console.log("");
    }
}
