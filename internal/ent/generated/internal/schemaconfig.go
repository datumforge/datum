// Code generated by ent, DO NOT EDIT.

package internal

import "context"

// SchemaConfig represents alternative schema names for all tables
// that can be passed at runtime.
type SchemaConfig struct {
	Group                string // Group table.
	GroupUsers           string // Group-users->User table.
	GroupSettings        string // GroupSettings table.
	Integration          string // Integration table.
	Organization         string // Organization table.
	OrganizationSettings string // OrganizationSettings table.
	RefreshToken         string // RefreshToken table.
	Session              string // Session table.
	Subscription         string // Subscription table.
	User                 string // User table.
	UserOrganizations    string // User-organizations->Organization table.
}

type schemaCtxKey struct{}

// SchemaConfigFromContext returns a SchemaConfig stored inside a context, or empty if there isn't one.
func SchemaConfigFromContext(ctx context.Context) SchemaConfig {
	config, _ := ctx.Value(schemaCtxKey{}).(SchemaConfig)
	return config
}

// NewSchemaConfigContext returns a new context with the given SchemaConfig attached.
func NewSchemaConfigContext(parent context.Context, config SchemaConfig) context.Context {
	return context.WithValue(parent, schemaCtxKey{}, config)
}
