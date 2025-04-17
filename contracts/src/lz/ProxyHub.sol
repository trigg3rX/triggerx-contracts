// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import { OApp, MessagingFee, Origin } from "@layerzero-v2/oapp/contracts/oapp/OApp.sol";
import { OAppOptionsType3 } from "@layerzero-v2/oapp/contracts/oapp/libs/OAppOptionsType3.sol";
import { Ownable } from "@openzeppelin-contracts/contracts/access/Ownable.sol";

contract ProxyHub is Ownable, OApp, OAppOptionsType3 {
    enum ActionType { REGISTER, UNREGISTER }

    mapping(address => bool) public isKeeper;
    address public avsGovernance;

    uint32 public immutable thisChainEid;
    uint32 public srcEid;

    uint32[] public dstEids;

    event KeeperRegistered(address indexed keeper);
    event KeeperUnregistered(address indexed keeper);
    event BroadcastSent(ActionType action, address keeper, uint32 dstEid);
    event FunctionExecuted(address indexed keeper, address indexed target, bytes data, uint256 value);
    event AVSGovernanceUpdated(address indexed oldGovernance, address indexed newGovernance);

    modifier onlyKeeper() {
        require(isKeeper[msg.sender], "Not a keeper");
        _;
    }

    constructor(
        address _endpoint,
        address _owner,
        uint32 _srcEid,         // L1 source
        uint32 _thisChainEid,   // This L2
        address[] memory _initialKeepers
    )
        OApp(_endpoint, _owner)
        Ownable(_owner)
    {
        srcEid = _srcEid;
        thisChainEid = _thisChainEid;

        // Initialize keepers
        for (uint i = 0; i < _initialKeepers.length; i++) {
            isKeeper[_initialKeepers[i]] = true;
            emit KeeperRegistered(_initialKeepers[i]);
        }
    }

    function setPeer(uint32 _srcEid, address _avsGovernanceLogic) external onlyOwner {
        _setPeer(_srcEid, bytes32(uint256(uint160(_avsGovernanceLogic))));
    }

    function addSpokes(uint32[] calldata _dstEids) external onlyOwner {
        for (uint i = 0; i < _dstEids.length; i++) {
            uint32 dstEid = _dstEids[i];
            if (dstEid != thisChainEid) {
                dstEids.push(dstEid);
                _setPeer(dstEid, bytes32(uint256(uint160(address(this)))));
            }
        }
    }

    function executeFunction(address target, bytes calldata data) external payable onlyKeeper {
        _executeFunction(target, data);
    }

    function _executeFunction(address target, bytes memory callData) internal returns (bytes memory) {
        (bool success, bytes memory result) = target.call{value: msg.value}(callData);
        require(success, "Execution failed");

        emit FunctionExecuted(msg.sender, target, callData, msg.value);
        return result;
    }

    function _lzReceive(
        Origin calldata _origin,
        bytes32,
        bytes calldata _payload,
        address,
        bytes calldata
    ) internal override {
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

    function _batchBroadcast(ActionType action, address keeper) internal {
        bytes memory payload = abi.encode(action, keeper);
        uint256 totalUsed;

        for (uint i = 0; i < dstEids.length; i++) {
            uint32 dstEid = dstEids[i];
            if (dstEid == thisChainEid) continue;

            bytes memory options = _buildExecutorOptions(1_000_000, 0); // 1M gas, no value
            MessagingFee memory fee = _quote(dstEid, payload, options, false);
            require(msg.value >= totalUsed + fee.nativeFee, "Insufficient fee");

            _lzSend(
                dstEid,
                payload,
                options,
                fee,
                payable(address(this)) // Pay from contract balance
            );

            emit BroadcastSent(action, keeper, dstEid);
            totalUsed += fee.nativeFee;
        }
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

    receive() external payable {}

    function withdraw(address payable to, uint256 amount) external onlyOwner {
        require(to != address(0), "Invalid recipient");
        require(amount <= address(this).balance, "Insufficient balance");
        to.transfer(amount);
    }


}
