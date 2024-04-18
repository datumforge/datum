package datumswitch

import (
	"context"
	"encoding/json"

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
		if err := switchorg(cmd.Context()); err != nil {
			return err

		}
		return nil
	},
}

func init() {
	datum.RootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringP("targetorg", "t", "", "target organization to switch to")
	datum.ViperBindFlag("switch.targetorg", loginCmd.Flags().Lookup("targetorg"))
}

func switchorg(ctx context.Context) error {
	var s []byte

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

	input := handlers.SwitchOrganizationRequest{
		TargetOrganizationID: targetorg,
	}

	SwitchOrganizationReply, err := datumclient.Switch(dc, ctx, input, cli.AccessToken)
	if err != nil {
		return err
	}

	s, err = json.Marshal(SwitchOrganizationReply)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
