package logger

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

var log *logrus.Logger

func InitLogger(env string) error {

	switch env {
	case envLocal:
		log = logrus.New()
		log.Formatter = &logrus.JSONFormatter{}
		log.Out = os.Stdout
		log.Level = logrus.DebugLevel
	case envDev:
		log = logrus.New()
		log.Formatter = &logrus.JSONFormatter{}
		log.Out = os.Stdout
		log.Level = logrus.DebugLevel
	case envProd:
		log = logrus.New()
		log.Formatter = &logrus.JSONFormatter{}
		log.Out = os.Stdout
		log.Level = logrus.InfoLevel

	default:
		log = logrus.New()
		log.Formatter = &logrus.JSONFormatter{}
		log.Out = os.Stdout
		log.Level = logrus.DebugLevel
		return fmt.Errorf("unknown env: %s", env)
	}

	log.WithFields(logrus.Fields{
		"env": env})

	return nil
}

func GetLogger() *logrus.Logger {
	return log
}
