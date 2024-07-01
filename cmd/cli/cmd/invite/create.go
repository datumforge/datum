package datuminvite

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/enums"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create and invitation to join a datum organization",
	Run: func(cmd *cobra.Command, args []string) {
		err := create(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(createCmd)

	createCmd.Flags().StringP("email", "e", "", "destination email for the invitation")
	createCmd.Flags().StringP("role", "r", "member", "role for the user in the organization (admin, member)")
}

// createValidation validates the required fields for the command
func createValidation() (input datumclient.CreateInviteInput, err error) {
	input.Recipient = datum.Config.String("email")
	if input.Recipient == "" {
		return input, datum.NewRequiredFieldMissingError("email")
	}

	input.Role = enums.ToRole(datum.Config.String("role"))

	return input, nil
}

// create an invitation in the datum platform
func create(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	input, err := createValidation()
	cobra.CheckErr(err)

	o, err := client.CreateInvite(ctx, input)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
