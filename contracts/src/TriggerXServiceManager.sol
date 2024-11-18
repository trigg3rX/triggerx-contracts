// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "./ITriggerXTaskManager.sol";
import "@eigenlayer-middleware/ServiceManagerBase.sol";
import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/proxy/utils/UUPSUpgradeable.sol";

contract TriggerXServiceManager is 
    Initializable,
    ServiceManagerBase,
    UUPSUpgradeable 
{
    address public taskManager;
    address public taskValidator;
    address public quorumManager;

    ITriggerXTaskManager public immutable taskManagerContract;

    modifier onlyTaskManagerContract() {
        require(
            msg.sender == address(taskManagerContract), 
            "TriggerXServiceManager: Only TriggerXTaskManagerContract can call this function");
        _;
    }

    constructor(
        IAVSDirectory _avsDirectory,
        IRewardsCoordinator _rewardsCoordinator,
        IRegistryCoordinator _registryCoordinator,
        IStakeRegistry _stakeRegistry,
        ITriggerXTaskManager _taskManagerContract
    ) ServiceManagerBase(_avsDirectory, _rewardsCoordinator, _registryCoordinator, _stakeRegistry) {
        _disableInitializers();
        taskManagerContract = _taskManagerContract;
    }

    function initialize(
        address initialOwner,
        address rewardsInitiator,
        address _taskManager,
        address _taskValidator,
        address _quorumManager
    ) external initializer {
        __ServiceManagerBase_init(initialOwner, rewardsInitiator);
        __Ownable_init();
        __UUPSUpgradeable_init();

        taskManager = _taskManager;
        taskValidator = _taskValidator;
        quorumManager = _quorumManager;
    }

    function setTaskManager(address _taskManager) external onlyOwner {
        taskManager = _taskManager;
    }

    function setTaskValidator(address _taskValidator) external onlyOwner {
        taskValidator = _taskValidator;
    }

    function setQuorumManager(address _quorumManager) external onlyOwner {
        quorumManager = _quorumManager;
    }

    /// @notice Function that authorizes an upgrade to a new implementation
    /// @dev Required by UUPSUpgradeable
    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[50] private __gap;
}
