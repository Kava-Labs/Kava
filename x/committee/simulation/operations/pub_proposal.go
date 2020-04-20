package operations

import (
	"fmt"
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govsimops "github.com/cosmos/cosmos-sdk/x/gov/simulation/operations"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/kava-labs/kava/x/committee"
	committeesims "github.com/kava-labs/kava/x/committee/simulation"
)

type PubProposalSimulator func(r *rand.Rand, ctx sdk.Context, accs []simulation.Account, perm committee.Permission) committee.PubProposal

// Generating pubProposals is different for committees compared to gov. Committees have permissions that limit the valid pubproposals.
// Solution here passes a permission into the pubproposal generation function. This function converts existing gov content generators into PubProposal simulators.
func SimulateAnyPubProposal(textSimulator govsimops.ContentSimulator, paramSimulator govsimops.ContentSimulator, paramFromPermSimulator PubProposalSimulator) PubProposalSimulator {
	return func(r *rand.Rand, ctx sdk.Context, accs []simulation.Account, permission committee.Permission) committee.PubProposal {
		switch perm := permission.(type) {
		case committee.GodPermission:
			switch r.Intn(2) {
			case 0:
				return textSimulator(r, ctx, accs)
			default:
				return paramSimulator(r, ctx, accs)
			}
		case committee.ParamChangePermission:
			return paramFromPermSimulator(r, ctx, accs, perm)
		case committee.TextPermission:
			return textSimulator(r, ctx, accs)
		default:
			panic(fmt.Sprintf("unexpected permission type %T", permission))
		}
	}
}

func SimulateParamChangePubProposal(paramChanges []simulation.ParamChange) PubProposalSimulator {
	return func(r *rand.Rand, ctx sdk.Context, accs []simulation.Account, permission committee.Permission) committee.PubProposal {
		perm, ok := permission.(committee.ParamChangePermission)
		if !ok {
			panic("expected permission to be of type ParamChangePermission")
		}

		// get available params that are allowed for the committee
		var availableParamChanges []simulation.ParamChange
		for _, pc := range paramChanges {
			for _, ap := range perm.AllowedParams {
				if paramChangeAllowed(pc, ap) {
					availableParamChanges = append(availableParamChanges, pc)
				}
				// this could produce duplicate if paramChanges, or AllowedParams contain duplicates
			}
		}

		// generate param changes
		numChanges := r.Intn(len(availableParamChanges) + 1)
		if numChanges < 1 {
			return nil // must be ≥ 1 param change to satisfy validate basic
		}
		indexes := r.Perm(len(availableParamChanges))[:numChanges]

		var changes []params.ParamChange
		for _, i := range indexes {
			pc := availableParamChanges[i]
			changes = append(changes, params.NewParamChangeWithSubkey(pc.Subspace, pc.Key, pc.Subkey, pc.SimValue(r)))
		}

		return params.NewParameterChangeProposal(
			simulation.RandStringOfLength(r, 140),
			simulation.RandStringOfLength(r, 5000),
			changes,
		)
	}
}

func paramChangeAllowed(pc simulation.ParamChange, ap committee.AllowedParam) bool {
	return ap.Subspace+ap.Key+ap.Subkey == pc.ComposedKey()
}

// SimulateCommitteeChangeProposalContent generates gov proposal contents that either:
// - create new committees
// - change existing committees
// - delete committees
func SimulateCommitteeChangeProposalContent(k committee.Keeper) govsimops.ContentSimulator {
	return func(r *rand.Rand, ctx sdk.Context, accs []simulation.Account) govtypes.Content {
		allowedParams := committeesims.GetAllowedParamKeys()

		// get current committees
		var committees []committee.Committee
		k.IterateCommittees(ctx, func(com committee.Committee) bool {
			committees = append(committees, com)
			return false
		})
		if len(committees) < 1 { // create a committee if none exist
			com, err := committeesims.RandomCommittee(r, accs, allowedParams)
			if err != nil {
				panic(err)
			}
			return committee.NewCommitteeChangeProposal(
				simulation.RandStringOfLength(r, 10),
				simulation.RandStringOfLength(r, 100),
				com,
			)
		}

		// create a proposal content

		var content govtypes.Content
		switch choice := r.Intn(100); {
		// create committee
		case choice < 20:
			com, err := committeesims.RandomCommittee(r, accs, allowedParams)
			if err != nil {
				panic(err)
			}
			content = committee.NewCommitteeChangeProposal(
				simulation.RandStringOfLength(r, 10),
				simulation.RandStringOfLength(r, 100),
				com,
			)
		// update committee
		case choice < 80:
			com := committees[r.Intn(len(committees))]

			// update members
			if r.Intn(100) < 50 {
				var members []sdk.AccAddress
				for len(members) < 1 {
					members = committeesims.RandomAddresses(r, accs)
				}
				com.Members = members
			}
			// update permissions
			if r.Intn(100) < 50 {
				var members []sdk.AccAddress
				for len(members) < 1 {
					members = committeesims.RandomAddresses(r, accs)
				}
				com.Permissions = committeesims.RandomPermissions(r, allowedParams)
			}
			// update proposal duration
			if r.Intn(100) < 50 {
				dur, err := RandomPositiveDuration(r, 0, committeesims.AverageBlockTime*100)
				if err != nil {
					panic(err)
				}
				com.MaxProposalDuration = dur
			}
			// update vote threshold
			if r.Intn(100) < 50 {
				com.VoteThreshold = simulation.RandomDecAmount(r, sdk.MustNewDecFromStr("1.00"))
			}

			content = committee.NewCommitteeChangeProposal(
				simulation.RandStringOfLength(r, 10),
				simulation.RandStringOfLength(r, 100),
				com,
			)
		// delete committee
		default:
			com := committees[r.Intn(len(committees))]
			content = committee.NewCommitteeDeleteProposal(
				simulation.RandStringOfLength(r, 10),
				simulation.RandStringOfLength(r, 100),
				com.ID,
			)
		}

		return content
	}
}
