package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgByName = "by_name"

var _ sdk.Msg = &MsgByName{}

func NewMsgByName(creator string, name string, bid string) *MsgByName {
	return &MsgByName{
		Creator: creator,
		Name:    name,
		Bid:     bid,
	}
}

func (msg *MsgByName) Route() string {
	return RouterKey
}


func (msg *MsgByName) Type() string {
	return TypeMsgByName
}

func (msg *MsgByName) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgByName) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgByName) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
