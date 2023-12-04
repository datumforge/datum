package config

import (
	"time"
)

type (
	// Config contains the configuration for the datum server
	Config struct {
		// RefreshInterval holds often to reload the config
		RefreshInterval time.Duration `yaml:"refreshInterval"`

		// Auth contains the authentication provider(s)
		Auth Auth `yaml:"auth"`

		// TLS contains the tls configuration settings
		TLS TLS `yaml:"tls"`

		// CORS contains settings to allow cross origin settings and insecure cookies
		CORS CORS `yaml:"cors"`
	}

	// Server settings
	Server struct {
		// StubVariable
		StubVariable time.Duration `yaml:"stubvariable"`
	}

	// Auth settings including providers and the ability to enable/disable auth all together
	Auth struct {
		// Enabled - checks this first before reading your provider config
		Enabled bool `yaml:"enabled"`
		// A list of auth providers. Currently enables only the first provider in the list.
		Providers []AuthProvider `yaml:"providers"`
	}

	// CORS settings
	CORS struct {
		// AllowOrigins is a list of allowed origin to indicate whether the response can be shared with
		// requesting code from the given origin
		AllowOrigins []string `yaml:"allowOrigins"`
		// CookieInsecure allows CSRF cookie to be sent to servers that the browser considers
		// unsecured. Useful for cases where the connection is secured via VPN rather than
		// HTTPS directly.
		CookieInsecure bool `yaml:"cookieInsecure"`
	}

	// TLS settings
	TLS struct {
		CaFile                 string `yaml:"caFile"`
		CertFile               string `yaml:"certFile"`
		KeyFile                string `yaml:"keyFile"`
		CaData                 string `yaml:"caData"`
		CertData               string `yaml:"certData"`
		KeyData                string `yaml:"keyData"`
		EnableHostVerification bool   `yaml:"enableHostVerification"`
		ServerName             string `yaml:"serverName"`
	}

	AuthProvider struct {
		// Label for the provider (optional)
		Label string `yaml:"label"`
		// Type of the auth provider, currently only OIDC is supported
		Type string `yaml:"type"`
		// OIDC .well-known/openid-configuration URL, ex. https://accounts.google.com/
		ProviderURL string `yaml:"providerUrl"`
		// IssuerURL is only needed when it differs from the ProviderURL (optional)
		IssuerURL string `yaml:"issuerUrl"`
		// ClientID of the oauth2 provider
		ClientID string `yaml:"clientId"`
		// ClientSecret is the private key that authenticates your integration when requesting an OAuth token (optional when using PKCE)
		ClientSecret string `yaml:"clientSecret"`
		// Scopes for authentication, typically [openid, profile, email]
		Scopes []string `yaml:"scopes"`
		// CallbackURL after a successful auth, e.g. https://localhost:8080/oauth/callback
		CallbackURL string `yaml:"callbackUrl"`
		// Options added as URL query params when redirecting to auth provider. Can be used to configure custom auth flows such as Auth0 invitation flow.
		Options map[string]interface{} `yaml:"options"`
	}
	//TODO: blow this out to hold all our other configs
)

// Ensure that *Config implements ConfigProvider interface.
var _ ConfigProvider = &Config{}

// GetConfig implements ConfigProvider.
func (c *Config) GetConfig() (*Config, error) {
	return c, nil
}
