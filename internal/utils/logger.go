package utils

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

func NewLogger() *Logger {

	logFile, err := os.OpenFile("agent.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}

	multiWriter := io.MultiWriter(os.Stdout, logFile)

	return &Logger{
		infoLog:  log.New(multiWriter, "[INFO] ", log.Ldate|log.Ltime),
		errorLog: log.New(multiWriter, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *Logger) Info(msg string) {
	l.infoLog.Println(msg)
}

func (l *Logger) Error(msg string) {
	l.errorLog.Println(msg)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.infoLog.Printf(format, args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.errorLog.Printf(format, args...)
}