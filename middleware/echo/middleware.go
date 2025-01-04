package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/quentinguidee/errors"
)

func Error(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			var target *errors.HTTPError
			if errors.As(err, &target) {
				if target.Code >= 500 {
					// internal server error must be logged
					logError(target)
				}
				return c.JSON(target.StatusCode(), target)
			}
			logError(err)
			return c.JSON(http.StatusInternalServerError, &errors.HTTPError{
				Code:    errors.ErrorCodeInternalServerError,
				Message: "An unknown error occurred.",
			})
		}
		return nil
	}
}

func logError(err error) {
	_, _ = fmt.Fprintln(os.Stderr, err)
}
