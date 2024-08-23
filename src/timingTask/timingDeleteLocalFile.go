package timingTask

import (
	"UploadFileProject/src/global"
	"UploadFileProject/src/mapper"
	"UploadFileProject/src/utils/process"
	"github.com/robfig/cron/v3"
	"os"
	"path/filepath"
	"time"
)

// timingDeleteLocalFile 开启定时删除本地文件任务 包括上传文件和缓存文件
func timingDeleteLocalFile() {
	// 创建一个新的 cron 实例，默认使用分钟调度
	c := cron.New(cron.WithSeconds())

	// 添加一个每 3 分钟执行一次的任务
	// 分 时 日 月 周
	_, err := c.AddFunc("*/50 * * * * *", deleteUploadFileTask)
	if err != nil {
		logTiming.Error("Error adding cron job:", err)
		return
	}

	_, err2 := c.AddFunc("*/50 * * * * *", deleteCacheFileTask)
	if err2 != nil {
		logTiming.Error("Error adding cron job:", err2)
		return
	}
	c.Start()
}

func deleteCacheFileTask() {
	dir := global.DownLoadsPath
	// 读取目录中的所有文件和文件夹
	entries, err := os.ReadDir(dir)
	if err != nil {
		logTiming.Error(err)
	}

	// 遍历每个条目并删除
	for _, entry := range entries {
		path := filepath.Join(dir, entry.Name())
		if entry.IsDir() {
			// 如果是文件夹，递归删除其内容
			err = os.RemoveAll(path)
		} else {
			// 如果是文件，直接删除
			err = os.Remove(path)
		}
		if err != nil {
			logTiming.Error(err)
		}
	}
	logTiming.Info("remove local cache file success")
}

func deleteUploadFileTask() {
	var filesUuid []string
	logTiming.Infof("定时删除文件任务执行，%v", time.Now())

	fileUploadWaitDeleteList := mapper.FuFileDeleteLocalMapperImpl.SelectUploadFileNotDelete()
	if len(fileUploadWaitDeleteList) == 0 {
		// 如果为空 就直接跳过
		logTiming.Info("upload 中需要删除的文件为空，没有需要删除的文件")
		return
	}

	for _, fileInfo := range fileUploadWaitDeleteList {
		filesUuid = append(filesUuid, fileInfo.FileUuid)
	}

	fileUploadWaitDeleteInformation := mapper.FuFileBOMapperImpl.GetBatchFileInformation(filesUuid)

	for _, fileInfo := range fileUploadWaitDeleteInformation {
		fileName := process.FileNameJoinSuffix(fileInfo.FileUuid, fileInfo.FileSuffix)
		filePath := filepath.Join(fileInfo.LocalGroup, fileInfo.OrgUuid, fileInfo.FileSuffix, fileName)
		logTiming.Info(filePath)
		if process.CheckFileExist(filePath) {
			// 如果为true 表示本地文件存在 需要删除
			// 为false 表示文件不存在 不需要删除
			// TODO: 文件如果删除失败 可以进入mq 等待重新删除
			if err := os.Remove(filePath); err != nil {
				logTiming.Errorf("文件删除失败，%#v", err)
				return
			}
			mapper.FuFileDeleteLocalMapperImpl.UpdateUploadFileDeletedStatue(fileInfo.FileUuid)
			logTiming.Infof("上传文件删除成功，filename is %#v , fileUuid is  %#v", fileName, fileInfo.FileUuid)
		} else {
			logTiming.Warnf("上传文件不存在，无需重复删除，filename is %#v , fileUuid is  %#v", fileName, fileInfo.FileUuid)
		}
	}
}
