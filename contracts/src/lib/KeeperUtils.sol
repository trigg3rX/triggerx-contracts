// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

/**
 * @title KeeperUtils
 * @dev Library for keeper/operator-related utility functions
 */
library KeeperUtils {
    /**
     * @dev Validates if an operator is active and not blacklisted
     * @param isBlacklisted Mapping of operator address to blacklist status
     * @param operator Address of the operator to check
     * @return bool Whether the operator is valid
     */
    function isValidOperator(
        mapping(address => bool) storage isBlacklisted,
        address operator
    ) internal view returns (bool) {
        return operator != address(0) && !isBlacklisted[operator];
    }

    /**
     * @dev Validates a quorum of operators
     * @param operators Array of operator addresses
     * @param quorumThreshold Minimum number of operators required
     * @return bool Whether the quorum is valid
     */
    function hasValidQuorum(
        address[] memory operators,
        uint256 quorumThreshold
    ) internal pure returns (bool) {
        return operators.length >= quorumThreshold;
    }

    /**
     * @dev Checks if an operator is part of a quorum
     * @param operators Array of operator addresses
     * @param operator Address of the operator to check
     * @return bool Whether the operator is in the quorum
     */
    function isOperatorInQuorum(
        address[] memory operators,
        address operator
    ) internal pure returns (bool) {
        for (uint256 i = 0; i < operators.length; i++) {
            if (operators[i] == operator) {
                return true;
            }
        }
        return false;
    }
} 