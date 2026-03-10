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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_taskExecutionHub\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execJobFromHub\",\"inputs\":[{\"name\":\"safeAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"actionTarget\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"actionValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actionData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operation\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"taskExecutionHub\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"TaskExecutedFromModule\",\"inputs\":[{\"name\":\"safeAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"executor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ExecFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotTaskExecutionHub\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
	Bin: "0x60a0346100aa57601f6103a338819003918201601f19168301916001600160401b038311848410176100ae578084926020946040528339810103126100aa57516001600160a01b0381168082036100aa5760015f551561007a576080526040516102e090816100c38239608051818181604701526101130152f35b60405162461bcd60e51b81526020600482015260086024820152673d32b93790343ab160c11b6044820152606490fd5b5f80fd5b634e487b7160e01b5f52604160045260245ffdfe6080806040526004361015610012575f80fd5b5f3560e01c908163b36e85181461007a575063c54d346e14610032575f80fd5b34610076575f366003190112610076576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b5f80fd5b346100765760a0366003190112610076576004356001600160a01b0381169190829003610076576024356001600160a01b03811690819003610076576064359067ffffffffffffffff821161007657366023830112156100765781600401359167ffffffffffffffff8311610076573660248483010111610076576084359160ff83168093036100765760025f541461029b5760025f557f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316330361028c57838593602095602460a495879563468721a760e01b87526004870152604435828701526080604487015282608487015201858501375f8483850101526064830152601f801991011681010301815f865af15f9181610225575b5061022057505f5b156101e45760405190600182527f528c00f3d7da8e9429355ebfd96f7db55a1b2363c6cf6eba8f256710c6f03c0760203293a360015f55602060405160018152f35b604051905f82527f528c00f3d7da8e9429355ebfd96f7db55a1b2363c6cf6eba8f256710c6f03c0760203293a3637349437160e11b5f5260045ffd5b6101a2565b90915060203d602011610285575b601f8101601f1916820167ffffffffffffffff811183821017610271576020918391604052810103126100765751801515810361007657908361019a565b634e487b7160e01b5f52604160045260245ffd5b503d610233565b633a6f757960e11b5f5260045ffd5b633ee5aeb560e01b5f5260045ffdfea2646970667358221220ed4bcb1042e978fc3e08b03e92693178262354f7c150057b4f13b7a7f8ae2a7b64736f6c634300081b0033",
}

// ContractTriggerXSafeModuleABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractTriggerXSafeModuleMetaData.ABI instead.
var ContractTriggerXSafeModuleABI = ContractTriggerXSafeModuleMetaData.ABI

// ContractTriggerXSafeModuleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractTriggerXSafeModuleMetaData.Bin instead.
var ContractTriggerXSafeModuleBin = ContractTriggerXSafeModuleMetaData.Bin

// DeployContractTriggerXSafeModule deploys a new Ethereum contract, binding an instance of ContractTriggerXSafeModule to it.
func DeployContractTriggerXSafeModule(auth *bind.TransactOpts, backend bind.ContractBackend, _taskExecutionHub common.Address) (common.Address, *types.Transaction, *ContractTriggerXSafeModule, error) {
	parsed, err := ContractTriggerXSafeModuleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractTriggerXSafeModuleBin), backend, _taskExecutionHub)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractTriggerXSafeModule{ContractTriggerXSafeModuleCaller: ContractTriggerXSafeModuleCaller{contract: contract}, ContractTriggerXSafeModuleTransactor: ContractTriggerXSafeModuleTransactor{contract: contract}, ContractTriggerXSafeModuleFilterer: ContractTriggerXSafeModuleFilterer{contract: contract}}, nil
}

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
