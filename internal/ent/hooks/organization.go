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

// HookOrganization runs on organization mutations to setup or remove relationship tuples
func HookOrganization() ent.Hook {
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

				// TODO: rollback mutation if tuple creation fails
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
