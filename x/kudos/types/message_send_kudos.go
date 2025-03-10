package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSendKudos{}

func NewMsgSendKudos(creator string, sender string, receiver string, message string) *MsgSendKudos {
	return &MsgSendKudos{
		Creator:  creator,
		Sender:   sender,
		Receiver: receiver,
		Message:  message,
	}
}

func (msg *MsgSendKudos) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
