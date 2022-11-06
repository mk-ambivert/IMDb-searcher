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
	logsFileName := "logs/logfile.log"
	pathToLogs := filepath.Join(projectDir, logsFileName)

	file, err := os.OpenFile(pathToLogs, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		msg := errors.New("Failed to initalize logging" + err.Error())
		panic(msg)
	} else {
		logger.SetOutput(file)
	}

	logger.SetLevel(logrus.DebugLevel)

	loggerEntry = *logrus.NewEntry(logger)
}
