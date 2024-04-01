//go:build ignore

// See Upstream docs for more details: https://entgo.io/docs/code-gen/#use-entc-as-a-package

package main

import (
	"log"
	"net/http"
	"os"

	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entoas"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/datumforge/fgax"
	"github.com/datumforge/fgax/entfga"
	"github.com/ogen-go/ogen"
	"go.uber.org/zap"
	"gocloud.dev/secrets"

	"github.com/datumforge/entx"

	"github.com/datumforge/datum/pkg/analytics"
	"github.com/datumforge/datum/pkg/utils/emails"
	"github.com/datumforge/datum/pkg/utils/marionette"
	"github.com/datumforge/datum/pkg/utils/totp"
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
					URL:         "https://api.datum.net/v1",
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
				URL:   "https://datum.net/support",
				Email: "support@datum.net",
			})
			spec.Info.SetLicense(&ogen.License{
				Name: "Apache 2.0",
				URL:  "https://www.apache.org/licenses/LICENSE-2.0",
			})
			spec.Info.SetTermsOfService("https://datum.net/tos")

			bearerSecurity := ogen.SecurityScheme{
				Type:         "http",
				Scheme:       "bearer",
				Description:  "Bearer Token Authentication",
				BearerFormat: "JWT",
			}

			oauth2Security := ogen.SecurityScheme{
				Type: "oauth2",
				Flows: &ogen.OAuthFlows{
					Implicit: &ogen.OAuthFlow{
						AuthorizationURL: "https://api.datum.net/oauth2/authorize",
						Scopes: map[string]string{
							"email":   "email",
							"profile": "profile",
						},
					},
				},
			}

			oidcSecurity := ogen.SecurityScheme{
				Type:             "openIdConnect",
				OpenIDConnectURL: "https://api.datum.net/.well-known/openid-configuration",
			}

			spec.Components.SecuritySchemes = map[string]*ogen.SecurityScheme{
				"BearerAuth": &bearerSecurity,
				"OAuth2":     &oauth2Security,
				"OIDC":       &oidcSecurity,
			}

			security := ogen.SecurityRequirement{
				"OAuth2":     {"email", "profile"},
				"BearerAuth": {},
				"OIDC":       {},
			}

			spec.Security = append(spec.Security, security)

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
			entx.GenSchema(graphSchemaDir),
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
		entc.Dependency(
			entc.DependencyName("TOTP"),
			entc.DependencyType(&totp.Manager{}),
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
