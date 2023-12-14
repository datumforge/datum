package store

import (
	"context"
	"fmt"
	"time"

	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/session"
)

// AuthSession has a single field `client` of type
// `*ent.Client`, which is a client for interacting with the database.
type AuthSession struct {
	client *ent.Client
}

// NewAuthSession function creates a new instance of the AuthSessions struct
func NewAuthSession(c *ent.Client) AuthSessions {
	return &AuthSession{
		client: c,
	}
}

// AuthSessions is defining an interface named AuthSessions
type AuthSessions interface {
	StoreSession(ctx context.Context, sessionID string, userID *ent.User) error
	GetUserIDFromSession(ctx context.Context, sessionID string) (string, error)
	DeleteSession(ctx context.Context, sessionID string) error
	GetExpiryFromSession(ctx context.Context, sessionID string) (time.Time, error)
}

// StoreSession is used to store a session in the database
func (sess *AuthSession) StoreSession(ctx context.Context, sessionID string, userID *ent.User) error {
	_, err := sess.client.Session.
		Create().
		SetUserID(userID.ID).
		SetID(sessionID).
		SetExpiresAt(time.Now().Add(time.Hour * 24 * 1)).
		Save(ctx)

	if err != nil {
		return fmt.Errorf("loginSessions.Create: %w", err)
	}

	return nil
}

// DeleteSession is used to delete a session from the database
func (sess *AuthSession) DeleteSession(ctx context.Context, sessionID string) error {
	_, err := sess.client.Session.
		Delete().
		Where(session.ID(sessionID)).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("loginsession.Delete: %w", err)
	}

	return nil
}

// GetUserIDFromSession is used to retrieve the user ID associated with a session from the database
func (sess *AuthSession) GetUserIDFromSession(ctx context.Context, sessionID string) (string, error) {
	session, err := sess.client.Session.
		Query().
		Where(session.ID(sessionID)).
		Only(ctx)

	if err != nil {
		return session.UserID, fmt.Errorf("loginSessions.Query: %w", err)
	}

	return session.UserID, nil
}

// GetExpiryFromSession is used to retrieve the expiration time of a session from the database
func (sess *AuthSession) GetExpiryFromSession(ctx context.Context, sessionID string) (time.Time, error) {
	session, err := sess.client.Session.
		Query().
		Where(session.ID(sessionID)).
		Only(ctx)

	if err != nil {
		return time.Time{}, fmt.Errorf("loginSessions.Query: %w", err)
	}

	return session.ExpiresAt, nil
}
