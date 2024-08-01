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

// OauthProviderHistory holds the schema definition for the OauthProviderHistory entity.
type OauthProviderHistory struct {
	ent.Schema
}

// Annotations of the OauthProviderHistory.
func (OauthProviderHistory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entx.SchemaGenSkip(true),
		entsql.Annotation{
			Table: "oauth_provider_history",
		},
		enthistory.Annotations{
			IsHistory: true,
			Exclude:   true,
		},
		entgql.QueryField(),
		entgql.RelayConnection(),
		entfga.Annotations{
			ObjectType:   "organization",
			IDField:      "OwnerID",
			IncludeHooks: false,
		},
	}
}

// Fields of the OauthProviderHistory.
func (OauthProviderHistory) Fields() []ent.Field {
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
	mixins := OauthProvider{}.Mixin()
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

	original := OauthProvider{}
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

// Indexes of the OauthProviderHistory
func (OauthProviderHistory) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("history_time"),
	}
}

// Interceptors of the OauthProviderHistory
func (OauthProviderHistory) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		interceptors.HistoryAccess("audit_log_viewer", true, false),
	}
}

// Policy of the OauthProviderHistory
func (OauthProviderHistory) Policy() ent.Policy {
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			privacy.OauthProviderHistoryQueryRuleFunc(func(ctx context.Context, q *generated.OauthProviderHistoryQuery) error {
				return q.CheckAccess(ctx)
			}),
			privacy.AlwaysDenyRule(),
		},
	}
}
