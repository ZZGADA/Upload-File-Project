package config

import (
	"UploadFileProject/src/global"
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

func initESClient() {
	esClient, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	if err != nil {
		panic(fmt.Sprintf("es init failed,%#v", err))
	}

	// 检查Elasticsearch是否可用
	info, code, err := esClient.Ping("http://localhost:9200").Do(context.Background())
	if err != nil {
		panic(fmt.Sprintf("Error pinging Elasticsearch: %s", err))
	}

	global.Log.Info("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
	global.ESClient = esClient
}
