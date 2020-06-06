package service

import (
	"go.uber.org/fx"

	"github.com/gomvn/gomvn/internal/service/user"
)

var Module = fx.Options(
	fx.Provide(NewPathService),
	fx.Provide(NewStorage),
	fx.Provide(user.New),
	fx.Invoke(user.Initialize),
)
