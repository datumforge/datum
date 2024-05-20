package datumgroupmembers

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/utils/cli/tables"
)

var groupMembersGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get existing members of a datum group",
	RunE: func(cmd *cobra.Command, args []string) error {
		return groupMembers(cmd.Context())
	},
}

func init() {
	groupMembersCmd.AddCommand(groupMembersGetCmd)

	groupMembersGetCmd.Flags().StringP("group-id", "g", "", "group id to query")
	datum.ViperBindFlag("groupmember.get.id", groupMembersGetCmd.Flags().Lookup("group-id"))
}

func groupMembers(ctx context.Context) error {
	// setup datum http client
	cli, err := datum.GetGraphClient(ctx)
	if err != nil {
		return err
	}

	// save session cookies on function exit
	client, _ := cli.Client.(*datumclient.Client)
	defer datum.StoreSessionCookies(client)
	// filter options
	gID := viper.GetString("groupmember.get.id")
	if gID == "" {
		return datum.NewRequiredFieldMissingError("group id")
	}

	where := datumclient.GroupMembershipWhereInput{
		GroupID: &gID,
	}

	var s []byte

	group, err := cli.Client.GetGroupMembersByGroupID(ctx, &where, cli.Interceptor)
	if err != nil {
		return err
	}

	if datum.OutputFormat == datum.JSONOutput {
		s, err = json.Marshal(group)
		if err != nil {
			return err
		}

		return datum.JSONPrint(s)
	}

	writer := tables.NewTableWriter(groupMembersCmd.OutOrStdout(), "UserID", "DisplayName", "FirstName", "LastName", "Email", "Role")

	for _, g := range group.GroupMemberships.Edges {
		writer.AddRow(
			g.Node.UserID,
			g.Node.User.DisplayName,
			*g.Node.User.FirstName,
			*g.Node.User.LastName,
			g.Node.User.Email,
			g.Node.Role,
		)
	}

	writer.Render()

	return nil
}
