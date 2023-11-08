package entx

import (
	"encoding/json"
	"io"

	"github.com/99designs/gqlgen/graphql"
)

// MarshalRawMessage provides a graphql.Marshaler for json.RawMessage
func MarshalRawMessage(t json.RawMessage) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		s, _ := t.MarshalJSON()
		_, _ = io.WriteString(w, string(s))
	})
}

// UnmarshalRawMessage provides a graphql.Unmarshaler for json.RawMessage
func UnmarshalRawMessage(v interface{}) (json.RawMessage, error) {
	switch j := v.(type) {
	case []byte:
		return json.RawMessage(j), nil
	case map[string]interface{}:
		js, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}

		return json.RawMessage(js), nil
	default:
		js, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}

		return json.RawMessage(js), nil
	}
}
