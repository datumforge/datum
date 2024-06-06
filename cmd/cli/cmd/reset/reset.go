package reset

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/models"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "reset a datum user password",
	RunE: func(cmd *cobra.Command, args []string) error {
		return resetPassword(cmd.Context())
	},
}

func init() {
	datum.RootCmd.AddCommand(resetCmd)

	resetCmd.Flags().StringP("token", "t", "", "reset token")
	resetCmd.Flags().StringP("password", "p", "", "new password of the user")
}

func resetPassword(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClient(ctx)
	if err != nil {
		return err
	}

	var s []byte

	password := datum.Config.String("password")
	if password == "" {
		return datum.NewRequiredFieldMissingError("password")
	}

	token := datum.Config.String("token")
	if token == "" {
		return datum.NewRequiredFieldMissingError("token")
	}

	reset := models.ResetPasswordRequest{
		Password: password,
		Token:    token,
	}

	reply, err := client.ResetPassword(ctx, &reset)
	if err != nil {
		return err
	}

	s, err = json.Marshal(reply)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
