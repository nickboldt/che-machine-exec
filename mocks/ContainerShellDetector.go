// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ContainerShellDetector is an autogenerated mock type for the ContainerShellDetector type
type ContainerShellDetector struct {
	mock.Mock
}

// DetectShell provides a mock function with given fields: containerInfo
func (_m *ContainerShellDetector) DetectShell(containerInfo map[string]string) (string, error) {
	ret := _m.Called(containerInfo)

	var r0 string
	if rf, ok := ret.Get(0).(func(map[string]string) string); ok {
		r0 = rf(containerInfo)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(map[string]string) error); ok {
		r1 = rf(containerInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
