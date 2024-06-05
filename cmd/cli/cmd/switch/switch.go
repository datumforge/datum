package datumswitch

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

	loginCmd.Flags().StringP("targetorg", "t", "", "target organization to switch to")
	datum.ViperBindFlag("switch.targetorg", loginCmd.Flags().Lookup("targetorg"))
}

func switchorg(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	if err != nil {
		return err
	}
	defer datum.StoreSessionCookies(client)

	targetorg := viper.GetString("switch.targetorg")
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

	if err := datum.StoreToken(&oauth2.Token{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
	}); err != nil {
		return err
	}

	fmt.Println("auth tokens successfully stored in keychain")

	return nil
}
