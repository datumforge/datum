package entfga

import (
	"database/sql/driver"
	"io"
	"strconv"
)

// OpType is the ent operation type in string form
type OpType string

const (
	// OpTypeInsert is the insert (create) operation
	OpTypeInsert OpType = "INSERT"
	// OpTypeUpdate is the update operation
	OpTypeUpdate OpType = "UPDATE"
	// OpTypeDelete is the delete operation
	OpTypeDelete OpType = "DELETE"
)

// opTypes are the possible values that can be used
var opTypes = []string{
	OpTypeInsert.String(),
	OpTypeUpdate.String(),
	OpTypeDelete.String(),
}

// Values provides list valid values for Enum.
func (OpType) Values() (kinds []string) {
	kinds = append(kinds, opTypes...)
	return
}

// Value of the operation type
func (op OpType) Value() (driver.Value, error) {
	return op.String(), nil
}

// String value of the operation
func (op OpType) String() string {
	return string(op)
}

// Scan implements the `database/sql.Scanner` interface for the `OpType` type
// and is used to convert a value from the database into an `OpType` value.
func (op *OpType) Scan(v any) error {
	if v == nil {
		*op = OpType("")
		return nil
	}

	switch src := v.(type) {
	case string:
		*op = OpType(src)
	case []byte:
		*op = OpType(string(src))
	case OpType:
		*op = src
	default:
		return ErrUnsupportedType
	}

	return nil
}

// MarshalGQL implement the Marshaler interface for gqlgen
func (op OpType) MarshalGQL(w io.Writer) {
	// graphql ID is a scalar which must be quoted
	io.WriteString(w, strconv.Quote(string(op))) //nolint:errcheck
}

// UnmarshalGQL implement the Unmarshaler interface for gqlgen
func (op *OpType) UnmarshalGQL(v interface{}) error {
	return op.Scan(v)
}
