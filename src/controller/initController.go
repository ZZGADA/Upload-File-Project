package controller

import (
	"UploadFileProject/src/global"
	"UploadFileProject/src/middleWare"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var logController *logrus.Logger

func InitController(router *gin.Engine) {
	logController = global.Log

	checkHealthRouterGroup(router)
	fileUploadRouterGroup(router)
}

// 健康检测路由组
func checkHealthRouterGroup(router *gin.Engine) {
	checkHealth := router.Group("/checkHealth")
	{
		checkHealthController(checkHealth)
	}
}

// fileUploadRouterGroup 文件上传路由组
func fileUploadRouterGroup(router *gin.Engine) {
	fileUploadGroup := router.Group("/uploadFile")
	{
		// 路由组初始化配置
		fileUploadGroup.Use(middleWare.HeaderInterceptor())

		fileSingleUploadController(fileUploadGroup)
	}
}
