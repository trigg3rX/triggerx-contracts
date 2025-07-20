// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractAVSGovernanceLogic

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

// ContractAVSGovernanceLogicMetaData contains all meta data concerning the ContractAVSGovernanceLogic contract.
var ContractAVSGovernanceLogicMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_endpoint\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_taskExecutionHub\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_dstEid\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_ownerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_avsGovernance\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidDelegate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEndpointCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LzTokenUnavailable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"}],\"name\":\"NoPeer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"}],\"name\":\"NotEnoughNative\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"OnlyEndpoint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"}],\"name\":\"OnlyPeer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"gasLimit\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"callValue\",\"type\":\"uint128\"}],\"name\":\"GasOptionsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"LowBalanceAlert\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"MessageFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"MessageSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"peer\",\"type\":\"bytes32\"}],\"name\":\"PeerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"UnWhitelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"Whitelisted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_operators\",\"type\":\"address[]\"}],\"name\":\"addToWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"\",\"type\":\"uint256[4]\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"afterOperatorRegistered\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"afterOperatorUnregistered\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"origin\",\"type\":\"tuple\"}],\"name\":\"allowInitializePath\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"avsGovernance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"\",\"type\":\"uint256[4]\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"beforeOperatorRegistered\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"beforeOperatorUnregistered\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"callValue\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dstEid\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endpoint\",\"outputs\":[{\"internalType\":\"contractILayerZeroEndpointV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasLimit\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"isComposeMsgSender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"_origin\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_guid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"lzReceive\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"nextNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oAppVersion\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"senderVersion\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"receiverVersion\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"}],\"name\":\"peers\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"peer\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taskExecutionHub\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_operators\",\"type\":\"address[]\"}],\"name\":\"removeFromWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegate\",\"type\":\"address\"}],\"name\":\"setDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_gasLimit\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_callValue\",\"type\":\"uint128\"}],\"name\":\"setGasOptions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_peer\",\"type\":\"bytes32\"}],\"name\":\"setPeer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_taskExecutionHub\",\"type\":\"address\"}],\"name\":\"setTaskExecutionHub\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052600436101561001a575b3615610018575f80fd5b005b5f5f3560e01c806313137d6514610f7c57806317442b7014610f5b5780633400288b14610f2e5780633af32abf14610ef1578063471f9e6414610ec2578063548db17414610dfb5780635e280f1114610db7578063715018a614610d60578063734ff5f314610d205780637d25a05e14610cfc5780637f64978314610bee57806382413eac14610b915780638d98e57914610b695780638da5cb5b14610b4257806394e7a3f014610a6b57806396e51b9c146109ba5780639dbe15381461092b578063bb0b6a53146108f6578063ca5eb5e11461084e578063d62457f61461082d578063e9ecc1cb14610631578063ed8ad83614610380578063f0b351d914610357578063f2fde38b146102d1578063f3fef3a3146101b1578063f68016b71461018a5763ff7bd03d1461014e575061000e565b346101875760603660031901126101875760209060409063ffffffff6101726111bf565b16815260018352205460405190602435148152f35b80fd5b503461018757806003193601126101875760206001600160801b0360045416604051908152f35b5034610187576040366003190112610187576004356001600160a01b038116908190036102cd57602435906101e46114a8565b801561029457811561024f578147106102135782808093819382f1156102075780f35b604051903d90823e3d90fd5b60405162461bcd60e51b8152602060048201526014602482015273496e73756666696369656e742062616c616e636560601b6044820152606490fd5b60405162461bcd60e51b815260206004820152601d60248201527f416d6f756e74206d7573742062652067726561746572207468616e20300000006044820152606490fd5b60405162461bcd60e51b8152602060048201526011602482015270125b9d985b1a59081c9958da5c1a595b9d607a1b6044820152606490fd5b5080fd5b5034610187576020366003190112610187576102eb6110fe565b6102f36114a8565b6001600160a01b031680156103435781546001600160a01b03198116821783556001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b631e4fbdf760e01b82526004829052602482fd5b50346101875780600319360112610187576002546040516001600160a01b039091168152602090f35b50346101875761038f3661117b565b5050506103a760018060a01b036003541633146111d2565b6040519082602083015260018060a01b0316806040830152604082526103ce6060836112d1565b6104616103ea6004546001600160801b038160801c9116611514565b7f0000000000000000000000000000000000000000000000000000000000009d359060018060a01b0360025416946040805161042581611269565b63ffffffff85169788825260208201528282820152836060820152886080820152815180968192631bb8518b60e31b835230906004840161133f565b03817f0000000000000000000000006edce65403992e310a62460808c4b910d972f10f6001600160a01b03165afa8794816105f2575b5061054c5761054860206105286031898b806104b1611439565b927f2254f63be429412a67ef70cf2ef65d61a88074bfd112333e0a449a9e4a7683c8604051888152806104e68a82018861131b565b0390a36040519485927002932b3b4b9ba32b9103330b4b632b21d1607d1b828501528051918291018585015e820190838201520301601f1981018352826112d1565b60405162461bcd60e51b815260206004820152918291602483019061131b565b0390fd5b8394959351600a8102818104600a14821517156105de576020936105ad936105a661059e7f0597c8258e620343bfa7f3cf82a7d6578378f08248fd958f32bcab61a1004ec3989560648c9604906113ad565b4710156113ce565b30936115f3565b519351604051908152a37f4d0eb1f4bac8744fd2be119845e23b3befc88094b42bcda1204c65694a00f9e58280a280f35b634e487b7160e01b89526011600452602489fd5b90945060403d60401161062a575b61060a81836112d1565b81016040828203126106265761061f916112f3565b935f610497565b8880fd5b503d610600565b50346101875760203660031901126101875761064b6110fe565b61066060018060a01b036003541633146111d2565b60408051600160208201526001600160a01b0390921682820181905290825261068a6060836112d1565b6106a66103ea6004546001600160801b038160801c9116611514565b03817f0000000000000000000000006edce65403992e310a62460808c4b910d972f10f6001600160a01b03165afa8794816107f2575b5061076f5761054860206105286033898b806106f6611439565b927f2254f63be429412a67ef70cf2ef65d61a88074bfd112333e0a449a9e4a7683c86040518881528061072b8a82018861131b565b0390a36040519485927202ab73932b3b4b9ba32b9103330b4b632b21d1606d1b828501528051918291018585015e820190838201520301601f1981018352826112d1565b8394959351600a8102818104600a14821517156105de576020936107c1936105a661059e7f0597c8258e620343bfa7f3cf82a7d6578378f08248fd958f32bcab61a1004ec3989560648c9604906113ad565b519351604051908152a37f6f42117a557500c705ddf040a619d86f39101e6b74ac20d7b3e5943ba473fc7f8280a280f35b90945060403d604011610826575b61080a81836112d1565b81016040828203126106265761081f916112f3565b935f6106dc565b503d610800565b5034610187578060031936011261018757602060045460801c604051908152f35b50346108f25760203660031901126108f2576108686110fe565b6108706114a8565b7f0000000000000000000000006edce65403992e310a62460808c4b910d972f10f6001600160a01b031690813b156108f25760405163ca5eb5e160e01b81526001600160a01b039091166004820152905f908290602490829084905af180156108e7576108db575080f35b61001891505f906112d1565b6040513d5f823e3d90fd5b5f80fd5b346108f25760203660031901126108f25763ffffffff610914611114565b165f526001602052602060405f2054604051908152f35b346108f25760403660031901126108f2576004356001600160801b0381168091036108f2576024356001600160801b038116918282036108f2577f87768cdb9e7a69da85a21ad47c9654958b5cbc0447a55b66d0e3e0d640430c2d926040926109926114a8565b60801b6fffffffffffffffffffffffffffffffff1916821760045582519182526020820152a1005b346108f25760203660031901126108f2576109d36110fe565b6109db6114a8565b6001600160a01b03168015610a2657600280546001600160a01b03191682179055610018907f0000000000000000000000000000000000000000000000000000000000009d356114ce565b60405162461bcd60e51b815260206004820152601960248201527f496e76616c69642070726f7879206875622061646472657373000000000000006044820152606490fd5b346108f257610a793661117b565b505050610a9160018060a01b036003541633146111d2565b662386f26fc100004710610b06575b6001600160a01b03165f9081526005602052604090205460ff1615610ac157005b60405162461bcd60e51b815260206004820152601b60248201527f4f70657261746f72206973206e6f742077686974656c697374656400000000006044820152606490fd5b7fb1b538a145d528cb9b8a428a3f47d5730430a03c2d98d7562357ee7db16dbd266040478151908152662386f26fc100006020820152a1610aa0565b346108f2575f3660031901126108f2575f546040516001600160a01b039091168152602090f35b346108f2575f3660031901126108f2576003546040516001600160a01b039091168152602090f35b346108f257366003190160a081126108f2576060136108f25760643567ffffffffffffffff81116108f257610bca9036906004016110ba565b50506084356001600160a01b038116908190036108f2576020906040519030148152f35b346108f257610bfc36611127565b90610c056114a8565b5f5b828110610c1057005b6001600160a01b03610c2b610c26838686611231565b611255565b16908115610cc557815f52600560205260ff60405f205416610c8a57816001925f52600560205260405f208360ff198254161790557faab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a545f80a201610c07565b60405162461bcd60e51b8152602060048201526013602482015272105b1c9958591e481dda1a5d195b1a5cdd1959606a1b6044820152606490fd5b60405162461bcd60e51b815260206004820152600f60248201526e496e76616c6964206164647265737360881b6044820152606490fd5b346108f25760403660031901126108f257610d15611114565b5060206040515f8152f35b346108f2575f3660031901126108f257602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000009d35168152f35b346108f2575f3660031901126108f257610d786114a8565b5f80546001600160a01b0319811682556001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b346108f2575f3660031901126108f2576040517f0000000000000000000000006edce65403992e310a62460808c4b910d972f10f6001600160a01b03168152602090f35b346108f257610e0936611127565b90610e126114a8565b5f5b828110610e1d57005b6001600160a01b03610e33610c26838686611231565b1690815f52600560205260ff60405f20541615610e8b57816001925f52600560205260405f2060ff1981541690557f7cdb51b0cc2e541ad7e9471c358de415f5dbaff6cca78e3393d445346c610c1b5f80a201610e14565b60405162461bcd60e51b815260206004820152600f60248201526e139bdd081dda1a5d195b1a5cdd1959608a1b6044820152606490fd5b346108f25760203660031901126108f257610edb6110fe565b5061001860018060a01b036003541633146111d2565b346108f25760203660031901126108f2576001600160a01b03610f126110fe565b165f526005602052602060ff60405f2054166040519015158152f35b346108f25760403660031901126108f257610018610f4a611114565b610f526114a8565b602435906114ce565b346108f2575f3660031901126108f257604080516001815260026020820152f35b366003190160e081126108f2576060136108f25760843567ffffffffffffffff81116108f257610fb09036906004016110ba565b5050610fba6110e8565b5060c43567ffffffffffffffff81116108f257610fdb9036906004016110ba565b5050337f0000000000000000000000006edce65403992e310a62460808c4b910d972f10f6001600160a01b0316036110a75761101d6110186111bf565b611478565b6024358091036110845760405162461bcd60e51b815260206004820152602f60248201527f417673476f7665726e616e63654c6f6769633a2073686f756c64206e6f74207260448201526e656365697665206d6573736167657360881b6064820152608490fd5b63ffffffff6110916111bf565b63309afaf360e21b5f521660045260245260445ffd5b6391ac5e4f60e01b5f523360045260245ffd5b9181601f840112156108f25782359167ffffffffffffffff83116108f257602083818601950101116108f257565b60a435906001600160a01b03821682036108f257565b600435906001600160a01b03821682036108f257565b6004359063ffffffff821682036108f257565b9060206003198301126108f25760043567ffffffffffffffff81116108f257826023820112156108f25780600401359267ffffffffffffffff84116108f25760248460051b830101116108f2576024019190565b60e06003198201126108f2576004356001600160a01b03811681036108f257916024359160c4116108f25760449060c4356001600160a01b03811681036108f25790565b60043563ffffffff811681036108f25790565b156111d957565b60405162461bcd60e51b815260206004820152602a60248201527f4f6e6c792041565320476f7665726e616e63652063616e2063616c6c207468696044820152693990333ab731ba34b7b760b11b6064820152608490fd5b91908110156112415760051b0190565b634e487b7160e01b5f52603260045260245ffd5b356001600160a01b03811681036108f25790565b60a0810190811067ffffffffffffffff82111761128557604052565b634e487b7160e01b5f52604160045260245ffd5b6040810190811067ffffffffffffffff82111761128557604052565b6060810190811067ffffffffffffffff82111761128557604052565b90601f8019910116810190811067ffffffffffffffff82111761128557604052565b91908260409103126108f25760405161130b81611299565b6020808294805184520151910152565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b906020909392936040835263ffffffff8151166040840152818101516060840152608061139361137d604084015160a08488015260e087019061131b565b6060840151868203603f190160a088015261131b565b910151151560c08401526001600160a01b03909416910152565b919082018092116113ba57565b634e487b7160e01b5f52601160045260245ffd5b156113d557565b60405162461bcd60e51b815260206004820152603660248201527f496e73756666696369656e742062616c616e636520666f72206d6573736167656044820152752066656520287769746820313025206275666665722960501b6064820152608490fd5b3d15611473573d9067ffffffffffffffff82116112855760405191611468601f8201601f1916602001846112d1565b82523d5f602084013e565b606090565b63ffffffff16805f52600160205260405f2054908115611496575090565b63f6ff4fb760e01b5f5260045260245ffd5b5f546001600160a01b031633036114bb57565b63118cdaa760e01b5f523360045260245ffd5b7f238399d427b947898edb290f5ff0f9109849b1c3ba196a42e35f00c50a54b98b9163ffffffff6040921690815f52600160205280835f205582519182526020820152a1565b906001600160801b0381166115b55750604051906001600160801b03199060801b1660208201526010815261154a6030826112d1565b8051600181018091116113ba5760266115b2916020604051948592600360f01b83850152600160f81b602285015261ffff60f01b9060f01b166023840152600160f81b60258401528051918291018484015e81015f838201520301601f1981018352826112d1565b90565b604051916001600160801b03199060801b1660208301526001600160801b03199060801b166030820152602081526115ee6040826112d1565b61154a565b91604051611600816112b5565b5f81525f6020820152604080519161161783611299565b5f83525f6020840152015283519384471061185e5760200191825180611740575b509183916080959361164d6116969896611478565b925115159263ffffffff6040519561166487611269565b168552602085015260408401526060830152848201526040518095819482936302637a4560e41b84526004840161133f565b03917f0000000000000000000000006edce65403992e310a62460808c4b910d972f10f6001600160a01b03165af19081156108e7575f916116d5575090565b905060803d608011611739575b6116ec81836112d1565b8101906080818303126108f25760405191611706836112b5565b81518352602082015167ffffffffffffffff811681036108f2576020840152611731916040016112f3565b604082015290565b503d6116e2565b60405163393f876560e21b81527f0000000000000000000000006edce65403992e310a62460808c4b910d972f10f6001600160a01b031690602081600481855afa9081156108e7575f9161181c575b506001600160a01b031691821561180d576020915f9160405190848201926323b872dd60e01b845233602484015260448301526064820152606481526117d66084826112d1565b519082855af1156108e7575f513d6118045750803b155b1561163857635274afe760e01b5f5260045260245ffd5b600114156117ed565b6329b99a9560e11b5f5260045ffd5b90506020813d602011611856575b81611837602093836112d1565b810103126108f257516001600160a01b03811681036108f2575f61178f565b3d915061182a565b60405162461bcd60e51b815260206004820152601d60248201527f496e73756666696369656e7420636f6e74726163742062616c616e63650000006044820152606490fdfea2646970667358221220302ee99f6d2dd4aee8c22b3650fcd1c9accc6258a22c4ee4b1877048c1e458b164736f6c634300081b0033",
}

// ContractAVSGovernanceLogicABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractAVSGovernanceLogicMetaData.ABI instead.
var ContractAVSGovernanceLogicABI = ContractAVSGovernanceLogicMetaData.ABI

// ContractAVSGovernanceLogicBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractAVSGovernanceLogicMetaData.Bin instead.
var ContractAVSGovernanceLogicBin = ContractAVSGovernanceLogicMetaData.Bin

// DeployContractAVSGovernanceLogic deploys a new Ethereum contract, binding an instance of ContractAVSGovernanceLogic to it.
func DeployContractAVSGovernanceLogic(auth *bind.TransactOpts, backend bind.ContractBackend, _endpoint common.Address, _taskExecutionHub common.Address, _dstEid uint32, _ownerAddress common.Address, _avsGovernance common.Address) (common.Address, *types.Transaction, *ContractAVSGovernanceLogic, error) {
	parsed, err := ContractAVSGovernanceLogicMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractAVSGovernanceLogicBin), backend, _endpoint, _taskExecutionHub, _dstEid, _ownerAddress, _avsGovernance)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractAVSGovernanceLogic{ContractAVSGovernanceLogicCaller: ContractAVSGovernanceLogicCaller{contract: contract}, ContractAVSGovernanceLogicTransactor: ContractAVSGovernanceLogicTransactor{contract: contract}, ContractAVSGovernanceLogicFilterer: ContractAVSGovernanceLogicFilterer{contract: contract}}, nil
}

// ContractAVSGovernanceLogicMethods is an auto generated interface around an Ethereum contract.
type ContractAVSGovernanceLogicMethods interface {
	ContractAVSGovernanceLogicCalls
	ContractAVSGovernanceLogicTransacts
	ContractAVSGovernanceLogicFilters
}

// ContractAVSGovernanceLogicCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractAVSGovernanceLogicCalls interface {
	AllowInitializePath(opts *bind.CallOpts, origin Origin) (bool, error)

	AvsGovernance(opts *bind.CallOpts) (common.Address, error)

	CallValue(opts *bind.CallOpts) (*big.Int, error)

	DstEid(opts *bind.CallOpts) (uint32, error)

	Endpoint(opts *bind.CallOpts) (common.Address, error)

	GasLimit(opts *bind.CallOpts) (*big.Int, error)

	IsComposeMsgSender(opts *bind.CallOpts, arg0 Origin, arg1 []byte, _sender common.Address) (bool, error)

	IsWhitelisted(opts *bind.CallOpts, arg0 common.Address) (bool, error)

	NextNonce(opts *bind.CallOpts, arg0 uint32, arg1 [32]byte) (uint64, error)

	OAppVersion(opts *bind.CallOpts) (struct {
		SenderVersion   uint64
		ReceiverVersion uint64
	}, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Peers(opts *bind.CallOpts, eid uint32) ([32]byte, error)

	TaskExecutionHub(opts *bind.CallOpts) (common.Address, error)
}

// ContractAVSGovernanceLogicTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractAVSGovernanceLogicTransacts interface {
	AddToWhitelist(opts *bind.TransactOpts, _operators []common.Address) (*types.Transaction, error)

	AfterOperatorRegistered(opts *bind.TransactOpts, _operator common.Address, arg1 *big.Int, arg2 [4]*big.Int, arg3 common.Address) (*types.Transaction, error)

	AfterOperatorUnregistered(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error)

	BeforeOperatorRegistered(opts *bind.TransactOpts, _operator common.Address, arg1 *big.Int, arg2 [4]*big.Int, arg3 common.Address) (*types.Transaction, error)

	BeforeOperatorUnregistered(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error)

	LzReceive(opts *bind.TransactOpts, _origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error)

	RemoveFromWhitelist(opts *bind.TransactOpts, _operators []common.Address) (*types.Transaction, error)

	RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetDelegate(opts *bind.TransactOpts, _delegate common.Address) (*types.Transaction, error)

	SetGasOptions(opts *bind.TransactOpts, _gasLimit *big.Int, _callValue *big.Int) (*types.Transaction, error)

	SetPeer(opts *bind.TransactOpts, _eid uint32, _peer [32]byte) (*types.Transaction, error)

	SetTaskExecutionHub(opts *bind.TransactOpts, _taskExecutionHub common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error)

	Withdraw(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error)
}

// ContractAVSGovernanceLogicFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractAVSGovernanceLogicFilters interface {
	FilterGasOptionsUpdated(opts *bind.FilterOpts) (*ContractAVSGovernanceLogicGasOptionsUpdatedIterator, error)
	WatchGasOptionsUpdated(opts *bind.WatchOpts, sink chan<- *ContractAVSGovernanceLogicGasOptionsUpdated) (event.Subscription, error)
	ParseGasOptionsUpdated(log types.Log) (*ContractAVSGovernanceLogicGasOptionsUpdated, error)

	FilterLowBalanceAlert(opts *bind.FilterOpts) (*ContractAVSGovernanceLogicLowBalanceAlertIterator, error)
	WatchLowBalanceAlert(opts *bind.WatchOpts, sink chan<- *ContractAVSGovernanceLogicLowBalanceAlert) (event.Subscription, error)
	ParseLowBalanceAlert(log types.Log) (*ContractAVSGovernanceLogicLowBalanceAlert, error)

	FilterMessageFailed(opts *bind.FilterOpts, dstEid []uint32, guid [][32]byte) (*ContractAVSGovernanceLogicMessageFailedIterator, error)
	WatchMessageFailed(opts *bind.WatchOpts, sink chan<- *ContractAVSGovernanceLogicMessageFailed, dstEid []uint32, guid [][32]byte) (event.Subscription, error)
	ParseMessageFailed(log types.Log) (*ContractAVSGovernanceLogicMessageFailed, error)

	FilterMessageSent(opts *bind.FilterOpts, dstEid []uint32, guid [][32]byte) (*ContractAVSGovernanceLogicMessageSentIterator, error)
	WatchMessageSent(opts *bind.WatchOpts, sink chan<- *ContractAVSGovernanceLogicMessageSent, dstEid []uint32, guid [][32]byte) (event.Subscription, error)
	ParseMessageSent(log types.Log) (*ContractAVSGovernanceLogicMessageSent, error)

	FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address) (*ContractAVSGovernanceLogicOperatorRegisteredIterator, error)
	WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *ContractAVSGovernanceLogicOperatorRegistered, operator []common.Address) (event.Subscription, error)
	ParseOperatorRegistered(log types.Log) (*ContractAVSGovernanceLogicOperatorRegistered, error)

	FilterOperatorUnregistered(opts *bind.FilterOpts, operator []common.Address) (*ContractAVSGovernanceLogicOperatorUnregisteredIterator, error)
	WatchOperatorUnregistered(opts *bind.WatchOpts, sink chan<- *ContractAVSGovernanceLogicOperatorUnregistered, operator []common.Address) (event.Subscription, error)
	ParseOperatorUnregistered(log types.Log) (*ContractAVSGovernanceLogicOperatorUnregistered, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractAVSGovernanceLogicOwnershipTransferredIterator, error)
	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractAVSGovernanceLogicOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error)
	ParseOwnershipTransferred(log types.Log) (*ContractAVSGovernanceLogicOwnershipTransferred, error)

	FilterPeerSet(opts *bind.FilterOpts) (*ContractAVSGovernanceLogicPeerSetIterator, error)
	WatchPeerSet(opts *bind.WatchOpts, sink chan<- *ContractAVSGovernanceLogicPeerSet) (event.Subscription, error)
	ParsePeerSet(log types.Log) (*ContractAVSGovernanceLogicPeerSet, error)

	FilterUnWhitelisted(opts *bind.FilterOpts, operator []common.Address) (*ContractAVSGovernanceLogicUnWhitelistedIterator, error)
	WatchUnWhitelisted(opts *bind.WatchOpts, sink chan<- *ContractAVSGovernanceLogicUnWhitelisted, operator []common.Address) (event.Subscription, error)
	ParseUnWhitelisted(log types.Log) (*ContractAVSGovernanceLogicUnWhitelisted, error)

	FilterWhitelisted(opts *bind.FilterOpts, operator []common.Address) (*ContractAVSGovernanceLogicWhitelistedIterator, error)
	WatchWhitelisted(opts *bind.WatchOpts, sink chan<- *ContractAVSGovernanceLogicWhitelisted, operator []common.Address) (event.Subscription, error)
	ParseWhitelisted(log types.Log) (*ContractAVSGovernanceLogicWhitelisted, error)
}

// ContractAVSGovernanceLogic is an auto generated Go binding around an Ethereum contract.
type ContractAVSGovernanceLogic struct {
	ContractAVSGovernanceLogicCaller     // Read-only binding to the contract
	ContractAVSGovernanceLogicTransactor // Write-only binding to the contract
	ContractAVSGovernanceLogicFilterer   // Log filterer for contract events
}

// ContractAVSGovernanceLogic implements the ContractAVSGovernanceLogicMethods interface.
var _ ContractAVSGovernanceLogicMethods = (*ContractAVSGovernanceLogic)(nil)

// ContractAVSGovernanceLogicCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractAVSGovernanceLogicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAVSGovernanceLogicCaller implements the ContractAVSGovernanceLogicCalls interface.
var _ ContractAVSGovernanceLogicCalls = (*ContractAVSGovernanceLogicCaller)(nil)

// ContractAVSGovernanceLogicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractAVSGovernanceLogicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAVSGovernanceLogicTransactor implements the ContractAVSGovernanceLogicTransacts interface.
var _ ContractAVSGovernanceLogicTransacts = (*ContractAVSGovernanceLogicTransactor)(nil)

// ContractAVSGovernanceLogicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractAVSGovernanceLogicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAVSGovernanceLogicFilterer implements the ContractAVSGovernanceLogicFilters interface.
var _ ContractAVSGovernanceLogicFilters = (*ContractAVSGovernanceLogicFilterer)(nil)

// ContractAVSGovernanceLogicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractAVSGovernanceLogicSession struct {
	Contract     *ContractAVSGovernanceLogic // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ContractAVSGovernanceLogicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractAVSGovernanceLogicCallerSession struct {
	Contract *ContractAVSGovernanceLogicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// ContractAVSGovernanceLogicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractAVSGovernanceLogicTransactorSession struct {
	Contract     *ContractAVSGovernanceLogicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// ContractAVSGovernanceLogicRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractAVSGovernanceLogicRaw struct {
	Contract *ContractAVSGovernanceLogic // Generic contract binding to access the raw methods on
}

// ContractAVSGovernanceLogicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractAVSGovernanceLogicCallerRaw struct {
	Contract *ContractAVSGovernanceLogicCaller // Generic read-only contract binding to access the raw methods on
}

// ContractAVSGovernanceLogicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractAVSGovernanceLogicTransactorRaw struct {
	Contract *ContractAVSGovernanceLogicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractAVSGovernanceLogic creates a new instance of ContractAVSGovernanceLogic, bound to a specific deployed contract.
func NewContractAVSGovernanceLogic(address common.Address, backend bind.ContractBackend) (*ContractAVSGovernanceLogic, error) {
	contract, err := bindContractAVSGovernanceLogic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractAVSGovernanceLogic{ContractAVSGovernanceLogicCaller: ContractAVSGovernanceLogicCaller{contract: contract}, ContractAVSGovernanceLogicTransactor: ContractAVSGovernanceLogicTransactor{contract: contract}, ContractAVSGovernanceLogicFilterer: ContractAVSGovernanceLogicFilterer{contract: contract}}, nil
}

// NewContractAVSGovernanceLogicCaller creates a new read-only instance of ContractAVSGovernanceLogic, bound to a specific deployed contract.
func NewContractAVSGovernanceLogicCaller(address common.Address, caller bind.ContractCaller) (*ContractAVSGovernanceLogicCaller, error) {
	contract, err := bindContractAVSGovernanceLogic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractAVSGovernanceLogicCaller{contract: contract}, nil
}

// NewContractAVSGovernanceLogicTransactor creates a new write-only instance of ContractAVSGovernanceLogic, bound to a specific deployed contract.
func NewContractAVSGovernanceLogicTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractAVSGovernanceLogicTransactor, error) {
	contract, err := bindContractAVSGovernanceLogic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractAVSGovernanceLogicTransactor{contract: contract}, nil
}

// NewContractAVSGovernanceLogicFilterer creates a new log filterer instance of ContractAVSGovernanceLogic, bound to a specific deployed contract.
func NewContractAVSGovernanceLogicFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractAVSGovernanceLogicFilterer, error) {
	contract, err := bindContractAVSGovernanceLogic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractAVSGovernanceLogicFilterer{contract: contract}, nil
}

// bindContractAVSGovernanceLogic binds a generic wrapper to an already deployed contract.
func bindContractAVSGovernanceLogic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractAVSGovernanceLogicMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractAVSGovernanceLogic.Contract.ContractAVSGovernanceLogicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.Contract.ContractAVSGovernanceLogicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.Contract.ContractAVSGovernanceLogicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractAVSGovernanceLogic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.Contract.contract.Transact(opts, method, params...)
}

// AllowInitializePath is a free data retrieval call binding the contract method 0xff7bd03d.
//
// Solidity: function allowInitializePath((uint32,bytes32,uint64) origin) view returns(bool)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicCaller) AllowInitializePath(opts *bind.CallOpts, origin Origin) (bool, error) {
	var out []interface{}
	err := _ContractAVSGovernanceLogic.contract.Call(opts, &out, "allowInitializePath", origin)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowInitializePath is a free data retrieval call binding the contract method 0xff7bd03d.
//
// Solidity: function allowInitializePath((uint32,bytes32,uint64) origin) view returns(bool)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicSession) AllowInitializePath(origin Origin) (bool, error) {
	return _ContractAVSGovernanceLogic.Contract.AllowInitializePath(&_ContractAVSGovernanceLogic.CallOpts, origin)
}

// AllowInitializePath is a free data retrieval call binding the contract method 0xff7bd03d.
//
// Solidity: function allowInitializePath((uint32,bytes32,uint64) origin) view returns(bool)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicCallerSession) AllowInitializePath(origin Origin) (bool, error) {
	return _ContractAVSGovernanceLogic.Contract.AllowInitializePath(&_ContractAVSGovernanceLogic.CallOpts, origin)
}

// AvsGovernance is a free data retrieval call binding the contract method 0x8d98e579.
//
// Solidity: function avsGovernance() view returns(address)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicCaller) AvsGovernance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAVSGovernanceLogic.contract.Call(opts, &out, "avsGovernance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsGovernance is a free data retrieval call binding the contract method 0x8d98e579.
//
// Solidity: function avsGovernance() view returns(address)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicSession) AvsGovernance() (common.Address, error) {
	return _ContractAVSGovernanceLogic.Contract.AvsGovernance(&_ContractAVSGovernanceLogic.CallOpts)
}

// AvsGovernance is a free data retrieval call binding the contract method 0x8d98e579.
//
// Solidity: function avsGovernance() view returns(address)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicCallerSession) AvsGovernance() (common.Address, error) {
	return _ContractAVSGovernanceLogic.Contract.AvsGovernance(&_ContractAVSGovernanceLogic.CallOpts)
}

// CallValue is a free data retrieval call binding the contract method 0xd62457f6.
//
// Solidity: function callValue() view returns(uint128)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicCaller) CallValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractAVSGovernanceLogic.contract.Call(opts, &out, "callValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CallValue is a free data retrieval call binding the contract method 0xd62457f6.
//
// Solidity: function callValue() view returns(uint128)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicSession) CallValue() (*big.Int, error) {
	return _ContractAVSGovernanceLogic.Contract.CallValue(&_ContractAVSGovernanceLogic.CallOpts)
}

// CallValue is a free data retrieval call binding the contract method 0xd62457f6.
//
// Solidity: function callValue() view returns(uint128)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicCallerSession) CallValue() (*big.Int, error) {
	return _ContractAVSGovernanceLogic.Contract.CallValue(&_ContractAVSGovernanceLogic.CallOpts)
}

// DstEid is a free data retrieval call binding the contract method 0x734ff5f3.
//
// Solidity: function dstEid() view returns(uint32)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicCaller) DstEid(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractAVSGovernanceLogic.contract.Call(opts, &out, "dstEid")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// DstEid is a free data retrieval call binding the contract method 0x734ff5f3.
//
// Solidity: function dstEid() view returns(uint32)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicSession) DstEid() (uint32, error) {
	return _ContractAVSGovernanceLogic.Contract.DstEid(&_ContractAVSGovernanceLogic.CallOpts)
}

// DstEid is a free data retrieval call binding the contract method 0x734ff5f3.
//
// Solidity: function dstEid() view returns(uint32)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicCallerSession) DstEid() (uint32, error) {
	return _ContractAVSGovernanceLogic.Contract.DstEid(&_ContractAVSGovernanceLogic.CallOpts)
}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicCaller) Endpoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAVSGovernanceLogic.contract.Call(opts, &out, "endpoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicSession) Endpoint() (common.Address, error) {
	return _ContractAVSGovernanceLogic.Contract.Endpoint(&_ContractAVSGovernanceLogic.CallOpts)
}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicCallerSession) Endpoint() (common.Address, error) {
	return _ContractAVSGovernanceLogic.Contract.Endpoint(&_ContractAVSGovernanceLogic.CallOpts)
}

// GasLimit is a free data retrieval call binding the contract method 0xf68016b7.
//
// Solidity: function gasLimit() view returns(uint128)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicCaller) GasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractAVSGovernanceLogic.contract.Call(opts, &out, "gasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GasLimit is a free data retrieval call binding the contract method 0xf68016b7.
//
// Solidity: function gasLimit() view returns(uint128)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicSession) GasLimit() (*big.Int, error) {
	return _ContractAVSGovernanceLogic.Contract.GasLimit(&_ContractAVSGovernanceLogic.CallOpts)
}

// GasLimit is a free data retrieval call binding the contract method 0xf68016b7.
//
// Solidity: function gasLimit() view returns(uint128)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicCallerSession) GasLimit() (*big.Int, error) {
	return _ContractAVSGovernanceLogic.Contract.GasLimit(&_ContractAVSGovernanceLogic.CallOpts)
}

// IsComposeMsgSender is a free data retrieval call binding the contract method 0x82413eac.
//
// Solidity: function isComposeMsgSender((uint32,bytes32,uint64) , bytes , address _sender) view returns(bool)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicCaller) IsComposeMsgSender(opts *bind.CallOpts, arg0 Origin, arg1 []byte, _sender common.Address) (bool, error) {
	var out []interface{}
	err := _ContractAVSGovernanceLogic.contract.Call(opts, &out, "isComposeMsgSender", arg0, arg1, _sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsComposeMsgSender is a free data retrieval call binding the contract method 0x82413eac.
//
// Solidity: function isComposeMsgSender((uint32,bytes32,uint64) , bytes , address _sender) view returns(bool)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicSession) IsComposeMsgSender(arg0 Origin, arg1 []byte, _sender common.Address) (bool, error) {
	return _ContractAVSGovernanceLogic.Contract.IsComposeMsgSender(&_ContractAVSGovernanceLogic.CallOpts, arg0, arg1, _sender)
}

// IsComposeMsgSender is a free data retrieval call binding the contract method 0x82413eac.
//
// Solidity: function isComposeMsgSender((uint32,bytes32,uint64) , bytes , address _sender) view returns(bool)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicCallerSession) IsComposeMsgSender(arg0 Origin, arg1 []byte, _sender common.Address) (bool, error) {
	return _ContractAVSGovernanceLogic.Contract.IsComposeMsgSender(&_ContractAVSGovernanceLogic.CallOpts, arg0, arg1, _sender)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address ) view returns(bool)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicCaller) IsWhitelisted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ContractAVSGovernanceLogic.contract.Call(opts, &out, "isWhitelisted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address ) view returns(bool)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicSession) IsWhitelisted(arg0 common.Address) (bool, error) {
	return _ContractAVSGovernanceLogic.Contract.IsWhitelisted(&_ContractAVSGovernanceLogic.CallOpts, arg0)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address ) view returns(bool)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicCallerSession) IsWhitelisted(arg0 common.Address) (bool, error) {
	return _ContractAVSGovernanceLogic.Contract.IsWhitelisted(&_ContractAVSGovernanceLogic.CallOpts, arg0)
}

// NextNonce is a free data retrieval call binding the contract method 0x7d25a05e.
//
// Solidity: function nextNonce(uint32 , bytes32 ) view returns(uint64 nonce)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicCaller) NextNonce(opts *bind.CallOpts, arg0 uint32, arg1 [32]byte) (uint64, error) {
	var out []interface{}
	err := _ContractAVSGovernanceLogic.contract.Call(opts, &out, "nextNonce", arg0, arg1)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NextNonce is a free data retrieval call binding the contract method 0x7d25a05e.
//
// Solidity: function nextNonce(uint32 , bytes32 ) view returns(uint64 nonce)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicSession) NextNonce(arg0 uint32, arg1 [32]byte) (uint64, error) {
	return _ContractAVSGovernanceLogic.Contract.NextNonce(&_ContractAVSGovernanceLogic.CallOpts, arg0, arg1)
}

// NextNonce is a free data retrieval call binding the contract method 0x7d25a05e.
//
// Solidity: function nextNonce(uint32 , bytes32 ) view returns(uint64 nonce)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicCallerSession) NextNonce(arg0 uint32, arg1 [32]byte) (uint64, error) {
	return _ContractAVSGovernanceLogic.Contract.NextNonce(&_ContractAVSGovernanceLogic.CallOpts, arg0, arg1)
}

// OAppVersion is a free data retrieval call binding the contract method 0x17442b70.
//
// Solidity: function oAppVersion() pure returns(uint64 senderVersion, uint64 receiverVersion)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicCaller) OAppVersion(opts *bind.CallOpts) (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	var out []interface{}
	err := _ContractAVSGovernanceLogic.contract.Call(opts, &out, "oAppVersion")

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
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicSession) OAppVersion() (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	return _ContractAVSGovernanceLogic.Contract.OAppVersion(&_ContractAVSGovernanceLogic.CallOpts)
}

// OAppVersion is a free data retrieval call binding the contract method 0x17442b70.
//
// Solidity: function oAppVersion() pure returns(uint64 senderVersion, uint64 receiverVersion)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicCallerSession) OAppVersion() (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	return _ContractAVSGovernanceLogic.Contract.OAppVersion(&_ContractAVSGovernanceLogic.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAVSGovernanceLogic.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicSession) Owner() (common.Address, error) {
	return _ContractAVSGovernanceLogic.Contract.Owner(&_ContractAVSGovernanceLogic.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicCallerSession) Owner() (common.Address, error) {
	return _ContractAVSGovernanceLogic.Contract.Owner(&_ContractAVSGovernanceLogic.CallOpts)
}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 eid) view returns(bytes32 peer)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicCaller) Peers(opts *bind.CallOpts, eid uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractAVSGovernanceLogic.contract.Call(opts, &out, "peers", eid)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 eid) view returns(bytes32 peer)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicSession) Peers(eid uint32) ([32]byte, error) {
	return _ContractAVSGovernanceLogic.Contract.Peers(&_ContractAVSGovernanceLogic.CallOpts, eid)
}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 eid) view returns(bytes32 peer)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicCallerSession) Peers(eid uint32) ([32]byte, error) {
	return _ContractAVSGovernanceLogic.Contract.Peers(&_ContractAVSGovernanceLogic.CallOpts, eid)
}

// TaskExecutionHub is a free data retrieval call binding the contract method 0xf0b351d9.
//
// Solidity: function taskExecutionHub() view returns(address)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicCaller) TaskExecutionHub(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAVSGovernanceLogic.contract.Call(opts, &out, "taskExecutionHub")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TaskExecutionHub is a free data retrieval call binding the contract method 0xf0b351d9.
//
// Solidity: function taskExecutionHub() view returns(address)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicSession) TaskExecutionHub() (common.Address, error) {
	return _ContractAVSGovernanceLogic.Contract.TaskExecutionHub(&_ContractAVSGovernanceLogic.CallOpts)
}

// TaskExecutionHub is a free data retrieval call binding the contract method 0xf0b351d9.
//
// Solidity: function taskExecutionHub() view returns(address)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicCallerSession) TaskExecutionHub() (common.Address, error) {
	return _ContractAVSGovernanceLogic.Contract.TaskExecutionHub(&_ContractAVSGovernanceLogic.CallOpts)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0x7f649783.
//
// Solidity: function addToWhitelist(address[] _operators) returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicTransactor) AddToWhitelist(opts *bind.TransactOpts, _operators []common.Address) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.contract.Transact(opts, "addToWhitelist", _operators)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0x7f649783.
//
// Solidity: function addToWhitelist(address[] _operators) returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicSession) AddToWhitelist(_operators []common.Address) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.Contract.AddToWhitelist(&_ContractAVSGovernanceLogic.TransactOpts, _operators)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0x7f649783.
//
// Solidity: function addToWhitelist(address[] _operators) returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicTransactorSession) AddToWhitelist(_operators []common.Address) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.Contract.AddToWhitelist(&_ContractAVSGovernanceLogic.TransactOpts, _operators)
}

// AfterOperatorRegistered is a paid mutator transaction binding the contract method 0xed8ad836.
//
// Solidity: function afterOperatorRegistered(address _operator, uint256 , uint256[4] , address ) returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicTransactor) AfterOperatorRegistered(opts *bind.TransactOpts, _operator common.Address, arg1 *big.Int, arg2 [4]*big.Int, arg3 common.Address) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.contract.Transact(opts, "afterOperatorRegistered", _operator, arg1, arg2, arg3)
}

// AfterOperatorRegistered is a paid mutator transaction binding the contract method 0xed8ad836.
//
// Solidity: function afterOperatorRegistered(address _operator, uint256 , uint256[4] , address ) returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicSession) AfterOperatorRegistered(_operator common.Address, arg1 *big.Int, arg2 [4]*big.Int, arg3 common.Address) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.Contract.AfterOperatorRegistered(&_ContractAVSGovernanceLogic.TransactOpts, _operator, arg1, arg2, arg3)
}

// AfterOperatorRegistered is a paid mutator transaction binding the contract method 0xed8ad836.
//
// Solidity: function afterOperatorRegistered(address _operator, uint256 , uint256[4] , address ) returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicTransactorSession) AfterOperatorRegistered(_operator common.Address, arg1 *big.Int, arg2 [4]*big.Int, arg3 common.Address) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.Contract.AfterOperatorRegistered(&_ContractAVSGovernanceLogic.TransactOpts, _operator, arg1, arg2, arg3)
}

// AfterOperatorUnregistered is a paid mutator transaction binding the contract method 0xe9ecc1cb.
//
// Solidity: function afterOperatorUnregistered(address _operator) returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicTransactor) AfterOperatorUnregistered(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.contract.Transact(opts, "afterOperatorUnregistered", _operator)
}

// AfterOperatorUnregistered is a paid mutator transaction binding the contract method 0xe9ecc1cb.
//
// Solidity: function afterOperatorUnregistered(address _operator) returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicSession) AfterOperatorUnregistered(_operator common.Address) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.Contract.AfterOperatorUnregistered(&_ContractAVSGovernanceLogic.TransactOpts, _operator)
}

// AfterOperatorUnregistered is a paid mutator transaction binding the contract method 0xe9ecc1cb.
//
// Solidity: function afterOperatorUnregistered(address _operator) returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicTransactorSession) AfterOperatorUnregistered(_operator common.Address) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.Contract.AfterOperatorUnregistered(&_ContractAVSGovernanceLogic.TransactOpts, _operator)
}

// BeforeOperatorRegistered is a paid mutator transaction binding the contract method 0x94e7a3f0.
//
// Solidity: function beforeOperatorRegistered(address _operator, uint256 , uint256[4] , address ) returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicTransactor) BeforeOperatorRegistered(opts *bind.TransactOpts, _operator common.Address, arg1 *big.Int, arg2 [4]*big.Int, arg3 common.Address) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.contract.Transact(opts, "beforeOperatorRegistered", _operator, arg1, arg2, arg3)
}

// BeforeOperatorRegistered is a paid mutator transaction binding the contract method 0x94e7a3f0.
//
// Solidity: function beforeOperatorRegistered(address _operator, uint256 , uint256[4] , address ) returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicSession) BeforeOperatorRegistered(_operator common.Address, arg1 *big.Int, arg2 [4]*big.Int, arg3 common.Address) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.Contract.BeforeOperatorRegistered(&_ContractAVSGovernanceLogic.TransactOpts, _operator, arg1, arg2, arg3)
}

// BeforeOperatorRegistered is a paid mutator transaction binding the contract method 0x94e7a3f0.
//
// Solidity: function beforeOperatorRegistered(address _operator, uint256 , uint256[4] , address ) returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicTransactorSession) BeforeOperatorRegistered(_operator common.Address, arg1 *big.Int, arg2 [4]*big.Int, arg3 common.Address) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.Contract.BeforeOperatorRegistered(&_ContractAVSGovernanceLogic.TransactOpts, _operator, arg1, arg2, arg3)
}

// BeforeOperatorUnregistered is a paid mutator transaction binding the contract method 0x471f9e64.
//
// Solidity: function beforeOperatorUnregistered(address ) returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicTransactor) BeforeOperatorUnregistered(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.contract.Transact(opts, "beforeOperatorUnregistered", arg0)
}

// BeforeOperatorUnregistered is a paid mutator transaction binding the contract method 0x471f9e64.
//
// Solidity: function beforeOperatorUnregistered(address ) returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicSession) BeforeOperatorUnregistered(arg0 common.Address) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.Contract.BeforeOperatorUnregistered(&_ContractAVSGovernanceLogic.TransactOpts, arg0)
}

// BeforeOperatorUnregistered is a paid mutator transaction binding the contract method 0x471f9e64.
//
// Solidity: function beforeOperatorUnregistered(address ) returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicTransactorSession) BeforeOperatorUnregistered(arg0 common.Address) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.Contract.BeforeOperatorUnregistered(&_ContractAVSGovernanceLogic.TransactOpts, arg0)
}

// LzReceive is a paid mutator transaction binding the contract method 0x13137d65.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicTransactor) LzReceive(opts *bind.TransactOpts, _origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.contract.Transact(opts, "lzReceive", _origin, _guid, _message, _executor, _extraData)
}

// LzReceive is a paid mutator transaction binding the contract method 0x13137d65.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicSession) LzReceive(_origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.Contract.LzReceive(&_ContractAVSGovernanceLogic.TransactOpts, _origin, _guid, _message, _executor, _extraData)
}

// LzReceive is a paid mutator transaction binding the contract method 0x13137d65.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicTransactorSession) LzReceive(_origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.Contract.LzReceive(&_ContractAVSGovernanceLogic.TransactOpts, _origin, _guid, _message, _executor, _extraData)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x548db174.
//
// Solidity: function removeFromWhitelist(address[] _operators) returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicTransactor) RemoveFromWhitelist(opts *bind.TransactOpts, _operators []common.Address) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.contract.Transact(opts, "removeFromWhitelist", _operators)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x548db174.
//
// Solidity: function removeFromWhitelist(address[] _operators) returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicSession) RemoveFromWhitelist(_operators []common.Address) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.Contract.RemoveFromWhitelist(&_ContractAVSGovernanceLogic.TransactOpts, _operators)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x548db174.
//
// Solidity: function removeFromWhitelist(address[] _operators) returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicTransactorSession) RemoveFromWhitelist(_operators []common.Address) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.Contract.RemoveFromWhitelist(&_ContractAVSGovernanceLogic.TransactOpts, _operators)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.Contract.RenounceOwnership(&_ContractAVSGovernanceLogic.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.Contract.RenounceOwnership(&_ContractAVSGovernanceLogic.TransactOpts)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicTransactor) SetDelegate(opts *bind.TransactOpts, _delegate common.Address) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.contract.Transact(opts, "setDelegate", _delegate)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicSession) SetDelegate(_delegate common.Address) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.Contract.SetDelegate(&_ContractAVSGovernanceLogic.TransactOpts, _delegate)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicTransactorSession) SetDelegate(_delegate common.Address) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.Contract.SetDelegate(&_ContractAVSGovernanceLogic.TransactOpts, _delegate)
}

// SetGasOptions is a paid mutator transaction binding the contract method 0x9dbe1538.
//
// Solidity: function setGasOptions(uint128 _gasLimit, uint128 _callValue) returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicTransactor) SetGasOptions(opts *bind.TransactOpts, _gasLimit *big.Int, _callValue *big.Int) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.contract.Transact(opts, "setGasOptions", _gasLimit, _callValue)
}

// SetGasOptions is a paid mutator transaction binding the contract method 0x9dbe1538.
//
// Solidity: function setGasOptions(uint128 _gasLimit, uint128 _callValue) returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicSession) SetGasOptions(_gasLimit *big.Int, _callValue *big.Int) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.Contract.SetGasOptions(&_ContractAVSGovernanceLogic.TransactOpts, _gasLimit, _callValue)
}

// SetGasOptions is a paid mutator transaction binding the contract method 0x9dbe1538.
//
// Solidity: function setGasOptions(uint128 _gasLimit, uint128 _callValue) returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicTransactorSession) SetGasOptions(_gasLimit *big.Int, _callValue *big.Int) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.Contract.SetGasOptions(&_ContractAVSGovernanceLogic.TransactOpts, _gasLimit, _callValue)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicTransactor) SetPeer(opts *bind.TransactOpts, _eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.contract.Transact(opts, "setPeer", _eid, _peer)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicSession) SetPeer(_eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.Contract.SetPeer(&_ContractAVSGovernanceLogic.TransactOpts, _eid, _peer)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicTransactorSession) SetPeer(_eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.Contract.SetPeer(&_ContractAVSGovernanceLogic.TransactOpts, _eid, _peer)
}

// SetTaskExecutionHub is a paid mutator transaction binding the contract method 0x96e51b9c.
//
// Solidity: function setTaskExecutionHub(address _taskExecutionHub) returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicTransactor) SetTaskExecutionHub(opts *bind.TransactOpts, _taskExecutionHub common.Address) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.contract.Transact(opts, "setTaskExecutionHub", _taskExecutionHub)
}

// SetTaskExecutionHub is a paid mutator transaction binding the contract method 0x96e51b9c.
//
// Solidity: function setTaskExecutionHub(address _taskExecutionHub) returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicSession) SetTaskExecutionHub(_taskExecutionHub common.Address) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.Contract.SetTaskExecutionHub(&_ContractAVSGovernanceLogic.TransactOpts, _taskExecutionHub)
}

// SetTaskExecutionHub is a paid mutator transaction binding the contract method 0x96e51b9c.
//
// Solidity: function setTaskExecutionHub(address _taskExecutionHub) returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicTransactorSession) SetTaskExecutionHub(_taskExecutionHub common.Address) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.Contract.SetTaskExecutionHub(&_ContractAVSGovernanceLogic.TransactOpts, _taskExecutionHub)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.Contract.TransferOwnership(&_ContractAVSGovernanceLogic.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.Contract.TransferOwnership(&_ContractAVSGovernanceLogic.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _to, uint256 _amount) returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicTransactor) Withdraw(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.contract.Transact(opts, "withdraw", _to, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _to, uint256 _amount) returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicSession) Withdraw(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.Contract.Withdraw(&_ContractAVSGovernanceLogic.TransactOpts, _to, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _to, uint256 _amount) returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicTransactorSession) Withdraw(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.Contract.Withdraw(&_ContractAVSGovernanceLogic.TransactOpts, _to, _amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicSession) Receive() (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.Contract.Receive(&_ContractAVSGovernanceLogic.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicTransactorSession) Receive() (*types.Transaction, error) {
	return _ContractAVSGovernanceLogic.Contract.Receive(&_ContractAVSGovernanceLogic.TransactOpts)
}

// ContractAVSGovernanceLogicGasOptionsUpdatedIterator is returned from FilterGasOptionsUpdated and is used to iterate over the raw logs and unpacked data for GasOptionsUpdated events raised by the ContractAVSGovernanceLogic contract.
type ContractAVSGovernanceLogicGasOptionsUpdatedIterator struct {
	Event *ContractAVSGovernanceLogicGasOptionsUpdated // Event containing the contract specifics and raw log

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
func (it *ContractAVSGovernanceLogicGasOptionsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAVSGovernanceLogicGasOptionsUpdated)
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
		it.Event = new(ContractAVSGovernanceLogicGasOptionsUpdated)
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
func (it *ContractAVSGovernanceLogicGasOptionsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAVSGovernanceLogicGasOptionsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAVSGovernanceLogicGasOptionsUpdated represents a GasOptionsUpdated event raised by the ContractAVSGovernanceLogic contract.
type ContractAVSGovernanceLogicGasOptionsUpdated struct {
	GasLimit  *big.Int
	CallValue *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterGasOptionsUpdated is a free log retrieval operation binding the contract event 0x87768cdb9e7a69da85a21ad47c9654958b5cbc0447a55b66d0e3e0d640430c2d.
//
// Solidity: event GasOptionsUpdated(uint128 gasLimit, uint128 callValue)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicFilterer) FilterGasOptionsUpdated(opts *bind.FilterOpts) (*ContractAVSGovernanceLogicGasOptionsUpdatedIterator, error) {

	logs, sub, err := _ContractAVSGovernanceLogic.contract.FilterLogs(opts, "GasOptionsUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractAVSGovernanceLogicGasOptionsUpdatedIterator{contract: _ContractAVSGovernanceLogic.contract, event: "GasOptionsUpdated", logs: logs, sub: sub}, nil
}

// WatchGasOptionsUpdated is a free log subscription operation binding the contract event 0x87768cdb9e7a69da85a21ad47c9654958b5cbc0447a55b66d0e3e0d640430c2d.
//
// Solidity: event GasOptionsUpdated(uint128 gasLimit, uint128 callValue)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicFilterer) WatchGasOptionsUpdated(opts *bind.WatchOpts, sink chan<- *ContractAVSGovernanceLogicGasOptionsUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractAVSGovernanceLogic.contract.WatchLogs(opts, "GasOptionsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAVSGovernanceLogicGasOptionsUpdated)
				if err := _ContractAVSGovernanceLogic.contract.UnpackLog(event, "GasOptionsUpdated", log); err != nil {
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

// ParseGasOptionsUpdated is a log parse operation binding the contract event 0x87768cdb9e7a69da85a21ad47c9654958b5cbc0447a55b66d0e3e0d640430c2d.
//
// Solidity: event GasOptionsUpdated(uint128 gasLimit, uint128 callValue)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicFilterer) ParseGasOptionsUpdated(log types.Log) (*ContractAVSGovernanceLogicGasOptionsUpdated, error) {
	event := new(ContractAVSGovernanceLogicGasOptionsUpdated)
	if err := _ContractAVSGovernanceLogic.contract.UnpackLog(event, "GasOptionsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAVSGovernanceLogicLowBalanceAlertIterator is returned from FilterLowBalanceAlert and is used to iterate over the raw logs and unpacked data for LowBalanceAlert events raised by the ContractAVSGovernanceLogic contract.
type ContractAVSGovernanceLogicLowBalanceAlertIterator struct {
	Event *ContractAVSGovernanceLogicLowBalanceAlert // Event containing the contract specifics and raw log

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
func (it *ContractAVSGovernanceLogicLowBalanceAlertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAVSGovernanceLogicLowBalanceAlert)
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
		it.Event = new(ContractAVSGovernanceLogicLowBalanceAlert)
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
func (it *ContractAVSGovernanceLogicLowBalanceAlertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAVSGovernanceLogicLowBalanceAlertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAVSGovernanceLogicLowBalanceAlert represents a LowBalanceAlert event raised by the ContractAVSGovernanceLogic contract.
type ContractAVSGovernanceLogicLowBalanceAlert struct {
	CurrentBalance *big.Int
	Threshold      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLowBalanceAlert is a free log retrieval operation binding the contract event 0xb1b538a145d528cb9b8a428a3f47d5730430a03c2d98d7562357ee7db16dbd26.
//
// Solidity: event LowBalanceAlert(uint256 currentBalance, uint256 threshold)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicFilterer) FilterLowBalanceAlert(opts *bind.FilterOpts) (*ContractAVSGovernanceLogicLowBalanceAlertIterator, error) {

	logs, sub, err := _ContractAVSGovernanceLogic.contract.FilterLogs(opts, "LowBalanceAlert")
	if err != nil {
		return nil, err
	}
	return &ContractAVSGovernanceLogicLowBalanceAlertIterator{contract: _ContractAVSGovernanceLogic.contract, event: "LowBalanceAlert", logs: logs, sub: sub}, nil
}

// WatchLowBalanceAlert is a free log subscription operation binding the contract event 0xb1b538a145d528cb9b8a428a3f47d5730430a03c2d98d7562357ee7db16dbd26.
//
// Solidity: event LowBalanceAlert(uint256 currentBalance, uint256 threshold)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicFilterer) WatchLowBalanceAlert(opts *bind.WatchOpts, sink chan<- *ContractAVSGovernanceLogicLowBalanceAlert) (event.Subscription, error) {

	logs, sub, err := _ContractAVSGovernanceLogic.contract.WatchLogs(opts, "LowBalanceAlert")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAVSGovernanceLogicLowBalanceAlert)
				if err := _ContractAVSGovernanceLogic.contract.UnpackLog(event, "LowBalanceAlert", log); err != nil {
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
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicFilterer) ParseLowBalanceAlert(log types.Log) (*ContractAVSGovernanceLogicLowBalanceAlert, error) {
	event := new(ContractAVSGovernanceLogicLowBalanceAlert)
	if err := _ContractAVSGovernanceLogic.contract.UnpackLog(event, "LowBalanceAlert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAVSGovernanceLogicMessageFailedIterator is returned from FilterMessageFailed and is used to iterate over the raw logs and unpacked data for MessageFailed events raised by the ContractAVSGovernanceLogic contract.
type ContractAVSGovernanceLogicMessageFailedIterator struct {
	Event *ContractAVSGovernanceLogicMessageFailed // Event containing the contract specifics and raw log

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
func (it *ContractAVSGovernanceLogicMessageFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAVSGovernanceLogicMessageFailed)
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
		it.Event = new(ContractAVSGovernanceLogicMessageFailed)
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
func (it *ContractAVSGovernanceLogicMessageFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAVSGovernanceLogicMessageFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAVSGovernanceLogicMessageFailed represents a MessageFailed event raised by the ContractAVSGovernanceLogic contract.
type ContractAVSGovernanceLogicMessageFailed struct {
	DstEid uint32
	Guid   [32]byte
	Reason []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMessageFailed is a free log retrieval operation binding the contract event 0x2254f63be429412a67ef70cf2ef65d61a88074bfd112333e0a449a9e4a7683c8.
//
// Solidity: event MessageFailed(uint32 indexed dstEid, bytes32 indexed guid, bytes reason)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicFilterer) FilterMessageFailed(opts *bind.FilterOpts, dstEid []uint32, guid [][32]byte) (*ContractAVSGovernanceLogicMessageFailedIterator, error) {

	var dstEidRule []interface{}
	for _, dstEidItem := range dstEid {
		dstEidRule = append(dstEidRule, dstEidItem)
	}
	var guidRule []interface{}
	for _, guidItem := range guid {
		guidRule = append(guidRule, guidItem)
	}

	logs, sub, err := _ContractAVSGovernanceLogic.contract.FilterLogs(opts, "MessageFailed", dstEidRule, guidRule)
	if err != nil {
		return nil, err
	}
	return &ContractAVSGovernanceLogicMessageFailedIterator{contract: _ContractAVSGovernanceLogic.contract, event: "MessageFailed", logs: logs, sub: sub}, nil
}

// WatchMessageFailed is a free log subscription operation binding the contract event 0x2254f63be429412a67ef70cf2ef65d61a88074bfd112333e0a449a9e4a7683c8.
//
// Solidity: event MessageFailed(uint32 indexed dstEid, bytes32 indexed guid, bytes reason)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicFilterer) WatchMessageFailed(opts *bind.WatchOpts, sink chan<- *ContractAVSGovernanceLogicMessageFailed, dstEid []uint32, guid [][32]byte) (event.Subscription, error) {

	var dstEidRule []interface{}
	for _, dstEidItem := range dstEid {
		dstEidRule = append(dstEidRule, dstEidItem)
	}
	var guidRule []interface{}
	for _, guidItem := range guid {
		guidRule = append(guidRule, guidItem)
	}

	logs, sub, err := _ContractAVSGovernanceLogic.contract.WatchLogs(opts, "MessageFailed", dstEidRule, guidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAVSGovernanceLogicMessageFailed)
				if err := _ContractAVSGovernanceLogic.contract.UnpackLog(event, "MessageFailed", log); err != nil {
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
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicFilterer) ParseMessageFailed(log types.Log) (*ContractAVSGovernanceLogicMessageFailed, error) {
	event := new(ContractAVSGovernanceLogicMessageFailed)
	if err := _ContractAVSGovernanceLogic.contract.UnpackLog(event, "MessageFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAVSGovernanceLogicMessageSentIterator is returned from FilterMessageSent and is used to iterate over the raw logs and unpacked data for MessageSent events raised by the ContractAVSGovernanceLogic contract.
type ContractAVSGovernanceLogicMessageSentIterator struct {
	Event *ContractAVSGovernanceLogicMessageSent // Event containing the contract specifics and raw log

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
func (it *ContractAVSGovernanceLogicMessageSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAVSGovernanceLogicMessageSent)
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
		it.Event = new(ContractAVSGovernanceLogicMessageSent)
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
func (it *ContractAVSGovernanceLogicMessageSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAVSGovernanceLogicMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAVSGovernanceLogicMessageSent represents a MessageSent event raised by the ContractAVSGovernanceLogic contract.
type ContractAVSGovernanceLogicMessageSent struct {
	DstEid uint32
	Guid   [32]byte
	Fee    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMessageSent is a free log retrieval operation binding the contract event 0x0597c8258e620343bfa7f3cf82a7d6578378f08248fd958f32bcab61a1004ec3.
//
// Solidity: event MessageSent(uint32 indexed dstEid, bytes32 indexed guid, uint256 fee)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicFilterer) FilterMessageSent(opts *bind.FilterOpts, dstEid []uint32, guid [][32]byte) (*ContractAVSGovernanceLogicMessageSentIterator, error) {

	var dstEidRule []interface{}
	for _, dstEidItem := range dstEid {
		dstEidRule = append(dstEidRule, dstEidItem)
	}
	var guidRule []interface{}
	for _, guidItem := range guid {
		guidRule = append(guidRule, guidItem)
	}

	logs, sub, err := _ContractAVSGovernanceLogic.contract.FilterLogs(opts, "MessageSent", dstEidRule, guidRule)
	if err != nil {
		return nil, err
	}
	return &ContractAVSGovernanceLogicMessageSentIterator{contract: _ContractAVSGovernanceLogic.contract, event: "MessageSent", logs: logs, sub: sub}, nil
}

// WatchMessageSent is a free log subscription operation binding the contract event 0x0597c8258e620343bfa7f3cf82a7d6578378f08248fd958f32bcab61a1004ec3.
//
// Solidity: event MessageSent(uint32 indexed dstEid, bytes32 indexed guid, uint256 fee)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicFilterer) WatchMessageSent(opts *bind.WatchOpts, sink chan<- *ContractAVSGovernanceLogicMessageSent, dstEid []uint32, guid [][32]byte) (event.Subscription, error) {

	var dstEidRule []interface{}
	for _, dstEidItem := range dstEid {
		dstEidRule = append(dstEidRule, dstEidItem)
	}
	var guidRule []interface{}
	for _, guidItem := range guid {
		guidRule = append(guidRule, guidItem)
	}

	logs, sub, err := _ContractAVSGovernanceLogic.contract.WatchLogs(opts, "MessageSent", dstEidRule, guidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAVSGovernanceLogicMessageSent)
				if err := _ContractAVSGovernanceLogic.contract.UnpackLog(event, "MessageSent", log); err != nil {
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

// ParseMessageSent is a log parse operation binding the contract event 0x0597c8258e620343bfa7f3cf82a7d6578378f08248fd958f32bcab61a1004ec3.
//
// Solidity: event MessageSent(uint32 indexed dstEid, bytes32 indexed guid, uint256 fee)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicFilterer) ParseMessageSent(log types.Log) (*ContractAVSGovernanceLogicMessageSent, error) {
	event := new(ContractAVSGovernanceLogicMessageSent)
	if err := _ContractAVSGovernanceLogic.contract.UnpackLog(event, "MessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAVSGovernanceLogicOperatorRegisteredIterator is returned from FilterOperatorRegistered and is used to iterate over the raw logs and unpacked data for OperatorRegistered events raised by the ContractAVSGovernanceLogic contract.
type ContractAVSGovernanceLogicOperatorRegisteredIterator struct {
	Event *ContractAVSGovernanceLogicOperatorRegistered // Event containing the contract specifics and raw log

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
func (it *ContractAVSGovernanceLogicOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAVSGovernanceLogicOperatorRegistered)
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
		it.Event = new(ContractAVSGovernanceLogicOperatorRegistered)
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
func (it *ContractAVSGovernanceLogicOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAVSGovernanceLogicOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAVSGovernanceLogicOperatorRegistered represents a OperatorRegistered event raised by the ContractAVSGovernanceLogic contract.
type ContractAVSGovernanceLogicOperatorRegistered struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0x4d0eb1f4bac8744fd2be119845e23b3befc88094b42bcda1204c65694a00f9e5.
//
// Solidity: event OperatorRegistered(address indexed operator)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicFilterer) FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address) (*ContractAVSGovernanceLogicOperatorRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAVSGovernanceLogic.contract.FilterLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractAVSGovernanceLogicOperatorRegisteredIterator{contract: _ContractAVSGovernanceLogic.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0x4d0eb1f4bac8744fd2be119845e23b3befc88094b42bcda1204c65694a00f9e5.
//
// Solidity: event OperatorRegistered(address indexed operator)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicFilterer) WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *ContractAVSGovernanceLogicOperatorRegistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAVSGovernanceLogic.contract.WatchLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAVSGovernanceLogicOperatorRegistered)
				if err := _ContractAVSGovernanceLogic.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
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

// ParseOperatorRegistered is a log parse operation binding the contract event 0x4d0eb1f4bac8744fd2be119845e23b3befc88094b42bcda1204c65694a00f9e5.
//
// Solidity: event OperatorRegistered(address indexed operator)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicFilterer) ParseOperatorRegistered(log types.Log) (*ContractAVSGovernanceLogicOperatorRegistered, error) {
	event := new(ContractAVSGovernanceLogicOperatorRegistered)
	if err := _ContractAVSGovernanceLogic.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAVSGovernanceLogicOperatorUnregisteredIterator is returned from FilterOperatorUnregistered and is used to iterate over the raw logs and unpacked data for OperatorUnregistered events raised by the ContractAVSGovernanceLogic contract.
type ContractAVSGovernanceLogicOperatorUnregisteredIterator struct {
	Event *ContractAVSGovernanceLogicOperatorUnregistered // Event containing the contract specifics and raw log

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
func (it *ContractAVSGovernanceLogicOperatorUnregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAVSGovernanceLogicOperatorUnregistered)
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
		it.Event = new(ContractAVSGovernanceLogicOperatorUnregistered)
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
func (it *ContractAVSGovernanceLogicOperatorUnregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAVSGovernanceLogicOperatorUnregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAVSGovernanceLogicOperatorUnregistered represents a OperatorUnregistered event raised by the ContractAVSGovernanceLogic contract.
type ContractAVSGovernanceLogicOperatorUnregistered struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorUnregistered is a free log retrieval operation binding the contract event 0x6f42117a557500c705ddf040a619d86f39101e6b74ac20d7b3e5943ba473fc7f.
//
// Solidity: event OperatorUnregistered(address indexed operator)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicFilterer) FilterOperatorUnregistered(opts *bind.FilterOpts, operator []common.Address) (*ContractAVSGovernanceLogicOperatorUnregisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAVSGovernanceLogic.contract.FilterLogs(opts, "OperatorUnregistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractAVSGovernanceLogicOperatorUnregisteredIterator{contract: _ContractAVSGovernanceLogic.contract, event: "OperatorUnregistered", logs: logs, sub: sub}, nil
}

// WatchOperatorUnregistered is a free log subscription operation binding the contract event 0x6f42117a557500c705ddf040a619d86f39101e6b74ac20d7b3e5943ba473fc7f.
//
// Solidity: event OperatorUnregistered(address indexed operator)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicFilterer) WatchOperatorUnregistered(opts *bind.WatchOpts, sink chan<- *ContractAVSGovernanceLogicOperatorUnregistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAVSGovernanceLogic.contract.WatchLogs(opts, "OperatorUnregistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAVSGovernanceLogicOperatorUnregistered)
				if err := _ContractAVSGovernanceLogic.contract.UnpackLog(event, "OperatorUnregistered", log); err != nil {
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

// ParseOperatorUnregistered is a log parse operation binding the contract event 0x6f42117a557500c705ddf040a619d86f39101e6b74ac20d7b3e5943ba473fc7f.
//
// Solidity: event OperatorUnregistered(address indexed operator)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicFilterer) ParseOperatorUnregistered(log types.Log) (*ContractAVSGovernanceLogicOperatorUnregistered, error) {
	event := new(ContractAVSGovernanceLogicOperatorUnregistered)
	if err := _ContractAVSGovernanceLogic.contract.UnpackLog(event, "OperatorUnregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAVSGovernanceLogicOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractAVSGovernanceLogic contract.
type ContractAVSGovernanceLogicOwnershipTransferredIterator struct {
	Event *ContractAVSGovernanceLogicOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractAVSGovernanceLogicOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAVSGovernanceLogicOwnershipTransferred)
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
		it.Event = new(ContractAVSGovernanceLogicOwnershipTransferred)
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
func (it *ContractAVSGovernanceLogicOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAVSGovernanceLogicOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAVSGovernanceLogicOwnershipTransferred represents a OwnershipTransferred event raised by the ContractAVSGovernanceLogic contract.
type ContractAVSGovernanceLogicOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractAVSGovernanceLogicOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractAVSGovernanceLogic.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractAVSGovernanceLogicOwnershipTransferredIterator{contract: _ContractAVSGovernanceLogic.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractAVSGovernanceLogicOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractAVSGovernanceLogic.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAVSGovernanceLogicOwnershipTransferred)
				if err := _ContractAVSGovernanceLogic.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicFilterer) ParseOwnershipTransferred(log types.Log) (*ContractAVSGovernanceLogicOwnershipTransferred, error) {
	event := new(ContractAVSGovernanceLogicOwnershipTransferred)
	if err := _ContractAVSGovernanceLogic.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAVSGovernanceLogicPeerSetIterator is returned from FilterPeerSet and is used to iterate over the raw logs and unpacked data for PeerSet events raised by the ContractAVSGovernanceLogic contract.
type ContractAVSGovernanceLogicPeerSetIterator struct {
	Event *ContractAVSGovernanceLogicPeerSet // Event containing the contract specifics and raw log

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
func (it *ContractAVSGovernanceLogicPeerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAVSGovernanceLogicPeerSet)
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
		it.Event = new(ContractAVSGovernanceLogicPeerSet)
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
func (it *ContractAVSGovernanceLogicPeerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAVSGovernanceLogicPeerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAVSGovernanceLogicPeerSet represents a PeerSet event raised by the ContractAVSGovernanceLogic contract.
type ContractAVSGovernanceLogicPeerSet struct {
	Eid  uint32
	Peer [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPeerSet is a free log retrieval operation binding the contract event 0x238399d427b947898edb290f5ff0f9109849b1c3ba196a42e35f00c50a54b98b.
//
// Solidity: event PeerSet(uint32 eid, bytes32 peer)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicFilterer) FilterPeerSet(opts *bind.FilterOpts) (*ContractAVSGovernanceLogicPeerSetIterator, error) {

	logs, sub, err := _ContractAVSGovernanceLogic.contract.FilterLogs(opts, "PeerSet")
	if err != nil {
		return nil, err
	}
	return &ContractAVSGovernanceLogicPeerSetIterator{contract: _ContractAVSGovernanceLogic.contract, event: "PeerSet", logs: logs, sub: sub}, nil
}

// WatchPeerSet is a free log subscription operation binding the contract event 0x238399d427b947898edb290f5ff0f9109849b1c3ba196a42e35f00c50a54b98b.
//
// Solidity: event PeerSet(uint32 eid, bytes32 peer)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicFilterer) WatchPeerSet(opts *bind.WatchOpts, sink chan<- *ContractAVSGovernanceLogicPeerSet) (event.Subscription, error) {

	logs, sub, err := _ContractAVSGovernanceLogic.contract.WatchLogs(opts, "PeerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAVSGovernanceLogicPeerSet)
				if err := _ContractAVSGovernanceLogic.contract.UnpackLog(event, "PeerSet", log); err != nil {
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
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicFilterer) ParsePeerSet(log types.Log) (*ContractAVSGovernanceLogicPeerSet, error) {
	event := new(ContractAVSGovernanceLogicPeerSet)
	if err := _ContractAVSGovernanceLogic.contract.UnpackLog(event, "PeerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAVSGovernanceLogicUnWhitelistedIterator is returned from FilterUnWhitelisted and is used to iterate over the raw logs and unpacked data for UnWhitelisted events raised by the ContractAVSGovernanceLogic contract.
type ContractAVSGovernanceLogicUnWhitelistedIterator struct {
	Event *ContractAVSGovernanceLogicUnWhitelisted // Event containing the contract specifics and raw log

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
func (it *ContractAVSGovernanceLogicUnWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAVSGovernanceLogicUnWhitelisted)
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
		it.Event = new(ContractAVSGovernanceLogicUnWhitelisted)
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
func (it *ContractAVSGovernanceLogicUnWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAVSGovernanceLogicUnWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAVSGovernanceLogicUnWhitelisted represents a UnWhitelisted event raised by the ContractAVSGovernanceLogic contract.
type ContractAVSGovernanceLogicUnWhitelisted struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUnWhitelisted is a free log retrieval operation binding the contract event 0x7cdb51b0cc2e541ad7e9471c358de415f5dbaff6cca78e3393d445346c610c1b.
//
// Solidity: event UnWhitelisted(address indexed operator)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicFilterer) FilterUnWhitelisted(opts *bind.FilterOpts, operator []common.Address) (*ContractAVSGovernanceLogicUnWhitelistedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAVSGovernanceLogic.contract.FilterLogs(opts, "UnWhitelisted", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractAVSGovernanceLogicUnWhitelistedIterator{contract: _ContractAVSGovernanceLogic.contract, event: "UnWhitelisted", logs: logs, sub: sub}, nil
}

// WatchUnWhitelisted is a free log subscription operation binding the contract event 0x7cdb51b0cc2e541ad7e9471c358de415f5dbaff6cca78e3393d445346c610c1b.
//
// Solidity: event UnWhitelisted(address indexed operator)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicFilterer) WatchUnWhitelisted(opts *bind.WatchOpts, sink chan<- *ContractAVSGovernanceLogicUnWhitelisted, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAVSGovernanceLogic.contract.WatchLogs(opts, "UnWhitelisted", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAVSGovernanceLogicUnWhitelisted)
				if err := _ContractAVSGovernanceLogic.contract.UnpackLog(event, "UnWhitelisted", log); err != nil {
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

// ParseUnWhitelisted is a log parse operation binding the contract event 0x7cdb51b0cc2e541ad7e9471c358de415f5dbaff6cca78e3393d445346c610c1b.
//
// Solidity: event UnWhitelisted(address indexed operator)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicFilterer) ParseUnWhitelisted(log types.Log) (*ContractAVSGovernanceLogicUnWhitelisted, error) {
	event := new(ContractAVSGovernanceLogicUnWhitelisted)
	if err := _ContractAVSGovernanceLogic.contract.UnpackLog(event, "UnWhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAVSGovernanceLogicWhitelistedIterator is returned from FilterWhitelisted and is used to iterate over the raw logs and unpacked data for Whitelisted events raised by the ContractAVSGovernanceLogic contract.
type ContractAVSGovernanceLogicWhitelistedIterator struct {
	Event *ContractAVSGovernanceLogicWhitelisted // Event containing the contract specifics and raw log

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
func (it *ContractAVSGovernanceLogicWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAVSGovernanceLogicWhitelisted)
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
		it.Event = new(ContractAVSGovernanceLogicWhitelisted)
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
func (it *ContractAVSGovernanceLogicWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAVSGovernanceLogicWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAVSGovernanceLogicWhitelisted represents a Whitelisted event raised by the ContractAVSGovernanceLogic contract.
type ContractAVSGovernanceLogicWhitelisted struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWhitelisted is a free log retrieval operation binding the contract event 0xaab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a54.
//
// Solidity: event Whitelisted(address indexed operator)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicFilterer) FilterWhitelisted(opts *bind.FilterOpts, operator []common.Address) (*ContractAVSGovernanceLogicWhitelistedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAVSGovernanceLogic.contract.FilterLogs(opts, "Whitelisted", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractAVSGovernanceLogicWhitelistedIterator{contract: _ContractAVSGovernanceLogic.contract, event: "Whitelisted", logs: logs, sub: sub}, nil
}

// WatchWhitelisted is a free log subscription operation binding the contract event 0xaab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a54.
//
// Solidity: event Whitelisted(address indexed operator)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicFilterer) WatchWhitelisted(opts *bind.WatchOpts, sink chan<- *ContractAVSGovernanceLogicWhitelisted, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAVSGovernanceLogic.contract.WatchLogs(opts, "Whitelisted", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAVSGovernanceLogicWhitelisted)
				if err := _ContractAVSGovernanceLogic.contract.UnpackLog(event, "Whitelisted", log); err != nil {
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

// ParseWhitelisted is a log parse operation binding the contract event 0xaab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a54.
//
// Solidity: event Whitelisted(address indexed operator)
func (_ContractAVSGovernanceLogic *ContractAVSGovernanceLogicFilterer) ParseWhitelisted(log types.Log) (*ContractAVSGovernanceLogicWhitelisted, error) {
	event := new(ContractAVSGovernanceLogicWhitelisted)
	if err := _ContractAVSGovernanceLogic.contract.UnpackLog(event, "Whitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
