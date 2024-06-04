package datumorgsetting

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var orgSettingUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing datum org setting",
	RunE: func(cmd *cobra.Command, args []string) error {
		return updateOrganizationSetting(cmd.Context())
	},
}

func init() {
	orgSettingCmd.AddCommand(orgSettingUpdateCmd)

	orgSettingUpdateCmd.Flags().StringP("id", "i", "", "org setting id to update")
	orgSettingUpdateCmd.Flags().StringSliceP("domains", "d", []string{}, "domains associated with the org")
	orgSettingUpdateCmd.Flags().StringP("billing-contact", "c", "", "billing contact for the org")
	orgSettingUpdateCmd.Flags().StringP("billing-email", "e", "", "billing email for the org")
	orgSettingUpdateCmd.Flags().StringP("billing-phone", "p", "", "billing phone for the org")
	orgSettingUpdateCmd.Flags().StringP("billing-address", "a", "", "billing address for the org")
	orgSettingUpdateCmd.Flags().StringP("tax-identifier", "x", "", "tax identifier for the org")
	orgSettingUpdateCmd.Flags().StringSliceP("tags", "t", []string{}, "tags associated with the org")
}

func updateOrganizationSetting(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	if err != nil {
		return err
	}
	defer datum.StoreSessionCookies(client)

	var s []byte

	settingsID := datum.Config.String("id")
	if settingsID == "" {
		return datum.NewRequiredFieldMissingError("setting id")
	}

	input := datumclient.UpdateOrganizationSettingInput{}

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

	o, err := client.UpdateOrganizationSetting(ctx, settingsID, input)
	if err != nil {
		return err
	}

	s, err = json.Marshal(o)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
