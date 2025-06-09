// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {IAvsGovernanceLogic} from "@othentic-core-contracts/NetworkManagement/L1/interfaces/IAvsGovernanceLogic.sol";
import {OApp, MessagingFee, Origin, MessagingReceipt} from "@layerzero-v2/oapp/contracts/oapp/OApp.sol";
import {ILayerZeroEndpointV2, MessagingParams} from "@layerzerolabs/lz-evm-protocol-v2/contracts/interfaces/ILayerZeroEndpointV2.sol";
import {Ownable} from "@openzeppelin-contracts/contracts/access/Ownable.sol";

/**
 * @title AvsGovernanceLogic
 * @notice A LayerZero-enabled contract that manages operator registration and cross-chain communication
 * @dev This contract handles operator registration and broadcasts updates to the L2 network
 * @custom:security-contact security@triggerx.com
 */
contract AvsGovernanceLogic is Ownable, IAvsGovernanceLogic, OApp {
    /**
     * @notice Enum defining the types of actions that can be performed on operators
     */
    enum ActionType {
        REGISTER,
        UNREGISTER
    }

    /**
     * @notice Structure to store failed message attempts for retry
     */
    struct FailedMessage {
        ActionType action;
        address operator;
        uint256 timestamp;
        uint8 retryCount;
    }

    /**
     * @notice The address of the L2 proxy hub contract
     */
    address public proxyHub;

    /**
     * @notice The LayerZero endpoint ID of the destination chain
     */
    uint32 public immutable dstEid;

    /**
     * @notice Default gas limit for cross-chain messages
     */
    uint128 public gasLimit = 1_000_000;

    /**
     * @notice Default value to be sent with cross-chain messages
     */
    uint128 public callValue = 1e15;

    /**
     * @notice Maximum retry attempts for failed messages
     */
    uint8 public maxRetryAttempts = 3;

    /**
     * @notice Retry delay in seconds
     */
    uint256 public retryDelay = 300; // 5 minutes

    /**
     * @notice Low balance threshold for Ethereum mainnet (0.01 ETH)
     */
    uint256 public lowBalanceThreshold = 1e16;

    /**
     * @notice Array to store failed messages for retry
     */
    FailedMessage[] public failedMessages;

    /**
     * @notice Mapping to track whitelisted operators
     */
    mapping(address => bool) public isWhitelisted;

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
    event MessageFailed(
        uint32 indexed dstEid,
        bytes32 indexed guid,
        bytes reason
    );

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
     * @notice Emitted when contract balance is low
     * @param currentBalance The current contract balance
     * @param threshold The low balance threshold
     */
    event LowBalanceAlert(uint256 currentBalance, uint256 threshold);

    /**
     * @notice Emitted when a message fails and is queued for retry
     * @param action The type of action that failed
     * @param operator The address of the operator
     * @param retryCount The current retry count
     */
    event MessageFailedRetryQueued(ActionType action, address operator, uint8 retryCount);

    /**
     * @notice Emitted when a retry attempt succeeds
     * @param action The type of action that was retried
     * @param operator The address of the operator
     * @param retryCount The retry count when it succeeded
     */
    event RetrySucceeded(ActionType action, address operator, uint8 retryCount);

    /**
     * @notice Emitted when maximum retry attempts are reached
     * @param action The type of action that permanently failed
     * @param operator The address of the operator
     */
    event MaxRetriesReached(ActionType action, address operator);

    /**
     * @notice Constructor for the AvsGovernanceLogic contract
     * @param _endpoint The LayerZero endpoint address
     * @param _proxyHub The address of the L2 proxy hub contract
     * @param _dstEid The destination chain endpoint ID
     * @param _ownerAddress The owner address
     */
    constructor(
        address _endpoint,
        address _proxyHub,
        uint32 _dstEid,
        address _ownerAddress
    ) OApp(_endpoint, _ownerAddress) Ownable(_ownerAddress) {
        require(_proxyHub != address(0), "Invalid proxyHub");
        proxyHub = _proxyHub;
        dstEid = _dstEid;

        _setPeer(dstEid, bytes32(uint256(uint160(_proxyHub))));
    }

    /**
     * @notice Updates the proxy hub address
     * @param _proxyHub The new proxy hub address
     */
    function setProxyHub(address _proxyHub) external onlyOwner {
        require(_proxyHub != address(0), "Invalid proxy hub address");
        proxyHub = _proxyHub;
        _setPeer(dstEid, bytes32(uint256(uint160(_proxyHub))));
    }

    /**
     * @notice Updates the retry configuration
     * @param _maxRetryAttempts The maximum number of retry attempts
     * @param _retryDelay The delay between retry attempts in seconds
     */
    function setRetryConfig(uint8 _maxRetryAttempts, uint256 _retryDelay) external onlyOwner {
        maxRetryAttempts = _maxRetryAttempts;
        retryDelay = _retryDelay;
    }

    /**
     * @notice Updates the low balance threshold
     * @param _threshold The new low balance threshold
     */
    function setLowBalanceThreshold(uint256 _threshold) external onlyOwner {
        lowBalanceThreshold = _threshold;
    }

    /**
     * @notice Estimates the fee required for sending a message
     * @param action The type of action to send
     * @param operator The address of the operator
     * @return fee The estimated messaging fee
     */
    function estimateFee(ActionType action, address operator) 
        external 
        view 
        returns (MessagingFee memory fee) 
    {
        bytes memory payload = abi.encode(action, operator);
        bytes memory options = _buildExecutorOptions(gasLimit, callValue);
        
        return endpoint.quote(
            MessagingParams({
                dstEid: dstEid,
                receiver: bytes32(uint256(uint160(proxyHub))),
                message: payload,
                options: options,
                payInLzToken: false
            }),
            address(this)
        );
    }

    /**
     * @notice Allows the owner to withdraw ETH from the contract
     * @param _to The address to send the ETH to
     * @param _amount The amount of ETH to withdraw
     */
    function withdraw(address payable _to, uint256 _amount) external onlyOwner {
        require(_to != address(0), "Invalid recipient");
        require(_amount > 0, "Amount must be greater than 0");
        require(address(this).balance >= _amount, "Insufficient balance");
        _to.transfer(_amount);
    }

    /**
     * @notice Hook called before an operator is registered
     * @param _operator The address of the operator to be registered
     */
    function beforeOperatorRegistered(
        address _operator,
        uint256 /* _votingPower */,
        uint256[4] calldata /* _blsKey */,
        address /* _rewardsReceiver */
    ) external override {
        require(isWhitelisted[_operator], "Operator is not whitelisted");
    }

    /**
     * @notice Hook called after an operator is registered
     * @param _operator The address of the registered operator
     */
    function afterOperatorRegistered(
        address _operator,
        uint256 /* _votingPower */,
        uint256[4] calldata /* _blsKey */,
        address /* _rewardsReceiver */
    ) external override {
        _checkLowBalance();
        _sendMessage(ActionType.REGISTER, _operator);
    }

    /**
     * @notice Hook called before an operator is unregistered
     */
    function beforeOperatorUnregistered(address) external override {}

    /**
     * @notice Hook called after an operator is unregistered
     * @param _operator The address of the unregistered operator
     */
    function afterOperatorUnregistered(address _operator) external override {
        _checkLowBalance();
        _sendMessage(ActionType.UNREGISTER, _operator);
    }

    /**
     * @notice Internal function to send a message
     * @param action The type of action to send
     * @param operator The address of the operator
     */
    function _sendMessage(ActionType action, address operator) internal {
        bytes memory payload = abi.encode(action, operator);
        bytes memory options = _buildExecutorOptions(gasLimit, callValue);

        try
            endpoint.quote(
                MessagingParams({
                    dstEid: dstEid,
                    receiver: bytes32(uint256(uint160(proxyHub))),
                    message: payload,
                    options: options,
                    payInLzToken: false
                }),
                address(this)
            )
        returns (MessagingFee memory fee) {
            if (address(this).balance < fee.nativeFee) {
                _queueForRetry(action, operator, 0);
                return;
            }

            try this._safeLzSend(action, operator, payload, options, fee) {
                if (action == ActionType.REGISTER) {
                    emit OperatorRegistered(operator);
                } else {
                    emit OperatorUnregistered(operator);
                }
            } catch (bytes memory reason) {
                emit MessageFailed(dstEid, bytes32(0), reason);
                _queueForRetry(action, operator, 0);
            }
        } catch (bytes memory reason) {
            emit MessageFailed(dstEid, bytes32(0), reason);
            _queueForRetry(action, operator, 0);
        }
    }

    /**
     * @notice Safe wrapper for _lzSend that can be used with try/catch
     * @param action The type of action being sent
     * @param operator The address of the operator
     * @param payload The message payload
     * @param options The message options
     * @param fee The messaging fee
     */
    function _safeLzSend(
        ActionType action,
        address operator,
        bytes memory payload,
        bytes memory options,
        MessagingFee memory fee
    ) external {
        require(msg.sender == address(this), "Only self");
        
        MessagingReceipt memory receipt = _lzSend(
            dstEid,
            payload,
            options,
            fee,
            payable(address(this))
        );

        emit MessageSent(dstEid, receipt.guid, fee.nativeFee);
    }

    /**
     * @notice Queues a failed message for retry
     * @param action The type of action that failed
     * @param operator The address of the operator
     * @param retryCount The current retry count
     */
    function _queueForRetry(ActionType action, address operator, uint8 retryCount) internal {
        if (retryCount >= maxRetryAttempts) {
            emit MaxRetriesReached(action, operator);
            return;
        }

        failedMessages.push(FailedMessage({
            action: action,
            operator: operator,
            timestamp: block.timestamp,
            retryCount: retryCount
        }));

        emit MessageFailedRetryQueued(action, operator, retryCount);
    }

    /**
     * @notice Retries failed messages
     * @param maxRetries The maximum number of failed messages to retry in this call
     */
    function retryFailedMessages(uint256 maxRetries) external {
        uint256 processed = 0;
        uint256 i = 0;
        
        while (i < failedMessages.length && processed < maxRetries) {
            FailedMessage memory failed = failedMessages[i];
            
            if (block.timestamp >= failed.timestamp + retryDelay) {
                bytes memory payload = abi.encode(failed.action, failed.operator);
                bytes memory options = _buildExecutorOptions(gasLimit, callValue);
                
                try
                    endpoint.quote(
                        MessagingParams({
                            dstEid: dstEid,
                            receiver: bytes32(uint256(uint160(proxyHub))),
                            message: payload,
                            options: options,
                            payInLzToken: false
                        }),
                        address(this)
                    )
                returns (MessagingFee memory fee) {
                    if (address(this).balance >= fee.nativeFee) {
                        try this._safeLzSend(failed.action, failed.operator, payload, options, fee) {
                            emit RetrySucceeded(failed.action, failed.operator, failed.retryCount);
                            
                            if (failed.action == ActionType.REGISTER) {
                                emit OperatorRegistered(failed.operator);
                            } else {
                                emit OperatorUnregistered(failed.operator);
                            }
                            
                            // Remove the failed message from the array
                            failedMessages[i] = failedMessages[failedMessages.length - 1];
                            failedMessages.pop();
                            continue; // Don't increment i since we moved an element down
                        } catch {
                            _queueForRetry(failed.action, failed.operator, failed.retryCount + 1);
                            
                            // Remove the old entry
                            failedMessages[i] = failedMessages[failedMessages.length - 1];
                            failedMessages.pop();
                            continue; // Don't increment i since we moved an element down
                        }
                    } else {
                        _queueForRetry(failed.action, failed.operator, failed.retryCount + 1);
                        
                        // Remove the old entry
                        failedMessages[i] = failedMessages[failedMessages.length - 1];
                        failedMessages.pop();
                        continue; // Don't increment i since we moved an element down
                    }
                } catch {
                    _queueForRetry(failed.action, failed.operator, failed.retryCount + 1);
                    
                    // Remove the old entry
                    failedMessages[i] = failedMessages[failedMessages.length - 1];
                    failedMessages.pop();
                    continue; // Don't increment i since we moved an element down
                }
                
                processed++;
            }
            i++;
        }
    }

    /**
     * @notice Gets the count of failed messages waiting for retry
     * @return The number of failed messages in the queue
     */
    function getFailedMessageCount() external view returns (uint256) {
        return failedMessages.length;
    }

    /**
     * @notice Checks if contract balance is low and emits alert if needed
     */
    function _checkLowBalance() internal {
        if (address(this).balance < lowBalanceThreshold) {
            emit LowBalanceAlert(address(this).balance, lowBalanceThreshold);
        }
    }

    /**
     * @notice Manually trigger low balance check
     */
    function checkLowBalance() external {
        _checkLowBalance();
    }

    /**
     * @notice Handles incoming LayerZero messages
     * @dev This contract should not receive messages
     */
    function _lzReceive(
        Origin calldata,
        bytes32,
        bytes calldata,
        address,
        bytes calldata
    ) internal override {
        revert("AvsGovernanceLogic: should not receive messages");
    }

    /**
     * @notice Builds executor options for LayerZero messages
     * @param gas The gas limit for the message
     * @param value The value to be sent with the message
     * @return The encoded options
     */
    function _buildExecutorOptions(
        uint128 gas,
        uint128 value
    ) internal pure returns (bytes memory) {
        uint16 TYPE_3 = 3;
        uint8 WORKER_ID = 1;
        uint8 OPTION_TYPE_LZRECEIVE = 1;

        bytes memory option = value == 0
            ? abi.encodePacked(gas)
            : abi.encodePacked(gas, value);

        uint16 optionLength = uint16(option.length + 1);

        return
            abi.encodePacked(
                TYPE_3,
                WORKER_ID,
                optionLength,
                OPTION_TYPE_LZRECEIVE,
                option
            );
    }

    /**
     * @notice Updates the gas configuration for cross-chain messages
     * @param _gasLimit The new gas limit
     * @param _callValue The new call value
     */
    function setGasOptions(
        uint128 _gasLimit,
        uint128 _callValue
    ) external onlyOwner {
        gasLimit = _gasLimit;
        callValue = _callValue;
        emit GasOptionsUpdated(_gasLimit, _callValue);
    }

    /**
     * @notice Override _payNative to use contract balance instead of msg.value
     * @param _nativeFee The native fee to be paid
     * @return nativeFee The amount of native currency paid
     */
    function _payNative(
        uint256 _nativeFee
    ) internal override returns (uint256 nativeFee) {
        require(
            address(this).balance >= _nativeFee,
            "Insufficient contract balance"
        );
        return _nativeFee;
    }

    /**
     * @notice Adds multiple operators to the whitelist
     * @param _operators Array of operator addresses to whitelist
     */
    function addToWhitelist(address[] calldata _operators) external onlyOwner {
        for (uint256 i = 0; i < _operators.length; i++) {
            _addToWhitelist(_operators[i]);
        }
    }

    /**
     * @notice Internal function to add an operator to the whitelist
     * @param _operator The address of the operator to whitelist
     */
    function _addToWhitelist(address _operator) internal {
        require(_operator != address(0), "Invalid address");
        require(!isWhitelisted[_operator], "Already whitelisted");
        isWhitelisted[_operator] = true;
        emit Whitelisted(_operator);
    }

    /**
     * @notice Removes multiple operators from the whitelist
     * @param _operators Array of operator addresses to remove from whitelist
     */
    function removeFromWhitelist(
        address[] calldata _operators
    ) external onlyOwner {
        for (uint256 i = 0; i < _operators.length; i++) {
            _removeFromWhitelist(_operators[i]);
        }
    }

    /**
     * @notice Internal function to remove an operator from the whitelist
     * @param _operator The address of the operator to remove from whitelist
     */
    function _removeFromWhitelist(address _operator) internal {
        require(isWhitelisted[_operator], "Not whitelisted");
        isWhitelisted[_operator] = false;
        emit UnWhitelisted(_operator);
    }

    /**
     * @notice Allows the contract to receive ETH
     */
    receive() external payable {}
}
