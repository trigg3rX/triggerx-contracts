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
    
    // Mock unpackJobId - returns current chain ID and jobId as counter
    function unpackJobId(uint256 jobId) external view returns (uint256 chainId, uint256 timestamp, uint256 jobCounter) {
        // For testing, return the current chain ID so tests pass the chain check
        // and return the jobId as the counter
        return (block.chainid, block.timestamp, jobId);
    }
}