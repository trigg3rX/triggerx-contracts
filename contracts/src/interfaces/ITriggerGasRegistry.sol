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
    function owner() external view returns (address);
    function getBalance(address user) external view returns (uint256 ethSpent, uint256 TGbalance);
    function balances(address) external view returns (uint256 ethSpent, uint256 TGbalance);
    function points(address) external view returns (uint256);
    
    // MIGRATION FUNCTIONS
    function migrateUserBalance(address user, uint256 ethSpent, uint256 tgBalance) external;
    function migrateUserPoints(address user, uint256 userPoints) external;
    function batchMigrateUsers(
        address[] calldata users,
        uint256[] calldata ethSpentAmounts,
        uint256[] calldata tgBalances,
        uint256[] calldata userPoints
    ) external;
} 