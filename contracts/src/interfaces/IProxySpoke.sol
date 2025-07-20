// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

/**
 * @title IProxySpoke
 * @notice Interface for the TaskExecutionSpoke contract that acts as a spoke in the keeper network
 */
interface IProxySpoke {
    /**
     * @notice Enum defining the types of actions that can be performed on keepers
     */
    enum ActionType { REGISTER, UNREGISTER }

    /**
     * @notice Emitted when a keeper's status is updated
     * @param action The type of action performed
     * @param keeper The address of the keeper
     */
    event KeeperUpdated(ActionType action, address keeper);

    /**
     * @notice Emitted when a function is executed
     * @param keeper The address of the keeper who executed the function
     * @param target The address of the target contract
     * @param data The calldata used for the function call
     * @param value The amount of ETH sent with the call
     */
    event FunctionExecuted(address indexed keeper, address indexed target, bytes data, uint256 value);

    /**
     * @notice Executes a function on a target contract
     * @param target The address of the target contract
     * @param data The calldata for the function call
     */
    function executeFunction(address target, bytes memory data) external payable;

    /**
     * @notice Returns whether an address is a registered keeper
     * @param keeper The address to check
     * @return bool True if the address is a keeper, false otherwise
     */
    function isKeeper(address keeper) external view returns (bool);
} 