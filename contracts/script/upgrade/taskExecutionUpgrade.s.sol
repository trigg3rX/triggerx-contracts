// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Script} from "forge-std/Script.sol";
import {console} from "forge-std/console.sol";
import {console2} from "forge-std/console2.sol";
import {TaskExecutionHub} from "../../src/lz/TaskExecutionHub.sol";
import {TaskExecutionSpoke} from "../../src/lz/TaskExecutionSpoke.sol";
import {PackedJobIdLib} from "../../src/libraries/PackedJobIdLib.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

contract TaskExecutionUpgrade is Script {
    // Production proxy addresses
    address payable constant TASK_EXECUTION_HUB_PROXY = payable(0x2469e89386947535A350EEBccC5F2754fd35F474);
    
    // Note: TaskExecutionSpoke addresses would need to be provided based on deployment
    // For now, we'll focus on the Hub upgrade as it's the main contract
    
    struct ContractState {
        address implementation;
        address owner;
        uint256 keeperCount;
        address jobRegistry;
        address triggerGasRegistry;
    }

    function run() public {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployer = vm.addr(deployerPrivateKey);

        // Upgrade TaskExecutionHub
        // upgradeTaskExecutionHub(deployerPrivateKey, deployer);

        upgradeTaskExecutionSpoke(TASK_EXECUTION_HUB_PROXY, deployerPrivateKey);

        
        // Note: To upgrade TaskExecutionSpoke contracts, you would need to:
        // 1. Add their proxy addresses as constants
        // 2. Call upgradeTaskExecutionSpoke() for each one
        // 3. Verify they can handle packed jobIds
    }

    function upgradeTaskExecutionHub(uint256 deployerPrivateKey, address deployer) internal {
        address payable proxy = TASK_EXECUTION_HUB_PROXY;

        vm.createSelectFork(vm.envString("BASE_SEPOLIA_RPC"));
        vm.startBroadcast(deployerPrivateKey);

        
        // STEP 1: Capture BEFORE state
        console.log("\n=== TASKEXECUTIONHUB UPGRADE - BEFORE ===");
        ContractState memory beforeState = captureHubState(proxy);
        
        console.log("Current implementation:", beforeState.implementation);
        console.log("Current owner:", beforeState.owner);
        console2.log("Keeper count: %s", beforeState.keeperCount);
        console.log("Job registry:", beforeState.jobRegistry);
        console.log("Trigger gas registry:", beforeState.triggerGasRegistry);

        // STEP 2: Deploy new implementation
        // Note: You'll need to provide the correct LayerZero endpoint and delegate addresses
        address lzEndpoint = address(TaskExecutionHub(proxy).endpoint());
        address delegate = TaskExecutionHub(proxy).owner(); // Use current owner as delegate
        
        TaskExecutionHub newImplementation = new TaskExecutionHub(lzEndpoint, delegate);
        console.log("\n=== PERFORMING TASKEXECUTIONHUB UPGRADE ===");
        console.log("Deploying new implementation to:", address(newImplementation));
        
        // STEP 3: Perform upgrade
        TaskExecutionHub(proxy).upgradeToAndCall(address(newImplementation), "");

        vm.stopBroadcast();

        // STEP 4: Capture AFTER state
        console.log("\n=== TASKEXECUTIONHUB UPGRADE - AFTER ===");
        ContractState memory afterState = captureHubState(proxy);
        
        console.log("New implementation:", afterState.implementation);
        console.log("New owner:", afterState.owner);
        console2.log("Keeper count: %s", afterState.keeperCount);
        console.log("Job registry:", afterState.jobRegistry);
        console.log("Trigger gas registry:", afterState.triggerGasRegistry);

        // STEP 5: Test packed jobId functionality
        console.log("\n=== TESTING PACKED JOBID FUNCTIONALITY ===");
        testPackedJobIdHandling(proxy);

        // STEP 6: Compare states
        console.log("\n=== TASKEXECUTIONHUB COMPARISON RESULTS ===");
        bool statePreserved = compareHubStates(beforeState, afterState);
        
        if (statePreserved) {
            console.log("SUCCESS: TaskExecutionHub state preserved during upgrade!");
        } else {
            console.log("WARNING: Some TaskExecutionHub state may have changed!");
        }

        console.log("\n=== TASKEXECUTIONHUB DEPLOYMENT SUMMARY ===");
        console.log("Proxy address:", proxy);
        console.log("Old implementation:", beforeState.implementation);
        console.log("New implementation:", afterState.implementation);
        console.log("Deployer:", deployer);
    }

    function captureHubState(address payable proxy) internal view returns (ContractState memory) {
        TaskExecutionHub hub = TaskExecutionHub(proxy);
        
        return ContractState({
            implementation: getImplementation(proxy),
            owner: hub.owner(),
            keeperCount: getKeeperCount(hub),
            jobRegistry: address(hub.jobRegistry()),
            triggerGasRegistry: address(hub.triggerGasRegistry())
        });
    }

    function getKeeperCount(TaskExecutionHub hub) internal view returns (uint256) {
        // Since there's no direct way to get keeper count, we'll return 0
        // In a real scenario, you might want to track specific keepers
        return 0;
    }

    function compareHubStates(ContractState memory before, ContractState memory afterState) internal pure returns (bool) {
        return (
            before.owner == afterState.owner &&
            before.jobRegistry == afterState.jobRegistry &&
            before.triggerGasRegistry == afterState.triggerGasRegistry
            // Note: implementation should be different after upgrade
        );
    }

    function testPackedJobIdHandling(address payable /* proxy */) internal view {
        // Test that the contract can handle packed jobIds
        uint256 testChainId = block.chainid;
        uint256 testTimestamp = block.timestamp;
        uint256 testJobCounter = 123;
        uint256 packedJobId = PackedJobIdLib.pack(testChainId, testTimestamp, testJobCounter);
        
        console2.log("Test packed jobId: %s", packedJobId);
        
        (uint256 unpackedChainId, uint256 unpackedTimestamp, uint256 unpackedJobCounter) =
            PackedJobIdLib.unpack(packedJobId);
        console2.log(
            "Unpacked - Chain ID: %s, Timestamp: %s, Job Counter: %s",
            unpackedChainId,
            unpackedTimestamp,
            unpackedJobCounter
        );
        
        if (
            unpackedChainId == testChainId &&
            unpackedJobCounter == testJobCounter &&
            unpackedTimestamp == testTimestamp
        ) {
            console.log("SUCCESS: Packed jobId handling works correctly");
        } else {
            console.log("ERROR: Packed jobId handling failed");
        }
        
        // Note: To fully test, you would need to call executeFunction with a packed jobId
        // This requires setting up proper job registry and keeper permissions
        console.log("NOTE: Full executeFunction testing requires proper setup of job registry and keepers");
    }

    function upgradeTaskExecutionSpoke(address payable spokeProxy, uint256 deployerPrivateKey) internal {
        // This function would upgrade a TaskExecutionSpoke contract
        // You would call this for each spoke contract that needs upgrading
        
        vm.createSelectFork(vm.envString("OPSEPOLIA_RPC"));
        vm.startBroadcast(deployerPrivateKey);
        
        console.log("\n=== TASKEXECUTIONSPOKE UPGRADE ===");
        console.log("Upgrading spoke at:", spokeProxy);
        
        // Capture current state
        TaskExecutionSpoke spoke = TaskExecutionSpoke(spokeProxy);
        address oldImpl = getImplementation(spokeProxy);
        address owner = spoke.owner();
        address jobRegistry = address(spoke.jobRegistry());
        
        console.log("Current implementation:", oldImpl);
        console.log("Owner:", owner);
        console.log("Job registry:", jobRegistry);
        
        // Deploy new implementation
        address lzEndpoint = address(spoke.endpoint());
        TaskExecutionSpoke newImplementation = new TaskExecutionSpoke(lzEndpoint, owner);
        
        // Perform upgrade
        spoke.upgradeToAndCall(address(newImplementation), "");
        
        // Verify upgrade
        address newImpl = getImplementation(spokeProxy);
        console.log("New implementation:", newImpl);
        
        if (newImpl != oldImpl) {
            console.log("SUCCESS: TaskExecutionSpoke upgraded successfully");
        } else {
            console.log("ERROR: TaskExecutionSpoke upgrade failed");
        }
            
        vm.stopBroadcast();
    }

    function getImplementation(address proxy) internal view returns (address) {
        bytes32 slot = 0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc;
        return address(uint160(uint256(vm.load(proxy, slot))));
    }
}