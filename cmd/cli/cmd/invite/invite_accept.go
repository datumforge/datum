package datuminvite

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/models"
)

var inviteAcceptCmd = &cobra.Command{
	Use:   "accept",
	Short: "accept an invite",
	RunE: func(cmd *cobra.Command, args []string) error {
		return inviteAccept(cmd.Context())
	},
}

func init() {
	inviteCmd.AddCommand(inviteAcceptCmd)

	inviteAcceptCmd.Flags().StringP("token", "t", "", "invite token")
	datum.ViperBindFlag("invite.accept.token", inviteAcceptCmd.Flags().Lookup("token"))
}

func inviteAccept(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	if err != nil {
		return err
	}
	defer datum.StoreSessionCookies(client)

	var s []byte

	token := viper.GetString("invite.accept.token")
	if token == "" {
		return datum.NewRequiredFieldMissingError("token")
	}

	invite := models.InviteRequest{
		Token: token,
	}

	resp, err := client.AcceptInvite(ctx, &invite)
	if err != nil {
		return err
	}

	if err := datum.StoreToken(&oauth2.Token{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
	}); err != nil {
		return err
	}

	s, err = json.Marshal(resp)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
