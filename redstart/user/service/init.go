package service

import (
	libdomain "github.com/kujilabo/cocotola-1.23/redstart/lib/domain"
)

const (
	UserServiceContextKey libdomain.ContextKey = "user_service"

	SystemAdminLoginID = "__system_admin"
	SystemOwnerLoginID = "__system_owner"

	SystemOwnerGroupKey = "__system_owner"
	OwnerGroupKey       = "__owner"

	SystemOwnerGroupName = "System Owner"
	OwnerGroupName       = "Owner"
)
