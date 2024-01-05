package serveropts

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/alexedwards/scs/v2/memstore"
	echo "github.com/datumforge/echox"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/entdb"
	"github.com/datumforge/datum/internal/fga"
	"github.com/datumforge/datum/internal/graphapi"
	"github.com/datumforge/datum/internal/httpserve/config"
	"github.com/datumforge/datum/internal/httpserve/server"
	"github.com/datumforge/datum/internal/tokens"
	"github.com/datumforge/datum/internal/utils/marionette"
	"github.com/datumforge/datum/internal/utils/ulids"
)

type ServerOption interface {
	apply(*ServerOptions)
}

type applyFunc struct {
	applyInternal func(*ServerOptions)
}

func (fso *applyFunc) apply(s *ServerOptions) {
	fso.applyInternal(s)
}

func newApplyFunc(apply func(option *ServerOptions)) *applyFunc {
	return &applyFunc{
		applyInternal: apply,
	}
}

// WithConfigProvider supplies the config for the server
func WithConfigProvider(cfgProvider config.ConfigProvider) ServerOption {
	return newApplyFunc(func(s *ServerOptions) {
		s.ConfigProvider = cfgProvider
	})
}

// WithLogger supplies the logger for the server
func WithLogger(l *zap.SugaredLogger) ServerOption {
	return newApplyFunc(func(s *ServerOptions) {
		// Add logger to main config
		s.Config.Logger = l
		// Add logger to the handlers config
		s.Config.Server.Handler.Logger = l
	})
}

// WithServer supplies the echo server config for the server
func WithServer() ServerOption {
	return newApplyFunc(func(s *ServerOptions) {
		serverConfig := config.NewServerConfig()

		s.Config = *serverConfig
	})
}

// WithHTTPS sets up TLS config settings for the server
func WithHTTPS() ServerOption {
	return newApplyFunc(func(s *ServerOptions) {
		if !s.Config.Server.TLS.Enabled {
			// this is set to enabled by WithServer
			// if TLS is not enabled, move on
			return
		}

		s.Config.WithTLSDefaults()

		if !s.Config.Server.TLS.AutoCert {
			s.Config.WithTLSCerts(s.Config.Server.TLS.CertFile, s.Config.Server.TLS.CertKey)
		}
	})
}

// WithSQLiteDB supplies the sqlite db config for the server
func WithSQLiteDB() ServerOption {
	return newApplyFunc(func(s *ServerOptions) {
		// Database Config Setup
		dbConfig := &entdb.Config{}

		// load defaults and env vars
		err := envconfig.Process("datum_db", dbConfig)
		if err != nil {
			panic(err)
		}

		s.Config.DB = *dbConfig
	})
}

// WithFGAAuthz supplies the FGA authz config for the server
func WithFGAAuthz() ServerOption {
	return newApplyFunc(func(s *ServerOptions) {
		config, err := fga.NewAuthzConfig(s.Config.Logger)
		if err != nil {
			panic(err)
		}

		s.Config.Authz = *config

		if !s.Config.Authz.Enabled {
			return
		}
	})
}

// WithGeneratedKeys will generate a public/private key pair
// that can be used for jwt signing.
// This should only be used in a development environment
func WithGeneratedKeys() ServerOption {
	return newApplyFunc(func(s *ServerOptions) {
		privFileName := "private_key.pem"

		// generate a new private key if one doesn't exist
		if _, err := os.Stat(privFileName); err != nil {
			// Generate a new RSA private key with 2048 bits
			privateKey, err := rsa.GenerateKey(rand.Reader, 2048) //nolint:gomnd
			if err != nil {
				s.Config.Logger.Panicf("Error generating RSA private key:", err)
			}

			// Encode the private key to the PEM format
			privateKeyPEM := &pem.Block{
				Type:  "RSA PRIVATE KEY",
				Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
			}

			privateKeyFile, err := os.Create(privFileName)
			if err != nil {
				s.Config.Logger.Panicf("Error creating private key file:", err)
			}

			if err := pem.Encode(privateKeyFile, privateKeyPEM); err != nil {
				s.Config.Logger.Panicw("unable to encode pem on startup", "error", err.Error())
			}

			privateKeyFile.Close()
		}

		keys := map[string]string{}

		// check if kid was passed in
		kidPriv := s.Config.Server.Token.KID

		// if we didn't get a kid in the settings, assign one
		if kidPriv == "" {
			kidPriv = ulids.New().String()
		}

		keys[kidPriv] = fmt.Sprintf("%v", privFileName)

		s.Config.Server.Token.Keys = keys
	})
}

// WithAuth supplies the authn and jwt config for the server
func WithAuth() ServerOption {
	return newApplyFunc(func(s *ServerOptions) {
		// Token Config Setup
		tokenConfig := &tokens.Config{}

		// load defaults and env vars
		if err := envconfig.Process("datum_token", tokenConfig); err != nil {
			panic(err)
		}

		s.Config.Server.Token = *tokenConfig

		// Token Config Setup
		authConfig := &config.Auth{}

		// load defaults and env vars
		if err := envconfig.Process("datum_auth", authConfig); err != nil {
			panic(err)
		}

		s.Config.Auth = *authConfig

		// TODO: currently not used, this needs to be updated
		// to allow for an array to be provided in envconfig
		authProviderConfig := &config.AuthProvider{}

		// load defaults and env vars
		if err := envconfig.Process("datum_auth_provider", authProviderConfig); err != nil {
			panic(err)
		}

		s.Config.Auth.Providers = []config.AuthProvider{*authProviderConfig}
	})
}

// WithReadyChecks adds readiness checks to the server
func WithReadyChecks(c *entdb.EntClientConfig, f *fga.Client) ServerOption {
	return newApplyFunc(func(s *ServerOptions) {
		// Always add a check to the primary db connection
		s.Config.Server.Handler.AddReadinessCheck("sqlite_db_primary", entdb.Healthcheck(c.GetPrimaryDB()))

		// Check the secondary db, if enabled
		if s.Config.DB.MultiWrite {
			s.Config.Server.Handler.AddReadinessCheck("sqlite_db_secondary", entdb.Healthcheck(c.GetSecondaryDB()))
		}

		// Check the connection to openFGA, if enabled
		if s.Config.Authz.Enabled {
			s.Config.Server.Handler.AddReadinessCheck("fga", fga.Healthcheck(*f))
		}
	})
}

// WithGraphRoute adds the graph handler to the server
func WithGraphRoute(srv *server.Server, c *generated.Client, mw []echo.MiddlewareFunc) ServerOption {
	return newApplyFunc(func(s *ServerOptions) {
		// Setup Graph API Handlers
		r := graphapi.NewResolver(c, s.Config.Auth.Enabled).
			WithLogger(s.Config.Logger.Named("resolvers"))

		handler := r.Handler(s.Config.Server.Dev, mw...)

		// Add Graph Handler
		srv.AddHandler(handler)
	})
}

// WithMiddleware adds the middleware to the server
func WithMiddleware(mw []echo.MiddlewareFunc) ServerOption {
	return newApplyFunc(func(s *ServerOptions) {
		// Initialize middleware if null
		if s.Config.Server.Middleware == nil {
			s.Config.Server.Middleware = []echo.MiddlewareFunc{}
		}

		s.Config.Server.Middleware = append(s.Config.Server.Middleware, mw...)
	})
}

// WithEmailManager sets up the default SendGrid email manager to be used to send emails to users
// on registration, password reset, etc
func WithEmailManager() ServerOption {
	return newApplyFunc(func(s *ServerOptions) {
		if err := s.Config.Server.Handler.NewEmailManager(); err != nil {
			s.Config.Logger.Panicw("unable to create email manager", "error", err.Error())
		}
	})
}

// WithTaskManager sets up the default Marionette task manager to be used for delegating background tasks
func WithTaskManager() ServerOption {
	return newApplyFunc(func(s *ServerOptions) {
		// Start task manager
		tmConfig := marionette.Config{
			Logger: s.Config.Logger,
		}

		tm := marionette.New(tmConfig)

		tm.Start()

		s.Config.Server.Handler.TaskMan = tm
	})
}

// WithSessionManager sets up the default session manager with a 15 minute timeout and stale sessions are cleaned every 5 minutes
func WithSessionManager() ServerOption {
	return newApplyFunc(func(s *ServerOptions) {
		sm := scs.New()
		sm.Lifetime = time.Hour
		sm.Store = memstore.NewWithCleanupInterval(5 * time.Minute) // nolint: gomnd
		sm.IdleTimeout = 15 * time.Minute                           // nolint: gomnd
		sm.Cookie.Name = "__Host-datum"
		sm.Cookie.HttpOnly = true
		sm.Cookie.Persist = false
		sm.Cookie.SameSite = http.SameSiteStrictMode
		sm.Cookie.Secure = true
		s.Config.Server.Handler.SM = sm
	})
}
