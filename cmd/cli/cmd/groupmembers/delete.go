package datumgroupmembers

import (
	"context"
	"errors"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "remove a user from a datum group",
	Run: func(cmd *cobra.Command, args []string) {
		err := delete(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringP("group-id", "g", "", "group id")
	deleteCmd.Flags().StringP("user-id", "u", "", "user id")
}

// deleteValidation validates the required fields for the command
func deleteValidation() (where datumclient.GroupMembershipWhereInput, err error) {
	groupID := datum.Config.String("group-id")
	if groupID == "" {
		return where, datum.NewRequiredFieldMissingError("group id")
	}

	userID := datum.Config.String("user-id")
	if userID == "" {
		return where, datum.NewRequiredFieldMissingError("user id")
	}

	where.GroupID = &groupID
	where.UserID = &userID

	return where, nil
}

// delete removes a user from a group in the datum platform
func delete(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	where, err := deleteValidation()
	cobra.CheckErr(err)

	groupMembers, err := client.GetGroupMembersByGroupID(ctx, &where)
	cobra.CheckErr(err)

	if len(groupMembers.GroupMemberships.Edges) != 1 {
		return errors.New("error getting existing relation") //nolint:err113
	}

	o, err := client.RemoveUserFromGroup(ctx, groupMembers.GroupMemberships.Edges[0].Node.ID)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
