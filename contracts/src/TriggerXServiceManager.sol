// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@eigenlayer-middleware/ServiceManagerBase.sol";
import {EnumerableSet} from "@openzeppelin-contracts/contracts/utils/structs/EnumerableSet.sol";
import "./ITriggerXTaskManager.sol";

contract TriggerXServiceManager is ServiceManagerBase {
    using EnumerableSet for EnumerableSet.AddressSet;

    ITriggerXTaskManager public immutable triggerXTaskManager;
    EnumerableSet.AddressSet internal _operators;
    mapping(address => bool) public operatorBlacklist;
    uint8 public quorumThresholdPercentage;

    event OperatorAdded(address indexed operator);
    event OperatorBlacklistStatusChanged(address indexed operator, bool blacklisted);
    event QuorumThresholdPercentageChanged(uint8 newThresholdPercentage);

    modifier onlyValidOperator() {
        if (!_operators.contains(msg.sender)) {
            "not onlyValidOperator: Only valid operator can call this function";
        }
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

    function registerOperatorToAVS(
        address operator,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) public override onlyRegistryCoordinator {
        _avsDirectory.registerOperatorToAVS(operator, operatorSignature);
        _operators.add(operator);
        emit OperatorAdded(operator);
    }

    function whitelistOperator(address operator) external onlyOwner {
        if (operator == address(0))
            "whitelistOperator: Operator address cannot be zero";
        operatorBlacklist[operator] = false;
        emit OperatorBlacklistStatusChanged(operator, false);
    }

    function blacklistOperator(address operator) external onlyOwner {
        if (operator == address(0))
            "blacklistOperator: Operator address cannot be zero";
        operatorBlacklist[operator] = true;    
        emit OperatorBlacklistStatusChanged(operator, true);
    }

    function isOperatorBlacklisted(address operator) external view returns (bool) {
        return operatorBlacklist[operator];
    }

    function updateQuorumThresholdPercentage(uint8 thresholdPercentage) external onlyOwner {
        if (thresholdPercentage > 100)
            "updateQuorumThresholdPercentage: Threshold percentage cannot be greater than 100";
        quorumThresholdPercentage = thresholdPercentage;
        emit QuorumThresholdPercentageChanged(thresholdPercentage);
    }
}