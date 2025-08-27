// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "forge-std/console.sol";
import {CREATE3} from "@solady/utils/CREATE3.sol";
import {ERC1967Proxy} from "openzeppelin-contracts/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {IAttestationCenter} from "../../src/interfaces/IAttestationCenter.sol";
import {TaskExecutionHub} from "../../src/lz/TaskExecutionHub.sol";
import {TaskExecutionSpoke} from "../../src/lz/TaskExecutionSpoke.sol";
import {TriggerGasRegistry} from "../../src/TriggerGasRegistry.sol";

contract DeployTaskExecutionSpoke is Script {
    uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
    address deployer = vm.addr(deployerPrivateKey);

    bytes32 SALT = keccak256(abi.encodePacked("put_salt_here"));
    bytes32 IMPL_SALT = keccak256(abi.encodePacked("put_salt_here"));

    address[] private operators;
    address private spokeAddress;
    address private spokeImpl;

    function fetchOperatorsFromAttestationCenter() internal returns (uint256) {
        // Create fork for Base to fetch operators
        vm.createSelectFork(vm.envString("BASE_RPC"));

        console.log("Attempting to fetch operators from AttestationCenter...");
        
        IAttestationCenter attestationCenter = IAttestationCenter(vm.envAddress("ATTESTATION_CENTER_ADDRESS"));
        
        // Use numOfOperators() which is defined in the interface
        try attestationCenter.numOfTotalOperators() returns (uint256 operatorCount) {
            console.log("Number of operators found:", operatorCount);
            
            // if (operatorCount == 0) {
            //     revert("No operators found in AttestationCenter");
            // }
            
            return operatorCount;
        } catch {
            console.log("numOfTotalOperators failed");
        }
        
        return 0;
    }
    
    function fetchOperatorsByCount(IAttestationCenter attestationCenter, uint256 operatorCount) internal returns (address[] memory) {
        vm.createSelectFork(vm.envString("BASE_RPC"));

        address[] memory tempOperators = new address[](operatorCount);
        uint256 validOperators = 0;
        
        for (uint256 i = 1; i <= operatorCount; i++) {
            try attestationCenter.getOperatorPaymentDetail(i) returns (IAttestationCenter.PaymentDetails memory details) {
                if (details.operator != address(0)) {
                    tempOperators[validOperators] = details.operator;
                    console.log("Operator", i, ":", details.operator);
                    validOperators++;
                }
            } catch {
                console.log("Failed to fetch operator", i, "- skipping");
            }
        }
        
        // Resize array to only include valid operators
        if (validOperators < operatorCount) {
            address[] memory validOperatorArray = new address[](validOperators);
            for (uint256 i = 0; i < validOperators; i++) {
                validOperatorArray[i] = tempOperators[i];
            }
            return validOperatorArray;
        }
        
        return tempOperators;
    }

    function deploySpoke() internal returns (address) {
        console.log("\n=== Deploying TaskExecutionSpoke ===");
        
        // Create fork for the deployment network
        // vm.createSelectFork(vm.envString("OP_RPC"));
        vm.createSelectFork(vm.envString("ARB_RPC"));
        vm.startBroadcast(deployerPrivateKey);

        // 1. Deploy implementation
        spokeImpl = address(new TaskExecutionSpoke(vm.envAddress("LZ_ENDPOINT_ARB"), deployer));
        console.log("TaskExecutionSpoke implementation deployed at:", spokeImpl);

        // 2. Prepare initialization calldata
        bytes memory initData = abi.encodeWithSelector(
            TaskExecutionSpoke.initialize.selector,
            deployer,    // _ownerAddress
            vm.envUint("LZ_EID_BASE"),   // _hubEid
            operators,          // _initialKeepers
            vm.envAddress("JOB_REGISTRY_ADDRESS"),       // _jobRegistryAddress (random)
            vm.envAddress("TRIGGER_GAS_REGISTRY_ADDRESS")        // _triggerGasRegistryAddress (random)
        );

        // 3. Prepare proxy bytecode
        bytes memory proxyBytecode = abi.encodePacked(
            type(ERC1967Proxy).creationCode,
            abi.encode(spokeImpl, initData)
        );
        
        // 4. Deploy proxy using CREATE3
        spokeAddress = CREATE3.deployDeterministic(proxyBytecode, SALT);
        console.log("TaskExecutionSpoke proxy deployed at:", spokeAddress);

        TriggerGasRegistry(vm.envAddress("TRIGGER_GAS_REGISTRY_ADDRESS")).setOperator(spokeAddress);

        console.log("Operator Role:", TriggerGasRegistry(vm.envAddress("TRIGGER_GAS_REGISTRY_ADDRESS")).operatorRole());

        vm.stopBroadcast();
        
        return spokeAddress;
    }

    function run() external {
        // Fetch operators from Attestati onCenter
        uint256 operatorCount;
        operatorCount = fetchOperatorsFromAttestationCenter();
        console.log("Successfully fetched", operatorCount, "operators");

        operators = fetchOperatorsByCount(IAttestationCenter(vm.envAddress("ATTESTATION_CENTER_ADDRESS")), operatorCount);
        console.log("Successfully fetched", operators.length, "operators");

        // Deploy Spoke on Base Sepolia
        deploySpoke();
        
        // Print final deployment summary
        console.log("");
        console.log("=== TASK EXECUTION SPOKE DEPLOYMENT SUMMARY ===");
        console.log("Proxy:", spokeAddress);
        console.log("Implementation:", spokeImpl);
        console.log("Deployer:", deployer);
        console.log("");
    }
}
