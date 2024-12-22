package usecase

import (
	"go.opentelemetry.io/otel"
)

// const (
// 	loggerKey = liblog.TatoebaUsecaseLoggerContextKey
// )

var (
	tracer = otel.Tracer("github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/usecase")
)
