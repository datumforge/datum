//go:build ignore

// See Upstream docs for more details: https://entgo.io/docs/code-gen/#use-entc-as-a-package

package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entoas"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/datumforge/fgax"
	"github.com/datumforge/fgax/entfga"
	"github.com/ogen-go/ogen"
	"github.com/stoewer/go-strcase"
	"go.uber.org/zap"
	"gocloud.dev/secrets"

	"github.com/datumforge/entx"

	"github.com/datumforge/datum/internal/utils/emails"
	"github.com/datumforge/datum/internal/utils/marionette"
	"github.com/datumforge/datum/pkg/analytics"
)

var (
	graphSchemaDir = "./schema/"
)

func main() {
	xExt, err := entx.NewExtension(
		entx.WithJSONScalar(),
	)
	if err != nil {
		log.Fatalf("creating entx extension: %v", err)
	}

	// Ensure the schema directory exists before running entc.
	_ = os.Mkdir("schema", 0755)

	ex, err := entoas.NewExtension(
		entoas.SimpleModels(),
		entoas.Mutations(func(graph *gen.Graph, spec *ogen.Spec) error {
			spec.SetOpenAPI("3.1.0")
			spec.SetServers([]ogen.Server{
				{
					URL:         "https://api.datum.net",
					Description: "Datum Production API Endpoint",
				},
				{
					URL:         "http://localhost:17608/v1",
					Description: "http localhost endpoint for testing purposes",
				}})
			spec.Info.SetTitle("Datum OpenAPI 3.1.0 Specifications").
				SetDescription("Programmatic interfaces for interacting with Datum Services").
				SetVersion("1.0.1")
			spec.Info.SetContact(&ogen.Contact{
				Name:  "Datum Support",
				URL:   "https://datum.net",
				Email: "support@datum.net",
			})
			spec.Info.SetLicense(&ogen.License{
				Name: "Apache 2.0",
				URL:  "https://www.apache.org/licenses/LICENSE-2.0",
			})
			spec.Info.SetTermsOfService("https://datum.net/tos")

			return nil
		}),
	)

	if err != nil {
		log.Fatalf("creating entoas extension: %v", err)
	}

	gqlExt, err := entgql.NewExtension(
		// Tell Ent to generate a GraphQL schema for
		// the Ent schema in a file named ent.graphql.
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("schema/ent.graphql"),
		entgql.WithConfigPath("gqlgen.yml"),
		entgql.WithWhereInputs(true),
		entgql.WithSchemaHook(xExt.GQLSchemaHooks()...),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}

	if err := entc.Generate("./internal/ent/schema", &gen.Config{
		Target:    "./internal/ent/generated",
		Templates: entgql.AllTemplates,
		Hooks: []gen.Hook{
			GenSchema(),
		},
		Package: "github.com/datumforge/datum/internal/ent/generated",
		Features: []gen.Feature{
			gen.FeatureVersionedMigration,
			gen.FeaturePrivacy,
			gen.FeatureSnapshot,
			gen.FeatureEntQL,
			gen.FeatureNamedEdges,
			gen.FeatureSchemaConfig,
			gen.FeatureIntercept,
		},
	},
		entc.Dependency(
			entc.DependencyType(&secrets.Keeper{}),
		),
		entc.Dependency(
			entc.DependencyName("Authz"),
			entc.DependencyType(fgax.Client{}),
		),
		entc.Dependency(
			entc.DependencyName("Logger"),
			entc.DependencyType(zap.SugaredLogger{}),
		),
		entc.Dependency(
			entc.DependencyType(&http.Client{}),
		),
		entc.Dependency(
			entc.DependencyName("Emails"),
			entc.DependencyType(&emails.EmailManager{}),
		),
		entc.Dependency(
			entc.DependencyName("Marionette"),
			entc.DependencyType(&marionette.TaskManager{}),
		),
		entc.Dependency(
			entc.DependencyName("Analytics"),
			entc.DependencyType(&analytics.EventManager{}),
		),
		entc.TemplateDir("./internal/ent/templates"),
		entc.Extensions(
			gqlExt,
			ex,
			entfga.NewFGAExtension(
				entfga.WithSoftDeletes(),
			),
		)); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}

// GenSchema generates graphql schemas when not specified to be skipped
func GenSchema() gen.Hook {
	return func(next gen.Generator) gen.Generator {
		return gen.GenerateFunc(func(g *gen.Graph) error {
			for _, node := range g.Nodes {
				if sg, ok := node.Annotations[entx.SchemaGenAnnotationName]; ok {
					val, _ := sg.(map[string]interface{})["Skip"]

					if val.(bool) {
						continue
					}
				}

				fm := template.FuncMap{
					"ToLowerCamel": strcase.LowerCamelCase,
				}

				tmpl, err := template.New("graph.tpl").Funcs(fm).ParseFiles("./scripts/templates/graph.tpl")
				if err != nil {
					log.Fatalf("Unable to parse template: %v", err)
				}

				file, err := os.Create(graphSchemaDir + strings.ToLower(node.Name) + ".graphql")
				if err != nil {
					log.Fatalf("Unable to create file: %v", err)
				}

				s := struct {
					Name string
				}{
					Name: node.Name,
				}

				err = tmpl.Execute(file, s)
				if err != nil {
					log.Fatalf("Unable to execute template: %v", err)
				}
			}
			return next.Generate(g)
		})
	}
}
