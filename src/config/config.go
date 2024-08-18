package config

import (
	"bytes"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"strings"
	"time"
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
	initNacosClient()
	pullNacosBootStrapConfig()
	initMySQLClient()

}

// 初始化NacosClient 包括动态配置和服务发现
func initNacosClient() {
	naocsConfig := ProjectConfig.NacosConfig

	sc := []constant.ServerConfig{
		*constant.NewServerConfig(naocsConfig.IP, naocsConfig.Port, constant.WithContextPath(naocsConfig.ContextPath)),
	}

	//create ClientConfig
	cc := *constant.NewClientConfig(
		constant.WithTimeoutMs(naocsConfig.TimeoutMs),
		constant.WithNotLoadCacheAtStart(naocsConfig.NotLoadCacheAtStart),
		constant.WithLogDir(naocsConfig.Dir.Log),
		constant.WithCacheDir(naocsConfig.Dir.Cache),
		constant.WithLogLevel(naocsConfig.LogLevel),
		constant.WithUsername(naocsConfig.Username),
		constant.WithPassword(naocsConfig.Password),
	)

	nacosConfigClient, errConfig := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	nacosServicesClient, errService := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		})

	if errConfig != nil {
		log.Panic(fmt.Sprintf("nacos config client 初始化失败 %v", errConfig))
	}

	if errService != nil {
		log.Panic(fmt.Sprintf("nacos services client 初始化失败 %v", errService))
	}
	NacosConfigClient = nacosConfigClient
	NacosServicesClient = nacosServicesClient

	log.Printf("Nacos config client has been created ")
}

// 拉去nacos的动态配置文件
func pullNacosBootStrapConfig() {
	bootstrapDynamicConfig, err := NacosConfigClient.GetConfig(
		vo.ConfigParam{
			Group:  ProjectConfig.NacosConfig.Bootstrap.Group,
			DataId: ProjectConfig.NacosConfig.Bootstrap.DataId,
		})
	if err != nil {
		log.Panicf("nacos 拉取动态文件失败 ：%v", err)
	}

	if err := yaml.Unmarshal([]byte(bootstrapDynamicConfig), &ProjectConfig); err != nil {
		log.Printf("动态配置解析到NacosBootstrapConfig失败，%#v", err)
	}

	// 存入viper viper 中的对象是key-value(map)的形式
	byteAll, _ := yaml.Marshal(ProjectConfig)
	if err := viper.ReadConfig(bytes.NewBuffer(byteAll)); err != nil {
		log.Panicf("nacos远程动态配置和本地配置合并后写入失败，%#v", err)
	}
	log.Printf("Nacos bootstrap config is %v", ProjectConfig)
}

// 初始化MySQLClient
func initMySQLClient() {

	const mysqlConnectStr string = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"

	dsn := fmt.Sprintf(mysqlConnectStr,
		ProjectConfig.DataBases.Mysql.Username,
		ProjectConfig.DataBases.Mysql.Password,
		ProjectConfig.DataBases.Mysql.Ip,
		ProjectConfig.DataBases.Mysql.Port,
		ProjectConfig.DataBases.Mysql.Database)

	MySQLClientDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",   // 表前缀
			SingularTable: true, // 禁用表名复数
		}})
	if err != nil {
		panic(err)
	}
	//
	sqlDB, _ := MySQLClientDB.DB()
	// SetMaxIdleConnections 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConnections 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	MySQLClient = MySQLClientDB
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
