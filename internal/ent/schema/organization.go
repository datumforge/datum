package schema

import (
	"context"
	"fmt"
	"strings"

	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/ogen-go/ogen"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/hook"
	"github.com/datumforge/datum/internal/ent/mixin"
)

const (
	orgNameMaxLen = 160

	objectType     = "organization"
	memberRelation = "member"
	ownerRelation  = "owner"
)

// Organization holds the schema definition for the Organization entity - organizations are the top level tenancy construct in the system
type Organization struct {
	ent.Schema
}

// Fields of the Organization
func (Organization) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().
			MaxLen(orgNameMaxLen).
			NotEmpty().
			Annotations(
				entgql.OrderField("name"),
				entgql.Skip(entgql.SkipWhereInput),
			),
		field.String("display_name").
			Comment("The organization's displayed 'friendly' name").
			MaxLen(nameMaxLen).
			NotEmpty().
			Default("unknown").
			Annotations(
				entgql.OrderField("display_name"),
			).
			Validate(
				func(s string) error {
					if strings.Contains(s, " ") {
						return ErrContainsSpaces
					}
					return nil
				},
			),
		field.String("description").
			Comment("An optional description of the Organization").
			Optional().
			Annotations(
				entgql.Skip(entgql.SkipWhereInput),
			),
		field.String("parent_organization_id").Optional().Immutable().
			Comment("The ID of the parent organization for the organization.").
			Annotations(
				entgql.Type("ID"),
				entgql.Skip(entgql.SkipMutationUpdateInput, entgql.SkipType),
				entoas.Schema(ogen.String()),
			),
	}
}

// Edges of the Organization
func (Organization) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", Organization.Type).
			Annotations(
				entgql.RelayConnection(),
				entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
			).
			From("parent").
			Field("parent_organization_id").
			Immutable().
			Unique(),
		// an org can have and belong to many users
		edge.From("users", User.Type).Ref("organizations"),
		edge.To("groups", Group.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("integrations", Integration.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("setting", OrganizationSettings.Type).Unique().Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("entitlements", Entitlement.Type),
		edge.To("oauthprovider", OauthProvider.Type),
	}
}

// Annotations of the Organization
func (Organization) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
	}
}

// Mixin of the Organization
func (Organization) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.IDMixin{},
	}
}

// Hooks of the Organization
func (Organization) Hooks() []ent.Hook {
	return []ent.Hook{
		HookOrganization(),
	}
}

// HookOrganization runs on organization mutations to setup or remove relationship tuples
func HookOrganization() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return hook.OrganizationFunc(func(ctx context.Context, m *generated.OrganizationMutation) (ent.Value, error) {
			// do the mutation, and then create/delete the relationship
			retValue, err := next.Mutate(ctx, m)
			if err != nil {
				// if we error, do not attempt to create the relationships
				return retValue, err
			}

			if m.Op().Is(ent.OpCreate) {
				// create the relationship tuple for the owner
				err = organizationCreateHook(ctx, m)
			} else if m.Op().Is(ent.OpDelete | ent.OpDeleteOne) {
				// delete all relationship tuples
				err = organizationDeleteHook(ctx, m)
			}

			return retValue, err
		})
	}
}

func organizationCreateHook(ctx context.Context, m *generated.OrganizationMutation) error {
	// Add relationship tuples if authz is enabled
	if m.Authz.Ofga != nil {
		objID, exists := m.ID()
		m.Logger.Infow("creating relationship tuples", "object_id", objID)

		if exists {
			if err := m.Authz.CreateRelationshipTupleWithUser(ctx, ownerRelation, fmt.Sprintf("%s:%s", objectType, objID)); err != nil {
				m.Logger.Errorw("failed to create relationship tuple", "error", err)

				// TODO: rollback mutation if tuple creation fails
				return ErrInternalServerError
			}
		}

		m.Logger.Infow("created relationship tuples", "object_id", objID)
	}

	return nil
}

func organizationDeleteHook(ctx context.Context, m *generated.OrganizationMutation) error {
	// Add relationship tuples if authz is enabled
	if m.Authz.Ofga != nil {
		objID, _ := m.ID()
		m.Logger.Infow("going to relationship tuples", "object_id", objID)

		// Add relationship tuples if authz is enabled
		if m.Authz.Ofga != nil {
			if err := m.Authz.DeleteAllObjectRelations(ctx, fmt.Sprintf("%s:%s", objectType, objID)); err != nil {
				m.Logger.Errorw("failed to delete relationship tuples", "error", err)

				return ErrInternalServerError
			}

			m.Logger.Infow("deleted relationship tuples", "object_id", objID)
		}
	}

	return nil
}
