package main

import (
	datum "github.com/datumforge/datum/cmd/cli/cmd"

	// since the cmds are no longer part of the same package
	// they must all be imported in main
	_ "github.com/datumforge/datum/cmd/cli/cmd/apitokens"
	_ "github.com/datumforge/datum/cmd/cli/cmd/events"
	_ "github.com/datumforge/datum/cmd/cli/cmd/group"
	_ "github.com/datumforge/datum/cmd/cli/cmd/groupmembers"
	_ "github.com/datumforge/datum/cmd/cli/cmd/groupsetting"
	_ "github.com/datumforge/datum/cmd/cli/cmd/integration"
	_ "github.com/datumforge/datum/cmd/cli/cmd/invite"
	_ "github.com/datumforge/datum/cmd/cli/cmd/login"
	_ "github.com/datumforge/datum/cmd/cli/cmd/org"
	_ "github.com/datumforge/datum/cmd/cli/cmd/orgmembers"
	_ "github.com/datumforge/datum/cmd/cli/cmd/orgsetting"
	_ "github.com/datumforge/datum/cmd/cli/cmd/personalaccesstokens"
	_ "github.com/datumforge/datum/cmd/cli/cmd/register"
	_ "github.com/datumforge/datum/cmd/cli/cmd/reset"
	_ "github.com/datumforge/datum/cmd/cli/cmd/subscriber"
	_ "github.com/datumforge/datum/cmd/cli/cmd/switch"
	_ "github.com/datumforge/datum/cmd/cli/cmd/template"
	_ "github.com/datumforge/datum/cmd/cli/cmd/user"
	_ "github.com/datumforge/datum/cmd/cli/cmd/usersetting"
	_ "github.com/datumforge/datum/cmd/cli/cmd/version"
	_ "github.com/datumforge/datum/cmd/cli/cmd/webhook"
	_ "github.com/datumforge/datum/cmd/cli/cmd/workspace"
)

func main() {
	datum.Execute()
}
