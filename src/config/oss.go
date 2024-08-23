package config

import (
	"UploadFileProject/src/global"
	ossLoacl "UploadFileProject/src/oss"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

const EndPoint = "https://oss-cn-beijing.aliyuncs.com"

func initOssClient() {

	// 从环境变量中获取访问凭证。运行本代码示例之前，请确保已设置环境变量OSS_ACCESS_KEY_ID和OSS_ACCESS_KEY_SECRET。
	provider, err := oss.NewEnvironmentVariableCredentialsProvider()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为
	//https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	client, err := oss.New(EndPoint, provider.GetCredentials().GetAccessKeyID(), provider.GetCredentials().GetAccessKeySecret(), oss.SetCredentialsProvider(&provider))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 创建存储空间。
	err = client.CreateBucket(ossLoacl.BucketName)
	if err != nil {
		ossLoacl.HandleError(err)
	}

	global.OssClient = client
}
