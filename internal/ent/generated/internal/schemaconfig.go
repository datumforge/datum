// Code generated by ent, DO NOT EDIT.

package internal

import "context"

// SchemaConfig represents alternative schema names for all tables
// that can be passed at runtime.
type SchemaConfig struct {
	EmailVerificationToken           string // EmailVerificationToken table.
	Entitlement                      string // Entitlement table.
	Group                            string // Group table.
	GroupMembership                  string // GroupMembership table.
	GroupSetting                     string // GroupSetting table.
	Integration                      string // Integration table.
	Invite                           string // Invite table.
	OauthProvider                    string // OauthProvider table.
	OhAuthTooToken                   string // OhAuthTooToken table.
	OrgMembership                    string // OrgMembership table.
	Organization                     string // Organization table.
	OrganizationPersonalAccessTokens string // Organization-personal_access_tokens->PersonalAccessToken table.
	OrganizationSetting              string // OrganizationSetting table.
	PasswordResetToken               string // PasswordResetToken table.
	PersonalAccessToken              string // PersonalAccessToken table.
	Subscriber                       string // Subscriber table.
	TFASettings                      string // TFASettings table.
	Tier                             string // Tier table.
	User                             string // User table.
	UserSetting                      string // UserSetting table.
	Webauthn                         string // Webauthn table.
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
