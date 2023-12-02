package utils

import (
	"os"

	"github.com/datumforge/datum/cmd/cli/logger"
)

func GetWorkingDir() string {
	workingDir, err := os.Getwd()
	if err != nil {
		logger.FatalErr("Unable to get working directory", err)

		return ""
	}

	return workingDir
}
