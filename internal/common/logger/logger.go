package logger

import (
	"io"
	"log"
	"os"
)

const (
	LogLevelInfo    = "INFO"
	LogLevelWarning = "WARNING"
	LogLevelError   = "ERROR"
	LogLevelDebug   = "DEBUG"
)

var (
	LoggerClient = newLogger(os.Stdout, LogLevelInfo)
)

type Logger struct {
	logger *log.Logger
	level  string
}

func newLogger(out io.Writer, level string) *Logger {
	return &Logger{
		logger: log.New(out, "", log.Ldate|log.Ltime),
		level:  level,
	}
}

func (l *Logger) Info(message string) {
	if l.level == LogLevelInfo {
		l.logger.Printf("[%s] %s\n", LogLevelInfo, message)
	}
}

func (l *Logger) Warning(message string) {
	if l.level == LogLevelInfo || l.level == LogLevelWarning {
		l.logger.Printf("[%s] %s\n", LogLevelWarning, message)
	}
}

func (l *Logger) Error(message string) {
	if l.level == LogLevelInfo || l.level == LogLevelWarning || l.level == LogLevelError {
		l.logger.Printf("[%s] %s\n", LogLevelError, message)
	}
}

func (l *Logger) Debug(message string) {
	if l.level == LogLevelDebug {
		l.logger.Printf("[%s] %s\n", LogLevelDebug, message)
	}
}
