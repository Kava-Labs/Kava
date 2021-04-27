package committee

// DO NOT EDIT - generated by aliasgen tool (github.com/rhuairahrighairidh/aliasgen)

import (
	"github.com/kava-labs/kava/x/committee/client"
	"github.com/kava-labs/kava/x/committee/keeper"
	"github.com/kava-labs/kava/x/committee/types"
)

const (
	AttributeKeyCommitteeID         = types.AttributeKeyCommitteeID
	AttributeKeyProposalCloseStatus = types.AttributeKeyProposalCloseStatus
	AttributeKeyProposalID          = types.AttributeKeyProposalID
	AttributeKeyVoter               = types.AttributeKeyVoter
	AttributeValueCategory          = types.AttributeValueCategory
	DefaultNextProposalID           = types.DefaultNextProposalID
	DefaultParamspace               = types.DefaultParamspace
	EventTypeProposalClose          = types.EventTypeProposalClose
	EventTypeProposalSubmit         = types.EventTypeProposalSubmit
	EventTypeProposalVote           = types.EventTypeProposalVote
	MaxCommitteeDescriptionLength   = types.MaxCommitteeDescriptionLength
	ModuleName                      = types.ModuleName
	No                              = types.No
	ProposalTypeCommitteeChange     = types.ProposalTypeCommitteeChange
	ProposalTypeCommitteeDelete     = types.ProposalTypeCommitteeDelete
	QuerierRoute                    = types.QuerierRoute
	QueryCommittee                  = types.QueryCommittee
	QueryCommittees                 = types.QueryCommittees
	QueryNextProposalID             = types.QueryNextProposalID
	QueryProposal                   = types.QueryProposal
	QueryProposals                  = types.QueryProposals
	QueryRawParams                  = types.QueryRawParams
	QueryTally                      = types.QueryTally
	QueryVote                       = types.QueryVote
	QueryVotes                      = types.QueryVotes
	RouterKey                       = types.RouterKey
	StoreKey                        = types.StoreKey
	TypeMsgSubmitProposal           = types.TypeMsgSubmitProposal
	TypeMsgVote                     = types.TypeMsgVote
	Yes                             = types.Yes
)

var (
	// function aliases
	NewKeeper                   = keeper.NewKeeper
	NewQuerier                  = keeper.NewQuerier
	RegisterInvariants          = keeper.RegisterInvariants
	ValidCommitteesInvariant    = keeper.ValidCommitteesInvariant
	ValidProposalsInvariant     = keeper.ValidProposalsInvariant
	ValidVotesInvariant         = keeper.ValidVotesInvariant
	DefaultGenesisState         = types.DefaultGenesisState
	GetKeyFromID                = types.GetKeyFromID
	GetVoteKey                  = types.GetVoteKey
	NewAllowedCollateralParam   = types.NewAllowedCollateralParam
	NewAllowedMoneyMarket       = types.NewAllowedMoneyMarket
	NewCommitteeChangeProposal  = types.NewCommitteeChangeProposal
	NewCommitteeDeleteProposal  = types.NewCommitteeDeleteProposal
	NewGenesisState             = types.NewGenesisState
	NewMemberCommittee          = types.NewMemberCommittee
	NewTokenCommittee           = types.NewTokenCommittee
	NewMsgSubmitProposal        = types.NewMsgSubmitProposal
	NewMsgVote                  = types.NewMsgVote
	NewProposal                 = types.NewProposal
	NewQueryCommitteeParams     = types.NewQueryCommitteeParams
	NewQueryProposalParams      = types.NewQueryProposalParams
	NewQueryRawParamsParams     = types.NewQueryRawParamsParams
	NewQueryVoteParams          = types.NewQueryVoteParams
	NewVote                     = types.NewVote
	RegisterCodec               = types.RegisterCodec
	RegisterPermissionTypeCodec = types.RegisterPermissionTypeCodec
	RegisterProposalTypeCodec   = types.RegisterProposalTypeCodec
	Uint64FromBytes             = types.Uint64FromBytes

	// variable aliases
	ProposalHandler            = client.ProposalHandler
	CommitteeKeyPrefix         = types.CommitteeKeyPrefix
	ErrInvalidCommittee        = types.ErrInvalidCommittee
	ErrInvalidGenesis          = types.ErrInvalidGenesis
	ErrInvalidPubProposal      = types.ErrInvalidPubProposal
	ErrNoProposalHandlerExists = types.ErrNoProposalHandlerExists
	ErrProposalExpired         = types.ErrProposalExpired
	ErrUnknownCommittee        = types.ErrUnknownCommittee
	ErrUnknownProposal         = types.ErrUnknownProposal
	ErrUnknownSubspace         = types.ErrUnknownSubspace
	ErrUnknownVote             = types.ErrUnknownVote
	ModuleCdc                  = types.ModuleCdc
	NextProposalIDKey          = types.NextProposalIDKey
	ProposalKeyPrefix          = types.ProposalKeyPrefix
	VoteKeyPrefix              = types.VoteKeyPrefix
)

type (
	Keeper                      = keeper.Keeper
	AllowedAssetParam           = types.AllowedAssetParam
	AllowedAssetParams          = types.AllowedAssetParams
	AllowedCollateralParam      = types.AllowedCollateralParam
	AllowedCollateralParams     = types.AllowedCollateralParams
	AllowedDebtParam            = types.AllowedDebtParam
	AllowedMarket               = types.AllowedMarket
	AllowedMarkets              = types.AllowedMarkets
	AllowedMoneyMarket          = types.AllowedMoneyMarket
	AllowedMoneyMarkets         = types.AllowedMoneyMarkets
	AllowedParam                = types.AllowedParam
	AllowedParams               = types.AllowedParams
	Committee                   = types.Committee
	BaseCommittee               = types.BaseCommittee
	CommitteeChangeProposal     = types.CommitteeChangeProposal
	CommitteeDeleteProposal     = types.CommitteeDeleteProposal
	GenesisState                = types.GenesisState
	GodPermission               = types.GodPermission
	MsgSubmitProposal           = types.MsgSubmitProposal
	MemberCommittee             = types.MemberCommittee
	MsgVote                     = types.MsgVote
	TokenCommittee              = types.TokenCommittee
	ParamKeeper                 = types.ParamKeeper
	Permission                  = types.Permission
	Proposal                    = types.Proposal
	PubProposal                 = types.PubProposal
	QueryCommitteeParams        = types.QueryCommitteeParams
	QueryProposalParams         = types.QueryProposalParams
	QueryRawParamsParams        = types.QueryRawParamsParams
	QueryVoteParams             = types.QueryVoteParams
	SimpleParamChangePermission = types.SimpleParamChangePermission
	SoftwareUpgradePermission   = types.SoftwareUpgradePermission
	SubParamChangePermission    = types.SubParamChangePermission
	TextPermission              = types.TextPermission
	Vote                        = types.Vote
)
