// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract CounterContract {
    address public owner;
    address public keeper;
    address public targetAddress;
    uint256 public counter;

    event CounterIncremented(uint256 newValue);
    event TargetAddressChanged(address oldTarget, address newTarget);
    event KeeperChanged(address oldKeeper, address newKeeper);

    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call this function");
        _;
    }

    modifier onlyKeeper() {
        require(msg.sender == keeper, "Only keeper can call this function");
        _;
    }

    constructor() {
        owner = msg.sender;
        keeper = 0x68605feB94a8FeBe5e1fBEF0A9D3fE6e80cEC126; // Initially set keeper as the deployer
        counter = 0;
    }

    function incrementCounter() external onlyKeeper {
        counter += 1;
        emit CounterIncremented(counter);
    }

    function getCounter() external view returns (uint256) {
        return counter;
    }

    function setTargetAddress(address _newTarget) external onlyOwner {
        require(_newTarget != address(0), "Invalid address");
        address oldTarget = targetAddress;
        targetAddress = _newTarget;
        emit TargetAddressChanged(oldTarget, _newTarget);
    }

    function setKeeper(address _newKeeper) external onlyOwner {
        require(_newKeeper != address(0), "Invalid keeper address");
        address oldKeeper = keeper;
        keeper = _newKeeper;
        emit KeeperChanged(oldKeeper, _newKeeper);
    }
}
