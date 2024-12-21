package helper

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleFunction(c *gin.Context, fn func(ctx context.Context) error, errorHandle func(ctx context.Context, c *gin.Context, err error) bool) {
	ctx := c.Request.Context()
	// ctx = rsliblog.WithLoggerName(ctx, loggerKey)

	// controllerLogger := rsliblog.GetLoggerFromContext(ctx, loggerKey)
	if err := fn(ctx); err != nil {
		if handled := errorHandle(ctx, c, err); !handled {
			c.JSON(http.StatusInternalServerError, gin.H{"message": http.StatusText(http.StatusInternalServerError)})
		}
	}
}
