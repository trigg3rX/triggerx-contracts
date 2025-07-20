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
    uint8 constant JOB_TYPE = 1;
    uint256 constant TIME_FRAME = 3600; // 1 hour
    bytes constant JOB_DATA = abi.encode(uint256(300)); // 5 minutes timeInterval for JobType 1

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
        uint256 newTimeFrame = TIME_FRAME + 3600;
        bytes memory newData = abi.encode(uint256(600)); // 10 minutes timeInterval
        bytes32 newJobHash = keccak256(
            abi.encode(newJobName, JOB_TYPE, newTimeFrame, targetContract, newData)
        );

        vm.expectEmit(true, true, false, true);
        emit JobUpdated(jobId, user1, newJobHash, block.timestamp);

        jobRegistry.updateJob(
            jobId,
            JOB_NAME,        // oldJobName
            JOB_TYPE,        // jobType
            TIME_FRAME,      // oldTimeFrame
            targetContract,  // targetContract
            JOB_DATA,        // oldData
            newJobName,      // newJobName
            newTimeFrame,    // newTimeFrame
            newData          // newData
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
            999,                // jobId (non-existent)
            JOB_NAME,           // oldJobName
            JOB_TYPE,           // jobType
            TIME_FRAME,         // oldTimeFrame
            targetContract,     // targetContract
            JOB_DATA,           // oldData
            "Updated Job",      // newJobName
            TIME_FRAME + 3600,  // newTimeFrame
            JOB_DATA            // newData
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

        bytes memory newData = abi.encode(uint256(600)); // 10 minutes timeInterval
        vm.expectRevert(abi.encodeWithSelector(JobRegistry.UnauthorizedJobAccess.selector, jobId, user2));
        jobRegistry.updateJob(
            jobId,              // jobId
            JOB_NAME,           // oldJobName
            JOB_TYPE,           // jobType
            TIME_FRAME,         // oldTimeFrame
            targetContract,     // targetContract
            JOB_DATA,           // oldData
            "Updated Job",      // newJobName
            TIME_FRAME + 3600,  // newTimeFrame
            newData             // newData
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

    // ========== Job Data Validation Tests ==========

    function test_CreateJob_JobType1_ValidTimeInterval() public {
        vm.startPrank(user1);

        // JobType 1 requires uint256 timeInterval (first 32 bytes)
        bytes memory validData = abi.encode(uint256(300)); // 5 minutes

        uint256 jobId = jobRegistry.createJob(
            "Job Type 1",
            1,
            TIME_FRAME,
            targetContract,
            validData
        );

        assertEq(jobId, 1);
        vm.stopPrank();
    }

    function test_CreateJob_JobType1_InvalidTimeInterval() public {
        vm.startPrank(user1);

        // JobType 1 with zero timeInterval should revert
        bytes memory invalidData = abi.encode(uint256(0));

        vm.expectRevert(JobRegistry.MissingTimeInterval.selector);
        jobRegistry.createJob(
            "Job Type 1",
            1,
            TIME_FRAME,
            targetContract,
            invalidData
        );

        vm.stopPrank();
    }

    function test_CreateJob_JobType1_InsufficientData() public {
        vm.startPrank(user1);

        // JobType 1 with insufficient data length should revert
        bytes memory insufficientData = "0x1234"; // Less than 32 bytes

        vm.expectRevert(JobRegistry.MissingTimeInterval.selector);
        jobRegistry.createJob(
            "Job Type 1",
            1,
            TIME_FRAME,
            targetContract,
            insufficientData
        );

        vm.stopPrank();
    }

    function test_CreateJob_JobType2_ValidTimeIntervalAndIpfsHash() public {
        vm.startPrank(user1);

        // JobType 2 requires uint256 timeInterval + bytes32 ipfsHash
        bytes32 ipfsHash = keccak256(abi.encodePacked("test-ipfs-hash"));
        bytes memory validData = abi.encode(uint256(600), ipfsHash); // 10 minutes + ipfsHash

        uint256 jobId = jobRegistry.createJob(
            "Job Type 2",
            2,
            TIME_FRAME,
            targetContract,
            validData
        );

        assertEq(jobId, 1);
        vm.stopPrank();
    }

    function test_CreateJob_JobType2_MissingIpfsHash() public {
        vm.startPrank(user1);

        // JobType 2 with zero ipfsHash should revert
        bytes memory invalidData = abi.encode(uint256(600), bytes32(0));

        vm.expectRevert(JobRegistry.MissingIpfsHash.selector);
        jobRegistry.createJob(
            "Job Type 2",
            2,
            TIME_FRAME,
            targetContract,
            invalidData
        );

        vm.stopPrank();
    }

    function test_CreateJob_JobType2_InsufficientDataForIpfsHash() public {
        vm.startPrank(user1);

        // JobType 2 with insufficient data for ipfsHash should revert
        bytes memory insufficientData = abi.encode(uint256(600)); // Only timeInterval, no ipfsHash

        vm.expectRevert(JobRegistry.MissingIpfsHash.selector);
        jobRegistry.createJob(
            "Job Type 2",
            2,
            TIME_FRAME,
            targetContract,
            insufficientData
        );

        vm.stopPrank();
    }

    function test_CreateJob_JobType3_ValidRecurringJob() public {
        vm.startPrank(user1);

        // JobType 3 requires bool recurringJob (first 32 bytes)
        bytes memory validData = abi.encode(true);

        uint256 jobId = jobRegistry.createJob(
            "Job Type 3",
            3,
            TIME_FRAME,
            targetContract,
            validData
        );

        assertEq(jobId, 1);
        vm.stopPrank();
    }

    function test_CreateJob_JobType3_InsufficientData() public {
        vm.startPrank(user1);

        // JobType 3 with insufficient data should revert
        bytes memory insufficientData = "0x1234"; // Less than 32 bytes

        vm.expectRevert(JobRegistry.InvalidJobData.selector);
        jobRegistry.createJob(
            "Job Type 3",
            3,
            TIME_FRAME,
            targetContract,
            insufficientData
        );

        vm.stopPrank();
    }

    function test_CreateJob_JobType4_ValidRecurringJobAndIpfsHash() public {
        vm.startPrank(user1);

        // JobType 4 requires bool recurringJob + bytes32 ipfsHash
        bytes32 ipfsHash = keccak256(abi.encodePacked("test-ipfs-hash"));
        bytes memory validData = abi.encode(true, ipfsHash);

        uint256 jobId = jobRegistry.createJob(
            "Job Type 4",
            4,
            TIME_FRAME,
            targetContract,
            validData
        );

        assertEq(jobId, 1);
        vm.stopPrank();
    }

    function test_CreateJob_JobType4_MissingIpfsHash() public {
        vm.startPrank(user1);

        // JobType 4 with zero ipfsHash should revert
        bytes memory invalidData = abi.encode(true, bytes32(0));

        vm.expectRevert(JobRegistry.MissingIpfsHash.selector);
        jobRegistry.createJob(
            "Job Type 4",
            4,
            TIME_FRAME,
            targetContract,
            invalidData
        );

        vm.stopPrank();
    }

    function test_CreateJob_JobType6_ValidRecurringJobAndIpfsHash() public {
        vm.startPrank(user1);

        // JobType 6 requires bool recurringJob + bytes32 ipfsHash
        bytes32 ipfsHash = keccak256(abi.encodePacked("test-ipfs-hash"));
        bytes memory validData = abi.encode(false, ipfsHash);

        uint256 jobId = jobRegistry.createJob(
            "Job Type 6",
            6,
            TIME_FRAME,
            targetContract,
            validData
        );

        assertEq(jobId, 1);
        vm.stopPrank();
    }

    function test_CreateJob_JobType5_ValidRecurringJob() public {
        vm.startPrank(user1);

        // JobType 5 requires only bool recurringJob (no ipfsHash)
        bytes memory validData = abi.encode(false);

        uint256 jobId = jobRegistry.createJob(
            "Job Type 5",
            5,
            TIME_FRAME,
            targetContract,
            validData
        );

        assertEq(jobId, 1);
        vm.stopPrank();
    }

    // ========== UpdateJob Validation Tests ==========

    function test_UpdateJob_JobType1_ValidTimeInterval() public {
        vm.startPrank(user1);

        // Create job with valid data
        bytes memory oldData = abi.encode(uint256(300));
        uint256 jobId = jobRegistry.createJob(
            "Original Job",
            1,
            TIME_FRAME,
            targetContract,
            oldData
        );

        // Update with new valid timeInterval
        bytes memory newData = abi.encode(uint256(600));
        jobRegistry.updateJob(
            jobId,
            "Original Job",      // oldJobName
            1,                   // jobType
            TIME_FRAME,          // oldTimeFrame
            targetContract,      // targetContract
            oldData,             // oldData
            "Updated Job",       // newJobName
            TIME_FRAME + 3600,   // newTimeFrame
            newData              // newData
        );

        JobRegistry.Job memory job = jobRegistry.getJob(jobId);
        assertEq(job.lastUpdatedAt, block.timestamp);
        vm.stopPrank();
    }

    function test_UpdateJob_JobType1_InvalidTimeInterval() public {
        vm.startPrank(user1);

        // Create job with valid data
        bytes memory oldData = abi.encode(uint256(300));
        uint256 jobId = jobRegistry.createJob(
            "Original Job",
            1,
            TIME_FRAME,
            targetContract,
            oldData
        );

        // Try to update with invalid timeInterval (zero)
        bytes memory newData = abi.encode(uint256(0));
        vm.expectRevert(JobRegistry.MissingTimeInterval.selector);
        jobRegistry.updateJob(
            jobId,
            "Original Job",      // oldJobName
            1,                   // jobType
            TIME_FRAME,          // oldTimeFrame
            targetContract,      // targetContract
            oldData,             // oldData
            "Updated Job",       // newJobName
            TIME_FRAME + 3600,   // newTimeFrame
            newData              // newData
        );

        vm.stopPrank();
    }

    function test_UpdateJob_JobType2_ValidTimeIntervalAndIpfsHash() public {
        vm.startPrank(user1);

        // Create job with valid data
        bytes32 oldIpfsHash = keccak256(abi.encodePacked("old-ipfs-hash"));
        bytes memory oldData = abi.encode(uint256(300), oldIpfsHash);
        uint256 jobId = jobRegistry.createJob(
            "Original Job",
            2,
            TIME_FRAME,
            targetContract,
            oldData
        );

        // Update with new valid timeInterval (same ipfsHash)
        bytes memory newData = abi.encode(uint256(600), oldIpfsHash);
        jobRegistry.updateJob(
            jobId,
            "Original Job",      // oldJobName
            2,                   // jobType
            TIME_FRAME,          // oldTimeFrame
            targetContract,      // targetContract
            oldData,             // oldData
            "Updated Job",       // newJobName
            TIME_FRAME + 3600,   // newTimeFrame
            newData              // newData
        );

        JobRegistry.Job memory job = jobRegistry.getJob(jobId);
        assertEq(job.lastUpdatedAt, block.timestamp);
        vm.stopPrank();
    }

    function test_UpdateJob_JobType2_MissingIpfsHash() public {
        vm.startPrank(user1);

        // Create job with valid data
        bytes32 oldIpfsHash = keccak256(abi.encodePacked("old-ipfs-hash"));
        bytes memory oldData = abi.encode(uint256(300), oldIpfsHash);
        uint256 jobId = jobRegistry.createJob(
            "Original Job",
            2,
            TIME_FRAME,
            targetContract,
            oldData
        );

        // Try to update with missing ipfsHash
        bytes memory newData = abi.encode(uint256(600), bytes32(0));
        vm.expectRevert(JobRegistry.MissingIpfsHash.selector);
        jobRegistry.updateJob(
            jobId,
            "Original Job",      // oldJobName
            2,                   // jobType
            TIME_FRAME,          // oldTimeFrame
            targetContract,      // targetContract
            oldData,             // oldData
            "Updated Job",       // newJobName
            TIME_FRAME + 3600,   // newTimeFrame
            newData              // newData
        );

        vm.stopPrank();
    }

    function test_UpdateJob_JobType3_ValidRecurringJob() public {
        vm.startPrank(user1);

        // Create job with valid data
        bytes memory oldData = abi.encode(true);
        uint256 jobId = jobRegistry.createJob(
            "Original Job",
            3,
            TIME_FRAME,
            targetContract,
            oldData
        );

        // Update with new valid recurringJob
        bytes memory newData = abi.encode(false);
        jobRegistry.updateJob(
            jobId,
            "Original Job",      // oldJobName
            3,                   // jobType
            TIME_FRAME,          // oldTimeFrame
            targetContract,      // targetContract
            oldData,             // oldData
            "Updated Job",       // newJobName
            TIME_FRAME + 3600,   // newTimeFrame
            newData              // newData
        );

        JobRegistry.Job memory job = jobRegistry.getJob(jobId);
        assertEq(job.lastUpdatedAt, block.timestamp);
        vm.stopPrank();
    }

    function test_UpdateJob_JobType4_ValidRecurringJobAndIpfsHash() public {
        vm.startPrank(user1);

        // Create job with valid data
        bytes32 oldIpfsHash = keccak256(abi.encodePacked("old-ipfs-hash"));
        bytes memory oldData = abi.encode(true, oldIpfsHash);
        uint256 jobId = jobRegistry.createJob(
            "Original Job",
            4,
            TIME_FRAME,
            targetContract,
            oldData
        );

        // Update with new valid recurringJob (same ipfsHash)
        bytes memory newData = abi.encode(false, oldIpfsHash);
        jobRegistry.updateJob(
            jobId,
            "Original Job",      // oldJobName
            4,                   // jobType
            TIME_FRAME,          // oldTimeFrame
            targetContract,      // targetContract
            oldData,             // oldData
            "Updated Job",       // newJobName
            TIME_FRAME + 3600,   // newTimeFrame
            newData              // newData
        );

        JobRegistry.Job memory job = jobRegistry.getJob(jobId);
        assertEq(job.lastUpdatedAt, block.timestamp);
        vm.stopPrank();
    }

    function test_UpdateJob_OldValuesMismatch() public {
        vm.startPrank(user1);

        // Create job with valid data
        bytes memory oldData = abi.encode(uint256(300));
        uint256 jobId = jobRegistry.createJob(
            "Original Job",
            1,
            TIME_FRAME,
            targetContract,
            oldData
        );

        // Try to update with wrong old values
        bytes memory newData = abi.encode(uint256(600));
        vm.expectRevert("OLD_VALUES_MISMATCH");
        jobRegistry.updateJob(
            jobId,
            "Wrong Old Name",    // oldJobName (wrong)
            1,                   // jobType
            TIME_FRAME,          // oldTimeFrame
            targetContract,      // targetContract
            oldData,             // oldData
            "Updated Job",       // newJobName
            TIME_FRAME + 3600,   // newTimeFrame
            newData              // newData
        );

        vm.stopPrank();
    }

    function test_UpdateJob_EmptyJobName() public {
        vm.startPrank(user1);

        // Create job with valid data
        bytes memory oldData = abi.encode(uint256(300));
        uint256 jobId = jobRegistry.createJob(
            "Original Job",
            1,
            TIME_FRAME,
            targetContract,
            oldData
        );

        // Try to update with empty job name
        bytes memory newData = abi.encode(uint256(600));
        vm.expectRevert(JobRegistry.EmptyJobName.selector);
        jobRegistry.updateJob(
            jobId,
            "Original Job",      // oldJobName
            1,                   // jobType
            TIME_FRAME,          // oldTimeFrame
            targetContract,      // targetContract
            oldData,             // oldData
            "",                  // newJobName (empty)
            TIME_FRAME + 3600,   // newTimeFrame
            newData              // newData
        );

        vm.stopPrank();
    }

    // ========== Edge Cases and Boundary Tests ==========

    function test_CreateJob_MaxTimeInterval() public {
        vm.startPrank(user1);

        // Test with maximum uint256 timeInterval
        bytes memory validData = abi.encode(type(uint256).max);

        uint256 jobId = jobRegistry.createJob(
            "Max Time Interval Job",
            1,
            TIME_FRAME,
            targetContract,
            validData
        );

        assertEq(jobId, 1);
        vm.stopPrank();
    }

    function test_CreateJob_MinTimeInterval() public {
        vm.startPrank(user1);

        // Test with minimum non-zero timeInterval
        bytes memory validData = abi.encode(uint256(1));

        uint256 jobId = jobRegistry.createJob(
            "Min Time Interval Job",
            1,
            TIME_FRAME,
            targetContract,
            validData
        );

        assertEq(jobId, 1);
        vm.stopPrank();
    }

    function test_CreateJob_ComplexIpfsHash() public {
        vm.startPrank(user1);

        // Test with complex ipfsHash
        bytes32 complexIpfsHash = keccak256(abi.encodePacked("very-long-complex-ipfs-hash-with-special-characters-!@#$%^&*()"));
        bytes memory validData = abi.encode(uint256(300), complexIpfsHash);

        uint256 jobId = jobRegistry.createJob(
            "Complex IPFS Job",
            2,
            TIME_FRAME,
            targetContract,
            validData
        );

        assertEq(jobId, 1);
        vm.stopPrank();
    }

    function test_CreateJob_AllJobTypes() public {
        vm.startPrank(user1);

        // Test all job types in sequence
        bytes32 ipfsHash = keccak256(abi.encodePacked("test-ipfs"));

        // JobType 1
        bytes memory data1 = abi.encode(uint256(100));
        uint256 jobId1 = jobRegistry.createJob("Job Type 1", 1, TIME_FRAME, targetContract, data1);

        // JobType 2
        bytes memory data2 = abi.encode(uint256(200), ipfsHash);
        uint256 jobId2 = jobRegistry.createJob("Job Type 2", 2, TIME_FRAME, targetContract, data2);

        // JobType 3
        bytes memory data3 = abi.encode(true);
        uint256 jobId3 = jobRegistry.createJob("Job Type 3", 3, TIME_FRAME, targetContract, data3);

        // JobType 4
        bytes memory data4 = abi.encode(false, ipfsHash);
        uint256 jobId4 = jobRegistry.createJob("Job Type 4", 4, TIME_FRAME, targetContract, data4);

        // JobType 5
        bytes memory data5 = abi.encode(true);
        uint256 jobId5 = jobRegistry.createJob("Job Type 5", 5, TIME_FRAME, targetContract, data5);

        // JobType 6
        bytes memory data6 = abi.encode(false, ipfsHash);
        uint256 jobId6 = jobRegistry.createJob("Job Type 6", 6, TIME_FRAME, targetContract, data6);

        assertEq(jobId1, 1);
        assertEq(jobId2, 2);
        assertEq(jobId3, 3);
        assertEq(jobId4, 4);
        assertEq(jobId5, 5);
        assertEq(jobId6, 6);
        assertEq(jobRegistry.getTotalJobsCount(), 6);

        vm.stopPrank();
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
        bytes32 ipfsHash = keccak256(abi.encodePacked("user2-ipfs-hash"));
        bytes memory jobData2 = abi.encode(uint256(300), ipfsHash); // timeInterval + ipfsHash for JobType 2
        uint256 jobId2 = jobRegistry.createJob(
            "User2 Job",
            JOB_TYPE + uint8(1),
            TIME_FRAME,
            targetContract,
            jobData2
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

    // Helper function to calculate job hash with custom data
    function _calculateJobHash(string memory jobName, bytes memory data) internal view returns (bytes32) {
        return keccak256(
            abi.encode(jobName, JOB_TYPE, TIME_FRAME, targetContract, data)
        );
    }
} 