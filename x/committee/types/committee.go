package types

import (
	"fmt"
	"time"

	yaml "gopkg.in/yaml.v2"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const MaxCommitteeDescriptionLength int = 512

type TallyOption int

const (
	NullTallyOption  TallyOption = iota
	FirstPastThePost TallyOption = iota // Votes are tallied each block and the proposal passes as soon as the vote threshold is reached
	Deadline         TallyOption = iota // Votes are tallied exactly once, when the deadline time is reached
)

const (
	MemberCommitteeType = "member" // Committee is composed of member addresses that vote to enact proposals within their permissions
	TokenCommitteeType  = "token"  // Committee is composed of token holders with voting power determined by total token balance
	BondDenom           = "ukava"
)

// Committee is an interface for handling common actions on committees
type Committee interface {
	GetID() uint64
	GetType() string
	GetDescription() string

	SetMembers([]sdk.AccAddress)
	GetMembers() []sdk.AccAddress
	HasMember(addr sdk.AccAddress) bool

	SetPermissions([]Permission)
	GetPermissions() []Permission
	HasPermissionsFor(ctx sdk.Context, appCdc *codec.Codec, pk ParamKeeper, proposal PubProposal) bool

	SetProposalDuration(time.Duration)
	GetProposalDuration() time.Duration

	SetVoteThreshold(sdk.Dec)
	GetVoteThreshold() sdk.Dec

	GetTallyOption() TallyOption
	Validate() error
}

// Committees is a slice of committees
type Committees []Committee

// BaseCommittee is a common type shared by all Committees
type BaseCommittee struct {
	ID               uint64           `json:"id" yaml:"id"`
	Description      string           `json:"description" yaml:"description"`
	Members          []sdk.AccAddress `json:"members" yaml:"members"`
	Permissions      []Permission     `json:"permissions" yaml:"permissions"`
	VoteThreshold    sdk.Dec          `json:"vote_threshold" yaml:"vote_threshold"`       // Smallest percentage that must vote for a proposal to pass
	ProposalDuration time.Duration    `json:"proposal_duration" yaml:"proposal_duration"` // The length of time a proposal remains active for. Proposals will close earlier if they get enough votes.
	TallyOption      TallyOption      `json:"tally_option" yaml:"tally_option"`
}

// GetID is a getter for committee ID
func (c BaseCommittee) GetID() uint64 { return c.ID }

// GetDescription is a getter for committee description
func (c BaseCommittee) GetDescription() string { return c.Description }

// SetMembers is a setter for committee members
func (c BaseCommittee) SetMembers(members []sdk.AccAddress) { c.Members = members }

// GetMembers is a getter for committee members
func (c BaseCommittee) GetMembers() []sdk.AccAddress { return c.Members }

// HasMember returns if a committee contains a given member address
func (c BaseCommittee) HasMember(addr sdk.AccAddress) bool {
	for _, m := range c.Members {
		if m.Equals(addr) {
			return true
		}
	}
	return false
}

// SetPermissions is a setter for committee permissions
func (c BaseCommittee) SetPermissions(permissions []Permission) { c.Permissions = permissions }

// GetPermissions is a getter for committee permissions
func (c BaseCommittee) GetPermissions() []Permission { return c.Permissions }

// HasPermissionsFor returns whether the committee is authorized to enact a proposal.
// As long as one permission allows the proposal then it goes through. Its the OR of all permissions.
func (c BaseCommittee) HasPermissionsFor(ctx sdk.Context, appCdc *codec.Codec, pk ParamKeeper, proposal PubProposal) bool {
	for _, p := range c.Permissions {
		if p.Allows(ctx, appCdc, pk, proposal) {
			return true
		}
	}
	return false
}

// SetVoteThreshold is a setter for committee VoteThreshold
func (c BaseCommittee) SetVoteThreshold(threshold sdk.Dec) { c.VoteThreshold = threshold }

// GetVoteThreshold is a getter for committee VoteThreshold
func (c BaseCommittee) GetVoteThreshold() sdk.Dec { return c.VoteThreshold }

// SetProposalDuration is a setter for committee ProposalDuration
func (c BaseCommittee) SetProposalDuration(duration time.Duration) { c.ProposalDuration = duration }

// GetProposalDuration is a getter for committee ProposalDuration
func (c BaseCommittee) GetProposalDuration() time.Duration { return c.ProposalDuration }

// GetTallyOption is a getter for committee TallyOption
func (c BaseCommittee) GetTallyOption() TallyOption { return c.TallyOption }

// Validate validates BaseCommittee fields
func (c BaseCommittee) Validate() error {
	if len(c.Description) > MaxCommitteeDescriptionLength {
		return fmt.Errorf("description length %d longer than max allowed %d", len(c.Description), MaxCommitteeDescriptionLength)
	}

	if len(c.Members) <= 0 {
		return fmt.Errorf("committee must have members")
	}

	addressMap := make(map[string]bool, len(c.Members))
	for _, m := range c.Members {
		// check there are no duplicate members
		if _, ok := addressMap[m.String()]; ok {
			return fmt.Errorf("committe cannot have duplicate members, %s", m)
		}
		// check for valid addresses
		if m.Empty() {
			return fmt.Errorf("committee cannot have empty member address")
		}
		addressMap[m.String()] = true
	}

	for _, p := range c.Permissions {
		if p == nil {
			return fmt.Errorf("committee cannot have a nil permission")
		}
	}

	if c.ProposalDuration < 0 {
		return fmt.Errorf("invalid proposal duration: %s", c.ProposalDuration)
	}

	// threshold must be in the range [0,1]
	if c.VoteThreshold.IsNil() || c.VoteThreshold.LTE(sdk.ZeroDec()) || c.VoteThreshold.GT(sdk.NewDec(1)) {
		return fmt.Errorf("invalid threshold: %s", c.VoteThreshold)
	}

	return nil
}

// String implements fmt.Stringer
func (c BaseCommittee) String() string {
	return fmt.Sprintf(`Committee %d:
  Description:              %s
  Members:               %s
  Permissions:               			%s
  VoteThreshold:            		  %s
  ProposalDuration:        						%s
  TallyOption:   						%d`,
		c.ID, c.Description, c.GetMembers(), c.Permissions,
		c.VoteThreshold.String(), c.ProposalDuration.String(),
		c.TallyOption,
	)
}

// MemberCommittee is an alias of BaseCommittee
type MemberCommittee struct {
	BaseCommittee `json:"base_committee" yaml:"base_committee"`
}

// NewMemberCommittee instantiates a new instance of MemberCommittee
func NewMemberCommittee(id uint64, description string, members []sdk.AccAddress, permissions []Permission,
	threshold sdk.Dec, duration time.Duration, tallyOption TallyOption) MemberCommittee {
	return MemberCommittee{
		BaseCommittee: BaseCommittee{
			ID:               id,
			Description:      description,
			Members:          members,
			Permissions:      permissions,
			VoteThreshold:    threshold,
			ProposalDuration: duration,
			TallyOption:      tallyOption,
		},
	}
}

// GetType returns the type of the committee
func (c MemberCommittee) GetType() string { return MemberCommitteeType }

// Validate validates the committee's fields
func (c MemberCommittee) Validate() error {
	return c.BaseCommittee.Validate()
}

// TokenCommittee supports voting on proposals by token holders
type TokenCommittee struct {
	BaseCommittee `json:"base_committee" yaml:"base_committee"`
	Quorum        sdk.Dec `json:"quorum" yaml:"quorum"`
	TallyDenom    string  `json:"tally_denom" yaml:"tally_denom"`
}

// NewTokenCommittee instantiates a new instance of TokenCommittee
func NewTokenCommittee(id uint64, description string, permissions []Permission, threshold sdk.Dec,
	duration time.Duration, tallyOption TallyOption, quorum sdk.Dec, tallyDenom string) TokenCommittee {
	return TokenCommittee{
		BaseCommittee: BaseCommittee{
			ID:               id,
			Description:      description,
			Permissions:      permissions,
			VoteThreshold:    threshold,
			ProposalDuration: duration,
			TallyOption:      tallyOption,
		},
		Quorum:     quorum,
		TallyDenom: tallyDenom,
	}
}

// GetType returns the type of the committee
func (c TokenCommittee) GetType() string { return TokenCommitteeType }

// GetQuorum returns the quorum of the committee
func (c TokenCommittee) GetQuorum() sdk.Dec { return c.Quorum }

// GetTallyDenom returns the tally denom of the committee
func (c TokenCommittee) GetTallyDenom() string { return c.TallyDenom }

// TODO: don't allow nil TallyOption
// Validate validates the committee's fields
func (c TokenCommittee) Validate() error {
	if c.TallyDenom == BondDenom {
		return fmt.Errorf("invalid tally denom: %s", c.TallyDenom)
	}

	err := sdk.ValidateDenom(c.TallyDenom)
	if err != nil {
		return err
	}

	if c.Quorum.IsNil() || c.Quorum.IsNegative() || c.Quorum.GT(sdk.NewDec(1)) {
		return fmt.Errorf("invalid quroum: %s", c.Quorum)
	}

	return c.BaseCommittee.Validate()
}

// String implements fmt.Stringer
func (c TokenCommittee) String() string {
	return fmt.Sprintf(`Committee %d:
  Type:               %s
  Description:              %s
  Permissions:               			%s
  VoteThreshold:            		  %s
  ProposalDuration:        						%s
  TallyOption:   						%d
  Quorum:               %s
  TallyDenom:               %s`,
		c.ID, c.GetType(), c.Description, c.Permissions,
		c.VoteThreshold.String(), c.ProposalDuration.String(),
		c.TallyOption, c.Quorum, c.TallyDenom,
	)
}

// ------------------------------------------
//				Proposals
// ------------------------------------------

// PubProposal is the interface that all proposals must fulfill to be submitted to a committee.
// Proposal types can be created external to this module. For example a ParamChangeProposal, or CommunityPoolSpendProposal.
// It is pinned to the equivalent type in the gov module to create compatibility between proposal types.
type PubProposal govtypes.Content

// Proposal is an internal record of a governance proposal submitted to a committee.
type Proposal struct {
	PubProposal `json:"pub_proposal" yaml:"pub_proposal"`
	ID          uint64    `json:"id" yaml:"id"`
	CommitteeID uint64    `json:"committee_id" yaml:"committee_id"`
	Deadline    time.Time `json:"deadline" yaml:"deadline"`
}

func NewProposal(pubProposal PubProposal, id uint64, committeeID uint64, deadline time.Time) Proposal {
	return Proposal{
		PubProposal: pubProposal,
		ID:          id,
		CommitteeID: committeeID,
		Deadline:    deadline,
	}
}

// HasExpiredBy calculates if the proposal will have expired by a certain time.
// All votes must be cast before deadline, those cast at time == deadline are not valid
func (p Proposal) HasExpiredBy(time time.Time) bool {
	return !time.Before(p.Deadline)
}

// String implements the fmt.Stringer interface, and importantly overrides the String methods inherited from the embedded PubProposal type.
func (p Proposal) String() string {
	bz, _ := yaml.Marshal(p)
	return string(bz)
}

// ------------------------------------------
//				Votes
// ------------------------------------------

type Vote struct {
	ProposalID uint64         `json:"proposal_id" yaml:"proposal_id"`
	Voter      sdk.AccAddress `json:"voter" yaml:"voter"`
	VoteType   VoteType       `json:"vote_type" yaml:"vote_type"`
}

func NewVote(proposalID uint64, voter sdk.AccAddress, voteType VoteType) Vote {
	return Vote{
		ProposalID: proposalID,
		Voter:      voter,
		VoteType:   voteType,
	}
}

func (v Vote) Validate() error {
	if v.Voter.Empty() {
		return fmt.Errorf("voter address cannot be empty")
	}

	if v.VoteType >= 3 { // 0 = Yes, 1 = No, 2 = Abstain
		return fmt.Errorf("invalid vote type: %d", v.VoteType)
	}
	return nil
}
