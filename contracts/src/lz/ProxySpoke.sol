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

    constructor(address _endpoint, address _owner) Ownable(_owner) OApp(_endpoint, _owner) {}

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
        bytes32,
        bytes calldata message,
        address,
        bytes calldata
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
