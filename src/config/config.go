package config

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
	"strings"
)

/*
全局ProjectConfig 配置文件
全局NacosClient  Nacos的动态配置中心
*/
var (
	ProjectConfig       Config
	NacosConfigClient   config_client.IConfigClient
	NacosServicesClient naming_client.INamingClient
	MySQLClient         *gorm.DB
	Log                 *logrus.Logger
)

/*
配置静态常量
*/
//const (
//	bootstrapName string = ""
//)

// LoadResource /*加载配置文件资源*/
func LoadResource(configFile string) {

	readResourceFile(configFile)
	NacosConfigClient, NacosServicesClient = initNacosClient()
	pullNacosBootStrapConfig()
	MySQLClient = initMySQLClient()
	Log = initLogConfig()
}

// 读取资源文件
func readResourceFile(configFileName string) {
	fileConfig := strings.Split(configFileName, ".")
	if len(fileConfig) == 2 {
		switch fileConfig[0] {
		case "application":
			readApplicationFile(fileConfig)
		default:
			log.Panic(fmt.Sprintf("sorry, this kind of file hasn't function to read"), configFileName)
		}

	} else {
		log.Panic("file name isn't completed")
	}
}

// yaml 文件读取
func readApplicationFile(fileConfig []string) {
	configName := strings.Join(fileConfig, ".")
	// 配置读取yaml 文件
	viper.SetConfigName(configName)     // 配置文件名称(无扩展名)
	viper.SetConfigType(fileConfig[1])  // 或viper.SetConfigType("YAML")
	viper.AddConfigPath("./src/config") // 配置文件路径
	if err := viper.ReadInConfig(); err != nil {
		// 处理读取配置文件的错误
		// 小写
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// 获取对象
	if err := viper.Unmarshal(&ProjectConfig); err != nil {
		panic("viper 转换对象错误 \n")
	}

	log.Printf("config has been read %#v", ProjectConfig)
}
