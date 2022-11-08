package logger

import (
	"errors"
	"os"
	"path/filepath"
	"sync"

	"github.com/sirupsen/logrus"
)

type ILogger interface {
	Debug(...interface{})
	Info(...interface{})
	Error(...interface{})
	Fatal(...interface{})
	Panic(...interface{})
}

var once sync.Once
var loggerEntry logrus.Entry

func GetLogger() ILogger {
	once.Do(Init)

	return &loggerEntry
}

func Init() {
	logger := logrus.New()

	logger.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	projectDir := os.Getenv("PROJECT_DIR")
	folderName := "logs"
	pathToLogsDir := filepath.Join(projectDir, folderName)

	err := os.MkdirAll(pathToLogsDir, 0644)
	if err != nil {
		msg := errors.New("failed to initialize log dir:" + err.Error())
		panic(msg)
	}

	pathToLogFile := filepath.Join(pathToLogsDir, "log.log")
	file, err := os.OpenFile(pathToLogFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		msg := errors.New("failed to initialize logging:" + err.Error())
		panic(msg)
	} else {
		logger.SetOutput(file)
	}

	logger.SetLevel(logrus.DebugLevel)

	loggerEntry = *logrus.NewEntry(logger)
}
