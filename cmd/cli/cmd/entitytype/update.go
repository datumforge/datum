package datumentitytype

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update an existing datum entityType",
	Run: func(cmd *cobra.Command, args []string) {
		err := update(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(updateCmd)

	updateCmd.Flags().StringP("id", "i", "", "entityType id to update")

	// command line flags for the update command
	updateCmd.Flags().StringP("name", "n", "", "name of the entityType")
}

// updateValidation validates the required fields for the command
func updateValidation() (id string, input datumclient.UpdateEntityTypeInput, err error) {
	id = datum.Config.String("id")
	if id == "" {
		return id, input, datum.NewRequiredFieldMissingError("entityType id")
	}

	// validation of required fields for the update command
	name := datum.Config.String("name")
	if name != "" {
		input.Name = &name
	}

	return id, input, nil
}

// update an existing entityType in the datum platform
func update(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	id, input, err := updateValidation()
	cobra.CheckErr(err)

	o, err := client.UpdateEntityType(ctx, id, input)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
