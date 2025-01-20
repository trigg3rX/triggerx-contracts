// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractOperatorStateRetriever

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

// OperatorStateRetrieverCheckSignaturesIndices is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverCheckSignaturesIndices struct {
	NonSignerQuorumBitmapIndices []uint32
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// OperatorStateRetrieverOperator is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverOperator struct {
	Operator   common.Address
	OperatorId [32]byte
	Stake      *big.Int
}

// ContractOperatorStateRetrieverMetaData contains all meta data concerning the ContractOperatorStateRetriever contract.
var ContractOperatorStateRetrieverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"OperatorNotRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIRegistryCoordinator\",\"name\":\"registryCoordinator\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"operatorIds\",\"type\":\"bytes32[]\"}],\"name\":\"getBatchOperatorFromId\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"operators\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRegistryCoordinator\",\"name\":\"registryCoordinator\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"operators\",\"type\":\"address[]\"}],\"name\":\"getBatchOperatorId\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"operatorIds\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRegistryCoordinator\",\"name\":\"registryCoordinator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"referenceBlockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\"}],\"name\":\"getCheckSignaturesIndices\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32[]\",\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[]\",\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[]\",\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[][]\",\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\"}],\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRegistryCoordinator\",\"name\":\"registryCoordinator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"}],\"name\":\"getOperatorState\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"internalType\":\"uint96\",\"name\":\"stake\",\"type\":\"uint96\"}],\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"name\":\"\",\"type\":\"tuple[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRegistryCoordinator\",\"name\":\"registryCoordinator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"}],\"name\":\"getOperatorState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"internalType\":\"uint96\",\"name\":\"stake\",\"type\":\"uint96\"}],\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"name\":\"\",\"type\":\"tuple[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRegistryCoordinator\",\"name\":\"registryCoordinator\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"operatorIds\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"}],\"name\":\"getQuorumBitmapsAtBlockNumber\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040526004361015610011575f80fd5b5f3560e01c806331b36bd914610c615780633563b0d114610bc85780634d2b57fe14610a8a5780634f739f741461046d5780635c155662146102b25763cefdc1d41461005b575f80fd5b346102ae5760603660031901126102ae57610074610db5565b60243590610080610e80565b916040926100d78451926100948685610dfa565b60018452601f198601366020860137806100ad85611002565b5285516361c8a12f60e11b81526001600160a01b0386169490925f918491829187600484016115e1565b0381875afa9182156102a45763ffffffff6100fe6020946064935f91610282575b50611002565b511691875195869384926304ec635160e01b8452600484015263ffffffff8716602484015260448301525afa918215610278575f92610247575b506001600160c01b038216915f83805b61020e575061ffff169261015b84610e65565b9361016887519586610dfa565b808552610177601f1991610e65565b013660208601375f925f5b8551851080610203575b156101da576001811b84166001600160c01b03166101b3575b6101ae906115d3565b610182565b9360016101ae9160ff60f81b8760f81b165f1a6101d0828a611072565b53019490506101a5565b87836101ff6101ea858a8c611083565b83519384938452806020850152830190610e93565b0390f35b50610100811061018c565b5f1981018181116102335761ffff9116911661ffff8114610233576001019080610148565b634e487b7160e01b5f52601160045260245ffd5b61026a91925060203d602011610271575b6102628183610dfa565b8101906115a8565b905f610138565b503d610258565b84513d5f823e3d90fd5b61029e91503d805f833e6102968183610dfa565b8101906114f7565b5f6100f8565b86513d5f823e3d90fd5b5f80fd5b346102ae5760603660031901126102ae576102cb610db5565b6024356001600160401b0381116102ae576102ea903690600401610f3a565b6102f2610e80565b6040516361c8a12f60e11b815290926001600160a01b03165f828061031b8688600484016115e1565b0381845afa918215610402575f92610451575b5082519361035461033e86610e1b565b9561034c6040519788610dfa565b808752610e1b565b602086019490601f19013686375f5b815181101561040d576103768183611023565b519060208463ffffffff61038a848a611023565b516040516304ec635160e01b8152600481019690965263ffffffff92831660248701521616604484015282606481885afa8015610402576001925f916103e4575b50828060c01b03166103dd828a611023565b5201610363565b6103fc915060203d8111610271576102628183610dfa565b896103cb565b6040513d5f823e3d90fd5b8587604051918291602083019060208452518091526040830191905f5b818110610438575050500390f35b825184528594506020938401939092019160010161042a565b6104669192503d805f833e6102968183610dfa565b908461032e565b346102ae5760803660031901126102ae57610486610db5565b60243563ffffffff8116908181036102ae57604435926001600160401b0384116102ae57366023850112156102ae578360040135936001600160401b0385116102ae57602481019060248636920101116102ae57606435926001600160401b0384116102ae57366023850112156102ae578360040135956001600160401b0387116102ae5760248501938760051b9560248736920101116102ae576105296114b2565b50604051636830483560e01b81526001600160a01b03919091169390602081600481885afa908115610402575f91610a6b575b506105656114b2565b604080516361c8a12f60e11b8152600481018b90526024810191909152604481018b905290976001600160fb1b038b116102ae5781606481835f948c848401378101030181895afa908115610402575f91610a51575b50875260018060a01b031691604051986340e03a8160e11b8a528860048b0152604060248b01525f8a806105f3604482018688611578565b0381875afa998a15610402575f9a610a35575b5060408801998a5261061782610e1b565b966106256040519889610dfa565b828852610641601f1961063785610e1b565b0160208a01611056565b606089019788525f5b60ff8116848110156108ac575f6106778261066487610fd0565b8d51906106718383611023565b52611023565b505f84868e5b8d8d8386106106fa5750505050505061069581610fd0565b905f5b8c8282106106ca5760ff959492506106b893915051906106718383611023565b501660ff81146102335760010161064a565b9063ffffffff6106e7826106e18860019651611023565b51611023565b51166106f38286611023565b5201610698565b63ffffffff61071b87610713816020986107559a611598565b359551611023565b516040516304ec635160e01b8152600481019590955263ffffffff9283166024860152161660448301529092839190829081906064820190565b03915afa908115610402575f9161088e575b506001600160c01b0316801561087f5760018091610786868b8d6115c7565b3560f81c1c161461079d575b60010184868e61067d565b908960206107ac848989611598565b356107b8868b8d6115c7565b60405163dd9846b960e01b815260048101929092523560f81c602482015263ffffffff929092166044830152816064818d5afa908115610402578d85915f93610829575b509163ffffffff610818856106e1600197956108219751611023565b911690526115d3565b919050610792565b925050506020813d8211610877575b8161084560209383610dfa565b810103126102ae57818d63ffffffff6108186001956106e18961086a610821986114e6565b97509550509550506107fc565b3d9150610838565b6325ec6c1f60e01b5f5260045ffd5b6108a6915060203d8111610271576102628183610dfa565b5f610767565b5089898c8e8760048a60208f60405193848092632efa2ca360e11b82525afa908115610402575f93610905938593610a04575b506040519687948593849363354952a360e21b8552604060048601526044850191611578565b602483019190915203916001600160a01b03165afa9182156104025761098392610970915f916109ea575b50908594939291602061095d970190815260405196879660208852516080602089015260a0880190610f97565b9051868203601f19016040880152610f97565b9051848203601f19016060860152610f97565b905190601f19838203016080840152815180825260208201916020808360051b8301019401925f915b8383106109b95786860387f35b9193955091936020806109d8600193601f198682030187528951610f97565b970193019301909286959492936109ac565b6109fe91503d805f833e6102968183610dfa565b86610930565b610a2791935060203d602011610a2e575b610a1f8183610dfa565b810190611037565b91896108df565b503d610a15565b610a4a919a503d805f833e6102968183610dfa565b988a610606565b610a6591503d805f833e6102968183610dfa565b8a6105bb565b610a84915060203d602011610a2e57610a1f8183610dfa565b8961055c565b346102ae5760403660031901126102ae57610aa3610db5565b6024356001600160401b0381116102ae57610ac2903690600401610f3a565b8051610ae6610ad082610e1b565b91610ade6040519384610dfa565b808352610e1b565b602082019290601f19013684376001600160a01b03909316925f5b8151811015610b7b57610b148183611023565b519060405191630a5aec1960e21b83526004830152602082602481895afa8015610402576001925f91610b5d575b50610b4d8286611023565b90838060a01b0316905201610b01565b610b75915060203d8111610a2e57610a1f8183610dfa565b87610b42565b8383604051918291602083019060208452518091526040830191905f5b818110610ba6575050500390f35b82516001600160a01b0316845285945060209384019390920191600101610b98565b346102ae5760603660031901126102ae57610be1610db5565b6024356001600160401b0381116102ae57366023820112156102ae57806004013591610c0c83610e65565b610c196040519182610dfa565b83815236602485850101116102ae575f6020856101ff966024610c4d97018386013783010152610c47610e80565b91611083565b604051918291602083526020830190610e93565b346102ae5760403660031901126102ae57610c7a610db5565b602435906001600160401b0382116102ae57366023830112156102ae578160040135610ca581610e1b565b92610cb36040519485610dfa565b8184526024602085019260051b820101903682116102ae57602401915b818310610d95578385610ce38151610fd0565b6001600160a01b03909216915f5b8251811015610d7f576001600160a01b03610d0c8285611023565b516040516309aa152760e11b81529116600482015290602082602481885afa8015610402575f90610d4d575b60019250610d468285611023565b5201610cf1565b506020823d8211610d77575b81610d6660209383610dfa565b810103126102ae5760019151610d38565b3d9150610d59565b604051602080825281906101ff90820185610e32565b82356001600160a01b03811681036102ae57815260209283019201610cd0565b600435906001600160a01b03821682036102ae57565b606081019081106001600160401b03821117610de657604052565b634e487b7160e01b5f52604160045260245ffd5b90601f801991011681019081106001600160401b03821117610de657604052565b6001600160401b038111610de65760051b60200190565b90602080835192838152019201905f5b818110610e4f5750505090565b8251845260209384019390920191600101610e42565b6001600160401b038111610de657601f01601f191660200190565b6044359063ffffffff821682036102ae57565b9080602083519182815201916020808360051b8301019401925f915b838310610ebe57505050505090565b9091929394601f19828203018352855190602080835192838152019201905f905b808210610efe5750505060208060019297019301930191939290610eaf565b909192602060606001926001600160601b0360408851868060a01b03815116845285810151868501520151166040820152019401920190610edf565b9080601f830112156102ae578135610f5181610e1b565b92610f5f6040519485610dfa565b81845260208085019260051b8201019283116102ae57602001905b828210610f875750505090565b8135815260209182019101610f7a565b90602080835192838152019201905f5b818110610fb45750505090565b825163ffffffff16845260209384019390920191600101610fa7565b90610fda82610e1b565b610fe76040519182610dfa565b8281528092610ff8601f1991610e1b565b0190602036910137565b80511561100f5760200190565b634e487b7160e01b5f52603260045260245ffd5b805182101561100f5760209160051b010190565b908160209103126102ae57516001600160a01b03811681036102ae5790565b5f5b82811061106457505050565b606082820152602001611058565b90815181101561100f570160200190565b604051636830483560e01b81526001600160a01b0390911692909190602083600481875afa928315610402575f93611491575b50604051634f4c91e160e11b815292602084600481885afa938415610402575f9461144c575b5060206004949560405195868092632efa2ca360e11b82525afa938415610402575f9461142b575b509194939085519261113f61111885610e1b565b946111266040519687610dfa565b808652611135601f1991610e1b565b0160208601611056565b5f965b8051881015611421576111558882611072565b51604051638902624560e01b815260f89190911c6004820181905263ffffffff851660248301529790945f866044816001600160a01b0385165afa958615610402575f9661138f575b5085516111aa81610e1b565b906111b86040519283610dfa565b8082526111c7601f1991610e1b565b015f5b8181106113665750506111dd8b89611023565b526111e88a88611023565b505f5b8651811015611355576111fe8188611023565b516040516308f6629d60e31b81526004810191909152906020826024816001600160a01b038e165afa918215610402575f92611335575b5086611241828a611023565b5160208d61124f858d611023565b5160405163fa28c62760e01b8152600481019190915260ff91909116602482015263ffffffff939093166044840152826064816001600160a01b038c165afa908115610402578e925f926112ee575b50936112db6112e7936001600160601b038694600198604051956112c187610dcb565b8a8060a01b0316865260208601521660408401528d611023565b51906106718383611023565b50016111eb565b915091506020813d821161132d575b8161130a60209383610dfa565b810103126102ae57516001600160601b03811681036102ae578d916112db61129e565b3d91506112fd565b61134e91925060203d8111610a2e57610a1f8183610dfa565b905f611235565b506001909901989097509350611142565b60209060405161137581610dcb565b5f81525f838201525f6040820152828286010152016111ca565b9095503d805f833e6113a18183610dfa565b8101906020818303126102ae578051906001600160401b0382116102ae57019080601f830112156102ae5781516113d781610e1b565b926113e56040519485610dfa565b81845260208085019260051b8201019283116102ae57602001905b82821061141157505050945f61119e565b8151815260209182019101611400565b5092955050505050565b61144591945060203d602011610a2e57610a1f8183610dfa565b925f611104565b9093506020813d602011611489575b8161146860209383610dfa565b810103126102ae5751926001600160a01b03841684036102ae5760206110dc565b3d915061145b565b6114ab91935060203d602011610a2e57610a1f8183610dfa565b915f6110b6565b60405190608082018281106001600160401b03821117610de657604052606080838181528160208201528160408201520152565b519063ffffffff821682036102ae57565b6020818303126102ae578051906001600160401b0382116102ae57019080601f830112156102ae57815161152a81610e1b565b926115386040519485610dfa565b81845260208085019260051b8201019283116102ae57602001905b8282106115605750505090565b6020809161156d846114e6565b815201910190611553565b908060209392818452848401375f828201840152601f01601f1916010190565b919081101561100f5760051b0190565b908160209103126102ae57516001600160c01b03811681036102ae5790565b9082101561100f570190565b5f1981146102335760010190565b60409063ffffffff6115fe94931681528160208201520190610e32565b9056fea2646970667358221220ce80381b82d29905d32e0849a254fb90917e564f884e763f0e76102f81c7653564736f6c634300081b0033",
}

// ContractOperatorStateRetrieverABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractOperatorStateRetrieverMetaData.ABI instead.
var ContractOperatorStateRetrieverABI = ContractOperatorStateRetrieverMetaData.ABI

// ContractOperatorStateRetrieverBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractOperatorStateRetrieverMetaData.Bin instead.
var ContractOperatorStateRetrieverBin = ContractOperatorStateRetrieverMetaData.Bin

// DeployContractOperatorStateRetriever deploys a new Ethereum contract, binding an instance of ContractOperatorStateRetriever to it.
func DeployContractOperatorStateRetriever(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContractOperatorStateRetriever, error) {
	parsed, err := ContractOperatorStateRetrieverMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractOperatorStateRetrieverBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractOperatorStateRetriever{ContractOperatorStateRetrieverCaller: ContractOperatorStateRetrieverCaller{contract: contract}, ContractOperatorStateRetrieverTransactor: ContractOperatorStateRetrieverTransactor{contract: contract}, ContractOperatorStateRetrieverFilterer: ContractOperatorStateRetrieverFilterer{contract: contract}}, nil
}

// ContractOperatorStateRetrieverMethods is an auto generated interface around an Ethereum contract.
type ContractOperatorStateRetrieverMethods interface {
	ContractOperatorStateRetrieverCalls
	ContractOperatorStateRetrieverTransacts
	ContractOperatorStateRetrieverFilters
}

// ContractOperatorStateRetrieverCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractOperatorStateRetrieverCalls interface {
	GetBatchOperatorFromId(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error)

	GetBatchOperatorId(opts *bind.CallOpts, registryCoordinator common.Address, operators []common.Address) ([][32]byte, error)

	GetCheckSignaturesIndices(opts *bind.CallOpts, registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error)

	GetOperatorState(opts *bind.CallOpts, registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error)

	GetOperatorState0(opts *bind.CallOpts, registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error)

	GetQuorumBitmapsAtBlockNumber(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error)
}

// ContractOperatorStateRetrieverTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractOperatorStateRetrieverTransacts interface {
}

// ContractOperatorStateRetrieverFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractOperatorStateRetrieverFilters interface {
}

// ContractOperatorStateRetriever is an auto generated Go binding around an Ethereum contract.
type ContractOperatorStateRetriever struct {
	ContractOperatorStateRetrieverCaller     // Read-only binding to the contract
	ContractOperatorStateRetrieverTransactor // Write-only binding to the contract
	ContractOperatorStateRetrieverFilterer   // Log filterer for contract events
}

// ContractOperatorStateRetriever implements the ContractOperatorStateRetrieverMethods interface.
var _ ContractOperatorStateRetrieverMethods = (*ContractOperatorStateRetriever)(nil)

// ContractOperatorStateRetrieverCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractOperatorStateRetrieverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractOperatorStateRetrieverCaller implements the ContractOperatorStateRetrieverCalls interface.
var _ ContractOperatorStateRetrieverCalls = (*ContractOperatorStateRetrieverCaller)(nil)

// ContractOperatorStateRetrieverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractOperatorStateRetrieverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractOperatorStateRetrieverTransactor implements the ContractOperatorStateRetrieverTransacts interface.
var _ ContractOperatorStateRetrieverTransacts = (*ContractOperatorStateRetrieverTransactor)(nil)

// ContractOperatorStateRetrieverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractOperatorStateRetrieverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractOperatorStateRetrieverFilterer implements the ContractOperatorStateRetrieverFilters interface.
var _ ContractOperatorStateRetrieverFilters = (*ContractOperatorStateRetrieverFilterer)(nil)

// ContractOperatorStateRetrieverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractOperatorStateRetrieverSession struct {
	Contract     *ContractOperatorStateRetriever // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                   // Call options to use throughout this session
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ContractOperatorStateRetrieverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractOperatorStateRetrieverCallerSession struct {
	Contract *ContractOperatorStateRetrieverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                         // Call options to use throughout this session
}

// ContractOperatorStateRetrieverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractOperatorStateRetrieverTransactorSession struct {
	Contract     *ContractOperatorStateRetrieverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                         // Transaction auth options to use throughout this session
}

// ContractOperatorStateRetrieverRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractOperatorStateRetrieverRaw struct {
	Contract *ContractOperatorStateRetriever // Generic contract binding to access the raw methods on
}

// ContractOperatorStateRetrieverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractOperatorStateRetrieverCallerRaw struct {
	Contract *ContractOperatorStateRetrieverCaller // Generic read-only contract binding to access the raw methods on
}

// ContractOperatorStateRetrieverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractOperatorStateRetrieverTransactorRaw struct {
	Contract *ContractOperatorStateRetrieverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractOperatorStateRetriever creates a new instance of ContractOperatorStateRetriever, bound to a specific deployed contract.
func NewContractOperatorStateRetriever(address common.Address, backend bind.ContractBackend) (*ContractOperatorStateRetriever, error) {
	contract, err := bindContractOperatorStateRetriever(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractOperatorStateRetriever{ContractOperatorStateRetrieverCaller: ContractOperatorStateRetrieverCaller{contract: contract}, ContractOperatorStateRetrieverTransactor: ContractOperatorStateRetrieverTransactor{contract: contract}, ContractOperatorStateRetrieverFilterer: ContractOperatorStateRetrieverFilterer{contract: contract}}, nil
}

// NewContractOperatorStateRetrieverCaller creates a new read-only instance of ContractOperatorStateRetriever, bound to a specific deployed contract.
func NewContractOperatorStateRetrieverCaller(address common.Address, caller bind.ContractCaller) (*ContractOperatorStateRetrieverCaller, error) {
	contract, err := bindContractOperatorStateRetriever(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractOperatorStateRetrieverCaller{contract: contract}, nil
}

// NewContractOperatorStateRetrieverTransactor creates a new write-only instance of ContractOperatorStateRetriever, bound to a specific deployed contract.
func NewContractOperatorStateRetrieverTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractOperatorStateRetrieverTransactor, error) {
	contract, err := bindContractOperatorStateRetriever(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractOperatorStateRetrieverTransactor{contract: contract}, nil
}

// NewContractOperatorStateRetrieverFilterer creates a new log filterer instance of ContractOperatorStateRetriever, bound to a specific deployed contract.
func NewContractOperatorStateRetrieverFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractOperatorStateRetrieverFilterer, error) {
	contract, err := bindContractOperatorStateRetriever(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractOperatorStateRetrieverFilterer{contract: contract}, nil
}

// bindContractOperatorStateRetriever binds a generic wrapper to an already deployed contract.
func bindContractOperatorStateRetriever(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractOperatorStateRetrieverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractOperatorStateRetriever.Contract.ContractOperatorStateRetrieverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractOperatorStateRetriever.Contract.ContractOperatorStateRetrieverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractOperatorStateRetriever.Contract.ContractOperatorStateRetrieverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractOperatorStateRetriever.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractOperatorStateRetriever.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractOperatorStateRetriever.Contract.contract.Transact(opts, method, params...)
}

// GetBatchOperatorFromId is a free data retrieval call binding the contract method 0x4d2b57fe.
//
// Solidity: function getBatchOperatorFromId(address registryCoordinator, bytes32[] operatorIds) view returns(address[] operators)
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCaller) GetBatchOperatorFromId(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _ContractOperatorStateRetriever.contract.Call(opts, &out, "getBatchOperatorFromId", registryCoordinator, operatorIds)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetBatchOperatorFromId is a free data retrieval call binding the contract method 0x4d2b57fe.
//
// Solidity: function getBatchOperatorFromId(address registryCoordinator, bytes32[] operatorIds) view returns(address[] operators)
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverSession) GetBatchOperatorFromId(registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error) {
	return _ContractOperatorStateRetriever.Contract.GetBatchOperatorFromId(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, operatorIds)
}

// GetBatchOperatorFromId is a free data retrieval call binding the contract method 0x4d2b57fe.
//
// Solidity: function getBatchOperatorFromId(address registryCoordinator, bytes32[] operatorIds) view returns(address[] operators)
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCallerSession) GetBatchOperatorFromId(registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error) {
	return _ContractOperatorStateRetriever.Contract.GetBatchOperatorFromId(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, operatorIds)
}

// GetBatchOperatorId is a free data retrieval call binding the contract method 0x31b36bd9.
//
// Solidity: function getBatchOperatorId(address registryCoordinator, address[] operators) view returns(bytes32[] operatorIds)
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCaller) GetBatchOperatorId(opts *bind.CallOpts, registryCoordinator common.Address, operators []common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _ContractOperatorStateRetriever.contract.Call(opts, &out, "getBatchOperatorId", registryCoordinator, operators)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetBatchOperatorId is a free data retrieval call binding the contract method 0x31b36bd9.
//
// Solidity: function getBatchOperatorId(address registryCoordinator, address[] operators) view returns(bytes32[] operatorIds)
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverSession) GetBatchOperatorId(registryCoordinator common.Address, operators []common.Address) ([][32]byte, error) {
	return _ContractOperatorStateRetriever.Contract.GetBatchOperatorId(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, operators)
}

// GetBatchOperatorId is a free data retrieval call binding the contract method 0x31b36bd9.
//
// Solidity: function getBatchOperatorId(address registryCoordinator, address[] operators) view returns(bytes32[] operatorIds)
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCallerSession) GetBatchOperatorId(registryCoordinator common.Address, operators []common.Address) ([][32]byte, error) {
	return _ContractOperatorStateRetriever.Contract.GetBatchOperatorId(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, operators)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCaller) GetCheckSignaturesIndices(opts *bind.CallOpts, registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	var out []interface{}
	err := _ContractOperatorStateRetriever.contract.Call(opts, &out, "getCheckSignaturesIndices", registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)

	if err != nil {
		return *new(OperatorStateRetrieverCheckSignaturesIndices), err
	}

	out0 := *abi.ConvertType(out[0], new(OperatorStateRetrieverCheckSignaturesIndices)).(*OperatorStateRetrieverCheckSignaturesIndices)

	return out0, err

}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractOperatorStateRetriever.Contract.GetCheckSignaturesIndices(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCallerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractOperatorStateRetriever.Contract.GetCheckSignaturesIndices(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCaller) GetOperatorState(opts *bind.CallOpts, registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractOperatorStateRetriever.contract.Call(opts, &out, "getOperatorState", registryCoordinator, quorumNumbers, blockNumber)

	if err != nil {
		return *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, err

}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractOperatorStateRetriever.Contract.GetOperatorState(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCallerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractOperatorStateRetriever.Contract.GetOperatorState(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCaller) GetOperatorState0(opts *bind.CallOpts, registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractOperatorStateRetriever.contract.Call(opts, &out, "getOperatorState0", registryCoordinator, operatorId, blockNumber)

	if err != nil {
		return *new(*big.Int), *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, out1, err

}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractOperatorStateRetriever.Contract.GetOperatorState0(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCallerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractOperatorStateRetriever.Contract.GetOperatorState0(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCaller) GetQuorumBitmapsAtBlockNumber(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	var out []interface{}
	err := _ContractOperatorStateRetriever.contract.Call(opts, &out, "getQuorumBitmapsAtBlockNumber", registryCoordinator, operatorIds, blockNumber)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractOperatorStateRetriever.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCallerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractOperatorStateRetriever.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, operatorIds, blockNumber)
}
