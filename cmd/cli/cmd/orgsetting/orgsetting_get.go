package datumorgsetting

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

var orgSettingGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get org settings",
	RunE: func(cmd *cobra.Command, args []string) error {
		return orgSettings(cmd.Context())
	},
}

func init() {
	orgSettingCmd.AddCommand(orgSettingGetCmd)

	orgSettingGetCmd.Flags().StringP("id", "i", "", "org setting id to retrieve")
	datum.ViperBindFlag("orgsetting.get.id", orgSettingGetCmd.Flags().Lookup("id"))
}

func orgSettings(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	if err != nil {
		return err
	}
	defer datum.StoreSessionCookies(client)

	// filter options
	settingsID := viper.GetString("orgsetting.get.id")

	var s []byte

	// if setting ID is not provided, get settings which will automatically filter by org id
	if settingsID == "" {
		settings, err := client.GetOrganizationSettings(ctx)
		if err != nil {
			return err
		}

		s, err = json.Marshal(settings)
		if err != nil {
			return err
		}
	} else {
		org, err := client.GetOrganizationSettingByID(ctx, settingsID)
		if err != nil {
			return err
		}

		s, err = json.Marshal(org)
		if err != nil {
			return err
		}
	}

	return datum.JSONPrint(s)
}
