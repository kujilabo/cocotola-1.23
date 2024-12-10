//go:generate mockery --output mock --name TatoebaLinkAddParameter
//go:generate mockery --output mock --name TatoebaLinkRepository
package service

import (
	"context"
	"errors"

	libdomain "github.com/kujilabo/cocotola-1.23/redstart/lib/domain"
)

var ErrTatoebaLinkAlreadyExists = errors.New("tatoebaLink already exists")
var ErrTatoebaLinkSourceNotFound = errors.New("tatoebaLink source not found")

type TatoebaLinkAddParameter interface {
	GetFrom() int
	GetTo() int
}

type tatoebaLinkAddParameter struct {
	From int `validate:"required"`
	To   int `validate:"required"`
}

func NewTatoebaLinkAddParameter(from, to int) (TatoebaLinkAddParameter, error) {
	m := &tatoebaLinkAddParameter{
		From: from,
		To:   to,
	}
	return m, libdomain.Validator.Struct(m)
}

func (p *tatoebaLinkAddParameter) GetFrom() int {
	return p.From
}

func (p *tatoebaLinkAddParameter) GetTo() int {
	return p.To
}

type TatoebaLinkRepository interface {
	Add(ctx context.Context, param TatoebaLinkAddParameter) error
}
