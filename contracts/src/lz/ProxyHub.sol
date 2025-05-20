// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

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
     * @notice Mapping to track registered keepers
     */
    mapping(address => bool) public isKeeper;

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
     * @notice Modifier to restrict function access to registered keepers
     */
    modifier onlyKeeper() {
        require(isKeeper[msg.sender], "Not a keeper");
        _;
    }

    /**
     * @notice Constructor for the ProxyHub contract
     * @param _endpoint The LayerZero endpoint address
     * @param _owner The owner address
     * @param _srcEid The source chain endpoint ID
     * @param _thisChainEid The current chain endpoint ID
     * @param _initialKeepers Array of initial keeper addresses
     */
    constructor(
        address _endpoint,
        address _owner,
        uint32 _srcEid,
        uint32 _thisChainEid,
        address[] memory _initialKeepers
    )
        OApp(_endpoint, _owner)
        Ownable(_owner)
    {
        srcEid = _srcEid;
        thisChainEid = _thisChainEid;

        for (uint i = 0; i < _initialKeepers.length; i++) {
            isKeeper[_initialKeepers[i]] = true;
            emit KeeperRegistered(_initialKeepers[i]);
        }
    }

    /**
     * @notice Sets the peer address for a given source chain
     * @param _srcEid The source chain endpoint ID
     * @param _avsGovernanceLogic The address of the AVS governance logic contract
     */
    function setPeer(uint32 _srcEid, address _avsGovernanceLogic) external onlyOwner {
        _setPeer(_srcEid, bytes32(uint256(uint160(_avsGovernanceLogic))));
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
        bytes memory payload = abi.encode(action, keeper);
        uint256 totalUsed;

        for (uint i = 0; i < dstEids.length; i++) {
            uint32 dstEid = dstEids[i];
            if (dstEid == thisChainEid) continue;

            bytes memory options = _buildExecutorOptions(defaultGas, defaultValue);
            MessagingFee memory fee = _quote(dstEid, payload, options, false);

            if (msg.value < totalUsed + fee.nativeFee) {
                continue;
            }

            _lzSend(
                dstEid,
                payload,
                options,
                fee,
                payable(address(this))
            );

            emit BroadcastSent(action, keeper, dstEid);
            emit FeeUsed(dstEid, fee.nativeFee);

            totalUsed += fee.nativeFee;
        }
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
