// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@eigenlayer-middleware/ServiceManagerBase.sol";
import "./ITriggerXTaskManager.sol";

contract TriggerXServiceManager is ServiceManagerBase {

    ITriggerXTaskManager public immutable triggerXTaskManager;

    modifier onlyTriggerXTaskManager() {
        require(
            msg.sender == address(triggerXTaskManager),
            "onlyTriggerXTaskManager: not from riggerX task manager"
        );
        _;
    }

    constructor(
        IAVSDirectory __avsDirectory,
        IRegistryCoordinator __registryCoordinator,
        IStakeRegistry __stakeRegistry,
        ITriggerXTaskManager __triggerXTaskManager
    ) ServiceManagerBase(
        __avsDirectory, 
        __registryCoordinator, 
        __stakeRegistry
    ) 
    {
        triggerXTaskManager = __triggerXTaskManager;
    }

    function freezeOperator(
        address operatorAddr
    ) external onlyTriggerXTaskManager {
        // slasher.freezeOperator(operatorAddr);
    }
    
}