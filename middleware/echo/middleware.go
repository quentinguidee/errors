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
			if errors.As(err, &target) && target.Code < 500 {
				return c.JSON(target.StatusCode(), target)
			}
			_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
			return c.JSON(http.StatusInternalServerError, &errors.HTTPError{
				Code:    errors.ErrorCodeInternalServerError,
				Name:    "ERR_UNKNOWN",
				Message: "An unknown error occurred.",
			})
		}
		return nil
	}
}
