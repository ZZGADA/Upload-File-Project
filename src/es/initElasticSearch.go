package es

import (
	"UploadFileProject/src/global"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

var logEs *logrus.Logger
var EsClient *elastic.Client

func InitElasticSearch() {
	logEs = global.Log
	EsClient = global.ESClient

}
