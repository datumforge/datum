//go:build ignore

// See Upstream docs for more details: https://entgo.io/docs/code-gen/#use-entc-as-a-package

package main

import (
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"text/template"

	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entoas"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/ogen-go/ogen"
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
		Package:   "github.com/datumforge/datum/internal/ent/generated",
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

	check, _ := strconv.ParseBool(os.Getenv("SCHEMAGEN"))
	if check {
		generateSchemaFuncs()
	}
}

func generateSchemaFuncs() {
	r, err := regexp.Compile("type ([a-zA-Z].*) struct")
	if err != nil {
		log.Fatalf("Unable to compile regex: %v", err)
	}

	skipFile, err := os.ReadFile("./scripts/files_to_skip.txt")
	if err != nil {
		log.Fatalf("Unable to read skip list: %v", err)
	}

	files, _ := os.ReadDir(entSchemaDir)
	for _, f := range files {
		skip, err := regexp.Match(f.Name(), skipFile)
		if err != nil {
			log.Fatalf("Unable to search file: %v", err)
		}

		if !skip {
			if _, err := os.Stat(graphSchemaDir + strings.Split(f.Name(), ".")[0] + ".graphql"); os.IsNotExist(err) {
				log.Printf("Generating schema %s\n", f.Name())

				t, err := os.ReadFile(entSchemaDir + f.Name())
				if err != nil {
					log.Fatalf("Unable to read file: %v", err)
				}

				m := strings.Split(string(r.Find(t)), "")[1]

				fm := template.FuncMap{
					"ToLower": strings.ToLower,
				}

				tmpl, err := template.New("graph.tpl").Funcs(fm).ParseFiles("./scripts/templates/graph.tpl")
				if err != nil {
					log.Fatalf("Unable to parse template: %v", err)
				}

				file, err := os.Create(graphSchemaDir + strings.ToLower(m) + ".graphql")
				if err != nil {
					log.Fatalf("Unable to create file: %v", err)
				}

				s := struct {
					Name string
				}{
					Name: m,
				}

				err = tmpl.Execute(file, s)
				if err != nil {
					log.Fatalf("Unable to execute template: %v", err)
				}
			} else {
				log.Printf("Schema exists %s\n", f.Name())
			}
		}
	}
}
