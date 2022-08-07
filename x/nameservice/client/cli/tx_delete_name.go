package cli

import (
    "strconv"
	
	"github.com/spf13/cobra"
    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"nameservice/x/nameservice/types"
)

var _ = strconv.Itoa(0)

func CmdDeleteName() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-name [name] [value]",
		Short: "Broadcast message delete-name",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
      		 argName := args[0]
             argValue := args[1]
            
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteName(
				clientCtx.GetFromAddress().String(),
				argName,
				argValue,
				
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