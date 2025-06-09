// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import { OApp, MessagingFee, Origin } from "@layerzero-v2/oapp/contracts/oapp/OApp.sol";
import { OAppOptionsType3 } from "@layerzero-v2/oapp/contracts/oapp/libs/OAppOptionsType3.sol";
import { Ownable } from "@openzeppelin-contracts/contracts/access/Ownable.sol";
import { ReentrancyGuard } from "solady/utils/ReentrancyGuard.sol";

/**
 * @title ProxyHub
 * @notice A LayerZero-enabled contract that manages keeper registration and function execution across multiple chains
 * @dev This contract acts as a hub for managing keepers and broadcasting their status across different chains
 * @custom:security-contact security@triggerx.com
 */
contract ProxyHub is Ownable, OApp, OAppOptionsType3, ReentrancyGuard {
    /**
     * @notice Enum defining the types of actions that can be performed on keepers
     */
    enum ActionType { REGISTER, UNREGISTER }

    /**
     * @notice Structure to store failed broadcast attempts for retry
     */
    struct FailedBroadcast {
        ActionType action;
        address keeper;
        uint32 dstEid;
        uint256 timestamp;
        uint8 retryCount;
    }

    /**
     * @notice Mapping to track registered keepers
     */
    mapping(address => bool) public isKeeper;

    /**
     * @notice Array to store failed broadcasts for retry
     */
    FailedBroadcast[] public failedBroadcasts;

    /**
     * @notice The LayerZero endpoint ID of this chain
     */
    uint32 public immutable thisChainEid;

    /**
     * @notice The LayerZero endpoint ID of the source chain
     */
    uint32 public immutable srcEid;

    /**
     * @notice Array of destination chain endpoint IDs
     */
    uint32[] public dstEids;

    /**
     * @notice Default gas limit for cross-chain messages
     */
    uint128 public defaultGas = 1_000_000;

    /**
     * @notice Default value to be sent with cross-chain messages
     */
    uint128 public defaultValue = 0;

    /**
     * @notice Maximum retry attempts for failed broadcasts
     */
    uint8 public maxRetryAttempts = 3;

    /**
     * @notice Retry delay in seconds
     */
    uint256 public retryDelay = 300; // 5 minutes

    /**
     * @notice Low balance threshold for Base mainnet (0.001 ETH)
     */
    uint256 public lowBalanceThreshold = 1e15;

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
     * @notice Emitted when a broadcast fails and is queued for retry
     * @param action The type of action that failed
     * @param keeper The address of the keeper
     * @param dstEid The destination chain endpoint ID
     * @param retryCount The current retry count
     */
    event BroadcastFailedRetryQueued(ActionType action, address keeper, uint32 dstEid, uint8 retryCount);

    /**
     * @notice Emitted when a retry attempt succeeds
     * @param action The type of action that was retried
     * @param keeper The address of the keeper
     * @param dstEid The destination chain endpoint ID
     * @param retryCount The retry count when it succeeded
     */
    event RetrySucceeded(ActionType action, address keeper, uint32 dstEid, uint8 retryCount);

    /**
     * @notice Emitted when maximum retry attempts are reached
     * @param action The type of action that permanently failed
     * @param keeper The address of the keeper
     * @param dstEid The destination chain endpoint ID
     */
    event MaxRetriesReached(ActionType action, address keeper, uint32 dstEid);

    /**
     * @notice Modifier to restrict function access to registered keepers
     */
    modifier onlyKeeper() {
        require(isKeeper[msg.sender], "Not a keeper");
        _;
    }

    /**
     * @notice Constructor for the ProxyHub contract
     * @param _endpoint The LayerZero endpoint address
     * @param _ownerAddress The owner address
     * @param _srcEid The source chain endpoint ID
     * @param _thisChainEid The current chain endpoint ID
     * @param _initialKeepers Array of initial keeper addresses
     */
    constructor(
        address _endpoint,
        address _ownerAddress,
        uint32 _srcEid,
        uint32 _thisChainEid,
        address[] memory _initialKeepers
    )
        OApp(_endpoint, _ownerAddress)
        Ownable(_ownerAddress)
    {
        srcEid = _srcEid;
        thisChainEid = _thisChainEid;

        for (uint i = 0; i < _initialKeepers.length; i++) {
            isKeeper[_initialKeepers[i]] = true;
            emit KeeperRegistered(_initialKeepers[i]);
        }
    }
    
    /**
     * @notice Adds new spoke chains to the network
     * @param _dstEids Array of destination chain endpoint IDs
     */
    function addSpokes(uint32[] calldata _dstEids) external onlyOwner {
        for (uint i = 0; i < _dstEids.length; i++) {
            uint32 dstEid = _dstEids[i];
            if (dstEid != thisChainEid) {
                dstEids.push(dstEid);
                _setPeer(dstEid, bytes32(uint256(uint160(address(this)))));
            }
        }
    }

    /**
     * @notice Updates the gas configuration for cross-chain messages
     * @param gas The new gas limit
     * @param value The new default value
     */
    function setGasConfig(uint128 gas, uint128 value) external onlyOwner {
        defaultGas = gas;
        defaultValue = value;
        emit GasConfigUpdated(gas, value);
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
     * @notice Estimates the total fees required for broadcasting to all destination chains
     * @param action The type of action to broadcast
     * @param keeper The address of the keeper
     * @return totalFee The total estimated fee for all destinations
     * @return fees Array of individual fees for each destination
     */
    function estimateTotalFees(ActionType action, address keeper) 
        external 
        view 
        returns (uint256 totalFee, MessagingFee[] memory fees) 
    {
        bytes memory payload = abi.encode(action, keeper);
        bytes memory options = _buildExecutorOptions(defaultGas, defaultValue);
        
        uint256 validDestinations = 0;
        for (uint i = 0; i < dstEids.length; i++) {
            if (dstEids[i] != thisChainEid) {
                validDestinations++;
            }
        }
        
        fees = new MessagingFee[](validDestinations);
        uint256 feeIndex = 0;
        
        for (uint i = 0; i < dstEids.length; i++) {
            uint32 dstEid = dstEids[i];
            if (dstEid == thisChainEid) continue;
            
            MessagingFee memory fee = _quote(dstEid, payload, options, false);
            fees[feeIndex] = fee;
            totalFee += fee.nativeFee;
            feeIndex++;
        }
    }

    /**
     * @notice Executes a function on a target contract
     * @param target The address of the target contract
     * @param data The calldata for the function call
     */
    function executeFunction(address target, bytes calldata data) external payable onlyKeeper nonReentrant {
        _executeFunction(target, data);
    }

    /**
     * @notice Internal function to execute a function call
     * @param target The address of the target contract
     * @param callData The calldata for the function call
     * @return The return data from the function call
     */
    function _executeFunction(address target, bytes memory callData) internal returns (bytes memory) {
        (bool success, bytes memory result) = target.call{value: msg.value}(callData);
        require(success, "Execution failed");

        emit FunctionExecuted(msg.sender, target, callData, msg.value);
        return result;
    }

    /**
     * @notice Handles incoming LayerZero messages
     * @param _origin The origin information of the message
     * @param _payload The message payload
     */
    function _lzReceive(
        Origin calldata _origin,
        bytes32, /* _guid */
        bytes calldata _payload,
        address, /* _executor */
        bytes calldata /* extraData */
    ) internal override nonReentrant {
        require(_origin.srcEid == srcEid, "Invalid source chain");
        (ActionType action, address keeper) = abi.decode(_payload, (ActionType, address));

        if (action == ActionType.REGISTER) {
            isKeeper[keeper] = true;
            emit KeeperRegistered(keeper);
            _batchBroadcast(ActionType.REGISTER, keeper);
        } else if (action == ActionType.UNREGISTER) {
            isKeeper[keeper] = false;
            emit KeeperUnregistered(keeper);
            _batchBroadcast(ActionType.UNREGISTER, keeper);
        } else {
            revert("Unknown action");
        }
    }

    /**
     * @notice Broadcasts a keeper action to all spoke chains
     * @param action The type of action to broadcast
     * @param keeper The address of the keeper
     */
    function _batchBroadcast(ActionType action, address keeper) internal {
        _checkLowBalance();
        
        bytes memory payload = abi.encode(action, keeper);
        uint256 totalUsed = 0;
        uint256 initialValue = msg.value;

        for (uint i = 0; i < dstEids.length; i++) {
            uint32 dstEid = dstEids[i];
            if (dstEid == thisChainEid) continue;

            bytes memory options = _buildExecutorOptions(defaultGas, defaultValue);
            MessagingFee memory fee = _quote(dstEid, payload, options, false);

            if (initialValue < totalUsed + fee.nativeFee) {
                _queueForRetry(action, keeper, dstEid, 0);
                continue;
            }

            try this._safeLzSend(dstEid, payload, options, fee) {
                emit BroadcastSent(action, keeper, dstEid);
                emit FeeUsed(dstEid, fee.nativeFee);
                totalUsed += fee.nativeFee;
            } catch {
                _queueForRetry(action, keeper, dstEid, 0);
            }
        }
    }

    /**
     * @notice Safe wrapper for _lzSend that can be used with try/catch
     * @param dstEid The destination chain endpoint ID
     * @param payload The message payload
     * @param options The message options
     * @param fee The messaging fee
     */
    function _safeLzSend(
        uint32 dstEid,
        bytes memory payload,
        bytes memory options,
        MessagingFee memory fee
    ) external {
        require(msg.sender == address(this), "Only self");
        
        _lzSend(
            dstEid,
            payload,
            options,
            fee,
            payable(address(this))
        );
    }

    /**
     * @notice Queues a failed broadcast for retry
     * @param action The type of action that failed
     * @param keeper The address of the keeper
     * @param dstEid The destination chain endpoint ID
     * @param retryCount The current retry count
     */
    function _queueForRetry(ActionType action, address keeper, uint32 dstEid, uint8 retryCount) internal {
        if (retryCount >= maxRetryAttempts) {
            emit MaxRetriesReached(action, keeper, dstEid);
            return;
        }

        failedBroadcasts.push(FailedBroadcast({
            action: action,
            keeper: keeper,
            dstEid: dstEid,
            timestamp: block.timestamp,
            retryCount: retryCount
        }));

        emit BroadcastFailedRetryQueued(action, keeper, dstEid, retryCount);
    }

    /**
     * @notice Retries failed broadcasts
     * @param maxRetries The maximum number of failed broadcasts to retry in this call
     */
    function retryFailedBroadcasts(uint256 maxRetries) external {
        uint256 processed = 0;
        uint256 i = 0;
        
        while (i < failedBroadcasts.length && processed < maxRetries) {
            FailedBroadcast memory failed = failedBroadcasts[i];
            
            if (block.timestamp >= failed.timestamp + retryDelay) {
                bytes memory payload = abi.encode(failed.action, failed.keeper);
                bytes memory options = _buildExecutorOptions(defaultGas, defaultValue);
                MessagingFee memory fee = _quote(failed.dstEid, payload, options, false);
                
                if (address(this).balance >= fee.nativeFee) {
                    try this._safeLzSend(failed.dstEid, payload, options, fee) {
                        emit RetrySucceeded(failed.action, failed.keeper, failed.dstEid, failed.retryCount);
                        emit BroadcastSent(failed.action, failed.keeper, failed.dstEid);
                        emit FeeUsed(failed.dstEid, fee.nativeFee);
                        
                        // Remove the failed broadcast from the array
                        failedBroadcasts[i] = failedBroadcasts[failedBroadcasts.length - 1];
                        failedBroadcasts.pop();
                        continue; // Don't increment i since we moved an element down
                    } catch {
                        _queueForRetry(failed.action, failed.keeper, failed.dstEid, failed.retryCount + 1);
                        
                        // Remove the old entry
                        failedBroadcasts[i] = failedBroadcasts[failedBroadcasts.length - 1];
                        failedBroadcasts.pop();
                        continue; // Don't increment i since we moved an element down
                    }
                } else {
                    _queueForRetry(failed.action, failed.keeper, failed.dstEid, failed.retryCount + 1);
                    
                    // Remove the old entry
                    failedBroadcasts[i] = failedBroadcasts[failedBroadcasts.length - 1];
                    failedBroadcasts.pop();
                    continue; // Don't increment i since we moved an element down
                }
                
                processed++;
            }
            i++;
        }
    }

    /**
     * @notice Gets the count of failed broadcasts waiting for retry
     * @return The number of failed broadcasts in the queue
     */
    function getFailedBroadcastCount() external view returns (uint256) {
        return failedBroadcasts.length;
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
     * @notice Builds executor options for LayerZero messages
     * @param gas The gas limit for the message
     * @param value The value to be sent with the message
     * @return The encoded options
     */
    function _buildExecutorOptions(uint128 gas, uint128 value) internal pure returns (bytes memory) {
        uint16 TYPE_3 = 3;
        uint8 WORKER_ID = 1;
        uint8 OPTION_TYPE_LZRECEIVE = 1;

        bytes memory option = value == 0
            ? abi.encodePacked(gas)
            : abi.encodePacked(gas, value);

        uint16 optionLength = uint16(option.length + 1);

        return abi.encodePacked(
            TYPE_3,
            WORKER_ID,
            optionLength,
            OPTION_TYPE_LZRECEIVE,
            option
        );
    }

    /**
     * @notice Override _payNative to use contract balance instead of msg.value
     * @param _nativeFee The native fee to be paid
     * @return nativeFee The amount of native currency paid
     */
    function _payNative(uint256 _nativeFee) internal override returns (uint256 nativeFee) {
        require(address(this).balance >= _nativeFee, "Insufficient contract balance");
        return _nativeFee;
    }

    /**
     * @notice Allows the contract to receive ETH
     */
    receive() external payable {}

    /**
     * @notice Allows the owner to withdraw ETH from the contract
     * @param to The address to send the ETH to
     * @param amount The amount of ETH to withdraw
     */
    function withdraw(address payable to, uint256 amount) external onlyOwner nonReentrant {
        require(to != address(0), "Invalid recipient");
        require(amount <= address(this).balance, "Insufficient balance");
        to.transfer(amount);
    }
}
