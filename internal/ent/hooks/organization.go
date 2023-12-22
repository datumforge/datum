package hooks

import (
	"context"
	"fmt"
	"strings"

	"entgo.io/ent"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/hook"
	"github.com/datumforge/datum/internal/ent/mixin"
	"github.com/datumforge/datum/internal/fga"
)

// HookOrganization runs on org mutations to set default values that are not provided
func HookOrganization() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.OrganizationFunc(func(ctx context.Context, mutation *generated.OrganizationMutation) (generated.Value, error) {
			if mutation.Op().Is(ent.OpCreate) {
				// if this is empty generate a default org setting schema
				settingID, _ := mutation.SettingID()
				if settingID == "" {
					// sets up default org settings using schema defaults
					orgSettingID, err := defaultOrganizationSettings(ctx, mutation)
					if err != nil {
						return nil, err
					}

					// add the org setting ID to the input
					mutation.SetSettingID(orgSettingID)
				}

				// check if this is a child org, error if parent org is a personal org
				parentOrgID, ok := mutation.ParentID()
				if ok {
					// check if parent org is a personal org
					parentOrg, err := mutation.Client().Organization.Get(ctx, parentOrgID)
					if err != nil {
						return nil, err
					}

					if parentOrg.PersonalOrg {
						return nil, ErrPersonalOrgsNoChildren
					}
				}
			}

			if name, ok := mutation.Name(); ok {
				if displayName, ok := mutation.DisplayName(); ok {
					if displayName == "" {
						mutation.SetDisplayName(name)
					}
				}
			}
			return next.Mutate(ctx, mutation)
		})
	}, ent.OpCreate|ent.OpUpdateOne)
}

// HookOrganizationAuthz runs on organization mutations to setup or remove relationship tuples
func HookOrganizationAuthz() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return hook.OrganizationFunc(func(ctx context.Context, m *generated.OrganizationMutation) (ent.Value, error) {
			// do the mutation, and then create/delete the relationship
			retValue, err := next.Mutate(ctx, m)
			if err != nil {
				// if we error, do not attempt to create the relationships
				return retValue, err
			}

			if m.Op().Is(ent.OpCreate) {
				// create the relationship tuple for the owner
				err = organizationCreateHook(ctx, m)
			} else if m.Op().Is(ent.OpDelete|ent.OpDeleteOne) || mixin.CheckIsSoftDelete(ctx) {
				// delete all relationship tuples on delete, or soft delete (Update Op)
				err = organizationDeleteHook(ctx, m)
			}

			return retValue, err
		})
	}
}

func organizationCreateHook(ctx context.Context, m *generated.OrganizationMutation) error {
	// Add relationship tuples if authz is enabled
	if m.Authz.Ofga != nil {
		objID, exists := m.ID()
		objType := strings.ToLower(m.Type())
		object := fmt.Sprintf("%s:%s", objType, objID)

		m.Logger.Infow("creating relationship tuples", "relation", fga.OwnerRelation, "object", object)

		if exists {
			tuples, err := createTuple(ctx, &m.Authz, fga.OwnerRelation, object)
			if err != nil {
				return err
			}

			if _, err := m.Authz.CreateRelationshipTuple(ctx, tuples); err != nil {
				m.Logger.Errorw("failed to create relationship tuple", "error", err)

				return ErrInternalServerError
			}
		}

		m.Logger.Infow("created relationship tuples", "relation", fga.OwnerRelation, "object", object)
	}

	return nil
}

func organizationDeleteHook(ctx context.Context, m *generated.OrganizationMutation) error {
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

// defaultOrganizationSettings creates the default organizations settings for a new org
func defaultOrganizationSettings(ctx context.Context, mutation *generated.OrganizationMutation) (string, error) {
	input := generated.CreateOrganizationSettingInput{}

	organizationSetting, err := mutation.Client().OrganizationSetting.Create().SetInput(input).Save(ctx)
	if err != nil {
		return "", err
	}

	return organizationSetting.ID, nil
}
