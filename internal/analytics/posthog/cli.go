package posthog

import (
	"strings"
	"time"

	"github.com/posthog/posthog-go"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// Capture is intended to be a wrapper around CLI commands to generate generic events associated with the CLI actions
func Capture(command *cobra.Command, userID string) {
	ph, err := posthog.NewWithConfig(PosthogAPIKey,
		posthog.Config{
			Endpoint: PosthogAPIHost,
		},
	)
	if err != nil {
		return
	}
	defer ph.Close()

	properties := ToPosthogProperties()
	properties["command"] = commandName(command)
	flags := []string{}
	command.Flags().VisitAll(func(flag *pflag.Flag) {
		if flag.Changed {
			flags = append(flags, flag.Name)
		}
	})
	properties["flags"] = strings.Join(flags, " ")

	err = ph.Enqueue(posthog.Capture{
		DistinctId: userID,
		Event:      "cli-command-execution",
		Timestamp:  time.Now(),
		Properties: properties,
	})
	if err != nil {
		return
	}
}

func commandName(command *cobra.Command) string {
	if command.HasParent() {
		return commandName(command.Parent()) + " " + command.Name()
	} else {
		return command.Name()
	}
}

func ToPosthogProperties() map[string]interface{} {
	return map[string]interface{}{
		"organization": "",
	}
}
