// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import {Test} from "forge-std/Test.sol";
import {TaskExecutionHub} from "../src/lz/TaskExecutionHub.sol";
import {MockEndpoint} from "./mocks/MockEndpoint.sol";
import {Origin} from "@layerzero-v2/oapp/contracts/oapp/OAppReceiver.sol";
import {console2} from "forge-std/console2.sol";
import {OwnableUpgradeable} from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import {MockJobRegistry} from "./mocks/MockJobRegistry.sol";
import {MockTriggerGasRegistry} from "./mocks/MockTriggerGasRegistry.sol";
import {ERC1967Proxy} from "@openzeppelin-contracts/contracts/proxy/ERC1967/ERC1967Proxy.sol";

// Test wrapper to expose internal functions
contract TaskExecutionHubForTest is TaskExecutionHub {
    constructor(
        address _endpoint,
        address _owner,
        uint32 _srcEid,
        uint32 _thisChainEid,
        address[] memory _initialKeepers,
        address _jobRegistryAddress,
        address _triggerGasRegistryAddress
    ) TaskExecutionHub(_endpoint, _owner) {}
    
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

contract TaskExecutionHubTest is Test {
    TaskExecutionHubForTest public taskExecutionHub;
    MockEndpoint public mockEndpoint;
    MockJobRegistry public mockJobRegistry;
    MockTriggerGasRegistry public mockTriggerGasRegistry;
    
    address public owner = address(0x1);
    address public keeper1 = address(0x100);
    address public keeper2 = address(0x101);
    address public randomUser = address(0x200);
    address public jobOwner = address(0x300);
    
    uint32 public constant SRC_EID = 10121; // L1 chain ID
    uint32 public constant THIS_EID = 20202; // This L2 chain ID
    uint32 public constant DST_EID_1 = 20203; // Another L2 chain
    uint32 public constant DST_EID_2 = 20204; // Another L2 chain
    
    event KeeperRegistered(address indexed keeper);
    event KeeperUnregistered(address indexed keeper);
    event BroadcastSent(TaskExecutionHub.ActionType action, address keeper, uint32 dstEid);
    event FunctionExecuted(address indexed keeper, address indexed target, bytes data, uint256 value);
    
    function setUp() public {
        console2.log("Starting TaskExecutionHub test setup");
        
        vm.startPrank(owner);
        
        // Deploy mock endpoint
        console2.log("Deploying MockEndpoint");
        mockEndpoint = new MockEndpoint();
        console2.log("MockEndpoint deployed at", address(mockEndpoint));
        
        // Deploy mock contracts
        console2.log("Deploying MockJobRegistry");
        mockJobRegistry = new MockJobRegistry();
        console2.log("MockJobRegistry deployed at", address(mockJobRegistry));
        
        console2.log("Deploying MockTriggerGasRegistry");
        mockTriggerGasRegistry = new MockTriggerGasRegistry();
        console2.log("MockTriggerGasRegistry deployed at", address(mockTriggerGasRegistry));
        
        // Setup initial keepers
        console2.log("Setting up initial keepers");
        address[] memory initialKeepers = new address[](1);
        initialKeepers[0] = keeper1;
        
        // Deploy TaskExecutionHub implementation
        console2.log("Deploying TaskExecutionHub");
        TaskExecutionHubForTest implementation = new TaskExecutionHubForTest(
            address(mockEndpoint), 
            owner, 
            SRC_EID, 
            THIS_EID, 
            initialKeepers, 
            address(mockJobRegistry), 
            address(mockTriggerGasRegistry)
        );
        
        // Deploy proxy with initialization
        ERC1967Proxy proxy = new ERC1967Proxy(
            address(implementation),
            abi.encodeWithSelector(
                TaskExecutionHub.initialize.selector,
                owner,
                SRC_EID,
                THIS_EID,
                initialKeepers,
                address(mockJobRegistry),
                address(mockTriggerGasRegistry)
            )
        );
        
        taskExecutionHub = TaskExecutionHubForTest(payable(address(proxy)));
        console2.log("TaskExecutionHub deployed at", address(taskExecutionHub));
        
        // Fund contract for message fees
        vm.deal(address(taskExecutionHub), 100 ether);
        
        vm.stopPrank();
        console2.log("TaskExecutionHub test setup completed");
    }
    
    function test_Constructor() public view {
        assertEq(taskExecutionHub.owner(), owner);
        assertTrue(taskExecutionHub.isKeeper(keeper1));
        assertFalse(taskExecutionHub.isKeeper(keeper2));
    }
    
    function test_AddSpokes() public {
        uint32[] memory newEids = new uint32[](2);
        newEids[0] = DST_EID_1;
        newEids[1] = DST_EID_2;
        
        vm.prank(owner);
        taskExecutionHub.addSpokes(newEids);
        
        // Check that both chains were added
        assertEq(taskExecutionHub.dstEids(0), DST_EID_1);
        assertEq(taskExecutionHub.dstEids(1), DST_EID_2);
    }
    
    function test_AddSpokes_OnlyOwner() public {
        uint32[] memory newEids = new uint32[](2);
        newEids[0] = DST_EID_1;
        newEids[1] = DST_EID_2;
        
        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, randomUser));
        vm.prank(randomUser);
        taskExecutionHub.addSpokes(newEids);
    }
    
    function test_ExecuteFunction() public {
        uint256 jobId = 1;
        uint256 tgAmount = 100;
        address target = address(0x400);
        bytes memory data = abi.encodeWithSignature("doSomething()");
        uint256 value = 1 ether;
        
        // Setup job and balance
        mockJobRegistry.setJobOwner(jobId, jobOwner);
        mockTriggerGasRegistry.setBalance(jobOwner, 1000);
        
        // Fund both the keeper and the contract
        vm.deal(keeper1, value);
        vm.deal(address(taskExecutionHub), value);
        
        // Create a mock contract that will always return success
        vm.etch(target, hex"600180600c6000396000f3006000fd"); // Simple bytecode that always returns true
        
        vm.expectEmit(true, true, true, true);
        emit FunctionExecuted(keeper1, target, data, value);
        
        vm.prank(keeper1);
        taskExecutionHub.executeFunction{value: value}(jobId, tgAmount, target, data);
        
        // Verify TG balance was deducted
        (, uint256 finalBalance) = mockTriggerGasRegistry.getBalance(jobOwner);
        assertEq(finalBalance, 900);
    }
    
    function test_ExecuteFunction_OnlyKeeper() public {
        uint256 jobId = 1;
        uint256 tgAmount = 100;
        address target = address(0x400);
        bytes memory data = abi.encodeWithSignature("doSomething()");
        
        // Setup job and balance
        mockJobRegistry.setJobOwner(jobId, jobOwner);
        mockTriggerGasRegistry.setBalance(jobOwner, 1000);
        
        vm.expectRevert("Not a keeper");
        vm.prank(randomUser);
        taskExecutionHub.executeFunction(jobId, tgAmount, target, data);
    }
    
    function test_ExecuteFunction_Failure() public {
        uint256 jobId = 1;
        uint256 tgAmount = 100;
        address target = address(0x400);
        bytes memory data = abi.encodeWithSignature("doSomething()");
        
        // Setup job and balance
        mockJobRegistry.setJobOwner(jobId, jobOwner);
        mockTriggerGasRegistry.setBalance(jobOwner, 1000);
        
        // Create a mock contract that will always revert
        vm.etch(target, hex"60006000fd"); // Simple bytecode that always reverts
        
        vm.expectRevert(bytes("Execution failed"));
        vm.prank(keeper1);
        taskExecutionHub.executeFunction(jobId, tgAmount, target, data);
    }

    function test_ExecuteFunction_JobNotFound() public {
        uint256 jobId = 999; // Non-existent job
        uint256 tgAmount = 100;
        address target = address(0x400);
        bytes memory data = abi.encodeWithSignature("doSomething()");
        
        vm.expectRevert(bytes("Job not found"));
        vm.prank(keeper1);
        taskExecutionHub.executeFunction(jobId, tgAmount, target, data);
    }

    function test_ExecuteFunction_InsufficientTGBalance() public {
        uint256 jobId = 1;
        uint256 tgAmount = 1000; // More than available balance
        address target = address(0x400);
        bytes memory data = abi.encodeWithSignature("doSomething()");
        
        // Setup job with insufficient balance
        mockJobRegistry.setJobOwner(jobId, jobOwner);
        mockTriggerGasRegistry.setBalance(jobOwner, 500); // Less than required
        
        vm.expectRevert(bytes("Insufficient TG balance"));
        vm.prank(keeper1);
        taskExecutionHub.executeFunction(jobId, tgAmount, target, data);
    }
    
    function test_LzReceive_Register() public {
        // Simulate a message from AvsGovernanceLogic to register a new keeper
        bytes memory payload = abi.encode(TaskExecutionHub.ActionType.REGISTER, keeper2);
        
        // Create Origin struct directly with the 3 required fields
        Origin memory origin = Origin({
            srcEid: SRC_EID,
            sender: bytes32(uint256(uint160(address(0x1)))),
            nonce: 1
        });
        
        vm.expectEmit(true, false, false, false);
        emit KeeperRegistered(keeper2);
        
        // Call our exposed function to test _lzReceive
        taskExecutionHub.exposed_lzReceive(
            origin,
            bytes32(0),
            payload,
            address(0),
            bytes("")
        );
        
        // Verify keeper was registered
        assertTrue(taskExecutionHub.isKeeper(keeper2));
    }
    
    function test_LzReceive_Unregister() public {
        // First register keeper2
        bytes memory registerPayload = abi.encode(TaskExecutionHub.ActionType.REGISTER, keeper2);
        
        Origin memory origin = Origin({
            srcEid: SRC_EID,
            sender: bytes32(uint256(uint160(address(0x1)))),
            nonce: 1
        });
        
        taskExecutionHub.exposed_lzReceive(
            origin,
            bytes32(0),
            registerPayload,
            address(0),
            bytes("")
        );
        
        assertTrue(taskExecutionHub.isKeeper(keeper2));
        
        // Now unregister
        bytes memory unregisterPayload = abi.encode(TaskExecutionHub.ActionType.UNREGISTER, keeper2);
        
        vm.expectEmit(true, false, false, false);
        emit KeeperUnregistered(keeper2);
        
        taskExecutionHub.exposed_lzReceive(
            origin,
            bytes32(0),
            unregisterPayload,
            address(0),
            bytes("")
        );
        
        // Verify keeper was unregistered
        assertFalse(taskExecutionHub.isKeeper(keeper2));
    }
    
    function test_Withdraw() public {
        uint256 initialBalance = 10 ether;
        address payable recipient = payable(address(0x300));
        
        // Fund contract
        vm.deal(address(taskExecutionHub), initialBalance);
        
        // Withdraw funds
        vm.prank(owner);
        taskExecutionHub.withdraw(recipient, 5 ether);
        
        assertEq(address(taskExecutionHub).balance, 5 ether);
        assertEq(recipient.balance, 5 ether);
    }
    
    function test_Withdraw_OnlyOwner() public {
        address payable recipient = payable(address(0x300));
        
        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, randomUser));
        vm.prank(randomUser);
        taskExecutionHub.withdraw(recipient, 1 ether);
    }
    
    function test_ReceiveEther() public {
        uint256 amount = 5 ether;
        address sender = address(0x123);
        
        // Fund the sender
        vm.deal(sender, amount);
        
        uint256 initialBalance = address(taskExecutionHub).balance;
        
        vm.prank(sender);
        (bool success,) = address(taskExecutionHub).call{value: amount}("");
        
        assertTrue(success);
        assertEq(address(taskExecutionHub).balance, initialBalance + amount);
    }
    
    function test_SetPeer() public {
        uint32 newSrcEid = 10122;
        bytes32 newAvsGovernance = bytes32(uint256(uint160(address(0x456))));
        
        vm.prank(owner);
        taskExecutionHub.setPeer(newSrcEid, newAvsGovernance); 
        
        // We would need a getter function to verify this
        // Instead, we can test that setPeer is called correctly by looking at emitted events
        // or checking that subsequent operations with the new peer work as expected
    }
    
    function test_SetPeer_OnlyOwner() public {
        uint32 newSrcEid = 10122;
        bytes32 newAvsGovernance = bytes32(uint256(uint160(address(0x456))));
        
        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, randomUser));
        vm.prank(randomUser);
        taskExecutionHub.setPeer(newSrcEid, newAvsGovernance);
    }
} 