package http

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
	"os"
	"syscall"
)

type (
	IServer interface {
		Run() error
		Shutdown() error
	}

	httpServer struct {
		sig chan os.Signal
		cfg Config
		opt options

		// dependencies
		ec *echo.Echo
	}
)

// init function to initialize dependencies
func (s *httpServer) init() error {
	s.ec = echo.New()
	return nil
}

// middlewares function to register middleware to http server
func (s *httpServer) middlewares() []echo.MiddlewareFunc {
	middlewares := []echo.MiddlewareFunc{
		middleware.Recover(),
		middleware.Gzip(),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowHeaders: []string{"*"},
			AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		}),
		middleware.RequestIDWithConfig(middleware.RequestIDConfig{Generator: uuid.New().String}),
		middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
			LogURI:    true,
			LogStatus: true,
			LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
				log.Info().
					Str("URI", v.URI).
					Int("status", v.Status).
					Msgf("%s request", s.cfg.logPrefix())
				return nil
			},
		}),
	}
	middlewares = append(middlewares, s.opt.middlewares...)
	return middlewares
}

// routes function to initialize http routes
func (s *httpServer) routes() {
	for _, route := range s.opt.routes {
		route(s.ec)
	}
}

// Run function to run http server
func (s *httpServer) Run() error {
	var parameters = map[string]interface{}{"server": s}
	defer func() {
		log.Info().
			Fields(parameters).
			Msgf("%s terminating server", s.cfg.logPrefix())
		s.sig <- syscall.SIGTERM
	}()

	// initialize http
	log.Info().Fields(parameters).Msgf("%s server initialized", s.cfg.logPrefix())
	s.ec.Use(s.middlewares()...)
	s.ec.HTTPErrorHandler = defaultErrorHandler

	// register routes
	log.Info().Fields(parameters).Msgf("%s registering server routes", s.cfg.logPrefix())
	s.routes()

	// run http
	log.Info().Fields(parameters).Msgf("%s starting server", s.cfg.logPrefix())
	if err := s.ec.Start(fmt.Sprintf(":%d", s.cfg.Port)); err != nil {
		log.Error().Err(err).Fields(parameters).Msgf("%s failed to start server", s.cfg.logPrefix())
		return err
	}
	return nil
}

// Shutdown function to close http server
func (s *httpServer) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.cfg.GracefulDuration)
	defer cancel()

	if err := s.ec.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}

// Server functions to initialize http server
func Server(sig chan os.Signal, cfg Config, opts ...Option) (IServer, error) {
	cfg = sanitizeConfig(cfg)
	option := defaultOption
	for _, opt := range opts {
		opt.Apply(&option)
	}

	s := &httpServer{sig: sig, cfg: cfg, opt: option}
	if err := s.init(); err != nil {
		log.Error().Err(err).
			Fields(map[string]interface{}{"config": cfg}).
			Msgf("%s failed to initialized server", cfg.logPrefix())
		return nil, err
	}
	return s, nil
}
