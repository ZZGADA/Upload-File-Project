package controller

import (
	"UploadFileProject/src/utils/resp"
	"github.com/gin-gonic/gin"
)

// InitCheckHealthController 初始化接口配置
func initCheckHealthController(router *gin.RouterGroup) {
	router.GET("", checkHealth) // 传入函数指针
}

// checkHealth // 健康检测
func checkHealth(c *gin.Context) {
	result := resp.NewResult(c)

	// 核心的业务代码
	resultString := "health"

	result.Success(&resultString)
}
