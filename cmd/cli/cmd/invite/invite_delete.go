package datuminvite

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"

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
}

func deleteInvite(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	if err != nil {
		return err
	}
	defer datum.StoreSessionCookies(client)

	invID := datum.Config.String("invite-id")
	if invID == "" {
		return datum.NewRequiredFieldMissingError("invite id")
	}

	var s []byte

	invite, err := client.DeleteInvite(ctx, invID)
	if err != nil {
		return err
	}

	s, err = json.Marshal(invite)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
