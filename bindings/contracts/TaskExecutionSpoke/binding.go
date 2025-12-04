// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractTaskExecutionSpoke

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

// Origin is an auto generated low-level Go binding around an user-defined struct.
type Origin struct {
	SrcEid uint32
	Sender [32]byte
	Nonce  uint64
}

// ContractTaskExecutionSpokeMetaData contains all meta data concerning the ContractTaskExecutionSpoke contract.
var ContractTaskExecutionSpokeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_endpoint\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_delegate\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDelegate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEndpointCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LzTokenUnavailable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"}],\"name\":\"NoPeer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"}],\"name\":\"NotEnoughNative\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"OnlyEndpoint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"}],\"name\":\"OnlyPeer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"FunctionExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumTaskExecutionSpoke.ActionType\",\"name\":\"action\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"}],\"name\":\"KeeperUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"peer\",\"type\":\"bytes32\"}],\"name\":\"PeerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"origin\",\"type\":\"tuple\"}],\"name\":\"allowInitializePath\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endpoint\",\"outputs\":[{\"internalType\":\"contractILayerZeroEndpointV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tgAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeFunction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ownerAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_hubEid\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"_initialKeepers\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_jobRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_triggerGasRegistryAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"isComposeMsgSender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isKeeper\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"jobRegistry\",\"outputs\":[{\"internalType\":\"contractIJobRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"_origin\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_guid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"lzReceive\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"nextNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oAppVersion\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"senderVersion\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"receiverVersion\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"}],\"name\":\"peers\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"peer\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegate\",\"type\":\"address\"}],\"name\":\"setDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_peer\",\"type\":\"bytes32\"}],\"name\":\"setPeer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"triggerGasRegistry\",\"outputs\":[{\"internalType\":\"contractITriggerGasRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// ContractTaskExecutionSpokeABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractTaskExecutionSpokeMetaData.ABI instead.
var ContractTaskExecutionSpokeABI = ContractTaskExecutionSpokeMetaData.ABI

// ContractTaskExecutionSpokeMethods is an auto generated interface around an Ethereum contract.
type ContractTaskExecutionSpokeMethods interface {
	ContractTaskExecutionSpokeCalls
	ContractTaskExecutionSpokeTransacts
	ContractTaskExecutionSpokeFilters
}

// ContractTaskExecutionSpokeCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractTaskExecutionSpokeCalls interface {
	UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error)

	AllowInitializePath(opts *bind.CallOpts, origin Origin) (bool, error)

	Endpoint(opts *bind.CallOpts) (common.Address, error)

	IsComposeMsgSender(opts *bind.CallOpts, arg0 Origin, arg1 []byte, _sender common.Address) (bool, error)

	IsKeeper(opts *bind.CallOpts, arg0 common.Address) (bool, error)

	JobRegistry(opts *bind.CallOpts) (common.Address, error)

	NextNonce(opts *bind.CallOpts, arg0 uint32, arg1 [32]byte) (uint64, error)

	OAppVersion(opts *bind.CallOpts) (struct {
		SenderVersion   uint64
		ReceiverVersion uint64
	}, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Peers(opts *bind.CallOpts, eid uint32) ([32]byte, error)

	ProxiableUUID(opts *bind.CallOpts) ([32]byte, error)

	TriggerGasRegistry(opts *bind.CallOpts) (common.Address, error)
}

// ContractTaskExecutionSpokeTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractTaskExecutionSpokeTransacts interface {
	ExecuteFunction(opts *bind.TransactOpts, jobId *big.Int, tgAmount *big.Int, target common.Address, data []byte) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts, _ownerAddress common.Address, _hubEid uint32, _initialKeepers []common.Address, _jobRegistryAddress common.Address, _triggerGasRegistryAddress common.Address) (*types.Transaction, error)

	LzReceive(opts *bind.TransactOpts, _origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error)

	RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetDelegate(opts *bind.TransactOpts, _delegate common.Address) (*types.Transaction, error)

	SetPeer(opts *bind.TransactOpts, _eid uint32, _peer [32]byte) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error)

	UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error)
}

// ContractTaskExecutionSpokeFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractTaskExecutionSpokeFilters interface {
	FilterFunctionExecuted(opts *bind.FilterOpts, keeper []common.Address, target []common.Address) (*ContractTaskExecutionSpokeFunctionExecutedIterator, error)
	WatchFunctionExecuted(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionSpokeFunctionExecuted, keeper []common.Address, target []common.Address) (event.Subscription, error)
	ParseFunctionExecuted(log types.Log) (*ContractTaskExecutionSpokeFunctionExecuted, error)

	FilterInitialized(opts *bind.FilterOpts) (*ContractTaskExecutionSpokeInitializedIterator, error)
	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionSpokeInitialized) (event.Subscription, error)
	ParseInitialized(log types.Log) (*ContractTaskExecutionSpokeInitialized, error)

	FilterKeeperUpdated(opts *bind.FilterOpts) (*ContractTaskExecutionSpokeKeeperUpdatedIterator, error)
	WatchKeeperUpdated(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionSpokeKeeperUpdated) (event.Subscription, error)
	ParseKeeperUpdated(log types.Log) (*ContractTaskExecutionSpokeKeeperUpdated, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractTaskExecutionSpokeOwnershipTransferredIterator, error)
	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionSpokeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error)
	ParseOwnershipTransferred(log types.Log) (*ContractTaskExecutionSpokeOwnershipTransferred, error)

	FilterPeerSet(opts *bind.FilterOpts) (*ContractTaskExecutionSpokePeerSetIterator, error)
	WatchPeerSet(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionSpokePeerSet) (event.Subscription, error)
	ParsePeerSet(log types.Log) (*ContractTaskExecutionSpokePeerSet, error)

	FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ContractTaskExecutionSpokeUpgradedIterator, error)
	WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionSpokeUpgraded, implementation []common.Address) (event.Subscription, error)
	ParseUpgraded(log types.Log) (*ContractTaskExecutionSpokeUpgraded, error)
}

// ContractTaskExecutionSpoke is an auto generated Go binding around an Ethereum contract.
type ContractTaskExecutionSpoke struct {
	ContractTaskExecutionSpokeCaller     // Read-only binding to the contract
	ContractTaskExecutionSpokeTransactor // Write-only binding to the contract
	ContractTaskExecutionSpokeFilterer   // Log filterer for contract events
}

// ContractTaskExecutionSpoke implements the ContractTaskExecutionSpokeMethods interface.
var _ ContractTaskExecutionSpokeMethods = (*ContractTaskExecutionSpoke)(nil)

// ContractTaskExecutionSpokeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractTaskExecutionSpokeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTaskExecutionSpokeCaller implements the ContractTaskExecutionSpokeCalls interface.
var _ ContractTaskExecutionSpokeCalls = (*ContractTaskExecutionSpokeCaller)(nil)

// ContractTaskExecutionSpokeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTaskExecutionSpokeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTaskExecutionSpokeTransactor implements the ContractTaskExecutionSpokeTransacts interface.
var _ ContractTaskExecutionSpokeTransacts = (*ContractTaskExecutionSpokeTransactor)(nil)

// ContractTaskExecutionSpokeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractTaskExecutionSpokeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTaskExecutionSpokeFilterer implements the ContractTaskExecutionSpokeFilters interface.
var _ ContractTaskExecutionSpokeFilters = (*ContractTaskExecutionSpokeFilterer)(nil)

// ContractTaskExecutionSpokeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractTaskExecutionSpokeSession struct {
	Contract     *ContractTaskExecutionSpoke // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ContractTaskExecutionSpokeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractTaskExecutionSpokeCallerSession struct {
	Contract *ContractTaskExecutionSpokeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// ContractTaskExecutionSpokeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTaskExecutionSpokeTransactorSession struct {
	Contract     *ContractTaskExecutionSpokeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// ContractTaskExecutionSpokeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractTaskExecutionSpokeRaw struct {
	Contract *ContractTaskExecutionSpoke // Generic contract binding to access the raw methods on
}

// ContractTaskExecutionSpokeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractTaskExecutionSpokeCallerRaw struct {
	Contract *ContractTaskExecutionSpokeCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTaskExecutionSpokeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTaskExecutionSpokeTransactorRaw struct {
	Contract *ContractTaskExecutionSpokeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractTaskExecutionSpoke creates a new instance of ContractTaskExecutionSpoke, bound to a specific deployed contract.
func NewContractTaskExecutionSpoke(address common.Address, backend bind.ContractBackend) (*ContractTaskExecutionSpoke, error) {
	contract, err := bindContractTaskExecutionSpoke(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractTaskExecutionSpoke{ContractTaskExecutionSpokeCaller: ContractTaskExecutionSpokeCaller{contract: contract}, ContractTaskExecutionSpokeTransactor: ContractTaskExecutionSpokeTransactor{contract: contract}, ContractTaskExecutionSpokeFilterer: ContractTaskExecutionSpokeFilterer{contract: contract}}, nil
}

// NewContractTaskExecutionSpokeCaller creates a new read-only instance of ContractTaskExecutionSpoke, bound to a specific deployed contract.
func NewContractTaskExecutionSpokeCaller(address common.Address, caller bind.ContractCaller) (*ContractTaskExecutionSpokeCaller, error) {
	contract, err := bindContractTaskExecutionSpoke(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTaskExecutionSpokeCaller{contract: contract}, nil
}

// NewContractTaskExecutionSpokeTransactor creates a new write-only instance of ContractTaskExecutionSpoke, bound to a specific deployed contract.
func NewContractTaskExecutionSpokeTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTaskExecutionSpokeTransactor, error) {
	contract, err := bindContractTaskExecutionSpoke(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTaskExecutionSpokeTransactor{contract: contract}, nil
}

// NewContractTaskExecutionSpokeFilterer creates a new log filterer instance of ContractTaskExecutionSpoke, bound to a specific deployed contract.
func NewContractTaskExecutionSpokeFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractTaskExecutionSpokeFilterer, error) {
	contract, err := bindContractTaskExecutionSpoke(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractTaskExecutionSpokeFilterer{contract: contract}, nil
}

// bindContractTaskExecutionSpoke binds a generic wrapper to an already deployed contract.
func bindContractTaskExecutionSpoke(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractTaskExecutionSpokeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractTaskExecutionSpoke.Contract.ContractTaskExecutionSpokeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.Contract.ContractTaskExecutionSpokeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.Contract.ContractTaskExecutionSpokeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractTaskExecutionSpoke.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ContractTaskExecutionSpoke.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ContractTaskExecutionSpoke.Contract.UPGRADEINTERFACEVERSION(&_ContractTaskExecutionSpoke.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ContractTaskExecutionSpoke.Contract.UPGRADEINTERFACEVERSION(&_ContractTaskExecutionSpoke.CallOpts)
}

// AllowInitializePath is a free data retrieval call binding the contract method 0xff7bd03d.
//
// Solidity: function allowInitializePath((uint32,bytes32,uint64) origin) view returns(bool)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCaller) AllowInitializePath(opts *bind.CallOpts, origin Origin) (bool, error) {
	var out []interface{}
	err := _ContractTaskExecutionSpoke.contract.Call(opts, &out, "allowInitializePath", origin)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowInitializePath is a free data retrieval call binding the contract method 0xff7bd03d.
//
// Solidity: function allowInitializePath((uint32,bytes32,uint64) origin) view returns(bool)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeSession) AllowInitializePath(origin Origin) (bool, error) {
	return _ContractTaskExecutionSpoke.Contract.AllowInitializePath(&_ContractTaskExecutionSpoke.CallOpts, origin)
}

// AllowInitializePath is a free data retrieval call binding the contract method 0xff7bd03d.
//
// Solidity: function allowInitializePath((uint32,bytes32,uint64) origin) view returns(bool)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCallerSession) AllowInitializePath(origin Origin) (bool, error) {
	return _ContractTaskExecutionSpoke.Contract.AllowInitializePath(&_ContractTaskExecutionSpoke.CallOpts, origin)
}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCaller) Endpoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTaskExecutionSpoke.contract.Call(opts, &out, "endpoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeSession) Endpoint() (common.Address, error) {
	return _ContractTaskExecutionSpoke.Contract.Endpoint(&_ContractTaskExecutionSpoke.CallOpts)
}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCallerSession) Endpoint() (common.Address, error) {
	return _ContractTaskExecutionSpoke.Contract.Endpoint(&_ContractTaskExecutionSpoke.CallOpts)
}

// IsComposeMsgSender is a free data retrieval call binding the contract method 0x82413eac.
//
// Solidity: function isComposeMsgSender((uint32,bytes32,uint64) , bytes , address _sender) view returns(bool)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCaller) IsComposeMsgSender(opts *bind.CallOpts, arg0 Origin, arg1 []byte, _sender common.Address) (bool, error) {
	var out []interface{}
	err := _ContractTaskExecutionSpoke.contract.Call(opts, &out, "isComposeMsgSender", arg0, arg1, _sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsComposeMsgSender is a free data retrieval call binding the contract method 0x82413eac.
//
// Solidity: function isComposeMsgSender((uint32,bytes32,uint64) , bytes , address _sender) view returns(bool)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeSession) IsComposeMsgSender(arg0 Origin, arg1 []byte, _sender common.Address) (bool, error) {
	return _ContractTaskExecutionSpoke.Contract.IsComposeMsgSender(&_ContractTaskExecutionSpoke.CallOpts, arg0, arg1, _sender)
}

// IsComposeMsgSender is a free data retrieval call binding the contract method 0x82413eac.
//
// Solidity: function isComposeMsgSender((uint32,bytes32,uint64) , bytes , address _sender) view returns(bool)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCallerSession) IsComposeMsgSender(arg0 Origin, arg1 []byte, _sender common.Address) (bool, error) {
	return _ContractTaskExecutionSpoke.Contract.IsComposeMsgSender(&_ContractTaskExecutionSpoke.CallOpts, arg0, arg1, _sender)
}

// IsKeeper is a free data retrieval call binding the contract method 0x6ba42aaa.
//
// Solidity: function isKeeper(address ) view returns(bool)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCaller) IsKeeper(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ContractTaskExecutionSpoke.contract.Call(opts, &out, "isKeeper", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsKeeper is a free data retrieval call binding the contract method 0x6ba42aaa.
//
// Solidity: function isKeeper(address ) view returns(bool)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeSession) IsKeeper(arg0 common.Address) (bool, error) {
	return _ContractTaskExecutionSpoke.Contract.IsKeeper(&_ContractTaskExecutionSpoke.CallOpts, arg0)
}

// IsKeeper is a free data retrieval call binding the contract method 0x6ba42aaa.
//
// Solidity: function isKeeper(address ) view returns(bool)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCallerSession) IsKeeper(arg0 common.Address) (bool, error) {
	return _ContractTaskExecutionSpoke.Contract.IsKeeper(&_ContractTaskExecutionSpoke.CallOpts, arg0)
}

// JobRegistry is a free data retrieval call binding the contract method 0x23682c47.
//
// Solidity: function jobRegistry() view returns(address)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCaller) JobRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTaskExecutionSpoke.contract.Call(opts, &out, "jobRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// JobRegistry is a free data retrieval call binding the contract method 0x23682c47.
//
// Solidity: function jobRegistry() view returns(address)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeSession) JobRegistry() (common.Address, error) {
	return _ContractTaskExecutionSpoke.Contract.JobRegistry(&_ContractTaskExecutionSpoke.CallOpts)
}

// JobRegistry is a free data retrieval call binding the contract method 0x23682c47.
//
// Solidity: function jobRegistry() view returns(address)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCallerSession) JobRegistry() (common.Address, error) {
	return _ContractTaskExecutionSpoke.Contract.JobRegistry(&_ContractTaskExecutionSpoke.CallOpts)
}

// NextNonce is a free data retrieval call binding the contract method 0x7d25a05e.
//
// Solidity: function nextNonce(uint32 , bytes32 ) view returns(uint64 nonce)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCaller) NextNonce(opts *bind.CallOpts, arg0 uint32, arg1 [32]byte) (uint64, error) {
	var out []interface{}
	err := _ContractTaskExecutionSpoke.contract.Call(opts, &out, "nextNonce", arg0, arg1)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NextNonce is a free data retrieval call binding the contract method 0x7d25a05e.
//
// Solidity: function nextNonce(uint32 , bytes32 ) view returns(uint64 nonce)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeSession) NextNonce(arg0 uint32, arg1 [32]byte) (uint64, error) {
	return _ContractTaskExecutionSpoke.Contract.NextNonce(&_ContractTaskExecutionSpoke.CallOpts, arg0, arg1)
}

// NextNonce is a free data retrieval call binding the contract method 0x7d25a05e.
//
// Solidity: function nextNonce(uint32 , bytes32 ) view returns(uint64 nonce)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCallerSession) NextNonce(arg0 uint32, arg1 [32]byte) (uint64, error) {
	return _ContractTaskExecutionSpoke.Contract.NextNonce(&_ContractTaskExecutionSpoke.CallOpts, arg0, arg1)
}

// OAppVersion is a free data retrieval call binding the contract method 0x17442b70.
//
// Solidity: function oAppVersion() pure returns(uint64 senderVersion, uint64 receiverVersion)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCaller) OAppVersion(opts *bind.CallOpts) (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	var out []interface{}
	err := _ContractTaskExecutionSpoke.contract.Call(opts, &out, "oAppVersion")

	outstruct := new(struct {
		SenderVersion   uint64
		ReceiverVersion uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SenderVersion = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.ReceiverVersion = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// OAppVersion is a free data retrieval call binding the contract method 0x17442b70.
//
// Solidity: function oAppVersion() pure returns(uint64 senderVersion, uint64 receiverVersion)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeSession) OAppVersion() (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	return _ContractTaskExecutionSpoke.Contract.OAppVersion(&_ContractTaskExecutionSpoke.CallOpts)
}

// OAppVersion is a free data retrieval call binding the contract method 0x17442b70.
//
// Solidity: function oAppVersion() pure returns(uint64 senderVersion, uint64 receiverVersion)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCallerSession) OAppVersion() (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	return _ContractTaskExecutionSpoke.Contract.OAppVersion(&_ContractTaskExecutionSpoke.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTaskExecutionSpoke.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeSession) Owner() (common.Address, error) {
	return _ContractTaskExecutionSpoke.Contract.Owner(&_ContractTaskExecutionSpoke.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCallerSession) Owner() (common.Address, error) {
	return _ContractTaskExecutionSpoke.Contract.Owner(&_ContractTaskExecutionSpoke.CallOpts)
}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 eid) view returns(bytes32 peer)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCaller) Peers(opts *bind.CallOpts, eid uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractTaskExecutionSpoke.contract.Call(opts, &out, "peers", eid)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 eid) view returns(bytes32 peer)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeSession) Peers(eid uint32) ([32]byte, error) {
	return _ContractTaskExecutionSpoke.Contract.Peers(&_ContractTaskExecutionSpoke.CallOpts, eid)
}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 eid) view returns(bytes32 peer)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCallerSession) Peers(eid uint32) ([32]byte, error) {
	return _ContractTaskExecutionSpoke.Contract.Peers(&_ContractTaskExecutionSpoke.CallOpts, eid)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractTaskExecutionSpoke.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeSession) ProxiableUUID() ([32]byte, error) {
	return _ContractTaskExecutionSpoke.Contract.ProxiableUUID(&_ContractTaskExecutionSpoke.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCallerSession) ProxiableUUID() ([32]byte, error) {
	return _ContractTaskExecutionSpoke.Contract.ProxiableUUID(&_ContractTaskExecutionSpoke.CallOpts)
}

// TriggerGasRegistry is a free data retrieval call binding the contract method 0xd908b2ef.
//
// Solidity: function triggerGasRegistry() view returns(address)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCaller) TriggerGasRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTaskExecutionSpoke.contract.Call(opts, &out, "triggerGasRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TriggerGasRegistry is a free data retrieval call binding the contract method 0xd908b2ef.
//
// Solidity: function triggerGasRegistry() view returns(address)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeSession) TriggerGasRegistry() (common.Address, error) {
	return _ContractTaskExecutionSpoke.Contract.TriggerGasRegistry(&_ContractTaskExecutionSpoke.CallOpts)
}

// TriggerGasRegistry is a free data retrieval call binding the contract method 0xd908b2ef.
//
// Solidity: function triggerGasRegistry() view returns(address)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCallerSession) TriggerGasRegistry() (common.Address, error) {
	return _ContractTaskExecutionSpoke.Contract.TriggerGasRegistry(&_ContractTaskExecutionSpoke.CallOpts)
}

// ExecuteFunction is a paid mutator transaction binding the contract method 0xfa9b1a80.
//
// Solidity: function executeFunction(uint256 jobId, uint256 tgAmount, address target, bytes data) payable returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeTransactor) ExecuteFunction(opts *bind.TransactOpts, jobId *big.Int, tgAmount *big.Int, target common.Address, data []byte) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.contract.Transact(opts, "executeFunction", jobId, tgAmount, target, data)
}

// ExecuteFunction is a paid mutator transaction binding the contract method 0xfa9b1a80.
//
// Solidity: function executeFunction(uint256 jobId, uint256 tgAmount, address target, bytes data) payable returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeSession) ExecuteFunction(jobId *big.Int, tgAmount *big.Int, target common.Address, data []byte) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.Contract.ExecuteFunction(&_ContractTaskExecutionSpoke.TransactOpts, jobId, tgAmount, target, data)
}

// ExecuteFunction is a paid mutator transaction binding the contract method 0xfa9b1a80.
//
// Solidity: function executeFunction(uint256 jobId, uint256 tgAmount, address target, bytes data) payable returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeTransactorSession) ExecuteFunction(jobId *big.Int, tgAmount *big.Int, target common.Address, data []byte) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.Contract.ExecuteFunction(&_ContractTaskExecutionSpoke.TransactOpts, jobId, tgAmount, target, data)
}

// Initialize is a paid mutator transaction binding the contract method 0x149e0da4.
//
// Solidity: function initialize(address _ownerAddress, uint32 _hubEid, address[] _initialKeepers, address _jobRegistryAddress, address _triggerGasRegistryAddress) returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeTransactor) Initialize(opts *bind.TransactOpts, _ownerAddress common.Address, _hubEid uint32, _initialKeepers []common.Address, _jobRegistryAddress common.Address, _triggerGasRegistryAddress common.Address) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.contract.Transact(opts, "initialize", _ownerAddress, _hubEid, _initialKeepers, _jobRegistryAddress, _triggerGasRegistryAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x149e0da4.
//
// Solidity: function initialize(address _ownerAddress, uint32 _hubEid, address[] _initialKeepers, address _jobRegistryAddress, address _triggerGasRegistryAddress) returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeSession) Initialize(_ownerAddress common.Address, _hubEid uint32, _initialKeepers []common.Address, _jobRegistryAddress common.Address, _triggerGasRegistryAddress common.Address) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.Contract.Initialize(&_ContractTaskExecutionSpoke.TransactOpts, _ownerAddress, _hubEid, _initialKeepers, _jobRegistryAddress, _triggerGasRegistryAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x149e0da4.
//
// Solidity: function initialize(address _ownerAddress, uint32 _hubEid, address[] _initialKeepers, address _jobRegistryAddress, address _triggerGasRegistryAddress) returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeTransactorSession) Initialize(_ownerAddress common.Address, _hubEid uint32, _initialKeepers []common.Address, _jobRegistryAddress common.Address, _triggerGasRegistryAddress common.Address) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.Contract.Initialize(&_ContractTaskExecutionSpoke.TransactOpts, _ownerAddress, _hubEid, _initialKeepers, _jobRegistryAddress, _triggerGasRegistryAddress)
}

// LzReceive is a paid mutator transaction binding the contract method 0x13137d65.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeTransactor) LzReceive(opts *bind.TransactOpts, _origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.contract.Transact(opts, "lzReceive", _origin, _guid, _message, _executor, _extraData)
}

// LzReceive is a paid mutator transaction binding the contract method 0x13137d65.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeSession) LzReceive(_origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.Contract.LzReceive(&_ContractTaskExecutionSpoke.TransactOpts, _origin, _guid, _message, _executor, _extraData)
}

// LzReceive is a paid mutator transaction binding the contract method 0x13137d65.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeTransactorSession) LzReceive(_origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.Contract.LzReceive(&_ContractTaskExecutionSpoke.TransactOpts, _origin, _guid, _message, _executor, _extraData)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.Contract.RenounceOwnership(&_ContractTaskExecutionSpoke.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.Contract.RenounceOwnership(&_ContractTaskExecutionSpoke.TransactOpts)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeTransactor) SetDelegate(opts *bind.TransactOpts, _delegate common.Address) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.contract.Transact(opts, "setDelegate", _delegate)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeSession) SetDelegate(_delegate common.Address) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.Contract.SetDelegate(&_ContractTaskExecutionSpoke.TransactOpts, _delegate)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeTransactorSession) SetDelegate(_delegate common.Address) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.Contract.SetDelegate(&_ContractTaskExecutionSpoke.TransactOpts, _delegate)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeTransactor) SetPeer(opts *bind.TransactOpts, _eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.contract.Transact(opts, "setPeer", _eid, _peer)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeSession) SetPeer(_eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.Contract.SetPeer(&_ContractTaskExecutionSpoke.TransactOpts, _eid, _peer)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeTransactorSession) SetPeer(_eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.Contract.SetPeer(&_ContractTaskExecutionSpoke.TransactOpts, _eid, _peer)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.Contract.TransferOwnership(&_ContractTaskExecutionSpoke.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.Contract.TransferOwnership(&_ContractTaskExecutionSpoke.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.Contract.UpgradeToAndCall(&_ContractTaskExecutionSpoke.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.Contract.UpgradeToAndCall(&_ContractTaskExecutionSpoke.TransactOpts, newImplementation, data)
}

// ContractTaskExecutionSpokeFunctionExecutedIterator is returned from FilterFunctionExecuted and is used to iterate over the raw logs and unpacked data for FunctionExecuted events raised by the ContractTaskExecutionSpoke contract.
type ContractTaskExecutionSpokeFunctionExecutedIterator struct {
	Event *ContractTaskExecutionSpokeFunctionExecuted // Event containing the contract specifics and raw log

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
func (it *ContractTaskExecutionSpokeFunctionExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTaskExecutionSpokeFunctionExecuted)
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
		it.Event = new(ContractTaskExecutionSpokeFunctionExecuted)
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
func (it *ContractTaskExecutionSpokeFunctionExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTaskExecutionSpokeFunctionExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTaskExecutionSpokeFunctionExecuted represents a FunctionExecuted event raised by the ContractTaskExecutionSpoke contract.
type ContractTaskExecutionSpokeFunctionExecuted struct {
	Keeper common.Address
	Target common.Address
	Data   []byte
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFunctionExecuted is a free log retrieval operation binding the contract event 0xf4448cdaf10358453fa19b2e0363f44780277619bd9bb676eef79d7488a7d274.
//
// Solidity: event FunctionExecuted(address indexed keeper, address indexed target, bytes data, uint256 value)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeFilterer) FilterFunctionExecuted(opts *bind.FilterOpts, keeper []common.Address, target []common.Address) (*ContractTaskExecutionSpokeFunctionExecutedIterator, error) {

	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _ContractTaskExecutionSpoke.contract.FilterLogs(opts, "FunctionExecuted", keeperRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &ContractTaskExecutionSpokeFunctionExecutedIterator{contract: _ContractTaskExecutionSpoke.contract, event: "FunctionExecuted", logs: logs, sub: sub}, nil
}

// WatchFunctionExecuted is a free log subscription operation binding the contract event 0xf4448cdaf10358453fa19b2e0363f44780277619bd9bb676eef79d7488a7d274.
//
// Solidity: event FunctionExecuted(address indexed keeper, address indexed target, bytes data, uint256 value)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeFilterer) WatchFunctionExecuted(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionSpokeFunctionExecuted, keeper []common.Address, target []common.Address) (event.Subscription, error) {

	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _ContractTaskExecutionSpoke.contract.WatchLogs(opts, "FunctionExecuted", keeperRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTaskExecutionSpokeFunctionExecuted)
				if err := _ContractTaskExecutionSpoke.contract.UnpackLog(event, "FunctionExecuted", log); err != nil {
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

// ParseFunctionExecuted is a log parse operation binding the contract event 0xf4448cdaf10358453fa19b2e0363f44780277619bd9bb676eef79d7488a7d274.
//
// Solidity: event FunctionExecuted(address indexed keeper, address indexed target, bytes data, uint256 value)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeFilterer) ParseFunctionExecuted(log types.Log) (*ContractTaskExecutionSpokeFunctionExecuted, error) {
	event := new(ContractTaskExecutionSpokeFunctionExecuted)
	if err := _ContractTaskExecutionSpoke.contract.UnpackLog(event, "FunctionExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTaskExecutionSpokeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractTaskExecutionSpoke contract.
type ContractTaskExecutionSpokeInitializedIterator struct {
	Event *ContractTaskExecutionSpokeInitialized // Event containing the contract specifics and raw log

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
func (it *ContractTaskExecutionSpokeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTaskExecutionSpokeInitialized)
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
		it.Event = new(ContractTaskExecutionSpokeInitialized)
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
func (it *ContractTaskExecutionSpokeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTaskExecutionSpokeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTaskExecutionSpokeInitialized represents a Initialized event raised by the ContractTaskExecutionSpoke contract.
type ContractTaskExecutionSpokeInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractTaskExecutionSpokeInitializedIterator, error) {

	logs, sub, err := _ContractTaskExecutionSpoke.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractTaskExecutionSpokeInitializedIterator{contract: _ContractTaskExecutionSpoke.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionSpokeInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractTaskExecutionSpoke.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTaskExecutionSpokeInitialized)
				if err := _ContractTaskExecutionSpoke.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeFilterer) ParseInitialized(log types.Log) (*ContractTaskExecutionSpokeInitialized, error) {
	event := new(ContractTaskExecutionSpokeInitialized)
	if err := _ContractTaskExecutionSpoke.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTaskExecutionSpokeKeeperUpdatedIterator is returned from FilterKeeperUpdated and is used to iterate over the raw logs and unpacked data for KeeperUpdated events raised by the ContractTaskExecutionSpoke contract.
type ContractTaskExecutionSpokeKeeperUpdatedIterator struct {
	Event *ContractTaskExecutionSpokeKeeperUpdated // Event containing the contract specifics and raw log

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
func (it *ContractTaskExecutionSpokeKeeperUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTaskExecutionSpokeKeeperUpdated)
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
		it.Event = new(ContractTaskExecutionSpokeKeeperUpdated)
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
func (it *ContractTaskExecutionSpokeKeeperUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTaskExecutionSpokeKeeperUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTaskExecutionSpokeKeeperUpdated represents a KeeperUpdated event raised by the ContractTaskExecutionSpoke contract.
type ContractTaskExecutionSpokeKeeperUpdated struct {
	Action uint8
	Keeper common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterKeeperUpdated is a free log retrieval operation binding the contract event 0xa16bcb5f8175f03e8484e9d840834e357a7ca38d877946862644bf1e078ee1b2.
//
// Solidity: event KeeperUpdated(uint8 action, address keeper)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeFilterer) FilterKeeperUpdated(opts *bind.FilterOpts) (*ContractTaskExecutionSpokeKeeperUpdatedIterator, error) {

	logs, sub, err := _ContractTaskExecutionSpoke.contract.FilterLogs(opts, "KeeperUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractTaskExecutionSpokeKeeperUpdatedIterator{contract: _ContractTaskExecutionSpoke.contract, event: "KeeperUpdated", logs: logs, sub: sub}, nil
}

// WatchKeeperUpdated is a free log subscription operation binding the contract event 0xa16bcb5f8175f03e8484e9d840834e357a7ca38d877946862644bf1e078ee1b2.
//
// Solidity: event KeeperUpdated(uint8 action, address keeper)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeFilterer) WatchKeeperUpdated(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionSpokeKeeperUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractTaskExecutionSpoke.contract.WatchLogs(opts, "KeeperUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTaskExecutionSpokeKeeperUpdated)
				if err := _ContractTaskExecutionSpoke.contract.UnpackLog(event, "KeeperUpdated", log); err != nil {
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

// ParseKeeperUpdated is a log parse operation binding the contract event 0xa16bcb5f8175f03e8484e9d840834e357a7ca38d877946862644bf1e078ee1b2.
//
// Solidity: event KeeperUpdated(uint8 action, address keeper)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeFilterer) ParseKeeperUpdated(log types.Log) (*ContractTaskExecutionSpokeKeeperUpdated, error) {
	event := new(ContractTaskExecutionSpokeKeeperUpdated)
	if err := _ContractTaskExecutionSpoke.contract.UnpackLog(event, "KeeperUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTaskExecutionSpokeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractTaskExecutionSpoke contract.
type ContractTaskExecutionSpokeOwnershipTransferredIterator struct {
	Event *ContractTaskExecutionSpokeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractTaskExecutionSpokeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTaskExecutionSpokeOwnershipTransferred)
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
		it.Event = new(ContractTaskExecutionSpokeOwnershipTransferred)
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
func (it *ContractTaskExecutionSpokeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTaskExecutionSpokeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTaskExecutionSpokeOwnershipTransferred represents a OwnershipTransferred event raised by the ContractTaskExecutionSpoke contract.
type ContractTaskExecutionSpokeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractTaskExecutionSpokeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractTaskExecutionSpoke.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractTaskExecutionSpokeOwnershipTransferredIterator{contract: _ContractTaskExecutionSpoke.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionSpokeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractTaskExecutionSpoke.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTaskExecutionSpokeOwnershipTransferred)
				if err := _ContractTaskExecutionSpoke.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeFilterer) ParseOwnershipTransferred(log types.Log) (*ContractTaskExecutionSpokeOwnershipTransferred, error) {
	event := new(ContractTaskExecutionSpokeOwnershipTransferred)
	if err := _ContractTaskExecutionSpoke.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTaskExecutionSpokePeerSetIterator is returned from FilterPeerSet and is used to iterate over the raw logs and unpacked data for PeerSet events raised by the ContractTaskExecutionSpoke contract.
type ContractTaskExecutionSpokePeerSetIterator struct {
	Event *ContractTaskExecutionSpokePeerSet // Event containing the contract specifics and raw log

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
func (it *ContractTaskExecutionSpokePeerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTaskExecutionSpokePeerSet)
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
		it.Event = new(ContractTaskExecutionSpokePeerSet)
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
func (it *ContractTaskExecutionSpokePeerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTaskExecutionSpokePeerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTaskExecutionSpokePeerSet represents a PeerSet event raised by the ContractTaskExecutionSpoke contract.
type ContractTaskExecutionSpokePeerSet struct {
	Eid  uint32
	Peer [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPeerSet is a free log retrieval operation binding the contract event 0x238399d427b947898edb290f5ff0f9109849b1c3ba196a42e35f00c50a54b98b.
//
// Solidity: event PeerSet(uint32 eid, bytes32 peer)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeFilterer) FilterPeerSet(opts *bind.FilterOpts) (*ContractTaskExecutionSpokePeerSetIterator, error) {

	logs, sub, err := _ContractTaskExecutionSpoke.contract.FilterLogs(opts, "PeerSet")
	if err != nil {
		return nil, err
	}
	return &ContractTaskExecutionSpokePeerSetIterator{contract: _ContractTaskExecutionSpoke.contract, event: "PeerSet", logs: logs, sub: sub}, nil
}

// WatchPeerSet is a free log subscription operation binding the contract event 0x238399d427b947898edb290f5ff0f9109849b1c3ba196a42e35f00c50a54b98b.
//
// Solidity: event PeerSet(uint32 eid, bytes32 peer)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeFilterer) WatchPeerSet(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionSpokePeerSet) (event.Subscription, error) {

	logs, sub, err := _ContractTaskExecutionSpoke.contract.WatchLogs(opts, "PeerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTaskExecutionSpokePeerSet)
				if err := _ContractTaskExecutionSpoke.contract.UnpackLog(event, "PeerSet", log); err != nil {
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

// ParsePeerSet is a log parse operation binding the contract event 0x238399d427b947898edb290f5ff0f9109849b1c3ba196a42e35f00c50a54b98b.
//
// Solidity: event PeerSet(uint32 eid, bytes32 peer)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeFilterer) ParsePeerSet(log types.Log) (*ContractTaskExecutionSpokePeerSet, error) {
	event := new(ContractTaskExecutionSpokePeerSet)
	if err := _ContractTaskExecutionSpoke.contract.UnpackLog(event, "PeerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTaskExecutionSpokeUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ContractTaskExecutionSpoke contract.
type ContractTaskExecutionSpokeUpgradedIterator struct {
	Event *ContractTaskExecutionSpokeUpgraded // Event containing the contract specifics and raw log

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
func (it *ContractTaskExecutionSpokeUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTaskExecutionSpokeUpgraded)
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
		it.Event = new(ContractTaskExecutionSpokeUpgraded)
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
func (it *ContractTaskExecutionSpokeUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTaskExecutionSpokeUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTaskExecutionSpokeUpgraded represents a Upgraded event raised by the ContractTaskExecutionSpoke contract.
type ContractTaskExecutionSpokeUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ContractTaskExecutionSpokeUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ContractTaskExecutionSpoke.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ContractTaskExecutionSpokeUpgradedIterator{contract: _ContractTaskExecutionSpoke.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionSpokeUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ContractTaskExecutionSpoke.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTaskExecutionSpokeUpgraded)
				if err := _ContractTaskExecutionSpoke.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeFilterer) ParseUpgraded(log types.Log) (*ContractTaskExecutionSpokeUpgraded, error) {
	event := new(ContractTaskExecutionSpokeUpgraded)
	if err := _ContractTaskExecutionSpoke.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
