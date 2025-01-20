// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractAllocationManager

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

// IAllocationManagerTypesAllocateParams is an auto generated low-level Go binding around an user-defined struct.
type IAllocationManagerTypesAllocateParams struct {
	OperatorSet   OperatorSet
	Strategies    []common.Address
	NewMagnitudes []uint64
}

// IAllocationManagerTypesAllocation is an auto generated low-level Go binding around an user-defined struct.
type IAllocationManagerTypesAllocation struct {
	CurrentMagnitude uint64
	PendingDiff      *big.Int
	EffectBlock      uint32
}

// IAllocationManagerTypesCreateSetParams is an auto generated low-level Go binding around an user-defined struct.
type IAllocationManagerTypesCreateSetParams struct {
	OperatorSetId uint32
	Strategies    []common.Address
}

// IAllocationManagerTypesDeregisterParams is an auto generated low-level Go binding around an user-defined struct.
type IAllocationManagerTypesDeregisterParams struct {
	Operator       common.Address
	Avs            common.Address
	OperatorSetIds []uint32
}

// IAllocationManagerTypesRegisterParams is an auto generated low-level Go binding around an user-defined struct.
type IAllocationManagerTypesRegisterParams struct {
	Avs            common.Address
	OperatorSetIds []uint32
	Data           []byte
}

// IAllocationManagerTypesSlashingParams is an auto generated low-level Go binding around an user-defined struct.
type IAllocationManagerTypesSlashingParams struct {
	Operator      common.Address
	OperatorSetId uint32
	Strategies    []common.Address
	WadsToSlash   []*big.Int
	Description   string
}

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// ContractAllocationManagerMetaData contains all meta data concerning the ContractAllocationManager contract.
var ContractAllocationManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIDelegationManager\",\"name\":\"_delegation\",\"type\":\"address\"},{\"internalType\":\"contractIPauserRegistry\",\"name\":\"_pauserRegistry\",\"type\":\"address\"},{\"internalType\":\"contractIPermissionController\",\"name\":\"_permissionController\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_DEALLOCATION_DELAY\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_ALLOCATION_CONFIGURATION_DELAY\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyMemberOfSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CurrentlyPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Empty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InputAddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InputArrayLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientMagnitude\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCaller\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNewPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOperatorSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPermissions\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSnapshotOrdering\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWadToSlash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ModificationAlreadyPending\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotMemberOfSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyUnpauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorNotSlashable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SameMagnitude\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StrategiesMustBeInAscendingOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StrategyAlreadyInOperatorSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StrategyNotInOperatorSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UninitializedAllocationDelay\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"avs\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"}],\"name\":\"AVSMetadataURIUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"avs\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIAVSRegistrar\",\"name\":\"registrar\",\"type\":\"address\"}],\"name\":\"AVSRegistrarSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"delay\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"effectBlock\",\"type\":\"uint32\"}],\"name\":\"AllocationDelaySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"avs\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structOperatorSet\",\"name\":\"operatorSet\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"magnitude\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"effectBlock\",\"type\":\"uint32\"}],\"name\":\"AllocationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"encumberedMagnitude\",\"type\":\"uint64\"}],\"name\":\"EncumberedMagnitudeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"maxMagnitude\",\"type\":\"uint64\"}],\"name\":\"MaxMagnitudeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"avs\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structOperatorSet\",\"name\":\"operatorSet\",\"type\":\"tuple\"}],\"name\":\"OperatorAddedToOperatorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"avs\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structOperatorSet\",\"name\":\"operatorSet\",\"type\":\"tuple\"}],\"name\":\"OperatorRemovedFromOperatorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"avs\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structOperatorSet\",\"name\":\"operatorSet\",\"type\":\"tuple\"}],\"name\":\"OperatorSetCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"avs\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structOperatorSet\",\"name\":\"operatorSet\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"contractIStrategy[]\",\"name\":\"strategies\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"wadSlashed\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"OperatorSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"avs\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structOperatorSet\",\"name\":\"operatorSet\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"}],\"name\":\"StrategyAddedToOperatorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"avs\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structOperatorSet\",\"name\":\"operatorSet\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"}],\"name\":\"StrategyRemovedFromOperatorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ALLOCATION_CONFIGURATION_DELAY\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEALLOCATION_DELAY\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"avs\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"operatorSetId\",\"type\":\"uint32\"},{\"internalType\":\"contractIStrategy[]\",\"name\":\"strategies\",\"type\":\"address[]\"}],\"name\":\"addStrategiesToOperatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"contractIStrategy[]\",\"name\":\"strategies\",\"type\":\"address[]\"},{\"internalType\":\"uint16[]\",\"name\":\"numToClear\",\"type\":\"uint16[]\"}],\"name\":\"clearDeallocationQueue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"avs\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"operatorSetId\",\"type\":\"uint32\"},{\"internalType\":\"contractIStrategy[]\",\"name\":\"strategies\",\"type\":\"address[]\"}],\"internalType\":\"structIAllocationManagerTypes.CreateSetParams[]\",\"name\":\"params\",\"type\":\"tuple[]\"}],\"name\":\"createOperatorSets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegation\",\"outputs\":[{\"internalType\":\"contractIDelegationManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"avs\",\"type\":\"address\"},{\"internalType\":\"uint32[]\",\"name\":\"operatorSetIds\",\"type\":\"uint32[]\"}],\"internalType\":\"structIAllocationManagerTypes.DeregisterParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"deregisterFromOperatorSets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"}],\"name\":\"encumberedMagnitude\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"avs\",\"type\":\"address\"}],\"name\":\"getAVSRegistrar\",\"outputs\":[{\"internalType\":\"contractIAVSRegistrar\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"}],\"name\":\"getAllocatableMagnitude\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getAllocatedSets\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"avs\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"}],\"internalType\":\"structOperatorSet[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"avs\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"}],\"internalType\":\"structOperatorSet\",\"name\":\"operatorSet\",\"type\":\"tuple\"}],\"name\":\"getAllocatedStrategies\",\"outputs\":[{\"internalType\":\"contractIStrategy[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"avs\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"}],\"internalType\":\"structOperatorSet\",\"name\":\"operatorSet\",\"type\":\"tuple\"},{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"}],\"name\":\"getAllocation\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"currentMagnitude\",\"type\":\"uint64\"},{\"internalType\":\"int128\",\"name\":\"pendingDiff\",\"type\":\"int128\"},{\"internalType\":\"uint32\",\"name\":\"effectBlock\",\"type\":\"uint32\"}],\"internalType\":\"structIAllocationManagerTypes.Allocation\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getAllocationDelay\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"operators\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"avs\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"}],\"internalType\":\"structOperatorSet\",\"name\":\"operatorSet\",\"type\":\"tuple\"},{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"}],\"name\":\"getAllocations\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"currentMagnitude\",\"type\":\"uint64\"},{\"internalType\":\"int128\",\"name\":\"pendingDiff\",\"type\":\"int128\"},{\"internalType\":\"uint32\",\"name\":\"effectBlock\",\"type\":\"uint32\"}],\"internalType\":\"structIAllocationManagerTypes.Allocation[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"}],\"name\":\"getMaxMagnitude\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"operators\",\"type\":\"address[]\"},{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"}],\"name\":\"getMaxMagnitudes\",\"outputs\":[{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"contractIStrategy[]\",\"name\":\"strategies\",\"type\":\"address[]\"}],\"name\":\"getMaxMagnitudes\",\"outputs\":[{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"contractIStrategy[]\",\"name\":\"strategies\",\"type\":\"address[]\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"}],\"name\":\"getMaxMagnitudesAtBlock\",\"outputs\":[{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"avs\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"}],\"internalType\":\"structOperatorSet\",\"name\":\"operatorSet\",\"type\":\"tuple\"}],\"name\":\"getMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"avs\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"}],\"internalType\":\"structOperatorSet\",\"name\":\"operatorSet\",\"type\":\"tuple\"}],\"name\":\"getMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"avs\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"}],\"internalType\":\"structOperatorSet\",\"name\":\"operatorSet\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"operators\",\"type\":\"address[]\"},{\"internalType\":\"contractIStrategy[]\",\"name\":\"strategies\",\"type\":\"address[]\"},{\"internalType\":\"uint32\",\"name\":\"futureBlock\",\"type\":\"uint32\"}],\"name\":\"getMinimumSlashableStake\",\"outputs\":[{\"internalType\":\"uint256[][]\",\"name\":\"slashableStake\",\"type\":\"uint256[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"avs\",\"type\":\"address\"}],\"name\":\"getOperatorSetCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getRegisteredSets\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"avs\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"}],\"internalType\":\"structOperatorSet[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"avs\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"}],\"internalType\":\"structOperatorSet\",\"name\":\"operatorSet\",\"type\":\"tuple\"}],\"name\":\"getStrategiesInOperatorSet\",\"outputs\":[{\"internalType\":\"contractIStrategy[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"}],\"name\":\"getStrategyAllocations\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"avs\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"}],\"internalType\":\"structOperatorSet[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"currentMagnitude\",\"type\":\"uint64\"},{\"internalType\":\"int128\",\"name\":\"pendingDiff\",\"type\":\"int128\"},{\"internalType\":\"uint32\",\"name\":\"effectBlock\",\"type\":\"uint32\"}],\"internalType\":\"structIAllocationManagerTypes.Allocation[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialPausedStatus\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"avs\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"}],\"internalType\":\"structOperatorSet\",\"name\":\"operatorSet\",\"type\":\"tuple\"}],\"name\":\"isMemberOfOperatorSet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"avs\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"}],\"internalType\":\"structOperatorSet\",\"name\":\"operatorSet\",\"type\":\"tuple\"}],\"name\":\"isOperatorSet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"avs\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"}],\"internalType\":\"structOperatorSet\",\"name\":\"operatorSet\",\"type\":\"tuple\"},{\"internalType\":\"contractIStrategy[]\",\"name\":\"strategies\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"newMagnitudes\",\"type\":\"uint64[]\"}],\"internalType\":\"structIAllocationManagerTypes.AllocateParams[]\",\"name\":\"params\",\"type\":\"tuple[]\"}],\"name\":\"modifyAllocations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"}],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauserRegistry\",\"outputs\":[{\"internalType\":\"contractIPauserRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"permissionController\",\"outputs\":[{\"internalType\":\"contractIPermissionController\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"avs\",\"type\":\"address\"},{\"internalType\":\"uint32[]\",\"name\":\"operatorSetIds\",\"type\":\"uint32[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIAllocationManagerTypes.RegisterParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"registerForOperatorSets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"avs\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"operatorSetId\",\"type\":\"uint32\"},{\"internalType\":\"contractIStrategy[]\",\"name\":\"strategies\",\"type\":\"address[]\"}],\"name\":\"removeStrategiesFromOperatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"avs\",\"type\":\"address\"},{\"internalType\":\"contractIAVSRegistrar\",\"name\":\"registrar\",\"type\":\"address\"}],\"name\":\"setAVSRegistrar\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"delay\",\"type\":\"uint32\"}],\"name\":\"setAllocationDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"avs\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"operatorSetId\",\"type\":\"uint32\"},{\"internalType\":\"contractIStrategy[]\",\"name\":\"strategies\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"wadsToSlash\",\"type\":\"uint256[]\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structIAllocationManagerTypes.SlashingParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"slashOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"avs\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"}],\"name\":\"updateAVSMetadataURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b5060043610610281575f3560e01c80636e875dba11610156578063a984eb3a116100ca578063c221d8ae11610084578063c221d8ae146106bc578063cd6dc687146106cf578063d3d96ff4146106e2578063df5cf723146106f5578063f2fde38b1461071c578063fabc1cbc1461072f575f5ffd5b8063a984eb3a1461060e578063adc2e3d914610641578063b2447af714610654578063b66bd98914610667578063b9fbaed11461067a578063ba1a84e5146106a9575f5ffd5b80638ce648541161011b5780638ce64854146105915780638da5cb5b146105b157806394d7d00c146105c2578063952899ee146105d5578063a9333ec8146105e8578063a9821821146105fb575f5ffd5b80636e875dba14610515578063715018a61461052857806379ae50cd146105305780637bc1ef6114610543578063886f11951461056a575f5ffd5b80634657e26a116101f8578063595c6a67116101b2578063595c6a67146104875780635ac86ab71461048f5780635c975abb146104b2578063670d3ba2146104c45780636cfb4481146104d75780636e3492b514610502575f5ffd5b80634657e26a146103f45780634a10ffe51461041b5780634b5046ef1461043b57806350feea201461044e578063547afb871461046157806356c483e614610474575f5ffd5b80632981eb77116102495780632981eb77146103195780632bab2c4a14610355578063304c10cd1461037557806336352057146103a057806340120dab146103b35780634177a87c146103d4575f5ffd5b806310e1b9b814610285578063136439dd146102ae57806315fe5028146102c3578063260dc758146102e3578063261f84e014610306575b5f5ffd5b610298610293366004614728565b610742565b6040516102a5919061476f565b60405180910390f35b6102c16102bc3660046147a2565b61077d565b005b6102d66102d13660046147b9565b610852565b6040516102a59190614837565b6102f66102f1366004614849565b610969565b60405190151581526020016102a5565b6102c16103143660046148a3565b6109a0565b6103407f000000000000000000000000000000000000000000000000000000000000003281565b60405163ffffffff90911681526020016102a5565b610368610363366004614988565b610c43565b6040516102a59190614a3c565b6103886103833660046147b9565b610f30565b6040516001600160a01b0390911681526020016102a5565b6102c16103ae366004614a9f565b610f5f565b6103c66103c1366004614af1565b611771565b6040516102a5929190614b7e565b6103e76103e2366004614849565b6118ec565b6040516102a59190614bdb565b6103887f000000000000000000000000598cb226b591155f767da17afe7a2241a68c5c1081565b61042e610429366004614bed565b611910565b6040516102a59190614c30565b6102c1610449366004614c7b565b6119b8565b6102c161045c366004614cfb565b611a72565b61042e61046f366004614d59565b611bd0565b6102c1610482366004614da5565b611c78565b6102c1611d7d565b6102f661049d366004614dd8565b606654600160ff9092169190911b9081161490565b6066545b6040519081526020016102a5565b6102f66104d2366004614df8565b611e2c565b6104ea6104e5366004614af1565b611e3d565b6040516001600160401b0390911681526020016102a5565b6102c1610510366004614e39565b611faa565b6103e7610523366004614849565b61237a565b6102c161238b565b6102d661053e3660046147b9565b61239c565b6103407f000000000000000000000000000000000000000000000000000000000000004b81565b6103887f00000000000000000000000041dbe7bbaca97d986fcf6f5203b98ec02412ec1d81565b6105a461059f366004614e6a565b612476565b6040516102a59190614ead565b6033546001600160a01b0316610388565b61042e6105d0366004614ebf565b61253b565b6102c16105e3366004614f1a565b612627565b6104ea6105f6366004614af1565b612aee565b6102c16106093660046150c3565b612b1d565b6104ea61061c366004614af1565b60a260209081525f92835260408084209091529082529020546001600160401b031681565b6102c161064f366004615141565b612b8d565b6104b6610662366004614849565b612edc565b6102c1610675366004614cfb565b612efe565b61068d6106883660046147b9565b613058565b60408051921515835263ffffffff9091166020830152016102a5565b6104b66106b73660046147b9565b6130f2565b6103e76106ca366004614df8565b613112565b6102c16106dd366004615183565b613143565b6102c16106f0366004614af1565b613260565b6103887f000000000000000000000000a44151489861fe9e3055d95adc98fbd462b948e781565b6102c161072a3660046147b9565b6132ff565b6102c161073d3660046147a2565b613378565b604080516060810182525f80825260208201819052918101829052906107718561076b8661348e565b856134f1565b925050505b9392505050565b60405163237dfb4760e11b81523360048201527f00000000000000000000000041dbe7bbaca97d986fcf6f5203b98ec02412ec1d6001600160a01b0316906346fbf68e90602401602060405180830381865afa1580156107df573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061080391906151ad565b61082057604051631d77d47760e21b815260040160405180910390fd5b60665481811681146108455760405163c61dca5d60e01b815260040160405180910390fd5b61084e8261365d565b5050565b6001600160a01b0381165f908152609d60205260408120606091906108769061369a565b90505f816001600160401b038111156108915761089161464c565b6040519080825280602002602001820160405280156108d557816020015b604080518082019091525f80825260208201528152602001906001900390816108af5790505b5090505f5b82811015610961576001600160a01b0385165f908152609d6020526040902061093c9061090790836136a3565b604080518082019091525f80825260208201525060408051808201909152606082901c815263ffffffff909116602082015290565b82828151811061094e5761094e6151cc565b60209081029190910101526001016108da565b509392505050565b60208082015182516001600160a01b03165f90815260989092526040822061099a9163ffffffff908116906136ae16565b92915050565b826109aa816136c5565b6109c75760405163932d94f760e01b815260040160405180910390fd5b5f5b82811015610c3c575f6040518060400160405280876001600160a01b031681526020018686858181106109fe576109fe6151cc565b9050602002810190610a1091906151e0565b610a1e9060208101906151fe565b63ffffffff168152509050610a68816020015163ffffffff1660985f896001600160a01b03166001600160a01b031681526020019081526020015f2061376f90919063ffffffff16565b610a8557604051631fb1705560e21b815260040160405180910390fd5b7f31629285ead2335ae0933f86ed2ae63321f7af77b4e6eaabc42c057880977e6c6040518060400160405280886001600160a01b03168152602001836020015163ffffffff16815250604051610adb9190615217565b60405180910390a15f610aed8261348e565b90505f5b868685818110610b0357610b036151cc565b9050602002810190610b1591906151e0565b610b23906020810190615225565b9050811015610c3157610b99878786818110610b4157610b416151cc565b9050602002810190610b5391906151e0565b610b61906020810190615225565b83818110610b7157610b716151cc565b9050602002016020810190610b8691906147b9565b5f8481526099602052604090209061377a565b507f7ab260fe0af193db5f4986770d831bda4ea46099dc817e8b6716dcae8af8e88b83888887818110610bce57610bce6151cc565b9050602002810190610be091906151e0565b610bee906020810190615225565b84818110610bfe57610bfe6151cc565b9050602002016020810190610c1391906147b9565b604051610c2192919061526a565b60405180910390a1600101610af1565b5050506001016109c9565b5050505050565b606083516001600160401b03811115610c5e57610c5e61464c565b604051908082528060200260200182016040528015610c9157816020015b6060815260200190600190039081610c7c5790505b5090505f7f000000000000000000000000a44151489861fe9e3055d95adc98fbd462b948e76001600160a01b031663f0e0e67686866040518363ffffffff1660e01b8152600401610ce3929190615290565b5f60405180830381865afa158015610cfd573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610d2491908101906152b4565b90505f5b8551811015610f26575f868281518110610d4457610d446151cc565b6020026020010151905085516001600160401b03811115610d6757610d6761464c565b604051908082528060200260200182016040528015610d90578160200160208202803683370190505b50848381518110610da357610da36151cc565b60209081029190910101525f5b8651811015610f1c575f878281518110610dcc57610dcc6151cc565b6020908102919091018101516001600160a01b038086165f90815260a1845260408082209284168252919093528220909250610e079061378e565b9050806001600160401b03165f03610e20575050610f14565b5f610e2c858d85610742565b90508863ffffffff16816040015163ffffffff1611158015610e5457505f8160200151600f0b125b15610e7657610e6a815f015182602001516137a1565b6001600160401b031681525b80515f90610e91906001600160401b039081169085166137b5565b9050610ed881898981518110610ea957610ea96151cc565b60200260200101518781518110610ec257610ec26151cc565b60200260200101516137c990919063ffffffff16565b898881518110610eea57610eea6151cc565b60200260200101518681518110610f0357610f036151cc565b602002602001018181525050505050505b600101610db0565b5050600101610d28565b5050949350505050565b6001600160a01b038082165f908152609760205260408120549091168015610f585780610776565b5090919050565b606654600190600290811603610f885760405163840a48d560e01b815260040160405180910390fd5b82610f92816136c5565b610faf5760405163932d94f760e01b815260040160405180910390fd5b5f6040518060400160405280866001600160a01b03168152602001856020016020810190610fdd91906151fe565b63ffffffff16905290505f610ffe610ff860208701876147b9565b836137dd565b905061100d6060860186615225565b905061101c6040870187615225565b90501461103c576040516343714afd60e01b815260040160405180910390fd5b60208083015183516001600160a01b03165f9081526098909252604090912061106e9163ffffffff908116906136ae16565b61108b57604051631fb1705560e21b815260040160405180910390fd5b806110a95760405163ebbff49760e01b815260040160405180910390fd5b5f6110b76040870187615225565b90506001600160401b038111156110d0576110d061464c565b6040519080825280602002602001820160405280156110f9578160200160208202803683370190505b5090505f5b61110b6040880188615225565b90508110156117025780158061119e57506111296040880188615225565b6111346001846153d4565b818110611143576111436151cc565b905060200201602081019061115891906147b9565b6001600160a01b031661116e6040890189615225565b8381811061117e5761117e6151cc565b905060200201602081019061119391906147b9565b6001600160a01b0316115b6111bb57604051639f1c805360e01b815260040160405180910390fd5b6111c86060880188615225565b828181106111d8576111d86151cc565b905060200201355f1080156112185750670de0b6b3a76400006111fe6060890189615225565b8381811061120e5761120e6151cc565b9050602002013511155b61123557604051631353603160e01b815260040160405180910390fd5b6112916112456040890189615225565b83818110611255576112556151cc565b905060200201602081019061126a91906147b9565b60995f6112768861348e565b81526020019081526020015f2061385390919063ffffffff16565b6112ae576040516331bc342760e11b815260040160405180910390fd5b5f806113006112c060208b018b6147b9565b6112c98861348e565b6112d660408d018d615225565b878181106112e6576112e66151cc565b90506020020160208101906112fb91906147b9565b6134f1565b805191935091506001600160401b03165f0361131d5750506116fa565b5f61135861132e60608c018c615225565b8681811061133e5761133e6151cc565b85516001600160401b031692602090910201359050613874565b83519091506113736001600160401b038084169083166137b5565b868681518110611385576113856151cc565b60200260200101818152505081835f018181516113a291906153e7565b6001600160401b03169052508351829085906113bf9083906153e7565b6001600160401b03169052506020840180518391906113df9083906153e7565b6001600160401b031690525060208301515f600f9190910b12156114fa575f61144261140e60608e018e615225565b8881811061141e5761141e6151cc565b90506020020135856020015161143390615406565b6001600160801b031690613874565b9050806001600160401b03168460200181815161145f919061542a565b600f0b9052507f1487af5418c47ee5ea45ef4a93398668120890774a9e13487e61e9dc3baf76dd61149360208e018e6147b9565b8a8e80604001906114a49190615225565b8a8181106114b4576114b46151cc565b90506020020160208101906114c991906147b9565b6114da885f015189602001516137a1565b88604001516040516114f0959493929190615457565b60405180910390a1505b61154c61150a60208d018d6147b9565b6115138a61348e565b61152060408f018f615225565b89818110611530576115306151cc565b905060200201602081019061154591906147b9565b878761388a565b7f1487af5418c47ee5ea45ef4a93398668120890774a9e13487e61e9dc3baf76dd61157a60208d018d6147b9565b8961158860408f018f615225565b89818110611598576115986151cc565b90506020020160208101906115ad91906147b9565b86516040516115c194939291904390615457565b60405180910390a16116126115d960208d018d6147b9565b6115e660408e018e615225565b888181106115f6576115f66151cc565b905060200201602081019061160b91906147b9565b8651613aca565b6001600160a01b037f000000000000000000000000a44151489861fe9e3055d95adc98fbd462b948e71663ee74937f61164e60208e018e6147b9565b61165b60408f018f615225565b8981811061166b5761166b6151cc565b905060200201602081019061168091906147b9565b875160405160e085901b6001600160e01b03191681526001600160a01b0393841660048201529290911660248301526001600160401b0380861660448401521660648201526084015f604051808303815f87803b1580156116df575f5ffd5b505af11580156116f1573d5f5f3e3d5ffd5b50505050505050505b6001016110fe565b507f80969ad29428d6797ee7aad084f9e4a42a82fc506dcd2ca3b6fb431f85ccebe561173160208801886147b9565b8461173f60408a018a615225565b8561174d60808d018d6154a8565b6040516117609796959493929190615512565b60405180910390a150505050505050565b6001600160a01b0382165f908152609d6020526040812060609182916117969061369a565b90505f816001600160401b038111156117b1576117b161464c565b6040519080825280602002602001820160405280156117f557816020015b604080518082019091525f80825260208201528152602001906001900390816117cf5790505b5090505f826001600160401b038111156118115761181161464c565b60405190808252806020026020018201604052801561185a57816020015b604080516060810182525f80825260208083018290529282015282525f1990920191018161182f5790505b5090505f5b838110156118dd576001600160a01b0388165f908152609d6020526040812061188c9061090790846136a3565b9050808483815181106118a1576118a16151cc565b60200260200101819052506118b789828a610742565b8383815181106118c9576118c96151cc565b60209081029190910101525060010161185f565b509093509150505b9250929050565b60605f61077660995f6118fe8661348e565b81526020019081526020015f20613b4c565b60605f83516001600160401b0381111561192c5761192c61464c565b604051908082528060200260200182016040528015611955578160200160208202803683370190505b5090505f5b845181101561096157611986858281518110611978576119786151cc565b602002602001015185612aee565b828281518110611998576119986151cc565b6001600160401b039092166020928302919091019091015260010161195a565b6066545f906001908116036119e05760405163840a48d560e01b815260040160405180910390fd5b838214611a00576040516343714afd60e01b815260040160405180910390fd5b5f5b84811015611a6957611a6187878784818110611a2057611a206151cc565b9050602002016020810190611a3591906147b9565b868685818110611a4757611a476151cc565b9050602002016020810190611a5c91906155a8565b613b58565b600101611a02565b50505050505050565b83611a7c816136c5565b611a995760405163932d94f760e01b815260040160405180910390fd5b604080518082019091526001600160a01b038616815263ffffffff851660208201525f611ac58261348e565b9050611b06826020015163ffffffff1660985f8a6001600160a01b03166001600160a01b031681526020019081526020015f206136ae90919063ffffffff16565b611b2357604051631fb1705560e21b815260040160405180910390fd5b5f5b84811015611bc657611b42868683818110610b7157610b716151cc565b611b5f5760405163585cfb2f60e01b815260040160405180910390fd5b7f7ab260fe0af193db5f4986770d831bda4ea46099dc817e8b6716dcae8af8e88b83878784818110611b9357611b936151cc565b9050602002016020810190611ba891906147b9565b604051611bb692919061526a565b60405180910390a1600101611b25565b5050505050505050565b60605f82516001600160401b03811115611bec57611bec61464c565b604051908082528060200260200182016040528015611c15578160200160208202803683370190505b5090505f5b835181101561096157611c4685858381518110611c3957611c396151cc565b6020026020010151612aee565b828281518110611c5857611c586151cc565b6001600160401b0390921660209283029190910190910152600101611c1a565b336001600160a01b037f000000000000000000000000a44151489861fe9e3055d95adc98fbd462b948e71614611d7357611cb1826136c5565b611cce576040516348f5c3ed60e01b815260040160405180910390fd5b6040516336b87bd760e11b81526001600160a01b0383811660048301527f000000000000000000000000a44151489861fe9e3055d95adc98fbd462b948e71690636d70f7ae90602401602060405180830381865afa158015611d32573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611d5691906151ad565b611d735760405163ccea9e6f60e01b815260040160405180910390fd5b61084e8282613c5c565b60405163237dfb4760e11b81523360048201527f00000000000000000000000041dbe7bbaca97d986fcf6f5203b98ec02412ec1d6001600160a01b0316906346fbf68e90602401602060405180830381865afa158015611ddf573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611e0391906151ad565b611e2057604051631d77d47760e21b815260040160405180910390fd5b611e2a5f1961365d565b565b5f61077683609a5f6112768661348e565b6001600160a01b038281165f81815260a2602090815260408083209486168084529482528083205493835260a38252808320948352939052918220546001600160401b0390911690600f81810b600160801b909204900b03825b81811015611f67576001600160a01b038087165f90815260a3602090815260408083209389168352929052908120611ecf9083613dfd565b6001600160a01b038881165f90815260a0602090815260408083208584528252808320938b16835292815290829020825160608101845290546001600160401b0381168252600160401b8104600f0b92820192909252600160c01b90910463ffffffff16918101829052919250431015611f4a575050611f67565b611f588582602001516137a1565b94505050806001019050611e97565b506001600160a01b038086165f90815260a1602090815260408083209388168352929052208290611f979061378e565b611fa191906153e7565b95945050505050565b606654600290600490811603611fd35760405163840a48d560e01b815260040160405180910390fd5b611fe8611fe360208401846147b9565b6136c5565b806120015750612001611fe360408401602085016147b9565b61201e576040516348f5c3ed60e01b815260040160405180910390fd5b5f5b61202d6040840184615225565b90508110156122ef575f604051806040016040528085602001602081019061205591906147b9565b6001600160a01b031681526020016120706040870187615225565b85818110612080576120806151cc565b905060200201602081019061209591906151fe565b63ffffffff1681525090506120e2816020015163ffffffff1660985f8760200160208101906120c491906147b9565b6001600160a01b0316815260208101919091526040015f20906136ae565b6120ff57604051631fb1705560e21b815260040160405180910390fd5b609e5f61210f60208701876147b9565b6001600160a01b03166001600160a01b031681526020019081526020015f205f6121388361348e565b815260208101919091526040015f205460ff16612168576040516325131d4f60e01b815260040160405180910390fd5b6121a26121748261348e565b609c5f61218460208901896147b9565b6001600160a01b0316815260208101919091526040015f2090613e6c565b506121da6121b360208601866147b9565b609a5f6121bf8561348e565b81526020019081526020015f20613e7790919063ffffffff16565b506121e860208501856147b9565b6001600160a01b03167fad34c3070be1dffbcaa499d000ba2b8d9848aefcac3059df245dd95c4ece14fe826040516122209190615217565b60405180910390a2604080518082019091525f8152602081016122637f0000000000000000000000000000000000000000000000000000000000000032436155c9565b63ffffffff169052609e5f61227b60208801886147b9565b6001600160a01b03166001600160a01b031681526020019081526020015f205f6122a48461348e565b81526020808201929092526040015f2082518154939092015163ffffffff166101000264ffffffff00199215159290921664ffffffffff199093169290921717905550600101612020565b5061230361038360408401602085016147b9565b6001600160a01b0316639d8e0c2361231e60208501856147b9565b61232b6040860186615225565b6040518463ffffffff1660e01b81526004016123499392919061561e565b5f604051808303815f87803b158015612360575f5ffd5b505af1925050508015612371575060015b1561084e575050565b606061099a609a5f6118fe8561348e565b612393613e8b565b611e2a5f613ee5565b6001600160a01b0381165f908152609c60205260408120606091906123c09061369a565b90505f816001600160401b038111156123db576123db61464c565b60405190808252806020026020018201604052801561241f57816020015b604080518082019091525f80825260208201528152602001906001900390816123f95790505b5090505f5b82811015610961576001600160a01b0385165f908152609c602052604090206124519061090790836136a3565b828281518110612463576124636151cc565b6020908102919091010152600101612424565b60605f84516001600160401b038111156124925761249261464c565b6040519080825280602002602001820160405280156124db57816020015b604080516060810182525f80825260208083018290529282015282525f199092019101816124b05790505b5090505f5b85518110156125325761250d8682815181106124fe576124fe6151cc565b60200260200101518686610742565b82828151811061251f5761251f6151cc565b60209081029190910101526001016124e0565b50949350505050565b60605f83516001600160401b038111156125575761255761464c565b604051908082528060200260200182016040528015612580578160200160208202803683370190505b5090505f5b8451811015612532576001600160a01b0386165f90815260a16020526040812086516125f5928792918990869081106125c0576125c06151cc565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020015f20613f3690919063ffffffff16565b828281518110612607576126076151cc565b6001600160401b0390921660209283029190910190910152600101612585565b6066545f9060019081160361264f5760405163840a48d560e01b815260040160405180910390fd5b612658836136c5565b612675576040516348f5c3ed60e01b815260040160405180910390fd5b5f5f5f61268186613058565b91509150816126a35760405163fa55fc8160e01b815260040160405180910390fd5b91505f90505b8351811015610c3c578381815181106126c4576126c46151cc565b602002602001015160400151518482815181106126e3576126e36151cc565b602002602001015160200151511461270e576040516343714afd60e01b815260040160405180910390fd5b5f848281518110612721576127216151cc565b602090810291909101810151518082015181516001600160a01b03165f908152609890935260409092209092506127619163ffffffff908116906136ae16565b61277e57604051631fb1705560e21b815260040160405180910390fd5b5f61278987836137dd565b90505f5b86848151811061279f5761279f6151cc565b60200260200101516020015151811015612ae3575f8785815181106127c6576127c66151cc565b60200260200101516020015182815181106127e3576127e36151cc565b602002602001015190506127fa898261ffff613b58565b5f5f6128098b61076b8861348e565b915091508060200151600f0b5f1461283457604051630d8fcbe360e41b815260040160405180910390fd5b5f61284187858489613f4a565b9050612886825f01518c8a8151811061285c5761285c6151cc565b6020026020010151604001518781518110612879576128796151cc565b6020026020010151613f80565b600f0b602083018190525f036128af57604051634606179360e11b815260040160405180910390fd5b5f8260200151600f0b12156129f3578015612975576129306128d08861348e565b6001600160a01b03808f165f90815260a360209081526040808320938a16835292905220908154600160801b90819004600f0b5f818152600180860160205260409091209390935583546001600160801b03908116939091011602179055565b61295a7f0000000000000000000000000000000000000000000000000000000000000032436155c9565b6129659060016155c9565b63ffffffff166040830152612a60565b612987836020015183602001516137a1565b6001600160401b031660208401528a518b90899081106129a9576129a96151cc565b60200260200101516040015185815181106129c6576129c66151cc565b6020908102919091018101516001600160401b031683525f9083015263ffffffff43166040830152612a60565b5f8260200151600f0b1315612a6057612a14836020015183602001516137a1565b6001600160401b039081166020850181905284519091161015612a4a57604051636c9be0bf60e01b815260040160405180910390fd5b612a5489436155c9565b63ffffffff1660408301525b612a758c612a6d8961348e565b86868661388a565b7f1487af5418c47ee5ea45ef4a93398668120890774a9e13487e61e9dc3baf76dd8c612aa36109078a61348e565b86612ab5865f015187602001516137a1565b8660400151604051612acb959493929190615457565b60405180910390a150506001909201915061278d9050565b5050506001016126a9565b6001600160a01b038083165f90815260a16020908152604080832093851683529290529081206107769061378e565b82612b27816136c5565b612b445760405163932d94f760e01b815260040160405180910390fd5b836001600160a01b03167fa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c9437138484604051612b7f929190615642565b60405180910390a250505050565b606654600290600490811603612bb65760405163840a48d560e01b815260040160405180910390fd5b82612bc0816136c5565b612bdd5760405163932d94f760e01b815260040160405180910390fd5b6040516336b87bd760e11b81526001600160a01b0385811660048301527f000000000000000000000000a44151489861fe9e3055d95adc98fbd462b948e71690636d70f7ae90602401602060405180830381865afa158015612c41573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612c6591906151ad565b612c825760405163ccea9e6f60e01b815260040160405180910390fd5b5f5b612c916020850185615225565b9050811015612e5957604080518082019091525f9080612cb460208801886147b9565b6001600160a01b03168152602001868060200190612cd29190615225565b85818110612ce257612ce26151cc565b9050602002016020810190612cf791906151fe565b63ffffffff90811690915260208083015183516001600160a01b03165f90815260989092526040909120929350612d339291908116906136ae16565b612d5057604051631fb1705560e21b815260040160405180910390fd5b612d5a86826137dd565b15612d7857604051636c6c6e2760e11b815260040160405180910390fd5b612da1612d848261348e565b6001600160a01b0388165f908152609c602052604090209061376f565b50612dcd86609a5f612db28561348e565b81526020019081526020015f2061377a90919063ffffffff16565b50856001600160a01b03167f43232edf9071753d2321e5fa7e018363ee248e5f2142e6c08edd3265bfb4895e82604051612e079190615217565b60405180910390a26001600160a01b0386165f908152609e60205260408120600191612e328461348e565b815260208101919091526040015f20805460ff191691151591909117905550600101612c84565b50612e6a61038360208501856147b9565b6001600160a01b031663adcf73f785612e866020870187615225565b612e9360408901896154a8565b6040518663ffffffff1660e01b8152600401612eb3959493929190615655565b5f604051808303815f87803b158015612eca575f5ffd5b505af1158015611bc6573d5f5f3e3d5ffd5b5f61099a609a5f612eec8561348e565b81526020019081526020015f2061369a565b83612f08816136c5565b612f255760405163932d94f760e01b815260040160405180910390fd5b6040805180820182526001600160a01b03871680825263ffffffff80881660208085018290525f93845260989052939091209192612f6492916136ae16565b612f8157604051631fb1705560e21b815260040160405180910390fd5b5f612f8b8261348e565b90505f5b84811015611bc657612fd4868683818110612fac57612fac6151cc565b9050602002016020810190612fc191906147b9565b5f84815260996020526040902090613e77565b612ff1576040516331bc342760e11b815260040160405180910390fd5b7f7b4b073d80dcac55a11177d8459ad9f664ceeb91f71f27167bb14f8152a7eeee83878784818110613025576130256151cc565b905060200201602081019061303a91906147b9565b60405161304892919061526a565b60405180910390a1600101612f8f565b6001600160a01b0381165f908152609b602090815260408083208151608081018352905463ffffffff80821680845260ff600160201b8404161515958401869052650100000000008304821694840194909452600160481b9091041660608201819052849391929190158015906130d95750826060015163ffffffff164310155b156130e8575050604081015160015b9590945092505050565b6001600160a01b0381165f90815260986020526040812061099a9061369a565b6001600160a01b0382165f908152609f602052604081206060919061313b90826118fe8661348e565b949350505050565b5f54610100900460ff161580801561316157505f54600160ff909116105b8061317a5750303b15801561317a57505f5460ff166001145b6131e25760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b5f805460ff191660011790558015613203575f805461ff0019166101001790555b61320c8261365d565b61321583613ee5565b801561325b575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498906020015b60405180910390a15b505050565b8161326a816136c5565b6132875760405163932d94f760e01b815260040160405180910390fd5b6001600160a01b038381165f90815260976020526040902080546001600160a01b0319169184169190911790557f2ae945c40c44dc0ec263f95609c3fdc6952e0aefa22d6374e44f2c997acedf85836132df81610f30565b604080516001600160a01b03938416815292909116602083015201613252565b613307613e8b565b6001600160a01b03811661336c5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016131d9565b61337581613ee5565b50565b7f00000000000000000000000041dbe7bbaca97d986fcf6f5203b98ec02412ec1d6001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156133d4573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906133f89190615698565b6001600160a01b0316336001600160a01b0316146134295760405163794821ff60e01b815260040160405180910390fd5b606654801982198116146134505760405163c61dca5d60e01b815260040160405180910390fd5b606682905560405182815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9060200160405180910390a25050565b5f815f0151826020015163ffffffff166040516020016134d992919060609290921b6bffffffffffffffffffffffff1916825260a01b6001600160a01b031916601482015260200190565b60405160208183030381529060405261099a906156b3565b6040805180820182525f80825260208083018290528351606081018552828152808201839052808501839052845180860186526001600160a01b03898116855260a18452868520908816855290925293822092939281906135519061378e565b6001600160401b0390811682526001600160a01b038981165f81815260a260209081526040808320948c168084529482528083205486169682019690965291815260a082528481208b8252825284812092815291815290839020835160608101855290549283168152600160401b8304600f0b91810191909152600160c01b90910463ffffffff169181018290529192504310156135f3579092509050613655565b613604815f015182602001516137a1565b6001600160401b0316815260208101515f600f9190910b121561364257613633826020015182602001516137a1565b6001600160401b031660208301525b5f60408201819052602082015290925090505b935093915050565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a250565b5f61099a825490565b5f6107768383613f97565b5f8181526001830160205260408120541515610776565b604051631beb2b9760e31b81526001600160a01b0382811660048301523360248301523060448301525f80356001600160e01b0319166064840152917f000000000000000000000000598cb226b591155f767da17afe7a2241a68c5c109091169063df595cb8906084016020604051808303815f875af115801561374b573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061099a91906151ad565b5f6107768383613fbd565b5f610776836001600160a01b038416613fbd565b5f61099a82670de0b6b3a7640000614009565b5f610776826001600160401b03851661542a565b5f61077683670de0b6b3a76400008461404d565b5f6107768383670de0b6b3a764000061404d565b6001600160a01b0382165f908152609e602052604081208190816138008561348e565b815260208082019290925260409081015f2081518083019092525460ff8116151580835261010090910463ffffffff169282019290925291508061313b57506020015163ffffffff164311159392505050565b6001600160a01b0381165f9081526001830160205260408120541515610776565b5f6107768383670de0b6b3a76400006001614132565b6020808301516001600160a01b038088165f90815260a284526040808220928816825291909352909120546001600160401b0390811691161461395057602082810180516001600160a01b038881165f81815260a286526040808220938a1680835293875290819020805467ffffffffffffffff19166001600160401b0395861617905593518451918252948101919091529216908201527facf9095feb3a370c9cf692421c69ef320d4db5c66e6a7d29c7694eb02364fc559060600160405180910390a15b6001600160a01b038086165f90815260a060209081526040808320888452825280832093871683529281529082902083518154928501519385015163ffffffff16600160c01b0263ffffffff60c01b196001600160801b038616600160401b026001600160c01b03199095166001600160401b03909316929092179390931716919091179055600f0b15613a32576001600160a01b0385165f908152609f602090815260408083208784529091529020613a0a908461377a565b506001600160a01b0385165f908152609d60205260409020613a2c908561376f565b50610c3c565b80516001600160401b03165f03610c3c576001600160a01b0385165f908152609f602090815260408083208784529091529020613a6f9084613e77565b506001600160a01b0385165f908152609f602090815260408083208784529091529020613a9b9061369a565b5f03610c3c576001600160a01b0385165f908152609d60205260409020613ac29085613e6c565b505050505050565b6001600160a01b038084165f90815260a160209081526040808320938616835292905220613af990438361418b565b604080516001600160a01b038086168252841660208201526001600160401b038316918101919091527f1c6458079a41077d003c11faf9bf097e693bd67979e4e6500bac7b29db779b5c90606001613252565b60605f6107768361419f565b6001600160a01b038381165f90815260a360209081526040808320938616835292905290812054600f81810b600160801b909204900b035b5f81118015613ba257508261ffff1682105b15610c3c576001600160a01b038086165f90815260a3602090815260408083209388168352929052908120613bd6906141f8565b90505f5f613be58884896134f1565b91509150806040015163ffffffff16431015613c0357505050610c3c565b613c10888489858561388a565b6001600160a01b038089165f90815260a360209081526040808320938b16835292905220613c3d9061424a565b50613c47856156d6565b9450613c52846156ee565b9350505050613b90565b6001600160a01b0382165f908152609b60209081526040918290208251608081018452905463ffffffff808216835260ff600160201b830416151593830193909352650100000000008104831693820193909352600160481b909204166060820181905215801590613cd85750806060015163ffffffff164310155b15613cf257604081015163ffffffff168152600160208201525b63ffffffff82166040820152613d287f000000000000000000000000000000000000000000000000000000000000004b436155c9565b63ffffffff90811660608381019182526001600160a01b0386165f818152609b602090815260409182902087518154838a0151858b01519851928a1664ffffffffff1990921691909117600160201b91151591909102176cffffffffffffffff0000000000191665010000000000978916979097026cffffffff000000000000000000191696909617600160481b968816968702179055815192835294871694820194909452928301919091527f4e85751d6331506c6c62335f207eb31f12a61e570f34f5c17640308785c6d4db9101613252565b5f5f613e1f613e0b846142c7565b8554613e1a9190600f0b615703565b614334565b8454909150600160801b9004600f90810b9082900b12613e5257604051632d0483c560e21b815260040160405180910390fd5b600f0b5f9081526001939093016020525050604090205490565b5f610776838361439d565b5f610776836001600160a01b03841661439d565b6033546001600160a01b03163314611e2a5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016131d9565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f6107768383670de0b6b3a7640000614480565b5f613f5b8460995f6112768961348e565b8015613f645750815b8015611fa157505090516001600160401b031615159392505050565b5f6107766001600160401b0380851690841661572a565b5f825f018281548110613fac57613fac6151cc565b905f5260205f200154905092915050565b5f81815260018301602052604081205461400257508154600181810184555f84815260208082209093018490558454848252828601909352604090209190915561099a565b505f61099a565b81545f9080156140455761402f846140226001846153d4565b5f91825260209091200190565b54600160201b90046001600160e01b031661313b565b509092915050565b5f80805f19858709858702925082811083820303915050805f036140845783828161407a5761407a615757565b0492505050610776565b8084116140cb5760405162461bcd60e51b81526020600482015260156024820152744d6174683a206d756c446976206f766572666c6f7760581b60448201526064016131d9565b5f8486880960026001871981018816978890046003810283188082028403028082028403028082028403028082028403028082028403029081029092039091025f889003889004909101858311909403939093029303949094049190911702949350505050565b5f5f61413f86868661404d565b905060018360028111156141555761415561576b565b14801561417157505f848061416c5761416c615757565b868809115b15611fa15761418160018261577f565b9695505050505050565b61325b83836001600160401b0384166144c8565b6060815f018054806020026020016040519081016040528092919081815260200182805480156141ec57602002820191905f5260205f20905b8154815260200190600101908083116141d8575b50505050509050919050565b5f6142128254600f81810b600160801b909204900b131590565b1561423057604051631ed9509560e11b815260040160405180910390fd5b508054600f0b5f9081526001909101602052604090205490565b5f6142648254600f81810b600160801b909204900b131590565b1561428257604051631ed9509560e11b815260040160405180910390fd5b508054600f0b5f818152600180840160205260408220805492905583546fffffffffffffffffffffffffffffffff191692016001600160801b03169190911790915590565b5f6001600160ff1b038211156143305760405162461bcd60e51b815260206004820152602860248201527f53616665436173743a2076616c756520646f65736e27742066697420696e2061604482015267371034b73a191a9b60c11b60648201526084016131d9565b5090565b80600f81900b81146143985760405162461bcd60e51b815260206004820152602760248201527f53616665436173743a2076616c756520646f65736e27742066697420696e20316044820152663238206269747360c81b60648201526084016131d9565b919050565b5f8181526001830160205260408120548015614477575f6143bf6001836153d4565b85549091505f906143d2906001906153d4565b9050818114614431575f865f0182815481106143f0576143f06151cc565b905f5260205f200154905080875f018481548110614410576144106151cc565b5f918252602080832090910192909255918252600188019052604090208390555b855486908061444257614442615792565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f90556001935050505061099a565b5f91505061099a565b82545f9081614491868683856145cb565b905080156144be576144a8866140226001846153d4565b54600160201b90046001600160e01b0316610771565b5091949350505050565b8254801561457e575f6144e0856140226001856153d4565b60408051808201909152905463ffffffff808216808452600160201b9092046001600160e01b0316602084015291925090851610156145325760405163151b8e3f60e11b815260040160405180910390fd5b805163ffffffff80861691160361457c5782614553866140226001866153d4565b80546001600160e01b0392909216600160201b0263ffffffff9092169190911790555050505050565b505b506040805180820190915263ffffffff92831681526001600160e01b03918216602080830191825285546001810187555f968752952091519051909216600160201b029190921617910155565b5f5b81831015610961575f6145e0848461461e565b5f8781526020902090915063ffffffff86169082015463ffffffff16111561460a57809250614618565b61461581600161577f565b93505b506145cd565b5f61462c60028484186157a6565b6107769084841661577f565b6001600160a01b0381168114613375575f5ffd5b634e487b7160e01b5f52604160045260245ffd5b604051606081016001600160401b03811182821017156146825761468261464c565b60405290565b604051601f8201601f191681016001600160401b03811182821017156146b0576146b061464c565b604052919050565b803563ffffffff81168114614398575f5ffd5b5f604082840312156146db575f5ffd5b604080519081016001600160401b03811182821017156146fd576146fd61464c565b604052905080823561470e81614638565b815261471c602084016146b8565b60208201525092915050565b5f5f5f6080848603121561473a575f5ffd5b833561474581614638565b925061475485602086016146cb565b9150606084013561476481614638565b809150509250925092565b81516001600160401b03168152602080830151600f0b9082015260408083015163ffffffff16908201526060810161099a565b5f602082840312156147b2575f5ffd5b5035919050565b5f602082840312156147c9575f5ffd5b813561077681614638565b80516001600160a01b0316825260209081015163ffffffff16910152565b5f8151808452602084019350602083015f5b8281101561482d576148178683516147d4565b6040959095019460209190910190600101614804565b5093949350505050565b602081525f61077660208301846147f2565b5f60408284031215614859575f5ffd5b61077683836146cb565b5f5f83601f840112614873575f5ffd5b5081356001600160401b03811115614889575f5ffd5b6020830191508360208260051b85010111156118e5575f5ffd5b5f5f5f604084860312156148b5575f5ffd5b83356148c081614638565b925060208401356001600160401b038111156148da575f5ffd5b6148e686828701614863565b9497909650939450505050565b5f6001600160401b0382111561490b5761490b61464c565b5060051b60200190565b5f82601f830112614924575f5ffd5b8135614937614932826148f3565b614688565b8082825260208201915060208360051b860101925085831115614958575f5ffd5b602085015b8381101561497e57803561497081614638565b83526020928301920161495d565b5095945050505050565b5f5f5f5f60a0858703121561499b575f5ffd5b6149a586866146cb565b935060408501356001600160401b038111156149bf575f5ffd5b6149cb87828801614915565b93505060608501356001600160401b038111156149e6575f5ffd5b6149f287828801614915565b925050614a01608086016146b8565b905092959194509250565b5f8151808452602084019350602083015f5b8281101561482d578151865260209586019590910190600101614a1e565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b82811015614a9357603f19878603018452614a7e858351614a0c565b94506020938401939190910190600101614a62565b50929695505050505050565b5f5f60408385031215614ab0575f5ffd5b8235614abb81614638565b915060208301356001600160401b03811115614ad5575f5ffd5b830160a08186031215614ae6575f5ffd5b809150509250929050565b5f5f60408385031215614b02575f5ffd5b8235614b0d81614638565b91506020830135614ae681614638565b5f8151808452602084019350602083015f5b8281101561482d57614b6886835180516001600160401b03168252602080820151600f0b9083015260409081015163ffffffff16910152565b6060959095019460209190910190600101614b2f565b604081525f614b9060408301856147f2565b8281036020840152611fa18185614b1d565b5f8151808452602084019350602083015f5b8281101561482d5781516001600160a01b0316865260209586019590910190600101614bb4565b602081525f6107766020830184614ba2565b5f5f60408385031215614bfe575f5ffd5b82356001600160401b03811115614c13575f5ffd5b614c1f85828601614915565b9250506020830135614ae681614638565b602080825282518282018190525f918401906040840190835b81811015614c705783516001600160401b0316835260209384019390920191600101614c49565b509095945050505050565b5f5f5f5f5f60608688031215614c8f575f5ffd5b8535614c9a81614638565b945060208601356001600160401b03811115614cb4575f5ffd5b614cc088828901614863565b90955093505060408601356001600160401b03811115614cde575f5ffd5b614cea88828901614863565b969995985093965092949392505050565b5f5f5f5f60608587031215614d0e575f5ffd5b8435614d1981614638565b9350614d27602086016146b8565b925060408501356001600160401b03811115614d41575f5ffd5b614d4d87828801614863565b95989497509550505050565b5f5f60408385031215614d6a575f5ffd5b8235614d7581614638565b915060208301356001600160401b03811115614d8f575f5ffd5b614d9b85828601614915565b9150509250929050565b5f5f60408385031215614db6575f5ffd5b8235614dc181614638565b9150614dcf602084016146b8565b90509250929050565b5f60208284031215614de8575f5ffd5b813560ff81168114610776575f5ffd5b5f5f60608385031215614e09575f5ffd5b8235614e1481614638565b9150614dcf84602085016146cb565b5f60608284031215614e33575f5ffd5b50919050565b5f60208284031215614e49575f5ffd5b81356001600160401b03811115614e5e575f5ffd5b61313b84828501614e23565b5f5f5f60808486031215614e7c575f5ffd5b83356001600160401b03811115614e91575f5ffd5b614e9d86828701614915565b93505061475485602086016146cb565b602081525f6107766020830184614b1d565b5f5f5f60608486031215614ed1575f5ffd5b8335614edc81614638565b925060208401356001600160401b03811115614ef6575f5ffd5b614f0286828701614915565b925050614f11604085016146b8565b90509250925092565b5f5f60408385031215614f2b575f5ffd5b8235614f3681614638565b915060208301356001600160401b03811115614f50575f5ffd5b8301601f81018513614f60575f5ffd5b8035614f6e614932826148f3565b8082825260208201915060208360051b850101925087831115614f8f575f5ffd5b602084015b838110156150b45780356001600160401b03811115614fb1575f5ffd5b85016080818b03601f19011215614fc6575f5ffd5b614fce614660565b614fdb8b602084016146cb565b815260608201356001600160401b03811115614ff5575f5ffd5b6150048c602083860101614915565b60208301525060808201356001600160401b03811115615022575f5ffd5b6020818401019250508a601f830112615039575f5ffd5b8135615047614932826148f3565b8082825260208201915060208360051b86010192508d831115615068575f5ffd5b6020850194505b8285101561509e5784356001600160401b038116811461508d575f5ffd5b82526020948501949091019061506f565b6040840152505084525060209283019201614f94565b50809450505050509250929050565b5f5f5f604084860312156150d5575f5ffd5b83356150e081614638565b925060208401356001600160401b038111156150fa575f5ffd5b8401601f8101861361510a575f5ffd5b80356001600160401b0381111561511f575f5ffd5b866020828401011115615130575f5ffd5b939660209190910195509293505050565b5f5f60408385031215615152575f5ffd5b823561515d81614638565b915060208301356001600160401b03811115615177575f5ffd5b614d9b85828601614e23565b5f5f60408385031215615194575f5ffd5b823561519f81614638565b946020939093013593505050565b5f602082840312156151bd575f5ffd5b81518015158114610776575f5ffd5b634e487b7160e01b5f52603260045260245ffd5b5f8235603e198336030181126151f4575f5ffd5b9190910192915050565b5f6020828403121561520e575f5ffd5b610776826146b8565b6040810161099a82846147d4565b5f5f8335601e1984360301811261523a575f5ffd5b8301803591506001600160401b03821115615253575f5ffd5b6020019150600581901b36038213156118e5575f5ffd5b6060810161527882856147d4565b6001600160a01b039290921660409190910152919050565b604081525f6152a26040830185614ba2565b8281036020840152611fa18185614ba2565b5f602082840312156152c4575f5ffd5b81516001600160401b038111156152d9575f5ffd5b8201601f810184136152e9575f5ffd5b80516152f7614932826148f3565b8082825260208201915060208360051b850101925086831115615318575f5ffd5b602084015b838110156153b55780516001600160401b0381111561533a575f5ffd5b8501603f8101891361534a575f5ffd5b602081015161535b614932826148f3565b808282526020820191506020808460051b8601010192508b83111561537e575f5ffd5b6040840193505b828410156153a0578351825260209384019390910190615385565b8652505060209384019391909101905061531d565b509695505050505050565b634e487b7160e01b5f52601160045260245ffd5b8181038181111561099a5761099a6153c0565b6001600160401b03828116828216039081111561099a5761099a6153c0565b5f81600f0b60016001607f1b03198103615422576154226153c0565b5f0392915050565b600f81810b9083900b0160016001607f1b03811360016001607f1b03198212171561099a5761099a6153c0565b6001600160a01b038616815260c0810161547460208301876147d4565b6001600160a01b039490941660608201526001600160401b0392909216608083015263ffffffff1660a09091015292915050565b5f5f8335601e198436030181126154bd575f5ffd5b8301803591506001600160401b038211156154d6575f5ffd5b6020019150368190038213156118e5575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b6001600160a01b03881681525f60c08201615530602084018a6147d4565b60c060608401528690528660e083015f5b8881101561557157823561555481614638565b6001600160a01b0316825260209283019290910190600101615541565b5083810360808501526155848188614a0c565b91505082810360a084015261559a8185876154ea565b9a9950505050505050505050565b5f602082840312156155b8575f5ffd5b813561ffff81168114610776575f5ffd5b63ffffffff818116838216019081111561099a5761099a6153c0565b8183526020830192505f815f5b8481101561482d5763ffffffff615608836146b8565b16865260209586019591909101906001016155f2565b6001600160a01b03841681526040602082018190525f90611fa190830184866155e5565b602081525f61313b6020830184866154ea565b6001600160a01b03861681526060602082018190525f9061567990830186886155e5565b828103604084015261568c8185876154ea565b98975050505050505050565b5f602082840312156156a8575f5ffd5b815161077681614638565b80516020808301519190811015614e33575f1960209190910360031b1b16919050565b5f600182016156e7576156e76153c0565b5060010190565b5f816156fc576156fc6153c0565b505f190190565b8082018281125f831280158216821582161715615722576157226153c0565b505092915050565b600f82810b9082900b0360016001607f1b0319811260016001607f1b038213171561099a5761099a6153c0565b634e487b7160e01b5f52601260045260245ffd5b634e487b7160e01b5f52602160045260245ffd5b8082018082111561099a5761099a6153c0565b634e487b7160e01b5f52603160045260245ffd5b5f826157c057634e487b7160e01b5f52601260045260245ffd5b50049056fea2646970667358221220049ddc48108b86f3be70590462d517d57c47bba9450f474f41b568a93ac1337d64736f6c634300081b0033",
}

// ContractAllocationManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractAllocationManagerMetaData.ABI instead.
var ContractAllocationManagerABI = ContractAllocationManagerMetaData.ABI

// ContractAllocationManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractAllocationManagerMetaData.Bin instead.
var ContractAllocationManagerBin = ContractAllocationManagerMetaData.Bin

// DeployContractAllocationManager deploys a new Ethereum contract, binding an instance of ContractAllocationManager to it.
func DeployContractAllocationManager(auth *bind.TransactOpts, backend bind.ContractBackend, _delegation common.Address, _pauserRegistry common.Address, _permissionController common.Address, _DEALLOCATION_DELAY uint32, _ALLOCATION_CONFIGURATION_DELAY uint32) (common.Address, *types.Transaction, *ContractAllocationManager, error) {
	parsed, err := ContractAllocationManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractAllocationManagerBin), backend, _delegation, _pauserRegistry, _permissionController, _DEALLOCATION_DELAY, _ALLOCATION_CONFIGURATION_DELAY)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractAllocationManager{ContractAllocationManagerCaller: ContractAllocationManagerCaller{contract: contract}, ContractAllocationManagerTransactor: ContractAllocationManagerTransactor{contract: contract}, ContractAllocationManagerFilterer: ContractAllocationManagerFilterer{contract: contract}}, nil
}

// ContractAllocationManagerMethods is an auto generated interface around an Ethereum contract.
type ContractAllocationManagerMethods interface {
	ContractAllocationManagerCalls
	ContractAllocationManagerTransacts
	ContractAllocationManagerFilters
}

// ContractAllocationManagerCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractAllocationManagerCalls interface {
	ALLOCATIONCONFIGURATIONDELAY(opts *bind.CallOpts) (uint32, error)

	DEALLOCATIONDELAY(opts *bind.CallOpts) (uint32, error)

	Delegation(opts *bind.CallOpts) (common.Address, error)

	EncumberedMagnitude(opts *bind.CallOpts, operator common.Address, strategy common.Address) (uint64, error)

	GetAVSRegistrar(opts *bind.CallOpts, avs common.Address) (common.Address, error)

	GetAllocatableMagnitude(opts *bind.CallOpts, operator common.Address, strategy common.Address) (uint64, error)

	GetAllocatedSets(opts *bind.CallOpts, operator common.Address) ([]OperatorSet, error)

	GetAllocatedStrategies(opts *bind.CallOpts, operator common.Address, operatorSet OperatorSet) ([]common.Address, error)

	GetAllocation(opts *bind.CallOpts, operator common.Address, operatorSet OperatorSet, strategy common.Address) (IAllocationManagerTypesAllocation, error)

	GetAllocationDelay(opts *bind.CallOpts, operator common.Address) (bool, uint32, error)

	GetAllocations(opts *bind.CallOpts, operators []common.Address, operatorSet OperatorSet, strategy common.Address) ([]IAllocationManagerTypesAllocation, error)

	GetMaxMagnitude(opts *bind.CallOpts, operator common.Address, strategy common.Address) (uint64, error)

	GetMaxMagnitudes(opts *bind.CallOpts, operators []common.Address, strategy common.Address) ([]uint64, error)

	GetMaxMagnitudes0(opts *bind.CallOpts, operator common.Address, strategies []common.Address) ([]uint64, error)

	GetMaxMagnitudesAtBlock(opts *bind.CallOpts, operator common.Address, strategies []common.Address, blockNumber uint32) ([]uint64, error)

	GetMemberCount(opts *bind.CallOpts, operatorSet OperatorSet) (*big.Int, error)

	GetMembers(opts *bind.CallOpts, operatorSet OperatorSet) ([]common.Address, error)

	GetMinimumSlashableStake(opts *bind.CallOpts, operatorSet OperatorSet, operators []common.Address, strategies []common.Address, futureBlock uint32) ([][]*big.Int, error)

	GetOperatorSetCount(opts *bind.CallOpts, avs common.Address) (*big.Int, error)

	GetRegisteredSets(opts *bind.CallOpts, operator common.Address) ([]OperatorSet, error)

	GetStrategiesInOperatorSet(opts *bind.CallOpts, operatorSet OperatorSet) ([]common.Address, error)

	GetStrategyAllocations(opts *bind.CallOpts, operator common.Address, strategy common.Address) ([]OperatorSet, []IAllocationManagerTypesAllocation, error)

	IsMemberOfOperatorSet(opts *bind.CallOpts, operator common.Address, operatorSet OperatorSet) (bool, error)

	IsOperatorSet(opts *bind.CallOpts, operatorSet OperatorSet) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts, index uint8) (bool, error)

	Paused0(opts *bind.CallOpts) (*big.Int, error)

	PauserRegistry(opts *bind.CallOpts) (common.Address, error)

	PermissionController(opts *bind.CallOpts) (common.Address, error)
}

// ContractAllocationManagerTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractAllocationManagerTransacts interface {
	AddStrategiesToOperatorSet(opts *bind.TransactOpts, avs common.Address, operatorSetId uint32, strategies []common.Address) (*types.Transaction, error)

	ClearDeallocationQueue(opts *bind.TransactOpts, operator common.Address, strategies []common.Address, numToClear []uint16) (*types.Transaction, error)

	CreateOperatorSets(opts *bind.TransactOpts, avs common.Address, params []IAllocationManagerTypesCreateSetParams) (*types.Transaction, error)

	DeregisterFromOperatorSets(opts *bind.TransactOpts, params IAllocationManagerTypesDeregisterParams) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts, initialOwner common.Address, initialPausedStatus *big.Int) (*types.Transaction, error)

	ModifyAllocations(opts *bind.TransactOpts, operator common.Address, params []IAllocationManagerTypesAllocateParams) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error)

	PauseAll(opts *bind.TransactOpts) (*types.Transaction, error)

	RegisterForOperatorSets(opts *bind.TransactOpts, operator common.Address, params IAllocationManagerTypesRegisterParams) (*types.Transaction, error)

	RemoveStrategiesFromOperatorSet(opts *bind.TransactOpts, avs common.Address, operatorSetId uint32, strategies []common.Address) (*types.Transaction, error)

	RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetAVSRegistrar(opts *bind.TransactOpts, avs common.Address, registrar common.Address) (*types.Transaction, error)

	SetAllocationDelay(opts *bind.TransactOpts, operator common.Address, delay uint32) (*types.Transaction, error)

	SlashOperator(opts *bind.TransactOpts, avs common.Address, params IAllocationManagerTypesSlashingParams) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error)

	UpdateAVSMetadataURI(opts *bind.TransactOpts, avs common.Address, metadataURI string) (*types.Transaction, error)
}

// ContractAllocationManagerFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractAllocationManagerFilters interface {
	FilterAVSMetadataURIUpdated(opts *bind.FilterOpts, avs []common.Address) (*ContractAllocationManagerAVSMetadataURIUpdatedIterator, error)
	WatchAVSMetadataURIUpdated(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerAVSMetadataURIUpdated, avs []common.Address) (event.Subscription, error)
	ParseAVSMetadataURIUpdated(log types.Log) (*ContractAllocationManagerAVSMetadataURIUpdated, error)

	FilterAVSRegistrarSet(opts *bind.FilterOpts) (*ContractAllocationManagerAVSRegistrarSetIterator, error)
	WatchAVSRegistrarSet(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerAVSRegistrarSet) (event.Subscription, error)
	ParseAVSRegistrarSet(log types.Log) (*ContractAllocationManagerAVSRegistrarSet, error)

	FilterAllocationDelaySet(opts *bind.FilterOpts) (*ContractAllocationManagerAllocationDelaySetIterator, error)
	WatchAllocationDelaySet(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerAllocationDelaySet) (event.Subscription, error)
	ParseAllocationDelaySet(log types.Log) (*ContractAllocationManagerAllocationDelaySet, error)

	FilterAllocationUpdated(opts *bind.FilterOpts) (*ContractAllocationManagerAllocationUpdatedIterator, error)
	WatchAllocationUpdated(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerAllocationUpdated) (event.Subscription, error)
	ParseAllocationUpdated(log types.Log) (*ContractAllocationManagerAllocationUpdated, error)

	FilterEncumberedMagnitudeUpdated(opts *bind.FilterOpts) (*ContractAllocationManagerEncumberedMagnitudeUpdatedIterator, error)
	WatchEncumberedMagnitudeUpdated(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerEncumberedMagnitudeUpdated) (event.Subscription, error)
	ParseEncumberedMagnitudeUpdated(log types.Log) (*ContractAllocationManagerEncumberedMagnitudeUpdated, error)

	FilterInitialized(opts *bind.FilterOpts) (*ContractAllocationManagerInitializedIterator, error)
	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerInitialized) (event.Subscription, error)
	ParseInitialized(log types.Log) (*ContractAllocationManagerInitialized, error)

	FilterMaxMagnitudeUpdated(opts *bind.FilterOpts) (*ContractAllocationManagerMaxMagnitudeUpdatedIterator, error)
	WatchMaxMagnitudeUpdated(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerMaxMagnitudeUpdated) (event.Subscription, error)
	ParseMaxMagnitudeUpdated(log types.Log) (*ContractAllocationManagerMaxMagnitudeUpdated, error)

	FilterOperatorAddedToOperatorSet(opts *bind.FilterOpts, operator []common.Address) (*ContractAllocationManagerOperatorAddedToOperatorSetIterator, error)
	WatchOperatorAddedToOperatorSet(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerOperatorAddedToOperatorSet, operator []common.Address) (event.Subscription, error)
	ParseOperatorAddedToOperatorSet(log types.Log) (*ContractAllocationManagerOperatorAddedToOperatorSet, error)

	FilterOperatorRemovedFromOperatorSet(opts *bind.FilterOpts, operator []common.Address) (*ContractAllocationManagerOperatorRemovedFromOperatorSetIterator, error)
	WatchOperatorRemovedFromOperatorSet(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerOperatorRemovedFromOperatorSet, operator []common.Address) (event.Subscription, error)
	ParseOperatorRemovedFromOperatorSet(log types.Log) (*ContractAllocationManagerOperatorRemovedFromOperatorSet, error)

	FilterOperatorSetCreated(opts *bind.FilterOpts) (*ContractAllocationManagerOperatorSetCreatedIterator, error)
	WatchOperatorSetCreated(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerOperatorSetCreated) (event.Subscription, error)
	ParseOperatorSetCreated(log types.Log) (*ContractAllocationManagerOperatorSetCreated, error)

	FilterOperatorSlashed(opts *bind.FilterOpts) (*ContractAllocationManagerOperatorSlashedIterator, error)
	WatchOperatorSlashed(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerOperatorSlashed) (event.Subscription, error)
	ParseOperatorSlashed(log types.Log) (*ContractAllocationManagerOperatorSlashed, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractAllocationManagerOwnershipTransferredIterator, error)
	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error)
	ParseOwnershipTransferred(log types.Log) (*ContractAllocationManagerOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractAllocationManagerPausedIterator, error)
	WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerPaused, account []common.Address) (event.Subscription, error)
	ParsePaused(log types.Log) (*ContractAllocationManagerPaused, error)

	FilterStrategyAddedToOperatorSet(opts *bind.FilterOpts) (*ContractAllocationManagerStrategyAddedToOperatorSetIterator, error)
	WatchStrategyAddedToOperatorSet(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerStrategyAddedToOperatorSet) (event.Subscription, error)
	ParseStrategyAddedToOperatorSet(log types.Log) (*ContractAllocationManagerStrategyAddedToOperatorSet, error)

	FilterStrategyRemovedFromOperatorSet(opts *bind.FilterOpts) (*ContractAllocationManagerStrategyRemovedFromOperatorSetIterator, error)
	WatchStrategyRemovedFromOperatorSet(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerStrategyRemovedFromOperatorSet) (event.Subscription, error)
	ParseStrategyRemovedFromOperatorSet(log types.Log) (*ContractAllocationManagerStrategyRemovedFromOperatorSet, error)

	FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractAllocationManagerUnpausedIterator, error)
	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerUnpaused, account []common.Address) (event.Subscription, error)
	ParseUnpaused(log types.Log) (*ContractAllocationManagerUnpaused, error)
}

// ContractAllocationManager is an auto generated Go binding around an Ethereum contract.
type ContractAllocationManager struct {
	ContractAllocationManagerCaller     // Read-only binding to the contract
	ContractAllocationManagerTransactor // Write-only binding to the contract
	ContractAllocationManagerFilterer   // Log filterer for contract events
}

// ContractAllocationManager implements the ContractAllocationManagerMethods interface.
var _ ContractAllocationManagerMethods = (*ContractAllocationManager)(nil)

// ContractAllocationManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractAllocationManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAllocationManagerCaller implements the ContractAllocationManagerCalls interface.
var _ ContractAllocationManagerCalls = (*ContractAllocationManagerCaller)(nil)

// ContractAllocationManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractAllocationManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAllocationManagerTransactor implements the ContractAllocationManagerTransacts interface.
var _ ContractAllocationManagerTransacts = (*ContractAllocationManagerTransactor)(nil)

// ContractAllocationManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractAllocationManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAllocationManagerFilterer implements the ContractAllocationManagerFilters interface.
var _ ContractAllocationManagerFilters = (*ContractAllocationManagerFilterer)(nil)

// ContractAllocationManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractAllocationManagerSession struct {
	Contract     *ContractAllocationManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ContractAllocationManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractAllocationManagerCallerSession struct {
	Contract *ContractAllocationManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// ContractAllocationManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractAllocationManagerTransactorSession struct {
	Contract     *ContractAllocationManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// ContractAllocationManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractAllocationManagerRaw struct {
	Contract *ContractAllocationManager // Generic contract binding to access the raw methods on
}

// ContractAllocationManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractAllocationManagerCallerRaw struct {
	Contract *ContractAllocationManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractAllocationManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractAllocationManagerTransactorRaw struct {
	Contract *ContractAllocationManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractAllocationManager creates a new instance of ContractAllocationManager, bound to a specific deployed contract.
func NewContractAllocationManager(address common.Address, backend bind.ContractBackend) (*ContractAllocationManager, error) {
	contract, err := bindContractAllocationManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractAllocationManager{ContractAllocationManagerCaller: ContractAllocationManagerCaller{contract: contract}, ContractAllocationManagerTransactor: ContractAllocationManagerTransactor{contract: contract}, ContractAllocationManagerFilterer: ContractAllocationManagerFilterer{contract: contract}}, nil
}

// NewContractAllocationManagerCaller creates a new read-only instance of ContractAllocationManager, bound to a specific deployed contract.
func NewContractAllocationManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractAllocationManagerCaller, error) {
	contract, err := bindContractAllocationManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractAllocationManagerCaller{contract: contract}, nil
}

// NewContractAllocationManagerTransactor creates a new write-only instance of ContractAllocationManager, bound to a specific deployed contract.
func NewContractAllocationManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractAllocationManagerTransactor, error) {
	contract, err := bindContractAllocationManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractAllocationManagerTransactor{contract: contract}, nil
}

// NewContractAllocationManagerFilterer creates a new log filterer instance of ContractAllocationManager, bound to a specific deployed contract.
func NewContractAllocationManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractAllocationManagerFilterer, error) {
	contract, err := bindContractAllocationManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractAllocationManagerFilterer{contract: contract}, nil
}

// bindContractAllocationManager binds a generic wrapper to an already deployed contract.
func bindContractAllocationManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractAllocationManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractAllocationManager *ContractAllocationManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractAllocationManager.Contract.ContractAllocationManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractAllocationManager *ContractAllocationManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.ContractAllocationManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractAllocationManager *ContractAllocationManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.ContractAllocationManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractAllocationManager *ContractAllocationManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractAllocationManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractAllocationManager *ContractAllocationManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractAllocationManager *ContractAllocationManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.contract.Transact(opts, method, params...)
}

// ALLOCATIONCONFIGURATIONDELAY is a free data retrieval call binding the contract method 0x7bc1ef61.
//
// Solidity: function ALLOCATION_CONFIGURATION_DELAY() view returns(uint32)
func (_ContractAllocationManager *ContractAllocationManagerCaller) ALLOCATIONCONFIGURATIONDELAY(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "ALLOCATION_CONFIGURATION_DELAY")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ALLOCATIONCONFIGURATIONDELAY is a free data retrieval call binding the contract method 0x7bc1ef61.
//
// Solidity: function ALLOCATION_CONFIGURATION_DELAY() view returns(uint32)
func (_ContractAllocationManager *ContractAllocationManagerSession) ALLOCATIONCONFIGURATIONDELAY() (uint32, error) {
	return _ContractAllocationManager.Contract.ALLOCATIONCONFIGURATIONDELAY(&_ContractAllocationManager.CallOpts)
}

// ALLOCATIONCONFIGURATIONDELAY is a free data retrieval call binding the contract method 0x7bc1ef61.
//
// Solidity: function ALLOCATION_CONFIGURATION_DELAY() view returns(uint32)
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) ALLOCATIONCONFIGURATIONDELAY() (uint32, error) {
	return _ContractAllocationManager.Contract.ALLOCATIONCONFIGURATIONDELAY(&_ContractAllocationManager.CallOpts)
}

// DEALLOCATIONDELAY is a free data retrieval call binding the contract method 0x2981eb77.
//
// Solidity: function DEALLOCATION_DELAY() view returns(uint32)
func (_ContractAllocationManager *ContractAllocationManagerCaller) DEALLOCATIONDELAY(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "DEALLOCATION_DELAY")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// DEALLOCATIONDELAY is a free data retrieval call binding the contract method 0x2981eb77.
//
// Solidity: function DEALLOCATION_DELAY() view returns(uint32)
func (_ContractAllocationManager *ContractAllocationManagerSession) DEALLOCATIONDELAY() (uint32, error) {
	return _ContractAllocationManager.Contract.DEALLOCATIONDELAY(&_ContractAllocationManager.CallOpts)
}

// DEALLOCATIONDELAY is a free data retrieval call binding the contract method 0x2981eb77.
//
// Solidity: function DEALLOCATION_DELAY() view returns(uint32)
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) DEALLOCATIONDELAY() (uint32, error) {
	return _ContractAllocationManager.Contract.DEALLOCATIONDELAY(&_ContractAllocationManager.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractAllocationManager *ContractAllocationManagerCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractAllocationManager *ContractAllocationManagerSession) Delegation() (common.Address, error) {
	return _ContractAllocationManager.Contract.Delegation(&_ContractAllocationManager.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) Delegation() (common.Address, error) {
	return _ContractAllocationManager.Contract.Delegation(&_ContractAllocationManager.CallOpts)
}

// EncumberedMagnitude is a free data retrieval call binding the contract method 0xa984eb3a.
//
// Solidity: function encumberedMagnitude(address operator, address strategy) view returns(uint64)
func (_ContractAllocationManager *ContractAllocationManagerCaller) EncumberedMagnitude(opts *bind.CallOpts, operator common.Address, strategy common.Address) (uint64, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "encumberedMagnitude", operator, strategy)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// EncumberedMagnitude is a free data retrieval call binding the contract method 0xa984eb3a.
//
// Solidity: function encumberedMagnitude(address operator, address strategy) view returns(uint64)
func (_ContractAllocationManager *ContractAllocationManagerSession) EncumberedMagnitude(operator common.Address, strategy common.Address) (uint64, error) {
	return _ContractAllocationManager.Contract.EncumberedMagnitude(&_ContractAllocationManager.CallOpts, operator, strategy)
}

// EncumberedMagnitude is a free data retrieval call binding the contract method 0xa984eb3a.
//
// Solidity: function encumberedMagnitude(address operator, address strategy) view returns(uint64)
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) EncumberedMagnitude(operator common.Address, strategy common.Address) (uint64, error) {
	return _ContractAllocationManager.Contract.EncumberedMagnitude(&_ContractAllocationManager.CallOpts, operator, strategy)
}

// GetAVSRegistrar is a free data retrieval call binding the contract method 0x304c10cd.
//
// Solidity: function getAVSRegistrar(address avs) view returns(address)
func (_ContractAllocationManager *ContractAllocationManagerCaller) GetAVSRegistrar(opts *bind.CallOpts, avs common.Address) (common.Address, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "getAVSRegistrar", avs)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAVSRegistrar is a free data retrieval call binding the contract method 0x304c10cd.
//
// Solidity: function getAVSRegistrar(address avs) view returns(address)
func (_ContractAllocationManager *ContractAllocationManagerSession) GetAVSRegistrar(avs common.Address) (common.Address, error) {
	return _ContractAllocationManager.Contract.GetAVSRegistrar(&_ContractAllocationManager.CallOpts, avs)
}

// GetAVSRegistrar is a free data retrieval call binding the contract method 0x304c10cd.
//
// Solidity: function getAVSRegistrar(address avs) view returns(address)
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) GetAVSRegistrar(avs common.Address) (common.Address, error) {
	return _ContractAllocationManager.Contract.GetAVSRegistrar(&_ContractAllocationManager.CallOpts, avs)
}

// GetAllocatableMagnitude is a free data retrieval call binding the contract method 0x6cfb4481.
//
// Solidity: function getAllocatableMagnitude(address operator, address strategy) view returns(uint64)
func (_ContractAllocationManager *ContractAllocationManagerCaller) GetAllocatableMagnitude(opts *bind.CallOpts, operator common.Address, strategy common.Address) (uint64, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "getAllocatableMagnitude", operator, strategy)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetAllocatableMagnitude is a free data retrieval call binding the contract method 0x6cfb4481.
//
// Solidity: function getAllocatableMagnitude(address operator, address strategy) view returns(uint64)
func (_ContractAllocationManager *ContractAllocationManagerSession) GetAllocatableMagnitude(operator common.Address, strategy common.Address) (uint64, error) {
	return _ContractAllocationManager.Contract.GetAllocatableMagnitude(&_ContractAllocationManager.CallOpts, operator, strategy)
}

// GetAllocatableMagnitude is a free data retrieval call binding the contract method 0x6cfb4481.
//
// Solidity: function getAllocatableMagnitude(address operator, address strategy) view returns(uint64)
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) GetAllocatableMagnitude(operator common.Address, strategy common.Address) (uint64, error) {
	return _ContractAllocationManager.Contract.GetAllocatableMagnitude(&_ContractAllocationManager.CallOpts, operator, strategy)
}

// GetAllocatedSets is a free data retrieval call binding the contract method 0x15fe5028.
//
// Solidity: function getAllocatedSets(address operator) view returns((address,uint32)[])
func (_ContractAllocationManager *ContractAllocationManagerCaller) GetAllocatedSets(opts *bind.CallOpts, operator common.Address) ([]OperatorSet, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "getAllocatedSets", operator)

	if err != nil {
		return *new([]OperatorSet), err
	}

	out0 := *abi.ConvertType(out[0], new([]OperatorSet)).(*[]OperatorSet)

	return out0, err

}

// GetAllocatedSets is a free data retrieval call binding the contract method 0x15fe5028.
//
// Solidity: function getAllocatedSets(address operator) view returns((address,uint32)[])
func (_ContractAllocationManager *ContractAllocationManagerSession) GetAllocatedSets(operator common.Address) ([]OperatorSet, error) {
	return _ContractAllocationManager.Contract.GetAllocatedSets(&_ContractAllocationManager.CallOpts, operator)
}

// GetAllocatedSets is a free data retrieval call binding the contract method 0x15fe5028.
//
// Solidity: function getAllocatedSets(address operator) view returns((address,uint32)[])
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) GetAllocatedSets(operator common.Address) ([]OperatorSet, error) {
	return _ContractAllocationManager.Contract.GetAllocatedSets(&_ContractAllocationManager.CallOpts, operator)
}

// GetAllocatedStrategies is a free data retrieval call binding the contract method 0xc221d8ae.
//
// Solidity: function getAllocatedStrategies(address operator, (address,uint32) operatorSet) view returns(address[])
func (_ContractAllocationManager *ContractAllocationManagerCaller) GetAllocatedStrategies(opts *bind.CallOpts, operator common.Address, operatorSet OperatorSet) ([]common.Address, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "getAllocatedStrategies", operator, operatorSet)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllocatedStrategies is a free data retrieval call binding the contract method 0xc221d8ae.
//
// Solidity: function getAllocatedStrategies(address operator, (address,uint32) operatorSet) view returns(address[])
func (_ContractAllocationManager *ContractAllocationManagerSession) GetAllocatedStrategies(operator common.Address, operatorSet OperatorSet) ([]common.Address, error) {
	return _ContractAllocationManager.Contract.GetAllocatedStrategies(&_ContractAllocationManager.CallOpts, operator, operatorSet)
}

// GetAllocatedStrategies is a free data retrieval call binding the contract method 0xc221d8ae.
//
// Solidity: function getAllocatedStrategies(address operator, (address,uint32) operatorSet) view returns(address[])
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) GetAllocatedStrategies(operator common.Address, operatorSet OperatorSet) ([]common.Address, error) {
	return _ContractAllocationManager.Contract.GetAllocatedStrategies(&_ContractAllocationManager.CallOpts, operator, operatorSet)
}

// GetAllocation is a free data retrieval call binding the contract method 0x10e1b9b8.
//
// Solidity: function getAllocation(address operator, (address,uint32) operatorSet, address strategy) view returns((uint64,int128,uint32))
func (_ContractAllocationManager *ContractAllocationManagerCaller) GetAllocation(opts *bind.CallOpts, operator common.Address, operatorSet OperatorSet, strategy common.Address) (IAllocationManagerTypesAllocation, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "getAllocation", operator, operatorSet, strategy)

	if err != nil {
		return *new(IAllocationManagerTypesAllocation), err
	}

	out0 := *abi.ConvertType(out[0], new(IAllocationManagerTypesAllocation)).(*IAllocationManagerTypesAllocation)

	return out0, err

}

// GetAllocation is a free data retrieval call binding the contract method 0x10e1b9b8.
//
// Solidity: function getAllocation(address operator, (address,uint32) operatorSet, address strategy) view returns((uint64,int128,uint32))
func (_ContractAllocationManager *ContractAllocationManagerSession) GetAllocation(operator common.Address, operatorSet OperatorSet, strategy common.Address) (IAllocationManagerTypesAllocation, error) {
	return _ContractAllocationManager.Contract.GetAllocation(&_ContractAllocationManager.CallOpts, operator, operatorSet, strategy)
}

// GetAllocation is a free data retrieval call binding the contract method 0x10e1b9b8.
//
// Solidity: function getAllocation(address operator, (address,uint32) operatorSet, address strategy) view returns((uint64,int128,uint32))
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) GetAllocation(operator common.Address, operatorSet OperatorSet, strategy common.Address) (IAllocationManagerTypesAllocation, error) {
	return _ContractAllocationManager.Contract.GetAllocation(&_ContractAllocationManager.CallOpts, operator, operatorSet, strategy)
}

// GetAllocationDelay is a free data retrieval call binding the contract method 0xb9fbaed1.
//
// Solidity: function getAllocationDelay(address operator) view returns(bool, uint32)
func (_ContractAllocationManager *ContractAllocationManagerCaller) GetAllocationDelay(opts *bind.CallOpts, operator common.Address) (bool, uint32, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "getAllocationDelay", operator)

	if err != nil {
		return *new(bool), *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return out0, out1, err

}

// GetAllocationDelay is a free data retrieval call binding the contract method 0xb9fbaed1.
//
// Solidity: function getAllocationDelay(address operator) view returns(bool, uint32)
func (_ContractAllocationManager *ContractAllocationManagerSession) GetAllocationDelay(operator common.Address) (bool, uint32, error) {
	return _ContractAllocationManager.Contract.GetAllocationDelay(&_ContractAllocationManager.CallOpts, operator)
}

// GetAllocationDelay is a free data retrieval call binding the contract method 0xb9fbaed1.
//
// Solidity: function getAllocationDelay(address operator) view returns(bool, uint32)
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) GetAllocationDelay(operator common.Address) (bool, uint32, error) {
	return _ContractAllocationManager.Contract.GetAllocationDelay(&_ContractAllocationManager.CallOpts, operator)
}

// GetAllocations is a free data retrieval call binding the contract method 0x8ce64854.
//
// Solidity: function getAllocations(address[] operators, (address,uint32) operatorSet, address strategy) view returns((uint64,int128,uint32)[])
func (_ContractAllocationManager *ContractAllocationManagerCaller) GetAllocations(opts *bind.CallOpts, operators []common.Address, operatorSet OperatorSet, strategy common.Address) ([]IAllocationManagerTypesAllocation, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "getAllocations", operators, operatorSet, strategy)

	if err != nil {
		return *new([]IAllocationManagerTypesAllocation), err
	}

	out0 := *abi.ConvertType(out[0], new([]IAllocationManagerTypesAllocation)).(*[]IAllocationManagerTypesAllocation)

	return out0, err

}

// GetAllocations is a free data retrieval call binding the contract method 0x8ce64854.
//
// Solidity: function getAllocations(address[] operators, (address,uint32) operatorSet, address strategy) view returns((uint64,int128,uint32)[])
func (_ContractAllocationManager *ContractAllocationManagerSession) GetAllocations(operators []common.Address, operatorSet OperatorSet, strategy common.Address) ([]IAllocationManagerTypesAllocation, error) {
	return _ContractAllocationManager.Contract.GetAllocations(&_ContractAllocationManager.CallOpts, operators, operatorSet, strategy)
}

// GetAllocations is a free data retrieval call binding the contract method 0x8ce64854.
//
// Solidity: function getAllocations(address[] operators, (address,uint32) operatorSet, address strategy) view returns((uint64,int128,uint32)[])
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) GetAllocations(operators []common.Address, operatorSet OperatorSet, strategy common.Address) ([]IAllocationManagerTypesAllocation, error) {
	return _ContractAllocationManager.Contract.GetAllocations(&_ContractAllocationManager.CallOpts, operators, operatorSet, strategy)
}

// GetMaxMagnitude is a free data retrieval call binding the contract method 0xa9333ec8.
//
// Solidity: function getMaxMagnitude(address operator, address strategy) view returns(uint64)
func (_ContractAllocationManager *ContractAllocationManagerCaller) GetMaxMagnitude(opts *bind.CallOpts, operator common.Address, strategy common.Address) (uint64, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "getMaxMagnitude", operator, strategy)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetMaxMagnitude is a free data retrieval call binding the contract method 0xa9333ec8.
//
// Solidity: function getMaxMagnitude(address operator, address strategy) view returns(uint64)
func (_ContractAllocationManager *ContractAllocationManagerSession) GetMaxMagnitude(operator common.Address, strategy common.Address) (uint64, error) {
	return _ContractAllocationManager.Contract.GetMaxMagnitude(&_ContractAllocationManager.CallOpts, operator, strategy)
}

// GetMaxMagnitude is a free data retrieval call binding the contract method 0xa9333ec8.
//
// Solidity: function getMaxMagnitude(address operator, address strategy) view returns(uint64)
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) GetMaxMagnitude(operator common.Address, strategy common.Address) (uint64, error) {
	return _ContractAllocationManager.Contract.GetMaxMagnitude(&_ContractAllocationManager.CallOpts, operator, strategy)
}

// GetMaxMagnitudes is a free data retrieval call binding the contract method 0x4a10ffe5.
//
// Solidity: function getMaxMagnitudes(address[] operators, address strategy) view returns(uint64[])
func (_ContractAllocationManager *ContractAllocationManagerCaller) GetMaxMagnitudes(opts *bind.CallOpts, operators []common.Address, strategy common.Address) ([]uint64, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "getMaxMagnitudes", operators, strategy)

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

// GetMaxMagnitudes is a free data retrieval call binding the contract method 0x4a10ffe5.
//
// Solidity: function getMaxMagnitudes(address[] operators, address strategy) view returns(uint64[])
func (_ContractAllocationManager *ContractAllocationManagerSession) GetMaxMagnitudes(operators []common.Address, strategy common.Address) ([]uint64, error) {
	return _ContractAllocationManager.Contract.GetMaxMagnitudes(&_ContractAllocationManager.CallOpts, operators, strategy)
}

// GetMaxMagnitudes is a free data retrieval call binding the contract method 0x4a10ffe5.
//
// Solidity: function getMaxMagnitudes(address[] operators, address strategy) view returns(uint64[])
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) GetMaxMagnitudes(operators []common.Address, strategy common.Address) ([]uint64, error) {
	return _ContractAllocationManager.Contract.GetMaxMagnitudes(&_ContractAllocationManager.CallOpts, operators, strategy)
}

// GetMaxMagnitudes0 is a free data retrieval call binding the contract method 0x547afb87.
//
// Solidity: function getMaxMagnitudes(address operator, address[] strategies) view returns(uint64[])
func (_ContractAllocationManager *ContractAllocationManagerCaller) GetMaxMagnitudes0(opts *bind.CallOpts, operator common.Address, strategies []common.Address) ([]uint64, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "getMaxMagnitudes0", operator, strategies)

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

// GetMaxMagnitudes0 is a free data retrieval call binding the contract method 0x547afb87.
//
// Solidity: function getMaxMagnitudes(address operator, address[] strategies) view returns(uint64[])
func (_ContractAllocationManager *ContractAllocationManagerSession) GetMaxMagnitudes0(operator common.Address, strategies []common.Address) ([]uint64, error) {
	return _ContractAllocationManager.Contract.GetMaxMagnitudes0(&_ContractAllocationManager.CallOpts, operator, strategies)
}

// GetMaxMagnitudes0 is a free data retrieval call binding the contract method 0x547afb87.
//
// Solidity: function getMaxMagnitudes(address operator, address[] strategies) view returns(uint64[])
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) GetMaxMagnitudes0(operator common.Address, strategies []common.Address) ([]uint64, error) {
	return _ContractAllocationManager.Contract.GetMaxMagnitudes0(&_ContractAllocationManager.CallOpts, operator, strategies)
}

// GetMaxMagnitudesAtBlock is a free data retrieval call binding the contract method 0x94d7d00c.
//
// Solidity: function getMaxMagnitudesAtBlock(address operator, address[] strategies, uint32 blockNumber) view returns(uint64[])
func (_ContractAllocationManager *ContractAllocationManagerCaller) GetMaxMagnitudesAtBlock(opts *bind.CallOpts, operator common.Address, strategies []common.Address, blockNumber uint32) ([]uint64, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "getMaxMagnitudesAtBlock", operator, strategies, blockNumber)

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

// GetMaxMagnitudesAtBlock is a free data retrieval call binding the contract method 0x94d7d00c.
//
// Solidity: function getMaxMagnitudesAtBlock(address operator, address[] strategies, uint32 blockNumber) view returns(uint64[])
func (_ContractAllocationManager *ContractAllocationManagerSession) GetMaxMagnitudesAtBlock(operator common.Address, strategies []common.Address, blockNumber uint32) ([]uint64, error) {
	return _ContractAllocationManager.Contract.GetMaxMagnitudesAtBlock(&_ContractAllocationManager.CallOpts, operator, strategies, blockNumber)
}

// GetMaxMagnitudesAtBlock is a free data retrieval call binding the contract method 0x94d7d00c.
//
// Solidity: function getMaxMagnitudesAtBlock(address operator, address[] strategies, uint32 blockNumber) view returns(uint64[])
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) GetMaxMagnitudesAtBlock(operator common.Address, strategies []common.Address, blockNumber uint32) ([]uint64, error) {
	return _ContractAllocationManager.Contract.GetMaxMagnitudesAtBlock(&_ContractAllocationManager.CallOpts, operator, strategies, blockNumber)
}

// GetMemberCount is a free data retrieval call binding the contract method 0xb2447af7.
//
// Solidity: function getMemberCount((address,uint32) operatorSet) view returns(uint256)
func (_ContractAllocationManager *ContractAllocationManagerCaller) GetMemberCount(opts *bind.CallOpts, operatorSet OperatorSet) (*big.Int, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "getMemberCount", operatorSet)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMemberCount is a free data retrieval call binding the contract method 0xb2447af7.
//
// Solidity: function getMemberCount((address,uint32) operatorSet) view returns(uint256)
func (_ContractAllocationManager *ContractAllocationManagerSession) GetMemberCount(operatorSet OperatorSet) (*big.Int, error) {
	return _ContractAllocationManager.Contract.GetMemberCount(&_ContractAllocationManager.CallOpts, operatorSet)
}

// GetMemberCount is a free data retrieval call binding the contract method 0xb2447af7.
//
// Solidity: function getMemberCount((address,uint32) operatorSet) view returns(uint256)
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) GetMemberCount(operatorSet OperatorSet) (*big.Int, error) {
	return _ContractAllocationManager.Contract.GetMemberCount(&_ContractAllocationManager.CallOpts, operatorSet)
}

// GetMembers is a free data retrieval call binding the contract method 0x6e875dba.
//
// Solidity: function getMembers((address,uint32) operatorSet) view returns(address[])
func (_ContractAllocationManager *ContractAllocationManagerCaller) GetMembers(opts *bind.CallOpts, operatorSet OperatorSet) ([]common.Address, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "getMembers", operatorSet)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetMembers is a free data retrieval call binding the contract method 0x6e875dba.
//
// Solidity: function getMembers((address,uint32) operatorSet) view returns(address[])
func (_ContractAllocationManager *ContractAllocationManagerSession) GetMembers(operatorSet OperatorSet) ([]common.Address, error) {
	return _ContractAllocationManager.Contract.GetMembers(&_ContractAllocationManager.CallOpts, operatorSet)
}

// GetMembers is a free data retrieval call binding the contract method 0x6e875dba.
//
// Solidity: function getMembers((address,uint32) operatorSet) view returns(address[])
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) GetMembers(operatorSet OperatorSet) ([]common.Address, error) {
	return _ContractAllocationManager.Contract.GetMembers(&_ContractAllocationManager.CallOpts, operatorSet)
}

// GetMinimumSlashableStake is a free data retrieval call binding the contract method 0x2bab2c4a.
//
// Solidity: function getMinimumSlashableStake((address,uint32) operatorSet, address[] operators, address[] strategies, uint32 futureBlock) view returns(uint256[][] slashableStake)
func (_ContractAllocationManager *ContractAllocationManagerCaller) GetMinimumSlashableStake(opts *bind.CallOpts, operatorSet OperatorSet, operators []common.Address, strategies []common.Address, futureBlock uint32) ([][]*big.Int, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "getMinimumSlashableStake", operatorSet, operators, strategies, futureBlock)

	if err != nil {
		return *new([][]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([][]*big.Int)).(*[][]*big.Int)

	return out0, err

}

// GetMinimumSlashableStake is a free data retrieval call binding the contract method 0x2bab2c4a.
//
// Solidity: function getMinimumSlashableStake((address,uint32) operatorSet, address[] operators, address[] strategies, uint32 futureBlock) view returns(uint256[][] slashableStake)
func (_ContractAllocationManager *ContractAllocationManagerSession) GetMinimumSlashableStake(operatorSet OperatorSet, operators []common.Address, strategies []common.Address, futureBlock uint32) ([][]*big.Int, error) {
	return _ContractAllocationManager.Contract.GetMinimumSlashableStake(&_ContractAllocationManager.CallOpts, operatorSet, operators, strategies, futureBlock)
}

// GetMinimumSlashableStake is a free data retrieval call binding the contract method 0x2bab2c4a.
//
// Solidity: function getMinimumSlashableStake((address,uint32) operatorSet, address[] operators, address[] strategies, uint32 futureBlock) view returns(uint256[][] slashableStake)
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) GetMinimumSlashableStake(operatorSet OperatorSet, operators []common.Address, strategies []common.Address, futureBlock uint32) ([][]*big.Int, error) {
	return _ContractAllocationManager.Contract.GetMinimumSlashableStake(&_ContractAllocationManager.CallOpts, operatorSet, operators, strategies, futureBlock)
}

// GetOperatorSetCount is a free data retrieval call binding the contract method 0xba1a84e5.
//
// Solidity: function getOperatorSetCount(address avs) view returns(uint256)
func (_ContractAllocationManager *ContractAllocationManagerCaller) GetOperatorSetCount(opts *bind.CallOpts, avs common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "getOperatorSetCount", avs)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorSetCount is a free data retrieval call binding the contract method 0xba1a84e5.
//
// Solidity: function getOperatorSetCount(address avs) view returns(uint256)
func (_ContractAllocationManager *ContractAllocationManagerSession) GetOperatorSetCount(avs common.Address) (*big.Int, error) {
	return _ContractAllocationManager.Contract.GetOperatorSetCount(&_ContractAllocationManager.CallOpts, avs)
}

// GetOperatorSetCount is a free data retrieval call binding the contract method 0xba1a84e5.
//
// Solidity: function getOperatorSetCount(address avs) view returns(uint256)
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) GetOperatorSetCount(avs common.Address) (*big.Int, error) {
	return _ContractAllocationManager.Contract.GetOperatorSetCount(&_ContractAllocationManager.CallOpts, avs)
}

// GetRegisteredSets is a free data retrieval call binding the contract method 0x79ae50cd.
//
// Solidity: function getRegisteredSets(address operator) view returns((address,uint32)[])
func (_ContractAllocationManager *ContractAllocationManagerCaller) GetRegisteredSets(opts *bind.CallOpts, operator common.Address) ([]OperatorSet, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "getRegisteredSets", operator)

	if err != nil {
		return *new([]OperatorSet), err
	}

	out0 := *abi.ConvertType(out[0], new([]OperatorSet)).(*[]OperatorSet)

	return out0, err

}

// GetRegisteredSets is a free data retrieval call binding the contract method 0x79ae50cd.
//
// Solidity: function getRegisteredSets(address operator) view returns((address,uint32)[])
func (_ContractAllocationManager *ContractAllocationManagerSession) GetRegisteredSets(operator common.Address) ([]OperatorSet, error) {
	return _ContractAllocationManager.Contract.GetRegisteredSets(&_ContractAllocationManager.CallOpts, operator)
}

// GetRegisteredSets is a free data retrieval call binding the contract method 0x79ae50cd.
//
// Solidity: function getRegisteredSets(address operator) view returns((address,uint32)[])
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) GetRegisteredSets(operator common.Address) ([]OperatorSet, error) {
	return _ContractAllocationManager.Contract.GetRegisteredSets(&_ContractAllocationManager.CallOpts, operator)
}

// GetStrategiesInOperatorSet is a free data retrieval call binding the contract method 0x4177a87c.
//
// Solidity: function getStrategiesInOperatorSet((address,uint32) operatorSet) view returns(address[])
func (_ContractAllocationManager *ContractAllocationManagerCaller) GetStrategiesInOperatorSet(opts *bind.CallOpts, operatorSet OperatorSet) ([]common.Address, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "getStrategiesInOperatorSet", operatorSet)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetStrategiesInOperatorSet is a free data retrieval call binding the contract method 0x4177a87c.
//
// Solidity: function getStrategiesInOperatorSet((address,uint32) operatorSet) view returns(address[])
func (_ContractAllocationManager *ContractAllocationManagerSession) GetStrategiesInOperatorSet(operatorSet OperatorSet) ([]common.Address, error) {
	return _ContractAllocationManager.Contract.GetStrategiesInOperatorSet(&_ContractAllocationManager.CallOpts, operatorSet)
}

// GetStrategiesInOperatorSet is a free data retrieval call binding the contract method 0x4177a87c.
//
// Solidity: function getStrategiesInOperatorSet((address,uint32) operatorSet) view returns(address[])
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) GetStrategiesInOperatorSet(operatorSet OperatorSet) ([]common.Address, error) {
	return _ContractAllocationManager.Contract.GetStrategiesInOperatorSet(&_ContractAllocationManager.CallOpts, operatorSet)
}

// GetStrategyAllocations is a free data retrieval call binding the contract method 0x40120dab.
//
// Solidity: function getStrategyAllocations(address operator, address strategy) view returns((address,uint32)[], (uint64,int128,uint32)[])
func (_ContractAllocationManager *ContractAllocationManagerCaller) GetStrategyAllocations(opts *bind.CallOpts, operator common.Address, strategy common.Address) ([]OperatorSet, []IAllocationManagerTypesAllocation, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "getStrategyAllocations", operator, strategy)

	if err != nil {
		return *new([]OperatorSet), *new([]IAllocationManagerTypesAllocation), err
	}

	out0 := *abi.ConvertType(out[0], new([]OperatorSet)).(*[]OperatorSet)
	out1 := *abi.ConvertType(out[1], new([]IAllocationManagerTypesAllocation)).(*[]IAllocationManagerTypesAllocation)

	return out0, out1, err

}

// GetStrategyAllocations is a free data retrieval call binding the contract method 0x40120dab.
//
// Solidity: function getStrategyAllocations(address operator, address strategy) view returns((address,uint32)[], (uint64,int128,uint32)[])
func (_ContractAllocationManager *ContractAllocationManagerSession) GetStrategyAllocations(operator common.Address, strategy common.Address) ([]OperatorSet, []IAllocationManagerTypesAllocation, error) {
	return _ContractAllocationManager.Contract.GetStrategyAllocations(&_ContractAllocationManager.CallOpts, operator, strategy)
}

// GetStrategyAllocations is a free data retrieval call binding the contract method 0x40120dab.
//
// Solidity: function getStrategyAllocations(address operator, address strategy) view returns((address,uint32)[], (uint64,int128,uint32)[])
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) GetStrategyAllocations(operator common.Address, strategy common.Address) ([]OperatorSet, []IAllocationManagerTypesAllocation, error) {
	return _ContractAllocationManager.Contract.GetStrategyAllocations(&_ContractAllocationManager.CallOpts, operator, strategy)
}

// IsMemberOfOperatorSet is a free data retrieval call binding the contract method 0x670d3ba2.
//
// Solidity: function isMemberOfOperatorSet(address operator, (address,uint32) operatorSet) view returns(bool)
func (_ContractAllocationManager *ContractAllocationManagerCaller) IsMemberOfOperatorSet(opts *bind.CallOpts, operator common.Address, operatorSet OperatorSet) (bool, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "isMemberOfOperatorSet", operator, operatorSet)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMemberOfOperatorSet is a free data retrieval call binding the contract method 0x670d3ba2.
//
// Solidity: function isMemberOfOperatorSet(address operator, (address,uint32) operatorSet) view returns(bool)
func (_ContractAllocationManager *ContractAllocationManagerSession) IsMemberOfOperatorSet(operator common.Address, operatorSet OperatorSet) (bool, error) {
	return _ContractAllocationManager.Contract.IsMemberOfOperatorSet(&_ContractAllocationManager.CallOpts, operator, operatorSet)
}

// IsMemberOfOperatorSet is a free data retrieval call binding the contract method 0x670d3ba2.
//
// Solidity: function isMemberOfOperatorSet(address operator, (address,uint32) operatorSet) view returns(bool)
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) IsMemberOfOperatorSet(operator common.Address, operatorSet OperatorSet) (bool, error) {
	return _ContractAllocationManager.Contract.IsMemberOfOperatorSet(&_ContractAllocationManager.CallOpts, operator, operatorSet)
}

// IsOperatorSet is a free data retrieval call binding the contract method 0x260dc758.
//
// Solidity: function isOperatorSet((address,uint32) operatorSet) view returns(bool)
func (_ContractAllocationManager *ContractAllocationManagerCaller) IsOperatorSet(opts *bind.CallOpts, operatorSet OperatorSet) (bool, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "isOperatorSet", operatorSet)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorSet is a free data retrieval call binding the contract method 0x260dc758.
//
// Solidity: function isOperatorSet((address,uint32) operatorSet) view returns(bool)
func (_ContractAllocationManager *ContractAllocationManagerSession) IsOperatorSet(operatorSet OperatorSet) (bool, error) {
	return _ContractAllocationManager.Contract.IsOperatorSet(&_ContractAllocationManager.CallOpts, operatorSet)
}

// IsOperatorSet is a free data retrieval call binding the contract method 0x260dc758.
//
// Solidity: function isOperatorSet((address,uint32) operatorSet) view returns(bool)
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) IsOperatorSet(operatorSet OperatorSet) (bool, error) {
	return _ContractAllocationManager.Contract.IsOperatorSet(&_ContractAllocationManager.CallOpts, operatorSet)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractAllocationManager *ContractAllocationManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractAllocationManager *ContractAllocationManagerSession) Owner() (common.Address, error) {
	return _ContractAllocationManager.Contract.Owner(&_ContractAllocationManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) Owner() (common.Address, error) {
	return _ContractAllocationManager.Contract.Owner(&_ContractAllocationManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractAllocationManager *ContractAllocationManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractAllocationManager *ContractAllocationManagerSession) Paused(index uint8) (bool, error) {
	return _ContractAllocationManager.Contract.Paused(&_ContractAllocationManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) Paused(index uint8) (bool, error) {
	return _ContractAllocationManager.Contract.Paused(&_ContractAllocationManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractAllocationManager *ContractAllocationManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractAllocationManager *ContractAllocationManagerSession) Paused0() (*big.Int, error) {
	return _ContractAllocationManager.Contract.Paused0(&_ContractAllocationManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) Paused0() (*big.Int, error) {
	return _ContractAllocationManager.Contract.Paused0(&_ContractAllocationManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractAllocationManager *ContractAllocationManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractAllocationManager *ContractAllocationManagerSession) PauserRegistry() (common.Address, error) {
	return _ContractAllocationManager.Contract.PauserRegistry(&_ContractAllocationManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractAllocationManager.Contract.PauserRegistry(&_ContractAllocationManager.CallOpts)
}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_ContractAllocationManager *ContractAllocationManagerCaller) PermissionController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "permissionController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_ContractAllocationManager *ContractAllocationManagerSession) PermissionController() (common.Address, error) {
	return _ContractAllocationManager.Contract.PermissionController(&_ContractAllocationManager.CallOpts)
}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) PermissionController() (common.Address, error) {
	return _ContractAllocationManager.Contract.PermissionController(&_ContractAllocationManager.CallOpts)
}

// AddStrategiesToOperatorSet is a paid mutator transaction binding the contract method 0x50feea20.
//
// Solidity: function addStrategiesToOperatorSet(address avs, uint32 operatorSetId, address[] strategies) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactor) AddStrategiesToOperatorSet(opts *bind.TransactOpts, avs common.Address, operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _ContractAllocationManager.contract.Transact(opts, "addStrategiesToOperatorSet", avs, operatorSetId, strategies)
}

// AddStrategiesToOperatorSet is a paid mutator transaction binding the contract method 0x50feea20.
//
// Solidity: function addStrategiesToOperatorSet(address avs, uint32 operatorSetId, address[] strategies) returns()
func (_ContractAllocationManager *ContractAllocationManagerSession) AddStrategiesToOperatorSet(avs common.Address, operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.AddStrategiesToOperatorSet(&_ContractAllocationManager.TransactOpts, avs, operatorSetId, strategies)
}

// AddStrategiesToOperatorSet is a paid mutator transaction binding the contract method 0x50feea20.
//
// Solidity: function addStrategiesToOperatorSet(address avs, uint32 operatorSetId, address[] strategies) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactorSession) AddStrategiesToOperatorSet(avs common.Address, operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.AddStrategiesToOperatorSet(&_ContractAllocationManager.TransactOpts, avs, operatorSetId, strategies)
}

// ClearDeallocationQueue is a paid mutator transaction binding the contract method 0x4b5046ef.
//
// Solidity: function clearDeallocationQueue(address operator, address[] strategies, uint16[] numToClear) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactor) ClearDeallocationQueue(opts *bind.TransactOpts, operator common.Address, strategies []common.Address, numToClear []uint16) (*types.Transaction, error) {
	return _ContractAllocationManager.contract.Transact(opts, "clearDeallocationQueue", operator, strategies, numToClear)
}

// ClearDeallocationQueue is a paid mutator transaction binding the contract method 0x4b5046ef.
//
// Solidity: function clearDeallocationQueue(address operator, address[] strategies, uint16[] numToClear) returns()
func (_ContractAllocationManager *ContractAllocationManagerSession) ClearDeallocationQueue(operator common.Address, strategies []common.Address, numToClear []uint16) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.ClearDeallocationQueue(&_ContractAllocationManager.TransactOpts, operator, strategies, numToClear)
}

// ClearDeallocationQueue is a paid mutator transaction binding the contract method 0x4b5046ef.
//
// Solidity: function clearDeallocationQueue(address operator, address[] strategies, uint16[] numToClear) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactorSession) ClearDeallocationQueue(operator common.Address, strategies []common.Address, numToClear []uint16) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.ClearDeallocationQueue(&_ContractAllocationManager.TransactOpts, operator, strategies, numToClear)
}

// CreateOperatorSets is a paid mutator transaction binding the contract method 0x261f84e0.
//
// Solidity: function createOperatorSets(address avs, (uint32,address[])[] params) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactor) CreateOperatorSets(opts *bind.TransactOpts, avs common.Address, params []IAllocationManagerTypesCreateSetParams) (*types.Transaction, error) {
	return _ContractAllocationManager.contract.Transact(opts, "createOperatorSets", avs, params)
}

// CreateOperatorSets is a paid mutator transaction binding the contract method 0x261f84e0.
//
// Solidity: function createOperatorSets(address avs, (uint32,address[])[] params) returns()
func (_ContractAllocationManager *ContractAllocationManagerSession) CreateOperatorSets(avs common.Address, params []IAllocationManagerTypesCreateSetParams) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.CreateOperatorSets(&_ContractAllocationManager.TransactOpts, avs, params)
}

// CreateOperatorSets is a paid mutator transaction binding the contract method 0x261f84e0.
//
// Solidity: function createOperatorSets(address avs, (uint32,address[])[] params) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactorSession) CreateOperatorSets(avs common.Address, params []IAllocationManagerTypesCreateSetParams) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.CreateOperatorSets(&_ContractAllocationManager.TransactOpts, avs, params)
}

// DeregisterFromOperatorSets is a paid mutator transaction binding the contract method 0x6e3492b5.
//
// Solidity: function deregisterFromOperatorSets((address,address,uint32[]) params) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactor) DeregisterFromOperatorSets(opts *bind.TransactOpts, params IAllocationManagerTypesDeregisterParams) (*types.Transaction, error) {
	return _ContractAllocationManager.contract.Transact(opts, "deregisterFromOperatorSets", params)
}

// DeregisterFromOperatorSets is a paid mutator transaction binding the contract method 0x6e3492b5.
//
// Solidity: function deregisterFromOperatorSets((address,address,uint32[]) params) returns()
func (_ContractAllocationManager *ContractAllocationManagerSession) DeregisterFromOperatorSets(params IAllocationManagerTypesDeregisterParams) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.DeregisterFromOperatorSets(&_ContractAllocationManager.TransactOpts, params)
}

// DeregisterFromOperatorSets is a paid mutator transaction binding the contract method 0x6e3492b5.
//
// Solidity: function deregisterFromOperatorSets((address,address,uint32[]) params) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactorSession) DeregisterFromOperatorSets(params IAllocationManagerTypesDeregisterParams) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.DeregisterFromOperatorSets(&_ContractAllocationManager.TransactOpts, params)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address initialOwner, uint256 initialPausedStatus) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAllocationManager.contract.Transact(opts, "initialize", initialOwner, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address initialOwner, uint256 initialPausedStatus) returns()
func (_ContractAllocationManager *ContractAllocationManagerSession) Initialize(initialOwner common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.Initialize(&_ContractAllocationManager.TransactOpts, initialOwner, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address initialOwner, uint256 initialPausedStatus) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactorSession) Initialize(initialOwner common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.Initialize(&_ContractAllocationManager.TransactOpts, initialOwner, initialPausedStatus)
}

// ModifyAllocations is a paid mutator transaction binding the contract method 0x952899ee.
//
// Solidity: function modifyAllocations(address operator, ((address,uint32),address[],uint64[])[] params) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactor) ModifyAllocations(opts *bind.TransactOpts, operator common.Address, params []IAllocationManagerTypesAllocateParams) (*types.Transaction, error) {
	return _ContractAllocationManager.contract.Transact(opts, "modifyAllocations", operator, params)
}

// ModifyAllocations is a paid mutator transaction binding the contract method 0x952899ee.
//
// Solidity: function modifyAllocations(address operator, ((address,uint32),address[],uint64[])[] params) returns()
func (_ContractAllocationManager *ContractAllocationManagerSession) ModifyAllocations(operator common.Address, params []IAllocationManagerTypesAllocateParams) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.ModifyAllocations(&_ContractAllocationManager.TransactOpts, operator, params)
}

// ModifyAllocations is a paid mutator transaction binding the contract method 0x952899ee.
//
// Solidity: function modifyAllocations(address operator, ((address,uint32),address[],uint64[])[] params) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactorSession) ModifyAllocations(operator common.Address, params []IAllocationManagerTypesAllocateParams) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.ModifyAllocations(&_ContractAllocationManager.TransactOpts, operator, params)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAllocationManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractAllocationManager *ContractAllocationManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.Pause(&_ContractAllocationManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.Pause(&_ContractAllocationManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAllocationManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractAllocationManager *ContractAllocationManagerSession) PauseAll() (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.PauseAll(&_ContractAllocationManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.PauseAll(&_ContractAllocationManager.TransactOpts)
}

// RegisterForOperatorSets is a paid mutator transaction binding the contract method 0xadc2e3d9.
//
// Solidity: function registerForOperatorSets(address operator, (address,uint32[],bytes) params) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactor) RegisterForOperatorSets(opts *bind.TransactOpts, operator common.Address, params IAllocationManagerTypesRegisterParams) (*types.Transaction, error) {
	return _ContractAllocationManager.contract.Transact(opts, "registerForOperatorSets", operator, params)
}

// RegisterForOperatorSets is a paid mutator transaction binding the contract method 0xadc2e3d9.
//
// Solidity: function registerForOperatorSets(address operator, (address,uint32[],bytes) params) returns()
func (_ContractAllocationManager *ContractAllocationManagerSession) RegisterForOperatorSets(operator common.Address, params IAllocationManagerTypesRegisterParams) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.RegisterForOperatorSets(&_ContractAllocationManager.TransactOpts, operator, params)
}

// RegisterForOperatorSets is a paid mutator transaction binding the contract method 0xadc2e3d9.
//
// Solidity: function registerForOperatorSets(address operator, (address,uint32[],bytes) params) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactorSession) RegisterForOperatorSets(operator common.Address, params IAllocationManagerTypesRegisterParams) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.RegisterForOperatorSets(&_ContractAllocationManager.TransactOpts, operator, params)
}

// RemoveStrategiesFromOperatorSet is a paid mutator transaction binding the contract method 0xb66bd989.
//
// Solidity: function removeStrategiesFromOperatorSet(address avs, uint32 operatorSetId, address[] strategies) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactor) RemoveStrategiesFromOperatorSet(opts *bind.TransactOpts, avs common.Address, operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _ContractAllocationManager.contract.Transact(opts, "removeStrategiesFromOperatorSet", avs, operatorSetId, strategies)
}

// RemoveStrategiesFromOperatorSet is a paid mutator transaction binding the contract method 0xb66bd989.
//
// Solidity: function removeStrategiesFromOperatorSet(address avs, uint32 operatorSetId, address[] strategies) returns()
func (_ContractAllocationManager *ContractAllocationManagerSession) RemoveStrategiesFromOperatorSet(avs common.Address, operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.RemoveStrategiesFromOperatorSet(&_ContractAllocationManager.TransactOpts, avs, operatorSetId, strategies)
}

// RemoveStrategiesFromOperatorSet is a paid mutator transaction binding the contract method 0xb66bd989.
//
// Solidity: function removeStrategiesFromOperatorSet(address avs, uint32 operatorSetId, address[] strategies) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactorSession) RemoveStrategiesFromOperatorSet(avs common.Address, operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.RemoveStrategiesFromOperatorSet(&_ContractAllocationManager.TransactOpts, avs, operatorSetId, strategies)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAllocationManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractAllocationManager *ContractAllocationManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.RenounceOwnership(&_ContractAllocationManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.RenounceOwnership(&_ContractAllocationManager.TransactOpts)
}

// SetAVSRegistrar is a paid mutator transaction binding the contract method 0xd3d96ff4.
//
// Solidity: function setAVSRegistrar(address avs, address registrar) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactor) SetAVSRegistrar(opts *bind.TransactOpts, avs common.Address, registrar common.Address) (*types.Transaction, error) {
	return _ContractAllocationManager.contract.Transact(opts, "setAVSRegistrar", avs, registrar)
}

// SetAVSRegistrar is a paid mutator transaction binding the contract method 0xd3d96ff4.
//
// Solidity: function setAVSRegistrar(address avs, address registrar) returns()
func (_ContractAllocationManager *ContractAllocationManagerSession) SetAVSRegistrar(avs common.Address, registrar common.Address) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.SetAVSRegistrar(&_ContractAllocationManager.TransactOpts, avs, registrar)
}

// SetAVSRegistrar is a paid mutator transaction binding the contract method 0xd3d96ff4.
//
// Solidity: function setAVSRegistrar(address avs, address registrar) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactorSession) SetAVSRegistrar(avs common.Address, registrar common.Address) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.SetAVSRegistrar(&_ContractAllocationManager.TransactOpts, avs, registrar)
}

// SetAllocationDelay is a paid mutator transaction binding the contract method 0x56c483e6.
//
// Solidity: function setAllocationDelay(address operator, uint32 delay) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactor) SetAllocationDelay(opts *bind.TransactOpts, operator common.Address, delay uint32) (*types.Transaction, error) {
	return _ContractAllocationManager.contract.Transact(opts, "setAllocationDelay", operator, delay)
}

// SetAllocationDelay is a paid mutator transaction binding the contract method 0x56c483e6.
//
// Solidity: function setAllocationDelay(address operator, uint32 delay) returns()
func (_ContractAllocationManager *ContractAllocationManagerSession) SetAllocationDelay(operator common.Address, delay uint32) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.SetAllocationDelay(&_ContractAllocationManager.TransactOpts, operator, delay)
}

// SetAllocationDelay is a paid mutator transaction binding the contract method 0x56c483e6.
//
// Solidity: function setAllocationDelay(address operator, uint32 delay) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactorSession) SetAllocationDelay(operator common.Address, delay uint32) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.SetAllocationDelay(&_ContractAllocationManager.TransactOpts, operator, delay)
}

// SlashOperator is a paid mutator transaction binding the contract method 0x36352057.
//
// Solidity: function slashOperator(address avs, (address,uint32,address[],uint256[],string) params) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactor) SlashOperator(opts *bind.TransactOpts, avs common.Address, params IAllocationManagerTypesSlashingParams) (*types.Transaction, error) {
	return _ContractAllocationManager.contract.Transact(opts, "slashOperator", avs, params)
}

// SlashOperator is a paid mutator transaction binding the contract method 0x36352057.
//
// Solidity: function slashOperator(address avs, (address,uint32,address[],uint256[],string) params) returns()
func (_ContractAllocationManager *ContractAllocationManagerSession) SlashOperator(avs common.Address, params IAllocationManagerTypesSlashingParams) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.SlashOperator(&_ContractAllocationManager.TransactOpts, avs, params)
}

// SlashOperator is a paid mutator transaction binding the contract method 0x36352057.
//
// Solidity: function slashOperator(address avs, (address,uint32,address[],uint256[],string) params) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactorSession) SlashOperator(avs common.Address, params IAllocationManagerTypesSlashingParams) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.SlashOperator(&_ContractAllocationManager.TransactOpts, avs, params)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractAllocationManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractAllocationManager *ContractAllocationManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.TransferOwnership(&_ContractAllocationManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.TransferOwnership(&_ContractAllocationManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAllocationManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractAllocationManager *ContractAllocationManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.Unpause(&_ContractAllocationManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.Unpause(&_ContractAllocationManager.TransactOpts, newPausedStatus)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa9821821.
//
// Solidity: function updateAVSMetadataURI(address avs, string metadataURI) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, avs common.Address, metadataURI string) (*types.Transaction, error) {
	return _ContractAllocationManager.contract.Transact(opts, "updateAVSMetadataURI", avs, metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa9821821.
//
// Solidity: function updateAVSMetadataURI(address avs, string metadataURI) returns()
func (_ContractAllocationManager *ContractAllocationManagerSession) UpdateAVSMetadataURI(avs common.Address, metadataURI string) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.UpdateAVSMetadataURI(&_ContractAllocationManager.TransactOpts, avs, metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa9821821.
//
// Solidity: function updateAVSMetadataURI(address avs, string metadataURI) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactorSession) UpdateAVSMetadataURI(avs common.Address, metadataURI string) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.UpdateAVSMetadataURI(&_ContractAllocationManager.TransactOpts, avs, metadataURI)
}

// ContractAllocationManagerAVSMetadataURIUpdatedIterator is returned from FilterAVSMetadataURIUpdated and is used to iterate over the raw logs and unpacked data for AVSMetadataURIUpdated events raised by the ContractAllocationManager contract.
type ContractAllocationManagerAVSMetadataURIUpdatedIterator struct {
	Event *ContractAllocationManagerAVSMetadataURIUpdated // Event containing the contract specifics and raw log

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
func (it *ContractAllocationManagerAVSMetadataURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAllocationManagerAVSMetadataURIUpdated)
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
		it.Event = new(ContractAllocationManagerAVSMetadataURIUpdated)
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
func (it *ContractAllocationManagerAVSMetadataURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAllocationManagerAVSMetadataURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAllocationManagerAVSMetadataURIUpdated represents a AVSMetadataURIUpdated event raised by the ContractAllocationManager contract.
type ContractAllocationManagerAVSMetadataURIUpdated struct {
	Avs         common.Address
	MetadataURI string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAVSMetadataURIUpdated is a free log retrieval operation binding the contract event 0xa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c943713.
//
// Solidity: event AVSMetadataURIUpdated(address indexed avs, string metadataURI)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) FilterAVSMetadataURIUpdated(opts *bind.FilterOpts, avs []common.Address) (*ContractAllocationManagerAVSMetadataURIUpdatedIterator, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ContractAllocationManager.contract.FilterLogs(opts, "AVSMetadataURIUpdated", avsRule)
	if err != nil {
		return nil, err
	}
	return &ContractAllocationManagerAVSMetadataURIUpdatedIterator{contract: _ContractAllocationManager.contract, event: "AVSMetadataURIUpdated", logs: logs, sub: sub}, nil
}

// WatchAVSMetadataURIUpdated is a free log subscription operation binding the contract event 0xa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c943713.
//
// Solidity: event AVSMetadataURIUpdated(address indexed avs, string metadataURI)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) WatchAVSMetadataURIUpdated(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerAVSMetadataURIUpdated, avs []common.Address) (event.Subscription, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ContractAllocationManager.contract.WatchLogs(opts, "AVSMetadataURIUpdated", avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAllocationManagerAVSMetadataURIUpdated)
				if err := _ContractAllocationManager.contract.UnpackLog(event, "AVSMetadataURIUpdated", log); err != nil {
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

// ParseAVSMetadataURIUpdated is a log parse operation binding the contract event 0xa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c943713.
//
// Solidity: event AVSMetadataURIUpdated(address indexed avs, string metadataURI)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) ParseAVSMetadataURIUpdated(log types.Log) (*ContractAllocationManagerAVSMetadataURIUpdated, error) {
	event := new(ContractAllocationManagerAVSMetadataURIUpdated)
	if err := _ContractAllocationManager.contract.UnpackLog(event, "AVSMetadataURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAllocationManagerAVSRegistrarSetIterator is returned from FilterAVSRegistrarSet and is used to iterate over the raw logs and unpacked data for AVSRegistrarSet events raised by the ContractAllocationManager contract.
type ContractAllocationManagerAVSRegistrarSetIterator struct {
	Event *ContractAllocationManagerAVSRegistrarSet // Event containing the contract specifics and raw log

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
func (it *ContractAllocationManagerAVSRegistrarSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAllocationManagerAVSRegistrarSet)
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
		it.Event = new(ContractAllocationManagerAVSRegistrarSet)
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
func (it *ContractAllocationManagerAVSRegistrarSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAllocationManagerAVSRegistrarSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAllocationManagerAVSRegistrarSet represents a AVSRegistrarSet event raised by the ContractAllocationManager contract.
type ContractAllocationManagerAVSRegistrarSet struct {
	Avs       common.Address
	Registrar common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAVSRegistrarSet is a free log retrieval operation binding the contract event 0x2ae945c40c44dc0ec263f95609c3fdc6952e0aefa22d6374e44f2c997acedf85.
//
// Solidity: event AVSRegistrarSet(address avs, address registrar)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) FilterAVSRegistrarSet(opts *bind.FilterOpts) (*ContractAllocationManagerAVSRegistrarSetIterator, error) {

	logs, sub, err := _ContractAllocationManager.contract.FilterLogs(opts, "AVSRegistrarSet")
	if err != nil {
		return nil, err
	}
	return &ContractAllocationManagerAVSRegistrarSetIterator{contract: _ContractAllocationManager.contract, event: "AVSRegistrarSet", logs: logs, sub: sub}, nil
}

// WatchAVSRegistrarSet is a free log subscription operation binding the contract event 0x2ae945c40c44dc0ec263f95609c3fdc6952e0aefa22d6374e44f2c997acedf85.
//
// Solidity: event AVSRegistrarSet(address avs, address registrar)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) WatchAVSRegistrarSet(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerAVSRegistrarSet) (event.Subscription, error) {

	logs, sub, err := _ContractAllocationManager.contract.WatchLogs(opts, "AVSRegistrarSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAllocationManagerAVSRegistrarSet)
				if err := _ContractAllocationManager.contract.UnpackLog(event, "AVSRegistrarSet", log); err != nil {
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

// ParseAVSRegistrarSet is a log parse operation binding the contract event 0x2ae945c40c44dc0ec263f95609c3fdc6952e0aefa22d6374e44f2c997acedf85.
//
// Solidity: event AVSRegistrarSet(address avs, address registrar)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) ParseAVSRegistrarSet(log types.Log) (*ContractAllocationManagerAVSRegistrarSet, error) {
	event := new(ContractAllocationManagerAVSRegistrarSet)
	if err := _ContractAllocationManager.contract.UnpackLog(event, "AVSRegistrarSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAllocationManagerAllocationDelaySetIterator is returned from FilterAllocationDelaySet and is used to iterate over the raw logs and unpacked data for AllocationDelaySet events raised by the ContractAllocationManager contract.
type ContractAllocationManagerAllocationDelaySetIterator struct {
	Event *ContractAllocationManagerAllocationDelaySet // Event containing the contract specifics and raw log

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
func (it *ContractAllocationManagerAllocationDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAllocationManagerAllocationDelaySet)
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
		it.Event = new(ContractAllocationManagerAllocationDelaySet)
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
func (it *ContractAllocationManagerAllocationDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAllocationManagerAllocationDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAllocationManagerAllocationDelaySet represents a AllocationDelaySet event raised by the ContractAllocationManager contract.
type ContractAllocationManagerAllocationDelaySet struct {
	Operator    common.Address
	Delay       uint32
	EffectBlock uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAllocationDelaySet is a free log retrieval operation binding the contract event 0x4e85751d6331506c6c62335f207eb31f12a61e570f34f5c17640308785c6d4db.
//
// Solidity: event AllocationDelaySet(address operator, uint32 delay, uint32 effectBlock)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) FilterAllocationDelaySet(opts *bind.FilterOpts) (*ContractAllocationManagerAllocationDelaySetIterator, error) {

	logs, sub, err := _ContractAllocationManager.contract.FilterLogs(opts, "AllocationDelaySet")
	if err != nil {
		return nil, err
	}
	return &ContractAllocationManagerAllocationDelaySetIterator{contract: _ContractAllocationManager.contract, event: "AllocationDelaySet", logs: logs, sub: sub}, nil
}

// WatchAllocationDelaySet is a free log subscription operation binding the contract event 0x4e85751d6331506c6c62335f207eb31f12a61e570f34f5c17640308785c6d4db.
//
// Solidity: event AllocationDelaySet(address operator, uint32 delay, uint32 effectBlock)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) WatchAllocationDelaySet(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerAllocationDelaySet) (event.Subscription, error) {

	logs, sub, err := _ContractAllocationManager.contract.WatchLogs(opts, "AllocationDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAllocationManagerAllocationDelaySet)
				if err := _ContractAllocationManager.contract.UnpackLog(event, "AllocationDelaySet", log); err != nil {
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

// ParseAllocationDelaySet is a log parse operation binding the contract event 0x4e85751d6331506c6c62335f207eb31f12a61e570f34f5c17640308785c6d4db.
//
// Solidity: event AllocationDelaySet(address operator, uint32 delay, uint32 effectBlock)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) ParseAllocationDelaySet(log types.Log) (*ContractAllocationManagerAllocationDelaySet, error) {
	event := new(ContractAllocationManagerAllocationDelaySet)
	if err := _ContractAllocationManager.contract.UnpackLog(event, "AllocationDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAllocationManagerAllocationUpdatedIterator is returned from FilterAllocationUpdated and is used to iterate over the raw logs and unpacked data for AllocationUpdated events raised by the ContractAllocationManager contract.
type ContractAllocationManagerAllocationUpdatedIterator struct {
	Event *ContractAllocationManagerAllocationUpdated // Event containing the contract specifics and raw log

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
func (it *ContractAllocationManagerAllocationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAllocationManagerAllocationUpdated)
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
		it.Event = new(ContractAllocationManagerAllocationUpdated)
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
func (it *ContractAllocationManagerAllocationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAllocationManagerAllocationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAllocationManagerAllocationUpdated represents a AllocationUpdated event raised by the ContractAllocationManager contract.
type ContractAllocationManagerAllocationUpdated struct {
	Operator    common.Address
	OperatorSet OperatorSet
	Strategy    common.Address
	Magnitude   uint64
	EffectBlock uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAllocationUpdated is a free log retrieval operation binding the contract event 0x1487af5418c47ee5ea45ef4a93398668120890774a9e13487e61e9dc3baf76dd.
//
// Solidity: event AllocationUpdated(address operator, (address,uint32) operatorSet, address strategy, uint64 magnitude, uint32 effectBlock)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) FilterAllocationUpdated(opts *bind.FilterOpts) (*ContractAllocationManagerAllocationUpdatedIterator, error) {

	logs, sub, err := _ContractAllocationManager.contract.FilterLogs(opts, "AllocationUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractAllocationManagerAllocationUpdatedIterator{contract: _ContractAllocationManager.contract, event: "AllocationUpdated", logs: logs, sub: sub}, nil
}

// WatchAllocationUpdated is a free log subscription operation binding the contract event 0x1487af5418c47ee5ea45ef4a93398668120890774a9e13487e61e9dc3baf76dd.
//
// Solidity: event AllocationUpdated(address operator, (address,uint32) operatorSet, address strategy, uint64 magnitude, uint32 effectBlock)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) WatchAllocationUpdated(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerAllocationUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractAllocationManager.contract.WatchLogs(opts, "AllocationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAllocationManagerAllocationUpdated)
				if err := _ContractAllocationManager.contract.UnpackLog(event, "AllocationUpdated", log); err != nil {
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

// ParseAllocationUpdated is a log parse operation binding the contract event 0x1487af5418c47ee5ea45ef4a93398668120890774a9e13487e61e9dc3baf76dd.
//
// Solidity: event AllocationUpdated(address operator, (address,uint32) operatorSet, address strategy, uint64 magnitude, uint32 effectBlock)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) ParseAllocationUpdated(log types.Log) (*ContractAllocationManagerAllocationUpdated, error) {
	event := new(ContractAllocationManagerAllocationUpdated)
	if err := _ContractAllocationManager.contract.UnpackLog(event, "AllocationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAllocationManagerEncumberedMagnitudeUpdatedIterator is returned from FilterEncumberedMagnitudeUpdated and is used to iterate over the raw logs and unpacked data for EncumberedMagnitudeUpdated events raised by the ContractAllocationManager contract.
type ContractAllocationManagerEncumberedMagnitudeUpdatedIterator struct {
	Event *ContractAllocationManagerEncumberedMagnitudeUpdated // Event containing the contract specifics and raw log

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
func (it *ContractAllocationManagerEncumberedMagnitudeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAllocationManagerEncumberedMagnitudeUpdated)
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
		it.Event = new(ContractAllocationManagerEncumberedMagnitudeUpdated)
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
func (it *ContractAllocationManagerEncumberedMagnitudeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAllocationManagerEncumberedMagnitudeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAllocationManagerEncumberedMagnitudeUpdated represents a EncumberedMagnitudeUpdated event raised by the ContractAllocationManager contract.
type ContractAllocationManagerEncumberedMagnitudeUpdated struct {
	Operator            common.Address
	Strategy            common.Address
	EncumberedMagnitude uint64
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterEncumberedMagnitudeUpdated is a free log retrieval operation binding the contract event 0xacf9095feb3a370c9cf692421c69ef320d4db5c66e6a7d29c7694eb02364fc55.
//
// Solidity: event EncumberedMagnitudeUpdated(address operator, address strategy, uint64 encumberedMagnitude)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) FilterEncumberedMagnitudeUpdated(opts *bind.FilterOpts) (*ContractAllocationManagerEncumberedMagnitudeUpdatedIterator, error) {

	logs, sub, err := _ContractAllocationManager.contract.FilterLogs(opts, "EncumberedMagnitudeUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractAllocationManagerEncumberedMagnitudeUpdatedIterator{contract: _ContractAllocationManager.contract, event: "EncumberedMagnitudeUpdated", logs: logs, sub: sub}, nil
}

// WatchEncumberedMagnitudeUpdated is a free log subscription operation binding the contract event 0xacf9095feb3a370c9cf692421c69ef320d4db5c66e6a7d29c7694eb02364fc55.
//
// Solidity: event EncumberedMagnitudeUpdated(address operator, address strategy, uint64 encumberedMagnitude)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) WatchEncumberedMagnitudeUpdated(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerEncumberedMagnitudeUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractAllocationManager.contract.WatchLogs(opts, "EncumberedMagnitudeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAllocationManagerEncumberedMagnitudeUpdated)
				if err := _ContractAllocationManager.contract.UnpackLog(event, "EncumberedMagnitudeUpdated", log); err != nil {
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

// ParseEncumberedMagnitudeUpdated is a log parse operation binding the contract event 0xacf9095feb3a370c9cf692421c69ef320d4db5c66e6a7d29c7694eb02364fc55.
//
// Solidity: event EncumberedMagnitudeUpdated(address operator, address strategy, uint64 encumberedMagnitude)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) ParseEncumberedMagnitudeUpdated(log types.Log) (*ContractAllocationManagerEncumberedMagnitudeUpdated, error) {
	event := new(ContractAllocationManagerEncumberedMagnitudeUpdated)
	if err := _ContractAllocationManager.contract.UnpackLog(event, "EncumberedMagnitudeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAllocationManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractAllocationManager contract.
type ContractAllocationManagerInitializedIterator struct {
	Event *ContractAllocationManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ContractAllocationManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAllocationManagerInitialized)
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
		it.Event = new(ContractAllocationManagerInitialized)
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
func (it *ContractAllocationManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAllocationManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAllocationManagerInitialized represents a Initialized event raised by the ContractAllocationManager contract.
type ContractAllocationManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractAllocationManagerInitializedIterator, error) {

	logs, sub, err := _ContractAllocationManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractAllocationManagerInitializedIterator{contract: _ContractAllocationManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractAllocationManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAllocationManagerInitialized)
				if err := _ContractAllocationManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractAllocationManager *ContractAllocationManagerFilterer) ParseInitialized(log types.Log) (*ContractAllocationManagerInitialized, error) {
	event := new(ContractAllocationManagerInitialized)
	if err := _ContractAllocationManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAllocationManagerMaxMagnitudeUpdatedIterator is returned from FilterMaxMagnitudeUpdated and is used to iterate over the raw logs and unpacked data for MaxMagnitudeUpdated events raised by the ContractAllocationManager contract.
type ContractAllocationManagerMaxMagnitudeUpdatedIterator struct {
	Event *ContractAllocationManagerMaxMagnitudeUpdated // Event containing the contract specifics and raw log

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
func (it *ContractAllocationManagerMaxMagnitudeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAllocationManagerMaxMagnitudeUpdated)
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
		it.Event = new(ContractAllocationManagerMaxMagnitudeUpdated)
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
func (it *ContractAllocationManagerMaxMagnitudeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAllocationManagerMaxMagnitudeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAllocationManagerMaxMagnitudeUpdated represents a MaxMagnitudeUpdated event raised by the ContractAllocationManager contract.
type ContractAllocationManagerMaxMagnitudeUpdated struct {
	Operator     common.Address
	Strategy     common.Address
	MaxMagnitude uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMaxMagnitudeUpdated is a free log retrieval operation binding the contract event 0x1c6458079a41077d003c11faf9bf097e693bd67979e4e6500bac7b29db779b5c.
//
// Solidity: event MaxMagnitudeUpdated(address operator, address strategy, uint64 maxMagnitude)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) FilterMaxMagnitudeUpdated(opts *bind.FilterOpts) (*ContractAllocationManagerMaxMagnitudeUpdatedIterator, error) {

	logs, sub, err := _ContractAllocationManager.contract.FilterLogs(opts, "MaxMagnitudeUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractAllocationManagerMaxMagnitudeUpdatedIterator{contract: _ContractAllocationManager.contract, event: "MaxMagnitudeUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxMagnitudeUpdated is a free log subscription operation binding the contract event 0x1c6458079a41077d003c11faf9bf097e693bd67979e4e6500bac7b29db779b5c.
//
// Solidity: event MaxMagnitudeUpdated(address operator, address strategy, uint64 maxMagnitude)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) WatchMaxMagnitudeUpdated(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerMaxMagnitudeUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractAllocationManager.contract.WatchLogs(opts, "MaxMagnitudeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAllocationManagerMaxMagnitudeUpdated)
				if err := _ContractAllocationManager.contract.UnpackLog(event, "MaxMagnitudeUpdated", log); err != nil {
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

// ParseMaxMagnitudeUpdated is a log parse operation binding the contract event 0x1c6458079a41077d003c11faf9bf097e693bd67979e4e6500bac7b29db779b5c.
//
// Solidity: event MaxMagnitudeUpdated(address operator, address strategy, uint64 maxMagnitude)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) ParseMaxMagnitudeUpdated(log types.Log) (*ContractAllocationManagerMaxMagnitudeUpdated, error) {
	event := new(ContractAllocationManagerMaxMagnitudeUpdated)
	if err := _ContractAllocationManager.contract.UnpackLog(event, "MaxMagnitudeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAllocationManagerOperatorAddedToOperatorSetIterator is returned from FilterOperatorAddedToOperatorSet and is used to iterate over the raw logs and unpacked data for OperatorAddedToOperatorSet events raised by the ContractAllocationManager contract.
type ContractAllocationManagerOperatorAddedToOperatorSetIterator struct {
	Event *ContractAllocationManagerOperatorAddedToOperatorSet // Event containing the contract specifics and raw log

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
func (it *ContractAllocationManagerOperatorAddedToOperatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAllocationManagerOperatorAddedToOperatorSet)
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
		it.Event = new(ContractAllocationManagerOperatorAddedToOperatorSet)
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
func (it *ContractAllocationManagerOperatorAddedToOperatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAllocationManagerOperatorAddedToOperatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAllocationManagerOperatorAddedToOperatorSet represents a OperatorAddedToOperatorSet event raised by the ContractAllocationManager contract.
type ContractAllocationManagerOperatorAddedToOperatorSet struct {
	Operator    common.Address
	OperatorSet OperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorAddedToOperatorSet is a free log retrieval operation binding the contract event 0x43232edf9071753d2321e5fa7e018363ee248e5f2142e6c08edd3265bfb4895e.
//
// Solidity: event OperatorAddedToOperatorSet(address indexed operator, (address,uint32) operatorSet)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) FilterOperatorAddedToOperatorSet(opts *bind.FilterOpts, operator []common.Address) (*ContractAllocationManagerOperatorAddedToOperatorSetIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAllocationManager.contract.FilterLogs(opts, "OperatorAddedToOperatorSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractAllocationManagerOperatorAddedToOperatorSetIterator{contract: _ContractAllocationManager.contract, event: "OperatorAddedToOperatorSet", logs: logs, sub: sub}, nil
}

// WatchOperatorAddedToOperatorSet is a free log subscription operation binding the contract event 0x43232edf9071753d2321e5fa7e018363ee248e5f2142e6c08edd3265bfb4895e.
//
// Solidity: event OperatorAddedToOperatorSet(address indexed operator, (address,uint32) operatorSet)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) WatchOperatorAddedToOperatorSet(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerOperatorAddedToOperatorSet, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAllocationManager.contract.WatchLogs(opts, "OperatorAddedToOperatorSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAllocationManagerOperatorAddedToOperatorSet)
				if err := _ContractAllocationManager.contract.UnpackLog(event, "OperatorAddedToOperatorSet", log); err != nil {
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

// ParseOperatorAddedToOperatorSet is a log parse operation binding the contract event 0x43232edf9071753d2321e5fa7e018363ee248e5f2142e6c08edd3265bfb4895e.
//
// Solidity: event OperatorAddedToOperatorSet(address indexed operator, (address,uint32) operatorSet)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) ParseOperatorAddedToOperatorSet(log types.Log) (*ContractAllocationManagerOperatorAddedToOperatorSet, error) {
	event := new(ContractAllocationManagerOperatorAddedToOperatorSet)
	if err := _ContractAllocationManager.contract.UnpackLog(event, "OperatorAddedToOperatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAllocationManagerOperatorRemovedFromOperatorSetIterator is returned from FilterOperatorRemovedFromOperatorSet and is used to iterate over the raw logs and unpacked data for OperatorRemovedFromOperatorSet events raised by the ContractAllocationManager contract.
type ContractAllocationManagerOperatorRemovedFromOperatorSetIterator struct {
	Event *ContractAllocationManagerOperatorRemovedFromOperatorSet // Event containing the contract specifics and raw log

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
func (it *ContractAllocationManagerOperatorRemovedFromOperatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAllocationManagerOperatorRemovedFromOperatorSet)
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
		it.Event = new(ContractAllocationManagerOperatorRemovedFromOperatorSet)
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
func (it *ContractAllocationManagerOperatorRemovedFromOperatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAllocationManagerOperatorRemovedFromOperatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAllocationManagerOperatorRemovedFromOperatorSet represents a OperatorRemovedFromOperatorSet event raised by the ContractAllocationManager contract.
type ContractAllocationManagerOperatorRemovedFromOperatorSet struct {
	Operator    common.Address
	OperatorSet OperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorRemovedFromOperatorSet is a free log retrieval operation binding the contract event 0xad34c3070be1dffbcaa499d000ba2b8d9848aefcac3059df245dd95c4ece14fe.
//
// Solidity: event OperatorRemovedFromOperatorSet(address indexed operator, (address,uint32) operatorSet)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) FilterOperatorRemovedFromOperatorSet(opts *bind.FilterOpts, operator []common.Address) (*ContractAllocationManagerOperatorRemovedFromOperatorSetIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAllocationManager.contract.FilterLogs(opts, "OperatorRemovedFromOperatorSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractAllocationManagerOperatorRemovedFromOperatorSetIterator{contract: _ContractAllocationManager.contract, event: "OperatorRemovedFromOperatorSet", logs: logs, sub: sub}, nil
}

// WatchOperatorRemovedFromOperatorSet is a free log subscription operation binding the contract event 0xad34c3070be1dffbcaa499d000ba2b8d9848aefcac3059df245dd95c4ece14fe.
//
// Solidity: event OperatorRemovedFromOperatorSet(address indexed operator, (address,uint32) operatorSet)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) WatchOperatorRemovedFromOperatorSet(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerOperatorRemovedFromOperatorSet, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAllocationManager.contract.WatchLogs(opts, "OperatorRemovedFromOperatorSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAllocationManagerOperatorRemovedFromOperatorSet)
				if err := _ContractAllocationManager.contract.UnpackLog(event, "OperatorRemovedFromOperatorSet", log); err != nil {
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

// ParseOperatorRemovedFromOperatorSet is a log parse operation binding the contract event 0xad34c3070be1dffbcaa499d000ba2b8d9848aefcac3059df245dd95c4ece14fe.
//
// Solidity: event OperatorRemovedFromOperatorSet(address indexed operator, (address,uint32) operatorSet)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) ParseOperatorRemovedFromOperatorSet(log types.Log) (*ContractAllocationManagerOperatorRemovedFromOperatorSet, error) {
	event := new(ContractAllocationManagerOperatorRemovedFromOperatorSet)
	if err := _ContractAllocationManager.contract.UnpackLog(event, "OperatorRemovedFromOperatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAllocationManagerOperatorSetCreatedIterator is returned from FilterOperatorSetCreated and is used to iterate over the raw logs and unpacked data for OperatorSetCreated events raised by the ContractAllocationManager contract.
type ContractAllocationManagerOperatorSetCreatedIterator struct {
	Event *ContractAllocationManagerOperatorSetCreated // Event containing the contract specifics and raw log

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
func (it *ContractAllocationManagerOperatorSetCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAllocationManagerOperatorSetCreated)
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
		it.Event = new(ContractAllocationManagerOperatorSetCreated)
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
func (it *ContractAllocationManagerOperatorSetCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAllocationManagerOperatorSetCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAllocationManagerOperatorSetCreated represents a OperatorSetCreated event raised by the ContractAllocationManager contract.
type ContractAllocationManagerOperatorSetCreated struct {
	OperatorSet OperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetCreated is a free log retrieval operation binding the contract event 0x31629285ead2335ae0933f86ed2ae63321f7af77b4e6eaabc42c057880977e6c.
//
// Solidity: event OperatorSetCreated((address,uint32) operatorSet)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) FilterOperatorSetCreated(opts *bind.FilterOpts) (*ContractAllocationManagerOperatorSetCreatedIterator, error) {

	logs, sub, err := _ContractAllocationManager.contract.FilterLogs(opts, "OperatorSetCreated")
	if err != nil {
		return nil, err
	}
	return &ContractAllocationManagerOperatorSetCreatedIterator{contract: _ContractAllocationManager.contract, event: "OperatorSetCreated", logs: logs, sub: sub}, nil
}

// WatchOperatorSetCreated is a free log subscription operation binding the contract event 0x31629285ead2335ae0933f86ed2ae63321f7af77b4e6eaabc42c057880977e6c.
//
// Solidity: event OperatorSetCreated((address,uint32) operatorSet)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) WatchOperatorSetCreated(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerOperatorSetCreated) (event.Subscription, error) {

	logs, sub, err := _ContractAllocationManager.contract.WatchLogs(opts, "OperatorSetCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAllocationManagerOperatorSetCreated)
				if err := _ContractAllocationManager.contract.UnpackLog(event, "OperatorSetCreated", log); err != nil {
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

// ParseOperatorSetCreated is a log parse operation binding the contract event 0x31629285ead2335ae0933f86ed2ae63321f7af77b4e6eaabc42c057880977e6c.
//
// Solidity: event OperatorSetCreated((address,uint32) operatorSet)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) ParseOperatorSetCreated(log types.Log) (*ContractAllocationManagerOperatorSetCreated, error) {
	event := new(ContractAllocationManagerOperatorSetCreated)
	if err := _ContractAllocationManager.contract.UnpackLog(event, "OperatorSetCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAllocationManagerOperatorSlashedIterator is returned from FilterOperatorSlashed and is used to iterate over the raw logs and unpacked data for OperatorSlashed events raised by the ContractAllocationManager contract.
type ContractAllocationManagerOperatorSlashedIterator struct {
	Event *ContractAllocationManagerOperatorSlashed // Event containing the contract specifics and raw log

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
func (it *ContractAllocationManagerOperatorSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAllocationManagerOperatorSlashed)
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
		it.Event = new(ContractAllocationManagerOperatorSlashed)
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
func (it *ContractAllocationManagerOperatorSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAllocationManagerOperatorSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAllocationManagerOperatorSlashed represents a OperatorSlashed event raised by the ContractAllocationManager contract.
type ContractAllocationManagerOperatorSlashed struct {
	Operator    common.Address
	OperatorSet OperatorSet
	Strategies  []common.Address
	WadSlashed  []*big.Int
	Description string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorSlashed is a free log retrieval operation binding the contract event 0x80969ad29428d6797ee7aad084f9e4a42a82fc506dcd2ca3b6fb431f85ccebe5.
//
// Solidity: event OperatorSlashed(address operator, (address,uint32) operatorSet, address[] strategies, uint256[] wadSlashed, string description)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) FilterOperatorSlashed(opts *bind.FilterOpts) (*ContractAllocationManagerOperatorSlashedIterator, error) {

	logs, sub, err := _ContractAllocationManager.contract.FilterLogs(opts, "OperatorSlashed")
	if err != nil {
		return nil, err
	}
	return &ContractAllocationManagerOperatorSlashedIterator{contract: _ContractAllocationManager.contract, event: "OperatorSlashed", logs: logs, sub: sub}, nil
}

// WatchOperatorSlashed is a free log subscription operation binding the contract event 0x80969ad29428d6797ee7aad084f9e4a42a82fc506dcd2ca3b6fb431f85ccebe5.
//
// Solidity: event OperatorSlashed(address operator, (address,uint32) operatorSet, address[] strategies, uint256[] wadSlashed, string description)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) WatchOperatorSlashed(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerOperatorSlashed) (event.Subscription, error) {

	logs, sub, err := _ContractAllocationManager.contract.WatchLogs(opts, "OperatorSlashed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAllocationManagerOperatorSlashed)
				if err := _ContractAllocationManager.contract.UnpackLog(event, "OperatorSlashed", log); err != nil {
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

// ParseOperatorSlashed is a log parse operation binding the contract event 0x80969ad29428d6797ee7aad084f9e4a42a82fc506dcd2ca3b6fb431f85ccebe5.
//
// Solidity: event OperatorSlashed(address operator, (address,uint32) operatorSet, address[] strategies, uint256[] wadSlashed, string description)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) ParseOperatorSlashed(log types.Log) (*ContractAllocationManagerOperatorSlashed, error) {
	event := new(ContractAllocationManagerOperatorSlashed)
	if err := _ContractAllocationManager.contract.UnpackLog(event, "OperatorSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAllocationManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractAllocationManager contract.
type ContractAllocationManagerOwnershipTransferredIterator struct {
	Event *ContractAllocationManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractAllocationManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAllocationManagerOwnershipTransferred)
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
		it.Event = new(ContractAllocationManagerOwnershipTransferred)
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
func (it *ContractAllocationManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAllocationManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAllocationManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractAllocationManager contract.
type ContractAllocationManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractAllocationManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractAllocationManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractAllocationManagerOwnershipTransferredIterator{contract: _ContractAllocationManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractAllocationManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAllocationManagerOwnershipTransferred)
				if err := _ContractAllocationManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractAllocationManager *ContractAllocationManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractAllocationManagerOwnershipTransferred, error) {
	event := new(ContractAllocationManagerOwnershipTransferred)
	if err := _ContractAllocationManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAllocationManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractAllocationManager contract.
type ContractAllocationManagerPausedIterator struct {
	Event *ContractAllocationManagerPaused // Event containing the contract specifics and raw log

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
func (it *ContractAllocationManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAllocationManagerPaused)
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
		it.Event = new(ContractAllocationManagerPaused)
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
func (it *ContractAllocationManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAllocationManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAllocationManagerPaused represents a Paused event raised by the ContractAllocationManager contract.
type ContractAllocationManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractAllocationManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractAllocationManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractAllocationManagerPausedIterator{contract: _ContractAllocationManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractAllocationManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAllocationManagerPaused)
				if err := _ContractAllocationManager.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ContractAllocationManager *ContractAllocationManagerFilterer) ParsePaused(log types.Log) (*ContractAllocationManagerPaused, error) {
	event := new(ContractAllocationManagerPaused)
	if err := _ContractAllocationManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAllocationManagerStrategyAddedToOperatorSetIterator is returned from FilterStrategyAddedToOperatorSet and is used to iterate over the raw logs and unpacked data for StrategyAddedToOperatorSet events raised by the ContractAllocationManager contract.
type ContractAllocationManagerStrategyAddedToOperatorSetIterator struct {
	Event *ContractAllocationManagerStrategyAddedToOperatorSet // Event containing the contract specifics and raw log

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
func (it *ContractAllocationManagerStrategyAddedToOperatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAllocationManagerStrategyAddedToOperatorSet)
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
		it.Event = new(ContractAllocationManagerStrategyAddedToOperatorSet)
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
func (it *ContractAllocationManagerStrategyAddedToOperatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAllocationManagerStrategyAddedToOperatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAllocationManagerStrategyAddedToOperatorSet represents a StrategyAddedToOperatorSet event raised by the ContractAllocationManager contract.
type ContractAllocationManagerStrategyAddedToOperatorSet struct {
	OperatorSet OperatorSet
	Strategy    common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStrategyAddedToOperatorSet is a free log retrieval operation binding the contract event 0x7ab260fe0af193db5f4986770d831bda4ea46099dc817e8b6716dcae8af8e88b.
//
// Solidity: event StrategyAddedToOperatorSet((address,uint32) operatorSet, address strategy)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) FilterStrategyAddedToOperatorSet(opts *bind.FilterOpts) (*ContractAllocationManagerStrategyAddedToOperatorSetIterator, error) {

	logs, sub, err := _ContractAllocationManager.contract.FilterLogs(opts, "StrategyAddedToOperatorSet")
	if err != nil {
		return nil, err
	}
	return &ContractAllocationManagerStrategyAddedToOperatorSetIterator{contract: _ContractAllocationManager.contract, event: "StrategyAddedToOperatorSet", logs: logs, sub: sub}, nil
}

// WatchStrategyAddedToOperatorSet is a free log subscription operation binding the contract event 0x7ab260fe0af193db5f4986770d831bda4ea46099dc817e8b6716dcae8af8e88b.
//
// Solidity: event StrategyAddedToOperatorSet((address,uint32) operatorSet, address strategy)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) WatchStrategyAddedToOperatorSet(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerStrategyAddedToOperatorSet) (event.Subscription, error) {

	logs, sub, err := _ContractAllocationManager.contract.WatchLogs(opts, "StrategyAddedToOperatorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAllocationManagerStrategyAddedToOperatorSet)
				if err := _ContractAllocationManager.contract.UnpackLog(event, "StrategyAddedToOperatorSet", log); err != nil {
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

// ParseStrategyAddedToOperatorSet is a log parse operation binding the contract event 0x7ab260fe0af193db5f4986770d831bda4ea46099dc817e8b6716dcae8af8e88b.
//
// Solidity: event StrategyAddedToOperatorSet((address,uint32) operatorSet, address strategy)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) ParseStrategyAddedToOperatorSet(log types.Log) (*ContractAllocationManagerStrategyAddedToOperatorSet, error) {
	event := new(ContractAllocationManagerStrategyAddedToOperatorSet)
	if err := _ContractAllocationManager.contract.UnpackLog(event, "StrategyAddedToOperatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAllocationManagerStrategyRemovedFromOperatorSetIterator is returned from FilterStrategyRemovedFromOperatorSet and is used to iterate over the raw logs and unpacked data for StrategyRemovedFromOperatorSet events raised by the ContractAllocationManager contract.
type ContractAllocationManagerStrategyRemovedFromOperatorSetIterator struct {
	Event *ContractAllocationManagerStrategyRemovedFromOperatorSet // Event containing the contract specifics and raw log

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
func (it *ContractAllocationManagerStrategyRemovedFromOperatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAllocationManagerStrategyRemovedFromOperatorSet)
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
		it.Event = new(ContractAllocationManagerStrategyRemovedFromOperatorSet)
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
func (it *ContractAllocationManagerStrategyRemovedFromOperatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAllocationManagerStrategyRemovedFromOperatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAllocationManagerStrategyRemovedFromOperatorSet represents a StrategyRemovedFromOperatorSet event raised by the ContractAllocationManager contract.
type ContractAllocationManagerStrategyRemovedFromOperatorSet struct {
	OperatorSet OperatorSet
	Strategy    common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStrategyRemovedFromOperatorSet is a free log retrieval operation binding the contract event 0x7b4b073d80dcac55a11177d8459ad9f664ceeb91f71f27167bb14f8152a7eeee.
//
// Solidity: event StrategyRemovedFromOperatorSet((address,uint32) operatorSet, address strategy)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) FilterStrategyRemovedFromOperatorSet(opts *bind.FilterOpts) (*ContractAllocationManagerStrategyRemovedFromOperatorSetIterator, error) {

	logs, sub, err := _ContractAllocationManager.contract.FilterLogs(opts, "StrategyRemovedFromOperatorSet")
	if err != nil {
		return nil, err
	}
	return &ContractAllocationManagerStrategyRemovedFromOperatorSetIterator{contract: _ContractAllocationManager.contract, event: "StrategyRemovedFromOperatorSet", logs: logs, sub: sub}, nil
}

// WatchStrategyRemovedFromOperatorSet is a free log subscription operation binding the contract event 0x7b4b073d80dcac55a11177d8459ad9f664ceeb91f71f27167bb14f8152a7eeee.
//
// Solidity: event StrategyRemovedFromOperatorSet((address,uint32) operatorSet, address strategy)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) WatchStrategyRemovedFromOperatorSet(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerStrategyRemovedFromOperatorSet) (event.Subscription, error) {

	logs, sub, err := _ContractAllocationManager.contract.WatchLogs(opts, "StrategyRemovedFromOperatorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAllocationManagerStrategyRemovedFromOperatorSet)
				if err := _ContractAllocationManager.contract.UnpackLog(event, "StrategyRemovedFromOperatorSet", log); err != nil {
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

// ParseStrategyRemovedFromOperatorSet is a log parse operation binding the contract event 0x7b4b073d80dcac55a11177d8459ad9f664ceeb91f71f27167bb14f8152a7eeee.
//
// Solidity: event StrategyRemovedFromOperatorSet((address,uint32) operatorSet, address strategy)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) ParseStrategyRemovedFromOperatorSet(log types.Log) (*ContractAllocationManagerStrategyRemovedFromOperatorSet, error) {
	event := new(ContractAllocationManagerStrategyRemovedFromOperatorSet)
	if err := _ContractAllocationManager.contract.UnpackLog(event, "StrategyRemovedFromOperatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAllocationManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractAllocationManager contract.
type ContractAllocationManagerUnpausedIterator struct {
	Event *ContractAllocationManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *ContractAllocationManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAllocationManagerUnpaused)
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
		it.Event = new(ContractAllocationManagerUnpaused)
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
func (it *ContractAllocationManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAllocationManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAllocationManagerUnpaused represents a Unpaused event raised by the ContractAllocationManager contract.
type ContractAllocationManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractAllocationManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractAllocationManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractAllocationManagerUnpausedIterator{contract: _ContractAllocationManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractAllocationManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAllocationManagerUnpaused)
				if err := _ContractAllocationManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ContractAllocationManager *ContractAllocationManagerFilterer) ParseUnpaused(log types.Log) (*ContractAllocationManagerUnpaused, error) {
	event := new(ContractAllocationManagerUnpaused)
	if err := _ContractAllocationManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
