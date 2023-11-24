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
	ofgaclient "github.com/openfga/go-sdk/client"

	"github.com/datumforge/datum/internal/echox"
	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/hook"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/mixin"
	"github.com/datumforge/datum/internal/ent/privacy/rule"
	"github.com/datumforge/datum/internal/fga"
)

const (
	orgNameMaxLen = 160
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
		edge.To("setting", OrganizationSetting.Type).Unique().Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
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

// Policy defines the privacy policy of the Organization.
func (Organization) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			rule.DenyIfNoSubject(),      // Requires a user to be authenticated with a valid JWT
			rule.HasOrgMutationAccess(), // Requires edit for Update, and delete for Delete mutations
			privacy.AlwaysAllowRule(),   // Allow all other users (e.g. a user with a JWT should be able to create a new org)
		},
		Query: privacy.QueryPolicy{
			rule.DenyIfNoSubject(),   // Requires a user to be authenticated with a valid JWT
			rule.HasOrgReadAccess(),  // Requires a user to have can_view access of the org
			privacy.AlwaysDenyRule(), // Deny all other users
		},
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
		objType := strings.ToLower(m.Type())
		object := fmt.Sprintf("%s:%s", objType, objID)

		m.Logger.Infow("creating relationship tuples", "relation", fga.OwnerRelation, "object", object)

		if exists {
			tuples, err := createTuple(ctx, &m.Authz, fga.OwnerRelation, object)
			if err != nil {
				return err
			}

			if _, err := m.Authz.CreateRelationshipTuple(ctx, tuples); err != nil {
				m.Logger.Errorw("failed to create relationship tuple", "error", err)

				// TODO: rollback mutation if tuple creation fails
				return ErrInternalServerError
			}
		}

		m.Logger.Infow("created relationship tuples", "relation", fga.OwnerRelation, "object", object)
	}

	return nil
}

func organizationDeleteHook(ctx context.Context, m *generated.OrganizationMutation) error {
	// Add relationship tuples if authz is enabled
	if m.Authz.Ofga != nil {
		objID, _ := m.ID()
		objType := strings.ToLower(m.Type())
		object := fmt.Sprintf("%s:%s", objType, objID)

		m.Logger.Infow("deleting relationship tuples", "object", object)

		// Add relationship tuples if authz is enabled
		if m.Authz.Ofga != nil {
			if err := m.Authz.DeleteAllObjectRelations(ctx, object); err != nil {
				m.Logger.Errorw("failed to delete relationship tuples", "error", err)

				return ErrInternalServerError
			}

			m.Logger.Infow("deleted relationship tuples", "object", object)
		}
	}

	return nil
}

func createTuple(ctx context.Context, c *fga.Client, relation, object string) ([]ofgaclient.ClientTupleKey, error) {
	ec, err := echox.EchoContextFromContext(ctx)
	if err != nil {
		c.Logger.Errorw("unable to get echo context", "error", err)

		return nil, err
	}

	actor, err := echox.GetActorSubject(*ec)
	if err != nil {
		return nil, err
	}

	// TODO: convert jwt sub --> uuid

	tuples := []ofgaclient.ClientTupleKey{{
		User:     fmt.Sprintf("user:%s", actor),
		Relation: relation,
		Object:   object,
	}}

	return tuples, nil
}
