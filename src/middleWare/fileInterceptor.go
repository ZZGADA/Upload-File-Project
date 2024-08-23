package middleWare

import (
	"UploadFileProject/src/global"
	"UploadFileProject/src/utils/resp"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SingleFileInterceptor 单文件上传拦截器
func SingleFileInterceptor() gin.HandlerFunc {
	return func(context *gin.Context) {
		_, err := context.FormFile(global.SingleFileName)
		if err != nil {
			context.AbortWithStatusJSON(
				http.StatusBadRequest,
				resp.NewResultCont(
					http.StatusBadRequest,
					"error , didn't get any file,please upload at least single file",
					"",
				))
			return
		}
		context.Next()
	}
}

// MultiFileInterceptor 多文件上传拦截器
func MultiFileInterceptor() gin.HandlerFunc {
	return func(context *gin.Context) {
		form, _ := context.MultipartForm()
		if len(form.File) == 0 {
			context.AbortWithStatusJSON(
				http.StatusBadRequest,
				resp.NewResultCont(
					http.StatusBadRequest,
					"error , didn't get any file,please upload at least single file",
					"",
				))
			return
		}

		context.Next()
	}
}
