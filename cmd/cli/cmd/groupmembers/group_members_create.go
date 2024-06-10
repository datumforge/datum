package datumgroupmembers

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var groupMembersCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Add user to a datum group",
	RunE: func(cmd *cobra.Command, args []string) error {
		return addGroupMember(cmd.Context())
	},
}

func init() {
	groupMembersCmd.AddCommand(groupMembersCreateCmd)

	groupMembersCreateCmd.Flags().StringP("group-id", "g", "", "group id")
	groupMembersCreateCmd.Flags().StringP("user-id", "u", "", "user id")
	groupMembersCreateCmd.Flags().StringP("role", "r", "member", "role to assign the user (member, admin)")
}

func addGroupMember(ctx context.Context) error {
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

	// role defaults to `member` so it is not required
	role := datum.Config.String("role")

	r, err := datum.GetRoleEnum(role)
	if err != nil {
		return err
	}

	input := datumclient.CreateGroupMembershipInput{
		GroupID: gID,
		UserID:  uID,
		Role:    &r,
	}

	var s []byte

	groupMember, err := client.AddUserToGroupWithRole(ctx, input)
	if err != nil {
		return err
	}

	s, err = json.Marshal(groupMember)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
