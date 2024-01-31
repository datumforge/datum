package datumorgmembers

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/internal/datumclient"
)

var orgMembersUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a user's role in a datum org",
	RunE: func(cmd *cobra.Command, args []string) error {
		return updateOrgMember(cmd.Context())
	},
}

func init() {
	orgMembersCmd.AddCommand(orgMembersUpdateCmd)

	orgMembersUpdateCmd.Flags().StringP("org-id", "o", "", "organization id")
	datum.ViperBindFlag("orgmember.update.orgid", orgMembersUpdateCmd.Flags().Lookup("org-id"))

	orgMembersUpdateCmd.Flags().StringP("user-id", "u", "", "user id")
	datum.ViperBindFlag("orgmember.update.userid", orgMembersUpdateCmd.Flags().Lookup("user-id"))

	orgMembersUpdateCmd.Flags().StringP("role", "r", "member", "role to assign the user (member, admin)")
	datum.ViperBindFlag("orgmember.update.role", orgMembersUpdateCmd.Flags().Lookup("role"))
}

func updateOrgMember(ctx context.Context) error {
	// setup datum http client
	cli, err := datum.GetClient(ctx)
	if err != nil {
		return err
	}

	// save session cookies on function exit
	client, _ := cli.Client.(*datumclient.Client)
	defer datum.StoreSessionCookies(client)

	oID := viper.GetString("orgmember.update.orgid")
	if oID == "" {
		return datum.NewRequiredFieldMissingError("organization id")
	}

	uID := viper.GetString("orgmember.update.userid")
	if uID == "" {
		return datum.NewRequiredFieldMissingError("user id")
	}

	role := viper.GetString("orgmember.update.role")
	if role == "" {
		return datum.NewRequiredFieldMissingError("role")
	}

	r, err := datum.GetRoleEnum(role)
	if err != nil {
		return err
	}

	// get the id of the org member
	where := datumclient.OrgMembershipWhereInput{
		OrganizationID: &oID,
		UserID:         &uID,
	}

	orgMembers, err := cli.Client.GetOrgMembersByOrgID(ctx, &where, cli.Interceptor)
	if err != nil {
		return err
	}

	if len(orgMembers.OrgMemberships.Edges) != 1 {
		return errors.New("error getting existing relation") //nolint:goerr113
	}

	id := orgMembers.OrgMemberships.Edges[0].Node.ID

	input := datumclient.UpdateOrgMembershipInput{
		Role: &r,
	}

	var s []byte

	orgMember, err := cli.Client.UpdateUserRoleInOrg(ctx, id, input, cli.Interceptor)
	if err != nil {
		return err
	}

	s, err = json.Marshal(orgMember)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
