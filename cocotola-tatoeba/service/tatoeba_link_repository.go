package service

import (
	"context"
	"errors"

	libdomain "github.com/kujilabo/cocotola-1.23/redstart/lib/domain"
)

var ErrTatoebaLinkAlreadyExists = errors.New("tatoebaLink already exists")
var ErrTatoebaLinkSourceNotFound = errors.New("tatoebaLink source not found")

type TatoebaLinkAddParameter interface {
	GetSrc() int
	GetDst() int
}

type tatoebaLinkAddParameter struct {
	Src int `validate:"required"`
	Dst int `validate:"required"`
}

func NewTatoebaLinkAddParameter(src, dst int) (TatoebaLinkAddParameter, error) {
	m := &tatoebaLinkAddParameter{
		Src: src,
		Dst: dst,
	}
	return m, libdomain.Validator.Struct(m)
}

func (p *tatoebaLinkAddParameter) GetSrc() int {
	return p.Src
}

func (p *tatoebaLinkAddParameter) GetDst() int {
	return p.Dst
}

type TatoebaLinkRepository interface {
	Add(ctx context.Context, param TatoebaLinkAddParameter) error
}
