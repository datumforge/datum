package datumorgsetting

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"

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
}

func orgSettings(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	// filter options
	settingsID := datum.Config.String("id")

	var s []byte

	// if setting ID is not provided, get settings which will automatically filter by org id
	if settingsID == "" {
		settings, err := client.GetOrganizationSettings(ctx)
		cobra.CheckErr(err)

		s, err = json.Marshal(settings)
		cobra.CheckErr(err)
	} else {
		org, err := client.GetOrganizationSettingByID(ctx, settingsID)
		cobra.CheckErr(err)

		s, err = json.Marshal(org)
		cobra.CheckErr(err)
	}

	return datum.JSONPrint(s)
}
