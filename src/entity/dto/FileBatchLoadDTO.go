package dto

// FileBatchLoadDTO // FileBatchLoadDTO 批量文件导入
type FileBatchLoadDTO struct {
	FileUuidList []string `json:"fileUuidList"`
}
