package entfga

import (
	"text/template"

	"entgo.io/ent/entc/gen"
)

// extractUpdatedByKey gets the key that is used for the updated_by field
func extractObjectType(val any) string {
	objectType, ok := val.(string)
	if !ok {
		return ""
	}

	return objectType
}

// parseTemplate parses the template and sets values in the template
func parseTemplate(name, path string) *gen.Template {
	t := gen.NewTemplate(name)
	t.Funcs(template.FuncMap{
		"objectType": extractObjectType,
	})

	return gen.MustParse(t.ParseFS(_templates, path))
}
