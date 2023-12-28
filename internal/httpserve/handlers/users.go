package handlers

import (
	"database/sql"
	"time"

	"github.com/oklog/ulid/v2"

	"github.com/datumforge/datum/internal/tokens"
)

// User holds data specific to the datum user for the REST handlers for
// login, registration, verification, etc
type User struct {
	Username                 string `json:"username"`
	Password                 string `json:"password"`
	userID                   string
	FirstName                string
	LastName                 string
	Name                     string
	Email                    string
	EmailVerificationExpires sql.NullString
	EmailVerificationToken   sql.NullString
	EmailVerificationSecret  []byte
}

// GetVerificationToken returns the verification token if its valid
func (u *User) GetVerificationToken() string {
	if u.EmailVerificationToken.Valid {
		return u.EmailVerificationToken.String
	}

	return ""
}

// GetVerificationExpires returns the expiration time of email verification token
func (u *User) GetVerificationExpires() (time.Time, error) {
	if u.EmailVerificationExpires.Valid {
		return time.Parse(time.RFC3339Nano, u.EmailVerificationExpires.String)
	}

	return time.Time{}, nil
}

// CreateVerificationToken creates a new email verification token for the user
func (u *User) CreateVerificationToken() error {
	// Create a unique token from the user's email address
	verify, err := tokens.NewVerificationToken(u.Email)
	if err != nil {
		return err
	}

	// Sign the token to ensure that we can verify it later
	token, secret, err := verify.Sign()
	if err != nil {
		return err
	}

	u.EmailVerificationToken = sql.NullString{Valid: true, String: token}
	u.EmailVerificationExpires = sql.NullString{Valid: true, String: verify.ExpiresAt.Format(time.RFC3339Nano)}
	u.EmailVerificationSecret = secret

	return nil
}

// CreateResetToken creates a new reset token for the user
func (u *User) CreateResetToken() error {
	uid, err := ulid.Parse(u.userID)
	if err != nil {
		return err
	}

	reset, err := tokens.NewResetToken(uid)
	if err != nil {
		return err
	}

	token, secret, err := reset.Sign()
	if err != nil {
		return err
	}

	u.EmailVerificationToken = sql.NullString{Valid: true, String: token}
	u.EmailVerificationExpires = sql.NullString{Valid: true, String: reset.ExpiresAt.Format(time.RFC3339Nano)}
	u.EmailVerificationSecret = secret

	return nil
}
