// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	objects "api/recipes/objects"

	mock "github.com/stretchr/testify/mock"
)

// RecipesRep is an autogenerated mock type for the RecipesRep type
type RecipesRep struct {
	mock.Mock
}

// AddGrade provides a mock function with given fields: id, login
func (_m *RecipesRep) AddGrade(id int, login string) error {
	ret := _m.Called(id, login)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, string) error); ok {
		r0 = rf(id, login)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: rcp
func (_m *RecipesRep) Create(rcp *objects.Recipe) error {
	ret := _m.Called(rcp)

	var r0 error
	if rf, ok := ret.Get(0).(func(*objects.Recipe) error); ok {
		r0 = rf(rcp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: id
func (_m *RecipesRep) Delete(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteGrade provides a mock function with given fields: id, login
func (_m *RecipesRep) DeleteGrade(id int, login string) error {
	ret := _m.Called(id, login)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, string) error); ok {
		r0 = rf(id, login)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindById provides a mock function with given fields: id
func (_m *RecipesRep) FindById(id int) (*objects.Recipe, error) {
	ret := _m.Called(id)

	var r0 *objects.Recipe
	if rf, ok := ret.Get(0).(func(int) *objects.Recipe); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*objects.Recipe)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByLogin provides a mock function with given fields: login
func (_m *RecipesRep) FindByLogin(login string) []objects.Recipe {
	ret := _m.Called(login)

	var r0 []objects.Recipe
	if rf, ok := ret.Get(0).(func(string) []objects.Recipe); ok {
		r0 = rf(login)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]objects.Recipe)
		}
	}

	return r0
}

// GetAmountGrades provides a mock function with given fields: id
func (_m *RecipesRep) GetAmountGrades(id int) int {
	ret := _m.Called(id)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetLikedByLogin provides a mock function with given fields: login
func (_m *RecipesRep) GetLikedByLogin(login string) ([]objects.Recipe, error) {
	ret := _m.Called(login)

	var r0 []objects.Recipe
	if rf, ok := ret.Get(0).(func(string) []objects.Recipe); ok {
		r0 = rf(login)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]objects.Recipe)
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

// List provides a mock function with given fields:
func (_m *RecipesRep) List() []objects.Recipe {
	ret := _m.Called()

	var r0 []objects.Recipe
	if rf, ok := ret.Get(0).(func() []objects.Recipe); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]objects.Recipe)
		}
	}

	return r0
}
