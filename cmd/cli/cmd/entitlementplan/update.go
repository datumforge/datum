package datumentitlementplan

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update an existing datum entitlement plan",
	Run: func(cmd *cobra.Command, args []string) {
		err := update(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(updateCmd)

	updateCmd.Flags().StringP("id", "i", "", "plan id to update")
	updateCmd.Flags().String("display-name", "", "human friendly name of the plan")
	updateCmd.Flags().StringP("description", "d", "", "description of the plan")
	updateCmd.Flags().StringSlice("tags", []string{}, "tags associated with the plan")
}

// updateValidation validates the required fields for the command
func updateValidation() (id string, input datumclient.UpdateEntitlementPlanInput, err error) {
	id = datum.Config.String("id")
	if id == "" {
		return id, input, datum.NewRequiredFieldMissingError("plan id")
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

	return id, input, nil
}

// update an existing datum entitlement plan
func update(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	id, input, err := updateValidation()
	cobra.CheckErr(err)

	p, err := client.UpdateEntitlementPlan(ctx, id, input)
	cobra.CheckErr(err)

	return consoleOutput(p)
}
