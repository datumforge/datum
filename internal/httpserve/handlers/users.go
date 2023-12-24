package handlers

import (
	"database/sql"
	"time"

	"github.com/datumforge/datum/internal/tokens"
)

func (u *User) GetVerificationToken() string {
	if u.EmailVerificationToken.Valid {
		return u.EmailVerificationToken.String
	}

	return ""
}

func (u *User) GetVerificationExpires() (time.Time, error) {
	if u.EmailVerificationExpires.Valid {
		return time.Parse(time.RFC3339Nano, u.EmailVerificationExpires.String)
	}

	return time.Time{}, nil
}

func (u *User) CreateVerificationToken() (err error) {
	var (
		verify *tokens.VerificationToken
		token  string
		secret []byte
	)

	// Create a unique token from the user's email address
	if verify, err = tokens.NewVerificationToken(u.Email); err != nil {
		return err
	}

	// Sign the token to ensure that we can verify it later
	if token, secret, err = verify.Sign(); err != nil {
		return err
	}

	u.EmailVerificationToken = sql.NullString{Valid: true, String: token}
	u.EmailVerificationExpires = sql.NullString{Valid: true, String: verify.ExpiresAt.Format(time.RFC3339Nano)}
	u.EmailVerificationSecret = secret

	return nil
}

func (u *User) CreateResetToken() (err error) {
	var (
		reset  *tokens.ResetToken
		token  string
		secret []byte
	)

	if reset, err = tokens.NewResetToken(u.ID); err != nil {
		return err
	}

	if token, secret, err = reset.Sign(); err != nil {
		return err
	}

	u.EmailVerificationToken = sql.NullString{Valid: true, String: token}
	u.EmailVerificationExpires = sql.NullString{Valid: true, String: reset.ExpiresAt.Format(time.RFC3339Nano)}
	u.EmailVerificationSecret = secret

	return nil
}
