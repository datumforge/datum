package datumgroupmembers

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/internal/datumclient"
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
	cli, err := datum.GetClient(ctx)
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

	s, err = json.Marshal(group)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
