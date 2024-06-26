package version

import (
	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/internal/constants"
	"github.com/datumforge/datum/pkg/utils/cli/useragent"
)

// VersionCmd is the version command
var cmd = &cobra.Command{
	Use:   "version",
	Short: "print datum CLI version",
	Long:  `The datum version command prints the version of the datum CLI`,
	Run: func(cmd *cobra.Command, _ []string) {
		cmd.Println(constants.VerboseCLIVersion)
		cmd.Printf("User Agent: %s\n", useragent.GetUserAgent())
	},
}

func init() {
	datum.RootCmd.AddCommand(cmd)
}
