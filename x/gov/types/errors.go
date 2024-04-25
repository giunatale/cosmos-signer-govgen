package types

import (
	errorsmod "cosmossdk.io/errors"
)

// x/gov module sentinel errors
var (
	ErrUnknownProposal         = errorsmod.Register(ModuleName, 1020, "unknown proposal")
	ErrInactiveProposal        = errorsmod.Register(ModuleName, 1030, "inactive proposal")
	ErrAlreadyActiveProposal   = errorsmod.Register(ModuleName, 1040, "proposal already active")
	ErrInvalidProposalContent  = errorsmod.Register(ModuleName, 1050, "invalid proposal content")
	ErrInvalidProposalType     = errorsmod.Register(ModuleName, 1060, "invalid proposal type")
	ErrInvalidVote             = errorsmod.Register(ModuleName, 1070, "invalid vote option")
	ErrInvalidGenesis          = errorsmod.Register(ModuleName, 1080, "invalid genesis state")
	ErrNoProposalHandlerExists = errorsmod.Register(ModuleName, 1090, "no handler exists for proposal type")
)
