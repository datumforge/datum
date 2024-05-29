package resolvergen

import (
	"embed"
	_ "embed"

	"bytes"
	"html/template"
	"strings"

	"github.com/99designs/gqlgen/codegen"
	"github.com/stoewer/go-strcase"
)

//go:embed templates/**.gotpl
var templates embed.FS

type crudResolver struct {
	Field *codegen.Field
}

func renderTemplate(templateName string, field *codegen.Field) string {
	t, err := template.New(templateName).Funcs(template.FuncMap{
		"getEntityName": getEntityName,
		"toLower":       strings.ToLower,
		"toLowerCamel":  strcase.LowerCamelCase,
	}).ParseFS(templates, "templates/"+templateName)
	if err != nil {
		panic(err)
	}

	var code bytes.Buffer

	if err = t.Execute(&code, &crudResolver{
		Field: field,
	}); err != nil {
		panic(err)
	}

	return strings.Trim(code.String(), "\t \n")
}

func renderCreate(field *codegen.Field) string {
	return renderTemplate("create.gotpl", field)
}

func renderUpdate(field *codegen.Field) string {
	return renderTemplate("update.gotpl", field)
}

func renderDelete(field *codegen.Field) string {
	return renderTemplate("delete.gotpl", field)
}

func renderBulkUpload(field *codegen.Field) string {
	return renderTemplate("upload.gotpl", field)
}

// crudTypes is a list of CRUD operations that are included in the resolver name
var crudTypes = []string{"Create", "Update", "Delete", "Bulk", "CSV"}

func getEntityName(name string) string {
	for _, crudType := range crudTypes {
		if strings.Contains(name, crudType) {
			name = strings.ReplaceAll(name, crudType, "")
		}
	}

	return name
}
