package vo

// FileMultiUploadVO // 批量文件上传返回结果
type FileMultiUploadVO struct {
	SuccessFilesUuid []string `json:"successFilesUuid"`
	FailedFilesUuid  []string `json:"failedFilesUuid"`
}
