package utils

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var REDIS_CTX = context.Background()

type ICacheHelper interface {
	IsLogoutSession(ssId string) bool
}

type RedisHelper struct {
	ICacheHelper
	redisClient            *redis.Client
	redisClientForceLogout *redis.Client
}

func (redis *RedisHelper) Init(client *redis.Client, clientForceLogout *redis.Client) {
	redis.redisClient = client
	redis.redisClientForceLogout = clientForceLogout
}

func (redis *RedisHelper) IsLogoutSession(ssId string) bool {
	// Check if session id is logout force
	jsonResult, err := redis.redisClientForceLogout.Get(context.Background(), "user-id:force:"+ssId).Result()
	if err != nil {
		_ = fmt.Errorf("%v", err)
	}
	return len(jsonResult) > 0
}
