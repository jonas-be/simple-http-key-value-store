// Code generated by mockery v2.22.1. DO NOT EDIT.

package endpoint

import mock "github.com/stretchr/testify/mock"

// MockDatabase is an autogenerated mock type for the Database type
type MockDatabase struct {
	mock.Mock
}

type MockDatabase_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDatabase) EXPECT() *MockDatabase_Expecter {
	return &MockDatabase_Expecter{mock: &_m.Mock}
}

// Contains provides a mock function with given fields: key
func (_m *MockDatabase) Contains(key string) bool {
	ret := _m.Called(key)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockDatabase_Contains_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Contains'
type MockDatabase_Contains_Call struct {
	*mock.Call
}

// Contains is a helper method to define mock.On call
//   - key string
func (_e *MockDatabase_Expecter) Contains(key interface{}) *MockDatabase_Contains_Call {
	return &MockDatabase_Contains_Call{Call: _e.mock.On("Contains", key)}
}

func (_c *MockDatabase_Contains_Call) Run(run func(key string)) *MockDatabase_Contains_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockDatabase_Contains_Call) Return(_a0 bool) *MockDatabase_Contains_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDatabase_Contains_Call) RunAndReturn(run func(string) bool) *MockDatabase_Contains_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: key
func (_m *MockDatabase) Delete(key string) {
	_m.Called(key)
}

// MockDatabase_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockDatabase_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - key string
func (_e *MockDatabase_Expecter) Delete(key interface{}) *MockDatabase_Delete_Call {
	return &MockDatabase_Delete_Call{Call: _e.mock.On("Delete", key)}
}

func (_c *MockDatabase_Delete_Call) Run(run func(key string)) *MockDatabase_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockDatabase_Delete_Call) Return() *MockDatabase_Delete_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockDatabase_Delete_Call) RunAndReturn(run func(string)) *MockDatabase_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: key
func (_m *MockDatabase) Get(key string) string {
	ret := _m.Called(key)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockDatabase_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockDatabase_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - key string
func (_e *MockDatabase_Expecter) Get(key interface{}) *MockDatabase_Get_Call {
	return &MockDatabase_Get_Call{Call: _e.mock.On("Get", key)}
}

func (_c *MockDatabase_Get_Call) Run(run func(key string)) *MockDatabase_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockDatabase_Get_Call) Return(_a0 string) *MockDatabase_Get_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDatabase_Get_Call) RunAndReturn(run func(string) string) *MockDatabase_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Set provides a mock function with given fields: key, value
func (_m *MockDatabase) Set(key string, value string) error {
	ret := _m.Called(key, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(key, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDatabase_Set_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Set'
type MockDatabase_Set_Call struct {
	*mock.Call
}

// Set is a helper method to define mock.On call
//   - key string
//   - value string
func (_e *MockDatabase_Expecter) Set(key interface{}, value interface{}) *MockDatabase_Set_Call {
	return &MockDatabase_Set_Call{Call: _e.mock.On("Set", key, value)}
}

func (_c *MockDatabase_Set_Call) Run(run func(key string, value string)) *MockDatabase_Set_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockDatabase_Set_Call) Return(_a0 error) *MockDatabase_Set_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDatabase_Set_Call) RunAndReturn(run func(string, string) error) *MockDatabase_Set_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockDatabase interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockDatabase creates a new instance of MockDatabase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockDatabase(t mockConstructorTestingTNewMockDatabase) *MockDatabase {
	mock := &MockDatabase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}