// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

/**
 * @title IAvsGovernanceLogic
 * @notice Interface for the AvsGovernanceLogic contract that manages operator registration and cross-chain communication
 */
interface IAvsGovernanceLogic {
    /**
     * @notice Enum defining the types of actions that can be performed on operators
     */
    enum ActionType { REGISTER, UNREGISTER }

    /**
     * @notice Emitted when an operator is registered
     * @param operator The address of the registered operator
     */
    event OperatorRegistered(address indexed operator);

    /**
     * @notice Emitted when an operator is unregistered
     * @param operator The address of the unregistered operator
     */
    event OperatorUnregistered(address indexed operator);

    /**
     * @notice Emitted when contract balance is low
     * @param currentBalance The current contract balance
     * @param threshold The low balance threshold
     */
    event LowBalanceAlert(uint256 currentBalance, uint256 threshold);

    /**
     * @notice Emitted when a message is sent to another chain
     * @param dstEid The destination chain endpoint ID
     * @param guid The message GUID
     * @param fee The fee paid for the message
     */
    event MessageSent(uint32 indexed dstEid, bytes32 indexed guid, uint256 fee);

    /**
     * @notice Emitted when a message fails to send
     * @param dstEid The destination chain endpoint ID
     * @param guid The message GUID
     * @param reason The reason for the failure
     */
    event MessageFailed(uint32 indexed dstEid, bytes32 indexed guid, bytes reason);

    /**
     * @notice Emitted when an operator is whitelisted
     * @param operator The address of the whitelisted operator
     */
    event Whitelisted(address indexed operator);

    /**
     * @notice Emitted when an operator is removed from the whitelist
     * @param operator The address of the unwhitelisted operator
     */
    event UnWhitelisted(address indexed operator);

    /**
     * @notice Emitted when the gas options are updated
     * @param gasLimit The new gas limit
     * @param callValue The new call value
     */
    event GasOptionsUpdated(uint128 gasLimit, uint128 callValue);

    /**
     * @notice Hook called before an operator is registered
     * @param _operator The address of the operator to be registered
     * @param _votingPower The voting power of the operator
     * @param _blsKey The BLS key of the operator
     * @param _rewardsReceiver The address that will receive rewards
     */
    function beforeOperatorRegistered(
        address _operator,
        uint256 _votingPower,
        uint256[4] calldata _blsKey,
        address _rewardsReceiver
    ) external;

    /**
     * @notice Hook called after an operator is registered
     * @param _operator The address of the registered operator
     * @param _votingPower The voting power of the operator
     * @param _blsKey The BLS key of the operator
     * @param _rewardsReceiver The address that will receive rewards
     */
    function afterOperatorRegistered(
        address _operator,
        uint256 _votingPower,
        uint256[4] calldata _blsKey,
        address _rewardsReceiver
    ) external;

    /**
     * @notice Hook called before an operator is unregistered
     * @param _operator The address of the operator to be unregistered
     */
    function beforeOperatorUnregistered(address _operator) external;

    /**
     * @notice Hook called after an operator is unregistered
     * @param _operator The address of the unregistered operator
     */
    function afterOperatorUnregistered(address _operator) external;

    /**
     * @notice Updates the proxy hub address
     * @param _taskExecutionHub The new proxy hub address
     */
    function setTaskExecutionHub(address _taskExecutionHub) external;

    /**
     * @notice Allows the owner to withdraw ETH from the contract
     * @param _to The address to send the ETH to
     * @param _amount The amount of ETH to withdraw
     */
    function withdraw(address payable _to, uint256 _amount) external;

    /**
     * @notice Updates the gas configuration for cross-chain messages
     * @param _gasLimit The new gas limit
     * @param _callValue The new call value
     */
    function setGasOptions(uint128 _gasLimit, uint128 _callValue) external;

    /**
     * @notice Adds multiple operators to the whitelist
     * @param _operators Array of operator addresses to whitelist
     */
    function addToWhitelist(address[] calldata _operators) external;

    /**
     * @notice Removes multiple operators from the whitelist
     * @param _operators Array of operator addresses to remove from whitelist
     */
    function removeFromWhitelist(address[] calldata _operators) external;

    /**
     * @notice Returns whether an address is whitelisted
     * @param _operator The address to check
     * @return bool True if the address is whitelisted, false otherwise
     */
    function isWhitelisted(address _operator) external view returns (bool);

    /**
     * @notice Returns the address of the proxy hub
     * @return address The proxy hub address
     */
    function taskExecutionHub() external view returns (address);

    /**
     * @notice Returns the LayerZero endpoint ID of the destination chain
     * @return uint32 The endpoint ID
     */
    function dstEid() external view returns (uint32);

    /**
     * @notice Returns the default gas limit for cross-chain messages
     * @return uint128 The gas limit
     */
    function gasLimit() external view returns (uint128);

    /**
     * @notice Returns the default value to be sent with cross-chain messages
     * @return uint128 The call value
     */
    function callValue() external view returns (uint128);
} 