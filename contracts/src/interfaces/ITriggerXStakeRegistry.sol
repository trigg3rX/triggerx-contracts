// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

interface ITriggerXStakeRegistry {
    struct Stake {
        uint256 amount;
        uint256 TGbalance;
    }

    // Events
    event Staked(address indexed user, uint256 amount);
    event Unstaked(address indexed user, uint256 amount);
    event StakeRemoved(address indexed user, uint256 amount, string reason);
    event TGClaimed(address indexed user, uint256 amount);
    event TaskFeeClaimed(address indexed user, uint256 amount);
    event RewardClaimed(address indexed user, uint256 reward);
    event TGTransferred(address indexed user, address indexed keeper, uint256 amount);
    event ETHWithdrawn(address indexed owner, uint256 amount, string reason);

    function stake(uint256 amount) external payable;
    function getStake(address user) external view returns (uint256 amount, uint256 TGbalance);
    function claimETHForTG(uint256 TGAmount) external;
    function updateTGBalances(address keeper, address user, uint256 TGAmount) external;
    function withdrawETH(uint256 amount, string memory reason) external;
    function stakes(address user) external view returns (uint256 amount, uint256 TGbalance);
    function points(address user) external view returns (uint256);
}