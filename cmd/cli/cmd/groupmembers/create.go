package datumgroupmembers

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "add a user to a datum group",
	Run: func(cmd *cobra.Command, args []string) {
		err := create(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(createCmd)

	createCmd.Flags().StringP("group-id", "g", "", "group id")
	createCmd.Flags().StringP("user-id", "u", "", "user id")
	createCmd.Flags().StringP("role", "r", "member", "role to assign the user (member, admin)")
}

// createValidation validates the required fields for the command
func createValidation() (input datumclient.CreateGroupMembershipInput, err error) {
	input.GroupID = datum.Config.String("group-id")
	if input.GroupID == "" {
		return input, datum.NewRequiredFieldMissingError("group id")
	}

	input.UserID = datum.Config.String("user-id")
	if input.UserID == "" {
		return input, datum.NewRequiredFieldMissingError("user id")
	}

	// role defaults to `member` so it is not required
	role := datum.Config.String("role")
	if role != "" {
		r, err := datum.GetRoleEnum(role)
		cobra.CheckErr(err)

		input.Role = &r
	}

	return input, nil
}

// create adds a user to a group in the datum platform
func create(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	input, err := createValidation()
	cobra.CheckErr(err)

	o, err := client.AddUserToGroupWithRole(ctx, input)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
