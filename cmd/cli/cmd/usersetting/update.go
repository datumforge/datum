package datumusersetting

import (
	"context"
	"time"

	"github.com/samber/lo"
	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/enums"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update an existing datum user setting",
	Run: func(cmd *cobra.Command, args []string) {
		err := update(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(updateCmd)

	updateCmd.Flags().StringP("id", "i", "", "user setting id to update")
	updateCmd.Flags().String("status", "", "status of the user - active, inactive, suspended")
	updateCmd.Flags().StringP("default-org", "o", "", "default organization id")
	updateCmd.Flags().StringSliceP("tags", "t", []string{}, "tags associated with the user")
	updateCmd.Flags().BoolP("silence-notifications", "s", false, "silence notifications from datum")
	updateCmd.Flags().Bool("enable-2fa", false, "enable 2fa authentication")
	updateCmd.Flags().Bool("disable-2fa", false, "disable 2fa authentication")
}

// updateValidation validates the input flags provided by the user
func updateValidation() (id string, input datumclient.UpdateUserSettingInput, err error) {
	id = datum.Config.String("id")

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

	return id, input, nil
}

// update an existing datum user setting
func update(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	id, input, err := updateValidation()
	cobra.CheckErr(err)

	if id == "" {
		// get the user settings id
		settings, err := client.GetAllUserSettings(ctx)
		cobra.CheckErr(err)

		// this should never happen, but just in case
		if len(settings.GetUserSettings().Edges) == 0 {
			return datum.ErrNotFound
		}

		id = settings.GetUserSettings().Edges[0].Node.ID
	}

	// update the user settings
	o, err := client.UpdateUserSetting(ctx, id, input)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
