package logger

import (
	"fmt"
	"os"
	"time"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

var levelNames = map[LogLevel]string{
	DEBUG: "DEBUG",
	INFO:  "INFO",
	WARN:  "WARN",
	ERROR: "ERROR",
	FATAL: "FATAL",
}

type Logger struct {
	level LogLevel
}

func NewLogger(level LogLevel) *Logger {
	return &Logger{level: level}
}

// Formata a mensagem de log com timestamp e nivel
func (l *Logger) formatLog(level LogLevel, message string, args ...interface{}) string {
	timestamp := time.Now().Format("01-02-2006 15:04:05")
	levelName := levelNames[level]

	var formattedMsg string
	if len(args) > 0 {
		formattedMsg = fmt.Sprintf(message, args...)
	} else {
		formattedMsg = message
	}

	return fmt.Sprintf("[%s] %s: %s", timestamp, levelName, formattedMsg)
}

// Regristra Mensagem de Debug
func (l *Logger) Debug(message string, args ...interface{}) {
	if l.level <= DEBUG {
		fmt.Println(l.formatLog(DEBUG, message, args...))
	}
}

// Regristra Mensagem de Debug
func (l *Logger) Info(message string, args ...interface{}) {
	if l.level <= INFO {
		fmt.Println(l.formatLog(INFO, message, args...))
	}
}

// Regristra Mensagem de Debug
func (l *Logger) Warm(message string, args ...interface{}) {
	if l.level <= WARN {
		fmt.Println(l.formatLog(WARN, message, args...))
	}
}

// Regristra Mensagem de Debug
func (l *Logger) Error(message string, args ...interface{}) {
	if l.level <= ERROR {
		fmt.Fprintf(os.Stderr, "%s\n", l.formatLog(ERROR, message, args...))
	}
}

// Regristra Mensagem de Debug
func (l *Logger) Fatal(message string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "%s\n", l.formatLog(FATAL, message, args...))
	os.Exit(1)
}

// Retorna o nivel de log configurado
func (l *Logger) GetLogLevel() LogLevel {
	return l.level
}

// Define novo nivel de log
func (l *Logger) SetLogLevel(level LogLevel) {
	l.level = level
}
