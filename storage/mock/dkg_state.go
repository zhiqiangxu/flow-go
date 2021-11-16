// Code generated by mockery v1.0.0. DO NOT EDIT.

package mock

import mock "github.com/stretchr/testify/mock"

// DKGState is an autogenerated mock type for the DKGState type
type DKGState struct {
	mock.Mock
}

// GetDKGStarted provides a mock function with given fields: epochCounter
func (_m *DKGState) GetDKGStarted(epochCounter uint64) (bool, error) {
	ret := _m.Called(epochCounter)

	var r0 bool
	if rf, ok := ret.Get(0).(func(uint64) bool); ok {
		r0 = rf(epochCounter)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(epochCounter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetDKGStarted provides a mock function with given fields: epochCounter
func (_m *DKGState) SetDKGStarted(epochCounter uint64) error {
	ret := _m.Called(epochCounter)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(epochCounter)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
