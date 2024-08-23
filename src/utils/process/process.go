package process

import (
	"UploadFileProject/src/global"
	"UploadFileProject/src/utils/resp"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
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

// CheckFileExist  检查文件是否存在
func CheckFileExist(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// 如果有err 表示文件不存在
		return false
	}
	// 否则表示文件存在
	return true
}

// fileExists // 判断file是否存在 如果存在返回true
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

// CheckFileDirExist 查看文件是否存在 不存在则创建文件夹
func CheckFileDirExist(filePath string) {
	dirPath := filepath.Dir(filePath)
	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// 如果文件不存在，检查文件夹是否存在
		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			// 如果文件夹不存在，则创建文件夹
			if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
				global.Log.Error(err)
				return
			}
		}
	}
}
