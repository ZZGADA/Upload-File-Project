package service

import (
	"UploadFileProject/src/entity/bo"
	"UploadFileProject/src/entity/dto"
	"UploadFileProject/src/entity/vo"
	"UploadFileProject/src/global"
	"UploadFileProject/src/global/enum"
	"UploadFileProject/src/mapper"
	"UploadFileProject/src/mq"
	"UploadFileProject/src/utils/resp"
	"encoding/json"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
	"sync"
)

type FileUploadService struct {
}

var FileUploadServiceImpl = &FileUploadService{}

// UploadMultiFileService 多任务上传
func (fileUpload *FileUploadService) UploadMultiFileService(files []*multipart.FileHeader, context *gin.Context, result *resp.Result, orgUuid string) {
	successChan := make(chan string, len(files))
	failChan := make(chan string, len(files))
	var wg sync.WaitGroup
	var fileProcess = len(files)

	for _, file := range files {
		wg.Add(1)
		go func(f *multipart.FileHeader) {
			defer wg.Done()
			fileUploadDTO := &dto.FileUploadDTO{
				OrganizationUuid: orgUuid,
				File:             f,
				Context:          context,
			}
			resultStr, resultCode := fileUpload.saveFile(fileUploadDTO)
			if resultCode != http.StatusOK {
				failChan <- resultStr
			} else {
				successChan <- resultStr
			}
		}(file)
	}

	var failedFilesUuid = make([]string, 0, len(files)*10)
	var successFilesUuid = make([]string, 0, len(files)*10)

	// 处理成功和失败的文件
	for {
		select {
		case successFile, ok := <-successChan:
			if ok {
				logService.Infof("Successfully uploaded: %s\n", successFile)
				successFilesUuid = append(successFilesUuid, successFile)
				fileProcess--
			}
		case failFile, ok := <-failChan:
			if ok {
				logService.Error("Failed to upload: %s\n", failFile)
				failedFilesUuid = append(failedFilesUuid, failFile)
				fileProcess--
			}
		}
		// 兜底
		if len(successChan) == 0 && len(failChan) == 0 && fileProcess == 0 {
			break
		}
	}
	result.Success(vo.FileMultiUploadVO{
		SuccessFilesUuid: successFilesUuid,
		FailedFilesUuid:  failedFilesUuid,
	})

	wg.Wait()
	close(successChan)
	close(failChan)
	logService.Info("channel 资源释放")
	return
}

// UploadSingleFileService 将前端文件存入本地
func (fileUpload *FileUploadService) UploadSingleFileService(fileUploadDTO *dto.FileUploadDTO) (resultStr interface{}, statusCode int) {
	// 检查文件类型是否为PDF
	//logService.Infof("file name is %s", fileUploadDTO.File.Filename)
	//if filepath.Ext(fileUploadDTO.File.Filename) != ".pdf" {
	//	resultStr = "Only PDF files are allowed"
	//	statusCode = http.StatusBadRequest
	//	return
	//}

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
	savePath := filepath.Join(global.UpLoadsPath, FuOrganizationBO.OrgPath, fileSuffix, fileUuidAndSuffix)
	if err := fileUploadDTO.Context.SaveUploadedFile(fileUploadDTO.File, savePath); err != nil {
		resultStr = fileUuidStr
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
		LocalGroup:       global.UpLoadsPath,
	})

	// 创建消息
	message := mq.NewMessage(&dto.UpLoadSingleFileOSSMqDTO{
		OrganizationUuid: FuOrganizationBO.OrgUuid,
		FileSuffix:       fileSuffix,
		FileUuid:         fileUuidStr,
		GroupId:          global.UpLoadsPath,
	}, "UpLoadSingleFileOSSMqDTO", enum.TaskSingleFileUpload.ToInt64())

	jsonData, err := json.Marshal(message)
	if err != nil {
		logService.Fatalf("Error occurred during marshaling. Error: %s", err.Error())
		resultStr = fileUuidStr
		statusCode = http.StatusInternalServerError
	}
	// 生产者发送
	mq.Producer(jsonData)

	resultStr = fileUuidStr
	statusCode = http.StatusOK
	return
}
