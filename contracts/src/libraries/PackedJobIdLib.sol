// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

/**
 * @title PackedJobIdLib
 * @dev Library for packing and unpacking chainId and jobCounter into a single uint256
 * @notice Uses 32 bits for chainId (high bits) and 224 bits for jobCounter (low bits)
 */
library PackedJobIdLib {
    // ----- Constants for bit allocation -----
    uint256 private constant CHAIN_ID_BITS = 32;
    uint256 private constant JOB_COUNTER_BITS = 224;
    uint256 private constant JOB_COUNTER_MASK = (1 << JOB_COUNTER_BITS) - 1;
    
    // Maximum values for validation
    uint256 private constant MAX_CHAIN_ID = (1 << CHAIN_ID_BITS) - 1; // 4,294,967,295
    uint256 private constant MAX_JOB_COUNTER = JOB_COUNTER_MASK; // 2^224 - 1

    // Custom errors
    error ChainIdOverflow(uint256 chainId, uint256 maxChainId);
    error JobCounterOverflow(uint256 jobCounter, uint256 maxJobCounter);

    /**
     * @dev Pack chainId and jobId into a single uint256
     * @param chainId The chain ID (must fit in 32 bits)
     * @param jobCounter The job counter (must fit in 224 bits)
     * @return packed The packed value with chainId in high bits and jobCounter in low bits
     */
    function pack(uint256 chainId, uint256 jobCounter) internal pure returns (uint256 packed) {
        // Validate inputs
        if (chainId > MAX_CHAIN_ID) {
            revert ChainIdOverflow(chainId, MAX_CHAIN_ID);
        }
        if (jobCounter > MAX_JOB_COUNTER) {
            revert JobCounterOverflow(jobCounter, MAX_JOB_COUNTER);
        }
        
        // Pack: chainId in high 32 bits, jobId in low 224 bits
        packed = (chainId << JOB_COUNTER_BITS) | jobCounter;
    }

    /**
     * @dev Unpack a packed value into chainId and jobCounter
     * @param packed The packed value to unpack
     * @return chainId The extracted chain ID
     * @return jobCounter The extracted job counter
     */
    function unpack(uint256 packed) internal pure returns (uint256 chainId, uint256 jobCounter) {
        jobCounter = packed & JOB_COUNTER_MASK;
        chainId = packed >> JOB_COUNTER_BITS;
    }

    /**
     * @dev Extract only the chainId from a packed value
     * @param packed The packed value
     * @return chainId The extracted chain ID
     */
    function getChainId(uint256 packed) internal pure returns (uint256 chainId) {
        chainId = packed >> JOB_COUNTER_BITS;
    }

    /**
     * @dev Extract only the jobCounter from a packed value
     * @param packed The packed value
     * @return jobCounter The extracted job counter
     */
    function getJobCounter(uint256 packed) internal pure returns (uint256 jobCounter) {
        jobCounter = packed & JOB_COUNTER_MASK;
    }

    /**
     * @dev Get the maximum allowed chain ID
     * @return maxChainId The maximum chain ID (2^32 - 1)
     */
    function getMaxChainId() internal pure returns (uint256 maxChainId) {
        maxChainId = MAX_CHAIN_ID;
    }

    /**
     * @dev Get the maximum allowed job counter
     * @return maxJobCounter The maximum job counter (2^224 - 1)
     */
    function getMaxJobCounter() internal pure returns (uint256 maxJobCounter) {
        maxJobCounter = MAX_JOB_COUNTER;
    }

    /**
     * @dev Check if a packed value is valid (contains valid chainId and jobCounter)
     * @param packed The packed value to validate
     * @return isValid True if the packed value is valid
     */
    function isValidPacked(uint256 packed) internal pure returns (bool isValid) {
        (uint256 chainId, uint256 jobCounter) = unpack(packed);
        isValid = (chainId <= MAX_CHAIN_ID) && (jobCounter <= MAX_JOB_COUNTER);
    }
} 