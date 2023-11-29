package logger

import (
	"github.com/hashicorp/go-hclog"
	"os"
	"sync"
)

var logger hclog.Logger

var once sync.Once

// GetLogger return new Logger
func GetLogger() hclog.Logger {
	once.Do(func() {
		loggerOption := &hclog.LoggerOptions{
			Name:            "Ozon",
			Level:           hclog.Info,
			Output:          os.Stderr,
			IncludeLocation: true,
			TimeFormat:      "2006-01-02 15:04:05.000",
		}
		logger = hclog.New(loggerOption)
	})
	return logger
}
