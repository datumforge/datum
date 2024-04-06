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

var groupMembersUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a user's role in a datum group",
	RunE: func(cmd *cobra.Command, args []string) error {
		return updateGroupMember(cmd.Context())
	},
}

func init() {
	groupMembersCmd.AddCommand(groupMembersUpdateCmd)

	groupMembersUpdateCmd.Flags().StringP("group-id", "g", "", "group id")
	datum.ViperBindFlag("groupmember.update.groupid", groupMembersUpdateCmd.Flags().Lookup("group-id"))

	groupMembersUpdateCmd.Flags().StringP("user-id", "u", "", "user id")
	datum.ViperBindFlag("groupmember.update.userid", groupMembersUpdateCmd.Flags().Lookup("user-id"))

	groupMembersUpdateCmd.Flags().StringP("role", "r", "member", "role to assign the user (member, admin)")
	datum.ViperBindFlag("groupmember.update.role", groupMembersUpdateCmd.Flags().Lookup("role"))
}

func updateGroupMember(ctx context.Context) error {
	// setup datum http client
	cli, err := datum.GetGraphClient(ctx)
	if err != nil {
		return err
	}

	// save session cookies on function exit
	client, _ := cli.Client.(*datumclient.Client)
	defer datum.StoreSessionCookies(client)

	gID := viper.GetString("groupmember.update.groupid")
	if gID == "" {
		return datum.NewRequiredFieldMissingError("group id")
	}

	uID := viper.GetString("groupmember.update.userid")
	if uID == "" {
		return datum.NewRequiredFieldMissingError("user id")
	}

	role := viper.GetString("groupmember.update.role")
	if role == "" {
		return datum.NewRequiredFieldMissingError("role")
	}

	r, err := datum.GetRoleEnum(role)
	if err != nil {
		return err
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

	input := datumclient.UpdateGroupMembershipInput{
		Role: &r,
	}

	var s []byte

	groupMember, err := cli.Client.UpdateUserRoleInGroup(ctx, id, input, cli.Interceptor)
	if err != nil {
		return err
	}

	s, err = json.Marshal(groupMember)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
