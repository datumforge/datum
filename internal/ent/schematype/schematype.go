package schematype

import (
	"database/sql/driver"
	"encoding/json"
)

// TemplateConfig implements the field.ValueScanner interface
type TemplateConfig struct {
	Testfield string `json:"testfield"`
}

func (t *TemplateConfig) Scan(v interface{}) (err error) {
	switch v := v.(type) {
	case string:
		err = json.Unmarshal([]byte(v), t)
	case []byte:
		err = json.Unmarshal(v, t)
	}

	return
}

func (t *TemplateConfig) Value() (driver.Value, error) {
	return json.Marshal(t)
}
