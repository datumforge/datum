package main

import (
	datum "github.com/datumforge/datum/cmd/cli/cmd"

	// since the cmds are not part of the same package
	// they must all be imported in main
	_ "github.com/datumforge/datum/cmd/cli/cmd/apitokens"
	_ "github.com/datumforge/datum/cmd/cli/cmd/contact"
	_ "github.com/datumforge/datum/cmd/cli/cmd/entitlementplan"
	_ "github.com/datumforge/datum/cmd/cli/cmd/entitlementplanfeatures"
	_ "github.com/datumforge/datum/cmd/cli/cmd/entitlements"
	_ "github.com/datumforge/datum/cmd/cli/cmd/entity"
	_ "github.com/datumforge/datum/cmd/cli/cmd/entitytype"
	_ "github.com/datumforge/datum/cmd/cli/cmd/events"
	_ "github.com/datumforge/datum/cmd/cli/cmd/features"
	_ "github.com/datumforge/datum/cmd/cli/cmd/group"
	_ "github.com/datumforge/datum/cmd/cli/cmd/groupmembers"
	_ "github.com/datumforge/datum/cmd/cli/cmd/groupsetting"
	_ "github.com/datumforge/datum/cmd/cli/cmd/integration"
	_ "github.com/datumforge/datum/cmd/cli/cmd/invite"
	_ "github.com/datumforge/datum/cmd/cli/cmd/login"
	_ "github.com/datumforge/datum/cmd/cli/cmd/organization"
	_ "github.com/datumforge/datum/cmd/cli/cmd/organizationsetting"
	_ "github.com/datumforge/datum/cmd/cli/cmd/orgmembers"
	_ "github.com/datumforge/datum/cmd/cli/cmd/personalaccesstokens"
	_ "github.com/datumforge/datum/cmd/cli/cmd/register"
	_ "github.com/datumforge/datum/cmd/cli/cmd/reset"
	_ "github.com/datumforge/datum/cmd/cli/cmd/search"
	_ "github.com/datumforge/datum/cmd/cli/cmd/subscriber"
	_ "github.com/datumforge/datum/cmd/cli/cmd/switch"
	_ "github.com/datumforge/datum/cmd/cli/cmd/template"
	_ "github.com/datumforge/datum/cmd/cli/cmd/user"
	_ "github.com/datumforge/datum/cmd/cli/cmd/usersetting"
	_ "github.com/datumforge/datum/cmd/cli/cmd/version"
	_ "github.com/datumforge/datum/cmd/cli/cmd/webhook"

	// history commands
	_ "github.com/datumforge/datum/cmd/cli/cmd/documentdatahistory"
	_ "github.com/datumforge/datum/cmd/cli/cmd/entitlementhistory"
	_ "github.com/datumforge/datum/cmd/cli/cmd/entityhistory"
	_ "github.com/datumforge/datum/cmd/cli/cmd/entitytypehistory"
	_ "github.com/datumforge/datum/cmd/cli/cmd/eventhistory"
	_ "github.com/datumforge/datum/cmd/cli/cmd/featurehistory"
	_ "github.com/datumforge/datum/cmd/cli/cmd/filehistory"
	_ "github.com/datumforge/datum/cmd/cli/cmd/grouphistory"
	_ "github.com/datumforge/datum/cmd/cli/cmd/groupmembershiphistory"
	_ "github.com/datumforge/datum/cmd/cli/cmd/groupsettinghistory"
	_ "github.com/datumforge/datum/cmd/cli/cmd/hushhistory"
	_ "github.com/datumforge/datum/cmd/cli/cmd/integrationhistory"
	_ "github.com/datumforge/datum/cmd/cli/cmd/oauthproviderhistory"
	_ "github.com/datumforge/datum/cmd/cli/cmd/organizationhistory"
	_ "github.com/datumforge/datum/cmd/cli/cmd/organizationsettinghistory"
	_ "github.com/datumforge/datum/cmd/cli/cmd/orgmembershiphistory"
	_ "github.com/datumforge/datum/cmd/cli/cmd/templatehistory"
	_ "github.com/datumforge/datum/cmd/cli/cmd/userhistory"
	_ "github.com/datumforge/datum/cmd/cli/cmd/usersettinghistory"
	_ "github.com/datumforge/datum/cmd/cli/cmd/webhookhistory"
)

func main() {
	datum.Execute()
}
