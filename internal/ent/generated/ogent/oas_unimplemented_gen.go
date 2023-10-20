// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// CreateIntegration implements createIntegration operation.
//
// Creates a new Integration and persists it to storage.
//
// POST /integrations
func (UnimplementedHandler) CreateIntegration(ctx context.Context, req *CreateIntegrationReq) (r CreateIntegrationRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateMembership implements createMembership operation.
//
// Creates a new Membership and persists it to storage.
//
// POST /memberships
func (UnimplementedHandler) CreateMembership(ctx context.Context, req *CreateMembershipReq) (r CreateMembershipRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateOrganization implements createOrganization operation.
//
// Creates a new Organization and persists it to storage.
//
// POST /organizations
func (UnimplementedHandler) CreateOrganization(ctx context.Context, req *CreateOrganizationReq) (r CreateOrganizationRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateSession implements createSession operation.
//
// Creates a new Session and persists it to storage.
//
// POST /sessions
func (UnimplementedHandler) CreateSession(ctx context.Context, req *CreateSessionReq) (r CreateSessionRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateUser implements createUser operation.
//
// Creates a new User and persists it to storage.
//
// POST /users
func (UnimplementedHandler) CreateUser(ctx context.Context, req *CreateUserReq) (r CreateUserRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteIntegration implements deleteIntegration operation.
//
// Deletes the Integration with the requested ID.
//
// DELETE /integrations/{id}
func (UnimplementedHandler) DeleteIntegration(ctx context.Context, params DeleteIntegrationParams) (r DeleteIntegrationRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteMembership implements deleteMembership operation.
//
// Deletes the Membership with the requested ID.
//
// DELETE /memberships/{id}
func (UnimplementedHandler) DeleteMembership(ctx context.Context, params DeleteMembershipParams) (r DeleteMembershipRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteOrganization implements deleteOrganization operation.
//
// Deletes the Organization with the requested ID.
//
// DELETE /organizations/{id}
func (UnimplementedHandler) DeleteOrganization(ctx context.Context, params DeleteOrganizationParams) (r DeleteOrganizationRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteSession implements deleteSession operation.
//
// Deletes the Session with the requested ID.
//
// DELETE /sessions/{id}
func (UnimplementedHandler) DeleteSession(ctx context.Context, params DeleteSessionParams) (r DeleteSessionRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteUser implements deleteUser operation.
//
// Deletes the User with the requested ID.
//
// DELETE /users/{id}
func (UnimplementedHandler) DeleteUser(ctx context.Context, params DeleteUserParams) (r DeleteUserRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListIntegration implements listIntegration operation.
//
// List Integrations.
//
// GET /integrations
func (UnimplementedHandler) ListIntegration(ctx context.Context, params ListIntegrationParams) (r ListIntegrationRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListMembership implements listMembership operation.
//
// List Memberships.
//
// GET /memberships
func (UnimplementedHandler) ListMembership(ctx context.Context, params ListMembershipParams) (r ListMembershipRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListOrganization implements listOrganization operation.
//
// List Organizations.
//
// GET /organizations
func (UnimplementedHandler) ListOrganization(ctx context.Context, params ListOrganizationParams) (r ListOrganizationRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListOrganizationIntegrations implements listOrganizationIntegrations operation.
//
// List attached Integrations.
//
// GET /organizations/{id}/integrations
func (UnimplementedHandler) ListOrganizationIntegrations(ctx context.Context, params ListOrganizationIntegrationsParams) (r ListOrganizationIntegrationsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListOrganizationMemberships implements listOrganizationMemberships operation.
//
// List attached Memberships.
//
// GET /organizations/{id}/memberships
func (UnimplementedHandler) ListOrganizationMemberships(ctx context.Context, params ListOrganizationMembershipsParams) (r ListOrganizationMembershipsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListSession implements listSession operation.
//
// List Sessions.
//
// GET /sessions
func (UnimplementedHandler) ListSession(ctx context.Context, params ListSessionParams) (r ListSessionRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListUser implements listUser operation.
//
// List Users.
//
// GET /users
func (UnimplementedHandler) ListUser(ctx context.Context, params ListUserParams) (r ListUserRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListUserMemberships implements listUserMemberships operation.
//
// List attached Memberships.
//
// GET /users/{id}/memberships
func (UnimplementedHandler) ListUserMemberships(ctx context.Context, params ListUserMembershipsParams) (r ListUserMembershipsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListUserSessions implements listUserSessions operation.
//
// List attached Sessions.
//
// GET /users/{id}/sessions
func (UnimplementedHandler) ListUserSessions(ctx context.Context, params ListUserSessionsParams) (r ListUserSessionsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadIntegration implements readIntegration operation.
//
// Finds the Integration with the requested ID and returns it.
//
// GET /integrations/{id}
func (UnimplementedHandler) ReadIntegration(ctx context.Context, params ReadIntegrationParams) (r ReadIntegrationRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadIntegrationOrganization implements readIntegrationOrganization operation.
//
// Find the attached Organization of the Integration with the given ID.
//
// GET /integrations/{id}/organization
func (UnimplementedHandler) ReadIntegrationOrganization(ctx context.Context, params ReadIntegrationOrganizationParams) (r ReadIntegrationOrganizationRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadMembership implements readMembership operation.
//
// Finds the Membership with the requested ID and returns it.
//
// GET /memberships/{id}
func (UnimplementedHandler) ReadMembership(ctx context.Context, params ReadMembershipParams) (r ReadMembershipRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadMembershipOrganization implements readMembershipOrganization operation.
//
// Find the attached Organization of the Membership with the given ID.
//
// GET /memberships/{id}/organization
func (UnimplementedHandler) ReadMembershipOrganization(ctx context.Context, params ReadMembershipOrganizationParams) (r ReadMembershipOrganizationRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadMembershipUser implements readMembershipUser operation.
//
// Find the attached User of the Membership with the given ID.
//
// GET /memberships/{id}/user
func (UnimplementedHandler) ReadMembershipUser(ctx context.Context, params ReadMembershipUserParams) (r ReadMembershipUserRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadOrganization implements readOrganization operation.
//
// Finds the Organization with the requested ID and returns it.
//
// GET /organizations/{id}
func (UnimplementedHandler) ReadOrganization(ctx context.Context, params ReadOrganizationParams) (r ReadOrganizationRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadSession implements readSession operation.
//
// Finds the Session with the requested ID and returns it.
//
// GET /sessions/{id}
func (UnimplementedHandler) ReadSession(ctx context.Context, params ReadSessionParams) (r ReadSessionRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadSessionUsers implements readSessionUsers operation.
//
// Find the attached User of the Session with the given ID.
//
// GET /sessions/{id}/users
func (UnimplementedHandler) ReadSessionUsers(ctx context.Context, params ReadSessionUsersParams) (r ReadSessionUsersRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadUser implements readUser operation.
//
// Finds the User with the requested ID and returns it.
//
// GET /users/{id}
func (UnimplementedHandler) ReadUser(ctx context.Context, params ReadUserParams) (r ReadUserRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateIntegration implements updateIntegration operation.
//
// Updates a Integration and persists changes to storage.
//
// PATCH /integrations/{id}
func (UnimplementedHandler) UpdateIntegration(ctx context.Context, req *UpdateIntegrationReq, params UpdateIntegrationParams) (r UpdateIntegrationRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateMembership implements updateMembership operation.
//
// Updates a Membership and persists changes to storage.
//
// PATCH /memberships/{id}
func (UnimplementedHandler) UpdateMembership(ctx context.Context, req *UpdateMembershipReq, params UpdateMembershipParams) (r UpdateMembershipRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateOrganization implements updateOrganization operation.
//
// Updates a Organization and persists changes to storage.
//
// PATCH /organizations/{id}
func (UnimplementedHandler) UpdateOrganization(ctx context.Context, req *UpdateOrganizationReq, params UpdateOrganizationParams) (r UpdateOrganizationRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateSession implements updateSession operation.
//
// Updates a Session and persists changes to storage.
//
// PATCH /sessions/{id}
func (UnimplementedHandler) UpdateSession(ctx context.Context, req *UpdateSessionReq, params UpdateSessionParams) (r UpdateSessionRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateUser implements updateUser operation.
//
// Updates a User and persists changes to storage.
//
// PATCH /users/{id}
func (UnimplementedHandler) UpdateUser(ctx context.Context, req *UpdateUserReq, params UpdateUserParams) (r UpdateUserRes, _ error) {
	return r, ht.ErrNotImplemented
}
