// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractAvsGovernance

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

// IAvsGovernanceInitializationParams is an auto generated low-level Go binding around an user-defined struct.
type IAvsGovernanceInitializationParams struct {
	AvsGovernanceMultisigOwner common.Address
	OperationsMultisig         common.Address
	CommunityMultisig          common.Address
	OthenticRegistry           common.Address
	MessageHandler             common.Address
	AvsTreasury                common.Address
	AvsDirectoryContract       common.Address
	AllowlistSigner            common.Address
	AvsName                    string
	BlsAuthSingleton           common.Address
	RedistributionManager      common.Address
}

// IAvsGovernanceSlashingConfig is an auto generated low-level Go binding around an user-defined struct.
type IAvsGovernanceSlashingConfig struct {
	Activated              bool
	EjectOperator          bool
	StakeSlashedPercentage *big.Int
}

// IAvsGovernanceStakingContractDetails is an auto generated low-level Go binding around an user-defined struct.
type IAvsGovernanceStakingContractDetails struct {
	StakingContract        common.Address
	MinStake               *big.Int
	MinSlashableStake      *big.Int
	SharedSecurityProvider uint8
}

// IAvsGovernanceStakingContractInfo is an auto generated low-level Go binding around an user-defined struct.
type IAvsGovernanceStakingContractInfo struct {
	StakingContract        common.Address
	SharedSecurityProvider uint8
}

// IAvsGovernanceVetoSlashRequest is an auto generated low-level Go binding around an user-defined struct.
type IAvsGovernanceVetoSlashRequest struct {
	Slasher    common.Address
	SlashIndex *big.Int
}

// IAvsGovernanceVotingPowerMultiplier is an auto generated low-level Go binding around an user-defined struct.
type IAvsGovernanceVotingPowerMultiplier struct {
	StakingContract        common.Address
	Multiplier             *big.Int
	SlashableStakeWeight   *big.Int
	SharedSecurityProvider uint8
}

// IRedistributionManagerSlashDetails is an auto generated low-level Go binding around an user-defined struct.
type IRedistributionManagerSlashDetails struct {
	Operator          common.Address
	EigenStrategies   []common.Address
	SharesSlashed     []*big.Int
	SlashingCondition uint8
}

// ISlashingConfigSlashingStakingContractInfo is an auto generated low-level Go binding around an user-defined struct.
type ISlashingConfigSlashingStakingContractInfo struct {
	StakingContract        common.Address
	SharedSecurityProvider uint8
	WadsToSlash            *big.Int
}

// ContractAvsGovernanceMetaData contains all meta data concerning the ContractAvsGovernance contract.
var ContractAvsGovernanceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_extensionImplementation\",\"type\":\"address\"},{\"internalType\":\"contractIOthenticRegistry\",\"name\":\"_othenticRegistry\",\"type\":\"address\"},{\"internalType\":\"contractIAllocationManager\",\"name\":\"_allocationManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccessControlInvalidMultiplierSyncer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllowlistDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllowlistEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyAvsName\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptySharedSecurityProvidersList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FlowIsCurrentlyPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FlowIsCurrentlyUnpaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMultiplier\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMultiplierNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRewardsReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"arrayIndex\",\"type\":\"uint256\"}],\"name\":\"InvalidSharedSecurityProviderList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSlashingRate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStakingContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidVetoSlashRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ModificationDelayNotPassed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_roles\",\"type\":\"bytes\"}],\"name\":\"NotAuthorizedRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughVotingPower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorNotRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PauseFlowIsAlreadyPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RedistributionManagerDeploymentFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RedistributionManagerNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SlashingConfigNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StakingContractsNotInAscendingOrder\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnpausingFlowIsAlreadyUnpaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedChainId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VetoSlashIndexOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"EjectOperatorFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"_pausableFlow\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_pauser\",\"type\":\"address\"}],\"name\":\"FlowPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"_pausableFlowFlag\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_unpauser\",\"type\":\"address\"}],\"name\":\"FlowUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"InvalidStakingContractsForSlashing\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"stakingContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minSlashableStake\",\"type\":\"uint256\"}],\"name\":\"MinSlashableStakePerStakingContractSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"stakingContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minStake\",\"type\":\"uint256\"}],\"name\":\"MinStakePerStakingContractSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashId\",\"type\":\"uint256\"}],\"name\":\"OperatorSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakingContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"PriceFeedSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"QueuedRewardsReceiverModification\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"RedistributionFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"avsGovernanceLogic\",\"type\":\"address\"}],\"name\":\"SetAvsGovernanceLogic\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"avsName\",\"type\":\"string\"}],\"name\":\"SetAvsName\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isAllowlisted\",\"type\":\"bool\"}],\"name\":\"SetIsAllowlisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_isEnabled\",\"type\":\"bool\"}],\"name\":\"SetP2pAuthenticationEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"SetRewardsReceiver\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"stakingContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakingContractMultiplier\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashableStakeWeight\",\"type\":\"uint256\"}],\"name\":\"SetStakingContractMultiplier\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SetToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"contractIStrategy[]\",\"name\":\"eigenStrategies\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"sharesSlashed\",\"type\":\"uint256[]\"},{\"internalType\":\"enumISlashingConfig.SlashingCondition\",\"name\":\"slashingCondition\",\"type\":\"uint8\"}],\"indexed\":false,\"internalType\":\"structIRedistributionManager.SlashDetails\",\"name\":\"slashDetails\",\"type\":\"tuple\"}],\"name\":\"SlashedFundsRedistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SlashingConfigNotActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SlashingConfigNotFound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"SlashingFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"SymbioticSlashingBypassedZeroAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashAmount\",\"type\":\"uint256\"}],\"name\":\"SymbioticSlashingExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"revertData\",\"type\":\"bytes\"}],\"name\":\"SymbioticSlashingReverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"SymbioticSlashingSkipped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"slasher\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"slashIndex\",\"type\":\"uint256\"}],\"name\":\"VetoSlashAlreadyCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"slasher\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"slashIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashedAmount\",\"type\":\"uint256\"}],\"name\":\"VetoSlashExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"slasher\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"slashIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"revertData\",\"type\":\"bytes\"}],\"name\":\"VetoSlashExecutionFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"slasher\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"slashIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"VetoSlashRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"stakingContracts\",\"type\":\"address[]\"}],\"name\":\"setNewSupportedStakingContracts\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXTENSION_IMPLEMENTATION\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"stakingContract\",\"type\":\"address\"},{\"internalType\":\"enumIAvsGovernance.SharedSecurityProvider\",\"name\":\"sharedSecurityProvider\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"wadsToSlash\",\"type\":\"uint256\"}],\"internalType\":\"structISlashingConfig.SlashingStakingContractInfo[]\",\"name\":\"_slashingStakingContractInfos\",\"type\":\"tuple[]\"}],\"name\":\"applyCustomSlashing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"avsGovernanceLogic\",\"outputs\":[{\"internalType\":\"contractIAvsGovernanceLogic\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"avsName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"avsTreasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"completeRewardsReceiverModification\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_to\",\"type\":\"uint256\"}],\"name\":\"executeVetoSlashRequests\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIsAllowlisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"getIsOperatorEjected\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"getOperatorRestakedStrategies\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPendingVetoSlashCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getPendingVetoSlashRequest\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"slasher\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slashIndex\",\"type\":\"uint256\"}],\"internalType\":\"structIAvsGovernance.VetoSlashRequest\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPendingVetoSlashRequests\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"slasher\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slashIndex\",\"type\":\"uint256\"}],\"internalType\":\"structIAvsGovernance.VetoSlashRequest[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRestakeableStrategies\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRestakeableVaults\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"getRewardsReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSlashableStrategies\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumISlashingConfig.SlashingCondition\",\"name\":\"_condition\",\"type\":\"uint8\"}],\"name\":\"getSlashingConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"activated\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"ejectOperator\",\"type\":\"bool\"},{\"internalType\":\"uint24\",\"name\":\"stakeSlashedPercentage\",\"type\":\"uint24\"}],\"internalType\":\"structIAvsGovernance.SlashingConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakingContractDetailsAndMultipliers\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"stakingContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSlashableStake\",\"type\":\"uint256\"},{\"internalType\":\"enumIAvsGovernance.SharedSecurityProvider\",\"name\":\"sharedSecurityProvider\",\"type\":\"uint8\"}],\"internalType\":\"structIAvsGovernance.StakingContractDetails[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"stakingContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"multiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slashableStakeWeight\",\"type\":\"uint256\"},{\"internalType\":\"enumIAvsGovernance.SharedSecurityProvider\",\"name\":\"sharedSecurityProvider\",\"type\":\"uint8\"}],\"internalType\":\"structIAvsGovernance.VotingPowerMultiplier[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"avsGovernanceMultisigOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operationsMultisig\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"communityMultisig\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"othenticRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"messageHandler\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"avsTreasury\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"avsDirectoryContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowlistSigner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"avsName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"blsAuthSingleton\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"redistributionManager\",\"type\":\"address\"}],\"internalType\":\"structIAvsGovernance.InitializationParams\",\"name\":\"_initializationParams\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_pausableFlow\",\"type\":\"bytes4\"}],\"name\":\"isFlowPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_isPaused\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isOperatorRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxEffectiveBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_metadataURI\",\"type\":\"string\"}],\"name\":\"migrateAvsToAllocationManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1AvsFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_avsGovernanceMultisigOwner\",\"type\":\"address\"}],\"name\":\"migrateRedistributionManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingContract\",\"type\":\"address\"}],\"name\":\"minSlashableStakePerStakingContract\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingContract\",\"type\":\"address\"}],\"name\":\"minStakePerStakingContract\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minVotingPower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingContract\",\"type\":\"address\"}],\"name\":\"multiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numOfActiveOperators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"p2pAuthenticationEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_pausableFlow\",\"type\":\"bytes4\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newRewardsReceiver\",\"type\":\"address\"}],\"name\":\"queueRewardsReceiverModification\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redistributionManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"}],\"name\":\"registerAvsToEigenLayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registerAvsToSymbiotic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAvsGovernanceLogic\",\"name\":\"_avsGovernanceLogic\",\"type\":\"address\"}],\"name\":\"setAvsGovernanceLogic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isAllowlisted\",\"type\":\"bool\"}],\"name\":\"setIsAllowlisted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_p2pAuthenticationEnabled\",\"type\":\"bool\"}],\"name\":\"setP2pAuthenticationEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"stakingContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"multiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slashableStakeWeight\",\"type\":\"uint256\"},{\"internalType\":\"enumIAvsGovernance.SharedSecurityProvider\",\"name\":\"sharedSecurityProvider\",\"type\":\"uint8\"}],\"internalType\":\"structIAvsGovernance.VotingPowerMultiplier\",\"name\":\"_votingPowerMultiplier\",\"type\":\"tuple\"}],\"name\":\"setStakingContractMultiplier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"stakingContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"multiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slashableStakeWeight\",\"type\":\"uint256\"},{\"internalType\":\"enumIAvsGovernance.SharedSecurityProvider\",\"name\":\"sharedSecurityProvider\",\"type\":\"uint8\"}],\"internalType\":\"structIAvsGovernance.VotingPowerMultiplier[]\",\"name\":\"_votingPowerMultipliers\",\"type\":\"tuple[]\"}],\"name\":\"setStakingContractMultiplierBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feed\",\"type\":\"address\"}],\"name\":\"setStakingContractPriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"stakingContract\",\"type\":\"address\"},{\"internalType\":\"enumIAvsGovernance.SharedSecurityProvider\",\"name\":\"sharedSecurityProvider\",\"type\":\"uint8\"}],\"internalType\":\"structIAvsGovernance.StakingContractInfo[]\",\"name\":\"_stakingContractsDetails\",\"type\":\"tuple[]\"}],\"name\":\"setSupportedStakingContracts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIVetoSlasher\",\"name\":\"_vetoSlasher\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_resolver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_hints\",\"type\":\"bytes\"}],\"name\":\"setSymbioticResolver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"enumISlashingConfig.SlashingCondition\",\"name\":\"_slashingCondition\",\"type\":\"uint8\"}],\"name\":\"slashOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingContract\",\"type\":\"address\"}],\"name\":\"slashableStakeWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingContract\",\"type\":\"address\"}],\"name\":\"stakingContractToFeed\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingContracts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_pausableFlow\",\"type\":\"bytes4\"}],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"votingPower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_stakingContracts\",\"type\":\"address[]\"}],\"name\":\"votingPowerPerStakingContracts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ContractAvsGovernanceABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractAvsGovernanceMetaData.ABI instead.
var ContractAvsGovernanceABI = ContractAvsGovernanceMetaData.ABI

// ContractAvsGovernanceMethods is an auto generated interface around an Ethereum contract.
type ContractAvsGovernanceMethods interface {
	ContractAvsGovernanceCalls
	ContractAvsGovernanceTransacts
	ContractAvsGovernanceFilters
}

// ContractAvsGovernanceCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractAvsGovernanceCalls interface {
	DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error)

	EXTENSIONIMPLEMENTATION(opts *bind.CallOpts) (common.Address, error)

	AvsGovernanceLogic(opts *bind.CallOpts) (common.Address, error)

	AvsName(opts *bind.CallOpts) (string, error)

	AvsTreasury(opts *bind.CallOpts) (common.Address, error)

	GetIsAllowlisted(opts *bind.CallOpts) (bool, error)

	GetIsOperatorEjected(opts *bind.CallOpts, _operator common.Address) (bool, error)

	GetOperatorRestakedStrategies(opts *bind.CallOpts, _operator common.Address) ([]common.Address, error)

	GetPendingVetoSlashCount(opts *bind.CallOpts) (*big.Int, error)

	GetPendingVetoSlashRequest(opts *bind.CallOpts, _index *big.Int) (IAvsGovernanceVetoSlashRequest, error)

	GetPendingVetoSlashRequests(opts *bind.CallOpts) ([]IAvsGovernanceVetoSlashRequest, error)

	GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error)

	GetRestakeableVaults(opts *bind.CallOpts) ([]common.Address, error)

	GetRewardsReceiver(opts *bind.CallOpts, _operator common.Address) (common.Address, error)

	GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error)

	GetSlashableStrategies(opts *bind.CallOpts) ([]common.Address, error)

	GetSlashingConfig(opts *bind.CallOpts, _condition uint8) (IAvsGovernanceSlashingConfig, error)

	GetStakingContractDetailsAndMultipliers(opts *bind.CallOpts) ([]IAvsGovernanceStakingContractDetails, []IAvsGovernanceVotingPowerMultiplier, error)

	HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error)

	IsFlowPaused(opts *bind.CallOpts, _pausableFlow [4]byte) (bool, error)

	IsOperatorRegistered(opts *bind.CallOpts, operator common.Address) (bool, error)

	MaxEffectiveBalance(opts *bind.CallOpts) (*big.Int, error)

	MinSlashableStakePerStakingContract(opts *bind.CallOpts, _stakingContract common.Address) (*big.Int, error)

	MinStakePerStakingContract(opts *bind.CallOpts, _stakingContract common.Address) (*big.Int, error)

	MinVotingPower(opts *bind.CallOpts) (*big.Int, error)

	Multiplier(opts *bind.CallOpts, _stakingContract common.Address) (*big.Int, error)

	NumOfActiveOperators(opts *bind.CallOpts) (*big.Int, error)

	P2pAuthenticationEnabled(opts *bind.CallOpts) (bool, error)

	RedistributionManager(opts *bind.CallOpts) (common.Address, error)

	SlashableStakeWeight(opts *bind.CallOpts, _stakingContract common.Address) (*big.Int, error)

	StakingContractToFeed(opts *bind.CallOpts, _stakingContract common.Address) (common.Address, error)

	StakingContracts(opts *bind.CallOpts) ([]common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	VotingPower(opts *bind.CallOpts, _operator common.Address) (*big.Int, error)

	VotingPowerPerStakingContracts(opts *bind.CallOpts, _operator common.Address, _stakingContracts []common.Address) (*big.Int, error)
}

// ContractAvsGovernanceTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractAvsGovernanceTransacts interface {
	ApplyCustomSlashing(opts *bind.TransactOpts, _operator common.Address, _slashingStakingContractInfos []ISlashingConfigSlashingStakingContractInfo) (*types.Transaction, error)

	CompleteRewardsReceiverModification(opts *bind.TransactOpts) (*types.Transaction, error)

	ExecuteVetoSlashRequests(opts *bind.TransactOpts, _from *big.Int, _to *big.Int) (*types.Transaction, error)

	GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts, _initializationParams IAvsGovernanceInitializationParams) (*types.Transaction, error)

	MigrateAvsToAllocationManager(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error)

	MigrateRedistributionManager(opts *bind.TransactOpts, _l1AvsFactory common.Address, _avsGovernanceMultisigOwner common.Address) (*types.Transaction, error)

	Migration(opts *bind.TransactOpts) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts, _pausableFlow [4]byte) (*types.Transaction, error)

	QueueRewardsReceiverModification(opts *bind.TransactOpts, _newRewardsReceiver common.Address) (*types.Transaction, error)

	RegisterAvsToEigenLayer(opts *bind.TransactOpts, metadataURI string) (*types.Transaction, error)

	RegisterAvsToSymbiotic(opts *bind.TransactOpts) (*types.Transaction, error)

	RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error)

	RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	SetAvsGovernanceLogic(opts *bind.TransactOpts, _avsGovernanceLogic common.Address) (*types.Transaction, error)

	SetIsAllowlisted(opts *bind.TransactOpts, _isAllowlisted bool) (*types.Transaction, error)

	SetP2pAuthenticationEnabled(opts *bind.TransactOpts, _p2pAuthenticationEnabled bool) (*types.Transaction, error)

	SetStakingContractMultiplier(opts *bind.TransactOpts, _votingPowerMultiplier IAvsGovernanceVotingPowerMultiplier) (*types.Transaction, error)

	SetStakingContractMultiplierBatch(opts *bind.TransactOpts, _votingPowerMultipliers []IAvsGovernanceVotingPowerMultiplier) (*types.Transaction, error)

	SetStakingContractPriceFeed(opts *bind.TransactOpts, _stakingContract common.Address, _feed common.Address) (*types.Transaction, error)

	SetSupportedStakingContracts(opts *bind.TransactOpts, _stakingContractsDetails []IAvsGovernanceStakingContractInfo) (*types.Transaction, error)

	SetSymbioticResolver(opts *bind.TransactOpts, _vetoSlasher common.Address, _resolver common.Address, _hints []byte) (*types.Transaction, error)

	SlashOperator(opts *bind.TransactOpts, _operator common.Address, _slashingCondition uint8) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts, _pausableFlow [4]byte) (*types.Transaction, error)
}

// ContractAvsGovernanceFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractAvsGovernanceFilters interface {
	FilterEjectOperatorFailed(opts *bind.FilterOpts) (*ContractAvsGovernanceEjectOperatorFailedIterator, error)
	WatchEjectOperatorFailed(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceEjectOperatorFailed) (event.Subscription, error)
	ParseEjectOperatorFailed(log types.Log) (*ContractAvsGovernanceEjectOperatorFailed, error)

	FilterFlowPaused(opts *bind.FilterOpts) (*ContractAvsGovernanceFlowPausedIterator, error)
	WatchFlowPaused(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceFlowPaused) (event.Subscription, error)
	ParseFlowPaused(log types.Log) (*ContractAvsGovernanceFlowPaused, error)

	FilterFlowUnpaused(opts *bind.FilterOpts) (*ContractAvsGovernanceFlowUnpausedIterator, error)
	WatchFlowUnpaused(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceFlowUnpaused) (event.Subscription, error)
	ParseFlowUnpaused(log types.Log) (*ContractAvsGovernanceFlowUnpaused, error)

	FilterInitialized(opts *bind.FilterOpts) (*ContractAvsGovernanceInitializedIterator, error)
	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceInitialized) (event.Subscription, error)
	ParseInitialized(log types.Log) (*ContractAvsGovernanceInitialized, error)

	FilterInvalidStakingContractsForSlashing(opts *bind.FilterOpts) (*ContractAvsGovernanceInvalidStakingContractsForSlashingIterator, error)
	WatchInvalidStakingContractsForSlashing(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceInvalidStakingContractsForSlashing) (event.Subscription, error)
	ParseInvalidStakingContractsForSlashing(log types.Log) (*ContractAvsGovernanceInvalidStakingContractsForSlashing, error)

	FilterMinSlashableStakePerStakingContractSet(opts *bind.FilterOpts) (*ContractAvsGovernanceMinSlashableStakePerStakingContractSetIterator, error)
	WatchMinSlashableStakePerStakingContractSet(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceMinSlashableStakePerStakingContractSet) (event.Subscription, error)
	ParseMinSlashableStakePerStakingContractSet(log types.Log) (*ContractAvsGovernanceMinSlashableStakePerStakingContractSet, error)

	FilterMinStakePerStakingContractSet(opts *bind.FilterOpts) (*ContractAvsGovernanceMinStakePerStakingContractSetIterator, error)
	WatchMinStakePerStakingContractSet(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceMinStakePerStakingContractSet) (event.Subscription, error)
	ParseMinStakePerStakingContractSet(log types.Log) (*ContractAvsGovernanceMinStakePerStakingContractSet, error)

	FilterOperatorSlashed(opts *bind.FilterOpts) (*ContractAvsGovernanceOperatorSlashedIterator, error)
	WatchOperatorSlashed(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceOperatorSlashed) (event.Subscription, error)
	ParseOperatorSlashed(log types.Log) (*ContractAvsGovernanceOperatorSlashed, error)

	FilterPriceFeedSet(opts *bind.FilterOpts, stakingContract []common.Address) (*ContractAvsGovernancePriceFeedSetIterator, error)
	WatchPriceFeedSet(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernancePriceFeedSet, stakingContract []common.Address) (event.Subscription, error)
	ParsePriceFeedSet(log types.Log) (*ContractAvsGovernancePriceFeedSet, error)

	FilterQueuedRewardsReceiverModification(opts *bind.FilterOpts, operator []common.Address) (*ContractAvsGovernanceQueuedRewardsReceiverModificationIterator, error)
	WatchQueuedRewardsReceiverModification(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceQueuedRewardsReceiverModification, operator []common.Address) (event.Subscription, error)
	ParseQueuedRewardsReceiverModification(log types.Log) (*ContractAvsGovernanceQueuedRewardsReceiverModification, error)

	FilterRedistributionFailed(opts *bind.FilterOpts) (*ContractAvsGovernanceRedistributionFailedIterator, error)
	WatchRedistributionFailed(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceRedistributionFailed) (event.Subscription, error)
	ParseRedistributionFailed(log types.Log) (*ContractAvsGovernanceRedistributionFailed, error)

	FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ContractAvsGovernanceRoleAdminChangedIterator, error)
	WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error)
	ParseRoleAdminChanged(log types.Log) (*ContractAvsGovernanceRoleAdminChanged, error)

	FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ContractAvsGovernanceRoleGrantedIterator, error)
	WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error)
	ParseRoleGranted(log types.Log) (*ContractAvsGovernanceRoleGranted, error)

	FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ContractAvsGovernanceRoleRevokedIterator, error)
	WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error)
	ParseRoleRevoked(log types.Log) (*ContractAvsGovernanceRoleRevoked, error)

	FilterSetAvsGovernanceLogic(opts *bind.FilterOpts) (*ContractAvsGovernanceSetAvsGovernanceLogicIterator, error)
	WatchSetAvsGovernanceLogic(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetAvsGovernanceLogic) (event.Subscription, error)
	ParseSetAvsGovernanceLogic(log types.Log) (*ContractAvsGovernanceSetAvsGovernanceLogic, error)

	FilterSetAvsName(opts *bind.FilterOpts) (*ContractAvsGovernanceSetAvsNameIterator, error)
	WatchSetAvsName(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetAvsName) (event.Subscription, error)
	ParseSetAvsName(log types.Log) (*ContractAvsGovernanceSetAvsName, error)

	FilterSetIsAllowlisted(opts *bind.FilterOpts) (*ContractAvsGovernanceSetIsAllowlistedIterator, error)
	WatchSetIsAllowlisted(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetIsAllowlisted) (event.Subscription, error)
	ParseSetIsAllowlisted(log types.Log) (*ContractAvsGovernanceSetIsAllowlisted, error)

	FilterSetP2pAuthenticationEnabled(opts *bind.FilterOpts) (*ContractAvsGovernanceSetP2pAuthenticationEnabledIterator, error)
	WatchSetP2pAuthenticationEnabled(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetP2pAuthenticationEnabled) (event.Subscription, error)
	ParseSetP2pAuthenticationEnabled(log types.Log) (*ContractAvsGovernanceSetP2pAuthenticationEnabled, error)

	FilterSetRewardsReceiver(opts *bind.FilterOpts, operator []common.Address) (*ContractAvsGovernanceSetRewardsReceiverIterator, error)
	WatchSetRewardsReceiver(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetRewardsReceiver, operator []common.Address) (event.Subscription, error)
	ParseSetRewardsReceiver(log types.Log) (*ContractAvsGovernanceSetRewardsReceiver, error)

	FilterSetStakingContractMultiplier(opts *bind.FilterOpts) (*ContractAvsGovernanceSetStakingContractMultiplierIterator, error)
	WatchSetStakingContractMultiplier(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetStakingContractMultiplier) (event.Subscription, error)
	ParseSetStakingContractMultiplier(log types.Log) (*ContractAvsGovernanceSetStakingContractMultiplier, error)

	FilterSetToken(opts *bind.FilterOpts) (*ContractAvsGovernanceSetTokenIterator, error)
	WatchSetToken(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetToken) (event.Subscription, error)
	ParseSetToken(log types.Log) (*ContractAvsGovernanceSetToken, error)

	FilterSlashedFundsRedistributed(opts *bind.FilterOpts) (*ContractAvsGovernanceSlashedFundsRedistributedIterator, error)
	WatchSlashedFundsRedistributed(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSlashedFundsRedistributed) (event.Subscription, error)
	ParseSlashedFundsRedistributed(log types.Log) (*ContractAvsGovernanceSlashedFundsRedistributed, error)

	FilterSlashingConfigNotActivated(opts *bind.FilterOpts) (*ContractAvsGovernanceSlashingConfigNotActivatedIterator, error)
	WatchSlashingConfigNotActivated(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSlashingConfigNotActivated) (event.Subscription, error)
	ParseSlashingConfigNotActivated(log types.Log) (*ContractAvsGovernanceSlashingConfigNotActivated, error)

	FilterSlashingConfigNotFound(opts *bind.FilterOpts) (*ContractAvsGovernanceSlashingConfigNotFoundIterator, error)
	WatchSlashingConfigNotFound(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSlashingConfigNotFound) (event.Subscription, error)
	ParseSlashingConfigNotFound(log types.Log) (*ContractAvsGovernanceSlashingConfigNotFound, error)

	FilterSlashingFailed(opts *bind.FilterOpts) (*ContractAvsGovernanceSlashingFailedIterator, error)
	WatchSlashingFailed(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSlashingFailed) (event.Subscription, error)
	ParseSlashingFailed(log types.Log) (*ContractAvsGovernanceSlashingFailed, error)

	FilterSymbioticSlashingBypassedZeroAmount(opts *bind.FilterOpts, vault []common.Address, operator []common.Address) (*ContractAvsGovernanceSymbioticSlashingBypassedZeroAmountIterator, error)
	WatchSymbioticSlashingBypassedZeroAmount(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSymbioticSlashingBypassedZeroAmount, vault []common.Address, operator []common.Address) (event.Subscription, error)
	ParseSymbioticSlashingBypassedZeroAmount(log types.Log) (*ContractAvsGovernanceSymbioticSlashingBypassedZeroAmount, error)

	FilterSymbioticSlashingExecuted(opts *bind.FilterOpts, vault []common.Address, operator []common.Address) (*ContractAvsGovernanceSymbioticSlashingExecutedIterator, error)
	WatchSymbioticSlashingExecuted(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSymbioticSlashingExecuted, vault []common.Address, operator []common.Address) (event.Subscription, error)
	ParseSymbioticSlashingExecuted(log types.Log) (*ContractAvsGovernanceSymbioticSlashingExecuted, error)

	FilterSymbioticSlashingReverted(opts *bind.FilterOpts, vault []common.Address, operator []common.Address) (*ContractAvsGovernanceSymbioticSlashingRevertedIterator, error)
	WatchSymbioticSlashingReverted(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSymbioticSlashingReverted, vault []common.Address, operator []common.Address) (event.Subscription, error)
	ParseSymbioticSlashingReverted(log types.Log) (*ContractAvsGovernanceSymbioticSlashingReverted, error)

	FilterSymbioticSlashingSkipped(opts *bind.FilterOpts, vault []common.Address, operator []common.Address) (*ContractAvsGovernanceSymbioticSlashingSkippedIterator, error)
	WatchSymbioticSlashingSkipped(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSymbioticSlashingSkipped, vault []common.Address, operator []common.Address) (event.Subscription, error)
	ParseSymbioticSlashingSkipped(log types.Log) (*ContractAvsGovernanceSymbioticSlashingSkipped, error)

	FilterVetoSlashAlreadyCompleted(opts *bind.FilterOpts, slasher []common.Address, slashIndex []*big.Int) (*ContractAvsGovernanceVetoSlashAlreadyCompletedIterator, error)
	WatchVetoSlashAlreadyCompleted(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceVetoSlashAlreadyCompleted, slasher []common.Address, slashIndex []*big.Int) (event.Subscription, error)
	ParseVetoSlashAlreadyCompleted(log types.Log) (*ContractAvsGovernanceVetoSlashAlreadyCompleted, error)

	FilterVetoSlashExecuted(opts *bind.FilterOpts, slasher []common.Address, slashIndex []*big.Int) (*ContractAvsGovernanceVetoSlashExecutedIterator, error)
	WatchVetoSlashExecuted(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceVetoSlashExecuted, slasher []common.Address, slashIndex []*big.Int) (event.Subscription, error)
	ParseVetoSlashExecuted(log types.Log) (*ContractAvsGovernanceVetoSlashExecuted, error)

	FilterVetoSlashExecutionFailed(opts *bind.FilterOpts, slasher []common.Address, slashIndex []*big.Int) (*ContractAvsGovernanceVetoSlashExecutionFailedIterator, error)
	WatchVetoSlashExecutionFailed(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceVetoSlashExecutionFailed, slasher []common.Address, slashIndex []*big.Int) (event.Subscription, error)
	ParseVetoSlashExecutionFailed(log types.Log) (*ContractAvsGovernanceVetoSlashExecutionFailed, error)

	FilterVetoSlashRequested(opts *bind.FilterOpts, slasher []common.Address, slashIndex []*big.Int, operator []common.Address) (*ContractAvsGovernanceVetoSlashRequestedIterator, error)
	WatchVetoSlashRequested(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceVetoSlashRequested, slasher []common.Address, slashIndex []*big.Int, operator []common.Address) (event.Subscription, error)
	ParseVetoSlashRequested(log types.Log) (*ContractAvsGovernanceVetoSlashRequested, error)

	FilterSetNewSupportedStakingContracts(opts *bind.FilterOpts) (*ContractAvsGovernanceSetNewSupportedStakingContractsIterator, error)
	WatchSetNewSupportedStakingContracts(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetNewSupportedStakingContracts) (event.Subscription, error)
	ParseSetNewSupportedStakingContracts(log types.Log) (*ContractAvsGovernanceSetNewSupportedStakingContracts, error)
}

// ContractAvsGovernance is an auto generated Go binding around an Ethereum contract.
type ContractAvsGovernance struct {
	ContractAvsGovernanceCaller     // Read-only binding to the contract
	ContractAvsGovernanceTransactor // Write-only binding to the contract
	ContractAvsGovernanceFilterer   // Log filterer for contract events
}

// ContractAvsGovernance implements the ContractAvsGovernanceMethods interface.
var _ ContractAvsGovernanceMethods = (*ContractAvsGovernance)(nil)

// ContractAvsGovernanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractAvsGovernanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAvsGovernanceCaller implements the ContractAvsGovernanceCalls interface.
var _ ContractAvsGovernanceCalls = (*ContractAvsGovernanceCaller)(nil)

// ContractAvsGovernanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractAvsGovernanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAvsGovernanceTransactor implements the ContractAvsGovernanceTransacts interface.
var _ ContractAvsGovernanceTransacts = (*ContractAvsGovernanceTransactor)(nil)

// ContractAvsGovernanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractAvsGovernanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAvsGovernanceFilterer implements the ContractAvsGovernanceFilters interface.
var _ ContractAvsGovernanceFilters = (*ContractAvsGovernanceFilterer)(nil)

// ContractAvsGovernanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractAvsGovernanceSession struct {
	Contract     *ContractAvsGovernance // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ContractAvsGovernanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractAvsGovernanceCallerSession struct {
	Contract *ContractAvsGovernanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// ContractAvsGovernanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractAvsGovernanceTransactorSession struct {
	Contract     *ContractAvsGovernanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ContractAvsGovernanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractAvsGovernanceRaw struct {
	Contract *ContractAvsGovernance // Generic contract binding to access the raw methods on
}

// ContractAvsGovernanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractAvsGovernanceCallerRaw struct {
	Contract *ContractAvsGovernanceCaller // Generic read-only contract binding to access the raw methods on
}

// ContractAvsGovernanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractAvsGovernanceTransactorRaw struct {
	Contract *ContractAvsGovernanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractAvsGovernance creates a new instance of ContractAvsGovernance, bound to a specific deployed contract.
func NewContractAvsGovernance(address common.Address, backend bind.ContractBackend) (*ContractAvsGovernance, error) {
	contract, err := bindContractAvsGovernance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernance{ContractAvsGovernanceCaller: ContractAvsGovernanceCaller{contract: contract}, ContractAvsGovernanceTransactor: ContractAvsGovernanceTransactor{contract: contract}, ContractAvsGovernanceFilterer: ContractAvsGovernanceFilterer{contract: contract}}, nil
}

// NewContractAvsGovernanceCaller creates a new read-only instance of ContractAvsGovernance, bound to a specific deployed contract.
func NewContractAvsGovernanceCaller(address common.Address, caller bind.ContractCaller) (*ContractAvsGovernanceCaller, error) {
	contract, err := bindContractAvsGovernance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceCaller{contract: contract}, nil
}

// NewContractAvsGovernanceTransactor creates a new write-only instance of ContractAvsGovernance, bound to a specific deployed contract.
func NewContractAvsGovernanceTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractAvsGovernanceTransactor, error) {
	contract, err := bindContractAvsGovernance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceTransactor{contract: contract}, nil
}

// NewContractAvsGovernanceFilterer creates a new log filterer instance of ContractAvsGovernance, bound to a specific deployed contract.
func NewContractAvsGovernanceFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractAvsGovernanceFilterer, error) {
	contract, err := bindContractAvsGovernance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceFilterer{contract: contract}, nil
}

// bindContractAvsGovernance binds a generic wrapper to an already deployed contract.
func bindContractAvsGovernance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractAvsGovernanceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractAvsGovernance *ContractAvsGovernanceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractAvsGovernance.Contract.ContractAvsGovernanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractAvsGovernance *ContractAvsGovernanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.ContractAvsGovernanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractAvsGovernance *ContractAvsGovernanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.ContractAvsGovernanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractAvsGovernance *ContractAvsGovernanceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractAvsGovernance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ContractAvsGovernance.Contract.DEFAULTADMINROLE(&_ContractAvsGovernance.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ContractAvsGovernance.Contract.DEFAULTADMINROLE(&_ContractAvsGovernance.CallOpts)
}

// EXTENSIONIMPLEMENTATION is a free data retrieval call binding the contract method 0x226def04.
//
// Solidity: function EXTENSION_IMPLEMENTATION() view returns(address)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) EXTENSIONIMPLEMENTATION(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "EXTENSION_IMPLEMENTATION")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EXTENSIONIMPLEMENTATION is a free data retrieval call binding the contract method 0x226def04.
//
// Solidity: function EXTENSION_IMPLEMENTATION() view returns(address)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) EXTENSIONIMPLEMENTATION() (common.Address, error) {
	return _ContractAvsGovernance.Contract.EXTENSIONIMPLEMENTATION(&_ContractAvsGovernance.CallOpts)
}

// EXTENSIONIMPLEMENTATION is a free data retrieval call binding the contract method 0x226def04.
//
// Solidity: function EXTENSION_IMPLEMENTATION() view returns(address)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) EXTENSIONIMPLEMENTATION() (common.Address, error) {
	return _ContractAvsGovernance.Contract.EXTENSIONIMPLEMENTATION(&_ContractAvsGovernance.CallOpts)
}

// AvsGovernanceLogic is a free data retrieval call binding the contract method 0x9e91c39e.
//
// Solidity: function avsGovernanceLogic() view returns(address)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) AvsGovernanceLogic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "avsGovernanceLogic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsGovernanceLogic is a free data retrieval call binding the contract method 0x9e91c39e.
//
// Solidity: function avsGovernanceLogic() view returns(address)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) AvsGovernanceLogic() (common.Address, error) {
	return _ContractAvsGovernance.Contract.AvsGovernanceLogic(&_ContractAvsGovernance.CallOpts)
}

// AvsGovernanceLogic is a free data retrieval call binding the contract method 0x9e91c39e.
//
// Solidity: function avsGovernanceLogic() view returns(address)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) AvsGovernanceLogic() (common.Address, error) {
	return _ContractAvsGovernance.Contract.AvsGovernanceLogic(&_ContractAvsGovernance.CallOpts)
}

// AvsName is a free data retrieval call binding the contract method 0x41b92a29.
//
// Solidity: function avsName() view returns(string)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) AvsName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "avsName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AvsName is a free data retrieval call binding the contract method 0x41b92a29.
//
// Solidity: function avsName() view returns(string)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) AvsName() (string, error) {
	return _ContractAvsGovernance.Contract.AvsName(&_ContractAvsGovernance.CallOpts)
}

// AvsName is a free data retrieval call binding the contract method 0x41b92a29.
//
// Solidity: function avsName() view returns(string)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) AvsName() (string, error) {
	return _ContractAvsGovernance.Contract.AvsName(&_ContractAvsGovernance.CallOpts)
}

// AvsTreasury is a free data retrieval call binding the contract method 0x1246193e.
//
// Solidity: function avsTreasury() view returns(address)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) AvsTreasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "avsTreasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsTreasury is a free data retrieval call binding the contract method 0x1246193e.
//
// Solidity: function avsTreasury() view returns(address)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) AvsTreasury() (common.Address, error) {
	return _ContractAvsGovernance.Contract.AvsTreasury(&_ContractAvsGovernance.CallOpts)
}

// AvsTreasury is a free data retrieval call binding the contract method 0x1246193e.
//
// Solidity: function avsTreasury() view returns(address)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) AvsTreasury() (common.Address, error) {
	return _ContractAvsGovernance.Contract.AvsTreasury(&_ContractAvsGovernance.CallOpts)
}

// GetIsAllowlisted is a free data retrieval call binding the contract method 0xb525fa88.
//
// Solidity: function getIsAllowlisted() view returns(bool)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) GetIsAllowlisted(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "getIsAllowlisted")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetIsAllowlisted is a free data retrieval call binding the contract method 0xb525fa88.
//
// Solidity: function getIsAllowlisted() view returns(bool)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) GetIsAllowlisted() (bool, error) {
	return _ContractAvsGovernance.Contract.GetIsAllowlisted(&_ContractAvsGovernance.CallOpts)
}

// GetIsAllowlisted is a free data retrieval call binding the contract method 0xb525fa88.
//
// Solidity: function getIsAllowlisted() view returns(bool)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) GetIsAllowlisted() (bool, error) {
	return _ContractAvsGovernance.Contract.GetIsAllowlisted(&_ContractAvsGovernance.CallOpts)
}

// GetIsOperatorEjected is a free data retrieval call binding the contract method 0x37d80e7d.
//
// Solidity: function getIsOperatorEjected(address _operator) view returns(bool)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) GetIsOperatorEjected(opts *bind.CallOpts, _operator common.Address) (bool, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "getIsOperatorEjected", _operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetIsOperatorEjected is a free data retrieval call binding the contract method 0x37d80e7d.
//
// Solidity: function getIsOperatorEjected(address _operator) view returns(bool)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) GetIsOperatorEjected(_operator common.Address) (bool, error) {
	return _ContractAvsGovernance.Contract.GetIsOperatorEjected(&_ContractAvsGovernance.CallOpts, _operator)
}

// GetIsOperatorEjected is a free data retrieval call binding the contract method 0x37d80e7d.
//
// Solidity: function getIsOperatorEjected(address _operator) view returns(bool)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) GetIsOperatorEjected(_operator common.Address) (bool, error) {
	return _ContractAvsGovernance.Contract.GetIsOperatorEjected(&_ContractAvsGovernance.CallOpts, _operator)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address _operator) view returns(address[])
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) GetOperatorRestakedStrategies(opts *bind.CallOpts, _operator common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "getOperatorRestakedStrategies", _operator)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address _operator) view returns(address[])
func (_ContractAvsGovernance *ContractAvsGovernanceSession) GetOperatorRestakedStrategies(_operator common.Address) ([]common.Address, error) {
	return _ContractAvsGovernance.Contract.GetOperatorRestakedStrategies(&_ContractAvsGovernance.CallOpts, _operator)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address _operator) view returns(address[])
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) GetOperatorRestakedStrategies(_operator common.Address) ([]common.Address, error) {
	return _ContractAvsGovernance.Contract.GetOperatorRestakedStrategies(&_ContractAvsGovernance.CallOpts, _operator)
}

// GetPendingVetoSlashCount is a free data retrieval call binding the contract method 0xa889f180.
//
// Solidity: function getPendingVetoSlashCount() view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) GetPendingVetoSlashCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "getPendingVetoSlashCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingVetoSlashCount is a free data retrieval call binding the contract method 0xa889f180.
//
// Solidity: function getPendingVetoSlashCount() view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) GetPendingVetoSlashCount() (*big.Int, error) {
	return _ContractAvsGovernance.Contract.GetPendingVetoSlashCount(&_ContractAvsGovernance.CallOpts)
}

// GetPendingVetoSlashCount is a free data retrieval call binding the contract method 0xa889f180.
//
// Solidity: function getPendingVetoSlashCount() view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) GetPendingVetoSlashCount() (*big.Int, error) {
	return _ContractAvsGovernance.Contract.GetPendingVetoSlashCount(&_ContractAvsGovernance.CallOpts)
}

// GetPendingVetoSlashRequest is a free data retrieval call binding the contract method 0xb9b00a34.
//
// Solidity: function getPendingVetoSlashRequest(uint256 _index) view returns((address,uint256))
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) GetPendingVetoSlashRequest(opts *bind.CallOpts, _index *big.Int) (IAvsGovernanceVetoSlashRequest, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "getPendingVetoSlashRequest", _index)

	if err != nil {
		return *new(IAvsGovernanceVetoSlashRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(IAvsGovernanceVetoSlashRequest)).(*IAvsGovernanceVetoSlashRequest)

	return out0, err

}

// GetPendingVetoSlashRequest is a free data retrieval call binding the contract method 0xb9b00a34.
//
// Solidity: function getPendingVetoSlashRequest(uint256 _index) view returns((address,uint256))
func (_ContractAvsGovernance *ContractAvsGovernanceSession) GetPendingVetoSlashRequest(_index *big.Int) (IAvsGovernanceVetoSlashRequest, error) {
	return _ContractAvsGovernance.Contract.GetPendingVetoSlashRequest(&_ContractAvsGovernance.CallOpts, _index)
}

// GetPendingVetoSlashRequest is a free data retrieval call binding the contract method 0xb9b00a34.
//
// Solidity: function getPendingVetoSlashRequest(uint256 _index) view returns((address,uint256))
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) GetPendingVetoSlashRequest(_index *big.Int) (IAvsGovernanceVetoSlashRequest, error) {
	return _ContractAvsGovernance.Contract.GetPendingVetoSlashRequest(&_ContractAvsGovernance.CallOpts, _index)
}

// GetPendingVetoSlashRequests is a free data retrieval call binding the contract method 0x7067c01f.
//
// Solidity: function getPendingVetoSlashRequests() view returns((address,uint256)[])
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) GetPendingVetoSlashRequests(opts *bind.CallOpts) ([]IAvsGovernanceVetoSlashRequest, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "getPendingVetoSlashRequests")

	if err != nil {
		return *new([]IAvsGovernanceVetoSlashRequest), err
	}

	out0 := *abi.ConvertType(out[0], new([]IAvsGovernanceVetoSlashRequest)).(*[]IAvsGovernanceVetoSlashRequest)

	return out0, err

}

// GetPendingVetoSlashRequests is a free data retrieval call binding the contract method 0x7067c01f.
//
// Solidity: function getPendingVetoSlashRequests() view returns((address,uint256)[])
func (_ContractAvsGovernance *ContractAvsGovernanceSession) GetPendingVetoSlashRequests() ([]IAvsGovernanceVetoSlashRequest, error) {
	return _ContractAvsGovernance.Contract.GetPendingVetoSlashRequests(&_ContractAvsGovernance.CallOpts)
}

// GetPendingVetoSlashRequests is a free data retrieval call binding the contract method 0x7067c01f.
//
// Solidity: function getPendingVetoSlashRequests() view returns((address,uint256)[])
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) GetPendingVetoSlashRequests() ([]IAvsGovernanceVetoSlashRequest, error) {
	return _ContractAvsGovernance.Contract.GetPendingVetoSlashRequests(&_ContractAvsGovernance.CallOpts)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "getRestakeableStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractAvsGovernance *ContractAvsGovernanceSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractAvsGovernance.Contract.GetRestakeableStrategies(&_ContractAvsGovernance.CallOpts)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractAvsGovernance.Contract.GetRestakeableStrategies(&_ContractAvsGovernance.CallOpts)
}

// GetRestakeableVaults is a free data retrieval call binding the contract method 0xdf098f1d.
//
// Solidity: function getRestakeableVaults() view returns(address[])
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) GetRestakeableVaults(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "getRestakeableVaults")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRestakeableVaults is a free data retrieval call binding the contract method 0xdf098f1d.
//
// Solidity: function getRestakeableVaults() view returns(address[])
func (_ContractAvsGovernance *ContractAvsGovernanceSession) GetRestakeableVaults() ([]common.Address, error) {
	return _ContractAvsGovernance.Contract.GetRestakeableVaults(&_ContractAvsGovernance.CallOpts)
}

// GetRestakeableVaults is a free data retrieval call binding the contract method 0xdf098f1d.
//
// Solidity: function getRestakeableVaults() view returns(address[])
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) GetRestakeableVaults() ([]common.Address, error) {
	return _ContractAvsGovernance.Contract.GetRestakeableVaults(&_ContractAvsGovernance.CallOpts)
}

// GetRewardsReceiver is a free data retrieval call binding the contract method 0x5e95cee2.
//
// Solidity: function getRewardsReceiver(address _operator) view returns(address)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) GetRewardsReceiver(opts *bind.CallOpts, _operator common.Address) (common.Address, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "getRewardsReceiver", _operator)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRewardsReceiver is a free data retrieval call binding the contract method 0x5e95cee2.
//
// Solidity: function getRewardsReceiver(address _operator) view returns(address)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) GetRewardsReceiver(_operator common.Address) (common.Address, error) {
	return _ContractAvsGovernance.Contract.GetRewardsReceiver(&_ContractAvsGovernance.CallOpts, _operator)
}

// GetRewardsReceiver is a free data retrieval call binding the contract method 0x5e95cee2.
//
// Solidity: function getRewardsReceiver(address _operator) view returns(address)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) GetRewardsReceiver(_operator common.Address) (common.Address, error) {
	return _ContractAvsGovernance.Contract.GetRewardsReceiver(&_ContractAvsGovernance.CallOpts, _operator)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ContractAvsGovernance.Contract.GetRoleAdmin(&_ContractAvsGovernance.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ContractAvsGovernance.Contract.GetRoleAdmin(&_ContractAvsGovernance.CallOpts, role)
}

// GetSlashableStrategies is a free data retrieval call binding the contract method 0x7b4a5917.
//
// Solidity: function getSlashableStrategies() view returns(address[])
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) GetSlashableStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "getSlashableStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSlashableStrategies is a free data retrieval call binding the contract method 0x7b4a5917.
//
// Solidity: function getSlashableStrategies() view returns(address[])
func (_ContractAvsGovernance *ContractAvsGovernanceSession) GetSlashableStrategies() ([]common.Address, error) {
	return _ContractAvsGovernance.Contract.GetSlashableStrategies(&_ContractAvsGovernance.CallOpts)
}

// GetSlashableStrategies is a free data retrieval call binding the contract method 0x7b4a5917.
//
// Solidity: function getSlashableStrategies() view returns(address[])
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) GetSlashableStrategies() ([]common.Address, error) {
	return _ContractAvsGovernance.Contract.GetSlashableStrategies(&_ContractAvsGovernance.CallOpts)
}

// GetSlashingConfig is a free data retrieval call binding the contract method 0x0a0c5af7.
//
// Solidity: function getSlashingConfig(uint8 _condition) view returns((bool,bool,uint24))
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) GetSlashingConfig(opts *bind.CallOpts, _condition uint8) (IAvsGovernanceSlashingConfig, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "getSlashingConfig", _condition)

	if err != nil {
		return *new(IAvsGovernanceSlashingConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IAvsGovernanceSlashingConfig)).(*IAvsGovernanceSlashingConfig)

	return out0, err

}

// GetSlashingConfig is a free data retrieval call binding the contract method 0x0a0c5af7.
//
// Solidity: function getSlashingConfig(uint8 _condition) view returns((bool,bool,uint24))
func (_ContractAvsGovernance *ContractAvsGovernanceSession) GetSlashingConfig(_condition uint8) (IAvsGovernanceSlashingConfig, error) {
	return _ContractAvsGovernance.Contract.GetSlashingConfig(&_ContractAvsGovernance.CallOpts, _condition)
}

// GetSlashingConfig is a free data retrieval call binding the contract method 0x0a0c5af7.
//
// Solidity: function getSlashingConfig(uint8 _condition) view returns((bool,bool,uint24))
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) GetSlashingConfig(_condition uint8) (IAvsGovernanceSlashingConfig, error) {
	return _ContractAvsGovernance.Contract.GetSlashingConfig(&_ContractAvsGovernance.CallOpts, _condition)
}

// GetStakingContractDetailsAndMultipliers is a free data retrieval call binding the contract method 0xeb243a1c.
//
// Solidity: function getStakingContractDetailsAndMultipliers() view returns((address,uint256,uint256,uint8)[], (address,uint256,uint256,uint8)[])
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) GetStakingContractDetailsAndMultipliers(opts *bind.CallOpts) ([]IAvsGovernanceStakingContractDetails, []IAvsGovernanceVotingPowerMultiplier, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "getStakingContractDetailsAndMultipliers")

	if err != nil {
		return *new([]IAvsGovernanceStakingContractDetails), *new([]IAvsGovernanceVotingPowerMultiplier), err
	}

	out0 := *abi.ConvertType(out[0], new([]IAvsGovernanceStakingContractDetails)).(*[]IAvsGovernanceStakingContractDetails)
	out1 := *abi.ConvertType(out[1], new([]IAvsGovernanceVotingPowerMultiplier)).(*[]IAvsGovernanceVotingPowerMultiplier)

	return out0, out1, err

}

// GetStakingContractDetailsAndMultipliers is a free data retrieval call binding the contract method 0xeb243a1c.
//
// Solidity: function getStakingContractDetailsAndMultipliers() view returns((address,uint256,uint256,uint8)[], (address,uint256,uint256,uint8)[])
func (_ContractAvsGovernance *ContractAvsGovernanceSession) GetStakingContractDetailsAndMultipliers() ([]IAvsGovernanceStakingContractDetails, []IAvsGovernanceVotingPowerMultiplier, error) {
	return _ContractAvsGovernance.Contract.GetStakingContractDetailsAndMultipliers(&_ContractAvsGovernance.CallOpts)
}

// GetStakingContractDetailsAndMultipliers is a free data retrieval call binding the contract method 0xeb243a1c.
//
// Solidity: function getStakingContractDetailsAndMultipliers() view returns((address,uint256,uint256,uint8)[], (address,uint256,uint256,uint8)[])
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) GetStakingContractDetailsAndMultipliers() ([]IAvsGovernanceStakingContractDetails, []IAvsGovernanceVotingPowerMultiplier, error) {
	return _ContractAvsGovernance.Contract.GetStakingContractDetailsAndMultipliers(&_ContractAvsGovernance.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ContractAvsGovernance.Contract.HasRole(&_ContractAvsGovernance.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ContractAvsGovernance.Contract.HasRole(&_ContractAvsGovernance.CallOpts, role, account)
}

// IsFlowPaused is a free data retrieval call binding the contract method 0xefd96978.
//
// Solidity: function isFlowPaused(bytes4 _pausableFlow) view returns(bool _isPaused)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) IsFlowPaused(opts *bind.CallOpts, _pausableFlow [4]byte) (bool, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "isFlowPaused", _pausableFlow)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFlowPaused is a free data retrieval call binding the contract method 0xefd96978.
//
// Solidity: function isFlowPaused(bytes4 _pausableFlow) view returns(bool _isPaused)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) IsFlowPaused(_pausableFlow [4]byte) (bool, error) {
	return _ContractAvsGovernance.Contract.IsFlowPaused(&_ContractAvsGovernance.CallOpts, _pausableFlow)
}

// IsFlowPaused is a free data retrieval call binding the contract method 0xefd96978.
//
// Solidity: function isFlowPaused(bytes4 _pausableFlow) view returns(bool _isPaused)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) IsFlowPaused(_pausableFlow [4]byte) (bool, error) {
	return _ContractAvsGovernance.Contract.IsFlowPaused(&_ContractAvsGovernance.CallOpts, _pausableFlow)
}

// IsOperatorRegistered is a free data retrieval call binding the contract method 0x6b1906f8.
//
// Solidity: function isOperatorRegistered(address operator) view returns(bool)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) IsOperatorRegistered(opts *bind.CallOpts, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "isOperatorRegistered", operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorRegistered is a free data retrieval call binding the contract method 0x6b1906f8.
//
// Solidity: function isOperatorRegistered(address operator) view returns(bool)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) IsOperatorRegistered(operator common.Address) (bool, error) {
	return _ContractAvsGovernance.Contract.IsOperatorRegistered(&_ContractAvsGovernance.CallOpts, operator)
}

// IsOperatorRegistered is a free data retrieval call binding the contract method 0x6b1906f8.
//
// Solidity: function isOperatorRegistered(address operator) view returns(bool)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) IsOperatorRegistered(operator common.Address) (bool, error) {
	return _ContractAvsGovernance.Contract.IsOperatorRegistered(&_ContractAvsGovernance.CallOpts, operator)
}

// MaxEffectiveBalance is a free data retrieval call binding the contract method 0xa88171ee.
//
// Solidity: function maxEffectiveBalance() view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) MaxEffectiveBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "maxEffectiveBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxEffectiveBalance is a free data retrieval call binding the contract method 0xa88171ee.
//
// Solidity: function maxEffectiveBalance() view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) MaxEffectiveBalance() (*big.Int, error) {
	return _ContractAvsGovernance.Contract.MaxEffectiveBalance(&_ContractAvsGovernance.CallOpts)
}

// MaxEffectiveBalance is a free data retrieval call binding the contract method 0xa88171ee.
//
// Solidity: function maxEffectiveBalance() view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) MaxEffectiveBalance() (*big.Int, error) {
	return _ContractAvsGovernance.Contract.MaxEffectiveBalance(&_ContractAvsGovernance.CallOpts)
}

// MinSlashableStakePerStakingContract is a free data retrieval call binding the contract method 0x2860b259.
//
// Solidity: function minSlashableStakePerStakingContract(address _stakingContract) view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) MinSlashableStakePerStakingContract(opts *bind.CallOpts, _stakingContract common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "minSlashableStakePerStakingContract", _stakingContract)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinSlashableStakePerStakingContract is a free data retrieval call binding the contract method 0x2860b259.
//
// Solidity: function minSlashableStakePerStakingContract(address _stakingContract) view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) MinSlashableStakePerStakingContract(_stakingContract common.Address) (*big.Int, error) {
	return _ContractAvsGovernance.Contract.MinSlashableStakePerStakingContract(&_ContractAvsGovernance.CallOpts, _stakingContract)
}

// MinSlashableStakePerStakingContract is a free data retrieval call binding the contract method 0x2860b259.
//
// Solidity: function minSlashableStakePerStakingContract(address _stakingContract) view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) MinSlashableStakePerStakingContract(_stakingContract common.Address) (*big.Int, error) {
	return _ContractAvsGovernance.Contract.MinSlashableStakePerStakingContract(&_ContractAvsGovernance.CallOpts, _stakingContract)
}

// MinStakePerStakingContract is a free data retrieval call binding the contract method 0x88160c4c.
//
// Solidity: function minStakePerStakingContract(address _stakingContract) view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) MinStakePerStakingContract(opts *bind.CallOpts, _stakingContract common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "minStakePerStakingContract", _stakingContract)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStakePerStakingContract is a free data retrieval call binding the contract method 0x88160c4c.
//
// Solidity: function minStakePerStakingContract(address _stakingContract) view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) MinStakePerStakingContract(_stakingContract common.Address) (*big.Int, error) {
	return _ContractAvsGovernance.Contract.MinStakePerStakingContract(&_ContractAvsGovernance.CallOpts, _stakingContract)
}

// MinStakePerStakingContract is a free data retrieval call binding the contract method 0x88160c4c.
//
// Solidity: function minStakePerStakingContract(address _stakingContract) view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) MinStakePerStakingContract(_stakingContract common.Address) (*big.Int, error) {
	return _ContractAvsGovernance.Contract.MinStakePerStakingContract(&_ContractAvsGovernance.CallOpts, _stakingContract)
}

// MinVotingPower is a free data retrieval call binding the contract method 0x36fffde0.
//
// Solidity: function minVotingPower() view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) MinVotingPower(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "minVotingPower")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinVotingPower is a free data retrieval call binding the contract method 0x36fffde0.
//
// Solidity: function minVotingPower() view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) MinVotingPower() (*big.Int, error) {
	return _ContractAvsGovernance.Contract.MinVotingPower(&_ContractAvsGovernance.CallOpts)
}

// MinVotingPower is a free data retrieval call binding the contract method 0x36fffde0.
//
// Solidity: function minVotingPower() view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) MinVotingPower() (*big.Int, error) {
	return _ContractAvsGovernance.Contract.MinVotingPower(&_ContractAvsGovernance.CallOpts)
}

// Multiplier is a free data retrieval call binding the contract method 0xa9b3f8b7.
//
// Solidity: function multiplier(address _stakingContract) view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) Multiplier(opts *bind.CallOpts, _stakingContract common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "multiplier", _stakingContract)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Multiplier is a free data retrieval call binding the contract method 0xa9b3f8b7.
//
// Solidity: function multiplier(address _stakingContract) view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) Multiplier(_stakingContract common.Address) (*big.Int, error) {
	return _ContractAvsGovernance.Contract.Multiplier(&_ContractAvsGovernance.CallOpts, _stakingContract)
}

// Multiplier is a free data retrieval call binding the contract method 0xa9b3f8b7.
//
// Solidity: function multiplier(address _stakingContract) view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) Multiplier(_stakingContract common.Address) (*big.Int, error) {
	return _ContractAvsGovernance.Contract.Multiplier(&_ContractAvsGovernance.CallOpts, _stakingContract)
}

// NumOfActiveOperators is a free data retrieval call binding the contract method 0x7897dec3.
//
// Solidity: function numOfActiveOperators() view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) NumOfActiveOperators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "numOfActiveOperators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumOfActiveOperators is a free data retrieval call binding the contract method 0x7897dec3.
//
// Solidity: function numOfActiveOperators() view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) NumOfActiveOperators() (*big.Int, error) {
	return _ContractAvsGovernance.Contract.NumOfActiveOperators(&_ContractAvsGovernance.CallOpts)
}

// NumOfActiveOperators is a free data retrieval call binding the contract method 0x7897dec3.
//
// Solidity: function numOfActiveOperators() view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) NumOfActiveOperators() (*big.Int, error) {
	return _ContractAvsGovernance.Contract.NumOfActiveOperators(&_ContractAvsGovernance.CallOpts)
}

// P2pAuthenticationEnabled is a free data retrieval call binding the contract method 0xcab1a5ef.
//
// Solidity: function p2pAuthenticationEnabled() view returns(bool)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) P2pAuthenticationEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "p2pAuthenticationEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// P2pAuthenticationEnabled is a free data retrieval call binding the contract method 0xcab1a5ef.
//
// Solidity: function p2pAuthenticationEnabled() view returns(bool)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) P2pAuthenticationEnabled() (bool, error) {
	return _ContractAvsGovernance.Contract.P2pAuthenticationEnabled(&_ContractAvsGovernance.CallOpts)
}

// P2pAuthenticationEnabled is a free data retrieval call binding the contract method 0xcab1a5ef.
//
// Solidity: function p2pAuthenticationEnabled() view returns(bool)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) P2pAuthenticationEnabled() (bool, error) {
	return _ContractAvsGovernance.Contract.P2pAuthenticationEnabled(&_ContractAvsGovernance.CallOpts)
}

// RedistributionManager is a free data retrieval call binding the contract method 0x56e25cf8.
//
// Solidity: function redistributionManager() view returns(address)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) RedistributionManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "redistributionManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RedistributionManager is a free data retrieval call binding the contract method 0x56e25cf8.
//
// Solidity: function redistributionManager() view returns(address)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) RedistributionManager() (common.Address, error) {
	return _ContractAvsGovernance.Contract.RedistributionManager(&_ContractAvsGovernance.CallOpts)
}

// RedistributionManager is a free data retrieval call binding the contract method 0x56e25cf8.
//
// Solidity: function redistributionManager() view returns(address)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) RedistributionManager() (common.Address, error) {
	return _ContractAvsGovernance.Contract.RedistributionManager(&_ContractAvsGovernance.CallOpts)
}

// SlashableStakeWeight is a free data retrieval call binding the contract method 0x26f1cda7.
//
// Solidity: function slashableStakeWeight(address _stakingContract) view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) SlashableStakeWeight(opts *bind.CallOpts, _stakingContract common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "slashableStakeWeight", _stakingContract)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlashableStakeWeight is a free data retrieval call binding the contract method 0x26f1cda7.
//
// Solidity: function slashableStakeWeight(address _stakingContract) view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) SlashableStakeWeight(_stakingContract common.Address) (*big.Int, error) {
	return _ContractAvsGovernance.Contract.SlashableStakeWeight(&_ContractAvsGovernance.CallOpts, _stakingContract)
}

// SlashableStakeWeight is a free data retrieval call binding the contract method 0x26f1cda7.
//
// Solidity: function slashableStakeWeight(address _stakingContract) view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) SlashableStakeWeight(_stakingContract common.Address) (*big.Int, error) {
	return _ContractAvsGovernance.Contract.SlashableStakeWeight(&_ContractAvsGovernance.CallOpts, _stakingContract)
}

// StakingContractToFeed is a free data retrieval call binding the contract method 0xbce2c28d.
//
// Solidity: function stakingContractToFeed(address _stakingContract) view returns(address)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) StakingContractToFeed(opts *bind.CallOpts, _stakingContract common.Address) (common.Address, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "stakingContractToFeed", _stakingContract)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingContractToFeed is a free data retrieval call binding the contract method 0xbce2c28d.
//
// Solidity: function stakingContractToFeed(address _stakingContract) view returns(address)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) StakingContractToFeed(_stakingContract common.Address) (common.Address, error) {
	return _ContractAvsGovernance.Contract.StakingContractToFeed(&_ContractAvsGovernance.CallOpts, _stakingContract)
}

// StakingContractToFeed is a free data retrieval call binding the contract method 0xbce2c28d.
//
// Solidity: function stakingContractToFeed(address _stakingContract) view returns(address)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) StakingContractToFeed(_stakingContract common.Address) (common.Address, error) {
	return _ContractAvsGovernance.Contract.StakingContractToFeed(&_ContractAvsGovernance.CallOpts, _stakingContract)
}

// StakingContracts is a free data retrieval call binding the contract method 0x560d5484.
//
// Solidity: function stakingContracts() view returns(address[])
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) StakingContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "stakingContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// StakingContracts is a free data retrieval call binding the contract method 0x560d5484.
//
// Solidity: function stakingContracts() view returns(address[])
func (_ContractAvsGovernance *ContractAvsGovernanceSession) StakingContracts() ([]common.Address, error) {
	return _ContractAvsGovernance.Contract.StakingContracts(&_ContractAvsGovernance.CallOpts)
}

// StakingContracts is a free data retrieval call binding the contract method 0x560d5484.
//
// Solidity: function stakingContracts() view returns(address[])
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) StakingContracts() ([]common.Address, error) {
	return _ContractAvsGovernance.Contract.StakingContracts(&_ContractAvsGovernance.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ContractAvsGovernance.Contract.SupportsInterface(&_ContractAvsGovernance.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ContractAvsGovernance.Contract.SupportsInterface(&_ContractAvsGovernance.CallOpts, interfaceId)
}

// VotingPower is a free data retrieval call binding the contract method 0xc07473f6.
//
// Solidity: function votingPower(address _operator) view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) VotingPower(opts *bind.CallOpts, _operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "votingPower", _operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingPower is a free data retrieval call binding the contract method 0xc07473f6.
//
// Solidity: function votingPower(address _operator) view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) VotingPower(_operator common.Address) (*big.Int, error) {
	return _ContractAvsGovernance.Contract.VotingPower(&_ContractAvsGovernance.CallOpts, _operator)
}

// VotingPower is a free data retrieval call binding the contract method 0xc07473f6.
//
// Solidity: function votingPower(address _operator) view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) VotingPower(_operator common.Address) (*big.Int, error) {
	return _ContractAvsGovernance.Contract.VotingPower(&_ContractAvsGovernance.CallOpts, _operator)
}

// VotingPowerPerStakingContracts is a free data retrieval call binding the contract method 0xcbd969c9.
//
// Solidity: function votingPowerPerStakingContracts(address _operator, address[] _stakingContracts) view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceCaller) VotingPowerPerStakingContracts(opts *bind.CallOpts, _operator common.Address, _stakingContracts []common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractAvsGovernance.contract.Call(opts, &out, "votingPowerPerStakingContracts", _operator, _stakingContracts)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingPowerPerStakingContracts is a free data retrieval call binding the contract method 0xcbd969c9.
//
// Solidity: function votingPowerPerStakingContracts(address _operator, address[] _stakingContracts) view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceSession) VotingPowerPerStakingContracts(_operator common.Address, _stakingContracts []common.Address) (*big.Int, error) {
	return _ContractAvsGovernance.Contract.VotingPowerPerStakingContracts(&_ContractAvsGovernance.CallOpts, _operator, _stakingContracts)
}

// VotingPowerPerStakingContracts is a free data retrieval call binding the contract method 0xcbd969c9.
//
// Solidity: function votingPowerPerStakingContracts(address _operator, address[] _stakingContracts) view returns(uint256)
func (_ContractAvsGovernance *ContractAvsGovernanceCallerSession) VotingPowerPerStakingContracts(_operator common.Address, _stakingContracts []common.Address) (*big.Int, error) {
	return _ContractAvsGovernance.Contract.VotingPowerPerStakingContracts(&_ContractAvsGovernance.CallOpts, _operator, _stakingContracts)
}

// ApplyCustomSlashing is a paid mutator transaction binding the contract method 0x24f8c18b.
//
// Solidity: function applyCustomSlashing(address _operator, (address,uint8,uint256)[] _slashingStakingContractInfos) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) ApplyCustomSlashing(opts *bind.TransactOpts, _operator common.Address, _slashingStakingContractInfos []ISlashingConfigSlashingStakingContractInfo) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "applyCustomSlashing", _operator, _slashingStakingContractInfos)
}

// ApplyCustomSlashing is a paid mutator transaction binding the contract method 0x24f8c18b.
//
// Solidity: function applyCustomSlashing(address _operator, (address,uint8,uint256)[] _slashingStakingContractInfos) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) ApplyCustomSlashing(_operator common.Address, _slashingStakingContractInfos []ISlashingConfigSlashingStakingContractInfo) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.ApplyCustomSlashing(&_ContractAvsGovernance.TransactOpts, _operator, _slashingStakingContractInfos)
}

// ApplyCustomSlashing is a paid mutator transaction binding the contract method 0x24f8c18b.
//
// Solidity: function applyCustomSlashing(address _operator, (address,uint8,uint256)[] _slashingStakingContractInfos) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) ApplyCustomSlashing(_operator common.Address, _slashingStakingContractInfos []ISlashingConfigSlashingStakingContractInfo) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.ApplyCustomSlashing(&_ContractAvsGovernance.TransactOpts, _operator, _slashingStakingContractInfos)
}

// CompleteRewardsReceiverModification is a paid mutator transaction binding the contract method 0xe6474b0f.
//
// Solidity: function completeRewardsReceiverModification() returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) CompleteRewardsReceiverModification(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "completeRewardsReceiverModification")
}

// CompleteRewardsReceiverModification is a paid mutator transaction binding the contract method 0xe6474b0f.
//
// Solidity: function completeRewardsReceiverModification() returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) CompleteRewardsReceiverModification() (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.CompleteRewardsReceiverModification(&_ContractAvsGovernance.TransactOpts)
}

// CompleteRewardsReceiverModification is a paid mutator transaction binding the contract method 0xe6474b0f.
//
// Solidity: function completeRewardsReceiverModification() returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) CompleteRewardsReceiverModification() (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.CompleteRewardsReceiverModification(&_ContractAvsGovernance.TransactOpts)
}

// ExecuteVetoSlashRequests is a paid mutator transaction binding the contract method 0xe9dbe87f.
//
// Solidity: function executeVetoSlashRequests(uint256 _from, uint256 _to) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) ExecuteVetoSlashRequests(opts *bind.TransactOpts, _from *big.Int, _to *big.Int) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "executeVetoSlashRequests", _from, _to)
}

// ExecuteVetoSlashRequests is a paid mutator transaction binding the contract method 0xe9dbe87f.
//
// Solidity: function executeVetoSlashRequests(uint256 _from, uint256 _to) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) ExecuteVetoSlashRequests(_from *big.Int, _to *big.Int) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.ExecuteVetoSlashRequests(&_ContractAvsGovernance.TransactOpts, _from, _to)
}

// ExecuteVetoSlashRequests is a paid mutator transaction binding the contract method 0xe9dbe87f.
//
// Solidity: function executeVetoSlashRequests(uint256 _from, uint256 _to) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) ExecuteVetoSlashRequests(_from *big.Int, _to *big.Int) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.ExecuteVetoSlashRequests(&_ContractAvsGovernance.TransactOpts, _from, _to)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.GrantRole(&_ContractAvsGovernance.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.GrantRole(&_ContractAvsGovernance.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xba3bf212.
//
// Solidity: function initialize((address,address,address,address,address,address,address,address,string,address,address) _initializationParams) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) Initialize(opts *bind.TransactOpts, _initializationParams IAvsGovernanceInitializationParams) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "initialize", _initializationParams)
}

// Initialize is a paid mutator transaction binding the contract method 0xba3bf212.
//
// Solidity: function initialize((address,address,address,address,address,address,address,address,string,address,address) _initializationParams) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) Initialize(_initializationParams IAvsGovernanceInitializationParams) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.Initialize(&_ContractAvsGovernance.TransactOpts, _initializationParams)
}

// Initialize is a paid mutator transaction binding the contract method 0xba3bf212.
//
// Solidity: function initialize((address,address,address,address,address,address,address,address,string,address,address) _initializationParams) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) Initialize(_initializationParams IAvsGovernanceInitializationParams) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.Initialize(&_ContractAvsGovernance.TransactOpts, _initializationParams)
}

// MigrateAvsToAllocationManager is a paid mutator transaction binding the contract method 0x7cfd4974.
//
// Solidity: function migrateAvsToAllocationManager(string _metadataURI) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) MigrateAvsToAllocationManager(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "migrateAvsToAllocationManager", _metadataURI)
}

// MigrateAvsToAllocationManager is a paid mutator transaction binding the contract method 0x7cfd4974.
//
// Solidity: function migrateAvsToAllocationManager(string _metadataURI) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) MigrateAvsToAllocationManager(_metadataURI string) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.MigrateAvsToAllocationManager(&_ContractAvsGovernance.TransactOpts, _metadataURI)
}

// MigrateAvsToAllocationManager is a paid mutator transaction binding the contract method 0x7cfd4974.
//
// Solidity: function migrateAvsToAllocationManager(string _metadataURI) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) MigrateAvsToAllocationManager(_metadataURI string) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.MigrateAvsToAllocationManager(&_ContractAvsGovernance.TransactOpts, _metadataURI)
}

// MigrateRedistributionManager is a paid mutator transaction binding the contract method 0x35391c12.
//
// Solidity: function migrateRedistributionManager(address _l1AvsFactory, address _avsGovernanceMultisigOwner) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) MigrateRedistributionManager(opts *bind.TransactOpts, _l1AvsFactory common.Address, _avsGovernanceMultisigOwner common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "migrateRedistributionManager", _l1AvsFactory, _avsGovernanceMultisigOwner)
}

// MigrateRedistributionManager is a paid mutator transaction binding the contract method 0x35391c12.
//
// Solidity: function migrateRedistributionManager(address _l1AvsFactory, address _avsGovernanceMultisigOwner) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) MigrateRedistributionManager(_l1AvsFactory common.Address, _avsGovernanceMultisigOwner common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.MigrateRedistributionManager(&_ContractAvsGovernance.TransactOpts, _l1AvsFactory, _avsGovernanceMultisigOwner)
}

// MigrateRedistributionManager is a paid mutator transaction binding the contract method 0x35391c12.
//
// Solidity: function migrateRedistributionManager(address _l1AvsFactory, address _avsGovernanceMultisigOwner) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) MigrateRedistributionManager(_l1AvsFactory common.Address, _avsGovernanceMultisigOwner common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.MigrateRedistributionManager(&_ContractAvsGovernance.TransactOpts, _l1AvsFactory, _avsGovernanceMultisigOwner)
}

// Migration is a paid mutator transaction binding the contract method 0x1705a3bd.
//
// Solidity: function migration() returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) Migration(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "migration")
}

// Migration is a paid mutator transaction binding the contract method 0x1705a3bd.
//
// Solidity: function migration() returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) Migration() (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.Migration(&_ContractAvsGovernance.TransactOpts)
}

// Migration is a paid mutator transaction binding the contract method 0x1705a3bd.
//
// Solidity: function migration() returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) Migration() (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.Migration(&_ContractAvsGovernance.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x3aa83ec7.
//
// Solidity: function pause(bytes4 _pausableFlow) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) Pause(opts *bind.TransactOpts, _pausableFlow [4]byte) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "pause", _pausableFlow)
}

// Pause is a paid mutator transaction binding the contract method 0x3aa83ec7.
//
// Solidity: function pause(bytes4 _pausableFlow) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) Pause(_pausableFlow [4]byte) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.Pause(&_ContractAvsGovernance.TransactOpts, _pausableFlow)
}

// Pause is a paid mutator transaction binding the contract method 0x3aa83ec7.
//
// Solidity: function pause(bytes4 _pausableFlow) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) Pause(_pausableFlow [4]byte) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.Pause(&_ContractAvsGovernance.TransactOpts, _pausableFlow)
}

// QueueRewardsReceiverModification is a paid mutator transaction binding the contract method 0x1b21ba72.
//
// Solidity: function queueRewardsReceiverModification(address _newRewardsReceiver) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) QueueRewardsReceiverModification(opts *bind.TransactOpts, _newRewardsReceiver common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "queueRewardsReceiverModification", _newRewardsReceiver)
}

// QueueRewardsReceiverModification is a paid mutator transaction binding the contract method 0x1b21ba72.
//
// Solidity: function queueRewardsReceiverModification(address _newRewardsReceiver) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) QueueRewardsReceiverModification(_newRewardsReceiver common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.QueueRewardsReceiverModification(&_ContractAvsGovernance.TransactOpts, _newRewardsReceiver)
}

// QueueRewardsReceiverModification is a paid mutator transaction binding the contract method 0x1b21ba72.
//
// Solidity: function queueRewardsReceiverModification(address _newRewardsReceiver) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) QueueRewardsReceiverModification(_newRewardsReceiver common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.QueueRewardsReceiverModification(&_ContractAvsGovernance.TransactOpts, _newRewardsReceiver)
}

// RegisterAvsToEigenLayer is a paid mutator transaction binding the contract method 0x851bb725.
//
// Solidity: function registerAvsToEigenLayer(string metadataURI) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) RegisterAvsToEigenLayer(opts *bind.TransactOpts, metadataURI string) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "registerAvsToEigenLayer", metadataURI)
}

// RegisterAvsToEigenLayer is a paid mutator transaction binding the contract method 0x851bb725.
//
// Solidity: function registerAvsToEigenLayer(string metadataURI) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) RegisterAvsToEigenLayer(metadataURI string) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.RegisterAvsToEigenLayer(&_ContractAvsGovernance.TransactOpts, metadataURI)
}

// RegisterAvsToEigenLayer is a paid mutator transaction binding the contract method 0x851bb725.
//
// Solidity: function registerAvsToEigenLayer(string metadataURI) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) RegisterAvsToEigenLayer(metadataURI string) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.RegisterAvsToEigenLayer(&_ContractAvsGovernance.TransactOpts, metadataURI)
}

// RegisterAvsToSymbiotic is a paid mutator transaction binding the contract method 0xaa0dadf0.
//
// Solidity: function registerAvsToSymbiotic() returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) RegisterAvsToSymbiotic(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "registerAvsToSymbiotic")
}

// RegisterAvsToSymbiotic is a paid mutator transaction binding the contract method 0xaa0dadf0.
//
// Solidity: function registerAvsToSymbiotic() returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) RegisterAvsToSymbiotic() (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.RegisterAvsToSymbiotic(&_ContractAvsGovernance.TransactOpts)
}

// RegisterAvsToSymbiotic is a paid mutator transaction binding the contract method 0xaa0dadf0.
//
// Solidity: function registerAvsToSymbiotic() returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) RegisterAvsToSymbiotic() (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.RegisterAvsToSymbiotic(&_ContractAvsGovernance.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.RenounceRole(&_ContractAvsGovernance.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.RenounceRole(&_ContractAvsGovernance.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.RevokeRole(&_ContractAvsGovernance.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.RevokeRole(&_ContractAvsGovernance.TransactOpts, role, account)
}

// SetAvsGovernanceLogic is a paid mutator transaction binding the contract method 0x8987c767.
//
// Solidity: function setAvsGovernanceLogic(address _avsGovernanceLogic) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) SetAvsGovernanceLogic(opts *bind.TransactOpts, _avsGovernanceLogic common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "setAvsGovernanceLogic", _avsGovernanceLogic)
}

// SetAvsGovernanceLogic is a paid mutator transaction binding the contract method 0x8987c767.
//
// Solidity: function setAvsGovernanceLogic(address _avsGovernanceLogic) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) SetAvsGovernanceLogic(_avsGovernanceLogic common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetAvsGovernanceLogic(&_ContractAvsGovernance.TransactOpts, _avsGovernanceLogic)
}

// SetAvsGovernanceLogic is a paid mutator transaction binding the contract method 0x8987c767.
//
// Solidity: function setAvsGovernanceLogic(address _avsGovernanceLogic) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) SetAvsGovernanceLogic(_avsGovernanceLogic common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetAvsGovernanceLogic(&_ContractAvsGovernance.TransactOpts, _avsGovernanceLogic)
}

// SetIsAllowlisted is a paid mutator transaction binding the contract method 0x9e965cc1.
//
// Solidity: function setIsAllowlisted(bool _isAllowlisted) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) SetIsAllowlisted(opts *bind.TransactOpts, _isAllowlisted bool) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "setIsAllowlisted", _isAllowlisted)
}

// SetIsAllowlisted is a paid mutator transaction binding the contract method 0x9e965cc1.
//
// Solidity: function setIsAllowlisted(bool _isAllowlisted) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) SetIsAllowlisted(_isAllowlisted bool) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetIsAllowlisted(&_ContractAvsGovernance.TransactOpts, _isAllowlisted)
}

// SetIsAllowlisted is a paid mutator transaction binding the contract method 0x9e965cc1.
//
// Solidity: function setIsAllowlisted(bool _isAllowlisted) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) SetIsAllowlisted(_isAllowlisted bool) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetIsAllowlisted(&_ContractAvsGovernance.TransactOpts, _isAllowlisted)
}

// SetP2pAuthenticationEnabled is a paid mutator transaction binding the contract method 0xc7a93c6f.
//
// Solidity: function setP2pAuthenticationEnabled(bool _p2pAuthenticationEnabled) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) SetP2pAuthenticationEnabled(opts *bind.TransactOpts, _p2pAuthenticationEnabled bool) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "setP2pAuthenticationEnabled", _p2pAuthenticationEnabled)
}

// SetP2pAuthenticationEnabled is a paid mutator transaction binding the contract method 0xc7a93c6f.
//
// Solidity: function setP2pAuthenticationEnabled(bool _p2pAuthenticationEnabled) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) SetP2pAuthenticationEnabled(_p2pAuthenticationEnabled bool) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetP2pAuthenticationEnabled(&_ContractAvsGovernance.TransactOpts, _p2pAuthenticationEnabled)
}

// SetP2pAuthenticationEnabled is a paid mutator transaction binding the contract method 0xc7a93c6f.
//
// Solidity: function setP2pAuthenticationEnabled(bool _p2pAuthenticationEnabled) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) SetP2pAuthenticationEnabled(_p2pAuthenticationEnabled bool) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetP2pAuthenticationEnabled(&_ContractAvsGovernance.TransactOpts, _p2pAuthenticationEnabled)
}

// SetStakingContractMultiplier is a paid mutator transaction binding the contract method 0x2ba09e51.
//
// Solidity: function setStakingContractMultiplier((address,uint256,uint256,uint8) _votingPowerMultiplier) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) SetStakingContractMultiplier(opts *bind.TransactOpts, _votingPowerMultiplier IAvsGovernanceVotingPowerMultiplier) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "setStakingContractMultiplier", _votingPowerMultiplier)
}

// SetStakingContractMultiplier is a paid mutator transaction binding the contract method 0x2ba09e51.
//
// Solidity: function setStakingContractMultiplier((address,uint256,uint256,uint8) _votingPowerMultiplier) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) SetStakingContractMultiplier(_votingPowerMultiplier IAvsGovernanceVotingPowerMultiplier) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetStakingContractMultiplier(&_ContractAvsGovernance.TransactOpts, _votingPowerMultiplier)
}

// SetStakingContractMultiplier is a paid mutator transaction binding the contract method 0x2ba09e51.
//
// Solidity: function setStakingContractMultiplier((address,uint256,uint256,uint8) _votingPowerMultiplier) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) SetStakingContractMultiplier(_votingPowerMultiplier IAvsGovernanceVotingPowerMultiplier) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetStakingContractMultiplier(&_ContractAvsGovernance.TransactOpts, _votingPowerMultiplier)
}

// SetStakingContractMultiplierBatch is a paid mutator transaction binding the contract method 0x94643671.
//
// Solidity: function setStakingContractMultiplierBatch((address,uint256,uint256,uint8)[] _votingPowerMultipliers) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) SetStakingContractMultiplierBatch(opts *bind.TransactOpts, _votingPowerMultipliers []IAvsGovernanceVotingPowerMultiplier) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "setStakingContractMultiplierBatch", _votingPowerMultipliers)
}

// SetStakingContractMultiplierBatch is a paid mutator transaction binding the contract method 0x94643671.
//
// Solidity: function setStakingContractMultiplierBatch((address,uint256,uint256,uint8)[] _votingPowerMultipliers) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) SetStakingContractMultiplierBatch(_votingPowerMultipliers []IAvsGovernanceVotingPowerMultiplier) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetStakingContractMultiplierBatch(&_ContractAvsGovernance.TransactOpts, _votingPowerMultipliers)
}

// SetStakingContractMultiplierBatch is a paid mutator transaction binding the contract method 0x94643671.
//
// Solidity: function setStakingContractMultiplierBatch((address,uint256,uint256,uint8)[] _votingPowerMultipliers) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) SetStakingContractMultiplierBatch(_votingPowerMultipliers []IAvsGovernanceVotingPowerMultiplier) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetStakingContractMultiplierBatch(&_ContractAvsGovernance.TransactOpts, _votingPowerMultipliers)
}

// SetStakingContractPriceFeed is a paid mutator transaction binding the contract method 0xf19b7b79.
//
// Solidity: function setStakingContractPriceFeed(address _stakingContract, address _feed) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) SetStakingContractPriceFeed(opts *bind.TransactOpts, _stakingContract common.Address, _feed common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "setStakingContractPriceFeed", _stakingContract, _feed)
}

// SetStakingContractPriceFeed is a paid mutator transaction binding the contract method 0xf19b7b79.
//
// Solidity: function setStakingContractPriceFeed(address _stakingContract, address _feed) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) SetStakingContractPriceFeed(_stakingContract common.Address, _feed common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetStakingContractPriceFeed(&_ContractAvsGovernance.TransactOpts, _stakingContract, _feed)
}

// SetStakingContractPriceFeed is a paid mutator transaction binding the contract method 0xf19b7b79.
//
// Solidity: function setStakingContractPriceFeed(address _stakingContract, address _feed) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) SetStakingContractPriceFeed(_stakingContract common.Address, _feed common.Address) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetStakingContractPriceFeed(&_ContractAvsGovernance.TransactOpts, _stakingContract, _feed)
}

// SetSupportedStakingContracts is a paid mutator transaction binding the contract method 0x1dd2f74b.
//
// Solidity: function setSupportedStakingContracts((address,uint8)[] _stakingContractsDetails) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) SetSupportedStakingContracts(opts *bind.TransactOpts, _stakingContractsDetails []IAvsGovernanceStakingContractInfo) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "setSupportedStakingContracts", _stakingContractsDetails)
}

// SetSupportedStakingContracts is a paid mutator transaction binding the contract method 0x1dd2f74b.
//
// Solidity: function setSupportedStakingContracts((address,uint8)[] _stakingContractsDetails) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) SetSupportedStakingContracts(_stakingContractsDetails []IAvsGovernanceStakingContractInfo) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetSupportedStakingContracts(&_ContractAvsGovernance.TransactOpts, _stakingContractsDetails)
}

// SetSupportedStakingContracts is a paid mutator transaction binding the contract method 0x1dd2f74b.
//
// Solidity: function setSupportedStakingContracts((address,uint8)[] _stakingContractsDetails) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) SetSupportedStakingContracts(_stakingContractsDetails []IAvsGovernanceStakingContractInfo) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetSupportedStakingContracts(&_ContractAvsGovernance.TransactOpts, _stakingContractsDetails)
}

// SetSymbioticResolver is a paid mutator transaction binding the contract method 0x438c20a1.
//
// Solidity: function setSymbioticResolver(address _vetoSlasher, address _resolver, bytes _hints) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) SetSymbioticResolver(opts *bind.TransactOpts, _vetoSlasher common.Address, _resolver common.Address, _hints []byte) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "setSymbioticResolver", _vetoSlasher, _resolver, _hints)
}

// SetSymbioticResolver is a paid mutator transaction binding the contract method 0x438c20a1.
//
// Solidity: function setSymbioticResolver(address _vetoSlasher, address _resolver, bytes _hints) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) SetSymbioticResolver(_vetoSlasher common.Address, _resolver common.Address, _hints []byte) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetSymbioticResolver(&_ContractAvsGovernance.TransactOpts, _vetoSlasher, _resolver, _hints)
}

// SetSymbioticResolver is a paid mutator transaction binding the contract method 0x438c20a1.
//
// Solidity: function setSymbioticResolver(address _vetoSlasher, address _resolver, bytes _hints) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) SetSymbioticResolver(_vetoSlasher common.Address, _resolver common.Address, _hints []byte) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SetSymbioticResolver(&_ContractAvsGovernance.TransactOpts, _vetoSlasher, _resolver, _hints)
}

// SlashOperator is a paid mutator transaction binding the contract method 0x7f166721.
//
// Solidity: function slashOperator(address _operator, uint8 _slashingCondition) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) SlashOperator(opts *bind.TransactOpts, _operator common.Address, _slashingCondition uint8) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "slashOperator", _operator, _slashingCondition)
}

// SlashOperator is a paid mutator transaction binding the contract method 0x7f166721.
//
// Solidity: function slashOperator(address _operator, uint8 _slashingCondition) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) SlashOperator(_operator common.Address, _slashingCondition uint8) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SlashOperator(&_ContractAvsGovernance.TransactOpts, _operator, _slashingCondition)
}

// SlashOperator is a paid mutator transaction binding the contract method 0x7f166721.
//
// Solidity: function slashOperator(address _operator, uint8 _slashingCondition) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) SlashOperator(_operator common.Address, _slashingCondition uint8) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.SlashOperator(&_ContractAvsGovernance.TransactOpts, _operator, _slashingCondition)
}

// Unpause is a paid mutator transaction binding the contract method 0xbac1e94b.
//
// Solidity: function unpause(bytes4 _pausableFlow) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) Unpause(opts *bind.TransactOpts, _pausableFlow [4]byte) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.Transact(opts, "unpause", _pausableFlow)
}

// Unpause is a paid mutator transaction binding the contract method 0xbac1e94b.
//
// Solidity: function unpause(bytes4 _pausableFlow) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) Unpause(_pausableFlow [4]byte) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.Unpause(&_ContractAvsGovernance.TransactOpts, _pausableFlow)
}

// Unpause is a paid mutator transaction binding the contract method 0xbac1e94b.
//
// Solidity: function unpause(bytes4 _pausableFlow) returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) Unpause(_pausableFlow [4]byte) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.Unpause(&_ContractAvsGovernance.TransactOpts, _pausableFlow)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _ContractAvsGovernance.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_ContractAvsGovernance *ContractAvsGovernanceSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.Fallback(&_ContractAvsGovernance.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_ContractAvsGovernance *ContractAvsGovernanceTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ContractAvsGovernance.Contract.Fallback(&_ContractAvsGovernance.TransactOpts, calldata)
}

// ContractAvsGovernanceEjectOperatorFailedIterator is returned from FilterEjectOperatorFailed and is used to iterate over the raw logs and unpacked data for EjectOperatorFailed events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceEjectOperatorFailedIterator struct {
	Event *ContractAvsGovernanceEjectOperatorFailed // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceEjectOperatorFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceEjectOperatorFailed)
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
		it.Event = new(ContractAvsGovernanceEjectOperatorFailed)
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
func (it *ContractAvsGovernanceEjectOperatorFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceEjectOperatorFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceEjectOperatorFailed represents a EjectOperatorFailed event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceEjectOperatorFailed struct {
	Operator common.Address
	Data     []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEjectOperatorFailed is a free log retrieval operation binding the contract event 0x23b717cc29cc815f01bb575983e2bc61f85927ab73ed7169c7c3e62b31429965.
//
// Solidity: event EjectOperatorFailed(address operator, bytes data)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterEjectOperatorFailed(opts *bind.FilterOpts) (*ContractAvsGovernanceEjectOperatorFailedIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "EjectOperatorFailed")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceEjectOperatorFailedIterator{contract: _ContractAvsGovernance.contract, event: "EjectOperatorFailed", logs: logs, sub: sub}, nil
}

// WatchEjectOperatorFailed is a free log subscription operation binding the contract event 0x23b717cc29cc815f01bb575983e2bc61f85927ab73ed7169c7c3e62b31429965.
//
// Solidity: event EjectOperatorFailed(address operator, bytes data)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchEjectOperatorFailed(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceEjectOperatorFailed) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "EjectOperatorFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceEjectOperatorFailed)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "EjectOperatorFailed", log); err != nil {
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

// ParseEjectOperatorFailed is a log parse operation binding the contract event 0x23b717cc29cc815f01bb575983e2bc61f85927ab73ed7169c7c3e62b31429965.
//
// Solidity: event EjectOperatorFailed(address operator, bytes data)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseEjectOperatorFailed(log types.Log) (*ContractAvsGovernanceEjectOperatorFailed, error) {
	event := new(ContractAvsGovernanceEjectOperatorFailed)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "EjectOperatorFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceFlowPausedIterator is returned from FilterFlowPaused and is used to iterate over the raw logs and unpacked data for FlowPaused events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceFlowPausedIterator struct {
	Event *ContractAvsGovernanceFlowPaused // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceFlowPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceFlowPaused)
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
		it.Event = new(ContractAvsGovernanceFlowPaused)
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
func (it *ContractAvsGovernanceFlowPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceFlowPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceFlowPaused represents a FlowPaused event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceFlowPaused struct {
	PausableFlow [4]byte
	Pauser       common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFlowPaused is a free log retrieval operation binding the contract event 0x95c3658c5e0c74e20cf12db371b9b67d26e97a1937f6d2284f88cc44d036b4f6.
//
// Solidity: event FlowPaused(bytes4 _pausableFlow, address _pauser)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterFlowPaused(opts *bind.FilterOpts) (*ContractAvsGovernanceFlowPausedIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "FlowPaused")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceFlowPausedIterator{contract: _ContractAvsGovernance.contract, event: "FlowPaused", logs: logs, sub: sub}, nil
}

// WatchFlowPaused is a free log subscription operation binding the contract event 0x95c3658c5e0c74e20cf12db371b9b67d26e97a1937f6d2284f88cc44d036b4f6.
//
// Solidity: event FlowPaused(bytes4 _pausableFlow, address _pauser)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchFlowPaused(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceFlowPaused) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "FlowPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceFlowPaused)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "FlowPaused", log); err != nil {
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

// ParseFlowPaused is a log parse operation binding the contract event 0x95c3658c5e0c74e20cf12db371b9b67d26e97a1937f6d2284f88cc44d036b4f6.
//
// Solidity: event FlowPaused(bytes4 _pausableFlow, address _pauser)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseFlowPaused(log types.Log) (*ContractAvsGovernanceFlowPaused, error) {
	event := new(ContractAvsGovernanceFlowPaused)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "FlowPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceFlowUnpausedIterator is returned from FilterFlowUnpaused and is used to iterate over the raw logs and unpacked data for FlowUnpaused events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceFlowUnpausedIterator struct {
	Event *ContractAvsGovernanceFlowUnpaused // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceFlowUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceFlowUnpaused)
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
		it.Event = new(ContractAvsGovernanceFlowUnpaused)
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
func (it *ContractAvsGovernanceFlowUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceFlowUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceFlowUnpaused represents a FlowUnpaused event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceFlowUnpaused struct {
	PausableFlowFlag [4]byte
	Unpauser         common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterFlowUnpaused is a free log retrieval operation binding the contract event 0xc7e56e17b0a6c4b467df6495e1eda1baecd7ba20604e80c1058ac06f4578d85e.
//
// Solidity: event FlowUnpaused(bytes4 _pausableFlowFlag, address _unpauser)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterFlowUnpaused(opts *bind.FilterOpts) (*ContractAvsGovernanceFlowUnpausedIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "FlowUnpaused")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceFlowUnpausedIterator{contract: _ContractAvsGovernance.contract, event: "FlowUnpaused", logs: logs, sub: sub}, nil
}

// WatchFlowUnpaused is a free log subscription operation binding the contract event 0xc7e56e17b0a6c4b467df6495e1eda1baecd7ba20604e80c1058ac06f4578d85e.
//
// Solidity: event FlowUnpaused(bytes4 _pausableFlowFlag, address _unpauser)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchFlowUnpaused(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceFlowUnpaused) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "FlowUnpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceFlowUnpaused)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "FlowUnpaused", log); err != nil {
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

// ParseFlowUnpaused is a log parse operation binding the contract event 0xc7e56e17b0a6c4b467df6495e1eda1baecd7ba20604e80c1058ac06f4578d85e.
//
// Solidity: event FlowUnpaused(bytes4 _pausableFlowFlag, address _unpauser)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseFlowUnpaused(log types.Log) (*ContractAvsGovernanceFlowUnpaused, error) {
	event := new(ContractAvsGovernanceFlowUnpaused)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "FlowUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceInitializedIterator struct {
	Event *ContractAvsGovernanceInitialized // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceInitialized)
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
		it.Event = new(ContractAvsGovernanceInitialized)
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
func (it *ContractAvsGovernanceInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceInitialized represents a Initialized event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractAvsGovernanceInitializedIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceInitializedIterator{contract: _ContractAvsGovernance.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceInitialized)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseInitialized(log types.Log) (*ContractAvsGovernanceInitialized, error) {
	event := new(ContractAvsGovernanceInitialized)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceInvalidStakingContractsForSlashingIterator is returned from FilterInvalidStakingContractsForSlashing and is used to iterate over the raw logs and unpacked data for InvalidStakingContractsForSlashing events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceInvalidStakingContractsForSlashingIterator struct {
	Event *ContractAvsGovernanceInvalidStakingContractsForSlashing // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceInvalidStakingContractsForSlashingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceInvalidStakingContractsForSlashing)
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
		it.Event = new(ContractAvsGovernanceInvalidStakingContractsForSlashing)
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
func (it *ContractAvsGovernanceInvalidStakingContractsForSlashingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceInvalidStakingContractsForSlashingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceInvalidStakingContractsForSlashing represents a InvalidStakingContractsForSlashing event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceInvalidStakingContractsForSlashing struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterInvalidStakingContractsForSlashing is a free log retrieval operation binding the contract event 0xe37900f24e06ac44ab1396d92b2242d16769efc0cd66a368423719d9c1015fea.
//
// Solidity: event InvalidStakingContractsForSlashing()
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterInvalidStakingContractsForSlashing(opts *bind.FilterOpts) (*ContractAvsGovernanceInvalidStakingContractsForSlashingIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "InvalidStakingContractsForSlashing")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceInvalidStakingContractsForSlashingIterator{contract: _ContractAvsGovernance.contract, event: "InvalidStakingContractsForSlashing", logs: logs, sub: sub}, nil
}

// WatchInvalidStakingContractsForSlashing is a free log subscription operation binding the contract event 0xe37900f24e06ac44ab1396d92b2242d16769efc0cd66a368423719d9c1015fea.
//
// Solidity: event InvalidStakingContractsForSlashing()
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchInvalidStakingContractsForSlashing(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceInvalidStakingContractsForSlashing) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "InvalidStakingContractsForSlashing")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceInvalidStakingContractsForSlashing)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "InvalidStakingContractsForSlashing", log); err != nil {
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

// ParseInvalidStakingContractsForSlashing is a log parse operation binding the contract event 0xe37900f24e06ac44ab1396d92b2242d16769efc0cd66a368423719d9c1015fea.
//
// Solidity: event InvalidStakingContractsForSlashing()
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseInvalidStakingContractsForSlashing(log types.Log) (*ContractAvsGovernanceInvalidStakingContractsForSlashing, error) {
	event := new(ContractAvsGovernanceInvalidStakingContractsForSlashing)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "InvalidStakingContractsForSlashing", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceMinSlashableStakePerStakingContractSetIterator is returned from FilterMinSlashableStakePerStakingContractSet and is used to iterate over the raw logs and unpacked data for MinSlashableStakePerStakingContractSet events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceMinSlashableStakePerStakingContractSetIterator struct {
	Event *ContractAvsGovernanceMinSlashableStakePerStakingContractSet // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceMinSlashableStakePerStakingContractSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceMinSlashableStakePerStakingContractSet)
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
		it.Event = new(ContractAvsGovernanceMinSlashableStakePerStakingContractSet)
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
func (it *ContractAvsGovernanceMinSlashableStakePerStakingContractSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceMinSlashableStakePerStakingContractSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceMinSlashableStakePerStakingContractSet represents a MinSlashableStakePerStakingContractSet event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceMinSlashableStakePerStakingContractSet struct {
	StakingContract   common.Address
	MinSlashableStake *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterMinSlashableStakePerStakingContractSet is a free log retrieval operation binding the contract event 0xc57761e5f61d8780398fffb71ecce8e2c362424f54cd9a9afee479a82b3c1cc2.
//
// Solidity: event MinSlashableStakePerStakingContractSet(address stakingContract, uint256 minSlashableStake)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterMinSlashableStakePerStakingContractSet(opts *bind.FilterOpts) (*ContractAvsGovernanceMinSlashableStakePerStakingContractSetIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "MinSlashableStakePerStakingContractSet")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceMinSlashableStakePerStakingContractSetIterator{contract: _ContractAvsGovernance.contract, event: "MinSlashableStakePerStakingContractSet", logs: logs, sub: sub}, nil
}

// WatchMinSlashableStakePerStakingContractSet is a free log subscription operation binding the contract event 0xc57761e5f61d8780398fffb71ecce8e2c362424f54cd9a9afee479a82b3c1cc2.
//
// Solidity: event MinSlashableStakePerStakingContractSet(address stakingContract, uint256 minSlashableStake)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchMinSlashableStakePerStakingContractSet(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceMinSlashableStakePerStakingContractSet) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "MinSlashableStakePerStakingContractSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceMinSlashableStakePerStakingContractSet)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "MinSlashableStakePerStakingContractSet", log); err != nil {
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

// ParseMinSlashableStakePerStakingContractSet is a log parse operation binding the contract event 0xc57761e5f61d8780398fffb71ecce8e2c362424f54cd9a9afee479a82b3c1cc2.
//
// Solidity: event MinSlashableStakePerStakingContractSet(address stakingContract, uint256 minSlashableStake)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseMinSlashableStakePerStakingContractSet(log types.Log) (*ContractAvsGovernanceMinSlashableStakePerStakingContractSet, error) {
	event := new(ContractAvsGovernanceMinSlashableStakePerStakingContractSet)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "MinSlashableStakePerStakingContractSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceMinStakePerStakingContractSetIterator is returned from FilterMinStakePerStakingContractSet and is used to iterate over the raw logs and unpacked data for MinStakePerStakingContractSet events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceMinStakePerStakingContractSetIterator struct {
	Event *ContractAvsGovernanceMinStakePerStakingContractSet // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceMinStakePerStakingContractSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceMinStakePerStakingContractSet)
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
		it.Event = new(ContractAvsGovernanceMinStakePerStakingContractSet)
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
func (it *ContractAvsGovernanceMinStakePerStakingContractSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceMinStakePerStakingContractSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceMinStakePerStakingContractSet represents a MinStakePerStakingContractSet event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceMinStakePerStakingContractSet struct {
	StakingContract common.Address
	MinStake        *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMinStakePerStakingContractSet is a free log retrieval operation binding the contract event 0xd0fb069b8b4156af8735f9d96d43dc06680897ecd16849e5f66f32beb4f60e26.
//
// Solidity: event MinStakePerStakingContractSet(address stakingContract, uint256 minStake)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterMinStakePerStakingContractSet(opts *bind.FilterOpts) (*ContractAvsGovernanceMinStakePerStakingContractSetIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "MinStakePerStakingContractSet")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceMinStakePerStakingContractSetIterator{contract: _ContractAvsGovernance.contract, event: "MinStakePerStakingContractSet", logs: logs, sub: sub}, nil
}

// WatchMinStakePerStakingContractSet is a free log subscription operation binding the contract event 0xd0fb069b8b4156af8735f9d96d43dc06680897ecd16849e5f66f32beb4f60e26.
//
// Solidity: event MinStakePerStakingContractSet(address stakingContract, uint256 minStake)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchMinStakePerStakingContractSet(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceMinStakePerStakingContractSet) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "MinStakePerStakingContractSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceMinStakePerStakingContractSet)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "MinStakePerStakingContractSet", log); err != nil {
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

// ParseMinStakePerStakingContractSet is a log parse operation binding the contract event 0xd0fb069b8b4156af8735f9d96d43dc06680897ecd16849e5f66f32beb4f60e26.
//
// Solidity: event MinStakePerStakingContractSet(address stakingContract, uint256 minStake)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseMinStakePerStakingContractSet(log types.Log) (*ContractAvsGovernanceMinStakePerStakingContractSet, error) {
	event := new(ContractAvsGovernanceMinStakePerStakingContractSet)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "MinStakePerStakingContractSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceOperatorSlashedIterator is returned from FilterOperatorSlashed and is used to iterate over the raw logs and unpacked data for OperatorSlashed events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceOperatorSlashedIterator struct {
	Event *ContractAvsGovernanceOperatorSlashed // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceOperatorSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceOperatorSlashed)
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
		it.Event = new(ContractAvsGovernanceOperatorSlashed)
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
func (it *ContractAvsGovernanceOperatorSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceOperatorSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceOperatorSlashed represents a OperatorSlashed event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceOperatorSlashed struct {
	Operator common.Address
	SlashId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorSlashed is a free log retrieval operation binding the contract event 0xa005d3b8f5c8149659f2afc06ed99732d5891377a370a723669418ca1ca29ced.
//
// Solidity: event OperatorSlashed(address operator, uint256 slashId)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterOperatorSlashed(opts *bind.FilterOpts) (*ContractAvsGovernanceOperatorSlashedIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "OperatorSlashed")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceOperatorSlashedIterator{contract: _ContractAvsGovernance.contract, event: "OperatorSlashed", logs: logs, sub: sub}, nil
}

// WatchOperatorSlashed is a free log subscription operation binding the contract event 0xa005d3b8f5c8149659f2afc06ed99732d5891377a370a723669418ca1ca29ced.
//
// Solidity: event OperatorSlashed(address operator, uint256 slashId)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchOperatorSlashed(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceOperatorSlashed) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "OperatorSlashed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceOperatorSlashed)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "OperatorSlashed", log); err != nil {
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

// ParseOperatorSlashed is a log parse operation binding the contract event 0xa005d3b8f5c8149659f2afc06ed99732d5891377a370a723669418ca1ca29ced.
//
// Solidity: event OperatorSlashed(address operator, uint256 slashId)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseOperatorSlashed(log types.Log) (*ContractAvsGovernanceOperatorSlashed, error) {
	event := new(ContractAvsGovernanceOperatorSlashed)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "OperatorSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernancePriceFeedSetIterator is returned from FilterPriceFeedSet and is used to iterate over the raw logs and unpacked data for PriceFeedSet events raised by the ContractAvsGovernance contract.
type ContractAvsGovernancePriceFeedSetIterator struct {
	Event *ContractAvsGovernancePriceFeedSet // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernancePriceFeedSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernancePriceFeedSet)
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
		it.Event = new(ContractAvsGovernancePriceFeedSet)
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
func (it *ContractAvsGovernancePriceFeedSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernancePriceFeedSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernancePriceFeedSet represents a PriceFeedSet event raised by the ContractAvsGovernance contract.
type ContractAvsGovernancePriceFeedSet struct {
	StakingContract common.Address
	Feed            common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPriceFeedSet is a free log retrieval operation binding the contract event 0xd2d8394cf7549a5ddbc2ba3dd7b2de8d53c891472d1f2907008ed6a10045fdae.
//
// Solidity: event PriceFeedSet(address indexed stakingContract, address feed)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterPriceFeedSet(opts *bind.FilterOpts, stakingContract []common.Address) (*ContractAvsGovernancePriceFeedSetIterator, error) {

	var stakingContractRule []interface{}
	for _, stakingContractItem := range stakingContract {
		stakingContractRule = append(stakingContractRule, stakingContractItem)
	}

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "PriceFeedSet", stakingContractRule)
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernancePriceFeedSetIterator{contract: _ContractAvsGovernance.contract, event: "PriceFeedSet", logs: logs, sub: sub}, nil
}

// WatchPriceFeedSet is a free log subscription operation binding the contract event 0xd2d8394cf7549a5ddbc2ba3dd7b2de8d53c891472d1f2907008ed6a10045fdae.
//
// Solidity: event PriceFeedSet(address indexed stakingContract, address feed)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchPriceFeedSet(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernancePriceFeedSet, stakingContract []common.Address) (event.Subscription, error) {

	var stakingContractRule []interface{}
	for _, stakingContractItem := range stakingContract {
		stakingContractRule = append(stakingContractRule, stakingContractItem)
	}

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "PriceFeedSet", stakingContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernancePriceFeedSet)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "PriceFeedSet", log); err != nil {
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

// ParsePriceFeedSet is a log parse operation binding the contract event 0xd2d8394cf7549a5ddbc2ba3dd7b2de8d53c891472d1f2907008ed6a10045fdae.
//
// Solidity: event PriceFeedSet(address indexed stakingContract, address feed)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParsePriceFeedSet(log types.Log) (*ContractAvsGovernancePriceFeedSet, error) {
	event := new(ContractAvsGovernancePriceFeedSet)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "PriceFeedSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceQueuedRewardsReceiverModificationIterator is returned from FilterQueuedRewardsReceiverModification and is used to iterate over the raw logs and unpacked data for QueuedRewardsReceiverModification events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceQueuedRewardsReceiverModificationIterator struct {
	Event *ContractAvsGovernanceQueuedRewardsReceiverModification // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceQueuedRewardsReceiverModificationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceQueuedRewardsReceiverModification)
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
		it.Event = new(ContractAvsGovernanceQueuedRewardsReceiverModification)
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
func (it *ContractAvsGovernanceQueuedRewardsReceiverModificationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceQueuedRewardsReceiverModificationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceQueuedRewardsReceiverModification represents a QueuedRewardsReceiverModification event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceQueuedRewardsReceiverModification struct {
	Operator common.Address
	Receiver common.Address
	Delay    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterQueuedRewardsReceiverModification is a free log retrieval operation binding the contract event 0x0d8cfa10a3087b28d3c226ad9a37314860e7c3c0505a25a39e3cdefb3118a98a.
//
// Solidity: event QueuedRewardsReceiverModification(address indexed operator, address receiver, uint256 delay)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterQueuedRewardsReceiverModification(opts *bind.FilterOpts, operator []common.Address) (*ContractAvsGovernanceQueuedRewardsReceiverModificationIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "QueuedRewardsReceiverModification", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceQueuedRewardsReceiverModificationIterator{contract: _ContractAvsGovernance.contract, event: "QueuedRewardsReceiverModification", logs: logs, sub: sub}, nil
}

// WatchQueuedRewardsReceiverModification is a free log subscription operation binding the contract event 0x0d8cfa10a3087b28d3c226ad9a37314860e7c3c0505a25a39e3cdefb3118a98a.
//
// Solidity: event QueuedRewardsReceiverModification(address indexed operator, address receiver, uint256 delay)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchQueuedRewardsReceiverModification(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceQueuedRewardsReceiverModification, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "QueuedRewardsReceiverModification", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceQueuedRewardsReceiverModification)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "QueuedRewardsReceiverModification", log); err != nil {
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

// ParseQueuedRewardsReceiverModification is a log parse operation binding the contract event 0x0d8cfa10a3087b28d3c226ad9a37314860e7c3c0505a25a39e3cdefb3118a98a.
//
// Solidity: event QueuedRewardsReceiverModification(address indexed operator, address receiver, uint256 delay)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseQueuedRewardsReceiverModification(log types.Log) (*ContractAvsGovernanceQueuedRewardsReceiverModification, error) {
	event := new(ContractAvsGovernanceQueuedRewardsReceiverModification)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "QueuedRewardsReceiverModification", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceRedistributionFailedIterator is returned from FilterRedistributionFailed and is used to iterate over the raw logs and unpacked data for RedistributionFailed events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceRedistributionFailedIterator struct {
	Event *ContractAvsGovernanceRedistributionFailed // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceRedistributionFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceRedistributionFailed)
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
		it.Event = new(ContractAvsGovernanceRedistributionFailed)
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
func (it *ContractAvsGovernanceRedistributionFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceRedistributionFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceRedistributionFailed represents a RedistributionFailed event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceRedistributionFailed struct {
	SlashId *big.Int
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRedistributionFailed is a free log retrieval operation binding the contract event 0x4283421d385d485c84e22eb466412b3f5fc0c64892d3a5b909c3395c23b766ba.
//
// Solidity: event RedistributionFailed(uint256 slashId, bytes data)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterRedistributionFailed(opts *bind.FilterOpts) (*ContractAvsGovernanceRedistributionFailedIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "RedistributionFailed")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceRedistributionFailedIterator{contract: _ContractAvsGovernance.contract, event: "RedistributionFailed", logs: logs, sub: sub}, nil
}

// WatchRedistributionFailed is a free log subscription operation binding the contract event 0x4283421d385d485c84e22eb466412b3f5fc0c64892d3a5b909c3395c23b766ba.
//
// Solidity: event RedistributionFailed(uint256 slashId, bytes data)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchRedistributionFailed(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceRedistributionFailed) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "RedistributionFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceRedistributionFailed)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "RedistributionFailed", log); err != nil {
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

// ParseRedistributionFailed is a log parse operation binding the contract event 0x4283421d385d485c84e22eb466412b3f5fc0c64892d3a5b909c3395c23b766ba.
//
// Solidity: event RedistributionFailed(uint256 slashId, bytes data)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseRedistributionFailed(log types.Log) (*ContractAvsGovernanceRedistributionFailed, error) {
	event := new(ContractAvsGovernanceRedistributionFailed)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "RedistributionFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceRoleAdminChangedIterator struct {
	Event *ContractAvsGovernanceRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceRoleAdminChanged)
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
		it.Event = new(ContractAvsGovernanceRoleAdminChanged)
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
func (it *ContractAvsGovernanceRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceRoleAdminChanged represents a RoleAdminChanged event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ContractAvsGovernanceRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceRoleAdminChangedIterator{contract: _ContractAvsGovernance.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceRoleAdminChanged)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseRoleAdminChanged(log types.Log) (*ContractAvsGovernanceRoleAdminChanged, error) {
	event := new(ContractAvsGovernanceRoleAdminChanged)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceRoleGrantedIterator struct {
	Event *ContractAvsGovernanceRoleGranted // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceRoleGranted)
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
		it.Event = new(ContractAvsGovernanceRoleGranted)
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
func (it *ContractAvsGovernanceRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceRoleGranted represents a RoleGranted event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ContractAvsGovernanceRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceRoleGrantedIterator{contract: _ContractAvsGovernance.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceRoleGranted)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseRoleGranted(log types.Log) (*ContractAvsGovernanceRoleGranted, error) {
	event := new(ContractAvsGovernanceRoleGranted)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceRoleRevokedIterator struct {
	Event *ContractAvsGovernanceRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceRoleRevoked)
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
		it.Event = new(ContractAvsGovernanceRoleRevoked)
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
func (it *ContractAvsGovernanceRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceRoleRevoked represents a RoleRevoked event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ContractAvsGovernanceRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceRoleRevokedIterator{contract: _ContractAvsGovernance.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceRoleRevoked)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseRoleRevoked(log types.Log) (*ContractAvsGovernanceRoleRevoked, error) {
	event := new(ContractAvsGovernanceRoleRevoked)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceSetAvsGovernanceLogicIterator is returned from FilterSetAvsGovernanceLogic and is used to iterate over the raw logs and unpacked data for SetAvsGovernanceLogic events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetAvsGovernanceLogicIterator struct {
	Event *ContractAvsGovernanceSetAvsGovernanceLogic // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceSetAvsGovernanceLogicIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceSetAvsGovernanceLogic)
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
		it.Event = new(ContractAvsGovernanceSetAvsGovernanceLogic)
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
func (it *ContractAvsGovernanceSetAvsGovernanceLogicIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceSetAvsGovernanceLogicIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceSetAvsGovernanceLogic represents a SetAvsGovernanceLogic event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetAvsGovernanceLogic struct {
	AvsGovernanceLogic common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSetAvsGovernanceLogic is a free log retrieval operation binding the contract event 0x7c36ee80df183e227956a9f387a48d26bbf4d2f1526410493d11126de5a8942c.
//
// Solidity: event SetAvsGovernanceLogic(address avsGovernanceLogic)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterSetAvsGovernanceLogic(opts *bind.FilterOpts) (*ContractAvsGovernanceSetAvsGovernanceLogicIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "SetAvsGovernanceLogic")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceSetAvsGovernanceLogicIterator{contract: _ContractAvsGovernance.contract, event: "SetAvsGovernanceLogic", logs: logs, sub: sub}, nil
}

// WatchSetAvsGovernanceLogic is a free log subscription operation binding the contract event 0x7c36ee80df183e227956a9f387a48d26bbf4d2f1526410493d11126de5a8942c.
//
// Solidity: event SetAvsGovernanceLogic(address avsGovernanceLogic)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchSetAvsGovernanceLogic(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetAvsGovernanceLogic) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "SetAvsGovernanceLogic")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceSetAvsGovernanceLogic)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetAvsGovernanceLogic", log); err != nil {
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

// ParseSetAvsGovernanceLogic is a log parse operation binding the contract event 0x7c36ee80df183e227956a9f387a48d26bbf4d2f1526410493d11126de5a8942c.
//
// Solidity: event SetAvsGovernanceLogic(address avsGovernanceLogic)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseSetAvsGovernanceLogic(log types.Log) (*ContractAvsGovernanceSetAvsGovernanceLogic, error) {
	event := new(ContractAvsGovernanceSetAvsGovernanceLogic)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetAvsGovernanceLogic", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceSetAvsNameIterator is returned from FilterSetAvsName and is used to iterate over the raw logs and unpacked data for SetAvsName events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetAvsNameIterator struct {
	Event *ContractAvsGovernanceSetAvsName // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceSetAvsNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceSetAvsName)
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
		it.Event = new(ContractAvsGovernanceSetAvsName)
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
func (it *ContractAvsGovernanceSetAvsNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceSetAvsNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceSetAvsName represents a SetAvsName event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetAvsName struct {
	AvsName string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetAvsName is a free log retrieval operation binding the contract event 0x7f63aacad63bc1693280450d5c3612ccd4efc53e46d69f3a537db102cd66290c.
//
// Solidity: event SetAvsName(string avsName)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterSetAvsName(opts *bind.FilterOpts) (*ContractAvsGovernanceSetAvsNameIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "SetAvsName")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceSetAvsNameIterator{contract: _ContractAvsGovernance.contract, event: "SetAvsName", logs: logs, sub: sub}, nil
}

// WatchSetAvsName is a free log subscription operation binding the contract event 0x7f63aacad63bc1693280450d5c3612ccd4efc53e46d69f3a537db102cd66290c.
//
// Solidity: event SetAvsName(string avsName)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchSetAvsName(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetAvsName) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "SetAvsName")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceSetAvsName)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetAvsName", log); err != nil {
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

// ParseSetAvsName is a log parse operation binding the contract event 0x7f63aacad63bc1693280450d5c3612ccd4efc53e46d69f3a537db102cd66290c.
//
// Solidity: event SetAvsName(string avsName)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseSetAvsName(log types.Log) (*ContractAvsGovernanceSetAvsName, error) {
	event := new(ContractAvsGovernanceSetAvsName)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetAvsName", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceSetIsAllowlistedIterator is returned from FilterSetIsAllowlisted and is used to iterate over the raw logs and unpacked data for SetIsAllowlisted events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetIsAllowlistedIterator struct {
	Event *ContractAvsGovernanceSetIsAllowlisted // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceSetIsAllowlistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceSetIsAllowlisted)
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
		it.Event = new(ContractAvsGovernanceSetIsAllowlisted)
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
func (it *ContractAvsGovernanceSetIsAllowlistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceSetIsAllowlistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceSetIsAllowlisted represents a SetIsAllowlisted event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetIsAllowlisted struct {
	IsAllowlisted bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetIsAllowlisted is a free log retrieval operation binding the contract event 0x2dcb3282f9b7aa18e1bf7fa254c45f3e270e8f26d9a37ae590d5d8125b58d1b1.
//
// Solidity: event SetIsAllowlisted(bool isAllowlisted)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterSetIsAllowlisted(opts *bind.FilterOpts) (*ContractAvsGovernanceSetIsAllowlistedIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "SetIsAllowlisted")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceSetIsAllowlistedIterator{contract: _ContractAvsGovernance.contract, event: "SetIsAllowlisted", logs: logs, sub: sub}, nil
}

// WatchSetIsAllowlisted is a free log subscription operation binding the contract event 0x2dcb3282f9b7aa18e1bf7fa254c45f3e270e8f26d9a37ae590d5d8125b58d1b1.
//
// Solidity: event SetIsAllowlisted(bool isAllowlisted)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchSetIsAllowlisted(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetIsAllowlisted) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "SetIsAllowlisted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceSetIsAllowlisted)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetIsAllowlisted", log); err != nil {
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

// ParseSetIsAllowlisted is a log parse operation binding the contract event 0x2dcb3282f9b7aa18e1bf7fa254c45f3e270e8f26d9a37ae590d5d8125b58d1b1.
//
// Solidity: event SetIsAllowlisted(bool isAllowlisted)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseSetIsAllowlisted(log types.Log) (*ContractAvsGovernanceSetIsAllowlisted, error) {
	event := new(ContractAvsGovernanceSetIsAllowlisted)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetIsAllowlisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceSetP2pAuthenticationEnabledIterator is returned from FilterSetP2pAuthenticationEnabled and is used to iterate over the raw logs and unpacked data for SetP2pAuthenticationEnabled events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetP2pAuthenticationEnabledIterator struct {
	Event *ContractAvsGovernanceSetP2pAuthenticationEnabled // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceSetP2pAuthenticationEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceSetP2pAuthenticationEnabled)
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
		it.Event = new(ContractAvsGovernanceSetP2pAuthenticationEnabled)
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
func (it *ContractAvsGovernanceSetP2pAuthenticationEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceSetP2pAuthenticationEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceSetP2pAuthenticationEnabled represents a SetP2pAuthenticationEnabled event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetP2pAuthenticationEnabled struct {
	IsEnabled bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetP2pAuthenticationEnabled is a free log retrieval operation binding the contract event 0x476bc7159a4e42477e0aa2e2e9709c55cb469b547976f91002b8db00bbd9b2f6.
//
// Solidity: event SetP2pAuthenticationEnabled(bool _isEnabled)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterSetP2pAuthenticationEnabled(opts *bind.FilterOpts) (*ContractAvsGovernanceSetP2pAuthenticationEnabledIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "SetP2pAuthenticationEnabled")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceSetP2pAuthenticationEnabledIterator{contract: _ContractAvsGovernance.contract, event: "SetP2pAuthenticationEnabled", logs: logs, sub: sub}, nil
}

// WatchSetP2pAuthenticationEnabled is a free log subscription operation binding the contract event 0x476bc7159a4e42477e0aa2e2e9709c55cb469b547976f91002b8db00bbd9b2f6.
//
// Solidity: event SetP2pAuthenticationEnabled(bool _isEnabled)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchSetP2pAuthenticationEnabled(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetP2pAuthenticationEnabled) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "SetP2pAuthenticationEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceSetP2pAuthenticationEnabled)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetP2pAuthenticationEnabled", log); err != nil {
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

// ParseSetP2pAuthenticationEnabled is a log parse operation binding the contract event 0x476bc7159a4e42477e0aa2e2e9709c55cb469b547976f91002b8db00bbd9b2f6.
//
// Solidity: event SetP2pAuthenticationEnabled(bool _isEnabled)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseSetP2pAuthenticationEnabled(log types.Log) (*ContractAvsGovernanceSetP2pAuthenticationEnabled, error) {
	event := new(ContractAvsGovernanceSetP2pAuthenticationEnabled)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetP2pAuthenticationEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceSetRewardsReceiverIterator is returned from FilterSetRewardsReceiver and is used to iterate over the raw logs and unpacked data for SetRewardsReceiver events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetRewardsReceiverIterator struct {
	Event *ContractAvsGovernanceSetRewardsReceiver // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceSetRewardsReceiverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceSetRewardsReceiver)
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
		it.Event = new(ContractAvsGovernanceSetRewardsReceiver)
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
func (it *ContractAvsGovernanceSetRewardsReceiverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceSetRewardsReceiverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceSetRewardsReceiver represents a SetRewardsReceiver event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetRewardsReceiver struct {
	Operator common.Address
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetRewardsReceiver is a free log retrieval operation binding the contract event 0xe906feea2ef60b474e22b4169bdd4de6906a84cd448cbcee99593526fe87082d.
//
// Solidity: event SetRewardsReceiver(address indexed operator, address receiver)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterSetRewardsReceiver(opts *bind.FilterOpts, operator []common.Address) (*ContractAvsGovernanceSetRewardsReceiverIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "SetRewardsReceiver", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceSetRewardsReceiverIterator{contract: _ContractAvsGovernance.contract, event: "SetRewardsReceiver", logs: logs, sub: sub}, nil
}

// WatchSetRewardsReceiver is a free log subscription operation binding the contract event 0xe906feea2ef60b474e22b4169bdd4de6906a84cd448cbcee99593526fe87082d.
//
// Solidity: event SetRewardsReceiver(address indexed operator, address receiver)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchSetRewardsReceiver(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetRewardsReceiver, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "SetRewardsReceiver", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceSetRewardsReceiver)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetRewardsReceiver", log); err != nil {
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

// ParseSetRewardsReceiver is a log parse operation binding the contract event 0xe906feea2ef60b474e22b4169bdd4de6906a84cd448cbcee99593526fe87082d.
//
// Solidity: event SetRewardsReceiver(address indexed operator, address receiver)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseSetRewardsReceiver(log types.Log) (*ContractAvsGovernanceSetRewardsReceiver, error) {
	event := new(ContractAvsGovernanceSetRewardsReceiver)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetRewardsReceiver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceSetStakingContractMultiplierIterator is returned from FilterSetStakingContractMultiplier and is used to iterate over the raw logs and unpacked data for SetStakingContractMultiplier events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetStakingContractMultiplierIterator struct {
	Event *ContractAvsGovernanceSetStakingContractMultiplier // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceSetStakingContractMultiplierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceSetStakingContractMultiplier)
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
		it.Event = new(ContractAvsGovernanceSetStakingContractMultiplier)
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
func (it *ContractAvsGovernanceSetStakingContractMultiplierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceSetStakingContractMultiplierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceSetStakingContractMultiplier represents a SetStakingContractMultiplier event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetStakingContractMultiplier struct {
	StakingContract           common.Address
	StakingContractMultiplier *big.Int
	SlashableStakeWeight      *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterSetStakingContractMultiplier is a free log retrieval operation binding the contract event 0x844cd4418680d8a9b7047e4a287bf70ed3de8cdde22a6f1745ded1985972d7ad.
//
// Solidity: event SetStakingContractMultiplier(address stakingContract, uint256 stakingContractMultiplier, uint256 slashableStakeWeight)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterSetStakingContractMultiplier(opts *bind.FilterOpts) (*ContractAvsGovernanceSetStakingContractMultiplierIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "SetStakingContractMultiplier")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceSetStakingContractMultiplierIterator{contract: _ContractAvsGovernance.contract, event: "SetStakingContractMultiplier", logs: logs, sub: sub}, nil
}

// WatchSetStakingContractMultiplier is a free log subscription operation binding the contract event 0x844cd4418680d8a9b7047e4a287bf70ed3de8cdde22a6f1745ded1985972d7ad.
//
// Solidity: event SetStakingContractMultiplier(address stakingContract, uint256 stakingContractMultiplier, uint256 slashableStakeWeight)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchSetStakingContractMultiplier(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetStakingContractMultiplier) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "SetStakingContractMultiplier")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceSetStakingContractMultiplier)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetStakingContractMultiplier", log); err != nil {
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

// ParseSetStakingContractMultiplier is a log parse operation binding the contract event 0x844cd4418680d8a9b7047e4a287bf70ed3de8cdde22a6f1745ded1985972d7ad.
//
// Solidity: event SetStakingContractMultiplier(address stakingContract, uint256 stakingContractMultiplier, uint256 slashableStakeWeight)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseSetStakingContractMultiplier(log types.Log) (*ContractAvsGovernanceSetStakingContractMultiplier, error) {
	event := new(ContractAvsGovernanceSetStakingContractMultiplier)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetStakingContractMultiplier", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceSetTokenIterator is returned from FilterSetToken and is used to iterate over the raw logs and unpacked data for SetToken events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetTokenIterator struct {
	Event *ContractAvsGovernanceSetToken // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceSetTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceSetToken)
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
		it.Event = new(ContractAvsGovernanceSetToken)
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
func (it *ContractAvsGovernanceSetTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceSetTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceSetToken represents a SetToken event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetToken struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetToken is a free log retrieval operation binding the contract event 0xefc1fd16ea80a922086ee4e995739d59b025c1bcea6d1f67855747480c83214b.
//
// Solidity: event SetToken(address token)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterSetToken(opts *bind.FilterOpts) (*ContractAvsGovernanceSetTokenIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "SetToken")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceSetTokenIterator{contract: _ContractAvsGovernance.contract, event: "SetToken", logs: logs, sub: sub}, nil
}

// WatchSetToken is a free log subscription operation binding the contract event 0xefc1fd16ea80a922086ee4e995739d59b025c1bcea6d1f67855747480c83214b.
//
// Solidity: event SetToken(address token)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchSetToken(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetToken) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "SetToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceSetToken)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetToken", log); err != nil {
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

// ParseSetToken is a log parse operation binding the contract event 0xefc1fd16ea80a922086ee4e995739d59b025c1bcea6d1f67855747480c83214b.
//
// Solidity: event SetToken(address token)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseSetToken(log types.Log) (*ContractAvsGovernanceSetToken, error) {
	event := new(ContractAvsGovernanceSetToken)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "SetToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceSlashedFundsRedistributedIterator is returned from FilterSlashedFundsRedistributed and is used to iterate over the raw logs and unpacked data for SlashedFundsRedistributed events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSlashedFundsRedistributedIterator struct {
	Event *ContractAvsGovernanceSlashedFundsRedistributed // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceSlashedFundsRedistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceSlashedFundsRedistributed)
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
		it.Event = new(ContractAvsGovernanceSlashedFundsRedistributed)
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
func (it *ContractAvsGovernanceSlashedFundsRedistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceSlashedFundsRedistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceSlashedFundsRedistributed represents a SlashedFundsRedistributed event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSlashedFundsRedistributed struct {
	SlashId      *big.Int
	SlashDetails IRedistributionManagerSlashDetails
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSlashedFundsRedistributed is a free log retrieval operation binding the contract event 0x377c736dc81d58fe7ce282712d1e86c28d818a97761ff330dab5bcef76d4665c.
//
// Solidity: event SlashedFundsRedistributed(uint256 slashId, (address,address[],uint256[],uint8) slashDetails)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterSlashedFundsRedistributed(opts *bind.FilterOpts) (*ContractAvsGovernanceSlashedFundsRedistributedIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "SlashedFundsRedistributed")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceSlashedFundsRedistributedIterator{contract: _ContractAvsGovernance.contract, event: "SlashedFundsRedistributed", logs: logs, sub: sub}, nil
}

// WatchSlashedFundsRedistributed is a free log subscription operation binding the contract event 0x377c736dc81d58fe7ce282712d1e86c28d818a97761ff330dab5bcef76d4665c.
//
// Solidity: event SlashedFundsRedistributed(uint256 slashId, (address,address[],uint256[],uint8) slashDetails)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchSlashedFundsRedistributed(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSlashedFundsRedistributed) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "SlashedFundsRedistributed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceSlashedFundsRedistributed)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "SlashedFundsRedistributed", log); err != nil {
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

// ParseSlashedFundsRedistributed is a log parse operation binding the contract event 0x377c736dc81d58fe7ce282712d1e86c28d818a97761ff330dab5bcef76d4665c.
//
// Solidity: event SlashedFundsRedistributed(uint256 slashId, (address,address[],uint256[],uint8) slashDetails)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseSlashedFundsRedistributed(log types.Log) (*ContractAvsGovernanceSlashedFundsRedistributed, error) {
	event := new(ContractAvsGovernanceSlashedFundsRedistributed)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "SlashedFundsRedistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceSlashingConfigNotActivatedIterator is returned from FilterSlashingConfigNotActivated and is used to iterate over the raw logs and unpacked data for SlashingConfigNotActivated events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSlashingConfigNotActivatedIterator struct {
	Event *ContractAvsGovernanceSlashingConfigNotActivated // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceSlashingConfigNotActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceSlashingConfigNotActivated)
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
		it.Event = new(ContractAvsGovernanceSlashingConfigNotActivated)
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
func (it *ContractAvsGovernanceSlashingConfigNotActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceSlashingConfigNotActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceSlashingConfigNotActivated represents a SlashingConfigNotActivated event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSlashingConfigNotActivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSlashingConfigNotActivated is a free log retrieval operation binding the contract event 0x296f51dbc0c5bb583ca31a4c06f2e70635fa00cfefb8ae2a8e3076f78932d765.
//
// Solidity: event SlashingConfigNotActivated()
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterSlashingConfigNotActivated(opts *bind.FilterOpts) (*ContractAvsGovernanceSlashingConfigNotActivatedIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "SlashingConfigNotActivated")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceSlashingConfigNotActivatedIterator{contract: _ContractAvsGovernance.contract, event: "SlashingConfigNotActivated", logs: logs, sub: sub}, nil
}

// WatchSlashingConfigNotActivated is a free log subscription operation binding the contract event 0x296f51dbc0c5bb583ca31a4c06f2e70635fa00cfefb8ae2a8e3076f78932d765.
//
// Solidity: event SlashingConfigNotActivated()
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchSlashingConfigNotActivated(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSlashingConfigNotActivated) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "SlashingConfigNotActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceSlashingConfigNotActivated)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "SlashingConfigNotActivated", log); err != nil {
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

// ParseSlashingConfigNotActivated is a log parse operation binding the contract event 0x296f51dbc0c5bb583ca31a4c06f2e70635fa00cfefb8ae2a8e3076f78932d765.
//
// Solidity: event SlashingConfigNotActivated()
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseSlashingConfigNotActivated(log types.Log) (*ContractAvsGovernanceSlashingConfigNotActivated, error) {
	event := new(ContractAvsGovernanceSlashingConfigNotActivated)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "SlashingConfigNotActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceSlashingConfigNotFoundIterator is returned from FilterSlashingConfigNotFound and is used to iterate over the raw logs and unpacked data for SlashingConfigNotFound events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSlashingConfigNotFoundIterator struct {
	Event *ContractAvsGovernanceSlashingConfigNotFound // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceSlashingConfigNotFoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceSlashingConfigNotFound)
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
		it.Event = new(ContractAvsGovernanceSlashingConfigNotFound)
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
func (it *ContractAvsGovernanceSlashingConfigNotFoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceSlashingConfigNotFoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceSlashingConfigNotFound represents a SlashingConfigNotFound event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSlashingConfigNotFound struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSlashingConfigNotFound is a free log retrieval operation binding the contract event 0x98d3bac35198400c534ef9af74b08e628bb1dc6d37292e1c188599208b3a4de5.
//
// Solidity: event SlashingConfigNotFound()
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterSlashingConfigNotFound(opts *bind.FilterOpts) (*ContractAvsGovernanceSlashingConfigNotFoundIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "SlashingConfigNotFound")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceSlashingConfigNotFoundIterator{contract: _ContractAvsGovernance.contract, event: "SlashingConfigNotFound", logs: logs, sub: sub}, nil
}

// WatchSlashingConfigNotFound is a free log subscription operation binding the contract event 0x98d3bac35198400c534ef9af74b08e628bb1dc6d37292e1c188599208b3a4de5.
//
// Solidity: event SlashingConfigNotFound()
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchSlashingConfigNotFound(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSlashingConfigNotFound) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "SlashingConfigNotFound")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceSlashingConfigNotFound)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "SlashingConfigNotFound", log); err != nil {
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

// ParseSlashingConfigNotFound is a log parse operation binding the contract event 0x98d3bac35198400c534ef9af74b08e628bb1dc6d37292e1c188599208b3a4de5.
//
// Solidity: event SlashingConfigNotFound()
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseSlashingConfigNotFound(log types.Log) (*ContractAvsGovernanceSlashingConfigNotFound, error) {
	event := new(ContractAvsGovernanceSlashingConfigNotFound)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "SlashingConfigNotFound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceSlashingFailedIterator is returned from FilterSlashingFailed and is used to iterate over the raw logs and unpacked data for SlashingFailed events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSlashingFailedIterator struct {
	Event *ContractAvsGovernanceSlashingFailed // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceSlashingFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceSlashingFailed)
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
		it.Event = new(ContractAvsGovernanceSlashingFailed)
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
func (it *ContractAvsGovernanceSlashingFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceSlashingFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceSlashingFailed represents a SlashingFailed event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSlashingFailed struct {
	Operator common.Address
	Data     []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSlashingFailed is a free log retrieval operation binding the contract event 0x17ea47b379a2daf87a43d9fe6e50400125dff97424a6a7686eca528970fffc0f.
//
// Solidity: event SlashingFailed(address operator, bytes data)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterSlashingFailed(opts *bind.FilterOpts) (*ContractAvsGovernanceSlashingFailedIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "SlashingFailed")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceSlashingFailedIterator{contract: _ContractAvsGovernance.contract, event: "SlashingFailed", logs: logs, sub: sub}, nil
}

// WatchSlashingFailed is a free log subscription operation binding the contract event 0x17ea47b379a2daf87a43d9fe6e50400125dff97424a6a7686eca528970fffc0f.
//
// Solidity: event SlashingFailed(address operator, bytes data)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchSlashingFailed(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSlashingFailed) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "SlashingFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceSlashingFailed)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "SlashingFailed", log); err != nil {
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

// ParseSlashingFailed is a log parse operation binding the contract event 0x17ea47b379a2daf87a43d9fe6e50400125dff97424a6a7686eca528970fffc0f.
//
// Solidity: event SlashingFailed(address operator, bytes data)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseSlashingFailed(log types.Log) (*ContractAvsGovernanceSlashingFailed, error) {
	event := new(ContractAvsGovernanceSlashingFailed)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "SlashingFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceSymbioticSlashingBypassedZeroAmountIterator is returned from FilterSymbioticSlashingBypassedZeroAmount and is used to iterate over the raw logs and unpacked data for SymbioticSlashingBypassedZeroAmount events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSymbioticSlashingBypassedZeroAmountIterator struct {
	Event *ContractAvsGovernanceSymbioticSlashingBypassedZeroAmount // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceSymbioticSlashingBypassedZeroAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceSymbioticSlashingBypassedZeroAmount)
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
		it.Event = new(ContractAvsGovernanceSymbioticSlashingBypassedZeroAmount)
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
func (it *ContractAvsGovernanceSymbioticSlashingBypassedZeroAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceSymbioticSlashingBypassedZeroAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceSymbioticSlashingBypassedZeroAmount represents a SymbioticSlashingBypassedZeroAmount event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSymbioticSlashingBypassedZeroAmount struct {
	Vault    common.Address
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSymbioticSlashingBypassedZeroAmount is a free log retrieval operation binding the contract event 0x5c5bd80f7b5442cf8fa1109fd5e73f37ab66780749ff0f7eaa9b5f137e6a0829.
//
// Solidity: event SymbioticSlashingBypassedZeroAmount(address indexed vault, address indexed operator)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterSymbioticSlashingBypassedZeroAmount(opts *bind.FilterOpts, vault []common.Address, operator []common.Address) (*ContractAvsGovernanceSymbioticSlashingBypassedZeroAmountIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "SymbioticSlashingBypassedZeroAmount", vaultRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceSymbioticSlashingBypassedZeroAmountIterator{contract: _ContractAvsGovernance.contract, event: "SymbioticSlashingBypassedZeroAmount", logs: logs, sub: sub}, nil
}

// WatchSymbioticSlashingBypassedZeroAmount is a free log subscription operation binding the contract event 0x5c5bd80f7b5442cf8fa1109fd5e73f37ab66780749ff0f7eaa9b5f137e6a0829.
//
// Solidity: event SymbioticSlashingBypassedZeroAmount(address indexed vault, address indexed operator)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchSymbioticSlashingBypassedZeroAmount(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSymbioticSlashingBypassedZeroAmount, vault []common.Address, operator []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "SymbioticSlashingBypassedZeroAmount", vaultRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceSymbioticSlashingBypassedZeroAmount)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "SymbioticSlashingBypassedZeroAmount", log); err != nil {
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

// ParseSymbioticSlashingBypassedZeroAmount is a log parse operation binding the contract event 0x5c5bd80f7b5442cf8fa1109fd5e73f37ab66780749ff0f7eaa9b5f137e6a0829.
//
// Solidity: event SymbioticSlashingBypassedZeroAmount(address indexed vault, address indexed operator)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseSymbioticSlashingBypassedZeroAmount(log types.Log) (*ContractAvsGovernanceSymbioticSlashingBypassedZeroAmount, error) {
	event := new(ContractAvsGovernanceSymbioticSlashingBypassedZeroAmount)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "SymbioticSlashingBypassedZeroAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceSymbioticSlashingExecutedIterator is returned from FilterSymbioticSlashingExecuted and is used to iterate over the raw logs and unpacked data for SymbioticSlashingExecuted events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSymbioticSlashingExecutedIterator struct {
	Event *ContractAvsGovernanceSymbioticSlashingExecuted // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceSymbioticSlashingExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceSymbioticSlashingExecuted)
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
		it.Event = new(ContractAvsGovernanceSymbioticSlashingExecuted)
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
func (it *ContractAvsGovernanceSymbioticSlashingExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceSymbioticSlashingExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceSymbioticSlashingExecuted represents a SymbioticSlashingExecuted event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSymbioticSlashingExecuted struct {
	Vault       common.Address
	Operator    common.Address
	SlashAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSymbioticSlashingExecuted is a free log retrieval operation binding the contract event 0x1ced844fcb8e184656bb631d13776c817857483d8cbd2fb7fa07d6c933a76a81.
//
// Solidity: event SymbioticSlashingExecuted(address indexed vault, address indexed operator, uint256 slashAmount)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterSymbioticSlashingExecuted(opts *bind.FilterOpts, vault []common.Address, operator []common.Address) (*ContractAvsGovernanceSymbioticSlashingExecutedIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "SymbioticSlashingExecuted", vaultRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceSymbioticSlashingExecutedIterator{contract: _ContractAvsGovernance.contract, event: "SymbioticSlashingExecuted", logs: logs, sub: sub}, nil
}

// WatchSymbioticSlashingExecuted is a free log subscription operation binding the contract event 0x1ced844fcb8e184656bb631d13776c817857483d8cbd2fb7fa07d6c933a76a81.
//
// Solidity: event SymbioticSlashingExecuted(address indexed vault, address indexed operator, uint256 slashAmount)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchSymbioticSlashingExecuted(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSymbioticSlashingExecuted, vault []common.Address, operator []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "SymbioticSlashingExecuted", vaultRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceSymbioticSlashingExecuted)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "SymbioticSlashingExecuted", log); err != nil {
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

// ParseSymbioticSlashingExecuted is a log parse operation binding the contract event 0x1ced844fcb8e184656bb631d13776c817857483d8cbd2fb7fa07d6c933a76a81.
//
// Solidity: event SymbioticSlashingExecuted(address indexed vault, address indexed operator, uint256 slashAmount)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseSymbioticSlashingExecuted(log types.Log) (*ContractAvsGovernanceSymbioticSlashingExecuted, error) {
	event := new(ContractAvsGovernanceSymbioticSlashingExecuted)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "SymbioticSlashingExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceSymbioticSlashingRevertedIterator is returned from FilterSymbioticSlashingReverted and is used to iterate over the raw logs and unpacked data for SymbioticSlashingReverted events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSymbioticSlashingRevertedIterator struct {
	Event *ContractAvsGovernanceSymbioticSlashingReverted // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceSymbioticSlashingRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceSymbioticSlashingReverted)
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
		it.Event = new(ContractAvsGovernanceSymbioticSlashingReverted)
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
func (it *ContractAvsGovernanceSymbioticSlashingRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceSymbioticSlashingRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceSymbioticSlashingReverted represents a SymbioticSlashingReverted event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSymbioticSlashingReverted struct {
	Vault       common.Address
	Operator    common.Address
	SlashAmount *big.Int
	RevertData  []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSymbioticSlashingReverted is a free log retrieval operation binding the contract event 0x2f745c06290fbdf2aeccacf7710029e14576b679b64a3391ee3572311987cf88.
//
// Solidity: event SymbioticSlashingReverted(address indexed vault, address indexed operator, uint256 slashAmount, bytes revertData)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterSymbioticSlashingReverted(opts *bind.FilterOpts, vault []common.Address, operator []common.Address) (*ContractAvsGovernanceSymbioticSlashingRevertedIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "SymbioticSlashingReverted", vaultRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceSymbioticSlashingRevertedIterator{contract: _ContractAvsGovernance.contract, event: "SymbioticSlashingReverted", logs: logs, sub: sub}, nil
}

// WatchSymbioticSlashingReverted is a free log subscription operation binding the contract event 0x2f745c06290fbdf2aeccacf7710029e14576b679b64a3391ee3572311987cf88.
//
// Solidity: event SymbioticSlashingReverted(address indexed vault, address indexed operator, uint256 slashAmount, bytes revertData)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchSymbioticSlashingReverted(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSymbioticSlashingReverted, vault []common.Address, operator []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "SymbioticSlashingReverted", vaultRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceSymbioticSlashingReverted)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "SymbioticSlashingReverted", log); err != nil {
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

// ParseSymbioticSlashingReverted is a log parse operation binding the contract event 0x2f745c06290fbdf2aeccacf7710029e14576b679b64a3391ee3572311987cf88.
//
// Solidity: event SymbioticSlashingReverted(address indexed vault, address indexed operator, uint256 slashAmount, bytes revertData)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseSymbioticSlashingReverted(log types.Log) (*ContractAvsGovernanceSymbioticSlashingReverted, error) {
	event := new(ContractAvsGovernanceSymbioticSlashingReverted)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "SymbioticSlashingReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceSymbioticSlashingSkippedIterator is returned from FilterSymbioticSlashingSkipped and is used to iterate over the raw logs and unpacked data for SymbioticSlashingSkipped events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSymbioticSlashingSkippedIterator struct {
	Event *ContractAvsGovernanceSymbioticSlashingSkipped // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceSymbioticSlashingSkippedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceSymbioticSlashingSkipped)
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
		it.Event = new(ContractAvsGovernanceSymbioticSlashingSkipped)
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
func (it *ContractAvsGovernanceSymbioticSlashingSkippedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceSymbioticSlashingSkippedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceSymbioticSlashingSkipped represents a SymbioticSlashingSkipped event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSymbioticSlashingSkipped struct {
	Vault    common.Address
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSymbioticSlashingSkipped is a free log retrieval operation binding the contract event 0x70201f2c2fc9df79327ee453894afca19249c6be2c7ee4dce5f7be735d981027.
//
// Solidity: event SymbioticSlashingSkipped(address indexed vault, address indexed operator)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterSymbioticSlashingSkipped(opts *bind.FilterOpts, vault []common.Address, operator []common.Address) (*ContractAvsGovernanceSymbioticSlashingSkippedIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "SymbioticSlashingSkipped", vaultRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceSymbioticSlashingSkippedIterator{contract: _ContractAvsGovernance.contract, event: "SymbioticSlashingSkipped", logs: logs, sub: sub}, nil
}

// WatchSymbioticSlashingSkipped is a free log subscription operation binding the contract event 0x70201f2c2fc9df79327ee453894afca19249c6be2c7ee4dce5f7be735d981027.
//
// Solidity: event SymbioticSlashingSkipped(address indexed vault, address indexed operator)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchSymbioticSlashingSkipped(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSymbioticSlashingSkipped, vault []common.Address, operator []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "SymbioticSlashingSkipped", vaultRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceSymbioticSlashingSkipped)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "SymbioticSlashingSkipped", log); err != nil {
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

// ParseSymbioticSlashingSkipped is a log parse operation binding the contract event 0x70201f2c2fc9df79327ee453894afca19249c6be2c7ee4dce5f7be735d981027.
//
// Solidity: event SymbioticSlashingSkipped(address indexed vault, address indexed operator)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseSymbioticSlashingSkipped(log types.Log) (*ContractAvsGovernanceSymbioticSlashingSkipped, error) {
	event := new(ContractAvsGovernanceSymbioticSlashingSkipped)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "SymbioticSlashingSkipped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceVetoSlashAlreadyCompletedIterator is returned from FilterVetoSlashAlreadyCompleted and is used to iterate over the raw logs and unpacked data for VetoSlashAlreadyCompleted events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceVetoSlashAlreadyCompletedIterator struct {
	Event *ContractAvsGovernanceVetoSlashAlreadyCompleted // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceVetoSlashAlreadyCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceVetoSlashAlreadyCompleted)
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
		it.Event = new(ContractAvsGovernanceVetoSlashAlreadyCompleted)
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
func (it *ContractAvsGovernanceVetoSlashAlreadyCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceVetoSlashAlreadyCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceVetoSlashAlreadyCompleted represents a VetoSlashAlreadyCompleted event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceVetoSlashAlreadyCompleted struct {
	Slasher    common.Address
	SlashIndex *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVetoSlashAlreadyCompleted is a free log retrieval operation binding the contract event 0x41f5d5d07d79ca6c17b25e703ceca8de971893b69031d4ca0b033de1c6bb6b01.
//
// Solidity: event VetoSlashAlreadyCompleted(address indexed slasher, uint256 indexed slashIndex)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterVetoSlashAlreadyCompleted(opts *bind.FilterOpts, slasher []common.Address, slashIndex []*big.Int) (*ContractAvsGovernanceVetoSlashAlreadyCompletedIterator, error) {

	var slasherRule []interface{}
	for _, slasherItem := range slasher {
		slasherRule = append(slasherRule, slasherItem)
	}
	var slashIndexRule []interface{}
	for _, slashIndexItem := range slashIndex {
		slashIndexRule = append(slashIndexRule, slashIndexItem)
	}

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "VetoSlashAlreadyCompleted", slasherRule, slashIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceVetoSlashAlreadyCompletedIterator{contract: _ContractAvsGovernance.contract, event: "VetoSlashAlreadyCompleted", logs: logs, sub: sub}, nil
}

// WatchVetoSlashAlreadyCompleted is a free log subscription operation binding the contract event 0x41f5d5d07d79ca6c17b25e703ceca8de971893b69031d4ca0b033de1c6bb6b01.
//
// Solidity: event VetoSlashAlreadyCompleted(address indexed slasher, uint256 indexed slashIndex)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchVetoSlashAlreadyCompleted(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceVetoSlashAlreadyCompleted, slasher []common.Address, slashIndex []*big.Int) (event.Subscription, error) {

	var slasherRule []interface{}
	for _, slasherItem := range slasher {
		slasherRule = append(slasherRule, slasherItem)
	}
	var slashIndexRule []interface{}
	for _, slashIndexItem := range slashIndex {
		slashIndexRule = append(slashIndexRule, slashIndexItem)
	}

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "VetoSlashAlreadyCompleted", slasherRule, slashIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceVetoSlashAlreadyCompleted)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "VetoSlashAlreadyCompleted", log); err != nil {
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

// ParseVetoSlashAlreadyCompleted is a log parse operation binding the contract event 0x41f5d5d07d79ca6c17b25e703ceca8de971893b69031d4ca0b033de1c6bb6b01.
//
// Solidity: event VetoSlashAlreadyCompleted(address indexed slasher, uint256 indexed slashIndex)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseVetoSlashAlreadyCompleted(log types.Log) (*ContractAvsGovernanceVetoSlashAlreadyCompleted, error) {
	event := new(ContractAvsGovernanceVetoSlashAlreadyCompleted)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "VetoSlashAlreadyCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceVetoSlashExecutedIterator is returned from FilterVetoSlashExecuted and is used to iterate over the raw logs and unpacked data for VetoSlashExecuted events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceVetoSlashExecutedIterator struct {
	Event *ContractAvsGovernanceVetoSlashExecuted // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceVetoSlashExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceVetoSlashExecuted)
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
		it.Event = new(ContractAvsGovernanceVetoSlashExecuted)
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
func (it *ContractAvsGovernanceVetoSlashExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceVetoSlashExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceVetoSlashExecuted represents a VetoSlashExecuted event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceVetoSlashExecuted struct {
	Slasher       common.Address
	SlashIndex    *big.Int
	SlashedAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterVetoSlashExecuted is a free log retrieval operation binding the contract event 0x4f50a2e4f91995f76aef55cac74ffac2c72c4de8b344e5209d50e5a655a4ad27.
//
// Solidity: event VetoSlashExecuted(address indexed slasher, uint256 indexed slashIndex, uint256 slashedAmount)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterVetoSlashExecuted(opts *bind.FilterOpts, slasher []common.Address, slashIndex []*big.Int) (*ContractAvsGovernanceVetoSlashExecutedIterator, error) {

	var slasherRule []interface{}
	for _, slasherItem := range slasher {
		slasherRule = append(slasherRule, slasherItem)
	}
	var slashIndexRule []interface{}
	for _, slashIndexItem := range slashIndex {
		slashIndexRule = append(slashIndexRule, slashIndexItem)
	}

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "VetoSlashExecuted", slasherRule, slashIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceVetoSlashExecutedIterator{contract: _ContractAvsGovernance.contract, event: "VetoSlashExecuted", logs: logs, sub: sub}, nil
}

// WatchVetoSlashExecuted is a free log subscription operation binding the contract event 0x4f50a2e4f91995f76aef55cac74ffac2c72c4de8b344e5209d50e5a655a4ad27.
//
// Solidity: event VetoSlashExecuted(address indexed slasher, uint256 indexed slashIndex, uint256 slashedAmount)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchVetoSlashExecuted(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceVetoSlashExecuted, slasher []common.Address, slashIndex []*big.Int) (event.Subscription, error) {

	var slasherRule []interface{}
	for _, slasherItem := range slasher {
		slasherRule = append(slasherRule, slasherItem)
	}
	var slashIndexRule []interface{}
	for _, slashIndexItem := range slashIndex {
		slashIndexRule = append(slashIndexRule, slashIndexItem)
	}

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "VetoSlashExecuted", slasherRule, slashIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceVetoSlashExecuted)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "VetoSlashExecuted", log); err != nil {
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

// ParseVetoSlashExecuted is a log parse operation binding the contract event 0x4f50a2e4f91995f76aef55cac74ffac2c72c4de8b344e5209d50e5a655a4ad27.
//
// Solidity: event VetoSlashExecuted(address indexed slasher, uint256 indexed slashIndex, uint256 slashedAmount)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseVetoSlashExecuted(log types.Log) (*ContractAvsGovernanceVetoSlashExecuted, error) {
	event := new(ContractAvsGovernanceVetoSlashExecuted)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "VetoSlashExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceVetoSlashExecutionFailedIterator is returned from FilterVetoSlashExecutionFailed and is used to iterate over the raw logs and unpacked data for VetoSlashExecutionFailed events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceVetoSlashExecutionFailedIterator struct {
	Event *ContractAvsGovernanceVetoSlashExecutionFailed // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceVetoSlashExecutionFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceVetoSlashExecutionFailed)
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
		it.Event = new(ContractAvsGovernanceVetoSlashExecutionFailed)
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
func (it *ContractAvsGovernanceVetoSlashExecutionFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceVetoSlashExecutionFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceVetoSlashExecutionFailed represents a VetoSlashExecutionFailed event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceVetoSlashExecutionFailed struct {
	Slasher    common.Address
	SlashIndex *big.Int
	RevertData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVetoSlashExecutionFailed is a free log retrieval operation binding the contract event 0x047efda7bcf14d7e1b38d333153ef235e3f39987459657e0d082b903006f6d8a.
//
// Solidity: event VetoSlashExecutionFailed(address indexed slasher, uint256 indexed slashIndex, bytes revertData)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterVetoSlashExecutionFailed(opts *bind.FilterOpts, slasher []common.Address, slashIndex []*big.Int) (*ContractAvsGovernanceVetoSlashExecutionFailedIterator, error) {

	var slasherRule []interface{}
	for _, slasherItem := range slasher {
		slasherRule = append(slasherRule, slasherItem)
	}
	var slashIndexRule []interface{}
	for _, slashIndexItem := range slashIndex {
		slashIndexRule = append(slashIndexRule, slashIndexItem)
	}

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "VetoSlashExecutionFailed", slasherRule, slashIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceVetoSlashExecutionFailedIterator{contract: _ContractAvsGovernance.contract, event: "VetoSlashExecutionFailed", logs: logs, sub: sub}, nil
}

// WatchVetoSlashExecutionFailed is a free log subscription operation binding the contract event 0x047efda7bcf14d7e1b38d333153ef235e3f39987459657e0d082b903006f6d8a.
//
// Solidity: event VetoSlashExecutionFailed(address indexed slasher, uint256 indexed slashIndex, bytes revertData)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchVetoSlashExecutionFailed(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceVetoSlashExecutionFailed, slasher []common.Address, slashIndex []*big.Int) (event.Subscription, error) {

	var slasherRule []interface{}
	for _, slasherItem := range slasher {
		slasherRule = append(slasherRule, slasherItem)
	}
	var slashIndexRule []interface{}
	for _, slashIndexItem := range slashIndex {
		slashIndexRule = append(slashIndexRule, slashIndexItem)
	}

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "VetoSlashExecutionFailed", slasherRule, slashIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceVetoSlashExecutionFailed)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "VetoSlashExecutionFailed", log); err != nil {
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

// ParseVetoSlashExecutionFailed is a log parse operation binding the contract event 0x047efda7bcf14d7e1b38d333153ef235e3f39987459657e0d082b903006f6d8a.
//
// Solidity: event VetoSlashExecutionFailed(address indexed slasher, uint256 indexed slashIndex, bytes revertData)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseVetoSlashExecutionFailed(log types.Log) (*ContractAvsGovernanceVetoSlashExecutionFailed, error) {
	event := new(ContractAvsGovernanceVetoSlashExecutionFailed)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "VetoSlashExecutionFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceVetoSlashRequestedIterator is returned from FilterVetoSlashRequested and is used to iterate over the raw logs and unpacked data for VetoSlashRequested events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceVetoSlashRequestedIterator struct {
	Event *ContractAvsGovernanceVetoSlashRequested // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceVetoSlashRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceVetoSlashRequested)
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
		it.Event = new(ContractAvsGovernanceVetoSlashRequested)
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
func (it *ContractAvsGovernanceVetoSlashRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceVetoSlashRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceVetoSlashRequested represents a VetoSlashRequested event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceVetoSlashRequested struct {
	Slasher    common.Address
	SlashIndex *big.Int
	Operator   common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVetoSlashRequested is a free log retrieval operation binding the contract event 0x431852c863b55b9e746d12bb87e301e9119145835594f62eaa2aaceab5938cbc.
//
// Solidity: event VetoSlashRequested(address indexed slasher, uint256 indexed slashIndex, address indexed operator)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterVetoSlashRequested(opts *bind.FilterOpts, slasher []common.Address, slashIndex []*big.Int, operator []common.Address) (*ContractAvsGovernanceVetoSlashRequestedIterator, error) {

	var slasherRule []interface{}
	for _, slasherItem := range slasher {
		slasherRule = append(slasherRule, slasherItem)
	}
	var slashIndexRule []interface{}
	for _, slashIndexItem := range slashIndex {
		slashIndexRule = append(slashIndexRule, slashIndexItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "VetoSlashRequested", slasherRule, slashIndexRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceVetoSlashRequestedIterator{contract: _ContractAvsGovernance.contract, event: "VetoSlashRequested", logs: logs, sub: sub}, nil
}

// WatchVetoSlashRequested is a free log subscription operation binding the contract event 0x431852c863b55b9e746d12bb87e301e9119145835594f62eaa2aaceab5938cbc.
//
// Solidity: event VetoSlashRequested(address indexed slasher, uint256 indexed slashIndex, address indexed operator)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchVetoSlashRequested(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceVetoSlashRequested, slasher []common.Address, slashIndex []*big.Int, operator []common.Address) (event.Subscription, error) {

	var slasherRule []interface{}
	for _, slasherItem := range slasher {
		slasherRule = append(slasherRule, slasherItem)
	}
	var slashIndexRule []interface{}
	for _, slashIndexItem := range slashIndex {
		slashIndexRule = append(slashIndexRule, slashIndexItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "VetoSlashRequested", slasherRule, slashIndexRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceVetoSlashRequested)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "VetoSlashRequested", log); err != nil {
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

// ParseVetoSlashRequested is a log parse operation binding the contract event 0x431852c863b55b9e746d12bb87e301e9119145835594f62eaa2aaceab5938cbc.
//
// Solidity: event VetoSlashRequested(address indexed slasher, uint256 indexed slashIndex, address indexed operator)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseVetoSlashRequested(log types.Log) (*ContractAvsGovernanceVetoSlashRequested, error) {
	event := new(ContractAvsGovernanceVetoSlashRequested)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "VetoSlashRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAvsGovernanceSetNewSupportedStakingContractsIterator is returned from FilterSetNewSupportedStakingContracts and is used to iterate over the raw logs and unpacked data for SetNewSupportedStakingContracts events raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetNewSupportedStakingContractsIterator struct {
	Event *ContractAvsGovernanceSetNewSupportedStakingContracts // Event containing the contract specifics and raw log

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
func (it *ContractAvsGovernanceSetNewSupportedStakingContractsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAvsGovernanceSetNewSupportedStakingContracts)
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
		it.Event = new(ContractAvsGovernanceSetNewSupportedStakingContracts)
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
func (it *ContractAvsGovernanceSetNewSupportedStakingContractsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAvsGovernanceSetNewSupportedStakingContractsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAvsGovernanceSetNewSupportedStakingContracts represents a SetNewSupportedStakingContracts event raised by the ContractAvsGovernance contract.
type ContractAvsGovernanceSetNewSupportedStakingContracts struct {
	StakingContracts []common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetNewSupportedStakingContracts is a free log retrieval operation binding the contract event 0xdc9dd92c894ed28df4cdf880b63e1f5d6a6a32870a392e08c13ddfe18488caea.
//
// Solidity: event setNewSupportedStakingContracts(address[] stakingContracts)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) FilterSetNewSupportedStakingContracts(opts *bind.FilterOpts) (*ContractAvsGovernanceSetNewSupportedStakingContractsIterator, error) {

	logs, sub, err := _ContractAvsGovernance.contract.FilterLogs(opts, "setNewSupportedStakingContracts")
	if err != nil {
		return nil, err
	}
	return &ContractAvsGovernanceSetNewSupportedStakingContractsIterator{contract: _ContractAvsGovernance.contract, event: "setNewSupportedStakingContracts", logs: logs, sub: sub}, nil
}

// WatchSetNewSupportedStakingContracts is a free log subscription operation binding the contract event 0xdc9dd92c894ed28df4cdf880b63e1f5d6a6a32870a392e08c13ddfe18488caea.
//
// Solidity: event setNewSupportedStakingContracts(address[] stakingContracts)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) WatchSetNewSupportedStakingContracts(opts *bind.WatchOpts, sink chan<- *ContractAvsGovernanceSetNewSupportedStakingContracts) (event.Subscription, error) {

	logs, sub, err := _ContractAvsGovernance.contract.WatchLogs(opts, "setNewSupportedStakingContracts")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAvsGovernanceSetNewSupportedStakingContracts)
				if err := _ContractAvsGovernance.contract.UnpackLog(event, "setNewSupportedStakingContracts", log); err != nil {
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

// ParseSetNewSupportedStakingContracts is a log parse operation binding the contract event 0xdc9dd92c894ed28df4cdf880b63e1f5d6a6a32870a392e08c13ddfe18488caea.
//
// Solidity: event setNewSupportedStakingContracts(address[] stakingContracts)
func (_ContractAvsGovernance *ContractAvsGovernanceFilterer) ParseSetNewSupportedStakingContracts(log types.Log) (*ContractAvsGovernanceSetNewSupportedStakingContracts, error) {
	event := new(ContractAvsGovernanceSetNewSupportedStakingContracts)
	if err := _ContractAvsGovernance.contract.UnpackLog(event, "setNewSupportedStakingContracts", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
