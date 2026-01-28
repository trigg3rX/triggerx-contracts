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

    uint32 public constant SRC_EID = 10_121; // L1 chain ID
    uint32 public constant THIS_EID = 20_202; // This L2 chain ID
    uint32 public constant DST_EID_1 = 20_203; // Another L2 chain
    uint32 public constant DST_EID_2 = 20_204; // Another L2 chain

    event KeeperRegistered(address indexed keeper);
    event KeeperUnregistered(address indexed keeper);
    event BroadcastSent(TaskExecutionHub.ActionType action, address keeper, uint32 dstEid);
    event FunctionExecuted(
        address indexed keeper, address indexed target, bytes data, uint256 value
    );
    event FunctionExecutionFailed(
        address indexed keeper, address indexed target, bytes data, uint256 value, bytes reason
    );
    event TaskDispatcherUpdated(
        address indexed oldTaskDispatcher, address indexed newTaskDispatcher
    );

    // Helper to create signature params
    function _dummySignatureParams(
        uint256 jobId,
        address target,
        bytes memory data,
        address keeper
    ) internal view returns (uint256 deadline, bytes memory signature) {
        deadline = block.timestamp + 1 hours;
        bytes32 hash = keccak256(abi.encode(jobId, target, deadline, keeper, block.chainid));
        bytes32 ethSignedHash =
            keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", hash));
        // Use the private key corresponding to the TaskDispatcher (set in setUp)
        // 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
        uint256 pk = 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80;
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(pk, ethSignedHash);
        signature = abi.encodePacked(r, s, v);
    }

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

        vm.expectRevert(
            abi.encodeWithSelector(
                OwnableUpgradeable.OwnableUnauthorizedAccount.selector, randomUser
            )
        );
        vm.prank(randomUser);
        taskExecutionHub.addSpokes(newEids);
    }

    function test_ExecuteFunction() public {
        // Setup task dispatcher
        address taskDispatcherAddress = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266;
        uint256 taskDispatcherPrivateKey =
            0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80;

        vm.prank(owner);
        taskExecutionHub.setTaskDispatcher(taskDispatcherAddress);

        uint256 jobId = 1;
        uint256 ethAmount = 0.1 ether;
        address target = address(0x400);
        bytes memory data = abi.encodeWithSignature("doSomething()");
        uint256 value = 1 ether;
        uint256 deadline = block.timestamp + 1 hours;

        // Setup job and balance
        mockJobRegistry.setJobOwner(jobId, jobOwner);
        mockTriggerGasRegistry.setBalance(jobOwner, 1 ether);

        // Fund both the keeper and the contract
        vm.deal(keeper1, value);
        vm.deal(address(taskExecutionHub), value);

        // Create a mock contract that will always return success
        vm.etch(target, hex"600180600c6000396000f3006000fd"); // Simple bytecode that always returns true

        // Create signature
        bytes memory signature;
        (deadline, signature) = _dummySignatureParams(jobId, target, data, keeper1);

        vm.expectEmit(true, true, true, true);
        emit FunctionExecuted(keeper1, target, data, value);

        vm.prank(keeper1);
        taskExecutionHub.executeFunction{value: value}(
            jobId, ethAmount, target, data, deadline, signature
        );

        // Verify ETH balance was deducted
        uint256 finalBalance = mockTriggerGasRegistry.getBalance(jobOwner);
        assertEq(finalBalance, 1 ether - ethAmount);
    }

    function test_ExecuteFunction_OnlyKeeper() public {
        uint256 jobId = 1;
        uint256 ethAmount = 0.1 ether;
        address target = address(0x400);
        bytes memory data = abi.encodeWithSignature("doSomething()");
        uint256 deadline = block.timestamp + 1 hours;
        bytes memory signature = hex"";

        // Setup job and balance
        mockJobRegistry.setJobOwner(jobId, jobOwner);
        mockTriggerGasRegistry.setBalance(jobOwner, 1 ether);

        vm.expectRevert("Not a keeper");
        vm.prank(randomUser);
        taskExecutionHub.executeFunction(jobId, ethAmount, target, data, deadline, signature);
    }

    function test_ExecuteFunction_Failure() public {
        // Setup task dispatcher
        address taskDispatcherAddress = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266;
        uint256 taskDispatcherPrivateKey =
            0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80;

        vm.prank(owner);
        taskExecutionHub.setTaskDispatcher(taskDispatcherAddress);

        uint256 jobId = 1;
        uint256 ethAmount = 0.1 ether;
        address target = address(0x400);
        bytes memory data = abi.encodeWithSignature("doSomething()");
        uint256 deadline = block.timestamp + 1 hours;

        // Setup job and balance
        mockJobRegistry.setJobOwner(jobId, jobOwner);
        mockTriggerGasRegistry.setBalance(jobOwner, 1 ether);

        // Create a mock contract that will always revert
        vm.etch(target, hex"60006000fd"); // Simple bytecode that always reverts

        // Create signature
        bytes memory signature;
        (deadline, signature) = _dummySignatureParams(jobId, target, data, keeper1);

        vm.expectEmit(true, true, true, true);
        emit FunctionExecutionFailed(keeper1, target, data, 0, hex"");

        vm.prank(keeper1);
        taskExecutionHub.executeFunction(jobId, ethAmount, target, data, deadline, signature);

        // Verify ETH balance was STILL deducted despite execution failure
        uint256 finalBalance = mockTriggerGasRegistry.getBalance(jobOwner);
        assertEq(finalBalance, 1 ether - ethAmount);
    }

    function test_ExecuteFunction_JobNotFound() public {
        // Setup task dispatcher
        address taskDispatcherAddress = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266;
        uint256 taskDispatcherPrivateKey =
            0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80;

        vm.prank(owner);
        taskExecutionHub.setTaskDispatcher(taskDispatcherAddress);

        uint256 jobId = 999; // Non-existent job
        uint256 ethAmount = 0.1 ether;
        address target = address(0x400);
        bytes memory data = abi.encodeWithSignature("doSomething()");
        uint256 deadline = block.timestamp + 1 hours;

        // Create signature (even though job doesn't exist)
        bytes32 hash = keccak256(
            abi.encodePacked(jobId, target, keccak256(data), deadline, keeper1, block.chainid)
        );
        // Create signature
        bytes memory signature;
        (deadline, signature) = _dummySignatureParams(jobId, target, data, keeper1);

        vm.expectRevert(TaskExecutionHub.JobNotFound.selector);
        vm.prank(keeper1);
        taskExecutionHub.executeFunction(jobId, ethAmount, target, data, deadline, signature);
    }

    function test_ExecuteFunction_InsufficientETHBalance() public {
        // Setup task dispatcher
        address taskDispatcherAddress = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266;
        uint256 taskDispatcherPrivateKey =
            0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80;

        vm.prank(owner);
        taskExecutionHub.setTaskDispatcher(taskDispatcherAddress);

        uint256 jobId = 1;
        uint256 ethAmount = 1 ether; // More than available balance
        address target = address(0x400);
        bytes memory data = abi.encodeWithSignature("doSomething()");
        uint256 deadline = block.timestamp + 1 hours;

        // Setup job with insufficient balance
        mockJobRegistry.setJobOwner(jobId, jobOwner);
        mockTriggerGasRegistry.setBalance(jobOwner, 0.5 ether); // Less than required

        // Create signature
        // Create signature
        bytes memory signature;
        (deadline, signature) = _dummySignatureParams(jobId, target, data, keeper1);

        vm.expectRevert(bytes("Insufficient ETH balance"));
        vm.prank(keeper1);
        taskExecutionHub.executeFunction(jobId, ethAmount, target, data, deadline, signature);
    }

    function test_LzReceive_Register() public {
        // Simulate a message from AvsGovernanceLogic to register a new keeper
        bytes memory payload = abi.encode(TaskExecutionHub.ActionType.REGISTER, keeper2);

        // Create Origin struct directly with the 3 required fields
        Origin memory origin =
            Origin({srcEid: SRC_EID, sender: bytes32(uint256(uint160(address(0x1)))), nonce: 1});

        vm.expectEmit(true, false, false, false);
        emit KeeperRegistered(keeper2);

        // Call our exposed function to test _lzReceive
        taskExecutionHub.exposed_lzReceive(origin, bytes32(0), payload, address(0), bytes(""));

        // Verify keeper was registered
        assertTrue(taskExecutionHub.isKeeper(keeper2));
    }

    function test_LzReceive_Unregister() public {
        // First register keeper2
        bytes memory registerPayload = abi.encode(TaskExecutionHub.ActionType.REGISTER, keeper2);

        Origin memory origin =
            Origin({srcEid: SRC_EID, sender: bytes32(uint256(uint160(address(0x1)))), nonce: 1});

        taskExecutionHub.exposed_lzReceive(
            origin, bytes32(0), registerPayload, address(0), bytes("")
        );

        assertTrue(taskExecutionHub.isKeeper(keeper2));

        // Now unregister
        bytes memory unregisterPayload = abi.encode(TaskExecutionHub.ActionType.UNREGISTER, keeper2);

        vm.expectEmit(true, false, false, false);
        emit KeeperUnregistered(keeper2);

        taskExecutionHub.exposed_lzReceive(
            origin, bytes32(0), unregisterPayload, address(0), bytes("")
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

        vm.expectRevert(
            abi.encodeWithSelector(
                OwnableUpgradeable.OwnableUnauthorizedAccount.selector, randomUser
            )
        );
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
        uint32 newSrcEid = 10_122;
        bytes32 newAvsGovernance = bytes32(uint256(uint160(address(0x456))));

        vm.prank(owner);
        taskExecutionHub.setPeer(newSrcEid, newAvsGovernance);

        // We would need a getter function to verify this
        // Instead, we can test that setPeer is called correctly by looking at emitted events
        // or checking that subsequent operations with the new peer work as expected
    }

    function test_SetPeer_OnlyOwner() public {
        uint32 newSrcEid = 10_122;
        bytes32 newAvsGovernance = bytes32(uint256(uint160(address(0x456))));

        vm.expectRevert(
            abi.encodeWithSelector(
                OwnableUpgradeable.OwnableUnauthorizedAccount.selector, randomUser
            )
        );
        vm.prank(randomUser);
        taskExecutionHub.setPeer(newSrcEid, newAvsGovernance);
    }

    function test_ExecuteFunction_WithValidDispatcherSignature() public {
        // Setup task dispatcher
        address taskDispatcherAddress = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266;
        uint256 taskDispatcherPrivateKey =
            0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80;

        vm.prank(owner);
        taskExecutionHub.setTaskDispatcher(taskDispatcherAddress);

        // Setup job parameters
        uint256 jobId = 1;
        uint256 ethAmount = 0.1 ether;
        address target = address(0x400);
        bytes memory data = abi.encodeWithSignature("doSomething()");
        uint256 deadline = block.timestamp + 1 hours;

        // Setup job and balance
        mockJobRegistry.setJobOwner(jobId, jobOwner);
        mockTriggerGasRegistry.setBalance(jobOwner, 1 ether);

        // Create a mock contract that will always return success
        vm.etch(target, hex"600180600c6000396000f3006000fd");

        // Create signature hash (same as contract)
        bytes memory signature;
        (deadline, signature) = _dummySignatureParams(jobId, target, data, keeper1);

        console2.log("Task Dispatcher:", taskDispatcherAddress);
        console2.log("Keeper:", keeper1);
        console2.log("Target:", target);
        console2.log("JobId:", jobId);
        console2.log("Deadline:", deadline);
        console2.log("ChainId:", block.chainid);

        // Execute with valid signature
        vm.expectEmit(true, true, true, true);
        emit FunctionExecuted(keeper1, target, data, 0);

        vm.prank(keeper1);
        taskExecutionHub.executeFunction(jobId, ethAmount, target, data, deadline, signature);

        // Verify ETH balance was deducted
        uint256 finalBalance = mockTriggerGasRegistry.getBalance(jobOwner);
        assertEq(finalBalance, 1 ether - ethAmount);
    }

    function test_ExecuteFunction_ExpiredDeadline() public {
        // Setup task dispatcher
        address taskDispatcherAddress = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266;
        uint256 taskDispatcherPrivateKey =
            0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80;

        vm.prank(owner);
        taskExecutionHub.setTaskDispatcher(taskDispatcherAddress);

        // Setup job parameters with past deadline
        uint256 jobId = 1;
        uint256 ethAmount = 0.1 ether;
        address target = address(0x400);
        bytes memory data = abi.encodeWithSignature("doSomething()");
        uint256 deadline = block.timestamp - 1; // Past deadline

        // Setup job and balance
        mockJobRegistry.setJobOwner(jobId, jobOwner);
        mockTriggerGasRegistry.setBalance(jobOwner, 1 ether);

        // Create signature
        bytes32 hash = keccak256(
            abi.encodePacked(jobId, target, keccak256(data), deadline, keeper1, block.chainid)
        );
        bytes32 ethSignedHash =
            keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", hash));
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(taskDispatcherPrivateKey, ethSignedHash);
        bytes memory signature = abi.encodePacked(r, s, v);

        // Should revert with expired signature
        vm.expectRevert(TaskExecutionHub.SignatureExpired.selector);
        vm.prank(keeper1);
        taskExecutionHub.executeFunction(jobId, ethAmount, target, data, deadline, signature);
    }

    function test_ExecuteFunction_InvalidSignature() public {
        // Setup task dispatcher
        address taskDispatcherAddress = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266;

        vm.prank(owner);
        taskExecutionHub.setTaskDispatcher(taskDispatcherAddress);

        // Setup job parameters
        uint256 jobId = 1;
        uint256 ethAmount = 0.1 ether;
        address target = address(0x400);
        bytes memory data = abi.encodeWithSignature("doSomething()");
        uint256 deadline = block.timestamp + 1 hours;

        // Setup job and balance
        mockJobRegistry.setJobOwner(jobId, jobOwner);
        mockTriggerGasRegistry.setBalance(jobOwner, 1 ether);

        // Create signature with WRONG private key
        uint256 wrongPrivateKey = 0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef;
        bytes32 hash = keccak256(
            abi.encodePacked(jobId, target, keccak256(data), deadline, keeper1, block.chainid)
        );
        bytes32 ethSignedHash =
            keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", hash));
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(wrongPrivateKey, ethSignedHash);
        bytes memory signature = abi.encodePacked(r, s, v);

        // Should revert with invalid signature
        vm.expectRevert(TaskExecutionHub.InvalidSignature.selector);
        vm.prank(keeper1);
        taskExecutionHub.executeFunction(jobId, ethAmount, target, data, deadline, signature);
    }

    function test_ExecuteFunction_WrongKeeper() public {
        // Setup task dispatcher
        address taskDispatcherAddress = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266;
        uint256 taskDispatcherPrivateKey =
            0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80;

        vm.prank(owner);
        taskExecutionHub.setTaskDispatcher(taskDispatcherAddress);

        // Setup job parameters
        uint256 jobId = 1;
        uint256 ethAmount = 0.1 ether;
        address target = address(0x400);
        bytes memory data = abi.encodeWithSignature("doSomething()");
        uint256 deadline = block.timestamp + 1 hours;

        // Setup job and balance
        mockJobRegistry.setJobOwner(jobId, jobOwner);
        mockTriggerGasRegistry.setBalance(jobOwner, 1 ether);

        // Create signature for keeper1
        bytes32 hash = keccak256(
            abi.encodePacked(
                jobId,
                target,
                keccak256(data),
                deadline,
                keeper1, // Signature is for keeper1
                block.chainid
            )
        );
        bytes32 ethSignedHash =
            keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", hash));
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(taskDispatcherPrivateKey, ethSignedHash);
        bytes memory signature = abi.encodePacked(r, s, v);

        // Register keeper2
        vm.prank(owner);
        taskExecutionHub.addKeeper(keeper2);

        // Try to execute with keeper2, but signature was for keeper1
        vm.expectRevert(TaskExecutionHub.InvalidSignature.selector);
        vm.prank(keeper2);
        taskExecutionHub.executeFunction(jobId, ethAmount, target, data, deadline, signature);
    }

    function test_ExecuteFunction_TaskDispatcherNotSet() public {
        // DO NOT set task dispatcher

        // Setup job parameters
        uint256 jobId = 1;
        uint256 ethAmount = 0.1 ether;
        address target = address(0x400);
        bytes memory data = abi.encodeWithSignature("doSomething()");
        uint256 deadline = block.timestamp + 1 hours;
        bytes memory signature = hex""; // Any signature

        // Setup job and balance
        mockJobRegistry.setJobOwner(jobId, jobOwner);
        mockTriggerGasRegistry.setBalance(jobOwner, 1 ether);

        // Should revert because taskDispatcher is not set
        vm.expectRevert(TaskExecutionHub.TaskDispatcherNotSet.selector);
        vm.prank(keeper1);
        taskExecutionHub.executeFunction(jobId, ethAmount, target, data, deadline, signature);
    }

    function test_ExecuteFunction_TamperedCalldata() public {
        // Setup task dispatcher
        address taskDispatcherAddress = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266;
        uint256 taskDispatcherPrivateKey =
            0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80;

        vm.prank(owner);
        taskExecutionHub.setTaskDispatcher(taskDispatcherAddress);

        // Setup job parameters
        uint256 jobId = 1;
        uint256 ethAmount = 0.1 ether;
        address target = address(0x400);
        bytes memory originalData = abi.encodeWithSignature("doSomething()");
        bytes memory tamperedData = abi.encodeWithSignature("doSomethingElse()");
        uint256 deadline = block.timestamp + 1 hours;

        // Setup job and balance
        mockJobRegistry.setJobOwner(jobId, jobOwner);
        mockTriggerGasRegistry.setBalance(jobOwner, 1 ether);

        // Create signature for original data
        bytes32 hash = keccak256(
            abi.encodePacked(
                jobId,
                target,
                keccak256(originalData), // Signature for original data
                deadline,
                keeper1,
                block.chainid
            )
        );
        bytes32 ethSignedHash =
            keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", hash));
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(taskDispatcherPrivateKey, ethSignedHash);
        bytes memory signature = abi.encodePacked(r, s, v);

        // Try to execute with tampered data
        vm.expectRevert(TaskExecutionHub.InvalidSignature.selector);
        vm.prank(keeper1);
        taskExecutionHub.executeFunction(
            jobId,
            ethAmount,
            target,
            tamperedData, // Using tampered data
            deadline,
            signature
        );
    }

    function test_SetTaskDispatcher_OnlyOwner() public {
        address newDispatcher = address(0x999);

        vm.expectRevert(
            abi.encodeWithSelector(
                OwnableUpgradeable.OwnableUnauthorizedAccount.selector, randomUser
            )
        );
        vm.prank(randomUser);
        taskExecutionHub.setTaskDispatcher(newDispatcher);
    }

    function test_SetTaskDispatcher_ZeroAddress() public {
        vm.expectRevert("Invalid taskDispatcher address");
        vm.prank(owner);
        taskExecutionHub.setTaskDispatcher(address(0));
    }

    function test_SetTaskDispatcher_Success() public {
        address newDispatcher = address(0x999);

        vm.expectEmit(true, true, false, false);
        emit TaskDispatcherUpdated(address(0), newDispatcher);

        vm.prank(owner);
        taskExecutionHub.setTaskDispatcher(newDispatcher);

        assertEq(taskExecutionHub.taskDispatcher(), newDispatcher);
    }

    // =========================================================================
    // ============ TriggerXSafeModule JobOwner Validation Tests ===============
    // =========================================================================

    function test_SafeModule_ValidJobOwner() public {
        // Setup task dispatcher
        address taskDispatcherAddress = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266;
        address safeModuleAddr = address(0x500);

        vm.startPrank(owner);
        taskExecutionHub.setTaskDispatcher(taskDispatcherAddress);
        taskExecutionHub.setTriggerXSafeModule(safeModuleAddr);
        vm.stopPrank();

        uint256 jobId = 1;
        uint256 ethAmount = 0.1 ether;

        // Setup job and balance
        mockJobRegistry.setJobOwner(jobId, jobOwner);
        mockTriggerGasRegistry.setBalance(jobOwner, 1 ether);

        // Create calldata with correct jobOwner as last param
        // execJobFromHub(address safeAddress, address actionTarget, uint256 actionValue, bytes actionData, uint8 operation, address jobOwner)
        bytes memory data = abi.encodeWithSelector(
            bytes4(keccak256("execJobFromHub(address,address,uint256,bytes,uint8,address)")),
            address(0x600), // safeAddress
            address(0x700), // actionTarget
            0, // actionValue
            "", // actionData
            0, // operation
            jobOwner // jobOwner - matches actual job owner (LAST param)
        );

        // Create a mock contract at safeModuleAddr that succeeds
        vm.etch(safeModuleAddr, hex"600180600c6000396000f3006000fd");

        // Create signature
        (uint256 deadline, bytes memory signature) =
            _dummySignatureParams(jobId, safeModuleAddr, data, keeper1);

        // Should succeed because jobOwner in calldata matches
        vm.prank(keeper1);
        taskExecutionHub.executeFunction(
            jobId, ethAmount, safeModuleAddr, data, deadline, signature
        );

        // Verify ETH balance was deducted
        uint256 finalBalance = mockTriggerGasRegistry.getBalance(jobOwner);
        assertEq(finalBalance, 1 ether - ethAmount);
    }

    function test_SafeModule_MismatchedJobOwner_Reverts() public {
        // Setup task dispatcher
        address taskDispatcherAddress = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266;
        address safeModuleAddr = address(0x500);
        address wrongJobOwner = address(0xDEAD);

        vm.startPrank(owner);
        taskExecutionHub.setTaskDispatcher(taskDispatcherAddress);
        taskExecutionHub.setTriggerXSafeModule(safeModuleAddr);
        vm.stopPrank();

        uint256 jobId = 1;
        uint256 ethAmount = 0.1 ether;

        // Setup job with jobOwner
        mockJobRegistry.setJobOwner(jobId, jobOwner);
        mockTriggerGasRegistry.setBalance(jobOwner, 1 ether);

        // Create calldata with WRONG jobOwner as last param
        bytes memory data = abi.encodeWithSelector(
            bytes4(keccak256("execJobFromHub(address,address,uint256,bytes,uint8,address)")),
            address(0x600), // safeAddress
            address(0x700), // actionTarget
            0, // actionValue
            "", // actionData
            0, // operation
            wrongJobOwner // WRONG jobOwner (LAST param)
        );

        // Create signature
        (uint256 deadline, bytes memory signature) =
            _dummySignatureParams(jobId, safeModuleAddr, data, keeper1);

        // Should revert with JobOwnerMismatch
        vm.expectRevert(
            abi.encodeWithSelector(
                TaskExecutionHub.JobOwnerMismatch.selector, wrongJobOwner, jobOwner
            )
        );
        vm.prank(keeper1);
        taskExecutionHub.executeFunction(
            jobId, ethAmount, safeModuleAddr, data, deadline, signature
        );
    }

    function test_SafeModule_NonSafeModuleTarget_SkipsValidation() public {
        // Setup task dispatcher
        address taskDispatcherAddress = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266;
        address safeModuleAddr = address(0x500);
        address regularTarget = address(0x400);

        vm.startPrank(owner);
        taskExecutionHub.setTaskDispatcher(taskDispatcherAddress);
        taskExecutionHub.setTriggerXSafeModule(safeModuleAddr);
        vm.stopPrank();

        uint256 jobId = 1;
        uint256 ethAmount = 0.1 ether;

        // Setup job and balance
        mockJobRegistry.setJobOwner(jobId, jobOwner);
        mockTriggerGasRegistry.setBalance(jobOwner, 1 ether);

        // Create regular calldata (not SafeModule format)
        bytes memory data = abi.encodeWithSignature("doSomething()");

        // Create a mock contract that succeeds
        vm.etch(regularTarget, hex"600180600c6000396000f3006000fd");

        // Create signature
        (uint256 deadline, bytes memory signature) =
            _dummySignatureParams(jobId, regularTarget, data, keeper1);

        // Should succeed - validation is skipped for non-SafeModule targets
        vm.prank(keeper1);
        taskExecutionHub.executeFunction(jobId, ethAmount, regularTarget, data, deadline, signature);

        // Verify ETH balance was deducted
        uint256 finalBalance = mockTriggerGasRegistry.getBalance(jobOwner);
        assertEq(finalBalance, 1 ether - ethAmount);
    }

    function test_SafeModule_InvalidCalldata_Reverts() public {
        // Setup task dispatcher
        address taskDispatcherAddress = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266;
        address safeModuleAddr = address(0x500);

        vm.startPrank(owner);
        taskExecutionHub.setTaskDispatcher(taskDispatcherAddress);
        taskExecutionHub.setTriggerXSafeModule(safeModuleAddr);
        vm.stopPrank();

        uint256 jobId = 1;
        uint256 ethAmount = 0.1 ether;

        mockJobRegistry.setJobOwner(jobId, jobOwner);
        mockTriggerGasRegistry.setBalance(jobOwner, 1 ether);

        // Create too-short calldata (less than 36 bytes)
        bytes memory data = hex"12345678"; // Only 4 bytes (selector only)

        // Create signature
        (uint256 deadline, bytes memory signature) =
            _dummySignatureParams(jobId, safeModuleAddr, data, keeper1);

        // Should revert with InvalidSafeModuleCalldata
        vm.expectRevert(TaskExecutionHub.InvalidSafeModuleCalldata.selector);
        vm.prank(keeper1);
        taskExecutionHub.executeFunction(
            jobId, ethAmount, safeModuleAddr, data, deadline, signature
        );
    }

    function test_SetTriggerXSafeModule() public {
        address newModule = address(0x999);

        vm.prank(owner);
        taskExecutionHub.setTriggerXSafeModule(newModule);

        assertEq(taskExecutionHub.triggerXSafeModule(), newModule);
    }
}
