package hooks

import (
	"context"
	"fmt"
	"strings"

	"entgo.io/ent"
	openfga "github.com/openfga/go-sdk"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/hook"
	"github.com/datumforge/datum/internal/ent/mixin"
	"github.com/datumforge/datum/internal/fga"
)

// HookGroup runs on group mutations to set default values that are not provided
func HookGroup() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.GroupFunc(func(ctx context.Context, mutation *generated.GroupMutation) (generated.Value, error) {
			if name, ok := mutation.Name(); ok {
				displayName, _ := mutation.DisplayName()

				if displayName == "" {
					mutation.SetDisplayName(name)
				}
			}

			if mutation.Op().Is(ent.OpCreate) {
				// if this is empty generate a default group setting schema
				settingID, _ := mutation.SettingID()
				if settingID == "" {
					// sets up default group settings using schema defaults
					groupSettingID, err := defaultGroupSettings(ctx, mutation)
					if err != nil {
						return nil, err
					}

					// add the group setting ID to the input
					mutation.SetSettingID(groupSettingID)
				}
			}

			return next.Mutate(ctx, mutation)
		})
	}, ent.OpCreate|ent.OpUpdateOne)
}

// HookGroupAuthz runs on group mutations to setup or remove relationship tuples
func HookGroupAuthz() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return hook.GroupFunc(func(ctx context.Context, m *generated.GroupMutation) (ent.Value, error) {
			// do the mutation, and then create/delete the relationship
			retValue, err := next.Mutate(ctx, m)
			if err != nil {
				// if we error, do not attempt to create the relationships
				return retValue, err
			}

			if m.Op().Is(ent.OpCreate) {
				// create the relationship tuple for the admin
				err = groupCreateHook(ctx, m)
			} else if m.Op().Is(ent.OpDelete|ent.OpDeleteOne) || mixin.CheckIsSoftDelete(ctx) {
				// delete all relationship tuples on delete, or soft delete (Update Op)
				err = groupDeleteHook(ctx, m)
			}

			return retValue, err
		})
	}
}

func groupCreateHook(ctx context.Context, m *generated.GroupMutation) error {
	if m.Authz.Ofga != nil {
		objID, exists := m.ID()
		objType := strings.ToLower(m.Type())
		object := fmt.Sprintf("%s:%s", objType, objID)
		org, orgexists := m.OwnerID()

		if exists && orgexists {
			m.Logger.Infow("creating admin relationship tuples", "relation", fga.AdminRelation, "object", object)

			tupleGroup := []openfga.TupleKey{}

			adminTuples, err := createTuple(ctx, &m.Authz, fga.AdminRelation, object)

			if err != nil {
				return err
			}

			tupleGroup = append(tupleGroup, adminTuples...)

			m.Logger.Infow("creating parent relationship tuples", "relation", fga.ParentRelation, "org", org, "object", object)

			orgTuples, err := createOrgTuple(ctx, &m.Authz, org, fga.ParentRelation, object)

			if err != nil {
				return err
			}

			tupleGroup = append(tupleGroup, orgTuples...)

			if _, err := m.Authz.CreateRelationshipTuple(ctx, tupleGroup); err != nil {
				m.Logger.Errorw("failed to create relationship tuple", "error", err)

				return ErrInternalServerError
			}
		}
	}

	return nil
}

func groupDeleteHook(ctx context.Context, m *generated.GroupMutation) error {
	// Add relationship tuples if authz is enabled
	if m.Authz.Ofga != nil {
		objID, _ := m.ID()
		objType := strings.ToLower(m.Type())
		object := fmt.Sprintf("%s:%s", objType, objID)

		m.Logger.Infow("deleting relationship tuples", "object", object)

		// Add relationship tuples if authz is enabled
		if m.Authz.Ofga != nil {
			if err := m.Authz.DeleteAllObjectRelations(ctx, object); err != nil {
				m.Logger.Errorw("failed to delete relationship tuples", "error", err)

				return ErrInternalServerError
			}

			m.Logger.Infow("deleted relationship tuples", "object", object)
		}
	}

	return nil
}

// defaultGroupSettings creates the default group settings for a new group
func defaultGroupSettings(ctx context.Context, group *generated.GroupMutation) (string, error) {
	input := generated.CreateGroupSettingInput{}

	groupSetting, err := group.Client().GroupSetting.Create().SetInput(input).Save(ctx)
	if err != nil {
		return "", err
	}

	return groupSetting.ID, nil
}
