// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@eigenlayer-middleware/libraries/BN254.sol";

interface ITriggerXTaskManager {
    // EVENTS
    event NewTaskCreated(uint32 indexed taskIndex, Task task);

    event TaskResponded(
        TaskResponse taskResponse,
        TaskResponseMetadata taskResponseMetadata
    );

    event TaskCompleted(uint32 indexed taskIndex);

    event TaskChallengedSuccessfully(
        uint32 indexed taskIndex,
        address indexed challenger
    );

    event TaskChallengedUnsuccessfully(
        uint32 indexed taskIndex,
        address indexed challenger
    );

    event AggregatorUpdated(address indexed oldAggregator, address indexed newAggregator);

    // STRUCTS
    struct Task {
        uint32 jobId;
        uint32 taskCreatedBlock;
        uint32 quorumThresholdPercentage;
        bytes quorumNumbers;
    }

    struct TaskResponse {
        uint32 referenceTaskIndex;
        bytes32 transactionHash;
    }

    struct TaskResponseMetadata {
        uint32 taskResponsedBlock;
        bytes32 hashOfNonSigners;
    }

    // FUNCTIONS
    function createNewTask(
        uint32 jobId,
        uint32 quorumThresholdPercentage,
        bytes calldata quorumNumbers
    ) external;

    function taskNumber() external view returns (uint32);

    function getTaskResponseWindowBlock() external view returns (uint32);
}