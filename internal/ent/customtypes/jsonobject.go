package customtypes

import (
	"encoding/json"
	"io"
	"log"
)

type JSONObject map[string]interface{}

// MarshalGQL implement the Marshaler interface for gqlgen
func (j JSONObject) MarshalGQL(w io.Writer) {
	byteData, err := json.Marshal(j)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	_, err = w.Write(byteData)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

// UnmarshalGQL implement the Unmarshaler interface for gqlgen
func (j *JSONObject) UnmarshalGQL(v interface{}) error {
	byteData, err := json.Marshal(v)
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteData, &j)
	if err != nil {
		return err
	}

	return err
}

// type Pair struct {
//	K, V []byte
//}
//
//// Value implements the driver Valuer interface
// func (p Pair) Value() (driver.Value, error) {
//	var b bytes.Buffer
//	if err := gob.NewEncoder(&b).Encode(p); err != nil {
//		return nil, err
//	}
//
//	return b.Bytes(), nil
//}
//
//// Scan implements the Scanner interface.
// func (p *Pair) Scan(value interface{}) (err error) {
//	switch v := value.(type) {
//	case nil:
//	case []byte:
//		err = gob.NewDecoder(bytes.NewBuffer(v)).Decode(p)
//	default:
//		err = fmt.Errorf("unexpected type %T", v)
//	}
//
//	return
//}
