package datuminvite

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete an invitation to join an organization",
	Run: func(cmd *cobra.Command, args []string) {
		err := delete(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringP("id", "i", "", "id of the invitation to delete")
}

// deleteValidation validates the required fields for the command
func deleteValidation() (string, error) {
	id := datum.Config.String("id")
	if id == "" {
		return "", datum.NewRequiredFieldMissingError("id")
	}

	return id, nil
}

// delete an invitation to join an organization
func delete(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	id, err := deleteValidation()
	cobra.CheckErr(err)

	o, err := client.DeleteInvite(ctx, id)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
