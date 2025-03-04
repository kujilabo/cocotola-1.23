package domain

import (
	libdomain "github.com/kujilabo/cocotola-1.23/redstart/lib/domain"
	liberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
)

type AppUserID struct {
	Value int `validate:"required,gte=0"`
}

func NewAppUserID(value int) (*AppUserID, error) {
	return &AppUserID{
		Value: value,
	}, nil
}

func (v *AppUserID) Int() int {
	return v.Value
}
func (v *AppUserID) IsAppUserID() bool {
	return true
}

type AppUserModel struct {
	*libdomain.BaseModel
	AppUserID      *AppUserID
	OrganizationID *OrganizationID
	LoginID        string `validate:"required"`
	Username       string `validate:"required"`
	UserGroups     []*UserGroupModel
}

func NewAppUserModel(baseModel *libdomain.BaseModel, appUserID *AppUserID, organizationID *OrganizationID, loginID, username string, userGroups []*UserGroupModel) (*AppUserModel, error) {
	m := &AppUserModel{
		BaseModel:      baseModel,
		AppUserID:      appUserID,
		OrganizationID: organizationID,
		LoginID:        loginID,
		Username:       username,
		UserGroups:     userGroups,
	}

	if err := libdomain.Validator.Struct(m); err != nil {
		return nil, liberrors.Errorf("libdomain.Validator.Struct. err: %w", err)
	}

	return m, nil
}
