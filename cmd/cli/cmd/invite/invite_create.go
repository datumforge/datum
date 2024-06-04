package datuminvite

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"

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
	inviteCreateCmd.Flags().StringP("role", "r", "member", "role for the user in the organization (admin, member)")
}

func createInvite(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	if err != nil {
		return err
	}
	defer datum.StoreSessionCookies(client)

	email := datum.Config.String("email")
	if email == "" {
		return datum.NewRequiredFieldMissingError("email")
	}

	role := enums.ToRole(datum.Config.String("role"))

	input := datumclient.CreateInviteInput{
		Recipient: email,
		Role:      role,
	}

	var s []byte

	invite, err := client.CreateInvite(ctx, input)
	if err != nil {
		return err
	}

	s, err = json.Marshal(invite)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
