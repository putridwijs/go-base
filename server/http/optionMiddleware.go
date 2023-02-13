package http

import "github.com/labstack/echo/v4"

type registerMiddleware []echo.MiddlewareFunc

// Apply function to implement Option
func (r registerMiddleware) Apply(o *options) {
	o.middlewares = append(o.middlewares, r...)
}

// RegisterMiddleware function to add middleware to server
func RegisterMiddleware(fn echo.MiddlewareFunc) Option {
	return registerMiddleware([]echo.MiddlewareFunc{fn})
}

// RegisterMiddlewares function to add multiple middleware to server
func RegisterMiddlewares(fn []echo.MiddlewareFunc) Option {
	return registerMiddleware(fn)
}
