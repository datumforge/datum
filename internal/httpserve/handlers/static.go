package handlers

import (
	"net/http"

	echo "github.com/datumforge/echox"
)

// SecurityHandler hosts the /security.txt endpoint https://securitytxt.org/, signed with our GPG key
func (h *Handler) SecurityHandler(ctx echo.Context) error {
	return ctx.String(http.StatusOK, securityTxt)
}

const securityTxt = `-----BEGIN PGP SIGNED MESSAGE-----
Hash: SHA512

Contact: mailto:security@datum.net
Expires: 2026-01-09T12:34:00.000Z
Encryption: https://keys.openpgp.org/vks/v1/by-fingerprint/CDD6CB47F5C8C90340CEB56F5024ED84F6BF803F
Preferred-Languages: en
Canonical: https://api.datum.net/.well-known/security.txt
Policy: https://github.com/datumforge/datum/security/policy
-----BEGIN PGP SIGNATURE-----

iHUEARYKAB0WIQTN1stH9cjJA0DOtW9QJO2E9r+APwUCZZ6RLgAKCRBQJO2E9r+A
PyjIAQC13xI75wq7o4MhparBplTo6ZCF+bJTyCrO5c5izag9IQD+KbAXRKq3pjXu
kwa/7CNQAyf7R//ZzA2npg2Ly5Jv4Qs=
=Uh5X
-----END PGP SIGNATURE-----
`

// RobotsHandler # https://www.robotstxt.org/robotstxt.html
func (h *Handler) RobotsHandler(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "User-agent: *\n"+"Disallow: /")
}

// JWKSWellKnownHandler provides the JWK used to verify all Datum-issued JWTs
func (h *Handler) JWKSWellKnownHandler(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, h.JWTKeys)
}
