package datumgroupmembers

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/spf13/cobra"

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
	groupMembersDeleteCmd.Flags().StringP("user-id", "u", "", "user id")
}

func deleteGroupMember(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	gID := datum.Config.String("group-id")
	if gID == "" {
		return datum.NewRequiredFieldMissingError("group id")
	}

	uID := datum.Config.String("user-id")
	if uID == "" {
		return datum.NewRequiredFieldMissingError("user id")
	}

	// get the id of the group member
	where := datumclient.GroupMembershipWhereInput{
		GroupID: &gID,
		UserID:  &uID,
	}

	groupMembers, err := client.GetGroupMembersByGroupID(ctx, &where)
	cobra.CheckErr(err)

	if len(groupMembers.GroupMemberships.Edges) != 1 {
		return errors.New("error getting existing relation") //nolint:goerr113
	}

	id := groupMembers.GroupMemberships.Edges[0].Node.ID

	var s []byte

	groupMember, err := client.RemoveUserFromGroup(ctx, id)
	cobra.CheckErr(err)

	s, err = json.Marshal(groupMember)
	cobra.CheckErr(err)

	return datum.JSONPrint(s)
}
