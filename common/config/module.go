package config

import (
	cfg "gateway/config"
	"go.uber.org/fx"
)

func NewConfig() cfg.Config {
	return cfg.NewViper(
		cfg.WithRequiredConfig([]string{
			"CONTEXT_PATH",
			"LIMITER_MAX",
			"LIMITER_EXPIRATION",
		}),
	)
}

var Module = fx.Provide(NewConfig)
