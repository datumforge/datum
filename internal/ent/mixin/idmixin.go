package mixin

import (
	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"

	"github.com/datumforge/datum/pkg/utils/ulids"
)

// IDMixin holds the schema definition for the ID
type IDMixin struct {
	mixin.Schema
}

// Fields of the IDMixin.
func (IDMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Immutable().
			Annotations(entoas.Annotation{ReadOnly: true}).
			DefaultFunc(func() string { return ulids.New().String() }),
	}
}
