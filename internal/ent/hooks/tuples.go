package hooks

import (
	"strings"

	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/fga"
)

// getTupleKey creates a Tuple key with the provided subject, object, and role
func getTupleKey(subjectID, subjectType, objectID, objectType string, role enums.Role) (fga.TupleKey, error) {
	fgaRelation, err := roleToRelation(role)
	if err != nil {
		return fga.NewTupleKey(), err
	}

	sub := fga.Entity{
		Kind:       fga.Kind(subjectID),
		Identifier: subjectID,
	}

	object := fga.Entity{
		Kind:       fga.Kind(objectType),
		Identifier: objectID,
	}

	return fga.TupleKey{
		Subject:  sub,
		Object:   object,
		Relation: fga.Relation(fgaRelation),
	}, nil
}

// getTupleKey creates a user Tuple key with the provided user ID, object, and role
func getUserTupleKey(userID, objectID, objectType string, role enums.Role) (fga.TupleKey, error) {
	return getTupleKey(userID, "user", objectID, objectType, role)
}

func roleToRelation(r enums.Role) (string, error) {
	switch r {
	case enums.RoleOwner, enums.RoleAdmin, enums.RoleMember:
		return strings.ToLower(r.String()), nil
	case fga.ParentRelation:
		return r.String(), nil
	default:
		return "", ErrUnsupportedFGARole
	}
}
