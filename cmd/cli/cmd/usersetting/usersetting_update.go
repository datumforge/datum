package datumusersetting

import (
	"context"
	"encoding/json"
	"time"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

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
	datum.ViperBindFlag("usersetting.update.id", userSettingUpdateCmd.Flags().Lookup("id"))

	userSettingUpdateCmd.Flags().String("status", "", "status of the user - active, inactive, suspended")
	datum.ViperBindFlag("usersetting.update.status", userSettingUpdateCmd.Flags().Lookup("status"))

	userSettingUpdateCmd.Flags().StringP("default-org", "o", "", "default organization id")
	datum.ViperBindFlag("usersetting.update.defaultorg", userSettingUpdateCmd.Flags().Lookup("default-org"))

	userSettingUpdateCmd.Flags().StringSliceP("tags", "t", []string{}, "tags associated with the user")
	datum.ViperBindFlag("usersetting.update.tags", userSettingUpdateCmd.Flags().Lookup("tags"))

	userSettingUpdateCmd.Flags().BoolP("silence-notifications", "s", false, "silence notifications from datum")
	datum.ViperBindFlag("usersetting.update.silence", userSettingUpdateCmd.Flags().Lookup("silence-notifications"))

	userSettingUpdateCmd.Flags().Bool("enable-2fa", false, "enable 2fa authentication")
	datum.ViperBindFlag("usersetting.update.enable-2fa", userSettingUpdateCmd.Flags().Lookup("enable-2fa"))

	userSettingUpdateCmd.Flags().Bool("disable-2fa", false, "disable 2fa authentication")
	datum.ViperBindFlag("usersetting.update.disable-2fa", userSettingUpdateCmd.Flags().Lookup("disable-2fa"))
}

func updateUserSetting(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	if err != nil {
		return err
	}
	defer datum.StoreSessionCookies(client)

	var s []byte

	settingsID := viper.GetString("usersetting.update.id")

	// initialize the input
	input := datumclient.UpdateUserSettingInput{}

	// set default org if provided
	defaultOrg := viper.GetString("usersetting.update.defaultorg")
	if defaultOrg != "" {
		input.DefaultOrgID = &defaultOrg
	}

	// set silenced at time if silence flag is set
	if viper.GetBool("usersetting.update.silence") {
		now := time.Now().UTC()
		input.SilencedAt = &now
	} else {
		input.SilencedAt = nil
	}

	// explicitly set 2fa if provided to avoid wiping setting unintentionally
	if viper.GetBool("usersetting.update.enable-2fa") {
		input.IsTfaEnabled = lo.ToPtr(true)
	} else if viper.GetBool("usersetting.update.disable-2fa") {
		input.IsTfaEnabled = lo.ToPtr(false)
	}

	// add tags to the input if provided
	tags := viper.GetStringSlice("usersetting.update.tags")
	if len(tags) > 0 {
		input.Tags = tags
	}

	// add status to the input if provided
	status := viper.GetString("usersetting.update.status")
	if status != "" {
		input.Status = enums.ToUserStatus(status)
	}

	if settingsID == "" {
		// get the user settings id
		settings, err := client.GetUserSettings(ctx, client.Config().Interceptors...)
		if err != nil {
			return err
		}

		// this should never happen, but just in case
		if len(settings.GetUserSettings().Edges) == 0 {
			return datum.ErrNotFound
		}

		settingsID = settings.GetUserSettings().Edges[0].Node.ID
	}

	// update the user settings
	o, err := client.UpdateUserSetting(ctx, settingsID, input, client.Config().Interceptors...)
	if err != nil {
		return err
	}

	// parse the output
	s, err = json.Marshal(o)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
