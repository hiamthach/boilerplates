package router

import (
	"go-gin/app/controllers"

	"github.com/gin-gonic/gin"
	"github.com/golobby/container/v3"
)

func userRouter(router *gin.RouterGroup) {
	users := router.Group("/users")

	var userController controllers.IUserController
	err := container.Resolve(&userController)

	if err != nil {
		panic(err)
	}

	users.GET("/", userController.GetUsers())
}
