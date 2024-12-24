package logger_1

import (
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func InitLogger() *logrus.Logger {
	log = logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.InfoLevel)
	return log
}
