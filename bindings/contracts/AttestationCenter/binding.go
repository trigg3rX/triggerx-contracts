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

// IAttestationCenterOperatorDetails is an auto generated low-level Go binding around an user-defined struct.
type IAttestationCenterOperatorDetails struct {
	Operator    common.Address
	OperatorId  *big.Int
	VotingPower *big.Int
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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_extensionImplementation\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EigenRewardsDurationExceedsMaximum\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EigenRewardsDurationNotMultipleOfInterval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"totalRewards\",\"type\":\"uint256\"}],\"name\":\"EigenRewardsMaxRewardsAmountExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EigenRewardsMustBeRetroactive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EigenRewardsNotSupportedOnL2\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EigenRewardsStartTimestampNotMultipleOfInterval\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EigenRewardsStartTimestampTooFarInPast\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FlowIsCurrentlyPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FlowIsCurrentlyUnpaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InactiveAggregator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InactiveTaskPerformer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"taskDefinitionId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"minVotingPower\",\"type\":\"uint256\"}],\"name\":\"InsufficientVotingPowerForTaskDefinition\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAttesterSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaximumNumberOfAttesters\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOperatorsForPayment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPerformerSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeForBatchPaymentRequest\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"taskDefinitionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"operatorIndex\",\"type\":\"uint256\"}],\"name\":\"InvalidRestrictedAttester\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTaskDefinition\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageAlreadySigned\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operatorAddress\",\"type\":\"address\"}],\"name\":\"OperatorNotRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PauseFlowIsAlreadyPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"taskDefinitionId\",\"type\":\"uint16\"}],\"name\":\"TaskDefinitionNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnpausingFlowIsAlreadyUnpaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestedTaskNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestedAmountClaimed\",\"type\":\"uint256\"}],\"name\":\"ClearPaymentRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"startTimestamp\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"duration\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeToClaim\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIAttestationCenter.PaymentRequestMessage[]\",\"name\":\"operators\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lastPaidTaskNumber\",\"type\":\"uint256\"}],\"name\":\"EigenPaymentsRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"_pausableFlow\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_pauser\",\"type\":\"address\"}],\"name\":\"FlowPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"_pausableFlowFlag\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_unpauser\",\"type\":\"address\"}],\"name\":\"FlowUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[4]\",\"name\":\"blsKey\",\"type\":\"uint256[4]\"}],\"name\":\"OperatorBlsKeyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"votingPower\",\"type\":\"uint256\"}],\"name\":\"OperatorRegisteredToNetwork\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"operatorId\",\"type\":\"uint256\"}],\"name\":\"OperatorUnregisteredFromNetwork\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeToClaim\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIAttestationCenter.PaymentRequestMessage[]\",\"name\":\"operators\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lastPaidTaskNumber\",\"type\":\"uint256\"}],\"name\":\"PaymentsRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_baseRewardFeeForOperator\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"_taskNumber\",\"type\":\"uint32\"}],\"name\":\"RewardAccumulated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newMessageHandler\",\"type\":\"address\"}],\"name\":\"SetMessageHandler\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"taskNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"proofOfTask\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"taskDefinitionId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"attestersIds\",\"type\":\"uint256[]\"}],\"name\":\"TaskRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"taskNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"proofOfTask\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"taskDefinitionId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"attestersIds\",\"type\":\"uint256[]\"}],\"name\":\"TaskSubmitted\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXTENSION_IMPLEMENTATION\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"avsLogic\",\"outputs\":[{\"internalType\":\"contractIAvsLogic\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"avsTreasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beforePaymentsLogic\",\"outputs\":[{\"internalType\":\"contractIBeforePaymentsLogic\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeToClaim\",\"type\":\"uint256\"}],\"internalType\":\"structIAttestationCenter.PaymentRequestMessage[]\",\"name\":\"_operators\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"_paidTaskNumber\",\"type\":\"uint256\"}],\"name\":\"clearBatchPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveOperatorsDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votingPower\",\"type\":\"uint256\"}],\"internalType\":\"structIAttestationCenter.OperatorDetails[]\",\"name\":\"_operators\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"getOperatorPaymentDetail\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lastPaidTaskNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeToClaim\",\"type\":\"uint256\"},{\"internalType\":\"enumIAttestationCenter.PaymentStatus\",\"name\":\"paymentStatus\",\"type\":\"uint8\"}],\"internalType\":\"structIAttestationCenter.PaymentDetails\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_taskDefinitionId\",\"type\":\"uint16\"}],\"name\":\"getTaskDefinitionMaximumNumberOfAttesters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_taskDefinitionId\",\"type\":\"uint16\"}],\"name\":\"getTaskDefinitionMinimumVotingPower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_taskDefinitionId\",\"type\":\"uint16\"}],\"name\":\"getTaskDefinitionRestrictedAttesters\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"avsGovernanceMultisigOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operationsMultisig\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"communityMultisig\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"messageHandler\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"obls\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"avsTreasury\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isRewardsOnL2\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"internalTaskHandler\",\"type\":\"address\"}],\"internalType\":\"structIAttestationCenter.InitializationParams\",\"name\":\"_initializationParams\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"internalTaskHandler\",\"outputs\":[{\"internalType\":\"contractIInternalTaskHandler\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_pausableFlow\",\"type\":\"bytes4\"}],\"name\":\"isFlowPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_isPaused\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextEigenRewardsBatchStartTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numOfActiveOperators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numOfTaskDefinitions\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numOfTotalOperators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"obls\",\"outputs\":[{\"internalType\":\"contractIOBLS\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"operatorsIdsByAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_pausableFlow\",\"type\":\"bytes4\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_votingPower\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"_blsKey\",\"type\":\"uint256[4]\"},{\"internalType\":\"address\",\"name\":\"_rewardsReceiver\",\"type\":\"address\"}],\"name\":\"registerToNetwork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_to\",\"type\":\"uint256\"}],\"name\":\"requestBatchPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestBatchPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_startTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_duration\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_to\",\"type\":\"uint256\"}],\"name\":\"requestEigenBatchPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oblsSharesSyncer\",\"type\":\"address\"}],\"name\":\"setOblsSharesSyncer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"proofOfTask\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"taskPerformer\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"taskDefinitionId\",\"type\":\"uint16\"}],\"internalType\":\"structIAttestationCenter.TaskInfo\",\"name\":\"_taskInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isApproved\",\"type\":\"bool\"},{\"internalType\":\"uint256[2]\",\"name\":\"tpSignature\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"taSignature\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[]\",\"name\":\"attestersIds\",\"type\":\"uint256[]\"}],\"internalType\":\"structIAttestationCenter.BlsTaskSubmissionDetails\",\"name\":\"_blsTaskSubmissionDetails\",\"type\":\"tuple\"}],\"name\":\"submitTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"proofOfTask\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"taskPerformer\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"taskDefinitionId\",\"type\":\"uint16\"}],\"internalType\":\"structIAttestationCenter.TaskInfo\",\"name\":\"_taskInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isApproved\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"tpSignature\",\"type\":\"bytes\"},{\"internalType\":\"uint256[2]\",\"name\":\"taSignature\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[]\",\"name\":\"attestersIds\",\"type\":\"uint256[]\"}],\"internalType\":\"structIAttestationCenter.EcdsaTaskSubmissionDetails\",\"name\":\"_ecdsaTaskSubmissionDetails\",\"type\":\"tuple\"}],\"name\":\"submitTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taskNumber\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"unRegisterOperatorFromNetwork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_pausableFlow\",\"type\":\"bytes4\"}],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"_blsKey\",\"type\":\"uint256[4]\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"signature\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBLSAuthLibrary.Signature\",\"name\":\"_authSignature\",\"type\":\"tuple\"}],\"name\":\"updateBlsKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_taskDefinitionId\",\"type\":\"uint16\"}],\"name\":\"verifyOperatorValidForTaskDefinition\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"votingPower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600436106102475760003560e01c80638355e33d1161013b578063b0817c44116100b8578063d547741f1161007c578063d547741f14610525578063efd9697814610538578063f28eff3914610572578063fcd4e66a1461057a578063fff768e31461058d57610247565b8063b0817c44146104e7578063b7aa2fdf146104ef578063bac1e94b146104f7578063c07473f61461050a578063c2f429f11461051d57610247565b806397b5f370116100ff57806397b5f370146104775780639878eccb146104975780639eb72d4c146104ac578063a217fddf146104cc578063ac8c16cf146104d457610247565b80638355e33d146104235780638a7ce8a41461042b578063915359fc1461043e57806391d148541461045157806395983fee1461046457610247565b806336568abe116101c95780636ba5aa461161018d5780636ba5aa46146103c55780636f382619146103d857806372d18e8d146103eb57806375d9aedf146104085780637897dec31461041b57610247565b806336568abe146103715780633aa83ec7146103845780635b15c568146103975780635b386be0146103aa578063659fa976146103bd57610247565b8063226def0411610210578063226def04146102f6578063248a9ca31461031d57806327bbb287146103305780632f2ff15d1461034357806334a7c3911461035657610247565b8062028b071461027257806301ffc9a71461028d5780630a1e7093146102b05780631164224e146102c35780631246193e146102d6575b6102707f0000000000000000000000006b1abcd2bdbc069803c3963bf5b20f929681bdde6105a0565b005b61027a6105c9565b6040519081526020015b60405180910390f35b6102a061029b366004614073565b6105dc565b6040519015158152602001610284565b6102706102be3660046140b6565b610613565b6102706102d136600461410f565b6108e2565b6102de61095f565b6040516001600160a01b039091168152602001610284565b6102de7f0000000000000000000000006b1abcd2bdbc069803c3963bf5b20f929681bdde81565b61027a61032b36600461412a565b610980565b61027061033e36600461410f565b6109a2565b610270610351366004614143565b610a7d565b61035e610a9f565b60405161ffff9091168152602001610284565b61027061037f366004614143565b610ab6565b610270610392366004614073565b610ae9565b61027a6103a536600461410f565b610b10565b61027a6103b8366004614181565b610b23565b6102de610b4f565b6102706103d336600461419c565b610b6b565b6102706103e63660046141e2565b610c98565b6103f3610ce5565b60405163ffffffff9091168152602001610284565b61027a610416366004614181565b610cfb565b61027a610d25565b61027a610d38565b61027061043936600461421c565b610d4b565b61027061044c366004614339565b610e89565b6102a061045f366004614143565b610f64565b6102706104723660046143ff565b610f9c565b61048a610485366004614181565b6110a9565b604051610284919061444e565b61049f61112a565b6040516102849190614461565b6104bf6104ba36600461412a565b611321565b60405161028491906144d9565b61027a600081565b6102706104e236600461452f565b6113d2565b6102de61163a565b610270611656565b610270610505366004614073565b6116ad565b61027a61051836600461410f565b6116d4565b6102de611765565b610270610533366004614143565b611781565b6102a0610546366004614073565b6001600160e01b031916600090815260008051602061524f833981519152602052604090205460ff1690565b6102de61179d565b610270610588366004614559565b6117b9565b61027061059b3660046145e3565b6119fb565b3660008037600080366000845af43d6000803e8080156105bf573d6000f35b3d6000fd5b505050565b60006105d3611aea565b60020154905090565b60006001600160e01b03198216637965db0b60e01b148061060d57506301ffc9a760e01b6001600160e01b03198316145b92915050565b7f4b5254ecaf62d67135de29b9998979c2380f8b9012725c83915824ad95f4e88c61063d81611af9565b6378b4401360e11b61064e81611b44565b6000610658611aea565b600e81015490915060ff16156106815760405163ac187e3160e01b815260040160405180910390fd5b61068b8787611b4e565b600080600061069b848989611c5b565b9250925092506f4b3b4ca85a86c47a098a223fffffffff8111156106da5760405163c179495760e01b8152600481018290526024015b60405180910390fd5b81156108a55760018211156108145760005b6106f7600184614658565b8110156108125760005b600161070d8386614658565b6107179190614658565b811015610809578461072a82600161466b565b8151811061073a5761073a61467e565b6020026020010151600001516001600160a01b03168582815181106107615761076161467e565b6020026020010151600001516001600160a01b03161115610801578461078882600161466b565b815181106107985761079861467e565b60200260200101518582815181106107b2576107b261467e565b60200260200101518683815181106107cc576107cc61467e565b60200260200101878460016107e1919061466b565b815181106107f1576107f161467e565b6020026020010182905282905250505b600101610701565b506001016106ec565b505b6040805163ffffffff8c811660208301528b16818301526060810183905260808082018590528251808303909101815260a090910190915282845261085a858583611e46565b84546040517f928efea964c83ee05474996615e6ed34feaad6ce29a94eeec47d3f69388f551391610897918e918e91899163ffffffff16906146d9565b60405180910390a1506108be565b6040516302e66f0160e31b815260040160405180910390fd5b6108c8898b614715565b63ffffffff16846011018190555050505050505050505050565b6378b4401360e11b6108f381611b44565b6108fb611aea565b600301546040516308b2112760e11b81526001600160a01b03848116600483015290911690631164224e90602401600060405180830381600087803b15801561094357600080fd5b505af1158015610957573d6000803e3d6000fd5b505050505050565b6000610969611aea565b600e015461010090046001600160a01b0316919050565b600090815260008051602061526f833981519152602052604090206001015490565b638a70a0eb60e01b6109b381611b44565b60006109bd611aea565b60038101549091506001600160a01b031660006109da8386611ee4565b905082600b01600081546109ed90614732565b909155506040516333ccef3360e21b8152600481018290526001600160a01b0383169063cf33bccc90602401600060405180830381600087803b158015610a3357600080fd5b505af1158015610a47573d6000803e3d6000fd5b50506040518392507fda04f7db725bc5a9ad418baf26d08e9f24561a7cc119bfe1dd26bfebfc175db39150600090a25050505050565b610a8682610980565b610a8f81611b44565b610a998383611fa8565b50505050565b6000610aa9611aea565b6006015461ffff16919050565b6001600160a01b0381163314610adf5760405163334bd91960e11b815260040160405180910390fd5b6105c48282612054565b80610af381611af9565b6001600160e01b03198216610b0781611b44565b6105c4836120d0565b600061060d610b1d611aea565b83611ee4565b600061060d82610b31611aea565b61ffff90911660009081526007909101602052604090206009015490565b6000610b59611aea565b600301546001600160a01b0316919050565b6000610b75611aea565b9050336000610b848383611ee4565b600384015460405163475d551f60e11b81529192506001600160a01b0316908190638ebaaa3e90610bbf908890879030908c90600401614749565b60006040518083038186803b158015610bd757600080fd5b505afa158015610beb573d6000803e3d6000fd5b50506040516386d897a760e01b81526001600160a01b03841692506386d897a79150610c1d9085908a9060040161477a565b600060405180830381600087803b158015610c3757600080fd5b505af1158015610c4b573d6000803e3d6000fd5b50505050826001600160a01b03167f764bc14e663abcee4585e1a92e552918c69d453c673e7161504ff52fc3d428c987604051610c889190614790565b60405180910390a2505050505050565b7fc6d727150e371ebc5958e5c678aae159bc7f1ac1d986edc36fecf6e51285a057610cc281611af9565b6378b4401360e11b610cd381611b44565b610a99610cde611aea565b858561217d565b6000610cef611aea565b5463ffffffff16919050565b600061060d82610d09611aea565b61ffff9091166000908152600791820160205260409020015490565b6000610d2f611aea565b600b0154905090565b6000610d42611aea565b60110154905090565b7f2aedfbf340ba26a0aa3515c5f66276ef1d8ff6c89d5dbc63de01abc44b017be6610d7581611af9565b610d7d61220f565b6040805160a0810190915260009080610d9960208601866147ad565b1515815260200160405180602001604052806000815250815260200184602001600280602002604051908101604052809291908260026020028082843760009201919091525050508152604080518082018252602090920191906060870190600290839083908082843760009201919091525050508152602001610e2060a08601866147ca565b8080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050509152509050610e5f8482612259565b506105c460017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b638a70a0eb60e01b610e9a81611b44565b6000610ea4611aea565b905060005b845181108015610ee9575060006001600160a01b0316858281518110610ed157610ed161467e565b6020026020010151600001516001600160a01b031614155b15610f5d576000858281518110610f0257610f0261467e565b60209081029190910181015180516001600160a01b03166000908152600986018352604080822054808352600a8801855291209282015191935091610f48918890612a22565b50508080610f559061481a565b915050610ea9565b5050505050565b600091825260008051602061526f833981519152602090815260408084206001600160a01b0393909316845291905290205460ff1690565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b0316600081158015610fe15750825b90506000826001600160401b03166001148015610ffd5750303b155b90508115801561100b575080155b156110295760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561105357845460ff60401b1916600160401b1785555b61105c86612ab6565b831561095757845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a1505050505050565b60606110d4826110b7611aea565b61ffff909116600090815260079091016020526040902060080190565b80548060200260200160405190810160405280929190818152602001828054801561111e57602002820191906000526020600020905b81548152602001906001019080831161110a575b50505050509050919050565b60606000611136611aea565b6003810154600b8201549192506001600160a01b0316906001600160401b0381111561116457611164614286565b6040519080825280602002602001820160405280156111c257816020015b6111af604051806060016040528060006001600160a01b0316815260200160008152602001600081525090565b8152602001906001900390816111825790505b509250600060015b8360020154811161131a576040516382afd23b60e01b8152600481018290526001600160a01b038416906382afd23b90602401602060405180830381865afa15801561121a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061123e9190614833565b1561130857604080516060810182526000838152600a8701602090815290839020546001600160a01b03908116835290820184905282516372c4a92760e01b8152600481018590529192830191908616906372c4a92790602401602060405180830381865afa1580156112b5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112d99190614850565b8152508583815181106112ee576112ee61467e565b602002602001018190525081806113049061481a565b9250505b806113128161481a565b9150506111ca565b5050505090565b6113496040805160808101825260008082526020820181905291810182905290606082015290565b611351611aea565b6000838152600a919091016020908152604091829020825160808101845281546001600160a01b0316815260018201549281019290925260028082015493830193909352600381015491929091606084019160ff909116908111156113b8576113b86144c3565b60028111156113c9576113c96144c3565b90525092915050565b60006113dc611aea565b905060006113ea8285611ee4565b905060006113f88385612ccf565b604080516101408101909152815461ffff16815260018201805491929160208401919061142490614869565b80601f016020809104026020016040519081016040528092919081815260200182805461145090614869565b801561149d5780601f106114725761010080835404028352916020019161149d565b820191906000526020600020905b81548152906001019060200180831161148057829003601f168201915b505050505081526020016002820154815260200160038201548152602001600482015481526020016005820154815260200160068201548152602001600782015481526020016008820180548060200260200160405190810160405280929190818152602001828054801561153157602002820191906000526020600020905b81548152602001906001019080831161151d575b505050918352505060099190910154602091820152604080516001808252818301909252929350600092918281019080368337019050509050828160008151811061157e5761157e61467e565b6020026020010181815250506115948282612d18565b60e082015160038501546040516372c4a92760e01b8152600481018690526001600160a01b03909116906372c4a92790602401602060405180830381865afa1580156115e4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116089190614850565b10156109575760e08201516040516349ba8abd60e11b815261ffff8716600482015260248101919091526044016106d1565b6000611644611aea565b600501546001600160a01b0316919050565b7fc6d727150e371ebc5958e5c678aae159bc7f1ac1d986edc36fecf6e51285a05761168081611af9565b6378b4401360e11b61169181611b44565b600061169b611aea565b90506105c4816001836002015461217d565b806116b781612de3565b6001600160e01b031982166116cb81611b44565b6105c483612e2a565b6000806116df611aea565b905060006116ed8285611ee4565b60038301546040516372c4a92760e01b8152600481018390529192506001600160a01b0316906372c4a92790602401602060405180830381865afa158015611739573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061175d9190614850565b949350505050565b600061176f611aea565b600f01546001600160a01b0316919050565b61178a82610980565b61179381611b44565b610a998383612054565b60006117a7611aea565b601001546001600160a01b0316919050565b638a70a0eb60e01b6117ca81611b44565b60006117d4611aea565b9050600081600201600081546117e99061481a565b91829055506001600160a01b038816600090815260098401602052604081205491925090818082156118425750506000818152600a850160205260408120600281018054929055600381015460019091015491935060ff165b85600b01600081546118539061481a565b9190508190555060405180608001604052808c6001600160a01b03168152602001838152602001858152602001826002811115611892576118926144c3565b90526000868152600a88016020908152604091829020835181546001600160a01b0319166001600160a01b039091161781559083015160018083019190915591830151600280830191909155606084015160038301805493949193909260ff1990911691908490811115611908576119086144c3565b021790555050506001600160a01b038b811660009081526009880160209081526040808320899055600d8a019091529081902080546001600160a01b0319168b84161790556003880154905163891a80bb60e01b815291169063891a80bb906119799088908e908e9060040161489d565b600060405180830381600087803b15801561199357600080fd5b505af11580156119a7573d6000803e3d6000fd5b505050508a6001600160a01b03167f16c1a2a8195923d655fe84191b37c746f4385a5c32c038578958b29f52daa1a88b6040516119e691815260200190565b60405180910390a25050505050505050505050565b7f2aedfbf340ba26a0aa3515c5f66276ef1d8ff6c89d5dbc63de01abc44b017be6611a2581611af9565b611a2d61220f565b6040805160a0810190915260009080611a4960208601866147ad565b15158152602001848060200190611a6091906148de565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052509385525050604080518082018252838152602081810194909452928401929092525080518082018252918101919086810190600290839083908082843760009201919091525050508152602001610e2060808601866147ca565b6000611af4612ec7565b905090565b6001600160e01b03198116600090815260008051602061524f833981519152602052604090205460ff1615611b415760405163722fdba960e01b815260040160405180910390fd5b50565b611b418133612ef5565b625c490063ffffffff82161115611b7857604051639464cf7960e01b815260040160405180910390fd5b611b8562093a808261493a565b63ffffffff1615611ba95760405163e9742bc360e01b815260040160405180910390fd5b611bb662093a808361493a565b63ffffffff1615611bda5760405163ba36d17960e01b815260040160405180910390fd5b611be762dd7c004261495d565b63ffffffff168263ffffffff161080611c0957506365fb788063ffffffff8316105b15611c2757604051634bf010dd60e01b815260040160405180910390fd5b42611c328284614715565b63ffffffff161115611c575760405163031b909b60e51b815260040160405180910390fd5b5050565b6060600080841580611c705750856002015484115b80611c7a57508385115b15611c9857604051639ac893ad60e01b815260040160405180910390fd5b855463ffffffff166000611cac8787614658565b611cb790600161466b565b6001600160401b03811115611cce57611cce614286565b604051908082528060200260200182016040528015611d1357816020015b6040805180820190915260008082526020820152815260200190600190039081611cec5790505b509050600080885b888111611e35576000818152600a8c016020526040902080546001600160a01b031615801590611d6357506000600382015460ff166002811115611d6157611d616144c3565b145b8015611d7857508563ffffffff168160010154105b8015611d88575060008160020154115b15611e2c576040805180820190915281546001600160a01b03168152600282015460208201528551869086908110611dc257611dc261467e565b602090810291909101015260038101805460ff1916600117905583611de68161481a565b945050806002015483611df9919061466b565b600e8d015490935060ff1615611e2c57611e2c8c82888f600e0160019054906101000a90046001600160a01b0316612f2e565b50600101611d1b565b509199909850909650945050505050565b6000611e7a83604051602001611e5c919061497a565b60408051601f19818403018152919052855463ffffffff1684613016565b60048086015460405163104c8d4b60e31b81529293506001600160a01b0316916382646a5891611eac918591016149d3565b600060405180830381600087803b158015611ec657600080fd5b505af1158015611eda573d6000803e3d6000fd5b5050505050505050565b60038201546001600160a01b03828116600090815260098501602052604081205490929190911690801580611f7f57506040516382afd23b60e01b8152600481018290526001600160a01b038316906382afd23b90602401602060405180830381865afa158015611f59573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f7d9190614833565b155b1561175d5760405163bd62013360e01b81526001600160a01b03851660048201526024016106d1565b600060008051602061526f833981519152611fc38484610f64565b612043576000848152602082815260408083206001600160a01b03871684529091529020805460ff19166001179055611ff93390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4600191505061060d565b600091505061060d565b5092915050565b600060008051602061526f83398151915261206f8484610f64565b15612043576000848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4600191505061060d565b6001600160e01b03198116600090815260008051602061524f833981519152602081905260409091205460ff161561211b5760405163dfe10d7d60e01b815260040160405180910390fd5b6001600160e01b0319821660009081526020829052604090819020805460ff19166001179055517f95c3658c5e0c74e20cf12db371b9b67d26e97a1937f6d2284f88cc44d036b4f69061217190849033906149e6565b60405180910390a15050565b60008061218b858585611c5b565b5091509150806000036121b1576040516302e66f0160e31b815260040160405180910390fd5b600e85015460ff166121c7576121c78583613089565b84546040517f3e17ccbb628e667c33839a666b60f15eaefb9db2cbae6f7cc9f3f223cd77fece9161220091859163ffffffff1690614a09565b60405180910390a15050505050565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0080546001190161225357604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b600061271161226e6080850160608601614181565b61ffff16101590506000826040015160006002811061228f5761228f61467e565b60200201511515905060006122a2611aea565b60058101549091506001600160a01b0316600081158015906122c2575084155b905080156123c057831561234a578551604080880151606089015160808a015192516379bc050160e01b81526001600160a01b038716946379bc050194612313948e94929390929190600401614b42565b600060405180830381600087803b15801561232d57600080fd5b505af1158015612341573d6000803e3d6000fd5b505050506123c0565b8551602087015160608801516080890151604051630502f5bd60e41b81526001600160a01b0387169463502f5bd09461238d948e94929391929091600401614b95565b600060405180830381600087803b1580156123a757600080fd5b505af11580156123bb573d6000803e3d6000fd5b505050505b60006123cc88806148de565b6123d960208b018b6148de565b6123e960608d0160408e0161410f565b6123f960808e0160608f01614181565b60405160200161240e96959493929190614be3565b60408051601f19818403018152918152815160209283012060008181526008880190935291205490915060ff161561245957604051634510302360e11b815260040160405180910390fd5b60038401546001600160a01b031685156125a5576000816001600160a01b031663a850a9097fb110049506439a07d78731efed3c809a6b13e0bab7cc9805f3bb45fb1e9a67e8856040516020016124b291815260200190565b6040516020818303038152906040526040518363ffffffff1660e01b81526004016124de929190614c32565b6040805180830381865afa1580156124fa573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061251e9190614c4b565b9050816001600160a01b03166353c051bb61254b888d6040016020810190612546919061410f565b611ee4565b838c604001516040518463ffffffff1660e01b815260040161256f93929190614cab565b60006040518083038186803b15801561258757600080fd5b505afa15801561259b573d6000803e3d6000fd5b50505050506125c3565b6125c36125b860608b0160408c0161410f565b838a6020015161311c565b6000816001600160a01b031663a850a9097fb110049506439a07d78731efed3c809a6b13e0bab7cc9805f3bb45fb1e9a67e86126038d8d6000015161315c565b60405160200161261591815260200190565b6040516020818303038152906040526040518363ffffffff1660e01b8152600401612641929190614c32565b6040805180830381865afa15801561265d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126819190614c4b565b90506126dd604051806101400160405280600061ffff1681526020016060815260200160008152602001600081526020016000815260200160008152602001600081526020016000815260200160608152602001600081525090565b6126f6876126f160808e0160608f01614181565b612ccf565b604080516101408101909152815461ffff16815260018201805491929160208401919061272290614869565b80601f016020809104026020016040519081016040528092919081815260200182805461274e90614869565b801561279b5780601f106127705761010080835404028352916020019161279b565b820191906000526020600020905b81548152906001019060200180831161277e57829003601f168201915b505050505081526020016002820154815260200160038201548152602001600482015481526020016005820154815260200160068201548152602001600782015481526020016008820180548060200260200160405190810160405280929190818152602001828054801561282f57602002820191906000526020600020905b81548152602001906001019080831161281b575b505050505081526020016009820154815250509050612852818b60800151612d18565b61287f838b600001518d606001602081019061286e9190614181565b848e606001518f60800151886131ce565b612894878c838d600001518e60800151613250565b60008481526008880160205260409020805460ff19166001179055841561297a57871561293657856001600160a01b031663a971515d8c8c600001518d604001518e606001518f608001516040518663ffffffff1660e01b81526004016128ff959493929190614b42565b600060405180830381600087803b15801561291957600080fd5b505af115801561292d573d6000803e3d6000fd5b505050506129ef565b856001600160a01b031663dd1a53878c8c600001518d602001518e606001518f608001516040518663ffffffff1660e01b81526004016128ff959493929190614b95565b888015612985575089515b156129ef576010870154604051639d00234560e01b81526001600160a01b0390911690639d002345906129bc908e90600401614ccc565b600060405180830381600087803b1580156129d657600080fd5b505af11580156129ea573d6000803e3d6000fd5b505050505b5050505050505050505050565b60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b612a2d838383613693565b15612a6a578015612a5b5781836001018190555080836002016000828254612a559190614658565b90915550505b5050600301805460ff19169055565b825460408051848152602081018490526001600160a01b03909216917f1e643658b8248efd3563f24d116430bf571d036bea3721d94e848890a00a1023910160405180910390a2505050565b612abe6136ee565b6000612acd602083018361410f565b90506000612ae1604084016020850161410f565b90506000612af5606085016040860161410f565b90506000612b09608086016060870161410f565b90506000612b1d60c0870160a0880161410f565b90506001600160a01b038116612b465760405163d92e233d60e01b815260040160405180910390fd5b612b51858585613739565b612b5c858585613776565b612b6461379c565b6000612b6e611aea565b805463ffffffff191660011781559050612b8e60a088016080890161410f565b6003820180546001600160a01b0319166001600160a01b0392909216919091179055612bc0608088016060890161410f565b6004820180546001600160a01b0319166001600160a01b0392909216919091179055612bf260e0880160c089016147ad565b600e820180546001600160a81b031916911515610100600160a81b031916919091176101006001600160a01b038516810291909117909155612c3990880160e0890161410f565b6010820180546001600160a01b0319166001600160a01b0392909216919091179055612c6862093a8042614cdf565b612c729042614658565b6011820155612c88638a70a0eb60e01b84611fa8565b5060008080526007820160205260409020612ca2906137ac565b61271160009081526007820160205260409020612cbe90613801565b612cc6613856565b50505050505050565b61ffff81811660008181526007850160205260408120805491939092911614612d11576040516321321e1960e11b815261ffff841660048201526024016106d1565b9392505050565b43826040015111612d3c57604051635a31d91b60e01b815260040160405180910390fd5b610100820151805115612dab57815181511015612d6c57604051630a8d477960e01b815260040160405180910390fd5b6000612d7883836138b9565b90508015612da9578351604051635966606760e11b815261ffff9091166004820152602481018290526044016106d1565b505b6000836101200151118015612dc557508151836101200151105b156105c457604051630236998b60e01b815260040160405180910390fd5b6001600160e01b03198116600090815260008051602061524f833981519152602052604090205460ff16611b41576040516368c87f3360e11b815260040160405180910390fd5b6001600160e01b03198116600090815260008051602061524f833981519152602081905260409091205460ff16612e7457604051635bfd2da760e11b815260040160405180910390fd5b6001600160e01b0319821660009081526020829052604090819020805460ff19169055517fc7e56e17b0a6c4b467df6495e1eda1baecd7ba20604e80c1058ac06f4578d85e9061217190849033906149e6565b60008061060d60017f47c1bdca9c75057bf1ca178a62fb0ef8908b21ee1ea9b9658ba5135fd3a35469614658565b612eff8282610f64565b611c575760405163e2517d3f60e01b81526001600160a01b0382166004820152602481018390526044016106d1565b82546001600160a01b039081166000908152600d86016020526040812054909116908115612f5c5781612f68565b84546001600160a01b03165b6002860154604051633256b4d160e01b81526001600160a01b03808416600483015263ffffffff88166024830152604482019290925291925060009190851690633256b4d1906064016020604051808303816000875af1158015612fd0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ff49190614833565b905060008115613005575060028601545b611eda878763ffffffff1683612a22565b60607f9d0737d982403370785b2d562a1c75961d85a320fe394f421fc3de24bbbce7b884848460405160240161304e93929190614cf3565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b03199093169290921790915290509392505050565b60006130bc8260405160200161309f919061497a565b60408051601f19818403018152919052845463ffffffff166139a6565b60048085015460405163104c8d4b60e31b81529293506001600160a01b0316916382646a58916130ee918591016149d3565b600060405180830381600087803b15801561310857600080fd5b505af1158015612cc6573d6000803e3d6000fd5b60006131288383613a16565b9050836001600160a01b0316816001600160a01b031614610a9957604051632be1e1cb60e11b815260040160405180910390fd5b600061316883806148de565b61317560208601866148de565b613185606088016040890161410f565b6131956080890160608a01614181565b3046896040516020016131b099989796959493929190614d28565b60405160208183030381529060405280519060200120905092915050565b60006131dc88888888613a40565b60e08601516040516365c4647560e01b81529192506001600160a01b038a16916365c4647591613216918691899189918891600401614d8f565b60006040518083038186803b15801561322e57600080fd5b505afa158015613242573d6000803e3d6000fd5b505050505050505050505050565b845463ffffffff1682156132e45761326e6080860160608701614181565b61ffff16613282606087016040880161410f565b6001600160a01b03167faed5b9322f1cce79090471045d025d7a0a841daf787ec955d17cdd32a3ebfdb1836132b789806148de565b6132c460208c018c6148de565b896040516132d796959493929190614dd0565b60405180910390a3613366565b6132f46080860160608701614181565b61ffff16613308606087016040880161410f565b6001600160a01b03167f834bcdc44f628888090e897856fa3d661504c545654ba804c61b293c36b6595b8361333d89806148de565b61334a60208c018c6148de565b8960405161335d96959493929190614dd0565b60405180910390a35b600061337c876125466060890160408a0161410f565b9050600061338988613b6b565b601089015460038a01546040516382afd23b60e01b815260048101869052929350600160a01b90910460ff16916001600160a01b039091169081906382afd23b90602401602060405180830381865afa1580156133ea573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061340e9190614833565b61342b576040516346d3a4a960e01b815260040160405180910390fd5b8115801561349f57506040516382afd23b60e01b8152600481018490526001600160a01b038216906382afd23b90602401602060405180830381865afa158015613479573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061349d9190614833565b155b156134bd57604051631385e31560e21b815260040160405180910390fd5b5060006040518060a001604052808a6134d590614e9b565b8152602081018590526040810186905260608101889052881515608090910152600c8b01549091506001600160a01b031680158015908061357b575080801561357b5750816001600160a01b031663d005a4f66040518163ffffffff1660e01b8152600401602060405180830381865afa158015613557573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061357b9190614833565b156135cd5760008060006135908f8e88613b8d565b9250925092506135a28f848d8d613c47565b8b156135b4576135b48f838b8d613c80565b866135c5576135c58f828a8d613c80565b505050613652565b6040516336684b6760e01b81526000906001600160a01b038416906336684b67906135fc908790600401614f3f565b6000604051808303816000875af115801561361b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526136439190810190614ff4565b90506136508d828a613ceb565b505b8b5463ffffffff168c6000613666836150a9565b91906101000a81548163ffffffff021916908363ffffffff16021790555050505050505050505050505050565b60006001600385015460ff1660028111156136b0576136b06144c3565b146136bd57506000612d11565b828460010154106136d057506000612d11565b81846002015410156136e457506000612d11565b5060019392505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661373757604051631afcd79f60e31b815260040160405180910390fd5b565b6137416136ee565b61375263d8a8b5c760e01b83611fa8565b506137646378b4401360e11b84611fa8565b50610a9963d3e319af60e01b82611fa8565b61377e6136ee565b613786613d51565b613791838383613d59565b6105c4600033610ab6565b6137a46136ee565b613737613e96565b604080518082019091526007815266191959985d5b1d60ca1b602082015260018201906137d99082615114565b506000196002820155678ac7230489e800006003820181905560048201819055600590910155565b604080518082019091526016815275566f74696e6720506f7765722053796e63205461736b60501b6020820152600182019061383d9082615114565b50805461ffff1916612711178155600019600290910155565b6000613860611aea565b612712600090815260078201602052604081209192505060405180606001604052806027815260200161528f60279139600182019061389f9082615114565b50805461ffff191661271217815560001960029091015550565b600080805b8451821080156138ce5750835181105b15613993578381815181106138e5576138e561467e565b60200260200101518583815181106138ff576138ff61467e565b60200260200101510361392c57816139168161481a565b92505080806139249061481a565b9150506138be565b83818151811061393e5761393e61467e565b60200260200101518583815181106139585761395861467e565b6020026020010151111561397057806139248161481a565b8482815181106139825761398261467e565b60200260200101519250505061060d565b845182036139705760009250505061060d565b60607fd136f2c682d3b91d0bd1dc36dc70e1a7ffbd859fd1ff2240257fdce81d74fd8383836040516024016139dc9291906151d3565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152905092915050565b600080600080613a268686613e9e565b925092509250613a368282613eeb565b5090949350505050565b6000806000836101000151511180613a5c575060008360e00151115b15613ad45760405163255da33160e01b815261ffff851660048201526001600160a01b0387169063255da33190602401602060405180830381865afa158015613aa9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613acd9190614850565b9050613b39565b856001600160a01b031663671b37936040518163ffffffff1660e01b8152600401602060405180830381865afa158015613b12573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613b369190614850565b90505b8415613b5e576003613b4c8260026151f5565b613b56919061520c565b91505061175d565b6003613b4c8260016151f5565b60108101546000908190600160a01b900460ff1661060d57612d118333611ee4565b600c830154600090819081906001600160a01b03168015801590613c265760405163156a159960e21b81526001600160a01b038316906355a8566490613bd7908990600401614f3f565b6060604051808303816000875af1158015613bf6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613c1a9190615220565b91965094509250613c3c565b86606001519450866080015193508660a0015192505b505093509350939050565b60005b8251811015610f5d57613c788585858481518110613c6a57613c6a61467e565b602002602001015185613c80565b600101613c4a565b6000828152600a8501602052604081206002018054859290613ca390849061466b565b909155505060405183815263ffffffff82169083907fd3f16e9d8d3fe0ea8a6e5f923fe57e1ae1af6d890ac6c371e8af6cc177a49b659060200160405180910390a350505050565b60005b8251811015610a99576000838281518110613d0b57613d0b61467e565b60200260200101516000015190506000848381518110613d2d57613d2d61467e565b6020026020010151602001519050613d4786828487613c80565b5050600101613cee565b6137376136ee565b60008051602061524f8339815191526020527f90d0e609be45121efcbf7f1d0ae974a20fcd627e07aa044a1f490249e1d327988054600160ff1991821681179092556312d4953b60e21b6000527fc1d73a3cdf3346e56fde516857b4d0894fb686182db72c29641f07b7f38122b880549091169091179055613de263c6d7271560e01b84611fa8565b50613df46312d4953b60e21b84611fa8565b50613e066362250a9560e11b84611fa8565b50613e186362250a9560e11b82611fa8565b50613e2a6362250a9560e11b83611fa8565b50613e3c634bb5f31f60e11b84611fa8565b50613e4e634bb5f31f60e11b83611fa8565b50613e60634bb5f31f60e11b82611fa8565b50613e72632aedfbf360e01b84611fa8565b50613e84632aedfbf360e01b83611fa8565b50610a99632aedfbf360e01b82611fa8565b6129fc6136ee565b60008060008351604103613ed85760208401516040850151606086015160001a613eca88828585613fa4565b955095509550505050613ee4565b50508151600091506002905b9250925092565b6000826003811115613eff57613eff6144c3565b03613f08575050565b6001826003811115613f1c57613f1c6144c3565b03613f3a5760405163f645eedf60e01b815260040160405180910390fd5b6002826003811115613f4e57613f4e6144c3565b03613f6f5760405163fce698f760e01b8152600481018290526024016106d1565b6003826003811115613f8357613f836144c3565b03611c57576040516335e2f38360e21b8152600481018290526024016106d1565b600080807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841115613fdf5750600091506003905082614069565b604080516000808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015614033573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b03811661405f57506000925060019150829050614069565b9250600091508190505b9450945094915050565b60006020828403121561408557600080fd5b81356001600160e01b031981168114612d1157600080fd5b803563ffffffff811681146140b157600080fd5b919050565b600080600080608085870312156140cc57600080fd5b6140d58561409d565b93506140e36020860161409d565b93969395505050506040820135916060013590565b80356001600160a01b03811681146140b157600080fd5b60006020828403121561412157600080fd5b612d11826140f8565b60006020828403121561413c57600080fd5b5035919050565b6000806040838503121561415657600080fd5b82359150614166602084016140f8565b90509250929050565b803561ffff811681146140b157600080fd5b60006020828403121561419357600080fd5b612d118261416f565b60008082840360c08112156141b057600080fd5b60808401858111156141c157600080fd5b8493506040607f19830112156141d657600080fd5b80925050509250929050565b600080604083850312156141f557600080fd5b50508035926020909101359150565b60006080828403121561421657600080fd5b50919050565b6000806040838503121561422f57600080fd5b82356001600160401b038082111561424657600080fd5b61425286838701614204565b9350602085013591508082111561426857600080fd5b50830160c0818603121561427b57600080fd5b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b03811182821017156142be576142be614286565b60405290565b604051608081016001600160401b03811182821017156142be576142be614286565b604051601f8201601f191681016001600160401b038111828210171561430e5761430e614286565b604052919050565b60006001600160401b0382111561432f5761432f614286565b5060051b60200190565b600080604080848603121561434d57600080fd5b83356001600160401b0381111561436357600080fd5b8401601f8101861361437457600080fd5b8035602061438961438483614316565b6142e6565b82815260069290921b830181019181810190898411156143a857600080fd5b938201935b838510156143ef5785858b0312156143c55760008081fd5b6143cd61429c565b6143d6866140f8565b81528584013584820152825293850193908201906143ad565b9997909101359750505050505050565b6000610100828403121561421657600080fd5b60008151808452602080850194506020840160005b8381101561444357815187529582019590820190600101614427565b509495945050505050565b602081526000612d116020830184614412565b602080825282518282018190526000919060409081850190868401855b828110156144b657815180516001600160a01b031685528681015187860152850151858501526060909301929085019060010161447e565b5091979650505050505050565b634e487b7160e01b600052602160045260246000fd5b81516001600160a01b031681526020808301519082015260408083015190820152606082015160808201906003811061452257634e487b7160e01b600052602160045260246000fd5b8060608401525092915050565b6000806040838503121561454257600080fd5b61454b836140f8565b91506141666020840161416f565b60008060008060e0858703121561456f57600080fd5b614578856140f8565b9350602080860135935086605f87011261459157600080fd5b6145996142c4565b8060c08801898111156145ab57600080fd5b604089015b818110156145c757803584529284019284016145b0565b508195506145d4816140f8565b94505050505092959194509250565b600080604083850312156145f657600080fd5b82356001600160401b038082111561460d57600080fd5b61461986838701614204565b9350602085013591508082111561462f57600080fd5b50830160a0818603121561427b57600080fd5b634e487b7160e01b600052601160045260246000fd5b8181038181111561060d5761060d614642565b8082018082111561060d5761060d614642565b634e487b7160e01b600052603260045260246000fd5b60008151808452602080850194506020840160005b8381101561444357815180516001600160a01b0316885283015183880152604090960195908201906001016146a9565b600063ffffffff80871683528086166020840152608060408401526147016080840186614694565b915080841660608401525095945050505050565b63ffffffff81811683821601908082111561204d5761204d614642565b60008161474157614741614642565b506000190190565b610100810160408683376001600160a01b038581166040840152841660608301526080838184013795945050505050565b82815260a0810160808360208401379392505050565b60808181019083833792915050565b8015158114611b4157600080fd5b6000602082840312156147bf57600080fd5b8135612d118161479f565b6000808335601e198436030181126147e157600080fd5b8301803591506001600160401b038211156147fb57600080fd5b6020019150600581901b360382131561481357600080fd5b9250929050565b60006001820161482c5761482c614642565b5060010190565b60006020828403121561484557600080fd5b8151612d118161479f565b60006020828403121561486257600080fd5b5051919050565b600181811c9082168061487d57607f821691505b60208210810361421657634e487b7160e01b600052602260045260246000fd5b838152602080820184905260c0820190604083018460005b60048110156148d2578151835291830191908301906001016148b5565b50505050949350505050565b6000808335601e198436030181126148f557600080fd5b8301803591506001600160401b0382111561490f57600080fd5b60200191503681900382131561481357600080fd5b634e487b7160e01b600052601260045260246000fd5b600063ffffffff8084168061495157614951614924565b92169190910692915050565b63ffffffff82811682821603908082111561204d5761204d614642565b602081526000612d116020830184614694565b6000815180845260005b818110156149b357602081850181015186830182015201614997565b506000602082860101526020601f19601f83011685010191505092915050565b602081526000612d11602083018461498d565b6001600160e01b03199290921682526001600160a01b0316602082015260400190565b604081526000614a1c6040830185614694565b905063ffffffff831660208301529392505050565b6000808335601e19843603018112614a4857600080fd5b83016020810192503590506001600160401b03811115614a6757600080fd5b80360382131561481357600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6000614aab8283614a31565b60808552614abd608086018284614a76565b915050614acd6020840184614a31565b8583036020870152614ae0838284614a76565b92506001600160a01b039150614afa9050604085016140f8565b16604085015261ffff614b0f6060850161416f565b1660608501528091505092915050565b8060005b6002811015610a99578151845260209384019390910190600101614b23565b60e081526000614b5560e0830188614a9f565b8615156020840152614b6a6040840187614b1f565b614b776080840186614b1f565b82810360c0840152614b898185614412565b98975050505050505050565b60c081526000614ba860c0830188614a9f565b86151560208401528281036040840152614bc2818761498d565b9050614bd16060840186614b1f565b82810360a0840152614b898185614412565b608081526000614bf760808301888a614a76565b8281036020840152614c0a818789614a76565b6001600160a01b03959095166040840152505061ffff91909116606090910152949350505050565b82815260406020820152600061175d604083018461498d565b600060408284031215614c5d57600080fd5b82601f830112614c6c57600080fd5b614c7461429c565b806040840185811115614c8657600080fd5b845b81811015614ca0578051845260209384019301614c88565b509095945050505050565b83815260a08101614cbf6020830185614b1f565b61175d6060830184614b1f565b602081526000612d116020830184614a9f565b600082614cee57614cee614924565b500690565b606081526000614d06606083018661498d565b8460208401528281036040840152614d1e818561498d565b9695505050505050565b60e081526000614d3c60e083018b8d614a76565b8281036020840152614d4f818a8c614a76565b6001600160a01b03988916604085015261ffff979097166060840152505092909416608083015260a082015291151560c090920191909152949350505050565b614d998187614b1f565b614da66040820186614b1f565b60e060808201526000614dbc60e0830186614412565b60a08301949094525060c001529392505050565b63ffffffff87168152608060208201526000614df0608083018789614a76565b8281036040840152614e03818688614a76565b90508281036060840152614e178185614412565b9998505050505050505050565b60006001600160401b03831115614e3d57614e3d614286565b614e50601f8401601f19166020016142e6565b9050828152838383011115614e6457600080fd5b828260208301376000602084830101529392505050565b600082601f830112614e8c57600080fd5b612d1183833560208501614e24565b600060808236031215614ead57600080fd5b614eb56142c4565b82356001600160401b0380821115614ecc57600080fd5b9084019036601f830112614edf57600080fd5b614eee36833560208501614e24565b83526020850135915080821115614f0457600080fd5b50614f1136828601614e7b565b602083015250614f23604084016140f8565b6040820152614f346060840161416f565b606082015292915050565b602081526000825160a060208401528051608060c0850152614f6561014085018261498d565b9050602082015160bf198583030160e0860152614f82828261498d565b6040848101516001600160a01b031661010088015260609485015161ffff166101208801526020880151878201528701518487015292860151858403601f19016080870152929150614fd690508183614412565b9150506080840151614fec60a085018215159052565b509392505050565b6000602080838503121561500757600080fd5b82516001600160401b0381111561501d57600080fd5b8301601f8101851361502e57600080fd5b805161503c61438482614316565b81815260069190911b8201830190838101908783111561505b57600080fd5b928401925b8284101561509e57604084890312156150795760008081fd5b61508161429c565b845181528585015186820152825260409093019290840190615060565b979650505050505050565b600063ffffffff8083168181036150c2576150c2614642565b6001019392505050565b601f8211156105c4576000816000526020600020601f850160051c810160208610156150f55750805b601f850160051c820191505b8181101561095757828155600101615101565b81516001600160401b0381111561512d5761512d614286565b6151418161513b8454614869565b846150cc565b602080601f831160018114615176576000841561515e5750858301515b600019600386901b1c1916600185901b178555610957565b600085815260208120601f198616915b828110156151a557888601518255948401946001909101908401615186565b50858210156151c35787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6040815260006151e6604083018561498d565b90508260208301529392505050565b808202811582820484141761060d5761060d614642565b60008261521b5761521b614924565b500490565b60008060006060848603121561523557600080fd5b835192506020840151915060408401519050925092509256fefe6065fb4e9872e2ad4479001655335380d83f70e163706cd65857449b84510002dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800546f74616c20566f74696e6720506f7765722053796e63205461736b20446566696e6974696f6ea26469706673582212203ed1829016d8736beb288e45466701754f5f7fd772e71ef0fe514334041305cc64736f6c63430008190033",
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

	GetActiveOperatorsDetails(opts *bind.CallOpts) ([]IAttestationCenterOperatorDetails, error)

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

	GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts, _initializationParams IAttestationCenterInitializationParams) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts, _pausableFlow [4]byte) (*types.Transaction, error)

	RegisterToNetwork(opts *bind.TransactOpts, _operator common.Address, _votingPower *big.Int, _blsKey [4]*big.Int, _rewardsReceiver common.Address) (*types.Transaction, error)

	RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error)

	RequestBatchPayment(opts *bind.TransactOpts, _from *big.Int, _to *big.Int) (*types.Transaction, error)

	RequestBatchPayment0(opts *bind.TransactOpts) (*types.Transaction, error)

	RequestEigenBatchPayment(opts *bind.TransactOpts, _startTimestamp uint32, _duration uint32, _from *big.Int, _to *big.Int) (*types.Transaction, error)

	RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	SetOblsSharesSyncer(opts *bind.TransactOpts, _oblsSharesSyncer common.Address) (*types.Transaction, error)

	SubmitTask(opts *bind.TransactOpts, _taskInfo IAttestationCenterTaskInfo, _blsTaskSubmissionDetails IAttestationCenterBlsTaskSubmissionDetails) (*types.Transaction, error)

	SubmitTask0(opts *bind.TransactOpts, _taskInfo IAttestationCenterTaskInfo, _ecdsaTaskSubmissionDetails IAttestationCenterEcdsaTaskSubmissionDetails) (*types.Transaction, error)

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

// GetActiveOperatorsDetails is a free data retrieval call binding the contract method 0x9878eccb.
//
// Solidity: function getActiveOperatorsDetails() view returns((address,uint256,uint256)[] _operators)
func (_ContractAttestationCenter *ContractAttestationCenterCaller) GetActiveOperatorsDetails(opts *bind.CallOpts) ([]IAttestationCenterOperatorDetails, error) {
	var out []interface{}
	err := _ContractAttestationCenter.contract.Call(opts, &out, "getActiveOperatorsDetails")

	if err != nil {
		return *new([]IAttestationCenterOperatorDetails), err
	}

	out0 := *abi.ConvertType(out[0], new([]IAttestationCenterOperatorDetails)).(*[]IAttestationCenterOperatorDetails)

	return out0, err

}

// GetActiveOperatorsDetails is a free data retrieval call binding the contract method 0x9878eccb.
//
// Solidity: function getActiveOperatorsDetails() view returns((address,uint256,uint256)[] _operators)
func (_ContractAttestationCenter *ContractAttestationCenterSession) GetActiveOperatorsDetails() ([]IAttestationCenterOperatorDetails, error) {
	return _ContractAttestationCenter.Contract.GetActiveOperatorsDetails(&_ContractAttestationCenter.CallOpts)
}

// GetActiveOperatorsDetails is a free data retrieval call binding the contract method 0x9878eccb.
//
// Solidity: function getActiveOperatorsDetails() view returns((address,uint256,uint256)[] _operators)
func (_ContractAttestationCenter *ContractAttestationCenterCallerSession) GetActiveOperatorsDetails() ([]IAttestationCenterOperatorDetails, error) {
	return _ContractAttestationCenter.Contract.GetActiveOperatorsDetails(&_ContractAttestationCenter.CallOpts)
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
