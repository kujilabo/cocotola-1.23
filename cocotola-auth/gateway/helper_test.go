package gateway_test

import (
	"testing"

	rsuserdomain "github.com/kujilabo/cocotola-1.23/redstart/user/domain"
	"github.com/stretchr/testify/require"
)

type organization struct {
	organizationID *rsuserdomain.OrganizationID
	name           string
}

func (m *organization) OrganizationID() *rsuserdomain.OrganizationID {
	return m.organizationID
}
func (m *organization) Name() string {
	return m.name
}

type appUser struct {
	appUserID      *rsuserdomain.AppUserID
	organizationID *rsuserdomain.OrganizationID
	loginID        string
	username       string
}

func (m *appUser) AppUserID() *rsuserdomain.AppUserID {
	return m.appUserID
}
func (m *appUser) OrganizationID() *rsuserdomain.OrganizationID {
	return m.organizationID
}
func (m *appUser) Username() string {
	return m.username
}
func (m *appUser) LoginID() string {
	return m.loginID
}

func organizationID(t *testing.T, organizationID int) *rsuserdomain.OrganizationID {
	t.Helper()
	id, err := rsuserdomain.NewOrganizationID(organizationID)
	require.NoError(t, err)
	return id
}

func appUserID(t *testing.T, appUserID int) *rsuserdomain.AppUserID {
	t.Helper()
	id, err := rsuserdomain.NewAppUserID(appUserID)
	require.NoError(t, err)
	return id
}
