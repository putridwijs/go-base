package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type (
	Option interface {
		Apply(o *options)
	}
	options struct {
		errorHandler func(err error, c echo.Context)
		middlewares  []echo.MiddlewareFunc
		routes       []func(ec *echo.Echo)
	}
)

var (
	defaultErrorHandler = func(err error, c echo.Context) {
		_ = c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	defaultOption = options{
		errorHandler: defaultErrorHandler,
		middlewares:  make([]echo.MiddlewareFunc, 0),
		routes:       make([]func(echo2 *echo.Echo), 0),
	}
)
