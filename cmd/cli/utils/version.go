package utils

import "github.com/datumforge/datum/cmd/cli/vars"

type VersionInformation struct {
	Version   string
	BuildDate string
	CommitID  string
	Branch    string
	Stage     Stage
}

type Stage string

const (
	Prod Stage = "prod"
)

func GetVersionInformation() VersionInformation {
	return VersionInformation{
		Version:   vars.Version,
		BuildDate: vars.BuildDate,
		CommitID:  vars.CommitID,
		Branch:    vars.Branch,
		Stage:     Stage(vars.Stage),
	}
}
