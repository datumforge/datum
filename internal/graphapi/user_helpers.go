package graphapi

import (
	"context"
	"fmt"
	"strings"

	petname "github.com/dustinkirkland/golang-petname"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/datumforge/datum/internal/ent/generated"
)

const (
	personalOrgPrefix = "Personal Organization"
)

// getPersonalOrgInput generates the input for a new personal organization
// personal orgs are assigned to all new users when registering with Datum
func getPersonalOrgInput(user *generated.User) generated.CreateOrganizationInput {
	// caser is used to capitalize the first letter of words
	caser := cases.Title(language.AmericanEnglish)

	// generate random name for personal orgs
	name := caser.String(petname.Generate(2, " ")) //nolint:gomnd

	desc := fmt.Sprintf("%s - %s %s", personalOrgPrefix, caser.String(user.FirstName), caser.String(user.LastName))

	displayName := user.DisplayName
	if displayName == "unknown" || displayName == "" {
		displayName = fmt.Sprintf("%s%s", strings.ToLower(user.FirstName), strings.ToTitle(user.LastName))
	}

	return generated.CreateOrganizationInput{
		Name:        name,
		DisplayName: &displayName,
		Description: &desc,
	}
}

// createPersonalOrg creates an org for a user with a unique random name
func (r *mutationResolver) createPersonalOrg(ctx context.Context, user *generated.User) error {
	orgInput := getPersonalOrgInput(user)

	_, err := r.createOrg(ctx, orgInput)
	if err != nil {
		// retry on unique constraint
		if generated.IsConstraintError(err) {
			return r.createPersonalOrg(ctx, user)
		}

		r.logger.Errorw("unable to create personal org")
	}

	return err
}

// defaultUserSettings creates the default user settings for a new user
func (r *mutationResolver) defaultUserSettings(ctx context.Context) (string, error) {
	input := generated.CreateUserSettingInput{}

	userSetting, err := r.client.UserSetting.Create().SetInput(input).Save(ctx)
	if err != nil {
		return "", err
	}

	return userSetting.ID, nil
}
