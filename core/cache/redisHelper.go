package core_cache

import (
	"context"
	"go-gin/config"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type RedisHelper struct {
	IRedisHelper
	logger *logrus.Logger
	client *redis.Client
	config *config.AppEnv
}

func (r *RedisHelper) Init(
	client *redis.Client,
	log *logrus.Logger,
	config *config.AppEnv,
) {
	r.logger = log
	r.client = client
	r.config = config
}

func (r *RedisHelper) GetClient() *redis.Client {
	return r.client
}

func (r *RedisHelper) Get(key string) (string, error) {
	return r.client.Get(context.Background(), key).Result()
}

func (r *RedisHelper) Set(key string, value interface{}) error {
	return r.client.Set(context.Background(), key, value, 0).Err()
}

func (r *RedisHelper) SetWithExpire(key string, value interface{}, expire int) error {
	return r.client.Set(context.Background(), key, value, 0).Err()
}

func (r *RedisHelper) Delete(key string) error {
	return r.client.Del(context.Background(), key).Err()
}

func (r *RedisHelper) DeleteWithPrefix(prefix string) error {
	keys, err := r.client.Keys(context.Background(), prefix+"*").Result()
	if err != nil {
		return err
	}
	return r.client.Del(context.Background(), keys...).Err()
}

func (r *RedisHelper) DeleteAll() error {
	return r.client.FlushAll(context.Background()).Err()
}

func (r *RedisHelper) Exists(key string) bool {
	return r.client.Exists(context.Background(), key).Val() > 0
}
