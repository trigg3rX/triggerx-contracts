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

// IAttestationCenterTaskSubmissionDetails is an auto generated low-level Go binding around an user-defined struct.
type IAttestationCenterTaskSubmissionDetails struct {
	IsApproved   bool
	TpSignature  []byte
	TaSignature  [2]*big.Int
	AttestersIds []*big.Int
}

// TaskDefinitionParams is an auto generated low-level Go binding around an user-defined struct.
type TaskDefinitionParams struct {
	BlockExpiry                *big.Int
	BaseRewardFeeForAttesters  *big.Int
	BaseRewardFeeForPerformer  *big.Int
	BaseRewardFeeForAggregator *big.Int
	DisputePeriodBlocks        *big.Int
	MinimumVotingPower         *big.Int
	RestrictedOperatorIndexes  []*big.Int
}

// ContractAttestationCenterMetaData contains all meta data concerning the ContractAttestationCenter contract.
var ContractAttestationCenterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"avsLogic\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAvsLogic\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"baseRewardFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"beforePaymentsLogic\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBeforePaymentsLogic\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"clearBatchPayment\",\"inputs\":[{\"name\":\"_operators\",\"type\":\"tuple[]\",\"internalType\":\"structIAttestationCenter.PaymentRequestMessage[]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeToClaim\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"_paidTaskNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"clearPayment\",\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_lastPaidTaskNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_amountClaimed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createNewTaskDefinition\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_taskDefinitionParams\",\"type\":\"tuple\",\"internalType\":\"structTaskDefinitionParams\",\"components\":[{\"name\":\"blockExpiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"baseRewardFeeForAttesters\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"baseRewardFeeForPerformer\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"baseRewardFeeForAggregator\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"disputePeriodBlocks\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minimumVotingPower\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"restrictedOperatorIndexes\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"outputs\":[{\"name\":\"_id\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorPaymentDetail\",\"inputs\":[{\"name\":\"_operatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIAttestationCenter.PaymentDetails\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"lastPaidTaskNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeToClaim\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymentStatus\",\"type\":\"uint8\",\"internalType\":\"enumIAttestationCenter.PaymentStatus\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskDefinitionMinimumVotingPower\",\"inputs\":[{\"name\":\"_taskDefinitionId\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskDefinitionRestrictedOperators\",\"inputs\":[{\"name\":\"_taskDefinitionId\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_avsGovernanceMultisigOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_operationsMultisig\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_communityMultisig\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_messageHandler\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_obls\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isRewardsOnL2\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isFlowPaused\",\"inputs\":[{\"name\":\"_pausableFlow\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"_isPaused\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"numOfActiveOperators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"numOfOperators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"numOfTaskDefinitions\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"numOfTotalOperators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"obls\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOBLS\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorsIdsByAddress\",\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"_pausableFlow\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerToNetwork\",\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_votingPower\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_blsKey\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"},{\"name\":\"_rewardsReceiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestBatchPayment\",\"inputs\":[{\"name\":\"_from\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_to\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestBatchPayment\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestPayment\",\"inputs\":[{\"name\":\"_operatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAvsLogic\",\"inputs\":[{\"name\":\"_avsLogic\",\"type\":\"address\",\"internalType\":\"contractIAvsLogic\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBeforePaymentsLogic\",\"inputs\":[{\"name\":\"_beforePaymentsLogic\",\"type\":\"address\",\"internalType\":\"contractIBeforePaymentsLogic\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeCalculator\",\"inputs\":[{\"name\":\"_feeCalculator\",\"type\":\"address\",\"internalType\":\"contractIFeeCalculator\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOblsSharesSyncer\",\"inputs\":[{\"name\":\"_oblsSharesSyncer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTaskDefinitionMinVotingPower\",\"inputs\":[{\"name\":\"_taskDefinitionId\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"_minimumVotingPower\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTaskDefinitionRestrictedOperators\",\"inputs\":[{\"name\":\"_taskDefinitionId\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"_restrictedOperatorIndexes\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitTask\",\"inputs\":[{\"name\":\"_taskInfo\",\"type\":\"tuple\",\"internalType\":\"structIAttestationCenter.TaskInfo\",\"components\":[{\"name\":\"proofOfTask\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"taskPerformer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taskDefinitionId\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"name\":\"_isApproved\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_tpSignature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_taSignature\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"_attestersIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitTask\",\"inputs\":[{\"name\":\"_taskInfo\",\"type\":\"tuple\",\"internalType\":\"structIAttestationCenter.TaskInfo\",\"components\":[{\"name\":\"proofOfTask\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"taskPerformer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taskDefinitionId\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"name\":\"_taskSubmissionDetails\",\"type\":\"tuple\",\"internalType\":\"structIAttestationCenter.TaskSubmissionDetails\",\"components\":[{\"name\":\"isApproved\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"tpSignature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"taSignature\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"attestersIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferAvsGovernanceMultisig\",\"inputs\":[{\"name\":\"_newAvsGovernanceMultisig\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferMessageHandler\",\"inputs\":[{\"name\":\"_newMessageHandler\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unRegisterOperatorFromNetwork\",\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"_pausableFlow\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateBlsKey\",\"inputs\":[{\"name\":\"_blsKey\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"},{\"name\":\"_authSignature\",\"type\":\"tuple\",\"internalType\":\"structBLSAuthLibrary.Signature\",\"components\":[{\"name\":\"signature\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"vault\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"votingPower\",\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ClearPaymentRejected\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"requestedTaskNumber\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"requestedAmountClaimed\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FlowPaused\",\"inputs\":[{\"name\":\"_pausableFlow\",\"type\":\"bytes4\",\"indexed\":false,\"internalType\":\"bytes4\"},{\"name\":\"_pauser\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FlowUnpaused\",\"inputs\":[{\"name\":\"_pausableFlowFlag\",\"type\":\"bytes4\",\"indexed\":false,\"internalType\":\"bytes4\"},{\"name\":\"_unpauser\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorBlsKeyUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"blsKey\",\"type\":\"uint256[4]\",\"indexed\":false,\"internalType\":\"uint256[4]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRegisteredToNetwork\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"votingPower\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorUnregisteredFromNetwork\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PaymentRequested\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"lastPaidTaskNumber\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"feeToClaim\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PaymentsRequested\",\"inputs\":[{\"name\":\"operators\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structIAttestationCenter.PaymentRequestMessage[]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeToClaim\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"lastPaidTaskNumber\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardAccumulated\",\"inputs\":[{\"name\":\"_operatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"_baseRewardFeeForOperator\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_taskNumber\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetAvsGovernanceMultisig\",\"inputs\":[{\"name\":\"newAvsGovernanceMultisig\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetAvsLogic\",\"inputs\":[{\"name\":\"avsLogic\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetBeforePaymentsLogic\",\"inputs\":[{\"name\":\"paymentsLogic\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetFeeCalculator\",\"inputs\":[{\"name\":\"feeCalculator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetMessageHandler\",\"inputs\":[{\"name\":\"newMessageHandler\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetMinimumTaskDefinitionVotingPower\",\"inputs\":[{\"name\":\"minimumVotingPower\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetRestrictedOperator\",\"inputs\":[{\"name\":\"taskDefinitionId\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"restrictedOperatorIndexes\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskDefinitionCreated\",\"inputs\":[{\"name\":\"taskDefinitionId\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"blockExpiry\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"baseRewardFeeForAttesters\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"baseRewardFeeForPerformer\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"baseRewardFeeForAggregator\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"disputePeriodBlocks\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"minimumVotingPower\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"restrictedOperatorIndexes\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskDefinitionRestrictedOperatorsModified\",\"inputs\":[{\"name\":\"taskDefinitionId\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"restrictedOperatorIndexes\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"isRestricted\",\"type\":\"bool[]\",\"indexed\":false,\"internalType\":\"bool[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskRejected\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"taskNumber\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"proofOfTask\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"taskDefinitionId\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskSubmitted\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"taskNumber\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"proofOfTask\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"taskDefinitionId\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"FlowIsCurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FlowIsCurrentlyUnpaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InactiveAggregator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InactiveTaskPerformer\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidArrayLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAttesterSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBlockExpiry\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBlsKeyUpdateSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorsForPayment\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPaymentClaim\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPerformerSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRangeForBatchPaymentRequest\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRestrictedOperator\",\"inputs\":[{\"name\":\"taskDefinitionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"operatorIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidRestrictedOperatorIndexes\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTaskDefinition\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageAlreadySigned\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotRegistered\",\"inputs\":[{\"name\":\"_operatorAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"PauseFlowIsAlreadyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PaymentClaimed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PaymentReedemed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TaskDefinitionNotFound\",\"inputs\":[{\"name\":\"taskDefinitionId\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"UnpausingFlowIsAlreadyUnpaused\",\"inputs\":[]}]",
	Bin: "0x608080604052346015576152e3908161001a8239f35b5f80fdfe60806040526004361015610011575f80fd5b5f5f3560e01c8062028b0714612c465780628fd38614612bb257806301ffc9a714612b5c5780630c62bf0d146126025780631164224e1461257657806311a95e38146124e0578063242a76a414612493578063248a9ca31461245b57806327bbb287146123385780632f2ff15d146122ed5780633428c126146122af57806334a7c3911461226d57806336568abe146122285780633aa83ec7146121315780634d07f6511461203f5780634e2ce53f14611f98578063513c52ba14611f2b5780635b15c56814611eee5780635de988ab14611bbb57806364ada5d014611aba578063659fa97614611a845780636ade02da146116785780636ba5aa461461193c5780636f3826191461170757806372d18e8d146116d657806375d9aedf1461167d5780637897dec3146116785780637d5f32bc1461155e5780638c66d04f146114c8578063915359fc1461136957806391d14854146113135780639eb72d4c1461124c578063a217fddf14611230578063b0817c44146111e7578063b7aa2fdf14610eb0578063bac1e94b14610dba578063c07473f614610cfa578063c2f429f114610cb1578063c8c9e7ab14610b0b578063d547741f14610ab7578063d9378a59146106ab578063efd969781461065e578063fbfa77cf14610624578063fcd4e66a146103745763fff768e314610207575f80fd5b34610371576040366003190112610371576004356001600160401b03811161036f576080600319823603011261036f576024356001600160401b03811161036b5760a0600319823603011261036b5761025e613188565b610266613abb565b6040519061027382612c8b565b8060040135801515810361033f57825260248101356001600160401b03811161033f576102a69060043691840101612d27565b60208301523660638201121561033b5760408051906102c59082612cc1565b8060848301913683116103535760448401905b838210610343575050604084015235906001600160401b03821161033f5701913660238401121561033b5761031a610327933690602460048201359101612e9e565b6060830152600401613de6565b60015f51602061520e5f395f51905f525580f35b8380fd5b8480fd5b81358152602091820191016102d8565b8680fd5b634e487b7160e01b5f52604160045260245ffd5b8280fd5b505b80fd5b50346103715760e03660031901126103715761038e612d45565b602435366063121561036b576040516103a8608082612cc1565b803660c41161033f576044905b60c4821061061457505060c435906001600160a01b038216820361033f576103db61301c565b6103f25f51602061510e5f395f51905f5254612f2e565b91825f51602061510e5f395f51905f525561041a5f51602061512e5f395f51905f5254612f2e565b5f51602061512e5f395f51905f5255610431612f64565b506040519461043f86612c8b565b6001600160a01b0390811680875260208088018981526040808a018b815260608b018c8152898d525f5160206151ce5f395f51905f52909452908b2099518a546001600160a01b031916951694909417895551600189015591516002880155905190956003908101919081101561060057815460ff191660ff919091161790555f8581527f47c1bdca9c75057bf1ca178a62fb0ef8908b21ee1ea9b9658ba5135fd3a35471602090815260408083208690557f47c1bdca9c75057bf1ca178a62fb0ef8908b21ee1ea9b9658ba5135fd3a35475909152902080546001600160a01b039283166001600160a01b03199091161790555f51602061518e5f395f51905f52541690813b156105fc57908592916040519263891a80bb60e01b845260048401528460248401526044830184905b600482106105e2575050508160c4818580945af180156105d7576105c2575b507f16c1a2a8195923d655fe84191b37c746f4385a5c32c038578958b29f52daa1a86040848482519182526020820152a180f35b816105cc91612cc1565b61036b57825f61058e565b6040513d84823e3d90fd5b82518152889550602092830192600192909201910161056f565b8580fd5b634e487b7160e01b88526021600452602488fd5b81358152602091820191016103b5565b5034610371578060031936011261037157505f51602061522e5f395f51905f525460405160089190911c6001600160a01b03168152602090f35b50346103715760203660031901126103715760043563ffffffff60e01b811680910361036f5760408260ff92602094525f51602061516e5f395f51905f5284522054166040519015158152f35b50346103715760e0366003190112610371576106c5612d45565b6106cd612d5b565b906044356001600160a01b038116810361033b57606435906001600160a01b038216908183036105fc576084356001600160a01b03811692908390036103535760a4356001600160a01b03811696909390878503610ab35760c43593841515809503610aaf575f51602061524e5f395f51905f52549860ff8a60401c1615996001600160401b03811680159081610aa7575b6001149081610a9d575b159081610a94575b50610a855767ffffffffffffffff1981166001175f51602061524e5f395f51905f52558a610a59575b506107a3614d38565b15610a15576108a26108a8926108a26109b59a60408e6107c1614d38565b6107ca856133bc565b506107d483613342565b506107de8761345c565b506107e7614d38565b6107ef614d38565b634958479d60e01b81525f51602061516e5f395f51905f52602052818120600160ff1982541617905563c6d7271560e01b81525f51602061516e5f395f51905f5260205220600160ff19825416179055610848816134fc565b50610852836134fc565b5061085c8161359c565b506108668361359c565b506108708161363c565b5061087a8561363c565b506108848361363c565b5061088e816136dc565b50610898836136dc565b506108a2856136dc565b5061377c565b506108b23361399d565b506108bb614d38565b6108c3614d38565b60015f51602061520e5f395f51905f528190555f51602061510e5f395f51905f528990555f51602061528e5f395f51905f52805463ffffffff19169091179055678ac7230489e800007f47c1bdca9c75057bf1ca178a62fb0ef8908b21ee1ea9b9658ba5135fd3a35469555f51602061518e5f395f51905f5280546001600160a01b03199081166001600160a01b03938416179091555f5160206150ee5f395f51905f528054909116929091169190911790555f51602061522e5f395f51905f5280546001600160a81b03191660ff929092169190911760089290921b610100600160a81b03169190911790556132c3565b506109bd5780f35b68ff0000000000000000195f51602061524e5f395f51905f5254165f51602061524e5f395f51905f52557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a180f35b606460405162461bcd60e51b815260206004820152602060248201527f4174746573746174696f6e43656e7465723a20496e76616c696420696e7075746044820152fd5b68ffffffffffffffffff191668010000000000000001175f51602061524e5f395f51905f52555f61079a565b63f92ee8a960e01b8c5260048cfd5b9050155f610771565b303b159150610769565b8c915061075f565b8980fd5b8880fd5b503461037157604036600319011261037157610b07600435610ad7612d5b565b90610b02610afd825f525f5160206151ae5f395f51905f52602052600160405f20015490565b613061565b613a32565b5080f35b503461037157604036600319011261037157610b25612d85565b6024356001600160401b03811161036b57610b44903690600401612e01565b610b4c612fd7565b610b546130fc565b610b67610b62368385612e9e565b61322a565b15610ca257610b83835f51602061528e5f395f51905f52613bed565b600881016001600160401b038311610c8e579081610ba2848894612f9e565b82526020822084835b858110610c775750505f51602061518e5f395f51905f525460079290920154916001600160a01b03169050803b1561036b5784839185938389610c046040519788968795869463ca87bf8f60e01b865260048601612f0e565b03925af180156105d757610c5e575b5092610c587f364aa2fa0cf7a32b9240f9ab2bdebc99f0222262852b3b25a87388acff5a5b149361ffff93604051948594168452604060208501526040840191612eea565b0390a180f35b81610c6b91959395612cc1565b61033b5791835f610c13565b813581840155889450602090910190600101610bab565b634e487b7160e01b86526041600452602486fd5b63d996449f60e01b8452600484fd5b5034610371578060031936011261037157507f47c1bdca9c75057bf1ca178a62fb0ef8908b21ee1ea9b9658ba5135fd3a35477546040516001600160a01b039091168152602090f35b503461037157602036600319011261037157610d2a610d17612d45565b5f51602061528e5f395f51905f52613281565b90602060018060a01b035f51602061518e5f395f51905f525416926024604051809581936372c4a92760e01b835260048301525afa908115610dae5790610d77575b602090604051908152f35b506020813d602011610da6575b81610d9160209383612cc1565b81010312610da25760209051610d6c565b5f80fd5b3d9150610d84565b604051903d90823e3d90fd5b5034610371576020366003190112610371576004356001600160e01b0319811680820361036b578083525f51602061516e5f395f51905f5260205260ff60408420541615610ea157610e0b81613061565b8083525f51602061516e5f395f51905f5260205260ff60408420541615610e925782525f51602061516e5f395f51905f5260209081526040808420805460ff1916905580516001600160e01b0319909316835233918301919091527fc7e56e17b0a6c4b467df6495e1eda1baecd7ba20604e80c1058ac06f4578d85e919081908101610c58565b635bfd2da760e11b8352600483fd5b6368c87f3360e11b8352600483fd5b5034610371578060031936011261037157610ec9613142565b610ed1612fd7565b5f51602061510e5f395f51905f52545f51602061528e5f395f51905f525f6111dd565b6111ce5763ffffffff81541690825f198101116111ba57610f1483612e31565b92610f226040519485612cc1565b808452601f19610f3182612e31565b01855b8181106111975750505f51602061522e5f395f51905f525460ff811691869160081c6001600160a01b031660015b828111156110b057505050156110a15790849115610fc2575b507f3e17ccbb628e667c33839a666b60f15eaefb9db2cbae6f7cc9f3f223cd77fece610fb68484604051928392604084526040840190613c50565b9060208301520390a180f35b61103161101d9163ffffffff60405191602080840152610ff783610fe9604082018b613c50565b03601f198101855284612cc1565b541660405193849263689b796360e11b6020850152604060248501526064840190613206565b90604483015203601f198101835282612cc1565b5f5160206150ee5f395f51905f52546001600160a01b0316803b1561036b5760405163104c8d4b60e31b815260206004820152918391839182908490829061107d906024830190613206565b03925af180156105d75715610f7b578161109691612cc1565b61036b57825f610f7b565b6302e66f0160e31b8552600485fd5b8089525f5160206151ce5f395f51905f526020528460408a20888a60018060a01b0383541680151580611188575b8061117b575b8061116e575b6110fc575b5050505050600101610f62565b88611135600197969a8995946111499460028901546040519161111e83612ca6565b8252602082015261112f8383612f3c565b52612f3c565b5060038501805460ff191688179055612f2e565b9761115c575b9050879293508a826110ef565b611166928a613b0b565b5f838961114f565b50600284015415156110ea565b50826001850154106110e4565b5061119284614bb6565b6110de565b6020906040516111a681612ca6565b888152888382015282828901015201610f34565b634e487b7160e01b84526011600452602484fd5b639ac893ad60e01b8352600483fd5b5081600111610ef4565b5034610371578060031936011261037157507f47c1bdca9c75057bf1ca178a62fb0ef8908b21ee1ea9b9658ba5135fd3a3546d546040516001600160a01b039091168152602090f35b5034610371578060031936011261037157602090604051908152f35b503461037157602036600319011261037157611266612f64565b5060043581525f5160206151ce5f395f51905f52602052604081206040519061128e82612c8b565b60018060a01b03815416825260018101546020830190815260ff600360028401549360408601948552015416606084019460038210156112ff575084526040805193516001600160a01b03168452905160208401529051908201529051608091906112fd906060830190612e48565bf35b634e487b7160e01b81526021600452602490fd5b503461037157604036600319011261037157604061132f612d5b565b9160043581525f5160206151ae5f395f51905f52602052209060018060a01b03165f52602052602060ff60405f2054166040519015158152f35b5034610371576040366003190112610371576004356001600160401b03811161036f573660238201121561036f5780600401356113a581612e31565b916113b36040519384612cc1565b8183526024602084019260061b8201019036821161033f57602401915b818310611490575050506024356113e561301c565b7f47c1bdca9c75057bf1ca178a62fb0ef8908b21ee1ea9b9658ba5135fd3a3547190835b8351811080611473575b1561146f578061146561142961146a9387612f3c565b5180516001600160a01b03165f9081526020878152604091829020548a525f5160206151ce5f395f51905f528152908920910151908590614b2e565b612f2e565b611409565b8480f35b506001600160a01b036114868286612f3c565b5151161515611413565b60408336031261033f57602060409182516114aa81612ca6565b6114b386612d71565b815282860135838201528152019201916113d0565b5034610371576020366003190112610371576004356001600160a01b0381169081900361036f5760207f83b9ee7f260088fdd4ee12aa07fa7daebc115d796b6bfb55bfb0fc839bccff2d9161151b612fd7565b6115236130a7565b7f47c1bdca9c75057bf1ca178a62fb0ef8908b21ee1ea9b9658ba5135fd3a3547480546001600160a01b03191682179055604051908152a180f35b50346103715760c036600319011261037157600435906001600160401b03821161037157608060031983360301126103715760243580151580910361036f57604435926001600160401b03841161036b573660238501121561036b5783600401356001600160401b03811161033b57366024828701011161033b573660a41161033b5760a4356001600160401b03811161033f5761160361162f913690600401612e01565b96909261160e613188565b611616613abb565b6040519561162387612c8b565b86526024369201612ce2565b60208401526040805195906116449087612cc1565b856064965b60a488106116685750604085015261032793929161031a913691612e9e565b8735815260209788019701611649565b612dd8565b50346103715760203660031901126103715761ffff9061169b612d85565b9050165f527f47c1bdca9c75057bf1ca178a62fb0ef8908b21ee1ea9b9658ba5135fd3a3546f6020526020600760405f200154604051908152f35b503461037157806003193601126103715750602063ffffffff5f51602061528e5f395f51905f525416604051908152f35b503461037157604036600319011261037157602435600435611727613142565b61172f612fd7565b5f51602061528e5f395f51905f5281158015611925575b801561191c575b61190d5763ffffffff815416916117648185613c43565b93600185018095116118f95761177985612e31565b946117876040519687612cc1565b808652611796601f1991612e31565b01865b8181106118d65750505f51602061522e5f395f51905f525460ff81169287929160081c6001600160a01b0316905b8281111561181a57505050156110a15790849115610fc257507f3e17ccbb628e667c33839a666b60f15eaefb9db2cbae6f7cc9f3f223cd77fece610fb68484604051928392604084526040840190613c50565b8089525f5160206151ce5f395f51905f526020528460408a20888a60018060a01b03835416801515806118c7575b806118ba575b806118ad575b611866575b50505050506001016117c7565b88611135600197969a8995946118889460028901546040519161111e83612ca6565b9761189b575b9050879293508a82611859565b6118a5928a613b0b565b5f838961188e565b5060028401541515611854565b508260018501541061184e565b506118d184614bb6565b611848565b6020906040516118e581612ca6565b898152898382015282828a01015201611799565b634e487b7160e01b86526011600452602486fd5b639ac893ad60e01b8452600484fd5b5082821161174d565b505f51602061510e5f395f51905f52548311611746565b50346103715760c03660031901126103715736608411610371576040366083190112610371578061197a335f51602061528e5f395f51905f52613281565b5f51602061518e5f395f51905f52546001600160a01b031690813b15611a60576040805163475d551f60e11b81529060846004830137336044820152306064820152608060046084830137838161010481865afa908115611a79578491611a64575b5050813b15611a6057829160a4839260405194859384926386d897a760e01b845260048401526080600460248501375af180156105d757611a4b575b506040516080600482377f764bc14e663abcee4585e1a92e552918c69d453c673e7161504ff52fc3d428c960803392a280f35b81611a5591612cc1565b61037157805f611a18565b5050fd5b81611a6e91612cc1565b611a6057825f6119dc565b6040513d86823e3d90fd5b5034610371578060031936011261037157505f51602061518e5f395f51905f52546040516001600160a01b039091168152602090f35b503461037157604036600319011261037157611ad4612d85565b8160243591611ae1612fd7565b611ae96130fc565b826007611b03835f51602061528e5f395f51905f52613bed565b01555f51602061518e5f395f51905f52545f51602061510e5f395f51905f52546001600160a01b0390911691823b1561033b5760405163e010f95760e01b815261ffff9190911660048201526024810191909152604481018490529082908290606490829084905af180156105d757611ba6575b507f255c174d5deb340ac0a0c908147d9c66ae7d94a6c7f969f722bf5d71b92e98f8602083604051908152a180f35b81611bb091612cc1565b61036f57815f611b77565b503461037157602036600319011261037157634958479d60e01b81525f51602061516e5f395f51905f5260205260408120546004359060ff16611edf57611c00613abb565b8082525f5160206151ce5f395f51905f52602052604082205f51602061528e5f395f51905f5280547f47c1bdca9c75057bf1ca178a62fb0ef8908b21ee1ea9b9658ba5135fd3a354775463ffffffff9091169385916001600160a01b03169081611e55575b505082546001600160a01b031633039050611e4657611c8382614bb6565b15611e37578260018301541015611e285760028201805415611e195760038301600160ff1982541617905554918460ff5f51602061522e5f395f51905f5254165f14611d45575083917f34682c7a1451dbbbc9e3be4912a8f466eba3a8c72e9bcb5cb3a61e423a9c697394611d129260018060a01b035f51602061522e5f395f51905f525460081c1692613b0b565b6040805133815263ffffffff929092166020830152810191909152606090a160015f51602061520e5f395f51905f525580f35b9163ffffffff9150541660405190632ab2353960e01b6020830152336024830152604482015282606482015260648152611d80608482612cc1565b5f5160206150ee5f395f51905f52546001600160a01b0316803b1561036b5760405163104c8d4b60e31b8152602060048201529183918391829084908290611dcc906024830190613206565b03925af180156105d757611e04575b50507f34682c7a1451dbbbc9e3be4912a8f466eba3a8c72e9bcb5cb3a61e423a9c697391611d12565b81611e0e91612cc1565b61036b57825f611ddb565b639613c55b60e01b8552600485fd5b63c854300f60e01b8452600484fd5b6323d86e2b60e21b8452600484fd5b63b9aa612760e01b8452600484fd5b813b1561036b57829160c483926040519485938492631e34d6a960e21b8452600484015260018060a01b038a5416602484015260018a0154604484015260028a01546064840152611eb160ff60038c0154166084850190612e48565b8a60a48401525af180156105d757611eca575b80611c65565b81611ed491612cc1565b61033b57835f611ec4565b63722fdba960e01b8252600482fd5b503461037157602036600319011261037157611f23602091611f0e612d45565b90505f51602061528e5f395f51905f52613281565b604051908152f35b5034610371576020366003190112610371577f024e98b7d808a3ddb028252dc95dfdcb165a0ca59fcff8984b4fecf9a72226496020611f68612d45565b611f70612fd7565b611f7933613925565b50611f8381613342565b506040516001600160a01b039091168152a180f35b50346103715760203660031901126103715761ffff611fb5612d85565b165f527f47c1bdca9c75057bf1ca178a62fb0ef8908b21ee1ea9b9658ba5135fd3a3546f602052600860405f2001604051918260208354918281520192825260208220915b818110612029576120258561201181870382612cc1565b604051918291602083526020830190612da5565b0390f35b8254845260209093019260019283019201611ffa565b503461037157602036600319011261037157612059612d45565b63d8a8b5c760e01b82525f5160206151ae5f395f51905f5260209081526040808420335f908152925290205460ff1615612113575f5160206150ee5f395f51905f52547f997f84b541d7b68e210e6f50e3402b51d8411dbbc4d44ed81e508383126e4e94916020916120d3906001600160a01b03166138ad565b506120dd816132c3565b505f5160206150ee5f395f51905f5280546001600160a01b0319166001600160a01b03929092169182179055604051908152a180f35b63e2517d3f60e01b82523360045263d8a8b5c760e01b602452604482fd5b5034610371576020366003190112610371576004356001600160e01b0319811680820361036b578083525f51602061516e5f395f51905f5260205260ff6040842054166122195761218181613061565b8083525f51602061516e5f395f51905f5260205260ff60408420541661220a5782525f51602061516e5f395f51905f5260209081526040808420805460ff1916600117905580516001600160e01b0319909316835233918301919091527f95c3658c5e0c74e20cf12db371b9b67d26e97a1937f6d2284f88cc44d036b4f6919081908101610c58565b63dfe10d7d60e01b8352600483fd5b63722fdba960e01b8352600483fd5b503461037157604036600319011261037157612242612d5b565b336001600160a01b0382160361225e57610b0790600435613a32565b63334bd91960e11b8252600482fd5b503461037157806003193601126103715750602061ffff7f47c1bdca9c75057bf1ca178a62fb0ef8908b21ee1ea9b9658ba5135fd3a3546e5416604051908152f35b50346103715780600319360112610371575060207f47c1bdca9c75057bf1ca178a62fb0ef8908b21ee1ea9b9658ba5135fd3a3546954604051908152f35b503461037157604036600319011261037157610b0760043561230d612d5b565b90612333610afd825f525f5160206151ae5f395f51905f52602052600160405f20015490565b61381c565b503461037157602036600319011261037157612352612d45565b61235a61301c565b5f51602061518e5f395f51905f52546001600160a01b0316612389825f51602061528e5f395f51905f52613281565b9160018060a01b03165f527f47c1bdca9c75057bf1ca178a62fb0ef8908b21ee1ea9b9658ba5135fd3a354716020528260405f20555f51602061512e5f395f51905f525480156111ba575f19015f51602061512e5f395f51905f52558290803b1561036f578180916024604051809481936333ccef3360e21b83528860048401525af180156105d757612446575b507fda04f7db725bc5a9ad418baf26d08e9f24561a7cc119bfe1dd26bfebfc175db3602083604051908152a180f35b8161245091612cc1565b61036f57815f612417565b5034610371576020366003190112610371576020611f236004355f525f5160206151ae5f395f51905f52602052600160405f20015490565b5034610371576060366003190112610371576124b86124b0612d45565b610d1761301c565b81525f5160206151ce5f395f51905f526020526124dd60443560243560408420614b2e565b80f35b5034610371576020366003190112610371576004356001600160a01b0381169081900361036f5760207f6da780d66fa2f1ae3eb780c2f39d31bec5c71c81a572640a6af0e8a44347779291612533612fd7565b61253b6130a7565b7f47c1bdca9c75057bf1ca178a62fb0ef8908b21ee1ea9b9658ba5135fd3a3547780546001600160a01b03191682179055604051908152a180f35b50346103715760203660031901126103715780612591612d45565b612599612fd7565b5f51602061518e5f395f51905f52546001600160a01b031690813b15611a60576040516308b2112760e11b81526001600160a01b0390911660048201529082908290602490829084905af180156105d7576125f15750f35b816125fb91612cc1565b6103715780f35b5034610da2576040366003190112610da2576004356001600160401b038111610da257612633903690600401612d27565b6024356001600160401b038111610da2578060040160e06003198336030112610da25761265e612fd7565b6126666130fc565b5f51602061510e5f395f51905f52549080359343851115612b4d5761ffff7f47c1bdca9c75057bf1ca178a62fb0ef8908b21ee1ea9b9658ba5135fd3a3546e541661ffff8114612b3957600101927f47c1bdca9c75057bf1ca178a62fb0ef8908b21ee1ea9b9658ba5135fd3a3546e61ffff851661ffff1982541617905560248501359160448601359660648701359760848801359460c460a48a01359901956127108789612e69565b9b9060405161271e81612c6f565b8c61ffff8d169e8f93848452602084018b8152604085018b815260608601908a825260808701928c845261276960a08901958b875260c08a01978d895260e08b01998a523691612e9e565b986101008901998a525f527f47c1bdca9c75057bf1ca178a62fb0ef8908b21ee1ea9b9658ba5135fd3a3546f60205261ffff8060405f209951161661ffff198954161788556001880190518051906001600160401b038211610357576127cf83546131ce565b601f8111612afe575b50602090601f8311600114612a915760089a9998979695949392915f9183612a86575b50508160011b915f199060031b1c19161790555b516002870155516003860155516004850155516005840155516006830155516007820155019051908151916001600160401b038311610357579060208f979998969594939261285e8484612f9e565b01905f5260205f205f5b838110612a5d5750505050917f4306a9df64b49fc07c0f1929d57cc2b5cdfc108656189460e6aa127754413ef196916128f594936128a68b8d612e69565b9490936128c76040519a8b9a8b5261012060208c01526101208b0190613206565b9660408a01526060890152608088015260a087015260c08601528b60e0860152848303610100860152612eea565b0390a16129028284612e69565b1590506129c45750612921610b6261291a8385612e69565b3691612e9e565b156129b5575f51602061518e5f395f51905f52548693926001600160a01b039091169161294d91612e69565b949092823b1561033f5784906129796040519788968795869463ca87bf8f60e01b865260048601612f0e565b03925af180156129aa57612995575b5060209150604051908152f35b6129a0838092612cc1565b61036f5781612988565b6040513d85823e3d90fd5b63d996449f60e01b8652600486fd5b9550505090806129da575b505060209150611f23565b5f51602061518e5f395f51905f52546001600160a01b0316803b15610da25760405163e010f95760e01b815261ffff93909316600484015260248301949094526044820152915f908390606490829084905af1918215612a5257602092612a42575b806129cf565b5f612a4c91612cc1565b5f612a3c565b6040513d5f823e3d90fd5b829394959697999a9850602060019293519401938184015501908f979998969594939291612868565b015190505f806127fb565b90601f19831691845f52815f20925f5b818110612ae6575091600193918560089e9d9c9b9a999897969410612ace575b505050811b01905561280f565b01515f1960f88460031b161c191690555f8080612ac1565b92936020600181928786015181550195019301612aa1565b612b2990845f5260205f20601f850160051c81019160208610612b2f575b601f0160051c0190612f88565b5f6127d8565b9091508190612b1c565b634e487b7160e01b5f52601160045260245ffd5b63fa63b3a960e01b5f5260045ffd5b34610da2576020366003190112610da25760043563ffffffff60e01b8116809103610da257602090637965db0b60e01b8114908115612ba1575b506040519015158152f35b6301ffc9a760e01b14905082612b96565b34610da2576020366003190112610da2576004356001600160a01b03811690819003610da25760207fdf0d3b0bf99a87fc195d045bc7ec61d3e2619e6a49dd3f5cb69102b8c970203491612c04612fd7565b612c0c6130a7565b7f47c1bdca9c75057bf1ca178a62fb0ef8908b21ee1ea9b9658ba5135fd3a3546d80546001600160a01b03191682179055604051908152a1005b34610da2575f366003190112610da25760205f51602061510e5f395f51905f5254604051908152f35b61012081019081106001600160401b0382111761035757604052565b608081019081106001600160401b0382111761035757604052565b604081019081106001600160401b0382111761035757604052565b90601f801991011681019081106001600160401b0382111761035757604052565b9291926001600160401b0382116103575760405191612d0b601f8201601f191660200184612cc1565b829481845281830111610da2578281602093845f960137010152565b9080601f83011215610da257816020612d4293359101612ce2565b90565b600435906001600160a01b0382168203610da257565b602435906001600160a01b0382168203610da257565b35906001600160a01b0382168203610da257565b6004359061ffff82168203610da257565b359061ffff82168203610da257565b90602080835192838152019201905f5b818110612dc25750505090565b8251845260209384019390920191600101612db5565b34610da2575f366003190112610da25760205f51602061512e5f395f51905f5254604051908152f35b9181601f84011215610da2578235916001600160401b038311610da2576020808501948460051b010111610da257565b6001600160401b0381116103575760051b60200190565b906003821015612e555752565b634e487b7160e01b5f52602160045260245ffd5b903590601e1981360301821215610da257018035906001600160401b038211610da257602001918160051b36038313610da257565b929190612eaa81612e31565b93612eb86040519586612cc1565b602085838152019160051b8101928311610da257905b828210612eda57505050565b8135815260209182019101612ece565b81835290916001600160fb1b038311610da25760209260051b809284830137010190565b612d42949261ffff60609316825260208201528160408201520191612eea565b5f198114612b395760010190565b8051821015612f505760209160051b010190565b634e487b7160e01b5f52603260045260245ffd5b60405190612f7182612c8b565b5f6060838281528260208201528260408201520152565b818110612f93575050565b5f8155600101612f88565b9068010000000000000000811161035757815490808355818110612fc157505050565b612fd5925f5260205f209182019101612f88565b565b335f9081525f51602061526e5f395f51905f52602052604090205460ff1615612ffc57565b63e2517d3f60e01b5f908152336004526378b4401360e11b602452604490fd5b335f9081525f51602061514e5f395f51905f52602052604090205460ff161561304157565b63e2517d3f60e01b5f90815233600452638a70a0eb60e01b602452604490fd5b5f8181525f5160206151ae5f395f51905f526020908152604080832033845290915290205460ff16156130915750565b63e2517d3f60e01b5f523360045260245260445ffd5b6362250a9560e11b5f525f51602061516e5f395f51905f526020527fa2bb07c4b8f3ee9c9761d047d22dc0dd0458ab83840e48f1bb05e9bd1da0ab295460ff166130ed57565b63722fdba960e01b5f5260045ffd5b634bb5f31f60e11b5f525f51602061516e5f395f51905f526020527fe769c088cb976c2f234c8b85baa7dbd97d56bcc14abf32b4bdf747a265b085345460ff166130ed57565b63c6d7271560e01b5f525f51602061516e5f395f51905f526020527f90d0e609be45121efcbf7f1d0ae974a20fcd627e07aa044a1f490249e1d327985460ff166130ed57565b632aedfbf360e01b5f525f51602061516e5f395f51905f526020527f6f639d33107d73d2076a485b42dcf1a5a4ddaafd8ef16403bbea8dc49a45e18a5460ff166130ed57565b90600182811c921680156131fc575b60208310146131e857565b634e487b7160e01b5f52602260045260245ffd5b91607f16916131dd565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b805190600182111561327a575f915f1901915b82811061324c57505050600190565b6132568183612f3c565b516132646001830184612f3c565b5111156132735760010161323d565b5050505f90565b5050600190565b60099060018060a01b0383165f520160205260405f20549081156132a3575090565b63bd62013360e01b5f9081526001600160a01b0391909116600452602490fd5b6001600160a01b0381165f9081525f51602061514e5f395f51905f52602052604090205460ff1661333d576001600160a01b03165f8181525f51602061514e5f395f51905f5260205260408120805460ff19166001179055339190638a70a0eb60e01b905f5160206150ce5f395f51905f529080a4600190565b505f90565b6001600160a01b0381165f9081525f51602061526e5f395f51905f52602052604090205460ff1661333d576001600160a01b03165f8181525f51602061526e5f395f51905f5260205260408120805460ff191660011790553391906378b4401360e11b905f5160206150ce5f395f51905f529080a4600190565b6001600160a01b0381165f9081527f274b5753bc2a873526e44bae648b363f47953a4f0c2234820809428fe7b7dafd602052604090205460ff1661333d576001600160a01b03165f8181527f274b5753bc2a873526e44bae648b363f47953a4f0c2234820809428fe7b7dafd60205260408120805460ff1916600117905533919063d8a8b5c760e01b905f5160206150ce5f395f51905f529080a4600190565b6001600160a01b0381165f9081527f55283b829f13bbb718097e9d9a0a5f887f18c6e64a14e5b6837b3d53eab0a8f3602052604090205460ff1661333d576001600160a01b03165f8181527f55283b829f13bbb718097e9d9a0a5f887f18c6e64a14e5b6837b3d53eab0a8f360205260408120805460ff1916600117905533919063d3e319af60e01b905f5160206150ce5f395f51905f529080a4600190565b6001600160a01b0381165f9081527f23cc607a5e9b003f25bdbf61a33e0b876542b92670e155052f357fc8d9d93a72602052604090205460ff1661333d576001600160a01b03165f8181527f23cc607a5e9b003f25bdbf61a33e0b876542b92670e155052f357fc8d9d93a7260205260408120805460ff19166001179055339190634958479d60e01b905f5160206150ce5f395f51905f529080a4600190565b6001600160a01b0381165f9081527f45511562325900b8d45e513a5f18f2d1dbd90118b7e5c34aa4c4f3b470e43fcd602052604090205460ff1661333d576001600160a01b03165f8181527f45511562325900b8d45e513a5f18f2d1dbd90118b7e5c34aa4c4f3b470e43fcd60205260408120805460ff1916600117905533919063c6d7271560e01b905f5160206150ce5f395f51905f529080a4600190565b6001600160a01b0381165f9081527ffec025993670f3cfe62fe0b8ba8699907af9264646da1759461a905da18f569c602052604090205460ff1661333d576001600160a01b03165f8181527ffec025993670f3cfe62fe0b8ba8699907af9264646da1759461a905da18f569c60205260408120805460ff191660011790553391906362250a9560e11b905f5160206150ce5f395f51905f529080a4600190565b6001600160a01b0381165f9081527f1dd6bf49940a66e2b9e775e4815e6e7139950009c60c13fc7c84fcd0a8976f2e602052604090205460ff1661333d576001600160a01b03165f8181527f1dd6bf49940a66e2b9e775e4815e6e7139950009c60c13fc7c84fcd0a8976f2e60205260408120805460ff19166001179055339190634bb5f31f60e11b905f5160206150ce5f395f51905f529080a4600190565b6001600160a01b0381165f9081527f0740732a750fc7c9be8bf07cf60eb78a7874ecd8bae4a47da79e539bcdcdd16b602052604090205460ff1661333d576001600160a01b03165f8181527f0740732a750fc7c9be8bf07cf60eb78a7874ecd8bae4a47da79e539bcdcdd16b60205260408120805460ff19166001179055339190632aedfbf360e01b905f5160206150ce5f395f51905f529080a4600190565b5f8181525f5160206151ae5f395f51905f52602090815260408083206001600160a01b038616845290915290205460ff166138a7575f8181525f5160206151ae5f395f51905f52602090815260408083206001600160a01b0395909516808452949091528120805460ff19166001179055339291905f5160206150ce5f395f51905f529080a4600190565b50505f90565b6001600160a01b0381165f9081525f51602061514e5f395f51905f52602052604090205460ff161561333d576001600160a01b03165f8181525f51602061514e5f395f51905f5260205260408120805460ff19169055339190638a70a0eb60e01b905f5160206151ee5f395f51905f529080a4600190565b6001600160a01b0381165f9081525f51602061526e5f395f51905f52602052604090205460ff161561333d576001600160a01b03165f8181525f51602061526e5f395f51905f5260205260408120805460ff191690553391906378b4401360e11b905f5160206151ee5f395f51905f529080a4600190565b6001600160a01b0381165f9081527fb7db2dd08fcb62d0c9e08c51941cae53c267786a0b75803fb7960902fc8ef97d602052604090205460ff161561333d576001600160a01b03165f8181527fb7db2dd08fcb62d0c9e08c51941cae53c267786a0b75803fb7960902fc8ef97d60205260408120805460ff191690553391905f5160206151ee5f395f51905f528180a4600190565b5f8181525f5160206151ae5f395f51905f52602090815260408083206001600160a01b038616845290915290205460ff16156138a7575f8181525f5160206151ae5f395f51905f52602090815260408083206001600160a01b0395909516808452949091528120805460ff19169055339291905f5160206151ee5f395f51905f529080a4600190565b60025f51602061520e5f395f51905f525414613ae45760025f51602061520e5f395f51905f5255565b633ee5aeb560e01b5f5260045ffd5b90816020910312610da257518015158103610da25790565b81546001600160a01b039081165f818152600d9093016020908152604090932054613b8093921680613be657505b600284018054604051633256b4d160e01b81526001600160a01b03909316600484015263ffffffff87166024840152604483015295909283919082905f9082906064820190565b03926001600160a01b03165af1908115612a52575f91613bb7575b50613ba557505050565b63ffffffff612fd59354921690614b2e565b613bd9915060203d602011613bdf575b613bd18183612cc1565b810190613af3565b5f613b9b565b503d613bc7565b9050613b39565b61ffff806006830154169216918211613c3057600790825f520160205260405f209061ffff82541615613c1e575090565b6321321e1960e11b5f5260045260245ffd5b506321321e1960e11b5f5260045260245ffd5b91908203918211612b3957565b90602080835192838152019201905f5b818110613c6d5750505090565b825180516001600160a01b031685526020908101518186015260409094019390920191600101613c60565b9035601e1982360301811215610da25701602081359101916001600160401b038211610da2578136038313610da257565b908060209392818452848401375f828201840152601f01601f1916010190565b905f905b60028210613cfa57505050565b6020806001928551815201930191019091613ced565b91613db691612d429694613dab9260c0865261ffff613d906060613d6e613d4f8a60c06080613d3f8b80613c98565b92909301526101408d0191613cc9565b613d5c6020890189613c98565b8c830360bf190160e08e015290613cc9565b956001600160a01b03613d8360408301612d71565b166101008b015201612d96565b16610120870152151560208601528482036040860152613206565b936060830190613ce9565b60a0818403910152612da5565b356001600160a01b0381168103610da25790565b3561ffff81168103610da25790565b7f47c1bdca9c75057bf1ca178a62fb0ef8908b21ee1ea9b9658ba5135fd3a3546d545f51602061528e5f395f51905f52935f936001600160a01b038316614ab0575b613eb8613e82613e388680614bcb565b92908761ffff613e94613e4e6020840184614bcb565b9190613e686060613e6160408801613dc3565b9601613dd7565b92604051988997602089019b60808d5260a08a0191613cc9565b878103601f1901604089015291613cc9565b6001600160a01b03909316606085015216608083015203601f198101835282612cc1565b519020938486527f47c1bdca9c75057bf1ca178a62fb0ef8908b21ee1ea9b9658ba5135fd3a3547060205260ff604087205416614aa157613f15613efe60408301613dc3565b613f0c602085015188614fb1565b90939193614feb565b6001600160a01b03908116911603614a92579061405093929160e082511515613fd661ffff613f448580614bcb565b919093613f546020880188614bcb565b613fa1613f6660408b96949601613dc3565b94613f7360608c01613dd7565b92613f8f6040519a8b9960208b019d8e526101008b0191613cc9565b888103601f190160408a015291613cc9565b6001600160a01b0390941660608601521660808401523060a08401524660c084015260e083015203601f198101835282612cc1565b51902093604060018060a01b035f51602061518e5f395f51905f5254169581519060208201526020815261400a8282612cc1565b81518098819263a850a90960e01b83527fb110049506439a07d78731efed3c809a6b13e0bab7cc9805f3bb45fb1e9a67e860048401528460248401526044830190613206565b0381885afa958615614a87578896614a0f575b5060405161407081612c6f565b888152606060208201528860408201528860608201528860808201528860a08201528860c08201528860e082015260606101008201529561ffff6140b660608501613dd7565b1661483e575b83511515906140cd60608501613dd7565b916140e4896040880151946060890151938b614e0e565b9060e08a015190893b1561475d57918c939161413761411c9461412660405198899788976365c4647560e01b89526004890190613ce9565b6044870190613ce9565b60e0608486015260e4850190612da5565b9160a484015260c48301520381895afa80156148335790899161481e575b5050825160608401518a5463ffffffff169890969091156147db577fde1faaa044a7216023ec32c06f91d9b5098d8bbd3b37f3662be3a729752ec9fc8961419e60408701613dc3565b6141d26141ab8880614bcb565b6141b860208b018b614bcb565b916141c560608d01613dd7565b9360405197889788614bfd565b0390a15b6141eb6141e560408601613dc3565b8c613281565b976141f6338d613281565b916040516382afd23b60e01b81528a6004820152602081602481855afa9081156147d0578d916147b1575b50156147a2576020602491604051928380926382afd23b60e01b82528760048301525afa908115614797578c91614778575b50156147695760405161426581612c8b565b6080863603126147655760405161427b81612c8b565b86356001600160401b038111614761576142989036908901612d27565b81526001600160401b0360208801351161475d576142bc3660208901358901612d27565b60208201526142cd60408801612d71565b60408201526142de60608801612d96565b6060820152815282602082015289604082015288606082015260018060a01b037f47c1bdca9c75057bf1ca178a62fb0ef8908b21ee1ea9b9658ba5135fd3a35474541680158015806146f8575b50156145b3575061433e60608701613dd7565b7f47c1bdca9c75057bf1ca178a62fb0ef8908b21ee1ea9b9658ba5135fd3a35474546001600160a01b03169290831561455c57505061439c8c9796959493926060926040519e8f80948193633fd6c07d60e11b835260048301614c4d565b03925af19a8b1561455157869b8791889161450f575b50909b995b875b8a518110156143e457808f8e6143de918f8f6001966143d791612f3c565b5191614f5a565b016143b9565b508b959a5061440192949950614407969b9193989d979c88614f5a565b84614f5a565b63ffffffff81541663ffffffff81146144fb57815463ffffffff1916600191820163ffffffff16179091559085527f47c1bdca9c75057bf1ca178a62fb0ef8908b21ee1ea9b9658ba5135fd3a3547060205260408520805460ff191690911790556001600160a01b03811661447d575b50505050565b825115156020840151926060604086015195015160018060a01b0384163b156103535786949392916144c786926040519889978896879563dd1a538760e01b875260048701613d10565b03926001600160a01b03165af180156105d7576144e6575b8080614477565b6144f1828092612cc1565b61037157806144df565b634e487b7160e01b87526011600452602487fd5b9c50505060608b3d606011614549575b8161452c60609383612cc1565b810103126105fc578a51604060208d01519c0151909b905f6143b2565b3d915061451f565b6040513d88823e3d90fd5b909c9796959493925061ffff161590506145865760608b01519860a060808d01519c0151906143b7565b7f47c1bdca9c75057bf1ca178a62fb0ef8908b21ee1ea9b9658ba5135fd3a35469549a508a9889906143b7565b91508b9a9c99508a9798506145e2925060409b939b519a8b8094819363e623f74760e01b835260048301614c4d565b03925af1968715611a7957849761463f575b50835b875181101561462f57806146298b8a60206146208561461860019885612f3c565b515193612f3c565b5101518b614f5a565b016145f7565b5093975093909450949094614407565b9096503d8085833e6146518183612cc1565b81019060208183031261033f578051906001600160401b0382116105fc570181601f8201121561033f5780519061468782612e31565b926146956040519485612cc1565b82845260208085019360061b830101918183116146f457602001925b8284106146c35750505050955f6145f4565b6040848303126146f457602060409182516146dd81612ca6565b8651815282870151838201528152019301926146b1565b8780fd5b905080614706575b5f61432b565b50604051636802d27b60e11b8152602081600481855afa908115614750578e91614731575b50614700565b61474a915060203d602011613bdf57613bd18183612cc1565b5f61472b565b8e604051903d90823e3d90fd5b8c80fd5b8d80fd5b8b80fd5b631385e31560e21b8b5260048bfd5b614791915060203d602011613bdf57613bd18183612cc1565b5f614253565b6040513d8e823e3d90fd5b6346d3a4a960e01b8c5260048cfd5b6147ca915060203d602011613bdf57613bd18183612cc1565b5f614221565b6040513d8f823e3d90fd5b7f5d681de577a7b3ff669022ffe837e270ef5b32ee4c1377045cf61b700d99e70a8961480960408701613dc3565b6148166141ab8880614bcb565b0390a16141d6565b8161482891612cc1565b6146f457875f614155565b6040513d8b823e3d90fd5b955061485561484f60608401613dd7565b8a613bed565b9560405161486281612c6f565b61ffff88541681526001880160405190818c825492614880846131ce565b80845293600181169081156149ed57506001146149ac575b506148a592500382612cc1565b6020820152600860028901549860408301998a526003810154606084015260048101546080840152600581015460a0840152600681015460c0840152600781015460e0840152018a6040518092839160208254918281520191845260208420935b81811061499357505061491b92500382612cc1565b61010082019081528198514310156149845751805161493c575b50506140bc565b60608601518151815111614975579061495491614d63565b908115614935575163fe88fc7160e01b8b5261ffff16600452602452604489fd5b630a8d477960e01b8c5260048cfd5b635a31d91b60e01b8b5260048bfd5b8454835260019485019486945060209093019201614906565b8e525060208d2090918d915b8183106149d15750509060206148a5928201015f614898565b60209193508060019154838588010152019101909183926149b8565b9050602092506148a594915060ff191682840152151560051b8201015f614898565b90955060403d604011614a80575b614a278183612cc1565b810190604081830312610ab35781601f82011215610ab35760405191614a4e604084612cc1565b82906040830111610aaf5781905b604083018210614a7057505050945f614063565b8151815260209182019101614a5c565b503d614a1d565b6040513d8a823e3d90fd5b632be1e1cb60e11b8652600486fd5b634510302360e11b8652600486fd5b8051602082015160408301516060840151909290919015156001600160a01b0387163b15610da257614af95f936040519586948594630502f5bd60e41b86528c60048701613d10565b0381836001600160a01b0389165af18015612a5257614b19575b50613e28565b614b269195505f90612cc1565b5f935f614b13565b90614b3a838284614d02565b15614b73579160039281614b56575b505001805460ff19169055565b6001830155614b6a60028301918254613c43565b90555f80614b49565b7f1e643658b8248efd3563f24d116430bf571d036bea3721d94e848890a00a10239260609260018060a01b039054169160405192835260208301526040820152a1565b600360ff910154166003811015612e55571590565b903590601e1981360301821215610da257018035906001600160401b038211610da257602001918136038313610da257565b9590614c469361ffff9563ffffffff60809895614c38949c9b9c60018060a01b03168b521660208a015260a060408a015260a0890191613cc9565b918683036060880152613cc9565b9416910152565b906020825260608151916080602085015261ffff82614c92614c7c8651608060a08a0152610120890190613206565b6020870151888203609f190160c08a0152613206565b9460018060a01b0360408201511660e0880152015116610100850152602081015160408501526040810151828501520151916080601f1982840301910152602080835192838152019201905f5b818110614cec5750505090565b8251845260209384019390920191600101614cdf565b9060ff6003830154166003811015612e555760010361327357600182015410156138a7576002015410614d3457600190565b5f90565b60ff5f51602061524e5f395f51905f525460401c1615614d5457565b631afcd79f60e31b5f5260045ffd5b5f90815b8151811080614e04575b15614de757614d808183612f3c565b51614d8b8486612f3c565b5103614dac57614d9d614da391612f2e565b92612f2e565b925b9291614d67565b91614db78383612f3c565b51614dc28286612f3c565b511015614dd857614dd290612f2e565b92614da5565b50614de39250612f3c565b5190565b91509150815181145f14614dfb5750505f90565b614de391612f3c565b5083518310614d71565b9092909161ffff16908115159081614f34575b5015614ec95760405163255da33160e01b8152600481019190915290602090829060249082906001600160a01b03165afa908115612a52575f91614e97575b50905b15614e82578060011b9080820460021490151715612b39576003900490565b80800460011481151715612b39576003900490565b90506020813d602011614ec1575b81614eb260209383612cc1565b81010312610da257515f614e60565b3d9150614ea5565b5060405163671b379360e01b815290602090829060049082906001600160a01b03165afa908115612a52575f91614f02575b5090614e63565b90506020813d602011614f2c575b81614f1d60209383612cc1565b81010312610da257515f614efb565b3d9150614f10565b610100810151511580159250614f4c575b505f614e21565b60e09150015115155f614f45565b600a90939193835f5201602052600260405f2001805490848201809211612b39575560405192835263ffffffff16917fd3f16e9d8d3fe0ea8a6e5f923fe57e1ae1af6d890ac6c371e8af6cc177a49b6590602090a3565b8151919060418303614fe157614fda9250602082015190606060408401519301515f1a9061504b565b9192909190565b50505f9160029190565b6004811015612e555780614ffd575050565b600181036150145763f645eedf60e01b5f5260045ffd5b6002810361502f575063fce698f760e01b5f5260045260245ffd5b6003146150395750565b6335e2f38360e21b5f5260045260245ffd5b91907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084116150c2579160209360809260ff5f9560405194855216868401526040830152606082015282805260015afa15612a52575f516001600160a01b038116156150b857905f905f90565b505f906001905f90565b5050505f916003919056fe2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d47c1bdca9c75057bf1ca178a62fb0ef8908b21ee1ea9b9658ba5135fd3a3546c47c1bdca9c75057bf1ca178a62fb0ef8908b21ee1ea9b9658ba5135fd3a3546a47c1bdca9c75057bf1ca178a62fb0ef8908b21ee1ea9b9658ba5135fd3a35473d480ef4a5d78515f50ce9d1a72eb9abf0b7795388f968bf633429bfc14cfe40bfe6065fb4e9872e2ad4479001655335380d83f70e163706cd65857449b84510047c1bdca9c75057bf1ca178a62fb0ef8908b21ee1ea9b9658ba5135fd3a3546b02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680047c1bdca9c75057bf1ca178a62fb0ef8908b21ee1ea9b9658ba5135fd3a35472f6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0047c1bdca9c75057bf1ca178a62fb0ef8908b21ee1ea9b9658ba5135fd3a35476f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00b8c1f5d3f58d95cae9dbd2762e9e87b212a41bbf0f42459bc1feb152e30e1cec47c1bdca9c75057bf1ca178a62fb0ef8908b21ee1ea9b9658ba5135fd3a35468a26469706673582212206eeff0aa0a50a0dab48a69a46182d9673195be31f2a9600908e1acda73e2c3a264736f6c634300081b0033",
}

// ContractAttestationCenterABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractAttestationCenterMetaData.ABI instead.
var ContractAttestationCenterABI = ContractAttestationCenterMetaData.ABI

// ContractAttestationCenterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractAttestationCenterMetaData.Bin instead.
var ContractAttestationCenterBin = ContractAttestationCenterMetaData.Bin

// DeployContractAttestationCenter deploys a new Ethereum contract, binding an instance of ContractAttestationCenter to it.
func DeployContractAttestationCenter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContractAttestationCenter, error) {
	parsed, err := ContractAttestationCenterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractAttestationCenterBin), backend)
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

	AvsLogic(opts *bind.CallOpts) (common.Address, error)

	BaseRewardFee(opts *bind.CallOpts) (*big.Int, error)

	BeforePaymentsLogic(opts *bind.CallOpts) (common.Address, error)

	GetOperatorPaymentDetail(opts *bind.CallOpts, _operatorId *big.Int) (IAttestationCenterPaymentDetails, error)

	GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error)

	GetTaskDefinitionMinimumVotingPower(opts *bind.CallOpts, _taskDefinitionId uint16) (*big.Int, error)

	GetTaskDefinitionRestrictedOperators(opts *bind.CallOpts, _taskDefinitionId uint16) ([]*big.Int, error)

	HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error)

	IsFlowPaused(opts *bind.CallOpts, _pausableFlow [4]byte) (bool, error)

	NumOfActiveOperators(opts *bind.CallOpts) (*big.Int, error)

	NumOfOperators(opts *bind.CallOpts) (*big.Int, error)

	NumOfTaskDefinitions(opts *bind.CallOpts) (uint16, error)

	NumOfTotalOperators(opts *bind.CallOpts) (*big.Int, error)

	Obls(opts *bind.CallOpts) (common.Address, error)

	OperatorsIdsByAddress(opts *bind.CallOpts, _operator common.Address) (*big.Int, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TaskNumber(opts *bind.CallOpts) (uint32, error)

	Vault(opts *bind.CallOpts) (common.Address, error)

	VotingPower(opts *bind.CallOpts, _operator common.Address) (*big.Int, error)
}

// ContractAttestationCenterTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractAttestationCenterTransacts interface {
	ClearBatchPayment(opts *bind.TransactOpts, _operators []IAttestationCenterPaymentRequestMessage, _paidTaskNumber *big.Int) (*types.Transaction, error)

	ClearPayment(opts *bind.TransactOpts, _operator common.Address, _lastPaidTaskNumber *big.Int, _amountClaimed *big.Int) (*types.Transaction, error)

	CreateNewTaskDefinition(opts *bind.TransactOpts, _name string, _taskDefinitionParams TaskDefinitionParams) (*types.Transaction, error)

	GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts, _avsGovernanceMultisigOwner common.Address, _operationsMultisig common.Address, _communityMultisig common.Address, _messageHandler common.Address, _obls common.Address, _vault common.Address, _isRewardsOnL2 bool) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts, _pausableFlow [4]byte) (*types.Transaction, error)

	RegisterToNetwork(opts *bind.TransactOpts, _operator common.Address, _votingPower *big.Int, _blsKey [4]*big.Int, _rewardsReceiver common.Address) (*types.Transaction, error)

	RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error)

	RequestBatchPayment(opts *bind.TransactOpts, _from *big.Int, _to *big.Int) (*types.Transaction, error)

	RequestBatchPayment0(opts *bind.TransactOpts) (*types.Transaction, error)

	RequestPayment(opts *bind.TransactOpts, _operatorId *big.Int) (*types.Transaction, error)

	RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	SetAvsLogic(opts *bind.TransactOpts, _avsLogic common.Address) (*types.Transaction, error)

	SetBeforePaymentsLogic(opts *bind.TransactOpts, _beforePaymentsLogic common.Address) (*types.Transaction, error)

	SetFeeCalculator(opts *bind.TransactOpts, _feeCalculator common.Address) (*types.Transaction, error)

	SetOblsSharesSyncer(opts *bind.TransactOpts, _oblsSharesSyncer common.Address) (*types.Transaction, error)

	SetTaskDefinitionMinVotingPower(opts *bind.TransactOpts, _taskDefinitionId uint16, _minimumVotingPower *big.Int) (*types.Transaction, error)

	SetTaskDefinitionRestrictedOperators(opts *bind.TransactOpts, _taskDefinitionId uint16, _restrictedOperatorIndexes []*big.Int) (*types.Transaction, error)

	SubmitTask(opts *bind.TransactOpts, _taskInfo IAttestationCenterTaskInfo, _isApproved bool, _tpSignature []byte, _taSignature [2]*big.Int, _attestersIds []*big.Int) (*types.Transaction, error)

	SubmitTask0(opts *bind.TransactOpts, _taskInfo IAttestationCenterTaskInfo, _taskSubmissionDetails IAttestationCenterTaskSubmissionDetails) (*types.Transaction, error)

	TransferAvsGovernanceMultisig(opts *bind.TransactOpts, _newAvsGovernanceMultisig common.Address) (*types.Transaction, error)

	TransferMessageHandler(opts *bind.TransactOpts, _newMessageHandler common.Address) (*types.Transaction, error)

	UnRegisterOperatorFromNetwork(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts, _pausableFlow [4]byte) (*types.Transaction, error)

	UpdateBlsKey(opts *bind.TransactOpts, _blsKey [4]*big.Int, _authSignature BLSAuthLibrarySignature) (*types.Transaction, error)
}

// ContractAttestationCenterFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractAttestationCenterFilters interface {
	FilterClearPaymentRejected(opts *bind.FilterOpts) (*ContractAttestationCenterClearPaymentRejectedIterator, error)
	WatchClearPaymentRejected(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterClearPaymentRejected) (event.Subscription, error)
	ParseClearPaymentRejected(log types.Log) (*ContractAttestationCenterClearPaymentRejected, error)

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

	FilterOperatorRegisteredToNetwork(opts *bind.FilterOpts) (*ContractAttestationCenterOperatorRegisteredToNetworkIterator, error)
	WatchOperatorRegisteredToNetwork(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterOperatorRegisteredToNetwork) (event.Subscription, error)
	ParseOperatorRegisteredToNetwork(log types.Log) (*ContractAttestationCenterOperatorRegisteredToNetwork, error)

	FilterOperatorUnregisteredFromNetwork(opts *bind.FilterOpts) (*ContractAttestationCenterOperatorUnregisteredFromNetworkIterator, error)
	WatchOperatorUnregisteredFromNetwork(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterOperatorUnregisteredFromNetwork) (event.Subscription, error)
	ParseOperatorUnregisteredFromNetwork(log types.Log) (*ContractAttestationCenterOperatorUnregisteredFromNetwork, error)

	FilterPaymentRequested(opts *bind.FilterOpts) (*ContractAttestationCenterPaymentRequestedIterator, error)
	WatchPaymentRequested(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterPaymentRequested) (event.Subscription, error)
	ParsePaymentRequested(log types.Log) (*ContractAttestationCenterPaymentRequested, error)

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

	FilterSetAvsGovernanceMultisig(opts *bind.FilterOpts) (*ContractAttestationCenterSetAvsGovernanceMultisigIterator, error)
	WatchSetAvsGovernanceMultisig(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterSetAvsGovernanceMultisig) (event.Subscription, error)
	ParseSetAvsGovernanceMultisig(log types.Log) (*ContractAttestationCenterSetAvsGovernanceMultisig, error)

	FilterSetAvsLogic(opts *bind.FilterOpts) (*ContractAttestationCenterSetAvsLogicIterator, error)
	WatchSetAvsLogic(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterSetAvsLogic) (event.Subscription, error)
	ParseSetAvsLogic(log types.Log) (*ContractAttestationCenterSetAvsLogic, error)

	FilterSetBeforePaymentsLogic(opts *bind.FilterOpts) (*ContractAttestationCenterSetBeforePaymentsLogicIterator, error)
	WatchSetBeforePaymentsLogic(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterSetBeforePaymentsLogic) (event.Subscription, error)
	ParseSetBeforePaymentsLogic(log types.Log) (*ContractAttestationCenterSetBeforePaymentsLogic, error)

	FilterSetFeeCalculator(opts *bind.FilterOpts) (*ContractAttestationCenterSetFeeCalculatorIterator, error)
	WatchSetFeeCalculator(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterSetFeeCalculator) (event.Subscription, error)
	ParseSetFeeCalculator(log types.Log) (*ContractAttestationCenterSetFeeCalculator, error)

	FilterSetMessageHandler(opts *bind.FilterOpts) (*ContractAttestationCenterSetMessageHandlerIterator, error)
	WatchSetMessageHandler(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterSetMessageHandler) (event.Subscription, error)
	ParseSetMessageHandler(log types.Log) (*ContractAttestationCenterSetMessageHandler, error)

	FilterSetMinimumTaskDefinitionVotingPower(opts *bind.FilterOpts) (*ContractAttestationCenterSetMinimumTaskDefinitionVotingPowerIterator, error)
	WatchSetMinimumTaskDefinitionVotingPower(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterSetMinimumTaskDefinitionVotingPower) (event.Subscription, error)
	ParseSetMinimumTaskDefinitionVotingPower(log types.Log) (*ContractAttestationCenterSetMinimumTaskDefinitionVotingPower, error)

	FilterSetRestrictedOperator(opts *bind.FilterOpts) (*ContractAttestationCenterSetRestrictedOperatorIterator, error)
	WatchSetRestrictedOperator(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterSetRestrictedOperator) (event.Subscription, error)
	ParseSetRestrictedOperator(log types.Log) (*ContractAttestationCenterSetRestrictedOperator, error)

	FilterTaskDefinitionCreated(opts *bind.FilterOpts) (*ContractAttestationCenterTaskDefinitionCreatedIterator, error)
	WatchTaskDefinitionCreated(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterTaskDefinitionCreated) (event.Subscription, error)
	ParseTaskDefinitionCreated(log types.Log) (*ContractAttestationCenterTaskDefinitionCreated, error)

	FilterTaskDefinitionRestrictedOperatorsModified(opts *bind.FilterOpts) (*ContractAttestationCenterTaskDefinitionRestrictedOperatorsModifiedIterator, error)
	WatchTaskDefinitionRestrictedOperatorsModified(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterTaskDefinitionRestrictedOperatorsModified) (event.Subscription, error)
	ParseTaskDefinitionRestrictedOperatorsModified(log types.Log) (*ContractAttestationCenterTaskDefinitionRestrictedOperatorsModified, error)

	FilterTaskRejected(opts *bind.FilterOpts) (*ContractAttestationCenterTaskRejectedIterator, error)
	WatchTaskRejected(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterTaskRejected) (event.Subscription, error)
	ParseTaskRejected(log types.Log) (*ContractAttestationCenterTaskRejected, error)

	FilterTaskSubmitted(opts *bind.FilterOpts) (*ContractAttestationCenterTaskSubmittedIterator, error)
	WatchTaskSubmitted(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterTaskSubmitted) (event.Subscription, error)
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

// BaseRewardFee is a free data retrieval call binding the contract method 0x3428c126.
//
// Solidity: function baseRewardFee() view returns(uint256)
func (_ContractAttestationCenter *ContractAttestationCenterCaller) BaseRewardFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractAttestationCenter.contract.Call(opts, &out, "baseRewardFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseRewardFee is a free data retrieval call binding the contract method 0x3428c126.
//
// Solidity: function baseRewardFee() view returns(uint256)
func (_ContractAttestationCenter *ContractAttestationCenterSession) BaseRewardFee() (*big.Int, error) {
	return _ContractAttestationCenter.Contract.BaseRewardFee(&_ContractAttestationCenter.CallOpts)
}

// BaseRewardFee is a free data retrieval call binding the contract method 0x3428c126.
//
// Solidity: function baseRewardFee() view returns(uint256)
func (_ContractAttestationCenter *ContractAttestationCenterCallerSession) BaseRewardFee() (*big.Int, error) {
	return _ContractAttestationCenter.Contract.BaseRewardFee(&_ContractAttestationCenter.CallOpts)
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

// GetTaskDefinitionRestrictedOperators is a free data retrieval call binding the contract method 0x4e2ce53f.
//
// Solidity: function getTaskDefinitionRestrictedOperators(uint16 _taskDefinitionId) view returns(uint256[])
func (_ContractAttestationCenter *ContractAttestationCenterCaller) GetTaskDefinitionRestrictedOperators(opts *bind.CallOpts, _taskDefinitionId uint16) ([]*big.Int, error) {
	var out []interface{}
	err := _ContractAttestationCenter.contract.Call(opts, &out, "getTaskDefinitionRestrictedOperators", _taskDefinitionId)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetTaskDefinitionRestrictedOperators is a free data retrieval call binding the contract method 0x4e2ce53f.
//
// Solidity: function getTaskDefinitionRestrictedOperators(uint16 _taskDefinitionId) view returns(uint256[])
func (_ContractAttestationCenter *ContractAttestationCenterSession) GetTaskDefinitionRestrictedOperators(_taskDefinitionId uint16) ([]*big.Int, error) {
	return _ContractAttestationCenter.Contract.GetTaskDefinitionRestrictedOperators(&_ContractAttestationCenter.CallOpts, _taskDefinitionId)
}

// GetTaskDefinitionRestrictedOperators is a free data retrieval call binding the contract method 0x4e2ce53f.
//
// Solidity: function getTaskDefinitionRestrictedOperators(uint16 _taskDefinitionId) view returns(uint256[])
func (_ContractAttestationCenter *ContractAttestationCenterCallerSession) GetTaskDefinitionRestrictedOperators(_taskDefinitionId uint16) ([]*big.Int, error) {
	return _ContractAttestationCenter.Contract.GetTaskDefinitionRestrictedOperators(&_ContractAttestationCenter.CallOpts, _taskDefinitionId)
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

// NumOfOperators is a free data retrieval call binding the contract method 0x6ade02da.
//
// Solidity: function numOfOperators() view returns(uint256)
func (_ContractAttestationCenter *ContractAttestationCenterCaller) NumOfOperators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractAttestationCenter.contract.Call(opts, &out, "numOfOperators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumOfOperators is a free data retrieval call binding the contract method 0x6ade02da.
//
// Solidity: function numOfOperators() view returns(uint256)
func (_ContractAttestationCenter *ContractAttestationCenterSession) NumOfOperators() (*big.Int, error) {
	return _ContractAttestationCenter.Contract.NumOfOperators(&_ContractAttestationCenter.CallOpts)
}

// NumOfOperators is a free data retrieval call binding the contract method 0x6ade02da.
//
// Solidity: function numOfOperators() view returns(uint256)
func (_ContractAttestationCenter *ContractAttestationCenterCallerSession) NumOfOperators() (*big.Int, error) {
	return _ContractAttestationCenter.Contract.NumOfOperators(&_ContractAttestationCenter.CallOpts)
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

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_ContractAttestationCenter *ContractAttestationCenterCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAttestationCenter.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_ContractAttestationCenter *ContractAttestationCenterSession) Vault() (common.Address, error) {
	return _ContractAttestationCenter.Contract.Vault(&_ContractAttestationCenter.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_ContractAttestationCenter *ContractAttestationCenterCallerSession) Vault() (common.Address, error) {
	return _ContractAttestationCenter.Contract.Vault(&_ContractAttestationCenter.CallOpts)
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

// ClearPayment is a paid mutator transaction binding the contract method 0x242a76a4.
//
// Solidity: function clearPayment(address _operator, uint256 _lastPaidTaskNumber, uint256 _amountClaimed) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) ClearPayment(opts *bind.TransactOpts, _operator common.Address, _lastPaidTaskNumber *big.Int, _amountClaimed *big.Int) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "clearPayment", _operator, _lastPaidTaskNumber, _amountClaimed)
}

// ClearPayment is a paid mutator transaction binding the contract method 0x242a76a4.
//
// Solidity: function clearPayment(address _operator, uint256 _lastPaidTaskNumber, uint256 _amountClaimed) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) ClearPayment(_operator common.Address, _lastPaidTaskNumber *big.Int, _amountClaimed *big.Int) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.ClearPayment(&_ContractAttestationCenter.TransactOpts, _operator, _lastPaidTaskNumber, _amountClaimed)
}

// ClearPayment is a paid mutator transaction binding the contract method 0x242a76a4.
//
// Solidity: function clearPayment(address _operator, uint256 _lastPaidTaskNumber, uint256 _amountClaimed) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) ClearPayment(_operator common.Address, _lastPaidTaskNumber *big.Int, _amountClaimed *big.Int) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.ClearPayment(&_ContractAttestationCenter.TransactOpts, _operator, _lastPaidTaskNumber, _amountClaimed)
}

// CreateNewTaskDefinition is a paid mutator transaction binding the contract method 0x0c62bf0d.
//
// Solidity: function createNewTaskDefinition(string _name, (uint256,uint256,uint256,uint256,uint256,uint256,uint256[]) _taskDefinitionParams) returns(uint16 _id)
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) CreateNewTaskDefinition(opts *bind.TransactOpts, _name string, _taskDefinitionParams TaskDefinitionParams) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "createNewTaskDefinition", _name, _taskDefinitionParams)
}

// CreateNewTaskDefinition is a paid mutator transaction binding the contract method 0x0c62bf0d.
//
// Solidity: function createNewTaskDefinition(string _name, (uint256,uint256,uint256,uint256,uint256,uint256,uint256[]) _taskDefinitionParams) returns(uint16 _id)
func (_ContractAttestationCenter *ContractAttestationCenterSession) CreateNewTaskDefinition(_name string, _taskDefinitionParams TaskDefinitionParams) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.CreateNewTaskDefinition(&_ContractAttestationCenter.TransactOpts, _name, _taskDefinitionParams)
}

// CreateNewTaskDefinition is a paid mutator transaction binding the contract method 0x0c62bf0d.
//
// Solidity: function createNewTaskDefinition(string _name, (uint256,uint256,uint256,uint256,uint256,uint256,uint256[]) _taskDefinitionParams) returns(uint16 _id)
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) CreateNewTaskDefinition(_name string, _taskDefinitionParams TaskDefinitionParams) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.CreateNewTaskDefinition(&_ContractAttestationCenter.TransactOpts, _name, _taskDefinitionParams)
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

// Initialize is a paid mutator transaction binding the contract method 0xd9378a59.
//
// Solidity: function initialize(address _avsGovernanceMultisigOwner, address _operationsMultisig, address _communityMultisig, address _messageHandler, address _obls, address _vault, bool _isRewardsOnL2) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) Initialize(opts *bind.TransactOpts, _avsGovernanceMultisigOwner common.Address, _operationsMultisig common.Address, _communityMultisig common.Address, _messageHandler common.Address, _obls common.Address, _vault common.Address, _isRewardsOnL2 bool) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "initialize", _avsGovernanceMultisigOwner, _operationsMultisig, _communityMultisig, _messageHandler, _obls, _vault, _isRewardsOnL2)
}

// Initialize is a paid mutator transaction binding the contract method 0xd9378a59.
//
// Solidity: function initialize(address _avsGovernanceMultisigOwner, address _operationsMultisig, address _communityMultisig, address _messageHandler, address _obls, address _vault, bool _isRewardsOnL2) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) Initialize(_avsGovernanceMultisigOwner common.Address, _operationsMultisig common.Address, _communityMultisig common.Address, _messageHandler common.Address, _obls common.Address, _vault common.Address, _isRewardsOnL2 bool) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.Initialize(&_ContractAttestationCenter.TransactOpts, _avsGovernanceMultisigOwner, _operationsMultisig, _communityMultisig, _messageHandler, _obls, _vault, _isRewardsOnL2)
}

// Initialize is a paid mutator transaction binding the contract method 0xd9378a59.
//
// Solidity: function initialize(address _avsGovernanceMultisigOwner, address _operationsMultisig, address _communityMultisig, address _messageHandler, address _obls, address _vault, bool _isRewardsOnL2) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) Initialize(_avsGovernanceMultisigOwner common.Address, _operationsMultisig common.Address, _communityMultisig common.Address, _messageHandler common.Address, _obls common.Address, _vault common.Address, _isRewardsOnL2 bool) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.Initialize(&_ContractAttestationCenter.TransactOpts, _avsGovernanceMultisigOwner, _operationsMultisig, _communityMultisig, _messageHandler, _obls, _vault, _isRewardsOnL2)
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

// RequestPayment is a paid mutator transaction binding the contract method 0x5de988ab.
//
// Solidity: function requestPayment(uint256 _operatorId) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) RequestPayment(opts *bind.TransactOpts, _operatorId *big.Int) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "requestPayment", _operatorId)
}

// RequestPayment is a paid mutator transaction binding the contract method 0x5de988ab.
//
// Solidity: function requestPayment(uint256 _operatorId) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) RequestPayment(_operatorId *big.Int) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.RequestPayment(&_ContractAttestationCenter.TransactOpts, _operatorId)
}

// RequestPayment is a paid mutator transaction binding the contract method 0x5de988ab.
//
// Solidity: function requestPayment(uint256 _operatorId) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) RequestPayment(_operatorId *big.Int) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.RequestPayment(&_ContractAttestationCenter.TransactOpts, _operatorId)
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

// SetAvsLogic is a paid mutator transaction binding the contract method 0x008fd386.
//
// Solidity: function setAvsLogic(address _avsLogic) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) SetAvsLogic(opts *bind.TransactOpts, _avsLogic common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "setAvsLogic", _avsLogic)
}

// SetAvsLogic is a paid mutator transaction binding the contract method 0x008fd386.
//
// Solidity: function setAvsLogic(address _avsLogic) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) SetAvsLogic(_avsLogic common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SetAvsLogic(&_ContractAttestationCenter.TransactOpts, _avsLogic)
}

// SetAvsLogic is a paid mutator transaction binding the contract method 0x008fd386.
//
// Solidity: function setAvsLogic(address _avsLogic) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) SetAvsLogic(_avsLogic common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SetAvsLogic(&_ContractAttestationCenter.TransactOpts, _avsLogic)
}

// SetBeforePaymentsLogic is a paid mutator transaction binding the contract method 0x11a95e38.
//
// Solidity: function setBeforePaymentsLogic(address _beforePaymentsLogic) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) SetBeforePaymentsLogic(opts *bind.TransactOpts, _beforePaymentsLogic common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "setBeforePaymentsLogic", _beforePaymentsLogic)
}

// SetBeforePaymentsLogic is a paid mutator transaction binding the contract method 0x11a95e38.
//
// Solidity: function setBeforePaymentsLogic(address _beforePaymentsLogic) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) SetBeforePaymentsLogic(_beforePaymentsLogic common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SetBeforePaymentsLogic(&_ContractAttestationCenter.TransactOpts, _beforePaymentsLogic)
}

// SetBeforePaymentsLogic is a paid mutator transaction binding the contract method 0x11a95e38.
//
// Solidity: function setBeforePaymentsLogic(address _beforePaymentsLogic) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) SetBeforePaymentsLogic(_beforePaymentsLogic common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SetBeforePaymentsLogic(&_ContractAttestationCenter.TransactOpts, _beforePaymentsLogic)
}

// SetFeeCalculator is a paid mutator transaction binding the contract method 0x8c66d04f.
//
// Solidity: function setFeeCalculator(address _feeCalculator) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) SetFeeCalculator(opts *bind.TransactOpts, _feeCalculator common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "setFeeCalculator", _feeCalculator)
}

// SetFeeCalculator is a paid mutator transaction binding the contract method 0x8c66d04f.
//
// Solidity: function setFeeCalculator(address _feeCalculator) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) SetFeeCalculator(_feeCalculator common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SetFeeCalculator(&_ContractAttestationCenter.TransactOpts, _feeCalculator)
}

// SetFeeCalculator is a paid mutator transaction binding the contract method 0x8c66d04f.
//
// Solidity: function setFeeCalculator(address _feeCalculator) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) SetFeeCalculator(_feeCalculator common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SetFeeCalculator(&_ContractAttestationCenter.TransactOpts, _feeCalculator)
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

// SetTaskDefinitionMinVotingPower is a paid mutator transaction binding the contract method 0x64ada5d0.
//
// Solidity: function setTaskDefinitionMinVotingPower(uint16 _taskDefinitionId, uint256 _minimumVotingPower) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) SetTaskDefinitionMinVotingPower(opts *bind.TransactOpts, _taskDefinitionId uint16, _minimumVotingPower *big.Int) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "setTaskDefinitionMinVotingPower", _taskDefinitionId, _minimumVotingPower)
}

// SetTaskDefinitionMinVotingPower is a paid mutator transaction binding the contract method 0x64ada5d0.
//
// Solidity: function setTaskDefinitionMinVotingPower(uint16 _taskDefinitionId, uint256 _minimumVotingPower) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) SetTaskDefinitionMinVotingPower(_taskDefinitionId uint16, _minimumVotingPower *big.Int) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SetTaskDefinitionMinVotingPower(&_ContractAttestationCenter.TransactOpts, _taskDefinitionId, _minimumVotingPower)
}

// SetTaskDefinitionMinVotingPower is a paid mutator transaction binding the contract method 0x64ada5d0.
//
// Solidity: function setTaskDefinitionMinVotingPower(uint16 _taskDefinitionId, uint256 _minimumVotingPower) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) SetTaskDefinitionMinVotingPower(_taskDefinitionId uint16, _minimumVotingPower *big.Int) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SetTaskDefinitionMinVotingPower(&_ContractAttestationCenter.TransactOpts, _taskDefinitionId, _minimumVotingPower)
}

// SetTaskDefinitionRestrictedOperators is a paid mutator transaction binding the contract method 0xc8c9e7ab.
//
// Solidity: function setTaskDefinitionRestrictedOperators(uint16 _taskDefinitionId, uint256[] _restrictedOperatorIndexes) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) SetTaskDefinitionRestrictedOperators(opts *bind.TransactOpts, _taskDefinitionId uint16, _restrictedOperatorIndexes []*big.Int) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "setTaskDefinitionRestrictedOperators", _taskDefinitionId, _restrictedOperatorIndexes)
}

// SetTaskDefinitionRestrictedOperators is a paid mutator transaction binding the contract method 0xc8c9e7ab.
//
// Solidity: function setTaskDefinitionRestrictedOperators(uint16 _taskDefinitionId, uint256[] _restrictedOperatorIndexes) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) SetTaskDefinitionRestrictedOperators(_taskDefinitionId uint16, _restrictedOperatorIndexes []*big.Int) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SetTaskDefinitionRestrictedOperators(&_ContractAttestationCenter.TransactOpts, _taskDefinitionId, _restrictedOperatorIndexes)
}

// SetTaskDefinitionRestrictedOperators is a paid mutator transaction binding the contract method 0xc8c9e7ab.
//
// Solidity: function setTaskDefinitionRestrictedOperators(uint16 _taskDefinitionId, uint256[] _restrictedOperatorIndexes) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) SetTaskDefinitionRestrictedOperators(_taskDefinitionId uint16, _restrictedOperatorIndexes []*big.Int) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SetTaskDefinitionRestrictedOperators(&_ContractAttestationCenter.TransactOpts, _taskDefinitionId, _restrictedOperatorIndexes)
}

// SubmitTask is a paid mutator transaction binding the contract method 0x7d5f32bc.
//
// Solidity: function submitTask((string,bytes,address,uint16) _taskInfo, bool _isApproved, bytes _tpSignature, uint256[2] _taSignature, uint256[] _attestersIds) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) SubmitTask(opts *bind.TransactOpts, _taskInfo IAttestationCenterTaskInfo, _isApproved bool, _tpSignature []byte, _taSignature [2]*big.Int, _attestersIds []*big.Int) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "submitTask", _taskInfo, _isApproved, _tpSignature, _taSignature, _attestersIds)
}

// SubmitTask is a paid mutator transaction binding the contract method 0x7d5f32bc.
//
// Solidity: function submitTask((string,bytes,address,uint16) _taskInfo, bool _isApproved, bytes _tpSignature, uint256[2] _taSignature, uint256[] _attestersIds) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) SubmitTask(_taskInfo IAttestationCenterTaskInfo, _isApproved bool, _tpSignature []byte, _taSignature [2]*big.Int, _attestersIds []*big.Int) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SubmitTask(&_ContractAttestationCenter.TransactOpts, _taskInfo, _isApproved, _tpSignature, _taSignature, _attestersIds)
}

// SubmitTask is a paid mutator transaction binding the contract method 0x7d5f32bc.
//
// Solidity: function submitTask((string,bytes,address,uint16) _taskInfo, bool _isApproved, bytes _tpSignature, uint256[2] _taSignature, uint256[] _attestersIds) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) SubmitTask(_taskInfo IAttestationCenterTaskInfo, _isApproved bool, _tpSignature []byte, _taSignature [2]*big.Int, _attestersIds []*big.Int) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SubmitTask(&_ContractAttestationCenter.TransactOpts, _taskInfo, _isApproved, _tpSignature, _taSignature, _attestersIds)
}

// SubmitTask0 is a paid mutator transaction binding the contract method 0xfff768e3.
//
// Solidity: function submitTask((string,bytes,address,uint16) _taskInfo, (bool,bytes,uint256[2],uint256[]) _taskSubmissionDetails) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) SubmitTask0(opts *bind.TransactOpts, _taskInfo IAttestationCenterTaskInfo, _taskSubmissionDetails IAttestationCenterTaskSubmissionDetails) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "submitTask0", _taskInfo, _taskSubmissionDetails)
}

// SubmitTask0 is a paid mutator transaction binding the contract method 0xfff768e3.
//
// Solidity: function submitTask((string,bytes,address,uint16) _taskInfo, (bool,bytes,uint256[2],uint256[]) _taskSubmissionDetails) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) SubmitTask0(_taskInfo IAttestationCenterTaskInfo, _taskSubmissionDetails IAttestationCenterTaskSubmissionDetails) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SubmitTask0(&_ContractAttestationCenter.TransactOpts, _taskInfo, _taskSubmissionDetails)
}

// SubmitTask0 is a paid mutator transaction binding the contract method 0xfff768e3.
//
// Solidity: function submitTask((string,bytes,address,uint16) _taskInfo, (bool,bytes,uint256[2],uint256[]) _taskSubmissionDetails) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) SubmitTask0(_taskInfo IAttestationCenterTaskInfo, _taskSubmissionDetails IAttestationCenterTaskSubmissionDetails) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SubmitTask0(&_ContractAttestationCenter.TransactOpts, _taskInfo, _taskSubmissionDetails)
}

// TransferAvsGovernanceMultisig is a paid mutator transaction binding the contract method 0x513c52ba.
//
// Solidity: function transferAvsGovernanceMultisig(address _newAvsGovernanceMultisig) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) TransferAvsGovernanceMultisig(opts *bind.TransactOpts, _newAvsGovernanceMultisig common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "transferAvsGovernanceMultisig", _newAvsGovernanceMultisig)
}

// TransferAvsGovernanceMultisig is a paid mutator transaction binding the contract method 0x513c52ba.
//
// Solidity: function transferAvsGovernanceMultisig(address _newAvsGovernanceMultisig) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) TransferAvsGovernanceMultisig(_newAvsGovernanceMultisig common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.TransferAvsGovernanceMultisig(&_ContractAttestationCenter.TransactOpts, _newAvsGovernanceMultisig)
}

// TransferAvsGovernanceMultisig is a paid mutator transaction binding the contract method 0x513c52ba.
//
// Solidity: function transferAvsGovernanceMultisig(address _newAvsGovernanceMultisig) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) TransferAvsGovernanceMultisig(_newAvsGovernanceMultisig common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.TransferAvsGovernanceMultisig(&_ContractAttestationCenter.TransactOpts, _newAvsGovernanceMultisig)
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
// Solidity: event ClearPaymentRejected(address operator, uint256 requestedTaskNumber, uint256 requestedAmountClaimed)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) FilterClearPaymentRejected(opts *bind.FilterOpts) (*ContractAttestationCenterClearPaymentRejectedIterator, error) {

	logs, sub, err := _ContractAttestationCenter.contract.FilterLogs(opts, "ClearPaymentRejected")
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterClearPaymentRejectedIterator{contract: _ContractAttestationCenter.contract, event: "ClearPaymentRejected", logs: logs, sub: sub}, nil
}

// WatchClearPaymentRejected is a free log subscription operation binding the contract event 0x1e643658b8248efd3563f24d116430bf571d036bea3721d94e848890a00a1023.
//
// Solidity: event ClearPaymentRejected(address operator, uint256 requestedTaskNumber, uint256 requestedAmountClaimed)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) WatchClearPaymentRejected(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterClearPaymentRejected) (event.Subscription, error) {

	logs, sub, err := _ContractAttestationCenter.contract.WatchLogs(opts, "ClearPaymentRejected")
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
// Solidity: event ClearPaymentRejected(address operator, uint256 requestedTaskNumber, uint256 requestedAmountClaimed)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) ParseClearPaymentRejected(log types.Log) (*ContractAttestationCenterClearPaymentRejected, error) {
	event := new(ContractAttestationCenterClearPaymentRejected)
	if err := _ContractAttestationCenter.contract.UnpackLog(event, "ClearPaymentRejected", log); err != nil {
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
// Solidity: event OperatorRegisteredToNetwork(address operator, uint256 votingPower)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) FilterOperatorRegisteredToNetwork(opts *bind.FilterOpts) (*ContractAttestationCenterOperatorRegisteredToNetworkIterator, error) {

	logs, sub, err := _ContractAttestationCenter.contract.FilterLogs(opts, "OperatorRegisteredToNetwork")
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterOperatorRegisteredToNetworkIterator{contract: _ContractAttestationCenter.contract, event: "OperatorRegisteredToNetwork", logs: logs, sub: sub}, nil
}

// WatchOperatorRegisteredToNetwork is a free log subscription operation binding the contract event 0x16c1a2a8195923d655fe84191b37c746f4385a5c32c038578958b29f52daa1a8.
//
// Solidity: event OperatorRegisteredToNetwork(address operator, uint256 votingPower)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) WatchOperatorRegisteredToNetwork(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterOperatorRegisteredToNetwork) (event.Subscription, error) {

	logs, sub, err := _ContractAttestationCenter.contract.WatchLogs(opts, "OperatorRegisteredToNetwork")
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
// Solidity: event OperatorRegisteredToNetwork(address operator, uint256 votingPower)
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
// Solidity: event OperatorUnregisteredFromNetwork(uint256 operatorId)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) FilterOperatorUnregisteredFromNetwork(opts *bind.FilterOpts) (*ContractAttestationCenterOperatorUnregisteredFromNetworkIterator, error) {

	logs, sub, err := _ContractAttestationCenter.contract.FilterLogs(opts, "OperatorUnregisteredFromNetwork")
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterOperatorUnregisteredFromNetworkIterator{contract: _ContractAttestationCenter.contract, event: "OperatorUnregisteredFromNetwork", logs: logs, sub: sub}, nil
}

// WatchOperatorUnregisteredFromNetwork is a free log subscription operation binding the contract event 0xda04f7db725bc5a9ad418baf26d08e9f24561a7cc119bfe1dd26bfebfc175db3.
//
// Solidity: event OperatorUnregisteredFromNetwork(uint256 operatorId)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) WatchOperatorUnregisteredFromNetwork(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterOperatorUnregisteredFromNetwork) (event.Subscription, error) {

	logs, sub, err := _ContractAttestationCenter.contract.WatchLogs(opts, "OperatorUnregisteredFromNetwork")
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
// Solidity: event OperatorUnregisteredFromNetwork(uint256 operatorId)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) ParseOperatorUnregisteredFromNetwork(log types.Log) (*ContractAttestationCenterOperatorUnregisteredFromNetwork, error) {
	event := new(ContractAttestationCenterOperatorUnregisteredFromNetwork)
	if err := _ContractAttestationCenter.contract.UnpackLog(event, "OperatorUnregisteredFromNetwork", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAttestationCenterPaymentRequestedIterator is returned from FilterPaymentRequested and is used to iterate over the raw logs and unpacked data for PaymentRequested events raised by the ContractAttestationCenter contract.
type ContractAttestationCenterPaymentRequestedIterator struct {
	Event *ContractAttestationCenterPaymentRequested // Event containing the contract specifics and raw log

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
func (it *ContractAttestationCenterPaymentRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAttestationCenterPaymentRequested)
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
		it.Event = new(ContractAttestationCenterPaymentRequested)
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
func (it *ContractAttestationCenterPaymentRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAttestationCenterPaymentRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAttestationCenterPaymentRequested represents a PaymentRequested event raised by the ContractAttestationCenter contract.
type ContractAttestationCenterPaymentRequested struct {
	Operator           common.Address
	LastPaidTaskNumber *big.Int
	FeeToClaim         *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterPaymentRequested is a free log retrieval operation binding the contract event 0x34682c7a1451dbbbc9e3be4912a8f466eba3a8c72e9bcb5cb3a61e423a9c6973.
//
// Solidity: event PaymentRequested(address operator, uint256 lastPaidTaskNumber, uint256 feeToClaim)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) FilterPaymentRequested(opts *bind.FilterOpts) (*ContractAttestationCenterPaymentRequestedIterator, error) {

	logs, sub, err := _ContractAttestationCenter.contract.FilterLogs(opts, "PaymentRequested")
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterPaymentRequestedIterator{contract: _ContractAttestationCenter.contract, event: "PaymentRequested", logs: logs, sub: sub}, nil
}

// WatchPaymentRequested is a free log subscription operation binding the contract event 0x34682c7a1451dbbbc9e3be4912a8f466eba3a8c72e9bcb5cb3a61e423a9c6973.
//
// Solidity: event PaymentRequested(address operator, uint256 lastPaidTaskNumber, uint256 feeToClaim)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) WatchPaymentRequested(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterPaymentRequested) (event.Subscription, error) {

	logs, sub, err := _ContractAttestationCenter.contract.WatchLogs(opts, "PaymentRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAttestationCenterPaymentRequested)
				if err := _ContractAttestationCenter.contract.UnpackLog(event, "PaymentRequested", log); err != nil {
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

// ParsePaymentRequested is a log parse operation binding the contract event 0x34682c7a1451dbbbc9e3be4912a8f466eba3a8c72e9bcb5cb3a61e423a9c6973.
//
// Solidity: event PaymentRequested(address operator, uint256 lastPaidTaskNumber, uint256 feeToClaim)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) ParsePaymentRequested(log types.Log) (*ContractAttestationCenterPaymentRequested, error) {
	event := new(ContractAttestationCenterPaymentRequested)
	if err := _ContractAttestationCenter.contract.UnpackLog(event, "PaymentRequested", log); err != nil {
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

// ContractAttestationCenterSetAvsGovernanceMultisigIterator is returned from FilterSetAvsGovernanceMultisig and is used to iterate over the raw logs and unpacked data for SetAvsGovernanceMultisig events raised by the ContractAttestationCenter contract.
type ContractAttestationCenterSetAvsGovernanceMultisigIterator struct {
	Event *ContractAttestationCenterSetAvsGovernanceMultisig // Event containing the contract specifics and raw log

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
func (it *ContractAttestationCenterSetAvsGovernanceMultisigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAttestationCenterSetAvsGovernanceMultisig)
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
		it.Event = new(ContractAttestationCenterSetAvsGovernanceMultisig)
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
func (it *ContractAttestationCenterSetAvsGovernanceMultisigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAttestationCenterSetAvsGovernanceMultisigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAttestationCenterSetAvsGovernanceMultisig represents a SetAvsGovernanceMultisig event raised by the ContractAttestationCenter contract.
type ContractAttestationCenterSetAvsGovernanceMultisig struct {
	NewAvsGovernanceMultisig common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSetAvsGovernanceMultisig is a free log retrieval operation binding the contract event 0x024e98b7d808a3ddb028252dc95dfdcb165a0ca59fcff8984b4fecf9a7222649.
//
// Solidity: event SetAvsGovernanceMultisig(address newAvsGovernanceMultisig)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) FilterSetAvsGovernanceMultisig(opts *bind.FilterOpts) (*ContractAttestationCenterSetAvsGovernanceMultisigIterator, error) {

	logs, sub, err := _ContractAttestationCenter.contract.FilterLogs(opts, "SetAvsGovernanceMultisig")
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterSetAvsGovernanceMultisigIterator{contract: _ContractAttestationCenter.contract, event: "SetAvsGovernanceMultisig", logs: logs, sub: sub}, nil
}

// WatchSetAvsGovernanceMultisig is a free log subscription operation binding the contract event 0x024e98b7d808a3ddb028252dc95dfdcb165a0ca59fcff8984b4fecf9a7222649.
//
// Solidity: event SetAvsGovernanceMultisig(address newAvsGovernanceMultisig)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) WatchSetAvsGovernanceMultisig(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterSetAvsGovernanceMultisig) (event.Subscription, error) {

	logs, sub, err := _ContractAttestationCenter.contract.WatchLogs(opts, "SetAvsGovernanceMultisig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAttestationCenterSetAvsGovernanceMultisig)
				if err := _ContractAttestationCenter.contract.UnpackLog(event, "SetAvsGovernanceMultisig", log); err != nil {
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
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) ParseSetAvsGovernanceMultisig(log types.Log) (*ContractAttestationCenterSetAvsGovernanceMultisig, error) {
	event := new(ContractAttestationCenterSetAvsGovernanceMultisig)
	if err := _ContractAttestationCenter.contract.UnpackLog(event, "SetAvsGovernanceMultisig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAttestationCenterSetAvsLogicIterator is returned from FilterSetAvsLogic and is used to iterate over the raw logs and unpacked data for SetAvsLogic events raised by the ContractAttestationCenter contract.
type ContractAttestationCenterSetAvsLogicIterator struct {
	Event *ContractAttestationCenterSetAvsLogic // Event containing the contract specifics and raw log

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
func (it *ContractAttestationCenterSetAvsLogicIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAttestationCenterSetAvsLogic)
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
		it.Event = new(ContractAttestationCenterSetAvsLogic)
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
func (it *ContractAttestationCenterSetAvsLogicIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAttestationCenterSetAvsLogicIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAttestationCenterSetAvsLogic represents a SetAvsLogic event raised by the ContractAttestationCenter contract.
type ContractAttestationCenterSetAvsLogic struct {
	AvsLogic common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetAvsLogic is a free log retrieval operation binding the contract event 0xdf0d3b0bf99a87fc195d045bc7ec61d3e2619e6a49dd3f5cb69102b8c9702034.
//
// Solidity: event SetAvsLogic(address avsLogic)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) FilterSetAvsLogic(opts *bind.FilterOpts) (*ContractAttestationCenterSetAvsLogicIterator, error) {

	logs, sub, err := _ContractAttestationCenter.contract.FilterLogs(opts, "SetAvsLogic")
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterSetAvsLogicIterator{contract: _ContractAttestationCenter.contract, event: "SetAvsLogic", logs: logs, sub: sub}, nil
}

// WatchSetAvsLogic is a free log subscription operation binding the contract event 0xdf0d3b0bf99a87fc195d045bc7ec61d3e2619e6a49dd3f5cb69102b8c9702034.
//
// Solidity: event SetAvsLogic(address avsLogic)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) WatchSetAvsLogic(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterSetAvsLogic) (event.Subscription, error) {

	logs, sub, err := _ContractAttestationCenter.contract.WatchLogs(opts, "SetAvsLogic")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAttestationCenterSetAvsLogic)
				if err := _ContractAttestationCenter.contract.UnpackLog(event, "SetAvsLogic", log); err != nil {
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

// ParseSetAvsLogic is a log parse operation binding the contract event 0xdf0d3b0bf99a87fc195d045bc7ec61d3e2619e6a49dd3f5cb69102b8c9702034.
//
// Solidity: event SetAvsLogic(address avsLogic)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) ParseSetAvsLogic(log types.Log) (*ContractAttestationCenterSetAvsLogic, error) {
	event := new(ContractAttestationCenterSetAvsLogic)
	if err := _ContractAttestationCenter.contract.UnpackLog(event, "SetAvsLogic", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAttestationCenterSetBeforePaymentsLogicIterator is returned from FilterSetBeforePaymentsLogic and is used to iterate over the raw logs and unpacked data for SetBeforePaymentsLogic events raised by the ContractAttestationCenter contract.
type ContractAttestationCenterSetBeforePaymentsLogicIterator struct {
	Event *ContractAttestationCenterSetBeforePaymentsLogic // Event containing the contract specifics and raw log

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
func (it *ContractAttestationCenterSetBeforePaymentsLogicIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAttestationCenterSetBeforePaymentsLogic)
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
		it.Event = new(ContractAttestationCenterSetBeforePaymentsLogic)
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
func (it *ContractAttestationCenterSetBeforePaymentsLogicIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAttestationCenterSetBeforePaymentsLogicIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAttestationCenterSetBeforePaymentsLogic represents a SetBeforePaymentsLogic event raised by the ContractAttestationCenter contract.
type ContractAttestationCenterSetBeforePaymentsLogic struct {
	PaymentsLogic common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetBeforePaymentsLogic is a free log retrieval operation binding the contract event 0x6da780d66fa2f1ae3eb780c2f39d31bec5c71c81a572640a6af0e8a443477792.
//
// Solidity: event SetBeforePaymentsLogic(address paymentsLogic)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) FilterSetBeforePaymentsLogic(opts *bind.FilterOpts) (*ContractAttestationCenterSetBeforePaymentsLogicIterator, error) {

	logs, sub, err := _ContractAttestationCenter.contract.FilterLogs(opts, "SetBeforePaymentsLogic")
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterSetBeforePaymentsLogicIterator{contract: _ContractAttestationCenter.contract, event: "SetBeforePaymentsLogic", logs: logs, sub: sub}, nil
}

// WatchSetBeforePaymentsLogic is a free log subscription operation binding the contract event 0x6da780d66fa2f1ae3eb780c2f39d31bec5c71c81a572640a6af0e8a443477792.
//
// Solidity: event SetBeforePaymentsLogic(address paymentsLogic)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) WatchSetBeforePaymentsLogic(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterSetBeforePaymentsLogic) (event.Subscription, error) {

	logs, sub, err := _ContractAttestationCenter.contract.WatchLogs(opts, "SetBeforePaymentsLogic")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAttestationCenterSetBeforePaymentsLogic)
				if err := _ContractAttestationCenter.contract.UnpackLog(event, "SetBeforePaymentsLogic", log); err != nil {
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

// ParseSetBeforePaymentsLogic is a log parse operation binding the contract event 0x6da780d66fa2f1ae3eb780c2f39d31bec5c71c81a572640a6af0e8a443477792.
//
// Solidity: event SetBeforePaymentsLogic(address paymentsLogic)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) ParseSetBeforePaymentsLogic(log types.Log) (*ContractAttestationCenterSetBeforePaymentsLogic, error) {
	event := new(ContractAttestationCenterSetBeforePaymentsLogic)
	if err := _ContractAttestationCenter.contract.UnpackLog(event, "SetBeforePaymentsLogic", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAttestationCenterSetFeeCalculatorIterator is returned from FilterSetFeeCalculator and is used to iterate over the raw logs and unpacked data for SetFeeCalculator events raised by the ContractAttestationCenter contract.
type ContractAttestationCenterSetFeeCalculatorIterator struct {
	Event *ContractAttestationCenterSetFeeCalculator // Event containing the contract specifics and raw log

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
func (it *ContractAttestationCenterSetFeeCalculatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAttestationCenterSetFeeCalculator)
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
		it.Event = new(ContractAttestationCenterSetFeeCalculator)
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
func (it *ContractAttestationCenterSetFeeCalculatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAttestationCenterSetFeeCalculatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAttestationCenterSetFeeCalculator represents a SetFeeCalculator event raised by the ContractAttestationCenter contract.
type ContractAttestationCenterSetFeeCalculator struct {
	FeeCalculator common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetFeeCalculator is a free log retrieval operation binding the contract event 0x83b9ee7f260088fdd4ee12aa07fa7daebc115d796b6bfb55bfb0fc839bccff2d.
//
// Solidity: event SetFeeCalculator(address feeCalculator)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) FilterSetFeeCalculator(opts *bind.FilterOpts) (*ContractAttestationCenterSetFeeCalculatorIterator, error) {

	logs, sub, err := _ContractAttestationCenter.contract.FilterLogs(opts, "SetFeeCalculator")
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterSetFeeCalculatorIterator{contract: _ContractAttestationCenter.contract, event: "SetFeeCalculator", logs: logs, sub: sub}, nil
}

// WatchSetFeeCalculator is a free log subscription operation binding the contract event 0x83b9ee7f260088fdd4ee12aa07fa7daebc115d796b6bfb55bfb0fc839bccff2d.
//
// Solidity: event SetFeeCalculator(address feeCalculator)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) WatchSetFeeCalculator(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterSetFeeCalculator) (event.Subscription, error) {

	logs, sub, err := _ContractAttestationCenter.contract.WatchLogs(opts, "SetFeeCalculator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAttestationCenterSetFeeCalculator)
				if err := _ContractAttestationCenter.contract.UnpackLog(event, "SetFeeCalculator", log); err != nil {
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

// ParseSetFeeCalculator is a log parse operation binding the contract event 0x83b9ee7f260088fdd4ee12aa07fa7daebc115d796b6bfb55bfb0fc839bccff2d.
//
// Solidity: event SetFeeCalculator(address feeCalculator)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) ParseSetFeeCalculator(log types.Log) (*ContractAttestationCenterSetFeeCalculator, error) {
	event := new(ContractAttestationCenterSetFeeCalculator)
	if err := _ContractAttestationCenter.contract.UnpackLog(event, "SetFeeCalculator", log); err != nil {
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

// ContractAttestationCenterSetMinimumTaskDefinitionVotingPowerIterator is returned from FilterSetMinimumTaskDefinitionVotingPower and is used to iterate over the raw logs and unpacked data for SetMinimumTaskDefinitionVotingPower events raised by the ContractAttestationCenter contract.
type ContractAttestationCenterSetMinimumTaskDefinitionVotingPowerIterator struct {
	Event *ContractAttestationCenterSetMinimumTaskDefinitionVotingPower // Event containing the contract specifics and raw log

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
func (it *ContractAttestationCenterSetMinimumTaskDefinitionVotingPowerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAttestationCenterSetMinimumTaskDefinitionVotingPower)
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
		it.Event = new(ContractAttestationCenterSetMinimumTaskDefinitionVotingPower)
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
func (it *ContractAttestationCenterSetMinimumTaskDefinitionVotingPowerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAttestationCenterSetMinimumTaskDefinitionVotingPowerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAttestationCenterSetMinimumTaskDefinitionVotingPower represents a SetMinimumTaskDefinitionVotingPower event raised by the ContractAttestationCenter contract.
type ContractAttestationCenterSetMinimumTaskDefinitionVotingPower struct {
	MinimumVotingPower *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSetMinimumTaskDefinitionVotingPower is a free log retrieval operation binding the contract event 0x255c174d5deb340ac0a0c908147d9c66ae7d94a6c7f969f722bf5d71b92e98f8.
//
// Solidity: event SetMinimumTaskDefinitionVotingPower(uint256 minimumVotingPower)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) FilterSetMinimumTaskDefinitionVotingPower(opts *bind.FilterOpts) (*ContractAttestationCenterSetMinimumTaskDefinitionVotingPowerIterator, error) {

	logs, sub, err := _ContractAttestationCenter.contract.FilterLogs(opts, "SetMinimumTaskDefinitionVotingPower")
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterSetMinimumTaskDefinitionVotingPowerIterator{contract: _ContractAttestationCenter.contract, event: "SetMinimumTaskDefinitionVotingPower", logs: logs, sub: sub}, nil
}

// WatchSetMinimumTaskDefinitionVotingPower is a free log subscription operation binding the contract event 0x255c174d5deb340ac0a0c908147d9c66ae7d94a6c7f969f722bf5d71b92e98f8.
//
// Solidity: event SetMinimumTaskDefinitionVotingPower(uint256 minimumVotingPower)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) WatchSetMinimumTaskDefinitionVotingPower(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterSetMinimumTaskDefinitionVotingPower) (event.Subscription, error) {

	logs, sub, err := _ContractAttestationCenter.contract.WatchLogs(opts, "SetMinimumTaskDefinitionVotingPower")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAttestationCenterSetMinimumTaskDefinitionVotingPower)
				if err := _ContractAttestationCenter.contract.UnpackLog(event, "SetMinimumTaskDefinitionVotingPower", log); err != nil {
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

// ParseSetMinimumTaskDefinitionVotingPower is a log parse operation binding the contract event 0x255c174d5deb340ac0a0c908147d9c66ae7d94a6c7f969f722bf5d71b92e98f8.
//
// Solidity: event SetMinimumTaskDefinitionVotingPower(uint256 minimumVotingPower)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) ParseSetMinimumTaskDefinitionVotingPower(log types.Log) (*ContractAttestationCenterSetMinimumTaskDefinitionVotingPower, error) {
	event := new(ContractAttestationCenterSetMinimumTaskDefinitionVotingPower)
	if err := _ContractAttestationCenter.contract.UnpackLog(event, "SetMinimumTaskDefinitionVotingPower", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAttestationCenterSetRestrictedOperatorIterator is returned from FilterSetRestrictedOperator and is used to iterate over the raw logs and unpacked data for SetRestrictedOperator events raised by the ContractAttestationCenter contract.
type ContractAttestationCenterSetRestrictedOperatorIterator struct {
	Event *ContractAttestationCenterSetRestrictedOperator // Event containing the contract specifics and raw log

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
func (it *ContractAttestationCenterSetRestrictedOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAttestationCenterSetRestrictedOperator)
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
		it.Event = new(ContractAttestationCenterSetRestrictedOperator)
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
func (it *ContractAttestationCenterSetRestrictedOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAttestationCenterSetRestrictedOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAttestationCenterSetRestrictedOperator represents a SetRestrictedOperator event raised by the ContractAttestationCenter contract.
type ContractAttestationCenterSetRestrictedOperator struct {
	TaskDefinitionId          uint16
	RestrictedOperatorIndexes []*big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterSetRestrictedOperator is a free log retrieval operation binding the contract event 0x364aa2fa0cf7a32b9240f9ab2bdebc99f0222262852b3b25a87388acff5a5b14.
//
// Solidity: event SetRestrictedOperator(uint16 taskDefinitionId, uint256[] restrictedOperatorIndexes)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) FilterSetRestrictedOperator(opts *bind.FilterOpts) (*ContractAttestationCenterSetRestrictedOperatorIterator, error) {

	logs, sub, err := _ContractAttestationCenter.contract.FilterLogs(opts, "SetRestrictedOperator")
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterSetRestrictedOperatorIterator{contract: _ContractAttestationCenter.contract, event: "SetRestrictedOperator", logs: logs, sub: sub}, nil
}

// WatchSetRestrictedOperator is a free log subscription operation binding the contract event 0x364aa2fa0cf7a32b9240f9ab2bdebc99f0222262852b3b25a87388acff5a5b14.
//
// Solidity: event SetRestrictedOperator(uint16 taskDefinitionId, uint256[] restrictedOperatorIndexes)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) WatchSetRestrictedOperator(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterSetRestrictedOperator) (event.Subscription, error) {

	logs, sub, err := _ContractAttestationCenter.contract.WatchLogs(opts, "SetRestrictedOperator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAttestationCenterSetRestrictedOperator)
				if err := _ContractAttestationCenter.contract.UnpackLog(event, "SetRestrictedOperator", log); err != nil {
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

// ParseSetRestrictedOperator is a log parse operation binding the contract event 0x364aa2fa0cf7a32b9240f9ab2bdebc99f0222262852b3b25a87388acff5a5b14.
//
// Solidity: event SetRestrictedOperator(uint16 taskDefinitionId, uint256[] restrictedOperatorIndexes)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) ParseSetRestrictedOperator(log types.Log) (*ContractAttestationCenterSetRestrictedOperator, error) {
	event := new(ContractAttestationCenterSetRestrictedOperator)
	if err := _ContractAttestationCenter.contract.UnpackLog(event, "SetRestrictedOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAttestationCenterTaskDefinitionCreatedIterator is returned from FilterTaskDefinitionCreated and is used to iterate over the raw logs and unpacked data for TaskDefinitionCreated events raised by the ContractAttestationCenter contract.
type ContractAttestationCenterTaskDefinitionCreatedIterator struct {
	Event *ContractAttestationCenterTaskDefinitionCreated // Event containing the contract specifics and raw log

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
func (it *ContractAttestationCenterTaskDefinitionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAttestationCenterTaskDefinitionCreated)
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
		it.Event = new(ContractAttestationCenterTaskDefinitionCreated)
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
func (it *ContractAttestationCenterTaskDefinitionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAttestationCenterTaskDefinitionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAttestationCenterTaskDefinitionCreated represents a TaskDefinitionCreated event raised by the ContractAttestationCenter contract.
type ContractAttestationCenterTaskDefinitionCreated struct {
	TaskDefinitionId           uint16
	Name                       string
	BlockExpiry                *big.Int
	BaseRewardFeeForAttesters  *big.Int
	BaseRewardFeeForPerformer  *big.Int
	BaseRewardFeeForAggregator *big.Int
	DisputePeriodBlocks        *big.Int
	MinimumVotingPower         *big.Int
	RestrictedOperatorIndexes  []*big.Int
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterTaskDefinitionCreated is a free log retrieval operation binding the contract event 0x4306a9df64b49fc07c0f1929d57cc2b5cdfc108656189460e6aa127754413ef1.
//
// Solidity: event TaskDefinitionCreated(uint16 taskDefinitionId, string name, uint256 blockExpiry, uint256 baseRewardFeeForAttesters, uint256 baseRewardFeeForPerformer, uint256 baseRewardFeeForAggregator, uint256 disputePeriodBlocks, uint256 minimumVotingPower, uint256[] restrictedOperatorIndexes)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) FilterTaskDefinitionCreated(opts *bind.FilterOpts) (*ContractAttestationCenterTaskDefinitionCreatedIterator, error) {

	logs, sub, err := _ContractAttestationCenter.contract.FilterLogs(opts, "TaskDefinitionCreated")
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterTaskDefinitionCreatedIterator{contract: _ContractAttestationCenter.contract, event: "TaskDefinitionCreated", logs: logs, sub: sub}, nil
}

// WatchTaskDefinitionCreated is a free log subscription operation binding the contract event 0x4306a9df64b49fc07c0f1929d57cc2b5cdfc108656189460e6aa127754413ef1.
//
// Solidity: event TaskDefinitionCreated(uint16 taskDefinitionId, string name, uint256 blockExpiry, uint256 baseRewardFeeForAttesters, uint256 baseRewardFeeForPerformer, uint256 baseRewardFeeForAggregator, uint256 disputePeriodBlocks, uint256 minimumVotingPower, uint256[] restrictedOperatorIndexes)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) WatchTaskDefinitionCreated(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterTaskDefinitionCreated) (event.Subscription, error) {

	logs, sub, err := _ContractAttestationCenter.contract.WatchLogs(opts, "TaskDefinitionCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAttestationCenterTaskDefinitionCreated)
				if err := _ContractAttestationCenter.contract.UnpackLog(event, "TaskDefinitionCreated", log); err != nil {
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

// ParseTaskDefinitionCreated is a log parse operation binding the contract event 0x4306a9df64b49fc07c0f1929d57cc2b5cdfc108656189460e6aa127754413ef1.
//
// Solidity: event TaskDefinitionCreated(uint16 taskDefinitionId, string name, uint256 blockExpiry, uint256 baseRewardFeeForAttesters, uint256 baseRewardFeeForPerformer, uint256 baseRewardFeeForAggregator, uint256 disputePeriodBlocks, uint256 minimumVotingPower, uint256[] restrictedOperatorIndexes)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) ParseTaskDefinitionCreated(log types.Log) (*ContractAttestationCenterTaskDefinitionCreated, error) {
	event := new(ContractAttestationCenterTaskDefinitionCreated)
	if err := _ContractAttestationCenter.contract.UnpackLog(event, "TaskDefinitionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAttestationCenterTaskDefinitionRestrictedOperatorsModifiedIterator is returned from FilterTaskDefinitionRestrictedOperatorsModified and is used to iterate over the raw logs and unpacked data for TaskDefinitionRestrictedOperatorsModified events raised by the ContractAttestationCenter contract.
type ContractAttestationCenterTaskDefinitionRestrictedOperatorsModifiedIterator struct {
	Event *ContractAttestationCenterTaskDefinitionRestrictedOperatorsModified // Event containing the contract specifics and raw log

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
func (it *ContractAttestationCenterTaskDefinitionRestrictedOperatorsModifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAttestationCenterTaskDefinitionRestrictedOperatorsModified)
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
		it.Event = new(ContractAttestationCenterTaskDefinitionRestrictedOperatorsModified)
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
func (it *ContractAttestationCenterTaskDefinitionRestrictedOperatorsModifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAttestationCenterTaskDefinitionRestrictedOperatorsModifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAttestationCenterTaskDefinitionRestrictedOperatorsModified represents a TaskDefinitionRestrictedOperatorsModified event raised by the ContractAttestationCenter contract.
type ContractAttestationCenterTaskDefinitionRestrictedOperatorsModified struct {
	TaskDefinitionId          uint16
	RestrictedOperatorIndexes []*big.Int
	IsRestricted              []bool
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterTaskDefinitionRestrictedOperatorsModified is a free log retrieval operation binding the contract event 0x3a6545f49055b62a6dec3f6d6116dfcde655324bc0c3a4947266f3d11bff8239.
//
// Solidity: event TaskDefinitionRestrictedOperatorsModified(uint16 taskDefinitionId, uint256[] restrictedOperatorIndexes, bool[] isRestricted)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) FilterTaskDefinitionRestrictedOperatorsModified(opts *bind.FilterOpts) (*ContractAttestationCenterTaskDefinitionRestrictedOperatorsModifiedIterator, error) {

	logs, sub, err := _ContractAttestationCenter.contract.FilterLogs(opts, "TaskDefinitionRestrictedOperatorsModified")
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterTaskDefinitionRestrictedOperatorsModifiedIterator{contract: _ContractAttestationCenter.contract, event: "TaskDefinitionRestrictedOperatorsModified", logs: logs, sub: sub}, nil
}

// WatchTaskDefinitionRestrictedOperatorsModified is a free log subscription operation binding the contract event 0x3a6545f49055b62a6dec3f6d6116dfcde655324bc0c3a4947266f3d11bff8239.
//
// Solidity: event TaskDefinitionRestrictedOperatorsModified(uint16 taskDefinitionId, uint256[] restrictedOperatorIndexes, bool[] isRestricted)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) WatchTaskDefinitionRestrictedOperatorsModified(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterTaskDefinitionRestrictedOperatorsModified) (event.Subscription, error) {

	logs, sub, err := _ContractAttestationCenter.contract.WatchLogs(opts, "TaskDefinitionRestrictedOperatorsModified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAttestationCenterTaskDefinitionRestrictedOperatorsModified)
				if err := _ContractAttestationCenter.contract.UnpackLog(event, "TaskDefinitionRestrictedOperatorsModified", log); err != nil {
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

// ParseTaskDefinitionRestrictedOperatorsModified is a log parse operation binding the contract event 0x3a6545f49055b62a6dec3f6d6116dfcde655324bc0c3a4947266f3d11bff8239.
//
// Solidity: event TaskDefinitionRestrictedOperatorsModified(uint16 taskDefinitionId, uint256[] restrictedOperatorIndexes, bool[] isRestricted)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) ParseTaskDefinitionRestrictedOperatorsModified(log types.Log) (*ContractAttestationCenterTaskDefinitionRestrictedOperatorsModified, error) {
	event := new(ContractAttestationCenterTaskDefinitionRestrictedOperatorsModified)
	if err := _ContractAttestationCenter.contract.UnpackLog(event, "TaskDefinitionRestrictedOperatorsModified", log); err != nil {
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
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTaskRejected is a free log retrieval operation binding the contract event 0x5d681de577a7b3ff669022ffe837e270ef5b32ee4c1377045cf61b700d99e70a.
//
// Solidity: event TaskRejected(address operator, uint32 taskNumber, string proofOfTask, bytes data, uint16 taskDefinitionId)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) FilterTaskRejected(opts *bind.FilterOpts) (*ContractAttestationCenterTaskRejectedIterator, error) {

	logs, sub, err := _ContractAttestationCenter.contract.FilterLogs(opts, "TaskRejected")
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterTaskRejectedIterator{contract: _ContractAttestationCenter.contract, event: "TaskRejected", logs: logs, sub: sub}, nil
}

// WatchTaskRejected is a free log subscription operation binding the contract event 0x5d681de577a7b3ff669022ffe837e270ef5b32ee4c1377045cf61b700d99e70a.
//
// Solidity: event TaskRejected(address operator, uint32 taskNumber, string proofOfTask, bytes data, uint16 taskDefinitionId)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) WatchTaskRejected(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterTaskRejected) (event.Subscription, error) {

	logs, sub, err := _ContractAttestationCenter.contract.WatchLogs(opts, "TaskRejected")
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

// ParseTaskRejected is a log parse operation binding the contract event 0x5d681de577a7b3ff669022ffe837e270ef5b32ee4c1377045cf61b700d99e70a.
//
// Solidity: event TaskRejected(address operator, uint32 taskNumber, string proofOfTask, bytes data, uint16 taskDefinitionId)
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
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTaskSubmitted is a free log retrieval operation binding the contract event 0xde1faaa044a7216023ec32c06f91d9b5098d8bbd3b37f3662be3a729752ec9fc.
//
// Solidity: event TaskSubmitted(address operator, uint32 taskNumber, string proofOfTask, bytes data, uint16 taskDefinitionId)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) FilterTaskSubmitted(opts *bind.FilterOpts) (*ContractAttestationCenterTaskSubmittedIterator, error) {

	logs, sub, err := _ContractAttestationCenter.contract.FilterLogs(opts, "TaskSubmitted")
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterTaskSubmittedIterator{contract: _ContractAttestationCenter.contract, event: "TaskSubmitted", logs: logs, sub: sub}, nil
}

// WatchTaskSubmitted is a free log subscription operation binding the contract event 0xde1faaa044a7216023ec32c06f91d9b5098d8bbd3b37f3662be3a729752ec9fc.
//
// Solidity: event TaskSubmitted(address operator, uint32 taskNumber, string proofOfTask, bytes data, uint16 taskDefinitionId)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) WatchTaskSubmitted(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterTaskSubmitted) (event.Subscription, error) {

	logs, sub, err := _ContractAttestationCenter.contract.WatchLogs(opts, "TaskSubmitted")
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

// ParseTaskSubmitted is a log parse operation binding the contract event 0xde1faaa044a7216023ec32c06f91d9b5098d8bbd3b37f3662be3a729752ec9fc.
//
// Solidity: event TaskSubmitted(address operator, uint32 taskNumber, string proofOfTask, bytes data, uint16 taskDefinitionId)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) ParseTaskSubmitted(log types.Log) (*ContractAttestationCenterTaskSubmitted, error) {
	event := new(ContractAttestationCenterTaskSubmitted)
	if err := _ContractAttestationCenter.contract.UnpackLog(event, "TaskSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
