package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/tss module sentinel errors
var (
	ErrInvalidAccAddressFormat = sdkerrors.Register(ModuleName, 2, "account address format is invalid")
	ErrGroupNotFound           = sdkerrors.Register(ModuleName, 3, "group not found")
	ErrMemberNotFound          = sdkerrors.Register(ModuleName, 4, "member not found")
	ErrAlreadySubmit           = sdkerrors.Register(ModuleName, 5, "member is already submit message")
	ErrRound1DataNotFound      = sdkerrors.Register(ModuleName, 6, "round 1 data not found")
	ErrDKGContextNotFound      = sdkerrors.Register(ModuleName, 7, "dkg context not found")
	ErrMemberNotAuthorized     = sdkerrors.Register(
		ModuleName,
		8,
		"member is not authorized for this group",
	)
	ErrRoundExpired                          = sdkerrors.Register(ModuleName, 9, "round expired")
	ErrVerifyOneTimeSigFailed                = sdkerrors.Register(ModuleName, 10, "fail to verify one time sign")
	ErrVerifyA0SigFailed                     = sdkerrors.Register(ModuleName, 11, "fail to verify a0 sign")
	ErrRound2DataNotFound                    = sdkerrors.Register(ModuleName, 12, "round 2 data not found")
	ErrEncryptedSecretSharesNotCorrectLength = sdkerrors.Register(
		ModuleName,
		13,
		"encrypted secret shares not correct length",
	)
	ErrMemberIsAlreadyMalicious         = sdkerrors.Register(ModuleName, 14, "member is already malicious")
	ErrMemberIsAlreadyComplainOrConfirm = sdkerrors.Register(ModuleName, 15, "member is already complain or confirm")
	ErrComplainFailed                   = sdkerrors.Register(ModuleName, 16, "complain failed")
	ErrConfirmFailed                    = sdkerrors.Register(ModuleName, 17, "confirm failed")
	ErrConfirmNotFound                  = sdkerrors.Register(ModuleName, 18, "confirm not found")
	ErrComplainsWithStatusNotFound      = sdkerrors.Register(ModuleName, 19, "complains with status not found")
)
