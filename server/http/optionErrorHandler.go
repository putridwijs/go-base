package http

import "github.com/labstack/echo/v4"

type errHandler func(err error, c echo.Context)

// Apply function to implement Option
func (e errHandler) Apply(o *options) {
	o.errorHandler = e
}

// ErrorHandler function to override server error handler
func ErrorHandler(fn func(err error, c echo.Context)) Option {
	return errHandler(fn)
}
