package datumorg

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new datum organization",
	Run: func(cmd *cobra.Command, args []string) {
		err := create(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(createCmd)

	createCmd.Flags().StringP("name", "n", "", "name of the organization")
	createCmd.Flags().StringP("display-name", "s", "", "display name of the organization")
	createCmd.Flags().StringP("description", "d", "", "description of the organization")
	createCmd.Flags().StringP("parent-org-id", "p", "", "parent organization id, leave empty to create a root org")

	// TODO: https://github.com/datumforge/datum/issues/734
	// remove flag once the feature is implemented
	createCmd.Flags().BoolP("dedicated-db", "D", false, "create a dedicated database for the organization")
}

// createValidation validates the required fields for the command
func createValidation() (input datumclient.CreateOrganizationInput, err error) {
	name := datum.Config.String("name")
	if name != "" {
		input.Name = &name
	}
	if input.Name == nil {
		return input, datum.NewRequiredFieldMissingError("organization name")
	}

	displayName := datum.Config.String("display-name")
	if displayName != "" {
		input.DisplayName = &displayName
	}

	description := datum.Config.String("description")
	if description != "" {
		input.Description = &description
	}

	parentOrgID := datum.Config.String("parent-org-id")
	if parentOrgID != "" {
		input.ParentID = &parentOrgID
	}

	dedicatedDB := datum.Config.Bool("dedicated-db")
	if dedicatedDB {
		input.DedicatedDb = &dedicatedDB
	}

	return input, nil
}

// create an organization in the datum platform
func create(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	input, err := createValidation()
	cobra.CheckErr(err)

	o, err := client.CreateOrganization(ctx, input)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
