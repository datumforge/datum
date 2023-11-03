// Package nanox provides a ID interface based on go-nanoid
package nanox

import (
	"database/sql"
	"database/sql/driver"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/jaevor/go-nanoid"
)

const (
	idLength = 21
)

// ID is a string based on the go-nanoid implementation
type ID string

// GetNewID returns an ID based on go-nanoid
func GetNewID() (ID, error) {
	canonicID, err := nanoid.Standard(idLength)
	if err != nil {
		return "", err
	}

	return ID(canonicID()), nil
}

// MustGetNewID returns an ID
func MustGetNewID() ID {
	v, err := GetNewID()
	if err != nil {
		panic(err)
	}

	return v
}

// MarshalID implements the graphql.Marshaler interface
func MarshalID(u ID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, strconv.Quote(string(u)))
	})
}

// UnmarshalID implements the graphql.Unmarshaler interface
func (u *ID) UnmarshalID(v interface{}) error {
	return u.Scan(v)
}

// Scan checks the type of the provided ID
func (u *ID) Scan(v any) error {
	if v == nil {
		*u = ID("")
		return nil
	}

	switch src := v.(type) {
	case string:
		*u = ID(src)
	case []byte:
		*u = ID(string(src))
	case ID:
		*u = src
	default:
		return ErrUnsupportedType
	}

	return nil
}

// Value implements the driver Valuer interface.
func (u ID) Value() (driver.Value, error) {
	return string(u), nil
}

// IsValid checks if the ID provided is not empty
func (u ID) IsValid() bool {
	return u != ""
}

var (
	_ driver.Valuer = ID("")
	_ sql.Scanner   = (*ID)(nil)
)
