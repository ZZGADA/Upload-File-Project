package config

import (
	"UploadFileProject/src/global"
	"UploadFileProject/src/middleWare"
	"github.com/gin-gonic/gin"
	"net/http"
)

var ServerAllConfig = &Server{}
var Router *gin.Engine

func initServer() {
	ServerAllConfig = &ProjectConfig.Server

	gin.SetMode(ServerAllConfig.GinMod)
	Router = gin.New()
	Router.Use(middleWare.LoggerMiddleware(global.Log), gin.Recovery())

	frontEnd()
}

func frontEnd() {
	Router.LoadHTMLFiles("./templates/uploadFrontEnd.html")
	Router.GET("/uploadFrontEnd", func(c *gin.Context) {
		c.HTML(http.StatusOK, "uploadFrontEnd.html", nil)
	})
}
