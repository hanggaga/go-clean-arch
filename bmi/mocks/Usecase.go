// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// CalculateBMI provides a mock function with given fields: height, weight
func (_m *Usecase) CalculateBMI(height float64, weight float64) (float64, error) {
	ret := _m.Called(height, weight)

	if len(ret) == 0 {
		panic("no return value specified for CalculateBMI")
	}

	var r0 float64
	var r1 error
	if rf, ok := ret.Get(0).(func(float64, float64) (float64, error)); ok {
		return rf(height, weight)
	}
	if rf, ok := ret.Get(0).(func(float64, float64) float64); ok {
		r0 = rf(height, weight)
	} else {
		r0 = ret.Get(0).(float64)
	}

	if rf, ok := ret.Get(1).(func(float64, float64) error); ok {
		r1 = rf(height, weight)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUsecase creates a new instance of Usecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *Usecase {
	mock := &Usecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
