// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractVetoableSlasher

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

// IAllocationManagerTypesSlashingParams is an auto generated low-level Go binding around an user-defined struct.
type IAllocationManagerTypesSlashingParams struct {
	Operator      common.Address
	OperatorSetId uint32
	Strategies    []common.Address
	WadsToSlash   []*big.Int
	Description   string
}

// ContractVetoableSlasherMetaData contains all meta data concerning the ContractVetoableSlasher contract.
var ContractVetoableSlasherMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIAllocationManager\",\"name\":\"_allocationManager\",\"type\":\"address\"},{\"internalType\":\"contractIServiceManager\",\"name\":\"_serviceManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_slasher\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"OnlySlasher\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyVetoCommittee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SlashingRequestIsCancelled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SlashingRequestNotRequested\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VetoPeriodNotPassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VetoPeriodPassed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"slashingRequestId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"operatorSetId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"wadsToSlash\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"OperatorSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"SlashingRequestCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"operatorSetId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"wadsToSlash\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"SlashingRequested\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"VETO_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allocationManager\",\"outputs\":[{\"internalType\":\"contractIAllocationManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"cancelSlashingRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"fulfillSlashingRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vetoCommittee\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_slasher\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"operatorSetId\",\"type\":\"uint32\"},{\"internalType\":\"contractIStrategy[]\",\"name\":\"strategies\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"wadsToSlash\",\"type\":\"uint256[]\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structIAllocationManagerTypes.SlashingParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"queueSlashingRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"serviceManager\",\"outputs\":[{\"internalType\":\"contractIServiceManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slasher\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"slashingRequests\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"operatorSetId\",\"type\":\"uint32\"},{\"internalType\":\"contractIStrategy[]\",\"name\":\"strategies\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"wadsToSlash\",\"type\":\"uint256[]\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structIAllocationManagerTypes.SlashingParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"requestTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumISlasherTypes.SlashingStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vetoCommittee\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080806040526004361015610012575f80fd5b5f905f3560e01c90816304e17704146106ec575080630b6d4a201461051e5780631f05cc8e146104195780633998fdd3146103d4578063485cc955146101e95780636a84a985146101cb5780636a8dcf54146101a2578063b134427114610177578063ca8aa7c714610132578063e861dc85146101145763f48ab27f14610097575f80fd5b346101115760203660031901126101115760043581526033602052604081206100bf81610c53565b9060ff60056004830154920154166100e260405193606085526060850190610e26565b91602084015260048110156100fd5782935060408301520390f35b634e487b7160e01b84526021600452602484fd5b80fd5b503461011157806003193601126101115760206040516203f4808152f35b50346101115780600319360112610111576040517f00000000000000000000000078469728304326cbc65f8f95fa756b0b731644626001600160a01b03168152602090f35b50346101115780600319360112610111575460405160109190911c6001600160a01b03168152602090f35b50346101115780600319360112610111576032546040516001600160a01b039091168152602090f35b50346101115780600319360112610111576020600154604051908152f35b5034610111576040366003190112610111576004356001600160a01b038116908190036103d0576024356001600160a01b03811681036103cc57825460ff8160081c1615908180926103bf575b80156103a8575b1561034c5760ff19811660011785558161033b575b5083549160ff8360081c16156102e25762010000600160b01b0319831660109190911b62010000600160b01b03169081178555603280546001600160a01b0319169094179093556102a1578280f35b610100600160b01b031916178155604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602090a15f808280f35b60405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b6064820152608490fd5b61ffff19166101011784555f610252565b60405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608490fd5b50303b15801561023d5750600160ff82161461023d565b50600160ff821610610236565b8280fd5b5080fd5b50346101115780600319360112610111576040517f000000000000000000000000c0aa79c13bc1eea0168f1a2975b7aa9ff64135ed6001600160a01b03168152602090f35b503461011157602036600319011261011157603254600435906001600160a01b0316330361050f57808252603360205260046040832001546203f48081018091116104fb574210156104ec57808252603360205260ff60056040842001541660048110156104d8576001036104c957808252603360205260408220600501805460ff191660031790557fb612b6d36795da938211ef5caee2c1d887a24430f76918733089eee28b7f54ac8280a280f35b632c0bd3bb60e11b8252600482fd5b634e487b7160e01b83526021600452602483fd5b6308bf43f760e01b8252600482fd5b634e487b7160e01b83526011600452602483fd5b633ef0720160e21b8252600482fd5b50346106a25760203660031901126106a25760043561053c33610ebd565b805f52603360205260405f2060048101546203f48081018091116106d85742106106c957600581019060ff82541660048110156106b5576001036106a65761058e91600260ff19825416179055610c53565b907f00000000000000000000000078469728304326cbc65f8f95fa756b0b731644626001600160a01b0316803b156106a25760408051633635205760e01b81527f000000000000000000000000c0aa79c13bc1eea0168f1a2975b7aa9ff64135ed6001600160a01b031660048201526024810191909152905f90829081838161061a604482018a610e26565b03925af1801561069757610682575b5060018060a01b03825116907f8a83cf9afb09a981314f4fb353b95b003451da170a99f48d8db6474b06d79f3b63ffffffff60208501511693608060608201519101519061067c60405192839283610f1e565b0390a480f35b61068f9193505f90610be1565b5f915f610629565b6040513d5f823e3d90fd5b5f80fd5b63086c862560e01b5f5260045ffd5b634e487b7160e01b5f52602160045260245ffd5b637c80029d60e11b5f5260045ffd5b634e487b7160e01b5f52601160045260245ffd5b346106a25760203660031901126106a2576004359067ffffffffffffffff82116106a25760a060031983360301126106a25761072781610bc5565b60048201356001600160a01b03811681036106a257815260248201359063ffffffff821682036106a25760208101918252604483013567ffffffffffffffff81116106a2578301366023820112156106a257600481013561078781610c03565b916107956040519384610be1565b818352602060048185019360051b83010101903682116106a257602401915b818310610ba5575050506040820152606483013567ffffffffffffffff81116106a257830192366023850112156106a2576004840135936107f485610c03565b946108026040519687610be1565b808652602060048188019260051b84010101913683116106a257602401905b828210610b95575050506060820193845260848101359067ffffffffffffffff82116106a2570190366023830112156106a257600482013567ffffffffffffffff8111610b505760405192610880601f8301601f191660200185610be1565b81845236602482840101116106a257815f926024602093018387013784010152608081019182526108b033610ebd565b600154915f1983146106d85760018301600155604051916060830183811067ffffffffffffffff821117610b5057604090815281845242602080860191825260018684018181525f89815260338452859020975180518954948201516001600160c01b03199095166001600160a01b03919091161760a09490941b63ffffffff60a01b169390931788559382015180519497949185019067ffffffffffffffff8311610b5057600160401b8311610b505760209061097384845481865585610efb565b01905f5260205f205f5b838110610b7857505050506002830160608201519081519167ffffffffffffffff8311610b5057600160401b8311610b50576020906109c184845481865585610efb565b01905f5260205f205f5b838110610b645750505050608001518051600384019167ffffffffffffffff8211610b50576109fa8354610c1b565b601f8111610b15575b50602090601f8311600114610aad57600595949392915f9183610aa2575b50508160011b915f199060031b1c19161790555b5160048201550192519060048210156106b5577fadd285945f652e749df3dab9a584be524ec7fbd2a2cad39851278950f9b732279363ffffffff9260ff8019835416911617905560018060a01b039051169451169451905190610a9d60405192839283610f1e565b0390a4005b015190508c80610a21565b90601f19831691845f52815f20925f5b818110610afd57509160019391856005999897969410610ae5575b505050811b019055610a35565b01515f1960f88460031b161c191690558c8080610ad8565b92936020600181928786015181550195019301610abd565b610b4090845f5260205f20601f850160051c81019160208610610b46575b601f0160051c0190610ee5565b8b610a03565b9091508190610b33565b634e487b7160e01b5f52604160045260245ffd5b6001906020845194019381840155016109cb565b82516001600160a01b03168183015560209092019160010161097d565b8135815260209182019101610821565b82356001600160a01b03811681036106a2578152602092830192016107b4565b60a0810190811067ffffffffffffffff821117610b5057604052565b90601f8019910116810190811067ffffffffffffffff821117610b5057604052565b67ffffffffffffffff8111610b505760051b60200190565b90600182811c92168015610c49575b6020831014610c3557565b634e487b7160e01b5f52602260045260245ffd5b91607f1691610c2a565b90604051610c6081610bc5565b809263ffffffff815460018060a01b038116845260a01c16602083015260018101604051808260208294549384815201905f5260205f20925f5b818110610dad575050610caf92500382610be1565b6040830152600281016040519081602082549182815201915f5260205f20905f5b818110610d975750505090610cea81600394930382610be1565b6060840152019060405180925f90805490610d0482610c1b565b8085529160018116908115610d705750600114610d31575b505060809291610d2d910384610be1565b0152565b5f908152602081209092505b818310610d54575050810160200181610d2d610d1c565b6020919350806001915483858901015201910190918492610d3d565b60ff191660208681019190915292151560051b85019092019250839150610d2d9050610d1c565b8254845260209093019260019283019201610cd0565b84546001600160a01b0316835260019485019486945060209093019201610c9a565b90602080835192838152019201905f5b818110610dec5750505090565b8251845260209384019390920191600101610ddf565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b80516001600160a01b0316825260208082015163ffffffff168184015260408083015160a0918501829052805191850182905260c0850195949201905f5b818110610e9e575050506080610e8a610e9b949560608501518482036060860152610dcf565b920151906080818403910152610e02565b90565b82516001600160a01b0316875260209687019690920191600101610e64565b5f5460101c6001600160a01b03908116911603610ed657565b637e57b1e160e01b5f5260045ffd5b818110610ef0575050565b5f8155600101610ee5565b91818110610f0857505050565b610f1c925f5260205f209182019101610ee5565b565b9091610f35610e9b93604084526040840190610dcf565b916020818403910152610e0256fea26469706673582212201ffcd2461f1f2bff7cbddd7d81ca968228da240f90e0269639a2cf6d1d72f7e264736f6c634300081b0033",
}

// ContractVetoableSlasherABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractVetoableSlasherMetaData.ABI instead.
var ContractVetoableSlasherABI = ContractVetoableSlasherMetaData.ABI

// ContractVetoableSlasherBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractVetoableSlasherMetaData.Bin instead.
var ContractVetoableSlasherBin = ContractVetoableSlasherMetaData.Bin

// DeployContractVetoableSlasher deploys a new Ethereum contract, binding an instance of ContractVetoableSlasher to it.
func DeployContractVetoableSlasher(auth *bind.TransactOpts, backend bind.ContractBackend, _allocationManager common.Address, _serviceManager common.Address, _slasher common.Address) (common.Address, *types.Transaction, *ContractVetoableSlasher, error) {
	parsed, err := ContractVetoableSlasherMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractVetoableSlasherBin), backend, _allocationManager, _serviceManager, _slasher)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractVetoableSlasher{ContractVetoableSlasherCaller: ContractVetoableSlasherCaller{contract: contract}, ContractVetoableSlasherTransactor: ContractVetoableSlasherTransactor{contract: contract}, ContractVetoableSlasherFilterer: ContractVetoableSlasherFilterer{contract: contract}}, nil
}

// ContractVetoableSlasherMethods is an auto generated interface around an Ethereum contract.
type ContractVetoableSlasherMethods interface {
	ContractVetoableSlasherCalls
	ContractVetoableSlasherTransacts
	ContractVetoableSlasherFilters
}

// ContractVetoableSlasherCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractVetoableSlasherCalls interface {
	VETOPERIOD(opts *bind.CallOpts) (*big.Int, error)

	AllocationManager(opts *bind.CallOpts) (common.Address, error)

	NextRequestId(opts *bind.CallOpts) (*big.Int, error)

	ServiceManager(opts *bind.CallOpts) (common.Address, error)

	Slasher(opts *bind.CallOpts) (common.Address, error)

	SlashingRequests(opts *bind.CallOpts, arg0 *big.Int) (struct {
		Params           IAllocationManagerTypesSlashingParams
		RequestTimestamp *big.Int
		Status           uint8
	}, error)

	VetoCommittee(opts *bind.CallOpts) (common.Address, error)
}

// ContractVetoableSlasherTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractVetoableSlasherTransacts interface {
	CancelSlashingRequest(opts *bind.TransactOpts, requestId *big.Int) (*types.Transaction, error)

	FulfillSlashingRequest(opts *bind.TransactOpts, requestId *big.Int) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts, _vetoCommittee common.Address, _slasher common.Address) (*types.Transaction, error)

	QueueSlashingRequest(opts *bind.TransactOpts, params IAllocationManagerTypesSlashingParams) (*types.Transaction, error)
}

// ContractVetoableSlasherFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractVetoableSlasherFilters interface {
	FilterInitialized(opts *bind.FilterOpts) (*ContractVetoableSlasherInitializedIterator, error)
	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractVetoableSlasherInitialized) (event.Subscription, error)
	ParseInitialized(log types.Log) (*ContractVetoableSlasherInitialized, error)

	FilterOperatorSlashed(opts *bind.FilterOpts, slashingRequestId []*big.Int, operator []common.Address, operatorSetId []uint32) (*ContractVetoableSlasherOperatorSlashedIterator, error)
	WatchOperatorSlashed(opts *bind.WatchOpts, sink chan<- *ContractVetoableSlasherOperatorSlashed, slashingRequestId []*big.Int, operator []common.Address, operatorSetId []uint32) (event.Subscription, error)
	ParseOperatorSlashed(log types.Log) (*ContractVetoableSlasherOperatorSlashed, error)

	FilterSlashingRequestCancelled(opts *bind.FilterOpts, requestId []*big.Int) (*ContractVetoableSlasherSlashingRequestCancelledIterator, error)
	WatchSlashingRequestCancelled(opts *bind.WatchOpts, sink chan<- *ContractVetoableSlasherSlashingRequestCancelled, requestId []*big.Int) (event.Subscription, error)
	ParseSlashingRequestCancelled(log types.Log) (*ContractVetoableSlasherSlashingRequestCancelled, error)

	FilterSlashingRequested(opts *bind.FilterOpts, requestId []*big.Int, operator []common.Address, operatorSetId []uint32) (*ContractVetoableSlasherSlashingRequestedIterator, error)
	WatchSlashingRequested(opts *bind.WatchOpts, sink chan<- *ContractVetoableSlasherSlashingRequested, requestId []*big.Int, operator []common.Address, operatorSetId []uint32) (event.Subscription, error)
	ParseSlashingRequested(log types.Log) (*ContractVetoableSlasherSlashingRequested, error)
}

// ContractVetoableSlasher is an auto generated Go binding around an Ethereum contract.
type ContractVetoableSlasher struct {
	ContractVetoableSlasherCaller     // Read-only binding to the contract
	ContractVetoableSlasherTransactor // Write-only binding to the contract
	ContractVetoableSlasherFilterer   // Log filterer for contract events
}

// ContractVetoableSlasher implements the ContractVetoableSlasherMethods interface.
var _ ContractVetoableSlasherMethods = (*ContractVetoableSlasher)(nil)

// ContractVetoableSlasherCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractVetoableSlasherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractVetoableSlasherCaller implements the ContractVetoableSlasherCalls interface.
var _ ContractVetoableSlasherCalls = (*ContractVetoableSlasherCaller)(nil)

// ContractVetoableSlasherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractVetoableSlasherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractVetoableSlasherTransactor implements the ContractVetoableSlasherTransacts interface.
var _ ContractVetoableSlasherTransacts = (*ContractVetoableSlasherTransactor)(nil)

// ContractVetoableSlasherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractVetoableSlasherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractVetoableSlasherFilterer implements the ContractVetoableSlasherFilters interface.
var _ ContractVetoableSlasherFilters = (*ContractVetoableSlasherFilterer)(nil)

// ContractVetoableSlasherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractVetoableSlasherSession struct {
	Contract     *ContractVetoableSlasher // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ContractVetoableSlasherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractVetoableSlasherCallerSession struct {
	Contract *ContractVetoableSlasherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// ContractVetoableSlasherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractVetoableSlasherTransactorSession struct {
	Contract     *ContractVetoableSlasherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// ContractVetoableSlasherRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractVetoableSlasherRaw struct {
	Contract *ContractVetoableSlasher // Generic contract binding to access the raw methods on
}

// ContractVetoableSlasherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractVetoableSlasherCallerRaw struct {
	Contract *ContractVetoableSlasherCaller // Generic read-only contract binding to access the raw methods on
}

// ContractVetoableSlasherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractVetoableSlasherTransactorRaw struct {
	Contract *ContractVetoableSlasherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractVetoableSlasher creates a new instance of ContractVetoableSlasher, bound to a specific deployed contract.
func NewContractVetoableSlasher(address common.Address, backend bind.ContractBackend) (*ContractVetoableSlasher, error) {
	contract, err := bindContractVetoableSlasher(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractVetoableSlasher{ContractVetoableSlasherCaller: ContractVetoableSlasherCaller{contract: contract}, ContractVetoableSlasherTransactor: ContractVetoableSlasherTransactor{contract: contract}, ContractVetoableSlasherFilterer: ContractVetoableSlasherFilterer{contract: contract}}, nil
}

// NewContractVetoableSlasherCaller creates a new read-only instance of ContractVetoableSlasher, bound to a specific deployed contract.
func NewContractVetoableSlasherCaller(address common.Address, caller bind.ContractCaller) (*ContractVetoableSlasherCaller, error) {
	contract, err := bindContractVetoableSlasher(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractVetoableSlasherCaller{contract: contract}, nil
}

// NewContractVetoableSlasherTransactor creates a new write-only instance of ContractVetoableSlasher, bound to a specific deployed contract.
func NewContractVetoableSlasherTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractVetoableSlasherTransactor, error) {
	contract, err := bindContractVetoableSlasher(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractVetoableSlasherTransactor{contract: contract}, nil
}

// NewContractVetoableSlasherFilterer creates a new log filterer instance of ContractVetoableSlasher, bound to a specific deployed contract.
func NewContractVetoableSlasherFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractVetoableSlasherFilterer, error) {
	contract, err := bindContractVetoableSlasher(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractVetoableSlasherFilterer{contract: contract}, nil
}

// bindContractVetoableSlasher binds a generic wrapper to an already deployed contract.
func bindContractVetoableSlasher(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractVetoableSlasherMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractVetoableSlasher *ContractVetoableSlasherRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractVetoableSlasher.Contract.ContractVetoableSlasherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractVetoableSlasher *ContractVetoableSlasherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractVetoableSlasher.Contract.ContractVetoableSlasherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractVetoableSlasher *ContractVetoableSlasherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractVetoableSlasher.Contract.ContractVetoableSlasherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractVetoableSlasher *ContractVetoableSlasherCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractVetoableSlasher.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractVetoableSlasher *ContractVetoableSlasherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractVetoableSlasher.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractVetoableSlasher *ContractVetoableSlasherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractVetoableSlasher.Contract.contract.Transact(opts, method, params...)
}

// VETOPERIOD is a free data retrieval call binding the contract method 0xe861dc85.
//
// Solidity: function VETO_PERIOD() view returns(uint256)
func (_ContractVetoableSlasher *ContractVetoableSlasherCaller) VETOPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractVetoableSlasher.contract.Call(opts, &out, "VETO_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VETOPERIOD is a free data retrieval call binding the contract method 0xe861dc85.
//
// Solidity: function VETO_PERIOD() view returns(uint256)
func (_ContractVetoableSlasher *ContractVetoableSlasherSession) VETOPERIOD() (*big.Int, error) {
	return _ContractVetoableSlasher.Contract.VETOPERIOD(&_ContractVetoableSlasher.CallOpts)
}

// VETOPERIOD is a free data retrieval call binding the contract method 0xe861dc85.
//
// Solidity: function VETO_PERIOD() view returns(uint256)
func (_ContractVetoableSlasher *ContractVetoableSlasherCallerSession) VETOPERIOD() (*big.Int, error) {
	return _ContractVetoableSlasher.Contract.VETOPERIOD(&_ContractVetoableSlasher.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_ContractVetoableSlasher *ContractVetoableSlasherCaller) AllocationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractVetoableSlasher.contract.Call(opts, &out, "allocationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_ContractVetoableSlasher *ContractVetoableSlasherSession) AllocationManager() (common.Address, error) {
	return _ContractVetoableSlasher.Contract.AllocationManager(&_ContractVetoableSlasher.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_ContractVetoableSlasher *ContractVetoableSlasherCallerSession) AllocationManager() (common.Address, error) {
	return _ContractVetoableSlasher.Contract.AllocationManager(&_ContractVetoableSlasher.CallOpts)
}

// NextRequestId is a free data retrieval call binding the contract method 0x6a84a985.
//
// Solidity: function nextRequestId() view returns(uint256)
func (_ContractVetoableSlasher *ContractVetoableSlasherCaller) NextRequestId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractVetoableSlasher.contract.Call(opts, &out, "nextRequestId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextRequestId is a free data retrieval call binding the contract method 0x6a84a985.
//
// Solidity: function nextRequestId() view returns(uint256)
func (_ContractVetoableSlasher *ContractVetoableSlasherSession) NextRequestId() (*big.Int, error) {
	return _ContractVetoableSlasher.Contract.NextRequestId(&_ContractVetoableSlasher.CallOpts)
}

// NextRequestId is a free data retrieval call binding the contract method 0x6a84a985.
//
// Solidity: function nextRequestId() view returns(uint256)
func (_ContractVetoableSlasher *ContractVetoableSlasherCallerSession) NextRequestId() (*big.Int, error) {
	return _ContractVetoableSlasher.Contract.NextRequestId(&_ContractVetoableSlasher.CallOpts)
}

// ServiceManager is a free data retrieval call binding the contract method 0x3998fdd3.
//
// Solidity: function serviceManager() view returns(address)
func (_ContractVetoableSlasher *ContractVetoableSlasherCaller) ServiceManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractVetoableSlasher.contract.Call(opts, &out, "serviceManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ServiceManager is a free data retrieval call binding the contract method 0x3998fdd3.
//
// Solidity: function serviceManager() view returns(address)
func (_ContractVetoableSlasher *ContractVetoableSlasherSession) ServiceManager() (common.Address, error) {
	return _ContractVetoableSlasher.Contract.ServiceManager(&_ContractVetoableSlasher.CallOpts)
}

// ServiceManager is a free data retrieval call binding the contract method 0x3998fdd3.
//
// Solidity: function serviceManager() view returns(address)
func (_ContractVetoableSlasher *ContractVetoableSlasherCallerSession) ServiceManager() (common.Address, error) {
	return _ContractVetoableSlasher.Contract.ServiceManager(&_ContractVetoableSlasher.CallOpts)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_ContractVetoableSlasher *ContractVetoableSlasherCaller) Slasher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractVetoableSlasher.contract.Call(opts, &out, "slasher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_ContractVetoableSlasher *ContractVetoableSlasherSession) Slasher() (common.Address, error) {
	return _ContractVetoableSlasher.Contract.Slasher(&_ContractVetoableSlasher.CallOpts)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_ContractVetoableSlasher *ContractVetoableSlasherCallerSession) Slasher() (common.Address, error) {
	return _ContractVetoableSlasher.Contract.Slasher(&_ContractVetoableSlasher.CallOpts)
}

// SlashingRequests is a free data retrieval call binding the contract method 0xf48ab27f.
//
// Solidity: function slashingRequests(uint256 ) view returns((address,uint32,address[],uint256[],string) params, uint256 requestTimestamp, uint8 status)
func (_ContractVetoableSlasher *ContractVetoableSlasherCaller) SlashingRequests(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Params           IAllocationManagerTypesSlashingParams
	RequestTimestamp *big.Int
	Status           uint8
}, error) {
	var out []interface{}
	err := _ContractVetoableSlasher.contract.Call(opts, &out, "slashingRequests", arg0)

	outstruct := new(struct {
		Params           IAllocationManagerTypesSlashingParams
		RequestTimestamp *big.Int
		Status           uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Params = *abi.ConvertType(out[0], new(IAllocationManagerTypesSlashingParams)).(*IAllocationManagerTypesSlashingParams)
	outstruct.RequestTimestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// SlashingRequests is a free data retrieval call binding the contract method 0xf48ab27f.
//
// Solidity: function slashingRequests(uint256 ) view returns((address,uint32,address[],uint256[],string) params, uint256 requestTimestamp, uint8 status)
func (_ContractVetoableSlasher *ContractVetoableSlasherSession) SlashingRequests(arg0 *big.Int) (struct {
	Params           IAllocationManagerTypesSlashingParams
	RequestTimestamp *big.Int
	Status           uint8
}, error) {
	return _ContractVetoableSlasher.Contract.SlashingRequests(&_ContractVetoableSlasher.CallOpts, arg0)
}

// SlashingRequests is a free data retrieval call binding the contract method 0xf48ab27f.
//
// Solidity: function slashingRequests(uint256 ) view returns((address,uint32,address[],uint256[],string) params, uint256 requestTimestamp, uint8 status)
func (_ContractVetoableSlasher *ContractVetoableSlasherCallerSession) SlashingRequests(arg0 *big.Int) (struct {
	Params           IAllocationManagerTypesSlashingParams
	RequestTimestamp *big.Int
	Status           uint8
}, error) {
	return _ContractVetoableSlasher.Contract.SlashingRequests(&_ContractVetoableSlasher.CallOpts, arg0)
}

// VetoCommittee is a free data retrieval call binding the contract method 0x6a8dcf54.
//
// Solidity: function vetoCommittee() view returns(address)
func (_ContractVetoableSlasher *ContractVetoableSlasherCaller) VetoCommittee(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractVetoableSlasher.contract.Call(opts, &out, "vetoCommittee")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VetoCommittee is a free data retrieval call binding the contract method 0x6a8dcf54.
//
// Solidity: function vetoCommittee() view returns(address)
func (_ContractVetoableSlasher *ContractVetoableSlasherSession) VetoCommittee() (common.Address, error) {
	return _ContractVetoableSlasher.Contract.VetoCommittee(&_ContractVetoableSlasher.CallOpts)
}

// VetoCommittee is a free data retrieval call binding the contract method 0x6a8dcf54.
//
// Solidity: function vetoCommittee() view returns(address)
func (_ContractVetoableSlasher *ContractVetoableSlasherCallerSession) VetoCommittee() (common.Address, error) {
	return _ContractVetoableSlasher.Contract.VetoCommittee(&_ContractVetoableSlasher.CallOpts)
}

// CancelSlashingRequest is a paid mutator transaction binding the contract method 0x1f05cc8e.
//
// Solidity: function cancelSlashingRequest(uint256 requestId) returns()
func (_ContractVetoableSlasher *ContractVetoableSlasherTransactor) CancelSlashingRequest(opts *bind.TransactOpts, requestId *big.Int) (*types.Transaction, error) {
	return _ContractVetoableSlasher.contract.Transact(opts, "cancelSlashingRequest", requestId)
}

// CancelSlashingRequest is a paid mutator transaction binding the contract method 0x1f05cc8e.
//
// Solidity: function cancelSlashingRequest(uint256 requestId) returns()
func (_ContractVetoableSlasher *ContractVetoableSlasherSession) CancelSlashingRequest(requestId *big.Int) (*types.Transaction, error) {
	return _ContractVetoableSlasher.Contract.CancelSlashingRequest(&_ContractVetoableSlasher.TransactOpts, requestId)
}

// CancelSlashingRequest is a paid mutator transaction binding the contract method 0x1f05cc8e.
//
// Solidity: function cancelSlashingRequest(uint256 requestId) returns()
func (_ContractVetoableSlasher *ContractVetoableSlasherTransactorSession) CancelSlashingRequest(requestId *big.Int) (*types.Transaction, error) {
	return _ContractVetoableSlasher.Contract.CancelSlashingRequest(&_ContractVetoableSlasher.TransactOpts, requestId)
}

// FulfillSlashingRequest is a paid mutator transaction binding the contract method 0x0b6d4a20.
//
// Solidity: function fulfillSlashingRequest(uint256 requestId) returns()
func (_ContractVetoableSlasher *ContractVetoableSlasherTransactor) FulfillSlashingRequest(opts *bind.TransactOpts, requestId *big.Int) (*types.Transaction, error) {
	return _ContractVetoableSlasher.contract.Transact(opts, "fulfillSlashingRequest", requestId)
}

// FulfillSlashingRequest is a paid mutator transaction binding the contract method 0x0b6d4a20.
//
// Solidity: function fulfillSlashingRequest(uint256 requestId) returns()
func (_ContractVetoableSlasher *ContractVetoableSlasherSession) FulfillSlashingRequest(requestId *big.Int) (*types.Transaction, error) {
	return _ContractVetoableSlasher.Contract.FulfillSlashingRequest(&_ContractVetoableSlasher.TransactOpts, requestId)
}

// FulfillSlashingRequest is a paid mutator transaction binding the contract method 0x0b6d4a20.
//
// Solidity: function fulfillSlashingRequest(uint256 requestId) returns()
func (_ContractVetoableSlasher *ContractVetoableSlasherTransactorSession) FulfillSlashingRequest(requestId *big.Int) (*types.Transaction, error) {
	return _ContractVetoableSlasher.Contract.FulfillSlashingRequest(&_ContractVetoableSlasher.TransactOpts, requestId)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _vetoCommittee, address _slasher) returns()
func (_ContractVetoableSlasher *ContractVetoableSlasherTransactor) Initialize(opts *bind.TransactOpts, _vetoCommittee common.Address, _slasher common.Address) (*types.Transaction, error) {
	return _ContractVetoableSlasher.contract.Transact(opts, "initialize", _vetoCommittee, _slasher)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _vetoCommittee, address _slasher) returns()
func (_ContractVetoableSlasher *ContractVetoableSlasherSession) Initialize(_vetoCommittee common.Address, _slasher common.Address) (*types.Transaction, error) {
	return _ContractVetoableSlasher.Contract.Initialize(&_ContractVetoableSlasher.TransactOpts, _vetoCommittee, _slasher)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _vetoCommittee, address _slasher) returns()
func (_ContractVetoableSlasher *ContractVetoableSlasherTransactorSession) Initialize(_vetoCommittee common.Address, _slasher common.Address) (*types.Transaction, error) {
	return _ContractVetoableSlasher.Contract.Initialize(&_ContractVetoableSlasher.TransactOpts, _vetoCommittee, _slasher)
}

// QueueSlashingRequest is a paid mutator transaction binding the contract method 0x04e17704.
//
// Solidity: function queueSlashingRequest((address,uint32,address[],uint256[],string) params) returns()
func (_ContractVetoableSlasher *ContractVetoableSlasherTransactor) QueueSlashingRequest(opts *bind.TransactOpts, params IAllocationManagerTypesSlashingParams) (*types.Transaction, error) {
	return _ContractVetoableSlasher.contract.Transact(opts, "queueSlashingRequest", params)
}

// QueueSlashingRequest is a paid mutator transaction binding the contract method 0x04e17704.
//
// Solidity: function queueSlashingRequest((address,uint32,address[],uint256[],string) params) returns()
func (_ContractVetoableSlasher *ContractVetoableSlasherSession) QueueSlashingRequest(params IAllocationManagerTypesSlashingParams) (*types.Transaction, error) {
	return _ContractVetoableSlasher.Contract.QueueSlashingRequest(&_ContractVetoableSlasher.TransactOpts, params)
}

// QueueSlashingRequest is a paid mutator transaction binding the contract method 0x04e17704.
//
// Solidity: function queueSlashingRequest((address,uint32,address[],uint256[],string) params) returns()
func (_ContractVetoableSlasher *ContractVetoableSlasherTransactorSession) QueueSlashingRequest(params IAllocationManagerTypesSlashingParams) (*types.Transaction, error) {
	return _ContractVetoableSlasher.Contract.QueueSlashingRequest(&_ContractVetoableSlasher.TransactOpts, params)
}

// ContractVetoableSlasherInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractVetoableSlasher contract.
type ContractVetoableSlasherInitializedIterator struct {
	Event *ContractVetoableSlasherInitialized // Event containing the contract specifics and raw log

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
func (it *ContractVetoableSlasherInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractVetoableSlasherInitialized)
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
		it.Event = new(ContractVetoableSlasherInitialized)
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
func (it *ContractVetoableSlasherInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractVetoableSlasherInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractVetoableSlasherInitialized represents a Initialized event raised by the ContractVetoableSlasher contract.
type ContractVetoableSlasherInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractVetoableSlasher *ContractVetoableSlasherFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractVetoableSlasherInitializedIterator, error) {

	logs, sub, err := _ContractVetoableSlasher.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractVetoableSlasherInitializedIterator{contract: _ContractVetoableSlasher.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractVetoableSlasher *ContractVetoableSlasherFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractVetoableSlasherInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractVetoableSlasher.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractVetoableSlasherInitialized)
				if err := _ContractVetoableSlasher.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractVetoableSlasher *ContractVetoableSlasherFilterer) ParseInitialized(log types.Log) (*ContractVetoableSlasherInitialized, error) {
	event := new(ContractVetoableSlasherInitialized)
	if err := _ContractVetoableSlasher.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractVetoableSlasherOperatorSlashedIterator is returned from FilterOperatorSlashed and is used to iterate over the raw logs and unpacked data for OperatorSlashed events raised by the ContractVetoableSlasher contract.
type ContractVetoableSlasherOperatorSlashedIterator struct {
	Event *ContractVetoableSlasherOperatorSlashed // Event containing the contract specifics and raw log

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
func (it *ContractVetoableSlasherOperatorSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractVetoableSlasherOperatorSlashed)
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
		it.Event = new(ContractVetoableSlasherOperatorSlashed)
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
func (it *ContractVetoableSlasherOperatorSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractVetoableSlasherOperatorSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractVetoableSlasherOperatorSlashed represents a OperatorSlashed event raised by the ContractVetoableSlasher contract.
type ContractVetoableSlasherOperatorSlashed struct {
	SlashingRequestId *big.Int
	Operator          common.Address
	OperatorSetId     uint32
	WadsToSlash       []*big.Int
	Description       string
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterOperatorSlashed is a free log retrieval operation binding the contract event 0x8a83cf9afb09a981314f4fb353b95b003451da170a99f48d8db6474b06d79f3b.
//
// Solidity: event OperatorSlashed(uint256 indexed slashingRequestId, address indexed operator, uint32 indexed operatorSetId, uint256[] wadsToSlash, string description)
func (_ContractVetoableSlasher *ContractVetoableSlasherFilterer) FilterOperatorSlashed(opts *bind.FilterOpts, slashingRequestId []*big.Int, operator []common.Address, operatorSetId []uint32) (*ContractVetoableSlasherOperatorSlashedIterator, error) {

	var slashingRequestIdRule []interface{}
	for _, slashingRequestIdItem := range slashingRequestId {
		slashingRequestIdRule = append(slashingRequestIdRule, slashingRequestIdItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorSetIdRule []interface{}
	for _, operatorSetIdItem := range operatorSetId {
		operatorSetIdRule = append(operatorSetIdRule, operatorSetIdItem)
	}

	logs, sub, err := _ContractVetoableSlasher.contract.FilterLogs(opts, "OperatorSlashed", slashingRequestIdRule, operatorRule, operatorSetIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractVetoableSlasherOperatorSlashedIterator{contract: _ContractVetoableSlasher.contract, event: "OperatorSlashed", logs: logs, sub: sub}, nil
}

// WatchOperatorSlashed is a free log subscription operation binding the contract event 0x8a83cf9afb09a981314f4fb353b95b003451da170a99f48d8db6474b06d79f3b.
//
// Solidity: event OperatorSlashed(uint256 indexed slashingRequestId, address indexed operator, uint32 indexed operatorSetId, uint256[] wadsToSlash, string description)
func (_ContractVetoableSlasher *ContractVetoableSlasherFilterer) WatchOperatorSlashed(opts *bind.WatchOpts, sink chan<- *ContractVetoableSlasherOperatorSlashed, slashingRequestId []*big.Int, operator []common.Address, operatorSetId []uint32) (event.Subscription, error) {

	var slashingRequestIdRule []interface{}
	for _, slashingRequestIdItem := range slashingRequestId {
		slashingRequestIdRule = append(slashingRequestIdRule, slashingRequestIdItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorSetIdRule []interface{}
	for _, operatorSetIdItem := range operatorSetId {
		operatorSetIdRule = append(operatorSetIdRule, operatorSetIdItem)
	}

	logs, sub, err := _ContractVetoableSlasher.contract.WatchLogs(opts, "OperatorSlashed", slashingRequestIdRule, operatorRule, operatorSetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractVetoableSlasherOperatorSlashed)
				if err := _ContractVetoableSlasher.contract.UnpackLog(event, "OperatorSlashed", log); err != nil {
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

// ParseOperatorSlashed is a log parse operation binding the contract event 0x8a83cf9afb09a981314f4fb353b95b003451da170a99f48d8db6474b06d79f3b.
//
// Solidity: event OperatorSlashed(uint256 indexed slashingRequestId, address indexed operator, uint32 indexed operatorSetId, uint256[] wadsToSlash, string description)
func (_ContractVetoableSlasher *ContractVetoableSlasherFilterer) ParseOperatorSlashed(log types.Log) (*ContractVetoableSlasherOperatorSlashed, error) {
	event := new(ContractVetoableSlasherOperatorSlashed)
	if err := _ContractVetoableSlasher.contract.UnpackLog(event, "OperatorSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractVetoableSlasherSlashingRequestCancelledIterator is returned from FilterSlashingRequestCancelled and is used to iterate over the raw logs and unpacked data for SlashingRequestCancelled events raised by the ContractVetoableSlasher contract.
type ContractVetoableSlasherSlashingRequestCancelledIterator struct {
	Event *ContractVetoableSlasherSlashingRequestCancelled // Event containing the contract specifics and raw log

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
func (it *ContractVetoableSlasherSlashingRequestCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractVetoableSlasherSlashingRequestCancelled)
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
		it.Event = new(ContractVetoableSlasherSlashingRequestCancelled)
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
func (it *ContractVetoableSlasherSlashingRequestCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractVetoableSlasherSlashingRequestCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractVetoableSlasherSlashingRequestCancelled represents a SlashingRequestCancelled event raised by the ContractVetoableSlasher contract.
type ContractVetoableSlasherSlashingRequestCancelled struct {
	RequestId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSlashingRequestCancelled is a free log retrieval operation binding the contract event 0xb612b6d36795da938211ef5caee2c1d887a24430f76918733089eee28b7f54ac.
//
// Solidity: event SlashingRequestCancelled(uint256 indexed requestId)
func (_ContractVetoableSlasher *ContractVetoableSlasherFilterer) FilterSlashingRequestCancelled(opts *bind.FilterOpts, requestId []*big.Int) (*ContractVetoableSlasherSlashingRequestCancelledIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _ContractVetoableSlasher.contract.FilterLogs(opts, "SlashingRequestCancelled", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractVetoableSlasherSlashingRequestCancelledIterator{contract: _ContractVetoableSlasher.contract, event: "SlashingRequestCancelled", logs: logs, sub: sub}, nil
}

// WatchSlashingRequestCancelled is a free log subscription operation binding the contract event 0xb612b6d36795da938211ef5caee2c1d887a24430f76918733089eee28b7f54ac.
//
// Solidity: event SlashingRequestCancelled(uint256 indexed requestId)
func (_ContractVetoableSlasher *ContractVetoableSlasherFilterer) WatchSlashingRequestCancelled(opts *bind.WatchOpts, sink chan<- *ContractVetoableSlasherSlashingRequestCancelled, requestId []*big.Int) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _ContractVetoableSlasher.contract.WatchLogs(opts, "SlashingRequestCancelled", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractVetoableSlasherSlashingRequestCancelled)
				if err := _ContractVetoableSlasher.contract.UnpackLog(event, "SlashingRequestCancelled", log); err != nil {
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

// ParseSlashingRequestCancelled is a log parse operation binding the contract event 0xb612b6d36795da938211ef5caee2c1d887a24430f76918733089eee28b7f54ac.
//
// Solidity: event SlashingRequestCancelled(uint256 indexed requestId)
func (_ContractVetoableSlasher *ContractVetoableSlasherFilterer) ParseSlashingRequestCancelled(log types.Log) (*ContractVetoableSlasherSlashingRequestCancelled, error) {
	event := new(ContractVetoableSlasherSlashingRequestCancelled)
	if err := _ContractVetoableSlasher.contract.UnpackLog(event, "SlashingRequestCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractVetoableSlasherSlashingRequestedIterator is returned from FilterSlashingRequested and is used to iterate over the raw logs and unpacked data for SlashingRequested events raised by the ContractVetoableSlasher contract.
type ContractVetoableSlasherSlashingRequestedIterator struct {
	Event *ContractVetoableSlasherSlashingRequested // Event containing the contract specifics and raw log

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
func (it *ContractVetoableSlasherSlashingRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractVetoableSlasherSlashingRequested)
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
		it.Event = new(ContractVetoableSlasherSlashingRequested)
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
func (it *ContractVetoableSlasherSlashingRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractVetoableSlasherSlashingRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractVetoableSlasherSlashingRequested represents a SlashingRequested event raised by the ContractVetoableSlasher contract.
type ContractVetoableSlasherSlashingRequested struct {
	RequestId     *big.Int
	Operator      common.Address
	OperatorSetId uint32
	WadsToSlash   []*big.Int
	Description   string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSlashingRequested is a free log retrieval operation binding the contract event 0xadd285945f652e749df3dab9a584be524ec7fbd2a2cad39851278950f9b73227.
//
// Solidity: event SlashingRequested(uint256 indexed requestId, address indexed operator, uint32 indexed operatorSetId, uint256[] wadsToSlash, string description)
func (_ContractVetoableSlasher *ContractVetoableSlasherFilterer) FilterSlashingRequested(opts *bind.FilterOpts, requestId []*big.Int, operator []common.Address, operatorSetId []uint32) (*ContractVetoableSlasherSlashingRequestedIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorSetIdRule []interface{}
	for _, operatorSetIdItem := range operatorSetId {
		operatorSetIdRule = append(operatorSetIdRule, operatorSetIdItem)
	}

	logs, sub, err := _ContractVetoableSlasher.contract.FilterLogs(opts, "SlashingRequested", requestIdRule, operatorRule, operatorSetIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractVetoableSlasherSlashingRequestedIterator{contract: _ContractVetoableSlasher.contract, event: "SlashingRequested", logs: logs, sub: sub}, nil
}

// WatchSlashingRequested is a free log subscription operation binding the contract event 0xadd285945f652e749df3dab9a584be524ec7fbd2a2cad39851278950f9b73227.
//
// Solidity: event SlashingRequested(uint256 indexed requestId, address indexed operator, uint32 indexed operatorSetId, uint256[] wadsToSlash, string description)
func (_ContractVetoableSlasher *ContractVetoableSlasherFilterer) WatchSlashingRequested(opts *bind.WatchOpts, sink chan<- *ContractVetoableSlasherSlashingRequested, requestId []*big.Int, operator []common.Address, operatorSetId []uint32) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorSetIdRule []interface{}
	for _, operatorSetIdItem := range operatorSetId {
		operatorSetIdRule = append(operatorSetIdRule, operatorSetIdItem)
	}

	logs, sub, err := _ContractVetoableSlasher.contract.WatchLogs(opts, "SlashingRequested", requestIdRule, operatorRule, operatorSetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractVetoableSlasherSlashingRequested)
				if err := _ContractVetoableSlasher.contract.UnpackLog(event, "SlashingRequested", log); err != nil {
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

// ParseSlashingRequested is a log parse operation binding the contract event 0xadd285945f652e749df3dab9a584be524ec7fbd2a2cad39851278950f9b73227.
//
// Solidity: event SlashingRequested(uint256 indexed requestId, address indexed operator, uint32 indexed operatorSetId, uint256[] wadsToSlash, string description)
func (_ContractVetoableSlasher *ContractVetoableSlasherFilterer) ParseSlashingRequested(log types.Log) (*ContractVetoableSlasherSlashingRequested, error) {
	event := new(ContractVetoableSlasherSlashingRequested)
	if err := _ContractVetoableSlasher.contract.UnpackLog(event, "SlashingRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
