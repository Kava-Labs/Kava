package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RewardPeriod stores the state of an ongoing reward
type RewardPeriod struct {
	Denom         string        `json:"denom" yaml:"denom"`
	Start         time.Time     `json:"start" yaml:"start"`
	End           time.Time     `json:"end" yaml:"end"`
	Reward        sdk.Coin      `json:"reward" yaml:"reward"` // per second reward payouts
	ClaimEnd      time.Time     `json:"claim_end" yaml:"claim_end"`
	ClaimTimeLock time.Duration `json:"claim_time_lock" yaml:"claim_time_lock"` // the amount of time rewards are timelocked once they are sent to users
}

// String implements fmt.Stringer
func (rp RewardPeriod) String() string {
	return fmt.Sprintf(`Reward Period:
	Denom: %s,
	Start: %s,
	End: %s,
	Reward: %s,
	Claim End: %s,
	Time Lock: %s`,
		rp.Denom, rp.Start, rp.End, rp.Reward, rp.ClaimEnd, rp.ClaimTimeLock)
}

// NewRewardPeriod returns a new RewardPeriod
func NewRewardPeriod(denom string, start time.Time, end time.Time, reward sdk.Coin, claimEnd time.Time, claimTimeLock time.Duration) RewardPeriod {
	return RewardPeriod{
		Denom:         denom,
		Start:         start,
		End:           end,
		Reward:        reward,
		ClaimEnd:      claimEnd,
		ClaimTimeLock: claimTimeLock,
	}
}

// RewardPeriods array of RewardPeriod
type RewardPeriods []RewardPeriod

// ClaimPeriod stores the state of an ongoing claim period
type ClaimPeriod struct {
	Denom    string        `json:"denom" yaml:"denom"`
	ID       uint64        `json:"id" yaml:"id"`
	End      time.Time     `json:"end" yaml:"end"`
	TimeLock time.Duration `json:"claim_time_lock" yaml:"claim_time_lock"`
}

// NewClaimPeriod returns a new ClaimPeriod
func NewClaimPeriod(denom string, id uint64, end time.Time, timeLock time.Duration) ClaimPeriod {
	return ClaimPeriod{
		Denom:    denom,
		ID:       id,
		End:      end,
		TimeLock: timeLock,
	}
}

// ClaimPeriods array of ClaimPeriod
type ClaimPeriods []ClaimPeriod

// Claim stores the rewards that can be claimed by owner
type Claim struct {
	Owner         sdk.AccAddress `json:"owner" yaml:"owner"`
	Reward        sdk.Coin       `json:"reward" yaml:"reward"`
	Denom         string         `json:"denom" yaml:"denom"`
	ClaimPeriodID uint64         `json:"id" yaml:"id"`
}

// NewClaim returns a new Claim
func NewClaim(owner sdk.AccAddress, reward sdk.Coin, denom string, id uint64) Claim {
	return Claim{
		Owner:         owner,
		Reward:        reward,
		Denom:         denom,
		ClaimPeriodID: id,
	}
}

// String implements fmt.Stringer
func (c Claim) String() string {
	return fmt.Sprintf(`Claim:
	Owner: %s,
	Denom: %s,
	Reward: %s,
	Caim Period ID: %d,`,
		c.Denom, c.Denom, c.Reward, c.ClaimPeriodID)
}

// Claims array of Claim
type Claims []Claim
