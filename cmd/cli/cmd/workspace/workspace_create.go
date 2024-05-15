package datumworkspace

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var workspaceCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new datum org",
	RunE: func(cmd *cobra.Command, args []string) error {
		return createWorkspace(cmd.Context())
	},
}

func init() {
	workspaceCmd.AddCommand(workspaceCreateCmd)

	workspaceCreateCmd.Flags().StringP("name", "n", "", "name of the workspace")
	datum.ViperBindFlag("workspace.create.name", workspaceCreateCmd.Flags().Lookup("name"))

	workspaceCreateCmd.Flags().StringP("description", "d", "", "description of the workspace")
	datum.ViperBindFlag("workspace.create.description", workspaceCreateCmd.Flags().Lookup("description"))
}

func createWorkspace(ctx context.Context) error {
	// setup datum http client
	cli, err := datum.GetGraphClient(ctx)
	if err != nil {
		return err
	}

	// save session cookies on function exit
	client, _ := cli.Client.(*datumclient.Client)
	defer datum.StoreSessionCookies(client)

	var s []byte

	name := viper.GetString("workspace.create.name")
	if name == "" {
		return datum.NewRequiredFieldMissingError("workspace name")
	}

	description := viper.GetString("org.create.description")

	input := datumclient.CreateWorkspaceInput{
		Name: name,
	}

	if description != "" {
		input.Description = &description
	}

	o, err := cli.Client.CreateWorkspace(ctx, input, cli.Interceptor)
	if err != nil {
		return err
	}

	s, err = json.Marshal(o)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
