package datumentitlementplan

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new datum entitlement plan",
	Run: func(cmd *cobra.Command, args []string) {
		err := create(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(createCmd)

	createCmd.Flags().StringP("name", "n", "", "short name of the plan, must be unique")
	createCmd.Flags().String("display-name", "", "human friendly name of the plan")
	createCmd.Flags().StringP("description", "d", "", "description of the plan")
	createCmd.Flags().StringP("version", "v", "", "version of the plan")
	createCmd.Flags().StringSlice("tags", []string{}, "tags associated with the plan")
}

// createValidation validates the required fields for the command
func createValidation() (input datumclient.CreateEntitlementPlanInput, err error) {
	name := datum.Config.String("name")
	if name == "" {
		return input, datum.NewRequiredFieldMissingError("plan name")
	}

	version := datum.Config.String("version")
	if version == "" {
		return input, datum.NewRequiredFieldMissingError("version")
	}

	input = datumclient.CreateEntitlementPlanInput{
		Name:    name,
		Version: version,
	}

	description := datum.Config.String("description")
	if description != "" {
		input.Description = &description
	}

	displayName := datum.Config.String("display-name")
	if displayName != "" {
		input.DisplayName = &displayName
	}

	tags := datum.Config.Strings("tags")
	if len(tags) > 0 {
		input.Tags = tags
	}

	return input, nil
}

// create a plan in the datum platform
func create(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	input, err := createValidation()
	cobra.CheckErr(err)

	p, err := client.CreateEntitlementPlan(ctx, input)
	cobra.CheckErr(err)

	return consoleOutput(p)
}
