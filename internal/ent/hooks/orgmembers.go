package hooks

import (
	"context"
	"fmt"
	"strings"

	"entgo.io/ent"
	openfga "github.com/openfga/go-sdk"
	ofgaclient "github.com/openfga/go-sdk/client"

	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/hook"
	"github.com/datumforge/datum/internal/ent/mixin"
	"github.com/datumforge/datum/internal/fga"
)

// HookOrgMembersAuthz runs on organization member mutations to setup or remove relationship tuples
func HookOrgMembersAuthz() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return hook.OrgMembershipFunc(func(ctx context.Context, m *generated.OrgMembershipMutation) (ent.Value, error) {
			var (
				ids []string
				err error
			)

			// get the IDs that will be updated based on the mutation predicate
			// this must happen before the mutation, and is not valid on OpCreate
			if !m.Op().Is(ent.OpCreate) {
				ids, err = m.IDs(ctx)
				if err != nil {
					return nil, err
				}
			}

			// do the mutation, and then create/delete the relationships in FGA
			retValue, err := next.Mutate(ctx, m)
			if err != nil {
				// if we error, do not attempt to create the relationships
				return retValue, err
			}

			switch op := m.Op(); op {
			case ent.OpCreate:
				// create the relationship tuple for the owner
				err = orgMemberCreateHook(ctx, m)
			case ent.OpDelete, ent.OpDeleteOne:
				// delete all relationship tuples on delete, or soft delete (Update Op)
				err = orgMemberDeleteHook(ctx, m, ids)
			case ent.OpUpdate, ent.OpUpdateOne:
				if mixin.CheckIsSoftDelete(ctx) {
					err = orgMemberDeleteHook(ctx, m, ids)
				} else {
					err = orgMemberUpdateHook(ctx, m, ids)
				}
			default:
				return nil, fmt.Errorf("unsupported operations")
			}

			return retValue, err
		})
	}
}

func orgMemberCreateHook(ctx context.Context, m *generated.OrgMembershipMutation) error {
	// Add relationship tuples if authz is enabled
	if m.Authz.Ofga != nil {
		tuples, err := getOrgMemberTuples(ctx, m)
		if err != nil {
			return err
		}

		if len(tuples) > 0 {
			if _, err := m.Authz.CreateRelationshipTuple(ctx, tuples); err != nil {
				m.Logger.Errorw("failed to create relationship tuple", "error", err)

				return err
			}

			m.Logger.Infow("created relationship tuples", "relation", fga.OwnerRelation, "object", tuples[0].Object)
		}
	}

	return nil
}

func orgMemberDeleteHook(ctx context.Context, m *generated.OrgMembershipMutation, ids []string) error {
	if m.Authz.Ofga != nil {
		tuples, err := getDeleteOrgMemberTuples(ctx, m, ids)
		if err != nil {
			return err
		}

		if len(tuples) > 0 {
			if _, err := m.Authz.DeleteRelationshipTuple(ctx, tuples); err != nil {
				m.Logger.Errorw("failed to delete relationship tuple", "error", err)

				return err
			}

			m.Logger.Infow("deleted relationship tuples", "relation", fga.OwnerRelation, "object", tuples[0].Object)
		}
	}

	return nil
}

func orgMemberUpdateHook(ctx context.Context, m *generated.OrgMembershipMutation, ids []string) error {
	if m.Authz.Ofga != nil {
		tuples, err := getUpdateOrgMemberTuples(ctx, m, ids)
		if err != nil {
			return err
		}

		if _, err := m.Authz.WriteTuples(ctx, tuples); err != nil {
			m.Logger.Errorw("failed to update relationship tuple", "error", err)

			return err
		}
	}

	return nil
}

func getOrgMemberTuples(ctx context.Context, m *generated.OrgMembershipMutation) (tuples []ofgaclient.ClientTupleKey, err error) {
	userID, _ := m.UserID()
	orgID, _ := m.OrgID()
	role, _ := m.Role()

	fgaRelation, err := roleToRelation(role)
	if err != nil {
		return nil, err
	}

	object := fmt.Sprintf("%s:%s", "organization", orgID)

	m.Logger.Infow("creating relationship tuples", "relation", fgaRelation, "object", object)

	tuples = []ofgaclient.ClientTupleKey{{
		User:     fmt.Sprintf("user:%s", userID),
		Relation: fgaRelation,
		Object:   object,
	}}

	return
}

// getDeleteOrgMemberTuples gets all tuples related to the orgMembership IDs that were deleted
func getDeleteOrgMemberTuples(ctx context.Context, m *generated.OrgMembershipMutation, ids []string) (tuples []openfga.TupleKeyWithoutCondition, err error) {
	tuples = []openfga.TupleKeyWithoutCondition{}

	// User the IDs of the org memberships and delete all related tuples
	for _, id := range ids {
		// this happens after soft-delete, allow the request to pull the record
		ctx := mixin.SkipSoftDelete(ctx)

		om, err := m.Client().OrgMembership.Get(ctx, id)
		if err != nil {
			return nil, err
		}

		fgaRelation, err := roleToRelation(om.Role)
		if err != nil {
			return nil, err
		}

		object := fmt.Sprintf("%s:%s", "organization", om.OrgID)

		m.Logger.Infow("deleting relationship tuples", "relation", fgaRelation, "object", object)

		tuples = append(tuples, openfga.TupleKeyWithoutCondition{
			User:     fmt.Sprintf("user:%s", om.UserID),
			Relation: fgaRelation,
			Object:   object,
		})
	}

	return tuples, nil
}

// getUpdateOrgMemberTuples gets all tuples related to the orgMembership IDs that were updated
func getUpdateOrgMemberTuples(ctx context.Context, m *generated.OrgMembershipMutation, ids []string) (tuples ofgaclient.ClientWriteRequest, err error) {
	tuples = ofgaclient.ClientWriteRequest{
		Writes:  []ofgaclient.ClientTupleKey{},
		Deletes: []openfga.TupleKeyWithoutCondition{},
	}

	oldRole, err := m.OldRole(ctx)
	if err != nil {
		return tuples, err
	}

	newRole, exists := m.Role()
	if !exists {
		return tuples, fmt.Errorf("no role provided")
	}

	// User the IDs of the org memberships and delete all related tuples
	for _, id := range ids {
		// this happens after soft-delete, allow the request to pull the record
		om, err := m.Client().OrgMembership.Get(ctx, id)
		if err != nil {
			return tuples, err
		}

		oldRelation, err := roleToRelation(oldRole)
		if err != nil {
			return tuples, err
		}

		newRelation, err := roleToRelation(newRole)
		if err != nil {
			return tuples, err
		}

		object := fmt.Sprintf("%s:%s", "organization", om.OrgID)

		m.Logger.Infow("deleting relationship tuples", "relation", oldRelation, "object", object)

		tuples.Deletes = append(tuples.Deletes, openfga.TupleKeyWithoutCondition{
			User:     fmt.Sprintf("user:%s", om.UserID),
			Relation: oldRelation,
			Object:   object,
		})

		tuples.Writes = append(tuples.Writes, ofgaclient.ClientTupleKey{
			User:     fmt.Sprintf("user:%s", om.UserID),
			Relation: newRelation,
			Object:   object,
		})
	}

	return tuples, nil
}

func roleToRelation(r enums.Role) (string, error) {
	switch r {
	case enums.RoleOwner, enums.RoleAdmin, enums.RoleMember:
		return strings.ToLower(r.String()), nil
	default:
		return "", ErrUnsupportedFGARole
	}
}
