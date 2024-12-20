package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	rslibdomain "github.com/kujilabo/cocotola-1.23/redstart/lib/domain"
	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
	rsliblog "github.com/kujilabo/cocotola-1.23/redstart/lib/log"
)

func AppServerProcess(ctx context.Context, loggerKey rslibdomain.ContextKey, router http.Handler, port int, readHeaderTimeout time.Duration, shutdownTime time.Duration) error {
	ctx = rsliblog.WithLoggerName(ctx, loggerKey)
	logger := rsliblog.GetLoggerFromContext(ctx, loggerKey)

	httpServer := http.Server{
		Addr:              ":" + strconv.Itoa(port),
		Handler:           router,
		ReadHeaderTimeout: readHeaderTimeout,
	}

	logger.InfoContext(ctx, fmt.Sprintf("http server listening at %v", httpServer.Addr))

	errCh := make(chan error)
	go func() {
		defer close(errCh)
		if err := httpServer.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			logger.InfoContext(ctx, fmt.Sprintf("failed to ListenAndServe. err: %v", err))
			errCh <- err
		}
	}()

	select {
	case <-ctx.Done():
		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), shutdownTime)
		defer shutdownCancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			logger.InfoContext(ctx, fmt.Sprintf("Server forced to shutdown. err: %v", err))
			return rsliberrors.Errorf("httpServer.Shutdown. err: %w", err)
		}
		return nil
	case err := <-errCh:
		return rsliberrors.Errorf("httpServer.ListenAndServe. err: %w", err)
	}
}
