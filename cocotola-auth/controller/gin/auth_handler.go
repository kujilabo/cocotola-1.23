package handler

import (
	"context"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	rsliblog "github.com/kujilabo/cocotola-1.23/redstart/lib/log"
	rsuserdomain "github.com/kujilabo/cocotola-1.23/redstart/user/domain"

	libapi "github.com/kujilabo/cocotola-1.23/lib/api"
)

type AuthenticationUsecase interface {
	GetUserInfo(ctx context.Context, bearerToken string) (*rsuserdomain.AppUserModel, error)
	RefreshToken(ctx context.Context, refreshToken string) (string, error)
}

type AuthHandler struct {
	authenticationUsecase AuthenticationUsecase
}

func NewAuthHandler(authenticationUsecase AuthenticationUsecase) *AuthHandler {
	return &AuthHandler{
		authenticationUsecase: authenticationUsecase,
	}
}

func (h *AuthHandler) GetUserInfo(c *gin.Context) {
	ctx := c.Request.Context()
	ctx = rsliblog.WithLoggerName(ctx, loggerKey)
	logger := rsliblog.GetLoggerFromContext(ctx, loggerKey)

	logger.InfoContext(ctx, "GetUserInfo")

	authorization := c.GetHeader("Authorization")
	if !strings.HasPrefix(authorization, "Bearer ") {
		logger.InfoContext(ctx, "invalid header. Bearer not found")
		c.JSON(http.StatusUnauthorized, gin.H{"message": http.StatusText(http.StatusUnauthorized)})
		return
	}

	bearerToken := authorization[len("Bearer "):]
	appUserInfo, err := h.authenticationUsecase.GetUserInfo(ctx, bearerToken)
	if err != nil {
		logger.InfoContext(ctx, "GetUserInfo", slog.Any("err", (err)))
		c.JSON(http.StatusUnauthorized, gin.H{"message": http.StatusText(http.StatusUnauthorized)})
		return
	}

	c.JSON(http.StatusOK, libapi.AppUserInfoResponse{
		AppUserID:      appUserInfo.AppUserID.Int(),
		OrganizationID: appUserInfo.OrganizationID.Int(),
		LoginID:        appUserInfo.LoginID,
		Username:       appUserInfo.Username,
	})
	// TODO: check if the token is registered
}

func (h *AuthHandler) RefreshToken(c *gin.Context) {
	ctx := c.Request.Context()
	ctx = rsliblog.WithLoggerName(ctx, loggerKey)
	logger := rsliblog.GetLoggerFromContext(ctx, loggerKey)

	logger.InfoContext(ctx, "Authorize")
	refreshTokenParameter := libapi.RefreshTokenParameter{}
	if err := c.ShouldBindJSON(&refreshTokenParameter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})
		return
	}

	accessToken, err := h.authenticationUsecase.RefreshToken(ctx, refreshTokenParameter.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": http.StatusText(http.StatusUnauthorized)})
		return
	}

	c.JSON(http.StatusOK, libapi.AuthResponse{
		AccessToken: &accessToken,
	})
}

func NewInitAuthRouterFunc(authenticationUsecase AuthenticationUsecase) InitRouterGroupFunc {
	return func(parentRouterGroup gin.IRouter, middleware ...gin.HandlerFunc) error {
		auth := parentRouterGroup.Group("auth")
		for _, m := range middleware {
			auth.Use(m)
		}

		authHandler := NewAuthHandler(authenticationUsecase)
		auth.POST("refresh_token", authHandler.RefreshToken)
		auth.GET("userinfo", authHandler.GetUserInfo)
		return nil
	}
}
