package hooks

import (
	"context"

	"entgo.io/ent"
	ph "github.com/posthog/posthog-go"

	"github.com/datumforge/datum/internal/analytics"
	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/hook"
	"github.com/datumforge/datum/pkg/auth"
)

// HookOrganization runs on org mutations to set default values that are not provided
func HookOrganization() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.OrganizationFunc(func(ctx context.Context, mutation *generated.OrganizationMutation) (generated.Value, error) {
			if mutation.Op().Is(ent.OpCreate) {
				// if this is empty generate a default org setting schema
				_, exists := mutation.SettingID()
				if !exists {
					// sets up default org settings using schema defaults
					orgSettingID, err := defaultOrganizationSettings(ctx, mutation)
					if err != nil {
						mutation.Logger.Errorw("error creating default organization settings", "error", err)

						return nil, err
					}

					// add the org setting ID to the input
					mutation.SetSettingID(orgSettingID)
				}

				// check if this is a child org, error if parent org is a personal org
				if err := personalOrgNoChildren(ctx, mutation); err != nil {
					return nil, err
				}
			}

			if name, ok := mutation.Name(); ok {
				if displayName, ok := mutation.DisplayName(); ok {
					if displayName == "" {
						mutation.SetDisplayName(name)
					}
				}
			}

			v, err := next.Mutate(ctx, mutation)
			if err != nil {
				return v, err
			}

			orgCreated, ok := v.(*generated.Organization)
			if !ok {
				return nil, err
			}

			if mutation.Op().Is(ent.OpCreate) {
				if err := createOrgMemberOwner(ctx, orgCreated.ID, mutation); err != nil {
					return v, err
				}

				props := ph.NewProperties().
					Set("organization_name", orgCreated.Name)

				analytics.NewOrganization(orgCreated.ID, orgCreated.CreatedBy, props)
				analytics.OrganizationProperties(orgCreated.ID, props)
			}

			return v, err
		})
	}, ent.OpCreate|ent.OpUpdateOne)
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

// personalOrgNoChildren checks if the mutation is for a child org, and if so returns an error
// if the parent org is a personal org
func personalOrgNoChildren(ctx context.Context, mutation *generated.OrganizationMutation) error {
	// check if this is a child org, error if parent org is a personal org
	parentOrgID, ok := mutation.ParentID()
	if ok {
		// check if parent org is a personal org
		parentOrg, err := mutation.Client().Organization.Get(ctx, parentOrgID)
		if err != nil {
			return err
		}

		if parentOrg.PersonalOrg {
			return ErrPersonalOrgsNoChildren
		}
	}

	return nil
}

func createOrgMemberOwner(ctx context.Context, oID string, m *generated.OrganizationMutation) error {
	// This is handled by the user create hook for personal orgs
	personalOrg, _ := m.PersonalOrg()
	if personalOrg {
		return nil
	}

	// get userID from context
	userID, err := auth.GetUserIDFromContext(ctx)
	if err != nil {
		m.Logger.Errorw("unable to get user id from echo context, unable to add user to organization")

		return err
	}

	// Add User as owner of organization
	owner := enums.RoleOwner
	input := generated.CreateOrgMembershipInput{
		UserID:         userID,
		OrganizationID: oID,
		Role:           &owner,
	}

	if _, err := m.Client().OrgMembership.Create().SetInput(input).Save(ctx); err != nil {
		m.Logger.Errorw("error creating org membership for owner", "error", err)

		return err
	}

	return nil
}
