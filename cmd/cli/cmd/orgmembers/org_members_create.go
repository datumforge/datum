package datumorg

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/internal/datumclient"
	"github.com/datumforge/datum/internal/ent/enums"
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
	cli, err := datum.GetClient(ctx)
	if err != nil {
		return err
	}

	oID := viper.GetString("orgmember.create.orgid")
	if oID == "" {
		return datum.NewRequiredFieldMissingError("organization id")
	}

	uID := viper.GetString("orgmember.create.userid")
	if uID == "" {
		return datum.NewRequiredFieldMissingError("user id")
	}

	// role defaults to `member` so it is not required
	role := viper.GetString("orgmember.create.role")

	r, err := getRoleEnum(role)
	if err != nil {
		return err
	}

	input := datumclient.CreateOrgMembershipInput{
		OrgID:  oID,
		UserID: uID,
		Role:   &r,
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

func getRoleEnum(role string) (enums.Role, error) {
	r := enums.Enum(role)

	if r == enums.Invalid {
		return r, datum.ErrInvalidRole
	}

	return r, nil
}
