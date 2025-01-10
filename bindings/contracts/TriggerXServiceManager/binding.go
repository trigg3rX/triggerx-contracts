// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractTriggerXServiceManager

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

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// IBLSSignatureCheckerNonSignerStakesAndSignature is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerNonSignerStakesAndSignature struct {
	NonSignerQuorumBitmapIndices []uint32
	NonSignerPubkeys             []BN254G1Point
	QuorumApks                   []BN254G1Point
	ApkG2                        BN254G2Point
	Sigma                        BN254G1Point
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// IBLSSignatureCheckerQuorumStakeTotals is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerQuorumStakeTotals struct {
	SignedStakeForQuorum []*big.Int
	TotalStakeForQuorum  []*big.Int
}

// IRewardsCoordinatorOperatorDirectedRewardsSubmission is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorOperatorDirectedRewardsSubmission struct {
	StrategiesAndMultipliers []IRewardsCoordinatorStrategyAndMultiplier
	Token                    common.Address
	OperatorRewards          []IRewardsCoordinatorOperatorReward
	StartTimestamp           uint32
	Duration                 uint32
	Description              string
}

// IRewardsCoordinatorOperatorReward is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorOperatorReward struct {
	Operator common.Address
	Amount   *big.Int
}

// IRewardsCoordinatorRewardsSubmission is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorRewardsSubmission struct {
	StrategiesAndMultipliers []IRewardsCoordinatorStrategyAndMultiplier
	Token                    common.Address
	Amount                   *big.Int
	StartTimestamp           uint32
	Duration                 uint32
}

// IRewardsCoordinatorStrategyAndMultiplier is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorStrategyAndMultiplier struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// ContractTriggerXServiceManagerMetaData contains all meta data concerning the ContractTriggerXServiceManager contract.
var ContractTriggerXServiceManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIAVSDirectory\",\"name\":\"_avsDirectory\",\"type\":\"address\"},{\"internalType\":\"contractIRewardsCoordinator\",\"name\":\"_rewardsCoordinator\",\"type\":\"address\"},{\"internalType\":\"contractIRegistryCoordinator\",\"name\":\"_registryCoordinator\",\"type\":\"address\"},{\"internalType\":\"contractIStakeRegistry\",\"name\":\"_stakeRegistry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"KeeperAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"KeeperBlacklisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"KeeperRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"KeeperUnblacklisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIPauserRegistry\",\"name\":\"pauserRegistry\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIPauserRegistry\",\"name\":\"newPauserRegistry\",\"type\":\"address\"}],\"name\":\"PauserRegistrySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldQuorumManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newQuorumManager\",\"type\":\"address\"}],\"name\":\"QuorumManagerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prevRewardsInitiator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newRewardsInitiator\",\"type\":\"address\"}],\"name\":\"RewardsInitiatorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"StaleStakesForbiddenUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldTaskManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newTaskManager\",\"type\":\"address\"}],\"name\":\"TaskManagerContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldTaskManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newTaskManager\",\"type\":\"address\"}],\"name\":\"TaskManagerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldTaskValidator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newTaskValidator\",\"type\":\"address\"}],\"name\":\"TaskValidatorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"avsDirectory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"blacklistKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blsApkRegistry\",\"outputs\":[{\"internalType\":\"contractIBLSApkRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"referenceBlockNumber\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"uint32[]\",\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point[]\",\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point[]\",\"name\":\"quorumApks\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"apkG2\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"sigma\",\"type\":\"tuple\"},{\"internalType\":\"uint32[]\",\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[]\",\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[][]\",\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\"}],\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"checkSignatures\",\"outputs\":[{\"components\":[{\"internalType\":\"uint96[]\",\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\"},{\"internalType\":\"uint96[]\",\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\"}],\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"multiplier\",\"type\":\"uint96\"}],\"internalType\":\"structIRewardsCoordinator.StrategyAndMultiplier[]\",\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"startTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"duration\",\"type\":\"uint32\"}],\"internalType\":\"structIRewardsCoordinator.RewardsSubmission[]\",\"name\":\"rewardsSubmissions\",\"type\":\"tuple[]\"}],\"name\":\"createAVSRewardsSubmission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"multiplier\",\"type\":\"uint96\"}],\"internalType\":\"structIRewardsCoordinator.StrategyAndMultiplier[]\",\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIRewardsCoordinator.OperatorReward[]\",\"name\":\"operatorRewards\",\"type\":\"tuple[]\"},{\"internalType\":\"uint32\",\"name\":\"startTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"duration\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structIRewardsCoordinator.OperatorDirectedRewardsSubmission[]\",\"name\":\"operatorDirectedRewardsSubmissions\",\"type\":\"tuple[]\"}],\"name\":\"createOperatorDirectedAVSRewardsSubmission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegation\",\"outputs\":[{\"internalType\":\"contractIDelegationManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"deregisterOperatorFromAVS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getOperatorRestakedStrategies\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRestakeableStrategies\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITriggerXTaskManager\",\"name\":\"_taskManagerContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardsInitiator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_taskManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_taskValidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_quorumManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isBlackListed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"}],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauserRegistry\",\"outputs\":[{\"internalType\":\"contractIPauserRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quorumManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"name\":\"operatorSignature\",\"type\":\"tuple\"}],\"name\":\"registerOperatorToAVS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registryCoordinator\",\"outputs\":[{\"internalType\":\"contractIRegistryCoordinator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsInitiator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"}],\"name\":\"setClaimerFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPauserRegistry\",\"name\":\"newPauserRegistry\",\"type\":\"address\"}],\"name\":\"setPauserRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_quorumManager\",\"type\":\"address\"}],\"name\":\"setQuorumManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRewardsInitiator\",\"type\":\"address\"}],\"name\":\"setRewardsInitiator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"setStaleStakesForbidden\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_taskManager\",\"type\":\"address\"}],\"name\":\"setTaskManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_taskValidator\",\"type\":\"address\"}],\"name\":\"setTaskValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeRegistry\",\"outputs\":[{\"internalType\":\"contractIStakeRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staleStakesForbidden\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taskManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taskManagerContract\",\"outputs\":[{\"internalType\":\"contractITriggerXTaskManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taskValidator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"apk\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"apkG2\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"sigma\",\"type\":\"tuple\"}],\"name\":\"trySignatureAndApkVerification\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"pairingSuccessful\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"siganatureIsValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"unblacklistKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_metadataURI\",\"type\":\"string\"}],\"name\":\"updateAVSMetadataURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_taskManager\",\"type\":\"address\"}],\"name\":\"updateTaskManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052600436101561001257600080fd5b60003560e01c806310d67a2f146102a7578063136439dd146102a2578063171f1d5b1461029d57806326a965f014610298578063327d0a601461029357806333cfb7b71461028e5780633bc28c8c14610289578063416c7e5e14610284578063500c8dd31461027f57806358ac4a1e1461027a578063595c6a67146102755780635ac86ab7146102705780635c975abb1461026b5780635df45946146102665780635edd78ea14610261578063683048351461025c5780636b3aa72e146102575780636d14a987146102525780636efb46361461024d578063715018a614610248578063886f1195146102435780638da5cb5b1461023e5780639926ee7d146102395780639cd0fb8614610234578063a0169ddd1461022f578063a20b99bf1461022a578063a364f4da14610225578063a41d3f9414610220578063a50a640e1461021b578063a98fb35514610216578063b5f7eb6b14610211578063b98d09081461020c578063cc2a9a5b14610207578063df5cf72314610202578063e47d6060146101fd578063e481af9d146101f8578063f2fde38b146101f3578063fabc1cbc146101ee578063fc299dee146101e9578063fce36c7d146101e45763fd38ec8c146101df57600080fd5b611c71565b611b9b565b611b72565b611a63565b6119d2565b6119b6565b611974565b61192f565b611803565b6117e0565b6117b6565b611700565b6116d7565b61166d565b611551565b61140d565b61130b565b61129c565b611200565b611185565b61115c565b6110ff565b61106a565b610d16565b610cd1565b610c8c565b610c1e565b610bd9565b610bbb565b610b88565b610af3565b610a84565b610a5a565b61091b565b6108e4565b6108ac565b6107fb565b61078e565b610726565b610447565b6102c2565b6001600160a01b038116036102bd57565b600080fd5b346102bd5760203660031901126102bd576004356102df816102ac565b60005460405163755b36bd60e11b81529091906020816004816001600160a01b0387165afa80156104425761032791600091610413575b506001600160a01b03163314611cbe565b6001600160a01b0381161561039657604080516001600160a01b0393841681529282166020840152610394927f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb69190a160018060a01b03166001600160601b0360a01b6000541617600055565b005b60405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a490fd5b610435915060203d60201161043b575b61042d81836105d3565b810190611c9a565b38610316565b503d610423565b611cb2565b346102bd5760203660031901126102bd5760043560005460405163237dfb4760e11b815233600482015290602090829060249082906001600160a01b03165afa80156104425761049f91600091610553575b50611d32565b600154818116036104e857806001557fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d604051806104e33394829190602083019252565b0390a2005b60405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608490fd5b610575915060203d60201161057b575b61056d81836105d3565b810190611d1d565b38610499565b503d610563565b634e487b7160e01b600052604160045260246000fd5b604081019081106001600160401b038211176105b357604052565b610582565b606081019081106001600160401b038211176105b357604052565b90601f801991011681019081106001600160401b038211176105b357604052565b60405190610604610100836105d3565b565b604051906106046040836105d3565b9061060460405192836105d3565b60409060e31901126102bd576040519061063c82610598565b60e4358252610104356020830152565b91908260409103126102bd5760405161066481610598565b6020808294803584520135910152565b9080601f830112156102bd57604080519261068f82856105d3565b839181019283116102bd57905b8282106106a95750505090565b813581526020918201910161069c565b9060806063198301126102bd576040516106d281610598565b60206106ed82946106e4816064610674565b845260a4610674565b910152565b91906080838203126102bd5760206106ed6040519261071084610598565b6040849661071e8382610674565b865201610674565b346102bd576101203660031901126102bd5760043560403660231901126102bd5761077e604091825161075881610598565b6024358152604435602082015261076e366106b9565b9061077836610623565b92611dd1565b8251911515825215156020820152f35b346102bd5760203660031901126102bd576004356107ab816102ac565b6107b3613ede565b6001600160a01b0316600081815260fd60205260408120805460ff191660011790557f79bf5efab685b39bdd2b7b1aba09ffa5e9dfbd86f7aaebab4b52d7ef679d28449080a2005b346102bd5760203660031901126102bd57600435610818816102ac565b610820613ede565b60fe80546001600160a01b039283166001600160a01b0319821681179092559091167fdae652729007c1e5ea13dfe9e126a1b72358db8db69c68b96f316dcea59c1c5f600080a3005b602060408183019282815284518094520192019060005b81811061088d5750505090565b82516001600160a01b0316845260209384019390920191600101610880565b346102bd5760203660031901126102bd576108e06108d46004356108cf816102ac565b612042565b60405191829182610869565b0390f35b346102bd5760203660031901126102bd57610394600435610904816102ac565b61090c613ede565b613fd6565b801515036102bd57565b346102bd5760203660031901126102bd5760043561093881610911565b604051638da5cb5b60e01b81526020816004817f000000000000000000000000ab42fae90d4ac5df10bc069b132203ef566acc596001600160a01b03165afa90811561044257600091610a30575b506001600160a01b0316330361099f5761039490614034565b60405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a490fd5b610a49915060203d60201161043b5761042d81836105d3565b38610986565b60009103126102bd57565b346102bd5760003660031901126102bd57610101546040516001600160a01b039091168152602090f35b346102bd5760203660031901126102bd57600435610aa1816102ac565b610aa9613ede565b61010180546001600160a01b039283166001600160a01b0319821681179092559091167fecc36ab0616c35e413d90730ada58214cbe0188c5f26632fbb32451c9a135784600080a3005b346102bd5760003660031901126102bd5760005460405163237dfb4760e11b815233600482015290602090829060249082906001600160a01b03165afa801561044257610b47916000916105535750611d32565b60001960015560405160001981527fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d60203392a2005b60ff8116036102bd57565b346102bd5760203660031901126102bd576020600160ff600435610bab81610b7d565b161b806001541614604051908152f35b346102bd5760003660031901126102bd576020600154604051908152f35b346102bd5760003660031901126102bd576040517f000000000000000000000000855b96fbf39b92acea95a5f2a9fa5b016f53e4a86001600160a01b03168152602090f35b346102bd5760203660031901126102bd57600435610c3b816102ac565b610c43613ede565b60ff80546001600160a01b039283166001600160a01b0319821681179092559091167f86cebbccc5bbc9ac544f7e5361e6c5e4acf14f39d85c855032a815125cb97fd8600080a3005b346102bd5760003660031901126102bd576040517f000000000000000000000000ff6de6ea349e8b59aca720eb3fbb2c6552e266e06001600160a01b03168152602090f35b346102bd5760003660031901126102bd576040517f000000000000000000000000055733000064333caddbc92763c58bf0192ffebf6001600160a01b03168152602090f35b346102bd5760003660031901126102bd576040517f000000000000000000000000ab42fae90d4ac5df10bc069b132203ef566acc596001600160a01b03168152602090f35b6044359063ffffffff821682036102bd57565b359063ffffffff821682036102bd57565b6001600160401b0381116105b35760051b60200190565b9080601f830112156102bd578135610dad81610d7f565b92610dbb60405194856105d3565b81845260208085019260051b8201019283116102bd57602001905b828210610de35750505090565b60208091610df084610d6e565b815201910190610dd6565b81601f820112156102bd578035610e1181610d7f565b92610e1f60405194856105d3565b81845260208085019260061b840101928184116102bd57602001915b838310610e49575050505090565b6020604091610e58848661064c565b815201920191610e3b565b9080601f830112156102bd578135610e7a81610d7f565b92610e8860405194856105d3565b81845260208085019260051b820101918383116102bd5760208201905b838210610eb457505050505090565b81356001600160401b0381116102bd57602091610ed687848094880101610d96565b815201910190610ea5565b919091610180818403126102bd57610ef76105f4565b9281356001600160401b0381116102bd5781610f14918401610d96565b845260208201356001600160401b0381116102bd5781610f35918401610dfb565b602085015260408201356001600160401b0381116102bd5781610f59918401610dfb565b6040850152610f6b81606084016106f2565b6060850152610f7d8160e0840161064c565b60808501526101208201356001600160401b0381116102bd5781610fa2918401610d96565b60a08501526101408201356001600160401b0381116102bd5781610fc7918401610d96565b60c08501526101608201356001600160401b0381116102bd57610fea9201610e63565b60e0830152565b906020808351928381520192019060005b81811061100f5750505090565b82516001600160601b0316845260209384019390920191600101611002565b929190611065602091604086528261105182516040808a01526080890190610ff1565b910151868203603f19016060880152610ff1565b930152565b346102bd5760803660031901126102bd576004356024356001600160401b0381116102bd57366023820112156102bd5780600401356001600160401b0381116102bd5736602482840101116102bd576110c1610d5b565b90606435936001600160401b0385116102bd5760246110e76110ef963690600401610ee1565b94019061288a565b906108e06040519283928361102e565b346102bd5760003660031901126102bd57611118613ede565b609780546001600160a01b031981169091556000906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b346102bd5760003660031901126102bd576000546040516001600160a01b039091168152602090f35b346102bd5760003660031901126102bd576097546040516001600160a01b039091168152602090f35b6001600160401b0381116105b357601f01601f191660200190565b9291926111d5826111ae565b916111e360405193846105d3565b8294818452818301116102bd578281602093846000960137010152565b346102bd5760403660031901126102bd5760043561121d816102ac565b602435906001600160401b0382116102bd57606060031983360301126102bd576040519061124a826105b8565b82600401356001600160401b0381116102bd578301366023820112156102bd576103949361128460449236906024600482013591016111c9565b84526024810135602085015201356040830152613166565b346102bd5760203660031901126102bd576004356112b9816102ac565b6112c1613ede565b61010080546001600160a01b039283166001600160a01b0319821681179092559091167f5d68f69aebb8a1d75f8b2a87dddd6e8db842f9fa3bb63745ea7f49ea5bdf03a9600080a3005b346102bd57600060203660031901126113b0578060043561132b816102ac565b611333613ede565b7f000000000000000000000000acc1fb458a1317e886db376fc8141540537e68fe6001600160a01b0316803b156113ac5760405163a0169ddd60e01b81526001600160a01b0390921660048301529091908290602490829084905af1801561044257829061139e5780f35b6113a7916105d3565b388180f35b5050fd5b80fd5b9060206003198301126102bd576004356001600160401b0381116102bd5760040160009280601f83011215611409578135936001600160401b0385116113b057506020808301928560051b0101116102bd579190565b8380fd5b346102bd5761141b366113b3565b906114246142e8565b7f000000000000000000000000acc1fb458a1317e886db376fc8141540537e68fe6001600160a01b03169160005b8181106114ab5750823b156102bd576114869260009283604051809681958294634e5cd2fd60e11b84523060048501613415565b03925af180156104425761149657005b806114a56000610394936105d3565b80610a4f565b916000939184915b6114cb6114c18684846132c0565b60408101906132e2565b90508310156115075760016114fd819760206114f5876114ef6114c18c8a8a6132c0565b90613317565b013590611fc7565b93019295506114b3565b939092946001925061154b90611535813088611530602061152a898c33956132c0565b01613327565b61437c565b86611546602061152a86898b6132c0565b6143c5565b01611452565b346102bd5760203660031901126102bd5760043561156e816102ac565b611576613ede565b7f000000000000000000000000055733000064333caddbc92763c58bf0192ffebf6001600160a01b0316803b156102bd576040516351b27a6d60e11b81526001600160a01b0383166004820152906000908290602490829084905af1801561044257611658575b506116006115fb6001600160a01b0383165b6001600160a01b031690565b614807565b506001600160a01b038116600090815260fd60205260409020611628905b805460ff19169055565b6001600160a01b03167fa7a775c2c8141f7985c111748ec31c11e5e44b83528e105c8d1d4e8e6b81cf80600080a2005b806114a56000611667936105d3565b386115dd565b346102bd5760203660031901126102bd5760043561168a816102ac565b611692613ede565b6001600160a01b0316600081815260fd60205260408120805460ff191690557f8511046d95f95689c89b4807d91820adbc266db891b2c5c25a42f812781dde579080a2005b346102bd5760003660031901126102bd5760fe546040516001600160a01b039091168152602090f35b346102bd57600060203660031901126113b057806004356001600160401b0381116117b357366023820112156117b3576117449036906024816004013591016111c9565b61174c613ede565b7f000000000000000000000000055733000064333caddbc92763c58bf0192ffebf6001600160a01b031690813b156113ac5782916117a09160405194858094819363a98fb35560e01b835260048301613587565b03925af1801561044257829061139e5780f35b50fd5b346102bd5760003660031901126102bd57610100546040516001600160a01b039091168152602090f35b346102bd5760003660031901126102bd57602060ff603254166040519015158152f35b346102bd5760c03660031901126102bd57600435611820816102ac565b6118a360243561182f816102ac565b60443561183b816102ac565b606435611847816102ac565b60843591611854836102ac565b60a43593611861856102ac565b6064549661188760ff60088a901c16158099819a611921575b8115611901575b50613598565b8761189a600160ff196064541617606455565b6118e8576135fb565b6118a957005b6118b961ff001960645416606455565b604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602090a1005b6118fc61010061ff00196064541617606455565b6135fb565b303b15915081611913575b5038611881565b60ff1660011490503861190c565b600160ff821610915061187a565b346102bd5760003660031901126102bd576040517f000000000000000000000000a44151489861fe9e3055d95adc98fbd462b948e76001600160a01b03168152602090f35b346102bd5760203660031901126102bd57600435611991816102ac565b60018060a01b031660005260fd602052602060ff604060002054166040519015158152f35b346102bd5760003660031901126102bd576108e06108d46136fd565b346102bd5760203660031901126102bd576004356119ef816102ac565b6119f7613ede565b6001600160a01b03811615611a0f576103949061429f565b60405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608490fd5b346102bd5760203660031901126102bd5760005460405163755b36bd60e11b81526004803592602091839182906001600160a01b03165afa801561044257611abd9160009161041357506001600160a01b03163314611cbe565b600154198119811603611b0757611ad381600155565b60405190815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9080602081016104e3565b60405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608490fd5b346102bd5760003660031901126102bd5760c9546040516001600160a01b039091168152602090f35b346102bd57611ba9366113b3565b90611bb26142e8565b7f000000000000000000000000acc1fb458a1317e886db376fc8141540537e68fe6001600160a01b03169160005b818110611c135750823b156102bd57611486926000928360405180968195829463fce36c7d60e01b845260048401613955565b80611c42611c29602061152a6001958789613933565b6040611c36848789613933565b0135903090339061437c565b611c6b611c55602061152a848789613933565b866040611c6385888a613933565b0135916143c5565b01611be0565b346102bd5760003660031901126102bd5760ff546040516001600160a01b039091168152602090f35b908160209103126102bd5751611caf816102ac565b90565b6040513d6000823e3d90fd5b15611cc557565b60405162461bcd60e51b815260206004820152602a60248201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160448201526939903ab73830bab9b2b960b11b6064820152608490fd5b908160209103126102bd5751611caf81610911565b15611d3957565b60405162461bcd60e51b815260206004820152602860248201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160448201526739903830bab9b2b960c11b6064820152608490fd5b634e487b7160e01b600052603260045260246000fd5b906002811015611db65760051b0190565b611d8f565b634e487b7160e01b600052601260045260246000fd5b611ead611e8a611eb395611e84611e7d85875160208901518a515160208c51015160208d016020815151915101519189519360208b0151956040519760208901998a5260208a015260408901526060880152608087015260a086015260c085015260e0840152610100830152611e5481610120840103601f1981018352826105d3565b5190207f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001900690565b8096613a6a565b90613ab2565b92611e84611e9f611e99613b3c565b94613c33565b91611ea8613d5a565b613a6a565b91613da4565b9091565b908160209103126102bd575190565b908160209103126102bd57516001600160c01b03811681036102bd5790565b908160209103126102bd5751611caf81610b7d565b60405190611f096020836105d3565b6000808352366020840137565b90611f2082610d7f565b611f2d60405191826105d3565b8281528092611f3e601f1991610d7f565b0190602036910137565b805115611db65760200190565b908151811015611db6570160200190565b634e487b7160e01b600052601160045260246000fd5b9060018201809211611f8a57565b611f66565b9060028201809211611f8a57565b9060038201809211611f8a57565b9060048201809211611f8a57565b9060058201809211611f8a57565b91908201809211611f8a57565b6001600160601b038116036102bd57565b908160409103126102bd57602060405191611fff83610598565b805161200a816102ac565b8352015161201781611fd4565b602082015290565b8051821015611db65760209160051b010190565b6000198114611f8a5760010190565b6040516309aa152760e11b81526001600160a01b0391821660048201527f000000000000000000000000ab42fae90d4ac5df10bc069b132203ef566acc5990911690602081602481855afa908115610442576120c3916020916000916123bc575b506040518093819263871ef04960e01b8352600483019190602083019252565b0381855afa9081156104425760009161238d575b506001600160c01b0316908115908115612329575b50612320576120fa90613f36565b600091907f000000000000000000000000ff6de6ea349e8b59aca720eb3fbb2c6552e266e06001600160a01b031690835b81518510156121d857612180602061215d6121576121498987611f55565b516001600160f81b03191690565b60f81c90565b604051633ca5a5f560e01b815260ff909116600482015291829081906024820190565b0381875afa8015610442576001926121a0926000926121a8575b50611fc7565b94019361212b565b6121ca91925060203d81116121d1575b6121c281836105d3565b810190611eb7565b903861219a565b503d6121b8565b6121e3919450611f16565b9260009060005b815181101561231a576122036121576121498385611f55565b604051633ca5a5f560e01b815260ff8216600482015290602082602481895afa918215610442576000926122fa575b50906000915b81831061224a575050506001016121ea565b604080516356e4026d60e11b815260ff83166004820152602481018590529396929391929190816044818b5afa918215610442576122be8b6122af836122a96115ef6001986122c3986000916122cc575b50516001600160a01b031690565b9261201f565b6001600160a01b039091169052565b612033565b95019190612238565b6122ed915060403d81116122f3575b6122e581836105d3565b810190611fe5565b3861229b565b503d6122db565b61231391925060203d81116121d1576121c281836105d3565b9038612232565b50505050565b50611caf611efa565b604051639aa1653d60e01b81529150602090829060049082905afa80156104425760ff9160009161235e575b501615386120ec565b612380915060203d602011612386575b61237881836105d3565b810190611ee5565b38612355565b503d61236e565b6123af915060203d6020116123b5575b6123a781836105d3565b810190611ec6565b386120d7565b503d61239d565b6123d39150823d84116121d1576121c281836105d3565b386120a3565b604051906123e682610598565b60606020838281520152565b156123f957565b60405162461bcd60e51b8152602060048201526037602482015260008051602061499083398151915260448201527f7265733a20656d7074792071756f72756d20696e7075740000000000000000006064820152608490fd5b1561245957565b60405162461bcd60e51b8152602060048201526041602482015260008051602061499083398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a490fd5b156124c357565b60a460405162461bcd60e51b8152602060048201526044602482015260008051602061499083398151915260448201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b6084820152fd5b1561252f57565b60405162461bcd60e51b815260206004820152603c602482015260008051602061499083398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b000000006064820152608490fd5b600019810191908211611f8a57565b1561259e57565b608460405162461bcd60e51b8152602060048201526040602482015260008051602061499083398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f727465646064820152fd5b90821015611db6570190565b1561260957565b60405162461bcd60e51b8152602060048201526066602482015260008051602061499083398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c490fd5b908160209103126102bd575167ffffffffffffffff19811681036102bd5790565b156126bf57565b60405162461bcd60e51b8152602060048201526061602482015260008051602061499083398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c490fd5b908160209103126102bd5751611caf81611fd4565b906001600160601b03809116911603906001600160601b038211611f8a57565b1561278457565b60405162461bcd60e51b8152602060048201526043602482015260008051602061499083398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a490fd5b156127f057565b60405162461bcd60e51b8152602060048201526039602482015260008051602061499083398151915260448201527f7265733a207369676e617475726520697320696e76616c6964000000000000006064820152608490fd5b60049163ffffffff60e01b9060e01b16815201602082519192019060005b8181106128745750505090565b8251845260209384019390920191600101612867565b9493929091936128986123d9565b506128a48515156123f2565b604084015151851480613158575b8061314a575b8061313c575b6128c790612452565b6128d9602085015151855151146124bc565b6128f063ffffffff431663ffffffff841610612528565b6128f8610606565b60008152600060208201529261290c6123d9565b61291587611f16565b602082015261292387611f16565b815261292d6123d9565b9261293c602088015151611f16565b845261294c602088015151611f16565b602085810191909152604051639aa1653d60e01b815290816004817f000000000000000000000000ab42fae90d4ac5df10bc069b132203ef566acc596001600160a01b03165afa8015610442576129b69160009161311d575b506129b1368b876111c9565b614072565b986000965b60208901518051891015612b3657602088612a2a612a208c612a188f96868e6129fd6129e886809561201f565b51805160005260200151602052604060002090565b612a0a848484015161201f565b5282612b03575b015161201f565b51955161201f565b5163ffffffff1690565b6040516304ec635160e01b8152600481019490945263ffffffff9182166024850152166044830152816064816001600160a01b037f000000000000000000000000ab42fae90d4ac5df10bc069b132203ef566acc59165afa91821561044257611e848a612ad88f612ad18f8460208f92612ac893612ac08460019e612ade9e600091612ae6575b508f8060c01b0316925161201f565b52015161201f565b51938d5161201f565b51166140f9565b9061412c565b9701966129bb565b612afd9150863d81116123b5576123a781836105d3565b38612ab1565b612b31612b13848484015161201f565b51612b2a84840151612b2487612588565b9061201f565b5110612597565b612a11565b50909597949650612b4b919893929950614215565b91612b5860325460ff1690565b908115613114576040516318891fd760e31b81526020816004817f000000000000000000000000a44151489861fe9e3055d95adc98fbd462b948e76001600160a01b03165afa908115610442576000916130f5575b5091905b6000925b818410612c1957505050505092612bf2612bed612be6612c139585612c059860806060602099015192015192611dd1565b919061277d565b6127e9565b0151604051928391602083019586612849565b03601f1981018352826105d3565b51902090565b92989596909399919794878b888c888d612ff1575b612a208260a0612c7c612157612c6e84612c8497612c68612c5a6129e88f9c604060209f9e015161201f565b67ffffffffffffffff191690565b9b6125f6565b356001600160f81b03191690565b97015161201f565b604051631a2f32ab60e21b815260ff95909516600486015263ffffffff9182166024860152166044840152826064816001600160a01b037f000000000000000000000000855b96fbf39b92acea95a5f2a9fa5b016f53e4a8165afa90811561044257612d49612a208f958f90612d418f978f96848f612d3b60c096612d34848f60209f90612a11612c6e996040936121579c600091612fc3575b5067ffffffffffffffff199182169116146126b8565b5190613ab2565b9c6125f6565b96015161201f565b604051636414a62b60e11b815260ff94909416600485015263ffffffff9182166024850152166044830152816064816001600160a01b037f000000000000000000000000ff6de6ea349e8b59aca720eb3fbb2c6552e266e0165afa90811561044257612dd7918c8f92600092612f9f575b506020612dc99293015161201f565b906001600160601b03169052565b612e048c612dc98c612dfd612df082602086015161201f565b516001600160601b031690565b925161201f565b600098895b60208a015151811015612f86578b8d612e4789612e3a612157612c6e868f89612e32915161201f565b5194876125f6565b60ff161c60019081161490565b612e56575b5050600101612e09565b8a8a612ede859f948f9686612e988f9360e0612e8f612a20956020612e87612157612c6e839f612e9e9c89916125f6565b9a015161201f565b519b015161201f565b5161201f565b60405163795f4a5760e11b815260ff909316600484015263ffffffff93841660248401526044830196909652919094166064850152839081906084820190565b03817f000000000000000000000000ff6de6ea349e8b59aca720eb3fbb2c6552e266e06001600160a01b03165afa908115610442578f612f45908f93600195948695600092612f50575b506122a9612dc992935193612f40612df0848761201f565b61275d565b019a90508b8d612e4c565b612dc99250612f786122a99160203d8111612f7f575b612f7081836105d3565b810190612748565b9250612f28565b503d612f66565b5093919796996001919699509a94929a01929190612bb5565b612dc99250612fbc602091823d8111612f7f57612f7081836105d3565b9250612dba565b6020612fe492503d8111612fea575b612fdc81836105d3565b810190612697565b38612d1e565b503d612fd2565b61302e945061300b925061215791612c6e916020956125f6565b60405163124d062160e11b815260ff909116600482015291829081906024820190565b03817f000000000000000000000000ab42fae90d4ac5df10bc069b132203ef566acc596001600160a01b03165afa801561044257602089612c848f938f60a08f97612157612c6e8f8f90612c68612c5a6129e88f60408b96918f6130b690612a209f8a95612c7c9e6000926130cc575b5063ffffffff6130b092931692611fc7565b11612602565b5050505050509750505050505092935050612c2e565b60206130b0935063ffffffff916130ee913d81116121d1576121c281836105d3565b925061309e565b61310e915060203d6020116121d1576121c281836105d3565b38612bad565b60009190612bb1565b613136915060203d6020116123865761237881836105d3565b386129a5565b5060e08401515185146128be565b5060c08401515185146128b8565b5060a08401515185146128b2565b61316e613ede565b7f000000000000000000000000055733000064333caddbc92763c58bf0192ffebf6001600160a01b031690813b156102bd5760006040518093639926ee7d60e01b825281838160018060a01b038716988960048301526040602483015260406131e382516060604486015260a485019061327f565b91602081015160648501520151608483015203925af1908115610442576132439261161e9261326a575b506132286132236001600160a01b0383166115ef565b614610565b506001600160a01b0316600090815260fd6020526040902090565b7f1584773458d98c71b34a270ee1100b3a42889bf91e3b7a858563b684c24d838e600080a2565b806114a56000613279936105d3565b3861320d565b919082519283825260005b8481106132ab575050826000602080949584010152601f8019910116010190565b8060208092840101518282860101520161328a565b9190811015611db65760051b8101359060be19813603018212156102bd570190565b903590601e19813603018212156102bd57018035906001600160401b0382116102bd57602001918160061b360383136102bd57565b9190811015611db65760061b0190565b35611caf816102ac565b9035601e19823603018112156102bd5701602081359101916001600160401b0382116102bd578160061b360383136102bd57565b9160209082815201919060005b81811061337f5750505090565b9091926040806001928635613393816102ac565b848060a01b031681526001600160601b0360208801356133b281611fd4565b166020820152019401929101613372565b9035601e19823603018112156102bd5701602081359101916001600160401b0382116102bd5781360383136102bd57565b908060209392818452848401376000828201840152601f01601f1916010190565b6001600160a01b0390911681526040602082018190528101839052600583901b810160609081019383923684900360be190192600091908101905b838310613461575050505050505090565b90919293949596605f198282030183528735868112156102bd5787019061349961348b8380613331565b60c0845260c0840191613365565b9160208101356134a8816102ac565b6001600160a01b03166020838101919091526134c76040830183613331565b8486036040860152808652949091019360005b818110613553575050506135426001936020936135348461350e6135016060899801610d6e565b63ffffffff166060850152565b61352a61351d60808301610d6e565b63ffffffff166080850152565b60a08101906133c3565b9160a08185039101526133f4565b990193019301919594939290613450565b9091946040806001928835613567816102ac565b848060a01b031681526020890135602082015201960191019190916134da565b906020611caf92818152019061327f565b1561359f57565b60405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608490fd5b94939190926136098461429f565b60ff60645460081c16156136a4576106049561362e6136689361090c6136859761429f565b60018060a01b03166001600160601b0360a01b6101015416176101015560018060a01b03166001600160601b0360a01b60fe54161760fe55565b60018060a01b03166001600160601b0360a01b60ff54161760ff55565b60018060a01b03166001600160601b0360a01b61010054161761010055565b60405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b6064820152608490fd5b604051639aa1653d60e01b81527f000000000000000000000000ab42fae90d4ac5df10bc069b132203ef566acc596001600160a01b031690602081600481855afa80156104425760ff91600091613914575b5016801561390a577f000000000000000000000000ff6de6ea349e8b59aca720eb3fbb2c6552e266e06001600160a01b03169060009081905b8083106138c5575061379a9150611f16565b9260009060005b604051639aa1653d60e01b8152602081600481895afa80156104425760ff916000916138a7575b50168110156138a057604051633ca5a5f560e01b815260ff821660048201819052602082602481895afa91821561044257600092613880575b50906000915b818310613819575050506001016137a1565b604080516356e4026d60e11b815260ff83166004820152602481018590529396929391929190816044818b5afa918215610442576122be8b6122af836122a96115ef600198613877986000916122cc5750516001600160a01b031690565b95019190613807565b61389991925060203d81116121d1576121c281836105d3565b9038613801565b5092505050565b6138bf915060203d81116123865761237881836105d3565b386137c8565b604051633ca5a5f560e01b815260ff84166004820152909190602081602481885afa801561044257600192613901926000926121a85750611fc7565b92019190613788565b5050611caf611efa565b61392d915060203d6020116123865761237881836105d3565b3861374f565b9190811015611db65760051b81013590609e19813603018212156102bd570190565b909180602083016020845252604082019260408260051b840101938193600091609e1984360301915b85841061398f575050505050505090565b90919293949596603f19828203018352873590848212156102bd5760208091886001940190608063ffffffff613a16826139da6139cc8780613331565b60a0885260a0880191613365565b95878101356139e8816102ac565b8a8060a01b0316888701526040810135604087015283613a0a60608301610d6e565b16606087015201610d6e565b1691015299019301940192919594939061397e565b60405190613a3882610598565b60006020838281520152565b60405190610180613a5581846105d3565b368337565b604051906020613a5581846105d3565b604090929192613a78613a2b565b9384916060916020855192613a8d85856105d3565b8436853780518452015160208301528482015260076107cf195a01fa15613ab057565bfe5b604090929192613ac0613a2b565b938491602060809281865193613ad686866105d3565b85368637805185520151828401528051868401520151606082015260066107cf195a01fa8015613ab05715613b0757565b60405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b6044820152606490fd5b604051613b4881610598565b6040908151613b5783826105d3565b8236823781526020825191613b6c84846105d3565b8336843701528051613b7e82826105d3565b7f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c281527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6020820152815190613bd483836105d3565b7f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d6020830152613c29835193846105d3565b8252602082015290565b60008051602061497083398151915290613c4b613a2b565b5006906000908192602060c0945b613d555760009360008051602061497083398151915260038185818180090908604051613c8684826105d3565b83368237838189604051613c9a82826105d3565b813682378381528360208201528360408201528560608201527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52608082015260008051602061497083398151915260a082015260056107cf195a01fa8015613ab057613d05906148af565b5191613d555760008051602061497083398151915282800914613d3e575060008051602061497083398151915260018593089193613c59565b929450909250613d4c610606565b92835282015290565b611dbb565b613d62613a2b565b50604051613d6f81610598565b600181526002602082015290565b90600682029180830460061490151715611f8a57565b90600c811015611db65760051b0190565b93929091613db26040610615565b9485526020850152613dc46040610615565b9182526020820152613dd4613a44565b9260005b60028110613e0257505050602061018092613df1613a5a565b93849160086201d4c0fa9151151590565b80613e0e600192613d7d565b613e188285611da5565b5151600090613e27838a613d93565b526020613e348487611da5565b510151613e49613e4384611f7c565b8a613d93565b52613e548387611da5565b515151613e63613e4384611f8f565b52613e79613e718488611da5565b515160200190565b51613e86613e4384611f9d565b526020613e938488611da5565b510151905051613eab613ea583611fab565b89613d93565b52613ed7613ed1613eca6020613ec1868a611da5565b51015160200190565b5192611fb9565b88613d93565b5201613dd8565b6097546001600160a01b03163303613ef257565b606460405162461bcd60e51b815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b61ffff613f42826140f9565b16613f4c816111ae565b90613f5a60405192836105d3565b808252613f69601f19916111ae565b013660208301376000805b8251821080613fcb575b15613fc4576001811b8416613f9c575b613f9790612033565b613f74565b906001613f979160ff60f81b8460f81b1660001a613fba8287611f55565b5301919050613f8e565b5050905090565b506101008110613f7e565b60c954604080516001600160a01b038084168252841660208201529192917fe11cddf1816a43318ca175bbc52cd0185436e9cbead7c83acc54a73e461717e39190a16001600160a01b03166001600160a01b0319919091161760c955565b60207f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc91151560ff196032541660ff821617603255604051908152a1565b90600161408060ff936144f2565b928392161b111561408e5790565b60405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c7565006064820152608490fd5b806000915b614106575090565b6000198101818111611f8a5761ffff9116911661ffff8114611f8a5760010190806140fe565b90614135613a2b565b5061ffff8116906102008210156141dd57600182146141d857614156610606565b600081526000602082015292906001906000925b61ffff831685101561417e57505050505090565b600161ffff831660ff86161c8116146141b8575b60016141ae6141a38360ff94613ab2565b9460011b61fffe1690565b940116929161416a565b9460016141ae6141a36141cd8960ff95613ab2565b989350505050614192565b505090565b60405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b6044820152606490fd5b61421d613a2b565b50805190811580614293575b1561424c57505060405161423e6040826105d3565b600081526000602082015290565b602060008051602061497083398151915291015106600080516020614970833981519152036000805160206149708339815191528111611f8a5760405191613c2983610598565b50602081015115614229565b609780546001600160a01b039283166001600160a01b0319821681179092559091167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3565b60c9546001600160a01b031633036142fc57565b60405162461bcd60e51b815260206004820152604c60248201527f536572766963654d616e61676572426173652e6f6e6c7952657761726473496e60448201527f69746961746f723a2063616c6c6572206973206e6f742074686520726577617260648201526b32399034b734ba34b0ba37b960a11b608482015260a490fd5b6040516323b872dd60e01b60208201526001600160a01b039283166024820152929091166044830152606480830193909352918152610604916143c06084836105d3565b6146f3565b604051636eb1769f60e11b81523060048201526001600160a01b0383166024820152602081806044810103816001600160a01b0386165afa90811561044257610604946143c09261441d926000916144515750611fc7565b60405163095ea7b360e01b60208201526001600160a01b0394909416602485015260448085019190915283526064836105d3565b61446a915060203d6020116121d1576121c281836105d3565b3861219a565b1561447757565b60405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a490fd5b906101008251116145635781511561455d5761452061451661215761214985611f48565b60ff600191161b90565b6001905b8351821015614558576001906145436145166121576121498689611f55565b9061454f818311614470565b17910190614524565b925050565b60009150565b60a460405162461bcd60e51b815260206004820152604460248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b6084820152fd5b8054821015611db65760005260206000200190600090565b9161460c9183549060031b91821b91600019901b19161790565b9055565b8060005260fc6020526040600020541560001461468e5760fb54680100000000000000008110156105b3576001810160fb55600060fb54821015611db65760fb90527f3e7257b7272bb46d49cd6019b04ddee20da7c0cb13f7c1ec3391291b2ccebabc0181905560fb549060005260fc602052604060002055600190565b50600090565b1561469b57565b60405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608490fd5b6040805192916001600160a01b03169061470d81856105d3565b602084527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020850152813b15614785575060008281928260206147609796519301915af161475a6148fb565b9061492b565b80518061476b575050565b81602080614780936106049501019101611d1d565b614694565b5162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b805480156147f15760001901906147df82826145da565b8154906000199060031b1b1916905555565b634e487b7160e01b600052603160045260246000fd5b600081815260fc60205260409020549081156148a857600019820190828211611f8a5760fb54600019810193908411611f8a578383614867946000960361486d575b50505061485660fb6147c8565b60fc90600052602052604060002090565b55600190565b6148566148999161488f61488561489f9560fb6145da565b90549060031b1c90565b92839160fb6145da565b906145f2565b55388080614849565b5050600090565b156148b657565b60405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c7572650000000000006044820152606490fd5b3d15614926573d9061490c826111ae565b9161491a60405193846105d3565b82523d6000602084013e565b606090565b90919015614937575090565b8151156149475750805190602001fd5b60405162461bcd60e51b81526020600482015290819061496b90602483019061327f565b0390fdfe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a264697066735822122009b9fbda769cedc6fd898390d56355ad6f1dd852613da6648fe45ba26bbd432f64736f6c634300081b0033",
}

// ContractTriggerXServiceManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractTriggerXServiceManagerMetaData.ABI instead.
var ContractTriggerXServiceManagerABI = ContractTriggerXServiceManagerMetaData.ABI

// ContractTriggerXServiceManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractTriggerXServiceManagerMetaData.Bin instead.
var ContractTriggerXServiceManagerBin = ContractTriggerXServiceManagerMetaData.Bin

// DeployContractTriggerXServiceManager deploys a new Ethereum contract, binding an instance of ContractTriggerXServiceManager to it.
func DeployContractTriggerXServiceManager(auth *bind.TransactOpts, backend bind.ContractBackend, _avsDirectory common.Address, _rewardsCoordinator common.Address, _registryCoordinator common.Address, _stakeRegistry common.Address) (common.Address, *types.Transaction, *ContractTriggerXServiceManager, error) {
	parsed, err := ContractTriggerXServiceManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractTriggerXServiceManagerBin), backend, _avsDirectory, _rewardsCoordinator, _registryCoordinator, _stakeRegistry)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractTriggerXServiceManager{ContractTriggerXServiceManagerCaller: ContractTriggerXServiceManagerCaller{contract: contract}, ContractTriggerXServiceManagerTransactor: ContractTriggerXServiceManagerTransactor{contract: contract}, ContractTriggerXServiceManagerFilterer: ContractTriggerXServiceManagerFilterer{contract: contract}}, nil
}

// ContractTriggerXServiceManagerMethods is an auto generated interface around an Ethereum contract.
type ContractTriggerXServiceManagerMethods interface {
	ContractTriggerXServiceManagerCalls
	ContractTriggerXServiceManagerTransacts
	ContractTriggerXServiceManagerFilters
}

// ContractTriggerXServiceManagerCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractTriggerXServiceManagerCalls interface {
	AvsDirectory(opts *bind.CallOpts) (common.Address, error)

	BlsApkRegistry(opts *bind.CallOpts) (common.Address, error)

	CheckSignatures(opts *bind.CallOpts, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error)

	Delegation(opts *bind.CallOpts) (common.Address, error)

	GetOperatorRestakedStrategies(opts *bind.CallOpts, operator common.Address) ([]common.Address, error)

	GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error)

	IsBlackListed(opts *bind.CallOpts, arg0 common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts, index uint8) (bool, error)

	Paused0(opts *bind.CallOpts) (*big.Int, error)

	PauserRegistry(opts *bind.CallOpts) (common.Address, error)

	QuorumManager(opts *bind.CallOpts) (common.Address, error)

	RegistryCoordinator(opts *bind.CallOpts) (common.Address, error)

	RewardsInitiator(opts *bind.CallOpts) (common.Address, error)

	StakeRegistry(opts *bind.CallOpts) (common.Address, error)

	StaleStakesForbidden(opts *bind.CallOpts) (bool, error)

	TaskManager(opts *bind.CallOpts) (common.Address, error)

	TaskManagerContract(opts *bind.CallOpts) (common.Address, error)

	TaskValidator(opts *bind.CallOpts) (common.Address, error)

	TrySignatureAndApkVerification(opts *bind.CallOpts, msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
		PairingSuccessful bool
		SiganatureIsValid bool
	}, error)
}

// ContractTriggerXServiceManagerTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractTriggerXServiceManagerTransacts interface {
	BlacklistKeeper(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error)

	CreateAVSRewardsSubmission(opts *bind.TransactOpts, rewardsSubmissions []IRewardsCoordinatorRewardsSubmission) (*types.Transaction, error)

	CreateOperatorDirectedAVSRewardsSubmission(opts *bind.TransactOpts, operatorDirectedRewardsSubmissions []IRewardsCoordinatorOperatorDirectedRewardsSubmission) (*types.Transaction, error)

	DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts, _taskManagerContract common.Address, initialOwner common.Address, rewardsInitiator common.Address, _taskManager common.Address, _taskValidator common.Address, _quorumManager common.Address) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error)

	PauseAll(opts *bind.TransactOpts) (*types.Transaction, error)

	RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error)

	RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetClaimerFor(opts *bind.TransactOpts, claimer common.Address) (*types.Transaction, error)

	SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error)

	SetQuorumManager(opts *bind.TransactOpts, _quorumManager common.Address) (*types.Transaction, error)

	SetRewardsInitiator(opts *bind.TransactOpts, newRewardsInitiator common.Address) (*types.Transaction, error)

	SetStaleStakesForbidden(opts *bind.TransactOpts, value bool) (*types.Transaction, error)

	SetTaskManager(opts *bind.TransactOpts, _taskManager common.Address) (*types.Transaction, error)

	SetTaskValidator(opts *bind.TransactOpts, _taskValidator common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error)

	UnblacklistKeeper(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error)

	UpdateAVSMetadataURI(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error)

	UpdateTaskManager(opts *bind.TransactOpts, _taskManager common.Address) (*types.Transaction, error)
}

// ContractTriggerXServiceManagerFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractTriggerXServiceManagerFilters interface {
	FilterInitialized(opts *bind.FilterOpts) (*ContractTriggerXServiceManagerInitializedIterator, error)
	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractTriggerXServiceManagerInitialized) (event.Subscription, error)
	ParseInitialized(log types.Log) (*ContractTriggerXServiceManagerInitialized, error)

	FilterKeeperAdded(opts *bind.FilterOpts, operator []common.Address) (*ContractTriggerXServiceManagerKeeperAddedIterator, error)
	WatchKeeperAdded(opts *bind.WatchOpts, sink chan<- *ContractTriggerXServiceManagerKeeperAdded, operator []common.Address) (event.Subscription, error)
	ParseKeeperAdded(log types.Log) (*ContractTriggerXServiceManagerKeeperAdded, error)

	FilterKeeperBlacklisted(opts *bind.FilterOpts, operator []common.Address) (*ContractTriggerXServiceManagerKeeperBlacklistedIterator, error)
	WatchKeeperBlacklisted(opts *bind.WatchOpts, sink chan<- *ContractTriggerXServiceManagerKeeperBlacklisted, operator []common.Address) (event.Subscription, error)
	ParseKeeperBlacklisted(log types.Log) (*ContractTriggerXServiceManagerKeeperBlacklisted, error)

	FilterKeeperRemoved(opts *bind.FilterOpts, operator []common.Address) (*ContractTriggerXServiceManagerKeeperRemovedIterator, error)
	WatchKeeperRemoved(opts *bind.WatchOpts, sink chan<- *ContractTriggerXServiceManagerKeeperRemoved, operator []common.Address) (event.Subscription, error)
	ParseKeeperRemoved(log types.Log) (*ContractTriggerXServiceManagerKeeperRemoved, error)

	FilterKeeperUnblacklisted(opts *bind.FilterOpts, operator []common.Address) (*ContractTriggerXServiceManagerKeeperUnblacklistedIterator, error)
	WatchKeeperUnblacklisted(opts *bind.WatchOpts, sink chan<- *ContractTriggerXServiceManagerKeeperUnblacklisted, operator []common.Address) (event.Subscription, error)
	ParseKeeperUnblacklisted(log types.Log) (*ContractTriggerXServiceManagerKeeperUnblacklisted, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractTriggerXServiceManagerOwnershipTransferredIterator, error)
	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractTriggerXServiceManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error)
	ParseOwnershipTransferred(log types.Log) (*ContractTriggerXServiceManagerOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractTriggerXServiceManagerPausedIterator, error)
	WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractTriggerXServiceManagerPaused, account []common.Address) (event.Subscription, error)
	ParsePaused(log types.Log) (*ContractTriggerXServiceManagerPaused, error)

	FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractTriggerXServiceManagerPauserRegistrySetIterator, error)
	WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractTriggerXServiceManagerPauserRegistrySet) (event.Subscription, error)
	ParsePauserRegistrySet(log types.Log) (*ContractTriggerXServiceManagerPauserRegistrySet, error)

	FilterQuorumManagerSet(opts *bind.FilterOpts, oldQuorumManager []common.Address, newQuorumManager []common.Address) (*ContractTriggerXServiceManagerQuorumManagerSetIterator, error)
	WatchQuorumManagerSet(opts *bind.WatchOpts, sink chan<- *ContractTriggerXServiceManagerQuorumManagerSet, oldQuorumManager []common.Address, newQuorumManager []common.Address) (event.Subscription, error)
	ParseQuorumManagerSet(log types.Log) (*ContractTriggerXServiceManagerQuorumManagerSet, error)

	FilterRewardsInitiatorUpdated(opts *bind.FilterOpts) (*ContractTriggerXServiceManagerRewardsInitiatorUpdatedIterator, error)
	WatchRewardsInitiatorUpdated(opts *bind.WatchOpts, sink chan<- *ContractTriggerXServiceManagerRewardsInitiatorUpdated) (event.Subscription, error)
	ParseRewardsInitiatorUpdated(log types.Log) (*ContractTriggerXServiceManagerRewardsInitiatorUpdated, error)

	FilterStaleStakesForbiddenUpdate(opts *bind.FilterOpts) (*ContractTriggerXServiceManagerStaleStakesForbiddenUpdateIterator, error)
	WatchStaleStakesForbiddenUpdate(opts *bind.WatchOpts, sink chan<- *ContractTriggerXServiceManagerStaleStakesForbiddenUpdate) (event.Subscription, error)
	ParseStaleStakesForbiddenUpdate(log types.Log) (*ContractTriggerXServiceManagerStaleStakesForbiddenUpdate, error)

	FilterTaskManagerContractUpdated(opts *bind.FilterOpts, oldTaskManager []common.Address, newTaskManager []common.Address) (*ContractTriggerXServiceManagerTaskManagerContractUpdatedIterator, error)
	WatchTaskManagerContractUpdated(opts *bind.WatchOpts, sink chan<- *ContractTriggerXServiceManagerTaskManagerContractUpdated, oldTaskManager []common.Address, newTaskManager []common.Address) (event.Subscription, error)
	ParseTaskManagerContractUpdated(log types.Log) (*ContractTriggerXServiceManagerTaskManagerContractUpdated, error)

	FilterTaskManagerSet(opts *bind.FilterOpts, oldTaskManager []common.Address, newTaskManager []common.Address) (*ContractTriggerXServiceManagerTaskManagerSetIterator, error)
	WatchTaskManagerSet(opts *bind.WatchOpts, sink chan<- *ContractTriggerXServiceManagerTaskManagerSet, oldTaskManager []common.Address, newTaskManager []common.Address) (event.Subscription, error)
	ParseTaskManagerSet(log types.Log) (*ContractTriggerXServiceManagerTaskManagerSet, error)

	FilterTaskValidatorSet(opts *bind.FilterOpts, oldTaskValidator []common.Address, newTaskValidator []common.Address) (*ContractTriggerXServiceManagerTaskValidatorSetIterator, error)
	WatchTaskValidatorSet(opts *bind.WatchOpts, sink chan<- *ContractTriggerXServiceManagerTaskValidatorSet, oldTaskValidator []common.Address, newTaskValidator []common.Address) (event.Subscription, error)
	ParseTaskValidatorSet(log types.Log) (*ContractTriggerXServiceManagerTaskValidatorSet, error)

	FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractTriggerXServiceManagerUnpausedIterator, error)
	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractTriggerXServiceManagerUnpaused, account []common.Address) (event.Subscription, error)
	ParseUnpaused(log types.Log) (*ContractTriggerXServiceManagerUnpaused, error)
}

// ContractTriggerXServiceManager is an auto generated Go binding around an Ethereum contract.
type ContractTriggerXServiceManager struct {
	ContractTriggerXServiceManagerCaller     // Read-only binding to the contract
	ContractTriggerXServiceManagerTransactor // Write-only binding to the contract
	ContractTriggerXServiceManagerFilterer   // Log filterer for contract events
}

// ContractTriggerXServiceManager implements the ContractTriggerXServiceManagerMethods interface.
var _ ContractTriggerXServiceManagerMethods = (*ContractTriggerXServiceManager)(nil)

// ContractTriggerXServiceManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractTriggerXServiceManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTriggerXServiceManagerCaller implements the ContractTriggerXServiceManagerCalls interface.
var _ ContractTriggerXServiceManagerCalls = (*ContractTriggerXServiceManagerCaller)(nil)

// ContractTriggerXServiceManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTriggerXServiceManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTriggerXServiceManagerTransactor implements the ContractTriggerXServiceManagerTransacts interface.
var _ ContractTriggerXServiceManagerTransacts = (*ContractTriggerXServiceManagerTransactor)(nil)

// ContractTriggerXServiceManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractTriggerXServiceManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTriggerXServiceManagerFilterer implements the ContractTriggerXServiceManagerFilters interface.
var _ ContractTriggerXServiceManagerFilters = (*ContractTriggerXServiceManagerFilterer)(nil)

// ContractTriggerXServiceManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractTriggerXServiceManagerSession struct {
	Contract     *ContractTriggerXServiceManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                   // Call options to use throughout this session
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ContractTriggerXServiceManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractTriggerXServiceManagerCallerSession struct {
	Contract *ContractTriggerXServiceManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                         // Call options to use throughout this session
}

// ContractTriggerXServiceManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTriggerXServiceManagerTransactorSession struct {
	Contract     *ContractTriggerXServiceManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                         // Transaction auth options to use throughout this session
}

// ContractTriggerXServiceManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractTriggerXServiceManagerRaw struct {
	Contract *ContractTriggerXServiceManager // Generic contract binding to access the raw methods on
}

// ContractTriggerXServiceManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractTriggerXServiceManagerCallerRaw struct {
	Contract *ContractTriggerXServiceManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTriggerXServiceManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTriggerXServiceManagerTransactorRaw struct {
	Contract *ContractTriggerXServiceManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractTriggerXServiceManager creates a new instance of ContractTriggerXServiceManager, bound to a specific deployed contract.
func NewContractTriggerXServiceManager(address common.Address, backend bind.ContractBackend) (*ContractTriggerXServiceManager, error) {
	contract, err := bindContractTriggerXServiceManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXServiceManager{ContractTriggerXServiceManagerCaller: ContractTriggerXServiceManagerCaller{contract: contract}, ContractTriggerXServiceManagerTransactor: ContractTriggerXServiceManagerTransactor{contract: contract}, ContractTriggerXServiceManagerFilterer: ContractTriggerXServiceManagerFilterer{contract: contract}}, nil
}

// NewContractTriggerXServiceManagerCaller creates a new read-only instance of ContractTriggerXServiceManager, bound to a specific deployed contract.
func NewContractTriggerXServiceManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractTriggerXServiceManagerCaller, error) {
	contract, err := bindContractTriggerXServiceManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXServiceManagerCaller{contract: contract}, nil
}

// NewContractTriggerXServiceManagerTransactor creates a new write-only instance of ContractTriggerXServiceManager, bound to a specific deployed contract.
func NewContractTriggerXServiceManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTriggerXServiceManagerTransactor, error) {
	contract, err := bindContractTriggerXServiceManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXServiceManagerTransactor{contract: contract}, nil
}

// NewContractTriggerXServiceManagerFilterer creates a new log filterer instance of ContractTriggerXServiceManager, bound to a specific deployed contract.
func NewContractTriggerXServiceManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractTriggerXServiceManagerFilterer, error) {
	contract, err := bindContractTriggerXServiceManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXServiceManagerFilterer{contract: contract}, nil
}

// bindContractTriggerXServiceManager binds a generic wrapper to an already deployed contract.
func bindContractTriggerXServiceManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractTriggerXServiceManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractTriggerXServiceManager.Contract.ContractTriggerXServiceManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.ContractTriggerXServiceManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.ContractTriggerXServiceManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractTriggerXServiceManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.contract.Transact(opts, method, params...)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCaller) AvsDirectory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTriggerXServiceManager.contract.Call(opts, &out, "avsDirectory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) AvsDirectory() (common.Address, error) {
	return _ContractTriggerXServiceManager.Contract.AvsDirectory(&_ContractTriggerXServiceManager.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCallerSession) AvsDirectory() (common.Address, error) {
	return _ContractTriggerXServiceManager.Contract.AvsDirectory(&_ContractTriggerXServiceManager.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTriggerXServiceManager.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractTriggerXServiceManager.Contract.BlsApkRegistry(&_ContractTriggerXServiceManager.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCallerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractTriggerXServiceManager.Contract.BlsApkRegistry(&_ContractTriggerXServiceManager.CallOpts)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCaller) CheckSignatures(opts *bind.CallOpts, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	var out []interface{}
	err := _ContractTriggerXServiceManager.contract.Call(opts, &out, "checkSignatures", msgHash, quorumNumbers, referenceBlockNumber, params)

	if err != nil {
		return *new(IBLSSignatureCheckerQuorumStakeTotals), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(IBLSSignatureCheckerQuorumStakeTotals)).(*IBLSSignatureCheckerQuorumStakeTotals)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractTriggerXServiceManager.Contract.CheckSignatures(&_ContractTriggerXServiceManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCallerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractTriggerXServiceManager.Contract.CheckSignatures(&_ContractTriggerXServiceManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTriggerXServiceManager.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) Delegation() (common.Address, error) {
	return _ContractTriggerXServiceManager.Contract.Delegation(&_ContractTriggerXServiceManager.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCallerSession) Delegation() (common.Address, error) {
	return _ContractTriggerXServiceManager.Contract.Delegation(&_ContractTriggerXServiceManager.CallOpts)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCaller) GetOperatorRestakedStrategies(opts *bind.CallOpts, operator common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _ContractTriggerXServiceManager.contract.Call(opts, &out, "getOperatorRestakedStrategies", operator)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _ContractTriggerXServiceManager.Contract.GetOperatorRestakedStrategies(&_ContractTriggerXServiceManager.CallOpts, operator)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCallerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _ContractTriggerXServiceManager.Contract.GetOperatorRestakedStrategies(&_ContractTriggerXServiceManager.CallOpts, operator)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCaller) GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ContractTriggerXServiceManager.contract.Call(opts, &out, "getRestakeableStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractTriggerXServiceManager.Contract.GetRestakeableStrategies(&_ContractTriggerXServiceManager.CallOpts)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCallerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractTriggerXServiceManager.Contract.GetRestakeableStrategies(&_ContractTriggerXServiceManager.CallOpts)
}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) view returns(bool)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCaller) IsBlackListed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ContractTriggerXServiceManager.contract.Call(opts, &out, "isBlackListed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) view returns(bool)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) IsBlackListed(arg0 common.Address) (bool, error) {
	return _ContractTriggerXServiceManager.Contract.IsBlackListed(&_ContractTriggerXServiceManager.CallOpts, arg0)
}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) view returns(bool)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCallerSession) IsBlackListed(arg0 common.Address) (bool, error) {
	return _ContractTriggerXServiceManager.Contract.IsBlackListed(&_ContractTriggerXServiceManager.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTriggerXServiceManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) Owner() (common.Address, error) {
	return _ContractTriggerXServiceManager.Contract.Owner(&_ContractTriggerXServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCallerSession) Owner() (common.Address, error) {
	return _ContractTriggerXServiceManager.Contract.Owner(&_ContractTriggerXServiceManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractTriggerXServiceManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) Paused(index uint8) (bool, error) {
	return _ContractTriggerXServiceManager.Contract.Paused(&_ContractTriggerXServiceManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCallerSession) Paused(index uint8) (bool, error) {
	return _ContractTriggerXServiceManager.Contract.Paused(&_ContractTriggerXServiceManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractTriggerXServiceManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) Paused0() (*big.Int, error) {
	return _ContractTriggerXServiceManager.Contract.Paused0(&_ContractTriggerXServiceManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCallerSession) Paused0() (*big.Int, error) {
	return _ContractTriggerXServiceManager.Contract.Paused0(&_ContractTriggerXServiceManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTriggerXServiceManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) PauserRegistry() (common.Address, error) {
	return _ContractTriggerXServiceManager.Contract.PauserRegistry(&_ContractTriggerXServiceManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractTriggerXServiceManager.Contract.PauserRegistry(&_ContractTriggerXServiceManager.CallOpts)
}

// QuorumManager is a free data retrieval call binding the contract method 0xb5f7eb6b.
//
// Solidity: function quorumManager() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCaller) QuorumManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTriggerXServiceManager.contract.Call(opts, &out, "quorumManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QuorumManager is a free data retrieval call binding the contract method 0xb5f7eb6b.
//
// Solidity: function quorumManager() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) QuorumManager() (common.Address, error) {
	return _ContractTriggerXServiceManager.Contract.QuorumManager(&_ContractTriggerXServiceManager.CallOpts)
}

// QuorumManager is a free data retrieval call binding the contract method 0xb5f7eb6b.
//
// Solidity: function quorumManager() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCallerSession) QuorumManager() (common.Address, error) {
	return _ContractTriggerXServiceManager.Contract.QuorumManager(&_ContractTriggerXServiceManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTriggerXServiceManager.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractTriggerXServiceManager.Contract.RegistryCoordinator(&_ContractTriggerXServiceManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCallerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractTriggerXServiceManager.Contract.RegistryCoordinator(&_ContractTriggerXServiceManager.CallOpts)
}

// RewardsInitiator is a free data retrieval call binding the contract method 0xfc299dee.
//
// Solidity: function rewardsInitiator() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCaller) RewardsInitiator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTriggerXServiceManager.contract.Call(opts, &out, "rewardsInitiator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsInitiator is a free data retrieval call binding the contract method 0xfc299dee.
//
// Solidity: function rewardsInitiator() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) RewardsInitiator() (common.Address, error) {
	return _ContractTriggerXServiceManager.Contract.RewardsInitiator(&_ContractTriggerXServiceManager.CallOpts)
}

// RewardsInitiator is a free data retrieval call binding the contract method 0xfc299dee.
//
// Solidity: function rewardsInitiator() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCallerSession) RewardsInitiator() (common.Address, error) {
	return _ContractTriggerXServiceManager.Contract.RewardsInitiator(&_ContractTriggerXServiceManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTriggerXServiceManager.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) StakeRegistry() (common.Address, error) {
	return _ContractTriggerXServiceManager.Contract.StakeRegistry(&_ContractTriggerXServiceManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCallerSession) StakeRegistry() (common.Address, error) {
	return _ContractTriggerXServiceManager.Contract.StakeRegistry(&_ContractTriggerXServiceManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCaller) StaleStakesForbidden(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractTriggerXServiceManager.contract.Call(opts, &out, "staleStakesForbidden")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) StaleStakesForbidden() (bool, error) {
	return _ContractTriggerXServiceManager.Contract.StaleStakesForbidden(&_ContractTriggerXServiceManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCallerSession) StaleStakesForbidden() (bool, error) {
	return _ContractTriggerXServiceManager.Contract.StaleStakesForbidden(&_ContractTriggerXServiceManager.CallOpts)
}

// TaskManager is a free data retrieval call binding the contract method 0xa50a640e.
//
// Solidity: function taskManager() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCaller) TaskManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTriggerXServiceManager.contract.Call(opts, &out, "taskManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TaskManager is a free data retrieval call binding the contract method 0xa50a640e.
//
// Solidity: function taskManager() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) TaskManager() (common.Address, error) {
	return _ContractTriggerXServiceManager.Contract.TaskManager(&_ContractTriggerXServiceManager.CallOpts)
}

// TaskManager is a free data retrieval call binding the contract method 0xa50a640e.
//
// Solidity: function taskManager() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCallerSession) TaskManager() (common.Address, error) {
	return _ContractTriggerXServiceManager.Contract.TaskManager(&_ContractTriggerXServiceManager.CallOpts)
}

// TaskManagerContract is a free data retrieval call binding the contract method 0x500c8dd3.
//
// Solidity: function taskManagerContract() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCaller) TaskManagerContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTriggerXServiceManager.contract.Call(opts, &out, "taskManagerContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TaskManagerContract is a free data retrieval call binding the contract method 0x500c8dd3.
//
// Solidity: function taskManagerContract() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) TaskManagerContract() (common.Address, error) {
	return _ContractTriggerXServiceManager.Contract.TaskManagerContract(&_ContractTriggerXServiceManager.CallOpts)
}

// TaskManagerContract is a free data retrieval call binding the contract method 0x500c8dd3.
//
// Solidity: function taskManagerContract() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCallerSession) TaskManagerContract() (common.Address, error) {
	return _ContractTriggerXServiceManager.Contract.TaskManagerContract(&_ContractTriggerXServiceManager.CallOpts)
}

// TaskValidator is a free data retrieval call binding the contract method 0xfd38ec8c.
//
// Solidity: function taskValidator() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCaller) TaskValidator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTriggerXServiceManager.contract.Call(opts, &out, "taskValidator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TaskValidator is a free data retrieval call binding the contract method 0xfd38ec8c.
//
// Solidity: function taskValidator() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) TaskValidator() (common.Address, error) {
	return _ContractTriggerXServiceManager.Contract.TaskValidator(&_ContractTriggerXServiceManager.CallOpts)
}

// TaskValidator is a free data retrieval call binding the contract method 0xfd38ec8c.
//
// Solidity: function taskValidator() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCallerSession) TaskValidator() (common.Address, error) {
	return _ContractTriggerXServiceManager.Contract.TaskValidator(&_ContractTriggerXServiceManager.CallOpts)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCaller) TrySignatureAndApkVerification(opts *bind.CallOpts, msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	var out []interface{}
	err := _ContractTriggerXServiceManager.contract.Call(opts, &out, "trySignatureAndApkVerification", msgHash, apk, apkG2, sigma)

	outstruct := new(struct {
		PairingSuccessful bool
		SiganatureIsValid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PairingSuccessful = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.SiganatureIsValid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractTriggerXServiceManager.Contract.TrySignatureAndApkVerification(&_ContractTriggerXServiceManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCallerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractTriggerXServiceManager.Contract.TrySignatureAndApkVerification(&_ContractTriggerXServiceManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// BlacklistKeeper is a paid mutator transaction binding the contract method 0x26a965f0.
//
// Solidity: function blacklistKeeper(address _operator) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactor) BlacklistKeeper(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.contract.Transact(opts, "blacklistKeeper", _operator)
}

// BlacklistKeeper is a paid mutator transaction binding the contract method 0x26a965f0.
//
// Solidity: function blacklistKeeper(address _operator) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) BlacklistKeeper(_operator common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.BlacklistKeeper(&_ContractTriggerXServiceManager.TransactOpts, _operator)
}

// BlacklistKeeper is a paid mutator transaction binding the contract method 0x26a965f0.
//
// Solidity: function blacklistKeeper(address _operator) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactorSession) BlacklistKeeper(_operator common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.BlacklistKeeper(&_ContractTriggerXServiceManager.TransactOpts, _operator)
}

// CreateAVSRewardsSubmission is a paid mutator transaction binding the contract method 0xfce36c7d.
//
// Solidity: function createAVSRewardsSubmission(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactor) CreateAVSRewardsSubmission(opts *bind.TransactOpts, rewardsSubmissions []IRewardsCoordinatorRewardsSubmission) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.contract.Transact(opts, "createAVSRewardsSubmission", rewardsSubmissions)
}

// CreateAVSRewardsSubmission is a paid mutator transaction binding the contract method 0xfce36c7d.
//
// Solidity: function createAVSRewardsSubmission(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) CreateAVSRewardsSubmission(rewardsSubmissions []IRewardsCoordinatorRewardsSubmission) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.CreateAVSRewardsSubmission(&_ContractTriggerXServiceManager.TransactOpts, rewardsSubmissions)
}

// CreateAVSRewardsSubmission is a paid mutator transaction binding the contract method 0xfce36c7d.
//
// Solidity: function createAVSRewardsSubmission(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactorSession) CreateAVSRewardsSubmission(rewardsSubmissions []IRewardsCoordinatorRewardsSubmission) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.CreateAVSRewardsSubmission(&_ContractTriggerXServiceManager.TransactOpts, rewardsSubmissions)
}

// CreateOperatorDirectedAVSRewardsSubmission is a paid mutator transaction binding the contract method 0xa20b99bf.
//
// Solidity: function createOperatorDirectedAVSRewardsSubmission(((address,uint96)[],address,(address,uint256)[],uint32,uint32,string)[] operatorDirectedRewardsSubmissions) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactor) CreateOperatorDirectedAVSRewardsSubmission(opts *bind.TransactOpts, operatorDirectedRewardsSubmissions []IRewardsCoordinatorOperatorDirectedRewardsSubmission) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.contract.Transact(opts, "createOperatorDirectedAVSRewardsSubmission", operatorDirectedRewardsSubmissions)
}

// CreateOperatorDirectedAVSRewardsSubmission is a paid mutator transaction binding the contract method 0xa20b99bf.
//
// Solidity: function createOperatorDirectedAVSRewardsSubmission(((address,uint96)[],address,(address,uint256)[],uint32,uint32,string)[] operatorDirectedRewardsSubmissions) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) CreateOperatorDirectedAVSRewardsSubmission(operatorDirectedRewardsSubmissions []IRewardsCoordinatorOperatorDirectedRewardsSubmission) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.CreateOperatorDirectedAVSRewardsSubmission(&_ContractTriggerXServiceManager.TransactOpts, operatorDirectedRewardsSubmissions)
}

// CreateOperatorDirectedAVSRewardsSubmission is a paid mutator transaction binding the contract method 0xa20b99bf.
//
// Solidity: function createOperatorDirectedAVSRewardsSubmission(((address,uint96)[],address,(address,uint256)[],uint32,uint32,string)[] operatorDirectedRewardsSubmissions) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactorSession) CreateOperatorDirectedAVSRewardsSubmission(operatorDirectedRewardsSubmissions []IRewardsCoordinatorOperatorDirectedRewardsSubmission) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.CreateOperatorDirectedAVSRewardsSubmission(&_ContractTriggerXServiceManager.TransactOpts, operatorDirectedRewardsSubmissions)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.contract.Transact(opts, "deregisterOperatorFromAVS", operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.DeregisterOperatorFromAVS(&_ContractTriggerXServiceManager.TransactOpts, operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactorSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.DeregisterOperatorFromAVS(&_ContractTriggerXServiceManager.TransactOpts, operator)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address _taskManagerContract, address initialOwner, address rewardsInitiator, address _taskManager, address _taskValidator, address _quorumManager) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactor) Initialize(opts *bind.TransactOpts, _taskManagerContract common.Address, initialOwner common.Address, rewardsInitiator common.Address, _taskManager common.Address, _taskValidator common.Address, _quorumManager common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.contract.Transact(opts, "initialize", _taskManagerContract, initialOwner, rewardsInitiator, _taskManager, _taskValidator, _quorumManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address _taskManagerContract, address initialOwner, address rewardsInitiator, address _taskManager, address _taskValidator, address _quorumManager) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) Initialize(_taskManagerContract common.Address, initialOwner common.Address, rewardsInitiator common.Address, _taskManager common.Address, _taskValidator common.Address, _quorumManager common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.Initialize(&_ContractTriggerXServiceManager.TransactOpts, _taskManagerContract, initialOwner, rewardsInitiator, _taskManager, _taskValidator, _quorumManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address _taskManagerContract, address initialOwner, address rewardsInitiator, address _taskManager, address _taskValidator, address _quorumManager) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactorSession) Initialize(_taskManagerContract common.Address, initialOwner common.Address, rewardsInitiator common.Address, _taskManager common.Address, _taskValidator common.Address, _quorumManager common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.Initialize(&_ContractTriggerXServiceManager.TransactOpts, _taskManagerContract, initialOwner, rewardsInitiator, _taskManager, _taskValidator, _quorumManager)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.Pause(&_ContractTriggerXServiceManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.Pause(&_ContractTriggerXServiceManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) PauseAll() (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.PauseAll(&_ContractTriggerXServiceManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.PauseAll(&_ContractTriggerXServiceManager.TransactOpts)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.contract.Transact(opts, "registerOperatorToAVS", operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.RegisterOperatorToAVS(&_ContractTriggerXServiceManager.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactorSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.RegisterOperatorToAVS(&_ContractTriggerXServiceManager.TransactOpts, operator, operatorSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.RenounceOwnership(&_ContractTriggerXServiceManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.RenounceOwnership(&_ContractTriggerXServiceManager.TransactOpts)
}

// SetClaimerFor is a paid mutator transaction binding the contract method 0xa0169ddd.
//
// Solidity: function setClaimerFor(address claimer) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactor) SetClaimerFor(opts *bind.TransactOpts, claimer common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.contract.Transact(opts, "setClaimerFor", claimer)
}

// SetClaimerFor is a paid mutator transaction binding the contract method 0xa0169ddd.
//
// Solidity: function setClaimerFor(address claimer) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) SetClaimerFor(claimer common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.SetClaimerFor(&_ContractTriggerXServiceManager.TransactOpts, claimer)
}

// SetClaimerFor is a paid mutator transaction binding the contract method 0xa0169ddd.
//
// Solidity: function setClaimerFor(address claimer) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactorSession) SetClaimerFor(claimer common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.SetClaimerFor(&_ContractTriggerXServiceManager.TransactOpts, claimer)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.SetPauserRegistry(&_ContractTriggerXServiceManager.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.SetPauserRegistry(&_ContractTriggerXServiceManager.TransactOpts, newPauserRegistry)
}

// SetQuorumManager is a paid mutator transaction binding the contract method 0x9cd0fb86.
//
// Solidity: function setQuorumManager(address _quorumManager) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactor) SetQuorumManager(opts *bind.TransactOpts, _quorumManager common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.contract.Transact(opts, "setQuorumManager", _quorumManager)
}

// SetQuorumManager is a paid mutator transaction binding the contract method 0x9cd0fb86.
//
// Solidity: function setQuorumManager(address _quorumManager) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) SetQuorumManager(_quorumManager common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.SetQuorumManager(&_ContractTriggerXServiceManager.TransactOpts, _quorumManager)
}

// SetQuorumManager is a paid mutator transaction binding the contract method 0x9cd0fb86.
//
// Solidity: function setQuorumManager(address _quorumManager) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactorSession) SetQuorumManager(_quorumManager common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.SetQuorumManager(&_ContractTriggerXServiceManager.TransactOpts, _quorumManager)
}

// SetRewardsInitiator is a paid mutator transaction binding the contract method 0x3bc28c8c.
//
// Solidity: function setRewardsInitiator(address newRewardsInitiator) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactor) SetRewardsInitiator(opts *bind.TransactOpts, newRewardsInitiator common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.contract.Transact(opts, "setRewardsInitiator", newRewardsInitiator)
}

// SetRewardsInitiator is a paid mutator transaction binding the contract method 0x3bc28c8c.
//
// Solidity: function setRewardsInitiator(address newRewardsInitiator) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) SetRewardsInitiator(newRewardsInitiator common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.SetRewardsInitiator(&_ContractTriggerXServiceManager.TransactOpts, newRewardsInitiator)
}

// SetRewardsInitiator is a paid mutator transaction binding the contract method 0x3bc28c8c.
//
// Solidity: function setRewardsInitiator(address newRewardsInitiator) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactorSession) SetRewardsInitiator(newRewardsInitiator common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.SetRewardsInitiator(&_ContractTriggerXServiceManager.TransactOpts, newRewardsInitiator)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactor) SetStaleStakesForbidden(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.contract.Transact(opts, "setStaleStakesForbidden", value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.SetStaleStakesForbidden(&_ContractTriggerXServiceManager.TransactOpts, value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactorSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.SetStaleStakesForbidden(&_ContractTriggerXServiceManager.TransactOpts, value)
}

// SetTaskManager is a paid mutator transaction binding the contract method 0x327d0a60.
//
// Solidity: function setTaskManager(address _taskManager) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactor) SetTaskManager(opts *bind.TransactOpts, _taskManager common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.contract.Transact(opts, "setTaskManager", _taskManager)
}

// SetTaskManager is a paid mutator transaction binding the contract method 0x327d0a60.
//
// Solidity: function setTaskManager(address _taskManager) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) SetTaskManager(_taskManager common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.SetTaskManager(&_ContractTriggerXServiceManager.TransactOpts, _taskManager)
}

// SetTaskManager is a paid mutator transaction binding the contract method 0x327d0a60.
//
// Solidity: function setTaskManager(address _taskManager) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactorSession) SetTaskManager(_taskManager common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.SetTaskManager(&_ContractTriggerXServiceManager.TransactOpts, _taskManager)
}

// SetTaskValidator is a paid mutator transaction binding the contract method 0x5edd78ea.
//
// Solidity: function setTaskValidator(address _taskValidator) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactor) SetTaskValidator(opts *bind.TransactOpts, _taskValidator common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.contract.Transact(opts, "setTaskValidator", _taskValidator)
}

// SetTaskValidator is a paid mutator transaction binding the contract method 0x5edd78ea.
//
// Solidity: function setTaskValidator(address _taskValidator) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) SetTaskValidator(_taskValidator common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.SetTaskValidator(&_ContractTriggerXServiceManager.TransactOpts, _taskValidator)
}

// SetTaskValidator is a paid mutator transaction binding the contract method 0x5edd78ea.
//
// Solidity: function setTaskValidator(address _taskValidator) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactorSession) SetTaskValidator(_taskValidator common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.SetTaskValidator(&_ContractTriggerXServiceManager.TransactOpts, _taskValidator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.TransferOwnership(&_ContractTriggerXServiceManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.TransferOwnership(&_ContractTriggerXServiceManager.TransactOpts, newOwner)
}

// UnblacklistKeeper is a paid mutator transaction binding the contract method 0xa41d3f94.
//
// Solidity: function unblacklistKeeper(address _operator) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactor) UnblacklistKeeper(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.contract.Transact(opts, "unblacklistKeeper", _operator)
}

// UnblacklistKeeper is a paid mutator transaction binding the contract method 0xa41d3f94.
//
// Solidity: function unblacklistKeeper(address _operator) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) UnblacklistKeeper(_operator common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.UnblacklistKeeper(&_ContractTriggerXServiceManager.TransactOpts, _operator)
}

// UnblacklistKeeper is a paid mutator transaction binding the contract method 0xa41d3f94.
//
// Solidity: function unblacklistKeeper(address _operator) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactorSession) UnblacklistKeeper(_operator common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.UnblacklistKeeper(&_ContractTriggerXServiceManager.TransactOpts, _operator)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.Unpause(&_ContractTriggerXServiceManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.Unpause(&_ContractTriggerXServiceManager.TransactOpts, newPausedStatus)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.contract.Transact(opts, "updateAVSMetadataURI", _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.UpdateAVSMetadataURI(&_ContractTriggerXServiceManager.TransactOpts, _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactorSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.UpdateAVSMetadataURI(&_ContractTriggerXServiceManager.TransactOpts, _metadataURI)
}

// UpdateTaskManager is a paid mutator transaction binding the contract method 0x58ac4a1e.
//
// Solidity: function updateTaskManager(address _taskManager) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactor) UpdateTaskManager(opts *bind.TransactOpts, _taskManager common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.contract.Transact(opts, "updateTaskManager", _taskManager)
}

// UpdateTaskManager is a paid mutator transaction binding the contract method 0x58ac4a1e.
//
// Solidity: function updateTaskManager(address _taskManager) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) UpdateTaskManager(_taskManager common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.UpdateTaskManager(&_ContractTriggerXServiceManager.TransactOpts, _taskManager)
}

// UpdateTaskManager is a paid mutator transaction binding the contract method 0x58ac4a1e.
//
// Solidity: function updateTaskManager(address _taskManager) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactorSession) UpdateTaskManager(_taskManager common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.UpdateTaskManager(&_ContractTriggerXServiceManager.TransactOpts, _taskManager)
}

// ContractTriggerXServiceManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractTriggerXServiceManager contract.
type ContractTriggerXServiceManagerInitializedIterator struct {
	Event *ContractTriggerXServiceManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ContractTriggerXServiceManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerXServiceManagerInitialized)
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
		it.Event = new(ContractTriggerXServiceManagerInitialized)
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
func (it *ContractTriggerXServiceManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerXServiceManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerXServiceManagerInitialized represents a Initialized event raised by the ContractTriggerXServiceManager contract.
type ContractTriggerXServiceManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractTriggerXServiceManagerInitializedIterator, error) {

	logs, sub, err := _ContractTriggerXServiceManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXServiceManagerInitializedIterator{contract: _ContractTriggerXServiceManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractTriggerXServiceManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractTriggerXServiceManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerXServiceManagerInitialized)
				if err := _ContractTriggerXServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) ParseInitialized(log types.Log) (*ContractTriggerXServiceManagerInitialized, error) {
	event := new(ContractTriggerXServiceManagerInitialized)
	if err := _ContractTriggerXServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTriggerXServiceManagerKeeperAddedIterator is returned from FilterKeeperAdded and is used to iterate over the raw logs and unpacked data for KeeperAdded events raised by the ContractTriggerXServiceManager contract.
type ContractTriggerXServiceManagerKeeperAddedIterator struct {
	Event *ContractTriggerXServiceManagerKeeperAdded // Event containing the contract specifics and raw log

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
func (it *ContractTriggerXServiceManagerKeeperAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerXServiceManagerKeeperAdded)
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
		it.Event = new(ContractTriggerXServiceManagerKeeperAdded)
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
func (it *ContractTriggerXServiceManagerKeeperAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerXServiceManagerKeeperAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerXServiceManagerKeeperAdded represents a KeeperAdded event raised by the ContractTriggerXServiceManager contract.
type ContractTriggerXServiceManagerKeeperAdded struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterKeeperAdded is a free log retrieval operation binding the contract event 0x1584773458d98c71b34a270ee1100b3a42889bf91e3b7a858563b684c24d838e.
//
// Solidity: event KeeperAdded(address indexed operator)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) FilterKeeperAdded(opts *bind.FilterOpts, operator []common.Address) (*ContractTriggerXServiceManagerKeeperAddedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractTriggerXServiceManager.contract.FilterLogs(opts, "KeeperAdded", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXServiceManagerKeeperAddedIterator{contract: _ContractTriggerXServiceManager.contract, event: "KeeperAdded", logs: logs, sub: sub}, nil
}

// WatchKeeperAdded is a free log subscription operation binding the contract event 0x1584773458d98c71b34a270ee1100b3a42889bf91e3b7a858563b684c24d838e.
//
// Solidity: event KeeperAdded(address indexed operator)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) WatchKeeperAdded(opts *bind.WatchOpts, sink chan<- *ContractTriggerXServiceManagerKeeperAdded, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractTriggerXServiceManager.contract.WatchLogs(opts, "KeeperAdded", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerXServiceManagerKeeperAdded)
				if err := _ContractTriggerXServiceManager.contract.UnpackLog(event, "KeeperAdded", log); err != nil {
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

// ParseKeeperAdded is a log parse operation binding the contract event 0x1584773458d98c71b34a270ee1100b3a42889bf91e3b7a858563b684c24d838e.
//
// Solidity: event KeeperAdded(address indexed operator)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) ParseKeeperAdded(log types.Log) (*ContractTriggerXServiceManagerKeeperAdded, error) {
	event := new(ContractTriggerXServiceManagerKeeperAdded)
	if err := _ContractTriggerXServiceManager.contract.UnpackLog(event, "KeeperAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTriggerXServiceManagerKeeperBlacklistedIterator is returned from FilterKeeperBlacklisted and is used to iterate over the raw logs and unpacked data for KeeperBlacklisted events raised by the ContractTriggerXServiceManager contract.
type ContractTriggerXServiceManagerKeeperBlacklistedIterator struct {
	Event *ContractTriggerXServiceManagerKeeperBlacklisted // Event containing the contract specifics and raw log

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
func (it *ContractTriggerXServiceManagerKeeperBlacklistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerXServiceManagerKeeperBlacklisted)
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
		it.Event = new(ContractTriggerXServiceManagerKeeperBlacklisted)
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
func (it *ContractTriggerXServiceManagerKeeperBlacklistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerXServiceManagerKeeperBlacklistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerXServiceManagerKeeperBlacklisted represents a KeeperBlacklisted event raised by the ContractTriggerXServiceManager contract.
type ContractTriggerXServiceManagerKeeperBlacklisted struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterKeeperBlacklisted is a free log retrieval operation binding the contract event 0x79bf5efab685b39bdd2b7b1aba09ffa5e9dfbd86f7aaebab4b52d7ef679d2844.
//
// Solidity: event KeeperBlacklisted(address indexed operator)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) FilterKeeperBlacklisted(opts *bind.FilterOpts, operator []common.Address) (*ContractTriggerXServiceManagerKeeperBlacklistedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractTriggerXServiceManager.contract.FilterLogs(opts, "KeeperBlacklisted", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXServiceManagerKeeperBlacklistedIterator{contract: _ContractTriggerXServiceManager.contract, event: "KeeperBlacklisted", logs: logs, sub: sub}, nil
}

// WatchKeeperBlacklisted is a free log subscription operation binding the contract event 0x79bf5efab685b39bdd2b7b1aba09ffa5e9dfbd86f7aaebab4b52d7ef679d2844.
//
// Solidity: event KeeperBlacklisted(address indexed operator)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) WatchKeeperBlacklisted(opts *bind.WatchOpts, sink chan<- *ContractTriggerXServiceManagerKeeperBlacklisted, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractTriggerXServiceManager.contract.WatchLogs(opts, "KeeperBlacklisted", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerXServiceManagerKeeperBlacklisted)
				if err := _ContractTriggerXServiceManager.contract.UnpackLog(event, "KeeperBlacklisted", log); err != nil {
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

// ParseKeeperBlacklisted is a log parse operation binding the contract event 0x79bf5efab685b39bdd2b7b1aba09ffa5e9dfbd86f7aaebab4b52d7ef679d2844.
//
// Solidity: event KeeperBlacklisted(address indexed operator)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) ParseKeeperBlacklisted(log types.Log) (*ContractTriggerXServiceManagerKeeperBlacklisted, error) {
	event := new(ContractTriggerXServiceManagerKeeperBlacklisted)
	if err := _ContractTriggerXServiceManager.contract.UnpackLog(event, "KeeperBlacklisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTriggerXServiceManagerKeeperRemovedIterator is returned from FilterKeeperRemoved and is used to iterate over the raw logs and unpacked data for KeeperRemoved events raised by the ContractTriggerXServiceManager contract.
type ContractTriggerXServiceManagerKeeperRemovedIterator struct {
	Event *ContractTriggerXServiceManagerKeeperRemoved // Event containing the contract specifics and raw log

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
func (it *ContractTriggerXServiceManagerKeeperRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerXServiceManagerKeeperRemoved)
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
		it.Event = new(ContractTriggerXServiceManagerKeeperRemoved)
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
func (it *ContractTriggerXServiceManagerKeeperRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerXServiceManagerKeeperRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerXServiceManagerKeeperRemoved represents a KeeperRemoved event raised by the ContractTriggerXServiceManager contract.
type ContractTriggerXServiceManagerKeeperRemoved struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterKeeperRemoved is a free log retrieval operation binding the contract event 0xa7a775c2c8141f7985c111748ec31c11e5e44b83528e105c8d1d4e8e6b81cf80.
//
// Solidity: event KeeperRemoved(address indexed operator)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) FilterKeeperRemoved(opts *bind.FilterOpts, operator []common.Address) (*ContractTriggerXServiceManagerKeeperRemovedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractTriggerXServiceManager.contract.FilterLogs(opts, "KeeperRemoved", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXServiceManagerKeeperRemovedIterator{contract: _ContractTriggerXServiceManager.contract, event: "KeeperRemoved", logs: logs, sub: sub}, nil
}

// WatchKeeperRemoved is a free log subscription operation binding the contract event 0xa7a775c2c8141f7985c111748ec31c11e5e44b83528e105c8d1d4e8e6b81cf80.
//
// Solidity: event KeeperRemoved(address indexed operator)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) WatchKeeperRemoved(opts *bind.WatchOpts, sink chan<- *ContractTriggerXServiceManagerKeeperRemoved, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractTriggerXServiceManager.contract.WatchLogs(opts, "KeeperRemoved", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerXServiceManagerKeeperRemoved)
				if err := _ContractTriggerXServiceManager.contract.UnpackLog(event, "KeeperRemoved", log); err != nil {
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

// ParseKeeperRemoved is a log parse operation binding the contract event 0xa7a775c2c8141f7985c111748ec31c11e5e44b83528e105c8d1d4e8e6b81cf80.
//
// Solidity: event KeeperRemoved(address indexed operator)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) ParseKeeperRemoved(log types.Log) (*ContractTriggerXServiceManagerKeeperRemoved, error) {
	event := new(ContractTriggerXServiceManagerKeeperRemoved)
	if err := _ContractTriggerXServiceManager.contract.UnpackLog(event, "KeeperRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTriggerXServiceManagerKeeperUnblacklistedIterator is returned from FilterKeeperUnblacklisted and is used to iterate over the raw logs and unpacked data for KeeperUnblacklisted events raised by the ContractTriggerXServiceManager contract.
type ContractTriggerXServiceManagerKeeperUnblacklistedIterator struct {
	Event *ContractTriggerXServiceManagerKeeperUnblacklisted // Event containing the contract specifics and raw log

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
func (it *ContractTriggerXServiceManagerKeeperUnblacklistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerXServiceManagerKeeperUnblacklisted)
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
		it.Event = new(ContractTriggerXServiceManagerKeeperUnblacklisted)
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
func (it *ContractTriggerXServiceManagerKeeperUnblacklistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerXServiceManagerKeeperUnblacklistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerXServiceManagerKeeperUnblacklisted represents a KeeperUnblacklisted event raised by the ContractTriggerXServiceManager contract.
type ContractTriggerXServiceManagerKeeperUnblacklisted struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterKeeperUnblacklisted is a free log retrieval operation binding the contract event 0x8511046d95f95689c89b4807d91820adbc266db891b2c5c25a42f812781dde57.
//
// Solidity: event KeeperUnblacklisted(address indexed operator)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) FilterKeeperUnblacklisted(opts *bind.FilterOpts, operator []common.Address) (*ContractTriggerXServiceManagerKeeperUnblacklistedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractTriggerXServiceManager.contract.FilterLogs(opts, "KeeperUnblacklisted", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXServiceManagerKeeperUnblacklistedIterator{contract: _ContractTriggerXServiceManager.contract, event: "KeeperUnblacklisted", logs: logs, sub: sub}, nil
}

// WatchKeeperUnblacklisted is a free log subscription operation binding the contract event 0x8511046d95f95689c89b4807d91820adbc266db891b2c5c25a42f812781dde57.
//
// Solidity: event KeeperUnblacklisted(address indexed operator)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) WatchKeeperUnblacklisted(opts *bind.WatchOpts, sink chan<- *ContractTriggerXServiceManagerKeeperUnblacklisted, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractTriggerXServiceManager.contract.WatchLogs(opts, "KeeperUnblacklisted", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerXServiceManagerKeeperUnblacklisted)
				if err := _ContractTriggerXServiceManager.contract.UnpackLog(event, "KeeperUnblacklisted", log); err != nil {
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

// ParseKeeperUnblacklisted is a log parse operation binding the contract event 0x8511046d95f95689c89b4807d91820adbc266db891b2c5c25a42f812781dde57.
//
// Solidity: event KeeperUnblacklisted(address indexed operator)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) ParseKeeperUnblacklisted(log types.Log) (*ContractTriggerXServiceManagerKeeperUnblacklisted, error) {
	event := new(ContractTriggerXServiceManagerKeeperUnblacklisted)
	if err := _ContractTriggerXServiceManager.contract.UnpackLog(event, "KeeperUnblacklisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTriggerXServiceManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractTriggerXServiceManager contract.
type ContractTriggerXServiceManagerOwnershipTransferredIterator struct {
	Event *ContractTriggerXServiceManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractTriggerXServiceManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerXServiceManagerOwnershipTransferred)
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
		it.Event = new(ContractTriggerXServiceManagerOwnershipTransferred)
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
func (it *ContractTriggerXServiceManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerXServiceManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerXServiceManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractTriggerXServiceManager contract.
type ContractTriggerXServiceManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractTriggerXServiceManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractTriggerXServiceManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXServiceManagerOwnershipTransferredIterator{contract: _ContractTriggerXServiceManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractTriggerXServiceManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractTriggerXServiceManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerXServiceManagerOwnershipTransferred)
				if err := _ContractTriggerXServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractTriggerXServiceManagerOwnershipTransferred, error) {
	event := new(ContractTriggerXServiceManagerOwnershipTransferred)
	if err := _ContractTriggerXServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTriggerXServiceManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractTriggerXServiceManager contract.
type ContractTriggerXServiceManagerPausedIterator struct {
	Event *ContractTriggerXServiceManagerPaused // Event containing the contract specifics and raw log

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
func (it *ContractTriggerXServiceManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerXServiceManagerPaused)
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
		it.Event = new(ContractTriggerXServiceManagerPaused)
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
func (it *ContractTriggerXServiceManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerXServiceManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerXServiceManagerPaused represents a Paused event raised by the ContractTriggerXServiceManager contract.
type ContractTriggerXServiceManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractTriggerXServiceManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractTriggerXServiceManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXServiceManagerPausedIterator{contract: _ContractTriggerXServiceManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractTriggerXServiceManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractTriggerXServiceManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerXServiceManagerPaused)
				if err := _ContractTriggerXServiceManager.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) ParsePaused(log types.Log) (*ContractTriggerXServiceManagerPaused, error) {
	event := new(ContractTriggerXServiceManagerPaused)
	if err := _ContractTriggerXServiceManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTriggerXServiceManagerPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the ContractTriggerXServiceManager contract.
type ContractTriggerXServiceManagerPauserRegistrySetIterator struct {
	Event *ContractTriggerXServiceManagerPauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *ContractTriggerXServiceManagerPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerXServiceManagerPauserRegistrySet)
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
		it.Event = new(ContractTriggerXServiceManagerPauserRegistrySet)
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
func (it *ContractTriggerXServiceManagerPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerXServiceManagerPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerXServiceManagerPauserRegistrySet represents a PauserRegistrySet event raised by the ContractTriggerXServiceManager contract.
type ContractTriggerXServiceManagerPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractTriggerXServiceManagerPauserRegistrySetIterator, error) {

	logs, sub, err := _ContractTriggerXServiceManager.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXServiceManagerPauserRegistrySetIterator{contract: _ContractTriggerXServiceManager.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractTriggerXServiceManagerPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _ContractTriggerXServiceManager.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerXServiceManagerPauserRegistrySet)
				if err := _ContractTriggerXServiceManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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

// ParsePauserRegistrySet is a log parse operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) ParsePauserRegistrySet(log types.Log) (*ContractTriggerXServiceManagerPauserRegistrySet, error) {
	event := new(ContractTriggerXServiceManagerPauserRegistrySet)
	if err := _ContractTriggerXServiceManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTriggerXServiceManagerQuorumManagerSetIterator is returned from FilterQuorumManagerSet and is used to iterate over the raw logs and unpacked data for QuorumManagerSet events raised by the ContractTriggerXServiceManager contract.
type ContractTriggerXServiceManagerQuorumManagerSetIterator struct {
	Event *ContractTriggerXServiceManagerQuorumManagerSet // Event containing the contract specifics and raw log

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
func (it *ContractTriggerXServiceManagerQuorumManagerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerXServiceManagerQuorumManagerSet)
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
		it.Event = new(ContractTriggerXServiceManagerQuorumManagerSet)
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
func (it *ContractTriggerXServiceManagerQuorumManagerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerXServiceManagerQuorumManagerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerXServiceManagerQuorumManagerSet represents a QuorumManagerSet event raised by the ContractTriggerXServiceManager contract.
type ContractTriggerXServiceManagerQuorumManagerSet struct {
	OldQuorumManager common.Address
	NewQuorumManager common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterQuorumManagerSet is a free log retrieval operation binding the contract event 0x5d68f69aebb8a1d75f8b2a87dddd6e8db842f9fa3bb63745ea7f49ea5bdf03a9.
//
// Solidity: event QuorumManagerSet(address indexed oldQuorumManager, address indexed newQuorumManager)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) FilterQuorumManagerSet(opts *bind.FilterOpts, oldQuorumManager []common.Address, newQuorumManager []common.Address) (*ContractTriggerXServiceManagerQuorumManagerSetIterator, error) {

	var oldQuorumManagerRule []interface{}
	for _, oldQuorumManagerItem := range oldQuorumManager {
		oldQuorumManagerRule = append(oldQuorumManagerRule, oldQuorumManagerItem)
	}
	var newQuorumManagerRule []interface{}
	for _, newQuorumManagerItem := range newQuorumManager {
		newQuorumManagerRule = append(newQuorumManagerRule, newQuorumManagerItem)
	}

	logs, sub, err := _ContractTriggerXServiceManager.contract.FilterLogs(opts, "QuorumManagerSet", oldQuorumManagerRule, newQuorumManagerRule)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXServiceManagerQuorumManagerSetIterator{contract: _ContractTriggerXServiceManager.contract, event: "QuorumManagerSet", logs: logs, sub: sub}, nil
}

// WatchQuorumManagerSet is a free log subscription operation binding the contract event 0x5d68f69aebb8a1d75f8b2a87dddd6e8db842f9fa3bb63745ea7f49ea5bdf03a9.
//
// Solidity: event QuorumManagerSet(address indexed oldQuorumManager, address indexed newQuorumManager)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) WatchQuorumManagerSet(opts *bind.WatchOpts, sink chan<- *ContractTriggerXServiceManagerQuorumManagerSet, oldQuorumManager []common.Address, newQuorumManager []common.Address) (event.Subscription, error) {

	var oldQuorumManagerRule []interface{}
	for _, oldQuorumManagerItem := range oldQuorumManager {
		oldQuorumManagerRule = append(oldQuorumManagerRule, oldQuorumManagerItem)
	}
	var newQuorumManagerRule []interface{}
	for _, newQuorumManagerItem := range newQuorumManager {
		newQuorumManagerRule = append(newQuorumManagerRule, newQuorumManagerItem)
	}

	logs, sub, err := _ContractTriggerXServiceManager.contract.WatchLogs(opts, "QuorumManagerSet", oldQuorumManagerRule, newQuorumManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerXServiceManagerQuorumManagerSet)
				if err := _ContractTriggerXServiceManager.contract.UnpackLog(event, "QuorumManagerSet", log); err != nil {
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

// ParseQuorumManagerSet is a log parse operation binding the contract event 0x5d68f69aebb8a1d75f8b2a87dddd6e8db842f9fa3bb63745ea7f49ea5bdf03a9.
//
// Solidity: event QuorumManagerSet(address indexed oldQuorumManager, address indexed newQuorumManager)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) ParseQuorumManagerSet(log types.Log) (*ContractTriggerXServiceManagerQuorumManagerSet, error) {
	event := new(ContractTriggerXServiceManagerQuorumManagerSet)
	if err := _ContractTriggerXServiceManager.contract.UnpackLog(event, "QuorumManagerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTriggerXServiceManagerRewardsInitiatorUpdatedIterator is returned from FilterRewardsInitiatorUpdated and is used to iterate over the raw logs and unpacked data for RewardsInitiatorUpdated events raised by the ContractTriggerXServiceManager contract.
type ContractTriggerXServiceManagerRewardsInitiatorUpdatedIterator struct {
	Event *ContractTriggerXServiceManagerRewardsInitiatorUpdated // Event containing the contract specifics and raw log

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
func (it *ContractTriggerXServiceManagerRewardsInitiatorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerXServiceManagerRewardsInitiatorUpdated)
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
		it.Event = new(ContractTriggerXServiceManagerRewardsInitiatorUpdated)
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
func (it *ContractTriggerXServiceManagerRewardsInitiatorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerXServiceManagerRewardsInitiatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerXServiceManagerRewardsInitiatorUpdated represents a RewardsInitiatorUpdated event raised by the ContractTriggerXServiceManager contract.
type ContractTriggerXServiceManagerRewardsInitiatorUpdated struct {
	PrevRewardsInitiator common.Address
	NewRewardsInitiator  common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterRewardsInitiatorUpdated is a free log retrieval operation binding the contract event 0xe11cddf1816a43318ca175bbc52cd0185436e9cbead7c83acc54a73e461717e3.
//
// Solidity: event RewardsInitiatorUpdated(address prevRewardsInitiator, address newRewardsInitiator)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) FilterRewardsInitiatorUpdated(opts *bind.FilterOpts) (*ContractTriggerXServiceManagerRewardsInitiatorUpdatedIterator, error) {

	logs, sub, err := _ContractTriggerXServiceManager.contract.FilterLogs(opts, "RewardsInitiatorUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXServiceManagerRewardsInitiatorUpdatedIterator{contract: _ContractTriggerXServiceManager.contract, event: "RewardsInitiatorUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardsInitiatorUpdated is a free log subscription operation binding the contract event 0xe11cddf1816a43318ca175bbc52cd0185436e9cbead7c83acc54a73e461717e3.
//
// Solidity: event RewardsInitiatorUpdated(address prevRewardsInitiator, address newRewardsInitiator)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) WatchRewardsInitiatorUpdated(opts *bind.WatchOpts, sink chan<- *ContractTriggerXServiceManagerRewardsInitiatorUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractTriggerXServiceManager.contract.WatchLogs(opts, "RewardsInitiatorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerXServiceManagerRewardsInitiatorUpdated)
				if err := _ContractTriggerXServiceManager.contract.UnpackLog(event, "RewardsInitiatorUpdated", log); err != nil {
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

// ParseRewardsInitiatorUpdated is a log parse operation binding the contract event 0xe11cddf1816a43318ca175bbc52cd0185436e9cbead7c83acc54a73e461717e3.
//
// Solidity: event RewardsInitiatorUpdated(address prevRewardsInitiator, address newRewardsInitiator)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) ParseRewardsInitiatorUpdated(log types.Log) (*ContractTriggerXServiceManagerRewardsInitiatorUpdated, error) {
	event := new(ContractTriggerXServiceManagerRewardsInitiatorUpdated)
	if err := _ContractTriggerXServiceManager.contract.UnpackLog(event, "RewardsInitiatorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTriggerXServiceManagerStaleStakesForbiddenUpdateIterator is returned from FilterStaleStakesForbiddenUpdate and is used to iterate over the raw logs and unpacked data for StaleStakesForbiddenUpdate events raised by the ContractTriggerXServiceManager contract.
type ContractTriggerXServiceManagerStaleStakesForbiddenUpdateIterator struct {
	Event *ContractTriggerXServiceManagerStaleStakesForbiddenUpdate // Event containing the contract specifics and raw log

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
func (it *ContractTriggerXServiceManagerStaleStakesForbiddenUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerXServiceManagerStaleStakesForbiddenUpdate)
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
		it.Event = new(ContractTriggerXServiceManagerStaleStakesForbiddenUpdate)
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
func (it *ContractTriggerXServiceManagerStaleStakesForbiddenUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerXServiceManagerStaleStakesForbiddenUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerXServiceManagerStaleStakesForbiddenUpdate represents a StaleStakesForbiddenUpdate event raised by the ContractTriggerXServiceManager contract.
type ContractTriggerXServiceManagerStaleStakesForbiddenUpdate struct {
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStaleStakesForbiddenUpdate is a free log retrieval operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) FilterStaleStakesForbiddenUpdate(opts *bind.FilterOpts) (*ContractTriggerXServiceManagerStaleStakesForbiddenUpdateIterator, error) {

	logs, sub, err := _ContractTriggerXServiceManager.contract.FilterLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXServiceManagerStaleStakesForbiddenUpdateIterator{contract: _ContractTriggerXServiceManager.contract, event: "StaleStakesForbiddenUpdate", logs: logs, sub: sub}, nil
}

// WatchStaleStakesForbiddenUpdate is a free log subscription operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) WatchStaleStakesForbiddenUpdate(opts *bind.WatchOpts, sink chan<- *ContractTriggerXServiceManagerStaleStakesForbiddenUpdate) (event.Subscription, error) {

	logs, sub, err := _ContractTriggerXServiceManager.contract.WatchLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerXServiceManagerStaleStakesForbiddenUpdate)
				if err := _ContractTriggerXServiceManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
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

// ParseStaleStakesForbiddenUpdate is a log parse operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) ParseStaleStakesForbiddenUpdate(log types.Log) (*ContractTriggerXServiceManagerStaleStakesForbiddenUpdate, error) {
	event := new(ContractTriggerXServiceManagerStaleStakesForbiddenUpdate)
	if err := _ContractTriggerXServiceManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTriggerXServiceManagerTaskManagerContractUpdatedIterator is returned from FilterTaskManagerContractUpdated and is used to iterate over the raw logs and unpacked data for TaskManagerContractUpdated events raised by the ContractTriggerXServiceManager contract.
type ContractTriggerXServiceManagerTaskManagerContractUpdatedIterator struct {
	Event *ContractTriggerXServiceManagerTaskManagerContractUpdated // Event containing the contract specifics and raw log

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
func (it *ContractTriggerXServiceManagerTaskManagerContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerXServiceManagerTaskManagerContractUpdated)
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
		it.Event = new(ContractTriggerXServiceManagerTaskManagerContractUpdated)
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
func (it *ContractTriggerXServiceManagerTaskManagerContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerXServiceManagerTaskManagerContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerXServiceManagerTaskManagerContractUpdated represents a TaskManagerContractUpdated event raised by the ContractTriggerXServiceManager contract.
type ContractTriggerXServiceManagerTaskManagerContractUpdated struct {
	OldTaskManager common.Address
	NewTaskManager common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTaskManagerContractUpdated is a free log retrieval operation binding the contract event 0xecc36ab0616c35e413d90730ada58214cbe0188c5f26632fbb32451c9a135784.
//
// Solidity: event TaskManagerContractUpdated(address indexed oldTaskManager, address indexed newTaskManager)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) FilterTaskManagerContractUpdated(opts *bind.FilterOpts, oldTaskManager []common.Address, newTaskManager []common.Address) (*ContractTriggerXServiceManagerTaskManagerContractUpdatedIterator, error) {

	var oldTaskManagerRule []interface{}
	for _, oldTaskManagerItem := range oldTaskManager {
		oldTaskManagerRule = append(oldTaskManagerRule, oldTaskManagerItem)
	}
	var newTaskManagerRule []interface{}
	for _, newTaskManagerItem := range newTaskManager {
		newTaskManagerRule = append(newTaskManagerRule, newTaskManagerItem)
	}

	logs, sub, err := _ContractTriggerXServiceManager.contract.FilterLogs(opts, "TaskManagerContractUpdated", oldTaskManagerRule, newTaskManagerRule)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXServiceManagerTaskManagerContractUpdatedIterator{contract: _ContractTriggerXServiceManager.contract, event: "TaskManagerContractUpdated", logs: logs, sub: sub}, nil
}

// WatchTaskManagerContractUpdated is a free log subscription operation binding the contract event 0xecc36ab0616c35e413d90730ada58214cbe0188c5f26632fbb32451c9a135784.
//
// Solidity: event TaskManagerContractUpdated(address indexed oldTaskManager, address indexed newTaskManager)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) WatchTaskManagerContractUpdated(opts *bind.WatchOpts, sink chan<- *ContractTriggerXServiceManagerTaskManagerContractUpdated, oldTaskManager []common.Address, newTaskManager []common.Address) (event.Subscription, error) {

	var oldTaskManagerRule []interface{}
	for _, oldTaskManagerItem := range oldTaskManager {
		oldTaskManagerRule = append(oldTaskManagerRule, oldTaskManagerItem)
	}
	var newTaskManagerRule []interface{}
	for _, newTaskManagerItem := range newTaskManager {
		newTaskManagerRule = append(newTaskManagerRule, newTaskManagerItem)
	}

	logs, sub, err := _ContractTriggerXServiceManager.contract.WatchLogs(opts, "TaskManagerContractUpdated", oldTaskManagerRule, newTaskManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerXServiceManagerTaskManagerContractUpdated)
				if err := _ContractTriggerXServiceManager.contract.UnpackLog(event, "TaskManagerContractUpdated", log); err != nil {
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

// ParseTaskManagerContractUpdated is a log parse operation binding the contract event 0xecc36ab0616c35e413d90730ada58214cbe0188c5f26632fbb32451c9a135784.
//
// Solidity: event TaskManagerContractUpdated(address indexed oldTaskManager, address indexed newTaskManager)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) ParseTaskManagerContractUpdated(log types.Log) (*ContractTriggerXServiceManagerTaskManagerContractUpdated, error) {
	event := new(ContractTriggerXServiceManagerTaskManagerContractUpdated)
	if err := _ContractTriggerXServiceManager.contract.UnpackLog(event, "TaskManagerContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTriggerXServiceManagerTaskManagerSetIterator is returned from FilterTaskManagerSet and is used to iterate over the raw logs and unpacked data for TaskManagerSet events raised by the ContractTriggerXServiceManager contract.
type ContractTriggerXServiceManagerTaskManagerSetIterator struct {
	Event *ContractTriggerXServiceManagerTaskManagerSet // Event containing the contract specifics and raw log

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
func (it *ContractTriggerXServiceManagerTaskManagerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerXServiceManagerTaskManagerSet)
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
		it.Event = new(ContractTriggerXServiceManagerTaskManagerSet)
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
func (it *ContractTriggerXServiceManagerTaskManagerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerXServiceManagerTaskManagerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerXServiceManagerTaskManagerSet represents a TaskManagerSet event raised by the ContractTriggerXServiceManager contract.
type ContractTriggerXServiceManagerTaskManagerSet struct {
	OldTaskManager common.Address
	NewTaskManager common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTaskManagerSet is a free log retrieval operation binding the contract event 0xdae652729007c1e5ea13dfe9e126a1b72358db8db69c68b96f316dcea59c1c5f.
//
// Solidity: event TaskManagerSet(address indexed oldTaskManager, address indexed newTaskManager)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) FilterTaskManagerSet(opts *bind.FilterOpts, oldTaskManager []common.Address, newTaskManager []common.Address) (*ContractTriggerXServiceManagerTaskManagerSetIterator, error) {

	var oldTaskManagerRule []interface{}
	for _, oldTaskManagerItem := range oldTaskManager {
		oldTaskManagerRule = append(oldTaskManagerRule, oldTaskManagerItem)
	}
	var newTaskManagerRule []interface{}
	for _, newTaskManagerItem := range newTaskManager {
		newTaskManagerRule = append(newTaskManagerRule, newTaskManagerItem)
	}

	logs, sub, err := _ContractTriggerXServiceManager.contract.FilterLogs(opts, "TaskManagerSet", oldTaskManagerRule, newTaskManagerRule)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXServiceManagerTaskManagerSetIterator{contract: _ContractTriggerXServiceManager.contract, event: "TaskManagerSet", logs: logs, sub: sub}, nil
}

// WatchTaskManagerSet is a free log subscription operation binding the contract event 0xdae652729007c1e5ea13dfe9e126a1b72358db8db69c68b96f316dcea59c1c5f.
//
// Solidity: event TaskManagerSet(address indexed oldTaskManager, address indexed newTaskManager)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) WatchTaskManagerSet(opts *bind.WatchOpts, sink chan<- *ContractTriggerXServiceManagerTaskManagerSet, oldTaskManager []common.Address, newTaskManager []common.Address) (event.Subscription, error) {

	var oldTaskManagerRule []interface{}
	for _, oldTaskManagerItem := range oldTaskManager {
		oldTaskManagerRule = append(oldTaskManagerRule, oldTaskManagerItem)
	}
	var newTaskManagerRule []interface{}
	for _, newTaskManagerItem := range newTaskManager {
		newTaskManagerRule = append(newTaskManagerRule, newTaskManagerItem)
	}

	logs, sub, err := _ContractTriggerXServiceManager.contract.WatchLogs(opts, "TaskManagerSet", oldTaskManagerRule, newTaskManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerXServiceManagerTaskManagerSet)
				if err := _ContractTriggerXServiceManager.contract.UnpackLog(event, "TaskManagerSet", log); err != nil {
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

// ParseTaskManagerSet is a log parse operation binding the contract event 0xdae652729007c1e5ea13dfe9e126a1b72358db8db69c68b96f316dcea59c1c5f.
//
// Solidity: event TaskManagerSet(address indexed oldTaskManager, address indexed newTaskManager)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) ParseTaskManagerSet(log types.Log) (*ContractTriggerXServiceManagerTaskManagerSet, error) {
	event := new(ContractTriggerXServiceManagerTaskManagerSet)
	if err := _ContractTriggerXServiceManager.contract.UnpackLog(event, "TaskManagerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTriggerXServiceManagerTaskValidatorSetIterator is returned from FilterTaskValidatorSet and is used to iterate over the raw logs and unpacked data for TaskValidatorSet events raised by the ContractTriggerXServiceManager contract.
type ContractTriggerXServiceManagerTaskValidatorSetIterator struct {
	Event *ContractTriggerXServiceManagerTaskValidatorSet // Event containing the contract specifics and raw log

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
func (it *ContractTriggerXServiceManagerTaskValidatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerXServiceManagerTaskValidatorSet)
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
		it.Event = new(ContractTriggerXServiceManagerTaskValidatorSet)
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
func (it *ContractTriggerXServiceManagerTaskValidatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerXServiceManagerTaskValidatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerXServiceManagerTaskValidatorSet represents a TaskValidatorSet event raised by the ContractTriggerXServiceManager contract.
type ContractTriggerXServiceManagerTaskValidatorSet struct {
	OldTaskValidator common.Address
	NewTaskValidator common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTaskValidatorSet is a free log retrieval operation binding the contract event 0x86cebbccc5bbc9ac544f7e5361e6c5e4acf14f39d85c855032a815125cb97fd8.
//
// Solidity: event TaskValidatorSet(address indexed oldTaskValidator, address indexed newTaskValidator)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) FilterTaskValidatorSet(opts *bind.FilterOpts, oldTaskValidator []common.Address, newTaskValidator []common.Address) (*ContractTriggerXServiceManagerTaskValidatorSetIterator, error) {

	var oldTaskValidatorRule []interface{}
	for _, oldTaskValidatorItem := range oldTaskValidator {
		oldTaskValidatorRule = append(oldTaskValidatorRule, oldTaskValidatorItem)
	}
	var newTaskValidatorRule []interface{}
	for _, newTaskValidatorItem := range newTaskValidator {
		newTaskValidatorRule = append(newTaskValidatorRule, newTaskValidatorItem)
	}

	logs, sub, err := _ContractTriggerXServiceManager.contract.FilterLogs(opts, "TaskValidatorSet", oldTaskValidatorRule, newTaskValidatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXServiceManagerTaskValidatorSetIterator{contract: _ContractTriggerXServiceManager.contract, event: "TaskValidatorSet", logs: logs, sub: sub}, nil
}

// WatchTaskValidatorSet is a free log subscription operation binding the contract event 0x86cebbccc5bbc9ac544f7e5361e6c5e4acf14f39d85c855032a815125cb97fd8.
//
// Solidity: event TaskValidatorSet(address indexed oldTaskValidator, address indexed newTaskValidator)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) WatchTaskValidatorSet(opts *bind.WatchOpts, sink chan<- *ContractTriggerXServiceManagerTaskValidatorSet, oldTaskValidator []common.Address, newTaskValidator []common.Address) (event.Subscription, error) {

	var oldTaskValidatorRule []interface{}
	for _, oldTaskValidatorItem := range oldTaskValidator {
		oldTaskValidatorRule = append(oldTaskValidatorRule, oldTaskValidatorItem)
	}
	var newTaskValidatorRule []interface{}
	for _, newTaskValidatorItem := range newTaskValidator {
		newTaskValidatorRule = append(newTaskValidatorRule, newTaskValidatorItem)
	}

	logs, sub, err := _ContractTriggerXServiceManager.contract.WatchLogs(opts, "TaskValidatorSet", oldTaskValidatorRule, newTaskValidatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerXServiceManagerTaskValidatorSet)
				if err := _ContractTriggerXServiceManager.contract.UnpackLog(event, "TaskValidatorSet", log); err != nil {
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

// ParseTaskValidatorSet is a log parse operation binding the contract event 0x86cebbccc5bbc9ac544f7e5361e6c5e4acf14f39d85c855032a815125cb97fd8.
//
// Solidity: event TaskValidatorSet(address indexed oldTaskValidator, address indexed newTaskValidator)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) ParseTaskValidatorSet(log types.Log) (*ContractTriggerXServiceManagerTaskValidatorSet, error) {
	event := new(ContractTriggerXServiceManagerTaskValidatorSet)
	if err := _ContractTriggerXServiceManager.contract.UnpackLog(event, "TaskValidatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTriggerXServiceManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractTriggerXServiceManager contract.
type ContractTriggerXServiceManagerUnpausedIterator struct {
	Event *ContractTriggerXServiceManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *ContractTriggerXServiceManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerXServiceManagerUnpaused)
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
		it.Event = new(ContractTriggerXServiceManagerUnpaused)
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
func (it *ContractTriggerXServiceManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerXServiceManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerXServiceManagerUnpaused represents a Unpaused event raised by the ContractTriggerXServiceManager contract.
type ContractTriggerXServiceManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractTriggerXServiceManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractTriggerXServiceManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerXServiceManagerUnpausedIterator{contract: _ContractTriggerXServiceManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractTriggerXServiceManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractTriggerXServiceManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerXServiceManagerUnpaused)
				if err := _ContractTriggerXServiceManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerFilterer) ParseUnpaused(log types.Log) (*ContractTriggerXServiceManagerUnpaused, error) {
	event := new(ContractTriggerXServiceManagerUnpaused)
	if err := _ContractTriggerXServiceManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
