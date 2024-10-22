package controller

import (
	"UploadFileProject/src/utils/resp"
	"github.com/gin-gonic/gin"
)

// InitCheckHealthController 初始化接口配置
func checkHealthController(router *gin.RouterGroup) {
	router.GET("", checkHealth) // 传入函数指针
}

// checkHealth // 健康检测
func checkHealth(c *gin.Context) {
	result := resp.NewResult(c)
	result.Success("health")
}
