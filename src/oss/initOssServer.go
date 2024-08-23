package oss

import (
	"UploadFileProject/src/global"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/sirupsen/logrus"
	"os"
)

// OssClient Oss操作端
var (
	ossClient *oss.Client
	logOss    *logrus.Logger
)

const BucketName = "zzgeda-oss"

// InitOssServer 初始化Oss服务
func InitOssServer() {
	ossClient = global.OssClient
	logOss = global.Log
}

// HandleError 异常处理
func HandleError(err error) {
	logOss.Error("OSS Error:", err)
	os.Exit(-1)
}
