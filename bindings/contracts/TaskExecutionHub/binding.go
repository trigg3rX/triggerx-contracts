// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractTaskExecutionHub

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

// EnforcedOptionParam is an auto generated low-level Go binding around an user-defined struct.
type EnforcedOptionParam struct {
	Eid     uint32
	MsgType uint16
	Options []byte
}

// MessagingFee is an auto generated low-level Go binding around an user-defined struct.
type MessagingFee struct {
	NativeFee  *big.Int
	LzTokenFee *big.Int
}

// Origin is an auto generated low-level Go binding around an user-defined struct.
type Origin struct {
	SrcEid uint32
	Sender [32]byte
	Nonce  uint64
}

// ContractTaskExecutionHubMetaData contains all meta data concerning the ContractTaskExecutionHub contract.
var ContractTaskExecutionHubMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_endpoint\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_delegate\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDelegate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEndpointCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"name\":\"InvalidOptions\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LzTokenUnavailable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"}],\"name\":\"NoPeer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"}],\"name\":\"NotEnoughNative\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"OnlyEndpoint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"}],\"name\":\"OnlyPeer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumTaskExecutionHub.ActionType\",\"name\":\"action\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"}],\"name\":\"BroadcastSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"msgType\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structEnforcedOptionParam[]\",\"name\":\"_enforcedOptions\",\"type\":\"tuple[]\"}],\"name\":\"EnforcedOptionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"FeeUsed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"FunctionExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"gas\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"}],\"name\":\"GasConfigUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"}],\"name\":\"KeeperRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"}],\"name\":\"KeeperUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"LowBalanceAlert\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"MessageFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"peer\",\"type\":\"bytes32\"}],\"name\":\"PeerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"name\":\"_quoteFee\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lzTokenFee\",\"type\":\"uint256\"}],\"internalType\":\"structMessagingFee\",\"name\":\"fee\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"}],\"name\":\"addKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32[]\",\"name\":\"_dstEids\",\"type\":\"uint32[]\"}],\"name\":\"addSpokes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"origin\",\"type\":\"tuple\"}],\"name\":\"allowInitializePath\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"_msgType\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_extraOptions\",\"type\":\"bytes\"}],\"name\":\"combineOptions\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultGas\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultValue\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"dstEids\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endpoint\",\"outputs\":[{\"internalType\":\"contractILayerZeroEndpointV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"msgType\",\"type\":\"uint16\"}],\"name\":\"enforcedOptions\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"enforcedOption\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tgAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeFunction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ownerAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_srcEid\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_originEid\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"_initialKeepers\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_jobRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_triggerGasRegistryAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"isComposeMsgSender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isKeeper\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"jobRegistry\",\"outputs\":[{\"internalType\":\"contractIJobRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"_origin\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_guid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"lzReceive\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"nextNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oAppVersion\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"senderVersion\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"receiverVersion\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"originEid\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"}],\"name\":\"peers\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"peer\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"}],\"name\":\"removeKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegate\",\"type\":\"address\"}],\"name\":\"setDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"msgType\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"internalType\":\"structEnforcedOptionParam[]\",\"name\":\"_enforcedOptions\",\"type\":\"tuple[]\"}],\"name\":\"setEnforcedOptions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"gas\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"}],\"name\":\"setGasConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_peer\",\"type\":\"bytes32\"}],\"name\":\"setPeer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"srcEid\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"triggerGasRegistry\",\"outputs\":[{\"internalType\":\"contractITriggerGasRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// ContractTaskExecutionHubABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractTaskExecutionHubMetaData.ABI instead.
var ContractTaskExecutionHubABI = ContractTaskExecutionHubMetaData.ABI

// ContractTaskExecutionHubMethods is an auto generated interface around an Ethereum contract.
type ContractTaskExecutionHubMethods interface {
	ContractTaskExecutionHubCalls
	ContractTaskExecutionHubTransacts
	ContractTaskExecutionHubFilters
}

// ContractTaskExecutionHubCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractTaskExecutionHubCalls interface {
	UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error)

	QuoteFee(opts *bind.CallOpts, dstEid uint32, payload []byte, options []byte) (MessagingFee, error)

	AllowInitializePath(opts *bind.CallOpts, origin Origin) (bool, error)

	CombineOptions(opts *bind.CallOpts, _eid uint32, _msgType uint16, _extraOptions []byte) ([]byte, error)

	DefaultGas(opts *bind.CallOpts) (*big.Int, error)

	DefaultValue(opts *bind.CallOpts) (*big.Int, error)

	DstEids(opts *bind.CallOpts, arg0 *big.Int) (uint32, error)

	Endpoint(opts *bind.CallOpts) (common.Address, error)

	EnforcedOptions(opts *bind.CallOpts, eid uint32, msgType uint16) ([]byte, error)

	IsComposeMsgSender(opts *bind.CallOpts, arg0 Origin, arg1 []byte, _sender common.Address) (bool, error)

	IsKeeper(opts *bind.CallOpts, arg0 common.Address) (bool, error)

	JobRegistry(opts *bind.CallOpts) (common.Address, error)

	NextNonce(opts *bind.CallOpts, arg0 uint32, arg1 [32]byte) (uint64, error)

	OAppVersion(opts *bind.CallOpts) (struct {
		SenderVersion   uint64
		ReceiverVersion uint64
	}, error)

	OriginEid(opts *bind.CallOpts) (uint32, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Peers(opts *bind.CallOpts, eid uint32) ([32]byte, error)

	ProxiableUUID(opts *bind.CallOpts) ([32]byte, error)

	SrcEid(opts *bind.CallOpts) (uint32, error)

	TriggerGasRegistry(opts *bind.CallOpts) (common.Address, error)
}

// ContractTaskExecutionHubTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractTaskExecutionHubTransacts interface {
	AddKeeper(opts *bind.TransactOpts, keeper common.Address) (*types.Transaction, error)

	AddSpokes(opts *bind.TransactOpts, _dstEids []uint32) (*types.Transaction, error)

	ExecuteFunction(opts *bind.TransactOpts, jobId *big.Int, tgAmount *big.Int, target common.Address, data []byte) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts, _ownerAddress common.Address, _srcEid uint32, _originEid uint32, _initialKeepers []common.Address, _jobRegistryAddress common.Address, _triggerGasRegistryAddress common.Address) (*types.Transaction, error)

	LzReceive(opts *bind.TransactOpts, _origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error)

	RemoveKeeper(opts *bind.TransactOpts, keeper common.Address) (*types.Transaction, error)

	RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetDelegate(opts *bind.TransactOpts, _delegate common.Address) (*types.Transaction, error)

	SetEnforcedOptions(opts *bind.TransactOpts, _enforcedOptions []EnforcedOptionParam) (*types.Transaction, error)

	SetGasConfig(opts *bind.TransactOpts, gas *big.Int, value *big.Int) (*types.Transaction, error)

	SetPeer(opts *bind.TransactOpts, _eid uint32, _peer [32]byte) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error)

	UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error)

	Withdraw(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error)
}

// ContractTaskExecutionHubFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractTaskExecutionHubFilters interface {
	FilterBroadcastSent(opts *bind.FilterOpts) (*ContractTaskExecutionHubBroadcastSentIterator, error)
	WatchBroadcastSent(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionHubBroadcastSent) (event.Subscription, error)
	ParseBroadcastSent(log types.Log) (*ContractTaskExecutionHubBroadcastSent, error)

	FilterEnforcedOptionSet(opts *bind.FilterOpts) (*ContractTaskExecutionHubEnforcedOptionSetIterator, error)
	WatchEnforcedOptionSet(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionHubEnforcedOptionSet) (event.Subscription, error)
	ParseEnforcedOptionSet(log types.Log) (*ContractTaskExecutionHubEnforcedOptionSet, error)

	FilterFeeUsed(opts *bind.FilterOpts) (*ContractTaskExecutionHubFeeUsedIterator, error)
	WatchFeeUsed(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionHubFeeUsed) (event.Subscription, error)
	ParseFeeUsed(log types.Log) (*ContractTaskExecutionHubFeeUsed, error)

	FilterFunctionExecuted(opts *bind.FilterOpts, keeper []common.Address, target []common.Address) (*ContractTaskExecutionHubFunctionExecutedIterator, error)
	WatchFunctionExecuted(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionHubFunctionExecuted, keeper []common.Address, target []common.Address) (event.Subscription, error)
	ParseFunctionExecuted(log types.Log) (*ContractTaskExecutionHubFunctionExecuted, error)

	FilterGasConfigUpdated(opts *bind.FilterOpts) (*ContractTaskExecutionHubGasConfigUpdatedIterator, error)
	WatchGasConfigUpdated(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionHubGasConfigUpdated) (event.Subscription, error)
	ParseGasConfigUpdated(log types.Log) (*ContractTaskExecutionHubGasConfigUpdated, error)

	FilterInitialized(opts *bind.FilterOpts) (*ContractTaskExecutionHubInitializedIterator, error)
	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionHubInitialized) (event.Subscription, error)
	ParseInitialized(log types.Log) (*ContractTaskExecutionHubInitialized, error)

	FilterKeeperRegistered(opts *bind.FilterOpts, keeper []common.Address) (*ContractTaskExecutionHubKeeperRegisteredIterator, error)
	WatchKeeperRegistered(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionHubKeeperRegistered, keeper []common.Address) (event.Subscription, error)
	ParseKeeperRegistered(log types.Log) (*ContractTaskExecutionHubKeeperRegistered, error)

	FilterKeeperUnregistered(opts *bind.FilterOpts, keeper []common.Address) (*ContractTaskExecutionHubKeeperUnregisteredIterator, error)
	WatchKeeperUnregistered(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionHubKeeperUnregistered, keeper []common.Address) (event.Subscription, error)
	ParseKeeperUnregistered(log types.Log) (*ContractTaskExecutionHubKeeperUnregistered, error)

	FilterLowBalanceAlert(opts *bind.FilterOpts) (*ContractTaskExecutionHubLowBalanceAlertIterator, error)
	WatchLowBalanceAlert(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionHubLowBalanceAlert) (event.Subscription, error)
	ParseLowBalanceAlert(log types.Log) (*ContractTaskExecutionHubLowBalanceAlert, error)

	FilterMessageFailed(opts *bind.FilterOpts, dstEid []uint32, guid [][32]byte) (*ContractTaskExecutionHubMessageFailedIterator, error)
	WatchMessageFailed(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionHubMessageFailed, dstEid []uint32, guid [][32]byte) (event.Subscription, error)
	ParseMessageFailed(log types.Log) (*ContractTaskExecutionHubMessageFailed, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractTaskExecutionHubOwnershipTransferredIterator, error)
	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionHubOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error)
	ParseOwnershipTransferred(log types.Log) (*ContractTaskExecutionHubOwnershipTransferred, error)

	FilterPeerSet(opts *bind.FilterOpts) (*ContractTaskExecutionHubPeerSetIterator, error)
	WatchPeerSet(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionHubPeerSet) (event.Subscription, error)
	ParsePeerSet(log types.Log) (*ContractTaskExecutionHubPeerSet, error)

	FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ContractTaskExecutionHubUpgradedIterator, error)
	WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionHubUpgraded, implementation []common.Address) (event.Subscription, error)
	ParseUpgraded(log types.Log) (*ContractTaskExecutionHubUpgraded, error)
}

// ContractTaskExecutionHub is an auto generated Go binding around an Ethereum contract.
type ContractTaskExecutionHub struct {
	ContractTaskExecutionHubCaller     // Read-only binding to the contract
	ContractTaskExecutionHubTransactor // Write-only binding to the contract
	ContractTaskExecutionHubFilterer   // Log filterer for contract events
}

// ContractTaskExecutionHub implements the ContractTaskExecutionHubMethods interface.
var _ ContractTaskExecutionHubMethods = (*ContractTaskExecutionHub)(nil)

// ContractTaskExecutionHubCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractTaskExecutionHubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTaskExecutionHubCaller implements the ContractTaskExecutionHubCalls interface.
var _ ContractTaskExecutionHubCalls = (*ContractTaskExecutionHubCaller)(nil)

// ContractTaskExecutionHubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTaskExecutionHubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTaskExecutionHubTransactor implements the ContractTaskExecutionHubTransacts interface.
var _ ContractTaskExecutionHubTransacts = (*ContractTaskExecutionHubTransactor)(nil)

// ContractTaskExecutionHubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractTaskExecutionHubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTaskExecutionHubFilterer implements the ContractTaskExecutionHubFilters interface.
var _ ContractTaskExecutionHubFilters = (*ContractTaskExecutionHubFilterer)(nil)

// ContractTaskExecutionHubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractTaskExecutionHubSession struct {
	Contract     *ContractTaskExecutionHub // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ContractTaskExecutionHubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractTaskExecutionHubCallerSession struct {
	Contract *ContractTaskExecutionHubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// ContractTaskExecutionHubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTaskExecutionHubTransactorSession struct {
	Contract     *ContractTaskExecutionHubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// ContractTaskExecutionHubRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractTaskExecutionHubRaw struct {
	Contract *ContractTaskExecutionHub // Generic contract binding to access the raw methods on
}

// ContractTaskExecutionHubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractTaskExecutionHubCallerRaw struct {
	Contract *ContractTaskExecutionHubCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTaskExecutionHubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTaskExecutionHubTransactorRaw struct {
	Contract *ContractTaskExecutionHubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractTaskExecutionHub creates a new instance of ContractTaskExecutionHub, bound to a specific deployed contract.
func NewContractTaskExecutionHub(address common.Address, backend bind.ContractBackend) (*ContractTaskExecutionHub, error) {
	contract, err := bindContractTaskExecutionHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractTaskExecutionHub{ContractTaskExecutionHubCaller: ContractTaskExecutionHubCaller{contract: contract}, ContractTaskExecutionHubTransactor: ContractTaskExecutionHubTransactor{contract: contract}, ContractTaskExecutionHubFilterer: ContractTaskExecutionHubFilterer{contract: contract}}, nil
}

// NewContractTaskExecutionHubCaller creates a new read-only instance of ContractTaskExecutionHub, bound to a specific deployed contract.
func NewContractTaskExecutionHubCaller(address common.Address, caller bind.ContractCaller) (*ContractTaskExecutionHubCaller, error) {
	contract, err := bindContractTaskExecutionHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTaskExecutionHubCaller{contract: contract}, nil
}

// NewContractTaskExecutionHubTransactor creates a new write-only instance of ContractTaskExecutionHub, bound to a specific deployed contract.
func NewContractTaskExecutionHubTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTaskExecutionHubTransactor, error) {
	contract, err := bindContractTaskExecutionHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTaskExecutionHubTransactor{contract: contract}, nil
}

// NewContractTaskExecutionHubFilterer creates a new log filterer instance of ContractTaskExecutionHub, bound to a specific deployed contract.
func NewContractTaskExecutionHubFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractTaskExecutionHubFilterer, error) {
	contract, err := bindContractTaskExecutionHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractTaskExecutionHubFilterer{contract: contract}, nil
}

// bindContractTaskExecutionHub binds a generic wrapper to an already deployed contract.
func bindContractTaskExecutionHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractTaskExecutionHubMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractTaskExecutionHub *ContractTaskExecutionHubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractTaskExecutionHub.Contract.ContractTaskExecutionHubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractTaskExecutionHub *ContractTaskExecutionHubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.Contract.ContractTaskExecutionHubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractTaskExecutionHub *ContractTaskExecutionHubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.Contract.ContractTaskExecutionHubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractTaskExecutionHub *ContractTaskExecutionHubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractTaskExecutionHub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractTaskExecutionHub *ContractTaskExecutionHubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractTaskExecutionHub *ContractTaskExecutionHubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ContractTaskExecutionHub.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ContractTaskExecutionHub.Contract.UPGRADEINTERFACEVERSION(&_ContractTaskExecutionHub.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ContractTaskExecutionHub.Contract.UPGRADEINTERFACEVERSION(&_ContractTaskExecutionHub.CallOpts)
}

// QuoteFee is a free data retrieval call binding the contract method 0xf0b2a95d.
//
// Solidity: function _quoteFee(uint32 dstEid, bytes payload, bytes options) view returns((uint256,uint256) fee)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubCaller) QuoteFee(opts *bind.CallOpts, dstEid uint32, payload []byte, options []byte) (MessagingFee, error) {
	var out []interface{}
	err := _ContractTaskExecutionHub.contract.Call(opts, &out, "_quoteFee", dstEid, payload, options)

	if err != nil {
		return *new(MessagingFee), err
	}

	out0 := *abi.ConvertType(out[0], new(MessagingFee)).(*MessagingFee)

	return out0, err

}

// QuoteFee is a free data retrieval call binding the contract method 0xf0b2a95d.
//
// Solidity: function _quoteFee(uint32 dstEid, bytes payload, bytes options) view returns((uint256,uint256) fee)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubSession) QuoteFee(dstEid uint32, payload []byte, options []byte) (MessagingFee, error) {
	return _ContractTaskExecutionHub.Contract.QuoteFee(&_ContractTaskExecutionHub.CallOpts, dstEid, payload, options)
}

// QuoteFee is a free data retrieval call binding the contract method 0xf0b2a95d.
//
// Solidity: function _quoteFee(uint32 dstEid, bytes payload, bytes options) view returns((uint256,uint256) fee)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubCallerSession) QuoteFee(dstEid uint32, payload []byte, options []byte) (MessagingFee, error) {
	return _ContractTaskExecutionHub.Contract.QuoteFee(&_ContractTaskExecutionHub.CallOpts, dstEid, payload, options)
}

// AllowInitializePath is a free data retrieval call binding the contract method 0xff7bd03d.
//
// Solidity: function allowInitializePath((uint32,bytes32,uint64) origin) view returns(bool)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubCaller) AllowInitializePath(opts *bind.CallOpts, origin Origin) (bool, error) {
	var out []interface{}
	err := _ContractTaskExecutionHub.contract.Call(opts, &out, "allowInitializePath", origin)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowInitializePath is a free data retrieval call binding the contract method 0xff7bd03d.
//
// Solidity: function allowInitializePath((uint32,bytes32,uint64) origin) view returns(bool)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubSession) AllowInitializePath(origin Origin) (bool, error) {
	return _ContractTaskExecutionHub.Contract.AllowInitializePath(&_ContractTaskExecutionHub.CallOpts, origin)
}

// AllowInitializePath is a free data retrieval call binding the contract method 0xff7bd03d.
//
// Solidity: function allowInitializePath((uint32,bytes32,uint64) origin) view returns(bool)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubCallerSession) AllowInitializePath(origin Origin) (bool, error) {
	return _ContractTaskExecutionHub.Contract.AllowInitializePath(&_ContractTaskExecutionHub.CallOpts, origin)
}

// CombineOptions is a free data retrieval call binding the contract method 0xbc70b354.
//
// Solidity: function combineOptions(uint32 _eid, uint16 _msgType, bytes _extraOptions) view returns(bytes)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubCaller) CombineOptions(opts *bind.CallOpts, _eid uint32, _msgType uint16, _extraOptions []byte) ([]byte, error) {
	var out []interface{}
	err := _ContractTaskExecutionHub.contract.Call(opts, &out, "combineOptions", _eid, _msgType, _extraOptions)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// CombineOptions is a free data retrieval call binding the contract method 0xbc70b354.
//
// Solidity: function combineOptions(uint32 _eid, uint16 _msgType, bytes _extraOptions) view returns(bytes)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubSession) CombineOptions(_eid uint32, _msgType uint16, _extraOptions []byte) ([]byte, error) {
	return _ContractTaskExecutionHub.Contract.CombineOptions(&_ContractTaskExecutionHub.CallOpts, _eid, _msgType, _extraOptions)
}

// CombineOptions is a free data retrieval call binding the contract method 0xbc70b354.
//
// Solidity: function combineOptions(uint32 _eid, uint16 _msgType, bytes _extraOptions) view returns(bytes)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubCallerSession) CombineOptions(_eid uint32, _msgType uint16, _extraOptions []byte) ([]byte, error) {
	return _ContractTaskExecutionHub.Contract.CombineOptions(&_ContractTaskExecutionHub.CallOpts, _eid, _msgType, _extraOptions)
}

// DefaultGas is a free data retrieval call binding the contract method 0x84e9a901.
//
// Solidity: function defaultGas() view returns(uint128)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubCaller) DefaultGas(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractTaskExecutionHub.contract.Call(opts, &out, "defaultGas")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultGas is a free data retrieval call binding the contract method 0x84e9a901.
//
// Solidity: function defaultGas() view returns(uint128)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubSession) DefaultGas() (*big.Int, error) {
	return _ContractTaskExecutionHub.Contract.DefaultGas(&_ContractTaskExecutionHub.CallOpts)
}

// DefaultGas is a free data retrieval call binding the contract method 0x84e9a901.
//
// Solidity: function defaultGas() view returns(uint128)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubCallerSession) DefaultGas() (*big.Int, error) {
	return _ContractTaskExecutionHub.Contract.DefaultGas(&_ContractTaskExecutionHub.CallOpts)
}

// DefaultValue is a free data retrieval call binding the contract method 0xb3ce5817.
//
// Solidity: function defaultValue() view returns(uint128)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubCaller) DefaultValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractTaskExecutionHub.contract.Call(opts, &out, "defaultValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultValue is a free data retrieval call binding the contract method 0xb3ce5817.
//
// Solidity: function defaultValue() view returns(uint128)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubSession) DefaultValue() (*big.Int, error) {
	return _ContractTaskExecutionHub.Contract.DefaultValue(&_ContractTaskExecutionHub.CallOpts)
}

// DefaultValue is a free data retrieval call binding the contract method 0xb3ce5817.
//
// Solidity: function defaultValue() view returns(uint128)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubCallerSession) DefaultValue() (*big.Int, error) {
	return _ContractTaskExecutionHub.Contract.DefaultValue(&_ContractTaskExecutionHub.CallOpts)
}

// DstEids is a free data retrieval call binding the contract method 0x05dc0922.
//
// Solidity: function dstEids(uint256 ) view returns(uint32)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubCaller) DstEids(opts *bind.CallOpts, arg0 *big.Int) (uint32, error) {
	var out []interface{}
	err := _ContractTaskExecutionHub.contract.Call(opts, &out, "dstEids", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// DstEids is a free data retrieval call binding the contract method 0x05dc0922.
//
// Solidity: function dstEids(uint256 ) view returns(uint32)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubSession) DstEids(arg0 *big.Int) (uint32, error) {
	return _ContractTaskExecutionHub.Contract.DstEids(&_ContractTaskExecutionHub.CallOpts, arg0)
}

// DstEids is a free data retrieval call binding the contract method 0x05dc0922.
//
// Solidity: function dstEids(uint256 ) view returns(uint32)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubCallerSession) DstEids(arg0 *big.Int) (uint32, error) {
	return _ContractTaskExecutionHub.Contract.DstEids(&_ContractTaskExecutionHub.CallOpts, arg0)
}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubCaller) Endpoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTaskExecutionHub.contract.Call(opts, &out, "endpoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubSession) Endpoint() (common.Address, error) {
	return _ContractTaskExecutionHub.Contract.Endpoint(&_ContractTaskExecutionHub.CallOpts)
}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubCallerSession) Endpoint() (common.Address, error) {
	return _ContractTaskExecutionHub.Contract.Endpoint(&_ContractTaskExecutionHub.CallOpts)
}

// EnforcedOptions is a free data retrieval call binding the contract method 0x5535d461.
//
// Solidity: function enforcedOptions(uint32 eid, uint16 msgType) view returns(bytes enforcedOption)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubCaller) EnforcedOptions(opts *bind.CallOpts, eid uint32, msgType uint16) ([]byte, error) {
	var out []interface{}
	err := _ContractTaskExecutionHub.contract.Call(opts, &out, "enforcedOptions", eid, msgType)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EnforcedOptions is a free data retrieval call binding the contract method 0x5535d461.
//
// Solidity: function enforcedOptions(uint32 eid, uint16 msgType) view returns(bytes enforcedOption)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubSession) EnforcedOptions(eid uint32, msgType uint16) ([]byte, error) {
	return _ContractTaskExecutionHub.Contract.EnforcedOptions(&_ContractTaskExecutionHub.CallOpts, eid, msgType)
}

// EnforcedOptions is a free data retrieval call binding the contract method 0x5535d461.
//
// Solidity: function enforcedOptions(uint32 eid, uint16 msgType) view returns(bytes enforcedOption)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubCallerSession) EnforcedOptions(eid uint32, msgType uint16) ([]byte, error) {
	return _ContractTaskExecutionHub.Contract.EnforcedOptions(&_ContractTaskExecutionHub.CallOpts, eid, msgType)
}

// IsComposeMsgSender is a free data retrieval call binding the contract method 0x82413eac.
//
// Solidity: function isComposeMsgSender((uint32,bytes32,uint64) , bytes , address _sender) view returns(bool)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubCaller) IsComposeMsgSender(opts *bind.CallOpts, arg0 Origin, arg1 []byte, _sender common.Address) (bool, error) {
	var out []interface{}
	err := _ContractTaskExecutionHub.contract.Call(opts, &out, "isComposeMsgSender", arg0, arg1, _sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsComposeMsgSender is a free data retrieval call binding the contract method 0x82413eac.
//
// Solidity: function isComposeMsgSender((uint32,bytes32,uint64) , bytes , address _sender) view returns(bool)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubSession) IsComposeMsgSender(arg0 Origin, arg1 []byte, _sender common.Address) (bool, error) {
	return _ContractTaskExecutionHub.Contract.IsComposeMsgSender(&_ContractTaskExecutionHub.CallOpts, arg0, arg1, _sender)
}

// IsComposeMsgSender is a free data retrieval call binding the contract method 0x82413eac.
//
// Solidity: function isComposeMsgSender((uint32,bytes32,uint64) , bytes , address _sender) view returns(bool)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubCallerSession) IsComposeMsgSender(arg0 Origin, arg1 []byte, _sender common.Address) (bool, error) {
	return _ContractTaskExecutionHub.Contract.IsComposeMsgSender(&_ContractTaskExecutionHub.CallOpts, arg0, arg1, _sender)
}

// IsKeeper is a free data retrieval call binding the contract method 0x6ba42aaa.
//
// Solidity: function isKeeper(address ) view returns(bool)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubCaller) IsKeeper(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ContractTaskExecutionHub.contract.Call(opts, &out, "isKeeper", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsKeeper is a free data retrieval call binding the contract method 0x6ba42aaa.
//
// Solidity: function isKeeper(address ) view returns(bool)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubSession) IsKeeper(arg0 common.Address) (bool, error) {
	return _ContractTaskExecutionHub.Contract.IsKeeper(&_ContractTaskExecutionHub.CallOpts, arg0)
}

// IsKeeper is a free data retrieval call binding the contract method 0x6ba42aaa.
//
// Solidity: function isKeeper(address ) view returns(bool)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubCallerSession) IsKeeper(arg0 common.Address) (bool, error) {
	return _ContractTaskExecutionHub.Contract.IsKeeper(&_ContractTaskExecutionHub.CallOpts, arg0)
}

// JobRegistry is a free data retrieval call binding the contract method 0x23682c47.
//
// Solidity: function jobRegistry() view returns(address)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubCaller) JobRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTaskExecutionHub.contract.Call(opts, &out, "jobRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// JobRegistry is a free data retrieval call binding the contract method 0x23682c47.
//
// Solidity: function jobRegistry() view returns(address)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubSession) JobRegistry() (common.Address, error) {
	return _ContractTaskExecutionHub.Contract.JobRegistry(&_ContractTaskExecutionHub.CallOpts)
}

// JobRegistry is a free data retrieval call binding the contract method 0x23682c47.
//
// Solidity: function jobRegistry() view returns(address)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubCallerSession) JobRegistry() (common.Address, error) {
	return _ContractTaskExecutionHub.Contract.JobRegistry(&_ContractTaskExecutionHub.CallOpts)
}

// NextNonce is a free data retrieval call binding the contract method 0x7d25a05e.
//
// Solidity: function nextNonce(uint32 , bytes32 ) view returns(uint64 nonce)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubCaller) NextNonce(opts *bind.CallOpts, arg0 uint32, arg1 [32]byte) (uint64, error) {
	var out []interface{}
	err := _ContractTaskExecutionHub.contract.Call(opts, &out, "nextNonce", arg0, arg1)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NextNonce is a free data retrieval call binding the contract method 0x7d25a05e.
//
// Solidity: function nextNonce(uint32 , bytes32 ) view returns(uint64 nonce)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubSession) NextNonce(arg0 uint32, arg1 [32]byte) (uint64, error) {
	return _ContractTaskExecutionHub.Contract.NextNonce(&_ContractTaskExecutionHub.CallOpts, arg0, arg1)
}

// NextNonce is a free data retrieval call binding the contract method 0x7d25a05e.
//
// Solidity: function nextNonce(uint32 , bytes32 ) view returns(uint64 nonce)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubCallerSession) NextNonce(arg0 uint32, arg1 [32]byte) (uint64, error) {
	return _ContractTaskExecutionHub.Contract.NextNonce(&_ContractTaskExecutionHub.CallOpts, arg0, arg1)
}

// OAppVersion is a free data retrieval call binding the contract method 0x17442b70.
//
// Solidity: function oAppVersion() pure returns(uint64 senderVersion, uint64 receiverVersion)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubCaller) OAppVersion(opts *bind.CallOpts) (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	var out []interface{}
	err := _ContractTaskExecutionHub.contract.Call(opts, &out, "oAppVersion")

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
func (_ContractTaskExecutionHub *ContractTaskExecutionHubSession) OAppVersion() (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	return _ContractTaskExecutionHub.Contract.OAppVersion(&_ContractTaskExecutionHub.CallOpts)
}

// OAppVersion is a free data retrieval call binding the contract method 0x17442b70.
//
// Solidity: function oAppVersion() pure returns(uint64 senderVersion, uint64 receiverVersion)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubCallerSession) OAppVersion() (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	return _ContractTaskExecutionHub.Contract.OAppVersion(&_ContractTaskExecutionHub.CallOpts)
}

// OriginEid is a free data retrieval call binding the contract method 0xfd416257.
//
// Solidity: function originEid() view returns(uint32)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubCaller) OriginEid(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractTaskExecutionHub.contract.Call(opts, &out, "originEid")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// OriginEid is a free data retrieval call binding the contract method 0xfd416257.
//
// Solidity: function originEid() view returns(uint32)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubSession) OriginEid() (uint32, error) {
	return _ContractTaskExecutionHub.Contract.OriginEid(&_ContractTaskExecutionHub.CallOpts)
}

// OriginEid is a free data retrieval call binding the contract method 0xfd416257.
//
// Solidity: function originEid() view returns(uint32)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubCallerSession) OriginEid() (uint32, error) {
	return _ContractTaskExecutionHub.Contract.OriginEid(&_ContractTaskExecutionHub.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTaskExecutionHub.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubSession) Owner() (common.Address, error) {
	return _ContractTaskExecutionHub.Contract.Owner(&_ContractTaskExecutionHub.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubCallerSession) Owner() (common.Address, error) {
	return _ContractTaskExecutionHub.Contract.Owner(&_ContractTaskExecutionHub.CallOpts)
}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 eid) view returns(bytes32 peer)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubCaller) Peers(opts *bind.CallOpts, eid uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractTaskExecutionHub.contract.Call(opts, &out, "peers", eid)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 eid) view returns(bytes32 peer)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubSession) Peers(eid uint32) ([32]byte, error) {
	return _ContractTaskExecutionHub.Contract.Peers(&_ContractTaskExecutionHub.CallOpts, eid)
}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 eid) view returns(bytes32 peer)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubCallerSession) Peers(eid uint32) ([32]byte, error) {
	return _ContractTaskExecutionHub.Contract.Peers(&_ContractTaskExecutionHub.CallOpts, eid)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractTaskExecutionHub.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubSession) ProxiableUUID() ([32]byte, error) {
	return _ContractTaskExecutionHub.Contract.ProxiableUUID(&_ContractTaskExecutionHub.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubCallerSession) ProxiableUUID() ([32]byte, error) {
	return _ContractTaskExecutionHub.Contract.ProxiableUUID(&_ContractTaskExecutionHub.CallOpts)
}

// SrcEid is a free data retrieval call binding the contract method 0x479cca5e.
//
// Solidity: function srcEid() view returns(uint32)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubCaller) SrcEid(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractTaskExecutionHub.contract.Call(opts, &out, "srcEid")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SrcEid is a free data retrieval call binding the contract method 0x479cca5e.
//
// Solidity: function srcEid() view returns(uint32)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubSession) SrcEid() (uint32, error) {
	return _ContractTaskExecutionHub.Contract.SrcEid(&_ContractTaskExecutionHub.CallOpts)
}

// SrcEid is a free data retrieval call binding the contract method 0x479cca5e.
//
// Solidity: function srcEid() view returns(uint32)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubCallerSession) SrcEid() (uint32, error) {
	return _ContractTaskExecutionHub.Contract.SrcEid(&_ContractTaskExecutionHub.CallOpts)
}

// TriggerGasRegistry is a free data retrieval call binding the contract method 0xd908b2ef.
//
// Solidity: function triggerGasRegistry() view returns(address)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubCaller) TriggerGasRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTaskExecutionHub.contract.Call(opts, &out, "triggerGasRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TriggerGasRegistry is a free data retrieval call binding the contract method 0xd908b2ef.
//
// Solidity: function triggerGasRegistry() view returns(address)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubSession) TriggerGasRegistry() (common.Address, error) {
	return _ContractTaskExecutionHub.Contract.TriggerGasRegistry(&_ContractTaskExecutionHub.CallOpts)
}

// TriggerGasRegistry is a free data retrieval call binding the contract method 0xd908b2ef.
//
// Solidity: function triggerGasRegistry() view returns(address)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubCallerSession) TriggerGasRegistry() (common.Address, error) {
	return _ContractTaskExecutionHub.Contract.TriggerGasRegistry(&_ContractTaskExecutionHub.CallOpts)
}

// AddKeeper is a paid mutator transaction binding the contract method 0x4032b72b.
//
// Solidity: function addKeeper(address keeper) returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubTransactor) AddKeeper(opts *bind.TransactOpts, keeper common.Address) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.contract.Transact(opts, "addKeeper", keeper)
}

// AddKeeper is a paid mutator transaction binding the contract method 0x4032b72b.
//
// Solidity: function addKeeper(address keeper) returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubSession) AddKeeper(keeper common.Address) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.Contract.AddKeeper(&_ContractTaskExecutionHub.TransactOpts, keeper)
}

// AddKeeper is a paid mutator transaction binding the contract method 0x4032b72b.
//
// Solidity: function addKeeper(address keeper) returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubTransactorSession) AddKeeper(keeper common.Address) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.Contract.AddKeeper(&_ContractTaskExecutionHub.TransactOpts, keeper)
}

// AddSpokes is a paid mutator transaction binding the contract method 0x14bc03fb.
//
// Solidity: function addSpokes(uint32[] _dstEids) returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubTransactor) AddSpokes(opts *bind.TransactOpts, _dstEids []uint32) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.contract.Transact(opts, "addSpokes", _dstEids)
}

// AddSpokes is a paid mutator transaction binding the contract method 0x14bc03fb.
//
// Solidity: function addSpokes(uint32[] _dstEids) returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubSession) AddSpokes(_dstEids []uint32) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.Contract.AddSpokes(&_ContractTaskExecutionHub.TransactOpts, _dstEids)
}

// AddSpokes is a paid mutator transaction binding the contract method 0x14bc03fb.
//
// Solidity: function addSpokes(uint32[] _dstEids) returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubTransactorSession) AddSpokes(_dstEids []uint32) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.Contract.AddSpokes(&_ContractTaskExecutionHub.TransactOpts, _dstEids)
}

// ExecuteFunction is a paid mutator transaction binding the contract method 0xfa9b1a80.
//
// Solidity: function executeFunction(uint256 jobId, uint256 tgAmount, address target, bytes data) payable returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubTransactor) ExecuteFunction(opts *bind.TransactOpts, jobId *big.Int, tgAmount *big.Int, target common.Address, data []byte) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.contract.Transact(opts, "executeFunction", jobId, tgAmount, target, data)
}

// ExecuteFunction is a paid mutator transaction binding the contract method 0xfa9b1a80.
//
// Solidity: function executeFunction(uint256 jobId, uint256 tgAmount, address target, bytes data) payable returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubSession) ExecuteFunction(jobId *big.Int, tgAmount *big.Int, target common.Address, data []byte) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.Contract.ExecuteFunction(&_ContractTaskExecutionHub.TransactOpts, jobId, tgAmount, target, data)
}

// ExecuteFunction is a paid mutator transaction binding the contract method 0xfa9b1a80.
//
// Solidity: function executeFunction(uint256 jobId, uint256 tgAmount, address target, bytes data) payable returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubTransactorSession) ExecuteFunction(jobId *big.Int, tgAmount *big.Int, target common.Address, data []byte) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.Contract.ExecuteFunction(&_ContractTaskExecutionHub.TransactOpts, jobId, tgAmount, target, data)
}

// Initialize is a paid mutator transaction binding the contract method 0x9c983124.
//
// Solidity: function initialize(address _ownerAddress, uint32 _srcEid, uint32 _originEid, address[] _initialKeepers, address _jobRegistryAddress, address _triggerGasRegistryAddress) returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubTransactor) Initialize(opts *bind.TransactOpts, _ownerAddress common.Address, _srcEid uint32, _originEid uint32, _initialKeepers []common.Address, _jobRegistryAddress common.Address, _triggerGasRegistryAddress common.Address) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.contract.Transact(opts, "initialize", _ownerAddress, _srcEid, _originEid, _initialKeepers, _jobRegistryAddress, _triggerGasRegistryAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x9c983124.
//
// Solidity: function initialize(address _ownerAddress, uint32 _srcEid, uint32 _originEid, address[] _initialKeepers, address _jobRegistryAddress, address _triggerGasRegistryAddress) returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubSession) Initialize(_ownerAddress common.Address, _srcEid uint32, _originEid uint32, _initialKeepers []common.Address, _jobRegistryAddress common.Address, _triggerGasRegistryAddress common.Address) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.Contract.Initialize(&_ContractTaskExecutionHub.TransactOpts, _ownerAddress, _srcEid, _originEid, _initialKeepers, _jobRegistryAddress, _triggerGasRegistryAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x9c983124.
//
// Solidity: function initialize(address _ownerAddress, uint32 _srcEid, uint32 _originEid, address[] _initialKeepers, address _jobRegistryAddress, address _triggerGasRegistryAddress) returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubTransactorSession) Initialize(_ownerAddress common.Address, _srcEid uint32, _originEid uint32, _initialKeepers []common.Address, _jobRegistryAddress common.Address, _triggerGasRegistryAddress common.Address) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.Contract.Initialize(&_ContractTaskExecutionHub.TransactOpts, _ownerAddress, _srcEid, _originEid, _initialKeepers, _jobRegistryAddress, _triggerGasRegistryAddress)
}

// LzReceive is a paid mutator transaction binding the contract method 0x13137d65.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubTransactor) LzReceive(opts *bind.TransactOpts, _origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.contract.Transact(opts, "lzReceive", _origin, _guid, _message, _executor, _extraData)
}

// LzReceive is a paid mutator transaction binding the contract method 0x13137d65.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubSession) LzReceive(_origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.Contract.LzReceive(&_ContractTaskExecutionHub.TransactOpts, _origin, _guid, _message, _executor, _extraData)
}

// LzReceive is a paid mutator transaction binding the contract method 0x13137d65.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubTransactorSession) LzReceive(_origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.Contract.LzReceive(&_ContractTaskExecutionHub.TransactOpts, _origin, _guid, _message, _executor, _extraData)
}

// RemoveKeeper is a paid mutator transaction binding the contract method 0x14ae9f2e.
//
// Solidity: function removeKeeper(address keeper) returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubTransactor) RemoveKeeper(opts *bind.TransactOpts, keeper common.Address) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.contract.Transact(opts, "removeKeeper", keeper)
}

// RemoveKeeper is a paid mutator transaction binding the contract method 0x14ae9f2e.
//
// Solidity: function removeKeeper(address keeper) returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubSession) RemoveKeeper(keeper common.Address) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.Contract.RemoveKeeper(&_ContractTaskExecutionHub.TransactOpts, keeper)
}

// RemoveKeeper is a paid mutator transaction binding the contract method 0x14ae9f2e.
//
// Solidity: function removeKeeper(address keeper) returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubTransactorSession) RemoveKeeper(keeper common.Address) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.Contract.RemoveKeeper(&_ContractTaskExecutionHub.TransactOpts, keeper)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractTaskExecutionHub.Contract.RenounceOwnership(&_ContractTaskExecutionHub.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractTaskExecutionHub.Contract.RenounceOwnership(&_ContractTaskExecutionHub.TransactOpts)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubTransactor) SetDelegate(opts *bind.TransactOpts, _delegate common.Address) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.contract.Transact(opts, "setDelegate", _delegate)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubSession) SetDelegate(_delegate common.Address) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.Contract.SetDelegate(&_ContractTaskExecutionHub.TransactOpts, _delegate)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubTransactorSession) SetDelegate(_delegate common.Address) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.Contract.SetDelegate(&_ContractTaskExecutionHub.TransactOpts, _delegate)
}

// SetEnforcedOptions is a paid mutator transaction binding the contract method 0xb98bd070.
//
// Solidity: function setEnforcedOptions((uint32,uint16,bytes)[] _enforcedOptions) returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubTransactor) SetEnforcedOptions(opts *bind.TransactOpts, _enforcedOptions []EnforcedOptionParam) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.contract.Transact(opts, "setEnforcedOptions", _enforcedOptions)
}

// SetEnforcedOptions is a paid mutator transaction binding the contract method 0xb98bd070.
//
// Solidity: function setEnforcedOptions((uint32,uint16,bytes)[] _enforcedOptions) returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubSession) SetEnforcedOptions(_enforcedOptions []EnforcedOptionParam) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.Contract.SetEnforcedOptions(&_ContractTaskExecutionHub.TransactOpts, _enforcedOptions)
}

// SetEnforcedOptions is a paid mutator transaction binding the contract method 0xb98bd070.
//
// Solidity: function setEnforcedOptions((uint32,uint16,bytes)[] _enforcedOptions) returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubTransactorSession) SetEnforcedOptions(_enforcedOptions []EnforcedOptionParam) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.Contract.SetEnforcedOptions(&_ContractTaskExecutionHub.TransactOpts, _enforcedOptions)
}

// SetGasConfig is a paid mutator transaction binding the contract method 0x76a9b1f8.
//
// Solidity: function setGasConfig(uint128 gas, uint128 value) returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubTransactor) SetGasConfig(opts *bind.TransactOpts, gas *big.Int, value *big.Int) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.contract.Transact(opts, "setGasConfig", gas, value)
}

// SetGasConfig is a paid mutator transaction binding the contract method 0x76a9b1f8.
//
// Solidity: function setGasConfig(uint128 gas, uint128 value) returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubSession) SetGasConfig(gas *big.Int, value *big.Int) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.Contract.SetGasConfig(&_ContractTaskExecutionHub.TransactOpts, gas, value)
}

// SetGasConfig is a paid mutator transaction binding the contract method 0x76a9b1f8.
//
// Solidity: function setGasConfig(uint128 gas, uint128 value) returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubTransactorSession) SetGasConfig(gas *big.Int, value *big.Int) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.Contract.SetGasConfig(&_ContractTaskExecutionHub.TransactOpts, gas, value)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubTransactor) SetPeer(opts *bind.TransactOpts, _eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.contract.Transact(opts, "setPeer", _eid, _peer)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubSession) SetPeer(_eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.Contract.SetPeer(&_ContractTaskExecutionHub.TransactOpts, _eid, _peer)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubTransactorSession) SetPeer(_eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.Contract.SetPeer(&_ContractTaskExecutionHub.TransactOpts, _eid, _peer)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.Contract.TransferOwnership(&_ContractTaskExecutionHub.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.Contract.TransferOwnership(&_ContractTaskExecutionHub.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.Contract.UpgradeToAndCall(&_ContractTaskExecutionHub.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.Contract.UpgradeToAndCall(&_ContractTaskExecutionHub.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 amount) returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubTransactor) Withdraw(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.contract.Transact(opts, "withdraw", to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 amount) returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubSession) Withdraw(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.Contract.Withdraw(&_ContractTaskExecutionHub.TransactOpts, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 amount) returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubTransactorSession) Withdraw(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.Contract.Withdraw(&_ContractTaskExecutionHub.TransactOpts, to, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractTaskExecutionHub.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubSession) Receive() (*types.Transaction, error) {
	return _ContractTaskExecutionHub.Contract.Receive(&_ContractTaskExecutionHub.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ContractTaskExecutionHub *ContractTaskExecutionHubTransactorSession) Receive() (*types.Transaction, error) {
	return _ContractTaskExecutionHub.Contract.Receive(&_ContractTaskExecutionHub.TransactOpts)
}

// ContractTaskExecutionHubBroadcastSentIterator is returned from FilterBroadcastSent and is used to iterate over the raw logs and unpacked data for BroadcastSent events raised by the ContractTaskExecutionHub contract.
type ContractTaskExecutionHubBroadcastSentIterator struct {
	Event *ContractTaskExecutionHubBroadcastSent // Event containing the contract specifics and raw log

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
func (it *ContractTaskExecutionHubBroadcastSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTaskExecutionHubBroadcastSent)
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
		it.Event = new(ContractTaskExecutionHubBroadcastSent)
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
func (it *ContractTaskExecutionHubBroadcastSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTaskExecutionHubBroadcastSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTaskExecutionHubBroadcastSent represents a BroadcastSent event raised by the ContractTaskExecutionHub contract.
type ContractTaskExecutionHubBroadcastSent struct {
	Action uint8
	Keeper common.Address
	DstEid uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBroadcastSent is a free log retrieval operation binding the contract event 0x43cd42d638b31afc582e32f34f92ec252d43317313f573b84478d17a4b1706cc.
//
// Solidity: event BroadcastSent(uint8 action, address keeper, uint32 dstEid)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubFilterer) FilterBroadcastSent(opts *bind.FilterOpts) (*ContractTaskExecutionHubBroadcastSentIterator, error) {

	logs, sub, err := _ContractTaskExecutionHub.contract.FilterLogs(opts, "BroadcastSent")
	if err != nil {
		return nil, err
	}
	return &ContractTaskExecutionHubBroadcastSentIterator{contract: _ContractTaskExecutionHub.contract, event: "BroadcastSent", logs: logs, sub: sub}, nil
}

// WatchBroadcastSent is a free log subscription operation binding the contract event 0x43cd42d638b31afc582e32f34f92ec252d43317313f573b84478d17a4b1706cc.
//
// Solidity: event BroadcastSent(uint8 action, address keeper, uint32 dstEid)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubFilterer) WatchBroadcastSent(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionHubBroadcastSent) (event.Subscription, error) {

	logs, sub, err := _ContractTaskExecutionHub.contract.WatchLogs(opts, "BroadcastSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTaskExecutionHubBroadcastSent)
				if err := _ContractTaskExecutionHub.contract.UnpackLog(event, "BroadcastSent", log); err != nil {
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

// ParseBroadcastSent is a log parse operation binding the contract event 0x43cd42d638b31afc582e32f34f92ec252d43317313f573b84478d17a4b1706cc.
//
// Solidity: event BroadcastSent(uint8 action, address keeper, uint32 dstEid)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubFilterer) ParseBroadcastSent(log types.Log) (*ContractTaskExecutionHubBroadcastSent, error) {
	event := new(ContractTaskExecutionHubBroadcastSent)
	if err := _ContractTaskExecutionHub.contract.UnpackLog(event, "BroadcastSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTaskExecutionHubEnforcedOptionSetIterator is returned from FilterEnforcedOptionSet and is used to iterate over the raw logs and unpacked data for EnforcedOptionSet events raised by the ContractTaskExecutionHub contract.
type ContractTaskExecutionHubEnforcedOptionSetIterator struct {
	Event *ContractTaskExecutionHubEnforcedOptionSet // Event containing the contract specifics and raw log

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
func (it *ContractTaskExecutionHubEnforcedOptionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTaskExecutionHubEnforcedOptionSet)
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
		it.Event = new(ContractTaskExecutionHubEnforcedOptionSet)
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
func (it *ContractTaskExecutionHubEnforcedOptionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTaskExecutionHubEnforcedOptionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTaskExecutionHubEnforcedOptionSet represents a EnforcedOptionSet event raised by the ContractTaskExecutionHub contract.
type ContractTaskExecutionHubEnforcedOptionSet struct {
	EnforcedOptions []EnforcedOptionParam
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterEnforcedOptionSet is a free log retrieval operation binding the contract event 0xbe4864a8e820971c0247f5992e2da559595f7bf076a21cb5928d443d2a13b674.
//
// Solidity: event EnforcedOptionSet((uint32,uint16,bytes)[] _enforcedOptions)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubFilterer) FilterEnforcedOptionSet(opts *bind.FilterOpts) (*ContractTaskExecutionHubEnforcedOptionSetIterator, error) {

	logs, sub, err := _ContractTaskExecutionHub.contract.FilterLogs(opts, "EnforcedOptionSet")
	if err != nil {
		return nil, err
	}
	return &ContractTaskExecutionHubEnforcedOptionSetIterator{contract: _ContractTaskExecutionHub.contract, event: "EnforcedOptionSet", logs: logs, sub: sub}, nil
}

// WatchEnforcedOptionSet is a free log subscription operation binding the contract event 0xbe4864a8e820971c0247f5992e2da559595f7bf076a21cb5928d443d2a13b674.
//
// Solidity: event EnforcedOptionSet((uint32,uint16,bytes)[] _enforcedOptions)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubFilterer) WatchEnforcedOptionSet(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionHubEnforcedOptionSet) (event.Subscription, error) {

	logs, sub, err := _ContractTaskExecutionHub.contract.WatchLogs(opts, "EnforcedOptionSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTaskExecutionHubEnforcedOptionSet)
				if err := _ContractTaskExecutionHub.contract.UnpackLog(event, "EnforcedOptionSet", log); err != nil {
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

// ParseEnforcedOptionSet is a log parse operation binding the contract event 0xbe4864a8e820971c0247f5992e2da559595f7bf076a21cb5928d443d2a13b674.
//
// Solidity: event EnforcedOptionSet((uint32,uint16,bytes)[] _enforcedOptions)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubFilterer) ParseEnforcedOptionSet(log types.Log) (*ContractTaskExecutionHubEnforcedOptionSet, error) {
	event := new(ContractTaskExecutionHubEnforcedOptionSet)
	if err := _ContractTaskExecutionHub.contract.UnpackLog(event, "EnforcedOptionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTaskExecutionHubFeeUsedIterator is returned from FilterFeeUsed and is used to iterate over the raw logs and unpacked data for FeeUsed events raised by the ContractTaskExecutionHub contract.
type ContractTaskExecutionHubFeeUsedIterator struct {
	Event *ContractTaskExecutionHubFeeUsed // Event containing the contract specifics and raw log

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
func (it *ContractTaskExecutionHubFeeUsedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTaskExecutionHubFeeUsed)
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
		it.Event = new(ContractTaskExecutionHubFeeUsed)
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
func (it *ContractTaskExecutionHubFeeUsedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTaskExecutionHubFeeUsedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTaskExecutionHubFeeUsed represents a FeeUsed event raised by the ContractTaskExecutionHub contract.
type ContractTaskExecutionHubFeeUsed struct {
	DstEid uint32
	Fee    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeeUsed is a free log retrieval operation binding the contract event 0x7f8a66afaab7a89758bc21c10c08f65792a914c8459da50f53d4d6a6584d9193.
//
// Solidity: event FeeUsed(uint32 dstEid, uint256 fee)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubFilterer) FilterFeeUsed(opts *bind.FilterOpts) (*ContractTaskExecutionHubFeeUsedIterator, error) {

	logs, sub, err := _ContractTaskExecutionHub.contract.FilterLogs(opts, "FeeUsed")
	if err != nil {
		return nil, err
	}
	return &ContractTaskExecutionHubFeeUsedIterator{contract: _ContractTaskExecutionHub.contract, event: "FeeUsed", logs: logs, sub: sub}, nil
}

// WatchFeeUsed is a free log subscription operation binding the contract event 0x7f8a66afaab7a89758bc21c10c08f65792a914c8459da50f53d4d6a6584d9193.
//
// Solidity: event FeeUsed(uint32 dstEid, uint256 fee)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubFilterer) WatchFeeUsed(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionHubFeeUsed) (event.Subscription, error) {

	logs, sub, err := _ContractTaskExecutionHub.contract.WatchLogs(opts, "FeeUsed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTaskExecutionHubFeeUsed)
				if err := _ContractTaskExecutionHub.contract.UnpackLog(event, "FeeUsed", log); err != nil {
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

// ParseFeeUsed is a log parse operation binding the contract event 0x7f8a66afaab7a89758bc21c10c08f65792a914c8459da50f53d4d6a6584d9193.
//
// Solidity: event FeeUsed(uint32 dstEid, uint256 fee)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubFilterer) ParseFeeUsed(log types.Log) (*ContractTaskExecutionHubFeeUsed, error) {
	event := new(ContractTaskExecutionHubFeeUsed)
	if err := _ContractTaskExecutionHub.contract.UnpackLog(event, "FeeUsed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTaskExecutionHubFunctionExecutedIterator is returned from FilterFunctionExecuted and is used to iterate over the raw logs and unpacked data for FunctionExecuted events raised by the ContractTaskExecutionHub contract.
type ContractTaskExecutionHubFunctionExecutedIterator struct {
	Event *ContractTaskExecutionHubFunctionExecuted // Event containing the contract specifics and raw log

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
func (it *ContractTaskExecutionHubFunctionExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTaskExecutionHubFunctionExecuted)
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
		it.Event = new(ContractTaskExecutionHubFunctionExecuted)
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
func (it *ContractTaskExecutionHubFunctionExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTaskExecutionHubFunctionExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTaskExecutionHubFunctionExecuted represents a FunctionExecuted event raised by the ContractTaskExecutionHub contract.
type ContractTaskExecutionHubFunctionExecuted struct {
	Keeper common.Address
	Target common.Address
	Data   []byte
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFunctionExecuted is a free log retrieval operation binding the contract event 0xf4448cdaf10358453fa19b2e0363f44780277619bd9bb676eef79d7488a7d274.
//
// Solidity: event FunctionExecuted(address indexed keeper, address indexed target, bytes data, uint256 value)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubFilterer) FilterFunctionExecuted(opts *bind.FilterOpts, keeper []common.Address, target []common.Address) (*ContractTaskExecutionHubFunctionExecutedIterator, error) {

	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _ContractTaskExecutionHub.contract.FilterLogs(opts, "FunctionExecuted", keeperRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &ContractTaskExecutionHubFunctionExecutedIterator{contract: _ContractTaskExecutionHub.contract, event: "FunctionExecuted", logs: logs, sub: sub}, nil
}

// WatchFunctionExecuted is a free log subscription operation binding the contract event 0xf4448cdaf10358453fa19b2e0363f44780277619bd9bb676eef79d7488a7d274.
//
// Solidity: event FunctionExecuted(address indexed keeper, address indexed target, bytes data, uint256 value)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubFilterer) WatchFunctionExecuted(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionHubFunctionExecuted, keeper []common.Address, target []common.Address) (event.Subscription, error) {

	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _ContractTaskExecutionHub.contract.WatchLogs(opts, "FunctionExecuted", keeperRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTaskExecutionHubFunctionExecuted)
				if err := _ContractTaskExecutionHub.contract.UnpackLog(event, "FunctionExecuted", log); err != nil {
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
func (_ContractTaskExecutionHub *ContractTaskExecutionHubFilterer) ParseFunctionExecuted(log types.Log) (*ContractTaskExecutionHubFunctionExecuted, error) {
	event := new(ContractTaskExecutionHubFunctionExecuted)
	if err := _ContractTaskExecutionHub.contract.UnpackLog(event, "FunctionExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTaskExecutionHubGasConfigUpdatedIterator is returned from FilterGasConfigUpdated and is used to iterate over the raw logs and unpacked data for GasConfigUpdated events raised by the ContractTaskExecutionHub contract.
type ContractTaskExecutionHubGasConfigUpdatedIterator struct {
	Event *ContractTaskExecutionHubGasConfigUpdated // Event containing the contract specifics and raw log

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
func (it *ContractTaskExecutionHubGasConfigUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTaskExecutionHubGasConfigUpdated)
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
		it.Event = new(ContractTaskExecutionHubGasConfigUpdated)
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
func (it *ContractTaskExecutionHubGasConfigUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTaskExecutionHubGasConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTaskExecutionHubGasConfigUpdated represents a GasConfigUpdated event raised by the ContractTaskExecutionHub contract.
type ContractTaskExecutionHubGasConfigUpdated struct {
	Gas   *big.Int
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGasConfigUpdated is a free log retrieval operation binding the contract event 0xd7a0084aec960f13956a34c7a5a0ce5e4620bca53376284f7a612f0ddcd1efd5.
//
// Solidity: event GasConfigUpdated(uint128 gas, uint128 value)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubFilterer) FilterGasConfigUpdated(opts *bind.FilterOpts) (*ContractTaskExecutionHubGasConfigUpdatedIterator, error) {

	logs, sub, err := _ContractTaskExecutionHub.contract.FilterLogs(opts, "GasConfigUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractTaskExecutionHubGasConfigUpdatedIterator{contract: _ContractTaskExecutionHub.contract, event: "GasConfigUpdated", logs: logs, sub: sub}, nil
}

// WatchGasConfigUpdated is a free log subscription operation binding the contract event 0xd7a0084aec960f13956a34c7a5a0ce5e4620bca53376284f7a612f0ddcd1efd5.
//
// Solidity: event GasConfigUpdated(uint128 gas, uint128 value)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubFilterer) WatchGasConfigUpdated(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionHubGasConfigUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractTaskExecutionHub.contract.WatchLogs(opts, "GasConfigUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTaskExecutionHubGasConfigUpdated)
				if err := _ContractTaskExecutionHub.contract.UnpackLog(event, "GasConfigUpdated", log); err != nil {
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

// ParseGasConfigUpdated is a log parse operation binding the contract event 0xd7a0084aec960f13956a34c7a5a0ce5e4620bca53376284f7a612f0ddcd1efd5.
//
// Solidity: event GasConfigUpdated(uint128 gas, uint128 value)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubFilterer) ParseGasConfigUpdated(log types.Log) (*ContractTaskExecutionHubGasConfigUpdated, error) {
	event := new(ContractTaskExecutionHubGasConfigUpdated)
	if err := _ContractTaskExecutionHub.contract.UnpackLog(event, "GasConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTaskExecutionHubInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractTaskExecutionHub contract.
type ContractTaskExecutionHubInitializedIterator struct {
	Event *ContractTaskExecutionHubInitialized // Event containing the contract specifics and raw log

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
func (it *ContractTaskExecutionHubInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTaskExecutionHubInitialized)
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
		it.Event = new(ContractTaskExecutionHubInitialized)
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
func (it *ContractTaskExecutionHubInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTaskExecutionHubInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTaskExecutionHubInitialized represents a Initialized event raised by the ContractTaskExecutionHub contract.
type ContractTaskExecutionHubInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractTaskExecutionHubInitializedIterator, error) {

	logs, sub, err := _ContractTaskExecutionHub.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractTaskExecutionHubInitializedIterator{contract: _ContractTaskExecutionHub.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionHubInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractTaskExecutionHub.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTaskExecutionHubInitialized)
				if err := _ContractTaskExecutionHub.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractTaskExecutionHub *ContractTaskExecutionHubFilterer) ParseInitialized(log types.Log) (*ContractTaskExecutionHubInitialized, error) {
	event := new(ContractTaskExecutionHubInitialized)
	if err := _ContractTaskExecutionHub.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTaskExecutionHubKeeperRegisteredIterator is returned from FilterKeeperRegistered and is used to iterate over the raw logs and unpacked data for KeeperRegistered events raised by the ContractTaskExecutionHub contract.
type ContractTaskExecutionHubKeeperRegisteredIterator struct {
	Event *ContractTaskExecutionHubKeeperRegistered // Event containing the contract specifics and raw log

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
func (it *ContractTaskExecutionHubKeeperRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTaskExecutionHubKeeperRegistered)
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
		it.Event = new(ContractTaskExecutionHubKeeperRegistered)
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
func (it *ContractTaskExecutionHubKeeperRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTaskExecutionHubKeeperRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTaskExecutionHubKeeperRegistered represents a KeeperRegistered event raised by the ContractTaskExecutionHub contract.
type ContractTaskExecutionHubKeeperRegistered struct {
	Keeper common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterKeeperRegistered is a free log retrieval operation binding the contract event 0x1ecf9d662f201ac523808305cfcb9f1f2a8241b7fb3444333b5b4ee8f7e50b07.
//
// Solidity: event KeeperRegistered(address indexed keeper)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubFilterer) FilterKeeperRegistered(opts *bind.FilterOpts, keeper []common.Address) (*ContractTaskExecutionHubKeeperRegisteredIterator, error) {

	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}

	logs, sub, err := _ContractTaskExecutionHub.contract.FilterLogs(opts, "KeeperRegistered", keeperRule)
	if err != nil {
		return nil, err
	}
	return &ContractTaskExecutionHubKeeperRegisteredIterator{contract: _ContractTaskExecutionHub.contract, event: "KeeperRegistered", logs: logs, sub: sub}, nil
}

// WatchKeeperRegistered is a free log subscription operation binding the contract event 0x1ecf9d662f201ac523808305cfcb9f1f2a8241b7fb3444333b5b4ee8f7e50b07.
//
// Solidity: event KeeperRegistered(address indexed keeper)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubFilterer) WatchKeeperRegistered(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionHubKeeperRegistered, keeper []common.Address) (event.Subscription, error) {

	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}

	logs, sub, err := _ContractTaskExecutionHub.contract.WatchLogs(opts, "KeeperRegistered", keeperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTaskExecutionHubKeeperRegistered)
				if err := _ContractTaskExecutionHub.contract.UnpackLog(event, "KeeperRegistered", log); err != nil {
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

// ParseKeeperRegistered is a log parse operation binding the contract event 0x1ecf9d662f201ac523808305cfcb9f1f2a8241b7fb3444333b5b4ee8f7e50b07.
//
// Solidity: event KeeperRegistered(address indexed keeper)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubFilterer) ParseKeeperRegistered(log types.Log) (*ContractTaskExecutionHubKeeperRegistered, error) {
	event := new(ContractTaskExecutionHubKeeperRegistered)
	if err := _ContractTaskExecutionHub.contract.UnpackLog(event, "KeeperRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTaskExecutionHubKeeperUnregisteredIterator is returned from FilterKeeperUnregistered and is used to iterate over the raw logs and unpacked data for KeeperUnregistered events raised by the ContractTaskExecutionHub contract.
type ContractTaskExecutionHubKeeperUnregisteredIterator struct {
	Event *ContractTaskExecutionHubKeeperUnregistered // Event containing the contract specifics and raw log

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
func (it *ContractTaskExecutionHubKeeperUnregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTaskExecutionHubKeeperUnregistered)
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
		it.Event = new(ContractTaskExecutionHubKeeperUnregistered)
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
func (it *ContractTaskExecutionHubKeeperUnregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTaskExecutionHubKeeperUnregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTaskExecutionHubKeeperUnregistered represents a KeeperUnregistered event raised by the ContractTaskExecutionHub contract.
type ContractTaskExecutionHubKeeperUnregistered struct {
	Keeper common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterKeeperUnregistered is a free log retrieval operation binding the contract event 0x3d3b6e859e3dde2c83580d0bc85d58188abff18c72fccd0e167b04d70bcbf6cf.
//
// Solidity: event KeeperUnregistered(address indexed keeper)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubFilterer) FilterKeeperUnregistered(opts *bind.FilterOpts, keeper []common.Address) (*ContractTaskExecutionHubKeeperUnregisteredIterator, error) {

	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}

	logs, sub, err := _ContractTaskExecutionHub.contract.FilterLogs(opts, "KeeperUnregistered", keeperRule)
	if err != nil {
		return nil, err
	}
	return &ContractTaskExecutionHubKeeperUnregisteredIterator{contract: _ContractTaskExecutionHub.contract, event: "KeeperUnregistered", logs: logs, sub: sub}, nil
}

// WatchKeeperUnregistered is a free log subscription operation binding the contract event 0x3d3b6e859e3dde2c83580d0bc85d58188abff18c72fccd0e167b04d70bcbf6cf.
//
// Solidity: event KeeperUnregistered(address indexed keeper)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubFilterer) WatchKeeperUnregistered(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionHubKeeperUnregistered, keeper []common.Address) (event.Subscription, error) {

	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}

	logs, sub, err := _ContractTaskExecutionHub.contract.WatchLogs(opts, "KeeperUnregistered", keeperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTaskExecutionHubKeeperUnregistered)
				if err := _ContractTaskExecutionHub.contract.UnpackLog(event, "KeeperUnregistered", log); err != nil {
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

// ParseKeeperUnregistered is a log parse operation binding the contract event 0x3d3b6e859e3dde2c83580d0bc85d58188abff18c72fccd0e167b04d70bcbf6cf.
//
// Solidity: event KeeperUnregistered(address indexed keeper)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubFilterer) ParseKeeperUnregistered(log types.Log) (*ContractTaskExecutionHubKeeperUnregistered, error) {
	event := new(ContractTaskExecutionHubKeeperUnregistered)
	if err := _ContractTaskExecutionHub.contract.UnpackLog(event, "KeeperUnregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTaskExecutionHubLowBalanceAlertIterator is returned from FilterLowBalanceAlert and is used to iterate over the raw logs and unpacked data for LowBalanceAlert events raised by the ContractTaskExecutionHub contract.
type ContractTaskExecutionHubLowBalanceAlertIterator struct {
	Event *ContractTaskExecutionHubLowBalanceAlert // Event containing the contract specifics and raw log

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
func (it *ContractTaskExecutionHubLowBalanceAlertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTaskExecutionHubLowBalanceAlert)
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
		it.Event = new(ContractTaskExecutionHubLowBalanceAlert)
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
func (it *ContractTaskExecutionHubLowBalanceAlertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTaskExecutionHubLowBalanceAlertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTaskExecutionHubLowBalanceAlert represents a LowBalanceAlert event raised by the ContractTaskExecutionHub contract.
type ContractTaskExecutionHubLowBalanceAlert struct {
	CurrentBalance *big.Int
	Threshold      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLowBalanceAlert is a free log retrieval operation binding the contract event 0xb1b538a145d528cb9b8a428a3f47d5730430a03c2d98d7562357ee7db16dbd26.
//
// Solidity: event LowBalanceAlert(uint256 currentBalance, uint256 threshold)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubFilterer) FilterLowBalanceAlert(opts *bind.FilterOpts) (*ContractTaskExecutionHubLowBalanceAlertIterator, error) {

	logs, sub, err := _ContractTaskExecutionHub.contract.FilterLogs(opts, "LowBalanceAlert")
	if err != nil {
		return nil, err
	}
	return &ContractTaskExecutionHubLowBalanceAlertIterator{contract: _ContractTaskExecutionHub.contract, event: "LowBalanceAlert", logs: logs, sub: sub}, nil
}

// WatchLowBalanceAlert is a free log subscription operation binding the contract event 0xb1b538a145d528cb9b8a428a3f47d5730430a03c2d98d7562357ee7db16dbd26.
//
// Solidity: event LowBalanceAlert(uint256 currentBalance, uint256 threshold)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubFilterer) WatchLowBalanceAlert(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionHubLowBalanceAlert) (event.Subscription, error) {

	logs, sub, err := _ContractTaskExecutionHub.contract.WatchLogs(opts, "LowBalanceAlert")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTaskExecutionHubLowBalanceAlert)
				if err := _ContractTaskExecutionHub.contract.UnpackLog(event, "LowBalanceAlert", log); err != nil {
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

// ParseLowBalanceAlert is a log parse operation binding the contract event 0xb1b538a145d528cb9b8a428a3f47d5730430a03c2d98d7562357ee7db16dbd26.
//
// Solidity: event LowBalanceAlert(uint256 currentBalance, uint256 threshold)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubFilterer) ParseLowBalanceAlert(log types.Log) (*ContractTaskExecutionHubLowBalanceAlert, error) {
	event := new(ContractTaskExecutionHubLowBalanceAlert)
	if err := _ContractTaskExecutionHub.contract.UnpackLog(event, "LowBalanceAlert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTaskExecutionHubMessageFailedIterator is returned from FilterMessageFailed and is used to iterate over the raw logs and unpacked data for MessageFailed events raised by the ContractTaskExecutionHub contract.
type ContractTaskExecutionHubMessageFailedIterator struct {
	Event *ContractTaskExecutionHubMessageFailed // Event containing the contract specifics and raw log

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
func (it *ContractTaskExecutionHubMessageFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTaskExecutionHubMessageFailed)
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
		it.Event = new(ContractTaskExecutionHubMessageFailed)
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
func (it *ContractTaskExecutionHubMessageFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTaskExecutionHubMessageFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTaskExecutionHubMessageFailed represents a MessageFailed event raised by the ContractTaskExecutionHub contract.
type ContractTaskExecutionHubMessageFailed struct {
	DstEid uint32
	Guid   [32]byte
	Reason []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMessageFailed is a free log retrieval operation binding the contract event 0x2254f63be429412a67ef70cf2ef65d61a88074bfd112333e0a449a9e4a7683c8.
//
// Solidity: event MessageFailed(uint32 indexed dstEid, bytes32 indexed guid, bytes reason)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubFilterer) FilterMessageFailed(opts *bind.FilterOpts, dstEid []uint32, guid [][32]byte) (*ContractTaskExecutionHubMessageFailedIterator, error) {

	var dstEidRule []interface{}
	for _, dstEidItem := range dstEid {
		dstEidRule = append(dstEidRule, dstEidItem)
	}
	var guidRule []interface{}
	for _, guidItem := range guid {
		guidRule = append(guidRule, guidItem)
	}

	logs, sub, err := _ContractTaskExecutionHub.contract.FilterLogs(opts, "MessageFailed", dstEidRule, guidRule)
	if err != nil {
		return nil, err
	}
	return &ContractTaskExecutionHubMessageFailedIterator{contract: _ContractTaskExecutionHub.contract, event: "MessageFailed", logs: logs, sub: sub}, nil
}

// WatchMessageFailed is a free log subscription operation binding the contract event 0x2254f63be429412a67ef70cf2ef65d61a88074bfd112333e0a449a9e4a7683c8.
//
// Solidity: event MessageFailed(uint32 indexed dstEid, bytes32 indexed guid, bytes reason)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubFilterer) WatchMessageFailed(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionHubMessageFailed, dstEid []uint32, guid [][32]byte) (event.Subscription, error) {

	var dstEidRule []interface{}
	for _, dstEidItem := range dstEid {
		dstEidRule = append(dstEidRule, dstEidItem)
	}
	var guidRule []interface{}
	for _, guidItem := range guid {
		guidRule = append(guidRule, guidItem)
	}

	logs, sub, err := _ContractTaskExecutionHub.contract.WatchLogs(opts, "MessageFailed", dstEidRule, guidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTaskExecutionHubMessageFailed)
				if err := _ContractTaskExecutionHub.contract.UnpackLog(event, "MessageFailed", log); err != nil {
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

// ParseMessageFailed is a log parse operation binding the contract event 0x2254f63be429412a67ef70cf2ef65d61a88074bfd112333e0a449a9e4a7683c8.
//
// Solidity: event MessageFailed(uint32 indexed dstEid, bytes32 indexed guid, bytes reason)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubFilterer) ParseMessageFailed(log types.Log) (*ContractTaskExecutionHubMessageFailed, error) {
	event := new(ContractTaskExecutionHubMessageFailed)
	if err := _ContractTaskExecutionHub.contract.UnpackLog(event, "MessageFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTaskExecutionHubOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractTaskExecutionHub contract.
type ContractTaskExecutionHubOwnershipTransferredIterator struct {
	Event *ContractTaskExecutionHubOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractTaskExecutionHubOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTaskExecutionHubOwnershipTransferred)
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
		it.Event = new(ContractTaskExecutionHubOwnershipTransferred)
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
func (it *ContractTaskExecutionHubOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTaskExecutionHubOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTaskExecutionHubOwnershipTransferred represents a OwnershipTransferred event raised by the ContractTaskExecutionHub contract.
type ContractTaskExecutionHubOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractTaskExecutionHubOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractTaskExecutionHub.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractTaskExecutionHubOwnershipTransferredIterator{contract: _ContractTaskExecutionHub.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionHubOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractTaskExecutionHub.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTaskExecutionHubOwnershipTransferred)
				if err := _ContractTaskExecutionHub.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractTaskExecutionHub *ContractTaskExecutionHubFilterer) ParseOwnershipTransferred(log types.Log) (*ContractTaskExecutionHubOwnershipTransferred, error) {
	event := new(ContractTaskExecutionHubOwnershipTransferred)
	if err := _ContractTaskExecutionHub.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTaskExecutionHubPeerSetIterator is returned from FilterPeerSet and is used to iterate over the raw logs and unpacked data for PeerSet events raised by the ContractTaskExecutionHub contract.
type ContractTaskExecutionHubPeerSetIterator struct {
	Event *ContractTaskExecutionHubPeerSet // Event containing the contract specifics and raw log

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
func (it *ContractTaskExecutionHubPeerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTaskExecutionHubPeerSet)
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
		it.Event = new(ContractTaskExecutionHubPeerSet)
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
func (it *ContractTaskExecutionHubPeerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTaskExecutionHubPeerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTaskExecutionHubPeerSet represents a PeerSet event raised by the ContractTaskExecutionHub contract.
type ContractTaskExecutionHubPeerSet struct {
	Eid  uint32
	Peer [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPeerSet is a free log retrieval operation binding the contract event 0x238399d427b947898edb290f5ff0f9109849b1c3ba196a42e35f00c50a54b98b.
//
// Solidity: event PeerSet(uint32 eid, bytes32 peer)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubFilterer) FilterPeerSet(opts *bind.FilterOpts) (*ContractTaskExecutionHubPeerSetIterator, error) {

	logs, sub, err := _ContractTaskExecutionHub.contract.FilterLogs(opts, "PeerSet")
	if err != nil {
		return nil, err
	}
	return &ContractTaskExecutionHubPeerSetIterator{contract: _ContractTaskExecutionHub.contract, event: "PeerSet", logs: logs, sub: sub}, nil
}

// WatchPeerSet is a free log subscription operation binding the contract event 0x238399d427b947898edb290f5ff0f9109849b1c3ba196a42e35f00c50a54b98b.
//
// Solidity: event PeerSet(uint32 eid, bytes32 peer)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubFilterer) WatchPeerSet(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionHubPeerSet) (event.Subscription, error) {

	logs, sub, err := _ContractTaskExecutionHub.contract.WatchLogs(opts, "PeerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTaskExecutionHubPeerSet)
				if err := _ContractTaskExecutionHub.contract.UnpackLog(event, "PeerSet", log); err != nil {
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
func (_ContractTaskExecutionHub *ContractTaskExecutionHubFilterer) ParsePeerSet(log types.Log) (*ContractTaskExecutionHubPeerSet, error) {
	event := new(ContractTaskExecutionHubPeerSet)
	if err := _ContractTaskExecutionHub.contract.UnpackLog(event, "PeerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTaskExecutionHubUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ContractTaskExecutionHub contract.
type ContractTaskExecutionHubUpgradedIterator struct {
	Event *ContractTaskExecutionHubUpgraded // Event containing the contract specifics and raw log

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
func (it *ContractTaskExecutionHubUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTaskExecutionHubUpgraded)
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
		it.Event = new(ContractTaskExecutionHubUpgraded)
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
func (it *ContractTaskExecutionHubUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTaskExecutionHubUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTaskExecutionHubUpgraded represents a Upgraded event raised by the ContractTaskExecutionHub contract.
type ContractTaskExecutionHubUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ContractTaskExecutionHubUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ContractTaskExecutionHub.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ContractTaskExecutionHubUpgradedIterator{contract: _ContractTaskExecutionHub.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ContractTaskExecutionHub *ContractTaskExecutionHubFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionHubUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ContractTaskExecutionHub.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTaskExecutionHubUpgraded)
				if err := _ContractTaskExecutionHub.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_ContractTaskExecutionHub *ContractTaskExecutionHubFilterer) ParseUpgraded(log types.Log) (*ContractTaskExecutionHubUpgraded, error) {
	event := new(ContractTaskExecutionHubUpgraded)
	if err := _ContractTaskExecutionHub.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
