// SPDX-License-Identifier: MIT
pragma solidity >=0.8.25;

interface IAvsGovernance {
    function isOperatorRegistered(
        address operator
    ) external view returns (bool);
}

/// @title Proxy Contract for Delegated Calls
/// @notice Allows authorized keepers to execute functions on a target contract with or without ETH
/// @dev Uses immutable variables for gas efficiency and security
contract Proxy {
    address public immutable target;
    address public immutable avsGovernance;

    event FunctionExecuted(
        address indexed operator, 
        bytes callData, 
        uint256 value
    );

    constructor(address _target, address _avsGovernance) {
        target = _target;
        avsGovernance = _avsGovernance;
    }

    /// @notice Executes a function call on the target contract
    /// @param callData The encoded function call data to execute
    /// @return result The bytes returned by the function call
    /// @dev Only registered keepers can execute functions
    function executeFunction(bytes memory callData) external payable returns (bytes memory) {
        require(
            IAvsGovernance(avsGovernance).isOperatorRegistered(msg.sender),
            "Unauthorized keeper"
        );

        (bool success, bytes memory result) = target.call{value: msg.value}(callData);
        require(success, "Function call failed");
        
        emit FunctionExecuted(msg.sender, callData, msg.value);
        return result;
    }

}