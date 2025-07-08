// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package taskdefinition

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

// TaskDefinitionLibraryMetaData contains all meta data concerning the TaskDefinitionLibrary contract.
var TaskDefinitionLibraryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"event\",\"name\":\"TaskDefinitionCreated\",\"inputs\":[{\"name\":\"taskDefinitionId\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"baseRewardFeeForAttesters\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"baseRewardFeeForPerformer\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"baseRewardFeeForAggregator\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"minimumVotingPower\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"restrictedOperatorIds\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"maximumNumberOfAttesters\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// TaskDefinitionLibraryABI is the input ABI used to generate the binding from.
// Deprecated: Use TaskDefinitionLibraryMetaData.ABI instead.
var TaskDefinitionLibraryABI = TaskDefinitionLibraryMetaData.ABI

// TaskDefinitionLibrary is an auto generated Go binding around an Ethereum contract.
type TaskDefinitionLibrary struct {
	TaskDefinitionLibraryCaller     // Read-only binding to the contract
	TaskDefinitionLibraryTransactor // Write-only binding to the contract
	TaskDefinitionLibraryFilterer   // Log filterer for contract events
}

// TaskDefinitionLibraryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TaskDefinitionLibraryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskDefinitionLibraryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TaskDefinitionLibraryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskDefinitionLibraryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TaskDefinitionLibraryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskDefinitionLibrarySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TaskDefinitionLibrarySession struct {
	Contract     *TaskDefinitionLibrary // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TaskDefinitionLibraryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TaskDefinitionLibraryCallerSession struct {
	Contract *TaskDefinitionLibraryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// TaskDefinitionLibraryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TaskDefinitionLibraryTransactorSession struct {
	Contract     *TaskDefinitionLibraryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// TaskDefinitionLibraryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TaskDefinitionLibraryRaw struct {
	Contract *TaskDefinitionLibrary // Generic contract binding to access the raw methods on
}

// TaskDefinitionLibraryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TaskDefinitionLibraryCallerRaw struct {
	Contract *TaskDefinitionLibraryCaller // Generic read-only contract binding to access the raw methods on
}

// TaskDefinitionLibraryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TaskDefinitionLibraryTransactorRaw struct {
	Contract *TaskDefinitionLibraryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTaskDefinitionLibrary creates a new instance of TaskDefinitionLibrary, bound to a specific deployed contract.
func NewTaskDefinitionLibrary(address common.Address, backend bind.ContractBackend) (*TaskDefinitionLibrary, error) {
	contract, err := bindTaskDefinitionLibrary(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TaskDefinitionLibrary{TaskDefinitionLibraryCaller: TaskDefinitionLibraryCaller{contract: contract}, TaskDefinitionLibraryTransactor: TaskDefinitionLibraryTransactor{contract: contract}, TaskDefinitionLibraryFilterer: TaskDefinitionLibraryFilterer{contract: contract}}, nil
}

// NewTaskDefinitionLibraryCaller creates a new read-only instance of TaskDefinitionLibrary, bound to a specific deployed contract.
func NewTaskDefinitionLibraryCaller(address common.Address, caller bind.ContractCaller) (*TaskDefinitionLibraryCaller, error) {
	contract, err := bindTaskDefinitionLibrary(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TaskDefinitionLibraryCaller{contract: contract}, nil
}

// NewTaskDefinitionLibraryTransactor creates a new write-only instance of TaskDefinitionLibrary, bound to a specific deployed contract.
func NewTaskDefinitionLibraryTransactor(address common.Address, transactor bind.ContractTransactor) (*TaskDefinitionLibraryTransactor, error) {
	contract, err := bindTaskDefinitionLibrary(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TaskDefinitionLibraryTransactor{contract: contract}, nil
}

// NewTaskDefinitionLibraryFilterer creates a new log filterer instance of TaskDefinitionLibrary, bound to a specific deployed contract.
func NewTaskDefinitionLibraryFilterer(address common.Address, filterer bind.ContractFilterer) (*TaskDefinitionLibraryFilterer, error) {
	contract, err := bindTaskDefinitionLibrary(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TaskDefinitionLibraryFilterer{contract: contract}, nil
}

// bindTaskDefinitionLibrary binds a generic wrapper to an already deployed contract.
func bindTaskDefinitionLibrary(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TaskDefinitionLibraryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaskDefinitionLibrary *TaskDefinitionLibraryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaskDefinitionLibrary.Contract.TaskDefinitionLibraryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaskDefinitionLibrary *TaskDefinitionLibraryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaskDefinitionLibrary.Contract.TaskDefinitionLibraryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaskDefinitionLibrary *TaskDefinitionLibraryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaskDefinitionLibrary.Contract.TaskDefinitionLibraryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaskDefinitionLibrary *TaskDefinitionLibraryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaskDefinitionLibrary.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaskDefinitionLibrary *TaskDefinitionLibraryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaskDefinitionLibrary.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaskDefinitionLibrary *TaskDefinitionLibraryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaskDefinitionLibrary.Contract.contract.Transact(opts, method, params...)
}

// TaskDefinitionLibraryTaskDefinitionCreatedIterator is returned from FilterTaskDefinitionCreated and is used to iterate over the raw logs and unpacked data for TaskDefinitionCreated events raised by the TaskDefinitionLibrary contract.
type TaskDefinitionLibraryTaskDefinitionCreatedIterator struct {
	Event *TaskDefinitionLibraryTaskDefinitionCreated // Event containing the contract specifics and raw log

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
func (it *TaskDefinitionLibraryTaskDefinitionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskDefinitionLibraryTaskDefinitionCreated)
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
		it.Event = new(TaskDefinitionLibraryTaskDefinitionCreated)
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
func (it *TaskDefinitionLibraryTaskDefinitionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskDefinitionLibraryTaskDefinitionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskDefinitionLibraryTaskDefinitionCreated represents a TaskDefinitionCreated event raised by the TaskDefinitionLibrary contract.
type TaskDefinitionLibraryTaskDefinitionCreated struct {
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

// FilterTaskDefinitionCreated is a free log retrieval operation binding the contract event 0x922bce103cd095cc3c1297fd9537ca88abcc2bdd3a03c9e0086fd9adff5389d2.
//
// Solidity: event TaskDefinitionCreated(uint8 taskDefinitionId, string name, uint256 baseRewardFeeForAttesters, uint256 baseRewardFeeForPerformer, uint256 baseRewardFeeForAggregator, uint256 minimumVotingPower, uint256[] restrictedOperatorIds, uint256 maximumNumberOfAttesters)
func (_TaskDefinitionLibrary *TaskDefinitionLibraryFilterer) FilterTaskDefinitionCreated(opts *bind.FilterOpts) (*TaskDefinitionLibraryTaskDefinitionCreatedIterator, error) {

	logs, sub, err := _TaskDefinitionLibrary.contract.FilterLogs(opts, "TaskDefinitionCreated")
	if err != nil {
		return nil, err
	}
	return &TaskDefinitionLibraryTaskDefinitionCreatedIterator{contract: _TaskDefinitionLibrary.contract, event: "TaskDefinitionCreated", logs: logs, sub: sub}, nil
}

// WatchTaskDefinitionCreated is a free log subscription operation binding the contract event 0x922bce103cd095cc3c1297fd9537ca88abcc2bdd3a03c9e0086fd9adff5389d2.
//
// Solidity: event TaskDefinitionCreated(uint8 taskDefinitionId, string name, uint256 baseRewardFeeForAttesters, uint256 baseRewardFeeForPerformer, uint256 baseRewardFeeForAggregator, uint256 minimumVotingPower, uint256[] restrictedOperatorIds, uint256 maximumNumberOfAttesters)
func (_TaskDefinitionLibrary *TaskDefinitionLibraryFilterer) WatchTaskDefinitionCreated(opts *bind.WatchOpts, sink chan<- *TaskDefinitionLibraryTaskDefinitionCreated) (event.Subscription, error) {

	logs, sub, err := _TaskDefinitionLibrary.contract.WatchLogs(opts, "TaskDefinitionCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskDefinitionLibraryTaskDefinitionCreated)
				if err := _TaskDefinitionLibrary.contract.UnpackLog(event, "TaskDefinitionCreated", log); err != nil {
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

// ParseTaskDefinitionCreated is a log parse operation binding the contract event 0x922bce103cd095cc3c1297fd9537ca88abcc2bdd3a03c9e0086fd9adff5389d2.
//
// Solidity: event TaskDefinitionCreated(uint8 taskDefinitionId, string name, uint256 baseRewardFeeForAttesters, uint256 baseRewardFeeForPerformer, uint256 baseRewardFeeForAggregator, uint256 minimumVotingPower, uint256[] restrictedOperatorIds, uint256 maximumNumberOfAttesters)
func (_TaskDefinitionLibrary *TaskDefinitionLibraryFilterer) ParseTaskDefinitionCreated(log types.Log) (*TaskDefinitionLibraryTaskDefinitionCreated, error) {
	event := new(TaskDefinitionLibraryTaskDefinitionCreated)
	if err := _TaskDefinitionLibrary.contract.UnpackLog(event, "TaskDefinitionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
