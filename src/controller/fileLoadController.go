package controller

import (
	"UploadFileProject/src/entity/dto"
	"UploadFileProject/src/global"
	"UploadFileProject/src/middleWare"
	"UploadFileProject/src/service"
	"UploadFileProject/src/utils/process"
	"UploadFileProject/src/utils/resp"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func fileLoadController(router *gin.RouterGroup) {
	router.POST("/uploadSingle", middleWare.SingleFileInterceptor(), fileSingleUpload)
	router.POST("/downloadSingle", downloadOneFile)
	router.POST("/downloadBatch", downLoadBatchFile)
}

func downloadOneFile(c *gin.Context) {
	result := resp.NewResult(c)

	// 解析请求体
	var fileDownloadDTO = dto.NewFileDownloadDTOWithContext(c)
	if err := c.ShouldBindBodyWithJSON(fileDownloadDTO); err != nil {
		logController.Errorf("json 解析失败，%#v", err)
	}
	logController.Infof("get download dto ,%#v", fileDownloadDTO)

	service.FileDownloadServiceImpl.DownloadSingleFile(fileDownloadDTO, result)
}

func downLoadBatchFile(c *gin.Context) {
	result := resp.NewResult(c)

	var fileBatchLoadDTO = &dto.FileBatchLoadDTO{}
	if err := process.JsonFormat(c, fileBatchLoadDTO, result); err != nil {
		return
	}
	organizationUuid, _ := c.Get(global.Organization)
	organizationUuidStr := fmt.Sprintf("%v", organizationUuid)

	service.FileDownloadServiceImpl.DownloadBatchFile(fileBatchLoadDTO, organizationUuidStr, result)
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
