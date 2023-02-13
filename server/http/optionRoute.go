package http

import "github.com/labstack/echo/v4"

type registerRoutes []func(ec *echo.Echo)

// Apply function to implement Option
func (r registerRoutes) Apply(o *options) {
	o.routes = append(o.routes, r...)
}

// RegisterRoute function to add route to server
func RegisterRoute(fn func(*echo.Echo)) Option {
	return registerRoutes([]func(ec *echo.Echo){fn})
}

// RegisterRoutes function to add multi route to server
func RegisterRoutes(fn []func(*echo.Echo)) Option {
	return registerRoutes(fn)
}
