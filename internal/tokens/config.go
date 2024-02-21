package tokens

import "time"

// Config defines the configuration settings for authentication tokens used in the server
type Config struct {
	// KID represents the Key ID used in the configuration.
	KID string `json:"kid" koanf:"kid" jsonschema:"required"`
	// Audience represents the target audience for the tokens.
	Audience string `json:"audience" koanf:"audience" jsonschema:"required" default:"https://datum.net"`
	// RefreshAudience represents the audience for refreshing tokens.
	RefreshAudience string `json:"refresh_audience" koanf:"refresh_audience"`
	// Issuer represents the issuer of the tokens
	Issuer string `json:"issuer" koanf:"issuer" jsonschema:"required" default:"https://auth.datum.net" `
	// AccessDuration represents the duration of the access token is valid for
	AccessDuration time.Duration `json:"access_duration" koanf:"access_duration" default:"1h"`
	// RefreshDuration represents the duration of the refresh token is valid for
	RefreshDuration time.Duration `json:"refresh_duration" koanf:"refresh_duration" default:"2h"`
	// RefreshOverlap represents the overlap time for a refresh and access token
	RefreshOverlap time.Duration `json:"refresh_overlap" koanf:"refresh_overlap" default:"-15m" `
	// JWKSEndpoint represents the endpoint for the JSON Web Key Set
	JWKSEndpoint string `json:"jwks_endpoint" koanf:"jwks_endpoint" default:"https://api.datum.net/.well-known/jwks.json"`
	// Keys represents the key pairs used for signing the tokens
	Keys map[string]string `json:"keys" koanf:"keys" jsonschema:"required"`
}
