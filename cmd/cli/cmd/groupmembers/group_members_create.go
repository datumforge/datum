package datumgroupmembers

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

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
	datum.ViperBindFlag("groupmember.create.groupid", groupMembersCreateCmd.Flags().Lookup("group-id"))

	groupMembersCreateCmd.Flags().StringP("user-id", "u", "", "user id")
	datum.ViperBindFlag("groupmember.create.userid", groupMembersCreateCmd.Flags().Lookup("user-id"))

	groupMembersCreateCmd.Flags().StringP("role", "r", "member", "role to assign the user (member, admin)")
	datum.ViperBindFlag("groupmember.create.role", groupMembersCreateCmd.Flags().Lookup("role"))
}

func addGroupMember(ctx context.Context) error {
	// setup datum http client
	cli, err := datum.GetGraphClient(ctx)
	if err != nil {
		return err
	}

	// save session cookies on function exit
	client, _ := cli.Client.(*datumclient.Client)
	defer datum.StoreSessionCookies(client)

	gID := viper.GetString("groupmember.create.groupid")
	if gID == "" {
		return datum.NewRequiredFieldMissingError("group id")
	}

	uID := viper.GetString("groupmember.create.userid")
	if uID == "" {
		return datum.NewRequiredFieldMissingError("user id")
	}

	// role defaults to `member` so it is not required
	role := viper.GetString("groupmember.create.role")

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

	groupMember, err := cli.Client.AddUserToGroupWithRole(ctx, input, cli.Interceptor)
	if err != nil {
		return err
	}

	s, err = json.Marshal(groupMember)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
