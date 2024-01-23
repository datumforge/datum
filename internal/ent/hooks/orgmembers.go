package hooks

import (
	"context"
	"fmt"

	"entgo.io/ent"

	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/hook"
	"github.com/datumforge/datum/internal/ent/mixin"
	"github.com/datumforge/datum/internal/fga"
)

func HookOrgMembers() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.OrgMembershipFunc(func(ctx context.Context, mutation *generated.OrgMembershipMutation) (generated.Value, error) {
			// check role, if its not set the default is member
			role, _ := mutation.Role()
			if role == enums.RoleOwner {
				return next.Mutate(ctx, mutation)
			}

			// get the organization based on input
			orgID, exists := mutation.OrgID()
			if exists {
				org, err := mutation.Client().Organization.Get(ctx, orgID)
				if err != nil {
					mutation.Logger.Errorw("error getting organization", "error", err)

					return nil, err
				}

				// do not allow members to be added to personal orgs
				if org.PersonalOrg {
					return nil, ErrPersonalOrgsNoMembers
				}
			}

			return next.Mutate(ctx, mutation)
		})
	}, ent.OpCreate)
}

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
				// we should never get here
				return nil, fmt.Errorf("unsupported operations") //nolint:goerr113
			}

			return retValue, err
		})
	}
}

func orgMemberCreateHook(ctx context.Context, m *generated.OrgMembershipMutation) error {
	// Add relationship tuples if authz is enabled
	tuple, err := getOrgMemberTuple(m)
	if err != nil {
		return err
	}

	m.Logger.Debugw("details for fga", "object", tuple.Object, "relation", tuple.Relation, "subject", tuple.Subject)

	if _, err := m.Authz.WriteTupleKeys(ctx, []fga.TupleKey{tuple}, nil); err != nil {
		m.Logger.Errorw("failed to create relationship tuple", "error", err)

		return err
	}

	m.Logger.Debugw("created organization relationship tuples")

	return nil
}

func orgMemberDeleteHook(ctx context.Context, m *generated.OrgMembershipMutation, ids []string) error {
	tuples, err := getDeleteOrgMemberTuples(ctx, m, ids)
	if err != nil {
		return err
	}

	if len(tuples) > 0 {
		if _, err := m.Authz.WriteTupleKeys(ctx, nil, tuples); err != nil {
			m.Logger.Errorw("failed to delete relationship tuple", "error", err)

			return err
		}

		m.Logger.Debugw("deleted relationship tuples", "relation", fga.OwnerRelation, "object", tuples[0].Object)
	}

	return nil
}

func orgMemberUpdateHook(ctx context.Context, m *generated.OrgMembershipMutation, ids []string) error {
	writes, deletes, err := getUpdateOrgMemberTuples(ctx, m, ids)
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

func getOrgMemberTuple(m *generated.OrgMembershipMutation) (tuple fga.TupleKey, err error) {
	userID, _ := m.UserID()
	orgID, _ := m.OrgID()
	role, _ := m.Role()

	return getUserTupleKey(userID, orgID, "organization", role)
}

// getDeleteOrgMemberTuples gets all tuples related to the orgMembership IDs that were deleted
func getDeleteOrgMemberTuples(ctx context.Context, m *generated.OrgMembershipMutation, ids []string) (tuples []fga.TupleKey, err error) {
	// User the IDs of the org memberships and delete all related tuples
	for _, id := range ids {
		// this happens after soft-delete, allow the request to pull the record
		ctx := mixin.SkipSoftDelete(ctx)

		om, err := m.Client().OrgMembership.Get(ctx, id)
		if err != nil {
			return nil, err
		}

		t, err := getUserTupleKey(om.UserID, om.OrgID, "organization", om.Role)
		if err != nil {
			return nil, err
		}

		tuples = append(tuples, t)
	}

	return tuples, nil
}

// getUpdateOrgMemberTuples gets all tuples related to the orgMembership IDs that were updated
func getUpdateOrgMemberTuples(ctx context.Context, m *generated.OrgMembershipMutation, ids []string) (writes []fga.TupleKey, deletes []fga.TupleKey, err error) {
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

	// User the IDs of the org memberships and delete all related tuples
	for _, id := range ids {
		// this happens after soft-delete, allow the request to pull the record
		om, err := m.Client().OrgMembership.Get(ctx, id)
		if err != nil {
			return writes, deletes, err
		}

		d, err := getUserTupleKey(om.UserID, om.OrgID, "organization", oldRole)
		if err != nil {
			return writes, deletes, err
		}

		deletes = append(deletes, d)

		w, err := getUserTupleKey(om.UserID, om.OrgID, "organization", newRole)
		if err != nil {
			return writes, deletes, err
		}

		writes = append(writes, w)
	}

	return writes, deletes, nil
}
