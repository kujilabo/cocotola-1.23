package service

import (
	"context"

	"github.com/kujilabo/cocotola-1.23/cocotola-auth/domain"
	rsuserdomain "github.com/kujilabo/cocotola-1.23/redstart/user/domain"
)

type AppUserInfo struct {
	LoginID          string
	AppUserID        int
	Username         string
	OrganizationID   int
	OrganizationName string
}

type AppUserInterface interface {
	AppUserID() *rsuserdomain.AppUserID
	OrganizationID() *rsuserdomain.OrganizationID
	LoginID() string
	Username() string
	// GetUserGroups() []domain.UserGroupModel
}

type OrganizationInterface interface {
	OrganizationID() *rsuserdomain.OrganizationID
	Name() string
}

type AuthTokenManager interface {
	GetUserInfo(ctx context.Context, tokenString string) (*AppUserInfo, error)

	CreateTokenSet(ctx context.Context, appUser AppUserInterface, organizationUsecase OrganizationInterface) (*domain.AuthTokenSet, error)
	RefreshToken(ctx context.Context, accessToken string) (string, error)
}
