package datumsearch

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/utils/cli/tables"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for organizations, groups, users, subscribers, etc in the datum system",
	RunE: func(cmd *cobra.Command, args []string) error {
		return search(cmd.Context())
	},
}

func init() {
	datum.RootCmd.AddCommand(searchCmd)

	searchCmd.Flags().StringP("query", "q", "", "query string to search for")
}

func search(ctx context.Context) error { // setup datum http client
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	// filter options
	query := datum.Config.String("query")
	if query == "" {
		return datum.NewRequiredFieldMissingError("query")
	}

	results, err := client.Search(ctx, query)
	cobra.CheckErr(err)

	// print results
	for _, r := range results.Search.Nodes {
		if len(r.OrganizationSearchResult.Organizations) > 0 {
			fmt.Println("Organization Results")

			writer := tables.NewTableWriter(datum.RootCmd.OutOrStdout(), "ID", "Name", "DisplayName", "Description")

			for _, o := range r.OrganizationSearchResult.Organizations {
				writer.AddRow(o.ID, o.Name, o.DisplayName, *o.Description)
			}

			writer.Render()
		}

		if len(r.GroupSearchResult.Groups) > 0 {
			fmt.Println("Group Results")

			writer := tables.NewTableWriter(datum.RootCmd.OutOrStdout(), "ID", "Name", "DisplayName", "Description")

			for _, g := range r.GroupSearchResult.Groups {
				writer.AddRow(g.ID, g.Name, g.DisplayName, *g.Description)
			}

			writer.Render()
		}

		if len(r.UserSearchResult.Users) > 0 {
			fmt.Println("User Results")

			writer := tables.NewTableWriter(datum.RootCmd.OutOrStdout(), "ID", "FirstName", "LastName", "DisplayName", "Email")

			for _, u := range r.UserSearchResult.Users {
				writer.AddRow(u.ID, *u.FirstName, *u.LastName, u.DisplayName, u.Email)
			}

			writer.Render()
		}

		if len(r.SubscriberSearchResult.Subscribers) > 0 {
			fmt.Println("Subscriber Results")

			writer := tables.NewTableWriter(datum.RootCmd.OutOrStdout(), "ID", "Email")

			for _, s := range r.SubscriberSearchResult.Subscribers {
				writer.AddRow(s.ID, s.Email)
			}

			writer.Render()
		}
	}

	return nil
}
