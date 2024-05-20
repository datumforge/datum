package datumgroup

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/utils/cli/tables"
)

var groupGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get an existing new datum group",
	RunE: func(cmd *cobra.Command, args []string) error {
		return getGroup(cmd.Context())
	},
}

func init() {
	groupCmd.AddCommand(groupGetCmd)

	groupGetCmd.Flags().StringP("id", "i", "", "group id to query")
	datum.ViperBindFlag("group.get.id", groupGetCmd.Flags().Lookup("id"))
}

func getGroup(ctx context.Context) error {
	// setup datum http client
	cli, err := datum.GetGraphClient(ctx)
	if err != nil {
		return err
	}

	// save session cookies on function exit
	client, _ := cli.Client.(*datumclient.Client)
	defer datum.StoreSessionCookies(client)

	// filter options
	gID := viper.GetString("group.get.id")

	writer := tables.NewTableWriter(groupCmd.OutOrStdout(), "ID", "Name", "Description", "Visibility", "Organization", "Members")

	// if an group ID is provided, filter on that group, otherwise get all
	if gID != "" {
		group, err := cli.Client.GetGroupByID(ctx, gID, cli.Interceptor)
		if err != nil {
			return err
		}

		if datum.OutputFormat == datum.JSONOutput {

			s, err := json.Marshal(group)
			if err != nil {
				return err
			}

			return datum.JSONPrint(s)
		}

		writer.AddRow(group.Group.ID, group.Group.DisplayName, *group.Group.Description, group.Group.Setting.Visibility, group.Group.Owner.DisplayName, len(group.Group.Members))
		writer.Render()

		return nil
	}

	// get all groups, will be filtered for the authorized organization(s)
	groups, err := cli.Client.GetAllGroups(ctx, cli.Interceptor)
	if err != nil {
		return err
	}

	if datum.OutputFormat == datum.JSONOutput {
		s, err := json.Marshal(groups)
		if err != nil {
			return err
		}

		return datum.JSONPrint(s)
	}

	for _, g := range groups.Groups.Edges {
		writer.AddRow(g.Node.ID, g.Node.DisplayName, *g.Node.Description, g.Node.Setting.Visibility, g.Node.Owner.DisplayName, len(g.Node.Members))
	}

	writer.Render()

	return nil
}
