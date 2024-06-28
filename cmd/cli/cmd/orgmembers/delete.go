package datumorgmembers

import (
	"context"
	"errors"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "remove a user from a datum organization",
	Run: func(cmd *cobra.Command, args []string) {
		err := delete(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringP("user-id", "u", "", "user id")
}

// deleteValidation validates the required fields for the command
func deleteValidation() (where datumclient.OrgMembershipWhereInput, err error) {
	uID := datum.Config.String("user-id")
	if uID == "" {
		return where, datum.NewRequiredFieldMissingError("user id")
	}

	where.UserID = &uID

	return where, nil
}

func delete(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	where, err := deleteValidation()
	cobra.CheckErr(err)

	orgMembers, err := client.GetOrgMembersByOrgID(ctx, &where)
	cobra.CheckErr(err)

	if len(orgMembers.OrgMemberships.Edges) != 1 {
		cobra.CheckErr(errors.New("error getting existing relation")) //nolint:err113
	}

	id := orgMembers.OrgMemberships.Edges[0].Node.ID

	o, err := client.RemoveUserFromOrg(ctx, id)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
