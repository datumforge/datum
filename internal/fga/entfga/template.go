package entfga

import (
	"strings"
	"text/template"

	"entgo.io/ent/entc/gen"
	"github.com/stoewer/go-strcase"
)

// extractObjectType gets the key that is used for the object type
func extractObjectType(val any) string {
	objectType, ok := val.(string)
	if !ok {
		return ""
	}

	return objectType
}

// useSoftDeletes checks the config properties for the Soft Delete setting
func useSoftDeletes(config Config) bool {
	return config.SoftDeletes
}

// parseTemplate parses the template and sets values in the template
func parseTemplate(name, path string) *gen.Template {
	t := gen.NewTemplate(name)
	t.Funcs(template.FuncMap{
		"extractObjectType": extractObjectType,
		"useSoftDeletes":    useSoftDeletes,
		"ToUpperCamel":      strcase.UpperCamelCase,
		"ToLower":           strings.ToLower,
	})

	return gen.MustParse(t.ParseFS(_templates, path))
}
