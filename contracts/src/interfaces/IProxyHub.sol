// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

/**
 * @title ITaskExecutionHub
 * @notice Interface for the TaskExecutionHub contract that manages keeper registration and function execution across multiple chains
 */
interface ITaskExecutionHub {
    /**
     * @notice Enum defining the types of actions that can be performed on keepers
     */
    enum ActionType { REGISTER, UNREGISTER }

    /**
     * @notice Emitted when a keeper is registered
     * @param keeper The address of the registered keeper
     */
    event KeeperRegistered(address indexed keeper);

    /**
     * @notice Emitted when a keeper is unregistered
     * @param keeper The address of the unregistered keeper
     */
    event KeeperUnregistered(address indexed keeper);

    /**
     * @notice Emitted when a broadcast message is sent
     * @param action The type of action being broadcast
     * @param keeper The address of the keeper
     * @param dstEid The destination chain endpoint ID
     */
    event BroadcastSent(ActionType action, address keeper, uint32 dstEid);

    /**
     * @notice Emitted when a function is executed
     * @param keeper The address of the keeper who executed the function
     * @param target The address of the target contract
     * @param data The calldata used for the function call
     * @param value The amount of ETH sent with the call
     */
    event FunctionExecuted(address indexed keeper, address indexed target, bytes data, uint256 value);

    /**
     * @notice Emitted when fees are used for a cross-chain message
     * @param dstEid The destination chain endpoint ID
     * @param fee The amount of fees used
     */
    event FeeUsed(uint32 dstEid, uint256 fee);

    /**
     * @notice Emitted when gas configuration is updated
     * @param gas The new gas limit
     * @param value The new default value
     */
    event GasConfigUpdated(uint128 gas, uint128 value);

    /**
     * @notice Emitted when contract balance is low
     * @param currentBalance The current contract balance
     * @param threshold The low balance threshold
     */
    event LowBalanceAlert(uint256 currentBalance, uint256 threshold);

    /**
     * @notice Emitted when a message fails to send
     * @param dstEid The destination chain endpoint ID
     * @param guid The message GUID (bytes32(0) if not available)
     * @param reason The reason for the failure
     */
    event MessageFailed(uint32 indexed dstEid, bytes32 indexed guid, bytes reason);

    /**
     * @notice Adds new spoke chains to the network
     * @param _dstEids Array of destination chain endpoint IDs
     */
    function addSpokes(uint32[] calldata _dstEids) external;

    /**
     * @notice Updates the gas configuration for cross-chain messages
     * @param gas The new gas limit
     * @param value The new default value
     */
    function setGasConfig(uint128 gas, uint128 value) external;

    /**
     * @notice Executes a function on a target contract
     * @param target The address of the target contract
     * @param data The calldata for the function call
     */
    function executeFunction(address target, bytes calldata data) external payable;

    /**
     * @notice Allows the owner to withdraw ETH from the contract
     * @param to The address to send the ETH to
     * @param amount The amount of ETH to withdraw
     */
    function withdraw(address payable to, uint256 amount) external;

    /**
     * @notice Returns whether an address is a registered keeper
     * @param keeper The address to check
     * @return bool True if the address is a keeper, false otherwise
     */
    function isKeeper(address keeper) external view returns (bool);

    /**
     * @notice Returns the LayerZero endpoint ID of this chain
     * @return uint32 The endpoint ID
     */
    function thisChainEid() external view returns (uint32);

    /**
     * @notice Returns the LayerZero endpoint ID of the source chain
     * @return uint32 The endpoint ID
     */
    function srcEid() external view returns (uint32);

    /**
     * @notice Returns the array of destination chain endpoint IDs
     * @return uint32[] Array of endpoint IDs
     */
    function dstEids() external view returns (uint32[] memory);

    /**
     * @notice Returns the default gas limit for cross-chain messages
     * @return uint128 The gas limit
     */
    function defaultGas() external view returns (uint128);

    /**
     * @notice Returns the default value to be sent with cross-chain messages
     * @return uint128 The default value
     */
    function defaultValue() external view returns (uint128);
} 