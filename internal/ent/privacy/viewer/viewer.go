package viewer

import (
	"context"

	"github.com/datumforge/datum/internal/ent/generated"

	"github.com/google/uuid"
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
	Admin() bool                     // If viewer is admin.
	Organization() (uuid.UUID, bool) // Tenant identifier.
}

// UserViewer describes a user-viewer.
type UserViewer struct {
	T    *generated.Organization
	Role Role // Attached roles.
}

// Admin of the UserViewer
func (v UserViewer) Admin() bool {
	return v.Role&Admin != 0
}

// Organization of the UserViewer
func (v UserViewer) Organization() (uuid.UUID, bool) {
	if v.T != nil {
		return v.T.ID, true
	}

	return uuid.UUID{}, false
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
