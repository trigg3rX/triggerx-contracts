// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractJobRegistry

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

// JobRegistryJob is an auto generated low-level Go binding around an user-defined struct.
type JobRegistryJob struct {
	JobId         *big.Int
	JobOwner      common.Address
	JobHash       [32]byte
	LastUpdatedAt *big.Int
	IsActive      bool
}

// ContractJobRegistryMetaData contains all meta data concerning the ContractJobRegistry contract.
var ContractJobRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createJob\",\"inputs\":[{\"name\":\"jobName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"jobType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"timeFrame\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"targetContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"jobId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deleteJob\",\"inputs\":[{\"name\":\"jobId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getJob\",\"inputs\":[{\"name\":\"jobId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"job\",\"type\":\"tuple\",\"internalType\":\"structJobRegistry.Job\",\"components\":[{\"name\":\"jobId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"jobOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"jobHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"lastUpdatedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getJobByCounter\",\"inputs\":[{\"name\":\"jobCounter\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"job\",\"type\":\"tuple\",\"internalType\":\"structJobRegistry.Job\",\"components\":[{\"name\":\"jobId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"jobOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"jobHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"lastUpdatedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getJobCounter\",\"inputs\":[],\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getJobOwner\",\"inputs\":[{\"name\":\"jobId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUserActiveJobIds\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"activeJobIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUserJobIds\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"jobIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isJobActive\",\"inputs\":[{\"name\":\"jobId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpackJobId\",\"inputs\":[{\"name\":\"jobId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"jobCounter\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"updateJob\",\"inputs\":[{\"name\":\"jobId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"oldJobName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"jobType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"oldTimeFrame\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"targetContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"oldData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"newJobName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"newTimeFrame\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"JobCreated\",\"inputs\":[{\"name\":\"jobId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"jobOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"jobHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"JobDeleted\",\"inputs\":[{\"name\":\"jobId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"jobOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"JobUpdated\",\"inputs\":[{\"name\":\"jobId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"jobOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newJobHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ChainIdOverflow\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyJobName\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidJobData\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidJobParameters\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTargetContract\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"JobAlreadyInactive\",\"inputs\":[{\"name\":\"jobId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"JobCounterOverflow\",\"inputs\":[{\"name\":\"jobCounter\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxJobCounter\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"JobNotFound\",\"inputs\":[{\"name\":\"jobId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MissingIpfsHash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MissingTimeInterval\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TimestampOverflow\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedJobAccess\",\"inputs\":[{\"name\":\"jobId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60a080604052346100c257306080525f5160206116025f395f51905f525460ff8160401c166100b3576002600160401b03196001600160401b03821601610060575b60405161153b90816100c782396080518181816102ce01526103f80152f35b6001600160401b0319166001600160401b039081175f5160206116025f395f51905f525581527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80610041565b63f92ee8a960e01b5f5260045ffd5b5f80fdfe60806040526004361015610011575f80fd5b5f3560e01c8063147f5511146101345780632257b8da1461012f5780634f1ef2861461012a57806352d1902d1461012557806353bc234f146101205780635523d8dc1461011b5780636358bb8c14610116578063715018a614610111578063893de3c51461010c5780638da5cb5b1461010757806391461d0a14610102578063ad3cb1cc146100fd578063bf22c457146100f8578063bf8e9d53146100f3578063c4d66de8146100ee578063ef178814146100e9578063f2fde38b146100e45763ff7b49a0146100df575f80fd5b610cf9565b610cd0565b610ca1565b610b3a565b610af8565b610a3e565b6109f7565b61091d565b6108e9565b6106f5565b61066e565b61059c565b61051e565b61043d565b6103e6565b610290565b61016d565b34610169576020366003190112610169576004355f526001602052602060018060a01b03600160405f20015416604051908152f35b5f80fd5b34610169575f3660031901126101695760205f54604051908152f35b600435906001600160a01b038216820361016957565b606435906001600160a01b038216820361016957565b608435906001600160a01b038216820361016957565b634e487b7160e01b5f52604160045260245ffd5b60a081019081106001600160401b038211176101fa57604052565b6101cb565b90601f801991011681019081106001600160401b038211176101fa57604052565b6040519061022f60a0836101ff565b565b6001600160401b0381116101fa57601f01601f191660200190565b81601f820112156101695760208135910161026682610231565b9261027460405194856101ff565b8284528282011161016957815f92602092838601378301015290565b6040366003190112610169576102a4610189565b6024356001600160401b038111610169576102c390369060040161024c565b906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000163081149081156103c4575b506103b5576103066111b7565b6040516352d1902d60e01b8152916020836004816001600160a01b0386165afa5f9381610384575b5061034f57634c9c8ce360e01b5f526001600160a01b03821660045260245ffd5b905f5160206114c65f395f51905f5283036103705761036e925061133a565b005b632a87526960e21b5f52600483905260245ffd5b6103a791945060203d6020116103ae575b61039f81836101ff565b8101906111a8565b925f61032e565b503d610395565b63703e46dd60e11b5f5260045ffd5b5f5160206114c65f395f51905f52546001600160a01b0316141590505f6102f9565b34610169575f366003190112610169577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031630036103b55760206040515f5160206114c65f395f51905f528152f35b3461016957602036600319011261016957600435805f52600160205260405f20906040519161046b836101df565b8054835260018101546001600160a01b031660208401819052600282015460408501526003820154606085015260049091015460ff1615156080909301928352156104d3576104cf6104bd8351151590565b60405190151581529081906020820190565b0390f35b6350c83b9560e01b5f5260045260245ffd5b60206040818301928281528451809452019201905f5b8181106105085750505090565b82518452602093840193909201916001016104fb565b34610169576020366003190112610169576001600160a01b0361053f610189565b165f52600260205260405f206040519081602082549182815201915f5260205f20905f5b818110610586576104cf8561057a818703826101ff565b604051918291826104e5565b8254845260209093019260019283019201610563565b34610169576020366003190112610169576004356105c2815f52600160205260405f2090565b60018101546001600160a01b0316801561065a573303610642576004016105f16105ed825460ff1690565b1590565b61062e57805460ff191690556040514281523391907f8bc0131dae2d0cbc363614dfcf9653c32125d8fe69ea2cc1f6be9bc43019853790602090a3005b63060a06f160e01b5f52600482905260245ffd5b6316aceee960e21b5f5260048290523360245260445ffd5b6350c83b9560e01b5f52600483905260245ffd5b34610169575f366003190112610169576106866111b7565b5f5160206114a65f395f51905f5280546001600160a01b031981169091555f906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b6024359060ff8216820361016957565b6044359060ff8216820361016957565b346101695760a0366003190112610169576004356001600160401b0381116101695761072590369060040161024c565b61072d6106d5565b906044359061073a61019f565b916084356001600160401b0381116101695761075a90369060040161024c565b8251156108da57600760ff86161415806108c9575b6108ba57846107b1826107886107bf946104cf996111ea565b61079f6107955f54610e42565b805f5542466112ba565b96604051958694602086019889610e64565b03601f1981018352826101ff565b51902061085d6107cd610220565b8381523360208201528260408201524260608201526107ef6080820160019052565b610801845f52600160205260405f2090565b906080600491805184556001840160018060a01b036020830151166bffffffffffffffffffffffff60a01b82541617905560408101516002850155606081015160038501550151151591019060ff801983541691151516179055565b335f908152600260205260409020610876908390610ebf565b60408051918252426020830152339183917f737fc62fbb05dd9fb7c799e68796a2d7c8324e310af7b656c0580e7b3cf8bf8a91a36040519081529081906020820190565b632ab4c28d60e11b5f5260045ffd5b506001600160a01b0384161561076f565b630895f00360e01b5f5260045ffd5b34610169575f366003190112610169575f5160206114a65f395f51905f52546040516001600160a01b039091168152602090f35b3461016957610120366003190112610169576004356024356001600160401b0381116101695761095190369060040161024c565b9061095a6106e5565b6064356109656101b5565b60a4356001600160401b0381116101695761098490369060040161024c565b9060c4356001600160401b038111610169576109a490369060040161024c565b9260e4359461010435976001600160401b038911610169576109cd61036e99369060040161024c565b97610eeb565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b34610169575f366003190112610169576104cf604051610a186040826101ff565b60058152640352e302e360dc1b60208201526040519182916020835260208301906109d3565b3461016957602036600319011261016957600435610a5a611079565b50805f52600160205260405f209060405191610a75836101df565b8054835260018101546001600160a01b031660208401819052600282015460408501526003820154606085015260049091015460ff1615156080840152156104d35760408051835181526020808501516001600160a01b03169082015283820151918101919091526060808401519082015260808084015115159082015260a090f35b3461016957602036600319011261016957606060043560018060a01b036001600160401b0391604051928160e01c84528160a01c166020840152166040820152f35b3461016957602036600319011261016957610b53610189565b5f5160206114e65f395f51905f5254906001600160401b03610b8460ff604085901c1615936001600160401b031690565b1680159081610c99575b6001149081610c8f575b159081610c86575b50610c7757610be39082610bda60016001600160401b03195f5160206114e65f395f51905f525416175f5160206114e65f395f51905f5255565b610c42576110a3565b610be957005b610c1360ff60401b195f5160206114e65f395f51905f5254165f5160206114e65f395f51905f5255565b604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1005b610c72600160401b60ff60401b195f5160206114e65f395f51905f525416175f5160206114e65f395f51905f5255565b6110a3565b63f92ee8a960e01b5f5260045ffd5b9050155f610ba0565b303b159150610b98565b839150610b8e565b3461016957602036600319011261016957600435610cbd611079565b506350c83b9560e01b5f5260045260245ffd5b346101695760203660031901126101695761036e610cec610189565b610cf46111b7565b6110da565b3461016957602036600319011261016957610d36610d31610d18610189565b6001600160a01b03165f90815260026020526040902090565b610dfa565b5f5f5b8251811015610d8c57610d6b6004610d63610d54848761114b565b515f52600160205260405f2090565b015460ff1690565b610d78575b600101610d39565b90610d84600191610e42565b919050610d70565b50610d9690611176565b905f5f5b8251811015610dec57610db56004610d63610d54848761114b565b610dc2575b600101610d9a565b90610de4600191610dd3848661114b565b51610dde828861114b565b52610e42565b919050610dba565b604051806104cf86826104e5565b90604051918281549182825260208201905f5260205f20925f5b818110610e2957505061022f925003836101ff565b8454835260019485019487945060209093019201610e14565b5f198114610e505760010190565b634e487b7160e01b5f52601160045260245ffd5b9193909260ff610e80610ea8979560a0865260a08601906109d3565b9516602084015260408301526001600160a01b031660608201528083036080909101526109d3565b90565b634e487b7160e01b5f52603260045260245ffd5b805490600160401b8210156101fa5760018201808255821015610ee6575f5260205f200155565b610eab565b96939792959495610f04885f52600160205260405f2090565b60018101549097906001600160a01b0316801561106557330361104d57610f326105ed60048a015460ff1690565b6110395789916107b1610f519260405194859388602086019889610e64565b519020966002860197885403610ffe578451156108da57600760ff8316141580610fed575b6108ba57826107b191610f8c610f9d95856111ea565b604051958694602086019889610e64565b5190208093556003429101557f7dac5d7787afd94b2012ccf48531da941b75ac7e9cfc8b6ea0b55cb3a93dbd1360405180610fe8339542908360209093929193604081019481520152565b0390a3565b506001600160a01b03811615610f76565b60405162461bcd60e51b815260206004820152601360248201527209e9888beac8298aa8aa6be9a92a69a82a8869606b1b6044820152606490fd5b63060a06f160e01b5f52600489905260245ffd5b6316aceee960e21b5f5260048990523360245260445ffd5b6350c83b9560e01b5f5260048a905260245ffd5b60405190611086826101df565b5f6080838281528260208201528260408201528260608201520152565b6001600160a01b038116156110cb576110c6906110be6113dc565b610cf46113dc565b5f8055565b639e1d143d60e01b5f5260045ffd5b6001600160a01b03168015611138575f5160206114a65f395f51905f5280546001600160a01b0319811683179091556001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3565b631e4fbdf760e01b5f525f60045260245ffd5b8051821015610ee65760209160051b010190565b6001600160401b0381116101fa5760051b60200190565b906111808261115f565b61118d60405191826101ff565b828152809261119e601f199161115f565b0190602036910137565b90816020910312610169575190565b5f5160206114a65f395f51905f52546001600160a01b031633036111d757565b63118cdaa760e01b5f523360045260245ffd5b60ff166001811480156112b0575b80156112a6575b1561128d57602082511061127e5760208201511561127e575b60028114908115611273575b8115611268575b811561125d575b5061123a5750565b604081511061124e57604001511561124e57565b6334eab45360e11b5f5260045ffd5b60079150145f611232565b60068114915061122b565b600481149150611224565b638485995760e01b5f5260045ffd5b6020825110156112185763a103823560e01b5f5260045ffd5b50600781146111ff565b50600281146111f8565b909163ffffffff80831161132457506001600160401b039081841161130d576001600160a01b03908181116112f85716921660a01b9060e01b171790565b6318d5148f60e11b5f5260045260245260445ffd5b5082631a13274f60e31b5f5260045260245260445ffd5b8263343cb90960e21b5f5260045260245260445ffd5b90813b156113bb575f5160206114c65f395f51905f5280546001600160a01b0319166001600160a01b0384169081179091557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b5f80a28051156113a3576113a091611407565b50565b5050346113ac57565b63b398979f60e01b5f5260045ffd5b50634c9c8ce360e01b5f9081526001600160a01b0391909116600452602490fd5b60ff5f5160206114e65f395f51905f525460401c16156113f857565b631afcd79f60e31b5f5260045ffd5b5f80610ea893602081519101845af43d15611443573d9161142783610231565b9261143560405194856101ff565b83523d5f602085013e611447565b6060915b9061146b575080511561145c57805190602001fd5b63d6bda27560e01b5f5260045ffd5b8151158061149c575b61147c575090565b639996b31560e01b5f9081526001600160a01b0391909116600452602490fd5b50803b1561147456fe9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbcf0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a264697066735822122088b241b0f92c0514b21de45707e7a0aa3b5d95ab44954917ef2e748d36322d3264736f6c634300081b0033f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00",
}

// ContractJobRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractJobRegistryMetaData.ABI instead.
var ContractJobRegistryABI = ContractJobRegistryMetaData.ABI

// ContractJobRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractJobRegistryMetaData.Bin instead.
var ContractJobRegistryBin = ContractJobRegistryMetaData.Bin

// DeployContractJobRegistry deploys a new Ethereum contract, binding an instance of ContractJobRegistry to it.
func DeployContractJobRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContractJobRegistry, error) {
	parsed, err := ContractJobRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractJobRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractJobRegistry{ContractJobRegistryCaller: ContractJobRegistryCaller{contract: contract}, ContractJobRegistryTransactor: ContractJobRegistryTransactor{contract: contract}, ContractJobRegistryFilterer: ContractJobRegistryFilterer{contract: contract}}, nil
}

// ContractJobRegistryMethods is an auto generated interface around an Ethereum contract.
type ContractJobRegistryMethods interface {
	ContractJobRegistryCalls
	ContractJobRegistryTransacts
	ContractJobRegistryFilters
}

// ContractJobRegistryCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractJobRegistryCalls interface {
	UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error)

	GetJob(opts *bind.CallOpts, jobId *big.Int) (JobRegistryJob, error)

	GetJobByCounter(opts *bind.CallOpts, jobCounter *big.Int) (JobRegistryJob, error)

	GetJobCounter(opts *bind.CallOpts) (*big.Int, error)

	GetJobOwner(opts *bind.CallOpts, jobId *big.Int) (common.Address, error)

	GetUserActiveJobIds(opts *bind.CallOpts, user common.Address) ([]*big.Int, error)

	GetUserJobIds(opts *bind.CallOpts, user common.Address) ([]*big.Int, error)

	IsJobActive(opts *bind.CallOpts, jobId *big.Int) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	ProxiableUUID(opts *bind.CallOpts) ([32]byte, error)

	UnpackJobId(opts *bind.CallOpts, jobId *big.Int) (struct {
		ChainId    *big.Int
		Timestamp  *big.Int
		JobCounter *big.Int
	}, error)
}

// ContractJobRegistryTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractJobRegistryTransacts interface {
	CreateJob(opts *bind.TransactOpts, jobName string, jobType uint8, timeFrame *big.Int, targetContract common.Address, data []byte) (*types.Transaction, error)

	DeleteJob(opts *bind.TransactOpts, jobId *big.Int) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts, initialOwner common.Address) (*types.Transaction, error)

	RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error)

	UpdateJob(opts *bind.TransactOpts, jobId *big.Int, oldJobName string, jobType uint8, oldTimeFrame *big.Int, targetContract common.Address, oldData []byte, newJobName string, newTimeFrame *big.Int, newData []byte) (*types.Transaction, error)

	UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error)
}

// ContractJobRegistryFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractJobRegistryFilters interface {
	FilterInitialized(opts *bind.FilterOpts) (*ContractJobRegistryInitializedIterator, error)
	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractJobRegistryInitialized) (event.Subscription, error)
	ParseInitialized(log types.Log) (*ContractJobRegistryInitialized, error)

	FilterJobCreated(opts *bind.FilterOpts, jobId []*big.Int, jobOwner []common.Address) (*ContractJobRegistryJobCreatedIterator, error)
	WatchJobCreated(opts *bind.WatchOpts, sink chan<- *ContractJobRegistryJobCreated, jobId []*big.Int, jobOwner []common.Address) (event.Subscription, error)
	ParseJobCreated(log types.Log) (*ContractJobRegistryJobCreated, error)

	FilterJobDeleted(opts *bind.FilterOpts, jobId []*big.Int, jobOwner []common.Address) (*ContractJobRegistryJobDeletedIterator, error)
	WatchJobDeleted(opts *bind.WatchOpts, sink chan<- *ContractJobRegistryJobDeleted, jobId []*big.Int, jobOwner []common.Address) (event.Subscription, error)
	ParseJobDeleted(log types.Log) (*ContractJobRegistryJobDeleted, error)

	FilterJobUpdated(opts *bind.FilterOpts, jobId []*big.Int, jobOwner []common.Address) (*ContractJobRegistryJobUpdatedIterator, error)
	WatchJobUpdated(opts *bind.WatchOpts, sink chan<- *ContractJobRegistryJobUpdated, jobId []*big.Int, jobOwner []common.Address) (event.Subscription, error)
	ParseJobUpdated(log types.Log) (*ContractJobRegistryJobUpdated, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractJobRegistryOwnershipTransferredIterator, error)
	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractJobRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error)
	ParseOwnershipTransferred(log types.Log) (*ContractJobRegistryOwnershipTransferred, error)

	FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ContractJobRegistryUpgradedIterator, error)
	WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ContractJobRegistryUpgraded, implementation []common.Address) (event.Subscription, error)
	ParseUpgraded(log types.Log) (*ContractJobRegistryUpgraded, error)
}

// ContractJobRegistry is an auto generated Go binding around an Ethereum contract.
type ContractJobRegistry struct {
	ContractJobRegistryCaller     // Read-only binding to the contract
	ContractJobRegistryTransactor // Write-only binding to the contract
	ContractJobRegistryFilterer   // Log filterer for contract events
}

// ContractJobRegistry implements the ContractJobRegistryMethods interface.
var _ ContractJobRegistryMethods = (*ContractJobRegistry)(nil)

// ContractJobRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractJobRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractJobRegistryCaller implements the ContractJobRegistryCalls interface.
var _ ContractJobRegistryCalls = (*ContractJobRegistryCaller)(nil)

// ContractJobRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractJobRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractJobRegistryTransactor implements the ContractJobRegistryTransacts interface.
var _ ContractJobRegistryTransacts = (*ContractJobRegistryTransactor)(nil)

// ContractJobRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractJobRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractJobRegistryFilterer implements the ContractJobRegistryFilters interface.
var _ ContractJobRegistryFilters = (*ContractJobRegistryFilterer)(nil)

// ContractJobRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractJobRegistrySession struct {
	Contract     *ContractJobRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractJobRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractJobRegistryCallerSession struct {
	Contract *ContractJobRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ContractJobRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractJobRegistryTransactorSession struct {
	Contract     *ContractJobRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ContractJobRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractJobRegistryRaw struct {
	Contract *ContractJobRegistry // Generic contract binding to access the raw methods on
}

// ContractJobRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractJobRegistryCallerRaw struct {
	Contract *ContractJobRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ContractJobRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractJobRegistryTransactorRaw struct {
	Contract *ContractJobRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractJobRegistry creates a new instance of ContractJobRegistry, bound to a specific deployed contract.
func NewContractJobRegistry(address common.Address, backend bind.ContractBackend) (*ContractJobRegistry, error) {
	contract, err := bindContractJobRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractJobRegistry{ContractJobRegistryCaller: ContractJobRegistryCaller{contract: contract}, ContractJobRegistryTransactor: ContractJobRegistryTransactor{contract: contract}, ContractJobRegistryFilterer: ContractJobRegistryFilterer{contract: contract}}, nil
}

// NewContractJobRegistryCaller creates a new read-only instance of ContractJobRegistry, bound to a specific deployed contract.
func NewContractJobRegistryCaller(address common.Address, caller bind.ContractCaller) (*ContractJobRegistryCaller, error) {
	contract, err := bindContractJobRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractJobRegistryCaller{contract: contract}, nil
}

// NewContractJobRegistryTransactor creates a new write-only instance of ContractJobRegistry, bound to a specific deployed contract.
func NewContractJobRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractJobRegistryTransactor, error) {
	contract, err := bindContractJobRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractJobRegistryTransactor{contract: contract}, nil
}

// NewContractJobRegistryFilterer creates a new log filterer instance of ContractJobRegistry, bound to a specific deployed contract.
func NewContractJobRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractJobRegistryFilterer, error) {
	contract, err := bindContractJobRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractJobRegistryFilterer{contract: contract}, nil
}

// bindContractJobRegistry binds a generic wrapper to an already deployed contract.
func bindContractJobRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractJobRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractJobRegistry *ContractJobRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractJobRegistry.Contract.ContractJobRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractJobRegistry *ContractJobRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractJobRegistry.Contract.ContractJobRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractJobRegistry *ContractJobRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractJobRegistry.Contract.ContractJobRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractJobRegistry *ContractJobRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractJobRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractJobRegistry *ContractJobRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractJobRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractJobRegistry *ContractJobRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractJobRegistry.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ContractJobRegistry *ContractJobRegistryCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ContractJobRegistry.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ContractJobRegistry *ContractJobRegistrySession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ContractJobRegistry.Contract.UPGRADEINTERFACEVERSION(&_ContractJobRegistry.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ContractJobRegistry *ContractJobRegistryCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ContractJobRegistry.Contract.UPGRADEINTERFACEVERSION(&_ContractJobRegistry.CallOpts)
}

// GetJob is a free data retrieval call binding the contract method 0xbf22c457.
//
// Solidity: function getJob(uint256 jobId) view returns((uint256,address,bytes32,uint256,bool) job)
func (_ContractJobRegistry *ContractJobRegistryCaller) GetJob(opts *bind.CallOpts, jobId *big.Int) (JobRegistryJob, error) {
	var out []interface{}
	err := _ContractJobRegistry.contract.Call(opts, &out, "getJob", jobId)

	if err != nil {
		return *new(JobRegistryJob), err
	}

	out0 := *abi.ConvertType(out[0], new(JobRegistryJob)).(*JobRegistryJob)

	return out0, err

}

// GetJob is a free data retrieval call binding the contract method 0xbf22c457.
//
// Solidity: function getJob(uint256 jobId) view returns((uint256,address,bytes32,uint256,bool) job)
func (_ContractJobRegistry *ContractJobRegistrySession) GetJob(jobId *big.Int) (JobRegistryJob, error) {
	return _ContractJobRegistry.Contract.GetJob(&_ContractJobRegistry.CallOpts, jobId)
}

// GetJob is a free data retrieval call binding the contract method 0xbf22c457.
//
// Solidity: function getJob(uint256 jobId) view returns((uint256,address,bytes32,uint256,bool) job)
func (_ContractJobRegistry *ContractJobRegistryCallerSession) GetJob(jobId *big.Int) (JobRegistryJob, error) {
	return _ContractJobRegistry.Contract.GetJob(&_ContractJobRegistry.CallOpts, jobId)
}

// GetJobByCounter is a free data retrieval call binding the contract method 0xef178814.
//
// Solidity: function getJobByCounter(uint256 jobCounter) view returns((uint256,address,bytes32,uint256,bool) job)
func (_ContractJobRegistry *ContractJobRegistryCaller) GetJobByCounter(opts *bind.CallOpts, jobCounter *big.Int) (JobRegistryJob, error) {
	var out []interface{}
	err := _ContractJobRegistry.contract.Call(opts, &out, "getJobByCounter", jobCounter)

	if err != nil {
		return *new(JobRegistryJob), err
	}

	out0 := *abi.ConvertType(out[0], new(JobRegistryJob)).(*JobRegistryJob)

	return out0, err

}

// GetJobByCounter is a free data retrieval call binding the contract method 0xef178814.
//
// Solidity: function getJobByCounter(uint256 jobCounter) view returns((uint256,address,bytes32,uint256,bool) job)
func (_ContractJobRegistry *ContractJobRegistrySession) GetJobByCounter(jobCounter *big.Int) (JobRegistryJob, error) {
	return _ContractJobRegistry.Contract.GetJobByCounter(&_ContractJobRegistry.CallOpts, jobCounter)
}

// GetJobByCounter is a free data retrieval call binding the contract method 0xef178814.
//
// Solidity: function getJobByCounter(uint256 jobCounter) view returns((uint256,address,bytes32,uint256,bool) job)
func (_ContractJobRegistry *ContractJobRegistryCallerSession) GetJobByCounter(jobCounter *big.Int) (JobRegistryJob, error) {
	return _ContractJobRegistry.Contract.GetJobByCounter(&_ContractJobRegistry.CallOpts, jobCounter)
}

// GetJobCounter is a free data retrieval call binding the contract method 0x2257b8da.
//
// Solidity: function getJobCounter() view returns(uint256 count)
func (_ContractJobRegistry *ContractJobRegistryCaller) GetJobCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractJobRegistry.contract.Call(opts, &out, "getJobCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetJobCounter is a free data retrieval call binding the contract method 0x2257b8da.
//
// Solidity: function getJobCounter() view returns(uint256 count)
func (_ContractJobRegistry *ContractJobRegistrySession) GetJobCounter() (*big.Int, error) {
	return _ContractJobRegistry.Contract.GetJobCounter(&_ContractJobRegistry.CallOpts)
}

// GetJobCounter is a free data retrieval call binding the contract method 0x2257b8da.
//
// Solidity: function getJobCounter() view returns(uint256 count)
func (_ContractJobRegistry *ContractJobRegistryCallerSession) GetJobCounter() (*big.Int, error) {
	return _ContractJobRegistry.Contract.GetJobCounter(&_ContractJobRegistry.CallOpts)
}

// GetJobOwner is a free data retrieval call binding the contract method 0x147f5511.
//
// Solidity: function getJobOwner(uint256 jobId) view returns(address)
func (_ContractJobRegistry *ContractJobRegistryCaller) GetJobOwner(opts *bind.CallOpts, jobId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ContractJobRegistry.contract.Call(opts, &out, "getJobOwner", jobId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetJobOwner is a free data retrieval call binding the contract method 0x147f5511.
//
// Solidity: function getJobOwner(uint256 jobId) view returns(address)
func (_ContractJobRegistry *ContractJobRegistrySession) GetJobOwner(jobId *big.Int) (common.Address, error) {
	return _ContractJobRegistry.Contract.GetJobOwner(&_ContractJobRegistry.CallOpts, jobId)
}

// GetJobOwner is a free data retrieval call binding the contract method 0x147f5511.
//
// Solidity: function getJobOwner(uint256 jobId) view returns(address)
func (_ContractJobRegistry *ContractJobRegistryCallerSession) GetJobOwner(jobId *big.Int) (common.Address, error) {
	return _ContractJobRegistry.Contract.GetJobOwner(&_ContractJobRegistry.CallOpts, jobId)
}

// GetUserActiveJobIds is a free data retrieval call binding the contract method 0xff7b49a0.
//
// Solidity: function getUserActiveJobIds(address user) view returns(uint256[] activeJobIds)
func (_ContractJobRegistry *ContractJobRegistryCaller) GetUserActiveJobIds(opts *bind.CallOpts, user common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _ContractJobRegistry.contract.Call(opts, &out, "getUserActiveJobIds", user)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetUserActiveJobIds is a free data retrieval call binding the contract method 0xff7b49a0.
//
// Solidity: function getUserActiveJobIds(address user) view returns(uint256[] activeJobIds)
func (_ContractJobRegistry *ContractJobRegistrySession) GetUserActiveJobIds(user common.Address) ([]*big.Int, error) {
	return _ContractJobRegistry.Contract.GetUserActiveJobIds(&_ContractJobRegistry.CallOpts, user)
}

// GetUserActiveJobIds is a free data retrieval call binding the contract method 0xff7b49a0.
//
// Solidity: function getUserActiveJobIds(address user) view returns(uint256[] activeJobIds)
func (_ContractJobRegistry *ContractJobRegistryCallerSession) GetUserActiveJobIds(user common.Address) ([]*big.Int, error) {
	return _ContractJobRegistry.Contract.GetUserActiveJobIds(&_ContractJobRegistry.CallOpts, user)
}

// GetUserJobIds is a free data retrieval call binding the contract method 0x5523d8dc.
//
// Solidity: function getUserJobIds(address user) view returns(uint256[] jobIds)
func (_ContractJobRegistry *ContractJobRegistryCaller) GetUserJobIds(opts *bind.CallOpts, user common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _ContractJobRegistry.contract.Call(opts, &out, "getUserJobIds", user)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetUserJobIds is a free data retrieval call binding the contract method 0x5523d8dc.
//
// Solidity: function getUserJobIds(address user) view returns(uint256[] jobIds)
func (_ContractJobRegistry *ContractJobRegistrySession) GetUserJobIds(user common.Address) ([]*big.Int, error) {
	return _ContractJobRegistry.Contract.GetUserJobIds(&_ContractJobRegistry.CallOpts, user)
}

// GetUserJobIds is a free data retrieval call binding the contract method 0x5523d8dc.
//
// Solidity: function getUserJobIds(address user) view returns(uint256[] jobIds)
func (_ContractJobRegistry *ContractJobRegistryCallerSession) GetUserJobIds(user common.Address) ([]*big.Int, error) {
	return _ContractJobRegistry.Contract.GetUserJobIds(&_ContractJobRegistry.CallOpts, user)
}

// IsJobActive is a free data retrieval call binding the contract method 0x53bc234f.
//
// Solidity: function isJobActive(uint256 jobId) view returns(bool isActive)
func (_ContractJobRegistry *ContractJobRegistryCaller) IsJobActive(opts *bind.CallOpts, jobId *big.Int) (bool, error) {
	var out []interface{}
	err := _ContractJobRegistry.contract.Call(opts, &out, "isJobActive", jobId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsJobActive is a free data retrieval call binding the contract method 0x53bc234f.
//
// Solidity: function isJobActive(uint256 jobId) view returns(bool isActive)
func (_ContractJobRegistry *ContractJobRegistrySession) IsJobActive(jobId *big.Int) (bool, error) {
	return _ContractJobRegistry.Contract.IsJobActive(&_ContractJobRegistry.CallOpts, jobId)
}

// IsJobActive is a free data retrieval call binding the contract method 0x53bc234f.
//
// Solidity: function isJobActive(uint256 jobId) view returns(bool isActive)
func (_ContractJobRegistry *ContractJobRegistryCallerSession) IsJobActive(jobId *big.Int) (bool, error) {
	return _ContractJobRegistry.Contract.IsJobActive(&_ContractJobRegistry.CallOpts, jobId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractJobRegistry *ContractJobRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractJobRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractJobRegistry *ContractJobRegistrySession) Owner() (common.Address, error) {
	return _ContractJobRegistry.Contract.Owner(&_ContractJobRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractJobRegistry *ContractJobRegistryCallerSession) Owner() (common.Address, error) {
	return _ContractJobRegistry.Contract.Owner(&_ContractJobRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ContractJobRegistry *ContractJobRegistryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractJobRegistry.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ContractJobRegistry *ContractJobRegistrySession) ProxiableUUID() ([32]byte, error) {
	return _ContractJobRegistry.Contract.ProxiableUUID(&_ContractJobRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ContractJobRegistry *ContractJobRegistryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _ContractJobRegistry.Contract.ProxiableUUID(&_ContractJobRegistry.CallOpts)
}

// UnpackJobId is a free data retrieval call binding the contract method 0xbf8e9d53.
//
// Solidity: function unpackJobId(uint256 jobId) pure returns(uint256 chainId, uint256 timestamp, uint256 jobCounter)
func (_ContractJobRegistry *ContractJobRegistryCaller) UnpackJobId(opts *bind.CallOpts, jobId *big.Int) (struct {
	ChainId    *big.Int
	Timestamp  *big.Int
	JobCounter *big.Int
}, error) {
	var out []interface{}
	err := _ContractJobRegistry.contract.Call(opts, &out, "unpackJobId", jobId)

	outstruct := new(struct {
		ChainId    *big.Int
		Timestamp  *big.Int
		JobCounter *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ChainId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.JobCounter = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UnpackJobId is a free data retrieval call binding the contract method 0xbf8e9d53.
//
// Solidity: function unpackJobId(uint256 jobId) pure returns(uint256 chainId, uint256 timestamp, uint256 jobCounter)
func (_ContractJobRegistry *ContractJobRegistrySession) UnpackJobId(jobId *big.Int) (struct {
	ChainId    *big.Int
	Timestamp  *big.Int
	JobCounter *big.Int
}, error) {
	return _ContractJobRegistry.Contract.UnpackJobId(&_ContractJobRegistry.CallOpts, jobId)
}

// UnpackJobId is a free data retrieval call binding the contract method 0xbf8e9d53.
//
// Solidity: function unpackJobId(uint256 jobId) pure returns(uint256 chainId, uint256 timestamp, uint256 jobCounter)
func (_ContractJobRegistry *ContractJobRegistryCallerSession) UnpackJobId(jobId *big.Int) (struct {
	ChainId    *big.Int
	Timestamp  *big.Int
	JobCounter *big.Int
}, error) {
	return _ContractJobRegistry.Contract.UnpackJobId(&_ContractJobRegistry.CallOpts, jobId)
}

// CreateJob is a paid mutator transaction binding the contract method 0x893de3c5.
//
// Solidity: function createJob(string jobName, uint8 jobType, uint256 timeFrame, address targetContract, bytes data) returns(uint256 jobId)
func (_ContractJobRegistry *ContractJobRegistryTransactor) CreateJob(opts *bind.TransactOpts, jobName string, jobType uint8, timeFrame *big.Int, targetContract common.Address, data []byte) (*types.Transaction, error) {
	return _ContractJobRegistry.contract.Transact(opts, "createJob", jobName, jobType, timeFrame, targetContract, data)
}

// CreateJob is a paid mutator transaction binding the contract method 0x893de3c5.
//
// Solidity: function createJob(string jobName, uint8 jobType, uint256 timeFrame, address targetContract, bytes data) returns(uint256 jobId)
func (_ContractJobRegistry *ContractJobRegistrySession) CreateJob(jobName string, jobType uint8, timeFrame *big.Int, targetContract common.Address, data []byte) (*types.Transaction, error) {
	return _ContractJobRegistry.Contract.CreateJob(&_ContractJobRegistry.TransactOpts, jobName, jobType, timeFrame, targetContract, data)
}

// CreateJob is a paid mutator transaction binding the contract method 0x893de3c5.
//
// Solidity: function createJob(string jobName, uint8 jobType, uint256 timeFrame, address targetContract, bytes data) returns(uint256 jobId)
func (_ContractJobRegistry *ContractJobRegistryTransactorSession) CreateJob(jobName string, jobType uint8, timeFrame *big.Int, targetContract common.Address, data []byte) (*types.Transaction, error) {
	return _ContractJobRegistry.Contract.CreateJob(&_ContractJobRegistry.TransactOpts, jobName, jobType, timeFrame, targetContract, data)
}

// DeleteJob is a paid mutator transaction binding the contract method 0x6358bb8c.
//
// Solidity: function deleteJob(uint256 jobId) returns()
func (_ContractJobRegistry *ContractJobRegistryTransactor) DeleteJob(opts *bind.TransactOpts, jobId *big.Int) (*types.Transaction, error) {
	return _ContractJobRegistry.contract.Transact(opts, "deleteJob", jobId)
}

// DeleteJob is a paid mutator transaction binding the contract method 0x6358bb8c.
//
// Solidity: function deleteJob(uint256 jobId) returns()
func (_ContractJobRegistry *ContractJobRegistrySession) DeleteJob(jobId *big.Int) (*types.Transaction, error) {
	return _ContractJobRegistry.Contract.DeleteJob(&_ContractJobRegistry.TransactOpts, jobId)
}

// DeleteJob is a paid mutator transaction binding the contract method 0x6358bb8c.
//
// Solidity: function deleteJob(uint256 jobId) returns()
func (_ContractJobRegistry *ContractJobRegistryTransactorSession) DeleteJob(jobId *big.Int) (*types.Transaction, error) {
	return _ContractJobRegistry.Contract.DeleteJob(&_ContractJobRegistry.TransactOpts, jobId)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_ContractJobRegistry *ContractJobRegistryTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address) (*types.Transaction, error) {
	return _ContractJobRegistry.contract.Transact(opts, "initialize", initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_ContractJobRegistry *ContractJobRegistrySession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _ContractJobRegistry.Contract.Initialize(&_ContractJobRegistry.TransactOpts, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_ContractJobRegistry *ContractJobRegistryTransactorSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _ContractJobRegistry.Contract.Initialize(&_ContractJobRegistry.TransactOpts, initialOwner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractJobRegistry *ContractJobRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractJobRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractJobRegistry *ContractJobRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractJobRegistry.Contract.RenounceOwnership(&_ContractJobRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractJobRegistry *ContractJobRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractJobRegistry.Contract.RenounceOwnership(&_ContractJobRegistry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractJobRegistry *ContractJobRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractJobRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractJobRegistry *ContractJobRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractJobRegistry.Contract.TransferOwnership(&_ContractJobRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractJobRegistry *ContractJobRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractJobRegistry.Contract.TransferOwnership(&_ContractJobRegistry.TransactOpts, newOwner)
}

// UpdateJob is a paid mutator transaction binding the contract method 0x91461d0a.
//
// Solidity: function updateJob(uint256 jobId, string oldJobName, uint8 jobType, uint256 oldTimeFrame, address targetContract, bytes oldData, string newJobName, uint256 newTimeFrame, bytes newData) returns()
func (_ContractJobRegistry *ContractJobRegistryTransactor) UpdateJob(opts *bind.TransactOpts, jobId *big.Int, oldJobName string, jobType uint8, oldTimeFrame *big.Int, targetContract common.Address, oldData []byte, newJobName string, newTimeFrame *big.Int, newData []byte) (*types.Transaction, error) {
	return _ContractJobRegistry.contract.Transact(opts, "updateJob", jobId, oldJobName, jobType, oldTimeFrame, targetContract, oldData, newJobName, newTimeFrame, newData)
}

// UpdateJob is a paid mutator transaction binding the contract method 0x91461d0a.
//
// Solidity: function updateJob(uint256 jobId, string oldJobName, uint8 jobType, uint256 oldTimeFrame, address targetContract, bytes oldData, string newJobName, uint256 newTimeFrame, bytes newData) returns()
func (_ContractJobRegistry *ContractJobRegistrySession) UpdateJob(jobId *big.Int, oldJobName string, jobType uint8, oldTimeFrame *big.Int, targetContract common.Address, oldData []byte, newJobName string, newTimeFrame *big.Int, newData []byte) (*types.Transaction, error) {
	return _ContractJobRegistry.Contract.UpdateJob(&_ContractJobRegistry.TransactOpts, jobId, oldJobName, jobType, oldTimeFrame, targetContract, oldData, newJobName, newTimeFrame, newData)
}

// UpdateJob is a paid mutator transaction binding the contract method 0x91461d0a.
//
// Solidity: function updateJob(uint256 jobId, string oldJobName, uint8 jobType, uint256 oldTimeFrame, address targetContract, bytes oldData, string newJobName, uint256 newTimeFrame, bytes newData) returns()
func (_ContractJobRegistry *ContractJobRegistryTransactorSession) UpdateJob(jobId *big.Int, oldJobName string, jobType uint8, oldTimeFrame *big.Int, targetContract common.Address, oldData []byte, newJobName string, newTimeFrame *big.Int, newData []byte) (*types.Transaction, error) {
	return _ContractJobRegistry.Contract.UpdateJob(&_ContractJobRegistry.TransactOpts, jobId, oldJobName, jobType, oldTimeFrame, targetContract, oldData, newJobName, newTimeFrame, newData)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ContractJobRegistry *ContractJobRegistryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ContractJobRegistry.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ContractJobRegistry *ContractJobRegistrySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ContractJobRegistry.Contract.UpgradeToAndCall(&_ContractJobRegistry.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ContractJobRegistry *ContractJobRegistryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ContractJobRegistry.Contract.UpgradeToAndCall(&_ContractJobRegistry.TransactOpts, newImplementation, data)
}

// ContractJobRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractJobRegistry contract.
type ContractJobRegistryInitializedIterator struct {
	Event *ContractJobRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *ContractJobRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractJobRegistryInitialized)
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
		it.Event = new(ContractJobRegistryInitialized)
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
func (it *ContractJobRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractJobRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractJobRegistryInitialized represents a Initialized event raised by the ContractJobRegistry contract.
type ContractJobRegistryInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ContractJobRegistry *ContractJobRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractJobRegistryInitializedIterator, error) {

	logs, sub, err := _ContractJobRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractJobRegistryInitializedIterator{contract: _ContractJobRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ContractJobRegistry *ContractJobRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractJobRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractJobRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractJobRegistryInitialized)
				if err := _ContractJobRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractJobRegistry *ContractJobRegistryFilterer) ParseInitialized(log types.Log) (*ContractJobRegistryInitialized, error) {
	event := new(ContractJobRegistryInitialized)
	if err := _ContractJobRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractJobRegistryJobCreatedIterator is returned from FilterJobCreated and is used to iterate over the raw logs and unpacked data for JobCreated events raised by the ContractJobRegistry contract.
type ContractJobRegistryJobCreatedIterator struct {
	Event *ContractJobRegistryJobCreated // Event containing the contract specifics and raw log

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
func (it *ContractJobRegistryJobCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractJobRegistryJobCreated)
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
		it.Event = new(ContractJobRegistryJobCreated)
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
func (it *ContractJobRegistryJobCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractJobRegistryJobCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractJobRegistryJobCreated represents a JobCreated event raised by the ContractJobRegistry contract.
type ContractJobRegistryJobCreated struct {
	JobId     *big.Int
	JobOwner  common.Address
	JobHash   [32]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterJobCreated is a free log retrieval operation binding the contract event 0x737fc62fbb05dd9fb7c799e68796a2d7c8324e310af7b656c0580e7b3cf8bf8a.
//
// Solidity: event JobCreated(uint256 indexed jobId, address indexed jobOwner, bytes32 jobHash, uint256 timestamp)
func (_ContractJobRegistry *ContractJobRegistryFilterer) FilterJobCreated(opts *bind.FilterOpts, jobId []*big.Int, jobOwner []common.Address) (*ContractJobRegistryJobCreatedIterator, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var jobOwnerRule []interface{}
	for _, jobOwnerItem := range jobOwner {
		jobOwnerRule = append(jobOwnerRule, jobOwnerItem)
	}

	logs, sub, err := _ContractJobRegistry.contract.FilterLogs(opts, "JobCreated", jobIdRule, jobOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractJobRegistryJobCreatedIterator{contract: _ContractJobRegistry.contract, event: "JobCreated", logs: logs, sub: sub}, nil
}

// WatchJobCreated is a free log subscription operation binding the contract event 0x737fc62fbb05dd9fb7c799e68796a2d7c8324e310af7b656c0580e7b3cf8bf8a.
//
// Solidity: event JobCreated(uint256 indexed jobId, address indexed jobOwner, bytes32 jobHash, uint256 timestamp)
func (_ContractJobRegistry *ContractJobRegistryFilterer) WatchJobCreated(opts *bind.WatchOpts, sink chan<- *ContractJobRegistryJobCreated, jobId []*big.Int, jobOwner []common.Address) (event.Subscription, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var jobOwnerRule []interface{}
	for _, jobOwnerItem := range jobOwner {
		jobOwnerRule = append(jobOwnerRule, jobOwnerItem)
	}

	logs, sub, err := _ContractJobRegistry.contract.WatchLogs(opts, "JobCreated", jobIdRule, jobOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractJobRegistryJobCreated)
				if err := _ContractJobRegistry.contract.UnpackLog(event, "JobCreated", log); err != nil {
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

// ParseJobCreated is a log parse operation binding the contract event 0x737fc62fbb05dd9fb7c799e68796a2d7c8324e310af7b656c0580e7b3cf8bf8a.
//
// Solidity: event JobCreated(uint256 indexed jobId, address indexed jobOwner, bytes32 jobHash, uint256 timestamp)
func (_ContractJobRegistry *ContractJobRegistryFilterer) ParseJobCreated(log types.Log) (*ContractJobRegistryJobCreated, error) {
	event := new(ContractJobRegistryJobCreated)
	if err := _ContractJobRegistry.contract.UnpackLog(event, "JobCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractJobRegistryJobDeletedIterator is returned from FilterJobDeleted and is used to iterate over the raw logs and unpacked data for JobDeleted events raised by the ContractJobRegistry contract.
type ContractJobRegistryJobDeletedIterator struct {
	Event *ContractJobRegistryJobDeleted // Event containing the contract specifics and raw log

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
func (it *ContractJobRegistryJobDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractJobRegistryJobDeleted)
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
		it.Event = new(ContractJobRegistryJobDeleted)
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
func (it *ContractJobRegistryJobDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractJobRegistryJobDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractJobRegistryJobDeleted represents a JobDeleted event raised by the ContractJobRegistry contract.
type ContractJobRegistryJobDeleted struct {
	JobId     *big.Int
	JobOwner  common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterJobDeleted is a free log retrieval operation binding the contract event 0x8bc0131dae2d0cbc363614dfcf9653c32125d8fe69ea2cc1f6be9bc430198537.
//
// Solidity: event JobDeleted(uint256 indexed jobId, address indexed jobOwner, uint256 timestamp)
func (_ContractJobRegistry *ContractJobRegistryFilterer) FilterJobDeleted(opts *bind.FilterOpts, jobId []*big.Int, jobOwner []common.Address) (*ContractJobRegistryJobDeletedIterator, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var jobOwnerRule []interface{}
	for _, jobOwnerItem := range jobOwner {
		jobOwnerRule = append(jobOwnerRule, jobOwnerItem)
	}

	logs, sub, err := _ContractJobRegistry.contract.FilterLogs(opts, "JobDeleted", jobIdRule, jobOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractJobRegistryJobDeletedIterator{contract: _ContractJobRegistry.contract, event: "JobDeleted", logs: logs, sub: sub}, nil
}

// WatchJobDeleted is a free log subscription operation binding the contract event 0x8bc0131dae2d0cbc363614dfcf9653c32125d8fe69ea2cc1f6be9bc430198537.
//
// Solidity: event JobDeleted(uint256 indexed jobId, address indexed jobOwner, uint256 timestamp)
func (_ContractJobRegistry *ContractJobRegistryFilterer) WatchJobDeleted(opts *bind.WatchOpts, sink chan<- *ContractJobRegistryJobDeleted, jobId []*big.Int, jobOwner []common.Address) (event.Subscription, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var jobOwnerRule []interface{}
	for _, jobOwnerItem := range jobOwner {
		jobOwnerRule = append(jobOwnerRule, jobOwnerItem)
	}

	logs, sub, err := _ContractJobRegistry.contract.WatchLogs(opts, "JobDeleted", jobIdRule, jobOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractJobRegistryJobDeleted)
				if err := _ContractJobRegistry.contract.UnpackLog(event, "JobDeleted", log); err != nil {
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

// ParseJobDeleted is a log parse operation binding the contract event 0x8bc0131dae2d0cbc363614dfcf9653c32125d8fe69ea2cc1f6be9bc430198537.
//
// Solidity: event JobDeleted(uint256 indexed jobId, address indexed jobOwner, uint256 timestamp)
func (_ContractJobRegistry *ContractJobRegistryFilterer) ParseJobDeleted(log types.Log) (*ContractJobRegistryJobDeleted, error) {
	event := new(ContractJobRegistryJobDeleted)
	if err := _ContractJobRegistry.contract.UnpackLog(event, "JobDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractJobRegistryJobUpdatedIterator is returned from FilterJobUpdated and is used to iterate over the raw logs and unpacked data for JobUpdated events raised by the ContractJobRegistry contract.
type ContractJobRegistryJobUpdatedIterator struct {
	Event *ContractJobRegistryJobUpdated // Event containing the contract specifics and raw log

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
func (it *ContractJobRegistryJobUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractJobRegistryJobUpdated)
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
		it.Event = new(ContractJobRegistryJobUpdated)
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
func (it *ContractJobRegistryJobUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractJobRegistryJobUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractJobRegistryJobUpdated represents a JobUpdated event raised by the ContractJobRegistry contract.
type ContractJobRegistryJobUpdated struct {
	JobId      *big.Int
	JobOwner   common.Address
	NewJobHash [32]byte
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterJobUpdated is a free log retrieval operation binding the contract event 0x7dac5d7787afd94b2012ccf48531da941b75ac7e9cfc8b6ea0b55cb3a93dbd13.
//
// Solidity: event JobUpdated(uint256 indexed jobId, address indexed jobOwner, bytes32 newJobHash, uint256 timestamp)
func (_ContractJobRegistry *ContractJobRegistryFilterer) FilterJobUpdated(opts *bind.FilterOpts, jobId []*big.Int, jobOwner []common.Address) (*ContractJobRegistryJobUpdatedIterator, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var jobOwnerRule []interface{}
	for _, jobOwnerItem := range jobOwner {
		jobOwnerRule = append(jobOwnerRule, jobOwnerItem)
	}

	logs, sub, err := _ContractJobRegistry.contract.FilterLogs(opts, "JobUpdated", jobIdRule, jobOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractJobRegistryJobUpdatedIterator{contract: _ContractJobRegistry.contract, event: "JobUpdated", logs: logs, sub: sub}, nil
}

// WatchJobUpdated is a free log subscription operation binding the contract event 0x7dac5d7787afd94b2012ccf48531da941b75ac7e9cfc8b6ea0b55cb3a93dbd13.
//
// Solidity: event JobUpdated(uint256 indexed jobId, address indexed jobOwner, bytes32 newJobHash, uint256 timestamp)
func (_ContractJobRegistry *ContractJobRegistryFilterer) WatchJobUpdated(opts *bind.WatchOpts, sink chan<- *ContractJobRegistryJobUpdated, jobId []*big.Int, jobOwner []common.Address) (event.Subscription, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var jobOwnerRule []interface{}
	for _, jobOwnerItem := range jobOwner {
		jobOwnerRule = append(jobOwnerRule, jobOwnerItem)
	}

	logs, sub, err := _ContractJobRegistry.contract.WatchLogs(opts, "JobUpdated", jobIdRule, jobOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractJobRegistryJobUpdated)
				if err := _ContractJobRegistry.contract.UnpackLog(event, "JobUpdated", log); err != nil {
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

// ParseJobUpdated is a log parse operation binding the contract event 0x7dac5d7787afd94b2012ccf48531da941b75ac7e9cfc8b6ea0b55cb3a93dbd13.
//
// Solidity: event JobUpdated(uint256 indexed jobId, address indexed jobOwner, bytes32 newJobHash, uint256 timestamp)
func (_ContractJobRegistry *ContractJobRegistryFilterer) ParseJobUpdated(log types.Log) (*ContractJobRegistryJobUpdated, error) {
	event := new(ContractJobRegistryJobUpdated)
	if err := _ContractJobRegistry.contract.UnpackLog(event, "JobUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractJobRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractJobRegistry contract.
type ContractJobRegistryOwnershipTransferredIterator struct {
	Event *ContractJobRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractJobRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractJobRegistryOwnershipTransferred)
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
		it.Event = new(ContractJobRegistryOwnershipTransferred)
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
func (it *ContractJobRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractJobRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractJobRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the ContractJobRegistry contract.
type ContractJobRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractJobRegistry *ContractJobRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractJobRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractJobRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractJobRegistryOwnershipTransferredIterator{contract: _ContractJobRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractJobRegistry *ContractJobRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractJobRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractJobRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractJobRegistryOwnershipTransferred)
				if err := _ContractJobRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractJobRegistry *ContractJobRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*ContractJobRegistryOwnershipTransferred, error) {
	event := new(ContractJobRegistryOwnershipTransferred)
	if err := _ContractJobRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractJobRegistryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ContractJobRegistry contract.
type ContractJobRegistryUpgradedIterator struct {
	Event *ContractJobRegistryUpgraded // Event containing the contract specifics and raw log

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
func (it *ContractJobRegistryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractJobRegistryUpgraded)
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
		it.Event = new(ContractJobRegistryUpgraded)
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
func (it *ContractJobRegistryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractJobRegistryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractJobRegistryUpgraded represents a Upgraded event raised by the ContractJobRegistry contract.
type ContractJobRegistryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ContractJobRegistry *ContractJobRegistryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ContractJobRegistryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ContractJobRegistry.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ContractJobRegistryUpgradedIterator{contract: _ContractJobRegistry.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ContractJobRegistry *ContractJobRegistryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ContractJobRegistryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ContractJobRegistry.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractJobRegistryUpgraded)
				if err := _ContractJobRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_ContractJobRegistry *ContractJobRegistryFilterer) ParseUpgraded(log types.Log) (*ContractJobRegistryUpgraded, error) {
	event := new(ContractJobRegistryUpgraded)
	if err := _ContractJobRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
