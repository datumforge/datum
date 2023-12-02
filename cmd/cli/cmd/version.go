package datum

import (
	"fmt"

	"github.com/datumforge/datum/cmd/cli/utils"
	"github.com/spf13/cobra"
)

// Print out some basic info about the running binary. Useful for debugging.
var versionCommand = &cobra.Command{ //nolint:gochecknoglobals
	Use: "version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Datum CLI") //nolint:forbidigo

		versionInfo := utils.GetVersionInformation()

		if versionInfo.Stage == utils.Prod {
			fmt.Println("version: " + versionInfo.Version) //nolint:forbidigo
		} else {
			fmt.Println("version: " + versionInfo.Version + " (" + string(versionInfo.Stage) + ")") //nolint:forbidigo
		}

		fmt.Println("build date: " + versionInfo.BuildDate) //nolint:forbidigo
		fmt.Println("commit: " + versionInfo.CommitID)      //nolint:forbidigo
		fmt.Println("branch: " + versionInfo.Branch)        //nolint:forbidigo
	},
}

func init() {
	rootCmd.AddCommand(versionCommand)
}
