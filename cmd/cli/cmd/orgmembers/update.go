package datumorgmembers

import (
	"context"
	"errors"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update a user's role in a datum org",
	Run: func(cmd *cobra.Command, args []string) {
		err := update(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(updateCmd)

	updateCmd.Flags().StringP("org-id", "o", "", "organization id")
	updateCmd.Flags().StringP("user-id", "u", "", "user id")
	updateCmd.Flags().StringP("role", "r", "member", "role to assign the user (member, admin)")
}

// updateValidation validates the required fields for the command
func updateValidation() (where datumclient.OrgMembershipWhereInput, input datumclient.UpdateOrgMembershipInput, err error) {
	userID := datum.Config.String("user-id")
	if userID == "" {
		return where, input, datum.NewRequiredFieldMissingError("user id")
	}

	where.UserID = &userID

	orgID := datum.Config.String("org-id")
	if orgID != "" {
		where.OrganizationID = &orgID
	}

	role := datum.Config.String("role")
	if role == "" {
		return where, input, datum.NewRequiredFieldMissingError("role")
	}

	r, err := datum.GetRoleEnum(role)
	cobra.CheckErr(err)

	input.Role = &r

	return where, input, nil
}

func update(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	where, input, err := updateValidation()
	cobra.CheckErr(err)

	orgMembers, err := client.GetOrgMembersByOrgID(ctx, &where)
	cobra.CheckErr(err)

	if len(orgMembers.OrgMemberships.Edges) != 1 {
		return errors.New("error getting existing relation") //nolint:err113
	}

	o, err := client.UpdateUserRoleInOrg(ctx, orgMembers.OrgMemberships.Edges[0].Node.ID, input)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
