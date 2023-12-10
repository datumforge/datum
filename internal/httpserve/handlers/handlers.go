package handlers

import "github.com/lestrrat-go/jwx/v2/jwk"

// Handler contains configuration options for handlers including ReadyChecks and JWTKeys
type Handler struct {
	ReadyChecks Checks
	JWTKeys     jwk.Set
}
