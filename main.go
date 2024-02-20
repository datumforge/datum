// package main is the entry point
package main

import (
	"github.com/datumforge/datum/cmd"
	_ "github.com/datumforge/datum/internal/ent/generated/runtime"
)

func main() {
	cmd.Execute()
}
