package logger

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

// Singleton Pattern
func GetLogger() *logrus.Logger {

	var once sync.Once
	if logger == nil {
		once.Do(
			func() {
				logger = logrus.New()
				err := godotenv.Load(".env")
				if err != nil {
					return
				}
				LOG_FILE := os.Getenv("LOG_FILE")
				src, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

				if err != nil {
					fmt.Print(err.Error())
					fmt.Print("unable to create user.log file")
				}

				multiWriter := io.MultiWriter(os.Stdout, src)

				logger.SetFormatter(&logrus.JSONFormatter{})
				logger.SetOutput(multiWriter)
			})
	}

	return logger
}
