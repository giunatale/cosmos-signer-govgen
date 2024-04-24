package simulation

// DONTCOVER

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/atomone-hub/govgen/x/gov/types"
)

// Simulation parameter constants
const (
	DepositParamsMinDeposit                 = "deposit_params_min_deposit"
	DepositParamsDepositPeriod              = "deposit_params_deposit_period"
	VotingParamsVotingPeriodDefault         = "voting_params_voting_period_default"
	VotingParamsVotingPeriodParameterChange = "voting_params_voting_period_parameter_change"
	VotingParamsVotingPeriodSoftwareUpgrade = "voting_params_voting_period_software_upgrade"
	VotingParamsVotingPeriodText            = "voting_params_voting_period_text"
	TallyParamsQuorum                       = "tally_params_quorum"
	TallyParamsThreshold                    = "tally_params_threshold"
	TallyParamsVeto                         = "tally_params_veto"
)

// GenDepositParamsDepositPeriod randomized DepositParamsDepositPeriod
func GenDepositParamsDepositPeriod(r *rand.Rand) time.Duration {
	return time.Duration(simulation.RandIntBetween(r, 1, 2*60*60*24*2)) * time.Second
}

// GenDepositParamsMinDeposit randomized DepositParamsMinDeposit
func GenDepositParamsMinDeposit(r *rand.Rand) sdk.Coins {
	return sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, int64(simulation.RandIntBetween(r, 1, 1e3))))
}

// GenVotingParamsVotingPeriod randomized VotingParamsVotingPeriod
func GenVotingParamsVotingPeriod(r *rand.Rand) time.Duration {
	return time.Duration(simulation.RandIntBetween(r, 1, 2*60*60*24*2)) * time.Second
}

// GenTallyParamsQuorum randomized TallyParamsQuorum
func GenTallyParamsQuorum(r *rand.Rand) sdk.Dec {
	return sdk.NewDecWithPrec(int64(simulation.RandIntBetween(r, 334, 500)), 3)
}

// GenTallyParamsThreshold randomized TallyParamsThreshold
func GenTallyParamsThreshold(r *rand.Rand) sdk.Dec {
	return sdk.NewDecWithPrec(int64(simulation.RandIntBetween(r, 450, 550)), 3)
}

// GenTallyParamsVeto randomized TallyParamsVeto
func GenTallyParamsVeto(r *rand.Rand) sdk.Dec {
	return sdk.NewDecWithPrec(int64(simulation.RandIntBetween(r, 250, 334)), 3)
}

// RandomizedGenState generates a random GenesisState for gov
func RandomizedGenState(simState *module.SimulationState) {
	startingProposalID := uint64(simState.Rand.Intn(100))

	var minDeposit sdk.Coins
	simState.AppParams.GetOrGenerate(
		simState.Cdc, DepositParamsMinDeposit, &minDeposit, simState.Rand,
		func(r *rand.Rand) { minDeposit = GenDepositParamsMinDeposit(r) },
	)

	var depositPeriod time.Duration
	simState.AppParams.GetOrGenerate(
		simState.Cdc, DepositParamsDepositPeriod, &depositPeriod, simState.Rand,
		func(r *rand.Rand) { depositPeriod = GenDepositParamsDepositPeriod(r) },
	)

	var votingPeriodDefault time.Duration
	simState.AppParams.GetOrGenerate(
		simState.Cdc, VotingParamsVotingPeriodDefault, &votingPeriodDefault, simState.Rand,
		func(r *rand.Rand) { votingPeriodDefault = GenVotingParamsVotingPeriod(r) },
	)
	var votingPeriodParameterChange time.Duration
	simState.AppParams.GetOrGenerate(
		simState.Cdc, VotingParamsVotingPeriodParameterChange, &votingPeriodParameterChange, simState.Rand,
		func(r *rand.Rand) { votingPeriodParameterChange = GenVotingParamsVotingPeriod(r) },
	)
	var votingPeriodSoftwareUpgrade time.Duration
	simState.AppParams.GetOrGenerate(
		simState.Cdc, VotingParamsVotingPeriodSoftwareUpgrade, &votingPeriodSoftwareUpgrade, simState.Rand,
		func(r *rand.Rand) { votingPeriodSoftwareUpgrade = GenVotingParamsVotingPeriod(r) },
	)
	var votingPeriodText time.Duration
	simState.AppParams.GetOrGenerate(
		simState.Cdc, VotingParamsVotingPeriodText, &votingPeriodText, simState.Rand,
		func(r *rand.Rand) { votingPeriodText = GenVotingParamsVotingPeriod(r) },
	)

	var quorum sdk.Dec
	simState.AppParams.GetOrGenerate(
		simState.Cdc, TallyParamsQuorum, &quorum, simState.Rand,
		func(r *rand.Rand) { quorum = GenTallyParamsQuorum(r) },
	)

	var threshold sdk.Dec
	simState.AppParams.GetOrGenerate(
		simState.Cdc, TallyParamsThreshold, &threshold, simState.Rand,
		func(r *rand.Rand) { threshold = GenTallyParamsThreshold(r) },
	)

	var veto sdk.Dec
	simState.AppParams.GetOrGenerate(
		simState.Cdc, TallyParamsVeto, &veto, simState.Rand,
		func(r *rand.Rand) { veto = GenTallyParamsVeto(r) },
	)

	govGenesis := types.NewGenesisState(
		startingProposalID,
		types.NewDepositParams(minDeposit, depositPeriod),
		types.NewVotingParams(votingPeriodDefault, votingPeriodParameterChange,
			votingPeriodSoftwareUpgrade, votingPeriodText),
		types.NewTallyParams(quorum, threshold, veto),
	)

	bz, err := json.MarshalIndent(&govGenesis, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated governance parameters:\n%s\n", bz)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(govGenesis)
}