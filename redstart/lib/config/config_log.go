package config

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
)

type LogConfig struct {
	Level string `yaml:"level"`
}

func InitLog(cfg *LogConfig) {
	defaultLogLevel := stringToLogLevel(cfg.Level)

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: defaultLogLevel,
	})))
}

func stringToLogLevel(value string) slog.Level {
	switch strings.ToLower(value) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		slog.Info(fmt.Sprintf("Unsupported log level: %s", value))
		return slog.LevelWarn
	}
}
