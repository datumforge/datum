package handlers

import (
	"net/http"

	echo "github.com/datumforge/echox"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

type Tokens struct {
	Keys jwk.Set
}

// JWKSWellKnownHandler provides the JWK used to verify all Datum-issued JWTs
func (t *Tokens) JWKSWellKnownHandler(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, t.Keys)
}
