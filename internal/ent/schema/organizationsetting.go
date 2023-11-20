package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/datumforge/datum/internal/ent/mixin"
)

// OrganizationSetting holds the schema definition for the OrganizationSetting entity
type OrganizationSetting struct {
	ent.Schema
}

// Fields of the OrganizationSetting
func (OrganizationSetting) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("domains", []string{}).
			Comment("domains associated with the organization"),
		field.Text("sso_cert").
			Default(""),
		field.String("sso_entrypoint").
			Default(""),
		field.String("sso_issuer").
			Default(""),
		field.String("billing_contact").
			NotEmpty().
			Comment("Name of the person to contact for billing"),
		field.String("billing_email").
			NotEmpty(),
		field.String("billing_phone").
			NotEmpty(),
		field.String("billing_address").
			NotEmpty(),
		field.String("tax_identifier").
			Comment("Usually government-issued tax ID or business ID such as ABN in Australia"),
		field.JSON("tags", []string{}).
			Comment("tags associated with the object").
			Default([]string{}).
			Optional(),
	}
}

// Edges of the OrganizationSetting
func (OrganizationSetting) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("orgnaization", Organization.Type).Ref("setting").Unique(),
	}
}

// Annotations of the OrganizationSetting
func (OrganizationSetting) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
	}
}

// Mixin of the OrganizationSetting
func (OrganizationSetting) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.IDMixin{},
	}
}
