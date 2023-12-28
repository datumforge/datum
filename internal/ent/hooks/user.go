package hooks

import (
	"context"
	"fmt"

	"entgo.io/ent"
	petname "github.com/dustinkirkland/golang-petname"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/hook"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/httpserve/middleware/auth"
	"github.com/datumforge/datum/internal/passwd"
	"github.com/datumforge/datum/internal/utils/gravatar"
)

const (
	personalOrgPrefix = "Personal Organization"
)

// HookUser runs on user mutations validate and hash the password and set default values that are not provided
func HookUser() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.UserFunc(func(ctx context.Context, mutation *generated.UserMutation) (generated.Value, error) {
			if password, ok := mutation.Password(); ok {
				// validate password before its encrypted
				if passwd.Strength(password) < passwd.Moderate {
					return nil, auth.ErrPasswordTooWeak
				}

				hash, err := passwd.CreateDerivedKey(password)
				if err != nil {
					return nil, err
				}

				mutation.SetPassword(hash)
			}

			if email, ok := mutation.Email(); ok {
				url := gravatar.New(email, nil)
				mutation.SetAvatarRemoteURL(url)

				// use the email as the display name, if not provided on creation
				if mutation.Op().Is(ent.OpCreate) {
					displayName, _ := mutation.DisplayName()
					if displayName == "" {
						mutation.SetDisplayName(email)
					}
				}
			}

			// user settings are required, if this is empty generate a default setting schema
			if mutation.Op().Is(ent.OpCreate) {
				settingID, _ := mutation.SettingID()
				if settingID == "" {
					// sets up default user settings using schema defaults
					userSettingID, err := defaultUserSettings(ctx, mutation)
					if err != nil {
						return nil, err
					}

					// add the user setting ID to the input
					mutation.SetSettingID(userSettingID)
				}
			}

			v, err := next.Mutate(ctx, mutation)
			if err != nil {
				return nil, err
			}

			userCreated, ok := v.(*generated.User)
			if !ok {
				return nil, ErrInternalServerError
			}

			if mutation.Op().Is(ent.OpCreate) {
				// if using password auth, set subject to user ID
				if userCreated.Password != nil {
					userCreated.Sub = userCreated.ID

					if _, err := mutation.Client().User.
						UpdateOneID(userCreated.ID).
						SetSub(userCreated.Sub).
						Save(ctx); err != nil {
						return nil, ErrInternalServerError
					}
				}

				// when a user is created, we create a personal user org
				if err := createPersonalOrg(ctx, mutation.Client(), userCreated); err != nil {
					return nil, ErrInternalServerError
				}
			}

			return v, err
		})
	}, ent.OpCreate|ent.OpUpdateOne)
}

// getPersonalOrgInput generates the input for a new personal organization
// personal orgs are assigned to all new users when registering with Datum
func getPersonalOrgInput(user *generated.User) generated.CreateOrganizationInput {
	// caser is used to capitalize the first letter of words
	caser := cases.Title(language.AmericanEnglish)

	// generate random name for personal orgs
	name := caser.String(petname.Generate(2, " ")) //nolint:gomnd
	displayName := name
	personalOrg := true
	desc := fmt.Sprintf("%s - %s %s", personalOrgPrefix, caser.String(user.FirstName), caser.String(user.LastName))

	// add user to the users of the personal organization
	users := []string{user.ID}

	return generated.CreateOrganizationInput{
		Name:        name,
		DisplayName: &displayName,
		Description: &desc,
		PersonalOrg: &personalOrg,
		UserIDs:     users,
	}
}

// createPersonalOrg creates an org for a user with a unique random name
func createPersonalOrg(ctx context.Context, dbClient *generated.Client, user *generated.User) error {
	// this prevents a privacy check that would be required for regular orgs, but not a personal org
	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	orgInput := getPersonalOrgInput(user)

	_, err := dbClient.Organization.Create().SetInput(orgInput).Save(ctx)
	if err != nil {
		// retry on unique constraint
		if generated.IsConstraintError(err) {
			return createPersonalOrg(ctx, dbClient, user)
		}

		user.Logger.Errorw("unable to create personal org", "error", err.Error())
	}

	return err
}

// defaultUserSettings creates the default user settings for a new user
func defaultUserSettings(ctx context.Context, user *generated.UserMutation) (string, error) {
	input := generated.CreateUserSettingInput{}

	userSetting, err := user.Client().UserSetting.Create().SetInput(input).Save(ctx)
	if err != nil {
		return "", err
	}

	return userSetting.ID, nil
}
