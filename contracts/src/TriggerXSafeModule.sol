// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

/**
 * TriggerX Safe Module
 * - Trusted TaskExecutionHub calls execJobFromHub
 * - Module executes the user action via SAFE.execTransactionFromModule(...)
 */

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

interface IGnosisSafe {
    function execTransactionFromModule(
        address to,
        uint256 value,
        bytes calldata data,
        uint8 operation
    ) external returns (bool success);
}

contract TriggerXSafeModule is ReentrancyGuard {
    using ECDSA for bytes32;

    // Addresses
    address public immutable taskExecutionHub;

    event TaskExecutedFromModule(address indexed safeAddress, address indexed executor, bool success);

    error NotTaskExecutionHub();
    error ExecFailed();

    modifier onlyHub() {
        if (msg.sender != taskExecutionHub) revert NotTaskExecutionHub();
        _;
    }

    constructor(address _taskExecutionHub) {
        require(_taskExecutionHub != address(0), "zero hub");
        taskExecutionHub = _taskExecutionHub;
    }

    /**
     * execJobFromHub
     * Called by TaskExecutionHub. Validates jobHash from JobRegistry and then instructs the Safe
     * to perform the desired action via execTransactionFromModule.
     *
     * @param safeAddress - the Safe contract address
     * @param actionTarget - target contract the user wants to call
     * @param actionValue - ETH value for the action
     * @param actionData - calldata for the actionTarget
     * @param operation - 0 = CALL, 1 = DELEGATECALL
     */
    function execJobFromHub(
        address safeAddress,
        address actionTarget,
        uint256 actionValue,
        bytes calldata actionData,
        uint8 operation
    ) external nonReentrant onlyHub returns (bool success) {

        IGnosisSafe safe = IGnosisSafe(safeAddress);
    
        bool ok;
        try safe.execTransactionFromModule(actionTarget, actionValue, actionData, operation) returns (bool _success) {
            ok = _success;
        } catch {
            ok = false;
        }

        if (!ok) {
            emit TaskExecutedFromModule(safeAddress, tx.origin, false);
            revert ExecFailed();
        }

        emit TaskExecutedFromModule(safeAddress, tx.origin, true);
        return true;
    }

    receive() external payable {}
    fallback() external payable {}
}
