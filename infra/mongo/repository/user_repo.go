package repository

import (
	"context"
	"go-gin/config"
	"go-gin/infra/mongo/dto"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const COLLECTION_USERS string = "users"

type IUserRepository interface {
	Gets() ([]dto.UserDto, error)
	CreateUser(user dto.UserDto) (dto.UserDto, error)
}

type UserRepository struct {
	BaseRepository
	IUserRepository
	config config.AppEnv
	log    *logrus.Logger
}

func (repo *UserRepository) Init(
	config config.AppEnv,
	log *logrus.Logger,
) {
	repo.config = config
	repo.log = log
	repo.DbName = config.MongoDb.DbName
	repo.Database = repo.GetDb()
	repo.Client = repo.GetClient()
	repo.Collection = COLLECTION_USERS
	repo.MyCollection = repo.GetCollection(COLLECTION_USERS)
}

func (repo *UserRepository) Gets() ([]dto.UserDto, error) {
	sort := bson.D{{Key: "name", Value: 1}}
	opts := options.Find().SetSort(sort)
	cursor, err := repo.MyCollection.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		return nil, err
	}

	data := make([]dto.UserDto, 0)
	if err = cursor.All(context.Background(), &data); err != nil {
		return nil, err
	}

	return data, nil
}

func (repo *UserRepository) CreateUser(user dto.UserDto) (dto.UserDto, error) {
	_, err := repo.MyCollection.InsertOne(context.Background(), user)
	if err != nil {
		return user, err
	}

	return user, nil
}
