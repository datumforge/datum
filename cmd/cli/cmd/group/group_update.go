package datumgroup

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var groupUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing datum group",
	RunE: func(cmd *cobra.Command, args []string) error {
		return updateGroup(cmd.Context())
	},
}

func init() {
	groupCmd.AddCommand(groupUpdateCmd)

	groupUpdateCmd.Flags().StringP("id", "i", "", "group id to update")
	groupUpdateCmd.Flags().StringP("name", "n", "", "name of the group")
	groupUpdateCmd.Flags().StringP("short-name", "s", "", "display name of the group")
	groupUpdateCmd.Flags().StringP("description", "d", "", "description of the group")
}

func updateGroup(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	var s []byte

	gID := datum.Config.String("id")
	if gID == "" {
		return datum.NewRequiredFieldMissingError("group id")
	}

	name := datum.Config.String("name")
	displayName := datum.Config.String("short-name")
	description := datum.Config.String("description")

	input := datumclient.UpdateGroupInput{}

	if name != "" {
		input.Name = &name
	}

	if displayName != "" {
		input.DisplayName = &displayName
	}

	if description != "" {
		input.Description = &description
	}

	o, err := client.UpdateGroup(ctx, gID, input)
	cobra.CheckErr(err)

	s, err = json.Marshal(o)
	cobra.CheckErr(err)

	return datum.JSONPrint(s)
}
