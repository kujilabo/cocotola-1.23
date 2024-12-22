package middleware

import (
	"log/slog"
	"strings"

	"github.com/gin-gonic/gin"

	rsliblog "github.com/kujilabo/cocotola-1.23/redstart/lib/log"

	"github.com/kujilabo/cocotola-1.23/cocotola-core/service"
)

func NewAuthMiddleware(cocotolaAuthClient service.CocotolaAuthClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		ctx, span := tracer.Start(ctx, "AuthMiddleware")
		defer span.End()

		logger := slog.Default().With(slog.String(rsliblog.LoggerNameKey, "AuthMiddleware"))

		authorization := c.GetHeader("Authorization")
		if !strings.HasPrefix(authorization, "Bearer ") {
			logger.InfoContext(ctx, "invalid header. Bearer not found")
			return
		}

		bearerToken := authorization[len("Bearer "):]
		appUserInfo, err := cocotolaAuthClient.RetrieveUserInfo(ctx, bearerToken)
		if err != nil {
			logger.WarnContext(ctx, "getUserInfo")
			return
		}

		c.Set("AuthorizedUser", appUserInfo.AppUserID)
		c.Set("OrganizationID", appUserInfo.OrganizationID)

		logger.WarnContext(ctx, "authenticated", slog.Int("app_user_id", appUserInfo.AppUserID), slog.Int("organization_id", appUserInfo.OrganizationID))
	}
}
