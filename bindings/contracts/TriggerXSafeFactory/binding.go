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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_safeProxyFactory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_safeSingleton\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createSafeWallet\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"safeAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getSafeWallets\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUserSaltNonce\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestSafeWallet\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"predictSafeAddress\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"predictedAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"safeProxyFactory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"safeSingleton\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"SafeWalletCreated\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"safeWallet\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"saltNonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x60c03461012457601f610a5138819003918201601f19168301916001600160401b0383118484101761012857808492604094855283398101031261012457610052602061004b8361013c565b920161013c565b906001600160a01b038116156100ea576001600160a01b038216156100b45760805260a052604051610900908161015182396080518181816102690152818161054601526105ea015260a05181818161020d015281816103e101526104f40152f35b60405162461bcd60e51b815260206004820152600e60248201526d2d32b9379039b4b733b632ba37b760911b6044820152606490fd5b60405162461bcd60e51b81526020600482015260126024820152715a65726f2070726f787920666163746f727960701b6044820152606490fd5b5f80fd5b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b03821682036101245756fe6080806040526004361015610012575f80fd5b5f3560e01c90816319964501146105d857508063850d26b61461059c5780639ee0ce2714610410578063ac7d146b146103cc578063bc63280914610144578063e1e66866146101035763e5e19ca914610069575f80fd5b346100ff5760203660031901126100ff576004356001600160a01b038116908190036100ff575f525f60205260405f206040519081602082549182815201915f5260205f20905f5b8181106100e0576100dc856100c88187038261068d565b604051918291602083526020830190610619565b0390f35b82546001600160a01b03168452602090930192600192830192016100b1565b5f80fd5b346100ff5760203660031901126100ff576004356001600160a01b03811681036100ff5761013260209161072f565b6040516001600160a01b039091168152f35b346100ff5760203660031901126100ff576004356001600160a01b038116908190036100ff57610175811515610655565b805f52600160205260405f2054600181018082116103b857825f52600160205260405f2055604080516101a8828261068d565b60018152601f198201366020830137836101c1826106af565b52602082516101f2816101e48482019563b63e800d60e01b8752602483016106d0565b03601f19810183528261068d565b6084845180948193631688f0b960e01b835260018060a01b037f0000000000000000000000000000000000000000000000000000000000000000166004840152606060248401525180918160648501528484015e5f838284010152876044830152601f801991011681010301815f60018060a01b037f0000000000000000000000000000000000000000000000000000000000000000165af19081156103ae575f9161036c575b506001600160a01b031691821561033b57835f525f602052815f20938454946801000000000000000086101561032757602085936102ff887fa0771853c946ed7d6b5434c55582f8a82e282036fb4b0512878e4fab361e605f946001859b01815561071a565b81546001600160a01b0360039290921b91821b19169087901b1790558551908152a351908152f35b634e487b7160e01b5f52604160045260245ffd5b815162461bcd60e51b815260206004820152600a6024820152695a65726f2070726f787960b01b6044820152606490fd5b90506020813d6020116103a6575b816103876020938361068d565b810103126100ff57516001600160a01b03811681036100ff5784610299565b3d915061037a565b82513d5f823e3d90fd5b634e487b7160e01b5f52601160045260245ffd5b346100ff575f3660031901126100ff576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b346100ff5760203660031901126100ff576004356001600160a01b038116908190036100ff57806104446020921515610655565b805f526001825260405f2054604091825190610460848361068d565b60018252601f1984013686840137610477826106af565b528251610498816101e48782019463b63e800d60e01b8652602483016106d0565b51902090825190848201928352838201528281526104b760608261068d565b519020815161014e6104cb8582018361068d565b8082528482019061077d823961052485808651809482820196518091885e810160018060a01b037f000000000000000000000000000000000000000000000000000000000000000016838201520301808452018261068d565b519020908251918483019160ff60f81b83526bffffffffffffffffffffffff197f000000000000000000000000000000000000000000000000000000000000000060601b166021850152603584015260558301526055825261058760758361068d565b9051902090516001600160a01b039091168152f35b346100ff5760203660031901126100ff576004356001600160a01b038116908190036100ff575f526001602052602060405f2054604051908152f35b346100ff575f3660031901126100ff577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b90602080835192838152019201905f5b8181106106365750505090565b82516001600160a01b0316845260209384019390920191600101610629565b1561065c57565b60405162461bcd60e51b81526020600482015260096024820152682d32b937903ab9b2b960b91b6044820152606490fd5b90601f8019910116810190811067ffffffffffffffff82111761032757604052565b8051156106bc5760200190565b634e487b7160e01b5f52603260045260245ffd5b919060e060206106eb5f936101008752610100870190610619565b600182870152836040870152858103606087015283815201938260808201528260a08201528260c08201520152565b80548210156106bc575f5260205f2001905f90565b6001600160a01b03165f90815260208190526040902080548015610776575f1981019081116103b8576107619161071a565b905460039190911b1c6001600160a01b031690565b50505f9056fe60803460c357601f61014e38819003918201601f19168301916001600160401b0383118484101760c75780849260209460405283398101031260c357516001600160a01b0381169081900360c35780156073575f80546001600160a01b031916919091179055604051607290816100dc8239f35b60405162461bcd60e51b815260206004820152602260248201527f496e76616c69642073696e676c65746f6e20616464726573732070726f766964604482015261195960f21b6064820152608490fd5b5f80fd5b634e487b7160e01b5f52604160045260245ffdfe60806040525f5463a619486e5f3560e01c146030575f8091368280378136915af43d5f803e15602c573d5ff35b3d5ffd5b60601b606c5260206060f3fea2646970667358221220a7f111436f7c80bd11946d0802a794a634c0243f2af12030ba4f1fc6c4d3dce464736f6c634300081b0033a26469706673582212208502aa814a63a2ab765853f63c8c0a2ce3435bae679e404d339304913f6ecad164736f6c634300081b0033",
}

// ContractTriggerXSafeFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractTriggerXSafeFactoryMetaData.ABI instead.
var ContractTriggerXSafeFactoryABI = ContractTriggerXSafeFactoryMetaData.ABI

// ContractTriggerXSafeFactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractTriggerXSafeFactoryMetaData.Bin instead.
var ContractTriggerXSafeFactoryBin = ContractTriggerXSafeFactoryMetaData.Bin

// DeployContractTriggerXSafeFactory deploys a new Ethereum contract, binding an instance of ContractTriggerXSafeFactory to it.
func DeployContractTriggerXSafeFactory(auth *bind.TransactOpts, backend bind.ContractBackend, _safeProxyFactory common.Address, _safeSingleton common.Address) (common.Address, *types.Transaction, *ContractTriggerXSafeFactory, error) {
	parsed, err := ContractTriggerXSafeFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractTriggerXSafeFactoryBin), backend, _safeProxyFactory, _safeSingleton)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractTriggerXSafeFactory{ContractTriggerXSafeFactoryCaller: ContractTriggerXSafeFactoryCaller{contract: contract}, ContractTriggerXSafeFactoryTransactor: ContractTriggerXSafeFactoryTransactor{contract: contract}, ContractTriggerXSafeFactoryFilterer: ContractTriggerXSafeFactoryFilterer{contract: contract}}, nil
}

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
