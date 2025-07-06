// SPDX-License-Identifier: MIT
pragma solidity 0.8.30;

import "forge-std/Test.sol";
import "../src/TriggerXAvs.sol";

/// Mock that stands in for the AVS-Manager precompile (same as earlier tests).
contract MockAVSManager is avs.IAVSManager {
    uint64 public taskCounter;

    function registerAVS(avs.AVSParams calldata) external pure returns (bool) { return true; }
    function updateAVS(avs.AVSParams calldata) external pure returns (bool) { return true; }
    function deregisterAVS(address, string calldata) external pure returns (bool) { return true; }
    function registerOperatorToAVS(address) external pure returns (bool) { return true; }
    function deregisterOperatorFromAVS(address) external pure returns (bool) { return true; }
    function createTask(address, string calldata, bytes calldata, uint64, uint64, uint8, uint64) external returns (uint64) { return ++taskCounter; }
    function registerBLSPublicKey(address, address, bytes calldata, bytes calldata) external pure returns (bool) { return true; }
    function operatorSubmitTask(address, uint64, bytes calldata, bytes calldata, address, uint8) external pure returns (bool) { return true; }
    function challenge(address, uint64, address, uint8, bool, address[] calldata, address[] calldata) external pure returns (bool) { return true; }

    // unused views
    function getRegisteredPubkey(address, address) external pure returns (bytes memory) { return ""; }
    function getOptInOperators(address) external pure returns (address[] memory operators) {}
    function getAVSUSDValue(address) external pure returns (uint256) { return 0; }
    function getOperatorOptedUSDValue(address, address) external pure returns (uint256) { return 0; }
    function getAVSEpochIdentifier(address) external pure returns (string memory) { return "hour"; }
    function getTaskInfo(address, uint64) external pure returns (avs.TaskInfo memory) { revert(); }
    function isOperator(address) external pure returns (bool) { return true; }
    function getCurrentEpoch(string calldata) external pure returns (int64) { return 0; }
    function getChallengeInfo(address, uint64) external pure returns (address) { return address(0); }
    function getOperatorTaskResponse(address, address, uint64) external pure returns (avs.TaskResultInfo memory) { revert(); }
    function getOperatorTaskResponseList(address, uint64) external pure returns (avs.OperatorResInfo[] memory) { revert(); }
}

contract TriggerXAvsOwnerTests is Test {
    TriggerXAvs public avsProxy;
    MockAVSManager public mock;

    address owner = address(this);
    address nonOwner = address(0xBEEF);

    function setUp() public {
        // Deploy and etch mock precompile
        mock = new MockAVSManager();
        vm.etch(avs.AVSMANAGER_PRECOMPILE_ADDRESS, address(mock).code);

        avsProxy = new TriggerXAvs();
        avsProxy.initialize(owner);
    }

    // --------------------------------------------------------
    // RewardManager & Slasher setters
    // --------------------------------------------------------

    function testOnlyOwnerSetRewardManager() public {
        address newRM = address(0xCAFE);
        // non-owner should revert
        vm.prank(nonOwner);
        vm.expectRevert("Ownable: caller is not the owner");
        avsProxy.setRewardManager(newRM);

        // owner succeeds & event emitted
        vm.expectEmit(true, false, false, false);
        emit TriggerXAvs.RewardManagerUpdated(newRM);
        avsProxy.setRewardManager(newRM);
        assertEq(avsProxy.rewardManager(), newRM);
    }

    function testOnlyOwnerSetSlasher() public {
        address newS = address(0xDEAD);
        vm.prank(nonOwner);
        vm.expectRevert("Ownable: caller is not the owner");
        avsProxy.setSlasher(newS);

        vm.expectEmit(true, false, false, false);
        emit TriggerXAvs.SlasherUpdated(newS);
        avsProxy.setSlasher(newS);
        assertEq(avsProxy.slasher(), newS);
    }

    function _defaultParams() internal view returns (TaskDefinitionParams memory p) {
        p.baseRewardFeeForAttesters = 1 ether;
        p.baseRewardFeeForPerformer = 0;
        p.baseRewardFeeForAggregator = 0;
        p.minimumVotingPower = 0;
        p.restrictedOperatorIds = new uint256[](0);
        p.maximumNumberOfAttesters = 100;
    }

    function testCreateTaskDefinitionOnlyOwner() public {
        TaskDefinitionParams memory params = _defaultParams();

        vm.prank(nonOwner);
        vm.expectRevert("Ownable: caller is not the owner");
        avsProxy.createTaskDefinition("foo", params);

        vm.expectEmit(true, false, false, true);
        emit TriggerXAvs.TaskDefinitionCreated(1, "foo");
        uint8 id = avsProxy.createTaskDefinition("foo", params);
        assertEq(id, 1);

        TaskDefinition memory stored = avsProxy.getTaskDefinition(id);
        assertEq(stored.name, "foo");
        assertEq(stored.maximumNumberOfAttesters, 100);
    }

    // --------------------------------------------------------
    // createTask OnlyOwner guard
    // --------------------------------------------------------

    function testCreateTaskAccessControl() public {
        TaskDefinition memory def = TaskDefinition({
            taskDefinitionId: 1,
            name: "task",
            baseRewardFeeForAttesters: 1 ether,
            baseRewardFeeForPerformer: 0,
            baseRewardFeeForAggregator: 0,
            minimumVotingPower: 0,
            restrictedOperatorIds: new uint256[](0),
            maximumNumberOfAttesters: 100
        });

        vm.prank(nonOwner);
        vm.expectRevert("Ownable: caller is not the owner");
        TriggerXAvs(address(avsProxy)).createTask("task", 1, 5, 3, 60, 6);
    }
} 