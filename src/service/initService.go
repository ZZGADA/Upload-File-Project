package service

import (
	"UploadFileProject/src/global"
	"github.com/sirupsen/logrus"
)

var (
	logService *logrus.Logger
)

func InitService() {
	logService = global.Log
}
