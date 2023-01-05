package keeper

import (
	"context"
	"encoding/base64"
	"errors"

	"dataocean/x/dataocean/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"

	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

const (
	contentProducerPercent = 30
)

func (k msgServer) SubmitPaySign(goCtx context.Context, msg *types.MsgSubmitPaySign) (*types.MsgSubmitPaySignResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	msgPaySign, err := k.parsePaySign(ctx, msg.PaySign)
	if err != nil {
		return nil, err
	}

	err = k.exchangePaySign(ctx, msg.Creator, msgPaySign)
	if err != nil {
		return nil, err
	}

	return &types.MsgSubmitPaySignResponse{}, nil
}

func (k msgServer) parsePaySign(ctx sdk.Context, paySignStr string) (*types.MsgPaySign, error) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	std.RegisterInterfaces(interfaceRegistry)
	interfaceRegistry.RegisterImplementations((*sdk.Msg)(nil), &types.MsgPaySign{})
	protoCodec := codec.NewProtoCodec(interfaceRegistry)
	txConfig := tx.NewTxConfig(protoCodec, tx.DefaultSignModes)

	txBytes, err := base64.StdEncoding.DecodeString(paySignStr)
	if err != nil {
		return nil, err
	}
	theTx, err := txConfig.TxDecoder()(txBytes)
	if err != nil {
		return nil, err
	}

	err = k.verifyPaySign(ctx, txConfig, theTx)
	if err != nil {
		return nil, err
	}

	msgs := theTx.GetMsgs()
	if len(msgs) == 0 {
		return nil, errors.New("signature message is empty")
	}
	return msgs[0].(*types.MsgPaySign), nil
}

func (k msgServer) verifyPaySign(ctx sdk.Context, txConfig client.TxConfig, theTx sdk.Tx) error {
	sigTx := theTx.(authsigning.SigVerifiableTx)
	signers := sigTx.GetSigners()
	sigs, err := sigTx.GetSignaturesV2()
	if err != nil {
		return err
	}
	if len(sigs) != len(signers) {
		return errors.New(" len(sigs) and len(signers) are not equal")
	}
	for i, sig := range sigs {
		sigAddr := sdk.AccAddress(sig.PubKey.Address())
		sigAccount := k.accountKeeper.GetAccount(ctx, sigAddr)

		if i >= len(signers) || !sigAddr.Equals(signers[i]) {
			return errors.New("signature does not match its respective signer")
		}

		signingData := authsigning.SignerData{
			Address:       sigAddr.String(),
			ChainID:       ctx.ChainID(),
			AccountNumber: sigAccount.GetAccountNumber(),
			Sequence:      sigAccount.GetSequence(),
			PubKey:        sig.PubKey,
		}
		err = authsigning.VerifySignature(sig.PubKey, signingData, sig.Data, txConfig.SignModeHandler(), sigTx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (k msgServer) exchangePaySign(ctx sdk.Context, submitAddr string, paySign *types.MsgPaySign) error {
	video, found := k.GetVideo(ctx, paySign.VideoId)
	if !found {
		return sdkerrors.ErrKeyNotFound
	}
	allAmount := int64(video.PriceMB * paySign.ReceivedSizeMB)
	cpAmount := allAmount * contentProducerPercent / 100

	moduleAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)

	_, err := k.authzKeeper.DispatchActions(ctx, moduleAddr, []sdk.Msg{
		&banktypes.MsgSend{
			Amount:      sdk.NewCoins(sdk.NewInt64Coin("token", allAmount-cpAmount)),
			FromAddress: paySign.Creator,
			ToAddress:   submitAddr,
		},
		&banktypes.MsgSend{
			Amount:      sdk.NewCoins(sdk.NewInt64Coin("token", cpAmount)),
			FromAddress: paySign.Creator,
			ToAddress:   video.Creator,
		},
	})
	return err
}
