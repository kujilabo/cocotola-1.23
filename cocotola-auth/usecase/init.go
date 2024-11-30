package usecase

import (
	"context"

	rsuserservice "github.com/kujilabo/cocotola-1.23/redstart/user/service"

	"github.com/kujilabo/cocotola-1.23/cocotola-auth/service"

	liblog "github.com/kujilabo/cocotola-1.23/lib/log"
)

const (
	loggerKey = liblog.AuthUsecaseLoggerContextKey
)

type SystemOwnerByOrganizationName interface {
	Get(ctx context.Context, rf service.RepositoryFactory, organizationName string) (*rsuserservice.SystemOwner, error)
}
