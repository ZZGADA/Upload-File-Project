package mapper

import (
	"UploadFileProject/src/config"
	"UploadFileProject/src/entity/bo"
)

type FuFileBOMapper struct {
}

var FuFileBOMapperImpl = &FuFileBOMapper{}

// 查询组织
func (mapper *FuFileBOMapper) InsertFile(fuFileBO *bo.FuFileBO) int64 {
	result := config.MySQLClient.Create(fuFileBO)
	if result.Error != nil {
		config.Log.Errorf("ful file 失败，%#v", result.Error)
	}
	return result.RowsAffected
}
