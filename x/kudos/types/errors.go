package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/kudos module sentinel errors
var (
	ErrInvalidSigner = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrSample        = sdkerrors.Register(ModuleName, 1101, "sample error")
	ErrBadRequest    = sdkerrors.Register(ModuleName, 1102, "bad request")
)
