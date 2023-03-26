package logger

import (
	"log"
	"os"
	"time"

	"github.com/ginger-go/env"
	"github.com/ginger-go/ginger"
)

func Info(msg ...any) {
	refreshFileLogger()
	stdLogger.Println("[INFO]", msg)
	fileLogger.Println("[INFO]", msg)
}

func Warn(msg ...any) {
	stdLogger.Println("[WARN]", msg)
}

func Err(msg ...any) {
	refreshFileLogger()
	stdLogger.Println("[ERR]", msg)
	fileLogger.Println("[ERR]", msg)
}

func Debug(msg ...any) {
	if env.GetEnv("GIN_MODE", "") != ginger.GIN_MODE_RELEASE {
		stdLogger.Println("[DEBUG]", msg)
	}
}

var lastFileName = ""
var stdLogger = getStdLogger()
var fileLogger *log.Logger

func refreshFileLogger() {
	filename := getLogFileName()
	if filename != lastFileName {
		fileLogger = getFileLogger(filename)
		lastFileName = filename
	}
}

func getLogFileName() string {
	now := time.Now()
	return now.Format("2006-01-02") + ".log"
}

func getStdLogger() *log.Logger {
	return log.New(os.Stdout, "", log.LstdFlags)
}

func getFileLogger(filename string) *log.Logger {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	return log.New(file, "", log.LstdFlags)
}
