package controllers

import (
	"go-gin/config"
	"go-gin/infra/mongo/repository"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type IUserController interface {
	Init(config config.AppEnv, log *logrus.Logger, userRepo repository.IUserRepository)
	GetUsers() gin.HandlerFunc
}

type UserController struct {
	IUserController
	userRepo repository.IUserRepository
	config   config.AppEnv
	log      *logrus.Logger
}

func (c *UserController) Init(config config.AppEnv, log *logrus.Logger, userRepo repository.IUserRepository) {
	c.config = config
	c.log = log
	c.userRepo = userRepo
}

func (c *UserController) GetUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		users, err := c.userRepo.Gets()
		if err != nil {
			c.log.Error(err)
			ctx.JSON(500, gin.H{
				"message": "error when get users",
			})
			return
		}
		ctx.JSON(200, users)
	}
}
