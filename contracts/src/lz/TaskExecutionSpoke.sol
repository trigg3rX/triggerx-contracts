// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import { Initializable } from "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import { ReentrancyGuardUpgradeable } from "@openzeppelin-upgrades/contracts/utils/ReentrancyGuardUpgradeable.sol";
import { UUPSUpgradeable } from "@openzeppelin-upgrades/contracts/proxy/utils/UUPSUpgradeable.sol";
import { OApp, Origin } from "@layerzero-v2/oapp/contracts/oapp/OApp.sol";
import { Ownable } from "@openzeppelin-contracts/contracts/access/Ownable.sol";

interface IJobRegistry {
    function getJobOwner(uint256 jobId) external view returns (address);
}

interface ITriggerGasRegistry {
    function deductTGBalance(address user, uint256 tgAmount) external;
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
    constructor(address _endpoint, address _delegate) OApp(_endpoint, _delegate) Ownable(_delegate) {
        _disableInitializers();
    }

    // ----------------------------------
    // --------  State Variables  -------
    // ----------------------------------

    mapping(address => bool) public isKeeper;

    IJobRegistry public jobRegistry;
    ITriggerGasRegistry public triggerGasRegistry;

    enum ActionType { REGISTER, UNREGISTER }

    // ----------------------------------
    // -------------  Events ------------
    // ----------------------------------

    event KeeperUpdated(ActionType action, address keeper);
    event FunctionExecuted(address indexed keeper, address indexed target, bytes data, uint256 value);

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
     * @notice Executes a function on a target contract
     * @param jobId The ID of the job
     * @param tgAmount The amount of TG to deduct from the job owner
     * @param target The address of the target contract
     * @param data The calldata for the function call
     */
    function executeFunction(uint256 jobId, uint256 tgAmount, address target, bytes memory data) external payable onlyKeeper nonReentrant {
        address jobOwner = jobRegistry.getJobOwner(jobId);
        require(jobOwner != address(0), "Job not found");

        triggerGasRegistry.deductTGBalance(jobOwner, tgAmount);

        _executeFunction(target, data);
    }

    /**
     * @notice Internal function to execute a function call
     * @param target The address of the target contract
     * @param callData The calldata for the function call
     * @return The return data from the function call
     */
    // slither-disable-next-line reentrancy-events
    function _executeFunction(address target, bytes memory callData) internal returns (bytes memory) {
        // slither-disable-next-line low-level-calls
        (bool success, bytes memory result) = target.call{value: msg.value}(callData);
        require(success, "Function execution failed");

        emit FunctionExecuted(msg.sender, target, callData, msg.value);
        return result;
    }

    // ---------------------------------------------------------------------
    // -----------------   LayerZero Receive Hook   ------------------------
    // ---------------------------------------------------------------------

    /**
     * @notice Handles incoming LayerZero messages from the hub
     * @param message The message payload containing the action and keeper address
     */
    function _lzReceive(
        Origin calldata /* _origin */,
        bytes32 /* _guid */,
        bytes calldata message,
        address /* _executor */,
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
    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}

    // Storage gap for future upgrades
    uint256[50] private __gap;
}