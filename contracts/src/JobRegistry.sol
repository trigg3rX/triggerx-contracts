// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Initializable} from "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import {UUPSUpgradeable} from "@openzeppelin-upgrades/contracts/proxy/utils/UUPSUpgradeable.sol";
import {OwnableUpgradeable} from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import {PackedJobIdLib} from "./libraries/PackedJobIdLib.sol";

/**
 * @title JobRegistry
 * @dev A UUPS upgradeable contract for managing job registrations
 * @notice This contract allows users to create, update, and delete jobs with associated metadata
 */
contract JobRegistry is Initializable, UUPSUpgradeable, OwnableUpgradeable {
    // State variables
    uint256 private _lastJobCounter;
    mapping(uint256 => Job) private _jobs;
    mapping(address => uint256[]) private _userJobIds;

    // Structs
    struct Job {
        uint256 jobId;
        address jobOwner;
        bytes32 jobHash;
        uint256 lastUpdatedAt;
        bool isActive;
    }

    // Events
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

    // Custom errors
    error JobNotFound(uint256 jobId);
    error UnauthorizedJobAccess(uint256 jobId, address caller);
    error JobAlreadyInactive(uint256 jobId);
    error InvalidJobParameters();
    error EmptyJobName();
    error InvalidTargetContract();
    error InvalidJobData();
    error MissingTimeInterval();
    error MissingIpfsHash();

    /**
     * @dev Constructor that disables initializers to prevent implementation contract from being initialized
     */
    constructor() {
        _disableInitializers();
    }

    /**
     * @dev Initialize the contract
     * @param initialOwner The address that will be set as the initial owner
     */
    function initialize(address initialOwner) public initializer {
        if (initialOwner == address(0)) {
            revert InvalidJobParameters();
        }

        __Ownable_init(initialOwner);
        _lastJobCounter = 0;
    }

    /**
     * @dev Create a new job
     * @param jobName The name of the job
     * @param jobType The type/category of the job
     * @param timeFrame The timeframe for job execution
     * @param targetContract The address of the target contract
     * @param data Additional data for the job
     * @return jobId The ID of the created job
     */
    function createJob(
        string memory jobName,
        uint8 jobType,
        uint256 timeFrame,
        address targetContract,
        bytes memory data
    ) external returns (uint256 jobId) {
        if (bytes(jobName).length == 0) revert EmptyJobName();
        if (targetContract == address(0)) revert InvalidTargetContract(); 

        // Validate data based on jobType
        _validateJobData(jobType, data);

        uint256 jobCounter = ++_lastJobCounter;
        jobId = PackedJobIdLib.pack(block.chainid, jobCounter);

        bytes32 jobHash = keccak256(
            abi.encode(jobName, jobType, timeFrame, targetContract, data)
        );

        Job memory newJob = Job({
            jobId: jobId,
            jobOwner: msg.sender,
            jobHash: jobHash,
            lastUpdatedAt: block.timestamp,
            isActive: true
        });

        _jobs[jobId] = newJob;
        _userJobIds[msg.sender].push(jobId);

        emit JobCreated(jobId, msg.sender, jobHash, block.timestamp);
    }

    /**
     * @dev Update an existing job
     * @param jobId The ID of the job to update
     * @param oldJobName The old name of the job
     * @param jobType The old type/category of the job
     * @param oldTimeFrame The old timeframe for job execution
     * @param targetContract The old address of the target contract
     * @param oldData The old additional data for the job
     * @param newJobName The new name of the job
     * @param newTimeFrame The new timeframe for job execution
     * @param newData The new additional data for the job
     */
    function updateJob(
        uint256 jobId,
        string memory oldJobName,
        uint8 jobType,
        uint256 oldTimeFrame,
        address targetContract,
        bytes memory oldData,
        string memory newJobName,
        uint256 newTimeFrame,
        bytes memory newData
    ) external {
        Job storage job = _jobs[jobId];

        if (job.jobOwner == address(0)) revert JobNotFound(jobId);
        if (job.jobOwner != msg.sender)
            revert UnauthorizedJobAccess(jobId, msg.sender);
        if (!job.isActive) revert JobAlreadyInactive(jobId);

        bytes32 currentJobHash = keccak256(
            abi.encode(
                oldJobName,
                jobType,
                oldTimeFrame,
                targetContract,
                oldData
            )
        );
        if (currentJobHash != job.jobHash) revert("OLD_VALUES_MISMATCH");

        if (bytes(newJobName).length == 0) revert EmptyJobName();
        if (targetContract == address(0)) revert InvalidTargetContract();

        // Validate data based on jobType
        _validateJobData(jobType, newData);

        bytes32 newJobHash = keccak256(
            abi.encode(
                newJobName,
                jobType,
                newTimeFrame,
                targetContract,
                newData
            )
        );

        job.jobHash = newJobHash;
        job.lastUpdatedAt = block.timestamp;

        emit JobUpdated(jobId, msg.sender, newJobHash, block.timestamp);
    }

    /**
     * @dev Validate job data based on jobType requirements
     * @param jobType The type of the job
     * @param data The data to validate
     */
    function _validateJobData(uint8 jobType, bytes memory data) internal pure {
        // For jobType 1 or 2, require uint256 timeInterval
        if (jobType == 1 || jobType == 2) {
            if (data.length < 32) {
                revert MissingTimeInterval();
            }
            // Extract timeInterval (first 32 bytes)
            uint256 timeInterval;
            assembly {
                timeInterval := mload(add(data, 32))
            }
            if (timeInterval == 0) {
                revert MissingTimeInterval();
            }
        } else {
            // For other jobTypes, require bool recurringJob
            if (data.length < 32) {
                revert InvalidJobData();
            }
            // Extract recurringJob (first 32 bytes as bool)
            bool recurringJob;
            assembly {
                recurringJob := mload(add(data, 32))
            }
        }

        // For jobType 2, 4, or 6, require bytes32 ipfsHash
        if (jobType == 2 || jobType == 4 || jobType == 6) {
            if (data.length < 64) {
                // Need 32 bytes for timeInterval/recurringJob + 32 bytes for ipfsHash
                revert MissingIpfsHash();
            }
            // Extract ipfsHash (second 32 bytes)
            bytes32 ipfsHash;
            assembly {
                ipfsHash := mload(add(data, 64))
            }
            if (ipfsHash == bytes32(0)) {
                revert MissingIpfsHash();
            }
        }
    }

    /**
     * @dev Delete (deactivate) a job
     * @param jobId The ID of the job to delete
     */
    function deleteJob(uint256 jobId) external {
        Job storage job = _jobs[jobId];

        if (job.jobOwner == address(0)) {
            revert JobNotFound(jobId);
        }
        if (job.jobOwner != msg.sender) {
            revert UnauthorizedJobAccess(jobId, msg.sender);
        }
        if (!job.isActive) {
            revert JobAlreadyInactive(jobId);
        }

        job.isActive = false;

        emit JobDeleted(jobId, msg.sender, block.timestamp);
    }

    // View functions

    /**
     * @dev Get job details by ID
     * @param jobId The ID of the job
     * @return job The job details
     */
    function getJob(uint256 jobId) external view returns (Job memory job) {
        job = _jobs[jobId];
        if (job.jobOwner == address(0)) {
            revert JobNotFound(jobId);
        }
    }

    function getJobByCounter(uint256 jobCounter) external view returns (Job memory job) {
        job = _jobs[PackedJobIdLib.pack(block.chainid, jobCounter)];
        if (job.jobOwner == address(0)) {
            revert JobNotFound(PackedJobIdLib.pack(block.chainid, jobCounter));
        }
    }

    /**
     * @dev Get the owner of a job
     * @param jobId The ID of the job
     * @return jobOwner The address of the job owner
     */
    function getJobOwner(uint256 jobId) external view returns (address) {
        return _jobs[jobId].jobOwner;
    }

    /**
     * @dev Get all job IDs for a user
     * @param user The address of the user
     * @return jobIds Array of job IDs owned by the user
     */
    function getUserJobIds(
        address user
    ) external view returns (uint256[] memory jobIds) {
        return _userJobIds[user];
    }

    /**
     * @dev Get all active job IDs for a user
     * @param user The address of the user
     * @return activeJobIds Array of active job IDs owned by the user
     */
    function getUserActiveJobIds(
        address user
    ) external view returns (uint256[] memory activeJobIds) {
        uint256[] memory allJobIds = _userJobIds[user];
        uint256 activeCount = 0;

        // Count active jobs
        for (uint256 i = 0; i < allJobIds.length; i++) {
            if (_jobs[allJobIds[i]].isActive) {
                activeCount++;
            }
        }

        // Create array with active jobs
        activeJobIds = new uint256[](activeCount);
        uint256 currentIndex = 0;
        for (uint256 i = 0; i < allJobIds.length; i++) {
            if (_jobs[allJobIds[i]].isActive) {
                activeJobIds[currentIndex] = allJobIds[i];
                currentIndex++;
            }
        }
    }

    /**
     * @dev Check if a job is active
     * @param jobId The ID of the job
     * @return isActive True if the job is active, false otherwise
     */
    function isJobActive(uint256 jobId) external view returns (bool isActive) {
        Job memory job = _jobs[jobId];
        if (job.jobOwner == address(0)) {
            revert JobNotFound(jobId);
        }
        return job.isActive;
    }

    /**
     * @dev Get the total number of jobs created
     * @return count The total number of jobs
     */
    function getJobCounter() external view returns (uint256 count) {
        return _lastJobCounter;
    }

    /**
     * @dev Unpack a jobId to get the chainId and jobCounter.
     * @param jobId The packed job ID.
     * @return chainId The chain ID from the packed job ID.
     * @return jobCounter The job counter from the packed job ID.
     */
    function unpackJobId(uint256 jobId) external pure returns (uint256 chainId, uint256 jobCounter) {
        (chainId, jobCounter) = PackedJobIdLib.unpack(jobId);
    }

    // UUPS upgrade authorization
    /**
     * @dev Authorize contract upgrades - only owner can upgrade
     * @param newImplementation The address of the new implementation
     */
    function _authorizeUpgrade(
        address newImplementation
    ) internal override onlyOwner {}
}
