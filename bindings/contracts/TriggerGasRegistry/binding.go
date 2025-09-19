// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractTriggerGasRegistry

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

// ContractTriggerGasRegistryMetaData contains all meta data concerning the ContractTriggerGasRegistry contract.
var ContractTriggerGasRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"ETHWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TGBalanceDeducted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"TGBalanceRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TGClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tgPerEth\",\"type\":\"uint256\"}],\"name\":\"TGPerETHUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tgAmount\",\"type\":\"uint256\"}],\"name\":\"TGPurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TGRefunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TGTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TaskFeeClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"TG_PER_ETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ethSpent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"TGbalance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tgAmount\",\"type\":\"uint256\"}],\"name\":\"claimETHForTG\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tgAmount\",\"type\":\"uint256\"}],\"name\":\"deductTGBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ethSpent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tgBalance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tgPerEth\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operatorRole\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ethAmount\",\"type\":\"uint256\"}],\"name\":\"purchaseTG\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operatorRole\",\"type\":\"address\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tgPerEth\",\"type\":\"uint256\"}],\"name\":\"setTGPerETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526004361015610011575f80fd5b5f3560e01c8063168d6c5014610bb55780631794bb3c146109b257806327e235e31461097257806347503b3d1461094a5780634f1ef2861461076b57806350caaf901461063857806352d1902d146105d2578063715018a61461056b57806379dd15ff146104795780638da5cb5b14610445578063989c6f6a14610428578063ad3cb1cc146103dd578063b3ab15fb1461034c578063cc113622146101cd578063d0737b7a14610181578063f2fde38b146101565763f8b2cb4f146100d4575f80fd5b34610152576020366003190112610152576001600160a01b036100f5610c6c565b165f525f60205260405f20604051906040820182811067ffffffffffffffff82111761013e57604092602091845260018354938483520154918291015282519182526020820152f35b634e487b7160e01b5f52604160045260245ffd5b5f80fd5b346101525760203660031901126101525761017f610172610c6c565b61017a610e9a565b610df1565b005b34610152576020366003190112610152577f18652b6a656730019776361074bff5dbd2c837755a539e612061263f50b8345760206004356101c0610e9a565b80600155604051908152a1005b346101525760403660031901126101525760043560243567ffffffffffffffff8111610152573660238201121561015257610212903690602481600401359101610cc0565b61021a610e9a565b610222610e62565b811561030f574782116102ca577f68fc7e0340ad067ed9d1bcd121870cf25bd80b18a94d133b8bec22f4ea86b485906102805f8080808760018060a01b035f516020610f575f395f51905f5254165af161027a610d73565b50610da2565b60018060a01b035f516020610f575f395f51905f525416926102b46040519283928352604060208401526040830190610cf6565b0390a260015f516020610f975f395f51905f5255005b60405162461bcd60e51b815260206004820152601d60248201527f496e73756666696369656e7420636f6e74726163742062616c616e63650000006044820152606490fd5b60405162461bcd60e51b8152602060048201526015602482015274086c2dcdcdee840eed2e8d0c8e4c2ee4060408aa89605b1b6044820152606490fd5b3461015257602036600319011261015257610365610c6c565b61036d610e9a565b6001600160a01b03168015610398576bffffffffffffffffffffffff60a01b60025416176002555f80f35b60405162461bcd60e51b815260206004820152601c60248201527f4f70657261746f722063616e6e6f7420626520302061646472657373000000006044820152606490fd5b34610152575f366003190112610152576104246040516103fe604082610c82565b60058152640352e302e360dc1b6020820152604051918291602083526020830190610cf6565b0390f35b34610152575f366003190112610152576020600154604051908152f35b34610152575f366003190112610152575f516020610f575f395f51905f52546040516001600160a01b039091168152602090f35b3461015257604036600319011261015257610492610c6c565b60025460243591906001600160a01b0316330361051a5760018060a01b031690815f525f6020526104cc81600160405f2001541015610d1a565b806104d357005b60207fef70e6d79855d76ed5d6ba11f8368a48ca19a9f11ffed71aa53caab2cf8c16a491835f525f8252600160405f200161050f828254610d66565b9055604051908152a2005b60405162461bcd60e51b8152602060048201526024808201527f4f6e6c79206f70657261746f722063616e2063616c6c20746869732066756e636044820152633a34b7b760e11b6064820152608490fd5b34610152575f36600319011261015257610583610e9a565b5f516020610f575f395f51905f5280546001600160a01b031981169091555f906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b34610152575f366003190112610152577f0000000000000000000000009f452702490e11af38c213de30c000a46e4da3996001600160a01b031630036106295760206040515f516020610f775f395f51905f528152f35b63703e46dd60e11b5f5260045ffd5b60203660031901126101525760043561064f610e62565b8015610731578034036106ec57335f525f60205260405f20610672828254610de4565b81556001549081830291838304036106d857600101610692828254610de4565b905560405191825260208201527fd29e7684851e87b600fbe9f3d04a3adda696b6768e9bde3c3ae69143b4efa2b560403392a260015f516020610f975f395f51905f5255005b634e487b7160e01b5f52601160045260245ffd5b60405162461bcd60e51b815260206004820152601a60248201527f53656e7420455448206d757374206d6174636820616d6f756e740000000000006044820152606490fd5b60405162461bcd60e51b8152602060048201526012602482015271086c2dcdcdee840e6e0cadcc84060408aa8960731b6044820152606490fd5b60403660031901126101525761077f610c6c565b60243567ffffffffffffffff81116101525736602382011215610152576107b0903690602481600401359101610cc0565b6001600160a01b037f0000000000000000000000009f452702490e11af38c213de30c000a46e4da39916308114908115610928575b50610629576107f2610e9a565b6040516352d1902d60e01b81526001600160a01b0383169290602081600481875afa5f91816108f4575b506108345783634c9c8ce360e01b5f5260045260245ffd5b805f516020610f775f395f51905f528592036108e25750813b156108d0575f516020610f775f395f51905f5280546001600160a01b031916821790557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b5f80a28151156108b8575f8083602061017f95519101845af46108b2610d73565b91610ef8565b5050346108c157005b63b398979f60e01b5f5260045ffd5b634c9c8ce360e01b5f5260045260245ffd5b632a87526960e21b5f5260045260245ffd5b9091506020813d602011610920575b8161091060209383610c82565b810103126101525751908561081c565b3d9150610903565b5f516020610f775f395f51905f52546001600160a01b031614159050836107e5565b34610152575f366003190112610152576002546040516001600160a01b039091168152602090f35b34610152576020366003190112610152576001600160a01b03610993610c6c565b165f525f6020526040805f206001815491015482519182526020820152f35b34610152576060366003190112610152576109cb610c6c565b6024356001600160a01b03811690819003610152575f516020610fb75f395f51905f52549160ff8360401c16159267ffffffffffffffff811680159081610bad575b6001149081610ba3575b159081610b9a575b50610b8b5767ffffffffffffffff1981166001175f516020610fb75f395f51905f525583610b5f575b506001600160a01b03811615610b1057610a8d90610a64610ecd565b610a6c610ecd565b60015f516020610f975f395f51905f5255610a85610ecd565b61017a610ecd565b610a95610ecd565b6044356001556bffffffffffffffffffffffff60a01b6002541617600255610ab957005b68ff0000000000000000195f516020610fb75f395f51905f5254165f516020610fb75f395f51905f52557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a1005b60405162461bcd60e51b815260206004820152602160248201527f496e697469616c206f776e65722063616e6e6f742062652030206164647265736044820152607360f81b6064820152608490fd5b68ffffffffffffffffff191668010000000000000001175f516020610fb75f395f51905f525583610a48565b63f92ee8a960e01b5f5260045ffd5b90501585610a1f565b303b159150610a17565b859150610a0d565b3461015257602036600319011261015257600435610bd1610e62565b335f525f602052600160405f20018054610bed83821015610d1a565b6001548015610c58575f80809392610c1995610c0a888496610d66565b90558604335af161027a610d73565b6040519081527fbddee7230372ecc0cb1780b66efe6d1461ee2a8e5b70239f3ec20a7d85f0e92860203392a260015f516020610f975f395f51905f5255005b634e487b7160e01b5f52601260045260245ffd5b600435906001600160a01b038216820361015257565b90601f8019910116810190811067ffffffffffffffff82111761013e57604052565b67ffffffffffffffff811161013e57601f01601f191660200190565b929192610ccc82610ca4565b91610cda6040519384610c82565b829481845281830111610152578281602093845f960137010152565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b15610d2157565b60405162461bcd60e51b815260206004820152601760248201527f496e73756666696369656e742054472062616c616e63650000000000000000006044820152606490fd5b919082039182116106d857565b3d15610d9d573d90610d8482610ca4565b91610d926040519384610c82565b82523d5f602084013e565b606090565b15610da957565b60405162461bcd60e51b8152602060048201526013602482015272115512081d1c985b9cd9995c8819985a5b1959606a1b6044820152606490fd5b919082018092116106d857565b6001600160a01b03168015610e4f575f516020610f575f395f51905f5280546001600160a01b0319811683179091556001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3565b631e4fbdf760e01b5f525f60045260245ffd5b60025f516020610f975f395f51905f525414610e8b5760025f516020610f975f395f51905f5255565b633ee5aeb560e01b5f5260045ffd5b5f516020610f575f395f51905f52546001600160a01b03163303610eba57565b63118cdaa760e01b5f523360045260245ffd5b60ff5f516020610fb75f395f51905f525460401c1615610ee957565b631afcd79f60e31b5f5260045ffd5b90610f1c5750805115610f0d57805190602001fd5b63d6bda27560e01b5f5260045ffd5b81511580610f4d575b610f2d575090565b639996b31560e01b5f9081526001600160a01b0391909116600452602490fd5b50803b15610f2556fe9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a26469706673582212205a7b062f929a30d0aafb16f690fc3f89f9ae0a5361c3c2446bac5df043ddcf3164736f6c634300081b0033",
}

// ContractTriggerGasRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractTriggerGasRegistryMetaData.ABI instead.
var ContractTriggerGasRegistryABI = ContractTriggerGasRegistryMetaData.ABI

// ContractTriggerGasRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractTriggerGasRegistryMetaData.Bin instead.
var ContractTriggerGasRegistryBin = ContractTriggerGasRegistryMetaData.Bin

// DeployContractTriggerGasRegistry deploys a new Ethereum contract, binding an instance of ContractTriggerGasRegistry to it.
func DeployContractTriggerGasRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContractTriggerGasRegistry, error) {
	parsed, err := ContractTriggerGasRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractTriggerGasRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractTriggerGasRegistry{ContractTriggerGasRegistryCaller: ContractTriggerGasRegistryCaller{contract: contract}, ContractTriggerGasRegistryTransactor: ContractTriggerGasRegistryTransactor{contract: contract}, ContractTriggerGasRegistryFilterer: ContractTriggerGasRegistryFilterer{contract: contract}}, nil
}

// ContractTriggerGasRegistryMethods is an auto generated interface around an Ethereum contract.
type ContractTriggerGasRegistryMethods interface {
	ContractTriggerGasRegistryCalls
	ContractTriggerGasRegistryTransacts
	ContractTriggerGasRegistryFilters
}

// ContractTriggerGasRegistryCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractTriggerGasRegistryCalls interface {
	TGPERETH(opts *bind.CallOpts) (*big.Int, error)

	UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error)

	Balances(opts *bind.CallOpts, arg0 common.Address) (struct {
		EthSpent  *big.Int
		TGbalance *big.Int
	}, error)

	GetBalance(opts *bind.CallOpts, user common.Address) (struct {
		EthSpent  *big.Int
		TgBalance *big.Int
	}, error)

	OperatorRole(opts *bind.CallOpts) (common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	ProxiableUUID(opts *bind.CallOpts) ([32]byte, error)
}

// ContractTriggerGasRegistryTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractTriggerGasRegistryTransacts interface {
	ClaimETHForTG(opts *bind.TransactOpts, tgAmount *big.Int) (*types.Transaction, error)

	DeductTGBalance(opts *bind.TransactOpts, user common.Address, tgAmount *big.Int) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts, initialOwner common.Address, _operator common.Address, _tgPerEth *big.Int) (*types.Transaction, error)

	PurchaseTG(opts *bind.TransactOpts, ethAmount *big.Int) (*types.Transaction, error)

	RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetOperator(opts *bind.TransactOpts, _operatorRole common.Address) (*types.Transaction, error)

	SetTGPerETH(opts *bind.TransactOpts, _tgPerEth *big.Int) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error)

	UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error)

	WithdrawETH(opts *bind.TransactOpts, amount *big.Int, reason string) (*types.Transaction, error)
}

// ContractTriggerGasRegistryFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractTriggerGasRegistryFilters interface {
	FilterETHWithdrawn(opts *bind.FilterOpts, owner []common.Address) (*ContractTriggerGasRegistryETHWithdrawnIterator, error)
	WatchETHWithdrawn(opts *bind.WatchOpts, sink chan<- *ContractTriggerGasRegistryETHWithdrawn, owner []common.Address) (event.Subscription, error)
	ParseETHWithdrawn(log types.Log) (*ContractTriggerGasRegistryETHWithdrawn, error)

	FilterInitialized(opts *bind.FilterOpts) (*ContractTriggerGasRegistryInitializedIterator, error)
	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractTriggerGasRegistryInitialized) (event.Subscription, error)
	ParseInitialized(log types.Log) (*ContractTriggerGasRegistryInitialized, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractTriggerGasRegistryOwnershipTransferredIterator, error)
	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractTriggerGasRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error)
	ParseOwnershipTransferred(log types.Log) (*ContractTriggerGasRegistryOwnershipTransferred, error)

	FilterRewardClaimed(opts *bind.FilterOpts, user []common.Address) (*ContractTriggerGasRegistryRewardClaimedIterator, error)
	WatchRewardClaimed(opts *bind.WatchOpts, sink chan<- *ContractTriggerGasRegistryRewardClaimed, user []common.Address) (event.Subscription, error)
	ParseRewardClaimed(log types.Log) (*ContractTriggerGasRegistryRewardClaimed, error)

	FilterTGBalanceDeducted(opts *bind.FilterOpts, user []common.Address) (*ContractTriggerGasRegistryTGBalanceDeductedIterator, error)
	WatchTGBalanceDeducted(opts *bind.WatchOpts, sink chan<- *ContractTriggerGasRegistryTGBalanceDeducted, user []common.Address) (event.Subscription, error)
	ParseTGBalanceDeducted(log types.Log) (*ContractTriggerGasRegistryTGBalanceDeducted, error)

	FilterTGBalanceRemoved(opts *bind.FilterOpts, user []common.Address) (*ContractTriggerGasRegistryTGBalanceRemovedIterator, error)
	WatchTGBalanceRemoved(opts *bind.WatchOpts, sink chan<- *ContractTriggerGasRegistryTGBalanceRemoved, user []common.Address) (event.Subscription, error)
	ParseTGBalanceRemoved(log types.Log) (*ContractTriggerGasRegistryTGBalanceRemoved, error)

	FilterTGClaimed(opts *bind.FilterOpts, user []common.Address) (*ContractTriggerGasRegistryTGClaimedIterator, error)
	WatchTGClaimed(opts *bind.WatchOpts, sink chan<- *ContractTriggerGasRegistryTGClaimed, user []common.Address) (event.Subscription, error)
	ParseTGClaimed(log types.Log) (*ContractTriggerGasRegistryTGClaimed, error)

	FilterTGPerETHUpdated(opts *bind.FilterOpts) (*ContractTriggerGasRegistryTGPerETHUpdatedIterator, error)
	WatchTGPerETHUpdated(opts *bind.WatchOpts, sink chan<- *ContractTriggerGasRegistryTGPerETHUpdated) (event.Subscription, error)
	ParseTGPerETHUpdated(log types.Log) (*ContractTriggerGasRegistryTGPerETHUpdated, error)

	FilterTGPurchased(opts *bind.FilterOpts, user []common.Address) (*ContractTriggerGasRegistryTGPurchasedIterator, error)
	WatchTGPurchased(opts *bind.WatchOpts, sink chan<- *ContractTriggerGasRegistryTGPurchased, user []common.Address) (event.Subscription, error)
	ParseTGPurchased(log types.Log) (*ContractTriggerGasRegistryTGPurchased, error)

	FilterTGRefunded(opts *bind.FilterOpts, user []common.Address) (*ContractTriggerGasRegistryTGRefundedIterator, error)
	WatchTGRefunded(opts *bind.WatchOpts, sink chan<- *ContractTriggerGasRegistryTGRefunded, user []common.Address) (event.Subscription, error)
	ParseTGRefunded(log types.Log) (*ContractTriggerGasRegistryTGRefunded, error)

	FilterTGTransferred(opts *bind.FilterOpts, user []common.Address, keeper []common.Address) (*ContractTriggerGasRegistryTGTransferredIterator, error)
	WatchTGTransferred(opts *bind.WatchOpts, sink chan<- *ContractTriggerGasRegistryTGTransferred, user []common.Address, keeper []common.Address) (event.Subscription, error)
	ParseTGTransferred(log types.Log) (*ContractTriggerGasRegistryTGTransferred, error)

	FilterTaskFeeClaimed(opts *bind.FilterOpts, user []common.Address) (*ContractTriggerGasRegistryTaskFeeClaimedIterator, error)
	WatchTaskFeeClaimed(opts *bind.WatchOpts, sink chan<- *ContractTriggerGasRegistryTaskFeeClaimed, user []common.Address) (event.Subscription, error)
	ParseTaskFeeClaimed(log types.Log) (*ContractTriggerGasRegistryTaskFeeClaimed, error)

	FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ContractTriggerGasRegistryUpgradedIterator, error)
	WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ContractTriggerGasRegistryUpgraded, implementation []common.Address) (event.Subscription, error)
	ParseUpgraded(log types.Log) (*ContractTriggerGasRegistryUpgraded, error)
}

// ContractTriggerGasRegistry is an auto generated Go binding around an Ethereum contract.
type ContractTriggerGasRegistry struct {
	ContractTriggerGasRegistryCaller     // Read-only binding to the contract
	ContractTriggerGasRegistryTransactor // Write-only binding to the contract
	ContractTriggerGasRegistryFilterer   // Log filterer for contract events
}

// ContractTriggerGasRegistry implements the ContractTriggerGasRegistryMethods interface.
var _ ContractTriggerGasRegistryMethods = (*ContractTriggerGasRegistry)(nil)

// ContractTriggerGasRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractTriggerGasRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTriggerGasRegistryCaller implements the ContractTriggerGasRegistryCalls interface.
var _ ContractTriggerGasRegistryCalls = (*ContractTriggerGasRegistryCaller)(nil)

// ContractTriggerGasRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTriggerGasRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTriggerGasRegistryTransactor implements the ContractTriggerGasRegistryTransacts interface.
var _ ContractTriggerGasRegistryTransacts = (*ContractTriggerGasRegistryTransactor)(nil)

// ContractTriggerGasRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractTriggerGasRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTriggerGasRegistryFilterer implements the ContractTriggerGasRegistryFilters interface.
var _ ContractTriggerGasRegistryFilters = (*ContractTriggerGasRegistryFilterer)(nil)

// ContractTriggerGasRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractTriggerGasRegistrySession struct {
	Contract     *ContractTriggerGasRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ContractTriggerGasRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractTriggerGasRegistryCallerSession struct {
	Contract *ContractTriggerGasRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// ContractTriggerGasRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTriggerGasRegistryTransactorSession struct {
	Contract     *ContractTriggerGasRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// ContractTriggerGasRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractTriggerGasRegistryRaw struct {
	Contract *ContractTriggerGasRegistry // Generic contract binding to access the raw methods on
}

// ContractTriggerGasRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractTriggerGasRegistryCallerRaw struct {
	Contract *ContractTriggerGasRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTriggerGasRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTriggerGasRegistryTransactorRaw struct {
	Contract *ContractTriggerGasRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractTriggerGasRegistry creates a new instance of ContractTriggerGasRegistry, bound to a specific deployed contract.
func NewContractTriggerGasRegistry(address common.Address, backend bind.ContractBackend) (*ContractTriggerGasRegistry, error) {
	contract, err := bindContractTriggerGasRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerGasRegistry{ContractTriggerGasRegistryCaller: ContractTriggerGasRegistryCaller{contract: contract}, ContractTriggerGasRegistryTransactor: ContractTriggerGasRegistryTransactor{contract: contract}, ContractTriggerGasRegistryFilterer: ContractTriggerGasRegistryFilterer{contract: contract}}, nil
}

// NewContractTriggerGasRegistryCaller creates a new read-only instance of ContractTriggerGasRegistry, bound to a specific deployed contract.
func NewContractTriggerGasRegistryCaller(address common.Address, caller bind.ContractCaller) (*ContractTriggerGasRegistryCaller, error) {
	contract, err := bindContractTriggerGasRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerGasRegistryCaller{contract: contract}, nil
}

// NewContractTriggerGasRegistryTransactor creates a new write-only instance of ContractTriggerGasRegistry, bound to a specific deployed contract.
func NewContractTriggerGasRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTriggerGasRegistryTransactor, error) {
	contract, err := bindContractTriggerGasRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerGasRegistryTransactor{contract: contract}, nil
}

// NewContractTriggerGasRegistryFilterer creates a new log filterer instance of ContractTriggerGasRegistry, bound to a specific deployed contract.
func NewContractTriggerGasRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractTriggerGasRegistryFilterer, error) {
	contract, err := bindContractTriggerGasRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerGasRegistryFilterer{contract: contract}, nil
}

// bindContractTriggerGasRegistry binds a generic wrapper to an already deployed contract.
func bindContractTriggerGasRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractTriggerGasRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractTriggerGasRegistry.Contract.ContractTriggerGasRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.Contract.ContractTriggerGasRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.Contract.ContractTriggerGasRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractTriggerGasRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.Contract.contract.Transact(opts, method, params...)
}

// TGPERETH is a free data retrieval call binding the contract method 0x989c6f6a.
//
// Solidity: function TG_PER_ETH() view returns(uint256)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryCaller) TGPERETH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractTriggerGasRegistry.contract.Call(opts, &out, "TG_PER_ETH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TGPERETH is a free data retrieval call binding the contract method 0x989c6f6a.
//
// Solidity: function TG_PER_ETH() view returns(uint256)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistrySession) TGPERETH() (*big.Int, error) {
	return _ContractTriggerGasRegistry.Contract.TGPERETH(&_ContractTriggerGasRegistry.CallOpts)
}

// TGPERETH is a free data retrieval call binding the contract method 0x989c6f6a.
//
// Solidity: function TG_PER_ETH() view returns(uint256)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryCallerSession) TGPERETH() (*big.Int, error) {
	return _ContractTriggerGasRegistry.Contract.TGPERETH(&_ContractTriggerGasRegistry.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ContractTriggerGasRegistry.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistrySession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ContractTriggerGasRegistry.Contract.UPGRADEINTERFACEVERSION(&_ContractTriggerGasRegistry.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ContractTriggerGasRegistry.Contract.UPGRADEINTERFACEVERSION(&_ContractTriggerGasRegistry.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256 ethSpent, uint256 TGbalance)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (struct {
	EthSpent  *big.Int
	TGbalance *big.Int
}, error) {
	var out []interface{}
	err := _ContractTriggerGasRegistry.contract.Call(opts, &out, "balances", arg0)

	outstruct := new(struct {
		EthSpent  *big.Int
		TGbalance *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EthSpent = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TGbalance = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256 ethSpent, uint256 TGbalance)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistrySession) Balances(arg0 common.Address) (struct {
	EthSpent  *big.Int
	TGbalance *big.Int
}, error) {
	return _ContractTriggerGasRegistry.Contract.Balances(&_ContractTriggerGasRegistry.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256 ethSpent, uint256 TGbalance)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryCallerSession) Balances(arg0 common.Address) (struct {
	EthSpent  *big.Int
	TGbalance *big.Int
}, error) {
	return _ContractTriggerGasRegistry.Contract.Balances(&_ContractTriggerGasRegistry.CallOpts, arg0)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address user) view returns(uint256 ethSpent, uint256 tgBalance)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryCaller) GetBalance(opts *bind.CallOpts, user common.Address) (struct {
	EthSpent  *big.Int
	TgBalance *big.Int
}, error) {
	var out []interface{}
	err := _ContractTriggerGasRegistry.contract.Call(opts, &out, "getBalance", user)

	outstruct := new(struct {
		EthSpent  *big.Int
		TgBalance *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EthSpent = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TgBalance = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address user) view returns(uint256 ethSpent, uint256 tgBalance)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistrySession) GetBalance(user common.Address) (struct {
	EthSpent  *big.Int
	TgBalance *big.Int
}, error) {
	return _ContractTriggerGasRegistry.Contract.GetBalance(&_ContractTriggerGasRegistry.CallOpts, user)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address user) view returns(uint256 ethSpent, uint256 tgBalance)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryCallerSession) GetBalance(user common.Address) (struct {
	EthSpent  *big.Int
	TgBalance *big.Int
}, error) {
	return _ContractTriggerGasRegistry.Contract.GetBalance(&_ContractTriggerGasRegistry.CallOpts, user)
}

// OperatorRole is a free data retrieval call binding the contract method 0x47503b3d.
//
// Solidity: function operatorRole() view returns(address)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryCaller) OperatorRole(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTriggerGasRegistry.contract.Call(opts, &out, "operatorRole")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OperatorRole is a free data retrieval call binding the contract method 0x47503b3d.
//
// Solidity: function operatorRole() view returns(address)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistrySession) OperatorRole() (common.Address, error) {
	return _ContractTriggerGasRegistry.Contract.OperatorRole(&_ContractTriggerGasRegistry.CallOpts)
}

// OperatorRole is a free data retrieval call binding the contract method 0x47503b3d.
//
// Solidity: function operatorRole() view returns(address)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryCallerSession) OperatorRole() (common.Address, error) {
	return _ContractTriggerGasRegistry.Contract.OperatorRole(&_ContractTriggerGasRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTriggerGasRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistrySession) Owner() (common.Address, error) {
	return _ContractTriggerGasRegistry.Contract.Owner(&_ContractTriggerGasRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryCallerSession) Owner() (common.Address, error) {
	return _ContractTriggerGasRegistry.Contract.Owner(&_ContractTriggerGasRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractTriggerGasRegistry.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistrySession) ProxiableUUID() ([32]byte, error) {
	return _ContractTriggerGasRegistry.Contract.ProxiableUUID(&_ContractTriggerGasRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _ContractTriggerGasRegistry.Contract.ProxiableUUID(&_ContractTriggerGasRegistry.CallOpts)
}

// ClaimETHForTG is a paid mutator transaction binding the contract method 0x168d6c50.
//
// Solidity: function claimETHForTG(uint256 tgAmount) returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryTransactor) ClaimETHForTG(opts *bind.TransactOpts, tgAmount *big.Int) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.contract.Transact(opts, "claimETHForTG", tgAmount)
}

// ClaimETHForTG is a paid mutator transaction binding the contract method 0x168d6c50.
//
// Solidity: function claimETHForTG(uint256 tgAmount) returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistrySession) ClaimETHForTG(tgAmount *big.Int) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.Contract.ClaimETHForTG(&_ContractTriggerGasRegistry.TransactOpts, tgAmount)
}

// ClaimETHForTG is a paid mutator transaction binding the contract method 0x168d6c50.
//
// Solidity: function claimETHForTG(uint256 tgAmount) returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryTransactorSession) ClaimETHForTG(tgAmount *big.Int) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.Contract.ClaimETHForTG(&_ContractTriggerGasRegistry.TransactOpts, tgAmount)
}

// DeductTGBalance is a paid mutator transaction binding the contract method 0x79dd15ff.
//
// Solidity: function deductTGBalance(address user, uint256 tgAmount) returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryTransactor) DeductTGBalance(opts *bind.TransactOpts, user common.Address, tgAmount *big.Int) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.contract.Transact(opts, "deductTGBalance", user, tgAmount)
}

// DeductTGBalance is a paid mutator transaction binding the contract method 0x79dd15ff.
//
// Solidity: function deductTGBalance(address user, uint256 tgAmount) returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistrySession) DeductTGBalance(user common.Address, tgAmount *big.Int) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.Contract.DeductTGBalance(&_ContractTriggerGasRegistry.TransactOpts, user, tgAmount)
}

// DeductTGBalance is a paid mutator transaction binding the contract method 0x79dd15ff.
//
// Solidity: function deductTGBalance(address user, uint256 tgAmount) returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryTransactorSession) DeductTGBalance(user common.Address, tgAmount *big.Int) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.Contract.DeductTGBalance(&_ContractTriggerGasRegistry.TransactOpts, user, tgAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address _operator, uint256 _tgPerEth) returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, _operator common.Address, _tgPerEth *big.Int) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.contract.Transact(opts, "initialize", initialOwner, _operator, _tgPerEth)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address _operator, uint256 _tgPerEth) returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistrySession) Initialize(initialOwner common.Address, _operator common.Address, _tgPerEth *big.Int) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.Contract.Initialize(&_ContractTriggerGasRegistry.TransactOpts, initialOwner, _operator, _tgPerEth)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address _operator, uint256 _tgPerEth) returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryTransactorSession) Initialize(initialOwner common.Address, _operator common.Address, _tgPerEth *big.Int) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.Contract.Initialize(&_ContractTriggerGasRegistry.TransactOpts, initialOwner, _operator, _tgPerEth)
}

// PurchaseTG is a paid mutator transaction binding the contract method 0x50caaf90.
//
// Solidity: function purchaseTG(uint256 ethAmount) payable returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryTransactor) PurchaseTG(opts *bind.TransactOpts, ethAmount *big.Int) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.contract.Transact(opts, "purchaseTG", ethAmount)
}

// PurchaseTG is a paid mutator transaction binding the contract method 0x50caaf90.
//
// Solidity: function purchaseTG(uint256 ethAmount) payable returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistrySession) PurchaseTG(ethAmount *big.Int) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.Contract.PurchaseTG(&_ContractTriggerGasRegistry.TransactOpts, ethAmount)
}

// PurchaseTG is a paid mutator transaction binding the contract method 0x50caaf90.
//
// Solidity: function purchaseTG(uint256 ethAmount) payable returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryTransactorSession) PurchaseTG(ethAmount *big.Int) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.Contract.PurchaseTG(&_ContractTriggerGasRegistry.TransactOpts, ethAmount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.Contract.RenounceOwnership(&_ContractTriggerGasRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.Contract.RenounceOwnership(&_ContractTriggerGasRegistry.TransactOpts)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operatorRole) returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryTransactor) SetOperator(opts *bind.TransactOpts, _operatorRole common.Address) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.contract.Transact(opts, "setOperator", _operatorRole)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operatorRole) returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistrySession) SetOperator(_operatorRole common.Address) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.Contract.SetOperator(&_ContractTriggerGasRegistry.TransactOpts, _operatorRole)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operatorRole) returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryTransactorSession) SetOperator(_operatorRole common.Address) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.Contract.SetOperator(&_ContractTriggerGasRegistry.TransactOpts, _operatorRole)
}

// SetTGPerETH is a paid mutator transaction binding the contract method 0xd0737b7a.
//
// Solidity: function setTGPerETH(uint256 _tgPerEth) returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryTransactor) SetTGPerETH(opts *bind.TransactOpts, _tgPerEth *big.Int) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.contract.Transact(opts, "setTGPerETH", _tgPerEth)
}

// SetTGPerETH is a paid mutator transaction binding the contract method 0xd0737b7a.
//
// Solidity: function setTGPerETH(uint256 _tgPerEth) returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistrySession) SetTGPerETH(_tgPerEth *big.Int) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.Contract.SetTGPerETH(&_ContractTriggerGasRegistry.TransactOpts, _tgPerEth)
}

// SetTGPerETH is a paid mutator transaction binding the contract method 0xd0737b7a.
//
// Solidity: function setTGPerETH(uint256 _tgPerEth) returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryTransactorSession) SetTGPerETH(_tgPerEth *big.Int) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.Contract.SetTGPerETH(&_ContractTriggerGasRegistry.TransactOpts, _tgPerEth)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.Contract.TransferOwnership(&_ContractTriggerGasRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.Contract.TransferOwnership(&_ContractTriggerGasRegistry.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistrySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.Contract.UpgradeToAndCall(&_ContractTriggerGasRegistry.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.Contract.UpgradeToAndCall(&_ContractTriggerGasRegistry.TransactOpts, newImplementation, data)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xcc113622.
//
// Solidity: function withdrawETH(uint256 amount, string reason) returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryTransactor) WithdrawETH(opts *bind.TransactOpts, amount *big.Int, reason string) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.contract.Transact(opts, "withdrawETH", amount, reason)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xcc113622.
//
// Solidity: function withdrawETH(uint256 amount, string reason) returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistrySession) WithdrawETH(amount *big.Int, reason string) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.Contract.WithdrawETH(&_ContractTriggerGasRegistry.TransactOpts, amount, reason)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xcc113622.
//
// Solidity: function withdrawETH(uint256 amount, string reason) returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryTransactorSession) WithdrawETH(amount *big.Int, reason string) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.Contract.WithdrawETH(&_ContractTriggerGasRegistry.TransactOpts, amount, reason)
}

// ContractTriggerGasRegistryETHWithdrawnIterator is returned from FilterETHWithdrawn and is used to iterate over the raw logs and unpacked data for ETHWithdrawn events raised by the ContractTriggerGasRegistry contract.
type ContractTriggerGasRegistryETHWithdrawnIterator struct {
	Event *ContractTriggerGasRegistryETHWithdrawn // Event containing the contract specifics and raw log

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
func (it *ContractTriggerGasRegistryETHWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerGasRegistryETHWithdrawn)
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
		it.Event = new(ContractTriggerGasRegistryETHWithdrawn)
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
func (it *ContractTriggerGasRegistryETHWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerGasRegistryETHWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerGasRegistryETHWithdrawn represents a ETHWithdrawn event raised by the ContractTriggerGasRegistry contract.
type ContractTriggerGasRegistryETHWithdrawn struct {
	Owner  common.Address
	Amount *big.Int
	Reason string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterETHWithdrawn is a free log retrieval operation binding the contract event 0x68fc7e0340ad067ed9d1bcd121870cf25bd80b18a94d133b8bec22f4ea86b485.
//
// Solidity: event ETHWithdrawn(address indexed owner, uint256 amount, string reason)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) FilterETHWithdrawn(opts *bind.FilterOpts, owner []common.Address) (*ContractTriggerGasRegistryETHWithdrawnIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ContractTriggerGasRegistry.contract.FilterLogs(opts, "ETHWithdrawn", ownerRule)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerGasRegistryETHWithdrawnIterator{contract: _ContractTriggerGasRegistry.contract, event: "ETHWithdrawn", logs: logs, sub: sub}, nil
}

// WatchETHWithdrawn is a free log subscription operation binding the contract event 0x68fc7e0340ad067ed9d1bcd121870cf25bd80b18a94d133b8bec22f4ea86b485.
//
// Solidity: event ETHWithdrawn(address indexed owner, uint256 amount, string reason)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) WatchETHWithdrawn(opts *bind.WatchOpts, sink chan<- *ContractTriggerGasRegistryETHWithdrawn, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ContractTriggerGasRegistry.contract.WatchLogs(opts, "ETHWithdrawn", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerGasRegistryETHWithdrawn)
				if err := _ContractTriggerGasRegistry.contract.UnpackLog(event, "ETHWithdrawn", log); err != nil {
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

// ParseETHWithdrawn is a log parse operation binding the contract event 0x68fc7e0340ad067ed9d1bcd121870cf25bd80b18a94d133b8bec22f4ea86b485.
//
// Solidity: event ETHWithdrawn(address indexed owner, uint256 amount, string reason)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) ParseETHWithdrawn(log types.Log) (*ContractTriggerGasRegistryETHWithdrawn, error) {
	event := new(ContractTriggerGasRegistryETHWithdrawn)
	if err := _ContractTriggerGasRegistry.contract.UnpackLog(event, "ETHWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTriggerGasRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractTriggerGasRegistry contract.
type ContractTriggerGasRegistryInitializedIterator struct {
	Event *ContractTriggerGasRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *ContractTriggerGasRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerGasRegistryInitialized)
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
		it.Event = new(ContractTriggerGasRegistryInitialized)
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
func (it *ContractTriggerGasRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerGasRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerGasRegistryInitialized represents a Initialized event raised by the ContractTriggerGasRegistry contract.
type ContractTriggerGasRegistryInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractTriggerGasRegistryInitializedIterator, error) {

	logs, sub, err := _ContractTriggerGasRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractTriggerGasRegistryInitializedIterator{contract: _ContractTriggerGasRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractTriggerGasRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractTriggerGasRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerGasRegistryInitialized)
				if err := _ContractTriggerGasRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) ParseInitialized(log types.Log) (*ContractTriggerGasRegistryInitialized, error) {
	event := new(ContractTriggerGasRegistryInitialized)
	if err := _ContractTriggerGasRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTriggerGasRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractTriggerGasRegistry contract.
type ContractTriggerGasRegistryOwnershipTransferredIterator struct {
	Event *ContractTriggerGasRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractTriggerGasRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerGasRegistryOwnershipTransferred)
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
		it.Event = new(ContractTriggerGasRegistryOwnershipTransferred)
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
func (it *ContractTriggerGasRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerGasRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerGasRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the ContractTriggerGasRegistry contract.
type ContractTriggerGasRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractTriggerGasRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractTriggerGasRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerGasRegistryOwnershipTransferredIterator{contract: _ContractTriggerGasRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractTriggerGasRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractTriggerGasRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerGasRegistryOwnershipTransferred)
				if err := _ContractTriggerGasRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*ContractTriggerGasRegistryOwnershipTransferred, error) {
	event := new(ContractTriggerGasRegistryOwnershipTransferred)
	if err := _ContractTriggerGasRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTriggerGasRegistryRewardClaimedIterator is returned from FilterRewardClaimed and is used to iterate over the raw logs and unpacked data for RewardClaimed events raised by the ContractTriggerGasRegistry contract.
type ContractTriggerGasRegistryRewardClaimedIterator struct {
	Event *ContractTriggerGasRegistryRewardClaimed // Event containing the contract specifics and raw log

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
func (it *ContractTriggerGasRegistryRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerGasRegistryRewardClaimed)
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
		it.Event = new(ContractTriggerGasRegistryRewardClaimed)
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
func (it *ContractTriggerGasRegistryRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerGasRegistryRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerGasRegistryRewardClaimed represents a RewardClaimed event raised by the ContractTriggerGasRegistry contract.
type ContractTriggerGasRegistryRewardClaimed struct {
	User   common.Address
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardClaimed is a free log retrieval operation binding the contract event 0x106f923f993c2149d49b4255ff723acafa1f2d94393f561d3eda32ae348f7241.
//
// Solidity: event RewardClaimed(address indexed user, uint256 reward)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) FilterRewardClaimed(opts *bind.FilterOpts, user []common.Address) (*ContractTriggerGasRegistryRewardClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ContractTriggerGasRegistry.contract.FilterLogs(opts, "RewardClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerGasRegistryRewardClaimedIterator{contract: _ContractTriggerGasRegistry.contract, event: "RewardClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardClaimed is a free log subscription operation binding the contract event 0x106f923f993c2149d49b4255ff723acafa1f2d94393f561d3eda32ae348f7241.
//
// Solidity: event RewardClaimed(address indexed user, uint256 reward)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) WatchRewardClaimed(opts *bind.WatchOpts, sink chan<- *ContractTriggerGasRegistryRewardClaimed, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ContractTriggerGasRegistry.contract.WatchLogs(opts, "RewardClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerGasRegistryRewardClaimed)
				if err := _ContractTriggerGasRegistry.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
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
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) ParseRewardClaimed(log types.Log) (*ContractTriggerGasRegistryRewardClaimed, error) {
	event := new(ContractTriggerGasRegistryRewardClaimed)
	if err := _ContractTriggerGasRegistry.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTriggerGasRegistryTGBalanceDeductedIterator is returned from FilterTGBalanceDeducted and is used to iterate over the raw logs and unpacked data for TGBalanceDeducted events raised by the ContractTriggerGasRegistry contract.
type ContractTriggerGasRegistryTGBalanceDeductedIterator struct {
	Event *ContractTriggerGasRegistryTGBalanceDeducted // Event containing the contract specifics and raw log

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
func (it *ContractTriggerGasRegistryTGBalanceDeductedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerGasRegistryTGBalanceDeducted)
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
		it.Event = new(ContractTriggerGasRegistryTGBalanceDeducted)
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
func (it *ContractTriggerGasRegistryTGBalanceDeductedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerGasRegistryTGBalanceDeductedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerGasRegistryTGBalanceDeducted represents a TGBalanceDeducted event raised by the ContractTriggerGasRegistry contract.
type ContractTriggerGasRegistryTGBalanceDeducted struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTGBalanceDeducted is a free log retrieval operation binding the contract event 0xef70e6d79855d76ed5d6ba11f8368a48ca19a9f11ffed71aa53caab2cf8c16a4.
//
// Solidity: event TGBalanceDeducted(address indexed user, uint256 amount)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) FilterTGBalanceDeducted(opts *bind.FilterOpts, user []common.Address) (*ContractTriggerGasRegistryTGBalanceDeductedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ContractTriggerGasRegistry.contract.FilterLogs(opts, "TGBalanceDeducted", userRule)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerGasRegistryTGBalanceDeductedIterator{contract: _ContractTriggerGasRegistry.contract, event: "TGBalanceDeducted", logs: logs, sub: sub}, nil
}

// WatchTGBalanceDeducted is a free log subscription operation binding the contract event 0xef70e6d79855d76ed5d6ba11f8368a48ca19a9f11ffed71aa53caab2cf8c16a4.
//
// Solidity: event TGBalanceDeducted(address indexed user, uint256 amount)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) WatchTGBalanceDeducted(opts *bind.WatchOpts, sink chan<- *ContractTriggerGasRegistryTGBalanceDeducted, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ContractTriggerGasRegistry.contract.WatchLogs(opts, "TGBalanceDeducted", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerGasRegistryTGBalanceDeducted)
				if err := _ContractTriggerGasRegistry.contract.UnpackLog(event, "TGBalanceDeducted", log); err != nil {
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

// ParseTGBalanceDeducted is a log parse operation binding the contract event 0xef70e6d79855d76ed5d6ba11f8368a48ca19a9f11ffed71aa53caab2cf8c16a4.
//
// Solidity: event TGBalanceDeducted(address indexed user, uint256 amount)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) ParseTGBalanceDeducted(log types.Log) (*ContractTriggerGasRegistryTGBalanceDeducted, error) {
	event := new(ContractTriggerGasRegistryTGBalanceDeducted)
	if err := _ContractTriggerGasRegistry.contract.UnpackLog(event, "TGBalanceDeducted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTriggerGasRegistryTGBalanceRemovedIterator is returned from FilterTGBalanceRemoved and is used to iterate over the raw logs and unpacked data for TGBalanceRemoved events raised by the ContractTriggerGasRegistry contract.
type ContractTriggerGasRegistryTGBalanceRemovedIterator struct {
	Event *ContractTriggerGasRegistryTGBalanceRemoved // Event containing the contract specifics and raw log

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
func (it *ContractTriggerGasRegistryTGBalanceRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerGasRegistryTGBalanceRemoved)
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
		it.Event = new(ContractTriggerGasRegistryTGBalanceRemoved)
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
func (it *ContractTriggerGasRegistryTGBalanceRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerGasRegistryTGBalanceRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerGasRegistryTGBalanceRemoved represents a TGBalanceRemoved event raised by the ContractTriggerGasRegistry contract.
type ContractTriggerGasRegistryTGBalanceRemoved struct {
	User   common.Address
	Amount *big.Int
	Reason string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTGBalanceRemoved is a free log retrieval operation binding the contract event 0x5bb4e07063f2bc366c9a7d34efca7a39d411136e9a6af6a211676ab7d74fd8e0.
//
// Solidity: event TGBalanceRemoved(address indexed user, uint256 amount, string reason)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) FilterTGBalanceRemoved(opts *bind.FilterOpts, user []common.Address) (*ContractTriggerGasRegistryTGBalanceRemovedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ContractTriggerGasRegistry.contract.FilterLogs(opts, "TGBalanceRemoved", userRule)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerGasRegistryTGBalanceRemovedIterator{contract: _ContractTriggerGasRegistry.contract, event: "TGBalanceRemoved", logs: logs, sub: sub}, nil
}

// WatchTGBalanceRemoved is a free log subscription operation binding the contract event 0x5bb4e07063f2bc366c9a7d34efca7a39d411136e9a6af6a211676ab7d74fd8e0.
//
// Solidity: event TGBalanceRemoved(address indexed user, uint256 amount, string reason)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) WatchTGBalanceRemoved(opts *bind.WatchOpts, sink chan<- *ContractTriggerGasRegistryTGBalanceRemoved, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ContractTriggerGasRegistry.contract.WatchLogs(opts, "TGBalanceRemoved", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerGasRegistryTGBalanceRemoved)
				if err := _ContractTriggerGasRegistry.contract.UnpackLog(event, "TGBalanceRemoved", log); err != nil {
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

// ParseTGBalanceRemoved is a log parse operation binding the contract event 0x5bb4e07063f2bc366c9a7d34efca7a39d411136e9a6af6a211676ab7d74fd8e0.
//
// Solidity: event TGBalanceRemoved(address indexed user, uint256 amount, string reason)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) ParseTGBalanceRemoved(log types.Log) (*ContractTriggerGasRegistryTGBalanceRemoved, error) {
	event := new(ContractTriggerGasRegistryTGBalanceRemoved)
	if err := _ContractTriggerGasRegistry.contract.UnpackLog(event, "TGBalanceRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTriggerGasRegistryTGClaimedIterator is returned from FilterTGClaimed and is used to iterate over the raw logs and unpacked data for TGClaimed events raised by the ContractTriggerGasRegistry contract.
type ContractTriggerGasRegistryTGClaimedIterator struct {
	Event *ContractTriggerGasRegistryTGClaimed // Event containing the contract specifics and raw log

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
func (it *ContractTriggerGasRegistryTGClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerGasRegistryTGClaimed)
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
		it.Event = new(ContractTriggerGasRegistryTGClaimed)
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
func (it *ContractTriggerGasRegistryTGClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerGasRegistryTGClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerGasRegistryTGClaimed represents a TGClaimed event raised by the ContractTriggerGasRegistry contract.
type ContractTriggerGasRegistryTGClaimed struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTGClaimed is a free log retrieval operation binding the contract event 0xbddee7230372ecc0cb1780b66efe6d1461ee2a8e5b70239f3ec20a7d85f0e928.
//
// Solidity: event TGClaimed(address indexed user, uint256 amount)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) FilterTGClaimed(opts *bind.FilterOpts, user []common.Address) (*ContractTriggerGasRegistryTGClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ContractTriggerGasRegistry.contract.FilterLogs(opts, "TGClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerGasRegistryTGClaimedIterator{contract: _ContractTriggerGasRegistry.contract, event: "TGClaimed", logs: logs, sub: sub}, nil
}

// WatchTGClaimed is a free log subscription operation binding the contract event 0xbddee7230372ecc0cb1780b66efe6d1461ee2a8e5b70239f3ec20a7d85f0e928.
//
// Solidity: event TGClaimed(address indexed user, uint256 amount)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) WatchTGClaimed(opts *bind.WatchOpts, sink chan<- *ContractTriggerGasRegistryTGClaimed, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ContractTriggerGasRegistry.contract.WatchLogs(opts, "TGClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerGasRegistryTGClaimed)
				if err := _ContractTriggerGasRegistry.contract.UnpackLog(event, "TGClaimed", log); err != nil {
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
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) ParseTGClaimed(log types.Log) (*ContractTriggerGasRegistryTGClaimed, error) {
	event := new(ContractTriggerGasRegistryTGClaimed)
	if err := _ContractTriggerGasRegistry.contract.UnpackLog(event, "TGClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTriggerGasRegistryTGPerETHUpdatedIterator is returned from FilterTGPerETHUpdated and is used to iterate over the raw logs and unpacked data for TGPerETHUpdated events raised by the ContractTriggerGasRegistry contract.
type ContractTriggerGasRegistryTGPerETHUpdatedIterator struct {
	Event *ContractTriggerGasRegistryTGPerETHUpdated // Event containing the contract specifics and raw log

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
func (it *ContractTriggerGasRegistryTGPerETHUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerGasRegistryTGPerETHUpdated)
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
		it.Event = new(ContractTriggerGasRegistryTGPerETHUpdated)
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
func (it *ContractTriggerGasRegistryTGPerETHUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerGasRegistryTGPerETHUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerGasRegistryTGPerETHUpdated represents a TGPerETHUpdated event raised by the ContractTriggerGasRegistry contract.
type ContractTriggerGasRegistryTGPerETHUpdated struct {
	TgPerEth *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTGPerETHUpdated is a free log retrieval operation binding the contract event 0x18652b6a656730019776361074bff5dbd2c837755a539e612061263f50b83457.
//
// Solidity: event TGPerETHUpdated(uint256 tgPerEth)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) FilterTGPerETHUpdated(opts *bind.FilterOpts) (*ContractTriggerGasRegistryTGPerETHUpdatedIterator, error) {

	logs, sub, err := _ContractTriggerGasRegistry.contract.FilterLogs(opts, "TGPerETHUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractTriggerGasRegistryTGPerETHUpdatedIterator{contract: _ContractTriggerGasRegistry.contract, event: "TGPerETHUpdated", logs: logs, sub: sub}, nil
}

// WatchTGPerETHUpdated is a free log subscription operation binding the contract event 0x18652b6a656730019776361074bff5dbd2c837755a539e612061263f50b83457.
//
// Solidity: event TGPerETHUpdated(uint256 tgPerEth)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) WatchTGPerETHUpdated(opts *bind.WatchOpts, sink chan<- *ContractTriggerGasRegistryTGPerETHUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractTriggerGasRegistry.contract.WatchLogs(opts, "TGPerETHUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerGasRegistryTGPerETHUpdated)
				if err := _ContractTriggerGasRegistry.contract.UnpackLog(event, "TGPerETHUpdated", log); err != nil {
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

// ParseTGPerETHUpdated is a log parse operation binding the contract event 0x18652b6a656730019776361074bff5dbd2c837755a539e612061263f50b83457.
//
// Solidity: event TGPerETHUpdated(uint256 tgPerEth)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) ParseTGPerETHUpdated(log types.Log) (*ContractTriggerGasRegistryTGPerETHUpdated, error) {
	event := new(ContractTriggerGasRegistryTGPerETHUpdated)
	if err := _ContractTriggerGasRegistry.contract.UnpackLog(event, "TGPerETHUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTriggerGasRegistryTGPurchasedIterator is returned from FilterTGPurchased and is used to iterate over the raw logs and unpacked data for TGPurchased events raised by the ContractTriggerGasRegistry contract.
type ContractTriggerGasRegistryTGPurchasedIterator struct {
	Event *ContractTriggerGasRegistryTGPurchased // Event containing the contract specifics and raw log

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
func (it *ContractTriggerGasRegistryTGPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerGasRegistryTGPurchased)
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
		it.Event = new(ContractTriggerGasRegistryTGPurchased)
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
func (it *ContractTriggerGasRegistryTGPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerGasRegistryTGPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerGasRegistryTGPurchased represents a TGPurchased event raised by the ContractTriggerGasRegistry contract.
type ContractTriggerGasRegistryTGPurchased struct {
	User      common.Address
	EthAmount *big.Int
	TgAmount  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTGPurchased is a free log retrieval operation binding the contract event 0xd29e7684851e87b600fbe9f3d04a3adda696b6768e9bde3c3ae69143b4efa2b5.
//
// Solidity: event TGPurchased(address indexed user, uint256 ethAmount, uint256 tgAmount)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) FilterTGPurchased(opts *bind.FilterOpts, user []common.Address) (*ContractTriggerGasRegistryTGPurchasedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ContractTriggerGasRegistry.contract.FilterLogs(opts, "TGPurchased", userRule)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerGasRegistryTGPurchasedIterator{contract: _ContractTriggerGasRegistry.contract, event: "TGPurchased", logs: logs, sub: sub}, nil
}

// WatchTGPurchased is a free log subscription operation binding the contract event 0xd29e7684851e87b600fbe9f3d04a3adda696b6768e9bde3c3ae69143b4efa2b5.
//
// Solidity: event TGPurchased(address indexed user, uint256 ethAmount, uint256 tgAmount)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) WatchTGPurchased(opts *bind.WatchOpts, sink chan<- *ContractTriggerGasRegistryTGPurchased, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ContractTriggerGasRegistry.contract.WatchLogs(opts, "TGPurchased", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerGasRegistryTGPurchased)
				if err := _ContractTriggerGasRegistry.contract.UnpackLog(event, "TGPurchased", log); err != nil {
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

// ParseTGPurchased is a log parse operation binding the contract event 0xd29e7684851e87b600fbe9f3d04a3adda696b6768e9bde3c3ae69143b4efa2b5.
//
// Solidity: event TGPurchased(address indexed user, uint256 ethAmount, uint256 tgAmount)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) ParseTGPurchased(log types.Log) (*ContractTriggerGasRegistryTGPurchased, error) {
	event := new(ContractTriggerGasRegistryTGPurchased)
	if err := _ContractTriggerGasRegistry.contract.UnpackLog(event, "TGPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTriggerGasRegistryTGRefundedIterator is returned from FilterTGRefunded and is used to iterate over the raw logs and unpacked data for TGRefunded events raised by the ContractTriggerGasRegistry contract.
type ContractTriggerGasRegistryTGRefundedIterator struct {
	Event *ContractTriggerGasRegistryTGRefunded // Event containing the contract specifics and raw log

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
func (it *ContractTriggerGasRegistryTGRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerGasRegistryTGRefunded)
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
		it.Event = new(ContractTriggerGasRegistryTGRefunded)
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
func (it *ContractTriggerGasRegistryTGRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerGasRegistryTGRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerGasRegistryTGRefunded represents a TGRefunded event raised by the ContractTriggerGasRegistry contract.
type ContractTriggerGasRegistryTGRefunded struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTGRefunded is a free log retrieval operation binding the contract event 0x0ec18e13bc6adfe1239be0ea7ec72bd748b32f99a9fa97d32268b5a04a73c62f.
//
// Solidity: event TGRefunded(address indexed user, uint256 amount)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) FilterTGRefunded(opts *bind.FilterOpts, user []common.Address) (*ContractTriggerGasRegistryTGRefundedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ContractTriggerGasRegistry.contract.FilterLogs(opts, "TGRefunded", userRule)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerGasRegistryTGRefundedIterator{contract: _ContractTriggerGasRegistry.contract, event: "TGRefunded", logs: logs, sub: sub}, nil
}

// WatchTGRefunded is a free log subscription operation binding the contract event 0x0ec18e13bc6adfe1239be0ea7ec72bd748b32f99a9fa97d32268b5a04a73c62f.
//
// Solidity: event TGRefunded(address indexed user, uint256 amount)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) WatchTGRefunded(opts *bind.WatchOpts, sink chan<- *ContractTriggerGasRegistryTGRefunded, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ContractTriggerGasRegistry.contract.WatchLogs(opts, "TGRefunded", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerGasRegistryTGRefunded)
				if err := _ContractTriggerGasRegistry.contract.UnpackLog(event, "TGRefunded", log); err != nil {
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

// ParseTGRefunded is a log parse operation binding the contract event 0x0ec18e13bc6adfe1239be0ea7ec72bd748b32f99a9fa97d32268b5a04a73c62f.
//
// Solidity: event TGRefunded(address indexed user, uint256 amount)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) ParseTGRefunded(log types.Log) (*ContractTriggerGasRegistryTGRefunded, error) {
	event := new(ContractTriggerGasRegistryTGRefunded)
	if err := _ContractTriggerGasRegistry.contract.UnpackLog(event, "TGRefunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTriggerGasRegistryTGTransferredIterator is returned from FilterTGTransferred and is used to iterate over the raw logs and unpacked data for TGTransferred events raised by the ContractTriggerGasRegistry contract.
type ContractTriggerGasRegistryTGTransferredIterator struct {
	Event *ContractTriggerGasRegistryTGTransferred // Event containing the contract specifics and raw log

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
func (it *ContractTriggerGasRegistryTGTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerGasRegistryTGTransferred)
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
		it.Event = new(ContractTriggerGasRegistryTGTransferred)
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
func (it *ContractTriggerGasRegistryTGTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerGasRegistryTGTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerGasRegistryTGTransferred represents a TGTransferred event raised by the ContractTriggerGasRegistry contract.
type ContractTriggerGasRegistryTGTransferred struct {
	User   common.Address
	Keeper common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTGTransferred is a free log retrieval operation binding the contract event 0xa613b7e352cf78a30d92dbc93071265034f43edbeb7833b8d8713a3b4f323670.
//
// Solidity: event TGTransferred(address indexed user, address indexed keeper, uint256 amount)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) FilterTGTransferred(opts *bind.FilterOpts, user []common.Address, keeper []common.Address) (*ContractTriggerGasRegistryTGTransferredIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}

	logs, sub, err := _ContractTriggerGasRegistry.contract.FilterLogs(opts, "TGTransferred", userRule, keeperRule)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerGasRegistryTGTransferredIterator{contract: _ContractTriggerGasRegistry.contract, event: "TGTransferred", logs: logs, sub: sub}, nil
}

// WatchTGTransferred is a free log subscription operation binding the contract event 0xa613b7e352cf78a30d92dbc93071265034f43edbeb7833b8d8713a3b4f323670.
//
// Solidity: event TGTransferred(address indexed user, address indexed keeper, uint256 amount)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) WatchTGTransferred(opts *bind.WatchOpts, sink chan<- *ContractTriggerGasRegistryTGTransferred, user []common.Address, keeper []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}

	logs, sub, err := _ContractTriggerGasRegistry.contract.WatchLogs(opts, "TGTransferred", userRule, keeperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerGasRegistryTGTransferred)
				if err := _ContractTriggerGasRegistry.contract.UnpackLog(event, "TGTransferred", log); err != nil {
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
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) ParseTGTransferred(log types.Log) (*ContractTriggerGasRegistryTGTransferred, error) {
	event := new(ContractTriggerGasRegistryTGTransferred)
	if err := _ContractTriggerGasRegistry.contract.UnpackLog(event, "TGTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTriggerGasRegistryTaskFeeClaimedIterator is returned from FilterTaskFeeClaimed and is used to iterate over the raw logs and unpacked data for TaskFeeClaimed events raised by the ContractTriggerGasRegistry contract.
type ContractTriggerGasRegistryTaskFeeClaimedIterator struct {
	Event *ContractTriggerGasRegistryTaskFeeClaimed // Event containing the contract specifics and raw log

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
func (it *ContractTriggerGasRegistryTaskFeeClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerGasRegistryTaskFeeClaimed)
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
		it.Event = new(ContractTriggerGasRegistryTaskFeeClaimed)
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
func (it *ContractTriggerGasRegistryTaskFeeClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerGasRegistryTaskFeeClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerGasRegistryTaskFeeClaimed represents a TaskFeeClaimed event raised by the ContractTriggerGasRegistry contract.
type ContractTriggerGasRegistryTaskFeeClaimed struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTaskFeeClaimed is a free log retrieval operation binding the contract event 0xae8ecfb3533b83d346a9f325b0e61ccdd5a8d8aaa0fa2dbfcbea06121120da40.
//
// Solidity: event TaskFeeClaimed(address indexed user, uint256 amount)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) FilterTaskFeeClaimed(opts *bind.FilterOpts, user []common.Address) (*ContractTriggerGasRegistryTaskFeeClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ContractTriggerGasRegistry.contract.FilterLogs(opts, "TaskFeeClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerGasRegistryTaskFeeClaimedIterator{contract: _ContractTriggerGasRegistry.contract, event: "TaskFeeClaimed", logs: logs, sub: sub}, nil
}

// WatchTaskFeeClaimed is a free log subscription operation binding the contract event 0xae8ecfb3533b83d346a9f325b0e61ccdd5a8d8aaa0fa2dbfcbea06121120da40.
//
// Solidity: event TaskFeeClaimed(address indexed user, uint256 amount)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) WatchTaskFeeClaimed(opts *bind.WatchOpts, sink chan<- *ContractTriggerGasRegistryTaskFeeClaimed, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ContractTriggerGasRegistry.contract.WatchLogs(opts, "TaskFeeClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerGasRegistryTaskFeeClaimed)
				if err := _ContractTriggerGasRegistry.contract.UnpackLog(event, "TaskFeeClaimed", log); err != nil {
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
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) ParseTaskFeeClaimed(log types.Log) (*ContractTriggerGasRegistryTaskFeeClaimed, error) {
	event := new(ContractTriggerGasRegistryTaskFeeClaimed)
	if err := _ContractTriggerGasRegistry.contract.UnpackLog(event, "TaskFeeClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTriggerGasRegistryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ContractTriggerGasRegistry contract.
type ContractTriggerGasRegistryUpgradedIterator struct {
	Event *ContractTriggerGasRegistryUpgraded // Event containing the contract specifics and raw log

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
func (it *ContractTriggerGasRegistryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerGasRegistryUpgraded)
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
		it.Event = new(ContractTriggerGasRegistryUpgraded)
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
func (it *ContractTriggerGasRegistryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerGasRegistryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerGasRegistryUpgraded represents a Upgraded event raised by the ContractTriggerGasRegistry contract.
type ContractTriggerGasRegistryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ContractTriggerGasRegistryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ContractTriggerGasRegistry.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerGasRegistryUpgradedIterator{contract: _ContractTriggerGasRegistry.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ContractTriggerGasRegistryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ContractTriggerGasRegistry.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerGasRegistryUpgraded)
				if err := _ContractTriggerGasRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) ParseUpgraded(log types.Log) (*ContractTriggerGasRegistryUpgraded, error) {
	event := new(ContractTriggerGasRegistryUpgraded)
	if err := _ContractTriggerGasRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
