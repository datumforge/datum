package datumsubscribers

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "remove a subscriber from a datum organization",
	Run: func(cmd *cobra.Command, args []string) {
		err := delete(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringP("email", "e", "", "email address of the subscriber to delete")
}

// deleteValidation validates the required fields for the command
func deleteValidation() (string, error) {
	email := datum.Config.String("email")
	if email == "" {
		return "", datum.NewRequiredFieldMissingError("email")
	}

	return email, nil
}

// delete an existing organization subscriber
func delete(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	email, err := deleteValidation()
	cobra.CheckErr(err)

	o, err := client.DeleteSubscriber(ctx, email)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
