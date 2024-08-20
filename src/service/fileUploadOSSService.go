package service

import "UploadFileProject/src/entity/dto"

type UploadFileOSSService struct{}

var UploadFileOSSServiceImpl = &UploadFileOSSService{}

// UploadSingleFileOSS  // 单文件上传OSS
func (fileUploadOSS *UploadFileOSSService) UploadSingleFileOSS(singleFileMessageDTO *dto.UpLoadSingleFileOSSMqDTO) {
	logService.Infof("断言成功，成功进入UploadSingleFileOSSService,%#v", singleFileMessageDTO)
}
