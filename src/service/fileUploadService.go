package service

import (
	"UploadFileProject/src/entity/bo"
	"UploadFileProject/src/entity/dto"
	"UploadFileProject/src/mapper"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"path/filepath"
	"strings"
)

const UpLoadsPath = "uploads"

type FileUploadService struct {
}

var FileUploadServiceImpl = &FileUploadService{}

// UploadSingleFileService 将前端文件存入本地
func (fileUpload *FileUploadService) UploadSingleFileService(fileUploadDTO *dto.FileUploadDTO) (resultStr interface{}, statusCode int) {
	// 检查文件类型是否为PDF
	logService.Infof("file name is %s", fileUploadDTO.File.Filename)
	if filepath.Ext(fileUploadDTO.File.Filename) != ".pdf" {
		resultStr = "Only PDF files are allowed"
		statusCode = http.StatusBadRequest
		return
	}

	resultStr, statusCode = fileUpload.saveFile(fileUploadDTO)
	return

}

// 将文件具体存入到本地的具体方法
func (fileUpload *FileUploadService) saveFile(fileUploadDTO *dto.FileUploadDTO) (resultStr string, statusCode int) {
	// 保存文件到服务器
	fileUuid := uuid.NewV1()
	fileUuidStr := fileUuid.String()
	fileSlice := strings.Split(fileUploadDTO.File.Filename, ".")

	fileOriginalName := fileSlice[0]
	fileSuffix := fileSlice[1]
	fileSlice[0] = fileUuidStr

	organizationUuid := fmt.Sprintf("%v", fileUploadDTO.OrganizationUuid)
	fileUuidAndSuffix := strings.Join(fileSlice, ".")

	// 获取organization对象
	FuOrganizationBO := mapper.FuOrganizationBOMapperImpl.SelectOrganization(organizationUuid)

	// 拼接路径并将文件存入本地
	savePath := filepath.Join(UpLoadsPath, FuOrganizationBO.OrgPath, fileSuffix, fileUuidAndSuffix)
	if err := fileUploadDTO.C.SaveUploadedFile(fileUploadDTO.File, savePath); err != nil {
		resultStr = "Unable to save the file"
		statusCode = http.StatusInternalServerError
		return
	}

	logService.Infof("file has been saved :%s", savePath)

	// 将文件信息插入到表中
	mapper.FuFileBOMapperImpl.InsertFile(&bo.FuFileBO{
		FileUuid:         fileUuidStr,
		FileSuffix:       fileSuffix,
		FileOriginalName: fileOriginalName,
		OrgId:            FuOrganizationBO.Id,
		OssPath:          "",
		IfUploadOss:      0,
	})

	resultStr = fileUuidStr
	statusCode = http.StatusOK

	return
}
