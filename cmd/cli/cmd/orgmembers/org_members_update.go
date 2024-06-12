package datumorgmembers

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
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
	orgMembersUpdateCmd.Flags().StringP("user-id", "u", "", "user id")
	orgMembersUpdateCmd.Flags().StringP("role", "r", "member", "role to assign the user (member, admin)")
}

func updateOrgMember(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	if err != nil {
		return err
	}
	defer datum.StoreSessionCookies(client)

	oID := datum.Config.String("org-id")

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

	// get the id of the org member
	where := datumclient.OrgMembershipWhereInput{
		UserID: &uID,
	}

	if oID != "" {
		where.OrganizationID = &oID
	}

	orgMembers, err := client.GetOrgMembersByOrgID(ctx, &where)
	if err != nil {
		return err
	}

	if len(orgMembers.OrgMemberships.Edges) != 1 {
		return errors.New("error getting existing relation") //nolint:err113
	}

	id := orgMembers.OrgMemberships.Edges[0].Node.ID

	input := datumclient.UpdateOrgMembershipInput{
		Role: &r,
	}

	var s []byte

	orgMember, err := client.UpdateUserRoleInOrg(ctx, id, input)
	if err != nil {
		return err
	}

	s, err = json.Marshal(orgMember)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
