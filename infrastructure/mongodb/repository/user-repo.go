package repository

import "go-microservices/infrastructure/mongodb/dtos"

type IUserRepository interface {
	FindAll() ([]*dtos.UserDTO, error)
	FindById(id string) (*dtos.UserDTO, error)
	FindByUsername(username string) (*dtos.UserDTO, error)
	FindByEmail(email string) (*dtos.UserDTO, error)
	Insert(user *dtos.UserDTO) (*dtos.UserDTO, error)
	Update(user *dtos.UserDTO) (*dtos.UserDTO, error)
}

// implement the interface
type UserRepository struct {
	IUserRepository
	BaseRepository
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		BaseRepository: BaseRepository{
			DbName:       "go-microservices",
			DbCollection: "users",
		},
	}
}
