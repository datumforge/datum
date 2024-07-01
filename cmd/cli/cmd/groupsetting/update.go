package datumgroupsetting

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/enums"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update an existing datum group setting",
	Run: func(cmd *cobra.Command, args []string) {
		err := update(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(updateCmd)

	updateCmd.Flags().StringP("id", "i", "", "group setting id to update")
	updateCmd.Flags().StringP("visibility", "v", "", "visibility of the group")
	updateCmd.Flags().StringP("join-policy", "j", "", "join policy of the group")
	updateCmd.Flags().BoolP("sync-to-slack", "s", false, "sync group members to slack")
	updateCmd.Flags().BoolP("sync-to-github", "g", false, "sync group members to github")
	updateCmd.Flags().StringSliceP("tags", "t", []string{}, "tags associated with the group")
}

// updateValidation validates the input flags provided by the user
func updateValidation() (id string, input datumclient.UpdateGroupSettingInput, err error) {
	id = datum.Config.String("id")
	if id == "" {
		return id, input, datum.NewRequiredFieldMissingError("setting id")
	}

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

	return id, input, nil
}

// update an existing group setting in the datum platform
func update(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	id, input, err := updateValidation()
	cobra.CheckErr(err)

	o, err := client.UpdateGroupSetting(ctx, id, input)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
