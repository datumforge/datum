package datumgroupsetting

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"

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
	groupSettingUpdateCmd.Flags().StringP("visibility", "v", "", "visibility of the group")
	groupSettingUpdateCmd.Flags().StringP("join-policy", "j", "", "join policy of the group")
	groupSettingUpdateCmd.Flags().BoolP("sync-to-slack", "s", false, "sync group members to slack")
	groupSettingUpdateCmd.Flags().BoolP("sync-to-github", "g", false, "sync group members to github")
	groupSettingUpdateCmd.Flags().StringSliceP("tags", "t", []string{}, "tags associated with the group")
}

func updateGroupSetting(ctx context.Context) error {
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

	input := datumclient.UpdateGroupSettingInput{}

	visibility := datum.Config.String("visibility")
	if visibility != "" {
		input.Visibility = enums.ToGroupVisibility(visibility)
	}

	joinPolicy := datum.Config.String("join-policy")
	if joinPolicy != "" {
		input.JoinPolicy = enums.ToGroupJoinPolicy(joinPolicy)
	}

	tags := datum.Config.Strings("tags")
	if len(tags) > 0 {
		input.Tags = tags
	}

	syncToSlack := datum.Config.Bool("sync-to-slack")
	if syncToSlack {
		input.SyncToSlack = &syncToSlack
	}

	syncToGithub := datum.Config.Bool("sync-to-github")
	if syncToGithub {
		input.SyncToGithub = &syncToGithub
	}

	o, err := client.UpdateGroupSetting(ctx, settingsID, input)
	if err != nil {
		return err
	}

	s, err = json.Marshal(o)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
