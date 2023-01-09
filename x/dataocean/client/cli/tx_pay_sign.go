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

func CmdPaySign() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pay-sign [video-id] [pay-public-key]",
		Short: "Broadcast message pay-sign",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argVideoId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argPayPublicKey := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgPaySign(
				clientCtx.GetFromAddress().String(),
				argVideoId,
				argPayPublicKey,
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
