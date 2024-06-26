package datumsubscribers

import (
	"context"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "add subscribers to a datum organization",
	Run: func(cmd *cobra.Command, args []string) {
		err := create(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(createCmd)

	createCmd.Flags().StringSliceP("emails", "e", []string{}, "email address of the subscriber()")
}

// createValidation validates the required fields for the command
func createValidation() (input []*datumclient.CreateSubscriberInput, err error) {
	email := datum.Config.Strings("emails")
	if len(email) == 0 {
		return input, datum.NewRequiredFieldMissingError("emails")
	}

	for _, e := range email {
		input = append(input, &datumclient.CreateSubscriberInput{
			Email: e,
		})
	}

	return input, nil
}

func create(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	if datum.InputFile != "" {
		input, err := os.OpenFile(datum.InputFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
		cobra.CheckErr(err)

		defer input.Close()

		in := graphql.Upload{
			File: input,
		}

		o, err := client.CreateBulkCSVSubscriber(ctx, in)
		cobra.CheckErr(err)

		return consoleOutput(o)
	}

	input, err := createValidation()
	cobra.CheckErr(err)

	o, err := client.CreateBulkSubscriber(ctx, input)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
