package cmd

import (
	"strings"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/spf13/cobra"

	"github.com/datumforge/datum/cmd/cli/gencmd"
	"github.com/datumforge/datum/cmd/cli/gencmd/generate/prompts"
)

const (
	relativeSchemaPath = "../../internal/ent/schema"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate is the command to generate the stub files for a given cli cmd",
	Run: func(cmd *cobra.Command, args []string) {
		err := generateStubFiles()
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringP("name", "n", "", "name of the command to generate")
	generateCmd.Flags().StringP("dir", "d", "cmd", "root directory location to generate the files")
	generateCmd.Flags().BoolP("read-only", "r", false, "only generate the read only commands, no create, update or delete commands")
	generateCmd.Flags().BoolP("interactive", "i", true, "interactive prompt, set to false to disable")
	generateCmd.Flags().BoolP("force", "f", false, "force overwrite of existing files")
}

func generateStubFiles() (err error) {
	interactive := Config.Bool("interactive")

	cmdName := Config.String("name")
	hasHistory := false

	if interactive {
		cmdName, err = prompts.Name(cmdName)
		cobra.CheckErr(err)

		if hasHistorySchema(cmdName) {
			hasHistory = prompts.GenerateHistory()
		}
	}

	dirName := Config.String("dir")
	readOnly := Config.Bool("read-only")
	force := Config.Bool("force")

	err = gencmd.Generate(cmdName, dirName, readOnly, force)
	cobra.CheckErr(err)

	if !hasHistory {
		return nil
	}

	return gencmd.Generate(cmdName+"History", dirName, true, force)
}

// hasHistorySchemas loads the schema and checks if the history schema exists
func hasHistorySchema(cmdName string) bool {
	graph, err := entc.LoadGraph(relativeSchemaPath, &gen.Config{})
	cobra.CheckErr(err)

	for _, s := range graph.Schemas {
		if strings.EqualFold(s.Name, cmdName+"history") {
			return true
		}
	}

	return false
}
