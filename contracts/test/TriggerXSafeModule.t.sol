// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

import "forge-std/Test.sol";

import "../src/TriggerXSafeModule.sol";

// Minimal mock Safe that allows module exec
contract MockSafe is IGnosisSafe {
    address public lastTo;
    uint256 public lastValue;
    bytes public lastData;
    uint8 public lastOp;
    mapping(address => bool) public owners;

    constructor() {
        // Default owner for tests
        owners[address(0xCAFE)] = true;
    }

    function isOwner(
        address owner
    ) external view override returns (bool) {
        return owners[owner];
    }

    function setOwner(
        address owner,
        bool status
    ) external {
        owners[owner] = status;
    }

    // execute: simply call target with provided data as this contract
    function execTransactionFromModule(
        address to,
        uint256 value,
        bytes calldata data,
        uint8 operation
    ) external override returns (bool success) {
        lastTo = to;
        lastValue = value;
        lastData = data;
        lastOp = operation;

        // perform call
        (bool ok,) = to.call{value: value}(data);
        return ok;
    }

    receive() external payable {}
}

// Simple target contract the Safe will call
contract DummyTarget {
    event ActionCalled(address caller, uint256 value, bytes data);
    uint256 public x;

    function doSomething(
        uint256 v
    ) external payable {
        x = v;
        emit ActionCalled(msg.sender, msg.value, abi.encode(v));
    }
}

// Mock Hub - only to set as hub address and call module via test (we'll impersonate hub in test)
contract MockHub {}

contract TriggerXSafeModuleTest is Test {
    MockSafe safe;
    DummyTarget target;
    MockHub hub;
    TriggerXSafeModule moduleContract;

    address keeper = address(0xBEEF);
    address jobOwner = address(0xCAFE);

    function setUp() public {
        safe = new MockSafe();
        target = new DummyTarget();
        hub = new MockHub();

        // TriggerXSafeModule now only takes hub address (no jobRegistry)
        moduleContract = new TriggerXSafeModule(address(hub));

        vm.deal(address(safe), 1 ether);
        vm.deal(address(this), 1 ether);
    }

    function testHappyPathExec() public {
        // prepare actionData
        bytes memory actionData = abi.encodeWithSelector(DummyTarget.doSomething.selector, 123);
        uint8 op = 0;

        // impersonate hub and call execJobFromHub with jobOwner (last param)
        vm.prank(address(hub));
        bool ok = moduleContract.execJobFromHub(
            address(safe),
            address(target),
            uint256(0),
            actionData,
            op,
            jobOwner // jobOwner is now LAST (validated by TaskExecutionHub)
        );

        assertTrue(ok, "module should return true");
        assertEq(target.x(), 123);
    }

    function testSafeNotOwnedByJobOwner() public {
        bytes memory actionData = abi.encodeWithSelector(DummyTarget.doSomething.selector, 999);
        uint8 op = 0;

        // Use a different jobOwner that doesn't own the Safe
        address wrongJobOwner = address(0xDEAD);

        vm.prank(address(hub));
        vm.expectRevert(
            abi.encodeWithSelector(
                TriggerXSafeModule.SafeNotOwnedByJobOwner.selector, address(safe), wrongJobOwner
            )
        );
        moduleContract.execJobFromHub(
            address(safe),
            address(target),
            uint256(0),
            actionData,
            op,
            wrongJobOwner // wrongJobOwner is now LAST
        );
    }

    function testMultipleCallsAllowed() public {
        // The contract doesn't have replay protection - multiple identical calls are allowed
        // This is by design since the hub is trusted and manages replay protection at a higher level
        bytes memory actionData = abi.encodeWithSelector(DummyTarget.doSomething.selector, 5);
        uint8 op = 0;

        vm.prank(address(hub));
        bool ok1 = moduleContract.execJobFromHub(
            address(safe), address(target), uint256(0), actionData, op, jobOwner
        );
        assertTrue(ok1);

        // second call should also succeed (no replay protection in module)
        vm.prank(address(hub));
        bool ok2 = moduleContract.execJobFromHub(
            address(safe), address(target), uint256(0), actionData, op, jobOwner
        );
        assertTrue(ok2);

        // Verify both calls executed
        assertEq(target.x(), 5);
    }

    function testOnlyHubCanCall() public {
        bytes memory actionData = abi.encodeWithSelector(DummyTarget.doSomething.selector, 123);

        // Try calling from non-hub address
        vm.prank(address(0x1234));
        vm.expectRevert(TriggerXSafeModule.NotTaskExecutionHub.selector);
        moduleContract.execJobFromHub(
            address(safe), address(target), uint256(0), actionData, 0, jobOwner
        );
    }
}
