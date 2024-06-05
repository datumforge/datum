package datum

import (
	"encoding/json"
	"fmt"

	"github.com/TylerBrock/colorjson"
)

func JSONPrint(s []byte) error {
	var obj map[string]interface{}

	err := json.Unmarshal(s, &obj)
	if err != nil {
		return err
	}

	f := colorjson.NewFormatter()
	f.Indent = 2

	o, err := f.Marshal(obj)
	if err != nil {
		return err
	}

	fmt.Println(string(o))

	return nil
}

// ParseJSON parses a JSON formatted string into a map
func ParseJSON(v string) (map[string]interface{}, error) {
	var m map[string]interface{}

	if err := json.Unmarshal([]byte(v), &m); err != nil {
		return nil, err
	}

	return m, nil
}

// ParseBytes parses buffered bytes into a map
func ParseBytes(v []byte) (map[string]interface{}, error) {
	var m map[string]interface{}

	if err := json.Unmarshal(v, &m); err != nil {
		return nil, err
	}

	return m, nil
}
