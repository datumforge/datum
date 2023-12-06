package datum

import "github.com/spf13/cobra"

// groupCmd represents the base group command when called without any subcommands
var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "The subcommands for working with datum groups",
}

func init() {
	rootCmd.AddCommand(groupCmd)
}
