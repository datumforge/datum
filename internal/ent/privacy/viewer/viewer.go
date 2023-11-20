package viewer

import (
	"context"
	"fmt"
	"time"

	ofgaclient "github.com/openfga/go-sdk/client"

	"github.com/datumforge/datum/internal/fga"
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
	Admin(ctx context.Context) bool // If viewer is admin.
}

// UserViewer describes a user-viewer.
type UserViewer struct {
	UserID string
	Authz  *fga.Client
	OrgID  string
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
func (u UserViewer) Admin(ctx context.Context) bool {
	object := fmt.Sprintf("organization:%s", u.OrgID)

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second) //nolint:gomnd
	defer cancel()

	key := &ofgaclient.ClientCheckRequest{
		User:             fmt.Sprintf("user:%s", u.GetUserID()),
		Relation:         fga.OwnerRelation,
		Object:           object,
		ContextualTuples: nil,
	}

	admin, _ := u.Authz.CheckTuple(ctx, *key)

	u.Authz.Logger.Infow("authz check", "admin", admin)

	return admin
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
