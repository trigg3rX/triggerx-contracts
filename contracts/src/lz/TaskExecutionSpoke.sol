// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Initializable} from "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import {
    ReentrancyGuardUpgradeable
} from "@openzeppelin-upgrades/contracts/utils/ReentrancyGuardUpgradeable.sol";
import {UUPSUpgradeable} from "@openzeppelin-upgrades/contracts/proxy/utils/UUPSUpgradeable.sol";
import {OApp, Origin} from "@layerzero-v2/oapp/contracts/oapp/OApp.sol";
import {Ownable} from "@openzeppelin-contracts/contracts/access/Ownable.sol";
import {ECDSA} from "@openzeppelin-contracts/contracts/utils/cryptography/ECDSA.sol";
import {
    MessageHashUtils
} from "@openzeppelin-contracts/contracts/utils/cryptography/MessageHashUtils.sol";

interface IJobRegistry {
    function getJobOwner(
        uint256 jobId
    ) external view returns (address);
}

interface ITriggerGasRegistry {
    function deductETHBalance(
        address user,
        uint256 ethAmount
    ) external;
}

/**
 * @title TaskExecutionSpoke
 * @notice A LayerZero-enabled contract that acts as a spoke in the keeper network
 * @dev This contract receives keeper registration updates from the hub and executes functions on the respective L2 chain
 */
contract TaskExecutionSpoke is Initializable, OApp, UUPSUpgradeable, ReentrancyGuardUpgradeable {
    /// @notice Constructor only runs on the implementation contract. It passes minimal arguments to
    ///         the OApp constructor (so the byte-code is valid) and immediately disables further
    ///         initializers to protect the logic contract.
    constructor(
        address _endpoint,
        address _delegate
    ) OApp(_endpoint, _delegate) Ownable(_delegate) {
        _disableInitializers();
    }

    // ----------------------------------
    // --------  State Variables  -------
    // ----------------------------------

    mapping(address => bool) public isKeeper;

    IJobRegistry public jobRegistry;
    ITriggerGasRegistry public triggerGasRegistry;

    /// @notice The authorized taskDispatcher address for signature verification
    address public taskDispatcher;
    address public triggerXSafeModule;

    enum ActionType {
        REGISTER,
        UNREGISTER
    }

    // ----------------------------------
    // -------------  Events ------------
    // ----------------------------------

    event KeeperUpdated(ActionType action, address keeper);
    event FunctionExecuted(
        address indexed keeper, address indexed target, bytes data, uint256 value
    );
    event FunctionExecutionFailed(
        address indexed keeper, address indexed target, bytes data, uint256 value, bytes result
    );
    event TaskDispatcherUpdated(
        address indexed oldTaskDispatcher, address indexed newTaskDispatcher
    );
    event TriggerXSafeModuleUpdated(address indexed oldModule, address indexed newModule);

    // Custom errors for gas optimization
    error SignatureExpired();
    error JobNotFound();
    error TaskDispatcherNotSet();
    error InvalidSignature();
    error InvalidSafeModuleCalldata();
    error JobOwnerMismatch(address passed, address expected);
    error WrongChain();

    // ----------------------------------
    // ------------  Modifiers ----------
    // ----------------------------------

    modifier onlyKeeper() {
        require(isKeeper[msg.sender], "Spoke: Keeper not registered");
        _;
    }

    // ---------------------------------------------------------------------
    // Initialization (replaces constructor)
    // ---------------------------------------------------------------------

    /**
     * @dev Initialize function (to be called via proxy).
     * @param _ownerAddress Contract owner.
     * @param _hubEid The hub chain endpoint ID.
     * @param _initialKeepers Array of initial keeper addresses.
     * @param _jobRegistryAddress The address of the job registry contract.
     * @param _triggerGasRegistryAddress The address of the trigger gas registry contract.
     */
    function initialize(
        address _ownerAddress,
        uint32 _hubEid,
        address[] calldata _initialKeepers,
        address _jobRegistryAddress,
        address _triggerGasRegistryAddress
    ) external initializer {
        // Init OZ upgradeable helpers
        __ReentrancyGuard_init();
        __UUPSUpgradeable_init();

        // Set peer connection to hub
        _setPeer(_hubEid, bytes32(uint256(uint160(address(this)))));

        // Initialize keepers
        for (uint256 i = 0; i < _initialKeepers.length; i++) {
            isKeeper[_initialKeepers[i]] = true;
            emit KeeperUpdated(ActionType.REGISTER, _initialKeepers[i]);
        }

        jobRegistry = IJobRegistry(_jobRegistryAddress);
        triggerGasRegistry = ITriggerGasRegistry(_triggerGasRegistryAddress);

        // Delegate & ownership wiring on the LayerZero endpoint
        endpoint.setDelegate(_ownerAddress);
        _transferOwnership(_ownerAddress);
    }

    // ---------------------------------------------------------------------
    // -----------------------    Main Logic     ---------------------------
    // ---------------------------------------------------------------------

    /**
     * @notice Execute a function call with taskDispatcher signature verification
     * @param jobId The ID of the job to execute
     * @param ethAmount The amount of ETH to deduct from the job owner
     * @param target The address of the target contract
     * @param data The calldata for the function call
     * @param deadline The signature expiration timestamp
     * @param signature The taskDispatcher's signature authorizing this execution
     */
    function executeFunction(
        uint256 jobId,
        uint256 ethAmount,
        address target,
        bytes calldata data,
        uint256 deadline,
        bytes calldata signature
    ) external payable onlyKeeper nonReentrant {
        if (block.timestamp > deadline) revert SignatureExpired();

        _verifyTaskDispatcherSignature(jobId, target, deadline, signature);
        address jobOwner = jobRegistry.getJobOwner(jobId);
        if (jobOwner == address(0)) revert JobNotFound();

        triggerGasRegistry.deductETHBalance(jobOwner, ethAmount);

        address _safeModule = triggerXSafeModule;
        if (target == _safeModule && _safeModule != address(0)) {
            _validateSafeModuleCalldata(data, jobOwner);
        }

        _executeFunction(target, data);
    }

    /**
     * @notice Verifies the taskDispatcher signature
     * @dev Security-critical: Ensures only authorized task dispatcher can approve executions
     */
    function _verifyTaskDispatcherSignature(
        uint256 jobId,
        address target,
        uint256 deadline,
        bytes calldata signature
    ) internal view {
        if (taskDispatcher == address(0)) revert TaskDispatcherNotSet();
        bytes32 hash = keccak256(abi.encode(jobId, target, deadline, msg.sender, block.chainid));
        bytes32 ethSignedHash = MessageHashUtils.toEthSignedMessageHash(hash);
        if (ECDSA.recover(ethSignedHash, signature) != taskDispatcher) {
            revert InvalidSignature();
        }
    }

    /**
     * @notice Validates jobOwner in Safe module calldata matches the actual job owner
     * @dev Security-critical: Prevents unauthorized Safe wallet execution
     * @dev Calldata format: selector (4) + safeAddress (32) + actionTarget (32) +
     *      actionValue (32) + actionData offset (32) + operation (32) + jobOwner (32)
     *      jobOwner is at offset 4 + 5*32 = 164
     */
    function _validateSafeModuleCalldata(
        bytes calldata data,
        address expectedJobOwner
    ) internal pure {
        // Minimum: 4 bytes selector + 6 params * 32 bytes = 196 bytes
        if (data.length < 196) revert InvalidSafeModuleCalldata();

        address passedJobOwner;
        assembly {
            // jobOwner is the 6th parameter (index 5), at offset 4 + 5*32 = 164
            passedJobOwner := calldataload(add(data.offset, 164))
        }

        if (passedJobOwner != expectedJobOwner) {
            revert JobOwnerMismatch(passedJobOwner, expectedJobOwner);
        }
    }

    function _executeFunction(
        address target,
        bytes memory callData
    ) internal returns (bytes memory) {
        (bool success, bytes memory result) = target.call{value: msg.value}(callData);

        if (success) {
            emit FunctionExecuted(msg.sender, target, callData, msg.value);
        } else {
            emit FunctionExecutionFailed(msg.sender, target, callData, msg.value, result);
        }
        return result;
    }

    function setJobRegistry(
        address _jobRegistryAddress
    ) external onlyOwner {
        jobRegistry = IJobRegistry(_jobRegistryAddress);
    }

    function setTriggerGasRegistry(
        address _triggerGasRegistryAddress
    ) external onlyOwner {
        triggerGasRegistry = ITriggerGasRegistry(_triggerGasRegistryAddress);
    }

    /**
     * @notice Set the taskDispatcher address for signature verification
     * @param _taskDispatcher The address of the authorized taskDispatcher
     */
    function setTaskDispatcher(
        address _taskDispatcher
    ) external onlyOwner {
        require(_taskDispatcher != address(0), "Invalid dispatcher");
        address oldTaskDispatcher = taskDispatcher;
        taskDispatcher = _taskDispatcher;
        emit TaskDispatcherUpdated(oldTaskDispatcher, _taskDispatcher);
    }

    /**
     * @notice Set the TriggerXSafeModule address for jobOwner validation
     * @param _triggerXSafeModule The address of the TriggerXSafeModule contract
     */
    function setTriggerXSafeModule(
        address _triggerXSafeModule
    ) external onlyOwner {
        address oldModule = triggerXSafeModule;
        triggerXSafeModule = _triggerXSafeModule;
        emit TriggerXSafeModuleUpdated(oldModule, _triggerXSafeModule);
    }

    // ---------------------------------------------------------------------
    // -----------------   LayerZero Receive Hook   ------------------------
    // ---------------------------------------------------------------------

    /**
     * @notice Handles incoming LayerZero messages from the hub
     * @param message The message payload containing the action and keeper address
     */
    function _lzReceive(
        Origin calldata,
        /* _origin */
        bytes32,
        /* _guid */
        bytes calldata message,
        address,
        /* _executor */
        bytes calldata /* _extraData */
    ) internal override {
        (ActionType action, address keeper) = abi.decode(message, (ActionType, address));

        if (action == ActionType.REGISTER) {
            isKeeper[keeper] = true;
        } else if (action == ActionType.UNREGISTER) {
            isKeeper[keeper] = false;
        } else {
            revert("Invalid action type");
        }

        emit KeeperUpdated(action, keeper);
    }

    // ---------------------------------------------------------------------
    // --------------------  UUPS Upgrade Authorisation  -------------------
    // ---------------------------------------------------------------------

    /// @dev Required by UUPS pattern. Restricts upgrades to the contract owner.
    function _authorizeUpgrade(
        address newImplementation
    ) internal override onlyOwner {}

    // Storage gap for future upgrades
    uint256[48] private __gap;
}
