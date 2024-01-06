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
	"github.com/ogen-go/ogen"
	"github.com/stoewer/go-strcase"
	"go.uber.org/zap"
	"gocloud.dev/secrets"

	"github.com/datumforge/datum/internal/entx"
	"github.com/datumforge/datum/internal/fga"
)

var (
	entSchemaDir   = "./internal/ent/schema/"
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

	// Add OpenAPI Gen extension
	spec := new(ogen.Spec)
	oas, err := entoas.NewExtension(
		entoas.Spec(spec),
		entoas.Mutations(func(graph *gen.Graph, spec *ogen.Spec) error {
			spec.Info.SetTitle("Datum API").
				SetDescription("Programmatic interfaces for interacting with Datum Services").
				SetVersion("0.0.1")
			// TODO: finish the remainder of our http endpoints
			spec.AddPathItem("/livez", ogen.NewPathItem().
				SetDescription("Check if the server is up").
				SetGet(ogen.NewOperation().
					SetOperationID("Livez").
					SetSummary("Simple endpoint to check if the server is up").
					AddResponse(
						"200",
						ogen.
							NewResponse().
							SetDescription("Server is reachable").
							SetJSONContent(
								ogen.NewSchema().
									SetType("object").
									AddRequiredProperties(
										ogen.String().ToProperty("status"),
									),
							),
					).
					AddResponse("503", ogen.NewResponse().SetDescription("Server is not reachable")),
				),
			)

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
			entc.DependencyType(fga.Client{}),
		),
		entc.Dependency(
			entc.DependencyName("Logger"),
			entc.DependencyType(zap.SugaredLogger{}),
		),
		entc.Dependency(
			entc.DependencyType(&http.Client{}),
		),
		entc.TemplateDir("./internal/ent/templates"),
		entc.Extensions(
			gqlExt,
			oas,
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
