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
			Comment("domains associated with the organization").
			Optional(),
		field.Text("sso_cert").
			Optional(),
		field.String("sso_entrypoint").
			Optional(),
		field.String("sso_issuer").
			Optional(),
		field.String("billing_contact").
			Comment("Name of the person to contact for billing").
			Optional(),
		field.String("billing_email").
			Optional(),
		field.String("billing_phone").
			Optional(),
		field.String("billing_address").
			Optional(),
		field.String("tax_identifier").
			Comment("Usually government-issued tax ID or business ID such as ABN in Australia").
			Optional(),
		field.JSON("tags", []string{}).
			Comment("tags associated with the object").
			Default([]string{}).
			Optional(),
	}
}

// Edges of the OrganizationSetting
func (OrganizationSetting) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("organization", Organization.Type).Ref("setting").Unique(),
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
		mixin.SoftDeleteMixin{},
	}
}
