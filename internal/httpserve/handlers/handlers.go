package handlers

import (
	"github.com/lestrrat-go/jwx/v2/jwk"
	"go.uber.org/zap"

	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/tokens"
)

// Handler contains configuration options for handlers including ReadyChecks and JWTKeys
type Handler struct {
	// DBClient to interact with the generated ent schema
	DBClient     *ent.Client
	TM           *tokens.TokenManager
	CookieDomain string
	Logger       *zap.SugaredLogger
	ReadyChecks  Checks
	JWTKeys      jwk.Set
}
