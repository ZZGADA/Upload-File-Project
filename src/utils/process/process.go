package process

import (
	"UploadFileProject/src/global"
	"UploadFileProject/src/utils/resp"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FileNameJoinSuffix(fileName string, fileSuffix string) string {
	return fmt.Sprintf("%s.%s", fileName, fileSuffix)
}

// JsonFormat // json对象解析
func JsonFormat(context *gin.Context, obj any, result *resp.Result) error {
	if err := context.ShouldBind(obj); err != nil {
		global.Log.Errorf("json 解析失败，%#v", err)
		result.Failed(http.StatusInternalServerError, "json 解析失败 格式化错误")
		return err
	}
	return nil
}
