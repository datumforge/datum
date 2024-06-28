package datumentitlement

import (
	"context"
	"time"

	"github.com/samber/lo"
	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new datum entitlement",
	Run: func(cmd *cobra.Command, args []string) {
		err := create(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(createCmd)

	createCmd.Flags().StringP("organization_id", "o", "", "organization id associated with the entitlement")
	createCmd.Flags().StringP("plan_id", "p", "", "plan_id associated with the entitlement")
	createCmd.Flags().DurationP("expires_at", "e", 0, "expiration time of the entitlement")
	createCmd.Flags().Bool("cancelled", false, "whether the entitlement is cancelled")
	createCmd.Flags().StringP("external_customer_id", "c", "", "external customer id")
	createCmd.Flags().StringP("external_subscription_id", "s", "", "external subscription id")
	createCmd.Flags().StringSlice("tags", []string{}, "tags associated with the entitlement")
}

// createValidation validates the required fields for the command
func createValidation() (input datumclient.CreateEntitlementInput, err error) {
	orgID := datum.Config.String("organization_id")
	if orgID == "" {
		return input, datum.NewRequiredFieldMissingError("organization id")
	}

	planID := datum.Config.String("plan_id")
	if planID == "" {
		return input, datum.NewRequiredFieldMissingError("plan id")
	}

	cancelled := datum.Config.Bool("cancelled")
	externalCustomerID := datum.Config.String("external_customer_id")
	externalSubscriptionID := datum.Config.String("external_subscription_id")
	expiresAt := datum.Config.Duration("expires_at")
	tags := datum.Config.Strings("tags")

	input = datumclient.CreateEntitlementInput{
		OrganizationID: orgID,
		PlanID:         planID,
		Cancelled:      &cancelled,
	}

	if externalCustomerID != "" {
		input.ExternalCustomerID = &externalCustomerID
	}

	if externalSubscriptionID != "" {
		input.ExternalSubscriptionID = &externalSubscriptionID
	}

	if expiresAt != 0 {
		input.ExpiresAt = lo.ToPtr(time.Now().Add(expiresAt))
	}

	if len(tags) > 0 {
		input.Tags = tags
	}

	return input, nil
}

// create a new entitlement in the datum platform
func create(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	input, err := createValidation()
	cobra.CheckErr(err)

	o, err := client.CreateEntitlement(ctx, input)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
