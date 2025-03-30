package config

import (
	"log"
)

var (
	logger = log.Default()

	commonRequiredConfig = []string{
		"PROFILE",
		"HTTP_PORT",
		"GRPC_PORT",
		"GRPC_TIMEOUT",
		"GRPC_LATENCY_SLA",
		"GRPC_MAX_HEADER_SIZE",
		"REDIS_HOST",
		"REDIS_PORT",
		"REDIS_DEFAULT_TTL_MS",
		"REDIS_MIN_CONN",
		"DOMAIN_FRONTEND",
		"DOMAIN_BACKEND",
	}

	defaultRequiredConfig = []string{
		"APPLICATION",
	}
)

type (
	options struct {
		requiredConfig []string
	}

	Option func(*options)
)

func WithRequiredConfig(requiredConfig []string) Option {
	return func(o *options) {
		configs := make([]string, 0)
		configs = append(configs, commonRequiredConfig...)
		configs = append(configs, defaultRequiredConfig...)

		if requiredConfig == nil {
			o.requiredConfig = configs
		} else {
			o.requiredConfig = append(configs, requiredConfig...)
		}
	}
}

func eval(opts []Option) *options {
	o := &options{
		requiredConfig: append(commonRequiredConfig, defaultRequiredConfig...),
	}

	for _, opt := range opts {
		opt(o)
	}

	return o
}
