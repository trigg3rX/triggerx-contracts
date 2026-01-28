// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

contract MockJobRegistry {
    mapping(uint256 => address) public jobOwners;
    mapping(uint256 => bool) public jobActive;
    mapping(uint256 => bool) private _jobActiveSet;

    function setJobOwner(uint256 jobId, address owner) external {
        jobOwners[jobId] = owner;
        // Default to active when setting owner
        if (!_jobActiveSet[jobId]) {
            jobActive[jobId] = true;
        }
    }

    function setJobActive(uint256 jobId, bool active) external {
        jobActive[jobId] = active;
        _jobActiveSet[jobId] = true;
    }

    function getJobOwner(uint256 jobId) external view returns (address) {
        return jobOwners[jobId];
    }

    function isJobActive(uint256 jobId) external view returns (bool) {
        // Default to true if owner is set and not explicitly deactivated
        if (jobOwners[jobId] == address(0)) {
            return false;
        }
        if (_jobActiveSet[jobId]) {
            return jobActive[jobId];
        }
        return true; // Default active if owner exists
    }

    // Mock unpackJobId - returns current chain ID and jobId as counter
    function unpackJobId(
        uint256 jobId
    ) external view returns (uint256 chainId, uint256 jobCounter) {
        // For testing, return the current chain ID so tests pass the chain check
        // and return the jobId as the counter
        return (block.chainid, jobId);
    }
}
