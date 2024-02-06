package handlers

import (
	"net/http"
	"net/url"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/tokens"
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
	return ctx.JSON(http.StatusOK, "User-agent: *\n"+"Disallow: /")
}

// JWKSWellKnownHandler provides the JWK used to verify all Datum-issued JWTs
func (h *Handler) JWKSWellKnownHandler(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, h.JWTKeys)
}

// OpenIDConfiguration returns a JSON document with the OpenID configuration as defined by the OpenID
// Connect standard: https://connect2id.com/learn/openid-connect. This document helps
// clients understand how to authenticate with Datum.
func (h *Handler) OpenIDConfiguration(ctx echo.Context) error {
	// Parse the token issuer for the OpenID configuration
	base, err := url.Parse(h.TM.Config().Issuer)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse("openid is not configured correctly"))
	}

	openid := &tokens.DiscoveryJSON{
		Issuer:                base.ResolveReference(&url.URL{Path: "/"}).String(),
		JwksURI:               base.ResolveReference(&url.URL{Path: "/.well-known/jwks.json"}).String(),
		AuthorizationEndpoint: base.ResolveReference(&url.URL{Path: "/authorize"}).String(),
		TokenEndpoint:         base.ResolveReference(&url.URL{Path: "/token"}).String(),
		// TODO: add querystring for user query
		UserinfoEndpoint:                  base.ResolveReference(&url.URL{Path: "/query"}).String(),
		ScopesSupported:                   []string{"openid", "profile", "email"},
		ResponseTypesSupported:            []string{"code", "token", "id_token"},
		TokenEndpointAuthMethodsSupported: []string{"client_secret_basic", "client_secret_post"},
	}

	return ctx.JSON(http.StatusOK, openid)
}
