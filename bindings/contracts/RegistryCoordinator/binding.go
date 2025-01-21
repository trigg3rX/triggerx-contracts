// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractRegistryCoordinator

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

// IBLSApkRegistryPubkeyRegistrationParams is an auto generated low-level Go binding around an user-defined struct.
type IBLSApkRegistryPubkeyRegistrationParams struct {
	PubkeyRegistrationSignature BN254G1Point
	PubkeyG1                    BN254G1Point
	PubkeyG2                    BN254G2Point
}

// IRegistryCoordinatorOperatorInfo is an auto generated low-level Go binding around an user-defined struct.
type IRegistryCoordinatorOperatorInfo struct {
	OperatorId [32]byte
	Status     uint8
}

// IRegistryCoordinatorOperatorKickParam is an auto generated low-level Go binding around an user-defined struct.
type IRegistryCoordinatorOperatorKickParam struct {
	QuorumNumber uint8
	Operator     common.Address
}

// IRegistryCoordinatorOperatorSetParam is an auto generated low-level Go binding around an user-defined struct.
type IRegistryCoordinatorOperatorSetParam struct {
	MaxOperatorCount        uint32
	KickBIPsOfOperatorStake uint16
	KickBIPsOfTotalStake    uint16
}

// IRegistryCoordinatorQuorumBitmapUpdate is an auto generated low-level Go binding around an user-defined struct.
type IRegistryCoordinatorQuorumBitmapUpdate struct {
	UpdateBlockNumber     uint32
	NextUpdateBlockNumber uint32
	QuorumBitmap          *big.Int
}

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// IStakeRegistryStrategyParams is an auto generated low-level Go binding around an user-defined struct.
type IStakeRegistryStrategyParams struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// ContractRegistryCoordinatorMetaData contains all meta data concerning the ContractRegistryCoordinator contract.
var ContractRegistryCoordinatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIServiceManager\",\"name\":\"_serviceManager\",\"type\":\"address\"},{\"internalType\":\"contractIStakeRegistry\",\"name\":\"_stakeRegistry\",\"type\":\"address\"},{\"internalType\":\"contractIBLSApkRegistry\",\"name\":\"_blsApkRegistry\",\"type\":\"address\"},{\"internalType\":\"contractIIndexRegistry\",\"name\":\"_indexRegistry\",\"type\":\"address\"},{\"internalType\":\"contractIAllocationManager\",\"name\":\"_allocationManager\",\"type\":\"address\"},{\"internalType\":\"contractIPauserRegistry\",\"name\":\"_pauserRegistry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyRegisteredForQuorums\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BitmapCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BitmapEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BitmapUpdateIsAfterBlockNumber\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BitmapValueTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BytesArrayLengthTooLong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BytesArrayNotOrdered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotChurnSelf\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotKickOperatorAboveThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotReregisterYet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChurnApproverSaltUsed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CurrentlyPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpModFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InputAddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InputLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientStakeForChurn\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNewPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxQuorumsReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NextBitmapUpdateIsBeforeBlockNumber\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotRegisteredForQuorum\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotSorted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAllocationManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEjector\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyUnpauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorSetsEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorSetsNotEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorSetsNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QuorumDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QuorumOperatorCountMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryCoordinatorSignatureExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SaltAlreadyUsed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignatureExpired\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prevChurnApprover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newChurnApprover\",\"type\":\"address\"}],\"name\":\"ChurnApproverUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prevEjector\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newEjector\",\"type\":\"address\"}],\"name\":\"EjectorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"}],\"name\":\"OperatorDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"}],\"name\":\"OperatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"maxOperatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structIRegistryCoordinator.OperatorSetParam\",\"name\":\"operatorSetParams\",\"type\":\"tuple\"}],\"name\":\"OperatorSetParamsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"socket\",\"type\":\"string\"}],\"name\":\"OperatorSocketUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blocknumber\",\"type\":\"uint256\"}],\"name\":\"QuorumBlockNumberUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"OPERATOR_CHURN_APPROVAL_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PUBKEY_REGISTRATION_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allocationManager\",\"outputs\":[{\"internalType\":\"contractIAllocationManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blsApkRegistry\",\"outputs\":[{\"internalType\":\"contractIBLSApkRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registeringOperator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"registeringOperatorId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"internalType\":\"structIRegistryCoordinator.OperatorKickParam[]\",\"name\":\"operatorKickParams\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"name\":\"calculateOperatorChurnApprovalDigestHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"churnApprover\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"maxOperatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\"}],\"internalType\":\"structIRegistryCoordinator.OperatorSetParam\",\"name\":\"operatorSetParams\",\"type\":\"tuple\"},{\"internalType\":\"uint96\",\"name\":\"minimumStake\",\"type\":\"uint96\"},{\"components\":[{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"multiplier\",\"type\":\"uint96\"}],\"internalType\":\"structIStakeRegistry.StrategyParams[]\",\"name\":\"strategyParams\",\"type\":\"tuple[]\"},{\"internalType\":\"uint32\",\"name\":\"lookAheadPeriod\",\"type\":\"uint32\"}],\"name\":\"createSlashableStakeQuorum\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"maxOperatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\"}],\"internalType\":\"structIRegistryCoordinator.OperatorSetParam\",\"name\":\"operatorSetParams\",\"type\":\"tuple\"},{\"internalType\":\"uint96\",\"name\":\"minimumStake\",\"type\":\"uint96\"},{\"components\":[{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"multiplier\",\"type\":\"uint96\"}],\"internalType\":\"structIStakeRegistry.StrategyParams[]\",\"name\":\"strategyParams\",\"type\":\"tuple[]\"}],\"name\":\"createTotalDelegatedStakeQuorum\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint32[]\",\"name\":\"operatorSetIds\",\"type\":\"uint32[]\"}],\"name\":\"deregisterOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"}],\"name\":\"deregisterOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"}],\"name\":\"ejectOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ejectionCooldown\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ejector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableOperatorSets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"}],\"name\":\"getCurrentQuorumBitmap\",\"outputs\":[{\"internalType\":\"uint192\",\"name\":\"\",\"type\":\"uint192\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getOperator\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"internalType\":\"enumIRegistryCoordinator.OperatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structIRegistryCoordinator.OperatorInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"}],\"name\":\"getOperatorFromId\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getOperatorId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"}],\"name\":\"getOperatorSetParams\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"maxOperatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\"}],\"internalType\":\"structIRegistryCoordinator.OperatorSetParam\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getOperatorStatus\",\"outputs\":[{\"internalType\":\"enumIRegistryCoordinator.OperatorStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getQuorumBitmapAtBlockNumberByIndex\",\"outputs\":[{\"internalType\":\"uint192\",\"name\":\"\",\"type\":\"uint192\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"}],\"name\":\"getQuorumBitmapHistoryLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32[]\",\"name\":\"operatorIds\",\"type\":\"bytes32[]\"}],\"name\":\"getQuorumBitmapIndicesAtBlockNumber\",\"outputs\":[{\"internalType\":\"uint32[]\",\"name\":\"\",\"type\":\"uint32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getQuorumBitmapUpdateByIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"updateBlockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"nextUpdateBlockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint192\",\"name\":\"quorumBitmap\",\"type\":\"uint192\"}],\"internalType\":\"structIRegistryCoordinator.QuorumBitmapUpdate\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"indexRegistry\",\"outputs\":[{\"internalType\":\"contractIIndexRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_initialOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_churnApprover\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ejector\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_initialPausedStatus\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"maxOperatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\"}],\"internalType\":\"structIRegistryCoordinator.OperatorSetParam[]\",\"name\":\"_operatorSetParams\",\"type\":\"tuple[]\"},{\"internalType\":\"uint96[]\",\"name\":\"_minimumStakes\",\"type\":\"uint96[]\"},{\"components\":[{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"multiplier\",\"type\":\"uint96\"}],\"internalType\":\"structIStakeRegistry.StrategyParams[][]\",\"name\":\"_strategyParams\",\"type\":\"tuple[][]\"},{\"internalType\":\"enumStakeType[]\",\"name\":\"_stakeTypes\",\"type\":\"uint8[]\"},{\"internalType\":\"uint32[]\",\"name\":\"_lookAheadPeriods\",\"type\":\"uint32[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"isChurnApproverSaltUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"isM2Quorum\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOperatorSetAVS\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isUsingOperatorSets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastEjectionTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numRegistries\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"}],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauserRegistry\",\"outputs\":[{\"internalType\":\"contractIPauserRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"pubkeyRegistrationMessageHash\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quorumCount\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"quorumUpdateBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"socket\",\"type\":\"string\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyG1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"pubkeyG2\",\"type\":\"tuple\"}],\"internalType\":\"structIBLSApkRegistry.PubkeyRegistrationParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"name\":\"operatorSignature\",\"type\":\"tuple\"}],\"name\":\"registerOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint32[]\",\"name\":\"operatorSetIds\",\"type\":\"uint32[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"registerOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"socket\",\"type\":\"string\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyG1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"pubkeyG2\",\"type\":\"tuple\"}],\"internalType\":\"structIBLSApkRegistry.PubkeyRegistrationParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"internalType\":\"structIRegistryCoordinator.OperatorKickParam[]\",\"name\":\"operatorKickParams\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"name\":\"churnApproverSignature\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"name\":\"operatorSignature\",\"type\":\"tuple\"}],\"name\":\"registerOperatorWithChurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"registries\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"serviceManager\",\"outputs\":[{\"internalType\":\"contractIServiceManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_churnApprover\",\"type\":\"address\"}],\"name\":\"setChurnApprover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ejectionCooldown\",\"type\":\"uint256\"}],\"name\":\"setEjectionCooldown\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ejector\",\"type\":\"address\"}],\"name\":\"setEjector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"maxOperatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\"}],\"internalType\":\"structIRegistryCoordinator.OperatorSetParam\",\"name\":\"operatorSetParams\",\"type\":\"tuple\"}],\"name\":\"setOperatorSetParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeRegistry\",\"outputs\":[{\"internalType\":\"contractIStakeRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"operators\",\"type\":\"address[]\"}],\"name\":\"updateOperators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[][]\",\"name\":\"operatorsPerQuorum\",\"type\":\"address[][]\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"}],\"name\":\"updateOperatorsForQuorum\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"socket\",\"type\":\"string\"}],\"name\":\"updateSocket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526004361015610011575f80fd5b5f3560e01c8062cf2ab51461039e57806303fd34921461039957806304ec635114610394578063054310e61461038f5780630cf4b7671461038a5780630d3f213414610385578063125e05841461038057806313542a4e1461037b578063136439dd146103765780631478851f146103715780631eb812da1461036c578063249a0c421461036757806328f61b3114610362578063296bb0641461035d57806329d1e0c3146103585780632cdd1e86146103535780633998fdd31461034e5780633c2a7f4c146103495780633eef3a51146103445780635140a5481461033f5780635865c60c1461033a578063595c6a67146103355780635ac86ab7146103305780635b0b829f1461032b5780635c975abb146103265780635df45946146103215780636347c9001461031c57806368304835146103175780636e3b17db14610312578063715018a61461030d5780637fc3f886146103085780638281ab751461030357806384ca5213146102fe578063871ef049146102f9578063886f1195146102f45780638da5cb5b146102ef5780639aa1653d146102ea5780639b5d177b146102e55780639d8e0c23146102e05780639e9923c2146102db5780639feab859146102d6578063a4d7871f146102d1578063a50857bf146102cc578063a96f783e146102c7578063adcf73f7146102c2578063bd33ee24146102a9578063c391425e146102bd578063ca0de882146102b8578063ca4f2d97146102b3578063ca8aa7c7146102ae578063cabbb17f146102a9578063d72d8dd6146102a4578063e65797ad1461029f578063ee3188211461029a578063f2fde38b14610295578063fabc1cbc146102905763fd39105a1461028b575f80fd5b6127ad565b6126cb565b61263a565b6125c5565b612531565b612514565b6122ea565b6124d0565b612429565b6123ef565b61234b565b6121c9565b6121ac565b612099565b612062565b612028565b611fe4565b611ecc565b611e06565b611cb2565b611c8a565b611c46565b611c16565b611bba565b61190c565b611835565b611585565b61147e565b61143a565b6113e0565b61136e565b611351565b6112bc565b61128d565b61121a565b6111ae565b6110bd565b610cfe565b610bbd565b610b79565b610b4c565b610b1f565b610a6c565b610a44565b610a12565b61098a565b61095b565b61089d565b610862565b610827565b610806565b610766565b6106cf565b6105ef565b6105b7565b6104ed565b634e487b7160e01b5f52604160045260245ffd5b604081019081106001600160401b038211176103d257604052565b6103a3565b606081019081106001600160401b038211176103d257604052565b90601f801991011681019081106001600160401b038211176103d257604052565b604051906104226040836103f2565b565b604051906104226060836103f2565b6001600160401b0381116103d25760051b60200190565b6001600160a01b0381160361045b57565b5f80fd5b600435906104228261044a565b602435906104228261044a565b604435906104228261044a565b9080601f8301121561045b57813561049d81610433565b926104ab60405194856103f2565b81845260208085019260051b82010192831161045b57602001905b8282106104d35750505090565b6020809183356104e28161044a565b8152019101906104c6565b3461045b57602036600319011261045b576004356001600160401b03811161045b5761051d903690600401610486565b61053461052e600480600154161490565b156127f3565b5f5b81518110156105b5576001906105af6001600160a01b036105578386612816565b5116805f52609960205260405f2061058860ff8660405193610578856103b7565b805485520154166020830161282a565b6105a96105a46105988351614e2d565b6001600160c01b031690565b61344b565b9161352d565b01610536565b005b3461045b57602036600319011261045b576004355f526098602052602060405f2054604051908152f35b63ffffffff81160361045b57565b3461045b57606036600319011261045b5760243561062f610629600435610615846105e1565b604435905f52609860205260405f206113c6565b50612914565b63ffffffff8082511692169182106106b65760408161066b6106939460206106799501519063ffffffff821615918215610697575b5050613603565b01516001600160c01b031690565b6040516001600160c01b0390911681529081906020820190565b0390f35b9091506106ae9063ffffffff165b63ffffffff1690565b115f80610664565b636cb19aff60e01b5f5260045ffd5b5f91031261045b57565b3461045b575f36600319011261045b57609d546040516001600160a01b039091168152602090f35b6001600160401b0381116103d257601f01601f191660200190565b92919261071e826106f7565b9161072c60405193846103f2565b82948184528183011161045b578281602093845f960137010152565b9080601f8301121561045b5781602061076393359101610712565b90565b3461045b57602036600319011261045b576004356001600160401b03811161045b57610796903690600401610748565b335f52609960205260ff600160405f2001541660038110156108015760016107be911461285c565b335f5260996020527fec2963ab21c1e50e1e582aa542af2e4bf7bf38e6e1403c27b42e1c5d6e621eaa6107fc60405f20549260405191829182612896565b0390a2005b611183565b3461045b57602036600319011261045b57600435610822613619565b60a055005b3461045b57602036600319011261045b576004356108448161044a565b60018060a01b03165f52609f602052602060405f2054604051908152f35b3461045b57602036600319011261045b5760043561087f8161044a565b60018060a01b03165f526099602052602060405f2054604051908152f35b3461045b57602036600319011261045b5760043560405163237dfb4760e11b8152336004820152906020826024817f0000000000000000000000003db5b2115c2c4619255c107c8daade9dd35b366c6001600160a01b03165afa918215610956576105b592610913915f91610927575b506128ca565b610922600154828116146128e0565b6136a5565b610949915060203d60201161094f575b61094181836103f2565b8101906128a7565b5f61090d565b503d610937565b6128bf565b3461045b57602036600319011261045b576004355f52609a602052602060ff60405f2054166040519015158152f35b3461045b57604036600319011261045b5760606109c26106296024356004356109b16128f6565b505f52609860205260405f206113c6565b6040519063ffffffff815116825263ffffffff6020820151166020830152604060018060c01b03910151166040820152f35b6004359060ff8216820361045b57565b359060ff8216820361045b57565b3461045b57602036600319011261045b5760ff610a2d6109f4565b165f52609b602052602060405f2054604051908152f35b3461045b575f36600319011261045b57609e546040516001600160a01b039091168152602090f35b3461045b57602036600319011261045b576040516308f6629d60e31b815260048035908201526020816024816001600160a01b037f000000000000000000000000c275cce555cc4bd954ff078621778368f9006f1c165afa801561095657610693915f91610af0575b506040516001600160a01b0390911681529081906020820190565b610b12915060203d602011610b18575b610b0a81836103f2565b810190612945565b5f610ad5565b503d610b00565b3461045b57602036600319011261045b576105b5600435610b3f8161044a565b610b47613619565b6136d7565b3461045b57602036600319011261045b576105b5600435610b6c8161044a565b610b74613619565b613735565b3461045b575f36600319011261045b576040517f000000000000000000000000c0aa79c13bc1eea0168f1a2975b7aa9ff64135ed6001600160a01b03168152602090f35b3461045b57602036600319011261045b576040610be4600435610bdf8161044a565b612972565b610bfa8251809260208091805184520151910152565bf35b359061ffff8216820361045b57565b606090600319011261045b5760405190610c24826103d7565b81600435610c31816105e1565b815260243561ffff8116810361045b5760208201526044359061ffff8216820361045b5760400152565b6001600160601b0381160361045b57565b81601f8201121561045b57803590610c8382610433565b92610c9160405194856103f2565b82845260208085019360061b8301019181831161045b57602001925b828410610cbb575050505090565b60408483031261045b5760206040918251610cd5816103b7565b8635610ce08161044a565b815282870135610cef81610c5b565b83820152815201930192610cad565b3461045b5760c036600319011261045b57610d1836610c0b565b606435610d2481610c5b565b6084356001600160401b03811161045b57610d43903690600401610c6c565b9060a43591610d51836105e1565b610d59613619565b610d6760ff60a154166129d8565b60965460ff16938490610da290610d8060c084106132d7565b610d9c610d8c886138fe565b60ff1660ff196096541617609655565b86613eaa565b60a15460ff1680611064575b610f2c575b50610dbe6001612e96565b610dc86001612e96565b7f000000000000000000000000d7c48ab86ab9390e09165013600d5e8721cb3e1d6001600160a01b0316803b1561045b57610e1e935f809460405196879586948593630662d3e160e51b85528b60048601613a91565b03925af1801561095657610f18575b507f00000000000000000000000090103b1c64ceeb183db98c7496af617c1585e73f6001600160a01b0316803b1561045b5760405163136ca0f960e11b815260ff83166004820152905f908290602490829084905af1801561095657610f04575b507f000000000000000000000000c275cce555cc4bd954ff078621778368f9006f1c6001600160a01b0316803b1561045b5760405163136ca0f960e11b815260ff83166004820152905f908290602490829084905af1801561095657610ef057005b80610efe5f6105b5936103f2565b806106c5565b80610efe5f610f12936103f2565b5f610e8e565b80610efe5f610f26936103f2565b5f610e2d565b92610f35613910565b92610f40835161395d565b935f5b8451811015610f8c5780610f86610f6d610f5f60019489612816565b51516001600160a01b031690565b610f77838a612816565b6001600160a01b039091169052565b01610f43565b5091949093610fa8610f9c610413565b63ffffffff9093168352565b6020820152610fb682612809565b52610fc081612809565b507f00000000000000000000000078469728304326cbc65f8f95fa756b0b731644626001600160a01b0316803b1561045b57604051630130fc2760e51b8152915f918391829084908290611041907f000000000000000000000000c0aa79c13bc1eea0168f1a2975b7aa9ff64135ed6001600160a01b031660048401613985565b03925af180156109565715610db35780610efe5f61105e936103f2565b5f610db3565b5061108b6110876110808760ff165f5260a260205260405f2090565b5460ff1690565b1590565b610dae565b9181601f8401121561045b578235916001600160401b03831161045b576020838186019501011161045b57565b3461045b57604036600319011261045b576004356001600160401b03811161045b573660238201121561045b5780600401356110f881610433565b9161110660405193846103f2565b8183526024602084019260051b8201019036821161045b5760248101925b82841061115457602435856001600160401b03821161045b5761114e6105b5923690600401611090565b916129ee565b83356001600160401b03811161045b57602091611178839260243691870101610486565b815201930192611124565b634e487b7160e01b5f52602160045260245ffd5b6003111561080157565b9060038210156108015752565b3461045b57602036600319011261045b576004356111cb8161044a565b6111d361295a565b5060018060a01b03165f52609960205260405f206111fb60ff600160405193610578856103b7565b60405180916106936020604084019280518552015160208401906111a1565b3461045b575f36600319011261045b5760405163237dfb4760e11b81523360048201526020816024817f0000000000000000000000003db5b2115c2c4619255c107c8daade9dd35b366c6001600160a01b03165afa801561095657611285915f9161092757506128ca565b6105b5613671565b3461045b57602036600319011261045b576020600160ff6112ac6109f4565b161b806001541614604051908152f35b3461045b57608036600319011261045b576112d56109f4565b606036602319011261045b576040516112ed816103d7565b6024356112f9816105e1565b815260443561ffff8116810361045b57602082015260643561ffff8116810361045b576040820152611329613619565b60ff6096541660ff83161015611342576105b591613eaa565b637310cff560e11b5f5260045ffd5b3461045b575f36600319011261045b576020600154604051908152f35b3461045b575f36600319011261045b576040517f000000000000000000000000c275cce555cc4bd954ff078621778368f9006f1c6001600160a01b03168152602090f35b634e487b7160e01b5f52603260045260245ffd5b80548210156113db575f5260205f2001905f90565b6113b2565b3461045b57602036600319011261045b57600435609c5481101561045b57609c5f527faf85b9071dfafeac1409d3f1d19bafc9bc7c37974cde8df0ee6168f0086e539c01546040516001600160a01b039091168152602090f35b3461045b575f36600319011261045b576040517f000000000000000000000000d7c48ab86ab9390e09165013600d5e8721cb3e1d6001600160a01b03168152602090f35b3461045b57604036600319011261045b5760043561149b8161044a565b6024356001600160401b03811161045b576114ba903690600401610748565b609e546001600160a01b03163303611576576001600160a01b0382165f908152609f6020908152604080832042905560999091529020805460016115228161151961151361059861150d60965460ff1690565b89613e7f565b94614e2d565b94015460ff1690565b61152b81611197565b149182611563575b8261154a575b505061154157005b6105b591613fbc565b81166001600160c01b0390811691161490505f80611539565b6001600160c01b03821615159250611533565b6376d8ab1760e11b5f5260045ffd5b3461045b575f36600319011261045b5761159d613619565b606480546001600160a01b031981169091555f906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b81601f8201121561045b578035906115f782610433565b9261160560405194856103f2565b8284526020606081860194028301019181831161045b57602001925b82841061162f575050505090565b60608483031261045b57602060609160405161164a816103d7565b8635611655816105e1565b8152611662838801610bfc565b8382015261167260408801610bfc565b6040820152815201930192611621565b9080601f8301121561045b57813561169981610433565b926116a760405194856103f2565b81845260208085019260051b82010192831161045b57602001905b8282106116cf5750505090565b6020809183356116de81610c5b565b8152019101906116c2565b9080601f8301121561045b57813561170081610433565b9261170e60405194856103f2565b81845260208085019260051b8201019183831161045b5760208201905b83821061173a57505050505090565b81356001600160401b03811161045b5760209161175c87848094880101610c6c565b81520191019061172b565b9080601f8301121561045b5781359061177f82610433565b9261178d60405194856103f2565b82845260208085019360051b82010191821161045b57602001915b8183106117b55750505090565b8235600281101561045b578152602092830192016117a8565b9080601f8301121561045b5781356117e581610433565b926117f360405194856103f2565b81845260208085019260051b82010192831161045b57602001905b82821061181b5750505090565b60208091833561182a816105e1565b81520191019061180e565b3461045b5761012036600319011261045b5761184f61045f565b61185761046c565b90611860610479565b6064356084356001600160401b03811161045b576118829036906004016115e0565b60a4356001600160401b03811161045b576118a1903690600401611682565b9060c4356001600160401b03811161045b576118c19036906004016116e9565b9260e4356001600160401b03811161045b576118e1903690600401611767565b9461010435976001600160401b03891161045b576119066105b59936906004016117ce565b97612d5f565b3461045b5760a036600319011261045b5761192636610c0b565b60643561193281610c5b565b6084356001600160401b03811161045b57611951903690600401610c6c565b9061195a613619565b60965460ff169283906119859061197360c084106132d7565b61197f610d8c876138fe565b85613eaa565b60a15460ff1680611b09575b6119f6575b506119a05f612e96565b7f000000000000000000000000d7c48ab86ab9390e09165013600d5e8721cb3e1d6001600160a01b031691823b1561045b57610e1e925f9283604051809681958294633aea0b9d60e11b84528a60048501613ac6565b919092611a01613910565b93611a0c835161395d565b945f5b8451811015611a3b5780611a35611a2b610f5f60019489612816565b610f77838b612816565b01611a0f565b50919493909293611a4d610f9c610413565b6020820152611a5b82612809565b52611a6581612809565b507f00000000000000000000000078469728304326cbc65f8f95fa756b0b731644626001600160a01b0316803b1561045b57604051630130fc2760e51b8152915f918391829084908290611ae6907f000000000000000000000000c0aa79c13bc1eea0168f1a2975b7aa9ff64135ed6001600160a01b031660048401613985565b03925af1801561095657156119965780610efe5f611b03936103f2565b5f611996565b50611b256110876110808660ff165f5260a260205260405f2090565b611991565b81601f8201121561045b57803590611b4182610433565b92611b4f60405194856103f2565b82845260208085019360061b8301019181831161045b57602001925b828410611b79575050505090565b60408483031261045b5760206040918251611b93816103b7565b611b9c87610a04565b815282870135611bab8161044a565b83820152815201930192611b6b565b3461045b5760a036600319011261045b57600435611bd78161044a565b60243590604435906001600160401b03821161045b57602092611c01611c0e933690600401611b2a565b6064359160843593613019565b604051908152f35b3461045b57602036600319011261045b576020611c34600435614e2d565b6040516001600160c01b039091168152f35b3461045b575f36600319011261045b576040517f0000000000000000000000003db5b2115c2c4619255c107c8daade9dd35b366c6001600160a01b03168152602090f35b3461045b575f36600319011261045b576064546040516001600160a01b039091168152602090f35b3461045b575f36600319011261045b57602060ff60965416604051908152f35b919082604091031261045b57604051611cea816103b7565b6020808294803584520135910152565b9080601f8301121561045b5760405191611d156040846103f2565b82906040810192831161045b57905b828210611d315750505090565b8135815260209182019101611d24565b9061010060431983011261045b5760405191611d5c836103d7565b82611d68826044611cd2565b8152611d75826084611cd2565b6020820152608060c31983011261045b57604090611dad825193611d98856103b7565b611da38160c4611cfa565b8552610104611cfa565b60208401520152565b91909160608184031261045b5760405190611dd0826103d7565b81938135916001600160401b03831161045b57611df36040939284938301610748565b8452602081013560208501520135910152565b3461045b576101a036600319011261045b576004356001600160401b03811161045b57611e37903690600401611090565b906024356001600160401b03811161045b57611e57903690600401610748565b611e6036611d41565b610144356001600160401b03811161045b57611e80903690600401611b2a565b90610164356001600160401b03811161045b57611ea1903690600401611db6565b9261018435956001600160401b03871161045b57611ec66105b5973690600401611db6565b956130d4565b3461045b57604036600319011261045b57600435611ee98161044a565b6024356001600160401b03811161045b57611f089036906004016117ce565b90611f1961052e6001808054161490565b611f2d611f2860ff60a1541690565b6129d8565b5f5b8251811015611f835780611f7d611f78611087611080611f67611f61611f576001988b612816565b5163ffffffff1690565b60ff1690565b60ff165f5260a260205260405f2090565b61327e565b01611f2f565b50611f8c614c0c565b611f968251613294565b5f5b8351811015611fda5780611fc7611fb7611f61611f5760019589612816565b60f81b6001600160f81b03191690565b5f1a611fd382856132c6565b5301611f98565b506105b591613fbc565b3461045b575f36600319011261045b576040517f00000000000000000000000090103b1c64ceeb183db98c7496af617c1585e73f6001600160a01b03168152602090f35b3461045b575f36600319011261045b5760206040517f2bd82124057f0913bc3b772ce7b83e8057c1ad1f3510fc83778be20f10ec5de68152f35b3461045b57602036600319011261045b5760ff61207d6109f4565b165f5260a2602052602060ff60405f2054166040519015158152f35b3461045b5761016036600319011261045b576004356001600160401b03811161045b576120ca903690600401610748565b6024356001600160401b03811161045b576120e9903690600401610748565b6120f236611d41565b61014435916001600160401b03831161045b578361214061211a612146953690600401611db6565b9361212b61052e6001808054161490565b61213a60ff60a1541615613236565b3361436b565b336146eb565b51905f5b81518110156105b557806121a6612163600193856132c6565b5160f81c63ffffffff61219d816121928161217e888c612816565b51169460ff165f52609760205260405f2090565b541663ffffffff1690565b911611156132d7565b0161214a565b3461045b575f36600319011261045b57602060a054604051908152f35b3461045b57606036600319011261045b576004356121e68161044a565b6024356001600160401b03811161045b576122059036906004016117ce565b906044356001600160401b03811161045b57612225903690600401610748565b9061223661052e6001808054161490565b612245611f2860ff60a1541690565b5f5b8351811015612275578061226f611f78611087611080611f67611f61611f576001988c612816565b01612247565b509061229561229e91612286614c0c565b6020808251830101910161335c565b8392919261436b565b906122a98451613294565b915f5b85518110156122de57806122cb611fb7611f61611f576001958b612816565b5f1a6122d782876132c6565b53016122ac565b50916105b59284614c4d565b3461045b575f36600319011261045b57602060ff60a154166040519015158152f35b60206040818301928281528451809452019201905f5b81811061232f5750505090565b825163ffffffff16845260209384019390920191600101612322565b3461045b57604036600319011261045b57600435612368816105e1565b602435906001600160401b03821161045b573660238301121561045b5781600401359161239483610433565b926123a260405194856103f2565b8084526024602085019160051b8301019136831161045b57602401905b8282106123df576106936123d38686614da9565b6040519182918261230c565b81358152602091820191016123bf565b3461045b575f36600319011261045b5760206040517f4d404e3276e7ac2163d8ee476afa6a41d1f68fb71f2d8b6546b24e55ce01b72a8152f35b3461045b57602036600319011261045b576004356001600160401b03811161045b57612459903690600401610748565b61246a61052e600280600154161490565b5f60ff60a15416158015915b83518110156124c6578061248c600192866132c6565b5160f81c83856124a7575b6124a19150613236565b01612476565b505f5260a26020526124a16124c160405f2060ff90541690565b612497565b6105b58433613fbc565b3461045b575f36600319011261045b576040517f00000000000000000000000078469728304326cbc65f8f95fa756b0b731644626001600160a01b03168152602090f35b3461045b575f36600319011261045b576020609c54604051908152f35b3461045b57602036600319011261045b5760ff61254c6109f4565b6125546128f6565b50165f52609760205261069360405f2061ffff60405191612574836103d7565b5463ffffffff81168352818160201c16602084015260301c16604082015260405191829182919091604061ffff81606084019563ffffffff8151168552826020820151166020860152015116910152565b3461045b575f36600319011261045b576125dd613619565b6125ec60ff60a1541615613236565b5f5b60ff6096541660ff8216908110156126285760ff916001915f5260a260205261262160405f20600160ff19825416179055565b01166125ee565b6105b5600160ff1960a154161760a155565b3461045b57602036600319011261045b576004356126578161044a565b61265f613619565b6001600160a01b03811615612677576105b590614272565b60405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608490fd5b3461045b57602036600319011261045b5760043560405163755b36bd60e11b81526020816004817f0000000000000000000000003db5b2115c2c4619255c107c8daade9dd35b366c6001600160a01b03165afa908115610956575f9161278e575b506001600160a01b0316330361277f5761274d6001541982198116146128e0565b806001556040519081527f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c60203392a2005b63794821ff60e01b5f5260045ffd5b6127a7915060203d602011610b1857610b0a81836103f2565b5f61272c565b3461045b57602036600319011261045b576004356127ca8161044a565b60018060a01b03165f526099602052602060ff600160405f20015416610bfa60405180926111a1565b156127fa57565b63840a48d560e01b5f5260045ffd5b8051156113db5760200190565b80518210156113db5760209160051b010190565b60038210156108015752565b90610422604051612846816103b7565b602060ff6001839680548552015416910161282a565b1561286357565b63aba4733960e01b5f5260045ffd5b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b906020610763928181520190612872565b9081602091031261045b5751801515810361045b5790565b6040513d5f823e3d90fd5b156128d157565b631d77d47760e21b5f5260045ffd5b156128e757565b63c61dca5d60e01b5f5260045ffd5b60405190612903826103d7565b5f6040838281528260208201520152565b90604051612921816103d7565b604081935463ffffffff8116835263ffffffff8160201c166020840152811c910152565b9081602091031261045b57516107638161044a565b60405190612967826103b7565b5f6020838281520152565b6129d36107639161298161295a565b50604080517f2bd82124057f0913bc3b772ce7b83e8057c1ad1f3510fc83778be20f10ec5de6602082019081526001600160a01b03909316818301529081526129cb6060826103f2565b519020613793565b6137e0565b156129df57565b635b77901960e01b5f5260045ffd5b909291612a0261052e600480600154161490565b612a21612a1160965460ff1690565b612a1c368488610712565b613e7f565b50612a2e81835114612c37565b7f00000000000000000000000090103b1c64ceeb183db98c7496af617c1585e73f6001600160a01b0316935f5b828110612a6a57505050509050565b612a8f612a89612a7b838686612c4d565b356001600160f81b03191690565b60f81c90565b92612a9a8286612816565b5180516040516379a0849160e11b815260ff87166004820152919791906020826024818d5afa91821561095657612ae09263ffffffff915f91612c09575b501614612c6e565b5f97885b88518a1015612b9d57600190612b95612b0d612b008d8d612816565b516001600160a01b031690565b91612b70612b33612b2e8560018060a01b03165f52609960205260405f2090565b612836565b91612b5b612b568d612b486105988751614e2d565b60ff600192161c1660011490565b612c84565b858060a01b0316858060a01b03851611612c9a565b612b8e612b87612b7f8a612cc4565b8a8a8d612ce4565b3691610712565b908361352d565b990198612ae4565b5096509650929060019194929443612bc08260ff165f52609b60205260405f2090565b557f46077d55330763f16269fd75e5761663f4192d2791747c0189b16ad31db07db460ff60405192169180612bfa43829190602083019252565b0390a201949394929092612a5b565b612c2a915060203d8111612c30575b612c2281836103f2565b810190612c59565b5f612ad8565b503d612c18565b15612c3e57565b63aaad13f760e01b5f5260045ffd5b908210156113db570190565b9081602091031261045b5751610763816105e1565b15612c7557565b638e5aeee760e01b5f5260045ffd5b15612c8b57565b63d053aa2160e01b5f5260045ffd5b15612ca157565b63ba50f91160e01b5f5260045ffd5b634e487b7160e01b5f52601160045260245ffd5b9060018201809211612cd257565b612cb0565b91908201809211612cd257565b9093929384831161045b57841161045b578101920390565b15612d0357565b60405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608490fd5b97959391612da9979593915f5499612d8f60ff8c60081c16151515809c81612e23575b8115612e03575b50612cfc565b8a612da0600160ff195f5416175f55565b612dec57612ead565b612daf57565b612dbd61ff00195f54165f55565b604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602090a1565b612dfe61010061ff00195f5416175f55565b612ead565b303b15915081612e15575b505f612d89565b60ff1660011490505f612e0e565b600160ff8216109150612d82565b609c54600160401b8110156103d25760018101609c55609c548110156113db57609c5f527faf85b9071dfafeac1409d3f1d19bafc9bc7c37974cde8df0ee6168f0086e539c0180546001600160a01b0319166001600160a01b03909216919091179055565b6002111561080157565b5160028110156108015790565b92610922610b7492610b47612eec969c9b9a99989c8d89519051809114908161300e575b5080613003575b80612ff8575b612ee790612c37565b614272565b612f1e7f000000000000000000000000d7c48ab86ab9390e09165013600d5e8721cb3e1d6001600160a01b0316612e31565b612f507f000000000000000000000000c275cce555cc4bd954ff078621778368f9006f1c6001600160a01b0316612e31565b612f827f00000000000000000000000090103b1c64ceeb183db98c7496af617c1585e73f6001600160a01b0316612e31565b5f5b8151811015612fef5780612fe9612f9d60019385612816565b51612fb8612fab848c612816565b516001600160601b031690565b612fc28488612816565b51612fd5612fd0868b612816565b612ea0565b91612fe3611f57878d612816565b93613aee565b01612f84565b50505050509050565b508a518c5114612ede565b5089518b5114612ed8565b90508a51145f612ed1565b919493909260405192602084019460e08501917f4d404e3276e7ac2163d8ee476afa6a41d1f68fb71f2d8b6546b24e55ce01b72a875260018060a01b03166040860152606085015260c060808501528651809152602061010085019701905f5b8181106130a65750505061076394956129cb92849260a084015260c083015203601f1981018352826103f2565b8251805160ff168a526020908101516001600160a01b0316818b015260409099019890920191600101613079565b946131206131176131329598949396986130f461052e6001808054161490565b61310b61310661108760ff60a1541690565b613236565b61213a888b5114612c37565b9188833361446d565b61312b368688610712565b90336146eb565b925f5b828110613143575050505050565b8061317361316e61315d612a89612a7b600196898b612c4d565b60ff165f52609760205260405f2090565b61324c565b613181611f57838951612816565b63ffffffff6131976106a5845163ffffffff1690565b9116116131a6575b5001613135565b6131f2906131bb612a89612a7b85898b612c4d565b6131cc612fab8560408c0151612816565b906131de612fab8660208d0151612816565b906131e98689612816565b51923391614ab6565b61323061321360206132048487612816565b5101516001600160a01b031690565b61322a612b8761322285612cc4565b85898b612ce4565b90613fbc565b5f61319f565b1561323d57565b630b88306f60e01b5f5260045ffd5b90604051613259816103d7565b604061ffff82945463ffffffff81168452818160201c16602085015260301c16910152565b1561328557565b63fd2c1f4d60e01b5f5260045ffd5b9061329e826106f7565b6132ab60405191826103f2565b82815280926132bc601f19916106f7565b0190602036910137565b9081518110156113db570160200190565b156132de57565b633cb89c9760e01b5f5260045ffd5b919082604091031261045b57604051613305816103b7565b6020808294805184520151910152565b9080601f8301121561045b57604051916133306040846103f2565b82906040810192831161045b57905b82821061334c5750505090565b815181526020918201910161333f565b91909180830390610120821261045b5780516001600160401b03811161045b57810184601f8201121561045b578051613394816106f7565b916133a260405193846103f2565b818352866020838301011161045b57815f9260208093018386015e8301015293610100601f1984011261045b576080604051936133de856103d7565b6133eb83602086016132ed565b85526133fa83606086016132ed565b6020860152609f19011261045b576134309060e06040519361341b856103b7565b6134288360a08301613315565b855201613315565b6020820152604082015290565b5f198114612cd25760010190565b5f81805b6134c557506134619061ffff16613294565b5f5f5b82518210806134ba575b156134b3576001811b841661348c575b6134879061343d565b613464565b9060016134879160ff60f81b8460f81b165f1a6134a982876132c6565b530191905061347e565b5050905090565b50610100811061346e565b5f198101818111612cd25761ffff9116911661ffff8114612cd257600101908061344f565b9081602091031261045b57516001600160c01b038116810361045b5790565b610763939260609260018060a01b0316825260208201528160408201520190612872565b91906001602082015161353f81611197565b61354881611197565b036135fe57516040516333567f7f60e11b8152916020918391829161357291908760048501613509565b03815f7f000000000000000000000000d7c48ab86ab9390e09165013600d5e8721cb3e1d6001600160a01b03165af1908115610956575f916135cf575b506001600160c01b031690816135c3575050565b61322a6104229261344b565b6135f1915060203d6020116135f7575b6135e981836103f2565b8101906134ea565b5f6135af565b503d6135df565b505050565b1561360a57565b63bbba60cb60e01b5f5260045ffd5b6064546001600160a01b0316330361362d57565b606460405162461bcd60e51b815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b5f196001556040515f1981527fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d60203392a2565b806001556040519081527fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d60203392a2565b609d54604080516001600160a01b038084168252841660208201529192917f315457d8a8fe60f04af17c16e2f5a5e1db612b31648e58030360759ef8f3528c9190a16001600160a01b03166001600160a01b03199190911617609d55565b609e54604080516001600160a01b038084168252841660208201529192917f8f30ab09f43a6c157d7fce7e0a13c003042c1c95e8a72e7a146a21c0caa24dc99190a16001600160a01b03166001600160a01b03199190911617609e55565b61379b614e71565b9060405190602082019261190160f01b845260228301526042820152604281526137c66062826103f2565b51902090565b634e487b7160e01b5f52601260045260245ffd5b5f51602061558c5f395f51905f52906137f761295a565b505f919006602060c0835b6138f7575f935f51602061558c5f395f51905f526003818681818009090860405161382d85826103f2565b8436823784818560405161384182826103f2565b813682378381528360208201528360408201528560608201527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808201525f51602061558c5f395f51905f5260a082015260056107cf195a01fa80156138fc576138ab906152cc565b51916138f7575f51602061558c5f395f51905f52828009146138e257505f51602061558c5f395f51905f5260015f94089293613802565b929350506138ee610413565b92835282015290565b6137cc565bfe5b60ff60019116019060ff8211612cd257565b6040805190919061392183826103f2565b6001815291601f1901825f5b82811061393957505050565b602090604051613948816103b7565b5f81526060838201528282850101520161392d565b9061396782610433565b61397460405191826103f2565b82815280926132bc601f1991610433565b90604082019060018060a01b031682526040602083015282518091526060820191602060608360051b8301019401925f915b8383106139c657505050505090565b9091929394605f1982820301835285516020606081604085019363ffffffff81511686520151936040838201528451809452019201905f905b808210613a1e57505050602080600192970193019301919392906139b7565b82516001600160a01b03168452602093840193909201916001909101906139ff565b90602080835192838152019201905f5b818110613a5d5750505090565b825180516001600160a01b031685526020908101516001600160601b03168186015260409094019390920191600101613a50565b9061076394936001600160601b0360809460ff63ffffffff941685521660208401521660408201528160608201520190613a40565b6001600160601b03610763949360ff6060941683521660208201528160408201520190613a40565b93909192613afe60965460ff1690565b94613b2560ff871691613b1360c084106132d7565b613b1f610d8c896138fe565b87613eaa565b60a15460ff1680613e5e575b613d47575b50613b4081612e96565b80613caa5750507f000000000000000000000000d7c48ab86ab9390e09165013600d5e8721cb3e1d6001600160a01b031691823b1561045b57613b9d925f9283604051809681958294633aea0b9d60e11b84528a60048501613ac6565b03925af1801561095657613c96575b505b7f00000000000000000000000090103b1c64ceeb183db98c7496af617c1585e73f6001600160a01b0316803b1561045b5760405163136ca0f960e11b815260ff83166004820152905f908290602490829084905af1801561095657613c82575b507f000000000000000000000000c275cce555cc4bd954ff078621778368f9006f1c6001600160a01b0316803b1561045b5760405163136ca0f960e11b815260ff90921660048301525f908290818381602481015b03925af1801561095657613c745750565b80610efe5f610422936103f2565b80610efe5f613c90936103f2565b5f613c0e565b80610efe5f613ca4936103f2565b5f613bac565b80613cb9600192959395612e96565b14613cc7575b505050613bae565b7f000000000000000000000000d7c48ab86ab9390e09165013600d5e8721cb3e1d6001600160a01b0316803b1561045b57613d1d935f809460405196879586948593630662d3e160e51b85528b60048601613a91565b03925af1801561095657613d33575b8080613cbf565b80610efe5f613d41936103f2565b5f613d2c565b9592909491613d54613910565b95613d5f865161395d565b965f5b8751811015613d8e5780613d88613d7e610f5f6001948c612816565b610f77838d612816565b01613d62565b509193969790929497613da2610f9c610413565b6020820152613db082612809565b52613dba81612809565b507f00000000000000000000000078469728304326cbc65f8f95fa756b0b731644626001600160a01b0316803b1561045b57604051630130fc2760e51b8152915f918391829084908290613e3b907f000000000000000000000000c0aa79c13bc1eea0168f1a2975b7aa9ff64135ed6001600160a01b031660048401613985565b03925af180156109565715613b365780610efe5f613e58936103f2565b5f613b36565b50613e7a6110876110808860ff165f5260a260205260405f2090565b613b31565b906001613e8d60ff93614f9d565b928392161b1115613e9b5790565b63ca95733360e01b5f5260045ffd5b613f6860ff7f3ee6fe8d54610244c3e9d3c066ae4aee997884aa28f10616ae821925401318ac921692835f52609760205260405f20613f0063ffffffff835116829063ffffffff1663ffffffff19825416179055565b6020820151815465ffff0000000067ffff000000000000604086015160301b169260201b169067ffffffff0000000019161717905560405191829182919091604061ffff81606084019563ffffffff8151168552826020820151166020860152015116910152565b0390a2565b15613f7457565b6368b6a87560e11b5f5260045ffd5b6001600160a01b03909116815260406020820181905261076392910190612872565b604090610763939281528160208201520190612872565b6001600160a01b0381165f90815260996020526040902090600182549201613ff96001613fea835460ff1690565b613ff381611197565b1461285c565b61405961401461059861400e60965460ff1690565b87613e7f565b61401d85614e2d565b6001600160c01b0390911690614034821515613f6d565b61404a8282166001600160c01b03168314612c84565b9019166001600160c01b031690565b6140638185615023565b6001600160c01b0316156141b7575b507f000000000000000000000000c275cce555cc4bd954ff078621778368f9006f1c6001600160a01b0316803b1561045b57835f916140c8938360405180968195829463f4e24fe560e01b845260048401613f83565b03925af18015610956576141a3575b507f000000000000000000000000d7c48ab86ab9390e09165013600d5e8721cb3e1d6001600160a01b0316803b1561045b575f604051809263bd29b8cd60e01b825281838161412a898960048401613fa5565b03925af180156109565761418f575b507f00000000000000000000000090103b1c64ceeb183db98c7496af617c1585e73f6001600160a01b031691823b1561045b57613c63925f928360405180968195829463bd29b8cd60e01b845260048401613fa5565b80610efe5f61419d936103f2565b5f614139565b80610efe5f6141b1936103f2565b5f6140d7565b805460ff191660021790557f000000000000000000000000c0aa79c13bc1eea0168f1a2975b7aa9ff64135ed6001600160a01b0316803b1561045b576040516351b27a6d60e11b81526001600160a01b0383166004820152905f908290602490829084905af180156109565761425e575b50816001600160a01b0382167f396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e45f80a35f614072565b80610efe5f61426c936103f2565b5f614228565b606480546001600160a01b039283166001600160a01b0319821681179092559091167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3565b9081602091031261045b575190565b905f905b600282106142da57505050565b60208060019285518152019301910190916142cd565b6101209061435a60206040610422969897959861016085019960018060a01b0316855261432a838601825160208091805184520151910152565b8083015180516060870152602001516080860152015161434e60a0850182516142c9565b015160e08301906142c9565b019060208091805184520151910152565b6040516309aa152760e11b81526001600160a01b0382811660048301529091907f000000000000000000000000c275cce555cc4bd954ff078621778368f9006f1c16602083602481845afa928315610956575f9361444c575b5082156143d2575050905090565b60209250614402935f6143e484612972565b6040516317ef39cb60e31b81529687958694859391600485016142f0565b03925af1908115610956575f9161441d575b50805f806134b3565b61443f915060203d602011614445575b61443781836103f2565b8101906142ba565b5f614414565b503d61442d565b61446691935060203d6020116144455761443781836103f2565b915f6143c4565b919290602082019283515f52609a60205260ff60405f2054166144e85760408301805142116144d957610422956144d19386515f52609a6020526144bb60405f20600160ff19825416179055565b609d546001600160a01b03169651925193613019565b90519161513d565b630819bdcd60e01b5f5260045ffd5b636fbefec360e11b5f5260045ffd5b60405190614504826103d7565b60606040838281528260208201520152565b1561451d57565b6313ca465760e01b5f5260045ffd5b1561453357565b630c6816cd60e01b5f5260045ffd5b1561454957565b631968677d60e11b5f5260045ffd5b6001602091835181550191015160038110156108015760ff80198354169116179055565b9060018060a01b0316815260406020820152608060406145a7845160608386015260a0850190612872565b9360208101516060850152015191015290565b9080601f8301121561045b5781516145d181610433565b926145df60405194856103f2565b81845260208085019260051b82010192831161045b57602001905b8282106146075750505090565b60208091835161461681610c5b565b8152019101906145fa565b91909160408184031261045b5780516001600160401b03811161045b578361464a9183016145ba565b9260208201516001600160401b03811161045b5761076392016145ba565b60208183031261045b578051906001600160401b03821161045b57019080601f8301121561045b57815161469b81610433565b926146a960405194856103f2565b81845260208085019260051b82010192831161045b57602001905b8282106146d15750505090565b6020809183516146e0816105e1565b8152019101906146c4565b90919293827fec2963ab21c1e50e1e582aa542af2e4bf7bf38e6e1403c27b42e1c5d6e621eaa6147b661471c6144f7565b976147aa61473861059861473260965460ff1690565b8b613e7f565b61474186614e2d565b6001600160c01b0390911690614758821515614516565b60018060c01b031661477261476d8284161590565b61452c565b6001600160a01b0389165f908152609f602052604090206147a39061479c905b5460a05490612cd7565b4211614542565b1785615023565b60405191829182612896565b0390a260016147e1816147d98560018060a01b03165f52609960205260405f2090565b015460ff1690565b6147ea81611197565b0361497e575b507f000000000000000000000000c275cce555cc4bd954ff078621778368f9006f1c6001600160a01b0316803b1561045b575f6040518092631fd93ca960e11b82528183816148438a8960048401613f83565b03925af180156109565784925f92859261496a575b506148776040519687938493632550477760e01b855260048501613509565b0381837f000000000000000000000000d7c48ab86ab9390e09165013600d5e8721cb3e1d6001600160a01b03165af1918215610956576148e1935f9384918591614946575b5060408701526020860152604051938492839262bff04d60e01b845260048401613fa5565b0381837f00000000000000000000000090103b1c64ceeb183db98c7496af617c1585e73f6001600160a01b03165af1908115610956575f91614924575b50815290565b61494091503d805f833e61493881836103f2565b810190614668565b5f61491e565b905061496491503d8086833e61495c81836103f2565b810190614621565b5f6148bc565b80610efe85614978936103f2565b5f614858565b6149af614989610413565b848152600160208201526001600160a01b0384165f908152609960205260409020614558565b7f000000000000000000000000c0aa79c13bc1eea0168f1a2975b7aa9ff64135ed6001600160a01b0316803b1561045b57604051639926ee7d60e01b8152915f918391829084908290614a0690896004840161457c565b03925af1801561095657614a4b575b50816001600160a01b0382167fe8e68cef1c3a761ed7be7e8463a375f27f7bc335e51824223cacce636ec5c3fe5f80a35f6147f0565b80610efe5f614a59936103f2565b5f614a15565b15614a6657565b6356168b4160e11b5f5260045ffd5b9081602091031261045b575161076381610c5b565b15614a9157565b634c44995d60e01b5f5260045ffd5b15614aa757565b63b187e86960e01b5f5260045ffd5b60209192614b13614b06614b3e989697614aff614adb8783015160018060a01b031690565b6001600160a01b039081165f81815260996020526040902054969091161415614a5f565b5160ff1690565b60ff808516911614612c6e565b604051635401ed2760e01b8152600481019190915260ff909116602482015294859081906044820190565b03817f000000000000000000000000d7c48ab86ab9390e09165013600d5e8721cb3e1d6001600160a01b03165afa92831561095657610422945f94614bc2575b5082614bba92614bb5614ba1936001600160601b03614bad614ba182998b615181565b6001600160601b031690565b911611614a8a565b6151a4565b911610614aa0565b614ba191945092614bba92614bb5614bf96001600160601b039660203d602011614c05575b614bf181836103f2565b810190614a75565b96935050925092614b7e565b503d614be7565b7f00000000000000000000000078469728304326cbc65f8f95fa756b0b731644626001600160a01b03163303614c3e57565b6323d871a560e01b5f5260045ffd5b817fec2963ab21c1e50e1e582aa542af2e4bf7bf38e6e1403c27b42e1c5d6e621eaa614cef614c7a6144f7565b966147aa614c96610598614c9060965460ff1690565b8a613e7f565b614c9f86614e2d565b6001600160c01b0390911690614cb6821515614516565b60018060c01b0316614ccb61476d8284161590565b6001600160a01b0388165f908152609f602052604090206147a39061479c90614792565b0390a26001614d12816147d98460018060a01b03165f52609960205260405f2090565b614d1b81611197565b03614d73575b7f000000000000000000000000c275cce555cc4bd954ff078621778368f9006f1c6001600160a01b0316803b1561045b575f6040518092631fd93ca960e11b82528183816148438a8960048401613f83565b614da4614d7e610413565b838152600160208201526001600160a01b0383165f908152609960205260409020614558565b614d21565b9190805190614db782610433565b91614dc560405193846103f2565b808352614dd4601f1991610433565b013660208401375f5b8151811015614e185780614dff614df660019385612816565b518760986151c2565b63ffffffff614e0e8387612816565b9116905201614ddd565b5090925050565b5f19810191908211612cd257565b805f52609860205260405f20549081155f14614e495750505f90565b5f52609860205260405f20905f198101908111612cd257614e69916113c6565b505460401c90565b307f0000000000000000000000008f78e56bd8371591ab61641376e47706e34208926001600160a01b03161480614f5e575b15614ecc577f6d55bf2dae39a43e197231b59a608bd3e9a503622baf42c30316be9754577ff990565b60405160208101907f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f82527f6ec8a99f0e7f9ebde7354a446dcb9423f3af9c58f386a53c59c5b384f9e82d1160408201527f6bda7e3f385e48841048390444cced5cc795af87758af67622e5f4f0882c4a9960608201524660808201523060a082015260a081526137c660c0826103f2565b507f00000000000000000000000000000000000000000000000000000000000042684614614ea3565b15614f8e57565b631019106960e31b5f5260045ffd5b906101008251116150145781511561500f57602082015160019060f81c81901b5b835182101561500a57600190614ff5614feb612a89614fdd86896132c6565b516001600160f81b03191690565b60ff600191161b90565b90615001818311614f87565b17910190614fbe565b925050565b5f9150565b637da54e4760e11b5f5260045ffd5b90615036825f52609860205260405f2090565b54806150835750615052610422925f52609860205260405f2090565b61507e61505d610424565b4363ffffffff168152925b5f60208501526001600160c01b03166040840152565b6152e2565b916150ae63ffffffff936150a86150a2845f52609860205260405f2090565b91614e1f565b906113c6565b50906150be825463ffffffff1690565b438516941684036150e957506104229250906001600160401b0382549181199060401b169116179055565b815467ffffffff000000001916602085901b67ffffffff00000000161790915561042292919061507e90615125905f52609860205260405f2090565b91615068615131610424565b63ffffffff9095168552565b9061514892916153a9565b1561514f57565b638baa579f60e01b5f5260045ffd5b906001600160601b03809116911602906001600160601b038216918203612cd257565b61519f6001600160601b039161ffff6020612710950151169061515e565b160490565b61519f6001600160601b039161ffff6040612710950151169061515e565b9190815f528260205260405f2054925f5b84811061526b5760405162461bcd60e51b815260206004820152605c60248201527f5265676973747279436f6f7264696e61746f722e67657451756f72756d42697460448201527f6d6170496e6465784174426c6f636b4e756d6265723a206e6f206269746d617060648201527f2075706461746520666f756e6420666f72206f70657261746f72496400000000608482015260a490fd5b808503858111612cd2576106a561528191614e1f565b6152aa61529f8261529a8887905f5260205260405f2090565b6113c6565b505463ffffffff1690565b63ffffffff808616911611156152c357506001016151d3565b94505050505090565b156152d357565b63d51edae360e01b5f5260045ffd5b8054600160401b8110156103d2576152ff916001820181556113c6565b61533d57815160208084015160409485015163ffffffff909316911b67ffffffff00000000161767ffffffffffffffff199190931b16919091179055565b634e487b7160e01b5f525f60045260245ffd5b6005111561080157565b3d15615384573d9061536b826106f7565b9161537960405193846103f2565b82523d5f602084013e565b606090565b9081602091031261045b57516001600160e01b03198116810361045b5790565b9190916153b68284615474565b6153bf81615350565b15908161545e575b50615456575f926153f461540285946040519283916020830195630b135d3f60e11b875260248401613fa5565b03601f1981018352826103f2565b51915afa61540e61535a565b8161544a575b8161541d575090565b8051630b135d3f60e11b92506001600160e01b03199161544591810160209081019101615389565b161490565b80516020149150615414565b505050600190565b6001600160a01b0383811691161490505f6153c7565b8151604181036154a057509061549c91602082015190606060408401519301515f1a906154e2565b9091565b6040036154d95760406020830151920151918260ff1c91601b8301809311612cd25761549c936001600160ff1b03169260ff16906154e2565b50505f90600290565b9291907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083116155805760ff16601b81141580615575575b61556a576020935f93604051938493608085019385528785015260408401526060830152838052039060015afa15610956575f516001600160a01b0381161561556257905f90565b505f90600190565b505050505f90600490565b50601c81141561551a565b505050505f9060039056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a2646970667358221220fb7a8ed824ee20e220263be5e9259676b146bc02693495b6121a6a56a3a5026464736f6c634300081b0033",
}

// ContractRegistryCoordinatorABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractRegistryCoordinatorMetaData.ABI instead.
var ContractRegistryCoordinatorABI = ContractRegistryCoordinatorMetaData.ABI

// ContractRegistryCoordinatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractRegistryCoordinatorMetaData.Bin instead.
var ContractRegistryCoordinatorBin = ContractRegistryCoordinatorMetaData.Bin

// DeployContractRegistryCoordinator deploys a new Ethereum contract, binding an instance of ContractRegistryCoordinator to it.
func DeployContractRegistryCoordinator(auth *bind.TransactOpts, backend bind.ContractBackend, _serviceManager common.Address, _stakeRegistry common.Address, _blsApkRegistry common.Address, _indexRegistry common.Address, _allocationManager common.Address, _pauserRegistry common.Address) (common.Address, *types.Transaction, *ContractRegistryCoordinator, error) {
	parsed, err := ContractRegistryCoordinatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractRegistryCoordinatorBin), backend, _serviceManager, _stakeRegistry, _blsApkRegistry, _indexRegistry, _allocationManager, _pauserRegistry)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractRegistryCoordinator{ContractRegistryCoordinatorCaller: ContractRegistryCoordinatorCaller{contract: contract}, ContractRegistryCoordinatorTransactor: ContractRegistryCoordinatorTransactor{contract: contract}, ContractRegistryCoordinatorFilterer: ContractRegistryCoordinatorFilterer{contract: contract}}, nil
}

// ContractRegistryCoordinatorMethods is an auto generated interface around an Ethereum contract.
type ContractRegistryCoordinatorMethods interface {
	ContractRegistryCoordinatorCalls
	ContractRegistryCoordinatorTransacts
	ContractRegistryCoordinatorFilters
}

// ContractRegistryCoordinatorCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractRegistryCoordinatorCalls interface {
	OPERATORCHURNAPPROVALTYPEHASH(opts *bind.CallOpts) ([32]byte, error)

	PUBKEYREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error)

	AllocationManager(opts *bind.CallOpts) (common.Address, error)

	BlsApkRegistry(opts *bind.CallOpts) (common.Address, error)

	CalculateOperatorChurnApprovalDigestHash(opts *bind.CallOpts, registeringOperator common.Address, registeringOperatorId [32]byte, operatorKickParams []IRegistryCoordinatorOperatorKickParam, salt [32]byte, expiry *big.Int) ([32]byte, error)

	ChurnApprover(opts *bind.CallOpts) (common.Address, error)

	EjectionCooldown(opts *bind.CallOpts) (*big.Int, error)

	Ejector(opts *bind.CallOpts) (common.Address, error)

	GetCurrentQuorumBitmap(opts *bind.CallOpts, operatorId [32]byte) (*big.Int, error)

	GetOperator(opts *bind.CallOpts, operator common.Address) (IRegistryCoordinatorOperatorInfo, error)

	GetOperatorFromId(opts *bind.CallOpts, operatorId [32]byte) (common.Address, error)

	GetOperatorId(opts *bind.CallOpts, operator common.Address) ([32]byte, error)

	GetOperatorSetParams(opts *bind.CallOpts, quorumNumber uint8) (IRegistryCoordinatorOperatorSetParam, error)

	GetOperatorStatus(opts *bind.CallOpts, operator common.Address) (uint8, error)

	GetQuorumBitmapAtBlockNumberByIndex(opts *bind.CallOpts, operatorId [32]byte, blockNumber uint32, index *big.Int) (*big.Int, error)

	GetQuorumBitmapHistoryLength(opts *bind.CallOpts, operatorId [32]byte) (*big.Int, error)

	GetQuorumBitmapIndicesAtBlockNumber(opts *bind.CallOpts, blockNumber uint32, operatorIds [][32]byte) ([]uint32, error)

	GetQuorumBitmapUpdateByIndex(opts *bind.CallOpts, operatorId [32]byte, index *big.Int) (IRegistryCoordinatorQuorumBitmapUpdate, error)

	IndexRegistry(opts *bind.CallOpts) (common.Address, error)

	IsChurnApproverSaltUsed(opts *bind.CallOpts, arg0 [32]byte) (bool, error)

	IsM2Quorum(opts *bind.CallOpts, arg0 uint8) (bool, error)

	IsOperatorSetAVS(opts *bind.CallOpts) (bool, error)

	IsUsingOperatorSets(opts *bind.CallOpts) (bool, error)

	LastEjectionTimestamp(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error)

	NumRegistries(opts *bind.CallOpts) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts, index uint8) (bool, error)

	Paused0(opts *bind.CallOpts) (*big.Int, error)

	PauserRegistry(opts *bind.CallOpts) (common.Address, error)

	PubkeyRegistrationMessageHash(opts *bind.CallOpts, operator common.Address) (BN254G1Point, error)

	QuorumCount(opts *bind.CallOpts) (uint8, error)

	QuorumUpdateBlockNumber(opts *bind.CallOpts, arg0 uint8) (*big.Int, error)

	Registries(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error)

	ServiceManager(opts *bind.CallOpts) (common.Address, error)

	StakeRegistry(opts *bind.CallOpts) (common.Address, error)
}

// ContractRegistryCoordinatorTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractRegistryCoordinatorTransacts interface {
	CreateSlashableStakeQuorum(opts *bind.TransactOpts, operatorSetParams IRegistryCoordinatorOperatorSetParam, minimumStake *big.Int, strategyParams []IStakeRegistryStrategyParams, lookAheadPeriod uint32) (*types.Transaction, error)

	CreateTotalDelegatedStakeQuorum(opts *bind.TransactOpts, operatorSetParams IRegistryCoordinatorOperatorSetParam, minimumStake *big.Int, strategyParams []IStakeRegistryStrategyParams) (*types.Transaction, error)

	DeregisterOperator(opts *bind.TransactOpts, operator common.Address, operatorSetIds []uint32) (*types.Transaction, error)

	DeregisterOperator0(opts *bind.TransactOpts, quorumNumbers []byte) (*types.Transaction, error)

	EjectOperator(opts *bind.TransactOpts, operator common.Address, quorumNumbers []byte) (*types.Transaction, error)

	EnableOperatorSets(opts *bind.TransactOpts) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts, _initialOwner common.Address, _churnApprover common.Address, _ejector common.Address, _initialPausedStatus *big.Int, _operatorSetParams []IRegistryCoordinatorOperatorSetParam, _minimumStakes []*big.Int, _strategyParams [][]IStakeRegistryStrategyParams, _stakeTypes []uint8, _lookAheadPeriods []uint32) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error)

	PauseAll(opts *bind.TransactOpts) (*types.Transaction, error)

	RegisterOperator(opts *bind.TransactOpts, quorumNumbers []byte, socket string, params IBLSApkRegistryPubkeyRegistrationParams, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error)

	RegisterOperator0(opts *bind.TransactOpts, operator common.Address, operatorSetIds []uint32, data []byte) (*types.Transaction, error)

	RegisterOperatorWithChurn(opts *bind.TransactOpts, quorumNumbers []byte, socket string, params IBLSApkRegistryPubkeyRegistrationParams, operatorKickParams []IRegistryCoordinatorOperatorKickParam, churnApproverSignature ISignatureUtilsSignatureWithSaltAndExpiry, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error)

	RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetChurnApprover(opts *bind.TransactOpts, _churnApprover common.Address) (*types.Transaction, error)

	SetEjectionCooldown(opts *bind.TransactOpts, _ejectionCooldown *big.Int) (*types.Transaction, error)

	SetEjector(opts *bind.TransactOpts, _ejector common.Address) (*types.Transaction, error)

	SetOperatorSetParams(opts *bind.TransactOpts, quorumNumber uint8, operatorSetParams IRegistryCoordinatorOperatorSetParam) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error)

	UpdateOperators(opts *bind.TransactOpts, operators []common.Address) (*types.Transaction, error)

	UpdateOperatorsForQuorum(opts *bind.TransactOpts, operatorsPerQuorum [][]common.Address, quorumNumbers []byte) (*types.Transaction, error)

	UpdateSocket(opts *bind.TransactOpts, socket string) (*types.Transaction, error)
}

// ContractRegistryCoordinatorFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractRegistryCoordinatorFilters interface {
	FilterChurnApproverUpdated(opts *bind.FilterOpts) (*ContractRegistryCoordinatorChurnApproverUpdatedIterator, error)
	WatchChurnApproverUpdated(opts *bind.WatchOpts, sink chan<- *ContractRegistryCoordinatorChurnApproverUpdated) (event.Subscription, error)
	ParseChurnApproverUpdated(log types.Log) (*ContractRegistryCoordinatorChurnApproverUpdated, error)

	FilterEjectorUpdated(opts *bind.FilterOpts) (*ContractRegistryCoordinatorEjectorUpdatedIterator, error)
	WatchEjectorUpdated(opts *bind.WatchOpts, sink chan<- *ContractRegistryCoordinatorEjectorUpdated) (event.Subscription, error)
	ParseEjectorUpdated(log types.Log) (*ContractRegistryCoordinatorEjectorUpdated, error)

	FilterInitialized(opts *bind.FilterOpts) (*ContractRegistryCoordinatorInitializedIterator, error)
	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractRegistryCoordinatorInitialized) (event.Subscription, error)
	ParseInitialized(log types.Log) (*ContractRegistryCoordinatorInitialized, error)

	FilterOperatorDeregistered(opts *bind.FilterOpts, operator []common.Address, operatorId [][32]byte) (*ContractRegistryCoordinatorOperatorDeregisteredIterator, error)
	WatchOperatorDeregistered(opts *bind.WatchOpts, sink chan<- *ContractRegistryCoordinatorOperatorDeregistered, operator []common.Address, operatorId [][32]byte) (event.Subscription, error)
	ParseOperatorDeregistered(log types.Log) (*ContractRegistryCoordinatorOperatorDeregistered, error)

	FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address, operatorId [][32]byte) (*ContractRegistryCoordinatorOperatorRegisteredIterator, error)
	WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *ContractRegistryCoordinatorOperatorRegistered, operator []common.Address, operatorId [][32]byte) (event.Subscription, error)
	ParseOperatorRegistered(log types.Log) (*ContractRegistryCoordinatorOperatorRegistered, error)

	FilterOperatorSetParamsUpdated(opts *bind.FilterOpts, quorumNumber []uint8) (*ContractRegistryCoordinatorOperatorSetParamsUpdatedIterator, error)
	WatchOperatorSetParamsUpdated(opts *bind.WatchOpts, sink chan<- *ContractRegistryCoordinatorOperatorSetParamsUpdated, quorumNumber []uint8) (event.Subscription, error)
	ParseOperatorSetParamsUpdated(log types.Log) (*ContractRegistryCoordinatorOperatorSetParamsUpdated, error)

	FilterOperatorSocketUpdate(opts *bind.FilterOpts, operatorId [][32]byte) (*ContractRegistryCoordinatorOperatorSocketUpdateIterator, error)
	WatchOperatorSocketUpdate(opts *bind.WatchOpts, sink chan<- *ContractRegistryCoordinatorOperatorSocketUpdate, operatorId [][32]byte) (event.Subscription, error)
	ParseOperatorSocketUpdate(log types.Log) (*ContractRegistryCoordinatorOperatorSocketUpdate, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractRegistryCoordinatorOwnershipTransferredIterator, error)
	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractRegistryCoordinatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error)
	ParseOwnershipTransferred(log types.Log) (*ContractRegistryCoordinatorOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractRegistryCoordinatorPausedIterator, error)
	WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractRegistryCoordinatorPaused, account []common.Address) (event.Subscription, error)
	ParsePaused(log types.Log) (*ContractRegistryCoordinatorPaused, error)

	FilterQuorumBlockNumberUpdated(opts *bind.FilterOpts, quorumNumber []uint8) (*ContractRegistryCoordinatorQuorumBlockNumberUpdatedIterator, error)
	WatchQuorumBlockNumberUpdated(opts *bind.WatchOpts, sink chan<- *ContractRegistryCoordinatorQuorumBlockNumberUpdated, quorumNumber []uint8) (event.Subscription, error)
	ParseQuorumBlockNumberUpdated(log types.Log) (*ContractRegistryCoordinatorQuorumBlockNumberUpdated, error)

	FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractRegistryCoordinatorUnpausedIterator, error)
	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractRegistryCoordinatorUnpaused, account []common.Address) (event.Subscription, error)
	ParseUnpaused(log types.Log) (*ContractRegistryCoordinatorUnpaused, error)
}

// ContractRegistryCoordinator is an auto generated Go binding around an Ethereum contract.
type ContractRegistryCoordinator struct {
	ContractRegistryCoordinatorCaller     // Read-only binding to the contract
	ContractRegistryCoordinatorTransactor // Write-only binding to the contract
	ContractRegistryCoordinatorFilterer   // Log filterer for contract events
}

// ContractRegistryCoordinator implements the ContractRegistryCoordinatorMethods interface.
var _ ContractRegistryCoordinatorMethods = (*ContractRegistryCoordinator)(nil)

// ContractRegistryCoordinatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractRegistryCoordinatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractRegistryCoordinatorCaller implements the ContractRegistryCoordinatorCalls interface.
var _ ContractRegistryCoordinatorCalls = (*ContractRegistryCoordinatorCaller)(nil)

// ContractRegistryCoordinatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractRegistryCoordinatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractRegistryCoordinatorTransactor implements the ContractRegistryCoordinatorTransacts interface.
var _ ContractRegistryCoordinatorTransacts = (*ContractRegistryCoordinatorTransactor)(nil)

// ContractRegistryCoordinatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractRegistryCoordinatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractRegistryCoordinatorFilterer implements the ContractRegistryCoordinatorFilters interface.
var _ ContractRegistryCoordinatorFilters = (*ContractRegistryCoordinatorFilterer)(nil)

// ContractRegistryCoordinatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractRegistryCoordinatorSession struct {
	Contract     *ContractRegistryCoordinator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ContractRegistryCoordinatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractRegistryCoordinatorCallerSession struct {
	Contract *ContractRegistryCoordinatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// ContractRegistryCoordinatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractRegistryCoordinatorTransactorSession struct {
	Contract     *ContractRegistryCoordinatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// ContractRegistryCoordinatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRegistryCoordinatorRaw struct {
	Contract *ContractRegistryCoordinator // Generic contract binding to access the raw methods on
}

// ContractRegistryCoordinatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractRegistryCoordinatorCallerRaw struct {
	Contract *ContractRegistryCoordinatorCaller // Generic read-only contract binding to access the raw methods on
}

// ContractRegistryCoordinatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractRegistryCoordinatorTransactorRaw struct {
	Contract *ContractRegistryCoordinatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractRegistryCoordinator creates a new instance of ContractRegistryCoordinator, bound to a specific deployed contract.
func NewContractRegistryCoordinator(address common.Address, backend bind.ContractBackend) (*ContractRegistryCoordinator, error) {
	contract, err := bindContractRegistryCoordinator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryCoordinator{ContractRegistryCoordinatorCaller: ContractRegistryCoordinatorCaller{contract: contract}, ContractRegistryCoordinatorTransactor: ContractRegistryCoordinatorTransactor{contract: contract}, ContractRegistryCoordinatorFilterer: ContractRegistryCoordinatorFilterer{contract: contract}}, nil
}

// NewContractRegistryCoordinatorCaller creates a new read-only instance of ContractRegistryCoordinator, bound to a specific deployed contract.
func NewContractRegistryCoordinatorCaller(address common.Address, caller bind.ContractCaller) (*ContractRegistryCoordinatorCaller, error) {
	contract, err := bindContractRegistryCoordinator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryCoordinatorCaller{contract: contract}, nil
}

// NewContractRegistryCoordinatorTransactor creates a new write-only instance of ContractRegistryCoordinator, bound to a specific deployed contract.
func NewContractRegistryCoordinatorTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractRegistryCoordinatorTransactor, error) {
	contract, err := bindContractRegistryCoordinator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryCoordinatorTransactor{contract: contract}, nil
}

// NewContractRegistryCoordinatorFilterer creates a new log filterer instance of ContractRegistryCoordinator, bound to a specific deployed contract.
func NewContractRegistryCoordinatorFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractRegistryCoordinatorFilterer, error) {
	contract, err := bindContractRegistryCoordinator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryCoordinatorFilterer{contract: contract}, nil
}

// bindContractRegistryCoordinator binds a generic wrapper to an already deployed contract.
func bindContractRegistryCoordinator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractRegistryCoordinatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractRegistryCoordinator.Contract.ContractRegistryCoordinatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.ContractRegistryCoordinatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.ContractRegistryCoordinatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractRegistryCoordinator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.contract.Transact(opts, method, params...)
}

// OPERATORCHURNAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0xca0de882.
//
// Solidity: function OPERATOR_CHURN_APPROVAL_TYPEHASH() view returns(bytes32)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) OPERATORCHURNAPPROVALTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "OPERATOR_CHURN_APPROVAL_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORCHURNAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0xca0de882.
//
// Solidity: function OPERATOR_CHURN_APPROVAL_TYPEHASH() view returns(bytes32)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) OPERATORCHURNAPPROVALTYPEHASH() ([32]byte, error) {
	return _ContractRegistryCoordinator.Contract.OPERATORCHURNAPPROVALTYPEHASH(&_ContractRegistryCoordinator.CallOpts)
}

// OPERATORCHURNAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0xca0de882.
//
// Solidity: function OPERATOR_CHURN_APPROVAL_TYPEHASH() view returns(bytes32)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) OPERATORCHURNAPPROVALTYPEHASH() ([32]byte, error) {
	return _ContractRegistryCoordinator.Contract.OPERATORCHURNAPPROVALTYPEHASH(&_ContractRegistryCoordinator.CallOpts)
}

// PUBKEYREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0x9feab859.
//
// Solidity: function PUBKEY_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) PUBKEYREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "PUBKEY_REGISTRATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PUBKEYREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0x9feab859.
//
// Solidity: function PUBKEY_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) PUBKEYREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _ContractRegistryCoordinator.Contract.PUBKEYREGISTRATIONTYPEHASH(&_ContractRegistryCoordinator.CallOpts)
}

// PUBKEYREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0x9feab859.
//
// Solidity: function PUBKEY_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) PUBKEYREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _ContractRegistryCoordinator.Contract.PUBKEYREGISTRATIONTYPEHASH(&_ContractRegistryCoordinator.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) AllocationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "allocationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) AllocationManager() (common.Address, error) {
	return _ContractRegistryCoordinator.Contract.AllocationManager(&_ContractRegistryCoordinator.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) AllocationManager() (common.Address, error) {
	return _ContractRegistryCoordinator.Contract.AllocationManager(&_ContractRegistryCoordinator.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) BlsApkRegistry() (common.Address, error) {
	return _ContractRegistryCoordinator.Contract.BlsApkRegistry(&_ContractRegistryCoordinator.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractRegistryCoordinator.Contract.BlsApkRegistry(&_ContractRegistryCoordinator.CallOpts)
}

// CalculateOperatorChurnApprovalDigestHash is a free data retrieval call binding the contract method 0x84ca5213.
//
// Solidity: function calculateOperatorChurnApprovalDigestHash(address registeringOperator, bytes32 registeringOperatorId, (uint8,address)[] operatorKickParams, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) CalculateOperatorChurnApprovalDigestHash(opts *bind.CallOpts, registeringOperator common.Address, registeringOperatorId [32]byte, operatorKickParams []IRegistryCoordinatorOperatorKickParam, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "calculateOperatorChurnApprovalDigestHash", registeringOperator, registeringOperatorId, operatorKickParams, salt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateOperatorChurnApprovalDigestHash is a free data retrieval call binding the contract method 0x84ca5213.
//
// Solidity: function calculateOperatorChurnApprovalDigestHash(address registeringOperator, bytes32 registeringOperatorId, (uint8,address)[] operatorKickParams, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) CalculateOperatorChurnApprovalDigestHash(registeringOperator common.Address, registeringOperatorId [32]byte, operatorKickParams []IRegistryCoordinatorOperatorKickParam, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _ContractRegistryCoordinator.Contract.CalculateOperatorChurnApprovalDigestHash(&_ContractRegistryCoordinator.CallOpts, registeringOperator, registeringOperatorId, operatorKickParams, salt, expiry)
}

// CalculateOperatorChurnApprovalDigestHash is a free data retrieval call binding the contract method 0x84ca5213.
//
// Solidity: function calculateOperatorChurnApprovalDigestHash(address registeringOperator, bytes32 registeringOperatorId, (uint8,address)[] operatorKickParams, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) CalculateOperatorChurnApprovalDigestHash(registeringOperator common.Address, registeringOperatorId [32]byte, operatorKickParams []IRegistryCoordinatorOperatorKickParam, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _ContractRegistryCoordinator.Contract.CalculateOperatorChurnApprovalDigestHash(&_ContractRegistryCoordinator.CallOpts, registeringOperator, registeringOperatorId, operatorKickParams, salt, expiry)
}

// ChurnApprover is a free data retrieval call binding the contract method 0x054310e6.
//
// Solidity: function churnApprover() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) ChurnApprover(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "churnApprover")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChurnApprover is a free data retrieval call binding the contract method 0x054310e6.
//
// Solidity: function churnApprover() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) ChurnApprover() (common.Address, error) {
	return _ContractRegistryCoordinator.Contract.ChurnApprover(&_ContractRegistryCoordinator.CallOpts)
}

// ChurnApprover is a free data retrieval call binding the contract method 0x054310e6.
//
// Solidity: function churnApprover() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) ChurnApprover() (common.Address, error) {
	return _ContractRegistryCoordinator.Contract.ChurnApprover(&_ContractRegistryCoordinator.CallOpts)
}

// EjectionCooldown is a free data retrieval call binding the contract method 0xa96f783e.
//
// Solidity: function ejectionCooldown() view returns(uint256)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) EjectionCooldown(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "ejectionCooldown")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EjectionCooldown is a free data retrieval call binding the contract method 0xa96f783e.
//
// Solidity: function ejectionCooldown() view returns(uint256)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) EjectionCooldown() (*big.Int, error) {
	return _ContractRegistryCoordinator.Contract.EjectionCooldown(&_ContractRegistryCoordinator.CallOpts)
}

// EjectionCooldown is a free data retrieval call binding the contract method 0xa96f783e.
//
// Solidity: function ejectionCooldown() view returns(uint256)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) EjectionCooldown() (*big.Int, error) {
	return _ContractRegistryCoordinator.Contract.EjectionCooldown(&_ContractRegistryCoordinator.CallOpts)
}

// Ejector is a free data retrieval call binding the contract method 0x28f61b31.
//
// Solidity: function ejector() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) Ejector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "ejector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ejector is a free data retrieval call binding the contract method 0x28f61b31.
//
// Solidity: function ejector() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) Ejector() (common.Address, error) {
	return _ContractRegistryCoordinator.Contract.Ejector(&_ContractRegistryCoordinator.CallOpts)
}

// Ejector is a free data retrieval call binding the contract method 0x28f61b31.
//
// Solidity: function ejector() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) Ejector() (common.Address, error) {
	return _ContractRegistryCoordinator.Contract.Ejector(&_ContractRegistryCoordinator.CallOpts)
}

// GetCurrentQuorumBitmap is a free data retrieval call binding the contract method 0x871ef049.
//
// Solidity: function getCurrentQuorumBitmap(bytes32 operatorId) view returns(uint192)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) GetCurrentQuorumBitmap(opts *bind.CallOpts, operatorId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "getCurrentQuorumBitmap", operatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentQuorumBitmap is a free data retrieval call binding the contract method 0x871ef049.
//
// Solidity: function getCurrentQuorumBitmap(bytes32 operatorId) view returns(uint192)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) GetCurrentQuorumBitmap(operatorId [32]byte) (*big.Int, error) {
	return _ContractRegistryCoordinator.Contract.GetCurrentQuorumBitmap(&_ContractRegistryCoordinator.CallOpts, operatorId)
}

// GetCurrentQuorumBitmap is a free data retrieval call binding the contract method 0x871ef049.
//
// Solidity: function getCurrentQuorumBitmap(bytes32 operatorId) view returns(uint192)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) GetCurrentQuorumBitmap(operatorId [32]byte) (*big.Int, error) {
	return _ContractRegistryCoordinator.Contract.GetCurrentQuorumBitmap(&_ContractRegistryCoordinator.CallOpts, operatorId)
}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address operator) view returns((bytes32,uint8))
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) GetOperator(opts *bind.CallOpts, operator common.Address) (IRegistryCoordinatorOperatorInfo, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "getOperator", operator)

	if err != nil {
		return *new(IRegistryCoordinatorOperatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IRegistryCoordinatorOperatorInfo)).(*IRegistryCoordinatorOperatorInfo)

	return out0, err

}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address operator) view returns((bytes32,uint8))
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) GetOperator(operator common.Address) (IRegistryCoordinatorOperatorInfo, error) {
	return _ContractRegistryCoordinator.Contract.GetOperator(&_ContractRegistryCoordinator.CallOpts, operator)
}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address operator) view returns((bytes32,uint8))
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) GetOperator(operator common.Address) (IRegistryCoordinatorOperatorInfo, error) {
	return _ContractRegistryCoordinator.Contract.GetOperator(&_ContractRegistryCoordinator.CallOpts, operator)
}

// GetOperatorFromId is a free data retrieval call binding the contract method 0x296bb064.
//
// Solidity: function getOperatorFromId(bytes32 operatorId) view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) GetOperatorFromId(opts *bind.CallOpts, operatorId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "getOperatorFromId", operatorId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperatorFromId is a free data retrieval call binding the contract method 0x296bb064.
//
// Solidity: function getOperatorFromId(bytes32 operatorId) view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) GetOperatorFromId(operatorId [32]byte) (common.Address, error) {
	return _ContractRegistryCoordinator.Contract.GetOperatorFromId(&_ContractRegistryCoordinator.CallOpts, operatorId)
}

// GetOperatorFromId is a free data retrieval call binding the contract method 0x296bb064.
//
// Solidity: function getOperatorFromId(bytes32 operatorId) view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) GetOperatorFromId(operatorId [32]byte) (common.Address, error) {
	return _ContractRegistryCoordinator.Contract.GetOperatorFromId(&_ContractRegistryCoordinator.CallOpts, operatorId)
}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) GetOperatorId(opts *bind.CallOpts, operator common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "getOperatorId", operator)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) GetOperatorId(operator common.Address) ([32]byte, error) {
	return _ContractRegistryCoordinator.Contract.GetOperatorId(&_ContractRegistryCoordinator.CallOpts, operator)
}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) GetOperatorId(operator common.Address) ([32]byte, error) {
	return _ContractRegistryCoordinator.Contract.GetOperatorId(&_ContractRegistryCoordinator.CallOpts, operator)
}

// GetOperatorSetParams is a free data retrieval call binding the contract method 0xe65797ad.
//
// Solidity: function getOperatorSetParams(uint8 quorumNumber) view returns((uint32,uint16,uint16))
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) GetOperatorSetParams(opts *bind.CallOpts, quorumNumber uint8) (IRegistryCoordinatorOperatorSetParam, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "getOperatorSetParams", quorumNumber)

	if err != nil {
		return *new(IRegistryCoordinatorOperatorSetParam), err
	}

	out0 := *abi.ConvertType(out[0], new(IRegistryCoordinatorOperatorSetParam)).(*IRegistryCoordinatorOperatorSetParam)

	return out0, err

}

// GetOperatorSetParams is a free data retrieval call binding the contract method 0xe65797ad.
//
// Solidity: function getOperatorSetParams(uint8 quorumNumber) view returns((uint32,uint16,uint16))
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) GetOperatorSetParams(quorumNumber uint8) (IRegistryCoordinatorOperatorSetParam, error) {
	return _ContractRegistryCoordinator.Contract.GetOperatorSetParams(&_ContractRegistryCoordinator.CallOpts, quorumNumber)
}

// GetOperatorSetParams is a free data retrieval call binding the contract method 0xe65797ad.
//
// Solidity: function getOperatorSetParams(uint8 quorumNumber) view returns((uint32,uint16,uint16))
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) GetOperatorSetParams(quorumNumber uint8) (IRegistryCoordinatorOperatorSetParam, error) {
	return _ContractRegistryCoordinator.Contract.GetOperatorSetParams(&_ContractRegistryCoordinator.CallOpts, quorumNumber)
}

// GetOperatorStatus is a free data retrieval call binding the contract method 0xfd39105a.
//
// Solidity: function getOperatorStatus(address operator) view returns(uint8)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) GetOperatorStatus(opts *bind.CallOpts, operator common.Address) (uint8, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "getOperatorStatus", operator)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetOperatorStatus is a free data retrieval call binding the contract method 0xfd39105a.
//
// Solidity: function getOperatorStatus(address operator) view returns(uint8)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) GetOperatorStatus(operator common.Address) (uint8, error) {
	return _ContractRegistryCoordinator.Contract.GetOperatorStatus(&_ContractRegistryCoordinator.CallOpts, operator)
}

// GetOperatorStatus is a free data retrieval call binding the contract method 0xfd39105a.
//
// Solidity: function getOperatorStatus(address operator) view returns(uint8)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) GetOperatorStatus(operator common.Address) (uint8, error) {
	return _ContractRegistryCoordinator.Contract.GetOperatorStatus(&_ContractRegistryCoordinator.CallOpts, operator)
}

// GetQuorumBitmapAtBlockNumberByIndex is a free data retrieval call binding the contract method 0x04ec6351.
//
// Solidity: function getQuorumBitmapAtBlockNumberByIndex(bytes32 operatorId, uint32 blockNumber, uint256 index) view returns(uint192)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) GetQuorumBitmapAtBlockNumberByIndex(opts *bind.CallOpts, operatorId [32]byte, blockNumber uint32, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "getQuorumBitmapAtBlockNumberByIndex", operatorId, blockNumber, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuorumBitmapAtBlockNumberByIndex is a free data retrieval call binding the contract method 0x04ec6351.
//
// Solidity: function getQuorumBitmapAtBlockNumberByIndex(bytes32 operatorId, uint32 blockNumber, uint256 index) view returns(uint192)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) GetQuorumBitmapAtBlockNumberByIndex(operatorId [32]byte, blockNumber uint32, index *big.Int) (*big.Int, error) {
	return _ContractRegistryCoordinator.Contract.GetQuorumBitmapAtBlockNumberByIndex(&_ContractRegistryCoordinator.CallOpts, operatorId, blockNumber, index)
}

// GetQuorumBitmapAtBlockNumberByIndex is a free data retrieval call binding the contract method 0x04ec6351.
//
// Solidity: function getQuorumBitmapAtBlockNumberByIndex(bytes32 operatorId, uint32 blockNumber, uint256 index) view returns(uint192)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) GetQuorumBitmapAtBlockNumberByIndex(operatorId [32]byte, blockNumber uint32, index *big.Int) (*big.Int, error) {
	return _ContractRegistryCoordinator.Contract.GetQuorumBitmapAtBlockNumberByIndex(&_ContractRegistryCoordinator.CallOpts, operatorId, blockNumber, index)
}

// GetQuorumBitmapHistoryLength is a free data retrieval call binding the contract method 0x03fd3492.
//
// Solidity: function getQuorumBitmapHistoryLength(bytes32 operatorId) view returns(uint256)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) GetQuorumBitmapHistoryLength(opts *bind.CallOpts, operatorId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "getQuorumBitmapHistoryLength", operatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuorumBitmapHistoryLength is a free data retrieval call binding the contract method 0x03fd3492.
//
// Solidity: function getQuorumBitmapHistoryLength(bytes32 operatorId) view returns(uint256)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) GetQuorumBitmapHistoryLength(operatorId [32]byte) (*big.Int, error) {
	return _ContractRegistryCoordinator.Contract.GetQuorumBitmapHistoryLength(&_ContractRegistryCoordinator.CallOpts, operatorId)
}

// GetQuorumBitmapHistoryLength is a free data retrieval call binding the contract method 0x03fd3492.
//
// Solidity: function getQuorumBitmapHistoryLength(bytes32 operatorId) view returns(uint256)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) GetQuorumBitmapHistoryLength(operatorId [32]byte) (*big.Int, error) {
	return _ContractRegistryCoordinator.Contract.GetQuorumBitmapHistoryLength(&_ContractRegistryCoordinator.CallOpts, operatorId)
}

// GetQuorumBitmapIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xc391425e.
//
// Solidity: function getQuorumBitmapIndicesAtBlockNumber(uint32 blockNumber, bytes32[] operatorIds) view returns(uint32[])
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) GetQuorumBitmapIndicesAtBlockNumber(opts *bind.CallOpts, blockNumber uint32, operatorIds [][32]byte) ([]uint32, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "getQuorumBitmapIndicesAtBlockNumber", blockNumber, operatorIds)

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// GetQuorumBitmapIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xc391425e.
//
// Solidity: function getQuorumBitmapIndicesAtBlockNumber(uint32 blockNumber, bytes32[] operatorIds) view returns(uint32[])
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) GetQuorumBitmapIndicesAtBlockNumber(blockNumber uint32, operatorIds [][32]byte) ([]uint32, error) {
	return _ContractRegistryCoordinator.Contract.GetQuorumBitmapIndicesAtBlockNumber(&_ContractRegistryCoordinator.CallOpts, blockNumber, operatorIds)
}

// GetQuorumBitmapIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xc391425e.
//
// Solidity: function getQuorumBitmapIndicesAtBlockNumber(uint32 blockNumber, bytes32[] operatorIds) view returns(uint32[])
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) GetQuorumBitmapIndicesAtBlockNumber(blockNumber uint32, operatorIds [][32]byte) ([]uint32, error) {
	return _ContractRegistryCoordinator.Contract.GetQuorumBitmapIndicesAtBlockNumber(&_ContractRegistryCoordinator.CallOpts, blockNumber, operatorIds)
}

// GetQuorumBitmapUpdateByIndex is a free data retrieval call binding the contract method 0x1eb812da.
//
// Solidity: function getQuorumBitmapUpdateByIndex(bytes32 operatorId, uint256 index) view returns((uint32,uint32,uint192))
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) GetQuorumBitmapUpdateByIndex(opts *bind.CallOpts, operatorId [32]byte, index *big.Int) (IRegistryCoordinatorQuorumBitmapUpdate, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "getQuorumBitmapUpdateByIndex", operatorId, index)

	if err != nil {
		return *new(IRegistryCoordinatorQuorumBitmapUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new(IRegistryCoordinatorQuorumBitmapUpdate)).(*IRegistryCoordinatorQuorumBitmapUpdate)

	return out0, err

}

// GetQuorumBitmapUpdateByIndex is a free data retrieval call binding the contract method 0x1eb812da.
//
// Solidity: function getQuorumBitmapUpdateByIndex(bytes32 operatorId, uint256 index) view returns((uint32,uint32,uint192))
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) GetQuorumBitmapUpdateByIndex(operatorId [32]byte, index *big.Int) (IRegistryCoordinatorQuorumBitmapUpdate, error) {
	return _ContractRegistryCoordinator.Contract.GetQuorumBitmapUpdateByIndex(&_ContractRegistryCoordinator.CallOpts, operatorId, index)
}

// GetQuorumBitmapUpdateByIndex is a free data retrieval call binding the contract method 0x1eb812da.
//
// Solidity: function getQuorumBitmapUpdateByIndex(bytes32 operatorId, uint256 index) view returns((uint32,uint32,uint192))
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) GetQuorumBitmapUpdateByIndex(operatorId [32]byte, index *big.Int) (IRegistryCoordinatorQuorumBitmapUpdate, error) {
	return _ContractRegistryCoordinator.Contract.GetQuorumBitmapUpdateByIndex(&_ContractRegistryCoordinator.CallOpts, operatorId, index)
}

// IndexRegistry is a free data retrieval call binding the contract method 0x9e9923c2.
//
// Solidity: function indexRegistry() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) IndexRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "indexRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IndexRegistry is a free data retrieval call binding the contract method 0x9e9923c2.
//
// Solidity: function indexRegistry() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) IndexRegistry() (common.Address, error) {
	return _ContractRegistryCoordinator.Contract.IndexRegistry(&_ContractRegistryCoordinator.CallOpts)
}

// IndexRegistry is a free data retrieval call binding the contract method 0x9e9923c2.
//
// Solidity: function indexRegistry() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) IndexRegistry() (common.Address, error) {
	return _ContractRegistryCoordinator.Contract.IndexRegistry(&_ContractRegistryCoordinator.CallOpts)
}

// IsChurnApproverSaltUsed is a free data retrieval call binding the contract method 0x1478851f.
//
// Solidity: function isChurnApproverSaltUsed(bytes32 ) view returns(bool)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) IsChurnApproverSaltUsed(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "isChurnApproverSaltUsed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsChurnApproverSaltUsed is a free data retrieval call binding the contract method 0x1478851f.
//
// Solidity: function isChurnApproverSaltUsed(bytes32 ) view returns(bool)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) IsChurnApproverSaltUsed(arg0 [32]byte) (bool, error) {
	return _ContractRegistryCoordinator.Contract.IsChurnApproverSaltUsed(&_ContractRegistryCoordinator.CallOpts, arg0)
}

// IsChurnApproverSaltUsed is a free data retrieval call binding the contract method 0x1478851f.
//
// Solidity: function isChurnApproverSaltUsed(bytes32 ) view returns(bool)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) IsChurnApproverSaltUsed(arg0 [32]byte) (bool, error) {
	return _ContractRegistryCoordinator.Contract.IsChurnApproverSaltUsed(&_ContractRegistryCoordinator.CallOpts, arg0)
}

// IsM2Quorum is a free data retrieval call binding the contract method 0xa4d7871f.
//
// Solidity: function isM2Quorum(uint8 ) view returns(bool)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) IsM2Quorum(opts *bind.CallOpts, arg0 uint8) (bool, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "isM2Quorum", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsM2Quorum is a free data retrieval call binding the contract method 0xa4d7871f.
//
// Solidity: function isM2Quorum(uint8 ) view returns(bool)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) IsM2Quorum(arg0 uint8) (bool, error) {
	return _ContractRegistryCoordinator.Contract.IsM2Quorum(&_ContractRegistryCoordinator.CallOpts, arg0)
}

// IsM2Quorum is a free data retrieval call binding the contract method 0xa4d7871f.
//
// Solidity: function isM2Quorum(uint8 ) view returns(bool)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) IsM2Quorum(arg0 uint8) (bool, error) {
	return _ContractRegistryCoordinator.Contract.IsM2Quorum(&_ContractRegistryCoordinator.CallOpts, arg0)
}

// IsOperatorSetAVS is a free data retrieval call binding the contract method 0xcabbb17f.
//
// Solidity: function isOperatorSetAVS() view returns(bool)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) IsOperatorSetAVS(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "isOperatorSetAVS")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorSetAVS is a free data retrieval call binding the contract method 0xcabbb17f.
//
// Solidity: function isOperatorSetAVS() view returns(bool)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) IsOperatorSetAVS() (bool, error) {
	return _ContractRegistryCoordinator.Contract.IsOperatorSetAVS(&_ContractRegistryCoordinator.CallOpts)
}

// IsOperatorSetAVS is a free data retrieval call binding the contract method 0xcabbb17f.
//
// Solidity: function isOperatorSetAVS() view returns(bool)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) IsOperatorSetAVS() (bool, error) {
	return _ContractRegistryCoordinator.Contract.IsOperatorSetAVS(&_ContractRegistryCoordinator.CallOpts)
}

// IsUsingOperatorSets is a free data retrieval call binding the contract method 0xbd33ee24.
//
// Solidity: function isUsingOperatorSets() view returns(bool)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) IsUsingOperatorSets(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "isUsingOperatorSets")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUsingOperatorSets is a free data retrieval call binding the contract method 0xbd33ee24.
//
// Solidity: function isUsingOperatorSets() view returns(bool)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) IsUsingOperatorSets() (bool, error) {
	return _ContractRegistryCoordinator.Contract.IsUsingOperatorSets(&_ContractRegistryCoordinator.CallOpts)
}

// IsUsingOperatorSets is a free data retrieval call binding the contract method 0xbd33ee24.
//
// Solidity: function isUsingOperatorSets() view returns(bool)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) IsUsingOperatorSets() (bool, error) {
	return _ContractRegistryCoordinator.Contract.IsUsingOperatorSets(&_ContractRegistryCoordinator.CallOpts)
}

// LastEjectionTimestamp is a free data retrieval call binding the contract method 0x125e0584.
//
// Solidity: function lastEjectionTimestamp(address ) view returns(uint256)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) LastEjectionTimestamp(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "lastEjectionTimestamp", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastEjectionTimestamp is a free data retrieval call binding the contract method 0x125e0584.
//
// Solidity: function lastEjectionTimestamp(address ) view returns(uint256)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) LastEjectionTimestamp(arg0 common.Address) (*big.Int, error) {
	return _ContractRegistryCoordinator.Contract.LastEjectionTimestamp(&_ContractRegistryCoordinator.CallOpts, arg0)
}

// LastEjectionTimestamp is a free data retrieval call binding the contract method 0x125e0584.
//
// Solidity: function lastEjectionTimestamp(address ) view returns(uint256)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) LastEjectionTimestamp(arg0 common.Address) (*big.Int, error) {
	return _ContractRegistryCoordinator.Contract.LastEjectionTimestamp(&_ContractRegistryCoordinator.CallOpts, arg0)
}

// NumRegistries is a free data retrieval call binding the contract method 0xd72d8dd6.
//
// Solidity: function numRegistries() view returns(uint256)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) NumRegistries(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "numRegistries")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumRegistries is a free data retrieval call binding the contract method 0xd72d8dd6.
//
// Solidity: function numRegistries() view returns(uint256)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) NumRegistries() (*big.Int, error) {
	return _ContractRegistryCoordinator.Contract.NumRegistries(&_ContractRegistryCoordinator.CallOpts)
}

// NumRegistries is a free data retrieval call binding the contract method 0xd72d8dd6.
//
// Solidity: function numRegistries() view returns(uint256)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) NumRegistries() (*big.Int, error) {
	return _ContractRegistryCoordinator.Contract.NumRegistries(&_ContractRegistryCoordinator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) Owner() (common.Address, error) {
	return _ContractRegistryCoordinator.Contract.Owner(&_ContractRegistryCoordinator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) Owner() (common.Address, error) {
	return _ContractRegistryCoordinator.Contract.Owner(&_ContractRegistryCoordinator.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) Paused(index uint8) (bool, error) {
	return _ContractRegistryCoordinator.Contract.Paused(&_ContractRegistryCoordinator.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) Paused(index uint8) (bool, error) {
	return _ContractRegistryCoordinator.Contract.Paused(&_ContractRegistryCoordinator.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) Paused0() (*big.Int, error) {
	return _ContractRegistryCoordinator.Contract.Paused0(&_ContractRegistryCoordinator.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) Paused0() (*big.Int, error) {
	return _ContractRegistryCoordinator.Contract.Paused0(&_ContractRegistryCoordinator.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) PauserRegistry() (common.Address, error) {
	return _ContractRegistryCoordinator.Contract.PauserRegistry(&_ContractRegistryCoordinator.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractRegistryCoordinator.Contract.PauserRegistry(&_ContractRegistryCoordinator.CallOpts)
}

// PubkeyRegistrationMessageHash is a free data retrieval call binding the contract method 0x3c2a7f4c.
//
// Solidity: function pubkeyRegistrationMessageHash(address operator) view returns((uint256,uint256))
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) PubkeyRegistrationMessageHash(opts *bind.CallOpts, operator common.Address) (BN254G1Point, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "pubkeyRegistrationMessageHash", operator)

	if err != nil {
		return *new(BN254G1Point), err
	}

	out0 := *abi.ConvertType(out[0], new(BN254G1Point)).(*BN254G1Point)

	return out0, err

}

// PubkeyRegistrationMessageHash is a free data retrieval call binding the contract method 0x3c2a7f4c.
//
// Solidity: function pubkeyRegistrationMessageHash(address operator) view returns((uint256,uint256))
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) PubkeyRegistrationMessageHash(operator common.Address) (BN254G1Point, error) {
	return _ContractRegistryCoordinator.Contract.PubkeyRegistrationMessageHash(&_ContractRegistryCoordinator.CallOpts, operator)
}

// PubkeyRegistrationMessageHash is a free data retrieval call binding the contract method 0x3c2a7f4c.
//
// Solidity: function pubkeyRegistrationMessageHash(address operator) view returns((uint256,uint256))
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) PubkeyRegistrationMessageHash(operator common.Address) (BN254G1Point, error) {
	return _ContractRegistryCoordinator.Contract.PubkeyRegistrationMessageHash(&_ContractRegistryCoordinator.CallOpts, operator)
}

// QuorumCount is a free data retrieval call binding the contract method 0x9aa1653d.
//
// Solidity: function quorumCount() view returns(uint8)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) QuorumCount(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "quorumCount")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// QuorumCount is a free data retrieval call binding the contract method 0x9aa1653d.
//
// Solidity: function quorumCount() view returns(uint8)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) QuorumCount() (uint8, error) {
	return _ContractRegistryCoordinator.Contract.QuorumCount(&_ContractRegistryCoordinator.CallOpts)
}

// QuorumCount is a free data retrieval call binding the contract method 0x9aa1653d.
//
// Solidity: function quorumCount() view returns(uint8)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) QuorumCount() (uint8, error) {
	return _ContractRegistryCoordinator.Contract.QuorumCount(&_ContractRegistryCoordinator.CallOpts)
}

// QuorumUpdateBlockNumber is a free data retrieval call binding the contract method 0x249a0c42.
//
// Solidity: function quorumUpdateBlockNumber(uint8 ) view returns(uint256)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) QuorumUpdateBlockNumber(opts *bind.CallOpts, arg0 uint8) (*big.Int, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "quorumUpdateBlockNumber", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuorumUpdateBlockNumber is a free data retrieval call binding the contract method 0x249a0c42.
//
// Solidity: function quorumUpdateBlockNumber(uint8 ) view returns(uint256)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) QuorumUpdateBlockNumber(arg0 uint8) (*big.Int, error) {
	return _ContractRegistryCoordinator.Contract.QuorumUpdateBlockNumber(&_ContractRegistryCoordinator.CallOpts, arg0)
}

// QuorumUpdateBlockNumber is a free data retrieval call binding the contract method 0x249a0c42.
//
// Solidity: function quorumUpdateBlockNumber(uint8 ) view returns(uint256)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) QuorumUpdateBlockNumber(arg0 uint8) (*big.Int, error) {
	return _ContractRegistryCoordinator.Contract.QuorumUpdateBlockNumber(&_ContractRegistryCoordinator.CallOpts, arg0)
}

// Registries is a free data retrieval call binding the contract method 0x6347c900.
//
// Solidity: function registries(uint256 ) view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) Registries(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "registries", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registries is a free data retrieval call binding the contract method 0x6347c900.
//
// Solidity: function registries(uint256 ) view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) Registries(arg0 *big.Int) (common.Address, error) {
	return _ContractRegistryCoordinator.Contract.Registries(&_ContractRegistryCoordinator.CallOpts, arg0)
}

// Registries is a free data retrieval call binding the contract method 0x6347c900.
//
// Solidity: function registries(uint256 ) view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) Registries(arg0 *big.Int) (common.Address, error) {
	return _ContractRegistryCoordinator.Contract.Registries(&_ContractRegistryCoordinator.CallOpts, arg0)
}

// ServiceManager is a free data retrieval call binding the contract method 0x3998fdd3.
//
// Solidity: function serviceManager() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) ServiceManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "serviceManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ServiceManager is a free data retrieval call binding the contract method 0x3998fdd3.
//
// Solidity: function serviceManager() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) ServiceManager() (common.Address, error) {
	return _ContractRegistryCoordinator.Contract.ServiceManager(&_ContractRegistryCoordinator.CallOpts)
}

// ServiceManager is a free data retrieval call binding the contract method 0x3998fdd3.
//
// Solidity: function serviceManager() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) ServiceManager() (common.Address, error) {
	return _ContractRegistryCoordinator.Contract.ServiceManager(&_ContractRegistryCoordinator.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractRegistryCoordinator.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) StakeRegistry() (common.Address, error) {
	return _ContractRegistryCoordinator.Contract.StakeRegistry(&_ContractRegistryCoordinator.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession) StakeRegistry() (common.Address, error) {
	return _ContractRegistryCoordinator.Contract.StakeRegistry(&_ContractRegistryCoordinator.CallOpts)
}

// CreateSlashableStakeQuorum is a paid mutator transaction binding the contract method 0x3eef3a51.
//
// Solidity: function createSlashableStakeQuorum((uint32,uint16,uint16) operatorSetParams, uint96 minimumStake, (address,uint96)[] strategyParams, uint32 lookAheadPeriod) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor) CreateSlashableStakeQuorum(opts *bind.TransactOpts, operatorSetParams IRegistryCoordinatorOperatorSetParam, minimumStake *big.Int, strategyParams []IStakeRegistryStrategyParams, lookAheadPeriod uint32) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.contract.Transact(opts, "createSlashableStakeQuorum", operatorSetParams, minimumStake, strategyParams, lookAheadPeriod)
}

// CreateSlashableStakeQuorum is a paid mutator transaction binding the contract method 0x3eef3a51.
//
// Solidity: function createSlashableStakeQuorum((uint32,uint16,uint16) operatorSetParams, uint96 minimumStake, (address,uint96)[] strategyParams, uint32 lookAheadPeriod) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) CreateSlashableStakeQuorum(operatorSetParams IRegistryCoordinatorOperatorSetParam, minimumStake *big.Int, strategyParams []IStakeRegistryStrategyParams, lookAheadPeriod uint32) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.CreateSlashableStakeQuorum(&_ContractRegistryCoordinator.TransactOpts, operatorSetParams, minimumStake, strategyParams, lookAheadPeriod)
}

// CreateSlashableStakeQuorum is a paid mutator transaction binding the contract method 0x3eef3a51.
//
// Solidity: function createSlashableStakeQuorum((uint32,uint16,uint16) operatorSetParams, uint96 minimumStake, (address,uint96)[] strategyParams, uint32 lookAheadPeriod) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession) CreateSlashableStakeQuorum(operatorSetParams IRegistryCoordinatorOperatorSetParam, minimumStake *big.Int, strategyParams []IStakeRegistryStrategyParams, lookAheadPeriod uint32) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.CreateSlashableStakeQuorum(&_ContractRegistryCoordinator.TransactOpts, operatorSetParams, minimumStake, strategyParams, lookAheadPeriod)
}

// CreateTotalDelegatedStakeQuorum is a paid mutator transaction binding the contract method 0x8281ab75.
//
// Solidity: function createTotalDelegatedStakeQuorum((uint32,uint16,uint16) operatorSetParams, uint96 minimumStake, (address,uint96)[] strategyParams) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor) CreateTotalDelegatedStakeQuorum(opts *bind.TransactOpts, operatorSetParams IRegistryCoordinatorOperatorSetParam, minimumStake *big.Int, strategyParams []IStakeRegistryStrategyParams) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.contract.Transact(opts, "createTotalDelegatedStakeQuorum", operatorSetParams, minimumStake, strategyParams)
}

// CreateTotalDelegatedStakeQuorum is a paid mutator transaction binding the contract method 0x8281ab75.
//
// Solidity: function createTotalDelegatedStakeQuorum((uint32,uint16,uint16) operatorSetParams, uint96 minimumStake, (address,uint96)[] strategyParams) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) CreateTotalDelegatedStakeQuorum(operatorSetParams IRegistryCoordinatorOperatorSetParam, minimumStake *big.Int, strategyParams []IStakeRegistryStrategyParams) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.CreateTotalDelegatedStakeQuorum(&_ContractRegistryCoordinator.TransactOpts, operatorSetParams, minimumStake, strategyParams)
}

// CreateTotalDelegatedStakeQuorum is a paid mutator transaction binding the contract method 0x8281ab75.
//
// Solidity: function createTotalDelegatedStakeQuorum((uint32,uint16,uint16) operatorSetParams, uint96 minimumStake, (address,uint96)[] strategyParams) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession) CreateTotalDelegatedStakeQuorum(operatorSetParams IRegistryCoordinatorOperatorSetParam, minimumStake *big.Int, strategyParams []IStakeRegistryStrategyParams) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.CreateTotalDelegatedStakeQuorum(&_ContractRegistryCoordinator.TransactOpts, operatorSetParams, minimumStake, strategyParams)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x9d8e0c23.
//
// Solidity: function deregisterOperator(address operator, uint32[] operatorSetIds) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor) DeregisterOperator(opts *bind.TransactOpts, operator common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.contract.Transact(opts, "deregisterOperator", operator, operatorSetIds)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x9d8e0c23.
//
// Solidity: function deregisterOperator(address operator, uint32[] operatorSetIds) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) DeregisterOperator(operator common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.DeregisterOperator(&_ContractRegistryCoordinator.TransactOpts, operator, operatorSetIds)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x9d8e0c23.
//
// Solidity: function deregisterOperator(address operator, uint32[] operatorSetIds) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession) DeregisterOperator(operator common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.DeregisterOperator(&_ContractRegistryCoordinator.TransactOpts, operator, operatorSetIds)
}

// DeregisterOperator0 is a paid mutator transaction binding the contract method 0xca4f2d97.
//
// Solidity: function deregisterOperator(bytes quorumNumbers) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor) DeregisterOperator0(opts *bind.TransactOpts, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.contract.Transact(opts, "deregisterOperator0", quorumNumbers)
}

// DeregisterOperator0 is a paid mutator transaction binding the contract method 0xca4f2d97.
//
// Solidity: function deregisterOperator(bytes quorumNumbers) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) DeregisterOperator0(quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.DeregisterOperator0(&_ContractRegistryCoordinator.TransactOpts, quorumNumbers)
}

// DeregisterOperator0 is a paid mutator transaction binding the contract method 0xca4f2d97.
//
// Solidity: function deregisterOperator(bytes quorumNumbers) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession) DeregisterOperator0(quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.DeregisterOperator0(&_ContractRegistryCoordinator.TransactOpts, quorumNumbers)
}

// EjectOperator is a paid mutator transaction binding the contract method 0x6e3b17db.
//
// Solidity: function ejectOperator(address operator, bytes quorumNumbers) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor) EjectOperator(opts *bind.TransactOpts, operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.contract.Transact(opts, "ejectOperator", operator, quorumNumbers)
}

// EjectOperator is a paid mutator transaction binding the contract method 0x6e3b17db.
//
// Solidity: function ejectOperator(address operator, bytes quorumNumbers) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) EjectOperator(operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.EjectOperator(&_ContractRegistryCoordinator.TransactOpts, operator, quorumNumbers)
}

// EjectOperator is a paid mutator transaction binding the contract method 0x6e3b17db.
//
// Solidity: function ejectOperator(address operator, bytes quorumNumbers) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession) EjectOperator(operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.EjectOperator(&_ContractRegistryCoordinator.TransactOpts, operator, quorumNumbers)
}

// EnableOperatorSets is a paid mutator transaction binding the contract method 0xee318821.
//
// Solidity: function enableOperatorSets() returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor) EnableOperatorSets(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.contract.Transact(opts, "enableOperatorSets")
}

// EnableOperatorSets is a paid mutator transaction binding the contract method 0xee318821.
//
// Solidity: function enableOperatorSets() returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) EnableOperatorSets() (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.EnableOperatorSets(&_ContractRegistryCoordinator.TransactOpts)
}

// EnableOperatorSets is a paid mutator transaction binding the contract method 0xee318821.
//
// Solidity: function enableOperatorSets() returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession) EnableOperatorSets() (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.EnableOperatorSets(&_ContractRegistryCoordinator.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x7fc3f886.
//
// Solidity: function initialize(address _initialOwner, address _churnApprover, address _ejector, uint256 _initialPausedStatus, (uint32,uint16,uint16)[] _operatorSetParams, uint96[] _minimumStakes, (address,uint96)[][] _strategyParams, uint8[] _stakeTypes, uint32[] _lookAheadPeriods) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor) Initialize(opts *bind.TransactOpts, _initialOwner common.Address, _churnApprover common.Address, _ejector common.Address, _initialPausedStatus *big.Int, _operatorSetParams []IRegistryCoordinatorOperatorSetParam, _minimumStakes []*big.Int, _strategyParams [][]IStakeRegistryStrategyParams, _stakeTypes []uint8, _lookAheadPeriods []uint32) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.contract.Transact(opts, "initialize", _initialOwner, _churnApprover, _ejector, _initialPausedStatus, _operatorSetParams, _minimumStakes, _strategyParams, _stakeTypes, _lookAheadPeriods)
}

// Initialize is a paid mutator transaction binding the contract method 0x7fc3f886.
//
// Solidity: function initialize(address _initialOwner, address _churnApprover, address _ejector, uint256 _initialPausedStatus, (uint32,uint16,uint16)[] _operatorSetParams, uint96[] _minimumStakes, (address,uint96)[][] _strategyParams, uint8[] _stakeTypes, uint32[] _lookAheadPeriods) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) Initialize(_initialOwner common.Address, _churnApprover common.Address, _ejector common.Address, _initialPausedStatus *big.Int, _operatorSetParams []IRegistryCoordinatorOperatorSetParam, _minimumStakes []*big.Int, _strategyParams [][]IStakeRegistryStrategyParams, _stakeTypes []uint8, _lookAheadPeriods []uint32) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.Initialize(&_ContractRegistryCoordinator.TransactOpts, _initialOwner, _churnApprover, _ejector, _initialPausedStatus, _operatorSetParams, _minimumStakes, _strategyParams, _stakeTypes, _lookAheadPeriods)
}

// Initialize is a paid mutator transaction binding the contract method 0x7fc3f886.
//
// Solidity: function initialize(address _initialOwner, address _churnApprover, address _ejector, uint256 _initialPausedStatus, (uint32,uint16,uint16)[] _operatorSetParams, uint96[] _minimumStakes, (address,uint96)[][] _strategyParams, uint8[] _stakeTypes, uint32[] _lookAheadPeriods) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession) Initialize(_initialOwner common.Address, _churnApprover common.Address, _ejector common.Address, _initialPausedStatus *big.Int, _operatorSetParams []IRegistryCoordinatorOperatorSetParam, _minimumStakes []*big.Int, _strategyParams [][]IStakeRegistryStrategyParams, _stakeTypes []uint8, _lookAheadPeriods []uint32) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.Initialize(&_ContractRegistryCoordinator.TransactOpts, _initialOwner, _churnApprover, _ejector, _initialPausedStatus, _operatorSetParams, _minimumStakes, _strategyParams, _stakeTypes, _lookAheadPeriods)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.Pause(&_ContractRegistryCoordinator.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.Pause(&_ContractRegistryCoordinator.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) PauseAll() (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.PauseAll(&_ContractRegistryCoordinator.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.PauseAll(&_ContractRegistryCoordinator.TransactOpts)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xa50857bf.
//
// Solidity: function registerOperator(bytes quorumNumbers, string socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor) RegisterOperator(opts *bind.TransactOpts, quorumNumbers []byte, socket string, params IBLSApkRegistryPubkeyRegistrationParams, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.contract.Transact(opts, "registerOperator", quorumNumbers, socket, params, operatorSignature)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xa50857bf.
//
// Solidity: function registerOperator(bytes quorumNumbers, string socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) RegisterOperator(quorumNumbers []byte, socket string, params IBLSApkRegistryPubkeyRegistrationParams, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.RegisterOperator(&_ContractRegistryCoordinator.TransactOpts, quorumNumbers, socket, params, operatorSignature)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xa50857bf.
//
// Solidity: function registerOperator(bytes quorumNumbers, string socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession) RegisterOperator(quorumNumbers []byte, socket string, params IBLSApkRegistryPubkeyRegistrationParams, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.RegisterOperator(&_ContractRegistryCoordinator.TransactOpts, quorumNumbers, socket, params, operatorSignature)
}

// RegisterOperator0 is a paid mutator transaction binding the contract method 0xadcf73f7.
//
// Solidity: function registerOperator(address operator, uint32[] operatorSetIds, bytes data) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor) RegisterOperator0(opts *bind.TransactOpts, operator common.Address, operatorSetIds []uint32, data []byte) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.contract.Transact(opts, "registerOperator0", operator, operatorSetIds, data)
}

// RegisterOperator0 is a paid mutator transaction binding the contract method 0xadcf73f7.
//
// Solidity: function registerOperator(address operator, uint32[] operatorSetIds, bytes data) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) RegisterOperator0(operator common.Address, operatorSetIds []uint32, data []byte) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.RegisterOperator0(&_ContractRegistryCoordinator.TransactOpts, operator, operatorSetIds, data)
}

// RegisterOperator0 is a paid mutator transaction binding the contract method 0xadcf73f7.
//
// Solidity: function registerOperator(address operator, uint32[] operatorSetIds, bytes data) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession) RegisterOperator0(operator common.Address, operatorSetIds []uint32, data []byte) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.RegisterOperator0(&_ContractRegistryCoordinator.TransactOpts, operator, operatorSetIds, data)
}

// RegisterOperatorWithChurn is a paid mutator transaction binding the contract method 0x9b5d177b.
//
// Solidity: function registerOperatorWithChurn(bytes quorumNumbers, string socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (uint8,address)[] operatorKickParams, (bytes,bytes32,uint256) churnApproverSignature, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor) RegisterOperatorWithChurn(opts *bind.TransactOpts, quorumNumbers []byte, socket string, params IBLSApkRegistryPubkeyRegistrationParams, operatorKickParams []IRegistryCoordinatorOperatorKickParam, churnApproverSignature ISignatureUtilsSignatureWithSaltAndExpiry, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.contract.Transact(opts, "registerOperatorWithChurn", quorumNumbers, socket, params, operatorKickParams, churnApproverSignature, operatorSignature)
}

// RegisterOperatorWithChurn is a paid mutator transaction binding the contract method 0x9b5d177b.
//
// Solidity: function registerOperatorWithChurn(bytes quorumNumbers, string socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (uint8,address)[] operatorKickParams, (bytes,bytes32,uint256) churnApproverSignature, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) RegisterOperatorWithChurn(quorumNumbers []byte, socket string, params IBLSApkRegistryPubkeyRegistrationParams, operatorKickParams []IRegistryCoordinatorOperatorKickParam, churnApproverSignature ISignatureUtilsSignatureWithSaltAndExpiry, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.RegisterOperatorWithChurn(&_ContractRegistryCoordinator.TransactOpts, quorumNumbers, socket, params, operatorKickParams, churnApproverSignature, operatorSignature)
}

// RegisterOperatorWithChurn is a paid mutator transaction binding the contract method 0x9b5d177b.
//
// Solidity: function registerOperatorWithChurn(bytes quorumNumbers, string socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (uint8,address)[] operatorKickParams, (bytes,bytes32,uint256) churnApproverSignature, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession) RegisterOperatorWithChurn(quorumNumbers []byte, socket string, params IBLSApkRegistryPubkeyRegistrationParams, operatorKickParams []IRegistryCoordinatorOperatorKickParam, churnApproverSignature ISignatureUtilsSignatureWithSaltAndExpiry, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.RegisterOperatorWithChurn(&_ContractRegistryCoordinator.TransactOpts, quorumNumbers, socket, params, operatorKickParams, churnApproverSignature, operatorSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.RenounceOwnership(&_ContractRegistryCoordinator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.RenounceOwnership(&_ContractRegistryCoordinator.TransactOpts)
}

// SetChurnApprover is a paid mutator transaction binding the contract method 0x29d1e0c3.
//
// Solidity: function setChurnApprover(address _churnApprover) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor) SetChurnApprover(opts *bind.TransactOpts, _churnApprover common.Address) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.contract.Transact(opts, "setChurnApprover", _churnApprover)
}

// SetChurnApprover is a paid mutator transaction binding the contract method 0x29d1e0c3.
//
// Solidity: function setChurnApprover(address _churnApprover) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) SetChurnApprover(_churnApprover common.Address) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.SetChurnApprover(&_ContractRegistryCoordinator.TransactOpts, _churnApprover)
}

// SetChurnApprover is a paid mutator transaction binding the contract method 0x29d1e0c3.
//
// Solidity: function setChurnApprover(address _churnApprover) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession) SetChurnApprover(_churnApprover common.Address) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.SetChurnApprover(&_ContractRegistryCoordinator.TransactOpts, _churnApprover)
}

// SetEjectionCooldown is a paid mutator transaction binding the contract method 0x0d3f2134.
//
// Solidity: function setEjectionCooldown(uint256 _ejectionCooldown) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor) SetEjectionCooldown(opts *bind.TransactOpts, _ejectionCooldown *big.Int) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.contract.Transact(opts, "setEjectionCooldown", _ejectionCooldown)
}

// SetEjectionCooldown is a paid mutator transaction binding the contract method 0x0d3f2134.
//
// Solidity: function setEjectionCooldown(uint256 _ejectionCooldown) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) SetEjectionCooldown(_ejectionCooldown *big.Int) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.SetEjectionCooldown(&_ContractRegistryCoordinator.TransactOpts, _ejectionCooldown)
}

// SetEjectionCooldown is a paid mutator transaction binding the contract method 0x0d3f2134.
//
// Solidity: function setEjectionCooldown(uint256 _ejectionCooldown) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession) SetEjectionCooldown(_ejectionCooldown *big.Int) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.SetEjectionCooldown(&_ContractRegistryCoordinator.TransactOpts, _ejectionCooldown)
}

// SetEjector is a paid mutator transaction binding the contract method 0x2cdd1e86.
//
// Solidity: function setEjector(address _ejector) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor) SetEjector(opts *bind.TransactOpts, _ejector common.Address) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.contract.Transact(opts, "setEjector", _ejector)
}

// SetEjector is a paid mutator transaction binding the contract method 0x2cdd1e86.
//
// Solidity: function setEjector(address _ejector) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) SetEjector(_ejector common.Address) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.SetEjector(&_ContractRegistryCoordinator.TransactOpts, _ejector)
}

// SetEjector is a paid mutator transaction binding the contract method 0x2cdd1e86.
//
// Solidity: function setEjector(address _ejector) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession) SetEjector(_ejector common.Address) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.SetEjector(&_ContractRegistryCoordinator.TransactOpts, _ejector)
}

// SetOperatorSetParams is a paid mutator transaction binding the contract method 0x5b0b829f.
//
// Solidity: function setOperatorSetParams(uint8 quorumNumber, (uint32,uint16,uint16) operatorSetParams) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor) SetOperatorSetParams(opts *bind.TransactOpts, quorumNumber uint8, operatorSetParams IRegistryCoordinatorOperatorSetParam) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.contract.Transact(opts, "setOperatorSetParams", quorumNumber, operatorSetParams)
}

// SetOperatorSetParams is a paid mutator transaction binding the contract method 0x5b0b829f.
//
// Solidity: function setOperatorSetParams(uint8 quorumNumber, (uint32,uint16,uint16) operatorSetParams) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) SetOperatorSetParams(quorumNumber uint8, operatorSetParams IRegistryCoordinatorOperatorSetParam) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.SetOperatorSetParams(&_ContractRegistryCoordinator.TransactOpts, quorumNumber, operatorSetParams)
}

// SetOperatorSetParams is a paid mutator transaction binding the contract method 0x5b0b829f.
//
// Solidity: function setOperatorSetParams(uint8 quorumNumber, (uint32,uint16,uint16) operatorSetParams) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession) SetOperatorSetParams(quorumNumber uint8, operatorSetParams IRegistryCoordinatorOperatorSetParam) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.SetOperatorSetParams(&_ContractRegistryCoordinator.TransactOpts, quorumNumber, operatorSetParams)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.TransferOwnership(&_ContractRegistryCoordinator.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.TransferOwnership(&_ContractRegistryCoordinator.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.Unpause(&_ContractRegistryCoordinator.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.Unpause(&_ContractRegistryCoordinator.TransactOpts, newPausedStatus)
}

// UpdateOperators is a paid mutator transaction binding the contract method 0x00cf2ab5.
//
// Solidity: function updateOperators(address[] operators) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor) UpdateOperators(opts *bind.TransactOpts, operators []common.Address) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.contract.Transact(opts, "updateOperators", operators)
}

// UpdateOperators is a paid mutator transaction binding the contract method 0x00cf2ab5.
//
// Solidity: function updateOperators(address[] operators) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) UpdateOperators(operators []common.Address) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.UpdateOperators(&_ContractRegistryCoordinator.TransactOpts, operators)
}

// UpdateOperators is a paid mutator transaction binding the contract method 0x00cf2ab5.
//
// Solidity: function updateOperators(address[] operators) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession) UpdateOperators(operators []common.Address) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.UpdateOperators(&_ContractRegistryCoordinator.TransactOpts, operators)
}

// UpdateOperatorsForQuorum is a paid mutator transaction binding the contract method 0x5140a548.
//
// Solidity: function updateOperatorsForQuorum(address[][] operatorsPerQuorum, bytes quorumNumbers) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor) UpdateOperatorsForQuorum(opts *bind.TransactOpts, operatorsPerQuorum [][]common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.contract.Transact(opts, "updateOperatorsForQuorum", operatorsPerQuorum, quorumNumbers)
}

// UpdateOperatorsForQuorum is a paid mutator transaction binding the contract method 0x5140a548.
//
// Solidity: function updateOperatorsForQuorum(address[][] operatorsPerQuorum, bytes quorumNumbers) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) UpdateOperatorsForQuorum(operatorsPerQuorum [][]common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.UpdateOperatorsForQuorum(&_ContractRegistryCoordinator.TransactOpts, operatorsPerQuorum, quorumNumbers)
}

// UpdateOperatorsForQuorum is a paid mutator transaction binding the contract method 0x5140a548.
//
// Solidity: function updateOperatorsForQuorum(address[][] operatorsPerQuorum, bytes quorumNumbers) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession) UpdateOperatorsForQuorum(operatorsPerQuorum [][]common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.UpdateOperatorsForQuorum(&_ContractRegistryCoordinator.TransactOpts, operatorsPerQuorum, quorumNumbers)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x0cf4b767.
//
// Solidity: function updateSocket(string socket) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor) UpdateSocket(opts *bind.TransactOpts, socket string) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.contract.Transact(opts, "updateSocket", socket)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x0cf4b767.
//
// Solidity: function updateSocket(string socket) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorSession) UpdateSocket(socket string) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.UpdateSocket(&_ContractRegistryCoordinator.TransactOpts, socket)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x0cf4b767.
//
// Solidity: function updateSocket(string socket) returns()
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession) UpdateSocket(socket string) (*types.Transaction, error) {
	return _ContractRegistryCoordinator.Contract.UpdateSocket(&_ContractRegistryCoordinator.TransactOpts, socket)
}

// ContractRegistryCoordinatorChurnApproverUpdatedIterator is returned from FilterChurnApproverUpdated and is used to iterate over the raw logs and unpacked data for ChurnApproverUpdated events raised by the ContractRegistryCoordinator contract.
type ContractRegistryCoordinatorChurnApproverUpdatedIterator struct {
	Event *ContractRegistryCoordinatorChurnApproverUpdated // Event containing the contract specifics and raw log

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
func (it *ContractRegistryCoordinatorChurnApproverUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRegistryCoordinatorChurnApproverUpdated)
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
		it.Event = new(ContractRegistryCoordinatorChurnApproverUpdated)
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
func (it *ContractRegistryCoordinatorChurnApproverUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRegistryCoordinatorChurnApproverUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRegistryCoordinatorChurnApproverUpdated represents a ChurnApproverUpdated event raised by the ContractRegistryCoordinator contract.
type ContractRegistryCoordinatorChurnApproverUpdated struct {
	PrevChurnApprover common.Address
	NewChurnApprover  common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterChurnApproverUpdated is a free log retrieval operation binding the contract event 0x315457d8a8fe60f04af17c16e2f5a5e1db612b31648e58030360759ef8f3528c.
//
// Solidity: event ChurnApproverUpdated(address prevChurnApprover, address newChurnApprover)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) FilterChurnApproverUpdated(opts *bind.FilterOpts) (*ContractRegistryCoordinatorChurnApproverUpdatedIterator, error) {

	logs, sub, err := _ContractRegistryCoordinator.contract.FilterLogs(opts, "ChurnApproverUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractRegistryCoordinatorChurnApproverUpdatedIterator{contract: _ContractRegistryCoordinator.contract, event: "ChurnApproverUpdated", logs: logs, sub: sub}, nil
}

// WatchChurnApproverUpdated is a free log subscription operation binding the contract event 0x315457d8a8fe60f04af17c16e2f5a5e1db612b31648e58030360759ef8f3528c.
//
// Solidity: event ChurnApproverUpdated(address prevChurnApprover, address newChurnApprover)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) WatchChurnApproverUpdated(opts *bind.WatchOpts, sink chan<- *ContractRegistryCoordinatorChurnApproverUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractRegistryCoordinator.contract.WatchLogs(opts, "ChurnApproverUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRegistryCoordinatorChurnApproverUpdated)
				if err := _ContractRegistryCoordinator.contract.UnpackLog(event, "ChurnApproverUpdated", log); err != nil {
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

// ParseChurnApproverUpdated is a log parse operation binding the contract event 0x315457d8a8fe60f04af17c16e2f5a5e1db612b31648e58030360759ef8f3528c.
//
// Solidity: event ChurnApproverUpdated(address prevChurnApprover, address newChurnApprover)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) ParseChurnApproverUpdated(log types.Log) (*ContractRegistryCoordinatorChurnApproverUpdated, error) {
	event := new(ContractRegistryCoordinatorChurnApproverUpdated)
	if err := _ContractRegistryCoordinator.contract.UnpackLog(event, "ChurnApproverUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRegistryCoordinatorEjectorUpdatedIterator is returned from FilterEjectorUpdated and is used to iterate over the raw logs and unpacked data for EjectorUpdated events raised by the ContractRegistryCoordinator contract.
type ContractRegistryCoordinatorEjectorUpdatedIterator struct {
	Event *ContractRegistryCoordinatorEjectorUpdated // Event containing the contract specifics and raw log

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
func (it *ContractRegistryCoordinatorEjectorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRegistryCoordinatorEjectorUpdated)
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
		it.Event = new(ContractRegistryCoordinatorEjectorUpdated)
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
func (it *ContractRegistryCoordinatorEjectorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRegistryCoordinatorEjectorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRegistryCoordinatorEjectorUpdated represents a EjectorUpdated event raised by the ContractRegistryCoordinator contract.
type ContractRegistryCoordinatorEjectorUpdated struct {
	PrevEjector common.Address
	NewEjector  common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEjectorUpdated is a free log retrieval operation binding the contract event 0x8f30ab09f43a6c157d7fce7e0a13c003042c1c95e8a72e7a146a21c0caa24dc9.
//
// Solidity: event EjectorUpdated(address prevEjector, address newEjector)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) FilterEjectorUpdated(opts *bind.FilterOpts) (*ContractRegistryCoordinatorEjectorUpdatedIterator, error) {

	logs, sub, err := _ContractRegistryCoordinator.contract.FilterLogs(opts, "EjectorUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractRegistryCoordinatorEjectorUpdatedIterator{contract: _ContractRegistryCoordinator.contract, event: "EjectorUpdated", logs: logs, sub: sub}, nil
}

// WatchEjectorUpdated is a free log subscription operation binding the contract event 0x8f30ab09f43a6c157d7fce7e0a13c003042c1c95e8a72e7a146a21c0caa24dc9.
//
// Solidity: event EjectorUpdated(address prevEjector, address newEjector)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) WatchEjectorUpdated(opts *bind.WatchOpts, sink chan<- *ContractRegistryCoordinatorEjectorUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractRegistryCoordinator.contract.WatchLogs(opts, "EjectorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRegistryCoordinatorEjectorUpdated)
				if err := _ContractRegistryCoordinator.contract.UnpackLog(event, "EjectorUpdated", log); err != nil {
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

// ParseEjectorUpdated is a log parse operation binding the contract event 0x8f30ab09f43a6c157d7fce7e0a13c003042c1c95e8a72e7a146a21c0caa24dc9.
//
// Solidity: event EjectorUpdated(address prevEjector, address newEjector)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) ParseEjectorUpdated(log types.Log) (*ContractRegistryCoordinatorEjectorUpdated, error) {
	event := new(ContractRegistryCoordinatorEjectorUpdated)
	if err := _ContractRegistryCoordinator.contract.UnpackLog(event, "EjectorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRegistryCoordinatorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractRegistryCoordinator contract.
type ContractRegistryCoordinatorInitializedIterator struct {
	Event *ContractRegistryCoordinatorInitialized // Event containing the contract specifics and raw log

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
func (it *ContractRegistryCoordinatorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRegistryCoordinatorInitialized)
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
		it.Event = new(ContractRegistryCoordinatorInitialized)
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
func (it *ContractRegistryCoordinatorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRegistryCoordinatorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRegistryCoordinatorInitialized represents a Initialized event raised by the ContractRegistryCoordinator contract.
type ContractRegistryCoordinatorInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractRegistryCoordinatorInitializedIterator, error) {

	logs, sub, err := _ContractRegistryCoordinator.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractRegistryCoordinatorInitializedIterator{contract: _ContractRegistryCoordinator.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractRegistryCoordinatorInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractRegistryCoordinator.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRegistryCoordinatorInitialized)
				if err := _ContractRegistryCoordinator.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) ParseInitialized(log types.Log) (*ContractRegistryCoordinatorInitialized, error) {
	event := new(ContractRegistryCoordinatorInitialized)
	if err := _ContractRegistryCoordinator.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRegistryCoordinatorOperatorDeregisteredIterator is returned from FilterOperatorDeregistered and is used to iterate over the raw logs and unpacked data for OperatorDeregistered events raised by the ContractRegistryCoordinator contract.
type ContractRegistryCoordinatorOperatorDeregisteredIterator struct {
	Event *ContractRegistryCoordinatorOperatorDeregistered // Event containing the contract specifics and raw log

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
func (it *ContractRegistryCoordinatorOperatorDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRegistryCoordinatorOperatorDeregistered)
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
		it.Event = new(ContractRegistryCoordinatorOperatorDeregistered)
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
func (it *ContractRegistryCoordinatorOperatorDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRegistryCoordinatorOperatorDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRegistryCoordinatorOperatorDeregistered represents a OperatorDeregistered event raised by the ContractRegistryCoordinator contract.
type ContractRegistryCoordinatorOperatorDeregistered struct {
	Operator   common.Address
	OperatorId [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorDeregistered is a free log retrieval operation binding the contract event 0x396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e4.
//
// Solidity: event OperatorDeregistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) FilterOperatorDeregistered(opts *bind.FilterOpts, operator []common.Address, operatorId [][32]byte) (*ContractRegistryCoordinatorOperatorDeregisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractRegistryCoordinator.contract.FilterLogs(opts, "OperatorDeregistered", operatorRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryCoordinatorOperatorDeregisteredIterator{contract: _ContractRegistryCoordinator.contract, event: "OperatorDeregistered", logs: logs, sub: sub}, nil
}

// WatchOperatorDeregistered is a free log subscription operation binding the contract event 0x396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e4.
//
// Solidity: event OperatorDeregistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) WatchOperatorDeregistered(opts *bind.WatchOpts, sink chan<- *ContractRegistryCoordinatorOperatorDeregistered, operator []common.Address, operatorId [][32]byte) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractRegistryCoordinator.contract.WatchLogs(opts, "OperatorDeregistered", operatorRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRegistryCoordinatorOperatorDeregistered)
				if err := _ContractRegistryCoordinator.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
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

// ParseOperatorDeregistered is a log parse operation binding the contract event 0x396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e4.
//
// Solidity: event OperatorDeregistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) ParseOperatorDeregistered(log types.Log) (*ContractRegistryCoordinatorOperatorDeregistered, error) {
	event := new(ContractRegistryCoordinatorOperatorDeregistered)
	if err := _ContractRegistryCoordinator.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRegistryCoordinatorOperatorRegisteredIterator is returned from FilterOperatorRegistered and is used to iterate over the raw logs and unpacked data for OperatorRegistered events raised by the ContractRegistryCoordinator contract.
type ContractRegistryCoordinatorOperatorRegisteredIterator struct {
	Event *ContractRegistryCoordinatorOperatorRegistered // Event containing the contract specifics and raw log

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
func (it *ContractRegistryCoordinatorOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRegistryCoordinatorOperatorRegistered)
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
		it.Event = new(ContractRegistryCoordinatorOperatorRegistered)
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
func (it *ContractRegistryCoordinatorOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRegistryCoordinatorOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRegistryCoordinatorOperatorRegistered represents a OperatorRegistered event raised by the ContractRegistryCoordinator contract.
type ContractRegistryCoordinatorOperatorRegistered struct {
	Operator   common.Address
	OperatorId [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0xe8e68cef1c3a761ed7be7e8463a375f27f7bc335e51824223cacce636ec5c3fe.
//
// Solidity: event OperatorRegistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address, operatorId [][32]byte) (*ContractRegistryCoordinatorOperatorRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractRegistryCoordinator.contract.FilterLogs(opts, "OperatorRegistered", operatorRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryCoordinatorOperatorRegisteredIterator{contract: _ContractRegistryCoordinator.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0xe8e68cef1c3a761ed7be7e8463a375f27f7bc335e51824223cacce636ec5c3fe.
//
// Solidity: event OperatorRegistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *ContractRegistryCoordinatorOperatorRegistered, operator []common.Address, operatorId [][32]byte) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractRegistryCoordinator.contract.WatchLogs(opts, "OperatorRegistered", operatorRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRegistryCoordinatorOperatorRegistered)
				if err := _ContractRegistryCoordinator.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
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

// ParseOperatorRegistered is a log parse operation binding the contract event 0xe8e68cef1c3a761ed7be7e8463a375f27f7bc335e51824223cacce636ec5c3fe.
//
// Solidity: event OperatorRegistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) ParseOperatorRegistered(log types.Log) (*ContractRegistryCoordinatorOperatorRegistered, error) {
	event := new(ContractRegistryCoordinatorOperatorRegistered)
	if err := _ContractRegistryCoordinator.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRegistryCoordinatorOperatorSetParamsUpdatedIterator is returned from FilterOperatorSetParamsUpdated and is used to iterate over the raw logs and unpacked data for OperatorSetParamsUpdated events raised by the ContractRegistryCoordinator contract.
type ContractRegistryCoordinatorOperatorSetParamsUpdatedIterator struct {
	Event *ContractRegistryCoordinatorOperatorSetParamsUpdated // Event containing the contract specifics and raw log

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
func (it *ContractRegistryCoordinatorOperatorSetParamsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRegistryCoordinatorOperatorSetParamsUpdated)
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
		it.Event = new(ContractRegistryCoordinatorOperatorSetParamsUpdated)
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
func (it *ContractRegistryCoordinatorOperatorSetParamsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRegistryCoordinatorOperatorSetParamsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRegistryCoordinatorOperatorSetParamsUpdated represents a OperatorSetParamsUpdated event raised by the ContractRegistryCoordinator contract.
type ContractRegistryCoordinatorOperatorSetParamsUpdated struct {
	QuorumNumber      uint8
	OperatorSetParams IRegistryCoordinatorOperatorSetParam
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetParamsUpdated is a free log retrieval operation binding the contract event 0x3ee6fe8d54610244c3e9d3c066ae4aee997884aa28f10616ae821925401318ac.
//
// Solidity: event OperatorSetParamsUpdated(uint8 indexed quorumNumber, (uint32,uint16,uint16) operatorSetParams)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) FilterOperatorSetParamsUpdated(opts *bind.FilterOpts, quorumNumber []uint8) (*ContractRegistryCoordinatorOperatorSetParamsUpdatedIterator, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _ContractRegistryCoordinator.contract.FilterLogs(opts, "OperatorSetParamsUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryCoordinatorOperatorSetParamsUpdatedIterator{contract: _ContractRegistryCoordinator.contract, event: "OperatorSetParamsUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorSetParamsUpdated is a free log subscription operation binding the contract event 0x3ee6fe8d54610244c3e9d3c066ae4aee997884aa28f10616ae821925401318ac.
//
// Solidity: event OperatorSetParamsUpdated(uint8 indexed quorumNumber, (uint32,uint16,uint16) operatorSetParams)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) WatchOperatorSetParamsUpdated(opts *bind.WatchOpts, sink chan<- *ContractRegistryCoordinatorOperatorSetParamsUpdated, quorumNumber []uint8) (event.Subscription, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _ContractRegistryCoordinator.contract.WatchLogs(opts, "OperatorSetParamsUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRegistryCoordinatorOperatorSetParamsUpdated)
				if err := _ContractRegistryCoordinator.contract.UnpackLog(event, "OperatorSetParamsUpdated", log); err != nil {
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

// ParseOperatorSetParamsUpdated is a log parse operation binding the contract event 0x3ee6fe8d54610244c3e9d3c066ae4aee997884aa28f10616ae821925401318ac.
//
// Solidity: event OperatorSetParamsUpdated(uint8 indexed quorumNumber, (uint32,uint16,uint16) operatorSetParams)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) ParseOperatorSetParamsUpdated(log types.Log) (*ContractRegistryCoordinatorOperatorSetParamsUpdated, error) {
	event := new(ContractRegistryCoordinatorOperatorSetParamsUpdated)
	if err := _ContractRegistryCoordinator.contract.UnpackLog(event, "OperatorSetParamsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRegistryCoordinatorOperatorSocketUpdateIterator is returned from FilterOperatorSocketUpdate and is used to iterate over the raw logs and unpacked data for OperatorSocketUpdate events raised by the ContractRegistryCoordinator contract.
type ContractRegistryCoordinatorOperatorSocketUpdateIterator struct {
	Event *ContractRegistryCoordinatorOperatorSocketUpdate // Event containing the contract specifics and raw log

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
func (it *ContractRegistryCoordinatorOperatorSocketUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRegistryCoordinatorOperatorSocketUpdate)
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
		it.Event = new(ContractRegistryCoordinatorOperatorSocketUpdate)
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
func (it *ContractRegistryCoordinatorOperatorSocketUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRegistryCoordinatorOperatorSocketUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRegistryCoordinatorOperatorSocketUpdate represents a OperatorSocketUpdate event raised by the ContractRegistryCoordinator contract.
type ContractRegistryCoordinatorOperatorSocketUpdate struct {
	OperatorId [32]byte
	Socket     string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorSocketUpdate is a free log retrieval operation binding the contract event 0xec2963ab21c1e50e1e582aa542af2e4bf7bf38e6e1403c27b42e1c5d6e621eaa.
//
// Solidity: event OperatorSocketUpdate(bytes32 indexed operatorId, string socket)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) FilterOperatorSocketUpdate(opts *bind.FilterOpts, operatorId [][32]byte) (*ContractRegistryCoordinatorOperatorSocketUpdateIterator, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractRegistryCoordinator.contract.FilterLogs(opts, "OperatorSocketUpdate", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryCoordinatorOperatorSocketUpdateIterator{contract: _ContractRegistryCoordinator.contract, event: "OperatorSocketUpdate", logs: logs, sub: sub}, nil
}

// WatchOperatorSocketUpdate is a free log subscription operation binding the contract event 0xec2963ab21c1e50e1e582aa542af2e4bf7bf38e6e1403c27b42e1c5d6e621eaa.
//
// Solidity: event OperatorSocketUpdate(bytes32 indexed operatorId, string socket)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) WatchOperatorSocketUpdate(opts *bind.WatchOpts, sink chan<- *ContractRegistryCoordinatorOperatorSocketUpdate, operatorId [][32]byte) (event.Subscription, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractRegistryCoordinator.contract.WatchLogs(opts, "OperatorSocketUpdate", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRegistryCoordinatorOperatorSocketUpdate)
				if err := _ContractRegistryCoordinator.contract.UnpackLog(event, "OperatorSocketUpdate", log); err != nil {
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

// ParseOperatorSocketUpdate is a log parse operation binding the contract event 0xec2963ab21c1e50e1e582aa542af2e4bf7bf38e6e1403c27b42e1c5d6e621eaa.
//
// Solidity: event OperatorSocketUpdate(bytes32 indexed operatorId, string socket)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) ParseOperatorSocketUpdate(log types.Log) (*ContractRegistryCoordinatorOperatorSocketUpdate, error) {
	event := new(ContractRegistryCoordinatorOperatorSocketUpdate)
	if err := _ContractRegistryCoordinator.contract.UnpackLog(event, "OperatorSocketUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRegistryCoordinatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractRegistryCoordinator contract.
type ContractRegistryCoordinatorOwnershipTransferredIterator struct {
	Event *ContractRegistryCoordinatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractRegistryCoordinatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRegistryCoordinatorOwnershipTransferred)
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
		it.Event = new(ContractRegistryCoordinatorOwnershipTransferred)
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
func (it *ContractRegistryCoordinatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRegistryCoordinatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRegistryCoordinatorOwnershipTransferred represents a OwnershipTransferred event raised by the ContractRegistryCoordinator contract.
type ContractRegistryCoordinatorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractRegistryCoordinatorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractRegistryCoordinator.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryCoordinatorOwnershipTransferredIterator{contract: _ContractRegistryCoordinator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractRegistryCoordinatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractRegistryCoordinator.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRegistryCoordinatorOwnershipTransferred)
				if err := _ContractRegistryCoordinator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) ParseOwnershipTransferred(log types.Log) (*ContractRegistryCoordinatorOwnershipTransferred, error) {
	event := new(ContractRegistryCoordinatorOwnershipTransferred)
	if err := _ContractRegistryCoordinator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRegistryCoordinatorPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractRegistryCoordinator contract.
type ContractRegistryCoordinatorPausedIterator struct {
	Event *ContractRegistryCoordinatorPaused // Event containing the contract specifics and raw log

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
func (it *ContractRegistryCoordinatorPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRegistryCoordinatorPaused)
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
		it.Event = new(ContractRegistryCoordinatorPaused)
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
func (it *ContractRegistryCoordinatorPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRegistryCoordinatorPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRegistryCoordinatorPaused represents a Paused event raised by the ContractRegistryCoordinator contract.
type ContractRegistryCoordinatorPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractRegistryCoordinatorPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractRegistryCoordinator.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryCoordinatorPausedIterator{contract: _ContractRegistryCoordinator.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractRegistryCoordinatorPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractRegistryCoordinator.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRegistryCoordinatorPaused)
				if err := _ContractRegistryCoordinator.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) ParsePaused(log types.Log) (*ContractRegistryCoordinatorPaused, error) {
	event := new(ContractRegistryCoordinatorPaused)
	if err := _ContractRegistryCoordinator.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRegistryCoordinatorQuorumBlockNumberUpdatedIterator is returned from FilterQuorumBlockNumberUpdated and is used to iterate over the raw logs and unpacked data for QuorumBlockNumberUpdated events raised by the ContractRegistryCoordinator contract.
type ContractRegistryCoordinatorQuorumBlockNumberUpdatedIterator struct {
	Event *ContractRegistryCoordinatorQuorumBlockNumberUpdated // Event containing the contract specifics and raw log

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
func (it *ContractRegistryCoordinatorQuorumBlockNumberUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRegistryCoordinatorQuorumBlockNumberUpdated)
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
		it.Event = new(ContractRegistryCoordinatorQuorumBlockNumberUpdated)
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
func (it *ContractRegistryCoordinatorQuorumBlockNumberUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRegistryCoordinatorQuorumBlockNumberUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRegistryCoordinatorQuorumBlockNumberUpdated represents a QuorumBlockNumberUpdated event raised by the ContractRegistryCoordinator contract.
type ContractRegistryCoordinatorQuorumBlockNumberUpdated struct {
	QuorumNumber uint8
	Blocknumber  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterQuorumBlockNumberUpdated is a free log retrieval operation binding the contract event 0x46077d55330763f16269fd75e5761663f4192d2791747c0189b16ad31db07db4.
//
// Solidity: event QuorumBlockNumberUpdated(uint8 indexed quorumNumber, uint256 blocknumber)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) FilterQuorumBlockNumberUpdated(opts *bind.FilterOpts, quorumNumber []uint8) (*ContractRegistryCoordinatorQuorumBlockNumberUpdatedIterator, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _ContractRegistryCoordinator.contract.FilterLogs(opts, "QuorumBlockNumberUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryCoordinatorQuorumBlockNumberUpdatedIterator{contract: _ContractRegistryCoordinator.contract, event: "QuorumBlockNumberUpdated", logs: logs, sub: sub}, nil
}

// WatchQuorumBlockNumberUpdated is a free log subscription operation binding the contract event 0x46077d55330763f16269fd75e5761663f4192d2791747c0189b16ad31db07db4.
//
// Solidity: event QuorumBlockNumberUpdated(uint8 indexed quorumNumber, uint256 blocknumber)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) WatchQuorumBlockNumberUpdated(opts *bind.WatchOpts, sink chan<- *ContractRegistryCoordinatorQuorumBlockNumberUpdated, quorumNumber []uint8) (event.Subscription, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _ContractRegistryCoordinator.contract.WatchLogs(opts, "QuorumBlockNumberUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRegistryCoordinatorQuorumBlockNumberUpdated)
				if err := _ContractRegistryCoordinator.contract.UnpackLog(event, "QuorumBlockNumberUpdated", log); err != nil {
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

// ParseQuorumBlockNumberUpdated is a log parse operation binding the contract event 0x46077d55330763f16269fd75e5761663f4192d2791747c0189b16ad31db07db4.
//
// Solidity: event QuorumBlockNumberUpdated(uint8 indexed quorumNumber, uint256 blocknumber)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) ParseQuorumBlockNumberUpdated(log types.Log) (*ContractRegistryCoordinatorQuorumBlockNumberUpdated, error) {
	event := new(ContractRegistryCoordinatorQuorumBlockNumberUpdated)
	if err := _ContractRegistryCoordinator.contract.UnpackLog(event, "QuorumBlockNumberUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRegistryCoordinatorUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractRegistryCoordinator contract.
type ContractRegistryCoordinatorUnpausedIterator struct {
	Event *ContractRegistryCoordinatorUnpaused // Event containing the contract specifics and raw log

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
func (it *ContractRegistryCoordinatorUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRegistryCoordinatorUnpaused)
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
		it.Event = new(ContractRegistryCoordinatorUnpaused)
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
func (it *ContractRegistryCoordinatorUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRegistryCoordinatorUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRegistryCoordinatorUnpaused represents a Unpaused event raised by the ContractRegistryCoordinator contract.
type ContractRegistryCoordinatorUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractRegistryCoordinatorUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractRegistryCoordinator.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryCoordinatorUnpausedIterator{contract: _ContractRegistryCoordinator.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractRegistryCoordinatorUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractRegistryCoordinator.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRegistryCoordinatorUnpaused)
				if err := _ContractRegistryCoordinator.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer) ParseUnpaused(log types.Log) (*ContractRegistryCoordinatorUnpaused, error) {
	event := new(ContractRegistryCoordinatorUnpaused)
	if err := _ContractRegistryCoordinator.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
