package datumentitlementplanfeatures

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update an existing datum entitlement plan feature",
	Run: func(cmd *cobra.Command, args []string) {
		err := update(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(updateCmd)

	updateCmd.Flags().StringP("id", "i", "", "feature id to update")

	// fields for the plan feature metadata
	updateCmd.Flags().StringP("limit-type", "t", "", "limit type for the plan feature, e.g. requests, storage, etc.")
	updateCmd.Flags().Int64P("limit", "l", 0, "limit value for the plan feature")

	updateCmd.Flags().StringSlice("tags", []string{}, "tags associated with the plan")
}

// updateValidation validates the required fields for the command
func updateValidation() (id string, input datumclient.UpdateEntitlementPlanFeatureInput, err error) {
	id = datum.Config.String("id")
	if id == "" {
		return id, input, datum.NewRequiredFieldMissingError("plan feature id")
	}

	limitType := datum.Config.String("limit-type")
	limit := datum.Config.Int64("limit")

	if limitType == "" && limit > 0 {
		return id, input, datum.NewRequiredFieldMissingError("limit type")
	}

	if limitType != "" && limit == 0 {
		return id, input, datum.NewRequiredFieldMissingError("limit")
	}

	if limitType != "" && limit > 0 {
		input.Metadata = map[string]interface{}{
			"limit_type": limitType,
			"limit":      limit,
		}
	}

	tags := datum.Config.Strings("tags")
	if len(tags) > 0 {
		input.Tags = tags
	}

	return id, input, nil
}

// update an existing plan feature in the datum platform
func update(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	id, input, err := updateValidation()
	cobra.CheckErr(err)

	f, err := client.UpdateEntitlementPlanFeature(ctx, id, input)
	cobra.CheckErr(err)

	return consoleOutput(f)
}
