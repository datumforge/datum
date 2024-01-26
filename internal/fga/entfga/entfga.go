package entfga

import (
	"embed"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

var (
	//go:embed templates/*
	_templates embed.FS
)

type Config struct {
}

func (c Config) Name() string {
	return "AuthzConfig"
}

// AuthzExtension implements entc.Extension.
type AuthzExtension struct {
	entc.DefaultExtension
	config *Config
}

type ExtensionOption = func(*AuthzExtension)

// NewFGAExtension creates a new fga extension
func NewFGAExtension(opts ...ExtensionOption) *AuthzExtension {
	extension := &AuthzExtension{
		// Set configuration defaults that can get overridden with ExtensionOption
		config: &Config{},
	}

	// for _, opt := range opts {
	// 	opt(extension)
	// }

	return extension
}

// Templates returns the generated templates which include the client, history query, history from mutation
// and an optional auditing template
func (h *AuthzExtension) Templates() []*gen.Template {
	templates := []*gen.Template{
		parseTemplate("authzFromMutation", "templates/authzFromMutation.tmpl"),
	}

	return templates
}
