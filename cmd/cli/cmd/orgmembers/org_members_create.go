package datumorgmembers

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var orgMembersCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Add user to a datum org",
	RunE: func(cmd *cobra.Command, args []string) error {
		return addOrgMember(cmd.Context())
	},
}

func init() {
	orgMembersCmd.AddCommand(orgMembersCreateCmd)

	orgMembersCreateCmd.Flags().StringP("org-id", "o", "", "organization id")
	datum.ViperBindFlag("orgmember.create.orgid", orgMembersCreateCmd.Flags().Lookup("org-id"))

	orgMembersCreateCmd.Flags().StringP("user-id", "u", "", "user id")
	datum.ViperBindFlag("orgmember.create.userid", orgMembersCreateCmd.Flags().Lookup("user-id"))

	orgMembersCreateCmd.Flags().StringP("role", "r", "member", "role to assign the user (member, admin)")
	datum.ViperBindFlag("orgmember.create.role", orgMembersCreateCmd.Flags().Lookup("role"))
}

func addOrgMember(ctx context.Context) error {
	// setup datum http client
	cli, err := datum.GetGraphClient(ctx)
	if err != nil {
		return err
	}

	// save session cookies on function exit
	client, _ := cli.Client.(*datumclient.Client)
	defer datum.StoreSessionCookies(client)

	oID := viper.GetString("orgmember.create.orgid")

	uID := viper.GetString("orgmember.create.userid")
	if uID == "" {
		return datum.NewRequiredFieldMissingError("user id")
	}

	// role defaults to `member` so it is not required
	role := viper.GetString("orgmember.create.role")

	r, err := datum.GetRoleEnum(role)
	if err != nil {
		return err
	}

	input := datumclient.CreateOrgMembershipInput{
		UserID: uID,
		Role:   &r,
	}

	if oID != "" {
		input.OrganizationID = oID
	}

	var s []byte

	orgMember, err := cli.Client.AddUserToOrgWithRole(ctx, input, cli.Interceptor)
	if err != nil {
		return err
	}

	s, err = json.Marshal(orgMember)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
