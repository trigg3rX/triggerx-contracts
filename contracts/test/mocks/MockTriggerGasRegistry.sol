// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

contract MockTriggerGasRegistry {
    mapping(address => uint256) public balances;
    
    function setBalance(address user, uint256 balance) external {
        balances[user] = balance;
    }
    
    function deductTGBalance(address user, uint256 tgAmount) external {
        require(balances[user] >= tgAmount, "Insufficient TG balance");
        balances[user] -= tgAmount;
    }
    
    function getBalance(address user) external view returns (uint256, uint256) {
        return (0, balances[user]); // Return (ethSpent, tgBalance) to match real contract
    }
}
