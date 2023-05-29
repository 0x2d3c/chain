package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateGroup{}

// Route Implements Msg.
func (m MsgCreateGroup) Route() string { return sdk.MsgTypeURL(&m) }

// Type Implements Msg.
func (m MsgCreateGroup) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes Implements Msg.
func (m MsgCreateGroup) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// GetSigners returns the expected signers for a MsgCreateGroup.
func (m MsgCreateGroup) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Sender)}
}

// ValidateBasic does a sanity check on the provided data
func (m MsgCreateGroup) ValidateBasic() error {
	// Validate members address
	for _, member := range m.Members {
		_, err := sdk.AccAddressFromBech32(member)
		if err != nil {
			return sdkerrors.Wrap(
				err,
				fmt.Sprintf("member: %s ", member),
			)
		}
	}

	// Check duplicate member
	if DuplicateInArray(m.Members) {
		return sdkerrors.Wrap(fmt.Errorf("members can not duplicate"), "members")
	}

	// Validate sender address
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return sdkerrors.Wrap(
			err,
			fmt.Sprintf("sender: %s", m.Sender),
		)
	}

	// Validate threshold must be less than or equal to members but more than zero
	if m.Threshold > uint64(len(m.Members)) || m.Threshold <= 0 {
		return sdkerrors.Wrap(
			fmt.Errorf("threshold must be less than or equal to the members but more than zero"),
			"threshold",
		)
	}

	return nil
}

var _ sdk.Msg = &MsgSubmitDKGRound1{}

// Route Implements Msg.
func (m MsgSubmitDKGRound1) Route() string { return sdk.MsgTypeURL(&m) }

// Type Implements Msg.
func (m MsgSubmitDKGRound1) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes Implements Msg.
func (m MsgSubmitDKGRound1) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// GetSigners returns the expected signers for a MsgCreateGroup.
func (m MsgSubmitDKGRound1) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Member)}
}

// ValidateBasic does a sanity check on the provided data
func (m MsgSubmitDKGRound1) ValidateBasic() error {
	// Validate member address
	_, err := sdk.AccAddressFromBech32(m.Member)
	if err != nil {
		return sdkerrors.Wrap(err, "member")
	}

	// Validate coefficients commit
	for _, c := range m.Round1Data.CoefficientsCommit {
		_, err := c.Parse()
		if err != nil {
			return sdkerrors.Wrap(err, "coefficients commit")
		}
	}

	// Validate one time pub key
	_, err = m.Round1Data.OneTimePubKey.Parse()
	if err != nil {
		return sdkerrors.Wrap(err, "one time pub key")
	}

	// Validate a0 signature
	_, err = m.Round1Data.A0Sig.Parse()
	if err != nil {
		return sdkerrors.Wrap(err, "a0 sig")
	}

	// Validate one time signature
	_, err = m.Round1Data.OneTimeSig.Parse()
	if err != nil {
		return sdkerrors.Wrap(err, "one time sig")
	}

	return nil
}

var _ sdk.Msg = &MsgSubmitDKGRound2{}

// Route Implements Msg.
func (m MsgSubmitDKGRound2) Route() string { return sdk.MsgTypeURL(&m) }

// Type Implements Msg.
func (m MsgSubmitDKGRound2) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes Implements Msg.
func (m MsgSubmitDKGRound2) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// GetSigners returns the expected signers for a MsgCreateGroup.
func (m MsgSubmitDKGRound2) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Member)}
}

// ValidateBasic does a sanity check on the provided data
func (m MsgSubmitDKGRound2) ValidateBasic() error {
	// Validate member address
	_, err := sdk.AccAddressFromBech32(m.Member)
	if err != nil {
		return sdkerrors.Wrap(err, "member")
	}

	// Validate encrypted secret shares
	for _, ess := range m.Round2Data.EncryptedSecretShares {
		_, err = ess.Parse()
		if err != nil {
			return sdkerrors.Wrap(err, "encrypted secret shares")
		}
	}

	return nil
}

var _ sdk.Msg = &MsgComplain{}

// Route Implements Msg.
func (m MsgComplain) Route() string { return sdk.MsgTypeURL(&m) }

// Type Implements Msg.
func (m MsgComplain) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes Implements Msg.
func (m MsgComplain) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// GetSigners returns the expected signers for a MsgCreateGroup.
func (m MsgComplain) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Member)}
}

// ValidateBasic does a sanity check on the provided data
func (m MsgComplain) ValidateBasic() error {
	// Validate member address
	_, err := sdk.AccAddressFromBech32(m.Member)
	if err != nil {
		return sdkerrors.Wrap(err, "member")
	}

	// Validate complains
	memberI := m.Complains[0].I
	for i, c := range m.Complains {
		// Validate member I
		if i > 0 && memberI != c.I {
			return sdkerrors.Wrap(fmt.Errorf("member I in the list of complains must be the same value"), "I")
		}

		// Validate key sym
		_, err := c.KeySym.Parse()
		if err != nil {
			return sdkerrors.Wrap(err, "key sym")
		}
		// Validate nonce sym
		_, err = c.NonceSym.Parse()
		if err != nil {
			return sdkerrors.Wrap(err, "nonce sym")
		}
		// Validate signature
		_, err = c.Signature.Parse()
		if err != nil {
			return sdkerrors.Wrap(err, "signature")
		}
	}

	return nil
}

var _ sdk.Msg = &MsgConfirm{}

// Route Implements Msg.
func (m MsgConfirm) Route() string { return sdk.MsgTypeURL(&m) }

// Type Implements Msg.
func (m MsgConfirm) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes Implements Msg.
func (m MsgConfirm) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// GetSigners returns the expected signers for a MsgCreateGroup.
func (m MsgConfirm) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Member)}
}

// ValidateBasic does a sanity check on the provided data
func (m MsgConfirm) ValidateBasic() error {
	// Validate member address
	_, err := sdk.AccAddressFromBech32(m.Member)
	if err != nil {
		return sdkerrors.Wrap(err, "member")
	}

	// Validate own pub key sig
	_, err = m.OwnPubKeySig.Parse()
	if err != nil {
		return sdkerrors.Wrap(err, "own pub key sig")
	}

	return nil
}
