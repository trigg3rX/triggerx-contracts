// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

/**
 * @title StakeUtils
 * @dev Library for stake-related calculations and utilities
 */
library StakeUtils {
    /**
     * @dev Calculates rewards based on claimed TG
     * @param claimedTG Amount of TG claimed
     * @return uint256 The calculated reward
     */
    function calculateReward(uint256 claimedTG) internal pure returns (uint256) {
        return (claimedTG * 7) / 40000; // Reward calculation with precision
    }

    /**
     * @dev Validates if a stake amount is sufficient
     * @param currentStake Current stake amount
     * @param requiredStake Required stake amount
     * @return bool Whether the stake is sufficient
     */
    function isStakeSufficient(
        uint256 currentStake,
        uint256 requiredStake
    ) internal pure returns (bool) {
        return currentStake >= requiredStake;
    }

    /**
     * @dev Calculates the total stake in a quorum
     * @param stakes Array of stake amounts
     * @return uint256 The total stake
     */
    function calculateTotalStake(
        uint256[] memory stakes
    ) internal pure returns (uint256) {
        uint256 total = 0;
        for (uint256 i = 0; i < stakes.length; i++) {
            total += stakes[i];
        }
        return total;
    }

    /**
     * @dev Validates if pool has sufficient balance
     * @param poolBalance Current pool balance
     * @param requiredAmount Required amount
     * @return bool Whether the pool has sufficient balance
     */
    function hasPoolSufficientBalance(
        uint256 poolBalance,
        uint256 requiredAmount
    ) internal pure returns (bool) {
        return poolBalance >= requiredAmount;
    }
} 