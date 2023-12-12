package handlers

import (
	"github.com/lestrrat-go/jwx/v2/jwk"

	ent "github.com/datumforge/datum/internal/ent/generated"
)

// Handler contains configuration options for handlers including ReadyChecks and JWTKeys
type Handler struct {
	// DBClient to interact with the generated ent schema
	DBClient    *ent.Client
	ReadyChecks Checks
	JWTKeys     jwk.Set
}
