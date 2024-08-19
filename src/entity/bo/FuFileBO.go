package bo

// FuFileBO 文件表
type FuFileBO struct {
	Id               int64  `gorm:"column:id;type:bigint;comment:id主键;primaryKey;" json:"id"`                                               // id主键
	FileUuid         string `gorm:"column:file_uuid;type:varchar(64);comment:文件唯一uuid;" json:"file_uuid"`                                   // 文件唯一uuid
	FileOriginalName string `gorm:"column:file_original_name;type:varchar(64);comment:文件名;not null;" json:"file_original_name"`             // 文件名
	FileSuffix       string `gorm:"column:file_suffix;type:varchar(32);comment:文件名后缀;not null;" json:"file_suffix"`                         // 文件名后缀
	IfUploadOss      int32  `gorm:"column:if_upload_oss;type:tinyint;comment:是否上传oss: 0 没有上传 1上传;not null;default:0;" json:"if_upload_oss"` // 是否上传oss: 0 没有上传 1上传
	OrgId            int64  `gorm:"column:org_id;type:bigint;comment:file文件所属的机构;not null;default:0;" json:"org_id"`                        // file文件所属的机构
	OssPath          string `gorm:"column:oss_path;type:varchar(255);comment:oss路径的存放地址;not null;" json:"oss_path"`                         // oss路径的存放地址
}

// TableName 指定表名
func (FuFileBO) TableName() string {
	return "fu_file"
}
