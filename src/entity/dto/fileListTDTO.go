package dto

import "time"

type FileListTDTO struct {
	FileName    string    `gorm:"column:fileName"`
	FileSuffix  string    `gorm:"column:fileSuffix"`
	FileUuid    string    `gorm:"column:fileUuid"`
	OrgName     string    `gorm:"column:orgName"`
	OrgUuid     string    `gorm:"column:orgUuid"`
	LocalGroup  string    `gorm:"column:localGroup"`
	IfUploadOss int32     `gorm:"column:ifUploadOss"`
	OssBucket   string    `gorm:"column:ossBucket"`
	OssPath     string    `gorm:"column:ossPath"`
	CreateTime  time.Time `gorm:"column:createTime"`
	UpdateTime  time.Time `gorm:"column:updateTime"`
}
