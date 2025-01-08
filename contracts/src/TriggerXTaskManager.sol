// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@eigenlayer-contracts/contracts/permissions/Pausable.sol";
import {BLSSignatureChecker, IRegistryCoordinator} from "@eigenlayer-middleware/BLSSignatureChecker.sol";
import {BN254} from "@eigenlayer-middleware/libraries/BN254.sol";
import "./interfaces/ITriggerXTaskManager.sol";
import "./TriggerXServiceManager.sol";

contract TriggerXTaskManager is
    Initializable,
    OwnableUpgradeable,
    Pausable,
    BLSSignatureChecker,
    ITriggerXTaskManager
{
    using BN254 for BN254.G1Point;

    uint32 public TASK_RESPONSE_WINDOW_BLOCK;
    uint256 internal constant _THRESHOLD_DENOMINATOR = 100;

    mapping(uint32 => uint32) public jobToTaskCounter;
    mapping(bytes8 => bytes32) public taskHashes;
    mapping(bytes8 => bytes32) public taskResponseHashes;

    modifier onlyTaskManager() {
        require(msg.sender == serviceManager.taskManager(), "Only the task manager can call this function");
        _;
    }

    modifier onlyTaskValidator() {
        require(msg.sender == serviceManager.taskValidator(), "Only the task validator can call this function");
        _;
    }

    TriggerXServiceManager public serviceManager;

    constructor(IRegistryCoordinator _registryCoordinator) 
        BLSSignatureChecker(_registryCoordinator)
    {
        _disableInitializers();
    }

    function initialize(
        IPauserRegistry _pauserRegistry,
        address initialOwner,
        uint32 _taskResponseWindowBlock,
        TriggerXServiceManager _serviceManager
    ) public initializer {
        _initializePauser(_pauserRegistry, UNPAUSE_ALL);
        _transferOwnership(initialOwner);

        TASK_RESPONSE_WINDOW_BLOCK = _taskResponseWindowBlock;
        serviceManager = _serviceManager;
    }

    function generateTaskId(uint32 jobId, uint32 taskNum) public pure returns (bytes8) {
        return bytes8(abi.encodePacked(jobId, taskNum));
    }

    function createNewTask(
        uint32 jobId,
        bytes calldata quorumNumbers,
        uint8 quorumThreshold
    ) external onlyTaskManager {
        Task memory newTask;
        newTask.jobId = jobId;
        newTask.taskNum = jobToTaskCounter[jobId];
        newTask.taskCreatedBlock = uint32(block.number);
        newTask.quorumNumbers = quorumNumbers;
        newTask.quorumThreshold = quorumThreshold;

        bytes8 taskId = generateTaskId(jobId, jobToTaskCounter[jobId]);

        taskHashes[taskId] = keccak256(abi.encode(newTask));
        emit TaskCreated(taskId, taskHashes[taskId]);
        jobToTaskCounter[jobId] = jobToTaskCounter[jobId] + 1;
    }

    function respondToTask(
        Task calldata task,
        TaskResponse calldata taskResponse,
        NonSignerStakesAndSignature memory nonSignerStakesAndSignature
    ) external onlyTaskValidator {
        // uint32 taskCreatedBlock = task.taskCreatedBlock;
        // bytes calldata quorumNumbers = task.quorumNumbers;
        // uint8 quorumThreshold = task.quorumThreshold;

        // check that the task is valid, hasn't been responded to yet, and is being responded to in time
        // require(
        //     keccak256(abi.encode(task)) == allTaskHashes[taskResponse.referenceTaskIndex],
        //     "supplied task does not match the one recorded in the contract"
        // );
        // require(
        //     taskResponseHashes[taskResponse.taskId] == bytes32(0),
        //     "Validator has already responded to the task"
        // );
        // require(
        //     uint32(block.number) <= taskCreatedBlock + TASK_RESPONSE_WINDOW_BLOCK,
        //     "Validator has responded to the task too late"
        // );

        bytes32 message = keccak256(abi.encode(taskResponse));
        
        (
            QuorumStakeTotals memory quorumStakeTotals,
            bytes32 hashOfNonSigners
        ) = checkSignatures(
                message,
                task.quorumNumbers,
                task.taskCreatedBlock,
                nonSignerStakesAndSignature
            );

        for (uint256 i = 0; i < task.quorumNumbers.length; i++) {
            require(
                quorumStakeTotals.signedStakeForQuorum[i] *
                    _THRESHOLD_DENOMINATOR >=
                    quorumStakeTotals.totalStakeForQuorum[i] *
                        task.quorumThreshold,
                "Signatories do not own at least threshold percentage of a quorum"
            );
        }

        TaskResponseMetadata memory taskResponseMetadata = TaskResponseMetadata(
            uint32(block.number),
            hashOfNonSigners
        );
        
        taskResponseHashes[taskResponse.taskId] = keccak256(
            abi.encode(taskResponse, taskResponseMetadata)
        );

        emit TaskResponded(taskResponse.taskId, taskResponseHashes[taskResponse.taskId]);
    }

    function updateServiceManager(address _serviceManager) external onlyOwner {
        serviceManager = TriggerXServiceManager(_serviceManager);
    }
}