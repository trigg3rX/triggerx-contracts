// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import "@openzeppelin-upgrades/contracts/utils/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";

contract TriggerGasRegistry is 
    Initializable,
    ReentrancyGuardUpgradeable,
    UUPSUpgradeable,
    OwnableUpgradeable 
{
    struct TGBalance {
        uint256 ethSpent;
        uint256 TGbalance;
    }

    mapping(address => TGBalance) public balances;

    // Add a mapping to track points for each user
    mapping(address => uint256) public points;

    // Events
    event TGPurchased(address indexed user, uint256 ethAmount, uint256 tgAmount);
    event TGRefunded(address indexed user, uint256 amount);
    event TGBalanceRemoved(address indexed user, uint256 amount, string reason);
    event TGClaimed(address indexed user, uint256 amount);
    event TaskFeeClaimed(address indexed user, uint256 amount);
    event RewardClaimed(address indexed user, uint256 reward);
    event TGTransferred(address indexed user, address indexed keeper, uint256 amount);
    event ETHWithdrawn(address indexed owner, uint256 amount, string reason);
    
    constructor() {
        _disableInitializers();
    }
    
    function initialize() public initializer {
        __ReentrancyGuard_init();
        __Ownable_init(msg.sender);
        __UUPSUpgradeable_init();
    }

    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}

    /**
     * @notice Allows a user to purchase TG tokens with ETH
     * @param ethAmount The amount of ETH to spend
     */
    function purchaseTG(uint256 ethAmount) external payable nonReentrant {
        require(ethAmount > 0, "Cannot spend 0 ETH");
        require(msg.value == ethAmount, "Sent ETH must match amount");
        
        TGBalance storage userBalance = balances[msg.sender];
        userBalance.ethSpent += ethAmount;
        
        // Calculate TG tokens (0.001 ETH per TG)
        uint256 tgAmount = ethAmount * 1000; // This gives 0.001 ratio
        userBalance.TGbalance += tgAmount;

        emit TGPurchased(msg.sender, ethAmount, tgAmount);
    }

    /**
     * @notice Returns the TG balance information for a given user
     * @param user The address of the user
     * @return ethSpent The amount of ETH spent
     * @return tgBalance The TG balance
     */
    function getBalance(address user) external view returns (uint256 ethSpent, uint256 tgBalance) {
        TGBalance memory userBalance = balances[user];
        return (userBalance.ethSpent, userBalance.TGbalance);
    }

    /**
     * @notice Allows users to claim ETH based on their TG balance
     * @param tgAmount The amount of TG to convert to ETH
     */
    function claimETHForTG(uint256 tgAmount) external nonReentrant {
        TGBalance storage userBalance = balances[msg.sender];
        require(userBalance.TGbalance >= tgAmount, "Insufficient TG balance");

        uint256 ethToReturn = tgAmount / 1000; // Convert TG to ETH (1 ETH = 1000 TG)
        userBalance.TGbalance -= tgAmount;

        (bool success, ) = msg.sender.call{value: ethToReturn}("");
        require(success, "ETH transfer failed");

        emit TGClaimed(msg.sender, tgAmount);
    }

    /**
     * @notice Updates TG balances when a task is completed
     * @param keeper The address of the keeper who completed the task
     * @param user The address of the user who created the task
     * @param tgAmount The amount of TG to transfer
     */
    function updateTGBalances(address keeper, address user, uint256 tgAmount) external onlyOwner nonReentrant {
        TGBalance storage userBalance = balances[user];
        TGBalance storage keeperBalance = balances[keeper];
        
        require(userBalance.TGbalance >= tgAmount, "Insufficient user TG balance");
        
        // Deduct TG from user
        userBalance.TGbalance -= tgAmount;
        
        // Add TG to keeper
        keeperBalance.TGbalance += tgAmount;
        
        emit TGTransferred(user, keeper, tgAmount);
    }

    /**
     * @notice Allows the owner to withdraw ETH from the contract
     * @param amount The amount of ETH to withdraw
     * @param reason The reason for withdrawal
     */
    function withdrawETH(uint256 amount, string memory reason) external onlyOwner nonReentrant {
        require(amount > 0, "Cannot withdraw 0 ETH");
        require(amount <= address(this).balance, "Insufficient contract balance");

        (bool success, ) = owner().call{value: amount}("");
        require(success, "ETH transfer failed");

        emit ETHWithdrawn(owner(), amount, reason);
    }
} 