package repository

import (
	"github.com/golobby/container/v3"
	"go.mongodb.org/mongo-driver/mongo"
)

type BaseRepository struct {
	DbName       string
	Client       *mongo.Client
	Database     *mongo.Database
	MyCollection *mongo.Collection
	Collection   string
}

func (b BaseRepository) GetClient() *mongo.Client {
	if b.Client != nil {
		return b.Client
	}

	container.Resolve(&b.Client)
	return b.Client
}

func (b BaseRepository) GetDb() *mongo.Database {
	if b.Database != nil {
		return b.Database
	}

	b.Database = b.GetClient().Database(b.DbName)
	return b.Database
}

func (b BaseRepository) GetCollection(name string) *mongo.Collection {
	collection := b.GetDb().Collection(name)
	return collection
}
