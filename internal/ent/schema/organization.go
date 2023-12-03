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
	"entgo.io/ent/schema/index"

	gen "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/hook"
	"github.com/datumforge/datum/internal/ent/generated/intercept"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/generated/user"
	"github.com/datumforge/datum/internal/ent/hooks"
	"github.com/datumforge/datum/internal/ent/interceptors"
	"github.com/datumforge/datum/internal/ent/mixin"
	"github.com/datumforge/datum/internal/ent/privacy/rule"
	"github.com/ogen-go/ogen"
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
		field.String("name").
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
			Comment("An optional description of the organization").
			Optional().
			Annotations(
				entgql.Skip(entgql.SkipWhereInput),
			),
		field.String("parent_organization_id").Optional().Immutable().
			Comment("The ID of the parent organization for the organization.").
			Default("0").
			Annotations(
				entgql.Type("ID"),
				entgql.Skip(entgql.SkipMutationUpdateInput, entgql.SkipType),
				entoas.Schema(ogen.String()),
			),
		field.Text("path").Optional().
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.Enum("kind").NamedValues(
			"root", "root",
			"organization", "org",
		).
			Default("org").Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.String("owner_id").Optional().Nillable(),
		field.String("code").MaxLen(45).Optional().
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
	}
}

// Edges of the Organization
func (Organization) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", Organization.Type).
			Annotations(
				entgql.RelayConnection(),
				entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
				entsql.Annotation{OnDelete: entsql.Cascade},
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
		edge.To("owner", User.Type).Field("owner_id").Unique(),
	}
}

func (Organization) Indexes() []ent.Index {
	return []ent.Index{
		// names should be unique, but ignore deleted names
		index.Fields("name").
			Unique().Annotations(entsql.IndexWhere("deleted_at is NULL")),
	}
}

// Annotations of the Organization
func (Organization) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
		entgql.QueryField(),
	}
}

// Mixin of the Organization
func (Organization) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.IDMixin{},
		mixin.NewSoftDeleteMixin[intercept.Query, *gen.Client](intercept.NewQuery),
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
			rule.HasOrgReadAccess(),  // Requires a user to have can_view access of the org
			privacy.AlwaysDenyRule(), // Deny all other users
		},
	}
}

// Hooks of the Organization
func (Organization) Hooks() []ent.Hook {
	return []ent.Hook{
		pathHook(),
		hook.On(checkDeleteHook(), ent.OpDeleteOne),
		hook.On(ownerCheckHook(), ent.OpCreate|ent.OpUpdateOne|ent.OpUpdate),
		hooks.HookOrganization(),
	}
}

// Interceptors of the Organization
func (Organization) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		interceptors.InterceptorOrganization(),
	}
}

func pathHook() ent.Hook {
	return hook.On(
		func(next ent.Mutator) ent.Mutator {
			return hook.OrganizationFunc(func(ctx context.Context, mutation *gen.OrganizationMutation) (gen.Value, error) {
				if _, ok := mutation.Path(); ok {
					return next.Mutate(ctx, mutation)
				}
				if pid, ok := mutation.ParentOrganizationID(); ok {
					id, _ := mutation.ID()

					if pid == "0" {
						mutation.SetPath(id)
					} else {
						parentPath := ""
						prow, err := mutation.Client().Organization.Query().Where(organization.ID(pid)).
							Select(organization.FieldPath).Only(ctx)
						if err != nil {
							if !gen.IsNotFound(err) {
								return nil, err
							}
							parentPath = ""
						} else {
							parentPath = prow.Path + "/"
						}
						mutation.SetPath(parentPath + id)
					}
					if c, _ := mutation.Code(); c == "" {
						mutation.SetCode(id)
					}
				}
				return next.Mutate(ctx, mutation)
			})
		}, ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne)
}

func checkDeleteHook() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return hook.OrganizationFunc(func(ctx context.Context, mutation *gen.OrganizationMutation) (gen.Value, error) {
			if mutation.Op() == ent.OpDeleteOne {
				if id, ok := mutation.ID(); ok {
					count, err := mutation.Client().Organization.Query().Where(
						organization.ParentOrganizationID(id),
					).Count(ctx)
					if err != nil {
						return nil, err
					}
					if count > 0 {
						return nil, fmt.Errorf("organization has children")
					}
				}
			}
			return next.Mutate(ctx, mutation)
		})
	}
}

func ownerCheckHook() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return hook.OrganizationFunc(func(ctx context.Context, m *gen.OrganizationMutation) (gen.Value, error) {
			if uid, ok := m.OwnerID(); ok {
				usr, err := m.Client().User.Get(ctx, uid)
				if err != nil {
					return nil, err
				}
				if usr.UserType != user.UserTypeAccount {
					return nil, fmt.Errorf("owner must be account: %s", usr.DisplayName)
				}
				m.SetKind(organization.KindRoot)
			}
			return next.Mutate(ctx, m)
		})
	}
}
