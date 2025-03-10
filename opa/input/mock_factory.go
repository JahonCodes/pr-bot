// Code generated by mockery v2.49.0. DO NOT EDIT.

package input

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockFactory is an autogenerated mock type for the Factory type
type MockFactory struct {
	mock.Mock
}

type MockFactory_Expecter struct {
	mock *mock.Mock
}

func (_m *MockFactory) EXPECT() *MockFactory_Expecter {
	return &MockFactory_Expecter{mock: &_m.Mock}
}

// CreateModel provides a mock function with given fields: ctx, ghe
func (_m *MockFactory) CreateModel(ctx context.Context, ghe GHE) (*Model, error) {
	ret := _m.Called(ctx, ghe)

	if len(ret) == 0 {
		panic("no return value specified for CreateModel")
	}

	var r0 *Model
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, GHE) (*Model, error)); ok {
		return rf(ctx, ghe)
	}
	if rf, ok := ret.Get(0).(func(context.Context, GHE) *Model); ok {
		r0 = rf(ctx, ghe)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Model)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, GHE) error); ok {
		r1 = rf(ctx, ghe)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFactory_CreateModel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateModel'
type MockFactory_CreateModel_Call struct {
	*mock.Call
}

// CreateModel is a helper method to define mock.On call
//   - ctx context.Context
//   - ghe GHE
func (_e *MockFactory_Expecter) CreateModel(ctx interface{}, ghe interface{}) *MockFactory_CreateModel_Call {
	return &MockFactory_CreateModel_Call{Call: _e.mock.On("CreateModel", ctx, ghe)}
}

func (_c *MockFactory_CreateModel_Call) Run(run func(ctx context.Context, ghe GHE)) *MockFactory_CreateModel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(GHE))
	})
	return _c
}

func (_c *MockFactory_CreateModel_Call) Return(_a0 *Model, _a1 error) *MockFactory_CreateModel_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockFactory_CreateModel_Call) RunAndReturn(run func(context.Context, GHE) (*Model, error)) *MockFactory_CreateModel_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockFactory creates a new instance of MockFactory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockFactory(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockFactory {
	mock := &MockFactory{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
