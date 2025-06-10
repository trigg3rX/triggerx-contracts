// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import "@openzeppelin-upgrades/contracts/utils/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";

contract TriggerXStakeRegistry is 
    Initializable,
    ReentrancyGuardUpgradeable,
    UUPSUpgradeable,
    OwnableUpgradeable 
{
    struct Stake {
        uint256 amount;
        uint256 TGbalance;
    }

    mapping(address => Stake) public stakes;

    // Add a mapping to track points for each user
    mapping(address => uint256) public points;

    // Events
    event Staked(address indexed user, uint256 amount);
    event Unstaked(address indexed user, uint256 amount);
    event StakeRemoved(address indexed user, uint256 amount, string reason);
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
     * @notice Allows a user to stake ETH
     * @param amount The amount of ETH to stake
     */
    function stake(uint256 amount) external payable nonReentrant {
        require(amount > 0, "Cannot stake 0 ETH");
        require(msg.value == amount, "Sent ETH must match amount");
        
        Stake storage userStake = stakes[msg.sender];
        userStake.amount += amount;
        
        // Calculate and add TG rewards (0.001 ETH per TG)
        uint256 TGReward = (amount * 1) * 1000; // This gives 0.001 ratio
        userStake.TGbalance += TGReward;

        emit Staked(msg.sender, amount);
    }

    /**
     * @notice Returns the stake information for a given user
     * @param user The address of the user
     * @return amount The staked amount
     * @return TGbalance The TG balance
     */
    function getStake(address user) external view returns (uint256 amount, uint256 TGbalance) {
        Stake memory userStake = stakes[user];
        return (userStake.amount, userStake.TGbalance);
    }

    /**
     * @notice Allows users to claim ETH based on their TG balance
     * @param TGAmount The amount of TG to convert to ETH
     */
    function claimETHForTG(uint256 TGAmount) external nonReentrant {
        Stake storage userStake = stakes[msg.sender];
        require(userStake.TGbalance >= TGAmount, "Insufficient TG balance");

        uint256 ethToReturn = TGAmount / 1000; // Convert TG to ETH (1 ETH = 1000 TG)
        userStake.TGbalance -= TGAmount;

        (bool success, ) = msg.sender.call{value: ethToReturn}("");
        require(success, "ETH transfer failed");

        emit TGClaimed(msg.sender, TGAmount);
    }

    /**
     * @notice Updates TG balances when a task is completed
     * @param keeper The address of the keeper who completed the task
     * @param user The address of the user who created the task
     * @param TGAmount The amount of TG to transfer
     */
    function updateTGBalances(address keeper, address user, uint256 TGAmount) external onlyOwner nonReentrant {
        Stake storage userStake = stakes[user];
        Stake storage keeperStake = stakes[keeper];
        
        require(userStake.TGbalance >= TGAmount, "Insufficient user TG balance");
        
        // Deduct TG from user
        userStake.TGbalance -= TGAmount;
        
        // Add TG to keeper
        keeperStake.TGbalance += TGAmount;
        
        emit TGTransferred(user, keeper, TGAmount);
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