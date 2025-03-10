package keeper

import (
	"context"
	"fmt"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/avobl/simple-chain/x/kudos/types"
)

func (k msgServer) SendKudos(goCtx context.Context, msg *types.MsgSendKudos) (*types.MsgSendKudosResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Sender == msg.Receiver {
		return nil, sdkerrors.Wrap(types.ErrBadRequest, "Sender and receiver cannot be the same")
	}

	kudos := types.Kudos{
		Sender:   msg.Sender,
		Receiver: msg.Receiver,
		Message:  msg.Message,
	}

	k.AppendKudos(ctx, kudos)

	return &types.MsgSendKudosResponse{
		Message: fmt.Sprintf("Sent: %q from: %s to: %s", msg.Message, msg.Sender, msg.Receiver),
	}, nil
}
