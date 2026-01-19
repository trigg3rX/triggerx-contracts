// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

/**
 * TriggerX Safe Module
 * - Trusted TaskExecutionHub calls execJobFromHub
 * - Module executes the user action via SAFE.execTransactionFromModule(...)
 * - Verifies Safe ownership belongs to the job owner
 */

import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

interface IGnosisSafe {
    function execTransactionFromModule(
        address to,
        uint256 value,
        bytes calldata data,
        uint8 operation
    ) external returns (bool success);

    function isOwner(address owner) external view returns (bool);
}

interface IJobRegistry {
    function getJobOwner(uint256 jobId) external view returns (address);
}

/**
 * @title TriggerXSafeModule
 * @notice Safe module with job ownership verification
 * @dev This module verifies that the Safe being acted upon belongs to the job owner
 */
contract TriggerXSafeModule is ReentrancyGuard, Ownable {
    // State variables
    address public taskExecutionHub;
    IJobRegistry public jobRegistry;

    // Events
    event TaskExecutedFromModule(
        uint256 indexed jobId,
        address indexed safeAddress,
        address indexed executor,
        bool success
    );
    event TaskExecutionHubUpdated(
        address indexed oldHub,
        address indexed newHub
    );
    event JobRegistryUpdated(
        address indexed oldRegistry,
        address indexed newRegistry
    );

    // Errors
    error NotTaskExecutionHub();
    error ExecFailed();
    error SafeNotOwnedByJobOwner(address safeAddress, address jobOwner);
    error InvalidAddress();

    modifier onlyHub() {
        if (msg.sender != taskExecutionHub) revert NotTaskExecutionHub();
        _;
    }

    /**
     * @notice Constructor
     * @param _taskExecutionHub The address of the TaskExecutionHub
     * @param _jobRegistry The address of the JobRegistry
     * @param _owner The owner of this contract
     */
    constructor(
        address _taskExecutionHub,
        address _jobRegistry,
        address _owner
    ) Ownable(_owner) {
        if (_taskExecutionHub == address(0) || _jobRegistry == address(0)) {
            revert InvalidAddress();
        }
        taskExecutionHub = _taskExecutionHub;
        jobRegistry = IJobRegistry(_jobRegistry);
    }

    /**
     * @notice Execute a job from the TaskExecutionHub with ownership verification
     * @param jobId The ID of the job being executed
     * @param safeAddress The Safe contract address
     * @param actionTarget Target contract the user wants to call
     * @param actionValue ETH value for the action
     * @param actionData Calldata for the actionTarget
     * @param operation 0 = CALL, 1 = DELEGATECALL
     * @return success Whether the execution was successful
     */
    function execJobFromHub(
        uint256 jobId,
        address safeAddress,
        address actionTarget,
        uint256 actionValue,
        bytes calldata actionData,
        uint8 operation
    ) external nonReentrant onlyHub returns (bool success) {
        // CRITICAL: Verify Safe belongs to job owner
        address jobOwner = jobRegistry.getJobOwner(jobId);
        IGnosisSafe safe = IGnosisSafe(safeAddress);

        if (!safe.isOwner(jobOwner)) {
            revert SafeNotOwnedByJobOwner(safeAddress, jobOwner);
        }

        bool ok;
        try
            safe.execTransactionFromModule(
                actionTarget,
                actionValue,
                actionData,
                operation
            )
        returns (bool _success) {
            ok = _success;
        } catch {
            ok = false;
        }

        if (!ok) {
            emit TaskExecutedFromModule(jobId, safeAddress, tx.origin, false);
            revert ExecFailed();
        }

        emit TaskExecutedFromModule(jobId, safeAddress, tx.origin, true);
        return true;
    }

    // Admin functions

    /**
     * @notice Update the TaskExecutionHub address
     * @param _taskExecutionHub New TaskExecutionHub address
     */
    function setTaskExecutionHub(address _taskExecutionHub) external onlyOwner {
        if (_taskExecutionHub == address(0)) revert InvalidAddress();
        address oldHub = taskExecutionHub;
        taskExecutionHub = _taskExecutionHub;
        emit TaskExecutionHubUpdated(oldHub, _taskExecutionHub);
    }

    /**
     * @notice Update the JobRegistry address
     * @param _jobRegistry New JobRegistry address
     */
    function setJobRegistry(address _jobRegistry) external onlyOwner {
        if (_jobRegistry == address(0)) revert InvalidAddress();
        address oldRegistry = address(jobRegistry);
        jobRegistry = IJobRegistry(_jobRegistry);
        emit JobRegistryUpdated(oldRegistry, _jobRegistry);
    }
}
