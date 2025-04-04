// SPDX-License-Identifier: MIT
pragma solidity >=0.8.25;

import "./Proxy.sol";

/// @title Proxy Factory Contract for Deterministic Proxy Deployment
/// @notice Deploys proxy contracts with deterministic addresses using CREATE2
/// @dev Uses CREATE2 opcode to enable predictable proxy addresses based on salt values
contract ProxyFactory {
    address public immutable avsGovernance;
    mapping(bytes32 => address) public deployedProxies;

    event ProxyDeployed(address indexed proxy, address indexed target, bytes32 salt);

    constructor(address _avsGovernance) {
        avsGovernance = _avsGovernance;
    }

    /// @notice Deploys a new proxy contract with a deterministic address
    /// @param target The address of the contract to be proxied
    /// @param salt Unique value for deterministic address generation
    /// @return proxyAddress The address of the deployed proxy contract
    /// @dev Uses CREATE2 for deterministic deployment. Will revert if proxy already exists with given salt
    function deployProxy(address target, bytes32 salt) external returns (address) {
        require(deployedProxies[salt] == address(0), "Proxy already deployed with this salt");
        
        address proxyAddress = _computeProxyAddress(target, salt);
        Proxy proxy = new Proxy{salt: salt}(target, avsGovernance);
        require(address(proxy) == proxyAddress, "Proxy address mismatch");
        
        deployedProxies[salt] = proxyAddress;
        emit ProxyDeployed(proxyAddress, target, salt);
        
        return proxyAddress;
    }

    /// @notice Computes the address where a proxy contract will be deployed
    /// @param target The address of the contract to be proxied
    /// @param salt Unique value for deterministic address generation
    /// @return The address where the proxy would be deployed
    /// @dev This function can be used to predict the proxy address before deployment
    function getProxyAddress(address target, bytes32 salt) external view returns (address) {
        return _computeProxyAddress(target, salt);
    }

    /// @dev Internal function to compute the deterministic proxy address
    /// @param target The target contract address
    /// @param salt The salt value for CREATE2
    /// @return The computed proxy address
    function _computeProxyAddress(address target, bytes32 salt) internal view returns (address) {
        bytes32 bytecodeHash = keccak256(
            abi.encodePacked(
                type(Proxy).creationCode,
                abi.encode(target, avsGovernance)
            )
        );
        
        bytes32 hash = keccak256(
            abi.encodePacked(
                bytes1(0xff),
                address(this),
                salt,
                bytecodeHash
            )
        );
        
        return address(uint160(uint256(hash)));
    }
}