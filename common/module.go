package common

import (
	"gateway/common/config"
	redisutil "gateway/common/redis"
	"go.uber.org/fx"
)

var Module = fx.Module("common",
	config.Module,
	redisutil.Module,
)
