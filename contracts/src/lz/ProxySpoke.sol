// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import { OApp, Origin } from "@layerzero-v2/oapp/contracts/oapp/OApp.sol";
import { Ownable } from "@openzeppelin-contracts/contracts/access/Ownable.sol";

interface IJobRegistry {
    function getJobOwner(uint256 jobId) external view returns (address);
}

interface ITriggerGasRegistry {
    function deductTGBalance(address user, uint256 tgAmount) external;
}

/**
 * @title ProxySpoke
 * @notice A LayerZero-enabled contract that acts as a spoke in the keeper network
 * @dev This contract receives keeper registration updates from the hub and executes functions on the respective L2 chain
 */
contract ProxySpoke is Ownable, OApp {
    /**
     * @notice Enum defining the types of actions that can be performed on keepers
     */
    enum ActionType { REGISTER, UNREGISTER }

    /**
     * @notice Mapping to track registered keepers
     */
    mapping(address => bool) public isKeeper;

    /**
     * @notice The address of the job registry contract
     */
    IJobRegistry public jobRegistry;

    /**
     * @notice The address of the trigger gas registry contract
     */
    ITriggerGasRegistry public triggerGasRegistry;

    /**
     * @notice Emitted when a keeper's status is updated
     * @param action The type of action performed
     * @param keeper The address of the keeper
     */
    event KeeperUpdated(ActionType action, address keeper);

    /**
     * @notice Emitted when a function is executed
     * @param keeper The address of the keeper who executed the function
     * @param target The address of the target contract
     * @param data The calldata used for the function call
     * @param value The amount of ETH sent with the call
     */
    event FunctionExecuted(address indexed keeper, address indexed target, bytes data, uint256 value);

    /**
     * @notice Modifier to restrict function access to registered keepers
     */
    modifier onlyKeeper() {
        require(isKeeper[msg.sender], "Spoke: Keeper not registered");
        _;
    }

    /**
     * @notice Constructor for the ProxySpoke contract
     * @param _endpoint The LayerZero endpoint address
     * @param _ownerAddress The owner address
     * @param _hubEid The hub chain endpoint ID
     * @param _initialKeepers Array of initial keeper addresses
     * @param _jobRegistryAddress The address of the job registry contract
     * @param _triggerGasRegistryAddress The address of the trigger gas registry contract
     */
    constructor(
        address _endpoint,
        address _ownerAddress,
        uint32 _hubEid,
        address[] memory _initialKeepers,
        address _jobRegistryAddress,
        address _triggerGasRegistryAddress
    ) Ownable(_ownerAddress) OApp(_endpoint, _ownerAddress) {     
        _setPeer(_hubEid, bytes32(uint256(uint160(address(this)))));

        // Initialize keepers
        for (uint i = 0; i < _initialKeepers.length; i++) {
            isKeeper[_initialKeepers[i]] = true;
            emit KeeperUpdated(ActionType.REGISTER, _initialKeepers[i]);
        }

        jobRegistry = IJobRegistry(_jobRegistryAddress);
        triggerGasRegistry = ITriggerGasRegistry(_triggerGasRegistryAddress);
    }

    /**
     * @notice Executes a function on a target contract
     * @param jobId The ID of the job
     * @param tgAmount The amount of TG to deduct from the job owner
     * @param target The address of the target contract
     * @param data The calldata for the function call
     */
    function executeFunction(uint256 jobId, uint256 tgAmount, address target, bytes memory data) external payable onlyKeeper {
        address jobOwner = jobRegistry.getJobOwner(jobId);
        require(jobOwner != address(0), "Job not found");

        triggerGasRegistry.deductTGBalance(jobOwner, tgAmount);

        _executeFunction(target, data);
    }

    /**
     * @notice Internal function to execute a function call
     * @param target The address of the target contract
     * @param callData The calldata for the function call
     * @return The return data from the function call
     */
    // slither-disable-next-line reentrancy-events
    function _executeFunction(address target, bytes memory callData) internal returns (bytes memory) {
        // slither-disable-next-line low-level-calls
        (bool success, bytes memory result) = target.call{value: msg.value}(callData);
        require(success, "Function execution failed");

        emit FunctionExecuted(msg.sender, target, callData, msg.value);
        return result;
    }

    /**
     * @notice Handles incoming LayerZero messages from the hub
     * @param message The message payload containing the action and keeper address
     */
    function _lzReceive(
        Origin calldata /* _origin */,
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
}
