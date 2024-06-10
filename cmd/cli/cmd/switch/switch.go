package datumswitch

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"golang.org/x/oauth2"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/models"
)

var loginCmd = &cobra.Command{
	Use:   "switch",
	Short: "switch organization contexts",
	RunE: func(cmd *cobra.Command, args []string) error {
		return switchorg(cmd.Context())
	},
}

func init() {
	datum.RootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringP("target-org", "t", "", "target organization to switch to")
}

func switchorg(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	if err != nil {
		return err
	}

	targetorg := datum.Config.String("target-org")
	if targetorg == "" {
		return datum.NewRequiredFieldMissingError("target organization")
	}

	input := models.SwitchOrganizationRequest{
		TargetOrganizationID: targetorg,
	}

	resp, err := client.Switch(ctx, &input)
	if err != nil {
		return err
	}

	fmt.Printf("Successfully switched to organization: %s!\n", targetorg)

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
