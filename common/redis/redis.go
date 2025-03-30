package redisutil

import "github.com/redis/go-redis/v9"

type RedisUtil struct {
	Redis *redis.Client
}

func NewRedisUtil(redis *redis.Client) *RedisUtil {
	return &RedisUtil{
		Redis: redis,
	}
}
