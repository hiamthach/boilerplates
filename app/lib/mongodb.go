package lib

import (
	"context"
	"fmt"
	"go-microservices/app/config"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoDB *mongo.Database
)

func GetMongoDBClient() *mongo.Database {
	if mongoDB == nil {
		connect()
	}
	return mongoDB
}

func init() {
	connect()
}
func connect() {
	start := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Minute)
	defer cancel()
	conn := config.Get().MongoDb.Connection
	clientOptions := options.Client().ApplyURI(conn)
	cli, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println("connect mongodb error : " + err.Error())
	}
	errorPing := cli.Ping(context.TODO(), nil)
	if errorPing != nil {
		log.Println("mongodb client ping error ", errorPing)
	} else {
		log.Println("connected to mongodb!")
	}
	if err != nil {
		log.Println("error when init *mongo.Database : " + err.Error())
	}
	mongoDB = cli.Database(config.Get().MongoDb.DbName)
	elapsed := time.Since(start)
	log.Println("database connection took: " + fmt.Sprintf("%s", elapsed))
}
