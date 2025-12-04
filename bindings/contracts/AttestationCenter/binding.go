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
	FeeToClaim  *big.Int
	BlsKey      [4]*big.Int
}

// IAttestationCenterPaymentDetails is an auto generated low-level Go binding around an user-defined struct.
type IAttestationCenterPaymentDetails struct {
	Operator           common.Address
	LastPaidTaskNumber *big.Int
	FeeToClaim         *big.Int
	PaymentStatus      uint8
}

// IInternalTaskHandlerInternalTaskConfig is an auto generated low-level Go binding around an user-defined struct.
type IInternalTaskHandlerInternalTaskConfig struct {
	IsInternalTaskActivated bool
	Interval                *big.Int
	LeaderElectionMechanism uint8
	Data                    []byte
}

// IRewardsDistributorPaymentRequestMessage is an auto generated low-level Go binding around an user-defined struct.
type IRewardsDistributorPaymentRequestMessage struct {
	Operator   common.Address
	FeeToClaim *big.Int
}

// ISlashingConfigSlashingStakingContractInfo is an auto generated low-level Go binding around an user-defined struct.
type ISlashingConfigSlashingStakingContractInfo struct {
	StakingContract        common.Address
	SharedSecurityProvider uint8
	WadsToSlash            *big.Int
}

// TaskDefinition is an auto generated low-level Go binding around an user-defined struct.
type TaskDefinition struct {
	TaskDefinitionId                      uint16
	IsRejectedTaskSlashingEnabled         bool
	IsIncorrectAttestationSlashingEnabled bool
	Name                                  string
	BlockExpiry                           *big.Int
	BaseRewardFeeForAttesters             *big.Int
	BaseRewardFeeForPerformer             *big.Int
	BaseRewardFeeForAggregator            *big.Int
	DisputePeriodBlocks                   *big.Int
	MinimumVotingPower                    *big.Int
	RestrictedAttesterIds                 []*big.Int
	MaximumNumberOfAttesters              *big.Int
}

// TaskDefinitionParamsV2 is an auto generated low-level Go binding around an user-defined struct.
type TaskDefinitionParamsV2 struct {
	BlockExpiry                *big.Int
	BaseRewardFeeForAttesters  *big.Int
	BaseRewardFeeForPerformer  *big.Int
	BaseRewardFeeForAggregator *big.Int
	DisputePeriodBlocks        *big.Int
	MinimumVotingPower         *big.Int
	RestrictedAttesterIds      []*big.Int
	MaximumNumberOfAttesters   *big.Int
}

// ContractAttestationCenterMetaData contains all meta data concerning the ContractAttestationCenter contract.
var ContractAttestationCenterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_extensionImplementation\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EigenRewardsDurationExceedsMaximum\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EigenRewardsDurationNotMultipleOfInterval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"totalRewards\",\"type\":\"uint256\"}],\"name\":\"EigenRewardsMaxRewardsAmountExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EigenRewardsMustBeRetroactive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EigenRewardsNotSupportedOnL2\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EigenRewardsStartTimestampNotMultipleOfInterval\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EigenRewardsStartTimestampTooFarInPast\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FlowIsCurrentlyPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FlowIsCurrentlyUnpaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"taskDefinitionId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"minVotingPower\",\"type\":\"uint256\"}],\"name\":\"InsufficientVotingPowerForTaskDefinition\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAttesterSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBlockExpiry\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaximumNumberOfAttesters\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOperatorsForPayment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeForBatchPaymentRequest\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"taskDefinitionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"operatorIndex\",\"type\":\"uint256\"}],\"name\":\"InvalidRestrictedAttester\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRestrictedAttesterIds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTaskDefinition\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWadToSlash\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_roles\",\"type\":\"bytes\"}],\"name\":\"NotAuthorizedRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operatorAddress\",\"type\":\"address\"}],\"name\":\"OperatorNotRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PauseFlowIsAlreadyPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SymbioticRewardsNotSupportedOnL2\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"taskDefinitionId\",\"type\":\"uint16\"}],\"name\":\"TaskDefinitionNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnpausingFlowIsAlreadyUnpaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestedTaskNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestedAmountClaimed\",\"type\":\"uint256\"}],\"name\":\"ClearPaymentRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"startTimestamp\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"duration\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeToClaim\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIRewardsDistributor.PaymentRequestMessage[]\",\"name\":\"operators\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lastPaidTaskNumber\",\"type\":\"uint256\"}],\"name\":\"EigenPaymentsRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"_pausableFlow\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_pauser\",\"type\":\"address\"}],\"name\":\"FlowPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"_pausableFlowFlag\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_unpauser\",\"type\":\"address\"}],\"name\":\"FlowUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isOpenAggregator\",\"type\":\"bool\"}],\"name\":\"IsOpenAggregatorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorEjectionRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"stakingContract\",\"type\":\"address\"},{\"internalType\":\"enumIAvsGovernance.SharedSecurityProvider\",\"name\":\"sharedSecurityProvider\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"wadsToSlash\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structISlashingConfig.SlashingStakingContractInfo[]\",\"name\":\"slashingStakingContractInfos\",\"type\":\"tuple[]\"}],\"name\":\"OperatorSlashRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeToClaim\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIRewardsDistributor.PaymentRequestMessage[]\",\"name\":\"operators\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lastPaidTaskNumber\",\"type\":\"uint256\"}],\"name\":\"PaymentsRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAvsGovernanceMultisig\",\"type\":\"address\"}],\"name\":\"SetAvsGovernanceMultisig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"avsLogic\",\"type\":\"address\"}],\"name\":\"SetAvsLogic\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"taskDefinitionId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseRewardFeeForAggregator\",\"type\":\"uint256\"}],\"name\":\"SetBaseRewardFeeForAggregator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"taskDefinitionId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseRewardFeeForAttesters\",\"type\":\"uint256\"}],\"name\":\"SetBaseRewardFeeForAttesters\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"taskDefinitionId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseRewardFeeForPerformer\",\"type\":\"uint256\"}],\"name\":\"SetBaseRewardFeeForPerformer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"paymentsLogic\",\"type\":\"address\"}],\"name\":\"SetBeforePaymentsLogic\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengerRewardFee\",\"type\":\"uint256\"}],\"name\":\"SetChallengerRewardFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeCalculator\",\"type\":\"address\"}],\"name\":\"SetFeeCalculator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newInternalTaskHandler\",\"type\":\"address\"}],\"name\":\"SetInternalTaskHandler\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"taskDefinitionId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isIncorrectAttestationSlashingEnabled\",\"type\":\"bool\"}],\"name\":\"SetIsIncorrectAttestationSlashingEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"taskDefinitionId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isRejectedTaskSlashingEnabled\",\"type\":\"bool\"}],\"name\":\"SetIsRejectedTaskSlashingEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"taskDefinitionId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maximumNumberOfAttesters\",\"type\":\"uint256\"}],\"name\":\"SetMaximumNumberOfAttesters\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newMessageHandler\",\"type\":\"address\"}],\"name\":\"SetMessageHandler\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumVotingPower\",\"type\":\"uint256\"}],\"name\":\"SetMinimumTaskDefinitionVotingPower\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"taskDefinitionId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"restrictedAttesterIds\",\"type\":\"uint256[]\"}],\"name\":\"SetRestrictedAttester\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"taskDefinitionId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockExpiry\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseRewardFeeForAttesters\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseRewardFeeForPerformer\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseRewardFeeForAggregator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"disputePeriodBlocks\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumVotingPower\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"restrictedAttesterIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maximumNumberOfAttesters\",\"type\":\"uint256\"}],\"name\":\"TaskDefinitionCreated\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXTENSION_IMPLEMENTATION\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"stakingContract\",\"type\":\"address\"},{\"internalType\":\"enumIAvsGovernance.SharedSecurityProvider\",\"name\":\"sharedSecurityProvider\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"wadsToSlash\",\"type\":\"uint256\"}],\"internalType\":\"structISlashingConfig.SlashingStakingContractInfo[]\",\"name\":\"_slashingStakingContractInfos\",\"type\":\"tuple[]\"}],\"name\":\"applyCustomSlashing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"avsLogic\",\"outputs\":[{\"internalType\":\"contractIAvsLogic\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"avsTreasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beforePaymentsLogic\",\"outputs\":[{\"internalType\":\"contractIBeforePaymentsLogic\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"blockExpiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseRewardFeeForAttesters\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseRewardFeeForPerformer\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseRewardFeeForAggregator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"disputePeriodBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumVotingPower\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"restrictedAttesterIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"maximumNumberOfAttesters\",\"type\":\"uint256\"}],\"internalType\":\"structTaskDefinitionParamsV2\",\"name\":\"_taskDefinitionParams\",\"type\":\"tuple\"}],\"name\":\"createNewTaskDefinition\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"_id\",\"type\":\"uint16\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"ejectOperatorFromNetwork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeCalculator\",\"outputs\":[{\"internalType\":\"contractIFeeCalculator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveOperatorsDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votingPower\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeToClaim\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"blsKey\",\"type\":\"uint256[4]\"}],\"internalType\":\"structIAttestationCenter.OperatorDetails[]\",\"name\":\"_operators\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getL2MessageHandler\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"getOperatorPaymentDetail\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lastPaidTaskNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeToClaim\",\"type\":\"uint256\"},{\"internalType\":\"enumIAttestationCenter.PaymentStatus\",\"name\":\"paymentStatus\",\"type\":\"uint8\"}],\"internalType\":\"structIAttestationCenter.PaymentDetails\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_taskDefinitionId\",\"type\":\"uint16\"}],\"name\":\"getTaskDefinitionDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"taskDefinitionId\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"isRejectedTaskSlashingEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isIncorrectAttestationSlashingEnabled\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"blockExpiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseRewardFeeForAttesters\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseRewardFeeForPerformer\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseRewardFeeForAggregator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"disputePeriodBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumVotingPower\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"restrictedAttesterIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"maximumNumberOfAttesters\",\"type\":\"uint256\"}],\"internalType\":\"structTaskDefinition\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"avsGovernanceMultisigOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operationsMultisig\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"communityMultisig\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"messageHandler\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"obls\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"avsTreasury\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isRewardsOnL2\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"internalTaskHandler\",\"type\":\"address\"}],\"internalType\":\"structIAttestationCenter.InitializationParams\",\"name\":\"_initializationParams\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"internalTaskHandler\",\"outputs\":[{\"internalType\":\"contractIInternalTaskHandler\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_pausableFlow\",\"type\":\"bytes4\"}],\"name\":\"isFlowPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_isPaused\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_avsGovernanceMultisigOwner\",\"type\":\"address\"}],\"name\":\"migrateSlashingFlow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextEigenRewardsBatchStartTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numOfActiveOperators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numOfTaskDefinitions\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numOfTotalOperators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"obls\",\"outputs\":[{\"internalType\":\"contractIOBLS\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"operatorsIdsByAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_pausableFlow\",\"type\":\"bytes4\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paymentRequestsRole\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_currentPaymentRequestsAdmin\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_to\",\"type\":\"uint256\"},{\"internalType\":\"enumIRewardsDistributor.SubmissionType\",\"name\":\"_submissionType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_distributionData\",\"type\":\"bytes\"}],\"name\":\"requestBatchPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIRewardsDistributor.SubmissionType\",\"name\":\"_submissionType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_distributionData\",\"type\":\"bytes\"}],\"name\":\"requestBatchPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_startTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_duration\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_to\",\"type\":\"uint256\"}],\"name\":\"requestEigenBatchPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ejector\",\"type\":\"address\"}],\"name\":\"revokeEjector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_slasher\",\"type\":\"address\"}],\"name\":\"revokeSlasher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAvsLogic\",\"name\":\"_avsLogic\",\"type\":\"address\"}],\"name\":\"setAvsLogic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_taskDefinitionId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_baseRewardFeeForAggregator\",\"type\":\"uint256\"}],\"name\":\"setBaseRewardFeeForAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_taskDefinitionId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_baseRewardFeeForAttesters\",\"type\":\"uint256\"}],\"name\":\"setBaseRewardFeeForAttesters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_taskDefinitionId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_baseRewardFeeForPerformer\",\"type\":\"uint256\"}],\"name\":\"setBaseRewardFeeForPerformer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBeforePaymentsLogic\",\"name\":\"_beforePaymentsLogic\",\"type\":\"address\"}],\"name\":\"setBeforePaymentsLogic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_challengerRewardFee\",\"type\":\"uint256\"}],\"name\":\"setChallengerRewardFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ejector\",\"type\":\"address\"}],\"name\":\"setEjector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIFeeCalculator\",\"name\":\"_feeCalculator\",\"type\":\"address\"}],\"name\":\"setFeeCalculator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIInternalTaskHandler\",\"name\":\"_internalTaskHandler\",\"type\":\"address\"}],\"name\":\"setInternalTaskHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_taskDefinitionId\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"_isIncorrectAttestationSlashingEnabled\",\"type\":\"bool\"}],\"name\":\"setIsIncorrectAttestationSlashingEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isOpenAggregator\",\"type\":\"bool\"}],\"name\":\"setIsOpenAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_taskDefinitionId\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"_isRejectedTaskSlashingEnabled\",\"type\":\"bool\"}],\"name\":\"setIsRejectedTaskSlashingEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oblsSharesSyncer\",\"type\":\"address\"}],\"name\":\"setOblsSharesSyncer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newPaymentRequestsAdmin\",\"type\":\"address\"}],\"name\":\"setPaymentRequestsRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_slasher\",\"type\":\"address\"}],\"name\":\"setSlasher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_taskDefinitionId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_maximumNumberOfAttesters\",\"type\":\"uint256\"}],\"name\":\"setTaskDefinitionMaximumNumberOfAttesters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_taskDefinitionId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_minimumVotingPower\",\"type\":\"uint256\"}],\"name\":\"setTaskDefinitionMinVotingPower\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_taskDefinitionId\",\"type\":\"uint16\"},{\"internalType\":\"uint256[]\",\"name\":\"_restrictedAttesterIds\",\"type\":\"uint256[]\"}],\"name\":\"setTaskDefinitionRestrictedAttesters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taskNumber\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAvsGovernanceMultisig\",\"type\":\"address\"}],\"name\":\"transferAvsGovernanceMultisig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newMessageHandler\",\"type\":\"address\"}],\"name\":\"transferMessageHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_pausableFlow\",\"type\":\"bytes4\"}],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_taskDefinitionId\",\"type\":\"uint16\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isInternalTaskActivated\",\"type\":\"bool\"},{\"internalType\":\"uint48\",\"name\":\"interval\",\"type\":\"uint48\"},{\"internalType\":\"enumIInternalTaskHandler.LeaderElectionMechanism\",\"name\":\"leaderElectionMechanism\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIInternalTaskHandler.InternalTaskConfig\",\"name\":\"_config\",\"type\":\"tuple\"}],\"name\":\"updateInternalTaskConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_taskDefinitionId\",\"type\":\"uint16\"}],\"name\":\"verifyOperatorValidForTaskDefinition\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"votingPower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ContractAttestationCenterABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractAttestationCenterMetaData.ABI instead.
var ContractAttestationCenterABI = ContractAttestationCenterMetaData.ABI

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

	FeeCalculator(opts *bind.CallOpts) (common.Address, error)

	GetActiveOperatorsDetails(opts *bind.CallOpts) ([]IAttestationCenterOperatorDetails, error)

	GetL2MessageHandler(opts *bind.CallOpts) (common.Address, error)

	GetOperatorPaymentDetail(opts *bind.CallOpts, _operatorId *big.Int) (IAttestationCenterPaymentDetails, error)

	GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error)

	GetTaskDefinitionDetails(opts *bind.CallOpts, _taskDefinitionId uint16) (TaskDefinition, error)

	HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error)

	InternalTaskHandler(opts *bind.CallOpts) (common.Address, error)

	IsFlowPaused(opts *bind.CallOpts, _pausableFlow [4]byte) (bool, error)

	NextEigenRewardsBatchStartTimestamp(opts *bind.CallOpts) (*big.Int, error)

	NumOfActiveOperators(opts *bind.CallOpts) (*big.Int, error)

	NumOfTaskDefinitions(opts *bind.CallOpts) (uint16, error)

	NumOfTotalOperators(opts *bind.CallOpts) (*big.Int, error)

	Obls(opts *bind.CallOpts) (common.Address, error)

	OperatorsIdsByAddress(opts *bind.CallOpts, _operator common.Address) (*big.Int, error)

	PaymentRequestsRole(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TaskNumber(opts *bind.CallOpts) (uint32, error)

	VerifyOperatorValidForTaskDefinition(opts *bind.CallOpts, _operator common.Address, _taskDefinitionId uint16) error

	VotingPower(opts *bind.CallOpts, _operator common.Address) (*big.Int, error)
}

// ContractAttestationCenterTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractAttestationCenterTransacts interface {
	ApplyCustomSlashing(opts *bind.TransactOpts, _operator common.Address, _slashingStakingContractInfos []ISlashingConfigSlashingStakingContractInfo) (*types.Transaction, error)

	CreateNewTaskDefinition(opts *bind.TransactOpts, _name string, _taskDefinitionParams TaskDefinitionParamsV2) (*types.Transaction, error)

	EjectOperatorFromNetwork(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error)

	GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts, _initializationParams IAttestationCenterInitializationParams) (*types.Transaction, error)

	MigrateSlashingFlow(opts *bind.TransactOpts, _avsGovernanceMultisigOwner common.Address) (*types.Transaction, error)

	Migration(opts *bind.TransactOpts) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts, _pausableFlow [4]byte) (*types.Transaction, error)

	RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error)

	RequestBatchPayment(opts *bind.TransactOpts, _from *big.Int, _to *big.Int, _submissionType uint8, _distributionData []byte) (*types.Transaction, error)

	RequestBatchPayment0(opts *bind.TransactOpts, _submissionType uint8, _distributionData []byte) (*types.Transaction, error)

	RequestEigenBatchPayment(opts *bind.TransactOpts, _startTimestamp uint32, _duration uint32, _from *big.Int, _to *big.Int) (*types.Transaction, error)

	RevokeEjector(opts *bind.TransactOpts, _ejector common.Address) (*types.Transaction, error)

	RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	RevokeSlasher(opts *bind.TransactOpts, _slasher common.Address) (*types.Transaction, error)

	SetAvsLogic(opts *bind.TransactOpts, _avsLogic common.Address) (*types.Transaction, error)

	SetBaseRewardFeeForAggregator(opts *bind.TransactOpts, _taskDefinitionId uint16, _baseRewardFeeForAggregator *big.Int) (*types.Transaction, error)

	SetBaseRewardFeeForAttesters(opts *bind.TransactOpts, _taskDefinitionId uint16, _baseRewardFeeForAttesters *big.Int) (*types.Transaction, error)

	SetBaseRewardFeeForPerformer(opts *bind.TransactOpts, _taskDefinitionId uint16, _baseRewardFeeForPerformer *big.Int) (*types.Transaction, error)

	SetBeforePaymentsLogic(opts *bind.TransactOpts, _beforePaymentsLogic common.Address) (*types.Transaction, error)

	SetChallengerRewardFee(opts *bind.TransactOpts, _challengerRewardFee *big.Int) (*types.Transaction, error)

	SetEjector(opts *bind.TransactOpts, _ejector common.Address) (*types.Transaction, error)

	SetFeeCalculator(opts *bind.TransactOpts, _feeCalculator common.Address) (*types.Transaction, error)

	SetInternalTaskHandler(opts *bind.TransactOpts, _internalTaskHandler common.Address) (*types.Transaction, error)

	SetIsIncorrectAttestationSlashingEnabled(opts *bind.TransactOpts, _taskDefinitionId uint16, _isIncorrectAttestationSlashingEnabled bool) (*types.Transaction, error)

	SetIsOpenAggregator(opts *bind.TransactOpts, _isOpenAggregator bool) (*types.Transaction, error)

	SetIsRejectedTaskSlashingEnabled(opts *bind.TransactOpts, _taskDefinitionId uint16, _isRejectedTaskSlashingEnabled bool) (*types.Transaction, error)

	SetOblsSharesSyncer(opts *bind.TransactOpts, _oblsSharesSyncer common.Address) (*types.Transaction, error)

	SetPaymentRequestsRole(opts *bind.TransactOpts, _newPaymentRequestsAdmin common.Address) (*types.Transaction, error)

	SetSlasher(opts *bind.TransactOpts, _slasher common.Address) (*types.Transaction, error)

	SetTaskDefinitionMaximumNumberOfAttesters(opts *bind.TransactOpts, _taskDefinitionId uint16, _maximumNumberOfAttesters *big.Int) (*types.Transaction, error)

	SetTaskDefinitionMinVotingPower(opts *bind.TransactOpts, _taskDefinitionId uint16, _minimumVotingPower *big.Int) (*types.Transaction, error)

	SetTaskDefinitionRestrictedAttesters(opts *bind.TransactOpts, _taskDefinitionId uint16, _restrictedAttesterIds []*big.Int) (*types.Transaction, error)

	TransferAvsGovernanceMultisig(opts *bind.TransactOpts, _newAvsGovernanceMultisig common.Address) (*types.Transaction, error)

	TransferMessageHandler(opts *bind.TransactOpts, _newMessageHandler common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts, _pausableFlow [4]byte) (*types.Transaction, error)

	UpdateInternalTaskConfig(opts *bind.TransactOpts, _taskDefinitionId uint16, _config IInternalTaskHandlerInternalTaskConfig) (*types.Transaction, error)
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

	FilterIsOpenAggregatorSet(opts *bind.FilterOpts) (*ContractAttestationCenterIsOpenAggregatorSetIterator, error)
	WatchIsOpenAggregatorSet(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterIsOpenAggregatorSet) (event.Subscription, error)
	ParseIsOpenAggregatorSet(log types.Log) (*ContractAttestationCenterIsOpenAggregatorSet, error)

	FilterOperatorEjectionRequested(opts *bind.FilterOpts) (*ContractAttestationCenterOperatorEjectionRequestedIterator, error)
	WatchOperatorEjectionRequested(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterOperatorEjectionRequested) (event.Subscription, error)
	ParseOperatorEjectionRequested(log types.Log) (*ContractAttestationCenterOperatorEjectionRequested, error)

	FilterOperatorSlashRequested(opts *bind.FilterOpts) (*ContractAttestationCenterOperatorSlashRequestedIterator, error)
	WatchOperatorSlashRequested(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterOperatorSlashRequested) (event.Subscription, error)
	ParseOperatorSlashRequested(log types.Log) (*ContractAttestationCenterOperatorSlashRequested, error)

	FilterPaymentsRequested(opts *bind.FilterOpts) (*ContractAttestationCenterPaymentsRequestedIterator, error)
	WatchPaymentsRequested(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterPaymentsRequested) (event.Subscription, error)
	ParsePaymentsRequested(log types.Log) (*ContractAttestationCenterPaymentsRequested, error)

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

	FilterSetBaseRewardFeeForAggregator(opts *bind.FilterOpts, taskDefinitionId []uint16) (*ContractAttestationCenterSetBaseRewardFeeForAggregatorIterator, error)
	WatchSetBaseRewardFeeForAggregator(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterSetBaseRewardFeeForAggregator, taskDefinitionId []uint16) (event.Subscription, error)
	ParseSetBaseRewardFeeForAggregator(log types.Log) (*ContractAttestationCenterSetBaseRewardFeeForAggregator, error)

	FilterSetBaseRewardFeeForAttesters(opts *bind.FilterOpts, taskDefinitionId []uint16) (*ContractAttestationCenterSetBaseRewardFeeForAttestersIterator, error)
	WatchSetBaseRewardFeeForAttesters(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterSetBaseRewardFeeForAttesters, taskDefinitionId []uint16) (event.Subscription, error)
	ParseSetBaseRewardFeeForAttesters(log types.Log) (*ContractAttestationCenterSetBaseRewardFeeForAttesters, error)

	FilterSetBaseRewardFeeForPerformer(opts *bind.FilterOpts, taskDefinitionId []uint16) (*ContractAttestationCenterSetBaseRewardFeeForPerformerIterator, error)
	WatchSetBaseRewardFeeForPerformer(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterSetBaseRewardFeeForPerformer, taskDefinitionId []uint16) (event.Subscription, error)
	ParseSetBaseRewardFeeForPerformer(log types.Log) (*ContractAttestationCenterSetBaseRewardFeeForPerformer, error)

	FilterSetBeforePaymentsLogic(opts *bind.FilterOpts) (*ContractAttestationCenterSetBeforePaymentsLogicIterator, error)
	WatchSetBeforePaymentsLogic(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterSetBeforePaymentsLogic) (event.Subscription, error)
	ParseSetBeforePaymentsLogic(log types.Log) (*ContractAttestationCenterSetBeforePaymentsLogic, error)

	FilterSetChallengerRewardFee(opts *bind.FilterOpts) (*ContractAttestationCenterSetChallengerRewardFeeIterator, error)
	WatchSetChallengerRewardFee(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterSetChallengerRewardFee) (event.Subscription, error)
	ParseSetChallengerRewardFee(log types.Log) (*ContractAttestationCenterSetChallengerRewardFee, error)

	FilterSetFeeCalculator(opts *bind.FilterOpts) (*ContractAttestationCenterSetFeeCalculatorIterator, error)
	WatchSetFeeCalculator(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterSetFeeCalculator) (event.Subscription, error)
	ParseSetFeeCalculator(log types.Log) (*ContractAttestationCenterSetFeeCalculator, error)

	FilterSetInternalTaskHandler(opts *bind.FilterOpts) (*ContractAttestationCenterSetInternalTaskHandlerIterator, error)
	WatchSetInternalTaskHandler(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterSetInternalTaskHandler) (event.Subscription, error)
	ParseSetInternalTaskHandler(log types.Log) (*ContractAttestationCenterSetInternalTaskHandler, error)

	FilterSetIsIncorrectAttestationSlashingEnabled(opts *bind.FilterOpts, taskDefinitionId []uint16) (*ContractAttestationCenterSetIsIncorrectAttestationSlashingEnabledIterator, error)
	WatchSetIsIncorrectAttestationSlashingEnabled(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterSetIsIncorrectAttestationSlashingEnabled, taskDefinitionId []uint16) (event.Subscription, error)
	ParseSetIsIncorrectAttestationSlashingEnabled(log types.Log) (*ContractAttestationCenterSetIsIncorrectAttestationSlashingEnabled, error)

	FilterSetIsRejectedTaskSlashingEnabled(opts *bind.FilterOpts, taskDefinitionId []uint16) (*ContractAttestationCenterSetIsRejectedTaskSlashingEnabledIterator, error)
	WatchSetIsRejectedTaskSlashingEnabled(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterSetIsRejectedTaskSlashingEnabled, taskDefinitionId []uint16) (event.Subscription, error)
	ParseSetIsRejectedTaskSlashingEnabled(log types.Log) (*ContractAttestationCenterSetIsRejectedTaskSlashingEnabled, error)

	FilterSetMaximumNumberOfAttesters(opts *bind.FilterOpts, taskDefinitionId []uint16) (*ContractAttestationCenterSetMaximumNumberOfAttestersIterator, error)
	WatchSetMaximumNumberOfAttesters(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterSetMaximumNumberOfAttesters, taskDefinitionId []uint16) (event.Subscription, error)
	ParseSetMaximumNumberOfAttesters(log types.Log) (*ContractAttestationCenterSetMaximumNumberOfAttesters, error)

	FilterSetMessageHandler(opts *bind.FilterOpts) (*ContractAttestationCenterSetMessageHandlerIterator, error)
	WatchSetMessageHandler(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterSetMessageHandler) (event.Subscription, error)
	ParseSetMessageHandler(log types.Log) (*ContractAttestationCenterSetMessageHandler, error)

	FilterSetMinimumTaskDefinitionVotingPower(opts *bind.FilterOpts) (*ContractAttestationCenterSetMinimumTaskDefinitionVotingPowerIterator, error)
	WatchSetMinimumTaskDefinitionVotingPower(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterSetMinimumTaskDefinitionVotingPower) (event.Subscription, error)
	ParseSetMinimumTaskDefinitionVotingPower(log types.Log) (*ContractAttestationCenterSetMinimumTaskDefinitionVotingPower, error)

	FilterSetRestrictedAttester(opts *bind.FilterOpts, taskDefinitionId []uint16) (*ContractAttestationCenterSetRestrictedAttesterIterator, error)
	WatchSetRestrictedAttester(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterSetRestrictedAttester, taskDefinitionId []uint16) (event.Subscription, error)
	ParseSetRestrictedAttester(log types.Log) (*ContractAttestationCenterSetRestrictedAttester, error)

	FilterTaskDefinitionCreated(opts *bind.FilterOpts) (*ContractAttestationCenterTaskDefinitionCreatedIterator, error)
	WatchTaskDefinitionCreated(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterTaskDefinitionCreated) (event.Subscription, error)
	ParseTaskDefinitionCreated(log types.Log) (*ContractAttestationCenterTaskDefinitionCreated, error)
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

// FeeCalculator is a free data retrieval call binding the contract method 0xb00eb9fe.
//
// Solidity: function feeCalculator() view returns(address)
func (_ContractAttestationCenter *ContractAttestationCenterCaller) FeeCalculator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAttestationCenter.contract.Call(opts, &out, "feeCalculator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeCalculator is a free data retrieval call binding the contract method 0xb00eb9fe.
//
// Solidity: function feeCalculator() view returns(address)
func (_ContractAttestationCenter *ContractAttestationCenterSession) FeeCalculator() (common.Address, error) {
	return _ContractAttestationCenter.Contract.FeeCalculator(&_ContractAttestationCenter.CallOpts)
}

// FeeCalculator is a free data retrieval call binding the contract method 0xb00eb9fe.
//
// Solidity: function feeCalculator() view returns(address)
func (_ContractAttestationCenter *ContractAttestationCenterCallerSession) FeeCalculator() (common.Address, error) {
	return _ContractAttestationCenter.Contract.FeeCalculator(&_ContractAttestationCenter.CallOpts)
}

// GetActiveOperatorsDetails is a free data retrieval call binding the contract method 0x9878eccb.
//
// Solidity: function getActiveOperatorsDetails() view returns((address,uint256,uint256,uint256,uint256[4])[] _operators)
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
// Solidity: function getActiveOperatorsDetails() view returns((address,uint256,uint256,uint256,uint256[4])[] _operators)
func (_ContractAttestationCenter *ContractAttestationCenterSession) GetActiveOperatorsDetails() ([]IAttestationCenterOperatorDetails, error) {
	return _ContractAttestationCenter.Contract.GetActiveOperatorsDetails(&_ContractAttestationCenter.CallOpts)
}

// GetActiveOperatorsDetails is a free data retrieval call binding the contract method 0x9878eccb.
//
// Solidity: function getActiveOperatorsDetails() view returns((address,uint256,uint256,uint256,uint256[4])[] _operators)
func (_ContractAttestationCenter *ContractAttestationCenterCallerSession) GetActiveOperatorsDetails() ([]IAttestationCenterOperatorDetails, error) {
	return _ContractAttestationCenter.Contract.GetActiveOperatorsDetails(&_ContractAttestationCenter.CallOpts)
}

// GetL2MessageHandler is a free data retrieval call binding the contract method 0xf0e058d1.
//
// Solidity: function getL2MessageHandler() view returns(address)
func (_ContractAttestationCenter *ContractAttestationCenterCaller) GetL2MessageHandler(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAttestationCenter.contract.Call(opts, &out, "getL2MessageHandler")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetL2MessageHandler is a free data retrieval call binding the contract method 0xf0e058d1.
//
// Solidity: function getL2MessageHandler() view returns(address)
func (_ContractAttestationCenter *ContractAttestationCenterSession) GetL2MessageHandler() (common.Address, error) {
	return _ContractAttestationCenter.Contract.GetL2MessageHandler(&_ContractAttestationCenter.CallOpts)
}

// GetL2MessageHandler is a free data retrieval call binding the contract method 0xf0e058d1.
//
// Solidity: function getL2MessageHandler() view returns(address)
func (_ContractAttestationCenter *ContractAttestationCenterCallerSession) GetL2MessageHandler() (common.Address, error) {
	return _ContractAttestationCenter.Contract.GetL2MessageHandler(&_ContractAttestationCenter.CallOpts)
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

// GetTaskDefinitionDetails is a free data retrieval call binding the contract method 0x2fd60f1d.
//
// Solidity: function getTaskDefinitionDetails(uint16 _taskDefinitionId) view returns((uint16,bool,bool,string,uint256,uint256,uint256,uint256,uint256,uint256,uint256[],uint256))
func (_ContractAttestationCenter *ContractAttestationCenterCaller) GetTaskDefinitionDetails(opts *bind.CallOpts, _taskDefinitionId uint16) (TaskDefinition, error) {
	var out []interface{}
	err := _ContractAttestationCenter.contract.Call(opts, &out, "getTaskDefinitionDetails", _taskDefinitionId)

	if err != nil {
		return *new(TaskDefinition), err
	}

	out0 := *abi.ConvertType(out[0], new(TaskDefinition)).(*TaskDefinition)

	return out0, err

}

// GetTaskDefinitionDetails is a free data retrieval call binding the contract method 0x2fd60f1d.
//
// Solidity: function getTaskDefinitionDetails(uint16 _taskDefinitionId) view returns((uint16,bool,bool,string,uint256,uint256,uint256,uint256,uint256,uint256,uint256[],uint256))
func (_ContractAttestationCenter *ContractAttestationCenterSession) GetTaskDefinitionDetails(_taskDefinitionId uint16) (TaskDefinition, error) {
	return _ContractAttestationCenter.Contract.GetTaskDefinitionDetails(&_ContractAttestationCenter.CallOpts, _taskDefinitionId)
}

// GetTaskDefinitionDetails is a free data retrieval call binding the contract method 0x2fd60f1d.
//
// Solidity: function getTaskDefinitionDetails(uint16 _taskDefinitionId) view returns((uint16,bool,bool,string,uint256,uint256,uint256,uint256,uint256,uint256,uint256[],uint256))
func (_ContractAttestationCenter *ContractAttestationCenterCallerSession) GetTaskDefinitionDetails(_taskDefinitionId uint16) (TaskDefinition, error) {
	return _ContractAttestationCenter.Contract.GetTaskDefinitionDetails(&_ContractAttestationCenter.CallOpts, _taskDefinitionId)
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

// PaymentRequestsRole is a free data retrieval call binding the contract method 0xc8e5bff8.
//
// Solidity: function paymentRequestsRole() view returns(address _currentPaymentRequestsAdmin)
func (_ContractAttestationCenter *ContractAttestationCenterCaller) PaymentRequestsRole(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAttestationCenter.contract.Call(opts, &out, "paymentRequestsRole")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PaymentRequestsRole is a free data retrieval call binding the contract method 0xc8e5bff8.
//
// Solidity: function paymentRequestsRole() view returns(address _currentPaymentRequestsAdmin)
func (_ContractAttestationCenter *ContractAttestationCenterSession) PaymentRequestsRole() (common.Address, error) {
	return _ContractAttestationCenter.Contract.PaymentRequestsRole(&_ContractAttestationCenter.CallOpts)
}

// PaymentRequestsRole is a free data retrieval call binding the contract method 0xc8e5bff8.
//
// Solidity: function paymentRequestsRole() view returns(address _currentPaymentRequestsAdmin)
func (_ContractAttestationCenter *ContractAttestationCenterCallerSession) PaymentRequestsRole() (common.Address, error) {
	return _ContractAttestationCenter.Contract.PaymentRequestsRole(&_ContractAttestationCenter.CallOpts)
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

// ApplyCustomSlashing is a paid mutator transaction binding the contract method 0x24f8c18b.
//
// Solidity: function applyCustomSlashing(address _operator, (address,uint8,uint256)[] _slashingStakingContractInfos) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) ApplyCustomSlashing(opts *bind.TransactOpts, _operator common.Address, _slashingStakingContractInfos []ISlashingConfigSlashingStakingContractInfo) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "applyCustomSlashing", _operator, _slashingStakingContractInfos)
}

// ApplyCustomSlashing is a paid mutator transaction binding the contract method 0x24f8c18b.
//
// Solidity: function applyCustomSlashing(address _operator, (address,uint8,uint256)[] _slashingStakingContractInfos) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) ApplyCustomSlashing(_operator common.Address, _slashingStakingContractInfos []ISlashingConfigSlashingStakingContractInfo) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.ApplyCustomSlashing(&_ContractAttestationCenter.TransactOpts, _operator, _slashingStakingContractInfos)
}

// ApplyCustomSlashing is a paid mutator transaction binding the contract method 0x24f8c18b.
//
// Solidity: function applyCustomSlashing(address _operator, (address,uint8,uint256)[] _slashingStakingContractInfos) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) ApplyCustomSlashing(_operator common.Address, _slashingStakingContractInfos []ISlashingConfigSlashingStakingContractInfo) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.ApplyCustomSlashing(&_ContractAttestationCenter.TransactOpts, _operator, _slashingStakingContractInfos)
}

// CreateNewTaskDefinition is a paid mutator transaction binding the contract method 0x306fba1c.
//
// Solidity: function createNewTaskDefinition(string _name, (uint256,uint256,uint256,uint256,uint256,uint256,uint256[],uint256) _taskDefinitionParams) returns(uint16 _id)
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) CreateNewTaskDefinition(opts *bind.TransactOpts, _name string, _taskDefinitionParams TaskDefinitionParamsV2) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "createNewTaskDefinition", _name, _taskDefinitionParams)
}

// CreateNewTaskDefinition is a paid mutator transaction binding the contract method 0x306fba1c.
//
// Solidity: function createNewTaskDefinition(string _name, (uint256,uint256,uint256,uint256,uint256,uint256,uint256[],uint256) _taskDefinitionParams) returns(uint16 _id)
func (_ContractAttestationCenter *ContractAttestationCenterSession) CreateNewTaskDefinition(_name string, _taskDefinitionParams TaskDefinitionParamsV2) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.CreateNewTaskDefinition(&_ContractAttestationCenter.TransactOpts, _name, _taskDefinitionParams)
}

// CreateNewTaskDefinition is a paid mutator transaction binding the contract method 0x306fba1c.
//
// Solidity: function createNewTaskDefinition(string _name, (uint256,uint256,uint256,uint256,uint256,uint256,uint256[],uint256) _taskDefinitionParams) returns(uint16 _id)
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) CreateNewTaskDefinition(_name string, _taskDefinitionParams TaskDefinitionParamsV2) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.CreateNewTaskDefinition(&_ContractAttestationCenter.TransactOpts, _name, _taskDefinitionParams)
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

// MigrateSlashingFlow is a paid mutator transaction binding the contract method 0x460fd1be.
//
// Solidity: function migrateSlashingFlow(address _avsGovernanceMultisigOwner) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) MigrateSlashingFlow(opts *bind.TransactOpts, _avsGovernanceMultisigOwner common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "migrateSlashingFlow", _avsGovernanceMultisigOwner)
}

// MigrateSlashingFlow is a paid mutator transaction binding the contract method 0x460fd1be.
//
// Solidity: function migrateSlashingFlow(address _avsGovernanceMultisigOwner) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) MigrateSlashingFlow(_avsGovernanceMultisigOwner common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.MigrateSlashingFlow(&_ContractAttestationCenter.TransactOpts, _avsGovernanceMultisigOwner)
}

// MigrateSlashingFlow is a paid mutator transaction binding the contract method 0x460fd1be.
//
// Solidity: function migrateSlashingFlow(address _avsGovernanceMultisigOwner) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) MigrateSlashingFlow(_avsGovernanceMultisigOwner common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.MigrateSlashingFlow(&_ContractAttestationCenter.TransactOpts, _avsGovernanceMultisigOwner)
}

// Migration is a paid mutator transaction binding the contract method 0x1705a3bd.
//
// Solidity: function migration() returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) Migration(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "migration")
}

// Migration is a paid mutator transaction binding the contract method 0x1705a3bd.
//
// Solidity: function migration() returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) Migration() (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.Migration(&_ContractAttestationCenter.TransactOpts)
}

// Migration is a paid mutator transaction binding the contract method 0x1705a3bd.
//
// Solidity: function migration() returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) Migration() (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.Migration(&_ContractAttestationCenter.TransactOpts)
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

// RequestBatchPayment is a paid mutator transaction binding the contract method 0x2bfcb2bc.
//
// Solidity: function requestBatchPayment(uint256 _from, uint256 _to, uint8 _submissionType, bytes _distributionData) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) RequestBatchPayment(opts *bind.TransactOpts, _from *big.Int, _to *big.Int, _submissionType uint8, _distributionData []byte) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "requestBatchPayment", _from, _to, _submissionType, _distributionData)
}

// RequestBatchPayment is a paid mutator transaction binding the contract method 0x2bfcb2bc.
//
// Solidity: function requestBatchPayment(uint256 _from, uint256 _to, uint8 _submissionType, bytes _distributionData) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) RequestBatchPayment(_from *big.Int, _to *big.Int, _submissionType uint8, _distributionData []byte) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.RequestBatchPayment(&_ContractAttestationCenter.TransactOpts, _from, _to, _submissionType, _distributionData)
}

// RequestBatchPayment is a paid mutator transaction binding the contract method 0x2bfcb2bc.
//
// Solidity: function requestBatchPayment(uint256 _from, uint256 _to, uint8 _submissionType, bytes _distributionData) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) RequestBatchPayment(_from *big.Int, _to *big.Int, _submissionType uint8, _distributionData []byte) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.RequestBatchPayment(&_ContractAttestationCenter.TransactOpts, _from, _to, _submissionType, _distributionData)
}

// RequestBatchPayment0 is a paid mutator transaction binding the contract method 0xd0c613a9.
//
// Solidity: function requestBatchPayment(uint8 _submissionType, bytes _distributionData) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) RequestBatchPayment0(opts *bind.TransactOpts, _submissionType uint8, _distributionData []byte) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "requestBatchPayment0", _submissionType, _distributionData)
}

// RequestBatchPayment0 is a paid mutator transaction binding the contract method 0xd0c613a9.
//
// Solidity: function requestBatchPayment(uint8 _submissionType, bytes _distributionData) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) RequestBatchPayment0(_submissionType uint8, _distributionData []byte) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.RequestBatchPayment0(&_ContractAttestationCenter.TransactOpts, _submissionType, _distributionData)
}

// RequestBatchPayment0 is a paid mutator transaction binding the contract method 0xd0c613a9.
//
// Solidity: function requestBatchPayment(uint8 _submissionType, bytes _distributionData) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) RequestBatchPayment0(_submissionType uint8, _distributionData []byte) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.RequestBatchPayment0(&_ContractAttestationCenter.TransactOpts, _submissionType, _distributionData)
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

// RevokeEjector is a paid mutator transaction binding the contract method 0x7de6e60d.
//
// Solidity: function revokeEjector(address _ejector) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) RevokeEjector(opts *bind.TransactOpts, _ejector common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "revokeEjector", _ejector)
}

// RevokeEjector is a paid mutator transaction binding the contract method 0x7de6e60d.
//
// Solidity: function revokeEjector(address _ejector) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) RevokeEjector(_ejector common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.RevokeEjector(&_ContractAttestationCenter.TransactOpts, _ejector)
}

// RevokeEjector is a paid mutator transaction binding the contract method 0x7de6e60d.
//
// Solidity: function revokeEjector(address _ejector) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) RevokeEjector(_ejector common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.RevokeEjector(&_ContractAttestationCenter.TransactOpts, _ejector)
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

// RevokeSlasher is a paid mutator transaction binding the contract method 0x74a91439.
//
// Solidity: function revokeSlasher(address _slasher) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) RevokeSlasher(opts *bind.TransactOpts, _slasher common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "revokeSlasher", _slasher)
}

// RevokeSlasher is a paid mutator transaction binding the contract method 0x74a91439.
//
// Solidity: function revokeSlasher(address _slasher) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) RevokeSlasher(_slasher common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.RevokeSlasher(&_ContractAttestationCenter.TransactOpts, _slasher)
}

// RevokeSlasher is a paid mutator transaction binding the contract method 0x74a91439.
//
// Solidity: function revokeSlasher(address _slasher) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) RevokeSlasher(_slasher common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.RevokeSlasher(&_ContractAttestationCenter.TransactOpts, _slasher)
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

// SetBaseRewardFeeForAggregator is a paid mutator transaction binding the contract method 0x6b3ba95f.
//
// Solidity: function setBaseRewardFeeForAggregator(uint16 _taskDefinitionId, uint256 _baseRewardFeeForAggregator) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) SetBaseRewardFeeForAggregator(opts *bind.TransactOpts, _taskDefinitionId uint16, _baseRewardFeeForAggregator *big.Int) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "setBaseRewardFeeForAggregator", _taskDefinitionId, _baseRewardFeeForAggregator)
}

// SetBaseRewardFeeForAggregator is a paid mutator transaction binding the contract method 0x6b3ba95f.
//
// Solidity: function setBaseRewardFeeForAggregator(uint16 _taskDefinitionId, uint256 _baseRewardFeeForAggregator) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) SetBaseRewardFeeForAggregator(_taskDefinitionId uint16, _baseRewardFeeForAggregator *big.Int) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SetBaseRewardFeeForAggregator(&_ContractAttestationCenter.TransactOpts, _taskDefinitionId, _baseRewardFeeForAggregator)
}

// SetBaseRewardFeeForAggregator is a paid mutator transaction binding the contract method 0x6b3ba95f.
//
// Solidity: function setBaseRewardFeeForAggregator(uint16 _taskDefinitionId, uint256 _baseRewardFeeForAggregator) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) SetBaseRewardFeeForAggregator(_taskDefinitionId uint16, _baseRewardFeeForAggregator *big.Int) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SetBaseRewardFeeForAggregator(&_ContractAttestationCenter.TransactOpts, _taskDefinitionId, _baseRewardFeeForAggregator)
}

// SetBaseRewardFeeForAttesters is a paid mutator transaction binding the contract method 0xd33cf7d8.
//
// Solidity: function setBaseRewardFeeForAttesters(uint16 _taskDefinitionId, uint256 _baseRewardFeeForAttesters) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) SetBaseRewardFeeForAttesters(opts *bind.TransactOpts, _taskDefinitionId uint16, _baseRewardFeeForAttesters *big.Int) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "setBaseRewardFeeForAttesters", _taskDefinitionId, _baseRewardFeeForAttesters)
}

// SetBaseRewardFeeForAttesters is a paid mutator transaction binding the contract method 0xd33cf7d8.
//
// Solidity: function setBaseRewardFeeForAttesters(uint16 _taskDefinitionId, uint256 _baseRewardFeeForAttesters) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) SetBaseRewardFeeForAttesters(_taskDefinitionId uint16, _baseRewardFeeForAttesters *big.Int) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SetBaseRewardFeeForAttesters(&_ContractAttestationCenter.TransactOpts, _taskDefinitionId, _baseRewardFeeForAttesters)
}

// SetBaseRewardFeeForAttesters is a paid mutator transaction binding the contract method 0xd33cf7d8.
//
// Solidity: function setBaseRewardFeeForAttesters(uint16 _taskDefinitionId, uint256 _baseRewardFeeForAttesters) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) SetBaseRewardFeeForAttesters(_taskDefinitionId uint16, _baseRewardFeeForAttesters *big.Int) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SetBaseRewardFeeForAttesters(&_ContractAttestationCenter.TransactOpts, _taskDefinitionId, _baseRewardFeeForAttesters)
}

// SetBaseRewardFeeForPerformer is a paid mutator transaction binding the contract method 0x19636e2a.
//
// Solidity: function setBaseRewardFeeForPerformer(uint16 _taskDefinitionId, uint256 _baseRewardFeeForPerformer) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) SetBaseRewardFeeForPerformer(opts *bind.TransactOpts, _taskDefinitionId uint16, _baseRewardFeeForPerformer *big.Int) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "setBaseRewardFeeForPerformer", _taskDefinitionId, _baseRewardFeeForPerformer)
}

// SetBaseRewardFeeForPerformer is a paid mutator transaction binding the contract method 0x19636e2a.
//
// Solidity: function setBaseRewardFeeForPerformer(uint16 _taskDefinitionId, uint256 _baseRewardFeeForPerformer) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) SetBaseRewardFeeForPerformer(_taskDefinitionId uint16, _baseRewardFeeForPerformer *big.Int) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SetBaseRewardFeeForPerformer(&_ContractAttestationCenter.TransactOpts, _taskDefinitionId, _baseRewardFeeForPerformer)
}

// SetBaseRewardFeeForPerformer is a paid mutator transaction binding the contract method 0x19636e2a.
//
// Solidity: function setBaseRewardFeeForPerformer(uint16 _taskDefinitionId, uint256 _baseRewardFeeForPerformer) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) SetBaseRewardFeeForPerformer(_taskDefinitionId uint16, _baseRewardFeeForPerformer *big.Int) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SetBaseRewardFeeForPerformer(&_ContractAttestationCenter.TransactOpts, _taskDefinitionId, _baseRewardFeeForPerformer)
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

// SetChallengerRewardFee is a paid mutator transaction binding the contract method 0x7c742fb6.
//
// Solidity: function setChallengerRewardFee(uint256 _challengerRewardFee) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) SetChallengerRewardFee(opts *bind.TransactOpts, _challengerRewardFee *big.Int) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "setChallengerRewardFee", _challengerRewardFee)
}

// SetChallengerRewardFee is a paid mutator transaction binding the contract method 0x7c742fb6.
//
// Solidity: function setChallengerRewardFee(uint256 _challengerRewardFee) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) SetChallengerRewardFee(_challengerRewardFee *big.Int) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SetChallengerRewardFee(&_ContractAttestationCenter.TransactOpts, _challengerRewardFee)
}

// SetChallengerRewardFee is a paid mutator transaction binding the contract method 0x7c742fb6.
//
// Solidity: function setChallengerRewardFee(uint256 _challengerRewardFee) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) SetChallengerRewardFee(_challengerRewardFee *big.Int) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SetChallengerRewardFee(&_ContractAttestationCenter.TransactOpts, _challengerRewardFee)
}

// SetEjector is a paid mutator transaction binding the contract method 0x2cdd1e86.
//
// Solidity: function setEjector(address _ejector) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) SetEjector(opts *bind.TransactOpts, _ejector common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "setEjector", _ejector)
}

// SetEjector is a paid mutator transaction binding the contract method 0x2cdd1e86.
//
// Solidity: function setEjector(address _ejector) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) SetEjector(_ejector common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SetEjector(&_ContractAttestationCenter.TransactOpts, _ejector)
}

// SetEjector is a paid mutator transaction binding the contract method 0x2cdd1e86.
//
// Solidity: function setEjector(address _ejector) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) SetEjector(_ejector common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SetEjector(&_ContractAttestationCenter.TransactOpts, _ejector)
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

// SetInternalTaskHandler is a paid mutator transaction binding the contract method 0x8f350ef4.
//
// Solidity: function setInternalTaskHandler(address _internalTaskHandler) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) SetInternalTaskHandler(opts *bind.TransactOpts, _internalTaskHandler common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "setInternalTaskHandler", _internalTaskHandler)
}

// SetInternalTaskHandler is a paid mutator transaction binding the contract method 0x8f350ef4.
//
// Solidity: function setInternalTaskHandler(address _internalTaskHandler) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) SetInternalTaskHandler(_internalTaskHandler common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SetInternalTaskHandler(&_ContractAttestationCenter.TransactOpts, _internalTaskHandler)
}

// SetInternalTaskHandler is a paid mutator transaction binding the contract method 0x8f350ef4.
//
// Solidity: function setInternalTaskHandler(address _internalTaskHandler) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) SetInternalTaskHandler(_internalTaskHandler common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SetInternalTaskHandler(&_ContractAttestationCenter.TransactOpts, _internalTaskHandler)
}

// SetIsIncorrectAttestationSlashingEnabled is a paid mutator transaction binding the contract method 0xcc021985.
//
// Solidity: function setIsIncorrectAttestationSlashingEnabled(uint16 _taskDefinitionId, bool _isIncorrectAttestationSlashingEnabled) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) SetIsIncorrectAttestationSlashingEnabled(opts *bind.TransactOpts, _taskDefinitionId uint16, _isIncorrectAttestationSlashingEnabled bool) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "setIsIncorrectAttestationSlashingEnabled", _taskDefinitionId, _isIncorrectAttestationSlashingEnabled)
}

// SetIsIncorrectAttestationSlashingEnabled is a paid mutator transaction binding the contract method 0xcc021985.
//
// Solidity: function setIsIncorrectAttestationSlashingEnabled(uint16 _taskDefinitionId, bool _isIncorrectAttestationSlashingEnabled) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) SetIsIncorrectAttestationSlashingEnabled(_taskDefinitionId uint16, _isIncorrectAttestationSlashingEnabled bool) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SetIsIncorrectAttestationSlashingEnabled(&_ContractAttestationCenter.TransactOpts, _taskDefinitionId, _isIncorrectAttestationSlashingEnabled)
}

// SetIsIncorrectAttestationSlashingEnabled is a paid mutator transaction binding the contract method 0xcc021985.
//
// Solidity: function setIsIncorrectAttestationSlashingEnabled(uint16 _taskDefinitionId, bool _isIncorrectAttestationSlashingEnabled) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) SetIsIncorrectAttestationSlashingEnabled(_taskDefinitionId uint16, _isIncorrectAttestationSlashingEnabled bool) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SetIsIncorrectAttestationSlashingEnabled(&_ContractAttestationCenter.TransactOpts, _taskDefinitionId, _isIncorrectAttestationSlashingEnabled)
}

// SetIsOpenAggregator is a paid mutator transaction binding the contract method 0x51720945.
//
// Solidity: function setIsOpenAggregator(bool _isOpenAggregator) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) SetIsOpenAggregator(opts *bind.TransactOpts, _isOpenAggregator bool) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "setIsOpenAggregator", _isOpenAggregator)
}

// SetIsOpenAggregator is a paid mutator transaction binding the contract method 0x51720945.
//
// Solidity: function setIsOpenAggregator(bool _isOpenAggregator) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) SetIsOpenAggregator(_isOpenAggregator bool) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SetIsOpenAggregator(&_ContractAttestationCenter.TransactOpts, _isOpenAggregator)
}

// SetIsOpenAggregator is a paid mutator transaction binding the contract method 0x51720945.
//
// Solidity: function setIsOpenAggregator(bool _isOpenAggregator) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) SetIsOpenAggregator(_isOpenAggregator bool) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SetIsOpenAggregator(&_ContractAttestationCenter.TransactOpts, _isOpenAggregator)
}

// SetIsRejectedTaskSlashingEnabled is a paid mutator transaction binding the contract method 0xeaa6867a.
//
// Solidity: function setIsRejectedTaskSlashingEnabled(uint16 _taskDefinitionId, bool _isRejectedTaskSlashingEnabled) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) SetIsRejectedTaskSlashingEnabled(opts *bind.TransactOpts, _taskDefinitionId uint16, _isRejectedTaskSlashingEnabled bool) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "setIsRejectedTaskSlashingEnabled", _taskDefinitionId, _isRejectedTaskSlashingEnabled)
}

// SetIsRejectedTaskSlashingEnabled is a paid mutator transaction binding the contract method 0xeaa6867a.
//
// Solidity: function setIsRejectedTaskSlashingEnabled(uint16 _taskDefinitionId, bool _isRejectedTaskSlashingEnabled) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) SetIsRejectedTaskSlashingEnabled(_taskDefinitionId uint16, _isRejectedTaskSlashingEnabled bool) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SetIsRejectedTaskSlashingEnabled(&_ContractAttestationCenter.TransactOpts, _taskDefinitionId, _isRejectedTaskSlashingEnabled)
}

// SetIsRejectedTaskSlashingEnabled is a paid mutator transaction binding the contract method 0xeaa6867a.
//
// Solidity: function setIsRejectedTaskSlashingEnabled(uint16 _taskDefinitionId, bool _isRejectedTaskSlashingEnabled) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) SetIsRejectedTaskSlashingEnabled(_taskDefinitionId uint16, _isRejectedTaskSlashingEnabled bool) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SetIsRejectedTaskSlashingEnabled(&_ContractAttestationCenter.TransactOpts, _taskDefinitionId, _isRejectedTaskSlashingEnabled)
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

// SetPaymentRequestsRole is a paid mutator transaction binding the contract method 0x3e3e99ab.
//
// Solidity: function setPaymentRequestsRole(address _newPaymentRequestsAdmin) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) SetPaymentRequestsRole(opts *bind.TransactOpts, _newPaymentRequestsAdmin common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "setPaymentRequestsRole", _newPaymentRequestsAdmin)
}

// SetPaymentRequestsRole is a paid mutator transaction binding the contract method 0x3e3e99ab.
//
// Solidity: function setPaymentRequestsRole(address _newPaymentRequestsAdmin) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) SetPaymentRequestsRole(_newPaymentRequestsAdmin common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SetPaymentRequestsRole(&_ContractAttestationCenter.TransactOpts, _newPaymentRequestsAdmin)
}

// SetPaymentRequestsRole is a paid mutator transaction binding the contract method 0x3e3e99ab.
//
// Solidity: function setPaymentRequestsRole(address _newPaymentRequestsAdmin) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) SetPaymentRequestsRole(_newPaymentRequestsAdmin common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SetPaymentRequestsRole(&_ContractAttestationCenter.TransactOpts, _newPaymentRequestsAdmin)
}

// SetSlasher is a paid mutator transaction binding the contract method 0xaabc2496.
//
// Solidity: function setSlasher(address _slasher) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) SetSlasher(opts *bind.TransactOpts, _slasher common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "setSlasher", _slasher)
}

// SetSlasher is a paid mutator transaction binding the contract method 0xaabc2496.
//
// Solidity: function setSlasher(address _slasher) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) SetSlasher(_slasher common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SetSlasher(&_ContractAttestationCenter.TransactOpts, _slasher)
}

// SetSlasher is a paid mutator transaction binding the contract method 0xaabc2496.
//
// Solidity: function setSlasher(address _slasher) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) SetSlasher(_slasher common.Address) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SetSlasher(&_ContractAttestationCenter.TransactOpts, _slasher)
}

// SetTaskDefinitionMaximumNumberOfAttesters is a paid mutator transaction binding the contract method 0x5fdd0ac7.
//
// Solidity: function setTaskDefinitionMaximumNumberOfAttesters(uint16 _taskDefinitionId, uint256 _maximumNumberOfAttesters) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) SetTaskDefinitionMaximumNumberOfAttesters(opts *bind.TransactOpts, _taskDefinitionId uint16, _maximumNumberOfAttesters *big.Int) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "setTaskDefinitionMaximumNumberOfAttesters", _taskDefinitionId, _maximumNumberOfAttesters)
}

// SetTaskDefinitionMaximumNumberOfAttesters is a paid mutator transaction binding the contract method 0x5fdd0ac7.
//
// Solidity: function setTaskDefinitionMaximumNumberOfAttesters(uint16 _taskDefinitionId, uint256 _maximumNumberOfAttesters) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) SetTaskDefinitionMaximumNumberOfAttesters(_taskDefinitionId uint16, _maximumNumberOfAttesters *big.Int) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SetTaskDefinitionMaximumNumberOfAttesters(&_ContractAttestationCenter.TransactOpts, _taskDefinitionId, _maximumNumberOfAttesters)
}

// SetTaskDefinitionMaximumNumberOfAttesters is a paid mutator transaction binding the contract method 0x5fdd0ac7.
//
// Solidity: function setTaskDefinitionMaximumNumberOfAttesters(uint16 _taskDefinitionId, uint256 _maximumNumberOfAttesters) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) SetTaskDefinitionMaximumNumberOfAttesters(_taskDefinitionId uint16, _maximumNumberOfAttesters *big.Int) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SetTaskDefinitionMaximumNumberOfAttesters(&_ContractAttestationCenter.TransactOpts, _taskDefinitionId, _maximumNumberOfAttesters)
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

// SetTaskDefinitionRestrictedAttesters is a paid mutator transaction binding the contract method 0x58e75c82.
//
// Solidity: function setTaskDefinitionRestrictedAttesters(uint16 _taskDefinitionId, uint256[] _restrictedAttesterIds) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) SetTaskDefinitionRestrictedAttesters(opts *bind.TransactOpts, _taskDefinitionId uint16, _restrictedAttesterIds []*big.Int) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "setTaskDefinitionRestrictedAttesters", _taskDefinitionId, _restrictedAttesterIds)
}

// SetTaskDefinitionRestrictedAttesters is a paid mutator transaction binding the contract method 0x58e75c82.
//
// Solidity: function setTaskDefinitionRestrictedAttesters(uint16 _taskDefinitionId, uint256[] _restrictedAttesterIds) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) SetTaskDefinitionRestrictedAttesters(_taskDefinitionId uint16, _restrictedAttesterIds []*big.Int) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SetTaskDefinitionRestrictedAttesters(&_ContractAttestationCenter.TransactOpts, _taskDefinitionId, _restrictedAttesterIds)
}

// SetTaskDefinitionRestrictedAttesters is a paid mutator transaction binding the contract method 0x58e75c82.
//
// Solidity: function setTaskDefinitionRestrictedAttesters(uint16 _taskDefinitionId, uint256[] _restrictedAttesterIds) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) SetTaskDefinitionRestrictedAttesters(_taskDefinitionId uint16, _restrictedAttesterIds []*big.Int) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.SetTaskDefinitionRestrictedAttesters(&_ContractAttestationCenter.TransactOpts, _taskDefinitionId, _restrictedAttesterIds)
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

// UpdateInternalTaskConfig is a paid mutator transaction binding the contract method 0x84efb8ea.
//
// Solidity: function updateInternalTaskConfig(uint16 _taskDefinitionId, (bool,uint48,uint8,bytes) _config) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactor) UpdateInternalTaskConfig(opts *bind.TransactOpts, _taskDefinitionId uint16, _config IInternalTaskHandlerInternalTaskConfig) (*types.Transaction, error) {
	return _ContractAttestationCenter.contract.Transact(opts, "updateInternalTaskConfig", _taskDefinitionId, _config)
}

// UpdateInternalTaskConfig is a paid mutator transaction binding the contract method 0x84efb8ea.
//
// Solidity: function updateInternalTaskConfig(uint16 _taskDefinitionId, (bool,uint48,uint8,bytes) _config) returns()
func (_ContractAttestationCenter *ContractAttestationCenterSession) UpdateInternalTaskConfig(_taskDefinitionId uint16, _config IInternalTaskHandlerInternalTaskConfig) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.UpdateInternalTaskConfig(&_ContractAttestationCenter.TransactOpts, _taskDefinitionId, _config)
}

// UpdateInternalTaskConfig is a paid mutator transaction binding the contract method 0x84efb8ea.
//
// Solidity: function updateInternalTaskConfig(uint16 _taskDefinitionId, (bool,uint48,uint8,bytes) _config) returns()
func (_ContractAttestationCenter *ContractAttestationCenterTransactorSession) UpdateInternalTaskConfig(_taskDefinitionId uint16, _config IInternalTaskHandlerInternalTaskConfig) (*types.Transaction, error) {
	return _ContractAttestationCenter.Contract.UpdateInternalTaskConfig(&_ContractAttestationCenter.TransactOpts, _taskDefinitionId, _config)
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
	Operators          []IRewardsDistributorPaymentRequestMessage
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

// ContractAttestationCenterIsOpenAggregatorSetIterator is returned from FilterIsOpenAggregatorSet and is used to iterate over the raw logs and unpacked data for IsOpenAggregatorSet events raised by the ContractAttestationCenter contract.
type ContractAttestationCenterIsOpenAggregatorSetIterator struct {
	Event *ContractAttestationCenterIsOpenAggregatorSet // Event containing the contract specifics and raw log

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
func (it *ContractAttestationCenterIsOpenAggregatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAttestationCenterIsOpenAggregatorSet)
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
		it.Event = new(ContractAttestationCenterIsOpenAggregatorSet)
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
func (it *ContractAttestationCenterIsOpenAggregatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAttestationCenterIsOpenAggregatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAttestationCenterIsOpenAggregatorSet represents a IsOpenAggregatorSet event raised by the ContractAttestationCenter contract.
type ContractAttestationCenterIsOpenAggregatorSet struct {
	IsOpenAggregator bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterIsOpenAggregatorSet is a free log retrieval operation binding the contract event 0x660f8962d35177409475cadba4b4e75639dc5b6af824076b1b50616a39ec565f.
//
// Solidity: event IsOpenAggregatorSet(bool isOpenAggregator)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) FilterIsOpenAggregatorSet(opts *bind.FilterOpts) (*ContractAttestationCenterIsOpenAggregatorSetIterator, error) {

	logs, sub, err := _ContractAttestationCenter.contract.FilterLogs(opts, "IsOpenAggregatorSet")
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterIsOpenAggregatorSetIterator{contract: _ContractAttestationCenter.contract, event: "IsOpenAggregatorSet", logs: logs, sub: sub}, nil
}

// WatchIsOpenAggregatorSet is a free log subscription operation binding the contract event 0x660f8962d35177409475cadba4b4e75639dc5b6af824076b1b50616a39ec565f.
//
// Solidity: event IsOpenAggregatorSet(bool isOpenAggregator)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) WatchIsOpenAggregatorSet(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterIsOpenAggregatorSet) (event.Subscription, error) {

	logs, sub, err := _ContractAttestationCenter.contract.WatchLogs(opts, "IsOpenAggregatorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAttestationCenterIsOpenAggregatorSet)
				if err := _ContractAttestationCenter.contract.UnpackLog(event, "IsOpenAggregatorSet", log); err != nil {
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

// ParseIsOpenAggregatorSet is a log parse operation binding the contract event 0x660f8962d35177409475cadba4b4e75639dc5b6af824076b1b50616a39ec565f.
//
// Solidity: event IsOpenAggregatorSet(bool isOpenAggregator)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) ParseIsOpenAggregatorSet(log types.Log) (*ContractAttestationCenterIsOpenAggregatorSet, error) {
	event := new(ContractAttestationCenterIsOpenAggregatorSet)
	if err := _ContractAttestationCenter.contract.UnpackLog(event, "IsOpenAggregatorSet", log); err != nil {
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

// ContractAttestationCenterOperatorSlashRequestedIterator is returned from FilterOperatorSlashRequested and is used to iterate over the raw logs and unpacked data for OperatorSlashRequested events raised by the ContractAttestationCenter contract.
type ContractAttestationCenterOperatorSlashRequestedIterator struct {
	Event *ContractAttestationCenterOperatorSlashRequested // Event containing the contract specifics and raw log

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
func (it *ContractAttestationCenterOperatorSlashRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAttestationCenterOperatorSlashRequested)
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
		it.Event = new(ContractAttestationCenterOperatorSlashRequested)
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
func (it *ContractAttestationCenterOperatorSlashRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAttestationCenterOperatorSlashRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAttestationCenterOperatorSlashRequested represents a OperatorSlashRequested event raised by the ContractAttestationCenter contract.
type ContractAttestationCenterOperatorSlashRequested struct {
	Operator                     common.Address
	SlashingStakingContractInfos []ISlashingConfigSlashingStakingContractInfo
	Raw                          types.Log // Blockchain specific contextual infos
}

// FilterOperatorSlashRequested is a free log retrieval operation binding the contract event 0xdc2e6e7be0395520e6d44a5a96b4b957de4ce57bd39ff7b4be319bdfad0d20e8.
//
// Solidity: event OperatorSlashRequested(address operator, (address,uint8,uint256)[] slashingStakingContractInfos)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) FilterOperatorSlashRequested(opts *bind.FilterOpts) (*ContractAttestationCenterOperatorSlashRequestedIterator, error) {

	logs, sub, err := _ContractAttestationCenter.contract.FilterLogs(opts, "OperatorSlashRequested")
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterOperatorSlashRequestedIterator{contract: _ContractAttestationCenter.contract, event: "OperatorSlashRequested", logs: logs, sub: sub}, nil
}

// WatchOperatorSlashRequested is a free log subscription operation binding the contract event 0xdc2e6e7be0395520e6d44a5a96b4b957de4ce57bd39ff7b4be319bdfad0d20e8.
//
// Solidity: event OperatorSlashRequested(address operator, (address,uint8,uint256)[] slashingStakingContractInfos)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) WatchOperatorSlashRequested(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterOperatorSlashRequested) (event.Subscription, error) {

	logs, sub, err := _ContractAttestationCenter.contract.WatchLogs(opts, "OperatorSlashRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAttestationCenterOperatorSlashRequested)
				if err := _ContractAttestationCenter.contract.UnpackLog(event, "OperatorSlashRequested", log); err != nil {
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

// ParseOperatorSlashRequested is a log parse operation binding the contract event 0xdc2e6e7be0395520e6d44a5a96b4b957de4ce57bd39ff7b4be319bdfad0d20e8.
//
// Solidity: event OperatorSlashRequested(address operator, (address,uint8,uint256)[] slashingStakingContractInfos)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) ParseOperatorSlashRequested(log types.Log) (*ContractAttestationCenterOperatorSlashRequested, error) {
	event := new(ContractAttestationCenterOperatorSlashRequested)
	if err := _ContractAttestationCenter.contract.UnpackLog(event, "OperatorSlashRequested", log); err != nil {
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
	Operators          []IRewardsDistributorPaymentRequestMessage
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

// ContractAttestationCenterSetBaseRewardFeeForAggregatorIterator is returned from FilterSetBaseRewardFeeForAggregator and is used to iterate over the raw logs and unpacked data for SetBaseRewardFeeForAggregator events raised by the ContractAttestationCenter contract.
type ContractAttestationCenterSetBaseRewardFeeForAggregatorIterator struct {
	Event *ContractAttestationCenterSetBaseRewardFeeForAggregator // Event containing the contract specifics and raw log

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
func (it *ContractAttestationCenterSetBaseRewardFeeForAggregatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAttestationCenterSetBaseRewardFeeForAggregator)
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
		it.Event = new(ContractAttestationCenterSetBaseRewardFeeForAggregator)
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
func (it *ContractAttestationCenterSetBaseRewardFeeForAggregatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAttestationCenterSetBaseRewardFeeForAggregatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAttestationCenterSetBaseRewardFeeForAggregator represents a SetBaseRewardFeeForAggregator event raised by the ContractAttestationCenter contract.
type ContractAttestationCenterSetBaseRewardFeeForAggregator struct {
	TaskDefinitionId           uint16
	BaseRewardFeeForAggregator *big.Int
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterSetBaseRewardFeeForAggregator is a free log retrieval operation binding the contract event 0x4793dc4cc042d7bb55de6974e13f051eaed18732c9f5fea5830bc76604b38dc2.
//
// Solidity: event SetBaseRewardFeeForAggregator(uint16 indexed taskDefinitionId, uint256 baseRewardFeeForAggregator)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) FilterSetBaseRewardFeeForAggregator(opts *bind.FilterOpts, taskDefinitionId []uint16) (*ContractAttestationCenterSetBaseRewardFeeForAggregatorIterator, error) {

	var taskDefinitionIdRule []interface{}
	for _, taskDefinitionIdItem := range taskDefinitionId {
		taskDefinitionIdRule = append(taskDefinitionIdRule, taskDefinitionIdItem)
	}

	logs, sub, err := _ContractAttestationCenter.contract.FilterLogs(opts, "SetBaseRewardFeeForAggregator", taskDefinitionIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterSetBaseRewardFeeForAggregatorIterator{contract: _ContractAttestationCenter.contract, event: "SetBaseRewardFeeForAggregator", logs: logs, sub: sub}, nil
}

// WatchSetBaseRewardFeeForAggregator is a free log subscription operation binding the contract event 0x4793dc4cc042d7bb55de6974e13f051eaed18732c9f5fea5830bc76604b38dc2.
//
// Solidity: event SetBaseRewardFeeForAggregator(uint16 indexed taskDefinitionId, uint256 baseRewardFeeForAggregator)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) WatchSetBaseRewardFeeForAggregator(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterSetBaseRewardFeeForAggregator, taskDefinitionId []uint16) (event.Subscription, error) {

	var taskDefinitionIdRule []interface{}
	for _, taskDefinitionIdItem := range taskDefinitionId {
		taskDefinitionIdRule = append(taskDefinitionIdRule, taskDefinitionIdItem)
	}

	logs, sub, err := _ContractAttestationCenter.contract.WatchLogs(opts, "SetBaseRewardFeeForAggregator", taskDefinitionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAttestationCenterSetBaseRewardFeeForAggregator)
				if err := _ContractAttestationCenter.contract.UnpackLog(event, "SetBaseRewardFeeForAggregator", log); err != nil {
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

// ParseSetBaseRewardFeeForAggregator is a log parse operation binding the contract event 0x4793dc4cc042d7bb55de6974e13f051eaed18732c9f5fea5830bc76604b38dc2.
//
// Solidity: event SetBaseRewardFeeForAggregator(uint16 indexed taskDefinitionId, uint256 baseRewardFeeForAggregator)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) ParseSetBaseRewardFeeForAggregator(log types.Log) (*ContractAttestationCenterSetBaseRewardFeeForAggregator, error) {
	event := new(ContractAttestationCenterSetBaseRewardFeeForAggregator)
	if err := _ContractAttestationCenter.contract.UnpackLog(event, "SetBaseRewardFeeForAggregator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAttestationCenterSetBaseRewardFeeForAttestersIterator is returned from FilterSetBaseRewardFeeForAttesters and is used to iterate over the raw logs and unpacked data for SetBaseRewardFeeForAttesters events raised by the ContractAttestationCenter contract.
type ContractAttestationCenterSetBaseRewardFeeForAttestersIterator struct {
	Event *ContractAttestationCenterSetBaseRewardFeeForAttesters // Event containing the contract specifics and raw log

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
func (it *ContractAttestationCenterSetBaseRewardFeeForAttestersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAttestationCenterSetBaseRewardFeeForAttesters)
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
		it.Event = new(ContractAttestationCenterSetBaseRewardFeeForAttesters)
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
func (it *ContractAttestationCenterSetBaseRewardFeeForAttestersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAttestationCenterSetBaseRewardFeeForAttestersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAttestationCenterSetBaseRewardFeeForAttesters represents a SetBaseRewardFeeForAttesters event raised by the ContractAttestationCenter contract.
type ContractAttestationCenterSetBaseRewardFeeForAttesters struct {
	TaskDefinitionId          uint16
	BaseRewardFeeForAttesters *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterSetBaseRewardFeeForAttesters is a free log retrieval operation binding the contract event 0x13e7ed4e07d552b0c14aaba5b1725d6e31eb89bc4d0e1d1f1103fc30a1d419e5.
//
// Solidity: event SetBaseRewardFeeForAttesters(uint16 indexed taskDefinitionId, uint256 baseRewardFeeForAttesters)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) FilterSetBaseRewardFeeForAttesters(opts *bind.FilterOpts, taskDefinitionId []uint16) (*ContractAttestationCenterSetBaseRewardFeeForAttestersIterator, error) {

	var taskDefinitionIdRule []interface{}
	for _, taskDefinitionIdItem := range taskDefinitionId {
		taskDefinitionIdRule = append(taskDefinitionIdRule, taskDefinitionIdItem)
	}

	logs, sub, err := _ContractAttestationCenter.contract.FilterLogs(opts, "SetBaseRewardFeeForAttesters", taskDefinitionIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterSetBaseRewardFeeForAttestersIterator{contract: _ContractAttestationCenter.contract, event: "SetBaseRewardFeeForAttesters", logs: logs, sub: sub}, nil
}

// WatchSetBaseRewardFeeForAttesters is a free log subscription operation binding the contract event 0x13e7ed4e07d552b0c14aaba5b1725d6e31eb89bc4d0e1d1f1103fc30a1d419e5.
//
// Solidity: event SetBaseRewardFeeForAttesters(uint16 indexed taskDefinitionId, uint256 baseRewardFeeForAttesters)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) WatchSetBaseRewardFeeForAttesters(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterSetBaseRewardFeeForAttesters, taskDefinitionId []uint16) (event.Subscription, error) {

	var taskDefinitionIdRule []interface{}
	for _, taskDefinitionIdItem := range taskDefinitionId {
		taskDefinitionIdRule = append(taskDefinitionIdRule, taskDefinitionIdItem)
	}

	logs, sub, err := _ContractAttestationCenter.contract.WatchLogs(opts, "SetBaseRewardFeeForAttesters", taskDefinitionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAttestationCenterSetBaseRewardFeeForAttesters)
				if err := _ContractAttestationCenter.contract.UnpackLog(event, "SetBaseRewardFeeForAttesters", log); err != nil {
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

// ParseSetBaseRewardFeeForAttesters is a log parse operation binding the contract event 0x13e7ed4e07d552b0c14aaba5b1725d6e31eb89bc4d0e1d1f1103fc30a1d419e5.
//
// Solidity: event SetBaseRewardFeeForAttesters(uint16 indexed taskDefinitionId, uint256 baseRewardFeeForAttesters)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) ParseSetBaseRewardFeeForAttesters(log types.Log) (*ContractAttestationCenterSetBaseRewardFeeForAttesters, error) {
	event := new(ContractAttestationCenterSetBaseRewardFeeForAttesters)
	if err := _ContractAttestationCenter.contract.UnpackLog(event, "SetBaseRewardFeeForAttesters", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAttestationCenterSetBaseRewardFeeForPerformerIterator is returned from FilterSetBaseRewardFeeForPerformer and is used to iterate over the raw logs and unpacked data for SetBaseRewardFeeForPerformer events raised by the ContractAttestationCenter contract.
type ContractAttestationCenterSetBaseRewardFeeForPerformerIterator struct {
	Event *ContractAttestationCenterSetBaseRewardFeeForPerformer // Event containing the contract specifics and raw log

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
func (it *ContractAttestationCenterSetBaseRewardFeeForPerformerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAttestationCenterSetBaseRewardFeeForPerformer)
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
		it.Event = new(ContractAttestationCenterSetBaseRewardFeeForPerformer)
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
func (it *ContractAttestationCenterSetBaseRewardFeeForPerformerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAttestationCenterSetBaseRewardFeeForPerformerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAttestationCenterSetBaseRewardFeeForPerformer represents a SetBaseRewardFeeForPerformer event raised by the ContractAttestationCenter contract.
type ContractAttestationCenterSetBaseRewardFeeForPerformer struct {
	TaskDefinitionId          uint16
	BaseRewardFeeForPerformer *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterSetBaseRewardFeeForPerformer is a free log retrieval operation binding the contract event 0xbbc9e65198877320a28bfeba97eec9813ff314fc72b3dc385a99a631cfa39ef4.
//
// Solidity: event SetBaseRewardFeeForPerformer(uint16 indexed taskDefinitionId, uint256 baseRewardFeeForPerformer)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) FilterSetBaseRewardFeeForPerformer(opts *bind.FilterOpts, taskDefinitionId []uint16) (*ContractAttestationCenterSetBaseRewardFeeForPerformerIterator, error) {

	var taskDefinitionIdRule []interface{}
	for _, taskDefinitionIdItem := range taskDefinitionId {
		taskDefinitionIdRule = append(taskDefinitionIdRule, taskDefinitionIdItem)
	}

	logs, sub, err := _ContractAttestationCenter.contract.FilterLogs(opts, "SetBaseRewardFeeForPerformer", taskDefinitionIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterSetBaseRewardFeeForPerformerIterator{contract: _ContractAttestationCenter.contract, event: "SetBaseRewardFeeForPerformer", logs: logs, sub: sub}, nil
}

// WatchSetBaseRewardFeeForPerformer is a free log subscription operation binding the contract event 0xbbc9e65198877320a28bfeba97eec9813ff314fc72b3dc385a99a631cfa39ef4.
//
// Solidity: event SetBaseRewardFeeForPerformer(uint16 indexed taskDefinitionId, uint256 baseRewardFeeForPerformer)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) WatchSetBaseRewardFeeForPerformer(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterSetBaseRewardFeeForPerformer, taskDefinitionId []uint16) (event.Subscription, error) {

	var taskDefinitionIdRule []interface{}
	for _, taskDefinitionIdItem := range taskDefinitionId {
		taskDefinitionIdRule = append(taskDefinitionIdRule, taskDefinitionIdItem)
	}

	logs, sub, err := _ContractAttestationCenter.contract.WatchLogs(opts, "SetBaseRewardFeeForPerformer", taskDefinitionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAttestationCenterSetBaseRewardFeeForPerformer)
				if err := _ContractAttestationCenter.contract.UnpackLog(event, "SetBaseRewardFeeForPerformer", log); err != nil {
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

// ParseSetBaseRewardFeeForPerformer is a log parse operation binding the contract event 0xbbc9e65198877320a28bfeba97eec9813ff314fc72b3dc385a99a631cfa39ef4.
//
// Solidity: event SetBaseRewardFeeForPerformer(uint16 indexed taskDefinitionId, uint256 baseRewardFeeForPerformer)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) ParseSetBaseRewardFeeForPerformer(log types.Log) (*ContractAttestationCenterSetBaseRewardFeeForPerformer, error) {
	event := new(ContractAttestationCenterSetBaseRewardFeeForPerformer)
	if err := _ContractAttestationCenter.contract.UnpackLog(event, "SetBaseRewardFeeForPerformer", log); err != nil {
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

// ContractAttestationCenterSetChallengerRewardFeeIterator is returned from FilterSetChallengerRewardFee and is used to iterate over the raw logs and unpacked data for SetChallengerRewardFee events raised by the ContractAttestationCenter contract.
type ContractAttestationCenterSetChallengerRewardFeeIterator struct {
	Event *ContractAttestationCenterSetChallengerRewardFee // Event containing the contract specifics and raw log

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
func (it *ContractAttestationCenterSetChallengerRewardFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAttestationCenterSetChallengerRewardFee)
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
		it.Event = new(ContractAttestationCenterSetChallengerRewardFee)
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
func (it *ContractAttestationCenterSetChallengerRewardFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAttestationCenterSetChallengerRewardFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAttestationCenterSetChallengerRewardFee represents a SetChallengerRewardFee event raised by the ContractAttestationCenter contract.
type ContractAttestationCenterSetChallengerRewardFee struct {
	ChallengerRewardFee *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetChallengerRewardFee is a free log retrieval operation binding the contract event 0xabfb14e76f333086004dff417c020e5e0edde86950bcc21ec5c64b590eee222b.
//
// Solidity: event SetChallengerRewardFee(uint256 challengerRewardFee)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) FilterSetChallengerRewardFee(opts *bind.FilterOpts) (*ContractAttestationCenterSetChallengerRewardFeeIterator, error) {

	logs, sub, err := _ContractAttestationCenter.contract.FilterLogs(opts, "SetChallengerRewardFee")
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterSetChallengerRewardFeeIterator{contract: _ContractAttestationCenter.contract, event: "SetChallengerRewardFee", logs: logs, sub: sub}, nil
}

// WatchSetChallengerRewardFee is a free log subscription operation binding the contract event 0xabfb14e76f333086004dff417c020e5e0edde86950bcc21ec5c64b590eee222b.
//
// Solidity: event SetChallengerRewardFee(uint256 challengerRewardFee)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) WatchSetChallengerRewardFee(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterSetChallengerRewardFee) (event.Subscription, error) {

	logs, sub, err := _ContractAttestationCenter.contract.WatchLogs(opts, "SetChallengerRewardFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAttestationCenterSetChallengerRewardFee)
				if err := _ContractAttestationCenter.contract.UnpackLog(event, "SetChallengerRewardFee", log); err != nil {
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

// ParseSetChallengerRewardFee is a log parse operation binding the contract event 0xabfb14e76f333086004dff417c020e5e0edde86950bcc21ec5c64b590eee222b.
//
// Solidity: event SetChallengerRewardFee(uint256 challengerRewardFee)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) ParseSetChallengerRewardFee(log types.Log) (*ContractAttestationCenterSetChallengerRewardFee, error) {
	event := new(ContractAttestationCenterSetChallengerRewardFee)
	if err := _ContractAttestationCenter.contract.UnpackLog(event, "SetChallengerRewardFee", log); err != nil {
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

// ContractAttestationCenterSetInternalTaskHandlerIterator is returned from FilterSetInternalTaskHandler and is used to iterate over the raw logs and unpacked data for SetInternalTaskHandler events raised by the ContractAttestationCenter contract.
type ContractAttestationCenterSetInternalTaskHandlerIterator struct {
	Event *ContractAttestationCenterSetInternalTaskHandler // Event containing the contract specifics and raw log

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
func (it *ContractAttestationCenterSetInternalTaskHandlerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAttestationCenterSetInternalTaskHandler)
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
		it.Event = new(ContractAttestationCenterSetInternalTaskHandler)
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
func (it *ContractAttestationCenterSetInternalTaskHandlerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAttestationCenterSetInternalTaskHandlerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAttestationCenterSetInternalTaskHandler represents a SetInternalTaskHandler event raised by the ContractAttestationCenter contract.
type ContractAttestationCenterSetInternalTaskHandler struct {
	NewInternalTaskHandler common.Address
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetInternalTaskHandler is a free log retrieval operation binding the contract event 0x422f68e5e2c91706d1dfc874a022d02851fa68385617d07a5efe61b07adca4c3.
//
// Solidity: event SetInternalTaskHandler(address newInternalTaskHandler)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) FilterSetInternalTaskHandler(opts *bind.FilterOpts) (*ContractAttestationCenterSetInternalTaskHandlerIterator, error) {

	logs, sub, err := _ContractAttestationCenter.contract.FilterLogs(opts, "SetInternalTaskHandler")
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterSetInternalTaskHandlerIterator{contract: _ContractAttestationCenter.contract, event: "SetInternalTaskHandler", logs: logs, sub: sub}, nil
}

// WatchSetInternalTaskHandler is a free log subscription operation binding the contract event 0x422f68e5e2c91706d1dfc874a022d02851fa68385617d07a5efe61b07adca4c3.
//
// Solidity: event SetInternalTaskHandler(address newInternalTaskHandler)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) WatchSetInternalTaskHandler(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterSetInternalTaskHandler) (event.Subscription, error) {

	logs, sub, err := _ContractAttestationCenter.contract.WatchLogs(opts, "SetInternalTaskHandler")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAttestationCenterSetInternalTaskHandler)
				if err := _ContractAttestationCenter.contract.UnpackLog(event, "SetInternalTaskHandler", log); err != nil {
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

// ParseSetInternalTaskHandler is a log parse operation binding the contract event 0x422f68e5e2c91706d1dfc874a022d02851fa68385617d07a5efe61b07adca4c3.
//
// Solidity: event SetInternalTaskHandler(address newInternalTaskHandler)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) ParseSetInternalTaskHandler(log types.Log) (*ContractAttestationCenterSetInternalTaskHandler, error) {
	event := new(ContractAttestationCenterSetInternalTaskHandler)
	if err := _ContractAttestationCenter.contract.UnpackLog(event, "SetInternalTaskHandler", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAttestationCenterSetIsIncorrectAttestationSlashingEnabledIterator is returned from FilterSetIsIncorrectAttestationSlashingEnabled and is used to iterate over the raw logs and unpacked data for SetIsIncorrectAttestationSlashingEnabled events raised by the ContractAttestationCenter contract.
type ContractAttestationCenterSetIsIncorrectAttestationSlashingEnabledIterator struct {
	Event *ContractAttestationCenterSetIsIncorrectAttestationSlashingEnabled // Event containing the contract specifics and raw log

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
func (it *ContractAttestationCenterSetIsIncorrectAttestationSlashingEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAttestationCenterSetIsIncorrectAttestationSlashingEnabled)
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
		it.Event = new(ContractAttestationCenterSetIsIncorrectAttestationSlashingEnabled)
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
func (it *ContractAttestationCenterSetIsIncorrectAttestationSlashingEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAttestationCenterSetIsIncorrectAttestationSlashingEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAttestationCenterSetIsIncorrectAttestationSlashingEnabled represents a SetIsIncorrectAttestationSlashingEnabled event raised by the ContractAttestationCenter contract.
type ContractAttestationCenterSetIsIncorrectAttestationSlashingEnabled struct {
	TaskDefinitionId                      uint16
	IsIncorrectAttestationSlashingEnabled bool
	Raw                                   types.Log // Blockchain specific contextual infos
}

// FilterSetIsIncorrectAttestationSlashingEnabled is a free log retrieval operation binding the contract event 0x85e223b0f927390a7ced3882ec7361e1701b985e1908393bf858abcfa23b2307.
//
// Solidity: event SetIsIncorrectAttestationSlashingEnabled(uint16 indexed taskDefinitionId, bool isIncorrectAttestationSlashingEnabled)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) FilterSetIsIncorrectAttestationSlashingEnabled(opts *bind.FilterOpts, taskDefinitionId []uint16) (*ContractAttestationCenterSetIsIncorrectAttestationSlashingEnabledIterator, error) {

	var taskDefinitionIdRule []interface{}
	for _, taskDefinitionIdItem := range taskDefinitionId {
		taskDefinitionIdRule = append(taskDefinitionIdRule, taskDefinitionIdItem)
	}

	logs, sub, err := _ContractAttestationCenter.contract.FilterLogs(opts, "SetIsIncorrectAttestationSlashingEnabled", taskDefinitionIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterSetIsIncorrectAttestationSlashingEnabledIterator{contract: _ContractAttestationCenter.contract, event: "SetIsIncorrectAttestationSlashingEnabled", logs: logs, sub: sub}, nil
}

// WatchSetIsIncorrectAttestationSlashingEnabled is a free log subscription operation binding the contract event 0x85e223b0f927390a7ced3882ec7361e1701b985e1908393bf858abcfa23b2307.
//
// Solidity: event SetIsIncorrectAttestationSlashingEnabled(uint16 indexed taskDefinitionId, bool isIncorrectAttestationSlashingEnabled)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) WatchSetIsIncorrectAttestationSlashingEnabled(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterSetIsIncorrectAttestationSlashingEnabled, taskDefinitionId []uint16) (event.Subscription, error) {

	var taskDefinitionIdRule []interface{}
	for _, taskDefinitionIdItem := range taskDefinitionId {
		taskDefinitionIdRule = append(taskDefinitionIdRule, taskDefinitionIdItem)
	}

	logs, sub, err := _ContractAttestationCenter.contract.WatchLogs(opts, "SetIsIncorrectAttestationSlashingEnabled", taskDefinitionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAttestationCenterSetIsIncorrectAttestationSlashingEnabled)
				if err := _ContractAttestationCenter.contract.UnpackLog(event, "SetIsIncorrectAttestationSlashingEnabled", log); err != nil {
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

// ParseSetIsIncorrectAttestationSlashingEnabled is a log parse operation binding the contract event 0x85e223b0f927390a7ced3882ec7361e1701b985e1908393bf858abcfa23b2307.
//
// Solidity: event SetIsIncorrectAttestationSlashingEnabled(uint16 indexed taskDefinitionId, bool isIncorrectAttestationSlashingEnabled)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) ParseSetIsIncorrectAttestationSlashingEnabled(log types.Log) (*ContractAttestationCenterSetIsIncorrectAttestationSlashingEnabled, error) {
	event := new(ContractAttestationCenterSetIsIncorrectAttestationSlashingEnabled)
	if err := _ContractAttestationCenter.contract.UnpackLog(event, "SetIsIncorrectAttestationSlashingEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAttestationCenterSetIsRejectedTaskSlashingEnabledIterator is returned from FilterSetIsRejectedTaskSlashingEnabled and is used to iterate over the raw logs and unpacked data for SetIsRejectedTaskSlashingEnabled events raised by the ContractAttestationCenter contract.
type ContractAttestationCenterSetIsRejectedTaskSlashingEnabledIterator struct {
	Event *ContractAttestationCenterSetIsRejectedTaskSlashingEnabled // Event containing the contract specifics and raw log

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
func (it *ContractAttestationCenterSetIsRejectedTaskSlashingEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAttestationCenterSetIsRejectedTaskSlashingEnabled)
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
		it.Event = new(ContractAttestationCenterSetIsRejectedTaskSlashingEnabled)
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
func (it *ContractAttestationCenterSetIsRejectedTaskSlashingEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAttestationCenterSetIsRejectedTaskSlashingEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAttestationCenterSetIsRejectedTaskSlashingEnabled represents a SetIsRejectedTaskSlashingEnabled event raised by the ContractAttestationCenter contract.
type ContractAttestationCenterSetIsRejectedTaskSlashingEnabled struct {
	TaskDefinitionId              uint16
	IsRejectedTaskSlashingEnabled bool
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterSetIsRejectedTaskSlashingEnabled is a free log retrieval operation binding the contract event 0xbfff858ae099e475f126ce7aed2c36b4ddb5eed760f2cac98d57cd5fc9ae34db.
//
// Solidity: event SetIsRejectedTaskSlashingEnabled(uint16 indexed taskDefinitionId, bool isRejectedTaskSlashingEnabled)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) FilterSetIsRejectedTaskSlashingEnabled(opts *bind.FilterOpts, taskDefinitionId []uint16) (*ContractAttestationCenterSetIsRejectedTaskSlashingEnabledIterator, error) {

	var taskDefinitionIdRule []interface{}
	for _, taskDefinitionIdItem := range taskDefinitionId {
		taskDefinitionIdRule = append(taskDefinitionIdRule, taskDefinitionIdItem)
	}

	logs, sub, err := _ContractAttestationCenter.contract.FilterLogs(opts, "SetIsRejectedTaskSlashingEnabled", taskDefinitionIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterSetIsRejectedTaskSlashingEnabledIterator{contract: _ContractAttestationCenter.contract, event: "SetIsRejectedTaskSlashingEnabled", logs: logs, sub: sub}, nil
}

// WatchSetIsRejectedTaskSlashingEnabled is a free log subscription operation binding the contract event 0xbfff858ae099e475f126ce7aed2c36b4ddb5eed760f2cac98d57cd5fc9ae34db.
//
// Solidity: event SetIsRejectedTaskSlashingEnabled(uint16 indexed taskDefinitionId, bool isRejectedTaskSlashingEnabled)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) WatchSetIsRejectedTaskSlashingEnabled(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterSetIsRejectedTaskSlashingEnabled, taskDefinitionId []uint16) (event.Subscription, error) {

	var taskDefinitionIdRule []interface{}
	for _, taskDefinitionIdItem := range taskDefinitionId {
		taskDefinitionIdRule = append(taskDefinitionIdRule, taskDefinitionIdItem)
	}

	logs, sub, err := _ContractAttestationCenter.contract.WatchLogs(opts, "SetIsRejectedTaskSlashingEnabled", taskDefinitionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAttestationCenterSetIsRejectedTaskSlashingEnabled)
				if err := _ContractAttestationCenter.contract.UnpackLog(event, "SetIsRejectedTaskSlashingEnabled", log); err != nil {
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

// ParseSetIsRejectedTaskSlashingEnabled is a log parse operation binding the contract event 0xbfff858ae099e475f126ce7aed2c36b4ddb5eed760f2cac98d57cd5fc9ae34db.
//
// Solidity: event SetIsRejectedTaskSlashingEnabled(uint16 indexed taskDefinitionId, bool isRejectedTaskSlashingEnabled)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) ParseSetIsRejectedTaskSlashingEnabled(log types.Log) (*ContractAttestationCenterSetIsRejectedTaskSlashingEnabled, error) {
	event := new(ContractAttestationCenterSetIsRejectedTaskSlashingEnabled)
	if err := _ContractAttestationCenter.contract.UnpackLog(event, "SetIsRejectedTaskSlashingEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAttestationCenterSetMaximumNumberOfAttestersIterator is returned from FilterSetMaximumNumberOfAttesters and is used to iterate over the raw logs and unpacked data for SetMaximumNumberOfAttesters events raised by the ContractAttestationCenter contract.
type ContractAttestationCenterSetMaximumNumberOfAttestersIterator struct {
	Event *ContractAttestationCenterSetMaximumNumberOfAttesters // Event containing the contract specifics and raw log

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
func (it *ContractAttestationCenterSetMaximumNumberOfAttestersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAttestationCenterSetMaximumNumberOfAttesters)
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
		it.Event = new(ContractAttestationCenterSetMaximumNumberOfAttesters)
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
func (it *ContractAttestationCenterSetMaximumNumberOfAttestersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAttestationCenterSetMaximumNumberOfAttestersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAttestationCenterSetMaximumNumberOfAttesters represents a SetMaximumNumberOfAttesters event raised by the ContractAttestationCenter contract.
type ContractAttestationCenterSetMaximumNumberOfAttesters struct {
	TaskDefinitionId         uint16
	MaximumNumberOfAttesters *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSetMaximumNumberOfAttesters is a free log retrieval operation binding the contract event 0x08b4fb8850cdb8d9c6d276d628973bff1146e12e74dabf396622beca8b5ce5ae.
//
// Solidity: event SetMaximumNumberOfAttesters(uint16 indexed taskDefinitionId, uint256 maximumNumberOfAttesters)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) FilterSetMaximumNumberOfAttesters(opts *bind.FilterOpts, taskDefinitionId []uint16) (*ContractAttestationCenterSetMaximumNumberOfAttestersIterator, error) {

	var taskDefinitionIdRule []interface{}
	for _, taskDefinitionIdItem := range taskDefinitionId {
		taskDefinitionIdRule = append(taskDefinitionIdRule, taskDefinitionIdItem)
	}

	logs, sub, err := _ContractAttestationCenter.contract.FilterLogs(opts, "SetMaximumNumberOfAttesters", taskDefinitionIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterSetMaximumNumberOfAttestersIterator{contract: _ContractAttestationCenter.contract, event: "SetMaximumNumberOfAttesters", logs: logs, sub: sub}, nil
}

// WatchSetMaximumNumberOfAttesters is a free log subscription operation binding the contract event 0x08b4fb8850cdb8d9c6d276d628973bff1146e12e74dabf396622beca8b5ce5ae.
//
// Solidity: event SetMaximumNumberOfAttesters(uint16 indexed taskDefinitionId, uint256 maximumNumberOfAttesters)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) WatchSetMaximumNumberOfAttesters(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterSetMaximumNumberOfAttesters, taskDefinitionId []uint16) (event.Subscription, error) {

	var taskDefinitionIdRule []interface{}
	for _, taskDefinitionIdItem := range taskDefinitionId {
		taskDefinitionIdRule = append(taskDefinitionIdRule, taskDefinitionIdItem)
	}

	logs, sub, err := _ContractAttestationCenter.contract.WatchLogs(opts, "SetMaximumNumberOfAttesters", taskDefinitionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAttestationCenterSetMaximumNumberOfAttesters)
				if err := _ContractAttestationCenter.contract.UnpackLog(event, "SetMaximumNumberOfAttesters", log); err != nil {
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

// ParseSetMaximumNumberOfAttesters is a log parse operation binding the contract event 0x08b4fb8850cdb8d9c6d276d628973bff1146e12e74dabf396622beca8b5ce5ae.
//
// Solidity: event SetMaximumNumberOfAttesters(uint16 indexed taskDefinitionId, uint256 maximumNumberOfAttesters)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) ParseSetMaximumNumberOfAttesters(log types.Log) (*ContractAttestationCenterSetMaximumNumberOfAttesters, error) {
	event := new(ContractAttestationCenterSetMaximumNumberOfAttesters)
	if err := _ContractAttestationCenter.contract.UnpackLog(event, "SetMaximumNumberOfAttesters", log); err != nil {
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

// ContractAttestationCenterSetRestrictedAttesterIterator is returned from FilterSetRestrictedAttester and is used to iterate over the raw logs and unpacked data for SetRestrictedAttester events raised by the ContractAttestationCenter contract.
type ContractAttestationCenterSetRestrictedAttesterIterator struct {
	Event *ContractAttestationCenterSetRestrictedAttester // Event containing the contract specifics and raw log

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
func (it *ContractAttestationCenterSetRestrictedAttesterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAttestationCenterSetRestrictedAttester)
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
		it.Event = new(ContractAttestationCenterSetRestrictedAttester)
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
func (it *ContractAttestationCenterSetRestrictedAttesterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAttestationCenterSetRestrictedAttesterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAttestationCenterSetRestrictedAttester represents a SetRestrictedAttester event raised by the ContractAttestationCenter contract.
type ContractAttestationCenterSetRestrictedAttester struct {
	TaskDefinitionId      uint16
	RestrictedAttesterIds []*big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSetRestrictedAttester is a free log retrieval operation binding the contract event 0x7d9be67d2c086f7b6d17bd6e99430a2757c5b4dc4519288a62d5f167e441c240.
//
// Solidity: event SetRestrictedAttester(uint16 indexed taskDefinitionId, uint256[] restrictedAttesterIds)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) FilterSetRestrictedAttester(opts *bind.FilterOpts, taskDefinitionId []uint16) (*ContractAttestationCenterSetRestrictedAttesterIterator, error) {

	var taskDefinitionIdRule []interface{}
	for _, taskDefinitionIdItem := range taskDefinitionId {
		taskDefinitionIdRule = append(taskDefinitionIdRule, taskDefinitionIdItem)
	}

	logs, sub, err := _ContractAttestationCenter.contract.FilterLogs(opts, "SetRestrictedAttester", taskDefinitionIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterSetRestrictedAttesterIterator{contract: _ContractAttestationCenter.contract, event: "SetRestrictedAttester", logs: logs, sub: sub}, nil
}

// WatchSetRestrictedAttester is a free log subscription operation binding the contract event 0x7d9be67d2c086f7b6d17bd6e99430a2757c5b4dc4519288a62d5f167e441c240.
//
// Solidity: event SetRestrictedAttester(uint16 indexed taskDefinitionId, uint256[] restrictedAttesterIds)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) WatchSetRestrictedAttester(opts *bind.WatchOpts, sink chan<- *ContractAttestationCenterSetRestrictedAttester, taskDefinitionId []uint16) (event.Subscription, error) {

	var taskDefinitionIdRule []interface{}
	for _, taskDefinitionIdItem := range taskDefinitionId {
		taskDefinitionIdRule = append(taskDefinitionIdRule, taskDefinitionIdItem)
	}

	logs, sub, err := _ContractAttestationCenter.contract.WatchLogs(opts, "SetRestrictedAttester", taskDefinitionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAttestationCenterSetRestrictedAttester)
				if err := _ContractAttestationCenter.contract.UnpackLog(event, "SetRestrictedAttester", log); err != nil {
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

// ParseSetRestrictedAttester is a log parse operation binding the contract event 0x7d9be67d2c086f7b6d17bd6e99430a2757c5b4dc4519288a62d5f167e441c240.
//
// Solidity: event SetRestrictedAttester(uint16 indexed taskDefinitionId, uint256[] restrictedAttesterIds)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) ParseSetRestrictedAttester(log types.Log) (*ContractAttestationCenterSetRestrictedAttester, error) {
	event := new(ContractAttestationCenterSetRestrictedAttester)
	if err := _ContractAttestationCenter.contract.UnpackLog(event, "SetRestrictedAttester", log); err != nil {
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
	RestrictedAttesterIds      []*big.Int
	MaximumNumberOfAttesters   *big.Int
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterTaskDefinitionCreated is a free log retrieval operation binding the contract event 0x7eae3686e408afa807ebe15ce7995105877e176b7384d0656633435f6e58eab8.
//
// Solidity: event TaskDefinitionCreated(uint16 taskDefinitionId, string name, uint256 blockExpiry, uint256 baseRewardFeeForAttesters, uint256 baseRewardFeeForPerformer, uint256 baseRewardFeeForAggregator, uint256 disputePeriodBlocks, uint256 minimumVotingPower, uint256[] restrictedAttesterIds, uint256 maximumNumberOfAttesters)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) FilterTaskDefinitionCreated(opts *bind.FilterOpts) (*ContractAttestationCenterTaskDefinitionCreatedIterator, error) {

	logs, sub, err := _ContractAttestationCenter.contract.FilterLogs(opts, "TaskDefinitionCreated")
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCenterTaskDefinitionCreatedIterator{contract: _ContractAttestationCenter.contract, event: "TaskDefinitionCreated", logs: logs, sub: sub}, nil
}

// WatchTaskDefinitionCreated is a free log subscription operation binding the contract event 0x7eae3686e408afa807ebe15ce7995105877e176b7384d0656633435f6e58eab8.
//
// Solidity: event TaskDefinitionCreated(uint16 taskDefinitionId, string name, uint256 blockExpiry, uint256 baseRewardFeeForAttesters, uint256 baseRewardFeeForPerformer, uint256 baseRewardFeeForAggregator, uint256 disputePeriodBlocks, uint256 minimumVotingPower, uint256[] restrictedAttesterIds, uint256 maximumNumberOfAttesters)
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

// ParseTaskDefinitionCreated is a log parse operation binding the contract event 0x7eae3686e408afa807ebe15ce7995105877e176b7384d0656633435f6e58eab8.
//
// Solidity: event TaskDefinitionCreated(uint16 taskDefinitionId, string name, uint256 blockExpiry, uint256 baseRewardFeeForAttesters, uint256 baseRewardFeeForPerformer, uint256 baseRewardFeeForAggregator, uint256 disputePeriodBlocks, uint256 minimumVotingPower, uint256[] restrictedAttesterIds, uint256 maximumNumberOfAttesters)
func (_ContractAttestationCenter *ContractAttestationCenterFilterer) ParseTaskDefinitionCreated(log types.Log) (*ContractAttestationCenterTaskDefinitionCreated, error) {
	event := new(ContractAttestationCenterTaskDefinitionCreated)
	if err := _ContractAttestationCenter.contract.UnpackLog(event, "TaskDefinitionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
