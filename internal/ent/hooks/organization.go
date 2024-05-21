package hooks

import (
	"context"

	"entgo.io/ent"

	"github.com/datumforge/fgax"
	geodeticenums "github.com/datumforge/geodetic/pkg/enums"
	geodetic "github.com/datumforge/geodetic/pkg/geodeticclient"

	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/hook"
	"github.com/datumforge/datum/pkg/auth"
	"github.com/datumforge/datum/pkg/utils/gravatar"
	"github.com/datumforge/datum/pkg/utils/marionette"
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

				url := gravatar.New(name, nil)
				mutation.SetAvatarRemoteURL(url)
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
				// create the admin organization member if not using an API token (which is not associated with a user)
				// otherwise add the API toke for admin access to the newly created organization
				if err := createOrgMemberOwner(ctx, orgCreated.ID, mutation); err != nil {
					return v, err
				}

				// create the database, if the org has a dedicated db and geodetic is available
				if orgCreated.DedicatedDb && mutation.Geodetic != nil {
					settings, err := orgCreated.Setting(ctx)
					if err != nil {
						mutation.Logger.Errorw("unable to get organization settings")

						return nil, err
					}

					if err := mutation.Marionette.Queue(marionette.TaskFunc(func(ctx context.Context) error {
						return createDatabase(ctx, orgCreated.ID, settings.GeoLocation.String(), mutation)
					}), marionette.WithErrorf("could not send create the database for %s", orgCreated.Name),
					); err != nil {
						mutation.Logger.Errorw("unable to queue database creation")

						return v, err
					}
				}
			}

			return v, err
		})
	}, ent.OpCreate|ent.OpUpdateOne)
}

func createDatabase(ctx context.Context, orgID, geo string, mutation *generated.OrganizationMutation) error {
	// set default geo if not provided
	if geo == "" {
		geo = enums.Amer.String()
	}

	input := geodetic.CreateDatabaseInput{
		OrganizationID: orgID,
		Geo:            &geo,
		Provider:       &geodeticenums.Turso,
	}

	mutation.Logger.Infow("creating database", "org", input.OrganizationID, "geo", input.Geo, "provider", input.Provider)

	if _, err := mutation.Geodetic.CreateDatabase(ctx, input); err != nil {
		mutation.Logger.Errorw("error creating database", "error", err)

		return err
	}

	// create the database
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

	// if this was created with an API token, do not create an owner but add the service tuple to fga
	if auth.IsAPITokenAuthentication(ctx) {
		return createServiceTuple(ctx, oID, m)
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

// createServiceTuple creates a service tuple for the organization and api key so the organization can be accessed
func createServiceTuple(ctx context.Context, oID string, m *generated.OrganizationMutation) error {
	// get userID from context
	subjectID, err := auth.GetUserIDFromContext(ctx)
	if err != nil {
		m.Logger.Errorw("unable to get user id from echo context, unable to add user to organization")

		return err
	}

	// allow the api token to edit the newly created organization, no other users will have access
	// so this is the minimum required access
	role := fgax.CanEdit

	// get tuple key
	tuple := fgax.GetTupleKey(subjectID, "service", oID, "organization", role)

	if _, err := m.Authz.WriteTupleKeys(ctx, []fgax.TupleKey{tuple}, nil); err != nil {
		m.Logger.Errorw("failed to create relationship tuple", "error", err)

		return err
	}

	m.Logger.Debugw("created relationship tuples", "relation", role, "object", tuple.Object)

	return nil
}
