// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract Target {
    uint256 public value;
    address public lastCaller;
    bytes public lastData;

    event ValueSet(address indexed caller, uint256 newValue);

    function setValue(uint256 _newValue) payable external {
        value = _newValue;
        lastCaller = msg.sender; // Will be the SpokeContract address
        lastData = msg.data;
        emit ValueSet(msg.sender, _newValue);
    }
}