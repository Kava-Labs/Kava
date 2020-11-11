package harvest

// DO NOT EDIT - generated by aliasgen tool (github.com/rhuairahrighairidh/aliasgen)

import (
	"github.com/kava-labs/kava/x/harvest/keeper"
	"github.com/kava-labs/kava/x/harvest/types"
)

const (
	BeginningOfMonth                      = keeper.BeginningOfMonth
	MidMonth                              = keeper.MidMonth
	PaymentHour                           = keeper.PaymentHour
	AttributeKeyBlockHeight               = types.AttributeKeyBlockHeight
	AttributeKeyClaimAmount               = types.AttributeKeyClaimAmount
	AttributeKeyClaimHolder               = types.AttributeKeyClaimHolder
	AttributeKeyClaimMultiplier           = types.AttributeKeyClaimMultiplier
	AttributeKeyClaimType                 = types.AttributeKeyClaimType
	AttributeKeyDeposit                   = types.AttributeKeyDeposit
	AttributeKeyDepositDenom              = types.AttributeKeyDepositDenom
	AttributeKeyDepositor                 = types.AttributeKeyDepositor
	AttributeKeyRewardsDistribution       = types.AttributeKeyRewardsDistribution
	AttributeValueCategory                = types.AttributeValueCategory
	DefaultParamspace                     = types.DefaultParamspace
	DelegatorAccount                      = types.DelegatorAccount
	EventTypeClaimHarvestReward           = types.EventTypeClaimHarvestReward
	EventTypeDeleteHarvestDeposit         = types.EventTypeDeleteHarvestDeposit
	EventTypeHarvestDelegatorDistribution = types.EventTypeHarvestDelegatorDistribution
	EventTypeHarvestDeposit               = types.EventTypeHarvestDeposit
	EventTypeHarvestLPDistribution        = types.EventTypeHarvestLPDistribution
	EventTypeHarvestWithdrawal            = types.EventTypeHarvestWithdrawal
	LP                                    = types.LP
	LPAccount                             = types.LPAccount
	Large                                 = types.Large
	Medium                                = types.Medium
	ModuleAccountName                     = types.ModuleAccountName
	ModuleName                            = types.ModuleName
	QuerierRoute                          = types.QuerierRoute
	QueryGetClaims                        = types.QueryGetClaims
	QueryGetDeposits                      = types.QueryGetDeposits
	QueryGetModuleAccounts                = types.QueryGetModuleAccounts
	QueryGetParams                        = types.QueryGetParams
	RouterKey                             = types.RouterKey
	Small                                 = types.Small
	Stake                                 = types.Stake
	StoreKey                              = types.StoreKey
)

var (
	// function aliases
	NewKeeper                        = keeper.NewKeeper
	NewQuerier                       = keeper.NewQuerier
	ClaimKey                         = types.ClaimKey
	DefaultGenesisState              = types.DefaultGenesisState
	DefaultParams                    = types.DefaultParams
	DepositKey                       = types.DepositKey
	DepositTypeIteratorKey           = types.DepositTypeIteratorKey
	GetTotalVestingPeriodLength      = types.GetTotalVestingPeriodLength
	NewClaim                         = types.NewClaim
	NewDelegatorDistributionSchedule = types.NewDelegatorDistributionSchedule
	NewDeposit                       = types.NewDeposit
	NewDistributionSchedule          = types.NewDistributionSchedule
	NewGenesisState                  = types.NewGenesisState
	NewMsgClaimReward                = types.NewMsgClaimReward
	NewMsgDeposit                    = types.NewMsgDeposit
	NewMsgWithdraw                   = types.NewMsgWithdraw
	NewMultiplier                    = types.NewMultiplier
	NewParams                        = types.NewParams
	NewPeriod                        = types.NewPeriod
	NewQueryAccountParams            = types.NewQueryAccountParams
	NewQueryClaimParams              = types.NewQueryClaimParams
	NewQueryDepositParams            = types.NewQueryDepositParams
	ParamKeyTable                    = types.ParamKeyTable
	RegisterCodec                    = types.RegisterCodec

	// variable aliases
	BorrowsKeyPrefix                  = types.BorrowsKeyPrefix
	ClaimsKeyPrefix                   = types.ClaimsKeyPrefix
	DefaultActive                     = types.DefaultActive
	DefaultDelegatorSchedules         = types.DefaultDelegatorSchedules
	DefaultDistributionTimes          = types.DefaultDistributionTimes
	DefaultGovSchedules               = types.DefaultGovSchedules
	DefaultLPSchedules                = types.DefaultLPSchedules
	DefaultPreviousBlockTime          = types.DefaultPreviousBlockTime
	ClaimTypesClaimQuery              = types.ClaimTypesClaimQuery
	DepositsKeyPrefix                 = types.DepositsKeyPrefix
	ErrAccountNotFound                = types.ErrAccountNotFound
	ErrClaimExpired                   = types.ErrClaimExpired
	ErrClaimNotFound                  = types.ErrClaimNotFound
	ErrDepositNotFound                = types.ErrDepositNotFound
	ErrGovScheduleNotFound            = types.ErrGovScheduleNotFound
	ErrInsufficientModAccountBalance  = types.ErrInsufficientModAccountBalance
	ErrInvaliWithdrawAmount           = types.ErrInvalidWithdrawAmount
	ErrInvalidAccountType             = types.ErrInvalidAccountType
	ErrInvalidDepositDenom            = types.ErrInvalidDepositDenom
	ErrInvalidClaimType               = types.ErrInvalidClaimType
	ErrInvalidMultiplier              = types.ErrInvalidMultiplier
	ErrLPScheduleNotFound             = types.ErrLPScheduleNotFound
	ErrZeroClaim                      = types.ErrZeroClaim
	GovDenom                          = types.GovDenom
	KeyActive                         = types.KeyActive
	KeyDelegatorSchedule              = types.KeyDelegatorSchedule
	KeyLPSchedules                    = types.KeyLPSchedules
	ModuleCdc                         = types.ModuleCdc
	PreviousBlockTimeKey              = types.PreviousBlockTimeKey
	PreviousDelegationDistributionKey = types.PreviousDelegationDistributionKey
)

type (
	Keeper                         = keeper.Keeper
	AccountKeeper                  = types.AccountKeeper
	Borrow                         = types.Borrow
	MoneyMarket                    = types.MoneyMarket
	MoneyMarkets                   = types.MoneyMarkets
	DelegatorDistributionSchedule  = types.DelegatorDistributionSchedule
	DelegatorDistributionSchedules = types.DelegatorDistributionSchedules
	Deposit                        = types.Deposit
	ClaimType                      = types.ClaimType
	DistributionSchedule           = types.DistributionSchedule
	DistributionSchedules          = types.DistributionSchedules
	GenesisDistributionTime        = types.GenesisDistributionTime
	GenesisDistributionTimes       = types.GenesisDistributionTimes
	GenesisState                   = types.GenesisState
	MsgClaimReward                 = types.MsgClaimReward
	MsgDeposit                     = types.MsgDeposit
	MsgWithdraw                    = types.MsgWithdraw
	Multiplier                     = types.Multiplier
	MultiplierName                 = types.MultiplierName
	Multipliers                    = types.Multipliers
	Params                         = types.Params
	QueryAccountParams             = types.QueryAccountParams
	QueryClaimParams               = types.QueryClaimParams
	QueryDepositParams             = types.QueryDepositParams
	StakingKeeper                  = types.StakingKeeper
	SupplyKeeper                   = types.SupplyKeeper
)
