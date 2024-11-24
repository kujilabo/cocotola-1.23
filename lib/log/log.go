package log

import (
	"context"
	"log/slog"

	rslibdomain "github.com/kujilabo/cocotola-1.23/redstart/lib/domain"
	rsliblog "github.com/kujilabo/cocotola-1.23/redstart/lib/log"
)

const (
	AppControllerLoggerContextKey         rslibdomain.ContextKey = "app_controller"
	AppGatewayLoggerContextKey            rslibdomain.ContextKey = "app_gateway"
	AppInitializeLoggerContextKey         rslibdomain.ContextKey = "app_initialize"
	AppMainLoggerContextKey               rslibdomain.ContextKey = "app_main"
	AppServiceLoggerContextKey            rslibdomain.ContextKey = "app_service"
	AppUsecaseLoggerContextKey            rslibdomain.ContextKey = "app_usecase"
	AuthControllerLoggerContextKey        rslibdomain.ContextKey = "auth_controller"
	AuthGatewayLoggerContextKey           rslibdomain.ContextKey = "auth_gateway"
	AuthInitializeLoggerContextKey        rslibdomain.ContextKey = "auth_initialize"
	AuthMainLoggerContextKey              rslibdomain.ContextKey = "auth_main"
	AuthServiceLoggerContextKey           rslibdomain.ContextKey = "auth_service"
	AuthUsecaseLoggerContextKey           rslibdomain.ContextKey = "auth_usecase"
	CoreControllerLoggerContextKey        rslibdomain.ContextKey = "core_controller"
	CoreGatewayLoggerContextKey           rslibdomain.ContextKey = "core_gateway"
	CoreInitializeLoggerContextKey        rslibdomain.ContextKey = "core_initialize"
	CoreMainLoggerContextKey              rslibdomain.ContextKey = "core_main"
	CoreServiceLoggerContextKey           rslibdomain.ContextKey = "core_service"
	CoreUsecaseLoggerContextKey           rslibdomain.ContextKey = "core_usecase"
	LibControllerLoggerContextKey         rslibdomain.ContextKey = "lib_controller"
	LibGatewayLoggerContextKey            rslibdomain.ContextKey = "lib_gateway"
	SynthesizerControllerLoggerContextKey rslibdomain.ContextKey = "synthesizer_controller"
	SynthesizerGatewayLoggerContextKey    rslibdomain.ContextKey = "synthesizer_gateway"
	SynthesizerInitializeLoggerContextKey rslibdomain.ContextKey = "synthesizer_initialize"
	SynthesizerMainLoggerContextKey       rslibdomain.ContextKey = "synthesizer_main"
	SynthesizerServiceLoggerContextKey    rslibdomain.ContextKey = "synthesizer_service"
	SynthesizerUsecaseLoggerContextKey    rslibdomain.ContextKey = "synthesizer_usecase"
)

var (
	LoggerKeys = []rslibdomain.ContextKey{
		AppControllerLoggerContextKey,
		AppGatewayLoggerContextKey,
		AppInitializeLoggerContextKey,
		AppMainLoggerContextKey,
		AppServiceLoggerContextKey,
		AppUsecaseLoggerContextKey,
		AuthControllerLoggerContextKey,
		AuthGatewayLoggerContextKey,
		AuthInitializeLoggerContextKey,
		AuthMainLoggerContextKey,
		AuthServiceLoggerContextKey,
		AuthUsecaseLoggerContextKey,
		CoreControllerLoggerContextKey,
		CoreGatewayLoggerContextKey,
		CoreInitializeLoggerContextKey,
		CoreMainLoggerContextKey,
		CoreServiceLoggerContextKey,
		CoreUsecaseLoggerContextKey,
		LibControllerLoggerContextKey,
		LibGatewayLoggerContextKey,
		SynthesizerControllerLoggerContextKey,
		SynthesizerGatewayLoggerContextKey,
		SynthesizerInitializeLoggerContextKey,
		SynthesizerMainLoggerContextKey,
		SynthesizerServiceLoggerContextKey,
		SynthesizerUsecaseLoggerContextKey,
	}
)

func InitLogger(ctx context.Context) context.Context {
	for _, key := range LoggerKeys {
		if _, ok := rsliblog.Loggers[key]; !ok {
			rsliblog.Loggers[key] = slog.New(rsliblog.LogHandlers[rsliblog.DefaultLogLevel])
		}
		ctx = context.WithValue(ctx, key, rsliblog.Loggers[key])
	}

	return ctx
}
