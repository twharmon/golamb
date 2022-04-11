package golamb

import (
	"log"
	"math"
	"os"
)

// LogLevel as defined in the RFC 5424 specification.
type LogLevel int

const (
	// LogLevelDebug as defined in the RFC 5424 specification.
	LogLevelDebug LogLevel = 0

	// LogLevelInfo as defined in the RFC 5424 specification.
	LogLevelInfo LogLevel = 1

	// LogLevelNotice as defined in the RFC 5424 specification.
	LogLevelNotice LogLevel = 2

	// LogLevelWarning as defined in the RFC 5424 specification.
	LogLevelWarning LogLevel = 3

	// LogLevelError as defined in the RFC 5424 specification.
	LogLevelError LogLevel = 4

	// LogLevelCritical as defined in the RFC 5424 specification.
	LogLevelCritical LogLevel = 5

	// LogLevelAlert as defined in the RFC 5424 specification.
	LogLevelAlert LogLevel = 6

	// LogLevelEmergency as defined in the RFC 5424 specification.
	LogLevelEmergency LogLevel = 7

	// LogLevelSilent logs nothing.
	LogLevelSilent LogLevel = math.MaxInt
)

func (l LogLevel) String() string {
	switch l {
	case LogLevelDebug:
		return "DEBUG"
	case LogLevelInfo:
		return "INFO"
	case LogLevelNotice:
		return "NOTICE"
	case LogLevelWarning:
		return "WARNING"
	case LogLevelError:
		return "ERROR"
	case LogLevelCritical:
		return "CRITICAL"
	case LogLevelAlert:
		return "ALERT"
	case LogLevelEmergency:
		return "EMERGENCY"
	case LogLevelSilent:
		return "SILENT"
	}
	return "NONE"
}

// Logger is used to log messages.
type Logger interface {
	Log(level LogLevel, message string)
}

// DefaultLogger logs messages to os.Stdout, which is sent to
// CloudWatch logs.
type DefaultLogger struct {
	logger *log.Logger
}

// NewDefaultLogger returns the default logger.
func NewDefaultLogger() Logger {
	return &DefaultLogger{
		logger: log.New(os.Stdout, "", 0),
	}
}

// Log implements the Logger interface.
func (l *DefaultLogger) Log(level LogLevel, message string) {
	l.logger.Printf("%s\n", message)
}
