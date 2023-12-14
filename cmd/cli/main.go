package main

import (
	datum "github.com/datumforge/datum/cmd/cli/cmd"

	// since the cmds are no longer part of the same package
	// they must all be imported in main
	_ "github.com/datumforge/datum/cmd/cli/cmd/group"
	_ "github.com/datumforge/datum/cmd/cli/cmd/login"
	_ "github.com/datumforge/datum/cmd/cli/cmd/org"
	_ "github.com/datumforge/datum/cmd/cli/cmd/tokens"
	_ "github.com/datumforge/datum/cmd/cli/cmd/user"
)

func main() {
	datum.Execute()
}
