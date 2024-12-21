package middleware

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"

	rsliblog "github.com/kujilabo/cocotola-1.23/redstart/lib/log"
)

func NewTraceLogMiddleware(appName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		sc := trace.SpanFromContext(c.Request.Context()).SpanContext()
		if !sc.TraceID().IsValid() || !sc.SpanID().IsValid() {
			return
		}
		otTraceID := sc.TraceID().String()

		ctx := c.Request.Context()
		savedCtx := ctx
		defer func() {
			c.Request = c.Request.WithContext(savedCtx)
		}()

		// ctx = rsliblog.WithLoggerName(ctx, loggerKey)
		// logger := rsliblog.GetLoggerFromContext(ctx, loggerKey)
		logger := slog.Default().With(slog.String(rsliblog.LoggerNameKey, "TraceLogMiddleware"))
		logger.InfoContext(ctx, "", slog.String("uri", c.Request.RequestURI), slog.String("method", c.Request.Method), slog.String("trace_id", otTraceID))

		ctx, span := tracer.Start(ctx, "TraceLog")
		defer span.End()

		c.Request = c.Request.WithContext(ctx)

		// serve the request to the next middleware
		c.Next()
	}
}
