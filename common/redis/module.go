package redisutil

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewRedisUtil),
)
