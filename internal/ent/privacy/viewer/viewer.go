package viewer

import (
	"context"
)

// ViewerContextKey is the context key for the viewer-context
var ViewerContextKey = &ContextKey{"ViewerContextKey"}

// ContextKey is the key name for the additional context
type ContextKey struct {
	name string
}

// Viewer describes the query/mutation viewer-context.
type Viewer interface {
	// OrganizationID returns the organization ID from the context
	OrganizationID() string
}

// UserViewer describes a user-viewer.
type UserViewer struct {
	OrgID string
}

// OrganizationID returns the ID of the organization.
func (u UserViewer) OrganizationID() string {
	return u.OrgID
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
