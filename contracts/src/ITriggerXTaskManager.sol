// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@eigenlayer-middleware/libraries/BN254.sol";

interface ITriggerXTaskManager {
    // EVENTS
    event TaskCreated(uint256 jobId, uint32 indexed taskId);
    event TaskDeleted(uint256 jobId, uint32 indexed taskId);
    event TaskStatusUpdated(uint256 jobId, uint32 indexed taskId, string status);
    event TaskAssigned(uint256 jobId, uint32 indexed taskId, address operator);
    event TaskCompleted(uint256 jobId, uint32 indexed taskId);
    event TaskResponded(
        uint256 jobId,
        uint32 indexed taskId,
        TaskResponse taskResponse,
        TaskResponseMetadata taskResponseMetadata
    );
    event TaskChallengedSuccessfully(
        uint256 jobId,
        uint32 indexed taskId,
        address indexed challenger
    );

    // STRUCTS
    struct Task {
        uint256 jobId;
        uint32 taskId;
        string status;
        uint256 blockNumber;
    }

    struct TaskResponse {
        uint256 referenceJobId;
        uint32 referenceTaskId;
        address completionTransactionHash;
    }

    struct TaskResponseMetadata {
        uint256 referenceJobId;
        uint32 referenceTaskId;
        uint256 taskResponsedBlock;
        bytes32 hashOfNonSigners;
    }

    // FUNCTIONS
    function createTask(
        uint256 jobId
    ) external;

    function deleteTask(uint32 taskId) external;

    function updateTaskStatus(uint32 taskId, string calldata status) external;

    function assignTask(uint32 taskId, address operator) external;

    function respondToTask(
        uint32 taskId,
        TaskResponse calldata taskResponse,
        TaskResponseMetadata calldata taskResponseMetadata,
        BN254.G1Point[] memory pubkeysOfNonSigningOperators
    ) external;

    // function raiseAndResolveChallenge(
    //     uint32 taskId,
    //     TaskResponse calldata taskResponse,
    //     TaskResponseMetadata calldata taskResponseMetadata,
    //     BN254.G1Point[] memory pubkeysOfNonSigningOperators
    // ) external;
}