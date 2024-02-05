package sessions

import (
	"net/http"

	echo "github.com/datumforge/echox"
	"github.com/datumforge/echox/middleware"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

// SessionConfig is used to configure session management
type SessionConfig struct {
	// Skipper is a function that determines whether a particular request should be skipped or not
	Skipper middleware.Skipper
	// SessionManager is responsible for managing the session cookies. It handles the creation, retrieval, and deletion of
	// session cookies for each user session
	SessionManager Store[map[string]string]
	// RedisStore is used to store and retrieve session data in a persistent manner such as to a redis backend
	RedisStore PersistentStore
	// RedisClient establishes a connection to a Redis server and perform operations such as storing and retrieving data
	RedisClient *redis.Client
	// Logger is used to log errors in the middleware
	Logger *zap.SugaredLogger
}

var DefaultSessionConfig = SessionConfig{
	Skipper: middleware.DefaultSkipper,
}

// NewSessionConfig creates a new session config
func NewSessionConfig(sm Store[map[string]string], rc *redis.Client, logger *zap.SugaredLogger) *SessionConfig {
	c := &DefaultSessionConfig
	c.SessionManager = sm
	c.RedisClient = rc
	c.Logger = logger
	c.RedisStore = NewStore(rc)

	return c
}

// SaveAndStoreSession saves the session to the cookie and to the persistent store (redis)
func (sc *SessionConfig) SaveAndStoreSession(ctx echo.Context, name string, userID string) error {
	session := sc.SessionManager.New(name)
	sessionID := GenerateSessionID()

	setSessionMap := map[string]string{}
	setSessionMap["userID"] = userID

	session.Set(sessionID, setSessionMap)

	// Add session to context
	c := session.addSessionDataToContext(ctx.Request().Context())
	ctx.SetRequest(ctx.Request().WithContext(c))

	if err := session.Save(ctx.Response().Writer); err != nil {
		return err
	}

	if err := sc.RedisStore.StoreSession(ctx.Request().Context(), sessionID, userID); err != nil {
		return err
	}

	return nil
}

// LoadAndSave is a middleware function that loads and saves session data using a
// provided session manager. It takes a `SessionManager` as input and returns a middleware function
// that can be used with an Echo framework application
func LoadAndSave(sessionManager Store[map[string]string], client *redis.Client, logger *zap.SugaredLogger) echo.MiddlewareFunc {
	c := NewSessionConfig(sessionManager, client, logger)

	return LoadAndSaveWithConfig(c)
}

// LoadAndSaveWithConfig is a middleware that loads and saves session data
// using a provided session manager configuration. It takes a `SessionConfig` struct as input, which
// contains the skipper function and the session manager
func LoadAndSaveWithConfig(config *SessionConfig) echo.MiddlewareFunc {
	if config.Skipper == nil {
		config.Skipper = DefaultSessionConfig.Skipper
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			// get sessionData from request cookies
			session, err := config.SessionManager.Get(c.Request(), DefaultCookieName)
			if err != nil {
				config.Logger.Errorw("unable to get session", "error", err)

				return err
			}

			// get the session id from the session data
			sessionID := config.SessionManager.GetSessionIDFromCookie(session)
			sessionData := config.SessionManager.GetSessionDataFromCookie(session)

			// check session token on request matches cache
			userIDFromCookie := sessionData.(map[string]string)[UserIDKey]

			// lookup userID in cache to ensure tokens match
			userID, err := config.RedisStore.GetSession(c.Request().Context(), sessionID)
			if err != nil {
				config.Logger.Errorw("unable to get session from store", "error", err)

				return err
			}

			if userIDFromCookie != userID {
				config.Logger.Errorw("sessions do not match", "cookie", userIDFromCookie, "store", userID)
			}

			// Add session to context to be used in request paths
			ctx := session.addSessionDataToContext(c.Request().Context())
			c.SetRequest(c.Request().WithContext(ctx))

			c.Response().Before(func() {
				// refresh and save session cookie
				if err := config.SaveAndStoreSession(c, DefaultCookieName, sessionID); err != nil {
					config.Logger.Errorw("unable to create and store new session", "error", err)

					panic(err)
				}

				addHeaderIfMissing(c.Response(), "Cache-Control", `no-cache="Set-Cookie"`)
				addHeaderIfMissing(c.Response(), "Vary", "Cookie")
			})

			return next(c)
		}
	}
}

// addHeaderIfMissing function is used to add a header to the HTTP response if it is not already
// present. It takes in the response writer (`http.ResponseWriter`), the header key, and the header
// value as parameters
func addHeaderIfMissing(w http.ResponseWriter, key, value string) {
	for _, h := range w.Header()[key] {
		if h == value {
			return
		}
	}

	w.Header().Add(key, value)
}
