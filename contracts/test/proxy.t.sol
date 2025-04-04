// SPDX-License-Identifier: MIT
pragma solidity >=0.8.25;

import "forge-std/Test.sol";
import "../src/Proxy.sol";
import "../src/ProxyFactory.sol";
import "../src/test/Target.sol";
import "../src/test/MockAvsGovernance.sol";

contract ProxyFactoryTest is Test {
    ProxyFactory factory;
    MockAvsGovernance avsGovernance;
    Target target;
    address keeper;
    address unauthorized;
    bytes32 constant TEST_SALT = keccak256("test.salt.1");

    event FunctionExecuted(
        address indexed operator, 
        bytes callData, 
        uint256 value
    );

    function setUp() public {
        avsGovernance = new MockAvsGovernance();
        factory = new ProxyFactory(address(avsGovernance));
        
        keeper = address(0x123);
        unauthorized = address(0x456);
        
        vm.deal(keeper, 100 ether); // Give keeper some ETH
        avsGovernance.setOperatorRegistered(keeper, true);
        target = new Target(address(0));
    }

    function testDeployProxy() public {
        address predictedAddr = factory.getProxyAddress(address(target), TEST_SALT);
        address proxyAddr = factory.deployProxy(address(target), TEST_SALT);
        
        assertEq(proxyAddr, predictedAddr, "Predicted address mismatch");
        assertTrue(proxyAddr != address(0), "Proxy deployment failed");
        
        Proxy proxy = Proxy(proxyAddr);
        assertEq(proxy.target(), address(target), "Target address mismatch");
        assertEq(proxy.avsGovernance(), address(avsGovernance), "AvsGovernance address mismatch");
    }

    function testCannotDeployWithSameSalt() public {
        factory.deployProxy(address(target), TEST_SALT);
        
        vm.expectRevert("Proxy already deployed with this salt");
        factory.deployProxy(address(target), TEST_SALT);
    }

    function testOnlyRegisteredKeeperCanCall() public {
        address proxyAddr = factory.deployProxy(address(target), TEST_SALT);
        Proxy proxy = Proxy(proxyAddr);
        target.setAllowedCaller(proxyAddr);

        bytes memory callData = abi.encodeWithSignature("increment(uint256)", 5);
        
        vm.prank(keeper);
        bytes memory result = proxy.executeFunction{value: 1 ether}(callData);
        uint256 decoded = abi.decode(result, (uint256));
        assertEq(decoded, 6, "Increment function failed");

        vm.prank(unauthorized);
        vm.expectRevert("Unauthorized keeper");
        proxy.executeFunction(callData);
    }

    function testFunctionExecutionWithValue() public {
        address proxyAddr = factory.deployProxy(address(target), TEST_SALT);
        Proxy proxy = Proxy(proxyAddr);
        target.setAllowedCaller(proxyAddr);

        uint256 initialBalance = address(target).balance;
        bytes memory callData = abi.encodeWithSignature("increment(uint256)", 10);
        uint256 ethValue = 1 ether;

        vm.prank(keeper);
        bytes memory result = proxy.executeFunction{value: ethValue}(callData);
        uint256 decoded = abi.decode(result, (uint256));
        
        assertEq(decoded, 11, "Function execution returned incorrect result");
        assertEq(
            address(target).balance, 
            initialBalance + ethValue, 
            "ETH value not transferred correctly"
        );
    }

    function testFunctionExecutionWithValueFailsOnInsufficientBalance() public {
        address proxyAddr = factory.deployProxy(address(target), TEST_SALT);
        Proxy proxy = Proxy(proxyAddr);
        target.setAllowedCaller(proxyAddr);

        bytes memory callData = abi.encodeWithSignature("increment(uint256)", 10);
        
        // Set keeper's balance to exactly 1 ETH
        vm.deal(keeper, 1 ether);
        
        vm.prank(keeper);
        // Try to send more ETH than keeper has
        vm.expectRevert("Insufficient balance");  // Remove this line since we expect a lower level revert
        (bool success,) = address(proxy).call{value: 2 ether}(
            abi.encodeWithSelector(Proxy.executeFunction.selector, callData)
        );
        assertFalse(success, "Transaction should fail due to insufficient funds");
    }

    function testFunctionExecutionWithoutValue() public {
        address proxyAddr = factory.deployProxy(address(target), TEST_SALT);
        Proxy proxy = Proxy(proxyAddr);
        target.setAllowedCaller(proxyAddr);

        bytes memory callData = abi.encodeWithSignature("increment(uint256)", 10);

        // Call without any ETH value
        vm.prank(keeper);
        bytes memory result = proxy.executeFunction(callData); // No value specified
        uint256 decoded = abi.decode(result, (uint256));
        
        assertEq(decoded, 11, "Function execution returned incorrect result");
        assertEq(address(target).balance, 0, "Target should not receive any ETH");
    }
}