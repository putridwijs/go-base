package main

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"go-base/cmd/http"
	"go-base/internal/config"
)

var (
	cfg config.Config
	cmd = &cobra.Command{
		Use:   "go-base",
		Short: "go-base",
	}
)

func main() {
	cfg = config.InitConfig()
	globalLogLevel := zerolog.InfoLevel
	zerolog.SetGlobalLevel(globalLogLevel)

	cmd.AddCommand(
		http.Cmd(cfg),
	)

	// execute command
	if err := cmd.Execute(); err != nil {
		panic("can't execute cmd")
	}
}
