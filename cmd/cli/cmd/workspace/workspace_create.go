package datumworkspace

import (
	"context"
	"fmt"
	"strings"

	"github.com/ermineaweb/pbar"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/cmd/cli/cmd/prompts"
	datumswitch "github.com/datumforge/datum/cmd/cli/cmd/switch"
	"github.com/datumforge/datum/pkg/datumclient"
)

var workspaceCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new datum org",
	RunE: func(cmd *cobra.Command, args []string) error {
		return createWorkspace(cmd.Context())
	},
}

func init() {
	workspaceCmd.AddCommand(workspaceCreateCmd)

	workspaceCreateCmd.Flags().StringP("name", "n", "", "name of the workspace")
	datum.ViperBindFlag("workspace.create.name", workspaceCreateCmd.Flags().Lookup("name"))

	workspaceCreateCmd.Flags().StringP("description", "d", "", "description of the workspace")
	datum.ViperBindFlag("workspace.create.description", workspaceCreateCmd.Flags().Lookup("description"))

	workspaceCreateCmd.Flags().StringSlice("domains", []string{}, "domains associated with the workspace")
	datum.ViperBindFlag("workspace.create.domains", workspaceCreateCmd.Flags().Lookup("domains"))

	workspaceCreateCmd.Flags().BoolP("interactive", "i", true, "interactive prompt, set to false to disable")
	datum.ViperBindFlag("workspace.create.interactive", workspaceCreateCmd.Flags().Lookup("interactive"))
}

func createWorkspace(ctx context.Context) error {
	// setup datum http client
	cli, err := datum.GetGraphClient(ctx)
	if err != nil {
		return err
	}

	// check if interactive flag is set, if it is, an some params are not set, prompt user for input
	interactive := viper.GetBool("workspace.create.interactive")

	name := viper.GetString("workspace.create.name")
	if name == "" && interactive {
		name, err = prompts.Name()
		if err != nil {
			return err
		}
	}

	description := viper.GetString("org.create.description")
	if description == "" && interactive {
		description, err = prompts.Description()
		if err != nil {
			return err
		}
	}

	domains := viper.GetStringSlice("org.create.domains")
	if len(domains) == 0 && interactive {
		domainString, err := prompts.Domains()
		if err != nil {
			return err
		}

		domains = strings.Split(domainString, ",")
	}

	input := datumclient.CreateWorkspaceInput{
		Name: name,
	}

	if description != "" {
		input.Description = &description
	}

	if len(domains) > 0 {
		input.Domains = domains
	}

	bar := pbar.NewCustomSpinner(pbar.ConfigSpinner{
		Spinner:      pbar.SPINNER_SNAKE,
		StartMessage: "creating workspace...",
		StopMessage:  "workspace created",
	})
	bar.Start()

	ws, err := cli.Client.CreateWorkspace(ctx, input, cli.Interceptor)
	if err != nil {
		return err
	}

	bar.Stop()

	workspace := ws.CreateWorkspace.Workspace

	fmt.Println("ID: ", workspace.ID)
	fmt.Println("Name: ", workspace.DisplayName)
	fmt.Println("Description: ", *workspace.Description)
	fmt.Println("Domains: ", strings.Join(workspace.Setting.Domains, ","))

	// switch to new workspace
	datumswitch.SwitchOrg(ctx, workspace.ID)
	fmt.Println("\nSwitched to workspace: ", workspace.DisplayName)

	return nil
}
