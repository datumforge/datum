package datumgroupmembers

import (
	"context"
	"errors"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update a user's role in a datum group",
	Run: func(cmd *cobra.Command, args []string) {
		err := update(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(updateCmd)

	updateCmd.Flags().StringP("group-id", "g", "", "group id")
	updateCmd.Flags().StringP("user-id", "u", "", "user id")
	updateCmd.Flags().StringP("role", "r", "member", "role to assign the user (member, admin)")
}

// updateValidation validates the required fields for the command
func updateValidation() (where datumclient.GroupMembershipWhereInput, input datumclient.UpdateGroupMembershipInput, err error) {
	groupID := datum.Config.String("group-id")
	if groupID == "" {
		return where, input, datum.NewRequiredFieldMissingError("group id")
	}

	userID := datum.Config.String("user-id")
	if userID == "" {
		return where, input, datum.NewRequiredFieldMissingError("user id")
	}

	role := datum.Config.String("role")
	if role == "" {
		return where, input, datum.NewRequiredFieldMissingError("role")
	}

	r, err := datum.GetRoleEnum(role)
	cobra.CheckErr(err)

	where = datumclient.GroupMembershipWhereInput{
		GroupID: &groupID,
		UserID:  &userID,
	}

	input = datumclient.UpdateGroupMembershipInput{
		Role: &r,
	}

	return where, input, nil
}

// update a user's role in a group in the datum platform
func update(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	where, input, err := updateValidation()
	cobra.CheckErr(err)

	groupMembers, err := client.GetGroupMembersByGroupID(ctx, &where)
	cobra.CheckErr(err)

	if len(groupMembers.GroupMemberships.Edges) != 1 {
		return errors.New("error getting existing relation") //nolint:err113
	}

	o, err := client.UpdateUserRoleInGroup(ctx, groupMembers.GroupMemberships.Edges[0].Node.ID, input)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
