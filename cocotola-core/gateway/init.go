package gateway

import (
	"go.opentelemetry.io/otel"
	// liblog "github.com/kujilabo/cocotola-1.23/lib/log"
)

const (
// loggerKey = liblog.CoreGatewayLoggerContextKey
)

var (
	tracer = otel.Tracer("github.com/kujilabo/cocotola-1.23/cocotola-core/gateway")
)
