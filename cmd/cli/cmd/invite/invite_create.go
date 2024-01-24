package datuminvite

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/internal/datumclient"
)

var inviteCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create invitation to join a datum organization",
	RunE: func(cmd *cobra.Command, args []string) error {
		return createInvite(cmd.Context())
	},
}

func init() {
	inviteCmd.AddCommand(inviteCreateCmd)

	inviteCreateCmd.Flags().StringP("org-id", "o", "", "org id")
	datum.ViperBindFlag("invite.create.orgid", inviteCreateCmd.Flags().Lookup("org-id"))

	inviteCreateCmd.Flags().StringP("email", "e", "", "destination email for the invitation")
	datum.ViperBindFlag("invite.create.email", inviteCreateCmd.Flags().Lookup("email"))
}

func createInvite(ctx context.Context) error {
	cli, err := datum.GetClient(ctx)
	if err != nil {
		return err
	}

	// save session cookies on function exit
	client, _ := cli.Client.(*datumclient.Client)
	defer datum.StoreSessionCookies(client)

	oID := viper.GetString("invite.create.orgid")
	if oID == "" {
		return datum.NewRequiredFieldMissingError("org id")
	}

	email := viper.GetString("invite.create.email")
	if email == "" {
		return datum.NewRequiredFieldMissingError("email")
	}

	input := datumclient.CreateInviteInput{
		OwnerID:   oID,
		Recipient: email,
	}

	var s []byte

	invite, err := cli.Client.CreateInvite(ctx, input, cli.Interceptor)
	if err != nil {
		return err
	}

	s, err = json.Marshal(invite)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
