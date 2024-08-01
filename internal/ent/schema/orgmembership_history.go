// Code generated by enthistory, DO NOT EDIT.
package schema

import (
	"context"
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/interceptors"
	"github.com/datumforge/enthistory"
	"github.com/datumforge/entx"
	"github.com/datumforge/fgax/entfga"
)

// OrgMembershipHistory holds the schema definition for the OrgMembershipHistory entity.
type OrgMembershipHistory struct {
	ent.Schema
}

// Annotations of the OrgMembershipHistory.
func (OrgMembershipHistory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entx.SchemaGenSkip(true),
		entsql.Annotation{
			Table: "org_membership_history",
		},
		enthistory.Annotations{
			IsHistory: true,
			Exclude:   true,
		},
		entgql.QueryField(),
		entgql.RelayConnection(),
		entfga.Annotations{
			ObjectType:   "organization",
			IDField:      "OrganizationID",
			IncludeHooks: false,
		},
	}
}

// Fields of the OrgMembershipHistory.
func (OrgMembershipHistory) Fields() []ent.Field {
	historyFields := []ent.Field{
		field.Time("history_time").
			Default(time.Now).
			Immutable(),
		field.String("ref").
			Immutable().
			Optional(),
		field.Enum("operation").
			GoType(enthistory.OpType("")).
			Immutable(),
	}

	// get the fields from the mixins
	// we only want to include mixin fields, not edges
	// so this prevents FKs back to the main tables
	mixins := OrgMembership{}.Mixin()
	for _, mixin := range mixins {
		for _, field := range mixin.Fields() {
			// make sure the mixed in fields do not have unique constraints
			field.Descriptor().Unique = false

			// make sure the mixed in fields do not have validators
			field.Descriptor().Validators = nil

			// append the mixed in field to the history fields
			historyFields = append(historyFields, field)
		}
	}

	original := OrgMembership{}
	for _, field := range original.Fields() {
		// make sure the fields do not have unique constraints
		field.Descriptor().Unique = false

		// make sure the mixed in fields do not have validators
		field.Descriptor().Validators = nil

		// append the field to the history fields
		historyFields = append(historyFields, field)
	}

	return historyFields
}

// Indexes of the OrgMembershipHistory
func (OrgMembershipHistory) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("history_time"),
	}
}

// Interceptors of the OrgMembershipHistory
func (OrgMembershipHistory) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		interceptors.HistoryAccess("audit_log_viewer", false, false),
	}
}

// Policy of the OrgMembershipHistory
func (OrgMembershipHistory) Policy() ent.Policy {
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			privacy.OrgMembershipHistoryQueryRuleFunc(func(ctx context.Context, q *generated.OrgMembershipHistoryQuery) error {
				return q.CheckAccess(ctx)
			}),
			privacy.AlwaysDenyRule(),
		},
	}
}
