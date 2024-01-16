package handlers

import (
	"github.com/lestrrat-go/jwx/v2/jwk"
	"go.uber.org/zap"

	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/tokens"
	"github.com/datumforge/datum/internal/utils/emails"
	"github.com/datumforge/datum/internal/utils/marionette"
)

// Handler contains configuration options for handlers
type Handler struct {
	// DBClient to interact with the generated ent schema
	DBClient *ent.Client
	// TM contains the token manager in order to validate auth requests
	TM *tokens.TokenManager
	// CookieDomain is the domain set in cookie for authenticated requests, defaults to datum.net
	CookieDomain string
	// Logger provides the zap logger to do logging things from the handlers
	Logger *zap.SugaredLogger
	// ReadyChecks is a set of checkFuncs to determine if the application is "ready" upon startup
	ReadyChecks Checks
	// JWTKeys contains the set of valid JWT authentication key
	JWTKeys jwk.Set
	// SendGridConfig containing the email configuration
	SendGridConfig *emails.Config
	// emailManager to handle sending emails
	emailManager *emails.EmailManager
	// EmailURL contains the urls used within emails
	EmailURL *URLConfig
	// TaskMan manages tasks in a separate goroutine to allow for non blocking operations
	TaskMan *marionette.TaskManager
}

type Response struct {
	StatusCode int         `json:"status,omitempty"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}
