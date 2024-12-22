package config

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
)

type LogConfig struct {
	Level map[string]string `yaml:"level"`
}

func InitLog(cfg *LogConfig) error {
	defaultLogLevel := slog.LevelWarn
	if rootLevel, ok := cfg.Level["default"]; ok {
		defaultLogLevel = stringToLogLevel(rootLevel)
	}

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: defaultLogLevel,
	})))

	return nil
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
