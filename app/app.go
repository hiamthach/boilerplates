package app

import (
	"go-gin/config"
	"go-gin/router"

	"github.com/sirupsen/logrus"
)

func Start(config *config.AppEnv, log *logrus.Logger) {

	r := router.InitRouter(
		config,
		log,
	)

	r.Run(config.Server.Port)
}
