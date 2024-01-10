package handlers

import (
	"net/http"

	echo "github.com/datumforge/echox"
)

// SecurityHandler provides the JWK used to verify all Datum-issued JWTs
func (h *Handler) SecurityHandler(ctx echo.Context) error {
	return ctx.String(http.StatusOK, securityTxt)
}

const securityTxt = `-----BEGIN PGP SIGNED MESSAGE-----
matt is the best!
-----END PGP SIGNATURE-----
`
