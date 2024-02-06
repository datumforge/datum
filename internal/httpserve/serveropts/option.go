package serveropts

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	echoprometheus "github.com/datumforge/echo-prometheus/v5"
	echo "github.com/datumforge/echox"
	"github.com/datumforge/echox/middleware"
	"github.com/datumforge/echozap"
	"github.com/datumforge/fgax"
	"github.com/kelseyhightower/envconfig"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	"github.com/datumforge/datum/internal/cache"
	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/entdb"
	"github.com/datumforge/datum/internal/graphapi"
	"github.com/datumforge/datum/internal/httpserve/config"
	"github.com/datumforge/datum/internal/httpserve/handlers"
	authmw "github.com/datumforge/datum/internal/httpserve/middleware/auth"
	"github.com/datumforge/datum/internal/httpserve/middleware/cachecontrol"
	"github.com/datumforge/datum/internal/httpserve/middleware/cors"
	"github.com/datumforge/datum/internal/httpserve/middleware/echocontext"
	"github.com/datumforge/datum/internal/httpserve/middleware/mime"
	"github.com/datumforge/datum/internal/httpserve/middleware/ratelimit"
	"github.com/datumforge/datum/internal/httpserve/server"
	"github.com/datumforge/datum/internal/otelx"
	"github.com/datumforge/datum/internal/providers/github"
	"github.com/datumforge/datum/internal/providers/google"
	"github.com/datumforge/datum/internal/sessions"
	"github.com/datumforge/datum/internal/tokens"
	"github.com/datumforge/datum/internal/utils/emails"
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

func WithTracer() ServerOption {
	return newApplyFunc(func(s *ServerOptions) {
		// Tracer Config Setup
		tracerConfig := &otelx.Config{}

		// load defaults and env vars
		if err := envconfig.Process("datum_tracing", tracerConfig); err != nil {
			panic(err)
		}

		s.Config.Tracer = *tracerConfig
	})
}

// WithFGAAuthz supplies the FGA authz config for the server
func WithFGAAuthz() ServerOption {
	return newApplyFunc(func(s *ServerOptions) {
		config, err := fgax.NewAuthzConfig(s.Config.Logger)
		if err != nil {
			panic(err)
		}

		s.Config.Authz = *config
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

		authProviderConfig := &handlers.OauthProviderConfig{}
		googleProvider := &handlers.GoogleConfig{}
		githubProvider := &handlers.GithubConfig{}

		// load defaults and env vars for GitHub provider
		if err := envconfig.Process("datum_auth_provider_github", githubProvider); err != nil {
			panic(err)
		}

		// load defaults and env vars for Google Provider
		if err := envconfig.Process("datum_auth_provider_google", googleProvider); err != nil {
			panic(err)
		}

		// load defaults and env vars for Oauth setup
		if err := envconfig.Process("datum_auth_provider", authProviderConfig); err != nil {
			panic(err)
		}

		// add supported providers if not set
		if len(s.Config.Auth.SupportedProviders) == 0 {
			s.Config.Auth.SupportedProviders = []string{github.ProviderName, google.ProviderName}
		}

		// add external providers
		authProviderConfig.GithubConfig = *githubProvider
		authProviderConfig.GoogleConfig = *googleProvider

		// add our oauth2 provider
		s.Config.Server.Handler.OauthProvider = *authProviderConfig

		// add auth middleware
		conf := authmw.NewAuthOptions(
			authmw.WithAudience(s.Config.Server.Token.Audience),
			authmw.WithIssuer(s.Config.Server.Token.Issuer),
			authmw.WithJWKSEndpoint(s.Config.Server.Token.JWKSEndpoint),
		)

		s.Config.Server.GraphMiddleware = append(s.Config.Server.GraphMiddleware, authmw.Authenticate(conf))
	})
}

// WithReadyChecks adds readiness checks to the server
func WithReadyChecks(c *entdb.EntClientConfig, f *fgax.Client, r *redis.Client) ServerOption {
	return newApplyFunc(func(s *ServerOptions) {
		// Always add a check to the primary db connection
		s.Config.Server.Handler.AddReadinessCheck("sqlite_db_primary", entdb.Healthcheck(c.GetPrimaryDB()))

		// Check the secondary db, if enabled
		if s.Config.DB.MultiWrite {
			s.Config.Server.Handler.AddReadinessCheck("sqlite_db_secondary", entdb.Healthcheck(c.GetSecondaryDB()))
		}

		// Check the connection to openFGA, if enabled
		if s.Config.Authz.Enabled {
			s.Config.Server.Handler.AddReadinessCheck("fga", fgax.Healthcheck(*f))
		}

		// Check the connection to redis, if enabled
		if s.Config.RedisConfig.Enabled {
			s.Config.Server.Handler.AddReadinessCheck("redis", cache.Healthcheck(r))
		}
	})
}

// WithGraphRoute adds the graph handler to the server
func WithGraphRoute(srv *server.Server, c *generated.Client) ServerOption {
	return newApplyFunc(func(s *ServerOptions) {
		// Setup Graph API Handlers
		r := graphapi.NewResolver(c).
			WithLogger(s.Config.Logger.Named("resolvers"))

		handler := r.Handler(s.Config.Server.Dev)

		// Add Graph Handler
		srv.AddHandler(handler)
	})
}

// WithMiddleware adds the middleware to the server
func WithMiddleware() ServerOption {
	return newApplyFunc(func(s *ServerOptions) {
		// Initialize middleware if null
		if s.Config.Server.DefaultMiddleware == nil {
			s.Config.Server.DefaultMiddleware = []echo.MiddlewareFunc{}
		}

		// default middleware
		s.Config.Server.DefaultMiddleware = append(s.Config.Server.DefaultMiddleware,
			middleware.RequestID(), // add request id
			middleware.Recover(),   // recover server from any panic/fatal error gracefully
			middleware.LoggerWithConfig(middleware.LoggerConfig{
				Format: "remote_ip=${remote_ip}, method=${method}, uri=${uri}, status=${status}, session=${header:Set-Cookie}, auth=${header:Authorization}\n",
			}),
			echoprometheus.MetricsMiddleware(),           // add prometheus metrics
			echozap.ZapLogger(s.Config.Logger.Desugar()), // add zap logger, middleware requires the "regular" zap logger
			echocontext.EchoContextToContextMiddleware(), // adds echo context to parent
			cors.New(),                     // add cors middleware
			mime.New(),                     // add mime middleware
			cachecontrol.New(),             // add cache control middleware
			ratelimit.DefaultRateLimiter(), // add ratelimit middleware
			middleware.Secure(),            // add XSS middleware
		)
	})
}

// WithEmailManager sets up the default SendGrid email manager to be used to send emails to users
// on registration, password reset, etc
func WithEmailManager() ServerOption {
	return newApplyFunc(func(s *ServerOptions) {
		emailConfig := &emails.Config{}

		// load defaults and env vars
		err := envconfig.Process("datum_email", emailConfig)
		if err != nil {
			panic(err)
		}

		em, err := emails.New(*emailConfig)
		if err != nil {
			panic(err)
		}

		urlConfig := &emails.URLConfig{}
		if err := envconfig.Process("datum_email_url", urlConfig); err != nil {
			panic(err)
		}

		if err := urlConfig.Validate(); err != nil {
			panic(err)
		}

		em.URLConfig = *urlConfig

		s.Config.Server.Handler.EmailManager = em
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

// WithRedisCache sets up the redis config use as a key-value store for things such as session management
func WithRedisCache() ServerOption {
	return newApplyFunc(func(s *ServerOptions) {
		config := &cache.Config{}

		// load defaults and env vars
		if err := envconfig.Process("datum_redis", config); err != nil {
			panic(err)
		}

		s.Config.RedisConfig = *config
	})
}

// WithSessionManager sets up the default session manager with a 10 minute ttl
// with persistence to redis
func WithSessionManager(rc *redis.Client) ServerOption {
	return newApplyFunc(func(s *ServerOptions) {
		config := &sessions.Config{}

		// load defaults and env vars
		if err := envconfig.Process("datum_sessions", config); err != nil {
			panic(err)
		}

		cc := sessions.DefaultCookieConfig

		// In order for things to work in dev mode with localhost
		// we need to se the debug cookie config
		if s.Config.Server.Dev {
			cc = &sessions.DebugOnlyCookieConfig
		} else {
			cc.Name = sessions.DefaultCookieName
		}

		sm := sessions.NewCookieStore[map[string]string](cc,
			[]byte(config.SigningKey),
			[]byte(config.EncryptionKey),
		)

		// add session middleware, this has to be added after the authMiddleware so we have the user id
		// when we get to the session. this is also added here so its only added to the graph routes
		// REST routes are expected to add the session middleware, as required
		sessionConfig := sessions.NewSessionConfig(
			sm,
			sessions.WithPersistence(rc),
			sessions.WithLogger(s.Config.Logger),
		)

		// set cookie config to be used
		sessionConfig.CookieConfig = cc

		// Make the cookie session store available
		// to graph and REST endpoints
		s.Config.Server.Handler.SessionConfig = &sessionConfig
		s.Config.Server.SessionConfig = &sessionConfig

		s.Config.Server.GraphMiddleware = append(s.Config.Server.GraphMiddleware,
			sessions.LoadAndSaveWithConfig(sessionConfig),
		)
	})
}
