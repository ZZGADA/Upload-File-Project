package service

import (
	"UploadFileProject/src/entity/bo"
	"UploadFileProject/src/entity/dto"
	"UploadFileProject/src/entity/vo"
	"UploadFileProject/src/global"
	"UploadFileProject/src/global/enum"
	"UploadFileProject/src/mapper"
	"UploadFileProject/src/oss"
	"UploadFileProject/src/utils/process"
	"UploadFileProject/src/utils/resp"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"sync"
)

type FileDownloadService struct{}

var FileDownloadServiceImpl = &FileDownloadService{}

// DownloadSingleFile 单文件下载
func (fileDownloadService *FileDownloadService) DownloadSingleFile(fileDownloadDTO *dto.FileDownloadDTO, result *resp.Result) {
	orgUuid := fileDownloadDTO.OrganizationUuid
	fileUuid := fileDownloadDTO.FileUuid

	fuFileBO := mapper.FuFileBOMapperImpl.GetOneFile(fileUuid)

	resData := FileDownloadServiceImpl.download(fuFileBO, orgUuid, fileUuid)
	result.Success(resData)

	return
}

func (fileDownloadService *FileDownloadService) DownloadBatchFile(
	fileBatchLoadDTO *dto.FileBatchLoadDTO,
	orgStr string,
	result *resp.Result) {
	fileUuidList := fileBatchLoadDTO.FileUuidList
	if len(fileUuidList) == 0 {
		// 传入参数为空
		result.Success("请选择文件信息")
		return
	}

	fileList := mapper.FuFileBOMapperImpl.GetBatchFileInformation(fileUuidList)
	ch := make(chan string, len(fileList))
	var wg sync.WaitGroup

	for _, fileInfo := range fileList {
		fuFileBo := &bo.FuFileBO{
			OssBucket:        fileInfo.OssBucket,
			OssPath:          fileInfo.OssPath,
			FileUuid:         fileInfo.FileUuid,
			FileOriginalName: fileInfo.FileName,
			FileSuffix:       fileInfo.FileSuffix,
			LocalGroup:       fileInfo.LocalGroup,
			IfUploadOss:      fileInfo.IfUploadOss,
		}
		wg.Add(1)
		go func(fuFileBo *bo.FuFileBO) {
			defer wg.Done()
			fileDownloadVO := FileDownloadServiceImpl.download(fuFileBo, orgStr, fuFileBo.FileUuid)
			ch <- fileDownloadVO.FileData
		}(fuFileBo)
	}

	// 新建协程 等待协程全部结束
	go func() {
		wg.Wait()
		close(ch)
	}()

	var fileBatchDownloadVo = &vo.FileBatchDownloadVO{OrganizationId: orgStr}
	// 读取channel
	for message := range ch {
		fmt.Println("Received:", message)
		fileBatchDownloadVo.FilesPath = append(fileBatchDownloadVo.FilesPath, message)
	}
	result.Success(fileBatchDownloadVo)
}

func (fileDownloadService *FileDownloadService) download(fuFileBO *bo.FuFileBO, orgUuid string, fileUuid string) *vo.FileDownloadVO {
	bucketName := fuFileBO.OssBucket
	fileName := process.FileNameJoinSuffix(fuFileBO.FileUuid, fuFileBO.FileSuffix)

	localDownloadPath := filepath.Join(
		global.DownLoadsPath,
		orgUuid,
		fuFileBO.FileSuffix,
		fileName)

	localUploadPath := filepath.Join(
		global.UpLoadsPath,
		orgUuid,
		fuFileBO.FileSuffix,
		fileName)

	var fileDownloadVO = &vo.FileDownloadVO{
		OrganizationUuid: orgUuid,
		FileUuid:         fileUuid,
	}

	if fuFileBO.IfUploadOss == enum.UploadOss.ToInt32() {
		// 如果成功上传了OSS就从OSS获取对象
		// 优先查询本地是否有已经下载的对象 ， 如果有就直接读本地 否则从还是要从OSS下载
		fileDownloadVO.FileData = localDownloadPath
		if process.FileExists(localDownloadPath) {
			logService.Info("本地有下载缓存，直接从本地下载")
			//readFileContext(localDownloadPath, result.Ctx)
			return fileDownloadVO
		}

		oss.OssServerImpl.DownLoadSingleFIle(fuFileBO.OssPath, localDownloadPath, bucketName)
		//readFileContext(localDownloadPath, result.Ctx)
		return fileDownloadVO
	} else {
		// 文件还没有上传OSS 直接读本地
		logService.Info("文件没有上传OSS，从uploads文件读取")
		//readFileContext(localUploadPath, result.Ctx)
		fileDownloadVO.FileData = localUploadPath
		return fileDownloadVO
	}
}

// readFile 读文件
/*
// if err := readFile(localUploadPath, &fileDownloadVO.FileData); err != nil {
//	 result.Failed(http.StatusInternalServerError, err)
// }
// fileDownloadDTO.Context.File(localDownloadPath)
*/
func readFile(path string, fileData *interface{}) error {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	if err2 := json.Unmarshal(fileContent, fileData); err2 != nil {
		return err
	}
	return nil
}

func readFileContext(path string, context *gin.Context) {
	context.File(path)
}
