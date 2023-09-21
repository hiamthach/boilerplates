package router

import (
	"go-gin/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func InitRouter(
	config *config.AppEnv,
	log *logrus.Logger,
) *gin.Engine {
	router := gin.New()

	router.Use(cors.New(
		cors.Config{
			AllowAllOrigins: config.Server.Cors == "*", // true
			AllowMethods:    []string{config.Server.CorsMethod},
			AllowHeaders:    []string{config.Server.CorsHeader},
		},
	))

	router.Use(gin.Recovery())

	apiRouter := router.Group("/api/v1")

	{
		userRouter(apiRouter)
	}

	apiRouter.GET("/ping", func(c *gin.Context) {
		log.Print("Ping")

		c.JSON(200, gin.H{
			"message": "Pong",
		})
	})

	return router
}
