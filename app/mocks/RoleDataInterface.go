// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	roles "event_ticket/features/roles"

	mock "github.com/stretchr/testify/mock"
)

// RoleDataInterface is an autogenerated mock type for the RoleDataInterface type
type RoleDataInterface struct {
	mock.Mock
}

// CreateRole provides a mock function with given fields: data
func (_m *RoleDataInterface) CreateRole(data roles.RoleCore) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(roles.RoleCore) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteRole provides a mock function with given fields: id
func (_m *RoleDataInterface) DeleteRole(id uint64) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReadAllRole provides a mock function with given fields:
func (_m *RoleDataInterface) ReadAllRole() ([]roles.RoleCore, error) {
	ret := _m.Called()

	var r0 []roles.RoleCore
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]roles.RoleCore, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []roles.RoleCore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]roles.RoleCore)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadSpecificRole provides a mock function with given fields: id
func (_m *RoleDataInterface) ReadSpecificRole(id string) (roles.RoleCore, error) {
	ret := _m.Called(id)

	var r0 roles.RoleCore
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (roles.RoleCore, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) roles.RoleCore); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(roles.RoleCore)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRoleDataInterface creates a new instance of RoleDataInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRoleDataInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *RoleDataInterface {
	mock := &RoleDataInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
