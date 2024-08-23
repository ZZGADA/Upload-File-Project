package oss

import (
	"os"
	"path/filepath"
)

type ossServer struct{}

var OssServerImpl = &ossServer{}

// UploadSingleFile // Oss 上传单一文件
func (ossServer *ossServer) UploadSingleFile(objectName string, localFileName string) {
	// 获取存储空间。
	bucket, err := ossClient.Bucket(BucketName)
	if err != nil {
		HandleError(err)
	}
	// 上传文件。
	err = bucket.PutObjectFromFile(objectName, localFileName)
	if err != nil {
		HandleError(err)
	}
}

// DownLoadSingleFIle
/*
yourBucketName填写存储空间名称。
yourObjectName填写Object完整路径，完整路径中不能包含Bucket名称
*/
func (ossServer *ossServer) DownLoadSingleFIle(objectName string, downloadedFileName string, bucketName string) {
	logOss.Info("从OSS下载文件")
	// 获取存储空间。
	bucket, err := ossClient.Bucket(bucketName)
	if err != nil {
		HandleError(err)
	}

	// 下载文件 检查文件夹是否存在，如果不存在则直接下载
	checkFileDirExist(downloadedFileName)
	err = bucket.GetObjectToFile(objectName, downloadedFileName)
	if err != nil {

		HandleError(err)
	}
	logOss.Infof("downloadedFileName success，%#v", downloadedFileName)
}

// DeleteSingleFile // 删除OSS文件
func (ossServer *ossServer) DeleteSingleFile(objectName string) {
	// 获取存储空间。
	bucket, err := ossClient.Bucket(BucketName)
	if err != nil {
		HandleError(err)
	}
	// 删除文件。
	err = bucket.DeleteObject(objectName)
	if err != nil {
		HandleError(err)
	}
	logOss.Infof("oss文件删除成功,fileName = %#v", objectName)
}

func checkFileDirExist(filePath string) {
	dirPath := filepath.Dir(filePath)
	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// 如果文件不存在，检查文件夹是否存在
		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			// 如果文件夹不存在，则创建文件夹
			if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
				HandleError(err)
				return
			}
		}
	}
}
