package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	rsuserdomain "github.com/kujilabo/cocotola-1.23/redstart/user/domain"

	libapi "github.com/kujilabo/cocotola-1.23/lib/api"
	libcontroller "github.com/kujilabo/cocotola-1.23/lib/controller/gin"
)

type SystemAdminInterface interface {
	AppUserID() *rsuserdomain.AppUserID
	IsSystemAdmin() bool
	// GetUserGroups() []domain.UserGroupModel
}

type RBACUsecase interface {
	AddPolicyToUser(ctx context.Context, organizationID *rsuserdomain.OrganizationID, subject rsuserdomain.RBACSubject, action rsuserdomain.RBACAction, object rsuserdomain.RBACObject, effect rsuserdomain.RBACEffect) error
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

	organizationID, err := rsuserdomain.NewOrganizationID(apiParam.OrganizationID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})
		return
	}

	subject := rsuserdomain.NewRBACUser(apiParam.Subject)
	action := rsuserdomain.NewRBACAction(apiParam.Action)
	object := rsuserdomain.NewRBACObject(apiParam.Object)
	effect := rsuserdomain.NewRBACEffect(apiParam.Effect)

	if err := h.rbacUsecase.AddPolicyToUser(ctx, organizationID, subject, action, object, effect); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": http.StatusText(http.StatusInternalServerError)})
		return
	}
}

func (h *RBACHandler) AddPolicyToGroup(c *gin.Context) {

}

func NewInitRBACRouterFunc(rbacUsecase RBACUsecase) libcontroller.InitRouterGroupFunc {
	return func(parentRouterGroup gin.IRouter, middleware ...gin.HandlerFunc) {
		rbac := parentRouterGroup.Group("rbac")
		for _, m := range middleware {
			rbac.Use(m)
		}

		rbacHandler := NewRBACHandler(rbacUsecase)
		rbac.PUT("policy/user", rbacHandler.AddPolicyToUser)
		rbac.PUT("policy/group", rbacHandler.AddPolicyToGroup)
	}
}
