// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@eigenlayer-contracts/contracts/permissions/Pausable.sol";
import "@eigenlayer-middleware/libraries/BN254.sol";
import {OperatorStateRetriever} from "@eigenlayer-middleware/OperatorStateRetriever.sol";
import {BLSSignatureChecker, IRegistryCoordinator} from "@eigenlayer-middleware/BLSSignatureChecker.sol";
import "@eigenlayer-middleware/libraries/BN254.sol";
import {ITriggerXTaskManager} from "./ITriggerXTaskManager.sol";

contract TriggerXTaskManager is Initializable,
    OwnableUpgradeable,
    Pausable,
    BLSSignatureChecker,
    OperatorStateRetriever,
    ITriggerXTaskManager
{
    using BN254 for BN254.G1Point;

    // STATE VARIABLES
    uint32 public immutable TASK_RESPONSE_WINDOW_BLOCK;
    uint32 public constant TASK_CHALLENGE_WINDOW_BLOCK = 100;
    uint256 internal constant _THRESHOLD_DENOMINATOR = 100;

    address public aggregator;

    mapping(uint32 => Task) public tasks;

    mapping(uint32 => bytes32) public allTaskHashes;

    mapping(uint32 => bytes32) public allTaskResponses;

    uint32 public latestTaskNum;

    modifier onlyAggregator() {
        require(msg.sender == aggregator, "Aggregator must be the caller");
        _;
    }

    constructor(
        IRegistryCoordinator _registryCoordinator, 
        uint32 _taskResponseWindowBlock
    ) BLSSignatureChecker (_registryCoordinator) {
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
        uint32 jobId,
        bytes calldata quorumNumbers,
        uint8 quorumThresholdPercentage 
    ) external override returns (uint32) {
        latestTaskNum++;
        uint32 taskId = latestTaskNum;

        tasks[taskId] = Task({
            jobId: jobId,
            blockNumber: uint32(block.number),
            quorumNumbers: quorumNumbers,
            quorumThresholdPercentage: quorumThresholdPercentage
        });

        allTaskHashes[taskId] = keccak256(abi.encode(tasks[taskId]));

        emit TaskCreated(tasks[taskId].jobId, taskId);

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
        // some logical checks
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
        
        allTaskResponses[taskResponse.referenceTaskId] = keccak256(
            abi.encode(taskResponse, taskResponseMetadata)
        );

        emit TaskResponded(task.jobId, taskResponse, taskResponseMetadata);
    }

    function taskNumber() external view returns (uint32) {
        return latestTaskNum;
    }

    function getTaskResponseWindowBlock() external view returns (uint32) {
        return TASK_RESPONSE_WINDOW_BLOCK;
    }
}