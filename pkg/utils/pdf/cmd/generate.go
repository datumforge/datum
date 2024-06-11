package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/datumforge/datum/pkg/utils/pdf/invoice"
)

// generaterCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate invoice",
	Long:  `Generate invoice to PDF file from YAML`,
	RunE: func(cmd *cobra.Command, args []string) error {
		yamlPath, _ := cmd.Flags().GetString("yaml")
		outputPath, err := cmd.Flags().GetString("out")
		if err != nil {
			fmt.Println("Could not get the value of flag", err)
			os.Exit(1)
		}
		err = generate(yamlPath, outputPath)

		return err
	},
}

// init wraps the cobra cmds - this probably won't stay as a CLI (maybe it will?) so keeping it basic with options for now
func init() {
	rootCmd.AddCommand(generateCmd)
	rootCmd.PersistentFlags().String("yaml", "invoice.yaml", "Path to the YAML file with invoice parameters")
	rootCmd.PersistentFlags().String("out", "invoice.pdf", "Path to where should it save pdf file")
}

// generate will take an input yaml and hopefully spit out a beatuiful PDF to the provided file
func generate(sourcePath string, outputPath string) error {
	file, err := os.ReadFile(sourcePath)
	if err != nil {
		return fmt.Errorf("could not read the file: %s", err) // nolint: goerr113
	}

	inv, err := invoice.NewInvoice(file)
	if err != nil {
		return fmt.Errorf("could not prepare the invoice: %s", err) // nolint: goerr113
	}

	err = inv.SaveToPdf(outputPath)
	if err != nil {
		return fmt.Errorf("could not save PDF: %s", err) // nolint: goerr113
	}

	return err
}
