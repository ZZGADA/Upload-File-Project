package service

import (
	"UploadFileProject/src/config"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	logService  *logrus.Logger
	MySQLClient *gorm.DB
)

func InitService() {
	logService = config.Log
	MySQLClient = config.MySQLClient
}
