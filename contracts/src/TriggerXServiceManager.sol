// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import "./interfaces/ITriggerXTaskManager.sol";
import {EnumerableSet} from "@openzeppelin-contracts/contracts/utils/structs/EnumerableSet.sol";
import {Pausable} from "@eigenlayer-contracts/contracts/permissions/Pausable.sol";
import {IPauserRegistry} from "@eigenlayer-contracts/contracts/interfaces/IPauserRegistry.sol";

import {ServiceManagerBase, IAVSDirectory, IRewardsCoordinator, IServiceManager, ISignatureUtils} from "@eigenlayer-middleware/ServiceManagerBase.sol";
import {BLSSignatureChecker} from "@eigenlayer-middleware/BLSSignatureChecker.sol";
import {IRegistryCoordinator} from "@eigenlayer-middleware/interfaces/IRegistryCoordinator.sol";
import {IStakeRegistry} from "@eigenlayer-middleware/interfaces/IStakeRegistry.sol";


contract TriggerXServiceManager is
    ServiceManagerBase,
    BLSSignatureChecker, 
    Pausable
{
    using EnumerableSet for EnumerableSet.AddressSet;

    event KeeperAdded(address indexed operator);
    event KeeperRemoved(address indexed operator);
    event TaskManagerSet(address indexed oldTaskManager, address indexed newTaskManager);
    event TaskValidatorSet(address indexed oldTaskValidator, address indexed newTaskValidator);
    event QuorumManagerSet(address indexed oldQuorumManager, address indexed newQuorumManager);
    event TaskManagerContractUpdated(address indexed oldTaskManager, address indexed newTaskManager);
    event KeeperBlacklisted(address indexed operator);
    event KeeperUnblacklisted(address indexed operator);

    EnumerableSet.AddressSet private _registeredOperators;

    mapping(address => bool) public isBlackListed;

    address public taskManager;
    address public taskValidator;
    address public quorumManager;

    ITriggerXTaskManager public taskManagerContract;

    modifier onlyTaskManagerContract() {
        require(
            msg.sender == address(taskManagerContract), 
            "TriggerXServiceManager: Only TriggerXTaskManagerContract can call this function");
        _;
    }

    constructor(
        IAVSDirectory _avsDirectory,
        IRewardsCoordinator _rewardsCoordinator,
        IRegistryCoordinator _registryCoordinator,
        IStakeRegistry _stakeRegistry       
    ) 
        BLSSignatureChecker(_registryCoordinator)
        ServiceManagerBase(_avsDirectory, _rewardsCoordinator, _registryCoordinator, _stakeRegistry) 
    {
        _disableInitializers();
    }

    function initialize(
        ITriggerXTaskManager _taskManagerContract,
        IPauserRegistry _pauserRegistry,
        uint256 _initialPausedStatus,
        address initialOwner,
        address rewardsInitiator,
        address _taskManager,
        address _taskValidator,
        address _quorumManager
    ) public initializer {
        _initializePauser(_pauserRegistry, _initialPausedStatus);
        _transferOwnership(initialOwner);
        __ServiceManagerBase_init(initialOwner, rewardsInitiator);

        taskManagerContract = _taskManagerContract;
        taskManager = _taskManager;
        taskValidator = _taskValidator;
        quorumManager = _quorumManager;
    }

    function setTaskManager(address _taskManager) external onlyOwner {
        address oldTaskManager = taskManager;
        taskManager = _taskManager;
        emit TaskManagerSet(oldTaskManager, _taskManager);
    }

    function setTaskValidator(address _taskValidator) external onlyOwner {
        address oldTaskValidator = taskValidator;
        taskValidator = _taskValidator;
        emit TaskValidatorSet(oldTaskValidator, _taskValidator);
    }

    function setQuorumManager(address _quorumManager) external onlyOwner {
        address oldQuorumManager = quorumManager;
        quorumManager = _quorumManager;
        emit QuorumManagerSet(oldQuorumManager, _quorumManager);
    }

    function updateTaskManager(address _taskManager) external onlyOwner {
        address oldTaskManager = address(taskManagerContract);
        taskManagerContract = ITriggerXTaskManager(_taskManager);
        emit TaskManagerContractUpdated(oldTaskManager, _taskManager);
    }

    function registerKeeperToTriggerX(
        address operator,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) public {
        _avsDirectory.registerOperatorToAVS(operator, operatorSignature);
        _registeredOperators.add(operator);
        isBlackListed[operator] = false;
        emit KeeperAdded(operator);
    }

    function deregisterKeeperFromTriggerX(address operator) external {
        _avsDirectory.deregisterOperatorFromAVS(operator);
        _registeredOperators.remove(operator);
        isBlackListed[operator] = false;
        emit KeeperRemoved(operator);
    }

    function blacklistKeeper(address _operator) external onlyOwner {
        isBlackListed[_operator] = true;
        emit KeeperBlacklisted(_operator);
    }

    function unblacklistKeeper(address _operator) external onlyOwner {
        isBlackListed[_operator] = false;
        emit KeeperUnblacklisted(_operator);
    }
}
