package service

import (
	"UploadFileProject/src/entity/dto"
	"UploadFileProject/src/entity/vo"
	"UploadFileProject/src/global/enum"
	"UploadFileProject/src/mapper"
	"UploadFileProject/src/oss"
	"UploadFileProject/src/utils/process"
	"UploadFileProject/src/utils/resp"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/xuri/excelize/v2"
	"math"
	"os"
	"path/filepath"
)

type fileSearchService struct{}

var FileServiceImpl = &fileSearchService{}

const ExcelPath = "excel"

// GetFileInfoList // 文件信息批量导出
func (fileSearchService *fileSearchService) GetFileInfoList(listDTO *dto.FileInfoListDTO, organizationUuidStr string, result *resp.Result) {
	fileUuidList := listDTO.FileUuidList
	if len(fileUuidList) == 0 {
		// 传入参数为空
		result.Success("请选择文件信息")
		return
	}

	fileList := mapper.FuFileBOMapperImpl.GetBatchFileInformation(fileUuidList)

	// 创建一个新的 Excel 文件
	f := excelize.NewFile()

	// 创建一个新的工作表
	sheetName := "Sheet1"
	index, _ := f.NewSheet(sheetName)

	// 设置表头
	headers := []string{"文件名", "文件名后缀", "文件uuid", "所属组织名称", "上传地址", "是否长传OSS", "OSS路径", "OSS bucket", "创建时间", "更新时间"}
	for i, header := range headers {
		cell := fmt.Sprintf("%s1", string(rune('A'+i)))
		err := f.SetCellValue(sheetName, cell, header)
		if err != nil {
			return
		}
	}

	// 插入数据
	for i, record := range fileList {
		var _ error
		var fileName = process.FileNameJoinSuffix(record.FileName, record.FileSuffix)
		row := i + 2 // 从第二行开始插入数据
		_ = f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), fileName)
		_ = f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), record.FileSuffix)
		_ = f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), record.FileUuid)
		_ = f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), record.OrgName)
		_ = f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), record.LocalGroup)
		_ = f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), record.IfUploadOss)
		_ = f.SetCellValue(sheetName, fmt.Sprintf("G%d", row), record.OssPath)
		_ = f.SetCellValue(sheetName, fmt.Sprintf("H%d", row), record.OssBucket)
		_ = f.SetCellValue(sheetName, fmt.Sprintf("I%d", row), record.CreateTime.Format("2006-01-02 15:04:05"))
		_ = f.SetCellValue(sheetName, fmt.Sprintf("J%d", row), record.UpdateTime.Format("2006-01-02 15:04:05"))
	}

	// 设置活动工作表
	f.SetActiveSheet(index)

	// 保存 Excel 文件
	filaPathExcel := filepath.Join(ExcelPath, organizationUuidStr, process.FileNameJoinSuffix(uuid.NewV1().String(), "xlsx"))
	process.CheckFileDirExist(filaPathExcel)
	if err := f.SaveAs(filaPathExcel); err != nil {
		fmt.Println(err)
	}
	//result.Ctx.File(filaPathExcel)
	result.SuccessMsg("excel download success", vo.FileInfoListVO{ExcelPath: filaPathExcel})
}

// DeleteFile // 文件删除
func (fileSearchService *fileSearchService) DeleteFile(deleteDTO *dto.FileDeleteDTO, result *resp.Result) {
	fuFileObj := mapper.FuFileBOMapperImpl.GetOneFile(deleteDTO.FileUuid)
	if fuFileObj.Id == 0 {
		// 为nil 表示已经删除了
		result.Success("文件已经删除，无需重复删除")
		return
	}

	// TODO: redis 加锁
	// 删除文件的时候 如果文件正在读取，则不能删除 所以需要加锁
	// 文件下载的时候 需要针对对象加锁 当文件删除的时候 判断是否有锁 没有锁就删除
	orgObj := mapper.FuOrganizationBOMapperImpl.SelectFuOrganizationByID(fuFileObj.OrgId)
	fileName := process.FileNameJoinSuffix(fuFileObj.FileUuid, fuFileObj.FileSuffix)

	localFilePath := filepath.Join(fuFileObj.LocalGroup, orgObj.OrgUuid, fuFileObj.FileSuffix, fileName)
	ossFilePath := localFilePath

	if fuFileObj.IfUploadOss == enum.UploadOss.ToInt32() {
		// 如果上传了OSS 就删除OSS  也要查看本地文件是否存在，如果存在表示上传了OSS但是删除的定时任务还没执行
		if process.CheckFileExist(localFilePath) {
			if err := os.Remove(localFilePath); err != nil {
				logService.Errorf("本地文件删除错误，%#v", err)
			}
		}

		// 直接删除OSS
		// 下载缓存区不需要删除 获取文件接口只读未删除文件
		oss.OssServerImpl.DeleteSingleFile(ossFilePath)

	} else {
		// 如果没有上传OSS，表示文件刚刚上传到系统，异步上传OSS还没有执行
		// TODO: 发送消息进入mq 然后消费者消费 如果消费的时候OSS文件仍然没有上传 就重新入队 等待消费 否则直接删除OSS
		result.Success("网络繁忙，请稍后重试哦")
		return
	}

	mapper.FuFileBOMapperImpl.DeleteFile(deleteDTO.FileUuid)
	result.Success("delete file success")
}

// UpdateFileName // 更新文件名
func (fileSearchService *fileSearchService) UpdateFileName(fileUpdateName *dto.FileUpdateName, result *resp.Result) {
	mapper.FuFileBOMapperImpl.UpdateFileName(fileUpdateName.FileUuid, fileUpdateName.NewName)
	result.Success(&vo.FileUpdateVO{
		FileUuid: fileUpdateName.FileUuid,
		FileName: fileUpdateName.NewName,
	})
}

// GetFileList 获取文件列表
func (fileSearchService *fileSearchService) GetFileList(fileSearchDTO *dto.FileSearchDTO, result *resp.Result) {
	// 进行分页查询
	// 如果pageCurrent为0 则从第一页开始查
	// 否则设定查询的数量
	fileSearchItem := fileSearchDTO.SearchItem
	pageCurrent := fileSearchDTO.PageCurrent
	pageSize := fileSearchDTO.PageSize
	if pageSize == 0 {
		pageSize = 10
	}
	if pageCurrent == 0 {
		pageCurrent = 1
	}

	var fuFilesVO []dto.FileListDTO
	limitStart := pageSize * (pageCurrent - 1)
	rowsTotal := int(mapper.FuFileBOMapperImpl.QueryAllData(fileSearchItem))
	pageTotal := int(math.Ceil(float64(rowsTotal) / float64(pageSize)))

	// 如果当前页大于最大页 就直接返回
	if pageCurrent <= pageTotal {
		fuFiles := mapper.FuFileBOMapperImpl.PageQuery(
			fileSearchItem,
			pageSize,
			limitStart,
		)
		for _, val := range fuFiles {
			fuFilesVO = append(fuFilesVO,
				dto.FileListDTO{
					FileName:   val.FileOriginalName,
					FileUuid:   val.FileUuid,
					FileSuffix: val.FileSuffix,
					UpdateTime: val.UpdateTime,
				})
		}
	} else {
		fuFilesVO = make([]dto.FileListDTO, 0)
	}

	resultVO := vo.FileListVO{
		FileSearchDTO: dto.FileSearchDTO{
			PageTotal:   pageTotal,
			PageSize:    pageSize,
			PageCurrent: pageCurrent,
			RowsTotal:   rowsTotal,
			SearchItem:  fileSearchItem,
		},
		FileList: fuFilesVO,
	}
	logService.Infof("pageTotal: %d  pageSize: %d ,rowsTotal: %d,pageCurrent:%d,SearchItem:%s",
		pageTotal,
		pageSize,
		rowsTotal,
		pageCurrent,
		fileSearchItem)

	result.Success(resultVO)
}
