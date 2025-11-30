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
    mapping(address => uint256) public balances;

    address public operatorRole;
    uint256 public totalDeductedBalance;

    // Events
    event ETHDeposited(address indexed user, uint256 ethAmount);
    event ETHWithdrawn(address indexed owner, uint256 amount, string reason);
    event ETHBalanceDeducted(address indexed user, uint256 amount);

    // Modifiers
    modifier onlyOperator() {
        require(msg.sender == operatorRole, "Only operator can call this function");
        _;
    }

    constructor() {
        _disableInitializers();
    }
    
    function initialize(address initialOwner, address _operator) public initializer {
        require(initialOwner != address(0), "Initial owner cannot be 0 address");
        
        __ReentrancyGuard_init();
        __Ownable_init(initialOwner);
        __UUPSUpgradeable_init();

        // set the operator role
        operatorRole = _operator;
    }

    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}

    /**
     * @notice Allows a user to deposit ETH
     * @param ethAmount The amount of ETH to deposit
     */
    function depositETH(uint256 ethAmount) external payable nonReentrant {
        require(ethAmount > 0, "Cannot deposit 0 ETH");
        require(msg.value == ethAmount, "Sent ETH must match amount");
        
        balances[msg.sender] += ethAmount;

        emit ETHDeposited(msg.sender, ethAmount);
    }

    /**
     * @notice Returns the ETH balance information for a given user
     * @param user The address of the user
     * @return ethAmount The amount of ETH available
     */
    function getBalance(address user) external view returns (uint256 ethAmount) {
        return balances[user];
    }

    /**
     * @notice Allows users to withdraw their ETH balance
     * @param ethAmount The amount of ETH to withdraw
     */
    function withdrawETHBalance(uint256 ethAmount) external nonReentrant {
        require(ethAmount > 0, "Cannot withdraw 0 ETH");
        require(balances[msg.sender] >= ethAmount, "Insufficient ETH balance");

        balances[msg.sender] -= ethAmount;

        // slither-disable-next-line low-level-calls
        (bool success, ) = msg.sender.call{value: ethAmount}("");
        require(success, "ETH transfer failed");
    }

    /**
     * @notice Allows the operator to deduct ETH balance from a user
     * @param user The address of the user
     * @param ethAmount The amount of ETH to deduct
     */
    function deductETHBalance(address user, uint256 ethAmount) external onlyOperator {
        require(balances[user] >= ethAmount, "Insufficient ETH balance");
        if(ethAmount > 0) {
            balances[user] -= ethAmount;
            totalDeductedBalance += ethAmount;
            emit ETHBalanceDeducted(user, ethAmount);
        }
    }

    /**
     * @notice Allows the owner to withdraw ETH from the contract (only deducted balance)
     * @param amount The amount of ETH to withdraw
     * @param reason The reason for withdrawal
     */
    // slither-disable-next-line reentrancy-events
    function withdrawETH(uint256 amount, string memory reason) external onlyOwner nonReentrant {
        require(amount > 0, "Cannot withdraw 0 ETH");
        require(amount <= totalDeductedBalance, "Cannot withdraw more than deducted balance");
        require(amount <= address(this).balance, "Insufficient contract balance");

        totalDeductedBalance -= amount;

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

    /**
     * @notice Batch migration function for efficient migration of multiple users
     * @param users Array of user addresses
     * @param ethAmounts Array of ETH amounts
     */
    function batchMigrateUsers(
        address[] calldata users,
        uint256[] calldata ethAmounts
    ) external payable onlyOwner {
        require(users.length == ethAmounts.length, "Length mismatch");

        uint256 totalEth = 0;
        for (uint i = 0; i < users.length; i++) {
            balances[users[i]] = ethAmounts[i];
            totalEth += ethAmounts[i];
            emit ETHDeposited(users[i], ethAmounts[i]);
        }
        require(msg.value >= totalEth, "Sent ETH must be greater than or equal to total ETH amount");
    }
} 