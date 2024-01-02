package viewer

import (
	"errors"
)

var (
	// ErrProcessingRequest is returned when the request cannot be processed
	ErrProcessingRequest = errors.New("error processing request, please try again")
)

// // Middleware returns a middleware function for to add viewer to the user's context
// func (d *Client) Middleware(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		// add to context
// 		v := resolveViewer(r)
// 		viewerCtx := viewer.NewContext(c.Request().Context(), viewer.NewUserViewerFromID(meowuser.ID, true))
// 		ctx := NewContext(c.Request().Context(), client)

// 		c.SetRequest(c.Request().WithContext(ctx))

// 		if err := next(c); err != nil {
// 			d.Logger.Debug("rolling back transaction in middleware")

// 			if err := client.Rollback(); err != nil {
// 				d.Logger.Errorw(rollbackErr, "error", err)

// 				return c.JSON(http.StatusInternalServerError, ErrProcessingRequest)
// 			}

// 			return err
// 		}

// 		d.Logger.Debug("committing transaction in middleware")

// 		if err := client.Commit(); err != nil {
// 			d.Logger.Errorw(transactionCommitErr, "error", err)

// 			return c.JSON(http.StatusInternalServerError, ErrProcessingRequest)
// 		}

// 		return nil
// 	}
// }
