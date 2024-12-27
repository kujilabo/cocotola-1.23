package main

import (
	"flag"
	"log"
	"os"

	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
	// rslibconfig "github.com/kujilabo/cocotola-1.23/redstart/lib/config"
	rslibgateway "github.com/kujilabo/cocotola-1.23/redstart/lib/gateway"
)

func getValue(values ...string) string {
	for _, v := range values {
		if len(v) != 0 {
			return v
		}
	}
	return ""
}

// func checkError(err error) {
// 	if err != nil {
// 		log.Fatalf("err: %v", err)
// 	}
// }

func main() {
	log.Println("Hello, World!!!!!!!!")

	// ctx := context.Background()
	env := flag.String("env", "", "environment")
	flag.Parse()
	appEnv := getValue(*env, os.Getenv("APP_ENV"), "local")

	rsliberrors.UseXerrorsErrorf()

	// // load config
	// cfg, err := config.LoadConfig(appEnv)
	// checkError(err)

	var _ = rslibgateway.MYSQL_ER_DUP_ENTRY
	// var _ = rslibconfig.DBConfig{}
	// var _ = ctx
	// var _ = cfg
	var _ = appEnv

	// // init log
	// rslibconfig.InitLog(cfg.Log)
	// logger := slog.Default().With(slog.String(rsliblog.LoggerNameKey, "main"))
	// logger.InfoContext(ctx, fmt.Sprintf("env: %s", appEnv))

	// // init tracer
	// tp, err := rslibconfig.InitTracerProvider(ctx, cfg.App.Name, cfg.Trace)
	// checkError(err)
	// otel.SetTracerProvider(tp)
	// otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
}
