package datuminvite

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/internal/datumclient"
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
	cli, err := datum.GetGraphClient(ctx)
	if err != nil {
		return err
	}

	// save session cookies on function exit
	client, _ := cli.Client.(*datumclient.Client)
	defer datum.StoreSessionCookies(client)

	// filter options
	invID := viper.GetString("invite.get.id")
	if invID == "" {
		return datum.NewRequiredFieldMissingError("invite id")
	}

	var s []byte

	invite, err := cli.Client.GetInvite(ctx, invID, cli.Interceptor)
	if err != nil {
		return err
	}

	s, err = json.Marshal(invite)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
