package viewer

import (
	"context"
)

// Viewer describes the query/mutation viewer-context.
type Viewer interface {
	// GetObjectID returns the object ID from the context
	GetObjectID() string
}

// UserViewer describes a user-viewer.
type UserViewer struct {
	ObjectID string
}

// GetObjectID returns the ID of the object.
func (u UserViewer) GetObjectID() string {
	return u.ObjectID
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
