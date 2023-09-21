package core_cache

import (
	"go-gin/config"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type IRedisHelper interface {
	GetClient() *redis.Client
	Init(
		client *redis.Client,
		log *logrus.Logger,
		config *config.AppEnv,
	)

	Get(key string) (string, error)

	Set(key string, value interface{}) error

	SetWithExpire(key string, value interface{}, expire int) error

	Delete(key string) error

	DeleteWithPrefix(prefix string) error

	DeleteAll() error

	IsExist(key string) bool

	Keys(pattern string) ([]string, error)
}
