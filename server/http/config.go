package http

import (
	"fmt"
	"time"
)

const (
	defaultPort             = 8080
	defaultGracefulDuration = 12 * time.Second
)

// Config http config server
type Config struct {
	// Port optional, http port to be exposed, 8080 by default
	Port int

	// Name optional, http server name to be exposed
	Name string

	// GracefulDuration optional, graceful duration to shut down the server, 12 second by default
	GracefulDuration time.Duration
}

func (c Config) logPrefix() string {
	if c.Name != "" {
		return fmt.Sprintf("[http-server: %s]", c.Name)
	}

	return `[http-server]`
}

func sanitizeConfig(cfg Config) Config {
	if cfg.Port == 0 {
		cfg.Port = defaultPort
	}

	if cfg.GracefulDuration == 0 {
		cfg.GracefulDuration = defaultGracefulDuration
	}

	return cfg
}
