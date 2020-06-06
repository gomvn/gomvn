package main

import (
	"flag"

	"go.uber.org/fx"

	"github.com/gomvn/gomvn/internal/config"
	"github.com/gomvn/gomvn/internal/server"
	"github.com/gomvn/gomvn/internal/service"
)

func main() {
	cf := flag.String("config", "config.yml", "path to config file")
	flag.Parse()

	app := fx.New(
		fx.NopLogger,
		config.Module(*cf),
		service.Module,
		server.Module,
	)
	app.Run()
}
