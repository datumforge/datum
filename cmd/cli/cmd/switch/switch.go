package datumswitch

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/pkg/datumclient"
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
	targetorg := viper.GetString("switch.targetorg")
	if targetorg == "" {
		return datum.NewRequiredFieldMissingError("target organization")
	}

	s, err := SwitchOrg(ctx, targetorg)
	if err != nil {
		return err
	}

	fmt.Println("auth tokens successfully stored in keychain")

	return datum.JSONPrint(s)
}

func SwitchOrg(ctx context.Context, targetOrg string) ([]byte, error) {
	var s []byte

	cli, err := datum.GetRestClient(ctx)
	if err != nil {
		return nil, err
	}

	dc := cli.Client.(*datumclient.Client)
	defer datum.StoreSessionCookies(dc)

	input := handlers.SwitchOrganizationRequest{
		TargetOrganizationID: targetOrg,
	}

	switchOrganizationReply, err := datumclient.Switch(dc, ctx, input, cli.AccessToken)
	if err != nil {
		return nil, err
	}

	s, err = json.Marshal(switchOrganizationReply)
	if err != nil {
		return nil, err
	}

	return s, datum.StoreToken(switchOrganizationReply)
}
