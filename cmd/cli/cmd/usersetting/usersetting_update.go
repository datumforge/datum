package datumusersetting

import (
	"context"
	"encoding/json"
	"time"

	"github.com/samber/lo"
	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/pkg/datumclient"
)

var userSettingUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing datum user setting",
	RunE: func(cmd *cobra.Command, args []string) error {
		return updateUserSetting(cmd.Context())
	},
}

func init() {
	userSettingCmd.AddCommand(userSettingUpdateCmd)

	userSettingUpdateCmd.Flags().StringP("id", "i", "", "user setting id to update")
	userSettingUpdateCmd.Flags().String("status", "", "status of the user - active, inactive, suspended")
	userSettingUpdateCmd.Flags().StringP("default-org", "o", "", "default organization id")
	userSettingUpdateCmd.Flags().StringSliceP("tags", "t", []string{}, "tags associated with the user")
	userSettingUpdateCmd.Flags().BoolP("silence-notifications", "s", false, "silence notifications from datum")
	userSettingUpdateCmd.Flags().Bool("enable-2fa", false, "enable 2fa authentication")
	userSettingUpdateCmd.Flags().Bool("disable-2fa", false, "disable 2fa authentication")
}

func updateUserSetting(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	var s []byte

	settingsID := datum.Config.String("id")

	// initialize the input
	input := datumclient.UpdateUserSettingInput{}

	// set default org if provided
	defaultOrg := datum.Config.String("default-org")
	if defaultOrg != "" {
		input.DefaultOrgID = &defaultOrg
	}

	// set silenced at time if silence flag is set
	if datum.Config.Bool("silence") {
		now := time.Now().UTC()
		input.SilencedAt = &now
	} else {
		input.SilencedAt = nil
	}

	// explicitly set 2fa if provided to avoid wiping setting unintentionally
	if datum.Config.Bool("enable-2fa") {
		input.IsTfaEnabled = lo.ToPtr(true)
	} else if datum.Config.Bool("disable-2fa") {
		input.IsTfaEnabled = lo.ToPtr(false)
	}

	// add tags to the input if provided
	tags := datum.Config.Strings("tags")
	if len(tags) > 0 {
		input.Tags = tags
	}

	// add status to the input if provided
	status := datum.Config.String("status")
	if status != "" {
		input.Status = enums.ToUserStatus(status)
	}

	if settingsID == "" {
		// get the user settings id
		settings, err := client.GetUserSettings(ctx)
		cobra.CheckErr(err)

		// this should never happen, but just in case
		if len(settings.GetUserSettings().Edges) == 0 {
			return datum.ErrNotFound
		}

		settingsID = settings.GetUserSettings().Edges[0].Node.ID
	}

	// update the user settings
	o, err := client.UpdateUserSetting(ctx, settingsID, input)
	cobra.CheckErr(err)

	// parse the output
	s, err = json.Marshal(o)
	cobra.CheckErr(err)

	return datum.JSONPrint(s)
}
