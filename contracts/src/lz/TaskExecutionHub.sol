// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Initializable} from "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import {ReentrancyGuardUpgradeable} from "@openzeppelin-upgrades/contracts/utils/ReentrancyGuardUpgradeable.sol";
import {UUPSUpgradeable} from "@openzeppelin-upgrades/contracts/proxy/utils/UUPSUpgradeable.sol";
import {OApp, MessagingFee, Origin} from "@layerzero-v2/oapp/contracts/oapp/OApp.sol";
import {OAppOptionsType3} from "@layerzero-v2/oapp/contracts/oapp/libs/OAppOptionsType3.sol";
import {Ownable} from "@openzeppelin-contracts/contracts/access/Ownable.sol";
import {ECDSA} from "@openzeppelin-contracts/contracts/utils/cryptography/ECDSA.sol";
import {MessageHashUtils} from "@openzeppelin-contracts/contracts/utils/cryptography/MessageHashUtils.sol";

interface IJobRegistry {
    function getJobOwner(uint256 jobId) external view returns (address);

    function unpackJobId(
        uint256 jobId
    ) external view returns (uint256 chainId, uint256 jobCounter);
}

interface ITriggerGasRegistry {
    function deductETHBalance(address user, uint256 ethAmount) external;
}

/**
 * @title TaskExecutionHub
 * @notice A LayerZero-enabled contract that manages task execution across multiple chains
 * @dev This contract acts as a hub for managing task execution and broadcasting their status across different chains
 */
contract TaskExecutionHub is
    Initializable,
    OApp,
    UUPSUpgradeable,
    OAppOptionsType3,
    ReentrancyGuardUpgradeable
{
    /// @custom:oz-upgrades-unsafe-allow constructor
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

    uint32 public originEid;
    uint32 public srcEid;
    uint32[] public dstEids;

    uint128 public defaultGas;
    uint128 public defaultValue;

    IJobRegistry public jobRegistry;
    ITriggerGasRegistry public triggerGasRegistry;

    /// @notice The authorized dispatcher address for signature verification
    address public dispatcher;

    // ----------------------------------
    // -------------  Events ------------
    // ----------------------------------

    event KeeperRegistered(address indexed keeper);
    event KeeperUnregistered(address indexed keeper);
    event BroadcastSent(ActionType action, address keeper, uint32 dstEid);
    event FunctionExecuted(
        address indexed keeper,
        address indexed target,
        bytes data,
        uint256 value
    );
    event FunctionExecutionFailed(
        address indexed keeper,
        address indexed target,
        bytes data,
        uint256 value,
        bytes result
    );
    event FeeUsed(uint32 dstEid, uint256 fee);
    event GasConfigUpdated(uint128 gas, uint128 value);
    event LowBalanceAlert(uint256 currentBalance, uint256 threshold);
    event DispatcherUpdated(
        address indexed oldDispatcher,
        address indexed newDispatcher
    );
    event MessageFailed(
        uint32 indexed dstEid,
        bytes32 indexed guid,
        bytes reason
    );

    enum ActionType {
        REGISTER,
        UNREGISTER
    }

    // ----------------------------------
    // ------------  Modifiers ----------
    // ----------------------------------

    modifier onlyKeeper() {
        require(isKeeper[msg.sender], "Not a keeper");
        _;
    }

    // ---------------------------------------------------------------------
    // Initialization (replaces constructor)
    // ---------------------------------------------------------------------

    /**
     * @dev Initialize function (to be called via proxy).
     * @param _ownerAddress Contract owner.
     * @param _srcEid Source chain endpoint ID.
     * @param _originEid Origin chain endpoint ID.
     * @param _initialKeepers Array of keeper addresses to pre-register.
     * @param _jobRegistryAddress JobRegistry contract.
     * @param _triggerGasRegistryAddress TriggerGasRegistry contract.
     */
    function initialize(
        address _ownerAddress,
        uint32 _srcEid,
        uint32 _originEid,
        address[] calldata _initialKeepers,
        address _jobRegistryAddress,
        address _triggerGasRegistryAddress
    ) external initializer {
        // Init OZ upgradeable helpers
        __ReentrancyGuard_init();
        __UUPSUpgradeable_init();

        // Store basic params
        srcEid = _srcEid;
        originEid = _originEid;
        defaultGas = 1_000_000;
        defaultValue = 0;

        for (uint256 i = 0; i < _initialKeepers.length; i++) {
            isKeeper[_initialKeepers[i]] = true;
            emit KeeperRegistered(_initialKeepers[i]);
        }

        jobRegistry = IJobRegistry(_jobRegistryAddress);
        triggerGasRegistry = ITriggerGasRegistry(_triggerGasRegistryAddress);

        // Delegate & ownership wiring on the LayerZero endpoint
        endpoint.setDelegate(_ownerAddress);
        _transferOwnership(_ownerAddress);
    }

    // ---------------------------------------------------------------------
    // -----------------------   Receive Ether   ---------------------------
    // ---------------------------------------------------------------------

    receive() external payable {}

    // ---------------------------------------------------------------------
    // -----------------------    Main Logic     ---------------------------
    // ---------------------------------------------------------------------

    /**
     * @notice Execute a function call with dispatcher signature verification
     * @param jobId The ID of the job to execute
     * @param ethAmount The amount of ETH to deduct from the job owner
     * @param target The address of the target contract
     * @param data The calldata for the function call
     * @param deadline The signature expiration timestamp
     * @param signature The dispatcher's signature authorizing this execution
     */
    function executeFunction(
        uint256 jobId,
        uint256 ethAmount,
        address target,
        bytes calldata data,
        uint256 deadline,
        bytes calldata signature
    ) external payable onlyKeeper nonReentrant {
        // 1. Deadline check
        require(block.timestamp <= deadline, "Signature expired");

        // 2. Signature verification
        require(dispatcher != address(0), "Dispatcher not set");
        bytes32 hash = keccak256(
            abi.encodePacked(
                jobId,
                target,
                keccak256(data),
                deadline,
                msg.sender,
                block.chainid
            )
        );
        bytes32 ethSignedHash = MessageHashUtils.toEthSignedMessageHash(hash);
        require(
            ECDSA.recover(ethSignedHash, signature) == dispatcher,
            "Invalid signature"
        );

        address jobOwner = jobRegistry.getJobOwner(jobId);
        require(jobOwner != address(0), "Job not found");

        triggerGasRegistry.deductETHBalance(jobOwner, ethAmount);

        _executeFunction(target, data);
    }

    function _executeFunction(
        address target,
        bytes memory callData
    ) internal returns (bytes memory) {
        (bool success, bytes memory result) = target.call{value: msg.value}(
            callData
        );

        if (success) {
            emit FunctionExecuted(msg.sender, target, callData, msg.value);
        } else {
            emit FunctionExecutionFailed(
                msg.sender,
                target,
                callData,
                msg.value,
                result
            );
        }
        return result;
    }

    // ---------------------------------------------------------------------
    // -----------------   LayerZero Receive Hook   ------------------------
    // ---------------------------------------------------------------------

    function _lzReceive(
        Origin calldata _origin,
        bytes32 /*_guid*/,
        bytes calldata _payload,
        address /*_executor*/,
        bytes calldata /*extraData*/
    ) internal override nonReentrant {
        if (address(this).balance < 3e15)
            emit LowBalanceAlert(address(this).balance, 3e15);

        require(_origin.srcEid == srcEid, "Invalid source chain");
        (ActionType action, address keeper) = abi.decode(
            _payload,
            (ActionType, address)
        );

        if (action == ActionType.REGISTER) {
            isKeeper[keeper] = true;
            emit KeeperRegistered(keeper);
            _batchBroadcast(ActionType.REGISTER, keeper);
        } else if (action == ActionType.UNREGISTER) {
            isKeeper[keeper] = false;
            emit KeeperUnregistered(keeper);
            _batchBroadcast(ActionType.UNREGISTER, keeper);
        } else {
            revert("Unknown action");
        }
    }

    // ---------------------------------------------------------------------
    // ------------------  Internal Helper Functions  ----------------------
    // ---------------------------------------------------------------------

    // slither-disable-next-line calls-loop,reentrancy-events
    function _batchBroadcast(ActionType action, address keeper) internal {
        bytes memory payload = abi.encode(action, keeper);
        uint256 totalUsed = 0;
        uint256 initialValue = address(this).balance;

        for (uint256 i = 0; i < dstEids.length; i++) {
            uint32 dstEid = dstEids[i];

            bytes memory options = _buildExecutorOptions(
                defaultGas,
                defaultValue
            );

            try this._quoteFee(dstEid, payload, options) returns (
                MessagingFee memory fee
            ) {
                uint256 feeWithBuffer = fee.nativeFee +
                    (fee.nativeFee * 10) /
                    100;

                if (initialValue < totalUsed + feeWithBuffer) {
                    emit MessageFailed(
                        dstEid,
                        bytes32(0),
                        "Insufficient balance for broadcast (with 10% buffer)"
                    );
                    continue;
                }

                _lzSend(dstEid, payload, options, fee, payable(address(this)));

                emit BroadcastSent(action, keeper, dstEid);
                emit FeeUsed(dstEid, fee.nativeFee);

                totalUsed += feeWithBuffer;
            } catch (bytes memory reason) {
                emit MessageFailed(dstEid, bytes32(0), reason);
            }
        }
    }

    function _quoteFee(
        uint32 dstEid,
        bytes memory payload,
        bytes memory options
    ) external view returns (MessagingFee memory fee) {
        require(msg.sender == address(this), "Only self");
        return _quote(dstEid, payload, options, false);
    }

    function _buildExecutorOptions(
        uint128 gas,
        uint128 value
    ) internal pure returns (bytes memory) {
        uint16 TYPE_3 = 3;
        uint8 WORKER_ID = 1;
        uint8 OPTION_TYPE_LZRECEIVE = 1;

        bytes memory option = value == 0
            ? abi.encodePacked(gas)
            : abi.encodePacked(gas, value);

        uint16 optionLength = uint16(option.length + 1);

        return
            abi.encodePacked(
                TYPE_3,
                WORKER_ID,
                optionLength,
                OPTION_TYPE_LZRECEIVE,
                option
            );
    }

    function _payNative(
        uint256 _nativeFee
    ) internal view override returns (uint256 nativeFee) {
        require(
            address(this).balance >= _nativeFee,
            "Insufficient contract balance"
        );
        return _nativeFee;
    }

    function addSpokes(uint32[] calldata _dstEids) external onlyOwner {
        for (uint256 i = 0; i < _dstEids.length; i++) {
            uint32 dstEid = _dstEids[i];
            if (dstEid != originEid) {
                dstEids.push(dstEid);
                _setPeer(dstEid, bytes32(uint256(uint160(address(this)))));
            }
        }
    }

    function setJobRegistry(address _jobRegistryAddress) external onlyOwner {
        jobRegistry = IJobRegistry(_jobRegistryAddress);
    }

    function setDispatcher(address _dispatcher) external onlyOwner {
        require(_dispatcher != address(0), "Invalid dispatcher address");
        address oldDispatcher = dispatcher;
        dispatcher = _dispatcher;
        emit DispatcherUpdated(oldDispatcher, _dispatcher);
    }

    function setTriggerGasRegistry(
        address _triggerGasRegistryAddress
    ) external onlyOwner {
        triggerGasRegistry = ITriggerGasRegistry(_triggerGasRegistryAddress);
    }

    function setGasConfig(uint128 gas, uint128 value) external onlyOwner {
        defaultGas = gas;
        defaultValue = value;
        emit GasConfigUpdated(gas, value);
    }

    function withdraw(
        address payable to,
        uint256 amount
    ) external onlyOwner nonReentrant {
        require(to != address(0), "Invalid recipient");
        require(amount <= address(this).balance, "Insufficient balance");
        to.transfer(amount);
    }

    function addKeeper(address keeper) external onlyOwner {
        isKeeper[keeper] = true;
        emit KeeperRegistered(keeper);
        _batchBroadcast(ActionType.REGISTER, keeper);
    }

    function removeKeeper(address keeper) external onlyOwner {
        isKeeper[keeper] = false;
        emit KeeperUnregistered(keeper);
        _batchBroadcast(ActionType.UNREGISTER, keeper);
    }

    // ---------------------------------------------------------------------
    // --------------------  UUPS Upgrade Authorisation  -------------------
    // ---------------------------------------------------------------------

    /// @dev Required by UUPS pattern. Restricts upgrades to the contract owner.
    function _authorizeUpgrade(
        address newImplementation
    ) internal override onlyOwner {}

    // Storage gap for future upgrades
    uint256[50] private __gap;
}
