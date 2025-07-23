// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Script} from "forge-std/Script.sol";
import {console} from "forge-std/console.sol";
import {console2} from "forge-std/console2.sol";
import {JobRegistry} from "../../src/JobRegistry.sol";
import {PackedJobIdLib} from "../../src/libraries/PackedJobIdLib.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

contract JobRegistryUpgrade is Script {
    // Production proxy address
    address payable constant JOB_REGISTRY_PROXY = payable(0xdB66c11221234C6B19cCBd29868310c31494C21C);

    struct JobData {
        uint256 jobId;
        address jobOwner;
        bool isActive;
        uint256 chainId;
        uint256 jobCounter;
    }

    function run() public {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployer = vm.addr(deployerPrivateKey);

        address payable proxy = JOB_REGISTRY_PROXY;

        vm.startBroadcast(deployerPrivateKey);
        
        // STEP 1: Capture BEFORE state
        console.log("\n=== BEFORE UPGRADE ===");
        address oldImplementation = getImplementation(proxy);
        console.log("Current implementation:", oldImplementation);
        console.log("Current owner:", JobRegistry(proxy).owner());
        
        
        // Sample some existing jobs to verify they still work after upgrade
        JobData[] memory beforeJobs = new JobData[](3);
        for (uint256 i = 1; i <= 3; i++) {
            try JobRegistry(proxy).getJobByCounter(i) returns (JobRegistry.Job memory job) {
                beforeJobs[i-1] = JobData({
                    jobId: job.jobId,
                    jobOwner: job.jobOwner,
                    isActive: job.isActive,
                    chainId: 0, // Will be 0 for old format
                    jobCounter: i
                });
                console2.log("Job %s: owner=%s, active=%s", i, job.jobOwner, job.isActive);
            } catch {
                console2.log("Job %s: not found", i);
                beforeJobs[i-1] = JobData({
                    jobId: 0,
                    jobOwner: address(0),
                    isActive: false,
                    chainId: 0,
                    jobCounter: i
                });
            }
        }

        // STEP 2: Deploy new implementation
        JobRegistry newImplementation = new JobRegistry();
        console.log("\n=== PERFORMING UPGRADE ===");
        console.log("Deploying new implementation to:", address(newImplementation));
        
        // STEP 3: Perform upgrade
        JobRegistry(proxy).upgradeToAndCall(address(newImplementation), "");

        vm.stopBroadcast();

        // STEP 4: Capture AFTER state
        console.log("\n=== AFTER UPGRADE ===");
        address finalImplementation = getImplementation(proxy);
        console.log("New implementation:", finalImplementation);
        console.log("New owner:", JobRegistry(proxy).owner());

        // Verify existing jobs still work
        JobData[] memory afterJobs = new JobData[](3);
        for (uint256 i = 1; i <= 3; i++) {
            try JobRegistry(proxy).getJobByCounter(i) returns (JobRegistry.Job memory job) {
                afterJobs[i-1] = JobData({
                    jobId: job.jobId,
                    jobOwner: job.jobOwner,
                    isActive: job.isActive,
                    chainId: 0,
                    jobCounter: i
                });
                console2.log("Job %s: owner=%s, active=%s", i, job.jobOwner, job.isActive);
            } catch {
                console2.log("Job %s: not found", i);
                afterJobs[i-1] = JobData({
                    jobId: 0,
                    jobOwner: address(0),
                    isActive: false,
                    chainId: 0,
                    jobCounter: i
                });
            }
        }

        // STEP 5: Test new packed jobId functionality
        console.log("\n=== TESTING NEW PACKED JOBID FUNCTIONALITY ===");
        vm.startBroadcast(deployerPrivateKey);
        
        try JobRegistry(proxy).createJob(
            "Test Packed Job",
            1,
            3600,
            address(0x123),
            abi.encode(uint256(300))
        ) returns (uint256 newJobId) {
            console2.log("Created new job with packed ID: %s", newJobId);
            
            // Unpack and verify
            (uint256 chainId, uint256 jobCounter) = PackedJobIdLib.unpack(newJobId);
            console2.log("Unpacked - Chain ID: %s, Job Counter: %s", chainId, jobCounter);
            console2.log("Current chain ID: %s", block.chainid);
            
            if (chainId == block.chainid) {
                console.log("SUCCESS: Packed jobId correctly includes chain ID");
            } else {
                console.log("ERROR: Packed jobId has wrong chain ID");
            }
        } catch Error(string memory reason) {
            console.log("ERROR creating test job:", reason);
        }
        
        vm.stopBroadcast();

        // STEP 6: Compare BEFORE vs AFTER
        console.log("\n=== COMPARISON RESULTS ===");
        bool allDataPreserved = true;
        uint256 jobsWithData = 0;
        
        for (uint256 i = 0; i < 3; i++) {
            if (beforeJobs[i].jobOwner != address(0)) {
                jobsWithData++;
                
                bool ownerMatches = beforeJobs[i].jobOwner == afterJobs[i].jobOwner;
                bool activeMatches = beforeJobs[i].isActive == afterJobs[i].isActive;
                
                if (ownerMatches && activeMatches) {
                    console2.log("[OK] Job %s: Data preserved", i + 1);
                } else {
                    console2.log("[ERROR] Job %s: Data changed!", i + 1);
                    console2.log("  Owner: %s -> %s", beforeJobs[i].jobOwner, afterJobs[i].jobOwner);
                    console2.log("  Active: %s -> %s", beforeJobs[i].isActive, afterJobs[i].isActive);
                    allDataPreserved = false;
                }
            }
        }

        console.log("\n=== FINAL VERIFICATION ===");
        console2.log("Jobs with data: %s", jobsWithData);
        console2.log("Implementation changed: %s", oldImplementation != finalImplementation);
        
        if (allDataPreserved) {
            console.log("SUCCESS: All job data preserved during upgrade!");
        } else {
            console.log("CRITICAL: Some job data was lost during upgrade!");
        }

        console.log("\n=== DEPLOYMENT SUMMARY ===");
        console.log("Proxy address:", proxy);
        console.log("Old implementation:", oldImplementation);
        console.log("New implementation:", finalImplementation);
        console.log("Deployer:", deployer);
    }

    function getImplementation(address proxy) internal view returns (address) {
        bytes32 slot = 0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc;
        return address(uint160(uint256(vm.load(proxy, slot))));
    }
}
