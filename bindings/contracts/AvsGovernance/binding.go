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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_extensionImplementation\",\"type\":\"address\"},{\"internalType\":\"contractIOthenticRegistry\",\"name\":\"_othenticRegistryAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccessControlInvalidMultiplierSyncer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllowlistDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllowlistEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyAvsName\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptySharedSecurityProvidersList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FlowIsCurrentlyPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FlowIsCurrentlyUnpaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAllowlistAuthToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBlsRegistrationSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMultiplier\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMultiplierNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRewardsReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"arrayIndex\",\"type\":\"uint256\"}],\"name\":\"InvalidSharedSecurityProviderList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSlashingRate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStakingContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"MissingAuthToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ModificationDelayNotPassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NativeCoinNotSupportedForEigenRewardsError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughVotingPower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numOfOperatorsLimit\",\"type\":\"uint256\"}],\"name\":\"NumOfOperatorsLimitReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorNotRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorStillRegisteredToSharedSecurityProviders\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PauseFlowIsAlreadyPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StakingContractsNotInAscendingOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TreasuryWithdrawRewardsFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnpausingFlowIsAlreadyUnpaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DepositRewardsBackFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"_pausableFlow\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_pauser\",\"type\":\"address\"}],\"name\":\"FlowPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"_pausableFlowFlag\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_unpauser\",\"type\":\"address\"}],\"name\":\"FlowUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"NativeCoinNotSupportedForEigenRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorEjectedFromNetwork\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[4]\",\"name\":\"blsKey\",\"type\":\"uint256[4]\"}],\"name\":\"OperatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorRegisteredToEigenLayer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorRegisteredToSymbiotic\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorUnregisteredToEigenLayer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorUnregisteredToSymbiotic\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"QueuedRewardsReceiverModification\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"revertData\",\"type\":\"bytes\"}],\"name\":\"RewardsCoordinatorReverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"avsName\",\"type\":\"string\"}],\"name\":\"SetAvsName\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isAllowlisted\",\"type\":\"bool\"}],\"name\":\"SetIsAllowlisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_isEnabled\",\"type\":\"bool\"}],\"name\":\"SetP2pAuthenticationEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"SetRewardsReceiver\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"stakingContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"multiplier\",\"type\":\"uint256\"}],\"name\":\"SetStakingContractMultiplier\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SetToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"stakingContracts\",\"type\":\"address[]\"}],\"name\":\"setNewSupportedStakingContracts\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXTENSION_IMPLEMENTATION\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"avsName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"avsTreasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"completeRewardsReceiverModification\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_startTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_duration\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"createEigenRewardsSubmission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIRewardsCoordinator.OperatorReward[]\",\"name\":\"_operators\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"_lastPayedTask\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_rewardsData\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"_remoteEid\",\"type\":\"uint32\"}],\"name\":\"createOperatorDirectedAVSRewardsSubmission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"ejectOperatorFromNetwork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIsAllowlisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"getOperatorRestakedStrategies\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRestakeableStrategies\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"getRewardsReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"avsGovernanceMultisigOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operationsMultisig\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"communityMultisig\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"othenticRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"messageHandler\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"avsTreasury\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"avsDirectoryContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowlistSigner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"avsName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"blsAuthSingleton\",\"type\":\"address\"}],\"internalType\":\"structIAvsGovernance.InitializationParams\",\"name\":\"_initializationParams\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_pausableFlow\",\"type\":\"bytes4\"}],\"name\":\"isFlowPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_isPaused\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isOperatorRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxEffectiveBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingContract\",\"type\":\"address\"}],\"name\":\"minStakeAmountPerStakingContract\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minVotingPower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingContract\",\"type\":\"address\"}],\"name\":\"multiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numOfActiveOperators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"p2pAuthenticationEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_pausableFlow\",\"type\":\"bytes4\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newRewardsReceiver\",\"type\":\"address\"}],\"name\":\"queueRewardsReceiverModification\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[4]\",\"name\":\"blsKey\",\"type\":\"uint256[4]\"},{\"internalType\":\"address\",\"name\":\"rewardsReceiver\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"signature\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBLSAuthLibrary.Signature\",\"name\":\"blsRegistrationSignature\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"authToken\",\"type\":\"bytes\"}],\"internalType\":\"structIAvsGovernance.OperatorRegistrationParams\",\"name\":\"_operatorRegistrationParams\",\"type\":\"tuple\"}],\"name\":\"registerAsOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"}],\"name\":\"registerAvsToEigenLayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registerAvsToSymbiotic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"name\":\"_eigenSig\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_authToken\",\"type\":\"bytes\"}],\"name\":\"registerOperatorToEigenLayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint48\",\"name\":\"deadline\",\"type\":\"uint48\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIAvsGovernance.SymbioticOptInSignature\",\"name\":\"_symbioticSig\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_authToken\",\"type\":\"bytes\"}],\"name\":\"registerOperatorToSymbiotic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_p2pAuthenticationEnabled\",\"type\":\"bool\"}],\"name\":\"setP2pAuthenticationEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"stakingContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"multiplier\",\"type\":\"uint256\"},{\"internalType\":\"enumIAvsGovernance.SharedSecurityProvider\",\"name\":\"sharedSecurityProvider\",\"type\":\"uint8\"}],\"internalType\":\"structIAvsGovernance.VotingPowerMultiplier\",\"name\":\"_votingPowerMultiplier\",\"type\":\"tuple\"}],\"name\":\"setStakingContractMultiplier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"stakingContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"multiplier\",\"type\":\"uint256\"},{\"internalType\":\"enumIAvsGovernance.SharedSecurityProvider\",\"name\":\"sharedSecurityProvider\",\"type\":\"uint8\"}],\"internalType\":\"structIAvsGovernance.VotingPowerMultiplier[]\",\"name\":\"_votingPowerMultipliers\",\"type\":\"tuple[]\"}],\"name\":\"setStakingContractMultiplierBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"stakingContract\",\"type\":\"address\"},{\"internalType\":\"enumIAvsGovernance.SharedSecurityProvider\",\"name\":\"sharedSecurityProvider\",\"type\":\"uint8\"}],\"internalType\":\"structIAvsGovernance.StakingContractInfo[]\",\"name\":\"_stakingContractsDetails\",\"type\":\"tuple[]\"}],\"name\":\"setSupportedStakingContracts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingContracts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_pausableFlow\",\"type\":\"bytes4\"}],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unregisterAsOperatorFromAvs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unregisterAsOperatorFromEigenLayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint48\",\"name\":\"deadline\",\"type\":\"uint48\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIAvsGovernance.SymbioticOptOutSignature\",\"name\":\"_unregistrationSignature\",\"type\":\"tuple\"}],\"name\":\"unregisterAsOperatorFromSymbiotic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"votingPower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_stakingContracts\",\"type\":\"address[]\"}],\"name\":\"votingPowerPerStakingContracts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506004361061027f5760003560e01c806382616a011161015c578063bac1e94b116100ce578063e481af9d11610087578063e481af9d14610574578063e6474b0f1461057c578063eab90f1c14610584578063efd9697814610597578063f6e258cc146105d1578063fab57b8f146105e45761027f565b8063bac1e94b1461050d578063c07473f614610520578063c7a93c6f14610533578063cab1a5ef14610546578063cbd969c91461054e578063d547741f146105615761027f565b8063a217fddf11610120578063a217fddf146104c7578063a88171ee146104cf578063a9b3f8b7146104d7578063aa0dadf0146104ea578063ab8cac54146104f2578063b525fa88146105055761027f565b806382616a011461046857806383e451701461047b578063851bb7251461048e57806391d14854146104a157806393652277146104b45761027f565b80632f2ff15d116101f557806341b92a29116101b957806341b92a291461040a578063560d54841461041f5780635e95cee21461042757806362de73931461043a5780636b1906f81461044d5780637897dec3146104605761027f565b80632f2ff15d146103a957806333cfb7b7146103bc57806336568abe146103dc57806336fffde0146103ef5780633aa83ec7146103f75761027f565b806313fe45211161024757806313fe4521146103205780631b21ba72146103415780631dd2f74b14610354578063226def0414610367578063248a9ca31461038e578063263fe7c7146103a15761027f565b806301ffc9a7146102aa5780630399f004146102d257806306255b01146102da5780631246193e146102ed578063130bdcd91461030d575b6102a87f0000000000000000000000003d2f87731394b1ebd0fdfd906fe56e8c132418706105f7565b005b6102bd6102b836600461499c565b610620565b60405190151581526020015b60405180910390f35b6102a8610657565b6102a86102e83660046149c6565b610766565b6102f561090e565b6040516001600160a01b0390911681526020016102c9565b6102a861031b366004614a00565b61092a565b61033361032e366004614a89565b610982565b6040519081526020016102c9565b6102a861034f366004614a89565b6109ae565b6102a8610362366004614b44565b610aff565b6102f57f0000000000000000000000003d2f87731394b1ebd0fdfd906fe56e8c1324187081565b61033361039c366004614c11565b610b4b565b6102a8610b6d565b6102a86103b7366004614c2a565b610e77565b6103cf6103ca366004614a89565b610e99565b6040516102c99190614c9f565b6102a86103ea366004614c2a565b610f37565b610333610f6a565b6102a861040536600461499c565b610f7d565b610412610fa4565b6040516102c99190614d02565b6103cf61103f565b6102f5610435366004614a89565b6110a9565b6102a8610448366004614da6565b6110d9565b6102bd61045b366004614a89565b6115aa565b6103336115d8565b6102a8610476366004614ea6565b6115eb565b6102a8610489366004614eff565b61160d565b6102a861049c366004614fc1565b611636565b6102bd6104af366004614c2a565b6116b9565b6102a86104c2366004615018565b6116f1565b610333600081565b610333611712565b6103336104e5366004614a89565b611725565b6102a8611751565b6102a86105003660046150a7565b61190f565b6102bd6119bf565b6102a861051b36600461499c565b6119dc565b61033361052e366004614a89565b611a03565b6102a86105413660046150f0565b611a8e565b6102bd611af3565b61033361055c36600461510d565b611b09565b6102a861056f366004614c2a565b611b6f565b6103cf611b8b565b6102a8611b9a565b6102a8610592366004615194565b611c9c565b6102bd6105a536600461499c565b6001600160e01b0319166000908152600080516020615bfc833981519152602052604090205460ff1690565b6102a86105df366004614a89565b611fe1565b6102a86105f23660046151d5565b612192565b3660008037600080366000845af43d6000803e808015610616573d6000f35b3d6000fd5b505050565b60006001600160e01b03198216637965db0b60e01b148061065157506301ffc9a760e01b6001600160e01b03198316145b92915050565b61065f6122a0565b33600090815260079190910160205260408120549003610692576040516325ec6c1f60e01b815260040160405180910390fd5b600080516020615c5c8339815191526106aa816122aa565b6106b26122f2565b6106ba6122a0565b600501546040516351b27a6d60e11b81523360048201526001600160a01b039091169063a364f4da90602401600060405180830381600087803b15801561070057600080fd5b505af1158015610714573d6000803e3d6000fd5b50506040513381527f7877897edb095f8cd999adcb582e04128ffb91d12bf5139ca7b8a1567a8453009250602001905060405180910390a16107636001600080516020615c3c83398151915255565b50565b61076e6122a0565b336000908152600791909101602052604081205490036107a1576040516325ec6c1f60e01b815260040160405180910390fd5b600080516020615c5c8339815191526107b9816122aa565b6107c16122f2565b7f00000000000000000000000052fdd2102aaca3399f818f19353b84d1dd771b2d6001600160a01b0316635ee43d376040518163ffffffff1660e01b8152600401602060405180830381865afa15801561081f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108439190615210565b6001600160a01b03166393f79bc33330610860602087018761522d565b61086d6020880188615248565b6040518663ffffffff1660e01b815260040161088d9594939291906152b7565b600060405180830381600087803b1580156108a757600080fd5b505af11580156108bb573d6000803e3d6000fd5b50506040513381527ffc7eb04654394ed6ffe862307d7824361d3481d69ea2775b77e85837befe77379250602001905060405180910390a161090a6001600080516020615c3c83398151915255565b5050565b60006109186122a0565b600401546001600160a01b0316919050565b63ef0892d160e01b61093b8161233e565b60006109456122a0565b905060005b8381101561097b5761097382868684818110610968576109686152f2565b905060600201612348565b60010161094a565b5050505050565b600061098c6122a0565b6001600160a01b03909216600090815260129290920160205250604090205490565b6109b66122a0565b336000908152600791909101602052604081205490036109e9576040516325ec6c1f60e01b815260040160405180910390fd5b7fd93c394bc0520b0dc9db435f0321020356363275d7d80fbdc9f5296937698fc6610a13816122aa565b6001600160a01b038216610a395760405162cc6ac760e01b815260040160405180910390fd5b6000610a436122a0565b336000908152600c820160205260408120805460ff19166001179055600b82015491925090610a72904261531e565b6040805180820182526001600160a01b038781168083526020808401868152336000818152600e8b018452879020955186546001600160a01b03191695169490941785555160019094019390935583519081529182018490529293507f0d8cfa10a3087b28d3c226ad9a37314860e7c3c0505a25a39e3cdefb3118a98a910160405180910390a250505050565b6378b4401360e11b610b108161233e565b7f37b1f16af49fd5d40e5dc679bf35afa8bb06df10cb8308e0aa633882d2a81201610b3a816122aa565b61061b610b456122a0565b84612402565b6000908152600080516020615c1c833981519152602052604090206001015490565b610b756122a0565b33600090815260079190910160205260408120549003610ba8576040516325ec6c1f60e01b815260040160405180910390fd5b600080516020615c5c833981519152610bc0816122aa565b610bc86122f2565b6000610bd26122a0565b9050600160058201546040516349075da360e01b81523060048201523360248201526001600160a01b03909116906349075da390604401602060405180830381865afa158015610c26573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c4a9190615347565b6001811115610c5b57610c5b615331565b14158015610d5857507f00000000000000000000000052fdd2102aaca3399f818f19353b84d1dd771b2d6001600160a01b0316635ee43d376040518163ffffffff1660e01b8152600401602060405180830381865afa158015610cc2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ce69190615210565b6040516308834cb560e21b81523360048201523060248201526001600160a01b03919091169063220d32d490604401602060405180830381865afa158015610d32573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d569190615364565b155b15610e3f5760088101546001600160a01b03168015801590610dce576040516311c7e79960e21b81523360048201526001600160a01b0383169063471f9e6490602401600060405180830381600087803b158015610db557600080fd5b505af1158015610dc9573d6000803e3d6000fd5b505050505b610dd8833361274b565b8015610e385760405163e9ecc1cb60e01b81523360048201526001600160a01b0383169063e9ecc1cb90602401600060405180830381600087803b158015610e1f57600080fd5b505af1158015610e33573d6000803e3d6000fd5b505050505b5050610e5f565b604051633c1d10e160e11b81523360048201526024015b60405180910390fd5b506107636001600080516020615c3c83398151915255565b610e8082610b4b565b610e898161233e565b610e9383836127b1565b50505050565b60607f00000000000000000000000052fdd2102aaca3399f818f19353b84d1dd771b2d6001600160a01b03166390dea2af83610ed3612856565b306040518463ffffffff1660e01b8152600401610ef293929190615381565b600060405180830381865afa158015610f0f573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261065191908101906153b6565b6001600160a01b0381163314610f605760405163334bd91960e11b815260040160405180910390fd5b61061b8282612862565b6000610f746122a0565b60100154905090565b80610f87816122aa565b6001600160e01b03198216610f9b8161233e565b61061b836128de565b6060610fae6122a0565b600a018054610fbc90615444565b80601f0160208091040260200160405190810160405280929190818152602001828054610fe890615444565b80156110355780601f1061100a57610100808354040283529160200191611035565b820191906000526020600020905b81548152906001019060200180831161101857829003601f168201915b5050505050905090565b60606110496122a0565b60090180548060200260200160405190810160405280929190818152602001828054801561103557602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611082575050505050905090565b60006110b36122a0565b6001600160a01b039283166000908152600d919091016020526040902054909116919050565b638a70a0eb60e01b6110ea8161233e565b60006110f46122a0565b905060008160040160009054906101000a90046001600160a01b0316905060007f00000000000000000000000052fdd2102aaca3399f818f19353b84d1dd771b2d6001600160a01b0316638a2fc4e36040518163ffffffff1660e01b8152600401602060405180830381865afa158015611172573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111969190615210565b90506111cf6040518060800160405280600063ffffffff168152602001600063ffffffff16815260200160008152602001600081525090565b868060200190518101906111e39190615478565b606085015260408085019190915263ffffffff91821660208086019190915292909116835260048087015482516321df0da760e01b815292516000946001600160a01b03909216936321df0da7938181019392918290030181865afa158015611250573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112749190615210565b9050600060608660040160009054906101000a90046001600160a01b03166001600160a01b031663a734f06e6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156112cf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112f39190615210565b6001600160a01b0316836001600160a01b0316146114d05761131c30308d87604001518a61297f565b915081156114cb57604084015161133f906001600160a01b038516908790612a97565b50846001600160a01b0316639cb9a5fa60e01b306113678f8860000151896020015189612b79565b604051602401611378929190615553565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199094169390931790925290516113b69190615634565b6000604051808303816000865af19150503d80600081146113f3576040519150601f19603f3d011682016040523d82523d6000602084013e6113f8565b606091505b509092509050816114cb57604084015161141e906001600160a01b038516908890612a97565b5060408085015190516356a8904960e11b81526000916001600160a01b0389169163ad512092916114559160040190815260200190565b6020604051808303816000875af1158015611474573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114989190615364565b9050806114c9576040517fed3143332c2d71fbe302b126ea59ee4fb74a1699cb820f835a913b15c027798d90600090a15b505b6114fa565b6040517fb3f3d8f7b39bd24844dec3cce1a06cfee42f663826e9b790c6a36fb8c05dbf8e90600090a15b816115715760005b84606001518110156115385760008d8281518110611522576115226152f2565b6020908102919091018101510152600101611502565b507fab7e71f561320fa16a226c9744bf4d230ce7faf885df7bbc890baf339faa0a4e816040516115689190614d02565b60405180910390a15b61159c878d6040516020016115869190615650565b6040516020818303038152906040528d8c612c84565b505050505050505050505050565b60006115b46122a0565b6001600160a01b039092166000908152600792909201602052506040902054151590565b60006115e26122a0565b60010154905090565b63ef0892d160e01b6115fc8161233e565b61090a6116076122a0565b83612348565b60006116176122a0565b9050611624818484612d1f565b610e9361162f6122a0565b3386612dd4565b6378b4401360e11b6116478161233e565b61164f6122a0565b6005015460405163a98fb35560e01b81526001600160a01b039091169063a98fb355906116829086908690600401615663565b600060405180830381600087803b15801561169c57600080fd5b505af11580156116b0573d6000803e3d6000fd5b50505050505050565b6000918252600080516020615c1c833981519152602090815260408084206001600160a01b0393909316845291905290205460ff1690565b60006116fb6122a0565b9050611708818484612d1f565b610e933385612e74565b600061171c6122a0565b60110154905090565b600061172f6122a0565b6001600160a01b03909216600090815260139290920160205250604090205490565b6378b4401360e11b6117628161233e565b7f00000000000000000000000052fdd2102aaca3399f818f19353b84d1dd771b2d6001600160a01b031663e45f40be6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156117c0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117e49190615210565b6001600160a01b03166387140b5b6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561181e57600080fd5b505af1158015611832573d6000803e3d6000fd5b505050507f00000000000000000000000052fdd2102aaca3399f818f19353b84d1dd771b2d6001600160a01b0316630bb79bc06040518163ffffffff1660e01b8152600401602060405180830381865afa158015611894573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118b89190615210565b60405163b7d8e1a960e01b81523060048201526001600160a01b03919091169063b7d8e1a990602401600060405180830381600087803b1580156118fb57600080fd5b505af115801561097b573d6000803e3d6000fd5b6119176122a0565b336000908152600791909101602052604090205415611949576040516342ee68b560e01b815260040160405180910390fd5b600080516020615c5c833981519152611961816122aa565b6119696122f2565b60006119736122a0565b905061198b8161198660e0860186615248565b612d1f565b6119a781338560a081016119a28160808401614a89565b612f9e565b5061090a6001600080516020615c3c83398151915255565b60006119c96122a0565b60050154600160a01b900460ff16919050565b806119e681613348565b6001600160e01b031982166119fa8161233e565b61061b8361338f565b600080611a0e6122a0565b90506000611a788283600901805480602002602001604051908101604052809291908181526020018280548015611a6e57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611a50575b505050505061342c565b915050611a8682858361364c565b949350505050565b6378b4401360e11b611a9f8161233e565b60405182151581527f476bc7159a4e42477e0aa2e2e9709c55cb469b547976f91002b8db00bbd9b2f69060200160405180910390a181611add6122a0565b601701805460ff19169115159190911790555050565b6000611afd6122a0565b6017015460ff16919050565b600080611b146122a0565b90506000611b558286868080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525061342c92505050565b915050611b6382878361364c565b925050505b9392505050565b611b7882610b4b565b611b818161233e565b610e938383612862565b6060611b95612856565b905090565b611ba26122a0565b33600090815260079190910160205260408120549003611bd5576040516325ec6c1f60e01b815260040160405180910390fd5b6000611bdf6122a0565b336000908152600e82016020526040902060010154909150421015611c1757604051638ce7a3f160e01b815260040160405180910390fd5b336000908152600e82016020526040902054611c3d9082906001600160a01b0316613707565b336000818152600c830160209081526040808320805460ff1916905580518082018252838152808301848152948452600e9095019091529020915182546001600160a01b0319166001600160a01b039091161782555160019190910155565b6378b4401360e11b611cad8161233e565b6000611cb76122a0565b905060008160040160009054906101000a90046001600160a01b0316905060007f00000000000000000000000052fdd2102aaca3399f818f19353b84d1dd771b2d6001600160a01b0316638a2fc4e36040518163ffffffff1660e01b8152600401602060405180830381865afa158015611d35573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d599190615210565b905060008360040160009054906101000a90046001600160a01b03166001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015611db2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611dd69190615210565b90508360040160009054906101000a90046001600160a01b03166001600160a01b031663a734f06e6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611e2d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e519190615210565b6001600160a01b0316816001600160a01b031603611e8257604051634acb03b560e01b815260040160405180910390fd5b611e9030306000898761297f565b611ead576040516359062f7f60e01b815260040160405180910390fd5b611ec16001600160a01b038216838861376c565b604080516001808252818301909252600091816020015b6040805160a08101825260608082526000602080840182905293830181905290820181905260808201528252600019909201910181611ed85790505090506040518060a00160405280611f2a876137f6565b8152602001836001600160a01b031681526020018881526020018a63ffffffff1681526020018963ffffffff1681525081600081518110611f6d57611f6d6152f2565b602090810291909101015260405163fce36c7d60e01b81526001600160a01b0384169063fce36c7d90611fa4908490600401615677565b600060405180830381600087803b158015611fbe57600080fd5b505af1158015611fd2573d6000803e3d6000fd5b50505050505050505050505050565b611fe96122f2565b638a70a0eb60e01b611ffa8161233e565b60006120046122a0565b9050600160058201546040516349075da360e01b81523060048201526001600160a01b038681166024830152909116906349075da390604401602060405180830381865afa15801561205a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061207e9190615347565b600181111561208f5761208f615331565b036121335760058101546040516351b27a6d60e11b81526001600160a01b0385811660048301529091169063a364f4da90602401600060405180830381600087803b1580156120dd57600080fd5b505af11580156120f1573d6000803e3d6000fd5b50506040516001600160a01b03861681527f7877897edb095f8cd999adcb582e04128ffb91d12bf5139ca7b8a1567a8453009250602001905060405180910390a15b61213d818461274b565b6040516001600160a01b03841681527f98ac533c958bf5d682a0590dc21d14cc6a8acaaa7886880124cfb4531266f2f29060200160405180910390a150506107636001600080516020615c3c83398151915255565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03166000811580156121d75750825b90506000826001600160401b031660011480156121f35750303b155b905081158015612201575080155b1561221f5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561224957845460ff60401b1916600160401b1785555b612252866139a5565b831561229857845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050565b6000611b95613d17565b6001600160e01b031981166000908152600080516020615bfc833981519152602052604090205460ff16156107635760405163722fdba960e01b815260040160405180910390fd5b600080516020615c3c83398151915280546001190161232457604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b6001600080516020615c3c83398151915255565b6107638133613d45565b6001600160601b036020820135111561237457604051631bc4bcf760e21b815260040160405180910390fd5b60208101803590601384019060009061238d9085614a89565b6001600160a01b03168152602080820192909252604001600020919091557fa717a4f9ede961a75886c212a98163df9b8ff2e23b83c9fe3eb8a73f56cf385f906123d990830183614a89565b604080516001600160a01b03909216825260208085013590830152015b60405180910390a15050565b61241060098301600061496a565b6000805b825181101561270b576000838281518110612431576124316152f2565b60200260200101519050826001600160a01b031681600001516001600160a01b031610156124725760405163b854d24160e01b815260040160405180910390fd5b805192506001600160a01b03831661249d576040516307b18df360e31b815260040160405180910390fd5b8051600986018054600180820183556000928352602080842090920180546001600160a01b0319166001600160a01b0395861617905581850151855190941683526016890190915260409091208054909160ff1990911690838181111561250657612506615331565b021790555060018160200151600181111561252357612523615331565b148015612619575083828151811061253d5761253d6152f2565b6020026020010151600001516001600160a01b031663ce9b79306040518163ffffffff1660e01b8152600401602060405180830381865afa158015612586573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125aa9190615210565b6040516368adba0760e11b81523060601b60048201526001600160a01b03919091169063d15b740e90602401602060405180830381865afa1580156125f3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126179190615723565b155b1561270257838281518110612630576126306152f2565b6020026020010151600001516001600160a01b031663ce9b79306040518163ffffffff1660e01b8152600401602060405180830381865afa158015612679573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061269d9190615210565b6040516323f752d560e01b81526000600482015260001960248201526001600160a01b0391909116906323f752d590604401600060405180830381600087803b1580156126e957600080fd5b505af11580156126fd573d6000803e3d6000fd5b505050505b50600101612414565b507fdc9dd92c894ed28df4cdf880b63e1f5d6a6a32870a392e08c13ddfe18488caea8360090160405161273e919061573c565b60405180910390a1505050565b6127558282613d7e565b81600101600081546127669061578c565b909155506001600160a01b0381166000818152600784016020526040808220829055517f6f42117a557500c705ddf040a619d86f39101e6b74ac20d7b3e5943ba473fc7f9190a25050565b6000600080516020615c1c8339815191526127cc84846116b9565b61284c576000848152602082815260408083206001600160a01b03871684529091529020805460ff191660011790556128023390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610651565b6000915050610651565b6060611b956000613e1b565b6000600080516020615c1c83398151915261287d84846116b9565b1561284c576000848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610651565b6001600160e01b031981166000908152600080516020615bfc833981519152602081905260409091205460ff16156129295760405163dfe10d7d60e01b815260040160405180910390fd5b6001600160e01b0319821660009081526020829052604090819020805460ff19166001179055517f95c3658c5e0c74e20cf12db371b9b67d26e97a1937f6d2284f88cc44d036b4f6906123f690849033906157a3565b60006001600160a01b03851615612a1157604051633256b4d160e01b81526001600160a01b0386811660048301526024820186905260448201859052831690633256b4d1906064016020604051808303816000875af11580156129e6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a0a9190615364565b9050612a8e565b604051633256b4d160e01b81526001600160a01b0387811660048301526024820186905260448201859052831690633256b4d1906064016020604051808303816000875af1158015612a67573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a8b9190615364565b90505b95945050505050565b604051636eb1769f60e11b81523060048201526001600160a01b038381166024830152600091829186169063dd62ed3e90604401602060405180830381865afa158015612ae8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b0c9190615723565b9050612a8e856001600160a01b03811663095ea7b387612b2c888761531e565b6040516001600160a01b0390921660248301526044820152606401604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050613f40565b60606000612b856122a0565b60408051600180825281830190925291925060009190816020015b612bee6040518060c001604052806060815260200160006001600160a01b0316815260200160608152602001600063ffffffff168152602001600063ffffffff168152602001606081525090565b815260200190600190039081612ba05790505090506040518060c00160405280612c17846137f6565b8152602001856001600160a01b031681526020018881526020018763ffffffff1681526020018663ffffffff1681526020016040518060200160405280600081525081525081600081518110612c6f57612c6f6152f2565b60209081029190910101529695505050505050565b6000612c908484613fd8565b60028601546040519192506001600160a01b0316906382646a5890612cbb90849086906020016157c6565b6040516020818303038152906040526040518263ffffffff1660e01b8152600401612ce69190614d02565b600060405180830381600087803b158015612d0057600080fd5b505af1158015612d14573d6000803e3d6000fd5b505050505050505050565b6005830154600160a01b900460ff161561061b576000819003612d5957818160405163471b4be960e11b8152600401610e56929190615663565b612db730338560060160009054906101000a90046001600160a01b031685858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594939250506140489050565b61061b5760405163d2342ec760e01b815260040160405180910390fd5b6005830154604051639926ee7d60e01b81526001600160a01b0390911690639926ee7d90612e0890859085906004016157ee565b600060405180830381600087803b158015612e2257600080fd5b505af1158015612e36573d6000803e3d6000fd5b50506040516001600160a01b03851681527f23c66044e8b2ab58f45f4aa7eeb03c47c231fec8b7cf140f5b709455e5c1aa1d9250602001905061273e565b60007f00000000000000000000000052fdd2102aaca3399f818f19353b84d1dd771b2d6001600160a01b0316635ee43d376040518163ffffffff1660e01b8152600401602060405180830381865afa158015612ed4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ef89190615210565b8251602084015160405163ced44ba760e01b81529293506001600160a01b0384169263ced44ba792612f3292889230929190600401615839565b600060405180830381600087803b158015612f4c57600080fd5b505af1158015612f60573d6000803e3d6000fd5b50506040516001600160a01b03861681527f6f6f010344f25a8f5774e47688e7db5f23ef1071452f6048e1147d8de8cfe2d19250602001905061273e565b6001600160a01b038116612fc45760405162cc6ac760e01b815260040160405180910390fd5b612fce8582613707565b600f85015460018601548111612ffa5760405163c77b407760e01b815260048101829052602401610e56565b6015860154604051633cf65e3b60e01b81526001600160a01b0390911690633cf65e3b90613032908690899030908a90600401615873565b602060405180830381865afa15801561304f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906130739190615364565b613090576040516314532dfd60e11b815260040160405180910390fd5b6008860154600987018054604080516020808402820181019092528281526001600160a01b039094169384151593600093849361310e938e9390929091830182828015611a6e576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611a5057505050505061342c565b91509150600061311f8b8b8461364c565b905060008b6010015482101580156131c357506040516329ea021d60e01b81526001600160a01b037f00000000000000000000000052fdd2102aaca3399f818f19353b84d1dd771b2d16906329ea021d90613182908e90889030906004016158e9565b602060405180830381865afa15801561319f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906131c39190615364565b9050806131e357604051631c33ce8d60e11b815260040160405180910390fd5b841561324e5760405163094e7a3f60e41b81526001600160a01b038716906394e7a3f09061321b908e9086908f908e90600401615955565b600060405180830381600087803b15801561323557600080fd5b505af1158015613249573d6000803e3d6000fd5b505050505b61325b8c8c848d8c614072565b8b6001016000815461326c90615988565b90915550506001600160a01b038a16600090815260078c01602052604090206001905583156132fa576040516376c56c1b60e11b81526001600160a01b0386169063ed8ad836906132c7908d9085908e908d90600401615955565b600060405180830381600087803b1580156132e157600080fd5b505af11580156132f5573d6000803e3d6000fd5b505050505b50505050856001600160a01b03167f54bc9cf83c2eb0f2ad1abf6e4fab882964404622ba2df6b5a9356a18d3aac0558660405161333791906159a1565b60405180910390a250505050505050565b6001600160e01b031981166000908152600080516020615bfc833981519152602052604090205460ff16610763576040516368c87f3360e11b815260040160405180910390fd5b6001600160e01b031981166000908152600080516020615bfc833981519152602081905260409091205460ff166133d957604051635bfd2da760e11b815260040160405180910390fd5b6001600160e01b0319821660009081526020829052604090819020805460ff19169055517fc7e56e17b0a6c4b467df6495e1eda1baecd7ba20604e80c1058ac06f4578d85e906123f690849033906157a3565b606080600083516001600160401b0381111561344a5761344a614aa6565b60405190808252806020026020018201604052801561349c57816020015b61348960408051606081018252600080825260208201819052909182015290565b8152602001906001900390816134685790505b509050600084516001600160401b038111156134ba576134ba614aa6565b60405190808252806020026020018201604052801561350c57816020015b6134f960408051606081018252600080825260208201819052909182015290565b8152602001906001900390816134d85790505b50905060005b855181101561363e57600086828151811061352f5761352f6152f2565b6020908102919091018101516001600160a01b038116600081815260168c018452604080822054815160608101835284815293835260128e01865291819020549483019490945291935060ff90911691810182600181111561359357613593615331565b8152508584815181106135a8576135a86152f2565b6020908102919091018101919091526001600160a01b038316600090815260138b019091526040812054908190036135de575060015b6040518060600160405280846001600160a01b0316815260200182815260200183600181111561361057613610615331565b815250858581518110613625576136256152f2565b6020026020010181905250836001019350505050613512565b5090925090505b9250929050565b6011830154604051630e1b4e1f60e11b81526000919082906001600160a01b037f00000000000000000000000052fdd2102aaca3399f818f19353b84d1dd771b2d1690631c369c3e906136a7908890889030906004016159b0565b602060405180830381865afa1580156136c4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906136e89190615723565b905081811180156136f95750600082115b15612a8e5750949350505050565b336000818152600d8401602090815260409182902080546001600160a01b0319166001600160a01b03861690811790915591519182527fe906feea2ef60b474e22b4169bdd4de6906a84cd448cbcee99593526fe87082d910160405180910390a25050565b604051636eb1769f60e11b81523060048201526001600160a01b0383811660248301526000919085169063dd62ed3e90604401602060405180830381865afa1580156137bc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906137e09190615723565b9050610e9384846137f1858561531e565b614111565b6060600061385e8384600901805480602002602001604051908101604052809291908181526020018280548015611a6e576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611a5057505050505061342c565b91505060008082516001600160401b0381111561387d5761387d614aa6565b6040519080825280602002602001820160405280156138c257816020015b604080518082019091526000808252602082015281526020019060019003908161389b5790505b50905060005b835181101561399a5760008482815181106138e5576138e56152f2565b602002602001015160400151600181111561390257613902615331565b03613992576040518060400160405280858381518110613924576139246152f2565b6020026020010151600001516001600160a01b03168152602001858381518110613950576139506152f2565b6020026020010151602001516001600160601b031681525082848061397490615988565b955081518110613986576139866152f2565b60200260200101819052505b6001016138c8565b509081529392505050565b6139ad6141c5565b60006139bc6020830183614a89565b6001600160a01b031614806139e9575060006139de6040830160208401614a89565b6001600160a01b0316145b80613a0c57506000613a016060830160408401614a89565b6001600160a01b0316145b80613a2f57506000613a2460a0830160808401614a89565b6001600160a01b0316145b80613a5257506000613a4760c0830160a08401614a89565b6001600160a01b0316145b80613a7557506000613a6a60e0830160c08401614a89565b6001600160a01b0316145b80613a9957506000613a8e610100830160e08401614a89565b6001600160a01b0316145b80613abe57506000613ab361014083016101208401614a89565b6001600160a01b0316145b15613adc5760405163d92e233d60e01b815260040160405180910390fd5b6000613aeb6020830183614a89565b90506000613aff6040840160208501614a89565b90506000613b136060850160408601614a89565b90506000613b2760a0860160808701614a89565b9050366000613b3a610100880188615248565b91509150613b49868686614210565b613b5486868661424d565b613b5c614273565b6000613b666122a0565b6002810180546001600160a01b0319166001600160a01b0387161790559050613b96638a70a0eb60e01b856127b1565b50613ba760c0890160a08a01614a89565b6004820180546001600160a01b0319166001600160a01b0392909216919091179055613bda610100890160e08a01614a89565b6006820180546001600160a01b0319166001600160a01b039290921691909117905562093a80600b820155613c1560e0890160c08a01614a89565b6005820180546001600160a01b0319166001600160a01b03929092169190911790556064600f820155613c5061014089016101208a01614a89565b6015820180546001600160a01b0319166001600160a01b0392909216919091179055613c7d818484614283565b60405162ee0ec160e61b81526001600160a01b037f00000000000000000000000052fdd2102aaca3399f818f19353b84d1dd771b2d1690633b83b04090613cca9086908690600401615663565b600060405180830381600087803b158015613ce457600080fd5b505af1158015613cf8573d6000803e3d6000fd5b50505050613d0d81613d086142e6565b612402565b5050505050505050565b60008061065160017f3e2bfe19d7b287e1320a5adb4ca1cb62f90c30e328e073ba40443179d690e11a615a02565b613d4f82826116b9565b61090a5760405163e2517d3f60e01b81526001600160a01b038216600482015260248101839052604401610e56565b604080516001600160a01b0383166024808301919091528251808303909101815260449091019091526020810180516001600160e01b031663edad0a1360e01b17905260009060028401546040519192506001600160a01b0316906382646a5890613df09084906000906020016157c6565b6040516020818303038152906040526040518263ffffffff1660e01b81526004016116829190614d02565b60606000613e276122a0565b60098101805491925090806001600160401b03811115613e4957613e49614aa6565b604051908082528060200260200182016040528015613e72578160200160208202803683370190505b5093506000805b82811015613f34576000848281548110613e9557613e956152f2565b6000918252602090912001546001600160a01b03169050876001811115613ebe57613ebe615331565b6001600160a01b038216600090815260168801602052604090205460ff166001811115613eed57613eed615331565b03613f2b5780878481518110613f0557613f056152f2565b6001600160a01b039092166020928302919091019091015282613f2781615988565b9350505b50600101613e79565b50845250919392505050565b6000806000846001600160a01b031684604051613f5d9190615634565b6000604051808303816000865af19150503d8060008114613f9a576040519150601f19603f3d011682016040523d82523d6000602084013e613f9f565b606091505b5091509150818015613fbe575080511580613fbe5750613fbe81614375565b8015612a8e5750505050506001600160a01b03163b151590565b60607f7ac0d49c0df8eaba20287b89637474262377d7fcc1bf84c337d672a9ed663c75838360405160240161400e929190615a15565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152905092915050565b6000816001600160a01b031661405f858588614393565b6001600160a01b03161495945050505050565b6000614080858585856143d5565b60028701546040519192506001600160a01b0316906382646a58906140ac9084906000906020016157c6565b6040516020818303038152906040526040518263ffffffff1660e01b81526004016140d79190614d02565b600060405180830381600087803b1580156140f157600080fd5b505af1158015614105573d6000803e3d6000fd5b50505050505050505050565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663095ea7b360e01b179052614162848261444b565b610e9357604080516001600160a01b038516602482015260006044808301919091528251808303909101815260649091019091526020810180516001600160e01b031663095ea7b360e01b1790526141bb9085906144d4565b610e9384826144d4565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661420e57604051631afcd79f60e31b815260040160405180910390fd5b565b6142186141c5565b61422963d8a8b5c760e01b836127b1565b5061423b6378b4401360e11b846127b1565b50610e9363d3e319af60e01b826127b1565b6142556141c5565b61425d614537565b61426883838361453f565b61061b600033610f37565b61427b6141c5565b61420e614604565b60008190036142a55760405163e77c93b160e01b815260040160405180910390fd5b600a83016142b4828483615a7f565b507f7f63aacad63bc1693280450d5c3612ccd4efc53e46d69f3a537db102cd66290c828260405161273e929190615663565b6040516301753ab960e31b81524660048201526060907f00000000000000000000000052fdd2102aaca3399f818f19353b84d1dd771b2d6001600160a01b031690630ba9d5c890602401600060405180830381865afa15801561434d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611b959190810190615b3e565b6000815160201461438857506000919050565b506020015160011490565b604080516001600160a01b038086166020808401919091529085168284015282518083038401815260609092019092528051910120600090611a86908361460c565b60607f58cc5acb742fd5eb9df21407aed105abca7d37584645d014636aa13c7b7bc3878585858560405160240161440f9493929190615955565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b0319909316929092179091529050949350505050565b6000806000846001600160a01b0316846040516144689190615634565b6000604051808303816000865af19150503d80600081146144a5576040519150601f19603f3d011682016040523d82523d6000602084013e6144aa565b606091505b5091509150818015613fbe575080511580613fbe575080806020019051810190613fbe9190615364565b60006144e96001600160a01b03841683614645565b9050805160001415801561450e57508080602001905181019061450c9190615364565b155b1561061b57604051635274afe760e01b81526001600160a01b0384166004820152602401610e56565b61420e6141c5565b61455063024b274760e61b846127b1565b5061456263024b274760e61b836127b1565b5061457463024b274760e61b826127b1565b5061458663d93c394b60e01b846127b1565b5061459863d93c394b60e01b836127b1565b506145aa631bd8f8b560e11b846127b1565b506145bc631bd8f8b560e11b836127b1565b506145ce631bd8f8b560e11b826127b1565b506145e06362250a9560e11b846127b1565b506145f26362250a9560e11b826127b1565b50610e936362250a9560e11b836127b1565b61232a6141c5565b7f19457468657265756d205369676e6564204d6573736167653a0a3332000000006000908152601c839052603c8120611b689083614653565b6060611b688383600061467d565b6000806000806146638686614710565b925092509250614673828261475d565b5090949350505050565b6060814710156146a25760405163cd78605960e01b8152306004820152602401610e56565b600080856001600160a01b031684866040516146be9190615634565b60006040518083038185875af1925050503d80600081146146fb576040519150601f19603f3d011682016040523d82523d6000602084013e614700565b606091505b5091509150611b63868383614816565b6000806000835160410361474a5760208401516040850151606086015160001a61473c88828585614872565b955095509550505050614756565b50508151600091506002905b9250925092565b600082600381111561477157614771615331565b0361477a575050565b600182600381111561478e5761478e615331565b036147ac5760405163f645eedf60e01b815260040160405180910390fd5b60028260038111156147c0576147c0615331565b036147e15760405163fce698f760e01b815260048101829052602401610e56565b60038260038111156147f5576147f5615331565b0361090a576040516335e2f38360e21b815260048101829052602401610e56565b60608261482b5761482682614941565b611b68565b815115801561484257506001600160a01b0384163b155b1561486b57604051639996b31560e01b81526001600160a01b0385166004820152602401610e56565b5080611b68565b600080807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411156148ad5750600091506003905082614937565b604080516000808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015614901573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b03811661492d57506000925060019150829050614937565b9250600091508190505b9450945094915050565b8051156149515780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b508054600082559060005260206000209081019061076391905b808211156149985760008155600101614984565b5090565b6000602082840312156149ae57600080fd5b81356001600160e01b031981168114611b6857600080fd5b6000602082840312156149d857600080fd5b81356001600160401b038111156149ee57600080fd5b820160408185031215611b6857600080fd5b60008060208385031215614a1357600080fd5b82356001600160401b0380821115614a2a57600080fd5b818501915085601f830112614a3e57600080fd5b813581811115614a4d57600080fd5b866020606083028501011115614a6257600080fd5b60209290920196919550909350505050565b6001600160a01b038116811461076357600080fd5b600060208284031215614a9b57600080fd5b8135611b6881614a74565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b0381118282101715614ade57614ade614aa6565b60405290565b604051601f8201601f191681016001600160401b0381118282101715614b0c57614b0c614aa6565b604052919050565b60006001600160401b03821115614b2d57614b2d614aa6565b5060051b60200190565b6002811061076357600080fd5b60006020808385031215614b5757600080fd5b82356001600160401b03811115614b6d57600080fd5b8301601f81018513614b7e57600080fd5b8035614b91614b8c82614b14565b614ae4565b81815260069190911b82018301908381019087831115614bb057600080fd5b928401925b82841015614c065760408489031215614bce5760008081fd5b614bd6614abc565b8435614be181614a74565b815284860135614bf081614b37565b8187015282526040939093019290840190614bb5565b979650505050505050565b600060208284031215614c2357600080fd5b5035919050565b60008060408385031215614c3d57600080fd5b823591506020830135614c4f81614a74565b809150509250929050565b60008151808452602080850194506020840160005b83811015614c945781516001600160a01b031687529582019590820190600101614c6f565b509495945050505050565b602081526000611b686020830184614c5a565b60005b83811015614ccd578181015183820152602001614cb5565b50506000910152565b60008151808452614cee816020860160208601614cb2565b601f01601f19169290920160200192915050565b602081526000611b686020830184614cd6565b600082601f830112614d2657600080fd5b81356001600160401b03811115614d3f57614d3f614aa6565b614d52601f8201601f1916602001614ae4565b818152846020838601011115614d6757600080fd5b816020850160208301376000918101602001919091529392505050565b63ffffffff8116811461076357600080fd5b8035614da181614d84565b919050565b60008060008060808587031215614dbc57600080fd5b84356001600160401b0380821115614dd357600080fd5b818701915087601f830112614de757600080fd5b81356020614df7614b8c83614b14565b82815260069290921b8401810191818101908b841115614e1657600080fd5b948201945b83861015614e62576040868d031215614e345760008081fd5b614e3c614abc565b8635614e4781614a74565b81528684013584820152825260409095019490820190614e1b565b9850508801359550506040870135915080821115614e7f57600080fd5b50614e8c87828801614d15565b925050614e9b60608601614d96565b905092959194509250565b600060608284031215614eb857600080fd5b50919050565b60008083601f840112614ed057600080fd5b5081356001600160401b03811115614ee757600080fd5b60208301915083602082850101111561364557600080fd5b600080600060408486031215614f1457600080fd5b83356001600160401b0380821115614f2b57600080fd5b9085019060608288031215614f3f57600080fd5b604051606081018181108382111715614f5a57614f5a614aa6565b604052823582811115614f6c57600080fd5b614f7889828601614d15565b8252506020830135602082015260408301356040820152809550506020860135915080821115614fa757600080fd5b50614fb486828701614ebe565b9497909650939450505050565b60008060208385031215614fd457600080fd5b82356001600160401b03811115614fea57600080fd5b614ff685828601614ebe565b90969095509350505050565b803565ffffffffffff81168114614da157600080fd5b60008060006040848603121561502d57600080fd5b83356001600160401b038082111561504457600080fd5b908501906040828803121561505857600080fd5b615060614abc565b61506983615002565b815260208301358281111561507d57600080fd5b61508989828601614d15565b602083015250809550506020860135915080821115614fa757600080fd5b6000602082840312156150b957600080fd5b81356001600160401b038111156150cf57600080fd5b82016101008185031215611b6857600080fd5b801515811461076357600080fd5b60006020828403121561510257600080fd5b8135611b68816150e2565b60008060006040848603121561512257600080fd5b833561512d81614a74565b925060208401356001600160401b038082111561514957600080fd5b818601915086601f83011261515d57600080fd5b81358181111561516c57600080fd5b8760208260051b850101111561518157600080fd5b6020830194508093505050509250925092565b6000806000606084860312156151a957600080fd5b83356151b481614d84565b925060208401356151c481614d84565b929592945050506040919091013590565b6000602082840312156151e757600080fd5b81356001600160401b038111156151fd57600080fd5b82016101408185031215611b6857600080fd5b60006020828403121561522257600080fd5b8151611b6881614a74565b60006020828403121561523f57600080fd5b611b6882615002565b6000808335601e1984360301811261525f57600080fd5b8301803591506001600160401b0382111561527957600080fd5b60200191503681900382131561364557600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6001600160a01b0386811682528516602082015265ffffffffffff84166040820152608060608201819052600090614c06908301848661528e565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b8082018082111561065157610651615308565b634e487b7160e01b600052602160045260246000fd5b60006020828403121561535957600080fd5b8151611b6881614b37565b60006020828403121561537657600080fd5b8151611b68816150e2565b600060018060a01b038086168352606060208401526153a36060840186614c5a565b9150808416604084015250949350505050565b600060208083850312156153c957600080fd5b82516001600160401b038111156153df57600080fd5b8301601f810185136153f057600080fd5b80516153fe614b8c82614b14565b81815260059190911b8201830190838101908783111561541d57600080fd5b928401925b82841015614c0657835161543581614a74565b82529284019290840190615422565b600181811c9082168061545857607f821691505b602082108103614eb857634e487b7160e01b600052602260045260246000fd5b6000806000806080858703121561548e57600080fd5b845161549981614d84565b60208601519094506154aa81614d84565b6040860151606090960151949790965092505050565b60008151808452602080850194506020840160005b83811015614c9457815180516001600160a01b031688528301516001600160601b031683880152604090960195908201906001016154d5565b60008151808452602080850194506020840160005b83811015614c9457815180516001600160a01b031688528301518388015260409096019590820190600101615523565b6001600160a01b038381168252604060208084018290528451848301819052600093606092909183870190600581901b88018501898401885b8381101561562357605f198b8403018552815160c081518186526155b2828701826154c0565b915050888883015116888601528a8201518582038c8701526155d4828261550e565b838c015163ffffffff908116888e01526080808601519091169088015260a09384015187820394880194909452915061560f90508183614cd6565b96880196945050509085019060010161558c565b50909b9a5050505050505050505050565b60008251615646818460208701614cb2565b9190910192915050565b602081526000611b68602083018461550e565b602081526000611a8660208301848661528e565b600060208083018184528085518083526040925060408601915060408160051b87010184880160005b8381101561571557603f19898403018552815160a081518186526156c6828701826154c0565b838b01516001600160a01b0316878c0152898401518a88015260608085015163ffffffff90811691890191909152608094850151169390960192909252505093860193908601906001016156a0565b509098975050505050505050565b60006020828403121561573557600080fd5b5051919050565b6020808252825482820181905260008481528281209092916040850190845b818110156157805783546001600160a01b03168352600193840193928501920161575b565b50909695505050505050565b60008161579b5761579b615308565b506000190190565b6001600160e01b03199290921682526001600160a01b0316602082015260400190565b6040815260006157d96040830185614cd6565b905063ffffffff831660208301529392505050565b60018060a01b038316815260406020820152600082516060604084015261581860a0840182614cd6565b90506020840151606084015260408401516080840152809150509392505050565b6001600160a01b0385811682528416602082015265ffffffffffff83166040820152608060608201819052600090612a8b90830184614cd6565b610100810160408683376001600160a01b038581166040840152841660608301526080838184013795945050505050565b80516001600160a01b03168252602080820151908301526040810151600281106158de57634e487b7160e01b600052602160045260246000fd5b806040840152505050565b6001600160a01b0384811682526060602080840182905285518483018190526000938783019290916080870190865b8181101561593b5761592b8387516158a4565b9483019491860191600101615918565b505080955050508086166040860152505050949350505050565b6001600160a01b0385811682526020820185905260e0820190608085604085013780841660c08401525095945050505050565b60006001820161599a5761599a615308565b5060010190565b60808181019083833792915050565b6001600160a01b0384811682526060602080840182905285518483018190526000938783019290916080870190865b8181101561593b576159f28387516158a4565b94830194918601916001016159df565b8181038181111561065157610651615308565b604081526000615a286040830185614cd6565b90508260208301529392505050565b601f82111561061b576000816000526020600020601f850160051c81016020861015615a605750805b601f850160051c820191505b8181101561229857828155600101615a6c565b6001600160401b03831115615a9657615a96614aa6565b615aaa83615aa48354615444565b83615a37565b6000601f841160018114615ade5760008515615ac65750838201355b600019600387901b1c1916600186901b17835561097b565b600083815260209020601f19861690835b82811015615b0f5786850135825560209485019460019092019101615aef565b5086821015615b2c5760001960f88860031b161c19848701351681555b505060018560011b0183555050505050565b60006020808385031215615b5157600080fd5b82516001600160401b03811115615b6757600080fd5b8301601f81018513615b7857600080fd5b8051615b86614b8c82614b14565b81815260069190911b82018301908381019087831115615ba557600080fd5b928401925b82841015614c065760408489031215615bc35760008081fd5b615bcb614abc565b8451615bd681614a74565b815284860151615be581614b37565b8187015282526040939093019290840190615baa56fefe6065fb4e9872e2ad4479001655335380d83f70e163706cd65857449b84510002dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0092c9d1c07198b9465c1d2907806f1150774e0758b4edc0d272f5df87a5a7492ca264697066735822122081807e166ec3a5caf99d1fdaaaac9caf9b33b2b6fe22300d699a895e5f6912da64736f6c63430008190033",
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

	CreateEigenRewardsSubmission(opts *bind.TransactOpts, _startTimestamp uint32, _duration uint32, _amount *big.Int) (*types.Transaction, error)

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

// CreateEigenRewardsSubmission is a paid mutator transaction binding the contract method 0xeab90f1c.
//
// Solidity: function createEigenRewardsSubmission(uint32 _startTimestamp, uint32 _duration, uint256 _amount) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) CreateEigenRewardsSubmission(opts *bind.TransactOpts, _startTimestamp uint32, _duration uint32, _amount *big.Int) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "createEigenRewardsSubmission", _startTimestamp, _duration, _amount)
}

// CreateEigenRewardsSubmission is a paid mutator transaction binding the contract method 0xeab90f1c.
//
// Solidity: function createEigenRewardsSubmission(uint32 _startTimestamp, uint32 _duration, uint256 _amount) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) CreateEigenRewardsSubmission(_startTimestamp uint32, _duration uint32, _amount *big.Int) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.CreateEigenRewardsSubmission(&_ContractAvsGovernance.TransactOpts, _startTimestamp, _duration, _amount)
}

// CreateEigenRewardsSubmission is a paid mutator transaction binding the contract method 0xeab90f1c.
//
// Solidity: function createEigenRewardsSubmission(uint32 _startTimestamp, uint32 _duration, uint256 _amount) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) CreateEigenRewardsSubmission(_startTimestamp uint32, _duration uint32, _amount *big.Int) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.CreateEigenRewardsSubmission(&_ContractAvsGovernance.TransactOpts, _startTimestamp, _duration, _amount)
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
