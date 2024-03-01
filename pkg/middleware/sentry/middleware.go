package sentry

import (
	"bytes"
	"io"
	"net/http"
	"strconv"
	"time"

	echo "github.com/datumforge/echox"
	"github.com/datumforge/echox/middleware"
	"github.com/getsentry/sentry-go"

	response "github.com/datumforge/datum/pkg/utils/dumper"
)

// Config defines config for the Sentry middleware
type Config struct {
	// Skipper defines a function to skip middleware
	Skipper middleware.Skipper
	// AreHeadersDump defines whether to add req headers & resp headers to tracing tags
	AreHeadersDump bool
	// IsBodyDump defines whether to add req body & resp body to attributes
	IsBodyDump bool
	// Repanic configures whether Sentry should repanic after recovery, in most cases it should be set to true,
	// as echo includes it's own Recover middleware what handles http responses
	Repanic bool
	// WaitForDelivery configures whether you want to block the request before moving forward with the response
	// Because Echo's Recover handler doesn't restart the application,
	// it's safe to either skip this option or set it to false
	WaitForDelivery bool
	// Timeout for the event delivery requests
	Timeout time.Duration
}

// DefaultConfig is the default Sentry middleware config
var DefaultConfig = Config{
	Skipper:        middleware.DefaultSkipper,
	AreHeadersDump: true,
	IsBodyDump:     true,
}

// New returns a new Sentry middleware
func New() echo.MiddlewareFunc {
	return NewWithConfig(DefaultConfig)
}

// NewWithConfig returns a new Sentry middleware with config
func NewWithConfig(config Config) echo.MiddlewareFunc {
	if config.Skipper == nil {
		config.Skipper = middleware.DefaultSkipper
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) || c.Request() == nil || c.Response() == nil {
				return next(c)
			}

			request, span, endSpan := createSpan(c)
			defer endSpan()

			ctx := span.Context()

			setTag(span, "client_ip", c.RealIP())
			setTag(span, "remote_addr", request.RemoteAddr)
			setTag(span, "request_uri", request.RequestURI)
			setTag(span, "path", c.Path())

			respDumper := dumpReq(c, config, span, request)

			// setup request context - add span
			c.SetRequest(request.WithContext(ctx))

			err := next(c)
			if err != nil {
				setTag(span, "echo.error", err.Error())
				c.Error(err)
			}

			dumpResp(c, config, span, respDumper)

			return err
		}
	}
}

func dumpResp(c echo.Context, config Config, span *sentry.Span, respDumper *response.Dumper) {
	setTag(span, "request_id", getRequestID(c))
	span.Status = sentry.HTTPtoSpanStatus(c.Response().Status)
	setTag(span, "resp.status", strconv.Itoa(c.Response().Status))

	// Dump response headers
	if config.AreHeadersDump {
		for k := range c.Response().Header() {
			setTag(span, "resp.header."+k, c.Response().Header().Get(k))
		}
	}

	// Dump response body
	if config.IsBodyDump {
		setTag(span, "resp.body", respDumper.GetResponse())
	}
}

func dumpReq(c echo.Context, config Config, span *sentry.Span, request *http.Request) *response.Dumper {
	if username, _, ok := request.BasicAuth(); ok {
		setTag(span, "user", username)
	}

	// Add path parameters
	for _, paramName := range c.PathParams() {
		setTag(span, "path."+paramName.Name, c.PathParam(paramName.Name))
	}

	// Dump request headers
	if config.AreHeadersDump {
		for k := range request.Header {
			setTag(span, "req.header."+k, request.Header.Get(k))
		}
	}

	// Dump request & response body
	var respDumper *response.Dumper

	if config.IsBodyDump {
		// request
		if request.Body != nil {
			reqBody, err := io.ReadAll(request.Body)
			if err == nil {
				setTag(span, "req.body", string(reqBody))

				_ = request.Body.Close()
				request.Body = io.NopCloser(bytes.NewBuffer(reqBody)) // reset original request body
			}
		}

		// response
		respDumper = response.NewDumper(c.Response())
		c.Response().Writer = respDumper
	}

	return respDumper
}

func createSpan(c echo.Context) (*http.Request, *sentry.Span, func()) {
	request := c.Request()
	savedCtx := request.Context()
	opname := "HTTP " + request.Method + " " + c.Path()
	tname := "HTTP " + request.Method + " " + c.Request().RequestURI
	span := sentry.StartSpan(savedCtx, opname, sentry.WithTransactionName(tname))

	return request, span, func() {
		request = request.WithContext(savedCtx)
		c.SetRequest(request)

		defer span.Finish()
	}
}
