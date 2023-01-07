package keeper

import (
	"context"
	"fmt"
	"math/rand"
	"net/url"
	"time"

	"dataocean/x/dataocean/types"

	"github.com/golang-module/dongle"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

const (
	minAmount          = 1000
	minValidTime       = 12 * time.Hour
	videoLinkValidTime = 12 * time.Hour
)

var servers = []struct {
	host   string
	aesKey string
}{
	{
		host:   "127.0.0.1",
		aesKey: "key_for_server_1",
	},
	{
		host:   "127.0.0.2",
		aesKey: "key_for_server_2",
	},
}

func (k msgServer) PlayVideo(goCtx context.Context, msg *types.MsgPlayVideo) (*types.MsgPlayVideoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetVideo(ctx, msg.VideoId)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	moduleAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)
	userAddr, _ := sdk.AccAddressFromBech32(msg.Creator)

	auth, exp := k.authzKeeper.GetAuthorization(ctx, moduleAddr, userAddr, sdk.MsgTypeURL(&banktypes.MsgSend{}))
	if auth == nil {
		return nil, fmt.Errorf("authorization not exists")
	}
	sendAuth := auth.(*banktypes.SendAuthorization)
	if exp != nil && (*exp).Before(ctx.BlockTime().Add(minAmount)) {
		return nil, fmt.Errorf("authorization valid time cannot be less than %.0f hours", minValidTime.Hours())
	}
	amount := sendAuth.SpendLimit.AmountOfNoDenomValidation("token").Uint64()
	if amount != 0 && amount < minAmount {
		return nil, fmt.Errorf("authorization amount cannot be less than %d", minAmount)
	}

	expTimestamp := time.Now().Add(videoLinkValidTime).Unix()
	link := k.makeVideoLink(msg.Creator, msg.VideoId, expTimestamp)

	videoLink := types.VideoLink{
		Index: fmt.Sprintf("%s-%d", msg.Creator, msg.VideoId),
		Url:   link,
		Exp:   uint64(expTimestamp),
	}
	k.SetVideoLink(ctx, videoLink)

	ctx.EventManager().EmitEvent(sdk.NewEvent(types.TypeMsgPlayVideo, sdk.NewAttribute("url", videoLink.Url)))

	return &types.MsgPlayVideoResponse{
		Url: link,
	}, nil
}

func (k msgServer) makeVideoLink(creator string, videoId uint64, exp int64) string {
	server := servers[rand.Intn(len(servers))]

	cipher := dongle.NewCipher()
	cipher.SetMode(dongle.ECB)
	cipher.SetPadding(dongle.PKCS7)
	cipher.SetKey(server.aesKey)

	path := []byte(fmt.Sprintf("addr=%s,video_id=%d,exp=%d", creator, videoId, exp))
	path = dongle.Encrypt.FromBytes(path).ByAes(cipher).ToBase64Bytes()
	pathStr := url.PathEscape(string(path))

	return fmt.Sprintf("http://%s/%s/%d.m3u8", server.host, pathStr, videoId)
}
