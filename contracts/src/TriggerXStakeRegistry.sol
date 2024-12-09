// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";
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
        bool exists;
    }

    mapping(address => Stake) public stakes;

    // Events
    event Staked(address indexed user, uint256 amount);
    event Unstaked(address indexed user, uint256 amount);
    event StakeRemoved(address indexed user, uint256 amount, string reason);

    constructor() {
        _disableInitializers();
    }
    
    function initialize() public initializer {
        __ReentrancyGuard_init();
        __Ownable_init();
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
        if (userStake.exists) {
            userStake.amount += amount;
        } else {
            stakes[msg.sender] = Stake({
                amount: amount,
                exists: true
            });
        }

        emit Staked(msg.sender, amount);
    }

    /**
     * @notice Allows a user to unstake their ETH
     * @param amount The amount of ETH to unstake
     */
    function unstake(uint256 amount) external nonReentrant {
        Stake storage userStake = stakes[msg.sender];
        require(userStake.exists, "No stake found");
        require(amount <= userStake.amount, "Insufficient stake");

        userStake.amount -= amount;
        
        if (userStake.amount == 0) {
            userStake.exists = false;
        }

        (bool success, ) = msg.sender.call{value: amount}("");
        require(success, "ETH transfer failed");

        emit Unstaked(msg.sender, amount);
    }

    /**
     * @notice Returns the stake information for a given user
     * @param user The address of the user
     * @return amount The staked amount
     * @return exists Whether the stake exists
     */
    function getStake(address user) external view returns (uint256 amount, bool exists) {
        Stake memory userStake = stakes[user];
        return (userStake.amount, userStake.exists);
    }

    /**
     * @notice Allows owner to remove stakes from a user (e.g., for task execution)
     * @param user The address of the user whose stake to remove
     * @param amount The amount of ETH to remove
     * @param reason The reason for removing the stake
     */
    function removeStake(address user, uint256 amount, string calldata reason) external onlyOwner nonReentrant {
        Stake storage userStake = stakes[user];
        require(userStake.exists, "No stake found");
        require(amount <= userStake.amount, "Insufficient stake");

        userStake.amount -= amount;
        
        if (userStake.amount == 0) {
            userStake.exists = false;
        }

        (bool success, ) = owner().call{value: amount}("");
        require(success, "ETH transfer failed");

        emit StakeRemoved(user, amount, reason);
    }
}
