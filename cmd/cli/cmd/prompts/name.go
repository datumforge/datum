package prompts

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func Name() (string, error) {
	validate := func(input string) error {
		if len(input) == 0 {
			return fmt.Errorf("name cannot be empty")
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
