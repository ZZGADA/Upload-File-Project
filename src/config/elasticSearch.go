package config

import (
	"UploadFileProject/src/global"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
)

const Url = "http://%s:%d"

func initESClient() {
	ESServerURL := fmt.Sprintf(Url, ProjectConfig.EsConfig.Host, ProjectConfig.EsConfig.Port)
	es, err := elasticsearch.NewClient(elasticsearch.Config{Addresses: []string{ESServerURL}})
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// 打印 Elasticsearch 版本信息
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	fmt.Println(res.String())
	global.ESClient = es
}
