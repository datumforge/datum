package handlers

import (
	"net/http"
	"net/url"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/rout"
	"github.com/datumforge/datum/internal/tokens"
)

// OpenIDConfiguration returns a JSON document with the OpenID configuration as defined by the OpenID
// Connect standard: https://connect2id.com/learn/openid-connect. This document helps
// clients understand how to authenticate with Datum.
func (h *Handler) OpenIDConfiguration(ctx echo.Context) error {
	// Parse the token issuer for the OpenID configuration
	base, err := url.Parse(h.TM.Config().Issuer)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse("openid is not configured correctly"))
	}

	openid := &tokens.DiscoveryJSON{
		Issuer:                        base.ResolveReference(&url.URL{Path: "/"}).String(),
		JWKSURI:                       base.ResolveReference(&url.URL{Path: "/.well-known/jwks.json"}).String(),
		AuthorizationEP:               base.ResolveReference(&url.URL{Path: "/oauth/authorize"}).String(),
		TokenEP:                       base.ResolveReference(&url.URL{Path: "/oauth/token"}).String(),
		UserInfoEP:                    base.ResolveReference(&url.URL{Path: "/oauth/userinfo"}).String(),
		ScopesSupported:               []string{"openid", "profile", "email"},
		ResponseTypesSupported:        []string{"code", "token", "id_token"},
		TokenEndpointAuthMethods:      []string{"client_secret_basic", "client_secret_post"},
		CodeChallengeMethodsSupported: []string{"S256", "plain"},
		ResponseModesSupported:        []string{"query", "fragment", "form_post"},
		SubjectTypesSupported:         []string{"public"},
		IDTokenSigningAlgValues:       []string{"HS256", "RS256"},
		ClaimsSupported:               []string{"aud", "email", "exp", "iat", "iss", "sub"},
		RequestURIParameterSupported:  false,
	}

	return ctx.JSON(http.StatusOK, openid)
}
