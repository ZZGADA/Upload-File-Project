package middleWare

import (
	"UploadFileProject/src/global"
	"UploadFileProject/src/utils/resp"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
