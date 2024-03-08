package datumorgsetting

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/internal/datumclient"
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
	datum.ViperBindFlag("orgsetting.update.id", orgSettingUpdateCmd.Flags().Lookup("id"))

	orgSettingUpdateCmd.Flags().StringSliceP("domains", "d", []string{}, "domains associated with the org")
	datum.ViperBindFlag("orgsetting.update.domains", orgSettingUpdateCmd.Flags().Lookup("domains"))

	orgSettingUpdateCmd.Flags().StringP("billing-contact", "c", "", "billing contact for the org")
	datum.ViperBindFlag("orgsetting.update.billingcontact", orgSettingUpdateCmd.Flags().Lookup("billing-contact"))

	orgSettingUpdateCmd.Flags().StringP("billing-email", "e", "", "billing email for the org")
	datum.ViperBindFlag("orgsetting.update.billingemail", orgSettingUpdateCmd.Flags().Lookup("billing-email"))

	orgSettingUpdateCmd.Flags().StringP("billing-phone", "p", "", "billing phone for the org")
	datum.ViperBindFlag("orgsetting.update.billingphone", orgSettingUpdateCmd.Flags().Lookup("billing-phone"))

	orgSettingUpdateCmd.Flags().StringP("billing-address", "a", "", "billing address for the org")
	datum.ViperBindFlag("orgsetting.update.billingaddress", orgSettingUpdateCmd.Flags().Lookup("billing-address"))

	orgSettingUpdateCmd.Flags().StringP("tax-identifier", "x", "", "tax identifier for the org")
	datum.ViperBindFlag("orgsetting.update.taxidentifier", orgSettingUpdateCmd.Flags().Lookup("tax-identifier"))

	orgSettingUpdateCmd.Flags().StringSliceP("tags", "t", []string{}, "tags associated with the org")
	datum.ViperBindFlag("orgsetting.update.tags", orgSettingUpdateCmd.Flags().Lookup("tags"))
}

func updateOrganizationSetting(ctx context.Context) error {
	cli, err := datum.GetGraphClient(ctx)
	if err != nil {
		return err
	}

	client, _ := cli.Client.(*datumclient.Client)
	defer datum.StoreSessionCookies(client)

	var s []byte

	settingsID := viper.GetString("orgsetting.update.id")
	if settingsID == "" {
		return datum.NewRequiredFieldMissingError("setting id")
	}

	input := datumclient.UpdateOrganizationSettingInput{}

	billingContact := viper.GetString("orgsetting.update.billingcontact")
	if billingContact != "" {
		input.BillingContact = &billingContact
	}

	billingEmail := viper.GetString("orgsetting.update.billingemail")
	if billingEmail != "" {
		input.BillingContact = &billingEmail
	}

	billingPhone := viper.GetString("orgsetting.update.billingphone")
	if billingPhone != "" {
		input.BillingContact = &billingPhone
	}

	billingAddress := viper.GetString("orgsetting.update.billingaddress")
	if billingAddress != "" {
		input.BillingContact = &billingAddress
	}

	taxIdentifier := viper.GetString("orgsetting.update.taxidentifier")
	if taxIdentifier != "" {
		input.BillingContact = &taxIdentifier
	}

	tags := viper.GetStringSlice("orgsetting.update.tags")
	if len(tags) > 0 {
		input.Tags = tags
	}

	domains := viper.GetStringSlice("orgsetting.update.domains")
	if len(domains) > 0 {
		input.Domains = domains
	}

	o, err := cli.Client.UpdateOrganizationSetting(ctx, settingsID, input, cli.Interceptor)
	if err != nil {
		return err
	}

	s, err = json.Marshal(o)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
