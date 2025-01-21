// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import "./interfaces/ITriggerXTaskManager.sol";
import "./interfaces/ITriggerXServiceManager.sol";
import {EnumerableSet} from "@openzeppelin-contracts/contracts/utils/structs/EnumerableSet.sol";
import {Pausable} from "@eigenlayer-contracts/contracts/permissions/Pausable.sol";
// import {IAllocationManager} from "@eigenlayer-contracts/contracts/interfaces/IAllocationManager.sol";
import {IPermissionController} from "@eigenlayer-contracts/contracts/interfaces/IPermissionController.sol";
import {IPauserRegistry} from "@eigenlayer-contracts/contracts/interfaces/IPauserRegistry.sol";
import {ServiceManagerBase, IAVSDirectory, IRewardsCoordinator, IServiceManager, ISignatureUtils} from "@eigenlayer-middleware/ServiceManagerBase.sol";
import {BLSSignatureChecker} from "@eigenlayer-middleware/BLSSignatureChecker.sol";
import {IRegistryCoordinator} from "@eigenlayer-middleware/interfaces/IRegistryCoordinator.sol";
import {IStakeRegistry} from "@eigenlayer-middleware/interfaces/IStakeRegistry.sol";
import {VetoableSlasher} from "@eigenlayer-middleware/slashers/VetoableSlasher.sol";


contract TriggerXServiceManager is
    Pausable,
    BLSSignatureChecker,
    ServiceManagerBase,
    ITriggerXServiceManager
{
    using EnumerableSet for EnumerableSet.AddressSet;

    EnumerableSet.AddressSet private _registeredOperators;

    mapping(address => bool) public override isBlackListed;

    address public override taskManager;
    address public override taskValidator;
    address public override quorumManager;
    address public vetoCommittee;

    ITriggerXTaskManager public override taskManagerContract;

    VetoableSlasher public immutable vetoableSlasher;


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
        IStakeRegistry _stakeRegistry,
        // IAllocationManager _allocationManager,
        IPermissionController _permissionController,
        IPauserRegistry _pauserRegistry,
        VetoableSlasher _vetoableSlasher
    ) 
        Pausable(_pauserRegistry)
        BLSSignatureChecker(_registryCoordinator)
        ServiceManagerBase(_avsDirectory, _rewardsCoordinator, _registryCoordinator, _stakeRegistry, _permissionController)
    {
        vetoableSlasher = _vetoableSlasher;
        _disableInitializers();
    }

    function initialize(
        ITriggerXTaskManager _taskManagerContract,
        address initialOwner,
        address rewardsInitiator,
        address _slasher,
        address _vetoCommittee,
        address _taskManager,
        address _taskValidator,
        address _quorumManager
    ) public initializer {
        _transferOwnership(initialOwner);
        __ServiceManagerBase_init(initialOwner, rewardsInitiator);

        taskManagerContract = _taskManagerContract;
        taskManager = _taskManager;
        taskValidator = _taskValidator;
        quorumManager = _quorumManager;
        slasher = _slasher;
        vetoCommittee = _vetoCommittee;
    }

    function setTaskManager(address _taskManager) external override onlyOwner {
        address oldTaskManager = taskManager;
        taskManager = _taskManager;
        emit TaskManagerSet(oldTaskManager, _taskManager);
    }

    function setTaskValidator(address _taskValidator) external override onlyOwner {
        address oldTaskValidator = taskValidator;
        taskValidator = _taskValidator;
        emit TaskValidatorSet(oldTaskValidator, _taskValidator);
    }

    function setQuorumManager(address _quorumManager) external override onlyOwner {
        address oldQuorumManager = quorumManager;
        quorumManager = _quorumManager;
        emit QuorumManagerSet(oldQuorumManager, _quorumManager);
    }

    function updateTaskManager(address _taskManager) external override onlyOwner {
        address oldTaskManager = address(taskManagerContract);
        taskManagerContract = ITriggerXTaskManager(_taskManager);
        emit TaskManagerContractUpdated(oldTaskManager, _taskManager);
    }

    function registerOperatorToAVS(
        address operator,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) public override(ServiceManagerBase, ITriggerXServiceManager) onlyRegistryCoordinator() {
        _avsDirectory.registerOperatorToAVS(operator, operatorSignature);
        _registeredOperators.add(operator);
        isBlackListed[operator] = false;
        emit KeeperAdded(operator);
    }

    function deregisterOperatorFromAVS(address operator) public override(ServiceManagerBase, ITriggerXServiceManager) onlyRegistryCoordinator() {
        _avsDirectory.deregisterOperatorFromAVS(operator);
        _registeredOperators.remove(operator);
        isBlackListed[operator] = false;
        emit KeeperRemoved(operator);
    }

    function blacklistKeeper(address _operator) external override onlyOwner {
        isBlackListed[_operator] = true;
        emit KeeperBlacklisted(_operator);
    }

    function unblacklistKeeper(address _operator) external override onlyOwner {
        isBlackListed[_operator] = false;
        emit KeeperUnblacklisted(_operator);
    }
}
