package vo

// FileBatchDownloadVO // 批量文件到处返回对象
type FileBatchDownloadVO struct {
	OrganizationId string   `json:"organizationId"`
	FilesPath      []string `json:"filesPath"`
}
