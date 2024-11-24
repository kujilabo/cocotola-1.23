package service

import (
	libdomain "github.com/kujilabo/cocotola-1.23/redstart/lib/domain"
	liberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"

	"github.com/kujilabo/cocotola-1.23/redstart/user/domain"
)

type Organization struct {
	*domain.OrganizationModel
}

func NewOrganization(organizationModel *domain.OrganizationModel) (*Organization, error) {
	m := &Organization{
		organizationModel,
	}

	if err := libdomain.Validator.Struct(m); err != nil {
		return nil, liberrors.Errorf("libdomain.Validator.Struct. err: %w", err)
	}

	return m, nil
}

func (m *Organization) OrganizationID() *domain.OrganizationID {
	return m.OrganizationModel.OrganizationID
}
func (m *Organization) Name() string {
	return m.OrganizationModel.Name
}
