// Code generated by mockery v2.49.0. DO NOT EDIT.

package mocks

import (
	service "github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/service"
	mock "github.com/stretchr/testify/mock"
)

// TatoebaSentencePair is an autogenerated mock type for the TatoebaSentencePair type
type TatoebaSentencePair struct {
	mock.Mock
}

type TatoebaSentencePair_Expecter struct {
	mock *mock.Mock
}

func (_m *TatoebaSentencePair) EXPECT() *TatoebaSentencePair_Expecter {
	return &TatoebaSentencePair_Expecter{mock: &_m.Mock}
}

// GetDst provides a mock function with given fields:
func (_m *TatoebaSentencePair) GetDst() service.TatoebaSentence {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetDst")
	}

	var r0 service.TatoebaSentence
	if rf, ok := ret.Get(0).(func() service.TatoebaSentence); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(service.TatoebaSentence)
		}
	}

	return r0
}

// TatoebaSentencePair_GetDst_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDst'
type TatoebaSentencePair_GetDst_Call struct {
	*mock.Call
}

// GetDst is a helper method to define mock.On call
func (_e *TatoebaSentencePair_Expecter) GetDst() *TatoebaSentencePair_GetDst_Call {
	return &TatoebaSentencePair_GetDst_Call{Call: _e.mock.On("GetDst")}
}

func (_c *TatoebaSentencePair_GetDst_Call) Run(run func()) *TatoebaSentencePair_GetDst_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *TatoebaSentencePair_GetDst_Call) Return(_a0 service.TatoebaSentence) *TatoebaSentencePair_GetDst_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TatoebaSentencePair_GetDst_Call) RunAndReturn(run func() service.TatoebaSentence) *TatoebaSentencePair_GetDst_Call {
	_c.Call.Return(run)
	return _c
}

// GetSrc provides a mock function with given fields:
func (_m *TatoebaSentencePair) GetSrc() service.TatoebaSentence {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetSrc")
	}

	var r0 service.TatoebaSentence
	if rf, ok := ret.Get(0).(func() service.TatoebaSentence); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(service.TatoebaSentence)
		}
	}

	return r0
}

// TatoebaSentencePair_GetSrc_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSrc'
type TatoebaSentencePair_GetSrc_Call struct {
	*mock.Call
}

// GetSrc is a helper method to define mock.On call
func (_e *TatoebaSentencePair_Expecter) GetSrc() *TatoebaSentencePair_GetSrc_Call {
	return &TatoebaSentencePair_GetSrc_Call{Call: _e.mock.On("GetSrc")}
}

func (_c *TatoebaSentencePair_GetSrc_Call) Run(run func()) *TatoebaSentencePair_GetSrc_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *TatoebaSentencePair_GetSrc_Call) Return(_a0 service.TatoebaSentence) *TatoebaSentencePair_GetSrc_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TatoebaSentencePair_GetSrc_Call) RunAndReturn(run func() service.TatoebaSentence) *TatoebaSentencePair_GetSrc_Call {
	_c.Call.Return(run)
	return _c
}

// NewTatoebaSentencePair creates a new instance of TatoebaSentencePair. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTatoebaSentencePair(t interface {
	mock.TestingT
	Cleanup(func())
}) *TatoebaSentencePair {
	mock := &TatoebaSentencePair{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}