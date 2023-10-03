package logger_init

import "github.com/sirupsen/logrus"

func LogRusInit() *logrus.Logger {
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}
	logger.SetLevel(logrus.DebugLevel)

	logger.Infoln("logrus initialized")
	return logger
}
