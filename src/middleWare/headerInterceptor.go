package middleWare

import (
	"UploadFileProject/src/global"
	"UploadFileProject/src/utils/resp"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HeaderInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求头信息
		headerValue := c.GetHeader(global.Authorization)

		// 	用户token 未登陆就直接拦截退出
		if headerValue == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, resp.NewResultCont(http.StatusUnauthorized, "please login", ""))
			return
		}

		// 将请求头信息存储在上下文中
		c.Set(global.Organization, headerValue)
		c.Next()
	}
}
