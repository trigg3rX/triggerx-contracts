// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

interface ITriggerXStakeRegistry {
    // STRUCTS
    struct Stake {
        uint256 amount;
        bool exists;
    }

    // EVENTS
    event Staked(address indexed user, uint256 amount);
    event Unstaked(address indexed user, uint256 amount);
    event StakeRemoved(address indexed user, uint256 amount, string reason);
    event TGClaimed(address indexed user, uint256 amount);
    event TaskFeeClaimed(address indexed user, uint256 amount);
    event RewardClaimed(address indexed user, uint256 reward);

    // FUNCTIONS
    function initialize() external;
    function stake(uint256 amount) external payable;
    function unstake(uint256 amount) external;
    function removeStake(address user, uint256 amount, string calldata reason) external;
    function claimTG() external;
    function getTaskFee(uint256 amount) external;
    function getReward(uint256 claimedTG) external;

    // VIEW FUNCTIONS
    function getStake(address user) external view returns (uint256 amount, bool exists);
    function stakes(address) external view returns (uint256 amount, bool exists);
    function points(address) external view returns (uint256);
    function pool() external view returns (uint256);
    function TG_AMOUNT() external view returns (uint256);
} 