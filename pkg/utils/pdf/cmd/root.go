package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "invoice-generator",
	Short: "Generate invoice to PDF",
	Long:  `Generate invoice in PDF format from YAML file`,
}

// Execute root command
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
