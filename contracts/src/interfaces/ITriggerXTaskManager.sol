// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {IBLSSignatureChecker} from "@eigenlayer-middleware/interfaces/IBLSSignatureChecker.sol";

interface ITriggerXTaskManager {
    // EVENTS
    event TaskCreated(
        bytes8 taskId,
        bytes32 taskHash
    );

    event TaskResponded(
        bytes8 taskId,
        bytes32 taskResponseHash
    );

    // STRUCTS
    struct Task {
        uint32 jobId;
        uint32 taskNum;
        uint32 taskCreatedBlock;
        bytes quorumNumbers;
        uint8 quorumThreshold;
    }

    struct TaskResponse {
        bytes8 taskId;
        bytes32 taskResponseHash;
    }

    struct TaskResponseMetadata {
        uint32 taskResponsedBlock;
        bytes32 hashOfNonSigners;
    }

    // FUNCTIONS
    function createNewTask(
        uint32 jobId,
        bytes calldata quorumNumbers,
        uint8 quorumThreshold
    ) external;

    function respondToTask(
        Task calldata task,
        TaskResponse calldata taskResponse,
        IBLSSignatureChecker.NonSignerStakesAndSignature memory nonSignerStakesAndSignature
    ) external;
}