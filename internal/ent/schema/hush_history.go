// Code generated by enthistory, DO NOT EDIT.
package schema

import (
	"time"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/datumforge/enthistory"
	"github.com/datumforge/entx"
)

// HushHistory holds the schema definition for the HushHistory entity.
type HushHistory struct {
	ent.Schema
}

// Annotations of the HushHistory.
func (HushHistory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entx.SchemaGenSkip(true),
		entsql.Annotation{
			Table: "hush_history",
		},
		enthistory.Annotations{
			IsHistory: true,
			Exclude:   true,
		},
		entgql.QueryField(),
		entgql.RelayConnection(),
	}
}

// Fields of the HushHistory.
func (HushHistory) Fields() []ent.Field {
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
	mixins := Hush{}.Mixin()
	for _, mixin := range mixins {
		for _, field := range mixin.Fields() {
			historyFields = append(historyFields, field)
		}
	}

	original := Hush{}
	for _, field := range original.Fields() {
		historyFields = append(historyFields, field)
	}

	return historyFields
}
func (HushHistory) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("history_time"),
	}
}

