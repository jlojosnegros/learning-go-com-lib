// Code generated by mockery v2.44.1. DO NOT EDIT.

package programming

import mock "github.com/stretchr/testify/mock"

// MockInterface is an autogenerated mock type for the Interface type
type MockInterface struct {
	mock.Mock
}

// NewUuid provides a mock function with given fields: withoutHyphen
func (_m *MockInterface) NewUuid(withoutHyphen bool) string {
	ret := _m.Called(withoutHyphen)

	if len(ret) == 0 {
		panic("no return value specified for NewUuid")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(bool) string); ok {
		r0 = rf(withoutHyphen)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewMockInterface creates a new instance of MockInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockInterface {
	mock := &MockInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
