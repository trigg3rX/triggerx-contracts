// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@eigenlayer-middleware/libraries/BN254.sol";

interface ITriggerXTaskManager {
    // EVENTS
    event TaskCreated(uint32 jobId, uint32 taskId);

    event TaskAssigned(Task task, address operator);
    
    event TaskResponded(
        uint32 jobId,
        TaskResponse taskResponse,
        TaskResponseMetadata taskResponseMetadata
    );
    
    event TaskCompleted(Task task);
    
    event TaskChallengedSuccessfully(
        Task task,
        address indexed challenger
    );

    event TaskChallengedUnsuccessfully(
        Task task,
        address indexed challenger
    );

    // STRUCTS
    struct Task {
        uint32 jobId;
        uint32 blockNumber;
        bytes quorumNumbers;
        uint8 quorumThresholdPercentage;
    }

    struct TaskResponse {
        uint32 referenceTaskId;
        bytes32 responseDataHash;
    }

    struct TaskResponseMetadata {
        uint32 taskResponsedBlock;
        bytes32 hashOfNonSigners;
    }

    // FUNCTIONS
    function taskNumber() external view returns (uint32);

    function createTask(
        uint32 jobId,
        bytes calldata quorumNumbers,
        uint8 quorumThresholdPercentage
    ) external returns (uint32);

    function getTaskResponseWindowBlock() external view returns (uint32);

    // function raiseAndResolveChallenge(
    //     uint32 taskId,
    //     TaskResponse calldata taskResponse,
    //     TaskResponseMetadata calldata taskResponseMetadata,
    //     BN254.G1Point[] memory pubkeysOfNonSigningOperators
    // ) external;
}