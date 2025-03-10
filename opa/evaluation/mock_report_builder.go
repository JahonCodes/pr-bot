// Code generated by mockery v2.49.0. DO NOT EDIT.

package evaluation

import (
	input "github.com/marqeta/pr-bot/opa/input"
	mock "github.com/stretchr/testify/mock"
)

// MockReportBuilder is an autogenerated mock type for the ReportBuilder type
type MockReportBuilder struct {
	mock.Mock
}

type MockReportBuilder_Expecter struct {
	mock *mock.Mock
}

func (_m *MockReportBuilder) EXPECT() *MockReportBuilder_Expecter {
	return &MockReportBuilder_Expecter{mock: &_m.Mock}
}

// AddModuleResult provides a mock function with given fields: module, result
func (_m *MockReportBuilder) AddModuleResult(module string, result Result) {
	_m.Called(module, result)
}

// MockReportBuilder_AddModuleResult_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddModuleResult'
type MockReportBuilder_AddModuleResult_Call struct {
	*mock.Call
}

// AddModuleResult is a helper method to define mock.On call
//   - module string
//   - result Result
func (_e *MockReportBuilder_Expecter) AddModuleResult(module interface{}, result interface{}) *MockReportBuilder_AddModuleResult_Call {
	return &MockReportBuilder_AddModuleResult_Call{Call: _e.mock.On("AddModuleResult", module, result)}
}

func (_c *MockReportBuilder_AddModuleResult_Call) Run(run func(module string, result Result)) *MockReportBuilder_AddModuleResult_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(Result))
	})
	return _c
}

func (_c *MockReportBuilder_AddModuleResult_Call) Return() *MockReportBuilder_AddModuleResult_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockReportBuilder_AddModuleResult_Call) RunAndReturn(run func(string, Result)) *MockReportBuilder_AddModuleResult_Call {
	_c.Call.Return(run)
	return _c
}

// GetReport provides a mock function with given fields:
func (_m *MockReportBuilder) GetReport() Report {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetReport")
	}

	var r0 Report
	if rf, ok := ret.Get(0).(func() Report); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(Report)
	}

	return r0
}

// MockReportBuilder_GetReport_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetReport'
type MockReportBuilder_GetReport_Call struct {
	*mock.Call
}

// GetReport is a helper method to define mock.On call
func (_e *MockReportBuilder_Expecter) GetReport() *MockReportBuilder_GetReport_Call {
	return &MockReportBuilder_GetReport_Call{Call: _e.mock.On("GetReport")}
}

func (_c *MockReportBuilder_GetReport_Call) Run(run func()) *MockReportBuilder_GetReport_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockReportBuilder_GetReport_Call) Return(_a0 Report) *MockReportBuilder_GetReport_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockReportBuilder_GetReport_Call) RunAndReturn(run func() Report) *MockReportBuilder_GetReport_Call {
	_c.Call.Return(run)
	return _c
}

// SetInput provides a mock function with given fields: _a0
func (_m *MockReportBuilder) SetInput(_a0 *input.Model) {
	_m.Called(_a0)
}

// MockReportBuilder_SetInput_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetInput'
type MockReportBuilder_SetInput_Call struct {
	*mock.Call
}

// SetInput is a helper method to define mock.On call
//   - _a0 *input.Model
func (_e *MockReportBuilder_Expecter) SetInput(_a0 interface{}) *MockReportBuilder_SetInput_Call {
	return &MockReportBuilder_SetInput_Call{Call: _e.mock.On("SetInput", _a0)}
}

func (_c *MockReportBuilder_SetInput_Call) Run(run func(_a0 *input.Model)) *MockReportBuilder_SetInput_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*input.Model))
	})
	return _c
}

func (_c *MockReportBuilder_SetInput_Call) Return() *MockReportBuilder_SetInput_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockReportBuilder_SetInput_Call) RunAndReturn(run func(*input.Model)) *MockReportBuilder_SetInput_Call {
	_c.Call.Return(run)
	return _c
}

// SetOutcome provides a mock function with given fields: result
func (_m *MockReportBuilder) SetOutcome(result Result) {
	_m.Called(result)
}

// MockReportBuilder_SetOutcome_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetOutcome'
type MockReportBuilder_SetOutcome_Call struct {
	*mock.Call
}

// SetOutcome is a helper method to define mock.On call
//   - result Result
func (_e *MockReportBuilder_Expecter) SetOutcome(result interface{}) *MockReportBuilder_SetOutcome_Call {
	return &MockReportBuilder_SetOutcome_Call{Call: _e.mock.On("SetOutcome", result)}
}

func (_c *MockReportBuilder_SetOutcome_Call) Run(run func(result Result)) *MockReportBuilder_SetOutcome_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(Result))
	})
	return _c
}

func (_c *MockReportBuilder_SetOutcome_Call) Return() *MockReportBuilder_SetOutcome_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockReportBuilder_SetOutcome_Call) RunAndReturn(run func(Result)) *MockReportBuilder_SetOutcome_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockReportBuilder creates a new instance of MockReportBuilder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockReportBuilder(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockReportBuilder {
	mock := &MockReportBuilder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
