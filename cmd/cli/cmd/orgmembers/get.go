package datumorgmembers

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get existing members of a datum organization",
	Run: func(cmd *cobra.Command, args []string) {
		err := get(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(getCmd)

	getCmd.Flags().StringP("org-id", "o", "", "org id to query")
}

func get(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	where := datumclient.OrgMembershipWhereInput{}

	// filter options
	id := datum.Config.String("id")

	if id != "" {
		where.OrganizationID = &id
	}

	o, err := client.GetOrgMembersByOrgID(ctx, &where)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
