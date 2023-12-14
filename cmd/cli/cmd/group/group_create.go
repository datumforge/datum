package datumgroup

import (
	"context"
	"encoding/json"

	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/internal/datumclient"
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

	groupCreateCmd.Flags().StringP("owner-id", "o", "", "owner org id")
	datum.ViperBindFlag("group.create.owner-id", groupCreateCmd.Flags().Lookup("owner-id"))
}

func createGroup(ctx context.Context) error {
	cli, err := datum.GetClient(ctx)
	if err != nil {
		return err
	}

	var s []byte

	name := viper.GetString("group.create.name")
	if name == "" {
		return datum.NewRequiredFieldMissingError("group name")
	}

	owner := viper.GetString("group.create.owner-id")
	if owner == "" {
		return datum.NewRequiredFieldMissingError("organization id")
	}

	displayName := viper.GetString("group.create.short-name")
	description := viper.GetString("group.create.description")

	input := datumclient.CreateGroupInput{
		Name:    name,
		OwnerID: owner,
	}

	if displayName != "" {
		input.DisplayName = &displayName
	}

	if description != "" {
		input.Description = &description
	}

	o, err := cli.Client.CreateGroup(ctx, input, cli.Interceptor)
	if err != nil {
		return err
	}

	s, err = json.Marshal(o)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
