package log

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitLogger() *logrus.Logger {
	logger := logrus.New()

	viper.AutomaticEnv()

	logLevel := viper.GetString("LOG_LEVEL")
	switch logLevel {
	case "debug":
		logger.SetLevel(logrus.DebugLevel)
	case "info":
		logger.SetLevel(logrus.InfoLevel)
	case "warn":
		logger.SetLevel(logrus.WarnLevel)
	case "error":
		logger.SetLevel(logrus.ErrorLevel)
	default:
		logger.SetLevel(logrus.InfoLevel)
	}

	logFormat := viper.GetString("LOG_FORMAT")
	if logFormat == "json" {
		logger.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	}

	logger.SetReportCaller(true)

	return logger
}
