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
        bytes memory payload = abi.encode(ActionType.REGISTER, _operator);
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
            require(
                address(this).balance >= fee.nativeFee,
                "Insufficient balance for message fee"
            );

            MessagingReceipt memory receipt = _lzSend(
                dstEid,
                payload,
                options,
                fee,
                payable(address(this))
            );

            emit MessageSent(dstEid, receipt.guid, fee.nativeFee);
            emit OperatorRegistered(_operator);
        } catch (bytes memory reason) {
            emit MessageFailed(dstEid, bytes32(0), reason);
            revert(
                string(abi.encodePacked("Register failed: ", string(reason)))
            );
        }
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
        bytes memory payload = abi.encode(ActionType.UNREGISTER, _operator);
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
            require(
                address(this).balance >= fee.nativeFee,
                "Insufficient balance for message fee"
            );

            MessagingReceipt memory receipt = _lzSend(
                dstEid,
                payload,
                options,
                fee,
                payable(address(this))
            );

            emit MessageSent(dstEid, receipt.guid, fee.nativeFee);
            emit OperatorUnregistered(_operator);
        } catch (bytes memory reason) {
            emit MessageFailed(dstEid, bytes32(0), reason);
            revert(
                string(abi.encodePacked("Unregister failed: ", string(reason)))
            );
        }
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
