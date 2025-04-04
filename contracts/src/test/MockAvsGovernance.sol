// SPDX-License-Identifier: MIT
pragma solidity >=0.8.25;

contract MockAvsGovernance {
    mapping(address => bool) public isOperatorRegistered;

    function setOperatorRegistered(address operator, bool status) external {
        isOperatorRegistered[operator] = status;
    }
}