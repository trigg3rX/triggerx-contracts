// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractTriggerXStakeRegistry

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

// ContractTriggerXStakeRegistryMetaData contains all meta data concerning the ContractTriggerXStakeRegistry contract.
var ContractTriggerXStakeRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"StakeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TGClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TGTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TaskFeeClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Unstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"TGAmount\",\"type\":\"uint256\"}],\"name\":\"claimETHForTG\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"TGbalance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"points\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"TGbalance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"TGAmount\",\"type\":\"uint256\"}],\"name\":\"updateTGBalances\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x6080604052600436101561001257600080fd5b60003560e01c8063168d6c5014610d4a57806316934fc414610d07578063358b816614610ccd5780633659cfe6146109f25780634f1ef286146106bb57806352d1902d146105f7578063715018a61461059a5780637a766460146105175780637da49f3a146104025780638129fc1c146102a65780638da5cb5b1461027d578063a694fc3a1461013f5763f2fde38b146100ab57600080fd5b3461013a57602036600319011261013a576100c4610e65565b6100cc611011565b6001600160a01b038116156100e6576100e490611069565b005b60405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608490fd5b600080fd5b602036600319011261013a5760043561015d60026001541415610eb9565b60026001558015610243578034036101fe573360005260fb6020526040600020610188828254611004565b81558180046001036101e8576103e88202908282046103e8036101e85760016101b49101918254611004565b90556040519081527f9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d60203392a260018055005b634e487b7160e01b600052601160045260246000fd5b60405162461bcd60e51b815260206004820152601a60248201527f53656e7420455448206d757374206d6174636820616d6f756e740000000000006044820152606490fd5b60405162461bcd60e51b8152602060048201526012602482015271086c2dcdcdee840e6e8c2d6ca4060408aa8960731b6044820152606490fd5b3461013a57600036600319011261013a5760c9546040516001600160a01b039091168152602090f35b3461013a57600036600319011261013a5760005460ff8160081c1615908180926103f5575b80156103de575b156103825760ff19811660011760005581610370575b5061031960ff60005460081c166102fe816110b2565b610307816110b2565b60018055610314816110b2565b6110b2565b61032233611069565b6000549061033560ff8360081c166110b2565b61033b57005b61ff0019166000557f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498602060405160018152a1005b61ffff191661010117600055816102e8565b60405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608490fd5b50303b1580156102d25750600160ff8216146102d2565b50600160ff8216106102cb565b3461013a57606036600319011261013a5761041b610e65565b6024356001600160a01b0381169081900361013a576044359161043c611011565b61044b60026001541415610eb9565b60026001558160005260fb60205260406000209060018060a01b0316928360005260fb6020526001604060002092019081548181106104d2577fa613b7e352cf78a30d92dbc93071265034f43edbeb7833b8d8713a3b4f323670936020936104b584600194610f05565b9055016104c3828254611004565b9055604051908152a360018055005b60405162461bcd60e51b815260206004820152601c60248201527f496e73756666696369656e7420757365722054472062616c616e6365000000006044820152606490fd5b3461013a57602036600319011261013a576001600160a01b03610538610e65565b1660005260fb6020526040600020604051906040820182811067ffffffffffffffff82111761058457604092602091845260018354938483520154918291015282519182526020820152f35b634e487b7160e01b600052604160045260246000fd5b3461013a57600036600319011261013a576105b3611011565b60c980546001600160a01b031981169091556000906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b3461013a57600036600319011261013a577f0000000000000000000000006d0be42ea76f9c6700a985b35c91cc8382b79a726001600160a01b031630036106505760206040516000805160206112288339815191528152f35b60405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c00000000000000006064820152608490fd5b604036600319011261013a576106cf610e65565b60243567ffffffffffffffff811161013a573660238201121561013a5780600401356106fa81610e9d565b906107086040519283610e7b565b8082526020820192366024838301011161013a57816000926024602093018637830101526107847f0000000000000000000000006d0be42ea76f9c6700a985b35c91cc8382b79a726001600160a01b031661076530821415610f42565b600080516020611228833981519152546001600160a01b031614610fa3565b61078c611011565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff16156107c15750506100e490611112565b6040516352d1902d60e01b81526001600160a01b03841690602081600481855afa600091816109be575b5061084c5760405162461bcd60e51b815260206004820152602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b6064820152608490fd5b600080516020611228833981519152036109675761086984611112565b604051907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b600080a281511580159061095f575b6108a357005b833b1561090e57506100e492600092839251915af46108c0610f12565b604051916108cf606084610e7b565b602783527f416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c6020840152660819985a5b195960ca1b60408401526111b0565b62461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b6064820152608490fd5b50600161089d565b60405162461bcd60e51b815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b6064820152608490fd5b9091506020813d6020116109ea575b816109da60209383610e7b565b8101031261013a575190866107eb565b3d91506109cd565b3461013a57602036600319011261013a57610a0b610e65565b610a447f0000000000000000000000006d0be42ea76f9c6700a985b35c91cc8382b79a726001600160a01b031661076530821415610f42565b610a4c611011565b602090604051610a5c8382610e7b565b6000815282810190601f1984013683377f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff1615610aa25750506100e49150611112565b6040516352d1902d60e01b81526001600160a01b038416908581600481855afa60009181610c9e575b50610b2c5760405162461bcd60e51b815260048101879052602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b6064820152608490fd5b60008051602061122883398151915203610c4757610b4984611112565b604051907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b600080a2815115801590610c3f575b610b8357005b833b15610bee57506100e49392600092839251915af4610ba1610f12565b907f416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c60405193610bd2606086610e7b565b60278552840152660819985a5b195960ca1b60408401526111b0565b62461bcd60e51b815260048101859052602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b6064820152608490fd5b506000610b7d565b60405162461bcd60e51b815260048101869052602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b6064820152608490fd5b9091508681813d8311610cc6575b610cb68183610e7b565b8101031261013a57519087610acb565b503d610cac565b3461013a57602036600319011261013a576001600160a01b03610cee610e65565b1660005260fc6020526020604060002054604051908152f35b3461013a57602036600319011261013a576001600160a01b03610d28610e65565b1660005260fb6020526040806000206001815491015482519182526020820152f35b3461013a57602036600319011261013a57600435610d6d60026001541415610eb9565b60026001553360005260fb60205260016040600020018054828110610e205782610d9691610f05565b905560008080806103e88504335af1610dad610f12565b5015610de5576040519081527fbddee7230372ecc0cb1780b66efe6d1461ee2a8e5b70239f3ec20a7d85f0e92860203392a260018055005b60405162461bcd60e51b8152602060048201526013602482015272115512081d1c985b9cd9995c8819985a5b1959606a1b6044820152606490fd5b60405162461bcd60e51b815260206004820152601760248201527f496e73756666696369656e742054472062616c616e63650000000000000000006044820152606490fd5b600435906001600160a01b038216820361013a57565b90601f8019910116810190811067ffffffffffffffff82111761058457604052565b67ffffffffffffffff811161058457601f01601f191660200190565b15610ec057565b60405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606490fd5b919082039182116101e857565b3d15610f3d573d90610f2382610e9d565b91610f316040519384610e7b565b82523d6000602084013e565b606090565b15610f4957565b60405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b19195b1959d85d1958d85b1b60a21b6064820152608490fd5b15610faa57565b60405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b6163746976652070726f787960a01b6064820152608490fd5b919082018092116101e857565b60c9546001600160a01b0316330361102557565b606460405162461bcd60e51b815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b60c980546001600160a01b039283166001600160a01b0319821681179092559091167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3565b156110b957565b60405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b6064820152608490fd5b803b156111555760018060a01b03166bffffffffffffffffffffffff60a01b60008051602061122883398151915254161760008051602061122883398151915255565b60405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608490fd5b909190156111bc575090565b8151156111cc5750805190602001fd5b6040519062461bcd60e51b8252602060048301528181519182602483015260005b83811061120f5750508160006044809484010152601f80199101168101030190fd5b602082820181015160448784010152859350016111ed56fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbca26469706673582212201a8f2e27c3e82b8a8f3d1dd44611e03d02c370776af2afb12dd5dfe351b7070a64736f6c634300081a0033",
}

// ContractTriggerXStakeRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractTriggerXStakeRegistryMetaData.ABI instead.
var ContractTriggerXStakeRegistryABI = ContractTriggerXStakeRegistryMetaData.ABI

// ContractTriggerXStakeRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractTriggerXStakeRegistryMetaData.Bin instead.
var ContractTriggerXStakeRegistryBin = ContractTriggerXStakeRegistryMetaData.Bin

// DeployContractTriggerXStakeRegistry deploys a new Ethereum contract, binding an instance of ContractTriggerXStakeRegistry to it.
func DeployContractTriggerXStakeRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContractTriggerXStakeRegistry, error) {
	parsed, err := ContractTriggerXStakeRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractTriggerXStakeRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractTriggerXStakeRegistry{ContractTriggerXStakeRegistryCaller: ContractTriggerXStakeRegistryCaller{contract: contract}, ContractTriggerXStakeRegistryTransactor: ContractTriggerXStakeRegistryTransactor{contract: contract}, ContractTriggerXStakeRegistryFilterer: ContractTriggerXStakeRegistryFilterer{contract: contract}}, nil
}

// ContractTriggerXStakeRegistryMethods is an auto generated interface around an Ethereum contract.
type ContractTriggerXStakeRegistryMethods interface {
	ContractTriggerXStakeRegistryCalls
	ContractTriggerXStakeRegistryTransacts
	ContractTriggerXStakeRegistryFilters
}

// ContractTriggerXStakeRegistryCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractTriggerXStakeRegistryCalls interface {
	GetStake(opts *bind.CallOpts, user common.Address) (struct {
		Amount    *big.Int
		TGbalance *big.Int
	}, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Points(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error)

	ProxiableUUID(opts *bind.CallOpts) ([32]byte, error)

	Stakes(opts *bind.CallOpts, arg0 common.Address) (struct {
		Amount    *big.Int
		TGbalance *big.Int
	}, error)
}

// ContractTriggerXStakeRegistryTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractTriggerXStakeRegistryTransacts interface {
	ClaimETHForTG(opts *bind.TransactOpts, TGAmount *big.Int) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts) (*types.Transaction, error)

	RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	Stake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error)

	UpdateTGBalances(opts *bind.TransactOpts, keeper common.Address, user common.Address, TGAmount *big.Int) (*types.Transaction, error)

	UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error)

	UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error)
}

// ContractTriggerXStakeRegistryFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractTriggerXStakeRegistryFilters interface {
	FilterAdminChanged(opts *bind.FilterOpts) (*ContractTriggerXStakeRegistryAdminChangedIterator, error)
	WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *ContractTriggerXStakeRegistryAdminChanged) (event.Subscription, error)
	ParseAdminChanged(log types.Log) (*ContractTriggerXStakeRegistryAdminChanged, error)

	FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*ContractTriggerXStakeRegistryBeaconUpgradedIterator, error)
	WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *ContractTriggerXStakeRegistryBeaconUpgraded, beacon []common.Address) (event.Subscription, error)
	ParseBeaconUpgraded(log types.Log) (*ContractTriggerXStakeRegistryBeaconUpgraded, error)

	FilterInitialized(opts *bind.FilterOpts) (*ContractTriggerXStakeRegistryInitializedIterator, error)
	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractTriggerXStakeRegistryInitialized) (event.Subscription, error)
	ParseInitialized(log types.Log) (*ContractTriggerXStakeRegistryInitialized, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractTriggerXStakeRegistryOwnershipTransferredIterator, error)
	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractTriggerXStakeRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error)
	ParseOwnershipTransferred(log types.Log) (*ContractTriggerXStakeRegistryOwnershipTransferred, error)

	FilterRewardClaimed(opts *bind.FilterOpts, user []common.Address) (*ContractTriggerXStakeRegistryRewardClaimedIterator, error)
	WatchRewardClaimed(opts *bind.WatchOpts, sink chan<- *ContractTriggerXStakeRegistryRewardClaimed, user []common.Address) (event.Subscription, error)
	ParseRewardClaimed(log types.Log) (*ContractTriggerXStakeRegistryRewardClaimed, error)

	FilterStakeRemoved(opts *bind.FilterOpts, user []common.Address) (*ContractTriggerXStakeRegistryStakeRemovedIterator, error)
	WatchStakeRemoved(opts *bind.WatchOpts, sink chan<- *ContractTriggerXStakeRegistryStakeRemoved, user []common.Address) (event.Subscription, error)
	ParseStakeRemoved(log types.Log) (*ContractTriggerXStakeRegistryStakeRemoved, error)

	FilterStaked(opts *bind.FilterOpts, user []common.Address) (*ContractTriggerXStakeRegistryStakedIterator, error)
	WatchStaked(opts *bind.WatchOpts, sink chan<- *ContractTriggerXStakeRegistryStaked, user []common.Address) (event.Subscription, error)
	ParseStaked(log types.Log) (*ContractTriggerXStakeRegistryStaked, error)

	FilterTGClaimed(opts *bind.FilterOpts, user []common.Address) (*ContractTriggerXStakeRegistryTGClaimedIterator, error)
	WatchTGClaimed(opts *bind.WatchOpts, sink chan<- *ContractTriggerXStakeRegistryTGClaimed, user []common.Address) (event.Subscription, error)
	ParseTGClaimed(log types.Log) (*ContractTriggerXStakeRegistryTGClaimed, error)

	FilterTGTransferred(opts *bind.FilterOpts, user []common.Address, keeper []common.Address) (*ContractTriggerXStakeRegistryTGTransferredIterator, error)
	WatchTGTransferred(opts *bind.WatchOpts, sink chan<- *ContractTriggerXStakeRegistryTGTransferred, user []common.Address, keeper []common.Address) (event.Subscription, error)
	ParseTGTransferred(log types.Log) (*ContractTriggerXStakeRegistryTGTransferred, error)

	FilterTaskFeeClaimed(opts *bind.FilterOpts, user []common.Address) (*ContractTriggerXStakeRegistryTaskFeeClaimedIterator, error)
	WatchTaskFeeClaimed(opts *bind.WatchOpts, sink chan<- *ContractTriggerXStakeRegistryTaskFeeClaimed, user []common.Address) (event.Subscription, error)
	ParseTaskFeeClaimed(log types.Log) (*ContractTriggerXStakeRegistryTaskFeeClaimed, error)

	FilterUnstaked(opts *bind.FilterOpts, user []common.Address) (*ContractTriggerXStakeRegistryUnstakedIterator, error)
	WatchUnstaked(opts *bind.WatchOpts, sink chan<- *ContractTriggerXStakeRegistryUnstaked, user []common.Address) (event.Subscription, error)
	ParseUnstaked(log types.Log) (*ContractTriggerXStakeRegistryUnstaked, error)

	FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ContractTriggerXStakeRegistryUpgradedIterator, error)
	WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ContractTriggerXStakeRegistryUpgraded, implementation []common.Address) (event.Subscription, error)
	ParseUpgraded(log types.Log) (*ContractTriggerXStakeRegistryUpgraded, error)
}

// ContractTriggerXStakeRegistry is an auto generated Go binding around an Ethereum contract.
type ContractTriggerXStakeRegistry struct {
	ContractTriggerXStakeRegistryCaller     // Read-only binding to the contract
	ContractTriggerXStakeRegistryTransactor // Write-only binding to the contract
	ContractTriggerXStakeRegistryFilterer   // Log filterer for contract events
}

// ContractTriggerXStakeRegistry implements the ContractTriggerXStakeRegistryMethods interface.
var _ ContractTriggerXStakeRegistryMethods = (*ContractTriggerXStakeRegistry)(nil)

// ContractTriggerXStakeRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractTriggerXStakeRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTriggerXStakeRegistryCaller implements the ContractTriggerXStakeRegistryCalls interface.
var _ ContractTriggerXStakeRegistryCalls = (*ContractTriggerXStakeRegistryCaller)(nil)

// ContractTriggerXStakeRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTriggerXStakeRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTriggerXStakeRegistryTransactor implements the ContractTriggerXStakeRegistryTransacts interface.
var _ ContractTriggerXStakeRegistryTransacts = (*ContractTriggerXStakeRegistryTransactor)(nil)

// ContractTriggerXStakeRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractTriggerXStakeRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTriggerXStakeRegistryFilterer implements the ContractTriggerXStakeRegistryFilters interface.
var _ ContractTriggerXStakeRegistryFilters = (*ContractTriggerXStakeRegistryFilterer)(nil)

// ContractTriggerXStakeRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractTriggerXStakeRegistrySession struct {
	Contract     *ContractTriggerXStakeRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ContractTriggerXStakeRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractTriggerXStakeRegistryCallerSession struct {
	Contract *ContractTriggerXStakeRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// ContractTriggerXStakeRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTriggerXStakeRegistryTransactorSession struct {
	Contract     *ContractTriggerXStakeRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// ContractTriggerXStakeRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractTriggerXStakeRegistryRaw struct {
	Contract *ContractTriggerXStakeRegistry // Generic contract binding to access the raw methods on
}

// ContractTriggerXStakeRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractTriggerXStakeRegistryCallerRaw struct {
	Contract *ContractTriggerXStakeRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTriggerXStakeRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTriggerXStakeRegistryTransactorRaw struct {
	Contract *ContractTriggerXStakeRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractTriggerXStakeRegistry creates a new instance of ContractTriggerXStakeRegistry, bound to a specific deployed contract.
func NewContractTriggerXStakeRegistry(address common.Address, backend bind.ContractBackend) (*ContractTriggerXStakeRegistry, error) {
	contract, err := bindContractTriggerXStakeRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXStakeRegistry{ContractTriggerXStakeRegistryCaller: ContractTriggerXStakeRegistryCaller{contract: contract}, ContractTriggerXStakeRegistryTransactor: ContractTriggerXStakeRegistryTransactor{contract: contract}, ContractTriggerXStakeRegistryFilterer: ContractTriggerXStakeRegistryFilterer{contract: contract}}, nil
}

// NewContractTriggerXStakeRegistryCaller creates a new read-only instance of ContractTriggerXStakeRegistry, bound to a specific deployed contract.
func NewContractTriggerXStakeRegistryCaller(address common.Address, caller bind.ContractCaller) (*ContractTriggerXStakeRegistryCaller, error) {
	contract, err := bindContractTriggerXStakeRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXStakeRegistryCaller{contract: contract}, nil
}

// NewContractTriggerXStakeRegistryTransactor creates a new write-only instance of ContractTriggerXStakeRegistry, bound to a specific deployed contract.
func NewContractTriggerXStakeRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTriggerXStakeRegistryTransactor, error) {
	contract, err := bindContractTriggerXStakeRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXStakeRegistryTransactor{contract: contract}, nil
}

// NewContractTriggerXStakeRegistryFilterer creates a new log filterer instance of ContractTriggerXStakeRegistry, bound to a specific deployed contract.
func NewContractTriggerXStakeRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractTriggerXStakeRegistryFilterer, error) {
	contract, err := bindContractTriggerXStakeRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXStakeRegistryFilterer{contract: contract}, nil
}

// bindContractTriggerXStakeRegistry binds a generic wrapper to an already deployed contract.
func bindContractTriggerXStakeRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractTriggerXStakeRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractTriggerXStakeRegistry.Contract.ContractTriggerXStakeRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractTriggerXStakeRegistry.Contract.ContractTriggerXStakeRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractTriggerXStakeRegistry.Contract.ContractTriggerXStakeRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractTriggerXStakeRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractTriggerXStakeRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractTriggerXStakeRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetStake is a free data retrieval call binding the contract method 0x7a766460.
//
// Solidity: function getStake(address user) view returns(uint256 amount, uint256 TGbalance)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryCaller) GetStake(opts *bind.CallOpts, user common.Address) (struct {
	Amount    *big.Int
	TGbalance *big.Int
}, error) {
	var out []interface{}
	err := _ContractTriggerXStakeRegistry.contract.Call(opts, &out, "getStake", user)

	outstruct := new(struct {
		Amount    *big.Int
		TGbalance *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TGbalance = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetStake is a free data retrieval call binding the contract method 0x7a766460.
//
// Solidity: function getStake(address user) view returns(uint256 amount, uint256 TGbalance)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistrySession) GetStake(user common.Address) (struct {
	Amount    *big.Int
	TGbalance *big.Int
}, error) {
	return _ContractTriggerXStakeRegistry.Contract.GetStake(&_ContractTriggerXStakeRegistry.CallOpts, user)
}

// GetStake is a free data retrieval call binding the contract method 0x7a766460.
//
// Solidity: function getStake(address user) view returns(uint256 amount, uint256 TGbalance)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryCallerSession) GetStake(user common.Address) (struct {
	Amount    *big.Int
	TGbalance *big.Int
}, error) {
	return _ContractTriggerXStakeRegistry.Contract.GetStake(&_ContractTriggerXStakeRegistry.CallOpts, user)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTriggerXStakeRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistrySession) Owner() (common.Address, error) {
	return _ContractTriggerXStakeRegistry.Contract.Owner(&_ContractTriggerXStakeRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryCallerSession) Owner() (common.Address, error) {
	return _ContractTriggerXStakeRegistry.Contract.Owner(&_ContractTriggerXStakeRegistry.CallOpts)
}

// Points is a free data retrieval call binding the contract method 0x358b8166.
//
// Solidity: function points(address ) view returns(uint256)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryCaller) Points(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractTriggerXStakeRegistry.contract.Call(opts, &out, "points", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Points is a free data retrieval call binding the contract method 0x358b8166.
//
// Solidity: function points(address ) view returns(uint256)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistrySession) Points(arg0 common.Address) (*big.Int, error) {
	return _ContractTriggerXStakeRegistry.Contract.Points(&_ContractTriggerXStakeRegistry.CallOpts, arg0)
}

// Points is a free data retrieval call binding the contract method 0x358b8166.
//
// Solidity: function points(address ) view returns(uint256)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryCallerSession) Points(arg0 common.Address) (*big.Int, error) {
	return _ContractTriggerXStakeRegistry.Contract.Points(&_ContractTriggerXStakeRegistry.CallOpts, arg0)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractTriggerXStakeRegistry.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistrySession) ProxiableUUID() ([32]byte, error) {
	return _ContractTriggerXStakeRegistry.Contract.ProxiableUUID(&_ContractTriggerXStakeRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _ContractTriggerXStakeRegistry.Contract.ProxiableUUID(&_ContractTriggerXStakeRegistry.CallOpts)
}

// Stakes is a free data retrieval call binding the contract method 0x16934fc4.
//
// Solidity: function stakes(address ) view returns(uint256 amount, uint256 TGbalance)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryCaller) Stakes(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount    *big.Int
	TGbalance *big.Int
}, error) {
	var out []interface{}
	err := _ContractTriggerXStakeRegistry.contract.Call(opts, &out, "stakes", arg0)

	outstruct := new(struct {
		Amount    *big.Int
		TGbalance *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TGbalance = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Stakes is a free data retrieval call binding the contract method 0x16934fc4.
//
// Solidity: function stakes(address ) view returns(uint256 amount, uint256 TGbalance)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistrySession) Stakes(arg0 common.Address) (struct {
	Amount    *big.Int
	TGbalance *big.Int
}, error) {
	return _ContractTriggerXStakeRegistry.Contract.Stakes(&_ContractTriggerXStakeRegistry.CallOpts, arg0)
}

// Stakes is a free data retrieval call binding the contract method 0x16934fc4.
//
// Solidity: function stakes(address ) view returns(uint256 amount, uint256 TGbalance)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryCallerSession) Stakes(arg0 common.Address) (struct {
	Amount    *big.Int
	TGbalance *big.Int
}, error) {
	return _ContractTriggerXStakeRegistry.Contract.Stakes(&_ContractTriggerXStakeRegistry.CallOpts, arg0)
}

// ClaimETHForTG is a paid mutator transaction binding the contract method 0x168d6c50.
//
// Solidity: function claimETHForTG(uint256 TGAmount) returns()
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryTransactor) ClaimETHForTG(opts *bind.TransactOpts, TGAmount *big.Int) (*types.Transaction, error) {
	return _ContractTriggerXStakeRegistry.contract.Transact(opts, "claimETHForTG", TGAmount)
}

// ClaimETHForTG is a paid mutator transaction binding the contract method 0x168d6c50.
//
// Solidity: function claimETHForTG(uint256 TGAmount) returns()
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistrySession) ClaimETHForTG(TGAmount *big.Int) (*types.Transaction, error) {
	return _ContractTriggerXStakeRegistry.Contract.ClaimETHForTG(&_ContractTriggerXStakeRegistry.TransactOpts, TGAmount)
}

// ClaimETHForTG is a paid mutator transaction binding the contract method 0x168d6c50.
//
// Solidity: function claimETHForTG(uint256 TGAmount) returns()
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryTransactorSession) ClaimETHForTG(TGAmount *big.Int) (*types.Transaction, error) {
	return _ContractTriggerXStakeRegistry.Contract.ClaimETHForTG(&_ContractTriggerXStakeRegistry.TransactOpts, TGAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractTriggerXStakeRegistry.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistrySession) Initialize() (*types.Transaction, error) {
	return _ContractTriggerXStakeRegistry.Contract.Initialize(&_ContractTriggerXStakeRegistry.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryTransactorSession) Initialize() (*types.Transaction, error) {
	return _ContractTriggerXStakeRegistry.Contract.Initialize(&_ContractTriggerXStakeRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractTriggerXStakeRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractTriggerXStakeRegistry.Contract.RenounceOwnership(&_ContractTriggerXStakeRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractTriggerXStakeRegistry.Contract.RenounceOwnership(&_ContractTriggerXStakeRegistry.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) payable returns()
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryTransactor) Stake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ContractTriggerXStakeRegistry.contract.Transact(opts, "stake", amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) payable returns()
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistrySession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _ContractTriggerXStakeRegistry.Contract.Stake(&_ContractTriggerXStakeRegistry.TransactOpts, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) payable returns()
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryTransactorSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _ContractTriggerXStakeRegistry.Contract.Stake(&_ContractTriggerXStakeRegistry.TransactOpts, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractTriggerXStakeRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractTriggerXStakeRegistry.Contract.TransferOwnership(&_ContractTriggerXStakeRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractTriggerXStakeRegistry.Contract.TransferOwnership(&_ContractTriggerXStakeRegistry.TransactOpts, newOwner)
}

// UpdateTGBalances is a paid mutator transaction binding the contract method 0x7da49f3a.
//
// Solidity: function updateTGBalances(address keeper, address user, uint256 TGAmount) returns()
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryTransactor) UpdateTGBalances(opts *bind.TransactOpts, keeper common.Address, user common.Address, TGAmount *big.Int) (*types.Transaction, error) {
	return _ContractTriggerXStakeRegistry.contract.Transact(opts, "updateTGBalances", keeper, user, TGAmount)
}

// UpdateTGBalances is a paid mutator transaction binding the contract method 0x7da49f3a.
//
// Solidity: function updateTGBalances(address keeper, address user, uint256 TGAmount) returns()
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistrySession) UpdateTGBalances(keeper common.Address, user common.Address, TGAmount *big.Int) (*types.Transaction, error) {
	return _ContractTriggerXStakeRegistry.Contract.UpdateTGBalances(&_ContractTriggerXStakeRegistry.TransactOpts, keeper, user, TGAmount)
}

// UpdateTGBalances is a paid mutator transaction binding the contract method 0x7da49f3a.
//
// Solidity: function updateTGBalances(address keeper, address user, uint256 TGAmount) returns()
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryTransactorSession) UpdateTGBalances(keeper common.Address, user common.Address, TGAmount *big.Int) (*types.Transaction, error) {
	return _ContractTriggerXStakeRegistry.Contract.UpdateTGBalances(&_ContractTriggerXStakeRegistry.TransactOpts, keeper, user, TGAmount)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _ContractTriggerXStakeRegistry.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistrySession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _ContractTriggerXStakeRegistry.Contract.UpgradeTo(&_ContractTriggerXStakeRegistry.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _ContractTriggerXStakeRegistry.Contract.UpgradeTo(&_ContractTriggerXStakeRegistry.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ContractTriggerXStakeRegistry.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistrySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ContractTriggerXStakeRegistry.Contract.UpgradeToAndCall(&_ContractTriggerXStakeRegistry.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ContractTriggerXStakeRegistry.Contract.UpgradeToAndCall(&_ContractTriggerXStakeRegistry.TransactOpts, newImplementation, data)
}

// ContractTriggerXStakeRegistryAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the ContractTriggerXStakeRegistry contract.
type ContractTriggerXStakeRegistryAdminChangedIterator struct {
	Event *ContractTriggerXStakeRegistryAdminChanged // Event containing the contract specifics and raw log

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
func (it *ContractTriggerXStakeRegistryAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerXStakeRegistryAdminChanged)
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
		it.Event = new(ContractTriggerXStakeRegistryAdminChanged)
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
func (it *ContractTriggerXStakeRegistryAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerXStakeRegistryAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerXStakeRegistryAdminChanged represents a AdminChanged event raised by the ContractTriggerXStakeRegistry contract.
type ContractTriggerXStakeRegistryAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*ContractTriggerXStakeRegistryAdminChangedIterator, error) {

	logs, sub, err := _ContractTriggerXStakeRegistry.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXStakeRegistryAdminChangedIterator{contract: _ContractTriggerXStakeRegistry.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *ContractTriggerXStakeRegistryAdminChanged) (event.Subscription, error) {

	logs, sub, err := _ContractTriggerXStakeRegistry.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerXStakeRegistryAdminChanged)
				if err := _ContractTriggerXStakeRegistry.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryFilterer) ParseAdminChanged(log types.Log) (*ContractTriggerXStakeRegistryAdminChanged, error) {
	event := new(ContractTriggerXStakeRegistryAdminChanged)
	if err := _ContractTriggerXStakeRegistry.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTriggerXStakeRegistryBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the ContractTriggerXStakeRegistry contract.
type ContractTriggerXStakeRegistryBeaconUpgradedIterator struct {
	Event *ContractTriggerXStakeRegistryBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *ContractTriggerXStakeRegistryBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerXStakeRegistryBeaconUpgraded)
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
		it.Event = new(ContractTriggerXStakeRegistryBeaconUpgraded)
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
func (it *ContractTriggerXStakeRegistryBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerXStakeRegistryBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerXStakeRegistryBeaconUpgraded represents a BeaconUpgraded event raised by the ContractTriggerXStakeRegistry contract.
type ContractTriggerXStakeRegistryBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*ContractTriggerXStakeRegistryBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _ContractTriggerXStakeRegistry.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXStakeRegistryBeaconUpgradedIterator{contract: _ContractTriggerXStakeRegistry.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *ContractTriggerXStakeRegistryBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _ContractTriggerXStakeRegistry.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerXStakeRegistryBeaconUpgraded)
				if err := _ContractTriggerXStakeRegistry.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryFilterer) ParseBeaconUpgraded(log types.Log) (*ContractTriggerXStakeRegistryBeaconUpgraded, error) {
	event := new(ContractTriggerXStakeRegistryBeaconUpgraded)
	if err := _ContractTriggerXStakeRegistry.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTriggerXStakeRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractTriggerXStakeRegistry contract.
type ContractTriggerXStakeRegistryInitializedIterator struct {
	Event *ContractTriggerXStakeRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *ContractTriggerXStakeRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerXStakeRegistryInitialized)
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
		it.Event = new(ContractTriggerXStakeRegistryInitialized)
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
func (it *ContractTriggerXStakeRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerXStakeRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerXStakeRegistryInitialized represents a Initialized event raised by the ContractTriggerXStakeRegistry contract.
type ContractTriggerXStakeRegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractTriggerXStakeRegistryInitializedIterator, error) {

	logs, sub, err := _ContractTriggerXStakeRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXStakeRegistryInitializedIterator{contract: _ContractTriggerXStakeRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractTriggerXStakeRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractTriggerXStakeRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerXStakeRegistryInitialized)
				if err := _ContractTriggerXStakeRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryFilterer) ParseInitialized(log types.Log) (*ContractTriggerXStakeRegistryInitialized, error) {
	event := new(ContractTriggerXStakeRegistryInitialized)
	if err := _ContractTriggerXStakeRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTriggerXStakeRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractTriggerXStakeRegistry contract.
type ContractTriggerXStakeRegistryOwnershipTransferredIterator struct {
	Event *ContractTriggerXStakeRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractTriggerXStakeRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerXStakeRegistryOwnershipTransferred)
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
		it.Event = new(ContractTriggerXStakeRegistryOwnershipTransferred)
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
func (it *ContractTriggerXStakeRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerXStakeRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerXStakeRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the ContractTriggerXStakeRegistry contract.
type ContractTriggerXStakeRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractTriggerXStakeRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractTriggerXStakeRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXStakeRegistryOwnershipTransferredIterator{contract: _ContractTriggerXStakeRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractTriggerXStakeRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractTriggerXStakeRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerXStakeRegistryOwnershipTransferred)
				if err := _ContractTriggerXStakeRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*ContractTriggerXStakeRegistryOwnershipTransferred, error) {
	event := new(ContractTriggerXStakeRegistryOwnershipTransferred)
	if err := _ContractTriggerXStakeRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTriggerXStakeRegistryRewardClaimedIterator is returned from FilterRewardClaimed and is used to iterate over the raw logs and unpacked data for RewardClaimed events raised by the ContractTriggerXStakeRegistry contract.
type ContractTriggerXStakeRegistryRewardClaimedIterator struct {
	Event *ContractTriggerXStakeRegistryRewardClaimed // Event containing the contract specifics and raw log

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
func (it *ContractTriggerXStakeRegistryRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerXStakeRegistryRewardClaimed)
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
		it.Event = new(ContractTriggerXStakeRegistryRewardClaimed)
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
func (it *ContractTriggerXStakeRegistryRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerXStakeRegistryRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerXStakeRegistryRewardClaimed represents a RewardClaimed event raised by the ContractTriggerXStakeRegistry contract.
type ContractTriggerXStakeRegistryRewardClaimed struct {
	User   common.Address
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardClaimed is a free log retrieval operation binding the contract event 0x106f923f993c2149d49b4255ff723acafa1f2d94393f561d3eda32ae348f7241.
//
// Solidity: event RewardClaimed(address indexed user, uint256 reward)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryFilterer) FilterRewardClaimed(opts *bind.FilterOpts, user []common.Address) (*ContractTriggerXStakeRegistryRewardClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ContractTriggerXStakeRegistry.contract.FilterLogs(opts, "RewardClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXStakeRegistryRewardClaimedIterator{contract: _ContractTriggerXStakeRegistry.contract, event: "RewardClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardClaimed is a free log subscription operation binding the contract event 0x106f923f993c2149d49b4255ff723acafa1f2d94393f561d3eda32ae348f7241.
//
// Solidity: event RewardClaimed(address indexed user, uint256 reward)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryFilterer) WatchRewardClaimed(opts *bind.WatchOpts, sink chan<- *ContractTriggerXStakeRegistryRewardClaimed, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ContractTriggerXStakeRegistry.contract.WatchLogs(opts, "RewardClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerXStakeRegistryRewardClaimed)
				if err := _ContractTriggerXStakeRegistry.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
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

// ParseRewardClaimed is a log parse operation binding the contract event 0x106f923f993c2149d49b4255ff723acafa1f2d94393f561d3eda32ae348f7241.
//
// Solidity: event RewardClaimed(address indexed user, uint256 reward)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryFilterer) ParseRewardClaimed(log types.Log) (*ContractTriggerXStakeRegistryRewardClaimed, error) {
	event := new(ContractTriggerXStakeRegistryRewardClaimed)
	if err := _ContractTriggerXStakeRegistry.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTriggerXStakeRegistryStakeRemovedIterator is returned from FilterStakeRemoved and is used to iterate over the raw logs and unpacked data for StakeRemoved events raised by the ContractTriggerXStakeRegistry contract.
type ContractTriggerXStakeRegistryStakeRemovedIterator struct {
	Event *ContractTriggerXStakeRegistryStakeRemoved // Event containing the contract specifics and raw log

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
func (it *ContractTriggerXStakeRegistryStakeRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerXStakeRegistryStakeRemoved)
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
		it.Event = new(ContractTriggerXStakeRegistryStakeRemoved)
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
func (it *ContractTriggerXStakeRegistryStakeRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerXStakeRegistryStakeRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerXStakeRegistryStakeRemoved represents a StakeRemoved event raised by the ContractTriggerXStakeRegistry contract.
type ContractTriggerXStakeRegistryStakeRemoved struct {
	User   common.Address
	Amount *big.Int
	Reason string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakeRemoved is a free log retrieval operation binding the contract event 0x10a70fabb79e2f945edb05c430e4487214db653fb7b327784d4095b6befc9ce6.
//
// Solidity: event StakeRemoved(address indexed user, uint256 amount, string reason)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryFilterer) FilterStakeRemoved(opts *bind.FilterOpts, user []common.Address) (*ContractTriggerXStakeRegistryStakeRemovedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ContractTriggerXStakeRegistry.contract.FilterLogs(opts, "StakeRemoved", userRule)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXStakeRegistryStakeRemovedIterator{contract: _ContractTriggerXStakeRegistry.contract, event: "StakeRemoved", logs: logs, sub: sub}, nil
}

// WatchStakeRemoved is a free log subscription operation binding the contract event 0x10a70fabb79e2f945edb05c430e4487214db653fb7b327784d4095b6befc9ce6.
//
// Solidity: event StakeRemoved(address indexed user, uint256 amount, string reason)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryFilterer) WatchStakeRemoved(opts *bind.WatchOpts, sink chan<- *ContractTriggerXStakeRegistryStakeRemoved, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ContractTriggerXStakeRegistry.contract.WatchLogs(opts, "StakeRemoved", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerXStakeRegistryStakeRemoved)
				if err := _ContractTriggerXStakeRegistry.contract.UnpackLog(event, "StakeRemoved", log); err != nil {
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

// ParseStakeRemoved is a log parse operation binding the contract event 0x10a70fabb79e2f945edb05c430e4487214db653fb7b327784d4095b6befc9ce6.
//
// Solidity: event StakeRemoved(address indexed user, uint256 amount, string reason)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryFilterer) ParseStakeRemoved(log types.Log) (*ContractTriggerXStakeRegistryStakeRemoved, error) {
	event := new(ContractTriggerXStakeRegistryStakeRemoved)
	if err := _ContractTriggerXStakeRegistry.contract.UnpackLog(event, "StakeRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTriggerXStakeRegistryStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the ContractTriggerXStakeRegistry contract.
type ContractTriggerXStakeRegistryStakedIterator struct {
	Event *ContractTriggerXStakeRegistryStaked // Event containing the contract specifics and raw log

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
func (it *ContractTriggerXStakeRegistryStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerXStakeRegistryStaked)
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
		it.Event = new(ContractTriggerXStakeRegistryStaked)
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
func (it *ContractTriggerXStakeRegistryStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerXStakeRegistryStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerXStakeRegistryStaked represents a Staked event raised by the ContractTriggerXStakeRegistry contract.
type ContractTriggerXStakeRegistryStaked struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address) (*ContractTriggerXStakeRegistryStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ContractTriggerXStakeRegistry.contract.FilterLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXStakeRegistryStakedIterator{contract: _ContractTriggerXStakeRegistry.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *ContractTriggerXStakeRegistryStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ContractTriggerXStakeRegistry.contract.WatchLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerXStakeRegistryStaked)
				if err := _ContractTriggerXStakeRegistry.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryFilterer) ParseStaked(log types.Log) (*ContractTriggerXStakeRegistryStaked, error) {
	event := new(ContractTriggerXStakeRegistryStaked)
	if err := _ContractTriggerXStakeRegistry.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTriggerXStakeRegistryTGClaimedIterator is returned from FilterTGClaimed and is used to iterate over the raw logs and unpacked data for TGClaimed events raised by the ContractTriggerXStakeRegistry contract.
type ContractTriggerXStakeRegistryTGClaimedIterator struct {
	Event *ContractTriggerXStakeRegistryTGClaimed // Event containing the contract specifics and raw log

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
func (it *ContractTriggerXStakeRegistryTGClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerXStakeRegistryTGClaimed)
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
		it.Event = new(ContractTriggerXStakeRegistryTGClaimed)
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
func (it *ContractTriggerXStakeRegistryTGClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerXStakeRegistryTGClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerXStakeRegistryTGClaimed represents a TGClaimed event raised by the ContractTriggerXStakeRegistry contract.
type ContractTriggerXStakeRegistryTGClaimed struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTGClaimed is a free log retrieval operation binding the contract event 0xbddee7230372ecc0cb1780b66efe6d1461ee2a8e5b70239f3ec20a7d85f0e928.
//
// Solidity: event TGClaimed(address indexed user, uint256 amount)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryFilterer) FilterTGClaimed(opts *bind.FilterOpts, user []common.Address) (*ContractTriggerXStakeRegistryTGClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ContractTriggerXStakeRegistry.contract.FilterLogs(opts, "TGClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXStakeRegistryTGClaimedIterator{contract: _ContractTriggerXStakeRegistry.contract, event: "TGClaimed", logs: logs, sub: sub}, nil
}

// WatchTGClaimed is a free log subscription operation binding the contract event 0xbddee7230372ecc0cb1780b66efe6d1461ee2a8e5b70239f3ec20a7d85f0e928.
//
// Solidity: event TGClaimed(address indexed user, uint256 amount)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryFilterer) WatchTGClaimed(opts *bind.WatchOpts, sink chan<- *ContractTriggerXStakeRegistryTGClaimed, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ContractTriggerXStakeRegistry.contract.WatchLogs(opts, "TGClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerXStakeRegistryTGClaimed)
				if err := _ContractTriggerXStakeRegistry.contract.UnpackLog(event, "TGClaimed", log); err != nil {
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

// ParseTGClaimed is a log parse operation binding the contract event 0xbddee7230372ecc0cb1780b66efe6d1461ee2a8e5b70239f3ec20a7d85f0e928.
//
// Solidity: event TGClaimed(address indexed user, uint256 amount)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryFilterer) ParseTGClaimed(log types.Log) (*ContractTriggerXStakeRegistryTGClaimed, error) {
	event := new(ContractTriggerXStakeRegistryTGClaimed)
	if err := _ContractTriggerXStakeRegistry.contract.UnpackLog(event, "TGClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTriggerXStakeRegistryTGTransferredIterator is returned from FilterTGTransferred and is used to iterate over the raw logs and unpacked data for TGTransferred events raised by the ContractTriggerXStakeRegistry contract.
type ContractTriggerXStakeRegistryTGTransferredIterator struct {
	Event *ContractTriggerXStakeRegistryTGTransferred // Event containing the contract specifics and raw log

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
func (it *ContractTriggerXStakeRegistryTGTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerXStakeRegistryTGTransferred)
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
		it.Event = new(ContractTriggerXStakeRegistryTGTransferred)
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
func (it *ContractTriggerXStakeRegistryTGTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerXStakeRegistryTGTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerXStakeRegistryTGTransferred represents a TGTransferred event raised by the ContractTriggerXStakeRegistry contract.
type ContractTriggerXStakeRegistryTGTransferred struct {
	User   common.Address
	Keeper common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTGTransferred is a free log retrieval operation binding the contract event 0xa613b7e352cf78a30d92dbc93071265034f43edbeb7833b8d8713a3b4f323670.
//
// Solidity: event TGTransferred(address indexed user, address indexed keeper, uint256 amount)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryFilterer) FilterTGTransferred(opts *bind.FilterOpts, user []common.Address, keeper []common.Address) (*ContractTriggerXStakeRegistryTGTransferredIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}

	logs, sub, err := _ContractTriggerXStakeRegistry.contract.FilterLogs(opts, "TGTransferred", userRule, keeperRule)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXStakeRegistryTGTransferredIterator{contract: _ContractTriggerXStakeRegistry.contract, event: "TGTransferred", logs: logs, sub: sub}, nil
}

// WatchTGTransferred is a free log subscription operation binding the contract event 0xa613b7e352cf78a30d92dbc93071265034f43edbeb7833b8d8713a3b4f323670.
//
// Solidity: event TGTransferred(address indexed user, address indexed keeper, uint256 amount)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryFilterer) WatchTGTransferred(opts *bind.WatchOpts, sink chan<- *ContractTriggerXStakeRegistryTGTransferred, user []common.Address, keeper []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}

	logs, sub, err := _ContractTriggerXStakeRegistry.contract.WatchLogs(opts, "TGTransferred", userRule, keeperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerXStakeRegistryTGTransferred)
				if err := _ContractTriggerXStakeRegistry.contract.UnpackLog(event, "TGTransferred", log); err != nil {
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

// ParseTGTransferred is a log parse operation binding the contract event 0xa613b7e352cf78a30d92dbc93071265034f43edbeb7833b8d8713a3b4f323670.
//
// Solidity: event TGTransferred(address indexed user, address indexed keeper, uint256 amount)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryFilterer) ParseTGTransferred(log types.Log) (*ContractTriggerXStakeRegistryTGTransferred, error) {
	event := new(ContractTriggerXStakeRegistryTGTransferred)
	if err := _ContractTriggerXStakeRegistry.contract.UnpackLog(event, "TGTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTriggerXStakeRegistryTaskFeeClaimedIterator is returned from FilterTaskFeeClaimed and is used to iterate over the raw logs and unpacked data for TaskFeeClaimed events raised by the ContractTriggerXStakeRegistry contract.
type ContractTriggerXStakeRegistryTaskFeeClaimedIterator struct {
	Event *ContractTriggerXStakeRegistryTaskFeeClaimed // Event containing the contract specifics and raw log

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
func (it *ContractTriggerXStakeRegistryTaskFeeClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerXStakeRegistryTaskFeeClaimed)
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
		it.Event = new(ContractTriggerXStakeRegistryTaskFeeClaimed)
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
func (it *ContractTriggerXStakeRegistryTaskFeeClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerXStakeRegistryTaskFeeClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerXStakeRegistryTaskFeeClaimed represents a TaskFeeClaimed event raised by the ContractTriggerXStakeRegistry contract.
type ContractTriggerXStakeRegistryTaskFeeClaimed struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTaskFeeClaimed is a free log retrieval operation binding the contract event 0xae8ecfb3533b83d346a9f325b0e61ccdd5a8d8aaa0fa2dbfcbea06121120da40.
//
// Solidity: event TaskFeeClaimed(address indexed user, uint256 amount)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryFilterer) FilterTaskFeeClaimed(opts *bind.FilterOpts, user []common.Address) (*ContractTriggerXStakeRegistryTaskFeeClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ContractTriggerXStakeRegistry.contract.FilterLogs(opts, "TaskFeeClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXStakeRegistryTaskFeeClaimedIterator{contract: _ContractTriggerXStakeRegistry.contract, event: "TaskFeeClaimed", logs: logs, sub: sub}, nil
}

// WatchTaskFeeClaimed is a free log subscription operation binding the contract event 0xae8ecfb3533b83d346a9f325b0e61ccdd5a8d8aaa0fa2dbfcbea06121120da40.
//
// Solidity: event TaskFeeClaimed(address indexed user, uint256 amount)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryFilterer) WatchTaskFeeClaimed(opts *bind.WatchOpts, sink chan<- *ContractTriggerXStakeRegistryTaskFeeClaimed, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ContractTriggerXStakeRegistry.contract.WatchLogs(opts, "TaskFeeClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerXStakeRegistryTaskFeeClaimed)
				if err := _ContractTriggerXStakeRegistry.contract.UnpackLog(event, "TaskFeeClaimed", log); err != nil {
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

// ParseTaskFeeClaimed is a log parse operation binding the contract event 0xae8ecfb3533b83d346a9f325b0e61ccdd5a8d8aaa0fa2dbfcbea06121120da40.
//
// Solidity: event TaskFeeClaimed(address indexed user, uint256 amount)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryFilterer) ParseTaskFeeClaimed(log types.Log) (*ContractTriggerXStakeRegistryTaskFeeClaimed, error) {
	event := new(ContractTriggerXStakeRegistryTaskFeeClaimed)
	if err := _ContractTriggerXStakeRegistry.contract.UnpackLog(event, "TaskFeeClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTriggerXStakeRegistryUnstakedIterator is returned from FilterUnstaked and is used to iterate over the raw logs and unpacked data for Unstaked events raised by the ContractTriggerXStakeRegistry contract.
type ContractTriggerXStakeRegistryUnstakedIterator struct {
	Event *ContractTriggerXStakeRegistryUnstaked // Event containing the contract specifics and raw log

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
func (it *ContractTriggerXStakeRegistryUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerXStakeRegistryUnstaked)
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
		it.Event = new(ContractTriggerXStakeRegistryUnstaked)
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
func (it *ContractTriggerXStakeRegistryUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerXStakeRegistryUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerXStakeRegistryUnstaked represents a Unstaked event raised by the ContractTriggerXStakeRegistry contract.
type ContractTriggerXStakeRegistryUnstaked struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnstaked is a free log retrieval operation binding the contract event 0x0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f75.
//
// Solidity: event Unstaked(address indexed user, uint256 amount)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryFilterer) FilterUnstaked(opts *bind.FilterOpts, user []common.Address) (*ContractTriggerXStakeRegistryUnstakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ContractTriggerXStakeRegistry.contract.FilterLogs(opts, "Unstaked", userRule)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXStakeRegistryUnstakedIterator{contract: _ContractTriggerXStakeRegistry.contract, event: "Unstaked", logs: logs, sub: sub}, nil
}

// WatchUnstaked is a free log subscription operation binding the contract event 0x0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f75.
//
// Solidity: event Unstaked(address indexed user, uint256 amount)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryFilterer) WatchUnstaked(opts *bind.WatchOpts, sink chan<- *ContractTriggerXStakeRegistryUnstaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ContractTriggerXStakeRegistry.contract.WatchLogs(opts, "Unstaked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerXStakeRegistryUnstaked)
				if err := _ContractTriggerXStakeRegistry.contract.UnpackLog(event, "Unstaked", log); err != nil {
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

// ParseUnstaked is a log parse operation binding the contract event 0x0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f75.
//
// Solidity: event Unstaked(address indexed user, uint256 amount)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryFilterer) ParseUnstaked(log types.Log) (*ContractTriggerXStakeRegistryUnstaked, error) {
	event := new(ContractTriggerXStakeRegistryUnstaked)
	if err := _ContractTriggerXStakeRegistry.contract.UnpackLog(event, "Unstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTriggerXStakeRegistryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ContractTriggerXStakeRegistry contract.
type ContractTriggerXStakeRegistryUpgradedIterator struct {
	Event *ContractTriggerXStakeRegistryUpgraded // Event containing the contract specifics and raw log

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
func (it *ContractTriggerXStakeRegistryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerXStakeRegistryUpgraded)
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
		it.Event = new(ContractTriggerXStakeRegistryUpgraded)
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
func (it *ContractTriggerXStakeRegistryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerXStakeRegistryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerXStakeRegistryUpgraded represents a Upgraded event raised by the ContractTriggerXStakeRegistry contract.
type ContractTriggerXStakeRegistryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ContractTriggerXStakeRegistryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ContractTriggerXStakeRegistry.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXStakeRegistryUpgradedIterator{contract: _ContractTriggerXStakeRegistry.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ContractTriggerXStakeRegistryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ContractTriggerXStakeRegistry.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerXStakeRegistryUpgraded)
				if err := _ContractTriggerXStakeRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_ContractTriggerXStakeRegistry *ContractTriggerXStakeRegistryFilterer) ParseUpgraded(log types.Log) (*ContractTriggerXStakeRegistryUpgraded, error) {
	event := new(ContractTriggerXStakeRegistryUpgraded)
	if err := _ContractTriggerXStakeRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
