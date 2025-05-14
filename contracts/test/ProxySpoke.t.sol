// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import {Test} from "forge-std/Test.sol";
import {ProxySpoke} from "../src/lz/ProxySpoke.sol";
import {MockEndpoint} from "./mocks/MockEndpoint.sol";
import {Origin} from "@layerzero-v2/oapp/contracts/oapp/OAppReceiver.sol";

// Test wrapper to expose internal functions
contract ProxySpokeForTest is ProxySpoke {
    constructor(
        address _endpoint, 
        address _owner, 
        uint32 _srcEid,
        address[] memory _initialKeepers
    ) ProxySpoke(_endpoint, _owner, _srcEid, _initialKeepers) {}
    
    // Expose internal _lzReceive function for testing
    function exposed_lzReceive(
        Origin calldata _origin,
        bytes32 _guid,
        bytes calldata _message,
        address _executor,
        bytes calldata _extraData
    ) external {
        _lzReceive(_origin, _guid, _message, _executor, _extraData);
    }
}

contract ProxySpokeTest is Test {
    ProxySpokeForTest public proxySpoke;
    MockEndpoint public mockEndpoint;
    
    address public owner = address(0x1);
    address public keeper1 = address(0x100);
    address public keeper2 = address(0x101);
    address public randomUser = address(0x200);
    
    uint32 public constant SRC_EID = 10121; // L1 or Hub chain ID
    
    event KeeperUpdated(ProxySpoke.ActionType action, address keeper);
    event FunctionExecuted(address indexed keeper, address indexed target, bytes data, uint256 value);
    
    function setUp() public {
        vm.startPrank(owner);
        
        // Deploy mock endpoint
        mockEndpoint = new MockEndpoint();
        
        // Setup initial keepers
        address[] memory initialKeepers = new address[](1);
        initialKeepers[0] = keeper1;
        
        // Deploy ProxySpoke
        proxySpoke = new ProxySpokeForTest(
            address(mockEndpoint),
            owner,
            SRC_EID,
            initialKeepers
        );
        
        vm.stopPrank();
    }
    
    function test_Constructor() public {
        assertEq(proxySpoke.owner(), owner);
        assertTrue(proxySpoke.isKeeper(keeper1));
        assertFalse(proxySpoke.isKeeper(keeper2));
    }
    
    function test_ExecuteFunction() public {
        address target = address(0x300);
        bytes memory data = abi.encodeWithSignature("doSomething()");
        uint256 value = 1 ether;
        
        // Fund the keeper
        vm.deal(keeper1, value);
        
        // Create a mock contract that will always return success
        vm.etch(target, hex"600180600c6000396000f3006000fd"); // Simple bytecode that always returns true
        
        vm.expectEmit(true, true, true, true);
        emit FunctionExecuted(keeper1, target, data, value);
        
        vm.prank(keeper1);
        proxySpoke.executeFunction{value: value}(target, data);
    }
    
    function test_ExecuteFunction_OnlyKeeper() public {
        address target = address(0x300);
        bytes memory data = abi.encodeWithSignature("doSomething()");
        
        vm.expectRevert("Spoke: Keeper not registered");
        vm.prank(randomUser);
        proxySpoke.executeFunction(target, data);
    }
    
    function test_ExecuteFunction_Failure() public {
        address target = address(0x300);
        bytes memory data = abi.encodeWithSignature("doSomething()");
        
        // Create a mock contract that will always revert
        vm.etch(target, hex"60006000fd"); // Simple bytecode that always reverts
        
        vm.expectRevert("Function execution failed");
        vm.prank(keeper1);
        proxySpoke.executeFunction(target, data);
    }
    
    function test_LzReceive_Register() public {
        // Simulate a message from ProxyHub to register a new keeper
        bytes memory payload = abi.encode(ProxySpoke.ActionType.REGISTER, keeper2);
        
        // Create Origin struct directly with the 3 required fields
        Origin memory origin = Origin({
            srcEid: SRC_EID,
            sender: bytes32(uint256(uint160(address(0x1)))),
            nonce: 1
        });
        
        vm.expectEmit(true, false, false, false);
        emit KeeperUpdated(ProxySpoke.ActionType.REGISTER, keeper2);
        
        // Call our exposed function to test _lzReceive
        proxySpoke.exposed_lzReceive(
            origin,
            bytes32(0),
            payload,
            address(0),
            bytes("")
        );
        
        // Verify keeper was registered
        assertTrue(proxySpoke.isKeeper(keeper2));
    }
    
    function test_LzReceive_Unregister() public {
        // First register keeper2
        bytes memory registerPayload = abi.encode(ProxySpoke.ActionType.REGISTER, keeper2);
        
        Origin memory origin = Origin({
            srcEid: SRC_EID,
            sender: bytes32(uint256(uint160(address(0x1)))),
            nonce: 1
        });
        
        proxySpoke.exposed_lzReceive(
            origin,
            bytes32(0),
            registerPayload,
            address(0),
            bytes("")
        );
        
        assertTrue(proxySpoke.isKeeper(keeper2));
        
        // Now unregister
        bytes memory unregisterPayload = abi.encode(ProxySpoke.ActionType.UNREGISTER, keeper2);
        
        vm.expectEmit(true, false, false, false);
        emit KeeperUpdated(ProxySpoke.ActionType.UNREGISTER, keeper2);
        
        proxySpoke.exposed_lzReceive(
            origin,
            bytes32(0),
            unregisterPayload,
            address(0),
            bytes("")
        );
        
        // Verify keeper was unregistered
        assertFalse(proxySpoke.isKeeper(keeper2));
    }
    
    function test_ReceiveEther() public {
        uint256 amount = 5 ether;
        address sender = address(0x123);
        
        // Fund the sender
        vm.deal(sender, amount);
        
        uint256 initialBalance = address(proxySpoke).balance;
        
        vm.prank(sender);
        (bool success,) = address(proxySpoke).call{value: amount}("");
        
        assertTrue(success);
        assertEq(address(proxySpoke).balance, initialBalance + amount);
    }
} 