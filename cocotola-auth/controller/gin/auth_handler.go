package controller

import (
	"context"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	rsliblog "github.com/kujilabo/cocotola-1.23/redstart/lib/log"
	rsuserdomain "github.com/kujilabo/cocotola-1.23/redstart/user/domain"

	libapi "github.com/kujilabo/cocotola-1.23/lib/api"
	libcontroller "github.com/kujilabo/cocotola-1.23/lib/controller/gin"

	"github.com/kujilabo/cocotola-1.23/cocotola-auth/domain"
)

type AuthenticationUsecase interface {
	SignInWithIDToken(ctx context.Context, idToken string) (*domain.AuthTokenSet, error)
	GetUserInfo(ctx context.Context, bearerToken string) (*rsuserdomain.AppUserModel, error)
	RefreshToken(ctx context.Context, refreshToken string) (string, error)
}

type AuthHandler struct {
	authenticationUsecase AuthenticationUsecase
	logger                *slog.Logger
}

func NewAuthHandler(authenticationUsecase AuthenticationUsecase) *AuthHandler {
	return &AuthHandler{
		authenticationUsecase: authenticationUsecase,
		logger:                slog.Default().With(slog.String(rsliblog.LoggerNameKey, "AuthHandler")),
	}
}

func (h *AuthHandler) SignInWithIDToken(c *gin.Context) {
	ctx := c.Request.Context()
	authorization := c.GetHeader("Authorization")
	if !strings.HasPrefix(authorization, "Bearer ") {
		h.logger.InfoContext(ctx, "invalid header. Bearer not found")
		c.JSON(http.StatusUnauthorized, gin.H{"message": http.StatusText(http.StatusUnauthorized)})
		return
	}

	idToken := authorization[len("Bearer "):]
	tokenSet, err := h.authenticationUsecase.SignInWithIDToken(ctx, idToken)
	if err != nil {
		h.logger.InfoContext(ctx, "SignInWithIDToken", slog.Any("err", (err)))
		c.JSON(http.StatusUnauthorized, gin.H{"message": http.StatusText(http.StatusUnauthorized)})
		return
	}

	c.JSON(http.StatusOK, libapi.AuthResponse{
		AccessToken:  &tokenSet.AccessToken,
		RefreshToken: &tokenSet.RefreshToken,
	})
}

func (h *AuthHandler) GetUserInfo(c *gin.Context) {
	ctx := c.Request.Context()

	h.logger.InfoContext(ctx, "GetUserInfo")

	authorization := c.GetHeader("Authorization")
	if !strings.HasPrefix(authorization, "Bearer ") {
		h.logger.InfoContext(ctx, "invalid header. Bearer not found")
		c.JSON(http.StatusUnauthorized, gin.H{"message": http.StatusText(http.StatusUnauthorized)})
		return
	}

	bearerToken := authorization[len("Bearer "):]
	appUserInfo, err := h.authenticationUsecase.GetUserInfo(ctx, bearerToken)
	if err != nil {
		h.logger.InfoContext(ctx, "GetUserInfo", slog.Any("err", (err)))
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

	h.logger.InfoContext(ctx, "Authorize")
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

func NewInitAuthRouterFunc(authenticationUsecase AuthenticationUsecase) libcontroller.InitRouterGroupFunc {
	return func(parentRouterGroup gin.IRouter, middleware ...gin.HandlerFunc) {
		for _, m := range middleware {
			parentRouterGroup.Use(m)
		}

		authHandler := NewAuthHandler(authenticationUsecase)
		parentRouterGroup.POST("refresh_token", authHandler.RefreshToken)
		parentRouterGroup.GET("userinfo", authHandler.GetUserInfo)
		parentRouterGroup.POST("sign_in_with_id_token", authHandler.SignInWithIDToken)
	}
}
