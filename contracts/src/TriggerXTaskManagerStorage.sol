// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@eigenlayer-contracts/contracts/permissions/Pausable.sol";
import "@eigenlayer-middleware/libraries/BN254.sol";
import {OperatorStateRetriever} from "@eigenlayer-middleware/OperatorStateRetriever.sol";
import {BLSSignatureChecker, IRegistryCoordinator} from "@eigenlayer-middleware/BLSSignatureChecker.sol";

contract TriggerXTaskManagerStorage is OwnableUpgradeable, Pausable, BLSSignatureChecker, OperatorStateRetriever {
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

    constructor(uint32 _taskResponseWindowBlock, IRegistryCoordinator _registryCoordinator) 
        BLSSignatureChecker(_registryCoordinator)
    {
        TASK_RESPONSE_WINDOW_BLOCK = _taskResponseWindowBlock;
    }

    function initializeStorage(address initialOwner, address _aggregator) public initializer {
        _transferOwnership(initialOwner);
        aggregator = _aggregator;
    }

    modifier onlyAggregator() {
        require(msg.sender == aggregator, "Aggregator must be the caller");
        _;
    }

    function incrementTaskNum() internal {
        latestTaskNum++;
    }

    function updateTask(uint32 taskId, Task memory task) internal {
        tasks[taskId] = task;
        allTaskHashes[taskId] = keccak256(abi.encode(task));
    }

    function recordTaskResponse(uint32 taskId, bytes32 responseHash) internal {
        allTaskResponses[taskId] = responseHash;
    }
}
