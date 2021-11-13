// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	objects "api/recipes/objects"

	mock "github.com/stretchr/testify/mock"
)

// CategoriesRep is an autogenerated mock type for the CategoriesRep type
type CategoriesRep struct {
	mock.Mock
}

// AddToRecipe provides a mock function with given fields: id_rcp, ctg
func (_m *CategoriesRep) AddToRecipe(id_rcp int, ctg string) error {
	ret := _m.Called(id_rcp, ctg)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, string) error); ok {
		r0 = rf(id_rcp, ctg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: obj
func (_m *CategoriesRep) Create(obj *objects.Category) error {
	ret := _m.Called(obj)

	var r0 error
	if rf, ok := ret.Get(0).(func(*objects.Category) error); ok {
		r0 = rf(obj)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateList provides a mock function with given fields: obj
func (_m *CategoriesRep) CreateList(obj []objects.Category) error {
	ret := _m.Called(obj)

	var r0 error
	if rf, ok := ret.Get(0).(func([]objects.Category) error); ok {
		r0 = rf(obj)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DelFromRecipe provides a mock function with given fields: id_rcp, ctg
func (_m *CategoriesRep) DelFromRecipe(id_rcp int, ctg string) error {
	ret := _m.Called(id_rcp, ctg)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, string) error); ok {
		r0 = rf(id_rcp, ctg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: ctg
func (_m *CategoriesRep) Find(ctg string) ([]objects.Category, error) {
	ret := _m.Called(ctg)

	var r0 []objects.Category
	if rf, ok := ret.Get(0).(func(string) []objects.Category); ok {
		r0 = rf(ctg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]objects.Category)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ctg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByRecipe provides a mock function with given fields: id_rcp
func (_m *CategoriesRep) FindByRecipe(id_rcp int) ([]objects.Category, error) {
	ret := _m.Called(id_rcp)

	var r0 []objects.Category
	if rf, ok := ret.Get(0).(func(int) []objects.Category); ok {
		r0 = rf(id_rcp)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]objects.Category)
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

// FindRecipes provides a mock function with given fields: ctg
func (_m *CategoriesRep) FindRecipes(ctg string) ([]objects.Recipe, error) {
	ret := _m.Called(ctg)

	var r0 []objects.Recipe
	if rf, ok := ret.Get(0).(func(string) []objects.Recipe); ok {
		r0 = rf(ctg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]objects.Recipe)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ctg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctg
func (_m *CategoriesRep) Get(ctg string) (*objects.Category, error) {
	ret := _m.Called(ctg)

	var r0 *objects.Category
	if rf, ok := ret.Get(0).(func(string) *objects.Category); ok {
		r0 = rf(ctg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*objects.Category)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ctg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields:
func (_m *CategoriesRep) List() []objects.Category {
	ret := _m.Called()

	var r0 []objects.Category
	if rf, ok := ret.Get(0).(func() []objects.Category); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]objects.Category)
		}
	}

	return r0
}

// ReplaceInRecipe provides a mock function with given fields: id_rcp, arr
func (_m *CategoriesRep) ReplaceInRecipe(id_rcp int, arr []objects.Category) error {
	ret := _m.Called(id_rcp, arr)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, []objects.Category) error); ok {
		r0 = rf(id_rcp, arr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
