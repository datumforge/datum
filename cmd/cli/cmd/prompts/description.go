package prompts

import (
	"github.com/manifoldco/promptui"
)

func Domains() (string, error) {
	prompt := promptui.Prompt{
		Label:     "Domains (optional):",
		Templates: templates,
	}

	return prompt.Run()
}
