package hooks

import (
	"errors"
)

var (
	// ErrInternalServerError is returned when an internal error occurs.
	ErrInternalServerError = errors.New("internal server error")

	// ErrPersonalOrgsNoChildren is returned when personal org attempts to add a child org
	ErrPersonalOrgsNoChildren = errors.New("personal organizations are not allowed to have child organizations")

	// ErrPersonalOrgsNoMembers is returned when personal org attempts to add members
	ErrPersonalOrgsNoMembers = errors.New("personal organizations are not allowed to have members other than the owner")

	// ErrPersonalOrgsNoUser is returned when personal org has no user associated, so no permissions can be added
	ErrPersonalOrgsNoUser = errors.New("personal organizations missing user association")

	// ErrUnsupportedFGARole is returned when a role is assigned that is not supported in our fine grained authorization system
	ErrUnsupportedFGARole = errors.New("unsupported role")

	// ErrMissingRole is returned when an update request is made that contains no role
	ErrMissingRole = errors.New("missing role in update")
)
