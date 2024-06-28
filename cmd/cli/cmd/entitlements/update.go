package datumentitlement

import (
	"context"
	"time"

	"github.com/samber/lo"
	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update an existing datum entitlement",
	Long: `The entitlement update command updates an existing entitlement in the datum platform.
The command requires the entitlement id associated with the entitlement and the fields to update.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := update(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(updateCmd)

	updateCmd.Flags().StringP("id", "i", "", "entitlement id to update")
	updateCmd.Flags().DurationP("expires_at", "e", 0, "expiration time of the entitlement")
	updateCmd.Flags().Bool("cancelled", false, "whether the entitlement is cancelled")
	updateCmd.Flags().StringP("external_customer_id", "c", "", "external customer id")
	updateCmd.Flags().StringP("external_subscription_id", "s", "", "external subscription id")
}

// updateValidation validates the required fields for the command
func updateValidation() (id string, input datumclient.UpdateEntitlementInput, err error) {
	id = datum.Config.String("id")
	if id == "" {
		return id, input, datum.NewRequiredFieldMissingError("entitlement id")
	}

	cancelled := datum.Config.Bool("cancelled")
	externalCustomerID := datum.Config.String("external_customer_id")
	externalSubscriptionID := datum.Config.String("external_subscription_id")
	expiresAt := datum.Config.Duration("expires_at")

	input.Cancelled = &cancelled

	if externalCustomerID != "" {
		input.ExternalCustomerID = &externalCustomerID
	}

	if externalSubscriptionID != "" {
		input.ExternalSubscriptionID = &externalSubscriptionID
	}

	if expiresAt != 0 {
		input.ExpiresAt = lo.ToPtr(time.Now().Add(expiresAt))
	}

	return id, input, nil
}

// update an existing datum entitlement
func update(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	id, input, err := updateValidation()
	cobra.CheckErr(err)

	o, err := client.UpdateEntitlement(ctx, id, input)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
