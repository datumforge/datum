package datuminvite

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/utils/cli/tables"
)

var inviteGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get organization invitation",
	RunE: func(cmd *cobra.Command, args []string) error {
		return invites(cmd.Context())
	},
}

func init() {
	inviteCmd.AddCommand(inviteGetCmd)

	inviteGetCmd.Flags().StringP("invite-id", "i", "", "invite id to query")
	datum.ViperBindFlag("invite.get.id", inviteGetCmd.Flags().Lookup("invite-id"))
}

func invites(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	if err != nil {
		return err
	}
	defer datum.StoreSessionCookies(client)

	// filter options
	invID := viper.GetString("invite.get.id")

	if invID != "" {
		invite, err := client.GetInvite(ctx, invID, client.Config().Interceptors...)
		if err != nil {
			return err
		}

		return printInvite(invite)
	}

	invites, err := client.GetInvites(ctx, client.Config().Interceptors...)
	if err != nil {
		return err
	}

	return printInvite(invites)
}

// printInviteTable prints the invite table to the console
func printInviteTable(i interface{}) {
	writer := tables.NewTableWriter(inviteCmd.OutOrStdout(), "ID", "Recipient", "Role", "Status")

	switch v := i.(type) {
	case *datumclient.GetInvites:
		for _, invite := range v.Invites.Edges {
			writer.AddRow(invite.Node.ID, invite.Node.Recipient, invite.Node.Role, invite.Node.Status)
		}
	case *datumclient.GetInvite:
		writer.AddRow(v.Invite.ID, v.Invite.Recipient, v.Invite.Role, v.Invite.Status)
	}

	writer.Render()
}

// printInvite prints the invite to the console either in table or json format
func printInvite(i interface{}) error {
	if datum.OutputFormat == datum.JSONOutput {
		s, err := json.Marshal(i)
		if err != nil {
			return err
		}

		return datum.JSONPrint(s)
	}

	printInviteTable(i)

	return nil
}
