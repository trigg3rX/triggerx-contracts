// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@eigenlayer-contracts/contracts/permissions/Pausable.sol";
import "@eigenlayer-middleware/libraries/BN254.sol";
import {IRegistryCoordinator} from "@eigenlayer-middleware/interfaces/IRegistryCoordinator.sol";
import {ITriggerXTaskManager} from "./ITriggerXTaskManager.sol";

contract TriggerXTaskManager is Initializable,
    OwnableUpgradeable,
    Pausable,
    ITriggerXTaskManager
{
    // STATE VARIABLES
    uint32 public immutable TASK_RESPONSE_WINDOW_BLOCK;
    uint32 public constant TASK_CHALLENGE_WINDOW_BLOCK = 100;
    uint256 internal constant _THRESHOLD_DENOMINATOR = 100;

    IRegistryCoordinator public registryCoordinator;
    address public aggregator;

    mapping(uint32 => Task) public tasks;
    uint32 public taskCount;

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
        string calldata taskType,
        string calldata status
    ) external {
        taskCount++;
        tasks[taskCount] = Task({
            taskId: taskCount,
            taskType: taskType,
            status: status,
            blockNumber: block.number
        });

        emit TaskCreated(taskCount, taskType);
    }

    function deleteTask(uint32 taskId) external {
        require(keccak256(abi.encodePacked(tasks[taskId].status)) == keccak256(abi.encodePacked("Completed")), "Task not completed");
        delete tasks[taskId];
        emit TaskDeleted(taskId);
    }

    function updateTaskStatus(uint32 taskId, string calldata status) external {
        require(keccak256(abi.encodePacked(tasks[taskId].status)) == keccak256(abi.encodePacked("Assigned")), "Task not assigned");
        tasks[taskId].status = status;
        emit TaskStatusUpdated(taskId, status);
    }

    function assignTask(uint32 taskId, address operator) external {
        require(keccak256(abi.encodePacked(tasks[taskId].status)) == keccak256(abi.encodePacked("Created")), "Task not created");
        tasks[taskId].status = "Assigned";
        tasks[taskId].blockNumber = block.number;
        emit TaskAssigned(taskId, operator);
    }

    function respondToTask(
        uint32 taskId,
        TaskResponse calldata taskResponse,
        TaskResponseMetadata calldata taskResponseMetadata,
        BN254.G1Point[] memory pubkeysOfNonSigningOperators
    ) external {
        require(keccak256(abi.encodePacked(tasks[taskId].status)) == keccak256(abi.encodePacked("Assigned")), "Task not assigned");
        require(tasks[taskId].blockNumber + TASK_RESPONSE_WINDOW_BLOCK < block.number, "Task response window not expired");
        require(taskResponse.referenceTaskId == taskId, "Reference task id does not match");
        // require(verifyNonSigner(taskResponseMetadata.hashOfNonSigners, pubkeysOfNonSigningOperators), "Non signers not verified");

        tasks[taskId].status = "Completed";
        emit TaskCompleted(taskId);
    }
}