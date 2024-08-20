package controller

import (
	"UploadFileProject/src/entity/dto"
	"UploadFileProject/src/global"
	"UploadFileProject/src/middleWare"
	"UploadFileProject/src/service"
	"UploadFileProject/src/utils/resp"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func fileSingleUploadController(router *gin.RouterGroup) {
	router.POST("/singleFile", middleWare.SingleFileInterceptor(), fileSingleUpload)
}

func fileSingleUpload(c *gin.Context) {
	result := resp.NewResult(c)

	// 获取请求头信息
	organizationUuid, _ := c.Get(global.Organization)
	organizationUuidStr := fmt.Sprintf("%v", organizationUuid)

	// 核心的业务代码
	file, _ := c.FormFile(global.SingleFileName)
	fileUploadDTO := dto.NewFileUploadDTO(file, organizationUuidStr, c)
	resultData, statusCode := service.FileUploadServiceImpl.UploadSingleFileService(fileUploadDTO)

	if statusCode != http.StatusOK {
		logController.Warnf("fileSingleUpload failed, please checking，msg：%v ", resultData)
		result.Failed(statusCode, resultData)
		return
	}
	result.Success(&resultData)
}
