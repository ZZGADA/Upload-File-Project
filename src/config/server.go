package config

import (
	"github.com/gin-gonic/gin"
)

var ServerAllConfig = &Server{}
var Router *gin.Engine

func initServer() {
	ServerAllConfig = &ProjectConfig.Server
	gin.SetMode(ServerAllConfig.GinMod)
	Router = gin.New()

	Router.Use(LoggerMiddleware(Log))
	Router.Use(gin.Recovery())
	Router.Use(gin.Logger())

	// /search?query=gin  路由携带参数
	Router.GET("/search", func(c *gin.Context) {
		query := c.Query("query")
		more := c.Query("more")
		Log.Warn("test-test-test-test-test-test-test-test")
		Log.Info("Info-Info-Info-Info-Info-Info-Info-Info-")
		c.String(200, "Search query is "+query+" more is "+more)
	})

}
