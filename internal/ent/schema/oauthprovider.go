package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/datumforge/datum/internal/ent/mixin"
)

// OauthProvider holds the schema definition for the OauthProvider entity
type OauthProvider struct {
	ent.Schema
}

// Fields of the OauthProvider
func (OauthProvider) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("the provider's name"),
		field.String("client_id").Comment("the client id"),
		field.String("client_secret").Sensitive().Comment("the client secret"),
		field.String("redirect_url").Comment("the redirect url"),
		field.String("scopes").Comment("the scopes"),
		field.String("auth_url").Comment("the auth url of the provider"),
		field.String("token_url").Comment("the token url of the provider"),
		field.Uint8("auth_style").Comment("the auth style, 0: auto detect 1: third party log in 2: log in with username and password"),
		field.String("info_url").Comment("the URL to request user information by token"),
	}
}

// Edges of the OauthProvider
func (OauthProvider) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Organization.Type).Ref("oauthprovider").Unique(),
	}
}

// Annotations of the OauthProvider
func (OauthProvider) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
	}
}

// Mixin of the OauthProvider
func (OauthProvider) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.IDMixin{},
	}
}
