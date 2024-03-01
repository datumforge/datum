package handlers

import (
	"context"
	"time"

	"github.com/datumforge/datum/internal/ent/enums"
	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/emailverificationtoken"
	"github.com/datumforge/datum/internal/ent/generated/invite"
	"github.com/datumforge/datum/internal/ent/generated/passwordresettoken"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/generated/user"
	"github.com/datumforge/datum/internal/ent/generated/usersetting"
	"github.com/datumforge/datum/pkg/middleware/transaction"
)

// updateUserLastSeen updates the last seen timestamp of the user
func (h *Handler) updateUserLastSeen(ctx context.Context, id string) error {
	if _, err := transaction.FromContext(ctx).
		User.
		UpdateOneID(id).
		SetLastSeen(time.Now()).
		Save(ctx); err != nil {
		h.Logger.Errorw("error updating user last seen", "error", err)

		return err
	}

	return nil
}

// createUser creates a user in the database based on the input and returns the user with user settings
func (h *Handler) createUser(ctx context.Context, input ent.CreateUserInput) (*ent.User, error) {
	meowuser, err := transaction.FromContext(ctx).User.Create().
		SetInput(input).
		Save(ctx)
	if err != nil {
		h.Logger.Errorw("error creating new user", "error", err)

		return nil, err
	}

	return meowuser, nil
}

// createEmailVerificationToken creates a new email verification for the user
func (h *Handler) createEmailVerificationToken(ctx context.Context, user *User) (*ent.EmailVerificationToken, error) {
	ttl, err := time.Parse(time.RFC3339Nano, user.EmailVerificationExpires.String)
	if err != nil {
		h.Logger.Errorw("unable to parse ttl", "error", err)
		return nil, err
	}

	meowtoken, err := transaction.FromContext(ctx).EmailVerificationToken.Create().
		SetOwnerID(user.ID).
		SetToken(user.EmailVerificationToken.String).
		SetTTL(ttl).
		SetEmail(user.Email).
		SetSecret(user.EmailVerificationSecret).
		Save(ctx)
	if err != nil {
		h.Logger.Errorw("error creating email verification token", "error", err)

		return nil, err
	}

	return meowtoken, nil
}

func (h *Handler) createPasswordResetToken(ctx context.Context, user *User) (*ent.PasswordResetToken, error) {
	ttl, err := time.Parse(time.RFC3339Nano, user.PasswordResetExpires.String)
	if err != nil {
		h.Logger.Errorw("unable to parse ttl", "error", err)
		return nil, err
	}

	meowtoken, err := transaction.FromContext(ctx).PasswordResetToken.Create().
		SetOwnerID(user.ID).
		SetToken(user.PasswordResetToken.String).
		SetTTL(ttl).
		SetEmail(user.Email).
		SetSecret(user.PasswordResetSecret).
		Save(ctx)
	if err != nil {
		h.Logger.Errorw("error creating password reset token", "error", err)

		return nil, err
	}

	return meowtoken, nil
}

// getUserByEVToken returns the ent user with the user settings and email verification token fields based on the
// token in the request
func (h *Handler) getUserByEVToken(ctx context.Context, token string) (*ent.User, error) {
	user, err := transaction.FromContext(ctx).EmailVerificationToken.Query().
		Where(
			emailverificationtoken.Token(token),
		).
		QueryOwner().WithSetting().WithEmailVerificationTokens().Only(ctx)
	if err != nil {
		h.Logger.Errorw("error obtaining user from email verification token", "error", err)

		return nil, err
	}

	return user, nil
}

// getUserByResetToken returns the ent user with the user settings and password reset tokens based on the
// token in the request
func (h *Handler) getUserByResetToken(ctx context.Context, token string) (*ent.User, error) {
	user, err := transaction.FromContext(ctx).PasswordResetToken.Query().
		Where(
			passwordresettoken.Token(token),
		).
		QueryOwner().WithSetting().WithPasswordResetTokens().Only(ctx)
	if err != nil {
		h.Logger.Errorw("error obtaining user from reset token", "error", err)

		return nil, err
	}

	return user, nil
}

// getUserByEmail returns the ent user with the user settings based on the email and auth provider in the request
func (h *Handler) getUserByEmail(ctx context.Context, email string, authProvider enums.AuthProvider) (*ent.User, error) {
	user, err := transaction.FromContext(ctx).User.Query().WithSetting().
		Where(user.Email(email)).
		Where(user.AuthProviderEQ(authProvider)).
		Only(ctx)
	if err != nil {
		h.Logger.Errorw("error obtaining user from email", "error", err)

		return nil, err
	}

	return user, nil
}

// getUserBySub returns the ent user with the user settings based on the subject in the claim
func (h *Handler) getUserBySub(ctx context.Context, subject string) (*ent.User, error) {
	// check user in the database, sub == claims subject and ensure only one record is returned
	user, err := transaction.FromContext(ctx).User.Query().WithSetting().Where(
		user.Sub(subject),
	).Only(ctx)
	if err != nil {
		h.Logger.Errorf("error retrieving user", "error", err)

		return nil, err
	}

	return user, nil
}

// getUserByInviteToken returns the ent user based on the invite token in the request
func (h *Handler) getUserByInviteToken(ctx context.Context, token string) (*ent.Invite, error) {
	recipient, err := transaction.FromContext(ctx).Invite.Query().
		Where(
			invite.Token(token),
		).WithOwner().Only(ctx)

	if err != nil {
		h.Logger.Errorw("error obtaining user from token", "error", err)

		return nil, err
	}

	return recipient, err
}

// countVerificationTokensUserByEmail counts number of existing email verification attempts before issuing a new one
func (h *Handler) countVerificationTokensUserByEmail(ctx context.Context, email string) (int, error) {
	attempts, err := transaction.FromContext(ctx).EmailVerificationToken.Query().WithOwner().Where(
		emailverificationtoken.And(
			emailverificationtoken.Email(email),
		)).Count(ctx)
	if err != nil {
		h.Logger.Errorw("error counting verification reset tokens", "error", err)

		return 0, err
	}

	return attempts, nil
}

// expireAllVerificationTokensUserByEmail expires all existing email verification tokens before issuing a new one
func (h *Handler) expireAllVerificationTokensUserByEmail(ctx context.Context, email string) error {
	prs, err := transaction.FromContext(ctx).EmailVerificationToken.Query().WithOwner().Where(
		emailverificationtoken.And(
			emailverificationtoken.Email(email),
			emailverificationtoken.TTLGT(time.Now()),
		)).All(ctx)
	if err != nil {
		h.Logger.Errorw("error obtaining verification reset tokens", "error", err)

		return err
	}

	for _, pr := range prs {
		if err := pr.Update().SetTTL(time.Now()).Exec(ctx); err != nil {
			h.Logger.Errorw("error expiring verification token", "error", err)

			return err
		}
	}

	return nil
}

// expireAllResetTokensUserByEmail expires all existing password reset tokens before issuing a new one
func (h *Handler) expireAllResetTokensUserByEmail(ctx context.Context, email string) error {
	prs, err := transaction.FromContext(ctx).PasswordResetToken.Query().WithOwner().Where(
		passwordresettoken.And(
			passwordresettoken.Email(email),
			passwordresettoken.TTLGT(time.Now()),
		)).All(ctx)
	if err != nil {
		h.Logger.Errorw("error obtaining password reset tokens", "error", err)

		return err
	}

	for _, pr := range prs {
		if err := pr.Update().SetTTL(time.Now()).Exec(ctx); err != nil {
			h.Logger.Errorw("error expiring password reset token", "error", err)

			return err
		}
	}

	return nil
}

// setEmailConfirmed sets the user setting field email_confirmed to true within a transaction
func (h *Handler) setEmailConfirmed(ctx context.Context, user *ent.User) error {
	if _, err := transaction.FromContext(ctx).UserSetting.Update().SetEmailConfirmed(true).
		Where(
			usersetting.ID(user.Edges.Setting.ID),
		).Save(ctx); err != nil {
		h.Logger.Errorw("error setting email confirmed", "error", err)

		return err
	}

	return nil
}

// updateUserPassword changes a updates a user's password in the database
func (h *Handler) updateUserPassword(ctx context.Context, id string, password string) error {
	if _, err := transaction.FromContext(ctx).User.UpdateOneID(id).SetPassword(password).Save(ctx); err != nil {
		h.Logger.Errorw("error updating user password", "error", err)

		return err
	}

	return nil
}

// addDefaultOrgToUserQuery adds the default org to the user object, user must be authenticated before calling this
func (h *Handler) addDefaultOrgToUserQuery(ctx context.Context, user *ent.User) error {
	// get the default org for the user, allow access, accessible orgs will be filtered by the interceptor
	orgCtx := privacy.DecisionContext(ctx, privacy.Allow)

	org, err := user.Edges.Setting.DefaultOrg(orgCtx)
	if err != nil {
		h.Logger.Errorw("error obtaining default org", "error", err)

		return err
	}

	// add default org to user object
	user.Edges.Setting.Edges.DefaultOrg = org

	return nil
}
