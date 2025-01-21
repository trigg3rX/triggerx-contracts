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
	Bin: "0x60806040526004361015610011575f80fd5b5f3560e01c8063136439dd14610334578063171f1d5b1461032f5780631785f53c1461032a5780631fdb0cfd1461032557806326a965f014610320578063279432eb1461031b578063327d0a601461031657806333cfb7b7146103115780633bc28c8c1461030c578063416c7e5e14610307578063500c8dd31461030257806358ac4a1e146102fd578063595c6a67146102f85780635ac86ab7146102f35780635c975abb146102ee5780635df45946146102e95780635edd78ea146102e457806367940c89146102df57806368304835146102da5780636a8dcf54146102d55780636b3aa72e146102d05780636d14a987146102cb5780636efb4636146102c6578063715018a6146102c1578063886f1195146102bc5780638a29e2de146102b75780638d68349a146102b25780638da5cb5b146102ad5780639926ee7d146102a85780639cd0fb86146102a35780639da16d8e1461029e578063a364f4da14610299578063a41d3f9414610294578063a50a640e1461028f578063a98fb3551461028a578063b134427114610285578063b5f7eb6b14610280578063b7d79c931461027b578063b98d090814610276578063ba55088014610271578063df5cf7231461026c578063e46f181614610267578063e47d606014610262578063e481af9d1461025d578063f2fde38b14610258578063fabc1cbc14610253578063fc299dee1461024e578063fcd1c37514610249578063fce36c7d146102445763fd38ec8c1461023f575f80fd5b611c2d565b611bd2565b611bb5565b611b8d565b611aab565b611a1a565b6119ff565b6119bf565b611997565b611953565b6118c7565b6118a5565b611861565b611838565b611810565b611768565b611740565b6116d7565b61158d565b611503565b611495565b6113f9565b611380565b61135e565b611218565b6111d4565b611179565b6110e4565b610d8e565b610d4a565b610d21565b610cdd565b610cc0565b610c53565b610c0f565b610bf2565b610bbf565b610b41565b610ad3565b610aaa565b6109aa565b610973565b61093b565b61088c565b610802565b610796565b6106f6565b610615565b61059c565b346103f45760203660031901126103f45760043560405163237dfb4760e11b8152336004820152906020826024817f0000000000000000000000003db5b2115c2c4619255c107c8daade9dd35b366c6001600160a01b03165afa9182156103ef576103be926103aa915f916103c0575b50611c78565b6103b960015482811614611c8e565b6136c8565b005b6103e2915060203d6020116103e8575b6103da8183610447565b810190611c55565b5f6103a4565b503d6103d0565b611c6d565b5f80fd5b634e487b7160e01b5f52604160045260245ffd5b604081019081106001600160401b0382111761042757604052565b6103f8565b606081019081106001600160401b0382111761042757604052565b90601f801991011681019081106001600160401b0382111761042757604052565b6040519061047861010083610447565b565b60405190610478604083610447565b906104786040519283610447565b60409060e31901126103f457604051906104b08261040c565b60e4358252610104356020830152565b91908260409103126103f4576040516104d88161040c565b6020808294803584520135910152565b9080601f830112156103f45760405191610503604084610447565b8290604081019283116103f457905b82821061051f5750505090565b8135815260209182019101610512565b9060806063198301126103f4576040516105488161040c565b6020610563829461055a8160646104e8565b845260a46104e8565b910152565b91906080838203126103f4576020610563604051926105868461040c565b6040849661059483826104e8565b8652016104e8565b346103f4576101203660031901126103f45760043560403660231901126103f4576105f460409182516105ce8161040c565b602435815260443560208201526105e43661052f565b906105ee36610497565b92611ce2565b8251911515825215156020820152f35b6001600160a01b038116036103f457565b346103f4575f60203660031901126103f45760043561063381610604565b61063b613b71565b7f000000000000000000000000598cb226b591155f767da17afe7a2241a68c5c106001600160a01b0316803b156103f45760405163268959e560e01b81523060048201526001600160a01b0390921660248301525f908290818381604481015b03925af180156103ef576106ad575080f35b6103be91505f90610447565b60609060031901126103f4576004356106d181610604565b906024356106de81610604565b906044356001600160e01b0319811681036103f45790565b346103f457610704366106b9565b61070c613b71565b7f000000000000000000000000598cb226b591155f767da17afe7a2241a68c5c106001600160a01b031692833b156103f457604051634a86c03760e11b81523060048201526001600160a01b039182166024820152921660448301526001600160e01b03191660648201525f81608481015b93818381819703925af180156103ef576106ad575080f35b346103f45760203660031901126103f4576004356107b381610604565b6107bb613b71565b6001600160a01b03165f81815260fd60205260408120805460ff191660011790557f79bf5efab685b39bdd2b7b1aba09ffa5e9dfbd86f7aaebab4b52d7ef679d28449080a2005b346103f4575f60203660031901126103f45760043561082081610604565b610828613b71565b7f000000000000000000000000598cb226b591155f767da17afe7a2241a68c5c106001600160a01b0316803b156103f45760405163eb5a4e8760e01b81523060048201526001600160a01b0390921660248301525f9082908183816044810161069b565b346103f45760203660031901126103f4576004356108a981610604565b6108b1613b71565b60fe80546001600160a01b039283166001600160a01b0319821681179092559091167fdae652729007c1e5ea13dfe9e126a1b72358db8db69c68b96f316dcea59c1c5f5f80a3005b60206040818301928281528451809452019201905f5b81811061091c5750505090565b82516001600160a01b031684526020938401939092019160010161090f565b346103f45760203660031901126103f45761096f61096360043561095e81610604565b611f42565b604051918291826108f9565b0390f35b346103f45760203660031901126103f4576103be60043561099381610604565b61099b613b71565b613c67565b801515036103f457565b346103f45760203660031901126103f4576004356109c7816109a0565b604051638da5cb5b60e01b81526020816004817f000000000000000000000000b438c6fc1652148bb758b939831f9a2cd59ce02b6001600160a01b03165afa9081156103ef575f91610a71575b506001600160a01b03163303610a625760207f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc91151560ff196032541660ff821617603255604051908152a1005b637070f3b160e11b5f5260045ffd5b610a93915060203d602011610a99575b610a8b8183610447565b8101906122cf565b5f610a14565b503d610a81565b5f9103126103f457565b346103f4575f3660031901126103f457610102546040516001600160a01b039091168152602090f35b346103f45760203660031901126103f457600435610af081610604565b610af8613b71565b61010280546001600160a01b039283166001600160a01b0319821681179092559091167fecc36ab0616c35e413d90730ada58214cbe0188c5f26632fbb32451c9a1357845f80a3005b346103f4575f3660031901126103f45760405163237dfb4760e11b81523360048201526020816024817f0000000000000000000000003db5b2115c2c4619255c107c8daade9dd35b366c6001600160a01b03165afa80156103ef57610bac915f916103c05750611c78565b6103be613694565b60ff8116036103f457565b346103f45760203660031901126103f4576020600160ff600435610be281610bb4565b161b806001541614604051908152f35b346103f4575f3660031901126103f4576020600154604051908152f35b346103f4575f3660031901126103f4576040517f000000000000000000000000c275cce555cc4bd954ff078621778368f9006f1c6001600160a01b03168152602090f35b346103f45760203660031901126103f457600435610c7081610604565b610c78613b71565b60ff80546001600160a01b039283166001600160a01b0319821681179092559091167f86cebbccc5bbc9ac544f7e5361e6c5e4acf14f39d85c855032a815125cb97fd85f80a3005b346103f4575f3660031901126103f457602060405162093a808152f35b346103f4575f3660031901126103f4576040517f000000000000000000000000d7c48ab86ab9390e09165013600d5e8721cb3e1d6001600160a01b03168152602090f35b346103f4575f3660031901126103f457610101546040516001600160a01b039091168152602090f35b346103f4575f3660031901126103f4576040517f000000000000000000000000055733000064333caddbc92763c58bf0192ffebf6001600160a01b03168152602090f35b346103f4575f3660031901126103f4576040517f000000000000000000000000b438c6fc1652148bb758b939831f9a2cd59ce02b6001600160a01b03168152602090f35b63ffffffff8116036103f457565b6044359061047882610dd2565b359061047882610dd2565b6001600160401b0381116104275760051b60200190565b9080601f830112156103f4578135610e2681610df8565b92610e346040519485610447565b81845260208085019260051b8201019283116103f457602001905b828210610e5c5750505090565b602080918335610e6b81610dd2565b815201910190610e4f565b81601f820112156103f4578035610e8c81610df8565b92610e9a6040519485610447565b81845260208085019260061b840101928184116103f457602001915b838310610ec4575050505090565b6020604091610ed384866104c0565b815201920191610eb6565b9080601f830112156103f4578135610ef581610df8565b92610f036040519485610447565b81845260208085019260051b820101918383116103f45760208201905b838210610f2f57505050505090565b81356001600160401b0381116103f457602091610f5187848094880101610e0f565b815201910190610f20565b919091610180818403126103f457610f72610468565b9281356001600160401b0381116103f45781610f8f918401610e0f565b845260208201356001600160401b0381116103f45781610fb0918401610e76565b602085015260408201356001600160401b0381116103f45781610fd4918401610e76565b6040850152610fe68160608401610568565b6060850152610ff88160e084016104c0565b60808501526101208201356001600160401b0381116103f4578161101d918401610e0f565b60a08501526101408201356001600160401b0381116103f45781611042918401610e0f565b60c08501526101608201356001600160401b0381116103f4576110659201610ede565b60e0830152565b90602080835192838152019201905f5b8181106110895750505090565b82516001600160601b031684526020938401939092019160010161107c565b9291906110df60209160408652826110cb82516040808a0152608089019061106c565b910151868203603f1901606088015261106c565b930152565b346103f45760803660031901126103f4576004356024356001600160401b0381116103f457366023820112156103f45780600401356001600160401b0381116103f45736602482840101116103f45761113b610de0565b90606435936001600160401b0385116103f4576024611161611169963690600401610f5c565b940190612487565b9061096f604051928392836110a8565b346103f4575f3660031901126103f457611191613b71565b609780546001600160a01b031981169091555f906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b346103f4575f3660031901126103f4576040517f0000000000000000000000003db5b2115c2c4619255c107c8daade9dd35b366c6001600160a01b03168152602090f35b346103f4576101003660031901126103f45760043561123681610604565b6112d260243561124581610604565b60443561125181610604565b60643561125d81610604565b60843561126981610604565b60a4359161127683610604565b60c4359361128385610604565b60e4359561129087610604565b606454986112b660ff8b60081c16151515809b81611350575b8115611330575b50612d67565b896112c9600160ff196064541617606455565b61131757612dca565b6112d857005b6112e861ff001960645416606455565b604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602090a1005b61132b61010061ff00196064541617606455565b612dca565b303b15915081611342575b505f6112b0565b60ff1660011490505f61133b565b600160ff82161091506112a9565b346103f4575f3660031901126103f457602060ff60cd54166040519015158152f35b346103f4575f3660031901126103f4576097546040516001600160a01b039091168152602090f35b6001600160401b03811161042757601f01601f191660200190565b9291926113cf826113a8565b916113dd6040519384610447565b8294818452818301116103f4578281602093845f960137010152565b346103f45760403660031901126103f45760043561141681610604565b602435906001600160401b0382116103f457606060031983360301126103f457604051906114438261042c565b82600401356001600160401b0381116103f4578301366023820112156103f4576103be9361147d60449236906024600482013591016113c3565b84526024810135602085015201356040830152612f26565b346103f45760203660031901126103f4576004356114b281610604565b6114ba613b71565b61010080546001600160a01b039283166001600160a01b0319821681179092559091167f5d68f69aebb8a1d75f8b2a87dddd6e8db842f9fa3bb63745ea7f49ea5bdf03a95f80a3005b346103f4575f60203660031901126103f45760043561152181610604565b611529613b71565b7f000000000000000000000000598cb226b591155f767da17afe7a2241a68c5c106001600160a01b0316803b156103f457604051634f906cf960e01b81523060048201526001600160a01b0390921660248301525f9082908183816044810161069b565b346103f45760203660031901126103f4576004356115aa81610604565b6115de337f000000000000000000000000b438c6fc1652148bb758b939831f9a2cd59ce02b6001600160a01b031614612f10565b7f000000000000000000000000055733000064333caddbc92763c58bf0192ffebf6001600160a01b0316803b156103f4576040516351b27a6d60e11b81526001600160a01b0383166004820152905f908290602490829084905af180156103ef576116bd575b506116676116626001600160a01b0383165b6001600160a01b031690565b614020565b506001600160a01b0381165f90815260fd6020526040902061168e905b805460ff19169055565b6001600160a01b03167fa7a775c2c8141f7985c111748ec31c11e5e44b83528e105c8d1d4e8e6b81cf805f80a2005b806116cb5f6116d193610447565b80610aa0565b5f611644565b346103f45760203660031901126103f4576004356116f481610604565b6116fc613b71565b6001600160a01b03165f81815260fd60205260408120805460ff191690557f8511046d95f95689c89b4807d91820adbc266db891b2c5c25a42f812781dde579080a2005b346103f4575f3660031901126103f45760fe546040516001600160a01b039091168152602090f35b346103f4575f60203660031901126103f4576004356001600160401b0381116103f457366023820112156103f4576117aa9036906024816004013591016113c3565b6117b2613b71565b7f000000000000000000000000055733000064333caddbc92763c58bf0192ffebf6001600160a01b0316803b156103f45760405163a98fb35560e01b815260206004820152915f91839182908490829061069b90602483019061306f565b346103f4575f3660031901126103f45760ca546040516001600160a01b039091168152602090f35b346103f4575f3660031901126103f457610100546040516001600160a01b039091168152602090f35b346103f4575f3660031901126103f4576040517f000000000000000000000000ea80328759f6269ca985e631944ba0d23c933bd46001600160a01b03168152602090f35b346103f4575f3660031901126103f457602060ff603254166040519015158152f35b346103f4576118d5366106b9565b6118dd613b71565b7f000000000000000000000000598cb226b591155f767da17afe7a2241a68c5c106001600160a01b031692833b156103f457604051630664120160e01b81523060048201526001600160a01b039182166024820152921660448301526001600160e01b03191660648201525f816084810161077e565b346103f4575f3660031901126103f4576040517f000000000000000000000000a44151489861fe9e3055d95adc98fbd462b948e76001600160a01b03168152602090f35b346103f4575f3660031901126103f45760cb546040516001600160a01b039091168152602090f35b346103f45760203660031901126103f4576004356119dc81610604565b60018060a01b03165f5260fd602052602060ff60405f2054166040519015158152f35b346103f4575f3660031901126103f45761096f610963613093565b346103f45760203660031901126103f457600435611a3781610604565b611a3f613b71565b6001600160a01b03811615611a57576103be90613e63565b60405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608490fd5b346103f45760203660031901126103f45760043560405163755b36bd60e11b81526020816004817f0000000000000000000000003db5b2115c2c4619255c107c8daade9dd35b366c6001600160a01b03165afa9081156103ef575f91611b6e575b506001600160a01b03163303611b5f57611b2d600154198219811614611c8e565b806001556040519081527f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c60203392a2005b63794821ff60e01b5f5260045ffd5b611b87915060203d602011610a9957610a8b8183610447565b5f611b0c565b346103f4575f3660031901126103f45760c9546040516001600160a01b039091168152602090f35b346103f4575f3660031901126103f457602060cc54604051908152f35b346103f45760203660031901126103f4576004356001600160401b0381116103f457366023820112156103f45780600401356001600160401b0381116103f4573660248260051b840101116103f45760246103be92016132c0565b346103f4575f3660031901126103f45760ff546040516001600160a01b039091168152602090f35b908160209103126103f45751611c6a816109a0565b90565b6040513d5f823e3d90fd5b15611c7f57565b631d77d47760e21b5f5260045ffd5b15611c9557565b63c61dca5d60e01b5f5260045ffd5b634e487b7160e01b5f52603260045260245ffd5b906002811015611cc95760051b0190565b611ca4565b634e487b7160e01b5f52601260045260245ffd5b611dbe611d9b611dc495611d95611d8e85875160208901518a515160208c51015160208d016020815151915101519189519360208b0151956040519760208901998a5260208a015260408901526060880152608087015260a086015260c085015260e0840152610100830152611d6581610120840103601f198101835282610447565b5190207f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001900690565b809661373e565b90613784565b92611d95611db0611daa6137e6565b946138dd565b91611db96139f9565b61373e565b91613a43565b9091565b908160209103126103f4575190565b908160209103126103f457516001600160c01b03811681036103f45790565b908160209103126103f45751611c6a81610bb4565b60405190611e1a602083610447565b5f808352366020840137565b90611e3082610df8565b611e3d6040519182610447565b8281528092611e4e601f1991610df8565b0190602036910137565b908151811015611cc9570160200190565b634e487b7160e01b5f52601160045260245ffd5b9060018201809211611e8b57565b611e69565b9060028201809211611e8b57565b9060038201809211611e8b57565b9060048201809211611e8b57565b9060058201809211611e8b57565b91908201809211611e8b57565b6001600160601b038116036103f457565b908160409103126103f457602060405191611f008361040c565b8051611f0b81610604565b83520151611f1881611ed5565b602082015290565b8051821015611cc95760209160051b010190565b5f198114611e8b5760010190565b6040516309aa152760e11b81526001600160a01b0391821660048201527f000000000000000000000000b438c6fc1652148bb758b939831f9a2cd59ce02b90911690602081602481855afa9081156103ef57611fc2916020915f916122b2575b506040518093819263871ef04960e01b8352600483019190602083019252565b0381855afa9081156103ef575f91612283575b506001600160c01b0316908115908115612220575b5061221757611ff890613bc9565b5f91907f000000000000000000000000d7c48ab86ab9390e09165013600d5e8721cb3e1d6001600160a01b031690835b81518510156120d45761207d602061205a6120546120468987611e58565b516001600160f81b03191690565b60f81c90565b604051633ca5a5f560e01b815260ff909116600482015291829081906024820190565b0381875afa80156103ef5760019261209c925f926120a4575b50611ec8565b940193612028565b6120c691925060203d81116120cd575b6120be8183610447565b810190611dc8565b905f612096565b503d6120b4565b6120df919450611e26565b925f905f5b8151811015612211576120fd6120546120468385611e58565b604051633ca5a5f560e01b815260ff8216600482015290602082602481895afa9182156103ef575f926121f1575b50905f915b818310612142575050506001016120e4565b604080516356e4026d60e11b815260ff83166004820152602481018590529396929391929190816044818b5afa9182156103ef576121b58b6121a6836121a06116566001986121ba985f916121c3575b50516001600160a01b031690565b92611f20565b6001600160a01b039091169052565b611f34565b95019190612130565b6121e4915060403d81116121ea575b6121dc8183610447565b810190611ee6565b5f612192565b503d6121d2565b61220a91925060203d81116120cd576120be8183610447565b905f61212b565b50505050565b50611c6a611e0b565b604051639aa1653d60e01b81529150602090829060049082905afa80156103ef5760ff915f91612254575b5016155f611fea565b612276915060203d60201161227c575b61226e8183610447565b810190611df6565b5f61224b565b503d612264565b6122a5915060203d6020116122ab575b61229d8183610447565b810190611dd7565b5f611fd5565b503d612293565b6122c99150823d84116120cd576120be8183610447565b5f611fa2565b908160209103126103f45751611c6a81610604565b604051906122f18261040c565b60606020838281520152565b1561230457565b62f8202d60e51b5f5260045ffd5b1561231957565b6343714afd60e01b5f5260045ffd5b1561232f57565b635f832f4160e01b5f5260045ffd5b1561234557565b634b874f4560e01b5f5260045ffd5b5f19810191908211611e8b57565b1561236957565b633fdc650560e21b5f5260045ffd5b908160209103126103f45751611c6a81610dd2565b90821015611cc9570190565b156123a057565b63affc5edb60e01b5f5260045ffd5b908160209103126103f4575167ffffffffffffffff19811681036103f45790565b156123d757565b63e1310aed60e01b5f5260045ffd5b908160209103126103f45751611c6a81611ed5565b906001600160601b03809116911603906001600160601b038211611e8b57565b1561242257565b6367988d3360e01b5f5260045ffd5b1561243857565b63ab1b236b60e01b5f5260045ffd5b60049163ffffffff60e01b9060e01b1681520160208251919201905f5b8181106124715750505090565b8251845260209384019390920191600101612464565b9493929091936124956122e4565b506124a18515156122fd565b604084015151851480612d59575b80612d4b575b80612d3d575b6124c490612312565b6124d660208501515185515114612328565b6124ed63ffffffff431663ffffffff84161061233e565b6124f561047a565b5f81525f6020820152926125076122e4565b61251087611e26565b602082015261251e87611e26565b81526125286122e4565b92612537602088015151611e26565b8452612547602088015151611e26565b602085810191909152604051639aa1653d60e01b815290816004817f000000000000000000000000b438c6fc1652148bb758b939831f9a2cd59ce02b6001600160a01b03165afa80156103ef576125b0915f91612d1e575b506125ab368b876113c3565b613cc5565b985f965b6020890151805189101561272c576020886126216126178c61260f8f96868e6125f46125e1868095611f20565b5180515f526020015160205260405f2090565b6126018484840151611f20565b52826126f9575b0151611f20565b519551611f20565b5163ffffffff1690565b6040516304ec635160e01b8152600481019490945263ffffffff9182166024850152166044830152816064816001600160a01b037f000000000000000000000000b438c6fc1652148bb758b939831f9a2cd59ce02b165afa9182156103ef57611d958a6126ce8f6126c78f8460208f926126be936126b68460019e6126d49e5f916126dc575b508f8060c01b03169251611f20565b520151611f20565b51938d51611f20565b5116613cf0565b90613d21565b9701966125b4565b6126f39150863d81116122ab5761229d8183610447565b5f6126a7565b6127276127098484840151611f20565b516127208484015161271a87612354565b90611f20565b5110612362565b612608565b50909597949650612741919893929950613dde565b9161274e60325460ff1690565b908115612d16576040516318891fd760e31b81526020816004817f000000000000000000000000a44151489861fe9e3055d95adc98fbd462b948e76001600160a01b03165afa9081156103ef575f91612ce7575b5091905b5f925b81841061280d575050505050926127e66127e16127da61280795856127f99860806060602099015192015192611ce2565b919061241b565b612431565b0151604051928391602083019586612447565b03601f198101835282610447565b51902090565b92989596909399919794878b888c888d612be1575b6126178260a0612870612054612862846128789761285c61284e6125e18f9c604060209f9e0151611f20565b67ffffffffffffffff191690565b9b61238d565b356001600160f81b03191690565b970151611f20565b604051631a2f32ab60e21b815260ff95909516600486015263ffffffff9182166024860152166044840152826064816001600160a01b037f000000000000000000000000c275cce555cc4bd954ff078621778368f9006f1c165afa9081156103ef5761293c6126178f958f906129348f978f96848f61292e60c096612927848f60209f90612608612862996040936120549c5f91612bb3575b5067ffffffffffffffff199182169116146123d0565b5190613784565b9c61238d565b960151611f20565b604051636414a62b60e11b815260ff94909416600485015263ffffffff9182166024850152166044830152816064816001600160a01b037f000000000000000000000000d7c48ab86ab9390e09165013600d5e8721cb3e1d165afa9081156103ef576129c9918c8f925f92612b8f575b5060206129bb92930151611f20565b906001600160601b03169052565b6129f68c6129bb8c6129ef6129e2826020860151611f20565b516001600160601b031690565b9251611f20565b5f985f5b60208a015151811015612b76578b8d612a3889612a2b612054612862868f89612a239151611f20565b51948761238d565b60ff161c60019081161490565b612a47575b50506001016129fa565b8a8a612acf859f948f9686612a898f9360e0612a80612617956020612a78612054612862839f612a8f9c899161238d565b9a0151611f20565b519b0151611f20565b51611f20565b60405163795f4a5760e11b815260ff909316600484015263ffffffff93841660248401526044830196909652919094166064850152839081906084820190565b03817f000000000000000000000000d7c48ab86ab9390e09165013600d5e8721cb3e1d6001600160a01b03165afa9081156103ef578f612b35908f936001959486955f92612b40575b506121a06129bb92935193612b306129e28487611f20565b6123fb565b019a90508b8d612a3d565b6129bb9250612b686121a09160203d8111612b6f575b612b608183610447565b8101906123e6565b9250612b18565b503d612b56565b5093919796996001919699509a94929a019291906127a9565b6129bb9250612bac602091823d8111612b6f57612b608183610447565b92506129ac565b6020612bd492503d8111612bda575b612bcc8183610447565b8101906123af565b5f612911565b503d612bc2565b612c1e9450612bfb9250612054916128629160209561238d565b60405163124d062160e11b815260ff909116600482015291829081906024820190565b03817f000000000000000000000000b438c6fc1652148bb758b939831f9a2cd59ce02b6001600160a01b03165afa80156103ef576020896128788f938f60a08f976120546128628f8f9061285c61284e6125e18f60408b96918f88936126179f612ca290612ca8936128709f5f92612cbe575b5063ffffffff809116931690611ec8565b11612399565b5050505050509750505050505092935050612822565b602063ffffffff9293508291612cdf913d81116120cd576120be8183610447565b929150612c91565b612d09915060203d602011612d0f575b612d018183610447565b810190612378565b5f6127a2565b503d612cf7565b5f91906127a6565b612d37915060203d60201161227c5761226e8183610447565b5f61259f565b5060e08401515185146124bb565b5060c08401515185146124b5565b5060a08401515185146124af565b15612d6e57565b60405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608490fd5b929395969196612dd982613e63565b60ff60645460081c1615612eb757612e3f612e9896612e7b95612e056104789b61099b612e5c97613e63565b60018060a01b03166001600160601b0360a01b6101025416176101025560018060a01b03166001600160601b0360a01b60fe54161760fe55565b60018060a01b03166001600160601b0360a01b60ff54161760ff55565b60018060a01b03166001600160601b0360a01b61010054161761010055565b60018060a01b03166001600160601b0360a01b60ca54161760ca55565b60018060a01b03166001600160601b0360a01b61010154161761010155565b60405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b6064820152608490fd5b15612f1757565b634394dbdf60e11b5f5260045ffd5b90612f5b337f000000000000000000000000b438c6fc1652148bb758b939831f9a2cd59ce02b6001600160a01b031614612f10565b7f000000000000000000000000055733000064333caddbc92763c58bf0192ffebf6001600160a01b0316803b156103f4575f9182916040518094818094639926ee7d60e01b82526004820160018060a01b038a1681526040602082015260806040612fd1845160608386015260a085019061306f565b9360208101516060850152015191015203925af180156103ef5761305b575b5061300b6130066001600160a01b038316611656565b613f6b565b506001600160a01b0381165f90815260fd6020526040902061302c90611684565b6001600160a01b03167f1584773458d98c71b34a270ee1100b3a42889bf91e3b7a858563b684c24d838e5f80a2565b806116cb5f61306993610447565b5f612ff0565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b604051639aa1653d60e01b81527f000000000000000000000000b438c6fc1652148bb758b939831f9a2cd59ce02b6001600160a01b031690602081600481855afa80156103ef5760ff915f916132a1575b50168015613297577f000000000000000000000000d7c48ab86ab9390e09165013600d5e8721cb3e1d6001600160a01b0316905f9081905b808310613253575061312e9150611e26565b925f905f5b604051639aa1653d60e01b8152602081600481895afa80156103ef5760ff915f91613235575b501681101561322e57604051633ca5a5f560e01b815260ff821660048201819052602082602481895afa9182156103ef575f9261320e575b50905f915b8183106131a857505050600101613133565b604080516356e4026d60e11b815260ff83166004820152602481018590529396929391929190816044818b5afa9182156103ef576121b58b6121a6836121a0611656600198613205985f916121c35750516001600160a01b031690565b95019190613196565b61322791925060203d81116120cd576120be8183610447565b905f613191565b5092505050565b61324d915060203d811161227c5761226e8183610447565b5f613159565b604051633ca5a5f560e01b815260ff84166004820152909190602081602481885afa80156103ef5760019261328e925f926120a45750611ec8565b9201919061311c565b5050611c6a611e0b565b6132ba915060203d60201161227c5761226e8183610447565b5f6130e4565b60c9546001600160a01b031633036134e2577f000000000000000000000000acc1fb458a1317e886db376fc8141540537e68fe6001600160a01b0316915f5b8181106133505750823b156103f457613331925f928360405180968195829463fce36c7d60e01b845260048401613585565b03925af180156103ef576133425750565b806116cb5f61047893610447565b5f60206133ab61336e6116568361336887898b6134f1565b01613513565b604061337b86888a6134f1565b6040516323b872dd60e01b8152336004820152306024820152910135604482015293849283919082906064820190565b03925af180156103ef576134c6575b506133ce61165660206133688486886134f1565b604051636eb1769f60e11b81523060048201526001600160a01b03861660248201529190602090839060449082905afa80156103ef576134716020915f9485916134a9575b5061343e61342961165685613368888b8d6134f1565b916040613437878a8c6134f1565b0135611ec8565b60405163095ea7b360e01b81526001600160a01b038a166004820152602481019190915294859283919082906044820190565b03925af19182156103ef5760019261348b575b50016132ff565b6134a29060203d81116103e8576103da8183610447565b505f613484565b6134c09150833d81116120cd576120be8183610447565b5f613413565b6134dd9060203d81116103e8576103da8183610447565b6133ba565b638e79fdb560e01b5f5260045ffd5b9190811015611cc95760051b81013590609e19813603018212156103f4570190565b35611c6a81610604565b916020908281520191905f5b8181106135365750505090565b909192604080600192863561354a81610604565b848060a01b031681526001600160601b03602088013561356981611ed5565b166020820152019401929101613529565b359061047882610604565b9180602084016020855252604083019060408160051b85010193835f91609e1982360301905b8484106135bc575050505050505090565b90919293949596603f198282030187528735838112156103f4578401908135601e19833603018112156103f457820191602083359301906001600160401b0384116103f4578360061b360382136103f45761368583608061367a816136306020989760019a60a08b9a5260a087019161351d565b9561364e61363f89830161357a565b6001600160a01b0316868a0152565b6040810135604086015261367461366760608301610ded565b63ffffffff166060870152565b01610ded565b63ffffffff16910152565b990197019594019291906135ab565b5f196001556040515f1981527fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d60203392a2565b806001556040519081527fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d60203392a2565b604051906137078261040c565b5f6020838281520152565b604051906101806137238184610447565b368337565b60405190613737602083610447565b6020368337565b9190604090606061374d6136fa565b948592602085519261375f8585610447565b8436853780518452015160208301528482015260076107cf195a01fa1561378257565bfe5b6020929160806040926137956136fa565b958693818651936137a68686610447565b85368637805185520151828401528051868401520151606082015260066107cf195a01fa801561378257156137d757565b63d4b68fd760e01b5f5260045ffd5b6040516137f28161040c565b60409081516138018382610447565b82368237815260208251916138168484610447565b83368437015280516138288282610447565b7f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c281527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed602082015281519061387e8383610447565b7f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208301526138d383519384610447565b8252602082015290565b5f5160206140d85f395f51905f52906138f46136fa565b505f919006602060c0835b6139f4575f935f5160206140d85f395f51905f526003818681818009090860405161392a8582610447565b8436823784818560405161393e8282610447565b813682378381528360208201528360408201528560608201527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808201525f5160206140d85f395f51905f5260a082015260056107cf195a01fa8015613782576139a8906140c1565b51916139f4575f5160206140d85f395f51905f52828009146139df57505f5160206140d85f395f51905f5260015f940892936138ff565b929350506139eb61047a565b92835282015290565b611cce565b613a016136fa565b50604051613a0e8161040c565b600181526002602082015290565b90600682029180830460061490151715611e8b57565b90600c811015611cc95760051b0190565b93929091613a516040610489565b9485526020850152613a636040610489565b9182526020820152613a73613712565b925f5b60028110613aa057505050602061018092613a8f613728565b93849160086201d4c0fa9151151590565b80613aac600192613a1c565b613ab68285611cb8565b5151613ac28289613a32565b526020613acf8386611cb8565b510151613ae4613ade83611e7d565b89613a32565b52613aef8286611cb8565b515151613afe613ade83611e90565b52613b14613b0c8387611cb8565b515160200190565b51613b21613ade83611e9e565b526020613b2e8387611cb8565b51015151613b3e613ade83611eac565b52613b6a613b64613b5d6020613b54868a611cb8565b51015160200190565b5192611eba565b88613a32565b5201613a76565b6097546001600160a01b03163303613b8557565b606460405162461bcd60e51b815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b61ffff613bd582613cf0565b16613bdf816113a8565b90613bed6040519283610447565b808252613bfc601f19916113a8565b013660208301375f5f5b8251821080613c5c575b15613c55576001811b8416613c2e575b613c2990611f34565b613c06565b906001613c299160ff60f81b8460f81b165f1a613c4b8287611e58565b5301919050613c20565b5050905090565b506101008110613c10565b60c954604080516001600160a01b038084168252841660208201529192917fe11cddf1816a43318ca175bbc52cd0185436e9cbead7c83acc54a73e461717e39190a16001600160a01b03166001600160a01b0319919091161760c955565b906001613cd360ff93613ec1565b928392161b1115613ce15790565b63ca95733360e01b5f5260045ffd5b805f915b613cfc575090565b5f198101818111611e8b5761ffff9116911661ffff8114611e8b576001019080613cf4565b90613d2a6136fa565b5061ffff811690610200821015613dcf5760018214613dca57613d4b61047a565b5f81525f602082015292906001905f925b61ffff8316851015613d7057505050505090565b600161ffff831660ff86161c811614613daa575b6001613da0613d958360ff94613784565b9460011b61fffe1690565b9401169291613d5c565b946001613da0613d95613dbf8960ff95613784565b989350505050613d84565b505090565b637fc4ea7d60e11b5f5260045ffd5b613de66136fa565b50805190811580613e57575b15613e13575050604051613e07604082610447565b5f81525f602082015290565b60205f5160206140d85f395f51905f52910151065f5160206140d85f395f51905f52035f5160206140d85f395f51905f528111611e8b57604051916138d38361040c565b50602081015115613df2565b609780546001600160a01b039283166001600160a01b0319821681179092559091167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3565b15613eb257565b631019106960e31b5f5260045ffd5b90610100825111613f2a57815115613f2557602082015160019060f81c81901b5b8351821015613f2057600190613f0b613f016120546120468689611e58565b60ff600191161b90565b90613f17818311613eab565b17910190613ee2565b925050565b5f9150565b637da54e4760e11b5f5260045ffd5b8054821015611cc9575f5260205f2001905f90565b91613f679183549060031b91821b915f19901b19161790565b9055565b805f5260fc60205260405f2054155f14613fe05760fb5468010000000000000000811015610427576001810160fb5560fb54811015611cc9577f3e7257b7272bb46d49cd6019b04ddee20da7c0cb13f7c1ec3391291b2ccebabc0181905560fb545f91825260fc602052604090912055600190565b505f90565b8054801561400c575f190190613ffb8282613f39565b8154905f199060031b1b1916905555565b634e487b7160e01b5f52603160045260245ffd5b5f81815260fc60205260409020549081156140bb575f19820190828211611e8b5760fb545f19810193908411611e8b57838361407a945f9603614080575b50505061406b60fb613fe5565b60fc905f5260205260405f2090565b55600190565b61406b6140ac916140a26140986140b29560fb613f39565b90549060031b1c90565b92839160fb613f39565b90613f4e565b555f808061405e565b50505f90565b156140c857565b63d51edae360e01b5f5260045ffdfe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a26469706673582212201c679e27d3765a30b5320664118022104964d2150d3e23e7dc75e9d4ff7bbcae64736f6c634300081b0033",
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
