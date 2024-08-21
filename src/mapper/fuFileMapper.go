package mapper

import (
	"UploadFileProject/src/entity/bo"
	"UploadFileProject/src/global"
	"UploadFileProject/src/global/enum"
)

type FuFileBOMapper struct{}

var FuFileBOMapperImpl = &FuFileBOMapper{}

// 查询组织
func (mapper *FuFileBOMapper) InsertFuFile(fuFileBO *bo.FuFileBO) int64 {
	result := mysqlClient.Create(fuFileBO)
	if result.Error != nil {
		global.Log.Errorf("fu_file create 失败，%#v", result.Error)
	}
	return result.RowsAffected
}

func (mapper *FuFileBOMapper) UpdateOssPath(fileUuid string, ossPath string, ossBucket string) {
	mysqlClient.Model(&bo.FuFileBO{}).
		Where("file_uuid = ?", fileUuid).
		Updates(bo.FuFileBO{
			OssPath:     ossPath,
			IfUploadOss: enum.UploadOss.ToInt32(),
			OssBucket:   ossBucket})
}
