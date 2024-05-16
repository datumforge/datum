package prompts

import (
	"github.com/manifoldco/promptui"
)

func Description() (string, error) {
	prompt := promptui.Prompt{
		Label:     "Description (optional):",
		Templates: templates,
	}

	return prompt.Run()
}
