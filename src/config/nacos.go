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

var (
	NacosConfigClient   config_client.IConfigClient
	NacosServicesClient naming_client.INamingClient
	NacosServerConfig   = &NacosConfig{}
)

// 初始化NacosClient 包括动态配置和服务发现
func initNacosClient() {
	NacosServerConfig = &ProjectConfig.NacosConfig
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(NacosServerConfig.IP, NacosServerConfig.Port, constant.WithContextPath(NacosServerConfig.ContextPath)),
	}

	//create ClientConfig
	cc := *constant.NewClientConfig(
		constant.WithTimeoutMs(NacosServerConfig.TimeoutMs),
		constant.WithNotLoadCacheAtStart(NacosServerConfig.NotLoadCacheAtStart),
		constant.WithLogDir(NacosServerConfig.Dir.Log),
		constant.WithCacheDir(NacosServerConfig.Dir.Cache),
		constant.WithLogLevel(NacosServerConfig.LogLevel),
		constant.WithUsername(NacosServerConfig.Username),
		constant.WithPassword(NacosServerConfig.Password),
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

	NacosConfigClient, NacosServicesClient = nacosConfigClient, nacosServicesClient
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

	if err := yaml.Unmarshal([]byte(bootstrapDynamicConfig), ProjectConfig); err != nil {
		log.Printf("动态配置解析到NacosBootstrapConfig失败，%#v", err)
	}

	// 存入viper viper 中的对象是key-value(map)的形式
	byteAll, _ := yaml.Marshal(ProjectConfig)
	if err := viper.ReadConfig(bytes.NewBuffer(byteAll)); err != nil {
		log.Panicf("nacos远程动态配置和本地配置合并后写入失败，%#v", err)
	}
	log.Printf("Nacos bootstrap config is %v", *ProjectConfig)
}

func registerServiceInstance(client naming_client.INamingClient, param vo.RegisterInstanceParam) {
	success, err := client.RegisterInstance(param)
	if !success || err != nil {
		panic("RegisterServiceInstance failed!" + err.Error())
	}
	fmt.Printf("RegisterServiceInstance,param:%+v,result:%+v \n\n", param, success)

}

func registerUploadFileService() {
	//Register
	registerServiceInstance(NacosServicesClient, vo.RegisterInstanceParam{
		Ip:          ServerAllConfig.IP,
		Port:        ServerAllConfig.Port,
		ServiceName: ServerAllConfig.Name,
		GroupName:   ServerAllConfig.Group,
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"idc": "TAL"},
	})
}
