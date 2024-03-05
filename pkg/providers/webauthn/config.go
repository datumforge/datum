package webauthn

import (
	"encoding/gob"
	"time"

	"github.com/go-webauthn/webauthn/webauthn"
)

const (
	ProviderName = "WEBAUTHN"
)

// ProviderConfig represents the configuration settings for a Webauthn Provider
type ProviderConfig struct {
	// Enabled is the provider enabled
	Enabled bool `json:"enabled" koanf:"enabled" default:"true"`
	// DisplayName is the site display name
	DisplayName string `json:"display_name" koanf:"display_name" jsonschema:"required" default:"Datum"`
	// RelyingPartyID is the relying party identifier
	// set to localhost for development, no port
	RelyingPartyID string `json:"relying_party_id" koanf:"relying_party_id" jsonschema:"required" default:"localhost"`
	// RequestOrigin the origin domain for authentication requests
	// include the scheme and port
	RequestOrigin string `json:"request_origin" koanf:"request_origin" jsonschema:"required"  default:"http://localhost:3001"`
	// MaxDevices is the maximum number of devices that can be associated with a user
	MaxDevices int `json:"max_devices" koanf:"max_devices" default:"10"`
	// EnforceTimeout at the Relying Party / Server. This means if enabled and the user takes too long that even if the browser does not
	// enforce a timeout, the server will
	EnforceTimeout bool `json:"enforce_timeout" koanf:"enforce" default:"true"`
	// Timeout is the timeout in seconds
	Timeout time.Duration `json:"timeout" koanf:"timeout" default:"60s"`
	// Debug enables debug mode
	Debug bool `json:"debug" koanf:"debug" default:"false"`
}

// NewWithConfig returns a configured Webauthn Provider
func NewWithConfig(config ProviderConfig) *webauthn.WebAuthn {
	if !config.Enabled {
		return nil
	}

	cfg := &webauthn.Config{
		RPID:          config.RelyingPartyID,
		RPOrigin:      config.RequestOrigin,
		RPDisplayName: config.DisplayName,
		Debug:         config.Debug,
		Timeouts: webauthn.TimeoutsConfig{
			Login: webauthn.TimeoutConfig{
				Enforce:    config.EnforceTimeout,
				Timeout:    config.Timeout,
				TimeoutUVD: config.Timeout,
			},
			Registration: webauthn.TimeoutConfig{
				Enforce:    config.EnforceTimeout,
				Timeout:    config.Timeout,
				TimeoutUVD: config.Timeout,
			},
		},
	}

	return &webauthn.WebAuthn{Config: cfg}
}

func init() {
	// Register the webauthn.SessionData type with gob so it can be used in sessions
	gob.Register(webauthn.SessionData{})
}
