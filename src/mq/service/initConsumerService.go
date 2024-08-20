package service

import "github.com/sirupsen/logrus"

var logCs *logrus.Logger

func InitConsumerService(log *logrus.Logger) {
	logCs = log
}
