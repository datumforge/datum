package datumgroup

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

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
	datum.ViperBindFlag("group.create.name", groupCreateCmd.Flags().Lookup("name"))

	groupCreateCmd.Flags().StringP("short-name", "s", "", "display name of the group")
	datum.ViperBindFlag("group.create.short-name", groupCreateCmd.Flags().Lookup("short-name"))

	groupCreateCmd.Flags().StringP("description", "d", "", "description of the group")
	datum.ViperBindFlag("group.create.description", groupCreateCmd.Flags().Lookup("description"))
}

func createGroup(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	if err != nil {
		return err
	}
	defer datum.StoreSessionCookies(client)

	var s []byte

	name := viper.GetString("group.create.name")
	if name == "" {
		return datum.NewRequiredFieldMissingError("group name")
	}

	displayName := viper.GetString("group.create.short-name")
	description := viper.GetString("group.create.description")

	input := datumclient.CreateGroupInput{
		Name: name,
	}

	if displayName != "" {
		input.DisplayName = &displayName
	}

	if description != "" {
		input.Description = &description
	}

	o, err := client.CreateGroup(ctx, input, client.Config().Interceptors...)
	if err != nil {
		return err
	}

	s, err = json.Marshal(o)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
