package responses

import "encoding/json"

// RemarshalJSON data into a more readable string
func RemarshalJSON(data []byte) (string, error) {
	var obj interface{}
	if err := json.Unmarshal(data, &obj); err != nil {
		return "", err
	}

	pretty, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return "", err
	}

	return string(pretty), nil
}
