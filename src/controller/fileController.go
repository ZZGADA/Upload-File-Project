package controller

import (
	"UploadFileProject/src/entity/dto"
	"UploadFileProject/src/global"
	"UploadFileProject/src/service"
	"UploadFileProject/src/utils/process"
	"UploadFileProject/src/utils/resp"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func fileController(router *gin.RouterGroup) {
	router.POST("/searchList", searchList)
	router.POST("/updateFileName", updateFileName)
	router.GET("/delete", deleteFile)
	router.POST("/getInfoList", getInfoList)
}

const FileUuid = "fileUuid"

func deleteFile(context *gin.Context) {
	result := resp.NewResult(context)

	fileUuid := context.Query(FileUuid)
	if fileUuid == "" {
		result.Failed(http.StatusBadRequest, "请携带file uuid 参数")
		return
	}
	var fileDeleteDTO = &dto.FileDeleteDTO{FileUuid: fileUuid}

	service.FileServiceImpl.DeleteFile(fileDeleteDTO, result)
}

func updateFileName(context *gin.Context) {
	result := resp.NewResult(context)

	var fileUpdateNameDTO = &dto.FileUpdateName{}
	if err := process.JsonFormat(context, fileUpdateNameDTO, result); err != nil {
		return
	}

	service.FileServiceImpl.UpdateFileName(fileUpdateNameDTO, result)
}

func searchList(context *gin.Context) {
	result := resp.NewResult(context)

	var fileSearchDTO = &dto.FileSearchDTO{}
	if err := process.JsonFormat(context, fileSearchDTO, result); err != nil {
		return
	}

	service.FileServiceImpl.GetFileList(fileSearchDTO, result)
}

func getInfoList(context *gin.Context) {
	result := resp.NewResult(context)

	var fileInfoListDTO = &dto.FileInfoListDTO{}
	if err := process.JsonFormat(context, fileInfoListDTO, result); err != nil {
		return
	}

	// 获取请求头信息
	organizationUuid, _ := context.Get(global.Organization)
	organizationUuidStr := fmt.Sprintf("%v", organizationUuid)

	service.FileServiceImpl.GetFileInfoList(fileInfoListDTO, organizationUuidStr, result)
}
