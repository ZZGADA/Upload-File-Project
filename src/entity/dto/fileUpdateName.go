package dto

// FileUpdateName 更新文件名字
type FileUpdateName struct {
	FileUuid string `json:"fileUuid"`
	NewName  string `json:"newName"`
}
