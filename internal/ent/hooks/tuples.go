package hooks

import (
	"strings"

	"github.com/datumforge/fgax"

	"github.com/datumforge/datum/internal/ent/enums"
)

// getTupleKey creates a Tuple key with the provided subject, object, and role
func getTupleKey(subjectID, subjectType, objectID, objectType string, role enums.Role) (fgax.TupleKey, error) {
	fgaRelation, err := roleToRelation(role)
	if err != nil {
		return fgax.NewTupleKey(), err
	}

	sub := fgax.Entity{
		Kind:       fgax.Kind(subjectType),
		Identifier: subjectID,
	}

	object := fgax.Entity{
		Kind:       fgax.Kind(objectType),
		Identifier: objectID,
	}

	return fgax.TupleKey{
		Subject:  sub,
		Object:   object,
		Relation: fgax.Relation(fgaRelation),
	}, nil
}

// getTupleKey creates a user Tuple key with the provided user ID, object, and role
func getUserTupleKey(userID, objectID, objectType string, role enums.Role) (fgax.TupleKey, error) {
	return getTupleKey(userID, "user", objectID, objectType, role)
}

func roleToRelation(r enums.Role) (string, error) {
	switch r {
	case enums.RoleOwner, enums.RoleAdmin, enums.RoleMember:
		return strings.ToLower(r.String()), nil
	case fgax.ParentRelation:
		return r.String(), nil
	default:
		return "", ErrUnsupportedFGARole
	}
}
