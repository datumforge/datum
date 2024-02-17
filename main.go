//go:generate swagger generate spec
package main

import (
	"github.com/datumforge/datum/cmd"
	_ "github.com/datumforge/datum/internal/ent/generated/runtime"
	_ "github.com/datumforge/datum/openapi"
)

func main() {
	cmd.Execute()
}
