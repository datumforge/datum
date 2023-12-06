package handlers

import (
	"context"
	"net/http"

	echo "github.com/datumforge/echox"
)

// CheckFunc is a function that can be used to check the status of a service
type CheckFunc func(ctx context.Context) error

type Checks struct {
	readinessChecks map[string]CheckFunc
}

// AddReadinessCheck will accept a function to be ran during calls to /readyz
// These functions should accept a context and only return an error. When adding
// a readiness check a name is also provided, this name will be used when returning
// the state of all the checks
func (c *Checks) AddReadinessCheck(name string, f CheckFunc) {
	// if this is null, create the struct before trying to add
	if c.readinessChecks == nil {
		c.readinessChecks = map[string]CheckFunc{}
	}

	c.readinessChecks[name] = f
}

func (c *Checks) ReadyHandler(ctx echo.Context) error {
	failed := false
	status := map[string]string{}

	for name, check := range c.readinessChecks {
		if err := check(ctx.Request().Context()); err != nil {
			failed = true
			status[name] = err.Error()
		} else {
			status[name] = "OK"
		}
	}

	if failed {
		return ctx.JSON(http.StatusServiceUnavailable, status)
	}

	return ctx.JSON(http.StatusOK, status)
}
