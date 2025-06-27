// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractAttestationCenter

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

// IAttestationCenterBlsTaskSubmissionDetails is an auto generated low-level Go binding around an user-defined struct.
type IAttestationCenterBlsTaskSubmissionDetails struct {
	IsApproved   bool
	TpSignature  [2]*big.Int
	TaSignature  [2]*big.Int
	AttestersIds []*big.Int
}

// IAttestationCenterEcdsaTaskSubmissionDetails is an auto generated low-level Go binding around an user-defined struct.
type IAttestationCenterEcdsaTaskSubmissionDetails struct {
	IsApproved   bool
	TpSignature  []byte
	TaSignature  [2]*big.Int
	AttestersIds []*big.Int
}

// IAttestationCenterInitializationParams is an auto generated low-level Go binding around an user-defined struct.
type IAttestationCenterInitializationParams struct {
	AvsGovernanceMultisigOwner common.Address
	OperationsMultisig         common.Address
	CommunityMultisig          common.Address
	MessageHandler             common.Address
	Obls                       common.Address
	AvsTreasury                common.Address
	IsRewardsOnL2              bool
	InternalTaskHandler        common.Address
}

// IAttestationCenterPaymentDetails is an auto generated low-level Go binding around an user-defined struct.
type IAttestationCenterPaymentDetails struct {
	Operator           common.Address
	LastPaidTaskNumber *big.Int
	FeeToClaim         *big.Int
	PaymentStatus      uint8
}

// IAttestationCenterPaymentRequestMessage is an auto generated low-level Go binding around an user-defined struct.
type IAttestationCenterPaymentRequestMessage struct {
	Operator   common.Address
	FeeToClaim *big.Int
}

// IAttestationCenterTaskInfo is an auto generated low-level Go binding around an user-defined struct.
type IAttestationCenterTaskInfo struct {
	ProofOfTask      string
	Data             []byte
	TaskPerformer    common.Address
	TaskDefinitionId uint16
}

// ContractAttestationCenterMetaData contains all meta data concerning the ContractAttestationCenter contract.
var ContractAttestationCenterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_extensionImplementation\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EigenRewardsDurationExceedsMaximum\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EigenRewardsDurationNotMultipleOfInterval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"totalRewards\",\"type\":\"uint256\"}],\"name\":\"EigenRewardsMaxRewardsAmountExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EigenRewardsMustBeRetroactive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EigenRewardsNotSupportedOnL2\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EigenRewardsStartTimestampNotMultipleOfInterval\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EigenRewardsStartTimestampTooFarInPast\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FlowIsCurrentlyPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FlowIsCurrentlyUnpaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InactiveAggregator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InactiveTaskPerformer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"taskDefinitionId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"minVotingPower\",\"type\":\"uint256\"}],\"name\":\"InsufficientVotingPowerForTaskDefinition\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAttesterSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaximumNumberOfAttesters\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOperatorsForPayment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPerformerSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeForBatchPaymentRequest\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"taskDefinitionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"operatorIndex\",\"type\":\"uint256\"}],\"name\":\"InvalidRestrictedAttester\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTaskDefinition\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageAlreadySigned\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAnEjector\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operatorAddress\",\"type\":\"address\"}],\"name\":\"OperatorNotRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PauseFlowIsAlreadyPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"taskDefinitionId\",\"type\":\"uint16\"}],\"name\":\"TaskDefinitionNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnpausingFlowIsAlreadyUnpaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestedTaskNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestedAmountClaimed\",\"type\":\"uint256\"}],\"name\":\"ClearPaymentRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"startTimestamp\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"duration\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeToClaim\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIAttestationCenter.PaymentRequestMessage[]\",\"name\":\"operators\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lastPaidTaskNumber\",\"type\":\"uint256\"}],\"name\":\"EigenPaymentsRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"_pausableFlow\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_pauser\",\"type\":\"address\"}],\"name\":\"FlowPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"_pausableFlowFlag\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_unpauser\",\"type\":\"address\"}],\"name\":\"FlowUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[4]\",\"name\":\"blsKey\",\"type\":\"uint256[4]\"}],\"name\":\"OperatorBlsKeyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorEjectionRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"votingPower\",\"type\":\"uint256\"}],\"name\":\"OperatorRegisteredToNetwork\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"operatorId\",\"type\":\"uint256\"}],\"name\":\"OperatorUnregisteredFromNetwork\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeToClaim\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIAttestationCenter.PaymentRequestMessage[]\",\"name\":\"operators\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lastPaidTaskNumber\",\"type\":\"uint256\"}],\"name\":\"PaymentsRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_baseRewardFeeForOperator\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"_taskNumber\",\"type\":\"uint32\"}],\"name\":\"RewardAccumulated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newMessageHandler\",\"type\":\"address\"}],\"name\":\"SetMessageHandler\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"taskNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"proofOfTask\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"taskDefinitionId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"attestersIds\",\"type\":\"uint256[]\"}],\"name\":\"TaskRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"taskNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"proofOfTask\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"taskDefinitionId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"attestersIds\",\"type\":\"uint256[]\"}],\"name\":\"TaskSubmitted\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXTENSION_IMPLEMENTATION\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"avsLogic\",\"outputs\":[{\"internalType\":\"contractIAvsLogic\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"avsTreasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beforePaymentsLogic\",\"outputs\":[{\"internalType\":\"contractIBeforePaymentsLogic\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeToClaim\",\"type\":\"uint256\"}],\"internalType\":\"structIAttestationCenter.PaymentRequestMessage[]\",\"name\":\"_operators\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"_paidTaskNumber\",\"type\":\"uint256\"}],\"name\":\"clearBatchPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"ejectOperatorFromNetwork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"getOperatorPaymentDetail\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lastPaidTaskNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeToClaim\",\"type\":\"uint256\"},{\"internalType\":\"enumIAttestationCenter.PaymentStatus\",\"name\":\"paymentStatus\",\"type\":\"uint8\"}],\"internalType\":\"structIAttestationCenter.PaymentDetails\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_taskDefinitionId\",\"type\":\"uint16\"}],\"name\":\"getTaskDefinitionMaximumNumberOfAttesters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_taskDefinitionId\",\"type\":\"uint16\"}],\"name\":\"getTaskDefinitionMinimumVotingPower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_taskDefinitionId\",\"type\":\"uint16\"}],\"name\":\"getTaskDefinitionRestrictedAttesters\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"avsGovernanceMultisigOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operationsMultisig\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"communityMultisig\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"messageHandler\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"obls\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"avsTreasury\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isRewardsOnL2\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"internalTaskHandler\",\"type\":\"address\"}],\"internalType\":\"structIAttestationCenter.InitializationParams\",\"name\":\"_initializationParams\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"internalTaskHandler\",\"outputs\":[{\"internalType\":\"contractIInternalTaskHandler\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_pausableFlow\",\"type\":\"bytes4\"}],\"name\":\"isFlowPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_isPaused\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextEigenRewardsBatchStartTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numOfActiveOperators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numOfTaskDefinitions\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numOfTotalOperators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"obls\",\"outputs\":[{\"internalType\":\"contractIOBLS\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"operatorsIdsByAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_pausableFlow\",\"type\":\"bytes4\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_votingPower\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"_blsKey\",\"type\":\"uint256[4]\"},{\"internalType\":\"address\",\"name\":\"_rewardsReceiver\",\"type\":\"address\"}],\"name\":\"registerToNetwork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_to\",\"type\":\"uint256\"}],\"name\":\"requestBatchPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestBatchPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_startTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_duration\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_to\",\"type\":\"uint256\"}],\"name\":\"requestEigenBatchPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isRewardsOnL2\",\"type\":\"bool\"}],\"name\":\"setIsRewardsOnL2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oblsSharesSyncer\",\"type\":\"address\"}],\"name\":\"setOblsSharesSyncer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"proofOfTask\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"taskPerformer\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"taskDefinitionId\",\"type\":\"uint16\"}],\"internalType\":\"structIAttestationCenter.TaskInfo\",\"name\":\"_taskInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isApproved\",\"type\":\"bool\"},{\"internalType\":\"uint256[2]\",\"name\":\"tpSignature\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"taSignature\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[]\",\"name\":\"attestersIds\",\"type\":\"uint256[]\"}],\"internalType\":\"structIAttestationCenter.BlsTaskSubmissionDetails\",\"name\":\"_blsTaskSubmissionDetails\",\"type\":\"tuple\"}],\"name\":\"submitTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"proofOfTask\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"taskPerformer\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"taskDefinitionId\",\"type\":\"uint16\"}],\"internalType\":\"structIAttestationCenter.TaskInfo\",\"name\":\"_taskInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isApproved\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"tpSignature\",\"type\":\"bytes\"},{\"internalType\":\"uint256[2]\",\"name\":\"taSignature\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[]\",\"name\":\"attestersIds\",\"type\":\"uint256[]\"}],\"internalType\":\"structIAttestationCenter.EcdsaTaskSubmissionDetails\",\"name\":\"_ecdsaTaskSubmissionDetails\",\"type\":\"tuple\"}],\"name\":\"submitTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taskNumber\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newMessageHandler\",\"type\":\"address\"}],\"name\":\"transferMessageHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"unRegisterOperatorFromNetwork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_pausableFlow\",\"type\":\"bytes4\"}],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"_blsKey\",\"type\":\"uint256[4]\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"signature\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBLSAuthLibrary.Signature\",\"name\":\"_authSignature\",\"type\":\"tuple\"}],\"name\":\"updateBlsKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_taskDefinitionId\",\"type\":\"uint16\"}],\"name\":\"verifyOperatorValidForTaskDefinition\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"votingPower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506004361061025d5760003560e01c80638355e33d11610146578063b7aa2fdf116100c3578063dba8323b11610087578063dba8323b1461054c578063efd969781461055f578063f28eff3914610599578063f6e258cc146105a1578063fcd4e66a146105b4578063fff768e3146105c75761025d565b8063b7aa2fdf14610503578063bac1e94b1461050b578063c07473f61461051e578063c2f429f114610531578063d547741f146105395761025d565b806397b5f3701161010a57806397b5f370146104a05780639eb72d4c146104c0578063a217fddf146104e0578063ac8c16cf146104e8578063b0817c44146104fb5761025d565b80638355e33d1461044c5780638a7ce8a414610454578063915359fc1461046757806391d148541461047a57806395983fee1461048d5761025d565b806336568abe116101df578063659fa976116101a3578063659fa976146103e65780636ba5aa46146103ee5780636f3826191461040157806372d18e8d1461041457806375d9aedf146104315780637897dec3146104445761025d565b806336568abe146103875780633aa83ec71461039a5780634d07f651146103ad5780635b15c568146103c05780635b386be0146103d35761025d565b8063226def0411610226578063226def041461030c578063248a9ca31461033357806327bbb287146103465780632f2ff15d1461035957806334a7c3911461036c5761025d565b8062028b071461028857806301ffc9a7146102a35780630a1e7093146102c65780631164224e146102d95780631246193e146102ec575b6102867f0000000000000000000000003f4d414c5378e6d123322819752506484e97c9ee6105da565b005b610290610603565b6040519081526020015b60405180910390f35b6102b66102b1366004614001565b610616565b604051901515815260200161029a565b6102866102d4366004614044565b61064d565b6102866102e736600461409d565b61091c565b6102f4610999565b6040516001600160a01b03909116815260200161029a565b6102f47f0000000000000000000000003f4d414c5378e6d123322819752506484e97c9ee81565b6102906103413660046140b8565b6109ba565b61028661035436600461409d565b6109dc565b6102866103673660046140d1565b610ad4565b610374610af6565b60405161ffff909116815260200161029a565b6102866103953660046140d1565b610b0d565b6102866103a8366004614001565b610b40565b6102866103bb36600461409d565b610b67565b6102906103ce36600461409d565b610c10565b6102906103e136600461410f565b610c23565b6102f4610c4f565b6102866103fc36600461412a565b610c6b565b61028661040f366004614170565b610d98565b61041c610de5565b60405163ffffffff909116815260200161029a565b61029061043f36600461410f565b610dfb565b610290610e25565b610290610e38565b6102866104623660046141aa565b610e4b565b6102866104753660046142c7565b610f89565b6102b66104883660046140d1565b611064565b61028661049b36600461438d565b61109c565b6104b36104ae36600461410f565b6111a9565b60405161029a91906143dc565b6104d36104ce3660046140b8565b61122a565b60405161029a9190614405565b610290600081565b6102866104f636600461445b565b6112db565b6102f4611543565b61028661155f565b610286610519366004614001565b6115b6565b61029061052c36600461409d565b6115dd565b6102f461166e565b6102866105473660046140d1565b61168a565b61028661055a366004614493565b6116a6565b6102b661056d366004614001565b6001600160e01b031916600090815260008051602061517b833981519152602052604090205460ff1690565b6102f46116d6565b6102866105af36600461409d565b6116f2565b6102866105c23660046144b0565b611803565b6102866105d536600461453a565b6119cf565b3660008037600080366000845af43d6000803e8080156105f9573d6000f35b3d6000fd5b505050565b600061060d611abe565b60020154905090565b60006001600160e01b03198216637965db0b60e01b148061064757506301ffc9a760e01b6001600160e01b03198316145b92915050565b7f4b5254ecaf62d67135de29b9998979c2380f8b9012725c83915824ad95f4e88c61067781611acd565b6378b4401360e11b61068881611b18565b6000610692611abe565b600e81015490915060ff16156106bb5760405163ac187e3160e01b815260040160405180910390fd5b6106c58787611b22565b60008060006106d5848989611c2f565b9250925092506f4b3b4ca85a86c47a098a223fffffffff8111156107145760405163c179495760e01b8152600481018290526024015b60405180910390fd5b81156108df57600182111561084e5760005b6107316001846145af565b81101561084c5760005b600161074783866145af565b61075191906145af565b81101561084357846107648260016145c2565b81518110610774576107746145d5565b6020026020010151600001516001600160a01b031685828151811061079b5761079b6145d5565b6020026020010151600001516001600160a01b0316111561083b57846107c28260016145c2565b815181106107d2576107d26145d5565b60200260200101518582815181106107ec576107ec6145d5565b6020026020010151868381518110610806576108066145d5565b602002602001018784600161081b91906145c2565b8151811061082b5761082b6145d5565b6020026020010182905282905250505b60010161073b565b50600101610726565b505b6040805163ffffffff8c811660208301528b16818301526060810183905260808082018590528251808303909101815260a0909101909152828452610894858583611e1a565b84546040517f928efea964c83ee05474996615e6ed34feaad6ce29a94eeec47d3f69388f5513916108d1918e918e91899163ffffffff1690614630565b60405180910390a1506108f8565b6040516302e66f0160e31b815260040160405180910390fd5b610902898b61466c565b63ffffffff16846011018190555050505050505050505050565b6378b4401360e11b61092d81611b18565b610935611abe565b600301546040516308b2112760e11b81526001600160a01b03848116600483015290911690631164224e90602401600060405180830381600087803b15801561097d57600080fd5b505af1158015610991573d6000803e3d6000fd5b505050505050565b60006109a3611abe565b600e015461010090046001600160a01b0316919050565b600090815260008051602061519b833981519152602052604090206001015490565b638a70a0eb60e01b6109ed81611b18565b60006109f7611abe565b60038101549091506001600160a01b03166000610a148386611eb8565b6001600160a01b03861660009081526009850160205260408120819055600b8501805492935091610a4490614689565b909155506040516333ccef3360e21b8152600481018290526001600160a01b0383169063cf33bccc90602401600060405180830381600087803b158015610a8a57600080fd5b505af1158015610a9e573d6000803e3d6000fd5b50506040518392507fda04f7db725bc5a9ad418baf26d08e9f24561a7cc119bfe1dd26bfebfc175db39150600090a25050505050565b610add826109ba565b610ae681611b18565b610af08383611f05565b50505050565b6000610b00611abe565b6006015461ffff16919050565b6001600160a01b0381163314610b365760405163334bd91960e11b815260040160405180910390fd5b6105fe8282611fb1565b80610b4a81611acd565b6001600160e01b03198216610b5e81611b18565b6105fe8361202d565b63d8a8b5c760e01b610b7881611b18565b6000610b82611abe565b6004810154909150610ba590638a70a0eb60e01b906001600160a01b0316611fb1565b50610bb7638a70a0eb60e01b84611f05565b506004810180546001600160a01b0319166001600160a01b0385169081179091556040519081527f997f84b541d7b68e210e6f50e3402b51d8411dbbc4d44ed81e508383126e4e949060200160405180910390a1505050565b6000610647610c1d611abe565b83611eb8565b600061064782610c31611abe565b61ffff90911660009081526007909101602052604090206009015490565b6000610c59611abe565b600301546001600160a01b0316919050565b6000610c75611abe565b9050336000610c848383611eb8565b600384015460405163475d551f60e11b81529192506001600160a01b0316908190638ebaaa3e90610cbf908890879030908c906004016146a0565b60006040518083038186803b158015610cd757600080fd5b505afa158015610ceb573d6000803e3d6000fd5b50506040516386d897a760e01b81526001600160a01b03841692506386d897a79150610d1d9085908a906004016146d1565b600060405180830381600087803b158015610d3757600080fd5b505af1158015610d4b573d6000803e3d6000fd5b50505050826001600160a01b03167f764bc14e663abcee4585e1a92e552918c69d453c673e7161504ff52fc3d428c987604051610d8891906146e7565b60405180910390a2505050505050565b7fc6d727150e371ebc5958e5c678aae159bc7f1ac1d986edc36fecf6e51285a057610dc281611acd565b6378b4401360e11b610dd381611b18565b610af0610dde611abe565b85856120da565b6000610def611abe565b5463ffffffff16919050565b600061064782610e09611abe565b61ffff9091166000908152600791820160205260409020015490565b6000610e2f611abe565b600b0154905090565b6000610e42611abe565b60110154905090565b7f2aedfbf340ba26a0aa3515c5f66276ef1d8ff6c89d5dbc63de01abc44b017be6610e7581611acd565b610e7d61215d565b6040805160a0810190915260009080610e996020860186614493565b1515815260200160405180602001604052806000815250815260200184602001600280602002604051908101604052809291908260026020028082843760009201919091525050508152604080518082018252602090920191906060870190600290839083908082843760009201919091525050508152602001610f2060a08601866146f6565b8080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050509152509050610f5f84826121a7565b506105fe60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b638a70a0eb60e01b610f9a81611b18565b6000610fa4611abe565b905060005b845181108015610fe9575060006001600160a01b0316858281518110610fd157610fd16145d5565b6020026020010151600001516001600160a01b031614155b1561105d576000858281518110611002576110026145d5565b60209081029190910181015180516001600160a01b03166000908152600986018352604080822054808352600a8801855291209282015191935091611048918890612970565b5050808061105590614746565b915050610fa9565b5050505050565b600091825260008051602061519b833981519152602090815260408084206001600160a01b0393909316845291905290205460ff1690565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03166000811580156110e15750825b90506000826001600160401b031660011480156110fd5750303b155b90508115801561110b575080155b156111295760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561115357845460ff60401b1916600160401b1785555b61115c86612a04565b831561099157845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a1505050505050565b60606111d4826111b7611abe565b61ffff909116600090815260079091016020526040902060080190565b80548060200260200160405190810160405280929190818152602001828054801561121e57602002820191906000526020600020905b81548152602001906001019080831161120a575b50505050509050919050565b6112526040805160808101825260008082526020820181905291810182905290606082015290565b61125a611abe565b6000838152600a919091016020908152604091829020825160808101845281546001600160a01b0316815260018201549281019290925260028082015493830193909352600381015491929091606084019160ff909116908111156112c1576112c16143ef565b60028111156112d2576112d26143ef565b90525092915050565b60006112e5611abe565b905060006112f38285611eb8565b905060006113018385612c1d565b604080516101408101909152815461ffff16815260018201805491929160208401919061132d9061475f565b80601f01602080910402602001604051908101604052809291908181526020018280546113599061475f565b80156113a65780601f1061137b576101008083540402835291602001916113a6565b820191906000526020600020905b81548152906001019060200180831161138957829003601f168201915b505050505081526020016002820154815260200160038201548152602001600482015481526020016005820154815260200160068201548152602001600782015481526020016008820180548060200260200160405190810160405280929190818152602001828054801561143a57602002820191906000526020600020905b815481526020019060010190808311611426575b5050509183525050600991909101546020918201526040805160018082528183019092529293506000929182810190803683370190505090508281600081518110611487576114876145d5565b60200260200101818152505061149d8282612c5f565b60e082015160038501546040516372c4a92760e01b8152600481018690526001600160a01b03909116906372c4a92790602401602060405180830381865afa1580156114ed573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115119190614793565b10156109915760e08201516040516349ba8abd60e11b815261ffff87166004820152602481019190915260440161070b565b600061154d611abe565b600501546001600160a01b0316919050565b7fc6d727150e371ebc5958e5c678aae159bc7f1ac1d986edc36fecf6e51285a05761158981611acd565b6378b4401360e11b61159a81611b18565b60006115a4611abe565b90506105fe81600183600201546120da565b806115c081612d2a565b6001600160e01b031982166115d481611b18565b6105fe83612d71565b6000806115e8611abe565b905060006115f68285611eb8565b60038301546040516372c4a92760e01b8152600481018390529192506001600160a01b0316906372c4a92790602401602060405180830381865afa158015611642573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116669190614793565b949350505050565b6000611678611abe565b600f01546001600160a01b0316919050565b611693826109ba565b61169c81611b18565b610af08383611fb1565b6378b4401360e11b6116b781611b18565b816116c0611abe565b600e01805460ff19169115159190911790555050565b60006116e0611abe565b601001546001600160a01b0316919050565b63297da43760e11b61170381611b18565b600061170d611abe565b9050600061171b8285611eb8565b60038301546040516382afd23b60e01b8152600481018390529192506001600160a01b03169081906382afd23b90602401602060405180830381865afa158015611769573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061178d91906147ac565b6117b55760405163bd62013360e01b81526001600160a01b038616600482015260240161070b565b6117bf8386612e0e565b6040516001600160a01b03861681527f4bbb3ae76ccd0feb230e3802a79c234a035ea750a918f60db5ae976d67c096d0906020015b60405180910390a15050505050565b638a70a0eb60e01b61181481611b18565b600061181e611abe565b90506000816002016000815461183390614746565b919050819055905081600b016000815461184c90614746565b90915550604080516080810182526001600160a01b0389811682526000602080840182815284860183815260608601848152888552600a8a0190935295909220845181546001600160a01b0319169416939093178355905160018084019190915593516002808401919091559051600383018054949593949193909260ff199092169184908111156118e0576118e06143ef565b021790555050506001600160a01b0387811660009081526009840160209081526040808320859055600d86019091529081902080546001600160a01b0319168784161790556003840154905163891a80bb60e01b815291169063891a80bb906119519084908a908a906004016147c9565b600060405180830381600087803b15801561196b57600080fd5b505af115801561197f573d6000803e3d6000fd5b50505050866001600160a01b03167f16c1a2a8195923d655fe84191b37c746f4385a5c32c038578958b29f52daa1a8876040516119be91815260200190565b60405180910390a250505050505050565b7f2aedfbf340ba26a0aa3515c5f66276ef1d8ff6c89d5dbc63de01abc44b017be66119f981611acd565b611a0161215d565b6040805160a0810190915260009080611a1d6020860186614493565b15158152602001848060200190611a34919061480a565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052509385525050604080518082018252838152602081810194909452928401929092525080518082018252918101919086810190600290839083908082843760009201919091525050508152602001610f2060808601866146f6565b6000611ac8612eb5565b905090565b6001600160e01b03198116600090815260008051602061517b833981519152602052604090205460ff1615611b155760405163722fdba960e01b815260040160405180910390fd5b50565b611b158133612ee3565b625c490063ffffffff82161115611b4c57604051639464cf7960e01b815260040160405180910390fd5b611b5962093a8082614866565b63ffffffff1615611b7d5760405163e9742bc360e01b815260040160405180910390fd5b611b8a62093a8083614866565b63ffffffff1615611bae5760405163ba36d17960e01b815260040160405180910390fd5b611bbb62dd7c0042614889565b63ffffffff168263ffffffff161080611bdd57506365fb788063ffffffff8316105b15611bfb57604051634bf010dd60e01b815260040160405180910390fd5b42611c06828461466c565b63ffffffff161115611c2b5760405163031b909b60e51b815260040160405180910390fd5b5050565b6060600080841580611c445750856002015484115b80611c4e57508385115b15611c6c57604051639ac893ad60e01b815260040160405180910390fd5b855463ffffffff166000611c8087876145af565b611c8b9060016145c2565b6001600160401b03811115611ca257611ca2614214565b604051908082528060200260200182016040528015611ce757816020015b6040805180820190915260008082526020820152815260200190600190039081611cc05790505b509050600080885b888111611e09576000818152600a8c016020526040902080546001600160a01b031615801590611d3757506000600382015460ff166002811115611d3557611d356143ef565b145b8015611d4c57508563ffffffff168160010154105b8015611d5c575060008160020154115b15611e00576040805180820190915281546001600160a01b03168152600282015460208201528551869086908110611d9657611d966145d5565b602090810291909101015260038101805460ff1916600117905583611dba81614746565b945050806002015483611dcd91906145c2565b600e8d015490935060ff1615611e0057611e008c82888f600e0160019054906101000a90046001600160a01b0316612f1c565b50600101611cef565b509199909850909650945050505050565b6000611e4e83604051602001611e3091906148a6565b60408051601f19818403018152919052855463ffffffff1684613004565b60048086015460405163104c8d4b60e31b81529293506001600160a01b0316916382646a5891611e80918591016148ff565b600060405180830381600087803b158015611e9a57600080fd5b505af1158015611eae573d6000803e3d6000fd5b5050505050505050565b6001600160a01b0381166000908152600983016020526040812054808203611efe5760405163bd62013360e01b81526001600160a01b038416600482015260240161070b565b9392505050565b600060008051602061519b833981519152611f208484611064565b611fa0576000848152602082815260408083206001600160a01b03871684529091529020805460ff19166001179055611f563390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610647565b6000915050610647565b5092915050565b600060008051602061519b833981519152611fcc8484611064565b15611fa0576000848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610647565b6001600160e01b03198116600090815260008051602061517b833981519152602081905260409091205460ff16156120785760405163dfe10d7d60e01b815260040160405180910390fd5b6001600160e01b0319821660009081526020829052604090819020805460ff19166001179055517f95c3658c5e0c74e20cf12db371b9b67d26e97a1937f6d2284f88cc44d036b4f6906120ce9084903390614912565b60405180910390a15050565b6000806120e8858585611c2f565b50915091508060000361210e576040516302e66f0160e31b815260040160405180910390fd5b600e85015460ff16612124576121248583613077565b84546040517f3e17ccbb628e667c33839a666b60f15eaefb9db2cbae6f7cc9f3f223cd77fece916117f491859163ffffffff1690614935565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f008054600119016121a157604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b60006127116121bc608085016060860161410f565b61ffff1610159050600082604001516000600281106121dd576121dd6145d5565b60200201511515905060006121f0611abe565b60058101549091506001600160a01b031660008115801590612210575084155b9050801561230e578315612298578551604080880151606089015160808a015192516379bc050160e01b81526001600160a01b038716946379bc050194612261948e94929390929190600401614a6e565b600060405180830381600087803b15801561227b57600080fd5b505af115801561228f573d6000803e3d6000fd5b5050505061230e565b8551602087015160608801516080890151604051630502f5bd60e41b81526001600160a01b0387169463502f5bd0946122db948e94929391929091600401614ac1565b600060405180830381600087803b1580156122f557600080fd5b505af1158015612309573d6000803e3d6000fd5b505050505b600061231a888061480a565b61232760208b018b61480a565b61233760608d0160408e0161409d565b61234760808e0160608f0161410f565b60405160200161235c96959493929190614b0f565b60408051601f19818403018152918152815160209283012060008181526008880190935291205490915060ff16156123a757604051634510302360e11b815260040160405180910390fd5b60038401546001600160a01b031685156124f3576000816001600160a01b031663a850a9097fb110049506439a07d78731efed3c809a6b13e0bab7cc9805f3bb45fb1e9a67e88560405160200161240091815260200190565b6040516020818303038152906040526040518363ffffffff1660e01b815260040161242c929190614b5e565b6040805180830381865afa158015612448573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061246c9190614b77565b9050816001600160a01b03166353c051bb612499888d6040016020810190612494919061409d565b611eb8565b838c604001516040518463ffffffff1660e01b81526004016124bd93929190614bd7565b60006040518083038186803b1580156124d557600080fd5b505afa1580156124e9573d6000803e3d6000fd5b5050505050612511565b61251161250660608b0160408c0161409d565b838a602001516130aa565b6000816001600160a01b031663a850a9097fb110049506439a07d78731efed3c809a6b13e0bab7cc9805f3bb45fb1e9a67e86125518d8d600001516130ea565b60405160200161256391815260200190565b6040516020818303038152906040526040518363ffffffff1660e01b815260040161258f929190614b5e565b6040805180830381865afa1580156125ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125cf9190614b77565b905061262b604051806101400160405280600061ffff1681526020016060815260200160008152602001600081526020016000815260200160008152602001600081526020016000815260200160608152602001600081525090565b6126448761263f60808e0160608f0161410f565b612c1d565b604080516101408101909152815461ffff1681526001820180549192916020840191906126709061475f565b80601f016020809104026020016040519081016040528092919081815260200182805461269c9061475f565b80156126e95780601f106126be576101008083540402835291602001916126e9565b820191906000526020600020905b8154815290600101906020018083116126cc57829003601f168201915b505050505081526020016002820154815260200160038201548152602001600482015481526020016005820154815260200160068201548152602001600782015481526020016008820180548060200260200160405190810160405280929190818152602001828054801561277d57602002820191906000526020600020905b815481526020019060010190808311612769575b5050505050815260200160098201548152505090506127a0818b60800151612c5f565b6127cd838b600001518d60600160208101906127bc919061410f565b848e606001518f608001518861315c565b6127e2878c838d600001518e608001516131de565b60008481526008880160205260409020805460ff1916600117905584156128c857871561288457856001600160a01b031663a971515d8c8c600001518d604001518e606001518f608001516040518663ffffffff1660e01b815260040161284d959493929190614a6e565b600060405180830381600087803b15801561286757600080fd5b505af115801561287b573d6000803e3d6000fd5b5050505061293d565b856001600160a01b031663dd1a53878c8c600001518d602001518e606001518f608001516040518663ffffffff1660e01b815260040161284d959493929190614ac1565b8880156128d3575089515b1561293d576010870154604051639d00234560e01b81526001600160a01b0390911690639d0023459061290a908e90600401614bf8565b600060405180830381600087803b15801561292457600080fd5b505af1158015612938573d6000803e3d6000fd5b505050505b5050505050505050505050565b60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b61297b838383613621565b156129b85780156129a957818360010181905550808360020160008282546129a391906145af565b90915550505b5050600301805460ff19169055565b825460408051848152602081018490526001600160a01b03909216917f1e643658b8248efd3563f24d116430bf571d036bea3721d94e848890a00a1023910160405180910390a2505050565b612a0c61367c565b6000612a1b602083018361409d565b90506000612a2f604084016020850161409d565b90506000612a43606085016040860161409d565b90506000612a57608086016060870161409d565b90506000612a6b60c0870160a0880161409d565b90506001600160a01b038116612a945760405163d92e233d60e01b815260040160405180910390fd5b612a9f8585856136c7565b612aaa858585613704565b612ab261372a565b6000612abc611abe565b805463ffffffff191660011781559050612adc60a088016080890161409d565b6003820180546001600160a01b0319166001600160a01b0392909216919091179055612b0e608088016060890161409d565b6004820180546001600160a01b0319166001600160a01b0392909216919091179055612b4060e0880160c08901614493565b600e820180546001600160a81b031916911515610100600160a81b031916919091176101006001600160a01b038516810291909117909155612b8790880160e0890161409d565b6010820180546001600160a01b0319166001600160a01b0392909216919091179055612bb662093a8042614c0b565b612bc090426145af565b6011820155612bd6638a70a0eb60e01b84611f05565b5060008080526007820160205260409020612bf09061373a565b61271160009081526007820160205260409020612c0c9061378f565b612c146137e4565b50505050505050565b61ffff81811660008181526007850160205260408120805491939092911614611efe576040516321321e1960e11b815261ffff8416600482015260240161070b565b43826040015111612c8357604051635a31d91b60e01b815260040160405180910390fd5b610100820151805115612cf257815181511015612cb357604051630a8d477960e01b815260040160405180910390fd5b6000612cbf8383613847565b90508015612cf0578351604051635966606760e11b815261ffff90911660048201526024810182905260440161070b565b505b6000836101200151118015612d0c57508151836101200151105b156105fe57604051630236998b60e01b815260040160405180910390fd5b6001600160e01b03198116600090815260008051602061517b833981519152602052604090205460ff16611b15576040516368c87f3360e11b815260040160405180910390fd5b6001600160e01b03198116600090815260008051602061517b833981519152602081905260409091205460ff16612dbb57604051635bfd2da760e11b815260040160405180910390fd5b6001600160e01b0319821660009081526020829052604090819020805460ff19169055517fc7e56e17b0a6c4b467df6495e1eda1baecd7ba20604e80c1058ac06f4578d85e906120ce9084903390614912565b604080516001600160a01b0383166024808301919091528251808303909101815260449091019091526020810180516001600160e01b0316636e05955960e01b1790526000905b60048085015460405163104c8d4b60e31b81529293506001600160a01b0316916382646a5891612e87918591016148ff565b600060405180830381600087803b158015612ea157600080fd5b505af1158015612c14573d6000803e3d6000fd5b60008061064760017f47c1bdca9c75057bf1ca178a62fb0ef8908b21ee1ea9b9658ba5135fd3a354696145af565b612eed8282611064565b611c2b5760405163e2517d3f60e01b81526001600160a01b03821660048201526024810183905260440161070b565b82546001600160a01b039081166000908152600d86016020526040812054909116908115612f4a5781612f56565b84546001600160a01b03165b6002860154604051633256b4d160e01b81526001600160a01b03808416600483015263ffffffff88166024830152604482019290925291925060009190851690633256b4d1906064016020604051808303816000875af1158015612fbe573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612fe291906147ac565b905060008115612ff3575060028601545b611eae878763ffffffff1683612970565b60607f9d0737d982403370785b2d562a1c75961d85a320fe394f421fc3de24bbbce7b884848460405160240161303c93929190614c1f565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b03199093169290921790915290509392505050565b6000612e558260405160200161308d91906148a6565b60408051601f19818403018152919052845463ffffffff16613934565b60006130b683836139a4565b9050836001600160a01b0316816001600160a01b031614610af057604051632be1e1cb60e11b815260040160405180910390fd5b60006130f6838061480a565b613103602086018661480a565b613113606088016040890161409d565b6131236080890160608a0161410f565b30468960405160200161313e99989796959493929190614c54565b60405160208183030381529060405280519060200120905092915050565b600061316a888888886139ce565b60e08601516040516365c4647560e01b81529192506001600160a01b038a16916365c46475916131a4918691899189918891600401614cbb565b60006040518083038186803b1580156131bc57600080fd5b505afa1580156131d0573d6000803e3d6000fd5b505050505050505050505050565b845463ffffffff168215613272576131fc608086016060870161410f565b61ffff16613210606087016040880161409d565b6001600160a01b03167faed5b9322f1cce79090471045d025d7a0a841daf787ec955d17cdd32a3ebfdb183613245898061480a565b61325260208c018c61480a565b8960405161326596959493929190614cfc565b60405180910390a36132f4565b613282608086016060870161410f565b61ffff16613296606087016040880161409d565b6001600160a01b03167f834bcdc44f628888090e897856fa3d661504c545654ba804c61b293c36b6595b836132cb898061480a565b6132d860208c018c61480a565b896040516132eb96959493929190614cfc565b60405180910390a35b600061330a876124946060890160408a0161409d565b9050600061331788613af9565b601089015460038a01546040516382afd23b60e01b815260048101869052929350600160a01b90910460ff16916001600160a01b039091169081906382afd23b90602401602060405180830381865afa158015613378573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061339c91906147ac565b6133b9576040516346d3a4a960e01b815260040160405180910390fd5b8115801561342d57506040516382afd23b60e01b8152600481018490526001600160a01b038216906382afd23b90602401602060405180830381865afa158015613407573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061342b91906147ac565b155b1561344b57604051631385e31560e21b815260040160405180910390fd5b5060006040518060a001604052808a61346390614dc7565b8152602081018590526040810186905260608101889052881515608090910152600c8b01549091506001600160a01b031680158015908061350957508080156135095750816001600160a01b031663d005a4f66040518163ffffffff1660e01b8152600401602060405180830381865afa1580156134e5573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061350991906147ac565b1561355b57600080600061351e8f8e88613b1b565b9250925092506135308f848d8d613bd5565b8b15613542576135428f838b8d613c0e565b86613553576135538f828a8d613c0e565b5050506135e0565b6040516336684b6760e01b81526000906001600160a01b038416906336684b679061358a908790600401614e6b565b6000604051808303816000875af11580156135a9573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526135d19190810190614f20565b90506135de8d828a613c79565b505b8b5463ffffffff168c60006135f483614fd5565b91906101000a81548163ffffffff021916908363ffffffff16021790555050505050505050505050505050565b60006001600385015460ff16600281111561363e5761363e6143ef565b1461364b57506000611efe565b8284600101541061365e57506000611efe565b818460020154101561367257506000611efe565b5060019392505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff166136c557604051631afcd79f60e31b815260040160405180910390fd5b565b6136cf61367c565b6136e063d8a8b5c760e01b83611f05565b506136f26378b4401360e11b84611f05565b50610af063d3e319af60e01b82611f05565b61370c61367c565b613714613cdf565b61371f838383613ce7565b6105fe600033610b0d565b61373261367c565b6136c5613e24565b604080518082019091526007815266191959985d5b1d60ca1b602082015260018201906137679082615040565b506000196002820155678ac7230489e800006003820181905560048201819055600590910155565b604080518082019091526016815275566f74696e6720506f7765722053796e63205461736b60501b602082015260018201906137cb9082615040565b50805461ffff1916612711178155600019600290910155565b60006137ee611abe565b61271260009081526007820160205260408120919250506040518060600160405280602781526020016151bb60279139600182019061382d9082615040565b50805461ffff191661271217815560001960029091015550565b600080805b84518210801561385c5750835181105b1561392157838181518110613873576138736145d5565b602002602001015185838151811061388d5761388d6145d5565b6020026020010151036138ba57816138a481614746565b92505080806138b290614746565b91505061384c565b8381815181106138cc576138cc6145d5565b60200260200101518583815181106138e6576138e66145d5565b602002602001015111156138fe57806138b281614746565b848281518110613910576139106145d5565b602002602001015192505050610647565b845182036138fe57600092505050610647565b60607fd136f2c682d3b91d0bd1dc36dc70e1a7ffbd859fd1ff2240257fdce81d74fd83838360405160240161396a9291906150ff565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152905092915050565b6000806000806139b48686613e2c565b9250925092506139c48282613e79565b5090949350505050565b60008060008361010001515111806139ea575060008360e00151115b15613a625760405163255da33160e01b815261ffff851660048201526001600160a01b0387169063255da33190602401602060405180830381865afa158015613a37573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613a5b9190614793565b9050613ac7565b856001600160a01b031663671b37936040518163ffffffff1660e01b8152600401602060405180830381865afa158015613aa0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613ac49190614793565b90505b8415613aec576003613ada826002615121565b613ae49190615138565b915050611666565b6003613ada826001615121565b60108101546000908190600160a01b900460ff1661064757611efe8333611eb8565b600c830154600090819081906001600160a01b03168015801590613bb45760405163156a159960e21b81526001600160a01b038316906355a8566490613b65908990600401614e6b565b6060604051808303816000875af1158015613b84573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613ba8919061514c565b91965094509250613bca565b86606001519450866080015193508660a0015192505b505093509350939050565b60005b825181101561105d57613c068585858481518110613bf857613bf86145d5565b602002602001015185613c0e565b600101613bd8565b6000828152600a8501602052604081206002018054859290613c319084906145c2565b909155505060405183815263ffffffff82169083907fd3f16e9d8d3fe0ea8a6e5f923fe57e1ae1af6d890ac6c371e8af6cc177a49b659060200160405180910390a350505050565b60005b8251811015610af0576000838281518110613c9957613c996145d5565b60200260200101516000015190506000848381518110613cbb57613cbb6145d5565b6020026020010151602001519050613cd586828487613c0e565b5050600101613c7c565b6136c561367c565b60008051602061517b8339815191526020527f90d0e609be45121efcbf7f1d0ae974a20fcd627e07aa044a1f490249e1d327988054600160ff1991821681179092556312d4953b60e21b6000527fc1d73a3cdf3346e56fde516857b4d0894fb686182db72c29641f07b7f38122b880549091169091179055613d7063c6d7271560e01b84611f05565b50613d826312d4953b60e21b84611f05565b50613d946362250a9560e11b84611f05565b50613da66362250a9560e11b82611f05565b50613db86362250a9560e11b83611f05565b50613dca634bb5f31f60e11b84611f05565b50613ddc634bb5f31f60e11b83611f05565b50613dee634bb5f31f60e11b82611f05565b50613e00632aedfbf360e01b84611f05565b50613e12632aedfbf360e01b83611f05565b50610af0632aedfbf360e01b82611f05565b61294a61367c565b60008060008351604103613e665760208401516040850151606086015160001a613e5888828585613f32565b955095509550505050613e72565b50508151600091506002905b9250925092565b6000826003811115613e8d57613e8d6143ef565b03613e96575050565b6001826003811115613eaa57613eaa6143ef565b03613ec85760405163f645eedf60e01b815260040160405180910390fd5b6002826003811115613edc57613edc6143ef565b03613efd5760405163fce698f760e01b81526004810182905260240161070b565b6003826003811115613f1157613f116143ef565b03611c2b576040516335e2f38360e21b81526004810182905260240161070b565b600080807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841115613f6d5750600091506003905082613ff7565b604080516000808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015613fc1573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116613fed57506000925060019150829050613ff7565b9250600091508190505b9450945094915050565b60006020828403121561401357600080fd5b81356001600160e01b031981168114611efe57600080fd5b803563ffffffff8116811461403f57600080fd5b919050565b6000806000806080858703121561405a57600080fd5b6140638561402b565b93506140716020860161402b565b93969395505050506040820135916060013590565b80356001600160a01b038116811461403f57600080fd5b6000602082840312156140af57600080fd5b611efe82614086565b6000602082840312156140ca57600080fd5b5035919050565b600080604083850312156140e457600080fd5b823591506140f460208401614086565b90509250929050565b803561ffff8116811461403f57600080fd5b60006020828403121561412157600080fd5b611efe826140fd565b60008082840360c081121561413e57600080fd5b608084018581111561414f57600080fd5b8493506040607f198301121561416457600080fd5b80925050509250929050565b6000806040838503121561418357600080fd5b50508035926020909101359150565b6000608082840312156141a457600080fd5b50919050565b600080604083850312156141bd57600080fd5b82356001600160401b03808211156141d457600080fd5b6141e086838701614192565b935060208501359150808211156141f657600080fd5b50830160c0818603121561420957600080fd5b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b038111828210171561424c5761424c614214565b60405290565b604051608081016001600160401b038111828210171561424c5761424c614214565b604051601f8201601f191681016001600160401b038111828210171561429c5761429c614214565b604052919050565b60006001600160401b038211156142bd576142bd614214565b5060051b60200190565b60008060408084860312156142db57600080fd5b83356001600160401b038111156142f157600080fd5b8401601f8101861361430257600080fd5b80356020614317614312836142a4565b614274565b82815260069290921b8301810191818101908984111561433657600080fd5b938201935b8385101561437d5785858b0312156143535760008081fd5b61435b61422a565b61436486614086565b815285840135848201528252938501939082019061433b565b9997909101359750505050505050565b600061010082840312156141a457600080fd5b60008151808452602080850194506020840160005b838110156143d1578151875295820195908201906001016143b5565b509495945050505050565b602081526000611efe60208301846143a0565b634e487b7160e01b600052602160045260246000fd5b81516001600160a01b031681526020808301519082015260408083015190820152606082015160808201906003811061444e57634e487b7160e01b600052602160045260246000fd5b8060608401525092915050565b6000806040838503121561446e57600080fd5b61447783614086565b91506140f4602084016140fd565b8015158114611b1557600080fd5b6000602082840312156144a557600080fd5b8135611efe81614485565b60008060008060e085870312156144c657600080fd5b6144cf85614086565b9350602080860135935086605f8701126144e857600080fd5b6144f0614252565b8060c088018981111561450257600080fd5b604089015b8181101561451e5780358452928401928401614507565b5081955061452b81614086565b94505050505092959194509250565b6000806040838503121561454d57600080fd5b82356001600160401b038082111561456457600080fd5b61457086838701614192565b9350602085013591508082111561458657600080fd5b50830160a0818603121561420957600080fd5b634e487b7160e01b600052601160045260246000fd5b8181038181111561064757610647614599565b8082018082111561064757610647614599565b634e487b7160e01b600052603260045260246000fd5b60008151808452602080850194506020840160005b838110156143d157815180516001600160a01b031688528301518388015260409096019590820190600101614600565b600063ffffffff808716835280861660208401526080604084015261465860808401866145eb565b915080841660608401525095945050505050565b63ffffffff818116838216019080821115611faa57611faa614599565b60008161469857614698614599565b506000190190565b610100810160408683376001600160a01b038581166040840152841660608301526080838184013795945050505050565b82815260a0810160808360208401379392505050565b60808181019083833792915050565b6000808335601e1984360301811261470d57600080fd5b8301803591506001600160401b0382111561472757600080fd5b6020019150600581901b360382131561473f57600080fd5b9250929050565b60006001820161475857614758614599565b5060010190565b600181811c9082168061477357607f821691505b6020821081036141a457634e487b7160e01b600052602260045260246000fd5b6000602082840312156147a557600080fd5b5051919050565b6000602082840312156147be57600080fd5b8151611efe81614485565b838152602080820184905260c0820190604083018460005b60048110156147fe578151835291830191908301906001016147e1565b50505050949350505050565b6000808335601e1984360301811261482157600080fd5b8301803591506001600160401b0382111561483b57600080fd5b60200191503681900382131561473f57600080fd5b634e487b7160e01b600052601260045260246000fd5b600063ffffffff8084168061487d5761487d614850565b92169190910692915050565b63ffffffff828116828216039080821115611faa57611faa614599565b602081526000611efe60208301846145eb565b6000815180845260005b818110156148df576020818501810151868301820152016148c3565b506000602082860101526020601f19601f83011685010191505092915050565b602081526000611efe60208301846148b9565b6001600160e01b03199290921682526001600160a01b0316602082015260400190565b60408152600061494860408301856145eb565b905063ffffffff831660208301529392505050565b6000808335601e1984360301811261497457600080fd5b83016020810192503590506001600160401b0381111561499357600080fd5b80360382131561473f57600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60006149d7828361495d565b608085526149e96080860182846149a2565b9150506149f9602084018461495d565b8583036020870152614a0c8382846149a2565b92506001600160a01b039150614a26905060408501614086565b16604085015261ffff614a3b606085016140fd565b1660608501528091505092915050565b8060005b6002811015610af0578151845260209384019390910190600101614a4f565b60e081526000614a8160e08301886149cb565b8615156020840152614a966040840187614a4b565b614aa36080840186614a4b565b82810360c0840152614ab581856143a0565b98975050505050505050565b60c081526000614ad460c08301886149cb565b86151560208401528281036040840152614aee81876148b9565b9050614afd6060840186614a4b565b82810360a0840152614ab581856143a0565b608081526000614b2360808301888a6149a2565b8281036020840152614b368187896149a2565b6001600160a01b03959095166040840152505061ffff91909116606090910152949350505050565b82815260406020820152600061166660408301846148b9565b600060408284031215614b8957600080fd5b82601f830112614b9857600080fd5b614ba061422a565b806040840185811115614bb257600080fd5b845b81811015614bcc578051845260209384019301614bb4565b509095945050505050565b83815260a08101614beb6020830185614a4b565b6116666060830184614a4b565b602081526000611efe60208301846149cb565b600082614c1a57614c1a614850565b500690565b606081526000614c3260608301866148b9565b8460208401528281036040840152614c4a81856148b9565b9695505050505050565b60e081526000614c6860e083018b8d6149a2565b8281036020840152614c7b818a8c6149a2565b6001600160a01b03988916604085015261ffff979097166060840152505092909416608083015260a082015291151560c090920191909152949350505050565b614cc58187614a4b565b614cd26040820186614a4b565b60e060808201526000614ce860e08301866143a0565b60a08301949094525060c001529392505050565b63ffffffff87168152608060208201526000614d1c6080830187896149a2565b8281036040840152614d2f8186886149a2565b90508281036060840152614d4381856143a0565b9998505050505050505050565b60006001600160401b03831115614d6957614d69614214565b614d7c601f8401601f1916602001614274565b9050828152838383011115614d9057600080fd5b828260208301376000602084830101529392505050565b600082601f830112614db857600080fd5b611efe83833560208501614d50565b600060808236031215614dd957600080fd5b614de1614252565b82356001600160401b0380821115614df857600080fd5b9084019036601f830112614e0b57600080fd5b614e1a36833560208501614d50565b83526020850135915080821115614e3057600080fd5b50614e3d36828601614da7565b602083015250614e4f60408401614086565b6040820152614e60606084016140fd565b606082015292915050565b602081526000825160a060208401528051608060c0850152614e916101408501826148b9565b9050602082015160bf198583030160e0860152614eae82826148b9565b6040848101516001600160a01b031661010088015260609485015161ffff166101208801526020880151878201528701518487015292860151858403601f19016080870152929150614f02905081836143a0565b9150506080840151614f1860a085018215159052565b509392505050565b60006020808385031215614f3357600080fd5b82516001600160401b03811115614f4957600080fd5b8301601f81018513614f5a57600080fd5b8051614f68614312826142a4565b81815260069190911b82018301908381019087831115614f8757600080fd5b928401925b82841015614fca5760408489031215614fa55760008081fd5b614fad61422a565b845181528585015186820152825260409093019290840190614f8c565b979650505050505050565b600063ffffffff808316818103614fee57614fee614599565b6001019392505050565b601f8211156105fe576000816000526020600020601f850160051c810160208610156150215750805b601f850160051c820191505b818110156109915782815560010161502d565b81516001600160401b0381111561505957615059614214565b61506d81615067845461475f565b84614ff8565b602080601f8311600181146150a2576000841561508a5750858301515b600019600386901b1c1916600185901b178555610991565b600085815260208120601f198616915b828110156150d1578886015182559484019460019091019084016150b2565b50858210156150ef5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60408152600061511260408301856148b9565b90508260208301529392505050565b808202811582820484141761064757610647614599565b60008261514757615147614850565b500490565b60008060006060848603121561516157600080fd5b835192506020840151915060408401519050925092509256fefe6065fb4e9872e2ad4479001655335380d83f70e163706cd65857449b84510002dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800546f74616c20566f74696e6720506f7765722053796e63205461736b20446566696e6974696f6ea26469706673582212202d4e6fd6300d2b5aee7e355809f3cebbc02684d5b23cde9dc8daa621f8aa963464736f6c63430008190033",
}

// ContractAttestationCenterABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractAttestationCenterMetaData.ABI instead.
var ContractAttestationCenterABI = ContractAttestationCenterMetaData.ABI

// ContractAttestationCenterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractAttestationCenterMetaData.Bin instead.
var ContractAttestationCenterBin = ContractAttestationCenterMetaData.Bin

// DeployContractAttestationCenter deploys a new Ethereum contract, binding an instance of ContractAttestationCenter to it.
func DeployContractAttestationCenter(auth *bind.TransactOpts, backend bind.ContractBackend, _extensionImplementation common.Address) (common.Address, *types.Transaction, *ContractAttestationCenter, error) {
	parsed, err := ContractAttestationCenterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractAttestationCenterBin), backend, _extensionImplementation)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractAttestationCenter{ContractAttestationCenterCaller: ContractAttestationCenterCaller{contract: contract}, ContractAttestationCenterTransactor: ContractAttestationCenterTransactor{contract: contract}, ContractAttestationCenterFilterer: ContractAttestationCenterFilterer{contract: contract}}, nil
}

// ContractAttestationCenterMethods is an auto generated interface around an Ethereum contract.
type ContractAttestationCenterMethods interface {
	ContractAttestationCenterCalls
	ContractAttestationCenterTransacts
	ContractAttestationCenterFilters
}

// ContractAttestationCenterCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractAttestationCenterCalls interface {
	DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error)

	EXTENSIONIMPLEMENTATION(opts *bind.CallOpts) (common.Address, error)

	AvsLogic(opts *bind.CallOpts) (common.Address, error)

	AvsTreasury(opts *bind.CallOpts) (common.Address, error)

	BeforePaymentsLogic(opts *bind.CallOpts) (common.Address, error)

	GetOperatorPaymentDetail(opts *bind.CallOpts, _operatorId *big.Int) (IAttestationCenterPaymentDetails, error)

	GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error)

	GetTaskDefinitionMaximumNumberOfAttesters(opts *bind.CallOpts, _taskDefinitionId uint16) (*big.Int, error)

	GetTaskDefinitionMinimumVotingPower(opts *bind.CallOpts, _taskDefinitionId uint16) (*big.Int, error)

	GetTaskDefinitionRestrictedAttesters(opts *bind.CallOpts, _taskDefinitionId uint16) ([]*big.Int, error)

	HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error)

	InternalTaskHandler(opts *bind.CallOpts) (common.Address, error)

	IsFlowPaused(opts *bind.CallOpts, _pausableFlow [4]byte) (bool, error)

	NextEigenRewardsBatchStartTimestamp(opts *bind.CallOpts) (*big.Int, error)

	NumOfActiveOperators(opts *bind.CallOpts) (*big.Int, error)

	NumOfTaskDefinitions(opts *bind.CallOpts) (uint16, error)

	NumOfTotalOperators(opts *bind.CallOpts) (*big.Int, error)

	Obls(opts *bind.CallOpts) (common.Address, error)

	OperatorsIdsByAddress(opts *bind.CallOpts, _operator common.Address) (*big.Int, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TaskNumber(opts *bind.CallOpts) (uint32, error)

	VerifyOperatorValidForTaskDefinition(opts *bind.CallOpts, _operator common.Address, _taskDefinitionId uint16) error

	VotingPower(opts *bind.CallOpts, _operator common.Address) (*big.Int, error)
}

// ContractAttestationCenterTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractAttestationCenterTransacts interface {
	ClearBatchPayment(opts *bind.TransactOpts, _operators []IAttestationCenterPaymentRequestMessage, _paidTaskNumber *big.Int) (*types.Transaction, error)

	EjectOperatorFromNetwork(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error)

	GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts, _initializationParams IAttestationCenterInitializationParams) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts, _pausableFlow [4]byte) (*types.Transaction, error)

	RegisterToNetwork(opts *bind.TransactOpts, _operator common.Address, _votingPower *big.Int, _blsKey [4]*big.Int, _rewardsReceiver common.Address) (*types.Transaction, error)

	RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error)

	RequestBatchPayment(opts *bind.TransactOpts, _from *big.Int, _to *big.Int) (*types.Transaction, error)

	RequestBatchPayment0(opts *bind.TransactOpts) (*types.Transaction, error)

	RequestEigenBatchPayment(opts *bind.TransactOpts, _startTimestamp uint32, _duration uint32, _from *big.Int, _to *big.Int) (*types.Transaction, error)

	RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	SetIsRewardsOnL2(opts *bind.TransactOpts, _isRewardsOnL2 bool) (*types.Transaction, error)

	SetOblsSharesSyncer(opts *bind.TransactOpts, _oblsSharesSyncer common.Address) (*types.Transaction, error)

	SubmitTask(opts *bind.TransactOpts, _taskInfo IAttestationCenterTaskInfo, _blsTaskSubmissionDetails IAttestationCenterBlsTaskSubmissionDetails) (*types.Transaction, error)

	SubmitTask0(opts *bind.TransactOpts, _taskInfo IAttestationCenterTaskInfo, _ecdsaTaskSubmissionDetails IAttestationCenterEcdsaTaskSubmissionDetails) (*types.Transaction, error)

	TransferMessageHandler(opts *bind.TransactOpts, _newMessageHandler common.Address) (*types.Transaction, error)

	UnRegisterOperatorFromNetwork(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts, _pausableFlow [4]byte) (*types.Transaction, error)

	UpdateBlsKey(opts *bind.TransactOpts, _blsKey [4]*big.Int, _authSignature BLSAuthLibrarySignature) (*types.Transaction, error)
}

// ContractAttestationCenterFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractAttestationCenterFilters interface {
	FilterClearPaymentRejected(opts *bind.FilterOpts, operator []common.Address) (*ContractAttestationCenterClearPaymentRejectedIterator, error)
	WatchClearPaymentRejected(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterClearPaymentRejected, operator []common.Address) (event.Subscription, error)
	ParseClearPaymentRejected(log types.Log) (*ContractAttestationCenterClearPaymentRejected, error)

	FilterEigenPaymentsRequested(opts *bind.FilterOpts) (*ContractAttestationCenterEigenPaymentsRequestedIterator, error)
	WatchEigenPaymentsRequested(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterEigenPaymentsRequested) (event.Subscription, error)
	ParseEigenPaymentsRequested(log types.Log) (*ContractAttestationCenterEigenPaymentsRequested, error)

	FilterFlowPaused(opts *bind.FilterOpts) (*ContractAttestationCenterFlowPausedIterator, error)
	WatchFlowPaused(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterFlowPaused) (event.Subscription, error)
	ParseFlowPaused(log types.Log) (*ContractAttestationCenterFlowPaused, error)

	FilterFlowUnpaused(opts *bind.FilterOpts) (*ContractAttestationCenterFlowUnpausedIterator, error)
	WatchFlowUnpaused(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterFlowUnpaused) (event.Subscription, error)
	ParseFlowUnpaused(log types.Log) (*ContractAttestationCenterFlowUnpaused, error)

	FilterInitialized(opts *bind.FilterOpts) (*ContractAttestationCenterInitializedIterator, error)
	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterInitialized) (event.Subscription, error)
	ParseInitialized(log types.Log) (*ContractAttestationCenterInitialized, error)

	FilterOperatorBlsKeyUpdated(opts *bind.FilterOpts, operator []common.Address) (*ContractAttestationCenterOperatorBlsKeyUpdatedIterator, error)
	WatchOperatorBlsKeyUpdated(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterOperatorBlsKeyUpdated, operator []common.Address) (event.Subscription, error)
	ParseOperatorBlsKeyUpdated(log types.Log) (*ContractAttestationCenterOperatorBlsKeyUpdated, error)

	FilterOperatorEjectionRequested(opts *bind.FilterOpts) (*ContractAttestationCenterOperatorEjectionRequestedIterator, error)
	WatchOperatorEjectionRequested(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterOperatorEjectionRequested) (event.Subscription, error)
	ParseOperatorEjectionRequested(log types.Log) (*ContractAttestationCenterOperatorEjectionRequested, error)

	FilterOperatorRegisteredToNetwork(opts *bind.FilterOpts, operator []common.Address) (*ContractAttestationCenterOperatorRegisteredToNetworkIterator, error)
	WatchOperatorRegisteredToNetwork(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterOperatorRegisteredToNetwork, operator []common.Address) (event.Subscription, error)
	ParseOperatorRegisteredToNetwork(log types.Log) (*ContractAttestationCenterOperatorRegisteredToNetwork, error)

	FilterOperatorUnregisteredFromNetwork(opts *bind.FilterOpts, operatorId []*big.Int) (*ContractAttestationCenterOperatorUnregisteredFromNetworkIterator, error)
	WatchOperatorUnregisteredFromNetwork(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterOperatorUnregisteredFromNetwork, operatorId []*big.Int) (event.Subscription, error)
	ParseOperatorUnregisteredFromNetwork(log types.Log) (*ContractAttestationCenterOperatorUnregisteredFromNetwork, error)

	FilterPaymentsRequested(opts *bind.FilterOpts) (*ContractAttestationCenterPaymentsRequestedIterator, error)
	WatchPaymentsRequested(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterPaymentsRequested) (event.Subscription, error)
	ParsePaymentsRequested(log types.Log) (*ContractAttestationCenterPaymentsRequested, error)

	FilterRewardAccumulated(opts *bind.FilterOpts, _operatorId []*big.Int, _taskNumber []uint32) (*ContractAttestationCenterRewardAccumulatedIterator, error)
	WatchRewardAccumulated(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterRewardAccumulated, _operatorId []*big.Int, _taskNumber []uint32) (event.Subscription, error)
	ParseRewardAccumulated(log types.Log) (*ContractAttestationCenterRewardAccumulated, error)

	FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ContractAttestationCenterRoleAdminChangedIterator, error)
	WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error)
	ParseRoleAdminChanged(log types.Log) (*ContractAttestationCenterRoleAdminChanged, error)

	FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ContractAttestationCenterRoleGrantedIterator, error)
	WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error)
	ParseRoleGranted(log types.Log) (*ContractAttestationCenterRoleGranted, error)

	FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ContractAttestationCenterRoleRevokedIterator, error)
	WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error)
	ParseRoleRevoked(log types.Log) (*ContractAttestationCenterRoleRevoked, error)

	FilterSetMessageHandler(opts *bind.FilterOpts) (*ContractAttestationCenterSetMessageHandlerIterator, error)
	WatchSetMessageHandler(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterSetMessageHandler) (event.Subscription, error)
	ParseSetMessageHandler(log types.Log) (*ContractAttestationCenterSetMessageHandler, error)

	FilterTaskRejected(opts *bind.FilterOpts, operator []common.Address, taskDefinitionId []uint16) (*ContractAttestationCenterTaskRejectedIterator, error)
	WatchTaskRejected(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterTaskRejected, operator []common.Address, taskDefinitionId []uint16) (event.Subscription, error)
	ParseTaskRejected(log types.Log) (*ContractAttestationCenterTaskRejected, error)

	FilterTaskSubmitted(opts *bind.FilterOpts, operator []common.Address, taskDefinitionId []uint16) (*ContractAttestationCenterTaskSubmittedIterator, error)
	WatchTaskSubmitted(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterTaskSubmitted, operator []common.Address, taskDefinitionId []uint16) (event.Subscription, error)
	ParseTaskSubmitted(log types.Log) (*ContractAttestationCenterTaskSubmitted, error)
}

// ContractAttestationCenter is an auto generated Go binding around an Ethereum contract.
type ContractAttestationCenter struct {
	ContractAttestationCenterCaller     // Read-only binding to the contract
	ContractAttestationCenterTransactor // Write-only binding to the contract
	ContractAttestationCenterFilterer   // Log filterer for contract events
}

// ContractAttestationCenter implements the ContractAttestationCenterMethods interface.
var _ ContractAttestationCenterMethods = (*ContractAttestationCenter)(nil)

// ContractAttestationCenterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractAttestationCenterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAttestationCenterCaller implements the ContractAttestationCenterCalls interface.
var _ ContractAttestationCenterCalls = (*ContractAttestationCenterCaller)(nil)

// ContractAttestationCenterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractAttestationCenterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAttestationCenterTransactor implements the ContractAttestationCenterTransacts interface.
var _ ContractAttestationCenterTransacts = (*ContractAttestationCenterTransactor)(nil)

// ContractAttestationCenterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractAttestationCenterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAttestationCenterFilterer implements the ContractAttestationCenterFilters interface.
var _ ContractAttestationCenterFilters = (*ContractAttestationCenterFilterer)(nil)

// ContractAttestationCenterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractAttestationCenterSession struct {
	Contract     *ContractAttestationCenter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ContractAttestationCenterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractAttestationCenterCallerSession struct {
	Contract *ContractAttestationCenterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// ContractAttestationCenterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractAttestationCenterTransactorSession struct {
	Contract     *ContractAttestationCenterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// ContractAttestationCenterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractAttestationCenterRaw struct {
	Contract *ContractAttestationCenter // Generic contract binding to access the raw methods on
}

// ContractAttestationCenterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractAttestationCenterCallerRaw struct {
	Contract *ContractAttestationCenterCaller // Generic read-only contract binding to access the raw methods on
}

// ContractAttestationCenterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractAttestationCenterTransactorRaw struct {
	Contract *ContractAttestationCenterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractAttestationCenter creates a new instance of ContractAttestationCenter, bound to a specific deployed contract.
func NewContractAttestationCenter(address common.Address, backend bind.ContractBackend) (*ContractAttestationCenter, error) {
	contract, err := bindContractAttestationCenter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenter{ContractAttestationCenterCaller: ContractAttestationCenterCaller{contract: contract}, ContractAttestationCenterTransactor: ContractAttestationCenterTransactor{contract: contract}, ContractAttestationCenterFilterer: ContractAttestationCenterFilterer{contract: contract}}, nil
}

// NewContractAttestationCenterCaller creates a new read-only instance of ContractAttestationCenter, bound to a specific deployed contract.
func NewContractAttestationCenterCaller(address common.Address, caller bind.ContractCaller) (*ContractAttestationCenterCaller, error) {
	contract, err := bindContractAttestationCenter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterCaller{contract: contract}, nil
}

// NewContractAttestationCenterTransactor creates a new write-only instance of ContractAttestationCenter, bound to a specific deployed contract.
func NewContractAttestationCenterTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractAttestationCenterTransactor, error) {
	contract, err := bindContractAttestationCenter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterTransactor{contract: contract}, nil
}

// NewContractAttestationCenterFilterer creates a new log filterer instance of ContractAttestationCenter, bound to a specific deployed contract.
func NewContractAttestationCenterFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractAttestationCenterFilterer, error) {
	contract, err := bindContractAttestationCenter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterFilterer{contract: contract}, nil
}

// bindContractAttestationCenter binds a generic wrapper to an already deployed contract.
func bindContractAttestationCenter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractAttestationCenterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractAttestationCenter *ContractAttestationCenterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractAttestationCenter.Contract.ContractAttestationCenterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractAttestationCenter *ContractAttestationCenterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.ContractAttestationCenterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractAttestationCenter *ContractAttestationCenterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.ContractAttestationCenterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractAttestationCenter *ContractAttestationCenterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractAttestationCenter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractAttestationCenter *ContractAttestationCenterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractAttestationCenter *ContractAttestationCenterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ContractAttestationCenter *ContractAttestationCenterCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractAttestationCenter.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ContractAttestationCenter *ContractAttestationCenterSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ContractAttestationCenter.Contract.DEFAULTADMINROLE(&_ContractAttestationCenter.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ContractAttestationCenter *ContractAttestationCenterCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ContractAttestationCenter.Contract.DEFAULTADMINROLE(&_ContractAttestationCenter.CallOpts)
}

// EXTENSIONIMPLEMENTATION is a free data retrieval call binding the contract method 0x226def04.
//
// Solidity: function EXTENSION_IMPLEMENTATION() view returns(address)
func (_ContractAttestationCenter *ContractAttestationCenterCaller) EXTENSIONIMPLEMENTATION(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAttestationCenter.contract.Call(opts, &out, "EXTENSION_IMPLEMENTATION")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EXTENSIONIMPLEMENTATION is a free data retrieval call binding the contract method 0x226def04.
//
// Solidity: function EXTENSION_IMPLEMENTATION() view returns(address)
func (_ContractAttestationCenter *ContractAttestationCenterSession) EXTENSIONIMPLEMENTATION() (common.Address, error) {
	return _ContractAttestationCenter.Contract.EXTENSIONIMPLEMENTATION(&_ContractAttestationCenter.CallOpts)
}

// EXTENSIONIMPLEMENTATION is a free data retrieval call binding the contract method 0x226def04.
//
// Solidity: function EXTENSION_IMPLEMENTATION() view returns(address)
func (_ContractAttestationCenter *ContractAttestationCenterCallerSession) EXTENSIONIMPLEMENTATION() (common.Address, error) {
	return _ContractAttestationCenter.Contract.EXTENSIONIMPLEMENTATION(&_ContractAttestationCenter.CallOpts)
}

// AvsLogic is a free data retrieval call binding the contract method 0xb0817c44.
//
// Solidity: function avsLogic() view returns(address)
func (_ContractAttestationCenter *ContractAttestationCenterCaller) AvsLogic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAttestationCenter.contract.Call(opts, &out, "avsLogic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsLogic is a free data retrieval call binding the contract method 0xb0817c44.
//
// Solidity: function avsLogic() view returns(address)
func (_ContractAttestationCenter *ContractAttestationCenterSession) AvsLogic() (common.Address, error) {
	return _ContractAttestationCenter.Contract.AvsLogic(&_ContractAttestationCenter.CallOpts)
}

// AvsLogic is a free data retrieval call binding the contract method 0xb0817c44.
//
// Solidity: function avsLogic() view returns(address)
func (_ContractAttestationCenter *ContractAttestationCenterCallerSession) AvsLogic() (common.Address, error) {
	return _ContractAttestationCenter.Contract.AvsLogic(&_ContractAttestationCenter.CallOpts)
}

// AvsTreasury is a free data retrieval call binding the contract method 0x1246193e.
//
// Solidity: function avsTreasury() view returns(address)
func (_ContractAttestationCenter *ContractAttestationCenterCaller) AvsTreasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAttestationCenter.contract.Call(opts, &out, "avsTreasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsTreasury is a free data retrieval call binding the contract method 0x1246193e.
//
// Solidity: function avsTreasury() view returns(address)
func (_ContractAttestationCenter *ContractAttestationCenterSession) AvsTreasury() (common.Address, error) {
	return _ContractAttestationCenter.Contract.AvsTreasury(&_ContractAttestationCenter.CallOpts)
}

// AvsTreasury is a free data retrieval call binding the contract method 0x1246193e.
//
// Solidity: function avsTreasury() view returns(address)
func (_ContractAttestationCenter *ContractAttestationCenterCallerSession) AvsTreasury() (common.Address, error) {
	return _ContractAttestationCenter.Contract.AvsTreasury(&_ContractAttestationCenter.CallOpts)
}

// BeforePaymentsLogic is a free data retrieval call binding the contract method 0xc2f429f1.
//
// Solidity: function beforePaymentsLogic() view returns(address)
func (_ContractAttestationCenter *ContractAttestationCenterCaller) BeforePaymentsLogic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAttestationCenter.contract.Call(opts, &out, "beforePaymentsLogic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BeforePaymentsLogic is a free data retrieval call binding the contract method 0xc2f429f1.
//
// Solidity: function beforePaymentsLogic() view returns(address)
func (_ContractAttestationCenter *ContractAttestationCenterSession) BeforePaymentsLogic() (common.Address, error) {
	return _ContractAttestationCenter.Contract.BeforePaymentsLogic(&_ContractAttestationCenter.CallOpts)
}

// BeforePaymentsLogic is a free data retrieval call binding the contract method 0xc2f429f1.
//
// Solidity: function beforePaymentsLogic() view returns(address)
func (_ContractAttestationCenter *ContractAttestationCenterCallerSession) BeforePaymentsLogic() (common.Address, error) {
	return _ContractAttestationCenter.Contract.BeforePaymentsLogic(&_ContractAttestationCenter.CallOpts)
}

// GetOperatorPaymentDetail is a free data retrieval call binding the contract method 0x9eb72d4c.
//
// Solidity: function getOperatorPaymentDetail(uint256 _operatorId) view returns((address,uint256,uint256,uint8))
func (_ContractAttestationCenter *ContractAttestationCenterCaller) GetOperatorPaymentDetail(opts *bind.CallOpts, _operatorId *big.Int) (IAttestationCenterPaymentDetails, error) {
	var out []interface{}
	err := _ContractAttestationCenter.contract.Call(opts, &out, "getOperatorPaymentDetail", _operatorId)

	if err != nil {
		return *new(IAttestationCenterPaymentDetails), err
	}

	out0 := *abi.ConvertType(out[0], new(IAttestationCenterPaymentDetails)).(*IAttestationCenterPaymentDetails)

	return out0, err

}

// GetOperatorPaymentDetail is a free data retrieval call binding the contract method 0x9eb72d4c.
//
// Solidity: function getOperatorPaymentDetail(uint256 _operatorId) view returns((address,uint256,uint256,uint8))
func (_ContractAttestationCenter *ContractAttestationCenterSession) GetOperatorPaymentDetail(_operatorId *big.Int) (IAttestationCenterPaymentDetails, error) {
	return _ContractAttestationCenter.Contract.GetOperatorPaymentDetail(&_ContractAttestationCenter.CallOpts, _operatorId)
}

// GetOperatorPaymentDetail is a free data retrieval call binding the contract method 0x9eb72d4c.
//
// Solidity: function getOperatorPaymentDetail(uint256 _operatorId) view returns((address,uint256,uint256,uint8))
func (_ContractAttestationCenter *ContractAttestationCenterCallerSession) GetOperatorPaymentDetail(_operatorId *big.Int) (IAttestationCenterPaymentDetails, error) {
	return _ContractAttestationCenter.Contract.GetOperatorPaymentDetail(&_ContractAttestationCenter.CallOpts, _operatorId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ContractAttestationCenter *ContractAttestationCenterCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ContractAttestationCenter.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ContractAttestationCenter *ContractAttestationCenterSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ContractAttestationCenter.Contract.GetRoleAdmin(&_ContractAttestationCenter.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ContractAttestationCenter *ContractAttestationCenterCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ContractAttestationCenter.Contract.GetRoleAdmin(&_ContractAttestationCenter.CallOpts, role)
}

// GetTaskDefinitionMaximumNumberOfAttesters is a free data retrieval call binding the contract method 0x5b386be0.
//
// Solidity: function getTaskDefinitionMaximumNumberOfAttesters(uint16 _taskDefinitionId) view returns(uint256)
func (_ContractAttestationCenter *ContractAttestationCenterCaller) GetTaskDefinitionMaximumNumberOfAttesters(opts *bind.CallOpts, _taskDefinitionId uint16) (*big.Int, error) {
	var out []interface{}
	err := _ContractAttestationCenter.contract.Call(opts, &out, "getTaskDefinitionMaximumNumberOfAttesters", _taskDefinitionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTaskDefinitionMaximumNumberOfAttesters is a free data retrieval call binding the contract method 0x5b386be0.
//
// Solidity: function getTaskDefinitionMaximumNumberOfAttesters(uint16 _taskDefinitionId) view returns(uint256)
func (_ContractAttestationCenter *ContractAttestationCenterSession) GetTaskDefinitionMaximumNumberOfAttesters(_taskDefinitionId uint16) (*big.Int, error) {
	return _ContractAttestationCenter.Contract.GetTaskDefinitionMaximumNumberOfAttesters(&_ContractAttestationCenter.CallOpts, _taskDefinitionId)
}

// GetTaskDefinitionMaximumNumberOfAttesters is a free data retrieval call binding the contract method 0x5b386be0.
//
// Solidity: function getTaskDefinitionMaximumNumberOfAttesters(uint16 _taskDefinitionId) view returns(uint256)
func (_ContractAttestationCenter *ContractAttestationCenterCallerSession) GetTaskDefinitionMaximumNumberOfAttesters(_taskDefinitionId uint16) (*big.Int, error) {
	return _ContractAttestationCenter.Contract.GetTaskDefinitionMaximumNumberOfAttesters(&_ContractAttestationCenter.CallOpts, _taskDefinitionId)
}

// GetTaskDefinitionMinimumVotingPower is a free data retrieval call binding the contract method 0x75d9aedf.
//
// Solidity: function getTaskDefinitionMinimumVotingPower(uint16 _taskDefinitionId) view returns(uint256)
func (_ContractAttestationCenter *ContractAttestationCenterCaller) GetTaskDefinitionMinimumVotingPower(opts *bind.CallOpts, _taskDefinitionId uint16) (*big.Int, error) {
	var out []interface{}
	err := _ContractAttestationCenter.contract.Call(opts, &out, "getTaskDefinitionMinimumVotingPower", _taskDefinitionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTaskDefinitionMinimumVotingPower is a free data retrieval call binding the contract method 0x75d9aedf.
//
// Solidity: function getTaskDefinitionMinimumVotingPower(uint16 _taskDefinitionId) view returns(uint256)
func (_ContractAttestationCenter *ContractAttestationCenterSession) GetTaskDefinitionMinimumVotingPower(_taskDefinitionId uint16) (*big.Int, error) {
	return _ContractAttestationCenter.Contract.GetTaskDefinitionMinimumVotingPower(&_ContractAttestationCenter.CallOpts, _taskDefinitionId)
}

// GetTaskDefinitionMinimumVotingPower is a free data retrieval call binding the contract method 0x75d9aedf.
//
// Solidity: function getTaskDefinitionMinimumVotingPower(uint16 _taskDefinitionId) view returns(uint256)
func (_ContractAttestationCenter *ContractAttestationCenterCallerSession) GetTaskDefinitionMinimumVotingPower(_taskDefinitionId uint16) (*big.Int, error) {
	return _ContractAttestationCenter.Contract.GetTaskDefinitionMinimumVotingPower(&_ContractAttestationCenter.CallOpts, _taskDefinitionId)
}

// GetTaskDefinitionRestrictedAttesters is a free data retrieval call binding the contract method 0x97b5f370.
//
// Solidity: function getTaskDefinitionRestrictedAttesters(uint16 _taskDefinitionId) view returns(uint256[])
func (_ContractAttestationCenter *ContractAttestationCenterCaller) GetTaskDefinitionRestrictedAttesters(opts *bind.CallOpts, _taskDefinitionId uint16) ([]*big.Int, error) {
	var out []interface{}
	err := _ContractAttestationCenter.contract.Call(opts, &out, "getTaskDefinitionRestrictedAttesters", _taskDefinitionId)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetTaskDefinitionRestrictedAttesters is a free data retrieval call binding the contract method 0x97b5f370.
//
// Solidity: function getTaskDefinitionRestrictedAttesters(uint16 _taskDefinitionId) view returns(uint256[])
func (_ContractAttestationCenter *ContractAttestationCenterSession) GetTaskDefinitionRestrictedAttesters(_taskDefinitionId uint16) ([]*big.Int, error) {
	return _ContractAttestationCenter.Contract.GetTaskDefinitionRestrictedAttesters(&_ContractAttestationCenter.CallOpts, _taskDefinitionId)
}

// GetTaskDefinitionRestrictedAttesters is a free data retrieval call binding the contract method 0x97b5f370.
//
// Solidity: function getTaskDefinitionRestrictedAttesters(uint16 _taskDefinitionId) view returns(uint256[])
func (_ContractAttestationCenter *ContractAttestationCenterCallerSession) GetTaskDefinitionRestrictedAttesters(_taskDefinitionId uint16) ([]*big.Int, error) {
	return _ContractAttestationCenter.Contract.GetTaskDefinitionRestrictedAttesters(&_ContractAttestationCenter.CallOpts, _taskDefinitionId)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ContractAttestationCenter *ContractAttestationCenterCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ContractAttestationCenter.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ContractAttestationCenter *ContractAttestationCenterSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ContractAttestationCenter.Contract.HasRole(&_ContractAttestationCenter.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ContractAttestationCenter *ContractAttestationCenterCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ContractAttestationCenter.Contract.HasRole(&_ContractAttestationCenter.CallOpts, role, account)
}

// InternalTaskHandler is a free data retrieval call binding the contract method 0xf28eff39.
//
// Solidity: function internalTaskHandler() view returns(address)
func (_ContractAttestationCenter *ContractAttestationCenterCaller) InternalTaskHandler(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAttestationCenter.contract.Call(opts, &out, "internalTaskHandler")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InternalTaskHandler is a free data retrieval call binding the contract method 0xf28eff39.
//
// Solidity: function internalTaskHandler() view returns(address)
func (_ContractAttestationCenter *ContractAttestationCenterSession) InternalTaskHandler() (common.Address, error) {
	return _ContractAttestationCenter.Contract.InternalTaskHandler(&_ContractAttestationCenter.CallOpts)
}

// InternalTaskHandler is a free data retrieval call binding the contract method 0xf28eff39.
//
// Solidity: function internalTaskHandler() view returns(address)
func (_ContractAttestationCenter *ContractAttestationCenterCallerSession) InternalTaskHandler() (common.Address, error) {
	return _ContractAttestationCenter.Contract.InternalTaskHandler(&_ContractAttestationCenter.CallOpts)
}

// IsFlowPaused is a free data retrieval call binding the contract method 0xefd96978.
//
// Solidity: function isFlowPaused(bytes4 _pausableFlow) view returns(bool _isPaused)
func (_ContractAttestationCenter *ContractAttestationCenterCaller) IsFlowPaused(opts *bind.CallOpts, _pausableFlow [4]byte) (bool, error) {
	var out []interface{}
	err := _ContractAttestationCenter.contract.Call(opts, &out, "isFlowPaused", _pausableFlow)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFlowPaused is a free data retrieval call binding the contract method 0xefd96978.
//
// Solidity: function isFlowPaused(bytes4 _pausableFlow) view returns(bool _isPaused)
func (_ContractAttestationCenter *ContractAttestationCenterSession) IsFlowPaused(_pausableFlow [4]byte) (bool, error) {
	return _ContractAttestationCenter.Contract.IsFlowPaused(&_ContractAttestationCenter.CallOpts, _pausableFlow)
}

// IsFlowPaused is a free data retrieval call binding the contract method 0xefd96978.
//
// Solidity: function isFlowPaused(bytes4 _pausableFlow) view returns(bool _isPaused)
func (_ContractAttestationCenter *ContractAttestationCenterCallerSession) IsFlowPaused(_pausableFlow [4]byte) (bool, error) {
	return _ContractAttestationCenter.Contract.IsFlowPaused(&_ContractAttestationCenter.CallOpts, _pausableFlow)
}

// NextEigenRewardsBatchStartTimestamp is a free data retrieval call binding the contract method 0x8355e33d.
//
// Solidity: function nextEigenRewardsBatchStartTimestamp() view returns(uint256)
func (_ContractAttestationCenter *ContractAttestationCenterCaller) NextEigenRewardsBatchStartTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractAttestationCenter.contract.Call(opts, &out, "nextEigenRewardsBatchStartTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextEigenRewardsBatchStartTimestamp is a free data retrieval call binding the contract method 0x8355e33d.
//
// Solidity: function nextEigenRewardsBatchStartTimestamp() view returns(uint256)
func (_ContractAttestationCenter *ContractAttestationCenterSession) NextEigenRewardsBatchStartTimestamp() (*big.Int, error) {
	return _ContractAttestationCenter.Contract.NextEigenRewardsBatchStartTimestamp(&_ContractAttestationCenter.CallOpts)
}

// NextEigenRewardsBatchStartTimestamp is a free data retrieval call binding the contract method 0x8355e33d.
//
// Solidity: function nextEigenRewardsBatchStartTimestamp() view returns(uint256)
func (_ContractAttestationCenter *ContractAttestationCenterCallerSession) NextEigenRewardsBatchStartTimestamp() (*big.Int, error) {
	return _ContractAttestationCenter.Contract.NextEigenRewardsBatchStartTimestamp(&_ContractAttestationCenter.CallOpts)
}

// NumOfActiveOperators is a free data retrieval call binding the contract method 0x7897dec3.
//
// Solidity: function numOfActiveOperators() view returns(uint256)
func (_ContractAttestationCenter *ContractAttestationCenterCaller) NumOfActiveOperators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractAttestationCenter.contract.Call(opts, &out, "numOfActiveOperators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumOfActiveOperators is a free data retrieval call binding the contract method 0x7897dec3.
//
// Solidity: function numOfActiveOperators() view returns(uint256)
func (_ContractAttestationCenter *ContractAttestationCenterSession) NumOfActiveOperators() (*big.Int, error) {
	return _ContractAttestationCenter.Contract.NumOfActiveOperators(&_ContractAttestationCenter.CallOpts)
}

// NumOfActiveOperators is a free data retrieval call binding the contract method 0x7897dec3.
//
// Solidity: function numOfActiveOperators() view returns(uint256)
func (_ContractAttestationCenter *ContractAttestationCenterCallerSession) NumOfActiveOperators() (*big.Int, error) {
	return _ContractAttestationCenter.Contract.NumOfActiveOperators(&_ContractAttestationCenter.CallOpts)
}

// NumOfTaskDefinitions is a free data retrieval call binding the contract method 0x34a7c391.
//
// Solidity: function numOfTaskDefinitions() view returns(uint16)
func (_ContractAttestationCenter *ContractAttestationCenterCaller) NumOfTaskDefinitions(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ContractAttestationCenter.contract.Call(opts, &out, "numOfTaskDefinitions")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// NumOfTaskDefinitions is a free data retrieval call binding the contract method 0x34a7c391.
//
// Solidity: function numOfTaskDefinitions() view returns(uint16)
func (_ContractAttestationCenter *ContractAttestationCenterSession) NumOfTaskDefinitions() (uint16, error) {
	return _ContractAttestationCenter.Contract.NumOfTaskDefinitions(&_ContractAttestationCenter.CallOpts)
}

// NumOfTaskDefinitions is a free data retrieval call binding the contract method 0x34a7c391.
//
// Solidity: function numOfTaskDefinitions() view returns(uint16)
func (_ContractAttestationCenter *ContractAttestationCenterCallerSession) NumOfTaskDefinitions() (uint16, error) {
	return _ContractAttestationCenter.Contract.NumOfTaskDefinitions(&_ContractAttestationCenter.CallOpts)
}

// NumOfTotalOperators is a free data retrieval call binding the contract method 0x00028b07.
//
// Solidity: function numOfTotalOperators() view returns(uint256)
func (_ContractAttestationCenter *ContractAttestationCenterCaller) NumOfTotalOperators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractAttestationCenter.contract.Call(opts, &out, "numOfTotalOperators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumOfTotalOperators is a free data retrieval call binding the contract method 0x00028b07.
//
// Solidity: function numOfTotalOperators() view returns(uint256)
func (_ContractAttestationCenter *ContractAttestationCenterSession) NumOfTotalOperators() (*big.Int, error) {
	return _ContractAttestationCenter.Contract.NumOfTotalOperators(&_ContractAttestationCenter.CallOpts)
}

// NumOfTotalOperators is a free data retrieval call binding the contract method 0x00028b07.
//
// Solidity: function numOfTotalOperators() view returns(uint256)
func (_ContractAttestationCenter *ContractAttestationCenterCallerSession) NumOfTotalOperators() (*big.Int, error) {
	return _ContractAttestationCenter.Contract.NumOfTotalOperators(&_ContractAttestationCenter.CallOpts)
}

// Obls is a free data retrieval call binding the contract method 0x659fa976.
//
// Solidity: function obls() view returns(address)
func (_ContractAttestationCenter *ContractAttestationCenterCaller) Obls(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAttestationCenter.contract.Call(opts, &out, "obls")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Obls is a free data retrieval call binding the contract method 0x659fa976.
//
// Solidity: function obls() view returns(address)
func (_ContractAttestationCenter *ContractAttestationCenterSession) Obls() (common.Address, error) {
	return _ContractAttestationCenter.Contract.Obls(&_ContractAttestationCenter.CallOpts)
}

// Obls is a free data retrieval call binding the contract method 0x659fa976.
//
// Solidity: function obls() view returns(address)
func (_ContractAttestationCenter *ContractAttestationCenterCallerSession) Obls() (common.Address, error) {
	return _ContractAttestationCenter.Contract.Obls(&_ContractAttestationCenter.CallOpts)
}

// OperatorsIdsByAddress is a free data retrieval call binding the contract method 0x5b15c568.
//
// Solidity: function operatorsIdsByAddress(address _operator) view returns(uint256)
func (_ContractAttestationCenter *ContractAttestationCenterCaller) OperatorsIdsByAddress(opts *bind.CallOpts, _operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractAttestationCenter.contract.Call(opts, &out, "operatorsIdsByAddress", _operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorsIdsByAddress is a free data retrieval call binding the contract method 0x5b15c568.
//
// Solidity: function operatorsIdsByAddress(address _operator) view returns(uint256)
func (_ContractAttestationCenter *ContractAttestationCenterSession) OperatorsIdsByAddress(_operator common.Address) (*big.Int, error) {
	return _ContractAttestationCenter.Contract.OperatorsIdsByAddress(&_ContractAttestationCenter.CallOpts, _operator)
}

// OperatorsIdsByAddress is a free data retrieval call binding the contract method 0x5b15c568.
//
// Solidity: function operatorsIdsByAddress(address _operator) view returns(uint256)
func (_ContractAttestationCenter *ContractAttestationCenterCallerSession) OperatorsIdsByAddress(_operator common.Address) (*big.Int, error) {
	return _ContractAttestationCenter.Contract.OperatorsIdsByAddress(&_ContractAttestationCenter.CallOpts, _operator)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ContractAttestationCenter *ContractAttestationCenterCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ContractAttestationCenter.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ContractAttestationCenter *ContractAttestationCenterSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ContractAttestationCenter.Contract.SupportsInterface(&_ContractAttestationCenter.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ContractAttestationCenter *ContractAttestationCenterCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ContractAttestationCenter.Contract.SupportsInterface(&_ContractAttestationCenter.CallOpts, interfaceId)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractAttestationCenter *ContractAttestationCenterCaller) TaskNumber(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractAttestationCenter.contract.Call(opts, &out, "taskNumber")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractAttestationCenter *ContractAttestationCenterSession) TaskNumber() (uint32, error) {
	return _ContractAttestationCenter.Contract.TaskNumber(&_ContractAttestationCenter.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractAttestationCenter *ContractAttestationCenterCallerSession) TaskNumber() (uint32, error) {
	return _ContractAttestationCenter.Contract.TaskNumber(&_ContractAttestationCenter.CallOpts)
}

// VerifyOperatorValidForTaskDefinition is a free data retrieval call binding the contract method 0xac8c16cf.
//
// Solidity: function verifyOperatorValidForTaskDefinition(address _operator, uint16 _taskDefinitionId) view returns()
func (_ContractAttestationCenter *ContractAttestationCenterCaller) VerifyOperatorValidForTaskDefinition(opts *bind.CallOpts, _operator common.Address, _taskDefinitionId uint16) error {
	var out []interface{}
	err := _ContractAttestationCenter.contract.Call(opts, &out, "verifyOperatorValidForTaskDefinition", _operator, _taskDefinitionId)

	if err != nil {
		return err
	}

	return err

}

// VerifyOperatorValidForTaskDefinition is a free data retrieval call binding the contract method 0xac8c16cf.
//
// Solidity: function verifyOperatorValidForTaskDefinition(address _operator, uint16 _taskDefinitionId) view returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) VerifyOperatorValidForTaskDefinition(_operator common.Address, _taskDefinitionId uint16) error {
	return _ContractAttestationCenter.Contract.VerifyOperatorValidForTaskDefinition(&_ContractAttestationCenter.CallOpts, _operator, _taskDefinitionId)
}

// VerifyOperatorValidForTaskDefinition is a free data retrieval call binding the contract method 0xac8c16cf.
//
// Solidity: function verifyOperatorValidForTaskDefinition(address _operator, uint16 _taskDefinitionId) view returns()
func (_ContractAttestationCenter *ContractAttestationCenterCallerSession) VerifyOperatorValidForTaskDefinition(_operator common.Address, _taskDefinitionId uint16) error {
	return _ContractAttestationCenter.Contract.VerifyOperatorValidForTaskDefinition(&_ContractAttestationCenter.CallOpts, _operator, _taskDefinitionId)
}

// VotingPower is a free data retrieval call binding the contract method 0xc07473f6.
//
// Solidity: function votingPower(address _operator) view returns(uint256)
func (_ContractAttestationCenter *ContractAttestationCenterCaller) VotingPower(opts *bind.CallOpts, _operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractAttestationCenter.contract.Call(opts, &out, "votingPower", _operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingPower is a free data retrieval call binding the contract method 0xc07473f6.
//
// Solidity: function votingPower(address _operator) view returns(uint256)
func (_ContractAttestationCenter *ContractAttestationCenterSession) VotingPower(_operator common.Address) (*big.Int, error) {
	return _ContractAttestationCenter.Contract.VotingPower(&_ContractAttestationCenter.CallOpts, _operator)
}

// VotingPower is a free data retrieval call binding the contract method 0xc07473f6.
//
// Solidity: function votingPower(address _operator) view returns(uint256)
func (_ContractAttestationCenter *ContractAttestationCenterCallerSession) VotingPower(_operator common.Address) (*big.Int, error) {
	return _ContractAttestationCenter.Contract.VotingPower(&_ContractAttestationCenter.CallOpts, _operator)
}

// ClearBatchPayment is a paid mutator transaction binding the contract method 0x915359fc.
//
// Solidity: function clearBatchPayment((address,uint256)[] _operators, uint256 _paidTaskNumber) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) ClearBatchPayment(opts *bind.TransactOpts, _operators []IAttestationCenterPaymentRequestMessage, _paidTaskNumber *big.Int) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "clearBatchPayment", _operators, _paidTaskNumber)
}

// ClearBatchPayment is a paid mutator transaction binding the contract method 0x915359fc.
//
// Solidity: function clearBatchPayment((address,uint256)[] _operators, uint256 _paidTaskNumber) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) ClearBatchPayment(_operators []IAttestationCenterPaymentRequestMessage, _paidTaskNumber *big.Int) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.ClearBatchPayment(&_ContractAttestationCenter.TransactOpts, _operators, _paidTaskNumber)
}

// ClearBatchPayment is a paid mutator transaction binding the contract method 0x915359fc.
//
// Solidity: function clearBatchPayment((address,uint256)[] _operators, uint256 _paidTaskNumber) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) ClearBatchPayment(_operators []IAttestationCenterPaymentRequestMessage, _paidTaskNumber *big.Int) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.ClearBatchPayment(&_ContractAttestationCenter.TransactOpts, _operators, _paidTaskNumber)
}

// EjectOperatorFromNetwork is a paid mutator transaction binding the contract method 0xf6e258cc.
//
// Solidity: function ejectOperatorFromNetwork(address _operator) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) EjectOperatorFromNetwork(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "ejectOperatorFromNetwork", _operator)
}

// EjectOperatorFromNetwork is a paid mutator transaction binding the contract method 0xf6e258cc.
//
// Solidity: function ejectOperatorFromNetwork(address _operator) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) EjectOperatorFromNetwork(_operator common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.EjectOperatorFromNetwork(&_ContractAttestationCenter.TransactOpts, _operator)
}

// EjectOperatorFromNetwork is a paid mutator transaction binding the contract method 0xf6e258cc.
//
// Solidity: function ejectOperatorFromNetwork(address _operator) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) EjectOperatorFromNetwork(_operator common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.EjectOperatorFromNetwork(&_ContractAttestationCenter.TransactOpts, _operator)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.GrantRole(&_ContractAttestationCenter.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.GrantRole(&_ContractAttestationCenter.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x95983fee.
//
// Solidity: function initialize((address,address,address,address,address,address,bool,address) _initializationParams) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) Initialize(opts *bind.TransactOpts, _initializationParams IAttestationCenterInitializationParams) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "initialize", _initializationParams)
}

// Initialize is a paid mutator transaction binding the contract method 0x95983fee.
//
// Solidity: function initialize((address,address,address,address,address,address,bool,address) _initializationParams) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) Initialize(_initializationParams IAttestationCenterInitializationParams) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.Initialize(&_ContractAttestationCenter.TransactOpts, _initializationParams)
}

// Initialize is a paid mutator transaction binding the contract method 0x95983fee.
//
// Solidity: function initialize((address,address,address,address,address,address,bool,address) _initializationParams) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) Initialize(_initializationParams IAttestationCenterInitializationParams) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.Initialize(&_ContractAttestationCenter.TransactOpts, _initializationParams)
}

// Pause is a paid mutator transaction binding the contract method 0x3aa83ec7.
//
// Solidity: function pause(bytes4 _pausableFlow) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) Pause(opts *bind.TransactOpts, _pausableFlow [4]byte) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "pause", _pausableFlow)
}

// Pause is a paid mutator transaction binding the contract method 0x3aa83ec7.
//
// Solidity: function pause(bytes4 _pausableFlow) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) Pause(_pausableFlow [4]byte) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.Pause(&_ContractAttestationCenter.TransactOpts, _pausableFlow)
}

// Pause is a paid mutator transaction binding the contract method 0x3aa83ec7.
//
// Solidity: function pause(bytes4 _pausableFlow) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) Pause(_pausableFlow [4]byte) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.Pause(&_ContractAttestationCenter.TransactOpts, _pausableFlow)
}

// RegisterToNetwork is a paid mutator transaction binding the contract method 0xfcd4e66a.
//
// Solidity: function registerToNetwork(address _operator, uint256 _votingPower, uint256[4] _blsKey, address _rewardsReceiver) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) RegisterToNetwork(opts *bind.TransactOpts, _operator common.Address, _votingPower *big.Int, _blsKey [4]*big.Int, _rewardsReceiver common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "registerToNetwork", _operator, _votingPower, _blsKey, _rewardsReceiver)
}

// RegisterToNetwork is a paid mutator transaction binding the contract method 0xfcd4e66a.
//
// Solidity: function registerToNetwork(address _operator, uint256 _votingPower, uint256[4] _blsKey, address _rewardsReceiver) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) RegisterToNetwork(_operator common.Address, _votingPower *big.Int, _blsKey [4]*big.Int, _rewardsReceiver common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.RegisterToNetwork(&_ContractAttestationCenter.TransactOpts, _operator, _votingPower, _blsKey, _rewardsReceiver)
}

// RegisterToNetwork is a paid mutator transaction binding the contract method 0xfcd4e66a.
//
// Solidity: function registerToNetwork(address _operator, uint256 _votingPower, uint256[4] _blsKey, address _rewardsReceiver) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) RegisterToNetwork(_operator common.Address, _votingPower *big.Int, _blsKey [4]*big.Int, _rewardsReceiver common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.RegisterToNetwork(&_ContractAttestationCenter.TransactOpts, _operator, _votingPower, _blsKey, _rewardsReceiver)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.RenounceRole(&_ContractAttestationCenter.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.RenounceRole(&_ContractAttestationCenter.TransactOpts, role, callerConfirmation)
}

// RequestBatchPayment is a paid mutator transaction binding the contract method 0x6f382619.
//
// Solidity: function requestBatchPayment(uint256 _from, uint256 _to) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) RequestBatchPayment(opts *bind.TransactOpts, _from *big.Int, _to *big.Int) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "requestBatchPayment", _from, _to)
}

// RequestBatchPayment is a paid mutator transaction binding the contract method 0x6f382619.
//
// Solidity: function requestBatchPayment(uint256 _from, uint256 _to) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) RequestBatchPayment(_from *big.Int, _to *big.Int) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.RequestBatchPayment(&_ContractAttestationCenter.TransactOpts, _from, _to)
}

// RequestBatchPayment is a paid mutator transaction binding the contract method 0x6f382619.
//
// Solidity: function requestBatchPayment(uint256 _from, uint256 _to) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) RequestBatchPayment(_from *big.Int, _to *big.Int) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.RequestBatchPayment(&_ContractAttestationCenter.TransactOpts, _from, _to)
}

// RequestBatchPayment0 is a paid mutator transaction binding the contract method 0xb7aa2fdf.
//
// Solidity: function requestBatchPayment() returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) RequestBatchPayment0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "requestBatchPayment0")
}

// RequestBatchPayment0 is a paid mutator transaction binding the contract method 0xb7aa2fdf.
//
// Solidity: function requestBatchPayment() returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) RequestBatchPayment0() (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.RequestBatchPayment0(&_ContractAttestationCenter.TransactOpts)
}

// RequestBatchPayment0 is a paid mutator transaction binding the contract method 0xb7aa2fdf.
//
// Solidity: function requestBatchPayment() returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) RequestBatchPayment0() (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.RequestBatchPayment0(&_ContractAttestationCenter.TransactOpts)
}

// RequestEigenBatchPayment is a paid mutator transaction binding the contract method 0x0a1e7093.
//
// Solidity: function requestEigenBatchPayment(uint32 _startTimestamp, uint32 _duration, uint256 _from, uint256 _to) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) RequestEigenBatchPayment(opts *bind.TransactOpts, _startTimestamp uint32, _duration uint32, _from *big.Int, _to *big.Int) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "requestEigenBatchPayment", _startTimestamp, _duration, _from, _to)
}

// RequestEigenBatchPayment is a paid mutator transaction binding the contract method 0x0a1e7093.
//
// Solidity: function requestEigenBatchPayment(uint32 _startTimestamp, uint32 _duration, uint256 _from, uint256 _to) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) RequestEigenBatchPayment(_startTimestamp uint32, _duration uint32, _from *big.Int, _to *big.Int) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.RequestEigenBatchPayment(&_ContractAttestationCenter.TransactOpts, _startTimestamp, _duration, _from, _to)
}

// RequestEigenBatchPayment is a paid mutator transaction binding the contract method 0x0a1e7093.
//
// Solidity: function requestEigenBatchPayment(uint32 _startTimestamp, uint32 _duration, uint256 _from, uint256 _to) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) RequestEigenBatchPayment(_startTimestamp uint32, _duration uint32, _from *big.Int, _to *big.Int) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.RequestEigenBatchPayment(&_ContractAttestationCenter.TransactOpts, _startTimestamp, _duration, _from, _to)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.RevokeRole(&_ContractAttestationCenter.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.RevokeRole(&_ContractAttestationCenter.TransactOpts, role, account)
}

// SetIsRewardsOnL2 is a paid mutator transaction binding the contract method 0xdba8323b.
//
// Solidity: function setIsRewardsOnL2(bool _isRewardsOnL2) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) SetIsRewardsOnL2(opts *bind.TransactOpts, _isRewardsOnL2 bool) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "setIsRewardsOnL2", _isRewardsOnL2)
}

// SetIsRewardsOnL2 is a paid mutator transaction binding the contract method 0xdba8323b.
//
// Solidity: function setIsRewardsOnL2(bool _isRewardsOnL2) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) SetIsRewardsOnL2(_isRewardsOnL2 bool) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SetIsRewardsOnL2(&_ContractAttestationCenter.TransactOpts, _isRewardsOnL2)
}

// SetIsRewardsOnL2 is a paid mutator transaction binding the contract method 0xdba8323b.
//
// Solidity: function setIsRewardsOnL2(bool _isRewardsOnL2) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) SetIsRewardsOnL2(_isRewardsOnL2 bool) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SetIsRewardsOnL2(&_ContractAttestationCenter.TransactOpts, _isRewardsOnL2)
}

// SetOblsSharesSyncer is a paid mutator transaction binding the contract method 0x1164224e.
//
// Solidity: function setOblsSharesSyncer(address _oblsSharesSyncer) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) SetOblsSharesSyncer(opts *bind.TransactOpts, _oblsSharesSyncer common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "setOblsSharesSyncer", _oblsSharesSyncer)
}

// SetOblsSharesSyncer is a paid mutator transaction binding the contract method 0x1164224e.
//
// Solidity: function setOblsSharesSyncer(address _oblsSharesSyncer) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) SetOblsSharesSyncer(_oblsSharesSyncer common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SetOblsSharesSyncer(&_ContractAttestationCenter.TransactOpts, _oblsSharesSyncer)
}

// SetOblsSharesSyncer is a paid mutator transaction binding the contract method 0x1164224e.
//
// Solidity: function setOblsSharesSyncer(address _oblsSharesSyncer) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) SetOblsSharesSyncer(_oblsSharesSyncer common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SetOblsSharesSyncer(&_ContractAttestationCenter.TransactOpts, _oblsSharesSyncer)
}

// SubmitTask is a paid mutator transaction binding the contract method 0x8a7ce8a4.
//
// Solidity: function submitTask((string,bytes,address,uint16) _taskInfo, (bool,uint256[2],uint256[2],uint256[]) _blsTaskSubmissionDetails) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) SubmitTask(opts *bind.TransactOpts, _taskInfo IAttestationCenterTaskInfo, _blsTaskSubmissionDetails IAttestationCenterBlsTaskSubmissionDetails) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "submitTask", _taskInfo, _blsTaskSubmissionDetails)
}

// SubmitTask is a paid mutator transaction binding the contract method 0x8a7ce8a4.
//
// Solidity: function submitTask((string,bytes,address,uint16) _taskInfo, (bool,uint256[2],uint256[2],uint256[]) _blsTaskSubmissionDetails) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) SubmitTask(_taskInfo IAttestationCenterTaskInfo, _blsTaskSubmissionDetails IAttestationCenterBlsTaskSubmissionDetails) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SubmitTask(&_ContractAttestationCenter.TransactOpts, _taskInfo, _blsTaskSubmissionDetails)
}

// SubmitTask is a paid mutator transaction binding the contract method 0x8a7ce8a4.
//
// Solidity: function submitTask((string,bytes,address,uint16) _taskInfo, (bool,uint256[2],uint256[2],uint256[]) _blsTaskSubmissionDetails) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) SubmitTask(_taskInfo IAttestationCenterTaskInfo, _blsTaskSubmissionDetails IAttestationCenterBlsTaskSubmissionDetails) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SubmitTask(&_ContractAttestationCenter.TransactOpts, _taskInfo, _blsTaskSubmissionDetails)
}

// SubmitTask0 is a paid mutator transaction binding the contract method 0xfff768e3.
//
// Solidity: function submitTask((string,bytes,address,uint16) _taskInfo, (bool,bytes,uint256[2],uint256[]) _ecdsaTaskSubmissionDetails) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) SubmitTask0(opts *bind.TransactOpts, _taskInfo IAttestationCenterTaskInfo, _ecdsaTaskSubmissionDetails IAttestationCenterEcdsaTaskSubmissionDetails) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "submitTask0", _taskInfo, _ecdsaTaskSubmissionDetails)
}

// SubmitTask0 is a paid mutator transaction binding the contract method 0xfff768e3.
//
// Solidity: function submitTask((string,bytes,address,uint16) _taskInfo, (bool,bytes,uint256[2],uint256[]) _ecdsaTaskSubmissionDetails) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) SubmitTask0(_taskInfo IAttestationCenterTaskInfo, _ecdsaTaskSubmissionDetails IAttestationCenterEcdsaTaskSubmissionDetails) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SubmitTask0(&_ContractAttestationCenter.TransactOpts, _taskInfo, _ecdsaTaskSubmissionDetails)
}

// SubmitTask0 is a paid mutator transaction binding the contract method 0xfff768e3.
//
// Solidity: function submitTask((string,bytes,address,uint16) _taskInfo, (bool,bytes,uint256[2],uint256[]) _ecdsaTaskSubmissionDetails) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) SubmitTask0(_taskInfo IAttestationCenterTaskInfo, _ecdsaTaskSubmissionDetails IAttestationCenterEcdsaTaskSubmissionDetails) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SubmitTask0(&_ContractAttestationCenter.TransactOpts, _taskInfo, _ecdsaTaskSubmissionDetails)
}

// TransferMessageHandler is a paid mutator transaction binding the contract method 0x4d07f651.
//
// Solidity: function transferMessageHandler(address _newMessageHandler) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) TransferMessageHandler(opts *bind.TransactOpts, _newMessageHandler common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "transferMessageHandler", _newMessageHandler)
}

// TransferMessageHandler is a paid mutator transaction binding the contract method 0x4d07f651.
//
// Solidity: function transferMessageHandler(address _newMessageHandler) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) TransferMessageHandler(_newMessageHandler common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.TransferMessageHandler(&_ContractAttestationCenter.TransactOpts, _newMessageHandler)
}

// TransferMessageHandler is a paid mutator transaction binding the contract method 0x4d07f651.
//
// Solidity: function transferMessageHandler(address _newMessageHandler) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) TransferMessageHandler(_newMessageHandler common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.TransferMessageHandler(&_ContractAttestationCenter.TransactOpts, _newMessageHandler)
}

// UnRegisterOperatorFromNetwork is a paid mutator transaction binding the contract method 0x27bbb287.
//
// Solidity: function unRegisterOperatorFromNetwork(address _operator) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) UnRegisterOperatorFromNetwork(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "unRegisterOperatorFromNetwork", _operator)
}

// UnRegisterOperatorFromNetwork is a paid mutator transaction binding the contract method 0x27bbb287.
//
// Solidity: function unRegisterOperatorFromNetwork(address _operator) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) UnRegisterOperatorFromNetwork(_operator common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.UnRegisterOperatorFromNetwork(&_ContractAttestationCenter.TransactOpts, _operator)
}

// UnRegisterOperatorFromNetwork is a paid mutator transaction binding the contract method 0x27bbb287.
//
// Solidity: function unRegisterOperatorFromNetwork(address _operator) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) UnRegisterOperatorFromNetwork(_operator common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.UnRegisterOperatorFromNetwork(&_ContractAttestationCenter.TransactOpts, _operator)
}

// Unpause is a paid mutator transaction binding the contract method 0xbac1e94b.
//
// Solidity: function unpause(bytes4 _pausableFlow) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) Unpause(opts *bind.TransactOpts, _pausableFlow [4]byte) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "unpause", _pausableFlow)
}

// Unpause is a paid mutator transaction binding the contract method 0xbac1e94b.
//
// Solidity: function unpause(bytes4 _pausableFlow) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) Unpause(_pausableFlow [4]byte) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.Unpause(&_ContractAttestationCenter.TransactOpts, _pausableFlow)
}

// Unpause is a paid mutator transaction binding the contract method 0xbac1e94b.
//
// Solidity: function unpause(bytes4 _pausableFlow) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) Unpause(_pausableFlow [4]byte) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.Unpause(&_ContractAttestationCenter.TransactOpts, _pausableFlow)
}

// UpdateBlsKey is a paid mutator transaction binding the contract method 0x6ba5aa46.
//
// Solidity: function updateBlsKey(uint256[4] _blsKey, (uint256[2]) _authSignature) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) UpdateBlsKey(opts *bind.TransactOpts, _blsKey [4]*big.Int, _authSignature BLSAuthLibrarySignature) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "updateBlsKey", _blsKey, _authSignature)
}

// UpdateBlsKey is a paid mutator transaction binding the contract method 0x6ba5aa46.
//
// Solidity: function updateBlsKey(uint256[4] _blsKey, (uint256[2]) _authSignature) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) UpdateBlsKey(_blsKey [4]*big.Int, _authSignature BLSAuthLibrarySignature) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.UpdateBlsKey(&_ContractAttestationCenter.TransactOpts, _blsKey, _authSignature)
}

// UpdateBlsKey is a paid mutator transaction binding the contract method 0x6ba5aa46.
//
// Solidity: function updateBlsKey(uint256[4] _blsKey, (uint256[2]) _authSignature) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) UpdateBlsKey(_blsKey [4]*big.Int, _authSignature BLSAuthLibrarySignature) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.UpdateBlsKey(&_ContractAttestationCenter.TransactOpts, _blsKey, _authSignature)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.Fallback(&_ContractAttestationCenter.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.Fallback(&_ContractAttestationCenter.TransactOpts, calldata)
}

// ContractAttestationCenterClearPaymentRejectedIterator is returned from FilterClearPaymentRejected and is used to iterate over the raw logs and unpacked data for ClearPaymentRejected events raised by the ContractAttestationCenter contract.
type ContractAttestationCenterClearPaymentRejectedIterator struct {
	Event *ContractAttestationCenterClearPaymentRejected // Event containing the contract specifics and raw log

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
func (it *ContractAttestationCenterClearPaymentRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAttestationCenterClearPaymentRejected)
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
		it.Event = new(ContractAttestationCenterClearPaymentRejected)
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
func (it *ContractAttestationCenterClearPaymentRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAttestationCenterClearPaymentRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAttestationCenterClearPaymentRejected represents a ClearPaymentRejected event raised by the ContractAttestationCenter contract.
type ContractAttestationCenterClearPaymentRejected struct {
	Operator               common.Address
	RequestedTaskNumber    *big.Int
	RequestedAmountClaimed *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterClearPaymentRejected is a free log retrieval operation binding the contract event 0x1e643658b8248efd3563f24d116430bf571d036bea3721d94e848890a00a1023.
//
// Solidity: event ClearPaymentRejected(address indexed operator, uint256 requestedTaskNumber, uint256 requestedAmountClaimed)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) FilterClearPaymentRejected(opts *bind.FilterOpts, operator []common.Address) (*ContractAttestationCenterClearPaymentRejectedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAttestationCenter.contract.FilterLogs(opts, "ClearPaymentRejected", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterClearPaymentRejectedIterator{contract: _ContractAttestationCenter.contract, event: "ClearPaymentRejected", logs: logs, sub: sub}, nil
}

// WatchClearPaymentRejected is a free log subscription operation binding the contract event 0x1e643658b8248efd3563f24d116430bf571d036bea3721d94e848890a00a1023.
//
// Solidity: event ClearPaymentRejected(address indexed operator, uint256 requestedTaskNumber, uint256 requestedAmountClaimed)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) WatchClearPaymentRejected(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterClearPaymentRejected, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAttestationCenter.contract.WatchLogs(opts, "ClearPaymentRejected", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAttestationCenterClearPaymentRejected)
				if err := _ContractAttestationCenter.contract.UnpackLog(event, "ClearPaymentRejected", log); err != nil {
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

// ParseClearPaymentRejected is a log parse operation binding the contract event 0x1e643658b8248efd3563f24d116430bf571d036bea3721d94e848890a00a1023.
//
// Solidity: event ClearPaymentRejected(address indexed operator, uint256 requestedTaskNumber, uint256 requestedAmountClaimed)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) ParseClearPaymentRejected(log types.Log) (*ContractAttestationCenterClearPaymentRejected, error) {
	event := new(ContractAttestationCenterClearPaymentRejected)
	if err := _ContractAttestationCenter.contract.UnpackLog(event, "ClearPaymentRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAttestationCenterEigenPaymentsRequestedIterator is returned from FilterEigenPaymentsRequested and is used to iterate over the raw logs and unpacked data for EigenPaymentsRequested events raised by the ContractAttestationCenter contract.
type ContractAttestationCenterEigenPaymentsRequestedIterator struct {
	Event *ContractAttestationCenterEigenPaymentsRequested // Event containing the contract specifics and raw log

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
func (it *ContractAttestationCenterEigenPaymentsRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAttestationCenterEigenPaymentsRequested)
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
		it.Event = new(ContractAttestationCenterEigenPaymentsRequested)
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
func (it *ContractAttestationCenterEigenPaymentsRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAttestationCenterEigenPaymentsRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAttestationCenterEigenPaymentsRequested represents a EigenPaymentsRequested event raised by the ContractAttestationCenter contract.
type ContractAttestationCenterEigenPaymentsRequested struct {
	StartTimestamp     uint32
	Duration           uint32
	Operators          []IAttestationCenterPaymentRequestMessage
	LastPaidTaskNumber *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterEigenPaymentsRequested is a free log retrieval operation binding the contract event 0x928efea964c83ee05474996615e6ed34feaad6ce29a94eeec47d3f69388f5513.
//
// Solidity: event EigenPaymentsRequested(uint32 startTimestamp, uint32 duration, (address,uint256)[] operators, uint256 lastPaidTaskNumber)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) FilterEigenPaymentsRequested(opts *bind.FilterOpts) (*ContractAttestationCenterEigenPaymentsRequestedIterator, error) {

	logs, sub, err := _ContractAttestationCenter.contract.FilterLogs(opts, "EigenPaymentsRequested")
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterEigenPaymentsRequestedIterator{contract: _ContractAttestationCenter.contract, event: "EigenPaymentsRequested", logs: logs, sub: sub}, nil
}

// WatchEigenPaymentsRequested is a free log subscription operation binding the contract event 0x928efea964c83ee05474996615e6ed34feaad6ce29a94eeec47d3f69388f5513.
//
// Solidity: event EigenPaymentsRequested(uint32 startTimestamp, uint32 duration, (address,uint256)[] operators, uint256 lastPaidTaskNumber)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) WatchEigenPaymentsRequested(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterEigenPaymentsRequested) (event.Subscription, error) {

	logs, sub, err := _ContractAttestationCenter.contract.WatchLogs(opts, "EigenPaymentsRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAttestationCenterEigenPaymentsRequested)
				if err := _ContractAttestationCenter.contract.UnpackLog(event, "EigenPaymentsRequested", log); err != nil {
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

// ParseEigenPaymentsRequested is a log parse operation binding the contract event 0x928efea964c83ee05474996615e6ed34feaad6ce29a94eeec47d3f69388f5513.
//
// Solidity: event EigenPaymentsRequested(uint32 startTimestamp, uint32 duration, (address,uint256)[] operators, uint256 lastPaidTaskNumber)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) ParseEigenPaymentsRequested(log types.Log) (*ContractAttestationCenterEigenPaymentsRequested, error) {
	event := new(ContractAttestationCenterEigenPaymentsRequested)
	if err := _ContractAttestationCenter.contract.UnpackLog(event, "EigenPaymentsRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAttestationCenterFlowPausedIterator is returned from FilterFlowPaused and is used to iterate over the raw logs and unpacked data for FlowPaused events raised by the ContractAttestationCenter contract.
type ContractAttestationCenterFlowPausedIterator struct {
	Event *ContractAttestationCenterFlowPaused // Event containing the contract specifics and raw log

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
func (it *ContractAttestationCenterFlowPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAttestationCenterFlowPaused)
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
		it.Event = new(ContractAttestationCenterFlowPaused)
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
func (it *ContractAttestationCenterFlowPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAttestationCenterFlowPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAttestationCenterFlowPaused represents a FlowPaused event raised by the ContractAttestationCenter contract.
type ContractAttestationCenterFlowPaused struct {
	PausableFlow [4]byte
	Pauser       common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFlowPaused is a free log retrieval operation binding the contract event 0x95c3658c5e0c74e20cf12db371b9b67d26e97a1937f6d2284f88cc44d036b4f6.
//
// Solidity: event FlowPaused(bytes4 _pausableFlow, address _pauser)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) FilterFlowPaused(opts *bind.FilterOpts) (*ContractAttestationCenterFlowPausedIterator, error) {

	logs, sub, err := _ContractAttestationCenter.contract.FilterLogs(opts, "FlowPaused")
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterFlowPausedIterator{contract: _ContractAttestationCenter.contract, event: "FlowPaused", logs: logs, sub: sub}, nil
}

// WatchFlowPaused is a free log subscription operation binding the contract event 0x95c3658c5e0c74e20cf12db371b9b67d26e97a1937f6d2284f88cc44d036b4f6.
//
// Solidity: event FlowPaused(bytes4 _pausableFlow, address _pauser)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) WatchFlowPaused(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterFlowPaused) (event.Subscription, error) {

	logs, sub, err := _ContractAttestationCenter.contract.WatchLogs(opts, "FlowPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAttestationCenterFlowPaused)
				if err := _ContractAttestationCenter.contract.UnpackLog(event, "FlowPaused", log); err != nil {
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
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) ParseFlowPaused(log types.Log) (*ContractAttestationCenterFlowPaused, error) {
	event := new(ContractAttestationCenterFlowPaused)
	if err := _ContractAttestationCenter.contract.UnpackLog(event, "FlowPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAttestationCenterFlowUnpausedIterator is returned from FilterFlowUnpaused and is used to iterate over the raw logs and unpacked data for FlowUnpaused events raised by the ContractAttestationCenter contract.
type ContractAttestationCenterFlowUnpausedIterator struct {
	Event *ContractAttestationCenterFlowUnpaused // Event containing the contract specifics and raw log

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
func (it *ContractAttestationCenterFlowUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAttestationCenterFlowUnpaused)
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
		it.Event = new(ContractAttestationCenterFlowUnpaused)
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
func (it *ContractAttestationCenterFlowUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAttestationCenterFlowUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAttestationCenterFlowUnpaused represents a FlowUnpaused event raised by the ContractAttestationCenter contract.
type ContractAttestationCenterFlowUnpaused struct {
	PausableFlowFlag [4]byte
	Unpauser         common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterFlowUnpaused is a free log retrieval operation binding the contract event 0xc7e56e17b0a6c4b467df6495e1eda1baecd7ba20604e80c1058ac06f4578d85e.
//
// Solidity: event FlowUnpaused(bytes4 _pausableFlowFlag, address _unpauser)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) FilterFlowUnpaused(opts *bind.FilterOpts) (*ContractAttestationCenterFlowUnpausedIterator, error) {

	logs, sub, err := _ContractAttestationCenter.contract.FilterLogs(opts, "FlowUnpaused")
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterFlowUnpausedIterator{contract: _ContractAttestationCenter.contract, event: "FlowUnpaused", logs: logs, sub: sub}, nil
}

// WatchFlowUnpaused is a free log subscription operation binding the contract event 0xc7e56e17b0a6c4b467df6495e1eda1baecd7ba20604e80c1058ac06f4578d85e.
//
// Solidity: event FlowUnpaused(bytes4 _pausableFlowFlag, address _unpauser)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) WatchFlowUnpaused(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterFlowUnpaused) (event.Subscription, error) {

	logs, sub, err := _ContractAttestationCenter.contract.WatchLogs(opts, "FlowUnpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAttestationCenterFlowUnpaused)
				if err := _ContractAttestationCenter.contract.UnpackLog(event, "FlowUnpaused", log); err != nil {
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
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) ParseFlowUnpaused(log types.Log) (*ContractAttestationCenterFlowUnpaused, error) {
	event := new(ContractAttestationCenterFlowUnpaused)
	if err := _ContractAttestationCenter.contract.UnpackLog(event, "FlowUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAttestationCenterInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractAttestationCenter contract.
type ContractAttestationCenterInitializedIterator struct {
	Event *ContractAttestationCenterInitialized // Event containing the contract specifics and raw log

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
func (it *ContractAttestationCenterInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAttestationCenterInitialized)
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
		it.Event = new(ContractAttestationCenterInitialized)
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
func (it *ContractAttestationCenterInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAttestationCenterInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAttestationCenterInitialized represents a Initialized event raised by the ContractAttestationCenter contract.
type ContractAttestationCenterInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractAttestationCenterInitializedIterator, error) {

	logs, sub, err := _ContractAttestationCenter.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterInitializedIterator{contract: _ContractAttestationCenter.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractAttestationCenter.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAttestationCenterInitialized)
				if err := _ContractAttestationCenter.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) ParseInitialized(log types.Log) (*ContractAttestationCenterInitialized, error) {
	event := new(ContractAttestationCenterInitialized)
	if err := _ContractAttestationCenter.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAttestationCenterOperatorBlsKeyUpdatedIterator is returned from FilterOperatorBlsKeyUpdated and is used to iterate over the raw logs and unpacked data for OperatorBlsKeyUpdated events raised by the ContractAttestationCenter contract.
type ContractAttestationCenterOperatorBlsKeyUpdatedIterator struct {
	Event *ContractAttestationCenterOperatorBlsKeyUpdated // Event containing the contract specifics and raw log

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
func (it *ContractAttestationCenterOperatorBlsKeyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAttestationCenterOperatorBlsKeyUpdated)
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
		it.Event = new(ContractAttestationCenterOperatorBlsKeyUpdated)
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
func (it *ContractAttestationCenterOperatorBlsKeyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAttestationCenterOperatorBlsKeyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAttestationCenterOperatorBlsKeyUpdated represents a OperatorBlsKeyUpdated event raised by the ContractAttestationCenter contract.
type ContractAttestationCenterOperatorBlsKeyUpdated struct {
	Operator common.Address
	BlsKey   [4]*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorBlsKeyUpdated is a free log retrieval operation binding the contract event 0x764bc14e663abcee4585e1a92e552918c69d453c673e7161504ff52fc3d428c9.
//
// Solidity: event OperatorBlsKeyUpdated(address indexed operator, uint256[4] blsKey)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) FilterOperatorBlsKeyUpdated(opts *bind.FilterOpts, operator []common.Address) (*ContractAttestationCenterOperatorBlsKeyUpdatedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAttestationCenter.contract.FilterLogs(opts, "OperatorBlsKeyUpdated", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterOperatorBlsKeyUpdatedIterator{contract: _ContractAttestationCenter.contract, event: "OperatorBlsKeyUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorBlsKeyUpdated is a free log subscription operation binding the contract event 0x764bc14e663abcee4585e1a92e552918c69d453c673e7161504ff52fc3d428c9.
//
// Solidity: event OperatorBlsKeyUpdated(address indexed operator, uint256[4] blsKey)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) WatchOperatorBlsKeyUpdated(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterOperatorBlsKeyUpdated, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAttestationCenter.contract.WatchLogs(opts, "OperatorBlsKeyUpdated", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAttestationCenterOperatorBlsKeyUpdated)
				if err := _ContractAttestationCenter.contract.UnpackLog(event, "OperatorBlsKeyUpdated", log); err != nil {
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

// ParseOperatorBlsKeyUpdated is a log parse operation binding the contract event 0x764bc14e663abcee4585e1a92e552918c69d453c673e7161504ff52fc3d428c9.
//
// Solidity: event OperatorBlsKeyUpdated(address indexed operator, uint256[4] blsKey)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) ParseOperatorBlsKeyUpdated(log types.Log) (*ContractAttestationCenterOperatorBlsKeyUpdated, error) {
	event := new(ContractAttestationCenterOperatorBlsKeyUpdated)
	if err := _ContractAttestationCenter.contract.UnpackLog(event, "OperatorBlsKeyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAttestationCenterOperatorEjectionRequestedIterator is returned from FilterOperatorEjectionRequested and is used to iterate over the raw logs and unpacked data for OperatorEjectionRequested events raised by the ContractAttestationCenter contract.
type ContractAttestationCenterOperatorEjectionRequestedIterator struct {
	Event *ContractAttestationCenterOperatorEjectionRequested // Event containing the contract specifics and raw log

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
func (it *ContractAttestationCenterOperatorEjectionRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAttestationCenterOperatorEjectionRequested)
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
		it.Event = new(ContractAttestationCenterOperatorEjectionRequested)
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
func (it *ContractAttestationCenterOperatorEjectionRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAttestationCenterOperatorEjectionRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAttestationCenterOperatorEjectionRequested represents a OperatorEjectionRequested event raised by the ContractAttestationCenter contract.
type ContractAttestationCenterOperatorEjectionRequested struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorEjectionRequested is a free log retrieval operation binding the contract event 0x4bbb3ae76ccd0feb230e3802a79c234a035ea750a918f60db5ae976d67c096d0.
//
// Solidity: event OperatorEjectionRequested(address operator)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) FilterOperatorEjectionRequested(opts *bind.FilterOpts) (*ContractAttestationCenterOperatorEjectionRequestedIterator, error) {

	logs, sub, err := _ContractAttestationCenter.contract.FilterLogs(opts, "OperatorEjectionRequested")
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterOperatorEjectionRequestedIterator{contract: _ContractAttestationCenter.contract, event: "OperatorEjectionRequested", logs: logs, sub: sub}, nil
}

// WatchOperatorEjectionRequested is a free log subscription operation binding the contract event 0x4bbb3ae76ccd0feb230e3802a79c234a035ea750a918f60db5ae976d67c096d0.
//
// Solidity: event OperatorEjectionRequested(address operator)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) WatchOperatorEjectionRequested(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterOperatorEjectionRequested) (event.Subscription, error) {

	logs, sub, err := _ContractAttestationCenter.contract.WatchLogs(opts, "OperatorEjectionRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAttestationCenterOperatorEjectionRequested)
				if err := _ContractAttestationCenter.contract.UnpackLog(event, "OperatorEjectionRequested", log); err != nil {
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

// ParseOperatorEjectionRequested is a log parse operation binding the contract event 0x4bbb3ae76ccd0feb230e3802a79c234a035ea750a918f60db5ae976d67c096d0.
//
// Solidity: event OperatorEjectionRequested(address operator)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) ParseOperatorEjectionRequested(log types.Log) (*ContractAttestationCenterOperatorEjectionRequested, error) {
	event := new(ContractAttestationCenterOperatorEjectionRequested)
	if err := _ContractAttestationCenter.contract.UnpackLog(event, "OperatorEjectionRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAttestationCenterOperatorRegisteredToNetworkIterator is returned from FilterOperatorRegisteredToNetwork and is used to iterate over the raw logs and unpacked data for OperatorRegisteredToNetwork events raised by the ContractAttestationCenter contract.
type ContractAttestationCenterOperatorRegisteredToNetworkIterator struct {
	Event *ContractAttestationCenterOperatorRegisteredToNetwork // Event containing the contract specifics and raw log

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
func (it *ContractAttestationCenterOperatorRegisteredToNetworkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAttestationCenterOperatorRegisteredToNetwork)
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
		it.Event = new(ContractAttestationCenterOperatorRegisteredToNetwork)
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
func (it *ContractAttestationCenterOperatorRegisteredToNetworkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAttestationCenterOperatorRegisteredToNetworkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAttestationCenterOperatorRegisteredToNetwork represents a OperatorRegisteredToNetwork event raised by the ContractAttestationCenter contract.
type ContractAttestationCenterOperatorRegisteredToNetwork struct {
	Operator    common.Address
	VotingPower *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegisteredToNetwork is a free log retrieval operation binding the contract event 0x16c1a2a8195923d655fe84191b37c746f4385a5c32c038578958b29f52daa1a8.
//
// Solidity: event OperatorRegisteredToNetwork(address indexed operator, uint256 votingPower)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) FilterOperatorRegisteredToNetwork(opts *bind.FilterOpts, operator []common.Address) (*ContractAttestationCenterOperatorRegisteredToNetworkIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAttestationCenter.contract.FilterLogs(opts, "OperatorRegisteredToNetwork", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterOperatorRegisteredToNetworkIterator{contract: _ContractAttestationCenter.contract, event: "OperatorRegisteredToNetwork", logs: logs, sub: sub}, nil
}

// WatchOperatorRegisteredToNetwork is a free log subscription operation binding the contract event 0x16c1a2a8195923d655fe84191b37c746f4385a5c32c038578958b29f52daa1a8.
//
// Solidity: event OperatorRegisteredToNetwork(address indexed operator, uint256 votingPower)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) WatchOperatorRegisteredToNetwork(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterOperatorRegisteredToNetwork, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAttestationCenter.contract.WatchLogs(opts, "OperatorRegisteredToNetwork", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAttestationCenterOperatorRegisteredToNetwork)
				if err := _ContractAttestationCenter.contract.UnpackLog(event, "OperatorRegisteredToNetwork", log); err != nil {
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

// ParseOperatorRegisteredToNetwork is a log parse operation binding the contract event 0x16c1a2a8195923d655fe84191b37c746f4385a5c32c038578958b29f52daa1a8.
//
// Solidity: event OperatorRegisteredToNetwork(address indexed operator, uint256 votingPower)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) ParseOperatorRegisteredToNetwork(log types.Log) (*ContractAttestationCenterOperatorRegisteredToNetwork, error) {
	event := new(ContractAttestationCenterOperatorRegisteredToNetwork)
	if err := _ContractAttestationCenter.contract.UnpackLog(event, "OperatorRegisteredToNetwork", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAttestationCenterOperatorUnregisteredFromNetworkIterator is returned from FilterOperatorUnregisteredFromNetwork and is used to iterate over the raw logs and unpacked data for OperatorUnregisteredFromNetwork events raised by the ContractAttestationCenter contract.
type ContractAttestationCenterOperatorUnregisteredFromNetworkIterator struct {
	Event *ContractAttestationCenterOperatorUnregisteredFromNetwork // Event containing the contract specifics and raw log

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
func (it *ContractAttestationCenterOperatorUnregisteredFromNetworkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAttestationCenterOperatorUnregisteredFromNetwork)
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
		it.Event = new(ContractAttestationCenterOperatorUnregisteredFromNetwork)
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
func (it *ContractAttestationCenterOperatorUnregisteredFromNetworkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAttestationCenterOperatorUnregisteredFromNetworkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAttestationCenterOperatorUnregisteredFromNetwork represents a OperatorUnregisteredFromNetwork event raised by the ContractAttestationCenter contract.
type ContractAttestationCenterOperatorUnregisteredFromNetwork struct {
	OperatorId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorUnregisteredFromNetwork is a free log retrieval operation binding the contract event 0xda04f7db725bc5a9ad418baf26d08e9f24561a7cc119bfe1dd26bfebfc175db3.
//
// Solidity: event OperatorUnregisteredFromNetwork(uint256 indexed operatorId)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) FilterOperatorUnregisteredFromNetwork(opts *bind.FilterOpts, operatorId []*big.Int) (*ContractAttestationCenterOperatorUnregisteredFromNetworkIterator, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractAttestationCenter.contract.FilterLogs(opts, "OperatorUnregisteredFromNetwork", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterOperatorUnregisteredFromNetworkIterator{contract: _ContractAttestationCenter.contract, event: "OperatorUnregisteredFromNetwork", logs: logs, sub: sub}, nil
}

// WatchOperatorUnregisteredFromNetwork is a free log subscription operation binding the contract event 0xda04f7db725bc5a9ad418baf26d08e9f24561a7cc119bfe1dd26bfebfc175db3.
//
// Solidity: event OperatorUnregisteredFromNetwork(uint256 indexed operatorId)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) WatchOperatorUnregisteredFromNetwork(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterOperatorUnregisteredFromNetwork, operatorId []*big.Int) (event.Subscription, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractAttestationCenter.contract.WatchLogs(opts, "OperatorUnregisteredFromNetwork", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAttestationCenterOperatorUnregisteredFromNetwork)
				if err := _ContractAttestationCenter.contract.UnpackLog(event, "OperatorUnregisteredFromNetwork", log); err != nil {
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

// ParseOperatorUnregisteredFromNetwork is a log parse operation binding the contract event 0xda04f7db725bc5a9ad418baf26d08e9f24561a7cc119bfe1dd26bfebfc175db3.
//
// Solidity: event OperatorUnregisteredFromNetwork(uint256 indexed operatorId)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) ParseOperatorUnregisteredFromNetwork(log types.Log) (*ContractAttestationCenterOperatorUnregisteredFromNetwork, error) {
	event := new(ContractAttestationCenterOperatorUnregisteredFromNetwork)
	if err := _ContractAttestationCenter.contract.UnpackLog(event, "OperatorUnregisteredFromNetwork", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAttestationCenterPaymentsRequestedIterator is returned from FilterPaymentsRequested and is used to iterate over the raw logs and unpacked data for PaymentsRequested events raised by the ContractAttestationCenter contract.
type ContractAttestationCenterPaymentsRequestedIterator struct {
	Event *ContractAttestationCenterPaymentsRequested // Event containing the contract specifics and raw log

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
func (it *ContractAttestationCenterPaymentsRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAttestationCenterPaymentsRequested)
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
		it.Event = new(ContractAttestationCenterPaymentsRequested)
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
func (it *ContractAttestationCenterPaymentsRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAttestationCenterPaymentsRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAttestationCenterPaymentsRequested represents a PaymentsRequested event raised by the ContractAttestationCenter contract.
type ContractAttestationCenterPaymentsRequested struct {
	Operators          []IAttestationCenterPaymentRequestMessage
	LastPaidTaskNumber *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterPaymentsRequested is a free log retrieval operation binding the contract event 0x3e17ccbb628e667c33839a666b60f15eaefb9db2cbae6f7cc9f3f223cd77fece.
//
// Solidity: event PaymentsRequested((address,uint256)[] operators, uint256 lastPaidTaskNumber)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) FilterPaymentsRequested(opts *bind.FilterOpts) (*ContractAttestationCenterPaymentsRequestedIterator, error) {

	logs, sub, err := _ContractAttestationCenter.contract.FilterLogs(opts, "PaymentsRequested")
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterPaymentsRequestedIterator{contract: _ContractAttestationCenter.contract, event: "PaymentsRequested", logs: logs, sub: sub}, nil
}

// WatchPaymentsRequested is a free log subscription operation binding the contract event 0x3e17ccbb628e667c33839a666b60f15eaefb9db2cbae6f7cc9f3f223cd77fece.
//
// Solidity: event PaymentsRequested((address,uint256)[] operators, uint256 lastPaidTaskNumber)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) WatchPaymentsRequested(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterPaymentsRequested) (event.Subscription, error) {

	logs, sub, err := _ContractAttestationCenter.contract.WatchLogs(opts, "PaymentsRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAttestationCenterPaymentsRequested)
				if err := _ContractAttestationCenter.contract.UnpackLog(event, "PaymentsRequested", log); err != nil {
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

// ParsePaymentsRequested is a log parse operation binding the contract event 0x3e17ccbb628e667c33839a666b60f15eaefb9db2cbae6f7cc9f3f223cd77fece.
//
// Solidity: event PaymentsRequested((address,uint256)[] operators, uint256 lastPaidTaskNumber)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) ParsePaymentsRequested(log types.Log) (*ContractAttestationCenterPaymentsRequested, error) {
	event := new(ContractAttestationCenterPaymentsRequested)
	if err := _ContractAttestationCenter.contract.UnpackLog(event, "PaymentsRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAttestationCenterRewardAccumulatedIterator is returned from FilterRewardAccumulated and is used to iterate over the raw logs and unpacked data for RewardAccumulated events raised by the ContractAttestationCenter contract.
type ContractAttestationCenterRewardAccumulatedIterator struct {
	Event *ContractAttestationCenterRewardAccumulated // Event containing the contract specifics and raw log

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
func (it *ContractAttestationCenterRewardAccumulatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAttestationCenterRewardAccumulated)
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
		it.Event = new(ContractAttestationCenterRewardAccumulated)
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
func (it *ContractAttestationCenterRewardAccumulatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAttestationCenterRewardAccumulatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAttestationCenterRewardAccumulated represents a RewardAccumulated event raised by the ContractAttestationCenter contract.
type ContractAttestationCenterRewardAccumulated struct {
	OperatorId               *big.Int
	BaseRewardFeeForOperator *big.Int
	TaskNumber               uint32
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterRewardAccumulated is a free log retrieval operation binding the contract event 0xd3f16e9d8d3fe0ea8a6e5f923fe57e1ae1af6d890ac6c371e8af6cc177a49b65.
//
// Solidity: event RewardAccumulated(uint256 indexed _operatorId, uint256 _baseRewardFeeForOperator, uint32 indexed _taskNumber)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) FilterRewardAccumulated(opts *bind.FilterOpts, _operatorId []*big.Int, _taskNumber []uint32) (*ContractAttestationCenterRewardAccumulatedIterator, error) {

	var _operatorIdRule []interface{}
	for _, _operatorIdItem := range _operatorId {
		_operatorIdRule = append(_operatorIdRule, _operatorIdItem)
	}

	var _taskNumberRule []interface{}
	for _, _taskNumberItem := range _taskNumber {
		_taskNumberRule = append(_taskNumberRule, _taskNumberItem)
	}

	logs, sub, err := _ContractAttestationCenter.contract.FilterLogs(opts, "RewardAccumulated", _operatorIdRule, _taskNumberRule)
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterRewardAccumulatedIterator{contract: _ContractAttestationCenter.contract, event: "RewardAccumulated", logs: logs, sub: sub}, nil
}

// WatchRewardAccumulated is a free log subscription operation binding the contract event 0xd3f16e9d8d3fe0ea8a6e5f923fe57e1ae1af6d890ac6c371e8af6cc177a49b65.
//
// Solidity: event RewardAccumulated(uint256 indexed _operatorId, uint256 _baseRewardFeeForOperator, uint32 indexed _taskNumber)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) WatchRewardAccumulated(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterRewardAccumulated, _operatorId []*big.Int, _taskNumber []uint32) (event.Subscription, error) {

	var _operatorIdRule []interface{}
	for _, _operatorIdItem := range _operatorId {
		_operatorIdRule = append(_operatorIdRule, _operatorIdItem)
	}

	var _taskNumberRule []interface{}
	for _, _taskNumberItem := range _taskNumber {
		_taskNumberRule = append(_taskNumberRule, _taskNumberItem)
	}

	logs, sub, err := _ContractAttestationCenter.contract.WatchLogs(opts, "RewardAccumulated", _operatorIdRule, _taskNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAttestationCenterRewardAccumulated)
				if err := _ContractAttestationCenter.contract.UnpackLog(event, "RewardAccumulated", log); err != nil {
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

// ParseRewardAccumulated is a log parse operation binding the contract event 0xd3f16e9d8d3fe0ea8a6e5f923fe57e1ae1af6d890ac6c371e8af6cc177a49b65.
//
// Solidity: event RewardAccumulated(uint256 indexed _operatorId, uint256 _baseRewardFeeForOperator, uint32 indexed _taskNumber)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) ParseRewardAccumulated(log types.Log) (*ContractAttestationCenterRewardAccumulated, error) {
	event := new(ContractAttestationCenterRewardAccumulated)
	if err := _ContractAttestationCenter.contract.UnpackLog(event, "RewardAccumulated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAttestationCenterRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ContractAttestationCenter contract.
type ContractAttestationCenterRoleAdminChangedIterator struct {
	Event *ContractAttestationCenterRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ContractAttestationCenterRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAttestationCenterRoleAdminChanged)
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
		it.Event = new(ContractAttestationCenterRoleAdminChanged)
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
func (it *ContractAttestationCenterRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAttestationCenterRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAttestationCenterRoleAdminChanged represents a RoleAdminChanged event raised by the ContractAttestationCenter contract.
type ContractAttestationCenterRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ContractAttestationCenterRoleAdminChangedIterator, error) {

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

	logs, sub, err := _ContractAttestationCenter.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterRoleAdminChangedIterator{contract: _ContractAttestationCenter.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _ContractAttestationCenter.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAttestationCenterRoleAdminChanged)
				if err := _ContractAttestationCenter.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) ParseRoleAdminChanged(log types.Log) (*ContractAttestationCenterRoleAdminChanged, error) {
	event := new(ContractAttestationCenterRoleAdminChanged)
	if err := _ContractAttestationCenter.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAttestationCenterRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ContractAttestationCenter contract.
type ContractAttestationCenterRoleGrantedIterator struct {
	Event *ContractAttestationCenterRoleGranted // Event containing the contract specifics and raw log

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
func (it *ContractAttestationCenterRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAttestationCenterRoleGranted)
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
		it.Event = new(ContractAttestationCenterRoleGranted)
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
func (it *ContractAttestationCenterRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAttestationCenterRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAttestationCenterRoleGranted represents a RoleGranted event raised by the ContractAttestationCenter contract.
type ContractAttestationCenterRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ContractAttestationCenterRoleGrantedIterator, error) {

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

	logs, sub, err := _ContractAttestationCenter.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterRoleGrantedIterator{contract: _ContractAttestationCenter.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ContractAttestationCenter.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAttestationCenterRoleGranted)
				if err := _ContractAttestationCenter.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) ParseRoleGranted(log types.Log) (*ContractAttestationCenterRoleGranted, error) {
	event := new(ContractAttestationCenterRoleGranted)
	if err := _ContractAttestationCenter.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAttestationCenterRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ContractAttestationCenter contract.
type ContractAttestationCenterRoleRevokedIterator struct {
	Event *ContractAttestationCenterRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ContractAttestationCenterRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAttestationCenterRoleRevoked)
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
		it.Event = new(ContractAttestationCenterRoleRevoked)
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
func (it *ContractAttestationCenterRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAttestationCenterRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAttestationCenterRoleRevoked represents a RoleRevoked event raised by the ContractAttestationCenter contract.
type ContractAttestationCenterRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ContractAttestationCenterRoleRevokedIterator, error) {

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

	logs, sub, err := _ContractAttestationCenter.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterRoleRevokedIterator{contract: _ContractAttestationCenter.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ContractAttestationCenter.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAttestationCenterRoleRevoked)
				if err := _ContractAttestationCenter.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) ParseRoleRevoked(log types.Log) (*ContractAttestationCenterRoleRevoked, error) {
	event := new(ContractAttestationCenterRoleRevoked)
	if err := _ContractAttestationCenter.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAttestationCenterSetMessageHandlerIterator is returned from FilterSetMessageHandler and is used to iterate over the raw logs and unpacked data for SetMessageHandler events raised by the ContractAttestationCenter contract.
type ContractAttestationCenterSetMessageHandlerIterator struct {
	Event *ContractAttestationCenterSetMessageHandler // Event containing the contract specifics and raw log

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
func (it *ContractAttestationCenterSetMessageHandlerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAttestationCenterSetMessageHandler)
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
		it.Event = new(ContractAttestationCenterSetMessageHandler)
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
func (it *ContractAttestationCenterSetMessageHandlerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAttestationCenterSetMessageHandlerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAttestationCenterSetMessageHandler represents a SetMessageHandler event raised by the ContractAttestationCenter contract.
type ContractAttestationCenterSetMessageHandler struct {
	NewMessageHandler common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSetMessageHandler is a free log retrieval operation binding the contract event 0x997f84b541d7b68e210e6f50e3402b51d8411dbbc4d44ed81e508383126e4e94.
//
// Solidity: event SetMessageHandler(address newMessageHandler)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) FilterSetMessageHandler(opts *bind.FilterOpts) (*ContractAttestationCenterSetMessageHandlerIterator, error) {

	logs, sub, err := _ContractAttestationCenter.contract.FilterLogs(opts, "SetMessageHandler")
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterSetMessageHandlerIterator{contract: _ContractAttestationCenter.contract, event: "SetMessageHandler", logs: logs, sub: sub}, nil
}

// WatchSetMessageHandler is a free log subscription operation binding the contract event 0x997f84b541d7b68e210e6f50e3402b51d8411dbbc4d44ed81e508383126e4e94.
//
// Solidity: event SetMessageHandler(address newMessageHandler)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) WatchSetMessageHandler(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterSetMessageHandler) (event.Subscription, error) {

	logs, sub, err := _ContractAttestationCenter.contract.WatchLogs(opts, "SetMessageHandler")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAttestationCenterSetMessageHandler)
				if err := _ContractAttestationCenter.contract.UnpackLog(event, "SetMessageHandler", log); err != nil {
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
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) ParseSetMessageHandler(log types.Log) (*ContractAttestationCenterSetMessageHandler, error) {
	event := new(ContractAttestationCenterSetMessageHandler)
	if err := _ContractAttestationCenter.contract.UnpackLog(event, "SetMessageHandler", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAttestationCenterTaskRejectedIterator is returned from FilterTaskRejected and is used to iterate over the raw logs and unpacked data for TaskRejected events raised by the ContractAttestationCenter contract.
type ContractAttestationCenterTaskRejectedIterator struct {
	Event *ContractAttestationCenterTaskRejected // Event containing the contract specifics and raw log

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
func (it *ContractAttestationCenterTaskRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAttestationCenterTaskRejected)
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
		it.Event = new(ContractAttestationCenterTaskRejected)
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
func (it *ContractAttestationCenterTaskRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAttestationCenterTaskRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAttestationCenterTaskRejected represents a TaskRejected event raised by the ContractAttestationCenter contract.
type ContractAttestationCenterTaskRejected struct {
	Operator         common.Address
	TaskNumber       uint32
	ProofOfTask      string
	Data             []byte
	TaskDefinitionId uint16
	AttestersIds     []*big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTaskRejected is a free log retrieval operation binding the contract event 0x834bcdc44f628888090e897856fa3d661504c545654ba804c61b293c36b6595b.
//
// Solidity: event TaskRejected(address indexed operator, uint32 taskNumber, string proofOfTask, bytes data, uint16 indexed taskDefinitionId, uint256[] attestersIds)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) FilterTaskRejected(opts *bind.FilterOpts, operator []common.Address, taskDefinitionId []uint16) (*ContractAttestationCenterTaskRejectedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	var taskDefinitionIdRule []interface{}
	for _, taskDefinitionIdItem := range taskDefinitionId {
		taskDefinitionIdRule = append(taskDefinitionIdRule, taskDefinitionIdItem)
	}

	logs, sub, err := _ContractAttestationCenter.contract.FilterLogs(opts, "TaskRejected", operatorRule, taskDefinitionIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterTaskRejectedIterator{contract: _ContractAttestationCenter.contract, event: "TaskRejected", logs: logs, sub: sub}, nil
}

// WatchTaskRejected is a free log subscription operation binding the contract event 0x834bcdc44f628888090e897856fa3d661504c545654ba804c61b293c36b6595b.
//
// Solidity: event TaskRejected(address indexed operator, uint32 taskNumber, string proofOfTask, bytes data, uint16 indexed taskDefinitionId, uint256[] attestersIds)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) WatchTaskRejected(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterTaskRejected, operator []common.Address, taskDefinitionId []uint16) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	var taskDefinitionIdRule []interface{}
	for _, taskDefinitionIdItem := range taskDefinitionId {
		taskDefinitionIdRule = append(taskDefinitionIdRule, taskDefinitionIdItem)
	}

	logs, sub, err := _ContractAttestationCenter.contract.WatchLogs(opts, "TaskRejected", operatorRule, taskDefinitionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAttestationCenterTaskRejected)
				if err := _ContractAttestationCenter.contract.UnpackLog(event, "TaskRejected", log); err != nil {
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

// ParseTaskRejected is a log parse operation binding the contract event 0x834bcdc44f628888090e897856fa3d661504c545654ba804c61b293c36b6595b.
//
// Solidity: event TaskRejected(address indexed operator, uint32 taskNumber, string proofOfTask, bytes data, uint16 indexed taskDefinitionId, uint256[] attestersIds)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) ParseTaskRejected(log types.Log) (*ContractAttestationCenterTaskRejected, error) {
	event := new(ContractAttestationCenterTaskRejected)
	if err := _ContractAttestationCenter.contract.UnpackLog(event, "TaskRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAttestationCenterTaskSubmittedIterator is returned from FilterTaskSubmitted and is used to iterate over the raw logs and unpacked data for TaskSubmitted events raised by the ContractAttestationCenter contract.
type ContractAttestationCenterTaskSubmittedIterator struct {
	Event *ContractAttestationCenterTaskSubmitted // Event containing the contract specifics and raw log

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
func (it *ContractAttestationCenterTaskSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAttestationCenterTaskSubmitted)
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
		it.Event = new(ContractAttestationCenterTaskSubmitted)
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
func (it *ContractAttestationCenterTaskSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAttestationCenterTaskSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAttestationCenterTaskSubmitted represents a TaskSubmitted event raised by the ContractAttestationCenter contract.
type ContractAttestationCenterTaskSubmitted struct {
	Operator         common.Address
	TaskNumber       uint32
	ProofOfTask      string
	Data             []byte
	TaskDefinitionId uint16
	AttestersIds     []*big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTaskSubmitted is a free log retrieval operation binding the contract event 0xaed5b9322f1cce79090471045d025d7a0a841daf787ec955d17cdd32a3ebfdb1.
//
// Solidity: event TaskSubmitted(address indexed operator, uint32 taskNumber, string proofOfTask, bytes data, uint16 indexed taskDefinitionId, uint256[] attestersIds)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) FilterTaskSubmitted(opts *bind.FilterOpts, operator []common.Address, taskDefinitionId []uint16) (*ContractAttestationCenterTaskSubmittedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	var taskDefinitionIdRule []interface{}
	for _, taskDefinitionIdItem := range taskDefinitionId {
		taskDefinitionIdRule = append(taskDefinitionIdRule, taskDefinitionIdItem)
	}

	logs, sub, err := _ContractAttestationCenter.contract.FilterLogs(opts, "TaskSubmitted", operatorRule, taskDefinitionIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterTaskSubmittedIterator{contract: _ContractAttestationCenter.contract, event: "TaskSubmitted", logs: logs, sub: sub}, nil
}

// WatchTaskSubmitted is a free log subscription operation binding the contract event 0xaed5b9322f1cce79090471045d025d7a0a841daf787ec955d17cdd32a3ebfdb1.
//
// Solidity: event TaskSubmitted(address indexed operator, uint32 taskNumber, string proofOfTask, bytes data, uint16 indexed taskDefinitionId, uint256[] attestersIds)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) WatchTaskSubmitted(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterTaskSubmitted, operator []common.Address, taskDefinitionId []uint16) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	var taskDefinitionIdRule []interface{}
	for _, taskDefinitionIdItem := range taskDefinitionId {
		taskDefinitionIdRule = append(taskDefinitionIdRule, taskDefinitionIdItem)
	}

	logs, sub, err := _ContractAttestationCenter.contract.WatchLogs(opts, "TaskSubmitted", operatorRule, taskDefinitionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAttestationCenterTaskSubmitted)
				if err := _ContractAttestationCenter.contract.UnpackLog(event, "TaskSubmitted", log); err != nil {
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

// ParseTaskSubmitted is a log parse operation binding the contract event 0xaed5b9322f1cce79090471045d025d7a0a841daf787ec955d17cdd32a3ebfdb1.
//
// Solidity: event TaskSubmitted(address indexed operator, uint32 taskNumber, string proofOfTask, bytes data, uint16 indexed taskDefinitionId, uint256[] attestersIds)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) ParseTaskSubmitted(log types.Log) (*ContractAttestationCenterTaskSubmitted, error) {
	event := new(ContractAttestationCenterTaskSubmitted)
	if err := _ContractAttestationCenter.contract.UnpackLog(event, "TaskSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
