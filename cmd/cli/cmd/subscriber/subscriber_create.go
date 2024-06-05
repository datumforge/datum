package datumsubscribers

import (
	"context"
	"encoding/json"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/utils/cli/tables"
)

var subscribersCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Add subscribers to a datum org",
	RunE: func(cmd *cobra.Command, args []string) error {
		return subscriberCreate(cmd.Context())
	},
}

func init() {
	subscribersCmd.AddCommand(subscribersCreateCmd)

	subscribersCreateCmd.Flags().StringSliceP("emails", "e", []string{}, "email address of the subscriber()")
	datum.ViperBindFlag("subscribers.create.emails", subscribersCreateCmd.Flags().Lookup("emails"))
}

func subscriberCreate(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	if err != nil {
		return err
	}
	defer datum.StoreSessionCookies(client)

	subscriberInput := []*datumclient.CreateSubscriberInput{}

	if datum.InputFile != "" {
		input, err := os.OpenFile(datum.InputFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
		if err != nil {
			panic(err)
		}
		defer input.Close()

		in := graphql.Upload{
			File: input,
		}

		sub, err := client.CreateBulkCSVSubscriber(ctx, in)
		if err != nil {
			return err
		}

		if datum.OutputFormat == datum.JSONOutput {
			s, err := json.Marshal(sub.CreateBulkCSVSubscriber)
			if err != nil {
				return err
			}

			return datum.JSONPrint(s)
		}

		writer := tables.NewTableWriter(subscribersCmd.OutOrStdout(), "Email", "Verified", "Active")
		for _, s := range sub.CreateBulkCSVSubscriber.Subscribers {
			writer.AddRow(s.Email, s.VerifiedEmail, s.Active)
		}

		writer.Render()
	} else {
		email := viper.GetStringSlice("subscribers.create.emails")
		if len(email) == 0 {
			return datum.NewRequiredFieldMissingError("emails")
		}

		for _, e := range email {
			subscriberInput = append(subscriberInput, &datumclient.CreateSubscriberInput{
				Email: e,
			})
		}

		sub, err := client.CreateBulkSubscriber(ctx, subscriberInput)
		if err != nil {
			return err
		}

		if datum.OutputFormat == datum.JSONOutput {
			s, err := json.Marshal(sub.CreateBulkSubscriber)
			if err != nil {
				return err
			}

			return datum.JSONPrint(s)
		}

		writer := tables.NewTableWriter(subscribersCmd.OutOrStdout(), "Email", "Verified", "Active")
		for _, s := range sub.CreateBulkSubscriber.Subscribers {
			writer.AddRow(s.Email, s.VerifiedEmail, s.Active)
		}

		writer.Render()
	}

	return nil
}
