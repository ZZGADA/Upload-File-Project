package global

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Log                 *logrus.Logger
	MySQLClient         *gorm.DB
	NacosConfigClient   config_client.IConfigClient
	NacosServicesClient naming_client.INamingClient
	OssClient           *oss.Client
	ESClient            *elastic.Client
)

const (
	Authorization string = "Authorization"
	Organization  string = "OrganizationID"
)

const (
	SingleFileName = "singleFile"
	MultiFileName  = "multiFile"
)

const (
	UpLoadsPath   = "uploads"
	DownLoadsPath = "downloads"
)
