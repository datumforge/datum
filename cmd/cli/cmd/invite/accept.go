package datuminvite

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"golang.org/x/oauth2"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/models"
)

var acceptCmd = &cobra.Command{
	Use:   "accept",
	Short: "accept an invite to join an organization",
	Run: func(cmd *cobra.Command, args []string) {
		err := accept(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(acceptCmd)

	acceptCmd.Flags().StringP("token", "t", "", "invite token")
}

// acceptValidation validates the input for the accept command
func acceptValidation() (input models.InviteRequest, err error) {
	input.Token = datum.Config.String("token")
	if input.Token == "" {
		return input, datum.NewRequiredFieldMissingError("token")
	}

	return input, nil
}

// accept an invite to join an organization
func accept(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)

	var s []byte

	input, err := acceptValidation()
	cobra.CheckErr(err)

	resp, err := client.AcceptInvite(ctx, &input)
	cobra.CheckErr(err)

	s, err = json.Marshal(resp)
	cobra.CheckErr(err)

	if err := datum.StoreToken(&oauth2.Token{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
	}); err != nil {
		cobra.CheckErr(err)
	}

	datum.StoreSessionCookies(client)

	return datum.JSONPrint(s)
}
