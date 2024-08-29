package dto

import "time"

type SynEsDTO struct {
	FileUuid         string    `gorm:"column:file_uuid;type:varchar(64);comment:文件唯一uuid;" json:"file_uuid"`                                                // 文件唯一uuid
	FileOriginalName string    `gorm:"column:file_original_name;type:varchar(64);comment:文件名;not null;" json:"file_original_name"`                          // 文件名
	FileSuffix       string    `gorm:"column:file_suffix;type:varchar(32);comment:文件名后缀;not null;" json:"file_suffix"`                                      // 文件名后缀
	LocalGroup       string    `gorm:"column:local_group;type:varchar(128);comment:本地存放地址的最大区域;not null;" json:"local_group"`                               // 本地存放地址的最大区域
	IfUploadOss      int32     `gorm:"column:if_upload_oss;type:tinyint;comment:是否上传oss: 0 没有上传 1上传;not null;default:0;" json:"if_upload_oss"`              // 是否上传oss: 0 没有上传 1上传
	OrgId            int64     `gorm:"column:org_id;type:bigint;comment:file文件所属的机构;not null;default:0;" json:"org_id"`                                     // file文件所属的机构
	OssPath          string    `gorm:"column:oss_path;type:varchar(512);comment:oss路径的存放地址;not null;" json:"oss_path"`                                      // oss路径的存放地址
	CreateTime       time.Time `gorm:"column:create_time;type:datetime;comment:创建时间;not null;default:CURRENT_TIMESTAMP;autoCreateTime;" json:"create_time"` // 创建时间
	UpdateTime       time.Time `gorm:"column:update_time;type:datetime;comment:创建时间;not null;default:CURRENT_TIMESTAMP;autoUpdateTime;" json:"update_time"` // 创建时间
	IsDeleted        int32     `gorm:"column:is_deleted;type:tinyint;comment:是否删除 0不是 1 是;not null;default:0;" json:"is_deleted"`                           // 是否删除 0不是 1 是
	OssBucket        string    `gorm:"column:oss_bucket;type:varchar(128);comment:oss_bucket;not null;" json:"oss_bucket"`                                  // oss_bucket
}
