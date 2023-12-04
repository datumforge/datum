package datum

import "github.com/spf13/cobra"

// orgCmd represents the base org command when called without any subcommands
var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "The subcommands for working with datum groups",
}

func init() {
	rootCmd.AddCommand(groupCmd)
}
