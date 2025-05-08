// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import { IAvsGovernanceLogic } from "@othentic-core-contracts/NetworkManagement/L1/interfaces/IAvsGovernanceLogic.sol";
import { OApp, MessagingFee, Origin, MessagingReceipt } from "@layerzero-v2/oapp/contracts/oapp/OApp.sol";
import { ILayerZeroEndpointV2, MessagingParams } from "@layerzerolabs/lz-evm-protocol-v2/contracts/interfaces/ILayerZeroEndpointV2.sol";
import { Ownable } from "@openzeppelin-contracts/contracts/access/Ownable.sol";

contract AvsGovernanceLogic is Ownable, IAvsGovernanceLogic, OApp {

    enum ActionType { REGISTER, UNREGISTER }

    address public proxyHub; // L2 handler
    uint32 public immutable dstEid;    // L2 chain ID

    // Whitelist storage
    mapping(address => bool) public isWhitelisted;

    event OperatorRegistered(address indexed operator);
    event OperatorUnregistered(address indexed operator);
    event MessageSent(uint32 indexed dstEid, bytes32 indexed guid, uint256 fee);
    event MessageFailed(uint32 indexed dstEid, bytes32 indexed guid, bytes reason);
    event Whitelisted(address indexed operator);
    event UnWhitelisted(address indexed operator);

    constructor(address _endpoint, address _proxyHub, uint32 _dstEid, address _owner)
        OApp(_endpoint, _owner)
        Ownable(_owner)
    {
        require(_proxyHub != address(0), "Invalid proxyHub");
        proxyHub = _proxyHub;
        dstEid = _dstEid;

        _setPeer(dstEid, bytes32(uint256(uint160(_proxyHub))));
    }

    function setProxyHub(address _proxyHub) external onlyOwner {
        proxyHub = _proxyHub;
        _setPeer(dstEid, bytes32(uint256(uint160(_proxyHub))));
    }

    // Allow the contract to receive Ether for funding fees
    receive() external payable {}

    // Withdraw Ether from the contract
    function withdraw(address payable _to, uint256 _amount) external onlyOwner {
        require(_to != address(0), "Invalid recipient");
        require(_amount > 0, "Amount must be greater than 0");
        require(address(this).balance >= _amount, "Insufficient balance");
        _to.transfer(_amount);
    }

    function beforeOperatorRegistered(
        address _operator,
        uint256,
        uint256[4] calldata,
        address
    ) external override {
        require(isWhitelisted[_operator], "Operator is not whitelisted");
    }

    function afterOperatorRegistered(
        address _operator,
        uint256,
        uint256[4] calldata,
        address
    ) external override {
        bytes memory payload = abi.encode(ActionType.REGISTER, _operator);
        bytes memory options = _buildExecutorOptions(1_000_000, 1e15);
        
        try endpoint.quote(
            MessagingParams({
                dstEid: dstEid,
                receiver: bytes32(uint256(uint160(proxyHub))),
                message: payload,
                options: options,
                payInLzToken: false
            }),
            address(this)
        ) returns (MessagingFee memory fee) {
            require(address(this).balance >= fee.nativeFee, "Insufficient balance for message fee");
            
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
            revert("Failed to send register message");
        }
    }

    function beforeOperatorUnregistered(address) external override {}

    function afterOperatorUnregistered(address _operator) external override {
        bytes memory payload = abi.encode(ActionType.UNREGISTER, _operator);
        bytes memory options = _buildExecutorOptions(1_000_000, 1e15);
        
        try endpoint.quote(
            MessagingParams({
                dstEid: dstEid,
                receiver: bytes32(uint256(uint160(proxyHub))),
                message: payload,
                options: options,
                payInLzToken: false
            }),
            address(this)
        ) returns (MessagingFee memory fee) {
            require(address(this).balance >= fee.nativeFee, "Insufficient balance for message fee");
            
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
            revert("Failed to send unregister message");
        }
    }

    function _lzReceive(
        Origin calldata,
        bytes32,
        bytes calldata,
        address,
        bytes calldata
    ) internal override {
        revert("AvsGovernanceLogic: should not receive messages");
    }

    function _buildExecutorOptions(uint128 gas, uint128 value) internal pure returns (bytes memory) {
        uint16 TYPE_3 = 3;
        uint8 WORKER_ID = 1;
        uint8 OPTION_TYPE_LZRECEIVE = 1;

        // Encode the option payload: gas + value (16 bytes if both)
        bytes memory option = value == 0
            ? abi.encodePacked(gas)
            : abi.encodePacked(gas, value);

        uint16 optionLength = uint16(option.length + 1); // +1 for optionType

        return abi.encodePacked(
            TYPE_3,                   // option type 3
            WORKER_ID,               // executor ID
            optionLength,            // size of option payload
            OPTION_TYPE_LZRECEIVE,   // receive option type
            option                   // the actual payload: gas + optional value
        );
    }

    /**
     * @dev Override _payNative to use contract balance instead of msg.value
     * @param _nativeFee The native fee to be paid
     * @return nativeFee The amount of native currency paid
     */
    function _payNative(uint256 _nativeFee) internal override returns (uint256 nativeFee) {
        require(address(this).balance >= _nativeFee, "Insufficient contract balance");
        return _nativeFee;
    }

    // Add operator to whitelist (onlyOwner)
    function addToWhitelist(address _operator) external onlyOwner {
        require(_operator != address(0), "Invalid address");
        require(!isWhitelisted[_operator], "Already whitelisted");
        isWhitelisted[_operator] = true;
        emit Whitelisted(_operator);
    }

    // Remove operator from whitelist (onlyOwner)
    function removeFromWhitelist(address _operator) external onlyOwner {
        require(isWhitelisted[_operator], "Not whitelisted");
        isWhitelisted[_operator] = false;
        emit UnWhitelisted(_operator);
    }
}