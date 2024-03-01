package webauthn

// ProviderConfig represents the configuration settings for a Github Oauth Provider
type ProviderConfig struct {
	// DisplayName is the site display name
	DisplayName string `json:"display_name" koanf:"display_name" jsonschema:"required"`
	// Domain is the domain of the site
	Domain string `json:"domain" koanf:"domain" jsonschema:"required" default:"datum.net"`
	// RequestOrigin the origin domain for authentication requests
	RequestOrigin string `json:"request_origin" koanf:"request_origin" jsonschema:"required"`
	// MaxDevices is the maximum number of devices that can be associated with a user
	MaxDevices int `json:"max_devices" koanf:"max_devices" jsonschema:"required"`
}
