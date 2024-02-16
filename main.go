//go:generate swagger generate spec

// @title Datum API
// @version 1.0
// @description API Specifications for Datum Services
// @termsOfService http://datum.net/terms/

// @contact.name API Support
// @contact.url http://datum.net/support
// @contact.email support@datum.net

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host datum.net
// @BasePath /v1
package main

import (
	"github.com/datumforge/datum/cmd"
	_ "github.com/datumforge/datum/internal/ent/generated/runtime"
	_ "github.com/datumforge/datum/internal/httpserve/handlers"
	_ "github.com/datumforge/datum/internal/httpserve/route"
	_ "github.com/datumforge/datum/internal/rout"
	_ "github.com/datumforge/datum/internal/utils/responses"
)

func main() {
	cmd.Execute()
}
