package config

import (
	"time"
)

type (
	// Config contains the configuration for the datum server
	Config struct {
		// How often to reload the config
		RefreshInterval time.Duration `yaml:"refreshInterval"`

		Auth Auth `yaml:"auth"`
		TLS  TLS  `yaml:"tls"`
		CORS CORS `yaml:"cors"`
	}

	Server struct {
		StubVariable time.Duration `yaml:"stubvariable"`
	}

	Auth struct {
		// Enabled - UI checks this first before reading your provider config
		Enabled bool `yaml:"enabled"`
		// A list of auth providers. Currently enables only the first provider in the list.
		Providers []AuthProvider `yaml:"providers"`
	}

	CORS struct {
		AllowOrigins []string `yaml:"allowOrigins"`
		// CookieInsecure allows CSRF cookie to be sent to servers that the browser considers
		// unsecured. Useful for cases where the connection is secured via VPN rather than
		// HTTPS directly.
		CookieInsecure bool `yaml:"cookieInsecure"`
	}

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
		// Label - optional label for the provider
		Label string `yaml:"label"`
		// Type of the auth provider. Only OIDC is supported today
		Type string `yaml:"type"`
		// OIDC .well-known/openid-configuration URL, ex. https://accounts.google.com/
		ProviderURL string `yaml:"providerUrl"`
		// IssuerUrl - optional. Needed only when differs from the auth provider URL
		IssuerUrl    string `yaml:"issuerUrl"`
		ClientID     string `yaml:"clientId"`
		ClientSecret string `yaml:"clientSecret"`
		// Scopes for auth. Typically [openid, profile, email]
		Scopes []string `yaml:"scopes"`
		// CallbackURL - URL for the callback URL, ex. https://localhost:8080/sso/callback
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
