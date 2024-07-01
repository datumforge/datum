package datumgroup

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new datum group",
	Run: func(cmd *cobra.Command, args []string) {
		err := create(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(createCmd)

	createCmd.Flags().StringP("name", "n", "", "name of the group")
	createCmd.Flags().StringP("display-name", "s", "", "display name of the group")
	createCmd.Flags().StringP("description", "d", "", "description of the group")
}

// createValidation validates the required fields for the command
func createValidation() (input datumclient.CreateGroupInput, err error) {
	input.Name = datum.Config.String("name")
	if input.Name == "" {
		return input, datum.NewRequiredFieldMissingError("group name")
	}

	displayName := datum.Config.String("display-name")
	if displayName != "" {
		input.DisplayName = &displayName
	}

	description := datum.Config.String("description")
	if description != "" {
		input.Description = &description
	}

	return input, nil
}

// create a new datum group
func create(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	input, err := createValidation()
	cobra.CheckErr(err)

	o, err := client.CreateGroup(ctx, input)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
