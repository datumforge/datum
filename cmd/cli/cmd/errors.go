package datum

import "errors"

var (
	// ErrOrgNameRequired is returned when no organization name is provided when creating a new organization
	ErrOrgNameRequired = errors.New("organization name is required")

	// ErrOrgIDRequired is returned when no organization id is provided when deleting or updating an organization
	ErrOrgIDRequired = errors.New("organization id is required")

	// ErrUserEmailRequired is returned when no user email is provided when creating a new user
	ErrUserEmailRequired = errors.New("email is required")

	// ErrUserFirstNameRequired is returned when no user first name is provided when creating a new user
	ErrUserFirstNameRequired = errors.New("first name is required")

	// ErrUserLastNameRequired is returned when no user last name is provided when creating a new user
	ErrUserLastNameRequired = errors.New("last name is required")

	// ErrUserIDRequired is returned when no user id is provided when deleting or updating an existing user
	ErrUserIDRequired = errors.New("user id is required")

	// ErrTokenRequired is returned when no authentication token is provided
	ErrTokenRequired = errors.New("DATUM_ACCESS_TOKEN not set")
)
