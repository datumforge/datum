package datumsubscribers

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

var subscribersDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove a subscriber from a datum org",
	RunE: func(cmd *cobra.Command, args []string) error {
		return subscriberDelete(cmd.Context())
	},
}

func init() {
	subscribersCmd.AddCommand(subscribersDeleteCmd)

	subscribersDeleteCmd.Flags().StringP("email", "e", "", "email address of the subscriber to delete")
}

func subscriberDelete(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	if err != nil {
		return err
	}
	defer datum.StoreSessionCookies(client)

	email := datum.Config.String("email")
	if email == "" {
		return datum.NewRequiredFieldMissingError("email")
	}

	var s []byte

	sub, err := client.DeleteSubscriber(ctx, email)
	if err != nil {
		return err
	}

	s, err = json.Marshal(sub)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
