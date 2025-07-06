// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.8.25;

struct TaskSubmissionParams {
    string name; 
    uint8 taskDefinitionId;
    uint64 taskResponsePeriod;
    uint64 taskChallengePeriod;
    uint8 thresholdPercentage;
    uint64 statisticalPeriod;
}

struct TaskDefinitionParams {
    uint256 baseRewardFeeForAttesters;
    uint256 baseRewardFeeForPerformer;
    uint256 baseRewardFeeForAggregator;
    uint256 minimumVotingPower;
    uint256[] restrictedOperatorIds;
    uint256 maximumNumberOfAttesters;
}

struct TaskDefinition {
    uint8 taskDefinitionId;
    string name;    // human readable name
    uint256 baseRewardFeeForAttesters;
    uint256 baseRewardFeeForPerformer;
    uint256 baseRewardFeeForAggregator;
    uint256 minimumVotingPower;
    uint256[] restrictedOperatorIds;
    uint256 maximumNumberOfAttesters;
}

struct TaskDefinitions {
    uint8 counter;
    mapping(uint8 => TaskDefinition) taskDefinitions;
}


library TaskDefinitionLibrary {
    event TaskDefinitionCreated(
        uint8 taskDefinitionId,
        string name,
        uint256 baseRewardFeeForAttesters,
        uint256 baseRewardFeeForPerformer,
        uint256 baseRewardFeeForAggregator,
        uint256 minimumVotingPower,
        uint256[] restrictedOperatorIds,
        uint256 maximumNumberOfAttesters
    );

    // keep the params according to the TaskDefinition

    function createNewTaskDefinition(
        TaskDefinitions storage self,
        string memory _name,
        TaskDefinitionParams memory _definitionParams
    ) internal returns (uint8 _id) {
        _id = ++self.counter;
        self.taskDefinitions[_id] = TaskDefinition(
            _id,
            _name,
            _definitionParams.baseRewardFeeForAttesters,
            _definitionParams.baseRewardFeeForPerformer,
            _definitionParams.baseRewardFeeForAggregator,
            _definitionParams.minimumVotingPower,
            _definitionParams.restrictedOperatorIds,
            _definitionParams.maximumNumberOfAttesters
        );
        emit TaskDefinitionCreated(
            _id,
            _name,
            _definitionParams.baseRewardFeeForAttesters,
            _definitionParams.baseRewardFeeForPerformer,
            _definitionParams.baseRewardFeeForAggregator,
            _definitionParams.minimumVotingPower,
            _definitionParams.restrictedOperatorIds,
            _definitionParams.maximumNumberOfAttesters
        );
    }   

    function getTaskDefinition(TaskDefinitions storage self, uint8 _taskDefinitionId)
        internal
        view
        returns (TaskDefinition storage)
    {
        return self.taskDefinitions[_taskDefinitionId];
    }

    function getMinimumVotingPower(TaskDefinitions storage self, uint8 _taskDefinitionId)
        internal
        view
        returns (uint256)
    {
        return self.taskDefinitions[_taskDefinitionId].minimumVotingPower;
    }

    function getRestrictedOperatorIds(TaskDefinitions storage self, uint8 _taskDefinitionId)
        internal
        view
        returns (uint256[] storage)
    {
        return self.taskDefinitions[_taskDefinitionId].restrictedOperatorIds;
    }

    function getMaximumNumberOfAttesters(TaskDefinitions storage self, uint8 _taskDefinitionId)
        internal
        view
        returns (uint256)
    {
        return self.taskDefinitions[_taskDefinitionId].maximumNumberOfAttesters;
    }
  
}