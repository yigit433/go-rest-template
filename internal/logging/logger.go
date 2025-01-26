package logging

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
)

type Logger struct {
	logger zerolog.Logger
}

var CustomLogger *Logger

func InitializeLogger(level string, appName string) {
	CustomLogger = NewLogger(level, appName)
}

func NewLogger(level string, appName string) *Logger {
	consoleWriter := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
		FormatMessage: func(i interface{}) string {
			return fmt.Sprintf("[App: %s] %v", appName, i)
		},
	}

	zerolog.SetGlobalLevel(parseLogLevel(level))
	logger := zerolog.New(consoleWriter).With().
		Timestamp().
		Logger()

	return &Logger{logger: logger}
}

func (l *Logger) Log(level zerolog.Level, message string, fields map[string]interface{}) {
	event := l.logger.WithLevel(level).Str("message", message)

	for k, v := range fields {
		event = event.Interface(k, v)
	}

	event.Msg(message)
}

func (l *Logger) Debug(message string, fields map[string]interface{}) {
	l.Log(zerolog.DebugLevel, message, fields)
}

func (l *Logger) Info(message string, fields map[string]interface{}) {
	l.Log(zerolog.InfoLevel, message, fields)
}

func (l *Logger) Error(message string, fields map[string]interface{}) {
	l.Log(zerolog.ErrorLevel, message, fields)
}