// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@eigenlayer-contracts/contracts/permissions/Pausable.sol";
import "@eigenlayer-middleware/libraries/BN254.sol";
import {OperatorStateRetriever} from "@eigenlayer-middleware/OperatorStateRetriever.sol";
import {IRegistryCoordinator} from "@eigenlayer-middleware/interfaces/IRegistryCoordinator.sol";
import {ITriggerXTaskManager} from "./ITriggerXTaskManager.sol";

contract TriggerXTaskManager is Initializable,
    OwnableUpgradeable,
    Pausable,
    OperatorStateRetriever,
    ITriggerXTaskManager
{
    // STATE VARIABLES
    uint32 public immutable TASK_RESPONSE_WINDOW_BLOCK;
    uint32 public constant TASK_CHALLENGE_WINDOW_BLOCK = 100;
    uint256 internal constant _THRESHOLD_DENOMINATOR = 100;

    IRegistryCoordinator public registryCoordinator;
    address public aggregator;

    mapping(uint32 => Task) public tasks;
    uint32 public latestTaskNum;

    constructor(IRegistryCoordinator _registryCoordinator, uint32 _taskResponseWindowBlock) 
    {
        registryCoordinator = _registryCoordinator;
        TASK_RESPONSE_WINDOW_BLOCK = _taskResponseWindowBlock;
    }

    function initialize(
        IPauserRegistry _pauserRegistry,
        address initialOwner,
        address _aggregator
    ) public initializer {
        _initializePauser(_pauserRegistry, UNPAUSE_ALL);
        _transferOwnership(initialOwner);
        aggregator = _aggregator;
    }

    function createTask(
        uint256 jobId
    ) external override {
        latestTaskNum++;
        uint32 taskId = latestTaskNum;
        tasks[taskId] = Task({
            jobId: jobId,
            taskId: taskId,
            status: "Created",
            blockNumber: block.number
        });

        emit TaskCreated(jobId, taskId);
    }

    function deleteTask(uint32 taskId) external override {
        delete tasks[taskId];
        emit TaskDeleted(tasks[taskId].jobId, taskId);
    }

    function updateTaskStatus(uint32 taskId, string calldata status) external override {
        tasks[taskId].status = status;
        emit TaskStatusUpdated(tasks[taskId].jobId, taskId, status);
    }

    function assignTask(uint32 taskId, address operator) external override {
        emit TaskAssigned(tasks[taskId].jobId, taskId, operator);
    }

    function respondToTask(
        uint32 taskId,
        TaskResponse calldata taskResponse,
        TaskResponseMetadata calldata taskResponseMetadata,
        BN254.G1Point[] memory pubkeysOfNonSigningOperators
    ) external override {
        emit TaskResponded(tasks[taskId].jobId, taskId, taskResponse, taskResponseMetadata);
    }
}