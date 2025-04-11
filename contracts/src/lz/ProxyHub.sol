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
        uint32 _thisChainEid    // This L2
    )
        OApp(_endpoint, _owner)
        Ownable(_owner)
    {
        srcEid = _srcEid;
        thisChainEid = _thisChainEid;
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

            bytes memory options = abi.encodePacked(uint16(0x0001), uint16(uint256(1000000)));
            MessagingFee memory fee = _quote(dstEid, payload, options, false);
            require(msg.value >= totalUsed + fee.nativeFee, "Insufficient fee");

            _lzSend(
                dstEid,
                payload,
                options,
                fee,
                payable(msg.sender)
            );

            emit BroadcastSent(action, keeper, dstEid);
            totalUsed += fee.nativeFee;
        }

    }

    receive() external payable {}

    function withdraw(address payable to, uint256 amount) external onlyOwner {
        require(to != address(0), "Invalid recipient");
        require(amount <= address(this).balance, "Insufficient balance");
        to.transfer(amount);
    }
}
