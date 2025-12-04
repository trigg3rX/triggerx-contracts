// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

/**
 * @title PackedJobIdLib
 * @dev Library for packing and unpacking job identifiers into a single uint256.
 *
 * Layout (from most-significant bits to least-significant bits):
 * - [255..224] 32 bits  : chainId
 * - [223..160] 64 bits  : timestamp (seconds since epoch, good until year ~584 billion)
 * - [159..0]   160 bits : job counter
 *
 * This ensures:
 * - Uniqueness across chains        → chainId
 * - Temporal information            → timestamp
 * - Plenty of room for counters     → 160-bit counter
 *
 * NOTE: The high 32 bits remain chainId (same position as the old layout) so that
 *       getChainId() on old packed IDs still returns the correct value.
 */
library PackedJobIdLib {
    // ----- Constants for bit allocation -----
    uint256 private constant CHAIN_ID_BITS = 32;
    uint256 private constant TIMESTAMP_BITS = 64;
    uint256 private constant COUNTER_BITS = 160;

    // ----- Derived constants -----
    uint256 private constant COUNTER_MASK = (uint256(1) << COUNTER_BITS) - 1;
    uint256 private constant TIMESTAMP_MASK = (uint256(1) << TIMESTAMP_BITS) - 1;

    uint256 private constant CHAIN_ID_SHIFT = TIMESTAMP_BITS + COUNTER_BITS; // 224
    uint256 private constant TIMESTAMP_SHIFT = COUNTER_BITS; // 160

    // Maximum values for validation
    uint256 private constant MAX_CHAIN_ID = (uint256(1) << CHAIN_ID_BITS) - 1; // 2^32 - 1
    uint256 private constant MAX_TIMESTAMP = (uint256(1) << TIMESTAMP_BITS) - 1; // 2^64 - 1
    uint256 private constant MAX_JOB_COUNTER = COUNTER_MASK; // 2^160 - 1

    // Custom errors
    error ChainIdOverflow(uint256 chainId, uint256 maxChainId);
    error JobCounterOverflow(uint256 jobCounter, uint256 maxJobCounter);
    error TimestampOverflow(uint256 timestamp, uint256 maxTimestamp);

    /**
     * @dev Pack chainId, timestamp, and jobCounter into a single uint256.
     * @param chainId The chain ID (must fit in 32 bits).
     * @param timestamp The timestamp in seconds (must fit in 64 bits).
     * @param jobCounter The job counter (must fit in 160 bits).
     * @return packed The packed value.
     */
    function pack(
        uint256 chainId,
        uint256 timestamp,
        uint256 jobCounter
    ) internal pure returns (uint256 packed) {
        // Validate inputs
        if (chainId > MAX_CHAIN_ID) {
            revert ChainIdOverflow(chainId, MAX_CHAIN_ID);
        }
        if (timestamp > MAX_TIMESTAMP) {
            revert TimestampOverflow(timestamp, MAX_TIMESTAMP);
        }
        if (jobCounter > MAX_JOB_COUNTER) {
            revert JobCounterOverflow(jobCounter, MAX_JOB_COUNTER);
        }

        // Pack into 256 bits according to layout described above
        packed =
            (chainId << CHAIN_ID_SHIFT) |
            ((timestamp & TIMESTAMP_MASK) << TIMESTAMP_SHIFT) |
            (jobCounter & COUNTER_MASK);
    }

    /**
     * @dev Unpack a packed value into its components.
     * @param packed The packed value to unpack
     * @return chainId The extracted chain ID
     * @return timestamp The extracted timestamp
     * @return jobCounter The extracted job counter
     */
    function unpack(
        uint256 packed
    ) internal pure returns (uint256 chainId, uint256 timestamp, uint256 jobCounter) {
        jobCounter = packed & COUNTER_MASK;
        timestamp = (packed >> TIMESTAMP_SHIFT) & TIMESTAMP_MASK;
        chainId = packed >> CHAIN_ID_SHIFT;
    }

    /**
     * @dev Extract only the chainId from a packed value
     * @param packed The packed value
     * @return chainId The extracted chain ID
     */
    function getChainId(uint256 packed) internal pure returns (uint256 chainId) {
        chainId = packed >> CHAIN_ID_SHIFT;
    }

    /**
     * @dev Extract only the timestamp from a packed value
     * @param packed The packed value
     * @return timestamp The extracted timestamp
     */
    function getTimestamp(uint256 packed) internal pure returns (uint256 timestamp) {
        timestamp = (packed >> TIMESTAMP_SHIFT) & TIMESTAMP_MASK;
    }

    /**
     * @dev Extract only the jobCounter from a packed value
     * @param packed The packed value
     * @return jobCounter The extracted job counter
     */
    function getJobCounter(uint256 packed) internal pure returns (uint256 jobCounter) {
        jobCounter = packed & COUNTER_MASK;
    }

    /**
     * @dev Get the maximum allowed chain ID
     * @return maxChainId The maximum chain ID (2^32 - 1)
     */
    function getMaxChainId() internal pure returns (uint256 maxChainId) {
        maxChainId = MAX_CHAIN_ID;
    }

    /**
     * @dev Get the maximum allowed timestamp
     * @return maxTimestamp The maximum timestamp (2^64 - 1)
     */
    function getMaxTimestamp() internal pure returns (uint256 maxTimestamp) {
        maxTimestamp = MAX_TIMESTAMP;
    }

    /**
     * @dev Get the maximum allowed job counter
     * @return maxJobCounter The maximum job counter (2^160 - 1)
     */
    function getMaxJobCounter() internal pure returns (uint256 maxJobCounter) {
        maxJobCounter = MAX_JOB_COUNTER;
    }

    /**
     * @dev Check if a packed value is valid (contains valid chainId, timestamp and jobCounter)
     * @param packed The packed value to validate
     * @return isValid True if the packed value is valid
     */
    function isValidPacked(uint256 packed) internal pure returns (bool isValid) {
        (uint256 chainId, uint256 timestamp, uint256 jobCounter) = unpack(packed);
        isValid =
            (chainId <= MAX_CHAIN_ID) &&
            (timestamp <= MAX_TIMESTAMP) &&
            (jobCounter <= MAX_JOB_COUNTER);
    }
}
