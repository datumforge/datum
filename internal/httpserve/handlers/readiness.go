package handlers

// // AddReadinessCheck will accept a function to be ran during calls to /readyz
// // These functions should accept a context and only return an error. When adding
// // a readiness check a name is also provided, this name will be used when returning
// // the state of all the checks
// func (s *Handler) AddReadinessCheck(name string, f CheckFunc) *Handler {
// 	s.readinessChecks[name] = f

// 	return s
// }

// // readinessCheckHandler ensures that the server is up and that we are able to process requests
// func (s *Handler) readinessCheckHandler(c echo.Context) error {
// 	failed := false
// 	status := map[string]string{}

// 	for name, check := range s.readinessChecks {
// 		if err := check(c.Request().Context()); err != nil {
// 			s.logger.Error("readiness check failed", zap.String("name", name), zap.Error(err))

// 			failed = true
// 			status[name] = err.Error()
// 		} else {
// 			status[name] = "OK"
// 		}
// 	}

// 	if failed {
// 		return c.JSON(http.StatusServiceUnavailable, status)
// 	}

// 	return c.JSON(http.StatusOK, status)
// }
