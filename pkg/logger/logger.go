package logger

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

type Logger struct {
	zerolog.Logger
}

func NewLogger(level string, format string) *Logger {
	var output io.Writer = os.Stdout

	// Log Level
	logLevel, error := zerolog.ParseLevel(level)
	if error != nil {
		logLevel = zerolog.InfoLevel
	}

	// Format Config
	if format == "json" {
		output = os.Stdout
	} else {
		output = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}
	}

	logger := zerolog.New(output).
		Level(logLevel).
		With().
		Timestamp().
		Caller().
		Logger()

	return &Logger{logger}
}

// helpers
func (logger *Logger) Info() *zerolog.Event {
	return logger.Logger.Info()
}

func (logger *Logger) APIRequest(method, path string, status int, duration time.Duration) {
	logger.Info().
		Str("method", method).
		Str("path", path).
		Int("status", status).
		Dur("duration", duration).
		Msg("api_request")
}

func (logger *Logger) DatabaseQeury(query string, duration time.Duration) {
	logger.Info().
		Str("query", query).
		Dur("duration", duration).
		Msg("database_query")
}

func (logger *Logger) AuthEvent(userID int, event string) {
	logger.Info().
		Int("user_id", userID).
		Str("event", event).
		Msg("auth_event")
}
