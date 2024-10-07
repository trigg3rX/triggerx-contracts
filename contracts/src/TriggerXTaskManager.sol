// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "./TriggerXTaskManagerStorage.sol";
import {ITriggerXTaskManager} from "./ITriggerXTaskManager.sol";

contract TriggerXTaskManager is TriggerXTaskManagerStorage, ITriggerXTaskManager {

    constructor(
        IRegistryCoordinator _registryCoordinator, 
        uint32 _taskResponseWindowBlock
    ) TriggerXTaskManagerStorage(_taskResponseWindowBlock, _registryCoordinator) {}

    function initialize(
        IPauserRegistry _pauserRegistry,
        address initialOwner,
        address _aggregator
    ) public initializer {
        _initializePauser(_pauserRegistry, UNPAUSE_ALL);
        initializeStorage(initialOwner, _aggregator);
    }

    function createTask(
        uint32 jobId,
        bytes calldata quorumNumbers,
        uint8 quorumThresholdPercentage 
    ) external override returns (uint32) {
        incrementTaskNum();
        uint32 taskId = latestTaskNum;

        Task memory newTask = Task({
            jobId: jobId,
            blockNumber: uint32(block.number),
            quorumNumbers: quorumNumbers,
            quorumThresholdPercentage: quorumThresholdPercentage
        });

        updateTask(taskId, newTask);

        emit TaskCreated(newTask.jobId, taskId);

        return taskId;
    }

    function respondToTask(
        Task calldata task,
        TaskResponse calldata taskResponse,
        NonSignerStakesAndSignature memory nonSignerStakesAndSignature
    ) external onlyAggregator {

        uint32 taskCreatedBlock = task.blockNumber;

        require(
            keccak256(abi.encode(task)) ==
                allTaskHashes[taskResponse.referenceTaskId],
            "supplied task does not match the one recorded in the contract"
        );

        require(
            allTaskResponses[taskResponse.referenceTaskId] == bytes32(0),
            "Aggregator has already responded to the task"
        );
        require(
            uint32(block.number) <=
                taskCreatedBlock + TASK_RESPONSE_WINDOW_BLOCK,
            "Aggregator has responded to the task too late"
        );

        bytes calldata quorumNumbers = task.quorumNumbers;
        uint8 quorumThresholdPercentage = task.quorumThresholdPercentage;
        bytes32 message = keccak256(abi.encode(taskResponse));

        (
            QuorumStakeTotals memory quorumStakeTotals,
            bytes32 hashOfNonSigners
        ) = checkSignatures(
            message, 
            quorumNumbers, 
            taskCreatedBlock, 
            nonSignerStakesAndSignature
        );

        for( uint i = 0; i < quorumNumbers.length; i++) {
            require(
                quorumStakeTotals.signedStakeForQuorum[i] * _THRESHOLD_DENOMINATOR >= quorumStakeTotals.totalStakeForQuorum[i] * quorumThresholdPercentage,
                "Signatories do not own at least threshold percentage of a quorum"
            );
        }

        TaskResponseMetadata memory taskResponseMetadata = TaskResponseMetadata(
            uint32(block.number),
            hashOfNonSigners
        );

        recordTaskResponse(taskResponse.referenceTaskId, keccak256(
            abi.encode(taskResponse, taskResponseMetadata)
        ));

        emit TaskResponded(task.jobId, taskResponse, taskResponseMetadata);
    }

    function taskNumber() external view returns (uint32) {
        return latestTaskNum;
    }

    function getTaskResponseWindowBlock() external view returns (uint32) {
        return TASK_RESPONSE_WINDOW_BLOCK;
    }
}