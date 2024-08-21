package service

import (
	"UploadFileProject/src/entity/dto"
	"UploadFileProject/src/mapper"
	"UploadFileProject/src/oss"
	"fmt"

	"path/filepath"
)

type UploadFileOssConsumer struct{}

var UploadFileOssServiceImpl = &UploadFileOssConsumer{}

// UploadSingleFileOSS  // 单文件上传OSS
func (fileUploadOSS *UploadFileOssConsumer) UploadSingleFileOSS(singleFileMessageDTO *dto.UpLoadSingleFileOSSMqDTO) {
	groupName := singleFileMessageDTO.GroupId
	orgUuid := singleFileMessageDTO.OrganizationUuid
	fileUuid := singleFileMessageDTO.FileUuid
	fileSuffix := singleFileMessageDTO.FileSuffix

	filePath := filepath.Join(
		groupName,
		orgUuid,
		fileSuffix,
		fileUuid)
	fileFullPath := fmt.Sprintf("%s.%s", filePath, fileSuffix)

	// 上传文件并更新DB
	oss.OssServerImpl.UploadSingleFile(fileFullPath, fileFullPath)
	mapper.FuFileBOMapperImpl.UpdateOssPath(fileUuid, fileFullPath, oss.BucketName)
}
