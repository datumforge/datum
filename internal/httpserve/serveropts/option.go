package serveropts

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"time"

	"entgo.io/ent/dialect"
	echo "github.com/datumforge/echox"
	"go.uber.org/zap"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/entdb"
	"github.com/datumforge/datum/internal/fga"
	"github.com/datumforge/datum/internal/graphapi"
	"github.com/datumforge/datum/internal/httpserve/config"
	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/internal/httpserve/server"
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
		s.Config.Logger = l
	})
}

// WithServer supplies the echo server config for the server
func WithServer(settings map[string]any) ServerOption {
	return newApplyFunc(func(s *ServerOptions) {
		serverSettings := settings["server"].(map[string]any)

		serverConfig := config.NewConfig().
			WithListen(serverSettings["listen"].(string)).                                    // set custom port
			WithHTTPS(serverSettings["https"].(bool)).                                        // enable https
			WithShutdownGracePeriod(serverSettings["shutdown-grace-period"].(time.Duration)). // override default grace period shutdown
			WithDebug(serverSettings["debug"].(bool)).                                        // enable debug mode
			WithDev(serverSettings["dev"].(bool)).                                            // enable dev mode
			SetDefaults()                                                                     // set defaults if not already set

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
		dbSettings := settings["db"].(map[string]any)

		// Database Settings
		dbConfig := config.DB{
			Debug:           serverSettings["debug"].(bool),
			MultiWrite:      dbSettings["multi-write"].(bool),
			DriverName:      dialect.SQLite,
			PrimaryDBSource: dbSettings["primary"].(string),
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
			s.Config.Authz = config.Authz{
				Enabled: false,
			}

			return
		}

		fgaSettings := settings["fga"].(map[string]any)

		// Authz Setup
		authzConfig := config.Authz{
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
func WithGeneratedKeys() ServerOption {
	return newApplyFunc(func(s *ServerOptions) {
		// Generate a new RSA private key with 2048 bits
		privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			s.Config.Logger.Panicf("Error generating RSA private key:", err)
		}

		// Encode the private key to the PEM format
		privateKeyPEM := &pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
		}

		privFileName := "private_key.pem"
		privateKeyFile, err := os.Create(privFileName)
		if err != nil {
			s.Config.Logger.Panicf("Error creating private key file:", err)
		}

		pem.Encode(privateKeyFile, privateKeyPEM)
		privateKeyFile.Close()

		// Extract the public key from the private key
		publicKey := &privateKey.PublicKey

		// Encode the public key to the PEM format
		publicKeyPEM := &pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(publicKey),
		}

		pubFileName := "public_key.pem"
		publicKeyFile, err := os.Create(pubFileName)
		if err != nil {
			fmt.Println("Error creating public key file:", err)
			os.Exit(1)
		}
		pem.Encode(publicKeyFile, publicKeyPEM)
		publicKeyFile.Close()

		fmt.Println("RSA key pair generated successfully!")

		keys := map[string]string{}

		kidPriv := ulids.New().String()
		kidPub := ulids.New().String()

		keys[kidPriv] = fmt.Sprintf("%v", privFileName)
		keys[kidPub] = fmt.Sprintf("%v", pubFileName)

		s.Config.Auth.Token.Keys = keys
	})
}

// WithAuth supplies the authn config for the server
// TODO: expand these settings
func WithAuth(settings map[string]any) ServerOption {
	return newApplyFunc(func(s *ServerOptions) {
		authEnabled := settings["auth"].(bool)
		// jwtSettings := settings["jwt"].(map[string]any)

		// Commenting out until this is implemented
		// oidcSettings := settings["oidc"].(map[string]any)

		s.Config.Auth.Enabled = authEnabled

		// add signing key for tokens
		//s.Config.Auth.JWTSigningKey = []byte(jwtSettings["signingkey"].(string))

		s.Config.Auth.Token.Issuer = "http://localhost:17608"
		s.Config.Auth.Token.Audience = "http://localhost:17608"
	})
}

// WithReadyChecks adds readiness checks to the server
func WithReadyChecks(c *entdb.EntClientConfig, f *fga.Client) ServerOption {
	return newApplyFunc(func(s *ServerOptions) {
		// Initialize checks
		s.Config.Server.Checks = handlers.Checks{}

		// Always add a check to the primary db connection
		s.Config.Server.Checks.AddReadinessCheck("sqlite_db_primary", entdb.Healthcheck(c.GetPrimaryDB()))

		// Check the secondary db, if enabled
		if s.Config.DB.MultiWrite {
			s.Config.Server.Checks.AddReadinessCheck("sqlite_db_secondary", entdb.Healthcheck(c.GetSecondaryDB()))
		}

		// Check the connection to openFGA, if enabled
		if s.Config.Authz.Enabled {
			s.Config.Server.Checks.AddReadinessCheck("fga", fga.Healthcheck(*f))
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
