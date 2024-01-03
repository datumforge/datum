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

// UserViewer describes a user-viewer.
type UserViewer struct {
	Viewer
	GroupID string
	OrgID   string
	id      string
	hasID   bool
}

func NewUserViewerFromUser(user *generated.User) *UserViewer {
	if user == nil {
		return NewUserViewerFromID("", false)
	}

	return NewUserViewerFromID(user.ID, true)
}

func NewUserViewerFromID(id string, hasID bool) *UserViewer {
	return &UserViewer{
		id:    id,
		hasID: hasID,
	}
}

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

func (u UserViewer) IsAdmin() bool {
	return false
}

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
