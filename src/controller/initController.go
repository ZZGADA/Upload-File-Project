package controller

import (
	"UploadFileProject/src/config"
	"UploadFileProject/src/mapper"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var logController *logrus.Logger

func InitController(router *gin.Engine) {
	logController = config.Log
	mapper.InitMapper()

	checkHealth := router.Group("/checkHealth")
	{
		initCheckHealthController(checkHealth)
	}

	fileUploadGroup := router.Group("/uploadFile")
	{
		initFileSingleUploadController(fileUploadGroup)
	}
}
