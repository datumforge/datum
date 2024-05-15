package graphapi

import (
	"context"
	"fmt"
	"strings"

	"github.com/datumforge/datum/internal/ent/generated"
)

var (
	environments  = []string{"production", "testing"}
	buckets       = []string{"assets", "customers", "orders", "relationships", "sales"}
	relationships = []string{"internal_users", "marketing_subscribers", "marketplaces", "partners", "vendors"}
)

func CreateWorkspace(ctx context.Context, input CreateWorkspaceInput) (*WorkspaceCreatePayload, error) {
	rootOrg, err := createRootOrganization(ctx, input)
	if err != nil {
		return nil, err
	}

	envOrgs, err := createEnvironments(ctx, rootOrg.ID, input)
	if err != nil {
		return nil, err
	}

	for _, envOrg := range envOrgs {
		bucketOrgs, err := createBuckets(ctx, envOrg.ID, envOrg.DisplayName, input)
		if err != nil {
			return nil, err
		}

		for _, bucketOrg := range bucketOrgs {
			if bucketOrg.DisplayName == "relationships" {
				_, err := createRelationships(ctx, bucketOrg.ID, envOrg.DisplayName, input)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	return &WorkspaceCreatePayload{
		Workspace: rootOrg,
	}, nil
}

// createRootOrganization creates the root organization for the workspace
func createRootOrganization(ctx context.Context, input CreateWorkspaceInput) (*generated.Organization, error) {
	// create the settings for the root organization
	settingsInput := generated.CreateOrganizationSettingInput{
		Tags:    []string{"root"},
		Domains: input.Domains,
	}

	settings, err := withTransactionalMutation(ctx).OrganizationSetting.Create().SetInput(settingsInput).Save(ctx)
	if err != nil {
		return nil, err
	}

	// create the root organization
	rootOrgInput := generated.CreateOrganizationInput{
		Name:        input.Name,
		DisplayName: &input.Name,
		Description: input.Description,
		SettingID:   &settings.ID,
	}

	return withTransactionalMutation(ctx).Organization.Create().SetInput(rootOrgInput).Save(ctx)
}

// createChildOrganizations creates the child organizations for the workspace
func createChildOrganizations(ctx context.Context, namePrefix, parentOrgID string, childNames, additionalTags []string) ([]*generated.Organization, error) {
	builder := make([]*generated.OrganizationCreate, len(childNames))
	for i, childName := range childNames {
		// create the settings for the root organization
		tags := append(additionalTags, childName)
		settingsInput := generated.CreateOrganizationSettingInput{
			Tags: tags,
		}

		settings, err := withTransactionalMutation(ctx).OrganizationSetting.Create().SetInput(settingsInput).Save(ctx)
		if err != nil {
			return nil, err
		}

		input := generated.CreateOrganizationInput{
			Name:        strings.ToLower(fmt.Sprintf("%s.%s", namePrefix, childName)),
			DisplayName: &childName,
			ParentID:    &parentOrgID,
			SettingID:   &settings.ID,
		}

		builder[i] = withTransactionalMutation(ctx).Organization.Create().SetInput(input)
	}

	return withTransactionalMutation(ctx).Organization.CreateBulk(builder...).Save(ctx)
}

func createEnvironments(ctx context.Context, rootOrgID string, input CreateWorkspaceInput) ([]*generated.Organization, error) {
	return createChildOrganizations(ctx, input.Name, rootOrgID, environments, []string{})
}

func createBuckets(ctx context.Context, envOrgID, environment string, input CreateWorkspaceInput) ([]*generated.Organization, error) {
	return createChildOrganizations(ctx, fmt.Sprintf("%s.%s", input.Name, environment), envOrgID, buckets, []string{environment})
}

func createRelationships(ctx context.Context, relationshipOrgID, environment string, input CreateWorkspaceInput) ([]*generated.Organization, error) {
	return createChildOrganizations(ctx, fmt.Sprintf("%s.%s.%s", input.Name, environment, "relationships"), relationshipOrgID, relationships, []string{environment, "relationships"})
}
