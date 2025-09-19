// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractAvsGovernanceLogic

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

// ContractAvsGovernanceLogicMetaData contains all meta data concerning the ContractAvsGovernanceLogic contract.
var ContractAvsGovernanceLogicMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_endpoint\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_taskExecutionHub\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_dstEid\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_ownerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_avsGovernance\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidDelegate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEndpointCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LzTokenUnavailable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"}],\"name\":\"NoPeer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"}],\"name\":\"NotEnoughNative\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"OnlyEndpoint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"}],\"name\":\"OnlyPeer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"gasLimit\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"callValue\",\"type\":\"uint128\"}],\"name\":\"GasOptionsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"LowBalanceAlert\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"MessageFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"MessageSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"peer\",\"type\":\"bytes32\"}],\"name\":\"PeerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"UnWhitelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"WhitelistManagerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"WhitelistManagerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"Whitelisted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_operators\",\"type\":\"address[]\"}],\"name\":\"addToWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"\",\"type\":\"uint256[4]\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"afterOperatorRegistered\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"afterOperatorUnregistered\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"origin\",\"type\":\"tuple\"}],\"name\":\"allowInitializePath\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"avsGovernance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"\",\"type\":\"uint256[4]\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"beforeOperatorRegistered\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"beforeOperatorUnregistered\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"callValue\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dstEid\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endpoint\",\"outputs\":[{\"internalType\":\"contractILayerZeroEndpointV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasLimit\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"isComposeMsgSender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isWhitelistManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"_origin\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_guid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"lzReceive\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"nextNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oAppVersion\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"senderVersion\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"receiverVersion\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"}],\"name\":\"peers\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"peer\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_operators\",\"type\":\"address[]\"}],\"name\":\"removeFromWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegate\",\"type\":\"address\"}],\"name\":\"setDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_gasLimit\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_callValue\",\"type\":\"uint128\"}],\"name\":\"setGasOptions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_peer\",\"type\":\"bytes32\"}],\"name\":\"setPeer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_taskExecutionHub\",\"type\":\"address\"}],\"name\":\"setTaskExecutionHub\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isManager\",\"type\":\"bool\"}],\"name\":\"setWhitelistManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taskExecutionHub\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052600436101561001a575b3615610018575f80fd5b005b5f5f3560e01c806313137d651461114d57806317442b701461112c5780633035e4f9146110705780633400288b146110435780633af32abf14611006578063471f9e6414610fb8578063548db17414610ec65780635e280f1114610e82578063715018a614610e2b578063734ff5f314610deb5780637d25a05e14610dc75780637f64978314610cc057806382413eac14610c635780638d98e57914610c1f5780638da5cb5b14610bf857806394e7a3f014610b01578063989b147714610ac45780639dbe153814610a35578063bb0b6a5314610a00578063c54d346e146109d8578063ca5eb5e114610930578063d62457f61461090f578063e64cc9da14610859578063e9ecc1cb1461063e578063ed8ad8361461036d578063f2fde38b146102e7578063f3fef3a3146101c7578063f68016b7146101a05763ff7bd03d14610164575061000e565b3461019d57606036600319011261019d5760209060409063ffffffff610188611390565b16815260018352205460405190602435148152f35b80fd5b503461019d578060031936011261019d5760206001600160801b0360035416604051908152f35b503461019d57604036600319011261019d576004356001600160a01b038116908190036102e357602435906101fa611716565b80156102aa578115610265578147106102295782808093819382f11561021d5780f35b604051903d90823e3d90fd5b60405162461bcd60e51b8152602060048201526014602482015273496e73756666696369656e742062616c616e636560601b6044820152606490fd5b60405162461bcd60e51b815260206004820152601d60248201527f416d6f756e74206d7573742062652067726561746572207468616e20300000006044820152606490fd5b60405162461bcd60e51b8152602060048201526011602482015270125b9d985b1a59081c9958da5c1a595b9d607a1b6044820152606490fd5b5080fd5b503461019d57602036600319011261019d576103016112cf565b610309611716565b6001600160a01b031680156103595781546001600160a01b03198116821783556001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b631e4fbdf760e01b82526004829052602482fd5b503461019d5761037c3661134c565b506103b4915050337f000000000000000000000000875b5ff698b74b26f39c223c4996871f28acddea6001600160a01b0316146113a3565b6040519082602083015260018060a01b0316806040830152604082526103db606083611501565b61046e6103f76003546001600160801b038160801c9116611782565b7f00000000000000000000000000000000000000000000000000000000000075e89060018060a01b0360025416946040805161043281611499565b63ffffffff85169788825260208201528282820152836060820152886080820152815180968192631bb8518b60e31b83523090600484016115ad565b03817f0000000000000000000000001a44076050125825900e736c501f859c50fe728c6001600160a01b03165afa8794816105ff575b506105595761055560206105356031898b806104be6116a7565b927f2254f63be429412a67ef70cf2ef65d61a88074bfd112333e0a449a9e4a7683c8604051888152806104f38a820188611589565b0390a36040519485927002932b3b4b9ba32b9103330b4b632b21d1607d1b828501528051918291018585015e820190838201520301601f198101835282611501565b60405162461bcd60e51b8152602060048201529182916024830190611589565b0390fd5b8394959351600a8102818104600a14821517156105eb576020936105ba936105b36105ab7f0597c8258e620343bfa7f3cf82a7d6578378f08248fd958f32bcab61a1004ec3989560648c96049061161b565b47101561163c565b3093611861565b519351604051908152a37f4d0eb1f4bac8744fd2be119845e23b3befc88094b42bcda1204c65694a00f9e58280a280f35b634e487b7160e01b89526011600452602489fd5b90945060403d604011610637575b6106178183611501565b81016040828203126106335761062c91611561565b935f6104a4565b8880fd5b503d61060d565b503461019d57602036600319011261019d576106586112cf565b61068c337f000000000000000000000000875b5ff698b74b26f39c223c4996871f28acddea6001600160a01b0316146113a3565b60408051600160208201526001600160a01b039092168282018190529082526106b6606083611501565b6106d26103f76003546001600160801b038160801c9116611782565b03817f0000000000000000000000001a44076050125825900e736c501f859c50fe728c6001600160a01b03165afa87948161081e575b5061079b5761055560206105356033898b806107226116a7565b927f2254f63be429412a67ef70cf2ef65d61a88074bfd112333e0a449a9e4a7683c8604051888152806107578a820188611589565b0390a36040519485927202ab73932b3b4b9ba32b9103330b4b632b21d1606d1b828501528051918291018585015e820190838201520301601f198101835282611501565b8394959351600a8102818104600a14821517156105eb576020936107ed936105b36105ab7f0597c8258e620343bfa7f3cf82a7d6578378f08248fd958f32bcab61a1004ec3989560648c96049061161b565b519351604051908152a37f6f42117a557500c705ddf040a619d86f39101e6b74ac20d7b3e5943ba473fc7f8280a280f35b90945060403d604011610852575b6108368183611501565b81016040828203126106335761084b91611561565b935f610708565b503d61082c565b503461019d57604036600319011261019d576108736112cf565b6024358015159182820361090b57610889611716565b6001600160a01b03169161089e831515611523565b8284526005602052604084209060ff801983541691161790555f146108e4577f15a94be709049b6d6ac9d90d9a1ee40b18b9071a625486fad8ab1d45cedcc24d8280a280f35b7fce4df0698b95faea0f80f49bc274673a82e9cb758eb9df7c2951b3ab6732350f8280a280f35b8380fd5b503461019d578060031936011261019d57602060035460801c604051908152f35b50346109d45760203660031901126109d45761094a6112cf565b610952611716565b7f0000000000000000000000001a44076050125825900e736c501f859c50fe728c6001600160a01b031690813b156109d45760405163ca5eb5e160e01b81526001600160a01b039091166004820152905f908290602490829084905af180156109c9576109bd575080f35b61001891505f90611501565b6040513d5f823e3d90fd5b5f80fd5b346109d4575f3660031901126109d4576002546040516001600160a01b039091168152602090f35b346109d45760203660031901126109d45763ffffffff610a1e6112e5565b165f526001602052602060405f2054604051908152f35b346109d45760403660031901126109d4576004356001600160801b0381168091036109d4576024356001600160801b038116918282036109d4577f87768cdb9e7a69da85a21ad47c9654958b5cbc0447a55b66d0e3e0d640430c2d92604092610a9c611716565b60801b6fffffffffffffffffffffffffffffffff1916821760035582519182526020820152a1005b346109d45760203660031901126109d4576001600160a01b03610ae56112cf565b165f526005602052602060ff60405f2054166040519015158152f35b346109d457610b0f3661134c565b50610b47915050337f000000000000000000000000875b5ff698b74b26f39c223c4996871f28acddea6001600160a01b0316146113a3565b662386f26fc100004710610bbc575b6001600160a01b03165f9081526004602052604090205460ff1615610b7757005b60405162461bcd60e51b815260206004820152601b60248201527f4f70657261746f72206973206e6f742077686974656c697374656400000000006044820152606490fd5b7fb1b538a145d528cb9b8a428a3f47d5730430a03c2d98d7562357ee7db16dbd266040478151908152662386f26fc100006020820152a1610b56565b346109d4575f3660031901126109d4575f546040516001600160a01b039091168152602090f35b346109d4575f3660031901126109d4576040517f000000000000000000000000875b5ff698b74b26f39c223c4996871f28acddea6001600160a01b03168152602090f35b346109d457366003190160a081126109d4576060136109d45760643567ffffffffffffffff81116109d457610c9c90369060040161128b565b50506084356001600160a01b038116908190036109d4576020906040519030148152f35b346109d457610cce366112f8565b90335f52600560205260ff60405f2054168015610db4575b610cef90611402565b5f5b828110610cfa57005b6001600160a01b03610d15610d10838686611461565b611485565b1690610d22821515611523565b815f52600460205260ff60405f205416610d7957816001925f52600460205260405f208360ff198254161790557faab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a545f80a201610cf1565b60405162461bcd60e51b8152602060048201526013602482015272105b1c9958591e481dda1a5d195b1a5cdd1959606a1b6044820152606490fd5b505f546001600160a01b03163314610ce6565b346109d45760403660031901126109d457610de06112e5565b5060206040515f8152f35b346109d4575f3660031901126109d457602060405163ffffffff7f00000000000000000000000000000000000000000000000000000000000075e8168152f35b346109d4575f3660031901126109d457610e43611716565b5f80546001600160a01b0319811682556001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b346109d4575f3660031901126109d4576040517f0000000000000000000000001a44076050125825900e736c501f859c50fe728c6001600160a01b03168152602090f35b346109d457610ed4366112f8565b90335f52600560205260ff60405f2054168015610fa5575b610ef590611402565b5f5b828110610f0057005b6001600160a01b03610f16610d10838686611461565b1690815f52600460205260ff60405f20541615610f6e57816001925f52600460205260405f2060ff1981541690557f7cdb51b0cc2e541ad7e9471c358de415f5dbaff6cca78e3393d445346c610c1b5f80a201610ef7565b60405162461bcd60e51b815260206004820152600f60248201526e139bdd081dda1a5d195b1a5cdd1959608a1b6044820152606490fd5b505f546001600160a01b03163314610eec565b346109d45760203660031901126109d457610fd16112cf565b50610018337f000000000000000000000000875b5ff698b74b26f39c223c4996871f28acddea6001600160a01b0316146113a3565b346109d45760203660031901126109d4576001600160a01b036110276112cf565b165f526004602052602060ff60405f2054166040519015158152f35b346109d45760403660031901126109d45761001861105f6112e5565b611067611716565b6024359061173c565b346109d45760203660031901126109d4576110896112cf565b611091611716565b6001600160a01b031680156110dc57600280546001600160a01b03191682179055610018907f00000000000000000000000000000000000000000000000000000000000075e861173c565b60405162461bcd60e51b815260206004820152602260248201527f496e76616c6964207461736b20657865637574696f6e20687562206164647265604482015261737360f01b6064820152608490fd5b346109d4575f3660031901126109d457604080516001815260026020820152f35b366003190160e081126109d4576060136109d45760843567ffffffffffffffff81116109d45761118190369060040161128b565b505061118b6112b9565b5060c43567ffffffffffffffff81116109d4576111ac90369060040161128b565b5050337f0000000000000000000000001a44076050125825900e736c501f859c50fe728c6001600160a01b031603611278576111ee6111e9611390565b6116e6565b6024358091036112555760405162461bcd60e51b815260206004820152602f60248201527f417673476f7665726e616e63654c6f6769633a2073686f756c64206e6f74207260448201526e656365697665206d6573736167657360881b6064820152608490fd5b63ffffffff611262611390565b63309afaf360e21b5f521660045260245260445ffd5b6391ac5e4f60e01b5f523360045260245ffd5b9181601f840112156109d45782359167ffffffffffffffff83116109d457602083818601950101116109d457565b60a435906001600160a01b03821682036109d457565b600435906001600160a01b03821682036109d457565b6004359063ffffffff821682036109d457565b9060206003198301126109d45760043567ffffffffffffffff81116109d457826023820112156109d45780600401359267ffffffffffffffff84116109d45760248460051b830101116109d4576024019190565b60e06003198201126109d4576004356001600160a01b03811681036109d457916024359160c4116109d45760449060c4356001600160a01b03811681036109d45790565b60043563ffffffff811681036109d45790565b156113aa57565b60405162461bcd60e51b815260206004820152602a60248201527f4f6e6c792041565320476f7665726e616e63652063616e2063616c6c207468696044820152693990333ab731ba34b7b760b11b6064820152608490fd5b1561140957565b60405162461bcd60e51b815260206004820152602a60248201527f43616c6c6572206973206e6f7420612077686974656c697374206d616e616765604482015269391037b91037bbb732b960b11b6064820152608490fd5b91908110156114715760051b0190565b634e487b7160e01b5f52603260045260245ffd5b356001600160a01b03811681036109d45790565b60a0810190811067ffffffffffffffff8211176114b557604052565b634e487b7160e01b5f52604160045260245ffd5b6040810190811067ffffffffffffffff8211176114b557604052565b6060810190811067ffffffffffffffff8211176114b557604052565b90601f8019910116810190811067ffffffffffffffff8211176114b557604052565b1561152a57565b60405162461bcd60e51b815260206004820152600f60248201526e496e76616c6964206164647265737360881b6044820152606490fd5b91908260409103126109d457604051611579816114c9565b6020808294805184520151910152565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b906020909392936040835263ffffffff815116604084015281810151606084015260806116016115eb604084015160a08488015260e0870190611589565b6060840151868203603f190160a0880152611589565b910151151560c08401526001600160a01b03909416910152565b9190820180921161162857565b634e487b7160e01b5f52601160045260245ffd5b1561164357565b60405162461bcd60e51b815260206004820152603660248201527f496e73756666696369656e742062616c616e636520666f72206d6573736167656044820152752066656520287769746820313025206275666665722960501b6064820152608490fd5b3d156116e1573d9067ffffffffffffffff82116114b557604051916116d6601f8201601f191660200184611501565b82523d5f602084013e565b606090565b63ffffffff16805f52600160205260405f2054908115611704575090565b63f6ff4fb760e01b5f5260045260245ffd5b5f546001600160a01b0316330361172957565b63118cdaa760e01b5f523360045260245ffd5b7f238399d427b947898edb290f5ff0f9109849b1c3ba196a42e35f00c50a54b98b9163ffffffff6040921690815f52600160205280835f205582519182526020820152a1565b906001600160801b0381166118235750604051906001600160801b03199060801b166020820152601081526117b8603082611501565b805160018101809111611628576026611820916020604051948592600360f01b83850152600160f81b602285015261ffff60f01b9060f01b166023840152600160f81b60258401528051918291018484015e81015f838201520301601f198101835282611501565b90565b604051916001600160801b03199060801b1660208301526001600160801b03199060801b1660308201526020815261185c604082611501565b6117b8565b9160405161186e816114e5565b5f81525f60208201526040805191611885836114c9565b5f83525f60208401520152835193844710611acc57602001918251806119ae575b50918391608095936118bb61190498966116e6565b925115159263ffffffff604051956118d287611499565b168552602085015260408401526060830152848201526040518095819482936302637a4560e41b8452600484016115ad565b03917f0000000000000000000000001a44076050125825900e736c501f859c50fe728c6001600160a01b03165af19081156109c9575f91611943575090565b905060803d6080116119a7575b61195a8183611501565b8101906080818303126109d45760405191611974836114e5565b81518352602082015167ffffffffffffffff811681036109d457602084015261199f91604001611561565b604082015290565b503d611950565b60405163393f876560e21b81527f0000000000000000000000001a44076050125825900e736c501f859c50fe728c6001600160a01b031690602081600481855afa9081156109c9575f91611a8a575b506001600160a01b0316918215611a7b576020915f9160405190848201926323b872dd60e01b84523360248401526044830152606482015260648152611a44608482611501565b519082855af1156109c9575f513d611a725750803b155b156118a657635274afe760e01b5f5260045260245ffd5b60011415611a5b565b6329b99a9560e11b5f5260045ffd5b90506020813d602011611ac4575b81611aa560209383611501565b810103126109d457516001600160a01b03811681036109d4575f6119fd565b3d9150611a98565b60405162461bcd60e51b815260206004820152601d60248201527f496e73756666696369656e7420636f6e74726163742062616c616e63650000006044820152606490fdfea264697066735822122029d9e25a2fce216f47e2a036081d8c692ccc13786062f5f36218fc53ff3065d664736f6c634300081b0033",
}

// ContractAvsGovernanceLogicABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractAvsGovernanceLogicMetaData.ABI instead.
var ContractAvsGovernanceLogicABI = ContractAvsGovernanceLogicMetaData.ABI

// ContractAvsGovernanceLogicBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractAvsGovernanceLogicMetaData.Bin instead.
var ContractAvsGovernanceLogicBin = ContractAvsGovernanceLogicMetaData.Bin

// DeployContractAvsGovernanceLogic deploys a new Ethereum contract, binding an instance of ContractAvsGovernanceLogic to it.
func DeployContractAvsGovernanceLogic(auth *bind.TransactOpts, backend bind.ContractBackend, _endpoint common.Address, _taskExecutionHub common.Address, _dstEid uint32, _ownerAddress common.Address, _avsGovernance common.Address) (common.Address, *types.Transaction, *ContractAvsGovernanceLogic, error) {
	parsed, err := ContractAvsGovernanceLogicMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractAvsGovernanceLogicBin), backend, _endpoint, _taskExecutionHub, _dstEid, _ownerAddress, _avsGovernance)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractAvsGovernanceLogic{ContractAvsGovernanceLogicCaller: ContractAvsGovernanceLogicCaller{contract: contract}, ContractAvsGovernanceLogicTransactor: ContractAvsGovernanceLogicTransactor{contract: contract}, ContractAvsGovernanceLogicFilterer: ContractAvsGovernanceLogicFilterer{contract: contract}}, nil
}

// ContractAvsGovernanceLogicMethods is an auto generated interface around an Ethereum contract.
type ContractAvsGovernanceLogicMethods interface {
	ContractAvsGovernanceLogicCalls
	ContractAvsGovernanceLogicTransacts
	ContractAvsGovernanceLogicFilters
}

// ContractAvsGovernanceLogicCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractAvsGovernanceLogicCalls interface {
	AllowInitializePath(opts *bind.CallOpts, origin Origin) (bool, error)

	AvsGovernance(opts *bind.CallOpts) (common.Address, error)

	CallValue(opts *bind.CallOpts) (*big.Int, error)

	DstEid(opts *bind.CallOpts) (uint32, error)

	Endpoint(opts *bind.CallOpts) (common.Address, error)

	GasLimit(opts *bind.CallOpts) (*big.Int, error)

	IsComposeMsgSender(opts *bind.CallOpts, arg0 Origin, arg1 []byte, _sender common.Address) (bool, error)

	IsWhitelistManager(opts *bind.CallOpts, arg0 common.Address) (bool, error)

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

// ContractAvsGovernanceLogicTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractAvsGovernanceLogicTransacts interface {
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

	SetWhitelistManager(opts *bind.TransactOpts, _manager common.Address, _isManager bool) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error)

	Withdraw(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error)
}

// ContractAvsGovernanceLogicFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractAvsGovernanceLogicFilters interface {
	FilterGasOptionsUpdated(opts *bind.FilterOpts) (*ContractAvsGovernanceLogicGasOptionsUpdatedIterator, error)
	WatchGasOptionsUpdated(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceLogicGasOptionsUpdated) (event.Subscription, error)
	ParseGasOptionsUpdated(log types.Log) (*ContractAvsGovernanceLogicGasOptionsUpdated, error)

	FilterLowBalanceAlert(opts *bind.FilterOpts) (*ContractAvsGovernanceLogicLowBalanceAlertIterator, error)
	WatchLowBalanceAlert(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceLogicLowBalanceAlert) (event.Subscription, error)
	ParseLowBalanceAlert(log types.Log) (*ContractAvsGovernanceLogicLowBalanceAlert, error)

	FilterMessageFailed(opts *bind.FilterOpts, dstEid []uint32, guid [][32]byte) (*ContractAvsGovernanceLogicMessageFailedIterator, error)
	WatchMessageFailed(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceLogicMessageFailed, dstEid []uint32, guid [][32]byte) (event.Subscription, error)
	ParseMessageFailed(log types.Log) (*ContractAvsGovernanceLogicMessageFailed, error)

	FilterMessageSent(opts *bind.FilterOpts, dstEid []uint32, guid [][32]byte) (*ContractAvsGovernanceLogicMessageSentIterator, error)
	WatchMessageSent(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceLogicMessageSent, dstEid []uint32, guid [][32]byte) (event.Subscription, error)
	ParseMessageSent(log types.Log) (*ContractAvsGovernanceLogicMessageSent, error)

	FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address) (*ContractAvsGovernanceLogicOperatorRegisteredIterator, error)
	WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceLogicOperatorRegistered, operator []common.Address) (event.Subscription, error)
	ParseOperatorRegistered(log types.Log) (*ContractAvsGovernanceLogicOperatorRegistered, error)

	FilterOperatorUnregistered(opts *bind.FilterOpts, operator []common.Address) (*ContractAvsGovernanceLogicOperatorUnregisteredIterator, error)
	WatchOperatorUnregistered(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceLogicOperatorUnregistered, operator []common.Address) (event.Subscription, error)
	ParseOperatorUnregistered(log types.Log) (*ContractAvsGovernanceLogicOperatorUnregistered, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractAvsGovernanceLogicOwnershipTransferredIterator, error)
	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceLogicOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error)
	ParseOwnershipTransferred(log types.Log) (*ContractAvsGovernanceLogicOwnershipTransferred, error)

	FilterPeerSet(opts *bind.FilterOpts) (*ContractAvsGovernanceLogicPeerSetIterator, error)
	WatchPeerSet(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceLogicPeerSet) (event.Subscription, error)
	ParsePeerSet(log types.Log) (*ContractAvsGovernanceLogicPeerSet, error)

	FilterUnWhitelisted(opts *bind.FilterOpts, operator []common.Address) (*ContractAvsGovernanceLogicUnWhitelistedIterator, error)
	WatchUnWhitelisted(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceLogicUnWhitelisted, operator []common.Address) (event.Subscription, error)
	ParseUnWhitelisted(log types.Log) (*ContractAvsGovernanceLogicUnWhitelisted, error)

	FilterWhitelistManagerAdded(opts *bind.FilterOpts, manager []common.Address) (*ContractAvsGovernanceLogicWhitelistManagerAddedIterator, error)
	WatchWhitelistManagerAdded(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceLogicWhitelistManagerAdded, manager []common.Address) (event.Subscription, error)
	ParseWhitelistManagerAdded(log types.Log) (*ContractAvsGovernanceLogicWhitelistManagerAdded, error)

	FilterWhitelistManagerRemoved(opts *bind.FilterOpts, manager []common.Address) (*ContractAvsGovernanceLogicWhitelistManagerRemovedIterator, error)
	WatchWhitelistManagerRemoved(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceLogicWhitelistManagerRemoved, manager []common.Address) (event.Subscription, error)
	ParseWhitelistManagerRemoved(log types.Log) (*ContractAvsGovernanceLogicWhitelistManagerRemoved, error)

	FilterWhitelisted(opts *bind.FilterOpts, operator []common.Address) (*ContractAvsGovernanceLogicWhitelistedIterator, error)
	WatchWhitelisted(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceLogicWhitelisted, operator []common.Address) (event.Subscription, error)
	ParseWhitelisted(log types.Log) (*ContractAvsGovernanceLogicWhitelisted, error)
}

// ContractAvsGovernanceLogic is an auto generated Go binding around an Ethereum contract.
type ContractAvsGovernanceLogic struct {
	ContractAvsGovernanceLogicCaller     // Read-only binding to the contract
	ContractAvsGovernanceLogicTransactor // Write-only binding to the contract
	ContractAvsGovernanceLogicFilterer   // Log filterer for contract events
}

// ContractAvsGovernanceLogic implements the ContractAvsGovernanceLogicMethods interface.
var _ ContractAvsGovernanceLogicMethods = (*ContractAvsGovernanceLogic)(nil)

// ContractAvsGovernanceLogicCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractAvsGovernanceLogicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAvsGovernanceLogicCaller implements the ContractAvsGovernanceLogicCalls interface.
var _ ContractAvsGovernanceLogicCalls = (*ContractAvsGovernanceLogicCaller)(nil)

// ContractAvsGovernanceLogicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractAvsGovernanceLogicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAvsGovernanceLogicTransactor implements the ContractAvsGovernanceLogicTransacts interface.
var _ ContractAvsGovernanceLogicTransacts = (*ContractAvsGovernanceLogicTransactor)(nil)

// ContractAvsGovernanceLogicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractAvsGovernanceLogicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAvsGovernanceLogicFilterer implements the ContractAvsGovernanceLogicFilters interface.
var _ ContractAvsGovernanceLogicFilters = (*ContractAvsGovernanceLogicFilterer)(nil)

// ContractAvsGovernanceLogicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractAvsGovernanceLogicSession struct {
	Contract     *ContractAvsGovernanceLogic // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ContractAvsGovernanceLogicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractAvsGovernanceLogicCallerSession struct {
	Contract *ContractAvsGovernanceLogicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// ContractAvsGovernanceLogicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractAvsGovernanceLogicTransactorSession struct {
	Contract     *ContractAvsGovernanceLogicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// ContractAvsGovernanceLogicRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractAvsGovernanceLogicRaw struct {
	Contract *ContractAvsGovernanceLogic // Generic contract binding to access the raw methods on
}

// ContractAvsGovernanceLogicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractAvsGovernanceLogicCallerRaw struct {
	Contract *ContractAvsGovernanceLogicCaller // Generic read-only contract binding to access the raw methods on
}

// ContractAvsGovernanceLogicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractAvsGovernanceLogicTransactorRaw struct {
	Contract *ContractAvsGovernanceLogicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractAvsGovernanceLogic creates a new instance of ContractAvsGovernanceLogic, bound to a specific deployed contract.
func NewContractAvsGovernanceLogic(address common.Address, backend bind.ContractBackend) (*ContractAvsGovernanceLogic, error) {
	contract, err := bindContractAvsGovernanceLogic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceLogic{ContractAvsGovernanceLogicCaller: ContractAvsGovernanceLogicCaller{contract: contract}, ContractAvsGovernanceLogicTransactor: ContractAvsGovernanceLogicTransactor{contract: contract}, ContractAvsGovernanceLogicFilterer: ContractAvsGovernanceLogicFilterer{contract: contract}}, nil
}

// NewContractAvsGovernanceLogicCaller creates a new read-only instance of ContractAvsGovernanceLogic, bound to a specific deployed contract.
func NewContractAvsGovernanceLogicCaller(address common.Address, caller bind.ContractCaller) (*ContractAvsGovernanceLogicCaller, error) {
	contract, err := bindContractAvsGovernanceLogic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceLogicCaller{contract: contract}, nil
}

// NewContractAvsGovernanceLogicTransactor creates a new write-only instance of ContractAvsGovernanceLogic, bound to a specific deployed contract.
func NewContractAvsGovernanceLogicTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractAvsGovernanceLogicTransactor, error) {
	contract, err := bindContractAvsGovernanceLogic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceLogicTransactor{contract: contract}, nil
}

// NewContractAvsGovernanceLogicFilterer creates a new log filterer instance of ContractAvsGovernanceLogic, bound to a specific deployed contract.
func NewContractAvsGovernanceLogicFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractAvsGovernanceLogicFilterer, error) {
	contract, err := bindContractAvsGovernanceLogic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceLogicFilterer{contract: contract}, nil
}

// bindContractAvsGovernanceLogic binds a generic wrapper to an already deployed contract.
func bindContractAvsGovernanceLogic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractAvsGovernanceLogicMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractAvsGovernanceLogic.Contract.ContractAvsGovernanceLogicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.Contract.ContractAvsGovernanceLogicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.Contract.ContractAvsGovernanceLogicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractAvsGovernanceLogic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.Contract.contract.Transact(opts, method, params...)
}

// AllowInitializePath is a free data retrieval call binding the contract method 0xff7bd03d.
//
// Solidity: function allowInitializePath((uint32,bytes32,uint64) origin) view returns(bool)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicCaller) AllowInitializePath(opts *bind.CallOpts, origin Origin) (bool, error) {
	var out []interface{}
	err := _ContractAvsGovernanceLogic.contract.Call(opts, &out, "allowInitializePath", origin)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowInitializePath is a free data retrieval call binding the contract method 0xff7bd03d.
//
// Solidity: function allowInitializePath((uint32,bytes32,uint64) origin) view returns(bool)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicSession) AllowInitializePath(origin Origin) (bool, error) {
	return _ContractAvsGovernanceLogic.Contract.AllowInitializePath(&_ContractAvsGovernanceLogic.CallOpts, origin)
}

// AllowInitializePath is a free data retrieval call binding the contract method 0xff7bd03d.
//
// Solidity: function allowInitializePath((uint32,bytes32,uint64) origin) view returns(bool)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicCallerSession) AllowInitializePath(origin Origin) (bool, error) {
	return _ContractAvsGovernanceLogic.Contract.AllowInitializePath(&_ContractAvsGovernanceLogic.CallOpts, origin)
}

// AvsGovernance is a free data retrieval call binding the contract method 0x8d98e579.
//
// Solidity: function avsGovernance() view returns(address)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicCaller) AvsGovernance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAvsGovernanceLogic.contract.Call(opts, &out, "avsGovernance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsGovernance is a free data retrieval call binding the contract method 0x8d98e579.
//
// Solidity: function avsGovernance() view returns(address)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicSession) AvsGovernance() (common.Address, error) {
	return _ContractAvsGovernanceLogic.Contract.AvsGovernance(&_ContractAvsGovernanceLogic.CallOpts)
}

// AvsGovernance is a free data retrieval call binding the contract method 0x8d98e579.
//
// Solidity: function avsGovernance() view returns(address)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicCallerSession) AvsGovernance() (common.Address, error) {
	return _ContractAvsGovernanceLogic.Contract.AvsGovernance(&_ContractAvsGovernanceLogic.CallOpts)
}

// CallValue is a free data retrieval call binding the contract method 0xd62457f6.
//
// Solidity: function callValue() view returns(uint128)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicCaller) CallValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractAvsGovernanceLogic.contract.Call(opts, &out, "callValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CallValue is a free data retrieval call binding the contract method 0xd62457f6.
//
// Solidity: function callValue() view returns(uint128)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicSession) CallValue() (*big.Int, error) {
	return _ContractAvsGovernanceLogic.Contract.CallValue(&_ContractAvsGovernanceLogic.CallOpts)
}

// CallValue is a free data retrieval call binding the contract method 0xd62457f6.
//
// Solidity: function callValue() view returns(uint128)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicCallerSession) CallValue() (*big.Int, error) {
	return _ContractAvsGovernanceLogic.Contract.CallValue(&_ContractAvsGovernanceLogic.CallOpts)
}

// DstEid is a free data retrieval call binding the contract method 0x734ff5f3.
//
// Solidity: function dstEid() view returns(uint32)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicCaller) DstEid(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractAvsGovernanceLogic.contract.Call(opts, &out, "dstEid")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// DstEid is a free data retrieval call binding the contract method 0x734ff5f3.
//
// Solidity: function dstEid() view returns(uint32)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicSession) DstEid() (uint32, error) {
	return _ContractAvsGovernanceLogic.Contract.DstEid(&_ContractAvsGovernanceLogic.CallOpts)
}

// DstEid is a free data retrieval call binding the contract method 0x734ff5f3.
//
// Solidity: function dstEid() view returns(uint32)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicCallerSession) DstEid() (uint32, error) {
	return _ContractAvsGovernanceLogic.Contract.DstEid(&_ContractAvsGovernanceLogic.CallOpts)
}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicCaller) Endpoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAvsGovernanceLogic.contract.Call(opts, &out, "endpoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicSession) Endpoint() (common.Address, error) {
	return _ContractAvsGovernanceLogic.Contract.Endpoint(&_ContractAvsGovernanceLogic.CallOpts)
}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicCallerSession) Endpoint() (common.Address, error) {
	return _ContractAvsGovernanceLogic.Contract.Endpoint(&_ContractAvsGovernanceLogic.CallOpts)
}

// GasLimit is a free data retrieval call binding the contract method 0xf68016b7.
//
// Solidity: function gasLimit() view returns(uint128)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicCaller) GasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractAvsGovernanceLogic.contract.Call(opts, &out, "gasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GasLimit is a free data retrieval call binding the contract method 0xf68016b7.
//
// Solidity: function gasLimit() view returns(uint128)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicSession) GasLimit() (*big.Int, error) {
	return _ContractAvsGovernanceLogic.Contract.GasLimit(&_ContractAvsGovernanceLogic.CallOpts)
}

// GasLimit is a free data retrieval call binding the contract method 0xf68016b7.
//
// Solidity: function gasLimit() view returns(uint128)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicCallerSession) GasLimit() (*big.Int, error) {
	return _ContractAvsGovernanceLogic.Contract.GasLimit(&_ContractAvsGovernanceLogic.CallOpts)
}

// IsComposeMsgSender is a free data retrieval call binding the contract method 0x82413eac.
//
// Solidity: function isComposeMsgSender((uint32,bytes32,uint64) , bytes , address _sender) view returns(bool)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicCaller) IsComposeMsgSender(opts *bind.CallOpts, arg0 Origin, arg1 []byte, _sender common.Address) (bool, error) {
	var out []interface{}
	err := _ContractAvsGovernanceLogic.contract.Call(opts, &out, "isComposeMsgSender", arg0, arg1, _sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsComposeMsgSender is a free data retrieval call binding the contract method 0x82413eac.
//
// Solidity: function isComposeMsgSender((uint32,bytes32,uint64) , bytes , address _sender) view returns(bool)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicSession) IsComposeMsgSender(arg0 Origin, arg1 []byte, _sender common.Address) (bool, error) {
	return _ContractAvsGovernanceLogic.Contract.IsComposeMsgSender(&_ContractAvsGovernanceLogic.CallOpts, arg0, arg1, _sender)
}

// IsComposeMsgSender is a free data retrieval call binding the contract method 0x82413eac.
//
// Solidity: function isComposeMsgSender((uint32,bytes32,uint64) , bytes , address _sender) view returns(bool)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicCallerSession) IsComposeMsgSender(arg0 Origin, arg1 []byte, _sender common.Address) (bool, error) {
	return _ContractAvsGovernanceLogic.Contract.IsComposeMsgSender(&_ContractAvsGovernanceLogic.CallOpts, arg0, arg1, _sender)
}

// IsWhitelistManager is a free data retrieval call binding the contract method 0x989b1477.
//
// Solidity: function isWhitelistManager(address ) view returns(bool)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicCaller) IsWhitelistManager(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ContractAvsGovernanceLogic.contract.Call(opts, &out, "isWhitelistManager", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelistManager is a free data retrieval call binding the contract method 0x989b1477.
//
// Solidity: function isWhitelistManager(address ) view returns(bool)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicSession) IsWhitelistManager(arg0 common.Address) (bool, error) {
	return _ContractAvsGovernanceLogic.Contract.IsWhitelistManager(&_ContractAvsGovernanceLogic.CallOpts, arg0)
}

// IsWhitelistManager is a free data retrieval call binding the contract method 0x989b1477.
//
// Solidity: function isWhitelistManager(address ) view returns(bool)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicCallerSession) IsWhitelistManager(arg0 common.Address) (bool, error) {
	return _ContractAvsGovernanceLogic.Contract.IsWhitelistManager(&_ContractAvsGovernanceLogic.CallOpts, arg0)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address ) view returns(bool)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicCaller) IsWhitelisted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ContractAvsGovernanceLogic.contract.Call(opts, &out, "isWhitelisted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address ) view returns(bool)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicSession) IsWhitelisted(arg0 common.Address) (bool, error) {
	return _ContractAvsGovernanceLogic.Contract.IsWhitelisted(&_ContractAvsGovernanceLogic.CallOpts, arg0)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address ) view returns(bool)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicCallerSession) IsWhitelisted(arg0 common.Address) (bool, error) {
	return _ContractAvsGovernanceLogic.Contract.IsWhitelisted(&_ContractAvsGovernanceLogic.CallOpts, arg0)
}

// NextNonce is a free data retrieval call binding the contract method 0x7d25a05e.
//
// Solidity: function nextNonce(uint32 , bytes32 ) view returns(uint64 nonce)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicCaller) NextNonce(opts *bind.CallOpts, arg0 uint32, arg1 [32]byte) (uint64, error) {
	var out []interface{}
	err := _ContractAvsGovernanceLogic.contract.Call(opts, &out, "nextNonce", arg0, arg1)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NextNonce is a free data retrieval call binding the contract method 0x7d25a05e.
//
// Solidity: function nextNonce(uint32 , bytes32 ) view returns(uint64 nonce)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicSession) NextNonce(arg0 uint32, arg1 [32]byte) (uint64, error) {
	return _ContractAvsGovernanceLogic.Contract.NextNonce(&_ContractAvsGovernanceLogic.CallOpts, arg0, arg1)
}

// NextNonce is a free data retrieval call binding the contract method 0x7d25a05e.
//
// Solidity: function nextNonce(uint32 , bytes32 ) view returns(uint64 nonce)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicCallerSession) NextNonce(arg0 uint32, arg1 [32]byte) (uint64, error) {
	return _ContractAvsGovernanceLogic.Contract.NextNonce(&_ContractAvsGovernanceLogic.CallOpts, arg0, arg1)
}

// OAppVersion is a free data retrieval call binding the contract method 0x17442b70.
//
// Solidity: function oAppVersion() pure returns(uint64 senderVersion, uint64 receiverVersion)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicCaller) OAppVersion(opts *bind.CallOpts) (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	var out []interface{}
	err := _ContractAvsGovernanceLogic.contract.Call(opts, &out, "oAppVersion")

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
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicSession) OAppVersion() (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	return _ContractAvsGovernanceLogic.Contract.OAppVersion(&_ContractAvsGovernanceLogic.CallOpts)
}

// OAppVersion is a free data retrieval call binding the contract method 0x17442b70.
//
// Solidity: function oAppVersion() pure returns(uint64 senderVersion, uint64 receiverVersion)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicCallerSession) OAppVersion() (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	return _ContractAvsGovernanceLogic.Contract.OAppVersion(&_ContractAvsGovernanceLogic.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAvsGovernanceLogic.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicSession) Owner() (common.Address, error) {
	return _ContractAvsGovernanceLogic.Contract.Owner(&_ContractAvsGovernanceLogic.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicCallerSession) Owner() (common.Address, error) {
	return _ContractAvsGovernanceLogic.Contract.Owner(&_ContractAvsGovernanceLogic.CallOpts)
}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 eid) view returns(bytes32 peer)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicCaller) Peers(opts *bind.CallOpts, eid uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractAvsGovernanceLogic.contract.Call(opts, &out, "peers", eid)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 eid) view returns(bytes32 peer)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicSession) Peers(eid uint32) ([32]byte, error) {
	return _ContractAvsGovernanceLogic.Contract.Peers(&_ContractAvsGovernanceLogic.CallOpts, eid)
}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 eid) view returns(bytes32 peer)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicCallerSession) Peers(eid uint32) ([32]byte, error) {
	return _ContractAvsGovernanceLogic.Contract.Peers(&_ContractAvsGovernanceLogic.CallOpts, eid)
}

// TaskExecutionHub is a free data retrieval call binding the contract method 0xc54d346e.
//
// Solidity: function taskExecutionHub() view returns(address)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicCaller) TaskExecutionHub(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAvsGovernanceLogic.contract.Call(opts, &out, "taskExecutionHub")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TaskExecutionHub is a free data retrieval call binding the contract method 0xc54d346e.
//
// Solidity: function taskExecutionHub() view returns(address)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicSession) TaskExecutionHub() (common.Address, error) {
	return _ContractAvsGovernanceLogic.Contract.TaskExecutionHub(&_ContractAvsGovernanceLogic.CallOpts)
}

// TaskExecutionHub is a free data retrieval call binding the contract method 0xc54d346e.
//
// Solidity: function taskExecutionHub() view returns(address)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicCallerSession) TaskExecutionHub() (common.Address, error) {
	return _ContractAvsGovernanceLogic.Contract.TaskExecutionHub(&_ContractAvsGovernanceLogic.CallOpts)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0x7f649783.
//
// Solidity: function addToWhitelist(address[] _operators) returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicTransactor) AddToWhitelist(opts *bind.TransactOpts, _operators []common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.contract.Transact(opts, "addToWhitelist", _operators)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0x7f649783.
//
// Solidity: function addToWhitelist(address[] _operators) returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicSession) AddToWhitelist(_operators []common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.Contract.AddToWhitelist(&_ContractAvsGovernanceLogic.TransactOpts, _operators)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0x7f649783.
//
// Solidity: function addToWhitelist(address[] _operators) returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicTransactorSession) AddToWhitelist(_operators []common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.Contract.AddToWhitelist(&_ContractAvsGovernanceLogic.TransactOpts, _operators)
}

// AfterOperatorRegistered is a paid mutator transaction binding the contract method 0xed8ad836.
//
// Solidity: function afterOperatorRegistered(address _operator, uint256 , uint256[4] , address ) returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicTransactor) AfterOperatorRegistered(opts *bind.TransactOpts, _operator common.Address, arg1 *big.Int, arg2 [4]*big.Int, arg3 common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.contract.Transact(opts, "afterOperatorRegistered", _operator, arg1, arg2, arg3)
}

// AfterOperatorRegistered is a paid mutator transaction binding the contract method 0xed8ad836.
//
// Solidity: function afterOperatorRegistered(address _operator, uint256 , uint256[4] , address ) returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicSession) AfterOperatorRegistered(_operator common.Address, arg1 *big.Int, arg2 [4]*big.Int, arg3 common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.Contract.AfterOperatorRegistered(&_ContractAvsGovernanceLogic.TransactOpts, _operator, arg1, arg2, arg3)
}

// AfterOperatorRegistered is a paid mutator transaction binding the contract method 0xed8ad836.
//
// Solidity: function afterOperatorRegistered(address _operator, uint256 , uint256[4] , address ) returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicTransactorSession) AfterOperatorRegistered(_operator common.Address, arg1 *big.Int, arg2 [4]*big.Int, arg3 common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.Contract.AfterOperatorRegistered(&_ContractAvsGovernanceLogic.TransactOpts, _operator, arg1, arg2, arg3)
}

// AfterOperatorUnregistered is a paid mutator transaction binding the contract method 0xe9ecc1cb.
//
// Solidity: function afterOperatorUnregistered(address _operator) returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicTransactor) AfterOperatorUnregistered(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.contract.Transact(opts, "afterOperatorUnregistered", _operator)
}

// AfterOperatorUnregistered is a paid mutator transaction binding the contract method 0xe9ecc1cb.
//
// Solidity: function afterOperatorUnregistered(address _operator) returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicSession) AfterOperatorUnregistered(_operator common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.Contract.AfterOperatorUnregistered(&_ContractAvsGovernanceLogic.TransactOpts, _operator)
}

// AfterOperatorUnregistered is a paid mutator transaction binding the contract method 0xe9ecc1cb.
//
// Solidity: function afterOperatorUnregistered(address _operator) returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicTransactorSession) AfterOperatorUnregistered(_operator common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.Contract.AfterOperatorUnregistered(&_ContractAvsGovernanceLogic.TransactOpts, _operator)
}

// BeforeOperatorRegistered is a paid mutator transaction binding the contract method 0x94e7a3f0.
//
// Solidity: function beforeOperatorRegistered(address _operator, uint256 , uint256[4] , address ) returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicTransactor) BeforeOperatorRegistered(opts *bind.TransactOpts, _operator common.Address, arg1 *big.Int, arg2 [4]*big.Int, arg3 common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.contract.Transact(opts, "beforeOperatorRegistered", _operator, arg1, arg2, arg3)
}

// BeforeOperatorRegistered is a paid mutator transaction binding the contract method 0x94e7a3f0.
//
// Solidity: function beforeOperatorRegistered(address _operator, uint256 , uint256[4] , address ) returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicSession) BeforeOperatorRegistered(_operator common.Address, arg1 *big.Int, arg2 [4]*big.Int, arg3 common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.Contract.BeforeOperatorRegistered(&_ContractAvsGovernanceLogic.TransactOpts, _operator, arg1, arg2, arg3)
}

// BeforeOperatorRegistered is a paid mutator transaction binding the contract method 0x94e7a3f0.
//
// Solidity: function beforeOperatorRegistered(address _operator, uint256 , uint256[4] , address ) returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicTransactorSession) BeforeOperatorRegistered(_operator common.Address, arg1 *big.Int, arg2 [4]*big.Int, arg3 common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.Contract.BeforeOperatorRegistered(&_ContractAvsGovernanceLogic.TransactOpts, _operator, arg1, arg2, arg3)
}

// BeforeOperatorUnregistered is a paid mutator transaction binding the contract method 0x471f9e64.
//
// Solidity: function beforeOperatorUnregistered(address ) returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicTransactor) BeforeOperatorUnregistered(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.contract.Transact(opts, "beforeOperatorUnregistered", arg0)
}

// BeforeOperatorUnregistered is a paid mutator transaction binding the contract method 0x471f9e64.
//
// Solidity: function beforeOperatorUnregistered(address ) returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicSession) BeforeOperatorUnregistered(arg0 common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.Contract.BeforeOperatorUnregistered(&_ContractAvsGovernanceLogic.TransactOpts, arg0)
}

// BeforeOperatorUnregistered is a paid mutator transaction binding the contract method 0x471f9e64.
//
// Solidity: function beforeOperatorUnregistered(address ) returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicTransactorSession) BeforeOperatorUnregistered(arg0 common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.Contract.BeforeOperatorUnregistered(&_ContractAvsGovernanceLogic.TransactOpts, arg0)
}

// LzReceive is a paid mutator transaction binding the contract method 0x13137d65.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicTransactor) LzReceive(opts *bind.TransactOpts, _origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.contract.Transact(opts, "lzReceive", _origin, _guid, _message, _executor, _extraData)
}

// LzReceive is a paid mutator transaction binding the contract method 0x13137d65.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicSession) LzReceive(_origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.Contract.LzReceive(&_ContractAvsGovernanceLogic.TransactOpts, _origin, _guid, _message, _executor, _extraData)
}

// LzReceive is a paid mutator transaction binding the contract method 0x13137d65.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicTransactorSession) LzReceive(_origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.Contract.LzReceive(&_ContractAvsGovernanceLogic.TransactOpts, _origin, _guid, _message, _executor, _extraData)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x548db174.
//
// Solidity: function removeFromWhitelist(address[] _operators) returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicTransactor) RemoveFromWhitelist(opts *bind.TransactOpts, _operators []common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.contract.Transact(opts, "removeFromWhitelist", _operators)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x548db174.
//
// Solidity: function removeFromWhitelist(address[] _operators) returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicSession) RemoveFromWhitelist(_operators []common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.Contract.RemoveFromWhitelist(&_ContractAvsGovernanceLogic.TransactOpts, _operators)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x548db174.
//
// Solidity: function removeFromWhitelist(address[] _operators) returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicTransactorSession) RemoveFromWhitelist(_operators []common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.Contract.RemoveFromWhitelist(&_ContractAvsGovernanceLogic.TransactOpts, _operators)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.Contract.RenounceOwnership(&_ContractAvsGovernanceLogic.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.Contract.RenounceOwnership(&_ContractAvsGovernanceLogic.TransactOpts)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicTransactor) SetDelegate(opts *bind.TransactOpts, _delegate common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.contract.Transact(opts, "setDelegate", _delegate)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicSession) SetDelegate(_delegate common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.Contract.SetDelegate(&_ContractAvsGovernanceLogic.TransactOpts, _delegate)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicTransactorSession) SetDelegate(_delegate common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.Contract.SetDelegate(&_ContractAvsGovernanceLogic.TransactOpts, _delegate)
}

// SetGasOptions is a paid mutator transaction binding the contract method 0x9dbe1538.
//
// Solidity: function setGasOptions(uint128 _gasLimit, uint128 _callValue) returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicTransactor) SetGasOptions(opts *bind.TransactOpts, _gasLimit *big.Int, _callValue *big.Int) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.contract.Transact(opts, "setGasOptions", _gasLimit, _callValue)
}

// SetGasOptions is a paid mutator transaction binding the contract method 0x9dbe1538.
//
// Solidity: function setGasOptions(uint128 _gasLimit, uint128 _callValue) returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicSession) SetGasOptions(_gasLimit *big.Int, _callValue *big.Int) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.Contract.SetGasOptions(&_ContractAvsGovernanceLogic.TransactOpts, _gasLimit, _callValue)
}

// SetGasOptions is a paid mutator transaction binding the contract method 0x9dbe1538.
//
// Solidity: function setGasOptions(uint128 _gasLimit, uint128 _callValue) returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicTransactorSession) SetGasOptions(_gasLimit *big.Int, _callValue *big.Int) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.Contract.SetGasOptions(&_ContractAvsGovernanceLogic.TransactOpts, _gasLimit, _callValue)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicTransactor) SetPeer(opts *bind.TransactOpts, _eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.contract.Transact(opts, "setPeer", _eid, _peer)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicSession) SetPeer(_eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.Contract.SetPeer(&_ContractAvsGovernanceLogic.TransactOpts, _eid, _peer)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicTransactorSession) SetPeer(_eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.Contract.SetPeer(&_ContractAvsGovernanceLogic.TransactOpts, _eid, _peer)
}

// SetTaskExecutionHub is a paid mutator transaction binding the contract method 0x3035e4f9.
//
// Solidity: function setTaskExecutionHub(address _taskExecutionHub) returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicTransactor) SetTaskExecutionHub(opts *bind.TransactOpts, _taskExecutionHub common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.contract.Transact(opts, "setTaskExecutionHub", _taskExecutionHub)
}

// SetTaskExecutionHub is a paid mutator transaction binding the contract method 0x3035e4f9.
//
// Solidity: function setTaskExecutionHub(address _taskExecutionHub) returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicSession) SetTaskExecutionHub(_taskExecutionHub common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.Contract.SetTaskExecutionHub(&_ContractAvsGovernanceLogic.TransactOpts, _taskExecutionHub)
}

// SetTaskExecutionHub is a paid mutator transaction binding the contract method 0x3035e4f9.
//
// Solidity: function setTaskExecutionHub(address _taskExecutionHub) returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicTransactorSession) SetTaskExecutionHub(_taskExecutionHub common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.Contract.SetTaskExecutionHub(&_ContractAvsGovernanceLogic.TransactOpts, _taskExecutionHub)
}

// SetWhitelistManager is a paid mutator transaction binding the contract method 0xe64cc9da.
//
// Solidity: function setWhitelistManager(address _manager, bool _isManager) returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicTransactor) SetWhitelistManager(opts *bind.TransactOpts, _manager common.Address, _isManager bool) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.contract.Transact(opts, "setWhitelistManager", _manager, _isManager)
}

// SetWhitelistManager is a paid mutator transaction binding the contract method 0xe64cc9da.
//
// Solidity: function setWhitelistManager(address _manager, bool _isManager) returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicSession) SetWhitelistManager(_manager common.Address, _isManager bool) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.Contract.SetWhitelistManager(&_ContractAvsGovernanceLogic.TransactOpts, _manager, _isManager)
}

// SetWhitelistManager is a paid mutator transaction binding the contract method 0xe64cc9da.
//
// Solidity: function setWhitelistManager(address _manager, bool _isManager) returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicTransactorSession) SetWhitelistManager(_manager common.Address, _isManager bool) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.Contract.SetWhitelistManager(&_ContractAvsGovernanceLogic.TransactOpts, _manager, _isManager)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.Contract.TransferOwnership(&_ContractAvsGovernanceLogic.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.Contract.TransferOwnership(&_ContractAvsGovernanceLogic.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _to, uint256 _amount) returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicTransactor) Withdraw(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.contract.Transact(opts, "withdraw", _to, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _to, uint256 _amount) returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicSession) Withdraw(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.Contract.Withdraw(&_ContractAvsGovernanceLogic.TransactOpts, _to, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _to, uint256 _amount) returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicTransactorSession) Withdraw(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.Contract.Withdraw(&_ContractAvsGovernanceLogic.TransactOpts, _to, _amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicSession) Receive() (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.Contract.Receive(&_ContractAvsGovernanceLogic.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicTransactorSession) Receive() (*types.Transaction, error) {
	return _ContractAvsGovernanceLogic.Contract.Receive(&_ContractAvsGovernanceLogic.TransactOpts)
}

// ContractAvsGovernanceLogicGasOptionsUpdatedIterator is returned from FilterGasOptionsUpdated and is used to iterate over the raw logs and unpacked data for GasOptionsUpdated events raised by the ContractAvsGovernanceLogic contract.
type ContractAvsGovernanceLogicGasOptionsUpdatedIterator struct {
	Event *ContractAvsGovernanceLogicGasOptionsUpdated // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceLogicGasOptionsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceLogicGasOptionsUpdated)
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
		it.Event = new(ContractAvsGovernanceLogicGasOptionsUpdated)
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
func (it *ContractAvsGovernanceLogicGasOptionsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceLogicGasOptionsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceLogicGasOptionsUpdated represents a GasOptionsUpdated event raised by the ContractAvsGovernanceLogic contract.
type ContractAvsGovernanceLogicGasOptionsUpdated struct {
	GasLimit  *big.Int
	CallValue *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterGasOptionsUpdated is a free log retrieval operation binding the contract event 0x87768cdb9e7a69da85a21ad47c9654958b5cbc0447a55b66d0e3e0d640430c2d.
//
// Solidity: event GasOptionsUpdated(uint128 gasLimit, uint128 callValue)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicFilterer) FilterGasOptionsUpdated(opts *bind.FilterOpts) (*ContractAvsGovernanceLogicGasOptionsUpdatedIterator, error) {

	logs, sub, err := _ContractAvsGovernanceLogic.contract.FilterLogs(opts, "GasOptionsUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceLogicGasOptionsUpdatedIterator{contract: _ContractAvsGovernanceLogic.contract, event: "GasOptionsUpdated", logs: logs, sub: sub}, nil
}

// WatchGasOptionsUpdated is a free log subscription operation binding the contract event 0x87768cdb9e7a69da85a21ad47c9654958b5cbc0447a55b66d0e3e0d640430c2d.
//
// Solidity: event GasOptionsUpdated(uint128 gasLimit, uint128 callValue)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicFilterer) WatchGasOptionsUpdated(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceLogicGasOptionsUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernanceLogic.contract.WatchLogs(opts, "GasOptionsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceLogicGasOptionsUpdated)
				if err := _ContractAvsGovernanceLogic.contract.UnpackLog(event, "GasOptionsUpdated", log); err != nil {
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
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicFilterer) ParseGasOptionsUpdated(log types.Log) (*ContractAvsGovernanceLogicGasOptionsUpdated, error) {
	event := new(ContractAvsGovernanceLogicGasOptionsUpdated)
	if err := _ContractAvsGovernanceLogic.contract.UnpackLog(event, "GasOptionsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceLogicLowBalanceAlertIterator is returned from FilterLowBalanceAlert and is used to iterate over the raw logs and unpacked data for LowBalanceAlert events raised by the ContractAvsGovernanceLogic contract.
type ContractAvsGovernanceLogicLowBalanceAlertIterator struct {
	Event *ContractAvsGovernanceLogicLowBalanceAlert // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceLogicLowBalanceAlertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceLogicLowBalanceAlert)
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
		it.Event = new(ContractAvsGovernanceLogicLowBalanceAlert)
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
func (it *ContractAvsGovernanceLogicLowBalanceAlertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceLogicLowBalanceAlertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceLogicLowBalanceAlert represents a LowBalanceAlert event raised by the ContractAvsGovernanceLogic contract.
type ContractAvsGovernanceLogicLowBalanceAlert struct {
	CurrentBalance *big.Int
	Threshold      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLowBalanceAlert is a free log retrieval operation binding the contract event 0xb1b538a145d528cb9b8a428a3f47d5730430a03c2d98d7562357ee7db16dbd26.
//
// Solidity: event LowBalanceAlert(uint256 currentBalance, uint256 threshold)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicFilterer) FilterLowBalanceAlert(opts *bind.FilterOpts) (*ContractAvsGovernanceLogicLowBalanceAlertIterator, error) {

	logs, sub, err := _ContractAvsGovernanceLogic.contract.FilterLogs(opts, "LowBalanceAlert")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceLogicLowBalanceAlertIterator{contract: _ContractAvsGovernanceLogic.contract, event: "LowBalanceAlert", logs: logs, sub: sub}, nil
}

// WatchLowBalanceAlert is a free log subscription operation binding the contract event 0xb1b538a145d528cb9b8a428a3f47d5730430a03c2d98d7562357ee7db16dbd26.
//
// Solidity: event LowBalanceAlert(uint256 currentBalance, uint256 threshold)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicFilterer) WatchLowBalanceAlert(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceLogicLowBalanceAlert) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernanceLogic.contract.WatchLogs(opts, "LowBalanceAlert")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceLogicLowBalanceAlert)
				if err := _ContractAvsGovernanceLogic.contract.UnpackLog(event, "LowBalanceAlert", log); err != nil {
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
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicFilterer) ParseLowBalanceAlert(log types.Log) (*ContractAvsGovernanceLogicLowBalanceAlert, error) {
	event := new(ContractAvsGovernanceLogicLowBalanceAlert)
	if err := _ContractAvsGovernanceLogic.contract.UnpackLog(event, "LowBalanceAlert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceLogicMessageFailedIterator is returned from FilterMessageFailed and is used to iterate over the raw logs and unpacked data for MessageFailed events raised by the ContractAvsGovernanceLogic contract.
type ContractAvsGovernanceLogicMessageFailedIterator struct {
	Event *ContractAvsGovernanceLogicMessageFailed // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceLogicMessageFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceLogicMessageFailed)
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
		it.Event = new(ContractAvsGovernanceLogicMessageFailed)
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
func (it *ContractAvsGovernanceLogicMessageFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceLogicMessageFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceLogicMessageFailed represents a MessageFailed event raised by the ContractAvsGovernanceLogic contract.
type ContractAvsGovernanceLogicMessageFailed struct {
	DstEid uint32
	Guid   [32]byte
	Reason []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMessageFailed is a free log retrieval operation binding the contract event 0x2254f63be429412a67ef70cf2ef65d61a88074bfd112333e0a449a9e4a7683c8.
//
// Solidity: event MessageFailed(uint32 indexed dstEid, bytes32 indexed guid, bytes reason)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicFilterer) FilterMessageFailed(opts *bind.FilterOpts, dstEid []uint32, guid [][32]byte) (*ContractAvsGovernanceLogicMessageFailedIterator, error) {

	var dstEidRule []interface{}
	for _, dstEidItem := range dstEid {
		dstEidRule = append(dstEidRule, dstEidItem)
	}
	var guidRule []interface{}
	for _, guidItem := range guid {
		guidRule = append(guidRule, guidItem)
	}

	logs, sub, err := _ContractAvsGovernanceLogic.contract.FilterLogs(opts, "MessageFailed", dstEidRule, guidRule)
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceLogicMessageFailedIterator{contract: _ContractAvsGovernanceLogic.contract, event: "MessageFailed", logs: logs, sub: sub}, nil
}

// WatchMessageFailed is a free log subscription operation binding the contract event 0x2254f63be429412a67ef70cf2ef65d61a88074bfd112333e0a449a9e4a7683c8.
//
// Solidity: event MessageFailed(uint32 indexed dstEid, bytes32 indexed guid, bytes reason)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicFilterer) WatchMessageFailed(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceLogicMessageFailed, dstEid []uint32, guid [][32]byte) (event.Subscription, error) {

	var dstEidRule []interface{}
	for _, dstEidItem := range dstEid {
		dstEidRule = append(dstEidRule, dstEidItem)
	}
	var guidRule []interface{}
	for _, guidItem := range guid {
		guidRule = append(guidRule, guidItem)
	}

	logs, sub, err := _ContractAvsGovernanceLogic.contract.WatchLogs(opts, "MessageFailed", dstEidRule, guidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceLogicMessageFailed)
				if err := _ContractAvsGovernanceLogic.contract.UnpackLog(event, "MessageFailed", log); err != nil {
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
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicFilterer) ParseMessageFailed(log types.Log) (*ContractAvsGovernanceLogicMessageFailed, error) {
	event := new(ContractAvsGovernanceLogicMessageFailed)
	if err := _ContractAvsGovernanceLogic.contract.UnpackLog(event, "MessageFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceLogicMessageSentIterator is returned from FilterMessageSent and is used to iterate over the raw logs and unpacked data for MessageSent events raised by the ContractAvsGovernanceLogic contract.
type ContractAvsGovernanceLogicMessageSentIterator struct {
	Event *ContractAvsGovernanceLogicMessageSent // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceLogicMessageSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceLogicMessageSent)
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
		it.Event = new(ContractAvsGovernanceLogicMessageSent)
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
func (it *ContractAvsGovernanceLogicMessageSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceLogicMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceLogicMessageSent represents a MessageSent event raised by the ContractAvsGovernanceLogic contract.
type ContractAvsGovernanceLogicMessageSent struct {
	DstEid uint32
	Guid   [32]byte
	Fee    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMessageSent is a free log retrieval operation binding the contract event 0x0597c8258e620343bfa7f3cf82a7d6578378f08248fd958f32bcab61a1004ec3.
//
// Solidity: event MessageSent(uint32 indexed dstEid, bytes32 indexed guid, uint256 fee)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicFilterer) FilterMessageSent(opts *bind.FilterOpts, dstEid []uint32, guid [][32]byte) (*ContractAvsGovernanceLogicMessageSentIterator, error) {

	var dstEidRule []interface{}
	for _, dstEidItem := range dstEid {
		dstEidRule = append(dstEidRule, dstEidItem)
	}
	var guidRule []interface{}
	for _, guidItem := range guid {
		guidRule = append(guidRule, guidItem)
	}

	logs, sub, err := _ContractAvsGovernanceLogic.contract.FilterLogs(opts, "MessageSent", dstEidRule, guidRule)
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceLogicMessageSentIterator{contract: _ContractAvsGovernanceLogic.contract, event: "MessageSent", logs: logs, sub: sub}, nil
}

// WatchMessageSent is a free log subscription operation binding the contract event 0x0597c8258e620343bfa7f3cf82a7d6578378f08248fd958f32bcab61a1004ec3.
//
// Solidity: event MessageSent(uint32 indexed dstEid, bytes32 indexed guid, uint256 fee)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicFilterer) WatchMessageSent(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceLogicMessageSent, dstEid []uint32, guid [][32]byte) (event.Subscription, error) {

	var dstEidRule []interface{}
	for _, dstEidItem := range dstEid {
		dstEidRule = append(dstEidRule, dstEidItem)
	}
	var guidRule []interface{}
	for _, guidItem := range guid {
		guidRule = append(guidRule, guidItem)
	}

	logs, sub, err := _ContractAvsGovernanceLogic.contract.WatchLogs(opts, "MessageSent", dstEidRule, guidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceLogicMessageSent)
				if err := _ContractAvsGovernanceLogic.contract.UnpackLog(event, "MessageSent", log); err != nil {
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
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicFilterer) ParseMessageSent(log types.Log) (*ContractAvsGovernanceLogicMessageSent, error) {
	event := new(ContractAvsGovernanceLogicMessageSent)
	if err := _ContractAvsGovernanceLogic.contract.UnpackLog(event, "MessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceLogicOperatorRegisteredIterator is returned from FilterOperatorRegistered and is used to iterate over the raw logs and unpacked data for OperatorRegistered events raised by the ContractAvsGovernanceLogic contract.
type ContractAvsGovernanceLogicOperatorRegisteredIterator struct {
	Event *ContractAvsGovernanceLogicOperatorRegistered // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceLogicOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceLogicOperatorRegistered)
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
		it.Event = new(ContractAvsGovernanceLogicOperatorRegistered)
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
func (it *ContractAvsGovernanceLogicOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceLogicOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceLogicOperatorRegistered represents a OperatorRegistered event raised by the ContractAvsGovernanceLogic contract.
type ContractAvsGovernanceLogicOperatorRegistered struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0x4d0eb1f4bac8744fd2be119845e23b3befc88094b42bcda1204c65694a00f9e5.
//
// Solidity: event OperatorRegistered(address indexed operator)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicFilterer) FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address) (*ContractAvsGovernanceLogicOperatorRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAvsGovernanceLogic.contract.FilterLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceLogicOperatorRegisteredIterator{contract: _ContractAvsGovernanceLogic.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0x4d0eb1f4bac8744fd2be119845e23b3befc88094b42bcda1204c65694a00f9e5.
//
// Solidity: event OperatorRegistered(address indexed operator)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicFilterer) WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceLogicOperatorRegistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAvsGovernanceLogic.contract.WatchLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceLogicOperatorRegistered)
				if err := _ContractAvsGovernanceLogic.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
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
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicFilterer) ParseOperatorRegistered(log types.Log) (*ContractAvsGovernanceLogicOperatorRegistered, error) {
	event := new(ContractAvsGovernanceLogicOperatorRegistered)
	if err := _ContractAvsGovernanceLogic.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceLogicOperatorUnregisteredIterator is returned from FilterOperatorUnregistered and is used to iterate over the raw logs and unpacked data for OperatorUnregistered events raised by the ContractAvsGovernanceLogic contract.
type ContractAvsGovernanceLogicOperatorUnregisteredIterator struct {
	Event *ContractAvsGovernanceLogicOperatorUnregistered // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceLogicOperatorUnregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceLogicOperatorUnregistered)
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
		it.Event = new(ContractAvsGovernanceLogicOperatorUnregistered)
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
func (it *ContractAvsGovernanceLogicOperatorUnregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceLogicOperatorUnregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceLogicOperatorUnregistered represents a OperatorUnregistered event raised by the ContractAvsGovernanceLogic contract.
type ContractAvsGovernanceLogicOperatorUnregistered struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorUnregistered is a free log retrieval operation binding the contract event 0x6f42117a557500c705ddf040a619d86f39101e6b74ac20d7b3e5943ba473fc7f.
//
// Solidity: event OperatorUnregistered(address indexed operator)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicFilterer) FilterOperatorUnregistered(opts *bind.FilterOpts, operator []common.Address) (*ContractAvsGovernanceLogicOperatorUnregisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAvsGovernanceLogic.contract.FilterLogs(opts, "OperatorUnregistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceLogicOperatorUnregisteredIterator{contract: _ContractAvsGovernanceLogic.contract, event: "OperatorUnregistered", logs: logs, sub: sub}, nil
}

// WatchOperatorUnregistered is a free log subscription operation binding the contract event 0x6f42117a557500c705ddf040a619d86f39101e6b74ac20d7b3e5943ba473fc7f.
//
// Solidity: event OperatorUnregistered(address indexed operator)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicFilterer) WatchOperatorUnregistered(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceLogicOperatorUnregistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAvsGovernanceLogic.contract.WatchLogs(opts, "OperatorUnregistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceLogicOperatorUnregistered)
				if err := _ContractAvsGovernanceLogic.contract.UnpackLog(event, "OperatorUnregistered", log); err != nil {
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
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicFilterer) ParseOperatorUnregistered(log types.Log) (*ContractAvsGovernanceLogicOperatorUnregistered, error) {
	event := new(ContractAvsGovernanceLogicOperatorUnregistered)
	if err := _ContractAvsGovernanceLogic.contract.UnpackLog(event, "OperatorUnregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceLogicOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractAvsGovernanceLogic contract.
type ContractAvsGovernanceLogicOwnershipTransferredIterator struct {
	Event *ContractAvsGovernanceLogicOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceLogicOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceLogicOwnershipTransferred)
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
		it.Event = new(ContractAvsGovernanceLogicOwnershipTransferred)
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
func (it *ContractAvsGovernanceLogicOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceLogicOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceLogicOwnershipTransferred represents a OwnershipTransferred event raised by the ContractAvsGovernanceLogic contract.
type ContractAvsGovernanceLogicOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractAvsGovernanceLogicOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractAvsGovernanceLogic.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceLogicOwnershipTransferredIterator{contract: _ContractAvsGovernanceLogic.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceLogicOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractAvsGovernanceLogic.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceLogicOwnershipTransferred)
				if err := _ContractAvsGovernanceLogic.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicFilterer) ParseOwnershipTransferred(log types.Log) (*ContractAvsGovernanceLogicOwnershipTransferred, error) {
	event := new(ContractAvsGovernanceLogicOwnershipTransferred)
	if err := _ContractAvsGovernanceLogic.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceLogicPeerSetIterator is returned from FilterPeerSet and is used to iterate over the raw logs and unpacked data for PeerSet events raised by the ContractAvsGovernanceLogic contract.
type ContractAvsGovernanceLogicPeerSetIterator struct {
	Event *ContractAvsGovernanceLogicPeerSet // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceLogicPeerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceLogicPeerSet)
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
		it.Event = new(ContractAvsGovernanceLogicPeerSet)
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
func (it *ContractAvsGovernanceLogicPeerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceLogicPeerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceLogicPeerSet represents a PeerSet event raised by the ContractAvsGovernanceLogic contract.
type ContractAvsGovernanceLogicPeerSet struct {
	Eid  uint32
	Peer [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPeerSet is a free log retrieval operation binding the contract event 0x238399d427b947898edb290f5ff0f9109849b1c3ba196a42e35f00c50a54b98b.
//
// Solidity: event PeerSet(uint32 eid, bytes32 peer)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicFilterer) FilterPeerSet(opts *bind.FilterOpts) (*ContractAvsGovernanceLogicPeerSetIterator, error) {

	logs, sub, err := _ContractAvsGovernanceLogic.contract.FilterLogs(opts, "PeerSet")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceLogicPeerSetIterator{contract: _ContractAvsGovernanceLogic.contract, event: "PeerSet", logs: logs, sub: sub}, nil
}

// WatchPeerSet is a free log subscription operation binding the contract event 0x238399d427b947898edb290f5ff0f9109849b1c3ba196a42e35f00c50a54b98b.
//
// Solidity: event PeerSet(uint32 eid, bytes32 peer)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicFilterer) WatchPeerSet(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceLogicPeerSet) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernanceLogic.contract.WatchLogs(opts, "PeerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceLogicPeerSet)
				if err := _ContractAvsGovernanceLogic.contract.UnpackLog(event, "PeerSet", log); err != nil {
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
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicFilterer) ParsePeerSet(log types.Log) (*ContractAvsGovernanceLogicPeerSet, error) {
	event := new(ContractAvsGovernanceLogicPeerSet)
	if err := _ContractAvsGovernanceLogic.contract.UnpackLog(event, "PeerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceLogicUnWhitelistedIterator is returned from FilterUnWhitelisted and is used to iterate over the raw logs and unpacked data for UnWhitelisted events raised by the ContractAvsGovernanceLogic contract.
type ContractAvsGovernanceLogicUnWhitelistedIterator struct {
	Event *ContractAvsGovernanceLogicUnWhitelisted // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceLogicUnWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceLogicUnWhitelisted)
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
		it.Event = new(ContractAvsGovernanceLogicUnWhitelisted)
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
func (it *ContractAvsGovernanceLogicUnWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceLogicUnWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceLogicUnWhitelisted represents a UnWhitelisted event raised by the ContractAvsGovernanceLogic contract.
type ContractAvsGovernanceLogicUnWhitelisted struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUnWhitelisted is a free log retrieval operation binding the contract event 0x7cdb51b0cc2e541ad7e9471c358de415f5dbaff6cca78e3393d445346c610c1b.
//
// Solidity: event UnWhitelisted(address indexed operator)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicFilterer) FilterUnWhitelisted(opts *bind.FilterOpts, operator []common.Address) (*ContractAvsGovernanceLogicUnWhitelistedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAvsGovernanceLogic.contract.FilterLogs(opts, "UnWhitelisted", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceLogicUnWhitelistedIterator{contract: _ContractAvsGovernanceLogic.contract, event: "UnWhitelisted", logs: logs, sub: sub}, nil
}

// WatchUnWhitelisted is a free log subscription operation binding the contract event 0x7cdb51b0cc2e541ad7e9471c358de415f5dbaff6cca78e3393d445346c610c1b.
//
// Solidity: event UnWhitelisted(address indexed operator)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicFilterer) WatchUnWhitelisted(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceLogicUnWhitelisted, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAvsGovernanceLogic.contract.WatchLogs(opts, "UnWhitelisted", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceLogicUnWhitelisted)
				if err := _ContractAvsGovernanceLogic.contract.UnpackLog(event, "UnWhitelisted", log); err != nil {
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
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicFilterer) ParseUnWhitelisted(log types.Log) (*ContractAvsGovernanceLogicUnWhitelisted, error) {
	event := new(ContractAvsGovernanceLogicUnWhitelisted)
	if err := _ContractAvsGovernanceLogic.contract.UnpackLog(event, "UnWhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceLogicWhitelistManagerAddedIterator is returned from FilterWhitelistManagerAdded and is used to iterate over the raw logs and unpacked data for WhitelistManagerAdded events raised by the ContractAvsGovernanceLogic contract.
type ContractAvsGovernanceLogicWhitelistManagerAddedIterator struct {
	Event *ContractAvsGovernanceLogicWhitelistManagerAdded // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceLogicWhitelistManagerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceLogicWhitelistManagerAdded)
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
		it.Event = new(ContractAvsGovernanceLogicWhitelistManagerAdded)
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
func (it *ContractAvsGovernanceLogicWhitelistManagerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceLogicWhitelistManagerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceLogicWhitelistManagerAdded represents a WhitelistManagerAdded event raised by the ContractAvsGovernanceLogic contract.
type ContractAvsGovernanceLogicWhitelistManagerAdded struct {
	Manager common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistManagerAdded is a free log retrieval operation binding the contract event 0x15a94be709049b6d6ac9d90d9a1ee40b18b9071a625486fad8ab1d45cedcc24d.
//
// Solidity: event WhitelistManagerAdded(address indexed manager)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicFilterer) FilterWhitelistManagerAdded(opts *bind.FilterOpts, manager []common.Address) (*ContractAvsGovernanceLogicWhitelistManagerAddedIterator, error) {

	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}

	logs, sub, err := _ContractAvsGovernanceLogic.contract.FilterLogs(opts, "WhitelistManagerAdded", managerRule)
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceLogicWhitelistManagerAddedIterator{contract: _ContractAvsGovernanceLogic.contract, event: "WhitelistManagerAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistManagerAdded is a free log subscription operation binding the contract event 0x15a94be709049b6d6ac9d90d9a1ee40b18b9071a625486fad8ab1d45cedcc24d.
//
// Solidity: event WhitelistManagerAdded(address indexed manager)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicFilterer) WatchWhitelistManagerAdded(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceLogicWhitelistManagerAdded, manager []common.Address) (event.Subscription, error) {

	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}

	logs, sub, err := _ContractAvsGovernanceLogic.contract.WatchLogs(opts, "WhitelistManagerAdded", managerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceLogicWhitelistManagerAdded)
				if err := _ContractAvsGovernanceLogic.contract.UnpackLog(event, "WhitelistManagerAdded", log); err != nil {
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

// ParseWhitelistManagerAdded is a log parse operation binding the contract event 0x15a94be709049b6d6ac9d90d9a1ee40b18b9071a625486fad8ab1d45cedcc24d.
//
// Solidity: event WhitelistManagerAdded(address indexed manager)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicFilterer) ParseWhitelistManagerAdded(log types.Log) (*ContractAvsGovernanceLogicWhitelistManagerAdded, error) {
	event := new(ContractAvsGovernanceLogicWhitelistManagerAdded)
	if err := _ContractAvsGovernanceLogic.contract.UnpackLog(event, "WhitelistManagerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceLogicWhitelistManagerRemovedIterator is returned from FilterWhitelistManagerRemoved and is used to iterate over the raw logs and unpacked data for WhitelistManagerRemoved events raised by the ContractAvsGovernanceLogic contract.
type ContractAvsGovernanceLogicWhitelistManagerRemovedIterator struct {
	Event *ContractAvsGovernanceLogicWhitelistManagerRemoved // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceLogicWhitelistManagerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceLogicWhitelistManagerRemoved)
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
		it.Event = new(ContractAvsGovernanceLogicWhitelistManagerRemoved)
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
func (it *ContractAvsGovernanceLogicWhitelistManagerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceLogicWhitelistManagerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceLogicWhitelistManagerRemoved represents a WhitelistManagerRemoved event raised by the ContractAvsGovernanceLogic contract.
type ContractAvsGovernanceLogicWhitelistManagerRemoved struct {
	Manager common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistManagerRemoved is a free log retrieval operation binding the contract event 0xce4df0698b95faea0f80f49bc274673a82e9cb758eb9df7c2951b3ab6732350f.
//
// Solidity: event WhitelistManagerRemoved(address indexed manager)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicFilterer) FilterWhitelistManagerRemoved(opts *bind.FilterOpts, manager []common.Address) (*ContractAvsGovernanceLogicWhitelistManagerRemovedIterator, error) {

	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}

	logs, sub, err := _ContractAvsGovernanceLogic.contract.FilterLogs(opts, "WhitelistManagerRemoved", managerRule)
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceLogicWhitelistManagerRemovedIterator{contract: _ContractAvsGovernanceLogic.contract, event: "WhitelistManagerRemoved", logs: logs, sub: sub}, nil
}

// WatchWhitelistManagerRemoved is a free log subscription operation binding the contract event 0xce4df0698b95faea0f80f49bc274673a82e9cb758eb9df7c2951b3ab6732350f.
//
// Solidity: event WhitelistManagerRemoved(address indexed manager)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicFilterer) WatchWhitelistManagerRemoved(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceLogicWhitelistManagerRemoved, manager []common.Address) (event.Subscription, error) {

	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}

	logs, sub, err := _ContractAvsGovernanceLogic.contract.WatchLogs(opts, "WhitelistManagerRemoved", managerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceLogicWhitelistManagerRemoved)
				if err := _ContractAvsGovernanceLogic.contract.UnpackLog(event, "WhitelistManagerRemoved", log); err != nil {
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

// ParseWhitelistManagerRemoved is a log parse operation binding the contract event 0xce4df0698b95faea0f80f49bc274673a82e9cb758eb9df7c2951b3ab6732350f.
//
// Solidity: event WhitelistManagerRemoved(address indexed manager)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicFilterer) ParseWhitelistManagerRemoved(log types.Log) (*ContractAvsGovernanceLogicWhitelistManagerRemoved, error) {
	event := new(ContractAvsGovernanceLogicWhitelistManagerRemoved)
	if err := _ContractAvsGovernanceLogic.contract.UnpackLog(event, "WhitelistManagerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceLogicWhitelistedIterator is returned from FilterWhitelisted and is used to iterate over the raw logs and unpacked data for Whitelisted events raised by the ContractAvsGovernanceLogic contract.
type ContractAvsGovernanceLogicWhitelistedIterator struct {
	Event *ContractAvsGovernanceLogicWhitelisted // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceLogicWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceLogicWhitelisted)
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
		it.Event = new(ContractAvsGovernanceLogicWhitelisted)
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
func (it *ContractAvsGovernanceLogicWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceLogicWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceLogicWhitelisted represents a Whitelisted event raised by the ContractAvsGovernanceLogic contract.
type ContractAvsGovernanceLogicWhitelisted struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWhitelisted is a free log retrieval operation binding the contract event 0xaab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a54.
//
// Solidity: event Whitelisted(address indexed operator)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicFilterer) FilterWhitelisted(opts *bind.FilterOpts, operator []common.Address) (*ContractAvsGovernanceLogicWhitelistedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAvsGovernanceLogic.contract.FilterLogs(opts, "Whitelisted", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceLogicWhitelistedIterator{contract: _ContractAvsGovernanceLogic.contract, event: "Whitelisted", logs: logs, sub: sub}, nil
}

// WatchWhitelisted is a free log subscription operation binding the contract event 0xaab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a54.
//
// Solidity: event Whitelisted(address indexed operator)
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicFilterer) WatchWhitelisted(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceLogicWhitelisted, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAvsGovernanceLogic.contract.WatchLogs(opts, "Whitelisted", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceLogicWhitelisted)
				if err := _ContractAvsGovernanceLogic.contract.UnpackLog(event, "Whitelisted", log); err != nil {
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
func (_ContractAvsGovernanceLogic *ContractAvsGovernanceLogicFilterer) ParseWhitelisted(log types.Log) (*ContractAvsGovernanceLogicWhitelisted, error) {
	event := new(ContractAvsGovernanceLogicWhitelisted)
	if err := _ContractAvsGovernanceLogic.contract.UnpackLog(event, "Whitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
