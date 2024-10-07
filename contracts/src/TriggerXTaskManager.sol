// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@eigenlayer-contracts/contracts/permissions/Pausable.sol";
import {BLSSignatureChecker, IRegistryCoordinator} from "@eigenlayer-middleware/BLSSignatureChecker.sol";
import "./ITriggerXTaskManager.sol";
import "./TriggerXServiceManager.sol";

contract TriggerXTaskManager is
    Initializable,
    OwnableUpgradeable,
    Pausable,
    BLSSignatureChecker,
    ITriggerXTaskManager
{
    using BN254 for BN254.G1Point;

    uint32 public immutable TASK_RESPONSE_WINDOW_BLOCK;
    uint32 public constant TASK_CHALLENGE_WINDOW_BLOCK = 100;
    uint256 internal constant _THRESHOLD_DENOMINATOR = 100;

    uint32 public latestTaskNum;
    mapping(uint32 => bytes32) public allTaskHashes;
    mapping(uint32 => bytes32) public allTaskResponses;
    mapping(uint32 => bool) public taskSuccesfullyChallenged;

    address public aggregator;

    TriggerXServiceManager public serviceManager;

    modifier onlyAggregator() {
        require(msg.sender == aggregator, "Aggregator must be the caller");
        _;
    }

    modifier operatorNotBlacklisted(address operator) {
        require(!serviceManager.operatorBlacklist(operator), "Operator is blacklisted");
        _;
    }

    constructor(
        IRegistryCoordinator _registryCoordinator,
        uint32 _taskResponseWindowBlock,
        TriggerXServiceManager _serviceManager
    ) BLSSignatureChecker(_registryCoordinator) {
        TASK_RESPONSE_WINDOW_BLOCK = _taskResponseWindowBlock;
        serviceManager = _serviceManager;
    }

    function initialize(
        IPauserRegistry _pauserRegistry,
        address initialOwner,
        address _aggregator
    ) public initializer {
        _initializePauser(_pauserRegistry, UNPAUSE_ALL);
        _transferOwnership(initialOwner);
        _setAggregator(_aggregator);
    }

    function createNewTask(
        uint32 jobId,
        bytes calldata quorumNumbers
    ) external {
        Task memory newTask;
        newTask.jobId = jobId;
        newTask.taskCreatedBlock = uint32(block.number);
        newTask.quorumNumbers = quorumNumbers;

        allTaskHashes[latestTaskNum] = keccak256(abi.encode(newTask));
        emit NewTaskCreated(latestTaskNum, newTask);
        latestTaskNum = latestTaskNum + 1;
    }

    function _validateTask(
        Task calldata task,
        TaskResponse calldata taskResponse
    ) internal view {
        require(
            keccak256(abi.encode(task)) ==
                allTaskHashes[taskResponse.referenceTaskIndex],
            "Task mismatch"
        );
        require(
            allTaskResponses[taskResponse.referenceTaskIndex] == bytes32(0),
            "Already responded"
        );
        require(
            uint32(block.number) <=
                task.taskCreatedBlock + TASK_RESPONSE_WINDOW_BLOCK,
            "Response too late"
        );
    }

    function _checkQuorumThresholds(
        QuorumStakeTotals memory quorumStakeTotals,
        bytes calldata quorumNumbers,
        uint32 quorumThresholdPercentage
    ) internal pure {
        
    }

    function respondToTask(
        Task calldata task,
        TaskResponse calldata taskResponse,
        NonSignerStakesAndSignature memory nonSignerStakesAndSignature
    ) external onlyAggregator operatorNotBlacklisted(taskResponse.operator) {
        _validateTask(task, taskResponse);

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

        for (uint i = 0; i < task.quorumNumbers.length; i++) {
            require(
                quorumStakeTotals.signedStakeForQuorum[i] *
                    _THRESHOLD_DENOMINATOR >=
                    quorumStakeTotals.totalStakeForQuorum[i] *
                        uint8(serviceManager.quorumThresholdPercentage()),
                "Threshold not met"
            );
        }

        TaskResponseMetadata memory taskResponseMetadata = TaskResponseMetadata(
            uint32(block.number),
            hashOfNonSigners
        );
        
        allTaskResponses[taskResponse.referenceTaskIndex] = keccak256(
            abi.encode(taskResponse, taskResponseMetadata)
        );

        emit TaskResponded(taskResponse, taskResponseMetadata);
    }

    function taskNumber() external view returns (uint32) {
        return latestTaskNum;
    }

    function getTaskResponseWindowBlock() external view returns (uint32) {
        return TASK_RESPONSE_WINDOW_BLOCK;
    }

    function _setAggregator(address newAggregator) internal {
        address oldAggregator = aggregator;
        aggregator = newAggregator;
        emit AggregatorUpdated(oldAggregator, newAggregator);
    }

    function setAggregator(address newAggregator) external onlyOwner {
        _setAggregator(newAggregator);
    }
}