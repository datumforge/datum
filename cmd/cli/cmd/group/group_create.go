package datumgroup

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var groupCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new datum group",
	RunE: func(cmd *cobra.Command, args []string) error {
		return createGroup(cmd.Context())
	},
}

func init() {
	groupCmd.AddCommand(groupCreateCmd)

	groupCreateCmd.Flags().StringP("name", "n", "", "name of the group")
	groupCreateCmd.Flags().StringP("short-name", "s", "", "display name of the group")
	groupCreateCmd.Flags().StringP("description", "d", "", "description of the group")
}

func createGroup(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	var s []byte

	name := datum.Config.String("name")
	if name == "" {
		return datum.NewRequiredFieldMissingError("group name")
	}

	displayName := datum.Config.String("short-name")
	description := datum.Config.String("description")

	input := datumclient.CreateGroupInput{
		Name: name,
	}

	if displayName != "" {
		input.DisplayName = &displayName
	}

	if description != "" {
		input.Description = &description
	}

	o, err := client.CreateGroup(ctx, input)
	cobra.CheckErr(err)

	s, err = json.Marshal(o)
	cobra.CheckErr(err)

	return datum.JSONPrint(s)
}
