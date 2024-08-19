package bo

// FuOrganizationBO 组织表
type FuOrganizationBO struct {
	Id      int64  `json:"id" mapstructure:"id" gorm:"column:id; PRIMARY_KEY; AUTO_INCREMENT"` // 自增id
	OrgUuid string `json:"org_uuid" mapstructure:"org_uuid" gorm:"column:org_uuid"`            // 组织唯一id
	OrgName string `json:"org_name" mapstructure:"org_name" gorm:"column:org_name"`            // 组织名称
	OrgPath string `json:"org_path" mapstructure:"org_path" gorm:"column:org_path"`            // 组织文件存储路径
}

// TableName 指定表名
func (FuOrganizationBO) TableName() string {
	return "fu_organization"
}
