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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"balances\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"batchMigrateUsers\",\"inputs\":[{\"name\":\"users\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"ethAmounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"deductETHBalance\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"ethAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositETH\",\"inputs\":[{\"name\":\"ethAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getBalance\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"ethAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operatorRole\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOperator\",\"inputs\":[{\"name\":\"_operatorRole\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalDeductedBalance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdrawETH\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reason\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawETHBalance\",\"inputs\":[{\"name\":\"ethAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ETHBalanceDeducted\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ETHDeposited\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"ethAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ETHWithdrawn\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"reason\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	Bin: "0x60a080604052346100c257306080525f5160206112585f395f51905f525460ff8160401c166100b3576002600160401b03196001600160401b03821601610060575b60405161119190816100c7823960805181818161084901526108ec0152f35b6001600160401b0319166001600160401b039081175f5160206112585f395f51905f525581527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80610041565b63f92ee8a960e01b5f5260045ffd5b5f80fdfe60806040526004361015610011575f80fd5b5f3560e01c8063176c510614610ca157806327e235e3146100d457806347503b3d14610c79578063485cc95514610a7c5780634f1ef2861461089d57806352d1902d146108375780635358fbda1461073b578063715018a6146106d45780638da5cb5b146106a0578063a3c4c129146104e4578063ad3cb1cc14610499578063b3ab15fb14610408578063cc11362214610252578063f1cc931014610235578063f2fde38b1461020a578063f7c786041461010f5763f8b2cb4f146100d4575f80fd5b3461010b57602036600319011261010b576001600160a01b036100f5610d1c565b165f525f602052602060405f2054604051908152f35b5f80fd5b3461010b57604036600319011261010b57610128610d1c565b60015460243591906001600160a01b031633036101b95760018060a01b031690815f525f60205261015f8160405f20541015610e53565b8061016657005b60207fcb86089062e14afd83ae7cdae0be8f14e3cc48bcb2dd2b09bf2e56c2fdee8fee91835f525f825260405f2061019f828254610e9f565b90556101ad81600254610f31565b600255604051908152a2005b60405162461bcd60e51b8152602060048201526024808201527f4f6e6c79206f70657261746f722063616e2063616c6c20746869732066756e636044820152633a34b7b760e11b6064820152608490fd5b3461010b57602036600319011261010b57610233610226610d1c565b61022e61101f565b610f76565b005b3461010b575f36600319011261010b576020600254604051908152f35b3461010b57604036600319011261010b5760043560243567ffffffffffffffff811161010b573660238201121561010b57610297903690602481600401359101610d84565b61029f61101f565b6102a7610fe7565b6102b2821515610e0f565b600254908183116103b05747831161036b576102ef837f68fc7e0340ad067ed9d1bcd121870cf25bd80b18a94d133b8bec22f4ea86b48593610e9f565b6002556103215f8080808760018060a01b035f5160206110dc5f395f51905f5254165af161031b610ec0565b50610eef565b60018060a01b035f5160206110dc5f395f51905f525416926103556040519283928352604060208401526040830190610deb565b0390a260015f51602061111c5f395f51905f5255005b60405162461bcd60e51b815260206004820152601d60248201527f496e73756666696369656e7420636f6e74726163742062616c616e63650000006044820152606490fd5b60405162461bcd60e51b815260206004820152602a60248201527f43616e6e6f74207769746864726177206d6f7265207468616e2064656475637460448201526965642062616c616e636560b01b6064820152608490fd5b3461010b57602036600319011261010b57610421610d1c565b61042961101f565b6001600160a01b03168015610454576bffffffffffffffffffffffff60a01b60015416176001555f80f35b60405162461bcd60e51b815260206004820152601c60248201527f4f70657261746f722063616e6e6f7420626520302061646472657373000000006044820152606490fd5b3461010b575f36600319011261010b576104e06040516104ba604082610d32565b60058152640352e302e360dc1b6020820152604051918291602083526020830190610deb565b0390f35b604036600319011261010b5760043567ffffffffffffffff811161010b57610510903690600401610dba565b9060243567ffffffffffffffff811161010b57610531903690600401610dba565b61053961101f565b808403610669579291905f935f935b8085106105c35785341061055857005b60405162461bcd60e51b815260206004820152603a60248201527f53656e7420455448206d7573742062652067726561746572207468616e206f7260448201527f20657175616c20746f20746f74616c2045544820616d6f756e740000000000006064820152608490fd5b90919293946106116001916105d9888688610f3e565b35838060a01b036105f36105ee8b888c610f3e565b610f62565b165f525f60205260405f205561060a888688610f3e565b3590610f31565b956106206105ee828589610f3e565b7f6c703791f399558807424f489ccd811c72b4ff0b74af547264fad7c646776df0602061064e84888a610f3e565b3592604051938452858060a01b031692a20193929190610548565b60405162461bcd60e51b815260206004820152600f60248201526e098cadccee8d040dad2e6dac2e8c6d608b1b6044820152606490fd5b3461010b575f36600319011261010b575f5160206110dc5f395f51905f52546040516001600160a01b039091168152602090f35b3461010b575f36600319011261010b576106ec61101f565b5f5160206110dc5f395f51905f5280546001600160a01b031981169091555f906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b602036600319011261010b57600435610752610fe7565b80156107fb578034036107b657335f525f60205260405f20610775828254610f31565b90556040519081527f6c703791f399558807424f489ccd811c72b4ff0b74af547264fad7c646776df060203392a260015f51602061111c5f395f51905f5255005b60405162461bcd60e51b815260206004820152601a60248201527f53656e7420455448206d757374206d6174636820616d6f756e740000000000006044820152606490fd5b60405162461bcd60e51b8152602060048201526014602482015273086c2dcdcdee840c8cae0dee6d2e84060408aa8960631b6044820152606490fd5b3461010b575f36600319011261010b577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316300361088e5760206040515f5160206110fc5f395f51905f528152f35b63703e46dd60e11b5f5260045ffd5b604036600319011261010b576108b1610d1c565b60243567ffffffffffffffff811161010b573660238201121561010b576108e2903690602481600401359101610d84565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016308114908115610a5a575b5061088e5761092461101f565b6040516352d1902d60e01b81526001600160a01b0383169290602081600481875afa5f9181610a26575b506109665783634c9c8ce360e01b5f5260045260245ffd5b805f5160206110fc5f395f51905f52859203610a145750813b15610a02575f5160206110fc5f395f51905f5280546001600160a01b031916821790557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b5f80a28151156109ea575f8083602061023395519101845af46109e4610ec0565b9161107d565b5050346109f357005b63b398979f60e01b5f5260045ffd5b634c9c8ce360e01b5f5260045260245ffd5b632a87526960e21b5f5260045260245ffd5b9091506020813d602011610a52575b81610a4260209383610d32565b8101031261010b5751908561094e565b3d9150610a35565b5f5160206110fc5f395f51905f52546001600160a01b03161415905083610917565b3461010b57604036600319011261010b57610a95610d1c565b6024356001600160a01b0381169081900361010b575f51602061113c5f395f51905f52549160ff8360401c16159267ffffffffffffffff811680159081610c71575b6001149081610c67575b159081610c5e575b50610c4f5767ffffffffffffffff1981166001175f51602061113c5f395f51905f525583610c23575b506001600160a01b03811615610bd457610b5790610b2e611052565b610b36611052565b60015f51602061111c5f395f51905f5255610b4f611052565b61022e611052565b610b5f611052565b6bffffffffffffffffffffffff60a01b6001541617600155610b7d57005b68ff0000000000000000195f51602061113c5f395f51905f5254165f51602061113c5f395f51905f52557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a1005b60405162461bcd60e51b815260206004820152602160248201527f496e697469616c206f776e65722063616e6e6f742062652030206164647265736044820152607360f81b6064820152608490fd5b68ffffffffffffffffff191668010000000000000001175f51602061113c5f395f51905f525583610b12565b63f92ee8a960e01b5f5260045ffd5b90501585610ae9565b303b159150610ae1565b859150610ad7565b3461010b575f36600319011261010b576001546040516001600160a01b039091168152602090f35b3461010b57602036600319011261010b57610d095f808080600435610cc4610fe7565b610ccf811515610e0f565b33825281602052610ce68160408420541015610e53565b3382528160205260408220610cfc828254610e9f565b9055335af161031b610ec0565b60015f51602061111c5f395f51905f5255005b600435906001600160a01b038216820361010b57565b90601f8019910116810190811067ffffffffffffffff821117610d5457604052565b634e487b7160e01b5f52604160045260245ffd5b67ffffffffffffffff8111610d5457601f01601f191660200190565b929192610d9082610d68565b91610d9e6040519384610d32565b82948184528183011161010b578281602093845f960137010152565b9181601f8401121561010b5782359167ffffffffffffffff831161010b576020808501948460051b01011161010b57565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b15610e1657565b60405162461bcd60e51b8152602060048201526015602482015274086c2dcdcdee840eed2e8d0c8e4c2ee4060408aa89605b1b6044820152606490fd5b15610e5a57565b60405162461bcd60e51b815260206004820152601860248201527f496e73756666696369656e74204554482062616c616e636500000000000000006044820152606490fd5b91908203918211610eac57565b634e487b7160e01b5f52601160045260245ffd5b3d15610eea573d90610ed182610d68565b91610edf6040519384610d32565b82523d5f602084013e565b606090565b15610ef657565b60405162461bcd60e51b8152602060048201526013602482015272115512081d1c985b9cd9995c8819985a5b1959606a1b6044820152606490fd5b91908201809211610eac57565b9190811015610f4e5760051b0190565b634e487b7160e01b5f52603260045260245ffd5b356001600160a01b038116810361010b5790565b6001600160a01b03168015610fd4575f5160206110dc5f395f51905f5280546001600160a01b0319811683179091556001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3565b631e4fbdf760e01b5f525f60045260245ffd5b60025f51602061111c5f395f51905f5254146110105760025f51602061111c5f395f51905f5255565b633ee5aeb560e01b5f5260045ffd5b5f5160206110dc5f395f51905f52546001600160a01b0316330361103f57565b63118cdaa760e01b5f523360045260245ffd5b60ff5f51602061113c5f395f51905f525460401c161561106e57565b631afcd79f60e31b5f5260045ffd5b906110a1575080511561109257805190602001fd5b63d6bda27560e01b5f5260045ffd5b815115806110d2575b6110b2575090565b639996b31560e01b5f9081526001600160a01b0391909116600452602490fd5b50803b156110aa56fe9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a2646970667358221220c78ee5f63e37a6855c29bf6d6cc19e2ff83c3de8be6fcc26559a14c4c745cedb64736f6c634300081b0033f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00",
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
	UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error)

	Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error)

	GetBalance(opts *bind.CallOpts, user common.Address) (*big.Int, error)

	OperatorRole(opts *bind.CallOpts) (common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	ProxiableUUID(opts *bind.CallOpts) ([32]byte, error)

	TotalDeductedBalance(opts *bind.CallOpts) (*big.Int, error)
}

// ContractTriggerGasRegistryTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractTriggerGasRegistryTransacts interface {
	BatchMigrateUsers(opts *bind.TransactOpts, users []common.Address, ethAmounts []*big.Int) (*types.Transaction, error)

	DeductETHBalance(opts *bind.TransactOpts, user common.Address, ethAmount *big.Int) (*types.Transaction, error)

	DepositETH(opts *bind.TransactOpts, ethAmount *big.Int) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts, initialOwner common.Address, _operator common.Address) (*types.Transaction, error)

	RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetOperator(opts *bind.TransactOpts, _operatorRole common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error)

	UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error)

	WithdrawETH(opts *bind.TransactOpts, amount *big.Int, reason string) (*types.Transaction, error)

	WithdrawETHBalance(opts *bind.TransactOpts, ethAmount *big.Int) (*types.Transaction, error)
}

// ContractTriggerGasRegistryFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractTriggerGasRegistryFilters interface {
	FilterETHBalanceDeducted(opts *bind.FilterOpts, user []common.Address) (*ContractTriggerGasRegistryETHBalanceDeductedIterator, error)
	WatchETHBalanceDeducted(opts *bind.WatchOpts, sink chan<- *ContractTriggerGasRegistryETHBalanceDeducted, user []common.Address) (event.Subscription, error)
	ParseETHBalanceDeducted(log types.Log) (*ContractTriggerGasRegistryETHBalanceDeducted, error)

	FilterETHDeposited(opts *bind.FilterOpts, user []common.Address) (*ContractTriggerGasRegistryETHDepositedIterator, error)
	WatchETHDeposited(opts *bind.WatchOpts, sink chan<- *ContractTriggerGasRegistryETHDeposited, user []common.Address) (event.Subscription, error)
	ParseETHDeposited(log types.Log) (*ContractTriggerGasRegistryETHDeposited, error)

	FilterETHWithdrawn(opts *bind.FilterOpts, owner []common.Address) (*ContractTriggerGasRegistryETHWithdrawnIterator, error)
	WatchETHWithdrawn(opts *bind.WatchOpts, sink chan<- *ContractTriggerGasRegistryETHWithdrawn, owner []common.Address) (event.Subscription, error)
	ParseETHWithdrawn(log types.Log) (*ContractTriggerGasRegistryETHWithdrawn, error)

	FilterInitialized(opts *bind.FilterOpts) (*ContractTriggerGasRegistryInitializedIterator, error)
	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractTriggerGasRegistryInitialized) (event.Subscription, error)
	ParseInitialized(log types.Log) (*ContractTriggerGasRegistryInitialized, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractTriggerGasRegistryOwnershipTransferredIterator, error)
	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractTriggerGasRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error)
	ParseOwnershipTransferred(log types.Log) (*ContractTriggerGasRegistryOwnershipTransferred, error)

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
// Solidity: function balances(address ) view returns(uint256)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractTriggerGasRegistry.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistrySession) Balances(arg0 common.Address) (*big.Int, error) {
	return _ContractTriggerGasRegistry.Contract.Balances(&_ContractTriggerGasRegistry.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _ContractTriggerGasRegistry.Contract.Balances(&_ContractTriggerGasRegistry.CallOpts, arg0)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address user) view returns(uint256 ethAmount)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryCaller) GetBalance(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractTriggerGasRegistry.contract.Call(opts, &out, "getBalance", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address user) view returns(uint256 ethAmount)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistrySession) GetBalance(user common.Address) (*big.Int, error) {
	return _ContractTriggerGasRegistry.Contract.GetBalance(&_ContractTriggerGasRegistry.CallOpts, user)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address user) view returns(uint256 ethAmount)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryCallerSession) GetBalance(user common.Address) (*big.Int, error) {
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

// TotalDeductedBalance is a free data retrieval call binding the contract method 0xf1cc9310.
//
// Solidity: function totalDeductedBalance() view returns(uint256)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryCaller) TotalDeductedBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractTriggerGasRegistry.contract.Call(opts, &out, "totalDeductedBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDeductedBalance is a free data retrieval call binding the contract method 0xf1cc9310.
//
// Solidity: function totalDeductedBalance() view returns(uint256)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistrySession) TotalDeductedBalance() (*big.Int, error) {
	return _ContractTriggerGasRegistry.Contract.TotalDeductedBalance(&_ContractTriggerGasRegistry.CallOpts)
}

// TotalDeductedBalance is a free data retrieval call binding the contract method 0xf1cc9310.
//
// Solidity: function totalDeductedBalance() view returns(uint256)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryCallerSession) TotalDeductedBalance() (*big.Int, error) {
	return _ContractTriggerGasRegistry.Contract.TotalDeductedBalance(&_ContractTriggerGasRegistry.CallOpts)
}

// BatchMigrateUsers is a paid mutator transaction binding the contract method 0xa3c4c129.
//
// Solidity: function batchMigrateUsers(address[] users, uint256[] ethAmounts) payable returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryTransactor) BatchMigrateUsers(opts *bind.TransactOpts, users []common.Address, ethAmounts []*big.Int) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.contract.Transact(opts, "batchMigrateUsers", users, ethAmounts)
}

// BatchMigrateUsers is a paid mutator transaction binding the contract method 0xa3c4c129.
//
// Solidity: function batchMigrateUsers(address[] users, uint256[] ethAmounts) payable returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistrySession) BatchMigrateUsers(users []common.Address, ethAmounts []*big.Int) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.Contract.BatchMigrateUsers(&_ContractTriggerGasRegistry.TransactOpts, users, ethAmounts)
}

// BatchMigrateUsers is a paid mutator transaction binding the contract method 0xa3c4c129.
//
// Solidity: function batchMigrateUsers(address[] users, uint256[] ethAmounts) payable returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryTransactorSession) BatchMigrateUsers(users []common.Address, ethAmounts []*big.Int) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.Contract.BatchMigrateUsers(&_ContractTriggerGasRegistry.TransactOpts, users, ethAmounts)
}

// DeductETHBalance is a paid mutator transaction binding the contract method 0xf7c78604.
//
// Solidity: function deductETHBalance(address user, uint256 ethAmount) returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryTransactor) DeductETHBalance(opts *bind.TransactOpts, user common.Address, ethAmount *big.Int) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.contract.Transact(opts, "deductETHBalance", user, ethAmount)
}

// DeductETHBalance is a paid mutator transaction binding the contract method 0xf7c78604.
//
// Solidity: function deductETHBalance(address user, uint256 ethAmount) returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistrySession) DeductETHBalance(user common.Address, ethAmount *big.Int) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.Contract.DeductETHBalance(&_ContractTriggerGasRegistry.TransactOpts, user, ethAmount)
}

// DeductETHBalance is a paid mutator transaction binding the contract method 0xf7c78604.
//
// Solidity: function deductETHBalance(address user, uint256 ethAmount) returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryTransactorSession) DeductETHBalance(user common.Address, ethAmount *big.Int) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.Contract.DeductETHBalance(&_ContractTriggerGasRegistry.TransactOpts, user, ethAmount)
}

// DepositETH is a paid mutator transaction binding the contract method 0x5358fbda.
//
// Solidity: function depositETH(uint256 ethAmount) payable returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryTransactor) DepositETH(opts *bind.TransactOpts, ethAmount *big.Int) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.contract.Transact(opts, "depositETH", ethAmount)
}

// DepositETH is a paid mutator transaction binding the contract method 0x5358fbda.
//
// Solidity: function depositETH(uint256 ethAmount) payable returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistrySession) DepositETH(ethAmount *big.Int) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.Contract.DepositETH(&_ContractTriggerGasRegistry.TransactOpts, ethAmount)
}

// DepositETH is a paid mutator transaction binding the contract method 0x5358fbda.
//
// Solidity: function depositETH(uint256 ethAmount) payable returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryTransactorSession) DepositETH(ethAmount *big.Int) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.Contract.DepositETH(&_ContractTriggerGasRegistry.TransactOpts, ethAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address initialOwner, address _operator) returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, _operator common.Address) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.contract.Transact(opts, "initialize", initialOwner, _operator)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address initialOwner, address _operator) returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistrySession) Initialize(initialOwner common.Address, _operator common.Address) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.Contract.Initialize(&_ContractTriggerGasRegistry.TransactOpts, initialOwner, _operator)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address initialOwner, address _operator) returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryTransactorSession) Initialize(initialOwner common.Address, _operator common.Address) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.Contract.Initialize(&_ContractTriggerGasRegistry.TransactOpts, initialOwner, _operator)
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

// WithdrawETHBalance is a paid mutator transaction binding the contract method 0x176c5106.
//
// Solidity: function withdrawETHBalance(uint256 ethAmount) returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryTransactor) WithdrawETHBalance(opts *bind.TransactOpts, ethAmount *big.Int) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.contract.Transact(opts, "withdrawETHBalance", ethAmount)
}

// WithdrawETHBalance is a paid mutator transaction binding the contract method 0x176c5106.
//
// Solidity: function withdrawETHBalance(uint256 ethAmount) returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistrySession) WithdrawETHBalance(ethAmount *big.Int) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.Contract.WithdrawETHBalance(&_ContractTriggerGasRegistry.TransactOpts, ethAmount)
}

// WithdrawETHBalance is a paid mutator transaction binding the contract method 0x176c5106.
//
// Solidity: function withdrawETHBalance(uint256 ethAmount) returns()
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryTransactorSession) WithdrawETHBalance(ethAmount *big.Int) (*types.Transaction, error) {
	return _ContractTriggerGasRegistry.Contract.WithdrawETHBalance(&_ContractTriggerGasRegistry.TransactOpts, ethAmount)
}

// ContractTriggerGasRegistryETHBalanceDeductedIterator is returned from FilterETHBalanceDeducted and is used to iterate over the raw logs and unpacked data for ETHBalanceDeducted events raised by the ContractTriggerGasRegistry contract.
type ContractTriggerGasRegistryETHBalanceDeductedIterator struct {
	Event *ContractTriggerGasRegistryETHBalanceDeducted // Event containing the contract specifics and raw log

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
func (it *ContractTriggerGasRegistryETHBalanceDeductedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerGasRegistryETHBalanceDeducted)
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
		it.Event = new(ContractTriggerGasRegistryETHBalanceDeducted)
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
func (it *ContractTriggerGasRegistryETHBalanceDeductedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerGasRegistryETHBalanceDeductedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerGasRegistryETHBalanceDeducted represents a ETHBalanceDeducted event raised by the ContractTriggerGasRegistry contract.
type ContractTriggerGasRegistryETHBalanceDeducted struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterETHBalanceDeducted is a free log retrieval operation binding the contract event 0xcb86089062e14afd83ae7cdae0be8f14e3cc48bcb2dd2b09bf2e56c2fdee8fee.
//
// Solidity: event ETHBalanceDeducted(address indexed user, uint256 amount)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) FilterETHBalanceDeducted(opts *bind.FilterOpts, user []common.Address) (*ContractTriggerGasRegistryETHBalanceDeductedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ContractTriggerGasRegistry.contract.FilterLogs(opts, "ETHBalanceDeducted", userRule)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerGasRegistryETHBalanceDeductedIterator{contract: _ContractTriggerGasRegistry.contract, event: "ETHBalanceDeducted", logs: logs, sub: sub}, nil
}

// WatchETHBalanceDeducted is a free log subscription operation binding the contract event 0xcb86089062e14afd83ae7cdae0be8f14e3cc48bcb2dd2b09bf2e56c2fdee8fee.
//
// Solidity: event ETHBalanceDeducted(address indexed user, uint256 amount)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) WatchETHBalanceDeducted(opts *bind.WatchOpts, sink chan<- *ContractTriggerGasRegistryETHBalanceDeducted, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ContractTriggerGasRegistry.contract.WatchLogs(opts, "ETHBalanceDeducted", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerGasRegistryETHBalanceDeducted)
				if err := _ContractTriggerGasRegistry.contract.UnpackLog(event, "ETHBalanceDeducted", log); err != nil {
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

// ParseETHBalanceDeducted is a log parse operation binding the contract event 0xcb86089062e14afd83ae7cdae0be8f14e3cc48bcb2dd2b09bf2e56c2fdee8fee.
//
// Solidity: event ETHBalanceDeducted(address indexed user, uint256 amount)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) ParseETHBalanceDeducted(log types.Log) (*ContractTriggerGasRegistryETHBalanceDeducted, error) {
	event := new(ContractTriggerGasRegistryETHBalanceDeducted)
	if err := _ContractTriggerGasRegistry.contract.UnpackLog(event, "ETHBalanceDeducted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTriggerGasRegistryETHDepositedIterator is returned from FilterETHDeposited and is used to iterate over the raw logs and unpacked data for ETHDeposited events raised by the ContractTriggerGasRegistry contract.
type ContractTriggerGasRegistryETHDepositedIterator struct {
	Event *ContractTriggerGasRegistryETHDeposited // Event containing the contract specifics and raw log

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
func (it *ContractTriggerGasRegistryETHDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerGasRegistryETHDeposited)
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
		it.Event = new(ContractTriggerGasRegistryETHDeposited)
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
func (it *ContractTriggerGasRegistryETHDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerGasRegistryETHDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerGasRegistryETHDeposited represents a ETHDeposited event raised by the ContractTriggerGasRegistry contract.
type ContractTriggerGasRegistryETHDeposited struct {
	User      common.Address
	EthAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterETHDeposited is a free log retrieval operation binding the contract event 0x6c703791f399558807424f489ccd811c72b4ff0b74af547264fad7c646776df0.
//
// Solidity: event ETHDeposited(address indexed user, uint256 ethAmount)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) FilterETHDeposited(opts *bind.FilterOpts, user []common.Address) (*ContractTriggerGasRegistryETHDepositedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ContractTriggerGasRegistry.contract.FilterLogs(opts, "ETHDeposited", userRule)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerGasRegistryETHDepositedIterator{contract: _ContractTriggerGasRegistry.contract, event: "ETHDeposited", logs: logs, sub: sub}, nil
}

// WatchETHDeposited is a free log subscription operation binding the contract event 0x6c703791f399558807424f489ccd811c72b4ff0b74af547264fad7c646776df0.
//
// Solidity: event ETHDeposited(address indexed user, uint256 ethAmount)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) WatchETHDeposited(opts *bind.WatchOpts, sink chan<- *ContractTriggerGasRegistryETHDeposited, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ContractTriggerGasRegistry.contract.WatchLogs(opts, "ETHDeposited", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerGasRegistryETHDeposited)
				if err := _ContractTriggerGasRegistry.contract.UnpackLog(event, "ETHDeposited", log); err != nil {
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

// ParseETHDeposited is a log parse operation binding the contract event 0x6c703791f399558807424f489ccd811c72b4ff0b74af547264fad7c646776df0.
//
// Solidity: event ETHDeposited(address indexed user, uint256 ethAmount)
func (_ContractTriggerGasRegistry *ContractTriggerGasRegistryFilterer) ParseETHDeposited(log types.Log) (*ContractTriggerGasRegistryETHDeposited, error) {
	event := new(ContractTriggerGasRegistryETHDeposited)
	if err := _ContractTriggerGasRegistry.contract.UnpackLog(event, "ETHDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
