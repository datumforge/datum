package nanox

import (
	"database/sql"
	"database/sql/driver"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/jaevor/go-nanoid"
)

type ID string

func GetNewID() (ID, error) {
	canonicID, err := nanoid.Standard(21)
	if err != nil {
		return "", err
	}

	return ID(canonicID()), nil
}

func MustGetNewID() ID {
	v, err := GetNewID()
	if err != nil {
		panic(err)
	}
	return v
}

// MarshalGQL implements the graphql.Marshaler interface
func MarshalID(u ID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, strconv.Quote(string(u)))
	})
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (u *ID) UnmarshalGQL(v interface{}) error {
	return u.Scan(v)
}

func (p *ID) Scan(v any) error {
	if v == nil {
		*p = ID("")
		return nil
	}

	switch src := v.(type) {
	case string:
		*p = ID(src)
	case []byte:
		*p = ID(string(src))
	case ID:
		*p = src
	default:
		return ErrUnsupportedType
	}

	return nil
}

// Value implements the driver Valuer interface.
func (u ID) Value() (driver.Value, error) {
	return string(u), nil
}

func (u ID) IsValid() bool {
	return u != ""
}

var (
	_ driver.Valuer = ID("")
	_ sql.Scanner   = (*ID)(nil)
)
