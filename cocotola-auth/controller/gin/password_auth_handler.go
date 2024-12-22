package handler

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	rsliblog "github.com/kujilabo/cocotola-1.23/redstart/lib/log"

	libapi "github.com/kujilabo/cocotola-1.23/lib/api"

	"github.com/kujilabo/cocotola-1.23/cocotola-auth/domain"
)

type PasswordUsecaseInterface interface {
	Authenticate(ctx context.Context, loginID, password, organizationName string) (*domain.AuthTokenSet, error)
}

type PasswordAuthHandler struct {
	passwordUsecase PasswordUsecaseInterface
	logger          *slog.Logger
}

func NewPasswordAuthHandler(passwordUsecase PasswordUsecaseInterface) *PasswordAuthHandler {
	return &PasswordAuthHandler{
		passwordUsecase: passwordUsecase,
		logger:          slog.Default().With(slog.String(rsliblog.LoggerNameKey, "PasswordAuthHandler")),
	}
}

func (h *PasswordAuthHandler) Authorize(c *gin.Context) {
	ctx := c.Request.Context()

	passwordAuthParameter := libapi.PasswordAuthParameter{}
	if err := c.ShouldBindJSON(&passwordAuthParameter); err != nil {
		h.logger.InfoContext(ctx, fmt.Sprintf("invalid parameter. err: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})
		return
	}

	authResult, err := h.passwordUsecase.Authenticate(ctx, passwordAuthParameter.LoginID, passwordAuthParameter.Password, passwordAuthParameter.OrganizationName)
	if err != nil {
		if errors.Is(err, domain.ErrUnauthenticated) {
			h.logger.InfoContext(ctx, fmt.Sprintf("invalid parameter. err: %v", err))
			c.JSON(http.StatusUnauthorized, gin.H{"message": http.StatusText(http.StatusUnauthorized)})
			return
		}

		h.logger.ErrorContext(ctx, fmt.Sprintf("passwordUsecase.Authenticate. err: %+v", err))
		c.JSON(http.StatusInternalServerError, gin.H{"message": http.StatusText(http.StatusInternalServerError)})
		return
	}

	c.JSON(http.StatusOK, libapi.AuthResponse{
		AccessToken:  &authResult.AccessToken,
		RefreshToken: &authResult.RefreshToken,
	})
}

func NewInitPasswordRouterFunc(password PasswordUsecaseInterface) InitRouterGroupFunc {
	return func(parentRouterGroup gin.IRouter, middleware ...gin.HandlerFunc) error {
		auth := parentRouterGroup.Group("password")
		for _, m := range middleware {
			auth.Use(m)
		}

		passwordAuthHandler := NewPasswordAuthHandler(password)
		auth.POST("authenticate", passwordAuthHandler.Authorize)
		return nil
	}
}
