// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

/**
 * TriggerX Safe Module
 * - Trusted TaskExecutionHub calls execJobFromHub
 * - Module executes the user action via SAFE.execTransactionFromModule(...)
 * - Verifies Safe ownership belongs to the job owner (validated by TaskExecutionHub)
 */

import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

interface IGnosisSafe {
    function execTransactionFromModule(
        address to,
        uint256 value,
        bytes calldata data,
        uint8 operation
    ) external returns (bool success);

    function isOwner(
        address owner
    ) external view returns (bool);
}

/**
 * @title TriggerXSafeModule
 * @notice Safe module with job ownership verification
 * @dev JobOwner validation is now done by TaskExecutionHub before calling this module
 */
contract TriggerXSafeModule is ReentrancyGuard {
    // State variables
    address public immutable taskExecutionHub;

    // Events
    event TaskExecutedFromModule(
        address indexed jobOwner,
        address indexed safeAddress,
        address indexed executor,
        bool success
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
     */
    constructor(
        address _taskExecutionHub
    ) {
        if (_taskExecutionHub == address(0)) {
            revert InvalidAddress();
        }
        taskExecutionHub = _taskExecutionHub;
    }

    /**
     * @notice Execute a job from the TaskExecutionHub with ownership verification
     * @param jobOwner The validated job owner address (validated by TaskExecutionHub)
     * @param safeAddress The Safe contract address
     * @param actionTarget Target contract the user wants to call
     * @param actionValue ETH value for the action
     * @param actionData Calldata for the actionTarget
     * @param operation 0 = CALL, 1 = DELEGATECALL
     * @return success Whether the execution was successful
     */
    function execJobFromHub(
        address safeAddress,
        address actionTarget,
        uint256 actionValue,
        bytes calldata actionData,
        uint8 operation,
        address jobOwner
    ) external nonReentrant onlyHub returns (bool success) {
        // CRITICAL: Verify Safe belongs to job owner
        IGnosisSafe safe = IGnosisSafe(safeAddress);

        if (!safe.isOwner(jobOwner)) {
            revert SafeNotOwnedByJobOwner(safeAddress, jobOwner);
        }

        bool ok;
        try safe.execTransactionFromModule(
            actionTarget, actionValue, actionData, operation
        ) returns (
            bool _success
        ) {
            ok = _success;
        } catch {
            ok = false;
        }

        if (!ok) {
            emit TaskExecutedFromModule(jobOwner, safeAddress, tx.origin, false);
            revert ExecFailed();
        }

        emit TaskExecutedFromModule(jobOwner, safeAddress, tx.origin, true);

        return true;
    }
}
