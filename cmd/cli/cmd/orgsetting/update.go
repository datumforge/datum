package datumorgsetting

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update an existing datum organization setting",
	Run: func(cmd *cobra.Command, args []string) {
		err := update(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(updateCmd)

	updateCmd.Flags().StringP("id", "i", "", "org setting id to update")
	updateCmd.Flags().StringSliceP("domains", "d", []string{}, "domains associated with the org")
	updateCmd.Flags().StringP("billing-contact", "c", "", "billing contact for the org")
	updateCmd.Flags().StringP("billing-email", "e", "", "billing email for the org")
	updateCmd.Flags().StringP("billing-phone", "p", "", "billing phone for the org")
	updateCmd.Flags().StringP("billing-address", "a", "", "billing address for the org")
	updateCmd.Flags().StringP("tax-identifier", "x", "", "tax identifier for the org")
	updateCmd.Flags().StringSliceP("tags", "t", []string{}, "tags associated with the org")
}

// updateValidation validates the input flags provided by the user
func updateValidation() (id string, input datumclient.UpdateOrganizationSettingInput, err error) {
	id = datum.Config.String("id")
	if id == "" {
		return id, input, datum.NewRequiredFieldMissingError("setting id")
	}

	billingContact := datum.Config.String("billingcontact")
	if billingContact != "" {
		input.BillingContact = &billingContact
	}

	billingEmail := datum.Config.String("billingemail")
	if billingEmail != "" {
		input.BillingContact = &billingEmail
	}

	billingPhone := datum.Config.String("billingphone")
	if billingPhone != "" {
		input.BillingContact = &billingPhone
	}

	billingAddress := datum.Config.String("billingaddress")
	if billingAddress != "" {
		input.BillingContact = &billingAddress
	}

	taxIdentifier := datum.Config.String("taxidentifier")
	if taxIdentifier != "" {
		input.BillingContact = &taxIdentifier
	}

	tags := datum.Config.Strings("tags")
	if len(tags) > 0 {
		input.Tags = tags
	}

	domains := datum.Config.Strings("domains")
	if len(domains) > 0 {
		input.Domains = domains
	}

	return id, input, nil
}

// update an organization setting in the datum platform
func update(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	id, input, err := updateValidation()
	cobra.CheckErr(err)

	o, err := client.UpdateOrganizationSetting(ctx, id, input)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
