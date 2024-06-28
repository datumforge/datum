package datumgroupmembers

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get existing members of a datum group",
	Run: func(cmd *cobra.Command, args []string) {
		err := get(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(getCmd)

	getCmd.Flags().StringP("group-id", "g", "", "group id to query")
}

// get existing members of a datum group
func get(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	// filter options
	id := datum.Config.String("group-id")
	if id == "" {
		return datum.NewRequiredFieldMissingError("group id")
	}

	where := datumclient.GroupMembershipWhereInput{
		GroupID: &id,
	}

	o, err := client.GetGroupMembersByGroupID(ctx, &where)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
