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

// IRewardsCoordinatorTypesRewardsSubmission is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorTypesRewardsSubmission struct {
	StrategiesAndMultipliers []IRewardsCoordinatorTypesStrategyAndMultiplier
	Token                    common.Address
	Amount                   *big.Int
	StartTimestamp           uint32
	Duration                 uint32
}

// IRewardsCoordinatorTypesStrategyAndMultiplier is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorTypesStrategyAndMultiplier struct {
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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIAVSDirectory\",\"name\":\"_avsDirectory\",\"type\":\"address\"},{\"internalType\":\"contractIRewardsCoordinator\",\"name\":\"_rewardsCoordinator\",\"type\":\"address\"},{\"internalType\":\"contractIRegistryCoordinator\",\"name\":\"_registryCoordinator\",\"type\":\"address\"},{\"internalType\":\"contractIStakeRegistry\",\"name\":\"_stakeRegistry\",\"type\":\"address\"},{\"internalType\":\"contractIPermissionController\",\"name\":\"_permissionController\",\"type\":\"address\"},{\"internalType\":\"contractIPauserRegistry\",\"name\":\"_pauserRegistry\",\"type\":\"address\"},{\"internalType\":\"contractVetoableSlasher\",\"name\":\"_vetoableSlasher\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BitmapValueTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BytesArrayLengthTooLong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BytesArrayNotOrdered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CurrentlyPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DelayPeriodNotPassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECAddFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECMulFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpModFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InputAddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InputArrayLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InputEmptyQuorumNumbers\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InputNonSignerLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBLSPairingKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBLSSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNewPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidQuorumApkHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidReferenceBlocknumber\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonSignerPubkeysNotSorted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRegistryCoordinator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRegistryCoordinatorOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRewardsInitiator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlySlasher\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyStakeRegistry\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyUnpauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ScalarTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StaleStakesForbidden\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"KeeperAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"KeeperBlacklisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"KeeperRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"KeeperUnblacklisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldQuorumManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newQuorumManager\",\"type\":\"address\"}],\"name\":\"QuorumManagerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prevRewardsInitiator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newRewardsInitiator\",\"type\":\"address\"}],\"name\":\"RewardsInitiatorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"StaleStakesForbiddenUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldTaskManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newTaskManager\",\"type\":\"address\"}],\"name\":\"TaskManagerContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldTaskManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newTaskManager\",\"type\":\"address\"}],\"name\":\"TaskManagerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldTaskValidator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newTaskValidator\",\"type\":\"address\"}],\"name\":\"TaskValidatorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SLASHER_PROPOSAL_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"addPendingAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"avsDirectory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"blacklistKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blsApkRegistry\",\"outputs\":[{\"internalType\":\"contractIBLSApkRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"referenceBlockNumber\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"uint32[]\",\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point[]\",\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point[]\",\"name\":\"quorumApks\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"apkG2\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"sigma\",\"type\":\"tuple\"},{\"internalType\":\"uint32[]\",\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[]\",\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[][]\",\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\"}],\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"checkSignatures\",\"outputs\":[{\"components\":[{\"internalType\":\"uint96[]\",\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\"},{\"internalType\":\"uint96[]\",\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\"}],\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"multiplier\",\"type\":\"uint96\"}],\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[]\",\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"startTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"duration\",\"type\":\"uint32\"}],\"internalType\":\"structIRewardsCoordinatorTypes.RewardsSubmission[]\",\"name\":\"rewardsSubmissions\",\"type\":\"tuple[]\"}],\"name\":\"createAVSRewardsSubmission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegation\",\"outputs\":[{\"internalType\":\"contractIDelegationManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"deregisterOperatorFromAVS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getOperatorRestakedStrategies\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRestakeableStrategies\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITriggerXTaskManager\",\"name\":\"_taskManagerContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardsInitiator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_slasher\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vetoCommittee\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_taskManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_taskValidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_quorumManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isBlackListed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrationFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"}],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauserRegistry\",\"outputs\":[{\"internalType\":\"contractIPauserRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposedSlasher\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quorumManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"name\":\"operatorSignature\",\"type\":\"tuple\"}],\"name\":\"registerOperatorToAVS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registryCoordinator\",\"outputs\":[{\"internalType\":\"contractIRegistryCoordinator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"appointee\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"removeAppointee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pendingAdmin\",\"type\":\"address\"}],\"name\":\"removePendingAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsInitiator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"appointee\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"setAppointee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_quorumManager\",\"type\":\"address\"}],\"name\":\"setQuorumManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRewardsInitiator\",\"type\":\"address\"}],\"name\":\"setRewardsInitiator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"setStaleStakesForbidden\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_taskManager\",\"type\":\"address\"}],\"name\":\"setTaskManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_taskValidator\",\"type\":\"address\"}],\"name\":\"setTaskValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slasher\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slasherProposalTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeRegistry\",\"outputs\":[{\"internalType\":\"contractIStakeRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staleStakesForbidden\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taskManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taskManagerContract\",\"outputs\":[{\"internalType\":\"contractITriggerXTaskManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taskValidator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"apk\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"apkG2\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"sigma\",\"type\":\"tuple\"}],\"name\":\"trySignatureAndApkVerification\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"pairingSuccessful\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"siganatureIsValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"unblacklistKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_metadataURI\",\"type\":\"string\"}],\"name\":\"updateAVSMetadataURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_taskManager\",\"type\":\"address\"}],\"name\":\"updateTaskManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vetoCommittee\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vetoableSlasher\",\"outputs\":[{\"internalType\":\"contractVetoableSlasher\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040526004361015610011575f80fd5b5f3560e01c8063136439dd14610334578063171f1d5b1461032f5780631785f53c1461032a5780631fdb0cfd1461032557806326a965f014610320578063279432eb1461031b578063327d0a601461031657806333cfb7b7146103115780633bc28c8c1461030c578063416c7e5e14610307578063500c8dd31461030257806358ac4a1e146102fd578063595c6a67146102f85780635ac86ab7146102f35780635c975abb146102ee5780635df45946146102e95780635edd78ea146102e457806367940c89146102df57806368304835146102da5780636a8dcf54146102d55780636b3aa72e146102d05780636d14a987146102cb5780636efb4636146102c6578063715018a6146102c1578063886f1195146102bc5780638a29e2de146102b75780638d68349a146102b25780638da5cb5b146102ad5780639926ee7d146102a85780639cd0fb86146102a35780639da16d8e1461029e578063a364f4da14610299578063a41d3f9414610294578063a50a640e1461028f578063a98fb3551461028a578063b134427114610285578063b5f7eb6b14610280578063b7d79c931461027b578063b98d090814610276578063ba55088014610271578063df5cf7231461026c578063e46f181614610267578063e47d606014610262578063e481af9d1461025d578063f2fde38b14610258578063fabc1cbc14610253578063fc299dee1461024e578063fcd1c37514610249578063fce36c7d146102445763fd38ec8c1461023f575f80fd5b611c01565b611ba6565b611b89565b611b61565b611a7f565b6119ee565b6119d3565b611993565b61196b565b611927565b61189b565b611879565b611835565b61180c565b6117e4565b61173c565b611714565b6116ab565b61158d565b611503565b611495565b6113f9565b611380565b61135e565b611218565b6111d4565b611179565b6110e4565b610d8e565b610d4a565b610d21565b610cdd565b610cc0565b610c53565b610c0f565b610bf2565b610bbf565b610b41565b610ad3565b610aaa565b6109aa565b610973565b61093b565b61088c565b610802565b610796565b6106f6565b610615565b61059c565b346103f45760203660031901126103f45760043560405163237dfb4760e11b8152336004820152906020826024817f0000000000000000000000003db5b2115c2c4619255c107c8daade9dd35b366c6001600160a01b03165afa9182156103ef576103be926103aa915f916103c0575b50611c4c565b6103b960015482811614611c62565b613652565b005b6103e2915060203d6020116103e8575b6103da8183610447565b810190611c29565b5f6103a4565b503d6103d0565b611c41565b5f80fd5b634e487b7160e01b5f52604160045260245ffd5b604081019081106001600160401b0382111761042757604052565b6103f8565b606081019081106001600160401b0382111761042757604052565b90601f801991011681019081106001600160401b0382111761042757604052565b6040519061047861010083610447565b565b60405190610478604083610447565b906104786040519283610447565b60409060e31901126103f457604051906104b08261040c565b60e4358252610104356020830152565b91908260409103126103f4576040516104d88161040c565b6020808294803584520135910152565b9080601f830112156103f45760405191610503604084610447565b8290604081019283116103f457905b82821061051f5750505090565b8135815260209182019101610512565b9060806063198301126103f4576040516105488161040c565b6020610563829461055a8160646104e8565b845260a46104e8565b910152565b91906080838203126103f4576020610563604051926105868461040c565b6040849661059483826104e8565b8652016104e8565b346103f4576101203660031901126103f45760043560403660231901126103f4576105f460409182516105ce8161040c565b602435815260443560208201526105e43661052f565b906105ee36610497565b92611cb6565b8251911515825215156020820152f35b6001600160a01b038116036103f457565b346103f4575f60203660031901126103f45760043561063381610604565b61063b613afb565b7f000000000000000000000000598cb226b591155f767da17afe7a2241a68c5c106001600160a01b0316803b156103f45760405163268959e560e01b81523060048201526001600160a01b0390921660248301525f908290818381604481015b03925af180156103ef576106ad575080f35b6103be91505f90610447565b60609060031901126103f4576004356106d181610604565b906024356106de81610604565b906044356001600160e01b0319811681036103f45790565b346103f457610704366106b9565b61070c613afb565b7f000000000000000000000000598cb226b591155f767da17afe7a2241a68c5c106001600160a01b031692833b156103f457604051634a86c03760e11b81523060048201526001600160a01b039182166024820152921660448301526001600160e01b03191660648201525f81608481015b93818381819703925af180156103ef576106ad575080f35b346103f45760203660031901126103f4576004356107b381610604565b6107bb613afb565b6001600160a01b03165f81815260fd60205260408120805460ff191660011790557f79bf5efab685b39bdd2b7b1aba09ffa5e9dfbd86f7aaebab4b52d7ef679d28449080a2005b346103f4575f60203660031901126103f45760043561082081610604565b610828613afb565b7f000000000000000000000000598cb226b591155f767da17afe7a2241a68c5c106001600160a01b0316803b156103f45760405163eb5a4e8760e01b81523060048201526001600160a01b0390921660248301525f9082908183816044810161069b565b346103f45760203660031901126103f4576004356108a981610604565b6108b1613afb565b60fe80546001600160a01b039283166001600160a01b0319821681179092559091167fdae652729007c1e5ea13dfe9e126a1b72358db8db69c68b96f316dcea59c1c5f5f80a3005b60206040818301928281528451809452019201905f5b81811061091c5750505090565b82516001600160a01b031684526020938401939092019160010161090f565b346103f45760203660031901126103f45761096f61096360043561095e81610604565b611f16565b604051918291826108f9565b0390f35b346103f45760203660031901126103f4576103be60043561099381610604565b61099b613afb565b613bf1565b801515036103f457565b346103f45760203660031901126103f4576004356109c7816109a0565b604051638da5cb5b60e01b81526020816004817f000000000000000000000000b438c6fc1652148bb758b939831f9a2cd59ce02b6001600160a01b03165afa9081156103ef575f91610a71575b506001600160a01b03163303610a625760207f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc91151560ff196032541660ff821617603255604051908152a1005b637070f3b160e11b5f5260045ffd5b610a93915060203d602011610a99575b610a8b8183610447565b8101906122a3565b5f610a14565b503d610a81565b5f9103126103f457565b346103f4575f3660031901126103f457610102546040516001600160a01b039091168152602090f35b346103f45760203660031901126103f457600435610af081610604565b610af8613afb565b61010280546001600160a01b039283166001600160a01b0319821681179092559091167fecc36ab0616c35e413d90730ada58214cbe0188c5f26632fbb32451c9a1357845f80a3005b346103f4575f3660031901126103f45760405163237dfb4760e11b81523360048201526020816024817f0000000000000000000000003db5b2115c2c4619255c107c8daade9dd35b366c6001600160a01b03165afa80156103ef57610bac915f916103c05750611c4c565b6103be61361e565b60ff8116036103f457565b346103f45760203660031901126103f4576020600160ff600435610be281610bb4565b161b806001541614604051908152f35b346103f4575f3660031901126103f4576020600154604051908152f35b346103f4575f3660031901126103f4576040517f000000000000000000000000c275cce555cc4bd954ff078621778368f9006f1c6001600160a01b03168152602090f35b346103f45760203660031901126103f457600435610c7081610604565b610c78613afb565b60ff80546001600160a01b039283166001600160a01b0319821681179092559091167f86cebbccc5bbc9ac544f7e5361e6c5e4acf14f39d85c855032a815125cb97fd85f80a3005b346103f4575f3660031901126103f457602060405162093a808152f35b346103f4575f3660031901126103f4576040517f000000000000000000000000d7c48ab86ab9390e09165013600d5e8721cb3e1d6001600160a01b03168152602090f35b346103f4575f3660031901126103f457610101546040516001600160a01b039091168152602090f35b346103f4575f3660031901126103f4576040517f000000000000000000000000055733000064333caddbc92763c58bf0192ffebf6001600160a01b03168152602090f35b346103f4575f3660031901126103f4576040517f000000000000000000000000b438c6fc1652148bb758b939831f9a2cd59ce02b6001600160a01b03168152602090f35b63ffffffff8116036103f457565b6044359061047882610dd2565b359061047882610dd2565b6001600160401b0381116104275760051b60200190565b9080601f830112156103f4578135610e2681610df8565b92610e346040519485610447565b81845260208085019260051b8201019283116103f457602001905b828210610e5c5750505090565b602080918335610e6b81610dd2565b815201910190610e4f565b81601f820112156103f4578035610e8c81610df8565b92610e9a6040519485610447565b81845260208085019260061b840101928184116103f457602001915b838310610ec4575050505090565b6020604091610ed384866104c0565b815201920191610eb6565b9080601f830112156103f4578135610ef581610df8565b92610f036040519485610447565b81845260208085019260051b820101918383116103f45760208201905b838210610f2f57505050505090565b81356001600160401b0381116103f457602091610f5187848094880101610e0f565b815201910190610f20565b919091610180818403126103f457610f72610468565b9281356001600160401b0381116103f45781610f8f918401610e0f565b845260208201356001600160401b0381116103f45781610fb0918401610e76565b602085015260408201356001600160401b0381116103f45781610fd4918401610e76565b6040850152610fe68160608401610568565b6060850152610ff88160e084016104c0565b60808501526101208201356001600160401b0381116103f4578161101d918401610e0f565b60a08501526101408201356001600160401b0381116103f45781611042918401610e0f565b60c08501526101608201356001600160401b0381116103f4576110659201610ede565b60e0830152565b90602080835192838152019201905f5b8181106110895750505090565b82516001600160601b031684526020938401939092019160010161107c565b9291906110df60209160408652826110cb82516040808a0152608089019061106c565b910151868203603f1901606088015261106c565b930152565b346103f45760803660031901126103f4576004356024356001600160401b0381116103f457366023820112156103f45780600401356001600160401b0381116103f45736602482840101116103f45761113b610de0565b90606435936001600160401b0385116103f4576024611161611169963690600401610f5c565b94019061245b565b9061096f604051928392836110a8565b346103f4575f3660031901126103f457611191613afb565b609780546001600160a01b031981169091555f906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b346103f4575f3660031901126103f4576040517f0000000000000000000000003db5b2115c2c4619255c107c8daade9dd35b366c6001600160a01b03168152602090f35b346103f4576101003660031901126103f45760043561123681610604565b6112d260243561124581610604565b60443561125181610604565b60643561125d81610604565b60843561126981610604565b60a4359161127683610604565b60c4359361128385610604565b60e4359561129087610604565b606454986112b660ff8b60081c16151515809b81611350575b8115611330575b50612d3b565b896112c9600160ff196064541617606455565b61131757612d9e565b6112d857005b6112e861ff001960645416606455565b604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602090a1005b61132b61010061ff00196064541617606455565b612d9e565b303b15915081611342575b505f6112b0565b60ff1660011490505f61133b565b600160ff82161091506112a9565b346103f4575f3660031901126103f457602060ff60cd54166040519015158152f35b346103f4575f3660031901126103f4576097546040516001600160a01b039091168152602090f35b6001600160401b03811161042757601f01601f191660200190565b9291926113cf826113a8565b916113dd6040519384610447565b8294818452818301116103f4578281602093845f960137010152565b346103f45760403660031901126103f45760043561141681610604565b602435906001600160401b0382116103f457606060031983360301126103f457604051906114438261042c565b82600401356001600160401b0381116103f4578301366023820112156103f4576103be9361147d60449236906024600482013591016113c3565b84526024810135602085015201356040830152612ee4565b346103f45760203660031901126103f4576004356114b281610604565b6114ba613afb565b61010080546001600160a01b039283166001600160a01b0319821681179092559091167f5d68f69aebb8a1d75f8b2a87dddd6e8db842f9fa3bb63745ea7f49ea5bdf03a95f80a3005b346103f4575f60203660031901126103f45760043561152181610604565b611529613afb565b7f000000000000000000000000598cb226b591155f767da17afe7a2241a68c5c106001600160a01b0316803b156103f457604051634f906cf960e01b81523060048201526001600160a01b0390921660248301525f9082908183816044810161069b565b346103f45760203660031901126103f4576004356115aa81610604565b6115b2613afb565b7f000000000000000000000000055733000064333caddbc92763c58bf0192ffebf6001600160a01b0316803b156103f4576040516351b27a6d60e11b81526001600160a01b0383166004820152905f908290602490829084905af180156103ef57611691575b5061163b6116366001600160a01b0383165b6001600160a01b031690565b613faa565b506001600160a01b0381165f90815260fd60205260409020611662905b805460ff19169055565b6001600160a01b03167fa7a775c2c8141f7985c111748ec31c11e5e44b83528e105c8d1d4e8e6b81cf805f80a2005b8061169f5f6116a593610447565b80610aa0565b5f611618565b346103f45760203660031901126103f4576004356116c881610604565b6116d0613afb565b6001600160a01b03165f81815260fd60205260408120805460ff191690557f8511046d95f95689c89b4807d91820adbc266db891b2c5c25a42f812781dde579080a2005b346103f4575f3660031901126103f45760fe546040516001600160a01b039091168152602090f35b346103f4575f60203660031901126103f4576004356001600160401b0381116103f457366023820112156103f45761177e9036906024816004013591016113c3565b611786613afb565b7f000000000000000000000000055733000064333caddbc92763c58bf0192ffebf6001600160a01b0316803b156103f45760405163a98fb35560e01b815260206004820152915f91839182908490829061069b906024830190612ff9565b346103f4575f3660031901126103f45760ca546040516001600160a01b039091168152602090f35b346103f4575f3660031901126103f457610100546040516001600160a01b039091168152602090f35b346103f4575f3660031901126103f4576040517f000000000000000000000000ea80328759f6269ca985e631944ba0d23c933bd46001600160a01b03168152602090f35b346103f4575f3660031901126103f457602060ff603254166040519015158152f35b346103f4576118a9366106b9565b6118b1613afb565b7f000000000000000000000000598cb226b591155f767da17afe7a2241a68c5c106001600160a01b031692833b156103f457604051630664120160e01b81523060048201526001600160a01b039182166024820152921660448301526001600160e01b03191660648201525f816084810161077e565b346103f4575f3660031901126103f4576040517f000000000000000000000000a44151489861fe9e3055d95adc98fbd462b948e76001600160a01b03168152602090f35b346103f4575f3660031901126103f45760cb546040516001600160a01b039091168152602090f35b346103f45760203660031901126103f4576004356119b081610604565b60018060a01b03165f5260fd602052602060ff60405f2054166040519015158152f35b346103f4575f3660031901126103f45761096f61096361301d565b346103f45760203660031901126103f457600435611a0b81610604565b611a13613afb565b6001600160a01b03811615611a2b576103be90613ded565b60405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608490fd5b346103f45760203660031901126103f45760043560405163755b36bd60e11b81526020816004817f0000000000000000000000003db5b2115c2c4619255c107c8daade9dd35b366c6001600160a01b03165afa9081156103ef575f91611b42575b506001600160a01b03163303611b3357611b01600154198219811614611c62565b806001556040519081527f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c60203392a2005b63794821ff60e01b5f5260045ffd5b611b5b915060203d602011610a9957610a8b8183610447565b5f611ae0565b346103f4575f3660031901126103f45760c9546040516001600160a01b039091168152602090f35b346103f4575f3660031901126103f457602060cc54604051908152f35b346103f45760203660031901126103f4576004356001600160401b0381116103f457366023820112156103f45780600401356001600160401b0381116103f4573660248260051b840101116103f45760246103be920161324a565b346103f4575f3660031901126103f45760ff546040516001600160a01b039091168152602090f35b908160209103126103f45751611c3e816109a0565b90565b6040513d5f823e3d90fd5b15611c5357565b631d77d47760e21b5f5260045ffd5b15611c6957565b63c61dca5d60e01b5f5260045ffd5b634e487b7160e01b5f52603260045260245ffd5b906002811015611c9d5760051b0190565b611c78565b634e487b7160e01b5f52601260045260245ffd5b611d92611d6f611d9895611d69611d6285875160208901518a515160208c51015160208d016020815151915101519189519360208b0151956040519760208901998a5260208a015260408901526060880152608087015260a086015260c085015260e0840152610100830152611d3981610120840103601f198101835282610447565b5190207f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001900690565b80966136c8565b9061370e565b92611d69611d84611d7e613770565b94613867565b91611d8d613983565b6136c8565b916139cd565b9091565b908160209103126103f4575190565b908160209103126103f457516001600160c01b03811681036103f45790565b908160209103126103f45751611c3e81610bb4565b60405190611dee602083610447565b5f808352366020840137565b90611e0482610df8565b611e116040519182610447565b8281528092611e22601f1991610df8565b0190602036910137565b908151811015611c9d570160200190565b634e487b7160e01b5f52601160045260245ffd5b9060018201809211611e5f57565b611e3d565b9060028201809211611e5f57565b9060038201809211611e5f57565b9060048201809211611e5f57565b9060058201809211611e5f57565b91908201809211611e5f57565b6001600160601b038116036103f457565b908160409103126103f457602060405191611ed48361040c565b8051611edf81610604565b83520151611eec81611ea9565b602082015290565b8051821015611c9d5760209160051b010190565b5f198114611e5f5760010190565b6040516309aa152760e11b81526001600160a01b0391821660048201527f000000000000000000000000b438c6fc1652148bb758b939831f9a2cd59ce02b90911690602081602481855afa9081156103ef57611f96916020915f91612286575b506040518093819263871ef04960e01b8352600483019190602083019252565b0381855afa9081156103ef575f91612257575b506001600160c01b03169081159081156121f4575b506121eb57611fcc90613b53565b5f91907f000000000000000000000000d7c48ab86ab9390e09165013600d5e8721cb3e1d6001600160a01b031690835b81518510156120a857612051602061202e61202861201a8987611e2c565b516001600160f81b03191690565b60f81c90565b604051633ca5a5f560e01b815260ff909116600482015291829081906024820190565b0381875afa80156103ef57600192612070925f92612078575b50611e9c565b940193611ffc565b61209a91925060203d81116120a1575b6120928183610447565b810190611d9c565b905f61206a565b503d612088565b6120b3919450611dfa565b925f905f5b81518110156121e5576120d161202861201a8385611e2c565b604051633ca5a5f560e01b815260ff8216600482015290602082602481895afa9182156103ef575f926121c5575b50905f915b818310612116575050506001016120b8565b604080516356e4026d60e11b815260ff83166004820152602481018590529396929391929190816044818b5afa9182156103ef576121898b61217a8361217461162a60019861218e985f91612197575b50516001600160a01b031690565b92611ef4565b6001600160a01b039091169052565b611f08565b95019190612104565b6121b8915060403d81116121be575b6121b08183610447565b810190611eba565b5f612166565b503d6121a6565b6121de91925060203d81116120a1576120928183610447565b905f6120ff565b50505050565b50611c3e611ddf565b604051639aa1653d60e01b81529150602090829060049082905afa80156103ef5760ff915f91612228575b5016155f611fbe565b61224a915060203d602011612250575b6122428183610447565b810190611dca565b5f61221f565b503d612238565b612279915060203d60201161227f575b6122718183610447565b810190611dab565b5f611fa9565b503d612267565b61229d9150823d84116120a1576120928183610447565b5f611f76565b908160209103126103f45751611c3e81610604565b604051906122c58261040c565b60606020838281520152565b156122d857565b62f8202d60e51b5f5260045ffd5b156122ed57565b6343714afd60e01b5f5260045ffd5b1561230357565b635f832f4160e01b5f5260045ffd5b1561231957565b634b874f4560e01b5f5260045ffd5b5f19810191908211611e5f57565b1561233d57565b633fdc650560e21b5f5260045ffd5b908160209103126103f45751611c3e81610dd2565b90821015611c9d570190565b1561237457565b63affc5edb60e01b5f5260045ffd5b908160209103126103f4575167ffffffffffffffff19811681036103f45790565b156123ab57565b63e1310aed60e01b5f5260045ffd5b908160209103126103f45751611c3e81611ea9565b906001600160601b03809116911603906001600160601b038211611e5f57565b156123f657565b6367988d3360e01b5f5260045ffd5b1561240c57565b63ab1b236b60e01b5f5260045ffd5b60049163ffffffff60e01b9060e01b1681520160208251919201905f5b8181106124455750505090565b8251845260209384019390920191600101612438565b9493929091936124696122b8565b506124758515156122d1565b604084015151851480612d2d575b80612d1f575b80612d11575b612498906122e6565b6124aa602085015151855151146122fc565b6124c163ffffffff431663ffffffff841610612312565b6124c961047a565b5f81525f6020820152926124db6122b8565b6124e487611dfa565b60208201526124f287611dfa565b81526124fc6122b8565b9261250b602088015151611dfa565b845261251b602088015151611dfa565b602085810191909152604051639aa1653d60e01b815290816004817f000000000000000000000000b438c6fc1652148bb758b939831f9a2cd59ce02b6001600160a01b03165afa80156103ef57612584915f91612cf2575b5061257f368b876113c3565b613c4f565b985f965b60208901518051891015612700576020886125f56125eb8c6125e38f96868e6125c86125b5868095611ef4565b5180515f526020015160205260405f2090565b6125d58484840151611ef4565b52826126cd575b0151611ef4565b519551611ef4565b5163ffffffff1690565b6040516304ec635160e01b8152600481019490945263ffffffff9182166024850152166044830152816064816001600160a01b037f000000000000000000000000b438c6fc1652148bb758b939831f9a2cd59ce02b165afa9182156103ef57611d698a6126a28f61269b8f8460208f926126929361268a8460019e6126a89e5f916126b0575b508f8060c01b03169251611ef4565b520151611ef4565b51938d51611ef4565b5116613c7a565b90613cab565b970196612588565b6126c79150863d811161227f576122718183610447565b5f61267b565b6126fb6126dd8484840151611ef4565b516126f4848401516126ee87612328565b90611ef4565b5110612336565b6125dc565b50909597949650612715919893929950613d68565b9161272260325460ff1690565b908115612cea576040516318891fd760e31b81526020816004817f000000000000000000000000a44151489861fe9e3055d95adc98fbd462b948e76001600160a01b03165afa9081156103ef575f91612cbb575b5091905b5f925b8184106127e1575050505050926127ba6127b56127ae6127db95856127cd9860806060602099015192015192611cb6565b91906123ef565b612405565b015160405192839160208301958661241b565b03601f198101835282610447565b51902090565b92989596909399919794878b888c888d612bb5575b6125eb8260a06128446120286128368461284c976128306128226125b58f9c604060209f9e0151611ef4565b67ffffffffffffffff191690565b9b612361565b356001600160f81b03191690565b970151611ef4565b604051631a2f32ab60e21b815260ff95909516600486015263ffffffff9182166024860152166044840152826064816001600160a01b037f000000000000000000000000c275cce555cc4bd954ff078621778368f9006f1c165afa9081156103ef576129106125eb8f958f906129088f978f96848f61290260c0966128fb848f60209f906125dc612836996040936120289c5f91612b87575b5067ffffffffffffffff199182169116146123a4565b519061370e565b9c612361565b960151611ef4565b604051636414a62b60e11b815260ff94909416600485015263ffffffff9182166024850152166044830152816064816001600160a01b037f000000000000000000000000d7c48ab86ab9390e09165013600d5e8721cb3e1d165afa9081156103ef5761299d918c8f925f92612b63575b50602061298f92930151611ef4565b906001600160601b03169052565b6129ca8c61298f8c6129c36129b6826020860151611ef4565b516001600160601b031690565b9251611ef4565b5f985f5b60208a015151811015612b4a578b8d612a0c896129ff612028612836868f896129f79151611ef4565b519487612361565b60ff161c60019081161490565b612a1b575b50506001016129ce565b8a8a612aa3859f948f9686612a5d8f9360e0612a546125eb956020612a4c612028612836839f612a639c8991612361565b9a0151611ef4565b519b0151611ef4565b51611ef4565b60405163795f4a5760e11b815260ff909316600484015263ffffffff93841660248401526044830196909652919094166064850152839081906084820190565b03817f000000000000000000000000d7c48ab86ab9390e09165013600d5e8721cb3e1d6001600160a01b03165afa9081156103ef578f612b09908f936001959486955f92612b14575b5061217461298f92935193612b046129b68487611ef4565b6123cf565b019a90508b8d612a11565b61298f9250612b3c6121749160203d8111612b43575b612b348183610447565b8101906123ba565b9250612aec565b503d612b2a565b5093919796996001919699509a94929a0192919061277d565b61298f9250612b80602091823d8111612b4357612b348183610447565b9250612980565b6020612ba892503d8111612bae575b612ba08183610447565b810190612383565b5f6128e5565b503d612b96565b612bf29450612bcf92506120289161283691602095612361565b60405163124d062160e11b815260ff909116600482015291829081906024820190565b03817f000000000000000000000000b438c6fc1652148bb758b939831f9a2cd59ce02b6001600160a01b03165afa80156103ef5760208961284c8f938f60a08f976120286128368f8f906128306128226125b58f60408b96918f88936125eb9f612c7690612c7c936128449f5f92612c92575b5063ffffffff809116931690611e9c565b1161236d565b50505050505097505050505050929350506127f6565b602063ffffffff9293508291612cb3913d81116120a1576120928183610447565b929150612c65565b612cdd915060203d602011612ce3575b612cd58183610447565b81019061234c565b5f612776565b503d612ccb565b5f919061277a565b612d0b915060203d602011612250576122428183610447565b5f612573565b5060e084015151851461248f565b5060c0840151518514612489565b5060a0840151518514612483565b15612d4257565b60405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608490fd5b929395969196612dad82613ded565b60ff60645460081c1615612e8b57612e13612e6c96612e4f95612dd96104789b61099b612e3097613ded565b60018060a01b03166001600160601b0360a01b6101025416176101025560018060a01b03166001600160601b0360a01b60fe54161760fe55565b60018060a01b03166001600160601b0360a01b60ff54161760ff55565b60018060a01b03166001600160601b0360a01b61010054161761010055565b60018060a01b03166001600160601b0360a01b60ca54161760ca55565b60018060a01b03166001600160601b0360a01b61010154161761010155565b60405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b6064820152608490fd5b612eec613afb565b7f000000000000000000000000055733000064333caddbc92763c58bf0192ffebf6001600160a01b031690813b156103f4575f6040518093639926ee7d60e01b825281838160018060a01b03871698896004830152604060248301526040612f6082516060604486015260a4850190612ff9565b91602081015160648501520151608483015203925af19081156103ef57612fbf9261165892612fe5575b50612fa5612fa06001600160a01b03831661162a565b613ef5565b506001600160a01b03165f90815260fd6020526040902090565b7f1584773458d98c71b34a270ee1100b3a42889bf91e3b7a858563b684c24d838e5f80a2565b8061169f5f612ff393610447565b5f612f8a565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b604051639aa1653d60e01b81527f000000000000000000000000b438c6fc1652148bb758b939831f9a2cd59ce02b6001600160a01b031690602081600481855afa80156103ef5760ff915f9161322b575b50168015613221577f000000000000000000000000d7c48ab86ab9390e09165013600d5e8721cb3e1d6001600160a01b0316905f9081905b8083106131dd57506130b89150611dfa565b925f905f5b604051639aa1653d60e01b8152602081600481895afa80156103ef5760ff915f916131bf575b50168110156131b857604051633ca5a5f560e01b815260ff821660048201819052602082602481895afa9182156103ef575f92613198575b50905f915b818310613132575050506001016130bd565b604080516356e4026d60e11b815260ff83166004820152602481018590529396929391929190816044818b5afa9182156103ef576121898b61217a8361217461162a60019861318f985f916121975750516001600160a01b031690565b95019190613120565b6131b191925060203d81116120a1576120928183610447565b905f61311b565b5092505050565b6131d7915060203d8111612250576122428183610447565b5f6130e3565b604051633ca5a5f560e01b815260ff84166004820152909190602081602481885afa80156103ef57600192613218925f926120785750611e9c565b920191906130a6565b5050611c3e611ddf565b613244915060203d602011612250576122428183610447565b5f61306e565b60c9546001600160a01b0316330361346c577f000000000000000000000000acc1fb458a1317e886db376fc8141540537e68fe6001600160a01b0316915f5b8181106132da5750823b156103f4576132bb925f928360405180968195829463fce36c7d60e01b84526004840161350f565b03925af180156103ef576132cc5750565b8061169f5f61047893610447565b5f60206133356132f861162a836132f287898b61347b565b0161349d565b604061330586888a61347b565b6040516323b872dd60e01b8152336004820152306024820152910135604482015293849283919082906064820190565b03925af180156103ef57613450575b5061335861162a60206132f284868861347b565b604051636eb1769f60e11b81523060048201526001600160a01b03861660248201529190602090839060449082905afa80156103ef576133fb6020915f948591613433575b506133c86133b361162a856132f2888b8d61347b565b9160406133c1878a8c61347b565b0135611e9c565b60405163095ea7b360e01b81526001600160a01b038a166004820152602481019190915294859283919082906044820190565b03925af19182156103ef57600192613415575b5001613289565b61342c9060203d81116103e8576103da8183610447565b505f61340e565b61344a9150833d81116120a1576120928183610447565b5f61339d565b6134679060203d81116103e8576103da8183610447565b613344565b638e79fdb560e01b5f5260045ffd5b9190811015611c9d5760051b81013590609e19813603018212156103f4570190565b35611c3e81610604565b916020908281520191905f5b8181106134c05750505090565b90919260408060019286356134d481610604565b848060a01b031681526001600160601b0360208801356134f381611ea9565b1660208201520194019291016134b3565b359061047882610604565b9180602084016020855252604083019060408160051b85010193835f91609e1982360301905b848410613546575050505050505090565b90919293949596603f198282030187528735838112156103f4578401908135601e19833603018112156103f457820191602083359301906001600160401b0384116103f4578360061b360382136103f45761360f836080613604816135ba6020989760019a60a08b9a5260a08701916134a7565b956135d86135c9898301613504565b6001600160a01b0316868a0152565b604081013560408601526135fe6135f160608301610ded565b63ffffffff166060870152565b01610ded565b63ffffffff16910152565b99019701959401929190613535565b5f196001556040515f1981527fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d60203392a2565b806001556040519081527fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d60203392a2565b604051906136918261040c565b5f6020838281520152565b604051906101806136ad8184610447565b368337565b604051906136c1602083610447565b6020368337565b919060409060606136d7613684565b94859260208551926136e98585610447565b8436853780518452015160208301528482015260076107cf195a01fa1561370c57565bfe5b60209291608060409261371f613684565b958693818651936137308686610447565b85368637805185520151828401528051868401520151606082015260066107cf195a01fa801561370c571561376157565b63d4b68fd760e01b5f5260045ffd5b60405161377c8161040c565b604090815161378b8382610447565b82368237815260208251916137a08484610447565b83368437015280516137b28282610447565b7f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c281527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed60208201528151906138088383610447565b7f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d602083015261385d83519384610447565b8252602082015290565b5f5160206140625f395f51905f529061387e613684565b505f919006602060c0835b61397e575f935f5160206140625f395f51905f52600381868181800909086040516138b48582610447565b843682378481856040516138c88282610447565b813682378381528360208201528360408201528560608201527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808201525f5160206140625f395f51905f5260a082015260056107cf195a01fa801561370c576139329061404b565b519161397e575f5160206140625f395f51905f528280091461396957505f5160206140625f395f51905f5260015f94089293613889565b9293505061397561047a565b92835282015290565b611ca2565b61398b613684565b506040516139988161040c565b600181526002602082015290565b90600682029180830460061490151715611e5f57565b90600c811015611c9d5760051b0190565b939290916139db6040610489565b94855260208501526139ed6040610489565b91825260208201526139fd61369c565b925f5b60028110613a2a57505050602061018092613a196136b2565b93849160086201d4c0fa9151151590565b80613a366001926139a6565b613a408285611c8c565b5151613a4c82896139bc565b526020613a598386611c8c565b510151613a6e613a6883611e51565b896139bc565b52613a798286611c8c565b515151613a88613a6883611e64565b52613a9e613a968387611c8c565b515160200190565b51613aab613a6883611e72565b526020613ab88387611c8c565b51015151613ac8613a6883611e80565b52613af4613aee613ae76020613ade868a611c8c565b51015160200190565b5192611e8e565b886139bc565b5201613a00565b6097546001600160a01b03163303613b0f57565b606460405162461bcd60e51b815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b61ffff613b5f82613c7a565b16613b69816113a8565b90613b776040519283610447565b808252613b86601f19916113a8565b013660208301375f5f5b8251821080613be6575b15613bdf576001811b8416613bb8575b613bb390611f08565b613b90565b906001613bb39160ff60f81b8460f81b165f1a613bd58287611e2c565b5301919050613baa565b5050905090565b506101008110613b9a565b60c954604080516001600160a01b038084168252841660208201529192917fe11cddf1816a43318ca175bbc52cd0185436e9cbead7c83acc54a73e461717e39190a16001600160a01b03166001600160a01b0319919091161760c955565b906001613c5d60ff93613e4b565b928392161b1115613c6b5790565b63ca95733360e01b5f5260045ffd5b805f915b613c86575090565b5f198101818111611e5f5761ffff9116911661ffff8114611e5f576001019080613c7e565b90613cb4613684565b5061ffff811690610200821015613d595760018214613d5457613cd561047a565b5f81525f602082015292906001905f925b61ffff8316851015613cfa57505050505090565b600161ffff831660ff86161c811614613d34575b6001613d2a613d1f8360ff9461370e565b9460011b61fffe1690565b9401169291613ce6565b946001613d2a613d1f613d498960ff9561370e565b989350505050613d0e565b505090565b637fc4ea7d60e11b5f5260045ffd5b613d70613684565b50805190811580613de1575b15613d9d575050604051613d91604082610447565b5f81525f602082015290565b60205f5160206140625f395f51905f52910151065f5160206140625f395f51905f52035f5160206140625f395f51905f528111611e5f576040519161385d8361040c565b50602081015115613d7c565b609780546001600160a01b039283166001600160a01b0319821681179092559091167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3565b15613e3c57565b631019106960e31b5f5260045ffd5b90610100825111613eb457815115613eaf57602082015160019060f81c81901b5b8351821015613eaa57600190613e95613e8b61202861201a8689611e2c565b60ff600191161b90565b90613ea1818311613e35565b17910190613e6c565b925050565b5f9150565b637da54e4760e11b5f5260045ffd5b8054821015611c9d575f5260205f2001905f90565b91613ef19183549060031b91821b915f19901b19161790565b9055565b805f5260fc60205260405f2054155f14613f6a5760fb5468010000000000000000811015610427576001810160fb5560fb54811015611c9d577f3e7257b7272bb46d49cd6019b04ddee20da7c0cb13f7c1ec3391291b2ccebabc0181905560fb545f91825260fc602052604090912055600190565b505f90565b80548015613f96575f190190613f858282613ec3565b8154905f199060031b1b1916905555565b634e487b7160e01b5f52603160045260245ffd5b5f81815260fc6020526040902054908115614045575f19820190828211611e5f5760fb545f19810193908411611e5f578383614004945f960361400a575b505050613ff560fb613f6f565b60fc905f5260205260405f2090565b55600190565b613ff56140369161402c61402261403c9560fb613ec3565b90549060031b1c90565b92839160fb613ec3565b90613ed8565b555f8080613fe8565b50505f90565b1561405257565b63d51edae360e01b5f5260045ffdfe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a2646970667358221220d4be2623ece19f474df47c8a238da6879a30314244b0858ef261ebe396fdbe8a64736f6c634300081b0033",
}

// ContractTriggerXServiceManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractTriggerXServiceManagerMetaData.ABI instead.
var ContractTriggerXServiceManagerABI = ContractTriggerXServiceManagerMetaData.ABI

// ContractTriggerXServiceManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractTriggerXServiceManagerMetaData.Bin instead.
var ContractTriggerXServiceManagerBin = ContractTriggerXServiceManagerMetaData.Bin

// DeployContractTriggerXServiceManager deploys a new Ethereum contract, binding an instance of ContractTriggerXServiceManager to it.
func DeployContractTriggerXServiceManager(auth *bind.TransactOpts, backend bind.ContractBackend, _avsDirectory common.Address, _rewardsCoordinator common.Address, _registryCoordinator common.Address, _stakeRegistry common.Address, _permissionController common.Address, _pauserRegistry common.Address, _vetoableSlasher common.Address) (common.Address, *types.Transaction, *ContractTriggerXServiceManager, error) {
	parsed, err := ContractTriggerXServiceManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractTriggerXServiceManagerBin), backend, _avsDirectory, _rewardsCoordinator, _registryCoordinator, _stakeRegistry, _permissionController, _pauserRegistry, _vetoableSlasher)
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
	SLASHERPROPOSALDELAY(opts *bind.CallOpts) (*big.Int, error)

	AvsDirectory(opts *bind.CallOpts) (common.Address, error)

	BlsApkRegistry(opts *bind.CallOpts) (common.Address, error)

	CheckSignatures(opts *bind.CallOpts, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error)

	Delegation(opts *bind.CallOpts) (common.Address, error)

	GetOperatorRestakedStrategies(opts *bind.CallOpts, operator common.Address) ([]common.Address, error)

	GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error)

	IsBlackListed(opts *bind.CallOpts, arg0 common.Address) (bool, error)

	MigrationFinalized(opts *bind.CallOpts) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts, index uint8) (bool, error)

	Paused0(opts *bind.CallOpts) (*big.Int, error)

	PauserRegistry(opts *bind.CallOpts) (common.Address, error)

	ProposedSlasher(opts *bind.CallOpts) (common.Address, error)

	QuorumManager(opts *bind.CallOpts) (common.Address, error)

	RegistryCoordinator(opts *bind.CallOpts) (common.Address, error)

	RewardsInitiator(opts *bind.CallOpts) (common.Address, error)

	Slasher(opts *bind.CallOpts) (common.Address, error)

	SlasherProposalTimestamp(opts *bind.CallOpts) (*big.Int, error)

	StakeRegistry(opts *bind.CallOpts) (common.Address, error)

	StaleStakesForbidden(opts *bind.CallOpts) (bool, error)

	TaskManager(opts *bind.CallOpts) (common.Address, error)

	TaskManagerContract(opts *bind.CallOpts) (common.Address, error)

	TaskValidator(opts *bind.CallOpts) (common.Address, error)

	TrySignatureAndApkVerification(opts *bind.CallOpts, msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
		PairingSuccessful bool
		SiganatureIsValid bool
	}, error)

	VetoCommittee(opts *bind.CallOpts) (common.Address, error)

	VetoableSlasher(opts *bind.CallOpts) (common.Address, error)
}

// ContractTriggerXServiceManagerTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractTriggerXServiceManagerTransacts interface {
	AddPendingAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error)

	BlacklistKeeper(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error)

	CreateAVSRewardsSubmission(opts *bind.TransactOpts, rewardsSubmissions []IRewardsCoordinatorTypesRewardsSubmission) (*types.Transaction, error)

	DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts, _taskManagerContract common.Address, initialOwner common.Address, rewardsInitiator common.Address, _slasher common.Address, _vetoCommittee common.Address, _taskManager common.Address, _taskValidator common.Address, _quorumManager common.Address) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error)

	PauseAll(opts *bind.TransactOpts) (*types.Transaction, error)

	RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error)

	RemoveAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error)

	RemoveAppointee(opts *bind.TransactOpts, appointee common.Address, target common.Address, selector [4]byte) (*types.Transaction, error)

	RemovePendingAdmin(opts *bind.TransactOpts, pendingAdmin common.Address) (*types.Transaction, error)

	RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetAppointee(opts *bind.TransactOpts, appointee common.Address, target common.Address, selector [4]byte) (*types.Transaction, error)

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

// SLASHERPROPOSALDELAY is a free data retrieval call binding the contract method 0x67940c89.
//
// Solidity: function SLASHER_PROPOSAL_DELAY() view returns(uint256)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCaller) SLASHERPROPOSALDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractTriggerXServiceManager.contract.Call(opts, &out, "SLASHER_PROPOSAL_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SLASHERPROPOSALDELAY is a free data retrieval call binding the contract method 0x67940c89.
//
// Solidity: function SLASHER_PROPOSAL_DELAY() view returns(uint256)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) SLASHERPROPOSALDELAY() (*big.Int, error) {
	return _ContractTriggerXServiceManager.Contract.SLASHERPROPOSALDELAY(&_ContractTriggerXServiceManager.CallOpts)
}

// SLASHERPROPOSALDELAY is a free data retrieval call binding the contract method 0x67940c89.
//
// Solidity: function SLASHER_PROPOSAL_DELAY() view returns(uint256)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCallerSession) SLASHERPROPOSALDELAY() (*big.Int, error) {
	return _ContractTriggerXServiceManager.Contract.SLASHERPROPOSALDELAY(&_ContractTriggerXServiceManager.CallOpts)
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

// MigrationFinalized is a free data retrieval call binding the contract method 0x8d68349a.
//
// Solidity: function migrationFinalized() view returns(bool)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCaller) MigrationFinalized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractTriggerXServiceManager.contract.Call(opts, &out, "migrationFinalized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MigrationFinalized is a free data retrieval call binding the contract method 0x8d68349a.
//
// Solidity: function migrationFinalized() view returns(bool)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) MigrationFinalized() (bool, error) {
	return _ContractTriggerXServiceManager.Contract.MigrationFinalized(&_ContractTriggerXServiceManager.CallOpts)
}

// MigrationFinalized is a free data retrieval call binding the contract method 0x8d68349a.
//
// Solidity: function migrationFinalized() view returns(bool)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCallerSession) MigrationFinalized() (bool, error) {
	return _ContractTriggerXServiceManager.Contract.MigrationFinalized(&_ContractTriggerXServiceManager.CallOpts)
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

// ProposedSlasher is a free data retrieval call binding the contract method 0xe46f1816.
//
// Solidity: function proposedSlasher() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCaller) ProposedSlasher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTriggerXServiceManager.contract.Call(opts, &out, "proposedSlasher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProposedSlasher is a free data retrieval call binding the contract method 0xe46f1816.
//
// Solidity: function proposedSlasher() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) ProposedSlasher() (common.Address, error) {
	return _ContractTriggerXServiceManager.Contract.ProposedSlasher(&_ContractTriggerXServiceManager.CallOpts)
}

// ProposedSlasher is a free data retrieval call binding the contract method 0xe46f1816.
//
// Solidity: function proposedSlasher() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCallerSession) ProposedSlasher() (common.Address, error) {
	return _ContractTriggerXServiceManager.Contract.ProposedSlasher(&_ContractTriggerXServiceManager.CallOpts)
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

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCaller) Slasher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTriggerXServiceManager.contract.Call(opts, &out, "slasher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) Slasher() (common.Address, error) {
	return _ContractTriggerXServiceManager.Contract.Slasher(&_ContractTriggerXServiceManager.CallOpts)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCallerSession) Slasher() (common.Address, error) {
	return _ContractTriggerXServiceManager.Contract.Slasher(&_ContractTriggerXServiceManager.CallOpts)
}

// SlasherProposalTimestamp is a free data retrieval call binding the contract method 0xfcd1c375.
//
// Solidity: function slasherProposalTimestamp() view returns(uint256)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCaller) SlasherProposalTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractTriggerXServiceManager.contract.Call(opts, &out, "slasherProposalTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlasherProposalTimestamp is a free data retrieval call binding the contract method 0xfcd1c375.
//
// Solidity: function slasherProposalTimestamp() view returns(uint256)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) SlasherProposalTimestamp() (*big.Int, error) {
	return _ContractTriggerXServiceManager.Contract.SlasherProposalTimestamp(&_ContractTriggerXServiceManager.CallOpts)
}

// SlasherProposalTimestamp is a free data retrieval call binding the contract method 0xfcd1c375.
//
// Solidity: function slasherProposalTimestamp() view returns(uint256)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCallerSession) SlasherProposalTimestamp() (*big.Int, error) {
	return _ContractTriggerXServiceManager.Contract.SlasherProposalTimestamp(&_ContractTriggerXServiceManager.CallOpts)
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

// VetoCommittee is a free data retrieval call binding the contract method 0x6a8dcf54.
//
// Solidity: function vetoCommittee() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCaller) VetoCommittee(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTriggerXServiceManager.contract.Call(opts, &out, "vetoCommittee")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VetoCommittee is a free data retrieval call binding the contract method 0x6a8dcf54.
//
// Solidity: function vetoCommittee() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) VetoCommittee() (common.Address, error) {
	return _ContractTriggerXServiceManager.Contract.VetoCommittee(&_ContractTriggerXServiceManager.CallOpts)
}

// VetoCommittee is a free data retrieval call binding the contract method 0x6a8dcf54.
//
// Solidity: function vetoCommittee() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCallerSession) VetoCommittee() (common.Address, error) {
	return _ContractTriggerXServiceManager.Contract.VetoCommittee(&_ContractTriggerXServiceManager.CallOpts)
}

// VetoableSlasher is a free data retrieval call binding the contract method 0xb7d79c93.
//
// Solidity: function vetoableSlasher() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCaller) VetoableSlasher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTriggerXServiceManager.contract.Call(opts, &out, "vetoableSlasher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VetoableSlasher is a free data retrieval call binding the contract method 0xb7d79c93.
//
// Solidity: function vetoableSlasher() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) VetoableSlasher() (common.Address, error) {
	return _ContractTriggerXServiceManager.Contract.VetoableSlasher(&_ContractTriggerXServiceManager.CallOpts)
}

// VetoableSlasher is a free data retrieval call binding the contract method 0xb7d79c93.
//
// Solidity: function vetoableSlasher() view returns(address)
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerCallerSession) VetoableSlasher() (common.Address, error) {
	return _ContractTriggerXServiceManager.Contract.VetoableSlasher(&_ContractTriggerXServiceManager.CallOpts)
}

// AddPendingAdmin is a paid mutator transaction binding the contract method 0x279432eb.
//
// Solidity: function addPendingAdmin(address admin) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactor) AddPendingAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.contract.Transact(opts, "addPendingAdmin", admin)
}

// AddPendingAdmin is a paid mutator transaction binding the contract method 0x279432eb.
//
// Solidity: function addPendingAdmin(address admin) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) AddPendingAdmin(admin common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.AddPendingAdmin(&_ContractTriggerXServiceManager.TransactOpts, admin)
}

// AddPendingAdmin is a paid mutator transaction binding the contract method 0x279432eb.
//
// Solidity: function addPendingAdmin(address admin) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactorSession) AddPendingAdmin(admin common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.AddPendingAdmin(&_ContractTriggerXServiceManager.TransactOpts, admin)
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
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactor) CreateAVSRewardsSubmission(opts *bind.TransactOpts, rewardsSubmissions []IRewardsCoordinatorTypesRewardsSubmission) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.contract.Transact(opts, "createAVSRewardsSubmission", rewardsSubmissions)
}

// CreateAVSRewardsSubmission is a paid mutator transaction binding the contract method 0xfce36c7d.
//
// Solidity: function createAVSRewardsSubmission(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) CreateAVSRewardsSubmission(rewardsSubmissions []IRewardsCoordinatorTypesRewardsSubmission) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.CreateAVSRewardsSubmission(&_ContractTriggerXServiceManager.TransactOpts, rewardsSubmissions)
}

// CreateAVSRewardsSubmission is a paid mutator transaction binding the contract method 0xfce36c7d.
//
// Solidity: function createAVSRewardsSubmission(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactorSession) CreateAVSRewardsSubmission(rewardsSubmissions []IRewardsCoordinatorTypesRewardsSubmission) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.CreateAVSRewardsSubmission(&_ContractTriggerXServiceManager.TransactOpts, rewardsSubmissions)
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

// Initialize is a paid mutator transaction binding the contract method 0x8a29e2de.
//
// Solidity: function initialize(address _taskManagerContract, address initialOwner, address rewardsInitiator, address _slasher, address _vetoCommittee, address _taskManager, address _taskValidator, address _quorumManager) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactor) Initialize(opts *bind.TransactOpts, _taskManagerContract common.Address, initialOwner common.Address, rewardsInitiator common.Address, _slasher common.Address, _vetoCommittee common.Address, _taskManager common.Address, _taskValidator common.Address, _quorumManager common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.contract.Transact(opts, "initialize", _taskManagerContract, initialOwner, rewardsInitiator, _slasher, _vetoCommittee, _taskManager, _taskValidator, _quorumManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x8a29e2de.
//
// Solidity: function initialize(address _taskManagerContract, address initialOwner, address rewardsInitiator, address _slasher, address _vetoCommittee, address _taskManager, address _taskValidator, address _quorumManager) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) Initialize(_taskManagerContract common.Address, initialOwner common.Address, rewardsInitiator common.Address, _slasher common.Address, _vetoCommittee common.Address, _taskManager common.Address, _taskValidator common.Address, _quorumManager common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.Initialize(&_ContractTriggerXServiceManager.TransactOpts, _taskManagerContract, initialOwner, rewardsInitiator, _slasher, _vetoCommittee, _taskManager, _taskValidator, _quorumManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x8a29e2de.
//
// Solidity: function initialize(address _taskManagerContract, address initialOwner, address rewardsInitiator, address _slasher, address _vetoCommittee, address _taskManager, address _taskValidator, address _quorumManager) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactorSession) Initialize(_taskManagerContract common.Address, initialOwner common.Address, rewardsInitiator common.Address, _slasher common.Address, _vetoCommittee common.Address, _taskManager common.Address, _taskValidator common.Address, _quorumManager common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.Initialize(&_ContractTriggerXServiceManager.TransactOpts, _taskManagerContract, initialOwner, rewardsInitiator, _slasher, _vetoCommittee, _taskManager, _taskValidator, _quorumManager)
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

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactor) RemoveAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.contract.Transact(opts, "removeAdmin", admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.RemoveAdmin(&_ContractTriggerXServiceManager.TransactOpts, admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactorSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.RemoveAdmin(&_ContractTriggerXServiceManager.TransactOpts, admin)
}

// RemoveAppointee is a paid mutator transaction binding the contract method 0xba550880.
//
// Solidity: function removeAppointee(address appointee, address target, bytes4 selector) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactor) RemoveAppointee(opts *bind.TransactOpts, appointee common.Address, target common.Address, selector [4]byte) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.contract.Transact(opts, "removeAppointee", appointee, target, selector)
}

// RemoveAppointee is a paid mutator transaction binding the contract method 0xba550880.
//
// Solidity: function removeAppointee(address appointee, address target, bytes4 selector) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) RemoveAppointee(appointee common.Address, target common.Address, selector [4]byte) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.RemoveAppointee(&_ContractTriggerXServiceManager.TransactOpts, appointee, target, selector)
}

// RemoveAppointee is a paid mutator transaction binding the contract method 0xba550880.
//
// Solidity: function removeAppointee(address appointee, address target, bytes4 selector) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactorSession) RemoveAppointee(appointee common.Address, target common.Address, selector [4]byte) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.RemoveAppointee(&_ContractTriggerXServiceManager.TransactOpts, appointee, target, selector)
}

// RemovePendingAdmin is a paid mutator transaction binding the contract method 0x9da16d8e.
//
// Solidity: function removePendingAdmin(address pendingAdmin) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactor) RemovePendingAdmin(opts *bind.TransactOpts, pendingAdmin common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.contract.Transact(opts, "removePendingAdmin", pendingAdmin)
}

// RemovePendingAdmin is a paid mutator transaction binding the contract method 0x9da16d8e.
//
// Solidity: function removePendingAdmin(address pendingAdmin) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) RemovePendingAdmin(pendingAdmin common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.RemovePendingAdmin(&_ContractTriggerXServiceManager.TransactOpts, pendingAdmin)
}

// RemovePendingAdmin is a paid mutator transaction binding the contract method 0x9da16d8e.
//
// Solidity: function removePendingAdmin(address pendingAdmin) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactorSession) RemovePendingAdmin(pendingAdmin common.Address) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.RemovePendingAdmin(&_ContractTriggerXServiceManager.TransactOpts, pendingAdmin)
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

// SetAppointee is a paid mutator transaction binding the contract method 0x1fdb0cfd.
//
// Solidity: function setAppointee(address appointee, address target, bytes4 selector) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactor) SetAppointee(opts *bind.TransactOpts, appointee common.Address, target common.Address, selector [4]byte) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.contract.Transact(opts, "setAppointee", appointee, target, selector)
}

// SetAppointee is a paid mutator transaction binding the contract method 0x1fdb0cfd.
//
// Solidity: function setAppointee(address appointee, address target, bytes4 selector) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerSession) SetAppointee(appointee common.Address, target common.Address, selector [4]byte) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.SetAppointee(&_ContractTriggerXServiceManager.TransactOpts, appointee, target, selector)
}

// SetAppointee is a paid mutator transaction binding the contract method 0x1fdb0cfd.
//
// Solidity: function setAppointee(address appointee, address target, bytes4 selector) returns()
func (_ContractTriggerXServiceManager *ContractTriggerXServiceManagerTransactorSession) SetAppointee(appointee common.Address, target common.Address, selector [4]byte) (*types.Transaction, error) {
	return _ContractTriggerXServiceManager.Contract.SetAppointee(&_ContractTriggerXServiceManager.TransactOpts, appointee, target, selector)
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
