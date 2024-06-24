//go:build ignore

// See Upstream docs for more details: https://entgo.io/docs/code-gen/#use-entc-as-a-package

package main

import (
	"log"
	"net/http"
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
	geodetic "github.com/datumforge/geodetic/pkg/geodeticclient"

	"github.com/datumforge/datum/pkg/analytics"
	"github.com/datumforge/datum/pkg/sessions"
	"github.com/datumforge/datum/pkg/tokens"
	"github.com/datumforge/datum/pkg/utils/emails"
	"github.com/datumforge/datum/pkg/utils/marionette"
	"github.com/datumforge/datum/pkg/utils/totp"
)

var (
	graphSchemaDir = "./schema/"
)

func main() {
	xExt, err := entx.NewExtension(entx.WithJSONScalar())
	if err != nil {
		log.Fatalf("creating entx extension: %v", err)
	}

	_ = os.Mkdir("schema", 0755)

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

	historyExt := enthistory.NewHistoryExtension(
		enthistory.WithAuditing(),
		enthistory.WithImmutableFields(),
		enthistory.WithHistoryTimeIndex(),
		enthistory.WithNillableFields(),
		enthistory.WithGQLQuery(),
		enthistory.WithSchemaPath("./internal/ent/schema"),
	)
	if err != nil {
		log.Fatalf("creating enthistory extension: %v", err)
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
		entc.Dependency(
			entc.DependencyName("Geodetic"),
			entc.DependencyType(&geodetic.Client{}),
		),
		entc.TemplateDir("./internal/ent/templates"),
		entc.Extensions(
			gqlExt,
			entfga.NewFGAExtension(
				entfga.WithSoftDeletes(),
			),
			historyExt,
		)); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
