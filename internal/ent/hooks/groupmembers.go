package hooks

import (
	"context"
	"fmt"

	"entgo.io/ent"
	"github.com/datumforge/fgax"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/hook"
	"github.com/datumforge/datum/internal/ent/mixin"
)

// HookGroupMembersAuthz runs on group member mutations to setup or remove relationship tuples
func HookGroupMembersAuthz() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return hook.GroupMembershipFunc(func(ctx context.Context, m *generated.GroupMembershipMutation) (ent.Value, error) {
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
				err = groupMemberCreateHook(ctx, m)
			case ent.OpDelete, ent.OpDeleteOne:
				// delete all relationship tuples on delete, or soft delete (Update Op)
				err = groupMemberDeleteHook(ctx, m, ids)
			case ent.OpUpdate, ent.OpUpdateOne:
				if mixin.CheckIsSoftDelete(ctx) {
					err = groupMemberDeleteHook(ctx, m, ids)
				} else {
					err = groupMemberUpdateHook(ctx, m, ids)
				}
			default:
				// we should never get here
				return nil, fmt.Errorf("unsupported operations") //nolint:goerr113
			}

			return retValue, err
		})
	}
}

func groupMemberCreateHook(ctx context.Context, m *generated.GroupMembershipMutation) error {
	// Add relationship tuples if authz is enabled
	tuple, err := getGroupMemberTuple(m)
	if err != nil {
		return err
	}

	if _, err := m.Authz.WriteTupleKeys(ctx, []fgax.TupleKey{tuple}, nil); err != nil {
		m.Logger.Errorw("failed to create relationship tuple", "error", err)

		return err
	}

	m.Logger.Debugw("created relationship tuples", "relation", fgax.OwnerRelation, "object", tuple.Object)

	return nil
}

func groupMemberDeleteHook(ctx context.Context, m *generated.GroupMembershipMutation, ids []string) error {
	tuples, err := getDeleteGroupMemberTuples(ctx, m, ids)
	if err != nil {
		return err
	}

	if len(tuples) > 0 {
		if _, err := m.Authz.WriteTupleKeys(ctx, nil, tuples); err != nil {
			m.Logger.Errorw("failed to delete relationship tuple", "error", err)

			return err
		}

		m.Logger.Debugw("deleted group relationship tuples")
	}

	return nil
}

func groupMemberUpdateHook(ctx context.Context, m *generated.GroupMembershipMutation, ids []string) error {
	writes, deletes, err := getUpdateGroupMemberTuples(ctx, m, ids)
	if err != nil {
		return err
	}

	if len(writes) == 0 && len(deletes) == 0 {
		m.Logger.Debugw("no relationships to create or delete")

		return nil
	}

	if _, err := m.Authz.WriteTupleKeys(ctx, writes, deletes); err != nil {
		m.Logger.Errorw("failed to update relationship tuple", "error", err)

		return err
	}

	return nil
}

func getGroupMemberTuple(m *generated.GroupMembershipMutation) (tuple fgax.TupleKey, err error) {
	userID, _ := m.UserID()
	groupID, _ := m.GroupID()
	role, _ := m.Role()

	return getUserTupleKey(userID, groupID, "group", role)
}

// getDeleteGroupMemberTuples gets all tuples related to the groupMembership IDs that were deleted
func getDeleteGroupMemberTuples(ctx context.Context, m *generated.GroupMembershipMutation, ids []string) (tuples []fgax.TupleKey, err error) {
	// User the IDs of the group memberships and delete all related tuples
	for _, id := range ids {
		// this happens after soft-delete, allow the request to pull the record
		ctx := mixin.SkipSoftDelete(ctx)

		om, err := m.Client().GroupMembership.Get(ctx, id)
		if err != nil {
			return nil, err
		}

		t, err := getUserTupleKey(om.UserID, om.GroupID, "group", om.Role)
		if err != nil {
			return nil, err
		}

		tuples = append(tuples, t)
	}

	return
}

// getUpdateGroupMemberTuples gets all tuples related to the groupMembership IDs that were updated
func getUpdateGroupMemberTuples(ctx context.Context, m *generated.GroupMembershipMutation, ids []string) (writes []fgax.TupleKey, deletes []fgax.TupleKey, err error) {
	oldRole, err := m.OldRole(ctx)
	if err != nil {
		return writes, deletes, err
	}

	newRole, exists := m.Role()
	if !exists {
		return writes, deletes, ErrMissingRole
	}

	if oldRole == newRole {
		m.Logger.Debugw("nothing to update, roles are the same", "old_role", oldRole, "new_role", newRole)

		return writes, deletes, nil
	}

	// User the IDs of the group memberships and delete all related tuples
	for _, id := range ids {
		// this happens after soft-delete, allow the request to pull the record
		om, err := m.Client().GroupMembership.Get(ctx, id)
		if err != nil {
			return writes, deletes, err
		}

		d, err := getUserTupleKey(om.UserID, om.GroupID, "group", oldRole)
		if err != nil {
			return writes, deletes, err
		}

		deletes = append(deletes, d)

		w, err := getUserTupleKey(om.UserID, om.GroupID, "group", newRole)
		if err != nil {
			return writes, deletes, err
		}

		writes = append(writes, w)
	}

	return writes, deletes, nil
}
