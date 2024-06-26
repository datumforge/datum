package datumorgmembers

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/utils/cli/tables"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get existing members of a datum organization",
	Run: func(cmd *cobra.Command, args []string) {
		err := orgMembers(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(getCmd)

	getCmd.Flags().StringP("org-id", "o", "", "org id to query")
}

func orgMembers(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	where := datumclient.OrgMembershipWhereInput{}

	// filter options
	oID := datum.Config.String("id")

	if oID != "" {
		where.OrganizationID = &oID
	}

	var s []byte

	org, err := client.GetOrgMembersByOrgID(ctx, &where)
	cobra.CheckErr(err)

	if datum.OutputFormat == datum.JSONOutput {
		s, err = json.Marshal(org)
		cobra.CheckErr(err)

		return datum.JSONPrint(s)
	}

	writer := tables.NewTableWriter(cmd.OutOrStdout(), "UserID", "DisplayName", "FirstName", "LastName", "Email", "Role")

	for _, o := range org.OrgMemberships.Edges {
		writer.AddRow(
			o.Node.UserID,
			o.Node.User.DisplayName,
			*o.Node.User.FirstName,
			*o.Node.User.LastName,
			o.Node.User.Email,
			o.Node.Role,
		)
	}

	writer.Render()

	return nil
}
