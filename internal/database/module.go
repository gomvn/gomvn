package database

import (
	"context"

	"github.com/jinzhu/gorm"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(New),
	fx.Invoke(register),
)

func register(lifecycle fx.Lifecycle, db *gorm.DB) {
	lifecycle.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return db.Close()
		},
	})
}
