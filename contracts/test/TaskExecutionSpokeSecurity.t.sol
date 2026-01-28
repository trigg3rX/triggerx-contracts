// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Test} from "forge-std/Test.sol";
import {TaskExecutionSpoke} from "../src/lz/TaskExecutionSpoke.sol";
import {MockEndpoint} from "./mocks/MockEndpoint.sol";
import {OwnableUpgradeable} from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import {Origin} from "@layerzero-v2/oapp/contracts/oapp/OApp.sol";
import {MockJobRegistry} from "./mocks/MockJobRegistry.sol";
import {MockTriggerGasRegistry} from "./mocks/MockTriggerGasRegistry.sol";
import {ERC1967Proxy} from "@openzeppelin-contracts/contracts/proxy/ERC1967/ERC1967Proxy.sol";

contract MaliciousSpokeTarget {
    bool public shouldRevert;
    bool public shouldConsumeMuchGas;
    TaskExecutionSpoke public spoke;

    constructor(address _spoke) {
        spoke = TaskExecutionSpoke(_spoke);
    }

    function setRevert(bool _shouldRevert) external {
        shouldRevert = _shouldRevert;
    }

    function setConsumeMuchGas(bool _shouldConsume) external {
        shouldConsumeMuchGas = _shouldConsume;
    }

    function maliciousFunction() external payable {
        if (shouldRevert) {
            revert("Malicious revert");
        }

        if (shouldConsumeMuchGas) {
            // Consume much gas
            for (uint256 i = 0; i < 10000; i++) {
                keccak256(abi.encode(i, block.timestamp));
            }
        }
    }

    receive() external payable {
        if (shouldRevert) {
            revert("Rejecting ETH");
        }
    }
}

contract TaskExecutionSpokeSecurityTest is Test {
    TaskExecutionSpoke public taskExecutionSpoke;
    MockEndpoint public mockEndpoint;
    address public owner = address(0x1);
    address public attacker = address(0x666);
    address public keeper1 = address(0x100);
    address public keeper2 = address(0x101);
    address public taskDispatcher = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266;
    uint256 public taskDispatcherPK =
        0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80;

    uint32 public constant HUB_EID = 10101;

    // Helper to create signature params
    function _dummySignatureParams(
        uint256 jobId,
        address target,
        bytes memory data,
        address keeper
    ) internal view returns (uint256 deadline, bytes memory signature) {
        deadline = block.timestamp + 1 hours;
        bytes32 hash = keccak256(
            abi.encode(jobId, target, deadline, keeper, block.chainid)
        );
        bytes32 ethSignedHash = keccak256(
            abi.encodePacked("\x19Ethereum Signed Message:\n32", hash)
        );
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(
            taskDispatcherPK,
            ethSignedHash
        );
        signature = abi.encodePacked(r, s, v);
    }

    function setUp() public {
        vm.startPrank(owner);

        mockEndpoint = new MockEndpoint();

        address[] memory initialKeepers = new address[](2);
        initialKeepers[0] = keeper1;
        initialKeepers[1] = keeper2;

        // Deploy implementation
        TaskExecutionSpoke implementation = new TaskExecutionSpoke(
            address(mockEndpoint),
            owner
        );

        // Deploy mock registries
        MockJobRegistry mockJobRegistry = new MockJobRegistry();
        MockTriggerGasRegistry mockTriggerGasRegistry = new MockTriggerGasRegistry();

        // Deploy proxy with initialization
        ERC1967Proxy proxy = new ERC1967Proxy(
            address(implementation),
            abi.encodeWithSelector(
                TaskExecutionSpoke.initialize.selector,
                owner,
                HUB_EID,
                initialKeepers,
                address(mockJobRegistry),
                address(mockTriggerGasRegistry)
            )
        );

        taskExecutionSpoke = TaskExecutionSpoke(payable(address(proxy)));

        // Set task dispatcher for signature verification
        taskExecutionSpoke.setTaskDispatcher(taskDispatcher);

        vm.stopPrank();

        vm.deal(keeper1, 10 ether);
        vm.deal(keeper2, 10 ether);
        vm.deal(attacker, 10 ether);
    }

    // ==================== ACCESS CONTROL TESTS ====================

    function test_Security_OnlyKeeperCanExecuteFunction() public {
        uint256 jobId = 1;
        uint256 tgAmount = 100;
        address target = address(0x999);
        bytes memory data = abi.encodeWithSignature("someFunction()");

        // Setup job and balance
        MockJobRegistry mockJobRegistry = MockJobRegistry(
            address(taskExecutionSpoke.jobRegistry())
        );
        MockTriggerGasRegistry mockTriggerGasRegistry = MockTriggerGasRegistry(
            address(taskExecutionSpoke.triggerGasRegistry())
        );
        mockJobRegistry.setJobOwner(jobId, address(0x300));
        mockTriggerGasRegistry.setBalance(address(0x300), 1000);

        (uint256 deadline, bytes memory signature) = _dummySignatureParams(
            jobId,
            target,
            data,
            attacker
        );
        vm.expectRevert("Spoke: Keeper not registered");
        vm.prank(attacker);
        taskExecutionSpoke.executeFunction(
            jobId,
            tgAmount,
            target,
            data,
            deadline,
            signature
        );
    }

    function test_Security_InitialKeepersAreSetCorrectly() public view {
        assertTrue(taskExecutionSpoke.isKeeper(keeper1));
        assertTrue(taskExecutionSpoke.isKeeper(keeper2));
        assertFalse(taskExecutionSpoke.isKeeper(attacker));
    }

    // ==================== CROSS-CHAIN MESSAGING SECURITY TESTS ====================

    function test_Security_OnlyValidActionTypesAccepted() public {
        // Set up peer relationship first
        vm.prank(owner);
        taskExecutionSpoke.setPeer(
            HUB_EID,
            bytes32(uint256(uint160(address(this))))
        );

        // Test invalid action type
        bytes memory invalidPayload = abi.encode(uint8(99), attacker); // Invalid action type

        vm.expectRevert();
        vm.prank(address(mockEndpoint));
        taskExecutionSpoke.lzReceive(
            Origin({
                srcEid: HUB_EID,
                sender: bytes32(uint256(uint160(address(this)))),
                nonce: uint64(0)
            }),
            bytes32(0),
            invalidPayload,
            address(0),
            bytes("")
        );
    }

    function test_Security_MalformedPayloadHandling() public {
        // Test malformed payload
        bytes memory malformedPayload = abi.encode("invalid", 123, true);

        vm.expectRevert();
        vm.prank(address(mockEndpoint));
        taskExecutionSpoke.lzReceive(
            Origin({srcEid: uint32(0), sender: bytes32(0), nonce: uint64(0)}),
            bytes32(0),
            malformedPayload,
            address(0),
            bytes("")
        );
    }

    // ==================== FUNCTION EXECUTION SECURITY TESTS ====================

    function test_Security_FunctionExecutionWithMaliciousTarget() public {
        uint256 jobId = 1;
        uint256 tgAmount = 100;
        MaliciousSpokeTarget maliciousTarget = new MaliciousSpokeTarget(
            address(taskExecutionSpoke)
        );
        maliciousTarget.setRevert(true);

        // Setup job and balance
        MockJobRegistry mockJobRegistry = MockJobRegistry(
            address(taskExecutionSpoke.jobRegistry())
        );
        MockTriggerGasRegistry mockTriggerGasRegistry = MockTriggerGasRegistry(
            address(taskExecutionSpoke.triggerGasRegistry())
        );
        mockJobRegistry.setJobOwner(jobId, address(0x300));
        mockTriggerGasRegistry.setBalance(address(0x300), 1000);

        bytes memory data = abi.encodeWithSelector(
            MaliciousSpokeTarget.maliciousFunction.selector
        );

        // Should NOT revert, but emit FunctionExecutionFailed
        (uint256 deadline, bytes memory signature) = _dummySignatureParams(
            jobId,
            address(maliciousTarget),
            data,
            keeper1
        );
        vm.prank(keeper1);
        taskExecutionSpoke.executeFunction(
            jobId,
            tgAmount,
            address(maliciousTarget),
            data,
            deadline,
            signature
        );

        // Verify ETH balance was deducted
        assertEq(mockTriggerGasRegistry.getBalance(address(0x300)), 900);
    }

    function test_Security_CannotExecuteOnSpokeItself() public {
        uint256 jobId = 1;
        uint256 tgAmount = 100;
        // Try to call sensitive functions on the TaskExecutionSpoke itself
        bytes memory data = abi.encodeWithSelector(
            TaskExecutionSpoke.executeFunction.selector,
            address(0),
            ""
        );

        // Setup job and balance
        MockJobRegistry mockJobRegistry = MockJobRegistry(
            address(taskExecutionSpoke.jobRegistry())
        );
        MockTriggerGasRegistry mockTriggerGasRegistry = MockTriggerGasRegistry(
            address(taskExecutionSpoke.triggerGasRegistry())
        );
        mockJobRegistry.setJobOwner(jobId, address(0x300));
        mockTriggerGasRegistry.setBalance(address(0x300), 1000);

        // This should fail gracefully (emit event)
        (uint256 deadline, bytes memory signature) = _dummySignatureParams(
            jobId,
            address(taskExecutionSpoke),
            data,
            keeper1
        );
        vm.prank(keeper1);
        taskExecutionSpoke.executeFunction(
            jobId,
            tgAmount,
            address(taskExecutionSpoke),
            data,
            deadline,
            signature
        );

        // Verify ETH balance was deducted
        assertEq(mockTriggerGasRegistry.getBalance(address(0x300)), 900);
    }

    // ==================== KEEPER MANAGEMENT SECURITY TESTS ====================

    function test_Security_KeeperRegistrationSecurity() public {
        // Set up peer relationship first
        vm.prank(owner);
        taskExecutionSpoke.setPeer(
            HUB_EID,
            bytes32(uint256(uint160(address(this))))
        );

        // Initially, attacker is not a keeper
        assertFalse(taskExecutionSpoke.isKeeper(attacker));

        // Register attacker via LayerZero message
        bytes memory payload = abi.encode(
            TaskExecutionSpoke.ActionType.REGISTER,
            attacker
        );

        vm.prank(address(mockEndpoint));
        taskExecutionSpoke.lzReceive(
            Origin({
                srcEid: HUB_EID,
                sender: bytes32(uint256(uint160(address(this)))),
                nonce: uint64(0)
            }),
            bytes32(0),
            payload,
            address(0),
            bytes("")
        );

        // Now attacker should be a keeper
        assertTrue(taskExecutionSpoke.isKeeper(attacker));
    }

    function test_Security_KeeperUnregistrationSecurity() public {
        address testKeeper = address(0x999);

        // Set up peer relationship first
        vm.prank(owner);
        taskExecutionSpoke.setPeer(
            HUB_EID,
            bytes32(uint256(uint160(address(this))))
        );

        // Register a keeper first
        bytes memory registerPayload = abi.encode(
            TaskExecutionSpoke.ActionType.REGISTER,
            testKeeper
        );

        vm.prank(address(mockEndpoint));
        taskExecutionSpoke.lzReceive(
            Origin({
                srcEid: HUB_EID,
                sender: bytes32(uint256(uint160(address(this)))),
                nonce: uint64(0)
            }),
            bytes32(0),
            registerPayload,
            address(0),
            bytes("")
        );

        assertTrue(taskExecutionSpoke.isKeeper(testKeeper));

        // Unregister the keeper
        bytes memory unregisterPayload = abi.encode(
            TaskExecutionSpoke.ActionType.UNREGISTER,
            testKeeper
        );

        vm.prank(address(mockEndpoint));
        taskExecutionSpoke.lzReceive(
            Origin({
                srcEid: HUB_EID,
                sender: bytes32(uint256(uint160(address(this)))),
                nonce: uint64(0)
            }),
            bytes32(0),
            unregisterPayload,
            address(0),
            bytes("")
        );

        assertFalse(taskExecutionSpoke.isKeeper(testKeeper));
    }

    // ==================== HELPER FUNCTIONS ====================

    function dummyFunction() external pure returns (bool) {
        return true;
    }

    receive() external payable {
        // Allow contract to receive ETH
    }
}
