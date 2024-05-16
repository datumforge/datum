package prompts

import (
	"github.com/manifoldco/promptui"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

func Name() (string, error) {
	validate := func(input string) error {
		if len(input) == 0 {
			return datum.NewRequiredFieldMissingError("name")
		}

		return nil
	}

	prompt := promptui.Prompt{
		Label:     "Name:",
		Templates: templates,
		Validate:  validate,
	}

	return prompt.Run()
}
