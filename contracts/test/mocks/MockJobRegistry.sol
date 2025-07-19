// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

contract MockJobRegistry {
    mapping(uint256 => address) public jobOwners;
    
    function setJobOwner(uint256 jobId, address owner) external {
        jobOwners[jobId] = owner;
    }
    
    function getJobOwner(uint256 jobId) external view returns (address) {
        return jobOwners[jobId];
    }
}