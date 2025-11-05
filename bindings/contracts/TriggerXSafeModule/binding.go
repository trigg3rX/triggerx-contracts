// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractTriggerXSafeModule

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

// ContractTriggerXSafeModuleMetaData contains all meta data concerning the ContractTriggerXSafeModule contract.
var ContractTriggerXSafeModuleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_taskExecutionHub\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ExecFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotTaskExecutionHub\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"safeAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"TaskExecutedFromModule\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"safeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actionTarget\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actionValue\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"actionData\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"operation\",\"type\":\"uint8\"}],\"name\":\"execJobFromHub\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taskExecutionHub\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// ContractTriggerXSafeModuleABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractTriggerXSafeModuleMetaData.ABI instead.
var ContractTriggerXSafeModuleABI = ContractTriggerXSafeModuleMetaData.ABI

// ContractTriggerXSafeModuleMethods is an auto generated interface around an Ethereum contract.
type ContractTriggerXSafeModuleMethods interface {
	ContractTriggerXSafeModuleCalls
	ContractTriggerXSafeModuleTransacts
	ContractTriggerXSafeModuleFilters
}

// ContractTriggerXSafeModuleCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractTriggerXSafeModuleCalls interface {
	TaskExecutionHub(opts *bind.CallOpts) (common.Address, error)
}

// ContractTriggerXSafeModuleTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractTriggerXSafeModuleTransacts interface {
	ExecJobFromHub(opts *bind.TransactOpts, safeAddress common.Address, actionTarget common.Address, actionValue *big.Int, actionData []byte, operation uint8) (*types.Transaction, error)
}

// ContractTriggerXSafeModuleFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractTriggerXSafeModuleFilters interface {
	FilterTaskExecutedFromModule(opts *bind.FilterOpts, safeAddress []common.Address, executor []common.Address) (*ContractTriggerXSafeModuleTaskExecutedFromModuleIterator, error)
	WatchTaskExecutedFromModule(opts *bind.WatchOpts, sink chan<- *ContractTriggerXSafeModuleTaskExecutedFromModule, safeAddress []common.Address, executor []common.Address) (event.Subscription, error)
	ParseTaskExecutedFromModule(log types.Log) (*ContractTriggerXSafeModuleTaskExecutedFromModule, error)
}

// ContractTriggerXSafeModule is an auto generated Go binding around an Ethereum contract.
type ContractTriggerXSafeModule struct {
	ContractTriggerXSafeModuleCaller     // Read-only binding to the contract
	ContractTriggerXSafeModuleTransactor // Write-only binding to the contract
	ContractTriggerXSafeModuleFilterer   // Log filterer for contract events
}

// ContractTriggerXSafeModule implements the ContractTriggerXSafeModuleMethods interface.
var _ ContractTriggerXSafeModuleMethods = (*ContractTriggerXSafeModule)(nil)

// ContractTriggerXSafeModuleCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractTriggerXSafeModuleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTriggerXSafeModuleCaller implements the ContractTriggerXSafeModuleCalls interface.
var _ ContractTriggerXSafeModuleCalls = (*ContractTriggerXSafeModuleCaller)(nil)

// ContractTriggerXSafeModuleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTriggerXSafeModuleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTriggerXSafeModuleTransactor implements the ContractTriggerXSafeModuleTransacts interface.
var _ ContractTriggerXSafeModuleTransacts = (*ContractTriggerXSafeModuleTransactor)(nil)

// ContractTriggerXSafeModuleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractTriggerXSafeModuleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTriggerXSafeModuleFilterer implements the ContractTriggerXSafeModuleFilters interface.
var _ ContractTriggerXSafeModuleFilters = (*ContractTriggerXSafeModuleFilterer)(nil)

// ContractTriggerXSafeModuleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractTriggerXSafeModuleSession struct {
	Contract     *ContractTriggerXSafeModule // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ContractTriggerXSafeModuleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractTriggerXSafeModuleCallerSession struct {
	Contract *ContractTriggerXSafeModuleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// ContractTriggerXSafeModuleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTriggerXSafeModuleTransactorSession struct {
	Contract     *ContractTriggerXSafeModuleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// ContractTriggerXSafeModuleRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractTriggerXSafeModuleRaw struct {
	Contract *ContractTriggerXSafeModule // Generic contract binding to access the raw methods on
}

// ContractTriggerXSafeModuleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractTriggerXSafeModuleCallerRaw struct {
	Contract *ContractTriggerXSafeModuleCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTriggerXSafeModuleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTriggerXSafeModuleTransactorRaw struct {
	Contract *ContractTriggerXSafeModuleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractTriggerXSafeModule creates a new instance of ContractTriggerXSafeModule, bound to a specific deployed contract.
func NewContractTriggerXSafeModule(address common.Address, backend bind.ContractBackend) (*ContractTriggerXSafeModule, error) {
	contract, err := bindContractTriggerXSafeModule(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXSafeModule{ContractTriggerXSafeModuleCaller: ContractTriggerXSafeModuleCaller{contract: contract}, ContractTriggerXSafeModuleTransactor: ContractTriggerXSafeModuleTransactor{contract: contract}, ContractTriggerXSafeModuleFilterer: ContractTriggerXSafeModuleFilterer{contract: contract}}, nil
}

// NewContractTriggerXSafeModuleCaller creates a new read-only instance of ContractTriggerXSafeModule, bound to a specific deployed contract.
func NewContractTriggerXSafeModuleCaller(address common.Address, caller bind.ContractCaller) (*ContractTriggerXSafeModuleCaller, error) {
	contract, err := bindContractTriggerXSafeModule(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXSafeModuleCaller{contract: contract}, nil
}

// NewContractTriggerXSafeModuleTransactor creates a new write-only instance of ContractTriggerXSafeModule, bound to a specific deployed contract.
func NewContractTriggerXSafeModuleTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTriggerXSafeModuleTransactor, error) {
	contract, err := bindContractTriggerXSafeModule(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXSafeModuleTransactor{contract: contract}, nil
}

// NewContractTriggerXSafeModuleFilterer creates a new log filterer instance of ContractTriggerXSafeModule, bound to a specific deployed contract.
func NewContractTriggerXSafeModuleFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractTriggerXSafeModuleFilterer, error) {
	contract, err := bindContractTriggerXSafeModule(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXSafeModuleFilterer{contract: contract}, nil
}

// bindContractTriggerXSafeModule binds a generic wrapper to an already deployed contract.
func bindContractTriggerXSafeModule(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractTriggerXSafeModuleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractTriggerXSafeModule *ContractTriggerXSafeModuleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractTriggerXSafeModule.Contract.ContractTriggerXSafeModuleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractTriggerXSafeModule *ContractTriggerXSafeModuleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractTriggerXSafeModule.Contract.ContractTriggerXSafeModuleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractTriggerXSafeModule *ContractTriggerXSafeModuleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractTriggerXSafeModule.Contract.ContractTriggerXSafeModuleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractTriggerXSafeModule *ContractTriggerXSafeModuleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractTriggerXSafeModule.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractTriggerXSafeModule *ContractTriggerXSafeModuleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractTriggerXSafeModule.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractTriggerXSafeModule *ContractTriggerXSafeModuleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractTriggerXSafeModule.Contract.contract.Transact(opts, method, params...)
}

// TaskExecutionHub is a free data retrieval call binding the contract method 0xc54d346e.
//
// Solidity: function taskExecutionHub() view returns(address)
func (_ContractTriggerXSafeModule *ContractTriggerXSafeModuleCaller) TaskExecutionHub(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTriggerXSafeModule.contract.Call(opts, &out, "taskExecutionHub")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TaskExecutionHub is a free data retrieval call binding the contract method 0xc54d346e.
//
// Solidity: function taskExecutionHub() view returns(address)
func (_ContractTriggerXSafeModule *ContractTriggerXSafeModuleSession) TaskExecutionHub() (common.Address, error) {
	return _ContractTriggerXSafeModule.Contract.TaskExecutionHub(&_ContractTriggerXSafeModule.CallOpts)
}

// TaskExecutionHub is a free data retrieval call binding the contract method 0xc54d346e.
//
// Solidity: function taskExecutionHub() view returns(address)
func (_ContractTriggerXSafeModule *ContractTriggerXSafeModuleCallerSession) TaskExecutionHub() (common.Address, error) {
	return _ContractTriggerXSafeModule.Contract.TaskExecutionHub(&_ContractTriggerXSafeModule.CallOpts)
}

// ExecJobFromHub is a paid mutator transaction binding the contract method 0xb36e8518.
//
// Solidity: function execJobFromHub(address safeAddress, address actionTarget, uint256 actionValue, bytes actionData, uint8 operation) returns(bool success)
func (_ContractTriggerXSafeModule *ContractTriggerXSafeModuleTransactor) ExecJobFromHub(opts *bind.TransactOpts, safeAddress common.Address, actionTarget common.Address, actionValue *big.Int, actionData []byte, operation uint8) (*types.Transaction, error) {
	return _ContractTriggerXSafeModule.contract.Transact(opts, "execJobFromHub", safeAddress, actionTarget, actionValue, actionData, operation)
}

// ExecJobFromHub is a paid mutator transaction binding the contract method 0xb36e8518.
//
// Solidity: function execJobFromHub(address safeAddress, address actionTarget, uint256 actionValue, bytes actionData, uint8 operation) returns(bool success)
func (_ContractTriggerXSafeModule *ContractTriggerXSafeModuleSession) ExecJobFromHub(safeAddress common.Address, actionTarget common.Address, actionValue *big.Int, actionData []byte, operation uint8) (*types.Transaction, error) {
	return _ContractTriggerXSafeModule.Contract.ExecJobFromHub(&_ContractTriggerXSafeModule.TransactOpts, safeAddress, actionTarget, actionValue, actionData, operation)
}

// ExecJobFromHub is a paid mutator transaction binding the contract method 0xb36e8518.
//
// Solidity: function execJobFromHub(address safeAddress, address actionTarget, uint256 actionValue, bytes actionData, uint8 operation) returns(bool success)
func (_ContractTriggerXSafeModule *ContractTriggerXSafeModuleTransactorSession) ExecJobFromHub(safeAddress common.Address, actionTarget common.Address, actionValue *big.Int, actionData []byte, operation uint8) (*types.Transaction, error) {
	return _ContractTriggerXSafeModule.Contract.ExecJobFromHub(&_ContractTriggerXSafeModule.TransactOpts, safeAddress, actionTarget, actionValue, actionData, operation)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ContractTriggerXSafeModule *ContractTriggerXSafeModuleTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _ContractTriggerXSafeModule.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ContractTriggerXSafeModule *ContractTriggerXSafeModuleSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ContractTriggerXSafeModule.Contract.Fallback(&_ContractTriggerXSafeModule.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ContractTriggerXSafeModule *ContractTriggerXSafeModuleTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ContractTriggerXSafeModule.Contract.Fallback(&_ContractTriggerXSafeModule.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ContractTriggerXSafeModule *ContractTriggerXSafeModuleTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractTriggerXSafeModule.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ContractTriggerXSafeModule *ContractTriggerXSafeModuleSession) Receive() (*types.Transaction, error) {
	return _ContractTriggerXSafeModule.Contract.Receive(&_ContractTriggerXSafeModule.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ContractTriggerXSafeModule *ContractTriggerXSafeModuleTransactorSession) Receive() (*types.Transaction, error) {
	return _ContractTriggerXSafeModule.Contract.Receive(&_ContractTriggerXSafeModule.TransactOpts)
}

// ContractTriggerXSafeModuleTaskExecutedFromModuleIterator is returned from FilterTaskExecutedFromModule and is used to iterate over the raw logs and unpacked data for TaskExecutedFromModule events raised by the ContractTriggerXSafeModule contract.
type ContractTriggerXSafeModuleTaskExecutedFromModuleIterator struct {
	Event *ContractTriggerXSafeModuleTaskExecutedFromModule // Event containing the contract specifics and raw log

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
func (it *ContractTriggerXSafeModuleTaskExecutedFromModuleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerXSafeModuleTaskExecutedFromModule)
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
		it.Event = new(ContractTriggerXSafeModuleTaskExecutedFromModule)
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
func (it *ContractTriggerXSafeModuleTaskExecutedFromModuleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerXSafeModuleTaskExecutedFromModuleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerXSafeModuleTaskExecutedFromModule represents a TaskExecutedFromModule event raised by the ContractTriggerXSafeModule contract.
type ContractTriggerXSafeModuleTaskExecutedFromModule struct {
	SafeAddress common.Address
	Executor    common.Address
	Success     bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTaskExecutedFromModule is a free log retrieval operation binding the contract event 0x528c00f3d7da8e9429355ebfd96f7db55a1b2363c6cf6eba8f256710c6f03c07.
//
// Solidity: event TaskExecutedFromModule(address indexed safeAddress, address indexed executor, bool success)
func (_ContractTriggerXSafeModule *ContractTriggerXSafeModuleFilterer) FilterTaskExecutedFromModule(opts *bind.FilterOpts, safeAddress []common.Address, executor []common.Address) (*ContractTriggerXSafeModuleTaskExecutedFromModuleIterator, error) {

	var safeAddressRule []interface{}
	for _, safeAddressItem := range safeAddress {
		safeAddressRule = append(safeAddressRule, safeAddressItem)
	}
	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _ContractTriggerXSafeModule.contract.FilterLogs(opts, "TaskExecutedFromModule", safeAddressRule, executorRule)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXSafeModuleTaskExecutedFromModuleIterator{contract: _ContractTriggerXSafeModule.contract, event: "TaskExecutedFromModule", logs: logs, sub: sub}, nil
}

// WatchTaskExecutedFromModule is a free log subscription operation binding the contract event 0x528c00f3d7da8e9429355ebfd96f7db55a1b2363c6cf6eba8f256710c6f03c07.
//
// Solidity: event TaskExecutedFromModule(address indexed safeAddress, address indexed executor, bool success)
func (_ContractTriggerXSafeModule *ContractTriggerXSafeModuleFilterer) WatchTaskExecutedFromModule(opts *bind.WatchOpts, sink chan<- *ContractTriggerXSafeModuleTaskExecutedFromModule, safeAddress []common.Address, executor []common.Address) (event.Subscription, error) {

	var safeAddressRule []interface{}
	for _, safeAddressItem := range safeAddress {
		safeAddressRule = append(safeAddressRule, safeAddressItem)
	}
	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _ContractTriggerXSafeModule.contract.WatchLogs(opts, "TaskExecutedFromModule", safeAddressRule, executorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerXSafeModuleTaskExecutedFromModule)
				if err := _ContractTriggerXSafeModule.contract.UnpackLog(event, "TaskExecutedFromModule", log); err != nil {
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

// ParseTaskExecutedFromModule is a log parse operation binding the contract event 0x528c00f3d7da8e9429355ebfd96f7db55a1b2363c6cf6eba8f256710c6f03c07.
//
// Solidity: event TaskExecutedFromModule(address indexed safeAddress, address indexed executor, bool success)
func (_ContractTriggerXSafeModule *ContractTriggerXSafeModuleFilterer) ParseTaskExecutedFromModule(log types.Log) (*ContractTriggerXSafeModuleTaskExecutedFromModule, error) {
	event := new(ContractTriggerXSafeModuleTaskExecutedFromModule)
	if err := _ContractTriggerXSafeModule.contract.UnpackLog(event, "TaskExecutedFromModule", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
