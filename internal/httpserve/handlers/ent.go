package handlers

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"

	ent "github.com/datumforge/datum/internal/ent/generated"
)

const (
	rollbackErr = "error rolling back transaction"
)

func (h *Handler) updateUserLastSeen(ctx context.Context, tx *ent.Tx, id string) error {
	if _, err := tx.User.Update().SetLastSeen(time.Now()).Where(func(s *sql.Selector) {
		s.Where(sql.EQ("id", id))
	}).Save(ctx); err != nil {
		if err := tx.Rollback(); err != nil {
			h.Logger.Errorw(rollbackErr, "error", err)
			return err
		}

		return err
	}

	return nil
}

func (h *Handler) createUser(ctx context.Context, tx *ent.Tx, input ent.CreateUserInput) (*ent.User, error) {
	meowuser, err := tx.User.Create().
		SetInput(input).
		Save(ctx)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			h.Logger.Errorw(rollbackErr, "error", err)
			return nil, err
		}

		return nil, err
	}

	return meowuser, nil
}

func (h *Handler) createEmailVerificationToken(ctx context.Context, tx *ent.Tx, user *User) (*ent.EmailVerificationToken, error) {
	ttl, err := time.Parse(time.RFC3339Nano, user.EmailVerificationExpires.String)
	if err != nil {
		h.Logger.Errorw("unable to parse ttl", "error", err)
		return nil, err
	}

	meowtoken, err := tx.EmailVerificationToken.Create().
		SetOwnerID(user.ID).
		SetToken(user.EmailVerificationToken.String).
		SetTTL(ttl).
		SetEmail(user.Email).
		SetSecret(user.EmailVerificationSecret).
		Save(ctx)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			h.Logger.Errorw(rollbackErr, "error", err)
			return nil, err
		}

		h.Logger.Errorw("error creating email verification token", "error", err)

		return nil, err
	}

	return meowtoken, nil
}

// getUserByToken returns the ent user with the user settings and email verification token fields based on the
// token in the request
func (h *Handler) getUserByToken(ctx context.Context, tx *ent.Tx, token string) (*ent.User, error) {
	user, err := tx.EmailVerificationToken.Query().WithOwner().Where(func(s *sql.Selector) {
		s.Where(sql.EQ("token", token))
	}).QueryOwner().WithSetting().WithEmailVerificationTokens().Only(ctx)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			h.Logger.Errorw(rollbackErr, "error", err)
			return nil, err
		}

		h.Logger.Errorw("error obtaining user from email verification token", "error", err)

		return nil, err
	}

	return user, nil
}

// getUserByEmail returns the ent user with the user settings based on the email in the request
func (h *Handler) getUserByEmail(ctx context.Context, tx *ent.Tx, email string) (*ent.User, error) {
	user, err := h.DBClient.User.Query().WithSetting().Where(func(s *sql.Selector) {
		s.Where(sql.EQ("email", email))
	}).Only(ctx)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			h.Logger.Errorw(rollbackErr, "error", err)
			return nil, err
		}

		h.Logger.Errorw("error obtaining user from email verification token", "error", err)

		return nil, err
	}

	return user, nil
}

// getTokenUserByEmail returns the ent user with the user settings and email verification token fields based on the
// email in the request
func (h *Handler) getTokenUserByEmail(ctx context.Context, tx *ent.Tx, email string) (*ent.User, error) {
	user, err := tx.EmailVerificationToken.Query().WithOwner().Where(func(s *sql.Selector) {
		s.Where(sql.EQ("email", email))
	}).QueryOwner().WithSetting().WithEmailVerificationTokens().Only(ctx)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			h.Logger.Errorw(rollbackErr, "error", err)
			return nil, err
		}

		h.Logger.Errorw("error obtaining user from email verification token", "error", err)

		return nil, err
	}

	return user, nil
}

// setEmailConfirmed sets the user setting field email_confirmed to true within a transaction
func (h *Handler) setEmailConfirmed(ctx context.Context, tx *ent.Tx, user *ent.User) error {
	if _, err := tx.UserSetting.Update().SetEmailConfirmed(true).Where(func(s *sql.Selector) {
		s.Where(sql.EQ("id", user.Edges.Setting.ID))
	}).Save(ctx); err != nil {
		if err := tx.Rollback(); err != nil {
			h.Logger.Errorw(rollbackErr, "error", err)
			return err
		}

		return err
	}

	return nil
}
