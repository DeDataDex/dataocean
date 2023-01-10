package cli

import (
	"strconv"

	"dataocean/x/dataocean/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSubmitPaySign() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-pay-sign [pay-sign] [pay-data]",
		Short: "Broadcast message submit-pay-sign",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPaySign := args[0]
			argPayData := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSubmitPaySign(
				clientCtx.GetFromAddress().String(),
				argPaySign,
				argPayData,
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
