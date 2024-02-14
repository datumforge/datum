package datuminvite

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/internal/datumclient"
	"github.com/datumforge/datum/internal/httpserve/handlers"
)

var inviteAcceptCmd = &cobra.Command{
	Use:   "accept",
	Short: "accept an invite",
	RunE: func(cmd *cobra.Command, args []string) error {
		return inviteAccept(cmd.Context())
	},
}

func init() {
	inviteCmd.AddCommand(inviteAcceptCmd)

	inviteAcceptCmd.Flags().StringP("token", "t", "", "invite token")
	datum.ViperBindFlag("invite.accept.token", inviteAcceptCmd.Flags().Lookup("token"))
}

func inviteAccept(ctx context.Context) error {
	var s []byte

	token := viper.GetString("invite.accept.token")
	if token == "" {
		return datum.NewRequiredFieldMissingError("token")
	}

	invite := handlers.Invite{
		Token: token,
	}

	// new client with params
	cli, err := datum.GetRestClient(ctx)
	if err != nil {
		return err
	}

	// this allows the use of the graph client to be used for the REST endpoints
	dc := cli.Client.(*datumclient.Client)

	defer datum.StoreSessionCookies(dc)

	registration, tokens, err := datumclient.OrgInvite(dc, ctx, invite, cli.AccessToken)
	if err != nil {
		return err
	}

	if err := datum.StoreToken(tokens); err != nil {
		return err
	}

	s, err = json.Marshal(registration)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
