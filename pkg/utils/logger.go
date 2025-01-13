package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	log "github.com/sirupsen/logrus"
)

type Logger struct {
	LogFilePath string
	LogLevel    string
	LogFormat   log.Formatter
}

func (l *Logger) InitConfig() {
	switch l.LogLevel {
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	case "INFO":
		log.SetLevel(log.InfoLevel)
	case "WARN":
		log.SetLevel(log.WarnLevel)
	case "ERROR":
		log.SetLevel(log.ErrorLevel)
	default:
		log.SetLevel(log.ErrorLevel)
	}

	// Enable reporting caller
	log.SetReportCaller(true)

	if l.LogFilePath != "" {

		logDir := filepath.Dir(l.LogFilePath)
		if err := os.MkdirAll(logDir, 0755); err != nil {
			log.Errorf("failed to create log directory: %+v", err)
		}
		// Configure file rotation
		file, err := rotatelogs.New(
			fmt.Sprintf("%s.%s", l.LogFilePath, "%Y-%m-%d.%H:%M:%S"),
			rotatelogs.WithLinkName(l.LogFilePath+".link"),
			rotatelogs.WithMaxAge(time.Hour*24*5),     // 5 days retention
			rotatelogs.WithRotationTime(time.Hour*24), // Rotate daily
			rotatelogs.WithRotationCount(7),           // Keep 7 log files
		)

		if err != nil {
			log.Errorf("error opening file: %v\n", err)
			return
		}

		// Set log output to both stdout and file
		mw := io.MultiWriter(os.Stdout, file)
		log.SetOutput(mw)
	} else {
		log.Info("Log file path is empty, no log file will be created")
		log.SetOutput(os.Stdout)
	}

	// Set log formatter
	log.SetFormatter(l.LogFormat)

	log.Infof("Log level: %s\n", l.LogLevel)
	log.Infof("Log file path: %s\n", l.LogFilePath)
}
