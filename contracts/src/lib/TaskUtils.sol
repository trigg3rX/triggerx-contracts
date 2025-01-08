// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

/**
 * @title TaskUtils
 * @dev Library for task-related utility functions
 */
library TaskUtils {
    /**
     * @dev Generates a task ID from job ID and task number
     * @param jobId The ID of the job
     * @param taskNum The task number
     * @return bytes8 The generated task ID
     */
    function generateTaskId(uint32 jobId, uint32 taskNum) internal pure returns (bytes8) {
        return bytes8(abi.encodePacked(jobId, taskNum));
    }

    /**
     * @dev Validates task response window
     * @param taskCreatedBlock The block number when task was created
     * @param responseWindowBlocks Number of blocks within which response is valid
     * @return bool Whether the response is within valid window
     */
    function isWithinResponseWindow(
        uint32 taskCreatedBlock,
        uint32 responseWindowBlocks
    ) internal view returns (bool) {
        return block.number <= taskCreatedBlock + responseWindowBlocks;
    }

    /**
     * @dev Checks if quorum threshold is met
     * @param signedStake Amount of stake that signed
     * @param totalStake Total stake in the quorum
     * @param thresholdBps Threshold in basis points (100 = 1%)
     * @return bool Whether the threshold is met
     */
    function isQuorumThresholdMet(
        uint256 signedStake,
        uint256 totalStake,
        uint256 thresholdBps
    ) internal pure returns (bool) {
        return signedStake * 10000 >= totalStake * thresholdBps;
    }
} 