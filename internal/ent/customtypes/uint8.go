package customtypes

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
)

type Uint8 uint8

func (u Uint8) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.FormatUint(uint64(u), 10))
}

func (u *Uint8) UnmarshalGQL(v interface{}) error {
	i, err := graphql.UnmarshalUint64(v)
	if err != nil {
		return err
	}
	*u = Uint8(i)
	return nil
}
