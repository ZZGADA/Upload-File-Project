package oss

type OssServer struct{}

var OssServerImpl = &OssServer{}

// UploadSingleFile // Oss 上传单一文件
func (ossServer *OssServer) UploadSingleFile(objectName string, localFileName string) {
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
func (ossServer *OssServer) DownLoadSingleFIle(objectName string, downloadedFileName string) {
	// 获取存储空间。
	bucket, err := ossClient.Bucket(BucketName)
	if err != nil {
		HandleError(err)
	}
	// 下载文件。
	err = bucket.GetObjectToFile(objectName, downloadedFileName)
	if err != nil {
		HandleError(err)
	}
}
