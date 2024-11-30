package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	libapi "github.com/kujilabo/cocotola-1.23/lib/api"
	"github.com/kujilabo/cocotola-1.23/redstart/user/domain"
)

type SystemAdminInterface interface {
	AppUserID() *domain.AppUserID
	IsSystemAdmin() bool
	// GetUserGroups() []domain.UserGroupModel
}

type RBACUsecase interface {
	AddPolicyToUser(ctx context.Context, organizationID *domain.OrganizationID, subject domain.RBACSubject, action domain.RBACAction, object domain.RBACObject, effect domain.RBACEffect) error
}

type RBACHandler struct {
	rbacUsecase RBACUsecase
}

func NewRBACHandler(rbacUsecase RBACUsecase) *RBACHandler {
	return &RBACHandler{
		rbacUsecase: rbacUsecase,
	}
}

func (h *RBACHandler) AddPolicyToUser(c *gin.Context) {
	ctx := c.Request.Context()
	apiParam := libapi.AddPolicyToUserParameter{}
	if err := c.ShouldBindJSON(&apiParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})
		return
	}

	organizationID, err := domain.NewOrganizationID(apiParam.OrganizationID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})
		return
	}

	subject := domain.NewRBACUser(apiParam.Subject)
	action := domain.NewRBACAction(apiParam.Action)
	object := domain.NewRBACObject(apiParam.Object)
	effect := domain.NewRBACEffect(apiParam.Effect)

	if err := h.rbacUsecase.AddPolicyToUser(ctx, organizationID, subject, action, object, effect); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": http.StatusText(http.StatusInternalServerError)})
		return
	}
}

func (h *RBACHandler) AddPolicyToGroup(c *gin.Context) {

}

func NewInitRBACRouterFunc(rbacUsecase RBACUsecase) InitRouterGroupFunc {
	return func(parentRouterGroup gin.IRouter, middleware ...gin.HandlerFunc) error {
		rbac := parentRouterGroup.Group("rbac")
		for _, m := range middleware {
			rbac.Use(m)
		}

		rbacHandler := NewRBACHandler(rbacUsecase)
		rbac.PUT("policy/user", rbacHandler.AddPolicyToUser)
		rbac.PUT("policy/group", rbacHandler.AddPolicyToGroup)

		return nil
	}
}
