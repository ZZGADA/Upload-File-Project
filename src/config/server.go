package config

import (
	"UploadFileProject/src/global"
	"UploadFileProject/src/utils/resp"
	"github.com/gin-gonic/gin"
	"net/http"
)

var ServerAllConfig = &Server{}
var Router *gin.Engine

func initServer() {
	ServerAllConfig = &ProjectConfig.Server

	gin.SetMode(ServerAllConfig.GinMod)
	Router = gin.New()
	Router.Use(LoggerMiddleware(Log))
	Router.Use(gin.Recovery())

	Router.LoadHTMLFiles("./templates/uploadFrontEnd.html")
	Router.GET("/uploadFrontEnd", func(c *gin.Context) {
		c.HTML(http.StatusOK, "uploadFrontEnd.html", nil)
	})
}

func HeaderInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求头信息
		headerValue := c.GetHeader(global.Authorization)

		if headerValue == "" {
			result := resp.NewResult(c)
			result.Failed(http.StatusForbidden, "please login")
			return
		}

		// 将请求头信息存储在上下文中
		c.Set(global.Organization, headerValue)
	}
}
