package helper

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	rsliblog "github.com/kujilabo/cocotola-1.23/redstart/lib/log"
	rsuserdomain "github.com/kujilabo/cocotola-1.23/redstart/user/domain"

	"github.com/kujilabo/cocotola-1.23/cocotola-core/service"
)

type operator struct {
	appUserID      *rsuserdomain.AppUserID
	organizationID *rsuserdomain.OrganizationID
}

func (o *operator) AppUserID() *rsuserdomain.AppUserID {
	return o.appUserID
}
func (o *operator) OrganizationID() *rsuserdomain.OrganizationID {
	return o.organizationID
}

func HandleSecuredFunction(c *gin.Context, fn func(ctx context.Context, operator service.OperatorInterface) error, errorHandle func(ctx context.Context, c *gin.Context, err error) bool) {
	ctx := c.Request.Context()
	logger := slog.Default().With(slog.String(rsliblog.LoggerNameKey, "HandleSecuredFunction"))

	organizationIDInt := c.GetInt("OrganizationID")
	if organizationIDInt == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"message": http.StatusText(http.StatusUnauthorized)})
		return
	}

	organizationID, err := rsuserdomain.NewOrganizationID(organizationIDInt)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": http.StatusText(http.StatusUnauthorized)})
		return
	}

	appUserID := c.GetInt("AuthorizedUser")
	if appUserID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"message": http.StatusText(http.StatusUnauthorized)})
		return
	}

	operatorID, err := rsuserdomain.NewAppUserID(appUserID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": http.StatusText(http.StatusUnauthorized)})
		return
	}

	logger.InfoContext(ctx, "", slog.Int("organization_id", organizationID.Int()), slog.Int("operator_id", operatorID.Int()))

	operator := &operator{
		appUserID:      operatorID,
		organizationID: organizationID,
	}

	if err := fn(ctx, operator); err != nil {
		if handled := errorHandle(ctx, c, err); !handled {
			c.JSON(http.StatusInternalServerError, gin.H{"message": http.StatusText(http.StatusInternalServerError)})
		}
	}
}
