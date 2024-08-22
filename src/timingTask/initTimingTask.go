package timingTask

import (
	"UploadFileProject/src/global"
	"github.com/sirupsen/logrus"
)

var logTiming *logrus.Logger

// InitTimingTask //定时任务的启动与创建
func InitTimingTask() {
	logTiming = global.Log
	timingDeleteLocalFile()
}
