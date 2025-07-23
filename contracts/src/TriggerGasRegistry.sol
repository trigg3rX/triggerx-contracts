// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

/**
 * 
 *   ╔╦╗╦═╗╦╔═╗╔═╗╔═╗╦═╗═╗ ╦
 *    ║ ╠╦╝║║ ╦║ ╦║╣ ╠╦╝╔╩╦╝ 
 *    ╩ ╩╚═╩╚═╝╚═╝╚═╝╩╚═╩ ╚═
 * 
 */

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

    // slither-disable-next-line naming-convention
    uint256 public TG_PER_ETH;

    address public operatorRole;

    // Events
    event TGPurchased(address indexed user, uint256 ethAmount, uint256 tgAmount);
    event TGRefunded(address indexed user, uint256 amount);
    event TGBalanceRemoved(address indexed user, uint256 amount, string reason);
    event TGClaimed(address indexed user, uint256 amount);
    event TaskFeeClaimed(address indexed user, uint256 amount);
    event RewardClaimed(address indexed user, uint256 reward);
    event TGTransferred(address indexed user, address indexed keeper, uint256 amount);
    event ETHWithdrawn(address indexed owner, uint256 amount, string reason);
    event TGBalanceDeducted(address indexed user, uint256 amount);
    event TGPerETHUpdated(uint256 tgPerEth);

    // Modifiers
    modifier onlyOperator() {
        require(msg.sender == operatorRole, "Only operator can call this function");
        _;
    }

    constructor() {
        _disableInitializers();
    }
    
    function initialize(address initialOwner, address _operator, uint256 _tgPerEth) public initializer {
        require(initialOwner != address(0), "Initial owner cannot be 0 address");
        
        __ReentrancyGuard_init();
        __Ownable_init(initialOwner);
        __UUPSUpgradeable_init();

        // set the operator role
        TG_PER_ETH = _tgPerEth;
        operatorRole = _operator;
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
        
        uint256 tgAmount = ethAmount * TG_PER_ETH;
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

        uint256 ethToReturn = tgAmount / TG_PER_ETH;
        userBalance.TGbalance -= tgAmount;

        // slither-disable-next-line low-level-calls
        (bool success, ) = msg.sender.call{value: ethToReturn}("");
        require(success, "ETH transfer failed");

        emit TGClaimed(msg.sender, tgAmount);
    }

    /**
     * @notice Allows the operator to deduct TG balance from a user
     * @param user The address of the user
     * @param tgAmount The amount of TG to deduct
     */
    function deductTGBalance(address user, uint256 tgAmount) external onlyOperator {
        require(balances[user].TGbalance >= tgAmount, "Insufficient TG balance");
        if(tgAmount > 0) {
            balances[user].TGbalance -= tgAmount;
            emit TGBalanceDeducted(user, tgAmount);
        }
    }

    /**
     * @notice Allows the owner to withdraw ETH from the contract
     * @param amount The amount of ETH to withdraw
     * @param reason The reason for withdrawal
     */
    // slither-disable-next-line reentrancy-events
    function withdrawETH(uint256 amount, string memory reason) external onlyOwner nonReentrant {
        require(amount > 0, "Cannot withdraw 0 ETH");
        require(amount <= address(this).balance, "Insufficient contract balance");

        // slither-disable-next-line low-level-calls
        (bool success, ) = owner().call{value: amount}("");
        require(success, "ETH transfer failed");

        emit ETHWithdrawn(owner(), amount, reason);
    }

    // slither-disable-next-line missing-events-access-control
    function setOperator(address _operatorRole) external onlyOwner {
        require(_operatorRole != address(0), "Operator cannot be 0 address");
        operatorRole = _operatorRole;
    }

    function setTGPerETH(uint256 _tgPerEth) external onlyOwner {
        TG_PER_ETH = _tgPerEth;
        emit TGPerETHUpdated(_tgPerEth);
    }
} 