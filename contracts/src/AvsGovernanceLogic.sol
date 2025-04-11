// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import { IAvsGovernanceLogic } from "@othentic-core-contracts/NetworkManagement/L1/interfaces/IAvsGovernanceLogic.sol";
import { OApp, MessagingFee, Origin } from "@layerzero-v2/oapp/contracts/oapp/OApp.sol";
import { Ownable } from "@openzeppelin-contracts/contracts/access/Ownable.sol";

contract AvsGovernanceLogic is Ownable, IAvsGovernanceLogic, OApp {
    enum ActionType { REGISTER, UNREGISTER }

    address public immutable proxyHub; // L2 handler
    uint32 public immutable dstEid;    // L2 chain ID

    constructor(address _endpoint, address _proxyHub, uint32 _dstEid, address _owner)
        OApp(_endpoint, _owner)
        Ownable(_owner)
    {
        require(_proxyHub != address(0), "Invalid proxyHub");
        proxyHub = _proxyHub;
        dstEid = _dstEid;

        _setPeer(dstEid, bytes32(uint256(uint160(_proxyHub))));
    }

    function beforeOperatorRegistered(
        address,
        uint256,
        uint256[4] calldata,
        address
    ) external override {}

    function afterOperatorRegistered(
        address _operator,
        uint256,
        uint256[4] calldata,
        address
    ) external override {
        bytes memory payload = abi.encode(ActionType.REGISTER, _operator);

        _lzSend(
            dstEid,
            payload,
            bytes(""),                 // no custom options
            MessagingFee(0, 0),        // pre-funded by relayer or AVS
            payable(msg.sender)        // refund address
        );
    }

    function beforeOperatorUnregistered(address) external override {}

    function afterOperatorUnregistered(address _operator) external override {
        bytes memory payload = abi.encode(ActionType.UNREGISTER, _operator);

        _lzSend(
            dstEid,
            payload,
            bytes(""),
            MessagingFee(0, 0),
            payable(msg.sender)
        );
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
}
