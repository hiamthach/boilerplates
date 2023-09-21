package apphost

import (
	"context"
	"go-gin/app/controllers"
	envConfig "go-gin/config"
	"go-gin/infra/mongo/repository"
	"log"
	"os"
	"time"

	"github.com/golobby/container/v3"
	"github.com/golobby/dotenv"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	cacheUtil "go-gin/core/cache"
)

var AppConfig = envConfig.AppEnv{}
var AppLog = logrus.New()

func IoCConfig() {
	container.SingletonLazy(func() envConfig.AppEnv {
		var config = envConfig.AppEnv{}
		file, errReadFile := os.Open("app.env")
		if errReadFile != nil {
			log.Println("error when read file ", errReadFile)
		} else {
			errorLoadCf := dotenv.NewDecoder(file).Decode(&config)
			if errorLoadCf != nil {
				log.Println("error when load app config ", errorLoadCf)
			}
		}
		return config
	})
}

func IoCLog() {
	// Log as JSON instead of the default ASCII formatter.
	AppLog.SetFormatter(&logrus.JSONFormatter{})

	AppLog.Out = os.Stdout
	AppLog.SetReportCaller(false)
	// Only log the warning severity or above.
	minLevel := AppConfig.Log.MinLevel
	AppLog.SetLevel(logrus.Level(minLevel))
	AppLog.WithField("AppName", AppConfig.App.Name)
	container.SingletonLazy(func() *logrus.Logger {
		return AppLog
	})
}

func IoCRedis() {
	container.Singleton(func() *redis.Client {
		return redis.NewClient(&redis.Options{
			Addr:     AppConfig.Redis.Host,
			Password: AppConfig.Redis.Password, // no password set
			DB:       0,                        // use default DB
		})
	})

	container.SingletonLazy(func() cacheUtil.IRedisHelper {
		start := time.Now()
		redisServer := AppConfig.Redis.Host

		var redisClient = redis.NewClient(&redis.Options{
			Addr:     redisServer,
			Password: AppConfig.Redis.Password, // no password set
			DB:       0,                        // use default DB
		})

		_, err := redisClient.Ping(context.TODO()).Result()
		if err != nil {
			AppLog.Error("error when connect to redis ", err)
		} else {
			AppLog.Info("connected to redis ", redisServer)
		}

		redisHelper := cacheUtil.RedisHelper{}
		redisHelper.Init(redisClient, AppLog, &AppConfig)
		elapsed := time.Since(start)
		AppLog.Info("elapsed time when init redis ", elapsed)
		return &redisHelper
	})
}

func IoCMongo() {
	container.SingletonLazy(func() *mongo.Client {
		start := time.Now()
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
		defer cancel()

		clientOpts := options.Client().ApplyURI(AppConfig.MongoDb.Connection)
		client, err := mongo.Connect(ctx, clientOpts)
		if err != nil {
			AppLog.Error("error when connect to mongo ", err)
		} else {
			AppLog.Info("connected to mongo ", AppConfig.MongoDb.Connection)
		}

		elapsed := time.Since(start)
		AppLog.Info("elapsed time when init mongo ", elapsed)
		return client
	})
}

func IoCServices() {
	container.TransientLazy(func() repository.IUserRepository {
		var repo repository.UserRepository
		repo.Init(AppConfig, AppLog)
		return &repo
	})
}

func IoCControllers() {
	container.TransientLazy(func() controllers.IUserController {
		var (
			userRepo repository.IUserRepository
		)
		container.Resolve(&userRepo)

		controller := controllers.UserController{}
		controller.Init(AppConfig, AppLog, userRepo)
		return &controller
	})
}

func IoC() {
	IoCConfig()
	err := container.Resolve(&AppConfig)
	if err != nil {
		AppLog.Error("error when resolve app config ", err)
	}

	IoCLog()
	IoCRedis()
	IoCMongo()
	IoCServices()
	IoCControllers()
}
