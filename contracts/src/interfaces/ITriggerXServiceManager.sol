// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {ISignatureUtils} from "@eigenlayer-contracts/contracts/interfaces/ISignatureUtils.sol";
import {IServiceManager} from "@eigenlayer-middleware/interfaces/IServiceManager.sol";
import {ITriggerXTaskManager} from "./ITriggerXTaskManager.sol";

interface ITriggerXServiceManager is IServiceManager {
    // Events
    event KeeperAdded(address indexed operator);
    event KeeperRemoved(address indexed operator);
    event TaskManagerSet(address indexed oldTaskManager, address indexed newTaskManager);
    event TaskValidatorSet(address indexed oldTaskValidator, address indexed newTaskValidator);
    event QuorumManagerSet(address indexed oldQuorumManager, address indexed newQuorumManager);
    event TaskManagerContractUpdated(address indexed oldTaskManager, address indexed newTaskManager);
    event KeeperBlacklisted(address indexed operator);
    event KeeperUnblacklisted(address indexed operator);

    // View functions
    function isBlackListed(address operator) external view returns (bool);
    function taskManager() external view returns (address);
    function taskValidator() external view returns (address);
    function quorumManager() external view returns (address);
    function taskManagerContract() external view returns (ITriggerXTaskManager);

    // State-changing functions
    function setTaskManager(address _taskManager) external;
    function setTaskValidator(address _taskValidator) external;
    function setQuorumManager(address _quorumManager) external;
    function updateTaskManager(address _taskManager) external;
    
    function registerOperatorToAVS(
        address operator,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) external;
    
    function deregisterOperatorFromAVS(address operator) external;
    function blacklistKeeper(address _operator) external;
    function unblacklistKeeper(address _operator) external;
}