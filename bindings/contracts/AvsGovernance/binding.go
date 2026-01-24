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
	Vault                      common.Address
	AvsDirectoryContract       common.Address
	AllowlistSigner            common.Address
	AvsName                    string
	BlsAuthSingleton           common.Address
}

// IAvsGovernancePaymentRequestMessage is an auto generated low-level Go binding around an user-defined struct.
type IAvsGovernancePaymentRequestMessage struct {
	Operator   common.Address
	FeeToClaim *big.Int
}

// IAvsGovernanceStrategyMultiplier is an auto generated low-level Go binding around an user-defined struct.
type IAvsGovernanceStrategyMultiplier struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// ContractAvsGovernanceMetaData contains all meta data concerning the ContractAvsGovernance contract.
var ContractAvsGovernanceMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"avsName\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"completeRewardsReceiverModification\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositERC20\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDefaultStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getIsAllowlisted\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNumOfOperatorsLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"numOfOperatorsLimitView\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRewardsReceiver\",\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_initializationParams\",\"type\":\"tuple\",\"internalType\":\"structIAvsGovernance.InitializationParams\",\"components\":[{\"name\":\"avsGovernanceMultisigOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operationsMultisig\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"communityMultisig\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"othenticRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"messageHandler\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avsDirectoryContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistSigner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avsName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"blsAuthSingleton\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isFlowPaused\",\"inputs\":[{\"name\":\"_pausableFlow\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"_isPaused\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxEffectiveBalance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minSharesForStrategy\",\"inputs\":[{\"name\":\"_strategy\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minVotingPower\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"numOfActiveOperators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"numOfOperators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"numOfShares\",\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"_pausableFlow\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"queueRewardsReceiverModification\",\"inputs\":[{\"name\":\"_newRewardsReceiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerAsAllowedOperator\",\"inputs\":[{\"name\":\"_blsKey\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"},{\"name\":\"_authToken\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_rewardsReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"_blsRegistrationSignature\",\"type\":\"tuple\",\"internalType\":\"structBLSAuthLibrary.Signature\",\"components\":[{\"name\":\"signature\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerAsOperator\",\"inputs\":[{\"name\":\"_blsKey\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"},{\"name\":\"_rewardsReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"_blsRegistrationSignature\",\"type\":\"tuple\",\"internalType\":\"structBLSAuthLibrary.Signature\",\"components\":[{\"name\":\"signature\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAllowlistSigner\",\"inputs\":[{\"name\":\"_allowlistSigner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAvsGovernanceLogic\",\"inputs\":[{\"name\":\"_avsGovernanceLogic\",\"type\":\"address\",\"internalType\":\"contractIAvsGovernanceLogic\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAvsGovernanceMultiplierSyncer\",\"inputs\":[{\"name\":\"_newAvsGovernanceMultiplierSyncer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAvsName\",\"inputs\":[{\"name\":\"_avsName\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBLSAuthSingleton\",\"inputs\":[{\"name\":\"_blsAuthSingleton\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setIsAllowlisted\",\"inputs\":[{\"name\":\"_isAllowlisted\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxEffectiveBalance\",\"inputs\":[{\"name\":\"_maxBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinSharesForStrategy\",\"inputs\":[{\"name\":\"_strategy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_minShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinVotingPower\",\"inputs\":[{\"name\":\"_minVotingPower\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setNumOfOperatorsLimit\",\"inputs\":[{\"name\":\"_newLimitOfNumOfOperators\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOthenticRegistry\",\"inputs\":[{\"name\":\"_othenticRegistry\",\"type\":\"address\",\"internalType\":\"contractIOthenticRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRewardsReceiverModificationDelay\",\"inputs\":[{\"name\":\"_rewardsReceiverModificationDelay\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStrategyMultiplier\",\"inputs\":[{\"name\":\"_strategyMultiplier\",\"type\":\"tuple\",\"internalType\":\"structIAvsGovernance.StrategyMultiplier\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"multiplier\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStrategyMultiplierBatch\",\"inputs\":[{\"name\":\"_strategyMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIAvsGovernance.StrategyMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"multiplier\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSupportedStrategies\",\"inputs\":[{\"name\":\"_strategies\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"strategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyMultiplier\",\"inputs\":[{\"name\":\"_strategy\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferAvsGovernanceMultisig\",\"inputs\":[{\"name\":\"_newAvsGovernanceMultisig\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferMessageHandler\",\"inputs\":[{\"name\":\"_newMessageHandler\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"_pausableFlow\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unregisterAsOperator\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"vault\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"votingPower\",\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawBatchRewards\",\"inputs\":[{\"name\":\"_operators\",\"type\":\"tuple[]\",\"internalType\":\"structIAvsGovernance.PaymentRequestMessage[]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeToClaim\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"_lastPayedTask\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawRewards\",\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_lastPayedTask\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_feeToClaim\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BLSAuthSingletonSet\",\"inputs\":[{\"name\":\"blsAuthSingleton\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FlowPaused\",\"inputs\":[{\"name\":\"_pausableFlow\",\"type\":\"bytes4\",\"indexed\":false,\"internalType\":\"bytes4\"},{\"name\":\"_pauser\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FlowUnpaused\",\"inputs\":[{\"name\":\"_pausableFlowFlag\",\"type\":\"bytes4\",\"indexed\":false,\"internalType\":\"bytes4\"},{\"name\":\"_unpauser\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxEffectiveBalanceSet\",\"inputs\":[{\"name\":\"maxEffectiveBalance\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinSharesPerStrategySet\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"minShares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinVotingPowerSet\",\"inputs\":[{\"name\":\"minVotingPower\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"blsKey\",\"type\":\"uint256[4]\",\"indexed\":false,\"internalType\":\"uint256[4]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorUnregistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"QueuedRewardsReceiverModification\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"delay\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetAllowlistSigner\",\"inputs\":[{\"name\":\"allowlistSigner\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetAvsGovernanceLogic\",\"inputs\":[{\"name\":\"avsGovernanceLogic\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetAvsGovernanceMultiplierSyncer\",\"inputs\":[{\"name\":\"avsGovernanceMultiplierSyncer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetAvsGovernanceMultisig\",\"inputs\":[{\"name\":\"newAvsGovernanceMultisig\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetAvsName\",\"inputs\":[{\"name\":\"avsName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetIsAllowlisted\",\"inputs\":[{\"name\":\"isAllowlisted\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetMessageHandler\",\"inputs\":[{\"name\":\"newMessageHandler\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetNumOfOperatorsLimit\",\"inputs\":[{\"name\":\"newLimitOfNumOfOperators\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetOthenticRegistry\",\"inputs\":[{\"name\":\"othenticRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetRewardsReceiver\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetRewardsReceiverModificationDelay\",\"inputs\":[{\"name\":\"modificationDelay\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetStrategyMultiplier\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"multiplier\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetSupportedStrategies\",\"inputs\":[{\"name\":\"strategies\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlInvalidMultiplierSyncer\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AllowlistDisabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AllowlistEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"FlowIsCurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FlowIsCurrentlyUnpaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAllowlistAuthToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBlsRegistrationSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMultiplierNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRewardsReceiver\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSlashingRate\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidStrategy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ModificationDelayNotPassed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnoughVotingPower\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NumOfActiveOperatorsIsGreaterThanNumOfOperatorLimit\",\"inputs\":[{\"name\":\"numOfOperatorsLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"numOfActiveOperators\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NumOfOperatorsLimitReached\",\"inputs\":[{\"name\":\"numOfOperatorsLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OperatorAlreadyRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PauseFlowIsAlreadyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"UnpausingFlowIsAlreadyUnpaused\",\"inputs\":[]}]",
	Bin: "0x60808060405234601557614a93908161001a8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f905f3560e01c90816301ffc9a7146130a157508063076400d5146130025780630986944214612d825780631b21ba7214612bf757806322609a4d146128e9578063248a9ca3146128b15780632f2ff15d14612866578063305df58a1461274d578063312c150b146126705780633256b4d1146124c257806333cfb7b7146124165780633425e8d81461234a57806336568abe1461230557806336fffde0146122da5780633aa83ec7146121e357806341b92a29146120c257806345a022fa146120475780634d07f65114611fa15780634ef1476e14611f27578063513c52ba14611eba57806355e4891814611e605780635e95cee214611dfe5780636a90780314611d385780636ade02da14611c495780636b1906f814611cf05780636b3aa72e14611cba57806376086c7014611c4e5780637897dec314611c495780637d38e92614611bf65780638987c76714611b445780638a70469a14611ad75780638f53bc5014611a7e57806391d1485414611a2857806393304a9d146113565780639d79e4a7146112cf5780639e965cc114611255578063a217fddf14611239578063a88171ee146111fb578063a98fb35514611178578063b525fa8814611145578063b79092fd146110e1578063bac1e94b14610fe9578063bc8be0c814610d44578063c07473f614610cfb578063c3814e5b14610ca2578063d547741f14610c4e578063d94a2e1d14610b4c578063d9f9027f14610aba578063e474def414610abf578063e481af9d14610aba578063e6474b0f14610979578063e86685d91461093e578063efd96978146108f1578063f251c9a6146108c6578063fab57b8f146102be5763fbfa77cf14610286575f80fd5b346102bb57806003193601126102bb57505f516020614a1e5f395f51905f52546040516001600160a01b039091168152602090f35b80fd5b50346102bb5760203660031901126102bb576004356001600160401b0381116108c2573681900360048201610140600319830112610750575f5160206149be5f395f51905f52549260ff8460401c1615936001600160401b038116801590816108ba575b60011490816108b0575b1590816108a7575b506108985767ffffffffffffffff1981166001175f5160206149be5f395f51905f52558461086c575b50610366614642565b6001600160a01b03610377836136e2565b1615158061084f575b80610832575b80610815575b806107f8575b806107db575b806107be575b806107a1575b1561075c576103b385926136e2565b906103c0602482016136e2565b6103cc604483016136e2565b6103d8608484016136e2565b956101048401359060221901811215610758578301916004830135926001600160401b038411610754576024019383360385136107545761064a92610124926104c1916104bb90816001600160a01b03610434606488016136e2565b169a61043e614642565b610447856139aa565b5061045181613930565b5061045b83613a4a565b50610464614642565b61046c614642565b61047581613aea565b5061047f85613aea565b5061048983613aea565b5061049381613b8a565b5061049d85613b8a565b506104a781613c2a565b506104b185613c2a565b506104bb83613c2a565b50613cca565b506104cb3361411a565b506104d4614642565b6104dc614642565b60015f51602061497e5f395f51905f52555f5160206148be5f395f51905f5280546001600160a01b038881166001600160a01b0319928316179092555f51602061487e5f395f51905f528054928b16929091169190911790555f51602061489e5f395f51905f529761054d906138b6565b506001600160a01b0361056260a483016136e2565b5f516020614a1e5f395f51905f528054919092166001600160a01b03166001600160a01b031990911617905561059a60e482016136e2565b7f3e2bfe19d7b287e1320a5adb4ca1cb62f90c30e328e073ba40443179d690e11f80546001600160a01b0319166001600160a01b0392831617905562093a807f3e2bfe19d7b287e1320a5adb4ca1cb62f90c30e328e073ba40443179d690e1245561060760c483016136e2565b5f5160206149de5f395f51905f528054919092166001600160a01b03166001600160a01b031990911617905560645f5160206148de5f395f51905f5255016136e2565b5f51602061481e5f395f51905f5280546001600160a01b0319166001600160a01b03909216919091179055610680818387614238565b823b15610750576106a99284928360405180968195829462ee0ec160e61b8452600484016134cf565b03925af180156107455761072c575b5050806106c76106cd926145f1565b90613e11565b6106d45780f35b68ff0000000000000000195f5160206149be5f395f51905f5254165f5160206149be5f395f51905f52557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a180f35b816107369161317e565b61074157825f6106b8565b8280fd5b6040513d84823e3d90fd5b8380fd5b8680fd5b8580fd5b60405162461bcd60e51b815260206004820152601c60248201527f417673476f7665726e616e63653a20496e76616c696420696e707574000000006044820152606490fd5b506001600160a01b036107b660e483016136e2565b1615156103a4565b506001600160a01b036107d360c483016136e2565b16151561039e565b506001600160a01b036107f060a483016136e2565b161515610398565b506001600160a01b0361080d608483016136e2565b161515610392565b506001600160a01b0361082a606483016136e2565b16151561038c565b506001600160a01b03610847604483016136e2565b161515610386565b506001600160a01b03610864602483016136e2565b161515610380565b68ffffffffffffffffff191668010000000000000001175f5160206149be5f395f51905f52555f61035d565b63f92ee8a960e01b8652600486fd5b9050155f610334565b303b15915061032c565b869150610322565b5080fd5b50346102bb57806003193601126102bb575060205f5160206148de5f395f51905f5254604051908152f35b50346102bb5760203660031901126102bb5760043563ffffffff60e01b81168091036108c25760408260ff92602094525f51602061491e5f395f51905f5284522054166040519015158152f35b50346102bb57806003193601126102bb57506109756109695f51602061489e5f395f51905f526145f1565b60405191829182613255565b0390f35b50346102bb57806003193601126102bb57335f9081525f5160206147be5f395f51905f5260205260409020545f51602061489e5f395f51905f529015610aab57335f9081525f516020614a3e5f395f51905f5260205260409020600101544210610a9c57335f9081525f516020614a3e5f395f51905f526020526040902054610a0d916001600160a01b0390911690613783565b335f9081527f3e2bfe19d7b287e1320a5adb4ca1cb62f90c30e328e073ba40443179d690e125602052604090819020805460ff1916905551610a4e81613134565b8181526020808201838152335f9081525f516020614a3e5f395f51905f529092526040909120915182546001600160a01b0319166001600160a01b0391909116178255516001919091015580f35b638ce7a3f160e01b8252600482fd5b6325ec6c1f60e01b8252600482fd5b613328565b50346102bb5760203660031901126102bb577ffa4acc0aaeb2714e420e9c8339167ddef7bc66c0f94a0c5a7722de21dcb7508c6020610afc6130f4565b610b04613644565b7f3e2bfe19d7b287e1320a5adb4ca1cb62f90c30e328e073ba40443179d690e11f80546001600160a01b0319166001600160a01b03929092169182179055604051908152a180f35b50346102bb5760203660031901126102bb576004356001600160401b0381116108c257366023820112156108c25780600401356001600160401b038111610741573660248260061b840101116107415790610ba5613575565b7f3e2bfe19d7b287e1320a5adb4ca1cb62f90c30e328e073ba40443179d690e12c90835b83811015610c4a576001907f8ae53ffd0ebc018acb19342fba690554d49ae9a467a9606a38b49cb5ad775c818160061b8401610c2960446024830192013591610c11816136e2565b868060a01b03165f52876020528260405f20556136e2565b604080516001600160a01b039290921682526020820192909252a101610bc9565b8480f35b50346102bb5760403660031901126102bb57610c9e600435610c6e61310a565b90610c99610c94825f525f51602061493e5f395f51905f52602052600160405f20015490565b61369c565b6141af565b5080f35b50346102bb5760203660031901126102bb57610cbc6130f4565b905060018060a01b03165f527f3e2bfe19d7b287e1320a5adb4ca1cb62f90c30e328e073ba40443179d690e12b602052602060405f2054604051908152f35b50346102bb5760203660031901126102bb57610d3c602091610d1b6130f4565b5f51602061489e5f395f51905f529150610d3482614382565b9290506144fe565b604051908152f35b50346102bb5760403660031901126102bb576004356001600160401b0381116108c257366023820112156108c257806004013590610d8182613311565b90610d8f604051928361317e565b82825260208201906024829460061b8201019036821161075857602401915b818310610fb15750505060243591610dc46135ff565b5f516020614a1e5f395f51905f52546001600160a01b031690610de56134f6565b507f3e2bfe19d7b287e1320a5adb4ca1cb62f90c30e328e073ba40443179d690e12693855b8451811015610fa657610e1d818661350e565b5180516001600160a01b0316908115610e82576001600160a01b038083165f90815260208a815260409091205492015160019493610e619389938892911690613eec565b15610e6d575b01610e0a565b876020610e7a838961350e565b510152610e67565b50505091509392505b60405191604083019060208085015251809152606083019490845b818110610f74578580610eee610f028888610eca828e03601f19810184528361317e565b604051938492631eb0352760e21b6020850152604060248501526064840190613297565b90604483015203601f19810183528261317e565b5f51602061487e5f395f51905f52546001600160a01b0316803b15610f705760405163104c8d4b60e31b8152602060048201529183918391829084908290610f4e906024830190613297565b03925af1801561074557610f5f5750f35b81610f699161317e565b6102bb5780f35b5050fd5b9091956020604082610f9b6001948b516020809160018060a01b0381511684520151910152565b019701929101610ea6565b509150939250610e8b565b6040833603126107585760206040918251610fcb81613134565b610fd486613120565b81528286013583820152815201920191610dae565b50346102bb5760203660031901126102bb576004356001600160e01b03198116808203610741578083525f51602061491e5f395f51905f5260205260ff604084205416156110d25761103a8161369c565b8083525f51602061491e5f395f51905f5260205260ff604084205416156110c35782525f51602061491e5f395f51905f5260209081526040808420805460ff1916905580516001600160e01b0319909316835233918301919091527fc7e56e17b0a6c4b467df6495e1eda1baecd7ba20604e80c1058ac06f4578d85e9190819081015b0390a180f35b635bfd2da760e11b8352600483fd5b6368c87f3360e11b8352600483fd5b50346102bb5760203660031901126102bb575f516020614a1e5f395f51905f525481906001600160a01b0316803b1561114257604051634bff5c9360e11b815233600480830191909152356024820152908290829081838160448101610f4e565b50fd5b50346102bb57806003193601126102bb5750602060ff5f5160206149de5f395f51905f525460a01c166040519015158152f35b50346102bb5760203660031901126102bb57806004356001600160401b038111611142576111aa9036906004016132e4565b906111b36135ba565b5f5160206149de5f395f51905f52546001600160a01b031691823b156111f657610f4e9284928360405180968195829463a98fb35560e01b8452600484016134cf565b505050fd5b50346102bb57806003193601126102bb575060207f3e2bfe19d7b287e1320a5adb4ca1cb62f90c30e328e073ba40443179d690e12a54604051908152f35b50346102bb57806003193601126102bb57602090604051908152f35b50346102bb5760203660031901126102bb576004358015158091036108c25760207f2dcb3282f9b7aa18e1bf7fa254c45f3e270e8f26d9a37ae590d5d8125b58d1b1916112a06135ba565b5f5160206149de5f395f51905f52805460ff60a01b191660a083901b60ff60a01b16179055604051908152a180f35b50346102bb5760203660031901126102bb576004356112ec6135ba565b5f51602061499e5f395f51905f525481811161134157506020817fc0dd1d82df4ae12576f7a7912395305cf73deae26c764dd74a945cd6ba81591b925f5160206148de5f395f51905f5255604051908152a180f35b63d2930ec560e01b8352600452602452604490fd5b50346102bb576101203660031901126102bb57366084116102bb576084356001600160401b0381116108c2576113909036906004016132e4565b60a4359291906001600160a01b0384168085036107505760c4356001600160401b03811161180b576113c69036906004016131e4565b9260403660e319011261180b57335f9081525f5160206147be5f395f51905f5260205260409020545f51602061489e5f395f51905f529390611a195761140a6136f6565b61141261374b565b60ff5f5160206149de5f395f51905f525460a01c1615611a0a579061143891369161319f565b7f3e2bfe19d7b287e1320a5adb4ca1cb62f90c30e328e073ba40443179d690e11f5460408051306020820190815233928201929092526001600160a01b03909216926114d5926114cc929061149a81606081015b03601f19810183528261317e565b5190207f19457468657265756d205369676e6564204d6573736167653a0a3332000000008952601c52603c882061466d565b909291926146a7565b6001600160a01b0316036119fb57156119ed576114f28482613783565b6040516114fe81613163565b3661010312156107505760405161151481613134565b8036610124116107585760e4905b61012482106119dd57505081525f5160206148de5f395f51905f5254805f51602061499e5f395f51905f525410156119cb57505f51602061481e5f395f51905f5254604051633cf65e3b60e01b815291516001600160a01b0390911690829086600483015b600282106119b157505050610104816020933360448301523060648301526080600460848401375afa9081156119a6578491611987575b5015611978575f5160206148fe5f395f51905f52546001600160a01b0316801515949091906115f86115ef83614382565b909333906144fe565b915f5160206147de5f395f51905f525483101590816118a9575b501561189a5785611843575b6040516358cc5acb60e01b60208201526116438161148c846004873360248601613805565b5f51602061487e5f395f51905f52546001600160a01b031690813b156107545761168e8792839260405194858094819363104c8d4b60e31b8352602060048401526024830190613297565b03925af1801561183857908691611823575b50505f51602061499e5f395f51905f5254955f19871461180f5760018697015f51602061499e5f395f51905f525560018060a01b0333165f525f5160206147be5f395f51905f52602052600160405f20556117a6575b5050506040516080600482377f54bc9cf83c2eb0f2ad1abf6e4fab882964404622ba2df6b5a9356a18d3aac05560803392a25f5160206149de5f395f51905f52546001600160a01b0316803b15610f7057604051639926ee7d60e01b8152918391839182908490829061176d903360048401613351565b03925af1801561074557611791575b5060015f51602061497e5f395f51905f525580f35b8161179b9161317e565b6102bb57805f61177c565b823b1561180b576117d3926004868094604051968795869485936376c56c1b60e11b855233828601613805565b03925af19081156118005783916117eb575b806116f6565b816117f59161317e565b61114257815f6117e5565b6040513d85823e3d90fd5b8480fd5b634e487b7160e01b86526011600452602486fd5b8161182d9161317e565b61180b57845f6116a0565b6040513d88823e3d90fd5b823b1561180b5760405163094e7a3f60e41b815285818061186a8560048833828601613805565b038183885af1801561183857908691611885575b505061161e565b8161188f9161317e565b61180b57845f61187e565b631c33ce8d60e11b8552600485fd5b905060018060a01b035f5160206148be5f395f51905f52541660405180926330959fcb60e11b8252604482013360048401526040602484015281518091526020606484019201908a5b81811061194a5750505091818060209403915afa90811561183857869161191b575b505f611612565b61193d915060203d602011611943575b611935818361317e565b8101906137ed565b5f611914565b503d61192b565b825180516001600160a01b0316855260209081015181860152879550604090940193909201916001016118f2565b6314532dfd60e11b8352600483fd5b6119a0915060203d60201161194357611935818361317e565b5f6115be565b6040513d86823e3d90fd5b829350602080916001939451815201930191018492611587565b63c77b407760e01b8552600452602484fd5b8135815260209182019101611522565b62cc6ac760e01b8352600483fd5b63d2342ec760e01b8452600484fd5b632d35c8d360e01b8652600486fd5b6342ee68b560e01b8652600486fd5b50346102bb5760403660031901126102bb576040611a4461310a565b9160043581525f51602061493e5f395f51905f52602052209060018060a01b03165f52602052602060ff60405f2054166040519015158152f35b50346102bb5760203660031901126102bb57611a986130f4565b905060018060a01b03165f527f3e2bfe19d7b287e1320a5adb4ca1cb62f90c30e328e073ba40443179d690e12c602052602060405f2054604051908152f35b50346102bb5760203660031901126102bb577f47c8c3268759fc47868c5e319217a2e85d47bd3935a4108debe246f6025fb88b6020600435611b176135ba565b807f3e2bfe19d7b287e1320a5adb4ca1cb62f90c30e328e073ba40443179d690e12455604051908152a180f35b50346102bb5760203660031901126102bb576004356001600160a01b038116908190036108c257611b736135ba565b6362250a9560e11b82525f51602061491e5f395f51905f52602052604082205460ff16611be7575f5160206148fe5f395f51905f5280546001600160a01b031916821790556040519081527f7c36ee80df183e227956a9f387a48d26bbf4d2f1526410493d11126de5a8942c90602090a180f35b63722fdba960e01b8252600482fd5b50346102bb5760203660031901126102bb576004356001600160401b0381116108c257611c2a611c469136906004016132e4565b90611c336135ba565b5f51602061489e5f395f51905f52614238565b80f35b6132bb565b50346102bb5760203660031901126102bb577ec6fb6db9c52d89a1eaf84e0470a3304db2086d0ac44d64ebf4ea35a905a7d06020600435611c8d6135ba565b807f3e2bfe19d7b287e1320a5adb4ca1cb62f90c30e328e073ba40443179d690e12a55604051908152a180f35b50346102bb57806003193601126102bb57505f5160206149de5f395f51905f52546040516001600160a01b039091168152602090f35b50346102bb5760203660031901126102bb57611d0a6130f4565b905060018060a01b03165f525f5160206147be5f395f51905f52602052602060405f20541515604051908152f35b50346102bb5760203660031901126102bb57611d9f9060205f51602061485e5f395f51905f52611d666130f4565b5f5160206148be5f395f51905f5254604051639004134760e01b81529586936001600160a01b0390921692849283929160048401613442565b03915afa908115611df25790611dbb575b602090604051908152f35b506020813d602011611dea575b81611dd56020938361317e565b81010312611de65760209051611db0565b5f80fd5b3d9150611dc8565b604051903d90823e3d90fd5b50346102bb5760203660031901126102bb57611e186130f4565b905060018060a01b03165f527f3e2bfe19d7b287e1320a5adb4ca1cb62f90c30e328e073ba40443179d690e126602052602060405f2060018060a01b03905416604051908152f35b50346102bb5760203660031901126102bb577f10203ddc048c86cf14172a6ea2565c805ce7320b22d6941b2eb396d0ee0779836020600435611ea06135ba565b805f5160206147de5f395f51905f5255604051908152a180f35b50346102bb5760203660031901126102bb577f024e98b7d808a3ddb028252dc95dfdcb165a0ca59fcff8984b4fecf9a72226496020611ef76130f4565b611eff6135ba565b611f08336140a2565b50611f1281613930565b506040516001600160a01b039091168152a180f35b50346102bb5760203660031901126102bb577f4cbffdecf3b5e4b22bfb2bdec99a66f8fcf81e19b060682afd9645c729da14726020611f646130f4565b611f6c613644565b5f51602061481e5f395f51905f5280546001600160a01b0319166001600160a01b03929092169182179055604051908152a180f35b50346102bb5760203660031901126102bb577f997f84b541d7b68e210e6f50e3402b51d8411dbbc4d44ed81e508383126e4e946020611fde6130f4565b611fe6613644565b5f51602061487e5f395f51905f5254612007906001600160a01b031661402a565b50612011816138b6565b505f51602061487e5f395f51905f5280546001600160a01b0319166001600160a01b03929092169182179055604051908152a180f35b50346102bb5760203660031901126102bb576004356001600160a01b038116908190036108c25760207ff9855cc914fefc396bdeb5a4dcb97a2f6c75f4d6f00a8e71d6f9a40e474afe8d9161209a613644565b5f5160206148be5f395f51905f5280546001600160a01b03191682179055604051908152a180f35b50346102bb57806003193601126102bb5760405190807f3e2bfe19d7b287e1320a5adb4ca1cb62f90c30e328e073ba40443179d690e123549061210482613497565b80855291600181169081156121bc5750600114612140575b6109758461212c8186038261317e565b604051918291602083526020830190613297565b7f3e2bfe19d7b287e1320a5adb4ca1cb62f90c30e328e073ba40443179d690e12381527f426338f094e72a54b4785bcef70a30e76c5fb1ceb50028dfc7385f39cbf28add939250905b8082106121a25750909150810160200161212c8261211c565b919260018160209254838588010152019101909291612189565b60ff191660208087019190915292151560051b8501909201925061212c915083905061211c565b50346102bb5760203660031901126102bb576004356001600160e01b03198116808203610741578083525f51602061491e5f395f51905f5260205260ff6040842054166122cb576122338161369c565b8083525f51602061491e5f395f51905f5260205260ff6040842054166122bc5782525f51602061491e5f395f51905f5260209081526040808420805460ff1916600117905580516001600160e01b0319909316835233918301919091527f95c3658c5e0c74e20cf12db371b9b67d26e97a1937f6d2284f88cc44d036b4f69190819081016110bd565b63dfe10d7d60e01b8352600483fd5b63722fdba960e01b8352600483fd5b50346102bb57806003193601126102bb575060205f5160206147de5f395f51905f5254604051908152f35b50346102bb5760403660031901126102bb5761231f61310a565b336001600160a01b0382160361233b57610c9e906004356141af565b63334bd91960e11b8252600482fd5b50346102bb5760203660031901126102bb577fb73a70f24733a9265231de5807eae76d1740a9974b31a142ef9e243508987bbe60206123876130f4565b61238f6135ba565b7f3e2bfe19d7b287e1320a5adb4ca1cb62f90c30e328e073ba40443179d690e12d546123c3906001600160a01b0316613fb2565b506123cd81613837565b507f3e2bfe19d7b287e1320a5adb4ca1cb62f90c30e328e073ba40443179d690e12d80546001600160a01b0319166001600160a01b03929092169182179055604051908152a180f35b50346102bb5760203660031901126102bb5761247b815f51602061485e5f395f51905f526124426130f4565b5f5160206148be5f395f51905f5254604051630776843760e31b81529485936001600160a01b0390921692849283929160048401613442565b03915afa9081156107455782610975939261249f575b505060405191829182613255565b6124bb92503d8091833e6124b3818361317e565b8101906133b8565b5f80612491565b50346102bb5760603660031901126102bb57806124dd6130f4565b6024356044356124eb6135ff565b6001600160a01b0383165f9081527f3e2bfe19d7b287e1320a5adb4ca1cb62f90c30e328e073ba40443179d690e1256020526040902054819060ff16156125c7575050604051630356129d60e11b60208201526001600160a01b03909216602483015260448201526064808201839052815261256860848261317e565b5f51602061487e5f395f51905f52546001600160a01b0316803b15610f705760405163104c8d4b60e31b81526020600482015291839183918290849082906125b4906024830190613297565b03925af1801561074557610f5f57505080f35b5f516020614a1e5f395f51905f52546001600160a01b038086165f9081527f3e2bfe19d7b287e1320a5adb4ca1cb62f90c30e328e073ba40443179d690e126602052604090205494956125689590949361262993831692909188911686613eec565b15612669575b604051630356129d60e11b60208201526001600160a01b03909216602483015260448201939093526064810192909252816084810161148c565b508361262f565b50346102bb5760203660031901126102bb57600435906001600160401b0382116102bb57366023830112156102bb5781600401356001600160401b0381116108c25760248160051b840101368111610741576126ca6135ba565b631bd8f8b560e11b83525f51602061491e5f395f51905f52602052604083205460ff166122cb575f51602061489e5f395f51905f529061270983613311565b92612717604051948561317e565b83526024602084019501945b8186106127355784611c468585613e11565b6020809161274288613120565b815201950194612723565b50346102bb5760403660031901126102bb576127676130f4565b906024356127736135ba565b5f51602061485e5f395f51905f52546001600160a01b03841693908390815b818110612823575b505015612814577f3a6c52328a7b3b726d0ec757d68f416b26ec2991ac4d4f95d450c504f5a0e52192935f527f3e2bfe19d7b287e1320a5adb4ca1cb62f90c30e328e073ba40443179d690e12b6020528160405f20556110bd6040519283928360209093929193604081019460018060a01b031681520152565b632711b74d60e11b8352600483fd5b8661283b825f51602061485e5f395f51905f5261338f565b905460039190911b1c6001600160a01b03161461285a57600101612792565b50505060015f8061279a565b50346102bb5760403660031901126102bb57610c9e60043561288661310a565b906128ac610c94825f525f51602061493e5f395f51905f52602052600160405f20015490565b613d6a565b50346102bb5760203660031901126102bb576020610d3c6004355f525f51602061493e5f395f51905f52602052600160405f20015490565b50346102bb576101003660031901126102bb57366084116102bb57608435906001600160a01b0382168083036108c25760a4356001600160401b038111610741576129389036906004016131e4565b9060403660c319011261074157335f9081525f5160206147be5f395f51905f5260205260409020545f51602061489e5f395f51905f529190612be85761297c6136f6565b61298461374b565b60ff5f5160206149de5f395f51905f525460a01c16612bd957156119ed576129ac8482613783565b6040516129b881613163565b3660e31215610750576040516129cd81613134565b8036610104116107585760c4905b6101048210612bc957505081525f5160206148de5f395f51905f5254805f51602061499e5f395f51905f525410156119cb57505f51602061481e5f395f51905f5254604051633cf65e3b60e01b815291516001600160a01b0390911690829086600483015b60028210612baf57505050610104816020933360448301523060648301526080600460848401375afa9081156119a6578491612b90575b5015611978575f5160206148fe5f395f51905f52546001600160a01b031680151594909190612aa86115ef83614382565b915f5160206147de5f395f51905f52548310159081612af157501561189a5785611843576040516358cc5acb60e01b60208201526116438161148c846004873360248601613805565b905060018060a01b035f5160206148be5f395f51905f52541660405180926330959fcb60e11b8252604482013360048401526040602484015281518091526020606484019201908a5b818110612b625750505091818060209403915afa90811561183857869161191b57505f611612565b825180516001600160a01b031685526020908101518186015287955060409094019390920191600101612b3a565b612ba9915060203d60201161194357611935818361317e565b5f612a77565b829350602080916001939451815201930191018492612a40565b81358152602091820191016129db565b638a943acd60e01b8452600484fd5b6342ee68b560e01b8452600484fd5b50346102bb5760203660031901126102bb57612c116130f4565b335f9081525f5160206147be5f395f51905f52602052604090205415610aab5763d93c394b60e01b82525f51602061491e5f395f51905f52602052604082205460ff16611be7576001600160a01b03168015612d7457335f9081527f3e2bfe19d7b287e1320a5adb4ca1cb62f90c30e328e073ba40443179d690e12560205260409020805460ff191660011790557f3e2bfe19d7b287e1320a5adb4ca1cb62f90c30e328e073ba40443179d690e1245442908101908110612d60577f0d8cfa10a3087b28d3c226ad9a37314860e7c3c0505a25a39e3cdefb3118a98a91606091604051612cfd81613134565b8281526020808201838152335f8181525f516020614a3e5f395f51905f528452604090819020945185546001600160a01b0319166001600160a01b039190911617855591516001949094019390935580519283529082019390935291820152a180f35b634e487b7160e01b83526011600452602483fd5b62cc6ac760e01b8252600482fd5b5034611de6575f366003190112611de657335f9081525f5160206147be5f395f51905f52602052604090205415612ff357612dbb6136f6565b612dc361374b565b5f5160206148fe5f395f51905f52546001600160a01b0316801515919082612f9f575b60405163edad0a1360e01b602082015233602482015260248152612e0b60448261317e565b5f51602061487e5f395f51905f52546001600160a01b031690813b1561075057612e568492839260405194858094819363104c8d4b60e31b8352602060048401526024830190613297565b03925af1801561180057908391612f8a575b50505f51602061499e5f395f51905f5254928315612d605782935f19015f51602061499e5f395f51905f525560018060a01b0333165f525f5160206147be5f395f51905f526020528260405f2055612f45575b50507f6f42117a557500c705ddf040a619d86f39101e6b74ac20d7b3e5943ba473fc7f6020604051338152a15f5160206149de5f395f51905f525481906001600160a01b0316803b15611142578180916024604051809481936351b27a6d60e11b83523360048401525af1801561074557611791575060015f51602061497e5f395f51905f525580f35b803b156111425781809160246040518094819363e9ecc1cb60e01b83523360048401525af180156107455715612ebb5781612f7f9161317e565b6102bb57805f612ebb565b81612f949161317e565b6108c257815f612e68565b803b15611de6576040516311c7e79960e21b81523360048201525f8160248183865af18015612fe857612fd3575b50612de6565b612fe09192505f9061317e565b5f905f612fcd565b6040513d5f823e3d90fd5b6325ec6c1f60e01b5f5260045ffd5b34611de6576040366003190112611de65761301b613575565b6004356024356001600160a01b0382168203611de6576001600160a01b03919091165f8181527f3e2bfe19d7b287e1320a5adb4ca1cb62f90c30e328e073ba40443179d690e12c602090815260409182902084905581519283528201929092527f8ae53ffd0ebc018acb19342fba690554d49ae9a467a9606a38b49cb5ad775c819190a1005b34611de6576020366003190112611de6576004359063ffffffff60e01b8216809203611de657602091637965db0b60e01b81149081156130e3575b5015158152f35b6301ffc9a760e01b149050836130dc565b600435906001600160a01b0382168203611de657565b602435906001600160a01b0382168203611de657565b35906001600160a01b0382168203611de657565b604081019081106001600160401b0382111761314f57604052565b634e487b7160e01b5f52604160045260245ffd5b602081019081106001600160401b0382111761314f57604052565b90601f801991011681019081106001600160401b0382111761314f57604052565b9291926001600160401b03821161314f57604051916131c8601f8201601f19166020018461317e565b829481845281830111611de6578281602093845f960137010152565b919091606081840312611de65760405190606082018281106001600160401b0382111761314f57604052819381356001600160401b038111611de65782019181601f84011215611de65761324260409392836020869535910161319f565b8452602081013560208501520135910152565b60206040818301928281528451809452019201905f5b8181106132785750505090565b82516001600160a01b031684526020938401939092019160010161326b565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b34611de6575f366003190112611de65760205f51602061499e5f395f51905f5254604051908152f35b9181601f84011215611de6578235916001600160401b038311611de65760208381860195010111611de657565b6001600160401b03811161314f5760051b60200190565b34611de6575f366003190112611de6576109756109695f51602061485e5f395f51905f52613522565b9060018060a01b03168152604060208201526080604061337c845160608386015260a0850190613297565b9360208101516060850152015191015290565b80548210156133a4575f5260205f2001905f90565b634e487b7160e01b5f52603260045260245ffd5b602081830312611de6578051906001600160401b038211611de657019080601f83011215611de6578151906133ec82613311565b926133fa604051948561317e565b82845260208085019360051b820101918211611de657602001915b8183106134225750505090565b82516001600160a01b0381168103611de657815260209283019201613415565b606090604081019260018060a01b0316815260406020820152835480935201915f5260205f20905f5b8181106134785750505090565b82546001600160a01b031684526020909301926001928301920161346b565b90600182811c921680156134c5575b60208310146134b157565b634e487b7160e01b5f52602260045260245ffd5b91607f16916134a6565b90918060409360208452816020850152848401375f828201840152601f01601f1916010190565b6040519061350382613134565b5f6020838281520152565b80518210156133a45760209160051b010190565b90604051918281549182825260208201905f5260205f20925f5b8181106135535750506135519250038361317e565b565b84546001600160a01b031683526001948501948794506020909301920161353c565b335f9081525f5160206147fe5f395f51905f52602052604090205460ff161561359a57565b63e2517d3f60e01b5f9081523360045263ef0892d160e01b602452604490fd5b335f9081525f5160206149fe5f395f51905f52602052604090205460ff16156135df57565b63e2517d3f60e01b5f908152336004526378b4401360e11b602452604490fd5b335f9081525f51602061483e5f395f51905f52602052604090205460ff161561362457565b63e2517d3f60e01b5f90815233600452638a70a0eb60e01b602452604490fd5b335f9081527f274b5753bc2a873526e44bae648b363f47953a4f0c2234820809428fe7b7dafd602052604090205460ff161561367c57565b63e2517d3f60e01b5f9081523360045263d8a8b5c760e01b602452604490fd5b5f8181525f51602061493e5f395f51905f526020908152604080832033845290915290205460ff16156136cc5750565b63e2517d3f60e01b5f523360045260245260445ffd5b356001600160a01b0381168103611de65790565b63024b274760e61b5f525f51602061491e5f395f51905f526020527f6fea4dda5e9bec14cdaf5c3123c3f3d726de9276141f697f28abf1342e0bf5f55460ff1661373c57565b63722fdba960e01b5f5260045ffd5b60025f51602061497e5f395f51905f5254146137745760025f51602061497e5f395f51905f5255565b633ee5aeb560e01b5f5260045ffd5b335f818152600d92909201602090815260409283902080546001600160a01b0319166001600160a01b03909516948517905582519182528101929092527fe906feea2ef60b474e22b4169bdd4de6906a84cd448cbcee99593526fe87082d9190819081015b0390a1565b90816020910312611de657518015158103611de65790565b9260809060c0939695929660e086019760018060a01b03168652602086015260408501376001600160a01b0316910152565b6001600160a01b0381165f9081525f5160206147fe5f395f51905f52602052604090205460ff166138b1576001600160a01b03165f8181525f5160206147fe5f395f51905f5260205260408120805460ff1916600117905533919063ef0892d160e01b905f51602061479e5f395f51905f529080a4600190565b505f90565b6001600160a01b0381165f9081525f51602061483e5f395f51905f52602052604090205460ff166138b1576001600160a01b03165f8181525f51602061483e5f395f51905f5260205260408120805460ff19166001179055339190638a70a0eb60e01b905f51602061479e5f395f51905f529080a4600190565b6001600160a01b0381165f9081525f5160206149fe5f395f51905f52602052604090205460ff166138b1576001600160a01b03165f8181525f5160206149fe5f395f51905f5260205260408120805460ff191660011790553391906378b4401360e11b905f51602061479e5f395f51905f529080a4600190565b6001600160a01b0381165f9081527f274b5753bc2a873526e44bae648b363f47953a4f0c2234820809428fe7b7dafd602052604090205460ff166138b1576001600160a01b03165f8181527f274b5753bc2a873526e44bae648b363f47953a4f0c2234820809428fe7b7dafd60205260408120805460ff1916600117905533919063d8a8b5c760e01b905f51602061479e5f395f51905f529080a4600190565b6001600160a01b0381165f9081527f55283b829f13bbb718097e9d9a0a5f887f18c6e64a14e5b6837b3d53eab0a8f3602052604090205460ff166138b1576001600160a01b03165f8181527f55283b829f13bbb718097e9d9a0a5f887f18c6e64a14e5b6837b3d53eab0a8f360205260408120805460ff1916600117905533919063d3e319af60e01b905f51602061479e5f395f51905f529080a4600190565b6001600160a01b0381165f9081527f05577d836c3af8e2cf1127dcd027858a1d8f3404907a19d463fc4a26a60be426602052604090205460ff166138b1576001600160a01b03165f8181527f05577d836c3af8e2cf1127dcd027858a1d8f3404907a19d463fc4a26a60be42660205260408120805460ff1916600117905533919063024b274760e61b905f51602061479e5f395f51905f529080a4600190565b6001600160a01b0381165f9081527f83fcafa020a3c9213bdc7a9147feab576ec07cc55f595b8fb0fb928f72fcf184602052604090205460ff166138b1576001600160a01b03165f8181527f83fcafa020a3c9213bdc7a9147feab576ec07cc55f595b8fb0fb928f72fcf18460205260408120805460ff1916600117905533919063d93c394b60e01b905f51602061479e5f395f51905f529080a4600190565b6001600160a01b0381165f9081527f3662b209663523262d9d456b8e64f3b7e6057f997cdc7568782ab5f831e15efc602052604090205460ff166138b1576001600160a01b03165f8181527f3662b209663523262d9d456b8e64f3b7e6057f997cdc7568782ab5f831e15efc60205260408120805460ff19166001179055339190631bd8f8b560e11b905f51602061479e5f395f51905f529080a4600190565b6001600160a01b0381165f9081527ffec025993670f3cfe62fe0b8ba8699907af9264646da1759461a905da18f569c602052604090205460ff166138b1576001600160a01b03165f8181527ffec025993670f3cfe62fe0b8ba8699907af9264646da1759461a905da18f569c60205260408120805460ff191660011790553391906362250a9560e11b905f51602061479e5f395f51905f529080a4600190565b5f8181525f51602061493e5f395f51905f52602090815260408083206001600160a01b038616845290915290205460ff16613df5575f8181525f51602061493e5f395f51905f52602090815260408083206001600160a01b0395909516808452949091528120805460ff19166001179055339291905f51602061479e5f395f51905f529080a4600190565b50505f90565b818110613e06575050565b5f8155600101613dfb565b60090180545f82559092919080613ed2575b505f5b8151811015613e9d576001600160a01b03613e41828461350e565b5116908115613e8e578454916801000000000000000083101561314f57613e6f83600180950188558761338f565b819291549060031b91821b91858060a01b03901b191617905501613e26565b632711b74d60e11b5f5260045ffd5b507ff009a6ffded424f714e8904d643a1ea4479453188faf08a3996121996b76684f9192506137e89060405191829182613255565b613ee690845f5260205f2090810190613dfb565b5f613e23565b9293926001600160a01b03821615613f735750604051633256b4d160e01b81526001600160a01b039091166004820152602481019190915260448101929092526020908290815f81606481015b03926001600160a01b03165af1908115612fe8575f91613f57575090565b613f70915060203d60201161194357611935818361317e565b90565b604051633256b4d160e01b81526001600160a01b03909116600482015260248101929092525060448101929092526020908290815f8160648101613f39565b6001600160a01b0381165f9081525f5160206147fe5f395f51905f52602052604090205460ff16156138b1576001600160a01b03165f8181525f5160206147fe5f395f51905f5260205260408120805460ff1916905533919063ef0892d160e01b905f51602061495e5f395f51905f529080a4600190565b6001600160a01b0381165f9081525f51602061483e5f395f51905f52602052604090205460ff16156138b1576001600160a01b03165f8181525f51602061483e5f395f51905f5260205260408120805460ff19169055339190638a70a0eb60e01b905f51602061495e5f395f51905f529080a4600190565b6001600160a01b0381165f9081525f5160206149fe5f395f51905f52602052604090205460ff16156138b1576001600160a01b03165f8181525f5160206149fe5f395f51905f5260205260408120805460ff191690553391906378b4401360e11b905f51602061495e5f395f51905f529080a4600190565b6001600160a01b0381165f9081527fb7db2dd08fcb62d0c9e08c51941cae53c267786a0b75803fb7960902fc8ef97d602052604090205460ff16156138b1576001600160a01b03165f8181527fb7db2dd08fcb62d0c9e08c51941cae53c267786a0b75803fb7960902fc8ef97d60205260408120805460ff191690553391905f51602061495e5f395f51905f528180a4600190565b5f8181525f51602061493e5f395f51905f52602090815260408083206001600160a01b038616845290915290205460ff1615613df5575f8181525f51602061493e5f395f51905f52602090815260408083206001600160a01b0395909516808452949091528120805460ff19169055339291905f51602061495e5f395f51905f529080a4600190565b600a019291906001600160401b03821161314f576142568454613497565b601f8111614347575b505f93601f83116001146142c557827f7f63aacad63bc1693280450d5c3612ccd4efc53e46d69f3a537db102cd66290c94955f916142ba575b508360011b905f198560031b1c19161790555b6137e8604051928392836134cf565b90508201355f614298565b601f19831694815f5260205f20905f5b87811061432f5750847f7f63aacad63bc1693280450d5c3612ccd4efc53e46d69f3a537db102cd66290c969710614316575b5050600183811b0190556142ab565b8301355f19600386901b60f8161c191690555f80614307565b909160206001819285880135815501930191016142d5565b61437290855f5260205f20601f850160051c81019160208610614378575b601f0160051c0190613dfb565b5f61425f565b9091508190614365565b9061438f60098301613522565b9182519061439c82613311565b916143aa604051938461317e565b8083526143b9601f1991613311565b015f5b8181106144db5750508351936143d185613311565b946143df604051968761317e565b8086526143ee601f1991613311565b015f5b8181106144b85750505f5b81518110156144b1576001906001600160a01b0361441a828561350e565b5116828060a01b0381165f526012850160205260405f20546040519061443f82613134565b8282526020820152614451838861350e565b5261445c828761350e565b50828060a01b0381165f526013850160205260405f205480156144aa575b6040519161448783613134565b82526020820152614498828961350e565b526144a3818861350e565b50016143fc565b508261447a565b5050509190565b6020906040516144c781613134565b5f81525f8382015282828a010152016143f1565b6020906040516144ea81613134565b5f81525f83820152828287010152016143bc565b60118101546003909101546040805163f679e15f60e01b81526001600160a01b0394851660048201526024810191909152845160448201819052909492939091169184916064830191602001905f5b8181106145c35750505091818060209403915afa918215612fe8575f9261458f575b5080821180614586575b614581575090565b905090565b50801515614579565b9091506020813d6020116145bb575b816145ab6020938361317e565b81010312611de65751905f61456f565b3d915061459e565b825180516001600160a01b03168552602090810151818601528895506040909401939092019160010161454d565b600301546040516301753ab960e31b8152466004820152905f90829060249082906001600160a01b03165afa908115612fe8575f9161462e575090565b613f7091503d805f833e6124b3818361317e565b60ff5f5160206149be5f395f51905f525460401c161561465e57565b631afcd79f60e31b5f5260045ffd5b815191906041830361469d576146969250602082015190606060408401519301515f1a9061471b565b9192909190565b50505f9160029190565b600481101561470757806146b9575050565b600181036146d05763f645eedf60e01b5f5260045ffd5b600281036146eb575063fce698f760e01b5f5260045260245ffd5b6003146146f55750565b6335e2f38360e21b5f5260045260245ffd5b634e487b7160e01b5f52602160045260245ffd5b91907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411614792579160209360809260ff5f9560405194855216868401526040830152606082015282805260015afa15612fe8575f516001600160a01b0381161561478857905f905f90565b505f906001905f90565b5050505f916003919056fe2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d3e2bfe19d7b287e1320a5adb4ca1cb62f90c30e328e073ba40443179d690e1203e2bfe19d7b287e1320a5adb4ca1cb62f90c30e328e073ba40443179d690e1290b373776dcbf64ceb4d34b4e217820b0691f48f308d42843c6080fb1fac6bd243e2bfe19d7b287e1320a5adb4ca1cb62f90c30e328e073ba40443179d690e12ed480ef4a5d78515f50ce9d1a72eb9abf0b7795388f968bf633429bfc14cfe40b3e2bfe19d7b287e1320a5adb4ca1cb62f90c30e328e073ba40443179d690e1223e2bfe19d7b287e1320a5adb4ca1cb62f90c30e328e073ba40443179d690e11b3e2bfe19d7b287e1320a5adb4ca1cb62f90c30e328e073ba40443179d690e1193e2bfe19d7b287e1320a5adb4ca1cb62f90c30e328e073ba40443179d690e11c3e2bfe19d7b287e1320a5adb4ca1cb62f90c30e328e073ba40443179d690e1283e2bfe19d7b287e1320a5adb4ca1cb62f90c30e328e073ba40443179d690e121fe6065fb4e9872e2ad4479001655335380d83f70e163706cd65857449b84510002dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800f6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f003e2bfe19d7b287e1320a5adb4ca1cb62f90c30e328e073ba40443179d690e11af0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a003e2bfe19d7b287e1320a5adb4ca1cb62f90c30e328e073ba40443179d690e11eb8c1f5d3f58d95cae9dbd2762e9e87b212a41bbf0f42459bc1feb152e30e1cec3e2bfe19d7b287e1320a5adb4ca1cb62f90c30e328e073ba40443179d690e11d3e2bfe19d7b287e1320a5adb4ca1cb62f90c30e328e073ba40443179d690e127a26469706673582212200d6e51287b44df1aadf4e72810fca548d70ec295b503b42c7707f3811d3c63ac64736f6c634300081b0033",
}

// ContractAvsGovernanceABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractAvsGovernanceMetaData.ABI instead.
var ContractAvsGovernanceABI = ContractAvsGovernanceMetaData.ABI

// ContractAvsGovernanceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractAvsGovernanceMetaData.Bin instead.
var ContractAvsGovernanceBin = ContractAvsGovernanceMetaData.Bin

// DeployContractAvsGovernance deploys a new Ethereum contract, binding an instance of ContractAvsGovernance to it.
func DeployContractAvsGovernance(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContractAvsGovernance, error) {
	parsed, err := ContractAvsGovernanceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractAvsGovernanceBin), backend)
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

	AvsDirectory(opts *bind.CallOpts) (common.Address, error)

	AvsName(opts *bind.CallOpts) (string, error)

	GetDefaultStrategies(opts *bind.CallOpts) ([]common.Address, error)

	GetIsAllowlisted(opts *bind.CallOpts) (bool, error)

	GetNumOfOperatorsLimit(opts *bind.CallOpts) (*big.Int, error)

	GetOperatorRestakedStrategies(opts *bind.CallOpts, _operator common.Address) ([]common.Address, error)

	GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error)

	GetRewardsReceiver(opts *bind.CallOpts, _operator common.Address) (common.Address, error)

	GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error)

	HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error)

	IsFlowPaused(opts *bind.CallOpts, _pausableFlow [4]byte) (bool, error)

	IsOperatorRegistered(opts *bind.CallOpts, operator common.Address) (bool, error)

	MaxEffectiveBalance(opts *bind.CallOpts) (*big.Int, error)

	MinSharesForStrategy(opts *bind.CallOpts, _strategy common.Address) (*big.Int, error)

	MinVotingPower(opts *bind.CallOpts) (*big.Int, error)

	NumOfActiveOperators(opts *bind.CallOpts) (*big.Int, error)

	NumOfOperators(opts *bind.CallOpts) (*big.Int, error)

	NumOfShares(opts *bind.CallOpts, _operator common.Address) (*big.Int, error)

	Strategies(opts *bind.CallOpts) ([]common.Address, error)

	StrategyMultiplier(opts *bind.CallOpts, _strategy common.Address) (*big.Int, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	Vault(opts *bind.CallOpts) (common.Address, error)

	VotingPower(opts *bind.CallOpts, _operator common.Address) (*big.Int, error)
}

// ContractAvsGovernanceTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractAvsGovernanceTransacts interface {
	CompleteRewardsReceiverModification(opts *bind.TransactOpts) (*types.Transaction, error)

	DepositERC20(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error)

	GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts, _initializationParams IAvsGovernanceInitializationParams) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts, _pausableFlow [4]byte) (*types.Transaction, error)

	QueueRewardsReceiverModification(opts *bind.TransactOpts, _newRewardsReceiver common.Address) (*types.Transaction, error)

	RegisterAsAllowedOperator(opts *bind.TransactOpts, _blsKey [4]*big.Int, _authToken []byte, _rewardsReceiver common.Address, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry, _blsRegistrationSignature BLSAuthLibrarySignature) (*types.Transaction, error)

	RegisterAsOperator(opts *bind.TransactOpts, _blsKey [4]*big.Int, _rewardsReceiver common.Address, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry, _blsRegistrationSignature BLSAuthLibrarySignature) (*types.Transaction, error)

	RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error)

	RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	SetAllowlistSigner(opts *bind.TransactOpts, _allowlistSigner common.Address) (*types.Transaction, error)

	SetAvsGovernanceLogic(opts *bind.TransactOpts, _avsGovernanceLogic common.Address) (*types.Transaction, error)

	SetAvsGovernanceMultiplierSyncer(opts *bind.TransactOpts, _newAvsGovernanceMultiplierSyncer common.Address) (*types.Transaction, error)

	SetAvsName(opts *bind.TransactOpts, _avsName string) (*types.Transaction, error)

	SetBLSAuthSingleton(opts *bind.TransactOpts, _blsAuthSingleton common.Address) (*types.Transaction, error)

	SetIsAllowlisted(opts *bind.TransactOpts, _isAllowlisted bool) (*types.Transaction, error)

	SetMaxEffectiveBalance(opts *bind.TransactOpts, _maxBalance *big.Int) (*types.Transaction, error)

	SetMinSharesForStrategy(opts *bind.TransactOpts, _strategy common.Address, _minShares *big.Int) (*types.Transaction, error)

	SetMinVotingPower(opts *bind.TransactOpts, _minVotingPower *big.Int) (*types.Transaction, error)

	SetNumOfOperatorsLimit(opts *bind.TransactOpts, _newLimitOfNumOfOperators *big.Int) (*types.Transaction, error)

	SetOthenticRegistry(opts *bind.TransactOpts, _othenticRegistry common.Address) (*types.Transaction, error)

	SetRewardsReceiverModificationDelay(opts *bind.TransactOpts, _rewardsReceiverModificationDelay *big.Int) (*types.Transaction, error)

	SetStrategyMultiplier(opts *bind.TransactOpts, _strategyMultiplier IAvsGovernanceStrategyMultiplier) (*types.Transaction, error)

	SetStrategyMultiplierBatch(opts *bind.TransactOpts, _strategyMultipliers []IAvsGovernanceStrategyMultiplier) (*types.Transaction, error)

	SetSupportedStrategies(opts *bind.TransactOpts, _strategies []common.Address) (*types.Transaction, error)

	TransferAvsGovernanceMultisig(opts *bind.TransactOpts, _newAvsGovernanceMultisig common.Address) (*types.Transaction, error)

	TransferMessageHandler(opts *bind.TransactOpts, _newMessageHandler common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts, _pausableFlow [4]byte) (*types.Transaction, error)

	UnregisterAsOperator(opts *bind.TransactOpts) (*types.Transaction, error)

	UpdateAVSMetadataURI(opts *bind.TransactOpts, metadataURI string) (*types.Transaction, error)

	WithdrawBatchRewards(opts *bind.TransactOpts, _operators []IAvsGovernancePaymentRequestMessage, _lastPayedTask *big.Int) (*types.Transaction, error)

	WithdrawRewards(opts *bind.TransactOpts, _operator common.Address, _lastPayedTask *big.Int, _feeToClaim *big.Int) (*types.Transaction, error)
}

// ContractAvsGovernanceFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractAvsGovernanceFilters interface {
	FilterBLSAuthSingletonSet(opts *bind.FilterOpts) (*ContractAvsGovernanceBLSAuthSingletonSetIterator, error)
	WatchBLSAuthSingletonSet(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceBLSAuthSingletonSet) (event.Subscription, error)
	ParseBLSAuthSingletonSet(log types.Log) (*ContractAvsGovernanceBLSAuthSingletonSet, error)

	FilterFlowPaused(opts *bind.FilterOpts) (*ContractAvsGovernanceFlowPausedIterator, error)
	WatchFlowPaused(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceFlowPaused) (event.Subscription, error)
	ParseFlowPaused(log types.Log) (*ContractAvsGovernanceFlowPaused, error)

	FilterFlowUnpaused(opts *bind.FilterOpts) (*ContractAvsGovernanceFlowUnpausedIterator, error)
	WatchFlowUnpaused(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceFlowUnpaused) (event.Subscription, error)
	ParseFlowUnpaused(log types.Log) (*ContractAvsGovernanceFlowUnpaused, error)

	FilterInitialized(opts *bind.FilterOpts) (*ContractAvsGovernanceInitializedIterator, error)
	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceInitialized) (event.Subscription, error)
	ParseInitialized(log types.Log) (*ContractAvsGovernanceInitialized, error)

	FilterMaxEffectiveBalanceSet(opts *bind.FilterOpts) (*ContractAvsGovernanceMaxEffectiveBalanceSetIterator, error)
	WatchMaxEffectiveBalanceSet(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceMaxEffectiveBalanceSet) (event.Subscription, error)
	ParseMaxEffectiveBalanceSet(log types.Log) (*ContractAvsGovernanceMaxEffectiveBalanceSet, error)

	FilterMinSharesPerStrategySet(opts *bind.FilterOpts) (*ContractAvsGovernanceMinSharesPerStrategySetIterator, error)
	WatchMinSharesPerStrategySet(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceMinSharesPerStrategySet) (event.Subscription, error)
	ParseMinSharesPerStrategySet(log types.Log) (*ContractAvsGovernanceMinSharesPerStrategySet, error)

	FilterMinVotingPowerSet(opts *bind.FilterOpts) (*ContractAvsGovernanceMinVotingPowerSetIterator, error)
	WatchMinVotingPowerSet(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceMinVotingPowerSet) (event.Subscription, error)
	ParseMinVotingPowerSet(log types.Log) (*ContractAvsGovernanceMinVotingPowerSet, error)

	FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address) (*ContractAvsGovernanceOperatorRegisteredIterator, error)
	WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceOperatorRegistered, operator []common.Address) (event.Subscription, error)
	ParseOperatorRegistered(log types.Log) (*ContractAvsGovernanceOperatorRegistered, error)

	FilterOperatorUnregistered(opts *bind.FilterOpts) (*ContractAvsGovernanceOperatorUnregisteredIterator, error)
	WatchOperatorUnregistered(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceOperatorUnregistered) (event.Subscription, error)
	ParseOperatorUnregistered(log types.Log) (*ContractAvsGovernanceOperatorUnregistered, error)

	FilterQueuedRewardsReceiverModification(opts *bind.FilterOpts) (*ContractAvsGovernanceQueuedRewardsReceiverModificationIterator, error)
	WatchQueuedRewardsReceiverModification(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceQueuedRewardsReceiverModification) (event.Subscription, error)
	ParseQueuedRewardsReceiverModification(log types.Log) (*ContractAvsGovernanceQueuedRewardsReceiverModification, error)

	FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ContractAvsGovernanceRoleAdminChangedIterator, error)
	WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error)
	ParseRoleAdminChanged(log types.Log) (*ContractAvsGovernanceRoleAdminChanged, error)

	FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ContractAvsGovernanceRoleGrantedIterator, error)
	WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error)
	ParseRoleGranted(log types.Log) (*ContractAvsGovernanceRoleGranted, error)

	FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ContractAvsGovernanceRoleRevokedIterator, error)
	WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error)
	ParseRoleRevoked(log types.Log) (*ContractAvsGovernanceRoleRevoked, error)

	FilterSetAllowlistSigner(opts *bind.FilterOpts) (*ContractAvsGovernanceSetAllowlistSignerIterator, error)
	WatchSetAllowlistSigner(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetAllowlistSigner) (event.Subscription, error)
	ParseSetAllowlistSigner(log types.Log) (*ContractAvsGovernanceSetAllowlistSigner, error)

	FilterSetAvsGovernanceLogic(opts *bind.FilterOpts) (*ContractAvsGovernanceSetAvsGovernanceLogicIterator, error)
	WatchSetAvsGovernanceLogic(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetAvsGovernanceLogic) (event.Subscription, error)
	ParseSetAvsGovernanceLogic(log types.Log) (*ContractAvsGovernanceSetAvsGovernanceLogic, error)

	FilterSetAvsGovernanceMultiplierSyncer(opts *bind.FilterOpts) (*ContractAvsGovernanceSetAvsGovernanceMultiplierSyncerIterator, error)
	WatchSetAvsGovernanceMultiplierSyncer(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetAvsGovernanceMultiplierSyncer) (event.Subscription, error)
	ParseSetAvsGovernanceMultiplierSyncer(log types.Log) (*ContractAvsGovernanceSetAvsGovernanceMultiplierSyncer, error)

	FilterSetAvsGovernanceMultisig(opts *bind.FilterOpts) (*ContractAvsGovernanceSetAvsGovernanceMultisigIterator, error)
	WatchSetAvsGovernanceMultisig(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetAvsGovernanceMultisig) (event.Subscription, error)
	ParseSetAvsGovernanceMultisig(log types.Log) (*ContractAvsGovernanceSetAvsGovernanceMultisig, error)

	FilterSetAvsName(opts *bind.FilterOpts) (*ContractAvsGovernanceSetAvsNameIterator, error)
	WatchSetAvsName(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetAvsName) (event.Subscription, error)
	ParseSetAvsName(log types.Log) (*ContractAvsGovernanceSetAvsName, error)

	FilterSetIsAllowlisted(opts *bind.FilterOpts) (*ContractAvsGovernanceSetIsAllowlistedIterator, error)
	WatchSetIsAllowlisted(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetIsAllowlisted) (event.Subscription, error)
	ParseSetIsAllowlisted(log types.Log) (*ContractAvsGovernanceSetIsAllowlisted, error)

	FilterSetMessageHandler(opts *bind.FilterOpts) (*ContractAvsGovernanceSetMessageHandlerIterator, error)
	WatchSetMessageHandler(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetMessageHandler) (event.Subscription, error)
	ParseSetMessageHandler(log types.Log) (*ContractAvsGovernanceSetMessageHandler, error)

	FilterSetNumOfOperatorsLimit(opts *bind.FilterOpts) (*ContractAvsGovernanceSetNumOfOperatorsLimitIterator, error)
	WatchSetNumOfOperatorsLimit(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetNumOfOperatorsLimit) (event.Subscription, error)
	ParseSetNumOfOperatorsLimit(log types.Log) (*ContractAvsGovernanceSetNumOfOperatorsLimit, error)

	FilterSetOthenticRegistry(opts *bind.FilterOpts) (*ContractAvsGovernanceSetOthenticRegistryIterator, error)
	WatchSetOthenticRegistry(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetOthenticRegistry) (event.Subscription, error)
	ParseSetOthenticRegistry(log types.Log) (*ContractAvsGovernanceSetOthenticRegistry, error)

	FilterSetRewardsReceiver(opts *bind.FilterOpts) (*ContractAvsGovernanceSetRewardsReceiverIterator, error)
	WatchSetRewardsReceiver(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetRewardsReceiver) (event.Subscription, error)
	ParseSetRewardsReceiver(log types.Log) (*ContractAvsGovernanceSetRewardsReceiver, error)

	FilterSetRewardsReceiverModificationDelay(opts *bind.FilterOpts) (*ContractAvsGovernanceSetRewardsReceiverModificationDelayIterator, error)
	WatchSetRewardsReceiverModificationDelay(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetRewardsReceiverModificationDelay) (event.Subscription, error)
	ParseSetRewardsReceiverModificationDelay(log types.Log) (*ContractAvsGovernanceSetRewardsReceiverModificationDelay, error)

	FilterSetStrategyMultiplier(opts *bind.FilterOpts) (*ContractAvsGovernanceSetStrategyMultiplierIterator, error)
	WatchSetStrategyMultiplier(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetStrategyMultiplier) (event.Subscription, error)
	ParseSetStrategyMultiplier(log types.Log) (*ContractAvsGovernanceSetStrategyMultiplier, error)

	FilterSetSupportedStrategies(opts *bind.FilterOpts) (*ContractAvsGovernanceSetSupportedStrategiesIterator, error)
	WatchSetSupportedStrategies(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetSupportedStrategies) (event.Subscription, error)
	ParseSetSupportedStrategies(log types.Log) (*ContractAvsGovernanceSetSupportedStrategies, error)

	FilterSetToken(opts *bind.FilterOpts) (*ContractAvsGovernanceSetTokenIterator, error)
	WatchSetToken(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetToken) (event.Subscription, error)
	ParseSetToken(log types.Log) (*ContractAvsGovernanceSetToken, error)
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

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) AvsDirectory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "avsDirectory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) AvsDirectory() (common.Address, error) {
	return _ContractAvsGovernance.Contract.AvsDirectory(&_ContractAvsGovernance.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) AvsDirectory() (common.Address, error) {
	return _ContractAvsGovernance.Contract.AvsDirectory(&_ContractAvsGovernance.CallOpts)
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

// GetDefaultStrategies is a free data retrieval call binding the contract method 0xe86685d9.
//
// Solidity: function getDefaultStrategies() view returns(address[])
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) GetDefaultStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "getDefaultStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetDefaultStrategies is a free data retrieval call binding the contract method 0xe86685d9.
//
// Solidity: function getDefaultStrategies() view returns(address[])
func (_ContractAvsGovernance *ContractAvsGovernanceSession) GetDefaultStrategies() ([]common.Address, error) {
	return _ContractAvsGovernance.Contract.GetDefaultStrategies(&_ContractAvsGovernance.CallOpts)
}

// GetDefaultStrategies is a free data retrieval call binding the contract method 0xe86685d9.
//
// Solidity: function getDefaultStrategies() view returns(address[])
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) GetDefaultStrategies() ([]common.Address, error) {
	return _ContractAvsGovernance.Contract.GetDefaultStrategies(&_ContractAvsGovernance.CallOpts)
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

// GetNumOfOperatorsLimit is a free data retrieval call binding the contract method 0xf251c9a6.
//
// Solidity: function getNumOfOperatorsLimit() view returns(uint256 numOfOperatorsLimitView)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) GetNumOfOperatorsLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "getNumOfOperatorsLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumOfOperatorsLimit is a free data retrieval call binding the contract method 0xf251c9a6.
//
// Solidity: function getNumOfOperatorsLimit() view returns(uint256 numOfOperatorsLimitView)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) GetNumOfOperatorsLimit() (*big.Int, error) {
	return _ContractAvsGovernance.Contract.GetNumOfOperatorsLimit(&_ContractAvsGovernance.CallOpts)
}

// GetNumOfOperatorsLimit is a free data retrieval call binding the contract method 0xf251c9a6.
//
// Solidity: function getNumOfOperatorsLimit() view returns(uint256 numOfOperatorsLimitView)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) GetNumOfOperatorsLimit() (*big.Int, error) {
	return _ContractAvsGovernance.Contract.GetNumOfOperatorsLimit(&_ContractAvsGovernance.CallOpts)
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

// MinSharesForStrategy is a free data retrieval call binding the contract method 0xc3814e5b.
//
// Solidity: function minSharesForStrategy(address _strategy) view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) MinSharesForStrategy(opts *bind.CallOpts, _strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "minSharesForStrategy", _strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinSharesForStrategy is a free data retrieval call binding the contract method 0xc3814e5b.
//
// Solidity: function minSharesForStrategy(address _strategy) view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) MinSharesForStrategy(_strategy common.Address) (*big.Int, error) {
	return _ContractAvsGovernance.Contract.MinSharesForStrategy(&_ContractAvsGovernance.CallOpts, _strategy)
}

// MinSharesForStrategy is a free data retrieval call binding the contract method 0xc3814e5b.
//
// Solidity: function minSharesForStrategy(address _strategy) view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) MinSharesForStrategy(_strategy common.Address) (*big.Int, error) {
	return _ContractAvsGovernance.Contract.MinSharesForStrategy(&_ContractAvsGovernance.CallOpts, _strategy)
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

// NumOfOperators is a free data retrieval call binding the contract method 0x6ade02da.
//
// Solidity: function numOfOperators() view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) NumOfOperators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "numOfOperators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumOfOperators is a free data retrieval call binding the contract method 0x6ade02da.
//
// Solidity: function numOfOperators() view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) NumOfOperators() (*big.Int, error) {
	return _ContractAvsGovernance.Contract.NumOfOperators(&_ContractAvsGovernance.CallOpts)
}

// NumOfOperators is a free data retrieval call binding the contract method 0x6ade02da.
//
// Solidity: function numOfOperators() view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) NumOfOperators() (*big.Int, error) {
	return _ContractAvsGovernance.Contract.NumOfOperators(&_ContractAvsGovernance.CallOpts)
}

// NumOfShares is a free data retrieval call binding the contract method 0x6a907803.
//
// Solidity: function numOfShares(address _operator) view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) NumOfShares(opts *bind.CallOpts, _operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "numOfShares", _operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumOfShares is a free data retrieval call binding the contract method 0x6a907803.
//
// Solidity: function numOfShares(address _operator) view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) NumOfShares(_operator common.Address) (*big.Int, error) {
	return _ContractAvsGovernance.Contract.NumOfShares(&_ContractAvsGovernance.CallOpts, _operator)
}

// NumOfShares is a free data retrieval call binding the contract method 0x6a907803.
//
// Solidity: function numOfShares(address _operator) view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) NumOfShares(_operator common.Address) (*big.Int, error) {
	return _ContractAvsGovernance.Contract.NumOfShares(&_ContractAvsGovernance.CallOpts, _operator)
}

// Strategies is a free data retrieval call binding the contract method 0xd9f9027f.
//
// Solidity: function strategies() view returns(address[])
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) Strategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "strategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// Strategies is a free data retrieval call binding the contract method 0xd9f9027f.
//
// Solidity: function strategies() view returns(address[])
func (_ContractAvsGovernance *ContractAvsGovernanceSession) Strategies() ([]common.Address, error) {
	return _ContractAvsGovernance.Contract.Strategies(&_ContractAvsGovernance.CallOpts)
}

// Strategies is a free data retrieval call binding the contract method 0xd9f9027f.
//
// Solidity: function strategies() view returns(address[])
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) Strategies() ([]common.Address, error) {
	return _ContractAvsGovernance.Contract.Strategies(&_ContractAvsGovernance.CallOpts)
}

// StrategyMultiplier is a free data retrieval call binding the contract method 0x8f53bc50.
//
// Solidity: function strategyMultiplier(address _strategy) view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) StrategyMultiplier(opts *bind.CallOpts, _strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "strategyMultiplier", _strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StrategyMultiplier is a free data retrieval call binding the contract method 0x8f53bc50.
//
// Solidity: function strategyMultiplier(address _strategy) view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) StrategyMultiplier(_strategy common.Address) (*big.Int, error) {
	return _ContractAvsGovernance.Contract.StrategyMultiplier(&_ContractAvsGovernance.CallOpts, _strategy)
}

// StrategyMultiplier is a free data retrieval call binding the contract method 0x8f53bc50.
//
// Solidity: function strategyMultiplier(address _strategy) view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) StrategyMultiplier(_strategy common.Address) (*big.Int, error) {
	return _ContractAvsGovernance.Contract.StrategyMultiplier(&_ContractAvsGovernance.CallOpts, _strategy)
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

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) Vault() (common.Address, error) {
	return _ContractAvsGovernance.Contract.Vault(&_ContractAvsGovernance.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) Vault() (common.Address, error) {
	return _ContractAvsGovernance.Contract.Vault(&_ContractAvsGovernance.CallOpts)
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

// DepositERC20 is a paid mutator transaction binding the contract method 0xb79092fd.
//
// Solidity: function depositERC20(uint256 _amount) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) DepositERC20(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "depositERC20", _amount)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0xb79092fd.
//
// Solidity: function depositERC20(uint256 _amount) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) DepositERC20(_amount *big.Int) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.DepositERC20(&_ContractAvsGovernance.TransactOpts, _amount)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0xb79092fd.
//
// Solidity: function depositERC20(uint256 _amount) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) DepositERC20(_amount *big.Int) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.DepositERC20(&_ContractAvsGovernance.TransactOpts, _amount)
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

// RegisterAsAllowedOperator is a paid mutator transaction binding the contract method 0x93304a9d.
//
// Solidity: function registerAsAllowedOperator(uint256[4] _blsKey, bytes _authToken, address _rewardsReceiver, (bytes,bytes32,uint256) _operatorSignature, (uint256[2]) _blsRegistrationSignature) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) RegisterAsAllowedOperator(opts *bind.TransactOpts, _blsKey [4]*big.Int, _authToken []byte, _rewardsReceiver common.Address, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry, _blsRegistrationSignature BLSAuthLibrarySignature) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "registerAsAllowedOperator", _blsKey, _authToken, _rewardsReceiver, _operatorSignature, _blsRegistrationSignature)
}

// RegisterAsAllowedOperator is a paid mutator transaction binding the contract method 0x93304a9d.
//
// Solidity: function registerAsAllowedOperator(uint256[4] _blsKey, bytes _authToken, address _rewardsReceiver, (bytes,bytes32,uint256) _operatorSignature, (uint256[2]) _blsRegistrationSignature) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) RegisterAsAllowedOperator(_blsKey [4]*big.Int, _authToken []byte, _rewardsReceiver common.Address, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry, _blsRegistrationSignature BLSAuthLibrarySignature) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.RegisterAsAllowedOperator(&_ContractAvsGovernance.TransactOpts, _blsKey, _authToken, _rewardsReceiver, _operatorSignature, _blsRegistrationSignature)
}

// RegisterAsAllowedOperator is a paid mutator transaction binding the contract method 0x93304a9d.
//
// Solidity: function registerAsAllowedOperator(uint256[4] _blsKey, bytes _authToken, address _rewardsReceiver, (bytes,bytes32,uint256) _operatorSignature, (uint256[2]) _blsRegistrationSignature) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) RegisterAsAllowedOperator(_blsKey [4]*big.Int, _authToken []byte, _rewardsReceiver common.Address, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry, _blsRegistrationSignature BLSAuthLibrarySignature) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.RegisterAsAllowedOperator(&_ContractAvsGovernance.TransactOpts, _blsKey, _authToken, _rewardsReceiver, _operatorSignature, _blsRegistrationSignature)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0x22609a4d.
//
// Solidity: function registerAsOperator(uint256[4] _blsKey, address _rewardsReceiver, (bytes,bytes32,uint256) _operatorSignature, (uint256[2]) _blsRegistrationSignature) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) RegisterAsOperator(opts *bind.TransactOpts, _blsKey [4]*big.Int, _rewardsReceiver common.Address, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry, _blsRegistrationSignature BLSAuthLibrarySignature) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "registerAsOperator", _blsKey, _rewardsReceiver, _operatorSignature, _blsRegistrationSignature)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0x22609a4d.
//
// Solidity: function registerAsOperator(uint256[4] _blsKey, address _rewardsReceiver, (bytes,bytes32,uint256) _operatorSignature, (uint256[2]) _blsRegistrationSignature) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) RegisterAsOperator(_blsKey [4]*big.Int, _rewardsReceiver common.Address, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry, _blsRegistrationSignature BLSAuthLibrarySignature) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.RegisterAsOperator(&_ContractAvsGovernance.TransactOpts, _blsKey, _rewardsReceiver, _operatorSignature, _blsRegistrationSignature)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0x22609a4d.
//
// Solidity: function registerAsOperator(uint256[4] _blsKey, address _rewardsReceiver, (bytes,bytes32,uint256) _operatorSignature, (uint256[2]) _blsRegistrationSignature) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) RegisterAsOperator(_blsKey [4]*big.Int, _rewardsReceiver common.Address, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry, _blsRegistrationSignature BLSAuthLibrarySignature) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.RegisterAsOperator(&_ContractAvsGovernance.TransactOpts, _blsKey, _rewardsReceiver, _operatorSignature, _blsRegistrationSignature)
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

// SetAllowlistSigner is a paid mutator transaction binding the contract method 0xe474def4.
//
// Solidity: function setAllowlistSigner(address _allowlistSigner) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) SetAllowlistSigner(opts *bind.TransactOpts, _allowlistSigner common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "setAllowlistSigner", _allowlistSigner)
}

// SetAllowlistSigner is a paid mutator transaction binding the contract method 0xe474def4.
//
// Solidity: function setAllowlistSigner(address _allowlistSigner) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) SetAllowlistSigner(_allowlistSigner common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetAllowlistSigner(&_ContractAvsGovernance.TransactOpts, _allowlistSigner)
}

// SetAllowlistSigner is a paid mutator transaction binding the contract method 0xe474def4.
//
// Solidity: function setAllowlistSigner(address _allowlistSigner) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) SetAllowlistSigner(_allowlistSigner common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetAllowlistSigner(&_ContractAvsGovernance.TransactOpts, _allowlistSigner)
}

// SetAvsGovernanceLogic is a paid mutator transaction binding the contract method 0x8987c767.
//
// Solidity: function setAvsGovernanceLogic(address _avsGovernanceLogic) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) SetAvsGovernanceLogic(opts *bind.TransactOpts, _avsGovernanceLogic common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "setAvsGovernanceLogic", _avsGovernanceLogic)
}

// SetAvsGovernanceLogic is a paid mutator transaction binding the contract method 0x8987c767.
//
// Solidity: function setAvsGovernanceLogic(address _avsGovernanceLogic) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) SetAvsGovernanceLogic(_avsGovernanceLogic common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetAvsGovernanceLogic(&_ContractAvsGovernance.TransactOpts, _avsGovernanceLogic)
}

// SetAvsGovernanceLogic is a paid mutator transaction binding the contract method 0x8987c767.
//
// Solidity: function setAvsGovernanceLogic(address _avsGovernanceLogic) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) SetAvsGovernanceLogic(_avsGovernanceLogic common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetAvsGovernanceLogic(&_ContractAvsGovernance.TransactOpts, _avsGovernanceLogic)
}

// SetAvsGovernanceMultiplierSyncer is a paid mutator transaction binding the contract method 0x3425e8d8.
//
// Solidity: function setAvsGovernanceMultiplierSyncer(address _newAvsGovernanceMultiplierSyncer) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) SetAvsGovernanceMultiplierSyncer(opts *bind.TransactOpts, _newAvsGovernanceMultiplierSyncer common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "setAvsGovernanceMultiplierSyncer", _newAvsGovernanceMultiplierSyncer)
}

// SetAvsGovernanceMultiplierSyncer is a paid mutator transaction binding the contract method 0x3425e8d8.
//
// Solidity: function setAvsGovernanceMultiplierSyncer(address _newAvsGovernanceMultiplierSyncer) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) SetAvsGovernanceMultiplierSyncer(_newAvsGovernanceMultiplierSyncer common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetAvsGovernanceMultiplierSyncer(&_ContractAvsGovernance.TransactOpts, _newAvsGovernanceMultiplierSyncer)
}

// SetAvsGovernanceMultiplierSyncer is a paid mutator transaction binding the contract method 0x3425e8d8.
//
// Solidity: function setAvsGovernanceMultiplierSyncer(address _newAvsGovernanceMultiplierSyncer) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) SetAvsGovernanceMultiplierSyncer(_newAvsGovernanceMultiplierSyncer common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetAvsGovernanceMultiplierSyncer(&_ContractAvsGovernance.TransactOpts, _newAvsGovernanceMultiplierSyncer)
}

// SetAvsName is a paid mutator transaction binding the contract method 0x7d38e926.
//
// Solidity: function setAvsName(string _avsName) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) SetAvsName(opts *bind.TransactOpts, _avsName string) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "setAvsName", _avsName)
}

// SetAvsName is a paid mutator transaction binding the contract method 0x7d38e926.
//
// Solidity: function setAvsName(string _avsName) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) SetAvsName(_avsName string) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetAvsName(&_ContractAvsGovernance.TransactOpts, _avsName)
}

// SetAvsName is a paid mutator transaction binding the contract method 0x7d38e926.
//
// Solidity: function setAvsName(string _avsName) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) SetAvsName(_avsName string) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetAvsName(&_ContractAvsGovernance.TransactOpts, _avsName)
}

// SetBLSAuthSingleton is a paid mutator transaction binding the contract method 0x4ef1476e.
//
// Solidity: function setBLSAuthSingleton(address _blsAuthSingleton) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) SetBLSAuthSingleton(opts *bind.TransactOpts, _blsAuthSingleton common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "setBLSAuthSingleton", _blsAuthSingleton)
}

// SetBLSAuthSingleton is a paid mutator transaction binding the contract method 0x4ef1476e.
//
// Solidity: function setBLSAuthSingleton(address _blsAuthSingleton) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) SetBLSAuthSingleton(_blsAuthSingleton common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetBLSAuthSingleton(&_ContractAvsGovernance.TransactOpts, _blsAuthSingleton)
}

// SetBLSAuthSingleton is a paid mutator transaction binding the contract method 0x4ef1476e.
//
// Solidity: function setBLSAuthSingleton(address _blsAuthSingleton) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) SetBLSAuthSingleton(_blsAuthSingleton common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetBLSAuthSingleton(&_ContractAvsGovernance.TransactOpts, _blsAuthSingleton)
}

// SetIsAllowlisted is a paid mutator transaction binding the contract method 0x9e965cc1.
//
// Solidity: function setIsAllowlisted(bool _isAllowlisted) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) SetIsAllowlisted(opts *bind.TransactOpts, _isAllowlisted bool) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "setIsAllowlisted", _isAllowlisted)
}

// SetIsAllowlisted is a paid mutator transaction binding the contract method 0x9e965cc1.
//
// Solidity: function setIsAllowlisted(bool _isAllowlisted) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) SetIsAllowlisted(_isAllowlisted bool) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetIsAllowlisted(&_ContractAvsGovernance.TransactOpts, _isAllowlisted)
}

// SetIsAllowlisted is a paid mutator transaction binding the contract method 0x9e965cc1.
//
// Solidity: function setIsAllowlisted(bool _isAllowlisted) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) SetIsAllowlisted(_isAllowlisted bool) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetIsAllowlisted(&_ContractAvsGovernance.TransactOpts, _isAllowlisted)
}

// SetMaxEffectiveBalance is a paid mutator transaction binding the contract method 0x76086c70.
//
// Solidity: function setMaxEffectiveBalance(uint256 _maxBalance) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) SetMaxEffectiveBalance(opts *bind.TransactOpts, _maxBalance *big.Int) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "setMaxEffectiveBalance", _maxBalance)
}

// SetMaxEffectiveBalance is a paid mutator transaction binding the contract method 0x76086c70.
//
// Solidity: function setMaxEffectiveBalance(uint256 _maxBalance) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) SetMaxEffectiveBalance(_maxBalance *big.Int) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetMaxEffectiveBalance(&_ContractAvsGovernance.TransactOpts, _maxBalance)
}

// SetMaxEffectiveBalance is a paid mutator transaction binding the contract method 0x76086c70.
//
// Solidity: function setMaxEffectiveBalance(uint256 _maxBalance) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) SetMaxEffectiveBalance(_maxBalance *big.Int) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetMaxEffectiveBalance(&_ContractAvsGovernance.TransactOpts, _maxBalance)
}

// SetMinSharesForStrategy is a paid mutator transaction binding the contract method 0x305df58a.
//
// Solidity: function setMinSharesForStrategy(address _strategy, uint256 _minShares) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) SetMinSharesForStrategy(opts *bind.TransactOpts, _strategy common.Address, _minShares *big.Int) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "setMinSharesForStrategy", _strategy, _minShares)
}

// SetMinSharesForStrategy is a paid mutator transaction binding the contract method 0x305df58a.
//
// Solidity: function setMinSharesForStrategy(address _strategy, uint256 _minShares) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) SetMinSharesForStrategy(_strategy common.Address, _minShares *big.Int) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetMinSharesForStrategy(&_ContractAvsGovernance.TransactOpts, _strategy, _minShares)
}

// SetMinSharesForStrategy is a paid mutator transaction binding the contract method 0x305df58a.
//
// Solidity: function setMinSharesForStrategy(address _strategy, uint256 _minShares) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) SetMinSharesForStrategy(_strategy common.Address, _minShares *big.Int) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetMinSharesForStrategy(&_ContractAvsGovernance.TransactOpts, _strategy, _minShares)
}

// SetMinVotingPower is a paid mutator transaction binding the contract method 0x55e48918.
//
// Solidity: function setMinVotingPower(uint256 _minVotingPower) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) SetMinVotingPower(opts *bind.TransactOpts, _minVotingPower *big.Int) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "setMinVotingPower", _minVotingPower)
}

// SetMinVotingPower is a paid mutator transaction binding the contract method 0x55e48918.
//
// Solidity: function setMinVotingPower(uint256 _minVotingPower) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) SetMinVotingPower(_minVotingPower *big.Int) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetMinVotingPower(&_ContractAvsGovernance.TransactOpts, _minVotingPower)
}

// SetMinVotingPower is a paid mutator transaction binding the contract method 0x55e48918.
//
// Solidity: function setMinVotingPower(uint256 _minVotingPower) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) SetMinVotingPower(_minVotingPower *big.Int) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetMinVotingPower(&_ContractAvsGovernance.TransactOpts, _minVotingPower)
}

// SetNumOfOperatorsLimit is a paid mutator transaction binding the contract method 0x9d79e4a7.
//
// Solidity: function setNumOfOperatorsLimit(uint256 _newLimitOfNumOfOperators) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) SetNumOfOperatorsLimit(opts *bind.TransactOpts, _newLimitOfNumOfOperators *big.Int) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "setNumOfOperatorsLimit", _newLimitOfNumOfOperators)
}

// SetNumOfOperatorsLimit is a paid mutator transaction binding the contract method 0x9d79e4a7.
//
// Solidity: function setNumOfOperatorsLimit(uint256 _newLimitOfNumOfOperators) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) SetNumOfOperatorsLimit(_newLimitOfNumOfOperators *big.Int) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetNumOfOperatorsLimit(&_ContractAvsGovernance.TransactOpts, _newLimitOfNumOfOperators)
}

// SetNumOfOperatorsLimit is a paid mutator transaction binding the contract method 0x9d79e4a7.
//
// Solidity: function setNumOfOperatorsLimit(uint256 _newLimitOfNumOfOperators) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) SetNumOfOperatorsLimit(_newLimitOfNumOfOperators *big.Int) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetNumOfOperatorsLimit(&_ContractAvsGovernance.TransactOpts, _newLimitOfNumOfOperators)
}

// SetOthenticRegistry is a paid mutator transaction binding the contract method 0x45a022fa.
//
// Solidity: function setOthenticRegistry(address _othenticRegistry) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) SetOthenticRegistry(opts *bind.TransactOpts, _othenticRegistry common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "setOthenticRegistry", _othenticRegistry)
}

// SetOthenticRegistry is a paid mutator transaction binding the contract method 0x45a022fa.
//
// Solidity: function setOthenticRegistry(address _othenticRegistry) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) SetOthenticRegistry(_othenticRegistry common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetOthenticRegistry(&_ContractAvsGovernance.TransactOpts, _othenticRegistry)
}

// SetOthenticRegistry is a paid mutator transaction binding the contract method 0x45a022fa.
//
// Solidity: function setOthenticRegistry(address _othenticRegistry) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) SetOthenticRegistry(_othenticRegistry common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetOthenticRegistry(&_ContractAvsGovernance.TransactOpts, _othenticRegistry)
}

// SetRewardsReceiverModificationDelay is a paid mutator transaction binding the contract method 0x8a70469a.
//
// Solidity: function setRewardsReceiverModificationDelay(uint256 _rewardsReceiverModificationDelay) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) SetRewardsReceiverModificationDelay(opts *bind.TransactOpts, _rewardsReceiverModificationDelay *big.Int) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "setRewardsReceiverModificationDelay", _rewardsReceiverModificationDelay)
}

// SetRewardsReceiverModificationDelay is a paid mutator transaction binding the contract method 0x8a70469a.
//
// Solidity: function setRewardsReceiverModificationDelay(uint256 _rewardsReceiverModificationDelay) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) SetRewardsReceiverModificationDelay(_rewardsReceiverModificationDelay *big.Int) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetRewardsReceiverModificationDelay(&_ContractAvsGovernance.TransactOpts, _rewardsReceiverModificationDelay)
}

// SetRewardsReceiverModificationDelay is a paid mutator transaction binding the contract method 0x8a70469a.
//
// Solidity: function setRewardsReceiverModificationDelay(uint256 _rewardsReceiverModificationDelay) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) SetRewardsReceiverModificationDelay(_rewardsReceiverModificationDelay *big.Int) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetRewardsReceiverModificationDelay(&_ContractAvsGovernance.TransactOpts, _rewardsReceiverModificationDelay)
}

// SetStrategyMultiplier is a paid mutator transaction binding the contract method 0x076400d5.
//
// Solidity: function setStrategyMultiplier((address,uint256) _strategyMultiplier) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) SetStrategyMultiplier(opts *bind.TransactOpts, _strategyMultiplier IAvsGovernanceStrategyMultiplier) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "setStrategyMultiplier", _strategyMultiplier)
}

// SetStrategyMultiplier is a paid mutator transaction binding the contract method 0x076400d5.
//
// Solidity: function setStrategyMultiplier((address,uint256) _strategyMultiplier) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) SetStrategyMultiplier(_strategyMultiplier IAvsGovernanceStrategyMultiplier) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetStrategyMultiplier(&_ContractAvsGovernance.TransactOpts, _strategyMultiplier)
}

// SetStrategyMultiplier is a paid mutator transaction binding the contract method 0x076400d5.
//
// Solidity: function setStrategyMultiplier((address,uint256) _strategyMultiplier) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) SetStrategyMultiplier(_strategyMultiplier IAvsGovernanceStrategyMultiplier) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetStrategyMultiplier(&_ContractAvsGovernance.TransactOpts, _strategyMultiplier)
}

// SetStrategyMultiplierBatch is a paid mutator transaction binding the contract method 0xd94a2e1d.
//
// Solidity: function setStrategyMultiplierBatch((address,uint256)[] _strategyMultipliers) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) SetStrategyMultiplierBatch(opts *bind.TransactOpts, _strategyMultipliers []IAvsGovernanceStrategyMultiplier) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "setStrategyMultiplierBatch", _strategyMultipliers)
}

// SetStrategyMultiplierBatch is a paid mutator transaction binding the contract method 0xd94a2e1d.
//
// Solidity: function setStrategyMultiplierBatch((address,uint256)[] _strategyMultipliers) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) SetStrategyMultiplierBatch(_strategyMultipliers []IAvsGovernanceStrategyMultiplier) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetStrategyMultiplierBatch(&_ContractAvsGovernance.TransactOpts, _strategyMultipliers)
}

// SetStrategyMultiplierBatch is a paid mutator transaction binding the contract method 0xd94a2e1d.
//
// Solidity: function setStrategyMultiplierBatch((address,uint256)[] _strategyMultipliers) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) SetStrategyMultiplierBatch(_strategyMultipliers []IAvsGovernanceStrategyMultiplier) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetStrategyMultiplierBatch(&_ContractAvsGovernance.TransactOpts, _strategyMultipliers)
}

// SetSupportedStrategies is a paid mutator transaction binding the contract method 0x312c150b.
//
// Solidity: function setSupportedStrategies(address[] _strategies) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) SetSupportedStrategies(opts *bind.TransactOpts, _strategies []common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "setSupportedStrategies", _strategies)
}

// SetSupportedStrategies is a paid mutator transaction binding the contract method 0x312c150b.
//
// Solidity: function setSupportedStrategies(address[] _strategies) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) SetSupportedStrategies(_strategies []common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetSupportedStrategies(&_ContractAvsGovernance.TransactOpts, _strategies)
}

// SetSupportedStrategies is a paid mutator transaction binding the contract method 0x312c150b.
//
// Solidity: function setSupportedStrategies(address[] _strategies) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) SetSupportedStrategies(_strategies []common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetSupportedStrategies(&_ContractAvsGovernance.TransactOpts, _strategies)
}

// TransferAvsGovernanceMultisig is a paid mutator transaction binding the contract method 0x513c52ba.
//
// Solidity: function transferAvsGovernanceMultisig(address _newAvsGovernanceMultisig) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) TransferAvsGovernanceMultisig(opts *bind.TransactOpts, _newAvsGovernanceMultisig common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "transferAvsGovernanceMultisig", _newAvsGovernanceMultisig)
}

// TransferAvsGovernanceMultisig is a paid mutator transaction binding the contract method 0x513c52ba.
//
// Solidity: function transferAvsGovernanceMultisig(address _newAvsGovernanceMultisig) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) TransferAvsGovernanceMultisig(_newAvsGovernanceMultisig common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.TransferAvsGovernanceMultisig(&_ContractAvsGovernance.TransactOpts, _newAvsGovernanceMultisig)
}

// TransferAvsGovernanceMultisig is a paid mutator transaction binding the contract method 0x513c52ba.
//
// Solidity: function transferAvsGovernanceMultisig(address _newAvsGovernanceMultisig) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) TransferAvsGovernanceMultisig(_newAvsGovernanceMultisig common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.TransferAvsGovernanceMultisig(&_ContractAvsGovernance.TransactOpts, _newAvsGovernanceMultisig)
}

// TransferMessageHandler is a paid mutator transaction binding the contract method 0x4d07f651.
//
// Solidity: function transferMessageHandler(address _newMessageHandler) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) TransferMessageHandler(opts *bind.TransactOpts, _newMessageHandler common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "transferMessageHandler", _newMessageHandler)
}

// TransferMessageHandler is a paid mutator transaction binding the contract method 0x4d07f651.
//
// Solidity: function transferMessageHandler(address _newMessageHandler) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) TransferMessageHandler(_newMessageHandler common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.TransferMessageHandler(&_ContractAvsGovernance.TransactOpts, _newMessageHandler)
}

// TransferMessageHandler is a paid mutator transaction binding the contract method 0x4d07f651.
//
// Solidity: function transferMessageHandler(address _newMessageHandler) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) TransferMessageHandler(_newMessageHandler common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.TransferMessageHandler(&_ContractAvsGovernance.TransactOpts, _newMessageHandler)
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

// UnregisterAsOperator is a paid mutator transaction binding the contract method 0x09869442.
//
// Solidity: function unregisterAsOperator() returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) UnregisterAsOperator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "unregisterAsOperator")
}

// UnregisterAsOperator is a paid mutator transaction binding the contract method 0x09869442.
//
// Solidity: function unregisterAsOperator() returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) UnregisterAsOperator() (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.UnregisterAsOperator(&_ContractAvsGovernance.TransactOpts)
}

// UnregisterAsOperator is a paid mutator transaction binding the contract method 0x09869442.
//
// Solidity: function unregisterAsOperator() returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) UnregisterAsOperator() (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.UnregisterAsOperator(&_ContractAvsGovernance.TransactOpts)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string metadataURI) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, metadataURI string) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "updateAVSMetadataURI", metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string metadataURI) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) UpdateAVSMetadataURI(metadataURI string) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.UpdateAVSMetadataURI(&_ContractAvsGovernance.TransactOpts, metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string metadataURI) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) UpdateAVSMetadataURI(metadataURI string) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.UpdateAVSMetadataURI(&_ContractAvsGovernance.TransactOpts, metadataURI)
}

// WithdrawBatchRewards is a paid mutator transaction binding the contract method 0xbc8be0c8.
//
// Solidity: function withdrawBatchRewards((address,uint256)[] _operators, uint256 _lastPayedTask) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) WithdrawBatchRewards(opts *bind.TransactOpts, _operators []IAvsGovernancePaymentRequestMessage, _lastPayedTask *big.Int) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "withdrawBatchRewards", _operators, _lastPayedTask)
}

// WithdrawBatchRewards is a paid mutator transaction binding the contract method 0xbc8be0c8.
//
// Solidity: function withdrawBatchRewards((address,uint256)[] _operators, uint256 _lastPayedTask) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) WithdrawBatchRewards(_operators []IAvsGovernancePaymentRequestMessage, _lastPayedTask *big.Int) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.WithdrawBatchRewards(&_ContractAvsGovernance.TransactOpts, _operators, _lastPayedTask)
}

// WithdrawBatchRewards is a paid mutator transaction binding the contract method 0xbc8be0c8.
//
// Solidity: function withdrawBatchRewards((address,uint256)[] _operators, uint256 _lastPayedTask) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) WithdrawBatchRewards(_operators []IAvsGovernancePaymentRequestMessage, _lastPayedTask *big.Int) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.WithdrawBatchRewards(&_ContractAvsGovernance.TransactOpts, _operators, _lastPayedTask)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0x3256b4d1.
//
// Solidity: function withdrawRewards(address _operator, uint256 _lastPayedTask, uint256 _feeToClaim) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) WithdrawRewards(opts *bind.TransactOpts, _operator common.Address, _lastPayedTask *big.Int, _feeToClaim *big.Int) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "withdrawRewards", _operator, _lastPayedTask, _feeToClaim)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0x3256b4d1.
//
// Solidity: function withdrawRewards(address _operator, uint256 _lastPayedTask, uint256 _feeToClaim) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) WithdrawRewards(_operator common.Address, _lastPayedTask *big.Int, _feeToClaim *big.Int) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.WithdrawRewards(&_ContractAvsGovernance.TransactOpts, _operator, _lastPayedTask, _feeToClaim)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0x3256b4d1.
//
// Solidity: function withdrawRewards(address _operator, uint256 _lastPayedTask, uint256 _feeToClaim) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) WithdrawRewards(_operator common.Address, _lastPayedTask *big.Int, _feeToClaim *big.Int) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.WithdrawRewards(&_ContractAvsGovernance.TransactOpts, _operator, _lastPayedTask, _feeToClaim)
}

// ContractAvsGovernanceBLSAuthSingletonSetIterator is returned from FilterBLSAuthSingletonSet and is used to iterate over the raw logs and unpacked data for BLSAuthSingletonSet events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceBLSAuthSingletonSetIterator struct {
	Event *ContractAvsGovernanceBLSAuthSingletonSet // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceBLSAuthSingletonSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceBLSAuthSingletonSet)
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
		it.Event = new(ContractAvsGovernanceBLSAuthSingletonSet)
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
func (it *ContractAvsGovernanceBLSAuthSingletonSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceBLSAuthSingletonSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceBLSAuthSingletonSet represents a BLSAuthSingletonSet event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceBLSAuthSingletonSet struct {
	BlsAuthSingleton common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterBLSAuthSingletonSet is a free log retrieval operation binding the contract event 0x4cbffdecf3b5e4b22bfb2bdec99a66f8fcf81e19b060682afd9645c729da1472.
//
// Solidity: event BLSAuthSingletonSet(address blsAuthSingleton)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterBLSAuthSingletonSet(opts *bind.FilterOpts) (*ContractAvsGovernanceBLSAuthSingletonSetIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "BLSAuthSingletonSet")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceBLSAuthSingletonSetIterator{contract: _ContractAvsGovernance.contract, event: "BLSAuthSingletonSet", logs: logs, sub: sub}, nil
}

// WatchBLSAuthSingletonSet is a free log subscription operation binding the contract event 0x4cbffdecf3b5e4b22bfb2bdec99a66f8fcf81e19b060682afd9645c729da1472.
//
// Solidity: event BLSAuthSingletonSet(address blsAuthSingleton)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchBLSAuthSingletonSet(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceBLSAuthSingletonSet) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "BLSAuthSingletonSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceBLSAuthSingletonSet)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "BLSAuthSingletonSet", log); err != nil {
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

// ParseBLSAuthSingletonSet is a log parse operation binding the contract event 0x4cbffdecf3b5e4b22bfb2bdec99a66f8fcf81e19b060682afd9645c729da1472.
//
// Solidity: event BLSAuthSingletonSet(address blsAuthSingleton)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseBLSAuthSingletonSet(log types.Log) (*ContractAvsGovernanceBLSAuthSingletonSet, error) {
	event := new(ContractAvsGovernanceBLSAuthSingletonSet)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "BLSAuthSingletonSet", log); err != nil {
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

// ContractAvsGovernanceMaxEffectiveBalanceSetIterator is returned from FilterMaxEffectiveBalanceSet and is used to iterate over the raw logs and unpacked data for MaxEffectiveBalanceSet events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceMaxEffectiveBalanceSetIterator struct {
	Event *ContractAvsGovernanceMaxEffectiveBalanceSet // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceMaxEffectiveBalanceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceMaxEffectiveBalanceSet)
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
		it.Event = new(ContractAvsGovernanceMaxEffectiveBalanceSet)
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
func (it *ContractAvsGovernanceMaxEffectiveBalanceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceMaxEffectiveBalanceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceMaxEffectiveBalanceSet represents a MaxEffectiveBalanceSet event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceMaxEffectiveBalanceSet struct {
	MaxEffectiveBalance *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterMaxEffectiveBalanceSet is a free log retrieval operation binding the contract event 0x00c6fb6db9c52d89a1eaf84e0470a3304db2086d0ac44d64ebf4ea35a905a7d0.
//
// Solidity: event MaxEffectiveBalanceSet(uint256 maxEffectiveBalance)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterMaxEffectiveBalanceSet(opts *bind.FilterOpts) (*ContractAvsGovernanceMaxEffectiveBalanceSetIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "MaxEffectiveBalanceSet")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceMaxEffectiveBalanceSetIterator{contract: _ContractAvsGovernance.contract, event: "MaxEffectiveBalanceSet", logs: logs, sub: sub}, nil
}

// WatchMaxEffectiveBalanceSet is a free log subscription operation binding the contract event 0x00c6fb6db9c52d89a1eaf84e0470a3304db2086d0ac44d64ebf4ea35a905a7d0.
//
// Solidity: event MaxEffectiveBalanceSet(uint256 maxEffectiveBalance)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchMaxEffectiveBalanceSet(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceMaxEffectiveBalanceSet) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "MaxEffectiveBalanceSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceMaxEffectiveBalanceSet)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "MaxEffectiveBalanceSet", log); err != nil {
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

// ParseMaxEffectiveBalanceSet is a log parse operation binding the contract event 0x00c6fb6db9c52d89a1eaf84e0470a3304db2086d0ac44d64ebf4ea35a905a7d0.
//
// Solidity: event MaxEffectiveBalanceSet(uint256 maxEffectiveBalance)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseMaxEffectiveBalanceSet(log types.Log) (*ContractAvsGovernanceMaxEffectiveBalanceSet, error) {
	event := new(ContractAvsGovernanceMaxEffectiveBalanceSet)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "MaxEffectiveBalanceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceMinSharesPerStrategySetIterator is returned from FilterMinSharesPerStrategySet and is used to iterate over the raw logs and unpacked data for MinSharesPerStrategySet events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceMinSharesPerStrategySetIterator struct {
	Event *ContractAvsGovernanceMinSharesPerStrategySet // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceMinSharesPerStrategySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceMinSharesPerStrategySet)
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
		it.Event = new(ContractAvsGovernanceMinSharesPerStrategySet)
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
func (it *ContractAvsGovernanceMinSharesPerStrategySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceMinSharesPerStrategySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceMinSharesPerStrategySet represents a MinSharesPerStrategySet event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceMinSharesPerStrategySet struct {
	Strategy  common.Address
	MinShares *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMinSharesPerStrategySet is a free log retrieval operation binding the contract event 0x3a6c52328a7b3b726d0ec757d68f416b26ec2991ac4d4f95d450c504f5a0e521.
//
// Solidity: event MinSharesPerStrategySet(address strategy, uint256 minShares)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterMinSharesPerStrategySet(opts *bind.FilterOpts) (*ContractAvsGovernanceMinSharesPerStrategySetIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "MinSharesPerStrategySet")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceMinSharesPerStrategySetIterator{contract: _ContractAvsGovernance.contract, event: "MinSharesPerStrategySet", logs: logs, sub: sub}, nil
}

// WatchMinSharesPerStrategySet is a free log subscription operation binding the contract event 0x3a6c52328a7b3b726d0ec757d68f416b26ec2991ac4d4f95d450c504f5a0e521.
//
// Solidity: event MinSharesPerStrategySet(address strategy, uint256 minShares)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchMinSharesPerStrategySet(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceMinSharesPerStrategySet) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "MinSharesPerStrategySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceMinSharesPerStrategySet)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "MinSharesPerStrategySet", log); err != nil {
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

// ParseMinSharesPerStrategySet is a log parse operation binding the contract event 0x3a6c52328a7b3b726d0ec757d68f416b26ec2991ac4d4f95d450c504f5a0e521.
//
// Solidity: event MinSharesPerStrategySet(address strategy, uint256 minShares)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseMinSharesPerStrategySet(log types.Log) (*ContractAvsGovernanceMinSharesPerStrategySet, error) {
	event := new(ContractAvsGovernanceMinSharesPerStrategySet)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "MinSharesPerStrategySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceMinVotingPowerSetIterator is returned from FilterMinVotingPowerSet and is used to iterate over the raw logs and unpacked data for MinVotingPowerSet events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceMinVotingPowerSetIterator struct {
	Event *ContractAvsGovernanceMinVotingPowerSet // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceMinVotingPowerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceMinVotingPowerSet)
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
		it.Event = new(ContractAvsGovernanceMinVotingPowerSet)
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
func (it *ContractAvsGovernanceMinVotingPowerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceMinVotingPowerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceMinVotingPowerSet represents a MinVotingPowerSet event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceMinVotingPowerSet struct {
	MinVotingPower *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMinVotingPowerSet is a free log retrieval operation binding the contract event 0x10203ddc048c86cf14172a6ea2565c805ce7320b22d6941b2eb396d0ee077983.
//
// Solidity: event MinVotingPowerSet(uint256 minVotingPower)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterMinVotingPowerSet(opts *bind.FilterOpts) (*ContractAvsGovernanceMinVotingPowerSetIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "MinVotingPowerSet")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceMinVotingPowerSetIterator{contract: _ContractAvsGovernance.contract, event: "MinVotingPowerSet", logs: logs, sub: sub}, nil
}

// WatchMinVotingPowerSet is a free log subscription operation binding the contract event 0x10203ddc048c86cf14172a6ea2565c805ce7320b22d6941b2eb396d0ee077983.
//
// Solidity: event MinVotingPowerSet(uint256 minVotingPower)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchMinVotingPowerSet(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceMinVotingPowerSet) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "MinVotingPowerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceMinVotingPowerSet)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "MinVotingPowerSet", log); err != nil {
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

// ParseMinVotingPowerSet is a log parse operation binding the contract event 0x10203ddc048c86cf14172a6ea2565c805ce7320b22d6941b2eb396d0ee077983.
//
// Solidity: event MinVotingPowerSet(uint256 minVotingPower)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseMinVotingPowerSet(log types.Log) (*ContractAvsGovernanceMinVotingPowerSet, error) {
	event := new(ContractAvsGovernanceMinVotingPowerSet)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "MinVotingPowerSet", log); err != nil {
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
// Solidity: event OperatorUnregistered(address operator)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterOperatorUnregistered(opts *bind.FilterOpts) (*ContractAvsGovernanceOperatorUnregisteredIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "OperatorUnregistered")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceOperatorUnregisteredIterator{contract: _ContractAvsGovernance.contract, event: "OperatorUnregistered", logs: logs, sub: sub}, nil
}

// WatchOperatorUnregistered is a free log subscription operation binding the contract event 0x6f42117a557500c705ddf040a619d86f39101e6b74ac20d7b3e5943ba473fc7f.
//
// Solidity: event OperatorUnregistered(address operator)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchOperatorUnregistered(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceOperatorUnregistered) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "OperatorUnregistered")
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
// Solidity: event OperatorUnregistered(address operator)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseOperatorUnregistered(log types.Log) (*ContractAvsGovernanceOperatorUnregistered, error) {
	event := new(ContractAvsGovernanceOperatorUnregistered)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "OperatorUnregistered", log); err != nil {
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
// Solidity: event QueuedRewardsReceiverModification(address operator, address receiver, uint256 delay)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterQueuedRewardsReceiverModification(opts *bind.FilterOpts) (*ContractAvsGovernanceQueuedRewardsReceiverModificationIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "QueuedRewardsReceiverModification")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceQueuedRewardsReceiverModificationIterator{contract: _ContractAvsGovernance.contract, event: "QueuedRewardsReceiverModification", logs: logs, sub: sub}, nil
}

// WatchQueuedRewardsReceiverModification is a free log subscription operation binding the contract event 0x0d8cfa10a3087b28d3c226ad9a37314860e7c3c0505a25a39e3cdefb3118a98a.
//
// Solidity: event QueuedRewardsReceiverModification(address operator, address receiver, uint256 delay)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchQueuedRewardsReceiverModification(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceQueuedRewardsReceiverModification) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "QueuedRewardsReceiverModification")
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
// Solidity: event QueuedRewardsReceiverModification(address operator, address receiver, uint256 delay)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseQueuedRewardsReceiverModification(log types.Log) (*ContractAvsGovernanceQueuedRewardsReceiverModification, error) {
	event := new(ContractAvsGovernanceQueuedRewardsReceiverModification)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "QueuedRewardsReceiverModification", log); err != nil {
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

// ContractAvsGovernanceSetAllowlistSignerIterator is returned from FilterSetAllowlistSigner and is used to iterate over the raw logs and unpacked data for SetAllowlistSigner events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetAllowlistSignerIterator struct {
	Event *ContractAvsGovernanceSetAllowlistSigner // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceSetAllowlistSignerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceSetAllowlistSigner)
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
		it.Event = new(ContractAvsGovernanceSetAllowlistSigner)
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
func (it *ContractAvsGovernanceSetAllowlistSignerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceSetAllowlistSignerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceSetAllowlistSigner represents a SetAllowlistSigner event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetAllowlistSigner struct {
	AllowlistSigner common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetAllowlistSigner is a free log retrieval operation binding the contract event 0xfa4acc0aaeb2714e420e9c8339167ddef7bc66c0f94a0c5a7722de21dcb7508c.
//
// Solidity: event SetAllowlistSigner(address allowlistSigner)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterSetAllowlistSigner(opts *bind.FilterOpts) (*ContractAvsGovernanceSetAllowlistSignerIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "SetAllowlistSigner")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceSetAllowlistSignerIterator{contract: _ContractAvsGovernance.contract, event: "SetAllowlistSigner", logs: logs, sub: sub}, nil
}

// WatchSetAllowlistSigner is a free log subscription operation binding the contract event 0xfa4acc0aaeb2714e420e9c8339167ddef7bc66c0f94a0c5a7722de21dcb7508c.
//
// Solidity: event SetAllowlistSigner(address allowlistSigner)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchSetAllowlistSigner(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetAllowlistSigner) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "SetAllowlistSigner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceSetAllowlistSigner)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetAllowlistSigner", log); err != nil {
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

// ParseSetAllowlistSigner is a log parse operation binding the contract event 0xfa4acc0aaeb2714e420e9c8339167ddef7bc66c0f94a0c5a7722de21dcb7508c.
//
// Solidity: event SetAllowlistSigner(address allowlistSigner)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseSetAllowlistSigner(log types.Log) (*ContractAvsGovernanceSetAllowlistSigner, error) {
	event := new(ContractAvsGovernanceSetAllowlistSigner)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetAllowlistSigner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceSetAvsGovernanceLogicIterator is returned from FilterSetAvsGovernanceLogic and is used to iterate over the raw logs and unpacked data for SetAvsGovernanceLogic events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetAvsGovernanceLogicIterator struct {
	Event *ContractAvsGovernanceSetAvsGovernanceLogic // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceSetAvsGovernanceLogicIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceSetAvsGovernanceLogic)
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
		it.Event = new(ContractAvsGovernanceSetAvsGovernanceLogic)
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
func (it *ContractAvsGovernanceSetAvsGovernanceLogicIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceSetAvsGovernanceLogicIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceSetAvsGovernanceLogic represents a SetAvsGovernanceLogic event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetAvsGovernanceLogic struct {
	AvsGovernanceLogic common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSetAvsGovernanceLogic is a free log retrieval operation binding the contract event 0x7c36ee80df183e227956a9f387a48d26bbf4d2f1526410493d11126de5a8942c.
//
// Solidity: event SetAvsGovernanceLogic(address avsGovernanceLogic)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterSetAvsGovernanceLogic(opts *bind.FilterOpts) (*ContractAvsGovernanceSetAvsGovernanceLogicIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "SetAvsGovernanceLogic")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceSetAvsGovernanceLogicIterator{contract: _ContractAvsGovernance.contract, event: "SetAvsGovernanceLogic", logs: logs, sub: sub}, nil
}

// WatchSetAvsGovernanceLogic is a free log subscription operation binding the contract event 0x7c36ee80df183e227956a9f387a48d26bbf4d2f1526410493d11126de5a8942c.
//
// Solidity: event SetAvsGovernanceLogic(address avsGovernanceLogic)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchSetAvsGovernanceLogic(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetAvsGovernanceLogic) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "SetAvsGovernanceLogic")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceSetAvsGovernanceLogic)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetAvsGovernanceLogic", log); err != nil {
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

// ParseSetAvsGovernanceLogic is a log parse operation binding the contract event 0x7c36ee80df183e227956a9f387a48d26bbf4d2f1526410493d11126de5a8942c.
//
// Solidity: event SetAvsGovernanceLogic(address avsGovernanceLogic)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseSetAvsGovernanceLogic(log types.Log) (*ContractAvsGovernanceSetAvsGovernanceLogic, error) {
	event := new(ContractAvsGovernanceSetAvsGovernanceLogic)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetAvsGovernanceLogic", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceSetAvsGovernanceMultiplierSyncerIterator is returned from FilterSetAvsGovernanceMultiplierSyncer and is used to iterate over the raw logs and unpacked data for SetAvsGovernanceMultiplierSyncer events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetAvsGovernanceMultiplierSyncerIterator struct {
	Event *ContractAvsGovernanceSetAvsGovernanceMultiplierSyncer // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceSetAvsGovernanceMultiplierSyncerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceSetAvsGovernanceMultiplierSyncer)
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
		it.Event = new(ContractAvsGovernanceSetAvsGovernanceMultiplierSyncer)
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
func (it *ContractAvsGovernanceSetAvsGovernanceMultiplierSyncerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceSetAvsGovernanceMultiplierSyncerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceSetAvsGovernanceMultiplierSyncer represents a SetAvsGovernanceMultiplierSyncer event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetAvsGovernanceMultiplierSyncer struct {
	AvsGovernanceMultiplierSyncer common.Address
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterSetAvsGovernanceMultiplierSyncer is a free log retrieval operation binding the contract event 0xb73a70f24733a9265231de5807eae76d1740a9974b31a142ef9e243508987bbe.
//
// Solidity: event SetAvsGovernanceMultiplierSyncer(address avsGovernanceMultiplierSyncer)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterSetAvsGovernanceMultiplierSyncer(opts *bind.FilterOpts) (*ContractAvsGovernanceSetAvsGovernanceMultiplierSyncerIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "SetAvsGovernanceMultiplierSyncer")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceSetAvsGovernanceMultiplierSyncerIterator{contract: _ContractAvsGovernance.contract, event: "SetAvsGovernanceMultiplierSyncer", logs: logs, sub: sub}, nil
}

// WatchSetAvsGovernanceMultiplierSyncer is a free log subscription operation binding the contract event 0xb73a70f24733a9265231de5807eae76d1740a9974b31a142ef9e243508987bbe.
//
// Solidity: event SetAvsGovernanceMultiplierSyncer(address avsGovernanceMultiplierSyncer)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchSetAvsGovernanceMultiplierSyncer(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetAvsGovernanceMultiplierSyncer) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "SetAvsGovernanceMultiplierSyncer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceSetAvsGovernanceMultiplierSyncer)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetAvsGovernanceMultiplierSyncer", log); err != nil {
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

// ParseSetAvsGovernanceMultiplierSyncer is a log parse operation binding the contract event 0xb73a70f24733a9265231de5807eae76d1740a9974b31a142ef9e243508987bbe.
//
// Solidity: event SetAvsGovernanceMultiplierSyncer(address avsGovernanceMultiplierSyncer)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseSetAvsGovernanceMultiplierSyncer(log types.Log) (*ContractAvsGovernanceSetAvsGovernanceMultiplierSyncer, error) {
	event := new(ContractAvsGovernanceSetAvsGovernanceMultiplierSyncer)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetAvsGovernanceMultiplierSyncer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceSetAvsGovernanceMultisigIterator is returned from FilterSetAvsGovernanceMultisig and is used to iterate over the raw logs and unpacked data for SetAvsGovernanceMultisig events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetAvsGovernanceMultisigIterator struct {
	Event *ContractAvsGovernanceSetAvsGovernanceMultisig // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceSetAvsGovernanceMultisigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceSetAvsGovernanceMultisig)
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
		it.Event = new(ContractAvsGovernanceSetAvsGovernanceMultisig)
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
func (it *ContractAvsGovernanceSetAvsGovernanceMultisigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceSetAvsGovernanceMultisigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceSetAvsGovernanceMultisig represents a SetAvsGovernanceMultisig event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetAvsGovernanceMultisig struct {
	NewAvsGovernanceMultisig common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSetAvsGovernanceMultisig is a free log retrieval operation binding the contract event 0x024e98b7d808a3ddb028252dc95dfdcb165a0ca59fcff8984b4fecf9a7222649.
//
// Solidity: event SetAvsGovernanceMultisig(address newAvsGovernanceMultisig)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterSetAvsGovernanceMultisig(opts *bind.FilterOpts) (*ContractAvsGovernanceSetAvsGovernanceMultisigIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "SetAvsGovernanceMultisig")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceSetAvsGovernanceMultisigIterator{contract: _ContractAvsGovernance.contract, event: "SetAvsGovernanceMultisig", logs: logs, sub: sub}, nil
}

// WatchSetAvsGovernanceMultisig is a free log subscription operation binding the contract event 0x024e98b7d808a3ddb028252dc95dfdcb165a0ca59fcff8984b4fecf9a7222649.
//
// Solidity: event SetAvsGovernanceMultisig(address newAvsGovernanceMultisig)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchSetAvsGovernanceMultisig(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetAvsGovernanceMultisig) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "SetAvsGovernanceMultisig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceSetAvsGovernanceMultisig)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetAvsGovernanceMultisig", log); err != nil {
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

// ParseSetAvsGovernanceMultisig is a log parse operation binding the contract event 0x024e98b7d808a3ddb028252dc95dfdcb165a0ca59fcff8984b4fecf9a7222649.
//
// Solidity: event SetAvsGovernanceMultisig(address newAvsGovernanceMultisig)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseSetAvsGovernanceMultisig(log types.Log) (*ContractAvsGovernanceSetAvsGovernanceMultisig, error) {
	event := new(ContractAvsGovernanceSetAvsGovernanceMultisig)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetAvsGovernanceMultisig", log); err != nil {
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

// ContractAvsGovernanceSetMessageHandlerIterator is returned from FilterSetMessageHandler and is used to iterate over the raw logs and unpacked data for SetMessageHandler events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetMessageHandlerIterator struct {
	Event *ContractAvsGovernanceSetMessageHandler // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceSetMessageHandlerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceSetMessageHandler)
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
		it.Event = new(ContractAvsGovernanceSetMessageHandler)
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
func (it *ContractAvsGovernanceSetMessageHandlerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceSetMessageHandlerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceSetMessageHandler represents a SetMessageHandler event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetMessageHandler struct {
	NewMessageHandler common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSetMessageHandler is a free log retrieval operation binding the contract event 0x997f84b541d7b68e210e6f50e3402b51d8411dbbc4d44ed81e508383126e4e94.
//
// Solidity: event SetMessageHandler(address newMessageHandler)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterSetMessageHandler(opts *bind.FilterOpts) (*ContractAvsGovernanceSetMessageHandlerIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "SetMessageHandler")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceSetMessageHandlerIterator{contract: _ContractAvsGovernance.contract, event: "SetMessageHandler", logs: logs, sub: sub}, nil
}

// WatchSetMessageHandler is a free log subscription operation binding the contract event 0x997f84b541d7b68e210e6f50e3402b51d8411dbbc4d44ed81e508383126e4e94.
//
// Solidity: event SetMessageHandler(address newMessageHandler)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchSetMessageHandler(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetMessageHandler) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "SetMessageHandler")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceSetMessageHandler)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetMessageHandler", log); err != nil {
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

// ParseSetMessageHandler is a log parse operation binding the contract event 0x997f84b541d7b68e210e6f50e3402b51d8411dbbc4d44ed81e508383126e4e94.
//
// Solidity: event SetMessageHandler(address newMessageHandler)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseSetMessageHandler(log types.Log) (*ContractAvsGovernanceSetMessageHandler, error) {
	event := new(ContractAvsGovernanceSetMessageHandler)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetMessageHandler", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceSetNumOfOperatorsLimitIterator is returned from FilterSetNumOfOperatorsLimit and is used to iterate over the raw logs and unpacked data for SetNumOfOperatorsLimit events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetNumOfOperatorsLimitIterator struct {
	Event *ContractAvsGovernanceSetNumOfOperatorsLimit // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceSetNumOfOperatorsLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceSetNumOfOperatorsLimit)
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
		it.Event = new(ContractAvsGovernanceSetNumOfOperatorsLimit)
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
func (it *ContractAvsGovernanceSetNumOfOperatorsLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceSetNumOfOperatorsLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceSetNumOfOperatorsLimit represents a SetNumOfOperatorsLimit event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetNumOfOperatorsLimit struct {
	NewLimitOfNumOfOperators *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSetNumOfOperatorsLimit is a free log retrieval operation binding the contract event 0xc0dd1d82df4ae12576f7a7912395305cf73deae26c764dd74a945cd6ba81591b.
//
// Solidity: event SetNumOfOperatorsLimit(uint256 newLimitOfNumOfOperators)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterSetNumOfOperatorsLimit(opts *bind.FilterOpts) (*ContractAvsGovernanceSetNumOfOperatorsLimitIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "SetNumOfOperatorsLimit")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceSetNumOfOperatorsLimitIterator{contract: _ContractAvsGovernance.contract, event: "SetNumOfOperatorsLimit", logs: logs, sub: sub}, nil
}

// WatchSetNumOfOperatorsLimit is a free log subscription operation binding the contract event 0xc0dd1d82df4ae12576f7a7912395305cf73deae26c764dd74a945cd6ba81591b.
//
// Solidity: event SetNumOfOperatorsLimit(uint256 newLimitOfNumOfOperators)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchSetNumOfOperatorsLimit(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetNumOfOperatorsLimit) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "SetNumOfOperatorsLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceSetNumOfOperatorsLimit)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetNumOfOperatorsLimit", log); err != nil {
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

// ParseSetNumOfOperatorsLimit is a log parse operation binding the contract event 0xc0dd1d82df4ae12576f7a7912395305cf73deae26c764dd74a945cd6ba81591b.
//
// Solidity: event SetNumOfOperatorsLimit(uint256 newLimitOfNumOfOperators)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseSetNumOfOperatorsLimit(log types.Log) (*ContractAvsGovernanceSetNumOfOperatorsLimit, error) {
	event := new(ContractAvsGovernanceSetNumOfOperatorsLimit)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetNumOfOperatorsLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceSetOthenticRegistryIterator is returned from FilterSetOthenticRegistry and is used to iterate over the raw logs and unpacked data for SetOthenticRegistry events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetOthenticRegistryIterator struct {
	Event *ContractAvsGovernanceSetOthenticRegistry // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceSetOthenticRegistryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceSetOthenticRegistry)
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
		it.Event = new(ContractAvsGovernanceSetOthenticRegistry)
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
func (it *ContractAvsGovernanceSetOthenticRegistryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceSetOthenticRegistryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceSetOthenticRegistry represents a SetOthenticRegistry event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetOthenticRegistry struct {
	OthenticRegistry common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetOthenticRegistry is a free log retrieval operation binding the contract event 0xf9855cc914fefc396bdeb5a4dcb97a2f6c75f4d6f00a8e71d6f9a40e474afe8d.
//
// Solidity: event SetOthenticRegistry(address othenticRegistry)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterSetOthenticRegistry(opts *bind.FilterOpts) (*ContractAvsGovernanceSetOthenticRegistryIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "SetOthenticRegistry")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceSetOthenticRegistryIterator{contract: _ContractAvsGovernance.contract, event: "SetOthenticRegistry", logs: logs, sub: sub}, nil
}

// WatchSetOthenticRegistry is a free log subscription operation binding the contract event 0xf9855cc914fefc396bdeb5a4dcb97a2f6c75f4d6f00a8e71d6f9a40e474afe8d.
//
// Solidity: event SetOthenticRegistry(address othenticRegistry)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchSetOthenticRegistry(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetOthenticRegistry) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "SetOthenticRegistry")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceSetOthenticRegistry)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetOthenticRegistry", log); err != nil {
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

// ParseSetOthenticRegistry is a log parse operation binding the contract event 0xf9855cc914fefc396bdeb5a4dcb97a2f6c75f4d6f00a8e71d6f9a40e474afe8d.
//
// Solidity: event SetOthenticRegistry(address othenticRegistry)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseSetOthenticRegistry(log types.Log) (*ContractAvsGovernanceSetOthenticRegistry, error) {
	event := new(ContractAvsGovernanceSetOthenticRegistry)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetOthenticRegistry", log); err != nil {
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
// Solidity: event SetRewardsReceiver(address operator, address receiver)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterSetRewardsReceiver(opts *bind.FilterOpts) (*ContractAvsGovernanceSetRewardsReceiverIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "SetRewardsReceiver")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceSetRewardsReceiverIterator{contract: _ContractAvsGovernance.contract, event: "SetRewardsReceiver", logs: logs, sub: sub}, nil
}

// WatchSetRewardsReceiver is a free log subscription operation binding the contract event 0xe906feea2ef60b474e22b4169bdd4de6906a84cd448cbcee99593526fe87082d.
//
// Solidity: event SetRewardsReceiver(address operator, address receiver)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchSetRewardsReceiver(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetRewardsReceiver) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "SetRewardsReceiver")
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
// Solidity: event SetRewardsReceiver(address operator, address receiver)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseSetRewardsReceiver(log types.Log) (*ContractAvsGovernanceSetRewardsReceiver, error) {
	event := new(ContractAvsGovernanceSetRewardsReceiver)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetRewardsReceiver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceSetRewardsReceiverModificationDelayIterator is returned from FilterSetRewardsReceiverModificationDelay and is used to iterate over the raw logs and unpacked data for SetRewardsReceiverModificationDelay events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetRewardsReceiverModificationDelayIterator struct {
	Event *ContractAvsGovernanceSetRewardsReceiverModificationDelay // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceSetRewardsReceiverModificationDelayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceSetRewardsReceiverModificationDelay)
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
		it.Event = new(ContractAvsGovernanceSetRewardsReceiverModificationDelay)
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
func (it *ContractAvsGovernanceSetRewardsReceiverModificationDelayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceSetRewardsReceiverModificationDelayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceSetRewardsReceiverModificationDelay represents a SetRewardsReceiverModificationDelay event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetRewardsReceiverModificationDelay struct {
	ModificationDelay *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSetRewardsReceiverModificationDelay is a free log retrieval operation binding the contract event 0x47c8c3268759fc47868c5e319217a2e85d47bd3935a4108debe246f6025fb88b.
//
// Solidity: event SetRewardsReceiverModificationDelay(uint256 modificationDelay)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterSetRewardsReceiverModificationDelay(opts *bind.FilterOpts) (*ContractAvsGovernanceSetRewardsReceiverModificationDelayIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "SetRewardsReceiverModificationDelay")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceSetRewardsReceiverModificationDelayIterator{contract: _ContractAvsGovernance.contract, event: "SetRewardsReceiverModificationDelay", logs: logs, sub: sub}, nil
}

// WatchSetRewardsReceiverModificationDelay is a free log subscription operation binding the contract event 0x47c8c3268759fc47868c5e319217a2e85d47bd3935a4108debe246f6025fb88b.
//
// Solidity: event SetRewardsReceiverModificationDelay(uint256 modificationDelay)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchSetRewardsReceiverModificationDelay(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetRewardsReceiverModificationDelay) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "SetRewardsReceiverModificationDelay")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceSetRewardsReceiverModificationDelay)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetRewardsReceiverModificationDelay", log); err != nil {
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

// ParseSetRewardsReceiverModificationDelay is a log parse operation binding the contract event 0x47c8c3268759fc47868c5e319217a2e85d47bd3935a4108debe246f6025fb88b.
//
// Solidity: event SetRewardsReceiverModificationDelay(uint256 modificationDelay)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseSetRewardsReceiverModificationDelay(log types.Log) (*ContractAvsGovernanceSetRewardsReceiverModificationDelay, error) {
	event := new(ContractAvsGovernanceSetRewardsReceiverModificationDelay)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetRewardsReceiverModificationDelay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceSetStrategyMultiplierIterator is returned from FilterSetStrategyMultiplier and is used to iterate over the raw logs and unpacked data for SetStrategyMultiplier events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetStrategyMultiplierIterator struct {
	Event *ContractAvsGovernanceSetStrategyMultiplier // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceSetStrategyMultiplierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceSetStrategyMultiplier)
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
		it.Event = new(ContractAvsGovernanceSetStrategyMultiplier)
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
func (it *ContractAvsGovernanceSetStrategyMultiplierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceSetStrategyMultiplierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceSetStrategyMultiplier represents a SetStrategyMultiplier event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetStrategyMultiplier struct {
	Strategy   common.Address
	Multiplier *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetStrategyMultiplier is a free log retrieval operation binding the contract event 0x8ae53ffd0ebc018acb19342fba690554d49ae9a467a9606a38b49cb5ad775c81.
//
// Solidity: event SetStrategyMultiplier(address strategy, uint256 multiplier)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterSetStrategyMultiplier(opts *bind.FilterOpts) (*ContractAvsGovernanceSetStrategyMultiplierIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "SetStrategyMultiplier")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceSetStrategyMultiplierIterator{contract: _ContractAvsGovernance.contract, event: "SetStrategyMultiplier", logs: logs, sub: sub}, nil
}

// WatchSetStrategyMultiplier is a free log subscription operation binding the contract event 0x8ae53ffd0ebc018acb19342fba690554d49ae9a467a9606a38b49cb5ad775c81.
//
// Solidity: event SetStrategyMultiplier(address strategy, uint256 multiplier)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchSetStrategyMultiplier(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetStrategyMultiplier) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "SetStrategyMultiplier")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceSetStrategyMultiplier)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetStrategyMultiplier", log); err != nil {
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

// ParseSetStrategyMultiplier is a log parse operation binding the contract event 0x8ae53ffd0ebc018acb19342fba690554d49ae9a467a9606a38b49cb5ad775c81.
//
// Solidity: event SetStrategyMultiplier(address strategy, uint256 multiplier)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseSetStrategyMultiplier(log types.Log) (*ContractAvsGovernanceSetStrategyMultiplier, error) {
	event := new(ContractAvsGovernanceSetStrategyMultiplier)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetStrategyMultiplier", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceSetSupportedStrategiesIterator is returned from FilterSetSupportedStrategies and is used to iterate over the raw logs and unpacked data for SetSupportedStrategies events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetSupportedStrategiesIterator struct {
	Event *ContractAvsGovernanceSetSupportedStrategies // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceSetSupportedStrategiesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceSetSupportedStrategies)
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
		it.Event = new(ContractAvsGovernanceSetSupportedStrategies)
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
func (it *ContractAvsGovernanceSetSupportedStrategiesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceSetSupportedStrategiesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceSetSupportedStrategies represents a SetSupportedStrategies event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetSupportedStrategies struct {
	Strategies []common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetSupportedStrategies is a free log retrieval operation binding the contract event 0xf009a6ffded424f714e8904d643a1ea4479453188faf08a3996121996b76684f.
//
// Solidity: event SetSupportedStrategies(address[] strategies)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterSetSupportedStrategies(opts *bind.FilterOpts) (*ContractAvsGovernanceSetSupportedStrategiesIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "SetSupportedStrategies")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceSetSupportedStrategiesIterator{contract: _ContractAvsGovernance.contract, event: "SetSupportedStrategies", logs: logs, sub: sub}, nil
}

// WatchSetSupportedStrategies is a free log subscription operation binding the contract event 0xf009a6ffded424f714e8904d643a1ea4479453188faf08a3996121996b76684f.
//
// Solidity: event SetSupportedStrategies(address[] strategies)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchSetSupportedStrategies(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetSupportedStrategies) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "SetSupportedStrategies")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceSetSupportedStrategies)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetSupportedStrategies", log); err != nil {
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

// ParseSetSupportedStrategies is a log parse operation binding the contract event 0xf009a6ffded424f714e8904d643a1ea4479453188faf08a3996121996b76684f.
//
// Solidity: event SetSupportedStrategies(address[] strategies)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseSetSupportedStrategies(log types.Log) (*ContractAvsGovernanceSetSupportedStrategies, error) {
	event := new(ContractAvsGovernanceSetSupportedStrategies)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetSupportedStrategies", log); err != nil {
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
