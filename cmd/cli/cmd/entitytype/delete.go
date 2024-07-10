package datumentitytype

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete an existing datum entity type",
	Run: func(cmd *cobra.Command, args []string) {
		err := delete(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringP("id", "i", "", "entity type id to delete")
}

// deleteValidation validates the required fields for the command
func deleteValidation() (string, error) {
	id := datum.Config.String("id")
	if id == "" {
		return "", datum.NewRequiredFieldMissingError("entity type id")
	}

	return id, nil
}

// delete an existing entityType in the datum platform
func delete(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	id, err := deleteValidation()
	cobra.CheckErr(err)

	o, err := client.DeleteEntityType(ctx, id)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
