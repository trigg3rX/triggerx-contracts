// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

contract MockTaskExecutionHub {
    address public owner;
    mapping(address => bool) public isOperator;
    
    constructor() {
        owner = msg.sender;
    }
    
    function registerOperator(address _operator) external {
        isOperator[_operator] = true;
    }
    
    function unregisterOperator(address _operator) external {
        isOperator[_operator] = false;
    }
} 