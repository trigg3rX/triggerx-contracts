// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

interface ITriggerXStakeRegistry {
    // STRUCTS
    struct Stake {
        uint256 amount;
        uint256 TGbalance;
    }

    // EVENTS
    event Staked(address indexed user, uint256 amount);
    // event Unstaked(address indexed user, uint256 amount);
    // event StakeRemoved(address indexed user, uint256 amount, string reason);
    event TGClaimed(address indexed user, uint256 amount);
    event TaskFeeClaimed(address indexed user, uint256 amount);
    // event RewardClaimed(address indexed user, uint256 reward);
    event TGTransferred(address indexed user, address indexed keeper, uint256 amount);

    // FUNCTIONS
    function initialize() external;
    function stake(uint256 amount) external payable;
    function unstake(uint256 amount) external;
    function removeStake(address user, uint256 amount, string calldata reason) external;
    function claimETHForTG(uint256 TGAmount) external;
    function getTaskFee(uint256 amount) external;
    function getReward(uint256 claimedTG) external;
    function updateTGBalances(address keeper, address user, uint256 TGAmount) external;

    // VIEW FUNCTIONS
    function getStake(address user) external view returns (uint256 amount, uint256 TGbalance);
    function stakes(address) external view returns (uint256 amount, uint256 TGbalance);
    function points(address) external view returns (uint256);
    // function pool() external view returns (uint256);
} 