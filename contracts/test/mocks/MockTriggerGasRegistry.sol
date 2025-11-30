// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

contract MockTriggerGasRegistry {
    mapping(address => uint256) public balances;
    
    function setBalance(address user, uint256 balance) external {
        balances[user] = balance;
    }
    
    function deductETHBalance(address user, uint256 ethAmount) external {
        require(balances[user] >= ethAmount, "Insufficient ETH balance");
        balances[user] -= ethAmount;
    }
    
    function getBalance(address user) external view returns (uint256) {
        return balances[user]; // Return ethAmount
    }
}
