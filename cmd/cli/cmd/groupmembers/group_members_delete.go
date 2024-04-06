package datumgroupmembers

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var groupMembersDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove a user from a datum group",
	RunE: func(cmd *cobra.Command, args []string) error {
		return deleteGroupMember(cmd.Context())
	},
}

func init() {
	groupMembersCmd.AddCommand(groupMembersDeleteCmd)

	groupMembersDeleteCmd.Flags().StringP("group-id", "g", "", "group id")
	datum.ViperBindFlag("groupmember.delete.groupid", groupMembersDeleteCmd.Flags().Lookup("group-id"))

	groupMembersDeleteCmd.Flags().StringP("user-id", "u", "", "user id")
	datum.ViperBindFlag("groupmember.delete.userid", groupMembersDeleteCmd.Flags().Lookup("user-id"))
}

func deleteGroupMember(ctx context.Context) error {
	// setup datum http client
	cli, err := datum.GetGraphClient(ctx)
	if err != nil {
		return err
	}

	// save session cookies on function exit
	client, _ := cli.Client.(*datumclient.Client)
	defer datum.StoreSessionCookies(client)

	gID := viper.GetString("groupmember.delete.groupid")
	if gID == "" {
		return datum.NewRequiredFieldMissingError("group id")
	}

	uID := viper.GetString("groupmember.delete.userid")
	if uID == "" {
		return datum.NewRequiredFieldMissingError("user id")
	}

	// get the id of the group member
	where := datumclient.GroupMembershipWhereInput{
		GroupID: &gID,
		UserID:  &uID,
	}

	groupMembers, err := cli.Client.GetGroupMembersByGroupID(ctx, &where, cli.Interceptor)
	if err != nil {
		return err
	}

	if len(groupMembers.GroupMemberships.Edges) != 1 {
		return errors.New("error getting existing relation") //nolint:goerr113
	}

	id := groupMembers.GroupMemberships.Edges[0].Node.ID

	var s []byte

	groupMember, err := cli.Client.RemoveUserFromGroup(ctx, id, cli.Interceptor)
	if err != nil {
		return err
	}

	s, err = json.Marshal(groupMember)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
