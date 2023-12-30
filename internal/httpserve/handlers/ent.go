package handlers

import (
	"context"
	"net/http"
	"time"

	echo "github.com/datumforge/echox"

	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/emailverificationtoken"
	"github.com/datumforge/datum/internal/ent/generated/passwordresettoken"
	"github.com/datumforge/datum/internal/ent/generated/user"
	"github.com/datumforge/datum/internal/ent/generated/usersetting"
)

const (
	rollbackErr          = "error rolling back transaction"
	transactionStartErr  = "error starting transaction"
	transactionCommitErr = "error committing transaction"
)

type entClientCtxKey struct{}

// TransactionFromContext returns a TX Client stored inside a context, or nil if there isn't one
func TransactionFromContext(ctx context.Context) *ent.Tx {
	c, _ := ctx.Value(entClientCtxKey{}).(*ent.Tx)
	return c
}

// NewContext returns a new context with the given TX Client attached
func NewContext(parent context.Context, c *ent.Tx) context.Context {
	return context.WithValue(parent, entClientCtxKey{}, c)
}

// Transaction returns a middleware function for transactions on REST endpoints
func (h *Handler) Transaction(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		h.Logger.Debug("starting transaction in middleware")

		var err error

		client, err := h.DBClient.Tx(c.Request().Context())
		if err != nil {
			h.Logger.Errorw(transactionStartErr, "error", err)

			return c.JSON(http.StatusInternalServerError, ErrProcessingRequest)
		}

		// add to context
		ctx := NewContext(c.Request().Context(), client)

		c.SetRequest(c.Request().WithContext(ctx))

		if err := next(c); err != nil {
			h.Logger.Debug("rolling back transaction in middleware")

			if err := client.Rollback(); err != nil {
				h.Logger.Errorw(rollbackErr, "error", err)
			}

			return c.JSON(http.StatusInternalServerError, ErrProcessingRequest)
		}

		h.Logger.Debug("committing transaction in middleware")

		if err := client.Commit(); err != nil {
			h.Logger.Errorw(transactionCommitErr, "error", err)

			return c.JSON(http.StatusInternalServerError, ErrProcessingRequest)
		}

		return nil
	}
}

func (h *Handler) updateUserLastSeen(ctx context.Context, id string) error {
	if _, err := TransactionFromContext(ctx).User.Update().SetLastSeen(time.Now()).
		Where(
			user.ID(id),
		).
		Save(ctx); err != nil {
		return err
	}

	return nil
}

func (h *Handler) createUser(ctx context.Context, input ent.CreateUserInput) (*ent.User, error) {
	meowuser, err := TransactionFromContext(ctx).User.Create().
		SetInput(input).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return meowuser, nil
}

func (h *Handler) createEmailVerificationToken(ctx context.Context, user *User) (*ent.EmailVerificationToken, error) {
	ttl, err := time.Parse(time.RFC3339Nano, user.EmailVerificationExpires.String)
	if err != nil {
		h.Logger.Errorw("unable to parse ttl", "error", err)
		return nil, err
	}

	meowtoken, err := TransactionFromContext(ctx).EmailVerificationToken.Create().
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

	meowtoken, err := TransactionFromContext(ctx).PasswordResetToken.Create().
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
	user, err := TransactionFromContext(ctx).EmailVerificationToken.Query().WithOwner().
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

// getUserByEmail returns the ent user with the user settings based on the email in the request
func (h *Handler) getUserByEmail(ctx context.Context, email string) (*ent.User, error) {
	user, err := TransactionFromContext(ctx).User.Query().WithSetting().
		Where(user.Email(email)).
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
	user, err := TransactionFromContext(ctx).User.Query().WithSetting().Where(
		user.Sub(subject),
	).Only(ctx)
	if err != nil {
		h.Logger.Errorf("error retrieving user", "error", err)

		return nil, err
	}

	return user, nil
}

// expireAllVerificationTokensUserByEmail expires all existing email verification tokens before issuing a new one
func (h *Handler) expireAllVerificationTokensUserByEmail(ctx context.Context, email string) error {
	prs, err := TransactionFromContext(ctx).EmailVerificationToken.Query().WithOwner().Where(
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
	prs, err := TransactionFromContext(ctx).PasswordResetToken.Query().WithOwner().Where(
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
			if err := TransactionFromContext(ctx).Rollback(); err != nil {
				h.Logger.Errorw(rollbackErr, "error", err)
				return err
			}

			h.Logger.Errorw("error expiring password reset token", "error", err)

			return err
		}
	}

	return nil
}

// setEmailConfirmed sets the user setting field email_confirmed to true within a transaction
func (h *Handler) setEmailConfirmed(ctx context.Context, user *ent.User) error {
	if _, err := TransactionFromContext(ctx).UserSetting.Update().SetEmailConfirmed(true).
		Where(
			usersetting.ID(user.Edges.Setting.ID),
		).Save(ctx); err != nil {
		return err
	}

	return nil
}
