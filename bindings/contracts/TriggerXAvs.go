// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package triggerxavs

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

// AVSParams is an auto generated low-level Go binding around an user-defined struct.
type AVSParams struct {
	Sender              common.Address
	AvsName             string
	MinStakeAmount      uint64
	TaskAddress         common.Address
	SlashAddress        common.Address
	RewardAddress       common.Address
	AvsOwnerAddresses   []common.Address
	WhitelistAddresses  []common.Address
	AssetIDs            []string
	AvsUnbondingPeriod  uint64
	MinSelfDelegation   uint64
	EpochIdentifier     string
	MiniOptInOperators  uint64
	MinTotalStakeAmount uint64
	AvsRewardProportion uint64
	AvsSlashProportion  uint64
}

// OperatorActivePower is an auto generated low-level Go binding around an user-defined struct.
type OperatorActivePower struct {
	Operator common.Address
	Power    *big.Int
}

// OperatorResInfo is an auto generated low-level Go binding around an user-defined struct.
type OperatorResInfo struct {
	TaskContractAddress common.Address
	TaskID              uint64
	OperatorAddress     common.Address
	TaskResponseHash    string
	TaskResponse        []byte
	BlsSignature        []byte
	Power               *big.Int
	Phase               uint8
}

// TaskDefinition is an auto generated low-level Go binding around an user-defined struct.
type TaskDefinition struct {
	TaskDefinitionId           uint8
	Name                       string
	BaseRewardFeeForAttesters  *big.Int
	BaseRewardFeeForPerformer  *big.Int
	BaseRewardFeeForAggregator *big.Int
	MinimumVotingPower         *big.Int
	RestrictedOperatorIds      []*big.Int
	MaximumNumberOfAttesters   *big.Int
}

// TaskDefinitionParams is an auto generated low-level Go binding around an user-defined struct.
type TaskDefinitionParams struct {
	BaseRewardFeeForAttesters  *big.Int
	BaseRewardFeeForPerformer  *big.Int
	BaseRewardFeeForAggregator *big.Int
	MinimumVotingPower         *big.Int
	RestrictedOperatorIds      []*big.Int
	MaximumNumberOfAttesters   *big.Int
}

// TaskInfo is an auto generated low-level Go binding around an user-defined struct.
type TaskInfo struct {
	TaskContractAddress     common.Address
	Name                    string
	Hash                    []byte
	TaskID                  uint64
	TaskResponsePeriod      uint64
	TaskStatisticalPeriod   uint64
	TaskChallengePeriod     uint64
	ThresholdPercentage     uint8
	StartingEpoch           uint64
	ActualThreshold         string
	OptInOperators          []common.Address
	SignedOperators         []common.Address
	NoSignedOperators       []common.Address
	ErrSignedOperators      []common.Address
	TaskTotalPower          string
	OperatorActivePower     []OperatorActivePower
	IsExpected              bool
	EligibleRewardOperators []common.Address
	EligibleSlashOperators  []common.Address
}

// TaskResultInfo is an auto generated low-level Go binding around an user-defined struct.
type TaskResultInfo struct {
	OperatorAddress     common.Address
	TaskResponseHash    string
	TaskResponse        []byte
	BlsSignature        []byte
	TaskContractAddress common.Address
	TaskID              uint64
	Phase               uint8
}

// TriggerXAvsMetaData contains all meta data concerning the TriggerXAvs contract.
var TriggerXAvsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"challenge\",\"inputs\":[{\"name\":\"taskID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"actualThreshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isExpected\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"eligibleRewardOperators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"eligibleSlashOperators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createTask\",\"inputs\":[{\"name\":\"taskName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"taskDefinitionId\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"taskResponsePeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"taskChallengePeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"thresholdPercentage\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"taskStatisticalPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"taskID\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createTaskDefinition\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"defParams\",\"type\":\"tuple\",\"internalType\":\"structTaskDefinitionParams\",\"components\":[{\"name\":\"baseRewardFeeForAttesters\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"baseRewardFeeForPerformer\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"baseRewardFeeForAggregator\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minimumVotingPower\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"restrictedOperatorIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"maximumNumberOfAttesters\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"taskDefinitionId\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAVSEpochIdentifier\",\"inputs\":[{\"name\":\"avsAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAVSUSDValue\",\"inputs\":[{\"name\":\"avsAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getChallengeInfo\",\"inputs\":[{\"name\":\"taskAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taskID\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentEpoch\",\"inputs\":[{\"name\":\"epochIdentifier\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int64\",\"internalType\":\"int64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorOptedUSDValue\",\"inputs\":[{\"name\":\"avsAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorTaskResponse\",\"inputs\":[{\"name\":\"taskAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taskID\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structTaskResultInfo\",\"components\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taskResponseHash\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"taskResponse\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blsSignature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"taskContractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taskID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"phase\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorTaskResponseList\",\"inputs\":[{\"name\":\"taskAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taskID\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorResInfo[]\",\"components\":[{\"name\":\"taskContractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taskID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"operatorAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taskResponseHash\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"taskResponse\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blsSignature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"power\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"phase\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOptInOperators\",\"inputs\":[{\"name\":\"avsAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRegisteredPubkey\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avsAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskDefinition\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structTaskDefinition\",\"components\":[{\"name\":\"taskDefinitionId\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"baseRewardFeeForAttesters\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"baseRewardFeeForPerformer\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"baseRewardFeeForAggregator\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minimumVotingPower\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"restrictedOperatorIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"maximumNumberOfAttesters\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskInfo\",\"inputs\":[{\"name\":\"taskAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taskID\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structTaskInfo\",\"components\":[{\"name\":\"taskContractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"hash\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"taskID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"taskResponsePeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"taskStatisticalPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"taskChallengePeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"thresholdPercentage\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"startingEpoch\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"actualThreshold\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"optInOperators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"signedOperators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"noSignedOperators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"errSignedOperators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"taskTotalPower\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"operatorActivePower\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorActivePower[]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"power\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"isExpected\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"eligibleRewardOperators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"eligibleSlashOperators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorSubmitTask\",\"inputs\":[{\"name\":\"taskID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"taskResponse\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blsSignature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"taskContractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"phase\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerAVS\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structAVSParams\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avsName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"minStakeAmount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"taskAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"slashAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rewardAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avsOwnerAddresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"whitelistAddresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"assetIDs\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"avsUnbondingPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"minSelfDelegation\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"epochIdentifier\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"miniOptInOperators\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"minTotalStakeAmount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"avsRewardProportion\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"avsSlashProportion\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerBLSPublicKey\",\"inputs\":[{\"name\":\"avsAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pubKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"pubKeyRegistrationSignature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rewardManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setRewardManager\",\"inputs\":[{\"name\":\"_rewardManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSlasher\",\"inputs\":[{\"name\":\"_slasher\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slasher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVS\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structAVSParams\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avsName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"minStakeAmount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"taskAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"slashAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rewardAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avsOwnerAddresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"whitelistAddresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"assetIDs\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"avsUnbondingPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"minSelfDelegation\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"epochIdentifier\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"miniOptInOperators\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"minTotalStakeAmount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"avsRewardProportion\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"avsSlashProportion\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"BLSPublicKeyRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"avsAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"pubKey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChallengeSubmitted\",\"inputs\":[{\"name\":\"taskID\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"isExpected\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorOptedIn\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorOptedOut\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardManagerUpdated\",\"inputs\":[{\"name\":\"newRewardManager\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlasherUpdated\",\"inputs\":[{\"name\":\"newSlasher\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCreated\",\"inputs\":[{\"name\":\"taskID\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"definitionHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"kind\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskDefinitionCreated\",\"inputs\":[{\"name\":\"taskDefinitionId\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskDefinitionCreated\",\"inputs\":[{\"name\":\"taskDefinitionId\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"baseRewardFeeForAttesters\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"baseRewardFeeForPerformer\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"baseRewardFeeForAggregator\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"minimumVotingPower\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"restrictedOperatorIds\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"maximumNumberOfAttesters\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskSubmitted\",\"inputs\":[{\"name\":\"taskID\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"phase\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
}

// TriggerXAvsABI is the input ABI used to generate the binding from.
// Deprecated: Use TriggerXAvsMetaData.ABI instead.
var TriggerXAvsABI = TriggerXAvsMetaData.ABI

// TriggerXAvs is an auto generated Go binding around an Ethereum contract.
type TriggerXAvs struct {
	TriggerXAvsCaller     // Read-only binding to the contract
	TriggerXAvsTransactor // Write-only binding to the contract
	TriggerXAvsFilterer   // Log filterer for contract events
}

// TriggerXAvsCaller is an auto generated read-only Go binding around an Ethereum contract.
type TriggerXAvsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TriggerXAvsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TriggerXAvsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TriggerXAvsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TriggerXAvsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TriggerXAvsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TriggerXAvsSession struct {
	Contract     *TriggerXAvs      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TriggerXAvsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TriggerXAvsCallerSession struct {
	Contract *TriggerXAvsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TriggerXAvsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TriggerXAvsTransactorSession struct {
	Contract     *TriggerXAvsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TriggerXAvsRaw is an auto generated low-level Go binding around an Ethereum contract.
type TriggerXAvsRaw struct {
	Contract *TriggerXAvs // Generic contract binding to access the raw methods on
}

// TriggerXAvsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TriggerXAvsCallerRaw struct {
	Contract *TriggerXAvsCaller // Generic read-only contract binding to access the raw methods on
}

// TriggerXAvsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TriggerXAvsTransactorRaw struct {
	Contract *TriggerXAvsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTriggerXAvs creates a new instance of TriggerXAvs, bound to a specific deployed contract.
func NewTriggerXAvs(address common.Address, backend bind.ContractBackend) (*TriggerXAvs, error) {
	contract, err := bindTriggerXAvs(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TriggerXAvs{TriggerXAvsCaller: TriggerXAvsCaller{contract: contract}, TriggerXAvsTransactor: TriggerXAvsTransactor{contract: contract}, TriggerXAvsFilterer: TriggerXAvsFilterer{contract: contract}}, nil
}

// NewTriggerXAvsCaller creates a new read-only instance of TriggerXAvs, bound to a specific deployed contract.
func NewTriggerXAvsCaller(address common.Address, caller bind.ContractCaller) (*TriggerXAvsCaller, error) {
	contract, err := bindTriggerXAvs(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TriggerXAvsCaller{contract: contract}, nil
}

// NewTriggerXAvsTransactor creates a new write-only instance of TriggerXAvs, bound to a specific deployed contract.
func NewTriggerXAvsTransactor(address common.Address, transactor bind.ContractTransactor) (*TriggerXAvsTransactor, error) {
	contract, err := bindTriggerXAvs(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TriggerXAvsTransactor{contract: contract}, nil
}

// NewTriggerXAvsFilterer creates a new log filterer instance of TriggerXAvs, bound to a specific deployed contract.
func NewTriggerXAvsFilterer(address common.Address, filterer bind.ContractFilterer) (*TriggerXAvsFilterer, error) {
	contract, err := bindTriggerXAvs(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TriggerXAvsFilterer{contract: contract}, nil
}

// bindTriggerXAvs binds a generic wrapper to an already deployed contract.
func bindTriggerXAvs(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TriggerXAvsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TriggerXAvs *TriggerXAvsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TriggerXAvs.Contract.TriggerXAvsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TriggerXAvs *TriggerXAvsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TriggerXAvs.Contract.TriggerXAvsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TriggerXAvs *TriggerXAvsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TriggerXAvs.Contract.TriggerXAvsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TriggerXAvs *TriggerXAvsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TriggerXAvs.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TriggerXAvs *TriggerXAvsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TriggerXAvs.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TriggerXAvs *TriggerXAvsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TriggerXAvs.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TriggerXAvs *TriggerXAvsCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TriggerXAvs.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TriggerXAvs *TriggerXAvsSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TriggerXAvs.Contract.UPGRADEINTERFACEVERSION(&_TriggerXAvs.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TriggerXAvs *TriggerXAvsCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TriggerXAvs.Contract.UPGRADEINTERFACEVERSION(&_TriggerXAvs.CallOpts)
}

// GetAVSEpochIdentifier is a free data retrieval call binding the contract method 0xe0938414.
//
// Solidity: function getAVSEpochIdentifier(address avsAddr) view returns(string)
func (_TriggerXAvs *TriggerXAvsCaller) GetAVSEpochIdentifier(opts *bind.CallOpts, avsAddr common.Address) (string, error) {
	var out []interface{}
	err := _TriggerXAvs.contract.Call(opts, &out, "getAVSEpochIdentifier", avsAddr)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetAVSEpochIdentifier is a free data retrieval call binding the contract method 0xe0938414.
//
// Solidity: function getAVSEpochIdentifier(address avsAddr) view returns(string)
func (_TriggerXAvs *TriggerXAvsSession) GetAVSEpochIdentifier(avsAddr common.Address) (string, error) {
	return _TriggerXAvs.Contract.GetAVSEpochIdentifier(&_TriggerXAvs.CallOpts, avsAddr)
}

// GetAVSEpochIdentifier is a free data retrieval call binding the contract method 0xe0938414.
//
// Solidity: function getAVSEpochIdentifier(address avsAddr) view returns(string)
func (_TriggerXAvs *TriggerXAvsCallerSession) GetAVSEpochIdentifier(avsAddr common.Address) (string, error) {
	return _TriggerXAvs.Contract.GetAVSEpochIdentifier(&_TriggerXAvs.CallOpts, avsAddr)
}

// GetAVSUSDValue is a free data retrieval call binding the contract method 0xdcf61b2c.
//
// Solidity: function getAVSUSDValue(address avsAddr) view returns(uint256)
func (_TriggerXAvs *TriggerXAvsCaller) GetAVSUSDValue(opts *bind.CallOpts, avsAddr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TriggerXAvs.contract.Call(opts, &out, "getAVSUSDValue", avsAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAVSUSDValue is a free data retrieval call binding the contract method 0xdcf61b2c.
//
// Solidity: function getAVSUSDValue(address avsAddr) view returns(uint256)
func (_TriggerXAvs *TriggerXAvsSession) GetAVSUSDValue(avsAddr common.Address) (*big.Int, error) {
	return _TriggerXAvs.Contract.GetAVSUSDValue(&_TriggerXAvs.CallOpts, avsAddr)
}

// GetAVSUSDValue is a free data retrieval call binding the contract method 0xdcf61b2c.
//
// Solidity: function getAVSUSDValue(address avsAddr) view returns(uint256)
func (_TriggerXAvs *TriggerXAvsCallerSession) GetAVSUSDValue(avsAddr common.Address) (*big.Int, error) {
	return _TriggerXAvs.Contract.GetAVSUSDValue(&_TriggerXAvs.CallOpts, avsAddr)
}

// GetChallengeInfo is a free data retrieval call binding the contract method 0x6d6ac37f.
//
// Solidity: function getChallengeInfo(address taskAddress, uint64 taskID) view returns(address)
func (_TriggerXAvs *TriggerXAvsCaller) GetChallengeInfo(opts *bind.CallOpts, taskAddress common.Address, taskID uint64) (common.Address, error) {
	var out []interface{}
	err := _TriggerXAvs.contract.Call(opts, &out, "getChallengeInfo", taskAddress, taskID)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetChallengeInfo is a free data retrieval call binding the contract method 0x6d6ac37f.
//
// Solidity: function getChallengeInfo(address taskAddress, uint64 taskID) view returns(address)
func (_TriggerXAvs *TriggerXAvsSession) GetChallengeInfo(taskAddress common.Address, taskID uint64) (common.Address, error) {
	return _TriggerXAvs.Contract.GetChallengeInfo(&_TriggerXAvs.CallOpts, taskAddress, taskID)
}

// GetChallengeInfo is a free data retrieval call binding the contract method 0x6d6ac37f.
//
// Solidity: function getChallengeInfo(address taskAddress, uint64 taskID) view returns(address)
func (_TriggerXAvs *TriggerXAvsCallerSession) GetChallengeInfo(taskAddress common.Address, taskID uint64) (common.Address, error) {
	return _TriggerXAvs.Contract.GetChallengeInfo(&_TriggerXAvs.CallOpts, taskAddress, taskID)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0x992907fb.
//
// Solidity: function getCurrentEpoch(string epochIdentifier) view returns(int64)
func (_TriggerXAvs *TriggerXAvsCaller) GetCurrentEpoch(opts *bind.CallOpts, epochIdentifier string) (int64, error) {
	var out []interface{}
	err := _TriggerXAvs.contract.Call(opts, &out, "getCurrentEpoch", epochIdentifier)

	if err != nil {
		return *new(int64), err
	}

	out0 := *abi.ConvertType(out[0], new(int64)).(*int64)

	return out0, err

}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0x992907fb.
//
// Solidity: function getCurrentEpoch(string epochIdentifier) view returns(int64)
func (_TriggerXAvs *TriggerXAvsSession) GetCurrentEpoch(epochIdentifier string) (int64, error) {
	return _TriggerXAvs.Contract.GetCurrentEpoch(&_TriggerXAvs.CallOpts, epochIdentifier)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0x992907fb.
//
// Solidity: function getCurrentEpoch(string epochIdentifier) view returns(int64)
func (_TriggerXAvs *TriggerXAvsCallerSession) GetCurrentEpoch(epochIdentifier string) (int64, error) {
	return _TriggerXAvs.Contract.GetCurrentEpoch(&_TriggerXAvs.CallOpts, epochIdentifier)
}

// GetOperatorOptedUSDValue is a free data retrieval call binding the contract method 0x4d568f24.
//
// Solidity: function getOperatorOptedUSDValue(address avsAddr, address operatorAddr) view returns(uint256)
func (_TriggerXAvs *TriggerXAvsCaller) GetOperatorOptedUSDValue(opts *bind.CallOpts, avsAddr common.Address, operatorAddr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TriggerXAvs.contract.Call(opts, &out, "getOperatorOptedUSDValue", avsAddr, operatorAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorOptedUSDValue is a free data retrieval call binding the contract method 0x4d568f24.
//
// Solidity: function getOperatorOptedUSDValue(address avsAddr, address operatorAddr) view returns(uint256)
func (_TriggerXAvs *TriggerXAvsSession) GetOperatorOptedUSDValue(avsAddr common.Address, operatorAddr common.Address) (*big.Int, error) {
	return _TriggerXAvs.Contract.GetOperatorOptedUSDValue(&_TriggerXAvs.CallOpts, avsAddr, operatorAddr)
}

// GetOperatorOptedUSDValue is a free data retrieval call binding the contract method 0x4d568f24.
//
// Solidity: function getOperatorOptedUSDValue(address avsAddr, address operatorAddr) view returns(uint256)
func (_TriggerXAvs *TriggerXAvsCallerSession) GetOperatorOptedUSDValue(avsAddr common.Address, operatorAddr common.Address) (*big.Int, error) {
	return _TriggerXAvs.Contract.GetOperatorOptedUSDValue(&_TriggerXAvs.CallOpts, avsAddr, operatorAddr)
}

// GetOperatorTaskResponse is a free data retrieval call binding the contract method 0x16395dc4.
//
// Solidity: function getOperatorTaskResponse(address taskAddress, address operator, uint64 taskID) view returns((address,string,bytes,bytes,address,uint64,uint8))
func (_TriggerXAvs *TriggerXAvsCaller) GetOperatorTaskResponse(opts *bind.CallOpts, taskAddress common.Address, operator common.Address, taskID uint64) (TaskResultInfo, error) {
	var out []interface{}
	err := _TriggerXAvs.contract.Call(opts, &out, "getOperatorTaskResponse", taskAddress, operator, taskID)

	if err != nil {
		return *new(TaskResultInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(TaskResultInfo)).(*TaskResultInfo)

	return out0, err

}

// GetOperatorTaskResponse is a free data retrieval call binding the contract method 0x16395dc4.
//
// Solidity: function getOperatorTaskResponse(address taskAddress, address operator, uint64 taskID) view returns((address,string,bytes,bytes,address,uint64,uint8))
func (_TriggerXAvs *TriggerXAvsSession) GetOperatorTaskResponse(taskAddress common.Address, operator common.Address, taskID uint64) (TaskResultInfo, error) {
	return _TriggerXAvs.Contract.GetOperatorTaskResponse(&_TriggerXAvs.CallOpts, taskAddress, operator, taskID)
}

// GetOperatorTaskResponse is a free data retrieval call binding the contract method 0x16395dc4.
//
// Solidity: function getOperatorTaskResponse(address taskAddress, address operator, uint64 taskID) view returns((address,string,bytes,bytes,address,uint64,uint8))
func (_TriggerXAvs *TriggerXAvsCallerSession) GetOperatorTaskResponse(taskAddress common.Address, operator common.Address, taskID uint64) (TaskResultInfo, error) {
	return _TriggerXAvs.Contract.GetOperatorTaskResponse(&_TriggerXAvs.CallOpts, taskAddress, operator, taskID)
}

// GetOperatorTaskResponseList is a free data retrieval call binding the contract method 0xb6f64d2a.
//
// Solidity: function getOperatorTaskResponseList(address taskAddress, uint64 taskID) view returns((address,uint64,address,string,bytes,bytes,uint256,uint8)[])
func (_TriggerXAvs *TriggerXAvsCaller) GetOperatorTaskResponseList(opts *bind.CallOpts, taskAddress common.Address, taskID uint64) ([]OperatorResInfo, error) {
	var out []interface{}
	err := _TriggerXAvs.contract.Call(opts, &out, "getOperatorTaskResponseList", taskAddress, taskID)

	if err != nil {
		return *new([]OperatorResInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]OperatorResInfo)).(*[]OperatorResInfo)

	return out0, err

}

// GetOperatorTaskResponseList is a free data retrieval call binding the contract method 0xb6f64d2a.
//
// Solidity: function getOperatorTaskResponseList(address taskAddress, uint64 taskID) view returns((address,uint64,address,string,bytes,bytes,uint256,uint8)[])
func (_TriggerXAvs *TriggerXAvsSession) GetOperatorTaskResponseList(taskAddress common.Address, taskID uint64) ([]OperatorResInfo, error) {
	return _TriggerXAvs.Contract.GetOperatorTaskResponseList(&_TriggerXAvs.CallOpts, taskAddress, taskID)
}

// GetOperatorTaskResponseList is a free data retrieval call binding the contract method 0xb6f64d2a.
//
// Solidity: function getOperatorTaskResponseList(address taskAddress, uint64 taskID) view returns((address,uint64,address,string,bytes,bytes,uint256,uint8)[])
func (_TriggerXAvs *TriggerXAvsCallerSession) GetOperatorTaskResponseList(taskAddress common.Address, taskID uint64) ([]OperatorResInfo, error) {
	return _TriggerXAvs.Contract.GetOperatorTaskResponseList(&_TriggerXAvs.CallOpts, taskAddress, taskID)
}

// GetOptInOperators is a free data retrieval call binding the contract method 0x1d4c8007.
//
// Solidity: function getOptInOperators(address avsAddress) view returns(address[])
func (_TriggerXAvs *TriggerXAvsCaller) GetOptInOperators(opts *bind.CallOpts, avsAddress common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _TriggerXAvs.contract.Call(opts, &out, "getOptInOperators", avsAddress)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOptInOperators is a free data retrieval call binding the contract method 0x1d4c8007.
//
// Solidity: function getOptInOperators(address avsAddress) view returns(address[])
func (_TriggerXAvs *TriggerXAvsSession) GetOptInOperators(avsAddress common.Address) ([]common.Address, error) {
	return _TriggerXAvs.Contract.GetOptInOperators(&_TriggerXAvs.CallOpts, avsAddress)
}

// GetOptInOperators is a free data retrieval call binding the contract method 0x1d4c8007.
//
// Solidity: function getOptInOperators(address avsAddress) view returns(address[])
func (_TriggerXAvs *TriggerXAvsCallerSession) GetOptInOperators(avsAddress common.Address) ([]common.Address, error) {
	return _TriggerXAvs.Contract.GetOptInOperators(&_TriggerXAvs.CallOpts, avsAddress)
}

// GetRegisteredPubkey is a free data retrieval call binding the contract method 0x9943aa27.
//
// Solidity: function getRegisteredPubkey(address operator, address avsAddr) view returns(bytes)
func (_TriggerXAvs *TriggerXAvsCaller) GetRegisteredPubkey(opts *bind.CallOpts, operator common.Address, avsAddr common.Address) ([]byte, error) {
	var out []interface{}
	err := _TriggerXAvs.contract.Call(opts, &out, "getRegisteredPubkey", operator, avsAddr)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetRegisteredPubkey is a free data retrieval call binding the contract method 0x9943aa27.
//
// Solidity: function getRegisteredPubkey(address operator, address avsAddr) view returns(bytes)
func (_TriggerXAvs *TriggerXAvsSession) GetRegisteredPubkey(operator common.Address, avsAddr common.Address) ([]byte, error) {
	return _TriggerXAvs.Contract.GetRegisteredPubkey(&_TriggerXAvs.CallOpts, operator, avsAddr)
}

// GetRegisteredPubkey is a free data retrieval call binding the contract method 0x9943aa27.
//
// Solidity: function getRegisteredPubkey(address operator, address avsAddr) view returns(bytes)
func (_TriggerXAvs *TriggerXAvsCallerSession) GetRegisteredPubkey(operator common.Address, avsAddr common.Address) ([]byte, error) {
	return _TriggerXAvs.Contract.GetRegisteredPubkey(&_TriggerXAvs.CallOpts, operator, avsAddr)
}

// GetTaskDefinition is a free data retrieval call binding the contract method 0xbd0d5737.
//
// Solidity: function getTaskDefinition(uint8 id) view returns((uint8,string,uint256,uint256,uint256,uint256,uint256[],uint256))
func (_TriggerXAvs *TriggerXAvsCaller) GetTaskDefinition(opts *bind.CallOpts, id uint8) (TaskDefinition, error) {
	var out []interface{}
	err := _TriggerXAvs.contract.Call(opts, &out, "getTaskDefinition", id)

	if err != nil {
		return *new(TaskDefinition), err
	}

	out0 := *abi.ConvertType(out[0], new(TaskDefinition)).(*TaskDefinition)

	return out0, err

}

// GetTaskDefinition is a free data retrieval call binding the contract method 0xbd0d5737.
//
// Solidity: function getTaskDefinition(uint8 id) view returns((uint8,string,uint256,uint256,uint256,uint256,uint256[],uint256))
func (_TriggerXAvs *TriggerXAvsSession) GetTaskDefinition(id uint8) (TaskDefinition, error) {
	return _TriggerXAvs.Contract.GetTaskDefinition(&_TriggerXAvs.CallOpts, id)
}

// GetTaskDefinition is a free data retrieval call binding the contract method 0xbd0d5737.
//
// Solidity: function getTaskDefinition(uint8 id) view returns((uint8,string,uint256,uint256,uint256,uint256,uint256[],uint256))
func (_TriggerXAvs *TriggerXAvsCallerSession) GetTaskDefinition(id uint8) (TaskDefinition, error) {
	return _TriggerXAvs.Contract.GetTaskDefinition(&_TriggerXAvs.CallOpts, id)
}

// GetTaskInfo is a free data retrieval call binding the contract method 0xe2906f3d.
//
// Solidity: function getTaskInfo(address taskAddress, uint64 taskID) view returns((address,string,bytes,uint64,uint64,uint64,uint64,uint8,uint64,string,address[],address[],address[],address[],string,(address,uint256)[],bool,address[],address[]))
func (_TriggerXAvs *TriggerXAvsCaller) GetTaskInfo(opts *bind.CallOpts, taskAddress common.Address, taskID uint64) (TaskInfo, error) {
	var out []interface{}
	err := _TriggerXAvs.contract.Call(opts, &out, "getTaskInfo", taskAddress, taskID)

	if err != nil {
		return *new(TaskInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(TaskInfo)).(*TaskInfo)

	return out0, err

}

// GetTaskInfo is a free data retrieval call binding the contract method 0xe2906f3d.
//
// Solidity: function getTaskInfo(address taskAddress, uint64 taskID) view returns((address,string,bytes,uint64,uint64,uint64,uint64,uint8,uint64,string,address[],address[],address[],address[],string,(address,uint256)[],bool,address[],address[]))
func (_TriggerXAvs *TriggerXAvsSession) GetTaskInfo(taskAddress common.Address, taskID uint64) (TaskInfo, error) {
	return _TriggerXAvs.Contract.GetTaskInfo(&_TriggerXAvs.CallOpts, taskAddress, taskID)
}

// GetTaskInfo is a free data retrieval call binding the contract method 0xe2906f3d.
//
// Solidity: function getTaskInfo(address taskAddress, uint64 taskID) view returns((address,string,bytes,uint64,uint64,uint64,uint64,uint8,uint64,string,address[],address[],address[],address[],string,(address,uint256)[],bool,address[],address[]))
func (_TriggerXAvs *TriggerXAvsCallerSession) GetTaskInfo(taskAddress common.Address, taskID uint64) (TaskInfo, error) {
	return _TriggerXAvs.Contract.GetTaskInfo(&_TriggerXAvs.CallOpts, taskAddress, taskID)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address operator) view returns(bool)
func (_TriggerXAvs *TriggerXAvsCaller) IsOperator(opts *bind.CallOpts, operator common.Address) (bool, error) {
	var out []interface{}
	err := _TriggerXAvs.contract.Call(opts, &out, "isOperator", operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address operator) view returns(bool)
func (_TriggerXAvs *TriggerXAvsSession) IsOperator(operator common.Address) (bool, error) {
	return _TriggerXAvs.Contract.IsOperator(&_TriggerXAvs.CallOpts, operator)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address operator) view returns(bool)
func (_TriggerXAvs *TriggerXAvsCallerSession) IsOperator(operator common.Address) (bool, error) {
	return _TriggerXAvs.Contract.IsOperator(&_TriggerXAvs.CallOpts, operator)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TriggerXAvs *TriggerXAvsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TriggerXAvs.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TriggerXAvs *TriggerXAvsSession) Owner() (common.Address, error) {
	return _TriggerXAvs.Contract.Owner(&_TriggerXAvs.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TriggerXAvs *TriggerXAvsCallerSession) Owner() (common.Address, error) {
	return _TriggerXAvs.Contract.Owner(&_TriggerXAvs.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TriggerXAvs *TriggerXAvsCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TriggerXAvs.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TriggerXAvs *TriggerXAvsSession) ProxiableUUID() ([32]byte, error) {
	return _TriggerXAvs.Contract.ProxiableUUID(&_TriggerXAvs.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TriggerXAvs *TriggerXAvsCallerSession) ProxiableUUID() ([32]byte, error) {
	return _TriggerXAvs.Contract.ProxiableUUID(&_TriggerXAvs.CallOpts)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_TriggerXAvs *TriggerXAvsCaller) RewardManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TriggerXAvs.contract.Call(opts, &out, "rewardManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_TriggerXAvs *TriggerXAvsSession) RewardManager() (common.Address, error) {
	return _TriggerXAvs.Contract.RewardManager(&_TriggerXAvs.CallOpts)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_TriggerXAvs *TriggerXAvsCallerSession) RewardManager() (common.Address, error) {
	return _TriggerXAvs.Contract.RewardManager(&_TriggerXAvs.CallOpts)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_TriggerXAvs *TriggerXAvsCaller) Slasher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TriggerXAvs.contract.Call(opts, &out, "slasher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_TriggerXAvs *TriggerXAvsSession) Slasher() (common.Address, error) {
	return _TriggerXAvs.Contract.Slasher(&_TriggerXAvs.CallOpts)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_TriggerXAvs *TriggerXAvsCallerSession) Slasher() (common.Address, error) {
	return _TriggerXAvs.Contract.Slasher(&_TriggerXAvs.CallOpts)
}

// Challenge is a paid mutator transaction binding the contract method 0xe20529ce.
//
// Solidity: function challenge(uint64 taskID, uint8 actualThreshold, bool isExpected, address[] eligibleRewardOperators, address[] eligibleSlashOperators) returns(bool success)
func (_TriggerXAvs *TriggerXAvsTransactor) Challenge(opts *bind.TransactOpts, taskID uint64, actualThreshold uint8, isExpected bool, eligibleRewardOperators []common.Address, eligibleSlashOperators []common.Address) (*types.Transaction, error) {
	return _TriggerXAvs.contract.Transact(opts, "challenge", taskID, actualThreshold, isExpected, eligibleRewardOperators, eligibleSlashOperators)
}

// Challenge is a paid mutator transaction binding the contract method 0xe20529ce.
//
// Solidity: function challenge(uint64 taskID, uint8 actualThreshold, bool isExpected, address[] eligibleRewardOperators, address[] eligibleSlashOperators) returns(bool success)
func (_TriggerXAvs *TriggerXAvsSession) Challenge(taskID uint64, actualThreshold uint8, isExpected bool, eligibleRewardOperators []common.Address, eligibleSlashOperators []common.Address) (*types.Transaction, error) {
	return _TriggerXAvs.Contract.Challenge(&_TriggerXAvs.TransactOpts, taskID, actualThreshold, isExpected, eligibleRewardOperators, eligibleSlashOperators)
}

// Challenge is a paid mutator transaction binding the contract method 0xe20529ce.
//
// Solidity: function challenge(uint64 taskID, uint8 actualThreshold, bool isExpected, address[] eligibleRewardOperators, address[] eligibleSlashOperators) returns(bool success)
func (_TriggerXAvs *TriggerXAvsTransactorSession) Challenge(taskID uint64, actualThreshold uint8, isExpected bool, eligibleRewardOperators []common.Address, eligibleSlashOperators []common.Address) (*types.Transaction, error) {
	return _TriggerXAvs.Contract.Challenge(&_TriggerXAvs.TransactOpts, taskID, actualThreshold, isExpected, eligibleRewardOperators, eligibleSlashOperators)
}

// CreateTask is a paid mutator transaction binding the contract method 0x2563adf9.
//
// Solidity: function createTask(string taskName, uint8 taskDefinitionId, uint64 taskResponsePeriod, uint64 taskChallengePeriod, uint8 thresholdPercentage, uint64 taskStatisticalPeriod) returns(uint64 taskID)
func (_TriggerXAvs *TriggerXAvsTransactor) CreateTask(opts *bind.TransactOpts, taskName string, taskDefinitionId uint8, taskResponsePeriod uint64, taskChallengePeriod uint64, thresholdPercentage uint8, taskStatisticalPeriod uint64) (*types.Transaction, error) {
	return _TriggerXAvs.contract.Transact(opts, "createTask", taskName, taskDefinitionId, taskResponsePeriod, taskChallengePeriod, thresholdPercentage, taskStatisticalPeriod)
}

// CreateTask is a paid mutator transaction binding the contract method 0x2563adf9.
//
// Solidity: function createTask(string taskName, uint8 taskDefinitionId, uint64 taskResponsePeriod, uint64 taskChallengePeriod, uint8 thresholdPercentage, uint64 taskStatisticalPeriod) returns(uint64 taskID)
func (_TriggerXAvs *TriggerXAvsSession) CreateTask(taskName string, taskDefinitionId uint8, taskResponsePeriod uint64, taskChallengePeriod uint64, thresholdPercentage uint8, taskStatisticalPeriod uint64) (*types.Transaction, error) {
	return _TriggerXAvs.Contract.CreateTask(&_TriggerXAvs.TransactOpts, taskName, taskDefinitionId, taskResponsePeriod, taskChallengePeriod, thresholdPercentage, taskStatisticalPeriod)
}

// CreateTask is a paid mutator transaction binding the contract method 0x2563adf9.
//
// Solidity: function createTask(string taskName, uint8 taskDefinitionId, uint64 taskResponsePeriod, uint64 taskChallengePeriod, uint8 thresholdPercentage, uint64 taskStatisticalPeriod) returns(uint64 taskID)
func (_TriggerXAvs *TriggerXAvsTransactorSession) CreateTask(taskName string, taskDefinitionId uint8, taskResponsePeriod uint64, taskChallengePeriod uint64, thresholdPercentage uint8, taskStatisticalPeriod uint64) (*types.Transaction, error) {
	return _TriggerXAvs.Contract.CreateTask(&_TriggerXAvs.TransactOpts, taskName, taskDefinitionId, taskResponsePeriod, taskChallengePeriod, thresholdPercentage, taskStatisticalPeriod)
}

// CreateTaskDefinition is a paid mutator transaction binding the contract method 0x459e1340.
//
// Solidity: function createTaskDefinition(string name, (uint256,uint256,uint256,uint256,uint256[],uint256) defParams) returns(uint8 taskDefinitionId)
func (_TriggerXAvs *TriggerXAvsTransactor) CreateTaskDefinition(opts *bind.TransactOpts, name string, defParams TaskDefinitionParams) (*types.Transaction, error) {
	return _TriggerXAvs.contract.Transact(opts, "createTaskDefinition", name, defParams)
}

// CreateTaskDefinition is a paid mutator transaction binding the contract method 0x459e1340.
//
// Solidity: function createTaskDefinition(string name, (uint256,uint256,uint256,uint256,uint256[],uint256) defParams) returns(uint8 taskDefinitionId)
func (_TriggerXAvs *TriggerXAvsSession) CreateTaskDefinition(name string, defParams TaskDefinitionParams) (*types.Transaction, error) {
	return _TriggerXAvs.Contract.CreateTaskDefinition(&_TriggerXAvs.TransactOpts, name, defParams)
}

// CreateTaskDefinition is a paid mutator transaction binding the contract method 0x459e1340.
//
// Solidity: function createTaskDefinition(string name, (uint256,uint256,uint256,uint256,uint256[],uint256) defParams) returns(uint8 taskDefinitionId)
func (_TriggerXAvs *TriggerXAvsTransactorSession) CreateTaskDefinition(name string, defParams TaskDefinitionParams) (*types.Transaction, error) {
	return _TriggerXAvs.Contract.CreateTaskDefinition(&_TriggerXAvs.TransactOpts, name, defParams)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xde16bf46.
//
// Solidity: function deregisterOperatorFromAVS() returns(bool success)
func (_TriggerXAvs *TriggerXAvsTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TriggerXAvs.contract.Transact(opts, "deregisterOperatorFromAVS")
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xde16bf46.
//
// Solidity: function deregisterOperatorFromAVS() returns(bool success)
func (_TriggerXAvs *TriggerXAvsSession) DeregisterOperatorFromAVS() (*types.Transaction, error) {
	return _TriggerXAvs.Contract.DeregisterOperatorFromAVS(&_TriggerXAvs.TransactOpts)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xde16bf46.
//
// Solidity: function deregisterOperatorFromAVS() returns(bool success)
func (_TriggerXAvs *TriggerXAvsTransactorSession) DeregisterOperatorFromAVS() (*types.Transaction, error) {
	return _TriggerXAvs.Contract.DeregisterOperatorFromAVS(&_TriggerXAvs.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_TriggerXAvs *TriggerXAvsTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address) (*types.Transaction, error) {
	return _TriggerXAvs.contract.Transact(opts, "initialize", initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_TriggerXAvs *TriggerXAvsSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _TriggerXAvs.Contract.Initialize(&_TriggerXAvs.TransactOpts, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_TriggerXAvs *TriggerXAvsTransactorSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _TriggerXAvs.Contract.Initialize(&_TriggerXAvs.TransactOpts, initialOwner)
}

// OperatorSubmitTask is a paid mutator transaction binding the contract method 0x26135e9d.
//
// Solidity: function operatorSubmitTask(uint64 taskID, bytes taskResponse, bytes blsSignature, address taskContractAddress, uint8 phase) returns(bool success)
func (_TriggerXAvs *TriggerXAvsTransactor) OperatorSubmitTask(opts *bind.TransactOpts, taskID uint64, taskResponse []byte, blsSignature []byte, taskContractAddress common.Address, phase uint8) (*types.Transaction, error) {
	return _TriggerXAvs.contract.Transact(opts, "operatorSubmitTask", taskID, taskResponse, blsSignature, taskContractAddress, phase)
}

// OperatorSubmitTask is a paid mutator transaction binding the contract method 0x26135e9d.
//
// Solidity: function operatorSubmitTask(uint64 taskID, bytes taskResponse, bytes blsSignature, address taskContractAddress, uint8 phase) returns(bool success)
func (_TriggerXAvs *TriggerXAvsSession) OperatorSubmitTask(taskID uint64, taskResponse []byte, blsSignature []byte, taskContractAddress common.Address, phase uint8) (*types.Transaction, error) {
	return _TriggerXAvs.Contract.OperatorSubmitTask(&_TriggerXAvs.TransactOpts, taskID, taskResponse, blsSignature, taskContractAddress, phase)
}

// OperatorSubmitTask is a paid mutator transaction binding the contract method 0x26135e9d.
//
// Solidity: function operatorSubmitTask(uint64 taskID, bytes taskResponse, bytes blsSignature, address taskContractAddress, uint8 phase) returns(bool success)
func (_TriggerXAvs *TriggerXAvsTransactorSession) OperatorSubmitTask(taskID uint64, taskResponse []byte, blsSignature []byte, taskContractAddress common.Address, phase uint8) (*types.Transaction, error) {
	return _TriggerXAvs.Contract.OperatorSubmitTask(&_TriggerXAvs.TransactOpts, taskID, taskResponse, blsSignature, taskContractAddress, phase)
}

// RegisterAVS is a paid mutator transaction binding the contract method 0x0b70f322.
//
// Solidity: function registerAVS((address,string,uint64,address,address,address,address[],address[],string[],uint64,uint64,string,uint64,uint64,uint64,uint64) params) returns(bool success)
func (_TriggerXAvs *TriggerXAvsTransactor) RegisterAVS(opts *bind.TransactOpts, params AVSParams) (*types.Transaction, error) {
	return _TriggerXAvs.contract.Transact(opts, "registerAVS", params)
}

// RegisterAVS is a paid mutator transaction binding the contract method 0x0b70f322.
//
// Solidity: function registerAVS((address,string,uint64,address,address,address,address[],address[],string[],uint64,uint64,string,uint64,uint64,uint64,uint64) params) returns(bool success)
func (_TriggerXAvs *TriggerXAvsSession) RegisterAVS(params AVSParams) (*types.Transaction, error) {
	return _TriggerXAvs.Contract.RegisterAVS(&_TriggerXAvs.TransactOpts, params)
}

// RegisterAVS is a paid mutator transaction binding the contract method 0x0b70f322.
//
// Solidity: function registerAVS((address,string,uint64,address,address,address,address[],address[],string[],uint64,uint64,string,uint64,uint64,uint64,uint64) params) returns(bool success)
func (_TriggerXAvs *TriggerXAvsTransactorSession) RegisterAVS(params AVSParams) (*types.Transaction, error) {
	return _TriggerXAvs.Contract.RegisterAVS(&_TriggerXAvs.TransactOpts, params)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0x5d9e941f.
//
// Solidity: function registerBLSPublicKey(address avsAddr, bytes pubKey, bytes pubKeyRegistrationSignature) returns(bool success)
func (_TriggerXAvs *TriggerXAvsTransactor) RegisterBLSPublicKey(opts *bind.TransactOpts, avsAddr common.Address, pubKey []byte, pubKeyRegistrationSignature []byte) (*types.Transaction, error) {
	return _TriggerXAvs.contract.Transact(opts, "registerBLSPublicKey", avsAddr, pubKey, pubKeyRegistrationSignature)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0x5d9e941f.
//
// Solidity: function registerBLSPublicKey(address avsAddr, bytes pubKey, bytes pubKeyRegistrationSignature) returns(bool success)
func (_TriggerXAvs *TriggerXAvsSession) RegisterBLSPublicKey(avsAddr common.Address, pubKey []byte, pubKeyRegistrationSignature []byte) (*types.Transaction, error) {
	return _TriggerXAvs.Contract.RegisterBLSPublicKey(&_TriggerXAvs.TransactOpts, avsAddr, pubKey, pubKeyRegistrationSignature)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0x5d9e941f.
//
// Solidity: function registerBLSPublicKey(address avsAddr, bytes pubKey, bytes pubKeyRegistrationSignature) returns(bool success)
func (_TriggerXAvs *TriggerXAvsTransactorSession) RegisterBLSPublicKey(avsAddr common.Address, pubKey []byte, pubKeyRegistrationSignature []byte) (*types.Transaction, error) {
	return _TriggerXAvs.Contract.RegisterBLSPublicKey(&_TriggerXAvs.TransactOpts, avsAddr, pubKey, pubKeyRegistrationSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0xc208dd99.
//
// Solidity: function registerOperatorToAVS() returns(bool success)
func (_TriggerXAvs *TriggerXAvsTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TriggerXAvs.contract.Transact(opts, "registerOperatorToAVS")
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0xc208dd99.
//
// Solidity: function registerOperatorToAVS() returns(bool success)
func (_TriggerXAvs *TriggerXAvsSession) RegisterOperatorToAVS() (*types.Transaction, error) {
	return _TriggerXAvs.Contract.RegisterOperatorToAVS(&_TriggerXAvs.TransactOpts)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0xc208dd99.
//
// Solidity: function registerOperatorToAVS() returns(bool success)
func (_TriggerXAvs *TriggerXAvsTransactorSession) RegisterOperatorToAVS() (*types.Transaction, error) {
	return _TriggerXAvs.Contract.RegisterOperatorToAVS(&_TriggerXAvs.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TriggerXAvs *TriggerXAvsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TriggerXAvs.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TriggerXAvs *TriggerXAvsSession) RenounceOwnership() (*types.Transaction, error) {
	return _TriggerXAvs.Contract.RenounceOwnership(&_TriggerXAvs.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TriggerXAvs *TriggerXAvsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TriggerXAvs.Contract.RenounceOwnership(&_TriggerXAvs.TransactOpts)
}

// SetRewardManager is a paid mutator transaction binding the contract method 0x153ee554.
//
// Solidity: function setRewardManager(address _rewardManager) returns()
func (_TriggerXAvs *TriggerXAvsTransactor) SetRewardManager(opts *bind.TransactOpts, _rewardManager common.Address) (*types.Transaction, error) {
	return _TriggerXAvs.contract.Transact(opts, "setRewardManager", _rewardManager)
}

// SetRewardManager is a paid mutator transaction binding the contract method 0x153ee554.
//
// Solidity: function setRewardManager(address _rewardManager) returns()
func (_TriggerXAvs *TriggerXAvsSession) SetRewardManager(_rewardManager common.Address) (*types.Transaction, error) {
	return _TriggerXAvs.Contract.SetRewardManager(&_TriggerXAvs.TransactOpts, _rewardManager)
}

// SetRewardManager is a paid mutator transaction binding the contract method 0x153ee554.
//
// Solidity: function setRewardManager(address _rewardManager) returns()
func (_TriggerXAvs *TriggerXAvsTransactorSession) SetRewardManager(_rewardManager common.Address) (*types.Transaction, error) {
	return _TriggerXAvs.Contract.SetRewardManager(&_TriggerXAvs.TransactOpts, _rewardManager)
}

// SetSlasher is a paid mutator transaction binding the contract method 0xaabc2496.
//
// Solidity: function setSlasher(address _slasher) returns()
func (_TriggerXAvs *TriggerXAvsTransactor) SetSlasher(opts *bind.TransactOpts, _slasher common.Address) (*types.Transaction, error) {
	return _TriggerXAvs.contract.Transact(opts, "setSlasher", _slasher)
}

// SetSlasher is a paid mutator transaction binding the contract method 0xaabc2496.
//
// Solidity: function setSlasher(address _slasher) returns()
func (_TriggerXAvs *TriggerXAvsSession) SetSlasher(_slasher common.Address) (*types.Transaction, error) {
	return _TriggerXAvs.Contract.SetSlasher(&_TriggerXAvs.TransactOpts, _slasher)
}

// SetSlasher is a paid mutator transaction binding the contract method 0xaabc2496.
//
// Solidity: function setSlasher(address _slasher) returns()
func (_TriggerXAvs *TriggerXAvsTransactorSession) SetSlasher(_slasher common.Address) (*types.Transaction, error) {
	return _TriggerXAvs.Contract.SetSlasher(&_TriggerXAvs.TransactOpts, _slasher)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TriggerXAvs *TriggerXAvsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TriggerXAvs.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TriggerXAvs *TriggerXAvsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TriggerXAvs.Contract.TransferOwnership(&_TriggerXAvs.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TriggerXAvs *TriggerXAvsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TriggerXAvs.Contract.TransferOwnership(&_TriggerXAvs.TransactOpts, newOwner)
}

// UpdateAVS is a paid mutator transaction binding the contract method 0x3a72b900.
//
// Solidity: function updateAVS((address,string,uint64,address,address,address,address[],address[],string[],uint64,uint64,string,uint64,uint64,uint64,uint64) params) returns(bool success)
func (_TriggerXAvs *TriggerXAvsTransactor) UpdateAVS(opts *bind.TransactOpts, params AVSParams) (*types.Transaction, error) {
	return _TriggerXAvs.contract.Transact(opts, "updateAVS", params)
}

// UpdateAVS is a paid mutator transaction binding the contract method 0x3a72b900.
//
// Solidity: function updateAVS((address,string,uint64,address,address,address,address[],address[],string[],uint64,uint64,string,uint64,uint64,uint64,uint64) params) returns(bool success)
func (_TriggerXAvs *TriggerXAvsSession) UpdateAVS(params AVSParams) (*types.Transaction, error) {
	return _TriggerXAvs.Contract.UpdateAVS(&_TriggerXAvs.TransactOpts, params)
}

// UpdateAVS is a paid mutator transaction binding the contract method 0x3a72b900.
//
// Solidity: function updateAVS((address,string,uint64,address,address,address,address[],address[],string[],uint64,uint64,string,uint64,uint64,uint64,uint64) params) returns(bool success)
func (_TriggerXAvs *TriggerXAvsTransactorSession) UpdateAVS(params AVSParams) (*types.Transaction, error) {
	return _TriggerXAvs.Contract.UpdateAVS(&_TriggerXAvs.TransactOpts, params)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_TriggerXAvs *TriggerXAvsTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _TriggerXAvs.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_TriggerXAvs *TriggerXAvsSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _TriggerXAvs.Contract.UpgradeToAndCall(&_TriggerXAvs.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_TriggerXAvs *TriggerXAvsTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _TriggerXAvs.Contract.UpgradeToAndCall(&_TriggerXAvs.TransactOpts, newImplementation, data)
}

// TriggerXAvsBLSPublicKeyRegisteredIterator is returned from FilterBLSPublicKeyRegistered and is used to iterate over the raw logs and unpacked data for BLSPublicKeyRegistered events raised by the TriggerXAvs contract.
type TriggerXAvsBLSPublicKeyRegisteredIterator struct {
	Event *TriggerXAvsBLSPublicKeyRegistered // Event containing the contract specifics and raw log

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
func (it *TriggerXAvsBLSPublicKeyRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TriggerXAvsBLSPublicKeyRegistered)
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
		it.Event = new(TriggerXAvsBLSPublicKeyRegistered)
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
func (it *TriggerXAvsBLSPublicKeyRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TriggerXAvsBLSPublicKeyRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TriggerXAvsBLSPublicKeyRegistered represents a BLSPublicKeyRegistered event raised by the TriggerXAvs contract.
type TriggerXAvsBLSPublicKeyRegistered struct {
	Operator   common.Address
	AvsAddress common.Address
	PubKey     []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBLSPublicKeyRegistered is a free log retrieval operation binding the contract event 0xd7a56e64a4cd1aeb35e575da573ffdbd29cbafdf2ef88c1772197d7c72be405f.
//
// Solidity: event BLSPublicKeyRegistered(address indexed operator, address indexed avsAddress, bytes pubKey)
func (_TriggerXAvs *TriggerXAvsFilterer) FilterBLSPublicKeyRegistered(opts *bind.FilterOpts, operator []common.Address, avsAddress []common.Address) (*TriggerXAvsBLSPublicKeyRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsAddressRule []interface{}
	for _, avsAddressItem := range avsAddress {
		avsAddressRule = append(avsAddressRule, avsAddressItem)
	}

	logs, sub, err := _TriggerXAvs.contract.FilterLogs(opts, "BLSPublicKeyRegistered", operatorRule, avsAddressRule)
	if err != nil {
		return nil, err
	}
	return &TriggerXAvsBLSPublicKeyRegisteredIterator{contract: _TriggerXAvs.contract, event: "BLSPublicKeyRegistered", logs: logs, sub: sub}, nil
}

// WatchBLSPublicKeyRegistered is a free log subscription operation binding the contract event 0xd7a56e64a4cd1aeb35e575da573ffdbd29cbafdf2ef88c1772197d7c72be405f.
//
// Solidity: event BLSPublicKeyRegistered(address indexed operator, address indexed avsAddress, bytes pubKey)
func (_TriggerXAvs *TriggerXAvsFilterer) WatchBLSPublicKeyRegistered(opts *bind.WatchOpts, sink chan<- *TriggerXAvsBLSPublicKeyRegistered, operator []common.Address, avsAddress []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsAddressRule []interface{}
	for _, avsAddressItem := range avsAddress {
		avsAddressRule = append(avsAddressRule, avsAddressItem)
	}

	logs, sub, err := _TriggerXAvs.contract.WatchLogs(opts, "BLSPublicKeyRegistered", operatorRule, avsAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TriggerXAvsBLSPublicKeyRegistered)
				if err := _TriggerXAvs.contract.UnpackLog(event, "BLSPublicKeyRegistered", log); err != nil {
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

// ParseBLSPublicKeyRegistered is a log parse operation binding the contract event 0xd7a56e64a4cd1aeb35e575da573ffdbd29cbafdf2ef88c1772197d7c72be405f.
//
// Solidity: event BLSPublicKeyRegistered(address indexed operator, address indexed avsAddress, bytes pubKey)
func (_TriggerXAvs *TriggerXAvsFilterer) ParseBLSPublicKeyRegistered(log types.Log) (*TriggerXAvsBLSPublicKeyRegistered, error) {
	event := new(TriggerXAvsBLSPublicKeyRegistered)
	if err := _TriggerXAvs.contract.UnpackLog(event, "BLSPublicKeyRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TriggerXAvsChallengeSubmittedIterator is returned from FilterChallengeSubmitted and is used to iterate over the raw logs and unpacked data for ChallengeSubmitted events raised by the TriggerXAvs contract.
type TriggerXAvsChallengeSubmittedIterator struct {
	Event *TriggerXAvsChallengeSubmitted // Event containing the contract specifics and raw log

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
func (it *TriggerXAvsChallengeSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TriggerXAvsChallengeSubmitted)
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
		it.Event = new(TriggerXAvsChallengeSubmitted)
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
func (it *TriggerXAvsChallengeSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TriggerXAvsChallengeSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TriggerXAvsChallengeSubmitted represents a ChallengeSubmitted event raised by the TriggerXAvs contract.
type TriggerXAvsChallengeSubmitted struct {
	TaskID     uint64
	Challenger common.Address
	IsExpected bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterChallengeSubmitted is a free log retrieval operation binding the contract event 0x9de5978cdd4b224702e27e43ab9f28631c0ab8ffe4b20e4cd8301d7a6da8ae72.
//
// Solidity: event ChallengeSubmitted(uint64 indexed taskID, address indexed challenger, bool isExpected)
func (_TriggerXAvs *TriggerXAvsFilterer) FilterChallengeSubmitted(opts *bind.FilterOpts, taskID []uint64, challenger []common.Address) (*TriggerXAvsChallengeSubmittedIterator, error) {

	var taskIDRule []interface{}
	for _, taskIDItem := range taskID {
		taskIDRule = append(taskIDRule, taskIDItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _TriggerXAvs.contract.FilterLogs(opts, "ChallengeSubmitted", taskIDRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &TriggerXAvsChallengeSubmittedIterator{contract: _TriggerXAvs.contract, event: "ChallengeSubmitted", logs: logs, sub: sub}, nil
}

// WatchChallengeSubmitted is a free log subscription operation binding the contract event 0x9de5978cdd4b224702e27e43ab9f28631c0ab8ffe4b20e4cd8301d7a6da8ae72.
//
// Solidity: event ChallengeSubmitted(uint64 indexed taskID, address indexed challenger, bool isExpected)
func (_TriggerXAvs *TriggerXAvsFilterer) WatchChallengeSubmitted(opts *bind.WatchOpts, sink chan<- *TriggerXAvsChallengeSubmitted, taskID []uint64, challenger []common.Address) (event.Subscription, error) {

	var taskIDRule []interface{}
	for _, taskIDItem := range taskID {
		taskIDRule = append(taskIDRule, taskIDItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _TriggerXAvs.contract.WatchLogs(opts, "ChallengeSubmitted", taskIDRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TriggerXAvsChallengeSubmitted)
				if err := _TriggerXAvs.contract.UnpackLog(event, "ChallengeSubmitted", log); err != nil {
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

// ParseChallengeSubmitted is a log parse operation binding the contract event 0x9de5978cdd4b224702e27e43ab9f28631c0ab8ffe4b20e4cd8301d7a6da8ae72.
//
// Solidity: event ChallengeSubmitted(uint64 indexed taskID, address indexed challenger, bool isExpected)
func (_TriggerXAvs *TriggerXAvsFilterer) ParseChallengeSubmitted(log types.Log) (*TriggerXAvsChallengeSubmitted, error) {
	event := new(TriggerXAvsChallengeSubmitted)
	if err := _TriggerXAvs.contract.UnpackLog(event, "ChallengeSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TriggerXAvsInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TriggerXAvs contract.
type TriggerXAvsInitializedIterator struct {
	Event *TriggerXAvsInitialized // Event containing the contract specifics and raw log

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
func (it *TriggerXAvsInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TriggerXAvsInitialized)
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
		it.Event = new(TriggerXAvsInitialized)
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
func (it *TriggerXAvsInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TriggerXAvsInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TriggerXAvsInitialized represents a Initialized event raised by the TriggerXAvs contract.
type TriggerXAvsInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TriggerXAvs *TriggerXAvsFilterer) FilterInitialized(opts *bind.FilterOpts) (*TriggerXAvsInitializedIterator, error) {

	logs, sub, err := _TriggerXAvs.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TriggerXAvsInitializedIterator{contract: _TriggerXAvs.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TriggerXAvs *TriggerXAvsFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TriggerXAvsInitialized) (event.Subscription, error) {

	logs, sub, err := _TriggerXAvs.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TriggerXAvsInitialized)
				if err := _TriggerXAvs.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_TriggerXAvs *TriggerXAvsFilterer) ParseInitialized(log types.Log) (*TriggerXAvsInitialized, error) {
	event := new(TriggerXAvsInitialized)
	if err := _TriggerXAvs.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TriggerXAvsOperatorOptedInIterator is returned from FilterOperatorOptedIn and is used to iterate over the raw logs and unpacked data for OperatorOptedIn events raised by the TriggerXAvs contract.
type TriggerXAvsOperatorOptedInIterator struct {
	Event *TriggerXAvsOperatorOptedIn // Event containing the contract specifics and raw log

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
func (it *TriggerXAvsOperatorOptedInIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TriggerXAvsOperatorOptedIn)
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
		it.Event = new(TriggerXAvsOperatorOptedIn)
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
func (it *TriggerXAvsOperatorOptedInIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TriggerXAvsOperatorOptedInIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TriggerXAvsOperatorOptedIn represents a OperatorOptedIn event raised by the TriggerXAvs contract.
type TriggerXAvsOperatorOptedIn struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorOptedIn is a free log retrieval operation binding the contract event 0x3eb9749f7e08a89f04537594ff4cf7885d8f0b56138c8cfa999b44771fb1a148.
//
// Solidity: event OperatorOptedIn(address indexed operator)
func (_TriggerXAvs *TriggerXAvsFilterer) FilterOperatorOptedIn(opts *bind.FilterOpts, operator []common.Address) (*TriggerXAvsOperatorOptedInIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TriggerXAvs.contract.FilterLogs(opts, "OperatorOptedIn", operatorRule)
	if err != nil {
		return nil, err
	}
	return &TriggerXAvsOperatorOptedInIterator{contract: _TriggerXAvs.contract, event: "OperatorOptedIn", logs: logs, sub: sub}, nil
}

// WatchOperatorOptedIn is a free log subscription operation binding the contract event 0x3eb9749f7e08a89f04537594ff4cf7885d8f0b56138c8cfa999b44771fb1a148.
//
// Solidity: event OperatorOptedIn(address indexed operator)
func (_TriggerXAvs *TriggerXAvsFilterer) WatchOperatorOptedIn(opts *bind.WatchOpts, sink chan<- *TriggerXAvsOperatorOptedIn, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TriggerXAvs.contract.WatchLogs(opts, "OperatorOptedIn", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TriggerXAvsOperatorOptedIn)
				if err := _TriggerXAvs.contract.UnpackLog(event, "OperatorOptedIn", log); err != nil {
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

// ParseOperatorOptedIn is a log parse operation binding the contract event 0x3eb9749f7e08a89f04537594ff4cf7885d8f0b56138c8cfa999b44771fb1a148.
//
// Solidity: event OperatorOptedIn(address indexed operator)
func (_TriggerXAvs *TriggerXAvsFilterer) ParseOperatorOptedIn(log types.Log) (*TriggerXAvsOperatorOptedIn, error) {
	event := new(TriggerXAvsOperatorOptedIn)
	if err := _TriggerXAvs.contract.UnpackLog(event, "OperatorOptedIn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TriggerXAvsOperatorOptedOutIterator is returned from FilterOperatorOptedOut and is used to iterate over the raw logs and unpacked data for OperatorOptedOut events raised by the TriggerXAvs contract.
type TriggerXAvsOperatorOptedOutIterator struct {
	Event *TriggerXAvsOperatorOptedOut // Event containing the contract specifics and raw log

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
func (it *TriggerXAvsOperatorOptedOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TriggerXAvsOperatorOptedOut)
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
		it.Event = new(TriggerXAvsOperatorOptedOut)
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
func (it *TriggerXAvsOperatorOptedOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TriggerXAvsOperatorOptedOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TriggerXAvsOperatorOptedOut represents a OperatorOptedOut event raised by the TriggerXAvs contract.
type TriggerXAvsOperatorOptedOut struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorOptedOut is a free log retrieval operation binding the contract event 0x70245217442e9b16c38a95d39c7516c388c244266dd7c91b7bd8dbe285adf1a2.
//
// Solidity: event OperatorOptedOut(address indexed operator)
func (_TriggerXAvs *TriggerXAvsFilterer) FilterOperatorOptedOut(opts *bind.FilterOpts, operator []common.Address) (*TriggerXAvsOperatorOptedOutIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TriggerXAvs.contract.FilterLogs(opts, "OperatorOptedOut", operatorRule)
	if err != nil {
		return nil, err
	}
	return &TriggerXAvsOperatorOptedOutIterator{contract: _TriggerXAvs.contract, event: "OperatorOptedOut", logs: logs, sub: sub}, nil
}

// WatchOperatorOptedOut is a free log subscription operation binding the contract event 0x70245217442e9b16c38a95d39c7516c388c244266dd7c91b7bd8dbe285adf1a2.
//
// Solidity: event OperatorOptedOut(address indexed operator)
func (_TriggerXAvs *TriggerXAvsFilterer) WatchOperatorOptedOut(opts *bind.WatchOpts, sink chan<- *TriggerXAvsOperatorOptedOut, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TriggerXAvs.contract.WatchLogs(opts, "OperatorOptedOut", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TriggerXAvsOperatorOptedOut)
				if err := _TriggerXAvs.contract.UnpackLog(event, "OperatorOptedOut", log); err != nil {
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

// ParseOperatorOptedOut is a log parse operation binding the contract event 0x70245217442e9b16c38a95d39c7516c388c244266dd7c91b7bd8dbe285adf1a2.
//
// Solidity: event OperatorOptedOut(address indexed operator)
func (_TriggerXAvs *TriggerXAvsFilterer) ParseOperatorOptedOut(log types.Log) (*TriggerXAvsOperatorOptedOut, error) {
	event := new(TriggerXAvsOperatorOptedOut)
	if err := _TriggerXAvs.contract.UnpackLog(event, "OperatorOptedOut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TriggerXAvsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TriggerXAvs contract.
type TriggerXAvsOwnershipTransferredIterator struct {
	Event *TriggerXAvsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TriggerXAvsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TriggerXAvsOwnershipTransferred)
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
		it.Event = new(TriggerXAvsOwnershipTransferred)
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
func (it *TriggerXAvsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TriggerXAvsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TriggerXAvsOwnershipTransferred represents a OwnershipTransferred event raised by the TriggerXAvs contract.
type TriggerXAvsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TriggerXAvs *TriggerXAvsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TriggerXAvsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TriggerXAvs.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TriggerXAvsOwnershipTransferredIterator{contract: _TriggerXAvs.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TriggerXAvs *TriggerXAvsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TriggerXAvsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TriggerXAvs.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TriggerXAvsOwnershipTransferred)
				if err := _TriggerXAvs.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TriggerXAvs *TriggerXAvsFilterer) ParseOwnershipTransferred(log types.Log) (*TriggerXAvsOwnershipTransferred, error) {
	event := new(TriggerXAvsOwnershipTransferred)
	if err := _TriggerXAvs.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TriggerXAvsRewardManagerUpdatedIterator is returned from FilterRewardManagerUpdated and is used to iterate over the raw logs and unpacked data for RewardManagerUpdated events raised by the TriggerXAvs contract.
type TriggerXAvsRewardManagerUpdatedIterator struct {
	Event *TriggerXAvsRewardManagerUpdated // Event containing the contract specifics and raw log

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
func (it *TriggerXAvsRewardManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TriggerXAvsRewardManagerUpdated)
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
		it.Event = new(TriggerXAvsRewardManagerUpdated)
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
func (it *TriggerXAvsRewardManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TriggerXAvsRewardManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TriggerXAvsRewardManagerUpdated represents a RewardManagerUpdated event raised by the TriggerXAvs contract.
type TriggerXAvsRewardManagerUpdated struct {
	NewRewardManager common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRewardManagerUpdated is a free log retrieval operation binding the contract event 0x3d94d9e8342a65edb95eef4f65059294d45e5192603632d8dddb2344e7078053.
//
// Solidity: event RewardManagerUpdated(address indexed newRewardManager)
func (_TriggerXAvs *TriggerXAvsFilterer) FilterRewardManagerUpdated(opts *bind.FilterOpts, newRewardManager []common.Address) (*TriggerXAvsRewardManagerUpdatedIterator, error) {

	var newRewardManagerRule []interface{}
	for _, newRewardManagerItem := range newRewardManager {
		newRewardManagerRule = append(newRewardManagerRule, newRewardManagerItem)
	}

	logs, sub, err := _TriggerXAvs.contract.FilterLogs(opts, "RewardManagerUpdated", newRewardManagerRule)
	if err != nil {
		return nil, err
	}
	return &TriggerXAvsRewardManagerUpdatedIterator{contract: _TriggerXAvs.contract, event: "RewardManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardManagerUpdated is a free log subscription operation binding the contract event 0x3d94d9e8342a65edb95eef4f65059294d45e5192603632d8dddb2344e7078053.
//
// Solidity: event RewardManagerUpdated(address indexed newRewardManager)
func (_TriggerXAvs *TriggerXAvsFilterer) WatchRewardManagerUpdated(opts *bind.WatchOpts, sink chan<- *TriggerXAvsRewardManagerUpdated, newRewardManager []common.Address) (event.Subscription, error) {

	var newRewardManagerRule []interface{}
	for _, newRewardManagerItem := range newRewardManager {
		newRewardManagerRule = append(newRewardManagerRule, newRewardManagerItem)
	}

	logs, sub, err := _TriggerXAvs.contract.WatchLogs(opts, "RewardManagerUpdated", newRewardManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TriggerXAvsRewardManagerUpdated)
				if err := _TriggerXAvs.contract.UnpackLog(event, "RewardManagerUpdated", log); err != nil {
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

// ParseRewardManagerUpdated is a log parse operation binding the contract event 0x3d94d9e8342a65edb95eef4f65059294d45e5192603632d8dddb2344e7078053.
//
// Solidity: event RewardManagerUpdated(address indexed newRewardManager)
func (_TriggerXAvs *TriggerXAvsFilterer) ParseRewardManagerUpdated(log types.Log) (*TriggerXAvsRewardManagerUpdated, error) {
	event := new(TriggerXAvsRewardManagerUpdated)
	if err := _TriggerXAvs.contract.UnpackLog(event, "RewardManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TriggerXAvsSlasherUpdatedIterator is returned from FilterSlasherUpdated and is used to iterate over the raw logs and unpacked data for SlasherUpdated events raised by the TriggerXAvs contract.
type TriggerXAvsSlasherUpdatedIterator struct {
	Event *TriggerXAvsSlasherUpdated // Event containing the contract specifics and raw log

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
func (it *TriggerXAvsSlasherUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TriggerXAvsSlasherUpdated)
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
		it.Event = new(TriggerXAvsSlasherUpdated)
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
func (it *TriggerXAvsSlasherUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TriggerXAvsSlasherUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TriggerXAvsSlasherUpdated represents a SlasherUpdated event raised by the TriggerXAvs contract.
type TriggerXAvsSlasherUpdated struct {
	NewSlasher common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSlasherUpdated is a free log retrieval operation binding the contract event 0x0adf62081dae4c128a0af3a933748637b1d874a033588518f810559e6bdb23ff.
//
// Solidity: event SlasherUpdated(address indexed newSlasher)
func (_TriggerXAvs *TriggerXAvsFilterer) FilterSlasherUpdated(opts *bind.FilterOpts, newSlasher []common.Address) (*TriggerXAvsSlasherUpdatedIterator, error) {

	var newSlasherRule []interface{}
	for _, newSlasherItem := range newSlasher {
		newSlasherRule = append(newSlasherRule, newSlasherItem)
	}

	logs, sub, err := _TriggerXAvs.contract.FilterLogs(opts, "SlasherUpdated", newSlasherRule)
	if err != nil {
		return nil, err
	}
	return &TriggerXAvsSlasherUpdatedIterator{contract: _TriggerXAvs.contract, event: "SlasherUpdated", logs: logs, sub: sub}, nil
}

// WatchSlasherUpdated is a free log subscription operation binding the contract event 0x0adf62081dae4c128a0af3a933748637b1d874a033588518f810559e6bdb23ff.
//
// Solidity: event SlasherUpdated(address indexed newSlasher)
func (_TriggerXAvs *TriggerXAvsFilterer) WatchSlasherUpdated(opts *bind.WatchOpts, sink chan<- *TriggerXAvsSlasherUpdated, newSlasher []common.Address) (event.Subscription, error) {

	var newSlasherRule []interface{}
	for _, newSlasherItem := range newSlasher {
		newSlasherRule = append(newSlasherRule, newSlasherItem)
	}

	logs, sub, err := _TriggerXAvs.contract.WatchLogs(opts, "SlasherUpdated", newSlasherRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TriggerXAvsSlasherUpdated)
				if err := _TriggerXAvs.contract.UnpackLog(event, "SlasherUpdated", log); err != nil {
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

// ParseSlasherUpdated is a log parse operation binding the contract event 0x0adf62081dae4c128a0af3a933748637b1d874a033588518f810559e6bdb23ff.
//
// Solidity: event SlasherUpdated(address indexed newSlasher)
func (_TriggerXAvs *TriggerXAvsFilterer) ParseSlasherUpdated(log types.Log) (*TriggerXAvsSlasherUpdated, error) {
	event := new(TriggerXAvsSlasherUpdated)
	if err := _TriggerXAvs.contract.UnpackLog(event, "SlasherUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TriggerXAvsTaskCreatedIterator is returned from FilterTaskCreated and is used to iterate over the raw logs and unpacked data for TaskCreated events raised by the TriggerXAvs contract.
type TriggerXAvsTaskCreatedIterator struct {
	Event *TriggerXAvsTaskCreated // Event containing the contract specifics and raw log

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
func (it *TriggerXAvsTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TriggerXAvsTaskCreated)
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
		it.Event = new(TriggerXAvsTaskCreated)
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
func (it *TriggerXAvsTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TriggerXAvsTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TriggerXAvsTaskCreated represents a TaskCreated event raised by the TriggerXAvs contract.
type TriggerXAvsTaskCreated struct {
	TaskID         uint64
	DefinitionHash [32]byte
	Kind           uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTaskCreated is a free log retrieval operation binding the contract event 0x0a4f8dccc5979424692085e700d543d100dbe873f13a263782ba0988203a3689.
//
// Solidity: event TaskCreated(uint64 indexed taskID, bytes32 definitionHash, uint8 kind)
func (_TriggerXAvs *TriggerXAvsFilterer) FilterTaskCreated(opts *bind.FilterOpts, taskID []uint64) (*TriggerXAvsTaskCreatedIterator, error) {

	var taskIDRule []interface{}
	for _, taskIDItem := range taskID {
		taskIDRule = append(taskIDRule, taskIDItem)
	}

	logs, sub, err := _TriggerXAvs.contract.FilterLogs(opts, "TaskCreated", taskIDRule)
	if err != nil {
		return nil, err
	}
	return &TriggerXAvsTaskCreatedIterator{contract: _TriggerXAvs.contract, event: "TaskCreated", logs: logs, sub: sub}, nil
}

// WatchTaskCreated is a free log subscription operation binding the contract event 0x0a4f8dccc5979424692085e700d543d100dbe873f13a263782ba0988203a3689.
//
// Solidity: event TaskCreated(uint64 indexed taskID, bytes32 definitionHash, uint8 kind)
func (_TriggerXAvs *TriggerXAvsFilterer) WatchTaskCreated(opts *bind.WatchOpts, sink chan<- *TriggerXAvsTaskCreated, taskID []uint64) (event.Subscription, error) {

	var taskIDRule []interface{}
	for _, taskIDItem := range taskID {
		taskIDRule = append(taskIDRule, taskIDItem)
	}

	logs, sub, err := _TriggerXAvs.contract.WatchLogs(opts, "TaskCreated", taskIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TriggerXAvsTaskCreated)
				if err := _TriggerXAvs.contract.UnpackLog(event, "TaskCreated", log); err != nil {
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

// ParseTaskCreated is a log parse operation binding the contract event 0x0a4f8dccc5979424692085e700d543d100dbe873f13a263782ba0988203a3689.
//
// Solidity: event TaskCreated(uint64 indexed taskID, bytes32 definitionHash, uint8 kind)
func (_TriggerXAvs *TriggerXAvsFilterer) ParseTaskCreated(log types.Log) (*TriggerXAvsTaskCreated, error) {
	event := new(TriggerXAvsTaskCreated)
	if err := _TriggerXAvs.contract.UnpackLog(event, "TaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TriggerXAvsTaskDefinitionCreatedIterator is returned from FilterTaskDefinitionCreated and is used to iterate over the raw logs and unpacked data for TaskDefinitionCreated events raised by the TriggerXAvs contract.
type TriggerXAvsTaskDefinitionCreatedIterator struct {
	Event *TriggerXAvsTaskDefinitionCreated // Event containing the contract specifics and raw log

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
func (it *TriggerXAvsTaskDefinitionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TriggerXAvsTaskDefinitionCreated)
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
		it.Event = new(TriggerXAvsTaskDefinitionCreated)
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
func (it *TriggerXAvsTaskDefinitionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TriggerXAvsTaskDefinitionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TriggerXAvsTaskDefinitionCreated represents a TaskDefinitionCreated event raised by the TriggerXAvs contract.
type TriggerXAvsTaskDefinitionCreated struct {
	TaskDefinitionId uint8
	Name             string
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTaskDefinitionCreated is a free log retrieval operation binding the contract event 0x92c18501d7bf134f9f011d635d4c6d7d630c69d346a230df26302f43e132bf55.
//
// Solidity: event TaskDefinitionCreated(uint8 indexed taskDefinitionId, string name)
func (_TriggerXAvs *TriggerXAvsFilterer) FilterTaskDefinitionCreated(opts *bind.FilterOpts, taskDefinitionId []uint8) (*TriggerXAvsTaskDefinitionCreatedIterator, error) {

	var taskDefinitionIdRule []interface{}
	for _, taskDefinitionIdItem := range taskDefinitionId {
		taskDefinitionIdRule = append(taskDefinitionIdRule, taskDefinitionIdItem)
	}

	logs, sub, err := _TriggerXAvs.contract.FilterLogs(opts, "TaskDefinitionCreated", taskDefinitionIdRule)
	if err != nil {
		return nil, err
	}
	return &TriggerXAvsTaskDefinitionCreatedIterator{contract: _TriggerXAvs.contract, event: "TaskDefinitionCreated", logs: logs, sub: sub}, nil
}

// WatchTaskDefinitionCreated is a free log subscription operation binding the contract event 0x92c18501d7bf134f9f011d635d4c6d7d630c69d346a230df26302f43e132bf55.
//
// Solidity: event TaskDefinitionCreated(uint8 indexed taskDefinitionId, string name)
func (_TriggerXAvs *TriggerXAvsFilterer) WatchTaskDefinitionCreated(opts *bind.WatchOpts, sink chan<- *TriggerXAvsTaskDefinitionCreated, taskDefinitionId []uint8) (event.Subscription, error) {

	var taskDefinitionIdRule []interface{}
	for _, taskDefinitionIdItem := range taskDefinitionId {
		taskDefinitionIdRule = append(taskDefinitionIdRule, taskDefinitionIdItem)
	}

	logs, sub, err := _TriggerXAvs.contract.WatchLogs(opts, "TaskDefinitionCreated", taskDefinitionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TriggerXAvsTaskDefinitionCreated)
				if err := _TriggerXAvs.contract.UnpackLog(event, "TaskDefinitionCreated", log); err != nil {
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

// ParseTaskDefinitionCreated is a log parse operation binding the contract event 0x92c18501d7bf134f9f011d635d4c6d7d630c69d346a230df26302f43e132bf55.
//
// Solidity: event TaskDefinitionCreated(uint8 indexed taskDefinitionId, string name)
func (_TriggerXAvs *TriggerXAvsFilterer) ParseTaskDefinitionCreated(log types.Log) (*TriggerXAvsTaskDefinitionCreated, error) {
	event := new(TriggerXAvsTaskDefinitionCreated)
	if err := _TriggerXAvs.contract.UnpackLog(event, "TaskDefinitionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TriggerXAvsTaskDefinitionCreated0Iterator is returned from FilterTaskDefinitionCreated0 and is used to iterate over the raw logs and unpacked data for TaskDefinitionCreated0 events raised by the TriggerXAvs contract.
type TriggerXAvsTaskDefinitionCreated0Iterator struct {
	Event *TriggerXAvsTaskDefinitionCreated0 // Event containing the contract specifics and raw log

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
func (it *TriggerXAvsTaskDefinitionCreated0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TriggerXAvsTaskDefinitionCreated0)
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
		it.Event = new(TriggerXAvsTaskDefinitionCreated0)
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
func (it *TriggerXAvsTaskDefinitionCreated0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TriggerXAvsTaskDefinitionCreated0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TriggerXAvsTaskDefinitionCreated0 represents a TaskDefinitionCreated0 event raised by the TriggerXAvs contract.
type TriggerXAvsTaskDefinitionCreated0 struct {
	TaskDefinitionId           uint8
	Name                       string
	BaseRewardFeeForAttesters  *big.Int
	BaseRewardFeeForPerformer  *big.Int
	BaseRewardFeeForAggregator *big.Int
	MinimumVotingPower         *big.Int
	RestrictedOperatorIds      []*big.Int
	MaximumNumberOfAttesters   *big.Int
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterTaskDefinitionCreated0 is a free log retrieval operation binding the contract event 0x922bce103cd095cc3c1297fd9537ca88abcc2bdd3a03c9e0086fd9adff5389d2.
//
// Solidity: event TaskDefinitionCreated(uint8 taskDefinitionId, string name, uint256 baseRewardFeeForAttesters, uint256 baseRewardFeeForPerformer, uint256 baseRewardFeeForAggregator, uint256 minimumVotingPower, uint256[] restrictedOperatorIds, uint256 maximumNumberOfAttesters)
func (_TriggerXAvs *TriggerXAvsFilterer) FilterTaskDefinitionCreated0(opts *bind.FilterOpts) (*TriggerXAvsTaskDefinitionCreated0Iterator, error) {

	logs, sub, err := _TriggerXAvs.contract.FilterLogs(opts, "TaskDefinitionCreated0")
	if err != nil {
		return nil, err
	}
	return &TriggerXAvsTaskDefinitionCreated0Iterator{contract: _TriggerXAvs.contract, event: "TaskDefinitionCreated0", logs: logs, sub: sub}, nil
}

// WatchTaskDefinitionCreated0 is a free log subscription operation binding the contract event 0x922bce103cd095cc3c1297fd9537ca88abcc2bdd3a03c9e0086fd9adff5389d2.
//
// Solidity: event TaskDefinitionCreated(uint8 taskDefinitionId, string name, uint256 baseRewardFeeForAttesters, uint256 baseRewardFeeForPerformer, uint256 baseRewardFeeForAggregator, uint256 minimumVotingPower, uint256[] restrictedOperatorIds, uint256 maximumNumberOfAttesters)
func (_TriggerXAvs *TriggerXAvsFilterer) WatchTaskDefinitionCreated0(opts *bind.WatchOpts, sink chan<- *TriggerXAvsTaskDefinitionCreated0) (event.Subscription, error) {

	logs, sub, err := _TriggerXAvs.contract.WatchLogs(opts, "TaskDefinitionCreated0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TriggerXAvsTaskDefinitionCreated0)
				if err := _TriggerXAvs.contract.UnpackLog(event, "TaskDefinitionCreated0", log); err != nil {
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

// ParseTaskDefinitionCreated0 is a log parse operation binding the contract event 0x922bce103cd095cc3c1297fd9537ca88abcc2bdd3a03c9e0086fd9adff5389d2.
//
// Solidity: event TaskDefinitionCreated(uint8 taskDefinitionId, string name, uint256 baseRewardFeeForAttesters, uint256 baseRewardFeeForPerformer, uint256 baseRewardFeeForAggregator, uint256 minimumVotingPower, uint256[] restrictedOperatorIds, uint256 maximumNumberOfAttesters)
func (_TriggerXAvs *TriggerXAvsFilterer) ParseTaskDefinitionCreated0(log types.Log) (*TriggerXAvsTaskDefinitionCreated0, error) {
	event := new(TriggerXAvsTaskDefinitionCreated0)
	if err := _TriggerXAvs.contract.UnpackLog(event, "TaskDefinitionCreated0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TriggerXAvsTaskSubmittedIterator is returned from FilterTaskSubmitted and is used to iterate over the raw logs and unpacked data for TaskSubmitted events raised by the TriggerXAvs contract.
type TriggerXAvsTaskSubmittedIterator struct {
	Event *TriggerXAvsTaskSubmitted // Event containing the contract specifics and raw log

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
func (it *TriggerXAvsTaskSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TriggerXAvsTaskSubmitted)
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
		it.Event = new(TriggerXAvsTaskSubmitted)
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
func (it *TriggerXAvsTaskSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TriggerXAvsTaskSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TriggerXAvsTaskSubmitted represents a TaskSubmitted event raised by the TriggerXAvs contract.
type TriggerXAvsTaskSubmitted struct {
	TaskID   uint64
	Operator common.Address
	Phase    uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTaskSubmitted is a free log retrieval operation binding the contract event 0x852ef925dfb4a67eaf327f40c6c555596a5c1b2099e05cad9362356ec8916ae3.
//
// Solidity: event TaskSubmitted(uint64 indexed taskID, address indexed operator, uint8 phase)
func (_TriggerXAvs *TriggerXAvsFilterer) FilterTaskSubmitted(opts *bind.FilterOpts, taskID []uint64, operator []common.Address) (*TriggerXAvsTaskSubmittedIterator, error) {

	var taskIDRule []interface{}
	for _, taskIDItem := range taskID {
		taskIDRule = append(taskIDRule, taskIDItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TriggerXAvs.contract.FilterLogs(opts, "TaskSubmitted", taskIDRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &TriggerXAvsTaskSubmittedIterator{contract: _TriggerXAvs.contract, event: "TaskSubmitted", logs: logs, sub: sub}, nil
}

// WatchTaskSubmitted is a free log subscription operation binding the contract event 0x852ef925dfb4a67eaf327f40c6c555596a5c1b2099e05cad9362356ec8916ae3.
//
// Solidity: event TaskSubmitted(uint64 indexed taskID, address indexed operator, uint8 phase)
func (_TriggerXAvs *TriggerXAvsFilterer) WatchTaskSubmitted(opts *bind.WatchOpts, sink chan<- *TriggerXAvsTaskSubmitted, taskID []uint64, operator []common.Address) (event.Subscription, error) {

	var taskIDRule []interface{}
	for _, taskIDItem := range taskID {
		taskIDRule = append(taskIDRule, taskIDItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TriggerXAvs.contract.WatchLogs(opts, "TaskSubmitted", taskIDRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TriggerXAvsTaskSubmitted)
				if err := _TriggerXAvs.contract.UnpackLog(event, "TaskSubmitted", log); err != nil {
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

// ParseTaskSubmitted is a log parse operation binding the contract event 0x852ef925dfb4a67eaf327f40c6c555596a5c1b2099e05cad9362356ec8916ae3.
//
// Solidity: event TaskSubmitted(uint64 indexed taskID, address indexed operator, uint8 phase)
func (_TriggerXAvs *TriggerXAvsFilterer) ParseTaskSubmitted(log types.Log) (*TriggerXAvsTaskSubmitted, error) {
	event := new(TriggerXAvsTaskSubmitted)
	if err := _TriggerXAvs.contract.UnpackLog(event, "TaskSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TriggerXAvsUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the TriggerXAvs contract.
type TriggerXAvsUpgradedIterator struct {
	Event *TriggerXAvsUpgraded // Event containing the contract specifics and raw log

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
func (it *TriggerXAvsUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TriggerXAvsUpgraded)
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
		it.Event = new(TriggerXAvsUpgraded)
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
func (it *TriggerXAvsUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TriggerXAvsUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TriggerXAvsUpgraded represents a Upgraded event raised by the TriggerXAvs contract.
type TriggerXAvsUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TriggerXAvs *TriggerXAvsFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*TriggerXAvsUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TriggerXAvs.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &TriggerXAvsUpgradedIterator{contract: _TriggerXAvs.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TriggerXAvs *TriggerXAvsFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *TriggerXAvsUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TriggerXAvs.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TriggerXAvsUpgraded)
				if err := _TriggerXAvs.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TriggerXAvs *TriggerXAvsFilterer) ParseUpgraded(log types.Log) (*TriggerXAvsUpgraded, error) {
	event := new(TriggerXAvsUpgraded)
	if err := _TriggerXAvs.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
