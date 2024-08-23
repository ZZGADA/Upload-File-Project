package dto

// UpLoadOSSMqDTO // Mq 传输文件的消息体
type UpLoadSingleFileOSSMqDTO struct {
	OrganizationUuid string `json:"organizationId" mapstructure:"organizationId"`
	FileUuid         string `json:"fileUuid" mapstructure:"fileUuid"`
	GroupId          string `json:"groupId" mapstructure:"groupId"` // 分组ID 最高级磁盘区域
	FileSuffix       string `json:"fileSuffix" mapstructure:"fileSuffix"`
}
