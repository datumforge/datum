package datumswitch

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"golang.org/x/oauth2"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/models"
)

var cmd = &cobra.Command{
	Use:   "switch",
	Short: "switch organization contexts",
	Run: func(cmd *cobra.Command, args []string) {
		err := switchOrg(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	datum.RootCmd.AddCommand(cmd)

	cmd.Flags().StringP("target-org", "t", "", "target organization to switch to")
}

// validate validates the required fields for the command
func validate() (*models.SwitchOrganizationRequest, error) {
	input := &models.SwitchOrganizationRequest{}

	input.TargetOrganizationID = datum.Config.String("target-org")
	if input.TargetOrganizationID == "" {
		return nil, datum.NewRequiredFieldMissingError("target organization")
	}

	return input, nil
}

// switchOrg switches the organization context
func switchOrg(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)

	input, err := validate()
	cobra.CheckErr(err)

	resp, err := client.Switch(ctx, input)
	cobra.CheckErr(err)

	fmt.Printf("Successfully switched to organization: %s!\n", input.TargetOrganizationID)

	// store auth tokens
	if err := datum.StoreToken(&oauth2.Token{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
	}); err != nil {
		return err
	}

	// store session cookies
	datum.StoreSessionCookies(client)

	fmt.Println("auth tokens successfully stored in keychain")

	return nil
}
