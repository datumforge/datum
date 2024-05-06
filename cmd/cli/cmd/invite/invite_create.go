package datuminvite

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/pkg/datumclient"
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

	inviteCreateCmd.Flags().StringP("email", "e", "", "destination email for the invitation")
	datum.ViperBindFlag("invite.create.email", inviteCreateCmd.Flags().Lookup("email"))

	inviteCreateCmd.Flags().StringP("role", "r", "member", "role for the user in the organization (admin, member)")
	datum.ViperBindFlag("invite.create.role", inviteCreateCmd.Flags().Lookup("role"))
}

func createInvite(ctx context.Context) error {
	cli, err := datum.GetGraphClient(ctx)
	if err != nil {
		return err
	}

	// save session cookies on function exit
	client, _ := cli.Client.(*datumclient.Client)
	defer datum.StoreSessionCookies(client)

	email := viper.GetString("invite.create.email")
	if email == "" {
		return datum.NewRequiredFieldMissingError("email")
	}

	role := enums.ToRole(viper.GetString("invite.create.role"))

	input := datumclient.CreateInviteInput{
		Recipient: email,
		Role:      role,
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
