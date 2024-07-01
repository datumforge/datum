package datumentitlementplanfeatures

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new datum entitlement plan feature",
	Run: func(cmd *cobra.Command, args []string) {
		err := create(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(createCmd)

	createCmd.Flags().StringP("plan-id", "p", "", "plan id to use for the plan feature")
	createCmd.Flags().StringP("feature-id", "f", "", "feature id to use for the plan feature")

	// fields for the plan feature metadata
	createCmd.Flags().StringP("limit-type", "t", "", "limit type for the plan feature, e.g. requests, storage, etc.")
	createCmd.Flags().Int64P("limit", "l", 0, "limit value for the plan feature")

	createCmd.Flags().StringSlice("tags", []string{}, "tags associated with the plan")
}

// createValidation validates the required fields for the command
func createValidation() (input datumclient.CreateEntitlementPlanFeatureInput, err error) {
	planID := datum.Config.String("plan-id")
	if planID == "" {
		return input, datum.NewRequiredFieldMissingError("plan id")
	}

	featureID := datum.Config.String("feature-id")
	if featureID == "" {
		return input, datum.NewRequiredFieldMissingError("feature id")
	}

	input = datumclient.CreateEntitlementPlanFeatureInput{
		PlanID:    planID,
		FeatureID: featureID,
	}

	limitType := datum.Config.String("limit-type")
	limit := datum.Config.Int64("limit")

	if limitType == "" && limit > 0 {
		return input, datum.NewRequiredFieldMissingError("limit type")
	}

	if limitType != "" && limit == 0 {
		return input, datum.NewRequiredFieldMissingError("limit")
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

	return input, nil
}

// create a new plan feature in the datum platform
func create(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	input, err := createValidation()
	cobra.CheckErr(err)

	f, err := client.CreateEntitlementPlanFeature(ctx, input)
	cobra.CheckErr(err)

	return consoleOutput(f)
}
