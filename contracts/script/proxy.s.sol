// SPDX-License-Identifier: MIT
pragma solidity >=0.8.25;

import "forge-std/Script.sol";
import "../src/Proxy.sol";
import "../src/ProxyFactory.sol";
import "../src/test/Target.sol";

/// @title Proxy Deployment and Interaction Script
/// @notice Demonstrates interaction with existing ProxyFactory and deployment of proxies
contract CallTargetViaProxy is Script {
    // Existing ProxyFactory address
    address constant PROXY_FACTORY = 0xFfF9083b5126D597F8Bb1A25726879b317d7a8F8;
    // Holesky AVS Governance address
    address constant AVS_GOVERNANCE = 0x0C77B6273F4852200b17193837960b2f253518FC;

    function run() external {
        // Load the keeper's private key from .env
        uint256 keeperPrivateKey = vm.envUint("KEEPER_PRIVATE_KEY");
        address keeper = vm.addr(keeperPrivateKey);

        // Generate a unique salt based on timestamp and keeper address
        bytes32 salt = keccak256(abi.encodePacked(block.timestamp, keeper));

        vm.startBroadcast(keeperPrivateKey);

        // Deploy the Target contract
        Target target = new Target(address(0));
        console.log("Target deployed at:", address(target));

        // Use existing ProxyFactory
        ProxyFactory factory = ProxyFactory(PROXY_FACTORY);
        console.log("Using ProxyFactory at:", address(factory));

        // Get predicted proxy address
        address predictedAddr = factory.getProxyAddress(address(target), salt);
        console.log("Predicted proxy address:", predictedAddr);

        // Deploy the proxy
        address proxyAddr = factory.deployProxy(address(target), salt);
        console.log("Proxy deployed at:", proxyAddr);
        require(proxyAddr == predictedAddr, "Proxy address mismatch");

        // Update the Target's allowedCaller
        target.setAllowedCaller(proxyAddr);
        console.log("Target allowedCaller set to:", target.allowedCaller());

        // Test the proxy with increment function
        bytes memory callData = abi.encodeWithSignature("increment(uint256)", 5);
        bytes memory result = Proxy(proxyAddr).executeFunction(callData);
        uint256 decodedResult = abi.decode(result, (uint256));
        console.log("Result of increment(5):", decodedResult);

        vm.stopBroadcast();
    }
}