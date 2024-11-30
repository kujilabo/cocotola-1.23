package domain_test

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	libdomain "github.com/kujilabo/cocotola-1.23/redstart/lib/domain"

	"github.com/kujilabo/cocotola-1.23/redstart/user/domain"
)

func TestNewSystemOwner(t *testing.T) {
	t.Parallel()
	model, err := libdomain.NewBaseModel(1, time.Now(), time.Now(), 1, 1)
	require.NoError(t, err)
	appUserID, err := domain.NewAppUserID(1)
	require.NoError(t, err)
	organizationID, err := domain.NewOrganizationID(1)
	require.NoError(t, err)
	appUser, err := domain.NewAppUserModel(model, appUserID, organizationID, "LOGIN_ID", "USERNAME", nil)
	assert.NoError(t, err)
	ower, err := domain.NewOwnerModel(appUser)
	assert.NoError(t, err)
	systemOwner, err := domain.NewSystemOwnerModel(ower)
	assert.NoError(t, err)
	log.Println(systemOwner)
}
