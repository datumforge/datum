package datumgroupsetting

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/pkg/datumclient"
)

var groupSettingUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing datum group setting",
	RunE: func(cmd *cobra.Command, args []string) error {
		return updateGroupSetting(cmd.Context())
	},
}

func init() {
	groupSettingCmd.AddCommand(groupSettingUpdateCmd)

	groupSettingUpdateCmd.Flags().StringP("id", "i", "", "group setting id to update")
	datum.ViperBindFlag("groupsetting.update.id", groupSettingUpdateCmd.Flags().Lookup("id"))

	groupSettingUpdateCmd.Flags().StringP("visibility", "v", "", "visibility of the group")
	datum.ViperBindFlag("groupsetting.update.visibility", groupSettingUpdateCmd.Flags().Lookup("visibility"))

	groupSettingUpdateCmd.Flags().StringP("join-policy", "j", "", "join policy of the group")
	datum.ViperBindFlag("groupsetting.update.joinpolicy", groupSettingUpdateCmd.Flags().Lookup("join-policy"))

	groupSettingUpdateCmd.Flags().BoolP("sync-to-slack", "s", false, "sync group members to slack")
	datum.ViperBindFlag("groupsetting.update.synctoslack", groupSettingUpdateCmd.Flags().Lookup("sync-to-slack"))

	groupSettingUpdateCmd.Flags().BoolP("sync-to-github", "g", false, "sync group members to github")
	datum.ViperBindFlag("groupsetting.update.synctogithub", groupSettingUpdateCmd.Flags().Lookup("sync-to-github"))

	groupSettingUpdateCmd.Flags().StringSliceP("tags", "t", []string{}, "tags associated with the group")
	datum.ViperBindFlag("groupsetting.update.tags", groupSettingUpdateCmd.Flags().Lookup("tags"))
}

func updateGroupSetting(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	if err != nil {
		return err
	}
	defer datum.StoreSessionCookies(client)

	var s []byte

	settingsID := viper.GetString("groupsetting.update.id")
	if settingsID == "" {
		return datum.NewRequiredFieldMissingError("setting id")
	}

	input := datumclient.UpdateGroupSettingInput{}

	visibility := viper.GetString("groupsetting.update.visibility")
	if visibility != "" {
		input.Visibility = enums.ToGroupVisibility(visibility)
	}

	joinPolicy := viper.GetString("groupsetting.update.joinpolicy")
	if joinPolicy != "" {
		input.JoinPolicy = enums.ToGroupJoinPolicy(joinPolicy)
	}

	tags := viper.GetStringSlice("groupsetting.update.tags")
	if len(tags) > 0 {
		input.Tags = tags
	}

	syncToSlack := viper.GetBool("groupsetting.update.synctoslack")
	if syncToSlack {
		input.SyncToSlack = &syncToSlack
	}

	syncToGithub := viper.GetBool("groupsetting.update.synctogithub")
	if syncToGithub {
		input.SyncToGithub = &syncToGithub
	}

	o, err := client.UpdateGroupSetting(ctx, settingsID, input, client.Config().Interceptors...)
	if err != nil {
		return err
	}

	s, err = json.Marshal(o)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
