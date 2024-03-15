package schema

import (
	"net/mail"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/datumforge/datum/internal/ent/mixin"
)

// Subscribers holds the schema definition for the Subscribers entity
type Subscribers struct {
	ent.Schema
}

// Fields of the Subscribers
func (Subscribers) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").
			Comment("email address of the subscriber").
			Validate(func(email string) error {
				_, err := mail.ParseAddress(email)
				return err
			}),
		field.Bool("active").
			Comment("indicates if the subscriber is active or not").
			Default(true),
		field.String("ip_address").
			Comment("IP address of the subscriber").
			Optional(),
		field.String("token").
			Comment("the token used to unsubscribe").
			Unique().
			NotEmpty(),
		field.Bytes("secret").
			Comment("the comparison secret to verify the token's signature").
			NotEmpty().
			Nillable().Annotations(entgql.Skip()),
	}
}

// Mixin of the Subscribers
func (Subscribers) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.IDMixin{},
		mixin.SoftDeleteMixin{},
		OrgOwnerMixin{
			Ref:        "subscribers",
			Optional:   true,
			AllowWhere: true,
		},
	}
}

// Edges of the Subscribers
func (Subscribers) Edges() []ent.Edge {
	return []ent.Edge{
		// Edges go here
	}
}

// Indexes of the Subscribers
func (Subscribers) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("email", "active", "owner_id").
			Unique().Annotations(entsql.IndexWhere("deleted_at is NULL")),
	}
}

// Annotations of the Subscribers
func (Subscribers) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
	}
}
