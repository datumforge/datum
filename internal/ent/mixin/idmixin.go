package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"

	"github.com/jaevor/go-nanoid"
)

const (
	idLength = 21
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
			DefaultFunc(mustGetNewID),
	}
}

// getNewID returns an ID based on go-nanoid
func getNewID() (string, error) {
	canonicID, err := nanoid.Standard(idLength)
	if err != nil {
		return "", err
	}

	return canonicID(), nil
}

// mustGetNewID returns an ID
func mustGetNewID() string {
	v, err := getNewID()
	if err != nil {
		panic(err)
	}

	return v
}
