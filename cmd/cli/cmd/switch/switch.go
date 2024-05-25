package datumswitch

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
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
	cli, err := datum.GetRestClient(ctx)
	if err != nil {
		return err
	}

	dc := cli.Client.(*datumclient.Client)
	defer datum.StoreSessionCookies(dc)

	targetorg := viper.GetString("switch.targetorg")
	if targetorg == "" {
		return datum.NewRequiredFieldMissingError("target organization")
	}

	input := models.SwitchOrganizationRequest{
		TargetOrganizationID: targetorg,
	}

	switchOrganizationReply, err := datumclient.Switch(dc, ctx, input, cli.AccessToken)
	if err != nil {
		return err
	}

	fmt.Printf("Successfully switched to organization: %s!\n", targetorg)

	if err := datum.StoreToken(switchOrganizationReply); err != nil {
		return err
	}

	fmt.Println("auth tokens successfully stored in keychain")

	return nil
}
