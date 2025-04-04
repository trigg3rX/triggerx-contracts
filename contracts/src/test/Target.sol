// SPDX-License-Identifier: MIT
pragma solidity >=0.8.25;

contract Target {
    address public allowedCaller;

    constructor(address _allowedCaller) {
        allowedCaller = _allowedCaller;
    }

    function setAllowedCaller(address _allowedCaller) external {
        allowedCaller = _allowedCaller;
    }

    // Remove the ETH requirement, just keep it payable
    function increment(uint256 x) external payable returns (uint256) {
        require(msg.sender == allowedCaller, "Unauthorized caller");
        return x + 1;
    }
}