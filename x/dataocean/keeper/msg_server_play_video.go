package keeper

import (
	"context"
	crand "crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
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
		host:   "18.141.197.172:9001",
		aesKey: "key_for_server_1",
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

	privateKey, publicKey, err := k.genRsaKey()
	if err != nil {
		return nil, fmt.Errorf("failed to generate rsa key: %s", err.Error())
	}

	videoLink := types.VideoLink{
		Index: fmt.Sprintf("%s-%d", msg.Creator, msg.VideoId),
		Url:   link,
		Exp:   uint64(expTimestamp),
	}
	k.SetVideoLink(ctx, videoLink)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(types.TypeMsgPlayVideo, sdk.NewAttribute("url", videoLink.Url)),
		sdk.NewEvent(types.TypeMsgPlayVideo, sdk.NewAttribute("payPrivateKey", privateKey)),
		sdk.NewEvent(types.TypeMsgPlayVideo, sdk.NewAttribute("payPublicKey", publicKey)),
	})

	return &types.MsgPlayVideoResponse{
		Url:           link,
		Exp:           uint64(expTimestamp),
		PayPrivateKey: privateKey,
		PayPublicKey:  publicKey,
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

func (k msgServer) genRsaKey() (string, string, error) {
	privateKey, err := rsa.GenerateKey(crand.Reader, 2048)
	if err != nil {
		return "", "", err
	}
	publicKey := &privateKey.PublicKey

	privateBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return "", "", err
	}
	publicBlock := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: derPkix,
	}

	privateKeyStr := string(pem.EncodeToMemory(privateBlock))
	publicKeyStr := string(pem.EncodeToMemory(publicBlock))

	return privateKeyStr, publicKeyStr, nil
}
