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
	"gopkg.in/yaml.v3"
	"log"
)

// 初始化NacosClient 包括动态配置和服务发现
func initNacosClient() (config_client.IConfigClient, naming_client.INamingClient) {
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

	log.Printf("Nacos config client has been created ")

	return nacosConfigClient, nacosServicesClient
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
