package datumsubscribers

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/utils/cli/tables"
)

var subscribersGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get existing subscribers of a datum org",
	RunE: func(cmd *cobra.Command, args []string) error {
		return subscribers(cmd.Context())
	},
}

func init() {
	subscribersCmd.AddCommand(subscribersGetCmd)

	subscribersGetCmd.Flags().BoolP("active", "a", true, "filter on active subscribers")
	subscribersGetCmd.Flags().StringP("email", "e", "", "email address of the subscriber to get")
}

func subscribers(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	// filter options
	where := datumclient.SubscriberWhereInput{}

	active := datum.Config.Bool("active")
	if active {
		where.Active = &active
	}

	email := datum.Config.String("email")

	writer := tables.NewTableWriter(subscribersCmd.OutOrStdout(), "Email", "Verified", "Active")

	if email != "" {
		sub, err := client.GetSubscriber(ctx, email)
		cobra.CheckErr(err)

		if datum.OutputFormat == datum.JSONOutput {
			s, err := json.Marshal(sub)
			if err != nil {
				return err
			}

			return datum.JSONPrint(s)
		} else {
			writer.AddRow(sub.Subscriber.Email, sub.Subscriber.VerifiedEmail, sub.Subscriber.Active)

			writer.Render()
		}
	} else {
		subs, err := client.Subscribers(ctx, &where)
		cobra.CheckErr(err)

		if datum.OutputFormat == datum.JSONOutput {
			s, err := json.Marshal(subs)
			if err != nil {
				return err
			}

			return datum.JSONPrint(s)
		} else {
			for _, s := range subs.Subscribers.Edges {
				writer.AddRow(s.Node.Email, s.Node.VerifiedEmail, s.Node.Active)
			}

			writer.Render()
		}
	}

	return nil
}
