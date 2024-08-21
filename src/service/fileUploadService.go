package service

import (
	"UploadFileProject/src/entity/bo"
	"UploadFileProject/src/entity/dto"
	"UploadFileProject/src/global/enum"
	"UploadFileProject/src/mapper"
	"UploadFileProject/src/mq"
	"encoding/json"
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

	organizationUuid := fileUploadDTO.OrganizationUuid
	fileUuidAndSuffix := strings.Join(fileSlice, ".")

	// 获取organization对象
	FuOrganizationBO := mapper.FuOrganizationBOMapperImpl.SelectFuOrganization(organizationUuid)

	// 拼接路径并将文件存入本地
	savePath := filepath.Join(UpLoadsPath, FuOrganizationBO.OrgPath, fileSuffix, fileUuidAndSuffix)
	if err := fileUploadDTO.Context.SaveUploadedFile(fileUploadDTO.File, savePath); err != nil {
		resultStr = "Unable to save the file"
		statusCode = http.StatusInternalServerError
		return
	}

	logService.Infof("file has been saved :%s", savePath)

	// 将文件信息插入到表中
	mapper.FuFileBOMapperImpl.InsertFuFile(&bo.FuFileBO{
		FileUuid:         fileUuidStr,
		FileSuffix:       fileSuffix,
		FileOriginalName: fileOriginalName,
		OrgId:            FuOrganizationBO.Id,
		OssPath:          enum.OssPathDefault.ToString(),
		IfUploadOss:      enum.NoneUploadOss.ToInt32(),
		LocalGroup:       UpLoadsPath,
	})

	// 创建消息
	message := mq.NewMessage(&dto.UpLoadSingleFileOSSMqDTO{
		OrganizationUuid: FuOrganizationBO.OrgUuid,
		FileSuffix:       fileSuffix,
		FileUuid:         fileUuidStr,
		GroupId:          UpLoadsPath,
	}, "UpLoadSingleFileOSSMqDTO", enum.TaskSingleFileUpload.ToInt64())

	jsonData, err := json.Marshal(message)
	if err != nil {
		logService.Fatalf("Error occurred during marshaling. Error: %s", err.Error())
	}
	// 生产者发送
	mq.Producer(jsonData)

	resultStr = fileUuidStr
	statusCode = http.StatusOK
	return
}
