// Code generated by mockery v2.30.16. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Assigner is an autogenerated mock type for the Assigner type
type Assigner struct {
	mock.Mock
}

type Assigner_Expecter struct {
	mock *mock.Mock
}

func (_m *Assigner) EXPECT() *Assigner_Expecter {
	return &Assigner_Expecter{mock: &_m.Mock}
}

// Assign provides a mock function with given fields: instanceID, zone, filter, orderBy
func (_m *Assigner) Assign(instanceID string, zone string, filter []string, orderBy string) error {
	ret := _m.Called(instanceID, zone, filter, orderBy)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, []string, string) error); ok {
		r0 = rf(instanceID, zone, filter, orderBy)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Assigner_Assign_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Assign'
type Assigner_Assign_Call struct {
	*mock.Call
}

// Assign is a helper method to define mock.On call
//   - instanceID string
//   - zone string
//   - filter []string
//   - orderBy string
func (_e *Assigner_Expecter) Assign(instanceID interface{}, zone interface{}, filter interface{}, orderBy interface{}) *Assigner_Assign_Call {
	return &Assigner_Assign_Call{Call: _e.mock.On("Assign", instanceID, zone, filter, orderBy)}
}

func (_c *Assigner_Assign_Call) Run(run func(instanceID string, zone string, filter []string, orderBy string)) *Assigner_Assign_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].([]string), args[3].(string))
	})
	return _c
}

func (_c *Assigner_Assign_Call) Return(_a0 error) *Assigner_Assign_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Assigner_Assign_Call) RunAndReturn(run func(string, string, []string, string) error) *Assigner_Assign_Call {
	_c.Call.Return(run)
	return _c
}

// NewAssigner creates a new instance of Assigner. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAssigner(t interface {
	mock.TestingT
	Cleanup(func())
}) *Assigner {
	mock := &Assigner{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
