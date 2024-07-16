package resolvergen

import (
	"embed"

	"bytes"
	"html/template"
	"strings"

	"github.com/99designs/gqlgen/codegen"
	gqltemplates "github.com/99designs/gqlgen/codegen/templates"

	"github.com/stoewer/go-strcase"
	gqlast "github.com/vektah/gqlparser/v2/ast"
)

//go:embed templates/**.gotpl
var templates embed.FS

// crudResolver is a struct to hold the field for the CRUD resolver
type crudResolver struct {
	Field *codegen.Field
}

// renderTemplate renders the template with the given name
func renderTemplate(templateName string, field *codegen.Field) string {
	t, err := template.New(templateName).Funcs(template.FuncMap{
		"getEntityName": getEntityName,
		"toLower":       strings.ToLower,
		"toLowerCamel":  strcase.LowerCamelCase,
		"hasArgument":   hasArgument,
		"hasOwnerField": hasOwnerField,
		"reserveImport": gqltemplates.CurrentImports.Reserve,
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

// renderCreate renders the create template
func renderCreate(field *codegen.Field) string {
	return renderTemplate("create.gotpl", field)
}

// renderUpdate renders the update template
func renderUpdate(field *codegen.Field) string {
	return renderTemplate("update.gotpl", field)
}

// renderDelete renders the delete template
func renderDelete(field *codegen.Field) string {
	return renderTemplate("delete.gotpl", field)
}

// renderBulkUpload renders the bulk upload template
func renderBulkUpload(field *codegen.Field) string {
	return renderTemplate("upload.gotpl", field)
}

// renderBulk renders the bulk template
func renderBulk(field *codegen.Field) string {
	return renderTemplate("bulk.gotpl", field)
}

// renderQuery renders the query template
func renderQuery(field *codegen.Field) string {
	return renderTemplate("get.gotpl", field)
}

// renderList renders the list template
func renderList(field *codegen.Field) string {
	return renderTemplate("list.gotpl", field)
}

// crudTypes is a list of CRUD operations that are included in the resolver name
var stripStrings = []string{"Create", "Update", "Delete", "Bulk", "CSV", "Connection", "Payload"}

// getEntityName returns the entity name by stripping the CRUD operation from the resolver name
func getEntityName(name string) string {
	for _, s := range stripStrings {
		if strings.Contains(name, s) {
			name = strings.ReplaceAll(name, s, "")
		}
	}

	return name
}

// hasArgument checks if the argument is present in the list of arguments
func hasArgument(arg string, args gqlast.ArgumentDefinitionList) bool {
	for _, a := range args {
		if a.Name == arg {
			return true
		}
	}

	return false
}

// hasOwnerField checks if the field has an owner field in the input arguments
func hasOwnerField(field *codegen.Field) bool {
	for _, arg := range field.Args {
		if arg.TypeReference.Definition.Kind == gqlast.InputObject {
			if arg.TypeReference.Definition.Fields.ForName("ownerID") != nil {
				return true
			}
		}
	}

	return false
}
