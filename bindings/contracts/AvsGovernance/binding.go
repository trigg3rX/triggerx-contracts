// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractAvsGovernance

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BLSAuthLibrarySignature is an auto generated low-level Go binding around an user-defined struct.
type BLSAuthLibrarySignature struct {
	Signature [2]*big.Int
}

// IAvsGovernanceInitializationParams is an auto generated low-level Go binding around an user-defined struct.
type IAvsGovernanceInitializationParams struct {
	AvsGovernanceMultisigOwner common.Address
	OperationsMultisig         common.Address
	CommunityMultisig          common.Address
	OthenticRegistry           common.Address
	MessageHandler             common.Address
	AvsTreasury                common.Address
	AvsDirectoryContract       common.Address
	AllowlistSigner            common.Address
	AvsName                    string
	BlsAuthSingleton           common.Address
}

// IAvsGovernanceOperatorRegistrationParams is an auto generated low-level Go binding around an user-defined struct.
type IAvsGovernanceOperatorRegistrationParams struct {
	BlsKey                   [4]*big.Int
	RewardsReceiver          common.Address
	BlsRegistrationSignature BLSAuthLibrarySignature
	AuthToken                []byte
}

// IAvsGovernanceStakingContractInfo is an auto generated low-level Go binding around an user-defined struct.
type IAvsGovernanceStakingContractInfo struct {
	StakingContract        common.Address
	SharedSecurityProvider uint8
}

// IAvsGovernanceSymbioticOptInSignature is an auto generated low-level Go binding around an user-defined struct.
type IAvsGovernanceSymbioticOptInSignature struct {
	Deadline *big.Int
	Data     []byte
}

// IAvsGovernanceSymbioticOptOutSignature is an auto generated low-level Go binding around an user-defined struct.
type IAvsGovernanceSymbioticOptOutSignature struct {
	Deadline *big.Int
	Data     []byte
}

// IAvsGovernanceVotingPowerMultiplier is an auto generated low-level Go binding around an user-defined struct.
type IAvsGovernanceVotingPowerMultiplier struct {
	StakingContract        common.Address
	Multiplier             *big.Int
	SharedSecurityProvider uint8
}

// IRewardsCoordinatorOperatorReward is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorOperatorReward struct {
	Operator common.Address
	Amount   *big.Int
}

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// ContractAvsGovernanceMetaData contains all meta data concerning the ContractAvsGovernance contract.
var ContractAvsGovernanceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_extensionImplementation\",\"type\":\"address\"},{\"internalType\":\"contractIOthenticRegistry\",\"name\":\"_othenticRegistryAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccessControlInvalidMultiplierSyncer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllowlistDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllowlistEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyAvsName\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptySharedSecurityProvidersList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FlowIsCurrentlyPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FlowIsCurrentlyUnpaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAllowlistAuthToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBlsRegistrationSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMultiplier\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMultiplierNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRewardsReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"arrayIndex\",\"type\":\"uint256\"}],\"name\":\"InvalidSharedSecurityProviderList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSlashingRate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStakingContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"MissingAuthToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ModificationDelayNotPassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NativeCoinNotSupportedForEigenRewardsError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughVotingPower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numOfOperatorsLimit\",\"type\":\"uint256\"}],\"name\":\"NumOfOperatorsLimitReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorNotRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorStillRegisteredToSharedSecurityProviders\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PauseFlowIsAlreadyPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StakingContractsNotInAscendingOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TreasuryWithdrawRewardsFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnpausingFlowIsAlreadyUnpaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DepositRewardsBackFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"_pausableFlow\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_pauser\",\"type\":\"address\"}],\"name\":\"FlowPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"_pausableFlowFlag\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_unpauser\",\"type\":\"address\"}],\"name\":\"FlowUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"NativeCoinNotSupportedForEigenRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorEjectedFromNetwork\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[4]\",\"name\":\"blsKey\",\"type\":\"uint256[4]\"}],\"name\":\"OperatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorRegisteredToEigenLayer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorRegisteredToSymbiotic\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorUnregisteredToEigenLayer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorUnregisteredToSymbiotic\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"QueuedRewardsReceiverModification\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"revertData\",\"type\":\"bytes\"}],\"name\":\"RewardsCoordinatorReverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"avsName\",\"type\":\"string\"}],\"name\":\"SetAvsName\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isAllowlisted\",\"type\":\"bool\"}],\"name\":\"SetIsAllowlisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_isEnabled\",\"type\":\"bool\"}],\"name\":\"SetP2pAuthenticationEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"SetRewardsReceiver\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"stakingContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"multiplier\",\"type\":\"uint256\"}],\"name\":\"SetStakingContractMultiplier\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SetToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"stakingContracts\",\"type\":\"address[]\"}],\"name\":\"setNewSupportedStakingContracts\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXTENSION_IMPLEMENTATION\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"avsName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"avsTreasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"completeRewardsReceiverModification\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_operators\",\"type\":\"address[]\"},{\"internalType\":\"uint32\",\"name\":\"_startTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_duration\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_totalAmount\",\"type\":\"uint256\"}],\"name\":\"createEigenRewardsSubmission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIRewardsCoordinator.OperatorReward[]\",\"name\":\"_operators\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"_lastPayedTask\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_rewardsData\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"_remoteEid\",\"type\":\"uint32\"}],\"name\":\"createOperatorDirectedAVSRewardsSubmission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"ejectOperatorFromNetwork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIsAllowlisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"getOperatorRestakedStrategies\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRestakeableStrategies\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"getRewardsReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"avsGovernanceMultisigOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operationsMultisig\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"communityMultisig\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"othenticRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"messageHandler\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"avsTreasury\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"avsDirectoryContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowlistSigner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"avsName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"blsAuthSingleton\",\"type\":\"address\"}],\"internalType\":\"structIAvsGovernance.InitializationParams\",\"name\":\"_initializationParams\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_pausableFlow\",\"type\":\"bytes4\"}],\"name\":\"isFlowPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_isPaused\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isOperatorRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxEffectiveBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingContract\",\"type\":\"address\"}],\"name\":\"minStakeAmountPerStakingContract\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minVotingPower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingContract\",\"type\":\"address\"}],\"name\":\"multiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numOfActiveOperators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"p2pAuthenticationEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_pausableFlow\",\"type\":\"bytes4\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newRewardsReceiver\",\"type\":\"address\"}],\"name\":\"queueRewardsReceiverModification\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[4]\",\"name\":\"blsKey\",\"type\":\"uint256[4]\"},{\"internalType\":\"address\",\"name\":\"rewardsReceiver\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"signature\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBLSAuthLibrary.Signature\",\"name\":\"blsRegistrationSignature\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"authToken\",\"type\":\"bytes\"}],\"internalType\":\"structIAvsGovernance.OperatorRegistrationParams\",\"name\":\"_operatorRegistrationParams\",\"type\":\"tuple\"}],\"name\":\"registerAsOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"}],\"name\":\"registerAvsToEigenLayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registerAvsToSymbiotic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"name\":\"_eigenSig\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_authToken\",\"type\":\"bytes\"}],\"name\":\"registerOperatorToEigenLayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint48\",\"name\":\"deadline\",\"type\":\"uint48\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIAvsGovernance.SymbioticOptInSignature\",\"name\":\"_symbioticSig\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_authToken\",\"type\":\"bytes\"}],\"name\":\"registerOperatorToSymbiotic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_p2pAuthenticationEnabled\",\"type\":\"bool\"}],\"name\":\"setP2pAuthenticationEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"stakingContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"multiplier\",\"type\":\"uint256\"},{\"internalType\":\"enumIAvsGovernance.SharedSecurityProvider\",\"name\":\"sharedSecurityProvider\",\"type\":\"uint8\"}],\"internalType\":\"structIAvsGovernance.VotingPowerMultiplier\",\"name\":\"_votingPowerMultiplier\",\"type\":\"tuple\"}],\"name\":\"setStakingContractMultiplier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"stakingContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"multiplier\",\"type\":\"uint256\"},{\"internalType\":\"enumIAvsGovernance.SharedSecurityProvider\",\"name\":\"sharedSecurityProvider\",\"type\":\"uint8\"}],\"internalType\":\"structIAvsGovernance.VotingPowerMultiplier[]\",\"name\":\"_votingPowerMultipliers\",\"type\":\"tuple[]\"}],\"name\":\"setStakingContractMultiplierBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"stakingContract\",\"type\":\"address\"},{\"internalType\":\"enumIAvsGovernance.SharedSecurityProvider\",\"name\":\"sharedSecurityProvider\",\"type\":\"uint8\"}],\"internalType\":\"structIAvsGovernance.StakingContractInfo[]\",\"name\":\"_stakingContractsDetails\",\"type\":\"tuple[]\"}],\"name\":\"setSupportedStakingContracts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingContracts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_pausableFlow\",\"type\":\"bytes4\"}],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unregisterAsOperatorFromAvs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unregisterAsOperatorFromEigenLayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint48\",\"name\":\"deadline\",\"type\":\"uint48\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIAvsGovernance.SymbioticOptOutSignature\",\"name\":\"_unregistrationSignature\",\"type\":\"tuple\"}],\"name\":\"unregisterAsOperatorFromSymbiotic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"votingPower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_stakingContracts\",\"type\":\"address[]\"}],\"name\":\"votingPowerPerStakingContracts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506004361061027f5760003560e01c80637897dec31161015c578063b525fa88116100ce578063d547741f11610087578063d547741f14610574578063e481af9d14610587578063e6474b0f1461058f578063efd9697814610597578063f6e258cc146105d1578063fab57b8f146105e45761027f565b8063b525fa8814610518578063bac1e94b14610520578063c07473f614610533578063c7a93c6f14610546578063cab1a5ef14610559578063cbd969c9146105615761027f565b8063936522771161012057806393652277146104c7578063a217fddf146104da578063a88171ee146104e2578063a9b3f8b7146104ea578063aa0dadf0146104fd578063ab8cac54146105055761027f565b80637897dec31461047357806382616a011461047b57806383e451701461048e578063851bb725146104a157806391d14854146104b45761027f565b8063263fe7c7116101f55780633aa83ec7116101b95780633aa83ec71461040a57806341b92a291461041d578063560d5484146104325780635e95cee21461043a57806362de73931461044d5780636b1906f8146104605761027f565b8063263fe7c7146103b45780632f2ff15d146103bc57806333cfb7b7146103cf57806336568abe146103ef57806336fffde0146104025761027f565b806313fe45211161024757806313fe4521146103205780631b21ba72146103415780631b409bf2146103545780631dd2f74b14610367578063226def041461037a578063248a9ca3146103a15761027f565b806301ffc9a7146102aa5780630399f004146102d257806306255b01146102da5780631246193e146102ed578063130bdcd91461030d575b6102a87f000000000000000000000000f80a16be05b9f4665268cd69fd9379f17b0636c86105f7565b005b6102bd6102b8366004614a23565b610620565b60405190151581526020015b60405180910390f35b6102a8610657565b6102a86102e8366004614a4d565b610766565b6102f561090e565b6040516001600160a01b0390911681526020016102c9565b6102a861031b366004614a87565b61092a565b61033361032e366004614b10565b610982565b6040519081526020016102c9565b6102a861034f366004614b10565b6109ae565b6102a8610362366004614b93565b610aff565b6102a8610375366004614ca0565b610ed0565b6102f57f000000000000000000000000f80a16be05b9f4665268cd69fd9379f17b0636c881565b6103336103af366004614d6d565b610f1c565b6102a8610f3e565b6102a86103ca366004614d86565b611243565b6103e26103dd366004614b10565b611265565b6040516102c99190614dfb565b6102a86103fd366004614d86565b611303565b610333611336565b6102a8610418366004614a23565b611349565b610425611370565b6040516102c99190614e5e565b6103e261140b565b6102f5610448366004614b10565b611475565b6102a861045b366004614ee0565b6114a5565b6102bd61046e366004614b10565b611976565b6103336119a4565b6102a8610489366004614fe0565b6119b7565b6102a861049c366004615039565b6119d9565b6102a86104af3660046150fb565b611a02565b6102bd6104c2366004614d86565b611a85565b6102a86104d5366004615152565b611abd565b610333600081565b610333611ade565b6103336104f8366004614b10565b611af1565b6102a8611b1d565b6102a86105133660046151e1565b611cdb565b6102bd611d8b565b6102a861052e366004614a23565b611da8565b610333610541366004614b10565b611dcf565b6102a861055436600461522a565b611e5a565b6102bd611ebf565b61033361056f366004615247565b611ed5565b6102a8610582366004614d86565b611f3b565b6103e2611f57565b6102a8611f66565b6102bd6105a5366004614a23565b6001600160e01b0319166000908152600080516020615c29833981519152602052604090205460ff1690565b6102a86105df366004614b10565b612068565b6102a86105f236600461528e565b612219565b3660008037600080366000845af43d6000803e808015610616573d6000f35b3d6000fd5b505050565b60006001600160e01b03198216637965db0b60e01b148061065157506301ffc9a760e01b6001600160e01b03198316145b92915050565b61065f612327565b33600090815260079190910160205260408120549003610692576040516325ec6c1f60e01b815260040160405180910390fd5b600080516020615c898339815191526106aa81612331565b6106b2612379565b6106ba612327565b600501546040516351b27a6d60e11b81523360048201526001600160a01b039091169063a364f4da90602401600060405180830381600087803b15801561070057600080fd5b505af1158015610714573d6000803e3d6000fd5b50506040513381527f7877897edb095f8cd999adcb582e04128ffb91d12bf5139ca7b8a1567a8453009250602001905060405180910390a16107636001600080516020615c6983398151915255565b50565b61076e612327565b336000908152600791909101602052604081205490036107a1576040516325ec6c1f60e01b815260040160405180910390fd5b600080516020615c898339815191526107b981612331565b6107c1612379565b7f000000000000000000000000bf844a9fcc63bb081415ef8cc3447e1cda33a3096001600160a01b0316635ee43d376040518163ffffffff1660e01b8152600401602060405180830381865afa15801561081f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061084391906152c9565b6001600160a01b03166393f79bc3333061086060208701876152e6565b61086d6020880188615301565b6040518663ffffffff1660e01b815260040161088d959493929190615370565b600060405180830381600087803b1580156108a757600080fd5b505af11580156108bb573d6000803e3d6000fd5b50506040513381527ffc7eb04654394ed6ffe862307d7824361d3481d69ea2775b77e85837befe77379250602001905060405180910390a161090a6001600080516020615c6983398151915255565b5050565b6000610918612327565b600401546001600160a01b0316919050565b63ef0892d160e01b61093b816123c5565b6000610945612327565b905060005b8381101561097b5761097382868684818110610968576109686153ab565b9050606002016123cf565b60010161094a565b5050505050565b600061098c612327565b6001600160a01b03909216600090815260129290920160205250604090205490565b6109b6612327565b336000908152600791909101602052604081205490036109e9576040516325ec6c1f60e01b815260040160405180910390fd5b7fd93c394bc0520b0dc9db435f0321020356363275d7d80fbdc9f5296937698fc6610a1381612331565b6001600160a01b038216610a395760405162cc6ac760e01b815260040160405180910390fd5b6000610a43612327565b336000908152600c820160205260408120805460ff19166001179055600b82015491925090610a7290426153d7565b6040805180820182526001600160a01b038781168083526020808401868152336000818152600e8b018452879020955186546001600160a01b03191695169490941785555160019094019390935583519081529182018490529293507f0d8cfa10a3087b28d3c226ad9a37314860e7c3c0505a25a39e3cdefb3118a98a910160405180910390a250505050565b6378b4401360e11b610b10816123c5565b6000859003610b5e5760405162461bcd60e51b8152602060048201526015602482015274139bc81bdc195c985d1bdc9cc81c1c9bdd9a591959605a1b60448201526064015b60405180910390fd5b6000610b68612327565b905060008160040160009054906101000a90046001600160a01b0316905060007f000000000000000000000000bf844a9fcc63bb081415ef8cc3447e1cda33a3096001600160a01b0316638a2fc4e36040518163ffffffff1660e01b8152600401602060405180830381865afa158015610be6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c0a91906152c9565b905060008360040160009054906101000a90046001600160a01b03166001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c63573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c8791906152c9565b90508360040160009054906101000a90046001600160a01b03166001600160a01b031663a734f06e6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610cde573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d0291906152c9565b6001600160a01b0316816001600160a01b031603610d3357604051634acb03b560e01b815260040160405180910390fd5b610d41303060008987612489565b610d5e576040516359062f7f60e01b815260040160405180910390fd5b6000610d6a8a886153ea565b905060008a6001600160401b03811115610d8657610d86614c02565b604051908082528060200260200182016040528015610dcb57816020015b6040805180820190915260008082526020820152815260200190600190039081610da45790505b50905060005b8b811015610e425760405180604001604052808e8e84818110610df657610df66153ab565b9050602002016020810190610e0b9190614b10565b6001600160a01b0316815260200184815250828281518110610e2f57610e2f6153ab565b6020908102919091010152600101610dd1565b50610e576001600160a01b038416858a6125a1565b836001600160a01b0316639cb9a5fa30610e73848e8e8961262b565b6040518363ffffffff1660e01b8152600401610e90929190615451565b600060405180830381600087803b158015610eaa57600080fd5b505af1158015610ebe573d6000803e3d6000fd5b50505050505050505050505050505050565b6378b4401360e11b610ee1816123c5565b7f37b1f16af49fd5d40e5dc679bf35afa8bb06df10cb8308e0aa633882d2a81201610f0b81612331565b61061b610f16612327565b84612736565b6000908152600080516020615c49833981519152602052604090206001015490565b610f46612327565b33600090815260079190910160205260408120549003610f79576040516325ec6c1f60e01b815260040160405180910390fd5b600080516020615c89833981519152610f9181612331565b610f99612379565b6000610fa3612327565b9050600160058201546040516349075da360e01b81523060048201523360248201526001600160a01b03909116906349075da390604401602060405180830381865afa158015610ff7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061101b9190615594565b600181111561102c5761102c61557e565b1415801561112957507f000000000000000000000000bf844a9fcc63bb081415ef8cc3447e1cda33a3096001600160a01b0316635ee43d376040518163ffffffff1660e01b8152600401602060405180830381865afa158015611093573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110b791906152c9565b6040516308834cb560e21b81523360048201523060248201526001600160a01b03919091169063220d32d490604401602060405180830381865afa158015611103573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061112791906155b1565b155b156112105760088101546001600160a01b0316801580159061119f576040516311c7e79960e21b81523360048201526001600160a01b0383169063471f9e6490602401600060405180830381600087803b15801561118657600080fd5b505af115801561119a573d6000803e3d6000fd5b505050505b6111a98333612a7f565b80156112095760405163e9ecc1cb60e01b81523360048201526001600160a01b0383169063e9ecc1cb90602401600060405180830381600087803b1580156111f057600080fd5b505af1158015611204573d6000803e3d6000fd5b505050505b505061122b565b604051633c1d10e160e11b8152336004820152602401610b55565b506107636001600080516020615c6983398151915255565b61124c82610f1c565b611255816123c5565b61125f8383612ae5565b50505050565b60607f000000000000000000000000bf844a9fcc63bb081415ef8cc3447e1cda33a3096001600160a01b03166390dea2af8361129f612b8a565b306040518463ffffffff1660e01b81526004016112be939291906155ce565b600060405180830381865afa1580156112db573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526106519190810190615603565b6001600160a01b038116331461132c5760405163334bd91960e11b815260040160405180910390fd5b61061b8282612b96565b6000611340612327565b60100154905090565b8061135381612331565b6001600160e01b03198216611367816123c5565b61061b83612c12565b606061137a612327565b600a01805461138890615691565b80601f01602080910402602001604051908101604052809291908181526020018280546113b490615691565b80156114015780601f106113d657610100808354040283529160200191611401565b820191906000526020600020905b8154815290600101906020018083116113e457829003601f168201915b5050505050905090565b6060611415612327565b60090180548060200260200160405190810160405280929190818152602001828054801561140157602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161144e575050505050905090565b600061147f612327565b6001600160a01b039283166000908152600d919091016020526040902054909116919050565b638a70a0eb60e01b6114b6816123c5565b60006114c0612327565b905060008160040160009054906101000a90046001600160a01b0316905060007f000000000000000000000000bf844a9fcc63bb081415ef8cc3447e1cda33a3096001600160a01b0316638a2fc4e36040518163ffffffff1660e01b8152600401602060405180830381865afa15801561153e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061156291906152c9565b905061159b6040518060800160405280600063ffffffff168152602001600063ffffffff16815260200160008152602001600081525090565b868060200190518101906115af91906156c5565b606085015260408085019190915263ffffffff91821660208086019190915292909116835260048087015482516321df0da760e01b815292516000946001600160a01b03909216936321df0da7938181019392918290030181865afa15801561161c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061164091906152c9565b9050600060608660040160009054906101000a90046001600160a01b03166001600160a01b031663a734f06e6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561169b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116bf91906152c9565b6001600160a01b0316836001600160a01b03161461189c576116e830308d87604001518a612489565b9150811561189757604084015161170b906001600160a01b038516908790612cb3565b50846001600160a01b0316639cb9a5fa60e01b306117338f886000015189602001518961262b565b604051602401611744929190615451565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909416939093179092529051611782919061570d565b6000604051808303816000865af19150503d80600081146117bf576040519150601f19603f3d011682016040523d82523d6000602084013e6117c4565b606091505b509092509050816118975760408401516117ea906001600160a01b038516908890612cb3565b5060408085015190516356a8904960e11b81526000916001600160a01b0389169163ad512092916118219160040190815260200190565b6020604051808303816000875af1158015611840573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061186491906155b1565b905080611895576040517fed3143332c2d71fbe302b126ea59ee4fb74a1699cb820f835a913b15c027798d90600090a15b505b6118c6565b6040517fb3f3d8f7b39bd24844dec3cce1a06cfee42f663826e9b790c6a36fb8c05dbf8e90600090a15b8161193d5760005b84606001518110156119045760008d82815181106118ee576118ee6153ab565b60209081029190910181015101526001016118ce565b507fab7e71f561320fa16a226c9744bf4d230ce7faf885df7bbc890baf339faa0a4e816040516119349190614e5e565b60405180910390a15b611968878d6040516020016119529190615729565b6040516020818303038152906040528d8c612d95565b505050505050505050505050565b6000611980612327565b6001600160a01b039092166000908152600792909201602052506040902054151590565b60006119ae612327565b60010154905090565b63ef0892d160e01b6119c8816123c5565b61090a6119d3612327565b836123cf565b60006119e3612327565b90506119f0818484612e30565b61125f6119fb612327565b3386612ee5565b6378b4401360e11b611a13816123c5565b611a1b612327565b6005015460405163a98fb35560e01b81526001600160a01b039091169063a98fb35590611a4e908690869060040161573c565b600060405180830381600087803b158015611a6857600080fd5b505af1158015611a7c573d6000803e3d6000fd5b50505050505050565b6000918252600080516020615c49833981519152602090815260408084206001600160a01b0393909316845291905290205460ff1690565b6000611ac7612327565b9050611ad4818484612e30565b61125f3385612f85565b6000611ae8612327565b60110154905090565b6000611afb612327565b6001600160a01b03909216600090815260139290920160205250604090205490565b6378b4401360e11b611b2e816123c5565b7f000000000000000000000000bf844a9fcc63bb081415ef8cc3447e1cda33a3096001600160a01b031663e45f40be6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611b8c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611bb091906152c9565b6001600160a01b03166387140b5b6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015611bea57600080fd5b505af1158015611bfe573d6000803e3d6000fd5b505050507f000000000000000000000000bf844a9fcc63bb081415ef8cc3447e1cda33a3096001600160a01b0316630bb79bc06040518163ffffffff1660e01b8152600401602060405180830381865afa158015611c60573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c8491906152c9565b60405163b7d8e1a960e01b81523060048201526001600160a01b03919091169063b7d8e1a990602401600060405180830381600087803b158015611cc757600080fd5b505af115801561097b573d6000803e3d6000fd5b611ce3612327565b336000908152600791909101602052604090205415611d15576040516342ee68b560e01b815260040160405180910390fd5b600080516020615c89833981519152611d2d81612331565b611d35612379565b6000611d3f612327565b9050611d5781611d5260e0860186615301565b612e30565b611d7381338560a08101611d6e8160808401614b10565b6130af565b5061090a6001600080516020615c6983398151915255565b6000611d95612327565b60050154600160a01b900460ff16919050565b80611db281613459565b6001600160e01b03198216611dc6816123c5565b61061b836134a0565b600080611dda612327565b90506000611e448283600901805480602002602001604051908101604052809291908181526020018280548015611e3a57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611e1c575b505050505061353d565b915050611e5282858361375d565b949350505050565b6378b4401360e11b611e6b816123c5565b60405182151581527f476bc7159a4e42477e0aa2e2e9709c55cb469b547976f91002b8db00bbd9b2f69060200160405180910390a181611ea9612327565b601701805460ff19169115159190911790555050565b6000611ec9612327565b6017015460ff16919050565b600080611ee0612327565b90506000611f218286868080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525061353d92505050565b915050611f2f82878361375d565b925050505b9392505050565b611f4482610f1c565b611f4d816123c5565b61125f8383612b96565b6060611f61612b8a565b905090565b611f6e612327565b33600090815260079190910160205260408120549003611fa1576040516325ec6c1f60e01b815260040160405180910390fd5b6000611fab612327565b336000908152600e82016020526040902060010154909150421015611fe357604051638ce7a3f160e01b815260040160405180910390fd5b336000908152600e820160205260409020546120099082906001600160a01b0316613818565b336000818152600c830160209081526040808320805460ff1916905580518082018252838152808301848152948452600e9095019091529020915182546001600160a01b0319166001600160a01b039091161782555160019190910155565b612070612379565b638a70a0eb60e01b612081816123c5565b600061208b612327565b9050600160058201546040516349075da360e01b81523060048201526001600160a01b038681166024830152909116906349075da390604401602060405180830381865afa1580156120e1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121059190615594565b60018111156121165761211661557e565b036121ba5760058101546040516351b27a6d60e11b81526001600160a01b0385811660048301529091169063a364f4da90602401600060405180830381600087803b15801561216457600080fd5b505af1158015612178573d6000803e3d6000fd5b50506040516001600160a01b03861681527f7877897edb095f8cd999adcb582e04128ffb91d12bf5139ca7b8a1567a8453009250602001905060405180910390a15b6121c48184612a7f565b6040516001600160a01b03841681527f98ac533c958bf5d682a0590dc21d14cc6a8acaaa7886880124cfb4531266f2f29060200160405180910390a150506107636001600080516020615c6983398151915255565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b031660008115801561225e5750825b90506000826001600160401b0316600114801561227a5750303b155b905081158015612288575080155b156122a65760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156122d057845460ff60401b1916600160401b1785555b6122d98661387d565b831561231f57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050565b6000611f61613bef565b6001600160e01b031981166000908152600080516020615c29833981519152602052604090205460ff16156107635760405163722fdba960e01b815260040160405180910390fd5b600080516020615c698339815191528054600119016123ab57604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b6001600080516020615c6983398151915255565b6107638133613c1d565b6001600160601b03602082013511156123fb57604051631bc4bcf760e21b815260040160405180910390fd5b6020810180359060138401906000906124149085614b10565b6001600160a01b03168152602080820192909252604001600020919091557fa717a4f9ede961a75886c212a98163df9b8ff2e23b83c9fe3eb8a73f56cf385f9061246090830183614b10565b604080516001600160a01b03909216825260208085013590830152015b60405180910390a15050565b60006001600160a01b0385161561251b57604051633256b4d160e01b81526001600160a01b0386811660048301526024820186905260448201859052831690633256b4d1906064016020604051808303816000875af11580156124f0573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061251491906155b1565b9050612598565b604051633256b4d160e01b81526001600160a01b0387811660048301526024820186905260448201859052831690633256b4d1906064016020604051808303816000875af1158015612571573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061259591906155b1565b90505b95945050505050565b604051636eb1769f60e11b81523060048201526001600160a01b0383811660248301526000919085169063dd62ed3e90604401602060405180830381865afa1580156125f1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126159190615750565b905061125f848461262685856153d7565b613c56565b60606000612637612327565b60408051600180825281830190925291925060009190816020015b6126a06040518060c001604052806060815260200160006001600160a01b0316815260200160608152602001600063ffffffff168152602001600063ffffffff168152602001606081525090565b8152602001906001900390816126525790505090506040518060c001604052806126c984613d0a565b8152602001856001600160a01b031681526020018881526020018763ffffffff1681526020018663ffffffff1681526020016040518060200160405280600081525081525081600081518110612721576127216153ab565b60209081029190910101529695505050505050565b6127446009830160006149f1565b6000805b8251811015612a3f576000838281518110612765576127656153ab565b60200260200101519050826001600160a01b031681600001516001600160a01b031610156127a65760405163b854d24160e01b815260040160405180910390fd5b805192506001600160a01b0383166127d1576040516307b18df360e31b815260040160405180910390fd5b8051600986018054600180820183556000928352602080842090920180546001600160a01b0319166001600160a01b0395861617905581850151855190941683526016890190915260409091208054909160ff1990911690838181111561283a5761283a61557e565b02179055506001816020015160018111156128575761285761557e565b14801561294d5750838281518110612871576128716153ab565b6020026020010151600001516001600160a01b031663ce9b79306040518163ffffffff1660e01b8152600401602060405180830381865afa1580156128ba573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128de91906152c9565b6040516368adba0760e11b81523060601b60048201526001600160a01b03919091169063d15b740e90602401602060405180830381865afa158015612927573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061294b9190615750565b155b15612a3657838281518110612964576129646153ab565b6020026020010151600001516001600160a01b031663ce9b79306040518163ffffffff1660e01b8152600401602060405180830381865afa1580156129ad573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129d191906152c9565b6040516323f752d560e01b81526000600482015260001960248201526001600160a01b0391909116906323f752d590604401600060405180830381600087803b158015612a1d57600080fd5b505af1158015612a31573d6000803e3d6000fd5b505050505b50600101612748565b507fdc9dd92c894ed28df4cdf880b63e1f5d6a6a32870a392e08c13ddfe18488caea83600901604051612a729190615769565b60405180910390a1505050565b612a898282613eb9565b8160010160008154612a9a906157b9565b909155506001600160a01b0381166000818152600784016020526040808220829055517f6f42117a557500c705ddf040a619d86f39101e6b74ac20d7b3e5943ba473fc7f9190a25050565b6000600080516020615c49833981519152612b008484611a85565b612b80576000848152602082815260408083206001600160a01b03871684529091529020805460ff19166001179055612b363390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610651565b6000915050610651565b6060611f616000613f56565b6000600080516020615c49833981519152612bb18484611a85565b15612b80576000848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610651565b6001600160e01b031981166000908152600080516020615c29833981519152602081905260409091205460ff1615612c5d5760405163dfe10d7d60e01b815260040160405180910390fd5b6001600160e01b0319821660009081526020829052604090819020805460ff19166001179055517f95c3658c5e0c74e20cf12db371b9b67d26e97a1937f6d2284f88cc44d036b4f69061247d90849033906157d0565b604051636eb1769f60e11b81523060048201526001600160a01b038381166024830152600091829186169063dd62ed3e90604401602060405180830381865afa158015612d04573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d289190615750565b9050612598856001600160a01b03811663095ea7b387612d4888876153d7565b6040516001600160a01b0390921660248301526044820152606401604051602081830303815290604052915060e01b6020820180516001600160e01b03838183161783525050505061407b565b6000612da18484614113565b60028601546040519192506001600160a01b0316906382646a5890612dcc90849086906020016157f3565b6040516020818303038152906040526040518263ffffffff1660e01b8152600401612df79190614e5e565b600060405180830381600087803b158015612e1157600080fd5b505af1158015612e25573d6000803e3d6000fd5b505050505050505050565b6005830154600160a01b900460ff161561061b576000819003612e6a57818160405163471b4be960e11b8152600401610b5592919061573c565b612ec830338560060160009054906101000a90046001600160a01b031685858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594939250506141839050565b61061b5760405163d2342ec760e01b815260040160405180910390fd5b6005830154604051639926ee7d60e01b81526001600160a01b0390911690639926ee7d90612f19908590859060040161581b565b600060405180830381600087803b158015612f3357600080fd5b505af1158015612f47573d6000803e3d6000fd5b50506040516001600160a01b03851681527f23c66044e8b2ab58f45f4aa7eeb03c47c231fec8b7cf140f5b709455e5c1aa1d92506020019050612a72565b60007f000000000000000000000000bf844a9fcc63bb081415ef8cc3447e1cda33a3096001600160a01b0316635ee43d376040518163ffffffff1660e01b8152600401602060405180830381865afa158015612fe5573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061300991906152c9565b8251602084015160405163ced44ba760e01b81529293506001600160a01b0384169263ced44ba79261304392889230929190600401615866565b600060405180830381600087803b15801561305d57600080fd5b505af1158015613071573d6000803e3d6000fd5b50506040516001600160a01b03861681527f6f6f010344f25a8f5774e47688e7db5f23ef1071452f6048e1147d8de8cfe2d192506020019050612a72565b6001600160a01b0381166130d55760405162cc6ac760e01b815260040160405180910390fd5b6130df8582613818565b600f8501546001860154811161310b5760405163c77b407760e01b815260048101829052602401610b55565b6015860154604051633cf65e3b60e01b81526001600160a01b0390911690633cf65e3b90613143908690899030908a906004016158a0565b602060405180830381865afa158015613160573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061318491906155b1565b6131a1576040516314532dfd60e11b815260040160405180910390fd5b6008860154600987018054604080516020808402820181019092528281526001600160a01b039094169384151593600093849361321f938e9390929091830182828015611e3a576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611e1c57505050505061353d565b9150915060006132308b8b8461375d565b905060008b6010015482101580156132d457506040516329ea021d60e01b81526001600160a01b037f000000000000000000000000bf844a9fcc63bb081415ef8cc3447e1cda33a30916906329ea021d90613293908e9088903090600401615916565b602060405180830381865afa1580156132b0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906132d491906155b1565b9050806132f457604051631c33ce8d60e11b815260040160405180910390fd5b841561335f5760405163094e7a3f60e41b81526001600160a01b038716906394e7a3f09061332c908e9086908f908e90600401615982565b600060405180830381600087803b15801561334657600080fd5b505af115801561335a573d6000803e3d6000fd5b505050505b61336c8c8c848d8c6141ad565b8b6001016000815461337d906159b5565b90915550506001600160a01b038a16600090815260078c016020526040902060019055831561340b576040516376c56c1b60e11b81526001600160a01b0386169063ed8ad836906133d8908d9085908e908d90600401615982565b600060405180830381600087803b1580156133f257600080fd5b505af1158015613406573d6000803e3d6000fd5b505050505b50505050856001600160a01b03167f54bc9cf83c2eb0f2ad1abf6e4fab882964404622ba2df6b5a9356a18d3aac0558660405161344891906159ce565b60405180910390a250505050505050565b6001600160e01b031981166000908152600080516020615c29833981519152602052604090205460ff16610763576040516368c87f3360e11b815260040160405180910390fd5b6001600160e01b031981166000908152600080516020615c29833981519152602081905260409091205460ff166134ea57604051635bfd2da760e11b815260040160405180910390fd5b6001600160e01b0319821660009081526020829052604090819020805460ff19169055517fc7e56e17b0a6c4b467df6495e1eda1baecd7ba20604e80c1058ac06f4578d85e9061247d90849033906157d0565b606080600083516001600160401b0381111561355b5761355b614c02565b6040519080825280602002602001820160405280156135ad57816020015b61359a60408051606081018252600080825260208201819052909182015290565b8152602001906001900390816135795790505b509050600084516001600160401b038111156135cb576135cb614c02565b60405190808252806020026020018201604052801561361d57816020015b61360a60408051606081018252600080825260208201819052909182015290565b8152602001906001900390816135e95790505b50905060005b855181101561374f576000868281518110613640576136406153ab565b6020908102919091018101516001600160a01b038116600081815260168c018452604080822054815160608101835284815293835260128e01865291819020549483019490945291935060ff9091169181018260018111156136a4576136a461557e565b8152508584815181106136b9576136b96153ab565b6020908102919091018101919091526001600160a01b038316600090815260138b019091526040812054908190036136ef575060015b6040518060600160405280846001600160a01b031681526020018281526020018360018111156137215761372161557e565b815250858581518110613736576137366153ab565b6020026020010181905250836001019350505050613623565b5090925090505b9250929050565b6011830154604051630e1b4e1f60e11b81526000919082906001600160a01b037f000000000000000000000000bf844a9fcc63bb081415ef8cc3447e1cda33a3091690631c369c3e906137b8908890889030906004016159dd565b602060405180830381865afa1580156137d5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906137f99190615750565b9050818111801561380a5750600082115b156125985750949350505050565b336000818152600d8401602090815260409182902080546001600160a01b0319166001600160a01b03861690811790915591519182527fe906feea2ef60b474e22b4169bdd4de6906a84cd448cbcee99593526fe87082d910160405180910390a25050565b61388561424c565b60006138946020830183614b10565b6001600160a01b031614806138c1575060006138b66040830160208401614b10565b6001600160a01b0316145b806138e4575060006138d96060830160408401614b10565b6001600160a01b0316145b80613907575060006138fc60a0830160808401614b10565b6001600160a01b0316145b8061392a5750600061391f60c0830160a08401614b10565b6001600160a01b0316145b8061394d5750600061394260e0830160c08401614b10565b6001600160a01b0316145b8061397157506000613966610100830160e08401614b10565b6001600160a01b0316145b806139965750600061398b61014083016101208401614b10565b6001600160a01b0316145b156139b45760405163d92e233d60e01b815260040160405180910390fd5b60006139c36020830183614b10565b905060006139d76040840160208501614b10565b905060006139eb6060850160408601614b10565b905060006139ff60a0860160808701614b10565b9050366000613a12610100880188615301565b91509150613a21868686614297565b613a2c8686866142d4565b613a346142fa565b6000613a3e612327565b6002810180546001600160a01b0319166001600160a01b0387161790559050613a6e638a70a0eb60e01b85612ae5565b50613a7f60c0890160a08a01614b10565b6004820180546001600160a01b0319166001600160a01b0392909216919091179055613ab2610100890160e08a01614b10565b6006820180546001600160a01b0319166001600160a01b039290921691909117905562093a80600b820155613aed60e0890160c08a01614b10565b6005820180546001600160a01b0319166001600160a01b03929092169190911790556064600f820155613b2861014089016101208a01614b10565b6015820180546001600160a01b0319166001600160a01b0392909216919091179055613b5581848461430a565b60405162ee0ec160e61b81526001600160a01b037f000000000000000000000000bf844a9fcc63bb081415ef8cc3447e1cda33a3091690633b83b04090613ba2908690869060040161573c565b600060405180830381600087803b158015613bbc57600080fd5b505af1158015613bd0573d6000803e3d6000fd5b50505050613be581613be061436d565b612736565b5050505050505050565b60008061065160017f3e2bfe19d7b287e1320a5adb4ca1cb62f90c30e328e073ba40443179d690e11a615a2f565b613c278282611a85565b61090a5760405163e2517d3f60e01b81526001600160a01b038216600482015260248101839052604401610b55565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663095ea7b360e01b179052613ca784826143fc565b61125f57604080516001600160a01b038516602482015260006044808301919091528251808303909101815260649091019091526020810180516001600160e01b031663095ea7b360e01b179052613d00908590614485565b61125f8482614485565b60606000613d728384600901805480602002602001604051908101604052809291908181526020018280548015611e3a576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611e1c57505050505061353d565b91505060008082516001600160401b03811115613d9157613d91614c02565b604051908082528060200260200182016040528015613dd657816020015b6040805180820190915260008082526020820152815260200190600190039081613daf5790505b50905060005b8351811015613eae576000848281518110613df957613df96153ab565b6020026020010151604001516001811115613e1657613e1661557e565b03613ea6576040518060400160405280858381518110613e3857613e386153ab565b6020026020010151600001516001600160a01b03168152602001858381518110613e6457613e646153ab565b6020026020010151602001516001600160601b0316815250828480613e88906159b5565b955081518110613e9a57613e9a6153ab565b60200260200101819052505b600101613ddc565b509081529392505050565b604080516001600160a01b0383166024808301919091528251808303909101815260449091019091526020810180516001600160e01b031663edad0a1360e01b17905260009060028401546040519192506001600160a01b0316906382646a5890613f2b9084906000906020016157f3565b6040516020818303038152906040526040518263ffffffff1660e01b8152600401611a4e9190614e5e565b60606000613f62612327565b60098101805491925090806001600160401b03811115613f8457613f84614c02565b604051908082528060200260200182016040528015613fad578160200160208202803683370190505b5093506000805b8281101561406f576000848281548110613fd057613fd06153ab565b6000918252602090912001546001600160a01b03169050876001811115613ff957613ff961557e565b6001600160a01b038216600090815260168801602052604090205460ff1660018111156140285761402861557e565b036140665780878481518110614040576140406153ab565b6001600160a01b039092166020928302919091019091015282614062816159b5565b9350505b50600101613fb4565b50845250919392505050565b6000806000846001600160a01b031684604051614098919061570d565b6000604051808303816000865af19150503d80600081146140d5576040519150601f19603f3d011682016040523d82523d6000602084013e6140da565b606091505b50915091508180156140f95750805115806140f957506140f9816144e8565b80156125985750505050506001600160a01b03163b151590565b60607f7ac0d49c0df8eaba20287b89637474262377d7fcc1bf84c337d672a9ed663c758383604051602401614149929190615a42565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152905092915050565b6000816001600160a01b031661419a858588614506565b6001600160a01b03161495945050505050565b60006141bb85858585614548565b60028701546040519192506001600160a01b0316906382646a58906141e79084906000906020016157f3565b6040516020818303038152906040526040518263ffffffff1660e01b81526004016142129190614e5e565b600060405180830381600087803b15801561422c57600080fd5b505af1158015614240573d6000803e3d6000fd5b50505050505050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661429557604051631afcd79f60e31b815260040160405180910390fd5b565b61429f61424c565b6142b063d8a8b5c760e01b83612ae5565b506142c26378b4401360e11b84612ae5565b5061125f63d3e319af60e01b82612ae5565b6142dc61424c565b6142e46145be565b6142ef8383836145c6565b61061b600033611303565b61430261424c565b61429561468b565b600081900361432c5760405163e77c93b160e01b815260040160405180910390fd5b600a830161433b828483615aac565b507f7f63aacad63bc1693280450d5c3612ccd4efc53e46d69f3a537db102cd66290c8282604051612a7292919061573c565b6040516301753ab960e31b81524660048201526060907f000000000000000000000000bf844a9fcc63bb081415ef8cc3447e1cda33a3096001600160a01b031690630ba9d5c890602401600060405180830381865afa1580156143d4573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611f619190810190615b6b565b6000806000846001600160a01b031684604051614419919061570d565b6000604051808303816000865af19150503d8060008114614456576040519150601f19603f3d011682016040523d82523d6000602084013e61445b565b606091505b50915091508180156140f95750805115806140f95750808060200190518101906140f991906155b1565b600061449a6001600160a01b03841683614693565b905080516000141580156144bf5750808060200190518101906144bd91906155b1565b155b1561061b57604051635274afe760e01b81526001600160a01b0384166004820152602401610b55565b600081516020146144fb57506000919050565b506020015160011490565b604080516001600160a01b038086166020808401919091529085168284015282518083038401815260609092019092528051910120600090611e5290836146a1565b60607f58cc5acb742fd5eb9df21407aed105abca7d37584645d014636aa13c7b7bc387858585856040516024016145829493929190615982565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b0319909316929092179091529050949350505050565b61429561424c565b6145d763024b274760e61b84612ae5565b506145e963024b274760e61b83612ae5565b506145fb63024b274760e61b82612ae5565b5061460d63d93c394b60e01b84612ae5565b5061461f63d93c394b60e01b83612ae5565b50614631631bd8f8b560e11b84612ae5565b50614643631bd8f8b560e11b83612ae5565b50614655631bd8f8b560e11b82612ae5565b506146676362250a9560e11b84612ae5565b506146796362250a9560e11b82612ae5565b5061125f6362250a9560e11b83612ae5565b6123b161424c565b6060611f34838360006146da565b7f19457468657265756d205369676e6564204d6573736167653a0a3332000000006000908152601c839052603c8120611f34908361476d565b6060814710156146ff5760405163cd78605960e01b8152306004820152602401610b55565b600080856001600160a01b0316848660405161471b919061570d565b60006040518083038185875af1925050503d8060008114614758576040519150601f19603f3d011682016040523d82523d6000602084013e61475d565b606091505b5091509150611f2f868383614797565b60008060008061477d86866147f3565b92509250925061478d8282614840565b5090949350505050565b6060826147ac576147a7826148f9565b611f34565b81511580156147c357506001600160a01b0384163b155b156147ec57604051639996b31560e01b81526001600160a01b0385166004820152602401610b55565b5080611f34565b6000806000835160410361482d5760208401516040850151606086015160001a61481f88828585614922565b955095509550505050614839565b50508151600091506002905b9250925092565b60008260038111156148545761485461557e565b0361485d575050565b60018260038111156148715761487161557e565b0361488f5760405163f645eedf60e01b815260040160405180910390fd5b60028260038111156148a3576148a361557e565b036148c45760405163fce698f760e01b815260048101829052602401610b55565b60038260038111156148d8576148d861557e565b0361090a576040516335e2f38360e21b815260048101829052602401610b55565b8051156149095780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b600080807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084111561495d57506000915060039050826149e7565b604080516000808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa1580156149b1573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166149dd575060009250600191508290506149e7565b9250600091508190505b9450945094915050565b508054600082559060005260206000209081019061076391905b80821115614a1f5760008155600101614a0b565b5090565b600060208284031215614a3557600080fd5b81356001600160e01b031981168114611f3457600080fd5b600060208284031215614a5f57600080fd5b81356001600160401b03811115614a7557600080fd5b820160408185031215611f3457600080fd5b60008060208385031215614a9a57600080fd5b82356001600160401b0380821115614ab157600080fd5b818501915085601f830112614ac557600080fd5b813581811115614ad457600080fd5b866020606083028501011115614ae957600080fd5b60209290920196919550909350505050565b6001600160a01b038116811461076357600080fd5b600060208284031215614b2257600080fd5b8135611f3481614afb565b60008083601f840112614b3f57600080fd5b5081356001600160401b03811115614b5657600080fd5b6020830191508360208260051b850101111561375657600080fd5b63ffffffff8116811461076357600080fd5b8035614b8e81614b71565b919050565b600080600080600060808688031215614bab57600080fd5b85356001600160401b03811115614bc157600080fd5b614bcd88828901614b2d565b9096509450506020860135614be181614b71565b92506040860135614bf181614b71565b949793965091946060013592915050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b0381118282101715614c3a57614c3a614c02565b60405290565b604051601f8201601f191681016001600160401b0381118282101715614c6857614c68614c02565b604052919050565b60006001600160401b03821115614c8957614c89614c02565b5060051b60200190565b6002811061076357600080fd5b60006020808385031215614cb357600080fd5b82356001600160401b03811115614cc957600080fd5b8301601f81018513614cda57600080fd5b8035614ced614ce882614c70565b614c40565b81815260069190911b82018301908381019087831115614d0c57600080fd5b928401925b82841015614d625760408489031215614d2a5760008081fd5b614d32614c18565b8435614d3d81614afb565b815284860135614d4c81614c93565b8187015282526040939093019290840190614d11565b979650505050505050565b600060208284031215614d7f57600080fd5b5035919050565b60008060408385031215614d9957600080fd5b823591506020830135614dab81614afb565b809150509250929050565b60008151808452602080850194506020840160005b83811015614df05781516001600160a01b031687529582019590820190600101614dcb565b509495945050505050565b602081526000611f346020830184614db6565b60005b83811015614e29578181015183820152602001614e11565b50506000910152565b60008151808452614e4a816020860160208601614e0e565b601f01601f19169290920160200192915050565b602081526000611f346020830184614e32565b600082601f830112614e8257600080fd5b81356001600160401b03811115614e9b57614e9b614c02565b614eae601f8201601f1916602001614c40565b818152846020838601011115614ec357600080fd5b816020850160208301376000918101602001919091529392505050565b60008060008060808587031215614ef657600080fd5b84356001600160401b0380821115614f0d57600080fd5b818701915087601f830112614f2157600080fd5b81356020614f31614ce883614c70565b82815260069290921b8401810191818101908b841115614f5057600080fd5b948201945b83861015614f9c576040868d031215614f6e5760008081fd5b614f76614c18565b8635614f8181614afb565b81528684013584820152825260409095019490820190614f55565b9850508801359550506040870135915080821115614fb957600080fd5b50614fc687828801614e71565b925050614fd560608601614b83565b905092959194509250565b600060608284031215614ff257600080fd5b50919050565b60008083601f84011261500a57600080fd5b5081356001600160401b0381111561502157600080fd5b60208301915083602082850101111561375657600080fd5b60008060006040848603121561504e57600080fd5b83356001600160401b038082111561506557600080fd5b908501906060828803121561507957600080fd5b60405160608101818110838211171561509457615094614c02565b6040528235828111156150a657600080fd5b6150b289828601614e71565b82525060208301356020820152604083013560408201528095505060208601359150808211156150e157600080fd5b506150ee86828701614ff8565b9497909650939450505050565b6000806020838503121561510e57600080fd5b82356001600160401b0381111561512457600080fd5b61513085828601614ff8565b90969095509350505050565b803565ffffffffffff81168114614b8e57600080fd5b60008060006040848603121561516757600080fd5b83356001600160401b038082111561517e57600080fd5b908501906040828803121561519257600080fd5b61519a614c18565b6151a38361513c565b81526020830135828111156151b757600080fd5b6151c389828601614e71565b6020830152508095505060208601359150808211156150e157600080fd5b6000602082840312156151f357600080fd5b81356001600160401b0381111561520957600080fd5b82016101008185031215611f3457600080fd5b801515811461076357600080fd5b60006020828403121561523c57600080fd5b8135611f348161521c565b60008060006040848603121561525c57600080fd5b833561526781614afb565b925060208401356001600160401b0381111561528257600080fd5b6150ee86828701614b2d565b6000602082840312156152a057600080fd5b81356001600160401b038111156152b657600080fd5b82016101408185031215611f3457600080fd5b6000602082840312156152db57600080fd5b8151611f3481614afb565b6000602082840312156152f857600080fd5b611f348261513c565b6000808335601e1984360301811261531857600080fd5b8301803591506001600160401b0382111561533257600080fd5b60200191503681900382131561375657600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6001600160a01b0386811682528516602082015265ffffffffffff84166040820152608060608201819052600090614d629083018486615347565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b80820180821115610651576106516153c1565b60008261540757634e487b7160e01b600052601260045260246000fd5b500490565b60008151808452602080850194506020840160005b83811015614df057815180516001600160a01b031688528301518388015260409096019590820190600101615421565b6001600160a01b038381168252604060208084018290528451848301819052600093606092909183870190600581901b8801850189840188805b8481101561556c578b8403605f190186528251805160c0808752815190870181905260e08701918a019085905b808210156154ec57825180518e1685528c01516001600160601b03168c850152928e0192918b0191600191909101906154b8565b505050888201516001600160a01b0316868a01528b8201518682038d880152615515828261540c565b9150508a82015161552d8c88018263ffffffff169052565b5060808281015163ffffffff169087015260a091820151868203928701929092526155588183614e32565b97890197955050509186019160010161548b565b50919c9b505050505050505050505050565b634e487b7160e01b600052602160045260246000fd5b6000602082840312156155a657600080fd5b8151611f3481614c93565b6000602082840312156155c357600080fd5b8151611f348161521c565b600060018060a01b038086168352606060208401526155f06060840186614db6565b9150808416604084015250949350505050565b6000602080838503121561561657600080fd5b82516001600160401b0381111561562c57600080fd5b8301601f8101851361563d57600080fd5b805161564b614ce882614c70565b81815260059190911b8201830190838101908783111561566a57600080fd5b928401925b82841015614d6257835161568281614afb565b8252928401929084019061566f565b600181811c908216806156a557607f821691505b602082108103614ff257634e487b7160e01b600052602260045260246000fd5b600080600080608085870312156156db57600080fd5b84516156e681614b71565b60208601519094506156f781614b71565b6040860151606090960151949790965092505050565b6000825161571f818460208701614e0e565b9190910192915050565b602081526000611f34602083018461540c565b602081526000611e52602083018486615347565b60006020828403121561576257600080fd5b5051919050565b6020808252825482820181905260008481528281209092916040850190845b818110156157ad5783546001600160a01b031683526001938401939285019201615788565b50909695505050505050565b6000816157c8576157c86153c1565b506000190190565b6001600160e01b03199290921682526001600160a01b0316602082015260400190565b6040815260006158066040830185614e32565b905063ffffffff831660208301529392505050565b60018060a01b038316815260406020820152600082516060604084015261584560a0840182614e32565b90506020840151606084015260408401516080840152809150509392505050565b6001600160a01b0385811682528416602082015265ffffffffffff8316604082015260806060820181905260009061259590830184614e32565b610100810160408683376001600160a01b038581166040840152841660608301526080838184013795945050505050565b80516001600160a01b031682526020808201519083015260408101516002811061590b57634e487b7160e01b600052602160045260246000fd5b806040840152505050565b6001600160a01b0384811682526060602080840182905285518483018190526000938783019290916080870190865b81811015615968576159588387516158d1565b9483019491860191600101615945565b505080955050508086166040860152505050949350505050565b6001600160a01b0385811682526020820185905260e0820190608085604085013780841660c08401525095945050505050565b6000600182016159c7576159c76153c1565b5060010190565b60808181019083833792915050565b6001600160a01b0384811682526060602080840182905285518483018190526000938783019290916080870190865b8181101561596857615a1f8387516158d1565b9483019491860191600101615a0c565b81810381811115610651576106516153c1565b604081526000615a556040830185614e32565b90508260208301529392505050565b601f82111561061b576000816000526020600020601f850160051c81016020861015615a8d5750805b601f850160051c820191505b8181101561231f57828155600101615a99565b6001600160401b03831115615ac357615ac3614c02565b615ad783615ad18354615691565b83615a64565b6000601f841160018114615b0b5760008515615af35750838201355b600019600387901b1c1916600186901b17835561097b565b600083815260209020601f19861690835b82811015615b3c5786850135825560209485019460019092019101615b1c565b5086821015615b595760001960f88860031b161c19848701351681555b505060018560011b0183555050505050565b60006020808385031215615b7e57600080fd5b82516001600160401b03811115615b9457600080fd5b8301601f81018513615ba557600080fd5b8051615bb3614ce882614c70565b81815260069190911b82018301908381019087831115615bd257600080fd5b928401925b82841015614d625760408489031215615bf05760008081fd5b615bf8614c18565b8451615c0381614afb565b815284860151615c1281614c93565b8187015282526040939093019290840190615bd756fefe6065fb4e9872e2ad4479001655335380d83f70e163706cd65857449b84510002dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0092c9d1c07198b9465c1d2907806f1150774e0758b4edc0d272f5df87a5a7492ca2646970667358221220e3c4a9ba586831d69056cc71ff3d53806c89d4512c6e7f73129953404f66379a64736f6c63430008190033",
}

// ContractAvsGovernanceABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractAvsGovernanceMetaData.ABI instead.
var ContractAvsGovernanceABI = ContractAvsGovernanceMetaData.ABI

// ContractAvsGovernanceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractAvsGovernanceMetaData.Bin instead.
var ContractAvsGovernanceBin = ContractAvsGovernanceMetaData.Bin

// DeployContractAvsGovernance deploys a new Ethereum contract, binding an instance of ContractAvsGovernance to it.
func DeployContractAvsGovernance(auth *bind.TransactOpts, backend bind.ContractBackend, _extensionImplementation common.Address, _othenticRegistryAddress common.Address) (common.Address, *types.Transaction, *ContractAvsGovernance, error) {
	parsed, err := ContractAvsGovernanceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractAvsGovernanceBin), backend, _extensionImplementation, _othenticRegistryAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractAvsGovernance{ContractAvsGovernanceCaller: ContractAvsGovernanceCaller{contract: contract}, ContractAvsGovernanceTransactor: ContractAvsGovernanceTransactor{contract: contract}, ContractAvsGovernanceFilterer: ContractAvsGovernanceFilterer{contract: contract}}, nil
}

// ContractAvsGovernanceMethods is an auto generated interface around an Ethereum contract.
type ContractAvsGovernanceMethods interface {
	ContractAvsGovernanceCalls
	ContractAvsGovernanceTransacts
	ContractAvsGovernanceFilters
}

// ContractAvsGovernanceCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractAvsGovernanceCalls interface {
	DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error)

	EXTENSIONIMPLEMENTATION(opts *bind.CallOpts) (common.Address, error)

	AvsName(opts *bind.CallOpts) (string, error)

	AvsTreasury(opts *bind.CallOpts) (common.Address, error)

	GetIsAllowlisted(opts *bind.CallOpts) (bool, error)

	GetOperatorRestakedStrategies(opts *bind.CallOpts, _operator common.Address) ([]common.Address, error)

	GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error)

	GetRewardsReceiver(opts *bind.CallOpts, _operator common.Address) (common.Address, error)

	GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error)

	HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error)

	IsFlowPaused(opts *bind.CallOpts, _pausableFlow [4]byte) (bool, error)

	IsOperatorRegistered(opts *bind.CallOpts, operator common.Address) (bool, error)

	MaxEffectiveBalance(opts *bind.CallOpts) (*big.Int, error)

	MinStakeAmountPerStakingContract(opts *bind.CallOpts, _stakingContract common.Address) (*big.Int, error)

	MinVotingPower(opts *bind.CallOpts) (*big.Int, error)

	Multiplier(opts *bind.CallOpts, _stakingContract common.Address) (*big.Int, error)

	NumOfActiveOperators(opts *bind.CallOpts) (*big.Int, error)

	P2pAuthenticationEnabled(opts *bind.CallOpts) (bool, error)

	StakingContracts(opts *bind.CallOpts) ([]common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	VotingPower(opts *bind.CallOpts, _operator common.Address) (*big.Int, error)

	VotingPowerPerStakingContracts(opts *bind.CallOpts, _operator common.Address, _stakingContracts []common.Address) (*big.Int, error)
}

// ContractAvsGovernanceTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractAvsGovernanceTransacts interface {
	CompleteRewardsReceiverModification(opts *bind.TransactOpts) (*types.Transaction, error)

	CreateEigenRewardsSubmission(opts *bind.TransactOpts, _operators []common.Address, _startTimestamp uint32, _duration uint32, _totalAmount *big.Int) (*types.Transaction, error)

	CreateOperatorDirectedAVSRewardsSubmission(opts *bind.TransactOpts, _operators []IRewardsCoordinatorOperatorReward, _lastPayedTask *big.Int, _rewardsData []byte, _remoteEid uint32) (*types.Transaction, error)

	EjectOperatorFromNetwork(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error)

	GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts, _initializationParams IAvsGovernanceInitializationParams) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts, _pausableFlow [4]byte) (*types.Transaction, error)

	QueueRewardsReceiverModification(opts *bind.TransactOpts, _newRewardsReceiver common.Address) (*types.Transaction, error)

	RegisterAsOperator(opts *bind.TransactOpts, _operatorRegistrationParams IAvsGovernanceOperatorRegistrationParams) (*types.Transaction, error)

	RegisterAvsToEigenLayer(opts *bind.TransactOpts, metadataURI string) (*types.Transaction, error)

	RegisterAvsToSymbiotic(opts *bind.TransactOpts) (*types.Transaction, error)

	RegisterOperatorToEigenLayer(opts *bind.TransactOpts, _eigenSig ISignatureUtilsSignatureWithSaltAndExpiry, _authToken []byte) (*types.Transaction, error)

	RegisterOperatorToSymbiotic(opts *bind.TransactOpts, _symbioticSig IAvsGovernanceSymbioticOptInSignature, _authToken []byte) (*types.Transaction, error)

	RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error)

	RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	SetP2pAuthenticationEnabled(opts *bind.TransactOpts, _p2pAuthenticationEnabled bool) (*types.Transaction, error)

	SetStakingContractMultiplier(opts *bind.TransactOpts, _votingPowerMultiplier IAvsGovernanceVotingPowerMultiplier) (*types.Transaction, error)

	SetStakingContractMultiplierBatch(opts *bind.TransactOpts, _votingPowerMultipliers []IAvsGovernanceVotingPowerMultiplier) (*types.Transaction, error)

	SetSupportedStakingContracts(opts *bind.TransactOpts, _stakingContractsDetails []IAvsGovernanceStakingContractInfo) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts, _pausableFlow [4]byte) (*types.Transaction, error)

	UnregisterAsOperatorFromAvs(opts *bind.TransactOpts) (*types.Transaction, error)

	UnregisterAsOperatorFromEigenLayer(opts *bind.TransactOpts) (*types.Transaction, error)

	UnregisterAsOperatorFromSymbiotic(opts *bind.TransactOpts, _unregistrationSignature IAvsGovernanceSymbioticOptOutSignature) (*types.Transaction, error)
}

// ContractAvsGovernanceFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractAvsGovernanceFilters interface {
	FilterDepositRewardsBackFailed(opts *bind.FilterOpts) (*ContractAvsGovernanceDepositRewardsBackFailedIterator, error)
	WatchDepositRewardsBackFailed(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceDepositRewardsBackFailed) (event.Subscription, error)
	ParseDepositRewardsBackFailed(log types.Log) (*ContractAvsGovernanceDepositRewardsBackFailed, error)

	FilterFlowPaused(opts *bind.FilterOpts) (*ContractAvsGovernanceFlowPausedIterator, error)
	WatchFlowPaused(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceFlowPaused) (event.Subscription, error)
	ParseFlowPaused(log types.Log) (*ContractAvsGovernanceFlowPaused, error)

	FilterFlowUnpaused(opts *bind.FilterOpts) (*ContractAvsGovernanceFlowUnpausedIterator, error)
	WatchFlowUnpaused(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceFlowUnpaused) (event.Subscription, error)
	ParseFlowUnpaused(log types.Log) (*ContractAvsGovernanceFlowUnpaused, error)

	FilterInitialized(opts *bind.FilterOpts) (*ContractAvsGovernanceInitializedIterator, error)
	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceInitialized) (event.Subscription, error)
	ParseInitialized(log types.Log) (*ContractAvsGovernanceInitialized, error)

	FilterNativeCoinNotSupportedForEigenRewards(opts *bind.FilterOpts) (*ContractAvsGovernanceNativeCoinNotSupportedForEigenRewardsIterator, error)
	WatchNativeCoinNotSupportedForEigenRewards(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceNativeCoinNotSupportedForEigenRewards) (event.Subscription, error)
	ParseNativeCoinNotSupportedForEigenRewards(log types.Log) (*ContractAvsGovernanceNativeCoinNotSupportedForEigenRewards, error)

	FilterOperatorEjectedFromNetwork(opts *bind.FilterOpts) (*ContractAvsGovernanceOperatorEjectedFromNetworkIterator, error)
	WatchOperatorEjectedFromNetwork(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceOperatorEjectedFromNetwork) (event.Subscription, error)
	ParseOperatorEjectedFromNetwork(log types.Log) (*ContractAvsGovernanceOperatorEjectedFromNetwork, error)

	FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address) (*ContractAvsGovernanceOperatorRegisteredIterator, error)
	WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceOperatorRegistered, operator []common.Address) (event.Subscription, error)
	ParseOperatorRegistered(log types.Log) (*ContractAvsGovernanceOperatorRegistered, error)

	FilterOperatorRegisteredToEigenLayer(opts *bind.FilterOpts) (*ContractAvsGovernanceOperatorRegisteredToEigenLayerIterator, error)
	WatchOperatorRegisteredToEigenLayer(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceOperatorRegisteredToEigenLayer) (event.Subscription, error)
	ParseOperatorRegisteredToEigenLayer(log types.Log) (*ContractAvsGovernanceOperatorRegisteredToEigenLayer, error)

	FilterOperatorRegisteredToSymbiotic(opts *bind.FilterOpts) (*ContractAvsGovernanceOperatorRegisteredToSymbioticIterator, error)
	WatchOperatorRegisteredToSymbiotic(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceOperatorRegisteredToSymbiotic) (event.Subscription, error)
	ParseOperatorRegisteredToSymbiotic(log types.Log) (*ContractAvsGovernanceOperatorRegisteredToSymbiotic, error)

	FilterOperatorUnregistered(opts *bind.FilterOpts, operator []common.Address) (*ContractAvsGovernanceOperatorUnregisteredIterator, error)
	WatchOperatorUnregistered(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceOperatorUnregistered, operator []common.Address) (event.Subscription, error)
	ParseOperatorUnregistered(log types.Log) (*ContractAvsGovernanceOperatorUnregistered, error)

	FilterOperatorUnregisteredToEigenLayer(opts *bind.FilterOpts) (*ContractAvsGovernanceOperatorUnregisteredToEigenLayerIterator, error)
	WatchOperatorUnregisteredToEigenLayer(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceOperatorUnregisteredToEigenLayer) (event.Subscription, error)
	ParseOperatorUnregisteredToEigenLayer(log types.Log) (*ContractAvsGovernanceOperatorUnregisteredToEigenLayer, error)

	FilterOperatorUnregisteredToSymbiotic(opts *bind.FilterOpts) (*ContractAvsGovernanceOperatorUnregisteredToSymbioticIterator, error)
	WatchOperatorUnregisteredToSymbiotic(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceOperatorUnregisteredToSymbiotic) (event.Subscription, error)
	ParseOperatorUnregisteredToSymbiotic(log types.Log) (*ContractAvsGovernanceOperatorUnregisteredToSymbiotic, error)

	FilterQueuedRewardsReceiverModification(opts *bind.FilterOpts, operator []common.Address) (*ContractAvsGovernanceQueuedRewardsReceiverModificationIterator, error)
	WatchQueuedRewardsReceiverModification(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceQueuedRewardsReceiverModification, operator []common.Address) (event.Subscription, error)
	ParseQueuedRewardsReceiverModification(log types.Log) (*ContractAvsGovernanceQueuedRewardsReceiverModification, error)

	FilterRewardsCoordinatorReverted(opts *bind.FilterOpts) (*ContractAvsGovernanceRewardsCoordinatorRevertedIterator, error)
	WatchRewardsCoordinatorReverted(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceRewardsCoordinatorReverted) (event.Subscription, error)
	ParseRewardsCoordinatorReverted(log types.Log) (*ContractAvsGovernanceRewardsCoordinatorReverted, error)

	FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ContractAvsGovernanceRoleAdminChangedIterator, error)
	WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error)
	ParseRoleAdminChanged(log types.Log) (*ContractAvsGovernanceRoleAdminChanged, error)

	FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ContractAvsGovernanceRoleGrantedIterator, error)
	WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error)
	ParseRoleGranted(log types.Log) (*ContractAvsGovernanceRoleGranted, error)

	FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ContractAvsGovernanceRoleRevokedIterator, error)
	WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error)
	ParseRoleRevoked(log types.Log) (*ContractAvsGovernanceRoleRevoked, error)

	FilterSetAvsName(opts *bind.FilterOpts) (*ContractAvsGovernanceSetAvsNameIterator, error)
	WatchSetAvsName(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetAvsName) (event.Subscription, error)
	ParseSetAvsName(log types.Log) (*ContractAvsGovernanceSetAvsName, error)

	FilterSetIsAllowlisted(opts *bind.FilterOpts) (*ContractAvsGovernanceSetIsAllowlistedIterator, error)
	WatchSetIsAllowlisted(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetIsAllowlisted) (event.Subscription, error)
	ParseSetIsAllowlisted(log types.Log) (*ContractAvsGovernanceSetIsAllowlisted, error)

	FilterSetP2pAuthenticationEnabled(opts *bind.FilterOpts) (*ContractAvsGovernanceSetP2pAuthenticationEnabledIterator, error)
	WatchSetP2pAuthenticationEnabled(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetP2pAuthenticationEnabled) (event.Subscription, error)
	ParseSetP2pAuthenticationEnabled(log types.Log) (*ContractAvsGovernanceSetP2pAuthenticationEnabled, error)

	FilterSetRewardsReceiver(opts *bind.FilterOpts, operator []common.Address) (*ContractAvsGovernanceSetRewardsReceiverIterator, error)
	WatchSetRewardsReceiver(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetRewardsReceiver, operator []common.Address) (event.Subscription, error)
	ParseSetRewardsReceiver(log types.Log) (*ContractAvsGovernanceSetRewardsReceiver, error)

	FilterSetStakingContractMultiplier(opts *bind.FilterOpts) (*ContractAvsGovernanceSetStakingContractMultiplierIterator, error)
	WatchSetStakingContractMultiplier(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetStakingContractMultiplier) (event.Subscription, error)
	ParseSetStakingContractMultiplier(log types.Log) (*ContractAvsGovernanceSetStakingContractMultiplier, error)

	FilterSetToken(opts *bind.FilterOpts) (*ContractAvsGovernanceSetTokenIterator, error)
	WatchSetToken(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetToken) (event.Subscription, error)
	ParseSetToken(log types.Log) (*ContractAvsGovernanceSetToken, error)

	FilterSetNewSupportedStakingContracts(opts *bind.FilterOpts) (*ContractAvsGovernanceSetNewSupportedStakingContractsIterator, error)
	WatchSetNewSupportedStakingContracts(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetNewSupportedStakingContracts) (event.Subscription, error)
	ParseSetNewSupportedStakingContracts(log types.Log) (*ContractAvsGovernanceSetNewSupportedStakingContracts, error)
}

// ContractAvsGovernance is an auto generated Go binding around an Ethereum contract.
type ContractAvsGovernance struct {
	ContractAvsGovernanceCaller     // Read-only binding to the contract
	ContractAvsGovernanceTransactor // Write-only binding to the contract
	ContractAvsGovernanceFilterer   // Log filterer for contract events
}

// ContractAvsGovernance implements the ContractAvsGovernanceMethods interface.
var _ ContractAvsGovernanceMethods = (*ContractAvsGovernance)(nil)

// ContractAvsGovernanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractAvsGovernanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAvsGovernanceCaller implements the ContractAvsGovernanceCalls interface.
var _ ContractAvsGovernanceCalls = (*ContractAvsGovernanceCaller)(nil)

// ContractAvsGovernanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractAvsGovernanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAvsGovernanceTransactor implements the ContractAvsGovernanceTransacts interface.
var _ ContractAvsGovernanceTransacts = (*ContractAvsGovernanceTransactor)(nil)

// ContractAvsGovernanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractAvsGovernanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAvsGovernanceFilterer implements the ContractAvsGovernanceFilters interface.
var _ ContractAvsGovernanceFilters = (*ContractAvsGovernanceFilterer)(nil)

// ContractAvsGovernanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractAvsGovernanceSession struct {
	Contract     *ContractAvsGovernance // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ContractAvsGovernanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractAvsGovernanceCallerSession struct {
	Contract *ContractAvsGovernanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// ContractAvsGovernanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractAvsGovernanceTransactorSession struct {
	Contract     *ContractAvsGovernanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ContractAvsGovernanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractAvsGovernanceRaw struct {
	Contract *ContractAvsGovernance // Generic contract binding to access the raw methods on
}

// ContractAvsGovernanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractAvsGovernanceCallerRaw struct {
	Contract *ContractAvsGovernanceCaller // Generic read-only contract binding to access the raw methods on
}

// ContractAvsGovernanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractAvsGovernanceTransactorRaw struct {
	Contract *ContractAvsGovernanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractAvsGovernance creates a new instance of ContractAvsGovernance, bound to a specific deployed contract.
func NewContractAvsGovernance(address common.Address, backend bind.ContractBackend) (*ContractAvsGovernance, error) {
	contract, err := bindContractAvsGovernance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernance{ContractAvsGovernanceCaller: ContractAvsGovernanceCaller{contract: contract}, ContractAvsGovernanceTransactor: ContractAvsGovernanceTransactor{contract: contract}, ContractAvsGovernanceFilterer: ContractAvsGovernanceFilterer{contract: contract}}, nil
}

// NewContractAvsGovernanceCaller creates a new read-only instance of ContractAvsGovernance, bound to a specific deployed contract.
func NewContractAvsGovernanceCaller(address common.Address, caller bind.ContractCaller) (*ContractAvsGovernanceCaller, error) {
	contract, err := bindContractAvsGovernance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceCaller{contract: contract}, nil
}

// NewContractAvsGovernanceTransactor creates a new write-only instance of ContractAvsGovernance, bound to a specific deployed contract.
func NewContractAvsGovernanceTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractAvsGovernanceTransactor, error) {
	contract, err := bindContractAvsGovernance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceTransactor{contract: contract}, nil
}

// NewContractAvsGovernanceFilterer creates a new log filterer instance of ContractAvsGovernance, bound to a specific deployed contract.
func NewContractAvsGovernanceFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractAvsGovernanceFilterer, error) {
	contract, err := bindContractAvsGovernance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceFilterer{contract: contract}, nil
}

// bindContractAvsGovernance binds a generic wrapper to an already deployed contract.
func bindContractAvsGovernance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractAvsGovernanceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractAvsGovernance *ContractAvsGovernanceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractAvsGovernance.Contract.ContractAvsGovernanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractAvsGovernance *ContractAvsGovernanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.ContractAvsGovernanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractAvsGovernance *ContractAvsGovernanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.ContractAvsGovernanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractAvsGovernance *ContractAvsGovernanceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractAvsGovernance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ContractAvsGovernance.Contract.DEFAULTADMINROLE(&_ContractAvsGovernance.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ContractAvsGovernance.Contract.DEFAULTADMINROLE(&_ContractAvsGovernance.CallOpts)
}

// EXTENSIONIMPLEMENTATION is a free data retrieval call binding the contract method 0x226def04.
//
// Solidity: function EXTENSION_IMPLEMENTATION() view returns(address)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) EXTENSIONIMPLEMENTATION(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "EXTENSION_IMPLEMENTATION")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EXTENSIONIMPLEMENTATION is a free data retrieval call binding the contract method 0x226def04.
//
// Solidity: function EXTENSION_IMPLEMENTATION() view returns(address)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) EXTENSIONIMPLEMENTATION() (common.Address, error) {
	return _ContractAvsGovernance.Contract.EXTENSIONIMPLEMENTATION(&_ContractAvsGovernance.CallOpts)
}

// EXTENSIONIMPLEMENTATION is a free data retrieval call binding the contract method 0x226def04.
//
// Solidity: function EXTENSION_IMPLEMENTATION() view returns(address)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) EXTENSIONIMPLEMENTATION() (common.Address, error) {
	return _ContractAvsGovernance.Contract.EXTENSIONIMPLEMENTATION(&_ContractAvsGovernance.CallOpts)
}

// AvsName is a free data retrieval call binding the contract method 0x41b92a29.
//
// Solidity: function avsName() view returns(string)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) AvsName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "avsName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AvsName is a free data retrieval call binding the contract method 0x41b92a29.
//
// Solidity: function avsName() view returns(string)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) AvsName() (string, error) {
	return _ContractAvsGovernance.Contract.AvsName(&_ContractAvsGovernance.CallOpts)
}

// AvsName is a free data retrieval call binding the contract method 0x41b92a29.
//
// Solidity: function avsName() view returns(string)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) AvsName() (string, error) {
	return _ContractAvsGovernance.Contract.AvsName(&_ContractAvsGovernance.CallOpts)
}

// AvsTreasury is a free data retrieval call binding the contract method 0x1246193e.
//
// Solidity: function avsTreasury() view returns(address)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) AvsTreasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "avsTreasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsTreasury is a free data retrieval call binding the contract method 0x1246193e.
//
// Solidity: function avsTreasury() view returns(address)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) AvsTreasury() (common.Address, error) {
	return _ContractAvsGovernance.Contract.AvsTreasury(&_ContractAvsGovernance.CallOpts)
}

// AvsTreasury is a free data retrieval call binding the contract method 0x1246193e.
//
// Solidity: function avsTreasury() view returns(address)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) AvsTreasury() (common.Address, error) {
	return _ContractAvsGovernance.Contract.AvsTreasury(&_ContractAvsGovernance.CallOpts)
}

// GetIsAllowlisted is a free data retrieval call binding the contract method 0xb525fa88.
//
// Solidity: function getIsAllowlisted() view returns(bool)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) GetIsAllowlisted(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "getIsAllowlisted")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetIsAllowlisted is a free data retrieval call binding the contract method 0xb525fa88.
//
// Solidity: function getIsAllowlisted() view returns(bool)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) GetIsAllowlisted() (bool, error) {
	return _ContractAvsGovernance.Contract.GetIsAllowlisted(&_ContractAvsGovernance.CallOpts)
}

// GetIsAllowlisted is a free data retrieval call binding the contract method 0xb525fa88.
//
// Solidity: function getIsAllowlisted() view returns(bool)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) GetIsAllowlisted() (bool, error) {
	return _ContractAvsGovernance.Contract.GetIsAllowlisted(&_ContractAvsGovernance.CallOpts)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address _operator) view returns(address[])
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) GetOperatorRestakedStrategies(opts *bind.CallOpts, _operator common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "getOperatorRestakedStrategies", _operator)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address _operator) view returns(address[])
func (_ContractAvsGovernance *ContractAvsGovernanceSession) GetOperatorRestakedStrategies(_operator common.Address) ([]common.Address, error) {
	return _ContractAvsGovernance.Contract.GetOperatorRestakedStrategies(&_ContractAvsGovernance.CallOpts, _operator)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address _operator) view returns(address[])
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) GetOperatorRestakedStrategies(_operator common.Address) ([]common.Address, error) {
	return _ContractAvsGovernance.Contract.GetOperatorRestakedStrategies(&_ContractAvsGovernance.CallOpts, _operator)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "getRestakeableStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractAvsGovernance *ContractAvsGovernanceSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractAvsGovernance.Contract.GetRestakeableStrategies(&_ContractAvsGovernance.CallOpts)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractAvsGovernance.Contract.GetRestakeableStrategies(&_ContractAvsGovernance.CallOpts)
}

// GetRewardsReceiver is a free data retrieval call binding the contract method 0x5e95cee2.
//
// Solidity: function getRewardsReceiver(address _operator) view returns(address)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) GetRewardsReceiver(opts *bind.CallOpts, _operator common.Address) (common.Address, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "getRewardsReceiver", _operator)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRewardsReceiver is a free data retrieval call binding the contract method 0x5e95cee2.
//
// Solidity: function getRewardsReceiver(address _operator) view returns(address)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) GetRewardsReceiver(_operator common.Address) (common.Address, error) {
	return _ContractAvsGovernance.Contract.GetRewardsReceiver(&_ContractAvsGovernance.CallOpts, _operator)
}

// GetRewardsReceiver is a free data retrieval call binding the contract method 0x5e95cee2.
//
// Solidity: function getRewardsReceiver(address _operator) view returns(address)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) GetRewardsReceiver(_operator common.Address) (common.Address, error) {
	return _ContractAvsGovernance.Contract.GetRewardsReceiver(&_ContractAvsGovernance.CallOpts, _operator)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ContractAvsGovernance.Contract.GetRoleAdmin(&_ContractAvsGovernance.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ContractAvsGovernance.Contract.GetRoleAdmin(&_ContractAvsGovernance.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ContractAvsGovernance.Contract.HasRole(&_ContractAvsGovernance.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ContractAvsGovernance.Contract.HasRole(&_ContractAvsGovernance.CallOpts, role, account)
}

// IsFlowPaused is a free data retrieval call binding the contract method 0xefd96978.
//
// Solidity: function isFlowPaused(bytes4 _pausableFlow) view returns(bool _isPaused)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) IsFlowPaused(opts *bind.CallOpts, _pausableFlow [4]byte) (bool, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "isFlowPaused", _pausableFlow)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFlowPaused is a free data retrieval call binding the contract method 0xefd96978.
//
// Solidity: function isFlowPaused(bytes4 _pausableFlow) view returns(bool _isPaused)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) IsFlowPaused(_pausableFlow [4]byte) (bool, error) {
	return _ContractAvsGovernance.Contract.IsFlowPaused(&_ContractAvsGovernance.CallOpts, _pausableFlow)
}

// IsFlowPaused is a free data retrieval call binding the contract method 0xefd96978.
//
// Solidity: function isFlowPaused(bytes4 _pausableFlow) view returns(bool _isPaused)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) IsFlowPaused(_pausableFlow [4]byte) (bool, error) {
	return _ContractAvsGovernance.Contract.IsFlowPaused(&_ContractAvsGovernance.CallOpts, _pausableFlow)
}

// IsOperatorRegistered is a free data retrieval call binding the contract method 0x6b1906f8.
//
// Solidity: function isOperatorRegistered(address operator) view returns(bool)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) IsOperatorRegistered(opts *bind.CallOpts, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "isOperatorRegistered", operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorRegistered is a free data retrieval call binding the contract method 0x6b1906f8.
//
// Solidity: function isOperatorRegistered(address operator) view returns(bool)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) IsOperatorRegistered(operator common.Address) (bool, error) {
	return _ContractAvsGovernance.Contract.IsOperatorRegistered(&_ContractAvsGovernance.CallOpts, operator)
}

// IsOperatorRegistered is a free data retrieval call binding the contract method 0x6b1906f8.
//
// Solidity: function isOperatorRegistered(address operator) view returns(bool)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) IsOperatorRegistered(operator common.Address) (bool, error) {
	return _ContractAvsGovernance.Contract.IsOperatorRegistered(&_ContractAvsGovernance.CallOpts, operator)
}

// MaxEffectiveBalance is a free data retrieval call binding the contract method 0xa88171ee.
//
// Solidity: function maxEffectiveBalance() view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) MaxEffectiveBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "maxEffectiveBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxEffectiveBalance is a free data retrieval call binding the contract method 0xa88171ee.
//
// Solidity: function maxEffectiveBalance() view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) MaxEffectiveBalance() (*big.Int, error) {
	return _ContractAvsGovernance.Contract.MaxEffectiveBalance(&_ContractAvsGovernance.CallOpts)
}

// MaxEffectiveBalance is a free data retrieval call binding the contract method 0xa88171ee.
//
// Solidity: function maxEffectiveBalance() view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) MaxEffectiveBalance() (*big.Int, error) {
	return _ContractAvsGovernance.Contract.MaxEffectiveBalance(&_ContractAvsGovernance.CallOpts)
}

// MinStakeAmountPerStakingContract is a free data retrieval call binding the contract method 0x13fe4521.
//
// Solidity: function minStakeAmountPerStakingContract(address _stakingContract) view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) MinStakeAmountPerStakingContract(opts *bind.CallOpts, _stakingContract common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "minStakeAmountPerStakingContract", _stakingContract)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStakeAmountPerStakingContract is a free data retrieval call binding the contract method 0x13fe4521.
//
// Solidity: function minStakeAmountPerStakingContract(address _stakingContract) view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) MinStakeAmountPerStakingContract(_stakingContract common.Address) (*big.Int, error) {
	return _ContractAvsGovernance.Contract.MinStakeAmountPerStakingContract(&_ContractAvsGovernance.CallOpts, _stakingContract)
}

// MinStakeAmountPerStakingContract is a free data retrieval call binding the contract method 0x13fe4521.
//
// Solidity: function minStakeAmountPerStakingContract(address _stakingContract) view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) MinStakeAmountPerStakingContract(_stakingContract common.Address) (*big.Int, error) {
	return _ContractAvsGovernance.Contract.MinStakeAmountPerStakingContract(&_ContractAvsGovernance.CallOpts, _stakingContract)
}

// MinVotingPower is a free data retrieval call binding the contract method 0x36fffde0.
//
// Solidity: function minVotingPower() view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) MinVotingPower(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "minVotingPower")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinVotingPower is a free data retrieval call binding the contract method 0x36fffde0.
//
// Solidity: function minVotingPower() view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) MinVotingPower() (*big.Int, error) {
	return _ContractAvsGovernance.Contract.MinVotingPower(&_ContractAvsGovernance.CallOpts)
}

// MinVotingPower is a free data retrieval call binding the contract method 0x36fffde0.
//
// Solidity: function minVotingPower() view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) MinVotingPower() (*big.Int, error) {
	return _ContractAvsGovernance.Contract.MinVotingPower(&_ContractAvsGovernance.CallOpts)
}

// Multiplier is a free data retrieval call binding the contract method 0xa9b3f8b7.
//
// Solidity: function multiplier(address _stakingContract) view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) Multiplier(opts *bind.CallOpts, _stakingContract common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "multiplier", _stakingContract)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Multiplier is a free data retrieval call binding the contract method 0xa9b3f8b7.
//
// Solidity: function multiplier(address _stakingContract) view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) Multiplier(_stakingContract common.Address) (*big.Int, error) {
	return _ContractAvsGovernance.Contract.Multiplier(&_ContractAvsGovernance.CallOpts, _stakingContract)
}

// Multiplier is a free data retrieval call binding the contract method 0xa9b3f8b7.
//
// Solidity: function multiplier(address _stakingContract) view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) Multiplier(_stakingContract common.Address) (*big.Int, error) {
	return _ContractAvsGovernance.Contract.Multiplier(&_ContractAvsGovernance.CallOpts, _stakingContract)
}

// NumOfActiveOperators is a free data retrieval call binding the contract method 0x7897dec3.
//
// Solidity: function numOfActiveOperators() view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) NumOfActiveOperators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "numOfActiveOperators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumOfActiveOperators is a free data retrieval call binding the contract method 0x7897dec3.
//
// Solidity: function numOfActiveOperators() view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) NumOfActiveOperators() (*big.Int, error) {
	return _ContractAvsGovernance.Contract.NumOfActiveOperators(&_ContractAvsGovernance.CallOpts)
}

// NumOfActiveOperators is a free data retrieval call binding the contract method 0x7897dec3.
//
// Solidity: function numOfActiveOperators() view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) NumOfActiveOperators() (*big.Int, error) {
	return _ContractAvsGovernance.Contract.NumOfActiveOperators(&_ContractAvsGovernance.CallOpts)
}

// P2pAuthenticationEnabled is a free data retrieval call binding the contract method 0xcab1a5ef.
//
// Solidity: function p2pAuthenticationEnabled() view returns(bool)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) P2pAuthenticationEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "p2pAuthenticationEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// P2pAuthenticationEnabled is a free data retrieval call binding the contract method 0xcab1a5ef.
//
// Solidity: function p2pAuthenticationEnabled() view returns(bool)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) P2pAuthenticationEnabled() (bool, error) {
	return _ContractAvsGovernance.Contract.P2pAuthenticationEnabled(&_ContractAvsGovernance.CallOpts)
}

// P2pAuthenticationEnabled is a free data retrieval call binding the contract method 0xcab1a5ef.
//
// Solidity: function p2pAuthenticationEnabled() view returns(bool)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) P2pAuthenticationEnabled() (bool, error) {
	return _ContractAvsGovernance.Contract.P2pAuthenticationEnabled(&_ContractAvsGovernance.CallOpts)
}

// StakingContracts is a free data retrieval call binding the contract method 0x560d5484.
//
// Solidity: function stakingContracts() view returns(address[])
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) StakingContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "stakingContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// StakingContracts is a free data retrieval call binding the contract method 0x560d5484.
//
// Solidity: function stakingContracts() view returns(address[])
func (_ContractAvsGovernance *ContractAvsGovernanceSession) StakingContracts() ([]common.Address, error) {
	return _ContractAvsGovernance.Contract.StakingContracts(&_ContractAvsGovernance.CallOpts)
}

// StakingContracts is a free data retrieval call binding the contract method 0x560d5484.
//
// Solidity: function stakingContracts() view returns(address[])
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) StakingContracts() ([]common.Address, error) {
	return _ContractAvsGovernance.Contract.StakingContracts(&_ContractAvsGovernance.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ContractAvsGovernance.Contract.SupportsInterface(&_ContractAvsGovernance.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ContractAvsGovernance.Contract.SupportsInterface(&_ContractAvsGovernance.CallOpts, interfaceId)
}

// VotingPower is a free data retrieval call binding the contract method 0xc07473f6.
//
// Solidity: function votingPower(address _operator) view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) VotingPower(opts *bind.CallOpts, _operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "votingPower", _operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingPower is a free data retrieval call binding the contract method 0xc07473f6.
//
// Solidity: function votingPower(address _operator) view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) VotingPower(_operator common.Address) (*big.Int, error) {
	return _ContractAvsGovernance.Contract.VotingPower(&_ContractAvsGovernance.CallOpts, _operator)
}

// VotingPower is a free data retrieval call binding the contract method 0xc07473f6.
//
// Solidity: function votingPower(address _operator) view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) VotingPower(_operator common.Address) (*big.Int, error) {
	return _ContractAvsGovernance.Contract.VotingPower(&_ContractAvsGovernance.CallOpts, _operator)
}

// VotingPowerPerStakingContracts is a free data retrieval call binding the contract method 0xcbd969c9.
//
// Solidity: function votingPowerPerStakingContracts(address _operator, address[] _stakingContracts) view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) VotingPowerPerStakingContracts(opts *bind.CallOpts, _operator common.Address, _stakingContracts []common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "votingPowerPerStakingContracts", _operator, _stakingContracts)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingPowerPerStakingContracts is a free data retrieval call binding the contract method 0xcbd969c9.
//
// Solidity: function votingPowerPerStakingContracts(address _operator, address[] _stakingContracts) view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) VotingPowerPerStakingContracts(_operator common.Address, _stakingContracts []common.Address) (*big.Int, error) {
	return _ContractAvsGovernance.Contract.VotingPowerPerStakingContracts(&_ContractAvsGovernance.CallOpts, _operator, _stakingContracts)
}

// VotingPowerPerStakingContracts is a free data retrieval call binding the contract method 0xcbd969c9.
//
// Solidity: function votingPowerPerStakingContracts(address _operator, address[] _stakingContracts) view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) VotingPowerPerStakingContracts(_operator common.Address, _stakingContracts []common.Address) (*big.Int, error) {
	return _ContractAvsGovernance.Contract.VotingPowerPerStakingContracts(&_ContractAvsGovernance.CallOpts, _operator, _stakingContracts)
}

// CompleteRewardsReceiverModification is a paid mutator transaction binding the contract method 0xe6474b0f.
//
// Solidity: function completeRewardsReceiverModification() returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) CompleteRewardsReceiverModification(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "completeRewardsReceiverModification")
}

// CompleteRewardsReceiverModification is a paid mutator transaction binding the contract method 0xe6474b0f.
//
// Solidity: function completeRewardsReceiverModification() returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) CompleteRewardsReceiverModification() (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.CompleteRewardsReceiverModification(&_ContractAvsGovernance.TransactOpts)
}

// CompleteRewardsReceiverModification is a paid mutator transaction binding the contract method 0xe6474b0f.
//
// Solidity: function completeRewardsReceiverModification() returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) CompleteRewardsReceiverModification() (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.CompleteRewardsReceiverModification(&_ContractAvsGovernance.TransactOpts)
}

// CreateEigenRewardsSubmission is a paid mutator transaction binding the contract method 0x1b409bf2.
//
// Solidity: function createEigenRewardsSubmission(address[] _operators, uint32 _startTimestamp, uint32 _duration, uint256 _totalAmount) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) CreateEigenRewardsSubmission(opts *bind.TransactOpts, _operators []common.Address, _startTimestamp uint32, _duration uint32, _totalAmount *big.Int) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "createEigenRewardsSubmission", _operators, _startTimestamp, _duration, _totalAmount)
}

// CreateEigenRewardsSubmission is a paid mutator transaction binding the contract method 0x1b409bf2.
//
// Solidity: function createEigenRewardsSubmission(address[] _operators, uint32 _startTimestamp, uint32 _duration, uint256 _totalAmount) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) CreateEigenRewardsSubmission(_operators []common.Address, _startTimestamp uint32, _duration uint32, _totalAmount *big.Int) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.CreateEigenRewardsSubmission(&_ContractAvsGovernance.TransactOpts, _operators, _startTimestamp, _duration, _totalAmount)
}

// CreateEigenRewardsSubmission is a paid mutator transaction binding the contract method 0x1b409bf2.
//
// Solidity: function createEigenRewardsSubmission(address[] _operators, uint32 _startTimestamp, uint32 _duration, uint256 _totalAmount) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) CreateEigenRewardsSubmission(_operators []common.Address, _startTimestamp uint32, _duration uint32, _totalAmount *big.Int) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.CreateEigenRewardsSubmission(&_ContractAvsGovernance.TransactOpts, _operators, _startTimestamp, _duration, _totalAmount)
}

// CreateOperatorDirectedAVSRewardsSubmission is a paid mutator transaction binding the contract method 0x62de7393.
//
// Solidity: function createOperatorDirectedAVSRewardsSubmission((address,uint256)[] _operators, uint256 _lastPayedTask, bytes _rewardsData, uint32 _remoteEid) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) CreateOperatorDirectedAVSRewardsSubmission(opts *bind.TransactOpts, _operators []IRewardsCoordinatorOperatorReward, _lastPayedTask *big.Int, _rewardsData []byte, _remoteEid uint32) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "createOperatorDirectedAVSRewardsSubmission", _operators, _lastPayedTask, _rewardsData, _remoteEid)
}

// CreateOperatorDirectedAVSRewardsSubmission is a paid mutator transaction binding the contract method 0x62de7393.
//
// Solidity: function createOperatorDirectedAVSRewardsSubmission((address,uint256)[] _operators, uint256 _lastPayedTask, bytes _rewardsData, uint32 _remoteEid) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) CreateOperatorDirectedAVSRewardsSubmission(_operators []IRewardsCoordinatorOperatorReward, _lastPayedTask *big.Int, _rewardsData []byte, _remoteEid uint32) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.CreateOperatorDirectedAVSRewardsSubmission(&_ContractAvsGovernance.TransactOpts, _operators, _lastPayedTask, _rewardsData, _remoteEid)
}

// CreateOperatorDirectedAVSRewardsSubmission is a paid mutator transaction binding the contract method 0x62de7393.
//
// Solidity: function createOperatorDirectedAVSRewardsSubmission((address,uint256)[] _operators, uint256 _lastPayedTask, bytes _rewardsData, uint32 _remoteEid) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) CreateOperatorDirectedAVSRewardsSubmission(_operators []IRewardsCoordinatorOperatorReward, _lastPayedTask *big.Int, _rewardsData []byte, _remoteEid uint32) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.CreateOperatorDirectedAVSRewardsSubmission(&_ContractAvsGovernance.TransactOpts, _operators, _lastPayedTask, _rewardsData, _remoteEid)
}

// EjectOperatorFromNetwork is a paid mutator transaction binding the contract method 0xf6e258cc.
//
// Solidity: function ejectOperatorFromNetwork(address _operator) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) EjectOperatorFromNetwork(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "ejectOperatorFromNetwork", _operator)
}

// EjectOperatorFromNetwork is a paid mutator transaction binding the contract method 0xf6e258cc.
//
// Solidity: function ejectOperatorFromNetwork(address _operator) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) EjectOperatorFromNetwork(_operator common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.EjectOperatorFromNetwork(&_ContractAvsGovernance.TransactOpts, _operator)
}

// EjectOperatorFromNetwork is a paid mutator transaction binding the contract method 0xf6e258cc.
//
// Solidity: function ejectOperatorFromNetwork(address _operator) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) EjectOperatorFromNetwork(_operator common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.EjectOperatorFromNetwork(&_ContractAvsGovernance.TransactOpts, _operator)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.GrantRole(&_ContractAvsGovernance.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.GrantRole(&_ContractAvsGovernance.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xfab57b8f.
//
// Solidity: function initialize((address,address,address,address,address,address,address,address,string,address) _initializationParams) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) Initialize(opts *bind.TransactOpts, _initializationParams IAvsGovernanceInitializationParams) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "initialize", _initializationParams)
}

// Initialize is a paid mutator transaction binding the contract method 0xfab57b8f.
//
// Solidity: function initialize((address,address,address,address,address,address,address,address,string,address) _initializationParams) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) Initialize(_initializationParams IAvsGovernanceInitializationParams) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.Initialize(&_ContractAvsGovernance.TransactOpts, _initializationParams)
}

// Initialize is a paid mutator transaction binding the contract method 0xfab57b8f.
//
// Solidity: function initialize((address,address,address,address,address,address,address,address,string,address) _initializationParams) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) Initialize(_initializationParams IAvsGovernanceInitializationParams) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.Initialize(&_ContractAvsGovernance.TransactOpts, _initializationParams)
}

// Pause is a paid mutator transaction binding the contract method 0x3aa83ec7.
//
// Solidity: function pause(bytes4 _pausableFlow) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) Pause(opts *bind.TransactOpts, _pausableFlow [4]byte) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "pause", _pausableFlow)
}

// Pause is a paid mutator transaction binding the contract method 0x3aa83ec7.
//
// Solidity: function pause(bytes4 _pausableFlow) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) Pause(_pausableFlow [4]byte) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.Pause(&_ContractAvsGovernance.TransactOpts, _pausableFlow)
}

// Pause is a paid mutator transaction binding the contract method 0x3aa83ec7.
//
// Solidity: function pause(bytes4 _pausableFlow) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) Pause(_pausableFlow [4]byte) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.Pause(&_ContractAvsGovernance.TransactOpts, _pausableFlow)
}

// QueueRewardsReceiverModification is a paid mutator transaction binding the contract method 0x1b21ba72.
//
// Solidity: function queueRewardsReceiverModification(address _newRewardsReceiver) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) QueueRewardsReceiverModification(opts *bind.TransactOpts, _newRewardsReceiver common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "queueRewardsReceiverModification", _newRewardsReceiver)
}

// QueueRewardsReceiverModification is a paid mutator transaction binding the contract method 0x1b21ba72.
//
// Solidity: function queueRewardsReceiverModification(address _newRewardsReceiver) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) QueueRewardsReceiverModification(_newRewardsReceiver common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.QueueRewardsReceiverModification(&_ContractAvsGovernance.TransactOpts, _newRewardsReceiver)
}

// QueueRewardsReceiverModification is a paid mutator transaction binding the contract method 0x1b21ba72.
//
// Solidity: function queueRewardsReceiverModification(address _newRewardsReceiver) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) QueueRewardsReceiverModification(_newRewardsReceiver common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.QueueRewardsReceiverModification(&_ContractAvsGovernance.TransactOpts, _newRewardsReceiver)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0xab8cac54.
//
// Solidity: function registerAsOperator((uint256[4],address,(uint256[2]),bytes) _operatorRegistrationParams) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) RegisterAsOperator(opts *bind.TransactOpts, _operatorRegistrationParams IAvsGovernanceOperatorRegistrationParams) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "registerAsOperator", _operatorRegistrationParams)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0xab8cac54.
//
// Solidity: function registerAsOperator((uint256[4],address,(uint256[2]),bytes) _operatorRegistrationParams) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) RegisterAsOperator(_operatorRegistrationParams IAvsGovernanceOperatorRegistrationParams) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.RegisterAsOperator(&_ContractAvsGovernance.TransactOpts, _operatorRegistrationParams)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0xab8cac54.
//
// Solidity: function registerAsOperator((uint256[4],address,(uint256[2]),bytes) _operatorRegistrationParams) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) RegisterAsOperator(_operatorRegistrationParams IAvsGovernanceOperatorRegistrationParams) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.RegisterAsOperator(&_ContractAvsGovernance.TransactOpts, _operatorRegistrationParams)
}

// RegisterAvsToEigenLayer is a paid mutator transaction binding the contract method 0x851bb725.
//
// Solidity: function registerAvsToEigenLayer(string metadataURI) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) RegisterAvsToEigenLayer(opts *bind.TransactOpts, metadataURI string) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "registerAvsToEigenLayer", metadataURI)
}

// RegisterAvsToEigenLayer is a paid mutator transaction binding the contract method 0x851bb725.
//
// Solidity: function registerAvsToEigenLayer(string metadataURI) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) RegisterAvsToEigenLayer(metadataURI string) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.RegisterAvsToEigenLayer(&_ContractAvsGovernance.TransactOpts, metadataURI)
}

// RegisterAvsToEigenLayer is a paid mutator transaction binding the contract method 0x851bb725.
//
// Solidity: function registerAvsToEigenLayer(string metadataURI) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) RegisterAvsToEigenLayer(metadataURI string) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.RegisterAvsToEigenLayer(&_ContractAvsGovernance.TransactOpts, metadataURI)
}

// RegisterAvsToSymbiotic is a paid mutator transaction binding the contract method 0xaa0dadf0.
//
// Solidity: function registerAvsToSymbiotic() returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) RegisterAvsToSymbiotic(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "registerAvsToSymbiotic")
}

// RegisterAvsToSymbiotic is a paid mutator transaction binding the contract method 0xaa0dadf0.
//
// Solidity: function registerAvsToSymbiotic() returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) RegisterAvsToSymbiotic() (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.RegisterAvsToSymbiotic(&_ContractAvsGovernance.TransactOpts)
}

// RegisterAvsToSymbiotic is a paid mutator transaction binding the contract method 0xaa0dadf0.
//
// Solidity: function registerAvsToSymbiotic() returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) RegisterAvsToSymbiotic() (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.RegisterAvsToSymbiotic(&_ContractAvsGovernance.TransactOpts)
}

// RegisterOperatorToEigenLayer is a paid mutator transaction binding the contract method 0x83e45170.
//
// Solidity: function registerOperatorToEigenLayer((bytes,bytes32,uint256) _eigenSig, bytes _authToken) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) RegisterOperatorToEigenLayer(opts *bind.TransactOpts, _eigenSig ISignatureUtilsSignatureWithSaltAndExpiry, _authToken []byte) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "registerOperatorToEigenLayer", _eigenSig, _authToken)
}

// RegisterOperatorToEigenLayer is a paid mutator transaction binding the contract method 0x83e45170.
//
// Solidity: function registerOperatorToEigenLayer((bytes,bytes32,uint256) _eigenSig, bytes _authToken) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) RegisterOperatorToEigenLayer(_eigenSig ISignatureUtilsSignatureWithSaltAndExpiry, _authToken []byte) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.RegisterOperatorToEigenLayer(&_ContractAvsGovernance.TransactOpts, _eigenSig, _authToken)
}

// RegisterOperatorToEigenLayer is a paid mutator transaction binding the contract method 0x83e45170.
//
// Solidity: function registerOperatorToEigenLayer((bytes,bytes32,uint256) _eigenSig, bytes _authToken) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) RegisterOperatorToEigenLayer(_eigenSig ISignatureUtilsSignatureWithSaltAndExpiry, _authToken []byte) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.RegisterOperatorToEigenLayer(&_ContractAvsGovernance.TransactOpts, _eigenSig, _authToken)
}

// RegisterOperatorToSymbiotic is a paid mutator transaction binding the contract method 0x93652277.
//
// Solidity: function registerOperatorToSymbiotic((uint48,bytes) _symbioticSig, bytes _authToken) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) RegisterOperatorToSymbiotic(opts *bind.TransactOpts, _symbioticSig IAvsGovernanceSymbioticOptInSignature, _authToken []byte) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "registerOperatorToSymbiotic", _symbioticSig, _authToken)
}

// RegisterOperatorToSymbiotic is a paid mutator transaction binding the contract method 0x93652277.
//
// Solidity: function registerOperatorToSymbiotic((uint48,bytes) _symbioticSig, bytes _authToken) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) RegisterOperatorToSymbiotic(_symbioticSig IAvsGovernanceSymbioticOptInSignature, _authToken []byte) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.RegisterOperatorToSymbiotic(&_ContractAvsGovernance.TransactOpts, _symbioticSig, _authToken)
}

// RegisterOperatorToSymbiotic is a paid mutator transaction binding the contract method 0x93652277.
//
// Solidity: function registerOperatorToSymbiotic((uint48,bytes) _symbioticSig, bytes _authToken) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) RegisterOperatorToSymbiotic(_symbioticSig IAvsGovernanceSymbioticOptInSignature, _authToken []byte) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.RegisterOperatorToSymbiotic(&_ContractAvsGovernance.TransactOpts, _symbioticSig, _authToken)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.RenounceRole(&_ContractAvsGovernance.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.RenounceRole(&_ContractAvsGovernance.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.RevokeRole(&_ContractAvsGovernance.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.RevokeRole(&_ContractAvsGovernance.TransactOpts, role, account)
}

// SetP2pAuthenticationEnabled is a paid mutator transaction binding the contract method 0xc7a93c6f.
//
// Solidity: function setP2pAuthenticationEnabled(bool _p2pAuthenticationEnabled) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) SetP2pAuthenticationEnabled(opts *bind.TransactOpts, _p2pAuthenticationEnabled bool) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "setP2pAuthenticationEnabled", _p2pAuthenticationEnabled)
}

// SetP2pAuthenticationEnabled is a paid mutator transaction binding the contract method 0xc7a93c6f.
//
// Solidity: function setP2pAuthenticationEnabled(bool _p2pAuthenticationEnabled) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) SetP2pAuthenticationEnabled(_p2pAuthenticationEnabled bool) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetP2pAuthenticationEnabled(&_ContractAvsGovernance.TransactOpts, _p2pAuthenticationEnabled)
}

// SetP2pAuthenticationEnabled is a paid mutator transaction binding the contract method 0xc7a93c6f.
//
// Solidity: function setP2pAuthenticationEnabled(bool _p2pAuthenticationEnabled) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) SetP2pAuthenticationEnabled(_p2pAuthenticationEnabled bool) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetP2pAuthenticationEnabled(&_ContractAvsGovernance.TransactOpts, _p2pAuthenticationEnabled)
}

// SetStakingContractMultiplier is a paid mutator transaction binding the contract method 0x82616a01.
//
// Solidity: function setStakingContractMultiplier((address,uint256,uint8) _votingPowerMultiplier) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) SetStakingContractMultiplier(opts *bind.TransactOpts, _votingPowerMultiplier IAvsGovernanceVotingPowerMultiplier) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "setStakingContractMultiplier", _votingPowerMultiplier)
}

// SetStakingContractMultiplier is a paid mutator transaction binding the contract method 0x82616a01.
//
// Solidity: function setStakingContractMultiplier((address,uint256,uint8) _votingPowerMultiplier) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) SetStakingContractMultiplier(_votingPowerMultiplier IAvsGovernanceVotingPowerMultiplier) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetStakingContractMultiplier(&_ContractAvsGovernance.TransactOpts, _votingPowerMultiplier)
}

// SetStakingContractMultiplier is a paid mutator transaction binding the contract method 0x82616a01.
//
// Solidity: function setStakingContractMultiplier((address,uint256,uint8) _votingPowerMultiplier) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) SetStakingContractMultiplier(_votingPowerMultiplier IAvsGovernanceVotingPowerMultiplier) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetStakingContractMultiplier(&_ContractAvsGovernance.TransactOpts, _votingPowerMultiplier)
}

// SetStakingContractMultiplierBatch is a paid mutator transaction binding the contract method 0x130bdcd9.
//
// Solidity: function setStakingContractMultiplierBatch((address,uint256,uint8)[] _votingPowerMultipliers) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) SetStakingContractMultiplierBatch(opts *bind.TransactOpts, _votingPowerMultipliers []IAvsGovernanceVotingPowerMultiplier) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "setStakingContractMultiplierBatch", _votingPowerMultipliers)
}

// SetStakingContractMultiplierBatch is a paid mutator transaction binding the contract method 0x130bdcd9.
//
// Solidity: function setStakingContractMultiplierBatch((address,uint256,uint8)[] _votingPowerMultipliers) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) SetStakingContractMultiplierBatch(_votingPowerMultipliers []IAvsGovernanceVotingPowerMultiplier) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetStakingContractMultiplierBatch(&_ContractAvsGovernance.TransactOpts, _votingPowerMultipliers)
}

// SetStakingContractMultiplierBatch is a paid mutator transaction binding the contract method 0x130bdcd9.
//
// Solidity: function setStakingContractMultiplierBatch((address,uint256,uint8)[] _votingPowerMultipliers) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) SetStakingContractMultiplierBatch(_votingPowerMultipliers []IAvsGovernanceVotingPowerMultiplier) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetStakingContractMultiplierBatch(&_ContractAvsGovernance.TransactOpts, _votingPowerMultipliers)
}

// SetSupportedStakingContracts is a paid mutator transaction binding the contract method 0x1dd2f74b.
//
// Solidity: function setSupportedStakingContracts((address,uint8)[] _stakingContractsDetails) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) SetSupportedStakingContracts(opts *bind.TransactOpts, _stakingContractsDetails []IAvsGovernanceStakingContractInfo) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "setSupportedStakingContracts", _stakingContractsDetails)
}

// SetSupportedStakingContracts is a paid mutator transaction binding the contract method 0x1dd2f74b.
//
// Solidity: function setSupportedStakingContracts((address,uint8)[] _stakingContractsDetails) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) SetSupportedStakingContracts(_stakingContractsDetails []IAvsGovernanceStakingContractInfo) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetSupportedStakingContracts(&_ContractAvsGovernance.TransactOpts, _stakingContractsDetails)
}

// SetSupportedStakingContracts is a paid mutator transaction binding the contract method 0x1dd2f74b.
//
// Solidity: function setSupportedStakingContracts((address,uint8)[] _stakingContractsDetails) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) SetSupportedStakingContracts(_stakingContractsDetails []IAvsGovernanceStakingContractInfo) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetSupportedStakingContracts(&_ContractAvsGovernance.TransactOpts, _stakingContractsDetails)
}

// Unpause is a paid mutator transaction binding the contract method 0xbac1e94b.
//
// Solidity: function unpause(bytes4 _pausableFlow) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) Unpause(opts *bind.TransactOpts, _pausableFlow [4]byte) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "unpause", _pausableFlow)
}

// Unpause is a paid mutator transaction binding the contract method 0xbac1e94b.
//
// Solidity: function unpause(bytes4 _pausableFlow) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) Unpause(_pausableFlow [4]byte) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.Unpause(&_ContractAvsGovernance.TransactOpts, _pausableFlow)
}

// Unpause is a paid mutator transaction binding the contract method 0xbac1e94b.
//
// Solidity: function unpause(bytes4 _pausableFlow) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) Unpause(_pausableFlow [4]byte) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.Unpause(&_ContractAvsGovernance.TransactOpts, _pausableFlow)
}

// UnregisterAsOperatorFromAvs is a paid mutator transaction binding the contract method 0x263fe7c7.
//
// Solidity: function unregisterAsOperatorFromAvs() returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) UnregisterAsOperatorFromAvs(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "unregisterAsOperatorFromAvs")
}

// UnregisterAsOperatorFromAvs is a paid mutator transaction binding the contract method 0x263fe7c7.
//
// Solidity: function unregisterAsOperatorFromAvs() returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) UnregisterAsOperatorFromAvs() (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.UnregisterAsOperatorFromAvs(&_ContractAvsGovernance.TransactOpts)
}

// UnregisterAsOperatorFromAvs is a paid mutator transaction binding the contract method 0x263fe7c7.
//
// Solidity: function unregisterAsOperatorFromAvs() returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) UnregisterAsOperatorFromAvs() (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.UnregisterAsOperatorFromAvs(&_ContractAvsGovernance.TransactOpts)
}

// UnregisterAsOperatorFromEigenLayer is a paid mutator transaction binding the contract method 0x0399f004.
//
// Solidity: function unregisterAsOperatorFromEigenLayer() returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) UnregisterAsOperatorFromEigenLayer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "unregisterAsOperatorFromEigenLayer")
}

// UnregisterAsOperatorFromEigenLayer is a paid mutator transaction binding the contract method 0x0399f004.
//
// Solidity: function unregisterAsOperatorFromEigenLayer() returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) UnregisterAsOperatorFromEigenLayer() (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.UnregisterAsOperatorFromEigenLayer(&_ContractAvsGovernance.TransactOpts)
}

// UnregisterAsOperatorFromEigenLayer is a paid mutator transaction binding the contract method 0x0399f004.
//
// Solidity: function unregisterAsOperatorFromEigenLayer() returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) UnregisterAsOperatorFromEigenLayer() (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.UnregisterAsOperatorFromEigenLayer(&_ContractAvsGovernance.TransactOpts)
}

// UnregisterAsOperatorFromSymbiotic is a paid mutator transaction binding the contract method 0x06255b01.
//
// Solidity: function unregisterAsOperatorFromSymbiotic((uint48,bytes) _unregistrationSignature) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) UnregisterAsOperatorFromSymbiotic(opts *bind.TransactOpts, _unregistrationSignature IAvsGovernanceSymbioticOptOutSignature) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "unregisterAsOperatorFromSymbiotic", _unregistrationSignature)
}

// UnregisterAsOperatorFromSymbiotic is a paid mutator transaction binding the contract method 0x06255b01.
//
// Solidity: function unregisterAsOperatorFromSymbiotic((uint48,bytes) _unregistrationSignature) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) UnregisterAsOperatorFromSymbiotic(_unregistrationSignature IAvsGovernanceSymbioticOptOutSignature) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.UnregisterAsOperatorFromSymbiotic(&_ContractAvsGovernance.TransactOpts, _unregistrationSignature)
}

// UnregisterAsOperatorFromSymbiotic is a paid mutator transaction binding the contract method 0x06255b01.
//
// Solidity: function unregisterAsOperatorFromSymbiotic((uint48,bytes) _unregistrationSignature) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) UnregisterAsOperatorFromSymbiotic(_unregistrationSignature IAvsGovernanceSymbioticOptOutSignature) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.UnregisterAsOperatorFromSymbiotic(&_ContractAvsGovernance.TransactOpts, _unregistrationSignature)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.Fallback(&_ContractAvsGovernance.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.Fallback(&_ContractAvsGovernance.TransactOpts, calldata)
}

// ContractAvsGovernanceDepositRewardsBackFailedIterator is returned from FilterDepositRewardsBackFailed and is used to iterate over the raw logs and unpacked data for DepositRewardsBackFailed events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceDepositRewardsBackFailedIterator struct {
	Event *ContractAvsGovernanceDepositRewardsBackFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAvsGovernanceDepositRewardsBackFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceDepositRewardsBackFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAvsGovernanceDepositRewardsBackFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAvsGovernanceDepositRewardsBackFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceDepositRewardsBackFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceDepositRewardsBackFailed represents a DepositRewardsBackFailed event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceDepositRewardsBackFailed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDepositRewardsBackFailed is a free log retrieval operation binding the contract event 0xed3143332c2d71fbe302b126ea59ee4fb74a1699cb820f835a913b15c027798d.
//
// Solidity: event DepositRewardsBackFailed()
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterDepositRewardsBackFailed(opts *bind.FilterOpts) (*ContractAvsGovernanceDepositRewardsBackFailedIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "DepositRewardsBackFailed")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceDepositRewardsBackFailedIterator{contract: _ContractAvsGovernance.contract, event: "DepositRewardsBackFailed", logs: logs, sub: sub}, nil
}

// WatchDepositRewardsBackFailed is a free log subscription operation binding the contract event 0xed3143332c2d71fbe302b126ea59ee4fb74a1699cb820f835a913b15c027798d.
//
// Solidity: event DepositRewardsBackFailed()
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchDepositRewardsBackFailed(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceDepositRewardsBackFailed) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "DepositRewardsBackFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceDepositRewardsBackFailed)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "DepositRewardsBackFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDepositRewardsBackFailed is a log parse operation binding the contract event 0xed3143332c2d71fbe302b126ea59ee4fb74a1699cb820f835a913b15c027798d.
//
// Solidity: event DepositRewardsBackFailed()
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseDepositRewardsBackFailed(log types.Log) (*ContractAvsGovernanceDepositRewardsBackFailed, error) {
	event := new(ContractAvsGovernanceDepositRewardsBackFailed)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "DepositRewardsBackFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceFlowPausedIterator is returned from FilterFlowPaused and is used to iterate over the raw logs and unpacked data for FlowPaused events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceFlowPausedIterator struct {
	Event *ContractAvsGovernanceFlowPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAvsGovernanceFlowPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceFlowPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAvsGovernanceFlowPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAvsGovernanceFlowPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceFlowPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceFlowPaused represents a FlowPaused event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceFlowPaused struct {
	PausableFlow [4]byte
	Pauser       common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFlowPaused is a free log retrieval operation binding the contract event 0x95c3658c5e0c74e20cf12db371b9b67d26e97a1937f6d2284f88cc44d036b4f6.
//
// Solidity: event FlowPaused(bytes4 _pausableFlow, address _pauser)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterFlowPaused(opts *bind.FilterOpts) (*ContractAvsGovernanceFlowPausedIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "FlowPaused")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceFlowPausedIterator{contract: _ContractAvsGovernance.contract, event: "FlowPaused", logs: logs, sub: sub}, nil
}

// WatchFlowPaused is a free log subscription operation binding the contract event 0x95c3658c5e0c74e20cf12db371b9b67d26e97a1937f6d2284f88cc44d036b4f6.
//
// Solidity: event FlowPaused(bytes4 _pausableFlow, address _pauser)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchFlowPaused(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceFlowPaused) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "FlowPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceFlowPaused)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "FlowPaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFlowPaused is a log parse operation binding the contract event 0x95c3658c5e0c74e20cf12db371b9b67d26e97a1937f6d2284f88cc44d036b4f6.
//
// Solidity: event FlowPaused(bytes4 _pausableFlow, address _pauser)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseFlowPaused(log types.Log) (*ContractAvsGovernanceFlowPaused, error) {
	event := new(ContractAvsGovernanceFlowPaused)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "FlowPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceFlowUnpausedIterator is returned from FilterFlowUnpaused and is used to iterate over the raw logs and unpacked data for FlowUnpaused events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceFlowUnpausedIterator struct {
	Event *ContractAvsGovernanceFlowUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAvsGovernanceFlowUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceFlowUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAvsGovernanceFlowUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAvsGovernanceFlowUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceFlowUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceFlowUnpaused represents a FlowUnpaused event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceFlowUnpaused struct {
	PausableFlowFlag [4]byte
	Unpauser         common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterFlowUnpaused is a free log retrieval operation binding the contract event 0xc7e56e17b0a6c4b467df6495e1eda1baecd7ba20604e80c1058ac06f4578d85e.
//
// Solidity: event FlowUnpaused(bytes4 _pausableFlowFlag, address _unpauser)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterFlowUnpaused(opts *bind.FilterOpts) (*ContractAvsGovernanceFlowUnpausedIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "FlowUnpaused")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceFlowUnpausedIterator{contract: _ContractAvsGovernance.contract, event: "FlowUnpaused", logs: logs, sub: sub}, nil
}

// WatchFlowUnpaused is a free log subscription operation binding the contract event 0xc7e56e17b0a6c4b467df6495e1eda1baecd7ba20604e80c1058ac06f4578d85e.
//
// Solidity: event FlowUnpaused(bytes4 _pausableFlowFlag, address _unpauser)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchFlowUnpaused(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceFlowUnpaused) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "FlowUnpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceFlowUnpaused)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "FlowUnpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFlowUnpaused is a log parse operation binding the contract event 0xc7e56e17b0a6c4b467df6495e1eda1baecd7ba20604e80c1058ac06f4578d85e.
//
// Solidity: event FlowUnpaused(bytes4 _pausableFlowFlag, address _unpauser)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseFlowUnpaused(log types.Log) (*ContractAvsGovernanceFlowUnpaused, error) {
	event := new(ContractAvsGovernanceFlowUnpaused)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "FlowUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceInitializedIterator struct {
	Event *ContractAvsGovernanceInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAvsGovernanceInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAvsGovernanceInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAvsGovernanceInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceInitialized represents a Initialized event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractAvsGovernanceInitializedIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceInitializedIterator{contract: _ContractAvsGovernance.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceInitialized)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseInitialized(log types.Log) (*ContractAvsGovernanceInitialized, error) {
	event := new(ContractAvsGovernanceInitialized)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceNativeCoinNotSupportedForEigenRewardsIterator is returned from FilterNativeCoinNotSupportedForEigenRewards and is used to iterate over the raw logs and unpacked data for NativeCoinNotSupportedForEigenRewards events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceNativeCoinNotSupportedForEigenRewardsIterator struct {
	Event *ContractAvsGovernanceNativeCoinNotSupportedForEigenRewards // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAvsGovernanceNativeCoinNotSupportedForEigenRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceNativeCoinNotSupportedForEigenRewards)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAvsGovernanceNativeCoinNotSupportedForEigenRewards)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAvsGovernanceNativeCoinNotSupportedForEigenRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceNativeCoinNotSupportedForEigenRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceNativeCoinNotSupportedForEigenRewards represents a NativeCoinNotSupportedForEigenRewards event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceNativeCoinNotSupportedForEigenRewards struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNativeCoinNotSupportedForEigenRewards is a free log retrieval operation binding the contract event 0xb3f3d8f7b39bd24844dec3cce1a06cfee42f663826e9b790c6a36fb8c05dbf8e.
//
// Solidity: event NativeCoinNotSupportedForEigenRewards()
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterNativeCoinNotSupportedForEigenRewards(opts *bind.FilterOpts) (*ContractAvsGovernanceNativeCoinNotSupportedForEigenRewardsIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "NativeCoinNotSupportedForEigenRewards")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceNativeCoinNotSupportedForEigenRewardsIterator{contract: _ContractAvsGovernance.contract, event: "NativeCoinNotSupportedForEigenRewards", logs: logs, sub: sub}, nil
}

// WatchNativeCoinNotSupportedForEigenRewards is a free log subscription operation binding the contract event 0xb3f3d8f7b39bd24844dec3cce1a06cfee42f663826e9b790c6a36fb8c05dbf8e.
//
// Solidity: event NativeCoinNotSupportedForEigenRewards()
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchNativeCoinNotSupportedForEigenRewards(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceNativeCoinNotSupportedForEigenRewards) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "NativeCoinNotSupportedForEigenRewards")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceNativeCoinNotSupportedForEigenRewards)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "NativeCoinNotSupportedForEigenRewards", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNativeCoinNotSupportedForEigenRewards is a log parse operation binding the contract event 0xb3f3d8f7b39bd24844dec3cce1a06cfee42f663826e9b790c6a36fb8c05dbf8e.
//
// Solidity: event NativeCoinNotSupportedForEigenRewards()
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseNativeCoinNotSupportedForEigenRewards(log types.Log) (*ContractAvsGovernanceNativeCoinNotSupportedForEigenRewards, error) {
	event := new(ContractAvsGovernanceNativeCoinNotSupportedForEigenRewards)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "NativeCoinNotSupportedForEigenRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceOperatorEjectedFromNetworkIterator is returned from FilterOperatorEjectedFromNetwork and is used to iterate over the raw logs and unpacked data for OperatorEjectedFromNetwork events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceOperatorEjectedFromNetworkIterator struct {
	Event *ContractAvsGovernanceOperatorEjectedFromNetwork // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAvsGovernanceOperatorEjectedFromNetworkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceOperatorEjectedFromNetwork)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAvsGovernanceOperatorEjectedFromNetwork)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAvsGovernanceOperatorEjectedFromNetworkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceOperatorEjectedFromNetworkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceOperatorEjectedFromNetwork represents a OperatorEjectedFromNetwork event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceOperatorEjectedFromNetwork struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorEjectedFromNetwork is a free log retrieval operation binding the contract event 0x98ac533c958bf5d682a0590dc21d14cc6a8acaaa7886880124cfb4531266f2f2.
//
// Solidity: event OperatorEjectedFromNetwork(address operator)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterOperatorEjectedFromNetwork(opts *bind.FilterOpts) (*ContractAvsGovernanceOperatorEjectedFromNetworkIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "OperatorEjectedFromNetwork")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceOperatorEjectedFromNetworkIterator{contract: _ContractAvsGovernance.contract, event: "OperatorEjectedFromNetwork", logs: logs, sub: sub}, nil
}

// WatchOperatorEjectedFromNetwork is a free log subscription operation binding the contract event 0x98ac533c958bf5d682a0590dc21d14cc6a8acaaa7886880124cfb4531266f2f2.
//
// Solidity: event OperatorEjectedFromNetwork(address operator)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchOperatorEjectedFromNetwork(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceOperatorEjectedFromNetwork) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "OperatorEjectedFromNetwork")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceOperatorEjectedFromNetwork)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "OperatorEjectedFromNetwork", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorEjectedFromNetwork is a log parse operation binding the contract event 0x98ac533c958bf5d682a0590dc21d14cc6a8acaaa7886880124cfb4531266f2f2.
//
// Solidity: event OperatorEjectedFromNetwork(address operator)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseOperatorEjectedFromNetwork(log types.Log) (*ContractAvsGovernanceOperatorEjectedFromNetwork, error) {
	event := new(ContractAvsGovernanceOperatorEjectedFromNetwork)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "OperatorEjectedFromNetwork", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceOperatorRegisteredIterator is returned from FilterOperatorRegistered and is used to iterate over the raw logs and unpacked data for OperatorRegistered events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceOperatorRegisteredIterator struct {
	Event *ContractAvsGovernanceOperatorRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAvsGovernanceOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceOperatorRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAvsGovernanceOperatorRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAvsGovernanceOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceOperatorRegistered represents a OperatorRegistered event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceOperatorRegistered struct {
	Operator common.Address
	BlsKey   [4]*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0x54bc9cf83c2eb0f2ad1abf6e4fab882964404622ba2df6b5a9356a18d3aac055.
//
// Solidity: event OperatorRegistered(address indexed operator, uint256[4] blsKey)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address) (*ContractAvsGovernanceOperatorRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceOperatorRegisteredIterator{contract: _ContractAvsGovernance.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0x54bc9cf83c2eb0f2ad1abf6e4fab882964404622ba2df6b5a9356a18d3aac055.
//
// Solidity: event OperatorRegistered(address indexed operator, uint256[4] blsKey)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceOperatorRegistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceOperatorRegistered)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorRegistered is a log parse operation binding the contract event 0x54bc9cf83c2eb0f2ad1abf6e4fab882964404622ba2df6b5a9356a18d3aac055.
//
// Solidity: event OperatorRegistered(address indexed operator, uint256[4] blsKey)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseOperatorRegistered(log types.Log) (*ContractAvsGovernanceOperatorRegistered, error) {
	event := new(ContractAvsGovernanceOperatorRegistered)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceOperatorRegisteredToEigenLayerIterator is returned from FilterOperatorRegisteredToEigenLayer and is used to iterate over the raw logs and unpacked data for OperatorRegisteredToEigenLayer events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceOperatorRegisteredToEigenLayerIterator struct {
	Event *ContractAvsGovernanceOperatorRegisteredToEigenLayer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAvsGovernanceOperatorRegisteredToEigenLayerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceOperatorRegisteredToEigenLayer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAvsGovernanceOperatorRegisteredToEigenLayer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAvsGovernanceOperatorRegisteredToEigenLayerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceOperatorRegisteredToEigenLayerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceOperatorRegisteredToEigenLayer represents a OperatorRegisteredToEigenLayer event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceOperatorRegisteredToEigenLayer struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegisteredToEigenLayer is a free log retrieval operation binding the contract event 0x23c66044e8b2ab58f45f4aa7eeb03c47c231fec8b7cf140f5b709455e5c1aa1d.
//
// Solidity: event OperatorRegisteredToEigenLayer(address operator)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterOperatorRegisteredToEigenLayer(opts *bind.FilterOpts) (*ContractAvsGovernanceOperatorRegisteredToEigenLayerIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "OperatorRegisteredToEigenLayer")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceOperatorRegisteredToEigenLayerIterator{contract: _ContractAvsGovernance.contract, event: "OperatorRegisteredToEigenLayer", logs: logs, sub: sub}, nil
}

// WatchOperatorRegisteredToEigenLayer is a free log subscription operation binding the contract event 0x23c66044e8b2ab58f45f4aa7eeb03c47c231fec8b7cf140f5b709455e5c1aa1d.
//
// Solidity: event OperatorRegisteredToEigenLayer(address operator)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchOperatorRegisteredToEigenLayer(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceOperatorRegisteredToEigenLayer) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "OperatorRegisteredToEigenLayer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceOperatorRegisteredToEigenLayer)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "OperatorRegisteredToEigenLayer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorRegisteredToEigenLayer is a log parse operation binding the contract event 0x23c66044e8b2ab58f45f4aa7eeb03c47c231fec8b7cf140f5b709455e5c1aa1d.
//
// Solidity: event OperatorRegisteredToEigenLayer(address operator)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseOperatorRegisteredToEigenLayer(log types.Log) (*ContractAvsGovernanceOperatorRegisteredToEigenLayer, error) {
	event := new(ContractAvsGovernanceOperatorRegisteredToEigenLayer)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "OperatorRegisteredToEigenLayer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceOperatorRegisteredToSymbioticIterator is returned from FilterOperatorRegisteredToSymbiotic and is used to iterate over the raw logs and unpacked data for OperatorRegisteredToSymbiotic events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceOperatorRegisteredToSymbioticIterator struct {
	Event *ContractAvsGovernanceOperatorRegisteredToSymbiotic // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAvsGovernanceOperatorRegisteredToSymbioticIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceOperatorRegisteredToSymbiotic)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAvsGovernanceOperatorRegisteredToSymbiotic)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAvsGovernanceOperatorRegisteredToSymbioticIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceOperatorRegisteredToSymbioticIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceOperatorRegisteredToSymbiotic represents a OperatorRegisteredToSymbiotic event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceOperatorRegisteredToSymbiotic struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegisteredToSymbiotic is a free log retrieval operation binding the contract event 0x6f6f010344f25a8f5774e47688e7db5f23ef1071452f6048e1147d8de8cfe2d1.
//
// Solidity: event OperatorRegisteredToSymbiotic(address operator)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterOperatorRegisteredToSymbiotic(opts *bind.FilterOpts) (*ContractAvsGovernanceOperatorRegisteredToSymbioticIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "OperatorRegisteredToSymbiotic")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceOperatorRegisteredToSymbioticIterator{contract: _ContractAvsGovernance.contract, event: "OperatorRegisteredToSymbiotic", logs: logs, sub: sub}, nil
}

// WatchOperatorRegisteredToSymbiotic is a free log subscription operation binding the contract event 0x6f6f010344f25a8f5774e47688e7db5f23ef1071452f6048e1147d8de8cfe2d1.
//
// Solidity: event OperatorRegisteredToSymbiotic(address operator)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchOperatorRegisteredToSymbiotic(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceOperatorRegisteredToSymbiotic) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "OperatorRegisteredToSymbiotic")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceOperatorRegisteredToSymbiotic)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "OperatorRegisteredToSymbiotic", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorRegisteredToSymbiotic is a log parse operation binding the contract event 0x6f6f010344f25a8f5774e47688e7db5f23ef1071452f6048e1147d8de8cfe2d1.
//
// Solidity: event OperatorRegisteredToSymbiotic(address operator)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseOperatorRegisteredToSymbiotic(log types.Log) (*ContractAvsGovernanceOperatorRegisteredToSymbiotic, error) {
	event := new(ContractAvsGovernanceOperatorRegisteredToSymbiotic)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "OperatorRegisteredToSymbiotic", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceOperatorUnregisteredIterator is returned from FilterOperatorUnregistered and is used to iterate over the raw logs and unpacked data for OperatorUnregistered events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceOperatorUnregisteredIterator struct {
	Event *ContractAvsGovernanceOperatorUnregistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAvsGovernanceOperatorUnregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceOperatorUnregistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAvsGovernanceOperatorUnregistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAvsGovernanceOperatorUnregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceOperatorUnregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceOperatorUnregistered represents a OperatorUnregistered event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceOperatorUnregistered struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorUnregistered is a free log retrieval operation binding the contract event 0x6f42117a557500c705ddf040a619d86f39101e6b74ac20d7b3e5943ba473fc7f.
//
// Solidity: event OperatorUnregistered(address indexed operator)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterOperatorUnregistered(opts *bind.FilterOpts, operator []common.Address) (*ContractAvsGovernanceOperatorUnregisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "OperatorUnregistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceOperatorUnregisteredIterator{contract: _ContractAvsGovernance.contract, event: "OperatorUnregistered", logs: logs, sub: sub}, nil
}

// WatchOperatorUnregistered is a free log subscription operation binding the contract event 0x6f42117a557500c705ddf040a619d86f39101e6b74ac20d7b3e5943ba473fc7f.
//
// Solidity: event OperatorUnregistered(address indexed operator)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchOperatorUnregistered(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceOperatorUnregistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "OperatorUnregistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceOperatorUnregistered)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "OperatorUnregistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorUnregistered is a log parse operation binding the contract event 0x6f42117a557500c705ddf040a619d86f39101e6b74ac20d7b3e5943ba473fc7f.
//
// Solidity: event OperatorUnregistered(address indexed operator)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseOperatorUnregistered(log types.Log) (*ContractAvsGovernanceOperatorUnregistered, error) {
	event := new(ContractAvsGovernanceOperatorUnregistered)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "OperatorUnregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceOperatorUnregisteredToEigenLayerIterator is returned from FilterOperatorUnregisteredToEigenLayer and is used to iterate over the raw logs and unpacked data for OperatorUnregisteredToEigenLayer events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceOperatorUnregisteredToEigenLayerIterator struct {
	Event *ContractAvsGovernanceOperatorUnregisteredToEigenLayer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAvsGovernanceOperatorUnregisteredToEigenLayerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceOperatorUnregisteredToEigenLayer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAvsGovernanceOperatorUnregisteredToEigenLayer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAvsGovernanceOperatorUnregisteredToEigenLayerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceOperatorUnregisteredToEigenLayerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceOperatorUnregisteredToEigenLayer represents a OperatorUnregisteredToEigenLayer event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceOperatorUnregisteredToEigenLayer struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorUnregisteredToEigenLayer is a free log retrieval operation binding the contract event 0x7877897edb095f8cd999adcb582e04128ffb91d12bf5139ca7b8a1567a845300.
//
// Solidity: event OperatorUnregisteredToEigenLayer(address operator)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterOperatorUnregisteredToEigenLayer(opts *bind.FilterOpts) (*ContractAvsGovernanceOperatorUnregisteredToEigenLayerIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "OperatorUnregisteredToEigenLayer")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceOperatorUnregisteredToEigenLayerIterator{contract: _ContractAvsGovernance.contract, event: "OperatorUnregisteredToEigenLayer", logs: logs, sub: sub}, nil
}

// WatchOperatorUnregisteredToEigenLayer is a free log subscription operation binding the contract event 0x7877897edb095f8cd999adcb582e04128ffb91d12bf5139ca7b8a1567a845300.
//
// Solidity: event OperatorUnregisteredToEigenLayer(address operator)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchOperatorUnregisteredToEigenLayer(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceOperatorUnregisteredToEigenLayer) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "OperatorUnregisteredToEigenLayer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceOperatorUnregisteredToEigenLayer)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "OperatorUnregisteredToEigenLayer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorUnregisteredToEigenLayer is a log parse operation binding the contract event 0x7877897edb095f8cd999adcb582e04128ffb91d12bf5139ca7b8a1567a845300.
//
// Solidity: event OperatorUnregisteredToEigenLayer(address operator)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseOperatorUnregisteredToEigenLayer(log types.Log) (*ContractAvsGovernanceOperatorUnregisteredToEigenLayer, error) {
	event := new(ContractAvsGovernanceOperatorUnregisteredToEigenLayer)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "OperatorUnregisteredToEigenLayer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceOperatorUnregisteredToSymbioticIterator is returned from FilterOperatorUnregisteredToSymbiotic and is used to iterate over the raw logs and unpacked data for OperatorUnregisteredToSymbiotic events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceOperatorUnregisteredToSymbioticIterator struct {
	Event *ContractAvsGovernanceOperatorUnregisteredToSymbiotic // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAvsGovernanceOperatorUnregisteredToSymbioticIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceOperatorUnregisteredToSymbiotic)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAvsGovernanceOperatorUnregisteredToSymbiotic)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAvsGovernanceOperatorUnregisteredToSymbioticIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceOperatorUnregisteredToSymbioticIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceOperatorUnregisteredToSymbiotic represents a OperatorUnregisteredToSymbiotic event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceOperatorUnregisteredToSymbiotic struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorUnregisteredToSymbiotic is a free log retrieval operation binding the contract event 0xfc7eb04654394ed6ffe862307d7824361d3481d69ea2775b77e85837befe7737.
//
// Solidity: event OperatorUnregisteredToSymbiotic(address operator)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterOperatorUnregisteredToSymbiotic(opts *bind.FilterOpts) (*ContractAvsGovernanceOperatorUnregisteredToSymbioticIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "OperatorUnregisteredToSymbiotic")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceOperatorUnregisteredToSymbioticIterator{contract: _ContractAvsGovernance.contract, event: "OperatorUnregisteredToSymbiotic", logs: logs, sub: sub}, nil
}

// WatchOperatorUnregisteredToSymbiotic is a free log subscription operation binding the contract event 0xfc7eb04654394ed6ffe862307d7824361d3481d69ea2775b77e85837befe7737.
//
// Solidity: event OperatorUnregisteredToSymbiotic(address operator)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchOperatorUnregisteredToSymbiotic(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceOperatorUnregisteredToSymbiotic) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "OperatorUnregisteredToSymbiotic")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceOperatorUnregisteredToSymbiotic)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "OperatorUnregisteredToSymbiotic", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorUnregisteredToSymbiotic is a log parse operation binding the contract event 0xfc7eb04654394ed6ffe862307d7824361d3481d69ea2775b77e85837befe7737.
//
// Solidity: event OperatorUnregisteredToSymbiotic(address operator)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseOperatorUnregisteredToSymbiotic(log types.Log) (*ContractAvsGovernanceOperatorUnregisteredToSymbiotic, error) {
	event := new(ContractAvsGovernanceOperatorUnregisteredToSymbiotic)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "OperatorUnregisteredToSymbiotic", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceQueuedRewardsReceiverModificationIterator is returned from FilterQueuedRewardsReceiverModification and is used to iterate over the raw logs and unpacked data for QueuedRewardsReceiverModification events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceQueuedRewardsReceiverModificationIterator struct {
	Event *ContractAvsGovernanceQueuedRewardsReceiverModification // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAvsGovernanceQueuedRewardsReceiverModificationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceQueuedRewardsReceiverModification)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAvsGovernanceQueuedRewardsReceiverModification)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAvsGovernanceQueuedRewardsReceiverModificationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceQueuedRewardsReceiverModificationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceQueuedRewardsReceiverModification represents a QueuedRewardsReceiverModification event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceQueuedRewardsReceiverModification struct {
	Operator common.Address
	Receiver common.Address
	Delay    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterQueuedRewardsReceiverModification is a free log retrieval operation binding the contract event 0x0d8cfa10a3087b28d3c226ad9a37314860e7c3c0505a25a39e3cdefb3118a98a.
//
// Solidity: event QueuedRewardsReceiverModification(address indexed operator, address receiver, uint256 delay)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterQueuedRewardsReceiverModification(opts *bind.FilterOpts, operator []common.Address) (*ContractAvsGovernanceQueuedRewardsReceiverModificationIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "QueuedRewardsReceiverModification", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceQueuedRewardsReceiverModificationIterator{contract: _ContractAvsGovernance.contract, event: "QueuedRewardsReceiverModification", logs: logs, sub: sub}, nil
}

// WatchQueuedRewardsReceiverModification is a free log subscription operation binding the contract event 0x0d8cfa10a3087b28d3c226ad9a37314860e7c3c0505a25a39e3cdefb3118a98a.
//
// Solidity: event QueuedRewardsReceiverModification(address indexed operator, address receiver, uint256 delay)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchQueuedRewardsReceiverModification(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceQueuedRewardsReceiverModification, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "QueuedRewardsReceiverModification", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceQueuedRewardsReceiverModification)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "QueuedRewardsReceiverModification", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseQueuedRewardsReceiverModification is a log parse operation binding the contract event 0x0d8cfa10a3087b28d3c226ad9a37314860e7c3c0505a25a39e3cdefb3118a98a.
//
// Solidity: event QueuedRewardsReceiverModification(address indexed operator, address receiver, uint256 delay)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseQueuedRewardsReceiverModification(log types.Log) (*ContractAvsGovernanceQueuedRewardsReceiverModification, error) {
	event := new(ContractAvsGovernanceQueuedRewardsReceiverModification)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "QueuedRewardsReceiverModification", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceRewardsCoordinatorRevertedIterator is returned from FilterRewardsCoordinatorReverted and is used to iterate over the raw logs and unpacked data for RewardsCoordinatorReverted events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceRewardsCoordinatorRevertedIterator struct {
	Event *ContractAvsGovernanceRewardsCoordinatorReverted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAvsGovernanceRewardsCoordinatorRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceRewardsCoordinatorReverted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAvsGovernanceRewardsCoordinatorReverted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAvsGovernanceRewardsCoordinatorRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceRewardsCoordinatorRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceRewardsCoordinatorReverted represents a RewardsCoordinatorReverted event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceRewardsCoordinatorReverted struct {
	RevertData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRewardsCoordinatorReverted is a free log retrieval operation binding the contract event 0xab7e71f561320fa16a226c9744bf4d230ce7faf885df7bbc890baf339faa0a4e.
//
// Solidity: event RewardsCoordinatorReverted(bytes revertData)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterRewardsCoordinatorReverted(opts *bind.FilterOpts) (*ContractAvsGovernanceRewardsCoordinatorRevertedIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "RewardsCoordinatorReverted")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceRewardsCoordinatorRevertedIterator{contract: _ContractAvsGovernance.contract, event: "RewardsCoordinatorReverted", logs: logs, sub: sub}, nil
}

// WatchRewardsCoordinatorReverted is a free log subscription operation binding the contract event 0xab7e71f561320fa16a226c9744bf4d230ce7faf885df7bbc890baf339faa0a4e.
//
// Solidity: event RewardsCoordinatorReverted(bytes revertData)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchRewardsCoordinatorReverted(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceRewardsCoordinatorReverted) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "RewardsCoordinatorReverted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceRewardsCoordinatorReverted)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "RewardsCoordinatorReverted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardsCoordinatorReverted is a log parse operation binding the contract event 0xab7e71f561320fa16a226c9744bf4d230ce7faf885df7bbc890baf339faa0a4e.
//
// Solidity: event RewardsCoordinatorReverted(bytes revertData)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseRewardsCoordinatorReverted(log types.Log) (*ContractAvsGovernanceRewardsCoordinatorReverted, error) {
	event := new(ContractAvsGovernanceRewardsCoordinatorReverted)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "RewardsCoordinatorReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceRoleAdminChangedIterator struct {
	Event *ContractAvsGovernanceRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAvsGovernanceRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAvsGovernanceRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAvsGovernanceRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceRoleAdminChanged represents a RoleAdminChanged event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ContractAvsGovernanceRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceRoleAdminChangedIterator{contract: _ContractAvsGovernance.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceRoleAdminChanged)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseRoleAdminChanged(log types.Log) (*ContractAvsGovernanceRoleAdminChanged, error) {
	event := new(ContractAvsGovernanceRoleAdminChanged)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceRoleGrantedIterator struct {
	Event *ContractAvsGovernanceRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAvsGovernanceRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAvsGovernanceRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAvsGovernanceRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceRoleGranted represents a RoleGranted event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ContractAvsGovernanceRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceRoleGrantedIterator{contract: _ContractAvsGovernance.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceRoleGranted)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseRoleGranted(log types.Log) (*ContractAvsGovernanceRoleGranted, error) {
	event := new(ContractAvsGovernanceRoleGranted)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceRoleRevokedIterator struct {
	Event *ContractAvsGovernanceRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAvsGovernanceRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAvsGovernanceRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAvsGovernanceRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceRoleRevoked represents a RoleRevoked event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ContractAvsGovernanceRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceRoleRevokedIterator{contract: _ContractAvsGovernance.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceRoleRevoked)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseRoleRevoked(log types.Log) (*ContractAvsGovernanceRoleRevoked, error) {
	event := new(ContractAvsGovernanceRoleRevoked)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceSetAvsNameIterator is returned from FilterSetAvsName and is used to iterate over the raw logs and unpacked data for SetAvsName events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetAvsNameIterator struct {
	Event *ContractAvsGovernanceSetAvsName // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAvsGovernanceSetAvsNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceSetAvsName)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAvsGovernanceSetAvsName)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAvsGovernanceSetAvsNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceSetAvsNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceSetAvsName represents a SetAvsName event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetAvsName struct {
	AvsName string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetAvsName is a free log retrieval operation binding the contract event 0x7f63aacad63bc1693280450d5c3612ccd4efc53e46d69f3a537db102cd66290c.
//
// Solidity: event SetAvsName(string avsName)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterSetAvsName(opts *bind.FilterOpts) (*ContractAvsGovernanceSetAvsNameIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "SetAvsName")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceSetAvsNameIterator{contract: _ContractAvsGovernance.contract, event: "SetAvsName", logs: logs, sub: sub}, nil
}

// WatchSetAvsName is a free log subscription operation binding the contract event 0x7f63aacad63bc1693280450d5c3612ccd4efc53e46d69f3a537db102cd66290c.
//
// Solidity: event SetAvsName(string avsName)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchSetAvsName(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetAvsName) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "SetAvsName")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceSetAvsName)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetAvsName", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetAvsName is a log parse operation binding the contract event 0x7f63aacad63bc1693280450d5c3612ccd4efc53e46d69f3a537db102cd66290c.
//
// Solidity: event SetAvsName(string avsName)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseSetAvsName(log types.Log) (*ContractAvsGovernanceSetAvsName, error) {
	event := new(ContractAvsGovernanceSetAvsName)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetAvsName", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceSetIsAllowlistedIterator is returned from FilterSetIsAllowlisted and is used to iterate over the raw logs and unpacked data for SetIsAllowlisted events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetIsAllowlistedIterator struct {
	Event *ContractAvsGovernanceSetIsAllowlisted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAvsGovernanceSetIsAllowlistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceSetIsAllowlisted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAvsGovernanceSetIsAllowlisted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAvsGovernanceSetIsAllowlistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceSetIsAllowlistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceSetIsAllowlisted represents a SetIsAllowlisted event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetIsAllowlisted struct {
	IsAllowlisted bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetIsAllowlisted is a free log retrieval operation binding the contract event 0x2dcb3282f9b7aa18e1bf7fa254c45f3e270e8f26d9a37ae590d5d8125b58d1b1.
//
// Solidity: event SetIsAllowlisted(bool isAllowlisted)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterSetIsAllowlisted(opts *bind.FilterOpts) (*ContractAvsGovernanceSetIsAllowlistedIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "SetIsAllowlisted")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceSetIsAllowlistedIterator{contract: _ContractAvsGovernance.contract, event: "SetIsAllowlisted", logs: logs, sub: sub}, nil
}

// WatchSetIsAllowlisted is a free log subscription operation binding the contract event 0x2dcb3282f9b7aa18e1bf7fa254c45f3e270e8f26d9a37ae590d5d8125b58d1b1.
//
// Solidity: event SetIsAllowlisted(bool isAllowlisted)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchSetIsAllowlisted(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetIsAllowlisted) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "SetIsAllowlisted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceSetIsAllowlisted)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetIsAllowlisted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetIsAllowlisted is a log parse operation binding the contract event 0x2dcb3282f9b7aa18e1bf7fa254c45f3e270e8f26d9a37ae590d5d8125b58d1b1.
//
// Solidity: event SetIsAllowlisted(bool isAllowlisted)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseSetIsAllowlisted(log types.Log) (*ContractAvsGovernanceSetIsAllowlisted, error) {
	event := new(ContractAvsGovernanceSetIsAllowlisted)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetIsAllowlisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceSetP2pAuthenticationEnabledIterator is returned from FilterSetP2pAuthenticationEnabled and is used to iterate over the raw logs and unpacked data for SetP2pAuthenticationEnabled events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetP2pAuthenticationEnabledIterator struct {
	Event *ContractAvsGovernanceSetP2pAuthenticationEnabled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAvsGovernanceSetP2pAuthenticationEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceSetP2pAuthenticationEnabled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAvsGovernanceSetP2pAuthenticationEnabled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAvsGovernanceSetP2pAuthenticationEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceSetP2pAuthenticationEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceSetP2pAuthenticationEnabled represents a SetP2pAuthenticationEnabled event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetP2pAuthenticationEnabled struct {
	IsEnabled bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetP2pAuthenticationEnabled is a free log retrieval operation binding the contract event 0x476bc7159a4e42477e0aa2e2e9709c55cb469b547976f91002b8db00bbd9b2f6.
//
// Solidity: event SetP2pAuthenticationEnabled(bool _isEnabled)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterSetP2pAuthenticationEnabled(opts *bind.FilterOpts) (*ContractAvsGovernanceSetP2pAuthenticationEnabledIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "SetP2pAuthenticationEnabled")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceSetP2pAuthenticationEnabledIterator{contract: _ContractAvsGovernance.contract, event: "SetP2pAuthenticationEnabled", logs: logs, sub: sub}, nil
}

// WatchSetP2pAuthenticationEnabled is a free log subscription operation binding the contract event 0x476bc7159a4e42477e0aa2e2e9709c55cb469b547976f91002b8db00bbd9b2f6.
//
// Solidity: event SetP2pAuthenticationEnabled(bool _isEnabled)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchSetP2pAuthenticationEnabled(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetP2pAuthenticationEnabled) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "SetP2pAuthenticationEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceSetP2pAuthenticationEnabled)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetP2pAuthenticationEnabled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetP2pAuthenticationEnabled is a log parse operation binding the contract event 0x476bc7159a4e42477e0aa2e2e9709c55cb469b547976f91002b8db00bbd9b2f6.
//
// Solidity: event SetP2pAuthenticationEnabled(bool _isEnabled)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseSetP2pAuthenticationEnabled(log types.Log) (*ContractAvsGovernanceSetP2pAuthenticationEnabled, error) {
	event := new(ContractAvsGovernanceSetP2pAuthenticationEnabled)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetP2pAuthenticationEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceSetRewardsReceiverIterator is returned from FilterSetRewardsReceiver and is used to iterate over the raw logs and unpacked data for SetRewardsReceiver events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetRewardsReceiverIterator struct {
	Event *ContractAvsGovernanceSetRewardsReceiver // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAvsGovernanceSetRewardsReceiverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceSetRewardsReceiver)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAvsGovernanceSetRewardsReceiver)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAvsGovernanceSetRewardsReceiverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceSetRewardsReceiverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceSetRewardsReceiver represents a SetRewardsReceiver event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetRewardsReceiver struct {
	Operator common.Address
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetRewardsReceiver is a free log retrieval operation binding the contract event 0xe906feea2ef60b474e22b4169bdd4de6906a84cd448cbcee99593526fe87082d.
//
// Solidity: event SetRewardsReceiver(address indexed operator, address receiver)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterSetRewardsReceiver(opts *bind.FilterOpts, operator []common.Address) (*ContractAvsGovernanceSetRewardsReceiverIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "SetRewardsReceiver", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceSetRewardsReceiverIterator{contract: _ContractAvsGovernance.contract, event: "SetRewardsReceiver", logs: logs, sub: sub}, nil
}

// WatchSetRewardsReceiver is a free log subscription operation binding the contract event 0xe906feea2ef60b474e22b4169bdd4de6906a84cd448cbcee99593526fe87082d.
//
// Solidity: event SetRewardsReceiver(address indexed operator, address receiver)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchSetRewardsReceiver(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetRewardsReceiver, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "SetRewardsReceiver", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceSetRewardsReceiver)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetRewardsReceiver", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetRewardsReceiver is a log parse operation binding the contract event 0xe906feea2ef60b474e22b4169bdd4de6906a84cd448cbcee99593526fe87082d.
//
// Solidity: event SetRewardsReceiver(address indexed operator, address receiver)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseSetRewardsReceiver(log types.Log) (*ContractAvsGovernanceSetRewardsReceiver, error) {
	event := new(ContractAvsGovernanceSetRewardsReceiver)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetRewardsReceiver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceSetStakingContractMultiplierIterator is returned from FilterSetStakingContractMultiplier and is used to iterate over the raw logs and unpacked data for SetStakingContractMultiplier events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetStakingContractMultiplierIterator struct {
	Event *ContractAvsGovernanceSetStakingContractMultiplier // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAvsGovernanceSetStakingContractMultiplierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceSetStakingContractMultiplier)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAvsGovernanceSetStakingContractMultiplier)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAvsGovernanceSetStakingContractMultiplierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceSetStakingContractMultiplierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceSetStakingContractMultiplier represents a SetStakingContractMultiplier event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetStakingContractMultiplier struct {
	StakingContract common.Address
	Multiplier      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetStakingContractMultiplier is a free log retrieval operation binding the contract event 0xa717a4f9ede961a75886c212a98163df9b8ff2e23b83c9fe3eb8a73f56cf385f.
//
// Solidity: event SetStakingContractMultiplier(address stakingContract, uint256 multiplier)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterSetStakingContractMultiplier(opts *bind.FilterOpts) (*ContractAvsGovernanceSetStakingContractMultiplierIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "SetStakingContractMultiplier")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceSetStakingContractMultiplierIterator{contract: _ContractAvsGovernance.contract, event: "SetStakingContractMultiplier", logs: logs, sub: sub}, nil
}

// WatchSetStakingContractMultiplier is a free log subscription operation binding the contract event 0xa717a4f9ede961a75886c212a98163df9b8ff2e23b83c9fe3eb8a73f56cf385f.
//
// Solidity: event SetStakingContractMultiplier(address stakingContract, uint256 multiplier)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchSetStakingContractMultiplier(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetStakingContractMultiplier) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "SetStakingContractMultiplier")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceSetStakingContractMultiplier)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetStakingContractMultiplier", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetStakingContractMultiplier is a log parse operation binding the contract event 0xa717a4f9ede961a75886c212a98163df9b8ff2e23b83c9fe3eb8a73f56cf385f.
//
// Solidity: event SetStakingContractMultiplier(address stakingContract, uint256 multiplier)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseSetStakingContractMultiplier(log types.Log) (*ContractAvsGovernanceSetStakingContractMultiplier, error) {
	event := new(ContractAvsGovernanceSetStakingContractMultiplier)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetStakingContractMultiplier", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceSetTokenIterator is returned from FilterSetToken and is used to iterate over the raw logs and unpacked data for SetToken events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetTokenIterator struct {
	Event *ContractAvsGovernanceSetToken // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAvsGovernanceSetTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceSetToken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAvsGovernanceSetToken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAvsGovernanceSetTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceSetTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceSetToken represents a SetToken event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetToken struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetToken is a free log retrieval operation binding the contract event 0xefc1fd16ea80a922086ee4e995739d59b025c1bcea6d1f67855747480c83214b.
//
// Solidity: event SetToken(address token)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterSetToken(opts *bind.FilterOpts) (*ContractAvsGovernanceSetTokenIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "SetToken")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceSetTokenIterator{contract: _ContractAvsGovernance.contract, event: "SetToken", logs: logs, sub: sub}, nil
}

// WatchSetToken is a free log subscription operation binding the contract event 0xefc1fd16ea80a922086ee4e995739d59b025c1bcea6d1f67855747480c83214b.
//
// Solidity: event SetToken(address token)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchSetToken(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetToken) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "SetToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceSetToken)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetToken", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetToken is a log parse operation binding the contract event 0xefc1fd16ea80a922086ee4e995739d59b025c1bcea6d1f67855747480c83214b.
//
// Solidity: event SetToken(address token)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseSetToken(log types.Log) (*ContractAvsGovernanceSetToken, error) {
	event := new(ContractAvsGovernanceSetToken)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceSetNewSupportedStakingContractsIterator is returned from FilterSetNewSupportedStakingContracts and is used to iterate over the raw logs and unpacked data for SetNewSupportedStakingContracts events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetNewSupportedStakingContractsIterator struct {
	Event *ContractAvsGovernanceSetNewSupportedStakingContracts // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAvsGovernanceSetNewSupportedStakingContractsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceSetNewSupportedStakingContracts)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAvsGovernanceSetNewSupportedStakingContracts)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAvsGovernanceSetNewSupportedStakingContractsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceSetNewSupportedStakingContractsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceSetNewSupportedStakingContracts represents a SetNewSupportedStakingContracts event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetNewSupportedStakingContracts struct {
	StakingContracts []common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetNewSupportedStakingContracts is a free log retrieval operation binding the contract event 0xdc9dd92c894ed28df4cdf880b63e1f5d6a6a32870a392e08c13ddfe18488caea.
//
// Solidity: event setNewSupportedStakingContracts(address[] stakingContracts)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterSetNewSupportedStakingContracts(opts *bind.FilterOpts) (*ContractAvsGovernanceSetNewSupportedStakingContractsIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "setNewSupportedStakingContracts")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceSetNewSupportedStakingContractsIterator{contract: _ContractAvsGovernance.contract, event: "setNewSupportedStakingContracts", logs: logs, sub: sub}, nil
}

// WatchSetNewSupportedStakingContracts is a free log subscription operation binding the contract event 0xdc9dd92c894ed28df4cdf880b63e1f5d6a6a32870a392e08c13ddfe18488caea.
//
// Solidity: event setNewSupportedStakingContracts(address[] stakingContracts)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchSetNewSupportedStakingContracts(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetNewSupportedStakingContracts) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "setNewSupportedStakingContracts")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceSetNewSupportedStakingContracts)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "setNewSupportedStakingContracts", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetNewSupportedStakingContracts is a log parse operation binding the contract event 0xdc9dd92c894ed28df4cdf880b63e1f5d6a6a32870a392e08c13ddfe18488caea.
//
// Solidity: event setNewSupportedStakingContracts(address[] stakingContracts)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseSetNewSupportedStakingContracts(log types.Log) (*ContractAvsGovernanceSetNewSupportedStakingContracts, error) {
	event := new(ContractAvsGovernanceSetNewSupportedStakingContracts)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "setNewSupportedStakingContracts", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
