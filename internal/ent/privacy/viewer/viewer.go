package viewer

import (
	"context"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/httpserve/middleware/auth"
)

// ViewerContextKey is the context key for the viewer-context
var ViewerContextKey = &ContextKey{"ViewerContextKey"}

// ContextKey is the key name for the additional context
type ContextKey struct {
	name string
}

// Viewer describes the query/mutation viewer-context
type Viewer interface {
	GetOrganizationID() string
	GetGroupID() string
	IsAdmin() bool
	GetID() (id string, exists bool)
}

// UserViewer describes a user-viewer
type UserViewer struct {
	Viewer
	GroupID string
	OrgID   string
	id      string
	hasID   bool
}

// NewUserViewerFromUser function is used to create a new `UserViewer` instance based on a
// `generated.User` object - this function is useful when you have a user object and want
// to create a `UserViewer` from it
func NewUserViewerFromUser(user *generated.User) *UserViewer {
	if user == nil {
		return NewUserViewerFromID("", false)
	}

	return NewUserViewerFromID(user.ID, true)
}

// NewUserViewerFromID  ets the `id` and `hasID` fields of the `UserViewer`
// struct and  is used to create a `UserViewer` when the user ID is known, but
// the actual user object is not available
func NewUserViewerFromID(id string, hasID bool) *UserViewer {
	return &UserViewer{
		id:    id,
		hasID: hasID,
	}
}

// NewUserViewerFromSubject function creates a new `UserViewer` instance based on the user ID obtained from the context. It uses the `auth.GetUserIDFromContext` function to
// retrieve the user ID from the context
func NewUserViewerFromSubject(c context.Context) *UserViewer {
	id, err := auth.GetUserIDFromContext(c)
	if err != nil {
		return &UserViewer{
			id:    id,
			hasID: false,
		}
	}

	return &UserViewer{
		id:    id,
		hasID: true,
	}
}

// GetOrganizationID returns the ID of the organization.
func (u UserViewer) GetOrganizationID() string {
	return u.OrgID
}

// GetGroupID returns the ID of the group
func (u UserViewer) GetGroupID() string {
	return u.GroupID
}

// IsAdmin returns a boolean value indicating whether the user is an admin or not. In this implementation, the function
// always returns `false`, indicating that the user is not an admin
func (u UserViewer) IsAdmin() bool {
	return false
}

// GetID() returns the ID of the user and a boolean value indicating whether the ID exists or not
func (u UserViewer) GetID() (string, bool) {
	return u.id, u.hasID
}

// FromContext returns the Viewer stored in a context.
func FromContext(ctx context.Context) Viewer {
	v, _ := ctx.Value(ViewerContextKey).(Viewer)

	return v
}

// NewContext returns a copy of parent context with the given Viewer attached with it.
func NewContext(parent context.Context, v Viewer) context.Context {
	return context.WithValue(parent, ViewerContextKey, v)
}
