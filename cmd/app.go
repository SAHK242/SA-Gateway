package cmd

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"gateway/api"
	"gateway/common"
	"gateway/grpc"
	"gateway/server"
	"github.com/redis/go-redis/v9"
)

func Start() {
	load()
}

func load() {
	fx.New(
		// Provide a zap logger first
		fx.Provide(
			func() *zap.SugaredLogger {
				logger, _ := zap.NewProduction()
				return logger.Sugar()
			},
			// Provide Redis client
			func() *redis.Client {
				client := redis.NewClient(&redis.Options{
					Addr: "localhost:6379", // Change if needed
					DB:   0,                // Default DB
				})
				return client
			},
		),
		// Then use it for Fx's logger
		fx.WithLogger(func(logger *zap.SugaredLogger) fxevent.Logger {
			return &fxevent.ZapLogger{
				Logger: logger.Desugar().WithOptions(zap.IncreaseLevel(zapcore.WarnLevel)),
			}
		}),
		common.Module,
		grpc.Module,
		api.Module,
		server.Module,
		fx.Invoke(func(*fiber.App) {}),
	).Run()
}
