package config

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: GetEnv("REDIS_ADDR", "localhost:6379"),
	})
	if _, err := RedisClient.Ping(context.Background()).Result(); err != nil {
		panic("Failed to connect to Redis: " + err.Error())
	}
}
