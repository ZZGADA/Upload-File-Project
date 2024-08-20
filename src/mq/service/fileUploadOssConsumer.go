package service

import (
	"UploadFileProject/src/entity/dto"
)

type UploadFileOssConsumer struct{}

var UploadFileOssServiceImpl = &UploadFileOssConsumer{}

// UploadSingleFileOSS  // 单文件上传OSS
func (fileUploadOSS *UploadFileOssConsumer) UploadSingleFileOSS(singleFileMessageDTO *dto.UpLoadSingleFileOSSMqDTO) {
	logCs.Infof("断言成功，成功进入UploadSingleFileOSSService,%#v", singleFileMessageDTO)

}
