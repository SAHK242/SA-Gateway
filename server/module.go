package server

import (
	"gateway/config"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"gateway/api/route/common"
)

type HttpServerProps struct {
	fx.In
	fx.Lifecycle
	Config  config.Config
	Logger  *zap.SugaredLogger
	Routers []common.Router `group:"routers"`
}

var Module = fx.Provide(
	NewHttpServer,
)
