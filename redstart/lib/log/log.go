package log

const (
	LoggerNameKey = "logger_name"
)

// import (
// 	"context"
// 	"fmt"
// 	"log/slog"
// 	"os"
// 	"sync"

// 	"github.com/kujilabo/cocotola-1.23/redstart/lib/domain"
// )

// var (
// 	lock            sync.Mutex
// 	DefaultLogger   *slog.Logger
// 	DefaultLogLevel slog.Level
// 	LogHandlers     map[slog.Level]slog.Handler        = make(map[slog.Level]slog.Handler)
// 	Loggers         map[domain.ContextKey]*slog.Logger = make(map[domain.ContextKey]*slog.Logger)
// )

// func init() {
// 	for _, level := range []slog.Level{slog.LevelDebug, slog.LevelInfo, slog.LevelWarn, slog.LevelError} {
// 		LogHandlers[level] = &LogHandler{Handler: slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
// 			Level: level,
// 		})}
// 	}

// 	DefaultLogger = slog.New(LogHandlers[slog.LevelWarn])
// }

// func WithLoggerName(ctx context.Context, val domain.ContextKey) context.Context {
// 	return context.WithValue(ctx, LoggerNameContextKey, string(val))
// }

// // GetLoggerFromContext Gets the logger from context
// func GetLoggerFromContext(ctx context.Context, key domain.ContextKey) *slog.Logger {
// 	logger, ok := ctx.Value(key).(*slog.Logger)
// 	if ok {
// 		return logger
// 	}

// 	lock.Lock()
// 	defer lock.Unlock()

// 	if _, ok := Loggers[key]; !ok {
// 		DefaultLogger.DebugContext(ctx, fmt.Sprintf("logger not found. logger: %s", key))
// 		return DefaultLogger
// 	}

// 	return Loggers[key]
// }
