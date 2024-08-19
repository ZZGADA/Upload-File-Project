package controller

import (
	"UploadFileProject/src/entity/dto"
	"UploadFileProject/src/global"
	"UploadFileProject/src/service"
	"UploadFileProject/src/utils/resp"
	"github.com/gin-gonic/gin"
	"net/http"
)

func initFileSingleUploadController(router *gin.RouterGroup) {
	service.InitService()
	router.POST("/singleFile", fileSingleUpload)
}

func fileSingleUpload(c *gin.Context) {
	result := resp.NewResult(c)

	HeaderInterceptor(c)
	organizationUuid, _ := c.Get(global.Organization)
	logController.Infof("organization id is %#v", organizationUuid)

	// 核心的业务代码
	file, err := c.FormFile("singleFile")
	if err != nil {
		result.Failed(http.StatusBadRequest, "error , didn't get any file")
		return
	}
	fileUploadDTO := &dto.FileUploadDTO{File: file, OrganizationUuid: organizationUuid, C: c}
	resultData, statusCode := service.FileUploadServiceImpl.UploadSingleFileService(fileUploadDTO)

	if statusCode != http.StatusOK {
		logController.Warnf("fileSingleUpload failed, please checking，msg：%v ", resultData)
		result.Failed(statusCode, resultData)
		return
	}
	result.Success(&resultData)
}

func HeaderInterceptor(c *gin.Context) {
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
