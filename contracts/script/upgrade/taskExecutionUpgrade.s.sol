// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Script} from "forge-std/Script.sol";
import {console} from "forge-std/console.sol";
import {console2} from "forge-std/console2.sol";
import {TaskExecutionHub} from "../../src/lz/TaskExecutionHub.sol";
import {TaskExecutionSpoke} from "../../src/lz/TaskExecutionSpoke.sol";
import {PackedJobIdLib} from "../../src/libraries/PackedJobIdLib.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {CREATE3} from "@solady/utils/CREATE3.sol";

contract TaskExecutionUpgrade is Script {
    // Production proxy addresses
    address payable constant TASK_EXECUTION_HUB_PROXY = payable(0x179c62e83c3f90981B65bc12176FdFB0f2efAD54);
    
    // Salt for deterministic deployment
    bytes32 constant IMPL_SALT = keccak256(abi.encodePacked("put_salt_here"));
    
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

        // Upgrade TaskExecutionHub on Base Sepolia
        upgradeTaskExecutionHub(deployerPrivateKey, deployer);

        // Upgrade TaskExecutionSpoke on multiple chains
        upgradeTaskExecutionSpokeOnChain("SEPOLIA_RPC_URL", "SEPOLIA", TASK_EXECUTION_HUB_PROXY, deployerPrivateKey);
        upgradeTaskExecutionSpokeOnChain("OPSEPOLIA_RPC", "OP SEPOLIA", TASK_EXECUTION_HUB_PROXY, deployerPrivateKey);
        upgradeTaskExecutionSpokeOnChain("ARB_SEPOLIA_RPC", "ARBITRUM SEPOLIA", TASK_EXECUTION_HUB_PROXY, deployerPrivateKey);
    }

    function upgradeTaskExecutionHub(uint256 deployerPrivateKey, address deployer) internal {
        address payable proxy = TASK_EXECUTION_HUB_PROXY;

        vm.createSelectFork(vm.envString("BASE_SEPOLIA_RPC"));
        vm.startBroadcast(deployerPrivateKey);

        
        // STEP 1: Capture BEFORE state
        console.log("\n=== TASKEXECUTIONHUB UPGRADE ON BASE SEPOLIA - BEFORE ===");
        ContractState memory beforeState = captureHubState(proxy);
        
        console.log("Current implementation:", beforeState.implementation);
        console.log("Current owner:", beforeState.owner);
        console2.log("Keeper count: %s", beforeState.keeperCount);
        console.log("Job registry:", beforeState.jobRegistry);
        console.log("Trigger gas registry:", beforeState.triggerGasRegistry);

        // STEP 2: Deploy new implementation using CREATE3
        address lzEndpoint = address(TaskExecutionHub(proxy).endpoint());
        address delegate = TaskExecutionHub(proxy).owner(); // Use current owner as delegate
        
        bytes memory implementationCode = abi.encodePacked(
            type(TaskExecutionHub).creationCode,
            abi.encode(lzEndpoint, delegate)
        );
        
        address newImplementation = CREATE3.deployDeterministic(implementationCode, IMPL_SALT);
        console.log("\n=== PERFORMING TASKEXECUTIONHUB UPGRADE ON BASE SEPOLIA ===");
        console.log("Deploying new implementation to:", newImplementation);
        
        // STEP 3: Perform upgrade
        TaskExecutionHub(proxy).upgradeToAndCall(newImplementation, "");

        vm.stopBroadcast();

        // STEP 4: Capture AFTER state
        console.log("\n=== TASKEXECUTIONHUB UPGRADE ON BASE SEPOLIA - AFTER ===");
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

    function upgradeTaskExecutionSpokeOnChain(
        string memory rpcEnvVar,
        string memory chainName,
        address payable spokeProxy,
        uint256 deployerPrivateKey
    ) internal {
        vm.createSelectFork(vm.envString(rpcEnvVar));
        vm.startBroadcast(deployerPrivateKey);
        
        console.log(string.concat("\n=== TASKEXECUTIONSPOKE UPGRADE ON ", chainName, " ==="));
        console.log("Upgrading spoke at:", spokeProxy);
        
        // Capture current state
        TaskExecutionSpoke spoke = TaskExecutionSpoke(spokeProxy);
        address oldImpl = getImplementation(spokeProxy);
        address owner = spoke.owner();
        address jobRegistry = address(spoke.jobRegistry());
        
        console.log("Current implementation:", oldImpl);
        console.log("Owner:", owner);
        console.log("Job registry:", jobRegistry);
        
        // Deploy new implementation using CREATE3
        address lzEndpoint = address(spoke.endpoint());
        
        bytes memory implementationCode = abi.encodePacked(
            type(TaskExecutionSpoke).creationCode,
            abi.encode(lzEndpoint, owner)
        );
        
        address newImplementation = CREATE3.deployDeterministic(implementationCode, IMPL_SALT);
        console.log("Deploying new TaskExecutionSpoke implementation to:", newImplementation);
        
        // Perform upgrade
        spoke.upgradeToAndCall(newImplementation, "");
        
        // Verify upgrade
        address newImpl = getImplementation(spokeProxy);
        console.log("New implementation:", newImpl);
        
        if (newImpl != oldImpl) {
            console.log(string.concat("SUCCESS: TaskExecutionSpoke upgraded successfully on ", chainName));
        } else {
            console.log(string.concat("ERROR: TaskExecutionSpoke upgrade failed on ", chainName));
        }
            
        vm.stopBroadcast();
    }

    function getImplementation(address proxy) internal view returns (address) {
        bytes32 slot = 0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc;
        return address(uint160(uint256(vm.load(proxy, slot))));
    }
}