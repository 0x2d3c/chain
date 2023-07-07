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
	for _, c := range m.Round1Info.CoefficientCommits {
		if err := c.Validate(); err != nil {
			return sdkerrors.Wrap(err, "coefficients commit")
		}
	}

	// Validate one time pub key
	if err := m.Round1Info.OneTimePubKey.Validate(); err != nil {
		return sdkerrors.Wrap(err, "one time pub key")
	}

	// Validate a0 signature
	if err := m.Round1Info.A0Sig.Validate(); err != nil {
		return sdkerrors.Wrap(err, "a0 sig")
	}

	// Validate one time signature
	if err := m.Round1Info.OneTimeSig.Validate(); err != nil {
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
	for _, ess := range m.Round2Info.EncryptedSecretShares {
		if err := ess.Validate(); err != nil {
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

	// Validate complaints size
	if len(m.Complaints) < 1 {
		return sdkerrors.Wrap(fmt.Errorf("must contain at least one complaint"), "complaints")
	}

	// Validate complaints
	memberI := m.Complaints[0].Complainer
	for i, c := range m.Complaints {
		// Validate member complainer
		if i > 0 && memberI != c.Complainer {
			return sdkerrors.Wrap(
				fmt.Errorf("memberID complainer in the list of complaints must be the same value"),
				"complainer",
			)
		}

		// Validate member complainer and complainant
		if c.Complainer == c.Complainant {
			return sdkerrors.Wrap(
				fmt.Errorf("memberID complainer and complainant can not be the same value"),
				"complainer, complainant",
			)
		}

		// Validate key sym
		if err := c.KeySym.Validate(); err != nil {
			return sdkerrors.Wrap(err, "key sym")
		}

		// Validate signature
		if err := c.Signature.Validate(); err != nil {
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
	if err = m.OwnPubKeySig.Validate(); err != nil {
		return sdkerrors.Wrap(err, "own pub key sig")
	}

	return nil
}

var _ sdk.Msg = &MsgSubmitDEs{}

// Route Implements Msg.
func (m MsgSubmitDEs) Route() string { return sdk.MsgTypeURL(&m) }

// Type Implements Msg.
func (m MsgSubmitDEs) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes Implements Msg.
func (m MsgSubmitDEs) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// GetSigners returns the expected signers for a MsgCreateGroup.
func (m MsgSubmitDEs) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Member)}
}

// ValidateBasic does a sanity check on the provided data
func (m MsgSubmitDEs) ValidateBasic() error {
	// Validate member address
	_, err := sdk.AccAddressFromBech32(m.Member)
	if err != nil {
		return sdkerrors.Wrap(err, "member")
	}

	// Validate DEs
	for i, de := range m.DEs {
		// Validate public key D
		if err = de.PubD.Validate(); err != nil {
			return sdkerrors.Wrap(err, fmt.Sprintf("pub D in DE index: %d", i))
		}

		// Validate public key E
		if err = de.PubE.Validate(); err != nil {
			return sdkerrors.Wrap(err, fmt.Sprintf("pub E in DE index: %d", i))
		}
	}

	return nil
}

var _ sdk.Msg = &MsgRequestSignature{}

// Route Implements Msg.
func (m MsgRequestSignature) Route() string { return sdk.MsgTypeURL(&m) }

// Type Implements Msg.
func (m MsgRequestSignature) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes Implements Msg.
func (m MsgRequestSignature) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// GetSigners returns the expected signers for a MsgCreateGroup.
func (m MsgRequestSignature) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Sender)}
}

// ValidateBasic does a sanity check on the provided data
func (m MsgRequestSignature) ValidateBasic() error {
	// Validate sender address
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return sdkerrors.Wrap(err, "sender")
	}

	return nil
}

var _ sdk.Msg = &MsgSign{}

// Route Implements Msg.
func (m MsgSign) Route() string { return sdk.MsgTypeURL(&m) }

// Type Implements Msg.
func (m MsgSign) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes Implements Msg.
func (m MsgSign) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// GetSigners returns the expected signers for a MsgCreateGroup.
func (m MsgSign) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Member)}
}

// ValidateBasic does a sanity check on the provided data
func (m MsgSign) ValidateBasic() error {
	// Validate member address
	_, err := sdk.AccAddressFromBech32(m.Member)
	if err != nil {
		return sdkerrors.Wrap(err, "member")
	}

	// Validate member signature
	if err = m.Signature.Validate(); err != nil {
		return sdkerrors.Wrap(err, "signature")
	}

	return nil
}
