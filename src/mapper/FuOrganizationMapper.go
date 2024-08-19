package mapper

import (
	"UploadFileProject/src/config"
	"UploadFileProject/src/entity/bo"
	"gorm.io/gorm"
)

type FuOrganizationBOMapper struct {
}

var FuOrganizationBOMapperImpl = &FuOrganizationBOMapper{}
var mySQLClient *gorm.DB

// 初始化initFuOrganizationBOMapper 用于操作数据库
func initFuOrganizationBOMapper() {
	mySQLClient = config.MySQLClient
}

// 查询组织
func (mapper *FuOrganizationBOMapper) SelectOrganization(uuid string) *bo.FuOrganizationBO {
	var FuOrganizationBOStructure = &bo.FuOrganizationBO{}
	mySQLClient.Where("org_uuid = ?", uuid).First(FuOrganizationBOStructure)
	return FuOrganizationBOStructure
}
