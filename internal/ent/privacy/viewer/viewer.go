package viewer

import (
	"context"
	"time"

	ofgaclient "github.com/openfga/go-sdk/client"

	"github.com/datumforge/datum/internal/fga"
)

// Viewer describes the query/mutation viewer-context.
type Viewer interface {
	// GetUserID returns the user ID from the context
	GetUserID() string
	// HasAccess uses the FGA client to determine access to the objcet
	HasAccess(ctx context.Context) bool
}

// UserViewer describes a user-viewer.
type UserViewer struct {
	UserID string
	Authz  *fga.Client
	Key    fga.TupleKey
}

// GetUserID returns the ID of the user.
func (u UserViewer) GetUserID() string {
	return u.UserID
}

// HasAccess of the UserViewer
func (u UserViewer) HasAccess(ctx context.Context) bool {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second) //nolint:gomnd
	defer cancel()

	check := ofgaclient.ClientCheckRequest{
		User:             u.Key.Subject.String(),
		Relation:         u.Key.Relation.String(),
		Object:           u.Key.Object.String(),
		ContextualTuples: nil, // TODO: allow contextual tuples
	}

	access, _ := u.Authz.CheckTuple(ctx, check)

	u.Authz.Logger.Infow("authz check",
		"user", u.Key.Subject.String(),
		"relation", u.Key.Relation.String(),
		"object", u.Key.Object.String(),
		"has_access", access,
	)

	return access
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
