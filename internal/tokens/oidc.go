package tokens

// DiscoveryJSON is the response from the OIDC discovery endpoint and contains supported scopes and claims,
// public keys used to sign the tokens, issuer, and other information.
// Clients can use this information to construct a request to the OpenID server.
// The field names and values are defined in the OpenID Connect Discovery Specification
// https://openid.net/specs/openid-connect-discovery-1_0.html
type DiscoveryJSON struct {
	Issuer                            string   `json:"issuer"`
	AuthorizationEndpoint             string   `json:"authorization_endpoint"`
	TokenEndpoint                     string   `json:"token_endpoint"`
	UserinfoEndpoint                  string   `json:"userinfo_endpoint"`
	JwksURI                           string   `json:"jwks_uri"`
	ScopesSupported                   []string `json:"scopes_supported"`
	ResponseTypesSupported            []string `json:"response_types_supported"`
	TokenEndpointAuthMethodsSupported []string `json:"token_endpoint_auth_methods_supported"`
}
