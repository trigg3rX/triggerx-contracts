// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractTriggerXSafeFactory

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

// ContractTriggerXSafeFactoryMetaData contains all meta data concerning the ContractTriggerXSafeFactory contract.
var ContractTriggerXSafeFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_taskExecutionHub\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_safeProxyFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_safeSingleton\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"safeWallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"saltNonce\",\"type\":\"uint256\"}],\"name\":\"SafeWalletCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"createSafeWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"safeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getSafeWallets\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserSaltNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"latestSafeWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"predictSafeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"predictedAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"safeProxyFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"safeSingleton\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taskExecutionHub\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ContractTriggerXSafeFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractTriggerXSafeFactoryMetaData.ABI instead.
var ContractTriggerXSafeFactoryABI = ContractTriggerXSafeFactoryMetaData.ABI

// ContractTriggerXSafeFactoryMethods is an auto generated interface around an Ethereum contract.
type ContractTriggerXSafeFactoryMethods interface {
	ContractTriggerXSafeFactoryCalls
	ContractTriggerXSafeFactoryTransacts
	ContractTriggerXSafeFactoryFilters
}

// ContractTriggerXSafeFactoryCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractTriggerXSafeFactoryCalls interface {
	GetSafeWallets(opts *bind.CallOpts, user common.Address) ([]common.Address, error)

	GetUserSaltNonce(opts *bind.CallOpts, user common.Address) (*big.Int, error)

	LatestSafeWallet(opts *bind.CallOpts, user common.Address) (common.Address, error)

	PredictSafeAddress(opts *bind.CallOpts, user common.Address) (common.Address, error)

	SafeProxyFactory(opts *bind.CallOpts) (common.Address, error)

	SafeSingleton(opts *bind.CallOpts) (common.Address, error)

	TaskExecutionHub(opts *bind.CallOpts) (common.Address, error)
}

// ContractTriggerXSafeFactoryTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractTriggerXSafeFactoryTransacts interface {
	CreateSafeWallet(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error)
}

// ContractTriggerXSafeFactoryFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractTriggerXSafeFactoryFilters interface {
	FilterSafeWalletCreated(opts *bind.FilterOpts, user []common.Address, safeWallet []common.Address) (*ContractTriggerXSafeFactorySafeWalletCreatedIterator, error)
	WatchSafeWalletCreated(opts *bind.WatchOpts, sink chan<- *ContractTriggerXSafeFactorySafeWalletCreated, user []common.Address, safeWallet []common.Address) (event.Subscription, error)
	ParseSafeWalletCreated(log types.Log) (*ContractTriggerXSafeFactorySafeWalletCreated, error)
}

// ContractTriggerXSafeFactory is an auto generated Go binding around an Ethereum contract.
type ContractTriggerXSafeFactory struct {
	ContractTriggerXSafeFactoryCaller     // Read-only binding to the contract
	ContractTriggerXSafeFactoryTransactor // Write-only binding to the contract
	ContractTriggerXSafeFactoryFilterer   // Log filterer for contract events
}

// ContractTriggerXSafeFactory implements the ContractTriggerXSafeFactoryMethods interface.
var _ ContractTriggerXSafeFactoryMethods = (*ContractTriggerXSafeFactory)(nil)

// ContractTriggerXSafeFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractTriggerXSafeFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTriggerXSafeFactoryCaller implements the ContractTriggerXSafeFactoryCalls interface.
var _ ContractTriggerXSafeFactoryCalls = (*ContractTriggerXSafeFactoryCaller)(nil)

// ContractTriggerXSafeFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTriggerXSafeFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTriggerXSafeFactoryTransactor implements the ContractTriggerXSafeFactoryTransacts interface.
var _ ContractTriggerXSafeFactoryTransacts = (*ContractTriggerXSafeFactoryTransactor)(nil)

// ContractTriggerXSafeFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractTriggerXSafeFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTriggerXSafeFactoryFilterer implements the ContractTriggerXSafeFactoryFilters interface.
var _ ContractTriggerXSafeFactoryFilters = (*ContractTriggerXSafeFactoryFilterer)(nil)

// ContractTriggerXSafeFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractTriggerXSafeFactorySession struct {
	Contract     *ContractTriggerXSafeFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ContractTriggerXSafeFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractTriggerXSafeFactoryCallerSession struct {
	Contract *ContractTriggerXSafeFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// ContractTriggerXSafeFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTriggerXSafeFactoryTransactorSession struct {
	Contract     *ContractTriggerXSafeFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// ContractTriggerXSafeFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractTriggerXSafeFactoryRaw struct {
	Contract *ContractTriggerXSafeFactory // Generic contract binding to access the raw methods on
}

// ContractTriggerXSafeFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractTriggerXSafeFactoryCallerRaw struct {
	Contract *ContractTriggerXSafeFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTriggerXSafeFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTriggerXSafeFactoryTransactorRaw struct {
	Contract *ContractTriggerXSafeFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractTriggerXSafeFactory creates a new instance of ContractTriggerXSafeFactory, bound to a specific deployed contract.
func NewContractTriggerXSafeFactory(address common.Address, backend bind.ContractBackend) (*ContractTriggerXSafeFactory, error) {
	contract, err := bindContractTriggerXSafeFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXSafeFactory{ContractTriggerXSafeFactoryCaller: ContractTriggerXSafeFactoryCaller{contract: contract}, ContractTriggerXSafeFactoryTransactor: ContractTriggerXSafeFactoryTransactor{contract: contract}, ContractTriggerXSafeFactoryFilterer: ContractTriggerXSafeFactoryFilterer{contract: contract}}, nil
}

// NewContractTriggerXSafeFactoryCaller creates a new read-only instance of ContractTriggerXSafeFactory, bound to a specific deployed contract.
func NewContractTriggerXSafeFactoryCaller(address common.Address, caller bind.ContractCaller) (*ContractTriggerXSafeFactoryCaller, error) {
	contract, err := bindContractTriggerXSafeFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXSafeFactoryCaller{contract: contract}, nil
}

// NewContractTriggerXSafeFactoryTransactor creates a new write-only instance of ContractTriggerXSafeFactory, bound to a specific deployed contract.
func NewContractTriggerXSafeFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTriggerXSafeFactoryTransactor, error) {
	contract, err := bindContractTriggerXSafeFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXSafeFactoryTransactor{contract: contract}, nil
}

// NewContractTriggerXSafeFactoryFilterer creates a new log filterer instance of ContractTriggerXSafeFactory, bound to a specific deployed contract.
func NewContractTriggerXSafeFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractTriggerXSafeFactoryFilterer, error) {
	contract, err := bindContractTriggerXSafeFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXSafeFactoryFilterer{contract: contract}, nil
}

// bindContractTriggerXSafeFactory binds a generic wrapper to an already deployed contract.
func bindContractTriggerXSafeFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractTriggerXSafeFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractTriggerXSafeFactory *ContractTriggerXSafeFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractTriggerXSafeFactory.Contract.ContractTriggerXSafeFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractTriggerXSafeFactory *ContractTriggerXSafeFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractTriggerXSafeFactory.Contract.ContractTriggerXSafeFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractTriggerXSafeFactory *ContractTriggerXSafeFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractTriggerXSafeFactory.Contract.ContractTriggerXSafeFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractTriggerXSafeFactory *ContractTriggerXSafeFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractTriggerXSafeFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractTriggerXSafeFactory *ContractTriggerXSafeFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractTriggerXSafeFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractTriggerXSafeFactory *ContractTriggerXSafeFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractTriggerXSafeFactory.Contract.contract.Transact(opts, method, params...)
}

// GetSafeWallets is a free data retrieval call binding the contract method 0xe5e19ca9.
//
// Solidity: function getSafeWallets(address user) view returns(address[])
func (_ContractTriggerXSafeFactory *ContractTriggerXSafeFactoryCaller) GetSafeWallets(opts *bind.CallOpts, user common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _ContractTriggerXSafeFactory.contract.Call(opts, &out, "getSafeWallets", user)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSafeWallets is a free data retrieval call binding the contract method 0xe5e19ca9.
//
// Solidity: function getSafeWallets(address user) view returns(address[])
func (_ContractTriggerXSafeFactory *ContractTriggerXSafeFactorySession) GetSafeWallets(user common.Address) ([]common.Address, error) {
	return _ContractTriggerXSafeFactory.Contract.GetSafeWallets(&_ContractTriggerXSafeFactory.CallOpts, user)
}

// GetSafeWallets is a free data retrieval call binding the contract method 0xe5e19ca9.
//
// Solidity: function getSafeWallets(address user) view returns(address[])
func (_ContractTriggerXSafeFactory *ContractTriggerXSafeFactoryCallerSession) GetSafeWallets(user common.Address) ([]common.Address, error) {
	return _ContractTriggerXSafeFactory.Contract.GetSafeWallets(&_ContractTriggerXSafeFactory.CallOpts, user)
}

// GetUserSaltNonce is a free data retrieval call binding the contract method 0x850d26b6.
//
// Solidity: function getUserSaltNonce(address user) view returns(uint256)
func (_ContractTriggerXSafeFactory *ContractTriggerXSafeFactoryCaller) GetUserSaltNonce(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractTriggerXSafeFactory.contract.Call(opts, &out, "getUserSaltNonce", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserSaltNonce is a free data retrieval call binding the contract method 0x850d26b6.
//
// Solidity: function getUserSaltNonce(address user) view returns(uint256)
func (_ContractTriggerXSafeFactory *ContractTriggerXSafeFactorySession) GetUserSaltNonce(user common.Address) (*big.Int, error) {
	return _ContractTriggerXSafeFactory.Contract.GetUserSaltNonce(&_ContractTriggerXSafeFactory.CallOpts, user)
}

// GetUserSaltNonce is a free data retrieval call binding the contract method 0x850d26b6.
//
// Solidity: function getUserSaltNonce(address user) view returns(uint256)
func (_ContractTriggerXSafeFactory *ContractTriggerXSafeFactoryCallerSession) GetUserSaltNonce(user common.Address) (*big.Int, error) {
	return _ContractTriggerXSafeFactory.Contract.GetUserSaltNonce(&_ContractTriggerXSafeFactory.CallOpts, user)
}

// LatestSafeWallet is a free data retrieval call binding the contract method 0xe1e66866.
//
// Solidity: function latestSafeWallet(address user) view returns(address)
func (_ContractTriggerXSafeFactory *ContractTriggerXSafeFactoryCaller) LatestSafeWallet(opts *bind.CallOpts, user common.Address) (common.Address, error) {
	var out []interface{}
	err := _ContractTriggerXSafeFactory.contract.Call(opts, &out, "latestSafeWallet", user)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LatestSafeWallet is a free data retrieval call binding the contract method 0xe1e66866.
//
// Solidity: function latestSafeWallet(address user) view returns(address)
func (_ContractTriggerXSafeFactory *ContractTriggerXSafeFactorySession) LatestSafeWallet(user common.Address) (common.Address, error) {
	return _ContractTriggerXSafeFactory.Contract.LatestSafeWallet(&_ContractTriggerXSafeFactory.CallOpts, user)
}

// LatestSafeWallet is a free data retrieval call binding the contract method 0xe1e66866.
//
// Solidity: function latestSafeWallet(address user) view returns(address)
func (_ContractTriggerXSafeFactory *ContractTriggerXSafeFactoryCallerSession) LatestSafeWallet(user common.Address) (common.Address, error) {
	return _ContractTriggerXSafeFactory.Contract.LatestSafeWallet(&_ContractTriggerXSafeFactory.CallOpts, user)
}

// PredictSafeAddress is a free data retrieval call binding the contract method 0x9ee0ce27.
//
// Solidity: function predictSafeAddress(address user) view returns(address predictedAddress)
func (_ContractTriggerXSafeFactory *ContractTriggerXSafeFactoryCaller) PredictSafeAddress(opts *bind.CallOpts, user common.Address) (common.Address, error) {
	var out []interface{}
	err := _ContractTriggerXSafeFactory.contract.Call(opts, &out, "predictSafeAddress", user)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PredictSafeAddress is a free data retrieval call binding the contract method 0x9ee0ce27.
//
// Solidity: function predictSafeAddress(address user) view returns(address predictedAddress)
func (_ContractTriggerXSafeFactory *ContractTriggerXSafeFactorySession) PredictSafeAddress(user common.Address) (common.Address, error) {
	return _ContractTriggerXSafeFactory.Contract.PredictSafeAddress(&_ContractTriggerXSafeFactory.CallOpts, user)
}

// PredictSafeAddress is a free data retrieval call binding the contract method 0x9ee0ce27.
//
// Solidity: function predictSafeAddress(address user) view returns(address predictedAddress)
func (_ContractTriggerXSafeFactory *ContractTriggerXSafeFactoryCallerSession) PredictSafeAddress(user common.Address) (common.Address, error) {
	return _ContractTriggerXSafeFactory.Contract.PredictSafeAddress(&_ContractTriggerXSafeFactory.CallOpts, user)
}

// SafeProxyFactory is a free data retrieval call binding the contract method 0x19964501.
//
// Solidity: function safeProxyFactory() view returns(address)
func (_ContractTriggerXSafeFactory *ContractTriggerXSafeFactoryCaller) SafeProxyFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTriggerXSafeFactory.contract.Call(opts, &out, "safeProxyFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SafeProxyFactory is a free data retrieval call binding the contract method 0x19964501.
//
// Solidity: function safeProxyFactory() view returns(address)
func (_ContractTriggerXSafeFactory *ContractTriggerXSafeFactorySession) SafeProxyFactory() (common.Address, error) {
	return _ContractTriggerXSafeFactory.Contract.SafeProxyFactory(&_ContractTriggerXSafeFactory.CallOpts)
}

// SafeProxyFactory is a free data retrieval call binding the contract method 0x19964501.
//
// Solidity: function safeProxyFactory() view returns(address)
func (_ContractTriggerXSafeFactory *ContractTriggerXSafeFactoryCallerSession) SafeProxyFactory() (common.Address, error) {
	return _ContractTriggerXSafeFactory.Contract.SafeProxyFactory(&_ContractTriggerXSafeFactory.CallOpts)
}

// SafeSingleton is a free data retrieval call binding the contract method 0xac7d146b.
//
// Solidity: function safeSingleton() view returns(address)
func (_ContractTriggerXSafeFactory *ContractTriggerXSafeFactoryCaller) SafeSingleton(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTriggerXSafeFactory.contract.Call(opts, &out, "safeSingleton")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SafeSingleton is a free data retrieval call binding the contract method 0xac7d146b.
//
// Solidity: function safeSingleton() view returns(address)
func (_ContractTriggerXSafeFactory *ContractTriggerXSafeFactorySession) SafeSingleton() (common.Address, error) {
	return _ContractTriggerXSafeFactory.Contract.SafeSingleton(&_ContractTriggerXSafeFactory.CallOpts)
}

// SafeSingleton is a free data retrieval call binding the contract method 0xac7d146b.
//
// Solidity: function safeSingleton() view returns(address)
func (_ContractTriggerXSafeFactory *ContractTriggerXSafeFactoryCallerSession) SafeSingleton() (common.Address, error) {
	return _ContractTriggerXSafeFactory.Contract.SafeSingleton(&_ContractTriggerXSafeFactory.CallOpts)
}

// TaskExecutionHub is a free data retrieval call binding the contract method 0xc54d346e.
//
// Solidity: function taskExecutionHub() view returns(address)
func (_ContractTriggerXSafeFactory *ContractTriggerXSafeFactoryCaller) TaskExecutionHub(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTriggerXSafeFactory.contract.Call(opts, &out, "taskExecutionHub")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TaskExecutionHub is a free data retrieval call binding the contract method 0xc54d346e.
//
// Solidity: function taskExecutionHub() view returns(address)
func (_ContractTriggerXSafeFactory *ContractTriggerXSafeFactorySession) TaskExecutionHub() (common.Address, error) {
	return _ContractTriggerXSafeFactory.Contract.TaskExecutionHub(&_ContractTriggerXSafeFactory.CallOpts)
}

// TaskExecutionHub is a free data retrieval call binding the contract method 0xc54d346e.
//
// Solidity: function taskExecutionHub() view returns(address)
func (_ContractTriggerXSafeFactory *ContractTriggerXSafeFactoryCallerSession) TaskExecutionHub() (common.Address, error) {
	return _ContractTriggerXSafeFactory.Contract.TaskExecutionHub(&_ContractTriggerXSafeFactory.CallOpts)
}

// CreateSafeWallet is a paid mutator transaction binding the contract method 0xbc632809.
//
// Solidity: function createSafeWallet(address user) returns(address safeAddress)
func (_ContractTriggerXSafeFactory *ContractTriggerXSafeFactoryTransactor) CreateSafeWallet(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _ContractTriggerXSafeFactory.contract.Transact(opts, "createSafeWallet", user)
}

// CreateSafeWallet is a paid mutator transaction binding the contract method 0xbc632809.
//
// Solidity: function createSafeWallet(address user) returns(address safeAddress)
func (_ContractTriggerXSafeFactory *ContractTriggerXSafeFactorySession) CreateSafeWallet(user common.Address) (*types.Transaction, error) {
	return _ContractTriggerXSafeFactory.Contract.CreateSafeWallet(&_ContractTriggerXSafeFactory.TransactOpts, user)
}

// CreateSafeWallet is a paid mutator transaction binding the contract method 0xbc632809.
//
// Solidity: function createSafeWallet(address user) returns(address safeAddress)
func (_ContractTriggerXSafeFactory *ContractTriggerXSafeFactoryTransactorSession) CreateSafeWallet(user common.Address) (*types.Transaction, error) {
	return _ContractTriggerXSafeFactory.Contract.CreateSafeWallet(&_ContractTriggerXSafeFactory.TransactOpts, user)
}

// ContractTriggerXSafeFactorySafeWalletCreatedIterator is returned from FilterSafeWalletCreated and is used to iterate over the raw logs and unpacked data for SafeWalletCreated events raised by the ContractTriggerXSafeFactory contract.
type ContractTriggerXSafeFactorySafeWalletCreatedIterator struct {
	Event *ContractTriggerXSafeFactorySafeWalletCreated // Event containing the contract specifics and raw log

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
func (it *ContractTriggerXSafeFactorySafeWalletCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerXSafeFactorySafeWalletCreated)
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
		it.Event = new(ContractTriggerXSafeFactorySafeWalletCreated)
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
func (it *ContractTriggerXSafeFactorySafeWalletCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerXSafeFactorySafeWalletCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerXSafeFactorySafeWalletCreated represents a SafeWalletCreated event raised by the ContractTriggerXSafeFactory contract.
type ContractTriggerXSafeFactorySafeWalletCreated struct {
	User       common.Address
	SafeWallet common.Address
	SaltNonce  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSafeWalletCreated is a free log retrieval operation binding the contract event 0xa0771853c946ed7d6b5434c55582f8a82e282036fb4b0512878e4fab361e605f.
//
// Solidity: event SafeWalletCreated(address indexed user, address indexed safeWallet, uint256 saltNonce)
func (_ContractTriggerXSafeFactory *ContractTriggerXSafeFactoryFilterer) FilterSafeWalletCreated(opts *bind.FilterOpts, user []common.Address, safeWallet []common.Address) (*ContractTriggerXSafeFactorySafeWalletCreatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var safeWalletRule []interface{}
	for _, safeWalletItem := range safeWallet {
		safeWalletRule = append(safeWalletRule, safeWalletItem)
	}

	logs, sub, err := _ContractTriggerXSafeFactory.contract.FilterLogs(opts, "SafeWalletCreated", userRule, safeWalletRule)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXSafeFactorySafeWalletCreatedIterator{contract: _ContractTriggerXSafeFactory.contract, event: "SafeWalletCreated", logs: logs, sub: sub}, nil
}

// WatchSafeWalletCreated is a free log subscription operation binding the contract event 0xa0771853c946ed7d6b5434c55582f8a82e282036fb4b0512878e4fab361e605f.
//
// Solidity: event SafeWalletCreated(address indexed user, address indexed safeWallet, uint256 saltNonce)
func (_ContractTriggerXSafeFactory *ContractTriggerXSafeFactoryFilterer) WatchSafeWalletCreated(opts *bind.WatchOpts, sink chan<- *ContractTriggerXSafeFactorySafeWalletCreated, user []common.Address, safeWallet []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var safeWalletRule []interface{}
	for _, safeWalletItem := range safeWallet {
		safeWalletRule = append(safeWalletRule, safeWalletItem)
	}

	logs, sub, err := _ContractTriggerXSafeFactory.contract.WatchLogs(opts, "SafeWalletCreated", userRule, safeWalletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerXSafeFactorySafeWalletCreated)
				if err := _ContractTriggerXSafeFactory.contract.UnpackLog(event, "SafeWalletCreated", log); err != nil {
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

// ParseSafeWalletCreated is a log parse operation binding the contract event 0xa0771853c946ed7d6b5434c55582f8a82e282036fb4b0512878e4fab361e605f.
//
// Solidity: event SafeWalletCreated(address indexed user, address indexed safeWallet, uint256 saltNonce)
func (_ContractTriggerXSafeFactory *ContractTriggerXSafeFactoryFilterer) ParseSafeWalletCreated(log types.Log) (*ContractTriggerXSafeFactorySafeWalletCreated, error) {
	event := new(ContractTriggerXSafeFactorySafeWalletCreated)
	if err := _ContractTriggerXSafeFactory.contract.UnpackLog(event, "SafeWalletCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
