package service

import (
	"context"

	"github.com/casbin/casbin/v2"

	"github.com/kujilabo/cocotola-1.23/redstart/user/domain"
)

type RBACRepository interface {
	Init() error
	GetEnforcer() casbin.IEnforcer
	// who can do what actions on which resources
	AddPolicy(ctx context.Context, domain domain.RBACDomain, subject domain.RBACSubject, action domain.RBACAction, object domain.RBACObject, effect domain.RBACEffect) error

	// add user to group
	AddSubjectGroupingPolicy(ctx context.Context, domain domain.RBACDomain, subject domain.RBACUser, object domain.RBACRole) error

	// add child object to parent object
	AddObjectGroupingPolicy(ctx context.Context, domain domain.RBACDomain, child domain.RBACObject, parent domain.RBACObject) error

	RemovePolicy(ctx context.Context, domain domain.RBACDomain, subject domain.RBACSubject, action domain.RBACAction, object domain.RBACObject, effect domain.RBACEffect) error
	// RemoveSubjectPolicy(domain domain.RBACDomain, subject domain.RBACSubject) error

	RemoveSubjectGroupingPolicy(ctx context.Context, domain domain.RBACDomain, subject domain.RBACUser, object domain.RBACRole) error
	RemoveObjectGroupingPolicy(ctx context.Context, domain domain.RBACDomain, child domain.RBACObject, parent domain.RBACObject) error

	NewEnforcerWithGroupsAndUsers(ctx context.Context, roles []domain.RBACRole, users []domain.RBACUser) (casbin.IEnforcer, error)
}
