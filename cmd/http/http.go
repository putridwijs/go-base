package http

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"go-base/internal/config"
	"go-base/internal/pkg/greeting/encoding"
	"go-base/internal/pkg/greeting/endpoint"
	echo2 "go-base/server/echo"
	"go-base/server/http"

	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type IServer interface {
	Run() error
	Shutdown() error
}

func validatedServerConfig(cfg config.Config) bool {
	return cfg.IsValidHTTP()
}

func initializeServer(sig chan os.Signal, cfg config.Config) (IServer, error) {
	return http.Server(sig,
		http.Config{
			Port:             cfg.HTTP.Port,
			GracefulDuration: cfg.HTTP.GracefulDuration,
		},
		http.RegisterRoute(func(ec *echo.Echo) {
			greeting := endpoint.NewEndpoint(cfg)
			greet := ec.Group("/greetings")
			greet.GET("", echo2.Handler(
				greeting.Greet(),
				echo2.WithDecoder(encoding.DecodeGreetingRequest()),
			))
		}),
	)
}

func runner(cfg config.Config) func(c *cobra.Command, args []string) error {
	return func(_ *cobra.Command, args []string) error {
		if !validatedServerConfig(cfg) {
			return fmt.Errorf("invalid required config for <no value>")
		}

		log.Info().Msgf("[http-server] starting server with [%s] log level", zerolog.GlobalLevel().String())

		sigChannel := make(chan os.Signal, 1)
		signal.Notify(sigChannel, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

		s, err := initializeServer(sigChannel, cfg)
		if err != nil {
			return err
		}

		go s.Run()

		sig := <-sigChannel
		log.Info().Msgf("[http-server] signal %s received, exiting", sig.String())
		if err = s.Shutdown(); err != nil {
			return err
		}
		return nil
	}
}

// Cms expose command runner
func Cmd(cfg config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "http",
		Short: "Run http server",
		RunE:  runner(cfg),
	}
}
