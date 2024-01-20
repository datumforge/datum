package datuminvite

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

var inviteDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove an invitation from an organization",
	RunE: func(cmd *cobra.Command, args []string) error {
		return deleteInvite(cmd.Context())
	},
}

func init() {
	inviteCmd.AddCommand(inviteDeleteCmd)

	inviteDeleteCmd.Flags().StringP("invite-id", "i", "", "invite id")
	datum.ViperBindFlag("invite.delete.inviteid", inviteDeleteCmd.Flags().Lookup("invite-id"))
}

func deleteInvite(ctx context.Context) error {
	// setup datum http client
	cli, err := datum.GetClient(ctx)
	if err != nil {
		return err
	}

	invID := viper.GetString("invite.delete.inviteid")
	if invID == "" {
		return datum.NewRequiredFieldMissingError("invite id")
	}

	var s []byte

	invite, err := cli.Client.DeleteInvite(ctx, invID, cli.Interceptor)
	if err != nil {
		return err
	}

	s, err = json.Marshal(invite)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
