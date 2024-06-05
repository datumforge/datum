package reset

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/models"
)

var registerCmd = &cobra.Command{
	Use:   "reset",
	Short: "reset a datum user password",
	RunE: func(cmd *cobra.Command, args []string) error {
		return resetPassword(cmd.Context())
	},
}

func init() {
	datum.RootCmd.AddCommand(registerCmd)

	registerCmd.Flags().StringP("token", "t", "", "reset token")
	datum.ViperBindFlag("reset.token", registerCmd.Flags().Lookup("token"))

	registerCmd.Flags().StringP("password", "p", "", "new password of the user")
	datum.ViperBindFlag("reset.password", registerCmd.Flags().Lookup("password"))
}

func resetPassword(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClient(ctx)
	if err != nil {
		return err
	}

	var s []byte

	password := viper.GetString("reset.password")
	if password == "" {
		return datum.NewRequiredFieldMissingError("password")
	}

	token := viper.GetString("reset.token")
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
