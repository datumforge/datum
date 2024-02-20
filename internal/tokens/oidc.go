package tokens

// DiscoveryJSON is the response from the OIDC discovery endpoint and contains supported scopes and claims,
// public keys used to sign the tokens, issuer, and other information.
// Clients can use this information to construct a request to the OpenID server.
// The field names and values are defined in the OpenID Connect Discovery Specification
// https://openid.net/specs/openid-connect-discovery-1_0.html
type DiscoveryJSON struct {
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
