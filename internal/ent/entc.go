//go:build ignore

// See Upstream docs for more details: https://entgo.io/docs/code-gen/#use-entc-as-a-package

package main

import (
	"log"
	"os"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/datumforge/fgax"
	"github.com/datumforge/fgax/entfga"
	"go.uber.org/zap"
	"gocloud.dev/secrets"

	"github.com/datumforge/enthistory"
	"github.com/datumforge/entx"
	"github.com/datumforge/entx/genhooks"
	geodetic "github.com/datumforge/geodetic/pkg/geodeticclient"

	"github.com/datumforge/datum/internal/ent/entconfig"
	"github.com/datumforge/datum/pkg/analytics"
	"github.com/datumforge/datum/pkg/sessions"
	"github.com/datumforge/datum/pkg/tokens"
	"github.com/datumforge/datum/pkg/utils/emails"
	"github.com/datumforge/datum/pkg/utils/marionette"
	"github.com/datumforge/datum/pkg/utils/totp"
)

const (
	graphSchemaDir = "./schema/"
	graphQueryDir  = "./query/"
	schemaPath     = "./internal/ent/schema"
)

func main() {
	if err := os.Mkdir("schema", 0755); err != nil && !os.IsExist(err) {
		log.Fatalf("creating schema directory: %v", err)
	}

	// generate the history schemas
	historyExt, entfgaExt := preRun()

	xExt, err := entx.NewExtension(entx.WithJSONScalar())
	if err != nil {
		log.Fatalf("creating entx extension: %v", err)
	}

	gqlExt, err := entgql.NewExtension(
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("schema/ent.graphql"),
		entgql.WithConfigPath("gqlgen.yml"),
		entgql.WithWhereInputs(true),
		entgql.WithSchemaHook(xExt.GQLSchemaHooks()...),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}

	if err := entc.Generate(schemaPath, &gen.Config{
		Target:    "./internal/ent/generated",
		Templates: entgql.AllTemplates,
		Hooks: []gen.Hook{
			genhooks.GenSchema(graphSchemaDir),
			genhooks.GenQuery(graphQueryDir),
		},
		Package: "github.com/datumforge/datum/internal/ent/generated",
		Features: []gen.Feature{
			gen.FeatureVersionedMigration,
			gen.FeaturePrivacy,
			gen.FeatureEntQL,
			gen.FeatureNamedEdges,
			gen.FeatureSchemaConfig,
			gen.FeatureIntercept,
		},
	},
		entc.Dependency(
			entc.DependencyName("EntConfig"),
			entc.DependencyType(&entconfig.Config{}),
		),
		entc.Dependency(
			entc.DependencyName("Secrets"),
			entc.DependencyType(&secrets.Keeper{}),
		),
		entc.Dependency(
			entc.DependencyName("Authz"),
			entc.DependencyType(fgax.Client{}),
		),
		entc.Dependency(
			entc.DependencyName("TokenManager"),
			entc.DependencyType(&tokens.TokenManager{}),
		),
		entc.Dependency(
			entc.DependencyName("SessionConfig"),
			entc.DependencyType(&sessions.SessionConfig{}),
		),
		entc.Dependency(
			entc.DependencyName("Logger"),
			entc.DependencyType(zap.SugaredLogger{}),
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
		entc.Dependency(
			entc.DependencyName("Geodetic"),
			entc.DependencyType(&geodetic.Client{}),
		),
		entc.TemplateDir("./internal/ent/templates"),
		entc.Extensions(
			gqlExt,
			historyExt,
			entfgaExt,
		)); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}

// preRun runs before the ent codegen to generate the history schemas and authz checks
// and returns the history and fga extensions to be used in the ent codegen
func preRun() (*enthistory.HistoryExtension, *entfga.AuthzExtension) {
	// generate the history schemas
	historyExt := enthistory.New(
		enthistory.WithAuditing(),
		enthistory.WithImmutableFields(),
		enthistory.WithHistoryTimeIndex(),
		enthistory.WithNillableFields(),
		enthistory.WithGQLQuery(),
		enthistory.WithAuthzPolicy(),
		enthistory.WithSchemaPath(schemaPath),
		enthistory.WithFirstRun(true),
		enthistory.WithAllowedRelation("audit_log_viewer"),
		enthistory.WithUpdatedByFromSchema(enthistory.ValueTypeString, false),
	)

	if err := historyExt.GenerateSchemas(); err != nil {
		log.Fatalf("generating history schema: %v", err)
	}

	// initialize the entfga extension
	entfgaExt := entfga.New(
		entfga.WithSoftDeletes(),
		entfga.WithSchemaPath(schemaPath),
	)

	// generate authz checks, this needs to happen before we regenerate the schema
	// because the authz checks are generated based on the schema
	if err := entfgaExt.GenerateAuthzChecks(); err != nil {
		log.Fatalf("generating authz checks: %v", err)
	}

	// run again with policy
	historyExt.SetFirstRun(false)

	// generate the updated history schemas with authz checks
	if err := historyExt.GenerateSchemas(); err != nil {
		log.Fatalf("generating history schema: %v", err)
	}

	return historyExt, entfgaExt
}
