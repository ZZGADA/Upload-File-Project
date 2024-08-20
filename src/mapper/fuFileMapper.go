package mapper

import (
	"UploadFileProject/src/entity/bo"
	"UploadFileProject/src/global"
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
