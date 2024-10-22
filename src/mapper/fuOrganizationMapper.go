package mapper

import (
	"UploadFileProject/src/entity/bo"
	"UploadFileProject/src/global"
)

type fuOrganizationBOMapper struct {
}

// FuOrganizationBOMapperImpl 对外暴露mapper服务
var FuOrganizationBOMapperImpl = &fuOrganizationBOMapper{}

// SelectFuOrganization 查询组织
func (mapper *fuOrganizationBOMapper) SelectFuOrganization(uuid string) *bo.FuOrganizationBO {
	var FuOrganizationBOStructure = &bo.FuOrganizationBO{}
	mysqlClient.Where("org_uuid = ?", uuid).First(FuOrganizationBOStructure)
	return FuOrganizationBOStructure
}

// SelectFuOrganizationByID 根据组织id 查询组织信息
func (mapper *fuOrganizationBOMapper) SelectFuOrganizationByID(orgId int64) *bo.FuOrganizationBO {
	var FuOrganizationBOStructure = &bo.FuOrganizationBO{}
	mysqlClient.Where("id = ?", orgId).First(FuOrganizationBOStructure)
	return FuOrganizationBOStructure
}

// CreateFuOrganization 创建组织
func (mapper *fuOrganizationBOMapper) CreateFuOrganization(fuOrganizationBO *bo.FuOrganizationBO) {
	result := mysqlClient.Create(fuOrganizationBO)
	if result.Error != nil {
		global.Log.Errorf("fu_organization create  失败，%#v", result.Error)
	}
}
