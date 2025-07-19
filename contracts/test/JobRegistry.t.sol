// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import "forge-std/Test.sol";
import {JobRegistry} from "../src/JobRegistry.sol";
import {ERC1967Proxy} from "@openzeppelin-contracts/contracts/proxy/ERC1967/ERC1967Proxy.sol";

contract JobRegistryTest is Test {
    JobRegistry public jobRegistry;
    JobRegistry public implementation;
    address public owner = address(0x1);
    address public user1 = address(0x2);
    address public user2 = address(0x3);
    address public targetContract = address(0x4);

    // Sample job data
    string constant JOB_NAME = "Test Job";
    uint256 constant JOB_TYPE = 1;
    uint256 constant TIME_FRAME = 3600; // 1 hour
    bytes constant JOB_DATA = "0x1234";

    event JobCreated(
        uint256 indexed jobId,
        address indexed jobOwner,
        bytes32 jobHash,
        uint256 timestamp
    );

    event JobUpdated(
        uint256 indexed jobId,
        address indexed jobOwner,
        bytes32 newJobHash,
        uint256 timestamp
    );

    event JobDeleted(
        uint256 indexed jobId,
        address indexed jobOwner,
        uint256 timestamp
    );

    function setUp() public {
        // Deploy implementation
        implementation = new JobRegistry();

        // Deploy proxy with initialization
        bytes memory initData = abi.encodeWithSelector(
            JobRegistry.initialize.selector,
            owner
        );
        
        ERC1967Proxy proxy = new ERC1967Proxy(
            address(implementation),
            initData
        );

        jobRegistry = JobRegistry(address(proxy));
    }

    function test_Initialize() public view {
        assertEq(jobRegistry.owner(), owner);
        assertEq(jobRegistry.getTotalJobsCount(), 0);
    }

    function test_Initialize_RevertIfZeroAddress() public {
        // Deploy new implementation
        JobRegistry newImpl = new JobRegistry();
        
        bytes memory initData = abi.encodeWithSelector(
            JobRegistry.initialize.selector,
            address(0)
        );

        vm.expectRevert(JobRegistry.InvalidJobParameters.selector);
        new ERC1967Proxy(address(newImpl), initData);
    }

    function test_CreateJob() public {
        vm.startPrank(user1);

        vm.expectEmit(true, true, false, true);
        emit JobCreated(1, user1, _calculateJobHash(), block.timestamp);

        uint256 jobId = jobRegistry.createJob(
            JOB_NAME,
            JOB_TYPE,
            TIME_FRAME,
            targetContract,
            JOB_DATA
        );

        assertEq(jobId, 1);
        assertEq(jobRegistry.getTotalJobsCount(), 1);

        JobRegistry.Job memory job = jobRegistry.getJob(jobId);
        assertEq(job.jobId, jobId);
        assertEq(job.jobOwner, user1);
        assertEq(job.jobHash, _calculateJobHash());
        assertEq(job.lastUpdatedAt, block.timestamp);
        assertTrue(job.isActive);

        uint256[] memory userJobs = jobRegistry.getUserJobIds(user1);
        assertEq(userJobs.length, 1);
        assertEq(userJobs[0], jobId);

        vm.stopPrank();
    }

    function test_CreateJob_RevertIfEmptyJobName() public {
        vm.startPrank(user1);

        vm.expectRevert(JobRegistry.EmptyJobName.selector);
        jobRegistry.createJob(
            "",
            JOB_TYPE,
            TIME_FRAME,
            targetContract,
            JOB_DATA
        );

        vm.stopPrank();
    }

    function test_CreateJob_RevertIfInvalidTargetContract() public {
        vm.startPrank(user1);

        vm.expectRevert(JobRegistry.InvalidTargetContract.selector);
        jobRegistry.createJob(
            JOB_NAME,
            JOB_TYPE,
            TIME_FRAME,
            address(0),
            JOB_DATA
        );

        vm.stopPrank();
    }

    function test_UpdateJob() public {
        // Create job first
        vm.startPrank(user1);
        uint256 jobId = jobRegistry.createJob(
            JOB_NAME,
            JOB_TYPE,
            TIME_FRAME,
            targetContract,
            JOB_DATA
        );

        // Update job
        string memory newJobName = "Updated Job";
        bytes32 newJobHash = _calculateJobHash(newJobName);

        vm.expectEmit(true, true, false, true);
        emit JobUpdated(jobId, user1, newJobHash, block.timestamp);

        jobRegistry.updateJob(
            jobId,
            newJobName,
            JOB_TYPE,
            TIME_FRAME,
            targetContract,
            JOB_DATA
        );

        JobRegistry.Job memory job = jobRegistry.getJob(jobId);
        assertEq(job.jobHash, newJobHash);
        assertEq(job.lastUpdatedAt, block.timestamp);

        vm.stopPrank();
    }

    function test_UpdateJob_RevertIfJobNotFound() public {
        vm.startPrank(user1);

        vm.expectRevert(abi.encodeWithSelector(JobRegistry.JobNotFound.selector, 999));
        jobRegistry.updateJob(
            999,
            JOB_NAME,
            JOB_TYPE,
            TIME_FRAME,
            targetContract,
            JOB_DATA
        );

        vm.stopPrank();
    }

    function test_UpdateJob_RevertIfNotOwner() public {
        // Create job as user1
        vm.prank(user1);
        uint256 jobId = jobRegistry.createJob(
            JOB_NAME,
            JOB_TYPE,
            TIME_FRAME,
            targetContract,
            JOB_DATA
        );

        // Try to update as user2
        vm.startPrank(user2);

        vm.expectRevert(abi.encodeWithSelector(JobRegistry.UnauthorizedJobAccess.selector, jobId, user2));
        jobRegistry.updateJob(
            jobId,
            "Updated Job",
            JOB_TYPE,
            TIME_FRAME,
            targetContract,
            JOB_DATA
        );

        vm.stopPrank();
    }

    function test_DeleteJob() public {
        // Create job first
        vm.startPrank(user1);
        uint256 jobId = jobRegistry.createJob(
            JOB_NAME,
            JOB_TYPE,
            TIME_FRAME,
            targetContract,
            JOB_DATA
        );

        assertTrue(jobRegistry.isJobActive(jobId));

        vm.expectEmit(true, true, false, true);
        emit JobDeleted(jobId, user1, block.timestamp);

        jobRegistry.deleteJob(jobId);

        assertFalse(jobRegistry.isJobActive(jobId));

        JobRegistry.Job memory job = jobRegistry.getJob(jobId);
        assertFalse(job.isActive);

        vm.stopPrank();
    }

    function test_DeleteJob_RevertIfAlreadyInactive() public {
        // Create and delete job first
        vm.startPrank(user1);
        uint256 jobId = jobRegistry.createJob(
            JOB_NAME,
            JOB_TYPE,
            TIME_FRAME,
            targetContract,
            JOB_DATA
        );
        jobRegistry.deleteJob(jobId);

        // Try to delete again
        vm.expectRevert(abi.encodeWithSelector(JobRegistry.JobAlreadyInactive.selector, jobId));
        jobRegistry.deleteJob(jobId);

        vm.stopPrank();
    }

    function test_GetUserActiveJobIds() public {
        vm.startPrank(user1);

        // Create multiple jobs
        uint256 jobId1 = jobRegistry.createJob(
            "Job 1",
            JOB_TYPE,
            TIME_FRAME,
            targetContract,
            JOB_DATA
        );

        uint256 jobId2 = jobRegistry.createJob(
            "Job 2",
            JOB_TYPE,
            TIME_FRAME,
            targetContract,
            JOB_DATA
        );

        uint256 jobId3 = jobRegistry.createJob(
            "Job 3",
            JOB_TYPE,
            TIME_FRAME,
            targetContract,
            JOB_DATA
        );

        // Delete one job
        jobRegistry.deleteJob(jobId2);

        uint256[] memory allJobs = jobRegistry.getUserJobIds(user1);
        assertEq(allJobs.length, 3);

        uint256[] memory activeJobs = jobRegistry.getUserActiveJobIds(user1);
        assertEq(activeJobs.length, 2);
        assertEq(activeJobs[0], jobId1);
        assertEq(activeJobs[1], jobId3);

        vm.stopPrank();
    }

    function test_GetJobHash() public {
        vm.startPrank(user1);

        uint256 jobId = jobRegistry.createJob(
            JOB_NAME,
            JOB_TYPE,
            TIME_FRAME,
            targetContract,
            JOB_DATA
        );

        bytes32 expectedHash = _calculateJobHash();
        bytes32 actualHash = jobRegistry.getJob(jobId).jobHash;
        assertEq(actualHash, expectedHash);

        vm.stopPrank();
    }

    function test_GetJobOwner() public {
        vm.startPrank(user1);

        uint256 jobId = jobRegistry.createJob(
            JOB_NAME,
            JOB_TYPE,
            TIME_FRAME,
            targetContract,
            JOB_DATA
        );

        address jobOwner = jobRegistry.getJobOwner(jobId);
        assertEq(jobOwner, user1);

        vm.stopPrank();
    }

    function test_GetJobOwner_NonExistentJob() public view {
        address jobOwner = jobRegistry.getJobOwner(999);
        assertEq(jobOwner, address(0)); // Non-existent job should return zero address
    }

    function test_Upgrade() public {
        // Deploy new implementation
        JobRegistry newImplementation = new JobRegistry();

        vm.startPrank(owner);
        jobRegistry.upgradeToAndCall(address(newImplementation), "");
        vm.stopPrank();

        // Contract should still work after upgrade
        assertEq(jobRegistry.owner(), owner);
    }

    function test_Upgrade_RevertIfNotOwner() public {
        JobRegistry newImplementation = new JobRegistry();

        vm.startPrank(user1);
        vm.expectRevert(); // OwnableUpgradeable uses different error messages
        jobRegistry.upgradeToAndCall(address(newImplementation), "");
        vm.stopPrank();
    }

    function test_MultipleUsersCanCreateJobs() public {
        // User1 creates job
        vm.prank(user1);
        uint256 jobId1 = jobRegistry.createJob(
            "User1 Job",
            JOB_TYPE,
            TIME_FRAME,
            targetContract,
            JOB_DATA
        );

        // User2 creates job
        vm.prank(user2);
        uint256 jobId2 = jobRegistry.createJob(
            "User2 Job",
            JOB_TYPE + 1,
            TIME_FRAME,
            targetContract,
            JOB_DATA
        );

        assertEq(jobId1, 1);
        assertEq(jobId2, 2);
        assertEq(jobRegistry.getTotalJobsCount(), 2);

        // Check job ownership
        JobRegistry.Job memory job1 = jobRegistry.getJob(jobId1);
        JobRegistry.Job memory job2 = jobRegistry.getJob(jobId2);

        assertEq(job1.jobOwner, user1);
        assertEq(job2.jobOwner, user2);

        // Check user job lists
        uint256[] memory user1Jobs = jobRegistry.getUserJobIds(user1);
        uint256[] memory user2Jobs = jobRegistry.getUserJobIds(user2);

        assertEq(user1Jobs.length, 1);
        assertEq(user2Jobs.length, 1);
        assertEq(user1Jobs[0], jobId1);
        assertEq(user2Jobs[0], jobId2);
    }

    // Helper function to calculate job hash
    function _calculateJobHash() internal view returns (bytes32) {
        return keccak256(
            abi.encode(JOB_NAME, JOB_TYPE, TIME_FRAME, targetContract, JOB_DATA)
        );
    }

    function _calculateJobHash(string memory jobName) internal view returns (bytes32) {
        return keccak256(
            abi.encode(jobName, JOB_TYPE, TIME_FRAME, targetContract, JOB_DATA)
        );
    }
} 