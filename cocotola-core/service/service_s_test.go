//go:build small

package service_test

import (
	"testing"

	"github.com/kujilabo/cocotola-1.23/cocotola-core/service"
)

func TestA(t *testing.T) {
	service.A()
}

func TestB(t *testing.T) {
	service.B()
}

func TestC(t *testing.T) {
	service.C()
}
