package datumentitytype

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new datum entity type",
	Run: func(cmd *cobra.Command, args []string) {
		err := create(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(createCmd)

	// command line flags for the create command
	createCmd.Flags().StringP("name", "n", "", "name of the entity type")
}

// createValidation validates the required fields for the command
func createValidation() (input datumclient.CreateEntityTypeInput, err error) {
	// validation of required fields for the create command
	input.Name = datum.Config.String("name")
	if input.Name == "" {
		return input, datum.NewRequiredFieldMissingError("entity type name")
	}

	return input, nil
}

// create a new datum entityType
func create(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	input, err := createValidation()
	cobra.CheckErr(err)

	o, err := client.CreateEntityType(ctx, input)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
