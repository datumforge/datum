package reset

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/models"
)

var cmd = &cobra.Command{
	Use:   "reset",
	Short: "reset a datum user password",
	Run: func(cmd *cobra.Command, args []string) {
		err := resetPassword(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	datum.RootCmd.AddCommand(cmd)

	cmd.Flags().StringP("token", "t", "", "reset token")
	cmd.Flags().StringP("password", "p", "", "new password of the user")
}

// validateReset validates the required fields for the command
func validateReset() (*models.ResetPasswordRequest, error) {
	password := datum.Config.String("password")
	if password == "" {
		return nil, datum.NewRequiredFieldMissingError("password")
	}

	token := datum.Config.String("token")
	if token == "" {
		return nil, datum.NewRequiredFieldMissingError("token")
	}

	return &models.ResetPasswordRequest{
		Password: password,
		Token:    token,
	}, nil
}

func resetPassword(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClient(ctx)
	cobra.CheckErr(err)

	input, err := validateReset()
	cobra.CheckErr(err)

	reply, err := client.ResetPassword(ctx, input)
	cobra.CheckErr(err)

	s, err := json.Marshal(reply)
	cobra.CheckErr(err)

	return datum.JSONPrint(s)
}
