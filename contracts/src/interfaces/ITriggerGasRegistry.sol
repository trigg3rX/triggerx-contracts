// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

interface ITriggerGasRegistry {
    // STRUCTS
    struct TGBalance {
        uint256 ethSpent;
        uint256 TGbalance;
    }

    // EVENTS
    event TGPurchased(address indexed user, uint256 ethAmount, uint256 tgAmount);
    event TGRefunded(address indexed user, uint256 amount);
    event TGBalanceRemoved(address indexed user, uint256 amount, string reason);
    event TGClaimed(address indexed user, uint256 amount);
    event TaskFeeClaimed(address indexed user, uint256 amount);
    event RewardClaimed(address indexed user, uint256 reward);
    event TGTransferred(address indexed user, address indexed keeper, uint256 amount);
    event ETHWithdrawn(address indexed owner, uint256 amount, string reason);
    event TGWithdrawn(address indexed owner, uint256 amount, string reason);

    // FUNCTIONS
    function initialize() external;
    function purchaseTG(uint256 ethAmount) external payable;
    function claimETHForTG(uint256 TGAmount) external;
    function updateTGBalances(address keeper, address user, uint256 TGAmount) external;
    function withdrawETH(uint256 amount, string calldata reason) external;

    // VIEW FUNCTIONS
    function getBalance(address user) external view returns (uint256 ethSpent, uint256 TGbalance);
    function balances(address) external view returns (uint256 ethSpent, uint256 TGbalance);
    function points(address) external view returns (uint256);
} 