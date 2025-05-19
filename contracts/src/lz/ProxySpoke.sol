// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import { OApp, Origin } from "@layerzero-v2/oapp/contracts/oapp/OApp.sol";
import { Ownable } from "@openzeppelin-contracts/contracts/access/Ownable.sol";

contract ProxySpoke is Ownable, OApp {
    enum ActionType { REGISTER, UNREGISTER }

    mapping(address => bool) public isKeeper;

    event KeeperUpdated(ActionType action, address keeper);
    event FunctionExecuted(address indexed keeper, address indexed target, bytes data, uint256 value);

    modifier onlyKeeper() {
        require(isKeeper[msg.sender], "Spoke: Keeper not registered");
        _;
    }

    constructor(
        address _endpoint, 
        uint32 _hubEid,
        address[] memory _initialKeepers
    ) Ownable(msg.sender) OApp(_endpoint, msg.sender) {
        _setPeer(_hubEid, bytes32(uint256(uint160(address(this)))));

        // Initialize keepers
        for (uint i = 0; i < _initialKeepers.length; i++) {
            isKeeper[_initialKeepers[i]] = true;
            emit KeeperUpdated(ActionType.REGISTER, _initialKeepers[i]);
        }
    }

    /**
     * @notice Executes a function locally if the caller is a registered keeper
     * @param target Target contract address
     * @param data Calldata for the function call
     */
    function executeFunction(address target, bytes memory data) external payable onlyKeeper {
        _executeFunction(target, data);
    }

    /**
     * @dev Internal function execution logic
     */
    function _executeFunction(address target, bytes memory callData) internal returns (bytes memory) {
        (bool success, bytes memory result) = target.call{value: msg.value}(callData);
        require(success, "Function execution failed");

        emit FunctionExecuted(msg.sender, target, callData, msg.value);
        return result;
    }

    function _lzReceive(
        Origin calldata _origin,
        bytes32 /* _guid */,
        bytes calldata message,
        address /* _executor */,
        bytes calldata /* _extraData */
    ) internal override {
        (ActionType action, address keeper) = abi.decode(message, (ActionType, address));

        if (action == ActionType.REGISTER) {
            isKeeper[keeper] = true;
        } else if (action == ActionType.UNREGISTER) {
            isKeeper[keeper] = false;
        } else {
            revert("Invalid action type");
        }

        emit KeeperUpdated(action, keeper);
    }

    receive() external payable {}
}
