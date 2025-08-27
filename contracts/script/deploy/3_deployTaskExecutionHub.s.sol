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

contract DeployTaskExecutionHub is Script {
    uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
    address deployer = vm.addr(deployerPrivateKey);

    bytes32 SALT = keccak256(abi.encodePacked("put_salt_here"));
    bytes32 IMPL_SALT = keccak256(abi.encodePacked("put_salt_here"));

    address[] private operators;
    address private hubAddress;
    address private hubImpl;

    function fetchOperatorsFromAttestationCenter() internal view returns (address[] memory) {
        console.log("Attempting to fetch operators from AttestationCenter...");
        
        IAttestationCenter attestationCenter = IAttestationCenter(vm.envAddress("ATTESTATION_CENTER_ADDRESS"));
        
        // Use numOfOperators() which is defined in the interface
        try attestationCenter.numOfTotalOperators() returns (uint256 operatorCount) {
            console.log("Number of operators found:", operatorCount);
            
            // if (operatorCount == 0) {
            //     revert("No operators found in AttestationCenter");
            // }
            
            return fetchOperatorsByCount(attestationCenter, operatorCount);
            
        } catch Error(string memory reason) {
            console.log("numOfOperators failed:", reason);
        } catch {
            console.log("numOfOperators failed with unknown error");
        }
        
        // If all methods fail, revert with error
        revert("Unable to fetch operators from AttestationCenter. Please check the contract address and network.");
    }
    
    function fetchOperatorsByCount(IAttestationCenter attestationCenter, uint256 operatorCount) internal view returns (address[] memory) {
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

    function deployHub() internal {
        vm.startBroadcast(deployerPrivateKey);

        // 1. Deploy the implementation contract
        hubImpl = address(new TaskExecutionHub(vm.envAddress("LZ_ENDPOINT_BASE"), deployer));

        // 2. Prepare the initialization calldata
        bytes memory initData = abi.encodeWithSelector(
            TaskExecutionHub.initialize.selector,
            deployer,                   // _ownerAddress
            vm.envUint("LZ_EID_ETH"),           // _srcEid
            vm.envUint("LZ_EID_BASE"),           // _originEid
            operators,                         // _initialKeepers
            vm.envAddress("JOB_REGISTRY_ADDRESS"),        // _jobRegistryAddress (random)
            vm.envAddress("TRIGGER_GAS_REGISTRY_ADDRESS") // _triggerGasRegistryAddress (random)
        );

        // 3. Prepare the proxy bytecode
        bytes memory proxyBytecode = abi.encodePacked(
            type(ERC1967Proxy).creationCode,
            abi.encode(hubImpl, initData)
        );
        
        // 4. Deploy proxy using CREATE3
        hubAddress = CREATE3.deployDeterministic(proxyBytecode, SALT);

        vm.stopBroadcast();
    }

    function depositEthToHub() internal {
        vm.startBroadcast(deployerPrivateKey);
        payable(hubAddress).transfer(0.1 ether);
        vm.stopBroadcast();
    }

    function run() external {
        // Create a fork using the Base Sepolia RPC to fetch operators
        vm.createSelectFork(vm.envString("BASE_RPC"));
        
        // Fetch operators from AttestationCenter
        operators = fetchOperatorsFromAttestationCenter();
        console.log("Successfully fetched", operators.length, "operators");

        // Deploy Hub on Base Sepolia
        deployHub();

        depositEthToHub();

        console.log("");
        console.log("=== TASK EXECUTION HUB DEPLOYMENT SUMMARY ===");
        console.log("Proxy:", hubAddress);
        console.log("Implementation:", hubImpl);
        console.log("Deployer:", deployer);
        console.log("");
    }
}

contract AddSokesToHub is Script {
    function run() public {
        vm.createSelectFork(vm.envString("BASE_RPC"));
        vm.startBroadcast(vm.envUint("PRIVATE_KEY"));

        uint32[] memory spokeEids = new uint32[](3);
        spokeEids[0] = uint32(vm.envUint("LZ_EID_ARB"));
        spokeEids[1] = uint32(vm.envUint("LZ_EID_ETH"));
        spokeEids[2] = uint32(vm.envUint("LZ_EID_OP"));

        TaskExecutionHub hub = TaskExecutionHub(payable(vm.envAddress("TASK_EXECUTION_ADDRESS")));
        hub.addSpokes(spokeEids);

        vm.stopBroadcast();
    }
}

// contract SetJobRegistryonHub is Script {
//     function run() public {
//         vm.createSelectFork(vm.envString("BASE_RPC"));
//         vm.startBroadcast(vm.envUint("PRIVATE_KEY"));

//         TaskExecutionHub hub = TaskExecutionHub(payable(vm.envAddress("TASK_EXECUTION_ADDRESS")));
//         hub.setJobRegistry(vm.envAddress("JOB_REGISTRY_ADDRESS"));

//         vm.stopBroadcast();
//     }
// }

// contract SetTriggerGasRegistryonHub is Script {
//     function run() public {
//         vm.createSelectFork(vm.envString("BASE_RPC"));
//         vm.startBroadcast(vm.envUint("PRIVATE_KEY"));

//         TaskExecutionHub hub = TaskExecutionHub(payable(vm.envAddress("TASK_EXECUTION_ADDRESS")));
//         hub.setTriggerGasRegistry(vm.envAddress("TRIGGER_GAS_REGISTRY_ADDRESS"));

//         vm.stopBroadcast();
//     }
// }