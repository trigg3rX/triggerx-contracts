// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@eigenlayer-middleware/src/ServiceManagerBase.sol";
import "./ItxTaskManager.sol";

contract txServiceManager is ServiceManagerBase {

    ItxTaskManager public immutable txTaskManager;

    modifier onlyTxTaskManager() {
        require(
            msg.sender == address(txTaskManager),
            "onlytxTaskManager: not from triggerX task manager"
        );
        _;
    }

    constructor(
        IAVSDirectory __avsDirectory,
        IRegistryCoordinator __registryCoordinator,
        IStakeRegistry __stakeRegistry,
        ItxTaskManager __txTaskManager
    ) ServiceManagerBase(
        __avsDirectory, 
        __registryCoordinator, 
        __stakeRegistry
    ) 
    {
        txTaskManager = __txTaskManager;
    }

    // functions for rewards / slashing
    
}