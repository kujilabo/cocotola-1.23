package helper

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	rsliblog "github.com/kujilabo/cocotola-1.23/redstart/lib/log"
)

func HandleFunction(c *gin.Context, fn func(ctx context.Context, logger *slog.Logger) error, errorHandle func(ctx context.Context, logger *slog.Logger, c *gin.Context, err error) bool) {
	ctx := c.Request.Context()
	ctx = rsliblog.WithLoggerName(ctx, loggerKey)

	controllerLogger := rsliblog.GetLoggerFromContext(ctx, loggerKey)
	if err := fn(ctx, controllerLogger); err != nil {
		if handled := errorHandle(ctx, controllerLogger, c, err); !handled {
			c.JSON(http.StatusInternalServerError, gin.H{"message": http.StatusText(http.StatusInternalServerError)})
		}
	}
}
