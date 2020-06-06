package server

import (
	"context"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(New),
	fx.Invoke(register),
)

func register(lifecycle fx.Lifecycle, server *Server) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return server.Listen()
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown()
		},
	})
}
