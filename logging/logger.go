package logging

import (
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Setup() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).With().Caller().Logger()

	zerolog.LevelFieldName = "severity"
	zerolog.TimestampFieldName = "timestamp"
	zerolog.TimeFieldFormat = time.RFC3339Nano

	// Set log level from environment variable, default to Info
	level := getLogLevel()
	zerolog.SetGlobalLevel(level)
	log.Debug().Str("level", level.String()).Msg("Log level set")
}

func getLogLevel() zerolog.Level {
	levelStr := strings.ToLower(os.Getenv("BAMBULAB_EXPORTER_LOG_LEVEL"))
	switch levelStr {
	case "trace":
		return zerolog.TraceLevel
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn", "warning":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "fatal":
		return zerolog.FatalLevel
	case "panic":
		return zerolog.PanicLevel
	default:
		return zerolog.InfoLevel // Default to Info level
	}
}
