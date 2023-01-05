package simulation

import (
	"math/rand"

	"dataocean/x/dataocean/keeper"
	"dataocean/x/dataocean/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgSubmitPaySign(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgSubmitPaySign{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the SubmitPaySign simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "SubmitPaySign simulation not implemented"), nil, nil
	}
}
