package bo

import "time"

type FuFileDeleteLocal struct {
	Id             int64     `gorm:"column:id;type:bigint;comment:id;primaryKey;" json:"id"`                                               // id
	FileDeleteUuid string    `gorm:"column:file_delete_uuid;type:varchar(64);comment:删除文件记录唯一id;not null;" json:"file_delete_uuid"`        // 删除文件记录唯一id
	CreateTime     time.Time `gorm:"column:create_time;type:datetime;comment:创建时间;not null;default:CURRENT_TIMESTAMP;" json:"create_time"` // 创建时间
	UpdateTime     time.Time `gorm:"column:update_time;type:datetime;comment:更新时间;default:CURRENT_TIMESTAMP;" json:"update_time"`          // 更新时间
	FileUuid       string    `gorm:"column:file_uuid;type:varchar(64);comment:文件uuid;not null;" json:"file_uuid"`                          // 文件uuid
	IsDeleted      int32     `gorm:"column:is_deleted;type:tinyint;comment:是否删除 0不是 1 是;not null;default:0;" json:"is_deleted"`            // 是否删除 0不是 1 是
}
