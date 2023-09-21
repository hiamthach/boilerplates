package repository

import "go.mongodb.org/mongo-driver/mongo"

type IBaseRepository interface {
	FindAll() ([]interface{}, error)
	FindById(id string) (interface{}, error)
	Insert(data interface{}) (interface{}, error)
	Update(data interface{}) (interface{}, error)
	Delete(id string) (interface{}, error)
}

type BaseRepository struct {
	DbName       string
	DbCollection string
	Client       *mongo.Client
	Database     *mongo.Database
	MyCollection *mongo.Collection

	IBaseRepository
}
