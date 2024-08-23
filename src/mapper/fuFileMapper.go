package mapper

import (
	"UploadFileProject/src/entity/bo"
	"UploadFileProject/src/entity/dto"
	"UploadFileProject/src/global"
	"UploadFileProject/src/global/enum"
	"time"
)

type fuFileBOMapper struct{}

// FuFileBOMapperImpl 对外暴露mapper服务
var FuFileBOMapperImpl = &fuFileBOMapper{}

// InsertFuFile 插入文件
func (mapper *fuFileBOMapper) InsertFuFile(fuFileBO *bo.FuFileBO) int64 {
	result := mysqlClient.Create(fuFileBO)
	if result.Error != nil {
		global.Log.Errorf("fu_file create 失败，%#v", result.Error)
	}
	return result.RowsAffected
}

// UpdateOssPath //更新OssPath
func (mapper *fuFileBOMapper) UpdateOssPath(fileUuid string, ossPath string, ossBucket string) {
	mysqlClient.Model(&bo.FuFileBO{}).
		Where("file_uuid = ?", fileUuid).
		Updates(bo.FuFileBO{
			OssPath:     ossPath,
			IfUploadOss: enum.UploadOss.ToInt32(),
			OssBucket:   ossBucket,
			UpdateTime:  time.Now()})
}

// UpdateFileName // 更新文件名
func (mapper *fuFileBOMapper) UpdateFileName(fileUuid string, newFileName string) {
	mysqlClient.Model(&bo.FuFileBO{}).
		Where("file_uuid = ?", fileUuid).
		Updates(bo.FuFileBO{
			FileOriginalName: newFileName,
			UpdateTime:       time.Now()})
}

// GetOneFile // 获取file信息
func (mapper *fuFileBOMapper) GetOneFile(fileUuid string) *bo.FuFileBO {
	var fuFileBO *bo.FuFileBO = &bo.FuFileBO{}
	mysqlClient.Where("is_deleted = ? and file_uuid = ?", enum.NoneDeleted.ToInt32(), fileUuid).First(fuFileBO)
	return fuFileBO
}

// GetOneFileOrg //获取文件和org信息
func (mapper *fuFileBOMapper) GetOneFileOrg(fileUuid string) *bo.FuFileBO {
	var fuFileBO *bo.FuFileBO = &bo.FuFileBO{}
	mysqlClient.Where("is_deleted = ? and file_uuid = ?", enum.NoneDeleted.ToInt32(), fileUuid).First(fuFileBO)
	return fuFileBO
}

// PageQuery 分页查询
func (mapper *fuFileBOMapper) PageQuery(fileSearchItem string,
	pageSize int,
	limitStart int) []bo.FuFileBO {

	var fuFiles []bo.FuFileBO
	mysqlClient.Model(&bo.FuFileBO{}).Select("file_uuid", "file_original_name", "file_suffix", "update_time").
		Where("is_deleted =? and file_original_name like ? ", enum.NoneDeleted.ToInt32(), "%"+fileSearchItem+"%").
		Limit(pageSize).
		Offset(limitStart).
		Find(&fuFiles)
	return fuFiles
}

// QueryAllData 查询总数量
func (mapper *fuFileBOMapper) QueryAllData(fileSearchItem string) int64 {
	var result int64
	mysqlClient.Model(&bo.FuFileBO{}).
		Where("is_deleted =? and file_original_name like ? ", enum.NoneDeleted.ToInt32(), "%"+fileSearchItem+"%").
		Count(&result)
	return result
}

func (mapper *fuFileBOMapper) DeleteFile(fileUuid string) {
	mysqlClient.Model(&bo.FuFileBO{}).
		Where("file_uuid = ?", fileUuid).Updates(bo.FuFileBO{
		IsDeleted:  enum.Deleted.ToInt32(),
		UpdateTime: time.Now()})
}

func (mapper *fuFileBOMapper) GetBatchFileInformation(fileUuidList []string) []dto.FileListTDTO {
	var result []dto.FileListTDTO
	mysqlClient.Table("fu_file as ff").
		Select("ff.file_original_name as fileName,"+
			"ff.file_suffix as fileSuffix,"+
			"ff.file_uuid as fileUuid,"+
			"fo.org_name as orgName,"+
			"fo.org_uuid as orgUuid,"+
			"ff.local_group as localGroup,"+
			"ff.if_upload_oss as ifUploadOss,"+
			"ff.oss_path as ossPath,"+
			"ff.oss_bucket as ossBucket,"+
			"ff.create_time as createTime,"+
			"ff.update_time as updateTime").
		Joins("inner join fu_organization as fo on ff.org_id = fo.id").
		Where("ff.file_uuid in ? ", fileUuidList).
		Order("ff.update_time desc").
		Scan(&result)
	return result
}
