package bo

import "time"

// FuOrganizationBO 组织表
type FuOrganizationBO struct {
	Id         int64     `gorm:"column:id;type:bigint;comment:自增id;primaryKey;" json:"id"`                                   // 自增id
	OrgUuid    string    `gorm:"column:org_uuid;type:varchar(64);comment:组织唯一id;not null;" json:"org_uuid"`                  // 组织唯一id
	OrgName    string    `gorm:"column:org_name;type:varchar(255);comment:组织名称;not null;" json:"org_name"`                   // 组织名称
	OrgPath    string    `gorm:"column:org_path;type:varchar(64);comment:组织文件存储路径;" json:"org_path"`                         // 组织文件存储路径
	CreateTime time.Time `gorm:"column:create_time;type:datetime;comment:创建时间;autoCreateTime" json:"create_time"`            // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;type:datetime;comment:更新时间;autoUpdateTime" json:"update_time"`            // 更新时间
	IsDeleted  int32     `gorm:"column:is_deleted;type:tinyint;comment:是否删除 0 没有 1 有;not null;default:0;" json:"is_deleted"` // 是否删除 0 没有 1 有
}

// TableName 指定表名
func (FuOrganizationBO) TableName() string {
	return "fu_organization"
}
