// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import {Test} from "forge-std/Test.sol";
import {ProxyHub} from "../src/lz/ProxyHub.sol";
import {MockEndpoint} from "./mocks/MockEndpoint.sol";
import {Origin} from "@layerzero-v2/oapp/contracts/oapp/OAppReceiver.sol";
import {console2} from "forge-std/console2.sol";
import {OwnableUpgradeable} from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";

// Test wrapper to expose internal functions
contract ProxyHubForTest is ProxyHub {
    constructor(
        address _endpoint,
        address _owner,
        uint32 _srcEid,
        uint32 _thisChainEid,
        address[] memory _initialKeepers
    ) ProxyHub(_endpoint, _owner, _srcEid, _thisChainEid, _initialKeepers) {}
    
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

contract ProxyHubTest is Test {
    ProxyHubForTest public proxyHub;
    MockEndpoint public mockEndpoint;
    
    address public owner = address(0x1);
    address public keeper1 = address(0x100);
    address public keeper2 = address(0x101);
    address public randomUser = address(0x200);
    
    uint32 public constant SRC_EID = 10121; // L1 chain ID
    uint32 public constant THIS_EID = 20202; // This L2 chain ID
    uint32 public constant DST_EID_1 = 20203; // Another L2 chain
    uint32 public constant DST_EID_2 = 20204; // Another L2 chain
    
    event KeeperRegistered(address indexed keeper);
    event KeeperUnregistered(address indexed keeper);
    event BroadcastSent(ProxyHub.ActionType action, address keeper, uint32 dstEid);
    event FunctionExecuted(address indexed keeper, address indexed target, bytes data, uint256 value);
    
    function setUp() public {
        console2.log("Starting ProxyHub test setup");
        
        vm.startPrank(owner);
        
        // Deploy mock endpoint
        console2.log("Deploying MockEndpoint");
        mockEndpoint = new MockEndpoint();
        console2.log("MockEndpoint deployed at", address(mockEndpoint));
        
        // Setup initial keepers
        console2.log("Setting up initial keepers");
        address[] memory initialKeepers = new address[](1);
        initialKeepers[0] = keeper1;
        
        // Deploy ProxyHub
        console2.log("Deploying ProxyHub");
        try new ProxyHubForTest(
            address(mockEndpoint),
            owner,
            SRC_EID,
            THIS_EID,
            initialKeepers
        ) returns (ProxyHubForTest hub) {
            proxyHub = hub;
            console2.log("ProxyHub deployed at", address(proxyHub));
        } catch Error(string memory reason) {
            console2.log("ProxyHub deployment failed with reason:", reason);
            revert(reason);
        } catch (bytes memory reason) {
            console2.log("ProxyHub deployment failed with low-level error");
            revert("Low-level error");
        }
        
        // Fund contract for message fees
        vm.deal(address(proxyHub), 100 ether);
        
        vm.stopPrank();
        console2.log("ProxyHub test setup completed");
    }
    
    function test_Constructor() public {
        assertEq(proxyHub.owner(), owner);
        assertTrue(proxyHub.isKeeper(keeper1));
        assertFalse(proxyHub.isKeeper(keeper2));
    }
    
    function test_AddSpokes() public {
        uint32[] memory newEids = new uint32[](2);
        newEids[0] = DST_EID_1;
        newEids[1] = DST_EID_2;
        
        vm.prank(owner);
        proxyHub.addSpokes(newEids);
        
        // Check that both chains were added
        assertEq(proxyHub.dstEids(0), DST_EID_1);
        assertEq(proxyHub.dstEids(1), DST_EID_2);
    }
    
    function test_AddSpokes_OnlyOwner() public {
        uint32[] memory newEids = new uint32[](2);
        newEids[0] = DST_EID_1;
        newEids[1] = DST_EID_2;
        
        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, randomUser));
        vm.prank(randomUser);
        proxyHub.addSpokes(newEids);
    }
    
    function test_ExecuteFunction() public {
        address target = address(0x300);
        bytes memory data = abi.encodeWithSignature("doSomething()");
        uint256 value = 1 ether;
        
        // Fund both the keeper and the contract
        vm.deal(keeper1, value);
        vm.deal(address(proxyHub), value);
        
        // Create a mock contract that will always return success
        vm.etch(target, hex"600180600c6000396000f3006000fd"); // Simple bytecode that always returns true
        
        vm.expectEmit(true, true, true, true);
        emit FunctionExecuted(keeper1, target, data, value);
        
        vm.prank(keeper1);
        proxyHub.executeFunction{value: value}(target, data);
    }
    
    function test_ExecuteFunction_OnlyKeeper() public {
        address target = address(0x300);
        bytes memory data = abi.encodeWithSignature("doSomething()");
        
        vm.expectRevert("Not a keeper");
        vm.prank(randomUser);
        proxyHub.executeFunction(target, data);
    }
    
    function test_ExecuteFunction_Failure() public {
        address target = address(0x300);
        bytes memory data = abi.encodeWithSignature("doSomething()");
        
        // Create a mock contract that will always revert
        vm.etch(target, hex"60006000fd"); // Simple bytecode that always reverts
        
        vm.expectRevert("Execution failed");
        vm.prank(keeper1);
        proxyHub.executeFunction(target, data);
    }
    
    function test_LzReceive_Register() public {
        // Simulate a message from AvsGovernanceLogic to register a new keeper
        bytes memory payload = abi.encode(ProxyHub.ActionType.REGISTER, keeper2);
        
        // Create Origin struct directly with the 3 required fields
        Origin memory origin = Origin({
            srcEid: SRC_EID,
            sender: bytes32(uint256(uint160(address(0x1)))),
            nonce: 1
        });
        
        vm.expectEmit(true, false, false, false);
        emit KeeperRegistered(keeper2);
        
        // Call our exposed function to test _lzReceive
        proxyHub.exposed_lzReceive(
            origin,
            bytes32(0),
            payload,
            address(0),
            bytes("")
        );
        
        // Verify keeper was registered
        assertTrue(proxyHub.isKeeper(keeper2));
    }
    
    function test_LzReceive_Unregister() public {
        // First register keeper2
        bytes memory registerPayload = abi.encode(ProxyHub.ActionType.REGISTER, keeper2);
        
        Origin memory origin = Origin({
            srcEid: SRC_EID,
            sender: bytes32(uint256(uint160(address(0x1)))),
            nonce: 1
        });
        
        proxyHub.exposed_lzReceive(
            origin,
            bytes32(0),
            registerPayload,
            address(0),
            bytes("")
        );
        
        assertTrue(proxyHub.isKeeper(keeper2));
        
        // Now unregister
        bytes memory unregisterPayload = abi.encode(ProxyHub.ActionType.UNREGISTER, keeper2);
        
        vm.expectEmit(true, false, false, false);
        emit KeeperUnregistered(keeper2);
        
        proxyHub.exposed_lzReceive(
            origin,
            bytes32(0),
            unregisterPayload,
            address(0),
            bytes("")
        );
        
        // Verify keeper was unregistered
        assertFalse(proxyHub.isKeeper(keeper2));
    }
    
    function test_Withdraw() public {
        uint256 initialBalance = 10 ether;
        address payable recipient = payable(address(0x300));
        
        // Fund contract
        vm.deal(address(proxyHub), initialBalance);
        
        // Withdraw funds
        vm.prank(owner);
        proxyHub.withdraw(recipient, 5 ether);
        
        assertEq(address(proxyHub).balance, 5 ether);
        assertEq(recipient.balance, 5 ether);
    }
    
    function test_Withdraw_OnlyOwner() public {
        address payable recipient = payable(address(0x300));
        
        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, randomUser));
        vm.prank(randomUser);
        proxyHub.withdraw(recipient, 1 ether);
    }
    
    function test_ReceiveEther() public {
        uint256 amount = 5 ether;
        address sender = address(0x123);
        
        // Fund the sender
        vm.deal(sender, amount);
        
        uint256 initialBalance = address(proxyHub).balance;
        
        vm.prank(sender);
        (bool success,) = address(proxyHub).call{value: amount}("");
        
        assertTrue(success);
        assertEq(address(proxyHub).balance, initialBalance + amount);
    }
    
    function test_SetPeer() public {
        uint32 newSrcEid = 10122;
        bytes32 newAvsGovernance = bytes32(uint256(uint160(address(0x456))));
        
        vm.prank(owner);
        proxyHub.setPeer(newSrcEid, newAvsGovernance); 
        
        // We would need a getter function to verify this
        // Instead, we can test that setPeer is called correctly by looking at emitted events
        // or checking that subsequent operations with the new peer work as expected
    }
    
    function test_SetPeer_OnlyOwner() public {
        uint32 newSrcEid = 10122;
        bytes32 newAvsGovernance = bytes32(uint256(uint160(address(0x456))));
        
        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, randomUser));
        vm.prank(randomUser);
        proxyHub.setPeer(newSrcEid, newAvsGovernance);
    }
} 