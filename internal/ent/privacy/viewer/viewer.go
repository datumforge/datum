package viewer

import (
	"context"

	"github.com/datumforge/datum/internal/ent/generated"
)

// Role for viewer actions.
type Role int

// List of roles.
const (
	_ Role = 1 << iota
	Admin
	View
)

// Viewer describes the query/mutation viewer-context.
type Viewer interface {
	GetUser() UserViewer
	GetUserID() string
	Admin() bool // If viewer is admin.
}

// UserViewer describes a user-viewer.
type UserViewer struct {
	UserID string
	T      *generated.Organization
	Role   Role // Attached roles.
}

// GetUser returns the user information.
func (u UserViewer) GetUser() UserViewer {
	return u
}

// GetUserID returns the ID of the user.
func (u UserViewer) GetUserID() string {
	return u.UserID
}

// Admin of the UserViewer
func (u UserViewer) Admin() bool {
	return u.Role&Admin != 0
}

type ctxKey struct{}

// FromContext returns the Viewer stored in a context.
func FromContext(ctx context.Context) Viewer {
	v, _ := ctx.Value(ctxKey{}).(Viewer)

	return v
}

// NewContext returns a copy of parent context with the given Viewer attached with it.
func NewContext(parent context.Context, v Viewer) context.Context {
	return context.WithValue(parent, ctxKey{}, v)
}
