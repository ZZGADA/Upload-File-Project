package mapper

import (
	"UploadFileProject/src/entity/bo"
	"UploadFileProject/src/global"
	"UploadFileProject/src/global/enum"
	"time"
)

type fuFileDeleteLocalMapper struct {
}

// FuFileDeleteLocalMapperImpl 对外暴露mapper服务
var FuFileDeleteLocalMapperImpl = &fuFileDeleteLocalMapper{}

// CreateFuFileDeleteLocal 创建组织
func (mapper *fuFileDeleteLocalMapper) CreateFuFileDeleteLocal(localBO *bo.FuFileDeleteLocalBO) {
	result := mysqlClient.Create(localBO)
	if result.Error != nil {
		global.Log.Errorf("fu_organization create  失败，%#v", result.Error)
	}
}

// SelectUploadFileNotDelete  查询上传文件没有删除的
func (mapper *fuFileDeleteLocalMapper) SelectUploadFileNotDelete() []bo.FuFileDeleteLocalBO {
	var result = []bo.FuFileDeleteLocalBO{}
	mysqlClient.Model(&bo.FuFileDeleteLocalBO{}).
		Where("upload_file_deleted = ? ", enum.NoneDeleted).Find(&result)
	return result
}

// UpdateUploadFileDeletedStatue 更新删除本地文件状态
func (mapper *fuFileDeleteLocalMapper) UpdateUploadFileDeletedStatue(fileUuid string) {
	mysqlClient.Model(&bo.FuFileDeleteLocalBO{}).
		Where("file_uuid = ?", fileUuid).
		Updates(bo.FuFileDeleteLocalBO{
			UploadFileDeleted: enum.Deleted.ToInt32(),
			UpdateTime:        time.Now()})
}
