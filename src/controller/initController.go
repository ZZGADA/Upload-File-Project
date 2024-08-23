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
	fileUploadGroup := router.Group("/file")
	{
		// 路由组初始化配置
		// 从上往下：上传下载controller、文件查询controller
		fileUploadGroup.Use(middleWare.HeaderInterceptor())

		fileLoadController(fileUploadGroup)
		fileController(fileUploadGroup)
	}
}
