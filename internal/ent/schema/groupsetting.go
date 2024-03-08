package schema

import (
	"context"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/mixin"
	"github.com/datumforge/datum/internal/ent/privacy/rule"
	"github.com/datumforge/fgax/entfga"
)

// GroupSetting holds the schema definition for the GroupSetting entity
type GroupSetting struct {
	ent.Schema
}

// Fields of the GroupSetting.
func (GroupSetting) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("visibility").
			Comment("whether the group is visible to it's members / owners only or if it's searchable by anyone within the organization").
			GoType(enums.Visibility("")).
			Default(string(enums.Public)),
		field.Enum("join_policy").
			Comment("the policy governing ability to freely join a group, whether it requires an invitation, application, or either").
			GoType(enums.JoinPolicy("")).
			Default(string(enums.InviteOrApplication)),
		field.JSON("tags", []string{}).
			Comment("tags associated with the object").
			Optional().
			Default([]string{}),
		field.Bool("sync_to_slack").
			Comment("whether to sync group members to slack groups").
			Optional().
			Default(false),
		field.Bool("sync_to_github").
			Comment("whether to sync group members to github groups").
			Optional().
			Default(false),
		field.String("group_id").
			Comment("the group id associated with the settings").
			Optional(),
	}
}

// Edges of the GroupSetting
func (GroupSetting) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("group", Group.Type).Ref("setting").Field("group_id").Unique(),
	}
}

// Annotations of the GroupSetting
func (GroupSetting) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
		entfga.Annotations{
			ObjectType:      "group",
			IncludeHooks:    false,
			IDField:         "GroupID",
			NillableIDField: true,
		},
	}
}

// Mixin of the GroupSetting
func (GroupSetting) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.IDMixin{},
		mixin.SoftDeleteMixin{},
	}
}

// Policy defines the privacy policy of the GroupSetting
func (GroupSetting) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			rule.DenyIfNoSubject(),
			privacy.GroupSettingMutationRuleFunc(func(ctx context.Context, m *generated.GroupSettingMutation) error {
				return m.CheckAccessForEdit(ctx)
			}),
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			privacy.GroupSettingQueryRuleFunc(func(ctx context.Context, q *generated.GroupSettingQuery) error {
				return q.CheckAccess(ctx)
			}),
			privacy.AlwaysDenyRule(),
		},
	}
}
