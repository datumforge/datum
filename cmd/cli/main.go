package main

import (
	datum "github.com/datumforge/datum/cmd/cli/cmd"

	// since the cmds are no longer part of the same package
	// they must all be imported in main
	_ "github.com/datumforge/datum/cmd/cli/cmd/group"
	_ "github.com/datumforge/datum/cmd/cli/cmd/groupmembers"
	_ "github.com/datumforge/datum/cmd/cli/cmd/groupsetting"
	_ "github.com/datumforge/datum/cmd/cli/cmd/invite"
	_ "github.com/datumforge/datum/cmd/cli/cmd/login"
	_ "github.com/datumforge/datum/cmd/cli/cmd/org"
	_ "github.com/datumforge/datum/cmd/cli/cmd/orgmembers"
	_ "github.com/datumforge/datum/cmd/cli/cmd/orgsetting"
	_ "github.com/datumforge/datum/cmd/cli/cmd/register"
	_ "github.com/datumforge/datum/cmd/cli/cmd/reset"
	_ "github.com/datumforge/datum/cmd/cli/cmd/subscribers"
	_ "github.com/datumforge/datum/cmd/cli/cmd/template"
	_ "github.com/datumforge/datum/cmd/cli/cmd/tokens"
	_ "github.com/datumforge/datum/cmd/cli/cmd/user"
	_ "github.com/datumforge/datum/cmd/cli/cmd/usersetting"
	_ "github.com/datumforge/datum/cmd/cli/cmd/version"
)

func main() {
	datum.Execute()
}
