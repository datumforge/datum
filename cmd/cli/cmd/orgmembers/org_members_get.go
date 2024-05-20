package datumorgmembers

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/utils/cli/tables"
)

var orgMembersGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get existing members of a datum org",
	RunE: func(cmd *cobra.Command, args []string) error {
		return orgMembers(cmd.Context())
	},
}

func init() {
	orgMembersCmd.AddCommand(orgMembersGetCmd)

	orgMembersGetCmd.Flags().StringP("org-id", "o", "", "org id to query")
	datum.ViperBindFlag("orgmember.get.id", orgMembersGetCmd.Flags().Lookup("org-id"))
}

func orgMembers(ctx context.Context) error {
	// setup datum http client
	cli, err := datum.GetGraphClient(ctx)
	if err != nil {
		return err
	}

	// save session cookies on function exit
	client, _ := cli.Client.(*datumclient.Client)
	defer datum.StoreSessionCookies(client)

	where := datumclient.OrgMembershipWhereInput{}

	// filter options
	oID := viper.GetString("orgmember.get.id")

	if oID != "" {
		where.OrganizationID = &oID
	}

	var s []byte

	org, err := cli.Client.GetOrgMembersByOrgID(ctx, &where, cli.Interceptor)
	if err != nil {
		return err
	}

	if datum.OutputFormat == datum.JSONOutput {
		s, err = json.Marshal(org)
		if err != nil {
			return err
		}

		return datum.JSONPrint(s)
	}

	writer := tables.NewTableWriter(orgMembersCmd.OutOrStdout(), "UserID", "DisplayName", "FirstName", "LastName", "Email", "Role")

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
