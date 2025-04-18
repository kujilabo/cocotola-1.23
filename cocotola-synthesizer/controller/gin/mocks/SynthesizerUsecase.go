// Code generated by mockery v2.52.2. DO NOT EDIT.

package mocks

import (
	context "context"

	cocotola_synthesizerdomain "github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/domain"

	domain "github.com/kujilabo/cocotola-1.23/lib/domain"

	mock "github.com/stretchr/testify/mock"
)

// SynthesizerUsecase is an autogenerated mock type for the SynthesizerUsecase type
type SynthesizerUsecase struct {
	mock.Mock
}

type SynthesizerUsecase_Expecter struct {
	mock *mock.Mock
}

func (_m *SynthesizerUsecase) EXPECT() *SynthesizerUsecase_Expecter {
	return &SynthesizerUsecase_Expecter{mock: &_m.Mock}
}

// Synthesize provides a mock function with given fields: ctx, lang5, void, text
func (_m *SynthesizerUsecase) Synthesize(ctx context.Context, lang5 *domain.Lang5, void string, text string) (*cocotola_synthesizerdomain.AudioModel, error) {
	ret := _m.Called(ctx, lang5, void, text)

	if len(ret) == 0 {
		panic("no return value specified for Synthesize")
	}

	var r0 *cocotola_synthesizerdomain.AudioModel
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Lang5, string, string) (*cocotola_synthesizerdomain.AudioModel, error)); ok {
		return rf(ctx, lang5, void, text)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Lang5, string, string) *cocotola_synthesizerdomain.AudioModel); ok {
		r0 = rf(ctx, lang5, void, text)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cocotola_synthesizerdomain.AudioModel)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *domain.Lang5, string, string) error); ok {
		r1 = rf(ctx, lang5, void, text)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SynthesizerUsecase_Synthesize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Synthesize'
type SynthesizerUsecase_Synthesize_Call struct {
	*mock.Call
}

// Synthesize is a helper method to define mock.On call
//   - ctx context.Context
//   - lang5 *domain.Lang5
//   - void string
//   - text string
func (_e *SynthesizerUsecase_Expecter) Synthesize(ctx interface{}, lang5 interface{}, void interface{}, text interface{}) *SynthesizerUsecase_Synthesize_Call {
	return &SynthesizerUsecase_Synthesize_Call{Call: _e.mock.On("Synthesize", ctx, lang5, void, text)}
}

func (_c *SynthesizerUsecase_Synthesize_Call) Run(run func(ctx context.Context, lang5 *domain.Lang5, void string, text string)) *SynthesizerUsecase_Synthesize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.Lang5), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *SynthesizerUsecase_Synthesize_Call) Return(_a0 *cocotola_synthesizerdomain.AudioModel, _a1 error) *SynthesizerUsecase_Synthesize_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SynthesizerUsecase_Synthesize_Call) RunAndReturn(run func(context.Context, *domain.Lang5, string, string) (*cocotola_synthesizerdomain.AudioModel, error)) *SynthesizerUsecase_Synthesize_Call {
	_c.Call.Return(run)
	return _c
}

// NewSynthesizerUsecase creates a new instance of SynthesizerUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSynthesizerUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *SynthesizerUsecase {
	mock := &SynthesizerUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
