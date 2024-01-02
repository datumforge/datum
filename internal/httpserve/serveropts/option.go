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

	"entgo.io/ent/dialect"
	"github.com/alexedwards/scs/v2"
	"github.com/alexedwards/scs/v2/memstore"
	echo "github.com/datumforge/echox"
	"go.uber.org/zap"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/entdb"
	"github.com/datumforge/datum/internal/fga"
	"github.com/datumforge/datum/internal/graphapi"
	"github.com/datumforge/datum/internal/httpserve/config"
	"github.com/datumforge/datum/internal/httpserve/server"
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
func WithServer(settings map[string]any) ServerOption {
	return newApplyFunc(func(s *ServerOptions) {
		serverSettings := settings["server"].(map[string]any)
		logSettings := serverSettings["logging"].(map[string]any)

		shutdown, err := time.ParseDuration(serverSettings["shutdown-grace-period"].(string))
		if err != nil {
			s.Config.Logger.Panicw("unable to parse shutdown duration", "error", err.Error())
		}

		serverConfig := config.NewConfig().
			WithListen(serverSettings["listen"].(string)). // set custom port
			WithHTTPS(serverSettings["https"].(bool)).     // enable https
			WithShutdownGracePeriod(shutdown).             // override default grace period shutdown
			WithDebug(logSettings["debug"].(bool)).        // enable debug mode
			WithDev(serverSettings["dev"].(bool)).         // enable dev mode
			SetDefaults()                                  // set defaults if not already set

		s.Config = *serverConfig
	})
}

// WithHTTPS sets up TLS config settings for the server
func WithHTTPS(settings map[string]any) ServerOption {
	return newApplyFunc(func(s *ServerOptions) {
		serverSettings := settings["server"].(map[string]any)

		if !s.Config.Server.TLS.Enabled {
			// this is set to enabled by WithServer
			// if TLS is not enabled, move on
			return
		}

		s.Config.WithTLSDefaults()

		autoCert := serverSettings["auto-cert"].(bool)

		if autoCert {
			s.Config.WithAutoCert(serverSettings["cert-host"].(string))
		} else {
			cf := serverSettings["ssl-cert"].(string)
			k := serverSettings["ssl-key"].(string)

			certFile, certKey, err := server.GetCertFiles(cf, k)
			if err != nil {
				// if this errors, we should panic because a required file is not found
				s.Config.Logger.Panicw("unable to start https server", "error", err.Error())
			}

			s.Config.WithTLSCerts(certFile, certKey)
		}
	})
}

// WithSQLiteDB supplies the sqlite db config for the server
func WithSQLiteDB(settings map[string]any) ServerOption {
	return newApplyFunc(func(s *ServerOptions) {
		serverSettings := settings["server"].(map[string]any)
		logSettings := serverSettings["logging"].(map[string]any)
		dbSettings := settings["db"].(map[string]any)

		// Database Settings
		dbConfig := config.DB{
			Debug:           logSettings["debug"].(bool),
			MultiWrite:      dbSettings["multi-write"].(bool),
			DriverName:      dialect.SQLite,
			PrimaryDBSource: dbSettings["primary"].(string),
			CacheTTL:        entdb.DefaultCacheTTL,
		}

		if dbConfig.MultiWrite {
			dbConfig.SecondaryDBSource = dbSettings["secondary"].(string)
		}

		s.Config.DB = dbConfig
	})
}

// WithFGAAuthz supplies the FGA authz config for the server
func WithFGAAuthz(settings map[string]any) ServerOption {
	return newApplyFunc(func(s *ServerOptions) {
		authzEnabled := settings["auth"].(bool)

		if !authzEnabled {
			s.Config.Authz = fga.Config{
				Enabled: false,
			}

			return
		}

		fgaSettings := settings["fga"].(map[string]any)

		// Authz Setup
		authzConfig := fga.Config{
			Enabled:        authzEnabled,
			StoreName:      "datum",
			Host:           fgaSettings["host"].(string),
			Scheme:         fgaSettings["scheme"].(string),
			StoreID:        fgaSettings["store"].(map[string]any)["id"].(string),
			ModelID:        fgaSettings["model"].(map[string]any)["id"].(string),
			CreateNewModel: fgaSettings["model"].(map[string]any)["create"].(bool),
		}

		s.Config.Authz = authzConfig
	})
}

// WithGeneratedKeys will generate a public/private key pair
// that can be used for jwt signing.
// This should only be used in a development environment
func WithGeneratedKeys(settings map[string]any) ServerOption {
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
		var kidPriv string
		jwtSettings, ok := settings["jwt"].(map[string]any)
		if ok {
			kid, ok := jwtSettings["kid"].(string)
			if ok {
				kidPriv = kid
			}
		}

		// if we didn't get a kid in the settings, assign one
		if kidPriv == "" {
			kidPriv = ulids.New().String()
		}

		keys[kidPriv] = fmt.Sprintf("%v", privFileName)

		s.Config.Server.Token.Keys = keys
	})
}

// WithAuth supplies the authn and jwt config for the server
func WithAuth(settings map[string]any) ServerOption {
	return newApplyFunc(func(s *ServerOptions) {
		authEnabled := settings["auth"].(bool)
		jwtSettings := settings["jwt"].(map[string]any)

		s.Config.Auth.Enabled = authEnabled

		s.Config.Server.Token.Issuer = jwtSettings["issuer"].(string)
		s.Config.Server.Token.Audience = jwtSettings["audience"].(string)
		s.Config.Server.Token.CookieDomain = jwtSettings["cookie-domain"].(string)

		// Set durations, flags have defaults values set
		accessDuration, err := time.ParseDuration(jwtSettings["access-duration"].(string))
		if err != nil {
			s.Config.Logger.Panicw("unable to parse shutdown duration", "error", err.Error())
		}

		refreshDuration, err := time.ParseDuration(jwtSettings["refresh-duration"].(string))
		if err != nil {
			s.Config.Logger.Panicw("unable to parse shutdown duration", "error", err.Error())
		}

		refreshOverlap, err := time.ParseDuration(jwtSettings["refresh-overlap"].(string))
		if err != nil {
			s.Config.Logger.Panicw("unable to parse shutdown duration", "error", err.Error())
		}

		s.Config.Server.Token.AccessDuration = accessDuration
		s.Config.Server.Token.RefreshDuration = refreshDuration
		s.Config.Server.Token.RefreshOverlap = refreshOverlap
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
func WithGraphRoute(srv *server.Server, c *generated.Client, settings map[string]any, mw []echo.MiddlewareFunc) ServerOption {
	return newApplyFunc(func(s *ServerOptions) {
		serverSettings := settings["server"].(map[string]any)

		// Setup Graph API Handlers
		r := graphapi.NewResolver(c, s.Config.Auth.Enabled).
			WithLogger(s.Config.Logger.Named("resolvers"))

		handler := r.Handler(serverSettings["dev"].(bool), mw...)

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
