// Code generated by mockery v2.49.0. DO NOT EDIT.

package mocks

import (
	context "context"

	api "github.com/kujilabo/cocotola-1.23/lib/api"

	domain "github.com/kujilabo/cocotola-1.23/cocotola-core/domain"

	mock "github.com/stretchr/testify/mock"

	service "github.com/kujilabo/cocotola-1.23/cocotola-core/service"
)

// WorkbookQueryUsecase is an autogenerated mock type for the WorkbookQueryUsecase type
type WorkbookQueryUsecase struct {
	mock.Mock
}

type WorkbookQueryUsecase_Expecter struct {
	mock *mock.Mock
}

func (_m *WorkbookQueryUsecase) EXPECT() *WorkbookQueryUsecase_Expecter {
	return &WorkbookQueryUsecase_Expecter{mock: &_m.Mock}
}

// FindWorkbooks provides a mock function with given fields: ctx, operator, param
func (_m *WorkbookQueryUsecase) FindWorkbooks(ctx context.Context, operator service.OperatorInterface, param *api.WorkbookFindParameter) (*api.WorkbookFindResult, error) {
	ret := _m.Called(ctx, operator, param)

	if len(ret) == 0 {
		panic("no return value specified for FindWorkbooks")
	}

	var r0 *api.WorkbookFindResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, service.OperatorInterface, *api.WorkbookFindParameter) (*api.WorkbookFindResult, error)); ok {
		return rf(ctx, operator, param)
	}
	if rf, ok := ret.Get(0).(func(context.Context, service.OperatorInterface, *api.WorkbookFindParameter) *api.WorkbookFindResult); ok {
		r0 = rf(ctx, operator, param)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.WorkbookFindResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, service.OperatorInterface, *api.WorkbookFindParameter) error); ok {
		r1 = rf(ctx, operator, param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WorkbookQueryUsecase_FindWorkbooks_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindWorkbooks'
type WorkbookQueryUsecase_FindWorkbooks_Call struct {
	*mock.Call
}

// FindWorkbooks is a helper method to define mock.On call
//   - ctx context.Context
//   - operator service.OperatorInterface
//   - param *api.WorkbookFindParameter
func (_e *WorkbookQueryUsecase_Expecter) FindWorkbooks(ctx interface{}, operator interface{}, param interface{}) *WorkbookQueryUsecase_FindWorkbooks_Call {
	return &WorkbookQueryUsecase_FindWorkbooks_Call{Call: _e.mock.On("FindWorkbooks", ctx, operator, param)}
}

func (_c *WorkbookQueryUsecase_FindWorkbooks_Call) Run(run func(ctx context.Context, operator service.OperatorInterface, param *api.WorkbookFindParameter)) *WorkbookQueryUsecase_FindWorkbooks_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(service.OperatorInterface), args[2].(*api.WorkbookFindParameter))
	})
	return _c
}

func (_c *WorkbookQueryUsecase_FindWorkbooks_Call) Return(_a0 *api.WorkbookFindResult, _a1 error) *WorkbookQueryUsecase_FindWorkbooks_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *WorkbookQueryUsecase_FindWorkbooks_Call) RunAndReturn(run func(context.Context, service.OperatorInterface, *api.WorkbookFindParameter) (*api.WorkbookFindResult, error)) *WorkbookQueryUsecase_FindWorkbooks_Call {
	_c.Call.Return(run)
	return _c
}

// RetrieveWorkbookByID provides a mock function with given fields: ctx, operator, workbookID
func (_m *WorkbookQueryUsecase) RetrieveWorkbookByID(ctx context.Context, operator service.OperatorInterface, workbookID *domain.WorkbookID) (*api.WorkbookRetrieveResult, error) {
	ret := _m.Called(ctx, operator, workbookID)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveWorkbookByID")
	}

	var r0 *api.WorkbookRetrieveResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, service.OperatorInterface, *domain.WorkbookID) (*api.WorkbookRetrieveResult, error)); ok {
		return rf(ctx, operator, workbookID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, service.OperatorInterface, *domain.WorkbookID) *api.WorkbookRetrieveResult); ok {
		r0 = rf(ctx, operator, workbookID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.WorkbookRetrieveResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, service.OperatorInterface, *domain.WorkbookID) error); ok {
		r1 = rf(ctx, operator, workbookID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WorkbookQueryUsecase_RetrieveWorkbookByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RetrieveWorkbookByID'
type WorkbookQueryUsecase_RetrieveWorkbookByID_Call struct {
	*mock.Call
}

// RetrieveWorkbookByID is a helper method to define mock.On call
//   - ctx context.Context
//   - operator service.OperatorInterface
//   - workbookID *domain.WorkbookID
func (_e *WorkbookQueryUsecase_Expecter) RetrieveWorkbookByID(ctx interface{}, operator interface{}, workbookID interface{}) *WorkbookQueryUsecase_RetrieveWorkbookByID_Call {
	return &WorkbookQueryUsecase_RetrieveWorkbookByID_Call{Call: _e.mock.On("RetrieveWorkbookByID", ctx, operator, workbookID)}
}

func (_c *WorkbookQueryUsecase_RetrieveWorkbookByID_Call) Run(run func(ctx context.Context, operator service.OperatorInterface, workbookID *domain.WorkbookID)) *WorkbookQueryUsecase_RetrieveWorkbookByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(service.OperatorInterface), args[2].(*domain.WorkbookID))
	})
	return _c
}

func (_c *WorkbookQueryUsecase_RetrieveWorkbookByID_Call) Return(_a0 *api.WorkbookRetrieveResult, _a1 error) *WorkbookQueryUsecase_RetrieveWorkbookByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *WorkbookQueryUsecase_RetrieveWorkbookByID_Call) RunAndReturn(run func(context.Context, service.OperatorInterface, *domain.WorkbookID) (*api.WorkbookRetrieveResult, error)) *WorkbookQueryUsecase_RetrieveWorkbookByID_Call {
	_c.Call.Return(run)
	return _c
}

// NewWorkbookQueryUsecase creates a new instance of WorkbookQueryUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewWorkbookQueryUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *WorkbookQueryUsecase {
	mock := &WorkbookQueryUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}