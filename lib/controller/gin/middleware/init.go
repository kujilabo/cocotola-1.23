package middleware

import (
	"go.opentelemetry.io/otel"

	liblog "github.com/kujilabo/cocotola-1.23/lib/log"
)

const (
	loggerKey = liblog.LibControllerLoggerContextKey
)

var (
	tracer = otel.Tracer("github.com/kujilabo/cocotola-1.23/lib/controller/gin/middleware")
)
