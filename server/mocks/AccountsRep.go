// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	objects "api/recipes/objects"

	mock "github.com/stretchr/testify/mock"
)

// AccountsRep is an autogenerated mock type for the AccountsRep type
type AccountsRep struct {
	mock.Mock
}

// Create provides a mock function with given fields: obj
func (_m *AccountsRep) Create(obj *objects.Account) error {
	ret := _m.Called(obj)

	var r0 error
	if rf, ok := ret.Get(0).(func(*objects.Account) error); ok {
		r0 = rf(obj)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateList provides a mock function with given fields: obj
func (_m *AccountsRep) CreateList(obj []objects.Account) error {
	ret := _m.Called(obj)

	var r0 error
	if rf, ok := ret.Get(0).(func([]objects.Account) error); ok {
		r0 = rf(obj)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: login
func (_m *AccountsRep) Find(login string) (*objects.Account, error) {
	ret := _m.Called(login)

	var r0 *objects.Account
	if rf, ok := ret.Get(0).(func(string) *objects.Account); ok {
		r0 = rf(login)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*objects.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(login)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindLikedRecipe provides a mock function with given fields: id_rcp
func (_m *AccountsRep) FindLikedRecipe(id_rcp int) ([]objects.Account, error) {
	ret := _m.Called(id_rcp)

	var r0 []objects.Account
	if rf, ok := ret.Get(0).(func(int) []objects.Account); ok {
		r0 = rf(id_rcp)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]objects.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id_rcp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateRole provides a mock function with given fields: login, role
func (_m *AccountsRep) UpdateRole(login string, role string) error {
	ret := _m.Called(login, role)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(login, role)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
