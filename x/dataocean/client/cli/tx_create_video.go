package cli

import (
	"strconv"

	"dataocean/x/dataocean/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateVideo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-video [title] [description] [cover-link] [video-link] [price-mb]",
		Short: "Broadcast message create-video",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTitle := args[0]
			argDescription := args[1]
			argCoverLink := args[2]
			argVideoLink := args[3]
			argPriceMB, err := cast.ToUint64E(args[4])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateVideo(
				clientCtx.GetFromAddress().String(),
				argTitle,
				argDescription,
				argCoverLink,
				argVideoLink,
				argPriceMB,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
