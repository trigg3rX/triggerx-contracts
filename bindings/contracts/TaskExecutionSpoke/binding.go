// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractTaskExecutionSpoke

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

// ContractTaskExecutionSpokeMetaData contains all meta data concerning the ContractTaskExecutionSpoke contract.
var ContractTaskExecutionSpokeMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_endpoint\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_delegate\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowInitializePath\",\"inputs\":[{\"name\":\"origin\",\"type\":\"tuple\",\"internalType\":\"structOrigin\",\"components\":[{\"name\":\"srcEid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"sender\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"endpoint\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractILayerZeroEndpointV2\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"executeFunction\",\"inputs\":[{\"name\":\"jobId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ethAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_ownerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_hubEid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_initialKeepers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_jobRegistryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_triggerGasRegistryAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isComposeMsgSender\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOrigin\",\"components\":[{\"name\":\"srcEid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"sender\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_sender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isKeeper\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"jobRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIJobRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lzReceive\",\"inputs\":[{\"name\":\"_origin\",\"type\":\"tuple\",\"internalType\":\"structOrigin\",\"components\":[{\"name\":\"srcEid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"sender\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"_guid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_executor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"nextNonce\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"oAppVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"senderVersion\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiverVersion\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"peers\",\"inputs\":[{\"name\":\"eid\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"peer\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDelegate\",\"inputs\":[{\"name\":\"_delegate\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setJobRegistry\",\"inputs\":[{\"name\":\"_jobRegistryAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPeer\",\"inputs\":[{\"name\":\"_eid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_peer\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTriggerGasRegistry\",\"inputs\":[{\"name\":\"_triggerGasRegistryAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"triggerGasRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITriggerGasRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"FunctionExecuted\",\"inputs\":[{\"name\":\"keeper\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FunctionExecutionFailed\",\"inputs\":[{\"name\":\"keeper\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"result\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"KeeperUpdated\",\"inputs\":[{\"name\":\"action\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumTaskExecutionSpoke.ActionType\"},{\"name\":\"keeper\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PeerSet\",\"inputs\":[{\"name\":\"eid\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"peer\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDelegate\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidEndpointCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LzTokenUnavailable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoPeer\",\"inputs\":[{\"name\":\"eid\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"NotEnoughNative\",\"inputs\":[{\"name\":\"msgValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyEndpoint\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyPeer\",\"inputs\":[{\"name\":\"eid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"sender\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	Bin: "0x60c080604052346101cb575f9060408161165b803803809161002182856101e2565b8339810103126101cb5761003481610219565b906001600160a01b039061004a90602001610219565b1680156101cf575f80546001600160a01b0319811683178255604051939183916001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09080a36001600160a01b03166080819052803b156101cb576024835f8193819563ca5eb5e160e01b845260048401525af180156101c0576101ad575b503060a0525f51602061163b5f395f51905f52549060ff8260401c1661019e57506002600160401b03196001600160401b03821601610148575b60405161140d908161022e82396080518181816105cb0152818161082b01528181610c890152610eaa015260a05181818161086d01526109010152f35b6001600160401b0319166001600160401b039081175f51602061163b5f395f51905f52556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f61010b565b63f92ee8a960e01b8152600490fd5b6101b991505f906101e2565b5f5f6100d1565b6040513d5f823e3d90fd5b5f80fd5b631e4fbdf760e01b5f525f60045260245ffd5b601f909101601f19168101906001600160401b0382119082101761020557604052565b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b03821682036101cb5756fe60806040526004361015610011575f80fd5b5f5f3560e01c8062dc34c11461104057806313137d6514610e48578063149e0da414610b2457806317442b7014610b0257806323682c4714610ad95780633400288b14610aab5780634f1ef286146108c157806352d1902d1461085a5780635e280f11146108155780636ba42aaa146107d6578063715018a61461077c5780637d25a05e1461075657806382413eac146106fc5780638da5cb5b146106d5578063ad3cb1cc14610688578063bb0b6a5314610653578063ca5eb5e1146105a6578063d908b2ef1461057d578063e4d53fef14610538578063f2fde38b146104e7578063fa9b1a80146101475763ff7bd03d1461010b575f80fd5b346101445760603660031901126101445760209060409063ffffffff61012f6111c2565b16815260018352205460405190602435148152f35b80fd5b506080366003190112610144576044356001600160a01b03811690600435908281036103835760643567ffffffffffffffff81116103f65761018d903690600401611158565b91338552600260205260ff604086205416156104a25760025f5160206113985f395f51905f5254146104935760025f5160206113985f395f51905f525560018060a01b03600354169060405163bf8e9d5360e01b8152816004820152604081602481865afa908115610488578791610452575b50460361040d5760209060246040518094819363147f551160e01b835260048301525afa9081156104025785916103bc575b506001600160a01b03169081156103875760045485926001600160a01b0390911690813b15610383578391604483926040519485938492633df1e18160e21b8452600484015260243560248401525af190811561037857839161035f575b5090835190602085019034905af16102a66112bf565b901561030557507ff4448cdaf10358453fa19b2e0363f44780277619bd9bb676eef79d7488a7d2746102e36040519260408452604084019061119e565b913460208201528033930390a35b60015f5160206113985f395f51905f525580f35b906103427f5339dbdca52ae73f492e64bedd4159a0f4fe9b166a572305256af078ac94f28a9161035760405192839260608452606084019061119e565b3460208401528281036040840152339561119e565b0390a36102f1565b8161036991611106565b61037457815f610290565b5080fd5b6040513d85823e3d90fd5b8380fd5b60405162461bcd60e51b815260206004820152600d60248201526c129bd8881b9bdd08199bdd5b99609a1b6044820152606490fd5b90506020813d6020116103fa575b816103d760209383611106565b810103126103f657516001600160a01b03811681036103f6575f610232565b8480fd5b3d91506103ca565b6040513d87823e3d90fd5b60405162461bcd60e51b815260206004820152601d60248201527f4a6f622069732066726f6d206120646966666572656e7420636861696e0000006044820152606490fd5b90506040813d604011610480575b8161046d60409383611106565b8101031261047c57515f610200565b8680fd5b3d9150610460565b6040513d89823e3d90fd5b633ee5aeb560e01b8552600485fd5b60405162461bcd60e51b815260206004820152601c60248201527f53706f6b653a204b6565706572206e6f742072656769737465726564000000006044820152606490fd5b503461014457602036600319011261014457610501611083565b61050961120d565b6001600160a01b038116156105245761052190611279565b80f35b631e4fbdf760e01b82526004829052602482fd5b503461014457602036600319011261014457610552611083565b61055a61120d565b60018060a01b03166bffffffffffffffffffffffff60a01b600454161760045580f35b50346101445780600319360112610144576004546040516001600160a01b039091168152602090f35b503461014457602036600319011261014457806105c1611083565b6105c961120d565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690813b1561064f5760405163ca5eb5e160e01b81526001600160a01b0390911660048201529082908290602490829084905af18015610644576106335750f35b8161063d91611106565b6101445780f35b6040513d84823e3d90fd5b5050fd5b503461014457602036600319011261014457604060209163ffffffff6106776110f3565b168152600183522054604051908152f35b5034610144578060031936011261014457506106d16040516106ab604082611106565b60058152640352e302e360dc1b602082015260405191829160208352602083019061119e565b0390f35b5034610144578060031936011261014457546040516001600160a01b039091168152602090f35b503461014457366003190160a08112610374576060136101445760643567ffffffffffffffff8111610374576107369036906004016110c5565b505060206107426110af565b6040516001600160a01b0390911630148152f35b5034610144576040366003190112610144576020906107736110f3565b50604051908152f35b503461014457806003193601126101445761079561120d565b80546001600160a01b03198116825581906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b50346101445760203660031901126101445760209060ff906040906001600160a01b03610801611083565b168152600284522054166040519015158152f35b50346101445780600319360112610144576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b50346101445780600319360112610144577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031630036108b25760206040515f5160206113785f395f51905f528152f35b63703e46dd60e11b8152600490fd5b506040366003190112610144576108d6611083565b9060243567ffffffffffffffff8111610374576108f7903690600401611158565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016308114908115610a89575b50610a7a5761093961120d565b6040516352d1902d60e01b8152926001600160a01b0381169190602085600481865afa80958596610a46575b5061097e57634c9c8ce360e01b84526004839052602484fd5b9091845f5160206113785f395f51905f528103610a345750813b15610a22575f5160206113785f395f51905f5280546001600160a01b031916821790557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b8480a28151839015610a085780836020610a0495519101845af46109fe6112bf565b91611319565b5080f35b50505034610a135780f35b63b398979f60e01b8152600490fd5b634c9c8ce360e01b8452600452602483fd5b632a87526960e21b8552600452602484fd5b9095506020813d602011610a72575b81610a6260209383611106565b810103126103f65751945f610965565b3d9150610a55565b63703e46dd60e11b8252600482fd5b5f5160206113785f395f51905f52546001600160a01b0316141590505f61092c565b503461014457604036600319011261014457610521610ac86110f3565b610ad061120d565b60243590611233565b50346101445780600319360112610144576003546040516001600160a01b039091168152602090f35b5034610144578060031936011261014457604080516001815260026020820152f35b5034610d705760a0366003190112610d7057610b3e611083565b6024359063ffffffff82168203610d705760443567ffffffffffffffff8111610d705736602382011215610d705780600401359067ffffffffffffffff8211610d70576024810190602436918460051b010111610d70576064356001600160a01b0381169190829003610d7057610bb36110af565b925f5160206113b85f395f51905f525460ff8160401c16159667ffffffffffffffff821680159081610e40575b6001149081610e36575b159081610e2d575b50610e1e5767ffffffffffffffff1982166001175f5160206113b85f395f51905f5255610c509188610df2575b50610c286112ee565b610c306112ee565b60015f5160206113985f395f51905f5255610c496112ee565b3090611233565b5f5b818110610d74575050600380546001600160a01b031990811693909317905550600480549091166001600160a01b039283161790557f000000000000000000000000000000000000000000000000000000000000000016803b15610d705760405163ca5eb5e160e01b81526001600160a01b0383166004820152905f908290602490829084905af18015610d6557610d4e575b50610cef90611279565b610cf65780f35b68ff0000000000000000195f5160206113b85f395f51905f5254165f5160206113b85f395f51905f52557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a180f35b610d5b9193505f90611106565b5f91610cef610ce5565b6040513d5f823e3d90fd5b5f80fd5b6001906001600160a01b03610d92610d8d8386886111d5565b6111f9565b165f52600260205260405f208260ff198254161790557fa16bcb5f8175f03e8484e9d840834e357a7ca38d877946862644bf1e078ee1b26040610dd9610d8d8487896111d5565b8151905f8252858060a01b03166020820152a101610c52565b68ffffffffffffffffff191668010000000000000001175f5160206113b85f395f51905f52555f610c1f565b63f92ee8a960e01b5f5260045ffd5b9050155f610bf2565b303b159150610bea565b899150610be0565b366003190160e08112610d7057606013610d705760843567ffffffffffffffff8111610d7057610e7c9036906004016110c5565b610e84611099565b5060c43567ffffffffffffffff8111610d7057610ea59036906004016110c5565b5050337f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03160361102d5763ffffffff610ee46111c2565b16805f52600160205260405f205490811561101b5750602435809103610ff857508160409181010312610d7057803560028110918215610d7057602001359160018060a01b038316809303610d705781610f9c57825f52600260205260405f20600160ff198254161790555b6040519015610f88577fa16bcb5f8175f03e8484e9d840834e357a7ca38d877946862644bf1e078ee1b29260409282526020820152a1005b634e487b7160e01b5f52602160045260245ffd5b60018203610fbd57825f52600260205260405f2060ff198154169055610f50565b60405162461bcd60e51b8152602060048201526013602482015272496e76616c696420616374696f6e207479706560681b6044820152606490fd5b63ffffffff6110056111c2565b63309afaf360e21b5f521660045260245260445ffd5b63f6ff4fb760e01b5f5260045260245ffd5b6391ac5e4f60e01b5f523360045260245ffd5b34610d70576020366003190112610d7057611059611083565b61106161120d565b600380546001600160a01b0319166001600160a01b0392909216919091179055005b600435906001600160a01b0382168203610d7057565b60a435906001600160a01b0382168203610d7057565b608435906001600160a01b0382168203610d7057565b9181601f84011215610d705782359167ffffffffffffffff8311610d705760208381860195010111610d7057565b6004359063ffffffff82168203610d7057565b90601f8019910116810190811067ffffffffffffffff82111761112857604052565b634e487b7160e01b5f52604160045260245ffd5b67ffffffffffffffff811161112857601f01601f191660200190565b81601f82011215610d705780359061116f8261113c565b9261117d6040519485611106565b82845260208383010111610d7057815f926020809301838601378301015290565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b60043563ffffffff81168103610d705790565b91908110156111e55760051b0190565b634e487b7160e01b5f52603260045260245ffd5b356001600160a01b0381168103610d705790565b5f546001600160a01b0316330361122057565b63118cdaa760e01b5f523360045260245ffd5b7f238399d427b947898edb290f5ff0f9109849b1c3ba196a42e35f00c50a54b98b9163ffffffff6040921690815f52600160205280835f205582519182526020820152a1565b5f80546001600160a01b039283166001600160a01b03198216811783559216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09080a3565b3d156112e9573d906112d08261113c565b916112de6040519384611106565b82523d5f602084013e565b606090565b60ff5f5160206113b85f395f51905f525460401c161561130a57565b631afcd79f60e31b5f5260045ffd5b9061133d575080511561132e57805190602001fd5b63d6bda27560e01b5f5260045ffd5b8151158061136e575b61134e575090565b639996b31560e01b5f9081526001600160a01b0391909116600452602490fd5b50803b1561134656fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a264697066735822122017c01bce9041629c65ff45a15964bbba546b9299f9531b75503b2a95a98e33d464736f6c634300081b0033f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00",
}

// ContractTaskExecutionSpokeABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractTaskExecutionSpokeMetaData.ABI instead.
var ContractTaskExecutionSpokeABI = ContractTaskExecutionSpokeMetaData.ABI

// ContractTaskExecutionSpokeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractTaskExecutionSpokeMetaData.Bin instead.
var ContractTaskExecutionSpokeBin = ContractTaskExecutionSpokeMetaData.Bin

// DeployContractTaskExecutionSpoke deploys a new Ethereum contract, binding an instance of ContractTaskExecutionSpoke to it.
func DeployContractTaskExecutionSpoke(auth *bind.TransactOpts, backend bind.ContractBackend, _endpoint common.Address, _delegate common.Address) (common.Address, *types.Transaction, *ContractTaskExecutionSpoke, error) {
	parsed, err := ContractTaskExecutionSpokeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractTaskExecutionSpokeBin), backend, _endpoint, _delegate)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractTaskExecutionSpoke{ContractTaskExecutionSpokeCaller: ContractTaskExecutionSpokeCaller{contract: contract}, ContractTaskExecutionSpokeTransactor: ContractTaskExecutionSpokeTransactor{contract: contract}, ContractTaskExecutionSpokeFilterer: ContractTaskExecutionSpokeFilterer{contract: contract}}, nil
}

// ContractTaskExecutionSpokeMethods is an auto generated interface around an Ethereum contract.
type ContractTaskExecutionSpokeMethods interface {
	ContractTaskExecutionSpokeCalls
	ContractTaskExecutionSpokeTransacts
	ContractTaskExecutionSpokeFilters
}

// ContractTaskExecutionSpokeCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractTaskExecutionSpokeCalls interface {
	UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error)

	AllowInitializePath(opts *bind.CallOpts, origin Origin) (bool, error)

	Endpoint(opts *bind.CallOpts) (common.Address, error)

	IsComposeMsgSender(opts *bind.CallOpts, arg0 Origin, arg1 []byte, _sender common.Address) (bool, error)

	IsKeeper(opts *bind.CallOpts, arg0 common.Address) (bool, error)

	JobRegistry(opts *bind.CallOpts) (common.Address, error)

	NextNonce(opts *bind.CallOpts, arg0 uint32, arg1 [32]byte) (uint64, error)

	OAppVersion(opts *bind.CallOpts) (struct {
		SenderVersion   uint64
		ReceiverVersion uint64
	}, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Peers(opts *bind.CallOpts, eid uint32) ([32]byte, error)

	ProxiableUUID(opts *bind.CallOpts) ([32]byte, error)

	TriggerGasRegistry(opts *bind.CallOpts) (common.Address, error)
}

// ContractTaskExecutionSpokeTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractTaskExecutionSpokeTransacts interface {
	ExecuteFunction(opts *bind.TransactOpts, jobId *big.Int, ethAmount *big.Int, target common.Address, data []byte) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts, _ownerAddress common.Address, _hubEid uint32, _initialKeepers []common.Address, _jobRegistryAddress common.Address, _triggerGasRegistryAddress common.Address) (*types.Transaction, error)

	LzReceive(opts *bind.TransactOpts, _origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error)

	RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetDelegate(opts *bind.TransactOpts, _delegate common.Address) (*types.Transaction, error)

	SetJobRegistry(opts *bind.TransactOpts, _jobRegistryAddress common.Address) (*types.Transaction, error)

	SetPeer(opts *bind.TransactOpts, _eid uint32, _peer [32]byte) (*types.Transaction, error)

	SetTriggerGasRegistry(opts *bind.TransactOpts, _triggerGasRegistryAddress common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error)

	UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error)
}

// ContractTaskExecutionSpokeFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractTaskExecutionSpokeFilters interface {
	FilterFunctionExecuted(opts *bind.FilterOpts, keeper []common.Address, target []common.Address) (*ContractTaskExecutionSpokeFunctionExecutedIterator, error)
	WatchFunctionExecuted(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionSpokeFunctionExecuted, keeper []common.Address, target []common.Address) (event.Subscription, error)
	ParseFunctionExecuted(log types.Log) (*ContractTaskExecutionSpokeFunctionExecuted, error)

	FilterFunctionExecutionFailed(opts *bind.FilterOpts, keeper []common.Address, target []common.Address) (*ContractTaskExecutionSpokeFunctionExecutionFailedIterator, error)
	WatchFunctionExecutionFailed(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionSpokeFunctionExecutionFailed, keeper []common.Address, target []common.Address) (event.Subscription, error)
	ParseFunctionExecutionFailed(log types.Log) (*ContractTaskExecutionSpokeFunctionExecutionFailed, error)

	FilterInitialized(opts *bind.FilterOpts) (*ContractTaskExecutionSpokeInitializedIterator, error)
	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionSpokeInitialized) (event.Subscription, error)
	ParseInitialized(log types.Log) (*ContractTaskExecutionSpokeInitialized, error)

	FilterKeeperUpdated(opts *bind.FilterOpts) (*ContractTaskExecutionSpokeKeeperUpdatedIterator, error)
	WatchKeeperUpdated(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionSpokeKeeperUpdated) (event.Subscription, error)
	ParseKeeperUpdated(log types.Log) (*ContractTaskExecutionSpokeKeeperUpdated, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractTaskExecutionSpokeOwnershipTransferredIterator, error)
	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionSpokeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error)
	ParseOwnershipTransferred(log types.Log) (*ContractTaskExecutionSpokeOwnershipTransferred, error)

	FilterPeerSet(opts *bind.FilterOpts) (*ContractTaskExecutionSpokePeerSetIterator, error)
	WatchPeerSet(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionSpokePeerSet) (event.Subscription, error)
	ParsePeerSet(log types.Log) (*ContractTaskExecutionSpokePeerSet, error)

	FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ContractTaskExecutionSpokeUpgradedIterator, error)
	WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionSpokeUpgraded, implementation []common.Address) (event.Subscription, error)
	ParseUpgraded(log types.Log) (*ContractTaskExecutionSpokeUpgraded, error)
}

// ContractTaskExecutionSpoke is an auto generated Go binding around an Ethereum contract.
type ContractTaskExecutionSpoke struct {
	ContractTaskExecutionSpokeCaller     // Read-only binding to the contract
	ContractTaskExecutionSpokeTransactor // Write-only binding to the contract
	ContractTaskExecutionSpokeFilterer   // Log filterer for contract events
}

// ContractTaskExecutionSpoke implements the ContractTaskExecutionSpokeMethods interface.
var _ ContractTaskExecutionSpokeMethods = (*ContractTaskExecutionSpoke)(nil)

// ContractTaskExecutionSpokeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractTaskExecutionSpokeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTaskExecutionSpokeCaller implements the ContractTaskExecutionSpokeCalls interface.
var _ ContractTaskExecutionSpokeCalls = (*ContractTaskExecutionSpokeCaller)(nil)

// ContractTaskExecutionSpokeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTaskExecutionSpokeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTaskExecutionSpokeTransactor implements the ContractTaskExecutionSpokeTransacts interface.
var _ ContractTaskExecutionSpokeTransacts = (*ContractTaskExecutionSpokeTransactor)(nil)

// ContractTaskExecutionSpokeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractTaskExecutionSpokeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTaskExecutionSpokeFilterer implements the ContractTaskExecutionSpokeFilters interface.
var _ ContractTaskExecutionSpokeFilters = (*ContractTaskExecutionSpokeFilterer)(nil)

// ContractTaskExecutionSpokeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractTaskExecutionSpokeSession struct {
	Contract     *ContractTaskExecutionSpoke // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ContractTaskExecutionSpokeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractTaskExecutionSpokeCallerSession struct {
	Contract *ContractTaskExecutionSpokeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// ContractTaskExecutionSpokeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTaskExecutionSpokeTransactorSession struct {
	Contract     *ContractTaskExecutionSpokeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// ContractTaskExecutionSpokeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractTaskExecutionSpokeRaw struct {
	Contract *ContractTaskExecutionSpoke // Generic contract binding to access the raw methods on
}

// ContractTaskExecutionSpokeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractTaskExecutionSpokeCallerRaw struct {
	Contract *ContractTaskExecutionSpokeCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTaskExecutionSpokeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTaskExecutionSpokeTransactorRaw struct {
	Contract *ContractTaskExecutionSpokeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractTaskExecutionSpoke creates a new instance of ContractTaskExecutionSpoke, bound to a specific deployed contract.
func NewContractTaskExecutionSpoke(address common.Address, backend bind.ContractBackend) (*ContractTaskExecutionSpoke, error) {
	contract, err := bindContractTaskExecutionSpoke(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractTaskExecutionSpoke{ContractTaskExecutionSpokeCaller: ContractTaskExecutionSpokeCaller{contract: contract}, ContractTaskExecutionSpokeTransactor: ContractTaskExecutionSpokeTransactor{contract: contract}, ContractTaskExecutionSpokeFilterer: ContractTaskExecutionSpokeFilterer{contract: contract}}, nil
}

// NewContractTaskExecutionSpokeCaller creates a new read-only instance of ContractTaskExecutionSpoke, bound to a specific deployed contract.
func NewContractTaskExecutionSpokeCaller(address common.Address, caller bind.ContractCaller) (*ContractTaskExecutionSpokeCaller, error) {
	contract, err := bindContractTaskExecutionSpoke(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTaskExecutionSpokeCaller{contract: contract}, nil
}

// NewContractTaskExecutionSpokeTransactor creates a new write-only instance of ContractTaskExecutionSpoke, bound to a specific deployed contract.
func NewContractTaskExecutionSpokeTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTaskExecutionSpokeTransactor, error) {
	contract, err := bindContractTaskExecutionSpoke(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTaskExecutionSpokeTransactor{contract: contract}, nil
}

// NewContractTaskExecutionSpokeFilterer creates a new log filterer instance of ContractTaskExecutionSpoke, bound to a specific deployed contract.
func NewContractTaskExecutionSpokeFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractTaskExecutionSpokeFilterer, error) {
	contract, err := bindContractTaskExecutionSpoke(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractTaskExecutionSpokeFilterer{contract: contract}, nil
}

// bindContractTaskExecutionSpoke binds a generic wrapper to an already deployed contract.
func bindContractTaskExecutionSpoke(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractTaskExecutionSpokeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractTaskExecutionSpoke.Contract.ContractTaskExecutionSpokeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.Contract.ContractTaskExecutionSpokeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.Contract.ContractTaskExecutionSpokeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractTaskExecutionSpoke.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ContractTaskExecutionSpoke.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ContractTaskExecutionSpoke.Contract.UPGRADEINTERFACEVERSION(&_ContractTaskExecutionSpoke.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ContractTaskExecutionSpoke.Contract.UPGRADEINTERFACEVERSION(&_ContractTaskExecutionSpoke.CallOpts)
}

// AllowInitializePath is a free data retrieval call binding the contract method 0xff7bd03d.
//
// Solidity: function allowInitializePath((uint32,bytes32,uint64) origin) view returns(bool)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCaller) AllowInitializePath(opts *bind.CallOpts, origin Origin) (bool, error) {
	var out []interface{}
	err := _ContractTaskExecutionSpoke.contract.Call(opts, &out, "allowInitializePath", origin)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowInitializePath is a free data retrieval call binding the contract method 0xff7bd03d.
//
// Solidity: function allowInitializePath((uint32,bytes32,uint64) origin) view returns(bool)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeSession) AllowInitializePath(origin Origin) (bool, error) {
	return _ContractTaskExecutionSpoke.Contract.AllowInitializePath(&_ContractTaskExecutionSpoke.CallOpts, origin)
}

// AllowInitializePath is a free data retrieval call binding the contract method 0xff7bd03d.
//
// Solidity: function allowInitializePath((uint32,bytes32,uint64) origin) view returns(bool)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCallerSession) AllowInitializePath(origin Origin) (bool, error) {
	return _ContractTaskExecutionSpoke.Contract.AllowInitializePath(&_ContractTaskExecutionSpoke.CallOpts, origin)
}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCaller) Endpoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTaskExecutionSpoke.contract.Call(opts, &out, "endpoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeSession) Endpoint() (common.Address, error) {
	return _ContractTaskExecutionSpoke.Contract.Endpoint(&_ContractTaskExecutionSpoke.CallOpts)
}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCallerSession) Endpoint() (common.Address, error) {
	return _ContractTaskExecutionSpoke.Contract.Endpoint(&_ContractTaskExecutionSpoke.CallOpts)
}

// IsComposeMsgSender is a free data retrieval call binding the contract method 0x82413eac.
//
// Solidity: function isComposeMsgSender((uint32,bytes32,uint64) , bytes , address _sender) view returns(bool)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCaller) IsComposeMsgSender(opts *bind.CallOpts, arg0 Origin, arg1 []byte, _sender common.Address) (bool, error) {
	var out []interface{}
	err := _ContractTaskExecutionSpoke.contract.Call(opts, &out, "isComposeMsgSender", arg0, arg1, _sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsComposeMsgSender is a free data retrieval call binding the contract method 0x82413eac.
//
// Solidity: function isComposeMsgSender((uint32,bytes32,uint64) , bytes , address _sender) view returns(bool)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeSession) IsComposeMsgSender(arg0 Origin, arg1 []byte, _sender common.Address) (bool, error) {
	return _ContractTaskExecutionSpoke.Contract.IsComposeMsgSender(&_ContractTaskExecutionSpoke.CallOpts, arg0, arg1, _sender)
}

// IsComposeMsgSender is a free data retrieval call binding the contract method 0x82413eac.
//
// Solidity: function isComposeMsgSender((uint32,bytes32,uint64) , bytes , address _sender) view returns(bool)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCallerSession) IsComposeMsgSender(arg0 Origin, arg1 []byte, _sender common.Address) (bool, error) {
	return _ContractTaskExecutionSpoke.Contract.IsComposeMsgSender(&_ContractTaskExecutionSpoke.CallOpts, arg0, arg1, _sender)
}

// IsKeeper is a free data retrieval call binding the contract method 0x6ba42aaa.
//
// Solidity: function isKeeper(address ) view returns(bool)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCaller) IsKeeper(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ContractTaskExecutionSpoke.contract.Call(opts, &out, "isKeeper", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsKeeper is a free data retrieval call binding the contract method 0x6ba42aaa.
//
// Solidity: function isKeeper(address ) view returns(bool)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeSession) IsKeeper(arg0 common.Address) (bool, error) {
	return _ContractTaskExecutionSpoke.Contract.IsKeeper(&_ContractTaskExecutionSpoke.CallOpts, arg0)
}

// IsKeeper is a free data retrieval call binding the contract method 0x6ba42aaa.
//
// Solidity: function isKeeper(address ) view returns(bool)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCallerSession) IsKeeper(arg0 common.Address) (bool, error) {
	return _ContractTaskExecutionSpoke.Contract.IsKeeper(&_ContractTaskExecutionSpoke.CallOpts, arg0)
}

// JobRegistry is a free data retrieval call binding the contract method 0x23682c47.
//
// Solidity: function jobRegistry() view returns(address)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCaller) JobRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTaskExecutionSpoke.contract.Call(opts, &out, "jobRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// JobRegistry is a free data retrieval call binding the contract method 0x23682c47.
//
// Solidity: function jobRegistry() view returns(address)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeSession) JobRegistry() (common.Address, error) {
	return _ContractTaskExecutionSpoke.Contract.JobRegistry(&_ContractTaskExecutionSpoke.CallOpts)
}

// JobRegistry is a free data retrieval call binding the contract method 0x23682c47.
//
// Solidity: function jobRegistry() view returns(address)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCallerSession) JobRegistry() (common.Address, error) {
	return _ContractTaskExecutionSpoke.Contract.JobRegistry(&_ContractTaskExecutionSpoke.CallOpts)
}

// NextNonce is a free data retrieval call binding the contract method 0x7d25a05e.
//
// Solidity: function nextNonce(uint32 , bytes32 ) view returns(uint64 nonce)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCaller) NextNonce(opts *bind.CallOpts, arg0 uint32, arg1 [32]byte) (uint64, error) {
	var out []interface{}
	err := _ContractTaskExecutionSpoke.contract.Call(opts, &out, "nextNonce", arg0, arg1)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NextNonce is a free data retrieval call binding the contract method 0x7d25a05e.
//
// Solidity: function nextNonce(uint32 , bytes32 ) view returns(uint64 nonce)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeSession) NextNonce(arg0 uint32, arg1 [32]byte) (uint64, error) {
	return _ContractTaskExecutionSpoke.Contract.NextNonce(&_ContractTaskExecutionSpoke.CallOpts, arg0, arg1)
}

// NextNonce is a free data retrieval call binding the contract method 0x7d25a05e.
//
// Solidity: function nextNonce(uint32 , bytes32 ) view returns(uint64 nonce)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCallerSession) NextNonce(arg0 uint32, arg1 [32]byte) (uint64, error) {
	return _ContractTaskExecutionSpoke.Contract.NextNonce(&_ContractTaskExecutionSpoke.CallOpts, arg0, arg1)
}

// OAppVersion is a free data retrieval call binding the contract method 0x17442b70.
//
// Solidity: function oAppVersion() pure returns(uint64 senderVersion, uint64 receiverVersion)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCaller) OAppVersion(opts *bind.CallOpts) (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	var out []interface{}
	err := _ContractTaskExecutionSpoke.contract.Call(opts, &out, "oAppVersion")

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
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeSession) OAppVersion() (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	return _ContractTaskExecutionSpoke.Contract.OAppVersion(&_ContractTaskExecutionSpoke.CallOpts)
}

// OAppVersion is a free data retrieval call binding the contract method 0x17442b70.
//
// Solidity: function oAppVersion() pure returns(uint64 senderVersion, uint64 receiverVersion)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCallerSession) OAppVersion() (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	return _ContractTaskExecutionSpoke.Contract.OAppVersion(&_ContractTaskExecutionSpoke.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTaskExecutionSpoke.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeSession) Owner() (common.Address, error) {
	return _ContractTaskExecutionSpoke.Contract.Owner(&_ContractTaskExecutionSpoke.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCallerSession) Owner() (common.Address, error) {
	return _ContractTaskExecutionSpoke.Contract.Owner(&_ContractTaskExecutionSpoke.CallOpts)
}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 eid) view returns(bytes32 peer)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCaller) Peers(opts *bind.CallOpts, eid uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractTaskExecutionSpoke.contract.Call(opts, &out, "peers", eid)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 eid) view returns(bytes32 peer)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeSession) Peers(eid uint32) ([32]byte, error) {
	return _ContractTaskExecutionSpoke.Contract.Peers(&_ContractTaskExecutionSpoke.CallOpts, eid)
}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 eid) view returns(bytes32 peer)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCallerSession) Peers(eid uint32) ([32]byte, error) {
	return _ContractTaskExecutionSpoke.Contract.Peers(&_ContractTaskExecutionSpoke.CallOpts, eid)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractTaskExecutionSpoke.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeSession) ProxiableUUID() ([32]byte, error) {
	return _ContractTaskExecutionSpoke.Contract.ProxiableUUID(&_ContractTaskExecutionSpoke.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCallerSession) ProxiableUUID() ([32]byte, error) {
	return _ContractTaskExecutionSpoke.Contract.ProxiableUUID(&_ContractTaskExecutionSpoke.CallOpts)
}

// TriggerGasRegistry is a free data retrieval call binding the contract method 0xd908b2ef.
//
// Solidity: function triggerGasRegistry() view returns(address)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCaller) TriggerGasRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTaskExecutionSpoke.contract.Call(opts, &out, "triggerGasRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TriggerGasRegistry is a free data retrieval call binding the contract method 0xd908b2ef.
//
// Solidity: function triggerGasRegistry() view returns(address)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeSession) TriggerGasRegistry() (common.Address, error) {
	return _ContractTaskExecutionSpoke.Contract.TriggerGasRegistry(&_ContractTaskExecutionSpoke.CallOpts)
}

// TriggerGasRegistry is a free data retrieval call binding the contract method 0xd908b2ef.
//
// Solidity: function triggerGasRegistry() view returns(address)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeCallerSession) TriggerGasRegistry() (common.Address, error) {
	return _ContractTaskExecutionSpoke.Contract.TriggerGasRegistry(&_ContractTaskExecutionSpoke.CallOpts)
}

// ExecuteFunction is a paid mutator transaction binding the contract method 0xfa9b1a80.
//
// Solidity: function executeFunction(uint256 jobId, uint256 ethAmount, address target, bytes data) payable returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeTransactor) ExecuteFunction(opts *bind.TransactOpts, jobId *big.Int, ethAmount *big.Int, target common.Address, data []byte) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.contract.Transact(opts, "executeFunction", jobId, ethAmount, target, data)
}

// ExecuteFunction is a paid mutator transaction binding the contract method 0xfa9b1a80.
//
// Solidity: function executeFunction(uint256 jobId, uint256 ethAmount, address target, bytes data) payable returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeSession) ExecuteFunction(jobId *big.Int, ethAmount *big.Int, target common.Address, data []byte) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.Contract.ExecuteFunction(&_ContractTaskExecutionSpoke.TransactOpts, jobId, ethAmount, target, data)
}

// ExecuteFunction is a paid mutator transaction binding the contract method 0xfa9b1a80.
//
// Solidity: function executeFunction(uint256 jobId, uint256 ethAmount, address target, bytes data) payable returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeTransactorSession) ExecuteFunction(jobId *big.Int, ethAmount *big.Int, target common.Address, data []byte) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.Contract.ExecuteFunction(&_ContractTaskExecutionSpoke.TransactOpts, jobId, ethAmount, target, data)
}

// Initialize is a paid mutator transaction binding the contract method 0x149e0da4.
//
// Solidity: function initialize(address _ownerAddress, uint32 _hubEid, address[] _initialKeepers, address _jobRegistryAddress, address _triggerGasRegistryAddress) returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeTransactor) Initialize(opts *bind.TransactOpts, _ownerAddress common.Address, _hubEid uint32, _initialKeepers []common.Address, _jobRegistryAddress common.Address, _triggerGasRegistryAddress common.Address) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.contract.Transact(opts, "initialize", _ownerAddress, _hubEid, _initialKeepers, _jobRegistryAddress, _triggerGasRegistryAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x149e0da4.
//
// Solidity: function initialize(address _ownerAddress, uint32 _hubEid, address[] _initialKeepers, address _jobRegistryAddress, address _triggerGasRegistryAddress) returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeSession) Initialize(_ownerAddress common.Address, _hubEid uint32, _initialKeepers []common.Address, _jobRegistryAddress common.Address, _triggerGasRegistryAddress common.Address) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.Contract.Initialize(&_ContractTaskExecutionSpoke.TransactOpts, _ownerAddress, _hubEid, _initialKeepers, _jobRegistryAddress, _triggerGasRegistryAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x149e0da4.
//
// Solidity: function initialize(address _ownerAddress, uint32 _hubEid, address[] _initialKeepers, address _jobRegistryAddress, address _triggerGasRegistryAddress) returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeTransactorSession) Initialize(_ownerAddress common.Address, _hubEid uint32, _initialKeepers []common.Address, _jobRegistryAddress common.Address, _triggerGasRegistryAddress common.Address) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.Contract.Initialize(&_ContractTaskExecutionSpoke.TransactOpts, _ownerAddress, _hubEid, _initialKeepers, _jobRegistryAddress, _triggerGasRegistryAddress)
}

// LzReceive is a paid mutator transaction binding the contract method 0x13137d65.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeTransactor) LzReceive(opts *bind.TransactOpts, _origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.contract.Transact(opts, "lzReceive", _origin, _guid, _message, _executor, _extraData)
}

// LzReceive is a paid mutator transaction binding the contract method 0x13137d65.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeSession) LzReceive(_origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.Contract.LzReceive(&_ContractTaskExecutionSpoke.TransactOpts, _origin, _guid, _message, _executor, _extraData)
}

// LzReceive is a paid mutator transaction binding the contract method 0x13137d65.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeTransactorSession) LzReceive(_origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.Contract.LzReceive(&_ContractTaskExecutionSpoke.TransactOpts, _origin, _guid, _message, _executor, _extraData)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.Contract.RenounceOwnership(&_ContractTaskExecutionSpoke.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.Contract.RenounceOwnership(&_ContractTaskExecutionSpoke.TransactOpts)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeTransactor) SetDelegate(opts *bind.TransactOpts, _delegate common.Address) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.contract.Transact(opts, "setDelegate", _delegate)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeSession) SetDelegate(_delegate common.Address) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.Contract.SetDelegate(&_ContractTaskExecutionSpoke.TransactOpts, _delegate)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeTransactorSession) SetDelegate(_delegate common.Address) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.Contract.SetDelegate(&_ContractTaskExecutionSpoke.TransactOpts, _delegate)
}

// SetJobRegistry is a paid mutator transaction binding the contract method 0x00dc34c1.
//
// Solidity: function setJobRegistry(address _jobRegistryAddress) returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeTransactor) SetJobRegistry(opts *bind.TransactOpts, _jobRegistryAddress common.Address) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.contract.Transact(opts, "setJobRegistry", _jobRegistryAddress)
}

// SetJobRegistry is a paid mutator transaction binding the contract method 0x00dc34c1.
//
// Solidity: function setJobRegistry(address _jobRegistryAddress) returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeSession) SetJobRegistry(_jobRegistryAddress common.Address) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.Contract.SetJobRegistry(&_ContractTaskExecutionSpoke.TransactOpts, _jobRegistryAddress)
}

// SetJobRegistry is a paid mutator transaction binding the contract method 0x00dc34c1.
//
// Solidity: function setJobRegistry(address _jobRegistryAddress) returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeTransactorSession) SetJobRegistry(_jobRegistryAddress common.Address) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.Contract.SetJobRegistry(&_ContractTaskExecutionSpoke.TransactOpts, _jobRegistryAddress)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeTransactor) SetPeer(opts *bind.TransactOpts, _eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.contract.Transact(opts, "setPeer", _eid, _peer)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeSession) SetPeer(_eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.Contract.SetPeer(&_ContractTaskExecutionSpoke.TransactOpts, _eid, _peer)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeTransactorSession) SetPeer(_eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.Contract.SetPeer(&_ContractTaskExecutionSpoke.TransactOpts, _eid, _peer)
}

// SetTriggerGasRegistry is a paid mutator transaction binding the contract method 0xe4d53fef.
//
// Solidity: function setTriggerGasRegistry(address _triggerGasRegistryAddress) returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeTransactor) SetTriggerGasRegistry(opts *bind.TransactOpts, _triggerGasRegistryAddress common.Address) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.contract.Transact(opts, "setTriggerGasRegistry", _triggerGasRegistryAddress)
}

// SetTriggerGasRegistry is a paid mutator transaction binding the contract method 0xe4d53fef.
//
// Solidity: function setTriggerGasRegistry(address _triggerGasRegistryAddress) returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeSession) SetTriggerGasRegistry(_triggerGasRegistryAddress common.Address) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.Contract.SetTriggerGasRegistry(&_ContractTaskExecutionSpoke.TransactOpts, _triggerGasRegistryAddress)
}

// SetTriggerGasRegistry is a paid mutator transaction binding the contract method 0xe4d53fef.
//
// Solidity: function setTriggerGasRegistry(address _triggerGasRegistryAddress) returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeTransactorSession) SetTriggerGasRegistry(_triggerGasRegistryAddress common.Address) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.Contract.SetTriggerGasRegistry(&_ContractTaskExecutionSpoke.TransactOpts, _triggerGasRegistryAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.Contract.TransferOwnership(&_ContractTaskExecutionSpoke.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.Contract.TransferOwnership(&_ContractTaskExecutionSpoke.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.Contract.UpgradeToAndCall(&_ContractTaskExecutionSpoke.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ContractTaskExecutionSpoke.Contract.UpgradeToAndCall(&_ContractTaskExecutionSpoke.TransactOpts, newImplementation, data)
}

// ContractTaskExecutionSpokeFunctionExecutedIterator is returned from FilterFunctionExecuted and is used to iterate over the raw logs and unpacked data for FunctionExecuted events raised by the ContractTaskExecutionSpoke contract.
type ContractTaskExecutionSpokeFunctionExecutedIterator struct {
	Event *ContractTaskExecutionSpokeFunctionExecuted // Event containing the contract specifics and raw log

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
func (it *ContractTaskExecutionSpokeFunctionExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTaskExecutionSpokeFunctionExecuted)
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
		it.Event = new(ContractTaskExecutionSpokeFunctionExecuted)
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
func (it *ContractTaskExecutionSpokeFunctionExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTaskExecutionSpokeFunctionExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTaskExecutionSpokeFunctionExecuted represents a FunctionExecuted event raised by the ContractTaskExecutionSpoke contract.
type ContractTaskExecutionSpokeFunctionExecuted struct {
	Keeper common.Address
	Target common.Address
	Data   []byte
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFunctionExecuted is a free log retrieval operation binding the contract event 0xf4448cdaf10358453fa19b2e0363f44780277619bd9bb676eef79d7488a7d274.
//
// Solidity: event FunctionExecuted(address indexed keeper, address indexed target, bytes data, uint256 value)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeFilterer) FilterFunctionExecuted(opts *bind.FilterOpts, keeper []common.Address, target []common.Address) (*ContractTaskExecutionSpokeFunctionExecutedIterator, error) {

	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _ContractTaskExecutionSpoke.contract.FilterLogs(opts, "FunctionExecuted", keeperRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &ContractTaskExecutionSpokeFunctionExecutedIterator{contract: _ContractTaskExecutionSpoke.contract, event: "FunctionExecuted", logs: logs, sub: sub}, nil
}

// WatchFunctionExecuted is a free log subscription operation binding the contract event 0xf4448cdaf10358453fa19b2e0363f44780277619bd9bb676eef79d7488a7d274.
//
// Solidity: event FunctionExecuted(address indexed keeper, address indexed target, bytes data, uint256 value)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeFilterer) WatchFunctionExecuted(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionSpokeFunctionExecuted, keeper []common.Address, target []common.Address) (event.Subscription, error) {

	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _ContractTaskExecutionSpoke.contract.WatchLogs(opts, "FunctionExecuted", keeperRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTaskExecutionSpokeFunctionExecuted)
				if err := _ContractTaskExecutionSpoke.contract.UnpackLog(event, "FunctionExecuted", log); err != nil {
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
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeFilterer) ParseFunctionExecuted(log types.Log) (*ContractTaskExecutionSpokeFunctionExecuted, error) {
	event := new(ContractTaskExecutionSpokeFunctionExecuted)
	if err := _ContractTaskExecutionSpoke.contract.UnpackLog(event, "FunctionExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTaskExecutionSpokeFunctionExecutionFailedIterator is returned from FilterFunctionExecutionFailed and is used to iterate over the raw logs and unpacked data for FunctionExecutionFailed events raised by the ContractTaskExecutionSpoke contract.
type ContractTaskExecutionSpokeFunctionExecutionFailedIterator struct {
	Event *ContractTaskExecutionSpokeFunctionExecutionFailed // Event containing the contract specifics and raw log

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
func (it *ContractTaskExecutionSpokeFunctionExecutionFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTaskExecutionSpokeFunctionExecutionFailed)
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
		it.Event = new(ContractTaskExecutionSpokeFunctionExecutionFailed)
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
func (it *ContractTaskExecutionSpokeFunctionExecutionFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTaskExecutionSpokeFunctionExecutionFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTaskExecutionSpokeFunctionExecutionFailed represents a FunctionExecutionFailed event raised by the ContractTaskExecutionSpoke contract.
type ContractTaskExecutionSpokeFunctionExecutionFailed struct {
	Keeper common.Address
	Target common.Address
	Data   []byte
	Value  *big.Int
	Result []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFunctionExecutionFailed is a free log retrieval operation binding the contract event 0x5339dbdca52ae73f492e64bedd4159a0f4fe9b166a572305256af078ac94f28a.
//
// Solidity: event FunctionExecutionFailed(address indexed keeper, address indexed target, bytes data, uint256 value, bytes result)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeFilterer) FilterFunctionExecutionFailed(opts *bind.FilterOpts, keeper []common.Address, target []common.Address) (*ContractTaskExecutionSpokeFunctionExecutionFailedIterator, error) {

	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _ContractTaskExecutionSpoke.contract.FilterLogs(opts, "FunctionExecutionFailed", keeperRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &ContractTaskExecutionSpokeFunctionExecutionFailedIterator{contract: _ContractTaskExecutionSpoke.contract, event: "FunctionExecutionFailed", logs: logs, sub: sub}, nil
}

// WatchFunctionExecutionFailed is a free log subscription operation binding the contract event 0x5339dbdca52ae73f492e64bedd4159a0f4fe9b166a572305256af078ac94f28a.
//
// Solidity: event FunctionExecutionFailed(address indexed keeper, address indexed target, bytes data, uint256 value, bytes result)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeFilterer) WatchFunctionExecutionFailed(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionSpokeFunctionExecutionFailed, keeper []common.Address, target []common.Address) (event.Subscription, error) {

	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _ContractTaskExecutionSpoke.contract.WatchLogs(opts, "FunctionExecutionFailed", keeperRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTaskExecutionSpokeFunctionExecutionFailed)
				if err := _ContractTaskExecutionSpoke.contract.UnpackLog(event, "FunctionExecutionFailed", log); err != nil {
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

// ParseFunctionExecutionFailed is a log parse operation binding the contract event 0x5339dbdca52ae73f492e64bedd4159a0f4fe9b166a572305256af078ac94f28a.
//
// Solidity: event FunctionExecutionFailed(address indexed keeper, address indexed target, bytes data, uint256 value, bytes result)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeFilterer) ParseFunctionExecutionFailed(log types.Log) (*ContractTaskExecutionSpokeFunctionExecutionFailed, error) {
	event := new(ContractTaskExecutionSpokeFunctionExecutionFailed)
	if err := _ContractTaskExecutionSpoke.contract.UnpackLog(event, "FunctionExecutionFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTaskExecutionSpokeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractTaskExecutionSpoke contract.
type ContractTaskExecutionSpokeInitializedIterator struct {
	Event *ContractTaskExecutionSpokeInitialized // Event containing the contract specifics and raw log

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
func (it *ContractTaskExecutionSpokeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTaskExecutionSpokeInitialized)
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
		it.Event = new(ContractTaskExecutionSpokeInitialized)
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
func (it *ContractTaskExecutionSpokeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTaskExecutionSpokeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTaskExecutionSpokeInitialized represents a Initialized event raised by the ContractTaskExecutionSpoke contract.
type ContractTaskExecutionSpokeInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractTaskExecutionSpokeInitializedIterator, error) {

	logs, sub, err := _ContractTaskExecutionSpoke.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractTaskExecutionSpokeInitializedIterator{contract: _ContractTaskExecutionSpoke.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionSpokeInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractTaskExecutionSpoke.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTaskExecutionSpokeInitialized)
				if err := _ContractTaskExecutionSpoke.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeFilterer) ParseInitialized(log types.Log) (*ContractTaskExecutionSpokeInitialized, error) {
	event := new(ContractTaskExecutionSpokeInitialized)
	if err := _ContractTaskExecutionSpoke.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTaskExecutionSpokeKeeperUpdatedIterator is returned from FilterKeeperUpdated and is used to iterate over the raw logs and unpacked data for KeeperUpdated events raised by the ContractTaskExecutionSpoke contract.
type ContractTaskExecutionSpokeKeeperUpdatedIterator struct {
	Event *ContractTaskExecutionSpokeKeeperUpdated // Event containing the contract specifics and raw log

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
func (it *ContractTaskExecutionSpokeKeeperUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTaskExecutionSpokeKeeperUpdated)
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
		it.Event = new(ContractTaskExecutionSpokeKeeperUpdated)
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
func (it *ContractTaskExecutionSpokeKeeperUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTaskExecutionSpokeKeeperUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTaskExecutionSpokeKeeperUpdated represents a KeeperUpdated event raised by the ContractTaskExecutionSpoke contract.
type ContractTaskExecutionSpokeKeeperUpdated struct {
	Action uint8
	Keeper common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterKeeperUpdated is a free log retrieval operation binding the contract event 0xa16bcb5f8175f03e8484e9d840834e357a7ca38d877946862644bf1e078ee1b2.
//
// Solidity: event KeeperUpdated(uint8 action, address keeper)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeFilterer) FilterKeeperUpdated(opts *bind.FilterOpts) (*ContractTaskExecutionSpokeKeeperUpdatedIterator, error) {

	logs, sub, err := _ContractTaskExecutionSpoke.contract.FilterLogs(opts, "KeeperUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractTaskExecutionSpokeKeeperUpdatedIterator{contract: _ContractTaskExecutionSpoke.contract, event: "KeeperUpdated", logs: logs, sub: sub}, nil
}

// WatchKeeperUpdated is a free log subscription operation binding the contract event 0xa16bcb5f8175f03e8484e9d840834e357a7ca38d877946862644bf1e078ee1b2.
//
// Solidity: event KeeperUpdated(uint8 action, address keeper)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeFilterer) WatchKeeperUpdated(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionSpokeKeeperUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractTaskExecutionSpoke.contract.WatchLogs(opts, "KeeperUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTaskExecutionSpokeKeeperUpdated)
				if err := _ContractTaskExecutionSpoke.contract.UnpackLog(event, "KeeperUpdated", log); err != nil {
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

// ParseKeeperUpdated is a log parse operation binding the contract event 0xa16bcb5f8175f03e8484e9d840834e357a7ca38d877946862644bf1e078ee1b2.
//
// Solidity: event KeeperUpdated(uint8 action, address keeper)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeFilterer) ParseKeeperUpdated(log types.Log) (*ContractTaskExecutionSpokeKeeperUpdated, error) {
	event := new(ContractTaskExecutionSpokeKeeperUpdated)
	if err := _ContractTaskExecutionSpoke.contract.UnpackLog(event, "KeeperUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTaskExecutionSpokeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractTaskExecutionSpoke contract.
type ContractTaskExecutionSpokeOwnershipTransferredIterator struct {
	Event *ContractTaskExecutionSpokeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractTaskExecutionSpokeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTaskExecutionSpokeOwnershipTransferred)
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
		it.Event = new(ContractTaskExecutionSpokeOwnershipTransferred)
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
func (it *ContractTaskExecutionSpokeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTaskExecutionSpokeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTaskExecutionSpokeOwnershipTransferred represents a OwnershipTransferred event raised by the ContractTaskExecutionSpoke contract.
type ContractTaskExecutionSpokeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractTaskExecutionSpokeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractTaskExecutionSpoke.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractTaskExecutionSpokeOwnershipTransferredIterator{contract: _ContractTaskExecutionSpoke.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionSpokeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractTaskExecutionSpoke.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTaskExecutionSpokeOwnershipTransferred)
				if err := _ContractTaskExecutionSpoke.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeFilterer) ParseOwnershipTransferred(log types.Log) (*ContractTaskExecutionSpokeOwnershipTransferred, error) {
	event := new(ContractTaskExecutionSpokeOwnershipTransferred)
	if err := _ContractTaskExecutionSpoke.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTaskExecutionSpokePeerSetIterator is returned from FilterPeerSet and is used to iterate over the raw logs and unpacked data for PeerSet events raised by the ContractTaskExecutionSpoke contract.
type ContractTaskExecutionSpokePeerSetIterator struct {
	Event *ContractTaskExecutionSpokePeerSet // Event containing the contract specifics and raw log

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
func (it *ContractTaskExecutionSpokePeerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTaskExecutionSpokePeerSet)
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
		it.Event = new(ContractTaskExecutionSpokePeerSet)
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
func (it *ContractTaskExecutionSpokePeerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTaskExecutionSpokePeerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTaskExecutionSpokePeerSet represents a PeerSet event raised by the ContractTaskExecutionSpoke contract.
type ContractTaskExecutionSpokePeerSet struct {
	Eid  uint32
	Peer [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPeerSet is a free log retrieval operation binding the contract event 0x238399d427b947898edb290f5ff0f9109849b1c3ba196a42e35f00c50a54b98b.
//
// Solidity: event PeerSet(uint32 eid, bytes32 peer)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeFilterer) FilterPeerSet(opts *bind.FilterOpts) (*ContractTaskExecutionSpokePeerSetIterator, error) {

	logs, sub, err := _ContractTaskExecutionSpoke.contract.FilterLogs(opts, "PeerSet")
	if err != nil {
		return nil, err
	}
	return &ContractTaskExecutionSpokePeerSetIterator{contract: _ContractTaskExecutionSpoke.contract, event: "PeerSet", logs: logs, sub: sub}, nil
}

// WatchPeerSet is a free log subscription operation binding the contract event 0x238399d427b947898edb290f5ff0f9109849b1c3ba196a42e35f00c50a54b98b.
//
// Solidity: event PeerSet(uint32 eid, bytes32 peer)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeFilterer) WatchPeerSet(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionSpokePeerSet) (event.Subscription, error) {

	logs, sub, err := _ContractTaskExecutionSpoke.contract.WatchLogs(opts, "PeerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTaskExecutionSpokePeerSet)
				if err := _ContractTaskExecutionSpoke.contract.UnpackLog(event, "PeerSet", log); err != nil {
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
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeFilterer) ParsePeerSet(log types.Log) (*ContractTaskExecutionSpokePeerSet, error) {
	event := new(ContractTaskExecutionSpokePeerSet)
	if err := _ContractTaskExecutionSpoke.contract.UnpackLog(event, "PeerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTaskExecutionSpokeUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ContractTaskExecutionSpoke contract.
type ContractTaskExecutionSpokeUpgradedIterator struct {
	Event *ContractTaskExecutionSpokeUpgraded // Event containing the contract specifics and raw log

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
func (it *ContractTaskExecutionSpokeUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTaskExecutionSpokeUpgraded)
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
		it.Event = new(ContractTaskExecutionSpokeUpgraded)
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
func (it *ContractTaskExecutionSpokeUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTaskExecutionSpokeUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTaskExecutionSpokeUpgraded represents a Upgraded event raised by the ContractTaskExecutionSpoke contract.
type ContractTaskExecutionSpokeUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ContractTaskExecutionSpokeUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ContractTaskExecutionSpoke.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ContractTaskExecutionSpokeUpgradedIterator{contract: _ContractTaskExecutionSpoke.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ContractTaskExecutionSpokeUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ContractTaskExecutionSpoke.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTaskExecutionSpokeUpgraded)
				if err := _ContractTaskExecutionSpoke.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_ContractTaskExecutionSpoke *ContractTaskExecutionSpokeFilterer) ParseUpgraded(log types.Log) (*ContractTaskExecutionSpokeUpgraded, error) {
	event := new(ContractTaskExecutionSpokeUpgraded)
	if err := _ContractTaskExecutionSpoke.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
