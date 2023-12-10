package handlers

import (
	"net/http"
	"net/url"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/httpserve/middleware/auth"
	"github.com/datumforge/datum/internal/tokens"
)

type OpenIDConfiguration struct {
	Issuer                        string   `json:"issuer"`
	AuthorizationEP               string   `json:"authorization_endpoint"`
	TokenEP                       string   `json:"token_endpoint"`
	DeviceAuthorizationEP         string   `json:"device_authorization_endpoint"`
	UserInfoEP                    string   `json:"userinfo_endpoint"`
	MFAChallengeEP                string   `json:"mfa_challenge_endpoint"`
	JWKSURI                       string   `json:"jwks_uri"`
	RegistrationEP                string   `json:"registration_endpoint"`
	RevocationEP                  string   `json:"revocation_endpoint"`
	ScopesSupported               []string `json:"scopes_supported"`
	ResponseTypesSupported        []string `json:"response_types_supported"`
	CodeChallengeMethodsSupported []string `json:"code_challenge_methods_supported"`
	ResponseModesSupported        []string `json:"response_modes_supported"`
	SubjectTypesSupported         []string `json:"subject_types_supported"`
	IDTokenSigningAlgValues       []string `json:"id_token_signing_alg_values_supported"`
	TokenEndpointAuthMethods      []string `json:"token_endpoint_auth_methods_supported"`
	ClaimsSupported               []string `json:"claims_supported"`
	RequestURIParameterSupported  bool     `json:"request_uri_parameter_supported"`
}

// JWKSWellKnownHandler provides the JWK used to verify all Datum-issued JWTs
func JWKSWellKnownHandler(ctx echo.Context) error {

	// Setup Token manager
	conf := tokens.TokenConfig{
		Issuer:   "http://localhost:17608",
		Audience: "http://localhost:17608",
	}

	tm, err := tokens.New(conf)
	if err != nil {
		return err
	}

	tm.Keys()

	base, err := url.Parse("http://localhost:17608")
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, auth.ErrorResponse("openid is not configured correctly"))
	}

	openid := &OpenIDConfiguration{
		Issuer:                        base.ResolveReference(&url.URL{Path: "/"}).String(),
		JWKSURI:                       base.ResolveReference(&url.URL{Path: "/.well-known/jwks.json"}).String(),
		ScopesSupported:               []string{"openid", "profile", "email"},
		ResponseTypesSupported:        []string{"token", "id_token"},
		CodeChallengeMethodsSupported: []string{"S256", "plain"},
		ResponseModesSupported:        []string{"query", "fragment", "form_post"},
		SubjectTypesSupported:         []string{"public"},
		IDTokenSigningAlgValues:       []string{"HS256", "RS256"},
		TokenEndpointAuthMethods:      []string{"client_secret_basic", "client_secret_post"},
		ClaimsSupported:               []string{"aud", "email", "exp", "iat", "iss", "sub"},
		RequestURIParameterSupported:  false,
	}

	return ctx.JSON(http.StatusOK, openid)
}
