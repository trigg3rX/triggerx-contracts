// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractOBLS

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

// IOBLSOperatorVotingPower is an auto generated low-level Go binding around an user-defined struct.
type IOBLSOperatorVotingPower struct {
	OperatorId  *big.Int
	VotingPower *big.Int
}

// ContractOBLSMetaData contains all meta data concerning the ContractOBLS contract.
var ContractOBLSMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BNAddCallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadFTMappingImplementation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"operator\",\"type\":\"uint256\"}],\"name\":\"InactiveOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientVotingPower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAuthSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFieldElement\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOBLSSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOperatorIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOperatorIndexes\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRequiredVotingPower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ModularInverseError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOBLSManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOBLSManagerOrShareSyncer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorIndex\",\"type\":\"uint256\"}],\"name\":\"OperatorDoesNotHaveMinimumVotingPower\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votingPower\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIOBLS.OperatorVotingPower[]\",\"name\":\"operatorsVotingPower\",\"type\":\"tuple[]\"}],\"name\":\"DecreaseBatchOperatorVotingPower\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votingPower\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIOBLS.OperatorVotingPower[]\",\"name\":\"operatorsVotingPower\",\"type\":\"tuple[]\"}],\"name\":\"IncreaseBatchOperatorVotingPower\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOBLSManager\",\"type\":\"address\"}],\"name\":\"SetOBLSManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"taskDefinitionId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumVotingPower\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"restrictedAttesterIds\",\"type\":\"uint256[]\"}],\"name\":\"SetTotalVotingPowerPerRestrictedTaskDefinition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"taskdefinitionId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numOfTotalOperators\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumVotingPower\",\"type\":\"uint256\"}],\"name\":\"SetTotalVotingPowerPerTaskDefinition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"syncer\",\"type\":\"address\"}],\"name\":\"SharesSyncerModified\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votingPower\",\"type\":\"uint256\"}],\"internalType\":\"structIOBLS.OperatorVotingPower[]\",\"name\":\"_operatorsVotingPower\",\"type\":\"tuple[]\"}],\"name\":\"decreaseBatchOperatorVotingPower\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_votingPower\",\"type\":\"uint256\"}],\"name\":\"decreaseOperatorVotingPower\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_taskDefinitionId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_votingPower\",\"type\":\"uint256\"}],\"name\":\"decreaseOperatorVotingPowerPerTaskDefinition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOblsManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getOperatorBLSPubKey\",\"outputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"\",\"type\":\"uint256[4]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"domain\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"hashToPoint\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votingPower\",\"type\":\"uint256\"}],\"internalType\":\"structIOBLS.OperatorVotingPower[]\",\"name\":\"_operatorsVotingPower\",\"type\":\"tuple[]\"}],\"name\":\"increaseBatchOperatorVotingPower\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_votingPower\",\"type\":\"uint256\"}],\"name\":\"increaseOperatorVotingPower\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_taskDefinitionId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_votingPower\",\"type\":\"uint256\"}],\"name\":\"increaseOperatorVotingPowerPerTaskDefinition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"isActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"isRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isRegistered\",\"type\":\"bool\"}],\"name\":\"modifyOperatorActiveStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"_blsKey\",\"type\":\"uint256[4]\"}],\"name\":\"modifyOperatorBlsKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_votingPower\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"_blsKey\",\"type\":\"uint256[4]\"}],\"name\":\"registerOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oblsManager\",\"type\":\"address\"}],\"name\":\"setOblsManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oblsSharesSyncer\",\"type\":\"address\"}],\"name\":\"setOblsSharesSyncer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_taskDefinitionId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_minimumVotingPower\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_restrictedAttesterIds\",\"type\":\"uint256[]\"}],\"name\":\"setTotalVotingPowerPerRestrictedTaskDefinition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_taskDefinitionId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_numOfTotalOperators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minimumVotingPower\",\"type\":\"uint256\"}],\"name\":\"setTotalVotingPowerPerTaskDefinition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_votingPowerSyncer\",\"type\":\"address\"}],\"name\":\"setVotingPowerSyncer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_votingPowers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[4][]\",\"name\":\"_blsKeys\",\"type\":\"uint256[4][]\"}],\"name\":\"syncOperatorDetails\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalVotingPower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"totalVotingPowerPerTaskDefinition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"unRegisterOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[2]\",\"name\":\"_message\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"_signature\",\"type\":\"uint256[2]\"}],\"name\":\"validateOperatorSignature\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"signature\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBLSAuthLibrary.Signature\",\"name\":\"_signature\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"},{\"internalType\":\"uint256[4]\",\"name\":\"_blsKey\",\"type\":\"uint256[4]\"}],\"name\":\"verifyAuthSignature\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"_message\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"_signature\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[]\",\"name\":\"_indexes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_requiredVotingPower\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minimumVotingPowerPerTaskDefinition\",\"type\":\"uint256\"}],\"name\":\"verifySignature\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"votingPower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ContractOBLSABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractOBLSMetaData.ABI instead.
var ContractOBLSABI = ContractOBLSMetaData.ABI

// ContractOBLSMethods is an auto generated interface around an Ethereum contract.
type ContractOBLSMethods interface {
	ContractOBLSCalls
	ContractOBLSTransacts
	ContractOBLSFilters
}

// ContractOBLSCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractOBLSCalls interface {
	DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error)

	GetOblsManager(opts *bind.CallOpts) (common.Address, error)

	GetOperatorBLSPubKey(opts *bind.CallOpts, _index *big.Int) ([4]*big.Int, error)

	GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error)

	HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error)

	HashToPoint(opts *bind.CallOpts, domain [32]byte, message []byte) ([2]*big.Int, error)

	IsActive(opts *bind.CallOpts, _index *big.Int) (bool, error)

	IsRegistered(opts *bind.CallOpts, _index *big.Int) (bool, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TotalVotingPower(opts *bind.CallOpts) (*big.Int, error)

	TotalVotingPowerPerTaskDefinition(opts *bind.CallOpts, _id *big.Int) (*big.Int, error)

	ValidateOperatorSignature(opts *bind.CallOpts, _operatorId *big.Int, _message [2]*big.Int, _signature [2]*big.Int) error

	VerifyAuthSignature(opts *bind.CallOpts, _signature BLSAuthLibrarySignature, _operator common.Address, _contract common.Address, _blsKey [4]*big.Int) error

	VerifySignature(opts *bind.CallOpts, _message [2]*big.Int, _signature [2]*big.Int, _indexes []*big.Int, _requiredVotingPower *big.Int, _minimumVotingPowerPerTaskDefinition *big.Int) error

	VotingPower(opts *bind.CallOpts, _index *big.Int) (*big.Int, error)
}

// ContractOBLSTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractOBLSTransacts interface {
	DecreaseBatchOperatorVotingPower(opts *bind.TransactOpts, _operatorsVotingPower []IOBLSOperatorVotingPower) (*types.Transaction, error)

	DecreaseOperatorVotingPower(opts *bind.TransactOpts, _index *big.Int, _votingPower *big.Int) (*types.Transaction, error)

	DecreaseOperatorVotingPowerPerTaskDefinition(opts *bind.TransactOpts, _taskDefinitionId uint16, _votingPower *big.Int) (*types.Transaction, error)

	GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	IncreaseBatchOperatorVotingPower(opts *bind.TransactOpts, _operatorsVotingPower []IOBLSOperatorVotingPower) (*types.Transaction, error)

	IncreaseOperatorVotingPower(opts *bind.TransactOpts, _index *big.Int, _votingPower *big.Int) (*types.Transaction, error)

	IncreaseOperatorVotingPowerPerTaskDefinition(opts *bind.TransactOpts, _taskDefinitionId uint16, _votingPower *big.Int) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts) (*types.Transaction, error)

	Migration(opts *bind.TransactOpts) (*types.Transaction, error)

	ModifyOperatorActiveStatus(opts *bind.TransactOpts, _index *big.Int, _isRegistered bool) (*types.Transaction, error)

	ModifyOperatorBlsKey(opts *bind.TransactOpts, _index *big.Int, _blsKey [4]*big.Int) (*types.Transaction, error)

	RegisterOperator(opts *bind.TransactOpts, _index *big.Int, _votingPower *big.Int, _blsKey [4]*big.Int) (*types.Transaction, error)

	RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error)

	RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	SetOblsManager(opts *bind.TransactOpts, _oblsManager common.Address) (*types.Transaction, error)

	SetOblsSharesSyncer(opts *bind.TransactOpts, _oblsSharesSyncer common.Address) (*types.Transaction, error)

	SetTotalVotingPowerPerRestrictedTaskDefinition(opts *bind.TransactOpts, _taskDefinitionId uint16, _minimumVotingPower *big.Int, _restrictedAttesterIds []*big.Int) (*types.Transaction, error)

	SetTotalVotingPowerPerTaskDefinition(opts *bind.TransactOpts, _taskDefinitionId uint16, _numOfTotalOperators *big.Int, _minimumVotingPower *big.Int) (*types.Transaction, error)

	SetVotingPowerSyncer(opts *bind.TransactOpts, _votingPowerSyncer common.Address) (*types.Transaction, error)

	SyncOperatorDetails(opts *bind.TransactOpts, _votingPowers []*big.Int, _blsKeys [][4]*big.Int) (*types.Transaction, error)

	UnRegisterOperator(opts *bind.TransactOpts, _index *big.Int) (*types.Transaction, error)
}

// ContractOBLSFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractOBLSFilters interface {
	FilterDecreaseBatchOperatorVotingPower(opts *bind.FilterOpts) (*ContractOBLSDecreaseBatchOperatorVotingPowerIterator, error)
	WatchDecreaseBatchOperatorVotingPower(opts *bind.WatchOpts, sink chan<- *ContractOBLSDecreaseBatchOperatorVotingPower) (event.Subscription, error)
	ParseDecreaseBatchOperatorVotingPower(log types.Log) (*ContractOBLSDecreaseBatchOperatorVotingPower, error)

	FilterIncreaseBatchOperatorVotingPower(opts *bind.FilterOpts) (*ContractOBLSIncreaseBatchOperatorVotingPowerIterator, error)
	WatchIncreaseBatchOperatorVotingPower(opts *bind.WatchOpts, sink chan<- *ContractOBLSIncreaseBatchOperatorVotingPower) (event.Subscription, error)
	ParseIncreaseBatchOperatorVotingPower(log types.Log) (*ContractOBLSIncreaseBatchOperatorVotingPower, error)

	FilterInitialized(opts *bind.FilterOpts) (*ContractOBLSInitializedIterator, error)
	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractOBLSInitialized) (event.Subscription, error)
	ParseInitialized(log types.Log) (*ContractOBLSInitialized, error)

	FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ContractOBLSRoleAdminChangedIterator, error)
	WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ContractOBLSRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error)
	ParseRoleAdminChanged(log types.Log) (*ContractOBLSRoleAdminChanged, error)

	FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ContractOBLSRoleGrantedIterator, error)
	WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ContractOBLSRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error)
	ParseRoleGranted(log types.Log) (*ContractOBLSRoleGranted, error)

	FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ContractOBLSRoleRevokedIterator, error)
	WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ContractOBLSRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error)
	ParseRoleRevoked(log types.Log) (*ContractOBLSRoleRevoked, error)

	FilterSetOBLSManager(opts *bind.FilterOpts) (*ContractOBLSSetOBLSManagerIterator, error)
	WatchSetOBLSManager(opts *bind.WatchOpts, sink chan<- *ContractOBLSSetOBLSManager) (event.Subscription, error)
	ParseSetOBLSManager(log types.Log) (*ContractOBLSSetOBLSManager, error)

	FilterSetTotalVotingPowerPerRestrictedTaskDefinition(opts *bind.FilterOpts) (*ContractOBLSSetTotalVotingPowerPerRestrictedTaskDefinitionIterator, error)
	WatchSetTotalVotingPowerPerRestrictedTaskDefinition(opts *bind.WatchOpts, sink chan<- *ContractOBLSSetTotalVotingPowerPerRestrictedTaskDefinition) (event.Subscription, error)
	ParseSetTotalVotingPowerPerRestrictedTaskDefinition(log types.Log) (*ContractOBLSSetTotalVotingPowerPerRestrictedTaskDefinition, error)

	FilterSetTotalVotingPowerPerTaskDefinition(opts *bind.FilterOpts) (*ContractOBLSSetTotalVotingPowerPerTaskDefinitionIterator, error)
	WatchSetTotalVotingPowerPerTaskDefinition(opts *bind.WatchOpts, sink chan<- *ContractOBLSSetTotalVotingPowerPerTaskDefinition) (event.Subscription, error)
	ParseSetTotalVotingPowerPerTaskDefinition(log types.Log) (*ContractOBLSSetTotalVotingPowerPerTaskDefinition, error)

	FilterSharesSyncerModified(opts *bind.FilterOpts) (*ContractOBLSSharesSyncerModifiedIterator, error)
	WatchSharesSyncerModified(opts *bind.WatchOpts, sink chan<- *ContractOBLSSharesSyncerModified) (event.Subscription, error)
	ParseSharesSyncerModified(log types.Log) (*ContractOBLSSharesSyncerModified, error)
}

// ContractOBLS is an auto generated Go binding around an Ethereum contract.
type ContractOBLS struct {
	ContractOBLSCaller     // Read-only binding to the contract
	ContractOBLSTransactor // Write-only binding to the contract
	ContractOBLSFilterer   // Log filterer for contract events
}

// ContractOBLS implements the ContractOBLSMethods interface.
var _ ContractOBLSMethods = (*ContractOBLS)(nil)

// ContractOBLSCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractOBLSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractOBLSCaller implements the ContractOBLSCalls interface.
var _ ContractOBLSCalls = (*ContractOBLSCaller)(nil)

// ContractOBLSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractOBLSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractOBLSTransactor implements the ContractOBLSTransacts interface.
var _ ContractOBLSTransacts = (*ContractOBLSTransactor)(nil)

// ContractOBLSFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractOBLSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractOBLSFilterer implements the ContractOBLSFilters interface.
var _ ContractOBLSFilters = (*ContractOBLSFilterer)(nil)

// ContractOBLSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractOBLSSession struct {
	Contract     *ContractOBLS     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractOBLSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractOBLSCallerSession struct {
	Contract *ContractOBLSCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ContractOBLSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractOBLSTransactorSession struct {
	Contract     *ContractOBLSTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ContractOBLSRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractOBLSRaw struct {
	Contract *ContractOBLS // Generic contract binding to access the raw methods on
}

// ContractOBLSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractOBLSCallerRaw struct {
	Contract *ContractOBLSCaller // Generic read-only contract binding to access the raw methods on
}

// ContractOBLSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractOBLSTransactorRaw struct {
	Contract *ContractOBLSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractOBLS creates a new instance of ContractOBLS, bound to a specific deployed contract.
func NewContractOBLS(address common.Address, backend bind.ContractBackend) (*ContractOBLS, error) {
	contract, err := bindContractOBLS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractOBLS{ContractOBLSCaller: ContractOBLSCaller{contract: contract}, ContractOBLSTransactor: ContractOBLSTransactor{contract: contract}, ContractOBLSFilterer: ContractOBLSFilterer{contract: contract}}, nil
}

// NewContractOBLSCaller creates a new read-only instance of ContractOBLS, bound to a specific deployed contract.
func NewContractOBLSCaller(address common.Address, caller bind.ContractCaller) (*ContractOBLSCaller, error) {
	contract, err := bindContractOBLS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractOBLSCaller{contract: contract}, nil
}

// NewContractOBLSTransactor creates a new write-only instance of ContractOBLS, bound to a specific deployed contract.
func NewContractOBLSTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractOBLSTransactor, error) {
	contract, err := bindContractOBLS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractOBLSTransactor{contract: contract}, nil
}

// NewContractOBLSFilterer creates a new log filterer instance of ContractOBLS, bound to a specific deployed contract.
func NewContractOBLSFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractOBLSFilterer, error) {
	contract, err := bindContractOBLS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractOBLSFilterer{contract: contract}, nil
}

// bindContractOBLS binds a generic wrapper to an already deployed contract.
func bindContractOBLS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractOBLSMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractOBLS *ContractOBLSRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractOBLS.Contract.ContractOBLSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractOBLS *ContractOBLSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractOBLS.Contract.ContractOBLSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractOBLS *ContractOBLSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractOBLS.Contract.ContractOBLSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractOBLS *ContractOBLSCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractOBLS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractOBLS *ContractOBLSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractOBLS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractOBLS *ContractOBLSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractOBLS.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ContractOBLS *ContractOBLSCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractOBLS.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ContractOBLS *ContractOBLSSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ContractOBLS.Contract.DEFAULTADMINROLE(&_ContractOBLS.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ContractOBLS *ContractOBLSCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ContractOBLS.Contract.DEFAULTADMINROLE(&_ContractOBLS.CallOpts)
}

// GetOblsManager is a free data retrieval call binding the contract method 0xb3dfe1e8.
//
// Solidity: function getOblsManager() view returns(address)
func (_ContractOBLS *ContractOBLSCaller) GetOblsManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractOBLS.contract.Call(opts, &out, "getOblsManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOblsManager is a free data retrieval call binding the contract method 0xb3dfe1e8.
//
// Solidity: function getOblsManager() view returns(address)
func (_ContractOBLS *ContractOBLSSession) GetOblsManager() (common.Address, error) {
	return _ContractOBLS.Contract.GetOblsManager(&_ContractOBLS.CallOpts)
}

// GetOblsManager is a free data retrieval call binding the contract method 0xb3dfe1e8.
//
// Solidity: function getOblsManager() view returns(address)
func (_ContractOBLS *ContractOBLSCallerSession) GetOblsManager() (common.Address, error) {
	return _ContractOBLS.Contract.GetOblsManager(&_ContractOBLS.CallOpts)
}

// GetOperatorBLSPubKey is a free data retrieval call binding the contract method 0xd4af0b9c.
//
// Solidity: function getOperatorBLSPubKey(uint256 _index) view returns(uint256[4])
func (_ContractOBLS *ContractOBLSCaller) GetOperatorBLSPubKey(opts *bind.CallOpts, _index *big.Int) ([4]*big.Int, error) {
	var out []interface{}
	err := _ContractOBLS.contract.Call(opts, &out, "getOperatorBLSPubKey", _index)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// GetOperatorBLSPubKey is a free data retrieval call binding the contract method 0xd4af0b9c.
//
// Solidity: function getOperatorBLSPubKey(uint256 _index) view returns(uint256[4])
func (_ContractOBLS *ContractOBLSSession) GetOperatorBLSPubKey(_index *big.Int) ([4]*big.Int, error) {
	return _ContractOBLS.Contract.GetOperatorBLSPubKey(&_ContractOBLS.CallOpts, _index)
}

// GetOperatorBLSPubKey is a free data retrieval call binding the contract method 0xd4af0b9c.
//
// Solidity: function getOperatorBLSPubKey(uint256 _index) view returns(uint256[4])
func (_ContractOBLS *ContractOBLSCallerSession) GetOperatorBLSPubKey(_index *big.Int) ([4]*big.Int, error) {
	return _ContractOBLS.Contract.GetOperatorBLSPubKey(&_ContractOBLS.CallOpts, _index)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ContractOBLS *ContractOBLSCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ContractOBLS.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ContractOBLS *ContractOBLSSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ContractOBLS.Contract.GetRoleAdmin(&_ContractOBLS.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ContractOBLS *ContractOBLSCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ContractOBLS.Contract.GetRoleAdmin(&_ContractOBLS.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ContractOBLS *ContractOBLSCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ContractOBLS.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ContractOBLS *ContractOBLSSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ContractOBLS.Contract.HasRole(&_ContractOBLS.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ContractOBLS *ContractOBLSCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ContractOBLS.Contract.HasRole(&_ContractOBLS.CallOpts, role, account)
}

// HashToPoint is a free data retrieval call binding the contract method 0xa850a909.
//
// Solidity: function hashToPoint(bytes32 domain, bytes message) view returns(uint256[2])
func (_ContractOBLS *ContractOBLSCaller) HashToPoint(opts *bind.CallOpts, domain [32]byte, message []byte) ([2]*big.Int, error) {
	var out []interface{}
	err := _ContractOBLS.contract.Call(opts, &out, "hashToPoint", domain, message)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// HashToPoint is a free data retrieval call binding the contract method 0xa850a909.
//
// Solidity: function hashToPoint(bytes32 domain, bytes message) view returns(uint256[2])
func (_ContractOBLS *ContractOBLSSession) HashToPoint(domain [32]byte, message []byte) ([2]*big.Int, error) {
	return _ContractOBLS.Contract.HashToPoint(&_ContractOBLS.CallOpts, domain, message)
}

// HashToPoint is a free data retrieval call binding the contract method 0xa850a909.
//
// Solidity: function hashToPoint(bytes32 domain, bytes message) view returns(uint256[2])
func (_ContractOBLS *ContractOBLSCallerSession) HashToPoint(domain [32]byte, message []byte) ([2]*big.Int, error) {
	return _ContractOBLS.Contract.HashToPoint(&_ContractOBLS.CallOpts, domain, message)
}

// IsActive is a free data retrieval call binding the contract method 0x82afd23b.
//
// Solidity: function isActive(uint256 _index) view returns(bool)
func (_ContractOBLS *ContractOBLSCaller) IsActive(opts *bind.CallOpts, _index *big.Int) (bool, error) {
	var out []interface{}
	err := _ContractOBLS.contract.Call(opts, &out, "isActive", _index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActive is a free data retrieval call binding the contract method 0x82afd23b.
//
// Solidity: function isActive(uint256 _index) view returns(bool)
func (_ContractOBLS *ContractOBLSSession) IsActive(_index *big.Int) (bool, error) {
	return _ContractOBLS.Contract.IsActive(&_ContractOBLS.CallOpts, _index)
}

// IsActive is a free data retrieval call binding the contract method 0x82afd23b.
//
// Solidity: function isActive(uint256 _index) view returns(bool)
func (_ContractOBLS *ContractOBLSCallerSession) IsActive(_index *big.Int) (bool, error) {
	return _ContractOBLS.Contract.IsActive(&_ContractOBLS.CallOpts, _index)
}

// IsRegistered is a free data retrieval call binding the contract method 0x579a6988.
//
// Solidity: function isRegistered(uint256 _index) view returns(bool)
func (_ContractOBLS *ContractOBLSCaller) IsRegistered(opts *bind.CallOpts, _index *big.Int) (bool, error) {
	var out []interface{}
	err := _ContractOBLS.contract.Call(opts, &out, "isRegistered", _index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegistered is a free data retrieval call binding the contract method 0x579a6988.
//
// Solidity: function isRegistered(uint256 _index) view returns(bool)
func (_ContractOBLS *ContractOBLSSession) IsRegistered(_index *big.Int) (bool, error) {
	return _ContractOBLS.Contract.IsRegistered(&_ContractOBLS.CallOpts, _index)
}

// IsRegistered is a free data retrieval call binding the contract method 0x579a6988.
//
// Solidity: function isRegistered(uint256 _index) view returns(bool)
func (_ContractOBLS *ContractOBLSCallerSession) IsRegistered(_index *big.Int) (bool, error) {
	return _ContractOBLS.Contract.IsRegistered(&_ContractOBLS.CallOpts, _index)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ContractOBLS *ContractOBLSCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ContractOBLS.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ContractOBLS *ContractOBLSSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ContractOBLS.Contract.SupportsInterface(&_ContractOBLS.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ContractOBLS *ContractOBLSCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ContractOBLS.Contract.SupportsInterface(&_ContractOBLS.CallOpts, interfaceId)
}

// TotalVotingPower is a free data retrieval call binding the contract method 0x671b3793.
//
// Solidity: function totalVotingPower() view returns(uint256)
func (_ContractOBLS *ContractOBLSCaller) TotalVotingPower(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractOBLS.contract.Call(opts, &out, "totalVotingPower")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalVotingPower is a free data retrieval call binding the contract method 0x671b3793.
//
// Solidity: function totalVotingPower() view returns(uint256)
func (_ContractOBLS *ContractOBLSSession) TotalVotingPower() (*big.Int, error) {
	return _ContractOBLS.Contract.TotalVotingPower(&_ContractOBLS.CallOpts)
}

// TotalVotingPower is a free data retrieval call binding the contract method 0x671b3793.
//
// Solidity: function totalVotingPower() view returns(uint256)
func (_ContractOBLS *ContractOBLSCallerSession) TotalVotingPower() (*big.Int, error) {
	return _ContractOBLS.Contract.TotalVotingPower(&_ContractOBLS.CallOpts)
}

// TotalVotingPowerPerTaskDefinition is a free data retrieval call binding the contract method 0x255da331.
//
// Solidity: function totalVotingPowerPerTaskDefinition(uint256 _id) view returns(uint256)
func (_ContractOBLS *ContractOBLSCaller) TotalVotingPowerPerTaskDefinition(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ContractOBLS.contract.Call(opts, &out, "totalVotingPowerPerTaskDefinition", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalVotingPowerPerTaskDefinition is a free data retrieval call binding the contract method 0x255da331.
//
// Solidity: function totalVotingPowerPerTaskDefinition(uint256 _id) view returns(uint256)
func (_ContractOBLS *ContractOBLSSession) TotalVotingPowerPerTaskDefinition(_id *big.Int) (*big.Int, error) {
	return _ContractOBLS.Contract.TotalVotingPowerPerTaskDefinition(&_ContractOBLS.CallOpts, _id)
}

// TotalVotingPowerPerTaskDefinition is a free data retrieval call binding the contract method 0x255da331.
//
// Solidity: function totalVotingPowerPerTaskDefinition(uint256 _id) view returns(uint256)
func (_ContractOBLS *ContractOBLSCallerSession) TotalVotingPowerPerTaskDefinition(_id *big.Int) (*big.Int, error) {
	return _ContractOBLS.Contract.TotalVotingPowerPerTaskDefinition(&_ContractOBLS.CallOpts, _id)
}

// ValidateOperatorSignature is a free data retrieval call binding the contract method 0x53c051bb.
//
// Solidity: function validateOperatorSignature(uint256 _operatorId, uint256[2] _message, uint256[2] _signature) view returns()
func (_ContractOBLS *ContractOBLSCaller) ValidateOperatorSignature(opts *bind.CallOpts, _operatorId *big.Int, _message [2]*big.Int, _signature [2]*big.Int) error {
	var out []interface{}
	err := _ContractOBLS.contract.Call(opts, &out, "validateOperatorSignature", _operatorId, _message, _signature)

	if err != nil {
		return err
	}

	return err

}

// ValidateOperatorSignature is a free data retrieval call binding the contract method 0x53c051bb.
//
// Solidity: function validateOperatorSignature(uint256 _operatorId, uint256[2] _message, uint256[2] _signature) view returns()
func (_ContractOBLS *ContractOBLSSession) ValidateOperatorSignature(_operatorId *big.Int, _message [2]*big.Int, _signature [2]*big.Int) error {
	return _ContractOBLS.Contract.ValidateOperatorSignature(&_ContractOBLS.CallOpts, _operatorId, _message, _signature)
}

// ValidateOperatorSignature is a free data retrieval call binding the contract method 0x53c051bb.
//
// Solidity: function validateOperatorSignature(uint256 _operatorId, uint256[2] _message, uint256[2] _signature) view returns()
func (_ContractOBLS *ContractOBLSCallerSession) ValidateOperatorSignature(_operatorId *big.Int, _message [2]*big.Int, _signature [2]*big.Int) error {
	return _ContractOBLS.Contract.ValidateOperatorSignature(&_ContractOBLS.CallOpts, _operatorId, _message, _signature)
}

// VerifyAuthSignature is a free data retrieval call binding the contract method 0x8ebaaa3e.
//
// Solidity: function verifyAuthSignature((uint256[2]) _signature, address _operator, address _contract, uint256[4] _blsKey) view returns()
func (_ContractOBLS *ContractOBLSCaller) VerifyAuthSignature(opts *bind.CallOpts, _signature BLSAuthLibrarySignature, _operator common.Address, _contract common.Address, _blsKey [4]*big.Int) error {
	var out []interface{}
	err := _ContractOBLS.contract.Call(opts, &out, "verifyAuthSignature", _signature, _operator, _contract, _blsKey)

	if err != nil {
		return err
	}

	return err

}

// VerifyAuthSignature is a free data retrieval call binding the contract method 0x8ebaaa3e.
//
// Solidity: function verifyAuthSignature((uint256[2]) _signature, address _operator, address _contract, uint256[4] _blsKey) view returns()
func (_ContractOBLS *ContractOBLSSession) VerifyAuthSignature(_signature BLSAuthLibrarySignature, _operator common.Address, _contract common.Address, _blsKey [4]*big.Int) error {
	return _ContractOBLS.Contract.VerifyAuthSignature(&_ContractOBLS.CallOpts, _signature, _operator, _contract, _blsKey)
}

// VerifyAuthSignature is a free data retrieval call binding the contract method 0x8ebaaa3e.
//
// Solidity: function verifyAuthSignature((uint256[2]) _signature, address _operator, address _contract, uint256[4] _blsKey) view returns()
func (_ContractOBLS *ContractOBLSCallerSession) VerifyAuthSignature(_signature BLSAuthLibrarySignature, _operator common.Address, _contract common.Address, _blsKey [4]*big.Int) error {
	return _ContractOBLS.Contract.VerifyAuthSignature(&_ContractOBLS.CallOpts, _signature, _operator, _contract, _blsKey)
}

// VerifySignature is a free data retrieval call binding the contract method 0x65c46475.
//
// Solidity: function verifySignature(uint256[2] _message, uint256[2] _signature, uint256[] _indexes, uint256 _requiredVotingPower, uint256 _minimumVotingPowerPerTaskDefinition) view returns()
func (_ContractOBLS *ContractOBLSCaller) VerifySignature(opts *bind.CallOpts, _message [2]*big.Int, _signature [2]*big.Int, _indexes []*big.Int, _requiredVotingPower *big.Int, _minimumVotingPowerPerTaskDefinition *big.Int) error {
	var out []interface{}
	err := _ContractOBLS.contract.Call(opts, &out, "verifySignature", _message, _signature, _indexes, _requiredVotingPower, _minimumVotingPowerPerTaskDefinition)

	if err != nil {
		return err
	}

	return err

}

// VerifySignature is a free data retrieval call binding the contract method 0x65c46475.
//
// Solidity: function verifySignature(uint256[2] _message, uint256[2] _signature, uint256[] _indexes, uint256 _requiredVotingPower, uint256 _minimumVotingPowerPerTaskDefinition) view returns()
func (_ContractOBLS *ContractOBLSSession) VerifySignature(_message [2]*big.Int, _signature [2]*big.Int, _indexes []*big.Int, _requiredVotingPower *big.Int, _minimumVotingPowerPerTaskDefinition *big.Int) error {
	return _ContractOBLS.Contract.VerifySignature(&_ContractOBLS.CallOpts, _message, _signature, _indexes, _requiredVotingPower, _minimumVotingPowerPerTaskDefinition)
}

// VerifySignature is a free data retrieval call binding the contract method 0x65c46475.
//
// Solidity: function verifySignature(uint256[2] _message, uint256[2] _signature, uint256[] _indexes, uint256 _requiredVotingPower, uint256 _minimumVotingPowerPerTaskDefinition) view returns()
func (_ContractOBLS *ContractOBLSCallerSession) VerifySignature(_message [2]*big.Int, _signature [2]*big.Int, _indexes []*big.Int, _requiredVotingPower *big.Int, _minimumVotingPowerPerTaskDefinition *big.Int) error {
	return _ContractOBLS.Contract.VerifySignature(&_ContractOBLS.CallOpts, _message, _signature, _indexes, _requiredVotingPower, _minimumVotingPowerPerTaskDefinition)
}

// VotingPower is a free data retrieval call binding the contract method 0x72c4a927.
//
// Solidity: function votingPower(uint256 _index) view returns(uint256)
func (_ContractOBLS *ContractOBLSCaller) VotingPower(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ContractOBLS.contract.Call(opts, &out, "votingPower", _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingPower is a free data retrieval call binding the contract method 0x72c4a927.
//
// Solidity: function votingPower(uint256 _index) view returns(uint256)
func (_ContractOBLS *ContractOBLSSession) VotingPower(_index *big.Int) (*big.Int, error) {
	return _ContractOBLS.Contract.VotingPower(&_ContractOBLS.CallOpts, _index)
}

// VotingPower is a free data retrieval call binding the contract method 0x72c4a927.
//
// Solidity: function votingPower(uint256 _index) view returns(uint256)
func (_ContractOBLS *ContractOBLSCallerSession) VotingPower(_index *big.Int) (*big.Int, error) {
	return _ContractOBLS.Contract.VotingPower(&_ContractOBLS.CallOpts, _index)
}

// DecreaseBatchOperatorVotingPower is a paid mutator transaction binding the contract method 0x6533df08.
//
// Solidity: function decreaseBatchOperatorVotingPower((uint256,uint256)[] _operatorsVotingPower) returns()
func (_ContractOBLS *ContractOBLSTransactor) DecreaseBatchOperatorVotingPower(opts *bind.TransactOpts, _operatorsVotingPower []IOBLSOperatorVotingPower) (*types.Transaction, error) {
	return _ContractOBLS.contract.Transact(opts, "decreaseBatchOperatorVotingPower", _operatorsVotingPower)
}

// DecreaseBatchOperatorVotingPower is a paid mutator transaction binding the contract method 0x6533df08.
//
// Solidity: function decreaseBatchOperatorVotingPower((uint256,uint256)[] _operatorsVotingPower) returns()
func (_ContractOBLS *ContractOBLSSession) DecreaseBatchOperatorVotingPower(_operatorsVotingPower []IOBLSOperatorVotingPower) (*types.Transaction, error) {
	return _ContractOBLS.Contract.DecreaseBatchOperatorVotingPower(&_ContractOBLS.TransactOpts, _operatorsVotingPower)
}

// DecreaseBatchOperatorVotingPower is a paid mutator transaction binding the contract method 0x6533df08.
//
// Solidity: function decreaseBatchOperatorVotingPower((uint256,uint256)[] _operatorsVotingPower) returns()
func (_ContractOBLS *ContractOBLSTransactorSession) DecreaseBatchOperatorVotingPower(_operatorsVotingPower []IOBLSOperatorVotingPower) (*types.Transaction, error) {
	return _ContractOBLS.Contract.DecreaseBatchOperatorVotingPower(&_ContractOBLS.TransactOpts, _operatorsVotingPower)
}

// DecreaseOperatorVotingPower is a paid mutator transaction binding the contract method 0x20f527ad.
//
// Solidity: function decreaseOperatorVotingPower(uint256 _index, uint256 _votingPower) returns()
func (_ContractOBLS *ContractOBLSTransactor) DecreaseOperatorVotingPower(opts *bind.TransactOpts, _index *big.Int, _votingPower *big.Int) (*types.Transaction, error) {
	return _ContractOBLS.contract.Transact(opts, "decreaseOperatorVotingPower", _index, _votingPower)
}

// DecreaseOperatorVotingPower is a paid mutator transaction binding the contract method 0x20f527ad.
//
// Solidity: function decreaseOperatorVotingPower(uint256 _index, uint256 _votingPower) returns()
func (_ContractOBLS *ContractOBLSSession) DecreaseOperatorVotingPower(_index *big.Int, _votingPower *big.Int) (*types.Transaction, error) {
	return _ContractOBLS.Contract.DecreaseOperatorVotingPower(&_ContractOBLS.TransactOpts, _index, _votingPower)
}

// DecreaseOperatorVotingPower is a paid mutator transaction binding the contract method 0x20f527ad.
//
// Solidity: function decreaseOperatorVotingPower(uint256 _index, uint256 _votingPower) returns()
func (_ContractOBLS *ContractOBLSTransactorSession) DecreaseOperatorVotingPower(_index *big.Int, _votingPower *big.Int) (*types.Transaction, error) {
	return _ContractOBLS.Contract.DecreaseOperatorVotingPower(&_ContractOBLS.TransactOpts, _index, _votingPower)
}

// DecreaseOperatorVotingPowerPerTaskDefinition is a paid mutator transaction binding the contract method 0xff1e2c28.
//
// Solidity: function decreaseOperatorVotingPowerPerTaskDefinition(uint16 _taskDefinitionId, uint256 _votingPower) returns()
func (_ContractOBLS *ContractOBLSTransactor) DecreaseOperatorVotingPowerPerTaskDefinition(opts *bind.TransactOpts, _taskDefinitionId uint16, _votingPower *big.Int) (*types.Transaction, error) {
	return _ContractOBLS.contract.Transact(opts, "decreaseOperatorVotingPowerPerTaskDefinition", _taskDefinitionId, _votingPower)
}

// DecreaseOperatorVotingPowerPerTaskDefinition is a paid mutator transaction binding the contract method 0xff1e2c28.
//
// Solidity: function decreaseOperatorVotingPowerPerTaskDefinition(uint16 _taskDefinitionId, uint256 _votingPower) returns()
func (_ContractOBLS *ContractOBLSSession) DecreaseOperatorVotingPowerPerTaskDefinition(_taskDefinitionId uint16, _votingPower *big.Int) (*types.Transaction, error) {
	return _ContractOBLS.Contract.DecreaseOperatorVotingPowerPerTaskDefinition(&_ContractOBLS.TransactOpts, _taskDefinitionId, _votingPower)
}

// DecreaseOperatorVotingPowerPerTaskDefinition is a paid mutator transaction binding the contract method 0xff1e2c28.
//
// Solidity: function decreaseOperatorVotingPowerPerTaskDefinition(uint16 _taskDefinitionId, uint256 _votingPower) returns()
func (_ContractOBLS *ContractOBLSTransactorSession) DecreaseOperatorVotingPowerPerTaskDefinition(_taskDefinitionId uint16, _votingPower *big.Int) (*types.Transaction, error) {
	return _ContractOBLS.Contract.DecreaseOperatorVotingPowerPerTaskDefinition(&_ContractOBLS.TransactOpts, _taskDefinitionId, _votingPower)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ContractOBLS *ContractOBLSTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractOBLS.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ContractOBLS *ContractOBLSSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractOBLS.Contract.GrantRole(&_ContractOBLS.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ContractOBLS *ContractOBLSTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractOBLS.Contract.GrantRole(&_ContractOBLS.TransactOpts, role, account)
}

// IncreaseBatchOperatorVotingPower is a paid mutator transaction binding the contract method 0x63f903cd.
//
// Solidity: function increaseBatchOperatorVotingPower((uint256,uint256)[] _operatorsVotingPower) returns()
func (_ContractOBLS *ContractOBLSTransactor) IncreaseBatchOperatorVotingPower(opts *bind.TransactOpts, _operatorsVotingPower []IOBLSOperatorVotingPower) (*types.Transaction, error) {
	return _ContractOBLS.contract.Transact(opts, "increaseBatchOperatorVotingPower", _operatorsVotingPower)
}

// IncreaseBatchOperatorVotingPower is a paid mutator transaction binding the contract method 0x63f903cd.
//
// Solidity: function increaseBatchOperatorVotingPower((uint256,uint256)[] _operatorsVotingPower) returns()
func (_ContractOBLS *ContractOBLSSession) IncreaseBatchOperatorVotingPower(_operatorsVotingPower []IOBLSOperatorVotingPower) (*types.Transaction, error) {
	return _ContractOBLS.Contract.IncreaseBatchOperatorVotingPower(&_ContractOBLS.TransactOpts, _operatorsVotingPower)
}

// IncreaseBatchOperatorVotingPower is a paid mutator transaction binding the contract method 0x63f903cd.
//
// Solidity: function increaseBatchOperatorVotingPower((uint256,uint256)[] _operatorsVotingPower) returns()
func (_ContractOBLS *ContractOBLSTransactorSession) IncreaseBatchOperatorVotingPower(_operatorsVotingPower []IOBLSOperatorVotingPower) (*types.Transaction, error) {
	return _ContractOBLS.Contract.IncreaseBatchOperatorVotingPower(&_ContractOBLS.TransactOpts, _operatorsVotingPower)
}

// IncreaseOperatorVotingPower is a paid mutator transaction binding the contract method 0xd66f643d.
//
// Solidity: function increaseOperatorVotingPower(uint256 _index, uint256 _votingPower) returns()
func (_ContractOBLS *ContractOBLSTransactor) IncreaseOperatorVotingPower(opts *bind.TransactOpts, _index *big.Int, _votingPower *big.Int) (*types.Transaction, error) {
	return _ContractOBLS.contract.Transact(opts, "increaseOperatorVotingPower", _index, _votingPower)
}

// IncreaseOperatorVotingPower is a paid mutator transaction binding the contract method 0xd66f643d.
//
// Solidity: function increaseOperatorVotingPower(uint256 _index, uint256 _votingPower) returns()
func (_ContractOBLS *ContractOBLSSession) IncreaseOperatorVotingPower(_index *big.Int, _votingPower *big.Int) (*types.Transaction, error) {
	return _ContractOBLS.Contract.IncreaseOperatorVotingPower(&_ContractOBLS.TransactOpts, _index, _votingPower)
}

// IncreaseOperatorVotingPower is a paid mutator transaction binding the contract method 0xd66f643d.
//
// Solidity: function increaseOperatorVotingPower(uint256 _index, uint256 _votingPower) returns()
func (_ContractOBLS *ContractOBLSTransactorSession) IncreaseOperatorVotingPower(_index *big.Int, _votingPower *big.Int) (*types.Transaction, error) {
	return _ContractOBLS.Contract.IncreaseOperatorVotingPower(&_ContractOBLS.TransactOpts, _index, _votingPower)
}

// IncreaseOperatorVotingPowerPerTaskDefinition is a paid mutator transaction binding the contract method 0x93f438be.
//
// Solidity: function increaseOperatorVotingPowerPerTaskDefinition(uint16 _taskDefinitionId, uint256 _votingPower) returns()
func (_ContractOBLS *ContractOBLSTransactor) IncreaseOperatorVotingPowerPerTaskDefinition(opts *bind.TransactOpts, _taskDefinitionId uint16, _votingPower *big.Int) (*types.Transaction, error) {
	return _ContractOBLS.contract.Transact(opts, "increaseOperatorVotingPowerPerTaskDefinition", _taskDefinitionId, _votingPower)
}

// IncreaseOperatorVotingPowerPerTaskDefinition is a paid mutator transaction binding the contract method 0x93f438be.
//
// Solidity: function increaseOperatorVotingPowerPerTaskDefinition(uint16 _taskDefinitionId, uint256 _votingPower) returns()
func (_ContractOBLS *ContractOBLSSession) IncreaseOperatorVotingPowerPerTaskDefinition(_taskDefinitionId uint16, _votingPower *big.Int) (*types.Transaction, error) {
	return _ContractOBLS.Contract.IncreaseOperatorVotingPowerPerTaskDefinition(&_ContractOBLS.TransactOpts, _taskDefinitionId, _votingPower)
}

// IncreaseOperatorVotingPowerPerTaskDefinition is a paid mutator transaction binding the contract method 0x93f438be.
//
// Solidity: function increaseOperatorVotingPowerPerTaskDefinition(uint16 _taskDefinitionId, uint256 _votingPower) returns()
func (_ContractOBLS *ContractOBLSTransactorSession) IncreaseOperatorVotingPowerPerTaskDefinition(_taskDefinitionId uint16, _votingPower *big.Int) (*types.Transaction, error) {
	return _ContractOBLS.Contract.IncreaseOperatorVotingPowerPerTaskDefinition(&_ContractOBLS.TransactOpts, _taskDefinitionId, _votingPower)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ContractOBLS *ContractOBLSTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractOBLS.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ContractOBLS *ContractOBLSSession) Initialize() (*types.Transaction, error) {
	return _ContractOBLS.Contract.Initialize(&_ContractOBLS.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ContractOBLS *ContractOBLSTransactorSession) Initialize() (*types.Transaction, error) {
	return _ContractOBLS.Contract.Initialize(&_ContractOBLS.TransactOpts)
}

// Migration is a paid mutator transaction binding the contract method 0x1705a3bd.
//
// Solidity: function migration() returns()
func (_ContractOBLS *ContractOBLSTransactor) Migration(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractOBLS.contract.Transact(opts, "migration")
}

// Migration is a paid mutator transaction binding the contract method 0x1705a3bd.
//
// Solidity: function migration() returns()
func (_ContractOBLS *ContractOBLSSession) Migration() (*types.Transaction, error) {
	return _ContractOBLS.Contract.Migration(&_ContractOBLS.TransactOpts)
}

// Migration is a paid mutator transaction binding the contract method 0x1705a3bd.
//
// Solidity: function migration() returns()
func (_ContractOBLS *ContractOBLSTransactorSession) Migration() (*types.Transaction, error) {
	return _ContractOBLS.Contract.Migration(&_ContractOBLS.TransactOpts)
}

// ModifyOperatorActiveStatus is a paid mutator transaction binding the contract method 0x67ad3232.
//
// Solidity: function modifyOperatorActiveStatus(uint256 _index, bool _isRegistered) returns()
func (_ContractOBLS *ContractOBLSTransactor) ModifyOperatorActiveStatus(opts *bind.TransactOpts, _index *big.Int, _isRegistered bool) (*types.Transaction, error) {
	return _ContractOBLS.contract.Transact(opts, "modifyOperatorActiveStatus", _index, _isRegistered)
}

// ModifyOperatorActiveStatus is a paid mutator transaction binding the contract method 0x67ad3232.
//
// Solidity: function modifyOperatorActiveStatus(uint256 _index, bool _isRegistered) returns()
func (_ContractOBLS *ContractOBLSSession) ModifyOperatorActiveStatus(_index *big.Int, _isRegistered bool) (*types.Transaction, error) {
	return _ContractOBLS.Contract.ModifyOperatorActiveStatus(&_ContractOBLS.TransactOpts, _index, _isRegistered)
}

// ModifyOperatorActiveStatus is a paid mutator transaction binding the contract method 0x67ad3232.
//
// Solidity: function modifyOperatorActiveStatus(uint256 _index, bool _isRegistered) returns()
func (_ContractOBLS *ContractOBLSTransactorSession) ModifyOperatorActiveStatus(_index *big.Int, _isRegistered bool) (*types.Transaction, error) {
	return _ContractOBLS.Contract.ModifyOperatorActiveStatus(&_ContractOBLS.TransactOpts, _index, _isRegistered)
}

// ModifyOperatorBlsKey is a paid mutator transaction binding the contract method 0x86d897a7.
//
// Solidity: function modifyOperatorBlsKey(uint256 _index, uint256[4] _blsKey) returns()
func (_ContractOBLS *ContractOBLSTransactor) ModifyOperatorBlsKey(opts *bind.TransactOpts, _index *big.Int, _blsKey [4]*big.Int) (*types.Transaction, error) {
	return _ContractOBLS.contract.Transact(opts, "modifyOperatorBlsKey", _index, _blsKey)
}

// ModifyOperatorBlsKey is a paid mutator transaction binding the contract method 0x86d897a7.
//
// Solidity: function modifyOperatorBlsKey(uint256 _index, uint256[4] _blsKey) returns()
func (_ContractOBLS *ContractOBLSSession) ModifyOperatorBlsKey(_index *big.Int, _blsKey [4]*big.Int) (*types.Transaction, error) {
	return _ContractOBLS.Contract.ModifyOperatorBlsKey(&_ContractOBLS.TransactOpts, _index, _blsKey)
}

// ModifyOperatorBlsKey is a paid mutator transaction binding the contract method 0x86d897a7.
//
// Solidity: function modifyOperatorBlsKey(uint256 _index, uint256[4] _blsKey) returns()
func (_ContractOBLS *ContractOBLSTransactorSession) ModifyOperatorBlsKey(_index *big.Int, _blsKey [4]*big.Int) (*types.Transaction, error) {
	return _ContractOBLS.Contract.ModifyOperatorBlsKey(&_ContractOBLS.TransactOpts, _index, _blsKey)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x891a80bb.
//
// Solidity: function registerOperator(uint256 _index, uint256 _votingPower, uint256[4] _blsKey) returns()
func (_ContractOBLS *ContractOBLSTransactor) RegisterOperator(opts *bind.TransactOpts, _index *big.Int, _votingPower *big.Int, _blsKey [4]*big.Int) (*types.Transaction, error) {
	return _ContractOBLS.contract.Transact(opts, "registerOperator", _index, _votingPower, _blsKey)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x891a80bb.
//
// Solidity: function registerOperator(uint256 _index, uint256 _votingPower, uint256[4] _blsKey) returns()
func (_ContractOBLS *ContractOBLSSession) RegisterOperator(_index *big.Int, _votingPower *big.Int, _blsKey [4]*big.Int) (*types.Transaction, error) {
	return _ContractOBLS.Contract.RegisterOperator(&_ContractOBLS.TransactOpts, _index, _votingPower, _blsKey)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x891a80bb.
//
// Solidity: function registerOperator(uint256 _index, uint256 _votingPower, uint256[4] _blsKey) returns()
func (_ContractOBLS *ContractOBLSTransactorSession) RegisterOperator(_index *big.Int, _votingPower *big.Int, _blsKey [4]*big.Int) (*types.Transaction, error) {
	return _ContractOBLS.Contract.RegisterOperator(&_ContractOBLS.TransactOpts, _index, _votingPower, _blsKey)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ContractOBLS *ContractOBLSTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ContractOBLS.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ContractOBLS *ContractOBLSSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ContractOBLS.Contract.RenounceRole(&_ContractOBLS.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ContractOBLS *ContractOBLSTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ContractOBLS.Contract.RenounceRole(&_ContractOBLS.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ContractOBLS *ContractOBLSTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractOBLS.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ContractOBLS *ContractOBLSSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractOBLS.Contract.RevokeRole(&_ContractOBLS.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ContractOBLS *ContractOBLSTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractOBLS.Contract.RevokeRole(&_ContractOBLS.TransactOpts, role, account)
}

// SetOblsManager is a paid mutator transaction binding the contract method 0xd9835b61.
//
// Solidity: function setOblsManager(address _oblsManager) returns()
func (_ContractOBLS *ContractOBLSTransactor) SetOblsManager(opts *bind.TransactOpts, _oblsManager common.Address) (*types.Transaction, error) {
	return _ContractOBLS.contract.Transact(opts, "setOblsManager", _oblsManager)
}

// SetOblsManager is a paid mutator transaction binding the contract method 0xd9835b61.
//
// Solidity: function setOblsManager(address _oblsManager) returns()
func (_ContractOBLS *ContractOBLSSession) SetOblsManager(_oblsManager common.Address) (*types.Transaction, error) {
	return _ContractOBLS.Contract.SetOblsManager(&_ContractOBLS.TransactOpts, _oblsManager)
}

// SetOblsManager is a paid mutator transaction binding the contract method 0xd9835b61.
//
// Solidity: function setOblsManager(address _oblsManager) returns()
func (_ContractOBLS *ContractOBLSTransactorSession) SetOblsManager(_oblsManager common.Address) (*types.Transaction, error) {
	return _ContractOBLS.Contract.SetOblsManager(&_ContractOBLS.TransactOpts, _oblsManager)
}

// SetOblsSharesSyncer is a paid mutator transaction binding the contract method 0x1164224e.
//
// Solidity: function setOblsSharesSyncer(address _oblsSharesSyncer) returns()
func (_ContractOBLS *ContractOBLSTransactor) SetOblsSharesSyncer(opts *bind.TransactOpts, _oblsSharesSyncer common.Address) (*types.Transaction, error) {
	return _ContractOBLS.contract.Transact(opts, "setOblsSharesSyncer", _oblsSharesSyncer)
}

// SetOblsSharesSyncer is a paid mutator transaction binding the contract method 0x1164224e.
//
// Solidity: function setOblsSharesSyncer(address _oblsSharesSyncer) returns()
func (_ContractOBLS *ContractOBLSSession) SetOblsSharesSyncer(_oblsSharesSyncer common.Address) (*types.Transaction, error) {
	return _ContractOBLS.Contract.SetOblsSharesSyncer(&_ContractOBLS.TransactOpts, _oblsSharesSyncer)
}

// SetOblsSharesSyncer is a paid mutator transaction binding the contract method 0x1164224e.
//
// Solidity: function setOblsSharesSyncer(address _oblsSharesSyncer) returns()
func (_ContractOBLS *ContractOBLSTransactorSession) SetOblsSharesSyncer(_oblsSharesSyncer common.Address) (*types.Transaction, error) {
	return _ContractOBLS.Contract.SetOblsSharesSyncer(&_ContractOBLS.TransactOpts, _oblsSharesSyncer)
}

// SetTotalVotingPowerPerRestrictedTaskDefinition is a paid mutator transaction binding the contract method 0xca87bf8f.
//
// Solidity: function setTotalVotingPowerPerRestrictedTaskDefinition(uint16 _taskDefinitionId, uint256 _minimumVotingPower, uint256[] _restrictedAttesterIds) returns()
func (_ContractOBLS *ContractOBLSTransactor) SetTotalVotingPowerPerRestrictedTaskDefinition(opts *bind.TransactOpts, _taskDefinitionId uint16, _minimumVotingPower *big.Int, _restrictedAttesterIds []*big.Int) (*types.Transaction, error) {
	return _ContractOBLS.contract.Transact(opts, "setTotalVotingPowerPerRestrictedTaskDefinition", _taskDefinitionId, _minimumVotingPower, _restrictedAttesterIds)
}

// SetTotalVotingPowerPerRestrictedTaskDefinition is a paid mutator transaction binding the contract method 0xca87bf8f.
//
// Solidity: function setTotalVotingPowerPerRestrictedTaskDefinition(uint16 _taskDefinitionId, uint256 _minimumVotingPower, uint256[] _restrictedAttesterIds) returns()
func (_ContractOBLS *ContractOBLSSession) SetTotalVotingPowerPerRestrictedTaskDefinition(_taskDefinitionId uint16, _minimumVotingPower *big.Int, _restrictedAttesterIds []*big.Int) (*types.Transaction, error) {
	return _ContractOBLS.Contract.SetTotalVotingPowerPerRestrictedTaskDefinition(&_ContractOBLS.TransactOpts, _taskDefinitionId, _minimumVotingPower, _restrictedAttesterIds)
}

// SetTotalVotingPowerPerRestrictedTaskDefinition is a paid mutator transaction binding the contract method 0xca87bf8f.
//
// Solidity: function setTotalVotingPowerPerRestrictedTaskDefinition(uint16 _taskDefinitionId, uint256 _minimumVotingPower, uint256[] _restrictedAttesterIds) returns()
func (_ContractOBLS *ContractOBLSTransactorSession) SetTotalVotingPowerPerRestrictedTaskDefinition(_taskDefinitionId uint16, _minimumVotingPower *big.Int, _restrictedAttesterIds []*big.Int) (*types.Transaction, error) {
	return _ContractOBLS.Contract.SetTotalVotingPowerPerRestrictedTaskDefinition(&_ContractOBLS.TransactOpts, _taskDefinitionId, _minimumVotingPower, _restrictedAttesterIds)
}

// SetTotalVotingPowerPerTaskDefinition is a paid mutator transaction binding the contract method 0xe010f957.
//
// Solidity: function setTotalVotingPowerPerTaskDefinition(uint16 _taskDefinitionId, uint256 _numOfTotalOperators, uint256 _minimumVotingPower) returns()
func (_ContractOBLS *ContractOBLSTransactor) SetTotalVotingPowerPerTaskDefinition(opts *bind.TransactOpts, _taskDefinitionId uint16, _numOfTotalOperators *big.Int, _minimumVotingPower *big.Int) (*types.Transaction, error) {
	return _ContractOBLS.contract.Transact(opts, "setTotalVotingPowerPerTaskDefinition", _taskDefinitionId, _numOfTotalOperators, _minimumVotingPower)
}

// SetTotalVotingPowerPerTaskDefinition is a paid mutator transaction binding the contract method 0xe010f957.
//
// Solidity: function setTotalVotingPowerPerTaskDefinition(uint16 _taskDefinitionId, uint256 _numOfTotalOperators, uint256 _minimumVotingPower) returns()
func (_ContractOBLS *ContractOBLSSession) SetTotalVotingPowerPerTaskDefinition(_taskDefinitionId uint16, _numOfTotalOperators *big.Int, _minimumVotingPower *big.Int) (*types.Transaction, error) {
	return _ContractOBLS.Contract.SetTotalVotingPowerPerTaskDefinition(&_ContractOBLS.TransactOpts, _taskDefinitionId, _numOfTotalOperators, _minimumVotingPower)
}

// SetTotalVotingPowerPerTaskDefinition is a paid mutator transaction binding the contract method 0xe010f957.
//
// Solidity: function setTotalVotingPowerPerTaskDefinition(uint16 _taskDefinitionId, uint256 _numOfTotalOperators, uint256 _minimumVotingPower) returns()
func (_ContractOBLS *ContractOBLSTransactorSession) SetTotalVotingPowerPerTaskDefinition(_taskDefinitionId uint16, _numOfTotalOperators *big.Int, _minimumVotingPower *big.Int) (*types.Transaction, error) {
	return _ContractOBLS.Contract.SetTotalVotingPowerPerTaskDefinition(&_ContractOBLS.TransactOpts, _taskDefinitionId, _numOfTotalOperators, _minimumVotingPower)
}

// SetVotingPowerSyncer is a paid mutator transaction binding the contract method 0x6c778aae.
//
// Solidity: function setVotingPowerSyncer(address _votingPowerSyncer) returns()
func (_ContractOBLS *ContractOBLSTransactor) SetVotingPowerSyncer(opts *bind.TransactOpts, _votingPowerSyncer common.Address) (*types.Transaction, error) {
	return _ContractOBLS.contract.Transact(opts, "setVotingPowerSyncer", _votingPowerSyncer)
}

// SetVotingPowerSyncer is a paid mutator transaction binding the contract method 0x6c778aae.
//
// Solidity: function setVotingPowerSyncer(address _votingPowerSyncer) returns()
func (_ContractOBLS *ContractOBLSSession) SetVotingPowerSyncer(_votingPowerSyncer common.Address) (*types.Transaction, error) {
	return _ContractOBLS.Contract.SetVotingPowerSyncer(&_ContractOBLS.TransactOpts, _votingPowerSyncer)
}

// SetVotingPowerSyncer is a paid mutator transaction binding the contract method 0x6c778aae.
//
// Solidity: function setVotingPowerSyncer(address _votingPowerSyncer) returns()
func (_ContractOBLS *ContractOBLSTransactorSession) SetVotingPowerSyncer(_votingPowerSyncer common.Address) (*types.Transaction, error) {
	return _ContractOBLS.Contract.SetVotingPowerSyncer(&_ContractOBLS.TransactOpts, _votingPowerSyncer)
}

// SyncOperatorDetails is a paid mutator transaction binding the contract method 0xc1515f8b.
//
// Solidity: function syncOperatorDetails(uint256[] _votingPowers, uint256[4][] _blsKeys) returns()
func (_ContractOBLS *ContractOBLSTransactor) SyncOperatorDetails(opts *bind.TransactOpts, _votingPowers []*big.Int, _blsKeys [][4]*big.Int) (*types.Transaction, error) {
	return _ContractOBLS.contract.Transact(opts, "syncOperatorDetails", _votingPowers, _blsKeys)
}

// SyncOperatorDetails is a paid mutator transaction binding the contract method 0xc1515f8b.
//
// Solidity: function syncOperatorDetails(uint256[] _votingPowers, uint256[4][] _blsKeys) returns()
func (_ContractOBLS *ContractOBLSSession) SyncOperatorDetails(_votingPowers []*big.Int, _blsKeys [][4]*big.Int) (*types.Transaction, error) {
	return _ContractOBLS.Contract.SyncOperatorDetails(&_ContractOBLS.TransactOpts, _votingPowers, _blsKeys)
}

// SyncOperatorDetails is a paid mutator transaction binding the contract method 0xc1515f8b.
//
// Solidity: function syncOperatorDetails(uint256[] _votingPowers, uint256[4][] _blsKeys) returns()
func (_ContractOBLS *ContractOBLSTransactorSession) SyncOperatorDetails(_votingPowers []*big.Int, _blsKeys [][4]*big.Int) (*types.Transaction, error) {
	return _ContractOBLS.Contract.SyncOperatorDetails(&_ContractOBLS.TransactOpts, _votingPowers, _blsKeys)
}

// UnRegisterOperator is a paid mutator transaction binding the contract method 0xcf33bccc.
//
// Solidity: function unRegisterOperator(uint256 _index) returns()
func (_ContractOBLS *ContractOBLSTransactor) UnRegisterOperator(opts *bind.TransactOpts, _index *big.Int) (*types.Transaction, error) {
	return _ContractOBLS.contract.Transact(opts, "unRegisterOperator", _index)
}

// UnRegisterOperator is a paid mutator transaction binding the contract method 0xcf33bccc.
//
// Solidity: function unRegisterOperator(uint256 _index) returns()
func (_ContractOBLS *ContractOBLSSession) UnRegisterOperator(_index *big.Int) (*types.Transaction, error) {
	return _ContractOBLS.Contract.UnRegisterOperator(&_ContractOBLS.TransactOpts, _index)
}

// UnRegisterOperator is a paid mutator transaction binding the contract method 0xcf33bccc.
//
// Solidity: function unRegisterOperator(uint256 _index) returns()
func (_ContractOBLS *ContractOBLSTransactorSession) UnRegisterOperator(_index *big.Int) (*types.Transaction, error) {
	return _ContractOBLS.Contract.UnRegisterOperator(&_ContractOBLS.TransactOpts, _index)
}

// ContractOBLSDecreaseBatchOperatorVotingPowerIterator is returned from FilterDecreaseBatchOperatorVotingPower and is used to iterate over the raw logs and unpacked data for DecreaseBatchOperatorVotingPower events raised by the ContractOBLS contract.
type ContractOBLSDecreaseBatchOperatorVotingPowerIterator struct {
	Event *ContractOBLSDecreaseBatchOperatorVotingPower // Event containing the contract specifics and raw log

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
func (it *ContractOBLSDecreaseBatchOperatorVotingPowerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOBLSDecreaseBatchOperatorVotingPower)
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
		it.Event = new(ContractOBLSDecreaseBatchOperatorVotingPower)
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
func (it *ContractOBLSDecreaseBatchOperatorVotingPowerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOBLSDecreaseBatchOperatorVotingPowerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOBLSDecreaseBatchOperatorVotingPower represents a DecreaseBatchOperatorVotingPower event raised by the ContractOBLS contract.
type ContractOBLSDecreaseBatchOperatorVotingPower struct {
	OperatorsVotingPower []IOBLSOperatorVotingPower
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterDecreaseBatchOperatorVotingPower is a free log retrieval operation binding the contract event 0x66c9c86bb0b2a5e43f482b8ae3ec409cd83c878dca97ceeb2024a49a344e431c.
//
// Solidity: event DecreaseBatchOperatorVotingPower((uint256,uint256)[] operatorsVotingPower)
func (_ContractOBLS *ContractOBLSFilterer) FilterDecreaseBatchOperatorVotingPower(opts *bind.FilterOpts) (*ContractOBLSDecreaseBatchOperatorVotingPowerIterator, error) {

	logs, sub, err := _ContractOBLS.contract.FilterLogs(opts, "DecreaseBatchOperatorVotingPower")
	if err != nil {
		return nil, err
	}
	return &ContractOBLSDecreaseBatchOperatorVotingPowerIterator{contract: _ContractOBLS.contract, event: "DecreaseBatchOperatorVotingPower", logs: logs, sub: sub}, nil
}

// WatchDecreaseBatchOperatorVotingPower is a free log subscription operation binding the contract event 0x66c9c86bb0b2a5e43f482b8ae3ec409cd83c878dca97ceeb2024a49a344e431c.
//
// Solidity: event DecreaseBatchOperatorVotingPower((uint256,uint256)[] operatorsVotingPower)
func (_ContractOBLS *ContractOBLSFilterer) WatchDecreaseBatchOperatorVotingPower(opts *bind.WatchOpts, sink chan<- *ContractOBLSDecreaseBatchOperatorVotingPower) (event.Subscription, error) {

	logs, sub, err := _ContractOBLS.contract.WatchLogs(opts, "DecreaseBatchOperatorVotingPower")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOBLSDecreaseBatchOperatorVotingPower)
				if err := _ContractOBLS.contract.UnpackLog(event, "DecreaseBatchOperatorVotingPower", log); err != nil {
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

// ParseDecreaseBatchOperatorVotingPower is a log parse operation binding the contract event 0x66c9c86bb0b2a5e43f482b8ae3ec409cd83c878dca97ceeb2024a49a344e431c.
//
// Solidity: event DecreaseBatchOperatorVotingPower((uint256,uint256)[] operatorsVotingPower)
func (_ContractOBLS *ContractOBLSFilterer) ParseDecreaseBatchOperatorVotingPower(log types.Log) (*ContractOBLSDecreaseBatchOperatorVotingPower, error) {
	event := new(ContractOBLSDecreaseBatchOperatorVotingPower)
	if err := _ContractOBLS.contract.UnpackLog(event, "DecreaseBatchOperatorVotingPower", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOBLSIncreaseBatchOperatorVotingPowerIterator is returned from FilterIncreaseBatchOperatorVotingPower and is used to iterate over the raw logs and unpacked data for IncreaseBatchOperatorVotingPower events raised by the ContractOBLS contract.
type ContractOBLSIncreaseBatchOperatorVotingPowerIterator struct {
	Event *ContractOBLSIncreaseBatchOperatorVotingPower // Event containing the contract specifics and raw log

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
func (it *ContractOBLSIncreaseBatchOperatorVotingPowerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOBLSIncreaseBatchOperatorVotingPower)
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
		it.Event = new(ContractOBLSIncreaseBatchOperatorVotingPower)
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
func (it *ContractOBLSIncreaseBatchOperatorVotingPowerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOBLSIncreaseBatchOperatorVotingPowerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOBLSIncreaseBatchOperatorVotingPower represents a IncreaseBatchOperatorVotingPower event raised by the ContractOBLS contract.
type ContractOBLSIncreaseBatchOperatorVotingPower struct {
	OperatorsVotingPower []IOBLSOperatorVotingPower
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterIncreaseBatchOperatorVotingPower is a free log retrieval operation binding the contract event 0xcb3fc06dd5b9108c2ccfca6812e638086d45dd7119dd53964eb3524e11a4b113.
//
// Solidity: event IncreaseBatchOperatorVotingPower((uint256,uint256)[] operatorsVotingPower)
func (_ContractOBLS *ContractOBLSFilterer) FilterIncreaseBatchOperatorVotingPower(opts *bind.FilterOpts) (*ContractOBLSIncreaseBatchOperatorVotingPowerIterator, error) {

	logs, sub, err := _ContractOBLS.contract.FilterLogs(opts, "IncreaseBatchOperatorVotingPower")
	if err != nil {
		return nil, err
	}
	return &ContractOBLSIncreaseBatchOperatorVotingPowerIterator{contract: _ContractOBLS.contract, event: "IncreaseBatchOperatorVotingPower", logs: logs, sub: sub}, nil
}

// WatchIncreaseBatchOperatorVotingPower is a free log subscription operation binding the contract event 0xcb3fc06dd5b9108c2ccfca6812e638086d45dd7119dd53964eb3524e11a4b113.
//
// Solidity: event IncreaseBatchOperatorVotingPower((uint256,uint256)[] operatorsVotingPower)
func (_ContractOBLS *ContractOBLSFilterer) WatchIncreaseBatchOperatorVotingPower(opts *bind.WatchOpts, sink chan<- *ContractOBLSIncreaseBatchOperatorVotingPower) (event.Subscription, error) {

	logs, sub, err := _ContractOBLS.contract.WatchLogs(opts, "IncreaseBatchOperatorVotingPower")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOBLSIncreaseBatchOperatorVotingPower)
				if err := _ContractOBLS.contract.UnpackLog(event, "IncreaseBatchOperatorVotingPower", log); err != nil {
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

// ParseIncreaseBatchOperatorVotingPower is a log parse operation binding the contract event 0xcb3fc06dd5b9108c2ccfca6812e638086d45dd7119dd53964eb3524e11a4b113.
//
// Solidity: event IncreaseBatchOperatorVotingPower((uint256,uint256)[] operatorsVotingPower)
func (_ContractOBLS *ContractOBLSFilterer) ParseIncreaseBatchOperatorVotingPower(log types.Log) (*ContractOBLSIncreaseBatchOperatorVotingPower, error) {
	event := new(ContractOBLSIncreaseBatchOperatorVotingPower)
	if err := _ContractOBLS.contract.UnpackLog(event, "IncreaseBatchOperatorVotingPower", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOBLSInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractOBLS contract.
type ContractOBLSInitializedIterator struct {
	Event *ContractOBLSInitialized // Event containing the contract specifics and raw log

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
func (it *ContractOBLSInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOBLSInitialized)
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
		it.Event = new(ContractOBLSInitialized)
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
func (it *ContractOBLSInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOBLSInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOBLSInitialized represents a Initialized event raised by the ContractOBLS contract.
type ContractOBLSInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ContractOBLS *ContractOBLSFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractOBLSInitializedIterator, error) {

	logs, sub, err := _ContractOBLS.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractOBLSInitializedIterator{contract: _ContractOBLS.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ContractOBLS *ContractOBLSFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractOBLSInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractOBLS.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOBLSInitialized)
				if err := _ContractOBLS.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractOBLS *ContractOBLSFilterer) ParseInitialized(log types.Log) (*ContractOBLSInitialized, error) {
	event := new(ContractOBLSInitialized)
	if err := _ContractOBLS.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOBLSRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ContractOBLS contract.
type ContractOBLSRoleAdminChangedIterator struct {
	Event *ContractOBLSRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ContractOBLSRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOBLSRoleAdminChanged)
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
		it.Event = new(ContractOBLSRoleAdminChanged)
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
func (it *ContractOBLSRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOBLSRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOBLSRoleAdminChanged represents a RoleAdminChanged event raised by the ContractOBLS contract.
type ContractOBLSRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ContractOBLS *ContractOBLSFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ContractOBLSRoleAdminChangedIterator, error) {

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

	logs, sub, err := _ContractOBLS.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ContractOBLSRoleAdminChangedIterator{contract: _ContractOBLS.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ContractOBLS *ContractOBLSFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ContractOBLSRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _ContractOBLS.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOBLSRoleAdminChanged)
				if err := _ContractOBLS.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_ContractOBLS *ContractOBLSFilterer) ParseRoleAdminChanged(log types.Log) (*ContractOBLSRoleAdminChanged, error) {
	event := new(ContractOBLSRoleAdminChanged)
	if err := _ContractOBLS.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOBLSRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ContractOBLS contract.
type ContractOBLSRoleGrantedIterator struct {
	Event *ContractOBLSRoleGranted // Event containing the contract specifics and raw log

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
func (it *ContractOBLSRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOBLSRoleGranted)
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
		it.Event = new(ContractOBLSRoleGranted)
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
func (it *ContractOBLSRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOBLSRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOBLSRoleGranted represents a RoleGranted event raised by the ContractOBLS contract.
type ContractOBLSRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ContractOBLS *ContractOBLSFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ContractOBLSRoleGrantedIterator, error) {

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

	logs, sub, err := _ContractOBLS.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractOBLSRoleGrantedIterator{contract: _ContractOBLS.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ContractOBLS *ContractOBLSFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ContractOBLSRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ContractOBLS.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOBLSRoleGranted)
				if err := _ContractOBLS.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_ContractOBLS *ContractOBLSFilterer) ParseRoleGranted(log types.Log) (*ContractOBLSRoleGranted, error) {
	event := new(ContractOBLSRoleGranted)
	if err := _ContractOBLS.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOBLSRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ContractOBLS contract.
type ContractOBLSRoleRevokedIterator struct {
	Event *ContractOBLSRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ContractOBLSRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOBLSRoleRevoked)
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
		it.Event = new(ContractOBLSRoleRevoked)
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
func (it *ContractOBLSRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOBLSRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOBLSRoleRevoked represents a RoleRevoked event raised by the ContractOBLS contract.
type ContractOBLSRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ContractOBLS *ContractOBLSFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ContractOBLSRoleRevokedIterator, error) {

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

	logs, sub, err := _ContractOBLS.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractOBLSRoleRevokedIterator{contract: _ContractOBLS.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ContractOBLS *ContractOBLSFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ContractOBLSRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ContractOBLS.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOBLSRoleRevoked)
				if err := _ContractOBLS.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_ContractOBLS *ContractOBLSFilterer) ParseRoleRevoked(log types.Log) (*ContractOBLSRoleRevoked, error) {
	event := new(ContractOBLSRoleRevoked)
	if err := _ContractOBLS.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOBLSSetOBLSManagerIterator is returned from FilterSetOBLSManager and is used to iterate over the raw logs and unpacked data for SetOBLSManager events raised by the ContractOBLS contract.
type ContractOBLSSetOBLSManagerIterator struct {
	Event *ContractOBLSSetOBLSManager // Event containing the contract specifics and raw log

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
func (it *ContractOBLSSetOBLSManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOBLSSetOBLSManager)
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
		it.Event = new(ContractOBLSSetOBLSManager)
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
func (it *ContractOBLSSetOBLSManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOBLSSetOBLSManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOBLSSetOBLSManager represents a SetOBLSManager event raised by the ContractOBLS contract.
type ContractOBLSSetOBLSManager struct {
	NewOBLSManager common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSetOBLSManager is a free log retrieval operation binding the contract event 0xff7dfc53b4a07266bc3bbaabbdb9992fc67be75384df0de05509a5cfdae75101.
//
// Solidity: event SetOBLSManager(address newOBLSManager)
func (_ContractOBLS *ContractOBLSFilterer) FilterSetOBLSManager(opts *bind.FilterOpts) (*ContractOBLSSetOBLSManagerIterator, error) {

	logs, sub, err := _ContractOBLS.contract.FilterLogs(opts, "SetOBLSManager")
	if err != nil {
		return nil, err
	}
	return &ContractOBLSSetOBLSManagerIterator{contract: _ContractOBLS.contract, event: "SetOBLSManager", logs: logs, sub: sub}, nil
}

// WatchSetOBLSManager is a free log subscription operation binding the contract event 0xff7dfc53b4a07266bc3bbaabbdb9992fc67be75384df0de05509a5cfdae75101.
//
// Solidity: event SetOBLSManager(address newOBLSManager)
func (_ContractOBLS *ContractOBLSFilterer) WatchSetOBLSManager(opts *bind.WatchOpts, sink chan<- *ContractOBLSSetOBLSManager) (event.Subscription, error) {

	logs, sub, err := _ContractOBLS.contract.WatchLogs(opts, "SetOBLSManager")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOBLSSetOBLSManager)
				if err := _ContractOBLS.contract.UnpackLog(event, "SetOBLSManager", log); err != nil {
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

// ParseSetOBLSManager is a log parse operation binding the contract event 0xff7dfc53b4a07266bc3bbaabbdb9992fc67be75384df0de05509a5cfdae75101.
//
// Solidity: event SetOBLSManager(address newOBLSManager)
func (_ContractOBLS *ContractOBLSFilterer) ParseSetOBLSManager(log types.Log) (*ContractOBLSSetOBLSManager, error) {
	event := new(ContractOBLSSetOBLSManager)
	if err := _ContractOBLS.contract.UnpackLog(event, "SetOBLSManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOBLSSetTotalVotingPowerPerRestrictedTaskDefinitionIterator is returned from FilterSetTotalVotingPowerPerRestrictedTaskDefinition and is used to iterate over the raw logs and unpacked data for SetTotalVotingPowerPerRestrictedTaskDefinition events raised by the ContractOBLS contract.
type ContractOBLSSetTotalVotingPowerPerRestrictedTaskDefinitionIterator struct {
	Event *ContractOBLSSetTotalVotingPowerPerRestrictedTaskDefinition // Event containing the contract specifics and raw log

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
func (it *ContractOBLSSetTotalVotingPowerPerRestrictedTaskDefinitionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOBLSSetTotalVotingPowerPerRestrictedTaskDefinition)
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
		it.Event = new(ContractOBLSSetTotalVotingPowerPerRestrictedTaskDefinition)
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
func (it *ContractOBLSSetTotalVotingPowerPerRestrictedTaskDefinitionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOBLSSetTotalVotingPowerPerRestrictedTaskDefinitionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOBLSSetTotalVotingPowerPerRestrictedTaskDefinition represents a SetTotalVotingPowerPerRestrictedTaskDefinition event raised by the ContractOBLS contract.
type ContractOBLSSetTotalVotingPowerPerRestrictedTaskDefinition struct {
	TaskDefinitionId      uint16
	MinimumVotingPower    *big.Int
	RestrictedAttesterIds []*big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSetTotalVotingPowerPerRestrictedTaskDefinition is a free log retrieval operation binding the contract event 0xc32e19b0b3c968ec2ad3ecf6420cfeac8f1a79d7870c6da28207d5ea499fde15.
//
// Solidity: event SetTotalVotingPowerPerRestrictedTaskDefinition(uint16 taskDefinitionId, uint256 minimumVotingPower, uint256[] restrictedAttesterIds)
func (_ContractOBLS *ContractOBLSFilterer) FilterSetTotalVotingPowerPerRestrictedTaskDefinition(opts *bind.FilterOpts) (*ContractOBLSSetTotalVotingPowerPerRestrictedTaskDefinitionIterator, error) {

	logs, sub, err := _ContractOBLS.contract.FilterLogs(opts, "SetTotalVotingPowerPerRestrictedTaskDefinition")
	if err != nil {
		return nil, err
	}
	return &ContractOBLSSetTotalVotingPowerPerRestrictedTaskDefinitionIterator{contract: _ContractOBLS.contract, event: "SetTotalVotingPowerPerRestrictedTaskDefinition", logs: logs, sub: sub}, nil
}

// WatchSetTotalVotingPowerPerRestrictedTaskDefinition is a free log subscription operation binding the contract event 0xc32e19b0b3c968ec2ad3ecf6420cfeac8f1a79d7870c6da28207d5ea499fde15.
//
// Solidity: event SetTotalVotingPowerPerRestrictedTaskDefinition(uint16 taskDefinitionId, uint256 minimumVotingPower, uint256[] restrictedAttesterIds)
func (_ContractOBLS *ContractOBLSFilterer) WatchSetTotalVotingPowerPerRestrictedTaskDefinition(opts *bind.WatchOpts, sink chan<- *ContractOBLSSetTotalVotingPowerPerRestrictedTaskDefinition) (event.Subscription, error) {

	logs, sub, err := _ContractOBLS.contract.WatchLogs(opts, "SetTotalVotingPowerPerRestrictedTaskDefinition")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOBLSSetTotalVotingPowerPerRestrictedTaskDefinition)
				if err := _ContractOBLS.contract.UnpackLog(event, "SetTotalVotingPowerPerRestrictedTaskDefinition", log); err != nil {
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

// ParseSetTotalVotingPowerPerRestrictedTaskDefinition is a log parse operation binding the contract event 0xc32e19b0b3c968ec2ad3ecf6420cfeac8f1a79d7870c6da28207d5ea499fde15.
//
// Solidity: event SetTotalVotingPowerPerRestrictedTaskDefinition(uint16 taskDefinitionId, uint256 minimumVotingPower, uint256[] restrictedAttesterIds)
func (_ContractOBLS *ContractOBLSFilterer) ParseSetTotalVotingPowerPerRestrictedTaskDefinition(log types.Log) (*ContractOBLSSetTotalVotingPowerPerRestrictedTaskDefinition, error) {
	event := new(ContractOBLSSetTotalVotingPowerPerRestrictedTaskDefinition)
	if err := _ContractOBLS.contract.UnpackLog(event, "SetTotalVotingPowerPerRestrictedTaskDefinition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOBLSSetTotalVotingPowerPerTaskDefinitionIterator is returned from FilterSetTotalVotingPowerPerTaskDefinition and is used to iterate over the raw logs and unpacked data for SetTotalVotingPowerPerTaskDefinition events raised by the ContractOBLS contract.
type ContractOBLSSetTotalVotingPowerPerTaskDefinitionIterator struct {
	Event *ContractOBLSSetTotalVotingPowerPerTaskDefinition // Event containing the contract specifics and raw log

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
func (it *ContractOBLSSetTotalVotingPowerPerTaskDefinitionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOBLSSetTotalVotingPowerPerTaskDefinition)
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
		it.Event = new(ContractOBLSSetTotalVotingPowerPerTaskDefinition)
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
func (it *ContractOBLSSetTotalVotingPowerPerTaskDefinitionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOBLSSetTotalVotingPowerPerTaskDefinitionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOBLSSetTotalVotingPowerPerTaskDefinition represents a SetTotalVotingPowerPerTaskDefinition event raised by the ContractOBLS contract.
type ContractOBLSSetTotalVotingPowerPerTaskDefinition struct {
	TaskdefinitionId    uint16
	NumOfTotalOperators *big.Int
	MinimumVotingPower  *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetTotalVotingPowerPerTaskDefinition is a free log retrieval operation binding the contract event 0x5a0cbf3f4ff24ea3dcc0680c9e2ab134f24e152cf4382343908507e4338d4b4c.
//
// Solidity: event SetTotalVotingPowerPerTaskDefinition(uint16 taskdefinitionId, uint256 numOfTotalOperators, uint256 minimumVotingPower)
func (_ContractOBLS *ContractOBLSFilterer) FilterSetTotalVotingPowerPerTaskDefinition(opts *bind.FilterOpts) (*ContractOBLSSetTotalVotingPowerPerTaskDefinitionIterator, error) {

	logs, sub, err := _ContractOBLS.contract.FilterLogs(opts, "SetTotalVotingPowerPerTaskDefinition")
	if err != nil {
		return nil, err
	}
	return &ContractOBLSSetTotalVotingPowerPerTaskDefinitionIterator{contract: _ContractOBLS.contract, event: "SetTotalVotingPowerPerTaskDefinition", logs: logs, sub: sub}, nil
}

// WatchSetTotalVotingPowerPerTaskDefinition is a free log subscription operation binding the contract event 0x5a0cbf3f4ff24ea3dcc0680c9e2ab134f24e152cf4382343908507e4338d4b4c.
//
// Solidity: event SetTotalVotingPowerPerTaskDefinition(uint16 taskdefinitionId, uint256 numOfTotalOperators, uint256 minimumVotingPower)
func (_ContractOBLS *ContractOBLSFilterer) WatchSetTotalVotingPowerPerTaskDefinition(opts *bind.WatchOpts, sink chan<- *ContractOBLSSetTotalVotingPowerPerTaskDefinition) (event.Subscription, error) {

	logs, sub, err := _ContractOBLS.contract.WatchLogs(opts, "SetTotalVotingPowerPerTaskDefinition")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOBLSSetTotalVotingPowerPerTaskDefinition)
				if err := _ContractOBLS.contract.UnpackLog(event, "SetTotalVotingPowerPerTaskDefinition", log); err != nil {
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

// ParseSetTotalVotingPowerPerTaskDefinition is a log parse operation binding the contract event 0x5a0cbf3f4ff24ea3dcc0680c9e2ab134f24e152cf4382343908507e4338d4b4c.
//
// Solidity: event SetTotalVotingPowerPerTaskDefinition(uint16 taskdefinitionId, uint256 numOfTotalOperators, uint256 minimumVotingPower)
func (_ContractOBLS *ContractOBLSFilterer) ParseSetTotalVotingPowerPerTaskDefinition(log types.Log) (*ContractOBLSSetTotalVotingPowerPerTaskDefinition, error) {
	event := new(ContractOBLSSetTotalVotingPowerPerTaskDefinition)
	if err := _ContractOBLS.contract.UnpackLog(event, "SetTotalVotingPowerPerTaskDefinition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOBLSSharesSyncerModifiedIterator is returned from FilterSharesSyncerModified and is used to iterate over the raw logs and unpacked data for SharesSyncerModified events raised by the ContractOBLS contract.
type ContractOBLSSharesSyncerModifiedIterator struct {
	Event *ContractOBLSSharesSyncerModified // Event containing the contract specifics and raw log

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
func (it *ContractOBLSSharesSyncerModifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOBLSSharesSyncerModified)
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
		it.Event = new(ContractOBLSSharesSyncerModified)
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
func (it *ContractOBLSSharesSyncerModifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOBLSSharesSyncerModifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOBLSSharesSyncerModified represents a SharesSyncerModified event raised by the ContractOBLS contract.
type ContractOBLSSharesSyncerModified struct {
	Syncer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSharesSyncerModified is a free log retrieval operation binding the contract event 0x797f7e8ad76b743fbff71cb58d982a1bd35e4df14315d7687b187c1c5c257e4d.
//
// Solidity: event SharesSyncerModified(address syncer)
func (_ContractOBLS *ContractOBLSFilterer) FilterSharesSyncerModified(opts *bind.FilterOpts) (*ContractOBLSSharesSyncerModifiedIterator, error) {

	logs, sub, err := _ContractOBLS.contract.FilterLogs(opts, "SharesSyncerModified")
	if err != nil {
		return nil, err
	}
	return &ContractOBLSSharesSyncerModifiedIterator{contract: _ContractOBLS.contract, event: "SharesSyncerModified", logs: logs, sub: sub}, nil
}

// WatchSharesSyncerModified is a free log subscription operation binding the contract event 0x797f7e8ad76b743fbff71cb58d982a1bd35e4df14315d7687b187c1c5c257e4d.
//
// Solidity: event SharesSyncerModified(address syncer)
func (_ContractOBLS *ContractOBLSFilterer) WatchSharesSyncerModified(opts *bind.WatchOpts, sink chan<- *ContractOBLSSharesSyncerModified) (event.Subscription, error) {

	logs, sub, err := _ContractOBLS.contract.WatchLogs(opts, "SharesSyncerModified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOBLSSharesSyncerModified)
				if err := _ContractOBLS.contract.UnpackLog(event, "SharesSyncerModified", log); err != nil {
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

// ParseSharesSyncerModified is a log parse operation binding the contract event 0x797f7e8ad76b743fbff71cb58d982a1bd35e4df14315d7687b187c1c5c257e4d.
//
// Solidity: event SharesSyncerModified(address syncer)
func (_ContractOBLS *ContractOBLSFilterer) ParseSharesSyncerModified(log types.Log) (*ContractOBLSSharesSyncerModified, error) {
	event := new(ContractOBLSSharesSyncerModified)
	if err := _ContractOBLS.contract.UnpackLog(event, "SharesSyncerModified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
