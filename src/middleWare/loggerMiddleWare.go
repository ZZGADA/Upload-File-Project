package middleWare

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

// LoggerMiddleware 是一个 Gin 中间件，用于使用 Logrus 记录 HTTP 请求日志
func LoggerMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		//
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()

		// 日志格式
		logInfoStr := fmt.Sprintf("%d----%s----[%s]:%s   latencyTime:%v", statusCode, reqMethod, clientIP, reqUri, latencyTime)
		logger.Info(logInfoStr)
	}
}
