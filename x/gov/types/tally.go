package types

import (
	yaml "gopkg.in/yaml.v2"

	cosmossdk_io_math "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ValidatorGovInfo used for tallying
type ValidatorGovInfo struct {
	Address             sdk.ValAddress              // address of the validator operator
	BondedTokens        cosmossdk_io_math.Int       // Power of a Validator
	DelegatorShares     cosmossdk_io_math.LegacyDec // Total outstanding delegator shares
	DelegatorDeductions cosmossdk_io_math.LegacyDec // Delegator deductions from validator's delegators voting independently
	Vote                WeightedVoteOptions         // Vote of the validator
}

// NewValidatorGovInfo creates a ValidatorGovInfo instance
func NewValidatorGovInfo(address sdk.ValAddress, bondedTokens cosmossdk_io_math.Int, delegatorShares,
	delegatorDeductions cosmossdk_io_math.LegacyDec, options WeightedVoteOptions,
) ValidatorGovInfo {
	return ValidatorGovInfo{
		Address:             address,
		BondedTokens:        bondedTokens,
		DelegatorShares:     delegatorShares,
		DelegatorDeductions: delegatorDeductions,
		Vote:                options,
	}
}

// NewTallyResult creates a new TallyResult instance
func NewTallyResult(yes, abstain, no, noWithVeto cosmossdk_io_math.Int) TallyResult {
	return TallyResult{
		Yes:        yes,
		Abstain:    abstain,
		No:         no,
		NoWithVeto: noWithVeto,
	}
}

// NewTallyResultFromMap creates a new TallyResult instance from a Option -> Dec map
func NewTallyResultFromMap(results map[VoteOption]cosmossdk_io_math.LegacyDec) TallyResult {
	return NewTallyResult(
		results[OptionYes].TruncateInt(),
		results[OptionAbstain].TruncateInt(),
		results[OptionNo].TruncateInt(),
		results[OptionNoWithVeto].TruncateInt(),
	)
}

// EmptyTallyResult returns an empty TallyResult.
func EmptyTallyResult() TallyResult {
	return NewTallyResult(cosmossdk_io_math.ZeroInt(), cosmossdk_io_math.ZeroInt(), cosmossdk_io_math.ZeroInt(), cosmossdk_io_math.ZeroInt())
}

// Equals returns if two proposals are equal.
func (tr TallyResult) Equals(comp TallyResult) bool {
	return tr.Yes.Equal(comp.Yes) &&
		tr.Abstain.Equal(comp.Abstain) &&
		tr.No.Equal(comp.No) &&
		tr.NoWithVeto.Equal(comp.NoWithVeto)
}

// String implements stringer interface
func (tr TallyResult) String() string {
	out, _ := yaml.Marshal(tr)
	return string(out)
}
