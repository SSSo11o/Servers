package logger

import (
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

func InitLumberjackLogger() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.InfoLevel)

	log.SetOutput(&lumberjack.Logger{
		Filename:   "./logs/app.log",
		MaxSize:    10,
		MaxBackups: 3,
		MaxAge:     28,
	})

	return log
}
