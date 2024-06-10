package datumgroupmembers

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/spf13/cobra"

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
	groupMembersUpdateCmd.Flags().StringP("user-id", "u", "", "user id")
	groupMembersUpdateCmd.Flags().StringP("role", "r", "member", "role to assign the user (member, admin)")
}

func updateGroupMember(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	if err != nil {
		return err
	}
	defer datum.StoreSessionCookies(client)

	gID := datum.Config.String("group-id")
	if gID == "" {
		return datum.NewRequiredFieldMissingError("group id")
	}

	uID := datum.Config.String("user-id")
	if uID == "" {
		return datum.NewRequiredFieldMissingError("user id")
	}

	role := datum.Config.String("role")
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

	groupMembers, err := client.GetGroupMembersByGroupID(ctx, &where)
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

	groupMember, err := client.UpdateUserRoleInGroup(ctx, id, input)
	if err != nil {
		return err
	}

	s, err = json.Marshal(groupMember)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
