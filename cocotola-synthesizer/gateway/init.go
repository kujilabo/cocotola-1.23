package gateway

import (
	"go.opentelemetry.io/otel"
)

var (
	tracer = otel.Tracer("github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/gateway")
)
