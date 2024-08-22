package mapper

import (
	"UploadFileProject/src/entity/bo"
	"UploadFileProject/src/global"
)

type FuOrganizationBOMapper struct {
}

var FuOrganizationBOMapperImpl = &FuOrganizationBOMapper{}

// SelectFuOrganization //查询组织
func (mapper *FuOrganizationBOMapper) SelectFuOrganization(uuid string) *bo.FuOrganizationBO {
	var FuOrganizationBOStructure = &bo.FuOrganizationBO{}
	mysqlClient.Where("org_uuid = ?", uuid).First(FuOrganizationBOStructure)
	return FuOrganizationBOStructure
}

func (mapper *FuOrganizationBOMapper) SelectFuOrganizationByID(orgId int64) *bo.FuOrganizationBO {
	var FuOrganizationBOStructure = &bo.FuOrganizationBO{}
	mysqlClient.Where("id = ?", orgId).First(FuOrganizationBOStructure)
	return FuOrganizationBOStructure
}

// CreateFuOrganization //创建组织
func (mapper *FuOrganizationBOMapper) CreateFuOrganization(fuOrganizationBO *bo.FuOrganizationBO) {
	result := mysqlClient.Create(fuOrganizationBO)
	if result.Error != nil {
		global.Log.Errorf("fu_organization create  失败，%#v", result.Error)
	}
}
