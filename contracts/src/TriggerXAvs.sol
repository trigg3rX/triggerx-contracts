// SPDX-License-Identifier: MIT
pragma solidity 0.8.30;

import "./interfaces/IAVSManager.sol" as avs;
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/proxy/utils/UUPSUpgradeable.sol";
import "./libraries/TaskDefinitionLibrary.sol";

/// @title TriggerX AVS (Imua compatible)
/// @notice Upgradeable contract acting as TriggerX's AVS proxy to the Imua AVS-Manager precompile.
///         Initially provides thin wrappers around IAVSManager with modular hooks for future
///         reward and slashing logic.
contract TriggerXAvs is OwnableUpgradeable, UUPSUpgradeable {
    // ---------------------------------------------------------------------
    // Storage
    // ---------------------------------------------------------------------

    /// @dev Future on-chain reward handler. 0x0 until decided.
    address public rewardManager;

    /// @dev Future on-chain slashing handler. 0x0 until decided.
    address public slasher;

    event OperatorOptedIn(address indexed operator);
    event OperatorOptedOut(address indexed operator);
    event BLSPublicKeyRegistered(address indexed operator, address indexed avsAddress, bytes pubKey);
    event TaskSubmitted(uint64 indexed taskID, address indexed operator, uint8 phase);
    event ChallengeSubmitted(uint64 indexed taskID, address indexed challenger, bool isExpected);
    event TaskDefinitionCreated(uint8 indexed taskDefinitionId, string name);
    event TaskCreated(uint64 indexed taskID, bytes32 definitionHash, uint8 kind);

    using TaskDefinitionLibrary for TaskDefinitions;
    TaskDefinitions private _taskDefs;

    // ---------------------------------------------------------------------
    // Initializer & Upgrade Authorization
    // ---------------------------------------------------------------------

    /// @dev Upgradeable initializer replacing constructor
    function initialize(address initialOwner) external initializer {
        __Ownable_init(initialOwner);
    }

    /// @inheritdoc UUPSUpgradeable
    function _authorizeUpgrade(address) internal override onlyOwner {}

    // ---------------------------------------------------------------------
    // Admin setters – keep modules pluggable
    // ---------------------------------------------------------------------

    event RewardManagerUpdated(address indexed newRewardManager);
    event SlasherUpdated(address indexed newSlasher);

    function setRewardManager(address _rewardManager) external onlyOwner {
        rewardManager = _rewardManager;
        emit RewardManagerUpdated(_rewardManager);
    }

    function setSlasher(address _slasher) external onlyOwner {
        slasher = _slasher;
        emit SlasherUpdated(_slasher);
    }

    // ---------------------------------------------------------------------
    // AVS-Manager wrappers
    // ---------------------------------------------------------------------

    function registerAVS(avs.AVSParams calldata params) external onlyOwner returns (bool success) {
        success = avs.AVSMANAGER_CONTRACT.registerAVS(params);
    }

    function updateAVS(avs.AVSParams calldata params) external onlyOwner returns (bool success) {
        success = avs.AVSMANAGER_CONTRACT.updateAVS(params);
    }

    function registerOperatorToAVS() external returns (bool success) {
        success = avs.AVSMANAGER_CONTRACT.registerOperatorToAVS(msg.sender);
        if (success) emit OperatorOptedIn(msg.sender);
    }

    function deregisterOperatorFromAVS() external returns (bool success) {
        success = avs.AVSMANAGER_CONTRACT.deregisterOperatorFromAVS(msg.sender);
        if (success) emit OperatorOptedOut(msg.sender);
    }

    function registerBLSPublicKey(
        address avsAddr,
        bytes calldata pubKey,
        bytes calldata pubKeyRegistrationSignature
    ) external returns (bool success) {
        success = avs.AVSMANAGER_CONTRACT.registerBLSPublicKey(
            msg.sender,
            avsAddr,
            pubKey,
            pubKeyRegistrationSignature
        );
        if (success) emit BLSPublicKeyRegistered(msg.sender, avsAddr, pubKey);
    }

     /**
     * @notice Register a new task with Imua AVS-Manager.
     * @param taskName The name of the task.
     * @param taskDefinitionId The ID of the task definition.
     * @param taskResponsePeriod The response period of the task.
     * @param taskChallengePeriod The challenge period of the task.
     * @param thresholdPercentage The threshold percentage of the task.
     * @param taskStatisticalPeriod The statistical period of the task.
     * @return taskID The task identifier returned by the precompile.
     */
    function createTask(
        string memory taskName,
        uint8 taskDefinitionId,
        uint64 taskResponsePeriod,
        uint64 taskChallengePeriod,
        uint8 thresholdPercentage,
        uint64 taskStatisticalPeriod        
    ) external returns (uint64 taskID) {

        TaskDefinition memory taskDef = _taskDefs.getTaskDefinition(taskDefinitionId);
        bytes32 defHash = keccak256(abi.encode(taskDef));

        require(
            thresholdPercentage <= 100,
            "The threshold cannot be greater than 100"
        );
        
        taskID = avs.AVSMANAGER_CONTRACT.createTask(
            msg.sender,
            taskName,
            abi.encodePacked(defHash),
            taskResponsePeriod,
            taskChallengePeriod,
            thresholdPercentage,
            taskStatisticalPeriod
        );
        
        emit TaskCreated(taskID, defHash, taskDef.taskDefinitionId);
    }

    function operatorSubmitTask(
        uint64 taskID,
        bytes calldata taskResponse,
        bytes calldata blsSignature,
        address taskContractAddress,
        uint8 phase
    ) external returns (bool success) {
        success = avs.AVSMANAGER_CONTRACT.operatorSubmitTask(
            msg.sender,
            taskID,
            taskResponse,
            blsSignature,
            taskContractAddress,
            phase
        );
        if (success) emit TaskSubmitted(taskID, msg.sender, phase);
    }

    /**
     * @notice Initiate a challenge for a specific task.
     * @param taskID               The task identifier on Imua.
     * @param actualThreshold      Observed threshold of signatures.
     * @param isExpected           Whether the task outcome meets expectations.
     * @param eligibleRewardOperators Operators eligible for rewards.
     * @param eligibleSlashOperators  Operators eligible for slash.
     */
    function challenge(
        uint64 taskID,
        uint8 actualThreshold,
        bool isExpected,
        address[] calldata eligibleRewardOperators,
        address[] calldata eligibleSlashOperators
    ) external returns (bool success) {
        success = avs.AVSMANAGER_CONTRACT.challenge(
            msg.sender,
            taskID,
            address(this),
            actualThreshold,
            isExpected,
            eligibleRewardOperators,
            eligibleSlashOperators
        );
        if (success) emit ChallengeSubmitted(taskID, msg.sender, isExpected);
    }

    // ---------------------------------------------------------------------
    // View helpers delegating to AVS-Manager
    // ---------------------------------------------------------------------

    function getOptInOperators(address avsAddress) external view returns (address[] memory) {
        return avs.AVSMANAGER_CONTRACT.getOptInOperators(avsAddress);
    }

    function getRegisteredPubkey(address operator, address avsAddr) external view returns (bytes memory) {
        return avs.AVSMANAGER_CONTRACT.getRegisteredPubkey(operator, avsAddr);
    }

    function getAVSUSDValue(address avsAddr) external view returns (uint256) {
        return avs.AVSMANAGER_CONTRACT.getAVSUSDValue(avsAddr);
    }

    function getOperatorOptedUSDValue(address avsAddr, address operatorAddr) external view returns (uint256) {
        return avs.AVSMANAGER_CONTRACT.getOperatorOptedUSDValue(avsAddr, operatorAddr);
    }

    function getAVSEpochIdentifier(address avsAddr) external view returns (string memory) {
        return avs.AVSMANAGER_CONTRACT.getAVSEpochIdentifier(avsAddr);
    }

    function getTaskInfo(address taskAddress, uint64 taskID) external view returns (avs.TaskInfo memory) {
        return avs.AVSMANAGER_CONTRACT.getTaskInfo(taskAddress, taskID);
    }

    function isOperator(address operator) external view returns (bool) {
        return avs.AVSMANAGER_CONTRACT.isOperator(operator);
    }

    function getCurrentEpoch(string calldata epochIdentifier) external view returns (int64) {
        return avs.AVSMANAGER_CONTRACT.getCurrentEpoch(epochIdentifier);
    }

    function getChallengeInfo(address taskAddress, uint64 taskID) external view returns (address) {
        return avs.AVSMANAGER_CONTRACT.getChallengeInfo(taskAddress, taskID);
    }

    function getOperatorTaskResponse(
        address taskAddress,
        address operator,
        uint64 taskID
    ) external view returns (avs.TaskResultInfo memory) {
        return avs.AVSMANAGER_CONTRACT.getOperatorTaskResponse(taskAddress, operator, taskID);
    }

    function getOperatorTaskResponseList(
        address taskAddress,
        uint64 taskID
    ) external view returns (avs.OperatorResInfo[] memory) {
        return avs.AVSMANAGER_CONTRACT.getOperatorTaskResponseList(taskAddress, taskID);
    }

    // ---------------------------------------------------------------------
    // Task Definition functions
    // ---------------------------------------------------------------------

    /**
     * @notice Creates TaskDefinition and stores it locally.
     */
         function createTaskDefinition(
         string calldata name,
         TaskDefinitionParams calldata defParams
     ) external onlyOwner returns (uint8 taskDefinitionId) {
         taskDefinitionId = _taskDefs.createNewTaskDefinition(name, defParams);
         emit TaskDefinitionCreated(taskDefinitionId, name);
     }

     function getTaskDefinition(uint8 id) external view returns (TaskDefinition memory) {
         return _taskDefs.getTaskDefinition(id);
     }

} 