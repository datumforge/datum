package datumorg

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update an existing datum organization",
	Run: func(cmd *cobra.Command, args []string) {
		err := update(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(updateCmd)

	updateCmd.Flags().StringP("id", "i", "", "org id to update")
	updateCmd.Flags().StringP("name", "n", "", "name of the organization")
	updateCmd.Flags().StringP("display-name", "s", "", "display name of the organization")
	updateCmd.Flags().StringP("description", "d", "", "description of the organization")
}

// updateValidation validates the required fields for the command
func updateValidation() (id string, input datumclient.UpdateOrganizationInput, err error) {
	id = datum.Config.String("id")
	if id == "" {
		return id, input, datum.NewRequiredFieldMissingError("organization id")
	}

	name := datum.Config.String("name")
	if name != "" {
		input.Name = &name
	}

	displayName := datum.Config.String("display-name")
	if displayName != "" {
		input.DisplayName = &displayName
	}

	description := datum.Config.String("description")
	if description != "" {
		input.Description = &description
	}

	return id, input, nil
}

// update an existing organization in the datum platform
func update(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	id, input, err := updateValidation()
	cobra.CheckErr(err)

	o, err := client.UpdateOrganization(ctx, id, input)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
